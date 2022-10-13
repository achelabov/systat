package models

type Battery struct {
	PercentFull float64
	State       string
}

type Cpu struct {
	CpuLoad float64
}

type Stats struct {
	Battery
	Cpu
}
