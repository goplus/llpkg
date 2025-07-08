package freertos

import _ "unsafe"

// llgo:type C
type EspFreertosIdleCbT func() bool

// llgo:type C
type EspFreertosTickCbT func()

/**
 * @brief  Register a callback to be called from the specified core's idle hook.
 *         The callback should return true if it should be called by the idle hook
 *         once per interrupt (or FreeRTOS tick), and return false if it should
 *         be called repeatedly as fast as possible by the idle hook.
 *
 * @warning Idle callbacks MUST NOT, UNDER ANY CIRCUMSTANCES, CALL
 *          A FUNCTION THAT MIGHT BLOCK.
 *
 * @param[in]  new_idle_cb     Callback to be called
 * @param[in]  cpuid           id of the core
 *
 * @return
 *     - ESP_OK:              Callback registered to the specified core's idle hook
 *     - ESP_ERR_NO_MEM:      No more space on the specified core's idle hook to register callback
 *     - ESP_ERR_INVALID_ARG: cpuid is invalid
 */
//go:linkname EspRegisterFreertosIdleHookForCpu C.esp_register_freertos_idle_hook_for_cpu
func EspRegisterFreertosIdleHookForCpu(new_idle_cb EspFreertosIdleCbT, cpuid UBaseTypeT) EspErrT

/**
 * @brief  Register a callback to the idle hook of the core that calls this function.
 *         The callback should return true if it should be called by the idle hook
 *         once per interrupt (or FreeRTOS tick), and return false if it should
 *         be called repeatedly as fast as possible by the idle hook.
 *
 * @warning Idle callbacks MUST NOT, UNDER ANY CIRCUMSTANCES, CALL
 *          A FUNCTION THAT MIGHT BLOCK.
 *
 * @param[in]  new_idle_cb     Callback to be called
 *
 * @return
 *     - ESP_OK:         Callback registered to the calling core's idle hook
 *     - ESP_ERR_NO_MEM: No more space on the calling core's idle hook to register callback
 */
//go:linkname EspRegisterFreertosIdleHook C.esp_register_freertos_idle_hook
func EspRegisterFreertosIdleHook(new_idle_cb EspFreertosIdleCbT) EspErrT

/**
 * @brief  Register a callback to be called from the specified core's tick hook.
 *
 * @param[in]  new_tick_cb     Callback to be called
 * @param[in]  cpuid           id of the core
 *
 * @return
 *     - ESP_OK:              Callback registered to specified core's tick hook
 *     - ESP_ERR_NO_MEM:      No more space on the specified core's tick hook to register the callback
 *     - ESP_ERR_INVALID_ARG: cpuid is invalid
 */
//go:linkname EspRegisterFreertosTickHookForCpu C.esp_register_freertos_tick_hook_for_cpu
func EspRegisterFreertosTickHookForCpu(new_tick_cb EspFreertosTickCbT, cpuid UBaseTypeT) EspErrT

/**
 * @brief  Register a callback to be called from the calling core's tick hook.
 *
 * @param[in]  new_tick_cb     Callback to be called
 *
 * @return
 *     - ESP_OK:         Callback registered to the calling core's tick hook
 *     - ESP_ERR_NO_MEM: No more space on the calling core's tick hook to register the callback
 */
//go:linkname EspRegisterFreertosTickHook C.esp_register_freertos_tick_hook
func EspRegisterFreertosTickHook(new_tick_cb EspFreertosTickCbT) EspErrT

/**
 * @brief  Unregister an idle callback from the idle hook of the specified core
 *
 * @param[in]  old_idle_cb     Callback to be unregistered
 * @param[in]  cpuid           id of the core
 */
//go:linkname EspDeregisterFreertosIdleHookForCpu C.esp_deregister_freertos_idle_hook_for_cpu
func EspDeregisterFreertosIdleHookForCpu(old_idle_cb EspFreertosIdleCbT, cpuid UBaseTypeT)

/**
 * @brief  Unregister an idle callback. If the idle callback is registered to
 *         the idle hooks of both cores, the idle hook will be unregistered from
 *         both cores
 *
 * @param[in]  old_idle_cb     Callback to be unregistered
 */
//go:linkname EspDeregisterFreertosIdleHook C.esp_deregister_freertos_idle_hook
func EspDeregisterFreertosIdleHook(old_idle_cb EspFreertosIdleCbT)

/**
 * @brief  Unregister a tick callback from the tick hook of the specified core
 *
 * @param[in]  old_tick_cb     Callback to be unregistered
 * @param[in]  cpuid           id of the core
 */
//go:linkname EspDeregisterFreertosTickHookForCpu C.esp_deregister_freertos_tick_hook_for_cpu
func EspDeregisterFreertosTickHookForCpu(old_tick_cb EspFreertosTickCbT, cpuid UBaseTypeT)

/**
 * @brief  Unregister a tick callback. If the tick callback is registered to the
 *         tick hooks of both cores, the tick hook will be unregistered from
 *         both cores
 *
 * @param[in]  old_tick_cb     Callback to be unregistered
 */
//go:linkname EspDeregisterFreertosTickHook C.esp_deregister_freertos_tick_hook
func EspDeregisterFreertosTickHook(old_tick_cb EspFreertosTickCbT)
