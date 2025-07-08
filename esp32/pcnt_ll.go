package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const PCNT_LL_MAX_GLITCH_WIDTH = 1023

type PcntLlWatchEventIdT c.Int

const (
	PCNT_LL_WATCH_EVENT_INVALID    PcntLlWatchEventIdT = -1
	PCNT_LL_WATCH_EVENT_THRES1     PcntLlWatchEventIdT = 0
	PCNT_LL_WATCH_EVENT_THRES0     PcntLlWatchEventIdT = 1
	PCNT_LL_WATCH_EVENT_LOW_LIMIT  PcntLlWatchEventIdT = 2
	PCNT_LL_WATCH_EVENT_HIGH_LIMIT PcntLlWatchEventIdT = 3
	PCNT_LL_WATCH_EVENT_ZERO_CROSS PcntLlWatchEventIdT = 4
	PCNT_LL_WATCH_EVENT_MAX        PcntLlWatchEventIdT = 5
)
