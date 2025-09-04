package widgets

import (
	"math"
	_ "unsafe"

	"github.com/goplus/llgo/c"
	core "github.com/goplus/llpkg/lvgl/core"
)

// Chart types
type LvChartTypeT uint8

const (
	LV_CHART_TYPE_NONE    LvChartTypeT = iota // Don't draw the series
	LV_CHART_TYPE_LINE                        // Connect the points with lines
	LV_CHART_TYPE_BAR                         // Draw columns
	LV_CHART_TYPE_SCATTER                     // Draw points and lines in 2D (x,y coordinates)
)

// Chart update mode
type LvChartUpdateModeT uint8

const (
	LV_CHART_UPDATE_MODE_SHIFT    LvChartUpdateModeT = iota // Shift old data to the left and add the new one the right
	LV_CHART_UPDATE_MODE_CIRCULAR                           // Add the new data in a circular way
)

// Chart axis
type LvChartAxisT uint8

const (
	LV_CHART_AXIS_PRIMARY_Y   LvChartAxisT = 0x00
	LV_CHART_AXIS_SECONDARY_Y LvChartAxisT = 0x01
	LV_CHART_AXIS_PRIMARY_X   LvChartAxisT = 0x02
	LV_CHART_AXIS_SECONDARY_X LvChartAxisT = 0x04
	LV_CHART_AXIS_LAST
)

const LV_CHART_POINT_NONE = math.MaxInt32

//go:linkname LvChartCreate C.lv_chart_create
func LvChartCreate(parent *core.LvObjT) *core.LvObjT

//go:linkname LvChartSetType C.lv_chart_set_type
func LvChartSetType(obj *core.LvObjT, type_ LvChartTypeT)

//go:linkname LvChartSetPointCount C.lv_chart_set_point_count
func LvChartSetPointCount(obj *core.LvObjT, cnt c.Uint32T)

//go:linkname LvChartSetAxisRange C.lv_chart_set_axis_range
func LvChartSetAxisRange(obj *core.LvObjT, axis LvChartAxisT, min c.Int32T, max c.Int32T)

//go:linkname LvChartSetUpdateMode C.lv_chart_set_update_mode
func LvChartSetUpdateMode(obj *core.LvObjT, update_mode LvChartUpdateModeT)

//go:linkname LvChartSetDivLineCount C.lv_chart_set_div_line_count
func LvChartSetDivLineCount(obj *core.LvObjT, hdiv c.Uint8T, vdiv c.Uint8T)

//go:linkname LvChartGetType C.lv_chart_get_type
func LvChartGetType(obj *core.LvObjT) LvChartTypeT

//go:linkname LvChartGetPointCount C.lv_chart_get_point_count
func LvChartGetPointCount(obj *core.LvObjT) c.Uint32T

//go:linkname LvChartGetXStartPoint C.lv_chart_get_x_start_point
func LvChartGetXStartPoint(obj *core.LvObjT, ser *core.LvChartSeriesT) c.Uint32T

//go:linkname LvChartGetPointPosByID C.lv_chart_get_point_pos_by_id
func LvChartGetPointPosByID(obj *core.LvObjT, ser *core.LvChartSeriesT, id c.Uint32T, p_out *core.LvPointT)

//go:linkname LvChartRefresh C.lv_chart_refresh
func LvChartRefresh(obj *core.LvObjT)

//go:linkname LvChartAddSeries C.lv_chart_add_series
func LvChartAddSeries(obj *core.LvObjT, color core.LvColorT, axis LvChartAxisT) *core.LvChartSeriesT

//go:linkname LvChartRemoveSeries C.lv_chart_remove_series
func LvChartRemoveSeries(obj *core.LvObjT, series *core.LvChartSeriesT)

//go:linkname LvChartHideSeries C.lv_chart_hide_series
func LvChartHideSeries(chart *core.LvObjT, series *core.LvChartSeriesT, hide bool)

//go:linkname LvChartSetSeriesColor C.lv_chart_set_series_color
func LvChartSetSeriesColor(chart *core.LvObjT, series *core.LvChartSeriesT, color core.LvColorT)

//go:linkname LvChartGetSeriesColor C.lv_chart_get_series_color
func LvChartGetSeriesColor(chart *core.LvObjT, series *core.LvChartSeriesT) core.LvColorT

