package core

import (
	"fmt"
	_ "unsafe"

	"github.com/goplus/llgo/c"
)

// #if LV_USE_FLOAT
//type lv_value_precise_t c.Float

// #else
type LvValuePreciseT c.Int

// #endif

/**
 * LVGL error codes.
 */

type LvResultT c.Int

const (
	LV_RESULT_INVALID LvResultT = 0 /*Typically indicates that the object is deleted (become invalid) in the action
	  function or an operation was failed*/
	LV_RESULT_OK LvResultT = 1 /*The object is valid (no deleted) after the action*/
)

type LvOpaT c.Int

const (
	LV_OPA_TRANSP LvOpaT = 0
	LV_OPA_0      LvOpaT = 0
	LV_OPA_10     LvOpaT = 25
	LV_OPA_20     LvOpaT = 51
	LV_OPA_30     LvOpaT = 76
	LV_OPA_40     LvOpaT = 102
	LV_OPA_50     LvOpaT = 127
	LV_OPA_60     LvOpaT = 153
	LV_OPA_70     LvOpaT = 178
	LV_OPA_80     LvOpaT = 204
	LV_OPA_90     LvOpaT = 229
	LV_OPA_100    LvOpaT = 255
	LV_OPA_COVER  LvOpaT = 255
)

const (
	LV_OPA_MIN LvOpaT = 2
	LV_OPA_MAX LvOpaT = 253
)

type LvRadiusT c.Int32T

const (
	LV_RADIUS_CIRCLE LvRadiusT = 0x7FFF
)

//llgo:type C
type LvStyleTransitionDscT struct {
	Unused [0]byte
}

//llgo:type C
type LvAnimT struct {
	Unused [0]byte
}

//llgo:type C
type LvColorFilterDscT struct {
	Unused [0]byte
}

// **渐变描述结构 (lv_grad_dsc_t)**
const (
	LV_GRADIENT_MAX_STOPS = 2
)

// **渐变停止点 (lv_grad_stop_t)**
//
//llgo:type C
type LvGradStopT struct {
	Color LvColorT // 颜色
	Opa   LvOpaT   // 透明度
	Frac  uint8    // 停止点位置（0-255）
}

// **渐变方向枚举 (lv_grad_dir_t)**
type LvGradDirT uint8

const (
	LV_GRAD_DIR_NONE    LvGradDirT = iota // 无渐变
	LV_GRAD_DIR_VER                       // 垂直渐变
	LV_GRAD_DIR_HOR                       // 水平渐变
	LV_GRAD_DIR_LINEAR                    // 线性渐变
	LV_GRAD_DIR_RADIAL                    // 放射渐变
	LV_GRAD_DIR_CONICAL                   // 锥形渐变
)

// **渐变超出范围时的行为 (lv_grad_extend_t)**
type LvGradExtendT uint8

const (
	LV_GRAD_EXTEND_PAD     LvGradExtendT = iota // 保持最后的颜色
	LV_GRAD_EXTEND_REPEAT                       // 重复渐变模式
	LV_GRAD_EXTEND_REFLECT                      // 反射渐变模式
)

//llgo:type C
type LvGradDscT struct {
	Stops      [LV_GRADIENT_MAX_STOPS]LvGradStopT // 渐变停止点数组
	StopsCount uint8                              // 渐变点数量
	Flags      uint8                              // 低 4 位存储 dir，高 3 位存储 extend
	//State      unsafe.Pointer                     // 指针，存储额外状态（如果有的话）
}

// **设置 dir（低 4 位）**
func (d *LvGradDscT) SetDir(dir LvGradDirT) {
	fmt.Println("Setting dir:", dir)
	d.Flags = (d.Flags & 0xF0) | (uint8(dir) & 0x0F)
}

// **获取 dir（低 4 位）**
func (d *LvGradDscT) GetDir() LvGradDirT {
	return LvGradDirT(d.Flags & 0x0F)
}

// **设置 extend（高 3 位，5-7 位）**
func (d *LvGradDscT) SetExtend(extend LvGradExtendT) {
	d.Flags = (d.Flags & 0x8F) | ((uint8(extend) & 0x07) << 4)
}

// **获取 extend（高 3 位，5-7 位）**
func (d *LvGradDscT) GetExtend() LvGradExtendT {
	return LvGradExtendT((d.Flags >> 4) & 0x07)
}

//llgo:type C
type LvFontT struct {
	Unused [0]byte
}

