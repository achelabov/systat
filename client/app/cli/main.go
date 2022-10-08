package main

import (
	"fmt"
	"time"

	"github.com/achelabov/systat/server/widgets"
)

func main() {
	battWidget := widgets.NewBatteryWidget()

	for tick := range time.Tick(2 * time.Second) {
		for _, v := range battWidget.GetBatteries() {
			fmt.Println(v)
		}

		_ = tick
	}
}
