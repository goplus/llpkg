package freertos

import _ "unsafe"

/**
 * @brief Whether to allow the TOP power domain to be powered off.
 *
 * In light sleep mode, only when the system can provide enough memory
 * for digital peripheral retention, the TOP power domain can be powered off.
 *
 * @return True to allow power off
 */
//go:linkname PeripheralDomainPdAllowed C.peripheral_domain_pd_allowed
func PeripheralDomainPdAllowed() bool
