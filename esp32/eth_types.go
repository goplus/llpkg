package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type EthDataInterfaceT c.Int

const (
	EMAC_DATA_INTERFACE_RMII EthDataInterfaceT = 0
	EMAC_DATA_INTERFACE_MII  EthDataInterfaceT = 1
)

type EthLinkT c.Int

const (
	ETH_LINK_UP   EthLinkT = 0
	ETH_LINK_DOWN EthLinkT = 1
)

type EthSpeedT c.Int

const (
	ETH_SPEED_10M  EthSpeedT = 0
	ETH_SPEED_100M EthSpeedT = 1
	ETH_SPEED_MAX  EthSpeedT = 2
)

type EthDuplexT c.Int

const (
	ETH_DUPLEX_HALF EthDuplexT = 0
	ETH_DUPLEX_FULL EthDuplexT = 1
)

type EthChecksumT c.Int

const (
	ETH_CHECKSUM_SW EthChecksumT = 0
	ETH_CHECKSUM_HW EthChecksumT = 1
)

type EthMacDmaBurstLenT c.Int

const (
	ETH_DMA_BURST_LEN_32 EthMacDmaBurstLenT = 0
	ETH_DMA_BURST_LEN_16 EthMacDmaBurstLenT = 1
	ETH_DMA_BURST_LEN_8  EthMacDmaBurstLenT = 2
	ETH_DMA_BURST_LEN_4  EthMacDmaBurstLenT = 3
	ETH_DMA_BURST_LEN_2  EthMacDmaBurstLenT = 4
	ETH_DMA_BURST_LEN_1  EthMacDmaBurstLenT = 5
)
