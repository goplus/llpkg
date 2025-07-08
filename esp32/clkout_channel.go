package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type ClockOutChannel c.Int

const (
	CLKOUT_CHANNEL_1   ClockOutChannel = 0
	CLKOUT_CHANNEL_2   ClockOutChannel = 1
	CLKOUT_CHANNEL_3   ClockOutChannel = 2
	CLKOUT_CHANNEL_MAX ClockOutChannel = 3
)

type ClockOutChannelT ClockOutChannel
