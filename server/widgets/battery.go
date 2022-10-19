package widgets

import (
	"log"
	"math"
	"sync"
	"time"

	models "github.com/achelabov/systat/server/models"
	"github.com/distatus/battery"
)

type BatteryWidget struct {
	batts          map[int]*models.Battery
	mutex          *sync.Mutex
	updateInterval time.Duration
	resp           chan *models.Battery
}

func NewBatteryWidget() *BatteryWidget {
	widget := &BatteryWidget{
		batts:          make(map[int]*models.Battery),
		mutex:          new(sync.Mutex),
		updateInterval: time.Second,
		resp:           make(chan *models.Battery, len(getBatteries())),
	}

	for i := range getBatteries() {
		widget.batts[i] = new(models.Battery)
	}

	go func() {
		for range time.NewTicker(widget.updateInterval).C {
			widget.update()
		}
	}()

	return widget
}

func (b *BatteryWidget) GetBatteries(cancel <-chan struct{}) <-chan []*models.Battery {
	//batteries := make([]*models.Battery, 0)

	//b.mutex.Lock()
	//for _, value := range b.batts {
	//	batteries = append(batteries, value)
	//}
	//b.mutex.Unlock()
	out := make(chan []*models.Battery)

	go func() {
		defer close(out)
		for {
			select {
			case out <- <-toBatts(cancel, b.resp):
			case <-cancel:
				return
			}
		}
	}()

	return out
}

func toBatts(cancel <-chan struct{}, batt <-chan *models.Battery) <-chan []*models.Battery {
	battAmount := len(getBatteries())
	batteries := make([]*models.Battery, battAmount)
	out := make(chan []*models.Battery)

	go func() {
		defer close(out)

		idx := 0
		for idx < battAmount {
			select {
			case batt := <-batt:
				batteries[idx] = batt
				//				fmt.Println("idx: ", idx, "batt: ", batt, "amount: ", battAmount)
				idx++
			case <-cancel:
				return
			}
		}
		out <- batteries
	}()

	return out
}

func (b *BatteryWidget) update() {
	batteries := getBatteries()

	b.mutex.Lock()
	defer b.mutex.Unlock()

	for i, battery := range batteries {
		b.batts[i].BatteryLoad = math.Abs(battery.Current/battery.Full) * 100.0
		b.batts[i].State = battery.State.String()

		b.resp <- b.batts[i]
	}
}

func getBatteries() []*battery.Battery {
	batteries, err := battery.GetAll()
	if err != nil {
		log.Fatal(err)
	}

	return batteries
}
