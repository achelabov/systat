package widgets

import (
	"context"
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
	battsAmount := battsCount()
	widget := &BatteryWidget{
		batts:          make(map[int]*models.Battery, battsAmount),
		mutex:          new(sync.Mutex),
		updateInterval: time.Second,
		resp:           make(chan *models.Battery, battsAmount),
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

func (b *BatteryWidget) GetBatteries(ctx context.Context) <-chan []*models.Battery {
	out := make(chan []*models.Battery)

	go func() {
		defer close(out)
		for {
			select {
			case out <- <-toBatts(ctx, b.resp):
			case <-ctx.Done():
				return
			}
		}
	}()

	return out
}

//receives batts one at a time until their number is less than the number of batteries
func toBatts(ctx context.Context, batt <-chan *models.Battery) <-chan []*models.Battery {
	battsAmount := battsCount()
	batteries := make([]*models.Battery, battsAmount)
	out := make(chan []*models.Battery)

	go func() {
		defer close(out)

		idx := 0
		for idx < battsAmount {
			select {
			case batt := <-batt:
				batteries[idx] = batt
				idx++
			case <-ctx.Done():
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

func (b *BatteryWidget) BattsCount() int {
	return len(getBatteries())
}

func battsCount() int {
	return len(getBatteries())
}
