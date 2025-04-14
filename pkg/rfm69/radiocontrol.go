package rfm69

type RadioController interface {
	Init() error
	//SetRfSwitchMode(mode int) error
	SetCs(state bool) error
	//WaitWhileBusy() error
	SetupInterrupts(handler func()) error
}
