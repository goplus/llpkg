package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/*
 * @brief The structure of the counter value in systimer
 *
 */

type SystimerCounterValueT struct {
	Unused [8]uint8
}
type SystimerAlarmModeT c.Int

const (
	SYSTIMER_ALARM_MODE_ONESHOT SystimerAlarmModeT = 0
	SYSTIMER_ALARM_MODE_PERIOD  SystimerAlarmModeT = 1
)
