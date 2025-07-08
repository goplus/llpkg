package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const CPIO_MODE_FILETYPE_MASK = 0xF000
const CPIO_MODE_FILETYPE_SOCKET = 0xC000
const CPIO_MODE_FILETYPE_SYMLINK = 0xA000
const CPIO_MODE_FILETYPE_REGULAR = 0x8000
const CPIO_MODE_FILETYPE_BLOCKDEV = 0x6000
const CPIO_MODE_FILETYPE_DIR = 0x4000
const CPIO_MODE_FILETYPE_CHARDEV = 0x2000
const CPIO_MODE_FILETYPE_FIFO = 0x1000
const CPIO_MODE_SUID = 0x0800
const CPIO_MODE_SGID = 0x0400
const CPIO_MODE_STICKY = 0x0200

type CpioFileT struct {
	Filesize c.SizeT
	Name     *c.Char
	Mode     c.Uint32T
	Check    c.Uint32T
}
type CpioRetT c.Int

const (
	CPIO_RET_MORE CpioRetT = 0
	CPIO_RET_DONE CpioRetT = 1
	CPIO_RET_ERR  CpioRetT = 2
)

type CpioHandleDataT struct {
	Unused [8]uint8
}
type CpioHandleT *CpioHandleDataT
type CpioCallbackReasonT c.Int

const (
	CPIO_RSN_FILE_ALL     CpioCallbackReasonT = 0
	CPIO_RSN_FILE_INITIAL CpioCallbackReasonT = 1
	CPIO_RSN_FILE_MORE    CpioCallbackReasonT = 2
	CPIO_RSN_FILE_END     CpioCallbackReasonT = 3
)

// llgo:type C
type CpioCallbackT func(CpioCallbackReasonT, *CpioFileT, c.SizeT, c.SizeT, *c.Char, c.Pointer)

/**
 * @brief      Initialize a cpio handle.
 *
 * Call this to start parsing a cpio archive. You can set the callback that handles the
 * files/data here.
 *
 * @param  callback The callback that will handle the data of the files inside the cpio archive
 *
 * @param  cbarg User-supplied argument. The callback will be called with this as an argument.
 *
 * @param  buflen Length of internal buffer used.
 *                If this is zero, the callback will be called with data that lives in the data buffer
 *                supplied to the cpio library by whomever called cpio_feed(). Because this library has
 *                no power over that buffer, the callback can be passed as little as 1 and as many as
 *                INT_MAX bytes at a time.
 *                If this is non-zero, the library will allocate an internal buffer of this size. All
 *                cpio_feed()-calls will be rebuffered, and the callback is guaranteed to only be called
 *                with this many bytes in the buffer, given there's enough data in the file to fill it.
 *
 * @param memchunk Chunk of memory to allocate everything (handle, I/O buffer, filename buffer) in. Minimum size
 *                 (estimate) is 160+buflen+sizeof(largest filename/path).
 * @param memchunklen Size of the mem chunk
 *
 * @return
 *     - Success: A pointer to a cpio handle
 *     - Error: NULL
 *
 */
//go:linkname CpioStart C.cpio_start
func CpioStart(callback CpioCallbackT, cbarg c.Pointer, buflen c.SizeT, memchunk c.Pointer, memchunklen c.Int) CpioHandleT

/**
 * @brief      Feed data from a cpio archive into the library
 *
 * This routine is used to feed consecutive data of the cpio archive into the library. While processing,
 * the library can call the callback function one or more times if needed.
 *
 * @param  cpio Handle obtained by calling cpio_start()
 *
 * @param  buffer Pointer to buffer containing cpio archive data
 *
 * @param  len Length of the buffer, in bytes
 *
 * @return
 *     - CPIO_RET_MORE: CPIO archive isn't done yet, please feed more data.
 *     - CPIO_RET_DONE: CPUI archive is finished.
 *     - CPIO_RET_ERR: Invalid CPIO archive data; decoding aborted.
 *
 */
//go:linkname CpioFeed C.cpio_feed
func CpioFeed(cpio CpioHandleT, buffer *c.Char, len c.Int) CpioRetT

/**
 * @brief      Indicate there is no more cpio data to be fed into the archive
 *
 * This call is to be called when the source data is exhausted. Normally, the library can find the end of the
 * cpio archive by looking for the end marker,
 *
 * @param  timer_conf Pointer of LEDC timer configure struct
 *
 *
 * @return
 *     - CPIO_RET_DONE on success
 *     - CPIO_RET_ERR when cpio archive is invalid
 *
 */
//go:linkname CpioDone C.cpio_done
func CpioDone(cpio CpioHandleT) CpioRetT

/**
 * @brief      Free the memory allocated for a cpio handle.
 *
 * @param  cpio Handle obtained by calling cpio_start()
 *
 * @return
 *     - CPIO_RET_DONE on success
 *
 */
//go:linkname CpioDestroy C.cpio_destroy
func CpioDestroy(cpio CpioHandleT) CpioRetT
