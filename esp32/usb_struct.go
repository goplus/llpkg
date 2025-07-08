package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/* USB IN EP Register block type */

type UsbInEpReg struct {
	Diepctl   c.Uint32T
	Reserved  c.Uint32T
	Diepint   c.Uint32T
	Reserved1 c.Uint32T
	Dieptsiz  c.Uint32T
	Diepdma   c.Uint32T
	Dtxfsts   c.Uint32T
	Reserved2 c.Uint32T
}
type UsbInEndpointT UsbInEpReg

/* USB OUT EP Register block type */

type UsbOutEpReg struct {
	Doepctl   c.Uint32T
	Reserved  c.Uint32T
	Doepint   c.Uint32T
	Reserved1 c.Uint32T
	Doeptsiz  c.Uint32T
	Doepdma   c.Uint32T
	Reserved2 c.Uint32T
	Reserved3 c.Uint32T
}
type UsbOutEndpointT UsbOutEpReg

type UsbReg struct {
	Gotgctl              c.Uint32T
	Gotgint              c.Uint32T
	Gahbcfg              c.Uint32T
	Gusbcfg              c.Uint32T
	Grstctl              c.Uint32T
	Gintsts              c.Uint32T
	Gintmsk              c.Uint32T
	Grxstsr              c.Uint32T
	Grxstsp              c.Uint32T
	Grxfsiz              c.Uint32T
	Gnptxfsiz            c.Uint32T
	Gnptxsts             c.Uint32T
	Reserved0x00300x0040 [4]c.Uint32T
	Gsnpsid              c.Uint32T
	Ghwcfg1              c.Uint32T
	Ghwcfg2              c.Uint32T
	Ghwcfg3              c.Uint32T
	Ghwcfg4              c.Uint32T
	Reserved0x00540x005c [2]c.Uint32T
	Gdfifocfg            c.Uint32T
	Reserved0x00600x0100 [40]c.Uint32T
	Hptxfsiz             c.Uint32T
	Dieptxf              [4]c.Uint32T
	Reserved0x01140x0140 [11]c.Uint32T
	Reserved0x01400x0400 [176]c.Uint32T
	Hcfg                 c.Uint32T
	Hfir                 c.Uint32T
	Hfnum                c.Uint32T
	Reserved0x40C        c.Uint32T
	Hptxsts              c.Uint32T
	Haint                c.Uint32T
	Haintmsk             c.Uint32T
	Hflbaddr             c.Uint32T
	Reserved0x04200x0440 [8]c.Uint32T
	Hprt                 c.Uint32T
	Reserved0x04440x0500 [47]c.Uint32T
	HostChanRegs         [128]c.Uint32T
	Reserved0x07000x0800 [64]c.Uint32T
	Dcfg                 c.Uint32T
	Dctl                 c.Uint32T
	Dsts                 c.Uint32T
	Reserved0x80c        c.Uint32T
	Diepmsk              c.Uint32T
	Doepmsk              c.Uint32T
	Daint                c.Uint32T
	Daintmsk             c.Uint32T
	Reserved0x08200x0828 [2]c.Uint32T
	Dvbusdis             c.Uint32T
	Dvbuspulse           c.Uint32T
	Dthrctl              c.Uint32T
	Dtknqr4Fifoemptymsk  c.Uint32T
	Reserved0x08380x0900 [50]c.Uint32T
	InEpReg              [7]UsbInEndpointT
	Reserved0x09e00x0b00 [72]c.Uint32T
	OutEpReg             [7]UsbOutEndpointT
	Reserved0x0be00x0d00 [72]c.Uint32T
	Reserved0x0d000x0e00 [64]c.Uint32T
	Pcgctrl              c.Uint32T
	Reserved0x0e04       c.Uint32T
	Reserved8            [504]c.Uint8T
	Fifo                 [16][1024]c.Uint32T
	Reserved0x11000      [61440]c.Uint8T
	DbgFifo              [131072]c.Uint32T
}
type UsbDevT UsbReg
