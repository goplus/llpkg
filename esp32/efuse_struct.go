package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: PGM Data Register */
/** Type of pgm_data0 register
 *  Register 0 that stores data to be programmed.
 */

type EfusePgmData0RegT struct {
	Val c.Uint32T
}

/** Type of pgm_data1 register
 *  Register 1 that stores data to be programmed.
 */

type EfusePgmData1RegT struct {
	Val c.Uint32T
}

/** Type of pgm_data2 register
 *  Register 2 that stores data to be programmed.
 */

type EfusePgmData2RegT struct {
	Val c.Uint32T
}

/** Type of pgm_data3 register
 *  Register 3 that stores data to be programmed.
 */

type EfusePgmData3RegT struct {
	Val c.Uint32T
}

/** Type of pgm_data4 register
 *  Register 4 that stores data to be programmed.
 */

type EfusePgmData4RegT struct {
	Val c.Uint32T
}

/** Type of pgm_data5 register
 *  Register 5 that stores data to be programmed.
 */

type EfusePgmData5RegT struct {
	Val c.Uint32T
}

/** Type of pgm_data6 register
 *  Register 6 that stores data to be programmed.
 */

type EfusePgmData6RegT struct {
	Val c.Uint32T
}

/** Type of pgm_data7 register
 *  Register 7 that stores data to be programmed.
 */

type EfusePgmData7RegT struct {
	Val c.Uint32T
}

/** Type of pgm_check_value0 register
 *  Register 0 that stores the RS code to be programmed.
 */

type EfusePgmCheckValue0RegT struct {
	Val c.Uint32T
}

/** Type of pgm_check_value1 register
 *  Register 1 that stores the RS code to be programmed.
 */

type EfusePgmCheckValue1RegT struct {
	Val c.Uint32T
}

/** Type of pgm_check_value2 register
 *  Register 2 that stores the RS code to be programmed.
 */

type EfusePgmCheckValue2RegT struct {
	Val c.Uint32T
}

/** Group: Read Data Register */
/** Type of rd_wr_dis register
 *  BLOCK0 data register 0.
 */

type EfuseRdWrDisRegT struct {
	Val c.Uint32T
}

/** Type of rd_repeat_data0 register
 *  BLOCK0 data register 1.
 */

type EfuseRdRepeatData0RegT struct {
	Val c.Uint32T
}

/** Type of rd_repeat_data1 register
 *  BLOCK0 data register 2.
 */

type EfuseRdRepeatData1RegT struct {
	Val c.Uint32T
}

/** Type of rd_repeat_data2 register
 *  BLOCK0 data register 3.
 */

type EfuseRdRepeatData2RegT struct {
	Val c.Uint32T
}

/** Type of rd_repeat_data3 register
 *  BLOCK0 data register 4.
 */

type EfuseRdRepeatData3RegT struct {
	Val c.Uint32T
}

/** Type of rd_repeat_data4 register
 *  BLOCK0 data register 5.
 */

type EfuseRdRepeatData4RegT struct {
	Val c.Uint32T
}

/** Type of rd_mac_spi_sys_0 register
 *  BLOCK1 data register 0.
 */

type EfuseRdMacSpiSys0RegT struct {
	Val c.Uint32T
}

/** Type of rd_mac_spi_sys_1 register
 *  BLOCK1 data register 1.
 */

type EfuseRdMacSpiSys1RegT struct {
	Val c.Uint32T
}

/** Type of rd_mac_spi_sys_2 register
 *  BLOCK1 data register 2.
 */

type EfuseRdMacSpiSys2RegT struct {
	Val c.Uint32T
}

/** Type of rd_mac_spi_sys_3 register
 *  BLOCK1 data register 3.
 */

type EfuseRdMacSpiSys3RegT struct {
	Val c.Uint32T
}

/** Type of rd_mac_spi_sys_4 register
 *  BLOCK1 data register 4.
 */

type EfuseRdMacSpiSys4RegT struct {
	Val c.Uint32T
}

/** Type of rd_mac_spi_sys_5 register
 *  BLOCK1 data register 5.
 */

