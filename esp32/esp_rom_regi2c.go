package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Read internal register
 *
 * @param block     Block of the register
 * @param host_id   Host of the register
 * @param reg_add   Address of the register
 * @return uint8_t  Register value
 */
//go:linkname EspRomRegi2cRead C.esp_rom_regi2c_read
func EspRomRegi2cRead(block c.Uint8T, host_id c.Uint8T, reg_add c.Uint8T) c.Uint8T

/**
 * @brief Read internal register, in bits
 *
 * @param block     Block of the register
 * @param host_id   Host of the register
 * @param reg_add   Address of the register
 * @param msb       MSB of the register
 * @param lsb       LSB of the register
 * @return uint8_t  Register value
 */
//go:linkname EspRomRegi2cReadMask C.esp_rom_regi2c_read_mask
func EspRomRegi2cReadMask(block c.Uint8T, host_id c.Uint8T, reg_add c.Uint8T, msb c.Uint8T, lsb c.Uint8T) c.Uint8T

/**
 * @brief Write internal register
 *
 * @param block     Block of the register
 * @param host_id   Host of the register
 * @param reg_add   Address of the register
 * @param data      Value to write
 */
//go:linkname EspRomRegi2cWrite C.esp_rom_regi2c_write
func EspRomRegi2cWrite(block c.Uint8T, host_id c.Uint8T, reg_add c.Uint8T, data c.Uint8T)

/**
 * @brief Write internal register, in bits
 *
 * @param block     Block of the register
 * @param host_id   Host of the register
 * @param reg_add   Address of the register
 * @param msb       MSB of the register
 * @param lsb       LSB of the register
 * @param data      Value to write
 */
//go:linkname EspRomRegi2cWriteMask C.esp_rom_regi2c_write_mask
func EspRomRegi2cWriteMask(block c.Uint8T, host_id c.Uint8T, reg_add c.Uint8T, msb c.Uint8T, lsb c.Uint8T, data c.Uint8T)
