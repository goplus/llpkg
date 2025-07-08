package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const PARLIO_TX_UNIT_MAX_DATA_WIDTH = 0
const PARLIO_RX_UNIT_MAX_DATA_WIDTH = 0

type ParlioSampleEdgeT c.Int

const (
	PARLIO_SAMPLE_EDGE_NEG ParlioSampleEdgeT = 0
	PARLIO_SAMPLE_EDGE_POS ParlioSampleEdgeT = 1
)

type ParlioBitPackOrderT c.Int

const (
	PARLIO_BIT_PACK_ORDER_LSB ParlioBitPackOrderT = 0
	PARLIO_BIT_PACK_ORDER_MSB ParlioBitPackOrderT = 1
)

type ParlioClockSourceT c.Int
