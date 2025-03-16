package widgets

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
	core "github.com/goplus/llpkg/lvgl/core"
)

//go:linkname LvTileviewCreate C.lv_tileview_create
func LvTileviewCreate(parent *core.LvObjT) *core.LvObjT

//go:linkname LvTileviewAddTile C.lv_tileview_add_tile
func LvTileviewAddTile(tv *core.LvObjT, colId c.Uint8T, rowId c.Uint8T, dir c.Int32T) *core.LvObjT

//go:linkname LvTileviewSetTile C.lv_tileview_set_tile
func LvTileviewSetTile(tv *core.LvObjT, tileObj *core.LvObjT, animEn c.Int32T)

//go:linkname LvTileviewSetTileByIndex C.lv_tileview_set_tile_by_index
func LvTileviewSetTileByIndex(tv *core.LvObjT, colId c.Uint32T, rowId c.Uint32T, animEn c.Int32T)

//go:linkname LvTileviewGetTileActive C.lv_tileview_get_tile_active
func LvTileviewGetTileActive(obj *core.LvObjT) *core.LvObjT
