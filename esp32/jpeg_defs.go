package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type JpegMarkerCodeT c.Int

const (
	JPEG_M_SOF0  JpegMarkerCodeT = 65472
	JPEG_M_SOF1  JpegMarkerCodeT = 65473
	JPEG_M_SOF2  JpegMarkerCodeT = 65474
	JPEG_M_SOF3  JpegMarkerCodeT = 65475
	JPEG_M_SOF5  JpegMarkerCodeT = 65477
	JPEG_M_SOF6  JpegMarkerCodeT = 65478
	JPEG_M_SOF7  JpegMarkerCodeT = 65479
	JPEG_M_JPG   JpegMarkerCodeT = 65480
	JPEG_M_SOF9  JpegMarkerCodeT = 65481
	JPEG_M_SOF10 JpegMarkerCodeT = 65482
	JPEG_M_SOF11 JpegMarkerCodeT = 65483
	JPEG_M_SOF13 JpegMarkerCodeT = 65485
	JPEG_M_SOF14 JpegMarkerCodeT = 65486
	JPEG_M_SOF15 JpegMarkerCodeT = 65487
	JPEG_M_DHT   JpegMarkerCodeT = 65476
	JPEG_M_DAC   JpegMarkerCodeT = 65484
	JPEG_M_RST0  JpegMarkerCodeT = 65488
	JPEG_M_RST1  JpegMarkerCodeT = 65489
	JPEG_M_RST2  JpegMarkerCodeT = 65490
	JPEG_M_RST3  JpegMarkerCodeT = 65491
	JPEG_M_RST4  JpegMarkerCodeT = 65492
	JPEG_M_RST5  JpegMarkerCodeT = 65493
	JPEG_M_RST6  JpegMarkerCodeT = 65494
	JPEG_M_RST7  JpegMarkerCodeT = 65495
	JPEG_M_SOI   JpegMarkerCodeT = 65496
	JPEG_M_EOI   JpegMarkerCodeT = 65497
	JPEG_M_SOS   JpegMarkerCodeT = 65498
	JPEG_M_DQT   JpegMarkerCodeT = 65499
	JPEG_M_DNL   JpegMarkerCodeT = 65500
	JPEG_M_DRI   JpegMarkerCodeT = 65501
	JPEG_M_DHP   JpegMarkerCodeT = 65502
	JPEG_M_EXP   JpegMarkerCodeT = 65503
	JPEG_M_APP0  JpegMarkerCodeT = 65504
	JPEG_M_APP1  JpegMarkerCodeT = 65505
	JPEG_M_APP2  JpegMarkerCodeT = 65506
	JPEG_M_APP3  JpegMarkerCodeT = 65507
	JPEG_M_APP4  JpegMarkerCodeT = 65508
	JPEG_M_APP5  JpegMarkerCodeT = 65509
	JPEG_M_APP6  JpegMarkerCodeT = 65510
	JPEG_M_APP7  JpegMarkerCodeT = 65511
	JPEG_M_APP8  JpegMarkerCodeT = 65512
	JPEG_M_APP9  JpegMarkerCodeT = 65513
	JPEG_M_APP10 JpegMarkerCodeT = 65514
	JPEG_M_APP11 JpegMarkerCodeT = 65515
	JPEG_M_APP12 JpegMarkerCodeT = 65516
	JPEG_M_APP13 JpegMarkerCodeT = 65517
	JPEG_M_APP14 JpegMarkerCodeT = 65518
	JPEG_M_APP15 JpegMarkerCodeT = 65519
	JPEG_M_JPG0  JpegMarkerCodeT = 65520
	JPEG_M_JPG13 JpegMarkerCodeT = 65533
	JPEG_M_COM   JpegMarkerCodeT = 65534
	JPEG_M_TEM   JpegMarkerCodeT = 65281
	JPEG_M_INV   JpegMarkerCodeT = 65535
)
