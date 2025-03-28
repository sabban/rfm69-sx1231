package main

import (
	"machine"
	"time"

	"github.com/sabban/rfm69-sx1231/pkg/rfm69" // Replace with the actual module path
)

func main() {
	// Configure SPI
	spi := machine.SPI0
	spi.Configure(machine.SPIConfig{
		Frequency: 8e6,
		SCK:       machine.SPI0_SCK_PIN,
		SDO:       machine.SPI0_SDO_PIN,
		SDI:       machine.SPI0_SDI_PIN,
	})

	// Define control pins (adjust to your board)
	csPin := machine.D10   // Chip select
	resetPin := machine.D9 // Reset pin

	// Create a new radio instance
	radio := rfm69.New(spi, csPin, resetPin)

	// Initialize the radio (includes chip revision check)
	if err := radio.Init(); err != nil {
		println("Failed to initialize radio:", err.Error())
		// Optionally, halt execution
		for {
			time.Sleep(time.Second)
		}
	}

	// Set frequency to 915 MHz (adjust if needed)
	radio.SetFrequency(915000000)

	println("Radio initialized successfully!")

	// Main loop: add further functionality such as sending or receiving data
	for {
		// For example, you could call radio.Send([]byte("Hello TinyGo!"))
		time.Sleep(time.Second)
	}
}
