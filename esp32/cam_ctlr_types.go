package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type CamCtlrColorT c.Int

const (
	CAM_CTLR_COLOR_RAW8   CamCtlrColorT = 16777216
	CAM_CTLR_COLOR_RAW10  CamCtlrColorT = 16777217
	CAM_CTLR_COLOR_RAW12  CamCtlrColorT = 16777218
	CAM_CTLR_COLOR_RGB565 CamCtlrColorT = 33554434
	CAM_CTLR_COLOR_RGB666 CamCtlrColorT = 33554433
	CAM_CTLR_COLOR_RGB888 CamCtlrColorT = 33554432
	CAM_CTLR_COLOR_YUV420 CamCtlrColorT = 50331650
	CAM_CTLR_COLOR_YUV422 CamCtlrColorT = 50331649
	CAM_CTLR_COLOR_GRAY4  CamCtlrColorT = 67108864
	CAM_CTLR_COLOR_GRAY8  CamCtlrColorT = 67108865
)

type CamCtlrDataWidthT c.Int

const (
	CAM_CTLR_DATA_WIDTH_8  CamCtlrDataWidthT = 8
	CAM_CTLR_DATA_WIDTH_10 CamCtlrDataWidthT = 10
	CAM_CTLR_DATA_WIDTH_12 CamCtlrDataWidthT = 12
	CAM_CTLR_DATA_WIDTH_16 CamCtlrDataWidthT = 16
)
