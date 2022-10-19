package widgets

type Widgets struct {
	BatteryWidget
	CpuWidget
}

func NewWidgets() *Widgets {
	return &Widgets{
		BatteryWidget: *NewBatteryWidget(),
		CpuWidget:     *NewCpuWidget(),
	}
}

func (w *Widgets) GetWidgets() {

}
