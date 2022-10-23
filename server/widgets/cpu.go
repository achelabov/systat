package widgets

import (
	"context"
	"log"
	"sync"
	"time"

	models "github.com/achelabov/systat/server/models"
	cpuInf "github.com/shirou/gopsutil/cpu"
)

type CpuWidget struct {
	averageLoad    float64
	cpus           map[int]*models.Cpu
	mutex          *sync.Mutex
	updateInterval time.Duration
	resp           chan *models.Cpu
}

func NewCpuWidget() *CpuWidget {
	cpuCount, err := cpuInf.Counts(true)
	if err != nil {
		log.Fatal("can't get cpu count")
	}

	widget := &CpuWidget{
		cpus:           make(map[int]*models.Cpu, cpuCount),
		mutex:          new(sync.Mutex),
		updateInterval: time.Second * 2,
		resp:           make(chan *models.Cpu),
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

func (c *CpuWidget) GetCpus(ctx context.Context) <-chan []*models.Cpu {
	out := make(chan []*models.Cpu)

	go func() {
		defer close(out)
		for {
			select {
			case out <- <-toCpus(ctx, c.resp):
			case <-ctx.Done():
				return
			}
		}
	}()

	return out
}

func toCpus(ctx context.Context, cpu <-chan *models.Cpu) <-chan []*models.Cpu {
	cpusAmount := cpusCount()
	cpus := make([]*models.Cpu, cpusAmount)
	out := make(chan []*models.Cpu)

	go func() {
		defer close(out)

		idx := 0
		for idx < cpusAmount {
			select {
			case cpu := <-cpu:
				cpus[idx] = cpu
				idx++
			case <-ctx.Done():
				return
			}
		}
		out <- cpus
	}()

	return out
}

func (c *CpuWidget) GetAverage() float64 {
	c.mutex.Lock()
	avg := c.averageLoad
	c.mutex.Unlock()

	return avg
}

func (c *CpuWidget) update() {
	//average load
	go func() {
		cpuPercent, err := cpuInf.Percent(c.updateInterval, false)
		if err != nil {
			log.Fatal("can't get average cpu usage percent")
		}

		c.mutex.Lock()
		c.averageLoad = cpuPercent[0]
		c.mutex.Unlock()
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
		defer c.mutex.Unlock()

		for i, percent := range cpuPercent {
			c.cpus[i].CpuLoad = percent

			c.resp <- c.cpus[i]
		}
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
