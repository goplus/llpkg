package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const MACSTR = "%02x:%02x:%02x:%02x:%02x:%02x"
const ONE_UNIVERSAL_MAC_ADDR = 1
const TWO_UNIVERSAL_MAC_ADDR = 2
const FOUR_UNIVERSAL_MAC_ADDR = 4

type EspMacTypeT c.Int

const (
	ESP_MAC_WIFI_STA      EspMacTypeT = 0
	ESP_MAC_WIFI_SOFTAP   EspMacTypeT = 1
	ESP_MAC_BT            EspMacTypeT = 2
	ESP_MAC_ETH           EspMacTypeT = 3
	ESP_MAC_IEEE802154    EspMacTypeT = 4
	ESP_MAC_BASE          EspMacTypeT = 5
	ESP_MAC_EFUSE_FACTORY EspMacTypeT = 6
	ESP_MAC_EFUSE_CUSTOM  EspMacTypeT = 7
	ESP_MAC_EFUSE_EXT     EspMacTypeT = 8
)

/**
 * @brief  Set base MAC address with the MAC address which is stored in BLK3 of EFUSE or
 *         external storage e.g. flash and EEPROM.
 *
 * Base MAC address is used to generate the MAC addresses used by network interfaces.
 *
 * If using a custom base MAC address, call this API before initializing any network interfaces.
 * Refer to the ESP-IDF Programming Guide for details about how the Base MAC is used.
 *
 * @note Base MAC must be a unicast MAC (least significant bit of first byte must be zero).
 *
 * @note If not using a valid OUI, set the "locally administered" bit
 *       (bit value 0x02 in the first byte) to avoid collisions.
 *
 * @param  mac base MAC address, length: 6 bytes.
 *         length: 6 bytes for MAC-48
 *
 * @return ESP_OK on success
 *         ESP_ERR_INVALID_ARG If mac is NULL or is not a unicast MAC
 */
//go:linkname EspBaseMacAddrSet C.esp_base_mac_addr_set
func EspBaseMacAddrSet(mac *c.Uint8T) EspErrT

/**
 * @brief  Return base MAC address which is set using esp_base_mac_addr_set.
 *
 * @note If no custom Base MAC has been set, this returns the pre-programmed Espressif base MAC address.
 *
 * @param  mac base MAC address, length: 6 bytes.
 *         length: 6 bytes for MAC-48
 *
 * @return ESP_OK on success
 *         ESP_ERR_INVALID_ARG mac is NULL
 *         ESP_ERR_INVALID_MAC base MAC address has not been set
 */
//go:linkname EspBaseMacAddrGet C.esp_base_mac_addr_get
func EspBaseMacAddrGet(mac *c.Uint8T) EspErrT

/**
 * @brief  Return base MAC address which was previously written to BLK3 of EFUSE.
 *
 * Base MAC address is used to generate the MAC addresses used by the networking interfaces.
 * This API returns the custom base MAC address which was previously written to EFUSE BLK3 in
 * a specified format.
 *
 * Writing this EFUSE allows setting of a different (non-Espressif) base MAC address. It is also
 * possible to store a custom base MAC address elsewhere, see esp_base_mac_addr_set() for details.
 *
 * @note This function is currently only supported on ESP32.
 *
 * @param  mac base MAC address, length: 6 bytes/8 bytes.
 *         length: 6 bytes for MAC-48
 *                 8 bytes for EUI-64(used for IEEE 802.15.4, if CONFIG_SOC_IEEE802154_SUPPORTED=y)
 *
 * @return ESP_OK on success
 *         ESP_ERR_INVALID_ARG mac is NULL
 *         ESP_ERR_INVALID_MAC CUSTOM_MAC address has not been set, all zeros (for esp32-xx)
 *         ESP_ERR_INVALID_VERSION An invalid MAC version field was read from BLK3 of EFUSE (for esp32)
 *         ESP_ERR_INVALID_CRC An invalid MAC CRC was read from BLK3 of EFUSE (for esp32)
 */
//go:linkname EspEfuseMacGetCustom C.esp_efuse_mac_get_custom
func EspEfuseMacGetCustom(mac *c.Uint8T) EspErrT

