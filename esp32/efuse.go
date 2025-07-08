package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const EFUSE_SPICONFIG_SPI_DEFAULTS = 0
const EFUSE_SPICONFIG_HSPI_DEFAULTS = 1
const EFUSE_SPICONFIG_RET_SPICLK_MASK = 0x3f
const EFUSE_SPICONFIG_RET_SPICLK_SHIFT = 0
const EFUSE_SPICONFIG_RET_SPIQ_MASK = 0x3f
const EFUSE_SPICONFIG_RET_SPIQ_SHIFT = 6
const EFUSE_SPICONFIG_RET_SPID_MASK = 0x3f
const EFUSE_SPICONFIG_RET_SPID_SHIFT = 12
const EFUSE_SPICONFIG_RET_SPICS0_MASK = 0x3f
const EFUSE_SPICONFIG_RET_SPICS0_SHIFT = 18
const EFUSE_SPICONFIG_RET_SPIHD_MASK = 0x3f
const EFUSE_SPICONFIG_RET_SPIHD_SHIFT = 24

type EtsEfusePurposeT c.Int

const (
	ETS_EFUSE_KEY_PURPOSE_USER                        EtsEfusePurposeT = 0
	ETS_EFUSE_KEY_PURPOSE_RESERVED                    EtsEfusePurposeT = 1
	ETS_EFUSE_KEY_PURPOSE_XTS_AES_256_KEY_1           EtsEfusePurposeT = 2
	ETS_EFUSE_KEY_PURPOSE_XTS_AES_256_KEY_2           EtsEfusePurposeT = 3
	ETS_EFUSE_KEY_PURPOSE_XTS_AES_128_KEY             EtsEfusePurposeT = 4
	ETS_EFUSE_KEY_PURPOSE_HMAC_DOWN_ALL               EtsEfusePurposeT = 5
	ETS_EFUSE_KEY_PURPOSE_HMAC_DOWN_JTAG              EtsEfusePurposeT = 6
	ETS_EFUSE_KEY_PURPOSE_HMAC_DOWN_DIGITAL_SIGNATURE EtsEfusePurposeT = 7
	ETS_EFUSE_KEY_PURPOSE_HMAC_UP                     EtsEfusePurposeT = 8
	ETS_EFUSE_KEY_PURPOSE_SECURE_BOOT_DIGEST0         EtsEfusePurposeT = 9
	ETS_EFUSE_KEY_PURPOSE_SECURE_BOOT_DIGEST1         EtsEfusePurposeT = 10
	ETS_EFUSE_KEY_PURPOSE_SECURE_BOOT_DIGEST2         EtsEfusePurposeT = 11
	ETS_EFUSE_KEY_PURPOSE_MAX                         EtsEfusePurposeT = 12
)

type EtsEfuseBlockT c.Int

const (
	ETS_EFUSE_BLOCK0         EtsEfuseBlockT = 0
	ETS_EFUSE_MAC_SPI_SYS_0  EtsEfuseBlockT = 1
	ETS_EFUSE_BLOCK_SYS_DATA EtsEfuseBlockT = 2
	ETS_EFUSE_BLOCK_USR_DATA EtsEfuseBlockT = 3
	ETS_EFUSE_BLOCK_KEY0     EtsEfuseBlockT = 4
	ETS_EFUSE_BLOCK_KEY1     EtsEfuseBlockT = 5
	ETS_EFUSE_BLOCK_KEY2     EtsEfuseBlockT = 6
	ETS_EFUSE_BLOCK_KEY3     EtsEfuseBlockT = 7
	ETS_EFUSE_BLOCK_KEY4     EtsEfuseBlockT = 8
	ETS_EFUSE_BLOCK_KEY5     EtsEfuseBlockT = 9
	ETS_EFUSE_BLOCK_KEY6     EtsEfuseBlockT = 10
	ETS_EFUSE_BLOCK_MAX      EtsEfuseBlockT = 11
)

/**
 * @brief set timing according the apb clock, so no read error or write error happens.
 *
 * @param clock: apb clock in HZ, only accept 5M(in FPGA), 10M(in FPGA), 20M, 40M, 80M.
 *
 * @return : 0 if success, others if clock not accepted
 */
