package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

// llgo:type C
type PoisonFillPfuncT func(c.Pointer, c.SizeT, bool)

// llgo:type C
type PoisonCheckPfuncT func(c.Pointer, c.SizeT, bool, bool) bool

/*!
 * @brief Set the function to call for filling memory region when
 * poisoning is configured.
 *
 * @note Please keep in mind that this function in ROM still accepts void*.
 *
 * @param pfunc The callback function to trigger for poisoning
 * a memory region.
 */
//go:linkname TlsfPoisonFillPfuncSet C.tlsf_poison_fill_pfunc_set
func TlsfPoisonFillPfuncSet(pfunc PoisonFillPfuncT)

/*!
 * @brief Set the function to call for checking memory region when
 * poisoning is configured.
 *
 * @param pfunc The callback function to trigger for checking
 * the content of a memory region.
 */
//go:linkname TlsfPoisonCheckPfuncSet C.tlsf_poison_check_pfunc_set
func TlsfPoisonCheckPfuncSet(pfunc PoisonCheckPfuncT)