/**
 * @brief  Return base MAC address which is factory-programmed by Espressif in EFUSE.
 *
 * @param  mac base MAC address, length: 6 bytes/8 bytes.
 *         length: 6 bytes for MAC-48
 *                 8 bytes for EUI-64(used for IEEE 802.15.4, if CONFIG_SOC_IEEE802154_SUPPORTED=y)
 *
 * @return ESP_OK on success
 *         ESP_ERR_INVALID_ARG mac is NULL
 */
//go:linkname EspEfuseMacGetDefault C.esp_efuse_mac_get_default
func EspEfuseMacGetDefault(mac *c.Uint8T) EspErrT

/**
 * @brief  Read base MAC address and set MAC address of the interface.
 *
 * This function first get base MAC address using esp_base_mac_addr_get().
 * Then calculates the MAC address of the specific interface requested,
 * refer to ESP-IDF Programming Guide for the algorithm.
 *
 * The MAC address set by the esp_iface_mac_addr_set() function will not depend on the base MAC address.
 *
 * @param  mac base MAC address, length: 6 bytes/8 bytes.
 *         length: 6 bytes for MAC-48
 *                 8 bytes for EUI-64(used for IEEE 802.15.4, if CONFIG_SOC_IEEE802154_SUPPORTED=y)
 * @param  type Type of MAC address to return
 *
 * @return ESP_OK on success
 */
//go:linkname EspReadMac C.esp_read_mac
func EspReadMac(mac *c.Uint8T, type_ EspMacTypeT) EspErrT

/**
 * @brief Derive local MAC address from universal MAC address.
 *
 * This function copies a universal MAC address and then sets the "locally
 * administered" bit (bit 0x2) in the first octet, creating a locally
 * administered MAC address.
 *
 * If the universal MAC address argument is already a locally administered MAC
 * address, then the first octet is XORed with 0x4 in order to create a different
 * locally administered MAC address.
 *
 * @param  local_mac base MAC address, length: 6 bytes.
 *         length: 6 bytes for MAC-48
 * @param  universal_mac  Source universal MAC address, length: 6 bytes.
 *
 * @return ESP_OK on success
 */
//go:linkname EspDeriveLocalMac C.esp_derive_local_mac
func EspDeriveLocalMac(local_mac *c.Uint8T, universal_mac *c.Uint8T) EspErrT

/**
 * @brief  Set custom MAC address of the interface. This function allows you to overwrite the MAC addresses
 *         of the interfaces set by the base MAC address.
 *
 * @param  mac  MAC address, length: 6 bytes/8 bytes.
 *         length: 6 bytes for MAC-48
 *                 8 bytes for EUI-64(used for ESP_MAC_IEEE802154 type, if CONFIG_SOC_IEEE802154_SUPPORTED=y)
 * @param  type  Type of MAC address
 *
 * @return ESP_OK on success
 */
//go:linkname EspIfaceMacAddrSet C.esp_iface_mac_addr_set
func EspIfaceMacAddrSet(mac *c.Uint8T, type_ EspMacTypeT) EspErrT

/**
 * @brief  Return the size of the MAC type in bytes.
 *
 * If CONFIG_SOC_IEEE802154_SUPPORTED is set then for these types:
 *  - ESP_MAC_IEEE802154 is 8 bytes.
 *  - ESP_MAC_BASE, ESP_MAC_EFUSE_FACTORY and ESP_MAC_EFUSE_CUSTOM the MAC size is 6 bytes.
 *  - ESP_MAC_EFUSE_EXT is 2 bytes.
 * If CONFIG_SOC_IEEE802154_SUPPORTED is not set then for all types it returns 6 bytes.
 *
 * @param  type  Type of MAC address
 *
 * @return 0 MAC type not found (not supported)
 *         6 bytes for MAC-48.
 *         8 bytes for EUI-64.
 */
// llgo:link EspMacTypeT.EspMacAddrLenGet C.esp_mac_addr_len_get
func (recv_ EspMacTypeT) EspMacAddrLenGet() c.SizeT {
	return 0
}
