package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const ESP_AES_ENCRYPT = 1
const ESP_AES_DECRYPT = 0

type EspAesModeT c.Int

const (
	ESP_AES_BLOCK_MODE_ECB    EspAesModeT = 0
	ESP_AES_BLOCK_MODE_CBC    EspAesModeT = 1
	ESP_AES_BLOCK_MODE_OFB    EspAesModeT = 2
	ESP_AES_BLOCK_MODE_CTR    EspAesModeT = 3
	ESP_AES_BLOCK_MODE_CFB8   EspAesModeT = 4
	ESP_AES_BLOCK_MODE_CFB128 EspAesModeT = 5
	ESP_AES_BLOCK_MODE_GCM    EspAesModeT = 6
	ESP_AES_BLOCK_MODE_MAX    EspAesModeT = 7
)
