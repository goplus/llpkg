package music

import (
	"github.com/goplus/llpkg/lvgl/core"
	"github.com/goplus/llpkg/lvgl/display"
	"github.com/goplus/llpkg/lvgl/layouts"
)
func lv_demo_music_list_create(parent *c.Void) *c.Void {
	font_small := lvgl.LV_FONT_DEFAULT
	font_medium := lvgl.LV_FONT_DEFAULT

	


    core.LvStyleInit(&style_scrollbar);
    core.LvStyleSetWidth(&style_scrollbar,  4);
    core.LvStyleSetBgOpa(&style_scrollbar, core.LV_OPA_COVER);
    core.LvStyleSetBgColor(&style_scrollbar, core.LvColorHex(0xeee));
    core.LvStyleSetRadius(&style_scrollbar, core.LV_RADIUS_CIRCLE);
    core.LvStyleSetPadRight(&style_scrollbar, 4);

    grid_cols := []c.Int{core.LV_GRID_CONTENT, core.LV_GRID_FR(1), core.LV_GRID_CONTENT, core.LV_GRID_TEMPLATE_LAST}
    grid_rows := []c.Int{35,  30, core.LV_GRID_TEMPLATE_LAST}

    core.LvStyleInit(&style_btn);
    core.LvStyleSetBgOpa(&style_btn, core.LV_OPA_TRANSP);
    core.LvStyleSetGridColumnDscArray(&style_btn, grid_cols);
    core.LvStyleSetGridRowDscArray(&style_btn, grid_rows);
    core.LvStyleSetGridRowAlign(&style_btn, core.LV_GRID_ALIGN_CENTER);
    core.LvStyleSetLayout(&style_btn, core.LV_LAYOUT_GRID);

    core.LvStyleSetPadRight(&style_btn, 30);

    core.LvStyleInit(&style_button_pr);
    core.LvStyleSetBgOpa(&style_button_pr, core.LV_OPA_COVER);
    core.LvStyleSetBgColor(&style_button_pr,  core.LvColorHex(0x4c4965));

    core.LvStyleInit(&style_button_chk);
    core.LvStyleSetBgOpa(&style_button_chk, core.LV_OPA_COVER);
    core.LvStyleSetBgColor(&style_button_chk, core.LvColorHex(0x4c4965));

    core.LvStyleInit(&style_button_dis);
    core.LvStyleSetTextOpa(&style_button_dis, core.LV_OPA_40);
    core.LvStyleSetImageOpa(&style_button_dis, core.LV_OPA_40);

    core.LvStyleInit(&style_title);
    core.LvStyleSetTextFont(&style_title, font_medium);
    core.LvStyleSetTextColor(&style_title, core.LvColorHex(0xffffff));

    core.LvStyleInit(&style_artist);
    core.LvStyleSetTextFont(&style_artist, font_small);
    core.LvStyleSetTextColor(&style_artist, core.LvColorHex(0xb1b0be));

    core.LvStyleInit(&style_time);
    core.LvStyleSetTextFont(&style_time, font_medium);
    core.LvStyleSetTextColor(&style_time, core.LvColorHex(0xffffff));

    /*Create an empty transparent container*/
    list := core.LvObjCreate(parent);
    core.LvObjAddEventCb(list, list_delete_event_cb, core.LV_EVENT_DELETE, nil);
    core.LvObjRemoveStyleAll(list);
    core.LvObjSetSize(list, core.LV_HOR_RES, core.LV_VER_RES - core.LV_DEMO_MUSIC_HANDLE_SIZE);
    core.LvObjSetY(list, core.LV_DEMO_MUSIC_HANDLE_SIZE);
    core.LvObjAddStyle(list, &style_scrollbar, core.LV_PART_SCROLLBAR);
    layouts.LvObjSetFlexFlow(list, core.LV_FLEX_FLOW_COLUMN);

    track_id := 0
    for(track_id < lv_demo_music_get_title(track_id); track_id++) {
        add_list_button(list,  track_id);
    }

// #if LV_DEMO_MUSIC_ROUND
//     lv_obj_set_scroll_snap_y(list, LV_SCROLL_SNAP_CENTER);
// #endif

    lv_demo_music_list_button_check(0, true);

    return list;
}



