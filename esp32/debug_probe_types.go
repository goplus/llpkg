package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const DEBUG_PROBE_MAX_OUTPUT_WIDTH = 16

type DebugProbeSplitU16T c.Int

const (
	DEBUG_PROBE_SPLIT_LOWER16 DebugProbeSplitU16T = 0
	DEBUG_PROBE_SPLIT_UPPER16 DebugProbeSplitU16T = 1
)

type DebugProbeTargetT c.Int
