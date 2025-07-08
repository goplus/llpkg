package freertos

import _ "unsafe"

//go:linkname VPortYieldFromISR C.vPortYieldFromISR
func VPortYieldFromISR()

//go:linkname XPortCheckIfInISR C.xPortCheckIfInISR
func XPortCheckIfInISR() BaseTypeT
