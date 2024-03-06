package videocontrol


type ControlEvent struct{}

func (e *ControlEvent) Run() {
}

func (e *ControlEvent) Spec() string {
	return "0 0 0 0 0 ?"
}