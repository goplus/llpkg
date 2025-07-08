package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type DacChannelT c.Int

const (
	DAC_CHAN_0    DacChannelT = 0
	DAC_CHAN_1    DacChannelT = 1
	DAC_CHANNEL_1 DacChannelT = 0
	DAC_CHANNEL_2 DacChannelT = 1
)

type DacCosineAttenT c.Int

const (
	DAC_COSINE_ATTEN_DEFAULT DacCosineAttenT = 0
	DAC_COSINE_ATTEN_DB_0    DacCosineAttenT = 0
	DAC_COSINE_ATTEN_DB_6    DacCosineAttenT = 1
	DAC_COSINE_ATTEN_DB_12   DacCosineAttenT = 2
	DAC_COSINE_ATTEN_DB_18   DacCosineAttenT = 3
)

type DacCosinePhaseT c.Int

const (
	DAC_COSINE_PHASE_0   DacCosinePhaseT = 2
	DAC_COSINE_PHASE_180 DacCosinePhaseT = 3
)
