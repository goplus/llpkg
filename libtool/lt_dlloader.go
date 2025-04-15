package libtool

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

const LT_DLLOADER_H = 1

type Dlloader unsafe.Pointer
type Module unsafe.Pointer
type UserData unsafe.Pointer

type Advise struct {
	Unused [8]uint8
}
type Dladvise *Advise

// llgo:type C
type ModuleOpen func(UserData, *int8, Dladvise) Module

// llgo:type C
type ModuleClose func(UserData, Module) c.Int

// llgo:type C
type FindSym func(UserData, Module, *int8) unsafe.Pointer

// llgo:type C
type DlloaderInit func(UserData) c.Int

// llgo:type C
type DlloaderExit func(UserData) c.Int
type DlloaderPriority c.Int

const (
	LTDLLOADERPREPEND DlloaderPriority = 0
	LTDLLOADERAPPEND  DlloaderPriority = 1
)

/*
This structure defines a module loader, as populated by the get_vtable

	entry point of each loader.
*/
type Dlvtable struct {
	Name         *int8
	SymPrefix    *int8
	ModuleOpen   *unsafe.Pointer
	ModuleClose  *unsafe.Pointer
	FindSym      *unsafe.Pointer
	DlloaderInit *unsafe.Pointer
	DlloaderExit *unsafe.Pointer
	DlloaderData UserData
	Priority     DlloaderPriority
}

// llgo:link (*Dlvtable).DlloaderAdd C.lt_dlloader_add
func (recv_ *Dlvtable) DlloaderAdd() c.Int {
	return 0
}

//go:linkname DlloaderNext C.lt_dlloader_next
func DlloaderNext(loader Dlloader) Dlloader

//go:linkname DlloaderRemove C.lt_dlloader_remove
func DlloaderRemove(name *int8) *Dlvtable

//go:linkname DlloaderFind C.lt_dlloader_find
func DlloaderFind(name *int8) *Dlvtable

//go:linkname DlloaderGet C.lt_dlloader_get
func DlloaderGet(loader Dlloader) *Dlvtable

// llgo:type C
type GetVtable func(UserData) *Dlvtable
