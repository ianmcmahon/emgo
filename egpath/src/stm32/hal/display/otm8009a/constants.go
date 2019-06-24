// +build f469xx
package otm8009a

const (
	LCD_OTM8009A_ID        uint32 = 0
	OTM8009A_CMD_NOP       uint8  = 0x00 // NOP command
	OTM8009A_CMD_SWRESET          = 0x01 // Sw reset command
	OTM8009A_CMD_RDDMADCTL        = 0x0B // Read Display MADCTR command : read memory display access ctrl
	OTM8009A_CMD_RDDCOLMOD        = 0x0C // Read Display pixel format
	OTM8009A_CMD_SLPIN            = 0x10 // Sleep In command
	OTM8009A_CMD_SLPOUT           = 0x11 // Sleep Out command
	OTM8009A_CMD_PTLON            = 0x12 // Partial mode On command
	OTM8009A_CMD_DISPOFF          = 0x28 // Display Off command
	OTM8009A_CMD_DISPON           = 0x29 // Display On command
	OTM8009A_CMD_CASET            = 0x2A // Column address set command
	OTM8009A_CMD_PASET            = 0x2B // Page address set command
	OTM8009A_CMD_RAMWR            = 0x2C // Memory (GRAM) write command
	OTM8009A_CMD_RAMRD            = 0x2E // Memory (GRAM) read command
	OTM8009A_CMD_PLTAR            = 0x30 // Partial area command (4 parameters)
	OTM8009A_CMD_TEOFF            = 0x34 // Tearing Effect Line Off command : command with no parameter
	OTM8009A_CMD_TEEON            = 0x35 // Tearing Effect Line On command : command with 1 parameter 'TELOM'

	// Parameter TELOM : Tearing Effect Line Output Mode : possible values
	OTM8009A_TEEON_TELOM_VBLANKING_INFO_ONLY          = 0x00
	OTM8009A_TEEON_TELOM_VBLANKING_AND_HBLANKING_INFO = 0x01

	OTM8009A_CMD_MADCTR = 0x36 // Memory Access write control command

	// Possible used values of MADCTR
	OTM8009A_MADCTR_MODE_PORTRAIT  = 0x00
	OTM8009A_MADCTR_MODE_LANDSCAPE = 0x60 // MY = 0, MX = 1, MV = 1, ML = 0, RGB = 0

	OTM8009A_CMD_IDMOFF = 0x38 // Idle mode Off command
	OTM8009A_CMD_IDMON  = 0x39 // Idle mode On command

	OTM8009A_CMD_COLMOD = 0x3A // Interface Pixel format command

	// Possible values of COLMOD parameter corresponding to used pixel formats
	OTM8009A_COLMOD_RGB565 = 0x55
	OTM8009A_COLMOD_RGB888 = 0x77

	OTM8009A_CMD_RAMWRC = 0x3C // Memory write continue command
	OTM8009A_CMD_RAMRDC = 0x3E // Memory read continue command

	OTM8009A_CMD_WRTESCN = 0x44 // Write Tearing Effect Scan line command
	OTM8009A_CMD_RDSCNL  = 0x45 // Read  Tearing Effect Scan line command

	// CABC Management : ie : Content Adaptive Back light Control in IC OTM8009a
	OTM8009A_CMD_WRDISBV  = 0x51 // Write Display Brightness command
	OTM8009A_CMD_WRCTRLD  = 0x53 // Write CTRL Display command
	OTM8009A_CMD_WRCABC   = 0x55 // Write Content Adaptive Brightness command
	OTM8009A_CMD_WRCABCMB = 0x5E // Write CABC Minimum Brightness command

	OTM8009A_ORIENTATION_PORTRAIT  uint32 = 0x00 // Portrait orientation choice of LCD screen
	OTM8009A_ORIENTATION_LANDSCAPE        = 0x01 // Landscape orientation choice of LCD screen

	OTM8009A_FORMAT_RGB888 uint32 = 0x00 // Pixel format chosen is RGB888 : 24 bpp
	OTM8009A_FORMAT_RBG565        = 0x02 // Pixel format chosen is RGB565 : 16 bpp

	// Width and Height in Portrait mode
	OTM8009A_480X800_WIDTH  uint16 = 480 // LCD PIXEL WIDTH
	OTM8009A_480X800_HEIGHT        = 800 // LCD PIXEL HEIGHT

	// Width and Height in Landscape mode
	OTM8009A_800X480_WIDTH  uint16 = 800 // LCD PIXEL WIDTH
	OTM8009A_800X480_HEIGHT        = 480 // LCD PIXEL HEIGHT

)

