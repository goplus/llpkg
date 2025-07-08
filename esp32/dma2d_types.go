package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type Dma2dDescriptorAlign8S struct {
	Buffer c.Pointer
	Next   *Dma2dDescriptorAlign8T
}
type Dma2dDescriptorAlign8T Dma2dDescriptorAlign8S
type Dma2dDescriptorT Dma2dDescriptorAlign8T
type Dma2dTriggerPeripheralT c.Int

const (
	DMA2D_TRIG_PERIPH_M2M          Dma2dTriggerPeripheralT = 0
	DMA2D_TRIG_PERIPH_JPEG_ENCODER Dma2dTriggerPeripheralT = 1
	DMA2D_TRIG_PERIPH_JPEG_DECODER Dma2dTriggerPeripheralT = 2
	DMA2D_TRIG_PERIPH_PPA_SRM      Dma2dTriggerPeripheralT = 3
	DMA2D_TRIG_PERIPH_PPA_BLEND    Dma2dTriggerPeripheralT = 4
)

type Dma2dChannelDirectionT c.Int

const (
	DMA2D_CHANNEL_DIRECTION_TX Dma2dChannelDirectionT = 0
	DMA2D_CHANNEL_DIRECTION_RX Dma2dChannelDirectionT = 1
)

type Dma2dDataBurstLengthT c.Int

const (
	DMA2D_DATA_BURST_LENGTH_8       Dma2dDataBurstLengthT = 1
	DMA2D_DATA_BURST_LENGTH_16      Dma2dDataBurstLengthT = 2
	DMA2D_DATA_BURST_LENGTH_32      Dma2dDataBurstLengthT = 3
	DMA2D_DATA_BURST_LENGTH_64      Dma2dDataBurstLengthT = 4
	DMA2D_DATA_BURST_LENGTH_128     Dma2dDataBurstLengthT = 5
	DMA2D_DATA_BURST_LENGTH_INVALID Dma2dDataBurstLengthT = 6
)

type Dma2dMacroBlockSizeT c.Int

const (
	DMA2D_MACRO_BLOCK_SIZE_NONE    Dma2dMacroBlockSizeT = 0
	DMA2D_MACRO_BLOCK_SIZE_8_8     Dma2dMacroBlockSizeT = 1
	DMA2D_MACRO_BLOCK_SIZE_8_16    Dma2dMacroBlockSizeT = 2
	DMA2D_MACRO_BLOCK_SIZE_16_16   Dma2dMacroBlockSizeT = 3
	DMA2D_MACRO_BLOCK_SIZE_INVALID Dma2dMacroBlockSizeT = 4
)

type Dma2dScrambleOrderT c.Int

const (
	DMA2D_SCRAMBLE_ORDER_BYTE2_1_0 Dma2dScrambleOrderT = 0
	DMA2D_SCRAMBLE_ORDER_BYTE2_0_1 Dma2dScrambleOrderT = 1
	DMA2D_SCRAMBLE_ORDER_BYTE1_0_2 Dma2dScrambleOrderT = 2
	DMA2D_SCRAMBLE_ORDER_BYTE1_2_0 Dma2dScrambleOrderT = 3
	DMA2D_SCRAMBLE_ORDER_BYTE0_2_1 Dma2dScrambleOrderT = 4
	DMA2D_SCRAMBLE_ORDER_BYTE0_1_2 Dma2dScrambleOrderT = 5
	DMA2D_SCRAMBLE_ORDER_INVALID   Dma2dScrambleOrderT = 6
)

type Dma2dCscTxOptionT c.Int

const (
	DMA2D_CSC_TX_NONE                 Dma2dCscTxOptionT = 0
	DMA2D_CSC_TX_SCRAMBLE             Dma2dCscTxOptionT = 1
	DMA2D_CSC_TX_RGB888_TO_RGB565     Dma2dCscTxOptionT = 2
	DMA2D_CSC_TX_RGB565_TO_RGB888     Dma2dCscTxOptionT = 3
	DMA2D_CSC_TX_RGB888_TO_YUV444_601 Dma2dCscTxOptionT = 4
	DMA2D_CSC_TX_RGB888_TO_YUV444_709 Dma2dCscTxOptionT = 5
	DMA2D_CSC_TX_RGB888_TO_YUV422_601 Dma2dCscTxOptionT = 6
	DMA2D_CSC_TX_RGB888_TO_YUV422_709 Dma2dCscTxOptionT = 7
	DMA2D_CSC_TX_YUV444_TO_RGB888_601 Dma2dCscTxOptionT = 8
	DMA2D_CSC_TX_YUV444_TO_RGB888_709 Dma2dCscTxOptionT = 9
	DMA2D_CSC_TX_YUV422_TO_RGB888_601 Dma2dCscTxOptionT = 10
	DMA2D_CSC_TX_YUV422_TO_RGB888_709 Dma2dCscTxOptionT = 11
	DMA2D_CSC_TX_INVALID              Dma2dCscTxOptionT = 12
)

type Dma2dCscRxOptionT c.Int

const (
	DMA2D_CSC_RX_NONE                 Dma2dCscRxOptionT = 0
	DMA2D_CSC_RX_SCRAMBLE             Dma2dCscRxOptionT = 1
	DMA2D_CSC_RX_YUV422_TO_YUV444     Dma2dCscRxOptionT = 2
	DMA2D_CSC_RX_YUV420_TO_YUV444     Dma2dCscRxOptionT = 3
	DMA2D_CSC_RX_YUV420_TO_RGB888_601 Dma2dCscRxOptionT = 4
	DMA2D_CSC_RX_YUV420_TO_RGB565_601 Dma2dCscRxOptionT = 5
	DMA2D_CSC_RX_YUV420_TO_RGB888_709 Dma2dCscRxOptionT = 6
	DMA2D_CSC_RX_YUV420_TO_RGB565_709 Dma2dCscRxOptionT = 7
	DMA2D_CSC_RX_YUV422_TO_RGB888_601 Dma2dCscRxOptionT = 8
	DMA2D_CSC_RX_YUV422_TO_RGB565_601 Dma2dCscRxOptionT = 9
	DMA2D_CSC_RX_YUV422_TO_RGB888_709 Dma2dCscRxOptionT = 10
	DMA2D_CSC_RX_YUV422_TO_RGB565_709 Dma2dCscRxOptionT = 11
	DMA2D_CSC_RX_YUV444_TO_RGB888_601 Dma2dCscRxOptionT = 12
	DMA2D_CSC_RX_YUV444_TO_RGB565_601 Dma2dCscRxOptionT = 13
	DMA2D_CSC_RX_YUV444_TO_RGB888_709 Dma2dCscRxOptionT = 14
	DMA2D_CSC_RX_YUV444_TO_RGB565_709 Dma2dCscRxOptionT = 15
	DMA2D_CSC_RX_INVALID              Dma2dCscRxOptionT = 16
)
