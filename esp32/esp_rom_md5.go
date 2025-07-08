package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const ESP_ROM_MD5_DIGEST_LEN = 16

/**
 * @brief Type defined for MD5 context
 *
 */

type MD5Context struct {
	Buf  [4]c.Uint32T
	Bits [2]c.Uint32T
	In   [64]c.Uint8T
}
type Md5ContextT MD5Context

/**
 * @brief Initialize the MD5 context
 *
 * @param context Context object allocated by user
 */
// llgo:link (*Md5ContextT).EspRomMd5Init C.esp_rom_md5_init
func (recv_ *Md5ContextT) EspRomMd5Init() {
}

/**
 * @brief Running MD5 algorithm over input data
 *
 * @param context MD5 context which has been initialized by `MD5Init`
 * @param buf Input buffer
 * @param len Buffer length in bytes
 */
// llgo:link (*Md5ContextT).EspRomMd5Update C.esp_rom_md5_update
func (recv_ *Md5ContextT) EspRomMd5Update(buf c.Pointer, len c.Uint32T) {
}

/**
 * @brief Extract the MD5 result, and erase the context
 *
 * @param digest Where to store the 128-bit digest value
 * @param context MD5 context
 */
//go:linkname EspRomMd5Final C.esp_rom_md5_final
func EspRomMd5Final(digest *c.Uint8T, context *Md5ContextT)
