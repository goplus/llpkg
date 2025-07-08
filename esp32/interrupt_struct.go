package freertos

import _ "unsafe"

type InterruptDevS struct {
	Unused [8]uint8
}
type InterruptDevT InterruptDevS
