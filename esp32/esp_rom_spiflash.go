package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type EspRomSpiflashReadModeT c.Int

const (
	ESP_ROM_SPIFLASH_QIO_MODE         EspRomSpiflashReadModeT = 0
	ESP_ROM_SPIFLASH_QOUT_MODE        EspRomSpiflashReadModeT = 1
	ESP_ROM_SPIFLASH_DIO_MODE         EspRomSpiflashReadModeT = 2
	ESP_ROM_SPIFLASH_DOUT_MODE        EspRomSpiflashReadModeT = 3
	ESP_ROM_SPIFLASH_FASTRD_MODE      EspRomSpiflashReadModeT = 4
	ESP_ROM_SPIFLASH_SLOWRD_MODE      EspRomSpiflashReadModeT = 5
	ESP_ROM_SPIFLASH_OPI_STR_MODE     EspRomSpiflashReadModeT = 6
	ESP_ROM_SPIFLASH_OPI_DTR_MODE     EspRomSpiflashReadModeT = 7
	ESP_ROM_SPIFLASH_OOUT_MODE        EspRomSpiflashReadModeT = 8
	ESP_ROM_SPIFLASH_OIO_STR_MODE     EspRomSpiflashReadModeT = 9
	ESP_ROM_SPIFLASH_OIO_DTR_MODE     EspRomSpiflashReadModeT = 10
	ESP_ROM_SPIFLASH_QPI_MODE         EspRomSpiflashReadModeT = 11
	ESP_ROM_SPIFLASH_OPI_HEX_DTR_MODE EspRomSpiflashReadModeT = 12
)

type EspRomSpiflashChipT struct {
	DeviceId   c.Uint32T
	ChipSize   c.Uint32T
	BlockSize  c.Uint32T
	SectorSize c.Uint32T
	PageSize   c.Uint32T
	StatusMask c.Uint32T
}
type EspRomSpiflashResultT c.Int

const (
	ESP_ROM_SPIFLASH_RESULT_OK      EspRomSpiflashResultT = 0
	ESP_ROM_SPIFLASH_RESULT_ERR     EspRomSpiflashResultT = 1
	ESP_ROM_SPIFLASH_RESULT_TIMEOUT EspRomSpiflashResultT = 2
)

/**
 * @brief SPI Flash init, clock divisor is 4, use 1 line Slow read mode.
 *    Please do not call this function in SDK.
 *
 * @param  uint32_t ishspi: 0 for spi, 1 for hspi, flash pad decided by strapping
 *              else, bit[5:0] spiclk, bit[11:6] spiq, bit[17:12] spid, bit[23:18] spics0, bit[29:24] spihd
 *
 * @param  uint8_t legacy: always keeping false.
 *
 * @return None
 */
//go:linkname EspRomSpiflashAttach C.esp_rom_spiflash_attach
func EspRomSpiflashAttach(ishspi c.Uint32T, legacy bool)

/**
 * @brief SPI Read Flash status register. We use CMD 0x05 (RDSR).
 *    Please do not call this function in SDK.
 *
 * @param  esp_rom_spiflash_chip_t *spi : The information for Flash, which is exported from ld file.
 *
 * @param  uint32_t *status : The pointer to which to return the Flash status value.
 *
 * @return ESP_ROM_SPIFLASH_RESULT_OK : read OK.
 *         ESP_ROM_SPIFLASH_RESULT_ERR : read error.
 *         ESP_ROM_SPIFLASH_RESULT_TIMEOUT : read timeout.
 */
// llgo:link (*EspRomSpiflashChipT).EspRomSpiflashReadStatus C.esp_rom_spiflash_read_status
func (recv_ *EspRomSpiflashChipT) EspRomSpiflashReadStatus(status *c.Uint32T) EspRomSpiflashResultT {
	return 0
}

/**
 * @brief SPI Read Flash status register bits 8-15. We use CMD 0x35 (RDSR2).
 *        Please do not call this function in SDK.
 *
 * @param  esp_rom_spiflash_chip_t *spi : The information for Flash, which is exported from ld file.
 *
 * @param  uint32_t *status : The pointer to which to return the Flash status value.
 *
 * @return ESP_ROM_SPIFLASH_RESULT_OK : read OK.
 *         ESP_ROM_SPIFLASH_RESULT_ERR : read error.
 *         ESP_ROM_SPIFLASH_RESULT_TIMEOUT : read timeout.
 */
// llgo:link (*EspRomSpiflashChipT).EspRomSpiflashReadStatushigh C.esp_rom_spiflash_read_statushigh
func (recv_ *EspRomSpiflashChipT) EspRomSpiflashReadStatushigh(status *c.Uint32T) EspRomSpiflashResultT {
	return 0
}

