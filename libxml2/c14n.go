package libxml2

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type C14NMode c.Int

const (
	C14N10          C14NMode = 0
	C14NEXCLUSIVE10 C14NMode = 1
	C14N11          C14NMode = 2
)

//go:linkname C14NDocSaveTo C.xmlC14NDocSaveTo
func C14NDocSaveTo(doc DocPtr, nodes NodeSetPtr, mode c.Int, inclusive_ns_prefixes **Char, with_comments c.Int, buf OutputBufferPtr) c.Int

//go:linkname C14NDocDumpMemory C.xmlC14NDocDumpMemory
func C14NDocDumpMemory(doc DocPtr, nodes NodeSetPtr, mode c.Int, inclusive_ns_prefixes **Char, with_comments c.Int, doc_txt_ptr **Char) c.Int

//go:linkname C14NDocSave C.xmlC14NDocSave
func C14NDocSave(doc DocPtr, nodes NodeSetPtr, mode c.Int, inclusive_ns_prefixes **Char, with_comments c.Int, filename *int8, compression c.Int) c.Int

// llgo:type C
type C14NIsVisibleCallback func(unsafe.Pointer, NodePtr, NodePtr) c.Int

//go:linkname C14NExecute C.xmlC14NExecute
func C14NExecute(doc DocPtr, is_visible_callback C14NIsVisibleCallback, user_data unsafe.Pointer, mode c.Int, inclusive_ns_prefixes **Char, with_comments c.Int, buf OutputBufferPtr) c.Int
