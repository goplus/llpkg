package freertos

import _ "unsafe"

//go:linkname EspBrownoutInit C.esp_brownout_init
func EspBrownoutInit()

//go:linkname EspBrownoutDisable C.esp_brownout_disable
func EspBrownoutDisable()
