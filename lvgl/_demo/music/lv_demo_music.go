package main

import (
	"github.com/goplus/llgo/c"
	"github.com/goplus/llpkg/lvgl/core"
	"github.com/goplus/llpkg/lvgl/display"
)


const LV_DEMO_MUSIC_HANDLE_SIZE  40

/**********************
 *  STATIC VARIABLES
 **********************/
var ctrl *core.LvObjT
var list *core.LvObjT

var title_list = []string{
	"Waiting for true love",
	"Need a Better Future",
	"Vibrations",
	"Why now?",
	"Never Look Back",
	"It happened Yesterday",
	"Feeling so High",
	"Go Deeper",
	"Find You There",
	"Until the End",
	"Unknown",
	"Unknown",
	"Unknown",
	"Unknown",
}

var artist_list = []str{
	"The John Smith Band",
	"My True Name",
	"Robotics",
	"John Smith",
	"My True Name",
	"Robotics",
	"Robotics",
	"Unknown artist",
	"Unknown artist",
	"Unknown artist",
	"Unknown artist",
	"Unknown artist",
	"Unknown artist",
	"Unknown artist",
	"Unknown artist",
}

var genre_list = []string{
	"Rock - 1997",
	"Drum'n bass - 2016",
	"Psy trance - 2020",
	"Metal - 2015",
	"Metal - 2015",
	"Metal - 2015",
	"Metal - 2015",
	"Metal - 2015",
	"Metal - 2015",
	"Metal - 2015",
	"Metal - 2015",
	"Metal - 2015",
	"Metal - 2015",
	"Metal - 2015",
}

var time_list = []c.Uint{
	1*60 + 14,
	2*60 + 26,
	1*60 + 54,
	2*60 + 24,
	2*60 + 37,
	3*60 + 33,
	1*60 + 56,
	3*60 + 31,
	2*60 + 20,
	2*60 + 19,
	2*60 + 20,
	2*60 + 19,
	2*60 + 20,
	2*60 + 19,
}

func lv_demo_music() {

	core.LvObjSetStyleBgColor(display.LvScreenActive(), core.LvColorHex(0x343247), 0)

	// list = _lv_demo_music_list_create(lv_scr_act())
	// ctrl = _lv_demo_music_main_create(lv_scr_act())

	//lv_timer_create(auto_step_cb, 1000, NULL)

}



func lv_demo_music_get_title(track_id int) *c.Char {
    if(track_id >= len(title_list) / len(title_list[0])) return nil;
    return title_list[track_id];
}

func lv_demo_music_get_artist(track_id int) *c.Char {
    if(track_id >= len(artist_list) / len(artist_list[0])) return nil;
    return artist_list[track_id];
}

func lv_demo_music_get_genre(track_id int) *c.Char {
    if(track_id >= len(genre_list) / len(genre_list[0])) return nil;
    return genre_list[track_id];
}

func lv_demo_music_get_track_length(track_id int) c.Uint {
    if(track_id >= len(time_list) / len(time_list[0])) return 0;
    return time_list[track_id];
}

