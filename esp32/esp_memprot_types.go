package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const MEMPROT_OP_NONE = 0x00000000
const MEMPROT_OP_READ = 0x00000001
const MEMPROT_OP_WRITE = 0x00000002
const MEMPROT_OP_EXEC = 0x00000004
const MEMPROT_OP_INVALID = 0x80000000

type EspMprotPmsWorldT c.Int

const (
	MEMPROT_PMS_WORLD_NONE    EspMprotPmsWorldT = 0
	MEMPROT_PMS_WORLD_0       EspMprotPmsWorldT = 1
	MEMPROT_PMS_WORLD_1       EspMprotPmsWorldT = 2
	MEMPROT_PMS_WORLD_2       EspMprotPmsWorldT = 4
	MEMPROT_PMS_WORLD_ALL     EspMprotPmsWorldT = 2147483647
	MEMPROT_PMS_WORLD_INVALID EspMprotPmsWorldT = -2147483648
)
