package zlib

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

const MAX_MEM_LEVEL = 9
const MAX_WBITS = 15

type ZSizeT uintptr
type Byte int8
type UInt c.Uint
type ULong c.Ulong
type Bytef Byte
type Charf int8
type Intf c.Int
type UIntf UInt
type ULongf ULong
type Voidpc unsafe.Pointer
type Voidpf unsafe.Pointer
type Voidp unsafe.Pointer
type ZCrcT c.Uint