//llgo:type C
type LvObjT struct {
	Unused [0]byte
}

//llgo:type C
type LvObjClassT struct {
	Unused [0]byte
}

//llgo:type C
type LvStateT struct {
	Unused [0]byte
}

//llgo:type C
type LvGroupT struct {
	Unused [0]byte
}

type LvCoverResT c.Int

const (
	LV_COVER_RES_COVER     LvCoverResT = 0
	LV_COVER_RES_NOT_COVER LvCoverResT = 1
	LV_COVER_RES_MASKED    LvCoverResT = 2
)

//llgo:type C
type LvIndevT struct {
	Unused [0]byte
}

//llgo:type C
type LvLayerT struct {
	Unused [0]byte
}

//llgo:type C
type LvAreaT struct {
	Unused [0]byte
}

//llgo:type C
type LvPointT struct {
	X c.Int32T
	Y c.Int32T
}

//llgo:type C
type LvHitTestInfoT struct {
	Unused [0]byte
}

//llgo:type C
type LvDrawTaskT struct {
	Unused [0]byte
}

type LvObjPointTransformFlagT c.Int

const (
	LV_OBJ_POINT_TRANSFORM_FLAG_NONE              LvObjPointTransformFlagT = 0x00
	LV_OBJ_POINT_TRANSFORM_FLAG_RECURSIVE         LvObjPointTransformFlagT = 0x01
	LV_OBJ_POINT_TRANSFORM_FLAG_INVERSE           LvObjPointTransformFlagT = 0x02
	LV_OBJ_POINT_TRANSFORM_FLAG_INVERSE_RECURSIVE LvObjPointTransformFlagT = 0x03
)

//llgo:type C
type LvMatrixT struct {
	Unused [0]byte
}

//llgo:type
type LvStyleT struct {
	// 	Unused [0]byte

	// 	#if LV_USE_ASSERT_STYLE
	//     uint32_t sentinel;
	// #endif

	ValuesAndProps *c.Void

	HasGroup c.Uint
	PropCnt  c.Uint
}

type LvStyleStateCmpT c.Int

const (
	LV_STYLE_STATE_CMP_SAME LvStyleStateCmpT = iota
	LV_STYLE_STATE_CMP_DIFF_REDRAW
	LV_STYLE_STATE_CMP_DIFF_DRAW_PAD
	LV_STYLE_STATE_CMP_DIFF_LAYOUT
)

type LvStyleSelectorT c.Uint

type LvAlignT c.Int

const (
	LV_ALIGN_DEFAULT LvAlignT = iota
	LV_ALIGN_TOP_LEFT
	LV_ALIGN_TOP_MID
	LV_ALIGN_TOP_RIGHT
	LV_ALIGN_BOTTOM_LEFT
	LV_ALIGN_BOTTOM_MID
	LV_ALIGN_BOTTOM_RIGHT
	LV_ALIGN_LEFT_MID
	LV_ALIGN_RIGHT_MID
	LV_ALIGN_CENTER
	LV_ALIGN_OUT_TOP_LEFT
	LV_ALIGN_OUT_TOP_MID
	LV_ALIGN_OUT_TOP_RIGHT
	LV_ALIGN_OUT_BOTTOM_LEFT
	LV_ALIGN_OUT_BOTTOM_MID
	LV_ALIGN_OUT_BOTTOM_RIGHT
	LV_ALIGN_OUT_LEFT_TOP
	LV_ALIGN_OUT_LEFT_MID
	LV_ALIGN_OUT_LEFT_BOTTOM
	LV_ALIGN_OUT_RIGHT_TOP
	LV_ALIGN_OUT_RIGHT_MID
	LV_ALIGN_OUT_RIGHT_BOTTOM
)

// llgo:type C
type LvColorT struct {
	Blue  c.Uint
	Green c.Uint
	Red   c.Uint
}

// lv_color16_t  LVGL（LittleVGL），lv_color16_t 就是用来存储 RGB565 颜色的结构体，
//
//llgo:type C
type LvColor16T struct {
	Unused [0]byte
}

//llgo:type C
type LvColor32T struct {
	Blue  c.Uint
	Green c.Uint
	Red   c.Uint
	Alpha c.Uint
}

//llgo:type C
type LvColorHsvT struct {
	Hue        c.Uint
	Saturation c.Uint
	Value      c.Uint
}

//llgo:type C
type LvColor16aT struct {
	Luminance c.Uint
	Alpha     c.Uint
}

