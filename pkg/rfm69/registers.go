package rfm69

const (
	// Common Configuration Registers (Table 22)
	REG_FIFO           = 0x00 // FIFO read/write access
	REG_OP_MODE        = 0x01 // Operating modes of the transceiver
	REG_DATA_MODUL     = 0x02 // Data operation mode and modulation settings
	REG_BITRATE_MSB    = 0x03 // Bit Rate setting, most significant bits
	REG_BITRATE_LSB    = 0x04 // Bit Rate setting, least significant bits
	REG_FDEV_MSB       = 0x05 // Frequency deviation setting, MSB
	REG_FDEV_LSB       = 0x06 // Frequency deviation setting, LSB
	REG_FRF_MSB        = 0x07 // RF carrier frequency, MSB
	REG_FRF_MID        = 0x08 // RF carrier frequency, intermediate bits
	REG_FRF_LSB        = 0x09 // RF carrier frequency, LSB
	REG_OSC1           = 0x0A // RC oscillator settings
	REG_AFC_CTRL       = 0x0B // AFC control in low modulation index situations
	REG_LOW_BAT        = 0x0C // Low battery indicator settings
	REG_LISTEN1        = 0x0D // Listen mode settings
	REG_LISTEN2        = 0x0E // Listen mode idle duration
	REG_LISTEN3        = 0x0F // Listen mode Rx duration
	REG_VERSION        = 0x10 // Silicon revision/version code
	REG_PA_LEVEL       = 0x11 // PA selection and output power control
	REG_PA_RAMP        = 0x12 // PA ramp time control in FSK mode
	REG_OCP            = 0x13 // Over current protection control
	RESERVED_14        = 0x14 // Reserved
	RESERVED_15        = 0x15 // Reserved
	RESERVED_16        = 0x16 // Reserved
	RESERVED_17        = 0x17 // Reserved
	REG_LNA            = 0x18 // LNA settings
	REG_RX_BW          = 0x19 // Channel filter bandwidth control
	REG_AFC_BW         = 0x1A // Channel filter BW control during AFC routine
	REG_OOK_PEAK       = 0x1B // OOK demodulator selection and control in peak mode
	REG_OOK_AVG        = 0x1C // Average threshold control of the OOK demodulator
	REG_OOK_FIX        = 0x1D // Fixed threshold control of the OOK demodulator
	REG_AFC_FEI        = 0x1E // AFC and FEI control and status
	REG_AFC_MSB        = 0x1F // MSB of the AFC value (frequency correction)
	REG_AFC_LSB        = 0x20 // LSB of the AFC value (frequency correction)
	REG_FEI_MSB        = 0x21 // MSB of the calculated frequency error
	REG_FEI_LSB        = 0x22 // LSB of the calculated frequency error
	REG_RSSI_CONFIG    = 0x23 // RSSI-related settings
	REG_RSSI_VALUE     = 0x24 // RSSI value (in dBm, 0.5 dB steps)
	REG_DIO_MAPPING1   = 0x25 // Mapping of digital I/O pins DIO0–DIO3
	REG_DIO_MAPPING2   = 0x26 // Mapping of DIO4–DIO5 and ClkOut frequency selection
	REG_IRQ_FLAGS1     = 0x27 // IRQ flags register 1 (e.g. PLL lock, timeout, RSSI threshold)
	REG_IRQ_FLAGS2     = 0x28 // IRQ flags register 2 (e.g. FIFO flags, low battery detection)
	REG_RSSI_THRESH    = 0x29 // RSSI threshold control
	REG_RX_TIMEOUT1    = 0x2A // Rx timeout (between Rx request and RSSI detection)
	REG_RX_TIMEOUT2    = 0x2B // Rx timeout (between RSSI detection and PayloadReady)
	REG_PREAMBLE_MSB   = 0x2C // Preamble length, MSB
	REG_PREAMBLE_LSB   = 0x2D // Preamble length, LSB
	REG_SYNC_CONFIG    = 0x2E // Sync word recognition control
	REG_SYNC_VALUE1    = 0x2F // Sync word byte 1 (MSB)
	REG_SYNC_VALUE2    = 0x30 // Sync word byte 2
	REG_SYNC_VALUE3    = 0x31 // Sync word byte 3
	REG_SYNC_VALUE4    = 0x32 // Sync word byte 4
	REG_SYNC_VALUE5    = 0x33 // Sync word byte 5
	REG_SYNC_VALUE6    = 0x34 // Sync word byte 6
	REG_SYNC_VALUE7    = 0x35 // Sync word byte 7
	REG_SYNC_VALUE8    = 0x36 // Sync word byte 8
	REG_PACKET_CONFIG1 = 0x37 // Packet mode settings (packet format, CRC, address filtering, etc.)
	REG_PAYLOAD_LENGTH = 0x38 // Payload length setting (fixed or maximum in variable mode)
	REG_NODE_ADRS      = 0x39 // Node address for filtering
	REG_BROADCAST_ADRS = 0x3A // Broadcast address for filtering
	REG_AUTO_MODES     = 0x3B // Automatic mode settings
	REG_FIFO_THRESH    = 0x3C // FIFO threshold (Tx start condition)
	REG_PACKET_CONFIG2 = 0x3D // Additional packet mode settings

	// AES KEY REGISTERS (16 bytes, Table 22)
	REG_AES_KEY1  = 0x3E // AES key byte 1 (MSB)
	REG_AES_KEY2  = 0x3F // AES key byte 2
	REG_AES_KEY3  = 0x40 // AES key byte 3
	REG_AES_KEY4  = 0x41 // AES key byte 4
	REG_AES_KEY5  = 0x42 // AES key byte 5
	REG_AES_KEY6  = 0x43 // AES key byte 6
	REG_AES_KEY7  = 0x44 // AES key byte 7
	REG_AES_KEY8  = 0x45 // AES key byte 8
	REG_AES_KEY9  = 0x46 // AES key byte 9
	REG_AES_KEY10 = 0x47 // AES key byte 10
	REG_AES_KEY11 = 0x48 // AES key byte 11
	REG_AES_KEY12 = 0x49 // AES key byte 12
	REG_AES_KEY13 = 0x4A // AES key byte 13
	REG_AES_KEY14 = 0x4B // AES key byte 14
	REG_AES_KEY15 = 0x4C // AES key byte 15
	REG_AES_KEY16 = 0x4D // AES key byte 16 (LSB)

	REG_TEMP1 = 0x4E // Temperature sensor control
	REG_TEMP2 = 0x4F // Temperature sensor readout

	// TEST AND SPECIAL FUNCTION REGISTERS (Tables 28 and 29)
	REG_TEST_LNA   = 0x58 // Sensitivity boost control
	REG_TEST_TCXO  = 0x59 // XTAL or TCXO input selection
	REG_TEST_LL_BW = 0x5F // PLL bandwidth setting
	REG_TEST_DAGC  = 0x6F // Continuous DAGC control / fading margin improvement
	REG_TEST_AFC   = 0x71 // AFC offset for low modulation index AFC
)
