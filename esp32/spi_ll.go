package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const SPI_LL_MOSI_FREE_LEVEL = 1

type SpiLlClockValT c.Uint32T
type SpiDmaDevT SpiDevT
type SpiLlIntrT c.Int

const (
	SPI_LL_INTR_TRANS_DONE SpiLlIntrT = 1
	SPI_LL_INTR_RDBUF      SpiLlIntrT = 64
	SPI_LL_INTR_WRBUF      SpiLlIntrT = 128
	SPI_LL_INTR_RDDMA      SpiLlIntrT = 256
	SPI_LL_INTR_WRDMA      SpiLlIntrT = 512
	SPI_LL_INTR_CMD7       SpiLlIntrT = 1024
	SPI_LL_INTR_CMD8       SpiLlIntrT = 2048
	SPI_LL_INTR_CMD9       SpiLlIntrT = 4096
	SPI_LL_INTR_CMDA       SpiLlIntrT = 8192
	SPI_LL_INTR_SEG_DONE   SpiLlIntrT = 16384
)

type SpiLlTransLenCondT c.Int

const (
	SPI_LL_TRANS_LEN_COND_WRBUF SpiLlTransLenCondT = 1
	SPI_LL_TRANS_LEN_COND_RDBUF SpiLlTransLenCondT = 2
	SPI_LL_TRANS_LEN_COND_WRDMA SpiLlTransLenCondT = 4
	SPI_LL_TRANS_LEN_COND_RDDMA SpiLlTransLenCondT = 8
)

type SpiLlBaseCommandT c.Int

const (
	SPI_LL_BASE_CMD_HD_WRBUF   SpiLlBaseCommandT = 1
	SPI_LL_BASE_CMD_HD_RDBUF   SpiLlBaseCommandT = 2
	SPI_LL_BASE_CMD_HD_WRDMA   SpiLlBaseCommandT = 3
	SPI_LL_BASE_CMD_HD_RDDMA   SpiLlBaseCommandT = 4
	SPI_LL_BASE_CMD_HD_SEG_END SpiLlBaseCommandT = 5
	SPI_LL_BASE_CMD_HD_EN_QPI  SpiLlBaseCommandT = 6
	SPI_LL_BASE_CMD_HD_WR_END  SpiLlBaseCommandT = 7
	SPI_LL_BASE_CMD_HD_INT0    SpiLlBaseCommandT = 8
	SPI_LL_BASE_CMD_HD_INT1    SpiLlBaseCommandT = 9
	SPI_LL_BASE_CMD_HD_INT2    SpiLlBaseCommandT = 10
)
