package core

import (
	"unsafe"
	_ "unsafe"

	"github.com/goplus/llgo/c"
)

// Object tree walk result
type LvObjTreeWalkResT uint8

const (
	LV_OBJ_TREE_WALK_NEXT          LvObjTreeWalkResT = iota // Continue the walk
	LV_OBJ_TREE_WALK_SKIP_CHILDREN                          // Skip the children of the current node
	LV_OBJ_TREE_WALK_END                                    // End the walk
)

// Object tree walk callback function type
//
//go:type C
type LvObjTreeWalkCbT func(obj *LvObjT, userData *c.Void) LvObjTreeWalkResT

//go:linkname LvObjDelete C.lv_obj_delete
func LvObjDelete(obj *LvObjT)

//go:linkname LvObjClean C.lv_obj_clean
func LvObjClean(obj *LvObjT)

//go:linkname LvObjDeleteDelayed C.lv_obj_delete_delayed
func LvObjDeleteDelayed(obj *LvObjT, delay_ms uint32)

//go:linkname LvObjDeleteAnimCompletedCb C.lv_obj_delete_anim_completed_cb
func LvObjDeleteAnimCompletedCb(a *LvAnimT)

//go:linkname LvObjDeleteAsync C.lv_obj_delete_async
func LvObjDeleteAsync(obj *LvObjT)

//go:linkname LvObjSetParent C.lv_obj_set_parent
func LvObjSetParent(obj *LvObjT, parent *LvObjT)

//go:linkname LvObjSwap C.lv_obj_swap
func LvObjSwap(obj1 *LvObjT, obj2 *LvObjT)

//go:linkname LvObjMoveToIndex C.lv_obj_move_to_index
func LvObjMoveToIndex(obj *LvObjT, index int32)

//go:linkname LvObjGetScreen C.lv_obj_get_screen
func LvObjGetScreen(obj *LvObjT) *LvObjT

//go:linkname LvObjGetDisplay C.lv_obj_get_display
func LvObjGetDisplay(obj *LvObjT) *LvDisplayT

//go:linkname LvObjGetParent C.lv_obj_get_parent
func LvObjGetParent(obj *LvObjT) *LvObjT

//go:linkname LvObjGetChild C.lv_obj_get_child
func LvObjGetChild(obj *LvObjT, idx int32) *LvObjT

//go:linkname LvObjGetChildByType C.lv_obj_get_child_by_type
func LvObjGetChildByType(obj *LvObjT, idx int32, class_p *LvObjClassT) *LvObjT

//go:linkname LvObjGetSibling C.lv_obj_get_sibling
func LvObjGetSibling(obj *LvObjT, idx int32) *LvObjT

//go:linkname LvObjGetSiblingByType C.lv_obj_get_sibling_by_type
func LvObjGetSiblingByType(obj *LvObjT, idx int32, class_p *LvObjClassT) *LvObjT

//go:linkname LvObjGetChildCount C.lv_obj_get_child_count
func LvObjGetChildCount(obj *LvObjT) uint32

//go:linkname LvObjGetChildCountByType C.lv_obj_get_child_count_by_type
func LvObjGetChildCountByType(obj *LvObjT, class_p *LvObjClassT) uint32

// LV_USE_OBJ_NAME functions
//
//go:linkname LvObjSetName C.lv_obj_set_name
func LvObjSetName(obj *LvObjT, name *c.Char)

//go:linkname LvObjSetNameStatic C.lv_obj_set_name_static
func LvObjSetNameStatic(obj *LvObjT, name *c.Char)

//go:linkname LvObjGetName C.lv_obj_get_name
func LvObjGetName(obj *LvObjT) *c.Char

//go:linkname LvObjGetNameResolved C.lv_obj_get_name_resolved
func LvObjGetNameResolved(obj *LvObjT, buf *c.Char, buf_size c.SizeT)

//go:linkname LvObjFindByName C.lv_obj_find_by_name
func LvObjFindByName(parent *LvObjT, name *c.Char) *LvObjT

//go:linkname LvObjGetChildByName C.lv_obj_get_child_by_name
func LvObjGetChildByName(parent *LvObjT, name_path *c.Char) *LvObjT

//go:linkname LvObjGetIndex C.lv_obj_get_index
func LvObjGetIndex(obj *LvObjT) int32

//go:linkname LvObjGetIndexByType C.lv_obj_get_index_by_type
func LvObjGetIndexByType(obj *LvObjT, class_p *LvObjClassT) int32

//go:linkname LvObjTreeWalk C.lv_obj_tree_walk
func LvObjTreeWalk(start_obj *LvObjT, cb LvObjTreeWalkCbT, user_data unsafe.Pointer)

//go:linkname LvObjDumpTree C.lv_obj_dump_tree
func LvObjDumpTree(start_obj *LvObjT)
