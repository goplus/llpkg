package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const MINIZ_X86_OR_X64_CPU = 0
const MINIZ_LITTLE_ENDIAN = 1
const MINIZ_USE_UNALIGNED_LOADS_AND_STORES = 0
const MINIZ_HAS_64BIT_REGISTERS = 0
const TINFL_USE_64BIT_BITBUF = 0
const MZ_DEFLATED = 8
const TINFL_LZ_DICT_SIZE = 32768
const TDEFL_LESS_MEMORY = 1

type MzUlong c.Ulong

// mz_free() internally uses the MZ_FREE() macro (which by default calls free() unless you've modified the MZ_MALLOC macro) to release a block allocated from the heap.
//
//go:linkname MzFree C.mz_free
func MzFree(p c.Pointer)

// mz_adler32() returns the initial adler-32 value to use when called with ptr==NULL.
// llgo:link MzUlong.MzAdler32 C.mz_adler32
func (recv_ MzUlong) MzAdler32(ptr *c.Char, buf_len c.SizeT) MzUlong {
	return 0
}

// mz_crc32() returns the initial CRC-32 value to use when called with ptr==NULL.
// llgo:link MzUlong.MzCrc32 C.mz_crc32
func (recv_ MzUlong) MzCrc32(ptr *c.Char, buf_len c.SizeT) MzUlong {
	return 0
}

const (
	MZ_DEFAULT_STRATEGY c.Int = 0
	MZ_FILTERED         c.Int = 1
	MZ_HUFFMAN_ONLY     c.Int = 2
	MZ_RLE              c.Int = 3
	MZ_FIXED            c.Int = 4
)

type MzUint8 c.Char
type MzInt16 int16
type MzUint16 uint16
type MzUint32 c.Uint
type MzUint c.Uint
type MzInt64 c.LongLong
type MzUint64 c.UlongLong
type MzBool c.Int

const (
	TINFL_FLAG_PARSE_ZLIB_HEADER             c.Int = 1
	TINFL_FLAG_HAS_MORE_INPUT                c.Int = 2
	TINFL_FLAG_USING_NON_WRAPPING_OUTPUT_BUF c.Int = 4
	TINFL_FLAG_COMPUTE_ADLER32               c.Int = 8
)

// High level decompression functions:
// tinfl_decompress_mem_to_heap() decompresses a block in memory to a heap block allocated via malloc().
// On entry:
//
//	pSrc_buf, src_buf_len: Pointer and size of the Deflate or zlib source data to decompress.
//
// On return:
//
//	Function returns a pointer to the decompressed data, or NULL on failure.
//	*pOut_len will be set to the decompressed data's size, which could be larger than src_buf_len on uncompressible data.
//	The caller must call mz_free() on the returned block when it's no longer needed.
//
//go:linkname TinflDecompressMemToHeap C.tinfl_decompress_mem_to_heap
func TinflDecompressMemToHeap(pSrc_buf c.Pointer, src_buf_len c.SizeT, pOut_len *c.SizeT, flags c.Int) c.Pointer

//go:linkname TinflDecompressMemToMem C.tinfl_decompress_mem_to_mem
func TinflDecompressMemToMem(pOut_buf c.Pointer, out_buf_len c.SizeT, pSrc_buf c.Pointer, src_buf_len c.SizeT, flags c.Int) c.SizeT

// llgo:type C
type TinflPutBufFuncPtr func(c.Pointer, c.Int, c.Pointer) c.Int

//go:linkname TinflDecompressMemToCallback C.tinfl_decompress_mem_to_callback
func TinflDecompressMemToCallback(pIn_buf c.Pointer, pIn_buf_size *c.SizeT, pPut_buf_func TinflPutBufFuncPtr, pPut_buf_user c.Pointer, flags c.Int) c.Int

