package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type SensitiveDevS struct {
	CacheDataarrayConnect0 struct {
		Val c.Uint32T
	}
	CacheDataarrayConnect1 struct {
		Val c.Uint32T
	}
	ApbPeripheralAccess0 struct {
		Val c.Uint32T
	}
	ApbPeripheralAccess1 struct {
		Val c.Uint32T
	}
	InternalSramUsage0 struct {
		Val c.Uint32T
	}
	InternalSramUsage1 struct {
		Val c.Uint32T
	}
	InternalSramUsage2 struct {
		Val c.Uint32T
	}
	InternalSramUsage3 struct {
		Val c.Uint32T
	}
	InternalSramUsage4 struct {
		Val c.Uint32T
	}
	RetentionDisable struct {
		Val c.Uint32T
	}
	CacheTagAccess0 struct {
		Val c.Uint32T
	}
	CacheTagAccess1 struct {
		Val c.Uint32T
	}
	CacheMmuAccess0 struct {
		Val c.Uint32T
	}
	CacheMmuAccess1 struct {
		Val c.Uint32T
	}
	DmaApbperiSpi2PmsConstrain0 struct {
		Val c.Uint32T
	}
	DmaApbperiSpi2PmsConstrain1 struct {
		Val c.Uint32T
	}
	DmaApbperiSpi3PmsConstrain0 struct {
		Val c.Uint32T
	}
	DmaApbperiSpi3PmsConstrain1 struct {
		Val c.Uint32T
	}
	DmaApbperiUhci0PmsConstrain0 struct {
		Val c.Uint32T
	}
	DmaApbperiUhci0PmsConstrain1 struct {
		Val c.Uint32T
	}
	DmaApbperiI2s0PmsConstrain0 struct {
		Val c.Uint32T
	}
	DmaApbperiI2s0PmsConstrain1 struct {
		Val c.Uint32T
	}
	DmaApbperiI2s1PmsConstrain0 struct {
		Val c.Uint32T
	}
	DmaApbperiI2s1PmsConstrain1 struct {
		Val c.Uint32T
	}
	DmaApbperiMacPmsConstrain0 struct {
		Val c.Uint32T
	}
	DmaApbperiMacPmsConstrain1 struct {
		Val c.Uint32T
	}
	DmaApbperiBackupPmsConstrain0 struct {
		Val c.Uint32T
	}
	DmaApbperiBackupPmsConstrain1 struct {
		Val c.Uint32T
	}
	DmaApbperiAesPmsConstrain0 struct {
		Val c.Uint32T
	}
	DmaApbperiAesPmsConstrain1 struct {
		Val c.Uint32T
	}
	DmaApbperiShaPmsConstrain0 struct {
		Val c.Uint32T
	}
	DmaApbperiShaPmsConstrain1 struct {
		Val c.Uint32T
	}
	DmaApbperiAdcDacPmsConstrain0 struct {
		Val c.Uint32T
	}
	DmaApbperiAdcDacPmsConstrain1 struct {
		Val c.Uint32T
	}
	DmaApbperiRmtPmsConstrain0 struct {
		Val c.Uint32T
	}
	DmaApbperiRmtPmsConstrain1 struct {
		Val c.Uint32T
	}
	DmaApbperiLcdCamPmsConstrain0 struct {
		Val c.Uint32T
	}
	DmaApbperiLcdCamPmsConstrain1 struct {
		Val c.Uint32T
	}
	DmaApbperiUsbPmsConstrain0 struct {
		Val c.Uint32T
	}
	DmaApbperiUsbPmsConstrain1 struct {
		Val c.Uint32T
	}
	DmaApbperiLcPmsConstrain0 struct {
		Val c.Uint32T
	}
	DmaApbperiLcPmsConstrain1 struct {
		Val c.Uint32T
	}
	DmaApbperiSdioPmsConstrain0 struct {
		Val c.Uint32T
	}
	DmaApbperiSdioPmsConstrain1 struct {
		Val c.Uint32T
	}
	DmaApbperiPmsMonitor0 struct {
		Val c.Uint32T
	}
	DmaApbperiPmsMonitor1 struct {
		Val c.Uint32T
	}
	DmaApbperiPmsMonitor2 struct {
		Val c.Uint32T
	}
	DmaApbperiPmsMonitor3 struct {
		Val c.Uint32T
	}
	CoreXIram0Dram0DmaSplitLineConstrain0 struct {
		Val c.Uint32T
	}
	CoreXIram0Dram0DmaSplitLineConstrain1 struct {
		Val c.Uint32T
	}
	CoreXIram0Dram0DmaSplitLineConstrain2 struct {
		Val c.Uint32T
	}
	CoreXIram0Dram0DmaSplitLineConstrain3 struct {
		Val c.Uint32T
	}
	CoreXIram0Dram0DmaSplitLineConstrain4 struct {
		Val c.Uint32T
	}
	CoreXIram0Dram0DmaSplitLineConstrain5 struct {
		Val c.Uint32T
	}
	CoreXIram0PmsConstrain0 struct {
		Val c.Uint32T
	}
	CoreXIram0PmsConstrain1 struct {
		Val c.Uint32T
	}
	CoreXIram0PmsConstrain2 struct {
		Val c.Uint32T
	}
	Core0Iram0PmsMonitor0 struct {
		Val c.Uint32T
	}
	Core0Iram0PmsMonitor1 struct {
		Val c.Uint32T
	}
	Core0Iram0PmsMonitor2 struct {
		Val c.Uint32T
	}
	Core1Iram0PmsMonitor0 struct {
		Val c.Uint32T
	}
	Core1Iram0PmsMonitor1 struct {
		Val c.Uint32T
	}
	Core1Iram0PmsMonitor2 struct {
		Val c.Uint32T
	}
	CoreXDram0PmsConstrain0 struct {
		Val c.Uint32T
	}
	CoreXDram0PmsConstrain1 struct {
		Val c.Uint32T
	}
	Core0Dram0PmsMonitor0 struct {
		Val c.Uint32T
	}
	Core0Dram0PmsMonitor1 struct {
		Val c.Uint32T
	}
	Core0Dram0PmsMonitor2 struct {
		Val c.Uint32T
	}
	Core0Dram0PmsMonitor3 struct {
		Val c.Uint32T
	}
	Core1Dram0PmsMonitor0 struct {
		Val c.Uint32T
	}
	Core1Dram0PmsMonitor1 struct {
		Val c.Uint32T
	}
	Core1Dram0PmsMonitor2 struct {
		Val c.Uint32T
	}
	Core1Dram0PmsMonitor3 struct {
		Val c.Uint32T
	}
	Core0PifPmsConstrain0 struct {
		Val c.Uint32T
	}
	Core0PifPmsConstrain1 struct {
		Val c.Uint32T
	}
	Core0PifPmsConstrain2 struct {
		Val c.Uint32T
	}
	Core0PifPmsConstrain3 struct {
		Val c.Uint32T
	}
	Core0PifPmsConstrain4 struct {
		Val c.Uint32T
	}
	Core0PifPmsConstrain5 struct {
		Val c.Uint32T
	}
	Core0PifPmsConstrain6 struct {
		Val c.Uint32T
	}
	Core0PifPmsConstrain7 struct {
		Val c.Uint32T
	}
	Core0PifPmsConstrain8 struct {
		Val c.Uint32T
	}
	Core0PifPmsConstrain9 struct {
		Val c.Uint32T
	}
	Core0PifPmsConstrain10 struct {
		Val c.Uint32T
	}
	Core0PifPmsConstrain11 struct {
		Val c.Uint32T
	}
	Core0PifPmsConstrain12 struct {
		Val c.Uint32T
	}
	Core0PifPmsConstrain13 struct {
		Val c.Uint32T
	}
	Core0PifPmsConstrain14 struct {
		Val c.Uint32T
	}
	Core0RegionPmsConstrain0 struct {
		Val c.Uint32T
	}
	Core0RegionPmsConstrain1 struct {
		Val c.Uint32T
	}
	Core0RegionPmsConstrain2 struct {
		Val c.Uint32T
	}
	Core0RegionPmsConstrain3 struct {
		Val c.Uint32T
	}
	Core0RegionPmsConstrain4 struct {
		Val c.Uint32T
	}
	Core0RegionPmsConstrain5 struct {
		Val c.Uint32T
	}
	Core0RegionPmsConstrain6 struct {
		Val c.Uint32T
	}
	Core0RegionPmsConstrain7 struct {
		Val c.Uint32T
	}
	Core0RegionPmsConstrain8 struct {
		Val c.Uint32T
	}
	Core0RegionPmsConstrain9 struct {
		Val c.Uint32T
	}
	Core0RegionPmsConstrain10 struct {
		Val c.Uint32T
	}
	Core0RegionPmsConstrain11 struct {
		Val c.Uint32T
	}
	Core0RegionPmsConstrain12 struct {
		Val c.Uint32T
	}
	Core0RegionPmsConstrain13 struct {
		Val c.Uint32T
	}
	Core0RegionPmsConstrain14 struct {
		Val c.Uint32T
	}
	Core0PifPmsMonitor0 struct {
		Val c.Uint32T
	}
	Core0PifPmsMonitor1 struct {
		Val c.Uint32T
	}
	Core0PifPmsMonitor2 struct {
		Val c.Uint32T
	}
	Core0PifPmsMonitor3 c.Uint32T
	Core0PifPmsMonitor4 struct {
		Val c.Uint32T
	}
	Core0PifPmsMonitor5 struct {
		Val c.Uint32T
	}
	Core0PifPmsMonitor6      c.Uint32T
	Core0VecbaseOverrideLock struct {
		Val c.Uint32T
	}
	Core0VecbaseOverride0 struct {
		Val c.Uint32T
	}
	Core0VecbaseOverride1 struct {
		Val c.Uint32T
	}
	Core0VecbaseOverride2 struct {
		Val c.Uint32T
	}
	Core0ToomanyexceptionsMOverride0 struct {
		Val c.Uint32T
	}
	Core0ToomanyexceptionsMOverride1 struct {
		Val c.Uint32T
	}
	Core1PifPmsConstrain0 struct {
		Val c.Uint32T
	}
	Core1PifPmsConstrain1 struct {
		Val c.Uint32T
	}
	Core1PifPmsConstrain2 struct {
		Val c.Uint32T
	}
	Core1PifPmsConstrain3 struct {
		Val c.Uint32T
	}
	Core1PifPmsConstrain4 struct {
		Val c.Uint32T
	}
	Core1PifPmsConstrain5 struct {
		Val c.Uint32T
	}
	Core1PifPmsConstrain6 struct {
		Val c.Uint32T
	}
	Core1PifPmsConstrain7 struct {
		Val c.Uint32T
	}
	Core1PifPmsConstrain8 struct {
		Val c.Uint32T
	}
	Core1PifPmsConstrain9 struct {
		Val c.Uint32T
	}
	Core1PifPmsConstrain10 struct {
		Val c.Uint32T
	}
	Core1PifPmsConstrain11 struct {
		Val c.Uint32T
	}
	Core1PifPmsConstrain12 struct {
		Val c.Uint32T
	}
	Core1PifPmsConstrain13 struct {
		Val c.Uint32T
	}
	Core1PifPmsConstrain14 struct {
		Val c.Uint32T
	}
	Core1RegionPmsConstrain0 struct {
		Val c.Uint32T
	}
	Core1RegionPmsConstrain1 struct {
		Val c.Uint32T
	}
	Core1RegionPmsConstrain2 struct {
		Val c.Uint32T
	}
	Core1RegionPmsConstrain3 struct {
		Val c.Uint32T
	}
	Core1RegionPmsConstrain4 struct {
		Val c.Uint32T
	}
	Core1RegionPmsConstrain5 struct {
		Val c.Uint32T
	}
	Core1RegionPmsConstrain6 struct {
		Val c.Uint32T
	}
	Core1RegionPmsConstrain7 struct {
		Val c.Uint32T
	}
	Core1RegionPmsConstrain8 struct {
		Val c.Uint32T
	}
	Core1RegionPmsConstrain9 struct {
		Val c.Uint32T
	}
	Core1RegionPmsConstrain10 struct {
		Val c.Uint32T
	}
	Core1RegionPmsConstrain11 struct {
		Val c.Uint32T
	}
	Core1RegionPmsConstrain12 struct {
		Val c.Uint32T
	}
	Core1RegionPmsConstrain13 struct {
		Val c.Uint32T
	}
	Core1RegionPmsConstrain14 struct {
		Val c.Uint32T
	}
	Core1PifPmsMonitor0 struct {
		Val c.Uint32T
	}
	Core1PifPmsMonitor1 struct {
		Val c.Uint32T
	}
	Core1PifPmsMonitor2 struct {
		Val c.Uint32T
	}
	Core1PifPmsMonitor3 c.Uint32T
	Core1PifPmsMonitor4 struct {
		Val c.Uint32T
	}
	Core1PifPmsMonitor5 struct {
		Val c.Uint32T
	}
	Core1PifPmsMonitor6      c.Uint32T
	Core1VecbaseOverrideLock struct {
		Val c.Uint32T
	}
	Core1VecbaseOverride0 struct {
		Val c.Uint32T
	}
	Core1VecbaseOverride1 struct {
		Val c.Uint32T
	}
	Core1VecbaseOverride2 struct {
		Val c.Uint32T
	}
	Core1ToomanyexceptionsMOverride0 struct {
		Val c.Uint32T
	}
	Core1ToomanyexceptionsMOverride1 struct {
		Val c.Uint32T
	}
	BackupBusPmsConstrain0 struct {
		Val c.Uint32T
	}
	BackupBusPmsConstrain1 struct {
		Val c.Uint32T
	}
	BackupBusPmsConstrain2 struct {
		Val c.Uint32T
	}
	BackupBusPmsConstrain3 struct {
		Val c.Uint32T
	}
	BackupBusPmsConstrain4 struct {
		Val c.Uint32T
	}
	BackupBusPmsConstrain5 struct {
		Val c.Uint32T
	}
	BackupBusPmsConstrain6 struct {
		Val c.Uint32T
	}
	BackupBusPmsMonitor0 struct {
		Val c.Uint32T
	}
	BackupBusPmsMonitor1 struct {
		Val c.Uint32T
	}
	BackupBusPmsMonitor2 struct {
		Val c.Uint32T
	}
	BackupBusPmsMonitor3 c.Uint32T
	EdmaBoundaryLock     struct {
		Val c.Uint32T
	}
	EdmaBoundary0 struct {
		Val c.Uint32T
	}
	EdmaBoundary1 struct {
		Val c.Uint32T
	}
	EdmaBoundary2 struct {
		Val c.Uint32T
	}
	EdmaPmsSpi2Lock struct {
		Val c.Uint32T
	}
	EdmaPmsSpi2 struct {
		Val c.Uint32T
	}
	EdmaPmsSpi3Lock struct {
		Val c.Uint32T
	}
	EdmaPmsSpi3 struct {
		Val c.Uint32T
	}
	EdmaPmsUhci0Lock struct {
		Val c.Uint32T
	}
	EdmaPmsUhci0 struct {
		Val c.Uint32T
	}
	EdmaPmsI2s0Lock struct {
		Val c.Uint32T
	}
	EdmaPmsI2s0 struct {
		Val c.Uint32T
	}
	EdmaPmsI2s1Lock struct {
		Val c.Uint32T
	}
	EdmaPmsI2s1 struct {
		Val c.Uint32T
	}
	EdmaPmsLcdCamLock struct {
		Val c.Uint32T
	}
	EdmaPmsLcdCam struct {
		Val c.Uint32T
	}
	EdmaPmsAesLock struct {
		Val c.Uint32T
	}
	EdmaPmsAes struct {
		Val c.Uint32T
	}
	EdmaPmsShaLock struct {
		Val c.Uint32T
	}
	EdmaPmsSha struct {
		Val c.Uint32T
	}
	EdmaPmsAdcDacLock struct {
		Val c.Uint32T
	}
	EdmaPmsAdcDac struct {
		Val c.Uint32T
	}
	EdmaPmsRmtLock struct {
		Val c.Uint32T
	}
	EdmaPmsRmt struct {
		Val c.Uint32T
	}
	ClockGate struct {
		Val c.Uint32T
	}
	RtcPms struct {
		Val c.Uint32T
	}
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
	Reserved3fc c.Uint32T
	Reserved400 c.Uint32T
	Reserved404 c.Uint32T
	Reserved408 c.Uint32T
	Reserved40c c.Uint32T
	Reserved410 c.Uint32T
	Reserved414 c.Uint32T
	Reserved418 c.Uint32T
	Reserved41c c.Uint32T
	Reserved420 c.Uint32T
	Reserved424 c.Uint32T
	Reserved428 c.Uint32T
	Reserved42c c.Uint32T
	Reserved430 c.Uint32T
	Reserved434 c.Uint32T
	Reserved438 c.Uint32T
	Reserved43c c.Uint32T
	Reserved440 c.Uint32T
	Reserved444 c.Uint32T
	Reserved448 c.Uint32T
	Reserved44c c.Uint32T
	Reserved450 c.Uint32T
	Reserved454 c.Uint32T
	Reserved458 c.Uint32T
	Reserved45c c.Uint32T
	Reserved460 c.Uint32T
	Reserved464 c.Uint32T
	Reserved468 c.Uint32T
	Reserved46c c.Uint32T
	Reserved470 c.Uint32T
	Reserved474 c.Uint32T
	Reserved478 c.Uint32T
	Reserved47c c.Uint32T
	Reserved480 c.Uint32T
	Reserved484 c.Uint32T
	Reserved488 c.Uint32T
	Reserved48c c.Uint32T
	Reserved490 c.Uint32T
	Reserved494 c.Uint32T
	Reserved498 c.Uint32T
	Reserved49c c.Uint32T
	Reserved4a0 c.Uint32T
	Reserved4a4 c.Uint32T
	Reserved4a8 c.Uint32T
	Reserved4ac c.Uint32T
	Reserved4b0 c.Uint32T
	Reserved4b4 c.Uint32T
	Reserved4b8 c.Uint32T
	Reserved4bc c.Uint32T
	Reserved4c0 c.Uint32T
	Reserved4c4 c.Uint32T
	Reserved4c8 c.Uint32T
	Reserved4cc c.Uint32T
	Reserved4d0 c.Uint32T
	Reserved4d4 c.Uint32T
	Reserved4d8 c.Uint32T
	Reserved4dc c.Uint32T
	Reserved4e0 c.Uint32T
	Reserved4e4 c.Uint32T
	Reserved4e8 c.Uint32T
	Reserved4ec c.Uint32T
	Reserved4f0 c.Uint32T
	Reserved4f4 c.Uint32T
	Reserved4f8 c.Uint32T
	Reserved4fc c.Uint32T
	Reserved500 c.Uint32T
	Reserved504 c.Uint32T
	Reserved508 c.Uint32T
	Reserved50c c.Uint32T
	Reserved510 c.Uint32T
	Reserved514 c.Uint32T
	Reserved518 c.Uint32T
	Reserved51c c.Uint32T
	Reserved520 c.Uint32T
	Reserved524 c.Uint32T
	Reserved528 c.Uint32T
	Reserved52c c.Uint32T
	Reserved530 c.Uint32T
	Reserved534 c.Uint32T
	Reserved538 c.Uint32T
	Reserved53c c.Uint32T
	Reserved540 c.Uint32T
	Reserved544 c.Uint32T
	Reserved548 c.Uint32T
	Reserved54c c.Uint32T
	Reserved550 c.Uint32T
	Reserved554 c.Uint32T
	Reserved558 c.Uint32T
	Reserved55c c.Uint32T
	Reserved560 c.Uint32T
	Reserved564 c.Uint32T
	Reserved568 c.Uint32T
	Reserved56c c.Uint32T
	Reserved570 c.Uint32T
	Reserved574 c.Uint32T
	Reserved578 c.Uint32T
	Reserved57c c.Uint32T
	Reserved580 c.Uint32T
	Reserved584 c.Uint32T
	Reserved588 c.Uint32T
	Reserved58c c.Uint32T
	Reserved590 c.Uint32T
	Reserved594 c.Uint32T
	Reserved598 c.Uint32T
	Reserved59c c.Uint32T
	Reserved5a0 c.Uint32T
	Reserved5a4 c.Uint32T
	Reserved5a8 c.Uint32T
	Reserved5ac c.Uint32T
	Reserved5b0 c.Uint32T
	Reserved5b4 c.Uint32T
	Reserved5b8 c.Uint32T
	Reserved5bc c.Uint32T
	Reserved5c0 c.Uint32T
	Reserved5c4 c.Uint32T
	Reserved5c8 c.Uint32T
	Reserved5cc c.Uint32T
	Reserved5d0 c.Uint32T
	Reserved5d4 c.Uint32T
	Reserved5d8 c.Uint32T
	Reserved5dc c.Uint32T
	Reserved5e0 c.Uint32T
	Reserved5e4 c.Uint32T
	Reserved5e8 c.Uint32T
	Reserved5ec c.Uint32T
	Reserved5f0 c.Uint32T
	Reserved5f4 c.Uint32T
	Reserved5f8 c.Uint32T
	Reserved5fc c.Uint32T
	Reserved600 c.Uint32T
	Reserved604 c.Uint32T
	Reserved608 c.Uint32T
	Reserved60c c.Uint32T
	Reserved610 c.Uint32T
	Reserved614 c.Uint32T
	Reserved618 c.Uint32T
	Reserved61c c.Uint32T
	Reserved620 c.Uint32T
	Reserved624 c.Uint32T
	Reserved628 c.Uint32T
	Reserved62c c.Uint32T
	Reserved630 c.Uint32T
	Reserved634 c.Uint32T
	Reserved638 c.Uint32T
	Reserved63c c.Uint32T
	Reserved640 c.Uint32T
	Reserved644 c.Uint32T
	Reserved648 c.Uint32T
	Reserved64c c.Uint32T
	Reserved650 c.Uint32T
	Reserved654 c.Uint32T
	Reserved658 c.Uint32T
	Reserved65c c.Uint32T
	Reserved660 c.Uint32T
	Reserved664 c.Uint32T
	Reserved668 c.Uint32T
	Reserved66c c.Uint32T
	Reserved670 c.Uint32T
	Reserved674 c.Uint32T
	Reserved678 c.Uint32T
	Reserved67c c.Uint32T
	Reserved680 c.Uint32T
	Reserved684 c.Uint32T
	Reserved688 c.Uint32T
	Reserved68c c.Uint32T
	Reserved690 c.Uint32T
	Reserved694 c.Uint32T
	Reserved698 c.Uint32T
	Reserved69c c.Uint32T
	Reserved6a0 c.Uint32T
	Reserved6a4 c.Uint32T
	Reserved6a8 c.Uint32T
	Reserved6ac c.Uint32T
	Reserved6b0 c.Uint32T
	Reserved6b4 c.Uint32T
	Reserved6b8 c.Uint32T
	Reserved6bc c.Uint32T
	Reserved6c0 c.Uint32T
	Reserved6c4 c.Uint32T
	Reserved6c8 c.Uint32T
	Reserved6cc c.Uint32T
	Reserved6d0 c.Uint32T
	Reserved6d4 c.Uint32T
	Reserved6d8 c.Uint32T
	Reserved6dc c.Uint32T
	Reserved6e0 c.Uint32T
	Reserved6e4 c.Uint32T
	Reserved6e8 c.Uint32T
	Reserved6ec c.Uint32T
	Reserved6f0 c.Uint32T
	Reserved6f4 c.Uint32T
	Reserved6f8 c.Uint32T
	Reserved6fc c.Uint32T
	Reserved700 c.Uint32T
	Reserved704 c.Uint32T
	Reserved708 c.Uint32T
	Reserved70c c.Uint32T
	Reserved710 c.Uint32T
	Reserved714 c.Uint32T
	Reserved718 c.Uint32T
	Reserved71c c.Uint32T
	Reserved720 c.Uint32T
	Reserved724 c.Uint32T
	Reserved728 c.Uint32T
	Reserved72c c.Uint32T
	Reserved730 c.Uint32T
	Reserved734 c.Uint32T
	Reserved738 c.Uint32T
	Reserved73c c.Uint32T
	Reserved740 c.Uint32T
	Reserved744 c.Uint32T
	Reserved748 c.Uint32T
	Reserved74c c.Uint32T
	Reserved750 c.Uint32T
	Reserved754 c.Uint32T
	Reserved758 c.Uint32T
	Reserved75c c.Uint32T
	Reserved760 c.Uint32T
	Reserved764 c.Uint32T
	Reserved768 c.Uint32T
	Reserved76c c.Uint32T
	Reserved770 c.Uint32T
	Reserved774 c.Uint32T
	Reserved778 c.Uint32T
	Reserved77c c.Uint32T
	Reserved780 c.Uint32T
	Reserved784 c.Uint32T
	Reserved788 c.Uint32T
	Reserved78c c.Uint32T
	Reserved790 c.Uint32T
	Reserved794 c.Uint32T
	Reserved798 c.Uint32T
	Reserved79c c.Uint32T
	Reserved7a0 c.Uint32T
	Reserved7a4 c.Uint32T
	Reserved7a8 c.Uint32T
	Reserved7ac c.Uint32T
	Reserved7b0 c.Uint32T
	Reserved7b4 c.Uint32T
	Reserved7b8 c.Uint32T
	Reserved7bc c.Uint32T
	Reserved7c0 c.Uint32T
	Reserved7c4 c.Uint32T
	Reserved7c8 c.Uint32T
	Reserved7cc c.Uint32T
	Reserved7d0 c.Uint32T
	Reserved7d4 c.Uint32T
	Reserved7d8 c.Uint32T
	Reserved7dc c.Uint32T
	Reserved7e0 c.Uint32T
	Reserved7e4 c.Uint32T
	Reserved7e8 c.Uint32T
	Reserved7ec c.Uint32T
	Reserved7f0 c.Uint32T
	Reserved7f4 c.Uint32T
	Reserved7f8 c.Uint32T
	Reserved7fc c.Uint32T
	Reserved800 c.Uint32T
	Reserved804 c.Uint32T
	Reserved808 c.Uint32T
	Reserved80c c.Uint32T
	Reserved810 c.Uint32T
	Reserved814 c.Uint32T
	Reserved818 c.Uint32T
	Reserved81c c.Uint32T
	Reserved820 c.Uint32T
	Reserved824 c.Uint32T
	Reserved828 c.Uint32T
	Reserved82c c.Uint32T
	Reserved830 c.Uint32T
	Reserved834 c.Uint32T
	Reserved838 c.Uint32T
	Reserved83c c.Uint32T
	Reserved840 c.Uint32T
	Reserved844 c.Uint32T
	Reserved848 c.Uint32T
	Reserved84c c.Uint32T
	Reserved850 c.Uint32T
	Reserved854 c.Uint32T
	Reserved858 c.Uint32T
	Reserved85c c.Uint32T
	Reserved860 c.Uint32T
	Reserved864 c.Uint32T
	Reserved868 c.Uint32T
	Reserved86c c.Uint32T
	Reserved870 c.Uint32T
	Reserved874 c.Uint32T
	Reserved878 c.Uint32T
	Reserved87c c.Uint32T
	Reserved880 c.Uint32T
	Reserved884 c.Uint32T
	Reserved888 c.Uint32T
	Reserved88c c.Uint32T
	Reserved890 c.Uint32T
	Reserved894 c.Uint32T
	Reserved898 c.Uint32T
	Reserved89c c.Uint32T
	Reserved8a0 c.Uint32T
	Reserved8a4 c.Uint32T
	Reserved8a8 c.Uint32T
	Reserved8ac c.Uint32T
	Reserved8b0 c.Uint32T
	Reserved8b4 c.Uint32T
	Reserved8b8 c.Uint32T
	Reserved8bc c.Uint32T
	Reserved8c0 c.Uint32T
	Reserved8c4 c.Uint32T
	Reserved8c8 c.Uint32T
	Reserved8cc c.Uint32T
	Reserved8d0 c.Uint32T
	Reserved8d4 c.Uint32T
	Reserved8d8 c.Uint32T
	Reserved8dc c.Uint32T
	Reserved8e0 c.Uint32T
	Reserved8e4 c.Uint32T
	Reserved8e8 c.Uint32T
	Reserved8ec c.Uint32T
	Reserved8f0 c.Uint32T
	Reserved8f4 c.Uint32T
	Reserved8f8 c.Uint32T
	Reserved8fc c.Uint32T
	Reserved900 c.Uint32T
	Reserved904 c.Uint32T
	Reserved908 c.Uint32T
	Reserved90c c.Uint32T
	Reserved910 c.Uint32T
	Reserved914 c.Uint32T
	Reserved918 c.Uint32T
	Reserved91c c.Uint32T
	Reserved920 c.Uint32T
	Reserved924 c.Uint32T
	Reserved928 c.Uint32T
	Reserved92c c.Uint32T
	Reserved930 c.Uint32T
	Reserved934 c.Uint32T
	Reserved938 c.Uint32T
	Reserved93c c.Uint32T
	Reserved940 c.Uint32T
	Reserved944 c.Uint32T
	Reserved948 c.Uint32T
	Reserved94c c.Uint32T
	Reserved950 c.Uint32T
	Reserved954 c.Uint32T
	Reserved958 c.Uint32T
	Reserved95c c.Uint32T
	Reserved960 c.Uint32T
	Reserved964 c.Uint32T
	Reserved968 c.Uint32T
	Reserved96c c.Uint32T
	Reserved970 c.Uint32T
	Reserved974 c.Uint32T
	Reserved978 c.Uint32T
	Reserved97c c.Uint32T
	Reserved980 c.Uint32T
	Reserved984 c.Uint32T
	Reserved988 c.Uint32T
	Reserved98c c.Uint32T
	Reserved990 c.Uint32T
	Reserved994 c.Uint32T
	Reserved998 c.Uint32T
	Reserved99c c.Uint32T
	Reserved9a0 c.Uint32T
	Reserved9a4 c.Uint32T
	Reserved9a8 c.Uint32T
	Reserved9ac c.Uint32T
	Reserved9b0 c.Uint32T
	Reserved9b4 c.Uint32T
	Reserved9b8 c.Uint32T
	Reserved9bc c.Uint32T
	Reserved9c0 c.Uint32T
	Reserved9c4 c.Uint32T
	Reserved9c8 c.Uint32T
	Reserved9cc c.Uint32T
	Reserved9d0 c.Uint32T
	Reserved9d4 c.Uint32T
	Reserved9d8 c.Uint32T
	Reserved9dc c.Uint32T
	Reserved9e0 c.Uint32T
	Reserved9e4 c.Uint32T
	Reserved9e8 c.Uint32T
	Reserved9ec c.Uint32T
	Reserved9f0 c.Uint32T
	Reserved9f4 c.Uint32T
	Reserved9f8 c.Uint32T
	Reserved9fc c.Uint32T
	ReservedA00 c.Uint32T
	ReservedA04 c.Uint32T
	ReservedA08 c.Uint32T
	ReservedA0c c.Uint32T
	ReservedA10 c.Uint32T
	ReservedA14 c.Uint32T
	ReservedA18 c.Uint32T
	ReservedA1c c.Uint32T
	ReservedA20 c.Uint32T
	ReservedA24 c.Uint32T
	ReservedA28 c.Uint32T
	ReservedA2c c.Uint32T
	ReservedA30 c.Uint32T
	ReservedA34 c.Uint32T
	ReservedA38 c.Uint32T
	ReservedA3c c.Uint32T
	ReservedA40 c.Uint32T
	ReservedA44 c.Uint32T
	ReservedA48 c.Uint32T
	ReservedA4c c.Uint32T
	ReservedA50 c.Uint32T
	ReservedA54 c.Uint32T
	ReservedA58 c.Uint32T
	ReservedA5c c.Uint32T
	ReservedA60 c.Uint32T
	ReservedA64 c.Uint32T
	ReservedA68 c.Uint32T
	ReservedA6c c.Uint32T
	ReservedA70 c.Uint32T
	ReservedA74 c.Uint32T
	ReservedA78 c.Uint32T
	ReservedA7c c.Uint32T
	ReservedA80 c.Uint32T
	ReservedA84 c.Uint32T
	ReservedA88 c.Uint32T
	ReservedA8c c.Uint32T
	ReservedA90 c.Uint32T
	ReservedA94 c.Uint32T
	ReservedA98 c.Uint32T
	ReservedA9c c.Uint32T
	ReservedAa0 c.Uint32T
	ReservedAa4 c.Uint32T
	ReservedAa8 c.Uint32T
	ReservedAac c.Uint32T
	ReservedAb0 c.Uint32T
	ReservedAb4 c.Uint32T
	ReservedAb8 c.Uint32T
	ReservedAbc c.Uint32T
	ReservedAc0 c.Uint32T
	ReservedAc4 c.Uint32T
	ReservedAc8 c.Uint32T
	ReservedAcc c.Uint32T
	ReservedAd0 c.Uint32T
	ReservedAd4 c.Uint32T
	ReservedAd8 c.Uint32T
	ReservedAdc c.Uint32T
	ReservedAe0 c.Uint32T
	ReservedAe4 c.Uint32T
	ReservedAe8 c.Uint32T
	ReservedAec c.Uint32T
	ReservedAf0 c.Uint32T
	ReservedAf4 c.Uint32T
	ReservedAf8 c.Uint32T
	ReservedAfc c.Uint32T
	ReservedB00 c.Uint32T
	ReservedB04 c.Uint32T
	ReservedB08 c.Uint32T
	ReservedB0c c.Uint32T
	ReservedB10 c.Uint32T
	ReservedB14 c.Uint32T
	ReservedB18 c.Uint32T
	ReservedB1c c.Uint32T
	ReservedB20 c.Uint32T
	ReservedB24 c.Uint32T
	ReservedB28 c.Uint32T
	ReservedB2c c.Uint32T
	ReservedB30 c.Uint32T
	ReservedB34 c.Uint32T
	ReservedB38 c.Uint32T
	ReservedB3c c.Uint32T
	ReservedB40 c.Uint32T
	ReservedB44 c.Uint32T
	ReservedB48 c.Uint32T
	ReservedB4c c.Uint32T
	ReservedB50 c.Uint32T
	ReservedB54 c.Uint32T
	ReservedB58 c.Uint32T
	ReservedB5c c.Uint32T
	ReservedB60 c.Uint32T
	ReservedB64 c.Uint32T
	ReservedB68 c.Uint32T
	ReservedB6c c.Uint32T
	ReservedB70 c.Uint32T
	ReservedB74 c.Uint32T
	ReservedB78 c.Uint32T
	ReservedB7c c.Uint32T
	ReservedB80 c.Uint32T
	ReservedB84 c.Uint32T
	ReservedB88 c.Uint32T
	ReservedB8c c.Uint32T
	ReservedB90 c.Uint32T
	ReservedB94 c.Uint32T
	ReservedB98 c.Uint32T
	ReservedB9c c.Uint32T
	ReservedBa0 c.Uint32T
	ReservedBa4 c.Uint32T
	ReservedBa8 c.Uint32T
	ReservedBac c.Uint32T
	ReservedBb0 c.Uint32T
	ReservedBb4 c.Uint32T
	ReservedBb8 c.Uint32T
	ReservedBbc c.Uint32T
	ReservedBc0 c.Uint32T
	ReservedBc4 c.Uint32T
	ReservedBc8 c.Uint32T
	ReservedBcc c.Uint32T
	ReservedBd0 c.Uint32T
	ReservedBd4 c.Uint32T
	ReservedBd8 c.Uint32T
	ReservedBdc c.Uint32T
	ReservedBe0 c.Uint32T
	ReservedBe4 c.Uint32T
	ReservedBe8 c.Uint32T
	ReservedBec c.Uint32T
	ReservedBf0 c.Uint32T
	ReservedBf4 c.Uint32T
	ReservedBf8 c.Uint32T
	ReservedBfc c.Uint32T
	ReservedC00 c.Uint32T
	ReservedC04 c.Uint32T
	ReservedC08 c.Uint32T
	ReservedC0c c.Uint32T
	ReservedC10 c.Uint32T
	ReservedC14 c.Uint32T
	ReservedC18 c.Uint32T
	ReservedC1c c.Uint32T
	ReservedC20 c.Uint32T
	ReservedC24 c.Uint32T
	ReservedC28 c.Uint32T
	ReservedC2c c.Uint32T
	ReservedC30 c.Uint32T
	ReservedC34 c.Uint32T
	ReservedC38 c.Uint32T
	ReservedC3c c.Uint32T
	ReservedC40 c.Uint32T
	ReservedC44 c.Uint32T
	ReservedC48 c.Uint32T
	ReservedC4c c.Uint32T
	ReservedC50 c.Uint32T
	ReservedC54 c.Uint32T
	ReservedC58 c.Uint32T
	ReservedC5c c.Uint32T
	ReservedC60 c.Uint32T
	ReservedC64 c.Uint32T
	ReservedC68 c.Uint32T
	ReservedC6c c.Uint32T
	ReservedC70 c.Uint32T
	ReservedC74 c.Uint32T
	ReservedC78 c.Uint32T
	ReservedC7c c.Uint32T
	ReservedC80 c.Uint32T
	ReservedC84 c.Uint32T
	ReservedC88 c.Uint32T
	ReservedC8c c.Uint32T
	ReservedC90 c.Uint32T
	ReservedC94 c.Uint32T
	ReservedC98 c.Uint32T
	ReservedC9c c.Uint32T
	ReservedCa0 c.Uint32T
	ReservedCa4 c.Uint32T
	ReservedCa8 c.Uint32T
	ReservedCac c.Uint32T
	ReservedCb0 c.Uint32T
	ReservedCb4 c.Uint32T
	ReservedCb8 c.Uint32T
	ReservedCbc c.Uint32T
	ReservedCc0 c.Uint32T
	ReservedCc4 c.Uint32T
	ReservedCc8 c.Uint32T
	ReservedCcc c.Uint32T
	ReservedCd0 c.Uint32T
	ReservedCd4 c.Uint32T
	ReservedCd8 c.Uint32T
	ReservedCdc c.Uint32T
	ReservedCe0 c.Uint32T
	ReservedCe4 c.Uint32T
	ReservedCe8 c.Uint32T
	ReservedCec c.Uint32T
	ReservedCf0 c.Uint32T
	ReservedCf4 c.Uint32T
	ReservedCf8 c.Uint32T
	ReservedCfc c.Uint32T
	ReservedD00 c.Uint32T
	ReservedD04 c.Uint32T
	ReservedD08 c.Uint32T
	ReservedD0c c.Uint32T
	ReservedD10 c.Uint32T
	ReservedD14 c.Uint32T
	ReservedD18 c.Uint32T
	ReservedD1c c.Uint32T
	ReservedD20 c.Uint32T
	ReservedD24 c.Uint32T
	ReservedD28 c.Uint32T
	ReservedD2c c.Uint32T
	ReservedD30 c.Uint32T
	ReservedD34 c.Uint32T
	ReservedD38 c.Uint32T
	ReservedD3c c.Uint32T
	ReservedD40 c.Uint32T
	ReservedD44 c.Uint32T
	ReservedD48 c.Uint32T
	ReservedD4c c.Uint32T
	ReservedD50 c.Uint32T
	ReservedD54 c.Uint32T
	ReservedD58 c.Uint32T
	ReservedD5c c.Uint32T
	ReservedD60 c.Uint32T
	ReservedD64 c.Uint32T
	ReservedD68 c.Uint32T
	ReservedD6c c.Uint32T
	ReservedD70 c.Uint32T
	ReservedD74 c.Uint32T
	ReservedD78 c.Uint32T
	ReservedD7c c.Uint32T
	ReservedD80 c.Uint32T
	ReservedD84 c.Uint32T
	ReservedD88 c.Uint32T
	ReservedD8c c.Uint32T
	ReservedD90 c.Uint32T
	ReservedD94 c.Uint32T
	ReservedD98 c.Uint32T
	ReservedD9c c.Uint32T
	ReservedDa0 c.Uint32T
	ReservedDa4 c.Uint32T
	ReservedDa8 c.Uint32T
	ReservedDac c.Uint32T
	ReservedDb0 c.Uint32T
	ReservedDb4 c.Uint32T
	ReservedDb8 c.Uint32T
	ReservedDbc c.Uint32T
	ReservedDc0 c.Uint32T
	ReservedDc4 c.Uint32T
	ReservedDc8 c.Uint32T
	ReservedDcc c.Uint32T
	ReservedDd0 c.Uint32T
	ReservedDd4 c.Uint32T
	ReservedDd8 c.Uint32T
	ReservedDdc c.Uint32T
	ReservedDe0 c.Uint32T
	ReservedDe4 c.Uint32T
	ReservedDe8 c.Uint32T
	ReservedDec c.Uint32T
	ReservedDf0 c.Uint32T
	ReservedDf4 c.Uint32T
	ReservedDf8 c.Uint32T
	ReservedDfc c.Uint32T
	ReservedE00 c.Uint32T
	ReservedE04 c.Uint32T
	ReservedE08 c.Uint32T
	ReservedE0c c.Uint32T
	ReservedE10 c.Uint32T
	ReservedE14 c.Uint32T
	ReservedE18 c.Uint32T
	ReservedE1c c.Uint32T
	ReservedE20 c.Uint32T
	ReservedE24 c.Uint32T
	ReservedE28 c.Uint32T
	ReservedE2c c.Uint32T
	ReservedE30 c.Uint32T
	ReservedE34 c.Uint32T
	ReservedE38 c.Uint32T
	ReservedE3c c.Uint32T
	ReservedE40 c.Uint32T
	ReservedE44 c.Uint32T
	ReservedE48 c.Uint32T
	ReservedE4c c.Uint32T
	ReservedE50 c.Uint32T
	ReservedE54 c.Uint32T
	ReservedE58 c.Uint32T
	ReservedE5c c.Uint32T
	ReservedE60 c.Uint32T
	ReservedE64 c.Uint32T
	ReservedE68 c.Uint32T
	ReservedE6c c.Uint32T
	ReservedE70 c.Uint32T
	ReservedE74 c.Uint32T
	ReservedE78 c.Uint32T
	ReservedE7c c.Uint32T
	ReservedE80 c.Uint32T
	ReservedE84 c.Uint32T
	ReservedE88 c.Uint32T
	ReservedE8c c.Uint32T
	ReservedE90 c.Uint32T
	ReservedE94 c.Uint32T
	ReservedE98 c.Uint32T
	ReservedE9c c.Uint32T
	ReservedEa0 c.Uint32T
	ReservedEa4 c.Uint32T
	ReservedEa8 c.Uint32T
	ReservedEac c.Uint32T
	ReservedEb0 c.Uint32T
	ReservedEb4 c.Uint32T
	ReservedEb8 c.Uint32T
	ReservedEbc c.Uint32T
	ReservedEc0 c.Uint32T
	ReservedEc4 c.Uint32T
	ReservedEc8 c.Uint32T
	ReservedEcc c.Uint32T
	ReservedEd0 c.Uint32T
	ReservedEd4 c.Uint32T
	ReservedEd8 c.Uint32T
	ReservedEdc c.Uint32T
	ReservedEe0 c.Uint32T
	ReservedEe4 c.Uint32T
	ReservedEe8 c.Uint32T
	ReservedEec c.Uint32T
	ReservedEf0 c.Uint32T
	ReservedEf4 c.Uint32T
	ReservedEf8 c.Uint32T
	ReservedEfc c.Uint32T
	ReservedF00 c.Uint32T
	ReservedF04 c.Uint32T
	ReservedF08 c.Uint32T
	ReservedF0c c.Uint32T
	ReservedF10 c.Uint32T
	ReservedF14 c.Uint32T
	ReservedF18 c.Uint32T
	ReservedF1c c.Uint32T
	ReservedF20 c.Uint32T
	ReservedF24 c.Uint32T
	ReservedF28 c.Uint32T
	ReservedF2c c.Uint32T
	ReservedF30 c.Uint32T
	ReservedF34 c.Uint32T
	ReservedF38 c.Uint32T
	ReservedF3c c.Uint32T
	ReservedF40 c.Uint32T
	ReservedF44 c.Uint32T
	ReservedF48 c.Uint32T
	ReservedF4c c.Uint32T
	ReservedF50 c.Uint32T
	ReservedF54 c.Uint32T
	ReservedF58 c.Uint32T
	ReservedF5c c.Uint32T
	ReservedF60 c.Uint32T
	ReservedF64 c.Uint32T
	ReservedF68 c.Uint32T
	ReservedF6c c.Uint32T
	ReservedF70 c.Uint32T
	ReservedF74 c.Uint32T
	ReservedF78 c.Uint32T
	ReservedF7c c.Uint32T
	ReservedF80 c.Uint32T
	ReservedF84 c.Uint32T
	ReservedF88 c.Uint32T
	ReservedF8c c.Uint32T
	ReservedF90 c.Uint32T
	ReservedF94 c.Uint32T
	ReservedF98 c.Uint32T
	ReservedF9c c.Uint32T
	ReservedFa0 c.Uint32T
	ReservedFa4 c.Uint32T
	ReservedFa8 c.Uint32T
	ReservedFac c.Uint32T
	ReservedFb0 c.Uint32T
	ReservedFb4 c.Uint32T
	ReservedFb8 c.Uint32T
	ReservedFbc c.Uint32T
	ReservedFc0 c.Uint32T
	ReservedFc4 c.Uint32T
	ReservedFc8 c.Uint32T
	ReservedFcc c.Uint32T
	ReservedFd0 c.Uint32T
	ReservedFd4 c.Uint32T
	ReservedFd8 c.Uint32T
	ReservedFdc c.Uint32T
	ReservedFe0 c.Uint32T
	ReservedFe4 c.Uint32T
	ReservedFe8 c.Uint32T
	ReservedFec c.Uint32T
	ReservedFf0 c.Uint32T
	ReservedFf4 c.Uint32T
	ReservedFf8 c.Uint32T
	RegDate     struct {
		Val c.Uint32T
	}
}
type SensitiveDevT SensitiveDevS
