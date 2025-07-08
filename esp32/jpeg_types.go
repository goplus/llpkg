package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type JpegCodecModeT c.Int

const (
	JPEG_CODEC_ENCODER JpegCodecModeT = 0
	JPEG_CODEC_DECODER JpegCodecModeT = 1
)

/**
 * @brief Structure for recording factor of component.
 */

type JpegComponentFactorT struct {
	Horizontal c.Uint32T
	Vertical   c.Uint32T
}
type JpegSampleModeT c.Int

const (
	JPEG_SAMPLE_MODE_YUV444 JpegSampleModeT = 50331648
	JPEG_SAMPLE_MODE_YUV422 JpegSampleModeT = 50331649
	JPEG_SAMPLE_MODE_YUV420 JpegSampleModeT = 50331650
)

/**
 * @brief Structure for huffman information
 */

type JpegHuffmanTableInfoT struct {
	Info c.Uint8T
}
type JpegDownSamplingTypeT c.Int

const (
	JPEG_DOWN_SAMPLING_YUV444 JpegDownSamplingTypeT = 50331648
	JPEG_DOWN_SAMPLING_YUV422 JpegDownSamplingTypeT = 50331649
	JPEG_DOWN_SAMPLING_YUV420 JpegDownSamplingTypeT = 50331650
	JPEG_DOWN_SAMPLING_GRAY   JpegDownSamplingTypeT = 67108865
)

type JpegEncSrcTypeT c.Int

const (
	JPEG_ENC_SRC_RGB888 JpegEncSrcTypeT = 33554432
	JPEG_ENC_SRC_YUV422 JpegEncSrcTypeT = 50331649
	JPEG_ENC_SRC_RGB565 JpegEncSrcTypeT = 33554434
	JPEG_ENC_SRC_GRAY   JpegEncSrcTypeT = 67108865
)
