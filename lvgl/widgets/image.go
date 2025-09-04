package widgets

import (
	"unsafe"
	_ "unsafe"

	"github.com/goplus/llgo/c"
	core "github.com/goplus/llpkg/lvgl/core"
)

// Image alignment mode
type LvImageAlignT uint8

const (
	LV_IMAGE_ALIGN_DEFAULT LvImageAlignT = iota
	LV_IMAGE_ALIGN_TOP_LEFT
	LV_IMAGE_ALIGN_TOP_MID
	LV_IMAGE_ALIGN_TOP_RIGHT
	LV_IMAGE_ALIGN_BOTTOM_LEFT
	LV_IMAGE_ALIGN_BOTTOM_MID
	LV_IMAGE_ALIGN_BOTTOM_RIGHT
	LV_IMAGE_ALIGN_LEFT_MID
	LV_IMAGE_ALIGN_RIGHT_MID
	LV_IMAGE_ALIGN_CENTER
	LV_IMAGE_ALIGN_AUTO_TRANSFORM
	LV_IMAGE_ALIGN_STRETCH
	LV_IMAGE_ALIGN_TILE
)

//go:linkname LvImageCreate C.lv_image_create
func LvImageCreate(parent *core.LvObjT) *core.LvObjT

//go:linkname LvImageSetSrc C.lv_image_set_src
func LvImageSetSrc(obj *core.LvObjT, src *c.Void)

//go:linkname LvImageSetOffsetX C.lv_image_set_offset_x
func LvImageSetOffsetX(obj *core.LvObjT, x c.Int32T)

//go:linkname LvImageSetOffsetY C.lv_image_set_offset_y
func LvImageSetOffsetY(obj *core.LvObjT, y c.Int32T)

//go:linkname LvImageSetRotation C.lv_image_set_rotation
func LvImageSetRotation(obj *core.LvObjT, angle c.Int32T)

//go:linkname LvImageSetPivot C.lv_image_set_pivot
func LvImageSetPivot(obj *core.LvObjT, x c.Int32T, y c.Int32T)

//go:linkname LvImageSetScale C.lv_image_set_scale
func LvImageSetScale(obj *core.LvObjT, zoom c.Uint32T)

//go:linkname LvImageSetScaleX C.lv_image_set_scale_x
func LvImageSetScaleX(obj *core.LvObjT, zoom c.Uint32T)

//go:linkname LvImageSetScaleY C.lv_image_set_scale_y
func LvImageSetScaleY(obj *core.LvObjT, zoom c.Uint32T)

//go:linkname LvImageSetBlendMode C.lv_image_set_blend_mode
func LvImageSetBlendMode(obj *core.LvObjT, blend_mode core.LvBlendModeT)

//go:linkname LvImageSetAntialias C.lv_image_set_antialias
func LvImageSetAntialias(obj *core.LvObjT, antialias bool)

//go:linkname LvImageSetInnerAlign C.lv_image_set_inner_align
func LvImageSetInnerAlign(obj *core.LvObjT, align LvImageAlignT)

//go:linkname LvImageSetBitmapMapSrc C.lv_image_set_bitmap_map_src
func LvImageSetBitmapMapSrc(obj *core.LvObjT, src *core.LvImageDscT)

//go:linkname LvImageGetSrc C.lv_image_get_src
func LvImageGetSrc(obj *core.LvObjT) unsafe.Pointer

//go:linkname LvImageGetOffsetX C.lv_image_get_offset_x
func LvImageGetOffsetX(obj *core.LvObjT) c.Int32T

//go:linkname LvImageGetOffsetY C.lv_image_get_offset_y
func LvImageGetOffsetY(obj *core.LvObjT) c.Int32T

//go:linkname LvImageGetRotation C.lv_image_get_rotation
func LvImageGetRotation(obj *core.LvObjT) c.Int32T

//go:linkname LvImageGetPivot C.lv_image_get_pivot
func LvImageGetPivot(obj *core.LvObjT, pivot *core.LvPointT)

//go:linkname LvImageGetScale C.lv_image_get_scale
func LvImageGetScale(obj *core.LvObjT) c.Int32T

//go:linkname LvImageGetScaleX C.lv_image_get_scale_x
func LvImageGetScaleX(obj *core.LvObjT) c.Int32T

//go:linkname LvImageGetScaleY C.lv_image_get_scale_y
func LvImageGetScaleY(obj *core.LvObjT) c.Int32T

//go:linkname LvImageGetBlendMode C.lv_image_get_blend_mode
func LvImageGetBlendMode(obj *core.LvObjT) core.LvBlendModeT

//go:linkname LvImageGetAntialias C.lv_image_get_antialias
func LvImageGetAntialias(obj *core.LvObjT) bool

//go:linkname LvImageGetInnerAlign C.lv_image_get_inner_align
func LvImageGetInnerAlign(obj *core.LvObjT) LvImageAlignT

//go:linkname LvImageGetBitmapMapSrc C.lv_image_get_bitmap_map_src
func LvImageGetBitmapMapSrc(obj *core.LvObjT) *core.LvImageDscT
