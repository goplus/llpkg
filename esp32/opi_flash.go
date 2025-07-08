package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const FLASH_OP_MODE_RDCMD_DOUT = 0x3B
const ESP_ROM_FLASH_SECTOR_SIZE = 0x1000
const ESP_ROM_FLASH_BLOCK_SIZE_64K = 0x10000
const ESP_ROM_FLASH_PAGE_SIZE = 256
const ROM_FLASH_CMD_RDID = 0x9F
const ROM_FLASH_CMD_WRSR = 0x01
const ROM_FLASH_CMD_WRSR2 = 0x31
const ROM_FLASH_CMD_WREN = 0x06
const ROM_FLASH_CMD_WRDI = 0x04
const ROM_FLASH_CMD_RDSR = 0x05
const ROM_FLASH_CMD_RDSR2 = 0x35
const ROM_FLASH_CMD_ERASE_SEC = 0x20
const ROM_FLASH_CMD_ERASE_BLK_32K = 0x52
const ROM_FLASH_CMD_ERASE_BLK_64K = 0xD8
const ROM_FLASH_CMD_OTPEN = 0x3A
const ROM_FLASH_CMD_RSTEN = 0x66
const ROM_FLASH_CMD_RST = 0x99
const ROM_FLASH_CMD_SE4B = 0x21
const ROM_FLASH_CMD_SE4B_OCT = 0xDE21
const ROM_FLASH_CMD_BE4B = 0xDC
const ROM_FLASH_CMD_BE4B_OCT = 0x23DC
const ROM_FLASH_CMD_RSTEN_OCT = 0x9966
const ROM_FLASH_CMD_RST_OCT = 0x6699
const ROM_FLASH_CMD_FSTRD4B_STR = 0x13EC
const ROM_FLASH_CMD_FSTRD4B_DTR = 0x11EE
const ROM_FLASH_CMD_FSTRD4B = 0x0C
const ROM_FLASH_CMD_PP4B = 0x12
const ROM_FLASH_CMD_PP4B_OCT = 0xED12
const ROM_FLASH_CMD_RDID_OCT = 0x609F
const ROM_FLASH_CMD_WREN_OCT = 0xF906
const ROM_FLASH_CMD_RDSR_OCT = 0xFA05
const ROM_FLASH_CMD_RDCR2 = 0x71
const ROM_FLASH_CMD_RDCR2_OCT = 0x8E71
const ROM_FLASH_CMD_WRCR2 = 0x72
const ROM_FLASH_CMD_WRCR2_OCT = 0x8D72
const ROM_FLASH_CMD_RDFSR_GD = 0x70
const ROM_FLASH_CMD_RD_GD = 0x03
const ROM_FLASH_CMD_RD4B_GD = 0x13
const ROM_FLASH_CMD_FSTRD_GD = 0x0B
const ROM_FLASH_CMD_FSTRD4B_GD = 0x0C
const ROM_FLASH_CMD_FSTRD_OOUT_GD = 0x8B
const ROM_FLASH_CMD_FSTRD4B_OOUT_GD = 0x7C
const ROM_FLASH_CMD_FSTRD_OIOSTR_GD = 0xCB
const ROM_FLASH_CMD_FSTRD4B_OIOSTR_GD = 0xCC
const ROM_FLASH_CMD_FSTRD4B_OIODTR_GD = 0xFD
const ROM_FLASH_CMD_PP_GD = 0x02
const ROM_FLASH_CMD_PP4B_GD = 0x12
const ROM_FLASH_CMD_PP_OOUT_GD = 0x82
const ROM_FLASH_CMD_PP4B_OOUT_GD = 0x84
const ROM_FLASH_CMD_PP_OIO_GD = 0xC2
const ROM_FLASH_CMD_PP4B_OIOSTR_GD = 0x8E
const ROM_FLASH_CMD_SE_GD = 0x20
const ROM_FLASH_CMD_SE4B_GD = 0x21
const ROM_FLASH_CMD_BE32K_GD = 0x52
const ROM_FLASH_CMD_BE32K4B_GD = 0x5C
const ROM_FLASH_CMD_BE64K_GD = 0xD8
const ROM_FLASH_CMD_BE64K4B_GD = 0xDC
const ROM_FLASH_CMD_EN4B_GD = 0xB7
const ROM_FLASH_CMD_DIS4B_GD = 0xE9

type EspRomOpiflashCmdT struct {
	Mode        c.Uint8T
	CmdBitLen   c.Uint8T
	Cmd         c.Uint16T
	Addr        c.Uint32T
	AddrBitLen  c.Uint8T
	DummyBitLen c.Uint8T
	DataBitLen  c.Uint8T
	CsSel       c.Uint8T
	IsPe        c.Uint8T
}

