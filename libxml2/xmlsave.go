package libxml2

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type SaveOption c.Int

const (
	SAVEFORMAT   SaveOption = 1
	SAVENODECL   SaveOption = 2
	SAVENOEMPTY  SaveOption = 4
	SAVENOXHTML  SaveOption = 8
	SAVEXHTML    SaveOption = 16
	SAVEASXML    SaveOption = 32
	SAVEASHTML   SaveOption = 64
	SAVEWSNONSIG SaveOption = 128
)

type X_xmlSaveCtxt struct {
	Unused [8]uint8
}
type SaveCtxt X_xmlSaveCtxt
type SaveCtxtPtr *SaveCtxt

//go:linkname SaveToFd C.xmlSaveToFd
func SaveToFd(fd c.Int, encoding *int8, options c.Int) SaveCtxtPtr

//go:linkname SaveToFilename C.xmlSaveToFilename
func SaveToFilename(filename *int8, encoding *int8, options c.Int) SaveCtxtPtr

//go:linkname SaveToBuffer C.xmlSaveToBuffer
func SaveToBuffer(buffer BufferPtr, encoding *int8, options c.Int) SaveCtxtPtr

//go:linkname SaveToIO C.xmlSaveToIO
func SaveToIO(iowrite OutputWriteCallback, ioclose OutputCloseCallback, ioctx unsafe.Pointer, encoding *int8, options c.Int) SaveCtxtPtr

//go:linkname SaveDoc C.xmlSaveDoc
func SaveDoc(ctxt SaveCtxtPtr, doc DocPtr) c.Long

//go:linkname SaveTree C.xmlSaveTree
func SaveTree(ctxt SaveCtxtPtr, node NodePtr) c.Long

//go:linkname SaveFlush C.xmlSaveFlush
func SaveFlush(ctxt SaveCtxtPtr) c.Int

//go:linkname SaveClose C.xmlSaveClose
func SaveClose(ctxt SaveCtxtPtr) c.Int

//go:linkname SaveFinish C.xmlSaveFinish
func SaveFinish(ctxt SaveCtxtPtr) c.Int

//go:linkname SaveSetEscape C.xmlSaveSetEscape
func SaveSetEscape(ctxt SaveCtxtPtr, escape CharEncodingOutputFunc) c.Int

//go:linkname SaveSetAttrEscape C.xmlSaveSetAttrEscape
func SaveSetAttrEscape(ctxt SaveCtxtPtr, escape CharEncodingOutputFunc) c.Int

//go:linkname ThrDefIndentTreeOutput C.xmlThrDefIndentTreeOutput
func ThrDefIndentTreeOutput(v c.Int) c.Int

//go:linkname ThrDefTreeIndentString C.xmlThrDefTreeIndentString
func ThrDefTreeIndentString(v *int8) *int8

//go:linkname ThrDefSaveNoEmptyTags C.xmlThrDefSaveNoEmptyTags
func ThrDefSaveNoEmptyTags(v c.Int) c.Int
