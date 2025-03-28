// Package rfm69 provides a minimal TinyGo driver for the RFM69/SX1231 radio.
package rfm69

import (
	"errors"
	"machine"
	"time"
)

// Some register definitions (see datasheet for details).
const (
	REG_FIFO   = 0x00
	REG_OPMODE = 0x01
	REG_FRFMSB = 0x07
	REG_FRFMID = 0x08
	REG_FRFLSB = 0x09
	// Chip revision register; for many RFM69/SX1231 chips this is 0x10.
	REG_VERSION = 0x10
	// … add additional registers as needed.
)

// Operating modes
const (
	MODE_SLEEP = 0x00
	MODE_STDBY = 0x04
	MODE_TX    = 0x0C
	MODE_RX    = 0x10
)

// Expected chip revision value (adjust if needed)
const EXPECTED_REVISION byte = 0x24

// Device represents an RFM69/SX1231 radio connected via SPI.
type Device struct {
	spi      machine.SPI
	csPin    machine.Pin
	resetPin machine.Pin
	// Optionally add an interrupt pin or other pins if needed.
}

// New returns a new instance of the Device.
// It expects an SPI bus, a chip-select (CS) pin, and a reset pin.
func New(spi machine.SPI, cs, reset machine.Pin) *Device {
	cs.Configure(machine.PinConfig{Mode: machine.PinOutput})
	reset.Configure(machine.PinConfig{Mode: machine.PinOutput})
	return &Device{
		spi:      spi,
		csPin:    cs,
		resetPin: reset,
	}
}

// writeReg writes a value to a register.
func (d *Device) writeReg(reg, value byte) {
	d.csPin.Low()
	// For writes, set the MSB.
	d.spi.Transfer(reg | 0x80)
	d.spi.Transfer(value)
	d.csPin.High()
}

// readReg reads a value from a register.
func (d *Device) readReg(reg byte) byte {
	d.csPin.Low()
	// Clear the MSB for reading.
	d.spi.Transfer(reg & 0x7F)
	val, _ := d.spi.Transfer(0x00)
	d.csPin.High()
	return val
}

// reset performs a hardware reset of the module.
func (d *Device) reset() {
	d.resetPin.Low()
	time.Sleep(100 * time.Millisecond)
	d.resetPin.High()
	time.Sleep(100 * time.Millisecond)
}

// Init initializes the radio.
// It resets the device, sets it to standby mode, checks the chip revision,
// and applies initial configuration.
func (d *Device) Init() error {
	d.reset()
	// Place the radio into standby mode.
	d.writeReg(REG_OPMODE, MODE_STDBY)
	time.Sleep(100 * time.Millisecond)

	// Check the chip revision.
	rev := d.readReg(REG_VERSION)
	if rev != EXPECTED_REVISION {
		return errors.New("rfm69: chip revision mismatch")
	}

	// Additional configuration can be performed here.
	// For example: setting data mode, modulation, bandwidth, etc.
	// d.writeReg(REGISTER, value)

	return nil
}

// SetFrequency sets the radio frequency (in Hz).
// The RFM69 uses registers REG_FRFMSB, REG_FRFMID, and REG_FRFLSB.
// The frequency is calculated as: FRF = frequency / FSTEP, where FSTEP ~61.035 Hz.
func (d *Device) SetFrequency(freqHz uint32) {
	const fStep = 61.035
	frf := uint32(float64(freqHz) / fStep)
	d.writeReg(REG_FRFMSB, byte((frf>>16)&0xFF))
	d.writeReg(REG_FRFMID, byte((frf>>8)&0xFF))
	d.writeReg(REG_FRFLSB, byte(frf&0xFF))
}

// Send transmits the provided data.
// This is a simplified implementation – a complete version should handle
// interrupts and check for transmission completion via IRQ flags.
func (d *Device) Send(data []byte) error {
	// Go to standby before filling the FIFO.
	d.writeReg(REG_OPMODE, MODE_STDBY)
	time.Sleep(10 * time.Millisecond)

	// Write data to FIFO.
	d.csPin.Low()
	// Write the FIFO register address with write flag.
	d.spi.Transfer(REG_FIFO | 0x80)
	for _, b := range data {
		d.spi.Transfer(b)
	}
	d.csPin.High()

	// Switch to transmit mode.
	d.writeReg(REG_OPMODE, MODE_TX)
	// Wait for transmission to complete.
	// (Ideally, poll an IRQ register or wait for a TX done signal.)
	time.Sleep(100 * time.Millisecond)

	// Return to standby.
	d.writeReg(REG_OPMODE, MODE_STDBY)
	return nil
}

// Receive listens for an incoming packet.
// This blocking function is a very basic example.
// A complete implementation should use IRQs and proper packet handling.
func (d *Device) Receive() ([]byte, error) {
	// Set the radio to receive mode.
	d.writeReg(REG_OPMODE, MODE_RX)

	// Wait for a packet.
	// In a real implementation, poll an IRQ flag indicating payload ready.
	time.Sleep(500 * time.Millisecond)

	// For demonstration, read a fixed number of bytes from FIFO.
	// The actual payload length should be read from a dedicated register.
	const fifoSize = 66
	buf := make([]byte, fifoSize)
	d.csPin.Low()
	d.spi.Transfer(REG_FIFO & 0x7F)
	for i := 0; i < fifoSize; i++ {
		buf[i], _ = d.spi.Transfer(0x00)
	}
	d.csPin.High()

	// Return the buffer (in a full implementation, you may wish to trim to the actual length).
	return buf, nil
}
