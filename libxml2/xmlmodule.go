package libxml2

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type X_xmlModule struct {
	Unused [8]uint8
}
type Module X_xmlModule
type ModulePtr *Module
type ModuleOption c.Int

const (
	MODULELAZY  ModuleOption = 1
	MODULELOCAL ModuleOption = 2
)

//go:linkname ModuleOpen C.xmlModuleOpen
func ModuleOpen(filename *int8, options c.Int) ModulePtr

//go:linkname ModuleSymbol C.xmlModuleSymbol
func ModuleSymbol(module ModulePtr, name *int8, result *unsafe.Pointer) c.Int

//go:linkname ModuleClose C.xmlModuleClose
func ModuleClose(module ModulePtr) c.Int

//go:linkname ModuleFree C.xmlModuleFree
func ModuleFree(module ModulePtr) c.Int
