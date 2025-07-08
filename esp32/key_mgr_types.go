package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const KEY_MGR_ASSIST_INFO_LEN = 64
const HUK_INFO_LEN = 384
const KEY_INFO_LEN = 64
const KEY_HUK_SECTOR_MAGIC = 0xDEA5CE5A

type EspKeyMgrStateT c.Int

const (
	ESP_KEY_MGR_STATE_IDLE EspKeyMgrStateT = 0
	ESP_KEY_MGR_STATE_LOAD EspKeyMgrStateT = 1
	ESP_KEY_MGR_STATE_GAIN EspKeyMgrStateT = 2
	ESP_KEY_MGR_STATE_BUSY EspKeyMgrStateT = 3
)

type EspKeyMgrXtsAesKeyLenT c.Int

const (
	ESP_KEY_MGR_XTS_AES_LEN_256 EspKeyMgrXtsAesKeyLenT = 0
	ESP_KEY_MGR_XTS_AES_LEN_512 EspKeyMgrXtsAesKeyLenT = 1
)

type EspKeyMgrKeyTypeT c.Int

const (
	ESP_KEY_MGR_ECDSA_KEY       EspKeyMgrKeyTypeT = 0
	ESP_KEY_MGR_XTS_AES_128_KEY EspKeyMgrKeyTypeT = 1
	ESP_KEY_MGR_XTS_AES_256_KEY EspKeyMgrKeyTypeT = 2
)

type EspKeyMgrKeyUsageT c.Int

const (
	ESP_KEY_MGR_USE_OWN_KEY   EspKeyMgrKeyUsageT = 0
	ESP_KEY_MGR_USE_EFUSE_KEY EspKeyMgrKeyUsageT = 1
	ESP_KEY_MGR_USAGE_INVALID EspKeyMgrKeyUsageT = 2
)

type EspKeyMgrKeyPurposeT c.Int

const (
	ESP_KEY_MGR_KEY_PURPOSE_ECDSA         EspKeyMgrKeyPurposeT = 1
	ESP_KEY_MGR_KEY_PURPOSE_XTS_AES_256_1 EspKeyMgrKeyPurposeT = 2
	ESP_KEY_MGR_KEY_PURPOSE_XTS_AES_256_2 EspKeyMgrKeyPurposeT = 3
	ESP_KEY_MGR_KEY_PURPOSE_XTS_AES_128   EspKeyMgrKeyPurposeT = 4
)

type EspKeyMgrKeyGeneratorModeT c.Int

const (
	ESP_KEY_MGR_KEYGEN_MODE_RANDOM  EspKeyMgrKeyGeneratorModeT = 0
	ESP_KEY_MGR_KEYGEN_MODE_AES     EspKeyMgrKeyGeneratorModeT = 1
	ESP_KEY_MGR_KEYGEN_MODE_ECDH0   EspKeyMgrKeyGeneratorModeT = 2
	ESP_KEY_MGR_KEYGEN_MODE_ECDH1   EspKeyMgrKeyGeneratorModeT = 3
	ESP_KEY_MGR_KEYGEN_MODE_RECOVER EspKeyMgrKeyGeneratorModeT = 4
	ESP_KEY_MGR_KEYGEN_MODE_EXPORT  EspKeyMgrKeyGeneratorModeT = 5
)

type EspKeyMgrInterruptTypeT c.Int

const (
	ESP_KEY_MGR_INT_PREP_DONE EspKeyMgrInterruptTypeT = 1
	ESP_KEY_MGR_INT_PROC_DONE EspKeyMgrInterruptTypeT = 2
	ESP_KEY_MGR_INT_POST_DONE EspKeyMgrInterruptTypeT = 3
)

// store huk info, occupy 96 words
type EspKeyMgrHukInfoT struct {
	Info [384]c.Uint8T
	Crc  c.Uint32T
}

// store key info, occupy 512 bits
type EspKeyMgrKeyInfoT struct {
	Info [64]c.Uint8T
	Crc  c.Uint32T
}

type EspKeyMgrKeyRecoveryInfoT struct {
	Magic    c.Uint32T
	Version  c.Uint32T
	KeyType  c.Uint8T
	Reserved [15]c.Uint8T
	HukInfo  EspKeyMgrHukInfoT
	KeyInfo  [2]EspKeyMgrKeyInfoT
}