/**
 * @brief Write status to Flash status register.
 *        Please do not call this function in SDK.
 *
 * @param  esp_rom_spiflash_chip_t *spi : The information for Flash, which is exported from ld file.
 *
 * @param  uint32_t status_value : Value to .
 *
 * @return ESP_ROM_SPIFLASH_RESULT_OK : write OK.
 *         ESP_ROM_SPIFLASH_RESULT_ERR : write error.
 *         ESP_ROM_SPIFLASH_RESULT_TIMEOUT : write timeout.
 */
// llgo:link (*EspRomSpiflashChipT).EspRomSpiflashWriteStatus C.esp_rom_spiflash_write_status
func (recv_ *EspRomSpiflashChipT) EspRomSpiflashWriteStatus(status_value c.Uint32T) EspRomSpiflashResultT {
	return 0
}

/**
 * @brief Use a command to Read Flash status register.
 *        Please do not call this function in SDK.
 *
 * @param  esp_rom_spiflash_chip_t *spi : The information for Flash, which is exported from ld file.
 *
 * @param  uint32_t*status : The pointer to which to return the Flash status value.
 *
 * @return ESP_ROM_SPIFLASH_RESULT_OK : read OK.
 *         ESP_ROM_SPIFLASH_RESULT_ERR : read error.
 *         ESP_ROM_SPIFLASH_RESULT_TIMEOUT : read timeout.
 */
//go:linkname EspRomSpiflashReadUserCmd C.esp_rom_spiflash_read_user_cmd
func EspRomSpiflashReadUserCmd(status *c.Uint32T, cmd c.Uint8T) EspRomSpiflashResultT

/**
 * @brief Config SPI Flash read mode when init.
 *        Please do not call this function in SDK.
 *
 * @param  esp_rom_spiflash_read_mode_t mode : QIO/QOUT/DIO/DOUT/FastRD/SlowRD.
 *
 * This function does not try to set the QIO Enable bit in the status register, caller is responsible for this.
 *
 * @return ESP_ROM_SPIFLASH_RESULT_OK : config OK.
 *         ESP_ROM_SPIFLASH_RESULT_ERR : config error.
 *         ESP_ROM_SPIFLASH_RESULT_TIMEOUT : config timeout.
 */
// llgo:link EspRomSpiflashReadModeT.EspRomSpiflashConfigReadmode C.esp_rom_spiflash_config_readmode
func (recv_ EspRomSpiflashReadModeT) EspRomSpiflashConfigReadmode() EspRomSpiflashResultT {
	return 0
}

/**
 * @brief Config SPI Flash clock divisor.
 *        Please do not call this function in SDK.
 *
 * @param  uint8_t freqdiv: clock divisor.
 *
 * @param  uint8_t spi: 0 for SPI0, 1 for SPI1.
 *
 * @return ESP_ROM_SPIFLASH_RESULT_OK : config OK.
 *         ESP_ROM_SPIFLASH_RESULT_ERR : config error.
 *         ESP_ROM_SPIFLASH_RESULT_TIMEOUT : config timeout.
 */
//go:linkname EspRomSpiflashConfigClk C.esp_rom_spiflash_config_clk
func EspRomSpiflashConfigClk(freqdiv c.Uint8T, spi c.Uint8T) EspRomSpiflashResultT

/**
 * @brief Update SPI Flash parameter.
 *        Please do not call this function in SDK.
 *
 * @param  uint32_t deviceId : Device ID read from SPI, the low 32 bit.
 *
 * @param  uint32_t chip_size : The Flash size.
 *
 * @param  uint32_t block_size : The Flash block size.
 *
 * @param  uint32_t sector_size : The Flash sector size.
 *
 * @param  uint32_t page_size : The Flash page size.
 *
 * @param  uint32_t status_mask : The Mask used when read status from Flash(use single CMD).
 *
 * @return ESP_ROM_SPIFLASH_RESULT_OK : Update OK.
 *         ESP_ROM_SPIFLASH_RESULT_ERR : Update error.
 *         ESP_ROM_SPIFLASH_RESULT_TIMEOUT : Update timeout.
 */
//go:linkname EspRomSpiflashConfigParam C.esp_rom_spiflash_config_param
func EspRomSpiflashConfigParam(deviceId c.Uint32T, chip_size c.Uint32T, block_size c.Uint32T, sector_size c.Uint32T, page_size c.Uint32T, status_mask c.Uint32T) EspRomSpiflashResultT

/**
 * @brief Erase whole flash chip.
 *        Please do not call this function in SDK.
 *
 * @param  None
 *
 * @return ESP_ROM_SPIFLASH_RESULT_OK : Erase OK.
 *         ESP_ROM_SPIFLASH_RESULT_ERR : Erase error.
 *         ESP_ROM_SPIFLASH_RESULT_TIMEOUT : Erase timeout.
 */
