package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type McpwmSocHandleT *McpwmDevT

/**
 * @brief HAL layer configuration
 */

type McpwmHalInitConfigT struct {
	GroupId c.Int
}

/**
 * Context that should be maintained by both the driver and the HAL
 */

type McpwmHalContextT struct {
	Dev McpwmSocHandleT
}

/**
 * @brief Initialize the internal state of the HAL.
 *
 * @param hal Context of the HAL layer.
 * @param init_config Configuration for the HAL to be used only once.
 */
// llgo:link (*McpwmHalContextT).McpwmHalInit C.mcpwm_hal_init
func (recv_ *McpwmHalContextT) McpwmHalInit(init_config *McpwmHalInitConfigT) {
}

/**
 * @brief Deinitialize the HAL driver.
 *
 * @param hal Context of the HAL layer.
 */
// llgo:link (*McpwmHalContextT).McpwmHalDeinit C.mcpwm_hal_deinit
func (recv_ *McpwmHalContextT) McpwmHalDeinit() {
}

/**
 * @brief Reset MCPWM timer
 *
 * @param hal Context of the HAL layer.
 * @param timer_id Timer ID
 */
// llgo:link (*McpwmHalContextT).McpwmHalTimerReset C.mcpwm_hal_timer_reset
func (recv_ *McpwmHalContextT) McpwmHalTimerReset(timer_id c.Int) {
}

/**
 * @brief Reset MCPWM operator
 *
 * @param hal Context of the HAL layer.
 * @param oper_id Operator ID
 */
// llgo:link (*McpwmHalContextT).McpwmHalOperatorReset C.mcpwm_hal_operator_reset
func (recv_ *McpwmHalContextT) McpwmHalOperatorReset(oper_id c.Int) {
}

/**
 * @brief Reset MCPWM generator
 *
 * @param hal Context of the HAL layer.
 * @param oper_id Operator ID
 * @param gen_id Generator ID
 */
// llgo:link (*McpwmHalContextT).McpwmHalGeneratorReset C.mcpwm_hal_generator_reset
func (recv_ *McpwmHalContextT) McpwmHalGeneratorReset(oper_id c.Int, gen_id c.Int) {
}
