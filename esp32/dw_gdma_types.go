package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type DwGdmaBlockTransferTypeT c.Int

const (
	DW_GDMA_BLOCK_TRANSFER_CONTIGUOUS DwGdmaBlockTransferTypeT = 0
	DW_GDMA_BLOCK_TRANSFER_RELOAD     DwGdmaBlockTransferTypeT = 1
	DW_GDMA_BLOCK_TRANSFER_SHADOW     DwGdmaBlockTransferTypeT = 2
	DW_GDMA_BLOCK_TRANSFER_LIST       DwGdmaBlockTransferTypeT = 3
)

type DwGdmaFlowControllerT c.Int

const (
	DW_GDMA_FLOW_CTRL_SELF DwGdmaFlowControllerT = 0
	DW_GDMA_FLOW_CTRL_SRC  DwGdmaFlowControllerT = 1
	DW_GDMA_FLOW_CTRL_DST  DwGdmaFlowControllerT = 2
)

type DwGdmaHandshakeTypeT c.Int

const (
	DW_GDMA_HANDSHAKE_HW DwGdmaHandshakeTypeT = 0
	DW_GDMA_HANDSHAKE_SW DwGdmaHandshakeTypeT = 1
)

type DwGdmaRoleT c.Int

const (
	DW_GDMA_ROLE_MEM        DwGdmaRoleT = 0
	DW_GDMA_ROLE_PERIPH_DSI DwGdmaRoleT = 1
	DW_GDMA_ROLE_PERIPH_CSI DwGdmaRoleT = 2
	DW_GDMA_ROLE_PERIPH_ISP DwGdmaRoleT = 3
)

type DwGdmaLockLevelT c.Int

const (
	DW_GDMA_LOCK_LEVEL_FULL_TRANS  DwGdmaLockLevelT = 0
	DW_GDMA_LOCK_LEVEL_BLOCK_TRANS DwGdmaLockLevelT = 1
)

type DwGdmaTransferWidthT c.Int

const (
	DW_GDMA_TRANS_WIDTH_8   DwGdmaTransferWidthT = 0
	DW_GDMA_TRANS_WIDTH_16  DwGdmaTransferWidthT = 1
	DW_GDMA_TRANS_WIDTH_32  DwGdmaTransferWidthT = 2
	DW_GDMA_TRANS_WIDTH_64  DwGdmaTransferWidthT = 3
	DW_GDMA_TRANS_WIDTH_128 DwGdmaTransferWidthT = 4
	DW_GDMA_TRANS_WIDTH_256 DwGdmaTransferWidthT = 5
	DW_GDMA_TRANS_WIDTH_512 DwGdmaTransferWidthT = 6
)

type DwGdmaBurstModeT c.Int

const (
	DW_GDMA_BURST_MODE_INCREMENT DwGdmaBurstModeT = 0
	DW_GDMA_BURST_MODE_FIXED     DwGdmaBurstModeT = 1
)

type DwGdmaBurstItemsT c.Int

const (
	DW_GDMA_BURST_ITEMS_1    DwGdmaBurstItemsT = 0
	DW_GDMA_BURST_ITEMS_4    DwGdmaBurstItemsT = 1
	DW_GDMA_BURST_ITEMS_8    DwGdmaBurstItemsT = 2
	DW_GDMA_BURST_ITEMS_16   DwGdmaBurstItemsT = 3
	DW_GDMA_BURST_ITEMS_32   DwGdmaBurstItemsT = 4
	DW_GDMA_BURST_ITEMS_64   DwGdmaBurstItemsT = 5
	DW_GDMA_BURST_ITEMS_128  DwGdmaBurstItemsT = 6
	DW_GDMA_BURST_ITEMS_256  DwGdmaBurstItemsT = 7
	DW_GDMA_BURST_ITEMS_512  DwGdmaBurstItemsT = 8
	DW_GDMA_BURST_ITEMS_1024 DwGdmaBurstItemsT = 9
)