type EfuseRdMacSpiSys5RegT struct {
	Val c.Uint32T
}

/** Type of rd_sys_part1_data0 register
 *  Register 0 of BLOCK2 (system).
 */

type EfuseRdSysPart1Data0RegT struct {
	Val c.Uint32T
}

/** Type of rd_sys_part1_data1 register
 *  Register 1 of BLOCK2 (system).
 */

type EfuseRdSysPart1Data1RegT struct {
	Val c.Uint32T
}

/** Type of rd_sys_part1_data2 register
 *  Register 2 of BLOCK2 (system).
 */

type EfuseRdSysPart1Data2RegT struct {
	Val c.Uint32T
}

/** Type of rd_sys_part1_data3 register
 *  Register 3 of BLOCK2 (system).
 */

type EfuseRdSysPart1Data3RegT struct {
	Val c.Uint32T
}

/** Type of rd_sys_part1_data4 register
 *  Register 4 of BLOCK2 (system).
 */

type EfuseRdSysPart1Data4RegT struct {
	Val c.Uint32T
}

/** Type of rd_sys_part1_data5 register
 *  Register 5 of BLOCK2 (system).
 */

type EfuseRdSysPart1Data5RegT struct {
	Val c.Uint32T
}

/** Type of rd_sys_part1_data6 register
 *  Register 6 of BLOCK2 (system).
 */

type EfuseRdSysPart1Data6RegT struct {
	Val c.Uint32T
}

/** Type of rd_sys_part1_data7 register
 *  Register 7 of BLOCK2 (system).
 */

type EfuseRdSysPart1Data7RegT struct {
	Val c.Uint32T
}

/** Type of rd_usr_data0 register
 *  Register 0 of BLOCK3 (user).
 */

type EfuseRdUsrData0RegT struct {
	Val c.Uint32T
}

/** Type of rd_usr_data1 register
 *  Register 1 of BLOCK3 (user).
 */

type EfuseRdUsrData1RegT struct {
	Val c.Uint32T
}

/** Type of rd_usr_data2 register
 *  Register 2 of BLOCK3 (user).
 */

type EfuseRdUsrData2RegT struct {
	Val c.Uint32T
}

/** Type of rd_usr_data3 register
 *  Register 3 of BLOCK3 (user).
 */

type EfuseRdUsrData3RegT struct {
	Val c.Uint32T
}

/** Type of rd_usr_data4 register
 *  Register 4 of BLOCK3 (user).
 */

type EfuseRdUsrData4RegT struct {
	Val c.Uint32T
}

/** Type of rd_usr_data5 register
 *  Register 5 of BLOCK3 (user).
 */

type EfuseRdUsrData5RegT struct {
	Val c.Uint32T
}

/** Type of rd_usr_data6 register
 *  Register 6 of BLOCK3 (user).
 */

type EfuseRdUsrData6RegT struct {
	Val c.Uint32T
}

/** Type of rd_usr_data7 register
 *  Register 7 of BLOCK3 (user).
 */

type EfuseRdUsrData7RegT struct {
	Val c.Uint32T
}

/** Type of rd_key0_data0 register
 *  Register 0 of BLOCK4 (KEY0).
 */

type EfuseRdKey0Data0RegT struct {
	Val c.Uint32T
}

/** Type of rd_key0_data1 register
 *  Register 1 of BLOCK4 (KEY0).
 */

type EfuseRdKey0Data1RegT struct {
	Val c.Uint32T
}

/** Type of rd_key0_data2 register
 *  Register 2 of BLOCK4 (KEY0).
 */

type EfuseRdKey0Data2RegT struct {
	Val c.Uint32T
}

/** Type of rd_key0_data3 register
 *  Register 3 of BLOCK4 (KEY0).
 */

type EfuseRdKey0Data3RegT struct {
	Val c.Uint32T
}

/** Type of rd_key0_data4 register
 *  Register 4 of BLOCK4 (KEY0).
 */

type EfuseRdKey0Data4RegT struct {
	Val c.Uint32T
}

/** Type of rd_key0_data5 register
 *  Register 5 of BLOCK4 (KEY0).
 */

