// +build f746xx

// Peripheral: USB_OTG_OUTEndpoint_Periph  USB_OTG_OUT_Endpoint-Specific_Registers.
// Instances:
// Registers:
//  0x00 32  DOEPCTL  dev OUT Endpoint Control Reg           B00h + (ep_num * 20h) + 00h.
//  0x08 32  DOEPINT  dev OUT Endpoint Itr Reg               B00h + (ep_num * 20h) + 08h.
//  0x10 32  DOEPTSIZ dev OUT Endpoint Txfer Size            B00h + (ep_num * 20h) + 10h.
//  0x14 32  DOEPDMA  dev OUT Endpoint DMA Address           B00h + (ep_num * 20h) + 14h.
// Import:
//  stm32/o/f746xx/mmap
package usb

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	MPSIZ          DOEPCTL = 0x7FF << 0 //+ Maximum packet size */          /*!<Bit 1.
	USBAEP         DOEPCTL = 0x01 << 15 //+ USB active endpoint.
	NAKSTS         DOEPCTL = 0x01 << 17 //+ NAK status.
	SD0PID_SEVNFRM DOEPCTL = 0x01 << 28 //+ Set DATA0 PID.
	SODDFRM        DOEPCTL = 0x01 << 29 //+ Set odd frame.
	EPTYP          DOEPCTL = 0x03 << 18 //+ Endpoint type.
	SNPM           DOEPCTL = 0x01 << 20 //+ Snoop mode.
	STALL          DOEPCTL = 0x01 << 21 //+ STALL handshake.
	CNAK           DOEPCTL = 0x01 << 26 //+ Clear NAK.
	SNAK           DOEPCTL = 0x01 << 27 //+ Set NAK.
	EPDIS          DOEPCTL = 0x01 << 30 //+ Endpoint disable.
	EPENA          DOEPCTL = 0x01 << 31 //+ Endpoint enable.
)

const (
	MPSIZn          = 0
	USBAEPn         = 15
	NAKSTSn         = 17
	SD0PID_SEVNFRMn = 28
	SODDFRMn        = 29
	EPTYPn          = 18
	SNPMn           = 20
	STALLn          = 21
	CNAKn           = 26
	SNAKn           = 27
	EPDISn          = 30
	EPENAn          = 31
)

const (
	XFRC    DOEPINT = 0x01 << 0  //+ Transfer completed interrupt.
	EPDISD  DOEPINT = 0x01 << 1  //+ Endpoint disabled interrupt.
	STUP    DOEPINT = 0x01 << 3  //+ SETUP phase done.
	OTEPDIS DOEPINT = 0x01 << 4  //+ OUT token received when endpoint disabled.
	OTEPSPR DOEPINT = 0x01 << 5  //+ Status Phase Received For Control Write.
	B2BSTUP DOEPINT = 0x01 << 6  //+ Back-to-back SETUP packets received.
	NYET    DOEPINT = 0x01 << 14 //+ NYET interrupt.
)

const (
	XFRCn    = 0
	EPDISDn  = 1
	STUPn    = 3
	OTEPDISn = 4
	OTEPSPRn = 5
	B2BSTUPn = 6
	NYETn    = 14
)

const (
	XFRSIZ  DOEPTSIZ = 0x7FFFF << 0 //+ Transfer size.
	PKTCNT  DOEPTSIZ = 0x3FF << 19  //+ Packet count.
	STUPCNT DOEPTSIZ = 0x03 << 29   //+ SETUP packet count.
)

const (
	XFRSIZn  = 0
	PKTCNTn  = 19
	STUPCNTn = 29
)