type EspRomOpiflashSpi0rdT struct {
	AddrBitLen  c.Uint8T
	DummyBitLen c.Uint8T
	Cmd         c.Uint16T
	CmdBitLen   c.Uint8T
	VarDummyEn  c.Uint8T
}

type EspRomOpiflashDefT struct {
	Rdid       EspRomOpiflashCmdT
	Rdsr       EspRomOpiflashCmdT
	Wren       EspRomOpiflashCmdT
	Se         EspRomOpiflashCmdT
	Be64k      EspRomOpiflashCmdT
	Read       EspRomOpiflashCmdT
	Pp         EspRomOpiflashCmdT
	CacheRdCmd EspRomOpiflashSpi0rdT
}

type EspRomSpiCmdT struct {
	Cmd          c.Uint16T
	CmdBitLen    c.Uint16T
	Addr         *c.Uint32T
	AddrBitLen   c.Uint32T
	TxData       *c.Uint32T
	TxDataBitLen c.Uint32T
	RxData       *c.Uint32T
	RxDataBitLen c.Uint32T
	DummyBitLen  c.Uint32T
}

/**
 * @brief init legacy driver for Octal Flash
 */
// llgo:link (*EspRomOpiflashDefT).EspRomOpiflashLegacyDriverInit C.esp_rom_opiflash_legacy_driver_init
func (recv_ *EspRomOpiflashDefT) EspRomOpiflashLegacyDriverInit() {
}

// spi user mode command config
/**
 * @brief Config the spi user command
 * @param spi_num spi port
 * @param pcmd pointer to accept the spi command struct
 */
//go:linkname EspRomSpiCmdConfig C.esp_rom_spi_cmd_config
func EspRomSpiCmdConfig(spi_num c.Int, pcmd *EspRomSpiCmdT)

/**
 * @brief Start a spi user command sequence
 * @param spi_num spi port
 * @param rx_buf buffer pointer to receive data
 * @param rx_len receive data length in byte
 * @param cs_en_mask decide which cs to use, 0 for cs0, 1 for cs1
 * @param is_write_erase to indicate whether this is a write or erase operation, since the CPU would check permission
 */
//go:linkname EspRomSpiCmdStart C.esp_rom_spi_cmd_start
func EspRomSpiCmdStart(spi_num c.Int, rx_buf *c.Uint8T, rx_len c.Uint16T, cs_en_mask c.Uint8T, is_write_erase bool)

/**
 * @brief Config opi flash pads according to efuse settings.
 */
//go:linkname EspRomOpiflashPinConfig C.esp_rom_opiflash_pin_config
func EspRomOpiflashPinConfig()

// set SPI read/write mode
/**
 * @brief Set SPI operation mode
 * @param spi_num spi port
 * @param mode Flash Read Mode
 */
//go:linkname EspRomSpiSetOpMode C.esp_rom_spi_set_op_mode
func EspRomSpiSetOpMode(spi_num c.Int, mode EspRomSpiflashReadModeT)

/**
 * @brief Set data swap mode in DTR(DDR) mode
 * @param spi_num spi port
 * @param wr_swap to decide whether to swap fifo data in dtr write operation
 * @param rd_swap to decide whether to swap fifo data in dtr read operation
 */
//go:linkname EspRomSpiSetDtrSwapMode C.esp_rom_spi_set_dtr_swap_mode
func EspRomSpiSetDtrSwapMode(spi c.Int, wr_swap bool, rd_swap bool)

/**
 * @brief to send reset command in spi/opi-str/opi-dtr mode(for MX25UM25645G)
 * @param spi_num spi port
 */
//go:linkname EspRomOpiflashModeReset C.esp_rom_opiflash_mode_reset
func EspRomOpiflashModeReset(spi_num c.Int)

/**
 * @brief To execute a flash operation command
 * @param spi_num spi port
 * @param mode Flash Read Mode
 * @param cmd data to send in command field
 * @param cmd_bit_len bit length of command field
 * @param addr data to send in address field
 * @param addr_bit_len bit length of address field
 * @param dummy_bits bit length of dummy field
 * @param mosi_data data buffer to be sent in mosi field
 * @param mosi_bit_len bit length of data buffer to be sent in mosi field
 * @param miso_data data buffer to accept data in miso field
 * @param miso_bit_len bit length of data buffer to accept data in miso field
 * @param cs_mark decide which cs pin to use. 0: cs0, 1: cs1
 * @param is_write_erase_operation to indicate whether this a write or erase flash operation
 */
