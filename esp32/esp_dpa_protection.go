package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type EspCryptoDpaSecLevelT c.Int

const (
	ESP_CRYPTO_DPA_SEC_LEVEL_OFF    EspCryptoDpaSecLevelT = 0
	ESP_CRYPTO_DPA_SEC_LEVEL_LOW    EspCryptoDpaSecLevelT = 1
	ESP_CRYPTO_DPA_SEC_LEVEL_MIDDLE EspCryptoDpaSecLevelT = 2
	ESP_CRYPTO_DPA_SEC_LEVEL_HIGH   EspCryptoDpaSecLevelT = 3
)

/**
 * @brief Enable DPA (Differential Power Analysis) related protection
 *
 * @note
 * Enabling the DPA protection can help to make it difficult to perform SCA
 * attacks on the crypto peripherals. However, based on the security level
 * set there will be a performance impact, higher the level higher the impact.
 * Please refer to the TRM for more details.
 *
 * @param level DPA Security Level of type `esp_crypto_dpa_sec_level_t`
 */
// llgo:link EspCryptoDpaSecLevelT.EspCryptoDpaProtectionEnable C.esp_crypto_dpa_protection_enable
func (recv_ EspCryptoDpaSecLevelT) EspCryptoDpaProtectionEnable() {
}

/**
 * @brief Disable DPA (Differential Power Analysis) related protection
 */
//go:linkname EspCryptoDpaProtectionDisable C.esp_crypto_dpa_protection_disable
func EspCryptoDpaProtectionDisable()
