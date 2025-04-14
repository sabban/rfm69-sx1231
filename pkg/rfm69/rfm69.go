// Package rfm69 provides a minimal TinyGo driver for the RFM69/SX1231 radio.
package rfm69

import (
	"errors"
	"machine"
	"time"

	"tinygo.org/x/drivers"
)

// Operating modes
const (
	MODE_SLEEP = 0x00
	MODE_STDBY = 0x04
	MODE_TX    = 0x0C
	MODE_RX    = 0x10

	MODE_READY_BIT = 0x80
)

const (
	PERIOD_PER_SEC      = (uint32)(1000000 / 15.625) // SX1261 DS 13.1.4
	SPI_BUFFER_SIZE     = 256
	RADIOEVENTCHAN_SIZE = 1

	timeout = 10000 * time.Millisecond
)

// Expected chip revision value (adjust if needed)
const EXPECTED_REVISION byte = 0x24

// Device represents an RFM69/SX1231 radio connected via SPI.
type Device struct {
	spi            drivers.SPI
	rstPin         machine.Pin
	radioEventChan chan uint8
	// Optionally add an interrupt pin or other pins if needed.
	controller RadioController // to manage interactions with the radio
	deepSleep  bool            // Internal Sleep state
	deviceType int             // sx1261,sx1262,sx1268 (defaults sx1261)
	spiTxBuf   []byte          // global Tx buffer to avoid heap allocations in interrupt
	spiRxBuf   []byte          // global Rx buffer to avoid heap allocations in interrupt
}

func New(spi drivers.SPI, resetPin machine.Pin, radio RadioController) *Device {
	return &Device{
		spi:            spi,
		rstPin:         resetPin,
		radioEventChan: make(chan uint8, RADIOEVENTCHAN_SIZE),
		spiTxBuf:       make([]byte, SPI_BUFFER_SIZE),
		spiRxBuf:       make([]byte, SPI_BUFFER_SIZE),
		controller:     radio,
	}
}

func (d *Device) Reset() {
	d.rstPin.Low()
	time.Sleep(100 * time.Millisecond)
	d.rstPin.High()
	time.Sleep(100 * time.Millisecond)
}

func (d *Device) SetStandbyMode() error {
	// Set to standby mode
	if err := d.WriteRegister(REG_OP_MODE, MODE_STDBY); err != nil {
		return errors.New("failed to set standby mode")
	}
	return nil
}

// WaitForChipReady polls the IRQ Flags register (RegIrqFlags1)
// until the ModeReady flag (bit 7) is set, indicating the chip is ready.
func (d *Device) IsReady() error {
	start := time.Now()
	for {
		value, err := d.ReadRegister(REG_IRQ_FLAGS1)
		if err != nil {
			return err
		}
		if value&MODE_READY_BIT != 0 {
			return nil // Chip is ready
		}
		if time.Since(start) > timeout {
			return errors.New("timeout waiting for chip to be ready")
		}
		time.Sleep(1 * time.Millisecond)
	}
}

func (d *Device) ReadRegister(address byte) (byte, error) {
	d.controller.SetCs(false)
	defer d.controller.SetCs(true)
	ret, err := d.spi.Transfer(address & 0x7F)
	if err != nil {
		return 0, err
	}
	return ret, nil
}

func (d *Device) DetectDevice() bool {
	version, err := d.ReadRegister(REG_VERSION)
	if version != EXPECTED_REVISION || err != nil {
		return false
	}
	return true
}

func (d *Device) WriteRegister(addr, value byte) error {
	d.controller.SetCs(false)
	defer d.controller.SetCs(true)
	_, err := d.spi.Transfer(addr | 0x80)
	if err != nil {
		return err
	}
	_, err = d.spi.Transfer(value)
	if err != nil {
		return err
	}
	return nil
}