//go:linkname EspRomOpiflashExecCmd C.esp_rom_opiflash_exec_cmd
func EspRomOpiflashExecCmd(spi_num c.Int, mode EspRomSpiflashReadModeT, cmd c.Uint32T, cmd_bit_len c.Int, addr c.Uint32T, addr_bit_len c.Int, dummy_bits c.Int, mosi_data *c.Uint8T, mosi_bit_len c.Int, miso_data *c.Uint8T, miso_bit_len c.Int, cs_mask c.Uint32T, is_write_erase_operation bool)

/**
 * @brief send reset command to opi flash
 * @param spi_num spi port
 * @param mode Flash Operation Mode
 */
//go:linkname EspRomOpiflashSoftReset C.esp_rom_opiflash_soft_reset
func EspRomOpiflashSoftReset(spi_num c.Int, mode EspRomSpiflashReadModeT)

/**
 * @brief to read opi flash ID
 * @note command format would be defined in initialization
 * @param[out] out_id buffer to accept id
 * @return flash operation result
 */
//go:linkname EspRomOpiflashReadId C.esp_rom_opiflash_read_id
func EspRomOpiflashReadId(out_id *c.Uint8T) EspRomSpiflashResultT

/**
 * @brief to read opi flash status register
 * @note command format would be defined in initialization
 * @return opi flash status value
 */
//go:linkname EspRomOpiflashRdsr C.esp_rom_opiflash_rdsr
func EspRomOpiflashRdsr() c.Uint8T

/**
 * @brief wait opi flash status register to be idle
 * @note command format would be defined in initialization
 * @return flash operation result
 */
//go:linkname EspRomOpiflashWaitIdle C.esp_rom_opiflash_wait_idle
func EspRomOpiflashWaitIdle() EspRomSpiflashResultT

/**
 * @brief to erase flash sector
 * @note command format would be defined in initialization
 * @param sector_num the sector to be erased
 * @return flash operation result
 */
//go:linkname EspRomOpiflashEraseSector C.esp_rom_opiflash_erase_sector
func EspRomOpiflashEraseSector(sector_num c.Uint32T) EspRomSpiflashResultT

/**
 * @brief to erase flash block
 * @note command format would be defined in initialization
 * @param block_num the block to be erased
 * @return flash operation result
 */
//go:linkname EspRomOpiflashEraseBlock64k C.esp_rom_opiflash_erase_block_64k
func EspRomOpiflashEraseBlock64k(block_num c.Uint32T) EspRomSpiflashResultT

/**
 * @brief to erase a flash area define by start address and length
 * @note command format would be defined in initialization
 * @param start_addr the start address to be erased
 * @param area_len the erea length to be erased
 * @return flash operation result
 */
//go:linkname EspRomOpiflashEraseArea C.esp_rom_opiflash_erase_area
func EspRomOpiflashEraseArea(start_addr c.Uint32T, area_len c.Uint32T) EspRomSpiflashResultT

/**
 * @brief to read data from opi flash
 * @note command format would be defined in initialization
 * @param flash_addr flash address to read data from
 * @param data_addr data buffer to accept the data
 * @param len data length to be read
 * @return flash operation result
 */
//go:linkname EspRomOpiflashRead C.esp_rom_opiflash_read
func EspRomOpiflashRead(flash_addr c.Uint32T, data_addr c.Pointer, len c.Int) EspRomSpiflashResultT

/**
 * @brief to write data to opi flash
 * @note command format would be defined in initialization
 * @param flash_addr flash address to write data to
 * @param data_addr data buffer to write to flash
 * @param len data length to write
 * @return flash operation result
 */
//go:linkname EspRomOpiflashWrite C.esp_rom_opiflash_write
func EspRomOpiflashWrite(flash_addr c.Uint32T, data_addr *c.Uint32T, len c.Int) EspRomSpiflashResultT

/**
 * @brief send WREN command
 * @note command format would be defined in initialization
 * @param arg not used, set to NULL
 * @return flash operation result
 */
//go:linkname EspRomOpiflashWren C.esp_rom_opiflash_wren
func EspRomOpiflashWren(arg c.Pointer) EspRomSpiflashResultT

/**
 * @brief to configure SPI0 read flash command format for cache
 * @note command format would be defined in initialization
 *
 */
// llgo:link EspRomSpiflashReadModeT.EspRomOpiflashCacheModeConfig C.esp_rom_opiflash_cache_mode_config
func (recv_ EspRomSpiflashReadModeT) EspRomOpiflashCacheModeConfig(cache *EspRomOpiflashSpi0rdT) {
}

//go:linkname EspRomOpiflashReadRaw C.esp_rom_opiflash_read_raw
func EspRomOpiflashReadRaw(flash_addr c.Uint32T, buf *c.Uint8T, len c.Int) EspRomSpiflashResultT
