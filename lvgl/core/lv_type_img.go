package core

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
)

// lv_image_header_t 结构体（使用 uint32 存储位域信息）
//
//llgo:type C
type LvImageHeaderT struct {
	Magic uint8  // 魔数. 必须为 LV_IMAGE_HEADER_MAGIC
	CF    uint8  // 颜色格式: 见 `lv_color_format_t`
	Flags uint16 // 图片标志, 见 `lv_image_flags_t`

	Width    uint16 // 宽度
	Height   uint16 // 高度
	Stride   uint16 // 行字节数
	Reserved uint16 // 保留字段

}

// **lv_image_dsc_t 结构体**
//
//llgo:type C
type LvImageDscT struct {
	Header   LvImageHeaderT // 图片基本信息
	DataSize c.Uint32T      // 图片数据大小（字节）
	Data     *c.Char        // 指向图片数据的指针
	Reserved [0]byte        // 预留字段（保持与 C 结构体大小一致）
}
