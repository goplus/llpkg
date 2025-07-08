package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type PcntChannelLevelActionT c.Int

const (
	PCNT_CHANNEL_LEVEL_ACTION_KEEP    PcntChannelLevelActionT = 0
	PCNT_CHANNEL_LEVEL_ACTION_INVERSE PcntChannelLevelActionT = 1
	PCNT_CHANNEL_LEVEL_ACTION_HOLD    PcntChannelLevelActionT = 2
)

type PcntChannelEdgeActionT c.Int

const (
	PCNT_CHANNEL_EDGE_ACTION_HOLD     PcntChannelEdgeActionT = 0
	PCNT_CHANNEL_EDGE_ACTION_INCREASE PcntChannelEdgeActionT = 1
	PCNT_CHANNEL_EDGE_ACTION_DECREASE PcntChannelEdgeActionT = 2
)

type PcntUnitZeroCrossModeT c.Int

const (
	PCNT_UNIT_ZERO_CROSS_POS_ZERO PcntUnitZeroCrossModeT = 0
	PCNT_UNIT_ZERO_CROSS_NEG_ZERO PcntUnitZeroCrossModeT = 1
	PCNT_UNIT_ZERO_CROSS_NEG_POS  PcntUnitZeroCrossModeT = 2
	PCNT_UNIT_ZERO_CROSS_POS_NEG  PcntUnitZeroCrossModeT = 3
	PCNT_UNIT_ZERO_CROSS_INVALID  PcntUnitZeroCrossModeT = 4
)
