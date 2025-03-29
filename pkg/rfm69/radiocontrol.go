package rfm69

type RadioController interface {
	Init() error
	SetRfSwitchMode(mode int) error
	SetNss(state bool) error
	WaitWhileBusy() error
	SetupInterrupts(handler func()) error
}
