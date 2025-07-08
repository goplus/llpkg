package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type ImageProcessDriverS struct {
	Unused [8]uint8
}
type ImageProcessDriverT ImageProcessDriverS

/**
 * @brief Image process flow
 * @note This API first reads the image header, then process the segments from the image header.
 *       This API can be further inserted with more steps about the image processing by registering
 *       more function pointer in `image_process_driver_t`.
 *
 * @return
 *        - ESP_OK
 *        - ESP_FAIL: image process flow fails
 */
//go:linkname ImageProcess C.image_process
func ImageProcess() EspErrT

/**
 * @brief get flash segments info, only available after image_process() has been called
 *
 * @param[out] out_drom_paddr_start  drom paddr start
 * @param[out] out_irom_paddr_start  irom paddr start
 */
//go:linkname ImageProcessGetFlashSegmentsInfo C.image_process_get_flash_segments_info
func ImageProcessGetFlashSegmentsInfo(out_drom_paddr_start *c.Uint32T, out_irom_paddr_start *c.Uint32T)
