package models

type Battery struct {
	BatteryLoad float64
	State       string
}

type Cpu struct {
	CpuLoad float64
}

type Memory struct {
	Total       int64
	Avaliable   int64
	Used        int64
	UsedPercent int64
}
