package main

import (
	"machine"
	"time"

	"github.com/sabban/rfm69-sx1231/pkg/rfm69" // Replace with the actual module path
)

func main() {
	// Configure SPI
	time.Sleep(3 * time.Second)
	spi := machine.SPI0
	spi.Configure(machine.SPIConfig{
		Frequency: 8e6,
		SCK:       machine.SPI0_SCK_PIN,
		SDO:       machine.SPI0_SDO_PIN,
		SDI:       machine.SPI0_SDI_PIN,
	})

	// Define control pins (adjust to your board)
	csPin := machine.Pin(16)  // Chip select
	rstPin := machine.Pin(17) // Reset pin
	io0Pin := machine.Pin(21) // Interrupt pin

	radioControl := rfm69.NewRadioControl(csPin, io0Pin)
	radioControl.Init()

	// Create a new radio instance
	radio := rfm69.New(spi, rstPin, radioControl)

	radio.Reset()

	err := radio.SetStandbyMode()
	if err != nil {
		println("Failed to set standby mode: ", err)
	}

	err = radio.IsReady()
	if err != nil {
		println("Radio is not ready: ", err)

	} else {
		println("Radio initialized successfully!")
	}

	if radio.DetectDevice() {
		println("Device detected!")
	}

	// Main loop: add further functionality such as sending or receiving data
	for {
		// For example, you could call radio.Send([]byte("Hello TinyGo!"))
		time.Sleep(time.Second)
	}
}
