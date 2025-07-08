package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief  Start a Trax trace on the current CPU with instructions as unit of delay.
 * Memory blocks to be used as Trax memory must be enabled before
 * calling this function, if needed.
 */
//go:linkname XtTraxStartTraceInstructions C.xt_trax_start_trace_instructions
func XtTraxStartTraceInstructions()

/**
 * @brief  Start a Trax trace on the current CPU with words as unit of delay.
 * Memory blocks to be used as Trax memory must be enabled before
 * calling this function, if needed.
 */
//go:linkname XtTraxStartTraceWords C.xt_trax_start_trace_words
func XtTraxStartTraceWords()

/**
 * @brief Check if Trax trace is active on current CPU.
 *
 * @return bool. Return true if trace is active.
 */
//go:linkname XtTraxTraceIsActive C.xt_trax_trace_is_active
func XtTraxTraceIsActive() bool

/**
 * @brief  Trigger a Trax trace stop after the indicated delay. If this is called
 *         before and the previous delay hasn't ended yet, this will overwrite
 *         that delay with the new value. The delay will always start at the time
 *         the function is called.
 *
 * @param  delay : The delay to stop the trace in, in the unit indicated to
 *              trax_start_trace. Note: the trace memory has 4K words available.
 */
//go:linkname XtTraxTriggerTraceendAfterDelay C.xt_trax_trigger_traceend_after_delay
func XtTraxTriggerTraceendAfterDelay(delay c.Int)
