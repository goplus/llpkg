package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

// llgo:type C
type UsbOsglueIntdisenaRoutineT func()

// llgo:type C
type UsbOsglueWaitRoutineT func(c.Int) c.Int

type UsbOsglueDataT struct {
	IntDisProc UsbOsglueIntdisenaRoutineT
	IntEnaProc UsbOsglueIntdisenaRoutineT
	WaitProc   UsbOsglueWaitRoutineT
}
