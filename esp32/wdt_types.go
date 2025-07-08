package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type WdtInstT c.Int

const (
	WDT_RWDT  WdtInstT = 0
	WDT_MWDT0 WdtInstT = 1
	WDT_MWDT1 WdtInstT = 2
)

type WdtStageT c.Int

const (
	WDT_STAGE0 WdtStageT = 0
	WDT_STAGE1 WdtStageT = 1
	WDT_STAGE2 WdtStageT = 2
	WDT_STAGE3 WdtStageT = 3
)

type WdtStageActionT c.Int

const (
	WDT_STAGE_ACTION_OFF          WdtStageActionT = 0
	WDT_STAGE_ACTION_INT          WdtStageActionT = 1
	WDT_STAGE_ACTION_RESET_CPU    WdtStageActionT = 2
	WDT_STAGE_ACTION_RESET_SYSTEM WdtStageActionT = 3
	WDT_STAGE_ACTION_RESET_RTC    WdtStageActionT = 4
)

type WdtResetSigLengthT c.Int

const (
	WDT_RESET_SIG_LENGTH_100ns WdtResetSigLengthT = 0
	WDT_RESET_SIG_LENGTH_200ns WdtResetSigLengthT = 1
	WDT_RESET_SIG_LENGTH_300ns WdtResetSigLengthT = 2
	WDT_RESET_SIG_LENGTH_400ns WdtResetSigLengthT = 3
	WDT_RESET_SIG_LENGTH_500ns WdtResetSigLengthT = 4
	WDT_RESET_SIG_LENGTH_800ns WdtResetSigLengthT = 5
	WDT_RESET_SIG_LENGTH_1_6us WdtResetSigLengthT = 6
	WDT_RESET_SIG_LENGTH_3_2us WdtResetSigLengthT = 7
)

type MwdtClockSourceT SocPeriphMwdtClkSrcT