//emgo:const
var (
	// Constant tables of register settings used to transmit DSI command packets as
	// power up initialization sequence of the KoD LCD (OTM8009A LCD Driver)
	// I believe the last byte of each array is the command, and actually needs to be sent first
	lcdRegData1  = []byte{0x80, 0x09, 0x01, 0xFF}
	lcdRegData2  = []byte{0x80, 0x09, 0xFF}
	lcdRegData3  = []byte{0x00, 0x09, 0x0F, 0x0E, 0x07, 0x10, 0x0B, 0x0A, 0x04, 0x07, 0x0B, 0x08, 0x0F, 0x10, 0x0A, 0x01, 0xE1}
	lcdRegData4  = []byte{0x00, 0x09, 0x0F, 0x0E, 0x07, 0x10, 0x0B, 0x0A, 0x04, 0x07, 0x0B, 0x08, 0x0F, 0x10, 0x0A, 0x01, 0xE2}
	lcdRegData5  = []byte{0x79, 0x79, 0xD8}
	lcdRegData6  = []byte{0x00, 0x01, 0xB3}
	lcdRegData7  = []byte{0x85, 0x01, 0x00, 0x84, 0x01, 0x00, 0xCE}
	lcdRegData8  = []byte{0x18, 0x04, 0x03, 0x39, 0x00, 0x00, 0x00, 0x18, 0x03, 0x03, 0x3A, 0x00, 0x00, 0x00, 0xCE}
	lcdRegData9  = []byte{0x18, 0x02, 0x03, 0x3B, 0x00, 0x00, 0x00, 0x18, 0x01, 0x03, 0x3C, 0x00, 0x00, 0x00, 0xCE}
	lcdRegData10 = []byte{0x01, 0x01, 0x20, 0x20, 0x00, 0x00, 0x01, 0x02, 0x00, 0x00, 0xCF}
	lcdRegData11 = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xCB}
	lcdRegData12 = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xCB}
	lcdRegData13 = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xCB}
	lcdRegData14 = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xCB}
	lcdRegData15 = []byte{0x00, 0x04, 0x04, 0x04, 0x04, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xCB}
	lcdRegData16 = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04, 0x04, 0x04, 0x04, 0x04, 0x00, 0x00, 0x00, 0x00, 0xCB}
	lcdRegData17 = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xCB}
	lcdRegData18 = []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xCB}
	lcdRegData19 = []byte{0x00, 0x26, 0x09, 0x0B, 0x01, 0x25, 0x00, 0x00, 0x00, 0x00, 0xCC}
	lcdRegData20 = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x26, 0x0A, 0x0C, 0x02, 0xCC}
	lcdRegData21 = []byte{0x25, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xCC}
	lcdRegData22 = []byte{0x00, 0x25, 0x0C, 0x0A, 0x02, 0x26, 0x00, 0x00, 0x00, 0x00, 0xCC}
	lcdRegData23 = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x25, 0x0B, 0x09, 0x01, 0xCC}
	lcdRegData24 = []byte{0x26, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xCC}
	lcdRegData25 = []byte{0xFF, 0xFF, 0xFF, 0xFF}

	// CASET value (Column Address Set) : X direction LCD GRAM boundaries
	// depending on LCD orientation mode and PASET value (Page Address Set) : Y direction
	// LCD GRAM boundaries depending on LCD orientation mode
	// XS[15:0] = 0x000 = 0, XE[15:0] = 0x31F = 799 for landscape mode : apply to CASET
	// YS[15:0] = 0x000 = 0, YE[15:0] = 0x31F = 799 for portrait mode : : apply to PASET
	lcdRegData27 = []byte{0x00, 0x00, 0x03, 0x1F, OTM8009A_CMD_CASET}
	// XS[15:0] = 0x000 = 0, XE[15:0] = 0x1DF = 479 for portrait mode : apply to CASET
	// YS[15:0] = 0x000 = 0, YE[15:0] = 0x1DF = 479 for landscape mode : apply to PASET
	lcdRegData28 = []byte{0x00, 0x00, 0x01, 0xDF, OTM8009A_CMD_PASET}

	ShortRegData1  = []byte{OTM8009A_CMD_NOP, 0x00}
	ShortRegData2  = []byte{OTM8009A_CMD_NOP, 0x80}
	ShortRegData3  = []byte{0xC4, 0x30}
	ShortRegData4  = []byte{OTM8009A_CMD_NOP, 0x8A}
	ShortRegData5  = []byte{0xC4, 0x40}
	ShortRegData6  = []byte{OTM8009A_CMD_NOP, 0xB1}
	ShortRegData7  = []byte{0xC5, 0xA9}
	ShortRegData8  = []byte{OTM8009A_CMD_NOP, 0x91}
	ShortRegData9  = []byte{0xC5, 0x34}
	ShortRegData10 = []byte{OTM8009A_CMD_NOP, 0xB4}
	ShortRegData11 = []byte{0xC0, 0x50}
	ShortRegData12 = []byte{0xD9, 0x4E}
	ShortRegData13 = []byte{OTM8009A_CMD_NOP, 0x81}
	ShortRegData14 = []byte{0xC1, 0x66}
	ShortRegData15 = []byte{OTM8009A_CMD_NOP, 0xA1}
	ShortRegData16 = []byte{0xC1, 0x08}
	ShortRegData17 = []byte{OTM8009A_CMD_NOP, 0x92}
	ShortRegData18 = []byte{0xC5, 0x01}
	ShortRegData19 = []byte{OTM8009A_CMD_NOP, 0x95}
	ShortRegData20 = []byte{OTM8009A_CMD_NOP, 0x94}
	ShortRegData21 = []byte{0xC5, 0x33}
	ShortRegData22 = []byte{OTM8009A_CMD_NOP, 0xA3}
	ShortRegData23 = []byte{0xC0, 0x1B}
	ShortRegData24 = []byte{OTM8009A_CMD_NOP, 0x82}
	ShortRegData25 = []byte{0xC5, 0x83}
	ShortRegData26 = []byte{0xC4, 0x83}
	ShortRegData27 = []byte{0xC1, 0x0E}
	ShortRegData28 = []byte{OTM8009A_CMD_NOP, 0xA6}
	ShortRegData29 = []byte{OTM8009A_CMD_NOP, 0xA0}
	ShortRegData30 = []byte{OTM8009A_CMD_NOP, 0xB0}
	ShortRegData31 = []byte{OTM8009A_CMD_NOP, 0xC0}
	ShortRegData32 = []byte{OTM8009A_CMD_NOP, 0xD0}
	ShortRegData33 = []byte{OTM8009A_CMD_NOP, 0x90}
	ShortRegData34 = []byte{OTM8009A_CMD_NOP, 0xE0}
	ShortRegData35 = []byte{OTM8009A_CMD_NOP, 0xF0}
	ShortRegData36 = []byte{OTM8009A_CMD_SLPOUT, 0x00}
	ShortRegData37 = []byte{OTM8009A_CMD_COLMOD, OTM8009A_COLMOD_RGB565}
	ShortRegData38 = []byte{OTM8009A_CMD_COLMOD, OTM8009A_COLMOD_RGB888}
	ShortRegData39 = []byte{OTM8009A_CMD_MADCTR, OTM8009A_MADCTR_MODE_LANDSCAPE}
	ShortRegData40 = []byte{OTM8009A_CMD_WRDISBV, 0x7F}
	ShortRegData41 = []byte{OTM8009A_CMD_WRCTRLD, 0x2C}
	ShortRegData42 = []byte{OTM8009A_CMD_WRCABC, 0x02}
	ShortRegData43 = []byte{OTM8009A_CMD_WRCABCMB, 0xFF}
	ShortRegData44 = []byte{OTM8009A_CMD_DISPON, 0x00}
	ShortRegData45 = []byte{OTM8009A_CMD_RAMWR, 0x00}
	ShortRegData46 = []byte{0xCF, 0x00}
	ShortRegData47 = []byte{0xC5, 0x66}
	ShortRegData48 = []byte{OTM8009A_CMD_NOP, 0xB6}
	ShortRegData49 = []byte{0xF5, 0x06}
	ShortRegData50 = []byte{OTM8009A_CMD_NOP, 0xB1}
	ShortRegData51 = []byte{0xC6, 0x06}
)