//go:linkname EspRomSpiflashEraseChip C.esp_rom_spiflash_erase_chip
func EspRomSpiflashEraseChip() EspRomSpiflashResultT

/**
 * @brief Erase a 64KB block of flash
 *        Uses SPI flash command D8H.
 *        Please do not call this function in SDK.
 *
 * @param  uint32_t block_num : Which block to erase.
 *
 * @return ESP_ROM_SPIFLASH_RESULT_OK : Erase OK.
 *         ESP_ROM_SPIFLASH_RESULT_ERR : Erase error.
 *         ESP_ROM_SPIFLASH_RESULT_TIMEOUT : Erase timeout.
 */
//go:linkname EspRomSpiflashEraseBlock C.esp_rom_spiflash_erase_block
func EspRomSpiflashEraseBlock(block_num c.Uint32T) EspRomSpiflashResultT

/**
 * @brief Erase a sector of flash.
 *        Uses SPI flash command 20H.
 *        Please do not call this function in SDK.
 *
 * @param  uint32_t sector_num : Which sector to erase.
 *
 * @return ESP_ROM_SPIFLASH_RESULT_OK : Erase OK.
 *         ESP_ROM_SPIFLASH_RESULT_ERR : Erase error.
 *         ESP_ROM_SPIFLASH_RESULT_TIMEOUT : Erase timeout.
 */
//go:linkname EspRomSpiflashEraseSector C.esp_rom_spiflash_erase_sector
func EspRomSpiflashEraseSector(sector_num c.Uint32T) EspRomSpiflashResultT

/**
 * @brief Erase some sectors.
 *        Please do not call this function in SDK.
 *
 * @param  uint32_t start_addr : Start addr to erase, should be sector aligned.
 *
 * @param  uint32_t area_len : Length to erase, should be sector aligned.
 *
 * @return ESP_ROM_SPIFLASH_RESULT_OK : Erase OK.
 *         ESP_ROM_SPIFLASH_RESULT_ERR : Erase error.
 *         ESP_ROM_SPIFLASH_RESULT_TIMEOUT : Erase timeout.
 */
//go:linkname EspRomSpiflashEraseArea C.esp_rom_spiflash_erase_area
func EspRomSpiflashEraseArea(start_addr c.Uint32T, area_len c.Uint32T) EspRomSpiflashResultT

/**
 * @brief Write Data to Flash, you should Erase it yourself if need.
 *        Please do not call this function in SDK.
 *
 * @param  uint32_t dest_addr : Address to write, should be 4 bytes aligned.
 *
 * @param  const uint32_t *src : The pointer to data which is to write.
 *
 * @param  uint32_t len : Length to write, should be 4 bytes aligned.
 *
 * @return ESP_ROM_SPIFLASH_RESULT_OK : Write OK.
 *         ESP_ROM_SPIFLASH_RESULT_ERR : Write error.
 *         ESP_ROM_SPIFLASH_RESULT_TIMEOUT : Write timeout.
 */
//go:linkname EspRomSpiflashWrite C.esp_rom_spiflash_write
func EspRomSpiflashWrite(dest_addr c.Uint32T, src *c.Uint32T, len c.Int32T) EspRomSpiflashResultT

/**
 * @brief Read Data from Flash, you should Erase it yourself if need.
 *        Please do not call this function in SDK.
 *
 * @param  uint32_t src_addr : Address to read, should be 4 bytes aligned.
 *
 * @param  uint32_t *dest : The buf to read the data.
 *
 * @param  uint32_t len : Length to read, should be 4 bytes aligned.
 *
 * @return ESP_ROM_SPIFLASH_RESULT_OK : Read OK.
 *         ESP_ROM_SPIFLASH_RESULT_ERR : Read error.
 *         ESP_ROM_SPIFLASH_RESULT_TIMEOUT : Read timeout.
 */
//go:linkname EspRomSpiflashRead C.esp_rom_spiflash_read
func EspRomSpiflashRead(src_addr c.Uint32T, dest *c.Uint32T, len c.Int32T) EspRomSpiflashResultT

/**
 * @brief SPI1 go into encrypto mode.
 *        Please do not call this function in SDK.
 *
 * @param  None
 *
 * @return None
 */
//go:linkname EspRomSpiflashWriteEncryptedEnable C.esp_rom_spiflash_write_encrypted_enable
func EspRomSpiflashWriteEncryptedEnable()

/**
 * @brief Prepare 32 Bytes data to encrpto writing, you should Erase it yourself if need.
 *        Please do not call this function in SDK.
 *
 * @param  uint32_t flash_addr : Address to write, should be 32 bytes aligned.
 *
 * @param  uint32_t *data : The pointer to data which is to write.
 *
 * @return ESP_ROM_SPIFLASH_RESULT_OK : Prepare OK.
 *         ESP_ROM_SPIFLASH_RESULT_ERR : Prepare error.
 *         ESP_ROM_SPIFLASH_RESULT_TIMEOUT : Prepare timeout.
 */