type EfuseRdKey0Data5RegT struct {
	Val c.Uint32T
}

/** Type of rd_key0_data6 register
 *  Register 6 of BLOCK4 (KEY0).
 */

type EfuseRdKey0Data6RegT struct {
	Val c.Uint32T
}

/** Type of rd_key0_data7 register
 *  Register 7 of BLOCK4 (KEY0).
 */

type EfuseRdKey0Data7RegT struct {
	Val c.Uint32T
}

/** Type of rd_key1_data0 register
 *  Register 0 of BLOCK5 (KEY1).
 */

type EfuseRdKey1Data0RegT struct {
	Val c.Uint32T
}

/** Type of rd_key1_data1 register
 *  Register 1 of BLOCK5 (KEY1).
 */

type EfuseRdKey1Data1RegT struct {
	Val c.Uint32T
}

/** Type of rd_key1_data2 register
 *  Register 2 of BLOCK5 (KEY1).
 */

type EfuseRdKey1Data2RegT struct {
	Val c.Uint32T
}

/** Type of rd_key1_data3 register
 *  Register 3 of BLOCK5 (KEY1).
 */

type EfuseRdKey1Data3RegT struct {
	Val c.Uint32T
}

/** Type of rd_key1_data4 register
 *  Register 4 of BLOCK5 (KEY1).
 */

type EfuseRdKey1Data4RegT struct {
	Val c.Uint32T
}

/** Type of rd_key1_data5 register
 *  Register 5 of BLOCK5 (KEY1).
 */

type EfuseRdKey1Data5RegT struct {
	Val c.Uint32T
}

/** Type of rd_key1_data6 register
 *  Register 6 of BLOCK5 (KEY1).
 */

type EfuseRdKey1Data6RegT struct {
	Val c.Uint32T
}

/** Type of rd_key1_data7 register
 *  Register 7 of BLOCK5 (KEY1).
 */

type EfuseRdKey1Data7RegT struct {
	Val c.Uint32T
}

/** Type of rd_key2_data0 register
 *  Register 0 of BLOCK6 (KEY2).
 */

type EfuseRdKey2Data0RegT struct {
	Val c.Uint32T
}

/** Type of rd_key2_data1 register
 *  Register 1 of BLOCK6 (KEY2).
 */

type EfuseRdKey2Data1RegT struct {
	Val c.Uint32T
}

/** Type of rd_key2_data2 register
 *  Register 2 of BLOCK6 (KEY2).
 */

type EfuseRdKey2Data2RegT struct {
	Val c.Uint32T
}

/** Type of rd_key2_data3 register
 *  Register 3 of BLOCK6 (KEY2).
 */

type EfuseRdKey2Data3RegT struct {
	Val c.Uint32T
}

/** Type of rd_key2_data4 register
 *  Register 4 of BLOCK6 (KEY2).
 */

type EfuseRdKey2Data4RegT struct {
	Val c.Uint32T
}

/** Type of rd_key2_data5 register
 *  Register 5 of BLOCK6 (KEY2).
 */

type EfuseRdKey2Data5RegT struct {
	Val c.Uint32T
}

/** Type of rd_key2_data6 register
 *  Register 6 of BLOCK6 (KEY2).
 */

type EfuseRdKey2Data6RegT struct {
	Val c.Uint32T
}

/** Type of rd_key2_data7 register
 *  Register 7 of BLOCK6 (KEY2).
 */

type EfuseRdKey2Data7RegT struct {
	Val c.Uint32T
}

/** Type of rd_key3_data0 register
 *  Register 0 of BLOCK7 (KEY3).
 */

type EfuseRdKey3Data0RegT struct {
	Val c.Uint32T
}

/** Type of rd_key3_data1 register
 *  Register 1 of BLOCK7 (KEY3).
 */

type EfuseRdKey3Data1RegT struct {
	Val c.Uint32T
}

/** Type of rd_key3_data2 register
 *  Register 2 of BLOCK7 (KEY3).
 */

type EfuseRdKey3Data2RegT struct {
	Val c.Uint32T
}

