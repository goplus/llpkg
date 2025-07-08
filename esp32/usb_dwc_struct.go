package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/* ---------------------------- Register Types ------------------------------ */

type UsbDwcGotgctlRegT struct {
	Val c.Uint32T
}

type UsbDwcGotgintRegT struct {
	Val c.Uint32T
}

type UsbDwcGahbcfgRegT struct {
	Val c.Uint32T
}

type UsbDwcGusbcfgRegT struct {
	Val c.Uint32T
}

type UsbDwcGrstctlRegT struct {
	Val c.Uint32T
}

type UsbDwcGintstsRegT struct {
	Val c.Uint32T
}

type UsbDwcGintmskRegT struct {
	Val c.Uint32T
}

type UsbDwcGrxstsrRegT struct {
	Val c.Uint32T
}

type UsbDwcGrxstspRegT struct {
	Val c.Uint32T
}

type UsbDwcGrxfsizRegT struct {
	Val c.Uint32T
}

type UsbDwcGnptxfsizRegT struct {
	Val c.Uint32T
}

type UsbDwcGnptxstsRegT struct {
	Val c.Uint32T
}

type UsbDwcGsnpsidRegT struct {
	Val c.Uint32T
}

type UsbDwcGhwcfg1RegT struct {
	Val c.Uint32T
}

type UsbDwcGhwcfg2RegT struct {
	Val c.Uint32T
}

type UsbDwcGhwcfg3RegT struct {
	Val c.Uint32T
}

type UsbDwcGhwcfg4RegT struct {
	Val c.Uint32T
}

type UsbDwcGdfifocfgRegT struct {
	Val c.Uint32T
}

type UsbDwcHptxfsizRegT struct {
	Val c.Uint32T
}

type UsbDwcDieptxfiRegT struct {
	Val c.Uint32T
}

type UsbDwcHcfgRegT struct {
	Val c.Uint32T
}

type UsbDwcHfirRegT struct {
	Val c.Uint32T
}

type UsbDwcHfnumRegT struct {
	Val c.Uint32T
}

type UsbDwcHptxstsRegT struct {
	Val c.Uint32T
}

type UsbDwcHaintRegT struct {
	Val c.Uint32T
}

type UsbDwcHaintmskRegT struct {
	Val c.Uint32T
}

type UsbDwcHflbaddrRegT struct {
	Val c.Uint32T
}

type UsbDwcHprtRegT struct {
	Val c.Uint32T
}

type UsbDwcHccharRegT struct {
	Val c.Uint32T
}

type UsbDwcHcintRegT struct {
	Val c.Uint32T
}

type UsbDwcHcintmskRegT struct {
	Val c.Uint32T
}

type UsbDwcHctsizRegT struct {
	Val c.Uint32T
}

type UsbDwcHcdmaRegT struct {
	Val c.Uint32T
}

type UsbDwcHcdmabRegT struct {
	Val c.Uint32T
}

type UsbDwcDcfgRegT struct {
	Val c.Uint32T
}

type UsbDwcDctlRegT struct {
	Val c.Uint32T
}

type UsbDwcDstsRegT struct {
	Val c.Uint32T
}

type UsbDwcDiepmskRegT struct {
	Val c.Uint32T
}

type UsbDwcDoepmskRegT struct {
	Val c.Uint32T
}

type UsbDwcDaintRegT struct {
	Val c.Uint32T
}

type UsbDwcDaintmskRegT struct {
	Val c.Uint32T
}

type UsbDwcDvbusdisRegT struct {
	Val c.Uint32T
}

type UsbDwcDvbuspulseRegT struct {
	Val c.Uint32T
}

type UsbDwcDthrctlRegT struct {
	Val c.Uint32T
}

type UsbDwcDiepempmskRegT struct {
	Val c.Uint32T
}

type UsbDwcDiepctl0RegT struct {
	Val c.Uint32T
}

type UsbDwcDiepint0RegT struct {
	Val c.Uint32T
}

type UsbDwcDieptsiz0RegT struct {
	Val c.Uint32T
}

type UsbDwcDiepdma0RegT struct {
	Val c.Uint32T
}

type UsbDwcDtxfsts0RegT struct {
	Val c.Uint32T
}

