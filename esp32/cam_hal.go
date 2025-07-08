package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type CamDevT LcdCamDevT

/**
 * @brief CAM hardware interface object data
 */

type CamHalContext struct {
	Hw *CamDevT
}
type CamHalContextT CamHalContext

/**
 * @brief CAM HAL driver configuration
 */

type CamHalConfig struct {
	Port       c.Int
	ByteSwapEn bool
}
type CamHalConfigT CamHalConfig

/**
 * @brief Initialize CAM hardware
 *
 * @param hal    CAM object data pointer
 * @param config CAM configuration
 *
 * @return None
 */
// llgo:link (*CamHalContextT).CamHalInit C.cam_hal_init
func (recv_ *CamHalContextT) CamHalInit(config *CamHalConfigT) {
}

/**
 * @brief De-initialize CAM hardware
 *
 * @param hal CAM object data pointer
 *
 * @return None
 */
// llgo:link (*CamHalContextT).CamHalDeinit C.cam_hal_deinit
func (recv_ *CamHalContextT) CamHalDeinit() {
}

/**
 * @brief Start CAM to receive frame data
 *
 * @param hal CAM object data pointer
 *
 * @return None
 */
// llgo:link (*CamHalContextT).CamHalStartStreaming C.cam_hal_start_streaming
func (recv_ *CamHalContextT) CamHalStartStreaming() {
}

/**
 * @brief Stop CAM receiving frame data
 *
 * @param hal CAM object data pointer
 *
 * @return None
 */
// llgo:link (*CamHalContextT).CamHalStopStreaming C.cam_hal_stop_streaming
func (recv_ *CamHalContextT) CamHalStopStreaming() {
}
