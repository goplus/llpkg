package freertos

import _ "unsafe"

/**
 * @brief Set up the SysTick interrupt
 */
//go:linkname VPortSetupTimer C.vPortSetupTimer
func VPortSetupTimer()