//go:linkname EtsEfuseSetTiming C.ets_efuse_set_timing
func EtsEfuseSetTiming(clock c.Uint32T) c.Int

/**
 * @brief  Efuse read operation: copies data from physical efuses to efuse read registers.
 *
 * @param  null
 *
 * @return : 0 if success, others if apb clock is not accepted
 */
//go:linkname EtsEfuseRead C.ets_efuse_read
func EtsEfuseRead() c.Int

/**
 * @brief  Efuse write operation: Copies data from efuse write registers to efuse. Operates on a single block of efuses at a time.
 *
 * @note This function does not update read efuses, call ets_efuse_read() once all programming is complete.
 *
 * @return : 0 if success, others if apb clock is not accepted
 */
// llgo:link EtsEfuseBlockT.EtsEfuseProgram C.ets_efuse_program
func (recv_ EtsEfuseBlockT) EtsEfuseProgram() c.Int {
	return 0
}

/**
 * @brief Set all Efuse program registers to zero.
 *
 * Call this before writing new data to the program registers.
 */
//go:linkname EtsEfuseClearProgramRegisters C.ets_efuse_clear_program_registers
func EtsEfuseClearProgramRegisters()

/**
 * @brief Program a block of key data to an efuse block
 *
 * @param key_block Block to read purpose for. Must be in range ETS_EFUSE_BLOCK_KEY0 to ETS_EFUSE_BLOCK_KEY6. Key block must be unused (@ref ets_efuse_key_block_unused).
 * @param purpose Purpose to set for this key. Purpose must be already unset.
 * @param data Pointer to data to write.
 * @param data_len Length of data to write.
 *
 * @note This function also calls ets_efuse_program() for the specified block, and for block 0 (setting the purpose)
 */
// llgo:link EtsEfuseBlockT.EtsEfuseWriteKey C.ets_efuse_write_key
func (recv_ EtsEfuseBlockT) EtsEfuseWriteKey(purpose EtsEfusePurposeT, data c.Pointer, data_len c.SizeT) c.Int {
	return 0
}

/* @brief Return the address of a particular efuse block's first read register
 *
 * @param block Index of efuse block to look up
 *
 * @return 0 if block is invalid, otherwise a numeric read register address
 * of the first word in the block.
 */
// llgo:link EtsEfuseBlockT.EtsEfuseGetReadRegisterAddress C.ets_efuse_get_read_register_address
func (recv_ EtsEfuseBlockT) EtsEfuseGetReadRegisterAddress() c.Uint32T {
	return 0
}

/**
 * @brief Return the current purpose set for an efuse key block
 *
 * @param key_block Block to read purpose for. Must be in range ETS_EFUSE_BLOCK_KEY0 to ETS_EFUSE_BLOCK_KEY6.
 */
// llgo:link EtsEfuseBlockT.EtsEfuseGetKeyPurpose C.ets_efuse_get_key_purpose
func (recv_ EtsEfuseBlockT) EtsEfuseGetKeyPurpose() EtsEfusePurposeT {
	return 0
}

/**
 * @brief Find a key block with the particular purpose set
 *
 * @param purpose Purpose to search for.
 * @param[out] key_block Pointer which will be set to the key block if found. Can be NULL, if only need to test the key block exists.
 * @return true if found, false if not found. If false, value at key_block pointer is unchanged.
 */
// llgo:link EtsEfusePurposeT.EtsEfuseFindPurpose C.ets_efuse_find_purpose
func (recv_ EtsEfusePurposeT) EtsEfuseFindPurpose(key_block *EtsEfuseBlockT) bool {
	return false
}

/**
 * Return true if the key block is unused, false otherwise.
 *
 * An unused key block is all zero content, not read or write protected,
 * and has purpose 0 (ETS_EFUSE_KEY_PURPOSE_USER)
 *
 * @param key_block key block to check.
 *
 * @return true if key block is unused, false if key block or used
 * or the specified block index is not a key block.
 */
// llgo:link EtsEfuseBlockT.EtsEfuseKeyBlockUnused C.ets_efuse_key_block_unused
func (recv_ EtsEfuseBlockT) EtsEfuseKeyBlockUnused() bool {
	return false
}

