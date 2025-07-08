package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type TraxDowncountUnitT c.Int

const (
	TRAX_DOWNCOUNT_WORDS        TraxDowncountUnitT = 0
	TRAX_DOWNCOUNT_INSTRUCTIONS TraxDowncountUnitT = 1
)

type TraxEnaSelectT c.Int

const (
	TRAX_ENA_NONE         TraxEnaSelectT = 0
	TRAX_ENA_PRO          TraxEnaSelectT = 1
	TRAX_ENA_APP          TraxEnaSelectT = 2
	TRAX_ENA_PRO_APP      TraxEnaSelectT = 3
	TRAX_ENA_PRO_APP_SWAP TraxEnaSelectT = 4
)

/**
 * @brief  Enable the trax memory blocks to be used as Trax memory.
 *
 * @param  pro_cpu_enable : true if Trax needs to be enabled for the pro CPU
 * @param  app_cpu_enable : true if Trax needs to be enabled for the pro CPU
 * @param  swap_regions : Normally, the pro CPU writes to Trax mem block 0 while
 *                        the app cpu writes to block 1. Setting this to true
 *                        inverts this.
 *
 * @return esp_err_t. Fails with ESP_ERR_NO_MEM if Trax enable is requested for 2 CPUs
 *                    but memmap only has room for 1, or if Trax memmap is disabled
 *                    entirely.
 */
// llgo:link TraxEnaSelectT.TraxEnable C.trax_enable
func (recv_ TraxEnaSelectT) TraxEnable() c.Int {
	return 0
}

/**
 * @brief  Start a Trax trace on the current CPU
 *
 * @param  units_until_stop : Set the units of the delay that gets passed to
 *              trax_trigger_traceend_after_delay. One of TRAX_DOWNCOUNT_WORDS
 *              or TRAX_DOWNCOUNT_INSTRUCTIONS.
 *
 * @return esp_err_t. Fails with ESP_ERR_NO_MEM if Trax is disabled.
 */
// llgo:link TraxDowncountUnitT.TraxStartTrace C.trax_start_trace
func (recv_ TraxDowncountUnitT) TraxStartTrace() c.Int {
	return 0
}

/**
 * @brief  Trigger a Trax trace stop after the indicated delay. If this is called
 *         before and the previous delay hasn't ended yet, this will overwrite
 *         that delay with the new value. The delay will always start at the time
 *         the function is called.
 *
 * @param  delay : The delay to stop the trace in, in the unit indicated to
 *              trax_start_trace. Note: the trace memory has 4K words available.
 *
 * @return esp_err_t
 */
//go:linkname TraxTriggerTraceendAfterDelay C.trax_trigger_traceend_after_delay
func TraxTriggerTraceendAfterDelay(delay c.Int) c.Int