type TinflDecompressorTag struct {
	MState               MzUint32
	MNumBits             MzUint32
	MZhdr0               MzUint32
	MZhdr1               MzUint32
	MZAdler32            MzUint32
	MFinal               MzUint32
	MType                MzUint32
	MCheckAdler32        MzUint32
	MDist                MzUint32
	MCounter             MzUint32
	MNumExtra            MzUint32
	MTableSizes          [3]MzUint32
	MBitBuf              TinflBitBufT
	MDistFromOutBufStart c.SizeT
	MTables              [3]TinflHuffTable
	MRawHeader           [4]MzUint8
	MLenCodes            [457]MzUint8
}
type TinflDecompressor TinflDecompressorTag
type TinflStatus c.Int

const (
	TINFL_STATUS_BAD_PARAM        TinflStatus = -3
	TINFL_STATUS_ADLER32_MISMATCH TinflStatus = -2
	TINFL_STATUS_FAILED           TinflStatus = -1
	TINFL_STATUS_DONE             TinflStatus = 0
	TINFL_STATUS_NEEDS_MORE_INPUT TinflStatus = 1
	TINFL_STATUS_HAS_MORE_OUTPUT  TinflStatus = 2
)

// Main low-level decompressor coroutine function. This is the only function actually needed for decompression. All the other functions are just high-level helpers for improved usability.
// This is a universal API, i.e. it can be used as a building block to build any desired higher level decompression API. In the limit case, it can be called once per every byte input or output.
// llgo:link (*TinflDecompressor).TinflDecompress C.tinfl_decompress
func (recv_ *TinflDecompressor) TinflDecompress(pIn_buf_next *MzUint8, pIn_buf_size *c.SizeT, pOut_buf_start *MzUint8, pOut_buf_next *MzUint8, pOut_buf_size *c.SizeT, decomp_flags MzUint32) TinflStatus {
	return 0
}

const (
	TINFL_MAX_HUFF_TABLES    c.Int = 3
	TINFL_MAX_HUFF_SYMBOLS_0 c.Int = 288
	TINFL_MAX_HUFF_SYMBOLS_1 c.Int = 32
	TINFL_MAX_HUFF_SYMBOLS_2 c.Int = 19
	TINFL_FAST_LOOKUP_BITS   c.Int = 10
	TINFL_FAST_LOOKUP_SIZE   c.Int = 1024
)

type TinflHuffTable struct {
	MCodeSize [288]MzUint8
	MLookUp   [1024]MzInt16
	MTree     [576]MzInt16
}
type TinflBitBufT MzUint64

const (
	TDEFL_HUFFMAN_ONLY       c.Int = 0
	TDEFL_DEFAULT_MAX_PROBES c.Int = 128
	TDEFL_MAX_PROBES_MASK    c.Int = 4095
)
const (
	TDEFL_WRITE_ZLIB_HEADER             c.Int = 4096
	TDEFL_COMPUTE_ADLER32               c.Int = 8192
	TDEFL_GREEDY_PARSING_FLAG           c.Int = 16384
	TDEFL_NONDETERMINISTIC_PARSING_FLAG c.Int = 32768
	TDEFL_RLE_MATCHES                   c.Int = 65536
	TDEFL_FILTER_MATCHES                c.Int = 131072
	TDEFL_FORCE_ALL_STATIC_BLOCKS       c.Int = 262144
	TDEFL_FORCE_ALL_RAW_BLOCKS          c.Int = 524288
)

// High level compression functions:
// tdefl_compress_mem_to_heap() compresses a block in memory to a heap block allocated via malloc().
// On entry:
//
//	pSrc_buf, src_buf_len: Pointer and size of source block to compress.
//	flags: The max match finder probes (default is 128) logically OR'd against the above flags. Higher probes are slower but improve compression.
//
// On return:
//
//	Function returns a pointer to the compressed data, or NULL on failure.
//	*pOut_len will be set to the compressed data's size, which could be larger than src_buf_len on uncompressible data.
//	The caller must free() the returned block when it's no longer needed.
//
//go:linkname TdeflCompressMemToHeap C.tdefl_compress_mem_to_heap
func TdeflCompressMemToHeap(pSrc_buf c.Pointer, src_buf_len c.SizeT, pOut_len *c.SizeT, flags c.Int) c.Pointer

