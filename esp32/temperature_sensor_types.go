package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type TemperatureSensorClkSrcT SocPeriphTemperatureSensorClkSrcT
type TemperatureSensorEtmEventTypeT c.Int

const (
	TEMPERATURE_SENSOR_EVENT_OVER_LIMIT TemperatureSensorEtmEventTypeT = 0
	TEMPERATURE_SENSOR_EVENT_MAX        TemperatureSensorEtmEventTypeT = 1
)

type TemperatureSensorEtmTaskTypeT c.Int

const (
	TEMPERATURE_SENSOR_TASK_START TemperatureSensorEtmTaskTypeT = 0
	TEMPERATURE_SENSOR_TASK_STOP  TemperatureSensorEtmTaskTypeT = 1
	TEMPERATURE_SENSOR_TASK_MAX   TemperatureSensorEtmTaskTypeT = 2
)
