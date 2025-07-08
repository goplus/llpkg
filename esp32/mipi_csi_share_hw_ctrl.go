package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type MipiCsiBrgUserT c.Int

const (
	MIPI_CSI_BRG_USER_CSI     MipiCsiBrgUserT = 0
	MIPI_CSI_BRG_USER_ISP_DVP MipiCsiBrgUserT = 1
	MIPI_CSI_BRG_USER_SHARE   MipiCsiBrgUserT = 2
)

/**
 * @brief Claim MIPI CSI Bridge peripheral
 *
 * @param[in]  user    CSI Bridge user
 * @param[out] out_id  ID of the CSI Bridge
 *
 * @return
 *        - ESP_OK
 *        - ESP_ERR_NOT_FOUND  No free CSI Bridge
 */
// llgo:link MipiCsiBrgUserT.MipiCsiBrgClaim C.mipi_csi_brg_claim
func (recv_ MipiCsiBrgUserT) MipiCsiBrgClaim(out_id *c.Int) EspErrT {
	return 0
}

/**
 * @brief Declaim MIPI CSI Bridge peripheral
 *
 * @param[in] id  ID of the CSI Bridge
 *
 * @return
 *        - ESP_OK
 *        - ESP_ERR_INVALID_STATE  CSI Bridge isn't claimed yet
 */
//go:linkname MipiCsiBrgDeclaim C.mipi_csi_brg_declaim
func MipiCsiBrgDeclaim(id c.Int) EspErrT