func add_list_button(parent *core.LvObjT, track_id int) *core.LvObjT {
    t := lv_demo_music_get_track_length(track_id);
    time := make([]c.Char, 32)
    c.Sprintf(time, "%"LV_PRIu32":%02"LV_PRIu32, t / 60, t % 60);
    title := lv_demo_music_get_title(track_id);
    artist := lv_demo_music_get_artist(track_id);

    btn := core.LvObjCreate(parent);
    core.LvObjRemoveStyleAll(btn);
    core.LvObjSetSize(btn, lv_pct(100), 110);


    core.LvObjAddStyle(btn, &style_btn, 0);
    core.LvObjAddStyle(btn, &style_button_pr, core.LV_STATE_PRESSED);
    core.LvObjAddStyle(btn, &style_button_chk, core.LV_STATE_CHECKED);
    core.LvObjAddStyle(btn, &style_button_dis, core.LV_STATE_DISABLED);
    core.LvObjAddEventCb(btn, btn_click_event_cb, core.LV_EVENT_CLICKED, nil);

    if(track_id >= 3) {
        core.LvObjAddState(btn, core.LV_STATE_DISABLED);
    }

    icon := widgets.LvImageCreate(btn);
    widgets.LvImageSetSrc(icon, &img_lv_demo_music_btn_list_play);
    layouts.LvObjSetGridCell(icon, core.LV_GRID_ALIGN_START, 0, 1, core.LV_GRID_ALIGN_CENTER, 0, 2);

    title_label := widgets.LvLabelCreate(btn);
    widgets.LvLabelSetText(title_label, title);
    layouts.LvObjSetGridCell(title_label, core.LV_GRID_ALIGN_START, 1, 1, core.LV_GRID_ALIGN_CENTER, 0, 1);
    core.LvObjAddStyle(title_label, &style_title, 0);

    artist_label := widgets.LvLabelCreate(btn);
    widgets.LvLabelSetText(artist_label, artist);
    core.LvObjAddStyle(artist_label, &style_artist, 0);
    layouts.LvObjSetGridCell(artist_label, core.LV_GRID_ALIGN_START, 1, 1, core.LV_GRID_ALIGN_CENTER, 1, 1);

    time_label := widgets.LvLabelCreate(btn);
    widgets.LvLabelSetText(time_label, time);
    core.LvObjAddStyle(time_label, &style_time, 0);
    layouts.LvObjSetGridCell(time_label, core.LV_GRID_ALIGN_END, 2, 1, core.LV_GRID_ALIGN_CENTER, 0, 2);

   
    border := widgets.LvImageCreate(btn);
    widgets.LvImageSetSrc(border, &img_lv_demo_music_list_border);
    widgets.LvImageSetInnerAlign(border, core.LV_IMAGE_ALIGN_TILE);
    core.LvObjSetWidth(border, core.LV_PCT(120));
    core.LvObjAlign(border, core.LV_ALIGN_BOTTOM_MID, 0, 0);
    core.LvObjAddFlag(border, core.LV_OBJ_FLAG_IGNORE_LAYOUT);

    return btn;
}

func btn_click_event_cb(e *core.LvEventT) {
    btn := core.LvEventGetTarget(e);

    idx := core.LvObjGetIndex(btn);

    lv_demo_music_play(idx);
}

func list_delete_event_cb(e *core.LvEventT) {
{
    code := core.LvEventGetCode(e);

    if(code == core.LV_EVENT_DELETE) {
        core.LvStyleReset(&style_scrollbar);
        core.LvStyleReset(&style_btn);
        core.LvStyleReset(&style_button_pr);
        core.LvStyleReset(&style_button_chk);
        core.LvStyleReset(&style_button_dis);
        core.LvStyleReset(&style_title);
        core.LvStyleReset(&style_artist);
        core.LvStyleReset(&style_time);
    }
}