type LvBorderSideT c.Int

/**
 * Selects on which sides border should be drawn
 * 'OR'ed values can be used.
 */
const (
	LV_BORDER_SIDE_NONE     LvBorderSideT = 0x00
	LV_BORDER_SIDE_BOTTOM   LvBorderSideT = 0x01
	LV_BORDER_SIDE_TOP      LvBorderSideT = 0x02
	LV_BORDER_SIDE_LEFT     LvBorderSideT = 0x04
	LV_BORDER_SIDE_RIGHT    LvBorderSideT = 0x08
	LV_BORDER_SIDE_FULL     LvBorderSideT = 0x0F
	LV_BORDER_SIDE_INTERNAL LvBorderSideT = 0x10 /**< FOR matrix-like objects (e.g. Button matrix)*/
)

type LvTextDecorT c.Int

/**
 * Some options to apply decorations on texts.
 * 'OR'ed values can be used.
 */
const (
	LV_TEXT_DECOR_NONE          LvTextDecorT = 0x00
	LV_TEXT_DECOR_UNDERLINE     LvTextDecorT = 0x01
	LV_TEXT_DECOR_STRIKETHROUGH LvTextDecorT = 0x02
)

type LvTextAlignT c.Int

/** Label align policy*/
const (
	LV_TEXT_ALIGN_AUTO   LvTextAlignT = 0x00
	LV_TEXT_ALIGN_LEFT   LvTextAlignT = 0x01
	LV_TEXT_ALIGN_CENTER LvTextAlignT = 0x02
	LV_TEXT_ALIGN_RIGHT  LvTextAlignT = 0x03
)

type LvBlendModeT c.Int

/**
 * Possible options for blending opaque drawings
 */
const (
	LV_BLEND_MODE_NORMAL      LvBlendModeT = 0x00
	LV_BLEND_MODE_ADDITIVE    LvBlendModeT = 0x01
	LV_BLEND_MODE_SUBTRACTIVE LvBlendModeT = 0x02
	LV_BLEND_MODE_MULTIPLY    LvBlendModeT = 0x03
	LV_BLEND_MODE_DIFFERENCE  LvBlendModeT = 0x04
)

type LvBaseDirT c.Int

const (
	LV_BASE_DIR_LTR  LvBaseDirT = 0x00
	LV_BASE_DIR_RTL  LvBaseDirT = 0x01
	LV_BASE_DIR_AUTO LvBaseDirT = 0x02

	LV_BASE_DIR_NEUTRAL LvBaseDirT = 0x20
	LV_BASE_DIR_WEAK    LvBaseDirT = 0x21
)

//llgo:type C
type LvStyleValueT struct {
	Unused [0]byte
}

//llgo:type C
type LvStylePropT struct {
	Unused [0]byte
}

//llgo:type C
type LvPartT struct {
	Unused [0]byte
}

//llgo:type C
type LvStyleResT struct {
	Unused [0]byte
}

//llgo:type C
type LvDrawBufT struct {
	Unused [0]byte
}

type LvColorFormatT c.Int