/**
 * @brief Search for an unused key block and return the first one found.
 *
 * See @ref ets_efuse_key_block_unused for a description of an unused key block.
 *
 * @return First unused key block, or ETS_EFUSE_BLOCK_MAX if no unused key block is found.
 */
//go:linkname EtsEfuseFindUnusedKeyBlock C.ets_efuse_find_unused_key_block
func EtsEfuseFindUnusedKeyBlock() EtsEfuseBlockT

/**
 * @brief Return the number of unused efuse key blocks (0-6)
 */
//go:linkname EtsEfuseCountUnusedKeyBlocks C.ets_efuse_count_unused_key_blocks
func EtsEfuseCountUnusedKeyBlocks() c.Uint

/**
 * @brief Calculate Reed-Solomon Encoding values for a block of efuse data.
 *
 * @param data Pointer to data buffer (length 32 bytes)
 * @param rs_values Pointer to write encoded data to (length 12 bytes)
 */
//go:linkname EtsEfuseRsCalculate C.ets_efuse_rs_calculate
func EtsEfuseRsCalculate(data c.Pointer, rs_values c.Pointer)

/**
 * @brief  Read spi flash pads configuration from Efuse
 *
 * @return
 * - 0 for default SPI pins.
 * - 1 for default HSPI pins.
 * - Other values define a custom pin configuration mask. Pins are encoded as per the EFUSE_SPICONFIG_RET_SPICLK,
 *   EFUSE_SPICONFIG_RET_SPIQ, EFUSE_SPICONFIG_RET_SPID, EFUSE_SPICONFIG_RET_SPICS0, EFUSE_SPICONFIG_RET_SPIHD macros.
 *   WP pin (for quad I/O modes) is not saved in efuse and not returned by this function.
 */
//go:linkname EtsEfuseGetSpiconfig C.ets_efuse_get_spiconfig
func EtsEfuseGetSpiconfig() c.Uint32T

/**
 * @brief  Read spi flash wp pad from Efuse
 *
 * @return
 * - 0x3f for invalid.
 * - 0~46 is valid.
 */
//go:linkname EtsEfuseGetWpPad C.ets_efuse_get_wp_pad
func EtsEfuseGetWpPad() c.Uint32T

/**
 * @brief  Read if download mode disabled from Efuse
 *
 * @return
 * - true for efuse disable download mode.
 * - false for efuse doesn't disable download mode.
 */
//go:linkname EtsEfuseDownloadModesDisabled C.ets_efuse_download_modes_disabled
func EtsEfuseDownloadModesDisabled() bool

/**
 * @brief  Read if legacy spi flash boot mode disabled from Efuse
 *
 * @return
 * - true for efuse disable legacy spi flash boot mode.
 * - false for efuse doesn't disable legacy spi flash boot mode.
 */
//go:linkname EtsEfuseLegacySpiBootModeDisabled C.ets_efuse_legacy_spi_boot_mode_disabled
func EtsEfuseLegacySpiBootModeDisabled() bool

/**
 * @brief  Read if uart print control value from Efuse
 *
 * @return
 * - 0 for uart force print.
 * - 1 for uart print when GPIO46 is low when digital reset.
 *   2 for uart print when GPIO46 is high when digital reset.
 *   3 for uart force silent
 */
//go:linkname EtsEfuseGetUartPrintControl C.ets_efuse_get_uart_print_control
func EtsEfuseGetUartPrintControl() c.Uint32T

/**
 * @brief  Read if USB-Serial-JTAG print during rom boot is disabled from Efuse
 *
 * @return
 * - 1 for efuse disable USB-Serial-JTAG print during rom boot.
 * - 0 for efuse doesn't disable USB-Serial-JTAG print during rom boot.
 */
//go:linkname EtsEfuseUsbSerialJtagPrintIsDisabled C.ets_efuse_usb_serial_jtag_print_is_disabled
func EtsEfuseUsbSerialJtagPrintIsDisabled() c.Uint32T