// tdefl_compress_mem_to_mem() compresses a block in memory to another block in memory.
// Returns 0 on failure.
//
//go:linkname TdeflCompressMemToMem C.tdefl_compress_mem_to_mem
func TdeflCompressMemToMem(pOut_buf c.Pointer, out_buf_len c.SizeT, pSrc_buf c.Pointer, src_buf_len c.SizeT, flags c.Int) c.SizeT

// Compresses an image to a compressed PNG file in memory.
// On entry:
//
//	pImage, w, h, and num_chans describe the image to compress. num_chans may be 1, 2, 3, or 4.
//	The image pitch in bytes per scanline will be w*num_chans. The leftmost pixel on the top scanline is stored first in memory.
//	level may range from [0,10], use MZ_NO_COMPRESSION, MZ_BEST_SPEED, MZ_BEST_COMPRESSION, etc. or a decent default is MZ_DEFAULT_LEVEL
//	If flip is true, the image will be flipped on the Y axis (useful for OpenGL apps).
//
// On return:
//
//	Function returns a pointer to the compressed data, or NULL on failure.
//	*pLen_out will be set to the size of the PNG image file.
//	The caller must mz_free() the returned heap block (which will typically be larger than *pLen_out) when it's no longer needed.
//
//go:linkname TdeflWriteImageToPngFileInMemoryEx C.tdefl_write_image_to_png_file_in_memory_ex
func TdeflWriteImageToPngFileInMemoryEx(pImage c.Pointer, w c.Int, h c.Int, num_chans c.Int, pLen_out *c.SizeT, level MzUint, flip MzBool) c.Pointer

//go:linkname TdeflWriteImageToPngFileInMemory C.tdefl_write_image_to_png_file_in_memory
func TdeflWriteImageToPngFileInMemory(pImage c.Pointer, w c.Int, h c.Int, num_chans c.Int, pLen_out *c.SizeT) c.Pointer

// llgo:type C
type TdeflPutBufFuncPtr func(c.Pointer, c.Int, c.Pointer) MzBool

// tdefl_compress_mem_to_output() compresses a block to an output stream. The above helpers use this function internally.
//
//go:linkname TdeflCompressMemToOutput C.tdefl_compress_mem_to_output
func TdeflCompressMemToOutput(pBuf c.Pointer, buf_len c.SizeT, pPut_buf_func TdeflPutBufFuncPtr, pPut_buf_user c.Pointer, flags c.Int) MzBool

const (
	TDEFL_MAX_HUFF_TABLES    c.Int = 3
	TDEFL_MAX_HUFF_SYMBOLS_0 c.Int = 288
	TDEFL_MAX_HUFF_SYMBOLS_1 c.Int = 32
	TDEFL_MAX_HUFF_SYMBOLS_2 c.Int = 19
	TDEFL_LZ_DICT_SIZE       c.Int = 32768
	TDEFL_LZ_DICT_SIZE_MASK  c.Int = 32767
	TDEFL_MIN_MATCH_LEN      c.Int = 3
	TDEFL_MAX_MATCH_LEN      c.Int = 258
)
const (
	TDEFL_LZ_CODE_BUF_SIZE      c.Int = 24576
	TDEFL_OUT_BUF_SIZE          c.Int = 31948
	TDEFL_MAX_HUFF_SYMBOLS      c.Int = 288
	TDEFL_LZ_HASH_BITS          c.Int = 12
	TDEFL_LEVEL1_HASH_SIZE_MASK c.Int = 4095
	TDEFL_LZ_HASH_SHIFT         c.Int = 4
	TDEFL_LZ_HASH_SIZE          c.Int = 4096
)

type TdeflStatus c.Int

const (
	TDEFL_STATUS_BAD_PARAM      TdeflStatus = -2
	TDEFL_STATUS_PUT_BUF_FAILED TdeflStatus = -1
	TDEFL_STATUS_OKAY           TdeflStatus = 0
	TDEFL_STATUS_DONE           TdeflStatus = 1
)

type TdeflFlush c.Int

const (
	TDEFL_NO_FLUSH   TdeflFlush = 0
	TDEFL_SYNC_FLUSH TdeflFlush = 2
	TDEFL_FULL_FLUSH TdeflFlush = 3
	TDEFL_FINISH     TdeflFlush = 4
)