const (
	LV_COLOR_FORMAT_UNKNOWN LvColorFormatT = 0

	LV_COLOR_FORMAT_RAW       LvColorFormatT = 0x01
	LV_COLOR_FORMAT_RAW_ALPHA LvColorFormatT = 0x02

	/*<=1 byte (+alpha) formats*/
	LV_COLOR_FORMAT_L8 LvColorFormatT = 0x06
	LV_COLOR_FORMAT_I1 LvColorFormatT = 0x07
	LV_COLOR_FORMAT_I2 LvColorFormatT = 0x08
	LV_COLOR_FORMAT_I4 LvColorFormatT = 0x09
	LV_COLOR_FORMAT_I8 LvColorFormatT = 0x0A
	LV_COLOR_FORMAT_A8 LvColorFormatT = 0x0E

	/*2 byte (+alpha) formats*/
	LV_COLOR_FORMAT_RGB565   LvColorFormatT = 0x12
	LV_COLOR_FORMAT_ARGB8565 LvColorFormatT = 0x13
	LV_COLOR_FORMAT_RGB565A8 LvColorFormatT = 0x14
	LV_COLOR_FORMAT_AL88     LvColorFormatT = 0x15

	/*3 byte (+alpha) formats*/
	LV_COLOR_FORMAT_RGB888   LvColorFormatT = 0x0F
	LV_COLOR_FORMAT_ARGB8888 LvColorFormatT = 0x10
	LV_COLOR_FORMAT_XRGB8888 LvColorFormatT = 0x11

	/*Formats not supported by software renderer but kept here so GPU can use it*/
	LV_COLOR_FORMAT_A1       LvColorFormatT = 0x0B
	LV_COLOR_FORMAT_A2       LvColorFormatT = 0x0C
	LV_COLOR_FORMAT_A4       LvColorFormatT = 0x0D
	LV_COLOR_FORMAT_ARGB1555 LvColorFormatT = 0x16
	LV_COLOR_FORMAT_ARGB4444 LvColorFormatT = 0x17
	LV_COLOR_FORMAT_ARGB2222 LvColorFormatT = 0x18

	/* reference to https://wiki.videolan.org/YUV/ */
	/*YUV planar formats*/
	LV_COLOR_FORMAT_YUV_START LvColorFormatT = 0x20
	LV_COLOR_FORMAT_I420      LvColorFormatT = LV_COLOR_FORMAT_YUV_START /*YUV420 planar(3 plane)*/
	LV_COLOR_FORMAT_I422      LvColorFormatT = 0x21                      /*YUV422 planar(3 plane)*/
	LV_COLOR_FORMAT_I444      LvColorFormatT = 0x22                      /*YUV444 planar(3 plane)*/
	LV_COLOR_FORMAT_I400      LvColorFormatT = 0x23                      /*YUV400 no chroma channel*/
	LV_COLOR_FORMAT_NV21      LvColorFormatT = 0x24                      /*YUV420 planar(2 plane), UV plane in 'V, U, V, U'*/
	LV_COLOR_FORMAT_NV12      LvColorFormatT = 0x25                      /*YUV420 planar(2 plane), UV plane in 'U, V, U, V'*/

	/*YUV packed formats*/
	LV_COLOR_FORMAT_YUY2 LvColorFormatT = 0x26 /*YUV422 packed like 'Y U Y V'*/
	LV_COLOR_FORMAT_UYVY LvColorFormatT = 0x27 /*YUV422 packed like 'U Y V Y'*/

	LV_COLOR_FORMAT_YUV_END LvColorFormatT = LV_COLOR_FORMAT_UYVY

	LV_COLOR_FORMAT_PROPRIETARY_START LvColorFormatT = 0x30

	LV_COLOR_FORMAT_NEMA_TSC_START LvColorFormatT = LV_COLOR_FORMAT_PROPRIETARY_START
	LV_COLOR_FORMAT_NEMA_TSC4      LvColorFormatT = LV_COLOR_FORMAT_NEMA_TSC_START
	LV_COLOR_FORMAT_NEMA_TSC6      LvColorFormatT = 0x31
	LV_COLOR_FORMAT_NEMA_TSC6A     LvColorFormatT = 0x32
	LV_COLOR_FORMAT_NEMA_TSC6AP    LvColorFormatT = 0x33
	LV_COLOR_FORMAT_NEMA_TSC12     LvColorFormatT = 0x34
	LV_COLOR_FORMAT_NEMA_TSC12A    LvColorFormatT = 0x35
	LV_COLOR_FORMAT_NEMA_TSC_END   LvColorFormatT = LV_COLOR_FORMAT_NEMA_TSC12A

	//     /*Color formats in which LVGL can render*/
	// #if LV_COLOR_DEPTH == 1
	// LV_COLOR_FORMAT_NATIVE            lv_color_format_t = LV_COLOR_FORMAT_I1
	// LV_COLOR_FORMAT_NATIVE_WITH_ALPHA lv_color_format_t = LV_COLOR_FORMAT_I1
	// // #elif LV_COLOR_DEPTH == 8
	// LV_COLOR_FORMAT_NATIVE            lv_color_format_t = LV_COLOR_FORMAT_L8
	// LV_COLOR_FORMAT_NATIVE_WITH_ALPHA lv_color_format_t = LV_COLOR_FORMAT_AL88
	// // #elif LV_COLOR_DEPTH == 16
	LV_COLOR_FORMAT_NATIVE            LvColorFormatT = LV_COLOR_FORMAT_RGB565
	LV_COLOR_FORMAT_NATIVE_WITH_ALPHA LvColorFormatT = LV_COLOR_FORMAT_RGB565A8
	// #elif LV_COLOR_DEPTH == 24
	//     LV_COLOR_FORMAT_NATIVE            = LV_COLOR_FORMAT_RGB888,
	//     LV_COLOR_FORMAT_NATIVE_WITH_ALPHA = LV_COLOR_FORMAT_ARGB8888,
// // #elif LV_COLOR_DEPTH == 32
// 	LV_COLOR_FORMAT_NATIVE            lv_color_format_t = LV_COLOR_FORMAT_XRGB8888
// 	LV_COLOR_FORMAT_NATIVE_WITH_ALPHA lv_color_format_t = LV_COLOR_FORMAT_ARGB8888
// #else
// #error "LV_COLOR_DEPTH should be 1, 8, 16, 24 or 32"
// #endif

)

