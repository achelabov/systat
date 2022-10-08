package widgets

import (
	"log"
	"math"
	"sync"
	"time"

	"github.com/distatus/battery"
)

type BatteryWidget struct {
	batt           map[int]*Battery
	mutex          *sync.Mutex
	updateInterval time.Duration
}

func NewBatteryWidget() *BatteryWidget {
	widget := &BatteryWidget{
		batt:           make(map[int]*Battery),
		mutex:          new(sync.Mutex),
		updateInterval: time.Second,
	}

	for i := range getBatteries() {
		widget.batt[i] = new(Battery)
	}

	go func() {
		for range time.NewTicker(widget.updateInterval).C {
			widget.mutex.Lock()
			widget.update()
			widget.mutex.Unlock()
		}
	}()

	return widget
}

func (b *BatteryWidget) GetBatteries() []*Battery {
	batteries := make([]*Battery, 0)

	b.mutex.Lock()
	for _, value := range b.batt {
		batteries = append(batteries, value)
	}
	b.mutex.Unlock()

	return batteries
}

func (b *BatteryWidget) update() {
	batteries := getBatteries()

	for i, battery := range batteries {
		b.batt[i].PercentFull = math.Abs(battery.Current/battery.Full) * 100.0
		b.batt[i].State = battery.State.String()
	}
}

func getBatteries() []*battery.Battery {
	batteries, err := battery.GetAll()
	if err != nil {
		log.Fatal(err)
	}

	return batteries
}
