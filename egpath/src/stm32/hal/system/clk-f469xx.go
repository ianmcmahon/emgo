// +build f469xx

package system

const HSIClk = 16e6 // Hz

const (
	maxAPB1Clk = 48e6 // Hz (slightly overclocked from 45 MHz)
	maxAPB2Clk = 96e6 // Hz (slightly overclocked from 90 MHz)
)

const (
	maxAPB1ClkOverDrive = 54e6  // Hz
	maxAPB2ClkOverDrive = 108e6 // Hz
)
