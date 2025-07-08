package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Initialize the IPC ISR interrupt for a specific CPU.
 *
 * This function initializes the IPC ISR (Inter-Processor Communication Interrupt)
 * for a specific CPU core. It configures the interrupt source and enables the
 * IPC ISR interrupt for the specified CPU.
 *
 * @param[in] cpuid The ID of the CPU core to initialize IPC ISR for.
 */
//go:linkname EspIpcIsrPortInit C.esp_ipc_isr_port_init
func EspIpcIsrPortInit(cpuid c.Int)

/**
 * @brief Trigger an interrupt on a specific CPU core.
 *
 * @param[in] cpuid The ID of the CPU core to trigger the interrupt on (0 or 1).
 */
//go:linkname EspIpcIsrPortIntTrigger C.esp_ipc_isr_port_int_trigger
func EspIpcIsrPortIntTrigger(cpuid c.Int)