/** Type of rd_key3_data3 register
 *  Register 3 of BLOCK7 (KEY3).
 */

type EfuseRdKey3Data3RegT struct {
	Val c.Uint32T
}

/** Type of rd_key3_data4 register
 *  Register 4 of BLOCK7 (KEY3).
 */

type EfuseRdKey3Data4RegT struct {
	Val c.Uint32T
}

/** Type of rd_key3_data5 register
 *  Register 5 of BLOCK7 (KEY3).
 */

type EfuseRdKey3Data5RegT struct {
	Val c.Uint32T
}

/** Type of rd_key3_data6 register
 *  Register 6 of BLOCK7 (KEY3).
 */

type EfuseRdKey3Data6RegT struct {
	Val c.Uint32T
}

/** Type of rd_key3_data7 register
 *  Register 7 of BLOCK7 (KEY3).
 */

type EfuseRdKey3Data7RegT struct {
	Val c.Uint32T
}

/** Type of rd_key4_data0 register
 *  Register 0 of BLOCK8 (KEY4).
 */

type EfuseRdKey4Data0RegT struct {
	Val c.Uint32T
}

/** Type of rd_key4_data1 register
 *  Register 1 of BLOCK8 (KEY4).
 */

type EfuseRdKey4Data1RegT struct {
	Val c.Uint32T
}

/** Type of rd_key4_data2 register
 *  Register 2 of BLOCK8 (KEY4).
 */

type EfuseRdKey4Data2RegT struct {
	Val c.Uint32T
}

/** Type of rd_key4_data3 register
 *  Register 3 of BLOCK8 (KEY4).
 */

type EfuseRdKey4Data3RegT struct {
	Val c.Uint32T
}

/** Type of rd_key4_data4 register
 *  Register 4 of BLOCK8 (KEY4).
 */

type EfuseRdKey4Data4RegT struct {
	Val c.Uint32T
}

/** Type of rd_key4_data5 register
 *  Register 5 of BLOCK8 (KEY4).
 */

type EfuseRdKey4Data5RegT struct {
	Val c.Uint32T
}

/** Type of rd_key4_data6 register
 *  Register 6 of BLOCK8 (KEY4).
 */

type EfuseRdKey4Data6RegT struct {
	Val c.Uint32T
}

/** Type of rd_key4_data7 register
 *  Register 7 of BLOCK8 (KEY4).
 */

type EfuseRdKey4Data7RegT struct {
	Val c.Uint32T
}

/** Type of rd_key5_data0 register
 *  Register 0 of BLOCK9 (KEY5).
 */

type EfuseRdKey5Data0RegT struct {
	Val c.Uint32T
}

/** Type of rd_key5_data1 register
 *  Register 1 of BLOCK9 (KEY5).
 */

type EfuseRdKey5Data1RegT struct {
	Val c.Uint32T
}

/** Type of rd_key5_data2 register
 *  Register 2 of BLOCK9 (KEY5).
 */

type EfuseRdKey5Data2RegT struct {
	Val c.Uint32T
}

/** Type of rd_key5_data3 register
 *  Register 3 of BLOCK9 (KEY5).
 */

type EfuseRdKey5Data3RegT struct {
	Val c.Uint32T
}

/** Type of rd_key5_data4 register
 *  Register 4 of BLOCK9 (KEY5).
 */

type EfuseRdKey5Data4RegT struct {
	Val c.Uint32T
}

/** Type of rd_key5_data5 register
 *  Register 5 of BLOCK9 (KEY5).
 */

type EfuseRdKey5Data5RegT struct {
	Val c.Uint32T
}

/** Type of rd_key5_data6 register
 *  Register 6 of BLOCK9 (KEY5).
 */

type EfuseRdKey5Data6RegT struct {
	Val c.Uint32T
}

/** Type of rd_key5_data7 register
 *  Register 7 of BLOCK9 (KEY5).
 */

type EfuseRdKey5Data7RegT struct {
	Val c.Uint32T
}

/** Type of rd_sys_part2_data0 register
 *  Register 0 of BLOCK10 (system).
 */

type EfuseRdSysPart2Data0RegT struct {
	Val c.Uint32T
}

