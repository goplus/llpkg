/**
 * @file lv_animimage.h
 *
 */

package widgets

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
	core "github.com/goplus/llpkg/lvgl/core"
)

type lv_animimg_part_t c.Int

const (
	LV_ANIM_IMAGE_PART_MAIN lv_animimg_part_t = iota
)

//go:linkname LvAnimimgCreate C.lv_animimg_create
func LvAnimimgCreate(parent *core.LvObjT) *core.LvObjT

//go:linkname LvAnimimgSetSrc C.lv_animimg_set_src
func LvAnimimgSetSrc(img *core.LvObjT, dsc []*c.Void, num c.SizeT)

//go:linkname LvAnimimgStart C.lv_animimg_start
func LvAnimimgStart(obj *core.LvObjT)

//go:linkname LvAnimimgSetDuration C.lv_animimg_set_duration
func LvAnimimgSetDuration(img *core.LvObjT, duration c.Uint32T)

//go:linkname LvAnimimgSetRepeatCount C.lv_animimg_set_repeat_count
func LvAnimimgSetRepeatCount(img *core.LvObjT, count c.Uint32T)

//go:linkname LvAnimimgGetSrc C.lv_animimg_get_src
func LvAnimimgGetSrc(img *core.LvObjT) []*c.Void

//go:linkname LvAnimimgGetSrcCount C.lv_animimg_get_src_count
func LvAnimimgGetSrcCount(img *core.LvObjT) c.Uint8T

//go:linkname LvAnimimgGetDuration C.lv_animimg_get_duration
func LvAnimimgGetDuration(img *core.LvObjT) c.Uint32T

//go:linkname LvAnimimgGetRepeatCount C.lv_animimg_get_repeat_count
func LvAnimimgGetRepeatCount(img *core.LvObjT) c.Uint32T

//go:linkname LvAnimimgGetAnim C.lv_animimg_get_anim
func LvAnimimgGetAnim(img *core.LvObjT) *core.LvAnimT
