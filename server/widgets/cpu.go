package widgets

import (
	"log"
	"sync"
	"time"

	cpuInf "github.com/shirou/gopsutil/cpu"
)

type CpuWidget struct {
	averageLoad    float64
	cpus           map[int]*Cpu
	mutex          *sync.Mutex
	updateInterval time.Duration
}

func NewCpuWidget() *CpuWidget {
	cpuCount, err := cpuInf.Counts(true)
	if err != nil {
		log.Fatal("can't get cpu count")
	}

	widget := &CpuWidget{
		cpus:           make(map[int]*Cpu),
		mutex:          new(sync.Mutex),
		updateInterval: time.Second,
	}

	for i := 0; i < cpuCount; i++ {
		widget.cpus[i] = new(Cpu)
	}

	go func() {
		for range time.NewTicker(widget.updateInterval).C {
			widget.update()
		}
	}()

	return widget
}

func (c *CpuWidget) GetCpus() []*Cpu {
	cpus := make([]*Cpu, 0)

	c.mutex.Lock()
	for _, value := range c.cpus {
		cpus = append(cpus, value)
	}
	c.mutex.Unlock()

	return cpus
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
		}
	}()
}
