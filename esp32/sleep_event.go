package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type EspSleepEventCbIndexT c.Int

const (
	SLEEP_EVENT_HW_EXIT_SLEEP       EspSleepEventCbIndexT = 0
	SLEEP_EVENT_SW_CLK_READY        EspSleepEventCbIndexT = 1
	SLEEP_EVENT_SW_EXIT_SLEEP       EspSleepEventCbIndexT = 2
	SLEEP_EVENT_SW_GOTO_SLEEP       EspSleepEventCbIndexT = 3
	SLEEP_EVENT_HW_TIME_START       EspSleepEventCbIndexT = 4
	SLEEP_EVENT_HW_GOTO_SLEEP       EspSleepEventCbIndexT = 5
	SLEEP_EVENT_SW_CPU_TO_MEM_START EspSleepEventCbIndexT = 6
	SLEEP_EVENT_SW_CPU_TO_MEM_END   EspSleepEventCbIndexT = 7
	SLEEP_EVENT_HW_PLL_EN_START     EspSleepEventCbIndexT = 8
	SLEEP_EVENT_HW_PLL_EN_STOP      EspSleepEventCbIndexT = 9
	SLEEP_EVENT_CB_INDEX_NUM        EspSleepEventCbIndexT = 10
)

// llgo:type C
type EspSleepEventCbT func(c.Pointer, c.Pointer) EspErrT

/**
 * @brief Function entry parameter types for light sleep event callback functions (if CONFIG_FREERTOS_USE_TICKLESS_IDLE)
 */

type X_espSleepEventCbConfigT struct {
	Cb      EspSleepEventCbT
	UserArg c.Pointer
	Prior   c.Uint32T
	Next    *X_espSleepEventCbConfigT
}
type EspSleepEventCbConfigT X_espSleepEventCbConfigT

type X_espSleepEventCbsConfigT struct {
	SleepEventCbConfig [10]*EspSleepEventCbConfigT
}
type EspSleepEventCbsConfigT X_espSleepEventCbsConfigT

/**
 * @brief Register event callbacks for light sleep internal events (if CONFIG_FREERTOS_USE_TICKLESS_IDLE)
 * @param event_id      Designed to register the corresponding event_cb in g_sleep_event_cbs_config
 * @param event_cb_conf Config struct containing event callback function and corresponding argument
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_ARG if the input parameter event_cb_conf is NULL or event_id is out of range
 *      - ESP_ERR_NO_MEM if the remaining memory is insufficient to support malloc
 *      - ESP_FAIL if register the same function repeatedly
 *
 * @note Some of these callback functions are called from IDLE task context hence they cannot call any blocking functions
 * @note Passing NULL value will not deregister the callbacks, it will silently ignore and return ESP_OK
 */
// llgo:link EspSleepEventCbIndexT.EspSleepRegisterEventCallback C.esp_sleep_register_event_callback
func (recv_ EspSleepEventCbIndexT) EspSleepRegisterEventCallback(event_cb_conf *EspSleepEventCbConfigT) EspErrT {
	return 0
}

/**
 * @brief Unregister event callbacks for light sleep internal events (if CONFIG_FREERTOS_USE_TICKLESS_IDLE)
 * @param event_id      Designed to unregister the corresponding event_cb in g_sleep_event_cbs_config
 * @param event_cb_conf Config struct containing event callback function and corresponding argument
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_ARG if the input parameter cb is NULL or event_id is out of range
 */
// llgo:link EspSleepEventCbIndexT.EspSleepUnregisterEventCallback C.esp_sleep_unregister_event_callback
func (recv_ EspSleepEventCbIndexT) EspSleepUnregisterEventCallback(cb EspSleepEventCbT) EspErrT {
	return 0
}

/**
 * @brief Designed to execute functions in the esp_sleep_event_cb_config_t linked list
 *
 * @param event_id   Designed to annotate the corresponding event_cb in g_sleep_event_cbs_config
 * @param ext_arg    Designed to pass external parameters
 * @return None
 */
// llgo:link EspSleepEventCbIndexT.EspSleepExecuteEventCallbacks C.esp_sleep_execute_event_callbacks
func (recv_ EspSleepEventCbIndexT) EspSleepExecuteEventCallbacks(ext_arg c.Pointer) {
}
