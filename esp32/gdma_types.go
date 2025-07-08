package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type GdmaTriggerPeripheralT c.Int

const (
	GDMA_TRIG_PERIPH_M2M    GdmaTriggerPeripheralT = 0
	GDMA_TRIG_PERIPH_UHCI   GdmaTriggerPeripheralT = 1
	GDMA_TRIG_PERIPH_SPI    GdmaTriggerPeripheralT = 2
	GDMA_TRIG_PERIPH_I2S    GdmaTriggerPeripheralT = 3
	GDMA_TRIG_PERIPH_AES    GdmaTriggerPeripheralT = 4
	GDMA_TRIG_PERIPH_SHA    GdmaTriggerPeripheralT = 5
	GDMA_TRIG_PERIPH_ADC    GdmaTriggerPeripheralT = 6
	GDMA_TRIG_PERIPH_DAC    GdmaTriggerPeripheralT = 7
	GDMA_TRIG_PERIPH_LCD    GdmaTriggerPeripheralT = 8
	GDMA_TRIG_PERIPH_CAM    GdmaTriggerPeripheralT = 9
	GDMA_TRIG_PERIPH_RMT    GdmaTriggerPeripheralT = 10
	GDMA_TRIG_PERIPH_PARLIO GdmaTriggerPeripheralT = 11
	GDMA_TRIG_PERIPH_I3C    GdmaTriggerPeripheralT = 12
)

type GdmaChannelDirectionT c.Int

const (
	GDMA_CHANNEL_DIRECTION_TX GdmaChannelDirectionT = 0
	GDMA_CHANNEL_DIRECTION_RX GdmaChannelDirectionT = 1
)

type GdmaEtmEventTypeT c.Int

const (
	GDMA_ETM_EVENT_EOF GdmaEtmEventTypeT = 0
	GDMA_ETM_EVENT_MAX GdmaEtmEventTypeT = 1
)

type GdmaEtmTaskTypeT c.Int

const (
	GDMA_ETM_TASK_START GdmaEtmTaskTypeT = 0
	GDMA_ETM_TASK_MAX   GdmaEtmTaskTypeT = 1
)
