// +build f46xxx
package otm8009a

import (
	"delay"
	"stm32/hal/dsi"
)

type LCD struct {
}

func WriteCmd(cmd []byte) {
	if len(cmd) > 2 {
		dsi.DSIPeriph.LongWrite(LCD_OTM8009A_ID, dsi.DSI_DCS_LONG_PKT_WRITE, cmd)
	} else {
		dsi.DSIPeriph.ShortWrite(LCD_OTM8009A_ID, dsi.DSI_DCS_SHORT_PKT_WRITE_P1, cmd[0], cmd[1])
	}
}

func Init(ColorCoding uint32, orientation uint32) {
	// Enable CMD2 to access vendor specific commands
	// Enter in command 2 mode and set EXTC to enable address shift function (0x00)
	WriteCmd(ShortRegData1)
	WriteCmd(lcdRegData1)

	// Enter ORISE Command 2
	WriteCmd(ShortRegData2) // Shift address to 0x80
	WriteCmd(lcdRegData2)

	/////////////////////////////////////////////////////////////////////
	// SD_PCH_CTRL - 0xC480h - 129th parameter - Default 0x00
	// Set SD_PT
	// -> Source output level during porch and non-display area to GND
	WriteCmd(ShortRegData2)
	WriteCmd(ShortRegData3)
	delay.Millisec(10)
	// Not documented
	WriteCmd(ShortRegData4)
	WriteCmd(ShortRegData5)
	delay.Millisec(10)
	/////////////////////////////////////////////////////////////////////

	// PWR_CTRL4 - 0xC4B0h - 178th parameter - Default 0xA8
	// Set gvdd_en_test
	// -> enable GVDD test mode !!!
	WriteCmd(ShortRegData6)
	WriteCmd(ShortRegData7)

	// PWR_CTRL2 - 0xC590h - 146th parameter - Default 0x79
	// Set pump 4 vgh voltage
	// -> from 15.0v down to 13.0v
	// Set pump 5 vgh voltage
	// -> from -12.0v downto -9.0v
	WriteCmd(ShortRegData8)
	WriteCmd(ShortRegData9)

	// P_DRV_M - 0xC0B4h - 181th parameter - Default 0x00
	// -> Column inversion
	WriteCmd(ShortRegData10)
	WriteCmd(ShortRegData11)

	// VCOMDC - 0xD900h - 1st parameter - Default 0x39h
	// VCOM Voltage settings
	// -> from -1.0000v downto -1.2625v
	WriteCmd(ShortRegData1)
	WriteCmd(ShortRegData12)

	// Oscillator adjustment for Idle/Normal mode (LPDT only) set to 65Hz (default is 60Hz)
	WriteCmd(ShortRegData13)
	WriteCmd(ShortRegData14)

	// Video mode internal
	WriteCmd(ShortRegData15)
	WriteCmd(ShortRegData16)

	// PWR_CTRL2 - 0xC590h - 147h parameter - Default 0x00
	// Set pump 4&5 x6
	// -> ONLY VALID when PUMP4_EN_ASDM_HV = "0"
	WriteCmd(ShortRegData17)
	WriteCmd(ShortRegData18)

	// PWR_CTRL2 - 0xC590h - 150th parameter - Default 0x33h
	// Change pump4 clock ratio
	// -> from 1 line to 1/2 line
	WriteCmd(ShortRegData19)
	WriteCmd(ShortRegData9)

	// GVDD/NGVDD settings
	WriteCmd(ShortRegData1)
	WriteCmd(lcdRegData5)

	// PWR_CTRL2 - 0xC590h - 149th parameter - Default 0x33h
	// Rewrite the default value !
	WriteCmd(ShortRegData20)
	WriteCmd(ShortRegData21)

	// Panel display timing Setting 3
	WriteCmd(ShortRegData22)
	WriteCmd(ShortRegData23)

	// Power control 1
	WriteCmd(ShortRegData24)
	WriteCmd(ShortRegData25)

	// Source driver precharge
	WriteCmd(ShortRegData13)
	WriteCmd(ShortRegData26)

	WriteCmd(ShortRegData15)
	WriteCmd(ShortRegData27)

	WriteCmd(ShortRegData28)
	WriteCmd(lcdRegData6)

	// GOAVST
	WriteCmd(ShortRegData2)
	WriteCmd(lcdRegData7)

	WriteCmd(ShortRegData29)
	WriteCmd(lcdRegData8)

	WriteCmd(ShortRegData30)
	WriteCmd(lcdRegData9)

	WriteCmd(ShortRegData31)
	WriteCmd(lcdRegData10)

	WriteCmd(ShortRegData32)
	WriteCmd(ShortRegData46)

	WriteCmd(ShortRegData2)
	WriteCmd(lcdRegData11)

	WriteCmd(ShortRegData33)
	WriteCmd(lcdRegData12)

	WriteCmd(ShortRegData29)
	WriteCmd(lcdRegData13)

	WriteCmd(ShortRegData30)
	WriteCmd(lcdRegData14)

	WriteCmd(ShortRegData31)
	WriteCmd(lcdRegData15)

	WriteCmd(ShortRegData32)
	WriteCmd(lcdRegData16)

	WriteCmd(ShortRegData34)
	WriteCmd(lcdRegData17)

	WriteCmd(ShortRegData35)
	WriteCmd(lcdRegData18)

	WriteCmd(ShortRegData2)
	WriteCmd(lcdRegData19)

	WriteCmd(ShortRegData33)
	WriteCmd(lcdRegData20)

	WriteCmd(ShortRegData29)
	WriteCmd(lcdRegData21)

	WriteCmd(ShortRegData30)
	WriteCmd(lcdRegData22)

	WriteCmd(ShortRegData31)
	WriteCmd(lcdRegData23)

	WriteCmd(ShortRegData32)
	WriteCmd(lcdRegData24)

	/////////////////////////////////////////////////////////////////////////////
	// PWR_CTRL1 - 0xc580h - 130th parameter - default 0x00
	// Pump 1 min and max DM
	WriteCmd(ShortRegData13)
	WriteCmd(ShortRegData47)
	WriteCmd(ShortRegData48)
	WriteCmd(ShortRegData49)
	/////////////////////////////////////////////////////////////////////////////

	// CABC LEDPWM frequency adjusted to 19,5kHz
	WriteCmd(ShortRegData50)
	WriteCmd(ShortRegData51)

	// Exit CMD2 mode
	WriteCmd(ShortRegData1)
	WriteCmd(lcdRegData25)

	///*************************************************************************
	// Standard DCS Initialization TO KEEP CAN BE DONE IN HSDT
	///*************************************************************************

	// NOP - goes back to DCS std command ?
	WriteCmd(ShortRegData1)

	// Gamma correction 2.2+ table (HSDT possible)
	WriteCmd(ShortRegData1)
	WriteCmd(lcdRegData3)

	// Gamma correction 2.2- table (HSDT possible)
	WriteCmd(ShortRegData1)
	WriteCmd(lcdRegData4)

	// Send Sleep Out command to display : no parameter
	WriteCmd(ShortRegData36)

	// Wait for sleep out exit
	delay.Millisec(120)

	switch ColorCoding {
	case OTM8009A_FORMAT_RBG565:
		// Set Pixel color format to RGB565
		WriteCmd(ShortRegData37)
		break
	case OTM8009A_FORMAT_RGB888:
		// Set Pixel color format to RGB888
		WriteCmd(ShortRegData38)
		break
	default:
		break
	}

	// Send command to configure display in landscape orientation mode.
	// By default the orientation mode is portrait
	if orientation == OTM8009A_ORIENTATION_LANDSCAPE {
		WriteCmd(ShortRegData39)
		WriteCmd(lcdRegData27)
		WriteCmd(lcdRegData28)
	}

	// CABC : Content Adaptive Backlight Control section start >>
	// Note : defaut is 0 (lowest Brightness), 0xFF is highest Brightness,
	// try 0x7F : intermediate value
	WriteCmd(ShortRegData40)

	// defaut is 0, try 0x2C - Brightness Control Block, Display Dimming & BackLight on
	WriteCmd(ShortRegData41)

	// defaut is 0, try 0x02 - image Content based Adaptive Brightness [Still Picture]
	WriteCmd(ShortRegData42)

	// defaut is 0 (lowest Brightness), 0xFF is highest Brightness
	WriteCmd(ShortRegData43)
	// CABC : Content Adaptive Backlight Control section end <<

	// Send Command Display On
	WriteCmd(ShortRegData44)

	// NOP command
	WriteCmd(ShortRegData1)

	// Send Command GRAM memory write (no parameters) : this initiates frame write via
	// other DSI commands sent by DSI host from LTDC incoming pixels in video mode
	WriteCmd(ShortRegData45)
}
