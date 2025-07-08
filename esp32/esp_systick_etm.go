package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Get the ETM event handle of systick hardware's alarm/heartbeat event
 *
 * @note The created ETM event object can be deleted later by calling `esp_etm_del_event`
 *
 * @param[in] core_id CPU core ID
 * @param[out] out_event Returned ETM event handle
 * @return
 *      - ESP_OK Success
 *      - ESP_ERR_INVALID_ARG Parameter error
 */
//go:linkname EspSystickNewEtmAlarmEvent C.esp_systick_new_etm_alarm_event
func EspSystickNewEtmAlarmEvent(core_id c.Int, out_event *EspEtmEventHandleT) EspErrT
