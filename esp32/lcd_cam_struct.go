package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: LCD configuration registers */
/** Type of lcd_clock register
 *  LCD clock configuration register
 */

type LcdCamLcdClockRegT struct {
	Val c.Uint32T
}

/** Type of lcd_rgb_yuv register
 *  LCD data format conversion register
 */

type LcdCamLcdRgbYuvRegT struct {
	Val c.Uint32T
}

/** Type of lcd_user register
 *  LCD user configuration register
 */

type LcdCamLcdUserRegT struct {
	Val c.Uint32T
}

/** Type of lcd_misc register
 *  LCD MISC configuration register
 */

type LcdCamLcdMiscRegT struct {
	Val c.Uint32T
}

/** Type of lcd_ctrl register
 *  LCD signal configuration register
 */

type LcdCamLcdCtrlRegT struct {
	Val c.Uint32T
}

/** Type of lcd_ctrl1 register
 *  LCD signal configuration register 1
 */

type LcdCamLcdCtrl1RegT struct {
	Val c.Uint32T
}

/** Type of lcd_ctrl2 register
 *  LCD signal configuration register 2
 */

type LcdCamLcdCtrl2RegT struct {
	Val c.Uint32T
}

/** Type of lcd_cmd_val register
 *  LCD command value configuration register
 */

type LcdCamLcdCmdValRegT struct {
	Val c.Uint32T
}

/** Type of lcd_dly_mode register
 *  LCD signal delay configuration register
 */

type LcdCamLcdDlyModeRegT struct {
	Val c.Uint32T
}

/** Type of lcd_data_dout_mode register
 *  LCD data delay configuration register
 */

type LcdCamLcdDataDoutModeRegT struct {
	Val c.Uint32T
}

/** Group: Camera configuration registers */
/** Type of cam_ctrl register
 *  Camera clock configuration register
 */

type LcdCamCamCtrlRegT struct {
	Val c.Uint32T
}

/** Type of cam_ctrl1 register
 *  Camera control register
 */

type LcdCamCamCtrl1RegT struct {
	Val c.Uint32T
}

/** Type of cam_rgb_yuv register
 *  Camera data format conversion register
 */

type LcdCamCamRgbYuvRegT struct {
	Val c.Uint32T
}

/** Group: Interrupt registers */
/** Type of lc_dma_int_ena register
 *  LCD_CAM GDMA interrupt enable register
 */

type LcdCamLcDmaIntEnaRegT struct {
	Val c.Uint32T
}

/** Type of lc_dma_int_raw register
 *  LCD_CAM GDMA raw interrupt status register
 */

type LcdCamLcDmaIntRawRegT struct {
	Val c.Uint32T
}

/** Type of lc_dma_int_st register
 *  LCD_CAM GDMA masked interrupt status register
 */

type LcdCamLcDmaIntStRegT struct {
	Val c.Uint32T
}

/** Type of lc_dma_int_clr register
 *  LCD_CAM GDMA interrupt clear register
 */

type LcdCamLcDmaIntClrRegT struct {
	Val c.Uint32T
}

/** Group: Version register */
/** Type of lc_reg_date register
 *  Version control register
 */

type LcdCamLcRegDateRegT struct {
	Val c.Uint32T
}

type LcdCamDevT struct {
	LcdClock        LcdCamLcdClockRegT
	CamCtrl         LcdCamCamCtrlRegT
	CamCtrl1        LcdCamCamCtrl1RegT
	CamRgbYuv       LcdCamCamRgbYuvRegT
	LcdRgbYuv       LcdCamLcdRgbYuvRegT
	LcdUser         LcdCamLcdUserRegT
	LcdMisc         LcdCamLcdMiscRegT
	LcdCtrl         LcdCamLcdCtrlRegT
	LcdCtrl1        LcdCamLcdCtrl1RegT
	LcdCtrl2        LcdCamLcdCtrl2RegT
	LcdCmdVal       LcdCamLcdCmdValRegT
	Reserved02c     c.Uint32T
	LcdDlyMode      LcdCamLcdDlyModeRegT
	Reserved034     c.Uint32T
	LcdDataDoutMode LcdCamLcdDataDoutModeRegT
	Reserved03c     [10]c.Uint32T
	LcDmaIntEna     LcdCamLcDmaIntEnaRegT
	LcDmaIntRaw     LcdCamLcDmaIntRawRegT
	LcDmaIntSt      LcdCamLcDmaIntStRegT
	LcDmaIntClr     LcdCamLcDmaIntClrRegT
	Reserved074     [34]c.Uint32T
	LcRegDate       LcdCamLcRegDateRegT
}
