package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Check if the pointer is in external ram dma capable region
 *
 * @param p pointer
 *
 * @return true: capable; false: not capable
 */
//go:linkname EspPtrDmaExtCapable C.esp_ptr_dma_ext_capable
func EspPtrDmaExtCapable(p c.Pointer) bool

/**
 * @brief Check if the pointer is executable
 *
 * @param p pointer
 *
 * @return true: is executable; false: not executable
 */
//go:linkname EspPtrExecutable C.esp_ptr_executable
func EspPtrExecutable(p c.Pointer) bool

/**
 * @brief Check if the pointer is byte accessible
 *
 * @param p pointer
 *
 * @return true: is byte accessible; false: not byte accessible
 */
//go:linkname EspPtrByteAccessible C.esp_ptr_byte_accessible
func EspPtrByteAccessible(p c.Pointer) bool

/**
 * @brief Check if the pointer is in external ram
 *
 * @param p pointer
 *
 * @return true: is in external ram; false: not in external ram
 */
//go:linkname EspPtrExternalRam C.esp_ptr_external_ram
func EspPtrExternalRam(p c.Pointer) bool