/** Type of rd_sys_part2_data1 register
 *  Register 1 of BLOCK9 (KEY5).
 */

type EfuseRdSysPart2Data1RegT struct {
	Val c.Uint32T
}

/** Type of rd_sys_part2_data2 register
 *  Register 2 of BLOCK10 (system).
 */

type EfuseRdSysPart2Data2RegT struct {
	Val c.Uint32T
}

/** Type of rd_sys_part2_data3 register
 *  Register 3 of BLOCK10 (system).
 */

type EfuseRdSysPart2Data3RegT struct {
	Val c.Uint32T
}

/** Type of rd_sys_part2_data4 register
 *  Register 4 of BLOCK10 (system).
 */

type EfuseRdSysPart2Data4RegT struct {
	Val c.Uint32T
}

/** Type of rd_sys_part2_data5 register
 *  Register 5 of BLOCK10 (system).
 */

type EfuseRdSysPart2Data5RegT struct {
	Val c.Uint32T
}

/** Type of rd_sys_part2_data6 register
 *  Register 6 of BLOCK10 (system).
 */

type EfuseRdSysPart2Data6RegT struct {
	Val c.Uint32T
}

/** Type of rd_sys_part2_data7 register
 *  Register 7 of BLOCK10 (system).
 */

type EfuseRdSysPart2Data7RegT struct {
	Val c.Uint32T
}

/** Group: Report Register */
/** Type of rd_repeat_err0 register
 *  Programming error record register 0 of BLOCK0.
 */

type EfuseRdRepeatErr0RegT struct {
	Val c.Uint32T
}

/** Type of rd_repeat_err1 register
 *  Programming error record register 1 of BLOCK0.
 */

type EfuseRdRepeatErr1RegT struct {
	Val c.Uint32T
}

/** Type of rd_repeat_err2 register
 *  Programming error record register 2 of BLOCK0.
 */

type EfuseRdRepeatErr2RegT struct {
	Val c.Uint32T
}

/** Type of rd_repeat_err3 register
 *  Programming error record register 3 of BLOCK0.
 */

type EfuseRdRepeatErr3RegT struct {
	Val c.Uint32T
}

/** Type of rd_repeat_err4 register
 *  Programming error record register 4 of BLOCK0.
 */

type EfuseRdRepeatErr4RegT struct {
	Val c.Uint32T
}

/** Type of rd_rs_err0 register
 *  Programming error record register 0 of BLOCK1-10.
 */

type EfuseRdRsErr0RegT struct {
	Val c.Uint32T
}

/** Type of rd_rs_err1 register
 *  Programming error record register 1 of BLOCK1-10.
 */

type EfuseRdRsErr1RegT struct {
	Val c.Uint32T
}

/** Group: Configuration Register */
/** Type of clk register
 *  eFuse clcok configuration register.
 */

type EfuseClkRegT struct {
	Val c.Uint32T
}

/** Type of conf register
 *  eFuse operation mode configuration register
 */

type EfuseConfRegT struct {
	Val c.Uint32T
}

/** Type of cmd register
 *  eFuse command register.
 */

type EfuseCmdRegT struct {
	Val c.Uint32T
}

/** Type of dac_conf register
 *  Controls the eFuse programming voltage.
 */

type EfuseDacConfRegT struct {
	Val c.Uint32T
}

/** Type of rd_tim_conf register
 *  Configures read timing parameters.
 */

type EfuseRdTimConfRegT struct {
	Val c.Uint32T
}

/** Type of wr_tim_conf1 register
 *  Configurarion register 1 of eFuse programming timing parameters.
 */

type EfuseWrTimConf1RegT struct {
	Val c.Uint32T
}

/** Type of wr_tim_conf2 register
 *  Configurarion register 2 of eFuse programming timing parameters.
 */

type EfuseWrTimConf2RegT struct {
	Val c.Uint32T
}

/** Group: Status Register */
/** Type of status register
 *  eFuse status register.
 */

type EfuseStatusRegT struct {
	Val c.Uint32T
}

/** Group: Interrupt Register */
/** Type of int_raw register
 *  eFuse raw interrupt register.
 */

