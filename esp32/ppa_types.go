package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type PpaEngineTypeT c.Int

const (
	PPA_ENGINE_TYPE_SRM   PpaEngineTypeT = 0
	PPA_ENGINE_TYPE_BLEND PpaEngineTypeT = 1
)

type PpaSrmRotationAngleT c.Int

const (
	PPA_SRM_ROTATION_ANGLE_0   PpaSrmRotationAngleT = 0
	PPA_SRM_ROTATION_ANGLE_90  PpaSrmRotationAngleT = 1
	PPA_SRM_ROTATION_ANGLE_180 PpaSrmRotationAngleT = 2
	PPA_SRM_ROTATION_ANGLE_270 PpaSrmRotationAngleT = 3
)

type PpaSrmColorModeT c.Int

const (
	PPA_SRM_COLOR_MODE_ARGB8888 PpaSrmColorModeT = 83886080
	PPA_SRM_COLOR_MODE_RGB888   PpaSrmColorModeT = 33554432
	PPA_SRM_COLOR_MODE_RGB565   PpaSrmColorModeT = 33554434
	PPA_SRM_COLOR_MODE_YUV420   PpaSrmColorModeT = 50331650
	PPA_SRM_COLOR_MODE_YUV444   PpaSrmColorModeT = 50331648
)

type PpaBlendColorModeT c.Int

const (
	PPA_BLEND_COLOR_MODE_ARGB8888 PpaBlendColorModeT = 83886080
	PPA_BLEND_COLOR_MODE_RGB888   PpaBlendColorModeT = 33554432
	PPA_BLEND_COLOR_MODE_RGB565   PpaBlendColorModeT = 33554434
	PPA_BLEND_COLOR_MODE_A8       PpaBlendColorModeT = 100663297
	PPA_BLEND_COLOR_MODE_A4       PpaBlendColorModeT = 100663296
)

type PpaFillColorModeT c.Int

const (
	PPA_FILL_COLOR_MODE_ARGB8888 PpaFillColorModeT = 83886080
	PPA_FILL_COLOR_MODE_RGB888   PpaFillColorModeT = 33554432
	PPA_FILL_COLOR_MODE_RGB565   PpaFillColorModeT = 33554434
)

type PpaAlphaUpdateModeT c.Int

const (
	PPA_ALPHA_NO_CHANGE PpaAlphaUpdateModeT = 0
	PPA_ALPHA_FIX_VALUE PpaAlphaUpdateModeT = 1
	PPA_ALPHA_SCALE     PpaAlphaUpdateModeT = 2
	PPA_ALPHA_INVERT    PpaAlphaUpdateModeT = 3
)

type PpaColorConvStdRgbYuvT c.Int

const (
	PPA_COLOR_CONV_STD_RGB_YUV_BT601 PpaColorConvStdRgbYuvT = 0
	PPA_COLOR_CONV_STD_RGB_YUV_BT709 PpaColorConvStdRgbYuvT = 1
)

type PpaColorRangeT c.Int

const (
	PPA_COLOR_RANGE_LIMIT PpaColorRangeT = 0
	PPA_COLOR_RANGE_FULL  PpaColorRangeT = 1
)

type PpaDataBurstLengthT c.Int

const (
	PPA_DATA_BURST_LENGTH_8   PpaDataBurstLengthT = 1
	PPA_DATA_BURST_LENGTH_16  PpaDataBurstLengthT = 2
	PPA_DATA_BURST_LENGTH_32  PpaDataBurstLengthT = 3
	PPA_DATA_BURST_LENGTH_64  PpaDataBurstLengthT = 4
	PPA_DATA_BURST_LENGTH_128 PpaDataBurstLengthT = 5
)