type UsbDwcDiepdmab0RegT struct {
	Val c.Uint32T
}

type UsbDwcDiepctlRegT struct {
	Val c.Uint32T
}

type UsbDwcDiepintRegT struct {
	Val c.Uint32T
}

type UsbDwcDieptsizRegT struct {
	Val c.Uint32T
}

type UsbDwcDiepdmaRegT struct {
	Val c.Uint32T
}

type UsbDwcDtxfstsRegT struct {
	Val c.Uint32T
}

type UsbDwcDiepdmabRegT struct {
	Val c.Uint32T
}

type UsbDwcDoepctl0RegT struct {
	Val c.Uint32T
}

type UsbDwcDoepint0RegT struct {
	Val c.Uint32T
}

type UsbDwcDoeptsiz0RegT struct {
	Val c.Uint32T
}

type UsbDwcDoepdma0RegT struct {
	Val c.Uint32T
}

type UsbDwcDoepdmab0RegT struct {
	Val c.Uint32T
}

type UsbDwcDoepctlRegT struct {
	Val c.Uint32T
}

type UsbDwcDoepintRegT struct {
	Val c.Uint32T
}

type UsbDwcDoeptsizRegT struct {
	Val c.Uint32T
}

type UsbDwcDoepdmaRegT struct {
	Val c.Uint32T
}

type UsbDwcDoepdmabRegT struct {
	Val c.Uint32T
}

type UsbDwcPcgcctlRegT struct {
	Val c.Uint32T
}

/* --------------------------- Register Groups ------------------------------ */

type UsbDwcHostChanRegsT struct {
	HccharReg    UsbDwcHccharRegT
	Reserved0x04 [1]c.Uint32T
	HcintReg     UsbDwcHcintRegT
	HcintmskReg  UsbDwcHcintmskRegT
	HctsizReg    UsbDwcHctsizRegT
	HcdmaReg     UsbDwcHcdmaRegT
	Reserved0x18 [1]c.Uint32T
	HcdmabReg    UsbDwcHcdmabRegT
}

type UsbDwcInEpRegsT struct {
	DiepctlReg   UsbDwcDiepctlRegT
	Reserved0x04 [1]c.Uint32T
	DiepintReg   UsbDwcDiepintRegT
	Reserved0x0c [1]c.Uint32T
	DieptsizReg  UsbDwcDieptsizRegT
	DiepdmaReg   UsbDwcDiepdmaRegT
	DtxfstsReg   UsbDwcDtxfstsRegT
	DiepdmabReg  UsbDwcDiepdmabRegT
}

type UsbDwcOutEpRegsT struct {
	DoepctlReg   UsbDwcDoepctlRegT
	Reserved0x04 [1]c.Uint32T
	DoepintReg   UsbDwcDoepintRegT
	Reserved0x0c [1]c.Uint32T
	DoeptsizReg  UsbDwcDoeptsizRegT
	DoepdmaReg   UsbDwcDoepdmaRegT
	Reserved0x18 [1]c.Uint32T
	DoepdmabReg  UsbDwcDoepdmabRegT
}

/* --------------------------- Register Layout ------------------------------ */

