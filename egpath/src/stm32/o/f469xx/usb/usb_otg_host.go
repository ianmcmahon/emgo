// Peripheral: USB_OTG_Host_Periph  USB_OTG_Host_Mode_Register_Structures.
// Instances:
// Registers:
//  0x00 32  HCFG     Host Configuration Register          400h.
//  0x04 32  HFIR     Host Frame Interval Register         404h.
//  0x08 32  HFNUM    Host Frame Nbr/Frame Remaining       408h.
//  0x10 32  HPTXSTS  Host Periodic Tx FIFO/ Queue Status  410h.
//  0x14 32  HAINT    Host All Channels Interrupt Register 414h.
//  0x18 32  HAINTMSK Host All Channels Interrupt Mask     418h.
// Import:
//  /Users/ian/code/src/github.com/ianmcmahon/emgo/egpath/src/stm32/o/f469xx/mmap
package usb

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	FSLSPCS HCFG = 0x03 << 0 //+ FS/LS PHY clock select.
	FSLSS   HCFG = 0x01 << 2 //+ FS- and LS-only support.
)

const (
	FSLSPCSn = 0
	FSLSSn   = 2
)

const (
	FRIVL HFIR = 0xFFFF << 0 //+ Frame interval.
)

const (
	FRIVLn = 0
)

const (
	FRNUM HFNUM = 0xFFFF << 0  //+ Frame number.
	FTREM HFNUM = 0xFFFF << 16 //+ Frame time remaining.
)

const (
	FRNUMn = 0
	FTREMn = 16
)

const (
	PTXFSAVL HPTXSTS = 0xFFFF << 0 //+ Periodic transmit data FIFO space available.
	PTXQSAV  HPTXSTS = 0xFF << 16  //+ Periodic transmit request queue space available.
	PTXQTOP  HPTXSTS = 0xFF << 24  //+ Top of the periodic transmit request queue.
)

const (
	PTXFSAVLn = 0
	PTXQSAVn  = 16
	PTXQTOPn  = 24
)

const (
	HAINTM HAINTMSK = 0xFFFF << 0 //+ Channel interrupt mask.
)

const (
	HAINTMn = 0
)