//llgo:type C
type LvImageDscT struct {
	Unused [0]byte
}

//llgo:type C
type LvChartSeriesT struct {
	Unused [0]byte
}

//llgo:type C
type LvChartCursorT struct {
	Unused [0]byte
}

type LvDirT c.Int

const (
	LV_DIR_NONE   LvDirT = 0x00
	LV_DIR_LEFT   LvDirT = (1 << 0)
	LV_DIR_RIGHT  LvDirT = (1 << 1)
	LV_DIR_TOP    LvDirT = (1 << 2)
	LV_DIR_BOTTOM LvDirT = (1 << 3)
	LV_DIR_HOR    LvDirT = LV_DIR_LEFT | LV_DIR_RIGHT
	LV_DIR_VER    LvDirT = LV_DIR_TOP | LV_DIR_BOTTOM
	LV_DIR_ALL    LvDirT = LV_DIR_HOR | LV_DIR_VER
)

type LvAnimEnableT c.Char

//llgo:type C
type LvPointPreciseT struct {
	X LvValuePreciseT
	Y LvValuePreciseT
}

//llgo:type C
type LvDisplayT struct {
	Unused [0]byte
}

//llgo:type C
type LvThemeT struct {
	Unused [0]byte
}

//llgo:type C
type LvTimerT struct {
	Unused [0]byte
}

type LvGridAlignT c.Int

const (
	LV_GRID_ALIGN_START LvGridAlignT = iota
	LV_GRID_ALIGN_CENTER
	LV_GRID_ALIGN_END
	LV_GRID_ALIGN_STRETCH
	LV_GRID_ALIGN_SPACE_EVENLY
	LV_GRID_ALIGN_SPACE_AROUND
	LV_GRID_ALIGN_SPACE_BETWEEN
)

const (
	LV_FLEX_COLUMN  = 1 << 0
	LV_FLEX_WRAP    = 1 << 2
	LV_FLEX_REVERSE = 1 << 3
)

type LvFlexAlignT c.Int

const (
	LV_FLEX_ALIGN_START LvFlexAlignT = iota
	LV_FLEX_ALIGN_END
	LV_FLEX_ALIGN_CENTER
	LV_FLEX_ALIGN_SPACE_EVENLY
	LV_FLEX_ALIGN_SPACE_AROUND
	LV_FLEX_ALIGN_SPACE_BETWEEN
)

type LvFlexFlowT c.Int

const (
	LV_FLEX_FLOW_ROW                 LvFlexFlowT = 0x00
	LV_FLEX_FLOW_COLUMN              LvFlexFlowT = LV_FLEX_COLUMN
	LV_FLEX_FLOW_ROW_WRAP            LvFlexFlowT = LV_FLEX_FLOW_ROW | LV_FLEX_WRAP
	LV_FLEX_FLOW_ROW_REVERSE         LvFlexFlowT = LV_FLEX_FLOW_ROW | LV_FLEX_REVERSE
	LV_FLEX_FLOW_ROW_WRAP_REVERSE    LvFlexFlowT = LV_FLEX_FLOW_ROW | LV_FLEX_WRAP | LV_FLEX_REVERSE
	LV_FLEX_FLOW_COLUMN_WRAP         LvFlexFlowT = LV_FLEX_FLOW_COLUMN | LV_FLEX_WRAP
	LV_FLEX_FLOW_COLUMN_REVERSE      LvFlexFlowT = LV_FLEX_FLOW_COLUMN | LV_FLEX_REVERSE
	LV_FLEX_FLOW_COLUMN_WRAP_REVERSE LvFlexFlowT = LV_FLEX_FLOW_COLUMN | LV_FLEX_WRAP | LV_FLEX_REVERSE
)