/**
 * @brief  Read if usb download mode disabled from Efuse
 *
 * (Also returns true if security download mode is enabled, as this mode
 * disables USB download.)
 *
 * @return
 * - true for efuse disable usb download mode.
 * - false for efuse doesn't disable usb download mode.
 */
//go:linkname EtsEfuseUsbDownloadModeDisabled C.ets_efuse_usb_download_mode_disabled
func EtsEfuseUsbDownloadModeDisabled() bool

/**
 * @brief  Read if usb module disabled from Efuse
 *
 * @return
 * - true for efuse disable usb module.
 * - false for efuse doesn't disable usb module.
 */
//go:linkname EtsEfuseUsbModuleDisabled C.ets_efuse_usb_module_disabled
func EtsEfuseUsbModuleDisabled() bool

/**
 * @brief  Read if security download modes enabled from Efuse
 *
 * @return
 * - true for efuse enable security download mode.
 * - false for efuse doesn't enable security download mode.
 */
//go:linkname EtsEfuseSecurityDownloadModesEnabled C.ets_efuse_security_download_modes_enabled
func EtsEfuseSecurityDownloadModesEnabled() bool

/**
 * @brief Return true if secure boot is enabled in EFuse
 */
//go:linkname EtsEfuseSecureBootEnabled C.ets_efuse_secure_boot_enabled
func EtsEfuseSecureBootEnabled() bool

/**
 * @brief Return true if secure boot aggressive revoke is enabled in EFuse
 */
//go:linkname EtsEfuseSecureBootAggressiveRevokeEnabled C.ets_efuse_secure_boot_aggressive_revoke_enabled
func EtsEfuseSecureBootAggressiveRevokeEnabled() bool

/**
 * @brief Return true if cache encryption (flash, PSRAM, etc) is enabled from boot via EFuse
 */
//go:linkname EtsEfuseCacheEncryptionEnabled C.ets_efuse_cache_encryption_enabled
func EtsEfuseCacheEncryptionEnabled() bool

/**
 * @brief Return true if OPI pins GPIO33-37 are powered by VDDSPI, otherwise by VDD33CPU
 */
//go:linkname EtsEfuseFlashOpi5padsPowerSelVddspi C.ets_efuse_flash_opi_5pads_power_sel_vddspi
func EtsEfuseFlashOpi5padsPowerSelVddspi() bool

/**
 * @brief Return true if EFuse indicates to send a flash resume command.
 */
//go:linkname EtsEfuseForceSendResume C.ets_efuse_force_send_resume
func EtsEfuseForceSendResume() bool

/**
 * @brief  return the time in us ROM boot need wait flash to power on from Efuse
 *
 * @return
 * - uint32_t the time in us.
 */
//go:linkname EtsEfuseGetFlashDelayUs C.ets_efuse_get_flash_delay_us
func EtsEfuseGetFlashDelayUs() c.Uint32T

/**
 * @brief Enable JTAG temporarily by writing a JTAG HMAC "key" into
 * the JTAG_CTRL registers.
 *
 * Works if JTAG has been "soft" disabled by burning the EFUSE_SOFT_DIS_JTAG efuse.
 *
 * Will enable the HMAC module to generate a "downstream" HMAC value from a key already saved in efuse, and then write the JTAG HMAC "key" which will enable JTAG if the two keys match.
 *
 * @param jtag_hmac_key Pointer to a 32 byte array containing a valid key. Supplied by user.
 * @param key_block Index of a key block containing the source for this key.
 *
 * @return ETS_FAILED if HMAC operation fails or invalid parameter, ETS_OK otherwise. ETS_OK doesn't necessarily mean that JTAG was enabled.
 */
//go:linkname EtsJtagEnableTemporarily C.ets_jtag_enable_temporarily
func EtsJtagEnableTemporarily(jtag_hmac_key *c.Uint8T, key_block EtsEfuseBlockT) c.Int

/**
 * @brief  A crc8 algorithm used for MAC addresses in efuse
 *
 * @param  unsigned char const *p : Pointer to original data.
 *
 * @param  unsigned int len : Data length in byte.
 *
 * @return unsigned char: Crc value.
 */
//go:linkname EspCrc8 C.esp_crc8
func EspCrc8(p *c.Char, len c.Uint) c.Char
