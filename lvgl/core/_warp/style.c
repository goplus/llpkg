#include "src/misc/lv_style.h"

void lv_style_set_size_inline(lv_style_t * style, int32_t width, int32_t height)
{
   lv_style_set_size(style, width, height);
}

void lv_style_set_pad_all_inline(lv_style_t * style, int32_t value)
{
   lv_style_set_pad_all(style, value);
}

void lv_style_set_pad_hor_inline(lv_style_t * style, int32_t value)
{
    lv_style_set_pad_hor(style, value);\
}

void lv_style_set_pad_ver_inline(lv_style_t * style, int32_t value)
{
    lv_style_set_pad_ver(style, value);
}

void lv_style_set_pad_gap_inline(lv_style_t * style, int32_t value)
{
    lv_style_set_pad_gap(style, value);
}

void lv_style_set_margin_hor_inline(lv_style_t * style, int32_t value)
{
    lv_style_set_margin_hor(style, value);
}

void lv_style_set_margin_ver_inline(lv_style_t * style, int32_t value)
{
    lv_style_set_margin_ver(style, value);
}

void lv_style_set_margin_all_inline(lv_style_t * style, int32_t value)
{
    lv_style_set_margin_all(style, value);
}

void lv_style_set_transform_scale_inline(lv_style_t * style, int32_t value)
{
    lv_style_set_transform_scale(style, value);
}
