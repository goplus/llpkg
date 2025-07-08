package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type ApbCtrlDevS struct {
	ClkConf struct {
		Val c.Uint32T
	}
	TickConf struct {
		Val c.Uint32T
	}
	ClkOutEn struct {
		Val c.Uint32T
	}
	WifiBbCfg  c.Uint32T
	WifiBbCfg2 c.Uint32T
	WifiClkEn  c.Uint32T
	WifiRstEn  c.Uint32T
	HostInfSel struct {
		Val c.Uint32T
	}
	ExtMemPmsLock struct {
		Val c.Uint32T
	}
	ExtMemWritebackBypass struct {
		Val c.Uint32T
	}
	FlashAce0Attr struct {
		Val c.Uint32T
	}
	FlashAce1Attr struct {
		Val c.Uint32T
	}
	FlashAce2Attr struct {
		Val c.Uint32T
	}
	FlashAce3Attr struct {
		Val c.Uint32T
	}
	FlashAce0Addr c.Uint32T
	FlashAce1Addr c.Uint32T
	FlashAce2Addr c.Uint32T
	FlashAce3Addr c.Uint32T
	FlashAce0Size struct {
		Val c.Uint32T
	}
	FlashAce1Size struct {
		Val c.Uint32T
	}
	FlashAce2Size struct {
		Val c.Uint32T
	}
	FlashAce3Size struct {
		Val c.Uint32T
	}
	SramAce0Attr struct {
		Val c.Uint32T
	}
	SramAce1Attr struct {
		Val c.Uint32T
	}
	SramAce2Attr struct {
		Val c.Uint32T
	}
	SramAce3Attr struct {
		Val c.Uint32T
	}
	SramAce0Addr c.Uint32T
	SramAce1Addr c.Uint32T
	SramAce2Addr c.Uint32T
	SramAce3Addr c.Uint32T
	SramAce0Size struct {
		Val c.Uint32T
	}
	SramAce1Size struct {
		Val c.Uint32T
	}
	SramAce2Size struct {
		Val c.Uint32T
	}
	SramAce3Size struct {
		Val c.Uint32T
	}
	SpiMemPmsCtrl struct {
		Val c.Uint32T
	}
	SpiMemRejectAddr c.Uint32T
	SdioCtrl         struct {
		Val c.Uint32T
	}
	RedcySig0 struct {
		Val c.Uint32T
	}
	RedcySig1 struct {
		Val c.Uint32T
	}
	FrontEndMemPd struct {
		Val c.Uint32T
	}
	SpiMemEccCtrl struct {
		Val c.Uint32T
	}
	ReservedA4     c.Uint32T
	ClkgateForceOn struct {
		Val c.Uint32T
	}
	MemPowerDown struct {
		Val c.Uint32T
	}
	MemPowerUp struct {
		Val c.Uint32T
	}
	RetentionCtrl struct {
		Val c.Uint32T
	}
	RetentionCtrl1 struct {
		Val c.Uint32T
	}
	RetentionCtrl2 struct {
		Val c.Uint32T
	}
	RetentionCtrl3 struct {
		Val c.Uint32T
	}
	RetentionCtrl4 c.Uint32T
	RetentionCtrl5 struct {
		Val c.Uint32T
	}
	ReservedCc  c.Uint32T
	ReservedD0  c.Uint32T
	ReservedD4  c.Uint32T
	ReservedD8  c.Uint32T
	ReservedDc  c.Uint32T
	ReservedE0  c.Uint32T
	ReservedE4  c.Uint32T
	ReservedE8  c.Uint32T
	ReservedEc  c.Uint32T
	ReservedF0  c.Uint32T
	ReservedF4  c.Uint32T
	ReservedF8  c.Uint32T
	ReservedFc  c.Uint32T
	Reserved100 c.Uint32T
	Reserved104 c.Uint32T
	Reserved108 c.Uint32T
	Reserved10c c.Uint32T
	Reserved110 c.Uint32T
	Reserved114 c.Uint32T
	Reserved118 c.Uint32T
	Reserved11c c.Uint32T
	Reserved120 c.Uint32T
	Reserved124 c.Uint32T
	Reserved128 c.Uint32T
	Reserved12c c.Uint32T
	Reserved130 c.Uint32T
	Reserved134 c.Uint32T
	Reserved138 c.Uint32T
	Reserved13c c.Uint32T
	Reserved140 c.Uint32T
	Reserved144 c.Uint32T
	Reserved148 c.Uint32T
	Reserved14c c.Uint32T
	Reserved150 c.Uint32T
	Reserved154 c.Uint32T
	Reserved158 c.Uint32T
	Reserved15c c.Uint32T
	Reserved160 c.Uint32T
	Reserved164 c.Uint32T
	Reserved168 c.Uint32T
	Reserved16c c.Uint32T
	Reserved170 c.Uint32T
	Reserved174 c.Uint32T
	Reserved178 c.Uint32T
	Reserved17c c.Uint32T
	Reserved180 c.Uint32T
	Reserved184 c.Uint32T
	Reserved188 c.Uint32T
	Reserved18c c.Uint32T
	Reserved190 c.Uint32T
	Reserved194 c.Uint32T
	Reserved198 c.Uint32T
	Reserved19c c.Uint32T
	Reserved1a0 c.Uint32T
	Reserved1a4 c.Uint32T
	Reserved1a8 c.Uint32T
	Reserved1ac c.Uint32T
	Reserved1b0 c.Uint32T
	Reserved1b4 c.Uint32T
	Reserved1b8 c.Uint32T
	Reserved1bc c.Uint32T
	Reserved1c0 c.Uint32T
	Reserved1c4 c.Uint32T
	Reserved1c8 c.Uint32T
	Reserved1cc c.Uint32T
	Reserved1d0 c.Uint32T
	Reserved1d4 c.Uint32T
	Reserved1d8 c.Uint32T
	Reserved1dc c.Uint32T
	Reserved1e0 c.Uint32T
	Reserved1e4 c.Uint32T
	Reserved1e8 c.Uint32T
	Reserved1ec c.Uint32T
	Reserved1f0 c.Uint32T
	Reserved1f4 c.Uint32T
	Reserved1f8 c.Uint32T
	Reserved1fc c.Uint32T
	Reserved200 c.Uint32T
	Reserved204 c.Uint32T
	Reserved208 c.Uint32T
	Reserved20c c.Uint32T
	Reserved210 c.Uint32T
	Reserved214 c.Uint32T
	Reserved218 c.Uint32T
	Reserved21c c.Uint32T
	Reserved220 c.Uint32T
	Reserved224 c.Uint32T
	Reserved228 c.Uint32T
	Reserved22c c.Uint32T
	Reserved230 c.Uint32T
	Reserved234 c.Uint32T
	Reserved238 c.Uint32T
	Reserved23c c.Uint32T
	Reserved240 c.Uint32T
	Reserved244 c.Uint32T
	Reserved248 c.Uint32T
	Reserved24c c.Uint32T
	Reserved250 c.Uint32T
	Reserved254 c.Uint32T
	Reserved258 c.Uint32T
	Reserved25c c.Uint32T
	Reserved260 c.Uint32T
	Reserved264 c.Uint32T
	Reserved268 c.Uint32T
	Reserved26c c.Uint32T
	Reserved270 c.Uint32T
	Reserved274 c.Uint32T
	Reserved278 c.Uint32T
	Reserved27c c.Uint32T
	Reserved280 c.Uint32T
	Reserved284 c.Uint32T
	Reserved288 c.Uint32T
	Reserved28c c.Uint32T
	Reserved290 c.Uint32T
	Reserved294 c.Uint32T
	Reserved298 c.Uint32T
	Reserved29c c.Uint32T
	Reserved2a0 c.Uint32T
	Reserved2a4 c.Uint32T
	Reserved2a8 c.Uint32T
	Reserved2ac c.Uint32T
	Reserved2b0 c.Uint32T
	Reserved2b4 c.Uint32T
	Reserved2b8 c.Uint32T
	Reserved2bc c.Uint32T
	Reserved2c0 c.Uint32T
	Reserved2c4 c.Uint32T
	Reserved2c8 c.Uint32T
	Reserved2cc c.Uint32T
	Reserved2d0 c.Uint32T
	Reserved2d4 c.Uint32T
	Reserved2d8 c.Uint32T
	Reserved2dc c.Uint32T
	Reserved2e0 c.Uint32T
	Reserved2e4 c.Uint32T
	Reserved2e8 c.Uint32T
	Reserved2ec c.Uint32T
	Reserved2f0 c.Uint32T
	Reserved2f4 c.Uint32T
	Reserved2f8 c.Uint32T
	Reserved2fc c.Uint32T
	Reserved300 c.Uint32T
	Reserved304 c.Uint32T
	Reserved308 c.Uint32T
	Reserved30c c.Uint32T
	Reserved310 c.Uint32T
	Reserved314 c.Uint32T
	Reserved318 c.Uint32T
	Reserved31c c.Uint32T
	Reserved320 c.Uint32T
	Reserved324 c.Uint32T
	Reserved328 c.Uint32T
	Reserved32c c.Uint32T
	Reserved330 c.Uint32T
	Reserved334 c.Uint32T
	Reserved338 c.Uint32T
	Reserved33c c.Uint32T
	Reserved340 c.Uint32T
	Reserved344 c.Uint32T
	Reserved348 c.Uint32T
	Reserved34c c.Uint32T
	Reserved350 c.Uint32T
	Reserved354 c.Uint32T
	Reserved358 c.Uint32T
	Reserved35c c.Uint32T
	Reserved360 c.Uint32T
	Reserved364 c.Uint32T
	Reserved368 c.Uint32T
	Reserved36c c.Uint32T
	Reserved370 c.Uint32T
	Reserved374 c.Uint32T
	Reserved378 c.Uint32T
	Reserved37c c.Uint32T
	Reserved380 c.Uint32T
	Reserved384 c.Uint32T
	Reserved388 c.Uint32T
	Reserved38c c.Uint32T
	Reserved390 c.Uint32T
	Reserved394 c.Uint32T
	Reserved398 c.Uint32T
	Reserved39c c.Uint32T
	Reserved3a0 c.Uint32T
	Reserved3a4 c.Uint32T
	Reserved3a8 c.Uint32T
	Reserved3ac c.Uint32T
	Reserved3b0 c.Uint32T
	Reserved3b4 c.Uint32T
	Reserved3b8 c.Uint32T
	Reserved3bc c.Uint32T
	Reserved3c0 c.Uint32T
	Reserved3c4 c.Uint32T
	Reserved3c8 c.Uint32T
	Reserved3cc c.Uint32T
	Reserved3d0 c.Uint32T
	Reserved3d4 c.Uint32T
	Reserved3d8 c.Uint32T
	Reserved3dc c.Uint32T
	Reserved3e0 c.Uint32T
	Reserved3e4 c.Uint32T
	Reserved3e8 c.Uint32T
	Reserved3ec c.Uint32T
	Reserved3f0 c.Uint32T
	Reserved3f4 c.Uint32T
	Reserved3f8 c.Uint32T
	Date        c.Uint32T
}
type ApbCtrlDevT ApbCtrlDevS
