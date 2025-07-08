package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type DmaDescriptorS struct {
	Dw0 struct {
		Size         c.Uint32T
		Length       c.Uint32T
		Reversed2427 c.Uint32T
		ErrEof       c.Uint32T
		Reserved29   c.Uint32T
		SucEof       c.Uint32T
		Owner        c.Uint32T
	}
	Buffer c.Pointer
	Next   *DmaDescriptorT
}
type DmaDescriptorT DmaDescriptorS
type DmaDescriptorAlign4T DmaDescriptorT

type DmaDescriptorAlign8S struct {
	Dw0 struct {
		Size         c.Uint32T
		Length       c.Uint32T
		Reversed2427 c.Uint32T
		ErrEof       c.Uint32T
		Reserved29   c.Uint32T
		SucEof       c.Uint32T
		Owner        c.Uint32T
	}
	Buffer c.Pointer
	Next   *DmaDescriptorAlign8T
}
type DmaDescriptorAlign8T DmaDescriptorAlign8S
