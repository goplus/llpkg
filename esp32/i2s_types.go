package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type I2sSlotModeT c.Int

const (
	I2S_SLOT_MODE_MONO   I2sSlotModeT = 1
	I2S_SLOT_MODE_STEREO I2sSlotModeT = 2
)

type I2sDirT c.Int

const (
	I2S_DIR_RX I2sDirT = 1
	I2S_DIR_TX I2sDirT = 2
)

type I2sRoleT c.Int

const (
	I2S_ROLE_MASTER I2sRoleT = 0
	I2S_ROLE_SLAVE  I2sRoleT = 1
)

type I2sDataBitWidthT c.Int

const (
	I2S_DATA_BIT_WIDTH_8BIT  I2sDataBitWidthT = 8
	I2S_DATA_BIT_WIDTH_16BIT I2sDataBitWidthT = 16
	I2S_DATA_BIT_WIDTH_24BIT I2sDataBitWidthT = 24
	I2S_DATA_BIT_WIDTH_32BIT I2sDataBitWidthT = 32
)

type I2sSlotBitWidthT c.Int

const (
	I2S_SLOT_BIT_WIDTH_AUTO  I2sSlotBitWidthT = 0
	I2S_SLOT_BIT_WIDTH_8BIT  I2sSlotBitWidthT = 8
	I2S_SLOT_BIT_WIDTH_16BIT I2sSlotBitWidthT = 16
	I2S_SLOT_BIT_WIDTH_24BIT I2sSlotBitWidthT = 24
	I2S_SLOT_BIT_WIDTH_32BIT I2sSlotBitWidthT = 32
)

type I2sClockSrcT SocPeriphI2sClkSrcT
type I2sPcmCompressT c.Int

const (
	I2S_PCM_DISABLE      I2sPcmCompressT = 0
	I2S_PCM_A_DECOMPRESS I2sPcmCompressT = 1
	I2S_PCM_A_COMPRESS   I2sPcmCompressT = 2
	I2S_PCM_U_DECOMPRESS I2sPcmCompressT = 3
	I2S_PCM_U_COMPRESS   I2sPcmCompressT = 4
)

type I2sPdmDsrT c.Int

const (
	I2S_PDM_DSR_8S  I2sPdmDsrT = 0
	I2S_PDM_DSR_16S I2sPdmDsrT = 1
	I2S_PDM_DSR_MAX I2sPdmDsrT = 2
)

type I2sPdmSigScaleT c.Int

const (
	I2S_PDM_SIG_SCALING_DIV_2 I2sPdmSigScaleT = 0
	I2S_PDM_SIG_SCALING_MUL_1 I2sPdmSigScaleT = 1
	I2S_PDM_SIG_SCALING_MUL_2 I2sPdmSigScaleT = 2
	I2S_PDM_SIG_SCALING_MUL_4 I2sPdmSigScaleT = 3
)

type I2sPdmTxLineModeT c.Int

const (
	I2S_PDM_TX_ONE_LINE_CODEC I2sPdmTxLineModeT = 0
	I2S_PDM_TX_ONE_LINE_DAC   I2sPdmTxLineModeT = 1
	I2S_PDM_TX_TWO_LINE_DAC   I2sPdmTxLineModeT = 2
)

type I2sStdSlotMaskT c.Int

const (
	I2S_STD_SLOT_LEFT  I2sStdSlotMaskT = 1
	I2S_STD_SLOT_RIGHT I2sStdSlotMaskT = 2
	I2S_STD_SLOT_BOTH  I2sStdSlotMaskT = 3
)

type I2sPdmSlotMaskT c.Int

const (
	I2S_PDM_SLOT_RIGHT          I2sPdmSlotMaskT = 1
	I2S_PDM_SLOT_LEFT           I2sPdmSlotMaskT = 2
	I2S_PDM_SLOT_BOTH           I2sPdmSlotMaskT = 3
	I2S_PDM_RX_LINE0_SLOT_RIGHT I2sPdmSlotMaskT = 1
	I2S_PDM_RX_LINE0_SLOT_LEFT  I2sPdmSlotMaskT = 2
	I2S_PDM_RX_LINE1_SLOT_RIGHT I2sPdmSlotMaskT = 4
	I2S_PDM_RX_LINE1_SLOT_LEFT  I2sPdmSlotMaskT = 8
	I2S_PDM_RX_LINE2_SLOT_RIGHT I2sPdmSlotMaskT = 16
	I2S_PDM_RX_LINE2_SLOT_LEFT  I2sPdmSlotMaskT = 32
	I2S_PDM_RX_LINE3_SLOT_RIGHT I2sPdmSlotMaskT = 64
	I2S_PDM_RX_LINE3_SLOT_LEFT  I2sPdmSlotMaskT = 128
	I2S_PDM_LINE_SLOT_ALL       I2sPdmSlotMaskT = 255
)

type I2sTdmSlotMaskT c.Int

const (
	I2S_TDM_SLOT0  I2sTdmSlotMaskT = 1
	I2S_TDM_SLOT1  I2sTdmSlotMaskT = 2
	I2S_TDM_SLOT2  I2sTdmSlotMaskT = 4
	I2S_TDM_SLOT3  I2sTdmSlotMaskT = 8
	I2S_TDM_SLOT4  I2sTdmSlotMaskT = 16
	I2S_TDM_SLOT5  I2sTdmSlotMaskT = 32
	I2S_TDM_SLOT6  I2sTdmSlotMaskT = 64
	I2S_TDM_SLOT7  I2sTdmSlotMaskT = 128
	I2S_TDM_SLOT8  I2sTdmSlotMaskT = 256
	I2S_TDM_SLOT9  I2sTdmSlotMaskT = 512
	I2S_TDM_SLOT10 I2sTdmSlotMaskT = 1024
	I2S_TDM_SLOT11 I2sTdmSlotMaskT = 2048
	I2S_TDM_SLOT12 I2sTdmSlotMaskT = 4096
	I2S_TDM_SLOT13 I2sTdmSlotMaskT = 8192
	I2S_TDM_SLOT14 I2sTdmSlotMaskT = 16384
	I2S_TDM_SLOT15 I2sTdmSlotMaskT = 32768
)

type I2sEtmEventTypeT c.Int

const (
	I2S_ETM_EVENT_DONE         I2sEtmEventTypeT = 0
	I2S_ETM_EVENT_REACH_THRESH I2sEtmEventTypeT = 1
	I2S_ETM_EVENT_MAX          I2sEtmEventTypeT = 2
)

type I2sEtmTaskTypeT c.Int

const (
	I2S_ETM_TASK_START I2sEtmTaskTypeT = 0
	I2S_ETM_TASK_STOP  I2sEtmTaskTypeT = 1
	I2S_ETM_TASK_MAX   I2sEtmTaskTypeT = 2
)
