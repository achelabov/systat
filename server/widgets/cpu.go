package widgets

import (
	"log"
	"sync"
	"time"

	models "github.com/achelabov/systat/server/models"
	cpuInf "github.com/shirou/gopsutil/cpu"
)

type CpuWidget struct {
	cpus           []*models.Cpu
	mutex          *sync.Mutex
	updateInterval time.Duration
	cpusLoad       chan []*models.Cpu
	avgLoad        chan float64
}

func NewCpuWidget() *CpuWidget {
	cpuCount, err := cpuInf.Counts(true)
	if err != nil {
		log.Fatal("can't get cpu count")
	}

	widget := &CpuWidget{
		cpus:           make([]*models.Cpu, cpuCount),
		mutex:          new(sync.Mutex),
		updateInterval: time.Second,
		cpusLoad:       make(chan []*models.Cpu),
		avgLoad:        make(chan float64),
	}

	for i := 0; i < cpuCount; i++ {
		widget.cpus[i] = new(models.Cpu)
	}

	go func() {
		for range time.NewTicker(widget.updateInterval).C {
			widget.update()
		}
	}()

	return widget
}

func (c *CpuWidget) GetCpus(cancel <-chan struct{}) <-chan []*models.Cpu {
	out := make(chan []*models.Cpu)

	go func() {
		defer close(out)
		for {
			select {
			case out <- <-c.cpusLoad:
			case <-cancel:
				return
			}
		}
	}()

	return out
}

//receives cpus one at a time as long as their number is less than the number of processor cores
//func toCpus(ctx context.Context, cpu <-chan *models.Cpu) <-chan []*models.Cpu {
//	cpusAmount := cpusCount()
//	cpus := make([]*models.Cpu, cpusAmount)
//	out := make(chan []*models.Cpu)
//
//	go func() {
//		defer close(out)
//
//		for idx := 0; idx < cpusAmount; idx++ {
//			select {
//			case cpu := <-cpu:
//				cpus[idx] = cpu
//			case <-ctx.Done():
//				return
//			}
//		}
//		out <- cpus
//	}()
//
//	return out
//}

func (c *CpuWidget) GetAverageLoad(cancel <-chan struct{}) <-chan float64 {
	out := make(chan float64)

	go func() {
		defer close(out)
		for {
			select {
			case out <- <-c.avgLoad:
			case <-cancel:
				return
			}
		}
	}()

	return out
}

func (c *CpuWidget) update() {
	//average load
	go func() {
		cpuPercent, err := cpuInf.Percent(c.updateInterval, false)
		if err != nil {
			log.Fatal("can't get average cpu usage percent")
		}

		c.avgLoad <- cpuPercent[0]
	}()

	//load per cpu
	go func() {
		cpuPercent, err := cpuInf.Percent(c.updateInterval, true)
		if err != nil {
			log.Fatal("can't get load per cpu")
		}

		if len(cpuPercent) != len(c.cpus) {
			log.Fatal("number of cpu usage percents doesn't match count")
		}

		c.mutex.Lock()
		for i, percent := range cpuPercent {
			c.cpus[i].CpuLoad = percent
		}
		c.mutex.Unlock()

		c.cpusLoad <- c.cpus
	}()
}

func (c *CpuWidget) CpusCount() int {
	return cpusCount()
}

func cpusCount() int {
	cnt, err := cpuInf.Counts(true)
	if err != nil {
		log.Fatal("can't get cpu count")
	}
	return cnt
}
