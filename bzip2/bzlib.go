package bzip2

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

const BZ_RUN = 0
const BZ_FLUSH = 1
const BZ_FINISH = 2
const BZ_OK = 0
const BZ_RUN_OK = 1
const BZ_FLUSH_OK = 2
const BZ_FINISH_OK = 3
const BZ_STREAM_END = 4
const BZ_MAX_UNUSED = 5000

type BzStream struct {
	NextIn       *int8
	AvailIn      c.Uint
	TotalInLo32  c.Uint
	TotalInHi32  c.Uint
	NextOut      *int8
	AvailOut     c.Uint
	TotalOutLo32 c.Uint
	TotalOutHi32 c.Uint
	State        unsafe.Pointer
	Bzalloc      unsafe.Pointer
	Bzfree       unsafe.Pointer
	Opaque       unsafe.Pointer
}

/*-- Core (low-level) library functions --*/
// llgo:link (*BzStream).BzCompressInit C.BZ2_bzCompressInit
func (recv_ *BzStream) BzCompressInit(blockSize100k c.Int, verbosity c.Int, workFactor c.Int) c.Int {
	return 0
}

// llgo:link (*BzStream).BzCompress C.BZ2_bzCompress
func (recv_ *BzStream) BzCompress(action c.Int) c.Int {
	return 0
}

// llgo:link (*BzStream).BzCompressEnd C.BZ2_bzCompressEnd
func (recv_ *BzStream) BzCompressEnd() c.Int {
	return 0
}

// llgo:link (*BzStream).BzDecompressInit C.BZ2_bzDecompressInit
func (recv_ *BzStream) BzDecompressInit(verbosity c.Int, small c.Int) c.Int {
	return 0
}

// llgo:link (*BzStream).BzDecompress C.BZ2_bzDecompress
func (recv_ *BzStream) BzDecompress() c.Int {
	return 0
}

// llgo:link (*BzStream).BzDecompressEnd C.BZ2_bzDecompressEnd
func (recv_ *BzStream) BzDecompressEnd() c.Int {
	return 0
}

type BZFILE [0]byte

//go:linkname BzReadOpen C.BZ2_bzReadOpen
func BzReadOpen(bzerror *c.Int, f *c.FILE, verbosity c.Int, small c.Int, unused unsafe.Pointer, nUnused c.Int) *BZFILE

//go:linkname BzReadClose C.BZ2_bzReadClose
func BzReadClose(bzerror *c.Int, b *BZFILE)

//go:linkname BzReadGetUnused C.BZ2_bzReadGetUnused
func BzReadGetUnused(bzerror *c.Int, b *BZFILE, unused *unsafe.Pointer, nUnused *c.Int)

//go:linkname BzRead C.BZ2_bzRead
func BzRead(bzerror *c.Int, b *BZFILE, buf unsafe.Pointer, len c.Int) c.Int

//go:linkname BzWriteOpen C.BZ2_bzWriteOpen
func BzWriteOpen(bzerror *c.Int, f *c.FILE, blockSize100k c.Int, verbosity c.Int, workFactor c.Int) *BZFILE

//go:linkname BzWrite C.BZ2_bzWrite
func BzWrite(bzerror *c.Int, b *BZFILE, buf unsafe.Pointer, len c.Int)

//go:linkname BzWriteClose C.BZ2_bzWriteClose
func BzWriteClose(bzerror *c.Int, b *BZFILE, abandon c.Int, nbytes_in *c.Uint, nbytes_out *c.Uint)

//go:linkname BzWriteClose64 C.BZ2_bzWriteClose64
func BzWriteClose64(bzerror *c.Int, b *BZFILE, abandon c.Int, nbytes_in_lo32 *c.Uint, nbytes_in_hi32 *c.Uint, nbytes_out_lo32 *c.Uint, nbytes_out_hi32 *c.Uint)

/*-- Utility functions --*/
//go:linkname BzBuffToBuffCompress C.BZ2_bzBuffToBuffCompress
func BzBuffToBuffCompress(dest *int8, destLen *c.Uint, source *int8, sourceLen c.Uint, blockSize100k c.Int, verbosity c.Int, workFactor c.Int) c.Int

//go:linkname BzBuffToBuffDecompress C.BZ2_bzBuffToBuffDecompress
func BzBuffToBuffDecompress(dest *int8, destLen *c.Uint, source *int8, sourceLen c.Uint, small c.Int, verbosity c.Int) c.Int

/*--
   Code contributed by Yoshioka Tsuneo (tsuneo@rr.iij4u.or.jp)
   to support better zlib compatibility.
   This code is not _officially_ part of libbzip2 (yet);
   I haven't tested it, documented it, or considered the
   threading-safeness of it.
   If this code breaks, please contact both Yoshioka and me.
--*/
//go:linkname BzlibVersion C.BZ2_bzlibVersion
func BzlibVersion() *int8

//go:linkname Bzopen C.BZ2_bzopen
func Bzopen(path *int8, mode *int8) *BZFILE

//go:linkname Bzdopen C.BZ2_bzdopen
func Bzdopen(fd c.Int, mode *int8) *BZFILE

// llgo:link (*BZFILE).Bzread C.BZ2_bzread
func (recv_ *BZFILE) Bzread(buf unsafe.Pointer, len c.Int) c.Int {
	return 0
}

// llgo:link (*BZFILE).Bzwrite C.BZ2_bzwrite
func (recv_ *BZFILE) Bzwrite(buf unsafe.Pointer, len c.Int) c.Int {
	return 0
}

// llgo:link (*BZFILE).Bzflush C.BZ2_bzflush
func (recv_ *BZFILE) Bzflush() c.Int {
	return 0
}

// llgo:link (*BZFILE).Bzclose C.BZ2_bzclose
func (recv_ *BZFILE) Bzclose() {
}

// llgo:link (*BZFILE).Bzerror C.BZ2_bzerror
func (recv_ *BZFILE) Bzerror(errnum *c.Int) *int8 {
	return nil
}