type EfuseIntRawRegT struct {
	Val c.Uint32T
}

/** Type of int_st register
 *  eFuse interrupt status register.
 */

type EfuseIntStRegT struct {
	Val c.Uint32T
}

/** Type of int_ena register
 *  eFuse interrupt enable register.
 */

type EfuseIntEnaRegT struct {
	Val c.Uint32T
}

/** Type of int_clr register
 *  eFuse interrupt clear register.
 */

type EfuseIntClrRegT struct {
	Val c.Uint32T
}

/** Group: Version Register */
/** Type of date register
 *  eFuse version register.
 */

type EfuseDateRegT struct {
	Val c.Uint32T
}

type EfuseDevT struct {
	PgmData0        EfusePgmData0RegT
	PgmData1        EfusePgmData1RegT
	PgmData2        EfusePgmData2RegT
	PgmData3        EfusePgmData3RegT
	PgmData4        EfusePgmData4RegT
	PgmData5        EfusePgmData5RegT
	PgmData6        EfusePgmData6RegT
	PgmData7        EfusePgmData7RegT
	PgmCheckValue0  EfusePgmCheckValue0RegT
	PgmCheckValue1  EfusePgmCheckValue1RegT
	PgmCheckValue2  EfusePgmCheckValue2RegT
	RdWrDis         EfuseRdWrDisRegT
	RdRepeatData0   EfuseRdRepeatData0RegT
	RdRepeatData1   EfuseRdRepeatData1RegT
	RdRepeatData2   EfuseRdRepeatData2RegT
	RdRepeatData3   EfuseRdRepeatData3RegT
	RdRepeatData4   EfuseRdRepeatData4RegT
	RdMacSpiSys0    EfuseRdMacSpiSys0RegT
	RdMacSpiSys1    EfuseRdMacSpiSys1RegT
	RdMacSpiSys2    EfuseRdMacSpiSys2RegT
	RdMacSpiSys3    EfuseRdMacSpiSys3RegT
	RdMacSpiSys4    EfuseRdMacSpiSys4RegT
	RdMacSpiSys5    EfuseRdMacSpiSys5RegT
	RdSysPart1Data0 EfuseRdSysPart1Data0RegT
	RdSysPart1Data1 EfuseRdSysPart1Data1RegT
	RdSysPart1Data2 EfuseRdSysPart1Data2RegT
	RdSysPart1Data3 EfuseRdSysPart1Data3RegT
	RdSysPart1Data4 EfuseRdSysPart1Data4RegT
	RdSysPart1Data5 EfuseRdSysPart1Data5RegT
	RdSysPart1Data6 EfuseRdSysPart1Data6RegT
	RdSysPart1Data7 EfuseRdSysPart1Data7RegT
	RdUsrData0      EfuseRdUsrData0RegT
	RdUsrData1      EfuseRdUsrData1RegT
	RdUsrData2      EfuseRdUsrData2RegT
	RdUsrData3      EfuseRdUsrData3RegT
	RdUsrData4      EfuseRdUsrData4RegT
	RdUsrData5      EfuseRdUsrData5RegT
	RdUsrData6      EfuseRdUsrData6RegT
	RdUsrData7      EfuseRdUsrData7RegT
	RdKey0Data0     EfuseRdKey0Data0RegT
	RdKey0Data1     EfuseRdKey0Data1RegT
	RdKey0Data2     EfuseRdKey0Data2RegT
	RdKey0Data3     EfuseRdKey0Data3RegT
	RdKey0Data4     EfuseRdKey0Data4RegT
	RdKey0Data5     EfuseRdKey0Data5RegT
	RdKey0Data6     EfuseRdKey0Data6RegT
	RdKey0Data7     EfuseRdKey0Data7RegT
	RdKey1Data0     EfuseRdKey1Data0RegT
	RdKey1Data1     EfuseRdKey1Data1RegT
	RdKey1Data2     EfuseRdKey1Data2RegT
	RdKey1Data3     EfuseRdKey1Data3RegT
	RdKey1Data4     EfuseRdKey1Data4RegT
	RdKey1Data5     EfuseRdKey1Data5RegT
	RdKey1Data6     EfuseRdKey1Data6RegT
	RdKey1Data7     EfuseRdKey1Data7RegT
	RdKey2Data0     EfuseRdKey2Data0RegT
	RdKey2Data1     EfuseRdKey2Data1RegT
	RdKey2Data2     EfuseRdKey2Data2RegT
	RdKey2Data3     EfuseRdKey2Data3RegT
	RdKey2Data4     EfuseRdKey2Data4RegT
	RdKey2Data5     EfuseRdKey2Data5RegT
	RdKey2Data6     EfuseRdKey2Data6RegT
	RdKey2Data7     EfuseRdKey2Data7RegT
	RdKey3Data0     EfuseRdKey3Data0RegT
	RdKey3Data1     EfuseRdKey3Data1RegT
	RdKey3Data2     EfuseRdKey3Data2RegT
	RdKey3Data3     EfuseRdKey3Data3RegT
	RdKey3Data4     EfuseRdKey3Data4RegT
	RdKey3Data5     EfuseRdKey3Data5RegT
	RdKey3Data6     EfuseRdKey3Data6RegT
	RdKey3Data7     EfuseRdKey3Data7RegT
	RdKey4Data0     EfuseRdKey4Data0RegT
	RdKey4Data1     EfuseRdKey4Data1RegT
	RdKey4Data2     EfuseRdKey4Data2RegT
	RdKey4Data3     EfuseRdKey4Data3RegT
	RdKey4Data4     EfuseRdKey4Data4RegT
	RdKey4Data5     EfuseRdKey4Data5RegT
	RdKey4Data6     EfuseRdKey4Data6RegT
	RdKey4Data7     EfuseRdKey4Data7RegT
	RdKey5Data0     EfuseRdKey5Data0RegT
	RdKey5Data1     EfuseRdKey5Data1RegT
	RdKey5Data2     EfuseRdKey5Data2RegT
	RdKey5Data3     EfuseRdKey5Data3RegT
	RdKey5Data4     EfuseRdKey5Data4RegT
	RdKey5Data5     EfuseRdKey5Data5RegT
	RdKey5Data6     EfuseRdKey5Data6RegT
	RdKey5Data7     EfuseRdKey5Data7RegT
	RdSysPart2Data0 EfuseRdSysPart2Data0RegT
	RdSysPart2Data1 EfuseRdSysPart2Data1RegT
	RdSysPart2Data2 EfuseRdSysPart2Data2RegT
	RdSysPart2Data3 EfuseRdSysPart2Data3RegT
	RdSysPart2Data4 EfuseRdSysPart2Data4RegT
	RdSysPart2Data5 EfuseRdSysPart2Data5RegT
	RdSysPart2Data6 EfuseRdSysPart2Data6RegT
	RdSysPart2Data7 EfuseRdSysPart2Data7RegT
	RdRepeatErr0    EfuseRdRepeatErr0RegT
	RdRepeatErr1    EfuseRdRepeatErr1RegT
	RdRepeatErr2    EfuseRdRepeatErr2RegT
	RdRepeatErr3    EfuseRdRepeatErr3RegT
	Reserved18c     c.Uint32T
	RdRepeatErr4    EfuseRdRepeatErr4RegT
	Reserved194     [11]c.Uint32T
	RdRsErr0        EfuseRdRsErr0RegT
	RdRsErr1        EfuseRdRsErr1RegT
	Clk             EfuseClkRegT
	Conf            EfuseConfRegT
	Status          EfuseStatusRegT
	Cmd             EfuseCmdRegT
	IntRaw          EfuseIntRawRegT
	IntSt           EfuseIntStRegT
	IntEna          EfuseIntEnaRegT
	IntClr          EfuseIntClrRegT
	DacConf         EfuseDacConfRegT
	RdTimConf       EfuseRdTimConfRegT
	Reserved1f0     c.Uint32T
	WrTimConf1      EfuseWrTimConf1RegT
	WrTimConf2      EfuseWrTimConf2RegT
	Date            EfuseDateRegT
}