//go:linkname EspRomSpiflashPrepareEncryptedData C.esp_rom_spiflash_prepare_encrypted_data
func EspRomSpiflashPrepareEncryptedData(flash_addr c.Uint32T, data *c.Uint32T) EspRomSpiflashResultT

/**
 * @brief SPI1 go out of encrypto mode.
 *        Please do not call this function in SDK.
 *
 * @param  None
 *
 * @return None
 */
//go:linkname EspRomSpiflashWriteEncryptedDisable C.esp_rom_spiflash_write_encrypted_disable
func EspRomSpiflashWriteEncryptedDisable()

/**
 * @brief Write data to flash with transparent encryption.
 * @note Sectors to be written should already be erased.
 *
 * @note Please do not call this function in SDK.
 *
 * @param  uint32_t flash_addr : Address to write, should be 32 byte aligned.
 *
 * @param  uint32_t *data : The pointer to data to write. Note, this pointer must
 *                          be 32 bit aligned and the content of the data will be
 *                          modified by the encryption function.
 *
 * @param  uint32_t len : Length to write, should be 32 bytes aligned.
 *
 * @return ESP_ROM_SPIFLASH_RESULT_OK : Data written successfully.
 *         ESP_ROM_SPIFLASH_RESULT_ERR : Encryption write error.
 *         ESP_ROM_SPIFLASH_RESULT_TIMEOUT : Encrypto write timeout.
 */
//go:linkname EspRomSpiflashWriteEncrypted C.esp_rom_spiflash_write_encrypted
func EspRomSpiflashWriteEncrypted(flash_addr c.Uint32T, data *c.Uint32T, len c.Uint32T) EspRomSpiflashResultT

/** @brief Wait until SPI flash write operation is complete
 *
 * @note Please do not call this function in SDK.
 *
 * Reads the Write In Progress bit of the SPI flash status register,
 * repeats until this bit is zero (indicating write complete).
 *
 * @return ESP_ROM_SPIFLASH_RESULT_OK : Write is complete
 *         ESP_ROM_SPIFLASH_RESULT_ERR : Error while reading status.
 */
// llgo:link (*EspRomSpiflashChipT).EspRomSpiflashWaitIdle C.esp_rom_spiflash_wait_idle
func (recv_ *EspRomSpiflashChipT) EspRomSpiflashWaitIdle() EspRomSpiflashResultT {
	return 0
}

/** @brief Enable Quad I/O pin functions
 *
 * @note Please do not call this function in SDK.
 *
 * Sets the HD & WP pin functions for Quad I/O modes, based on the
 * efuse SPI pin configuration.
 *
 * @param wp_gpio_num - Number of the WP pin to reconfigure for quad I/O.
 *
 * @param spiconfig - Pin configuration, as returned from ets_efuse_get_spiconfig().
 * - If this parameter is 0, default SPI pins are used and wp_gpio_num parameter is ignored.
 * - If this parameter is 1, default HSPI pins are used and wp_gpio_num parameter is ignored.
 * - For other values, this parameter encodes the HD pin number and also the CLK pin number. CLK pin selection is used
 *   to determine if HSPI or SPI peripheral will be used (use HSPI if CLK pin is the HSPI clock pin, otherwise use SPI).
 *   Both HD & WP pins are configured via GPIO matrix to map to the selected peripheral.
 */
//go:linkname EspRomSpiflashSelectQioPins C.esp_rom_spiflash_select_qio_pins
func EspRomSpiflashSelectQioPins(wp_gpio_num c.Uint8T, spiconfig c.Uint32T)

/**
 * @brief Clear WEL bit unconditionally.
 *
 * @return always ESP_ROM_SPIFLASH_RESULT_OK
 */
//go:linkname EspRomSpiflashWriteDisable C.esp_rom_spiflash_write_disable
func EspRomSpiflashWriteDisable() EspRomSpiflashResultT

/**
 * @brief Set WREN bit.
 *
 * @param  esp_rom_spiflash_chip_t *spi : The information for Flash, which is exported from ld file.
 *
 * @return always ESP_ROM_SPIFLASH_RESULT_OK
 */
// llgo:link (*EspRomSpiflashChipT).EspRomSpiflashWriteEnable C.esp_rom_spiflash_write_enable
func (recv_ *EspRomSpiflashChipT) EspRomSpiflashWriteEnable() EspRomSpiflashResultT {
	return 0
}

/* Flash data defined in ROM*/

type EspRomSpiflashLegacyDataT struct {
	Chip         EspRomSpiflashChipT
	DummyLenPlus [3]c.Uint8T
	SigMatrix    c.Uint8T
}
