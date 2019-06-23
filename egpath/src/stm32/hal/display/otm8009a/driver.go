package otm8009a

import "github.com/ianmcmahon/emgo/egroot/src/delay"

type LCD struct {
}

func (p *LCD) Init() {
	// Enable CMD2 to access vendor specific commands
	// Enter in command 2 mode and set EXTC to enable address shift function (0x00)
	p.WriteCmd(0, ShortRegData1)
	p.WriteCmd(3, lcdRegData1)

	// Enter ORISE Command 2
	p.WriteCmd(0, ShortRegData2) // Shift address to 0x80
	p.WriteCmd(2, lcdRegData2)

	/////////////////////////////////////////////////////////////////////
	// SD_PCH_CTRL - 0xC480h - 129th parameter - Default 0x00
	// Set SD_PT
	// -> Source output level during porch and non-display area to GND
	p.WriteCmd(0, ShortRegData2)
	p.WriteCmd(0, ShortRegData3)
	delay.Millisec(10)
	// Not documented
	p.WriteCmd(0, ShortRegData4)
	p.WriteCmd(0, ShortRegData5)
	delay.Millisec(10)
	/////////////////////////////////////////////////////////////////////

	// PWR_CTRL4 - 0xC4B0h - 178th parameter - Default 0xA8
	// Set gvdd_en_test
	// -> enable GVDD test mode !!!
	p.WriteCmd(0, ShortRegData6)
	p.WriteCmd(0, ShortRegData7)

	// PWR_CTRL2 - 0xC590h - 146th parameter - Default 0x79
	// Set pump 4 vgh voltage
	// -> from 15.0v down to 13.0v
	// Set pump 5 vgh voltage
	// -> from -12.0v downto -9.0v
	p.WriteCmd(0, ShortRegData8)
	p.WriteCmd(0, ShortRegData9)

	// P_DRV_M - 0xC0B4h - 181th parameter - Default 0x00
	// -> Column inversion
	p.WriteCmd(0, ShortRegData10)
	p.WriteCmd(0, ShortRegData11)

	// VCOMDC - 0xD900h - 1st parameter - Default 0x39h
	// VCOM Voltage settings
	// -> from -1.0000v downto -1.2625v
	p.WriteCmd(0, ShortRegData1)
	p.WriteCmd(0, ShortRegData12)

	// Oscillator adjustment for Idle/Normal mode (LPDT only) set to 65Hz (default is 60Hz)
	p.WriteCmd(0, ShortRegData13)
	p.WriteCmd(0, ShortRegData14)

	// Video mode internal
	p.WriteCmd(0, ShortRegData15)
	p.WriteCmd(0, ShortRegData16)

	// PWR_CTRL2 - 0xC590h - 147h parameter - Default 0x00
	// Set pump 4&5 x6
	// -> ONLY VALID when PUMP4_EN_ASDM_HV = "0"
	p.WriteCmd(0, ShortRegData17)
	p.WriteCmd(0, ShortRegData18)

	// PWR_CTRL2 - 0xC590h - 150th parameter - Default 0x33h
	// Change pump4 clock ratio
	// -> from 1 line to 1/2 line
	p.WriteCmd(0, ShortRegData19)
	p.WriteCmd(0, ShortRegData9)

	// GVDD/NGVDD settings
	p.WriteCmd(0, ShortRegData1)
	p.WriteCmd(2, lcdRegData5)

	// PWR_CTRL2 - 0xC590h - 149th parameter - Default 0x33h
	// Rewrite the default value !
	p.WriteCmd(0, ShortRegData20)
	p.WriteCmd(0, ShortRegData21)

	// Panel display timing Setting 3
	p.WriteCmd(0, ShortRegData22)
	p.WriteCmd(0, ShortRegData23)

	// Power control 1
	p.WriteCmd(0, ShortRegData24)
	p.WriteCmd(0, ShortRegData25)

	// Source driver precharge
	p.WriteCmd(0, ShortRegData13)
	p.WriteCmd(0, ShortRegData26)

	p.WriteCmd(0, ShortRegData15)
	p.WriteCmd(0, ShortRegData27)

	p.WriteCmd(0, ShortRegData28)
	p.WriteCmd(2, lcdRegData6)

	// GOAVST
	p.WriteCmd(0, ShortRegData2)
	p.WriteCmd(6, lcdRegData7)

	p.WriteCmd(0, ShortRegData29)
	p.WriteCmd(14, lcdRegData8)

	p.WriteCmd(0, ShortRegData30)
	p.WriteCmd(14, lcdRegData9)

	p.WriteCmd(0, ShortRegData31)
	p.WriteCmd(10, lcdRegData10)

	p.WriteCmd(0, ShortRegData32)
	p.WriteCmd(0, ShortRegData46)

	p.WriteCmd(0, ShortRegData2)
	p.WriteCmd(10, lcdRegData11)

	p.WriteCmd(0, ShortRegData33)
	p.WriteCmd(15, lcdRegData12)

	p.WriteCmd(0, ShortRegData29)
	p.WriteCmd(15, lcdRegData13)

	p.WriteCmd(0, ShortRegData30)
	p.WriteCmd(10, lcdRegData14)

	p.WriteCmd(0, ShortRegData31)
	p.WriteCmd(15, lcdRegData15)

	p.WriteCmd(0, ShortRegData32)
	p.WriteCmd(15, lcdRegData16)

	p.WriteCmd(0, ShortRegData34)
	p.WriteCmd(10, lcdRegData17)

	p.WriteCmd(0, ShortRegData35)
	p.WriteCmd(10, lcdRegData18)

	p.WriteCmd(0, ShortRegData2)
	p.WriteCmd(10, lcdRegData19)

	p.WriteCmd(0, ShortRegData33)
	p.WriteCmd(15, lcdRegData20)

	p.WriteCmd(0, ShortRegData29)
	p.WriteCmd(15, lcdRegData21)

	p.WriteCmd(0, ShortRegData30)
	p.WriteCmd(10, lcdRegData22)

	p.WriteCmd(0, ShortRegData31)
	p.WriteCmd(15, lcdRegData23)

	p.WriteCmd(0, ShortRegData32)
	p.WriteCmd(15, lcdRegData24)

	/////////////////////////////////////////////////////////////////////////////
	// PWR_CTRL1 - 0xc580h - 130th parameter - default 0x00
	// Pump 1 min and max DM
	p.WriteCmd(0, ShortRegData13)
	p.WriteCmd(0, ShortRegData47)
	p.WriteCmd(0, ShortRegData48)
	p.WriteCmd(0, ShortRegData49)
	/////////////////////////////////////////////////////////////////////////////

	// CABC LEDPWM frequency adjusted to 19,5kHz
	p.WriteCmd(0, ShortRegData50)
	p.WriteCmd(0, ShortRegData51)

	// Exit CMD2 mode
	p.WriteCmd(0, ShortRegData1)
	p.WriteCmd(3, lcdRegData25)

	///*************************************************************************
	// Standard DCS Initialization TO KEEP CAN BE DONE IN HSDT
	///*************************************************************************

	// NOP - goes back to DCS std command ?
	p.WriteCmd(0, ShortRegData1)

	// Gamma correction 2.2+ table (HSDT possible)
	p.WriteCmd(0, ShortRegData1)
	p.WriteCmd(16, lcdRegData3)

	// Gamma correction 2.2- table (HSDT possible)
	p.WriteCmd(0, ShortRegData1)
	p.WriteCmd(16, lcdRegData4)

	// Send Sleep Out command to display : no parameter
	p.WriteCmd(0, ShortRegData36)

	// Wait for sleep out exit
	delay.Millisec(120)

	switch ColorCoding {
	case OTM8009A_FORMAT_RBG565:
		// Set Pixel color format to RGB565
		p.WriteCmd(0, ShortRegData37)
		break
	case OTM8009A_FORMAT_RGB888:
		// Set Pixel color format to RGB888
		p.WriteCmd(0, ShortRegData38)
		break
	default:
		break
	}

	// Send command to configure display in landscape orientation mode.
	// By default the orientation mode is portrait
	if orientation == OTM8009A_ORIENTATION_LANDSCAPE {
		p.WriteCmd(0, ShortRegData39)
		p.WriteCmd(4, lcdRegData27)
		p.WriteCmd(4, lcdRegData28)
	}

	// CABC : Content Adaptive Backlight Control section start >>
	// Note : defaut is 0 (lowest Brightness), 0xFF is highest Brightness,
	// try 0x7F : intermediate value
	p.WriteCmd(0, ShortRegData40)

	// defaut is 0, try 0x2C - Brightness Control Block, Display Dimming & BackLight on
	p.WriteCmd(0, ShortRegData41)

	// defaut is 0, try 0x02 - image Content based Adaptive Brightness [Still Picture]
	p.WriteCmd(0, ShortRegData42)

	// defaut is 0 (lowest Brightness), 0xFF is highest Brightness
	p.WriteCmd(0, ShortRegData43)
	// CABC : Content Adaptive Backlight Control section end <<

	// Send Command Display On
	p.WriteCmd(0, ShortRegData44)

	// NOP command
	p.WriteCmd(0, ShortRegData1)

	// Send Command GRAM memory write (no parameters) : this initiates frame write via
	// other DSI commands sent by DSI host from LTDC incoming pixels in video mode
	p.WriteCmd(0, ShortRegData45)
}
