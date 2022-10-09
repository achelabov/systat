package main

import (
	"fmt"
	"time"

	"github.com/achelabov/systat/server/widgets"
)

func main() {
	battWidget := widgets.NewBatteryWidget()
	cpuWidget := widgets.NewCpuWidget()

	for range time.Tick(2 * time.Second) {
		fmt.Println("--------------------------------------")
		for i, v := range battWidget.GetBatteries() {
			fmt.Println("batt id: ", i, "load: ", v.PercentFull, "state: ", v.State)
		}
		fmt.Println("--------------------------------------")
		for i, v := range cpuWidget.GetCpus() {
			fmt.Println("cpu id: ", i, "load: ", v.CpuLoad)
		}
		fmt.Println("**************************************")
		fmt.Println("average load: ", cpuWidget.GetAverage())
		fmt.Println("--------------------------------------")
	}
}
