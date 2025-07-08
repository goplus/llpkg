package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Whether to allow the TOP power domain to be powered off.
 *
 * In light sleep mode, only when the system can provide enough memory
 * for digital peripheral clock retention, the TOP power domain can be
 * powered off.
 *
 * @return True to allow power off
 */
//go:linkname ClockDomainPdAllowed C.clock_domain_pd_allowed
func ClockDomainPdAllowed() bool

/**
 * @brief SoC system clock retention initialize.
 *
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_NO_MEM not enough memory for system clock retention
 *      - ESP_ERR_INVALID_ARG if either of the arguments is out of range
 */
//go:linkname SleepClockSystemRetentionInit C.sleep_clock_system_retention_init
func SleepClockSystemRetentionInit(arg c.Pointer) EspErrT
