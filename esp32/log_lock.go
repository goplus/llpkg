package freertos

import _ "unsafe"

//go:linkname EspLogImplLock C.esp_log_impl_lock
func EspLogImplLock()

//go:linkname EspLogImplLockTimeout C.esp_log_impl_lock_timeout
func EspLogImplLockTimeout() bool

//go:linkname EspLogImplUnlock C.esp_log_impl_unlock
func EspLogImplUnlock()
