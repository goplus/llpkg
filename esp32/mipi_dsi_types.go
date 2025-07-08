package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type MipiDsiDataTypeT c.Int

const (
	MIPI_DSI_DT_VSYNC_START                 MipiDsiDataTypeT = 1
	MIPI_DSI_DT_VSYNC_END                   MipiDsiDataTypeT = 17
	MIPI_DSI_DT_HSYNC_START                 MipiDsiDataTypeT = 33
	MIPI_DSI_DT_HSYNC_END                   MipiDsiDataTypeT = 49
	MIPI_DSI_DT_EOT_PACKET                  MipiDsiDataTypeT = 8
	MIPI_DSI_DT_COLOR_MODE_OFF              MipiDsiDataTypeT = 2
	MIPI_DSI_DT_COLOR_MODE_ON               MipiDsiDataTypeT = 18
	MIPI_DSI_DT_SHUTDOWN_PERIPHERAL         MipiDsiDataTypeT = 34
	MIPI_DSI_DT_TURN_ON_PERIPHERAL          MipiDsiDataTypeT = 50
	MIPI_DSI_DT_GENERIC_SHORT_WRITE_0       MipiDsiDataTypeT = 3
	MIPI_DSI_DT_GENERIC_SHORT_WRITE_1       MipiDsiDataTypeT = 19
	MIPI_DSI_DT_GENERIC_SHORT_WRITE_2       MipiDsiDataTypeT = 35
	MIPI_DSI_DT_GENERIC_READ_REQUEST_0      MipiDsiDataTypeT = 4
	MIPI_DSI_DT_GENERIC_READ_REQUEST_1      MipiDsiDataTypeT = 20
	MIPI_DSI_DT_GENERIC_READ_REQUEST_2      MipiDsiDataTypeT = 36
	MIPI_DSI_DT_DCS_SHORT_WRITE_0           MipiDsiDataTypeT = 5
	MIPI_DSI_DT_DCS_SHORT_WRITE_1           MipiDsiDataTypeT = 21
	MIPI_DSI_DT_DCS_READ_0                  MipiDsiDataTypeT = 6
	MIPI_DSI_DT_SET_MAXIMUM_RETURN_PKT      MipiDsiDataTypeT = 55
	MIPI_DSI_DT_NULL_PACKET                 MipiDsiDataTypeT = 9
	MIPI_DSI_DT_BLANKING_PACKET             MipiDsiDataTypeT = 25
	MIPI_DSI_DT_GENERIC_LONG_WRITE          MipiDsiDataTypeT = 41
	MIPI_DSI_DT_DCS_LONG_WRITE              MipiDsiDataTypeT = 57
	MIPI_DSI_DT_PACKED_PIXEL_STREAM_RGB_16  MipiDsiDataTypeT = 14
	MIPI_DSI_DT_PACKED_PIXEL_STREAM_RGB_18  MipiDsiDataTypeT = 30
	MIPI_DSI_DT_LOOSELY_PIXEL_STREAM_RGB_18 MipiDsiDataTypeT = 46
	MIPI_DSI_DT_PACKED_PIXEL_STREAM_RGB_24  MipiDsiDataTypeT = 62
)

type MipiDsiPatternTypeT c.Int

const (
	MIPI_DSI_PATTERN_NONE           MipiDsiPatternTypeT = 0
	MIPI_DSI_PATTERN_BAR_VERTICAL   MipiDsiPatternTypeT = 1
	MIPI_DSI_PATTERN_BAR_HORIZONTAL MipiDsiPatternTypeT = 2
	MIPI_DSI_PATTERN_BER_VERTICAL   MipiDsiPatternTypeT = 3
)

type MipiDsiPhyClockSourceT c.Int
type MipiDsiDpiClockSourceT c.Int