// tdefl's compression state structure.
type TdeflCompressor struct {
	MPPutBufFunc          TdeflPutBufFuncPtr
	MPPutBufUser          c.Pointer
	MFlags                MzUint
	MMaxProbes            [2]MzUint
	MGreedyParsing        c.Int
	MAdler32              MzUint
	MLookaheadPos         MzUint
	MLookaheadSize        MzUint
	MDictSize             MzUint
	MPLZCodeBuf           *MzUint8
	MPLZFlags             *MzUint8
	MPOutputBuf           *MzUint8
	MPOutputBufEnd        *MzUint8
	MNumFlagsLeft         MzUint
	MTotalLzBytes         MzUint
	MLzCodeBufDictPos     MzUint
	MBitsIn               MzUint
	MBitBuffer            MzUint
	MSavedMatchDist       MzUint
	MSavedMatchLen        MzUint
	MSavedLit             MzUint
	MOutputFlushOfs       MzUint
	MOutputFlushRemaining MzUint
	MFinished             MzUint
	MBlockIndex           MzUint
	MWantsToFinish        MzUint
	MPrevReturnStatus     TdeflStatus
	MPInBuf               c.Pointer
	MPOutBuf              c.Pointer
	MPInBufSize           *c.SizeT
	MPOutBufSize          *c.SizeT
	MFlush                TdeflFlush
	MPSrc                 *MzUint8
	MSrcBufLeft           c.SizeT
	MOutBufOfs            c.SizeT
	MDict                 [33025]MzUint8
	MHuffCount            [3][288]MzUint16
	MHuffCodes            [3][288]MzUint16
	MHuffCodeSizes        [3][288]MzUint8
	MLzCodeBuf            [24576]MzUint8
	MNext                 [32768]MzUint16
	MHash                 [4096]MzUint16
	MOutputBuf            [31948]MzUint8
}

// Initializes the compressor.
// There is no corresponding deinit() function because the tdefl API's do not dynamically allocate memory.
// pBut_buf_func: If **not** NULL, output data will be supplied to the specified callback. In this case, the user should call the tdefl_compress_buffer() API for compression.
// If pBut_buf_func is NULL the user should always call the tdefl_compress() API.
// flags: See the above enums (TDEFL_HUFFMAN_ONLY, TDEFL_WRITE_ZLIB_HEADER, etc.)
// llgo:link (*TdeflCompressor).TdeflInit C.tdefl_init
func (recv_ *TdeflCompressor) TdeflInit(pPut_buf_func TdeflPutBufFuncPtr, pPut_buf_user c.Pointer, flags c.Int) TdeflStatus {
	return 0
}

// Compresses a block of data, consuming as much of the specified input buffer as possible, and writing as much compressed data to the specified output buffer as possible.
// llgo:link (*TdeflCompressor).TdeflCompress C.tdefl_compress
func (recv_ *TdeflCompressor) TdeflCompress(pIn_buf c.Pointer, pIn_buf_size *c.SizeT, pOut_buf c.Pointer, pOut_buf_size *c.SizeT, flush TdeflFlush) TdeflStatus {
	return 0
}

// tdefl_compress_buffer() is only usable when the tdefl_init() is called with a non-NULL tdefl_put_buf_func_ptr.
// tdefl_compress_buffer() always consumes the entire input buffer.
// llgo:link (*TdeflCompressor).TdeflCompressBuffer C.tdefl_compress_buffer
func (recv_ *TdeflCompressor) TdeflCompressBuffer(pIn_buf c.Pointer, in_buf_size c.SizeT, flush TdeflFlush) TdeflStatus {
	return 0
}

// llgo:link (*TdeflCompressor).TdeflGetPrevReturnStatus C.tdefl_get_prev_return_status
func (recv_ *TdeflCompressor) TdeflGetPrevReturnStatus() TdeflStatus {
	return 0
}

// llgo:link (*TdeflCompressor).TdeflGetAdler32 C.tdefl_get_adler32
func (recv_ *TdeflCompressor) TdeflGetAdler32() MzUint32 {
	return 0
}