type UsbDwcDevT struct {
	GotgctlReg           UsbDwcGotgctlRegT
	GotgintReg           UsbDwcGotgintRegT
	GahbcfgReg           UsbDwcGahbcfgRegT
	GusbcfgReg           UsbDwcGusbcfgRegT
	GrstctlReg           UsbDwcGrstctlRegT
	GintstsReg           UsbDwcGintstsRegT
	GintmskReg           UsbDwcGintmskRegT
	GrxstsrReg           UsbDwcGrxstsrRegT
	GrxstspReg           UsbDwcGrxstspRegT
	GrxfsizReg           UsbDwcGrxfsizRegT
	GnptxfsizReg         UsbDwcGnptxfsizRegT
	GnptxstsReg          UsbDwcGnptxstsRegT
	Reserved0x0030       c.Uint32T
	Reserved0x0034       c.Uint32T
	Reserved0x0038       c.Uint32T
	Reserved0x003c       c.Uint32T
	GsnpsidReg           UsbDwcGsnpsidRegT
	Ghwcfg1Reg           UsbDwcGhwcfg1RegT
	Ghwcfg2Reg           UsbDwcGhwcfg2RegT
	Ghwcfg3Reg           UsbDwcGhwcfg3RegT
	Ghwcfg4Reg           UsbDwcGhwcfg4RegT
	Reserved0x0054       c.Uint32T
	Reserved0x0058       c.Uint32T
	GdfifocfgReg         UsbDwcGdfifocfgRegT
	Reserved0x0060       c.Uint32T
	Reserved0x00640x0100 [39]c.Uint32T
	HptxfsizReg          UsbDwcHptxfsizRegT
	DieptxfRegs          [4]UsbDwcDieptxfiRegT
	Reserved0x01140x013c [11]UsbDwcDieptxfiRegT
	Reserved0x1400x3fc   [176]c.Uint32T
	HcfgReg              UsbDwcHcfgRegT
	HfirReg              UsbDwcHfirRegT
	HfnumReg             UsbDwcHfnumRegT
	Reserved0x40c        [1]c.Uint32T
	HptxstsReg           UsbDwcHptxstsRegT
	HaintReg             UsbDwcHaintRegT
	HaintmskReg          UsbDwcHaintmskRegT
	HflbaddrReg          UsbDwcHflbaddrRegT
	Reserved0x4200x43c   [8]c.Uint32T
	HprtReg              UsbDwcHprtRegT
	Reserved0x04440x04fc [47]c.Uint32T
	HostChans            [8]UsbDwcHostChanRegsT
	Reserved0x06000x06fc [8]UsbDwcHostChanRegsT
	Reserved0x07000x07fc [64]c.Uint32T
	DcfgReg              UsbDwcDcfgRegT
	DctlReg              UsbDwcDctlRegT
	DstsReg              UsbDwcDstsRegT
	Reserved0x080c       [1]c.Uint32T
	DiepmskReg           UsbDwcDiepmskRegT
	DoepmskReg           UsbDwcDoepmskRegT
	DaintReg             UsbDwcDaintRegT
	DaintmskReg          UsbDwcDaintmskRegT
	Reserved0x0820       c.Uint32T
	Reserved0x0824       c.Uint32T
	DvbusdisReg          UsbDwcDvbusdisRegT
	DvbuspulseReg        UsbDwcDvbuspulseRegT
	DthrctlReg           UsbDwcDthrctlRegT
	DiepempmskReg        UsbDwcDiepempmskRegT
	Reserved0x0838       c.Uint32T
	Reserved0x083c       c.Uint32T
	Reserved0x0840       c.Uint32T
	Reserved0x08440x087c [15]c.Uint32T
	Reserved0x0880       c.Uint32T
	Reserved0x08840x08c0 [15]c.Uint32T
	Reserved0x08c40x08fc [16]c.Uint32T
	Diepctl0Reg          UsbDwcDiepctl0RegT
	Reserved0x0904       [1]c.Uint32T
	Diepint0Reg          UsbDwcDiepint0RegT
	Reserved0x090c       [1]c.Uint32T
	Dieptsiz0Reg         UsbDwcDieptsiz0RegT
	Diepdma0Reg          UsbDwcDiepdma0RegT
	Dtxfsts0Reg          UsbDwcDtxfsts0RegT
	Diepdmab0Reg         UsbDwcDiepdmab0RegT
	InEps                [6]UsbDwcInEpRegsT
	Reserved0x09e00x0afc [9]UsbDwcInEpRegsT
	Doepctl0Reg          UsbDwcDoepctl0RegT
	Reserved0x0b04       [1]c.Uint32T
	Doepint0Reg          UsbDwcDoepint0RegT
	Reserved0x0b0c       [1]c.Uint32T
	Doeptsiz0Reg         UsbDwcDoeptsiz0RegT
	Doepdma0Reg          UsbDwcDoepdma0RegT
	Reserved0x0b18       [1]c.Uint32T
	Doepdmab0Reg         UsbDwcDoepdmab0RegT
	OutEps               [6]UsbDwcOutEpRegsT
	Reserved0x0be00x0d00 [9]UsbDwcOutEpRegsT
	Reserved0x0d000x0dfc [64]c.Uint32T
	PcgcctlReg           UsbDwcPcgcctlRegT
	Reserved0x0e04       [1]c.Uint32T
}
