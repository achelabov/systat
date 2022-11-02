package widgets

import (
	"math"
	"sync"
	"time"

	models "github.com/achelabov/systat/server/models"
	"github.com/distatus/battery"
)

type BatteryWidget struct {
	batts          []*models.Battery
	mutex          *sync.Mutex
	updateInterval time.Duration
	battsLoad      chan []*models.Battery
}

func NewBatteryWidget() *BatteryWidget {
	battsAmount := battsCount()
	widget := &BatteryWidget{
		batts:          make([]*models.Battery, battsAmount),
		mutex:          new(sync.Mutex),
		updateInterval: time.Second,
		battsLoad:      make(chan []*models.Battery),
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
	out := make(chan []*models.Battery)

	go func() {
		defer close(out)
		for {
			select {
			case out <- <-b.battsLoad:
			case <-cancel:
				return
			}
		}
	}()

	return out
}

//receives batts one at a time until their number is less than the number of batteries
//func toBatts(ctx context.Context, batt <-chan *models.Battery) <-chan []*models.Battery {
//	battsAmount := battsCount()
//	batteries := make([]*models.Battery, battsAmount)
//	out := make(chan []*models.Battery)
//
//	go func() {
//		defer close(out)
//
//		idx := 0
//		for idx < battsAmount {
//			select {
//			case batt := <-batt:
//				batteries[idx] = batt
//				idx++
//			case <-ctx.Done():
//				return
//			}
//		}
//		out <- batteries
//	}()
//
//	return out
//}

func (b *BatteryWidget) update() {
	batteries := getBatteries()

	b.mutex.Lock()
	for i, battery := range batteries {
		b.batts[i].BatteryLoad = math.Abs(battery.Current/battery.Full) * 100.0
		b.batts[i].State = battery.State.String()
	}
	b.mutex.Unlock()

	b.battsLoad <- b.batts
}

func getBatteries() []*battery.Battery {
	batteries, _ := battery.GetAll()

	return batteries
}

func (b *BatteryWidget) BattsCount() int {
	return len(getBatteries())
}

func battsCount() int {
	return len(getBatteries())
}
