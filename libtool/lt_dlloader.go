package libtool

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

const LT_DLLOADER_H = 1

type Loader unsafe.Pointer
type LtModule unsafe.Pointer
type LtUserData unsafe.Pointer

type LtAdvise struct {
	Unused [8]uint8
}
type Advise *LtAdvise

// llgo:type C
type LtModuleOpen func(LtUserData, *int8, Advise) LtModule

// llgo:type C
type LtModuleClose func(LtUserData, LtModule) c.Int

// llgo:type C
type LtFindSym func(LtUserData, LtModule, *int8) unsafe.Pointer

// llgo:type C
type LoaderInit func(LtUserData) c.Int

// llgo:type C
type LoaderExit func(LtUserData) c.Int
type LoaderPriority c.Int

const (
	LTDLLOADERPREPEND LoaderPriority = 0
	LTDLLOADERAPPEND  LoaderPriority = 1
)

/*
This structure defines a module loader, as populated by the get_vtable

	entry point of each loader.
*/
type Vtable struct {
	Name         *int8
	SymPrefix    *int8
	ModuleOpen   *unsafe.Pointer
	ModuleClose  *unsafe.Pointer
	FindSym      *unsafe.Pointer
	DlloaderInit *unsafe.Pointer
	DlloaderExit *unsafe.Pointer
	DlloaderData LtUserData
	Priority     LoaderPriority
}

// llgo:link (*Vtable).LoaderAdd C.lt_dlloader_add
func (recv_ *Vtable) LoaderAdd() c.Int {
	return 0
}

//go:linkname LoaderNext C.lt_dlloader_next
func LoaderNext(loader Loader) Loader

//go:linkname LoaderRemove C.lt_dlloader_remove
func LoaderRemove(name *int8) *Vtable

//go:linkname LoaderFind C.lt_dlloader_find
func LoaderFind(name *int8) *Vtable

//go:linkname LoaderGet C.lt_dlloader_get
func LoaderGet(loader Loader) *Vtable

// llgo:type C
type LtGetVtable func(LtUserData) *Vtable
