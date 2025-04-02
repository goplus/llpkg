package libtool

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

const LTDL_H = 1

type LtHandle struct {
	Unused [8]uint8
}
type Handle *LtHandle

/* Initialisation and finalisation functions for libltdl. */
//go:linkname Init C.lt_dlinit
func Init() c.Int

//go:linkname Exit C.lt_dlexit
func Exit() c.Int

/* Module search path manipulation.  */
//go:linkname Addsearchdir C.lt_dladdsearchdir
func Addsearchdir(search_dir *int8) c.Int

//go:linkname Insertsearchdir C.lt_dlinsertsearchdir
func Insertsearchdir(before *int8, search_dir *int8) c.Int

//go:linkname Setsearchpath C.lt_dlsetsearchpath
func Setsearchpath(search_path *int8) c.Int

//go:linkname Getsearchpath C.lt_dlgetsearchpath
func Getsearchpath() *int8

//go:linkname Foreachfile C.lt_dlforeachfile
func Foreachfile(search_path *int8, func_ func(*int8, unsafe.Pointer) c.Int, data unsafe.Pointer) c.Int

/* User module loading advisors.  */
//go:linkname AdviseInit C.lt_dladvise_init
func AdviseInit(advise *Advise) c.Int

//go:linkname AdviseDestroy C.lt_dladvise_destroy
func AdviseDestroy(advise *Advise) c.Int

//go:linkname AdviseExt C.lt_dladvise_ext
func AdviseExt(advise *Advise) c.Int

//go:linkname AdviseResident C.lt_dladvise_resident
func AdviseResident(advise *Advise) c.Int

//go:linkname AdviseLocal C.lt_dladvise_local
func AdviseLocal(advise *Advise) c.Int

//go:linkname AdviseGlobal C.lt_dladvise_global
func AdviseGlobal(advise *Advise) c.Int

//go:linkname AdvisePreload C.lt_dladvise_preload
func AdvisePreload(advise *Advise) c.Int

/* Portable libltdl versions of the system dlopen() API. */
//go:linkname Open C.lt_dlopen
func Open(filename *int8) Handle

//go:linkname Openext C.lt_dlopenext
func Openext(filename *int8) Handle

//go:linkname Openadvise C.lt_dlopenadvise
func Openadvise(filename *int8, advise Advise) Handle

//go:linkname Sym C.lt_dlsym
func Sym(handle Handle, name *int8) unsafe.Pointer

//go:linkname Error C.lt_dlerror
func Error() *int8

//go:linkname Close C.lt_dlclose
func Close(handle Handle) c.Int

/*
A preopened symbol. Arrays of this type comprise the exported

	symbols for a dlpreopened module.
*/
type Symlist struct {
	Name    *int8
	Address unsafe.Pointer
}

// llgo:type C
type PreloadCallbackFunc func(Handle) c.Int

// llgo:link (*Symlist).Preload C.lt_dlpreload
func (recv_ *Symlist) Preload() c.Int {
	return 0
}

// llgo:link (*Symlist).PreloadDefault C.lt_dlpreload_default
func (recv_ *Symlist) PreloadDefault() c.Int {
	return 0
}

//go:linkname PreloadOpen C.lt_dlpreload_open
func PreloadOpen(originator *int8, func_ PreloadCallbackFunc) c.Int

type InterfaceId unsafe.Pointer

// llgo:type C
type HandleInterface func(Handle, *int8) c.Int

//go:linkname InterfaceRegister C.lt_dlinterface_register
func InterfaceRegister(id_string *int8, iface HandleInterface) InterfaceId

//go:linkname InterfaceFree C.lt_dlinterface_free
func InterfaceFree(key InterfaceId)

//go:linkname CallerSetData C.lt_dlcaller_set_data
func CallerSetData(key InterfaceId, handle Handle, data unsafe.Pointer) unsafe.Pointer

//go:linkname CallerGetData C.lt_dlcaller_get_data
func CallerGetData(key InterfaceId, handle Handle) unsafe.Pointer

/* Read only information pertaining to a loaded module. */

type Info struct {
	Filename    *int8
	Name        *int8
	RefCount    c.Int
	IsResident  c.Uint
	IsSymglobal c.Uint
	IsSymlocal  c.Uint
}

//go:linkname Getinfo C.lt_dlgetinfo
func Getinfo(handle Handle) *Info

//go:linkname HandleIterate C.lt_dlhandle_iterate
func HandleIterate(iface InterfaceId, place Handle) Handle

//go:linkname HandleFetch C.lt_dlhandle_fetch
func HandleFetch(iface InterfaceId, module_name *int8) Handle

//go:linkname HandleMap C.lt_dlhandle_map
func HandleMap(iface InterfaceId, func_ func(Handle, unsafe.Pointer) c.Int, data unsafe.Pointer) c.Int

/* Deprecated module residency management API. */
//go:linkname Makeresident C.lt_dlmakeresident
func Makeresident(handle Handle) c.Int

//go:linkname Isresident C.lt_dlisresident
func Isresident(handle Handle) c.Int
