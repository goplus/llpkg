package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type EspChipModelT c.Int

const (
	CHIP_ESP32       EspChipModelT = 1
	CHIP_ESP32S2     EspChipModelT = 2
	CHIP_ESP32S3     EspChipModelT = 9
	CHIP_ESP32C3     EspChipModelT = 5
	CHIP_ESP32C2     EspChipModelT = 12
	CHIP_ESP32C6     EspChipModelT = 13
	CHIP_ESP32H2     EspChipModelT = 16
	CHIP_ESP32P4     EspChipModelT = 18
	CHIP_ESP32C61    EspChipModelT = 20
	CHIP_ESP32C5     EspChipModelT = 23
	CHIP_POSIX_LINUX EspChipModelT = 999
)

/**
 * @brief The structure represents information about the chip
 */

type EspChipInfoT struct {
	Model    EspChipModelT
	Features c.Uint32T
	Revision c.Uint16T
	Cores    c.Uint8T
}

/**
 * @brief Fill an esp_chip_info_t structure with information about the chip
 * @param[out] out_info structure to be filled
 */
// llgo:link (*EspChipInfoT).EspChipInfo C.esp_chip_info
func (recv_ *EspChipInfoT) EspChipInfo() {
}