//go:linkname LvChartSetXStartPoint C.lv_chart_set_x_start_point
func LvChartSetXStartPoint(obj *core.LvObjT, ser *core.LvChartSeriesT, id c.Uint32T)

//go:linkname LvChartGetSeriesNext C.lv_chart_get_series_next
func LvChartGetSeriesNext(chart *core.LvObjT, ser *core.LvChartSeriesT) *core.LvChartSeriesT

//go:linkname LvChartAddCursor C.lv_chart_add_cursor
func LvChartAddCursor(obj *core.LvObjT, color core.LvColorT, dir core.LvDirT) *core.LvChartCursorT

//go:linkname LvChartSetCursorPos C.lv_chart_set_cursor_pos
func LvChartSetCursorPos(chart *core.LvObjT, cursor *core.LvChartCursorT, pos *core.LvPointT)

//go:linkname LvChartSetCursorPoint C.lv_chart_set_cursor_point
func LvChartSetCursorPoint(chart *core.LvObjT, cursor *core.LvChartCursorT, ser *core.LvChartSeriesT, point_id c.Uint32T)

//go:linkname LvChartGetCursorPoint C.lv_chart_get_cursor_point
func LvChartGetCursorPoint(chart *core.LvObjT, cursor *core.LvChartCursorT) core.LvPointT

//go:linkname LvChartSetAllValues C.lv_chart_set_all_values
func LvChartSetAllValues(obj *core.LvObjT, ser *core.LvChartSeriesT, value c.Int32T)

//go:linkname LvChartSetNextValue C.lv_chart_set_next_value
func LvChartSetNextValue(obj *core.LvObjT, ser *core.LvChartSeriesT, value c.Int32T)

//go:linkname LvChartSetNextValue2 C.lv_chart_set_next_value2
func LvChartSetNextValue2(obj *core.LvObjT, ser *core.LvChartSeriesT, x_value c.Int32T, y_value c.Int32T)

//go:linkname LvChartSetSeriesValues C.lv_chart_set_series_values
func LvChartSetSeriesValues(obj *core.LvObjT, ser *core.LvChartSeriesT, values *c.Int32T, values_cnt c.SizeT)

//go:linkname LvChartSetSeriesValues2 C.lv_chart_set_series_values2
func LvChartSetSeriesValues2(obj *core.LvObjT, ser *core.LvChartSeriesT, x_values *c.Int32T, y_values *c.Int32T, values_cnt c.SizeT)

//go:linkname LvChartSetSeriesValueByID C.lv_chart_set_series_value_by_id
func LvChartSetSeriesValueByID(obj *core.LvObjT, ser *core.LvChartSeriesT, id c.Uint32T, value c.Int32T)

//go:linkname LvChartSetSeriesValueByID2 C.lv_chart_set_series_value_by_id2
func LvChartSetSeriesValueByID2(obj *core.LvObjT, ser *core.LvChartSeriesT, id c.Uint32T, x_value c.Int32T, y_value c.Int32T)

//go:linkname LvChartSetSeriesExtYArray C.lv_chart_set_series_ext_y_array
func LvChartSetSeriesExtYArray(obj *core.LvObjT, ser *core.LvChartSeriesT, array *c.Int32T)

//go:linkname LvChartSetSeriesExtXArray C.lv_chart_set_series_ext_x_array
func LvChartSetSeriesExtXArray(obj *core.LvObjT, ser *core.LvChartSeriesT, array *c.Int32T)

//go:linkname LvChartGetSeriesYArray C.lv_chart_get_series_y_array
func LvChartGetSeriesYArray(obj *core.LvObjT, ser *core.LvChartSeriesT) *c.Int32T

//go:linkname LvChartGetSeriesXArray C.lv_chart_get_series_x_array
func LvChartGetSeriesXArray(obj *core.LvObjT, ser *core.LvChartSeriesT) *c.Int32T

//go:linkname LvChartGetPressedPoint C.lv_chart_get_pressed_point
func LvChartGetPressedPoint(obj *core.LvObjT) c.Uint32T

//go:linkname LvChartGetFirstPointCenterOffset C.lv_chart_get_first_point_center_offset
func LvChartGetFirstPointCenterOffset(obj *core.LvObjT) c.Int32T
