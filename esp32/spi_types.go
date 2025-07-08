package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type SpiHostDeviceT c.Int

const (
	SPI1_HOST    SpiHostDeviceT = 0
	SPI2_HOST    SpiHostDeviceT = 1
	SPI3_HOST    SpiHostDeviceT = 2
	SPI_HOST_MAX SpiHostDeviceT = 3
)

type SpiClockSourceT SocPeriphSpiClkSrcT
type SpiEventT c.Int

const (
	SPI_EV_BUF_TX         SpiEventT = 1
	SPI_EV_BUF_RX         SpiEventT = 2
	SPI_EV_SEND_DMA_READY SpiEventT = 4
	SPI_EV_SEND           SpiEventT = 8
	SPI_EV_RECV_DMA_READY SpiEventT = 16
	SPI_EV_RECV           SpiEventT = 32
	SPI_EV_CMD9           SpiEventT = 64
	SPI_EV_CMDA           SpiEventT = 128
	SPI_EV_TRANS          SpiEventT = 256
)

/**
 * @brief Line mode of SPI transaction phases: CMD, ADDR, DOUT/DIN.
 */

type SpiLineModeT struct {
	CmdLines  c.Uint8T
	AddrLines c.Uint8T
	DataLines c.Uint8T
}
type SpiCommandT c.Int

const (
	SPI_CMD_HD_WRBUF   SpiCommandT = 1
	SPI_CMD_HD_RDBUF   SpiCommandT = 2
	SPI_CMD_HD_WRDMA   SpiCommandT = 4
	SPI_CMD_HD_RDDMA   SpiCommandT = 8
	SPI_CMD_HD_SEG_END SpiCommandT = 16
	SPI_CMD_HD_EN_QPI  SpiCommandT = 32
	SPI_CMD_HD_WR_END  SpiCommandT = 64
	SPI_CMD_HD_INT0    SpiCommandT = 128
	SPI_CMD_HD_INT1    SpiCommandT = 256
	SPI_CMD_HD_INT2    SpiCommandT = 512
)

type SpiSamplingPointT c.Int

const (
	SPI_SAMPLING_POINT_PHASE_0 SpiSamplingPointT = 0
	SPI_SAMPLING_POINT_PHASE_1 SpiSamplingPointT = 1
)
