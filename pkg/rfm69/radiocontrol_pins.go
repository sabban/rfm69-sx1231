package rfm69

import (
	"machine"
)

// RadioControl for boards that are connected using normal pins.
type RadioControl struct {
	csPin, rstPin, dio0Pin machine.Pin
}

func NewRadioControl(csPin, dio0Pin machine.Pin) *RadioControl {
	return &RadioControl{
		csPin:   csPin,
		dio0Pin: dio0Pin,
	}
}

// SetNss sets the NSS line aka chip select for SPI.
func (rc *RadioControl) SetCs(state bool) error {
	rc.csPin.Set(state)
	return nil
}

// Init() configures whatever needed for sx127x radio control
func (rc *RadioControl) Init() error {
	rc.csPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	rc.dio0Pin.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	return nil
}

// add interrupt handlers for Radio IRQs for pins
func (rc *RadioControl) SetupInterrupts(handler func()) error {
	irqHandler = handler

	// Setup DIO0 interrupt Handling
	if err := rc.dio0Pin.SetInterrupt(machine.PinRising, handleInterrupt); err != nil {
		return err
	}

	return nil
}

var irqHandler func()

func handleInterrupt(machine.Pin) {
	irqHandler()
}
