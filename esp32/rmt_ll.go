package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const RMT_LL_MAX_LOOP_COUNT_PER_BATCH = 1023
const RMT_LL_MAX_FILTER_VALUE = 255
const RMT_LL_MAX_IDLE_VALUE = 32767
const RMT_LL_CHANNEL_CLOCK_MAX_PRESCALE = 256
const RMT_LL_GROUP_CLOCK_MAX_INTEGER_PRESCALE = 256
const RMT_LL_GROUP_CLOCK_MAX_FRACTAL_PRESCALE = 64

type RmtLlMemOwnerT c.Int

const (
	RMT_LL_MEM_OWNER_SW RmtLlMemOwnerT = 0
	RMT_LL_MEM_OWNER_HW RmtLlMemOwnerT = 1
)
