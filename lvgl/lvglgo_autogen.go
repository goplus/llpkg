package lvgl

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const LV_OS_NONE = 0
const LV_OS_PTHREAD = 1
const LV_OS_FREERTOS = 2
const LV_OS_CMSIS_RTOS2 = 3
const LV_OS_RTTHREAD = 4
const LV_OS_WINDOWS = 5
const LV_OS_MQX = 6
const LV_OS_SDL2 = 7
const LV_OS_CUSTOM = 255
const LV_STDLIB_BUILTIN = 0
const LV_STDLIB_CLIB = 1
const LV_STDLIB_MICROPYTHON = 2
const LV_STDLIB_RTTHREAD = 3
const LV_STDLIB_CUSTOM = 255
const LV_DRAW_SW_ASM_NONE = 0
const LV_DRAW_SW_ASM_NEON = 1
const LV_DRAW_SW_ASM_HELIUM = 2
const LV_DRAW_SW_ASM_CUSTOM = 255
const LV_NEMA_HAL_CUSTOM = 0
const LV_NEMA_HAL_STM32 = 1
const LV_LOG_TRACE_MEM = 0
const LV_LOG_TRACE_TIMER = 0
const LV_LOG_TRACE_INDEV = 0
const LV_LOG_TRACE_DISP_REFR = 0
const LV_LOG_TRACE_EVENT = 0
const LV_LOG_TRACE_OBJ_CREATE = 0
const LV_LOG_TRACE_LAYOUT = 0
const LV_LOG_TRACE_ANIM = 0
const LV_WAYLAND_USE_DMABUF = 0
const LV_WAYLAND_WINDOW_DECORATIONS = 0
const LV_WAYLAND_WL_SHELL = 0
const LV_USE_PERF_MONITOR = 0
const LV_USE_MEM_MONITOR = 0
const LV_USE_PERF_MONITOR_LOG_MODE = 0
const LV_USE_LZ4 = 0
const LV_USE_THORVG = 0
const LV_SYMBOL_BULLET = "•"
const LV_SYMBOL_AUDIO = "\uf001"
const LV_SYMBOL_VIDEO = "\uf008"
const LV_SYMBOL_LIST = "\uf00b"
const LV_SYMBOL_OK = "\uf00c"
const LV_SYMBOL_CLOSE = "\uf00d"
const LV_SYMBOL_POWER = "\uf011"
const LV_SYMBOL_SETTINGS = "\uf013"
const LV_SYMBOL_HOME = "\uf015"
const LV_SYMBOL_DOWNLOAD = "\uf019"
const LV_SYMBOL_DRIVE = "\uf01c"
const LV_SYMBOL_REFRESH = "\uf021"
const LV_SYMBOL_MUTE = "\uf026"
const LV_SYMBOL_VOLUME_MID = "\uf027"
const LV_SYMBOL_VOLUME_MAX = "\uf028"
const LV_SYMBOL_IMAGE = "\uf03e"
const LV_SYMBOL_TINT = "\uf043"
const LV_SYMBOL_PREV = "\uf048"
const LV_SYMBOL_PLAY = "\uf04b"
const LV_SYMBOL_PAUSE = "\uf04c"
const LV_SYMBOL_STOP = "\uf04d"
const LV_SYMBOL_NEXT = "\uf051"
const LV_SYMBOL_EJECT = "\uf052"
const LV_SYMBOL_LEFT = "\uf053"
const LV_SYMBOL_RIGHT = "\uf054"
const LV_SYMBOL_PLUS = "\uf067"
const LV_SYMBOL_MINUS = "\uf068"
const LV_SYMBOL_EYE_OPEN = "\uf06e"
const LV_SYMBOL_EYE_CLOSE = "\uf070"
const LV_SYMBOL_WARNING = "\uf071"
const LV_SYMBOL_SHUFFLE = "\uf074"
const LV_SYMBOL_UP = "\uf077"
const LV_SYMBOL_DOWN = "\uf078"
const LV_SYMBOL_LOOP = "\uf079"
const LV_SYMBOL_DIRECTORY = "\uf07b"
const LV_SYMBOL_UPLOAD = "\uf093"
const LV_SYMBOL_CALL = "\uf095"
const LV_SYMBOL_CUT = "\uf0c4"
const LV_SYMBOL_COPY = "\uf0c5"
const LV_SYMBOL_SAVE = "\uf0c7"
const LV_SYMBOL_BARS = "\uf0c9"
const LV_SYMBOL_ENVELOPE = "\uf0e0"
const LV_SYMBOL_CHARGE = "\uf0e7"
const LV_SYMBOL_PASTE = "\uf0ea"
const LV_SYMBOL_BELL = "\uf0f3"
const LV_SYMBOL_KEYBOARD = "\uf11c"
const LV_SYMBOL_GPS = "\uf124"
const LV_SYMBOL_FILE = "\uf15b"
const LV_SYMBOL_WIFI = "\uf1eb"
const LV_SYMBOL_BATTERY_FULL = "\uf240"
const LV_SYMBOL_BATTERY_3 = "\uf241"
const LV_SYMBOL_BATTERY_2 = "\uf242"
const LV_SYMBOL_BATTERY_1 = "\uf243"
const LV_SYMBOL_BATTERY_EMPTY = "\uf244"
const LV_SYMBOL_USB = "\uf287"
const LV_SYMBOL_BLUETOOTH = "\uf293"
const LV_SYMBOL_TRASH = "\uf2ed"
const LV_SYMBOL_EDIT = "\uf304"
const LV_SYMBOL_BACKSPACE = "\uf55a"
const LV_SYMBOL_SD_CARD = "\uf7c2"
const LV_SYMBOL_NEW_LINE = "\uf8a2"
const LV_SYMBOL_DUMMY = "\uf8ff"
const LV_TRIGO_SIN_MAX = 32768
const LV_TRIGO_SHIFT = 15
const LV_BEZIER_VAL_SHIFT = 10
const LV_LOG_LEVEL_TRACE = 0
const LV_LOG_LEVEL_INFO = 1
const LV_LOG_LEVEL_WARN = 2
const LV_LOG_LEVEL_ERROR = 3
const LV_LOG_LEVEL_USER = 4
const LV_LOG_LEVEL_NONE = 5
const LV_LOG_LEVEL_NUM = 5
const LV_COLOR_NATIVE_WITH_ALPHA_SIZE = 3
const LV_OPA_MIN = 2
const LV_OPA_MAX = 253
const LV_STRIDE_AUTO = 0
const LV_NO_TIMER_READY = 0xFFFFFFFF
const LV_ANIM_REPEAT_INFINITE = 0xFFFFFFFF
const LV_ANIM_PLAYTIME_INFINITE = 0xFFFFFFFF
const LV_ANIM_PAUSE_FOREVER = 0xFFFFFFFF
const LV_TXT_ENC_UTF8 = 1
const LV_TXT_ENC_ASCII = 2
const LV_BIDI_LRO = "\u202d"
const LV_BIDI_RLO = "\u202e"
const LV_STYLE_SENTINEL_VALUE = 0xAABBCCDD
const LV_SCALE_NONE = 256
const LV_FS_MAX_FN_LENGTH = 64
const LV_FS_MAX_PATH_LENGTH = 256
const LV_DRAW_UNIT_NONE = 0
const LV_ARRAY_DEFAULT_CAPACITY = 4
const LV_ARRAY_DEFAULT_SHRINK_RATIO = 2
const LV_RADIUS_CIRCLE = 0x7FFF
const LV_MASK_MAX_NUM = 16
const ZERO_MEM_SENTINEL = 0xa1b2c3d4
const LV_INV_BUF_SIZE = 32
const LV_INDEV_VECT_HIST_SIZE = 8
const LV_BUTTONMATRIX_BUTTON_NONE = 0xFFFF
const LV_LABEL_DOT_NUM = 3
const LV_LABEL_POS_LAST = 0xFFFF
const LV_LABEL_DEFAULT_TEXT = "Text"
const LV_SWITCH_KNOB_EXT_AREA_CORRECTION = 2
const LV_TABLE_CELL_NONE = 0xFFFF
const LV_DROPDOWN_POS_LAST = 0xFFFF
const LV_SCALE_LABEL_ROTATE_MATCH_TICKS = 0x100000
const LV_SCALE_LABEL_ROTATE_KEEP_UPRIGHT = 0x80000
const LV_SCALE_ROTATION_ANGLE_MASK = 0x7FFFF
const LV_LED_BRIGHT_MIN = 80
const LV_LED_BRIGHT_MAX = 255
const LV_SPINBOX_MAX_DIGIT_COUNT = 10
const LV_ANIM_TIMELINE_PROGRESS_MAX = 0xFFFF
const LV_ARCLABEL_DOT_NUM = 3
const LV_ARCLABEL_DEFAULT_TEXT = "Arced Text"
const LV_FS_MAX_PATH_LEN = 256
const LV_SDL_MOUSEWHEEL_MODE_ENCODER = 0
const LV_SDL_MOUSEWHEEL_MODE_CROWN = 1
const KEYBOARD_BUFFER_SIZE = 32

type X_silenceGccWarning struct {
	Unused [8]uint8
}
type LvResultT c.Int

const (
	LV_RESULT_INVALID LvResultT = 0
	LV_RESULT_OK      LvResultT = 1
)

type LvUintptrT c.UintptrT
type LvIntptrT c.IntptrT
type LvValuePreciseT c.Int32T

type X_lvObjT struct {
	ClassP                    *LvObjClassT
	Parent                    *LvObjT
	SpecAttr                  *LvObjSpecAttrT
	Styles                    *LvObjStyleT
	UserData                  c.Pointer
	Coords                    LvAreaT
	Flags                     LvObjFlagT
	State                     LvStateT
	LayoutInv                 c.Uint16T
	ReadjustScrollAfterLayout c.Uint16T
	ScrLayoutInv              c.Uint16T
	SkipTrans                 c.Uint16T
	StyleCnt                  c.Uint16T
	HLayout                   c.Uint16T
	WLayout                   c.Uint16T
	IsDeleting                c.Uint16T
}
type LvObjT X_lvObjT

// llgo:type C
type LvScreenCreateCbT func() *LvObjT
type LvStateT c.Uint16T
type LvPartT c.Uint32T
type LvOpaT c.Uint8T
type LvStylePropT c.Uint8T

type X_lvObjClassT struct {
	BaseClass        *LvObjClassT
	ConstructorCb    c.Pointer
	DestructorCb     c.Pointer
	EventCb          c.Pointer
	UserData         c.Pointer
	Name             *c.Char
	WidthDef         c.Int32T
	HeightDef        c.Int32T
	Editable         c.Uint32T
	GroupDef         c.Uint32T
	InstanceSize     c.Uint32T
	ThemeInheritable c.Uint32T
}
type LvObjClassT X_lvObjClassT

type X_lvGroupT struct {
	ObjLl         LvLlT
	ObjFocus      **LvObjT
	FocusCb       LvGroupFocusCbT
	EdgeCb        LvGroupEdgeCbT
	UserData      c.Pointer
	Frozen        c.Uint8T
	Editing       c.Uint8T
	RefocusPolicy c.Uint8T
	Wrap          c.Uint8T
}
type LvGroupT X_lvGroupT

type X_lvDisplayT struct {
	HorRes              c.Int32T
	VerRes              c.Int32T
	PhysicalHorRes      c.Int32T
	PhysicalVerRes      c.Int32T
	OffsetX             c.Int32T
	OffsetY             c.Int32T
	Dpi                 c.Uint32T
	Buf1                *LvDrawBufT
	Buf2                *LvDrawBufT
	Buf3                *LvDrawBufT
	BufAct              *LvDrawBufT
	FlushCb             LvDisplayFlushCbT
	FlushWaitCb         LvDisplayFlushWaitCbT
	Flushing            c.Int
	FlushingLast        c.Int
	LastArea            c.Uint32T
	LastPart            c.Uint32T
	RenderMode          LvDisplayRenderModeT
	Antialiasing        c.Uint32T
	TileCnt             c.Uint32T
	StrideIsAuto        c.Uint32T
	RenderingInProgress c.Uint32T
	ColorFormat         LvColorFormatT
	InvAreas            [32]LvAreaT
	InvAreaJoined       [32]c.Uint8T
	InvP                c.Uint32T
	InvEnCnt            c.Int32T
	SyncAreas           LvLlT
	X_staticBuf1        LvDrawBufT
	X_staticBuf2        LvDrawBufT
	LayerHead           *LvLayerT
	LayerInit           c.Pointer
	LayerDeinit         c.Pointer
	Screens             **LvObjT
	SysLayer            *LvObjT
	TopLayer            *LvObjT
	ActScr              *LvObjT
	BottomLayer         *LvObjT
	PrevScr             *LvObjT
	ScrToLoad           *LvObjT
	ScreenCnt           c.Uint32T
	DrawPrevOverAct     c.Uint8T
	DelPrev             c.Uint8T
	DriverData          c.Pointer
	UserData            c.Pointer
	EventList           LvEventListT
	Rotation            c.Uint32T
	MatrixRotation      c.Uint32T
	Theme               *LvThemeT
	RefrTimer           *LvTimerT
	LastActivityTime    c.Uint32T
	RefreshedArea       LvAreaT
	VsyncCount          c.Uint32T
}
type LvDisplayT X_lvDisplayT

type X_lvLayerT struct {
	DrawBuf        *LvDrawBufT
	DrawTaskHead   *LvDrawTaskT
	Parent         *LvLayerT
	Next           *LvLayerT
	UserData       c.Pointer
	BufArea        LvAreaT
	PhyClipArea    LvAreaT
	X_clipArea     LvAreaT
	PartialYOffset c.Int32T
	Recolor        LvColor32T
	ColorFormat    LvColorFormatT
	AllTasksAdded  bool
	Opa            LvOpaT
}
type LvLayerT X_lvLayerT

type X_lvDrawUnitT struct {
	Next            *LvDrawUnitT
	Name            *c.Char
	Idx             c.Int32T
	DispatchCb      c.Pointer
	EvaluateCb      c.Pointer
	WaitForFinishCb c.Pointer
	DeleteCb        c.Pointer
}
type LvDrawUnitT X_lvDrawUnitT

type X_lvDrawTaskT struct {
	Next                *LvDrawTaskT
	Type                LvDrawTaskTypeT
	Area                LvAreaT
	X_realArea          LvAreaT
	ClipAreaOriginal    LvAreaT
	ClipArea            LvAreaT
	TargetLayer         *LvLayerT
	DrawUnit            *LvDrawUnitT
	State               c.Int
	DrawDsc             c.Pointer
	Opa                 LvOpaT
	PreferredDrawUnitId c.Uint8T
	PreferenceScore     c.Uint8T
}
type LvDrawTaskT X_lvDrawTaskT

type X_lvIndevT struct {
	Type                LvIndevTypeT
	ReadCb              LvIndevReadCbT
	State               LvIndevStateT
	PrevState           LvIndevStateT
	Mode                LvIndevModeT
	LongPrSent          c.Uint8T
	ResetQuery          c.Uint8T
	Enabled             c.Uint8T
	WaitUntilRelease    c.Uint8T
	StopProcessingQuery c.Uint8T
	Timestamp           c.Uint32T
	PrTimestamp         c.Uint32T
	LongprRepTimestamp  c.Uint32T
	DriverData          c.Pointer
	UserData            c.Pointer
	Disp                *LvDisplayT
	ReadTimer           *LvTimerT
	ScrollLimit         c.Uint8T
	ScrollThrow         c.Uint8T
	GestureMinVelocity  c.Uint8T
	GestureLimit        c.Uint8T
	LongPressTime       c.Uint16T
	LongPressRepeatTime c.Uint16T
	RotarySensitivity   c.Int32T
	Pointer             struct {
		ActPoint                LvPointT
		LastPoint               LvPointT
		LastRawPoint            LvPointT
		Vect                    LvPointT
		VectHist                [8]LvPointT
		VectHistTimestamp       [8]c.Uint32T
		VectHistIndex           c.Uint8T
		ScrollSum               LvPointT
		ScrollThrowVect         LvPointT
		ScrollThrowVectOri      LvPointT
		ActObj                  *LvObjT
		LastObj                 *LvObjT
		ScrollObj               *LvObjT
		LastPressed             *LvObjT
		LastHovered             *LvObjT
		ScrollArea              LvAreaT
		GestureSum              LvPointT
		Diff                    c.Int32T
		ShortClickStreak        c.Uint8T
		LastShortClickPoint     LvPointT
		LastShortClickTimestamp c.Uint32T
		ScrollDir               c.Uint8T
		GestureDir              c.Uint8T
		GestureSent             c.Uint8T
		PressMoved              c.Uint8T
		Pressed                 c.Uint8T
	}
	Keypad struct {
		LastState LvIndevStateT
		LastKey   c.Uint32T
	}
	Cursor          *LvObjT
	Group           *LvGroupT
	BtnPoints       *LvPointT
	EventList       LvEventListT
	ScrollThrowAnim *LvAnimT
}
type LvIndevT X_lvIndevT

type X_lvEventT struct {
	CurrentTarget  c.Pointer
	OriginalTarget c.Pointer
	Code           LvEventCodeT
	UserData       c.Pointer
	Param          c.Pointer
	Prev           *LvEventT
	Deleted        c.Uint8T
	StopProcessing c.Uint8T
	StopBubbling   c.Uint8T
}
type LvEventT X_lvEventT

type X_lvTimerT struct {
	Period      c.Uint32T
	LastRun     c.Uint32T
	TimerCb     LvTimerCbT
	UserData    c.Pointer
	RepeatCount c.Int32T
	Paused      c.Uint32T
	AutoDelete  c.Uint32T
}
type LvTimerT X_lvTimerT

type X_lvThemeT struct {
	ApplyCb        LvThemeApplyCbT
	Parent         *LvThemeT
	UserData       c.Pointer
	Disp           *LvDisplayT
	ColorPrimary   LvColorT
	ColorSecondary LvColorT
	FontSmall      *LvFontT
	FontNormal     *LvFontT
	FontLarge      *LvFontT
	Flags          c.Uint32T
}
type LvThemeT X_lvThemeT

type X_lvAnimT struct {
	Var                   c.Pointer
	ExecCb                LvAnimExecXcbT
	CustomExecCb          LvAnimCustomExecCbT
	StartCb               LvAnimStartCbT
	CompletedCb           LvAnimCompletedCbT
	DeletedCb             LvAnimDeletedCbT
	GetValueCb            LvAnimGetValueCbT
	UserData              c.Pointer
	PathCb                LvAnimPathCbT
	StartValue            c.Int32T
	CurrentValue          c.Int32T
	EndValue              c.Int32T
	Duration              c.Int32T
	ActTime               c.Int32T
	ReverseDelay          c.Uint32T
	ReverseDuration       c.Uint32T
	RepeatDelay           c.Uint32T
	RepeatCnt             c.Uint32T
	Parameter             X_lvAnimPathParaT
	LastTimerRun          c.Uint32T
	PauseTime             c.Uint32T
	PauseDuration         c.Uint32T
	IsPaused              c.Uint8T
	ReversePlayInProgress c.Uint8T
	RunRound              c.Uint8T
	StartCbCalled         c.Uint8T
	EarlyApply            c.Uint8T
}
type LvAnimT X_lvAnimT

type X_lvFontT struct {
	GetGlyphDsc        c.Pointer
	GetGlyphBitmap     c.Pointer
	ReleaseGlyph       c.Pointer
	LineHeight         c.Int32T
	BaseLine           c.Int32T
	Subpx              c.Uint8T
	Kerning            c.Uint8T
	StaticBitmap       c.Uint8T
	UnderlinePosition  c.Int8T
	UnderlineThickness c.Int8T
	Dsc                c.Pointer
	Fallback           *LvFontT
	UserData           c.Pointer
}
type LvFontT X_lvFontT

type X_lvFontClassT struct {
	CreateCb  c.Pointer
	DeleteCb  c.Pointer
	DupSrcCb  c.Pointer
	FreeSrcCb c.Pointer
}
type LvFontClassT X_lvFontClassT

type X_lvFontInfoT struct {
	Name       *c.Char
	ClassP     *LvFontClassT
	Size       c.Uint32T
	RenderMode c.Uint32T
	Style      c.Uint32T
	Kerning    LvFontKerningT
}
type LvFontInfoT X_lvFontInfoT

type X_lvFontManagerT struct {
	Unused [8]uint8
}
type LvFontManagerT X_lvFontManagerT

type X_lvImageDecoderT struct {
	InfoCb       LvImageDecoderInfoFT
	OpenCb       LvImageDecoderOpenFT
	GetAreaCb    LvImageDecoderGetAreaCbT
	CloseCb      LvImageDecoderCloseFT
	CustomDrawCb LvImageDecoderCustomDrawT
	Name         *c.Char
	UserData     c.Pointer
}
type LvImageDecoderT X_lvImageDecoderT

type X_lvImageDecoderDscT struct {
	Decoder     *LvImageDecoderT
	Args        LvImageDecoderArgsT
	Src         c.Pointer
	SrcType     LvImageSrcT
	File        LvFsFileT
	Header      LvImageHeaderT
	Decoded     *LvDrawBufT
	Palette     *LvColor32T
	PaletteSize c.Uint32T
	TimeToOpen  c.Uint32T
	ErrorMsg    *c.Char
	Cache       *LvCacheT
	CacheEntry  *LvCacheEntryT
	UserData    c.Pointer
}
type LvImageDecoderDscT X_lvImageDecoderDscT

type X_lvDrawImageDscT struct {
	Base          LvDrawDscBaseT
	Src           c.Pointer
	Header        LvImageHeaderT
	ClipRadius    c.Int32T
	Rotation      c.Int32T
	ScaleX        c.Int32T
	ScaleY        c.Int32T
	SkewX         c.Int32T
	SkewY         c.Int32T
	Pivot         LvPointT
	Recolor       LvColorT
	RecolorOpa    LvOpaT
	Opa           LvOpaT
	BlendMode     LvBlendModeT
	Antialias     c.Uint16T
	Tile          c.Uint16T
	Sup           *LvDrawImageSupT
	ImageArea     LvAreaT
	BitmapMaskSrc *LvImageDscT
}
type LvDrawImageDscT X_lvDrawImageDscT

type X_lvFragmentT struct {
	Unused [8]uint8
}
type LvFragmentT X_lvFragmentT

type X_lvFragmentClassT struct {
	Unused [8]uint8
}
type LvFragmentClassT X_lvFragmentClassT

type X_lvFragmentManagedStatesT struct {
	Unused [8]uint8
}
type LvFragmentManagedStatesT X_lvFragmentManagedStatesT

type X_lvProfilerBuiltinConfigT struct {
	Unused [8]uint8
}
type LvProfilerBuiltinConfigT X_lvProfilerBuiltinConfigT

type X_lvRbNodeT struct {
	Parent *X_lvRbNodeT
	Left   *X_lvRbNodeT
	Right  *X_lvRbNodeT
	Color  LvRbColorT
	Data   c.Pointer
}
type LvRbNodeT X_lvRbNodeT

type X_lvRbT struct {
	Root    *LvRbNodeT
	Compare LvRbCompareT
	Size    c.SizeT
}
type LvRbT X_lvRbT

type X_lvColorFilterDscT struct {
	FilterCb LvColorFilterCbT
	UserData c.Pointer
}
type LvColorFilterDscT X_lvColorFilterDscT

type X_lvEventDscT struct {
	Cb       LvEventCbT
	UserData c.Pointer
	Filter   c.Uint32T
}
type LvEventDscT X_lvEventDscT

type X_lvCacheT struct {
	Clz      *LvCacheClassT
	NodeSize c.Uint32T
	MaxSize  c.Uint32T
	Size     c.Uint32T
	Ops      LvCacheOpsT
	Lock     LvMutexT
	Name     *c.Char
}
type LvCacheT X_lvCacheT

type X_lvCacheEntryT struct {
	Unused [8]uint8
}
type LvCacheEntryT X_lvCacheEntryT

type X_lvFsFileCacheT struct {
	Start        c.Uint32T
	End          c.Uint32T
	FilePosition c.Uint32T
	Buffer       c.Pointer
}
type LvFsFileCacheT X_lvFsFileCacheT

type X_lvFsPathExT struct {
	Path   [4]c.Char
	Buffer c.Pointer
	Size   c.Uint32T
}
type LvFsPathExT X_lvFsPathExT

type X_lvImageDecoderArgsT struct {
	StrideAlign bool
	Premultiply bool
	NoCache     bool
	UseIndexed  bool
	FlushCache  bool
}
type LvImageDecoderArgsT X_lvImageDecoderArgsT

type X_lvImageCacheDataT struct {
	Slot     LvCacheSlotSizeT
	Src      c.Pointer
	SrcType  LvImageSrcT
	Decoded  *LvDrawBufT
	Decoder  *LvImageDecoderT
	UserData c.Pointer
}
type LvImageCacheDataT X_lvImageCacheDataT

type X_lvImageHeaderCacheDataT struct {
	Src     c.Pointer
	SrcType LvImageSrcT
	Header  LvImageHeaderT
	Decoder *LvImageDecoderT
}
type LvImageHeaderCacheDataT X_lvImageHeaderCacheDataT

type X_lvDrawMaskT struct {
	UserData c.Pointer
}
type LvDrawMaskT X_lvDrawMaskT

type X_lvDrawLabelHintT struct {
	LineStart c.Int32T
	Y         c.Int32T
	CoordY    c.Int32T
}
type LvDrawLabelHintT X_lvDrawLabelHintT

type X_lvDrawGlyphDscT struct {
	GlyphData          c.Pointer
	Format             LvFontGlyphFormatT
	LetterCoords       *LvAreaT
	BgCoords           *LvAreaT
	G                  *LvFontGlyphDscT
	Color              LvColorT
	Opa                LvOpaT
	OutlineStrokeColor LvColorT
	OutlineStrokeOpa   LvOpaT
	OutlineStrokeWidth c.Int32T
	Rotation           c.Int32T
	Pivot              LvPointT
	X_drawBuf          *LvDrawBufT
}
type LvDrawGlyphDscT X_lvDrawGlyphDscT

type X_lvDrawImageSupT struct {
	AlphaColor  LvColorT
	Palette     *LvColor32T
	PaletteSize c.Uint32T
}
type LvDrawImageSupT X_lvDrawImageSupT

type X_lvDrawMaskRectDscT struct {
	Base        LvDrawDscBaseT
	Area        LvAreaT
	Radius      c.Int32T
	KeepOutside c.Uint32T
}
type LvDrawMaskRectDscT X_lvDrawMaskRectDscT

type X_lvObjStyleT struct {
	Style      *LvStyleT
	Selector   c.Uint32T
	IsLocal    c.Uint32T
	IsTrans    c.Uint32T
	IsDisabled c.Uint32T
}
type LvObjStyleT X_lvObjStyleT

type X_lvObjStyleTransitionDscT struct {
	Time     c.Uint16T
	Delay    c.Uint16T
	Selector LvStyleSelectorT
	Prop     LvStylePropT
	PathCb   LvAnimPathCbT
	UserData c.Pointer
}
type LvObjStyleTransitionDscT X_lvObjStyleTransitionDscT

type X_lvHitTestInfoT struct {
	Point *LvPointT
	Res   bool
}
type LvHitTestInfoT X_lvHitTestInfoT

type X_lvCoverCheckInfoT struct {
	Res  LvCoverResT
	Area *LvAreaT
}
type LvCoverCheckInfoT X_lvCoverCheckInfoT

type X_lvObjSpecAttrT struct {
	Children      **LvObjT
	GroupP        *LvGroupT
	EventList     LvEventListT
	Scroll        LvPointT
	ExtClickPad   c.Int32T
	ExtDrawSize   c.Int32T
	ChildCnt      c.Uint16T
	ScrollbarMode c.Uint16T
	ScrollSnapX   c.Uint16T
	ScrollSnapY   c.Uint16T
	ScrollDir     c.Uint16T
	LayerType     c.Uint16T
	NameStatic    c.Uint16T
}
type LvObjSpecAttrT X_lvObjSpecAttrT

type X_lvImageT struct {
	Obj           LvObjT
	Src           c.Pointer
	BitmapMaskSrc *LvImageDscT
	Offset        LvPointT
	W             c.Int32T
	H             c.Int32T
	Rotation      c.Uint32T
	ScaleX        c.Uint32T
	ScaleY        c.Uint32T
	Pivot         LvPointT
	SrcType       c.Uint32T
	Cf            c.Uint32T
	Antialias     c.Uint32T
	Align         c.Uint32T
	BlendMode     c.Uint32T
}
type LvImageT X_lvImageT

type X_lvAnimimgT struct {
	Img      LvImageT
	Anim     LvAnimT
	Dsc      *c.Pointer
	PicCount c.Int8T
}
type LvAnimimgT X_lvAnimimgT

type X_lvArcT struct {
	Obj             LvObjT
	Rotation        c.Int32T
	IndicAngleStart LvValuePreciseT
	IndicAngleEnd   LvValuePreciseT
	BgAngleStart    LvValuePreciseT
	BgAngleEnd      LvValuePreciseT
	Value           c.Int32T
	MinValue        c.Int32T
	MaxValue        c.Int32T
	Dragging        c.Uint32T
	Type            c.Uint32T
	MinClose        c.Uint32T
	InOut           c.Uint32T
	ChgRate         c.Uint32T
	LastTick        c.Uint32T
	LastAngle       LvValuePreciseT
	KnobOffset      c.Int16T
}
type LvArcT X_lvArcT

type X_lvArclabelT struct {
	Unused [8]uint8
}
type LvArclabelT X_lvArclabelT

type X_lvLabelT struct {
	Obj              LvObjT
	Text             *c.Char
	Dot              [4]c.Char
	DotBegin         c.Uint32T
	Hint             LvDrawLabelHintT
	SelStart         c.Uint32T
	SelEnd           c.Uint32T
	SizeCache        LvPointT
	Offset           LvPointT
	LongMode         LvLabelLongModeT
	StaticTxt        c.Uint8T
	Recolor          c.Uint8T
	Expand           c.Uint8T
	InvalidSizeCache c.Uint8T
	TextSize         LvPointT
}
type LvLabelT X_lvLabelT

type X_lvBarAnimT struct {
	Bar       *LvObjT
	AnimStart c.Int32T
	AnimEnd   c.Int32T
	AnimState c.Int32T
}
type LvBarAnimT X_lvBarAnimT

type X_lvBarT struct {
	Obj            LvObjT
	CurValue       c.Int32T
	MinValue       c.Int32T
	MaxValue       c.Int32T
	StartValue     c.Int32T
	IndicArea      LvAreaT
	ValReversed    bool
	CurValueAnim   LvBarAnimT
	StartValueAnim LvBarAnimT
	Mode           LvBarModeT
	Orientation    LvBarOrientationT
}
type LvBarT X_lvBarT

type X_lvButtonT struct {
	Obj LvObjT
}
type LvButtonT X_lvButtonT

type X_lvButtonmatrixT struct {
	Obj         LvObjT
	MapP        **c.Char
	ButtonAreas *LvAreaT
	CtrlBits    *LvButtonmatrixCtrlT
	BtnCnt      c.Uint32T
	RowCnt      c.Uint32T
	BtnIdSel    c.Uint32T
	OneCheck    c.Uint32T
	AutoFreeMap c.Uint32T
}
type LvButtonmatrixT X_lvButtonmatrixT

type X_lvCalendarT struct {
	Obj                 LvObjT
	Btnm                *LvObjT
	Today               LvCalendarDateT
	ShowedDate          LvCalendarDateT
	HighlightedDates    *LvCalendarDateT
	HighlightedDatesNum c.SizeT
	Map                 [56]*c.Char
	UseChineseCalendar  bool
	Nums                [42][20]c.Char
}
type LvCalendarT X_lvCalendarT

type X_lvCanvasT struct {
	Img       LvImageT
	DrawBuf   *LvDrawBufT
	StaticBuf LvDrawBufT
}
type LvCanvasT X_lvCanvasT

type X_lvChartSeriesT struct {
	XPoints         *c.Int32T
	YPoints         *c.Int32T
	Color           LvColorT
	StartPoint      c.Uint32T
	Hidden          c.Uint32T
	XExtBufAssigned c.Uint32T
	YExtBufAssigned c.Uint32T
	XAxisSec        c.Uint32T
	YAxisSec        c.Uint32T
}
type LvChartSeriesT X_lvChartSeriesT

type X_lvChartCursorT struct {
	Pos     LvPointT
	PointId c.Int32T
	Color   LvColorT
	Ser     *LvChartSeriesT
	Dir     LvDirT
	PosSet  c.Uint32T
}
type LvChartCursorT X_lvChartCursorT

type X_lvChartT struct {
	Obj            LvObjT
	SeriesLl       LvLlT
	CursorLl       LvLlT
	Ymin           [2]c.Int32T
	Ymax           [2]c.Int32T
	Xmin           [2]c.Int32T
	Xmax           [2]c.Int32T
	PressedPointId c.Int32T
	HdivCnt        c.Uint32T
	VdivCnt        c.Uint32T
	PointCnt       c.Uint32T
	Type           LvChartTypeT
	UpdateMode     LvChartUpdateModeT
}
type LvChartT X_lvChartT

type X_lvCheckboxT struct {
	Obj       LvObjT
	Txt       *c.Char
	StaticTxt c.Uint32T
}
type LvCheckboxT X_lvCheckboxT

type X_lvDropdownT struct {
	Obj               LvObjT
	List              *LvObjT
	Text              *c.Char
	Symbol            c.Pointer
	Options           *c.Char
	OptionCnt         c.Uint32T
	SelOptId          c.Uint32T
	SelOptIdOrig      c.Uint32T
	PrOptId           c.Uint32T
	Dir               c.Uint8T
	StaticTxt         c.Uint8T
	SelectedHighlight c.Uint8T
}
type LvDropdownT X_lvDropdownT

type X_lvDropdownListT struct {
	Obj      LvObjT
	Dropdown *LvObjT
}
type LvDropdownListT X_lvDropdownListT

type X_lvImagebuttonSrcInfoT struct {
	ImgSrc c.Pointer
	Header LvImageHeaderT
}
type LvImagebuttonSrcInfoT X_lvImagebuttonSrcInfoT

type X_lvImagebuttonT struct {
	Obj      LvObjT
	SrcMid   [6]LvImagebuttonSrcInfoT
	SrcLeft  [6]LvImagebuttonSrcInfoT
	SrcRight [6]LvImagebuttonSrcInfoT
}
type LvImagebuttonT X_lvImagebuttonT

type X_lvKeyboardT struct {
	Btnm     LvButtonmatrixT
	Ta       *LvObjT
	Mode     LvKeyboardModeT
	Popovers c.Uint8T
}
type LvKeyboardT X_lvKeyboardT

type X_lvLedT struct {
	Obj    LvObjT
	Color  LvColorT
	Bright c.Uint8T
}
type LvLedT X_lvLedT

type X_lvLineT struct {
	Obj        LvObjT
	PointArray struct {
		Constant *LvPointPreciseT
	}
	PointNum            c.Uint32T
	YInv                c.Uint32T
	PointArrayIsMutable c.Uint32T
}
type LvLineT X_lvLineT

type X_lvMenuLoadPageEventDataT struct {
	Menu *LvObjT
	Page *LvObjT
}
type LvMenuLoadPageEventDataT X_lvMenuLoadPageEventDataT

type X_lvMenuHistoryT struct {
	Page *LvObjT
}
type LvMenuHistoryT X_lvMenuHistoryT

type X_lvMenuT struct {
	Obj                  LvObjT
	Storage              *LvObjT
	Main                 *LvObjT
	MainPage             *LvObjT
	MainHeader           *LvObjT
	MainHeaderBackBtn    *LvObjT
	MainHeaderTitle      *LvObjT
	Sidebar              *LvObjT
	SidebarPage          *LvObjT
	SidebarHeader        *LvObjT
	SidebarHeaderBackBtn *LvObjT
	SidebarHeaderTitle   *LvObjT
	SelectedTab          *LvObjT
	HistoryLl            LvLlT
	CurDepth             c.Uint8T
	PrevDepth            c.Uint8T
	SidebarGenerated     c.Uint8T
	ModeHeader           LvMenuModeHeaderT
	ModeRootBackBtn      LvMenuModeRootBackButtonT
}
type LvMenuT X_lvMenuT

type X_lvMenuPageT struct {
	Obj         LvObjT
	Title       *c.Char
	StaticTitle bool
}
type LvMenuPageT X_lvMenuPageT

type X_lvMsgboxT struct {
	Obj     LvObjT
	Header  *LvObjT
	Content *LvObjT
	Footer  *LvObjT
	Title   *LvObjT
}
type LvMsgboxT X_lvMsgboxT

type X_lvRollerT struct {
	Obj         LvObjT
	OptionCnt   c.Uint32T
	SelOptId    c.Uint32T
	SelOptIdOri c.Uint32T
	InfPageCnt  c.Uint32T
	Mode        LvRollerModeT
	Moved       c.Uint32T
}
type LvRollerT X_lvRollerT

type X_lvScaleSectionT struct {
	MainStyle               *LvStyleT
	IndicatorStyle          *LvStyleT
	ItemsStyle              *LvStyleT
	RangeMin                c.Int32T
	RangeMax                c.Int32T
	FirstTickIdxInSection   c.Uint32T
	LastTickIdxInSection    c.Uint32T
	FirstTickInSectionWidth c.Int32T
	LastTickInSectionWidth  c.Int32T
	FirstTickInSection      LvPointT
	LastTickInSection       LvPointT
	FirstTickIdxIsMajor     c.Uint32T
	LastTickIdxIsMajor      c.Uint32T
}
type LvScaleSectionT X_lvScaleSectionT

type X_lvScaleT struct {
	Obj            LvObjT
	SectionLl      LvLlT
	TxtSrc         **c.Char
	Mode           LvScaleModeT
	RangeMin       c.Int32T
	RangeMax       c.Int32T
	TotalTickCount c.Uint32T
	MajorTickEvery c.Uint32T
	LabelEnabled   c.Uint32T
	PostDraw       c.Uint32T
	DrawTicksOnTop c.Uint32T
	AngleRange     c.Uint32T
	Rotation       c.Int32T
	CustomLabelCnt c.Int32T
	LastTickWidth  c.Int32T
	FirstTickWidth c.Int32T
}
type LvScaleT X_lvScaleT

type X_lvSliderT struct {
	Bar           LvBarT
	LeftKnobArea  LvAreaT
	RightKnobArea LvAreaT
	PressedPoint  LvPointT
	ValueToSet    *c.Int32T
	Dragging      c.Uint8T
	LeftKnobFocus c.Uint8T
}
type LvSliderT X_lvSliderT

type X_lvSpanT struct {
	Txt            *c.Char
	Style          LvStyleT
	StaticFlag     c.Uint32T
	TrailingPos    LvPointT
	TrailingHeight c.Int32T
}
type LvSpanT X_lvSpanT

type X_lvSpangroupT struct {
	Obj      LvObjT
	Lines    c.Int32T
	Indent   c.Int32T
	CacheW   c.Int32T
	CacheH   c.Int32T
	ChildLl  LvLlT
	Overflow c.Uint32T
	Refresh  c.Uint32T
}
type LvSpangroupT X_lvSpangroupT

type X_lvTextareaT struct {
	Obj            LvObjT
	Label          *LvObjT
	PlaceholderTxt *c.Char
	PwdTmp         *c.Char
	PwdBullet      *c.Char
	AcceptedChars  *c.Char
	MaxLength      c.Uint32T
	PwdShowTime    c.Uint32T
	Cursor         struct {
		ValidX     c.Int32T
		Pos        c.Uint32T
		Area       LvAreaT
		TxtBytePos c.Uint32T
		Show       c.Uint8T
		ClickPos   c.Uint8T
	}
	SelStart      c.Uint32T
	SelEnd        c.Uint32T
	TextSelInProg c.Uint8T
	TextSelEn     c.Uint8T
	PwdMode       c.Uint8T
	OneLine       c.Uint8T
}
type LvTextareaT X_lvTextareaT

type X_lvSpinboxT struct {
	Ta           LvTextareaT
	Value        c.Int32T
	RangeMax     c.Int32T
	RangeMin     c.Int32T
	Step         c.Int32T
	DigitCount   c.Uint32T
	DecPointPos  c.Uint32T
	Rollover     c.Uint32T
	DigitStepDir c.Uint32T
}
type LvSpinboxT X_lvSpinboxT

type X_lvSwitchT struct {
	Obj         LvObjT
	AnimState   c.Int32T
	Orientation LvSwitchOrientationT
}
type LvSwitchT X_lvSwitchT

type X_lvTableCellT struct {
	Ctrl     LvTableCellCtrlT
	UserData c.Pointer
	Txt      [1]c.Char
}
type LvTableCellT X_lvTableCellT

type X_lvTableT struct {
	Obj      LvObjT
	ColCnt   c.Uint32T
	RowCnt   c.Uint32T
	CellData **LvTableCellT
	RowH     *c.Int32T
	ColW     *c.Int32T
	ColAct   c.Uint32T
	RowAct   c.Uint32T
}
type LvTableT X_lvTableT

type X_lvTabviewT struct {
	Obj    LvObjT
	TabCur c.Uint32T
	TabPos LvDirT
}
type LvTabviewT X_lvTabviewT

type X_lvTileviewT struct {
	Obj     LvObjT
	TileAct *LvObjT
}
type LvTileviewT X_lvTileviewT

type X_lvTileviewTileT struct {
	Obj LvObjT
	Dir LvDirT
}
type LvTileviewTileT X_lvTileviewTileT

type X_lvWinT struct {
	Obj LvObjT
}
type LvWinT X_lvWinT

type X_lv3dtextureT struct {
	Unused [8]uint8
}
type Lv3dtextureT X_lv3dtextureT

type X_lvObserverT struct {
	Subject          *LvSubjectT
	Cb               LvObserverCbT
	Target           c.Pointer
	UserData         c.Pointer
	AutoFreeUserData c.Uint32T
	Notified         c.Uint32T
	ForObj           c.Uint32T
}
type LvObserverT X_lvObserverT

type X_lvMonkeyConfigT struct {
	Unused [8]uint8
}
type LvMonkeyConfigT X_lvMonkeyConfigT

type X_lvImePinyinT struct {
	Unused [8]uint8
}
type LvImePinyinT X_lvImePinyinT

type X_lvFileExplorerT struct {
	Unused [8]uint8
}
type LvFileExplorerT X_lvFileExplorerT

type X_lvBarcodeT struct {
	Unused [8]uint8
}
type LvBarcodeT X_lvBarcodeT

type X_lvGifT struct {
	Unused [8]uint8
}
type LvGifT X_lvGifT

type X_lvQrcodeT struct {
	Unused [8]uint8
}
type LvQrcodeT X_lvQrcodeT

type X_lvFreetypeOutlineVectorT struct {
	Unused [8]uint8
}
type LvFreetypeOutlineVectorT X_lvFreetypeOutlineVectorT

type X_lvFreetypeOutlineEventParamT struct {
	Unused [8]uint8
}
type LvFreetypeOutlineEventParamT X_lvFreetypeOutlineEventParamT

type X_lvFpointT struct {
	Unused [8]uint8
}
type LvFpointT X_lvFpointT

type X_lvMatrixT struct {
	Unused [8]uint8
}
type LvMatrixT X_lvMatrixT

type X_lvVectorPathT struct {
	Unused [8]uint8
}
type LvVectorPathT X_lvVectorPathT

type X_lvVectorGradientT struct {
	Unused [8]uint8
}
type LvVectorGradientT X_lvVectorGradientT

type X_lvVectorFillDscT struct {
	Unused [8]uint8
}
type LvVectorFillDscT X_lvVectorFillDscT

type X_lvVectorStrokeDscT struct {
	Unused [8]uint8
}
type LvVectorStrokeDscT X_lvVectorStrokeDscT

type X_lvVectorDrawDscT struct {
	Unused [8]uint8
}
type LvVectorDrawDscT X_lvVectorDrawDscT

type X_lvDrawVectorTaskDscT struct {
	Unused [8]uint8
}
type LvDrawVectorTaskDscT X_lvDrawVectorTaskDscT

type X_lvVectorDscT struct {
	Unused [8]uint8
}
type LvVectorDscT X_lvVectorDscT

type X_lvXkbT struct {
	Unused [8]uint8
}
type LvXkbT X_lvXkbT

type X_lvLibinputEventT struct {
	Unused [8]uint8
}
type LvLibinputEventT X_lvLibinputEventT

type X_lvLibinputT struct {
	Unused [8]uint8
}
type LvLibinputT X_lvLibinputT

type X_lvDrawSwUnitT struct {
	BaseUnit LvDrawUnitT
	TaskAct  *LvDrawTaskT
}
type LvDrawSwUnitT X_lvDrawSwUnitT

type X_lvDrawSwMaskCommonDscT struct {
	Cb   LvDrawSwMaskXcbT
	Type LvDrawSwMaskTypeT
}
type LvDrawSwMaskCommonDscT X_lvDrawSwMaskCommonDscT

type X_lvDrawSwMaskLineParamT struct {
	Dsc LvDrawSwMaskCommonDscT
	Cfg struct {
		P1   LvPointT
		P2   LvPointT
		Side LvDrawSwMaskLineSideT
	}
	Origo   LvPointT
	XySteep c.Int32T
	YxSteep c.Int32T
	Steep   c.Int32T
	Spx     c.Int32T
	Flat    c.Uint8T
	Inv     c.Uint8T
}
type LvDrawSwMaskLineParamT X_lvDrawSwMaskLineParamT

type X_lvDrawSwMaskAngleParamT struct {
	Dsc LvDrawSwMaskCommonDscT
	Cfg struct {
		VertexP    LvPointT
		StartAngle c.Int32T
		EndAngle   c.Int32T
	}
	StartLine LvDrawSwMaskLineParamT
	EndLine   LvDrawSwMaskLineParamT
	DeltaDeg  c.Uint16T
}
type LvDrawSwMaskAngleParamT X_lvDrawSwMaskAngleParamT

type X_lvDrawSwMaskRadiusParamT struct {
	Dsc LvDrawSwMaskCommonDscT
	Cfg struct {
		Rect   LvAreaT
		Radius c.Int32T
		Outer  c.Uint8T
	}
	Circle *LvDrawSwMaskRadiusCircleDscT
}
type LvDrawSwMaskRadiusParamT X_lvDrawSwMaskRadiusParamT

type X_lvDrawSwMaskFadeParamT struct {
	Dsc LvDrawSwMaskCommonDscT
	Cfg struct {
		Coords    LvAreaT
		YTop      c.Int32T
		YBottom   c.Int32T
		OpaTop    LvOpaT
		OpaBottom LvOpaT
	}
}
type LvDrawSwMaskFadeParamT X_lvDrawSwMaskFadeParamT

type X_lvDrawSwMaskMapParamT struct {
	Dsc LvDrawSwMaskCommonDscT
	Cfg struct {
		Coords LvAreaT
		Map    *LvOpaT
	}
}
type LvDrawSwMaskMapParamT X_lvDrawSwMaskMapParamT

type X_lvDrawSwBlendDscT struct {
	BlendArea      *LvAreaT
	SrcBuf         c.Pointer
	SrcStride      c.Uint32T
	SrcColorFormat LvColorFormatT
	SrcArea        *LvAreaT
	Opa            LvOpaT
	Color          LvColorT
	MaskBuf        *LvOpaT
	MaskRes        LvDrawSwMaskResT
	MaskArea       *LvAreaT
	MaskStride     c.Int32T
	BlendMode      LvBlendModeT
}
type LvDrawSwBlendDscT X_lvDrawSwBlendDscT

type X_lvDrawSwBlendFillDscT struct {
	DestBuf      c.Pointer
	DestW        c.Int32T
	DestH        c.Int32T
	DestStride   c.Int32T
	MaskBuf      *LvOpaT
	MaskStride   c.Int32T
	Color        LvColorT
	Opa          LvOpaT
	RelativeArea LvAreaT
}
type LvDrawSwBlendFillDscT X_lvDrawSwBlendFillDscT

type X_lvDrawSwBlendImageDscT struct {
	DestBuf        c.Pointer
	DestW          c.Int32T
	DestH          c.Int32T
	DestStride     c.Int32T
	MaskBuf        *LvOpaT
	MaskStride     c.Int32T
	SrcBuf         c.Pointer
	SrcStride      c.Int32T
	SrcColorFormat LvColorFormatT
	Opa            LvOpaT
	BlendMode      LvBlendModeT
	RelativeArea   LvAreaT
	SrcArea        LvAreaT
}
type LvDrawSwBlendImageDscT X_lvDrawSwBlendImageDscT

type X_lvDrawBufHandlersT struct {
	BufMallocCb       LvDrawBufMallocCb
	BufFreeCb         LvDrawBufFreeCb
	AlignPointerCb    LvDrawBufAlignCb
	InvalidateCacheCb LvDrawBufCacheOperationCb
	FlushCacheCb      LvDrawBufCacheOperationCb
	WidthToStrideCb   LvDrawBufWidthToStrideCb
}
type LvDrawBufHandlersT X_lvDrawBufHandlersT

type X_lvRlottieT struct {
	Unused [8]uint8
}
type LvRlottieT X_lvRlottieT

type X_lvFfmpegPlayerT struct {
	Unused [8]uint8
}
type LvFfmpegPlayerT X_lvFfmpegPlayerT

type X_lvGlfwWindowT struct {
	Unused [8]uint8
}
type LvGlfwWindowT X_lvGlfwWindowT

type X_lvGlfwTextureT struct {
	Unused [8]uint8
}
type LvGlfwTextureT X_lvGlfwTextureT
type LvPropIdT c.Uint32T

type X_lvArrayT struct {
	Data        *c.Uint8T
	Size        c.Uint32T
	Capacity    c.Uint32T
	ElementSize c.Uint32T
	InnerAlloc  bool
}
type LvArrayT X_lvArrayT

type X_lvIterT struct {
	Unused [8]uint8
}
type LvIterT X_lvIterT

type X_lvCircleBufT struct {
	Unused [8]uint8
}
type LvCircleBufT X_lvCircleBufT

type X_lvDrawBufT struct {
	Header        LvImageHeaderT
	DataSize      c.Uint32T
	Data          *c.Uint8T
	UnalignedData c.Pointer
	Handlers      *LvDrawBufHandlersT
}
type LvDrawBufT X_lvDrawBufT

type X_lvXmlComponentScopeT struct {
	Unused [8]uint8
}
type LvXmlComponentScopeT X_lvXmlComponentScopeT

type X_lvXmlParserStateT struct {
	Unused [8]uint8
}
type LvXmlParserStateT X_lvXmlParserStateT
type X_lvStrSymbolIdT c.Int

const (
	LV_STR_SYMBOL_BULLET        X_lvStrSymbolIdT = 0
	LV_STR_SYMBOL_AUDIO         X_lvStrSymbolIdT = 1
	LV_STR_SYMBOL_VIDEO         X_lvStrSymbolIdT = 2
	LV_STR_SYMBOL_LIST          X_lvStrSymbolIdT = 3
	LV_STR_SYMBOL_OK            X_lvStrSymbolIdT = 4
	LV_STR_SYMBOL_CLOSE         X_lvStrSymbolIdT = 5
	LV_STR_SYMBOL_POWER         X_lvStrSymbolIdT = 6
	LV_STR_SYMBOL_SETTINGS      X_lvStrSymbolIdT = 7
	LV_STR_SYMBOL_HOME          X_lvStrSymbolIdT = 8
	LV_STR_SYMBOL_DOWNLOAD      X_lvStrSymbolIdT = 9
	LV_STR_SYMBOL_DRIVE         X_lvStrSymbolIdT = 10
	LV_STR_SYMBOL_REFRESH       X_lvStrSymbolIdT = 11
	LV_STR_SYMBOL_MUTE          X_lvStrSymbolIdT = 12
	LV_STR_SYMBOL_VOLUME_MID    X_lvStrSymbolIdT = 13
	LV_STR_SYMBOL_VOLUME_MAX    X_lvStrSymbolIdT = 14
	LV_STR_SYMBOL_IMAGE         X_lvStrSymbolIdT = 15
	LV_STR_SYMBOL_TINT          X_lvStrSymbolIdT = 16
	LV_STR_SYMBOL_PREV          X_lvStrSymbolIdT = 17
	LV_STR_SYMBOL_PLAY          X_lvStrSymbolIdT = 18
	LV_STR_SYMBOL_PAUSE         X_lvStrSymbolIdT = 19
	LV_STR_SYMBOL_STOP          X_lvStrSymbolIdT = 20
	LV_STR_SYMBOL_NEXT          X_lvStrSymbolIdT = 21
	LV_STR_SYMBOL_EJECT         X_lvStrSymbolIdT = 22
	LV_STR_SYMBOL_LEFT          X_lvStrSymbolIdT = 23
	LV_STR_SYMBOL_RIGHT         X_lvStrSymbolIdT = 24
	LV_STR_SYMBOL_PLUS          X_lvStrSymbolIdT = 25
	LV_STR_SYMBOL_MINUS         X_lvStrSymbolIdT = 26
	LV_STR_SYMBOL_EYE_OPEN      X_lvStrSymbolIdT = 27
	LV_STR_SYMBOL_EYE_CLOSE     X_lvStrSymbolIdT = 28
	LV_STR_SYMBOL_WARNING       X_lvStrSymbolIdT = 29
	LV_STR_SYMBOL_SHUFFLE       X_lvStrSymbolIdT = 30
	LV_STR_SYMBOL_UP            X_lvStrSymbolIdT = 31
	LV_STR_SYMBOL_DOWN          X_lvStrSymbolIdT = 32
	LV_STR_SYMBOL_LOOP          X_lvStrSymbolIdT = 33
	LV_STR_SYMBOL_DIRECTORY     X_lvStrSymbolIdT = 34
	LV_STR_SYMBOL_UPLOAD        X_lvStrSymbolIdT = 35
	LV_STR_SYMBOL_CALL          X_lvStrSymbolIdT = 36
	LV_STR_SYMBOL_CUT           X_lvStrSymbolIdT = 37
	LV_STR_SYMBOL_COPY          X_lvStrSymbolIdT = 38
	LV_STR_SYMBOL_SAVE          X_lvStrSymbolIdT = 39
	LV_STR_SYMBOL_BARS          X_lvStrSymbolIdT = 40
	LV_STR_SYMBOL_ENVELOPE      X_lvStrSymbolIdT = 41
	LV_STR_SYMBOL_CHARGE        X_lvStrSymbolIdT = 42
	LV_STR_SYMBOL_PASTE         X_lvStrSymbolIdT = 43
	LV_STR_SYMBOL_BELL          X_lvStrSymbolIdT = 44
	LV_STR_SYMBOL_KEYBOARD      X_lvStrSymbolIdT = 45
	LV_STR_SYMBOL_GPS           X_lvStrSymbolIdT = 46
	LV_STR_SYMBOL_FILE          X_lvStrSymbolIdT = 47
	LV_STR_SYMBOL_WIFI          X_lvStrSymbolIdT = 48
	LV_STR_SYMBOL_BATTERY_FULL  X_lvStrSymbolIdT = 49
	LV_STR_SYMBOL_BATTERY_3     X_lvStrSymbolIdT = 50
	LV_STR_SYMBOL_BATTERY_2     X_lvStrSymbolIdT = 51
	LV_STR_SYMBOL_BATTERY_1     X_lvStrSymbolIdT = 52
	LV_STR_SYMBOL_BATTERY_EMPTY X_lvStrSymbolIdT = 53
	LV_STR_SYMBOL_USB           X_lvStrSymbolIdT = 54
	LV_STR_SYMBOL_BLUETOOTH     X_lvStrSymbolIdT = 55
	LV_STR_SYMBOL_TRASH         X_lvStrSymbolIdT = 56
	LV_STR_SYMBOL_EDIT          X_lvStrSymbolIdT = 57
	LV_STR_SYMBOL_BACKSPACE     X_lvStrSymbolIdT = 58
	LV_STR_SYMBOL_SD_CARD       X_lvStrSymbolIdT = 59
	LV_STR_SYMBOL_NEW_LINE      X_lvStrSymbolIdT = 60
	LV_STR_SYMBOL_DUMMY         X_lvStrSymbolIdT = 61
)

/**********************
 *      TYPEDEFS
 **********************/

type LvSqrtResT struct {
	I c.Uint16T
	F c.Uint16T
}

//go:linkname LvTrigoSin C.lv_trigo_sin
func LvTrigoSin(angle c.Int16T) c.Int32T

//go:linkname LvTrigoCos C.lv_trigo_cos
func LvTrigoCos(angle c.Int16T) c.Int32T

/**
 * Calculate the y value of cubic-bezier(x1, y1, x2, y2) function as specified x.
 * @param x time in range of [0..LV_BEZIER_VAL_MAX]
 * @param x1 x of control point 1 in range of [0..LV_BEZIER_VAL_MAX]
 * @param y1 y of control point 1 in range of [0..LV_BEZIER_VAL_MAX]
 * @param x2 x of control point 2 in range of [0..LV_BEZIER_VAL_MAX]
 * @param y2 y of control point 2 in range of [0..LV_BEZIER_VAL_MAX]
 * @return the value calculated
 */
//go:linkname LvCubicBezier C.lv_cubic_bezier
func LvCubicBezier(x c.Int32T, x1 c.Int32T, y1 c.Int32T, x2 c.Int32T, y2 c.Int32T) c.Int32T

/**
 * Calculate a value of a Cubic Bezier function.
 * @param t time in range of [0..LV_BEZIER_VAL_MAX]
 * @param u0 must be 0
 * @param u1 control value 1 values in range of [0..LV_BEZIER_VAL_MAX]
 * @param u2 control value 2 in range of [0..LV_BEZIER_VAL_MAX]
 * @param u3 must be LV_BEZIER_VAL_MAX
 * @return the value calculated from the given parameters in range of [0..LV_BEZIER_VAL_MAX]
 */
//go:linkname LvBezier3 C.lv_bezier3
func LvBezier3(t c.Int32T, u0 c.Int32T, u1 c.Uint32T, u2 c.Int32T, u3 c.Int32T) c.Int32T

/**
 * Calculate the atan2 of a vector.
 * @param x
 * @param y
 * @return the angle in degree calculated from the given parameters in range of [0..360]
 */
//go:linkname LvAtan2 C.lv_atan2
func LvAtan2(x c.Int, y c.Int) c.Uint16T

//go:linkname LvSqrt C.lv_sqrt
func LvSqrt(x c.Uint32T, q *LvSqrtResT, mask c.Uint32T)

//go:linkname LvSqrt32 C.lv_sqrt32
func LvSqrt32(x c.Uint32T) c.Int32T

/**
 * Calculate the integer exponents.
 * @param base
 * @param exp
 * @return base raised to the power exponent
 */
//go:linkname LvPow C.lv_pow
func LvPow(base c.Int64T, exp c.Int8T) c.Int64T

/**
 * Get the mapped of a number given an input and output range
 * @param x integer which mapped value should be calculated
 * @param min_in min input range
 * @param max_in max input range
 * @param min_out max output range
 * @param max_out max output range
 * @return the mapped number
 */
//go:linkname LvMap C.lv_map
func LvMap(x c.Int32T, min_in c.Int32T, max_in c.Int32T, min_out c.Int32T, max_out c.Int32T) c.Int32T

/**
 * Set the seed of the pseudo random number generator
 * @param seed a number to initialize the random generator
 */
//go:linkname LvRandSetSeed C.lv_rand_set_seed
func LvRandSetSeed(seed c.Uint32T)

/**
 * Get a pseudo random number in the given range
 * @param min   the minimum value
 * @param max   the maximum value
 * @return return the random number. min <= return_value <= max
 */
//go:linkname LvRand C.lv_rand
func LvRand(min c.Uint32T, max c.Uint32T) c.Uint32T

/**
 * Represents a point on the screen.
 */

type LvPointT struct {
	X c.Int32T
	Y c.Int32T
}

type LvPointPreciseT struct {
	X LvValuePreciseT
	Y LvValuePreciseT
}

/** Represents an area of the screen.*/

type LvAreaT struct {
	X1 c.Int32T
	Y1 c.Int32T
	X2 c.Int32T
	Y2 c.Int32T
}
type LvAlignT c.Int

const (
	LV_ALIGN_DEFAULT          LvAlignT = 0
	LV_ALIGN_TOP_LEFT         LvAlignT = 1
	LV_ALIGN_TOP_MID          LvAlignT = 2
	LV_ALIGN_TOP_RIGHT        LvAlignT = 3
	LV_ALIGN_BOTTOM_LEFT      LvAlignT = 4
	LV_ALIGN_BOTTOM_MID       LvAlignT = 5
	LV_ALIGN_BOTTOM_RIGHT     LvAlignT = 6
	LV_ALIGN_LEFT_MID         LvAlignT = 7
	LV_ALIGN_RIGHT_MID        LvAlignT = 8
	LV_ALIGN_CENTER           LvAlignT = 9
	LV_ALIGN_OUT_TOP_LEFT     LvAlignT = 10
	LV_ALIGN_OUT_TOP_MID      LvAlignT = 11
	LV_ALIGN_OUT_TOP_RIGHT    LvAlignT = 12
	LV_ALIGN_OUT_BOTTOM_LEFT  LvAlignT = 13
	LV_ALIGN_OUT_BOTTOM_MID   LvAlignT = 14
	LV_ALIGN_OUT_BOTTOM_RIGHT LvAlignT = 15
	LV_ALIGN_OUT_LEFT_TOP     LvAlignT = 16
	LV_ALIGN_OUT_LEFT_MID     LvAlignT = 17
	LV_ALIGN_OUT_LEFT_BOTTOM  LvAlignT = 18
	LV_ALIGN_OUT_RIGHT_TOP    LvAlignT = 19
	LV_ALIGN_OUT_RIGHT_MID    LvAlignT = 20
	LV_ALIGN_OUT_RIGHT_BOTTOM LvAlignT = 21
)

type LvDirT c.Int

const (
	LV_DIR_NONE   LvDirT = 0
	LV_DIR_LEFT   LvDirT = 1
	LV_DIR_RIGHT  LvDirT = 2
	LV_DIR_TOP    LvDirT = 4
	LV_DIR_BOTTOM LvDirT = 8
	LV_DIR_HOR    LvDirT = 3
	LV_DIR_VER    LvDirT = 12
	LV_DIR_ALL    LvDirT = 15
)

/**
 * Initialize an area
 * @param area_p pointer to an area
 * @param x1 left coordinate of the area
 * @param y1 top coordinate of the area
 * @param x2 right coordinate of the area
 * @param y2 bottom coordinate of the area
 */
// llgo:link (*LvAreaT).LvAreaSet C.lv_area_set
func (recv_ *LvAreaT) LvAreaSet(x1 c.Int32T, y1 c.Int32T, x2 c.Int32T, y2 c.Int32T) {
}

/**
 * Get the width of an area
 * @param area_p pointer to an area
 * @return the width of the area (if x1 == x2 -> width = 1)
 */
// llgo:link (*LvAreaT).LvAreaGetWidth C.lv_area_get_width
func (recv_ *LvAreaT) LvAreaGetWidth() c.Int32T {
	return 0
}

/**
 * Get the height of an area
 * @param area_p pointer to an area
 * @return the height of the area (if y1 == y2 -> height = 1)
 */
// llgo:link (*LvAreaT).LvAreaGetHeight C.lv_area_get_height
func (recv_ *LvAreaT) LvAreaGetHeight() c.Int32T {
	return 0
}

/**
 * Set the width of an area
 * @param area_p pointer to an area
 * @param w the new width of the area (w == 1 makes x1 == x2)
 */
// llgo:link (*LvAreaT).LvAreaSetWidth C.lv_area_set_width
func (recv_ *LvAreaT) LvAreaSetWidth(w c.Int32T) {
}

/**
 * Set the height of an area
 * @param area_p pointer to an area
 * @param h the new height of the area (h == 1 makes y1 == y2)
 */
// llgo:link (*LvAreaT).LvAreaSetHeight C.lv_area_set_height
func (recv_ *LvAreaT) LvAreaSetHeight(h c.Int32T) {
}

/**
 * Return with area of an area (x * y)
 * @param area_p pointer to an area
 * @return size of area
 */
// llgo:link (*LvAreaT).LvAreaGetSize C.lv_area_get_size
func (recv_ *LvAreaT) LvAreaGetSize() c.Uint32T {
	return 0
}

// llgo:link (*LvAreaT).LvAreaIncrease C.lv_area_increase
func (recv_ *LvAreaT) LvAreaIncrease(w_extra c.Int32T, h_extra c.Int32T) {
}

// llgo:link (*LvAreaT).LvAreaMove C.lv_area_move
func (recv_ *LvAreaT) LvAreaMove(x_ofs c.Int32T, y_ofs c.Int32T) {
}

/**
 * Align an area to another
 * @param base an area where the other will be aligned
 * @param to_align the area to align
 * @param align `LV_ALIGN_...`
 * @param ofs_x X offset
 * @param ofs_y Y offset
 */
// llgo:link (*LvAreaT).LvAreaAlign C.lv_area_align
func (recv_ *LvAreaT) LvAreaAlign(to_align *LvAreaT, align LvAlignT, ofs_x c.Int32T, ofs_y c.Int32T) {
}

/**
 * Transform a point
 * @param point         pointer to a point
 * @param angle         angle with 0.1 resolutions (123 means 12.3°)
 * @param scale_x       horizontal zoom, 256 means 100%
 * @param scale_y       vertical zoom, 256 means 100%
 * @param pivot         pointer to the pivot point of the transformation
 * @param zoom_first    true: zoom first and rotate after that; else: opposite order
 */
// llgo:link (*LvPointT).LvPointTransform C.lv_point_transform
func (recv_ *LvPointT) LvPointTransform(angle c.Int32T, scale_x c.Int32T, scale_y c.Int32T, pivot *LvPointT, zoom_first bool) {
}

/**
 * Transform an array of points
 * @param points        pointer to an array of points
 * @param count         number of points in the array
 * @param angle         angle with 0.1 resolutions (123 means 12.3°)
 * @param scale_x       horizontal zoom, 256 means 100%
 * @param scale_y       vertical zoom, 256 means 100%
 * @param pivot         pointer to the pivot point of the transformation
 * @param zoom_first    true: zoom first and rotate after that; else: opposite order
 */
// llgo:link (*LvPointT).LvPointArrayTransform C.lv_point_array_transform
func (recv_ *LvPointT) LvPointArrayTransform(count c.SizeT, angle c.Int32T, scale_x c.Int32T, scale_y c.Int32T, pivot *LvPointT, zoom_first bool) {
}

// llgo:link (*LvPointPreciseT).LvPointFromPrecise C.lv_point_from_precise
func (recv_ *LvPointPreciseT) LvPointFromPrecise() LvPointT {
	return LvPointT{}
}

// llgo:link (*LvPointT).LvPointToPrecise C.lv_point_to_precise
func (recv_ *LvPointT) LvPointToPrecise() LvPointPreciseT {
	return LvPointPreciseT{}
}

// llgo:link (*LvPointT).LvPointSet C.lv_point_set
func (recv_ *LvPointT) LvPointSet(x c.Int32T, y c.Int32T) {
}

// llgo:link (*LvPointPreciseT).LvPointPreciseSet C.lv_point_precise_set
func (recv_ *LvPointPreciseT) LvPointPreciseSet(x LvValuePreciseT, y LvValuePreciseT) {
}

// llgo:link (*LvPointT).LvPointSwap C.lv_point_swap
func (recv_ *LvPointT) LvPointSwap(p2 *LvPointT) {
}

// llgo:link (*LvPointPreciseT).LvPointPreciseSwap C.lv_point_precise_swap
func (recv_ *LvPointPreciseT) LvPointPreciseSwap(p2 *LvPointPreciseT) {
}

/**
 * Convert a percentage value to `int32_t`.
 * Percentage values are stored in special range
 * @param x the percentage (0..1000)
 * @return a coordinate that stores the percentage
 */
//go:linkname LvPct C.lv_pct
func LvPct(x c.Int32T) c.Int32T

//go:linkname LvPctToPx C.lv_pct_to_px
func LvPctToPx(v c.Int32T, base c.Int32T) c.Int32T

type LvLogLevelT c.Int8T

/**
 * @brief Copies a block of memory from a source address to a destination address.
 * @param dst Pointer to the destination array where the content is to be copied.
 * @param src Pointer to the source of data to be copied.
 * @param len Number of bytes to copy.
 * @return Pointer to the destination array.
 * @note The function does not check for any overlapping of the source and destination memory blocks.
 */
//go:linkname LvMemcpy C.lv_memcpy
func LvMemcpy(dst c.Pointer, src c.Pointer, len c.SizeT) c.Pointer

/**
 * @brief Fills a block of memory with a specified value.
 * @param dst Pointer to the destination array to fill with the specified value.
 * @param v Value to be set. The value is passed as an int, but the function fills
 *          the block of memory using the unsigned char conversion of this value.
 * @param len Number of bytes to be set to the value.
 */
//go:linkname LvMemset C.lv_memset
func LvMemset(dst c.Pointer, v c.Uint8T, len c.SizeT)

/**
 * @brief Move a block of memory from source to destination
 * @param dst Pointer to the destination array where the content is to be copied.
 * @param src Pointer to the source of data to be copied.
 * @param len Number of bytes to copy
 * @return Pointer to the destination array.
 */
//go:linkname LvMemmove C.lv_memmove
func LvMemmove(dst c.Pointer, src c.Pointer, len c.SizeT) c.Pointer

/**
 * @brief This function will compare two memory blocks
 * @param p1 Pointer to the first memory block
 * @param p2 Pointer to the second memory block
 * @param len Number of bytes to compare
 * @return The difference between the value of the first unmatching byte.
 */
//go:linkname LvMemcmp C.lv_memcmp
func LvMemcmp(p1 c.Pointer, p2 c.Pointer, len c.SizeT) c.Int

/**
 * @brief Computes the length of the string str up to (but not including) the terminating null character.
 * @param str Pointer to the null-terminated byte string to be examined.
 * @return The length of the string in bytes.
 */
//go:linkname LvStrlen C.lv_strlen
func LvStrlen(str *c.Char) c.SizeT

/**
 * @brief Computes the length of the string str up to (but not including) the terminating null character,
 *        or the given maximum length.
 * @param str Pointer to byte string that is null-terminated or at least max_len bytes long.
 * @param max_len Maximum number of characters to examine.
 * @return The length of the string in bytes.
 */
//go:linkname LvStrnlen C.lv_strnlen
func LvStrnlen(str *c.Char, max_len c.SizeT) c.SizeT

/**
 * @brief Copies up to dst_size-1 (non-null) characters from src to dst. A null terminator is always added.
 * @param dst Pointer to the destination array where the content is to be copied.
 * @param src Pointer to the source of data to be copied.
 * @param dst_size Maximum number of characters to be copied to dst, including the null character.
 * @return The length of src. The return value is equivalent to the value returned by lv_strlen(src)
 */
//go:linkname LvStrlcpy C.lv_strlcpy
func LvStrlcpy(dst *c.Char, src *c.Char, dst_size c.SizeT) c.SizeT

/**
 * @brief Copies up to dest_size characters from the string pointed to by src to the character array pointed to by dst
 *        and fills the remaining length with null bytes.
 * @param dst Pointer to the destination array where the content is to be copied.
 * @param src Pointer to the source of data to be copied.
 * @param dest_size Maximum number of characters to be copied to dst.
 * @return A pointer to the destination array, which is dst.
 * @note dst will not be null terminated if dest_size bytes were copied from src before the end of src was reached.
 */
//go:linkname LvStrncpy C.lv_strncpy
func LvStrncpy(dst *c.Char, src *c.Char, dest_size c.SizeT) *c.Char

/**
 * @brief Copies the string pointed to by src, including the terminating null character,
 *        to the character array pointed to by dst.
 * @param dst Pointer to the destination array where the content is to be copied.
 * @param src Pointer to the source of data to be copied.
 * @return A pointer to the destination array, which is dst.
 */
//go:linkname LvStrcpy C.lv_strcpy
func LvStrcpy(dst *c.Char, src *c.Char) *c.Char

/**
 * @brief  This function will compare two strings without specified length.
 * @param s1    pointer to the first string
 * @param s2    pointer to the second string
 * @return      the difference between the value of the first unmatching character.
 */
//go:linkname LvStrcmp C.lv_strcmp
func LvStrcmp(s1 *c.Char, s2 *c.Char) c.Int

/**
 * @brief  This function will compare two strings up to the given length.
 * @param s1    pointer to the first string
 * @param s2    pointer to the second string
 * @param len   the maximum amount of characters to compare
 * @return      the difference between the value of the first unmatching character.
 */
//go:linkname LvStrncmp C.lv_strncmp
func LvStrncmp(s1 *c.Char, s2 *c.Char, len c.SizeT) c.Int

/**
 * @brief Duplicate a string by allocating a new one and copying the content.
 * @param src Pointer to the source of data to be copied.
 * @return A pointer to the new allocated string. NULL if failed.
 */
//go:linkname LvStrdup C.lv_strdup
func LvStrdup(src *c.Char) *c.Char

/**
 * @brief Duplicate a string by allocating a new one and copying the content
 *        up to the end or the specified maximum length, whichever comes first.
 * @param src Pointer to the source of data to be copied.
 * @param max_len Maximum number of characters to be copied.
 * @return Pointer to a newly allocated null-terminated string. NULL if failed.
 */
//go:linkname LvStrndup C.lv_strndup
func LvStrndup(src *c.Char, max_len c.SizeT) *c.Char

/**
 * @brief Copies the string pointed to by src, including the terminating null character,
 *        to the end of the string pointed to by dst.
 * @param dst Pointer to the destination string where the content is to be appended.
 * @param src Pointer to the source of data to be copied.
 * @return A pointer to the destination string, which is dst.
 */
//go:linkname LvStrcat C.lv_strcat
func LvStrcat(dst *c.Char, src *c.Char) *c.Char

/**
 * @brief Copies up to src_len characters from the string pointed to by src
 *        to the end of the string pointed to by dst.
 *        A terminating null character is appended to dst even if no null character
 *        was encountered in src after src_len characters were copied.
 * @param dst Pointer to the destination string where the content is to be appended.
 * @param src Pointer to the source of data to be copied.
 * @param src_len Maximum number of characters from src to be copied to the end of dst.
 * @return A pointer to the destination string, which is dst.
 */
//go:linkname LvStrncat C.lv_strncat
func LvStrncat(dst *c.Char, src *c.Char, src_len c.SizeT) *c.Char

/**
 * @brief Searches for the first occurrence of character c in the string str.
 * @param str Pointer to the null-terminated byte string to be searched.
 * @param c The character to be searched for.
 * @return A pointer to the first occurrence of character c in the string str, or a null pointer if c is not found.
 */
//go:linkname LvStrchr C.lv_strchr
func LvStrchr(str *c.Char, c c.Int) *c.Char

type LvMemPoolT c.Pointer

/**
 * Heap information structure.
 */

type LvMemMonitorT struct {
	TotalSize       c.SizeT
	FreeCnt         c.SizeT
	FreeSize        c.SizeT
	FreeBiggestSize c.SizeT
	UsedCnt         c.SizeT
	MaxUsed         c.SizeT
	UsedPct         c.Uint8T
	FragPct         c.Uint8T
}

/**
 * Initialize to use malloc/free/realloc etc
 */
//go:linkname LvMemInit C.lv_mem_init
func LvMemInit()

/**
 * Drop all dynamically allocated memory and reset the memory pools' state
 */
//go:linkname LvMemDeinit C.lv_mem_deinit
func LvMemDeinit()

//go:linkname LvMemAddPool C.lv_mem_add_pool
func LvMemAddPool(mem c.Pointer, bytes c.SizeT) LvMemPoolT

//go:linkname LvMemRemovePool C.lv_mem_remove_pool
func LvMemRemovePool(pool LvMemPoolT)

/**
 * Allocate memory dynamically
 * @param size requested size in bytes
 * @return pointer to allocated uninitialized memory, or NULL on failure
 */
//go:linkname LvMalloc C.lv_malloc
func LvMalloc(size c.SizeT) c.Pointer

/**
 * Allocate a block of zeroed memory dynamically
 * @param num requested number of element to be allocated.
 * @param size requested size of each element in bytes.
 * @return pointer to allocated zeroed memory, or NULL on failure
 */
//go:linkname LvCalloc C.lv_calloc
func LvCalloc(num c.SizeT, size c.SizeT) c.Pointer

/**
 * Allocate zeroed memory dynamically
 * @param size requested size in bytes
 * @return pointer to allocated zeroed memory, or NULL on failure
 */
//go:linkname LvZalloc C.lv_zalloc
func LvZalloc(size c.SizeT) c.Pointer

/**
 * Allocate zeroed memory dynamically
 * @param size requested size in bytes
 * @return pointer to allocated zeroed memory, or NULL on failure
 */
//go:linkname LvMallocZeroed C.lv_malloc_zeroed
func LvMallocZeroed(size c.SizeT) c.Pointer

/**
 * Free an allocated data
 * @param data pointer to an allocated memory
 */
//go:linkname LvFree C.lv_free
func LvFree(data c.Pointer)

/**
 * Reallocate a memory with a new size. The old content will be kept.
 * @param data_p pointer to an allocated memory.
 *               Its content will be copied to the new memory block and freed
 * @param new_size the desired new size in byte
 * @return pointer to the new memory, NULL on failure
 */
//go:linkname LvRealloc C.lv_realloc
func LvRealloc(data_p c.Pointer, new_size c.SizeT) c.Pointer

/**
 * Reallocate a memory with a new size. The old content will be kept.
 * In case of failure, the old pointer is free'd.
 * @param data_p pointer to an allocated memory.
 *               Its content will be copied to the new memory block and freed
 * @param new_size the desired new size in byte
 * @return pointer to the new memory, NULL on failure
 */
//go:linkname LvReallocf C.lv_reallocf
func LvReallocf(data_p c.Pointer, new_size c.SizeT) c.Pointer

/**
 * Used internally to execute a plain `malloc` operation
 * @param size      size in bytes to `malloc`
 */
//go:linkname LvMallocCore C.lv_malloc_core
func LvMallocCore(size c.SizeT) c.Pointer

/**
 * Used internally to execute a plain `free` operation
 * @param p      memory address to free
 */
//go:linkname LvFreeCore C.lv_free_core
func LvFreeCore(p c.Pointer)

/**
 * Used internally to execute a plain realloc operation
 * @param p         memory address to realloc
 * @param new_size  size in bytes to realloc
 */
//go:linkname LvReallocCore C.lv_realloc_core
func LvReallocCore(p c.Pointer, new_size c.SizeT) c.Pointer

/**
 * Used internally by lv_mem_monitor() to gather LVGL heap state information.
 * @param mon_p      pointer to lv_mem_monitor_t object to be populated.
 */
// llgo:link (*LvMemMonitorT).LvMemMonitorCore C.lv_mem_monitor_core
func (recv_ *LvMemMonitorT) LvMemMonitorCore() {
}

//go:linkname LvMemTestCore C.lv_mem_test_core
func LvMemTestCore() LvResultT

/**
 * @brief Tests the memory allocation system by allocating and freeing a block of memory.
 * @return LV_RESULT_OK if the memory allocation system is working properly, or LV_RESULT_INVALID if there is an error.
 */
//go:linkname LvMemTest C.lv_mem_test
func LvMemTest() LvResultT

/**
 * Give information about the work memory of dynamic allocation
 * @param mon_p pointer to a lv_mem_monitor_t variable,
 *              the result of the analysis will be stored here
 */
// llgo:link (*LvMemMonitorT).LvMemMonitor C.lv_mem_monitor
func (recv_ *LvMemMonitorT) LvMemMonitor() {
}

/**********************
 * GLOBAL PROTOTYPES
 **********************/
//go:linkname LvAssertHandler C.lv_assert_handler
func LvAssertHandler()

type X_lvOpacityLevelT c.Int

const (
	LV_OPA_TRANSP X_lvOpacityLevelT = 0
	LV_OPA_0      X_lvOpacityLevelT = 0
	LV_OPA_10     X_lvOpacityLevelT = 25
	LV_OPA_20     X_lvOpacityLevelT = 51
	LV_OPA_30     X_lvOpacityLevelT = 76
	LV_OPA_40     X_lvOpacityLevelT = 102
	LV_OPA_50     X_lvOpacityLevelT = 127
	LV_OPA_60     X_lvOpacityLevelT = 153
	LV_OPA_70     X_lvOpacityLevelT = 178
	LV_OPA_80     X_lvOpacityLevelT = 204
	LV_OPA_90     X_lvOpacityLevelT = 229
	LV_OPA_100    X_lvOpacityLevelT = 255
	LV_OPA_COVER  X_lvOpacityLevelT = 255
)

/**********************
 *      TYPEDEFS
 **********************/

type LvColorT struct {
	Blue  c.Uint8T
	Green c.Uint8T
	Red   c.Uint8T
}

type LvColor16T struct {
	Blue  c.Uint16T
	Green c.Uint16T
	Red   c.Uint16T
}

type LvColor32T struct {
	Blue  c.Uint8T
	Green c.Uint8T
	Red   c.Uint8T
	Alpha c.Uint8T
}

type LvColorHsvT struct {
	H c.Uint16T
	S c.Uint8T
	V c.Uint8T
}

type LvColor16aT struct {
	Lumi  c.Uint8T
	Alpha c.Uint8T
}
type LvColorFormatT c.Int

const (
	LV_COLOR_FORMAT_UNKNOWN                LvColorFormatT = 0
	LV_COLOR_FORMAT_RAW                    LvColorFormatT = 1
	LV_COLOR_FORMAT_RAW_ALPHA              LvColorFormatT = 2
	LV_COLOR_FORMAT_L8                     LvColorFormatT = 6
	LV_COLOR_FORMAT_I1                     LvColorFormatT = 7
	LV_COLOR_FORMAT_I2                     LvColorFormatT = 8
	LV_COLOR_FORMAT_I4                     LvColorFormatT = 9
	LV_COLOR_FORMAT_I8                     LvColorFormatT = 10
	LV_COLOR_FORMAT_A8                     LvColorFormatT = 14
	LV_COLOR_FORMAT_RGB565                 LvColorFormatT = 18
	LV_COLOR_FORMAT_ARGB8565               LvColorFormatT = 19
	LV_COLOR_FORMAT_RGB565A8               LvColorFormatT = 20
	LV_COLOR_FORMAT_AL88                   LvColorFormatT = 21
	LV_COLOR_FORMAT_RGB565_SWAPPED         LvColorFormatT = 27
	LV_COLOR_FORMAT_RGB888                 LvColorFormatT = 15
	LV_COLOR_FORMAT_ARGB8888               LvColorFormatT = 16
	LV_COLOR_FORMAT_XRGB8888               LvColorFormatT = 17
	LV_COLOR_FORMAT_ARGB8888_PREMULTIPLIED LvColorFormatT = 26
	LV_COLOR_FORMAT_A1                     LvColorFormatT = 11
	LV_COLOR_FORMAT_A2                     LvColorFormatT = 12
	LV_COLOR_FORMAT_A4                     LvColorFormatT = 13
	LV_COLOR_FORMAT_ARGB1555               LvColorFormatT = 22
	LV_COLOR_FORMAT_ARGB4444               LvColorFormatT = 23
	LV_COLOR_FORMAT_ARGB2222               LvColorFormatT = 24
	LV_COLOR_FORMAT_YUV_START              LvColorFormatT = 32
	LV_COLOR_FORMAT_I420                   LvColorFormatT = 32
	LV_COLOR_FORMAT_I422                   LvColorFormatT = 33
	LV_COLOR_FORMAT_I444                   LvColorFormatT = 34
	LV_COLOR_FORMAT_I400                   LvColorFormatT = 35
	LV_COLOR_FORMAT_NV21                   LvColorFormatT = 36
	LV_COLOR_FORMAT_NV12                   LvColorFormatT = 37
	LV_COLOR_FORMAT_YUY2                   LvColorFormatT = 38
	LV_COLOR_FORMAT_UYVY                   LvColorFormatT = 39
	LV_COLOR_FORMAT_YUV_END                LvColorFormatT = 39
	LV_COLOR_FORMAT_PROPRIETARY_START      LvColorFormatT = 48
	LV_COLOR_FORMAT_NEMA_TSC_START         LvColorFormatT = 48
	LV_COLOR_FORMAT_NEMA_TSC4              LvColorFormatT = 48
	LV_COLOR_FORMAT_NEMA_TSC6              LvColorFormatT = 49
	LV_COLOR_FORMAT_NEMA_TSC6A             LvColorFormatT = 50
	LV_COLOR_FORMAT_NEMA_TSC6AP            LvColorFormatT = 51
	LV_COLOR_FORMAT_NEMA_TSC12             LvColorFormatT = 52
	LV_COLOR_FORMAT_NEMA_TSC12A            LvColorFormatT = 53
	LV_COLOR_FORMAT_NEMA_TSC_END           LvColorFormatT = 53
	LV_COLOR_FORMAT_NATIVE                 LvColorFormatT = 18
	LV_COLOR_FORMAT_NATIVE_WITH_ALPHA      LvColorFormatT = 20
)

/**
 * Get the pixel size of a color format in bits, bpp
 * @param cf        a color format (`LV_COLOR_FORMAT_...`)
 * @return          the pixel size in bits
 * @sa              LV_COLOR_FORMAT_GET_BPP
 */
// llgo:link LvColorFormatT.LvColorFormatGetBpp C.lv_color_format_get_bpp
func (recv_ LvColorFormatT) LvColorFormatGetBpp() c.Uint8T {
	return 0
}

/**
 * Get the pixel size of a color format in bytes
 * @param cf        a color format (`LV_COLOR_FORMAT_...`)
 * @return          the pixel size in bytes
 * @sa              LV_COLOR_FORMAT_GET_SIZE
 */
// llgo:link LvColorFormatT.LvColorFormatGetSize C.lv_color_format_get_size
func (recv_ LvColorFormatT) LvColorFormatGetSize() c.Uint8T {
	return 0
}

/**
 * Check if a color format has alpha channel or not
 * @param src_cf    a color format (`LV_COLOR_FORMAT_...`)
 * @return          true: has alpha channel; false: doesn't have alpha channel
 */
// llgo:link LvColorFormatT.LvColorFormatHasAlpha C.lv_color_format_has_alpha
func (recv_ LvColorFormatT) LvColorFormatHasAlpha() bool {
	return false
}

/**
 * Create an ARGB8888 color from RGB888 + alpha
 * @param color     an RGB888 color
 * @param opa       the alpha value
 * @return          the ARGB8888 color
 */
// llgo:link LvColorT.LvColorTo32 C.lv_color_to_32
func (recv_ LvColorT) LvColorTo32(opa LvOpaT) LvColor32T {
	return LvColor32T{}
}

/**
 * Convert an RGB888 color to an integer
 * @param c     an RGB888 color
 * @return      `c` as an integer
 */
// llgo:link LvColorT.LvColorToInt C.lv_color_to_int
func (recv_ LvColorT) LvColorToInt() c.Uint32T {
	return 0
}

/**
 * Check if two RGB888 color are equal
 * @param c1    the first color
 * @param c2    the second color
 * @return      true: equal
 */
// llgo:link LvColorT.LvColorEq C.lv_color_eq
func (recv_ LvColorT) LvColorEq(c2 LvColorT) bool {
	return false
}

/**
 * Check if two ARGB8888 color are equal
 * @param c1    the first color
 * @param c2    the second color
 * @return      true: equal
 */
// llgo:link LvColor32T.LvColor32Eq C.lv_color32_eq
func (recv_ LvColor32T) LvColor32Eq(c2 LvColor32T) bool {
	return false
}

/**
 * Create a color from 0x000000..0xffffff input
 * @param c     the hex input
 * @return      the color
 */
//go:linkname LvColorHex C.lv_color_hex
func LvColorHex(c c.Uint32T) LvColorT

/**
 * Create an RGB888 color
 * @param r     the red channel (0..255)
 * @param g     the green channel (0..255)
 * @param b     the blue channel (0..255)
 * @return      the color
 */
//go:linkname LvColorMake C.lv_color_make
func LvColorMake(r c.Uint8T, g c.Uint8T, b c.Uint8T) LvColorT

/**
 * Create an ARGB8888 color
 * @param r     the red channel (0..255)
 * @param g     the green channel (0..255)
 * @param b     the blue channel (0..255)
 * @param a     the alpha channel (0..255)
 * @return      the color
 */
//go:linkname LvColor32Make C.lv_color32_make
func LvColor32Make(r c.Uint8T, g c.Uint8T, b c.Uint8T, a c.Uint8T) LvColor32T

/**
 * Create a color from 0x000..0xfff input
 * @param c     the hex input (e.g. 0x123 will be 0x112233)
 * @return      the color
 */
//go:linkname LvColorHex3 C.lv_color_hex3
func LvColorHex3(c c.Uint32T) LvColorT

/**
 * Convert am RGB888 color to RGB565 stored in `uint16_t`
 * @param color     and RGB888 color
 * @return          `color` as RGB565 on `uin16_t`
 */
// llgo:link LvColorT.LvColorToU16 C.lv_color_to_u16
func (recv_ LvColorT) LvColorToU16() c.Uint16T {
	return 0
}

/**
 * Convert am RGB888 color to XRGB8888 stored in `uint32_t`
 * @param color     and RGB888 color
 * @return          `color` as XRGB8888 on `uin32_t` (the alpha channel is always set to 0xFF)
 */
// llgo:link LvColorT.LvColorToU32 C.lv_color_to_u32
func (recv_ LvColorT) LvColorToU32() c.Uint32T {
	return 0
}

/**
 * Mix two RGB565 colors
 * @param c1        the first color (typically the foreground color)
 * @param c2        the second color  (typically the background color)
 * @param mix       0..255, or LV_OPA_0/10/20...
 * @return          mix == 0: c2
 *                  mix == 255: c1
 *                  mix == 128: 0.5 x c1 + 0.5 x c2
 */
//go:linkname LvColor1616Mix C.lv_color_16_16_mix
func LvColor1616Mix(c1 c.Uint16T, c2 c.Uint16T, mix c.Uint8T) c.Uint16T

/**
 * Mix white to a color
 * @param c     the base color
 * @param lvl   the intensity of white (0: no change, 255: fully white)
 * @return      the mixed color
 */
// llgo:link LvColorT.LvColorLighten C.lv_color_lighten
func (recv_ LvColorT) LvColorLighten(lvl LvOpaT) LvColorT {
	return LvColorT{}
}

/**
 * Mix black to a color
 * @param c     the base color
 * @param lvl   the intensity of black (0: no change, 255: fully black)
 * @return      the mixed color
 */
// llgo:link LvColorT.LvColorDarken C.lv_color_darken
func (recv_ LvColorT) LvColorDarken(lvl LvOpaT) LvColorT {
	return LvColorT{}
}

/**
 * Convert a HSV color to RGB
 * @param h hue [0..359]
 * @param s saturation [0..100]
 * @param v value [0..100]
 * @return the given RGB color in RGB (with LV_COLOR_DEPTH depth)
 */
//go:linkname LvColorHsvToRgb C.lv_color_hsv_to_rgb
func LvColorHsvToRgb(h c.Uint16T, s c.Uint8T, v c.Uint8T) LvColorT

/**
 * Convert a 32-bit RGB color to HSV
 * @param r8 8-bit red
 * @param g8 8-bit green
 * @param b8 8-bit blue
 * @return the given RGB color in HSV
 */
//go:linkname LvColorRgbToHsv C.lv_color_rgb_to_hsv
func LvColorRgbToHsv(r8 c.Uint8T, g8 c.Uint8T, b8 c.Uint8T) LvColorHsvT

/**
 * Convert a color to HSV
 * @param color color
 * @return the given color in HSV
 */
// llgo:link LvColorT.LvColorToHsv C.lv_color_to_hsv
func (recv_ LvColorT) LvColorToHsv() LvColorHsvT {
	return LvColorHsvT{}
}

/**
 * A helper for white color
 * @return      a white color
 */
//go:linkname LvColorWhite C.lv_color_white
func LvColorWhite() LvColorT

/**
 * A helper for black color
 * @return      a black color
 */
//go:linkname LvColorBlack C.lv_color_black
func LvColorBlack() LvColorT

// llgo:link (*LvColor32T).LvColorPremultiply C.lv_color_premultiply
func (recv_ *LvColor32T) LvColorPremultiply() {
}

// llgo:link (*LvColor16T).LvColor16Premultiply C.lv_color16_premultiply
func (recv_ *LvColor16T) LvColor16Premultiply(a LvOpaT) {
}

/**
 * Get the luminance of a color: luminance = 0.3 R + 0.59 G + 0.11 B
 * @param c a color
 * @return the brightness [0..255]
 */
// llgo:link LvColorT.LvColorLuminance C.lv_color_luminance
func (recv_ LvColorT) LvColorLuminance() c.Uint8T {
	return 0
}

/**
 * Get the luminance of a color16: luminance = 0.3 R + 0.59 G + 0.11 B
 * @param c a color
 * @return the brightness [0..255]
 */
// llgo:link LvColor16T.LvColor16Luminance C.lv_color16_luminance
func (recv_ LvColor16T) LvColor16Luminance() c.Uint8T {
	return 0
}

/**
 * Get the luminance of a color24: luminance = 0.3 R + 0.59 G + 0.11 B
 * @param c a color
 * @return the brightness [0..255]
 */
//go:linkname LvColor24Luminance C.lv_color24_luminance
func LvColor24Luminance(c *c.Uint8T) c.Uint8T

/**
 * Get the luminance of a color32: luminance = 0.3 R + 0.59 G + 0.11 B
 * @param c a color
 * @return the brightness [0..255]
 */
// llgo:link LvColor32T.LvColor32Luminance C.lv_color32_luminance
func (recv_ LvColor32T) LvColor32Luminance() c.Uint8T {
	return 0
}

type LvPaletteT c.Int

const (
	LV_PALETTE_RED         LvPaletteT = 0
	LV_PALETTE_PINK        LvPaletteT = 1
	LV_PALETTE_PURPLE      LvPaletteT = 2
	LV_PALETTE_DEEP_PURPLE LvPaletteT = 3
	LV_PALETTE_INDIGO      LvPaletteT = 4
	LV_PALETTE_BLUE        LvPaletteT = 5
	LV_PALETTE_LIGHT_BLUE  LvPaletteT = 6
	LV_PALETTE_CYAN        LvPaletteT = 7
	LV_PALETTE_TEAL        LvPaletteT = 8
	LV_PALETTE_GREEN       LvPaletteT = 9
	LV_PALETTE_LIGHT_GREEN LvPaletteT = 10
	LV_PALETTE_LIME        LvPaletteT = 11
	LV_PALETTE_YELLOW      LvPaletteT = 12
	LV_PALETTE_AMBER       LvPaletteT = 13
	LV_PALETTE_ORANGE      LvPaletteT = 14
	LV_PALETTE_DEEP_ORANGE LvPaletteT = 15
	LV_PALETTE_BROWN       LvPaletteT = 16
	LV_PALETTE_BLUE_GREY   LvPaletteT = 17
	LV_PALETTE_GREY        LvPaletteT = 18
	LV_PALETTE_LAST        LvPaletteT = 19
	LV_PALETTE_NONE        LvPaletteT = 255
)

/*Source: https://vuetifyjs.com/en/styles/colors/#material-colors*/
// llgo:link LvPaletteT.LvPaletteMain C.lv_palette_main
func (recv_ LvPaletteT) LvPaletteMain() LvColorT {
	return LvColorT{}
}

// llgo:link LvPaletteT.LvPaletteLighten C.lv_palette_lighten
func (recv_ LvPaletteT) LvPaletteLighten(lvl c.Uint8T) LvColorT {
	return LvColorT{}
}

// llgo:link LvPaletteT.LvPaletteDarken C.lv_palette_darken
func (recv_ LvPaletteT) LvPaletteDarken(lvl c.Uint8T) LvColorT {
	return LvColorT{}
}

// llgo:type C
type LvColorFilterCbT func(*X_lvColorFilterDscT, LvColorT, LvOpaT) LvColorT

/**
 * Mix two colors with a given ratio.
 * @param c1 the first color to mix (usually the foreground)
 * @param c2 the second color to mix (usually the background)
 * @param mix The ratio of the colors. 0: full `c2`, 255: full `c1`, 127: half `c1` and half`c2`
 * @return the mixed color
 */
// llgo:link LvColorT.LvColorMix C.lv_color_mix
func (recv_ LvColorT) LvColorMix(c2 LvColorT, mix c.Uint8T) LvColorT {
	return LvColorT{}
}

/**
 *
 * @param fg
 * @param bg
 * @return
 * @note Use bg.alpha in the return value
 * @note Use fg.alpha as mix ratio
 */
// llgo:link LvColor32T.LvColorMix32 C.lv_color_mix32
func (recv_ LvColor32T) LvColorMix32(bg LvColor32T) LvColor32T {
	return LvColor32T{}
}

/**
 * @brief Blends two premultiplied ARGB8888 colors while maintaining correct alpha compositing.
 *
 * This function correctly blends the foreground (fg) and background (bg) colors,
 * ensuring that the output remains in a premultiplied alpha format.
 *
 * @param fg The foreground color in premultiplied ARGB8888 format.
 * @param bg The background color in premultiplied ARGB8888 format.
 * @return The resulting blended color in premultiplied ARGB8888 format.
 *
 * @note If the foreground is fully opaque, it is returned as is.
 * @note If the foreground is fully transparent, the background is returned.
 */
// llgo:link LvColor32T.LvColorMix32Premultiplied C.lv_color_mix32_premultiplied
func (recv_ LvColor32T) LvColorMix32Premultiplied(bg LvColor32T) LvColor32T {
	return LvColor32T{}
}

/**
 * Get the brightness of a color
 * @param c   a color
 * @return brightness in range [0..255]
 */
// llgo:link LvColorT.LvColorBrightness C.lv_color_brightness
func (recv_ LvColorT) LvColorBrightness() c.Uint8T {
	return 0
}

// llgo:link (*LvColorFilterDscT).LvColorFilterDscInit C.lv_color_filter_dsc_init
func (recv_ *LvColorFilterDscT) LvColorFilterDscInit(cb LvColorFilterCbT) {
}

/**
 * Blend two colors that have not been pre-multiplied using their alpha values
 * @param fg the foreground color
 * @param bg the background color
 * @return result color
 */
// llgo:link LvColor32T.LvColorOver32 C.lv_color_over32
func (recv_ LvColor32T) LvColorOver32(bg LvColor32T) LvColor32T {
	return LvColor32T{}
}

type X_lvimageFlagsT c.Int

const (
	LV_IMAGE_FLAGS_PREMULTIPLIED X_lvimageFlagsT = 1
	LV_IMAGE_FLAGS_COMPRESSED    X_lvimageFlagsT = 8
	LV_IMAGE_FLAGS_ALLOCATED     X_lvimageFlagsT = 16
	LV_IMAGE_FLAGS_MODIFIABLE    X_lvimageFlagsT = 32
	LV_IMAGE_FLAGS_CUSTOM_DRAW   X_lvimageFlagsT = 64
	LV_IMAGE_FLAGS_USER1         X_lvimageFlagsT = 256
	LV_IMAGE_FLAGS_USER2         X_lvimageFlagsT = 512
	LV_IMAGE_FLAGS_USER3         X_lvimageFlagsT = 1024
	LV_IMAGE_FLAGS_USER4         X_lvimageFlagsT = 2048
	LV_IMAGE_FLAGS_USER5         X_lvimageFlagsT = 4096
	LV_IMAGE_FLAGS_USER6         X_lvimageFlagsT = 8192
	LV_IMAGE_FLAGS_USER7         X_lvimageFlagsT = 16384
	LV_IMAGE_FLAGS_USER8         X_lvimageFlagsT = 32768
)

type LvImageFlagsT X_lvimageFlagsT
type LvImageCompressT c.Int

const (
	LV_IMAGE_COMPRESS_NONE LvImageCompressT = 0
	LV_IMAGE_COMPRESS_RLE  LvImageCompressT = 1
	LV_IMAGE_COMPRESS_LZ4  LvImageCompressT = 2
)

type LvImageHeaderT struct {
	Magic     c.Uint32T
	Cf        c.Uint32T
	Flags     c.Uint32T
	W         c.Uint32T
	H         c.Uint32T
	Stride    c.Uint32T
	Reserved2 c.Uint32T
}

type LvYuvPlaneT struct {
	Buf    c.Pointer
	Stride c.Uint32T
}

type LvYuvBufT struct {
	Planar struct {
		Y LvYuvPlaneT
		U LvYuvPlaneT
		V LvYuvPlaneT
	}
}

/**
 * Struct to describe a constant image resource.
 * It's similar to lv_draw_buf_t, but the data is constant.
 */

type LvImageDscT struct {
	Header    LvImageHeaderT
	DataSize  c.Uint32T
	Data      *c.Uint8T
	Reserved  c.Pointer
	Reserved2 c.Pointer
}

// llgo:type C
type LvDrawBufMallocCb func(c.SizeT, LvColorFormatT) c.Pointer

// llgo:type C
type LvDrawBufFreeCb func(c.Pointer)

// llgo:type C
type LvDrawBufAlignCb func(c.Pointer, LvColorFormatT) c.Pointer

// llgo:type C
type LvDrawBufCacheOperationCb func(*LvDrawBufT, *LvAreaT)

// llgo:type C
type LvDrawBufWidthToStrideCb func(c.Uint32T, LvColorFormatT) c.Uint32T

/**
 * Initialize the draw buffer with the default handlers.
 *
 * @param handlers  the draw buffer handlers to set
 */
// llgo:link (*LvDrawBufHandlersT).LvDrawBufInitWithDefaultHandlers C.lv_draw_buf_init_with_default_handlers
func (recv_ *LvDrawBufHandlersT) LvDrawBufInitWithDefaultHandlers() {
}

/**
 * Initialize the draw buffer with given handlers.
 *
 * @param handlers             the draw buffer handlers to set
 * @param buf_malloc_cb        the callback to allocate memory for the buffer
 * @param buf_free_cb          the callback to free memory of the buffer
 * @param align_pointer_cb     the callback to align the buffer
 * @param invalidate_cache_cb  the callback to invalidate the cache of the buffer
 * @param flush_cache_cb       the callback to flush buffer
 * @param width_to_stride_cb   the callback to calculate the stride based on the width and color format
 */
// llgo:link (*LvDrawBufHandlersT).LvDrawBufHandlersInit C.lv_draw_buf_handlers_init
func (recv_ *LvDrawBufHandlersT) LvDrawBufHandlersInit(buf_malloc_cb LvDrawBufMallocCb, buf_free_cb LvDrawBufFreeCb, align_pointer_cb LvDrawBufAlignCb, invalidate_cache_cb LvDrawBufCacheOperationCb, flush_cache_cb LvDrawBufCacheOperationCb, width_to_stride_cb LvDrawBufWidthToStrideCb) {
}

/**
 * Get the struct which holds the callbacks for draw buf management.
 * Custom callback can be set on the returned value
 * @return                  pointer to the struct of handlers
 */
//go:linkname LvDrawBufGetHandlers C.lv_draw_buf_get_handlers
func LvDrawBufGetHandlers() *LvDrawBufHandlersT

//go:linkname LvDrawBufGetFontHandlers C.lv_draw_buf_get_font_handlers
func LvDrawBufGetFontHandlers() *LvDrawBufHandlersT

//go:linkname LvDrawBufGetImageHandlers C.lv_draw_buf_get_image_handlers
func LvDrawBufGetImageHandlers() *LvDrawBufHandlersT

/**
 * Align the address of a buffer. The buffer needs to be large enough for the real data after alignment
 * @param buf           the data to align
 * @param color_format  the color format of the buffer
 * @return              the aligned buffer
 */
//go:linkname LvDrawBufAlign C.lv_draw_buf_align
func LvDrawBufAlign(buf c.Pointer, color_format LvColorFormatT) c.Pointer

/**
 * Align the address of a buffer with custom draw buffer handlers.
 * The buffer needs to be large enough for the real data after alignment
 * @param handlers      the draw buffer handlers
 * @param buf           the data to align
 * @param color_format  the color format of the buffer
 * @return              the aligned buffer
 */
// llgo:link (*LvDrawBufHandlersT).LvDrawBufAlignEx C.lv_draw_buf_align_ex
func (recv_ *LvDrawBufHandlersT) LvDrawBufAlignEx(buf c.Pointer, color_format LvColorFormatT) c.Pointer {
	return nil
}

/**
 * Invalidate the cache of the buffer
 * @param draw_buf     the draw buffer needs to be invalidated
 * @param area         the area to invalidate in the buffer,
 *                     use NULL to invalidate the whole draw buffer address range
 */
// llgo:link (*LvDrawBufT).LvDrawBufInvalidateCache C.lv_draw_buf_invalidate_cache
func (recv_ *LvDrawBufT) LvDrawBufInvalidateCache(area *LvAreaT) {
}

/**
 * Flush the cache of the buffer
 * @param draw_buf     the draw buffer needs to be flushed
 * @param area         the area to flush in the buffer,
 *                     use NULL to flush the whole draw buffer address range
 */
// llgo:link (*LvDrawBufT).LvDrawBufFlushCache C.lv_draw_buf_flush_cache
func (recv_ *LvDrawBufT) LvDrawBufFlushCache(area *LvAreaT) {
}

/**
 * Calculate the stride in bytes based on a width and color format
 * @param w                 the width in pixels
 * @param color_format      the color format
 * @return                  the stride in bytes
 */
//go:linkname LvDrawBufWidthToStride C.lv_draw_buf_width_to_stride
func LvDrawBufWidthToStride(w c.Uint32T, color_format LvColorFormatT) c.Uint32T

/**
 * Calculate the stride in bytes based on a width and color format
 * @param handlers          the draw buffer handlers
 * @param w                 the width in pixels
 * @param color_format      the color format
 * @return                  the stride in bytes
 */
// llgo:link (*LvDrawBufHandlersT).LvDrawBufWidthToStrideEx C.lv_draw_buf_width_to_stride_ex
func (recv_ *LvDrawBufHandlersT) LvDrawBufWidthToStrideEx(w c.Uint32T, color_format LvColorFormatT) c.Uint32T {
	return 0
}

/**
 * Clear an area on the buffer
 * @param draw_buf          pointer to draw buffer
 * @param a                 the area to clear, or NULL to clear the whole buffer
 */
// llgo:link (*LvDrawBufT).LvDrawBufClear C.lv_draw_buf_clear
func (recv_ *LvDrawBufT) LvDrawBufClear(a *LvAreaT) {
}

/**
 * Copy an area from a buffer to another
 * @param dest      pointer to the destination draw buffer
 * @param dest_area the area to copy from the destination buffer, if NULL, use the whole buffer
 * @param src       pointer to the source draw buffer
 * @param src_area  the area to copy from the destination buffer, if NULL, use the whole buffer
 * @note `dest_area` and `src_area` should have the same width and height
 * @note  `dest` and `src` should have same color format. Color converting is not supported fow now.
 */
// llgo:link (*LvDrawBufT).LvDrawBufCopy C.lv_draw_buf_copy
func (recv_ *LvDrawBufT) LvDrawBufCopy(dest_area *LvAreaT, src *LvDrawBufT, src_area *LvAreaT) {
}

/**
 * Note: Eventually, lv_draw_buf_malloc/free will be kept as private.
 *       For now, we use `create` to distinguish with malloc.
 *
 * Create an draw buf by allocating struct for `lv_draw_buf_t` and allocating a buffer for it
 * that meets specified requirements.
 *
 * @param w         the buffer width in pixels
 * @param h         the buffer height in pixels
 * @param cf        the color format for image
 * @param stride    the stride in bytes for image. Use 0 for automatic calculation based on
 *                  w, cf, and global stride alignment configuration.
 */
//go:linkname LvDrawBufCreate C.lv_draw_buf_create
func LvDrawBufCreate(w c.Uint32T, h c.Uint32T, cf LvColorFormatT, stride c.Uint32T) *LvDrawBufT

/**
 * Note: Eventually, lv_draw_buf_malloc/free will be kept as private.
 *       For now, we use `create` to distinguish with malloc.
 *
 * Create an draw buf by allocating struct for `lv_draw_buf_t` and allocating a buffer for it
 * that meets specified requirements.
 *
 * @param handlers  the draw buffer handlers
 * @param w         the buffer width in pixels
 * @param h         the buffer height in pixels
 * @param cf        the color format for image
 * @param stride    the stride in bytes for image. Use 0 for automatic calculation based on
 *                  w, cf, and global stride alignment configuration.
 */
// llgo:link (*LvDrawBufHandlersT).LvDrawBufCreateEx C.lv_draw_buf_create_ex
func (recv_ *LvDrawBufHandlersT) LvDrawBufCreateEx(w c.Uint32T, h c.Uint32T, cf LvColorFormatT, stride c.Uint32T) *LvDrawBufT {
	return nil
}

/**
 * Duplicate a draw buf with same image size, stride and color format. Copy the image data too.
 * @param draw_buf  the draw buf to duplicate
 * @return          the duplicated draw buf on success, NULL if failed
 */
// llgo:link (*LvDrawBufT).LvDrawBufDup C.lv_draw_buf_dup
func (recv_ *LvDrawBufT) LvDrawBufDup() *LvDrawBufT {
	return nil
}

/**
 * Duplicate a draw buf with same image size, stride and color format. Copy the image data too.
 * @param handlers  the draw buffer handlers
 * @param draw_buf  the draw buf to duplicate
 * @return          the duplicated draw buf on success, NULL if failed
 */
// llgo:link (*LvDrawBufHandlersT).LvDrawBufDupEx C.lv_draw_buf_dup_ex
func (recv_ *LvDrawBufHandlersT) LvDrawBufDupEx(draw_buf *LvDrawBufT) *LvDrawBufT {
	return nil
}

/**
 * Initialize a draw buf with the given buffer and parameters. Clear draw buffer flag to zero.
 * @param draw_buf  the draw buf to initialize
 * @param w         the buffer width in pixels
 * @param h         the buffer height in pixels
 * @param cf        the color format
 * @param stride    the stride in bytes. Use 0 for automatic calculation
 * @param data      the buffer used for drawing. Unaligned `data` will be aligned internally
 * @param data_size the size of the buffer in bytes
 * @return          return LV_RESULT_OK on success, LV_RESULT_INVALID otherwise
 */
// llgo:link (*LvDrawBufT).LvDrawBufInit C.lv_draw_buf_init
func (recv_ *LvDrawBufT) LvDrawBufInit(w c.Uint32T, h c.Uint32T, cf LvColorFormatT, stride c.Uint32T, data c.Pointer, data_size c.Uint32T) LvResultT {
	return 0
}

/**
 * Keep using the existing memory, reshape the draw buffer to the given width and height.
 * Return NULL if data_size is smaller than the required size.
 * @param draw_buf  pointer to a draw buffer
 * @param cf        the new color format, use 0 or LV_COLOR_FORMAT_UNKNOWN to keep using the original color format.
 * @param w         the new width in pixels
 * @param h         the new height in pixels
 * @param stride    the stride in bytes for image. Use 0 for automatic calculation.
 */
// llgo:link (*LvDrawBufT).LvDrawBufReshape C.lv_draw_buf_reshape
func (recv_ *LvDrawBufT) LvDrawBufReshape(cf LvColorFormatT, w c.Uint32T, h c.Uint32T, stride c.Uint32T) *LvDrawBufT {
	return nil
}

/**
 * Destroy a draw buf by freeing the actual buffer if it's marked as LV_IMAGE_FLAGS_ALLOCATED in header.
 * Then free the lv_draw_buf_t struct.
 *
 * @param draw_buf  the draw buffer to destroy
 */
// llgo:link (*LvDrawBufT).LvDrawBufDestroy C.lv_draw_buf_destroy
func (recv_ *LvDrawBufT) LvDrawBufDestroy() {
}

/**
 * Return pointer to the buffer at the given coordinates
 */
// llgo:link (*LvDrawBufT).LvDrawBufGotoXy C.lv_draw_buf_goto_xy
func (recv_ *LvDrawBufT) LvDrawBufGotoXy(x c.Uint32T, y c.Uint32T) c.Pointer {
	return nil
}

/**
 * Adjust the stride of a draw buf in place.
 * @param src       pointer to a draw buffer
 * @param stride    the new stride in bytes for image. Use LV_STRIDE_AUTO for automatic calculation.
 * @return          LV_RESULT_OK: success or LV_RESULT_INVALID: failed
 */
// llgo:link (*LvDrawBufT).LvDrawBufAdjustStride C.lv_draw_buf_adjust_stride
func (recv_ *LvDrawBufT) LvDrawBufAdjustStride(stride c.Uint32T) LvResultT {
	return 0
}

/**
 * Premultiply draw buffer color with alpha channel.
 * If it's already premultiplied, return directly.
 * Only color formats with alpha channel will be processed.
 *
 * @return LV_RESULT_OK: premultiply success
 */
// llgo:link (*LvDrawBufT).LvDrawBufPremultiply C.lv_draw_buf_premultiply
func (recv_ *LvDrawBufT) LvDrawBufPremultiply() LvResultT {
	return 0
}

// llgo:link (*LvDrawBufT).LvDrawBufHasFlag C.lv_draw_buf_has_flag
func (recv_ *LvDrawBufT) LvDrawBufHasFlag(flag LvImageFlagsT) bool {
	return false
}

// llgo:link (*LvDrawBufT).LvDrawBufSetFlag C.lv_draw_buf_set_flag
func (recv_ *LvDrawBufT) LvDrawBufSetFlag(flag LvImageFlagsT) {
}

// llgo:link (*LvDrawBufT).LvDrawBufClearFlag C.lv_draw_buf_clear_flag
func (recv_ *LvDrawBufT) LvDrawBufClearFlag(flag LvImageFlagsT) {
}

/**
 * As of now, draw buf share same definition as `lv_image_dsc_t`.
 * And is interchangeable with `lv_image_dsc_t`.
 */
// llgo:link (*LvDrawBufT).LvDrawBufFromImage C.lv_draw_buf_from_image
func (recv_ *LvDrawBufT) LvDrawBufFromImage(img *LvImageDscT) LvResultT {
	return 0
}

// llgo:link (*LvDrawBufT).LvDrawBufToImage C.lv_draw_buf_to_image
func (recv_ *LvDrawBufT) LvDrawBufToImage(img *LvImageDscT) {
}

/**
 * Set the palette color of an indexed image. Valid only for `LV_COLOR_FORMAT_I1/2/4/8`
 * @param draw_buf pointer to an image descriptor
 * @param index the palette color to set:
 *   - for `LV_COLOR_FORMAT_I1`: 0..1
 *   - for `LV_COLOR_FORMAT_I2`: 0..3
 *   - for `LV_COLOR_FORMAT_I4`: 0..15
 *   - for `LV_COLOR_FORMAT_I8`: 0..255
 * @param color the color to set in lv_color32_t format
 */
// llgo:link (*LvDrawBufT).LvDrawBufSetPalette C.lv_draw_buf_set_palette
func (recv_ *LvDrawBufT) LvDrawBufSetPalette(index c.Uint8T, color LvColor32T) {
}

/**
 * @deprecated Use lv_draw_buf_set_palette instead.
 */
// llgo:link (*LvImageDscT).LvImageBufSetPalette C.lv_image_buf_set_palette
func (recv_ *LvImageDscT) LvImageBufSetPalette(id c.Uint8T, c LvColor32T) {
}

/**
 * @deprecated Use lv_draw_buffer_create/destroy instead.
 * Free the data pointer and dsc struct of an image.
 */
// llgo:link (*LvImageDscT).LvImageBufFree C.lv_image_buf_free
func (recv_ *LvImageDscT) LvImageBufFree() {
}

type LvFontGlyphFormatT c.Int

const (
	LV_FONT_GLYPH_FORMAT_NONE   LvFontGlyphFormatT = 0
	LV_FONT_GLYPH_FORMAT_A1     LvFontGlyphFormatT = 1
	LV_FONT_GLYPH_FORMAT_A2     LvFontGlyphFormatT = 2
	LV_FONT_GLYPH_FORMAT_A3     LvFontGlyphFormatT = 3
	LV_FONT_GLYPH_FORMAT_A4     LvFontGlyphFormatT = 4
	LV_FONT_GLYPH_FORMAT_A8     LvFontGlyphFormatT = 8
	LV_FONT_GLYPH_FORMAT_IMAGE  LvFontGlyphFormatT = 25
	LV_FONT_GLYPH_FORMAT_VECTOR LvFontGlyphFormatT = 26
	LV_FONT_GLYPH_FORMAT_SVG    LvFontGlyphFormatT = 27
	LV_FONT_GLYPH_FORMAT_CUSTOM LvFontGlyphFormatT = 255
)

/** Describes the properties of a glyph.*/

type LvFontGlyphDscT struct {
	ResolvedFont       *LvFontT
	AdvW               c.Uint16T
	BoxW               c.Uint16T
	BoxH               c.Uint16T
	OfsX               c.Int16T
	OfsY               c.Int16T
	Stride             c.Uint16T
	Format             LvFontGlyphFormatT
	IsPlaceholder      c.Uint8T
	ReqRawBitmap       c.Uint8T
	OutlineStrokeWidth c.Int32T
	Gid                struct {
		Src c.Pointer
	}
	Entry *LvCacheEntryT
}
type LvFontSubpxT c.Int

const (
	LV_FONT_SUBPX_NONE LvFontSubpxT = 0
	LV_FONT_SUBPX_HOR  LvFontSubpxT = 1
	LV_FONT_SUBPX_VER  LvFontSubpxT = 2
	LV_FONT_SUBPX_BOTH LvFontSubpxT = 3
)

type LvFontKerningT c.Int

const (
	LV_FONT_KERNING_NORMAL LvFontKerningT = 0
	LV_FONT_KERNING_NONE   LvFontKerningT = 1
)

/**
 * Return with the bitmap of a font.
 * It always converts the normal fonts to A8 format in a draw_buf with
 * LV_DRAW_BUF_ALIGN and LV_DRAW_BUF_STRIDE_ALIGN
 * @note You must call lv_font_get_glyph_dsc() to get `g_dsc` (lv_font_glyph_dsc_t)
 *       before you can call this function.
 * @param g_dsc         the glyph descriptor including which font to use, which supply the glyph_index
 *                      and the format.
 * @param draw_buf      a draw buffer that can be used to store the bitmap of the glyph.
 * @return              pointer to the glyph's data.
 *                      It can be a draw buffer for bitmap fonts or an image source for imgfonts.
 */
// llgo:link (*LvFontGlyphDscT).LvFontGetGlyphBitmap C.lv_font_get_glyph_bitmap
func (recv_ *LvFontGlyphDscT) LvFontGetGlyphBitmap(draw_buf *LvDrawBufT) c.Pointer {
	return nil
}

/**
 * Return the bitmap as it is. It works only if the font stores the bitmap in
 * a non-volitile memory.
 * @param g_dsc         the glyph descriptor including which font to use, which supply the glyph_index
 *                      and the format.
 * @return              the bitmap as it is
 */
// llgo:link (*LvFontGlyphDscT).LvFontGetGlyphStaticBitmap C.lv_font_get_glyph_static_bitmap
func (recv_ *LvFontGlyphDscT) LvFontGetGlyphStaticBitmap() c.Pointer {
	return nil
}

/**
 * Get the descriptor of a glyph
 * @param font          pointer to font
 * @param dsc_out       store the result descriptor here
 * @param letter        a UNICODE letter code
 * @param letter_next   the next letter after `letter`. Used for kerning
 * @return true: descriptor is successfully loaded into `dsc_out`.
 *         false: the letter was not found, no data is loaded to `dsc_out`
 */
// llgo:link (*LvFontT).LvFontGetGlyphDsc C.lv_font_get_glyph_dsc
func (recv_ *LvFontT) LvFontGetGlyphDsc(dsc_out *LvFontGlyphDscT, letter c.Uint32T, letter_next c.Uint32T) bool {
	return false
}

/**
 * Release the bitmap of a font.
 * @note You must call lv_font_get_glyph_dsc() to get `g_dsc` (lv_font_glyph_dsc_t) before you can call this function.
 * @param g_dsc         the glyph descriptor including which font to use, which supply the glyph_index and the format.
 */
// llgo:link (*LvFontGlyphDscT).LvFontGlyphReleaseDrawData C.lv_font_glyph_release_draw_data
func (recv_ *LvFontGlyphDscT) LvFontGlyphReleaseDrawData() {
}

/**
 * Get the width of a glyph with kerning
 * @param font          pointer to a font
 * @param letter        a UNICODE letter
 * @param letter_next   the next letter after `letter`. Used for kerning
 * @return the width of the glyph
 */
// llgo:link (*LvFontT).LvFontGetGlyphWidth C.lv_font_get_glyph_width
func (recv_ *LvFontT) LvFontGetGlyphWidth(letter c.Uint32T, letter_next c.Uint32T) c.Uint16T {
	return 0
}

/**
 * Get the line height of a font. All characters fit into this height
 * @param font      pointer to a font
 * @return the height of a font
 */
// llgo:link (*LvFontT).LvFontGetLineHeight C.lv_font_get_line_height
func (recv_ *LvFontT) LvFontGetLineHeight() c.Int32T {
	return 0
}

/**
 * Configure the use of kerning information stored in a font
 * @param font    pointer to a font
 * @param kerning `LV_FONT_KERNING_NORMAL` (default) or `LV_FONT_KERNING_NONE`
 */
// llgo:link (*LvFontT).LvFontSetKerning C.lv_font_set_kerning
func (recv_ *LvFontT) LvFontSetKerning(kerning LvFontKerningT) {
}

/**
 * Get the default font, defined by LV_FONT_DEFAULT
 * @return  return      pointer to the default font
 */
//go:linkname LvFontGetDefault C.lv_font_get_default
func LvFontGetDefault() *LvFontT

/**
 * Compare font information.
 * @param ft_info_1 font information 1.
 * @param ft_info_2 font information 2.
 * @return return true if the fonts are equal.
 */
// llgo:link (*LvFontInfoT).LvFontInfoIsEqual C.lv_font_info_is_equal
func (recv_ *LvFontInfoT) LvFontInfoIsEqual(ft_info_2 *LvFontInfoT) bool {
	return false
}

/**
 * Checks if a font has a static rendering bitmap.
 * @param font    pointer to a font
 * @return return true if the font has a bitmap generated for static rendering.
 */
// llgo:link (*LvFontT).LvFontHasStaticBitmap C.lv_font_has_static_bitmap
func (recv_ *LvFontT) LvFontHasStaticBitmap() bool {
	return false
}

// llgo:type C
type LvTickGetCbT func() c.Uint32T

// llgo:type C
type LvDelayCbT func(c.Uint32T)

/**
 * You have to call this function periodically
 * @param tick_period   the call period of this function in milliseconds
 */
//go:linkname LvTickInc C.lv_tick_inc
func LvTickInc(tick_period c.Uint32T)

/**
 * Get the elapsed milliseconds since start up
 * @return          the elapsed milliseconds
 */
//go:linkname LvTickGet C.lv_tick_get
func LvTickGet() c.Uint32T

/**
 * Get the elapsed milliseconds since a previous time stamp
 * @param prev_tick     a previous time stamp (return value of lv_tick_get() )
 * @return              the elapsed milliseconds since 'prev_tick'
 */
//go:linkname LvTickElaps C.lv_tick_elaps
func LvTickElaps(prev_tick c.Uint32T) c.Uint32T

/**
 * Get the elapsed milliseconds between two time stamps
 * @param tick          a time stamp
 * @param prev_tick     a time stamp before `tick`
 * @return              the elapsed milliseconds between `prev_tick` and `tick`
 */
//go:linkname LvTickDiff C.lv_tick_diff
func LvTickDiff(tick c.Uint32T, prev_tick c.Uint32T) c.Uint32T

/**
 * Delay for the given milliseconds.
 * By default it's a blocking delay, but with `lv_delay_set_cb()`
 * a custom delay function can be set too
 * @param ms        the number of milliseconds to delay
 */
//go:linkname LvDelayMs C.lv_delay_ms
func LvDelayMs(ms c.Uint32T)

/**
 * Set a callback for a blocking delay
 * @param cb        pointer to a callback
 */
//go:linkname LvDelaySetCb C.lv_delay_set_cb
func LvDelaySetCb(cb LvDelayCbT)

/**
 * Set the custom callback for 'lv_tick_get'
 * @param cb        call this callback on 'lv_tick_get'
 */
//go:linkname LvTickSetCb C.lv_tick_set_cb
func LvTickSetCb(cb LvTickGetCbT)

/**
 * Get the custom callback for 'lv_tick_get'
 * @return      call this callback on 'lv_tick_get'
 */
//go:linkname LvTickGetCb C.lv_tick_get_cb
func LvTickGetCb() LvTickGetCbT

type LvLlNodeT c.Uint8T

/** Description of a linked list*/

type LvLlT struct {
	NSize c.Uint32T
	Head  *LvLlNodeT
	Tail  *LvLlNodeT
}

/**
 * Initialize linked list
 * @param ll_p pointer to lv_ll_t variable
 * @param node_size the size of 1 node in bytes
 */
// llgo:link (*LvLlT).LvLlInit C.lv_ll_init
func (recv_ *LvLlT) LvLlInit(node_size c.Uint32T) {
}

/**
 * Add a new head to a linked list
 * @param ll_p pointer to linked list
 * @return pointer to the new head
 */
// llgo:link (*LvLlT).LvLlInsHead C.lv_ll_ins_head
func (recv_ *LvLlT) LvLlInsHead() c.Pointer {
	return nil
}

/**
 * Insert a new node in front of the n_act node
 * @param ll_p pointer to linked list
 * @param n_act pointer a node
 * @return pointer to the new node
 */
// llgo:link (*LvLlT).LvLlInsPrev C.lv_ll_ins_prev
func (recv_ *LvLlT) LvLlInsPrev(n_act c.Pointer) c.Pointer {
	return nil
}

/**
 * Add a new tail to a linked list
 * @param ll_p pointer to linked list
 * @return pointer to the new tail
 */
// llgo:link (*LvLlT).LvLlInsTail C.lv_ll_ins_tail
func (recv_ *LvLlT) LvLlInsTail() c.Pointer {
	return nil
}

/**
 * Remove the node 'node_p' from 'll_p' linked list.
 * It does not free the memory of node.
 * @param ll_p pointer to the linked list of 'node_p'
 * @param node_p pointer to node in 'll_p' linked list
 */
// llgo:link (*LvLlT).LvLlRemove C.lv_ll_remove
func (recv_ *LvLlT) LvLlRemove(node_p c.Pointer) {
}

// llgo:link (*LvLlT).LvLlClearCustom C.lv_ll_clear_custom
func (recv_ *LvLlT) LvLlClearCustom(cleanup func(c.Pointer)) {
}

/**
 * Remove and free all elements from a linked list. The list remain valid but become empty.
 * @param ll_p pointer to linked list
 */
// llgo:link (*LvLlT).LvLlClear C.lv_ll_clear
func (recv_ *LvLlT) LvLlClear() {
}

/**
 * Move a node to a new linked list
 * @param ll_ori_p pointer to the original (old) linked list
 * @param ll_new_p pointer to the new linked list
 * @param node pointer to a node
 * @param head true: be the head in the new list
 *             false be the tail in the new list
 */
// llgo:link (*LvLlT).LvLlChgList C.lv_ll_chg_list
func (recv_ *LvLlT) LvLlChgList(ll_new_p *LvLlT, node c.Pointer, head bool) {
}

/**
 * Return with head node of the linked list
 * @param ll_p pointer to linked list
 * @return pointer to the head of 'll_p'
 */
// llgo:link (*LvLlT).LvLlGetHead C.lv_ll_get_head
func (recv_ *LvLlT) LvLlGetHead() c.Pointer {
	return nil
}

/**
 * Return with tail node of the linked list
 * @param ll_p pointer to linked list
 * @return pointer to the tail of 'll_p'
 */
// llgo:link (*LvLlT).LvLlGetTail C.lv_ll_get_tail
func (recv_ *LvLlT) LvLlGetTail() c.Pointer {
	return nil
}

/**
 * Return with the pointer of the next node after 'n_act'
 * @param ll_p pointer to linked list
 * @param n_act pointer a node
 * @return pointer to the next node
 */
// llgo:link (*LvLlT).LvLlGetNext C.lv_ll_get_next
func (recv_ *LvLlT) LvLlGetNext(n_act c.Pointer) c.Pointer {
	return nil
}

/**
 * Return with the pointer of the previous node after 'n_act'
 * @param ll_p pointer to linked list
 * @param n_act pointer a node
 * @return pointer to the previous node
 */
// llgo:link (*LvLlT).LvLlGetPrev C.lv_ll_get_prev
func (recv_ *LvLlT) LvLlGetPrev(n_act c.Pointer) c.Pointer {
	return nil
}

/**
 * Return the length of the linked list.
 * @param ll_p pointer to linked list
 * @return length of the linked list
 */
// llgo:link (*LvLlT).LvLlGetLen C.lv_ll_get_len
func (recv_ *LvLlT) LvLlGetLen() c.Uint32T {
	return 0
}

/**
 * Move a node before another node in the same linked list
 *
 * @param ll_p pointer to a linked list
 * @param n_act pointer to node to move
 * @param n_after pointer to a node which should be after `n_act`
 */
// llgo:link (*LvLlT).LvLlMoveBefore C.lv_ll_move_before
func (recv_ *LvLlT) LvLlMoveBefore(n_act c.Pointer, n_after c.Pointer) {
}

/**
 * Check if a linked list is empty
 * @param ll_p pointer to a linked list
 * @return true: the linked list is empty; false: not empty
 */
// llgo:link (*LvLlT).LvLlIsEmpty C.lv_ll_is_empty
func (recv_ *LvLlT) LvLlIsEmpty() bool {
	return false
}

// llgo:type C
type LvTimerCbT func(*LvTimerT)

// llgo:type C
type LvTimerHandlerResumeCbT func(c.Pointer)

/**
 * Call it periodically to handle lv_timers.
 * @return time till it needs to be run next (in ms)
 */
//go:linkname LvTimerHandler C.lv_timer_handler
func LvTimerHandler() c.Uint32T

/**
 * Call it in the super-loop of main() or threads. It will run lv_timer_handler()
 * with a given period in ms. You can use it with sleep or delay in OS environment.
 * This function is used to simplify the porting.
 * @param period the period for running lv_timer_handler()
 * @return the time after which it must be called again
 */
//go:linkname LvTimerHandlerRunInPeriod C.lv_timer_handler_run_in_period
func LvTimerHandlerRunInPeriod(period c.Uint32T) c.Uint32T

/**
 * Call it in the super-loop of main() or threads. It will automatically call lv_timer_handler() at the right time.
 * This function is used to simplify the porting.
 */
//go:linkname LvTimerPeriodicHandler C.lv_timer_periodic_handler
func LvTimerPeriodicHandler()

/**
 * Set the resume callback to the timer handler
 * @param cb the function to call when timer handler is resumed
 * @param data pointer to a resume data
 */
//go:linkname LvTimerHandlerSetResumeCb C.lv_timer_handler_set_resume_cb
func LvTimerHandlerSetResumeCb(cb LvTimerHandlerResumeCbT, data c.Pointer)

/**
 * Create an "empty" timer. It needs to be initialized with at least
 * `lv_timer_set_cb` and `lv_timer_set_period`
 * @return pointer to the created timer
 */
//go:linkname LvTimerCreateBasic C.lv_timer_create_basic
func LvTimerCreateBasic() *LvTimerT

/**
 * Create a new lv_timer
 * @param timer_xcb a callback to call periodically.
 *                 (the 'x' in the argument name indicates that it's not a fully generic function because it not follows
 *                  the `func_name(object, callback, ...)` convention)
 * @param period call period in ms unit
 * @param user_data custom parameter
 * @return pointer to the new timer
 */
//go:linkname LvTimerCreate C.lv_timer_create
func LvTimerCreate(timer_xcb LvTimerCbT, period c.Uint32T, user_data c.Pointer) *LvTimerT

/**
 * Delete a lv_timer
 * @param timer pointer to an lv_timer
 */
// llgo:link (*LvTimerT).LvTimerDelete C.lv_timer_delete
func (recv_ *LvTimerT) LvTimerDelete() {
}

/**
 * Pause a timer.
 * @param timer pointer to an lv_timer
 */
// llgo:link (*LvTimerT).LvTimerPause C.lv_timer_pause
func (recv_ *LvTimerT) LvTimerPause() {
}

/**
 * Resume a timer.
 * @param timer pointer to an lv_timer
 */
// llgo:link (*LvTimerT).LvTimerResume C.lv_timer_resume
func (recv_ *LvTimerT) LvTimerResume() {
}

/**
 * Set the callback to the timer (the function to call periodically)
 * @param timer pointer to a timer
 * @param timer_cb the function to call periodically
 */
// llgo:link (*LvTimerT).LvTimerSetCb C.lv_timer_set_cb
func (recv_ *LvTimerT) LvTimerSetCb(timer_cb LvTimerCbT) {
}

/**
 * Set new period for a lv_timer
 * @param timer pointer to a lv_timer
 * @param period the new period
 */
// llgo:link (*LvTimerT).LvTimerSetPeriod C.lv_timer_set_period
func (recv_ *LvTimerT) LvTimerSetPeriod(period c.Uint32T) {
}

/**
 * Make a lv_timer ready. It will not wait its period.
 * @param timer pointer to a lv_timer.
 */
// llgo:link (*LvTimerT).LvTimerReady C.lv_timer_ready
func (recv_ *LvTimerT) LvTimerReady() {
}

/**
 * Set the number of times a timer will repeat.
 * @param timer pointer to a lv_timer.
 * @param repeat_count -1 : infinity;  0 : stop ;  n>0: residual times
 */
// llgo:link (*LvTimerT).LvTimerSetRepeatCount C.lv_timer_set_repeat_count
func (recv_ *LvTimerT) LvTimerSetRepeatCount(repeat_count c.Int32T) {
}

/**
 * Set whether a lv_timer will be deleted automatically when it is called `repeat_count` times.
 * @param timer pointer to a lv_timer.
 * @param auto_delete true: auto delete; false: timer will be paused when it is called `repeat_count` times.
 */
// llgo:link (*LvTimerT).LvTimerSetAutoDelete C.lv_timer_set_auto_delete
func (recv_ *LvTimerT) LvTimerSetAutoDelete(auto_delete bool) {
}

/**
 * Set custom parameter to the lv_timer.
 * @param timer pointer to a lv_timer.
 * @param user_data custom parameter
 */
// llgo:link (*LvTimerT).LvTimerSetUserData C.lv_timer_set_user_data
func (recv_ *LvTimerT) LvTimerSetUserData(user_data c.Pointer) {
}

/**
 * Reset a lv_timer.
 * It will be called the previously set period milliseconds later.
 * @param timer pointer to a lv_timer.
 */
// llgo:link (*LvTimerT).LvTimerReset C.lv_timer_reset
func (recv_ *LvTimerT) LvTimerReset() {
}

/**
 * Enable or disable the whole lv_timer handling
 * @param en true: lv_timer handling is running, false: lv_timer handling is suspended
 */
//go:linkname LvTimerEnable C.lv_timer_enable
func LvTimerEnable(en bool)

/**
 * Get idle percentage
 * @return the lv_timer idle in percentage
 */
//go:linkname LvTimerGetIdle C.lv_timer_get_idle
func LvTimerGetIdle() c.Uint32T

/**
 * Get the time remaining until the next timer will run
 * @return the time remaining in ms
 */
//go:linkname LvTimerGetTimeUntilNext C.lv_timer_get_time_until_next
func LvTimerGetTimeUntilNext() c.Uint32T

/**
 * Iterate through the timers
 * @param timer NULL to start iteration or the previous return value to get the next timer
 * @return the next timer or NULL if there is no more timer
 */
// llgo:link (*LvTimerT).LvTimerGetNext C.lv_timer_get_next
func (recv_ *LvTimerT) LvTimerGetNext() *LvTimerT {
	return nil
}

/**
 * Get the user_data passed when the timer was created
 * @param timer pointer to the lv_timer
 * @return pointer to the user_data
 */
// llgo:link (*LvTimerT).LvTimerGetUserData C.lv_timer_get_user_data
func (recv_ *LvTimerT) LvTimerGetUserData() c.Pointer {
	return nil
}

/**
 * Get the pause state of a timer
 * @param timer pointer to a lv_timer
 * @return true: timer is paused; false: timer is running
 */
// llgo:link (*LvTimerT).LvTimerGetPaused C.lv_timer_get_paused
func (recv_ *LvTimerT) LvTimerGetPaused() bool {
	return false
}

type LvAnimEnableT bool

// llgo:type C
type LvAnimPathCbT func(*LvAnimT) c.Int32T

// llgo:type C
type LvAnimExecXcbT func(c.Pointer, c.Int32T)

// llgo:type C
type LvAnimCustomExecCbT func(*LvAnimT, c.Int32T)

// llgo:type C
type LvAnimCompletedCbT func(*LvAnimT)

// llgo:type C
type LvAnimStartCbT func(*LvAnimT)

// llgo:type C
type LvAnimGetValueCbT func(*LvAnimT) c.Int32T

// llgo:type C
type LvAnimDeletedCbT func(*LvAnimT)

/** Parameter used when path is custom_bezier */

type LvAnimBezier3ParaT struct {
	X1 c.Int16T
	Y1 c.Int16T
	X2 c.Int16T
	Y2 c.Int16T
}

type X_lvAnimPathParaT struct {
	Unused [8]uint8
}

/**
 * Initialize an animation variable.
 * E.g.:
 * lv_anim_t a;
 * lv_anim_init(&a);
 * lv_anim_set_...(&a);
 * lv_anim_start(&a);
 * @param a     pointer to an `lv_anim_t` variable to initialize
 */
// llgo:link (*LvAnimT).LvAnimInit C.lv_anim_init
func (recv_ *LvAnimT) LvAnimInit() {
}

/**
 * Set a variable to animate
 * @param a     pointer to an initialized `lv_anim_t` variable
 * @param var   pointer to a variable to animate
 */
// llgo:link (*LvAnimT).LvAnimSetVar C.lv_anim_set_var
func (recv_ *LvAnimT) LvAnimSetVar(var_ c.Pointer) {
}

/**
 * Set a function to animate `var`
 * @param a         pointer to an initialized `lv_anim_t` variable
 * @param exec_cb   a function to execute during animation
 *                  LVGL's built-in functions can be used.
 *                  E.g. lv_obj_set_x
 */
// llgo:link (*LvAnimT).LvAnimSetExecCb C.lv_anim_set_exec_cb
func (recv_ *LvAnimT) LvAnimSetExecCb(exec_cb LvAnimExecXcbT) {
}

/**
 * Set the duration of an animation
 * @param a         pointer to an initialized `lv_anim_t` variable
 * @param duration  duration of the animation in milliseconds
 */
// llgo:link (*LvAnimT).LvAnimSetDuration C.lv_anim_set_duration
func (recv_ *LvAnimT) LvAnimSetDuration(duration c.Uint32T) {
}

/**
 * Set a delay before starting the animation
 * @param a         pointer to an initialized `lv_anim_t` variable
 * @param delay     delay before the animation in milliseconds
 */
// llgo:link (*LvAnimT).LvAnimSetDelay C.lv_anim_set_delay
func (recv_ *LvAnimT) LvAnimSetDelay(delay c.Uint32T) {
}

/**
 * Resumes a paused animation
 * @param a         pointer to an initialized `lv_anim_t` variable
 */
// llgo:link (*LvAnimT).LvAnimResume C.lv_anim_resume
func (recv_ *LvAnimT) LvAnimResume() {
}

/**
 * Pauses the animation
 * @param a         pointer to an initialized `lv_anim_t` variable
 */
// llgo:link (*LvAnimT).LvAnimPause C.lv_anim_pause
func (recv_ *LvAnimT) LvAnimPause() {
}

/**
 * Pauses the animation for ms milliseconds
 * @param a         pointer to an initialized `lv_anim_t` variable
 * @param ms        the pause time in milliseconds
 */
// llgo:link (*LvAnimT).LvAnimPauseFor C.lv_anim_pause_for
func (recv_ *LvAnimT) LvAnimPauseFor(ms c.Uint32T) {
}

/**
 * Check if the animation is paused
 * @param a         pointer to an initialized `lv_anim_t` variable
 * @return          true if the animation is paused else false
 */
// llgo:link (*LvAnimT).LvAnimIsPaused C.lv_anim_is_paused
func (recv_ *LvAnimT) LvAnimIsPaused() bool {
	return false
}

/**
 * Set the start and end values of an animation
 * @param a         pointer to an initialized `lv_anim_t` variable
 * @param start     the start value
 * @param end       the end value
 */
// llgo:link (*LvAnimT).LvAnimSetValues C.lv_anim_set_values
func (recv_ *LvAnimT) LvAnimSetValues(start c.Int32T, end c.Int32T) {
}

/**
 * Similar to `lv_anim_set_exec_cb` but `lv_anim_custom_exec_cb_t` receives
 * `lv_anim_t * ` as its first parameter instead of `void *`.
 * This function might be used when LVGL is bound to other languages because
 * it's more consistent to have `lv_anim_t *` as first parameter.
 * @param a         pointer to an initialized `lv_anim_t` variable
 * @param exec_cb   a function to execute.
 */
// llgo:link (*LvAnimT).LvAnimSetCustomExecCb C.lv_anim_set_custom_exec_cb
func (recv_ *LvAnimT) LvAnimSetCustomExecCb(exec_cb LvAnimCustomExecCbT) {
}

/**
 * Set the path (curve) of the animation.
 * @param a         pointer to an initialized `lv_anim_t` variable
 * @param path_cb a function to set the current value of the animation.
 */
// llgo:link (*LvAnimT).LvAnimSetPathCb C.lv_anim_set_path_cb
func (recv_ *LvAnimT) LvAnimSetPathCb(path_cb LvAnimPathCbT) {
}

/**
 * Set a function call when the animation really starts (considering `delay`)
 * @param a         pointer to an initialized `lv_anim_t` variable
 * @param start_cb  a function call when the animation starts
 */
// llgo:link (*LvAnimT).LvAnimSetStartCb C.lv_anim_set_start_cb
func (recv_ *LvAnimT) LvAnimSetStartCb(start_cb LvAnimStartCbT) {
}

/**
 * Set a function to use the current value of the variable and make start and end value
 * relative to the returned current value.
 * @param a             pointer to an initialized `lv_anim_t` variable
 * @param get_value_cb  a function call when the animation starts
 */
// llgo:link (*LvAnimT).LvAnimSetGetValueCb C.lv_anim_set_get_value_cb
func (recv_ *LvAnimT) LvAnimSetGetValueCb(get_value_cb LvAnimGetValueCbT) {
}

/**
 * Set a function call when the animation is completed
 * @param a             pointer to an initialized `lv_anim_t` variable
 * @param completed_cb  a function call when the animation is fully completed
 */
// llgo:link (*LvAnimT).LvAnimSetCompletedCb C.lv_anim_set_completed_cb
func (recv_ *LvAnimT) LvAnimSetCompletedCb(completed_cb LvAnimCompletedCbT) {
}

/**
 * Set a function call when the animation is deleted.
 * @param a         pointer to an initialized `lv_anim_t` variable
 * @param deleted_cb  a function call when the animation is deleted
 */
// llgo:link (*LvAnimT).LvAnimSetDeletedCb C.lv_anim_set_deleted_cb
func (recv_ *LvAnimT) LvAnimSetDeletedCb(deleted_cb LvAnimDeletedCbT) {
}

/**
 * Make the animation to play back to when the forward direction is ready
 * @param a         pointer to an initialized `lv_anim_t` variable
 * @param duration  duration of playback animation in milliseconds. 0: disable playback
 */
// llgo:link (*LvAnimT).LvAnimSetReverseDuration C.lv_anim_set_reverse_duration
func (recv_ *LvAnimT) LvAnimSetReverseDuration(duration c.Uint32T) {
}

/**
 * Legacy `lv_anim_set_reverse_time` API will be removed soon, use `lv_anim_set_reverse_duration` instead.
 */
// llgo:link (*LvAnimT).LvAnimSetReverseTime C.lv_anim_set_reverse_time
func (recv_ *LvAnimT) LvAnimSetReverseTime(duration c.Uint32T) {
}

/**
 * Make the animation to play back to when the forward direction is ready
 * @param a         pointer to an initialized `lv_anim_t` variable
 * @param delay     delay in milliseconds before starting the playback animation.
 */
// llgo:link (*LvAnimT).LvAnimSetReverseDelay C.lv_anim_set_reverse_delay
func (recv_ *LvAnimT) LvAnimSetReverseDelay(delay c.Uint32T) {
}

/**
 * Make the animation repeat itself.
 * @param a         pointer to an initialized `lv_anim_t` variable
 * @param cnt       repeat count or `LV_ANIM_REPEAT_INFINITE` for infinite repetition. 0: to disable repetition.
 */
// llgo:link (*LvAnimT).LvAnimSetRepeatCount C.lv_anim_set_repeat_count
func (recv_ *LvAnimT) LvAnimSetRepeatCount(cnt c.Uint32T) {
}

/**
 * Set a delay before repeating the animation.
 * @param a         pointer to an initialized `lv_anim_t` variable
 * @param delay     delay in milliseconds before repeating the animation.
 */
// llgo:link (*LvAnimT).LvAnimSetRepeatDelay C.lv_anim_set_repeat_delay
func (recv_ *LvAnimT) LvAnimSetRepeatDelay(delay c.Uint32T) {
}

/**
 * Set a whether the animation's should be applied immediately or only when the delay expired.
 * @param a         pointer to an initialized `lv_anim_t` variable
 * @param en        true: apply the start value immediately in `lv_anim_start`;
 *                  false: apply the start value only when `delay` ms is elapsed and the animations really starts
 */
// llgo:link (*LvAnimT).LvAnimSetEarlyApply C.lv_anim_set_early_apply
func (recv_ *LvAnimT) LvAnimSetEarlyApply(en bool) {
}

/**
 * Set the custom user data field of the animation.
 * @param a           pointer to an initialized `lv_anim_t` variable
 * @param user_data   pointer to the new user_data.
 */
// llgo:link (*LvAnimT).LvAnimSetUserData C.lv_anim_set_user_data
func (recv_ *LvAnimT) LvAnimSetUserData(user_data c.Pointer) {
}

/**
 * Set parameter for cubic bezier path
 * @param a         pointer to an initialized `lv_anim_t` variable
 * @param x1        first control point X
 * @param y1        first control point Y
 * @param x2        second control point X
 * @param y2        second control point Y
 */
// llgo:link (*LvAnimT).LvAnimSetBezier3Param C.lv_anim_set_bezier3_param
func (recv_ *LvAnimT) LvAnimSetBezier3Param(x1 c.Int16T, y1 c.Int16T, x2 c.Int16T, y2 c.Int16T) {
}

/**
 * Create an animation
 * @param a         an initialized 'anim_t' variable. Not required after call.
 * @return          pointer to the created animation (different from the `a` parameter)
 */
// llgo:link (*LvAnimT).LvAnimStart C.lv_anim_start
func (recv_ *LvAnimT) LvAnimStart() *LvAnimT {
	return nil
}

/**
 * Get a delay before starting the animation
 * @param a pointer to an initialized `lv_anim_t` variable
 * @return delay before the animation in milliseconds
 */
// llgo:link (*LvAnimT).LvAnimGetDelay C.lv_anim_get_delay
func (recv_ *LvAnimT) LvAnimGetDelay() c.Uint32T {
	return 0
}

/**
 * Get the time used to play the animation.
 * @param a pointer to an animation.
 * @return the play time in milliseconds.
 */
// llgo:link (*LvAnimT).LvAnimGetPlaytime C.lv_anim_get_playtime
func (recv_ *LvAnimT) LvAnimGetPlaytime() c.Uint32T {
	return 0
}

/**
 * Get the duration of an animation
 * @param a         pointer to an initialized `lv_anim_t` variable
 * @return the duration of the animation in milliseconds
 */
// llgo:link (*LvAnimT).LvAnimGetTime C.lv_anim_get_time
func (recv_ *LvAnimT) LvAnimGetTime() c.Uint32T {
	return 0
}

/**
 * Get the repeat count of the animation.
 * @param a         pointer to an initialized `lv_anim_t` variable
 * @return the repeat count or `LV_ANIM_REPEAT_INFINITE` for infinite repetition. 0: disabled repetition.
 */
// llgo:link (*LvAnimT).LvAnimGetRepeatCount C.lv_anim_get_repeat_count
func (recv_ *LvAnimT) LvAnimGetRepeatCount() c.Uint32T {
	return 0
}

/**
 * Get the user_data field of the animation
 * @param   a pointer to an initialized `lv_anim_t` variable
 * @return  the pointer to the custom user_data of the animation
 */
// llgo:link (*LvAnimT).LvAnimGetUserData C.lv_anim_get_user_data
func (recv_ *LvAnimT) LvAnimGetUserData() c.Pointer {
	return nil
}

/**
 * Delete animation(s) of a variable with a given animator function
 * @param var       pointer to variable
 * @param exec_cb   a function pointer which is animating 'var',
 *                  or NULL to ignore it and delete all the animations of 'var
 * @return          true: at least 1 animation is deleted, false: no animation is deleted
 */
//go:linkname LvAnimDelete C.lv_anim_delete
func LvAnimDelete(var_ c.Pointer, exec_cb LvAnimExecXcbT) bool

/**
 * Delete all the animations
 */
//go:linkname LvAnimDeleteAll C.lv_anim_delete_all
func LvAnimDeleteAll()

/**
 * Get the animation of a variable and its `exec_cb`.
 * @param var       pointer to variable
 * @param exec_cb   a function pointer which is animating 'var', or NULL to return first matching 'var'
 * @return          pointer to the animation.
 */
//go:linkname LvAnimGet C.lv_anim_get
func LvAnimGet(var_ c.Pointer, exec_cb LvAnimExecXcbT) *LvAnimT

/**
 * Get global animation refresher timer.
 * @return pointer to the animation refresher timer.
 */
//go:linkname LvAnimGetTimer C.lv_anim_get_timer
func LvAnimGetTimer() *LvTimerT

/**
 * Delete an animation by getting the animated variable from `a`.
 * Only animations with `exec_cb` will be deleted.
 * This function exists because it's logical that all anim. functions receives an
 * `lv_anim_t` as their first parameter. It's not practical in C but might make
 * the API more consequent and makes easier to generate bindings.
 * @param a         pointer to an animation.
 * @param exec_cb   a function pointer which is animating 'var',
 *                  or NULL to ignore it and delete all the animations of 'var
 * @return          true: at least 1 animation is deleted, false: no animation is deleted
 */
// llgo:link (*LvAnimT).LvAnimCustomDelete C.lv_anim_custom_delete
func (recv_ *LvAnimT) LvAnimCustomDelete(exec_cb LvAnimCustomExecCbT) bool {
	return false
}

/**
 * Get the animation of a variable and its `exec_cb`.
 * This function exists because it's logical that all anim. functions receives an
 * `lv_anim_t` as their first parameter. It's not practical in C but might make
 * the API more consequent and makes easier to generate bindings.
 * @param a         pointer to an animation.
 * @param exec_cb   a function pointer which is animating 'var', or NULL to return first matching 'var'
 * @return          pointer to the animation.
 */
// llgo:link (*LvAnimT).LvAnimCustomGet C.lv_anim_custom_get
func (recv_ *LvAnimT) LvAnimCustomGet(exec_cb LvAnimCustomExecCbT) *LvAnimT {
	return nil
}

/**
 * Get the number of currently running animations
 * @return      the number of running animations
 */
//go:linkname LvAnimCountRunning C.lv_anim_count_running
func LvAnimCountRunning() c.Uint16T

/**
 * Store the speed as a special value which can be used as time in animations.
 * It will be converted to time internally based on the start and end values.
 * The return value can be used as a constant with multiple animations
 * and let LVGL convert the speed to time based on the actual values.
 * LIMITATION: the max time stored this way can be 10,000 ms.
 * @param speed         the speed of the animation in with unit / sec resolution in 0..10k range
 * @return              a special value which can be used as an animation time
 * @note                internally speed is stored as 10 unit/sec
 */
//go:linkname LvAnimSpeed C.lv_anim_speed
func LvAnimSpeed(speed c.Uint32T) c.Uint32T

/**
 * Store the speed as a special value which can be used as time in animations.
 * It will be converted to time internally based on the start and end values.
 * The return value can be used as a constant with multiple animations
 * and let LVGL convert the speed to time based on the actual values.
 * @param speed         the speed of the animation in as unit / sec resolution in 0..10k range
 * @param min_time      the minimum time in 0..10k range
 * @param max_time      the maximum time in 0..10k range
 * @return              a special value in where all three values are stored and can be used as an animation time
 * @note                internally speed is stored as 10 unit/sec
 * @note                internally min/max_time are stored with 10 ms unit
 *
 */
//go:linkname LvAnimSpeedClamped C.lv_anim_speed_clamped
func LvAnimSpeedClamped(speed c.Uint32T, min_time c.Uint32T, max_time c.Uint32T) c.Uint32T

/**
 * Resolve the speed (created with `lv_anim_speed` or `lv_anim_speed_clamped`) to time
 * based on start and end values.
 * @param speed     return values of `lv_anim_speed` or `lv_anim_speed_clamped`
 * @param start     the start value of the animation
 * @param end       the end value of the animation
 * @return          the time required to get from `start` to `end` with the given `speed` setting
 */
//go:linkname LvAnimResolveSpeed C.lv_anim_resolve_speed
func LvAnimResolveSpeed(speed c.Uint32T, start c.Int32T, end c.Int32T) c.Uint32T

/**
 * Calculate the time of an animation based on its speed, start and end values.
 * It simpler than `lv_anim_speed` or `lv_anim_speed_clamped` as it converts
 * speed, start, and end to a time immediately.
 * As it's simpler there is no limit on the maximum time.
 * @param speed         the speed of the animation
 * @param start         the start value
 * @param end           the end value
 * @return              the time of the animation in milliseconds
 */
//go:linkname LvAnimSpeedToTime C.lv_anim_speed_to_time
func LvAnimSpeedToTime(speed c.Uint32T, start c.Int32T, end c.Int32T) c.Uint32T

/**
 * Manually refresh the state of the animations.
 * Useful to make the animations running in a blocking process where
 * `lv_timer_handler` can't run for a while.
 * Shouldn't be used directly because it is called in `lv_refr_now()`.
 */
//go:linkname LvAnimRefrNow C.lv_anim_refr_now
func LvAnimRefrNow()

/**
 * Calculate the current value of an animation applying linear characteristic
 * @param a     pointer to an animation
 * @return      the current value to set
 */
// llgo:link (*LvAnimT).LvAnimPathLinear C.lv_anim_path_linear
func (recv_ *LvAnimT) LvAnimPathLinear() c.Int32T {
	return 0
}

/**
 * Calculate the current value of an animation slowing down the start phase
 * @param a     pointer to an animation
 * @return      the current value to set
 */
// llgo:link (*LvAnimT).LvAnimPathEaseIn C.lv_anim_path_ease_in
func (recv_ *LvAnimT) LvAnimPathEaseIn() c.Int32T {
	return 0
}

/**
 * Calculate the current value of an animation slowing down the end phase
 * @param a     pointer to an animation
 * @return      the current value to set
 */
// llgo:link (*LvAnimT).LvAnimPathEaseOut C.lv_anim_path_ease_out
func (recv_ *LvAnimT) LvAnimPathEaseOut() c.Int32T {
	return 0
}

/**
 * Calculate the current value of an animation applying an "S" characteristic (cosine)
 * @param a     pointer to an animation
 * @return      the current value to set
 */
// llgo:link (*LvAnimT).LvAnimPathEaseInOut C.lv_anim_path_ease_in_out
func (recv_ *LvAnimT) LvAnimPathEaseInOut() c.Int32T {
	return 0
}

/**
 * Calculate the current value of an animation with overshoot at the end
 * @param a     pointer to an animation
 * @return      the current value to set
 */
// llgo:link (*LvAnimT).LvAnimPathOvershoot C.lv_anim_path_overshoot
func (recv_ *LvAnimT) LvAnimPathOvershoot() c.Int32T {
	return 0
}

/**
 * Calculate the current value of an animation with 3 bounces
 * @param a     pointer to an animation
 * @return      the current value to set
 */
// llgo:link (*LvAnimT).LvAnimPathBounce C.lv_anim_path_bounce
func (recv_ *LvAnimT) LvAnimPathBounce() c.Int32T {
	return 0
}

/**
 * Calculate the current value of an animation applying step characteristic.
 * (Set end value on the end of the animation)
 * @param a     pointer to an animation
 * @return      the current value to set
 */
// llgo:link (*LvAnimT).LvAnimPathStep C.lv_anim_path_step
func (recv_ *LvAnimT) LvAnimPathStep() c.Int32T {
	return 0
}

/**
 * A custom cubic bezier animation path, need to specify cubic-parameters in a->parameter.bezier3
 * @param a     pointer to an animation
 * @return      the current value to set
 */
// llgo:link (*LvAnimT).LvAnimPathCustomBezier3 C.lv_anim_path_custom_bezier3
func (recv_ *LvAnimT) LvAnimPathCustomBezier3() c.Int32T {
	return 0
}

//go:linkname LvSnprintf C.lv_snprintf
func LvSnprintf(buffer *c.Char, count c.SizeT, format *c.Char, __llgo_va_list ...interface{}) c.Int

//go:linkname LvVsnprintf C.lv_vsnprintf
func LvVsnprintf(buffer *c.Char, count c.SizeT, format *c.Char, va c.VaList) c.Int

type LvTextFlagT c.Int

const (
	LV_TEXT_FLAG_NONE      LvTextFlagT = 0
	LV_TEXT_FLAG_EXPAND    LvTextFlagT = 1
	LV_TEXT_FLAG_FIT       LvTextFlagT = 2
	LV_TEXT_FLAG_BREAK_ALL LvTextFlagT = 4
	LV_TEXT_FLAG_RECOLOR   LvTextFlagT = 8
)

type LvTextAlignT c.Int

const (
	LV_TEXT_ALIGN_AUTO   LvTextAlignT = 0
	LV_TEXT_ALIGN_LEFT   LvTextAlignT = 1
	LV_TEXT_ALIGN_CENTER LvTextAlignT = 2
	LV_TEXT_ALIGN_RIGHT  LvTextAlignT = 3
)

type LvTextCmdStateT c.Int

const (
	LV_TEXT_CMD_STATE_WAIT LvTextCmdStateT = 0
	LV_TEXT_CMD_STATE_PAR  LvTextCmdStateT = 1
	LV_TEXT_CMD_STATE_IN   LvTextCmdStateT = 2
)

/**
 * Get size of a text
 * @param size_res pointer to a 'point_t' variable to store the result
 * @param text pointer to a text
 * @param font pointer to font of the text
 * @param letter_space letter space of the text
 * @param line_space line space of the text
 * @param max_width max width of the text (break the lines to fit this size). Set COORD_MAX to avoid
 * @param flag settings for the text from ::lv_text_flag_t

 * line breaks
 */
// llgo:link (*LvPointT).LvTextGetSize C.lv_text_get_size
func (recv_ *LvPointT) LvTextGetSize(text *c.Char, font *LvFontT, letter_space c.Int32T, line_space c.Int32T, max_width c.Int32T, flag LvTextFlagT) {
}

/**
 * Give the length of a text with a given font
 * @param txt a '\0' terminate string
 * @param length length of 'txt' in byte count and not characters (Á is 1 character but 2 bytes in
 * UTF-8)
 * @param font pointer to a font
 * @param letter_space letter space
 * @return length of a char_num long text
 */
//go:linkname LvTextGetWidth C.lv_text_get_width
func LvTextGetWidth(txt *c.Char, length c.Uint32T, font *LvFontT, letter_space c.Int32T) c.Int32T

/**
 * Give the length of a text with a given font with text flags
 * @param txt a '\0' terminate string
 * @param length length of 'txt' in byte count and not characters (Á is 1 character but 2 bytes in
 * UTF-8)
 * @param font pointer to a font
 * @param letter_space letter space
 * @param flags settings for the text from ::lv_text_flag_t
 * @return length of a char_num long text
 */
//go:linkname LvTextGetWidthWithFlags C.lv_text_get_width_with_flags
func LvTextGetWidthWithFlags(txt *c.Char, length c.Uint32T, font *LvFontT, letter_space c.Int32T, flags LvTextFlagT) c.Int32T

/**
 * Check if c is command state
 * @param state
 * @param c
 * @return True if c is state
 */
// llgo:link (*LvTextCmdStateT).LvTextIsCmd C.lv_text_is_cmd
func (recv_ *LvTextCmdStateT) LvTextIsCmd(c c.Uint32T) bool {
	return false
}

type LvBaseDirT c.Int

const (
	LV_BASE_DIR_LTR     LvBaseDirT = 0
	LV_BASE_DIR_RTL     LvBaseDirT = 1
	LV_BASE_DIR_AUTO    LvBaseDirT = 2
	LV_BASE_DIR_NEUTRAL LvBaseDirT = 32
	LV_BASE_DIR_WEAK    LvBaseDirT = 33
)

type LvGradDirT c.Int

const (
	LV_GRAD_DIR_NONE    LvGradDirT = 0
	LV_GRAD_DIR_VER     LvGradDirT = 1
	LV_GRAD_DIR_HOR     LvGradDirT = 2
	LV_GRAD_DIR_LINEAR  LvGradDirT = 3
	LV_GRAD_DIR_RADIAL  LvGradDirT = 4
	LV_GRAD_DIR_CONICAL LvGradDirT = 5
)

type LvGradExtendT c.Int

const (
	LV_GRAD_EXTEND_PAD     LvGradExtendT = 0
	LV_GRAD_EXTEND_REPEAT  LvGradExtendT = 1
	LV_GRAD_EXTEND_REFLECT LvGradExtendT = 2
)

/** A gradient stop definition.
 *  This matches a color and a position in a virtual 0-255 scale.
 */

type LvGradStopT struct {
	Color LvColorT
	Opa   LvOpaT
	Frac  c.Uint8T
}

/** A descriptor of a gradient. */

type LvGradDscT struct {
	Stops      [2]LvGradStopT
	StopsCount c.Uint8T
	Dir        LvGradDirT
	Extend     LvGradExtendT
}

/**
 * Initialize gradient color map from a table
 * @param grad      pointer to a gradient descriptor
 * @param colors    color array
 * @param fracs     position array (0..255): if NULL, then colors are distributed evenly
 * @param opa       opacity array: if NULL, then LV_OPA_COVER is assumed
 * @param num_stops number of gradient stops (1..LV_GRADIENT_MAX_STOPS)
 */
// llgo:link (*LvGradDscT).LvGradInitStops C.lv_grad_init_stops
func (recv_ *LvGradDscT) LvGradInitStops(colors *LvColorT, opa *LvOpaT, fracs *c.Uint8T, num_stops c.Int) {
}

/**
 * Helper function to initialize a horizontal gradient.
 * @param dsc      gradient descriptor
 */
// llgo:link (*LvGradDscT).LvGradHorizontalInit C.lv_grad_horizontal_init
func (recv_ *LvGradDscT) LvGradHorizontalInit() {
}

/**
 * Helper function to initialize a vertical gradient.
 * @param dsc      gradient descriptor
 */
// llgo:link (*LvGradDscT).LvGradVerticalInit C.lv_grad_vertical_init
func (recv_ *LvGradDscT) LvGradVerticalInit() {
}

// llgo:type C
type LvLayoutUpdateCbT func(*LvObjT, c.Pointer)
type LvLayoutT c.Int

const (
	LV_LAYOUT_NONE LvLayoutT = 0
	LV_LAYOUT_FLEX LvLayoutT = 1
	LV_LAYOUT_GRID LvLayoutT = 2
	LV_LAYOUT_LAST LvLayoutT = 3
)

/**
 * Register a new layout
 * @param cb        the layout update callback
 * @param user_data custom data that will be passed to `cb`
 * @return          the ID of the new layout
 */
//go:linkname LvLayoutRegister C.lv_layout_register
func LvLayoutRegister(cb LvLayoutUpdateCbT, user_data c.Pointer) c.Uint32T

type LvFlexAlignT c.Int

const (
	LV_FLEX_ALIGN_START         LvFlexAlignT = 0
	LV_FLEX_ALIGN_END           LvFlexAlignT = 1
	LV_FLEX_ALIGN_CENTER        LvFlexAlignT = 2
	LV_FLEX_ALIGN_SPACE_EVENLY  LvFlexAlignT = 3
	LV_FLEX_ALIGN_SPACE_AROUND  LvFlexAlignT = 4
	LV_FLEX_ALIGN_SPACE_BETWEEN LvFlexAlignT = 5
)

type LvFlexFlowT c.Int

const (
	LV_FLEX_FLOW_ROW                 LvFlexFlowT = 0
	LV_FLEX_FLOW_COLUMN              LvFlexFlowT = 1
	LV_FLEX_FLOW_ROW_WRAP            LvFlexFlowT = 4
	LV_FLEX_FLOW_ROW_REVERSE         LvFlexFlowT = 8
	LV_FLEX_FLOW_ROW_WRAP_REVERSE    LvFlexFlowT = 12
	LV_FLEX_FLOW_COLUMN_WRAP         LvFlexFlowT = 5
	LV_FLEX_FLOW_COLUMN_REVERSE      LvFlexFlowT = 9
	LV_FLEX_FLOW_COLUMN_WRAP_REVERSE LvFlexFlowT = 13
)

/**
 * Initialize a flex layout to default values
 */
//go:linkname LvFlexInit C.lv_flex_init
func LvFlexInit()

/**
 * Set how the item should flow
 * @param obj pointer to an object. The parent must have flex layout else nothing will happen.
 * @param flow an element of `lv_flex_flow_t`.
 */
// llgo:link (*LvObjT).LvObjSetFlexFlow C.lv_obj_set_flex_flow
func (recv_ *LvObjT) LvObjSetFlexFlow(flow LvFlexFlowT) {
}

/**
 * Set how to place (where to align) the items and tracks
 * @param obj pointer to an object. The parent must have flex layout else nothing will happen.
 * @param main_place where to place the items on main axis (in their track). Any value of `lv_flex_align_t`.
 * @param cross_place where to place the item in their track on the cross axis. `LV_FLEX_ALIGN_START/END/CENTER`
 * @param track_cross_place where to place the tracks in the cross direction. Any value of `lv_flex_align_t`.
 */
// llgo:link (*LvObjT).LvObjSetFlexAlign C.lv_obj_set_flex_align
func (recv_ *LvObjT) LvObjSetFlexAlign(main_place LvFlexAlignT, cross_place LvFlexAlignT, track_cross_place LvFlexAlignT) {
}

/**
 * Sets the width or height (on main axis) to grow the object in order fill the free space
 * @param obj pointer to an object. The parent must have flex layout else nothing will happen.
 * @param grow a value to set how much free space to take proportionally to other growing items.
 */
// llgo:link (*LvObjT).LvObjSetFlexGrow C.lv_obj_set_flex_grow
func (recv_ *LvObjT) LvObjSetFlexGrow(grow c.Uint8T) {
}

type LvGridAlignT c.Int

const (
	LV_GRID_ALIGN_START         LvGridAlignT = 0
	LV_GRID_ALIGN_CENTER        LvGridAlignT = 1
	LV_GRID_ALIGN_END           LvGridAlignT = 2
	LV_GRID_ALIGN_STRETCH       LvGridAlignT = 3
	LV_GRID_ALIGN_SPACE_EVENLY  LvGridAlignT = 4
	LV_GRID_ALIGN_SPACE_AROUND  LvGridAlignT = 5
	LV_GRID_ALIGN_SPACE_BETWEEN LvGridAlignT = 6
)

/**********************
 * GLOBAL PROTOTYPES
 **********************/
//go:linkname LvGridInit C.lv_grid_init
func LvGridInit()

// llgo:link (*LvObjT).LvObjSetGridDscArray C.lv_obj_set_grid_dsc_array
func (recv_ *LvObjT) LvObjSetGridDscArray(col_dsc *c.Int32T, row_dsc *c.Int32T) {
}

// llgo:link (*LvObjT).LvObjSetGridAlign C.lv_obj_set_grid_align
func (recv_ *LvObjT) LvObjSetGridAlign(column_align LvGridAlignT, row_align LvGridAlignT) {
}

/**
 * Set the cell of an object. The object's parent needs to have grid layout, else nothing will happen
 * @param obj pointer to an object
 * @param column_align the vertical alignment in the cell. `LV_GRID_START/END/CENTER/STRETCH`
 * @param col_pos column ID
 * @param col_span number of columns to take (>= 1)
 * @param row_align the horizontal alignment in the cell. `LV_GRID_START/END/CENTER/STRETCH`
 * @param row_pos row ID
 * @param row_span number of rows to take (>= 1)
 */
// llgo:link (*LvObjT).LvObjSetGridCell C.lv_obj_set_grid_cell
func (recv_ *LvObjT) LvObjSetGridCell(column_align LvGridAlignT, col_pos c.Int32T, col_span c.Int32T, row_align LvGridAlignT, row_pos c.Int32T, row_span c.Int32T) {
}

/**
 * Just a wrapper to `LV_GRID_FR` for bindings.
 */
//go:linkname LvGridFr C.lv_grid_fr
func LvGridFr(x c.Uint8T) c.Int32T

type LvBlendModeT c.Int

const (
	LV_BLEND_MODE_NORMAL      LvBlendModeT = 0
	LV_BLEND_MODE_ADDITIVE    LvBlendModeT = 1
	LV_BLEND_MODE_SUBTRACTIVE LvBlendModeT = 2
	LV_BLEND_MODE_MULTIPLY    LvBlendModeT = 3
	LV_BLEND_MODE_DIFFERENCE  LvBlendModeT = 4
)

type LvTextDecorT c.Int

const (
	LV_TEXT_DECOR_NONE          LvTextDecorT = 0
	LV_TEXT_DECOR_UNDERLINE     LvTextDecorT = 1
	LV_TEXT_DECOR_STRIKETHROUGH LvTextDecorT = 2
)

type LvBorderSideT c.Int

const (
	LV_BORDER_SIDE_NONE     LvBorderSideT = 0
	LV_BORDER_SIDE_BOTTOM   LvBorderSideT = 1
	LV_BORDER_SIDE_TOP      LvBorderSideT = 2
	LV_BORDER_SIDE_LEFT     LvBorderSideT = 4
	LV_BORDER_SIDE_RIGHT    LvBorderSideT = 8
	LV_BORDER_SIDE_FULL     LvBorderSideT = 15
	LV_BORDER_SIDE_INTERNAL LvBorderSideT = 16
)

/**
 * A common type to handle all the property types in the same way.
 */

type LvStyleValueT struct {
	Ptr c.Pointer
}
type X_lvStyleIdT c.Int

const (
	LV_STYLE_PROP_INV                  X_lvStyleIdT = 0
	LV_STYLE_WIDTH                     X_lvStyleIdT = 1
	LV_STYLE_HEIGHT                    X_lvStyleIdT = 2
	LV_STYLE_LENGTH                    X_lvStyleIdT = 3
	LV_STYLE_MIN_WIDTH                 X_lvStyleIdT = 4
	LV_STYLE_MAX_WIDTH                 X_lvStyleIdT = 5
	LV_STYLE_MIN_HEIGHT                X_lvStyleIdT = 6
	LV_STYLE_MAX_HEIGHT                X_lvStyleIdT = 7
	LV_STYLE_X                         X_lvStyleIdT = 8
	LV_STYLE_Y                         X_lvStyleIdT = 9
	LV_STYLE_ALIGN                     X_lvStyleIdT = 10
	LV_STYLE_RADIUS                    X_lvStyleIdT = 12
	LV_STYLE_RADIAL_OFFSET             X_lvStyleIdT = 13
	LV_STYLE_PAD_RADIAL                X_lvStyleIdT = 14
	LV_STYLE_PAD_TOP                   X_lvStyleIdT = 16
	LV_STYLE_PAD_BOTTOM                X_lvStyleIdT = 17
	LV_STYLE_PAD_LEFT                  X_lvStyleIdT = 18
	LV_STYLE_PAD_RIGHT                 X_lvStyleIdT = 19
	LV_STYLE_PAD_ROW                   X_lvStyleIdT = 20
	LV_STYLE_PAD_COLUMN                X_lvStyleIdT = 21
	LV_STYLE_LAYOUT                    X_lvStyleIdT = 22
	LV_STYLE_MARGIN_TOP                X_lvStyleIdT = 24
	LV_STYLE_MARGIN_BOTTOM             X_lvStyleIdT = 25
	LV_STYLE_MARGIN_LEFT               X_lvStyleIdT = 26
	LV_STYLE_MARGIN_RIGHT              X_lvStyleIdT = 27
	LV_STYLE_BG_COLOR                  X_lvStyleIdT = 28
	LV_STYLE_BG_OPA                    X_lvStyleIdT = 29
	LV_STYLE_BG_GRAD_DIR               X_lvStyleIdT = 32
	LV_STYLE_BG_MAIN_STOP              X_lvStyleIdT = 33
	LV_STYLE_BG_GRAD_STOP              X_lvStyleIdT = 34
	LV_STYLE_BG_GRAD_COLOR             X_lvStyleIdT = 35
	LV_STYLE_BG_MAIN_OPA               X_lvStyleIdT = 36
	LV_STYLE_BG_GRAD_OPA               X_lvStyleIdT = 37
	LV_STYLE_BG_GRAD                   X_lvStyleIdT = 38
	LV_STYLE_BASE_DIR                  X_lvStyleIdT = 39
	LV_STYLE_BG_IMAGE_SRC              X_lvStyleIdT = 40
	LV_STYLE_BG_IMAGE_OPA              X_lvStyleIdT = 41
	LV_STYLE_BG_IMAGE_RECOLOR          X_lvStyleIdT = 42
	LV_STYLE_BG_IMAGE_RECOLOR_OPA      X_lvStyleIdT = 43
	LV_STYLE_BG_IMAGE_TILED            X_lvStyleIdT = 44
	LV_STYLE_CLIP_CORNER               X_lvStyleIdT = 45
	LV_STYLE_BORDER_WIDTH              X_lvStyleIdT = 48
	LV_STYLE_BORDER_COLOR              X_lvStyleIdT = 49
	LV_STYLE_BORDER_OPA                X_lvStyleIdT = 50
	LV_STYLE_BORDER_SIDE               X_lvStyleIdT = 52
	LV_STYLE_BORDER_POST               X_lvStyleIdT = 53
	LV_STYLE_OUTLINE_WIDTH             X_lvStyleIdT = 56
	LV_STYLE_OUTLINE_COLOR             X_lvStyleIdT = 57
	LV_STYLE_OUTLINE_OPA               X_lvStyleIdT = 58
	LV_STYLE_OUTLINE_PAD               X_lvStyleIdT = 59
	LV_STYLE_SHADOW_WIDTH              X_lvStyleIdT = 60
	LV_STYLE_SHADOW_COLOR              X_lvStyleIdT = 61
	LV_STYLE_SHADOW_OPA                X_lvStyleIdT = 62
	LV_STYLE_SHADOW_OFFSET_X           X_lvStyleIdT = 64
	LV_STYLE_SHADOW_OFFSET_Y           X_lvStyleIdT = 65
	LV_STYLE_SHADOW_SPREAD             X_lvStyleIdT = 66
	LV_STYLE_IMAGE_OPA                 X_lvStyleIdT = 68
	LV_STYLE_IMAGE_RECOLOR             X_lvStyleIdT = 69
	LV_STYLE_IMAGE_RECOLOR_OPA         X_lvStyleIdT = 70
	LV_STYLE_LINE_WIDTH                X_lvStyleIdT = 72
	LV_STYLE_LINE_DASH_WIDTH           X_lvStyleIdT = 73
	LV_STYLE_LINE_DASH_GAP             X_lvStyleIdT = 74
	LV_STYLE_LINE_ROUNDED              X_lvStyleIdT = 75
	LV_STYLE_LINE_COLOR                X_lvStyleIdT = 76
	LV_STYLE_LINE_OPA                  X_lvStyleIdT = 77
	LV_STYLE_ARC_WIDTH                 X_lvStyleIdT = 80
	LV_STYLE_ARC_ROUNDED               X_lvStyleIdT = 81
	LV_STYLE_ARC_COLOR                 X_lvStyleIdT = 82
	LV_STYLE_ARC_OPA                   X_lvStyleIdT = 83
	LV_STYLE_ARC_IMAGE_SRC             X_lvStyleIdT = 84
	LV_STYLE_TEXT_COLOR                X_lvStyleIdT = 88
	LV_STYLE_TEXT_OPA                  X_lvStyleIdT = 89
	LV_STYLE_TEXT_FONT                 X_lvStyleIdT = 90
	LV_STYLE_TEXT_LETTER_SPACE         X_lvStyleIdT = 91
	LV_STYLE_TEXT_LINE_SPACE           X_lvStyleIdT = 92
	LV_STYLE_TEXT_DECOR                X_lvStyleIdT = 93
	LV_STYLE_TEXT_ALIGN                X_lvStyleIdT = 94
	LV_STYLE_TEXT_OUTLINE_STROKE_WIDTH X_lvStyleIdT = 95
	LV_STYLE_TEXT_OUTLINE_STROKE_OPA   X_lvStyleIdT = 96
	LV_STYLE_TEXT_OUTLINE_STROKE_COLOR X_lvStyleIdT = 97
	LV_STYLE_OPA                       X_lvStyleIdT = 98
	LV_STYLE_OPA_LAYERED               X_lvStyleIdT = 99
	LV_STYLE_COLOR_FILTER_DSC          X_lvStyleIdT = 100
	LV_STYLE_COLOR_FILTER_OPA          X_lvStyleIdT = 101
	LV_STYLE_ANIM                      X_lvStyleIdT = 102
	LV_STYLE_ANIM_DURATION             X_lvStyleIdT = 103
	LV_STYLE_TRANSITION                X_lvStyleIdT = 104
	LV_STYLE_BLEND_MODE                X_lvStyleIdT = 105
	LV_STYLE_TRANSFORM_WIDTH           X_lvStyleIdT = 106
	LV_STYLE_TRANSFORM_HEIGHT          X_lvStyleIdT = 107
	LV_STYLE_TRANSLATE_X               X_lvStyleIdT = 108
	LV_STYLE_TRANSLATE_Y               X_lvStyleIdT = 109
	LV_STYLE_TRANSFORM_SCALE_X         X_lvStyleIdT = 110
	LV_STYLE_TRANSFORM_SCALE_Y         X_lvStyleIdT = 111
	LV_STYLE_TRANSFORM_ROTATION        X_lvStyleIdT = 112
	LV_STYLE_TRANSFORM_PIVOT_X         X_lvStyleIdT = 113
	LV_STYLE_TRANSFORM_PIVOT_Y         X_lvStyleIdT = 114
	LV_STYLE_TRANSFORM_SKEW_X          X_lvStyleIdT = 115
	LV_STYLE_TRANSFORM_SKEW_Y          X_lvStyleIdT = 116
	LV_STYLE_BITMAP_MASK_SRC           X_lvStyleIdT = 117
	LV_STYLE_ROTARY_SENSITIVITY        X_lvStyleIdT = 118
	LV_STYLE_TRANSLATE_RADIAL          X_lvStyleIdT = 119
	LV_STYLE_RECOLOR                   X_lvStyleIdT = 120
	LV_STYLE_RECOLOR_OPA               X_lvStyleIdT = 121
	LV_STYLE_FLEX_FLOW                 X_lvStyleIdT = 122
	LV_STYLE_FLEX_MAIN_PLACE           X_lvStyleIdT = 123
	LV_STYLE_FLEX_CROSS_PLACE          X_lvStyleIdT = 124
	LV_STYLE_FLEX_TRACK_PLACE          X_lvStyleIdT = 125
	LV_STYLE_FLEX_GROW                 X_lvStyleIdT = 126
	LV_STYLE_GRID_COLUMN_ALIGN         X_lvStyleIdT = 127
	LV_STYLE_GRID_ROW_ALIGN            X_lvStyleIdT = 128
	LV_STYLE_GRID_ROW_DSC_ARRAY        X_lvStyleIdT = 129
	LV_STYLE_GRID_COLUMN_DSC_ARRAY     X_lvStyleIdT = 130
	LV_STYLE_GRID_CELL_COLUMN_POS      X_lvStyleIdT = 131
	LV_STYLE_GRID_CELL_COLUMN_SPAN     X_lvStyleIdT = 132
	LV_STYLE_GRID_CELL_X_ALIGN         X_lvStyleIdT = 133
	LV_STYLE_GRID_CELL_ROW_POS         X_lvStyleIdT = 134
	LV_STYLE_GRID_CELL_ROW_SPAN        X_lvStyleIdT = 135
	LV_STYLE_GRID_CELL_Y_ALIGN         X_lvStyleIdT = 136
	LV_STYLE_LAST_BUILT_IN_PROP        X_lvStyleIdT = 137
	LV_STYLE_NUM_BUILT_IN_PROPS        X_lvStyleIdT = 138
	LV_STYLE_PROP_ANY                  X_lvStyleIdT = 255
	LV_STYLE_PROP_CONST                X_lvStyleIdT = 255
)

type LvStyleResT c.Int

const (
	LV_STYLE_RES_NOT_FOUND LvStyleResT = 0
	LV_STYLE_RES_FOUND     LvStyleResT = 1
)

/**
 * Descriptor for style transitions
 */

type LvStyleTransitionDscT struct {
	Props    *LvStylePropT
	UserData c.Pointer
	PathXcb  LvAnimPathCbT
	Time     c.Uint32T
	Delay    c.Uint32T
}

/**
 * Descriptor of a constant style property.
 */

type LvStyleConstPropT struct {
	Prop  LvStylePropT
	Value LvStyleValueT
}

/**
 * Descriptor of a style (a collection of properties and values).
 */

type LvStyleT struct {
	ValuesAndProps c.Pointer
	HasGroup       c.Uint32T
	PropCnt        c.Uint8T
}

/**
 * Initialize a style
 * @param style pointer to a style to initialize
 * @note Do not call `lv_style_init` on styles that already have some properties
 *       because this function won't free the used memory, just sets a default state for the style.
 *       In other words be sure to initialize styles only once!
 */
// llgo:link (*LvStyleT).LvStyleInit C.lv_style_init
func (recv_ *LvStyleT) LvStyleInit() {
}

/**
 * Clear all properties from a style and free all allocated memories.
 * @param style pointer to a style
 */
// llgo:link (*LvStyleT).LvStyleReset C.lv_style_reset
func (recv_ *LvStyleT) LvStyleReset() {
}

/**
 * Copy all properties of a style to an other.
 * It has the same affect callying the same `lv_set_style_...`
 * functions on both styles.
 * It means new memory will be allocated to store the properties in
 * the destination style.
 * After the copy the destination style is fully independent of the source
 * and source can removed without affecting the destination style.
 * @param dst   the destination to copy into (can not the a constant style)
 * @param src   the source style to copy from.
 */
// llgo:link (*LvStyleT).LvStyleCopy C.lv_style_copy
func (recv_ *LvStyleT) LvStyleCopy(src *LvStyleT) {
}

/**
 * Register a new style property for custom usage
 * @return a new property ID, or LV_STYLE_PROP_INV if there are no more available.
 *
 * Example:
 * @code
 * lv_style_prop_t MY_PROP;
 * static inline void lv_style_set_my_prop(lv_style_t * style, lv_color_t value) {
 * lv_style_value_t v = {.color = value}; lv_style_set_prop(style, MY_PROP, v); }
 *
 * ...
 * MY_PROP = lv_style_register_prop();
 * ...
 * lv_style_set_my_prop(&style1, lv_palette_main(LV_PALETTE_RED));
 * @endcode
 */
//go:linkname LvStyleRegisterProp C.lv_style_register_prop
func LvStyleRegisterProp(flag c.Uint8T) LvStylePropT

/**
 * Get the number of custom properties that have been registered thus far.
 */
//go:linkname LvStyleGetNumCustomProps C.lv_style_get_num_custom_props
func LvStyleGetNumCustomProps() LvStylePropT

/**
 * Remove a property from a style
 * @param style pointer to a style
 * @param prop  a style property ORed with a state.
 * @return true: the property was found and removed; false: the property wasn't found
 */
// llgo:link (*LvStyleT).LvStyleRemoveProp C.lv_style_remove_prop
func (recv_ *LvStyleT) LvStyleRemoveProp(prop LvStylePropT) bool {
	return false
}

/**
 * Set the value of property in a style.
 * This function shouldn't be used directly by the user.
 * Instead use `lv_style_set_<prop_name>()`. E.g. `lv_style_set_bg_color()`
 * @param style pointer to style
 * @param prop the ID of a property (e.g. `LV_STYLE_BG_COLOR`)
 * @param value `lv_style_value_t` variable in which a field is set according to the type of `prop`
 */
// llgo:link (*LvStyleT).LvStyleSetProp C.lv_style_set_prop
func (recv_ *LvStyleT) LvStyleSetProp(prop LvStylePropT, value LvStyleValueT) {
}

/**
 * Get the value of a property
 * @param style pointer to a style
 * @param prop  the ID of a property
 * @param value pointer to a `lv_style_value_t` variable to store the value
 * @return LV_RESULT_INVALID: the property wasn't found in the style (`value` is unchanged)
 *         LV_RESULT_OK: the property was fond, and `value` is set accordingly
 * @note For performance reasons there are no sanity check on `style`
 */
// llgo:link (*LvStyleT).LvStyleGetProp C.lv_style_get_prop
func (recv_ *LvStyleT) LvStyleGetProp(prop LvStylePropT, value *LvStyleValueT) LvStyleResT {
	return 0
}

/**
 * Initialize a transition descriptor.
 * @param tr        pointer to a transition descriptor to initialize
 * @param props     an array with the properties to transition. The last element must be zero.
 * @param path_cb   an animation path (ease) callback. If `NULL` liner path will be used.
 * @param time      duration of the transition in [ms]
 * @param delay     delay before the transition in [ms]
 * @param user_data any custom data that will be saved in the transition animation and will be available when `path_cb` is called
 *
 * Example:
 * @code
 * const static lv_style_prop_t trans_props[] = { LV_STYLE_BG_OPA, LV_STYLE_BG_COLOR, 0 };
 * static lv_style_transition_dsc_t trans1;
 * lv_style_transition_dsc_init(&trans1, trans_props, NULL, 300, 0, NULL);
 * @endcode
 */
// llgo:link (*LvStyleTransitionDscT).LvStyleTransitionDscInit C.lv_style_transition_dsc_init
func (recv_ *LvStyleTransitionDscT) LvStyleTransitionDscInit(props *LvStylePropT, path_cb LvAnimPathCbT, time c.Uint32T, delay c.Uint32T, user_data c.Pointer) {
}

/**
 * Get the default value of a property
 * @param prop the ID of a property
 * @return the default value
 */
// llgo:link LvStylePropT.LvStylePropGetDefault C.lv_style_prop_get_default
func (recv_ LvStylePropT) LvStylePropGetDefault() LvStyleValueT {
	return LvStyleValueT{}
}

/**
 * Checks if a style is empty (has no properties)
 * @param style pointer to a style
 * @return true if the style is empty
 */
// llgo:link (*LvStyleT).LvStyleIsEmpty C.lv_style_is_empty
func (recv_ *LvStyleT) LvStyleIsEmpty() bool {
	return false
}

/**
 * Get the flags of a built-in or custom property.
 *
 * @param prop a style property
 * @return the flags of the property
 */
// llgo:link LvStylePropT.LvStylePropLookupFlags C.lv_style_prop_lookup_flags
func (recv_ LvStylePropT) LvStylePropLookupFlags() c.Uint8T {
	return 0
}

// llgo:link (*LvStyleT).LvStyleSetWidth C.lv_style_set_width
func (recv_ *LvStyleT) LvStyleSetWidth(value c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetMinWidth C.lv_style_set_min_width
func (recv_ *LvStyleT) LvStyleSetMinWidth(value c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetMaxWidth C.lv_style_set_max_width
func (recv_ *LvStyleT) LvStyleSetMaxWidth(value c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetHeight C.lv_style_set_height
func (recv_ *LvStyleT) LvStyleSetHeight(value c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetMinHeight C.lv_style_set_min_height
func (recv_ *LvStyleT) LvStyleSetMinHeight(value c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetMaxHeight C.lv_style_set_max_height
func (recv_ *LvStyleT) LvStyleSetMaxHeight(value c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetLength C.lv_style_set_length
func (recv_ *LvStyleT) LvStyleSetLength(value c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetX C.lv_style_set_x
func (recv_ *LvStyleT) LvStyleSetX(value c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetY C.lv_style_set_y
func (recv_ *LvStyleT) LvStyleSetY(value c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetAlign C.lv_style_set_align
func (recv_ *LvStyleT) LvStyleSetAlign(value LvAlignT) {
}

// llgo:link (*LvStyleT).LvStyleSetTransformWidth C.lv_style_set_transform_width
func (recv_ *LvStyleT) LvStyleSetTransformWidth(value c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetTransformHeight C.lv_style_set_transform_height
func (recv_ *LvStyleT) LvStyleSetTransformHeight(value c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetTranslateX C.lv_style_set_translate_x
func (recv_ *LvStyleT) LvStyleSetTranslateX(value c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetTranslateY C.lv_style_set_translate_y
func (recv_ *LvStyleT) LvStyleSetTranslateY(value c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetTranslateRadial C.lv_style_set_translate_radial
func (recv_ *LvStyleT) LvStyleSetTranslateRadial(value c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetTransformScaleX C.lv_style_set_transform_scale_x
func (recv_ *LvStyleT) LvStyleSetTransformScaleX(value c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetTransformScaleY C.lv_style_set_transform_scale_y
func (recv_ *LvStyleT) LvStyleSetTransformScaleY(value c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetTransformRotation C.lv_style_set_transform_rotation
func (recv_ *LvStyleT) LvStyleSetTransformRotation(value c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetTransformPivotX C.lv_style_set_transform_pivot_x
func (recv_ *LvStyleT) LvStyleSetTransformPivotX(value c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetTransformPivotY C.lv_style_set_transform_pivot_y
func (recv_ *LvStyleT) LvStyleSetTransformPivotY(value c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetTransformSkewX C.lv_style_set_transform_skew_x
func (recv_ *LvStyleT) LvStyleSetTransformSkewX(value c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetTransformSkewY C.lv_style_set_transform_skew_y
func (recv_ *LvStyleT) LvStyleSetTransformSkewY(value c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetPadTop C.lv_style_set_pad_top
func (recv_ *LvStyleT) LvStyleSetPadTop(value c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetPadBottom C.lv_style_set_pad_bottom
func (recv_ *LvStyleT) LvStyleSetPadBottom(value c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetPadLeft C.lv_style_set_pad_left
func (recv_ *LvStyleT) LvStyleSetPadLeft(value c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetPadRight C.lv_style_set_pad_right
func (recv_ *LvStyleT) LvStyleSetPadRight(value c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetPadRow C.lv_style_set_pad_row
func (recv_ *LvStyleT) LvStyleSetPadRow(value c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetPadColumn C.lv_style_set_pad_column
func (recv_ *LvStyleT) LvStyleSetPadColumn(value c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetPadRadial C.lv_style_set_pad_radial
func (recv_ *LvStyleT) LvStyleSetPadRadial(value c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetMarginTop C.lv_style_set_margin_top
func (recv_ *LvStyleT) LvStyleSetMarginTop(value c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetMarginBottom C.lv_style_set_margin_bottom
func (recv_ *LvStyleT) LvStyleSetMarginBottom(value c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetMarginLeft C.lv_style_set_margin_left
func (recv_ *LvStyleT) LvStyleSetMarginLeft(value c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetMarginRight C.lv_style_set_margin_right
func (recv_ *LvStyleT) LvStyleSetMarginRight(value c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetBgColor C.lv_style_set_bg_color
func (recv_ *LvStyleT) LvStyleSetBgColor(value LvColorT) {
}

// llgo:link (*LvStyleT).LvStyleSetBgOpa C.lv_style_set_bg_opa
func (recv_ *LvStyleT) LvStyleSetBgOpa(value LvOpaT) {
}

// llgo:link (*LvStyleT).LvStyleSetBgGradColor C.lv_style_set_bg_grad_color
func (recv_ *LvStyleT) LvStyleSetBgGradColor(value LvColorT) {
}

// llgo:link (*LvStyleT).LvStyleSetBgGradDir C.lv_style_set_bg_grad_dir
func (recv_ *LvStyleT) LvStyleSetBgGradDir(value LvGradDirT) {
}

// llgo:link (*LvStyleT).LvStyleSetBgMainStop C.lv_style_set_bg_main_stop
func (recv_ *LvStyleT) LvStyleSetBgMainStop(value c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetBgGradStop C.lv_style_set_bg_grad_stop
func (recv_ *LvStyleT) LvStyleSetBgGradStop(value c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetBgMainOpa C.lv_style_set_bg_main_opa
func (recv_ *LvStyleT) LvStyleSetBgMainOpa(value LvOpaT) {
}

// llgo:link (*LvStyleT).LvStyleSetBgGradOpa C.lv_style_set_bg_grad_opa
func (recv_ *LvStyleT) LvStyleSetBgGradOpa(value LvOpaT) {
}

// llgo:link (*LvStyleT).LvStyleSetBgGrad C.lv_style_set_bg_grad
func (recv_ *LvStyleT) LvStyleSetBgGrad(value *LvGradDscT) {
}

// llgo:link (*LvStyleT).LvStyleSetBgImageSrc C.lv_style_set_bg_image_src
func (recv_ *LvStyleT) LvStyleSetBgImageSrc(value c.Pointer) {
}

// llgo:link (*LvStyleT).LvStyleSetBgImageOpa C.lv_style_set_bg_image_opa
func (recv_ *LvStyleT) LvStyleSetBgImageOpa(value LvOpaT) {
}

// llgo:link (*LvStyleT).LvStyleSetBgImageRecolor C.lv_style_set_bg_image_recolor
func (recv_ *LvStyleT) LvStyleSetBgImageRecolor(value LvColorT) {
}

// llgo:link (*LvStyleT).LvStyleSetBgImageRecolorOpa C.lv_style_set_bg_image_recolor_opa
func (recv_ *LvStyleT) LvStyleSetBgImageRecolorOpa(value LvOpaT) {
}

// llgo:link (*LvStyleT).LvStyleSetBgImageTiled C.lv_style_set_bg_image_tiled
func (recv_ *LvStyleT) LvStyleSetBgImageTiled(value bool) {
}

// llgo:link (*LvStyleT).LvStyleSetBorderColor C.lv_style_set_border_color
func (recv_ *LvStyleT) LvStyleSetBorderColor(value LvColorT) {
}

// llgo:link (*LvStyleT).LvStyleSetBorderOpa C.lv_style_set_border_opa
func (recv_ *LvStyleT) LvStyleSetBorderOpa(value LvOpaT) {
}

// llgo:link (*LvStyleT).LvStyleSetBorderWidth C.lv_style_set_border_width
func (recv_ *LvStyleT) LvStyleSetBorderWidth(value c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetBorderSide C.lv_style_set_border_side
func (recv_ *LvStyleT) LvStyleSetBorderSide(value LvBorderSideT) {
}

// llgo:link (*LvStyleT).LvStyleSetBorderPost C.lv_style_set_border_post
func (recv_ *LvStyleT) LvStyleSetBorderPost(value bool) {
}

// llgo:link (*LvStyleT).LvStyleSetOutlineWidth C.lv_style_set_outline_width
func (recv_ *LvStyleT) LvStyleSetOutlineWidth(value c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetOutlineColor C.lv_style_set_outline_color
func (recv_ *LvStyleT) LvStyleSetOutlineColor(value LvColorT) {
}

// llgo:link (*LvStyleT).LvStyleSetOutlineOpa C.lv_style_set_outline_opa
func (recv_ *LvStyleT) LvStyleSetOutlineOpa(value LvOpaT) {
}

// llgo:link (*LvStyleT).LvStyleSetOutlinePad C.lv_style_set_outline_pad
func (recv_ *LvStyleT) LvStyleSetOutlinePad(value c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetShadowWidth C.lv_style_set_shadow_width
func (recv_ *LvStyleT) LvStyleSetShadowWidth(value c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetShadowOffsetX C.lv_style_set_shadow_offset_x
func (recv_ *LvStyleT) LvStyleSetShadowOffsetX(value c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetShadowOffsetY C.lv_style_set_shadow_offset_y
func (recv_ *LvStyleT) LvStyleSetShadowOffsetY(value c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetShadowSpread C.lv_style_set_shadow_spread
func (recv_ *LvStyleT) LvStyleSetShadowSpread(value c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetShadowColor C.lv_style_set_shadow_color
func (recv_ *LvStyleT) LvStyleSetShadowColor(value LvColorT) {
}

// llgo:link (*LvStyleT).LvStyleSetShadowOpa C.lv_style_set_shadow_opa
func (recv_ *LvStyleT) LvStyleSetShadowOpa(value LvOpaT) {
}

// llgo:link (*LvStyleT).LvStyleSetImageOpa C.lv_style_set_image_opa
func (recv_ *LvStyleT) LvStyleSetImageOpa(value LvOpaT) {
}

// llgo:link (*LvStyleT).LvStyleSetImageRecolor C.lv_style_set_image_recolor
func (recv_ *LvStyleT) LvStyleSetImageRecolor(value LvColorT) {
}

// llgo:link (*LvStyleT).LvStyleSetImageRecolorOpa C.lv_style_set_image_recolor_opa
func (recv_ *LvStyleT) LvStyleSetImageRecolorOpa(value LvOpaT) {
}

// llgo:link (*LvStyleT).LvStyleSetLineWidth C.lv_style_set_line_width
func (recv_ *LvStyleT) LvStyleSetLineWidth(value c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetLineDashWidth C.lv_style_set_line_dash_width
func (recv_ *LvStyleT) LvStyleSetLineDashWidth(value c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetLineDashGap C.lv_style_set_line_dash_gap
func (recv_ *LvStyleT) LvStyleSetLineDashGap(value c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetLineRounded C.lv_style_set_line_rounded
func (recv_ *LvStyleT) LvStyleSetLineRounded(value bool) {
}

// llgo:link (*LvStyleT).LvStyleSetLineColor C.lv_style_set_line_color
func (recv_ *LvStyleT) LvStyleSetLineColor(value LvColorT) {
}

// llgo:link (*LvStyleT).LvStyleSetLineOpa C.lv_style_set_line_opa
func (recv_ *LvStyleT) LvStyleSetLineOpa(value LvOpaT) {
}

// llgo:link (*LvStyleT).LvStyleSetArcWidth C.lv_style_set_arc_width
func (recv_ *LvStyleT) LvStyleSetArcWidth(value c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetArcRounded C.lv_style_set_arc_rounded
func (recv_ *LvStyleT) LvStyleSetArcRounded(value bool) {
}

// llgo:link (*LvStyleT).LvStyleSetArcColor C.lv_style_set_arc_color
func (recv_ *LvStyleT) LvStyleSetArcColor(value LvColorT) {
}

// llgo:link (*LvStyleT).LvStyleSetArcOpa C.lv_style_set_arc_opa
func (recv_ *LvStyleT) LvStyleSetArcOpa(value LvOpaT) {
}

// llgo:link (*LvStyleT).LvStyleSetArcImageSrc C.lv_style_set_arc_image_src
func (recv_ *LvStyleT) LvStyleSetArcImageSrc(value c.Pointer) {
}

// llgo:link (*LvStyleT).LvStyleSetTextColor C.lv_style_set_text_color
func (recv_ *LvStyleT) LvStyleSetTextColor(value LvColorT) {
}

// llgo:link (*LvStyleT).LvStyleSetTextOpa C.lv_style_set_text_opa
func (recv_ *LvStyleT) LvStyleSetTextOpa(value LvOpaT) {
}

// llgo:link (*LvStyleT).LvStyleSetTextFont C.lv_style_set_text_font
func (recv_ *LvStyleT) LvStyleSetTextFont(value *LvFontT) {
}

// llgo:link (*LvStyleT).LvStyleSetTextLetterSpace C.lv_style_set_text_letter_space
func (recv_ *LvStyleT) LvStyleSetTextLetterSpace(value c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetTextLineSpace C.lv_style_set_text_line_space
func (recv_ *LvStyleT) LvStyleSetTextLineSpace(value c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetTextDecor C.lv_style_set_text_decor
func (recv_ *LvStyleT) LvStyleSetTextDecor(value LvTextDecorT) {
}

// llgo:link (*LvStyleT).LvStyleSetTextAlign C.lv_style_set_text_align
func (recv_ *LvStyleT) LvStyleSetTextAlign(value LvTextAlignT) {
}

// llgo:link (*LvStyleT).LvStyleSetTextOutlineStrokeColor C.lv_style_set_text_outline_stroke_color
func (recv_ *LvStyleT) LvStyleSetTextOutlineStrokeColor(value LvColorT) {
}

// llgo:link (*LvStyleT).LvStyleSetTextOutlineStrokeWidth C.lv_style_set_text_outline_stroke_width
func (recv_ *LvStyleT) LvStyleSetTextOutlineStrokeWidth(value c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetTextOutlineStrokeOpa C.lv_style_set_text_outline_stroke_opa
func (recv_ *LvStyleT) LvStyleSetTextOutlineStrokeOpa(value LvOpaT) {
}

// llgo:link (*LvStyleT).LvStyleSetRadius C.lv_style_set_radius
func (recv_ *LvStyleT) LvStyleSetRadius(value c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetRadialOffset C.lv_style_set_radial_offset
func (recv_ *LvStyleT) LvStyleSetRadialOffset(value c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetClipCorner C.lv_style_set_clip_corner
func (recv_ *LvStyleT) LvStyleSetClipCorner(value bool) {
}

// llgo:link (*LvStyleT).LvStyleSetOpa C.lv_style_set_opa
func (recv_ *LvStyleT) LvStyleSetOpa(value LvOpaT) {
}

// llgo:link (*LvStyleT).LvStyleSetOpaLayered C.lv_style_set_opa_layered
func (recv_ *LvStyleT) LvStyleSetOpaLayered(value LvOpaT) {
}

// llgo:link (*LvStyleT).LvStyleSetColorFilterDsc C.lv_style_set_color_filter_dsc
func (recv_ *LvStyleT) LvStyleSetColorFilterDsc(value *LvColorFilterDscT) {
}

// llgo:link (*LvStyleT).LvStyleSetColorFilterOpa C.lv_style_set_color_filter_opa
func (recv_ *LvStyleT) LvStyleSetColorFilterOpa(value LvOpaT) {
}

// llgo:link (*LvStyleT).LvStyleSetRecolor C.lv_style_set_recolor
func (recv_ *LvStyleT) LvStyleSetRecolor(value LvColorT) {
}

// llgo:link (*LvStyleT).LvStyleSetRecolorOpa C.lv_style_set_recolor_opa
func (recv_ *LvStyleT) LvStyleSetRecolorOpa(value LvOpaT) {
}

// llgo:link (*LvStyleT).LvStyleSetAnim C.lv_style_set_anim
func (recv_ *LvStyleT) LvStyleSetAnim(value *LvAnimT) {
}

// llgo:link (*LvStyleT).LvStyleSetAnimDuration C.lv_style_set_anim_duration
func (recv_ *LvStyleT) LvStyleSetAnimDuration(value c.Uint32T) {
}

// llgo:link (*LvStyleT).LvStyleSetTransition C.lv_style_set_transition
func (recv_ *LvStyleT) LvStyleSetTransition(value *LvStyleTransitionDscT) {
}

// llgo:link (*LvStyleT).LvStyleSetBlendMode C.lv_style_set_blend_mode
func (recv_ *LvStyleT) LvStyleSetBlendMode(value LvBlendModeT) {
}

// llgo:link (*LvStyleT).LvStyleSetLayout C.lv_style_set_layout
func (recv_ *LvStyleT) LvStyleSetLayout(value c.Uint16T) {
}

// llgo:link (*LvStyleT).LvStyleSetBaseDir C.lv_style_set_base_dir
func (recv_ *LvStyleT) LvStyleSetBaseDir(value LvBaseDirT) {
}

// llgo:link (*LvStyleT).LvStyleSetBitmapMaskSrc C.lv_style_set_bitmap_mask_src
func (recv_ *LvStyleT) LvStyleSetBitmapMaskSrc(value c.Pointer) {
}

// llgo:link (*LvStyleT).LvStyleSetRotarySensitivity C.lv_style_set_rotary_sensitivity
func (recv_ *LvStyleT) LvStyleSetRotarySensitivity(value c.Uint32T) {
}

// llgo:link (*LvStyleT).LvStyleSetFlexFlow C.lv_style_set_flex_flow
func (recv_ *LvStyleT) LvStyleSetFlexFlow(value LvFlexFlowT) {
}

// llgo:link (*LvStyleT).LvStyleSetFlexMainPlace C.lv_style_set_flex_main_place
func (recv_ *LvStyleT) LvStyleSetFlexMainPlace(value LvFlexAlignT) {
}

// llgo:link (*LvStyleT).LvStyleSetFlexCrossPlace C.lv_style_set_flex_cross_place
func (recv_ *LvStyleT) LvStyleSetFlexCrossPlace(value LvFlexAlignT) {
}

// llgo:link (*LvStyleT).LvStyleSetFlexTrackPlace C.lv_style_set_flex_track_place
func (recv_ *LvStyleT) LvStyleSetFlexTrackPlace(value LvFlexAlignT) {
}

// llgo:link (*LvStyleT).LvStyleSetFlexGrow C.lv_style_set_flex_grow
func (recv_ *LvStyleT) LvStyleSetFlexGrow(value c.Uint8T) {
}

// llgo:link (*LvStyleT).LvStyleSetGridColumnDscArray C.lv_style_set_grid_column_dsc_array
func (recv_ *LvStyleT) LvStyleSetGridColumnDscArray(value *c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetGridColumnAlign C.lv_style_set_grid_column_align
func (recv_ *LvStyleT) LvStyleSetGridColumnAlign(value LvGridAlignT) {
}

// llgo:link (*LvStyleT).LvStyleSetGridRowDscArray C.lv_style_set_grid_row_dsc_array
func (recv_ *LvStyleT) LvStyleSetGridRowDscArray(value *c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetGridRowAlign C.lv_style_set_grid_row_align
func (recv_ *LvStyleT) LvStyleSetGridRowAlign(value LvGridAlignT) {
}

// llgo:link (*LvStyleT).LvStyleSetGridCellColumnPos C.lv_style_set_grid_cell_column_pos
func (recv_ *LvStyleT) LvStyleSetGridCellColumnPos(value c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetGridCellXAlign C.lv_style_set_grid_cell_x_align
func (recv_ *LvStyleT) LvStyleSetGridCellXAlign(value LvGridAlignT) {
}

// llgo:link (*LvStyleT).LvStyleSetGridCellColumnSpan C.lv_style_set_grid_cell_column_span
func (recv_ *LvStyleT) LvStyleSetGridCellColumnSpan(value c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetGridCellRowPos C.lv_style_set_grid_cell_row_pos
func (recv_ *LvStyleT) LvStyleSetGridCellRowPos(value c.Int32T) {
}

// llgo:link (*LvStyleT).LvStyleSetGridCellYAlign C.lv_style_set_grid_cell_y_align
func (recv_ *LvStyleT) LvStyleSetGridCellYAlign(value LvGridAlignT) {
}

// llgo:link (*LvStyleT).LvStyleSetGridCellRowSpan C.lv_style_set_grid_cell_row_span
func (recv_ *LvStyleT) LvStyleSetGridCellRowSpan(value c.Int32T) {
}

type LvFsResT c.Int

const (
	LV_FS_RES_OK         LvFsResT = 0
	LV_FS_RES_HW_ERR     LvFsResT = 1
	LV_FS_RES_FS_ERR     LvFsResT = 2
	LV_FS_RES_NOT_EX     LvFsResT = 3
	LV_FS_RES_FULL       LvFsResT = 4
	LV_FS_RES_LOCKED     LvFsResT = 5
	LV_FS_RES_DENIED     LvFsResT = 6
	LV_FS_RES_BUSY       LvFsResT = 7
	LV_FS_RES_TOUT       LvFsResT = 8
	LV_FS_RES_NOT_IMP    LvFsResT = 9
	LV_FS_RES_OUT_OF_MEM LvFsResT = 10
	LV_FS_RES_INV_PARAM  LvFsResT = 11
	LV_FS_RES_UNKNOWN    LvFsResT = 12
)

type LvFsModeT c.Int

const (
	LV_FS_MODE_WR LvFsModeT = 1
	LV_FS_MODE_RD LvFsModeT = 2
)

type LvFsWhenceT c.Int

const (
	LV_FS_SEEK_SET LvFsWhenceT = 0
	LV_FS_SEEK_CUR LvFsWhenceT = 1
	LV_FS_SEEK_END LvFsWhenceT = 2
)

type X_lvFsDrvT struct {
	Letter     c.Char
	CacheSize  c.Uint32T
	ReadyCb    c.Pointer
	OpenCb     c.Pointer
	CloseCb    c.Pointer
	ReadCb     c.Pointer
	WriteCb    c.Pointer
	SeekCb     c.Pointer
	TellCb     c.Pointer
	DirOpenCb  c.Pointer
	DirReadCb  c.Pointer
	DirCloseCb c.Pointer
	UserData   c.Pointer
}
type LvFsDrvT X_lvFsDrvT

type LvFsFileT struct {
	FileD c.Pointer
	Drv   *LvFsDrvT
	Cache *LvFsFileCacheT
}

type LvFsDirT struct {
	DirD c.Pointer
	Drv  *LvFsDrvT
}

/**
 * Initialize a file system driver with default values.
 * It is used to ensure all fields have known values and not memory junk.
 * After it you can set the fields.
 * @param drv     pointer to driver variable to initialize
 */
// llgo:link (*LvFsDrvT).LvFsDrvInit C.lv_fs_drv_init
func (recv_ *LvFsDrvT) LvFsDrvInit() {
}

/**
 * Add a new drive
 * @param drv       pointer to an lv_fs_drv_t structure which is inited with the
 *                  corresponding function pointers. Only pointer is saved, so the
 *                  driver should be static or dynamically allocated.
 */
// llgo:link (*LvFsDrvT).LvFsDrvRegister C.lv_fs_drv_register
func (recv_ *LvFsDrvT) LvFsDrvRegister() {
}

/**
 * Give a pointer to a driver from its letter
 * @param letter    the driver-identifier letter
 * @return          pointer to a driver or NULL if not found
 */
//go:linkname LvFsGetDrv C.lv_fs_get_drv
func LvFsGetDrv(letter c.Char) *LvFsDrvT

/**
 * Test if a drive is ready or not. If the `ready` function was not initialized `true` will be
 * returned.
 * @param letter    letter of the drive
 * @return          true: drive is ready; false: drive is not ready
 */
//go:linkname LvFsIsReady C.lv_fs_is_ready
func LvFsIsReady(letter c.Char) bool

/**
 * Open a file
 * @param file_p    pointer to a lv_fs_file_t variable
 * @param path      path to the file beginning with the driver letter (e.g. S:/folder/file.txt)
 * @param mode      read: FS_MODE_RD, write: FS_MODE_WR, both: FS_MODE_RD | FS_MODE_WR
 * @return          LV_FS_RES_OK or any error from lv_fs_res_t enum
 */
// llgo:link (*LvFsFileT).LvFsOpen C.lv_fs_open
func (recv_ *LvFsFileT) LvFsOpen(path *c.Char, mode LvFsModeT) LvFsResT {
	return 0
}

/**
 * Make a path object for the memory-mapped file compatible with the file system interface
 * @param path      path to a lv_fs_path_ex object
 * @param letter    the identifier letter of the driver. E.g. `LV_FS_MEMFS_LETTER`
 * @param buf       address of the memory buffer
 * @param size      size of the memory buffer in bytes
 */
// llgo:link (*LvFsPathExT).LvFsMakePathFromBuffer C.lv_fs_make_path_from_buffer
func (recv_ *LvFsPathExT) LvFsMakePathFromBuffer(letter c.Char, buf c.Pointer, size c.Uint32T) {
}

/**
 * Close an already opened file
 * @param file_p    pointer to a lv_fs_file_t variable
 * @return          LV_FS_RES_OK or any error from lv_fs_res_t enum
 */
// llgo:link (*LvFsFileT).LvFsClose C.lv_fs_close
func (recv_ *LvFsFileT) LvFsClose() LvFsResT {
	return 0
}

/**
 * Read from a file
 * @param file_p    pointer to a lv_fs_file_t variable
 * @param buf       pointer to a buffer where the read bytes are stored
 * @param btr       Bytes To Read
 * @param br        the number of real read bytes (Bytes Read). NULL if unused.
 * @return          LV_FS_RES_OK or any error from lv_fs_res_t enum
 */
// llgo:link (*LvFsFileT).LvFsRead C.lv_fs_read
func (recv_ *LvFsFileT) LvFsRead(buf c.Pointer, btr c.Uint32T, br *c.Uint32T) LvFsResT {
	return 0
}

/**
 * Write into a file
 * @param file_p    pointer to a lv_fs_file_t variable
 * @param buf       pointer to a buffer with the bytes to write
 * @param btw       Bytes To Write
 * @param bw        the number of real written bytes (Bytes Written). NULL if unused.
 * @return          LV_FS_RES_OK or any error from lv_fs_res_t enum
 */
// llgo:link (*LvFsFileT).LvFsWrite C.lv_fs_write
func (recv_ *LvFsFileT) LvFsWrite(buf c.Pointer, btw c.Uint32T, bw *c.Uint32T) LvFsResT {
	return 0
}

/**
 * Set the position of the 'cursor' (read write pointer) in a file
 * @param file_p    pointer to a lv_fs_file_t variable
 * @param pos       the new position expressed in bytes index (0: start of file)
 * @param whence    tells from where to set position. See lv_fs_whence_t
 * @return          LV_FS_RES_OK or any error from lv_fs_res_t enum
 */
// llgo:link (*LvFsFileT).LvFsSeek C.lv_fs_seek
func (recv_ *LvFsFileT) LvFsSeek(pos c.Uint32T, whence LvFsWhenceT) LvFsResT {
	return 0
}

/**
 * Give the position of the read write pointer
 * @param file_p    pointer to a lv_fs_file_t variable
 * @param pos       pointer to store the position of the read write pointer
 * @return          LV_FS_RES_OK or any error from 'fs_res_t'
 */
// llgo:link (*LvFsFileT).LvFsTell C.lv_fs_tell
func (recv_ *LvFsFileT) LvFsTell(pos *c.Uint32T) LvFsResT {
	return 0
}

/**
 * Initialize a 'fs_dir_t' variable for directory reading
 * @param rddir_p   pointer to a 'lv_fs_dir_t' variable
 * @param path      path to a directory
 * @return          LV_FS_RES_OK or any error from lv_fs_res_t enum
 */
// llgo:link (*LvFsDirT).LvFsDirOpen C.lv_fs_dir_open
func (recv_ *LvFsDirT) LvFsDirOpen(path *c.Char) LvFsResT {
	return 0
}

/**
 * Read the next filename form a directory.
 * The name of the directories will begin with '/'
 * @param rddir_p   pointer to an initialized 'fs_dir_t' variable
 * @param fn        pointer to a buffer to store the filename
 * @param fn_len    length of the buffer to store the filename
 * @return          LV_FS_RES_OK or any error from lv_fs_res_t enum
 */
// llgo:link (*LvFsDirT).LvFsDirRead C.lv_fs_dir_read
func (recv_ *LvFsDirT) LvFsDirRead(fn *c.Char, fn_len c.Uint32T) LvFsResT {
	return 0
}

/**
 * Close the directory reading
 * @param rddir_p   pointer to an initialized 'fs_dir_t' variable
 * @return          LV_FS_RES_OK or any error from lv_fs_res_t enum
 */
// llgo:link (*LvFsDirT).LvFsDirClose C.lv_fs_dir_close
func (recv_ *LvFsDirT) LvFsDirClose() LvFsResT {
	return 0
}

/**
 * Fill a buffer with the letters of existing drivers
 * @param buf       buffer to store the letters ('\0' added after the last letter)
 * @return          the buffer
 */
//go:linkname LvFsGetLetters C.lv_fs_get_letters
func LvFsGetLetters(buf *c.Char) *c.Char

/**
 * Return with the extension of the filename
 * @param fn        string with a filename
 * @return          pointer to the beginning extension or empty string if no extension
 */
//go:linkname LvFsGetExt C.lv_fs_get_ext
func LvFsGetExt(fn *c.Char) *c.Char

/**
 * Step up one level
 * @param path      pointer to a file name
 * @return          the truncated file name
 */
//go:linkname LvFsUp C.lv_fs_up
func LvFsUp(path *c.Char) *c.Char

/**
 * Get the last element of a path (e.g. U:/folder/file -> file)
 * @param path      pointer to a file name
 * @return          pointer to the beginning of the last element in the path
 */
//go:linkname LvFsGetLast C.lv_fs_get_last
func LvFsGetLast(path *c.Char) *c.Char

type LvImageSrcT c.Int

const (
	LV_IMAGE_SRC_VARIABLE LvImageSrcT = 0
	LV_IMAGE_SRC_FILE     LvImageSrcT = 1
	LV_IMAGE_SRC_SYMBOL   LvImageSrcT = 2
	LV_IMAGE_SRC_UNKNOWN  LvImageSrcT = 3
)

// llgo:type C
type LvImageDecoderInfoFT func(*LvImageDecoderT, *LvImageDecoderDscT, *LvImageHeaderT) LvResultT

// llgo:type C
type LvImageDecoderOpenFT func(*LvImageDecoderT, *LvImageDecoderDscT) LvResultT

// llgo:type C
type LvImageDecoderGetAreaCbT func(*LvImageDecoderT, *LvImageDecoderDscT, *LvAreaT, *LvAreaT) LvResultT

// llgo:type C
type LvImageDecoderCloseFT func(*LvImageDecoderT, *LvImageDecoderDscT)

// llgo:type C
type LvImageDecoderCustomDrawT func(*LvLayerT, *LvImageDecoderDscT, *LvAreaT, *LvDrawImageDscT, *LvAreaT)

/**
 * Get information about an image.
 * Try the created image decoder one by one. Once one is able to get info that info will be used.
 * @param src the image source. Can be
 *  1) File name: E.g. "S:folder/img1.png" (The drivers needs to registered via `lv_fs_drv_register()`)
 *  2) Variable: Pointer to an `lv_image_dsc_t` variable
 *  3) Symbol: E.g. `LV_SYMBOL_OK`
 * @param header the image info will be stored here
 * @return LV_RESULT_OK: success; LV_RESULT_INVALID: wasn't able to get info about the image
 */
//go:linkname LvImageDecoderGetInfo C.lv_image_decoder_get_info
func LvImageDecoderGetInfo(src c.Pointer, header *LvImageHeaderT) LvResultT

/**
 * Open an image.
 * Try the created image decoders one by one. Once one is able to open the image that decoder is saved in `dsc`
 * @param dsc    describes a decoding session. Simply a pointer to an `lv_image_decoder_dsc_t` variable.
 * @param src    the image source. Can be
 *                 1) File name: E.g. "S:folder/img1.png" (The drivers needs to registered via `lv_fs_drv_register())`)
 *                 2) Variable: Pointer to an `lv_image_dsc_t` variable
 *                 3) Symbol: E.g. `LV_SYMBOL_OK`
 * @param args   args about how the image should be opened.
 * @return LV_RESULT_OK: opened the image. `dsc->decoded` and `dsc->header` are set.
 *         LV_RESULT_INVALID: none of the registered image decoders were able to open the image.
 */
// llgo:link (*LvImageDecoderDscT).LvImageDecoderOpen C.lv_image_decoder_open
func (recv_ *LvImageDecoderDscT) LvImageDecoderOpen(src c.Pointer, args *LvImageDecoderArgsT) LvResultT {
	return 0
}

/**
 * Decode `full_area` pixels incrementally by calling in a loop. Set `decoded_area` to `LV_COORD_MIN` on first call.
 * @param dsc           image decoder descriptor
 * @param full_area     input parameter. the full area to decode after enough subsequent calls
 * @param decoded_area  input+output parameter. set the values to `LV_COORD_MIN` for the first call and to reset decoding.
 *                      the decoded area is stored here after each call.
 * @return              LV_RESULT_OK: success; LV_RESULT_INVALID: an error occurred or there is nothing left to decode
 */
// llgo:link (*LvImageDecoderDscT).LvImageDecoderGetArea C.lv_image_decoder_get_area
func (recv_ *LvImageDecoderDscT) LvImageDecoderGetArea(full_area *LvAreaT, decoded_area *LvAreaT) LvResultT {
	return 0
}

/**
 * Close a decoding session
 * @param dsc pointer to `lv_image_decoder_dsc_t` used in `lv_image_decoder_open`
 */
// llgo:link (*LvImageDecoderDscT).LvImageDecoderClose C.lv_image_decoder_close
func (recv_ *LvImageDecoderDscT) LvImageDecoderClose() {
}

/**
 * Create a new image decoder
 * @return pointer to the new image decoder
 */
//go:linkname LvImageDecoderCreate C.lv_image_decoder_create
func LvImageDecoderCreate() *LvImageDecoderT

/**
 * Delete an image decoder
 * @param decoder pointer to an image decoder
 */
// llgo:link (*LvImageDecoderT).LvImageDecoderDelete C.lv_image_decoder_delete
func (recv_ *LvImageDecoderT) LvImageDecoderDelete() {
}

/**
 * Get the next image decoder in the linked list of image decoders
 * @param decoder pointer to an image decoder or NULL to get the first one
 * @return the next image decoder or NULL if no more image decoder exists
 */
// llgo:link (*LvImageDecoderT).LvImageDecoderGetNext C.lv_image_decoder_get_next
func (recv_ *LvImageDecoderT) LvImageDecoderGetNext() *LvImageDecoderT {
	return nil
}

/**
 * Set a callback to get information about the image
 * @param decoder pointer to an image decoder
 * @param info_cb a function to collect info about an image (fill an `lv_image_header_t` struct)
 */
// llgo:link (*LvImageDecoderT).LvImageDecoderSetInfoCb C.lv_image_decoder_set_info_cb
func (recv_ *LvImageDecoderT) LvImageDecoderSetInfoCb(info_cb LvImageDecoderInfoFT) {
}

/**
 * Set a callback to open an image
 * @param decoder pointer to an image decoder
 * @param open_cb a function to open an image
 */
// llgo:link (*LvImageDecoderT).LvImageDecoderSetOpenCb C.lv_image_decoder_set_open_cb
func (recv_ *LvImageDecoderT) LvImageDecoderSetOpenCb(open_cb LvImageDecoderOpenFT) {
}

/**
 * Set a callback to a decoded line of an image
 * @param decoder pointer to an image decoder
 * @param read_line_cb a function to read a line of an image
 */
// llgo:link (*LvImageDecoderT).LvImageDecoderSetGetAreaCb C.lv_image_decoder_set_get_area_cb
func (recv_ *LvImageDecoderT) LvImageDecoderSetGetAreaCb(read_line_cb LvImageDecoderGetAreaCbT) {
}

/**
 * Set a callback to close a decoding session. E.g. close files and free other resources.
 * @param decoder pointer to an image decoder
 * @param close_cb a function to close a decoding session
 */
// llgo:link (*LvImageDecoderT).LvImageDecoderSetCloseCb C.lv_image_decoder_set_close_cb
func (recv_ *LvImageDecoderT) LvImageDecoderSetCloseCb(close_cb LvImageDecoderCloseFT) {
}

// llgo:link (*LvImageDecoderT).LvImageDecoderAddToCache C.lv_image_decoder_add_to_cache
func (recv_ *LvImageDecoderT) LvImageDecoderAddToCache(search_key *LvImageCacheDataT, decoded *LvDrawBufT, user_data c.Pointer) *LvCacheEntryT {
	return nil
}

/**
 * Check the decoded image, make any modification if decoder `args` requires.
 * @note A new draw buf will be allocated if provided `decoded` is not modifiable or stride mismatch etc.
 * @param dsc       pointer to a decoder descriptor
 * @param decoded   pointer to a decoded image to post process to meet dsc->args requirement.
 * @return          post processed draw buffer, when it differs with `decoded`, it's newly allocated.
 */
// llgo:link (*LvImageDecoderDscT).LvImageDecoderPostProcess C.lv_image_decoder_post_process
func (recv_ *LvImageDecoderDscT) LvImageDecoderPostProcess(decoded *LvDrawBufT) *LvDrawBufT {
	return nil
}

type LvDrawTaskTypeT c.Int

const (
	LV_DRAW_TASK_TYPE_NONE           LvDrawTaskTypeT = 0
	LV_DRAW_TASK_TYPE_FILL           LvDrawTaskTypeT = 1
	LV_DRAW_TASK_TYPE_BORDER         LvDrawTaskTypeT = 2
	LV_DRAW_TASK_TYPE_BOX_SHADOW     LvDrawTaskTypeT = 3
	LV_DRAW_TASK_TYPE_LETTER         LvDrawTaskTypeT = 4
	LV_DRAW_TASK_TYPE_LABEL          LvDrawTaskTypeT = 5
	LV_DRAW_TASK_TYPE_IMAGE          LvDrawTaskTypeT = 6
	LV_DRAW_TASK_TYPE_LAYER          LvDrawTaskTypeT = 7
	LV_DRAW_TASK_TYPE_LINE           LvDrawTaskTypeT = 8
	LV_DRAW_TASK_TYPE_ARC            LvDrawTaskTypeT = 9
	LV_DRAW_TASK_TYPE_TRIANGLE       LvDrawTaskTypeT = 10
	LV_DRAW_TASK_TYPE_MASK_RECTANGLE LvDrawTaskTypeT = 11
	LV_DRAW_TASK_TYPE_MASK_BITMAP    LvDrawTaskTypeT = 12
)

type LvDrawTaskStateT c.Int

const (
	LV_DRAW_TASK_STATE_WAITING     LvDrawTaskStateT = 0
	LV_DRAW_TASK_STATE_QUEUED      LvDrawTaskStateT = 1
	LV_DRAW_TASK_STATE_IN_PROGRESS LvDrawTaskStateT = 2
	LV_DRAW_TASK_STATE_READY       LvDrawTaskStateT = 3
)

type LvDrawDscBaseT struct {
	Obj      *LvObjT
	Part     LvPartT
	Id1      c.Uint32T
	Id2      c.Uint32T
	Layer    *LvLayerT
	DscSize  c.SizeT
	UserData c.Pointer
}

/**
 * Used internally to initialize the drawing module
 */
//go:linkname LvDrawInit C.lv_draw_init
func LvDrawInit()

/**
 * Deinitialize the drawing module
 */
//go:linkname LvDrawDeinit C.lv_draw_deinit
func LvDrawDeinit()

/**
 * Allocate a new draw unit with the given size and appends it to the list of draw units
 * @param size      the size to allocate. E.g. `sizeof(my_draw_unit_t)`,
 *                  where the first element of `my_draw_unit_t` is `lv_draw_unit_t`.
 */
//go:linkname LvDrawCreateUnit C.lv_draw_create_unit
func LvDrawCreateUnit(size c.SizeT) c.Pointer

/**
 * Add an empty draw task to the draw task list of a layer.
 * @param layer     pointer to a layer
 * @param coords    the coordinates of the draw task
 * @return          the created draw task which needs to be
 *                  further configured e.g. by added a draw descriptor
 */
// llgo:link (*LvLayerT).LvDrawAddTask C.lv_draw_add_task
func (recv_ *LvLayerT) LvDrawAddTask(coords *LvAreaT, type_ LvDrawTaskTypeT) *LvDrawTaskT {
	return nil
}

/**
 * Needs to be called when a draw task is created and configured.
 * It will send an event about the new draw task to the widget
 * and assign it to a draw unit.
 * @param layer     pointer to a layer
 * @param t         pointer to a draw task
 */
// llgo:link (*LvLayerT).LvDrawFinalizeTaskCreation C.lv_draw_finalize_task_creation
func (recv_ *LvLayerT) LvDrawFinalizeTaskCreation(t *LvDrawTaskT) {
}

/**
 * Try dispatching draw tasks to draw units
 */
//go:linkname LvDrawDispatch C.lv_draw_dispatch
func LvDrawDispatch()

/**
 * Used internally to try dispatching draw tasks of a specific layer
 * @param disp      pointer to a display on which the dispatching was requested
 * @param layer     pointer to a layer
 * @return          at least one draw task is being rendered (maybe it was taken earlier)
 */
// llgo:link (*LvDisplayT).LvDrawDispatchLayer C.lv_draw_dispatch_layer
func (recv_ *LvDisplayT) LvDrawDispatchLayer(layer *LvLayerT) bool {
	return false
}

/**
 * Wait for a new dispatch request.
 * It's blocking if `LV_USE_OS == 0` else it yields
 */
//go:linkname LvDrawDispatchWaitForRequest C.lv_draw_dispatch_wait_for_request
func LvDrawDispatchWaitForRequest()

/**
 * Wait for draw finish in case of asynchronous task execution.
 * If `LV_USE_OS == 0` it just return.
 */
//go:linkname LvDrawWaitForFinish C.lv_draw_wait_for_finish
func LvDrawWaitForFinish()

/**
 * When a draw unit finished a draw task it needs to request dispatching
 * to let LVGL assign a new draw task to it
 */
//go:linkname LvDrawDispatchRequest C.lv_draw_dispatch_request
func LvDrawDispatchRequest()

/**
 * Get the total number of draw units.
 */
//go:linkname LvDrawGetUnitCount C.lv_draw_get_unit_count
func LvDrawGetUnitCount() c.Uint32T

/**
 * If there is only one draw unit check the first draw task if it's available.
 * If there are multiple draw units call `lv_draw_get_next_available_task` to find a task.
 * @param layer             the draw layer to search in
 * @param t_prev            continue searching from this task
 * @param draw_unit_id      check the task where `preferred_draw_unit_id` equals this value or `LV_DRAW_UNIT_NONE`
 * @return                  an available draw task or NULL if there is not any
 */
// llgo:link (*LvLayerT).LvDrawGetAvailableTask C.lv_draw_get_available_task
func (recv_ *LvLayerT) LvDrawGetAvailableTask(t_prev *LvDrawTaskT, draw_unit_id c.Uint8T) *LvDrawTaskT {
	return nil
}

/**
 * Find and available draw task
 * @param layer             the draw layer to search in
 * @param t_prev            continue searching from this task
 * @param draw_unit_id      check the task where `preferred_draw_unit_id` equals this value or `LV_DRAW_UNIT_NONE`
 * @return                  an available draw task or NULL if there is not any
 */
// llgo:link (*LvLayerT).LvDrawGetNextAvailableTask C.lv_draw_get_next_available_task
func (recv_ *LvLayerT) LvDrawGetNextAvailableTask(t_prev *LvDrawTaskT, draw_unit_id c.Uint8T) *LvDrawTaskT {
	return nil
}

/**
 * Tell how many draw task are waiting to be drawn on the area of `t_check`.
 * It can be used to determine if a GPU shall combine many draw tasks into one or not.
 * If a lot of tasks are waiting for the current ones it makes sense to draw them one-by-one
 * to not block the dependent tasks' rendering
 * @param t_check   the task whose dependent tasks shall be counted
 * @return          number of tasks depending on `t_check`
 */
// llgo:link (*LvDrawTaskT).LvDrawGetDependentCount C.lv_draw_get_dependent_count
func (recv_ *LvDrawTaskT) LvDrawGetDependentCount() c.Uint32T {
	return 0
}

/**
 * Initialize a layer
 * @param layer pointer to a layer to initialize
 */
// llgo:link (*LvLayerT).LvLayerInit C.lv_layer_init
func (recv_ *LvLayerT) LvLayerInit() {
}

/**
 * Reset the layer to a drawable state
 * @param layer pointer to a layer to reset
 */
// llgo:link (*LvLayerT).LvLayerReset C.lv_layer_reset
func (recv_ *LvLayerT) LvLayerReset() {
}

/**
 * Create (allocate) a new layer on a parent layer
 * @param parent_layer      the parent layer to which the layer will be merged when it's rendered
 * @param color_format      the color format of the layer
 * @param area              the areas of the layer (absolute coordinates)
 * @return                  the new target_layer or NULL on error
 */
// llgo:link (*LvLayerT).LvDrawLayerCreate C.lv_draw_layer_create
func (recv_ *LvLayerT) LvDrawLayerCreate(color_format LvColorFormatT, area *LvAreaT) *LvLayerT {
	return nil
}

/**
 * Initialize a layer which is allocated by the user
 * @param layer             pointer the layer to initialize (its lifetime needs to be managed by the user)
 * @param parent_layer      the parent layer to which the layer will be merged when it's rendered
 * @param color_format      the color format of the layer
 * @param area              the areas of the layer (absolute coordinates)
 * @return                  the new target_layer or NULL on error
 */
// llgo:link (*LvLayerT).LvDrawLayerInit C.lv_draw_layer_init
func (recv_ *LvLayerT) LvDrawLayerInit(parent_layer *LvLayerT, color_format LvColorFormatT, area *LvAreaT) {
}

/**
 * Try to allocate a buffer for the layer.
 * @param layer             pointer to a layer
 * @return                  pointer to the allocated aligned buffer or NULL on failure
 */
// llgo:link (*LvLayerT).LvDrawLayerAllocBuf C.lv_draw_layer_alloc_buf
func (recv_ *LvLayerT) LvDrawLayerAllocBuf() c.Pointer {
	return nil
}

/**
 * Got to a pixel at X and Y coordinate on a layer
 * @param layer             pointer to a layer
 * @param x                 the target X coordinate
 * @param y                 the target X coordinate
 * @return                  `buf` offset to point to the given X and Y coordinate
 */
// llgo:link (*LvLayerT).LvDrawLayerGoToXy C.lv_draw_layer_go_to_xy
func (recv_ *LvLayerT) LvDrawLayerGoToXy(x c.Int32T, y c.Int32T) c.Pointer {
	return nil
}

/**
 * Get the type of a draw task
 * @param t   the draw task to get the type of
 * @return    the draw task type
 */
// llgo:link (*LvDrawTaskT).LvDrawTaskGetType C.lv_draw_task_get_type
func (recv_ *LvDrawTaskT) LvDrawTaskGetType() LvDrawTaskTypeT {
	return 0
}

/**
 * Get the draw descriptor of a draw task
 * @param t   the draw task to get the draw descriptor of
 * @return    a void pointer to the draw descriptor
 */
// llgo:link (*LvDrawTaskT).LvDrawTaskGetDrawDsc C.lv_draw_task_get_draw_dsc
func (recv_ *LvDrawTaskT) LvDrawTaskGetDrawDsc() c.Pointer {
	return nil
}

/**
 * Get the draw area of a draw task
 * @param t      the draw task to get the draw area of
 * @param area   the destination where the draw area will be stored
 */
// llgo:link (*LvDrawTaskT).LvDrawTaskGetArea C.lv_draw_task_get_area
func (recv_ *LvDrawTaskT) LvDrawTaskGetArea(area *LvAreaT) {
}

/**
 * Init an array.
 * @param array pointer to an `lv_array_t` variable to initialize
 * @param capacity the initial capacity of the array
 * @param element_size the size of an element in bytes
 */
// llgo:link (*LvArrayT).LvArrayInit C.lv_array_init
func (recv_ *LvArrayT) LvArrayInit(capacity c.Uint32T, element_size c.Uint32T) {
}

/**
 * Init an array from a buffer.
 * @note The buffer must be large enough to store `capacity` elements. The array will not release the buffer and reallocate it.
 *       The user must ensure that the buffer is valid during the lifetime of the array. And release the buffer when the array is no longer needed.
 * @param array pointer to an `lv_array_t` variable to initialize
 * @param buf pointer to a buffer to use as the array's data
 * @param capacity the initial capacity of the array
 * @param element_size the size of an element in bytes
 */
// llgo:link (*LvArrayT).LvArrayInitFromBuf C.lv_array_init_from_buf
func (recv_ *LvArrayT) LvArrayInitFromBuf(buf c.Pointer, capacity c.Uint32T, element_size c.Uint32T) {
}

/**
 * Resize the array to the given capacity.
 * @note if the new capacity is smaller than the current size, the array will be truncated.
 * @param array pointer to an `lv_array_t` variable
 * @param new_capacity the new capacity of the array
 */
// llgo:link (*LvArrayT).LvArrayResize C.lv_array_resize
func (recv_ *LvArrayT) LvArrayResize(new_capacity c.Uint32T) bool {
	return false
}

/**
 * Deinit the array, and free the allocated memory
 * @param array pointer to an `lv_array_t` variable to deinitialize
 */
// llgo:link (*LvArrayT).LvArrayDeinit C.lv_array_deinit
func (recv_ *LvArrayT) LvArrayDeinit() {
}

/**
 * Copy an array to another.
 * @note this will create a new array with the same capacity and size as the source array.
 * @param target pointer to an `lv_array_t` variable to copy to
 * @param source pointer to an `lv_array_t` variable to copy from
 */
// llgo:link (*LvArrayT).LvArrayCopy C.lv_array_copy
func (recv_ *LvArrayT) LvArrayCopy(source *LvArrayT) {
}

/**
 * Shrink the memory capacity of array if necessary.
 * @param array pointer to an `lv_array_t` variable
 */
// llgo:link (*LvArrayT).LvArrayShrink C.lv_array_shrink
func (recv_ *LvArrayT) LvArrayShrink() {
}

/**
 * Remove the element at the specified position in the array.
 * @param array pointer to an `lv_array_t` variable
 * @param index the index of the element to remove
 * @return LV_RESULT_OK: success, otherwise: error
 */
// llgo:link (*LvArrayT).LvArrayRemove C.lv_array_remove
func (recv_ *LvArrayT) LvArrayRemove(index c.Uint32T) LvResultT {
	return 0
}

/**
 * Remove from the array either a single element or a range of elements ([start, end)).
 * @note This effectively reduces the container size by the number of elements removed.
 * @note When start equals to end, the function has no effect.
 * @param array pointer to an `lv_array_t` variable
 * @param start the index of the first element to be removed
 * @param end the index of the first element that is not to be removed
 * @return LV_RESULT_OK: success, otherwise: error
 */
// llgo:link (*LvArrayT).LvArrayErase C.lv_array_erase
func (recv_ *LvArrayT) LvArrayErase(start c.Uint32T, end c.Uint32T) LvResultT {
	return 0
}

/**
 * Concatenate two arrays. Adds new elements to the end of the array.
 * @note The destination array is automatically expanded as necessary.
 * @param array pointer to an `lv_array_t` variable
 * @param other pointer to the array to concatenate
 * @return LV_RESULT_OK: success, otherwise: error
 */
// llgo:link (*LvArrayT).LvArrayConcat C.lv_array_concat
func (recv_ *LvArrayT) LvArrayConcat(other *LvArrayT) LvResultT {
	return 0
}

/**
 * Push back element. Adds a new element to the end of the array.
 * If the array capacity is not enough for the new element, the array will be resized automatically.
 * @note If the element is NULL, it will be added as an empty element.
 * @param array pointer to an `lv_array_t` variable
 * @param element pointer to the element to add. NULL to push an empty element.
 * @return LV_RESULT_OK: success, otherwise: error
 */
// llgo:link (*LvArrayT).LvArrayPushBack C.lv_array_push_back
func (recv_ *LvArrayT) LvArrayPushBack(element c.Pointer) LvResultT {
	return 0
}

/**
 * Assigns one content to the array, replacing its current content.
 * @param array pointer to an `lv_array_t` variable
 * @param index the index of the element to replace
 * @param value pointer to the elements to add
 * @return true: success; false: error
 */
// llgo:link (*LvArrayT).LvArrayAssign C.lv_array_assign
func (recv_ *LvArrayT) LvArrayAssign(index c.Uint32T, value c.Pointer) LvResultT {
	return 0
}

/**
 * Returns a pointer to the element at position n in the array.
 * @param array pointer to an `lv_array_t` variable
 * @param index the index of the element to return
 * @return a pointer to the requested element, NULL if `index` is out of range
 */
// llgo:link (*LvArrayT).LvArrayAt C.lv_array_at
func (recv_ *LvArrayT) LvArrayAt(index c.Uint32T) c.Pointer {
	return nil
}

// llgo:type C
type LvEventCbT func(*LvEventT)
type LvEventCodeT c.Int

const (
	LV_EVENT_ALL                  LvEventCodeT = 0
	LV_EVENT_PRESSED              LvEventCodeT = 1
	LV_EVENT_PRESSING             LvEventCodeT = 2
	LV_EVENT_PRESS_LOST           LvEventCodeT = 3
	LV_EVENT_SHORT_CLICKED        LvEventCodeT = 4
	LV_EVENT_SINGLE_CLICKED       LvEventCodeT = 5
	LV_EVENT_DOUBLE_CLICKED       LvEventCodeT = 6
	LV_EVENT_TRIPLE_CLICKED       LvEventCodeT = 7
	LV_EVENT_LONG_PRESSED         LvEventCodeT = 8
	LV_EVENT_LONG_PRESSED_REPEAT  LvEventCodeT = 9
	LV_EVENT_CLICKED              LvEventCodeT = 10
	LV_EVENT_RELEASED             LvEventCodeT = 11
	LV_EVENT_SCROLL_BEGIN         LvEventCodeT = 12
	LV_EVENT_SCROLL_THROW_BEGIN   LvEventCodeT = 13
	LV_EVENT_SCROLL_END           LvEventCodeT = 14
	LV_EVENT_SCROLL               LvEventCodeT = 15
	LV_EVENT_GESTURE              LvEventCodeT = 16
	LV_EVENT_KEY                  LvEventCodeT = 17
	LV_EVENT_ROTARY               LvEventCodeT = 18
	LV_EVENT_FOCUSED              LvEventCodeT = 19
	LV_EVENT_DEFOCUSED            LvEventCodeT = 20
	LV_EVENT_LEAVE                LvEventCodeT = 21
	LV_EVENT_HIT_TEST             LvEventCodeT = 22
	LV_EVENT_INDEV_RESET          LvEventCodeT = 23
	LV_EVENT_HOVER_OVER           LvEventCodeT = 24
	LV_EVENT_HOVER_LEAVE          LvEventCodeT = 25
	LV_EVENT_COVER_CHECK          LvEventCodeT = 26
	LV_EVENT_REFR_EXT_DRAW_SIZE   LvEventCodeT = 27
	LV_EVENT_DRAW_MAIN_BEGIN      LvEventCodeT = 28
	LV_EVENT_DRAW_MAIN            LvEventCodeT = 29
	LV_EVENT_DRAW_MAIN_END        LvEventCodeT = 30
	LV_EVENT_DRAW_POST_BEGIN      LvEventCodeT = 31
	LV_EVENT_DRAW_POST            LvEventCodeT = 32
	LV_EVENT_DRAW_POST_END        LvEventCodeT = 33
	LV_EVENT_DRAW_TASK_ADDED      LvEventCodeT = 34
	LV_EVENT_VALUE_CHANGED        LvEventCodeT = 35
	LV_EVENT_INSERT               LvEventCodeT = 36
	LV_EVENT_REFRESH              LvEventCodeT = 37
	LV_EVENT_READY                LvEventCodeT = 38
	LV_EVENT_CANCEL               LvEventCodeT = 39
	LV_EVENT_CREATE               LvEventCodeT = 40
	LV_EVENT_DELETE               LvEventCodeT = 41
	LV_EVENT_CHILD_CHANGED        LvEventCodeT = 42
	LV_EVENT_CHILD_CREATED        LvEventCodeT = 43
	LV_EVENT_CHILD_DELETED        LvEventCodeT = 44
	LV_EVENT_SCREEN_UNLOAD_START  LvEventCodeT = 45
	LV_EVENT_SCREEN_LOAD_START    LvEventCodeT = 46
	LV_EVENT_SCREEN_LOADED        LvEventCodeT = 47
	LV_EVENT_SCREEN_UNLOADED      LvEventCodeT = 48
	LV_EVENT_SIZE_CHANGED         LvEventCodeT = 49
	LV_EVENT_STYLE_CHANGED        LvEventCodeT = 50
	LV_EVENT_LAYOUT_CHANGED       LvEventCodeT = 51
	LV_EVENT_GET_SELF_SIZE        LvEventCodeT = 52
	LV_EVENT_INVALIDATE_AREA      LvEventCodeT = 53
	LV_EVENT_RESOLUTION_CHANGED   LvEventCodeT = 54
	LV_EVENT_COLOR_FORMAT_CHANGED LvEventCodeT = 55
	LV_EVENT_REFR_REQUEST         LvEventCodeT = 56
	LV_EVENT_REFR_START           LvEventCodeT = 57
	LV_EVENT_REFR_READY           LvEventCodeT = 58
	LV_EVENT_RENDER_START         LvEventCodeT = 59
	LV_EVENT_RENDER_READY         LvEventCodeT = 60
	LV_EVENT_FLUSH_START          LvEventCodeT = 61
	LV_EVENT_FLUSH_FINISH         LvEventCodeT = 62
	LV_EVENT_FLUSH_WAIT_START     LvEventCodeT = 63
	LV_EVENT_FLUSH_WAIT_FINISH    LvEventCodeT = 64
	LV_EVENT_VSYNC                LvEventCodeT = 65
	LV_EVENT_VSYNC_REQUEST        LvEventCodeT = 66
	LV_EVENT_LAST                 LvEventCodeT = 67
	LV_EVENT_PREPROCESS           LvEventCodeT = 32768
	LV_EVENT_MARKED_DELETING      LvEventCodeT = 65536
)

type LvEventListT struct {
	Array             LvArrayT
	IsTraversing      c.Uint8T
	HasMarkedDeleting c.Uint8T
}

/**
 * @brief Event callback.
 * Events are used to notify the user of some action being taken on Widget.
 * For details, see ::lv_event_t.
 */
// llgo:link (*LvEventListT).LvEventSend C.lv_event_send
func (recv_ *LvEventListT) LvEventSend(e *LvEventT, preprocess bool) LvResultT {
	return 0
}

// llgo:link (*LvEventListT).LvEventAdd C.lv_event_add
func (recv_ *LvEventListT) LvEventAdd(cb LvEventCbT, filter LvEventCodeT, user_data c.Pointer) *LvEventDscT {
	return nil
}

// llgo:link (*LvEventListT).LvEventRemoveDsc C.lv_event_remove_dsc
func (recv_ *LvEventListT) LvEventRemoveDsc(dsc *LvEventDscT) bool {
	return false
}

// llgo:link (*LvEventListT).LvEventGetCount C.lv_event_get_count
func (recv_ *LvEventListT) LvEventGetCount() c.Uint32T {
	return 0
}

// llgo:link (*LvEventListT).LvEventGetDsc C.lv_event_get_dsc
func (recv_ *LvEventListT) LvEventGetDsc(index c.Uint32T) *LvEventDscT {
	return nil
}

// llgo:link (*LvEventDscT).LvEventDscGetCb C.lv_event_dsc_get_cb
func (recv_ *LvEventDscT) LvEventDscGetCb() LvEventCbT {
	return nil
}

// llgo:link (*LvEventDscT).LvEventDscGetUserData C.lv_event_dsc_get_user_data
func (recv_ *LvEventDscT) LvEventDscGetUserData() c.Pointer {
	return nil
}

// llgo:link (*LvEventListT).LvEventRemove C.lv_event_remove
func (recv_ *LvEventListT) LvEventRemove(index c.Uint32T) bool {
	return false
}

// llgo:link (*LvEventListT).LvEventRemoveAll C.lv_event_remove_all
func (recv_ *LvEventListT) LvEventRemoveAll() {
}

/**
 * Get Widget originally targeted by the event. It's the same even if event was bubbled.
 * @param e     pointer to the event descriptor
 * @return      the target of the event_code
 */
// llgo:link (*LvEventT).LvEventGetTarget C.lv_event_get_target
func (recv_ *LvEventT) LvEventGetTarget() c.Pointer {
	return nil
}

/**
 * Get current target of the event. It's the Widget for which the event handler being called.
 * If the event is not bubbled it's the same as "normal" target.
 * @param e     pointer to the event descriptor
 * @return      pointer to the current target of the event_code
 */
// llgo:link (*LvEventT).LvEventGetCurrentTarget C.lv_event_get_current_target
func (recv_ *LvEventT) LvEventGetCurrentTarget() c.Pointer {
	return nil
}

/**
 * Get event code of an event.
 * @param e     pointer to the event descriptor
 * @return      the event code. (E.g. `LV_EVENT_CLICKED`, `LV_EVENT_FOCUSED`, etc)
 */
// llgo:link (*LvEventT).LvEventGetCode C.lv_event_get_code
func (recv_ *LvEventT) LvEventGetCode() LvEventCodeT {
	return 0
}

/**
 * Get parameter passed when event was sent.
 * @param e     pointer to the event descriptor
 * @return      pointer to the parameter
 */
// llgo:link (*LvEventT).LvEventGetParam C.lv_event_get_param
func (recv_ *LvEventT) LvEventGetParam() c.Pointer {
	return nil
}

/**
 * Get user_data passed when event was registered on Widget.
 * @param e     pointer to the event descriptor
 * @return      pointer to the user_data
 */
// llgo:link (*LvEventT).LvEventGetUserData C.lv_event_get_user_data
func (recv_ *LvEventT) LvEventGetUserData() c.Pointer {
	return nil
}

/**
 * Stop event from bubbling.
 * This is only valid when called in the middle of an event processing chain.
 * @param e     pointer to the event descriptor
 */
// llgo:link (*LvEventT).LvEventStopBubbling C.lv_event_stop_bubbling
func (recv_ *LvEventT) LvEventStopBubbling() {
}

/**
 * Stop processing this event.
 * This is only valid when called in the middle of an event processing chain.
 * @param e     pointer to the event descriptor
 */
// llgo:link (*LvEventT).LvEventStopProcessing C.lv_event_stop_processing
func (recv_ *LvEventT) LvEventStopProcessing() {
}

/**
 * Register a new, custom event ID.
 * It can be used the same way as e.g. `LV_EVENT_CLICKED` to send custom events
 * @return      the new event id
 *
 * Example:
 * @code
 * uint32_t LV_EVENT_MINE = 0;
 * ...
 * e = lv_event_register_id();
 * ...
 * lv_obj_send_event(obj, LV_EVENT_MINE, &some_data);
 * @endcode
 */
//go:linkname LvEventRegisterId C.lv_event_register_id
func LvEventRegisterId() c.Uint32T

/**
 * Get the name of an event code.
 * @param code  the event code
 * @return      the name of the event code as a string
 */
// llgo:link LvEventCodeT.LvEventCodeGetName C.lv_event_code_get_name
func (recv_ LvEventCodeT) LvEventCodeGetName() *c.Char {
	return nil
}

type LvDisplayRotationT c.Int

const (
	LV_DISPLAY_ROTATION_0   LvDisplayRotationT = 0
	LV_DISPLAY_ROTATION_90  LvDisplayRotationT = 1
	LV_DISPLAY_ROTATION_180 LvDisplayRotationT = 2
	LV_DISPLAY_ROTATION_270 LvDisplayRotationT = 3
)

type LvDisplayRenderModeT c.Int

const (
	LV_DISPLAY_RENDER_MODE_PARTIAL LvDisplayRenderModeT = 0
	LV_DISPLAY_RENDER_MODE_DIRECT  LvDisplayRenderModeT = 1
	LV_DISPLAY_RENDER_MODE_FULL    LvDisplayRenderModeT = 2
)

type LvScreenLoadAnimT c.Int

const (
	LV_SCREEN_LOAD_ANIM_NONE        LvScreenLoadAnimT = 0
	LV_SCREEN_LOAD_ANIM_OVER_LEFT   LvScreenLoadAnimT = 1
	LV_SCREEN_LOAD_ANIM_OVER_RIGHT  LvScreenLoadAnimT = 2
	LV_SCREEN_LOAD_ANIM_OVER_TOP    LvScreenLoadAnimT = 3
	LV_SCREEN_LOAD_ANIM_OVER_BOTTOM LvScreenLoadAnimT = 4
	LV_SCREEN_LOAD_ANIM_MOVE_LEFT   LvScreenLoadAnimT = 5
	LV_SCREEN_LOAD_ANIM_MOVE_RIGHT  LvScreenLoadAnimT = 6
	LV_SCREEN_LOAD_ANIM_MOVE_TOP    LvScreenLoadAnimT = 7
	LV_SCREEN_LOAD_ANIM_MOVE_BOTTOM LvScreenLoadAnimT = 8
	LV_SCREEN_LOAD_ANIM_FADE_IN     LvScreenLoadAnimT = 9
	LV_SCREEN_LOAD_ANIM_FADE_ON     LvScreenLoadAnimT = 9
	LV_SCREEN_LOAD_ANIM_FADE_OUT    LvScreenLoadAnimT = 10
	LV_SCREEN_LOAD_ANIM_OUT_LEFT    LvScreenLoadAnimT = 11
	LV_SCREEN_LOAD_ANIM_OUT_RIGHT   LvScreenLoadAnimT = 12
	LV_SCREEN_LOAD_ANIM_OUT_TOP     LvScreenLoadAnimT = 13
	LV_SCREEN_LOAD_ANIM_OUT_BOTTOM  LvScreenLoadAnimT = 14
)

// llgo:type C
type LvDisplayFlushCbT func(*LvDisplayT, *LvAreaT, *c.Uint8T)

// llgo:type C
type LvDisplayFlushWaitCbT func(*LvDisplayT)

/**
 * Create a new display with the given resolution
 * @param hor_res   horizontal resolution in pixels
 * @param ver_res   vertical resolution in pixels
 * @return          pointer to a display object or `NULL` on error
 */
//go:linkname LvDisplayCreate C.lv_display_create
func LvDisplayCreate(hor_res c.Int32T, ver_res c.Int32T) *LvDisplayT

/**
 * Remove a display
 * @param disp      pointer to display
 */
// llgo:link (*LvDisplayT).LvDisplayDelete C.lv_display_delete
func (recv_ *LvDisplayT) LvDisplayDelete() {
}

/**
 * Set a default display. The new screens will be created on it by default.
 * @param disp      pointer to a display
 */
// llgo:link (*LvDisplayT).LvDisplaySetDefault C.lv_display_set_default
func (recv_ *LvDisplayT) LvDisplaySetDefault() {
}

/**
 * Get the default display
 * @return          pointer to the default display
 */
//go:linkname LvDisplayGetDefault C.lv_display_get_default
func LvDisplayGetDefault() *LvDisplayT

/**
 * Get the next display.
 * @param disp      pointer to the current display. NULL to initialize.
 * @return          the next display or NULL if no more. Gives the first display when the parameter is NULL.
 */
// llgo:link (*LvDisplayT).LvDisplayGetNext C.lv_display_get_next
func (recv_ *LvDisplayT) LvDisplayGetNext() *LvDisplayT {
	return nil
}

/**
 * Sets the resolution of a display. `LV_EVENT_RESOLUTION_CHANGED` event will be sent.
 * Here the native resolution of the device should be set. If the display will be rotated later with
 * `lv_display_set_rotation` LVGL will swap the hor. and ver. resolution automatically.
 * @param disp      pointer to a display
 * @param hor_res   the new horizontal resolution
 * @param ver_res   the new vertical resolution
 */
// llgo:link (*LvDisplayT).LvDisplaySetResolution C.lv_display_set_resolution
func (recv_ *LvDisplayT) LvDisplaySetResolution(hor_res c.Int32T, ver_res c.Int32T) {
}

/**
 * It's not mandatory to use the whole display for LVGL, however in some cases physical resolution is important.
 * For example the touchpad still sees whole resolution and the values needs to be converted
 * to the active LVGL display area.
 * @param disp      pointer to a display
 * @param hor_res   the new physical horizontal resolution, or -1 to assume it's the same as the normal hor. res.
 * @param ver_res   the new physical vertical resolution, or -1 to assume it's the same as the normal hor. res.
 */
// llgo:link (*LvDisplayT).LvDisplaySetPhysicalResolution C.lv_display_set_physical_resolution
func (recv_ *LvDisplayT) LvDisplaySetPhysicalResolution(hor_res c.Int32T, ver_res c.Int32T) {
}

/**
 * If physical resolution is not the same as the normal resolution
 * the offset of the active display area can be set here.
 * @param disp      pointer to a display
 * @param x         X offset
 * @param y         Y offset
 */
// llgo:link (*LvDisplayT).LvDisplaySetOffset C.lv_display_set_offset
func (recv_ *LvDisplayT) LvDisplaySetOffset(x c.Int32T, y c.Int32T) {
}

/**
 * Set the rotation of this display. LVGL will swap the horizontal and vertical resolutions internally.
 * @param disp      pointer to a display (NULL to use the default display)
 * @param rotation  `LV_DISPLAY_ROTATION_0/90/180/270`
 */
// llgo:link (*LvDisplayT).LvDisplaySetRotation C.lv_display_set_rotation
func (recv_ *LvDisplayT) LvDisplaySetRotation(rotation LvDisplayRotationT) {
}

/**
 * Use matrix rotation for the display. This function is depended on `LV_DRAW_TRANSFORM_USE_MATRIX`
 * @param disp      pointer to a display (NULL to use the default display)
 * @param enable    true: enable matrix rotation, false: disable
 */
// llgo:link (*LvDisplayT).LvDisplaySetMatrixRotation C.lv_display_set_matrix_rotation
func (recv_ *LvDisplayT) LvDisplaySetMatrixRotation(enable bool) {
}

/**
 * Set the DPI (dot per inch) of the display.
 * dpi = sqrt(hor_res^2 + ver_res^2) / diagonal"
 * @param disp      pointer to a display
 * @param dpi       the new DPI
 */
// llgo:link (*LvDisplayT).LvDisplaySetDpi C.lv_display_set_dpi
func (recv_ *LvDisplayT) LvDisplaySetDpi(dpi c.Int32T) {
}

/**
 * Get the horizontal resolution of a display.
 * @param disp      pointer to a display (NULL to use the default display)
 * @return          the horizontal resolution of the display.
 */
// llgo:link (*LvDisplayT).LvDisplayGetHorizontalResolution C.lv_display_get_horizontal_resolution
func (recv_ *LvDisplayT) LvDisplayGetHorizontalResolution() c.Int32T {
	return 0
}

/**
 * Get the vertical resolution of a display
 * @param disp      pointer to a display (NULL to use the default display)
 * @return          the vertical resolution of the display
 */
// llgo:link (*LvDisplayT).LvDisplayGetVerticalResolution C.lv_display_get_vertical_resolution
func (recv_ *LvDisplayT) LvDisplayGetVerticalResolution() c.Int32T {
	return 0
}

/**
 * Get the original horizontal resolution of a display without considering rotation
 * @param disp      pointer to a display (NULL to use the default display)
 * @return          the horizontal resolution of the display.
 */
// llgo:link (*LvDisplayT).LvDisplayGetOriginalHorizontalResolution C.lv_display_get_original_horizontal_resolution
func (recv_ *LvDisplayT) LvDisplayGetOriginalHorizontalResolution() c.Int32T {
	return 0
}

/**
 * Get the original vertical resolution of a display without considering rotation
 * @param disp      pointer to a display (NULL to use the default display)
 * @return          the vertical resolution of the display
 */
// llgo:link (*LvDisplayT).LvDisplayGetOriginalVerticalResolution C.lv_display_get_original_vertical_resolution
func (recv_ *LvDisplayT) LvDisplayGetOriginalVerticalResolution() c.Int32T {
	return 0
}

/**
 * Get the physical horizontal resolution of a display
 * @param disp      pointer to a display (NULL to use the default display)
 * @return the      physical horizontal resolution of the display
 */
// llgo:link (*LvDisplayT).LvDisplayGetPhysicalHorizontalResolution C.lv_display_get_physical_horizontal_resolution
func (recv_ *LvDisplayT) LvDisplayGetPhysicalHorizontalResolution() c.Int32T {
	return 0
}

/**
 * Get the physical vertical resolution of a display
 * @param disp      pointer to a display (NULL to use the default display)
 * @return          the physical vertical resolution of the display
 */
// llgo:link (*LvDisplayT).LvDisplayGetPhysicalVerticalResolution C.lv_display_get_physical_vertical_resolution
func (recv_ *LvDisplayT) LvDisplayGetPhysicalVerticalResolution() c.Int32T {
	return 0
}

/**
 * Get the horizontal offset from the full / physical display
 * @param disp      pointer to a display (NULL to use the default display)
 * @return          the horizontal offset from the physical display
 */
// llgo:link (*LvDisplayT).LvDisplayGetOffsetX C.lv_display_get_offset_x
func (recv_ *LvDisplayT) LvDisplayGetOffsetX() c.Int32T {
	return 0
}

/**
 * Get the vertical offset from the full / physical display
 * @param disp      pointer to a display (NULL to use the default display)
 * @return          the horizontal offset from the physical display
 */
// llgo:link (*LvDisplayT).LvDisplayGetOffsetY C.lv_display_get_offset_y
func (recv_ *LvDisplayT) LvDisplayGetOffsetY() c.Int32T {
	return 0
}

/**
 * Get the current rotation of this display.
 * @param disp      pointer to a display (NULL to use the default display)
 * @return          the current rotation
 */
// llgo:link (*LvDisplayT).LvDisplayGetRotation C.lv_display_get_rotation
func (recv_ *LvDisplayT) LvDisplayGetRotation() LvDisplayRotationT {
	return 0
}

/**
 * Get if matrix rotation is enabled for a display or not
 * @param disp      pointer to a display (NULL to use the default display)
 * @return          true: matrix rotation is enabled; false: disabled
 */
// llgo:link (*LvDisplayT).LvDisplayGetMatrixRotation C.lv_display_get_matrix_rotation
func (recv_ *LvDisplayT) LvDisplayGetMatrixRotation() bool {
	return false
}

/**
 * Get the DPI of the display
 * @param disp      pointer to a display (NULL to use the default display)
 * @return          dpi of the display
 */
// llgo:link (*LvDisplayT).LvDisplayGetDpi C.lv_display_get_dpi
func (recv_ *LvDisplayT) LvDisplayGetDpi() c.Int32T {
	return 0
}

/**
 * Set the buffers for a display, similarly to `lv_display_set_draw_buffers`, but accept the raw buffer pointers.
 * For DIRECT/FULL rending modes, the buffer size must be at least
 * `hor_res * ver_res * lv_color_format_get_size(lv_display_get_color_format(disp))`
 * @param disp              pointer to a display
 * @param buf1              first buffer
 * @param buf2              second buffer (can be `NULL`)
 * @param buf_size          buffer size in byte
 * @param render_mode       LV_DISPLAY_RENDER_MODE_PARTIAL/DIRECT/FULL
 */
// llgo:link (*LvDisplayT).LvDisplaySetBuffers C.lv_display_set_buffers
func (recv_ *LvDisplayT) LvDisplaySetBuffers(buf1 c.Pointer, buf2 c.Pointer, buf_size c.Uint32T, render_mode LvDisplayRenderModeT) {
}

/**
 * Set the frame buffers for a display, similarly to `lv_display_set_buffers`, but allow
 * for a custom stride as required by a display controller.
 * This allows the frame buffers to have a stride alignment different from the rest of
 * the buffers`
 * @param disp              pointer to a display
 * @param buf1              first buffer
 * @param buf2              second buffer (can be `NULL`)
 * @param buf_size          buffer size in byte
 * @param stride            buffer stride in bytes
 * @param render_mode       LV_DISPLAY_RENDER_MODE_PARTIAL/DIRECT/FULL
 */
// llgo:link (*LvDisplayT).LvDisplaySetBuffersWithStride C.lv_display_set_buffers_with_stride
func (recv_ *LvDisplayT) LvDisplaySetBuffersWithStride(buf1 c.Pointer, buf2 c.Pointer, buf_size c.Uint32T, stride c.Uint32T, render_mode LvDisplayRenderModeT) {
}

/**
 * Set the buffers for a display, accept a draw buffer pointer.
 * Normally use `lv_display_set_buffers` is enough for most cases.
 * Use this function when an existing lv_draw_buf_t is available.
 * @param disp              pointer to a display
 * @param buf1              first buffer
 * @param buf2              second buffer (can be `NULL`)
 */
// llgo:link (*LvDisplayT).LvDisplaySetDrawBuffers C.lv_display_set_draw_buffers
func (recv_ *LvDisplayT) LvDisplaySetDrawBuffers(buf1 *LvDrawBufT, buf2 *LvDrawBufT) {
}

/**
 * Set the third draw buffer for a display.
 * @param disp              pointer to a display
 * @param buf3              third buffer
 */
// llgo:link (*LvDisplayT).LvDisplaySet3rdDrawBuffer C.lv_display_set_3rd_draw_buffer
func (recv_ *LvDisplayT) LvDisplaySet3rdDrawBuffer(buf3 *LvDrawBufT) {
}

/**
 * Set display render mode
 * @param disp              pointer to a display
 * @param render_mode       LV_DISPLAY_RENDER_MODE_PARTIAL/DIRECT/FULL
 */
// llgo:link (*LvDisplayT).LvDisplaySetRenderMode C.lv_display_set_render_mode
func (recv_ *LvDisplayT) LvDisplaySetRenderMode(render_mode LvDisplayRenderModeT) {
}

/**
 * Set the flush callback which will be called to copy the rendered image to the display.
 * @param disp      pointer to a display
 * @param flush_cb  the flush callback (`px_map` contains the rendered image as raw pixel map and it should be copied to `area` on the display)
 */
// llgo:link (*LvDisplayT).LvDisplaySetFlushCb C.lv_display_set_flush_cb
func (recv_ *LvDisplayT) LvDisplaySetFlushCb(flush_cb LvDisplayFlushCbT) {
}

/**
 * Set a callback to be used while LVGL is waiting flushing to be finished.
 * It can do any complex logic to wait, including semaphores, mutexes, polling flags, etc.
 * If not set the `disp->flushing` flag is used which can be cleared with `lv_display_flush_ready()`
 * @param disp      pointer to a display
 * @param wait_cb   a callback to call while LVGL is waiting for flush ready.
 *                  If NULL `lv_display_flush_ready()` can be used to signal that flushing is ready.
 */
// llgo:link (*LvDisplayT).LvDisplaySetFlushWaitCb C.lv_display_set_flush_wait_cb
func (recv_ *LvDisplayT) LvDisplaySetFlushWaitCb(wait_cb LvDisplayFlushWaitCbT) {
}

/**
 * Set the color format of the display.
 * @param disp              pointer to a display
 * @param color_format      Possible values are
 *                          - LV_COLOR_FORMAT_RGB565
 *                          - LV_COLOR_FORMAT_RGB888
 *                          - LV_COLOR_FORMAT_XRGB888
 *                          - LV_COLOR_FORMAT_ARGB888
 *@note To change the endianness of the rendered image in case of RGB565 format
 *      (i.e. swap the 2 bytes) call `lv_draw_sw_rgb565_swap` in the flush_cb
 */
// llgo:link (*LvDisplayT).LvDisplaySetColorFormat C.lv_display_set_color_format
func (recv_ *LvDisplayT) LvDisplaySetColorFormat(color_format LvColorFormatT) {
}

/**
 * Get the color format of the display
 * @param disp              pointer to a display
 * @return                  the color format
 */
// llgo:link (*LvDisplayT).LvDisplayGetColorFormat C.lv_display_get_color_format
func (recv_ *LvDisplayT) LvDisplayGetColorFormat() LvColorFormatT {
	return 0
}

/**
 * Set the number of tiles for parallel rendering.
 * @param disp              pointer to a display
 * @param tile_cnt          number of tiles (1 =< tile_cnt < 256)
 */
// llgo:link (*LvDisplayT).LvDisplaySetTileCnt C.lv_display_set_tile_cnt
func (recv_ *LvDisplayT) LvDisplaySetTileCnt(tile_cnt c.Uint32T) {
}

/**
 * Get the number of tiles used for parallel rendering
 * @param disp              pointer to a display
 * @return                  number of tiles
 */
// llgo:link (*LvDisplayT).LvDisplayGetTileCnt C.lv_display_get_tile_cnt
func (recv_ *LvDisplayT) LvDisplayGetTileCnt() c.Uint32T {
	return 0
}

/**
 * Enable anti-aliasing for the render engine
 * @param disp      pointer to a display
 * @param en        true/false
 */
// llgo:link (*LvDisplayT).LvDisplaySetAntialiasing C.lv_display_set_antialiasing
func (recv_ *LvDisplayT) LvDisplaySetAntialiasing(en bool) {
}

/**
 * Get if anti-aliasing is enabled for a display or not
 * @param disp      pointer to a display (NULL to use the default display)
 * @return          true/false
 */
// llgo:link (*LvDisplayT).LvDisplayGetAntialiasing C.lv_display_get_antialiasing
func (recv_ *LvDisplayT) LvDisplayGetAntialiasing() bool {
	return false
}

/**
 * Call from the display driver when the flushing is finished
 * @param disp      pointer to display whose `flush_cb` was called
 */
// llgo:link (*LvDisplayT).LvDisplayFlushReady C.lv_display_flush_ready
func (recv_ *LvDisplayT) LvDisplayFlushReady() {
}

/**
 * Tell if it's the last area of the refreshing process.
 * Can be called from `flush_cb` to execute some special display refreshing if needed when all areas area flushed.
 * @param disp      pointer to display
 * @return          true: it's the last area to flush;
 *                  false: there are other areas too which will be refreshed soon
 */
// llgo:link (*LvDisplayT).LvDisplayFlushIsLast C.lv_display_flush_is_last
func (recv_ *LvDisplayT) LvDisplayFlushIsLast() bool {
	return false
}

// llgo:link (*LvDisplayT).LvDisplayIsDoubleBuffered C.lv_display_is_double_buffered
func (recv_ *LvDisplayT) LvDisplayIsDoubleBuffered() bool {
	return false
}

/**
 * Return a pointer to the active screen on a display
 * @param disp      pointer to display which active screen should be get.
 *                  (NULL to use the default screen)
 * @return          pointer to the active screen object (loaded by 'lv_screen_load()')
 */
// llgo:link (*LvDisplayT).LvDisplayGetScreenActive C.lv_display_get_screen_active
func (recv_ *LvDisplayT) LvDisplayGetScreenActive() *LvObjT {
	return nil
}

/**
 * Return with a pointer to the previous screen. Only used during screen transitions.
 * @param disp      pointer to display which previous screen should be get.
 *                  (NULL to use the default screen)
 * @return          pointer to the previous screen object or NULL if not used now
 */
// llgo:link (*LvDisplayT).LvDisplayGetScreenPrev C.lv_display_get_screen_prev
func (recv_ *LvDisplayT) LvDisplayGetScreenPrev() *LvObjT {
	return nil
}

/**
 * Return the top layer. The top layer is the same on all screens and it is above the normal screen layer.
 * @param disp      pointer to display which top layer should be get. (NULL to use the default screen)
 * @return          pointer to the top layer object
 */
// llgo:link (*LvDisplayT).LvDisplayGetLayerTop C.lv_display_get_layer_top
func (recv_ *LvDisplayT) LvDisplayGetLayerTop() *LvObjT {
	return nil
}

/**
 * Return the sys. layer. The system layer is the same on all screen and it is above the normal screen and the top layer.
 * @param disp      pointer to display which sys. layer should be retrieved. (NULL to use the default screen)
 * @return          pointer to the sys layer object
 */
// llgo:link (*LvDisplayT).LvDisplayGetLayerSys C.lv_display_get_layer_sys
func (recv_ *LvDisplayT) LvDisplayGetLayerSys() *LvObjT {
	return nil
}

/**
 * Return the bottom layer. The bottom layer is the same on all screen and it is under the normal screen layer.
 * It's visible only if the screen is transparent.
 * @param disp      pointer to display (NULL to use the default screen)
 * @return          pointer to the bottom layer object
 */
// llgo:link (*LvDisplayT).LvDisplayGetLayerBottom C.lv_display_get_layer_bottom
func (recv_ *LvDisplayT) LvDisplayGetLayerBottom() *LvObjT {
	return nil
}

/**
 * Load a screen on the default display
 * @param scr       pointer to a screen
 */
// llgo:link (*X_lvObjT).LvScreenLoad C.lv_screen_load
func (recv_ *X_lvObjT) LvScreenLoad() {
}

/**
 * Switch screen with animation
 * @param scr       pointer to the new screen to load
 * @param anim_type type of the animation from `lv_screen_load_anim_t`, e.g. `LV_SCREEN_LOAD_ANIM_MOVE_LEFT`
 * @param time      time of the animation
 * @param delay     delay before the transition
 * @param auto_del  true: automatically delete the old screen
 */
// llgo:link (*LvObjT).LvScreenLoadAnim C.lv_screen_load_anim
func (recv_ *LvObjT) LvScreenLoadAnim(anim_type LvScreenLoadAnimT, time c.Uint32T, delay c.Uint32T, auto_del bool) {
}

/**
 * Get the active screen of the default display
 * @return          pointer to the active screen
 */
//go:linkname LvScreenActive C.lv_screen_active
func LvScreenActive() *LvObjT

/**
 * Get the top layer  of the default display
 * @return          pointer to the top layer
 */
//go:linkname LvLayerTop C.lv_layer_top
func LvLayerTop() *LvObjT

/**
 * Get the system layer  of the default display
 * @return          pointer to the sys layer
 */
//go:linkname LvLayerSys C.lv_layer_sys
func LvLayerSys() *LvObjT

/**
 * Get the bottom layer  of the default display
 * @return          pointer to the bottom layer
 */
//go:linkname LvLayerBottom C.lv_layer_bottom
func LvLayerBottom() *LvObjT

/**
 * Add an event handler to the display
 * @param disp          pointer to a display
 * @param event_cb      an event callback
 * @param filter        event code to react or `LV_EVENT_ALL`
 * @param user_data     optional user_data
 */
// llgo:link (*LvDisplayT).LvDisplayAddEventCb C.lv_display_add_event_cb
func (recv_ *LvDisplayT) LvDisplayAddEventCb(event_cb LvEventCbT, filter LvEventCodeT, user_data c.Pointer) {
}

/**
 * Get the number of event attached to a display
 * @param disp          pointer to a display
 * @return              number of events
 */
// llgo:link (*LvDisplayT).LvDisplayGetEventCount C.lv_display_get_event_count
func (recv_ *LvDisplayT) LvDisplayGetEventCount() c.Uint32T {
	return 0
}

/**
 * Get an event descriptor for an event
 * @param disp          pointer to a display
 * @param index         the index of the event
 * @return              the event descriptor
 */
// llgo:link (*LvDisplayT).LvDisplayGetEventDsc C.lv_display_get_event_dsc
func (recv_ *LvDisplayT) LvDisplayGetEventDsc(index c.Uint32T) *LvEventDscT {
	return nil
}

/**
 * Remove an event
 * @param disp          pointer to a display
 * @param index         the index of the event to remove
 * @return              true: and event was removed; false: no event was removed
 */
// llgo:link (*LvDisplayT).LvDisplayDeleteEvent C.lv_display_delete_event
func (recv_ *LvDisplayT) LvDisplayDeleteEvent(index c.Uint32T) bool {
	return false
}

/**
 * Remove an event_cb with user_data
 * @param disp          pointer to a display
 * @param event_cb      the event_cb of the event to remove
 * @param user_data     user_data
 * @return              the count of the event removed
 */
// llgo:link (*LvDisplayT).LvDisplayRemoveEventCbWithUserData C.lv_display_remove_event_cb_with_user_data
func (recv_ *LvDisplayT) LvDisplayRemoveEventCbWithUserData(event_cb LvEventCbT, user_data c.Pointer) c.Uint32T {
	return 0
}

/**
 * Send an event to a display
 * @param disp          pointer to a display
 * @param code          an event code. LV_EVENT_...
 * @param param         optional param
 * @return              LV_RESULT_OK: disp wasn't deleted in the event.
 */
// llgo:link (*LvDisplayT).LvDisplaySendEvent C.lv_display_send_event
func (recv_ *LvDisplayT) LvDisplaySendEvent(code LvEventCodeT, param c.Pointer) LvResultT {
	return 0
}

/**
 * Get the area to be invalidated. Can be used in `LV_EVENT_INVALIDATE_AREA`
 * @param e     pointer to an event
 * @return      the area to invalidated (can be modified as required)
 */
// llgo:link (*LvEventT).LvEventGetInvalidatedArea C.lv_event_get_invalidated_area
func (recv_ *LvEventT) LvEventGetInvalidatedArea() *LvAreaT {
	return nil
}

/**
 * Set the theme of a display. If there are no user created widgets yet the screens' theme will be updated
 * @param disp      pointer to a display
 * @param th        pointer to a theme
 */
// llgo:link (*LvDisplayT).LvDisplaySetTheme C.lv_display_set_theme
func (recv_ *LvDisplayT) LvDisplaySetTheme(th *LvThemeT) {
}

/**
 * Get the theme of a display
 * @param disp      pointer to a display
 * @return          the display's theme (can be NULL)
 */
// llgo:link (*LvDisplayT).LvDisplayGetTheme C.lv_display_get_theme
func (recv_ *LvDisplayT) LvDisplayGetTheme() *LvThemeT {
	return nil
}

/**
 * Get elapsed time since last user activity on a display (e.g. click)
 * @param disp      pointer to a display (NULL to get the overall smallest inactivity)
 * @return          elapsed ticks (milliseconds) since the last activity
 */
// llgo:link (*LvDisplayT).LvDisplayGetInactiveTime C.lv_display_get_inactive_time
func (recv_ *LvDisplayT) LvDisplayGetInactiveTime() c.Uint32T {
	return 0
}

/**
 * Manually trigger an activity on a display
 * @param disp      pointer to a display (NULL to use the default display)
 */
// llgo:link (*LvDisplayT).LvDisplayTriggerActivity C.lv_display_trigger_activity
func (recv_ *LvDisplayT) LvDisplayTriggerActivity() {
}

/**
 * Temporarily enable and disable the invalidation of the display.
 * @param disp      pointer to a display (NULL to use the default display)
 * @param en        true: enable invalidation; false: invalidation
 */
// llgo:link (*LvDisplayT).LvDisplayEnableInvalidation C.lv_display_enable_invalidation
func (recv_ *LvDisplayT) LvDisplayEnableInvalidation(en bool) {
}

/**
 * Get display invalidation is enabled.
 * @param disp      pointer to a display (NULL to use the default display)
 * @return return   true if invalidation is enabled
 */
// llgo:link (*LvDisplayT).LvDisplayIsInvalidationEnabled C.lv_display_is_invalidation_enabled
func (recv_ *LvDisplayT) LvDisplayIsInvalidationEnabled() bool {
	return false
}

/**
 * Get a pointer to the screen refresher timer to
 * modify its parameters with `lv_timer_...` functions.
 * @param disp      pointer to a display
 * @return          pointer to the display refresher timer. (NULL on error)
 */
// llgo:link (*LvDisplayT).LvDisplayGetRefrTimer C.lv_display_get_refr_timer
func (recv_ *LvDisplayT) LvDisplayGetRefrTimer() *LvTimerT {
	return nil
}

/**
 * Delete screen refresher timer
 * @param disp      pointer to a display
 */
// llgo:link (*LvDisplayT).LvDisplayDeleteRefrTimer C.lv_display_delete_refr_timer
func (recv_ *LvDisplayT) LvDisplayDeleteRefrTimer() {
}

/**
 * Register vsync event of a display. `LV_EVENT_VSYNC` event will be sent periodically.
 * Please don't use it in display event listeners, as it may cause memory leaks and illegal access issues.
 *
 * @param disp      pointer to a display
 * @param event_cb      an event callback
 * @param user_data     optional user_data
 */
// llgo:link (*LvDisplayT).LvDisplayRegisterVsyncEvent C.lv_display_register_vsync_event
func (recv_ *LvDisplayT) LvDisplayRegisterVsyncEvent(event_cb LvEventCbT, user_data c.Pointer) bool {
	return false
}

/**
 * Unregister vsync event of a display. `LV_EVENT_VSYNC` event won't be sent periodically.
 * Please don't use it in display event listeners, as it may cause memory leaks and illegal access issues.
 * @param disp      pointer to a display
 * @param event_cb      an event callback
 * @param user_data     optional user_data
 */
// llgo:link (*LvDisplayT).LvDisplayUnregisterVsyncEvent C.lv_display_unregister_vsync_event
func (recv_ *LvDisplayT) LvDisplayUnregisterVsyncEvent(event_cb LvEventCbT, user_data c.Pointer) bool {
	return false
}

/**
 * Send an vsync event to a display
 * @param disp          pointer to a display
 * @param param         optional param
 * @return              LV_RESULT_OK: disp wasn't deleted in the event.
 */
// llgo:link (*LvDisplayT).LvDisplaySendVsyncEvent C.lv_display_send_vsync_event
func (recv_ *LvDisplayT) LvDisplaySendVsyncEvent(param c.Pointer) LvResultT {
	return 0
}

// llgo:link (*LvDisplayT).LvDisplaySetUserData C.lv_display_set_user_data
func (recv_ *LvDisplayT) LvDisplaySetUserData(user_data c.Pointer) {
}

// llgo:link (*LvDisplayT).LvDisplaySetDriverData C.lv_display_set_driver_data
func (recv_ *LvDisplayT) LvDisplaySetDriverData(driver_data c.Pointer) {
}

// llgo:link (*LvDisplayT).LvDisplayGetUserData C.lv_display_get_user_data
func (recv_ *LvDisplayT) LvDisplayGetUserData() c.Pointer {
	return nil
}

// llgo:link (*LvDisplayT).LvDisplayGetDriverData C.lv_display_get_driver_data
func (recv_ *LvDisplayT) LvDisplayGetDriverData() c.Pointer {
	return nil
}

// llgo:link (*LvDisplayT).LvDisplayGetBufActive C.lv_display_get_buf_active
func (recv_ *LvDisplayT) LvDisplayGetBufActive() *LvDrawBufT {
	return nil
}

/**
 * Rotate an area in-place according to the display's rotation
 * @param disp      pointer to a display
 * @param area      pointer to an area to rotate
 */
// llgo:link (*LvDisplayT).LvDisplayRotateArea C.lv_display_rotate_area
func (recv_ *LvDisplayT) LvDisplayRotateArea(area *LvAreaT) {
}

/**
 * Get the size of the draw buffers
 * @param disp      pointer to a display
 * @return          the size of the draw buffer in bytes for valid display, 0 otherwise
 */
// llgo:link (*LvDisplayT).LvDisplayGetDrawBufSize C.lv_display_get_draw_buf_size
func (recv_ *LvDisplayT) LvDisplayGetDrawBufSize() c.Uint32T {
	return 0
}

/**
 * Get the size of the invalidated draw buffer. Can be used in the flush callback
 * to get the number of bytes used in the current render buffer.
 * @param disp      pointer to a display
 * @param width     the width of the invalidated area
 * @param height    the height of the invalidated area
 * @return          the size of the invalidated draw buffer in bytes, not accounting for
 *                  any preceding palette information for a valid display, 0 otherwise
 */
// llgo:link (*LvDisplayT).LvDisplayGetInvalidatedDrawBufSize C.lv_display_get_invalidated_draw_buf_size
func (recv_ *LvDisplayT) LvDisplayGetInvalidatedDrawBufSize(width c.Uint32T, height c.Uint32T) c.Uint32T {
	return 0
}

/**
 * For default display, computes the number of pixels (a distance or size) as if the
 * display had 160 DPI.  This allows you to specify 1/160-th fractions of an inch to
 * get real distance on the display that will be consistent regardless of its current
 * DPI.  It ensures `lv_dpx(100)`, for example, will have the same physical size
 * regardless to the DPI of the display.
 * @param n     number of 1/160-th-inch units to compute with
 * @return      number of pixels to use to make that distance
 */
//go:linkname LvDpx C.lv_dpx
func LvDpx(n c.Int32T) c.Int32T

/**
 * For specified display, computes the number of pixels (a distance or size) as if the
 * display had 160 DPI.  This allows you to specify 1/160-th fractions of an inch to
 * get real distance on the display that will be consistent regardless of its current
 * DPI.  It ensures `lv_dpx(100)`, for example, will have the same physical size
 * regardless to the DPI of the display.
 * @param disp  pointer to display whose dpi should be considered
 * @param n     number of 1/160-th-inch units to compute with
 * @return      number of pixels to use to make that distance
 */
// llgo:link (*LvDisplayT).LvDisplayDpx C.lv_display_dpx
func (recv_ *LvDisplayT) LvDisplayDpx(n c.Int32T) c.Int32T {
	return 0
}

type LvMutexT c.Int
type LvThreadT c.Int
type LvThreadSyncT c.Int
type LvThreadPrioT c.Int

const (
	LV_THREAD_PRIO_LOWEST  LvThreadPrioT = 0
	LV_THREAD_PRIO_LOW     LvThreadPrioT = 1
	LV_THREAD_PRIO_MID     LvThreadPrioT = 2
	LV_THREAD_PRIO_HIGH    LvThreadPrioT = 3
	LV_THREAD_PRIO_HIGHEST LvThreadPrioT = 4
)

/**
 * Set it for `LV_SYSMON_GET_IDLE` to show the CPU usage
 * @return the idle percentage since the last call
 */
//go:linkname LvOsGetIdlePercent C.lv_os_get_idle_percent
func LvOsGetIdlePercent() c.Uint32T

// llgo:type C
type LvDrawImageCoreCb func(*LvDrawTaskT, *LvDrawImageDscT, *LvImageDecoderDscT, *LvDrawImageSupT, *LvAreaT, *LvAreaT)

/**
 * Initialize an image draw descriptor.
 * @param dsc       pointer to a draw descriptor
 */
// llgo:link (*LvDrawImageDscT).LvDrawImageDscInit C.lv_draw_image_dsc_init
func (recv_ *LvDrawImageDscT) LvDrawImageDscInit() {
}

/**
 * Try to get an image draw descriptor from a draw task.
 * @param task      draw task
 * @return          the task's draw descriptor or NULL if the task is not of type LV_DRAW_TASK_TYPE_IMAGE
 */
// llgo:link (*LvDrawTaskT).LvDrawTaskGetImageDsc C.lv_draw_task_get_image_dsc
func (recv_ *LvDrawTaskT) LvDrawTaskGetImageDsc() *LvDrawImageDscT {
	return nil
}

/**
 * Create an image draw task
 * @param layer         pointer to a layer
 * @param dsc           pointer to an initialized draw descriptor
 * @param coords        the coordinates of the image
 * @note                `coords` can be small than the real image area
 *                      (if only a part of the image is rendered)
 *                      or can be larger (in case of tiled images).   .
 */
// llgo:link (*LvLayerT).LvDrawImage C.lv_draw_image
func (recv_ *LvLayerT) LvDrawImage(dsc *LvDrawImageDscT, coords *LvAreaT) {
}

/**
 * Create a draw task to blend a layer to another layer
 * @param layer         pointer to a layer
 * @param dsc           pointer to an initialized draw descriptor. `src` must be set to the layer to blend
 * @param coords        the coordinates of the layer.
 * @note                `coords` can be small than the total widget area from which the layer is created
 *                      (if only a part of the widget was rendered to a layer)
 */
// llgo:link (*LvLayerT).LvDrawLayer C.lv_draw_layer
func (recv_ *LvLayerT) LvDrawLayer(dsc *LvDrawImageDscT, coords *LvAreaT) {
}

/**
 * Get the type of an image source
 * @param src pointer to an image source:
 *  - pointer to an 'lv_image_t' variable (image stored internally and compiled into the code)
 *  - a path to a file (e.g. "S:/folder/image.bin")
 *  - or a symbol (e.g. LV_SYMBOL_CLOSE)
 * @return type of the image source LV_IMAGE_SRC_VARIABLE/FILE/SYMBOL/UNKNOWN
 */
//go:linkname LvImageSrcGetType C.lv_image_src_get_type
func LvImageSrcGetType(src c.Pointer) LvImageSrcT

/**********************
 *      TYPEDEFS
 **********************/

type LvDrawRectDscT struct {
	Base              LvDrawDscBaseT
	Radius            c.Int32T
	BgImageSrc        c.Pointer
	BgImageSymbolFont c.Pointer
	BgImageRecolor    LvColorT
	BgImageOpa        LvOpaT
	BgImageRecolorOpa LvOpaT
	BgImageTiled      c.Uint8T
	BgOpa             LvOpaT
	BorderOpa         LvOpaT
	OutlineOpa        LvOpaT
	ShadowOpa         LvOpaT
	BgColor           LvColorT
	BgGrad            LvGradDscT
	BorderColor       LvColorT
	BorderWidth       c.Int32T
	BorderSide        LvBorderSideT
	BorderPost        c.Uint8T
	OutlineColor      LvColorT
	OutlineWidth      c.Int32T
	OutlinePad        c.Int32T
	ShadowColor       LvColorT
	ShadowWidth       c.Int32T
	ShadowOffsetX     c.Int32T
	ShadowOffsetY     c.Int32T
	ShadowSpread      c.Int32T
}

type LvDrawFillDscT struct {
	Base   LvDrawDscBaseT
	Radius c.Int32T
	Opa    LvOpaT
	Color  LvColorT
	Grad   LvGradDscT
}

type LvDrawBorderDscT struct {
	Base   LvDrawDscBaseT
	Radius c.Int32T
	Color  LvColorT
	Width  c.Int32T
	Opa    LvOpaT
	Side   LvBorderSideT
}

type LvDrawBoxShadowDscT struct {
	Base    LvDrawDscBaseT
	Radius  c.Int32T
	Color   LvColorT
	Width   c.Int32T
	Spread  c.Int32T
	OfsX    c.Int32T
	OfsY    c.Int32T
	Opa     LvOpaT
	BgCover c.Uint8T
}

// llgo:link (*LvDrawRectDscT).LvDrawRectDscInit C.lv_draw_rect_dsc_init
func (recv_ *LvDrawRectDscT) LvDrawRectDscInit() {
}

/**
 * Initialize a fill draw descriptor.
 * @param dsc       pointer to a draw descriptor
 */
// llgo:link (*LvDrawFillDscT).LvDrawFillDscInit C.lv_draw_fill_dsc_init
func (recv_ *LvDrawFillDscT) LvDrawFillDscInit() {
}

/**
 * Try to get a fill draw descriptor from a draw task.
 * @param task      draw task
 * @return          the task's draw descriptor or NULL if the task is not of type LV_DRAW_TASK_TYPE_FILL
 */
// llgo:link (*LvDrawTaskT).LvDrawTaskGetFillDsc C.lv_draw_task_get_fill_dsc
func (recv_ *LvDrawTaskT) LvDrawTaskGetFillDsc() *LvDrawFillDscT {
	return nil
}

/**
 * Fill an area
 * @param layer         pointer to a layer
 * @param dsc           pointer to an initialized draw descriptor variable
 * @param coords        the coordinates of the rectangle
 */
// llgo:link (*LvLayerT).LvDrawFill C.lv_draw_fill
func (recv_ *LvLayerT) LvDrawFill(dsc *LvDrawFillDscT, coords *LvAreaT) {
}

/**
 * Initialize a border draw descriptor.
 * @param dsc       pointer to a draw descriptor
 */
// llgo:link (*LvDrawBorderDscT).LvDrawBorderDscInit C.lv_draw_border_dsc_init
func (recv_ *LvDrawBorderDscT) LvDrawBorderDscInit() {
}

/**
 * Try to get a border draw descriptor from a draw task.
 * @param task      draw task
 * @return          the task's draw descriptor or NULL if the task is not of type LV_DRAW_TASK_TYPE_BORDER
 */
// llgo:link (*LvDrawTaskT).LvDrawTaskGetBorderDsc C.lv_draw_task_get_border_dsc
func (recv_ *LvDrawTaskT) LvDrawTaskGetBorderDsc() *LvDrawBorderDscT {
	return nil
}

/**
 * Draw a border
 * @param layer         pointer to a layer
 * @param dsc           pointer to an initialized draw descriptor variable
 * @param coords        the coordinates of the rectangle
 */
// llgo:link (*LvLayerT).LvDrawBorder C.lv_draw_border
func (recv_ *LvLayerT) LvDrawBorder(dsc *LvDrawBorderDscT, coords *LvAreaT) {
}

/**
 * Initialize a box shadow draw descriptor.
 * @param dsc       pointer to a draw descriptor
 */
// llgo:link (*LvDrawBoxShadowDscT).LvDrawBoxShadowDscInit C.lv_draw_box_shadow_dsc_init
func (recv_ *LvDrawBoxShadowDscT) LvDrawBoxShadowDscInit() {
}

/**
 * Try to get a box shadow draw descriptor from a draw task.
 * @param task      draw task
 * @return          the task's draw descriptor or NULL if the task is not of type LV_DRAW_TASK_TYPE_BOX_SHADOW
 */
// llgo:link (*LvDrawTaskT).LvDrawTaskGetBoxShadowDsc C.lv_draw_task_get_box_shadow_dsc
func (recv_ *LvDrawTaskT) LvDrawTaskGetBoxShadowDsc() *LvDrawBoxShadowDscT {
	return nil
}

/**
 * Draw a box shadow
 * @param layer         pointer to a layer
 * @param dsc           pointer to an initialized draw descriptor variable
 * @param coords        the coordinates of the rectangle
 */
// llgo:link (*LvLayerT).LvDrawBoxShadow C.lv_draw_box_shadow
func (recv_ *LvLayerT) LvDrawBoxShadow(dsc *LvDrawBoxShadowDscT, coords *LvAreaT) {
}

/**
 * The rectangle is a wrapper for fill, border, bg. image and box shadow.
 * Internally fill, border, image and box shadow draw tasks will be created.
 * @param layer         pointer to a layer
 * @param dsc           pointer to an initialized draw descriptor variable
 * @param coords        the coordinates of the rectangle
 */
// llgo:link (*LvLayerT).LvDrawRect C.lv_draw_rect
func (recv_ *LvLayerT) LvDrawRect(dsc *LvDrawRectDscT, coords *LvAreaT) {
}

/**********************
 *      TYPEDEFS
 **********************/

type LvDrawTriangleDscT struct {
	Base  LvDrawDscBaseT
	P     [3]LvPointPreciseT
	Color LvColorT
	Opa   LvOpaT
	Grad  LvGradDscT
}

/**
 * Initialize a triangle draw descriptor
 * @param draw_dsc  pointer to a draw descriptor
 */
// llgo:link (*LvDrawTriangleDscT).LvDrawTriangleDscInit C.lv_draw_triangle_dsc_init
func (recv_ *LvDrawTriangleDscT) LvDrawTriangleDscInit() {
}

/**
 * Try to get a triangle draw descriptor from a draw task.
 * @param task      draw task
 * @return          the task's draw descriptor or NULL if the task is not of type LV_DRAW_TASK_TYPE_TRIANGLE
 */
// llgo:link (*LvDrawTaskT).LvDrawTaskGetTriangleDsc C.lv_draw_task_get_triangle_dsc
func (recv_ *LvDrawTaskT) LvDrawTaskGetTriangleDsc() *LvDrawTriangleDscT {
	return nil
}

/**
 * Create a triangle draw task
 * @param layer     pointer to a layer
 * @param draw_dsc  pointer to an initialized `lv_draw_triangle_dsc_t` object
 */
// llgo:link (*LvLayerT).LvDrawTriangle C.lv_draw_triangle
func (recv_ *LvLayerT) LvDrawTriangle(draw_dsc *LvDrawTriangleDscT) {
}

/**********************
 *      TYPEDEFS
 **********************/

type LvDrawLabelDscT struct {
	Base               LvDrawDscBaseT
	Text               *c.Char
	TextSize           LvPointT
	Font               *LvFontT
	Color              LvColorT
	LineSpace          c.Int32T
	LetterSpace        c.Int32T
	OfsX               c.Int32T
	OfsY               c.Int32T
	Rotation           c.Int32T
	SelStart           c.Uint32T
	SelEnd             c.Uint32T
	SelColor           LvColorT
	SelBgColor         LvColorT
	TextLength         c.Uint32T
	Align              LvTextAlignT
	BidiDir            LvBaseDirT
	Opa                LvOpaT
	OutlineStrokeOpa   LvOpaT
	Decor              LvTextDecorT
	Flag               LvTextFlagT
	TextLocal          c.Uint8T
	TextStatic         c.Uint8T
	HasBided           c.Uint8T
	Hint               *LvDrawLabelHintT
	OutlineStrokeColor LvColorT
	OutlineStrokeWidth c.Int32T
}

type LvDrawLetterDscT struct {
	Base               LvDrawDscBaseT
	Unicode            c.Uint32T
	Font               *LvFontT
	Color              LvColorT
	Rotation           c.Int32T
	ScaleX             c.Int32T
	ScaleY             c.Int32T
	SkewX              c.Int32T
	SkewY              c.Int32T
	Pivot              LvPointT
	Opa                LvOpaT
	Decor              LvTextDecorT
	BlendMode          LvBlendModeT
	OutlineStrokeOpa   LvOpaT
	OutlineStrokeWidth c.Int32T
	OutlineStrokeColor LvColorT
}

// llgo:type C
type LvDrawGlyphCbT func(*LvDrawTaskT, *LvDrawGlyphDscT, *LvDrawFillDscT, *LvAreaT)

// llgo:link (*LvDrawLetterDscT).LvDrawLetterDscInit C.lv_draw_letter_dsc_init
func (recv_ *LvDrawLetterDscT) LvDrawLetterDscInit() {
}

// llgo:link (*LvDrawLabelDscT).LvDrawLabelDscInit C.lv_draw_label_dsc_init
func (recv_ *LvDrawLabelDscT) LvDrawLabelDscInit() {
}

/**
 * Try to get a label draw descriptor from a draw task.
 * @param task      draw task
 * @return          the task's draw descriptor or NULL if the task is not of type LV_DRAW_TASK_TYPE_LABEL
 */
// llgo:link (*LvDrawTaskT).LvDrawTaskGetLabelDsc C.lv_draw_task_get_label_dsc
func (recv_ *LvDrawTaskT) LvDrawTaskGetLabelDsc() *LvDrawLabelDscT {
	return nil
}

/**
 * Initialize a glyph draw descriptor.
 * Used internally.
 * @param dsc       pointer to a draw descriptor
 */
// llgo:link (*LvDrawGlyphDscT).LvDrawGlyphDscInit C.lv_draw_glyph_dsc_init
func (recv_ *LvDrawGlyphDscT) LvDrawGlyphDscInit() {
}

// llgo:link (*LvLayerT).LvDrawLabel C.lv_draw_label
func (recv_ *LvLayerT) LvDrawLabel(dsc *LvDrawLabelDscT, coords *LvAreaT) {
}

// llgo:link (*LvLayerT).LvDrawCharacter C.lv_draw_character
func (recv_ *LvLayerT) LvDrawCharacter(dsc *LvDrawLabelDscT, point *LvPointT, unicode_letter c.Uint32T) {
}

// llgo:link (*LvLayerT).LvDrawLetter C.lv_draw_letter
func (recv_ *LvLayerT) LvDrawLetter(dsc *LvDrawLetterDscT, point *LvPointT) {
}

/**
 * Should be used during rendering the characters to get the position and other
 * parameters of the characters
 * @param t             pointer to a draw task
 * @param dsc           pointer to draw descriptor
 * @param coords        coordinates of the label
 * @param cb            a callback to call to draw each glyphs one by one
 */
// llgo:link (*LvDrawTaskT).LvDrawLabelIterateCharacters C.lv_draw_label_iterate_characters
func (recv_ *LvDrawTaskT) LvDrawLabelIterateCharacters(dsc *LvDrawLabelDscT, coords *LvAreaT, cb LvDrawGlyphCbT) {
}

/**
 * @brief Draw a single letter using the provided draw unit, glyph descriptor, position, font, and callback.
 *
 * This function is responsible for rendering a single character from a text string,
 * applying the necessary styling described by the glyph descriptor (`dsc`). It handles
 * the retrieval of the glyph's description, checks its visibility within the clipping area,
 * and invokes the callback (`cb`) to render the glyph at the specified position (`pos`)
 * using the given font (`font`).
 *
 * @param t             Pointer to the drawing task.
 * @param dsc           Pointer to the descriptor containing styling for the glyph to be drawn.
 * @param pos           Pointer to the point coordinates where the letter should be drawn.
 * @param font          Pointer to the font containing the glyph.
 * @param letter        The Unicode code point of the letter to be drawn.
 * @param cb            Callback function to execute the actual rendering of the glyph.
 */
// llgo:link (*LvDrawTaskT).LvDrawUnitDrawLetter C.lv_draw_unit_draw_letter
func (recv_ *LvDrawTaskT) LvDrawUnitDrawLetter(dsc *LvDrawGlyphDscT, pos *LvPointT, font *LvFontT, letter c.Uint32T, cb LvDrawGlyphCbT) {
}

/**********************
 *      TYPEDEFS
 **********************/

type LvDrawLineDscT struct {
	Base       LvDrawDscBaseT
	P1         LvPointPreciseT
	P2         LvPointPreciseT
	Color      LvColorT
	Width      c.Int32T
	DashWidth  c.Int32T
	DashGap    c.Int32T
	Opa        LvOpaT
	RoundStart c.Uint8T
	RoundEnd   c.Uint8T
	RawEnd     c.Uint8T
}

/**
 * Initialize a line draw descriptor
 * @param dsc       pointer to a draw descriptor
 */
// llgo:link (*LvDrawLineDscT).LvDrawLineDscInit C.lv_draw_line_dsc_init
func (recv_ *LvDrawLineDscT) LvDrawLineDscInit() {
}

/**
 * Try to get a line draw descriptor from a draw task.
 * @param task      draw task
 * @return          the task's draw descriptor or NULL if the task is not of type LV_DRAW_TASK_TYPE_LINE
 */
// llgo:link (*LvDrawTaskT).LvDrawTaskGetLineDsc C.lv_draw_task_get_line_dsc
func (recv_ *LvDrawTaskT) LvDrawTaskGetLineDsc() *LvDrawLineDscT {
	return nil
}

/**
 * Create a line draw task
 * @param layer     pointer to a layer
 * @param dsc       pointer to an initialized `lv_draw_line_dsc_t` variable
 */
// llgo:link (*LvLayerT).LvDrawLine C.lv_draw_line
func (recv_ *LvLayerT) LvDrawLine(dsc *LvDrawLineDscT) {
}

/**********************
 *      TYPEDEFS
 **********************/

type LvDrawArcDscT struct {
	Base       LvDrawDscBaseT
	Color      LvColorT
	Width      c.Int32T
	StartAngle LvValuePreciseT
	EndAngle   LvValuePreciseT
	Center     LvPointT
	ImgSrc     c.Pointer
	Radius     c.Uint16T
	Opa        LvOpaT
	Rounded    c.Uint8T
}

/**
 * Initialize an arc draw descriptor.
 * @param dsc       pointer to a draw descriptor
 */
// llgo:link (*LvDrawArcDscT).LvDrawArcDscInit C.lv_draw_arc_dsc_init
func (recv_ *LvDrawArcDscT) LvDrawArcDscInit() {
}

/**
 * Try to get an arc draw descriptor from a draw task.
 * @param task      draw task
 * @return          the task's draw descriptor or NULL if the task is not of type LV_DRAW_TASK_TYPE_ARC
 */
// llgo:link (*LvDrawTaskT).LvDrawTaskGetArcDsc C.lv_draw_task_get_arc_dsc
func (recv_ *LvDrawTaskT) LvDrawTaskGetArcDsc() *LvDrawArcDscT {
	return nil
}

/**
 * Create an arc draw task.
 * @param layer         pointer to a layer
 * @param dsc           pointer to an initialized draw descriptor variable
 */
// llgo:link (*LvLayerT).LvDrawArc C.lv_draw_arc
func (recv_ *LvLayerT) LvDrawArc(dsc *LvDrawArcDscT) {
}

/**
 * Get an area the should be invalidated when the arcs angle changed between start_angle and end_ange
 * @param x             the x coordinate of the center of the arc
 * @param y             the y coordinate of the center of the arc
 * @param radius        the radius of the arc
 * @param start_angle   the start angle of the arc (0 deg on the bottom, 90 deg on the right)
 * @param end_angle     the end angle of the arc
 * @param w             width of the arc
 * @param rounded       true: the arc is rounded
 * @param area          store the area to invalidate here
 */
//go:linkname LvDrawArcGetArea C.lv_draw_arc_get_area
func LvDrawArcGetArea(x c.Int32T, y c.Int32T, radius c.Uint16T, start_angle LvValuePreciseT, end_angle LvValuePreciseT, w c.Int32T, rounded bool, area *LvAreaT)

/**
 * Converts an I1 buffer to ARGB8888 format.
 * @param buf_i1              pointer to buffer with I1 formatted render
 * @param buf_argb8888        pointer to buffer for ARGB8888 render
 * @param width               width in pixels of the area.
 *                            must be a multiple of 8.
 * @param height              height in pixels of the area
 * @param buf_i1_stride       stride of i1 buffer in bytes
 * @param buf_argb8888_stride stride of argb8888 buffer in bytes
 * @param index0_color        color of the 0 bits of i1 buf
 * @param index1_color        color of the 1 bits of i1 buf
 */
//go:linkname LvDrawSwI1ToArgb8888 C.lv_draw_sw_i1_to_argb8888
func LvDrawSwI1ToArgb8888(buf_i1 c.Pointer, buf_argb8888 c.Pointer, width c.Uint32T, height c.Uint32T, buf_i1_stride c.Uint32T, buf_argb8888_stride c.Uint32T, index0_color c.Uint32T, index1_color c.Uint32T)

/**
 * Swap the upper and lower byte of an RGB565 buffer.
 * Might be required if a 8bit parallel port or an SPI port send the bytes in the wrong order.
 * The bytes will be swapped in place.
 * @param buf           pointer to buffer
 * @param buf_size_px   number of pixels in the buffer
 */
//go:linkname LvDrawSwRgb565Swap C.lv_draw_sw_rgb565_swap
func LvDrawSwRgb565Swap(buf c.Pointer, buf_size_px c.Uint32T)

/**
 * Invert a draw buffer in the I1 color format.
 * Conventionally, a bit is set to 1 during blending if the luminance is greater than 127.
 * Depending on the display controller used, you might want to have different behavior.
 * The inversion will be performed in place.
 * @param buf          pointer to the buffer to be inverted
 * @param buf_size     size of the buffer in bytes
 */
//go:linkname LvDrawSwI1Invert C.lv_draw_sw_i1_invert
func LvDrawSwI1Invert(buf c.Pointer, buf_size c.Uint32T)

/**
 * Convert a draw buffer in I1 color format from htiled (row-wise)
 * to vtiled (column-wise) buffer layout. The conversion assumes that the buffer width
 * and height is rounded to a multiple of 8.
 * @param buf           pointer to the buffer to be converted
 * @param buf_size      size of the buffer in bytes
 * @param width         width of the buffer
 * @param height        height of the buffer
 * @param out_buf       pointer to the output buffer
 * @param out_buf_size  size of the output buffer in bytes
 * @param bit_order_lsb bit order of the resulting vtiled buffer
 */
//go:linkname LvDrawSwI1ConvertToVtiled C.lv_draw_sw_i1_convert_to_vtiled
func LvDrawSwI1ConvertToVtiled(buf c.Pointer, buf_size c.Uint32T, width c.Uint32T, height c.Uint32T, out_buf c.Pointer, out_buf_size c.Uint32T, bit_order_lsb bool)

/**
 * Rotate a buffer into another buffer
 * @param src           the source buffer
 * @param dest          the destination buffer
 * @param src_width     source width in pixels
 * @param src_height    source height in pixels
 * @param src_stride     source stride in bytes (number of bytes in a row)
 * @param dest_stride   destination stride in bytes (number of bytes in a row)
 * @param rotation      LV_DISPLAY_ROTATION_0/90/180/270
 * @param color_format  LV_COLOR_FORMAT_RGB565/RGB888/XRGB8888/ARGB8888
 */
//go:linkname LvDrawSwRotate C.lv_draw_sw_rotate
func LvDrawSwRotate(src c.Pointer, dest c.Pointer, src_width c.Int32T, src_height c.Int32T, src_stride c.Int32T, dest_stride c.Int32T, rotation LvDisplayRotationT, color_format LvColorFormatT)

type LvDrawSwMaskResT c.Int

const (
	LV_DRAW_SW_MASK_RES_TRANSP     LvDrawSwMaskResT = 0
	LV_DRAW_SW_MASK_RES_FULL_COVER LvDrawSwMaskResT = 1
	LV_DRAW_SW_MASK_RES_CHANGED    LvDrawSwMaskResT = 2
	LV_DRAW_SW_MASK_RES_UNKNOWN    LvDrawSwMaskResT = 3
)

type LvDrawSwMaskTypeT c.Int

const (
	LV_DRAW_SW_MASK_TYPE_LINE   LvDrawSwMaskTypeT = 0
	LV_DRAW_SW_MASK_TYPE_ANGLE  LvDrawSwMaskTypeT = 1
	LV_DRAW_SW_MASK_TYPE_RADIUS LvDrawSwMaskTypeT = 2
	LV_DRAW_SW_MASK_TYPE_FADE   LvDrawSwMaskTypeT = 3
	LV_DRAW_SW_MASK_TYPE_MAP    LvDrawSwMaskTypeT = 4
)

type LvDrawSwMaskLineSideT c.Int

const (
	LV_DRAW_SW_MASK_LINE_SIDE_LEFT   LvDrawSwMaskLineSideT = 0
	LV_DRAW_SW_MASK_LINE_SIDE_RIGHT  LvDrawSwMaskLineSideT = 1
	LV_DRAW_SW_MASK_LINE_SIDE_TOP    LvDrawSwMaskLineSideT = 2
	LV_DRAW_SW_MASK_LINE_SIDE_BOTTOM LvDrawSwMaskLineSideT = 3
)

// llgo:type C
type LvDrawSwMaskXcbT func(*LvOpaT, c.Int32T, c.Int32T, c.Int32T, c.Pointer) LvDrawSwMaskResT

/**********************
 * GLOBAL PROTOTYPES
 **********************/
//go:linkname LvDrawSwMaskInit C.lv_draw_sw_mask_init
func LvDrawSwMaskInit()

//go:linkname LvDrawSwMaskDeinit C.lv_draw_sw_mask_deinit
func LvDrawSwMaskDeinit()

//go:linkname LvDrawSwMaskApply C.lv_draw_sw_mask_apply
func LvDrawSwMaskApply(masks *c.Pointer, mask_buf *LvOpaT, abs_x c.Int32T, abs_y c.Int32T, len c.Int32T) LvDrawSwMaskResT

/**
 * Free the data from the parameter.
 * It's called inside `lv_draw_sw_mask_remove_id` and `lv_draw_sw_mask_remove_custom`
 * Needs to be called only in special cases when the mask is not added by `lv_draw_mask_add`
 * and not removed by `lv_draw_mask_remove_id` or `lv_draw_mask_remove_custom`
 * @param p pointer to a mask parameter
 */
//go:linkname LvDrawSwMaskFreeParam C.lv_draw_sw_mask_free_param
func LvDrawSwMaskFreeParam(p c.Pointer)

/**
 *Initialize a line mask from two points.
 * @param param pointer to a `lv_draw_mask_param_t` to initialize
 * @param p1x X coordinate of the first point of the line
 * @param p1y Y coordinate of the first point of the line
 * @param p2x X coordinate of the second point of the line
 * @param p2y y coordinate of the second point of the line
 * @param side and element of `lv_draw_mask_line_side_t` to describe which side to keep.
 * With `LV_DRAW_MASK_LINE_SIDE_LEFT/RIGHT` and horizontal line all pixels are kept
 * With `LV_DRAW_MASK_LINE_SIDE_TOP/BOTTOM` and vertical line all pixels are kept
 */
// llgo:link (*LvDrawSwMaskLineParamT).LvDrawSwMaskLinePointsInit C.lv_draw_sw_mask_line_points_init
func (recv_ *LvDrawSwMaskLineParamT) LvDrawSwMaskLinePointsInit(p1x c.Int32T, p1y c.Int32T, p2x c.Int32T, p2y c.Int32T, side LvDrawSwMaskLineSideT) {
}

/**
 *Initialize a line mask from a point and an angle.
 * @param param  pointer to a `lv_draw_mask_param_t` to initialize
 * @param px     X coordinate of a point of the line
 * @param py     X coordinate of a point of the line
 * @param angle  right 0 deg, bottom: 90
 * @param side   an element of `lv_draw_mask_line_side_t` to describe which side to keep.
 * With `LV_DRAW_MASK_LINE_SIDE_LEFT/RIGHT` and horizontal line all pixels are kept
 * With `LV_DRAW_MASK_LINE_SIDE_TOP/BOTTOM` and vertical line all pixels are kept
 */
// llgo:link (*LvDrawSwMaskLineParamT).LvDrawSwMaskLineAngleInit C.lv_draw_sw_mask_line_angle_init
func (recv_ *LvDrawSwMaskLineParamT) LvDrawSwMaskLineAngleInit(px c.Int32T, py c.Int32T, angle c.Int16T, side LvDrawSwMaskLineSideT) {
}

/**
 * Initialize an angle mask.
 * @param param pointer to a `lv_draw_mask_param_t` to initialize
 * @param vertex_x X coordinate of the angle vertex (absolute coordinates)
 * @param vertex_y Y coordinate of the angle vertex (absolute coordinates)
 * @param start_angle start angle in degrees. 0 deg on the right, 90 deg, on the bottom
 * @param end_angle end angle
 */
// llgo:link (*LvDrawSwMaskAngleParamT).LvDrawSwMaskAngleInit C.lv_draw_sw_mask_angle_init
func (recv_ *LvDrawSwMaskAngleParamT) LvDrawSwMaskAngleInit(vertex_x c.Int32T, vertex_y c.Int32T, start_angle c.Int32T, end_angle c.Int32T) {
}

/**
 * Initialize a fade mask.
 * @param param pointer to an `lv_draw_mask_radius_param_t` to initialize
 * @param rect coordinates of the rectangle to affect (absolute coordinates)
 * @param radius radius of the rectangle
 * @param inv true: keep the pixels inside the rectangle; keep the pixels outside of the rectangle
 */
// llgo:link (*LvDrawSwMaskRadiusParamT).LvDrawSwMaskRadiusInit C.lv_draw_sw_mask_radius_init
func (recv_ *LvDrawSwMaskRadiusParamT) LvDrawSwMaskRadiusInit(rect *LvAreaT, radius c.Int32T, inv bool) {
}

/**
 * Initialize a fade mask.
 * @param param pointer to a `lv_draw_mask_param_t` to initialize
 * @param coords coordinates of the area to affect (absolute coordinates)
 * @param opa_top opacity on the top
 * @param y_top at which coordinate start to change to opacity to `opa_bottom`
 * @param opa_bottom opacity at the bottom
 * @param y_bottom at which coordinate reach `opa_bottom`.
 */
// llgo:link (*LvDrawSwMaskFadeParamT).LvDrawSwMaskFadeInit C.lv_draw_sw_mask_fade_init
func (recv_ *LvDrawSwMaskFadeParamT) LvDrawSwMaskFadeInit(coords *LvAreaT, opa_top LvOpaT, y_top c.Int32T, opa_bottom LvOpaT, y_bottom c.Int32T) {
}

/**
 * Initialize a map mask.
 * @param param pointer to a `lv_draw_mask_param_t` to initialize
 * @param coords coordinates of the map (absolute coordinates)
 * @param map array of bytes with the mask values
 */
// llgo:link (*LvDrawSwMaskMapParamT).LvDrawSwMaskMapInit C.lv_draw_sw_mask_map_init
func (recv_ *LvDrawSwMaskMapParamT) LvDrawSwMaskMapInit(coords *LvAreaT, map_ *LvOpaT) {
}

// llgo:type C
type LvDrawSwBlendHandlerT func(*LvDrawTaskT, *LvDrawSwBlendDscT)

type LvDrawSwCustomBlendHandlerT struct {
	DestCf  LvColorFormatT
	Handler LvDrawSwBlendHandlerT
}

/**
 * Call the blend function of the `layer`.
 * @param draw_unit     pointer to a draw unit
 * @param dsc           pointer to an initialized blend descriptor
 */
// llgo:link (*LvDrawTaskT).LvDrawSwBlend C.lv_draw_sw_blend
func (recv_ *LvDrawTaskT) LvDrawSwBlend(dsc *LvDrawSwBlendDscT) {
}

/**
 * Initialize the SW renderer. Called in internally.
 * It creates as many SW renderers as defined in LV_DRAW_SW_DRAW_UNIT_CNT
 */
//go:linkname LvDrawSwInit C.lv_draw_sw_init
func LvDrawSwInit()

/**
 * Deinitialize the SW renderers
 */
//go:linkname LvDrawSwDeinit C.lv_draw_sw_deinit
func LvDrawSwDeinit()

/**
 * Fill an area using SW render. Handle gradient and radius.
 * @param t             pointer to a draw task
 * @param dsc           the draw descriptor
 * @param coords        the coordinates of the rectangle
 */
// llgo:link (*LvDrawTaskT).LvDrawSwFill C.lv_draw_sw_fill
func (recv_ *LvDrawTaskT) LvDrawSwFill(dsc *LvDrawFillDscT, coords *LvAreaT) {
}

/**
 * Draw border with SW render.
 * @param t             pointer to a draw task
 * @param dsc           the draw descriptor
 * @param coords        the coordinates of the rectangle
 */
// llgo:link (*LvDrawTaskT).LvDrawSwBorder C.lv_draw_sw_border
func (recv_ *LvDrawTaskT) LvDrawSwBorder(dsc *LvDrawBorderDscT, coords *LvAreaT) {
}

/**
 * Draw box shadow with SW render.
 * @param t             pointer to a draw task
 * @param dsc           the draw descriptor
 * @param coords        the coordinates of the rectangle for which the box shadow should be drawn
 */
// llgo:link (*LvDrawTaskT).LvDrawSwBoxShadow C.lv_draw_sw_box_shadow
func (recv_ *LvDrawTaskT) LvDrawSwBoxShadow(dsc *LvDrawBoxShadowDscT, coords *LvAreaT) {
}

/**
 * Draw an image with SW render. It handles image decoding, tiling, transformations, and recoloring.
 * @param t             pointer to a draw task
 * @param draw_dsc      the draw descriptor
 * @param coords        the coordinates of the image
 */
// llgo:link (*LvDrawTaskT).LvDrawSwImage C.lv_draw_sw_image
func (recv_ *LvDrawTaskT) LvDrawSwImage(draw_dsc *LvDrawImageDscT, coords *LvAreaT) {
}

// llgo:link (*LvDrawTaskT).LvDrawSwLetter C.lv_draw_sw_letter
func (recv_ *LvDrawTaskT) LvDrawSwLetter(dsc *LvDrawLetterDscT, coords *LvAreaT) {
}

/**
 * Draw a label with SW render.
 * @param t             pointer to a draw task
 * @param dsc           the draw descriptor
 * @param coords        the coordinates of the label
 */
// llgo:link (*LvDrawTaskT).LvDrawSwLabel C.lv_draw_sw_label
func (recv_ *LvDrawTaskT) LvDrawSwLabel(dsc *LvDrawLabelDscT, coords *LvAreaT) {
}

/**
 * Draw an arc with SW render.
 * @param t             pointer to a draw task
 * @param dsc           the draw descriptor
 * @param coords        the coordinates of the arc
 */
// llgo:link (*LvDrawTaskT).LvDrawSwArc C.lv_draw_sw_arc
func (recv_ *LvDrawTaskT) LvDrawSwArc(dsc *LvDrawArcDscT, coords *LvAreaT) {
}

/**
 * Draw a line with SW render.
 * @param t             pointer to a draw task
 * @param dsc           the draw descriptor
 */
// llgo:link (*LvDrawTaskT).LvDrawSwLine C.lv_draw_sw_line
func (recv_ *LvDrawTaskT) LvDrawSwLine(dsc *LvDrawLineDscT) {
}

/**
 * Blend a layer with SW render
 * @param t             pointer to a draw task
 * @param draw_dsc      the draw descriptor
 * @param coords        the coordinates of the layer
 */
// llgo:link (*LvDrawTaskT).LvDrawSwLayer C.lv_draw_sw_layer
func (recv_ *LvDrawTaskT) LvDrawSwLayer(draw_dsc *LvDrawImageDscT, coords *LvAreaT) {
}

/**
 * Draw a triangle with SW render.
 * @param t             pointer to a draw task
 * @param dsc           the draw descriptor
 */
// llgo:link (*LvDrawTaskT).LvDrawSwTriangle C.lv_draw_sw_triangle
func (recv_ *LvDrawTaskT) LvDrawSwTriangle(dsc *LvDrawTriangleDscT) {
}

/**
 * Mask out a rectangle with radius from a current layer
 * @param t             pointer to a draw task
 * @param dsc           the draw descriptor
 * @param coords        the coordinates of the mask
 */
// llgo:link (*LvDrawTaskT).LvDrawSwMaskRect C.lv_draw_sw_mask_rect
func (recv_ *LvDrawTaskT) LvDrawSwMaskRect(dsc *LvDrawMaskRectDscT) {
}

/**
 * Used internally to get a transformed are of an image
 * @param dest_area     area to calculate, i.e. get this area from the transformed image
 * @param src_buf       source buffer
 * @param src_w         source buffer width in pixels
 * @param src_h         source buffer height in pixels
 * @param src_stride    source buffer stride in bytes
 * @param draw_dsc      draw descriptor
 * @param sup           supplementary data
 * @param cf            color format of the source buffer
 * @param dest_buf      the destination buffer
 */
// llgo:link (*LvAreaT).LvDrawSwTransform C.lv_draw_sw_transform
func (recv_ *LvAreaT) LvDrawSwTransform(src_buf c.Pointer, src_w c.Int32T, src_h c.Int32T, src_stride c.Int32T, draw_dsc *LvDrawImageDscT, sup *LvDrawImageSupT, cf LvColorFormatT, dest_buf c.Pointer) {
}

/**
 * Register a custom blend handler for a color format.
 * Handler will be called when blending a color or an
 * image to a buffer with the given color format.
 * At most one handler can be registered for a color format.
 * Subsequent registrations will overwrite the previous handler.
 *
 * @param handler pointer to a blend handler
 * @return true if the handler was registered, false if the handler could not be registered
 */
// llgo:link (*LvDrawSwCustomBlendHandlerT).LvDrawSwRegisterBlendHandler C.lv_draw_sw_register_blend_handler
func (recv_ *LvDrawSwCustomBlendHandlerT) LvDrawSwRegisterBlendHandler() bool {
	return false
}

/**
 * Unregister a custom blend handler for a color format.
 * @param dest_cf color format
 * @return true if a handler was unregistered, false if no handler was registered
 */
// llgo:link LvColorFormatT.LvDrawSwUnregisterBlendHandler C.lv_draw_sw_unregister_blend_handler
func (recv_ LvColorFormatT) LvDrawSwUnregisterBlendHandler() bool {
	return false
}

/**
 * Get the blend handler for a color format.
 * @param dest_cf color format
 * @return pointer to the blend handler or NULL if no handler is registered
 */
// llgo:link LvColorFormatT.LvDrawSwGetBlendHandler C.lv_draw_sw_get_blend_handler
func (recv_ LvColorFormatT) LvDrawSwGetBlendHandler() LvDrawSwBlendHandlerT {
	return nil
}

type LvObjTreeWalkResT c.Int

const (
	LV_OBJ_TREE_WALK_NEXT          LvObjTreeWalkResT = 0
	LV_OBJ_TREE_WALK_SKIP_CHILDREN LvObjTreeWalkResT = 1
	LV_OBJ_TREE_WALK_END           LvObjTreeWalkResT = 2
)

// llgo:type C
type LvObjTreeWalkCbT func(*LvObjT, c.Pointer) LvObjTreeWalkResT

/**
 * Delete an object and all of its children.
 * Also remove the objects from their group and remove all animations (if any).
 * Send `LV_EVENT_DELETED` to deleted objects.
 * @param obj       pointer to an object
 */
// llgo:link (*LvObjT).LvObjDelete C.lv_obj_delete
func (recv_ *LvObjT) LvObjDelete() {
}

/**
 * Delete all children of an object.
 * Also remove the objects from their group and remove all animations (if any).
 * Send `LV_EVENT_DELETED` to deleted objects.
 * @param obj       pointer to an object
 */
// llgo:link (*LvObjT).LvObjClean C.lv_obj_clean
func (recv_ *LvObjT) LvObjClean() {
}

/**
 * Delete an object after some delay
 * @param obj       pointer to an object
 * @param delay_ms  time to wait before delete in milliseconds
 */
// llgo:link (*LvObjT).LvObjDeleteDelayed C.lv_obj_delete_delayed
func (recv_ *LvObjT) LvObjDeleteDelayed(delay_ms c.Uint32T) {
}

/**
 * A function to be easily used in animation ready callback to delete an object when the animation is ready
 * @param a         pointer to the animation
 */
// llgo:link (*LvAnimT).LvObjDeleteAnimCompletedCb C.lv_obj_delete_anim_completed_cb
func (recv_ *LvAnimT) LvObjDeleteAnimCompletedCb() {
}

/**
 * Helper function for asynchronously deleting objects.
 * Useful for cases where you can't delete an object directly in an `LV_EVENT_DELETE` handler (i.e. parent).
 * @param obj       object to delete
 * @see lv_async_call
 */
// llgo:link (*LvObjT).LvObjDeleteAsync C.lv_obj_delete_async
func (recv_ *LvObjT) LvObjDeleteAsync() {
}

/**
 * Move the parent of an object. The relative coordinates will be kept.
 *
 * @param obj       pointer to an object whose parent needs to be changed
 * @param parent pointer to the new parent
 */
// llgo:link (*LvObjT).LvObjSetParent C.lv_obj_set_parent
func (recv_ *LvObjT) LvObjSetParent(parent *LvObjT) {
}

/**
 * Swap the positions of two objects.
 * When used in listboxes, it can be used to sort the listbox items.
 * @param obj1  pointer to the first object
 * @param obj2  pointer to the second object
 */
// llgo:link (*LvObjT).LvObjSwap C.lv_obj_swap
func (recv_ *LvObjT) LvObjSwap(obj2 *LvObjT) {
}

/**
 * moves the object to the given index in its parent.
 * When used in listboxes, it can be used to sort the listbox items.
 * @param obj  pointer to the object to be moved.
 * @param index  new index in parent. -1 to count from the back
 * @note to move to the background: lv_obj_move_to_index(obj, 0)
 * @note to move forward (up): lv_obj_move_to_index(obj, lv_obj_get_index(obj) - 1)
 */
// llgo:link (*LvObjT).LvObjMoveToIndex C.lv_obj_move_to_index
func (recv_ *LvObjT) LvObjMoveToIndex(index c.Int32T) {
}

/**
 * Get the screen of an object
 * @param obj       pointer to an object
 * @return          pointer to the object's screen
 */
// llgo:link (*LvObjT).LvObjGetScreen C.lv_obj_get_screen
func (recv_ *LvObjT) LvObjGetScreen() *LvObjT {
	return nil
}

/**
 * Get the display of the object
 * @param obj       pointer to an object
 * @return          pointer to the object's display
 */
// llgo:link (*LvObjT).LvObjGetDisplay C.lv_obj_get_display
func (recv_ *LvObjT) LvObjGetDisplay() *LvDisplayT {
	return nil
}

/**
 * Get the parent of an object
 * @param obj       pointer to an object
 * @return          the parent of the object. (NULL if `obj` was a screen)
 */
// llgo:link (*LvObjT).LvObjGetParent C.lv_obj_get_parent
func (recv_ *LvObjT) LvObjGetParent() *LvObjT {
	return nil
}

/**
 * Get the child of an object by the child's index.
 * @param obj       pointer to an object whose child should be get
 * @param idx       the index of the child.
 *                  0: the oldest (firstly created) child
 *                  1: the second oldest
 *                  child count-1: the youngest
 *                  -1: the youngest
 *                  -2: the second youngest
 * @return          pointer to the child or NULL if the index was invalid
 */
// llgo:link (*LvObjT).LvObjGetChild C.lv_obj_get_child
func (recv_ *LvObjT) LvObjGetChild(idx c.Int32T) *LvObjT {
	return nil
}

/**
 * Get the child of an object by the child's index. Consider the children only with a given type.
 * @param obj       pointer to an object whose child should be get
 * @param idx       the index of the child.
 *                  0: the oldest (firstly created) child
 *                  1: the second oldest
 *                  child count-1: the youngest
 *                  -1: the youngest
 *                  -2: the second youngest
 * @param class_p   the type of the children to check
 * @return          pointer to the child or NULL if the index was invalid
 */
// llgo:link (*LvObjT).LvObjGetChildByType C.lv_obj_get_child_by_type
func (recv_ *LvObjT) LvObjGetChildByType(idx c.Int32T, class_p *LvObjClassT) *LvObjT {
	return nil
}

/**
 * Return a sibling of an object
 * @param obj       pointer to an object whose sibling should be get
 * @param idx       0: `obj` itself
 *                  -1: the first older sibling
 *                  -2: the next older sibling
 *                  1: the first younger sibling
 *                  2: the next younger sibling
 *                  etc
 * @return          pointer to the requested sibling  or NULL if there is no such sibling
 */
// llgo:link (*LvObjT).LvObjGetSibling C.lv_obj_get_sibling
func (recv_ *LvObjT) LvObjGetSibling(idx c.Int32T) *LvObjT {
	return nil
}

/**
 * Return a sibling of an object. Consider the siblings only with a given type.
 * @param obj       pointer to an object whose sibling should be get
 * @param idx       0: `obj` itself
 *                  -1: the first older sibling
 *                  -2: the next older sibling
 *                  1: the first younger sibling
 *                  2: the next younger sibling
 *                  etc
 * @param class_p   the type of the children to check
 * @return          pointer to the requested sibling  or NULL if there is no such sibling
 */
// llgo:link (*LvObjT).LvObjGetSiblingByType C.lv_obj_get_sibling_by_type
func (recv_ *LvObjT) LvObjGetSiblingByType(idx c.Int32T, class_p *LvObjClassT) *LvObjT {
	return nil
}

/**
 * Get the number of children
 * @param obj       pointer to an object
 * @return          the number of children
 */
// llgo:link (*LvObjT).LvObjGetChildCount C.lv_obj_get_child_count
func (recv_ *LvObjT) LvObjGetChildCount() c.Uint32T {
	return 0
}

/**
 * Get the number of children having a given type.
 * @param obj       pointer to an object
 * @param class_p   the type of the children to check
 * @return          the number of children
 */
// llgo:link (*LvObjT).LvObjGetChildCountByType C.lv_obj_get_child_count_by_type
func (recv_ *LvObjT) LvObjGetChildCountByType(class_p *LvObjClassT) c.Uint32T {
	return 0
}

/**
 * Get the index of a child.
 * @param obj       pointer to an object
 * @return          the child index of the object.
 *                  E.g. 0: the oldest (firstly created child).
 *                  (-1 if child could not be found or no parent exists)
 */
// llgo:link (*LvObjT).LvObjGetIndex C.lv_obj_get_index
func (recv_ *LvObjT) LvObjGetIndex() c.Int32T {
	return 0
}

/**
 * Get the index of a child. Consider the children only with a given type.
 * @param obj       pointer to an object
 * @param class_p   the type of the children to check
 * @return          the child index of the object.
 *                  E.g. 0: the oldest (firstly created child with the given class).
 *                  (-1 if child could not be found or no parent exists)
 */
// llgo:link (*LvObjT).LvObjGetIndexByType C.lv_obj_get_index_by_type
func (recv_ *LvObjT) LvObjGetIndexByType(class_p *LvObjClassT) c.Int32T {
	return 0
}

/**
 * Iterate through all children of any object.
 * @param start_obj     start integrating from this object
 * @param cb            call this callback on the objects
 * @param user_data     pointer to any user related data (will be passed to `cb`)
 */
// llgo:link (*LvObjT).LvObjTreeWalk C.lv_obj_tree_walk
func (recv_ *LvObjT) LvObjTreeWalk(cb LvObjTreeWalkCbT, user_data c.Pointer) {
}

/**
 * Iterate through all children of any object and print their ID.
 * @param start_obj     start integrating from this object
 */
// llgo:link (*LvObjT).LvObjDumpTree C.lv_obj_dump_tree
func (recv_ *LvObjT) LvObjDumpTree() {
}

type LvObjPointTransformFlagT c.Int

const (
	LV_OBJ_POINT_TRANSFORM_FLAG_NONE              LvObjPointTransformFlagT = 0
	LV_OBJ_POINT_TRANSFORM_FLAG_RECURSIVE         LvObjPointTransformFlagT = 1
	LV_OBJ_POINT_TRANSFORM_FLAG_INVERSE           LvObjPointTransformFlagT = 2
	LV_OBJ_POINT_TRANSFORM_FLAG_INVERSE_RECURSIVE LvObjPointTransformFlagT = 3
)

/**
 * Set the position of an object relative to the set alignment.
 * @param obj       pointer to an object
 * @param x         new x coordinate
 * @param y         new y coordinate
 * @note            With default alignment it's the distance from the top left corner
 * @note            E.g. LV_ALIGN_CENTER alignment it's the offset from the center of the parent
 * @note            The position is interpreted on the content area of the parent
 * @note            The values can be set in pixel or in percentage of parent size with `lv_pct(v)`
 */
// llgo:link (*LvObjT).LvObjSetPos C.lv_obj_set_pos
func (recv_ *LvObjT) LvObjSetPos(x c.Int32T, y c.Int32T) {
}

/**
 * Set the x coordinate of an object
 * @param obj       pointer to an object
 * @param x         new x coordinate
 * @note            With default alignment it's the distance from the top left corner
 * @note            E.g. LV_ALIGN_CENTER alignment it's the offset from the center of the parent
 * @note            The position is interpreted on the content area of the parent
 * @note            The values can be set in pixel or in percentage of parent size with `lv_pct(v)`
 */
// llgo:link (*LvObjT).LvObjSetX C.lv_obj_set_x
func (recv_ *LvObjT) LvObjSetX(x c.Int32T) {
}

/**
 * Set the y coordinate of an object
 * @param obj       pointer to an object
 * @param y         new y coordinate
 * @note            With default alignment it's the distance from the top left corner
 * @note            E.g. LV_ALIGN_CENTER alignment it's the offset from the center of the parent
 * @note            The position is interpreted on the content area of the parent
 * @note            The values can be set in pixel or in percentage of parent size with `lv_pct(v)`
 */
// llgo:link (*LvObjT).LvObjSetY C.lv_obj_set_y
func (recv_ *LvObjT) LvObjSetY(y c.Int32T) {
}

/**
 * Set the size of an object.
 * @param obj       pointer to an object
 * @param w         the new width
 * @param h         the new height
 * @note            possible values are:
 *                  pixel               simple set the size accordingly
 *                  LV_SIZE_CONTENT     set the size to involve all children in the given direction
 *                  lv_pct(x)           to set size in percentage of the parent's content area size (the size without paddings).
 *                                      x should be in [0..1000]% range
 */
// llgo:link (*LvObjT).LvObjSetSize C.lv_obj_set_size
func (recv_ *LvObjT) LvObjSetSize(w c.Int32T, h c.Int32T) {
}

/**
 * Recalculate the size of the object
 * @param obj       pointer to an object
 * @return          true: the size has been changed
 */
// llgo:link (*LvObjT).LvObjRefrSize C.lv_obj_refr_size
func (recv_ *LvObjT) LvObjRefrSize() bool {
	return false
}

/**
 * Set the width of an object
 * @param obj       pointer to an object
 * @param w         the new width
 * @note            possible values are:
 *                  pixel               simple set the size accordingly
 *                  LV_SIZE_CONTENT     set the size to involve all children in the given direction
 *                  lv_pct(x)           to set size in percentage of the parent's content area size (the size without paddings).
 *                                      x should be in [0..1000]% range
 */
// llgo:link (*LvObjT).LvObjSetWidth C.lv_obj_set_width
func (recv_ *LvObjT) LvObjSetWidth(w c.Int32T) {
}

/**
 * Set the height of an object
 * @param obj       pointer to an object
 * @param h         the new height
 * @note            possible values are:
 *                  pixel               simple set the size accordingly
 *                  LV_SIZE_CONTENT     set the size to involve all children in the given direction
 *                  lv_pct(x)           to set size in percentage of the parent's content area size (the size without paddings).
 *                                      x should be in [0..1000]% range
 */
// llgo:link (*LvObjT).LvObjSetHeight C.lv_obj_set_height
func (recv_ *LvObjT) LvObjSetHeight(h c.Int32T) {
}

/**
 * Set the width reduced by the left and right padding and the border width.
 * @param obj       pointer to an object
 * @param w         the width without paddings in pixels
 */
// llgo:link (*LvObjT).LvObjSetContentWidth C.lv_obj_set_content_width
func (recv_ *LvObjT) LvObjSetContentWidth(w c.Int32T) {
}

/**
 * Set the height reduced by the top and bottom padding and the border width.
 * @param obj       pointer to an object
 * @param h         the height without paddings in pixels
 */
// llgo:link (*LvObjT).LvObjSetContentHeight C.lv_obj_set_content_height
func (recv_ *LvObjT) LvObjSetContentHeight(h c.Int32T) {
}

/**
 * Set a layout for an object
 * @param obj       pointer to an object
 * @param layout    pointer to a layout descriptor to set
 */
// llgo:link (*LvObjT).LvObjSetLayout C.lv_obj_set_layout
func (recv_ *LvObjT) LvObjSetLayout(layout c.Uint32T) {
}

/**
 * Test whether the and object is positioned by a layout or not
 * @param obj       pointer to an object to test
 * @return true:    positioned by a layout; false: not positioned by a layout
 */
// llgo:link (*LvObjT).LvObjIsLayoutPositioned C.lv_obj_is_layout_positioned
func (recv_ *LvObjT) LvObjIsLayoutPositioned() bool {
	return false
}

/**
 * Mark the object for layout update.
 * @param obj      pointer to an object whose children need to be updated
 */
// llgo:link (*LvObjT).LvObjMarkLayoutAsDirty C.lv_obj_mark_layout_as_dirty
func (recv_ *LvObjT) LvObjMarkLayoutAsDirty() {
}

/**
 * Update the layout of an object.
 * @param obj      pointer to an object whose position and size needs to be updated
 */
// llgo:link (*LvObjT).LvObjUpdateLayout C.lv_obj_update_layout
func (recv_ *LvObjT) LvObjUpdateLayout() {
}

/**
 * Change the alignment of an object.
 * @param obj       pointer to an object to align
 * @param align     type of alignment (see 'lv_align_t' enum) `LV_ALIGN_OUT_...` can't be used.
 */
// llgo:link (*LvObjT).LvObjSetAlign C.lv_obj_set_align
func (recv_ *LvObjT) LvObjSetAlign(align LvAlignT) {
}

/**
 * Change the alignment of an object and set new coordinates.
 * Equivalent to:
 * lv_obj_set_align(obj, align);
 * lv_obj_set_pos(obj, x_ofs, y_ofs);
 * @param obj       pointer to an object to align
 * @param align     type of alignment (see 'lv_align_t' enum) `LV_ALIGN_OUT_...` can't be used.
 * @param x_ofs     x coordinate offset after alignment
 * @param y_ofs     y coordinate offset after alignment
 */
// llgo:link (*LvObjT).LvObjAlign C.lv_obj_align
func (recv_ *LvObjT) LvObjAlign(align LvAlignT, x_ofs c.Int32T, y_ofs c.Int32T) {
}

/**
 * Align an object to another object.
 * @param obj       pointer to an object to align
 * @param base      pointer to another object (if NULL `obj`s parent is used). 'obj' will be aligned to it.
 * @param align     type of alignment (see 'lv_align_t' enum)
 * @param x_ofs     x coordinate offset after alignment
 * @param y_ofs     y coordinate offset after alignment
 * @note            if the position or size of `base` changes `obj` needs to be aligned manually again
 */
// llgo:link (*LvObjT).LvObjAlignTo C.lv_obj_align_to
func (recv_ *LvObjT) LvObjAlignTo(base *LvObjT, align LvAlignT, x_ofs c.Int32T, y_ofs c.Int32T) {
}

/**
 * Align an object to the center on its parent.
 * @param obj       pointer to an object to align
 * @note            if the parent size changes `obj` needs to be aligned manually again
 */
// llgo:link (*LvObjT).LvObjCenter C.lv_obj_center
func (recv_ *LvObjT) LvObjCenter() {
}

/**
 * Set the transform matrix of an object
 * @param obj       pointer to an object
 * @param matrix    pointer to a matrix to set
 * @note `LV_DRAW_TRANSFORM_USE_MATRIX` needs to be enabled.
 */
// llgo:link (*LvObjT).LvObjSetTransform C.lv_obj_set_transform
func (recv_ *LvObjT) LvObjSetTransform(matrix *LvMatrixT) {
}

/**
 * Reset the transform matrix of an object to identity matrix
 * @param obj       pointer to an object
 * @note `LV_DRAW_TRANSFORM_USE_MATRIX` needs to be enabled.
 */
// llgo:link (*LvObjT).LvObjResetTransform C.lv_obj_reset_transform
func (recv_ *LvObjT) LvObjResetTransform() {
}

/**
 * Copy the coordinates of an object to an area
 * @param obj       pointer to an object
 * @param coords    pointer to an area to store the coordinates
 */
// llgo:link (*LvObjT).LvObjGetCoords C.lv_obj_get_coords
func (recv_ *LvObjT) LvObjGetCoords(coords *LvAreaT) {
}

/**
 * Get the x coordinate of object.
 * @param obj       pointer to an object
 * @return          distance of `obj` from the left side of its parent plus the parent's left padding
 * @note            The position of the object is recalculated only on the next redraw. To force coordinate recalculation
 *                  call `lv_obj_update_layout(obj)`.
 * @note            Zero return value means the object is on the left padding of the parent, and not on the left edge.
 * @note            Scrolling of the parent doesn't change the returned value.
 * @note            The returned value is always the distance from the parent even if `obj` is positioned by a layout.
 */
// llgo:link (*LvObjT).LvObjGetX C.lv_obj_get_x
func (recv_ *LvObjT) LvObjGetX() c.Int32T {
	return 0
}

/**
 * Get the x2 coordinate of object.
 * @param obj       pointer to an object
 * @return          distance of `obj` from the right side of its parent plus the parent's right padding
 * @note            The position of the object is recalculated only on the next redraw. To force coordinate recalculation
 *                  call `lv_obj_update_layout(obj)`.
 * @note            Zero return value means the object is on the right padding of the parent, and not on the right edge.
 * @note            Scrolling of the parent doesn't change the returned value.
 * @note            The returned value is always the distance from the parent even if `obj` is positioned by a layout.
 */
// llgo:link (*LvObjT).LvObjGetX2 C.lv_obj_get_x2
func (recv_ *LvObjT) LvObjGetX2() c.Int32T {
	return 0
}

/**
 * Get the y coordinate of object.
 * @param obj       pointer to an object
 * @return          distance of `obj` from the top side of its parent plus the parent's top padding
 * @note            The position of the object is recalculated only on the next redraw. To force coordinate recalculation
 *                  call `lv_obj_update_layout(obj)`.
 * @note            Zero return value means the object is on the top padding of the parent, and not on the top edge.
 * @note            Scrolling of the parent doesn't change the returned value.
 * @note            The returned value is always the distance from the parent even if `obj` is positioned by a layout.
 */
// llgo:link (*LvObjT).LvObjGetY C.lv_obj_get_y
func (recv_ *LvObjT) LvObjGetY() c.Int32T {
	return 0
}

/**
 * Get the y2 coordinate of object.
 * @param obj       pointer to an object
 * @return          distance of `obj` from the bottom side of its parent plus the parent's bottom padding
 * @note            The position of the object is recalculated only on the next redraw. To force coordinate recalculation
 *                  call `lv_obj_update_layout(obj)`.
 * @note            Zero return value means the object is on the bottom padding of the parent, and not on the bottom edge.
 * @note            Scrolling of the parent doesn't change the returned value.
 * @note            The returned value is always the distance from the parent even if `obj` is positioned by a layout.
 */
// llgo:link (*LvObjT).LvObjGetY2 C.lv_obj_get_y2
func (recv_ *LvObjT) LvObjGetY2() c.Int32T {
	return 0
}

/**
 * Get the actually set x coordinate of object, i.e. the offset from the set alignment
 * @param obj       pointer to an object
 * @return          the set x coordinate
 */
// llgo:link (*LvObjT).LvObjGetXAligned C.lv_obj_get_x_aligned
func (recv_ *LvObjT) LvObjGetXAligned() c.Int32T {
	return 0
}

/**
 * Get the actually set y coordinate of object, i.e. the offset from the set alignment
 * @param obj       pointer to an object
 * @return          the set y coordinate
 */
// llgo:link (*LvObjT).LvObjGetYAligned C.lv_obj_get_y_aligned
func (recv_ *LvObjT) LvObjGetYAligned() c.Int32T {
	return 0
}

/**
 * Get the width of an object
 * @param obj       pointer to an object
 * @note            The position of the object is recalculated only on the next redraw. To force coordinate recalculation
 *                  call `lv_obj_update_layout(obj)`.
 * @return          the width in pixels
 */
// llgo:link (*LvObjT).LvObjGetWidth C.lv_obj_get_width
func (recv_ *LvObjT) LvObjGetWidth() c.Int32T {
	return 0
}

/**
 * Get the height of an object
 * @param obj       pointer to an object
 * @note            The position of the object is recalculated only on the next redraw. To force coordinate recalculation
 *                  call `lv_obj_update_layout(obj)`.
 * @return          the height in pixels
 */
// llgo:link (*LvObjT).LvObjGetHeight C.lv_obj_get_height
func (recv_ *LvObjT) LvObjGetHeight() c.Int32T {
	return 0
}

/**
 * Get the width reduced by the left and right padding and the border width.
 * @param obj       pointer to an object
 * @note            The position of the object is recalculated only on the next redraw. To force coordinate recalculation
 *                  call `lv_obj_update_layout(obj)`.
 * @return          the width which still fits into its parent without causing overflow (making the parent scrollable)
 */
// llgo:link (*LvObjT).LvObjGetContentWidth C.lv_obj_get_content_width
func (recv_ *LvObjT) LvObjGetContentWidth() c.Int32T {
	return 0
}

/**
 * Get the height reduced by the top and bottom padding and the border width.
 * @param obj       pointer to an object
 * @note            The position of the object is recalculated only on the next redraw. To force coordinate recalculation
 *                  call `lv_obj_update_layout(obj)`.
 * @return          the height which still fits into the parent without causing overflow (making the parent scrollable)
 */
// llgo:link (*LvObjT).LvObjGetContentHeight C.lv_obj_get_content_height
func (recv_ *LvObjT) LvObjGetContentHeight() c.Int32T {
	return 0
}

/**
 * Get the area reduced by the paddings and the border width.
 * @param obj       pointer to an object
 * @note            The position of the object is recalculated only on the next redraw. To force coordinate recalculation
 *                  call `lv_obj_update_layout(obj)`.
 * @param area      the area which still fits into the parent without causing overflow (making the parent scrollable)
 */
// llgo:link (*LvObjT).LvObjGetContentCoords C.lv_obj_get_content_coords
func (recv_ *LvObjT) LvObjGetContentCoords(area *LvAreaT) {
}

/**
 * Get the width occupied by the "parts" of the widget. E.g. the width of all columns of a table.
 * @param obj       pointer to an objects
 * @return          the width of the virtually drawn content
 * @note            This size independent from the real size of the widget.
 *                  It just tells how large the internal ("virtual") content is.
 */
// llgo:link (*LvObjT).LvObjGetSelfWidth C.lv_obj_get_self_width
func (recv_ *LvObjT) LvObjGetSelfWidth() c.Int32T {
	return 0
}

/**
 * Get the height occupied by the "parts" of the widget. E.g. the height of all rows of a table.
 * @param obj       pointer to an objects
 * @return          the width of the virtually drawn content
 * @note            This size independent from the real size of the widget.
 *                  It just tells how large the internal ("virtual") content is.
 */
// llgo:link (*LvObjT).LvObjGetSelfHeight C.lv_obj_get_self_height
func (recv_ *LvObjT) LvObjGetSelfHeight() c.Int32T {
	return 0
}

/**
 * Handle if the size of the internal ("virtual") content of an object has changed.
 * @param obj       pointer to an object
 * @return          false: nothing happened; true: refresh happened
 */
// llgo:link (*LvObjT).LvObjRefreshSelfSize C.lv_obj_refresh_self_size
func (recv_ *LvObjT) LvObjRefreshSelfSize() bool {
	return false
}

// llgo:link (*LvObjT).LvObjRefrPos C.lv_obj_refr_pos
func (recv_ *LvObjT) LvObjRefrPos() {
}

// llgo:link (*LvObjT).LvObjMoveTo C.lv_obj_move_to
func (recv_ *LvObjT) LvObjMoveTo(x c.Int32T, y c.Int32T) {
}

// llgo:link (*LvObjT).LvObjMoveChildrenBy C.lv_obj_move_children_by
func (recv_ *LvObjT) LvObjMoveChildrenBy(x_diff c.Int32T, y_diff c.Int32T, ignore_floating bool) {
}

/**
 * Get the transform matrix of an object
 * @param obj       pointer to an object
 * @return          pointer to the transform matrix or NULL if not set
 */
// llgo:link (*LvObjT).LvObjGetTransform C.lv_obj_get_transform
func (recv_ *LvObjT) LvObjGetTransform() *LvMatrixT {
	return nil
}

/**
 * Transform a point using the angle and zoom style properties of an object
 * @param obj           pointer to an object whose style properties should be used
 * @param p             a point to transform, the result will be written back here too
 * @param flags         OR-ed valued of :cpp:enum:`lv_obj_point_transform_flag_t`
 */
// llgo:link (*LvObjT).LvObjTransformPoint C.lv_obj_transform_point
func (recv_ *LvObjT) LvObjTransformPoint(p *LvPointT, flags LvObjPointTransformFlagT) {
}

/**
 * Transform an array of points using the angle and zoom style properties of an object
 * @param obj           pointer to an object whose style properties should be used
 * @param points        the array of points to transform, the result will be written back here too
 * @param count         number of points in the array
 * @param flags         OR-ed valued of :cpp:enum:`lv_obj_point_transform_flag_t`
 */
// llgo:link (*LvObjT).LvObjTransformPointArray C.lv_obj_transform_point_array
func (recv_ *LvObjT) LvObjTransformPointArray(points *LvPointT, count c.SizeT, flags LvObjPointTransformFlagT) {
}

/**
 * Transform an area using the angle and zoom style properties of an object
 * @param obj           pointer to an object whose style properties should be used
 * @param area          an area to transform, the result will be written back here too
 * @param flags         OR-ed valued of :cpp:enum:`lv_obj_point_transform_flag_t`
 */
// llgo:link (*LvObjT).LvObjGetTransformedArea C.lv_obj_get_transformed_area
func (recv_ *LvObjT) LvObjGetTransformedArea(area *LvAreaT, flags LvObjPointTransformFlagT) {
}

/**
 * Mark an area of an object as invalid.
 * The area will be truncated to the object's area and marked for redraw.
 * @param obj       pointer to an object
 * @param           area the area to redraw
 */
// llgo:link (*LvObjT).LvObjInvalidateArea C.lv_obj_invalidate_area
func (recv_ *LvObjT) LvObjInvalidateArea(area *LvAreaT) {
}

/**
 * Mark the object as invalid to redrawn its area
 * @param obj       pointer to an object
 */
// llgo:link (*LvObjT).LvObjInvalidate C.lv_obj_invalidate
func (recv_ *LvObjT) LvObjInvalidate() {
}

/**
 * Tell whether an area of an object is visible (even partially) now or not
 * @param obj       pointer to an object
 * @param area      the are to check. The visible part of the area will be written back here.
 * @return true     visible; false not visible (hidden, out of parent, on other screen, etc)
 */
// llgo:link (*LvObjT).LvObjAreaIsVisible C.lv_obj_area_is_visible
func (recv_ *LvObjT) LvObjAreaIsVisible(area *LvAreaT) bool {
	return false
}

/**
 * Tell whether an object is visible (even partially) now or not
 * @param obj       pointer to an object
 * @return      true: visible; false not visible (hidden, out of parent, on other screen, etc)
 */
// llgo:link (*LvObjT).LvObjIsVisible C.lv_obj_is_visible
func (recv_ *LvObjT) LvObjIsVisible() bool {
	return false
}

/**
 * Set the size of an extended clickable area
 * @param obj       pointer to an object
 * @param size      extended clickable area in all 4 directions [px]
 */
// llgo:link (*LvObjT).LvObjSetExtClickArea C.lv_obj_set_ext_click_area
func (recv_ *LvObjT) LvObjSetExtClickArea(size c.Int32T) {
}

/**
 * Get the an area where to object can be clicked.
 * It's the object's normal area plus the extended click area.
 * @param obj       pointer to an object
 * @param area      store the result area here
 */
// llgo:link (*LvObjT).LvObjGetClickArea C.lv_obj_get_click_area
func (recv_ *LvObjT) LvObjGetClickArea(area *LvAreaT) {
}

/**
 * Hit-test an object given a particular point in screen space.
 * @param obj       object to hit-test
 * @param point     screen-space point (absolute coordinate)
 * @return          true: if the object is considered under the point
 */
// llgo:link (*LvObjT).LvObjHitTest C.lv_obj_hit_test
func (recv_ *LvObjT) LvObjHitTest(point *LvPointT) bool {
	return false
}

/**
 * Clamp a width between min and max width. If the min/max width is in percentage value use the ref_width
 * @param width         width to clamp
 * @param min_width     the minimal width
 * @param max_width     the maximal width
 * @param ref_width     the reference width used when min/max width is in percentage
 * @return              the clamped width
 */
//go:linkname LvClampWidth C.lv_clamp_width
func LvClampWidth(width c.Int32T, min_width c.Int32T, max_width c.Int32T, ref_width c.Int32T) c.Int32T

/**
 * Clamp a height between min and max height. If the min/max height is in percentage value use the ref_height
 * @param height         height to clamp
 * @param min_height     the minimal height
 * @param max_height     the maximal height
 * @param ref_height     the reference height used when min/max height is in percentage
 * @return              the clamped height
 */
//go:linkname LvClampHeight C.lv_clamp_height
func LvClampHeight(height c.Int32T, min_height c.Int32T, max_height c.Int32T, ref_height c.Int32T) c.Int32T

type LvScrollbarModeT c.Int

const (
	LV_SCROLLBAR_MODE_OFF    LvScrollbarModeT = 0
	LV_SCROLLBAR_MODE_ON     LvScrollbarModeT = 1
	LV_SCROLLBAR_MODE_ACTIVE LvScrollbarModeT = 2
	LV_SCROLLBAR_MODE_AUTO   LvScrollbarModeT = 3
)

type LvScrollSnapT c.Int

const (
	LV_SCROLL_SNAP_NONE   LvScrollSnapT = 0
	LV_SCROLL_SNAP_START  LvScrollSnapT = 1
	LV_SCROLL_SNAP_END    LvScrollSnapT = 2
	LV_SCROLL_SNAP_CENTER LvScrollSnapT = 3
)

/**
 * Set how the scrollbars should behave.
 * @param obj       pointer to Widget
 * @param mode      LV_SCROLL_MODE_ON/OFF/AUTO/ACTIVE
 */
// llgo:link (*LvObjT).LvObjSetScrollbarMode C.lv_obj_set_scrollbar_mode
func (recv_ *LvObjT) LvObjSetScrollbarMode(mode LvScrollbarModeT) {
}

/**
 * Set direction Widget can be scrolled
 * @param obj       pointer to Widget
 * @param dir       one or more bit-wise OR-ed values of `lv_dir_t` enumeration
 */
// llgo:link (*LvObjT).LvObjSetScrollDir C.lv_obj_set_scroll_dir
func (recv_ *LvObjT) LvObjSetScrollDir(dir LvDirT) {
}

/**
 * Set where to snap the children when scrolling ends horizontally
 * @param obj       pointer to Widget
 * @param align     value from `lv_scroll_snap_t` enumeration
 */
// llgo:link (*LvObjT).LvObjSetScrollSnapX C.lv_obj_set_scroll_snap_x
func (recv_ *LvObjT) LvObjSetScrollSnapX(align LvScrollSnapT) {
}

/**
 * Set where to snap the children when scrolling ends vertically
 * @param obj       pointer to Widget
 * @param align     value from `lv_scroll_snap_t` enumeration
 */
// llgo:link (*LvObjT).LvObjSetScrollSnapY C.lv_obj_set_scroll_snap_y
func (recv_ *LvObjT) LvObjSetScrollSnapY(align LvScrollSnapT) {
}

/**
 * Get the current scroll mode (when to hide the scrollbars)
 * @param obj       pointer to Widget
 * @return          the current scroll mode from `lv_scrollbar_mode_t`
 */
// llgo:link (*LvObjT).LvObjGetScrollbarMode C.lv_obj_get_scrollbar_mode
func (recv_ *LvObjT) LvObjGetScrollbarMode() LvScrollbarModeT {
	return 0
}

/**
 * Get directions Widget can be scrolled (set with `lv_obj_set_scroll_dir()`)
 * @param obj       pointer to Widget
 * @return          current scroll direction bit(s)
 */
// llgo:link (*LvObjT).LvObjGetScrollDir C.lv_obj_get_scroll_dir
func (recv_ *LvObjT) LvObjGetScrollDir() LvDirT {
	return 0
}

/**
 * Get where to snap child Widgets when horizontal scrolling ends.
 * @param obj       pointer to Widget
 * @return          current snap value from `lv_scroll_snap_t`
 */
// llgo:link (*LvObjT).LvObjGetScrollSnapX C.lv_obj_get_scroll_snap_x
func (recv_ *LvObjT) LvObjGetScrollSnapX() LvScrollSnapT {
	return 0
}

/**
 * Get where to snap child Widgets when vertical scrolling ends.
 * @param  obj      pointer to Widget
 * @return          current snap value from `lv_scroll_snap_t`
 */
// llgo:link (*LvObjT).LvObjGetScrollSnapY C.lv_obj_get_scroll_snap_y
func (recv_ *LvObjT) LvObjGetScrollSnapY() LvScrollSnapT {
	return 0
}

/**
 * Get current X scroll position.  Identical to `lv_obj_get_scroll_left()`.
 * @param obj       pointer to scrollable container Widget
 * @return          current scroll position from left edge
 *                      - If Widget is not scrolled return 0.
 *                      - If scrolled return > 0.
 *                      - If scrolled inside (elastic scroll) return < 0.
 */
// llgo:link (*LvObjT).LvObjGetScrollX C.lv_obj_get_scroll_x
func (recv_ *LvObjT) LvObjGetScrollX() c.Int32T {
	return 0
}

/**
 * Get current Y scroll position.  Identical to `lv_obj_get_scroll_top()`.
 * @param obj       pointer to scrollable container Widget
 * @return          current scroll position from top edge
 *                      - If Widget is not scrolled return 0.
 *                      - If scrolled return > 0.
 *                      - If scrolled inside (elastic scroll) return < 0.
 */
// llgo:link (*LvObjT).LvObjGetScrollY C.lv_obj_get_scroll_y
func (recv_ *LvObjT) LvObjGetScrollY() c.Int32T {
	return 0
}

/**
 * Number of pixels a scrollable container Widget can be scrolled down
 * before its top edge appears.  When LV_OBJ_FLAG_SCROLL_ELASTIC flag
 * is set in Widget, this value can go negative while Widget is being
 * dragged below its normal top-edge boundary.
 * @param obj       pointer to scrollable container Widget
 * @return          pixels Widget can be scrolled down before its top edge appears
 */
// llgo:link (*LvObjT).LvObjGetScrollTop C.lv_obj_get_scroll_top
func (recv_ *LvObjT) LvObjGetScrollTop() c.Int32T {
	return 0
}

/**
 * Number of pixels a scrollable container Widget can be scrolled up
 * before its bottom edge appears.  When LV_OBJ_FLAG_SCROLL_ELASTIC flag
 * is set in Widget, this value can go negative while Widget is being
 * dragged above its normal bottom-edge boundary.
 * @param obj       pointer to scrollable container Widget
 * @return          pixels Widget can be scrolled up before its bottom edge appears
 */
// llgo:link (*LvObjT).LvObjGetScrollBottom C.lv_obj_get_scroll_bottom
func (recv_ *LvObjT) LvObjGetScrollBottom() c.Int32T {
	return 0
}

/**
 * Number of pixels a scrollable container Widget can be scrolled right
 * before its left edge appears.  When LV_OBJ_FLAG_SCROLL_ELASTIC flag
 * is set in Widget, this value can go negative while Widget is being
 * dragged farther right than its normal left-edge boundary.
 * @param obj       pointer to scrollable container Widget
 * @return          pixels Widget can be scrolled right before its left edge appears
 */
// llgo:link (*LvObjT).LvObjGetScrollLeft C.lv_obj_get_scroll_left
func (recv_ *LvObjT) LvObjGetScrollLeft() c.Int32T {
	return 0
}

/**
 * Number of pixels a scrollable container Widget can be scrolled left
 * before its right edge appears.  When LV_OBJ_FLAG_SCROLL_ELASTIC flag
 * is set in Widget, this value can go negative while Widget is being
 * dragged farther left than its normal right-edge boundary.
 * @param obj       pointer to scrollable container Widget
 * @return          pixels Widget can be scrolled left before its right edge appears
 */
// llgo:link (*LvObjT).LvObjGetScrollRight C.lv_obj_get_scroll_right
func (recv_ *LvObjT) LvObjGetScrollRight() c.Int32T {
	return 0
}

/**
 * Get the X and Y coordinates where the scrolling will end for Widget if a scrolling animation is in progress.
 * If no scrolling animation, give the current `x` or `y` scroll position.
 * @param obj       pointer to scrollable Widget
 * @param end       pointer to `lv_point_t` object in which to store result
 */
// llgo:link (*LvObjT).LvObjGetScrollEnd C.lv_obj_get_scroll_end
func (recv_ *LvObjT) LvObjGetScrollEnd(end *LvPointT) {
}

/**
 * Scroll by given amount of pixels.
 * @param obj       pointer to scrollable Widget to scroll
 * @param dx        pixels to scroll horizontally
 * @param dy        pixels to scroll vertically
 * @param anim_en   LV_ANIM_ON: scroll with animation; LV_ANIM_OFF: scroll immediately
 * @note            > 0 value means scroll right/bottom (show the more content on the right/bottom)
 * @note            e.g. dy = -20 means scroll down 20 px
 */
// llgo:link (*LvObjT).LvObjScrollBy C.lv_obj_scroll_by
func (recv_ *LvObjT) LvObjScrollBy(dx c.Int32T, dy c.Int32T, anim_en LvAnimEnableT) {
}

/**
 * Scroll by given amount of pixels.
 * `dx` and `dy` will be limited internally to allow scrolling only on the content area.
 * @param obj       pointer to scrollable Widget to scroll
 * @param dx        pixels to scroll horizontally
 * @param dy        pixels to scroll vertically
 * @param anim_en   LV_ANIM_ON: scroll with animation; LV_ANIM_OFF: scroll immediately
 * @note            e.g. dy = -20 means scroll down 20 px
 */
// llgo:link (*LvObjT).LvObjScrollByBounded C.lv_obj_scroll_by_bounded
func (recv_ *LvObjT) LvObjScrollByBounded(dx c.Int32T, dy c.Int32T, anim_en LvAnimEnableT) {
}

/**
 * Scroll to given coordinate on Widget.
 * `x` and `y` will be limited internally to allow scrolling only on the content area.
 * @param obj       pointer to scrollable Widget to scroll
 * @param x         pixels to scroll horizontally
 * @param y         pixels to scroll vertically
 * @param anim_en   LV_ANIM_ON: scroll with animation; LV_ANIM_OFF: scroll immediately
 */
// llgo:link (*LvObjT).LvObjScrollTo C.lv_obj_scroll_to
func (recv_ *LvObjT) LvObjScrollTo(x c.Int32T, y c.Int32T, anim_en LvAnimEnableT) {
}

/**
 * Scroll to X coordinate on Widget.
 * `x` will be limited internally to allow scrolling only on the content area.
 * @param obj       pointer to scrollable Widget to scroll
 * @param x         pixels to scroll horizontally
 * @param anim_en   LV_ANIM_ON: scroll with animation; LV_ANIM_OFF: scroll immediately
 */
// llgo:link (*LvObjT).LvObjScrollToX C.lv_obj_scroll_to_x
func (recv_ *LvObjT) LvObjScrollToX(x c.Int32T, anim_en LvAnimEnableT) {
}

/**
 * Scroll to Y coordinate on Widget.
 * `y` will be limited internally to allow scrolling only on the content area.
 * @param obj       pointer to scrollable Widget to scroll
 * @param y         pixels to scroll vertically
 * @param anim_en   LV_ANIM_ON: scroll with animation; LV_ANIM_OFF: scroll immediately
 */
// llgo:link (*LvObjT).LvObjScrollToY C.lv_obj_scroll_to_y
func (recv_ *LvObjT) LvObjScrollToY(y c.Int32T, anim_en LvAnimEnableT) {
}

/**
 * Scroll `obj`'s parent Widget until `obj` becomes visible.
 * @param obj       pointer to Widget to scroll into view
 * @param anim_en   LV_ANIM_ON: scroll with animation; LV_ANIM_OFF: scroll immediately
 */
// llgo:link (*LvObjT).LvObjScrollToView C.lv_obj_scroll_to_view
func (recv_ *LvObjT) LvObjScrollToView(anim_en LvAnimEnableT) {
}

/**
 * Scroll `obj`'s parent Widgets recursively until `obj` becomes visible.
 * Widget will be scrolled into view even it has nested scrollable parents.
 * @param obj       pointer to Widget to scroll into view
 * @param anim_en   LV_ANIM_ON: scroll with animation; LV_ANIM_OFF: scroll immediately
 */
// llgo:link (*LvObjT).LvObjScrollToViewRecursive C.lv_obj_scroll_to_view_recursive
func (recv_ *LvObjT) LvObjScrollToViewRecursive(anim_en LvAnimEnableT) {
}

/**
 * Tell whether Widget is being scrolled or not at this moment
 * @param obj   pointer to Widget
 * @return      true: `obj` is being scrolled
 */
// llgo:link (*LvObjT).LvObjIsScrolling C.lv_obj_is_scrolling
func (recv_ *LvObjT) LvObjIsScrolling() bool {
	return false
}

/**
 * Stop scrolling the current object
 *
 * @param obj The object being scrolled
 */
// llgo:link (*LvObjT).LvObjStopScrollAnim C.lv_obj_stop_scroll_anim
func (recv_ *LvObjT) LvObjStopScrollAnim() {
}

/**
 * Check children of `obj` and scroll `obj` to fulfill scroll_snap settings.
 * @param obj       Widget whose children need to be checked and snapped
 * @param anim_en   LV_ANIM_ON/OFF
 */
// llgo:link (*LvObjT).LvObjUpdateSnap C.lv_obj_update_snap
func (recv_ *LvObjT) LvObjUpdateSnap(anim_en LvAnimEnableT) {
}

/**
 * Get the area of the scrollbars
 * @param obj   pointer to Widget
 * @param hor   pointer to store the area of the horizontal scrollbar
 * @param ver   pointer to store the area of the vertical  scrollbar
 */
// llgo:link (*LvObjT).LvObjGetScrollbarArea C.lv_obj_get_scrollbar_area
func (recv_ *LvObjT) LvObjGetScrollbarArea(hor *LvAreaT, ver *LvAreaT) {
}

/**
 * Invalidate the area of the scrollbars
 * @param obj       pointer to Widget
 */
// llgo:link (*LvObjT).LvObjScrollbarInvalidate C.lv_obj_scrollbar_invalidate
func (recv_ *LvObjT) LvObjScrollbarInvalidate() {
}

/**
 * Checks if the content is scrolled "in" and adjusts it to a normal position.
 * @param obj       pointer to Widget
 * @param anim_en   LV_ANIM_ON/OFF
 */
// llgo:link (*LvObjT).LvObjReadjustScroll C.lv_obj_readjust_scroll
func (recv_ *LvObjT) LvObjReadjustScroll(anim_en LvAnimEnableT) {
}

type LvStyleStateCmpT c.Int

const (
	LV_STYLE_STATE_CMP_SAME          LvStyleStateCmpT = 0
	LV_STYLE_STATE_CMP_DIFF_REDRAW   LvStyleStateCmpT = 1
	LV_STYLE_STATE_CMP_DIFF_DRAW_PAD LvStyleStateCmpT = 2
	LV_STYLE_STATE_CMP_DIFF_LAYOUT   LvStyleStateCmpT = 3
)

type LvStyleSelectorT c.Uint32T

/**
 * Add a style to an object.
 * @param obj       pointer to an object
 * @param style     pointer to a style to add
 * @param selector  OR-ed value of parts and state to which the style should be added
 *
 * Examples:
 * @code
 * lv_obj_add_style(btn, &style_btn, 0); //Default button style
 *
 * lv_obj_add_style(btn, &btn_red, LV_STATE_PRESSED); //Overwrite only some colors to red when pressed
 * @endcode
 */
// llgo:link (*LvObjT).LvObjAddStyle C.lv_obj_add_style
func (recv_ *LvObjT) LvObjAddStyle(style *LvStyleT, selector LvStyleSelectorT) {
}

/**
 * Replaces a style of an object, preserving the order of the style stack (local styles and transitions are ignored).
 * @param obj           pointer to an object
 * @param old_style     pointer to a style to replace.
 * @param new_style     pointer to a style to replace the old style with.
 * @param selector      OR-ed values of states and a part to replace only styles with matching selectors. LV_STATE_ANY and LV_PART_ANY can be used
 *
 * Examples:
 * @code
 * lv_obj_replace_style(obj, &yellow_style, &blue_style, LV_PART_ANY | LV_STATE_ANY); //Replace a specific style
 *
 * lv_obj_replace_style(obj, &yellow_style, &blue_style, LV_PART_MAIN | LV_STATE_PRESSED); //Replace a specific style assigned to the main part when it is pressed
 * @endcode
 */
// llgo:link (*LvObjT).LvObjReplaceStyle C.lv_obj_replace_style
func (recv_ *LvObjT) LvObjReplaceStyle(old_style *LvStyleT, new_style *LvStyleT, selector LvStyleSelectorT) bool {
	return false
}

/**
 * Remove a style from an object.
 * @param obj       pointer to an object
 * @param style     pointer to a style to remove. Can be NULL to check only the selector
 * @param selector  OR-ed values of states and a part to remove only styles with matching selectors. LV_STATE_ANY and LV_PART_ANY can be used
 *
 * Examples:
 * @code
 * lv_obj_remove_style(obj, &style, LV_PART_ANY | LV_STATE_ANY); //Remove a specific style
 *
 * lv_obj_remove_style(obj, NULL, LV_PART_MAIN | LV_STATE_ANY); //Remove all styles from the main part
 *
 * lv_obj_remove_style(obj, NULL, LV_PART_ANY | LV_STATE_ANY); //Remove all styles
 * @endcode
 */
// llgo:link (*LvObjT).LvObjRemoveStyle C.lv_obj_remove_style
func (recv_ *LvObjT) LvObjRemoveStyle(style *LvStyleT, selector LvStyleSelectorT) {
}

/**
 * Remove all styles from an object
 * @param obj       pointer to an object
 */
// llgo:link (*LvObjT).LvObjRemoveStyleAll C.lv_obj_remove_style_all
func (recv_ *LvObjT) LvObjRemoveStyleAll() {
}

/**
 * Notify all object if a style is modified
 * @param style     pointer to a style. Only the objects with this style will be notified
 *                  (NULL to notify all objects)
 */
// llgo:link (*LvStyleT).LvObjReportStyleChange C.lv_obj_report_style_change
func (recv_ *LvStyleT) LvObjReportStyleChange() {
}

/**
 * Notify an object and its children about its style is modified.
 * @param obj       pointer to an object
 * @param part      the part whose style was changed. E.g. `LV_PART_ANY`, `LV_PART_MAIN`
 * @param prop      `LV_STYLE_PROP_ANY` or an `LV_STYLE_...` property.
 *                  It is used to optimize what needs to be refreshed.
 *                  `LV_STYLE_PROP_INV` to perform only a style cache update
 */
// llgo:link (*LvObjT).LvObjRefreshStyle C.lv_obj_refresh_style
func (recv_ *LvObjT) LvObjRefreshStyle(part LvPartT, prop LvStylePropT) {
}

/**
 * Temporary disable a style for a selector. It will look like is the style wasn't added
 * @param obj       pointer to an object
 * @param style     pointer to a style
 * @param selector  the selector of a style (e.g. LV_STATE_PRESSED | LV_PART_KNOB)
 * @param dis       true: disable the style, false: enable the style
 */
// llgo:link (*LvObjT).LvObjStyleSetDisabled C.lv_obj_style_set_disabled
func (recv_ *LvObjT) LvObjStyleSetDisabled(style *LvStyleT, selector LvStyleSelectorT, dis bool) {
}

/**
 * Get if a given style is disabled on an object.
 * @param obj       pointer to an object
 * @param style     pointer to a style
 * @param selector  the selector of a style (e.g. LV_STATE_PRESSED | LV_PART_KNOB)
 * @return          true: disable the style, false: enable the style
 */
// llgo:link (*LvObjT).LvObjStyleGetDisabled C.lv_obj_style_get_disabled
func (recv_ *LvObjT) LvObjStyleGetDisabled(style *LvStyleT, selector LvStyleSelectorT) bool {
	return false
}

/**
 * Enable or disable automatic style refreshing when a new style is added/removed to/from an object
 * or any other style change happens.
 * @param en        true: enable refreshing; false: disable refreshing
 */
//go:linkname LvObjEnableStyleRefresh C.lv_obj_enable_style_refresh
func LvObjEnableStyleRefresh(en bool)

/**
 * Get the value of a style property. The current state of the object will be considered.
 * Inherited properties will be inherited.
 * If a property is not set a default value will be returned.
 * @param obj       pointer to an object
 * @param part      a part from which the property should be get
 * @param prop      the property to get
 * @return          the value of the property.
 *                  Should be read from the correct field of the `lv_style_value_t` according to the type of the property.
 */
// llgo:link (*LvObjT).LvObjGetStyleProp C.lv_obj_get_style_prop
func (recv_ *LvObjT) LvObjGetStyleProp(part LvPartT, prop LvStylePropT) LvStyleValueT {
	return LvStyleValueT{}
}

/**
 * Check if an object has a specified style property for a given style selector.
 * @param obj       pointer to an object
 * @param selector  the style selector to be checked, defining the scope of the style to be examined.
 * @param prop      the property to be checked.
 * @return          true if the object has the specified selector and property, false otherwise.
 */
// llgo:link (*LvObjT).LvObjHasStyleProp C.lv_obj_has_style_prop
func (recv_ *LvObjT) LvObjHasStyleProp(selector LvStyleSelectorT, prop LvStylePropT) bool {
	return false
}

/**
 * Set local style property on an object's part and state.
 * @param obj       pointer to an object
 * @param prop      the property
 * @param value     value of the property. The correct element should be set according to the type of the property
 * @param selector  OR-ed value of parts and state for which the style should be set
 */
// llgo:link (*LvObjT).LvObjSetLocalStyleProp C.lv_obj_set_local_style_prop
func (recv_ *LvObjT) LvObjSetLocalStyleProp(prop LvStylePropT, value LvStyleValueT, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjGetLocalStyleProp C.lv_obj_get_local_style_prop
func (recv_ *LvObjT) LvObjGetLocalStyleProp(prop LvStylePropT, value *LvStyleValueT, selector LvStyleSelectorT) LvStyleResT {
	return 0
}

/**
 * Remove a local style property from a part of an object with a given state.
 * @param obj       pointer to an object
 * @param prop      a style property to remove.
 * @param selector  OR-ed value of parts and state for which the style should be removed
 * @return true     the property was found and removed; false: the property was not found
 */
// llgo:link (*LvObjT).LvObjRemoveLocalStyleProp C.lv_obj_remove_local_style_prop
func (recv_ *LvObjT) LvObjRemoveLocalStyleProp(prop LvStylePropT, selector LvStyleSelectorT) bool {
	return false
}

/**
 * Used internally for color filtering
 */
// llgo:link (*LvObjT).LvObjStyleApplyColorFilter C.lv_obj_style_apply_color_filter
func (recv_ *LvObjT) LvObjStyleApplyColorFilter(part LvPartT, v LvStyleValueT) LvStyleValueT {
	return LvStyleValueT{}
}

/**
 * Fade in an an object and all its children.
 * @param obj       the object to fade in
 * @param time      time of fade
 * @param delay     delay to start the animation
 */
// llgo:link (*LvObjT).LvObjFadeIn C.lv_obj_fade_in
func (recv_ *LvObjT) LvObjFadeIn(time c.Uint32T, delay c.Uint32T) {
}

/**
 * Fade out an an object and all its children.
 * @param obj       the object to fade out
 * @param time      time of fade
 * @param delay     delay to start the animation
 */
// llgo:link (*LvObjT).LvObjFadeOut C.lv_obj_fade_out
func (recv_ *LvObjT) LvObjFadeOut(time c.Uint32T, delay c.Uint32T) {
}

// llgo:link (*LvObjT).LvObjSetStyleWidth C.lv_obj_set_style_width
func (recv_ *LvObjT) LvObjSetStyleWidth(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleMinWidth C.lv_obj_set_style_min_width
func (recv_ *LvObjT) LvObjSetStyleMinWidth(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleMaxWidth C.lv_obj_set_style_max_width
func (recv_ *LvObjT) LvObjSetStyleMaxWidth(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleHeight C.lv_obj_set_style_height
func (recv_ *LvObjT) LvObjSetStyleHeight(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleMinHeight C.lv_obj_set_style_min_height
func (recv_ *LvObjT) LvObjSetStyleMinHeight(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleMaxHeight C.lv_obj_set_style_max_height
func (recv_ *LvObjT) LvObjSetStyleMaxHeight(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleLength C.lv_obj_set_style_length
func (recv_ *LvObjT) LvObjSetStyleLength(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleX C.lv_obj_set_style_x
func (recv_ *LvObjT) LvObjSetStyleX(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleY C.lv_obj_set_style_y
func (recv_ *LvObjT) LvObjSetStyleY(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleAlign C.lv_obj_set_style_align
func (recv_ *LvObjT) LvObjSetStyleAlign(value LvAlignT, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleTransformWidth C.lv_obj_set_style_transform_width
func (recv_ *LvObjT) LvObjSetStyleTransformWidth(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleTransformHeight C.lv_obj_set_style_transform_height
func (recv_ *LvObjT) LvObjSetStyleTransformHeight(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleTranslateX C.lv_obj_set_style_translate_x
func (recv_ *LvObjT) LvObjSetStyleTranslateX(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleTranslateY C.lv_obj_set_style_translate_y
func (recv_ *LvObjT) LvObjSetStyleTranslateY(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleTranslateRadial C.lv_obj_set_style_translate_radial
func (recv_ *LvObjT) LvObjSetStyleTranslateRadial(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleTransformScaleX C.lv_obj_set_style_transform_scale_x
func (recv_ *LvObjT) LvObjSetStyleTransformScaleX(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleTransformScaleY C.lv_obj_set_style_transform_scale_y
func (recv_ *LvObjT) LvObjSetStyleTransformScaleY(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleTransformRotation C.lv_obj_set_style_transform_rotation
func (recv_ *LvObjT) LvObjSetStyleTransformRotation(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleTransformPivotX C.lv_obj_set_style_transform_pivot_x
func (recv_ *LvObjT) LvObjSetStyleTransformPivotX(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleTransformPivotY C.lv_obj_set_style_transform_pivot_y
func (recv_ *LvObjT) LvObjSetStyleTransformPivotY(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleTransformSkewX C.lv_obj_set_style_transform_skew_x
func (recv_ *LvObjT) LvObjSetStyleTransformSkewX(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleTransformSkewY C.lv_obj_set_style_transform_skew_y
func (recv_ *LvObjT) LvObjSetStyleTransformSkewY(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStylePadTop C.lv_obj_set_style_pad_top
func (recv_ *LvObjT) LvObjSetStylePadTop(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStylePadBottom C.lv_obj_set_style_pad_bottom
func (recv_ *LvObjT) LvObjSetStylePadBottom(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStylePadLeft C.lv_obj_set_style_pad_left
func (recv_ *LvObjT) LvObjSetStylePadLeft(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStylePadRight C.lv_obj_set_style_pad_right
func (recv_ *LvObjT) LvObjSetStylePadRight(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStylePadRow C.lv_obj_set_style_pad_row
func (recv_ *LvObjT) LvObjSetStylePadRow(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStylePadColumn C.lv_obj_set_style_pad_column
func (recv_ *LvObjT) LvObjSetStylePadColumn(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStylePadRadial C.lv_obj_set_style_pad_radial
func (recv_ *LvObjT) LvObjSetStylePadRadial(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleMarginTop C.lv_obj_set_style_margin_top
func (recv_ *LvObjT) LvObjSetStyleMarginTop(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleMarginBottom C.lv_obj_set_style_margin_bottom
func (recv_ *LvObjT) LvObjSetStyleMarginBottom(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleMarginLeft C.lv_obj_set_style_margin_left
func (recv_ *LvObjT) LvObjSetStyleMarginLeft(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleMarginRight C.lv_obj_set_style_margin_right
func (recv_ *LvObjT) LvObjSetStyleMarginRight(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleBgColor C.lv_obj_set_style_bg_color
func (recv_ *LvObjT) LvObjSetStyleBgColor(value LvColorT, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleBgOpa C.lv_obj_set_style_bg_opa
func (recv_ *LvObjT) LvObjSetStyleBgOpa(value LvOpaT, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleBgGradColor C.lv_obj_set_style_bg_grad_color
func (recv_ *LvObjT) LvObjSetStyleBgGradColor(value LvColorT, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleBgGradDir C.lv_obj_set_style_bg_grad_dir
func (recv_ *LvObjT) LvObjSetStyleBgGradDir(value LvGradDirT, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleBgMainStop C.lv_obj_set_style_bg_main_stop
func (recv_ *LvObjT) LvObjSetStyleBgMainStop(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleBgGradStop C.lv_obj_set_style_bg_grad_stop
func (recv_ *LvObjT) LvObjSetStyleBgGradStop(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleBgMainOpa C.lv_obj_set_style_bg_main_opa
func (recv_ *LvObjT) LvObjSetStyleBgMainOpa(value LvOpaT, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleBgGradOpa C.lv_obj_set_style_bg_grad_opa
func (recv_ *LvObjT) LvObjSetStyleBgGradOpa(value LvOpaT, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleBgGrad C.lv_obj_set_style_bg_grad
func (recv_ *LvObjT) LvObjSetStyleBgGrad(value *LvGradDscT, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleBgImageSrc C.lv_obj_set_style_bg_image_src
func (recv_ *LvObjT) LvObjSetStyleBgImageSrc(value c.Pointer, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleBgImageOpa C.lv_obj_set_style_bg_image_opa
func (recv_ *LvObjT) LvObjSetStyleBgImageOpa(value LvOpaT, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleBgImageRecolor C.lv_obj_set_style_bg_image_recolor
func (recv_ *LvObjT) LvObjSetStyleBgImageRecolor(value LvColorT, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleBgImageRecolorOpa C.lv_obj_set_style_bg_image_recolor_opa
func (recv_ *LvObjT) LvObjSetStyleBgImageRecolorOpa(value LvOpaT, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleBgImageTiled C.lv_obj_set_style_bg_image_tiled
func (recv_ *LvObjT) LvObjSetStyleBgImageTiled(value bool, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleBorderColor C.lv_obj_set_style_border_color
func (recv_ *LvObjT) LvObjSetStyleBorderColor(value LvColorT, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleBorderOpa C.lv_obj_set_style_border_opa
func (recv_ *LvObjT) LvObjSetStyleBorderOpa(value LvOpaT, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleBorderWidth C.lv_obj_set_style_border_width
func (recv_ *LvObjT) LvObjSetStyleBorderWidth(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleBorderSide C.lv_obj_set_style_border_side
func (recv_ *LvObjT) LvObjSetStyleBorderSide(value LvBorderSideT, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleBorderPost C.lv_obj_set_style_border_post
func (recv_ *LvObjT) LvObjSetStyleBorderPost(value bool, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleOutlineWidth C.lv_obj_set_style_outline_width
func (recv_ *LvObjT) LvObjSetStyleOutlineWidth(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleOutlineColor C.lv_obj_set_style_outline_color
func (recv_ *LvObjT) LvObjSetStyleOutlineColor(value LvColorT, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleOutlineOpa C.lv_obj_set_style_outline_opa
func (recv_ *LvObjT) LvObjSetStyleOutlineOpa(value LvOpaT, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleOutlinePad C.lv_obj_set_style_outline_pad
func (recv_ *LvObjT) LvObjSetStyleOutlinePad(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleShadowWidth C.lv_obj_set_style_shadow_width
func (recv_ *LvObjT) LvObjSetStyleShadowWidth(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleShadowOffsetX C.lv_obj_set_style_shadow_offset_x
func (recv_ *LvObjT) LvObjSetStyleShadowOffsetX(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleShadowOffsetY C.lv_obj_set_style_shadow_offset_y
func (recv_ *LvObjT) LvObjSetStyleShadowOffsetY(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleShadowSpread C.lv_obj_set_style_shadow_spread
func (recv_ *LvObjT) LvObjSetStyleShadowSpread(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleShadowColor C.lv_obj_set_style_shadow_color
func (recv_ *LvObjT) LvObjSetStyleShadowColor(value LvColorT, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleShadowOpa C.lv_obj_set_style_shadow_opa
func (recv_ *LvObjT) LvObjSetStyleShadowOpa(value LvOpaT, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleImageOpa C.lv_obj_set_style_image_opa
func (recv_ *LvObjT) LvObjSetStyleImageOpa(value LvOpaT, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleImageRecolor C.lv_obj_set_style_image_recolor
func (recv_ *LvObjT) LvObjSetStyleImageRecolor(value LvColorT, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleImageRecolorOpa C.lv_obj_set_style_image_recolor_opa
func (recv_ *LvObjT) LvObjSetStyleImageRecolorOpa(value LvOpaT, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleLineWidth C.lv_obj_set_style_line_width
func (recv_ *LvObjT) LvObjSetStyleLineWidth(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleLineDashWidth C.lv_obj_set_style_line_dash_width
func (recv_ *LvObjT) LvObjSetStyleLineDashWidth(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleLineDashGap C.lv_obj_set_style_line_dash_gap
func (recv_ *LvObjT) LvObjSetStyleLineDashGap(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleLineRounded C.lv_obj_set_style_line_rounded
func (recv_ *LvObjT) LvObjSetStyleLineRounded(value bool, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleLineColor C.lv_obj_set_style_line_color
func (recv_ *LvObjT) LvObjSetStyleLineColor(value LvColorT, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleLineOpa C.lv_obj_set_style_line_opa
func (recv_ *LvObjT) LvObjSetStyleLineOpa(value LvOpaT, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleArcWidth C.lv_obj_set_style_arc_width
func (recv_ *LvObjT) LvObjSetStyleArcWidth(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleArcRounded C.lv_obj_set_style_arc_rounded
func (recv_ *LvObjT) LvObjSetStyleArcRounded(value bool, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleArcColor C.lv_obj_set_style_arc_color
func (recv_ *LvObjT) LvObjSetStyleArcColor(value LvColorT, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleArcOpa C.lv_obj_set_style_arc_opa
func (recv_ *LvObjT) LvObjSetStyleArcOpa(value LvOpaT, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleArcImageSrc C.lv_obj_set_style_arc_image_src
func (recv_ *LvObjT) LvObjSetStyleArcImageSrc(value c.Pointer, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleTextColor C.lv_obj_set_style_text_color
func (recv_ *LvObjT) LvObjSetStyleTextColor(value LvColorT, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleTextOpa C.lv_obj_set_style_text_opa
func (recv_ *LvObjT) LvObjSetStyleTextOpa(value LvOpaT, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleTextFont C.lv_obj_set_style_text_font
func (recv_ *LvObjT) LvObjSetStyleTextFont(value *LvFontT, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleTextLetterSpace C.lv_obj_set_style_text_letter_space
func (recv_ *LvObjT) LvObjSetStyleTextLetterSpace(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleTextLineSpace C.lv_obj_set_style_text_line_space
func (recv_ *LvObjT) LvObjSetStyleTextLineSpace(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleTextDecor C.lv_obj_set_style_text_decor
func (recv_ *LvObjT) LvObjSetStyleTextDecor(value LvTextDecorT, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleTextAlign C.lv_obj_set_style_text_align
func (recv_ *LvObjT) LvObjSetStyleTextAlign(value LvTextAlignT, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleTextOutlineStrokeColor C.lv_obj_set_style_text_outline_stroke_color
func (recv_ *LvObjT) LvObjSetStyleTextOutlineStrokeColor(value LvColorT, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleTextOutlineStrokeWidth C.lv_obj_set_style_text_outline_stroke_width
func (recv_ *LvObjT) LvObjSetStyleTextOutlineStrokeWidth(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleTextOutlineStrokeOpa C.lv_obj_set_style_text_outline_stroke_opa
func (recv_ *LvObjT) LvObjSetStyleTextOutlineStrokeOpa(value LvOpaT, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleRadius C.lv_obj_set_style_radius
func (recv_ *LvObjT) LvObjSetStyleRadius(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleRadialOffset C.lv_obj_set_style_radial_offset
func (recv_ *LvObjT) LvObjSetStyleRadialOffset(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleClipCorner C.lv_obj_set_style_clip_corner
func (recv_ *LvObjT) LvObjSetStyleClipCorner(value bool, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleOpa C.lv_obj_set_style_opa
func (recv_ *LvObjT) LvObjSetStyleOpa(value LvOpaT, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleOpaLayered C.lv_obj_set_style_opa_layered
func (recv_ *LvObjT) LvObjSetStyleOpaLayered(value LvOpaT, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleColorFilterDsc C.lv_obj_set_style_color_filter_dsc
func (recv_ *LvObjT) LvObjSetStyleColorFilterDsc(value *LvColorFilterDscT, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleColorFilterOpa C.lv_obj_set_style_color_filter_opa
func (recv_ *LvObjT) LvObjSetStyleColorFilterOpa(value LvOpaT, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleRecolor C.lv_obj_set_style_recolor
func (recv_ *LvObjT) LvObjSetStyleRecolor(value LvColorT, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleRecolorOpa C.lv_obj_set_style_recolor_opa
func (recv_ *LvObjT) LvObjSetStyleRecolorOpa(value LvOpaT, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleAnim C.lv_obj_set_style_anim
func (recv_ *LvObjT) LvObjSetStyleAnim(value *LvAnimT, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleAnimDuration C.lv_obj_set_style_anim_duration
func (recv_ *LvObjT) LvObjSetStyleAnimDuration(value c.Uint32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleTransition C.lv_obj_set_style_transition
func (recv_ *LvObjT) LvObjSetStyleTransition(value *LvStyleTransitionDscT, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleBlendMode C.lv_obj_set_style_blend_mode
func (recv_ *LvObjT) LvObjSetStyleBlendMode(value LvBlendModeT, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleLayout C.lv_obj_set_style_layout
func (recv_ *LvObjT) LvObjSetStyleLayout(value c.Uint16T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleBaseDir C.lv_obj_set_style_base_dir
func (recv_ *LvObjT) LvObjSetStyleBaseDir(value LvBaseDirT, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleBitmapMaskSrc C.lv_obj_set_style_bitmap_mask_src
func (recv_ *LvObjT) LvObjSetStyleBitmapMaskSrc(value c.Pointer, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleRotarySensitivity C.lv_obj_set_style_rotary_sensitivity
func (recv_ *LvObjT) LvObjSetStyleRotarySensitivity(value c.Uint32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleFlexFlow C.lv_obj_set_style_flex_flow
func (recv_ *LvObjT) LvObjSetStyleFlexFlow(value LvFlexFlowT, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleFlexMainPlace C.lv_obj_set_style_flex_main_place
func (recv_ *LvObjT) LvObjSetStyleFlexMainPlace(value LvFlexAlignT, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleFlexCrossPlace C.lv_obj_set_style_flex_cross_place
func (recv_ *LvObjT) LvObjSetStyleFlexCrossPlace(value LvFlexAlignT, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleFlexTrackPlace C.lv_obj_set_style_flex_track_place
func (recv_ *LvObjT) LvObjSetStyleFlexTrackPlace(value LvFlexAlignT, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleFlexGrow C.lv_obj_set_style_flex_grow
func (recv_ *LvObjT) LvObjSetStyleFlexGrow(value c.Uint8T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleGridColumnDscArray C.lv_obj_set_style_grid_column_dsc_array
func (recv_ *LvObjT) LvObjSetStyleGridColumnDscArray(value *c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleGridColumnAlign C.lv_obj_set_style_grid_column_align
func (recv_ *LvObjT) LvObjSetStyleGridColumnAlign(value LvGridAlignT, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleGridRowDscArray C.lv_obj_set_style_grid_row_dsc_array
func (recv_ *LvObjT) LvObjSetStyleGridRowDscArray(value *c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleGridRowAlign C.lv_obj_set_style_grid_row_align
func (recv_ *LvObjT) LvObjSetStyleGridRowAlign(value LvGridAlignT, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleGridCellColumnPos C.lv_obj_set_style_grid_cell_column_pos
func (recv_ *LvObjT) LvObjSetStyleGridCellColumnPos(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleGridCellXAlign C.lv_obj_set_style_grid_cell_x_align
func (recv_ *LvObjT) LvObjSetStyleGridCellXAlign(value LvGridAlignT, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleGridCellColumnSpan C.lv_obj_set_style_grid_cell_column_span
func (recv_ *LvObjT) LvObjSetStyleGridCellColumnSpan(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleGridCellRowPos C.lv_obj_set_style_grid_cell_row_pos
func (recv_ *LvObjT) LvObjSetStyleGridCellRowPos(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleGridCellYAlign C.lv_obj_set_style_grid_cell_y_align
func (recv_ *LvObjT) LvObjSetStyleGridCellYAlign(value LvGridAlignT, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjSetStyleGridCellRowSpan C.lv_obj_set_style_grid_cell_row_span
func (recv_ *LvObjT) LvObjSetStyleGridCellRowSpan(value c.Int32T, selector LvStyleSelectorT) {
}

// llgo:link (*LvObjT).LvObjCalculateStyleTextAlign C.lv_obj_calculate_style_text_align
func (recv_ *LvObjT) LvObjCalculateStyleTextAlign(part LvPartT, txt *c.Char) LvTextAlignT {
	return 0
}

/**
 * Get the `opa` style property from all parents and multiply and `>> 8` them.
 * @param obj       the object whose opacity should be get
 * @param part      the part whose opacity should be get. Non-MAIN parts will consider the `opa` of the MAIN part too
 * @return          the final opacity considering the parents' opacity too
 */
// llgo:link (*LvObjT).LvObjGetStyleOpaRecursive C.lv_obj_get_style_opa_recursive
func (recv_ *LvObjT) LvObjGetStyleOpaRecursive(part LvPartT) LvOpaT {
	return 0
}

/**
 * Apply recolor effect to the input color based on the object's style properties.
 * @param obj       the target object containing recolor style properties
 * @param part      the part to retrieve recolor styles.
 * @param color     the original color to be modified
 * @return          the blended color after applying recolor and opacity
 */
// llgo:link (*LvObjT).LvObjStyleApplyRecolor C.lv_obj_style_apply_recolor
func (recv_ *LvObjT) LvObjStyleApplyRecolor(part LvPartT, color LvColor32T) LvColor32T {
	return LvColor32T{}
}

/**
 * Get the `recolor` style property from all parents and blend them recursively.
 * @param obj       the object whose recolor value should be retrieved
 * @param part      the target part to check. Non-MAIN parts will also consider
 *                  the `recolor` value from the MAIN part during calculation
 * @return          the final blended recolor value combining all parent's recolor values
 */
// llgo:link (*LvObjT).LvObjGetStyleRecolorRecursive C.lv_obj_get_style_recolor_recursive
func (recv_ *LvObjT) LvObjGetStyleRecolorRecursive(part LvPartT) LvColor32T {
	return LvColor32T{}
}

type LvLayerTypeT c.Int

const (
	LV_LAYER_TYPE_NONE      LvLayerTypeT = 0
	LV_LAYER_TYPE_SIMPLE    LvLayerTypeT = 1
	LV_LAYER_TYPE_TRANSFORM LvLayerTypeT = 2
)

/**
 * Initialize a rectangle draw descriptor from an object's styles in its current state
 * @param obj       pointer to an object
 * @param part      part of the object, e.g. `LV_PART_MAIN`, `LV_PART_SCROLLBAR`, `LV_PART_KNOB`, etc
 * @param draw_dsc  the descriptor to initialize.
 *                  If an `..._opa` field is set to `LV_OPA_TRANSP` the related properties won't be initialized.
 *                  Should be initialized with `lv_draw_rect_dsc_init(draw_dsc)`.
 * @note Only the relevant fields will be set.
 *       E.g. if `border width == 0` the other border properties won't be evaluated.
 */
// llgo:link (*LvObjT).LvObjInitDrawRectDsc C.lv_obj_init_draw_rect_dsc
func (recv_ *LvObjT) LvObjInitDrawRectDsc(part LvPartT, draw_dsc *LvDrawRectDscT) {
}

/**
 * Initialize a label draw descriptor from an object's styles in its current state
 * @param obj       pointer to an object
 * @param part      part of the object, e.g. `LV_PART_MAIN`, `LV_PART_SCROLLBAR`, `LV_PART_KNOB`, etc
 * @param draw_dsc  the descriptor to initialize.
 *                  If the `opa` field is set to or the property is equal to `LV_OPA_TRANSP` the rest won't be initialized.
 *                  Should be initialized with `lv_draw_label_dsc_init(draw_dsc)`.
 */
// llgo:link (*LvObjT).LvObjInitDrawLabelDsc C.lv_obj_init_draw_label_dsc
func (recv_ *LvObjT) LvObjInitDrawLabelDsc(part LvPartT, draw_dsc *LvDrawLabelDscT) {
}

/**
 * Initialize an image draw descriptor from an object's styles in its current state
 * @param obj       pointer to an object
 * @param part      part of the object, e.g. `LV_PART_MAIN`, `LV_PART_SCROLLBAR`, `LV_PART_KNOB`, etc
 * @param draw_dsc  the descriptor to initialize.
 *                  Should be initialized with `lv_draw_image_dsc_init(draw_dsc)`.
 */
// llgo:link (*LvObjT).LvObjInitDrawImageDsc C.lv_obj_init_draw_image_dsc
func (recv_ *LvObjT) LvObjInitDrawImageDsc(part LvPartT, draw_dsc *LvDrawImageDscT) {
}

/**
 * Initialize a line draw descriptor from an object's styles in its current state
 * @param obj pointer to an object
 * @param part      part of the object, e.g. `LV_PART_MAIN`, `LV_PART_SCROLLBAR`, `LV_PART_KNOB`, etc
 * @param draw_dsc  the descriptor to initialize.
 *                  Should be initialized with `lv_draw_line_dsc_init(draw_dsc)`.
 */
// llgo:link (*LvObjT).LvObjInitDrawLineDsc C.lv_obj_init_draw_line_dsc
func (recv_ *LvObjT) LvObjInitDrawLineDsc(part LvPartT, draw_dsc *LvDrawLineDscT) {
}

/**
 * Initialize an arc draw descriptor from an object's styles in its current state
 * @param obj       pointer to an object
 * @param part      part of the object, e.g. `LV_PART_MAIN`, `LV_PART_SCROLLBAR`, `LV_PART_KNOB`, etc
 * @param draw_dsc  the descriptor to initialize.
 *                  Should be initialized with `lv_draw_arc_dsc_init(draw_dsc)`.
 */
// llgo:link (*LvObjT).LvObjInitDrawArcDsc C.lv_obj_init_draw_arc_dsc
func (recv_ *LvObjT) LvObjInitDrawArcDsc(part LvPartT, draw_dsc *LvDrawArcDscT) {
}

/**
 * Get the required extra size (around the object's part) to draw shadow, outline, value etc.
 * @param obj       pointer to an object
 * @param part      part of the object
 * @return          the extra size required around the object
 */
// llgo:link (*LvObjT).LvObjCalculateExtDrawSize C.lv_obj_calculate_ext_draw_size
func (recv_ *LvObjT) LvObjCalculateExtDrawSize(part LvPartT) c.Int32T {
	return 0
}

/**
 * Send a 'LV_EVENT_REFR_EXT_DRAW_SIZE' Call the ancestor's event handler to the object to refresh the value of the extended draw size.
 * The result will be saved in `obj`.
 * @param obj       pointer to an object
 */
// llgo:link (*LvObjT).LvObjRefreshExtDrawSize C.lv_obj_refresh_ext_draw_size
func (recv_ *LvObjT) LvObjRefreshExtDrawSize() {
}

type LvObjClassEditableT c.Int

const (
	LV_OBJ_CLASS_EDITABLE_INHERIT LvObjClassEditableT = 0
	LV_OBJ_CLASS_EDITABLE_TRUE    LvObjClassEditableT = 1
	LV_OBJ_CLASS_EDITABLE_FALSE   LvObjClassEditableT = 2
)

type LvObjClassGroupDefT c.Int

const (
	LV_OBJ_CLASS_GROUP_DEF_INHERIT LvObjClassGroupDefT = 0
	LV_OBJ_CLASS_GROUP_DEF_TRUE    LvObjClassGroupDefT = 1
	LV_OBJ_CLASS_GROUP_DEF_FALSE   LvObjClassGroupDefT = 2
)

type LvObjClassThemeInheritableT c.Int

const (
	LV_OBJ_CLASS_THEME_INHERITABLE_FALSE LvObjClassThemeInheritableT = 0
	LV_OBJ_CLASS_THEME_INHERITABLE_TRUE  LvObjClassThemeInheritableT = 1
)

// llgo:type C
type LvObjClassEventCbT func(*LvObjClassT, *LvEventT)

/**
 * Create an object form a class descriptor
 * @param class_p   pointer to a class
 * @param parent    pointer to an object where the new object should be created
 * @return          pointer to the created object
 */
// llgo:link (*LvObjClassT).LvObjClassCreateObj C.lv_obj_class_create_obj
func (recv_ *LvObjClassT) LvObjClassCreateObj(parent *LvObjT) *LvObjT {
	return nil
}

// llgo:link (*LvObjT).LvObjClassInitObj C.lv_obj_class_init_obj
func (recv_ *LvObjT) LvObjClassInitObj() {
}

// llgo:link (*LvObjT).LvObjIsEditable C.lv_obj_is_editable
func (recv_ *LvObjT) LvObjIsEditable() bool {
	return false
}

// llgo:link (*LvObjT).LvObjIsGroupDef C.lv_obj_is_group_def
func (recv_ *LvObjT) LvObjIsGroupDef() bool {
	return false
}

type LvKeyT c.Int

const (
	LV_KEY_UP        LvKeyT = 17
	LV_KEY_DOWN      LvKeyT = 18
	LV_KEY_RIGHT     LvKeyT = 19
	LV_KEY_LEFT      LvKeyT = 20
	LV_KEY_ESC       LvKeyT = 27
	LV_KEY_DEL       LvKeyT = 127
	LV_KEY_BACKSPACE LvKeyT = 8
	LV_KEY_ENTER     LvKeyT = 10
	LV_KEY_NEXT      LvKeyT = 9
	LV_KEY_PREV      LvKeyT = 11
	LV_KEY_HOME      LvKeyT = 2
	LV_KEY_END       LvKeyT = 3
)

// llgo:type C
type LvGroupFocusCbT func(*LvGroupT)

// llgo:type C
type LvGroupEdgeCbT func(*LvGroupT, bool)
type LvGroupRefocusPolicyT c.Int

const (
	LV_GROUP_REFOCUS_POLICY_NEXT LvGroupRefocusPolicyT = 0
	LV_GROUP_REFOCUS_POLICY_PREV LvGroupRefocusPolicyT = 1
)

/**
 * Create new Widget group.
 * @return          pointer to the new Widget group
 */
//go:linkname LvGroupCreate C.lv_group_create
func LvGroupCreate() *LvGroupT

/**
 * Delete group object.
 * @param group     pointer to a group
 */
// llgo:link (*LvGroupT).LvGroupDelete C.lv_group_delete
func (recv_ *LvGroupT) LvGroupDelete() {
}

/**
 * Set default group. New Widgets will be added to this group if it's enabled in
 * their class with `add_to_def_group = true`.
 * @param group     pointer to a group (can be `NULL`)
 */
// llgo:link (*LvGroupT).LvGroupSetDefault C.lv_group_set_default
func (recv_ *LvGroupT) LvGroupSetDefault() {
}

/**
 * Get default group.
 * @return          pointer to the default group
 */
//go:linkname LvGroupGetDefault C.lv_group_get_default
func LvGroupGetDefault() *LvGroupT

/**
 * Add an Widget to group.
 * @param group     pointer to a group
 * @param obj       pointer to a Widget to add
 */
// llgo:link (*LvGroupT).LvGroupAddObj C.lv_group_add_obj
func (recv_ *LvGroupT) LvGroupAddObj(obj *LvObjT) {
}

/**
 * Swap 2 Widgets in group.  Widgets must be in the same group.
 * @param obj1  pointer to a Widget
 * @param obj2  pointer to another Widget
 */
// llgo:link (*LvObjT).LvGroupSwapObj C.lv_group_swap_obj
func (recv_ *LvObjT) LvGroupSwapObj(obj2 *LvObjT) {
}

/**
 * Remove a Widget from its group.
 * @param obj       pointer to Widget to remove
 */
// llgo:link (*LvObjT).LvGroupRemoveObj C.lv_group_remove_obj
func (recv_ *LvObjT) LvGroupRemoveObj() {
}

/**
 * Remove all Widgets from a group.
 * @param group     pointer to a group
 */
// llgo:link (*LvGroupT).LvGroupRemoveAllObjs C.lv_group_remove_all_objs
func (recv_ *LvGroupT) LvGroupRemoveAllObjs() {
}

/**
 * Focus on a Widget (defocus the current).
 * @param obj       pointer to Widget to focus on
 */
// llgo:link (*LvObjT).LvGroupFocusObj C.lv_group_focus_obj
func (recv_ *LvObjT) LvGroupFocusObj() {
}

/**
 * Focus on next Widget in a group (defocus the current).
 * @param group     pointer to a group
 */
// llgo:link (*LvGroupT).LvGroupFocusNext C.lv_group_focus_next
func (recv_ *LvGroupT) LvGroupFocusNext() {
}

/**
 * Focus on previous Widget in a group (defocus the current).
 * @param group     pointer to a group
 */
// llgo:link (*LvGroupT).LvGroupFocusPrev C.lv_group_focus_prev
func (recv_ *LvGroupT) LvGroupFocusPrev() {
}

/**
 * Do not allow changing focus from current Widget.
 * @param group     pointer to a group
 * @param en        true: freeze, false: release freezing (normal mode)
 */
// llgo:link (*LvGroupT).LvGroupFocusFreeze C.lv_group_focus_freeze
func (recv_ *LvGroupT) LvGroupFocusFreeze(en bool) {
}

/**
 * Send a control character to Widget that has focus in a group.
 * @param group     pointer to a group
 * @param c         a character (use LV_KEY_.. to navigate)
 * @return          result of Widget with focus in group.
 */
// llgo:link (*LvGroupT).LvGroupSendData C.lv_group_send_data
func (recv_ *LvGroupT) LvGroupSendData(c c.Uint32T) LvResultT {
	return 0
}

/**
 * Set a function for a group which will be called when a new Widget has focus.
 * @param group         pointer to a group
 * @param focus_cb      the call back function or NULL if unused
 */
// llgo:link (*LvGroupT).LvGroupSetFocusCb C.lv_group_set_focus_cb
func (recv_ *LvGroupT) LvGroupSetFocusCb(focus_cb LvGroupFocusCbT) {
}

/**
 * Set a function for a group which will be called when a focus edge is reached
 * @param group         pointer to a group
 * @param edge_cb      the call back function or NULL if unused
 */
// llgo:link (*LvGroupT).LvGroupSetEdgeCb C.lv_group_set_edge_cb
func (recv_ *LvGroupT) LvGroupSetEdgeCb(edge_cb LvGroupEdgeCbT) {
}

/**
 * Set whether the next or previous Widget in a group gets focus when Widget that has
 * focus is deleted.
 * @param group         pointer to a group
 * @param policy        new refocus policy enum
 */
// llgo:link (*LvGroupT).LvGroupSetRefocusPolicy C.lv_group_set_refocus_policy
func (recv_ *LvGroupT) LvGroupSetRefocusPolicy(policy LvGroupRefocusPolicyT) {
}

/**
 * Manually set the current mode (edit or navigate).
 * @param group         pointer to group
 * @param edit          true: edit mode; false: navigate mode
 */
// llgo:link (*LvGroupT).LvGroupSetEditing C.lv_group_set_editing
func (recv_ *LvGroupT) LvGroupSetEditing(edit bool) {
}

/**
 * Set whether moving focus to next/previous Widget will allow wrapping from
 * first->last or last->first Widget.
 * @param group         pointer to group
 * @param               en true: wrapping enabled; false: wrapping disabled
 */
// llgo:link (*LvGroupT).LvGroupSetWrap C.lv_group_set_wrap
func (recv_ *LvGroupT) LvGroupSetWrap(en bool) {
}

/**
 * Get Widget that has focus, or NULL if there isn't one.
 * @param group         pointer to a group
 * @return              pointer to Widget with focus
 */
// llgo:link (*LvGroupT).LvGroupGetFocused C.lv_group_get_focused
func (recv_ *LvGroupT) LvGroupGetFocused() *LvObjT {
	return nil
}

/**
 * Get focus callback function of a group.
 * @param group pointer to a group
 * @return the call back function or NULL if not set
 */
// llgo:link (*LvGroupT).LvGroupGetFocusCb C.lv_group_get_focus_cb
func (recv_ *LvGroupT) LvGroupGetFocusCb() LvGroupFocusCbT {
	return nil
}

/**
 * Get edge callback function of a group.
 * @param group pointer to a group
 * @return the call back function or NULL if not set
 */
// llgo:link (*LvGroupT).LvGroupGetEdgeCb C.lv_group_get_edge_cb
func (recv_ *LvGroupT) LvGroupGetEdgeCb() LvGroupEdgeCbT {
	return nil
}

/**
 * Get current mode (edit or navigate).
 * @param group         pointer to group
 * @return              true: edit mode; false: navigate mode
 */
// llgo:link (*LvGroupT).LvGroupGetEditing C.lv_group_get_editing
func (recv_ *LvGroupT) LvGroupGetEditing() bool {
	return false
}

/**
 * Get whether moving focus to next/previous Widget will allow wrapping from
 * first->last or last->first Widget.
 * @param group         pointer to group
 */
// llgo:link (*LvGroupT).LvGroupGetWrap C.lv_group_get_wrap
func (recv_ *LvGroupT) LvGroupGetWrap() bool {
	return false
}

/**
 * Get number of Widgets in group.
 * @param group         pointer to a group
 * @return              number of Widgets in the group
 */
// llgo:link (*LvGroupT).LvGroupGetObjCount C.lv_group_get_obj_count
func (recv_ *LvGroupT) LvGroupGetObjCount() c.Uint32T {
	return 0
}

/**
 * Get nth Widget within group.
 * @param group         pointer to a group
 * @param index         index of Widget within the group
 * @return              pointer to Widget
 */
// llgo:link (*LvGroupT).LvGroupGetObjByIndex C.lv_group_get_obj_by_index
func (recv_ *LvGroupT) LvGroupGetObjByIndex(index c.Uint32T) *LvObjT {
	return nil
}

/**
 * Get the number of groups.
 * @return              number of groups
 */
//go:linkname LvGroupGetCount C.lv_group_get_count
func LvGroupGetCount() c.Uint32T

/**
 * Get a group by its index.
 * @param index         index of the group
 * @return              pointer to the group
 */
//go:linkname LvGroupByIndex C.lv_group_by_index
func LvGroupByIndex(index c.Uint32T) *LvGroupT

type LvIndevTypeT c.Int

const (
	LV_INDEV_TYPE_NONE    LvIndevTypeT = 0
	LV_INDEV_TYPE_POINTER LvIndevTypeT = 1
	LV_INDEV_TYPE_KEYPAD  LvIndevTypeT = 2
	LV_INDEV_TYPE_BUTTON  LvIndevTypeT = 3
	LV_INDEV_TYPE_ENCODER LvIndevTypeT = 4
)

type LvIndevStateT c.Int

const (
	LV_INDEV_STATE_RELEASED LvIndevStateT = 0
	LV_INDEV_STATE_PRESSED  LvIndevStateT = 1
)

type LvIndevModeT c.Int

const (
	LV_INDEV_MODE_NONE  LvIndevModeT = 0
	LV_INDEV_MODE_TIMER LvIndevModeT = 1
	LV_INDEV_MODE_EVENT LvIndevModeT = 2
)

type LvIndevGestureTypeT c.Int

const (
	LV_INDEV_GESTURE_NONE              LvIndevGestureTypeT = 0
	LV_INDEV_GESTURE_PINCH             LvIndevGestureTypeT = 1
	LV_INDEV_GESTURE_SWIPE             LvIndevGestureTypeT = 2
	LV_INDEV_GESTURE_ROTATE            LvIndevGestureTypeT = 3
	LV_INDEV_GESTURE_TWO_FINGERS_SWIPE LvIndevGestureTypeT = 4
	LV_INDEV_GESTURE_SCROLL            LvIndevGestureTypeT = 5
	LV_INDEV_GESTURE_CNT               LvIndevGestureTypeT = 6
)

/** Data structure passed to an input driver to fill*/

type LvIndevDataT struct {
	GestureType     [6]LvIndevGestureTypeT
	GestureData     [6]c.Pointer
	State           LvIndevStateT
	Point           LvPointT
	Key             c.Uint32T
	BtnId           c.Uint32T
	EncDiff         c.Int16T
	Timestamp       c.Uint32T
	ContinueReading bool
}

// llgo:type C
type LvIndevReadCbT func(*LvIndevT, *LvIndevDataT)

/**
 * Create an indev
 * @return Pointer to the created indev or NULL when allocation failed
 */
//go:linkname LvIndevCreate C.lv_indev_create
func LvIndevCreate() *LvIndevT

/**
 * Remove the provided input device. Make sure not to use the provided input device afterwards anymore.
 * @param indev pointer to delete
 */
// llgo:link (*LvIndevT).LvIndevDelete C.lv_indev_delete
func (recv_ *LvIndevT) LvIndevDelete() {
}

/**
 * Get the next input device.
 * @param indev pointer to the current input device. NULL to initialize.
 * @return the next input device or NULL if there are no more. Provide the first input device when
 * the parameter is NULL
 */
// llgo:link (*LvIndevT).LvIndevGetNext C.lv_indev_get_next
func (recv_ *LvIndevT) LvIndevGetNext() *LvIndevT {
	return nil
}

/**
 * Read data from an input device.
 * @param indev pointer to an input device
 */
// llgo:link (*LvIndevT).LvIndevRead C.lv_indev_read
func (recv_ *LvIndevT) LvIndevRead() {
}

/**
 * Called periodically to read the input devices
 * @param timer pointer to a timer to read
 */
// llgo:link (*LvTimerT).LvIndevReadTimerCb C.lv_indev_read_timer_cb
func (recv_ *LvTimerT) LvIndevReadTimerCb() {
}

/**
 * Enable or disable one or all input devices (default enabled)
 * @param indev pointer to an input device or NULL to enable/disable all of them
 * @param enable true to enable, false to disable
 */
// llgo:link (*LvIndevT).LvIndevEnable C.lv_indev_enable
func (recv_ *LvIndevT) LvIndevEnable(enable bool) {
}

/**
 * Get the currently processed input device. Can be used in action functions too.
 * @return pointer to the currently processed input device or NULL if no input device processing
 * right now
 */
//go:linkname LvIndevActive C.lv_indev_active
func LvIndevActive() *LvIndevT

/**
 * Set the type of an input device
 * @param indev pointer to an input device
 * @param indev_type the type of the input device from `lv_indev_type_t` (`LV_INDEV_TYPE_...`)
 */
// llgo:link (*LvIndevT).LvIndevSetType C.lv_indev_set_type
func (recv_ *LvIndevT) LvIndevSetType(indev_type LvIndevTypeT) {
}

/**
 * Set a callback function to read input device data to the indev
 * @param indev pointer to an input device
 * @param read_cb pointer to callback function to read input device data
 */
// llgo:link (*LvIndevT).LvIndevSetReadCb C.lv_indev_set_read_cb
func (recv_ *LvIndevT) LvIndevSetReadCb(read_cb LvIndevReadCbT) {
}

/**
 * Set user data to the indev
 * @param indev pointer to an input device
 * @param user_data pointer to user data
 */
// llgo:link (*LvIndevT).LvIndevSetUserData C.lv_indev_set_user_data
func (recv_ *LvIndevT) LvIndevSetUserData(user_data c.Pointer) {
}

/**
 * Set driver data to the indev
 * @param indev pointer to an input device
 * @param driver_data pointer to driver data
 */
// llgo:link (*LvIndevT).LvIndevSetDriverData C.lv_indev_set_driver_data
func (recv_ *LvIndevT) LvIndevSetDriverData(driver_data c.Pointer) {
}

/**
 * Assign a display to the indev
 * @param indev pointer to an input device
 * @param disp pointer to an display
 */
// llgo:link (*LvIndevT).LvIndevSetDisplay C.lv_indev_set_display
func (recv_ *LvIndevT) LvIndevSetDisplay(disp *X_lvDisplayT) {
}

/**
 * Set long press time to indev
 * @param  indev            pointer to input device
 * @param  long_press_time  time long press time in ms
 */
// llgo:link (*LvIndevT).LvIndevSetLongPressTime C.lv_indev_set_long_press_time
func (recv_ *LvIndevT) LvIndevSetLongPressTime(long_press_time c.Uint16T) {
}

/**
 * Set long press repeat time to indev
 * @param  indev            pointer to input device
 * @param  long_press_repeat_time  long press repeat time in ms
 */
// llgo:link (*LvIndevT).LvIndevSetLongPressRepeatTime C.lv_indev_set_long_press_repeat_time
func (recv_ *LvIndevT) LvIndevSetLongPressRepeatTime(long_press_repeat_time c.Uint16T) {
}

/**
 * Set scroll limit to the input device
 * @param indev pointer to an input device
 * @param scroll_limit the number of pixels to slide before actually drag the object
 */
// llgo:link (*LvIndevT).LvIndevSetScrollLimit C.lv_indev_set_scroll_limit
func (recv_ *LvIndevT) LvIndevSetScrollLimit(scroll_limit c.Uint8T) {
}

/**
 * Set scroll throw slow-down to the indev. Greater value means faster slow-down
 * @param indev pointer to an input device
 * @param scroll_throw the slow-down in [%]
 */
// llgo:link (*LvIndevT).LvIndevSetScrollThrow C.lv_indev_set_scroll_throw
func (recv_ *LvIndevT) LvIndevSetScrollThrow(scroll_throw c.Uint8T) {
}

/**
 * Get the type of an input device
 * @param indev pointer to an input device
 * @return the type of the input device from `lv_hal_indev_type_t` (`LV_INDEV_TYPE_...`)
 */
// llgo:link (*LvIndevT).LvIndevGetType C.lv_indev_get_type
func (recv_ *LvIndevT) LvIndevGetType() LvIndevTypeT {
	return 0
}

/**
 * Get the callback function to read input device data to the indev
 * @param indev pointer to an input device
 * @return Pointer to callback function to read input device data or NULL if indev is NULL
 */
// llgo:link (*LvIndevT).LvIndevGetReadCb C.lv_indev_get_read_cb
func (recv_ *LvIndevT) LvIndevGetReadCb() LvIndevReadCbT {
	return nil
}

/**
 * Get the indev state
 * @param indev pointer to an input device
 * @return Indev state or LV_INDEV_STATE_RELEASED if indev is NULL
 */
// llgo:link (*LvIndevT).LvIndevGetState C.lv_indev_get_state
func (recv_ *LvIndevT) LvIndevGetState() LvIndevStateT {
	return 0
}

/**
 * Get the indev assigned group
 * @param indev pointer to an input device
 * @return Pointer to indev assigned group or NULL if indev is NULL
 */
// llgo:link (*LvIndevT).LvIndevGetGroup C.lv_indev_get_group
func (recv_ *LvIndevT) LvIndevGetGroup() *LvGroupT {
	return nil
}

/**
 * Get a pointer to the assigned display of the indev
 * @param indev pointer to an input device
 * @return pointer to the assigned display or NULL if indev is NULL
 */
// llgo:link (*LvIndevT).LvIndevGetDisplay C.lv_indev_get_display
func (recv_ *LvIndevT) LvIndevGetDisplay() *LvDisplayT {
	return nil
}

/**
 * Get a pointer to the user data of the indev
 * @param indev pointer to an input device
 * @return pointer to the user data or NULL if indev is NULL
 */
// llgo:link (*LvIndevT).LvIndevGetUserData C.lv_indev_get_user_data
func (recv_ *LvIndevT) LvIndevGetUserData() c.Pointer {
	return nil
}

/**
 * Get a pointer to the driver data of the indev
 * @param indev pointer to an input device
 * @return pointer to the driver data or NULL if indev is NULL
 */
// llgo:link (*LvIndevT).LvIndevGetDriverData C.lv_indev_get_driver_data
func (recv_ *LvIndevT) LvIndevGetDriverData() c.Pointer {
	return nil
}

/**
 * Get whether indev is moved while pressed
 * @param indev pointer to an input device
 * @return true: indev is moved while pressed; false: indev is not moved while pressed
 */
// llgo:link (*LvIndevT).LvIndevGetPressMoved C.lv_indev_get_press_moved
func (recv_ *LvIndevT) LvIndevGetPressMoved() bool {
	return false
}

/**
 * Reset one or all input devices
 * @param indev pointer to an input device to reset or NULL to reset all of them
 * @param obj pointer to an object which triggers the reset.
 */
// llgo:link (*LvIndevT).LvIndevReset C.lv_indev_reset
func (recv_ *LvIndevT) LvIndevReset(obj *LvObjT) {
}

/**
 * Touch and key related events are sent to the input device first and to the widget after that.
 * If this functions called in an indev event, the event won't be sent to the widget.
 * @param indev pointer to an input device
 */
// llgo:link (*LvIndevT).LvIndevStopProcessing C.lv_indev_stop_processing
func (recv_ *LvIndevT) LvIndevStopProcessing() {
}

/**
 * Reset the long press state of an input device
 * @param indev pointer to an input device
 */
// llgo:link (*LvIndevT).LvIndevResetLongPress C.lv_indev_reset_long_press
func (recv_ *LvIndevT) LvIndevResetLongPress() {
}

/**
 * Set a cursor for a pointer input device (for LV_INPUT_TYPE_POINTER and LV_INPUT_TYPE_BUTTON)
 * @param indev pointer to an input device
 * @param cur_obj pointer to an object to be used as cursor
 */
// llgo:link (*LvIndevT).LvIndevSetCursor C.lv_indev_set_cursor
func (recv_ *LvIndevT) LvIndevSetCursor(cur_obj *LvObjT) {
}

/**
 * Set a destination group for a keypad input device (for LV_INDEV_TYPE_KEYPAD)
 * @param indev pointer to an input device
 * @param group pointer to a group
 */
// llgo:link (*LvIndevT).LvIndevSetGroup C.lv_indev_set_group
func (recv_ *LvIndevT) LvIndevSetGroup(group *LvGroupT) {
}

/**
 * Set the an array of points for LV_INDEV_TYPE_BUTTON.
 * These points will be assigned to the buttons to press a specific point on the screen
 * @param indev pointer to an input device
 * @param points array of points
 */
// llgo:link (*LvIndevT).LvIndevSetButtonPoints C.lv_indev_set_button_points
func (recv_ *LvIndevT) LvIndevSetButtonPoints(points *LvPointT) {
}

/**
 * Get the last point of an input device (for LV_INDEV_TYPE_POINTER and LV_INDEV_TYPE_BUTTON)
 * @param indev pointer to an input device
 * @param point pointer to a point to store the result
 */
// llgo:link (*LvIndevT).LvIndevGetPoint C.lv_indev_get_point
func (recv_ *LvIndevT) LvIndevGetPoint(point *LvPointT) {
}

/**
* Get the current gesture direct
* @param indev pointer to an input device
* @return current gesture direct
 */
// llgo:link (*LvIndevT).LvIndevGetGestureDir C.lv_indev_get_gesture_dir
func (recv_ *LvIndevT) LvIndevGetGestureDir() LvDirT {
	return 0
}

/**
 * Get the last pressed key of an input device (for LV_INDEV_TYPE_KEYPAD)
 * @param indev pointer to an input device
 * @return the last pressed key (0 on error)
 */
// llgo:link (*LvIndevT).LvIndevGetKey C.lv_indev_get_key
func (recv_ *LvIndevT) LvIndevGetKey() c.Uint32T {
	return 0
}

/**
 * Get the counter for consecutive clicks within a short distance and time.
 * The counter is updated before LV_EVENT_SHORT_CLICKED is fired.
 * @param indev pointer to an input device
 * @return short click streak counter
 */
// llgo:link (*LvIndevT).LvIndevGetShortClickStreak C.lv_indev_get_short_click_streak
func (recv_ *LvIndevT) LvIndevGetShortClickStreak() c.Uint8T {
	return 0
}

/**
 * Check the current scroll direction of an input device (for LV_INDEV_TYPE_POINTER and
 * LV_INDEV_TYPE_BUTTON)
 * @param indev pointer to an input device
 * @return LV_DIR_NONE: no scrolling now
 *         LV_DIR_HOR/VER
 */
// llgo:link (*LvIndevT).LvIndevGetScrollDir C.lv_indev_get_scroll_dir
func (recv_ *LvIndevT) LvIndevGetScrollDir() LvDirT {
	return 0
}

/**
 * Get the currently scrolled object (for LV_INDEV_TYPE_POINTER and
 * LV_INDEV_TYPE_BUTTON)
 * @param indev pointer to an input device
 * @return pointer to the currently scrolled object or NULL if no scrolling by this indev
 */
// llgo:link (*LvIndevT).LvIndevGetScrollObj C.lv_indev_get_scroll_obj
func (recv_ *LvIndevT) LvIndevGetScrollObj() *LvObjT {
	return nil
}

/**
 * Get the movement vector of an input device (for LV_INDEV_TYPE_POINTER and
 * LV_INDEV_TYPE_BUTTON)
 * @param indev pointer to an input device
 * @param point pointer to a point to store the types.pointer.vector
 */
// llgo:link (*LvIndevT).LvIndevGetVect C.lv_indev_get_vect
func (recv_ *LvIndevT) LvIndevGetVect(point *LvPointT) {
}

/**
 * Get the cursor object of an input device (for LV_INDEV_TYPE_POINTER only)
 * @param indev pointer to an input device
 * @return pointer to the cursor object
 */
// llgo:link (*LvIndevT).LvIndevGetCursor C.lv_indev_get_cursor
func (recv_ *LvIndevT) LvIndevGetCursor() *LvObjT {
	return nil
}

/**
 * Do nothing until the next release
 * @param indev pointer to an input device
 */
// llgo:link (*LvIndevT).LvIndevWaitRelease C.lv_indev_wait_release
func (recv_ *LvIndevT) LvIndevWaitRelease() {
}

/**
 * Gets a pointer to the currently active object in the currently processed input device.
 * @return pointer to currently active object or NULL if no active object
 */
//go:linkname LvIndevGetActiveObj C.lv_indev_get_active_obj
func LvIndevGetActiveObj() *LvObjT

/**
 * Get a pointer to the indev read timer to
 * modify its parameters with `lv_timer_...` functions.
 * @param indev pointer to an input device
 * @return pointer to the indev read refresher timer. (NULL on error)
 */
// llgo:link (*LvIndevT).LvIndevGetReadTimer C.lv_indev_get_read_timer
func (recv_ *LvIndevT) LvIndevGetReadTimer() *LvTimerT {
	return nil
}

/**
* Set the input device's event model: event-driven mode or timer mode.
* @param indev pointer to an input device
* @param mode the mode of input device
 */
// llgo:link (*LvIndevT).LvIndevSetMode C.lv_indev_set_mode
func (recv_ *LvIndevT) LvIndevSetMode(mode LvIndevModeT) {
}

/**
 * Get the input device's running mode.
 * @param indev pointer to an input device
 * @return the running mode for the specified input device.
 */
// llgo:link (*LvIndevT).LvIndevGetMode C.lv_indev_get_mode
func (recv_ *LvIndevT) LvIndevGetMode() LvIndevModeT {
	return 0
}

/**
 * Search the most top, clickable object by a point
 * @param obj pointer to a start object, typically the screen
 * @param point pointer to a point for searching the most top child
 * @return pointer to the found object or NULL if there was no suitable object
 */
// llgo:link (*LvObjT).LvIndevSearchObj C.lv_indev_search_obj
func (recv_ *LvObjT) LvIndevSearchObj(point *LvPointT) *LvObjT {
	return nil
}

/**
 * Add an event handler to the indev
 * @param indev          pointer to an indev
 * @param event_cb      an event callback
 * @param filter        event code to react or `LV_EVENT_ALL`
 * @param user_data     optional user_data
 */
// llgo:link (*LvIndevT).LvIndevAddEventCb C.lv_indev_add_event_cb
func (recv_ *LvIndevT) LvIndevAddEventCb(event_cb LvEventCbT, filter LvEventCodeT, user_data c.Pointer) {
}

/**
 * Get the number of event attached to an indev
 * @param indev          pointer to an indev
 * @return              number of events
 */
// llgo:link (*LvIndevT).LvIndevGetEventCount C.lv_indev_get_event_count
func (recv_ *LvIndevT) LvIndevGetEventCount() c.Uint32T {
	return 0
}

/**
 * Get an event descriptor for an event
 * @param indev          pointer to an indev
 * @param index         the index of the event
 * @return              the event descriptor
 */
// llgo:link (*LvIndevT).LvIndevGetEventDsc C.lv_indev_get_event_dsc
func (recv_ *LvIndevT) LvIndevGetEventDsc(index c.Uint32T) *LvEventDscT {
	return nil
}

/**
 * Remove an event
 * @param indev         pointer to an indev
 * @param index         the index of the event to remove
 * @return              true: and event was removed; false: no event was removed
 */
// llgo:link (*LvIndevT).LvIndevRemoveEvent C.lv_indev_remove_event
func (recv_ *LvIndevT) LvIndevRemoveEvent(index c.Uint32T) bool {
	return false
}

/**
 * Remove an event_cb with user_data
 * @param indev         pointer to a indev
 * @param event_cb      the event_cb of the event to remove
 * @param user_data     user_data
 * @return              the count of the event removed
 */
// llgo:link (*LvIndevT).LvIndevRemoveEventCbWithUserData C.lv_indev_remove_event_cb_with_user_data
func (recv_ *LvIndevT) LvIndevRemoveEventCbWithUserData(event_cb LvEventCbT, user_data c.Pointer) c.Uint32T {
	return 0
}

/**
 * Send an event to an indev
 * @param indev         pointer to an indev
 * @param code          an event code. LV_EVENT_...
 * @param param         optional param
 * @return              LV_RESULT_OK: indev wasn't deleted in the event.
 */
// llgo:link (*LvIndevT).LvIndevSendEvent C.lv_indev_send_event
func (recv_ *LvIndevT) LvIndevSendEvent(code LvEventCodeT, param c.Pointer) LvResultT {
	return 0
}

type LvCoverResT c.Int

const (
	LV_COVER_RES_COVER     LvCoverResT = 0
	LV_COVER_RES_NOT_COVER LvCoverResT = 1
	LV_COVER_RES_MASKED    LvCoverResT = 2
)

/**
 * Send an event to the object
 * @param obj           pointer to an object
 * @param event_code    the type of the event from `lv_event_t`
 * @param param         arbitrary data depending on the widget type and the event. (Usually `NULL`)
 * @return LV_RESULT_OK: `obj` was not deleted in the event; LV_RESULT_INVALID: `obj` was deleted in the event_code
 */
// llgo:link (*LvObjT).LvObjSendEvent C.lv_obj_send_event
func (recv_ *LvObjT) LvObjSendEvent(event_code LvEventCodeT, param c.Pointer) LvResultT {
	return 0
}

/**
 * Used by the widgets internally to call the ancestor widget types's event handler
 * @param class_p   pointer to the class of the widget (NOT the ancestor class)
 * @param e         pointer to the event descriptor
 * @return          LV_RESULT_OK: the target object was not deleted in the event; LV_RESULT_INVALID: it was deleted in the event_code
 */
// llgo:link (*LvObjClassT).LvObjEventBase C.lv_obj_event_base
func (recv_ *LvObjClassT) LvObjEventBase(e *LvEventT) LvResultT {
	return 0
}

/**
 * Get the current target of the event. It's the object which event handler being called.
 * If the event is not bubbled it's the same as "original" target.
 * @param e     pointer to the event descriptor
 * @return      the target of the event_code
 */
// llgo:link (*LvEventT).LvEventGetCurrentTargetObj C.lv_event_get_current_target_obj
func (recv_ *LvEventT) LvEventGetCurrentTargetObj() *LvObjT {
	return nil
}

/**
 * Get the object originally targeted by the event. It's the same even if the event is bubbled.
 * @param e     pointer to the event descriptor
 * @return      pointer to the original target of the event_code
 */
// llgo:link (*LvEventT).LvEventGetTargetObj C.lv_event_get_target_obj
func (recv_ *LvEventT) LvEventGetTargetObj() *LvObjT {
	return nil
}

/**
 * Add an event handler function for an object.
 * Used by the user to react on event which happens with the object.
 * An object can have multiple event handler. They will be called in the same order as they were added.
 * @param obj       pointer to an object
 * @param filter    an event code (e.g. `LV_EVENT_CLICKED`) on which the event should be called. `LV_EVENT_ALL` can be used to receive all the events.
 * @param event_cb  the new event function
 * @param           user_data custom data will be available in `event_cb`
 * @return          handler to the event. It can be used in `lv_obj_remove_event_dsc`.
 */
// llgo:link (*LvObjT).LvObjAddEventCb C.lv_obj_add_event_cb
func (recv_ *LvObjT) LvObjAddEventCb(event_cb LvEventCbT, filter LvEventCodeT, user_data c.Pointer) *LvEventDscT {
	return nil
}

// llgo:link (*LvObjT).LvObjGetEventCount C.lv_obj_get_event_count
func (recv_ *LvObjT) LvObjGetEventCount() c.Uint32T {
	return 0
}

// llgo:link (*LvObjT).LvObjGetEventDsc C.lv_obj_get_event_dsc
func (recv_ *LvObjT) LvObjGetEventDsc(index c.Uint32T) *LvEventDscT {
	return nil
}

// llgo:link (*LvObjT).LvObjRemoveEvent C.lv_obj_remove_event
func (recv_ *LvObjT) LvObjRemoveEvent(index c.Uint32T) bool {
	return false
}

// llgo:link (*LvObjT).LvObjRemoveEventDsc C.lv_obj_remove_event_dsc
func (recv_ *LvObjT) LvObjRemoveEventDsc(dsc *LvEventDscT) bool {
	return false
}

/**
 * Remove an event_cb from an object
 * @param obj           pointer to a obj
 * @param event_cb      the event_cb of the event to remove
 * @return              the count of the event removed
 */
// llgo:link (*LvObjT).LvObjRemoveEventCb C.lv_obj_remove_event_cb
func (recv_ *LvObjT) LvObjRemoveEventCb(event_cb LvEventCbT) c.Uint32T {
	return 0
}

/**
 * Remove an event_cb with user_data
 * @param obj           pointer to a obj
 * @param event_cb      the event_cb of the event to remove
 * @param user_data     user_data
 * @return              the count of the event removed
 */
// llgo:link (*LvObjT).LvObjRemoveEventCbWithUserData C.lv_obj_remove_event_cb_with_user_data
func (recv_ *LvObjT) LvObjRemoveEventCbWithUserData(event_cb LvEventCbT, user_data c.Pointer) c.Uint32T {
	return 0
}

/**
 * Get the input device passed as parameter to indev related events.
 * @param e     pointer to an event
 * @return      the indev that triggered the event or NULL if called on a not indev related event
 */
// llgo:link (*LvEventT).LvEventGetIndev C.lv_event_get_indev
func (recv_ *LvEventT) LvEventGetIndev() *LvIndevT {
	return nil
}

/**
 * Get the draw context which should be the first parameter of the draw functions.
 * Namely: `LV_EVENT_DRAW_MAIN/POST`, `LV_EVENT_DRAW_MAIN/POST_BEGIN`, `LV_EVENT_DRAW_MAIN/POST_END`
 * @param e     pointer to an event
 * @return      pointer to a draw context or NULL if called on an unrelated event
 */
// llgo:link (*LvEventT).LvEventGetLayer C.lv_event_get_layer
func (recv_ *LvEventT) LvEventGetLayer() *LvLayerT {
	return nil
}

/**
 * Get the old area of the object before its size was changed. Can be used in `LV_EVENT_SIZE_CHANGED`
 * @param e     pointer to an event
 * @return      the old absolute area of the object or NULL if called on an unrelated event
 */
// llgo:link (*LvEventT).LvEventGetOldSize C.lv_event_get_old_size
func (recv_ *LvEventT) LvEventGetOldSize() *LvAreaT {
	return nil
}

/**
 * Get the key passed as parameter to an event. Can be used in `LV_EVENT_KEY`
 * @param e     pointer to an event
 * @return      the triggering key or NULL if called on an unrelated event
 */
// llgo:link (*LvEventT).LvEventGetKey C.lv_event_get_key
func (recv_ *LvEventT) LvEventGetKey() c.Uint32T {
	return 0
}

/**
 * Get the signed rotary encoder diff. passed as parameter to an event. Can be used in `LV_EVENT_ROTARY`
 * @param e     pointer to an event
 * @return      the triggering key or NULL if called on an unrelated event
 */
// llgo:link (*LvEventT).LvEventGetRotaryDiff C.lv_event_get_rotary_diff
func (recv_ *LvEventT) LvEventGetRotaryDiff() c.Int32T {
	return 0
}

/**
 * Get the animation descriptor of a scrolling. Can be used in `LV_EVENT_SCROLL_BEGIN`
 * @param e     pointer to an event
 * @return      the animation that will scroll the object. (can be modified as required)
 */
// llgo:link (*LvEventT).LvEventGetScrollAnim C.lv_event_get_scroll_anim
func (recv_ *LvEventT) LvEventGetScrollAnim() *LvAnimT {
	return nil
}

/**
 * Set the new extra draw size. Can be used in `LV_EVENT_REFR_EXT_DRAW_SIZE`
 * @param e     pointer to an event
 * @param size  The new extra draw size
 */
// llgo:link (*LvEventT).LvEventSetExtDrawSize C.lv_event_set_ext_draw_size
func (recv_ *LvEventT) LvEventSetExtDrawSize(size c.Int32T) {
}

/**
 * Get a pointer to an `lv_point_t` variable in which the self size should be saved (width in `point->x` and height `point->y`).
 * Can be used in `LV_EVENT_GET_SELF_SIZE`
 * @param e     pointer to an event
 * @return      pointer to `lv_point_t` or NULL if called on an unrelated event
 */
// llgo:link (*LvEventT).LvEventGetSelfSizeInfo C.lv_event_get_self_size_info
func (recv_ *LvEventT) LvEventGetSelfSizeInfo() *LvPointT {
	return nil
}

/**
 * Get a pointer to an `lv_hit_test_info_t` variable in which the hit test result should be saved. Can be used in `LV_EVENT_HIT_TEST`
 * @param e     pointer to an event
 * @return      pointer to `lv_hit_test_info_t` or NULL if called on an unrelated event
 */
// llgo:link (*LvEventT).LvEventGetHitTestInfo C.lv_event_get_hit_test_info
func (recv_ *LvEventT) LvEventGetHitTestInfo() *LvHitTestInfoT {
	return nil
}

/**
 * Get a pointer to an area which should be examined whether the object fully covers it or not.
 * Can be used in `LV_EVENT_HIT_TEST`
 * @param e     pointer to an event
 * @return      an area with absolute coordinates to check
 */
// llgo:link (*LvEventT).LvEventGetCoverArea C.lv_event_get_cover_area
func (recv_ *LvEventT) LvEventGetCoverArea() *LvAreaT {
	return nil
}

/**
 * Set the result of cover checking. Can be used in `LV_EVENT_COVER_CHECK`
 * @param e     pointer to an event
 * @param res   an element of ::lv_cover_check_info_t
 */
// llgo:link (*LvEventT).LvEventSetCoverRes C.lv_event_set_cover_res
func (recv_ *LvEventT) LvEventSetCoverRes(res LvCoverResT) {
}

/**
 * Get the draw task which was just added.
 * Can be used in `LV_EVENT_DRAW_TASK_ADDED event`
 * @param e     pointer to an event
 * @return      the added draw task
 */
// llgo:link (*LvEventT).LvEventGetDrawTask C.lv_event_get_draw_task
func (recv_ *LvEventT) LvEventGetDrawTask() *LvDrawTaskT {
	return nil
}

type X_lvStateT c.Int

const (
	LV_STATE_DEFAULT   X_lvStateT = 0
	LV_STATE_CHECKED   X_lvStateT = 1
	LV_STATE_FOCUSED   X_lvStateT = 2
	LV_STATE_FOCUS_KEY X_lvStateT = 4
	LV_STATE_EDITED    X_lvStateT = 8
	LV_STATE_HOVERED   X_lvStateT = 16
	LV_STATE_PRESSED   X_lvStateT = 32
	LV_STATE_SCROLLED  X_lvStateT = 64
	LV_STATE_DISABLED  X_lvStateT = 128
	LV_STATE_USER_1    X_lvStateT = 4096
	LV_STATE_USER_2    X_lvStateT = 8192
	LV_STATE_USER_3    X_lvStateT = 16384
	LV_STATE_USER_4    X_lvStateT = 32768
	LV_STATE_ANY       X_lvStateT = 65535
)

type X_lvPartT c.Int

const (
	LV_PART_MAIN         X_lvPartT = 0
	LV_PART_SCROLLBAR    X_lvPartT = 65536
	LV_PART_INDICATOR    X_lvPartT = 131072
	LV_PART_KNOB         X_lvPartT = 196608
	LV_PART_SELECTED     X_lvPartT = 262144
	LV_PART_ITEMS        X_lvPartT = 327680
	LV_PART_CURSOR       X_lvPartT = 393216
	LV_PART_CUSTOM_FIRST X_lvPartT = 524288
	LV_PART_ANY          X_lvPartT = 983040
)

type LvObjFlagT c.Int

const (
	LV_OBJ_FLAG_HIDDEN                LvObjFlagT = 1
	LV_OBJ_FLAG_CLICKABLE             LvObjFlagT = 2
	LV_OBJ_FLAG_CLICK_FOCUSABLE       LvObjFlagT = 4
	LV_OBJ_FLAG_CHECKABLE             LvObjFlagT = 8
	LV_OBJ_FLAG_SCROLLABLE            LvObjFlagT = 16
	LV_OBJ_FLAG_SCROLL_ELASTIC        LvObjFlagT = 32
	LV_OBJ_FLAG_SCROLL_MOMENTUM       LvObjFlagT = 64
	LV_OBJ_FLAG_SCROLL_ONE            LvObjFlagT = 128
	LV_OBJ_FLAG_SCROLL_CHAIN_HOR      LvObjFlagT = 256
	LV_OBJ_FLAG_SCROLL_CHAIN_VER      LvObjFlagT = 512
	LV_OBJ_FLAG_SCROLL_CHAIN          LvObjFlagT = 768
	LV_OBJ_FLAG_SCROLL_ON_FOCUS       LvObjFlagT = 1024
	LV_OBJ_FLAG_SCROLL_WITH_ARROW     LvObjFlagT = 2048
	LV_OBJ_FLAG_SNAPPABLE             LvObjFlagT = 4096
	LV_OBJ_FLAG_PRESS_LOCK            LvObjFlagT = 8192
	LV_OBJ_FLAG_EVENT_BUBBLE          LvObjFlagT = 16384
	LV_OBJ_FLAG_GESTURE_BUBBLE        LvObjFlagT = 32768
	LV_OBJ_FLAG_ADV_HITTEST           LvObjFlagT = 65536
	LV_OBJ_FLAG_IGNORE_LAYOUT         LvObjFlagT = 131072
	LV_OBJ_FLAG_FLOATING              LvObjFlagT = 262144
	LV_OBJ_FLAG_SEND_DRAW_TASK_EVENTS LvObjFlagT = 524288
	LV_OBJ_FLAG_OVERFLOW_VISIBLE      LvObjFlagT = 1048576
	LV_OBJ_FLAG_FLEX_IN_NEW_TRACK     LvObjFlagT = 2097152
	LV_OBJ_FLAG_LAYOUT_1              LvObjFlagT = 8388608
	LV_OBJ_FLAG_LAYOUT_2              LvObjFlagT = 16777216
	LV_OBJ_FLAG_WIDGET_1              LvObjFlagT = 33554432
	LV_OBJ_FLAG_WIDGET_2              LvObjFlagT = 67108864
	LV_OBJ_FLAG_USER_1                LvObjFlagT = 134217728
	LV_OBJ_FLAG_USER_2                LvObjFlagT = 268435456
	LV_OBJ_FLAG_USER_3                LvObjFlagT = 536870912
	LV_OBJ_FLAG_USER_4                LvObjFlagT = 1073741824
)

/**
 * Create a base object (a rectangle)
 * @param parent    pointer to a parent object. If NULL then a screen will be created.
 * @return          pointer to the new object
 */
// llgo:link (*LvObjT).LvObjCreate C.lv_obj_create
func (recv_ *LvObjT) LvObjCreate() *LvObjT {
	return nil
}

/**
 * Set one or more flags
 * @param obj   pointer to an object
 * @param f     OR-ed values from `lv_obj_flag_t` to set.
 */
// llgo:link (*LvObjT).LvObjAddFlag C.lv_obj_add_flag
func (recv_ *LvObjT) LvObjAddFlag(f LvObjFlagT) {
}

/**
 * Remove one or more flags
 * @param obj   pointer to an object
 * @param f     OR-ed values from `lv_obj_flag_t` to clear.
 */
// llgo:link (*LvObjT).LvObjRemoveFlag C.lv_obj_remove_flag
func (recv_ *LvObjT) LvObjRemoveFlag(f LvObjFlagT) {
}

/**
 * Set add or remove one or more flags.
 * @param obj   pointer to an object
 * @param f     OR-ed values from `lv_obj_flag_t` to update.
 * @param v     true: add the flags; false: remove the flags
 */
// llgo:link (*LvObjT).LvObjSetFlag C.lv_obj_set_flag
func (recv_ *LvObjT) LvObjSetFlag(f LvObjFlagT, v bool) {
}

/**
 * Add one or more states to the object. The other state bits will remain unchanged.
 * If specified in the styles, transition animation will be started from the previous state to the current.
 * @param obj       pointer to an object
 * @param state     the states to add. E.g `LV_STATE_PRESSED | LV_STATE_FOCUSED`
 */
// llgo:link (*LvObjT).LvObjAddState C.lv_obj_add_state
func (recv_ *LvObjT) LvObjAddState(state LvStateT) {
}

/**
 * Remove one or more states to the object. The other state bits will remain unchanged.
 * If specified in the styles, transition animation will be started from the previous state to the current.
 * @param obj       pointer to an object
 * @param state     the states to add. E.g `LV_STATE_PRESSED | LV_STATE_FOCUSED`
 */
// llgo:link (*LvObjT).LvObjRemoveState C.lv_obj_remove_state
func (recv_ *LvObjT) LvObjRemoveState(state LvStateT) {
}

/**
 * Add or remove one or more states to the object. The other state bits will remain unchanged.
 * @param obj       pointer to an object
 * @param state     the states to add. E.g `LV_STATE_PRESSED | LV_STATE_FOCUSED`
 * @param v         true: add the states; false: remove the states
 */
// llgo:link (*LvObjT).LvObjSetState C.lv_obj_set_state
func (recv_ *LvObjT) LvObjSetState(state LvStateT, v bool) {
}

/**
 * Set the user_data field of the object
 * @param obj   pointer to an object
 * @param user_data   pointer to the new user_data.
 */
// llgo:link (*LvObjT).LvObjSetUserData C.lv_obj_set_user_data
func (recv_ *LvObjT) LvObjSetUserData(user_data c.Pointer) {
}

/**
 * Check if a given flag or all the given flags are set on an object.
 * @param obj   pointer to an object
 * @param f     the flag(s) to check (OR-ed values can be used)
 * @return      true: all flags are set; false: not all flags are set
 */
// llgo:link (*LvObjT).LvObjHasFlag C.lv_obj_has_flag
func (recv_ *LvObjT) LvObjHasFlag(f LvObjFlagT) bool {
	return false
}

/**
 * Check if a given flag or any of the flags are set on an object.
 * @param obj   pointer to an object
 * @param f     the flag(s) to check (OR-ed values can be used)
 * @return      true: at least one flag is set; false: none of the flags are set
 */
// llgo:link (*LvObjT).LvObjHasFlagAny C.lv_obj_has_flag_any
func (recv_ *LvObjT) LvObjHasFlagAny(f LvObjFlagT) bool {
	return false
}

/**
 * Get the state of an object
 * @param obj   pointer to an object
 * @return      the state (OR-ed values from `lv_state_t`)
 */
// llgo:link (*LvObjT).LvObjGetState C.lv_obj_get_state
func (recv_ *LvObjT) LvObjGetState() LvStateT {
	return 0
}

/**
 * Check if the object is in a given state or not.
 * @param obj       pointer to an object
 * @param state     a state or combination of states to check
 * @return          true: `obj` is in `state`; false: `obj` is not in `state`
 */
// llgo:link (*LvObjT).LvObjHasState C.lv_obj_has_state
func (recv_ *LvObjT) LvObjHasState(state LvStateT) bool {
	return false
}

/**
 * Get the group of the object
 * @param       obj pointer to an object
 * @return      the pointer to group of the object
 */
// llgo:link (*LvObjT).LvObjGetGroup C.lv_obj_get_group
func (recv_ *LvObjT) LvObjGetGroup() *LvGroupT {
	return nil
}

/**
 * Get the user_data field of the object
 * @param obj   pointer to an object
 * @return      the pointer to the user_data of the object
 */
// llgo:link (*LvObjT).LvObjGetUserData C.lv_obj_get_user_data
func (recv_ *LvObjT) LvObjGetUserData() c.Pointer {
	return nil
}

/**
 * Allocate special data for an object if not allocated yet.
 * @param obj   pointer to an object
 */
// llgo:link (*LvObjT).LvObjAllocateSpecAttr C.lv_obj_allocate_spec_attr
func (recv_ *LvObjT) LvObjAllocateSpecAttr() {
}

/**
 * Check the type of obj.
 * @param obj       pointer to an object
 * @param class_p   a class to check (e.g. `lv_slider_class`)
 * @return          true: `class_p` is the `obj` class.
 */
// llgo:link (*LvObjT).LvObjCheckType C.lv_obj_check_type
func (recv_ *LvObjT) LvObjCheckType(class_p *LvObjClassT) bool {
	return false
}

/**
 * Check if any object has a given class (type).
 * It checks the ancestor classes too.
 * @param obj       pointer to an object
 * @param class_p   a class to check (e.g. `lv_slider_class`)
 * @return          true: `obj` has the given class
 */
// llgo:link (*LvObjT).LvObjHasClass C.lv_obj_has_class
func (recv_ *LvObjT) LvObjHasClass(class_p *LvObjClassT) bool {
	return false
}

/**
 * Get the class (type) of the object
 * @param obj   pointer to an object
 * @return      the class (type) of the object
 */
// llgo:link (*LvObjT).LvObjGetClass C.lv_obj_get_class
func (recv_ *LvObjT) LvObjGetClass() *LvObjClassT {
	return nil
}

/**
 * Check if any object is still "alive".
 * @param obj       pointer to an object
 * @return          true: valid
 */
// llgo:link (*LvObjT).LvObjIsValid C.lv_obj_is_valid
func (recv_ *LvObjT) LvObjIsValid() bool {
	return false
}

/**
 * Utility to set an object reference to NULL when it gets deleted.
 * The reference should be in a location that will not become invalid
 * during the object's lifetime, i.e. static or allocated.
 * @param obj_ptr   a pointer to a pointer to an object
 */
//go:linkname LvObjNullOnDelete C.lv_obj_null_on_delete
func LvObjNullOnDelete(obj_ptr **LvObjT)

/**
 * Add an event handler to a widget that will load a screen on a trigger.
 * @param obj           pointer to widget which should load the screen
 * @param trigger       an event code, e.g. `LV_EVENT_CLICKED`
 * @param screen        the screen to load (must be a valid widget)
 * @param anim_type     element of `lv_screen_load_anim_t` the screen load animation
 * @param duration      duration of the animation in milliseconds
 * @param delay         delay before the screen load in milliseconds
 */
// llgo:link (*LvObjT).LvObjAddScreenLoadEvent C.lv_obj_add_screen_load_event
func (recv_ *LvObjT) LvObjAddScreenLoadEvent(trigger LvEventCodeT, screen *LvObjT, anim_type LvScreenLoadAnimT, duration c.Uint32T, delay c.Uint32T) {
}

/**
 * Add an event handler to a widget that will create a screen on a trigger.
 * The created screen will be deleted when it's unloaded
 * @param obj               pointer to widget which should load the screen
 * @param trigger           an event code, e.g. `LV_EVENT_CLICKED`
 * @param screen_create_cb  a callback to create the screen, e.g. `lv_obj_t * myscreen_create(void)`
 * @param anim_type         element of `lv_screen_load_anim_t` the screen load animation
 * @param duration          duration of the animation in milliseconds
 * @param delay             delay before the screen load in milliseconds
 */
// llgo:link (*LvObjT).LvObjAddScreenCreateEvent C.lv_obj_add_screen_create_event
func (recv_ *LvObjT) LvObjAddScreenCreateEvent(trigger LvEventCodeT, screen_create_cb LvScreenCreateCbT, anim_type LvScreenLoadAnimT, duration c.Uint32T, delay c.Uint32T) {
}

type LvSubjectTypeT c.Int

const (
	LV_SUBJECT_TYPE_INVALID LvSubjectTypeT = 0
	LV_SUBJECT_TYPE_NONE    LvSubjectTypeT = 1
	LV_SUBJECT_TYPE_INT     LvSubjectTypeT = 2
	LV_SUBJECT_TYPE_FLOAT   LvSubjectTypeT = 3
	LV_SUBJECT_TYPE_POINTER LvSubjectTypeT = 4
	LV_SUBJECT_TYPE_COLOR   LvSubjectTypeT = 5
	LV_SUBJECT_TYPE_GROUP   LvSubjectTypeT = 6
	LV_SUBJECT_TYPE_STRING  LvSubjectTypeT = 7
)

/**
 * A common type to handle all the various observable types in the same way
 */

type LvSubjectValueT struct {
	Pointer c.Pointer
}

/**
 * The Subject (an observable value)
 */

type LvSubjectT struct {
	SubsLl             LvLlT
	Value              LvSubjectValueT
	PrevValue          LvSubjectValueT
	UserData           c.Pointer
	Type               c.Uint32T
	Size               c.Uint32T
	NotifyRestartQuery c.Uint32T
}

// llgo:type C
type LvObserverCbT func(*LvObserverT, *LvSubjectT)

/**
 * Initialize an integer-type Subject.
 * @param subject   pointer to Subject
 * @param value     initial value
 */
// llgo:link (*LvSubjectT).LvSubjectInitInt C.lv_subject_init_int
func (recv_ *LvSubjectT) LvSubjectInitInt(value c.Int32T) {
}

/**
 * Set value of an integer Subject and notify Observers.
 * @param subject   pointer to Subject
 * @param value     new value
 */
// llgo:link (*LvSubjectT).LvSubjectSetInt C.lv_subject_set_int
func (recv_ *LvSubjectT) LvSubjectSetInt(value c.Int32T) {
}

/**
 * Get current value of an integer Subject.
 * @param subject   pointer to Subject
 * @return          current value
 */
// llgo:link (*LvSubjectT).LvSubjectGetInt C.lv_subject_get_int
func (recv_ *LvSubjectT) LvSubjectGetInt() c.Int32T {
	return 0
}

/**
 * Get previous value of an integer Subject.
 * @param subject   pointer to Subject
 * @return          current value
 */
// llgo:link (*LvSubjectT).LvSubjectGetPreviousInt C.lv_subject_get_previous_int
func (recv_ *LvSubjectT) LvSubjectGetPreviousInt() c.Int32T {
	return 0
}

/**
 * Initialize a string-type Subject.
 * @param subject   pointer to Subject
 * @param buf       pointer to buffer to store string
 * @param prev_buf  pointer to buffer to store previous string; can be NULL if not used
 * @param size      size of buffer(s)
 * @param value     initial value of string, e.g. "hello"
 * @note            A string Subject stores its own copy of the string, not just the pointer.
 */
// llgo:link (*LvSubjectT).LvSubjectInitString C.lv_subject_init_string
func (recv_ *LvSubjectT) LvSubjectInitString(buf *c.Char, prev_buf *c.Char, size c.SizeT, value *c.Char) {
}

/**
 * Copy a string to a Subject and notify Observers if it changed.
 * @param subject   pointer to Subject
 * @param buf       new string
 */
// llgo:link (*LvSubjectT).LvSubjectCopyString C.lv_subject_copy_string
func (recv_ *LvSubjectT) LvSubjectCopyString(buf *c.Char) {
}

/**
 * Format a new string, updating Subject, and notify Observers if it changed.
 * @param subject   pointer to Subject
 * @param format    format string
 */
// llgo:link (*LvSubjectT).LvSubjectSnprintf C.lv_subject_snprintf
func (recv_ *LvSubjectT) LvSubjectSnprintf(format *c.Char, __llgo_va_list ...interface{}) {
}

/**
 * Get current value of a string Subject.
 * @param subject   pointer to Subject
 * @return          pointer to buffer containing current value
 */
// llgo:link (*LvSubjectT).LvSubjectGetString C.lv_subject_get_string
func (recv_ *LvSubjectT) LvSubjectGetString() *c.Char {
	return nil
}

/**
 * Get previous value of a string Subject.
 * @param subject   pointer to Subject
 * @return          pointer to buffer containing previous value
 * @note            NULL will be returned if NULL was passed in `lv_subject_init_string()`
 *                  as `prev_buf`.
 */
// llgo:link (*LvSubjectT).LvSubjectGetPreviousString C.lv_subject_get_previous_string
func (recv_ *LvSubjectT) LvSubjectGetPreviousString() *c.Char {
	return nil
}

/**
 * Initialize a pointer-type Subject.
 * @param subject   pointer to Subject
 * @param value     initial value
 */
// llgo:link (*LvSubjectT).LvSubjectInitPointer C.lv_subject_init_pointer
func (recv_ *LvSubjectT) LvSubjectInitPointer(value c.Pointer) {
}

/**
 * Set value of a pointer Subject and notify Observers (regardless of whether it changed).
 * @param subject   pointer to Subject
 * @param ptr       new value
 */
// llgo:link (*LvSubjectT).LvSubjectSetPointer C.lv_subject_set_pointer
func (recv_ *LvSubjectT) LvSubjectSetPointer(ptr c.Pointer) {
}

/**
 * Get current value of a pointer Subject.
 * @param subject   pointer to Subject
 * @return          current value
 */
// llgo:link (*LvSubjectT).LvSubjectGetPointer C.lv_subject_get_pointer
func (recv_ *LvSubjectT) LvSubjectGetPointer() c.Pointer {
	return nil
}

/**
 * Get previous value of a pointer Subject.
 * @param subject   pointer to Subject
 * @return          previous value
 */
// llgo:link (*LvSubjectT).LvSubjectGetPreviousPointer C.lv_subject_get_previous_pointer
func (recv_ *LvSubjectT) LvSubjectGetPreviousPointer() c.Pointer {
	return nil
}

/**
 * Initialize a color-type Subject.
 * @param subject   pointer to Subject
 * @param color     initial value
 */
// llgo:link (*LvSubjectT).LvSubjectInitColor C.lv_subject_init_color
func (recv_ *LvSubjectT) LvSubjectInitColor(color LvColorT) {
}

/**
 * Set value of a color Subject and notify Observers if it changed.
 * @param subject   pointer to Subject
 * @param color     new value
 */
// llgo:link (*LvSubjectT).LvSubjectSetColor C.lv_subject_set_color
func (recv_ *LvSubjectT) LvSubjectSetColor(color LvColorT) {
}

/**
 * Get current value of a color Subject.
 * @param subject   pointer to Subject
 * @return          current value
 */
// llgo:link (*LvSubjectT).LvSubjectGetColor C.lv_subject_get_color
func (recv_ *LvSubjectT) LvSubjectGetColor() LvColorT {
	return LvColorT{}
}

/**
 * Get previous value of a color Subject.
 * @param subject   pointer to Subject
 * @return          previous value
 */
// llgo:link (*LvSubjectT).LvSubjectGetPreviousColor C.lv_subject_get_previous_color
func (recv_ *LvSubjectT) LvSubjectGetPreviousColor() LvColorT {
	return LvColorT{}
}

/**
* Initialize a Group-type Subject.
* @param group_subject  pointer to Group-type Subject
* @param list           list of other Subject addresses; when any of these have values
                            updated, Observers of `group_subject` will be notified.
* @param list_len       number of elements in `list[]`
*/
// llgo:link (*LvSubjectT).LvSubjectInitGroup C.lv_subject_init_group
func (recv_ *LvSubjectT) LvSubjectInitGroup(list **LvSubjectT, list_len c.Uint32T) {
}

/**
 * Remove all Observers from a Subject and free allocated memory, and delete
 * any associated Widget-Binding events.  This leaves `subject` "disconnected" from
 * all Observers and all associated Widget events established through Widget Binding.
 * @param subject   pointer to Subject
 * @note            This can safely be called regardless of whether any Observers
 *                  added with `lv_subject_add_observer_obj()` or bound to a Widget Property
 *                  with one of the `..._bind_...()` functions.
 */
// llgo:link (*LvSubjectT).LvSubjectDeinit C.lv_subject_deinit
func (recv_ *LvSubjectT) LvSubjectDeinit() {
}

/**
 * Get an element from Subject Group's list.
 * @param subject   pointer to Group-type Subject
 * @param index     index of element to get
 * @return          pointer to indexed Subject from list, or NULL if index is out of bounds
 */
// llgo:link (*LvSubjectT).LvSubjectGetGroupElement C.lv_subject_get_group_element
func (recv_ *LvSubjectT) LvSubjectGetGroupElement(index c.Int32T) *LvSubjectT {
	return nil
}

/**
 * Add Observer to Subject. When Subject's value changes `observer_cb` will be called.
 * @param subject       pointer to Subject
 * @param observer_cb   notification callback
 * @param user_data     optional user data
 * @return              pointer to newly-created Observer
 */
// llgo:link (*LvSubjectT).LvSubjectAddObserver C.lv_subject_add_observer
func (recv_ *LvSubjectT) LvSubjectAddObserver(observer_cb LvObserverCbT, user_data c.Pointer) *LvObserverT {
	return nil
}

/**
 * Add Observer to Subject for a Widget.
 * When the Widget is deleted, Observer will be unsubscribed from Subject automatically.
 * @param subject       pointer to Subject
 * @param observer_cb   notification callback
 * @param obj           pointer to Widget
 * @param user_data     optional user data
 * @return              pointer to newly-created Observer
 * @note                Do not call `lv_observer_remove()` on Observers created this way.
 *                      Only clean up such Observers by either:
 *                      - deleting the Widget, or
 *                      - calling `lv_subject_deinit()` to gracefully de-couple and
 *                        remove all Observers.
 */
// llgo:link (*LvSubjectT).LvSubjectAddObserverObj C.lv_subject_add_observer_obj
func (recv_ *LvSubjectT) LvSubjectAddObserverObj(observer_cb LvObserverCbT, obj *LvObjT, user_data c.Pointer) *LvObserverT {
	return nil
}

/**
 * Add an Observer to a Subject and also save a target pointer.
 * @param subject       pointer to Subject
 * @param observer_cb   notification callback
 * @param target        any pointer (NULL is okay)
 * @param user_data     optional user data
 * @return              pointer to newly-created Observer
 */
// llgo:link (*LvSubjectT).LvSubjectAddObserverWithTarget C.lv_subject_add_observer_with_target
func (recv_ *LvSubjectT) LvSubjectAddObserverWithTarget(observer_cb LvObserverCbT, target c.Pointer, user_data c.Pointer) *LvObserverT {
	return nil
}

/**
 * Remove Observer from its Subject.
 * @param observer      pointer to Observer
 */
// llgo:link (*LvObserverT).LvObserverRemove C.lv_observer_remove
func (recv_ *LvObserverT) LvObserverRemove() {
}

/**
 * Remove Observers associated with Widget `obj` from specified `subject` or all Subjects.
 * @param obj       pointer to Widget whose Observers should be removed
 * @param subject   Subject to remove Widget from, or NULL to remove from all Subjects
 * @note            This function can be used e.g. when a Widget's Subject(s) needs to
 *                  be replaced by other Subject(s)
 */
// llgo:link (*LvObjT).LvObjRemoveFromSubject C.lv_obj_remove_from_subject
func (recv_ *LvObjT) LvObjRemoveFromSubject(subject *LvSubjectT) {
}

/**
 * Get target of an Observer.
 * @param observer      pointer to Observer
 * @return              pointer to saved target
 */
// llgo:link (*LvObserverT).LvObserverGetTarget C.lv_observer_get_target
func (recv_ *LvObserverT) LvObserverGetTarget() c.Pointer {
	return nil
}

/**
 * Get target Widget of Observer.
 * This is the same as `lv_observer_get_target()`, except it returns `target`
 * as an `lv_obj_t *`.
 * @param observer      pointer to Observer
 * @return              pointer to saved Widget target
 */
// llgo:link (*LvObserverT).LvObserverGetTargetObj C.lv_observer_get_target_obj
func (recv_ *LvObserverT) LvObserverGetTargetObj() *LvObjT {
	return nil
}

/**
 * Get Observer's user data.
 * @param observer      pointer to Observer
 * @return              void pointer to saved user data
 */
// llgo:link (*LvObserverT).LvObserverGetUserData C.lv_observer_get_user_data
func (recv_ *LvObserverT) LvObserverGetUserData() c.Pointer {
	return nil
}

/**
 * Notify all Observers of Subject.
 * @param subject       pointer to Subject
 */
// llgo:link (*LvSubjectT).LvSubjectNotify C.lv_subject_notify
func (recv_ *LvSubjectT) LvSubjectNotify() {
}

/**
 * Add an event handler to increment (or decrement) the value of a subject on a trigger.
 * @param obj       pointer to a widget
 * @param subject   pointer to a subject to change
 * @param trigger   the trigger on which the subject should be changed
 * @param step      value to add on trigger
 * @param min       the minimum value
 * @param max       the maximum value
 */
// llgo:link (*LvObjT).LvObjAddSubjectIncrementEvent C.lv_obj_add_subject_increment_event
func (recv_ *LvObjT) LvObjAddSubjectIncrementEvent(subject *LvSubjectT, trigger LvEventCodeT, step c.Int32T, min c.Int32T, max c.Int32T) {
}

/**
 * Set the value of an integer subject.
 * @param obj       pointer to a widget
 * @param subject   pointer to a subject to change
 * @param trigger   the trigger on which the subject should be changed
 * @param value     the value to set
 */
// llgo:link (*LvObjT).LvObjAddSubjectSetIntEvent C.lv_obj_add_subject_set_int_event
func (recv_ *LvObjT) LvObjAddSubjectSetIntEvent(subject *LvSubjectT, trigger LvEventCodeT, value c.Int32T) {
}

/**
 * Set the value of a string subject.
 * @param obj       pointer to a widget
 * @param subject   pointer to a subject to change
 * @param trigger   the trigger on which the subject should be changed
 * @param value     the value to set
 */
// llgo:link (*LvObjT).LvObjAddSubjectSetStringEvent C.lv_obj_add_subject_set_string_event
func (recv_ *LvObjT) LvObjAddSubjectSetStringEvent(subject *LvSubjectT, trigger LvEventCodeT, value *c.Char) {
}

/**
 * Disable a style if a subject's value is not equal to a reference value
 * @param obj           pointer to Widget
 * @param style         pointer to a style
 * @param selector      pointer to a selector
 * @param subject       pointer to Subject
 * @param ref_value     reference value to compare Subject's value with
 * @return              pointer to newly-created Observer
 */
// llgo:link (*LvObjT).LvObjBindStyle C.lv_obj_bind_style
func (recv_ *LvObjT) LvObjBindStyle(style *LvStyleT, selector LvStyleSelectorT, subject *LvSubjectT, ref_value c.Int32T) *LvObserverT {
	return nil
}

/**
 * Set Widget's flag(s) if an integer Subject's value is equal to a reference value, clear flag otherwise.
 * @param obj           pointer to Widget
 * @param subject       pointer to Subject
 * @param flag          flag(s) (can be bit-wise OR-ed) to set or clear (e.g. `LV_OBJ_FLAG_HIDDEN`)
 * @param ref_value     reference value to compare Subject's value with
 * @return              pointer to newly-created Observer
 */
// llgo:link (*LvObjT).LvObjBindFlagIfEq C.lv_obj_bind_flag_if_eq
func (recv_ *LvObjT) LvObjBindFlagIfEq(subject *LvSubjectT, flag LvObjFlagT, ref_value c.Int32T) *LvObserverT {
	return nil
}

/**
 * Set Widget's flag(s) if an integer Subject's value is not equal to a reference value, clear flag otherwise.
 * @param obj           pointer to Widget
 * @param subject       pointer to Subject
 * @param flag          flag(s) (can be bit-wise OR-ed) to set or clear (e.g. `LV_OBJ_FLAG_HIDDEN`)
 * @param ref_value     reference value to compare Subject's value with
 * @return              pointer to newly-created Observer
 */
// llgo:link (*LvObjT).LvObjBindFlagIfNotEq C.lv_obj_bind_flag_if_not_eq
func (recv_ *LvObjT) LvObjBindFlagIfNotEq(subject *LvSubjectT, flag LvObjFlagT, ref_value c.Int32T) *LvObserverT {
	return nil
}

/**
 * Set Widget's flag(s) if an integer Subject's value is greater than a reference value, clear flag otherwise.
 * @param obj           pointer to Widget
 * @param subject       pointer to Subject
 * @param flag          flag(s) (can be bit-wise OR-ed) to set or clear (e.g. `LV_OBJ_FLAG_HIDDEN`)
 * @param ref_value     reference value to compare Subject's value with
 * @return              pointer to newly-created Observer
 */
// llgo:link (*LvObjT).LvObjBindFlagIfGt C.lv_obj_bind_flag_if_gt
func (recv_ *LvObjT) LvObjBindFlagIfGt(subject *LvSubjectT, flag LvObjFlagT, ref_value c.Int32T) *LvObserverT {
	return nil
}

/**
 * Set Widget's flag(s) if an integer Subject's value is greater than or equal to a reference value, clear flag otherwise.
 * @param obj           pointer to Widget
 * @param subject       pointer to Subject
 * @param flag          flag(s) (can be bit-wise OR-ed) to set or clear (e.g. `LV_OBJ_FLAG_HIDDEN`)
 * @param ref_value     reference value to compare Subject's value with
 * @return              pointer to newly-created Observer
 */
// llgo:link (*LvObjT).LvObjBindFlagIfGe C.lv_obj_bind_flag_if_ge
func (recv_ *LvObjT) LvObjBindFlagIfGe(subject *LvSubjectT, flag LvObjFlagT, ref_value c.Int32T) *LvObserverT {
	return nil
}

/**
 * Set Widget's flag(s) if an integer Subject's value is less than a reference value, clear flag otherwise.
 * @param obj           pointer to Widget
 * @param subject       pointer to Subject
 * @param flag          flag(s) (can be bit-wise OR-ed) to set or clear (e.g. `LV_OBJ_FLAG_HIDDEN`)
 * @param ref_value     reference value to compare Subject's value with
 * @return              pointer to newly-created Observer
 */
// llgo:link (*LvObjT).LvObjBindFlagIfLt C.lv_obj_bind_flag_if_lt
func (recv_ *LvObjT) LvObjBindFlagIfLt(subject *LvSubjectT, flag LvObjFlagT, ref_value c.Int32T) *LvObserverT {
	return nil
}

/**
 * Set Widget's flag(s) if an integer Subject's value is less than or equal to a reference value, clear flag otherwise.
 * @param obj           pointer to Widget
 * @param subject       pointer to Subject
 * @param flag          flag(s) (can be bit-wise OR-ed) to set or clear (e.g. `LV_OBJ_FLAG_HIDDEN`)
 * @param ref_value     reference value to compare Subject's value with
 * @return              pointer to newly-created Observer
 */
// llgo:link (*LvObjT).LvObjBindFlagIfLe C.lv_obj_bind_flag_if_le
func (recv_ *LvObjT) LvObjBindFlagIfLe(subject *LvSubjectT, flag LvObjFlagT, ref_value c.Int32T) *LvObserverT {
	return nil
}

/**
 * Set Widget's state(s) if an integer Subject's value is equal to a reference value, clear flag otherwise.
 * @param obj           pointer to Widget
 * @param subject       pointer to Subject
 * @param state         state(s) (can be bit-wise OR-ed) to set or clear (e.g. `LV_STATE_CHECKED`)
 * @param ref_value     reference value to compare Subject's value with
 * @return              pointer to newly-created Observer
 */
// llgo:link (*LvObjT).LvObjBindStateIfEq C.lv_obj_bind_state_if_eq
func (recv_ *LvObjT) LvObjBindStateIfEq(subject *LvSubjectT, state LvStateT, ref_value c.Int32T) *LvObserverT {
	return nil
}

/**
 * Set a Widget's state(s) if an integer Subject's value is not equal to a reference value, clear flag otherwise
 * @param obj           pointer to Widget
 * @param subject       pointer to Subject
 * @param state         state(s) (can be bit-wise OR-ed) to set or clear (e.g. `LV_STATE_CHECKED`)
 * @param ref_value     reference value to compare Subject's value with
 * @return              pointer to newly-created Observer
 */
// llgo:link (*LvObjT).LvObjBindStateIfNotEq C.lv_obj_bind_state_if_not_eq
func (recv_ *LvObjT) LvObjBindStateIfNotEq(subject *LvSubjectT, state LvStateT, ref_value c.Int32T) *LvObserverT {
	return nil
}

/**
 * Set Widget's state(s) if an integer Subject's value is greater than a reference value, clear flag otherwise.
 * @param obj           pointer to Widget
 * @param subject       pointer to Subject
 * @param state         state(s) (can be bit-wise OR-ed) to set or clear (e.g. `LV_STATE_CHECKED`)
 * @param ref_value     reference value to compare Subject's value with
 * @return              pointer to newly-created Observer
 */
// llgo:link (*LvObjT).LvObjBindStateIfGt C.lv_obj_bind_state_if_gt
func (recv_ *LvObjT) LvObjBindStateIfGt(subject *LvSubjectT, state LvStateT, ref_value c.Int32T) *LvObserverT {
	return nil
}

/**
 * Set Widget's state(s) if an integer Subject's value is greater than or equal to a reference value, clear flag otherwise.
 * @param obj           pointer to Widget
 * @param subject       pointer to Subject
 * @param state         state(s) (can be bit-wise OR-ed) to set or clear (e.g. `LV_STATE_CHECKED`)
 * @param ref_value     reference value to compare Subject's value with
 * @return              pointer to newly-created Observer
 */
// llgo:link (*LvObjT).LvObjBindStateIfGe C.lv_obj_bind_state_if_ge
func (recv_ *LvObjT) LvObjBindStateIfGe(subject *LvSubjectT, state LvStateT, ref_value c.Int32T) *LvObserverT {
	return nil
}

/**
 * Set Widget's state(s) if an integer Subject's value is less than a reference value, clear flag otherwise.
 * @param obj           pointer to Widget
 * @param subject       pointer to Subject
 * @param state         state(s) (can be bit-wise OR-ed) to set or clear (e.g. `LV_STATE_CHECKED`)
 * @param ref_value     reference value to compare Subject's value with
 * @return              pointer to newly-created Observer
 */
// llgo:link (*LvObjT).LvObjBindStateIfLt C.lv_obj_bind_state_if_lt
func (recv_ *LvObjT) LvObjBindStateIfLt(subject *LvSubjectT, state LvStateT, ref_value c.Int32T) *LvObserverT {
	return nil
}

/**
 * Set Widget's state(s) if an integer Subject's value is less than or equal to a reference value, clear flag otherwise.
 * @param obj           pointer to Widget
 * @param subject       pointer to Subject
 * @param state         state(s) (can be bit-wise OR-ed) to set or clear (e.g. `LV_STATE_CHECKED`)
 * @param ref_value     reference value to compare Subject's value with
 * @return              pointer to newly-created Observer
 */
// llgo:link (*LvObjT).LvObjBindStateIfLe C.lv_obj_bind_state_if_le
func (recv_ *LvObjT) LvObjBindStateIfLe(subject *LvSubjectT, state LvStateT, ref_value c.Int32T) *LvObserverT {
	return nil
}

/**
 * Set an integer Subject to 1 when a Widget is checked and set it 0 when unchecked, and
 * clear Widget's checked state when Subject's value changes to 0 and set it when non-zero.
 * @param obj       pointer to Widget
 * @param subject   pointer to a Subject
 * @return          pointer to newly-created Observer
 * @note            Ensure Widget's `LV_OBJ_FLAG_CHECKABLE` flag is set.
 */
// llgo:link (*LvObjT).LvObjBindChecked C.lv_obj_bind_checked
func (recv_ *LvObjT) LvObjBindChecked(subject *LvSubjectT) *LvObserverT {
	return nil
}

/**
 * Bind an integer, string, or pointer Subject to a Label.
 * @param obj       pointer to Label
 * @param subject   pointer to Subject
 * @param fmt       optional printf-like format string with 1 format specifier (e.g. "%d °C")
 *                  or NULL to bind to the value directly.
 * @return          pointer to newly-created Observer
 * @note            `fmt == NULL` can be used only with string and pointer Subjects.
 * @note            If Subject is a pointer and `fmt == NULL`, pointer must point
 *                  to a `\0` terminated string.
 */
// llgo:link (*LvObjT).LvLabelBindText C.lv_label_bind_text
func (recv_ *LvObjT) LvLabelBindText(subject *LvSubjectT, fmt *c.Char) *LvObserverT {
	return nil
}

/**
 * Bind an integer subject to an Arc's value.
 * @param obj       pointer to Arc
 * @param subject   pointer to Subject
 * @return          pointer to newly-created Observer
 */
// llgo:link (*LvObjT).LvArcBindValue C.lv_arc_bind_value
func (recv_ *LvObjT) LvArcBindValue(subject *LvSubjectT) *LvObserverT {
	return nil
}

/**
 * Bind an integer Subject to a Slider's value.
 * @param obj       pointer to Slider
 * @param subject   pointer to Subject
 * @return          pointer to newly-created Observer
 */
// llgo:link (*LvObjT).LvSliderBindValue C.lv_slider_bind_value
func (recv_ *LvObjT) LvSliderBindValue(subject *LvSubjectT) *LvObserverT {
	return nil
}

/**
 * Bind an integer Subject to a Roller's value.
 * @param obj       pointer to Roller
 * @param subject   pointer to Subject
 * @return          pointer to newly-created Observer
 */
// llgo:link (*LvObjT).LvRollerBindValue C.lv_roller_bind_value
func (recv_ *LvObjT) LvRollerBindValue(subject *LvSubjectT) *LvObserverT {
	return nil
}

/**
 * Bind an integer Subject to a Dropdown's value.
 * @param obj       pointer to Dropdown
 * @param subject   pointer to Subject
 * @return          pointer to newly-created Observer
 */
// llgo:link (*LvObjT).LvDropdownBindValue C.lv_dropdown_bind_value
func (recv_ *LvObjT) LvDropdownBindValue(subject *LvSubjectT) *LvObserverT {
	return nil
}

type LvTlsfT c.Pointer
type LvPoolT c.Pointer

/* Create/destroy a memory pool. */
//go:linkname LvTlsfCreate C.lv_tlsf_create
func LvTlsfCreate(mem c.Pointer) LvTlsfT

//go:linkname LvTlsfCreateWithPool C.lv_tlsf_create_with_pool
func LvTlsfCreateWithPool(mem c.Pointer, bytes c.SizeT) LvTlsfT

//go:linkname LvTlsfDestroy C.lv_tlsf_destroy
func LvTlsfDestroy(tlsf LvTlsfT)

//go:linkname LvTlsfGetPool C.lv_tlsf_get_pool
func LvTlsfGetPool(tlsf LvTlsfT) LvPoolT

/* Add/remove memory pools. */
//go:linkname LvTlsfAddPool C.lv_tlsf_add_pool
func LvTlsfAddPool(tlsf LvTlsfT, mem c.Pointer, bytes c.SizeT) LvPoolT

//go:linkname LvTlsfRemovePool C.lv_tlsf_remove_pool
func LvTlsfRemovePool(tlsf LvTlsfT, pool LvPoolT)

/* malloc/memalign/realloc/free replacements. */
//go:linkname LvTlsfMalloc C.lv_tlsf_malloc
func LvTlsfMalloc(tlsf LvTlsfT, bytes c.SizeT) c.Pointer

//go:linkname LvTlsfMemalign C.lv_tlsf_memalign
func LvTlsfMemalign(tlsf LvTlsfT, align c.SizeT, bytes c.SizeT) c.Pointer

//go:linkname LvTlsfRealloc C.lv_tlsf_realloc
func LvTlsfRealloc(tlsf LvTlsfT, ptr c.Pointer, size c.SizeT) c.Pointer

//go:linkname LvTlsfFree C.lv_tlsf_free
func LvTlsfFree(tlsf LvTlsfT, ptr c.Pointer) c.SizeT

/* Returns internal block size, not original request size */
//go:linkname LvTlsfBlockSize C.lv_tlsf_block_size
func LvTlsfBlockSize(ptr c.Pointer) c.SizeT

/* Overheads/limits of internal structures. */
//go:linkname LvTlsfSize C.lv_tlsf_size
func LvTlsfSize() c.SizeT

//go:linkname LvTlsfAlignSize C.lv_tlsf_align_size
func LvTlsfAlignSize() c.SizeT

//go:linkname LvTlsfBlockSizeMin C.lv_tlsf_block_size_min
func LvTlsfBlockSizeMin() c.SizeT

//go:linkname LvTlsfBlockSizeMax C.lv_tlsf_block_size_max
func LvTlsfBlockSizeMax() c.SizeT

//go:linkname LvTlsfPoolOverhead C.lv_tlsf_pool_overhead
func LvTlsfPoolOverhead() c.SizeT

//go:linkname LvTlsfAllocOverhead C.lv_tlsf_alloc_overhead
func LvTlsfAllocOverhead() c.SizeT

// llgo:type C
type LvTlsfWalker func(c.Pointer, c.SizeT, c.Int, c.Pointer)

//go:linkname LvTlsfWalkPool C.lv_tlsf_walk_pool
func LvTlsfWalkPool(pool LvPoolT, walker LvTlsfWalker, user c.Pointer)

/* Returns nonzero if any internal consistency check fails. */
//go:linkname LvTlsfCheck C.lv_tlsf_check
func LvTlsfCheck(tlsf LvTlsfT) c.Int

//go:linkname LvTlsfCheckPool C.lv_tlsf_check_pool
func LvTlsfCheckPool(pool LvPoolT) c.Int

type LvTimerStateT struct {
	TimerLl            LvLlT
	LvTimerRun         bool
	IdleLast           c.Uint8T
	TimerDeleted       bool
	TimerCreated       bool
	TimerTimeUntilNext c.Uint32T
	AlreadyRunning     bool
	PeriodicLastTick   c.Uint32T
	BusyTime           c.Uint32T
	IdlePeriodStart    c.Uint32T
	RunCnt             c.Uint32T
	ResumeCb           LvTimerHandlerResumeCbT
	ResumeData         c.Pointer
}

/**
 * Init the lv_timer module
 */
//go:linkname LvTimerCoreInit C.lv_timer_core_init
func LvTimerCoreInit()

/**
 * Deinit the lv_timer module
 */
//go:linkname LvTimerCoreDeinit C.lv_timer_core_deinit
func LvTimerCoreDeinit()

/**********************
 *      TYPEDEFS
 **********************/

type LvAnimStateT struct {
	AnimListChanged     bool
	AnimRunRound        bool
	AnimVsyncRegistered bool
	Timer               *LvTimerT
	AnimLl              LvLlT
}

/**
 * Init the animation module
 */
//go:linkname LvAnimCoreInit C.lv_anim_core_init
func LvAnimCoreInit()

/**
 * Deinit the animation module
 */
//go:linkname LvAnimCoreDeinit C.lv_anim_core_deinit
func LvAnimCoreDeinit()

/*
 * Set animation use vsync mode.
 * @param enable true: use vsync mode, false: use timer mode.
 */
//go:linkname LvAnimEnableVsyncMode C.lv_anim_enable_vsync_mode
func LvAnimEnableVsyncMode(enable bool)

/**********************
 *      TYPEDEFS
 **********************/

type LvTickStateT struct {
	SysTime    c.Uint32T
	SysIrqFlag c.Uint8T
	TickGetCb  LvTickGetCbT
	DelayCb    LvDelayCbT
}

/**
 * Called internally to initialize the draw_buf_handlers in lv_global
 */
//go:linkname LvDrawBufInitHandlers C.lv_draw_buf_init_handlers
func LvDrawBufInitHandlers()

/**
 * Get the size of a cache entry.
 * @param node_size     The size of the node in the cache.
 * @return              The size of the cache entry.
 */
//go:linkname LvCacheEntryGetSize C.lv_cache_entry_get_size
func LvCacheEntryGetSize(node_size c.Uint32T) c.Uint32T

/**
 * Get the reference count of a cache entry.
 * @param entry        The cache entry to get the reference count of.
 * @return             The reference count of the cache entry.
 */
// llgo:link (*LvCacheEntryT).LvCacheEntryGetRef C.lv_cache_entry_get_ref
func (recv_ *LvCacheEntryT) LvCacheEntryGetRef() c.Int32T {
	return 0
}

/**
 * Get the node size of a cache entry. Which is the same size with lv_cache_entry_get_size()'s node_size parameter.
 * @param entry        The cache entry to get the node size of.
 * @return             The node size of the cache entry.
 */
// llgo:link (*LvCacheEntryT).LvCacheEntryGetNodeSize C.lv_cache_entry_get_node_size
func (recv_ *LvCacheEntryT) LvCacheEntryGetNodeSize() c.Uint32T {
	return 0
}

/**
 * Check if a cache entry is invalid.
 * @param entry        The cache entry to check.
 * @return             True: the cache entry is invalid. False: the cache entry is valid.
 */
// llgo:link (*LvCacheEntryT).LvCacheEntryIsInvalid C.lv_cache_entry_is_invalid
func (recv_ *LvCacheEntryT) LvCacheEntryIsInvalid() bool {
	return false
}

/**
 * Get the data of a cache entry.
 * @param entry        The cache entry to get the data of.
 * @return             The pointer to the data of the cache entry.
 */
// llgo:link (*LvCacheEntryT).LvCacheEntryGetData C.lv_cache_entry_get_data
func (recv_ *LvCacheEntryT) LvCacheEntryGetData() c.Pointer {
	return nil
}

/**
 * Get the cache instance of a cache entry.
 * @param entry        The cache entry to get the cache instance of.
 * @return             The pointer to the cache instance of the cache entry.
 */
// llgo:link (*LvCacheEntryT).LvCacheEntryGetCache C.lv_cache_entry_get_cache
func (recv_ *LvCacheEntryT) LvCacheEntryGetCache() *LvCacheT {
	return nil
}

/**
 * Get the cache entry of a data. The data should be allocated by the cache instance.
 * @param data         The data to get the cache entry of.
 * @param node_size    The size of the node in the cache.
 * @return             The pointer to the cache entry of the data.
 */
//go:linkname LvCacheEntryGetEntry C.lv_cache_entry_get_entry
func LvCacheEntryGetEntry(data c.Pointer, node_size c.Uint32T) *LvCacheEntryT

/**
 * Allocate a cache entry.
 * @param node_size    The size of the node in the cache.
 * @param cache        The cache instance to allocate the cache entry from.
 * @return             The pointer to the allocated cache entry.
 */
//go:linkname LvCacheEntryAlloc C.lv_cache_entry_alloc
func LvCacheEntryAlloc(node_size c.Uint32T, cache *LvCacheT) *LvCacheEntryT

/**
 * Initialize a cache entry.
 * @param entry        The cache entry to initialize.
 * @param cache        The cache instance to allocate the cache entry from.
 * @param node_size    The size of the node in the cache.
 */
// llgo:link (*LvCacheEntryT).LvCacheEntryInit C.lv_cache_entry_init
func (recv_ *LvCacheEntryT) LvCacheEntryInit(cache *LvCacheT, node_size c.Uint32T) {
}

/**
 * Deallocate a cache entry. And the data of the cache entry will be freed.
 * @param entry        The cache entry to deallocate.
 */
// llgo:link (*LvCacheEntryT).LvCacheEntryDelete C.lv_cache_entry_delete
func (recv_ *LvCacheEntryT) LvCacheEntryDelete() {
}

type LvCacheReserveCondResT c.Int

const (
	LV_CACHE_RESERVE_COND_OK          LvCacheReserveCondResT = 0
	LV_CACHE_RESERVE_COND_TOO_LARGE   LvCacheReserveCondResT = 1
	LV_CACHE_RESERVE_COND_NEED_VICTIM LvCacheReserveCondResT = 2
	LV_CACHE_RESERVE_COND_ERROR       LvCacheReserveCondResT = 3
)

type X_lvCacheOpsT struct {
	CompareCb LvCacheCompareCbT
	CreateCb  LvCacheCreateCbT
	FreeCb    LvCacheFreeCbT
}

type X_lvCacheClassT struct {
	AllocCb       LvCacheAllocCbT
	InitCb        LvCacheInitCbT
	DestroyCb     LvCacheDestroyCbT
	GetCb         LvCacheGetCbT
	AddCb         LvCacheAddCbT
	RemoveCb      LvCacheRemoveCbT
	DropCb        LvCacheDropCbT
	DropAllCb     LvCacheDropAllCbT
	GetVictimCb   LvCacheGetVictimCb
	ReserveCondCb LvCacheReserveCondCb
	IterCreateCb  LvCacheIterCreateCb
}
type LvCacheOpsT X_lvCacheOpsT
type LvCacheClassT X_lvCacheClassT
type LvCacheCompareResT c.Int8T

// llgo:type C
type LvCacheCreateCbT func(c.Pointer, c.Pointer) bool

// llgo:type C
type LvCacheFreeCbT func(c.Pointer, c.Pointer)

// llgo:type C
type LvCacheCompareCbT func(c.Pointer, c.Pointer) LvCacheCompareResT

// llgo:type C
type LvCacheAllocCbT func() c.Pointer

// llgo:type C
type LvCacheInitCbT func(*LvCacheT) bool

// llgo:type C
type LvCacheDestroyCbT func(*LvCacheT, c.Pointer)

// llgo:type C
type LvCacheGetCbT func(*LvCacheT, c.Pointer, c.Pointer) *LvCacheEntryT

// llgo:type C
type LvCacheAddCbT func(*LvCacheT, c.Pointer, c.Pointer) *LvCacheEntryT

// llgo:type C
type LvCacheRemoveCbT func(*LvCacheT, *LvCacheEntryT, c.Pointer)

// llgo:type C
type LvCacheDropCbT func(*LvCacheT, c.Pointer, c.Pointer)

// llgo:type C
type LvCacheDropAllCbT func(*LvCacheT, c.Pointer)

// llgo:type C
type LvCacheGetVictimCb func(*LvCacheT, c.Pointer) *LvCacheEntryT

// llgo:type C
type LvCacheReserveCondCb func(*LvCacheT, c.Pointer, c.SizeT, c.Pointer) LvCacheReserveCondResT

// llgo:type C
type LvCacheIterCreateCb func(*LvCacheT) *LvIterT

/*-----------------
 * Cache entry slot
 *----------------*/

type X_lvCacheSlotSizeT struct {
	Size c.SizeT
}
type LvCacheSlotSizeT X_lvCacheSlotSizeT

/**
 * Initialize image header cache.
 * @param  count initial size of the cache in count of image headers.
 * @return LV_RESULT_OK: initialization succeeded, LV_RESULT_INVALID: failed.
 */
//go:linkname LvImageHeaderCacheInit C.lv_image_header_cache_init
func LvImageHeaderCacheInit(count c.Uint32T) LvResultT

/**
 * Resize image header cache.
 * If set to 0, the cache is disabled.
 * @param count  new max count of cached image headers.
 * @param evict_now true: evict the image headers should be removed by the eviction policy, false: wait for the next cache cleanup.
 */
//go:linkname LvImageHeaderCacheResize C.lv_image_header_cache_resize
func LvImageHeaderCacheResize(count c.Uint32T, evict_now bool)

/**
 * Invalidate image header cache. Use NULL to invalidate all image headers.
 * It's also automatically called when an image is invalidated.
 * @param src pointer to an image source.
 */
//go:linkname LvImageHeaderCacheDrop C.lv_image_header_cache_drop
func LvImageHeaderCacheDrop(src c.Pointer)

/**
 * Return true if the image header cache is enabled.
 * @return true: enabled, false: disabled.
 */
//go:linkname LvImageHeaderCacheIsEnabled C.lv_image_header_cache_is_enabled
func LvImageHeaderCacheIsEnabled() bool

/**
 * Create an iterator to iterate over the image header cache.
 * @return an iterator to iterate over the image header cache.
 */
//go:linkname LvImageHeaderCacheIterCreate C.lv_image_header_cache_iter_create
func LvImageHeaderCacheIterCreate() *LvIterT

/**
 * Dump the content of the image header cache in a human-readable format with cache order.
 */
//go:linkname LvImageHeaderCacheDump C.lv_image_header_cache_dump
func LvImageHeaderCacheDump()

/**
 * Initialize image cache.
 * @param  size size of the cache in bytes.
 * @return LV_RESULT_OK: initialization succeeded, LV_RESULT_INVALID: failed.
 */
//go:linkname LvImageCacheInit C.lv_image_cache_init
func LvImageCacheInit(size c.Uint32T) LvResultT

/**
 * Resize image cache.
 * If set to 0, the cache will be disabled.
 * @param new_size  new size of the cache in bytes.
 * @param evict_now true: evict the images should be removed by the eviction policy, false: wait for the next cache cleanup.
 */
//go:linkname LvImageCacheResize C.lv_image_cache_resize
func LvImageCacheResize(new_size c.Uint32T, evict_now bool)

/**
 * Invalidate image cache. Use NULL to invalidate all images.
 * @param src pointer to an image source.
 */
//go:linkname LvImageCacheDrop C.lv_image_cache_drop
func LvImageCacheDrop(src c.Pointer)

/**
 * Return true if the image cache is enabled.
 * @return true: enabled, false: disabled.
 */
//go:linkname LvImageCacheIsEnabled C.lv_image_cache_is_enabled
func LvImageCacheIsEnabled() bool

/**
 * Create an iterator to iterate over the image cache.
 * @return an iterator to iterate over the image cache.
 */
//go:linkname LvImageCacheIterCreate C.lv_image_cache_iter_create
func LvImageCacheIterCreate() *LvIterT

/**
 * Dump the content of the image cache in a human-readable format with cache order.
 */
//go:linkname LvImageCacheDump C.lv_image_cache_dump
func LvImageCacheDump()

/**
 * Create a cache object with the given parameters.
 * @param cache_class   The class of the cache. Currently only support one two builtin classes:
 *                        - lv_cache_class_lru_rb_count for LRU-based cache with count-based eviction policy.
 *                        - lv_cache_class_lru_rb_size for LRU-based cache with size-based eviction policy.
 * @param node_size     The node size is the size of the data stored in the cache..
 * @param max_size      The max size is the maximum amount of memory or count that the cache can hold.
 *                        - lv_cache_class_lru_rb_count: max_size is the maximum count of nodes in the cache.
 *                        - lv_cache_class_lru_rb_size: max_size is the maximum size of the cache in bytes.
 * @param ops           A set of operations that can be performed on the cache. See lv_cache_ops_t for details.
 * @return              Returns a pointer to the created cache object on success, `NULL` on error.
 */
// llgo:link (*LvCacheClassT).LvCacheCreate C.lv_cache_create
func (recv_ *LvCacheClassT) LvCacheCreate(node_size c.SizeT, max_size c.SizeT, ops LvCacheOpsT) *LvCacheT {
	return nil
}

/**
 * Destroy a cache object.
 * @param cache         The cache object pointer to destroy.
 * @param user_data     A user data pointer that will be passed to the free callback.
 */
// llgo:link (*LvCacheT).LvCacheDestroy C.lv_cache_destroy
func (recv_ *LvCacheT) LvCacheDestroy(user_data c.Pointer) {
}

/**
 * Acquire a cache entry with the given key. If entry not in cache, it will return `NULL` (not found).
 * If the entry is found, it's priority will be changed by the cache's policy. And the `lv_cache_entry_t::ref_cnt` will be incremented.
 * @param cache         The cache object pointer to acquire the entry.
 * @param key           The key of the entry to acquire.
 * @param user_data     A user data pointer that will be passed to the create callback.
 * @return              Returns a pointer to the acquired cache entry on success with `lv_cache_entry_t::ref_cnt` incremented, `NULL` on error.
 */
// llgo:link (*LvCacheT).LvCacheAcquire C.lv_cache_acquire
func (recv_ *LvCacheT) LvCacheAcquire(key c.Pointer, user_data c.Pointer) *LvCacheEntryT {
	return nil
}

/**
 * Acquire a cache entry with the given key. If the entry is not in the cache, it will create a new entry with the given key.
 * If the entry is found, it's priority will be changed by the cache's policy. And the `lv_cache_entry_t::ref_cnt` will be incremented.
 * If you want to use this API to simplify the code, you should provide a `lv_cache_ops_t::create_cb` that creates a new entry with the given key.
 * This API is a combination of lv_cache_acquire() and lv_cache_add(). The effect is the same as calling lv_cache_acquire() and lv_cache_add() separately.
 * And the internal impact on cache is also consistent with these two APIs.
 * @param cache         The cache object pointer to acquire the entry.
 * @param key           The key of the entry to acquire or create.
 * @param user_data     A user data pointer that will be passed to the create callback.
 * @return              Returns a pointer to the acquired or created cache entry on success with `lv_cache_entry_t::ref_cnt` incremented, `NULL` on error.
 */
// llgo:link (*LvCacheT).LvCacheAcquireOrCreate C.lv_cache_acquire_or_create
func (recv_ *LvCacheT) LvCacheAcquireOrCreate(key c.Pointer, user_data c.Pointer) *LvCacheEntryT {
	return nil
}

/**
 * Add a new cache entry with the given key and data. If the cache is full, the cache's policy will be used to evict an entry.
 * @param cache         The cache object pointer to add the entry.
 * @param key           The key of the entry to add.
 * @param user_data     A user data pointer that will be passed to the create callback.
 * @return              Returns a pointer to the added cache entry on success with `lv_cache_entry_t::ref_cnt` incremented, `NULL` on error.
 */
// llgo:link (*LvCacheT).LvCacheAdd C.lv_cache_add
func (recv_ *LvCacheT) LvCacheAdd(key c.Pointer, user_data c.Pointer) *LvCacheEntryT {
	return nil
}

/**
 * Release a cache entry. The `lv_cache_entry_t::ref_cnt` will be decremented. If the `lv_cache_entry_t::ref_cnt` is zero, it will issue an error.
 * If the entry passed to this function is the last reference to the data and the entry is marked as invalid, the cache's policy will be used to evict the entry.
 * @param cache         The cache object pointer to release the entry.
 * @param entry         The cache entry pointer to release.
 * @param user_data     A user data pointer that will be passed to the free callback.
 */
// llgo:link (*LvCacheT).LvCacheRelease C.lv_cache_release
func (recv_ *LvCacheT) LvCacheRelease(entry *LvCacheEntryT, user_data c.Pointer) {
}

/**
 * Reserve a certain amount of memory/count in the cache. This function is useful when you want to reserve a certain amount of memory/count in advance,
 * for example, when you know that you will need it later.
 * When the current cache size is max than the reserved size, the function will evict entries until the reserved size is reached.
 * @param cache         The cache object pointer to reserve.
 * @param reserved_size The amount of memory/count to reserve.
 * @param user_data     A user data pointer that will be passed to the free callback.
 */
// llgo:link (*LvCacheT).LvCacheReserve C.lv_cache_reserve
func (recv_ *LvCacheT) LvCacheReserve(reserved_size c.Uint32T, user_data c.Pointer) {
}

/**
 * Drop a cache entry with the given key. If the entry is not in the cache, nothing will happen to it.
 * If the entry is found, it will be removed from the cache and its data will be freed when the last reference to it is released.
 * @note The data will not be freed immediately but when the last reference to it is released. But this entry will not be found by lv_cache_acquire().
 *       If you want cache a same key again, you should use lv_cache_add() or lv_cache_acquire_or_create().
 * @param cache         The cache object pointer to drop the entry.
 * @param key           The key of the entry to drop.
 * @param user_data     A user data pointer that will be passed to the free callback.
 */
// llgo:link (*LvCacheT).LvCacheDrop C.lv_cache_drop
func (recv_ *LvCacheT) LvCacheDrop(key c.Pointer, user_data c.Pointer) {
}

/**
 * Drop all cache entries. All entries will be removed from the cache and their data will be freed when the last reference to them is released.
 * @note If some entries are still referenced by other objects, it will issue an error. And this case shouldn't happen in normal cases..
 * @param cache         The cache object pointer to drop all entries.
 * @param user_data     A user data pointer that will be passed to the free callback.
 */
// llgo:link (*LvCacheT).LvCacheDropAll C.lv_cache_drop_all
func (recv_ *LvCacheT) LvCacheDropAll(user_data c.Pointer) {
}

/**
 * Evict one entry from the cache. The eviction policy will be used to select the entry to evict.
 * @param cache         The cache object pointer to evict an entry.
 * @param user_data     A user data pointer that will be passed to the free callback.
 * @return              Returns true if an entry is evicted, false if no entry is evicted.
 */
// llgo:link (*LvCacheT).LvCacheEvictOne C.lv_cache_evict_one
func (recv_ *LvCacheT) LvCacheEvictOne(user_data c.Pointer) bool {
	return false
}

/**
 * Set the maximum size of the cache.
 * If the current cache size is greater than the new maximum size, the cache's policy will be used to evict entries until the new maximum size is reached.
 * If set to 0, the cache will be disabled.
 * @note But this behavior will happen only new entries are added to the cache.
 * @param cache         The cache object pointer to set the maximum size.
 * @param max_size      The new maximum size of the cache.
 * @param user_data     A user data pointer that will be passed to the free callback.
 */
// llgo:link (*LvCacheT).LvCacheSetMaxSize C.lv_cache_set_max_size
func (recv_ *LvCacheT) LvCacheSetMaxSize(max_size c.SizeT, user_data c.Pointer) {
}

/**
 * Get the maximum size of the cache.
 * @param cache         The cache object pointer to get the maximum size.
 * @param user_data     A user data pointer that will be passed to the free callback.
 * @return              Returns the maximum size of the cache.
 */
// llgo:link (*LvCacheT).LvCacheGetMaxSize C.lv_cache_get_max_size
func (recv_ *LvCacheT) LvCacheGetMaxSize(user_data c.Pointer) c.SizeT {
	return 0
}

/**
 * Get the current size of the cache.
 * @param cache         The cache object pointer to get the current size.
 * @param user_data     A user data pointer that will be passed to the free callback.
 * @return              Returns the current size of the cache.
 */
// llgo:link (*LvCacheT).LvCacheGetSize C.lv_cache_get_size
func (recv_ *LvCacheT) LvCacheGetSize(user_data c.Pointer) c.SizeT {
	return 0
}

/**
 * Get the free size of the cache.
 * @param cache         The cache object pointer to get the free size.
 * @param user_data     A user data pointer that will be passed to the free callback.
 * @return              Returns the free size of the cache.
 */
// llgo:link (*LvCacheT).LvCacheGetFreeSize C.lv_cache_get_free_size
func (recv_ *LvCacheT) LvCacheGetFreeSize(user_data c.Pointer) c.SizeT {
	return 0
}

/**
 * Return true if the cache is enabled.
 * Disabled cache means that when the max_size of the cache is 0. In this case, all cache operations will be no-op.
 * @param cache         The cache object pointer to check if it's disabled.
 * @return              Returns true if the cache is enabled, false otherwise.
 */
// llgo:link (*LvCacheT).LvCacheIsEnabled C.lv_cache_is_enabled
func (recv_ *LvCacheT) LvCacheIsEnabled() bool {
	return false
}

/**
 * Set the compare callback of the cache.
 * @param cache         The cache object pointer to set the compare callback.
 * @param compare_cb    The compare callback to set.
 * @param user_data     A user data pointer.
 */
// llgo:link (*LvCacheT).LvCacheSetCompareCb C.lv_cache_set_compare_cb
func (recv_ *LvCacheT) LvCacheSetCompareCb(compare_cb LvCacheCompareCbT, user_data c.Pointer) {
}

/**
 * Set the create callback of the cache.
 * @param cache         The cache object pointer to set the create callback.
 * @param alloc_cb      The create callback to set.
 * @param user_data     A user data pointer.
 */
// llgo:link (*LvCacheT).LvCacheSetCreateCb C.lv_cache_set_create_cb
func (recv_ *LvCacheT) LvCacheSetCreateCb(alloc_cb LvCacheCreateCbT, user_data c.Pointer) {
}

/**
 * Set the free callback of the cache.
 * @param cache         The cache object pointer to set the free callback.
 * @param free_cb       The free callback to set.
 * @param user_data     A user data pointer.
 */
// llgo:link (*LvCacheT).LvCacheSetFreeCb C.lv_cache_set_free_cb
func (recv_ *LvCacheT) LvCacheSetFreeCb(free_cb LvCacheFreeCbT, user_data c.Pointer) {
}

/**
 * Give a name for a cache object. Only the pointer of the string is saved.
 * @param cache         The cache object pointer to set the name.
 * @param name          The name of the cache.
 */
// llgo:link (*LvCacheT).LvCacheSetName C.lv_cache_set_name
func (recv_ *LvCacheT) LvCacheSetName(name *c.Char) {
}

/**
 * Get the name of a cache object.
 * @param cache         The cache object pointer to get the name.
 * @return              Returns the name of the cache.
 */
// llgo:link (*LvCacheT).LvCacheGetName C.lv_cache_get_name
func (recv_ *LvCacheT) LvCacheGetName() *c.Char {
	return nil
}

/**
 * Create an iterator for the cache object. The iterator is used to iterate over all cache entries.
 * @param cache         The cache object pointer to create the iterator.
 * @return              Returns a pointer to the created iterator on success, `NULL` on error.
 */
// llgo:link (*LvCacheT).LvCacheIterCreate C.lv_cache_iter_create
func (recv_ *LvCacheT) LvCacheIterCreate() *LvIterT {
	return nil
}

type LvDrawGlobalInfoT struct {
	UnitHead            *LvDrawUnitT
	UnitCnt             c.Uint32T
	UsedMemoryForLayers c.Uint32T
	DispatchReq         c.Int
	CircleCacheMutex    LvMutexT
	TaskRunning         bool
}

/**********************
 *      TYPEDEFS
 **********************/

type LvDrawSwThreadDscT struct {
	TaskAct    *LvDrawTaskT
	Thread     LvThreadT
	Sync       LvThreadSyncT
	DrawUnit   *LvDrawUnitT
	Idx        c.Uint32T
	Inited     bool
	ExitStatus bool
}

/**********************
 *      TYPEDEFS
 **********************/

type LvDrawSwMaskRadiusCircleDscT struct {
	Buf         *c.Uint8T
	CirOpa      *LvOpaT
	XStartOnY   *c.Uint16T
	OpaStartOnY *c.Uint16T
	Life        c.Int32T
	UsedCnt     c.Uint32T
	Radius      c.Int32T
}
type LvDrawSwMaskRadiusCircleDscArrT [4]LvDrawSwMaskRadiusCircleDscT

/**
 * Called by LVGL the rendering of a screen is ready to clean up
 * the temporal (cache) data of the masks
 */
//go:linkname LvDrawSwMaskCleanup C.lv_draw_sw_mask_cleanup
func LvDrawSwMaskCleanup()

/**********************
 *      TYPEDEFS
 **********************/

type LvTlsfStateT struct {
	Tlsf    LvTlsfT
	CurUsed c.SizeT
	MaxUsed c.SizeT
	PoolLl  LvLlT
}

/**********************
 *      TYPEDEFS
 **********************/

type LvLayoutDscT struct {
	Cb       LvLayoutUpdateCbT
	UserData c.Pointer
}

/**********************
 * GLOBAL PROTOTYPES
 **********************/
//go:linkname LvLayoutInit C.lv_layout_init
func LvLayoutInit()

//go:linkname LvLayoutDeinit C.lv_layout_deinit
func LvLayoutDeinit()

/**
 * Update the layout of a widget
 * @param obj   pointer to a widget
 */
// llgo:link (*LvObjT).LvLayoutApply C.lv_layout_apply
func (recv_ *LvObjT) LvLayoutApply() {
}

/**********************
 *      TYPEDEFS
 **********************/

type X_snippetStack struct {
	Unused [8]uint8
}

type X_lvGlobalT struct {
	Inited                         bool
	DeinitInProgress               bool
	DispLl                         LvLlT
	DispRefresh                    *LvDisplayT
	DispDefault                    *LvDisplayT
	StyleTransLl                   LvLlT
	StyleRefresh                   bool
	StyleCustomTableSize           c.Uint32T
	StyleLastCustomPropId          c.Uint32T
	StyleCustomPropFlagLookupTable *c.Uint8T
	GroupLl                        LvLlT
	GroupDefault                   *LvGroupT
	IndevLl                        LvLlT
	IndevActive                    *LvIndevT
	IndevObjActive                 *LvObjT
	LayoutCount                    c.Uint32T
	LayoutList                     *LvLayoutDscT
	LayoutUpdateMutex              bool
	MemoryZero                     c.Uint32T
	MathRandSeed                   c.Uint32T
	EventHeader                    *LvEventT
	EventLastRegisterId            c.Uint32T
	TimerState                     LvTimerStateT
	AnimState                      LvAnimStateT
	TickState                      LvTickStateT
	DrawBufHandlers                LvDrawBufHandlersT
	FontDrawBufHandlers            LvDrawBufHandlersT
	ImageCacheDrawBufHandlers      LvDrawBufHandlersT
	ImgDecoderLl                   LvLlT
	ImgCache                       *LvCacheT
	ImgHeaderCache                 *LvCacheT
	DrawInfo                       LvDrawGlobalInfoT
	DrawSwBlendHandlerLl           LvLlT
	SwCircleCache                  LvDrawSwMaskRadiusCircleDscArrT
	ThemeSimple                    c.Pointer
	ThemeDefault                   c.Pointer
	ThemeMono                      c.Pointer
	TlsfState                      LvTlsfStateT
	FsdrvLl                        LvLlT
	SpanSnippetStack               *X_snippetStack
	ObjidArray                     c.Pointer
	ObjidCount                     c.Uint32T
	UserData                       c.Pointer
}
type LvGlobalT X_lvGlobalT

/**
 * Handle scrolling. Called by LVGL during input device processing
 * @param indev      pointer to an input device
 */
// llgo:link (*LvIndevT).LvIndevScrollHandler C.lv_indev_scroll_handler
func (recv_ *LvIndevT) LvIndevScrollHandler() {
}

/**
 * Handle throwing after scrolling. Called by LVGL during input device processing
 * @param indev      pointer to an input device
 */
// llgo:link (*LvIndevT).LvIndevScrollThrowHandler C.lv_indev_scroll_throw_handler
func (recv_ *LvIndevT) LvIndevScrollThrowHandler() {
}

/**
 * Predict where would a scroll throw end
 * @param indev     pointer to an input device
 * @param dir `     LV_DIR_VER` or `LV_DIR_HOR`
 * @return          the difference compared to the current position when the throw would be finished
 */
// llgo:link (*LvIndevT).LvIndevScrollThrowPredict C.lv_indev_scroll_throw_predict
func (recv_ *LvIndevT) LvIndevScrollThrowPredict(dir LvDirT) c.Int32T {
	return 0
}

/**
 * Get the distance of the nearest snap point
 * @param obj       the object on which snap points should be found
 * @param p         save the distance of the found snap point there
 */
// llgo:link (*LvObjT).LvIndevScrollGetSnapDist C.lv_indev_scroll_get_snap_dist
func (recv_ *LvObjT) LvIndevScrollGetSnapDist(p *LvPointT) {
}

/**
 * Find a scrollable object based on the current scroll vector in the indev.
 * In handles scroll propagation to the parent if needed, and scroll directions too.
 * @param indev     pointer to an indev
 * @return          the found scrollable object or NULL if not found.
 */
// llgo:link (*LvIndevT).LvIndevFindScrollObj C.lv_indev_find_scroll_obj
func (recv_ *LvIndevT) LvIndevFindScrollObj() *LvObjT {
	return nil
}

/**
 * Get the next line of text. Check line length and break chars too.
 * @param txt a '\0' terminated string
 * @param len length of 'txt' in bytes
 * @param font pointer to a font
 * @param letter_space letter space
 * @param max_width max width of the text (break the lines to fit this size). Set COORD_MAX to avoid
 * line breaks
 * @param used_width When used_width != NULL, save the width of this line if
 * flag == LV_TEXT_FLAG_NONE, otherwise save -1.
 * @param flag settings for the text from 'txt_flag_type' enum
 * @return the index of the first char of the new line (in byte index not letter index. With UTF-8
 * they are different)
 */
//go:linkname LvTextGetNextLine C.lv_text_get_next_line
func LvTextGetNextLine(txt *c.Char, len c.Uint32T, font *LvFontT, letter_space c.Int32T, max_width c.Int32T, used_width *c.Int32T, flag LvTextFlagT) c.Uint32T

/**
 * Insert a string into another
 * @param txt_buf the original text (must be big enough for the result text and NULL terminated)
 * @param pos position to insert (0: before the original text, 1: after the first char etc.)
 * @param ins_txt text to insert, must be '\0' terminated
 */
//go:linkname LvTextIns C.lv_text_ins
func LvTextIns(txt_buf *c.Char, pos c.Uint32T, ins_txt *c.Char)

/**
 * Delete a part of a string
 * @param txt string to modify, must be '\0' terminated and should point to a heap or stack frame, not read-only memory.
 * @param pos position where to start the deleting (0: before the first char, 1: after the first
 * char etc.)
 * @param len number of characters to delete
 */
//go:linkname LvTextCut C.lv_text_cut
func LvTextCut(txt *c.Char, pos c.Uint32T, len c.Uint32T)

/**
 * return a new formatted text. Memory will be allocated to store the text.
 * @param fmt `printf`-like format
 * @param ap items to print

 * @return pointer to the allocated text string.
 */
//go:linkname LvTextSetTextVfmt C.lv_text_set_text_vfmt
func LvTextSetTextVfmt(fmt *c.Char, ap c.VaList) *c.Char

/**
 * Decode two encoded character from a string.
 * @param txt pointer to '\0' terminated string
 * @param letter the first decoded Unicode character or 0 on invalid data code
 * @param letter_next the second decoded Unicode character or 0 on invalid data code
 * @param ofs start index in 'txt' where to start.
 *                After the call it will point to the next encoded char in 'txt'.
 *                NULL to use txt[0] as index
 */
//go:linkname LvTextEncodedLetterNext2 C.lv_text_encoded_letter_next_2
func LvTextEncodedLetterNext2(txt *c.Char, letter *c.Uint32T, letter_next *c.Uint32T, ofs *c.Uint32T)

/**********************
 * GLOBAL PROTOTYPES
 **********************/
// llgo:link (*LvCacheEntryT).LvCacheEntryResetRef C.lv_cache_entry_reset_ref
func (recv_ *LvCacheEntryT) LvCacheEntryResetRef() {
}

// llgo:link (*LvCacheEntryT).LvCacheEntryIncRef C.lv_cache_entry_inc_ref
func (recv_ *LvCacheEntryT) LvCacheEntryIncRef() {
}

// llgo:link (*LvCacheEntryT).LvCacheEntryDecRef C.lv_cache_entry_dec_ref
func (recv_ *LvCacheEntryT) LvCacheEntryDecRef() {
}

// llgo:link (*LvCacheEntryT).LvCacheEntrySetNodeSize C.lv_cache_entry_set_node_size
func (recv_ *LvCacheEntryT) LvCacheEntrySetNodeSize(node_size c.Uint32T) {
}

// llgo:link (*LvCacheEntryT).LvCacheEntrySetInvalid C.lv_cache_entry_set_invalid
func (recv_ *LvCacheEntryT) LvCacheEntrySetInvalid(is_invalid bool) {
}

// llgo:link (*LvCacheEntryT).LvCacheEntrySetCache C.lv_cache_entry_set_cache
func (recv_ *LvCacheEntryT) LvCacheEntrySetCache(cache *LvCacheT) {
}

// llgo:link (*LvCacheEntryT).LvCacheEntryAcquireData C.lv_cache_entry_acquire_data
func (recv_ *LvCacheEntryT) LvCacheEntryAcquireData() c.Pointer {
	return nil
}

// llgo:link (*LvCacheEntryT).LvCacheEntryReleaseData C.lv_cache_entry_release_data
func (recv_ *LvCacheEntryT) LvCacheEntryReleaseData(user_data c.Pointer) {
}

type LvImageAlignT c.Int

const (
	LV_IMAGE_ALIGN_DEFAULT        LvImageAlignT = 0
	LV_IMAGE_ALIGN_TOP_LEFT       LvImageAlignT = 1
	LV_IMAGE_ALIGN_TOP_MID        LvImageAlignT = 2
	LV_IMAGE_ALIGN_TOP_RIGHT      LvImageAlignT = 3
	LV_IMAGE_ALIGN_BOTTOM_LEFT    LvImageAlignT = 4
	LV_IMAGE_ALIGN_BOTTOM_MID     LvImageAlignT = 5
	LV_IMAGE_ALIGN_BOTTOM_RIGHT   LvImageAlignT = 6
	LV_IMAGE_ALIGN_LEFT_MID       LvImageAlignT = 7
	LV_IMAGE_ALIGN_RIGHT_MID      LvImageAlignT = 8
	LV_IMAGE_ALIGN_CENTER         LvImageAlignT = 9
	LV_IMAGE_ALIGN_AUTO_TRANSFORM LvImageAlignT = 10
	LV_IMAGE_ALIGN_STRETCH        LvImageAlignT = 11
	LV_IMAGE_ALIGN_TILE           LvImageAlignT = 12
	LV_IMAGE_ALIGN_CONTAIN        LvImageAlignT = 13
	LV_IMAGE_ALIGN_COVER          LvImageAlignT = 14
)

/**
 * Create an image object
 * @param parent pointer to an object, it will be the parent of the new image
 * @return pointer to the created image
 */
// llgo:link (*LvObjT).LvImageCreate C.lv_image_create
func (recv_ *LvObjT) LvImageCreate() *LvObjT {
	return nil
}

/**
 * Set the image data to display on the object
 * @param obj       pointer to an image object
 * @param src       1) pointer to an ::lv_image_dsc_t descriptor (converted by LVGL's image converter) (e.g. &my_img) or
 *                  2) path to an image file (e.g. "S:/dir/img.bin")or
 *                  3) a SYMBOL (e.g. LV_SYMBOL_OK)
 */
// llgo:link (*LvObjT).LvImageSetSrc C.lv_image_set_src
func (recv_ *LvObjT) LvImageSetSrc(src c.Pointer) {
}

/**
 * Set an offset for the source of an image so the image will be displayed from the new origin.
 * @param obj       pointer to an image
 * @param x         the new offset along x axis.
 */
// llgo:link (*LvObjT).LvImageSetOffsetX C.lv_image_set_offset_x
func (recv_ *LvObjT) LvImageSetOffsetX(x c.Int32T) {
}

/**
 * Set an offset for the source of an image.
 * so the image will be displayed from the new origin.
 * @param obj       pointer to an image
 * @param y         the new offset along y axis.
 */
// llgo:link (*LvObjT).LvImageSetOffsetY C.lv_image_set_offset_y
func (recv_ *LvObjT) LvImageSetOffsetY(y c.Int32T) {
}

/**
 * Set the rotation angle of the image.
 * The image will be rotated around the set pivot set by `lv_image_set_pivot()`
 * Note that indexed and alpha only images can't be transformed.
 * @param obj       pointer to an image object
 * @param angle     rotation in degree with 0.1 degree resolution (0..3600: clock wise)
 * @note            if image_align is `LV_IMAGE_ALIGN_STRETCH` or `LV_IMAGE_ALIGN_FIT`
 *                  rotation will be set to 0 automatically.
 *
 */
// llgo:link (*LvObjT).LvImageSetRotation C.lv_image_set_rotation
func (recv_ *LvObjT) LvImageSetRotation(angle c.Int32T) {
}

/**
 * Set the rotation center of the image.
 * The image will be rotated around this point.
 * x, y can be set with value of LV_PCT, lv_image_get_pivot will return the true pixel coordinate of pivot in this case.
 * @param obj       pointer to an image object
 * @param x         rotation center x of the image
 * @param y         rotation center y of the image
 */
// llgo:link (*LvObjT).LvImageSetPivot C.lv_image_set_pivot
func (recv_ *LvObjT) LvImageSetPivot(x c.Int32T, y c.Int32T) {
}

/**
 * Set the rotation horizontal center of the image.
 * @param obj       pointer to an image object
 * @param x         rotation center x of the image, or lv_pct()
 */
// llgo:link (*LvObjT).LvImageSetPivotX C.lv_image_set_pivot_x
func (recv_ *LvObjT) LvImageSetPivotX(x c.Int32T) {
}

/**
 * Set the rotation vertical center of the image.
 * @param obj       pointer to an image object
 * @param y         rotation center y of the image, or lv_pct()
 */
// llgo:link (*LvObjT).LvImageSetPivotY C.lv_image_set_pivot_y
func (recv_ *LvObjT) LvImageSetPivotY(y c.Int32T) {
}

/**
 * Set the zoom factor of the image.
 * Note that indexed and alpha only images can't be transformed.
 * @param obj       pointer to an image object
 * @param zoom      the zoom factor.  Example values:
 *                      - 256 or LV_ZOOM_IMAGE_NONE:  no zoom
 *                      - <256:  scale down
 *                      - >256:  scale up
 *                      - 128:  half size
 *                      - 512:  double size
 */
// llgo:link (*LvObjT).LvImageSetScale C.lv_image_set_scale
func (recv_ *LvObjT) LvImageSetScale(zoom c.Uint32T) {
}

/**
 * Set the horizontal zoom factor of the image.
 * Note that indexed and alpha only images can't be transformed.
 * @param obj       pointer to an image object
 * @param zoom      the zoom factor.  Example values:
 *                      - 256 or LV_ZOOM_IMAGE_NONE:  no zoom
 *                      - <256:  scale down
 *                      - >256:  scale up
 *                      - 128:  half size
 *                      - 512:  double size
 */
// llgo:link (*LvObjT).LvImageSetScaleX C.lv_image_set_scale_x
func (recv_ *LvObjT) LvImageSetScaleX(zoom c.Uint32T) {
}

/**
 * Set the vertical zoom factor of the image.
 * Note that indexed and alpha only images can't be transformed.
 * @param obj       pointer to an image object
 * @param zoom      the zoom factor.  Example values:
 *                      - 256 or LV_ZOOM_IMAGE_NONE:  no zoom
 *                      - <256:  scale down
 *                      - >256:  scale up
 *                      - 128:  half size
 *                      - 512:  double size
 */
// llgo:link (*LvObjT).LvImageSetScaleY C.lv_image_set_scale_y
func (recv_ *LvObjT) LvImageSetScaleY(zoom c.Uint32T) {
}

/**
 * Set the blend mode of an image.
 * @param obj           pointer to an image object
 * @param blend_mode    the new blend mode
 */
// llgo:link (*LvObjT).LvImageSetBlendMode C.lv_image_set_blend_mode
func (recv_ *LvObjT) LvImageSetBlendMode(blend_mode LvBlendModeT) {
}

/**
 * Enable/disable anti-aliasing for the transformations (rotate, zoom) or not.
 * The quality is better with anti-aliasing looks better but slower.
 * @param obj       pointer to an image object
 * @param antialias true: anti-aliased; false: not anti-aliased
 */
// llgo:link (*LvObjT).LvImageSetAntialias C.lv_image_set_antialias
func (recv_ *LvObjT) LvImageSetAntialias(antialias bool) {
}

/**
 * Set the image object size mode.
 * @param obj       pointer to an image object
 * @param align     the new align mode.
 * @note            if image_align is `LV_IMAGE_ALIGN_STRETCH` or `LV_IMAGE_ALIGN_FIT`
 *                  rotation, scale and pivot will be overwritten and controlled internally.
 */
// llgo:link (*LvObjT).LvImageSetInnerAlign C.lv_image_set_inner_align
func (recv_ *LvObjT) LvImageSetInnerAlign(align LvImageAlignT) {
}

/**
 * Set an A8 bitmap mask for the image.
 * @param obj       pointer to an image object
 * @param src       an lv_image_dsc_t bitmap mask source.
 */
// llgo:link (*LvObjT).LvImageSetBitmapMapSrc C.lv_image_set_bitmap_map_src
func (recv_ *LvObjT) LvImageSetBitmapMapSrc(src *LvImageDscT) {
}

/**
 * Get the source of the image
 * @param obj       pointer to an image object
 * @return          the image source (symbol, file name or ::lv-img_dsc_t for C arrays)
 */
// llgo:link (*LvObjT).LvImageGetSrc C.lv_image_get_src
func (recv_ *LvObjT) LvImageGetSrc() c.Pointer {
	return nil
}

/**
 * Get the offset's x attribute of the image object.
 * @param obj       pointer to an image
 * @return          offset X value.
 */
// llgo:link (*LvObjT).LvImageGetOffsetX C.lv_image_get_offset_x
func (recv_ *LvObjT) LvImageGetOffsetX() c.Int32T {
	return 0
}

/**
 * Get the offset's y attribute of the image object.
 * @param obj       pointer to an image
 * @return          offset Y value.
 */
// llgo:link (*LvObjT).LvImageGetOffsetY C.lv_image_get_offset_y
func (recv_ *LvObjT) LvImageGetOffsetY() c.Int32T {
	return 0
}

/**
 * Get the rotation of the image.
 * @param obj       pointer to an image object
 * @return          rotation in 0.1 degrees (0..3600)
 * @note            if image_align is `LV_IMAGE_ALIGN_STRETCH` or  `LV_IMAGE_ALIGN_FIT`
 *                  rotation will be set to 0 automatically.
 */
// llgo:link (*LvObjT).LvImageGetRotation C.lv_image_get_rotation
func (recv_ *LvObjT) LvImageGetRotation() c.Int32T {
	return 0
}

/**
 * Get the pivot (rotation center) of the image.
 * If pivot is set with LV_PCT, convert it to px before return.
 * @param obj       pointer to an image object
 * @param pivot     store the rotation center here
 */
// llgo:link (*LvObjT).LvImageGetPivot C.lv_image_get_pivot
func (recv_ *LvObjT) LvImageGetPivot(pivot *LvPointT) {
}

/**
 * Get the zoom factor of the image.
 * @param obj       pointer to an image object
 * @return          zoom factor (256: no zoom)
 */
// llgo:link (*LvObjT).LvImageGetScale C.lv_image_get_scale
func (recv_ *LvObjT) LvImageGetScale() c.Int32T {
	return 0
}

/**
 * Get the horizontal zoom factor of the image.
 * @param obj       pointer to an image object
 * @return          zoom factor (256: no zoom)
 */
// llgo:link (*LvObjT).LvImageGetScaleX C.lv_image_get_scale_x
func (recv_ *LvObjT) LvImageGetScaleX() c.Int32T {
	return 0
}

/**
 * Get the vertical zoom factor of the image.
 * @param obj       pointer to an image object
 * @return          zoom factor (256: no zoom)
 */
// llgo:link (*LvObjT).LvImageGetScaleY C.lv_image_get_scale_y
func (recv_ *LvObjT) LvImageGetScaleY() c.Int32T {
	return 0
}

/**
 * Get the width of an image before any transformations.
 * @param obj Pointer to an image object.
 * @return The width of the image.
 */
// llgo:link (*LvObjT).LvImageGetSrcWidth C.lv_image_get_src_width
func (recv_ *LvObjT) LvImageGetSrcWidth() c.Int32T {
	return 0
}

/**
 * Get the height of an image before any transformations.
 * @param obj Pointer to an image object.
 * @return The height of the image.
 */
// llgo:link (*LvObjT).LvImageGetSrcHeight C.lv_image_get_src_height
func (recv_ *LvObjT) LvImageGetSrcHeight() c.Int32T {
	return 0
}

/**
 * Get the transformed width of an image object.
 * @param obj Pointer to an image object.
 * @return The transformed width of the image.
 */
// llgo:link (*LvObjT).LvImageGetTransformedWidth C.lv_image_get_transformed_width
func (recv_ *LvObjT) LvImageGetTransformedWidth() c.Int32T {
	return 0
}

/**
 * Get the transformed height of an image object.
 * @param obj Pointer to an image object.
 * @return The transformed height of the image.
 */
// llgo:link (*LvObjT).LvImageGetTransformedHeight C.lv_image_get_transformed_height
func (recv_ *LvObjT) LvImageGetTransformedHeight() c.Int32T {
	return 0
}

/**
 * Get the current blend mode of the image
 * @param obj       pointer to an image object
 * @return          the current blend mode
 */
// llgo:link (*LvObjT).LvImageGetBlendMode C.lv_image_get_blend_mode
func (recv_ *LvObjT) LvImageGetBlendMode() LvBlendModeT {
	return 0
}

/**
 * Get whether the transformations (rotate, zoom) are anti-aliased or not
 * @param obj       pointer to an image object
 * @return          true: anti-aliased; false: not anti-aliased
 */
// llgo:link (*LvObjT).LvImageGetAntialias C.lv_image_get_antialias
func (recv_ *LvObjT) LvImageGetAntialias() bool {
	return false
}

/**
 * Get the size mode of the image
 * @param obj       pointer to an image object
 * @return          element of `lv_image_align_t`
 */
// llgo:link (*LvObjT).LvImageGetInnerAlign C.lv_image_get_inner_align
func (recv_ *LvObjT) LvImageGetInnerAlign() LvImageAlignT {
	return 0
}

/**
 * Get the bitmap mask source.
 * @param obj       pointer to an image object
 * @return          an lv_image_dsc_t bitmap mask source.
 */
// llgo:link (*LvObjT).LvImageGetBitmapMapSrc C.lv_image_get_bitmap_map_src
func (recv_ *LvObjT) LvImageGetBitmapMapSrc() *LvImageDscT {
	return nil
}

/**
 * Create a canvas object
 * @param parent     pointer to an object, it will be the parent of the new canvas
 * @return           pointer to the created canvas
 */
// llgo:link (*LvObjT).LvCanvasCreate C.lv_canvas_create
func (recv_ *LvObjT) LvCanvasCreate() *LvObjT {
	return nil
}

/**
 * Set a buffer for the canvas.
 *
 * Use lv_canvas_set_draw_buf() instead if you need to set a buffer with alignment requirement.
 *
 * @param obj    pointer to a canvas object
 * @param buf    buffer where content of canvas will be.
 *                 The required size is (lv_image_color_format_get_px_size(cf) * w) / 8 * h)
 *                 It can be allocated with `lv_malloc()` or
 *                 it can be statically allocated array (e.g. static lv_color_t buf[100*50]) or
 *                 it can be an address in RAM or external SRAM
 * @param w      width of canvas
 * @param h      height of canvas
 * @param cf     color format. `LV_COLOR_FORMAT...`
 */
// llgo:link (*LvObjT).LvCanvasSetBuffer C.lv_canvas_set_buffer
func (recv_ *LvObjT) LvCanvasSetBuffer(buf c.Pointer, w c.Int32T, h c.Int32T, cf LvColorFormatT) {
}

/**
 * Set a draw buffer for the canvas. A draw buffer either can be allocated by `lv_draw_buf_create()`
 * or defined statically by `LV_DRAW_BUF_DEFINE_STATIC`. When buffer start address and stride has alignment
 * requirement, it's recommended to use `lv_draw_buf_create`.
 * @param obj       pointer to a canvas object
 * @param draw_buf  pointer to a draw buffer
 */
// llgo:link (*LvObjT).LvCanvasSetDrawBuf C.lv_canvas_set_draw_buf
func (recv_ *LvObjT) LvCanvasSetDrawBuf(draw_buf *LvDrawBufT) {
}

/**
 * Set a pixel's color and opacity
 * @param obj   pointer to a canvas
 * @param x     X coordinate of the pixel
 * @param y     Y coordinate of the pixel
 * @param color the color
 * @param opa   the opacity
 * @note        The following color formats are supported
 *              LV_COLOR_FORMAT_I1/2/4/8, LV_COLOR_FORMAT_A8,
 *              LV_COLOR_FORMAT_RGB565, LV_COLOR_FORMAT_RGB888,
 *              LV_COLOR_FORMAT_XRGB8888, LV_COLOR_FORMAT_ARGB8888
 */
// llgo:link (*LvObjT).LvCanvasSetPx C.lv_canvas_set_px
func (recv_ *LvObjT) LvCanvasSetPx(x c.Int32T, y c.Int32T, color LvColorT, opa LvOpaT) {
}

/**
 * Set the palette color of a canvas for index format. Valid only for `LV_COLOR_FORMAT_I1/2/4/8`
 * @param obj       pointer to canvas object
 * @param index     the palette color to set:
 *                  - for `LV_COLOR_FORMAT_I1`: 0..1
 *                  - for `LV_COLOR_FORMAT_I2`: 0..3
 *                  - for `LV_COLOR_FORMAT_I4`: 0..15
 *                  - for `LV_COLOR_FORMAT_I8`: 0..255
 * @param color     the color to set
 */
// llgo:link (*LvObjT).LvCanvasSetPalette C.lv_canvas_set_palette
func (recv_ *LvObjT) LvCanvasSetPalette(index c.Uint8T, color LvColor32T) {
}

/*=====================
 * Getter functions
 *====================*/
// llgo:link (*LvObjT).LvCanvasGetDrawBuf C.lv_canvas_get_draw_buf
func (recv_ *LvObjT) LvCanvasGetDrawBuf() *LvDrawBufT {
	return nil
}

/**
 * Get a pixel's color and opacity
 * @param obj   pointer to a canvas
 * @param x     X coordinate of the pixel
 * @param y     Y coordinate of the pixel
 * @return      ARGB8888 color of the pixel
 */
// llgo:link (*LvObjT).LvCanvasGetPx C.lv_canvas_get_px
func (recv_ *LvObjT) LvCanvasGetPx(x c.Int32T, y c.Int32T) LvColor32T {
	return LvColor32T{}
}

/**
 * Get the image of the canvas as a pointer to an `lv_image_dsc_t` variable.
 * @param canvas    pointer to a canvas object
 * @return          pointer to the image descriptor.
 */
// llgo:link (*LvObjT).LvCanvasGetImage C.lv_canvas_get_image
func (recv_ *LvObjT) LvCanvasGetImage() *LvImageDscT {
	return nil
}

/**
 * Return the pointer for the buffer.
 * It's recommended to use this function instead of the buffer form the
 * return value of lv_canvas_get_image() as is can be aligned
 * @param canvas    pointer to a canvas object
 * @return          pointer to the buffer
 */
// llgo:link (*LvObjT).LvCanvasGetBuf C.lv_canvas_get_buf
func (recv_ *LvObjT) LvCanvasGetBuf() c.Pointer {
	return nil
}

/**
 * Copy a buffer to the canvas
 * @param obj           pointer to a canvas object
 * @param canvas_area   the area of the canvas to copy
 * @param dest_buf      pointer to a buffer to store the copied data
 * @param dest_area     the area of the destination buffer to copy to. If omitted NULL, copy to the whole `dest_buf`
 */
// llgo:link (*LvObjT).LvCanvasCopyBuf C.lv_canvas_copy_buf
func (recv_ *LvObjT) LvCanvasCopyBuf(canvas_area *LvAreaT, dest_buf *LvDrawBufT, dest_area *LvAreaT) {
}

/**
 * Fill the canvas with color
 * @param obj       pointer to a canvas
 * @param color     the background color
 * @param opa       the desired opacity
 */
// llgo:link (*LvObjT).LvCanvasFillBg C.lv_canvas_fill_bg
func (recv_ *LvObjT) LvCanvasFillBg(color LvColorT, opa LvOpaT) {
}

/**
 * Initialize a layer to use LVGL's generic draw functions (lv_draw_rect/label/...) on the canvas.
 * Needs to be usd in pair with `lv_canvas_finish_layer`.
 * @param canvas    pointer to a canvas
 * @param layer     pointer to a layer variable to initialize
 */
// llgo:link (*LvObjT).LvCanvasInitLayer C.lv_canvas_init_layer
func (recv_ *LvObjT) LvCanvasInitLayer(layer *LvLayerT) {
}

/**
 * Wait until all the drawings are finished on layer.
 * Needs to be usd in pair with `lv_canvas_init_layer`.
 * @param canvas    pointer to a canvas
 * @param layer     pointer to a layer to finalize
 */
// llgo:link (*LvObjT).LvCanvasFinishLayer C.lv_canvas_finish_layer
func (recv_ *LvObjT) LvCanvasFinishLayer(layer *LvLayerT) {
}

/**
 * Just a wrapper to `LV_CANVAS_BUF_SIZE` for bindings.
 */
//go:linkname LvCanvasBufSize C.lv_canvas_buf_size
func LvCanvasBufSize(w c.Int32T, h c.Int32T, bpp c.Uint8T, stride c.Uint8T) c.Uint32T

/**
 * Can be used by draw units to handle the decoding and
 * prepare everything for the actual image rendering
 * @param t             pointer to a draw task
 * @param draw_dsc      the draw descriptor of the image
 * @param coords        the absolute coordinates of the image
 * @param draw_core_cb  a callback to perform the actual rendering
 */
// llgo:link (*LvDrawTaskT).LvDrawImageNormalHelper C.lv_draw_image_normal_helper
func (recv_ *LvDrawTaskT) LvDrawImageNormalHelper(draw_dsc *LvDrawImageDscT, coords *LvAreaT, draw_core_cb LvDrawImageCoreCb) {
}

/**
 * Can be used by draw units for TILED images to handle the decoding and
 * prepare everything for the actual image rendering
 * @param t             pointer to a draw task
 * @param draw_dsc      the draw descriptor of the image
 * @param coords        the absolute coordinates of the image
 * @param draw_core_cb  a callback to perform the actual rendering
 */
// llgo:link (*LvDrawTaskT).LvDrawImageTiledHelper C.lv_draw_image_tiled_helper
func (recv_ *LvDrawTaskT) LvDrawImageTiledHelper(draw_dsc *LvDrawImageDscT, coords *LvAreaT, draw_core_cb LvDrawImageCoreCb) {
}

/**
 * Get the area of a rectangle if its rotated and scaled
 * @param res store the coordinates here
 * @param w width of the rectangle to transform
 * @param h height of the rectangle to transform
 * @param angle angle of rotation
 * @param scale_x zoom in x direction, (256 no zoom)
 * @param scale_y zoom in y direction, (256 no zoom)
 * @param pivot x,y pivot coordinates of rotation
 */
// llgo:link (*LvAreaT).LvImageBufGetTransformedArea C.lv_image_buf_get_transformed_area
func (recv_ *LvAreaT) LvImageBufGetTransformedArea(w c.Int32T, h c.Int32T, angle c.Int32T, scale_x c.Uint16T, scale_y c.Uint16T, pivot *LvPointT) {
}

/**
 * Initialize the image decoder module
 * @param image_cache_size    Image cache size in bytes. 0 to disable cache.
 * @param image_header_count  Number of header cache entries. 0 to disable header cache.
 */
//go:linkname LvImageDecoderInit C.lv_image_decoder_init
func LvImageDecoderInit(image_cache_size c.Uint32T, image_header_count c.Uint32T)

/**
 * Deinitialize the image decoder module
 */
//go:linkname LvImageDecoderDeinit C.lv_image_decoder_deinit
func LvImageDecoderDeinit()

// llgo:link (*LvDrawMaskRectDscT).LvDrawMaskRectDscInit C.lv_draw_mask_rect_dsc_init
func (recv_ *LvDrawMaskRectDscT) LvDrawMaskRectDscInit() {
}

/**
 * Try to get a rectangle mask draw descriptor from a draw task.
 * @param task      draw task
 * @return          the task's draw descriptor or NULL if the task is not of type LV_DRAW_TASK_TYPE_MASK_RECTANGLE
 */
// llgo:link (*LvDrawTaskT).LvDrawTaskGetMaskRectDsc C.lv_draw_task_get_mask_rect_dsc
func (recv_ *LvDrawTaskT) LvDrawTaskGetMaskRectDsc() *LvDrawMaskRectDscT {
	return nil
}

/**
 * Create a draw task to mask a rectangle from the buffer
 * @param layer     pointer to a layer
 * @param dsc       pointer to a draw descriptor
 */
// llgo:link (*LvLayerT).LvDrawMaskRect C.lv_draw_mask_rect
func (recv_ *LvLayerT) LvDrawMaskRect(dsc *LvDrawMaskRectDscT) {
}

/** This describes a glyph.*/

type LvFontFmtTxtGlyphDscT struct {
	BitmapIndex c.Uint32T
	AdvW        c.Uint32T
	BoxW        c.Uint8T
	BoxH        c.Uint8T
	OfsX        c.Int8T
	OfsY        c.Int8T
}
type LvFontFmtTxtCmapTypeT c.Int

const (
	LV_FONT_FMT_TXT_CMAP_FORMAT0_FULL LvFontFmtTxtCmapTypeT = 0
	LV_FONT_FMT_TXT_CMAP_SPARSE_FULL  LvFontFmtTxtCmapTypeT = 1
	LV_FONT_FMT_TXT_CMAP_FORMAT0_TINY LvFontFmtTxtCmapTypeT = 2
	LV_FONT_FMT_TXT_CMAP_SPARSE_TINY  LvFontFmtTxtCmapTypeT = 3
)

/**
 * Map codepoints to a `glyph_dsc`s
 * Several formats are supported to optimize memory usage
 * See https://github.com/lvgl/lv_font_conv/blob/master/doc/font_spec.md
 */

type LvFontFmtTxtCmapT struct {
	RangeStart     c.Uint32T
	RangeLength    c.Uint16T
	GlyphIdStart   c.Uint16T
	UnicodeList    *c.Uint16T
	GlyphIdOfsList c.Pointer
	ListLength     c.Uint16T
	Type           LvFontFmtTxtCmapTypeT
}

/** A simple mapping of kern values from pairs*/

type LvFontFmtTxtKernPairT struct {
	GlyphIds     c.Pointer
	Values       *c.Int8T
	PairCnt      c.Uint32T
	GlyphIdsSize c.Uint32T
}

/** More complex but more optimal class based kern value storage*/

type LvFontFmtTxtKernClassesT struct {
	ClassPairValues   *c.Int8T
	LeftClassMapping  *c.Uint8T
	RightClassMapping *c.Uint8T
	LeftClassCnt      c.Uint8T
	RightClassCnt     c.Uint8T
}
type LvFontFmtTxtBitmapFormatT c.Int

const (
	LV_FONT_FMT_TXT_PLAIN                   LvFontFmtTxtBitmapFormatT = 0
	LV_FONT_FMT_TXT_COMPRESSED              LvFontFmtTxtBitmapFormatT = 1
	LV_FONT_FMT_TXT_COMPRESSED_NO_PREFILTER LvFontFmtTxtBitmapFormatT = 2
)

/** Describe store for additional data for fonts */

type LvFontFmtTxtDscT struct {
	GlyphBitmap  *c.Uint8T
	GlyphDsc     *LvFontFmtTxtGlyphDscT
	Cmaps        *LvFontFmtTxtCmapT
	KernDsc      c.Pointer
	KernScale    c.Uint16T
	CmapNum      c.Uint16T
	Bpp          c.Uint16T
	KernClasses  c.Uint16T
	BitmapFormat c.Uint16T
	Stride       c.Uint8T
}

type LvBuiltinFontSrcT struct {
	FontP *LvFontT
	Size  c.Uint32T
}

/**
 * Used as `get_glyph_bitmap` callback in lvgl's native font format if the font is uncompressed.
 * @param g_dsc         the glyph descriptor including which font to use, which supply the glyph_index and format.
 * @param draw_buf      a draw buffer that can be used to store the bitmap of the glyph, it's OK not to use it.
 * @return pointer to an A8 bitmap (not necessarily bitmap_out) or NULL if `unicode_letter` not found
 */
// llgo:link (*LvFontGlyphDscT).LvFontGetBitmapFmtTxt C.lv_font_get_bitmap_fmt_txt
func (recv_ *LvFontGlyphDscT) LvFontGetBitmapFmtTxt(draw_buf *LvDrawBufT) c.Pointer {
	return nil
}

/**
 * Used as `get_glyph_dsc` callback in lvgl's native font format if the font is uncompressed.
 * @param font pointer to font
 * @param dsc_out store the result descriptor here
 * @param unicode_letter a UNICODE letter code
 * @param unicode_letter_next the unicode letter succeeding the letter under test
 * @return true: descriptor is successfully loaded into `dsc_out`.
 *         false: the letter was not found, no data is loaded to `dsc_out`
 */
// llgo:link (*LvFontT).LvFontGetGlyphDscFmtTxt C.lv_font_get_glyph_dsc_fmt_txt
func (recv_ *LvFontT) LvFontGetGlyphDscFmtTxt(dsc_out *LvFontGlyphDscT, unicode_letter c.Uint32T, unicode_letter_next c.Uint32T) bool {
	return false
}

// llgo:type C
type LvThemeApplyCbT func(*LvThemeT, *LvObjT)

/**
 * Get the theme assigned to the display of the object
 * @param obj       pointer to a theme object
 * @return          the theme of the object's display (can be NULL)
 */
// llgo:link (*LvObjT).LvThemeGetFromObj C.lv_theme_get_from_obj
func (recv_ *LvObjT) LvThemeGetFromObj() *LvThemeT {
	return nil
}

/**
 * Apply the active theme on an object
 * @param obj pointer to an object
 */
// llgo:link (*LvObjT).LvThemeApply C.lv_theme_apply
func (recv_ *LvObjT) LvThemeApply() {
}

/**
 * Set a base theme for a theme.
 * The styles from the base them will be added before the styles of the current theme.
 * Arbitrary long chain of themes can be created by setting base themes.
 * @param new_theme pointer to theme which base should be set
 * @param parent pointer to the base theme
 */
// llgo:link (*LvThemeT).LvThemeSetParent C.lv_theme_set_parent
func (recv_ *LvThemeT) LvThemeSetParent(parent *LvThemeT) {
}

/**
 * Set an apply callback for a theme.
 * The apply callback is used to add styles to different objects
 * @param theme pointer to theme which callback should be set
 * @param apply_cb pointer to the callback
 */
// llgo:link (*LvThemeT).LvThemeSetApplyCb C.lv_theme_set_apply_cb
func (recv_ *LvThemeT) LvThemeSetApplyCb(apply_cb LvThemeApplyCbT) {
}

/**
 * Get the small font of the theme
 * @param obj pointer to an object
 * @return pointer to the font
 */
// llgo:link (*LvObjT).LvThemeGetFontSmall C.lv_theme_get_font_small
func (recv_ *LvObjT) LvThemeGetFontSmall() *LvFontT {
	return nil
}

/**
 * Get the normal font of the theme
 * @param obj pointer to an object
 * @return pointer to the font
 */
// llgo:link (*LvObjT).LvThemeGetFontNormal C.lv_theme_get_font_normal
func (recv_ *LvObjT) LvThemeGetFontNormal() *LvFontT {
	return nil
}

/**
 * Get the subtitle font of the theme
 * @param obj pointer to an object
 * @return pointer to the font
 */
// llgo:link (*LvObjT).LvThemeGetFontLarge C.lv_theme_get_font_large
func (recv_ *LvObjT) LvThemeGetFontLarge() *LvFontT {
	return nil
}

/**
 * Get the primary color of the theme
 * @param obj pointer to an object
 * @return the color
 */
// llgo:link (*LvObjT).LvThemeGetColorPrimary C.lv_theme_get_color_primary
func (recv_ *LvObjT) LvThemeGetColorPrimary() LvColorT {
	return LvColorT{}
}

/**
 * Get the secondary color of the theme
 * @param obj pointer to an object
 * @return the color
 */
// llgo:link (*LvObjT).LvThemeGetColorSecondary C.lv_theme_get_color_secondary
func (recv_ *LvObjT) LvThemeGetColorSecondary() LvColorT {
	return LvColorT{}
}

/**
 * Initialize the theme
 * @param disp pointer to display
 * @param color_primary the primary color of the theme
 * @param color_secondary the secondary color for the theme
 * @param dark
 * @param font pointer to a font to use.
 * @return a pointer to reference this theme later
 */
// llgo:link (*LvDisplayT).LvThemeDefaultInit C.lv_theme_default_init
func (recv_ *LvDisplayT) LvThemeDefaultInit(color_primary LvColorT, color_secondary LvColorT, dark bool, font *LvFontT) *LvThemeT {
	return nil
}

/**
 * Check if default theme is initialized
 * @return true if default theme is initialized, false otherwise
 */
//go:linkname LvThemeDefaultIsInited C.lv_theme_default_is_inited
func LvThemeDefaultIsInited() bool

/**
 * Get default theme
 * @return a pointer to default theme, or NULL if this is not initialized
 */
//go:linkname LvThemeDefaultGet C.lv_theme_default_get
func LvThemeDefaultGet() *LvThemeT

/**
 * Deinitialize the default theme
 */
//go:linkname LvThemeDefaultDeinit C.lv_theme_default_deinit
func LvThemeDefaultDeinit()

/**
 * Initialize the theme
 * @param disp pointer to display
 * @param dark_bg
 * @param font pointer to a font to use.
 * @return a pointer to reference this theme later
 */
// llgo:link (*LvDisplayT).LvThemeMonoInit C.lv_theme_mono_init
func (recv_ *LvDisplayT) LvThemeMonoInit(dark_bg bool, font *LvFontT) *LvThemeT {
	return nil
}

/**
* Check if the theme is initialized
* @return true if default theme is initialized, false otherwise
 */
//go:linkname LvThemeMonoIsInited C.lv_theme_mono_is_inited
func LvThemeMonoIsInited() bool

/**
 * Get mono theme
 * @return a pointer to mono theme, or NULL if this is not initialized
 */
//go:linkname LvThemeMonoGet C.lv_theme_mono_get
func LvThemeMonoGet() *LvThemeT

/**
 * Deinitialize the mono theme
 */
//go:linkname LvThemeMonoDeinit C.lv_theme_mono_deinit
func LvThemeMonoDeinit()

/**
 * Initialize the theme
 * @param disp pointer to display
 * @return a pointer to reference this theme later
 */
// llgo:link (*LvDisplayT).LvThemeSimpleInit C.lv_theme_simple_init
func (recv_ *LvDisplayT) LvThemeSimpleInit() *LvThemeT {
	return nil
}

/**
* Check if the theme is initialized
* @return true if default theme is initialized, false otherwise
 */
//go:linkname LvThemeSimpleIsInited C.lv_theme_simple_is_inited
func LvThemeSimpleIsInited() bool

/**
 * Get simple theme
 * @return a pointer to simple theme, or NULL if this is not initialized
 */
//go:linkname LvThemeSimpleGet C.lv_theme_simple_get
func LvThemeSimpleGet() *LvThemeT

/**
 * Deinitialize the simple theme
 */
//go:linkname LvThemeSimpleDeinit C.lv_theme_simple_deinit
func LvThemeSimpleDeinit()

/**
 * Redraw the invalidated areas now.
 * Normally the redrawing is periodically executed in `lv_timer_handler` but a long blocking process
 * can prevent the call of `lv_timer_handler`. In this case if the GUI is updated in the process
 * (e.g. progress bar) this function can be called when the screen should be updated.
 * @param disp pointer to display to refresh. NULL to refresh all displays.
 */
// llgo:link (*LvDisplayT).LvRefrNow C.lv_refr_now
func (recv_ *LvDisplayT) LvRefrNow() {
}

/**
 * Redrawn on object and all its children using the passed draw context
 * @param layer pointer to a layer where to draw.
 * @param obj   the start object from the redraw should start
 */
// llgo:link (*LvLayerT).LvObjRedraw C.lv_obj_redraw
func (recv_ *LvLayerT) LvObjRedraw(obj *LvObjT) {
}

/**
 * Called periodically to handle the refreshing
 * @param timer pointer to the timer itself, or `NULL`
 */
// llgo:link (*LvTimerT).LvDisplayRefrTimer C.lv_display_refr_timer
func (recv_ *LvTimerT) LvDisplayRefrTimer() {
}

/**
 * Initialize the screen refresh subsystem
 */
//go:linkname LvRefrInit C.lv_refr_init
func LvRefrInit()

/**
 * Deinitialize the screen refresh subsystem
 */
//go:linkname LvRefrDeinit C.lv_refr_deinit
func LvRefrDeinit()

/**
 * Invalidate an area on display to redraw it
 * @param area_p pointer to area which should be invalidated (NULL: delete the invalidated areas)
 * @param disp pointer to display where the area should be invalidated (NULL can be used if there is
 * only one display)
 */
// llgo:link (*LvDisplayT).LvInvArea C.lv_inv_area
func (recv_ *LvDisplayT) LvInvArea(area_p *LvAreaT) {
}

/**
 * Get the display which is being refreshed
 * @return the display being refreshed
 */
//go:linkname LvRefrGetDispRefreshing C.lv_refr_get_disp_refreshing
func LvRefrGetDispRefreshing() *LvDisplayT

/**
 * Set the display which is being refreshed
 * @param disp the display being refreshed
 */
// llgo:link (*LvDisplayT).LvRefrSetDispRefreshing C.lv_refr_set_disp_refreshing
func (recv_ *LvDisplayT) LvRefrSetDispRefreshing() {
}

/**
 * Search the most top object which fully covers an area
 * @param area_p pointer to an area
 * @param obj the first object to start the searching (typically a screen)
 * @return
 */
// llgo:link (*LvAreaT).LvRefrGetTopObj C.lv_refr_get_top_obj
func (recv_ *LvAreaT) LvRefrGetTopObj(obj *LvObjT) *LvObjT {
	return nil
}

/**
 * Initialize the object related style manager module.
 * Called by LVGL in `lv_init()`
 */
//go:linkname LvObjStyleInit C.lv_obj_style_init
func LvObjStyleInit()

/**
 * Deinitialize the object related style manager module.
 * Called by LVGL in `lv_deinit()`
 */
//go:linkname LvObjStyleDeinit C.lv_obj_style_deinit
func LvObjStyleDeinit()

/**
 * Used internally to create a style transition
 * @param obj
 * @param part
 * @param prev_state
 * @param new_state
 * @param tr
 */
// llgo:link (*LvObjT).LvObjStyleCreateTransition C.lv_obj_style_create_transition
func (recv_ *LvObjT) LvObjStyleCreateTransition(part LvPartT, prev_state LvStateT, new_state LvStateT, tr *LvObjStyleTransitionDscT) {
}

/**
 * Used internally to compare the appearance of an object in 2 states
 * @param obj
 * @param state1
 * @param state2
 * @return
 */
// llgo:link (*LvObjT).LvObjStyleStateCompare C.lv_obj_style_state_compare
func (recv_ *LvObjT) LvObjStyleStateCompare(state1 LvStateT, state2 LvStateT) LvStyleStateCmpT {
	return 0
}

/**
 * Update the layer type of a widget bayed on its current styles.
 * The result will be stored in `obj->spec_attr->layer_type`
 * @param obj       the object whose layer should be updated
 */
// llgo:link (*LvObjT).LvObjUpdateLayerType C.lv_obj_update_layer_type
func (recv_ *LvObjT) LvObjUpdateLayerType() {
}

/**
 * Low level function to scroll by given x and y coordinates.
 * `LV_EVENT_SCROLL` is sent.
 * @param obj       pointer to an object to scroll
 * @param x         pixels to scroll horizontally
 * @param y         pixels to scroll vertically
 * @return          `LV_RESULT_INVALID`: to object was deleted in `LV_EVENT_SCROLL`;
 *                  `LV_RESULT_OK`: if the object is still valid
 */
// llgo:link (*LvObjT).LvObjScrollByRaw C.lv_obj_scroll_by_raw
func (recv_ *LvObjT) LvObjScrollByRaw(x c.Int32T, y c.Int32T) LvResultT {
	return 0
}

/**
 * Get the extended draw area of an object.
 * @param obj       pointer to an object
 * @return          the size extended draw area around the real coordinates
 */
// llgo:link (*LvObjT).LvObjGetExtDrawSize C.lv_obj_get_ext_draw_size
func (recv_ *LvObjT) LvObjGetExtDrawSize() c.Int32T {
	return 0
}

// llgo:link (*LvObjT).LvObjGetLayerType C.lv_obj_get_layer_type
func (recv_ *LvObjT) LvObjGetLayerType() LvLayerTypeT {
	return 0
}

/**********************
 * GLOBAL PROTOTYPES
 **********************/
// llgo:link (*LvObjT).LvObjDestruct C.lv_obj_destruct
func (recv_ *LvObjT) LvObjDestruct() {
}

/**
 * Init the group module
 * @remarks Internal function, do not call directly.
 */
//go:linkname LvGroupInit C.lv_group_init
func LvGroupInit()

/**
 * Deinit the group module
 * @remarks Internal function, do not call directly.
 */
//go:linkname LvGroupDeinit C.lv_group_deinit
func LvGroupDeinit()

/**
 * Set the position of an area (width and height will be kept)
 * @param area_p pointer to an area
 * @param x the new x coordinate of the area
 * @param y the new y coordinate of the area
 */
// llgo:link (*LvAreaT).LvAreaSetPos C.lv_area_set_pos
func (recv_ *LvAreaT) LvAreaSetPos(x c.Int32T, y c.Int32T) {
}

/**
 * Get the common parts of two areas
 * @param res_p pointer to an area, the result will be stored her
 * @param a1_p pointer to the first area
 * @param a2_p pointer to the second area
 * @return false: the two area has NO common parts, res_p is invalid
 */
// llgo:link (*LvAreaT).LvAreaIntersect C.lv_area_intersect
func (recv_ *LvAreaT) LvAreaIntersect(a1_p *LvAreaT, a2_p *LvAreaT) bool {
	return false
}

/**
 * Get resulting sub areas after removing the common parts of two areas from the first area
 * @param res_p pointer to an array of areas with a count of 4, the resulting areas will be stored here
 * @param a1_p pointer to the first area
 * @param a2_p pointer to the second area
 * @return number of results (max 4) or -1 if no intersect
 */
//go:linkname LvAreaDiff C.lv_area_diff
func LvAreaDiff(res_p *LvAreaT, a1_p *LvAreaT, a2_p *LvAreaT) c.Int8T

/**
 * Join two areas into a third which involves the other two
 * @param a_res_p pointer to an area, the result will be stored here
 * @param a1_p pointer to the first area
 * @param a2_p pointer to the second area
 */
// llgo:link (*LvAreaT).LvAreaJoin C.lv_area_join
func (recv_ *LvAreaT) LvAreaJoin(a1_p *LvAreaT, a2_p *LvAreaT) {
}

/**
 * Check if a point is on an area
 * @param a_p pointer to an area
 * @param p_p pointer to a point
 * @param radius radius of area (e.g. for rounded rectangle)
 * @return false:the point is out of the area
 */
// llgo:link (*LvAreaT).LvAreaIsPointOn C.lv_area_is_point_on
func (recv_ *LvAreaT) LvAreaIsPointOn(p_p *LvPointT, radius c.Int32T) bool {
	return false
}

/**
 * Check if two area has common parts
 * @param a1_p pointer to an area.
 * @param a2_p pointer to another area
 * @return false: a1_p and a2_p has no common parts
 */
// llgo:link (*LvAreaT).LvAreaIsOn C.lv_area_is_on
func (recv_ *LvAreaT) LvAreaIsOn(a2_p *LvAreaT) bool {
	return false
}

/**
 * Check if an area is fully on another
 * @param ain_p pointer to an area which could be in 'aholder_p'
 * @param aholder_p pointer to an area which could involve 'ain_p'
 * @param radius radius of `aholder_p` (e.g. for rounded rectangle)
 * @return true: `ain_p` is fully inside `aholder_p`
 */
// llgo:link (*LvAreaT).LvAreaIsIn C.lv_area_is_in
func (recv_ *LvAreaT) LvAreaIsIn(aholder_p *LvAreaT, radius c.Int32T) bool {
	return false
}

/**
 * Check if an area is fully out of another
 * @param aout_p pointer to an area which could be in 'aholder_p'
 * @param aholder_p pointer to an area which could involve 'ain_p'
 * @param radius radius of `aholder_p` (e.g. for rounded rectangle)
 * @return true: `aout_p` is fully outside `aholder_p`
 */
// llgo:link (*LvAreaT).LvAreaIsOut C.lv_area_is_out
func (recv_ *LvAreaT) LvAreaIsOut(aholder_p *LvAreaT, radius c.Int32T) bool {
	return false
}

/**
 * Check if 2 area is the same
 * @param a pointer to an area
 * @param b pointer to another area
 */
// llgo:link (*LvAreaT).LvAreaIsEqual C.lv_area_is_equal
func (recv_ *LvAreaT) LvAreaIsEqual(b *LvAreaT) bool {
	return false
}

/**
 * Initialize the File system interface
 */
//go:linkname LvFsInit C.lv_fs_init
func LvFsInit()

/**
 * Deinitialize the File system interface
 */
//go:linkname LvFsDeinit C.lv_fs_deinit
func LvFsDeinit()

/**********************
 * GLOBAL PROTOTYPES
 **********************/
// llgo:link (*LvEventT).LvEventPush C.lv_event_push
func (recv_ *LvEventT) LvEventPush() {
}

// llgo:link (*LvEventT).LvEventPop C.lv_event_pop
func (recv_ *LvEventT) LvEventPop() {
}

/**
 * Nested events can be called and one of them might belong to an object that is being deleted.
 * Mark this object's `event_temp_data` deleted to know that its `lv_obj_send_event` should return `LV_RESULT_INVALID`
 * @param target     pointer to an event target which was deleted
 */
//go:linkname LvEventMarkDeleted C.lv_event_mark_deleted
func LvEventMarkDeleted(target c.Pointer)

type LvRbColorT c.Int

const (
	LV_RB_COLOR_RED   LvRbColorT = 0
	LV_RB_COLOR_BLACK LvRbColorT = 1
)

type LvRbCompareResT c.Int8T

// llgo:type C
type LvRbCompareT func(c.Pointer, c.Pointer) LvRbCompareResT

/**********************
 * GLOBAL PROTOTYPES
 **********************/
// llgo:link (*LvRbT).LvRbInit C.lv_rb_init
func (recv_ *LvRbT) LvRbInit(compare LvRbCompareT, node_size c.SizeT) bool {
	return false
}

// llgo:link (*LvRbT).LvRbInsert C.lv_rb_insert
func (recv_ *LvRbT) LvRbInsert(key c.Pointer) *LvRbNodeT {
	return nil
}

// llgo:link (*LvRbT).LvRbFind C.lv_rb_find
func (recv_ *LvRbT) LvRbFind(key c.Pointer) *LvRbNodeT {
	return nil
}

// llgo:link (*LvRbT).LvRbRemoveNode C.lv_rb_remove_node
func (recv_ *LvRbT) LvRbRemoveNode(node *LvRbNodeT) c.Pointer {
	return nil
}

// llgo:link (*LvRbT).LvRbRemove C.lv_rb_remove
func (recv_ *LvRbT) LvRbRemove(key c.Pointer) c.Pointer {
	return nil
}

// llgo:link (*LvRbT).LvRbDropNode C.lv_rb_drop_node
func (recv_ *LvRbT) LvRbDropNode(node *LvRbNodeT) bool {
	return false
}

// llgo:link (*LvRbT).LvRbDrop C.lv_rb_drop
func (recv_ *LvRbT) LvRbDrop(key c.Pointer) bool {
	return false
}

// llgo:link (*LvRbT).LvRbMinimum C.lv_rb_minimum
func (recv_ *LvRbT) LvRbMinimum() *LvRbNodeT {
	return nil
}

// llgo:link (*LvRbT).LvRbMaximum C.lv_rb_maximum
func (recv_ *LvRbT) LvRbMaximum() *LvRbNodeT {
	return nil
}

// llgo:link (*LvRbNodeT).LvRbMinimumFrom C.lv_rb_minimum_from
func (recv_ *LvRbNodeT) LvRbMinimumFrom() *LvRbNodeT {
	return nil
}

// llgo:link (*LvRbNodeT).LvRbMaximumFrom C.lv_rb_maximum_from
func (recv_ *LvRbNodeT) LvRbMaximumFrom() *LvRbNodeT {
	return nil
}

// llgo:link (*LvRbT).LvRbDestroy C.lv_rb_destroy
func (recv_ *LvRbT) LvRbDestroy() {
}

/**
 * Create an empty message box
 * @param parent        the parent or NULL to create a modal msgbox
 * @return              the created message box
 */
// llgo:link (*LvObjT).LvMsgboxCreate C.lv_msgbox_create
func (recv_ *LvObjT) LvMsgboxCreate() *LvObjT {
	return nil
}

/**
 * Add title to the message box. It also creates a header for the title.
 * @param obj           pointer to a message box
 * @param title         the text of the tile
 * @return              the created title label
 */
// llgo:link (*LvObjT).LvMsgboxAddTitle C.lv_msgbox_add_title
func (recv_ *LvObjT) LvMsgboxAddTitle(title *c.Char) *LvObjT {
	return nil
}

/**
 * Add a button to the header of to the message box. It also creates a header.
 * @param obj           pointer to a message box
 * @param icon          the icon of the button
 * @return              the created button
 */
// llgo:link (*LvObjT).LvMsgboxAddHeaderButton C.lv_msgbox_add_header_button
func (recv_ *LvObjT) LvMsgboxAddHeaderButton(icon c.Pointer) *LvObjT {
	return nil
}

/**
 * Add a text to the content area of message box. Multiple texts will be created below each other.
 * @param obj           pointer to a message box
 * @param text          text to add
 * @return              the created button
 */
// llgo:link (*LvObjT).LvMsgboxAddText C.lv_msgbox_add_text
func (recv_ *LvObjT) LvMsgboxAddText(text *c.Char) *LvObjT {
	return nil
}

/**
 * Add a button to the footer of to the message box. It also creates a footer.
 * @param obj           pointer to a message box
 * @param text          the text of the button
 * @return              the created button
 */
// llgo:link (*LvObjT).LvMsgboxAddFooterButton C.lv_msgbox_add_footer_button
func (recv_ *LvObjT) LvMsgboxAddFooterButton(text *c.Char) *LvObjT {
	return nil
}

/**
 * Add a close button to the message box. It also creates a header.
 * @param obj           pointer to a message box
 * @return              the created close button
 */
// llgo:link (*LvObjT).LvMsgboxAddCloseButton C.lv_msgbox_add_close_button
func (recv_ *LvObjT) LvMsgboxAddCloseButton() *LvObjT {
	return nil
}

/**
 * Get the header widget
 * @param obj           pointer to a message box
 * @return              the header, or NULL if not exists
 */
// llgo:link (*LvObjT).LvMsgboxGetHeader C.lv_msgbox_get_header
func (recv_ *LvObjT) LvMsgboxGetHeader() *LvObjT {
	return nil
}

/**
 * Get the footer widget
 * @param obj           pointer to a message box
 * @return              the footer, or NULL if not exists
 */
// llgo:link (*LvObjT).LvMsgboxGetFooter C.lv_msgbox_get_footer
func (recv_ *LvObjT) LvMsgboxGetFooter() *LvObjT {
	return nil
}

/**
 * Get the content widget
 * @param obj           pointer to a message box
 * @return              the content
 */
// llgo:link (*LvObjT).LvMsgboxGetContent C.lv_msgbox_get_content
func (recv_ *LvObjT) LvMsgboxGetContent() *LvObjT {
	return nil
}

/**
 * Get the title label
 * @param obj           pointer to a message box
 * @return              the title, or NULL if it does not exist
 */
// llgo:link (*LvObjT).LvMsgboxGetTitle C.lv_msgbox_get_title
func (recv_ *LvObjT) LvMsgboxGetTitle() *LvObjT {
	return nil
}

/**
 * Close a message box
 * @param mbox           pointer to a message box
 */
// llgo:link (*LvObjT).LvMsgboxClose C.lv_msgbox_close
func (recv_ *LvObjT) LvMsgboxClose() {
}

/**
 * Close a message box in the next call of the message box
 * @param mbox           pointer to a message box
 */
// llgo:link (*LvObjT).LvMsgboxCloseAsync C.lv_msgbox_close_async
func (recv_ *LvObjT) LvMsgboxCloseAsync() {
}

type LvButtonmatrixCtrlT c.Int

const (
	LV_BUTTONMATRIX_CTRL_NONE       LvButtonmatrixCtrlT = 0
	LV_BUTTONMATRIX_CTRL_WIDTH_1    LvButtonmatrixCtrlT = 1
	LV_BUTTONMATRIX_CTRL_WIDTH_2    LvButtonmatrixCtrlT = 2
	LV_BUTTONMATRIX_CTRL_WIDTH_3    LvButtonmatrixCtrlT = 3
	LV_BUTTONMATRIX_CTRL_WIDTH_4    LvButtonmatrixCtrlT = 4
	LV_BUTTONMATRIX_CTRL_WIDTH_5    LvButtonmatrixCtrlT = 5
	LV_BUTTONMATRIX_CTRL_WIDTH_6    LvButtonmatrixCtrlT = 6
	LV_BUTTONMATRIX_CTRL_WIDTH_7    LvButtonmatrixCtrlT = 7
	LV_BUTTONMATRIX_CTRL_WIDTH_8    LvButtonmatrixCtrlT = 8
	LV_BUTTONMATRIX_CTRL_WIDTH_9    LvButtonmatrixCtrlT = 9
	LV_BUTTONMATRIX_CTRL_WIDTH_10   LvButtonmatrixCtrlT = 10
	LV_BUTTONMATRIX_CTRL_WIDTH_11   LvButtonmatrixCtrlT = 11
	LV_BUTTONMATRIX_CTRL_WIDTH_12   LvButtonmatrixCtrlT = 12
	LV_BUTTONMATRIX_CTRL_WIDTH_13   LvButtonmatrixCtrlT = 13
	LV_BUTTONMATRIX_CTRL_WIDTH_14   LvButtonmatrixCtrlT = 14
	LV_BUTTONMATRIX_CTRL_WIDTH_15   LvButtonmatrixCtrlT = 15
	LV_BUTTONMATRIX_CTRL_HIDDEN     LvButtonmatrixCtrlT = 16
	LV_BUTTONMATRIX_CTRL_NO_REPEAT  LvButtonmatrixCtrlT = 32
	LV_BUTTONMATRIX_CTRL_DISABLED   LvButtonmatrixCtrlT = 64
	LV_BUTTONMATRIX_CTRL_CHECKABLE  LvButtonmatrixCtrlT = 128
	LV_BUTTONMATRIX_CTRL_CHECKED    LvButtonmatrixCtrlT = 256
	LV_BUTTONMATRIX_CTRL_CLICK_TRIG LvButtonmatrixCtrlT = 512
	LV_BUTTONMATRIX_CTRL_POPOVER    LvButtonmatrixCtrlT = 1024
	LV_BUTTONMATRIX_CTRL_RECOLOR    LvButtonmatrixCtrlT = 2048
	LV_BUTTONMATRIX_CTRL_RESERVED_1 LvButtonmatrixCtrlT = 4096
	LV_BUTTONMATRIX_CTRL_RESERVED_2 LvButtonmatrixCtrlT = 8192
	LV_BUTTONMATRIX_CTRL_CUSTOM_1   LvButtonmatrixCtrlT = 16384
	LV_BUTTONMATRIX_CTRL_CUSTOM_2   LvButtonmatrixCtrlT = 32768
)

// llgo:type C
type LvButtonmatrixButtonDrawCbT func(*LvObjT, c.Uint32T, *LvAreaT, *LvAreaT) bool

/**
 * Create a button matrix object
 * @param parent    pointer to an object, it will be the parent of the new button matrix
 * @return          pointer to the created button matrix
 */
// llgo:link (*LvObjT).LvButtonmatrixCreate C.lv_buttonmatrix_create
func (recv_ *LvObjT) LvButtonmatrixCreate() *LvObjT {
	return nil
}

/**
 * Set a new map. Buttons will be created/deleted according to the map. The
 * button matrix keeps a reference to the map and so the string array must not
 * be deallocated during the life of the matrix.
 * @param obj       pointer to a button matrix object
 * @param map       pointer a string array. The last string has to be: "". Use "\n" to make a line break.
 */
// llgo:link (*LvObjT).LvButtonmatrixSetMap C.lv_buttonmatrix_set_map
func (recv_ *LvObjT) LvButtonmatrixSetMap(map_ **c.Char) {
}

/**
 * Set the button control map (hidden, disabled etc.) for a button matrix.
 * The control map array will be copied and so may be deallocated after this
 * function returns.
 * @param obj       pointer to a button matrix object
 * @param ctrl_map  pointer to an array of `lv_button_ctrl_t` control bytes. The
 *                  length of the array and position of the elements must match
 *                  the number and order of the individual buttons (i.e. excludes
 *                  newline entries).
 *                  An element of the map should look like e.g.:
 *                 `ctrl_map[0] = width | LV_BUTTONMATRIX_CTRL_NO_REPEAT |  LV_BUTTONMATRIX_CTRL_TGL_ENABLE`
 */
// llgo:link (*LvObjT).LvButtonmatrixSetCtrlMap C.lv_buttonmatrix_set_ctrl_map
func (recv_ *LvObjT) LvButtonmatrixSetCtrlMap(ctrl_map *LvButtonmatrixCtrlT) {
}

/**
 * Set the selected buttons
 * @param obj        pointer to button matrix object
 * @param btn_id     0 based index of the button to modify. (Not counting new lines)
 */
// llgo:link (*LvObjT).LvButtonmatrixSetSelectedButton C.lv_buttonmatrix_set_selected_button
func (recv_ *LvObjT) LvButtonmatrixSetSelectedButton(btn_id c.Uint32T) {
}

/**
 * Set the attributes of a button of the button matrix
 * @param obj       pointer to button matrix object
 * @param btn_id    0 based index of the button to modify. (Not counting new lines)
 * @param ctrl      OR-ed attributes. E.g. `LV_BUTTONMATRIX_CTRL_NO_REPEAT | LV_BUTTONMATRIX_CTRL_CHECKABLE`
 */
// llgo:link (*LvObjT).LvButtonmatrixSetButtonCtrl C.lv_buttonmatrix_set_button_ctrl
func (recv_ *LvObjT) LvButtonmatrixSetButtonCtrl(btn_id c.Uint32T, ctrl LvButtonmatrixCtrlT) {
}

/**
 * Clear the attributes of a button of the button matrix
 * @param obj       pointer to button matrix object
 * @param btn_id    0 based index of the button to modify. (Not counting new lines)
 * @param ctrl      OR-ed attributes. E.g. `LV_BUTTONMATRIX_CTRL_NO_REPEAT | LV_BUTTONMATRIX_CTRL_CHECKABLE`
 */
// llgo:link (*LvObjT).LvButtonmatrixClearButtonCtrl C.lv_buttonmatrix_clear_button_ctrl
func (recv_ *LvObjT) LvButtonmatrixClearButtonCtrl(btn_id c.Uint32T, ctrl LvButtonmatrixCtrlT) {
}

/**
 * Set attributes of all buttons of a button matrix
 * @param obj       pointer to a button matrix object
 * @param ctrl      attribute(s) to set from `lv_buttonmatrix_ctrl_t`. Values can be ORed.
 */
// llgo:link (*LvObjT).LvButtonmatrixSetButtonCtrlAll C.lv_buttonmatrix_set_button_ctrl_all
func (recv_ *LvObjT) LvButtonmatrixSetButtonCtrlAll(ctrl LvButtonmatrixCtrlT) {
}

/**
 * Clear the attributes of all buttons of a button matrix
 * @param obj       pointer to a button matrix object
 * @param ctrl      attribute(s) to set from `lv_buttonmatrix_ctrl_t`. Values can be ORed.
 */
// llgo:link (*LvObjT).LvButtonmatrixClearButtonCtrlAll C.lv_buttonmatrix_clear_button_ctrl_all
func (recv_ *LvObjT) LvButtonmatrixClearButtonCtrlAll(ctrl LvButtonmatrixCtrlT) {
}

/**
 * Set a single button's relative width.
 * This method will cause the matrix be regenerated and is a relatively
 * expensive operation. It is recommended that initial width be specified using
 * `lv_buttonmatrix_set_ctrl_map` and this method only be used for dynamic changes.
 * @param obj       pointer to button matrix object
 * @param btn_id    0 based index of the button to modify.
 * @param width     relative width compared to the buttons in the same row. [1..15]
 */
// llgo:link (*LvObjT).LvButtonmatrixSetButtonWidth C.lv_buttonmatrix_set_button_width
func (recv_ *LvObjT) LvButtonmatrixSetButtonWidth(btn_id c.Uint32T, width c.Uint32T) {
}

/**
 * Make the button matrix like a selector widget (only one button may be checked at a time).
 * `LV_BUTTONMATRIX_CTRL_CHECKABLE` must be enabled on the buttons to be selected using
 * `lv_buttonmatrix_set_ctrl()` or `lv_buttonmatrix_set_button_ctrl_all()`.
 * @param obj       pointer to a button matrix object
 * @param en        whether "one check" mode is enabled
 */
// llgo:link (*LvObjT).LvButtonmatrixSetOneChecked C.lv_buttonmatrix_set_one_checked
func (recv_ *LvObjT) LvButtonmatrixSetOneChecked(en bool) {
}

/**
 * Get the current map of a button matrix
 * @param obj       pointer to a button matrix object
 * @return          the current map
 */
// llgo:link (*LvObjT).LvButtonmatrixGetMap C.lv_buttonmatrix_get_map
func (recv_ *LvObjT) LvButtonmatrixGetMap() **c.Char {
	return nil
}

/**
 * Get the index of the lastly "activated" button by the user (pressed, released, focused etc)
 * Useful in the `event_cb` to get the text of the button, check if hidden etc.
 * @param obj       pointer to button matrix object
 * @return          index of the last released button (LV_BUTTONMATRIX_BUTTON_NONE: if unset)
 */
// llgo:link (*LvObjT).LvButtonmatrixGetSelectedButton C.lv_buttonmatrix_get_selected_button
func (recv_ *LvObjT) LvButtonmatrixGetSelectedButton() c.Uint32T {
	return 0
}

/**
 * Get the button's text
 * @param obj       pointer to button matrix object
 * @param btn_id    the index a button not counting new line characters.
 * @return          text of btn_index` button
 */
// llgo:link (*LvObjT).LvButtonmatrixGetButtonText C.lv_buttonmatrix_get_button_text
func (recv_ *LvObjT) LvButtonmatrixGetButtonText(btn_id c.Uint32T) *c.Char {
	return nil
}

/**
 * Get the whether a control value is enabled or disabled for button of a button matrix
 * @param obj       pointer to a button matrix object
 * @param btn_id    the index of a button not counting new line characters.
 * @param ctrl      control values to check (ORed value can be used)
 * @return          true: the control attribute is enabled false: disabled
 */
// llgo:link (*LvObjT).LvButtonmatrixHasButtonCtrl C.lv_buttonmatrix_has_button_ctrl
func (recv_ *LvObjT) LvButtonmatrixHasButtonCtrl(btn_id c.Uint32T, ctrl LvButtonmatrixCtrlT) bool {
	return false
}

/**
 * Tell whether "one check" mode is enabled or not.
 * @param obj       Button matrix object
 * @return          true: "one check" mode is enabled; false: disabled
 */
// llgo:link (*LvObjT).LvButtonmatrixGetOneChecked C.lv_buttonmatrix_get_one_checked
func (recv_ *LvObjT) LvButtonmatrixGetOneChecked() bool {
	return false
}

type LvLabelLongModeT c.Int

const (
	LV_LABEL_LONG_MODE_WRAP            LvLabelLongModeT = 0
	LV_LABEL_LONG_MODE_DOTS            LvLabelLongModeT = 1
	LV_LABEL_LONG_MODE_SCROLL          LvLabelLongModeT = 2
	LV_LABEL_LONG_MODE_SCROLL_CIRCULAR LvLabelLongModeT = 3
	LV_LABEL_LONG_MODE_CLIP            LvLabelLongModeT = 4
)

/**
 * Create a label object
 * @param parent    pointer to an object, it will be the parent of the new label.
 * @return          pointer to the created button
 */
// llgo:link (*LvObjT).LvLabelCreate C.lv_label_create
func (recv_ *LvObjT) LvLabelCreate() *LvObjT {
	return nil
}

/**
 * Set a new text for a label. Memory will be allocated to store the text by the label.
 * @param obj           pointer to a label object
 * @param text          '\0' terminated character string. NULL to refresh with the current text.
 */
// llgo:link (*LvObjT).LvLabelSetText C.lv_label_set_text
func (recv_ *LvObjT) LvLabelSetText(text *c.Char) {
}

/**
 * Set a new formatted text for a label. Memory will be allocated to store the text by the label.
 * @param obj           pointer to a label object
 * @param fmt           `printf`-like format
 *
 * Example:
 * @code
 * lv_label_set_text_fmt(label1, "%d user", user_num);
 * @endcode
 */
// llgo:link (*LvObjT).LvLabelSetTextFmt C.lv_label_set_text_fmt
func (recv_ *LvObjT) LvLabelSetTextFmt(fmt *c.Char, __llgo_va_list ...interface{}) {
}

/**
 * Set a static text. It will not be saved by the label so the 'text' variable
 * has to be 'alive' while the label exists.
 * @param obj           pointer to a label object
 * @param text          pointer to a text. NULL to refresh with the current text.
 */
// llgo:link (*LvObjT).LvLabelSetTextStatic C.lv_label_set_text_static
func (recv_ *LvObjT) LvLabelSetTextStatic(text *c.Char) {
}

/**
 * Set the behavior of the label with text longer than the object size
 * @param obj           pointer to a label object
 * @param long_mode     the new mode from 'lv_label_long_mode' enum.
 *                      In LV_LONG_WRAP/DOT/SCROLL/SCROLL_CIRC the size of the label should be set AFTER this function
 */
// llgo:link (*LvObjT).LvLabelSetLongMode C.lv_label_set_long_mode
func (recv_ *LvObjT) LvLabelSetLongMode(long_mode LvLabelLongModeT) {
}

/**
 * Set where text selection should start
 * @param obj       pointer to a label object
 * @param index     character index from where selection should start. `LV_LABEL_TEXT_SELECTION_OFF` for no selection
 */
// llgo:link (*LvObjT).LvLabelSetTextSelectionStart C.lv_label_set_text_selection_start
func (recv_ *LvObjT) LvLabelSetTextSelectionStart(index c.Uint32T) {
}

/**
 * Set where text selection should end
 * @param obj       pointer to a label object
 * @param index     character index where selection should end. `LV_LABEL_TEXT_SELECTION_OFF` for no selection
 */
// llgo:link (*LvObjT).LvLabelSetTextSelectionEnd C.lv_label_set_text_selection_end
func (recv_ *LvObjT) LvLabelSetTextSelectionEnd(index c.Uint32T) {
}

/**
 * Enable the recoloring by in-line commands
 * @param obj           pointer to a label object
 * @param en            true: enable recoloring, false: disable
 * Example: "This is a #ff0000 red# word"
 */
// llgo:link (*LvObjT).LvLabelSetRecolor C.lv_label_set_recolor
func (recv_ *LvObjT) LvLabelSetRecolor(en bool) {
}

/**
 * Get the text of a label
 * @param obj       pointer to a label object
 * @return          the text of the label
 */
// llgo:link (*LvObjT).LvLabelGetText C.lv_label_get_text
func (recv_ *LvObjT) LvLabelGetText() *c.Char {
	return nil
}

/**
 * Get the long mode of a label
 * @param obj       pointer to a label object
 * @return          the current long mode
 */
// llgo:link (*LvObjT).LvLabelGetLongMode C.lv_label_get_long_mode
func (recv_ *LvObjT) LvLabelGetLongMode() LvLabelLongModeT {
	return 0
}

/**
 * Get the relative x and y coordinates of a letter
 * @param obj       pointer to a label object
 * @param char_id   index of the character [0 ... text length - 1].
 *                  Expressed in character index, not byte index (different in UTF-8)
 * @param pos       store the result here (E.g. index = 0 gives 0;0 coordinates if the text if aligned to the left)
 */
// llgo:link (*LvObjT).LvLabelGetLetterPos C.lv_label_get_letter_pos
func (recv_ *LvObjT) LvLabelGetLetterPos(char_id c.Uint32T, pos *LvPointT) {
}

/**
 * Get the index of letter on a relative point of a label.
 * @param obj       pointer to label object
 * @param pos_in    pointer to point with coordinates on a the label
 * @param bidi      whether to use bidi processed
 * @return          The index of the letter on the 'pos_p' point (E.g. on 0;0 is the 0. letter if aligned to the left)
 *                  Expressed in character index and not byte index (different in UTF-8)
 */
// llgo:link (*LvObjT).LvLabelGetLetterOn C.lv_label_get_letter_on
func (recv_ *LvObjT) LvLabelGetLetterOn(pos_in *LvPointT, bidi bool) c.Uint32T {
	return 0
}

/**
 * Check if a character is drawn under a point.
 * @param obj       pointer to a label object
 * @param pos       Point to check for character under
 * @return          whether a character is drawn under the point
 */
// llgo:link (*LvObjT).LvLabelIsCharUnderPos C.lv_label_is_char_under_pos
func (recv_ *LvObjT) LvLabelIsCharUnderPos(pos *LvPointT) bool {
	return false
}

/**
 * @brief Get the selection start index.
 * @param obj       pointer to a label object.
 * @return          selection start index. `LV_LABEL_TEXT_SELECTION_OFF` if nothing is selected.
 */
// llgo:link (*LvObjT).LvLabelGetTextSelectionStart C.lv_label_get_text_selection_start
func (recv_ *LvObjT) LvLabelGetTextSelectionStart() c.Uint32T {
	return 0
}

/**
 * @brief Get the selection end index.
 * @param obj       pointer to a label object.
 * @return          selection end index. `LV_LABEL_TXT_SEL_OFF` if nothing is selected.
 */
// llgo:link (*LvObjT).LvLabelGetTextSelectionEnd C.lv_label_get_text_selection_end
func (recv_ *LvObjT) LvLabelGetTextSelectionEnd() c.Uint32T {
	return 0
}

/**
 * @brief Get the recoloring attribute
 * @param obj       pointer to a label object.
 * @return          true: recoloring is enabled, false: recoloring is disabled
 */
// llgo:link (*LvObjT).LvLabelGetRecolor C.lv_label_get_recolor
func (recv_ *LvObjT) LvLabelGetRecolor() bool {
	return false
}

/**
 * Insert a text to a label. The label text cannot be static.
 * @param obj       pointer to a label object
 * @param pos       character index to insert. Expressed in character index and not byte index.
 *                  0: before first char. LV_LABEL_POS_LAST: after last char.
 * @param txt       pointer to the text to insert
 */
// llgo:link (*LvObjT).LvLabelInsText C.lv_label_ins_text
func (recv_ *LvObjT) LvLabelInsText(pos c.Uint32T, txt *c.Char) {
}

/**
 * Delete characters from a label. The label text cannot be static.
 * @param obj       pointer to a label object
 * @param pos       character index from where to cut. Expressed in character index and not byte index.
 *                  0: start in front of the first character
 * @param cnt       number of characters to cut
 */
// llgo:link (*LvObjT).LvLabelCutText C.lv_label_cut_text
func (recv_ *LvObjT) LvLabelCutText(pos c.Uint32T, cnt c.Uint32T) {
}

type LvBarModeT c.Int

const (
	LV_BAR_MODE_NORMAL      LvBarModeT = 0
	LV_BAR_MODE_SYMMETRICAL LvBarModeT = 1
	LV_BAR_MODE_RANGE       LvBarModeT = 2
)

type LvBarOrientationT c.Int

const (
	LV_BAR_ORIENTATION_AUTO       LvBarOrientationT = 0
	LV_BAR_ORIENTATION_HORIZONTAL LvBarOrientationT = 1
	LV_BAR_ORIENTATION_VERTICAL   LvBarOrientationT = 2
)

/**
 * Create a bar object
 * @param parent        pointer to an object, it will be the parent of the new bar
 * @return              pointer to the created bar
 */
// llgo:link (*LvObjT).LvBarCreate C.lv_bar_create
func (recv_ *LvObjT) LvBarCreate() *LvObjT {
	return nil
}

/**
 * Set a new value on the bar
 * @param obj           pointer to a bar object
 * @param value         new value
 * @param anim          LV_ANIM_ON: set the value with an animation; LV_ANIM_OFF: change the value immediately
 */
// llgo:link (*LvObjT).LvBarSetValue C.lv_bar_set_value
func (recv_ *LvObjT) LvBarSetValue(value c.Int32T, anim LvAnimEnableT) {
}

/**
 * Set a new start value on the bar
 * @param obj             pointer to a bar object
 * @param start_value     new start value
 * @param anim            LV_ANIM_ON: set the value with an animation; LV_ANIM_OFF: change the value immediately
 */
// llgo:link (*LvObjT).LvBarSetStartValue C.lv_bar_set_start_value
func (recv_ *LvObjT) LvBarSetStartValue(start_value c.Int32T, anim LvAnimEnableT) {
}

/**
 * Set minimum and the maximum values of a bar
 * @param obj       pointer to the bar object
 * @param min       minimum value
 * @param max       maximum value
 * @note If min is greater than max, the drawing direction becomes to the opposite direction.
 */
// llgo:link (*LvObjT).LvBarSetRange C.lv_bar_set_range
func (recv_ *LvObjT) LvBarSetRange(min c.Int32T, max c.Int32T) {
}

/**
 * Set minimum value of a bar
 * @param obj       pointer to the bar object
 * @param min       minimum value
 */
// llgo:link (*LvObjT).LvBarSetMinValue C.lv_bar_set_min_value
func (recv_ *LvObjT) LvBarSetMinValue(min c.Int32T) {
}

/**
 * Set maximum value of a bar
 * @param obj       pointer to the bar object
 * @param max       maximum value
 */
// llgo:link (*LvObjT).LvBarSetMaxValue C.lv_bar_set_max_value
func (recv_ *LvObjT) LvBarSetMaxValue(max c.Int32T) {
}

/**
 * Set the type of bar.
 * @param obj       pointer to bar object
 * @param mode      bar type from `lv_bar_mode_t`
 */
// llgo:link (*LvObjT).LvBarSetMode C.lv_bar_set_mode
func (recv_ *LvObjT) LvBarSetMode(mode LvBarModeT) {
}

/**
 * Set the orientation of bar.
 * @param obj           pointer to bar object
 * @param orientation   bar orientation from `lv_bar_orientation_t`
 */
// llgo:link (*LvObjT).LvBarSetOrientation C.lv_bar_set_orientation
func (recv_ *LvObjT) LvBarSetOrientation(orientation LvBarOrientationT) {
}

/**
 * Get the value of a bar
 * @param obj       pointer to a bar object
 * @return          the value of the bar
 */
// llgo:link (*LvObjT).LvBarGetValue C.lv_bar_get_value
func (recv_ *LvObjT) LvBarGetValue() c.Int32T {
	return 0
}

/**
 * Get the start value of a bar
 * @param obj       pointer to a bar object
 * @return          the start value of the bar
 */
// llgo:link (*LvObjT).LvBarGetStartValue C.lv_bar_get_start_value
func (recv_ *LvObjT) LvBarGetStartValue() c.Int32T {
	return 0
}

/**
 * Get the minimum value of a bar
 * @param obj       pointer to a bar object
 * @return          the minimum value of the bar
 */
// llgo:link (*LvObjT).LvBarGetMinValue C.lv_bar_get_min_value
func (recv_ *LvObjT) LvBarGetMinValue() c.Int32T {
	return 0
}

/**
 * Get the maximum value of a bar
 * @param obj       pointer to a bar object
 * @return          the maximum value of the bar
 */
// llgo:link (*LvObjT).LvBarGetMaxValue C.lv_bar_get_max_value
func (recv_ *LvObjT) LvBarGetMaxValue() c.Int32T {
	return 0
}

/**
 * Get the type of bar.
 * @param obj       pointer to bar object
 * @return          bar type from `lv_bar_mode_t`
 */
// llgo:link (*LvObjT).LvBarGetMode C.lv_bar_get_mode
func (recv_ *LvObjT) LvBarGetMode() LvBarModeT {
	return 0
}

/**
 * Get the orientation of bar.
 * @param obj       pointer to bar object
 * @return          bar orientation from `lv_bar_orientation_t`
 */
// llgo:link (*LvObjT).LvBarGetOrientation C.lv_bar_get_orientation
func (recv_ *LvObjT) LvBarGetOrientation() LvBarOrientationT {
	return 0
}

/**
 * Give the bar is in symmetrical mode or not
 * @param obj       pointer to bar object
 * @return          true: in symmetrical mode false : not in
 */
// llgo:link (*LvObjT).LvBarIsSymmetrical C.lv_bar_is_symmetrical
func (recv_ *LvObjT) LvBarIsSymmetrical() bool {
	return false
}

type LvSliderModeT c.Int

const (
	LV_SLIDER_MODE_NORMAL      LvSliderModeT = 0
	LV_SLIDER_MODE_SYMMETRICAL LvSliderModeT = 1
	LV_SLIDER_MODE_RANGE       LvSliderModeT = 2
)

type LvSliderOrientationT c.Int

const (
	LV_SLIDER_ORIENTATION_AUTO       LvSliderOrientationT = 0
	LV_SLIDER_ORIENTATION_HORIZONTAL LvSliderOrientationT = 1
	LV_SLIDER_ORIENTATION_VERTICAL   LvSliderOrientationT = 2
)

/**
 * Create a slider object
 * @param parent    pointer to an object, it will be the parent of the new slider.
 * @return          pointer to the created slider
 */
// llgo:link (*LvObjT).LvSliderCreate C.lv_slider_create
func (recv_ *LvObjT) LvSliderCreate() *LvObjT {
	return nil
}

/**
 * Set a new value on the slider
 * @param obj       pointer to a slider object
 * @param value     the new value
 * @param anim      LV_ANIM_ON: set the value with an animation; LV_ANIM_OFF: change the value immediately
 */
// llgo:link (*LvObjT).LvSliderSetValue C.lv_slider_set_value
func (recv_ *LvObjT) LvSliderSetValue(value c.Int32T, anim LvAnimEnableT) {
}

/**
 * Set a new value for the left knob of a slider
 * @param obj       pointer to a slider object
 * @param value     new value
 * @param anim      LV_ANIM_ON: set the value with an animation; LV_ANIM_OFF: change the value immediately
 */
// llgo:link (*LvObjT).LvSliderSetStartValue C.lv_slider_set_start_value
func (recv_ *LvObjT) LvSliderSetStartValue(value c.Int32T, anim LvAnimEnableT) {
}

/**
 * Set the minimum and the maximum values of a bar
 * @param obj       pointer to the slider object
 * @param min       minimum value
 * @param max       maximum value
 */
// llgo:link (*LvObjT).LvSliderSetRange C.lv_slider_set_range
func (recv_ *LvObjT) LvSliderSetRange(min c.Int32T, max c.Int32T) {
}

/**
 * Set the minimum values of a bar
 * @param obj       pointer to the slider object
 * @param min       minimum value
 */
// llgo:link (*LvObjT).LvSliderSetMinValue C.lv_slider_set_min_value
func (recv_ *LvObjT) LvSliderSetMinValue(min c.Int32T) {
}

/**
 * Set the maximum values of a bar
 * @param obj       pointer to the slider object
 * @param max       maximum value
 */
// llgo:link (*LvObjT).LvSliderSetMaxValue C.lv_slider_set_max_value
func (recv_ *LvObjT) LvSliderSetMaxValue(max c.Int32T) {
}

/**
 * Set the mode of slider.
 * @param obj       pointer to a slider object
 * @param mode      the mode of the slider. See `lv_slider_mode_t`
 */
// llgo:link (*LvObjT).LvSliderSetMode C.lv_slider_set_mode
func (recv_ *LvObjT) LvSliderSetMode(mode LvSliderModeT) {
}

/**
 * Set the orientation of slider.
 * @param obj           pointer to a slider object
 * @param orientation   slider  orientation from `lv_slider_orientation_t`
 */
// llgo:link (*LvObjT).LvSliderSetOrientation C.lv_slider_set_orientation
func (recv_ *LvObjT) LvSliderSetOrientation(orientation LvSliderOrientationT) {
}

/**
 * Get the value of the main knob of a slider
 * @param obj       pointer to a slider object
 * @return          the value of the main knob of the slider
 */
// llgo:link (*LvObjT).LvSliderGetValue C.lv_slider_get_value
func (recv_ *LvObjT) LvSliderGetValue() c.Int32T {
	return 0
}

/**
 * Get the value of the left knob of a slider
 * @param obj       pointer to a slider object
 * @return          the value of the left knob of the slider
 */
// llgo:link (*LvObjT).LvSliderGetLeftValue C.lv_slider_get_left_value
func (recv_ *LvObjT) LvSliderGetLeftValue() c.Int32T {
	return 0
}

/**
 * Get the minimum value of a slider
 * @param obj       pointer to a slider object
 * @return          the minimum value of the slider
 */
// llgo:link (*LvObjT).LvSliderGetMinValue C.lv_slider_get_min_value
func (recv_ *LvObjT) LvSliderGetMinValue() c.Int32T {
	return 0
}

/**
 * Get the maximum value of a slider
 * @param obj       pointer to a slider object
 * @return          the maximum value of the slider
 */
// llgo:link (*LvObjT).LvSliderGetMaxValue C.lv_slider_get_max_value
func (recv_ *LvObjT) LvSliderGetMaxValue() c.Int32T {
	return 0
}

/**
 * Give the slider is being dragged or not
 * @param obj       pointer to a slider object
 * @return          true: drag in progress false: not dragged
 */
// llgo:link (*LvObjT).LvSliderIsDragged C.lv_slider_is_dragged
func (recv_ *LvObjT) LvSliderIsDragged() bool {
	return false
}

/**
 * Get the mode of the slider.
 * @param slider       pointer to a slider object
 * @return          see `lv_slider_mode_t`
 */
// llgo:link (*LvObjT).LvSliderGetMode C.lv_slider_get_mode
func (recv_ *LvObjT) LvSliderGetMode() LvSliderModeT {
	return 0
}

/**
 * Get the orientation of slider.
 * @param obj       pointer to a slider object
 * @return          slider orientation from `lv_slider_orientation_t`
 */
// llgo:link (*LvObjT).LvSliderGetOrientation C.lv_slider_get_orientation
func (recv_ *LvObjT) LvSliderGetOrientation() LvSliderOrientationT {
	return 0
}

/**
 * Give the slider is in symmetrical mode or not
 * @param obj       pointer to slider object
 * @return          true: in symmetrical mode false : not in
 */
// llgo:link (*LvObjT).LvSliderIsSymmetrical C.lv_slider_is_symmetrical
func (recv_ *LvObjT) LvSliderIsSymmetrical() bool {
	return false
}

type LvSwitchOrientationT c.Int

const (
	LV_SWITCH_ORIENTATION_AUTO       LvSwitchOrientationT = 0
	LV_SWITCH_ORIENTATION_HORIZONTAL LvSwitchOrientationT = 1
	LV_SWITCH_ORIENTATION_VERTICAL   LvSwitchOrientationT = 2
)

/**
 * Create a switch object
 * @param parent    pointer to an object, it will be the parent of the new switch
 * @return          pointer to the created switch
 */
// llgo:link (*LvObjT).LvSwitchCreate C.lv_switch_create
func (recv_ *LvObjT) LvSwitchCreate() *LvObjT {
	return nil
}

/**
 * Set the orientation of switch.
 * @param obj           pointer to switch object
 * @param orientation   switch orientation from `lv_switch_orientation_t`
 */
// llgo:link (*LvObjT).LvSwitchSetOrientation C.lv_switch_set_orientation
func (recv_ *LvObjT) LvSwitchSetOrientation(orientation LvSwitchOrientationT) {
}

/**
 * Get the orientation of switch.
 * @param obj       pointer to switch object
 * @return          switch orientation from ::lv_switch_orientation_t
 */
// llgo:link (*LvObjT).LvSwitchGetOrientation C.lv_switch_get_orientation
func (recv_ *LvObjT) LvSwitchGetOrientation() LvSwitchOrientationT {
	return 0
}

/**
 * Represents a date on the calendar object (platform-agnostic).
 */

type LvCalendarDateT struct {
	Year  c.Uint16T
	Month c.Uint8T
	Day   c.Uint8T
}

/**
 * Create a calendar widget
 * @param parent    pointer to an object, it will be the parent of the new calendar
 * @return          pointer the created calendar
 */
// llgo:link (*LvObjT).LvCalendarCreate C.lv_calendar_create
func (recv_ *LvObjT) LvCalendarCreate() *LvObjT {
	return nil
}

/**
 * Set the today's year, month and day at once
 * @param obj  pointer to a calendar object
 * @param year      today's year
 * @param month     today's month [1..12]
 * @param day       today's day [1..31]
 */
// llgo:link (*LvObjT).LvCalendarSetTodayDate C.lv_calendar_set_today_date
func (recv_ *LvObjT) LvCalendarSetTodayDate(year c.Uint32T, month c.Uint32T, day c.Uint32T) {
}

/**
 * Set the today's year
 * @param obj  pointer to a calendar object
 * @param year      today's year
 */
// llgo:link (*LvObjT).LvCalendarSetTodayYear C.lv_calendar_set_today_year
func (recv_ *LvObjT) LvCalendarSetTodayYear(year c.Uint32T) {
}

/**
 * Set the today's year
 * @param obj  pointer to a calendar object
 * @param month     today's month [1..12]
 */
// llgo:link (*LvObjT).LvCalendarSetTodayMonth C.lv_calendar_set_today_month
func (recv_ *LvObjT) LvCalendarSetTodayMonth(month c.Uint32T) {
}

/**
 * Set the today's year
 * @param obj  pointer to a calendar object
 * @param day       today's day [1..31]
 */
// llgo:link (*LvObjT).LvCalendarSetTodayDay C.lv_calendar_set_today_day
func (recv_ *LvObjT) LvCalendarSetTodayDay(day c.Uint32T) {
}

/**
 * Set the currently shown year and month at once
 * @param obj           pointer to a calendar object
 * @param year          shown year
 * @param month         shown month [1..12]
 */
// llgo:link (*LvObjT).LvCalendarSetMonthShown C.lv_calendar_set_month_shown
func (recv_ *LvObjT) LvCalendarSetMonthShown(year c.Uint32T, month c.Uint32T) {
}

/**
 * Set the currently shown year
 * @param obj           pointer to a calendar object
 * @param year          shown year
 */
// llgo:link (*LvObjT).LvCalendarSetShownYear C.lv_calendar_set_shown_year
func (recv_ *LvObjT) LvCalendarSetShownYear(year c.Uint32T) {
}

/**
 * Set the currently shown month
 * @param obj           pointer to a calendar object
 * @param month         shown month [1..12]
 */
// llgo:link (*LvObjT).LvCalendarSetShownMonth C.lv_calendar_set_shown_month
func (recv_ *LvObjT) LvCalendarSetShownMonth(month c.Uint32T) {
}

/**
 * Set the highlighted dates
 * @param obj           pointer to a calendar object
 * @param highlighted   pointer to an `lv_calendar_date_t` array containing the dates.
 *                      Only the pointer will be saved so this variable can't be local which will be destroyed later.
 * @param date_num number of dates in the array
 */
// llgo:link (*LvObjT).LvCalendarSetHighlightedDates C.lv_calendar_set_highlighted_dates
func (recv_ *LvObjT) LvCalendarSetHighlightedDates(highlighted *LvCalendarDateT, date_num c.SizeT) {
}

/**
 * Set the name of the days
 * @param obj           pointer to a calendar object
 * @param day_names     pointer to an array with the names.
 *                      E.g. `const char * days[7] = {"Sun", "Mon", ...}`
 *                      Only the pointer will be saved so this variable can't be local which will be destroyed later.
 */
// llgo:link (*LvObjT).LvCalendarSetDayNames C.lv_calendar_set_day_names
func (recv_ *LvObjT) LvCalendarSetDayNames(day_names **c.Char) {
}

/**
 * Get the button matrix object of the calendar.
 * It shows the dates and day names.
 * @param obj       pointer to a calendar object
 * @return          pointer to a the button matrix
 */
// llgo:link (*LvObjT).LvCalendarGetBtnmatrix C.lv_calendar_get_btnmatrix
func (recv_ *LvObjT) LvCalendarGetBtnmatrix() *LvObjT {
	return nil
}

/**
 * Get the today's date
 * @param calendar  pointer to a calendar object
 * @return          return pointer to an `lv_calendar_date_t` variable containing the date of today.
 */
// llgo:link (*LvObjT).LvCalendarGetTodayDate C.lv_calendar_get_today_date
func (recv_ *LvObjT) LvCalendarGetTodayDate() *LvCalendarDateT {
	return nil
}

/**
 * Get the currently showed
 * @param calendar  pointer to a calendar object
 * @return          pointer to an `lv_calendar_date_t` variable containing the date is being shown.
 */
// llgo:link (*LvObjT).LvCalendarGetShowedDate C.lv_calendar_get_showed_date
func (recv_ *LvObjT) LvCalendarGetShowedDate() *LvCalendarDateT {
	return nil
}

/**
 * Get the highlighted dates
 * @param calendar  pointer to a calendar object
 * @return          pointer to an `lv_calendar_date_t` array containing the dates.
 */
// llgo:link (*LvObjT).LvCalendarGetHighlightedDates C.lv_calendar_get_highlighted_dates
func (recv_ *LvObjT) LvCalendarGetHighlightedDates() *LvCalendarDateT {
	return nil
}

/**
 * Get the number of the highlighted dates
 * @param calendar  pointer to a calendar object
 * @return          number of highlighted days
 */
// llgo:link (*LvObjT).LvCalendarGetHighlightedDatesNum C.lv_calendar_get_highlighted_dates_num
func (recv_ *LvObjT) LvCalendarGetHighlightedDatesNum() c.SizeT {
	return 0
}

/**
 * Get the currently pressed day
 * @param calendar  pointer to a calendar object
 * @param date      store the pressed date here
 * @return          LV_RESULT_OK: there is a valid pressed date
 *                  LV_RESULT_INVALID: there is no pressed data
 */
// llgo:link (*LvObjT).LvCalendarGetPressedDate C.lv_calendar_get_pressed_date
func (recv_ *LvObjT) LvCalendarGetPressedDate(date *LvCalendarDateT) LvResultT {
	return 0
}

/**
 * Create a calendar header with drop-drowns to select the year and month
 * @param parent    pointer to a calendar object.
 * @return          the created header
 */
// llgo:link (*LvObjT).LvCalendarAddHeaderArrow C.lv_calendar_add_header_arrow
func (recv_ *LvObjT) LvCalendarAddHeaderArrow() *LvObjT {
	return nil
}

/**
 * Create a calendar header with drop-drowns to select the year and month
 * @param parent    pointer to a calendar object.
 * @return          the created header
 */
// llgo:link (*LvObjT).LvCalendarAddHeaderDropdown C.lv_calendar_add_header_dropdown
func (recv_ *LvObjT) LvCalendarAddHeaderDropdown() *LvObjT {
	return nil
}

/**
 * Sets a custom calendar year list
 * @param parent        pointer to a calendar object
 * @param years_list    pointer to an const char array with the years list, see lv_dropdown set_options for more information.
 *                      E.g. `const char * years = "2023\n2022\n2021\n2020\n2019"
 *                      Only the pointer will be saved so this variable can't be local which will be destroyed later.
 */
// llgo:link (*LvObjT).LvCalendarHeaderDropdownSetYearList C.lv_calendar_header_dropdown_set_year_list
func (recv_ *LvObjT) LvCalendarHeaderDropdownSetYearList(years_list *c.Char) {
}

type LvImagebuttonStateT c.Int

const (
	LV_IMAGEBUTTON_STATE_RELEASED         LvImagebuttonStateT = 0
	LV_IMAGEBUTTON_STATE_PRESSED          LvImagebuttonStateT = 1
	LV_IMAGEBUTTON_STATE_DISABLED         LvImagebuttonStateT = 2
	LV_IMAGEBUTTON_STATE_CHECKED_RELEASED LvImagebuttonStateT = 3
	LV_IMAGEBUTTON_STATE_CHECKED_PRESSED  LvImagebuttonStateT = 4
	LV_IMAGEBUTTON_STATE_CHECKED_DISABLED LvImagebuttonStateT = 5
	LV_IMAGEBUTTON_STATE_NUM              LvImagebuttonStateT = 6
)

/**
 * Create an image button object
 * @param parent pointer to an object, it will be the parent of the new image button
 * @return pointer to the created image button
 */
// llgo:link (*LvObjT).LvImagebuttonCreate C.lv_imagebutton_create
func (recv_ *LvObjT) LvImagebuttonCreate() *LvObjT {
	return nil
}

/**
 * Set images for a state of the image button
 * @param imagebutton   pointer to an image button object
 * @param state         for which state set the new image
 * @param src_left      pointer to an image source for the left side of the button (a C array or path to
 * a file)
 * @param src_mid       pointer to an image source for the middle of the button (ideally 1px wide) (a C
 * array or path to a file)
 * @param src_right     pointer to an image source for the right side of the button (a C array or path
 * to a file)
 */
// llgo:link (*LvObjT).LvImagebuttonSetSrc C.lv_imagebutton_set_src
func (recv_ *LvObjT) LvImagebuttonSetSrc(state LvImagebuttonStateT, src_left c.Pointer, src_mid c.Pointer, src_right c.Pointer) {
}

/**
 * Use this function instead of `lv_obj_add/remove_state` to set a state manually
 * @param imagebutton   pointer to an image button object
 * @param state         the new state
 */
// llgo:link (*LvObjT).LvImagebuttonSetState C.lv_imagebutton_set_state
func (recv_ *LvObjT) LvImagebuttonSetState(state LvImagebuttonStateT) {
}

/**
 * Get the left image in a given state
 * @param imagebutton   pointer to an image button object
 * @param state         the state where to get the image (from `lv_button_state_t`) `
 * @return              pointer to the left image source (a C array or path to a file)
 */
// llgo:link (*LvObjT).LvImagebuttonGetSrcLeft C.lv_imagebutton_get_src_left
func (recv_ *LvObjT) LvImagebuttonGetSrcLeft(state LvImagebuttonStateT) c.Pointer {
	return nil
}

/**
 * Get the middle image in a given state
 * @param imagebutton   pointer to an image button object
 * @param state         the state where to get the image (from `lv_button_state_t`) `
 * @return              pointer to the middle image source (a C array or path to a file)
 */
// llgo:link (*LvObjT).LvImagebuttonGetSrcMiddle C.lv_imagebutton_get_src_middle
func (recv_ *LvObjT) LvImagebuttonGetSrcMiddle(state LvImagebuttonStateT) c.Pointer {
	return nil
}

/**
 * Get the right image in a given state
 * @param imagebutton   pointer to an image button object
 * @param state         the state where to get the image (from `lv_button_state_t`) `
 * @return              pointer to the left image source (a C array or path to a file)
 */
// llgo:link (*LvObjT).LvImagebuttonGetSrcRight C.lv_imagebutton_get_src_right
func (recv_ *LvObjT) LvImagebuttonGetSrcRight(state LvImagebuttonStateT) c.Pointer {
	return nil
}

type X_lvPartTextareaIdT c.Int

const LV_PART_TEXTAREA_PLACEHOLDER X_lvPartTextareaIdT = 524288

/**
 * Create a text area object
 * @param parent    pointer to an object, it will be the parent of the new text area
 * @return          pointer to the created text area
 */
// llgo:link (*LvObjT).LvTextareaCreate C.lv_textarea_create
func (recv_ *LvObjT) LvTextareaCreate() *LvObjT {
	return nil
}

/**
 * Insert a character to the current cursor position.
 * To add a wide char, e.g. 'Á' use `lv_text_encoded_conv_wc('Á')`
 * @param obj       pointer to a text area object
 * @param c         a character (e.g. 'a')
 */
// llgo:link (*LvObjT).LvTextareaAddChar C.lv_textarea_add_char
func (recv_ *LvObjT) LvTextareaAddChar(c c.Uint32T) {
}

/**
 * Insert a text to the current cursor position
 * @param obj       pointer to a text area object
 * @param txt       a '\0' terminated string to insert
 */
// llgo:link (*LvObjT).LvTextareaAddText C.lv_textarea_add_text
func (recv_ *LvObjT) LvTextareaAddText(txt *c.Char) {
}

/**
 * Delete a the left character from the current cursor position
 * @param obj       pointer to a text area object
 */
// llgo:link (*LvObjT).LvTextareaDeleteChar C.lv_textarea_delete_char
func (recv_ *LvObjT) LvTextareaDeleteChar() {
}

/**
 * Delete the right character from the current cursor position
 * @param obj       pointer to a text area object
 */
// llgo:link (*LvObjT).LvTextareaDeleteCharForward C.lv_textarea_delete_char_forward
func (recv_ *LvObjT) LvTextareaDeleteCharForward() {
}

/**
 * Set the text of a text area
 * @param obj       pointer to a text area object
 * @param txt       pointer to the text
 */
// llgo:link (*LvObjT).LvTextareaSetText C.lv_textarea_set_text
func (recv_ *LvObjT) LvTextareaSetText(txt *c.Char) {
}

/**
 * Set the placeholder text of a text area
 * @param obj       pointer to a text area object
 * @param txt       pointer to the text
 */
// llgo:link (*LvObjT).LvTextareaSetPlaceholderText C.lv_textarea_set_placeholder_text
func (recv_ *LvObjT) LvTextareaSetPlaceholderText(txt *c.Char) {
}

/**
 * Set the cursor position
 * @param obj       pointer to a text area object
 * @param pos       the new cursor position in character index
 *                  < 0 : index from the end of the text
 *                  LV_TEXTAREA_CURSOR_LAST: go after the last character
 */
// llgo:link (*LvObjT).LvTextareaSetCursorPos C.lv_textarea_set_cursor_pos
func (recv_ *LvObjT) LvTextareaSetCursorPos(pos c.Int32T) {
}

/**
 * Enable/Disable the positioning of the cursor by clicking the text on the text area.
 * @param obj       pointer to a text area object
 * @param en        true: enable click positions; false: disable
 */
// llgo:link (*LvObjT).LvTextareaSetCursorClickPos C.lv_textarea_set_cursor_click_pos
func (recv_ *LvObjT) LvTextareaSetCursorClickPos(en bool) {
}

/**
 * Enable/Disable password mode
 * @param obj       pointer to a text area object
 * @param en        true: enable, false: disable
 */
// llgo:link (*LvObjT).LvTextareaSetPasswordMode C.lv_textarea_set_password_mode
func (recv_ *LvObjT) LvTextareaSetPasswordMode(en bool) {
}

/**
 * Set the replacement characters to show in password mode
 * @param obj       pointer to a text area object
 * @param bullet    pointer to the replacement text
 */
// llgo:link (*LvObjT).LvTextareaSetPasswordBullet C.lv_textarea_set_password_bullet
func (recv_ *LvObjT) LvTextareaSetPasswordBullet(bullet *c.Char) {
}

/**
 * Configure the text area to one line or back to normal
 * @param obj       pointer to a text area object
 * @param en        true: one line, false: normal
 */
// llgo:link (*LvObjT).LvTextareaSetOneLine C.lv_textarea_set_one_line
func (recv_ *LvObjT) LvTextareaSetOneLine(en bool) {
}

/**
 * Set a list of characters. Only these characters will be accepted by the text area
 * @param obj       pointer to a text area object
 * @param list      list of characters. Only the pointer is saved. E.g. "+-.,0123456789"
 */
// llgo:link (*LvObjT).LvTextareaSetAcceptedChars C.lv_textarea_set_accepted_chars
func (recv_ *LvObjT) LvTextareaSetAcceptedChars(list *c.Char) {
}

/**
 * Set max length of a Text Area.
 * @param obj       pointer to a text area object
 * @param num       the maximal number of characters can be added (`lv_textarea_set_text` ignores it)
 */
// llgo:link (*LvObjT).LvTextareaSetMaxLength C.lv_textarea_set_max_length
func (recv_ *LvObjT) LvTextareaSetMaxLength(num c.Uint32T) {
}

/**
 * In `LV_EVENT_INSERT` the text which planned to be inserted can be replaced by another text.
 * It can be used to add automatic formatting to the text area.
 * @param obj       pointer to a text area object
 * @param txt       pointer to a new string to insert. If `""` no text will be added.
 *                  The variable must be live after the `event_cb` exists. (Should be `global` or `static`)
 */
// llgo:link (*LvObjT).LvTextareaSetInsertReplace C.lv_textarea_set_insert_replace
func (recv_ *LvObjT) LvTextareaSetInsertReplace(txt *c.Char) {
}

/**
 * Enable/disable selection mode.
 * @param obj       pointer to a text area object
 * @param en        true or false to enable/disable selection mode
 */
// llgo:link (*LvObjT).LvTextareaSetTextSelection C.lv_textarea_set_text_selection
func (recv_ *LvObjT) LvTextareaSetTextSelection(en bool) {
}

/**
 * Set how long show the password before changing it to '*'
 * @param obj       pointer to a text area object
 * @param time      show time in milliseconds. 0: hide immediately.
 */
// llgo:link (*LvObjT).LvTextareaSetPasswordShowTime C.lv_textarea_set_password_show_time
func (recv_ *LvObjT) LvTextareaSetPasswordShowTime(time c.Uint32T) {
}

/**
 * @deprecated Use the normal text_align style property instead
 * Set the label's alignment.
 * It sets where the label is aligned (in one line mode it can be smaller than the text area)
 * and how the lines of the area align in case of multiline text area
 * @param obj       pointer to a text area object
 * @param align     the align mode from ::lv_text_align_t
 */
// llgo:link (*LvObjT).LvTextareaSetAlign C.lv_textarea_set_align
func (recv_ *LvObjT) LvTextareaSetAlign(align LvTextAlignT) {
}

/**
 * Get the text of a text area. In password mode it gives the real text (not '*'s).
 * @param obj       pointer to a text area object
 * @return          pointer to the text
 */
// llgo:link (*LvObjT).LvTextareaGetText C.lv_textarea_get_text
func (recv_ *LvObjT) LvTextareaGetText() *c.Char {
	return nil
}

/**
 * Get the placeholder text of a text area
 * @param obj       pointer to a text area object
 * @return          pointer to the text
 */
// llgo:link (*LvObjT).LvTextareaGetPlaceholderText C.lv_textarea_get_placeholder_text
func (recv_ *LvObjT) LvTextareaGetPlaceholderText() *c.Char {
	return nil
}

/**
 * Get the label of a text area
 * @param obj       pointer to a text area object
 * @return          pointer to the label object
 */
// llgo:link (*LvObjT).LvTextareaGetLabel C.lv_textarea_get_label
func (recv_ *LvObjT) LvTextareaGetLabel() *LvObjT {
	return nil
}

/**
 * Get the current cursor position in character index
 * @param obj       pointer to a text area object
 * @return          the cursor position
 */
// llgo:link (*LvObjT).LvTextareaGetCursorPos C.lv_textarea_get_cursor_pos
func (recv_ *LvObjT) LvTextareaGetCursorPos() c.Uint32T {
	return 0
}

/**
 * Get whether the cursor click positioning is enabled or not.
 * @param obj       pointer to a text area object
 * @return          true: enable click positions; false: disable
 */
// llgo:link (*LvObjT).LvTextareaGetCursorClickPos C.lv_textarea_get_cursor_click_pos
func (recv_ *LvObjT) LvTextareaGetCursorClickPos() bool {
	return false
}

/**
 * Get the password mode attribute
 * @param obj       pointer to a text area object
 * @return          true: password mode is enabled, false: disabled
 */
// llgo:link (*LvObjT).LvTextareaGetPasswordMode C.lv_textarea_get_password_mode
func (recv_ *LvObjT) LvTextareaGetPasswordMode() bool {
	return false
}

/**
 * Get the replacement characters to show in password mode
 * @param obj       pointer to a text area object
 * @return          pointer to the replacement text
 */
// llgo:link (*LvObjT).LvTextareaGetPasswordBullet C.lv_textarea_get_password_bullet
func (recv_ *LvObjT) LvTextareaGetPasswordBullet() *c.Char {
	return nil
}

/**
 * Get the one line configuration attribute
 * @param obj       pointer to a text area object
 * @return          true: one line configuration is enabled, false: disabled
 */
// llgo:link (*LvObjT).LvTextareaGetOneLine C.lv_textarea_get_one_line
func (recv_ *LvObjT) LvTextareaGetOneLine() bool {
	return false
}

/**
 * Get a list of accepted characters.
 * @param obj       pointer to a text area object
 * @return          list of accented characters.
 */
// llgo:link (*LvObjT).LvTextareaGetAcceptedChars C.lv_textarea_get_accepted_chars
func (recv_ *LvObjT) LvTextareaGetAcceptedChars() *c.Char {
	return nil
}

/**
 * Get max length of a Text Area.
 * @param obj       pointer to a text area object
 * @return          the maximal number of characters to be add
 */
// llgo:link (*LvObjT).LvTextareaGetMaxLength C.lv_textarea_get_max_length
func (recv_ *LvObjT) LvTextareaGetMaxLength() c.Uint32T {
	return 0
}

/**
 * Find whether text is selected or not.
 * @param obj       pointer to a text area object
 * @return          whether text is selected or not
 */
// llgo:link (*LvObjT).LvTextareaTextIsSelected C.lv_textarea_text_is_selected
func (recv_ *LvObjT) LvTextareaTextIsSelected() bool {
	return false
}

/**
 * Find whether selection mode is enabled.
 * @param obj       pointer to a text area object
 * @return          true: selection mode is enabled, false: disabled
 */
// llgo:link (*LvObjT).LvTextareaGetTextSelection C.lv_textarea_get_text_selection
func (recv_ *LvObjT) LvTextareaGetTextSelection() bool {
	return false
}

/**
 * Set how long show the password before changing it to '*'
 * @param obj       pointer to a text area object
 * @return          show time in milliseconds. 0: hide immediately.
 */
// llgo:link (*LvObjT).LvTextareaGetPasswordShowTime C.lv_textarea_get_password_show_time
func (recv_ *LvObjT) LvTextareaGetPasswordShowTime() c.Uint32T {
	return 0
}

/**
 * Get a the character from the current cursor position
 * @param obj       pointer to a text area object
 * @return          a the character or 0
 */
// llgo:link (*LvObjT).LvTextareaGetCurrentChar C.lv_textarea_get_current_char
func (recv_ *LvObjT) LvTextareaGetCurrentChar() c.Uint32T {
	return 0
}

/**
 * Clear the selection on the text area.
 * @param obj       pointer to a text area object
 */
// llgo:link (*LvObjT).LvTextareaClearSelection C.lv_textarea_clear_selection
func (recv_ *LvObjT) LvTextareaClearSelection() {
}

/**
 * Move the cursor one character right
 * @param obj       pointer to a text area object
 */
// llgo:link (*LvObjT).LvTextareaCursorRight C.lv_textarea_cursor_right
func (recv_ *LvObjT) LvTextareaCursorRight() {
}

/**
 * Move the cursor one character left
 * @param obj       pointer to a text area object
 */
// llgo:link (*LvObjT).LvTextareaCursorLeft C.lv_textarea_cursor_left
func (recv_ *LvObjT) LvTextareaCursorLeft() {
}

/**
 * Move the cursor one line down
 * @param obj       pointer to a text area object
 */
// llgo:link (*LvObjT).LvTextareaCursorDown C.lv_textarea_cursor_down
func (recv_ *LvObjT) LvTextareaCursorDown() {
}

/**
 * Move the cursor one line up
 * @param obj       pointer to a text area object
 */
// llgo:link (*LvObjT).LvTextareaCursorUp C.lv_textarea_cursor_up
func (recv_ *LvObjT) LvTextareaCursorUp() {
}

type LvTableCellCtrlT c.Int

const (
	LV_TABLE_CELL_CTRL_NONE        LvTableCellCtrlT = 0
	LV_TABLE_CELL_CTRL_MERGE_RIGHT LvTableCellCtrlT = 1
	LV_TABLE_CELL_CTRL_TEXT_CROP   LvTableCellCtrlT = 2
	LV_TABLE_CELL_CTRL_CUSTOM_1    LvTableCellCtrlT = 16
	LV_TABLE_CELL_CTRL_CUSTOM_2    LvTableCellCtrlT = 32
	LV_TABLE_CELL_CTRL_CUSTOM_3    LvTableCellCtrlT = 64
	LV_TABLE_CELL_CTRL_CUSTOM_4    LvTableCellCtrlT = 128
)

/**
 * Create a table object
 * @param parent        pointer to an object, it will be the parent of the new table
 * @return              pointer to the created table
 */
// llgo:link (*LvObjT).LvTableCreate C.lv_table_create
func (recv_ *LvObjT) LvTableCreate() *LvObjT {
	return nil
}

/**
 * Set the value of a cell.
 * @param obj           pointer to a Table object
 * @param row           id of the row [0 .. row_cnt -1]
 * @param col           id of the column [0 .. col_cnt -1]
 * @param txt           text to display in the cell. It will be copied and saved so this variable is not required after this function call.
 * @note                New roes/columns are added automatically if required
 */
// llgo:link (*LvObjT).LvTableSetCellValue C.lv_table_set_cell_value
func (recv_ *LvObjT) LvTableSetCellValue(row c.Uint32T, col c.Uint32T, txt *c.Char) {
}

/**
 * Set the value of a cell.  Memory will be allocated to store the text by the table.
 * @param obj           pointer to a Table object
 * @param row           id of the row [0 .. row_cnt -1]
 * @param col           id of the column [0 .. col_cnt -1]
 * @param fmt           `printf`-like format
 * @note                New roes/columns are added automatically if required
 */
// llgo:link (*LvObjT).LvTableSetCellValueFmt C.lv_table_set_cell_value_fmt
func (recv_ *LvObjT) LvTableSetCellValueFmt(row c.Uint32T, col c.Uint32T, fmt *c.Char, __llgo_va_list ...interface{}) {
}

/**
 * Set the number of rows
 * @param obj           table pointer to a Table object
 * @param row_cnt       number of rows
 */
// llgo:link (*LvObjT).LvTableSetRowCount C.lv_table_set_row_count
func (recv_ *LvObjT) LvTableSetRowCount(row_cnt c.Uint32T) {
}

/**
 * Set the number of columns
 * @param obj       table pointer to a Table object
 * @param col_cnt   number of columns.
 */
// llgo:link (*LvObjT).LvTableSetColumnCount C.lv_table_set_column_count
func (recv_ *LvObjT) LvTableSetColumnCount(col_cnt c.Uint32T) {
}

/**
 * Set the width of a column
 * @param obj       table pointer to a Table object
 * @param col_id    id of the column [0 .. LV_TABLE_COL_MAX -1]
 * @param w         width of the column
 */
// llgo:link (*LvObjT).LvTableSetColumnWidth C.lv_table_set_column_width
func (recv_ *LvObjT) LvTableSetColumnWidth(col_id c.Uint32T, w c.Int32T) {
}

/**
 * Add control bits to the cell.
 * @param obj       pointer to a Table object
 * @param row       id of the row [0 .. row_cnt -1]
 * @param col       id of the column [0 .. col_cnt -1]
 * @param ctrl      OR-ed values from ::lv_table_cell_ctrl_t
 */
// llgo:link (*LvObjT).LvTableSetCellCtrl C.lv_table_set_cell_ctrl
func (recv_ *LvObjT) LvTableSetCellCtrl(row c.Uint32T, col c.Uint32T, ctrl LvTableCellCtrlT) {
}

/**
 * Clear control bits of the cell.
 * @param obj       pointer to a Table object
 * @param row       id of the row [0 .. row_cnt -1]
 * @param col       id of the column [0 .. col_cnt -1]
 * @param ctrl      OR-ed values from ::lv_table_cell_ctrl_t
 */
// llgo:link (*LvObjT).LvTableClearCellCtrl C.lv_table_clear_cell_ctrl
func (recv_ *LvObjT) LvTableClearCellCtrl(row c.Uint32T, col c.Uint32T, ctrl LvTableCellCtrlT) {
}

/**
 * Add custom user data to the cell.
 * @param obj       pointer to a Table object
 * @param row       id of the row [0 .. row_cnt -1]
 * @param col       id of the column [0 .. col_cnt -1]
 * @param user_data pointer to the new user_data.
 *                  Should be allocated by `lv_malloc`,
 *                  and it will be freed automatically when the table is deleted or
 *                  when the cell is dropped due to lower row or column count.
 */
// llgo:link (*LvObjT).LvTableSetCellUserData C.lv_table_set_cell_user_data
func (recv_ *LvObjT) LvTableSetCellUserData(row c.Uint16T, col c.Uint16T, user_data c.Pointer) {
}

/**
 * Set the selected cell
 * @param obj       pointer to a table object
 * @param row       id of the cell row to select
 * @param col       id of the cell column to select
 */
// llgo:link (*LvObjT).LvTableSetSelectedCell C.lv_table_set_selected_cell
func (recv_ *LvObjT) LvTableSetSelectedCell(row c.Uint16T, col c.Uint16T) {
}

/**
 * Get the value of a cell.
 * @param obj       pointer to a Table object
 * @param row       id of the row [0 .. row_cnt -1]
 * @param col       id of the column [0 .. col_cnt -1]
 * @return          text in the cell
 */
// llgo:link (*LvObjT).LvTableGetCellValue C.lv_table_get_cell_value
func (recv_ *LvObjT) LvTableGetCellValue(row c.Uint32T, col c.Uint32T) *c.Char {
	return nil
}

/**
 * Get the number of rows.
 * @param obj       table pointer to a Table object
 * @return          number of rows.
 */
// llgo:link (*LvObjT).LvTableGetRowCount C.lv_table_get_row_count
func (recv_ *LvObjT) LvTableGetRowCount() c.Uint32T {
	return 0
}

/**
 * Get the number of columns.
 * @param obj       table pointer to a Table object
 * @return          number of columns.
 */
// llgo:link (*LvObjT).LvTableGetColumnCount C.lv_table_get_column_count
func (recv_ *LvObjT) LvTableGetColumnCount() c.Uint32T {
	return 0
}

/**
 * Get the width of a column
 * @param obj       table pointer to a Table object
 * @param col       id of the column [0 .. LV_TABLE_COL_MAX -1]
 * @return          width of the column
 */
// llgo:link (*LvObjT).LvTableGetColumnWidth C.lv_table_get_column_width
func (recv_ *LvObjT) LvTableGetColumnWidth(col c.Uint32T) c.Int32T {
	return 0
}

/**
 * Get whether a cell has the control bits
 * @param obj       pointer to a Table object
 * @param row       id of the row [0 .. row_cnt -1]
 * @param col       id of the column [0 .. col_cnt -1]
 * @param ctrl      OR-ed values from ::lv_table_cell_ctrl_t
 * @return          true: all control bits are set; false: not all control bits are set
 */
// llgo:link (*LvObjT).LvTableHasCellCtrl C.lv_table_has_cell_ctrl
func (recv_ *LvObjT) LvTableHasCellCtrl(row c.Uint32T, col c.Uint32T, ctrl LvTableCellCtrlT) bool {
	return false
}

/**
 * Get the selected cell (pressed and or focused)
 * @param obj       pointer to a table object
 * @param row       pointer to variable to store the selected row (LV_TABLE_CELL_NONE: if no cell selected)
 * @param col       pointer to variable to store the selected column  (LV_TABLE_CELL_NONE: if no cell selected)
 */
// llgo:link (*LvObjT).LvTableGetSelectedCell C.lv_table_get_selected_cell
func (recv_ *LvObjT) LvTableGetSelectedCell(row *c.Uint32T, col *c.Uint32T) {
}

/**
 * Get custom user data to the cell.
 * @param obj       pointer to a Table object
 * @param row       id of the row [0 .. row_cnt -1]
 * @param col       id of the column [0 .. col_cnt -1]
 */
// llgo:link (*LvObjT).LvTableGetCellUserData C.lv_table_get_cell_user_data
func (recv_ *LvObjT) LvTableGetCellUserData(row c.Uint16T, col c.Uint16T) c.Pointer {
	return nil
}

/**
 * Create a check box object
 * @param parent    pointer to an object, it will be the parent of the new button
 * @return          pointer to the created check box
 */
// llgo:link (*LvObjT).LvCheckboxCreate C.lv_checkbox_create
func (recv_ *LvObjT) LvCheckboxCreate() *LvObjT {
	return nil
}

/**
 * Set the text of a check box. `txt` will be copied and may be deallocated
 * after this function returns.
 * @param obj   pointer to a check box
 * @param txt   the text of the check box. NULL to refresh with the current text.
 */
// llgo:link (*LvObjT).LvCheckboxSetText C.lv_checkbox_set_text
func (recv_ *LvObjT) LvCheckboxSetText(txt *c.Char) {
}

/**
 * Set the text of a check box. `txt` must not be deallocated during the life
 * of this checkbox.
 * @param obj   pointer to a check box
 * @param txt   the text of the check box.
 */
// llgo:link (*LvObjT).LvCheckboxSetTextStatic C.lv_checkbox_set_text_static
func (recv_ *LvObjT) LvCheckboxSetTextStatic(txt *c.Char) {
}

/**
 * Get the text of a check box
 * @param obj   pointer to check box object
 * @return      pointer to the text of the check box
 */
// llgo:link (*LvObjT).LvCheckboxGetText C.lv_checkbox_get_text
func (recv_ *LvObjT) LvCheckboxGetText() *c.Char {
	return nil
}

type LvRollerModeT c.Int

const (
	LV_ROLLER_MODE_NORMAL   LvRollerModeT = 0
	LV_ROLLER_MODE_INFINITE LvRollerModeT = 1
)

/**
 * Create a roller object
 * @param parent    pointer to an object, it will be the parent of the new roller.
 * @return          pointer to the created roller
 */
// llgo:link (*LvObjT).LvRollerCreate C.lv_roller_create
func (recv_ *LvObjT) LvRollerCreate() *LvObjT {
	return nil
}

/**
 * Set the options on a roller
 * @param obj       pointer to roller object
 * @param options   a string with '\n' separated options. E.g. "One\nTwo\nThree"
 * @param mode      `LV_ROLLER_MODE_NORMAL` or `LV_ROLLER_MODE_INFINITE`
 */
// llgo:link (*LvObjT).LvRollerSetOptions C.lv_roller_set_options
func (recv_ *LvObjT) LvRollerSetOptions(options *c.Char, mode LvRollerModeT) {
}

/**
 * Set the selected option
 * @param obj       pointer to a roller object
 * @param sel_opt   index of the selected option (0 ... number of option - 1);
 * @param anim   LV_ANIM_ON: set with animation; LV_ANIM_OFF set immediately
 */
// llgo:link (*LvObjT).LvRollerSetSelected C.lv_roller_set_selected
func (recv_ *LvObjT) LvRollerSetSelected(sel_opt c.Uint32T, anim LvAnimEnableT) {
}

/**
 * Sets the given string as the selection on the roller. Does not alter the current selection on failure.
 * @param obj               pointer to roller object
 * @param sel_opt   pointer to the string you want to set as an option
 * @param anim          LV_ANIM_ON: set with animation; LV_ANIM_OFF set immediately
 * @return                  `true` if set successfully and `false` if the given string does not exist as an option in the roller
 */
// llgo:link (*LvObjT).LvRollerSetSelectedStr C.lv_roller_set_selected_str
func (recv_ *LvObjT) LvRollerSetSelectedStr(sel_opt *c.Char, anim LvAnimEnableT) bool {
	return false
}

/**
 * Set the height to show the given number of rows (options)
 * @param obj       pointer to a roller object
 * @param row_cnt   number of desired visible rows
 */
// llgo:link (*LvObjT).LvRollerSetVisibleRowCount C.lv_roller_set_visible_row_count
func (recv_ *LvObjT) LvRollerSetVisibleRowCount(row_cnt c.Uint32T) {
}

/**
 * Get the index of the selected option
 * @param obj       pointer to a roller object
 * @return          index of the selected option (0 ... number of option - 1);
 */
// llgo:link (*LvObjT).LvRollerGetSelected C.lv_roller_get_selected
func (recv_ *LvObjT) LvRollerGetSelected() c.Uint32T {
	return 0
}

/**
 * Get the current selected option as a string.
 * @param obj       pointer to roller object
 * @param buf       pointer to an array to store the string
 * @param buf_size  size of `buf` in bytes. 0: to ignore it.
 */
// llgo:link (*LvObjT).LvRollerGetSelectedStr C.lv_roller_get_selected_str
func (recv_ *LvObjT) LvRollerGetSelectedStr(buf *c.Char, buf_size c.Uint32T) {
}

/**
 * Get the options of a roller
 * @param obj       pointer to roller object
 * @return          the options separated by '\n'-s (E.g. "Option1\nOption2\nOption3")
 */
// llgo:link (*LvObjT).LvRollerGetOptions C.lv_roller_get_options
func (recv_ *LvObjT) LvRollerGetOptions() *c.Char {
	return nil
}

/**
 * Get the total number of options
 * @param obj   pointer to a roller object
 * @return      the total number of options
 */
// llgo:link (*LvObjT).LvRollerGetOptionCount C.lv_roller_get_option_count
func (recv_ *LvObjT) LvRollerGetOptionCount() c.Uint32T {
	return 0
}

/**
 * Create a window widget
 * @param parent    pointer to a parent widget
 * @return          the created window
 */
// llgo:link (*LvObjT).LvWinCreate C.lv_win_create
func (recv_ *LvObjT) LvWinCreate() *LvObjT {
	return nil
}

/**
 * Add a title to the window
 * @param obj       pointer to a window widget
 * @param txt       the text of the title
 * @return          the widget where the content of the title can be created
 */
// llgo:link (*LvObjT).LvWinAddTitle C.lv_win_add_title
func (recv_ *LvObjT) LvWinAddTitle(txt *c.Char) *LvObjT {
	return nil
}

/**
 * Add a button to the window
 * @param obj       pointer to a window widget
 * @param icon      an icon to be displayed on the button
 * @param btn_w     width of the button
 * @return          the widget where the content of the button can be created
 */
// llgo:link (*LvObjT).LvWinAddButton C.lv_win_add_button
func (recv_ *LvObjT) LvWinAddButton(icon c.Pointer, btn_w c.Int32T) *LvObjT {
	return nil
}

/**
 * Get the header of the window
 * @param win       pointer to a window widget
 * @return          the header of the window
 */
// llgo:link (*LvObjT).LvWinGetHeader C.lv_win_get_header
func (recv_ *LvObjT) LvWinGetHeader() *LvObjT {
	return nil
}

/**
 * Get the content of the window
 * @param win       pointer to a window widget
 * @return          the content of the window
 */
// llgo:link (*LvObjT).LvWinGetContent C.lv_win_get_content
func (recv_ *LvObjT) LvWinGetContent() *LvObjT {
	return nil
}

type LvKeyboardModeT c.Int

const (
	LV_KEYBOARD_MODE_TEXT_LOWER LvKeyboardModeT = 0
	LV_KEYBOARD_MODE_TEXT_UPPER LvKeyboardModeT = 1
	LV_KEYBOARD_MODE_SPECIAL    LvKeyboardModeT = 2
	LV_KEYBOARD_MODE_NUMBER     LvKeyboardModeT = 3
	LV_KEYBOARD_MODE_USER_1     LvKeyboardModeT = 4
	LV_KEYBOARD_MODE_USER_2     LvKeyboardModeT = 5
	LV_KEYBOARD_MODE_USER_3     LvKeyboardModeT = 6
	LV_KEYBOARD_MODE_USER_4     LvKeyboardModeT = 7
)

/**
 * Create a Keyboard object
 * @param parent    pointer to an object, it will be the parent of the new keyboard
 * @return          pointer to the created keyboard
 */
// llgo:link (*LvObjT).LvKeyboardCreate C.lv_keyboard_create
func (recv_ *LvObjT) LvKeyboardCreate() *LvObjT {
	return nil
}

/**
 * Assign a Text Area to the Keyboard. The pressed characters will be put there.
 * @param kb        pointer to a Keyboard object
 * @param ta        pointer to a Text Area object to write there
 */
// llgo:link (*LvObjT).LvKeyboardSetTextarea C.lv_keyboard_set_textarea
func (recv_ *LvObjT) LvKeyboardSetTextarea(ta *LvObjT) {
}

/**
 * Set a new a mode (text or number map)
 * @param kb        pointer to a Keyboard object
 * @param mode      the mode from 'lv_keyboard_mode_t'
 */
// llgo:link (*LvObjT).LvKeyboardSetMode C.lv_keyboard_set_mode
func (recv_ *LvObjT) LvKeyboardSetMode(mode LvKeyboardModeT) {
}

/**
 * Show the button title in a popover when pressed.
 * @param kb        pointer to a Keyboard object
 * @param en        whether "popovers" mode is enabled
 */
// llgo:link (*LvObjT).LvKeyboardSetPopovers C.lv_keyboard_set_popovers
func (recv_ *LvObjT) LvKeyboardSetPopovers(en bool) {
}

/**
 * Set a new map for the keyboard
 * @param kb        pointer to a Keyboard object
 * @param mode      keyboard map to alter 'lv_keyboard_mode_t'
 * @param map       pointer to a string array to describe the map.
 *                  See 'lv_buttonmatrix_set_map()' for more info.
 * @param ctrl_map  See 'lv_buttonmatrix_set_ctrl_map()' for more info.

 */
// llgo:link (*LvObjT).LvKeyboardSetMap C.lv_keyboard_set_map
func (recv_ *LvObjT) LvKeyboardSetMap(mode LvKeyboardModeT, map_ **c.Char, ctrl_map *LvButtonmatrixCtrlT) {
}

/**
 * Assign a Text Area to the Keyboard. The pressed characters will be put there.
 * @param kb        pointer to a Keyboard object
 * @return          pointer to the assigned Text Area object
 */
// llgo:link (*LvObjT).LvKeyboardGetTextarea C.lv_keyboard_get_textarea
func (recv_ *LvObjT) LvKeyboardGetTextarea() *LvObjT {
	return nil
}

/**
 * Set a new a mode (text or number map)
 * @param kb        pointer to a Keyboard object
 * @return          the current mode from 'lv_keyboard_mode_t'
 */
// llgo:link (*LvObjT).LvKeyboardGetMode C.lv_keyboard_get_mode
func (recv_ *LvObjT) LvKeyboardGetMode() LvKeyboardModeT {
	return 0
}

/**
 * Tell whether "popovers" mode is enabled or not.
 * @param obj       pointer to a Keyboard object
 * @return          true: "popovers" mode is enabled; false: disabled
 */
// llgo:link (*LvObjT).LvKeyboardGetPopovers C.lv_keyboard_get_popovers
func (recv_ *LvObjT) LvKeyboardGetPopovers() bool {
	return false
}

/**
 * Get the current map of a keyboard
 * @param kb        pointer to a keyboard object
 * @return          the current map
 */
// llgo:link (*LvObjT).LvKeyboardGetMapArray C.lv_keyboard_get_map_array
func (recv_ *LvObjT) LvKeyboardGetMapArray() **c.Char {
	return nil
}

/**
 * Get the index of the lastly "activated" button by the user (pressed, released, focused etc)
 * Useful in the `event_cb` to get the text of the button, check if hidden etc.
 * @param obj       pointer to button matrix object
 * @return          index of the last released button (LV_BUTTONMATRIX_BUTTON_NONE: if unset)
 */
// llgo:link (*LvObjT).LvKeyboardGetSelectedButton C.lv_keyboard_get_selected_button
func (recv_ *LvObjT) LvKeyboardGetSelectedButton() c.Uint32T {
	return 0
}

/**
 * Get the button's text
 * @param obj       pointer to button matrix object
 * @param btn_id    the index a button not counting new line characters.
 * @return          text of btn_index` button
 */
// llgo:link (*LvObjT).LvKeyboardGetButtonText C.lv_keyboard_get_button_text
func (recv_ *LvObjT) LvKeyboardGetButtonText(btn_id c.Uint32T) *c.Char {
	return nil
}

/**
 * Default keyboard event to add characters to the Text area and change the map.
 * If a custom `event_cb` is added to the keyboard this function can be called from it to handle the
 * button clicks
 * @param e the triggering event
 */
// llgo:link (*LvEventT).LvKeyboardDefEventCb C.lv_keyboard_def_event_cb
func (recv_ *LvEventT) LvKeyboardDefEventCb() {
}

/**
 * Create a line object
 * @param parent pointer to an object, it will be the parent of the new line
 * @return pointer to the created line
 */
// llgo:link (*LvObjT).LvLineCreate C.lv_line_create
func (recv_ *LvObjT) LvLineCreate() *LvObjT {
	return nil
}

/**
 * Set an array of points. The line object will connect these points.
 * @param obj           pointer to a line object
 * @param points        an array of points. Only the address is saved, so the array needs to be alive while the line exists
 * @param point_num     number of points in 'point_a'
 */
// llgo:link (*LvObjT).LvLineSetPoints C.lv_line_set_points
func (recv_ *LvObjT) LvLineSetPoints(points *LvPointPreciseT, point_num c.Uint32T) {
}

/**
 * Set a non-const array of points. Identical to `lv_line_set_points` except the array may be retrieved by `lv_line_get_points_mutable`.
 * @param obj           pointer to a line object
 * @param points        a non-const array of points. Only the address is saved, so the array needs to be alive while the line exists.
 * @param point_num     number of points in 'point_a'
 */
// llgo:link (*LvObjT).LvLineSetPointsMutable C.lv_line_set_points_mutable
func (recv_ *LvObjT) LvLineSetPointsMutable(points *LvPointPreciseT, point_num c.Uint32T) {
}

/**
 * Enable (or disable) the y coordinate inversion.
 * If enabled then y will be subtracted from the height of the object,
 * therefore the y = 0 coordinate will be on the bottom.
 * @param obj       pointer to a line object
 * @param en        true: enable the y inversion, false:disable the y inversion
 */
// llgo:link (*LvObjT).LvLineSetYInvert C.lv_line_set_y_invert
func (recv_ *LvObjT) LvLineSetYInvert(en bool) {
}

/**
 * Get the pointer to the array of points.
 * @param obj           pointer to a line object
 * @return              const pointer to the array of points
 */
// llgo:link (*LvObjT).LvLineGetPoints C.lv_line_get_points
func (recv_ *LvObjT) LvLineGetPoints() *LvPointPreciseT {
	return nil
}

/**
 * Get the number of points in the array of points.
 * @param obj           pointer to a line object
 * @return              number of points in array of points
 */
// llgo:link (*LvObjT).LvLineGetPointCount C.lv_line_get_point_count
func (recv_ *LvObjT) LvLineGetPointCount() c.Uint32T {
	return 0
}

/**
 * Check the mutability of the stored point array pointer.
 * @param obj           pointer to a line object
 * @return              true: the point array pointer is mutable, false: constant
 */
// llgo:link (*LvObjT).LvLineIsPointArrayMutable C.lv_line_is_point_array_mutable
func (recv_ *LvObjT) LvLineIsPointArrayMutable() bool {
	return false
}

/**
 * Get a pointer to the mutable array of points or NULL if it is not mutable
 * @param obj           pointer to a line object
 * @return              pointer to the array of points. NULL if not mutable.
 */
// llgo:link (*LvObjT).LvLineGetPointsMutable C.lv_line_get_points_mutable
func (recv_ *LvObjT) LvLineGetPointsMutable() *LvPointPreciseT {
	return nil
}

/**
 * Get the y inversion attribute
 * @param obj       pointer to a line object
 * @return          true: y inversion is enabled, false: disabled
 */
// llgo:link (*LvObjT).LvLineGetYInvert C.lv_line_get_y_invert
func (recv_ *LvObjT) LvLineGetYInvert() bool {
	return false
}

type LvAnimimgPartT c.Int

const LV_ANIM_IMAGE_PART_MAIN LvAnimimgPartT = 0

/**
 * Create an animation image objects
 * @param parent pointer to an object, it will be the parent of the new button
 * @return pointer to the created animation image object
 */
// llgo:link (*LvObjT).LvAnimimgCreate C.lv_animimg_create
func (recv_ *LvObjT) LvAnimimgCreate() *LvObjT {
	return nil
}

/**
 * Set the image animation images source.
 * @param obj       pointer to an animation image object
 * @param dsc       pointer to a series images
 * @param num       images' number
 */
// llgo:link (*LvObjT).LvAnimimgSetSrc C.lv_animimg_set_src
func (recv_ *LvObjT) LvAnimimgSetSrc(dsc *c.Pointer, num c.SizeT) {
}

/**
 *  Set the images source for flip playback of animation image.
 * @param obj       pointer to an animation image object
 * @param dsc       pointer to a series images
 * @param num       images' number
 */
// llgo:link (*LvObjT).LvAnimimgSetSrcReverse C.lv_animimg_set_src_reverse
func (recv_ *LvObjT) LvAnimimgSetSrcReverse(dsc *c.Pointer, num c.SizeT) {
}

/**
 * Startup the image animation.
 * @param obj   pointer to an animation image object
 */
// llgo:link (*LvObjT).LvAnimimgStart C.lv_animimg_start
func (recv_ *LvObjT) LvAnimimgStart() {
}

/**
 * Delete the image animation.
 * @param obj   pointer to an animation image object
 */
// llgo:link (*LvObjT).LvAnimimgDelete C.lv_animimg_delete
func (recv_ *LvObjT) LvAnimimgDelete() bool {
	return false
}

/**
 * Set the image animation duration time. unit:ms
 * @param obj       pointer to an animation image object
 * @param duration  the duration in milliseconds
 */
// llgo:link (*LvObjT).LvAnimimgSetDuration C.lv_animimg_set_duration
func (recv_ *LvObjT) LvAnimimgSetDuration(duration c.Uint32T) {
}

/**
 * Set the image animation repeatedly play times.
 * @param obj       pointer to an animation image object
 * @param count     the number of times to repeat the animation
 */
// llgo:link (*LvObjT).LvAnimimgSetRepeatCount C.lv_animimg_set_repeat_count
func (recv_ *LvObjT) LvAnimimgSetRepeatCount(count c.Uint32T) {
}

/**
 * Make the image animation to play back to when the forward direction is ready.
 * @param obj   pointer to an animation image object
 * @param duration   the duration of the playback image animation in milliseconds. 0: disable playback
 */
// llgo:link (*LvObjT).LvAnimimgSetReverseDuration C.lv_animimg_set_reverse_duration
func (recv_ *LvObjT) LvAnimimgSetReverseDuration(duration c.Uint32T) {
}

/**
 * Make the image animation to play back to when the forward direction is ready.
 * @param obj   pointer to an animation image object
 * @param duration   delay in milliseconds before starting the playback image animation.
 */
// llgo:link (*LvObjT).LvAnimimgSetReverseDelay C.lv_animimg_set_reverse_delay
func (recv_ *LvObjT) LvAnimimgSetReverseDelay(duration c.Uint32T) {
}

/**
 * Set a function call when the animation image really starts (considering `delay`)
 * @param obj   pointer to an animation image object
 * @param start_cb   a function call when the animation is start
 */
// llgo:link (*LvObjT).LvAnimimgSetStartCb C.lv_animimg_set_start_cb
func (recv_ *LvObjT) LvAnimimgSetStartCb(start_cb LvAnimStartCbT) {
}

/**
 * Set a function call when the animation is completed
 * @param obj pointer to an animation image object
 * @param completed_cb  a function call when the animation is completed
 */
// llgo:link (*LvObjT).LvAnimimgSetCompletedCb C.lv_animimg_set_completed_cb
func (recv_ *LvObjT) LvAnimimgSetCompletedCb(completed_cb LvAnimCompletedCbT) {
}

/**
 * Get the image animation images source.
 * @param obj   pointer to an animation image object
 * @return a     pointer that will point to a series images
 */
// llgo:link (*LvObjT).LvAnimimgGetSrc C.lv_animimg_get_src
func (recv_ *LvObjT) LvAnimimgGetSrc() *c.Pointer {
	return nil
}

/**
 * Get the image animation images source.
 * @param obj   pointer to an animation image object
 * @return      the number of source images
 */
// llgo:link (*LvObjT).LvAnimimgGetSrcCount C.lv_animimg_get_src_count
func (recv_ *LvObjT) LvAnimimgGetSrcCount() c.Uint8T {
	return 0
}

/**
 * Get the image animation duration time. unit:ms
 * @param obj   pointer to an animation image object
 * @return      the animation duration time
 */
// llgo:link (*LvObjT).LvAnimimgGetDuration C.lv_animimg_get_duration
func (recv_ *LvObjT) LvAnimimgGetDuration() c.Uint32T {
	return 0
}

/**
 * Get the image animation repeat play times.
 * @param obj   pointer to an animation image object
 * @return      the repeat count
 */
// llgo:link (*LvObjT).LvAnimimgGetRepeatCount C.lv_animimg_get_repeat_count
func (recv_ *LvObjT) LvAnimimgGetRepeatCount() c.Uint32T {
	return 0
}

/**
 * Get the image animation underlying animation.
 * @param obj   pointer to an animation image object
 * @return      the animation reference
 */
// llgo:link (*LvObjT).LvAnimimgGetAnim C.lv_animimg_get_anim
func (recv_ *LvObjT) LvAnimimgGetAnim() *LvAnimT {
	return nil
}

/**
 * Create a drop-down list object
 * @param parent pointer to an object, it will be the parent of the new drop-down list
 * @return pointer to the created drop-down list
 */
// llgo:link (*LvObjT).LvDropdownCreate C.lv_dropdown_create
func (recv_ *LvObjT) LvDropdownCreate() *LvObjT {
	return nil
}

/**
 * Set text of the drop-down list's button.
 * If set to `NULL` the selected option's text will be displayed on the button.
 * If set to a specific text then that text will be shown regardless of the selected option.
 * @param obj       pointer to a drop-down list object
 * @param txt       the text as a string (Only its pointer is saved)
 */
// llgo:link (*LvObjT).LvDropdownSetText C.lv_dropdown_set_text
func (recv_ *LvObjT) LvDropdownSetText(txt *c.Char) {
}

/**
 * Set the options in a drop-down list from a string.
 * The options will be copied and saved in the object so the `options` can be destroyed after calling this function
 * @param obj       pointer to drop-down list object
 * @param options   a string with '\n' separated options. E.g. "One\nTwo\nThree"
 */
// llgo:link (*LvObjT).LvDropdownSetOptions C.lv_dropdown_set_options
func (recv_ *LvObjT) LvDropdownSetOptions(options *c.Char) {
}

/**
 * Set the options in a drop-down list from a static string (global, static or dynamically allocated).
 * Only the pointer of the option string will be saved.
 * @param obj       pointer to drop-down list object
 * @param options   a static string with '\n' separated options. E.g. "One\nTwo\nThree"
 */
// llgo:link (*LvObjT).LvDropdownSetOptionsStatic C.lv_dropdown_set_options_static
func (recv_ *LvObjT) LvDropdownSetOptionsStatic(options *c.Char) {
}

/**
 * Add an options to a drop-down list from a string.  Only works for non-static options.
 * @param obj       pointer to drop-down list object
 * @param option    a string without '\n'. E.g. "Four"
 * @param pos       the insert position, indexed from 0, LV_DROPDOWN_POS_LAST = end of string
 */
// llgo:link (*LvObjT).LvDropdownAddOption C.lv_dropdown_add_option
func (recv_ *LvObjT) LvDropdownAddOption(option *c.Char, pos c.Uint32T) {
}

/**
 * Clear all options in a drop-down list.  Works with both static and dynamic options.
 * @param obj       pointer to drop-down list object
 */
// llgo:link (*LvObjT).LvDropdownClearOptions C.lv_dropdown_clear_options
func (recv_ *LvObjT) LvDropdownClearOptions() {
}

/**
 * Set the selected option
 * @param obj       pointer to drop-down list object
 * @param sel_opt   id of the selected option (0 ... number of option - 1);
 */
// llgo:link (*LvObjT).LvDropdownSetSelected C.lv_dropdown_set_selected
func (recv_ *LvObjT) LvDropdownSetSelected(sel_opt c.Uint32T) {
}

/**
 * Set the direction of the a drop-down list
 * @param obj       pointer to a drop-down list object
 * @param dir       LV_DIR_LEFT/RIGHT/TOP/BOTTOM
 */
// llgo:link (*LvObjT).LvDropdownSetDir C.lv_dropdown_set_dir
func (recv_ *LvObjT) LvDropdownSetDir(dir LvDirT) {
}

/**
 * Set an arrow or other symbol to display when on drop-down list's button. Typically a down caret or arrow.
 * @param obj       pointer to drop-down list object
 * @param symbol    a text like `LV_SYMBOL_DOWN`, an image (pointer or path) or NULL to not draw symbol icon
 * @note angle and zoom transformation can be applied if the symbol is an image.
 * E.g. when drop down is checked (opened) rotate the symbol by 180 degree
 */
// llgo:link (*LvObjT).LvDropdownSetSymbol C.lv_dropdown_set_symbol
func (recv_ *LvObjT) LvDropdownSetSymbol(symbol c.Pointer) {
}

/**
 * Set whether the selected option in the list should be highlighted or not
 * @param obj       pointer to drop-down list object
 * @param en        true: highlight enabled; false: disabled
 */
// llgo:link (*LvObjT).LvDropdownSetSelectedHighlight C.lv_dropdown_set_selected_highlight
func (recv_ *LvObjT) LvDropdownSetSelectedHighlight(en bool) {
}

/**
 * Get the list of a drop-down to allow styling or other modifications
 * @param obj       pointer to a drop-down list object
 * @return          pointer to the list of the drop-down
 */
// llgo:link (*LvObjT).LvDropdownGetList C.lv_dropdown_get_list
func (recv_ *LvObjT) LvDropdownGetList() *LvObjT {
	return nil
}

/**
 * Get text of the drop-down list's button.
 * @param obj   pointer to a drop-down list object
 * @return      the text as string, `NULL` if no text
 */
// llgo:link (*LvObjT).LvDropdownGetText C.lv_dropdown_get_text
func (recv_ *LvObjT) LvDropdownGetText() *c.Char {
	return nil
}

/**
 * Get the options of a drop-down list
 * @param obj       pointer to drop-down list object
 * @return          the options separated by '\n'-s (E.g. "Option1\nOption2\nOption3")
 */
// llgo:link (*LvObjT).LvDropdownGetOptions C.lv_dropdown_get_options
func (recv_ *LvObjT) LvDropdownGetOptions() *c.Char {
	return nil
}

/**
 * Get the index of the selected option
 * @param obj       pointer to drop-down list object
 * @return          index of the selected option (0 ... number of option - 1);
 */
// llgo:link (*LvObjT).LvDropdownGetSelected C.lv_dropdown_get_selected
func (recv_ *LvObjT) LvDropdownGetSelected() c.Uint32T {
	return 0
}

/**
 * Get the total number of options
 * @param obj       pointer to drop-down list object
 * @return          the total number of options in the list
 */
// llgo:link (*LvObjT).LvDropdownGetOptionCount C.lv_dropdown_get_option_count
func (recv_ *LvObjT) LvDropdownGetOptionCount() c.Uint32T {
	return 0
}

/**
 * Get the current selected option as a string
 * @param obj       pointer to drop-down object
 * @param buf       pointer to an array to store the string
 * @param buf_size  size of `buf` in bytes. 0: to ignore it.
 */
// llgo:link (*LvObjT).LvDropdownGetSelectedStr C.lv_dropdown_get_selected_str
func (recv_ *LvObjT) LvDropdownGetSelectedStr(buf *c.Char, buf_size c.Uint32T) {
}

/**
 * Get the index of an option.
 * @param obj       pointer to drop-down object
 * @param option    an option as string
 * @return          index of `option` in the list of all options. -1 if not found.
 */
// llgo:link (*LvObjT).LvDropdownGetOptionIndex C.lv_dropdown_get_option_index
func (recv_ *LvObjT) LvDropdownGetOptionIndex(option *c.Char) c.Int32T {
	return 0
}

/**
 * Get the symbol on the drop-down list. Typically a down caret or arrow.
 * @param obj       pointer to drop-down list object
 * @return          the symbol or NULL if not enabled
 */
// llgo:link (*LvObjT).LvDropdownGetSymbol C.lv_dropdown_get_symbol
func (recv_ *LvObjT) LvDropdownGetSymbol() *c.Char {
	return nil
}

/**
 * Get whether the selected option in the list should be highlighted or not
 * @param obj       pointer to drop-down list object
 * @return          true: highlight enabled; false: disabled
 */
// llgo:link (*LvObjT).LvDropdownGetSelectedHighlight C.lv_dropdown_get_selected_highlight
func (recv_ *LvObjT) LvDropdownGetSelectedHighlight() bool {
	return false
}

/**
 * Get the direction of the drop-down list
 * @param obj       pointer to a drop-down list object
 * @return          LV_DIR_LEF/RIGHT/TOP/BOTTOM
 */
// llgo:link (*LvObjT).LvDropdownGetDir C.lv_dropdown_get_dir
func (recv_ *LvObjT) LvDropdownGetDir() LvDirT {
	return 0
}

/**
 * Open the drop.down list
 * @param dropdown_obj       pointer to drop-down list object
 */
// llgo:link (*LvObjT).LvDropdownOpen C.lv_dropdown_open
func (recv_ *LvObjT) LvDropdownOpen() {
}

/**
 * Close (Collapse) the drop-down list
 * @param obj       pointer to drop-down list object
 */
// llgo:link (*LvObjT).LvDropdownClose C.lv_dropdown_close
func (recv_ *LvObjT) LvDropdownClose() {
}

/**
 * Tells whether the list is opened or not
 * @param obj       pointer to a drop-down list object
 * @return          true if the list os opened
 */
// llgo:link (*LvObjT).LvDropdownIsOpen C.lv_dropdown_is_open
func (recv_ *LvObjT) LvDropdownIsOpen() bool {
	return false
}

type LvMenuModeHeaderT c.Int

const (
	LV_MENU_HEADER_TOP_FIXED    LvMenuModeHeaderT = 0
	LV_MENU_HEADER_TOP_UNFIXED  LvMenuModeHeaderT = 1
	LV_MENU_HEADER_BOTTOM_FIXED LvMenuModeHeaderT = 2
)

type LvMenuModeRootBackButtonT c.Int

const (
	LV_MENU_ROOT_BACK_BUTTON_DISABLED LvMenuModeRootBackButtonT = 0
	LV_MENU_ROOT_BACK_BUTTON_ENABLED  LvMenuModeRootBackButtonT = 1
)

/**
 * Create a menu object
 * @param parent    pointer to an object, it will be the parent of the new menu
 * @return          pointer to the created menu
 */
// llgo:link (*LvObjT).LvMenuCreate C.lv_menu_create
func (recv_ *LvObjT) LvMenuCreate() *LvObjT {
	return nil
}

/**
 * Create a menu page object.
 *
 * This call inserts the new page under menu->storage as its parent, which is itself a
 * child of the menu, so the resulting object hierarchy is: menu => storage => new_page
 * where `storage` is a Base Widget.
 * @param menu      pointer to menu object.
 * @param title     pointer to text for title in header (NULL to not display title)
 * @return          pointer to the created menu page
 */
// llgo:link (*LvObjT).LvMenuPageCreate C.lv_menu_page_create
func (recv_ *LvObjT) LvMenuPageCreate(title *c.Char) *LvObjT {
	return nil
}

/**
 * Create a menu cont object
 * @param parent    pointer to a menu page object, it will be the parent of the new menu cont object
 * @return          pointer to the created menu cont
 */
// llgo:link (*LvObjT).LvMenuContCreate C.lv_menu_cont_create
func (recv_ *LvObjT) LvMenuContCreate() *LvObjT {
	return nil
}

/**
 * Create a menu section object
 * @param parent    pointer to a menu page object, it will be the parent of the new menu section object
 * @return          pointer to the created menu section
 */
// llgo:link (*LvObjT).LvMenuSectionCreate C.lv_menu_section_create
func (recv_ *LvObjT) LvMenuSectionCreate() *LvObjT {
	return nil
}

/**
 * Create a menu separator object
 * @param parent    pointer to a menu page object, it will be the parent of the new menu separator object
 * @return          pointer to the created menu separator
 */
// llgo:link (*LvObjT).LvMenuSeparatorCreate C.lv_menu_separator_create
func (recv_ *LvObjT) LvMenuSeparatorCreate() *LvObjT {
	return nil
}

/*=====================
 * Setter functions
 *====================*/
/**
 * Set menu page to display in main
 * @param obj       pointer to the menu
 * @param page      pointer to the menu page to set (NULL to clear main and clear menu history)
 */
// llgo:link (*LvObjT).LvMenuSetPage C.lv_menu_set_page
func (recv_ *LvObjT) LvMenuSetPage(page *LvObjT) {
}

/**
 * Set menu page title
 * @param page      pointer to the menu page
 * @param title     pointer to text for title in header (NULL to not display title)
 */
// llgo:link (*LvObjT).LvMenuSetPageTitle C.lv_menu_set_page_title
func (recv_ *LvObjT) LvMenuSetPageTitle(title *c.Char) {
}

/**
 * Set menu page title with a static text. It will not be saved by the label so the 'text' variable
 * has to be 'alive' while the page exists.
 * @param page      pointer to the menu page
 * @param title     pointer to text for title in header (NULL to not display title)
 */
// llgo:link (*LvObjT).LvMenuSetPageTitleStatic C.lv_menu_set_page_title_static
func (recv_ *LvObjT) LvMenuSetPageTitleStatic(title *c.Char) {
}

/**
 * Set menu page to display in sidebar
 * @param obj       pointer to the menu
 * @param page      pointer to the menu page to set (NULL to clear sidebar)
 */
// llgo:link (*LvObjT).LvMenuSetSidebarPage C.lv_menu_set_sidebar_page
func (recv_ *LvObjT) LvMenuSetSidebarPage(page *LvObjT) {
}

/**
 * Set the how the header should behave and its position
 * @param obj       pointer to a menu
 * @param mode      LV_MENU_HEADER_TOP_FIXED/TOP_UNFIXED/BOTTOM_FIXED
 */
// llgo:link (*LvObjT).LvMenuSetModeHeader C.lv_menu_set_mode_header
func (recv_ *LvObjT) LvMenuSetModeHeader(mode LvMenuModeHeaderT) {
}

/**
 * Set whether back button should appear at root
 * @param obj       pointer to a menu
 * @param mode      LV_MENU_ROOT_BACK_BUTTON_DISABLED/ENABLED
 */
// llgo:link (*LvObjT).LvMenuSetModeRootBackButton C.lv_menu_set_mode_root_back_button
func (recv_ *LvObjT) LvMenuSetModeRootBackButton(mode LvMenuModeRootBackButtonT) {
}

/**
 * Add menu to the menu item
 * @param menu      pointer to the menu
 * @param obj       pointer to the obj
 * @param page      pointer to the page to load when obj is clicked
 */
// llgo:link (*LvObjT).LvMenuSetLoadPageEvent C.lv_menu_set_load_page_event
func (recv_ *LvObjT) LvMenuSetLoadPageEvent(obj *LvObjT, page *LvObjT) {
}

/*=====================
 * Getter functions
 *====================*/
/**
* Get a pointer to menu page that is currently displayed in main
* @param obj        pointer to the menu
* @return           pointer to current page
 */
// llgo:link (*LvObjT).LvMenuGetCurMainPage C.lv_menu_get_cur_main_page
func (recv_ *LvObjT) LvMenuGetCurMainPage() *LvObjT {
	return nil
}

/**
* Get a pointer to menu page that is currently displayed in sidebar
* @param obj        pointer to the menu
* @return           pointer to current page
 */
// llgo:link (*LvObjT).LvMenuGetCurSidebarPage C.lv_menu_get_cur_sidebar_page
func (recv_ *LvObjT) LvMenuGetCurSidebarPage() *LvObjT {
	return nil
}

/**
* Get a pointer to main header obj
* @param obj        pointer to the menu
* @return           pointer to main header obj
 */
// llgo:link (*LvObjT).LvMenuGetMainHeader C.lv_menu_get_main_header
func (recv_ *LvObjT) LvMenuGetMainHeader() *LvObjT {
	return nil
}

/**
* Get a pointer to main header back btn obj
* @param obj        pointer to the menu
* @return           pointer to main header back btn obj
 */
// llgo:link (*LvObjT).LvMenuGetMainHeaderBackButton C.lv_menu_get_main_header_back_button
func (recv_ *LvObjT) LvMenuGetMainHeaderBackButton() *LvObjT {
	return nil
}

/**
* Get a pointer to sidebar header obj
* @param obj        pointer to the menu
* @return           pointer to sidebar header obj
 */
// llgo:link (*LvObjT).LvMenuGetSidebarHeader C.lv_menu_get_sidebar_header
func (recv_ *LvObjT) LvMenuGetSidebarHeader() *LvObjT {
	return nil
}

/**
* Get a pointer to sidebar header obj
* @param obj        pointer to the menu
* @return           pointer to sidebar header back btn obj
 */
// llgo:link (*LvObjT).LvMenuGetSidebarHeaderBackButton C.lv_menu_get_sidebar_header_back_button
func (recv_ *LvObjT) LvMenuGetSidebarHeaderBackButton() *LvObjT {
	return nil
}

/**
 * Check if an obj is a root back btn
 * @param menu      pointer to the menu
 * @param obj       pointer to the back button
 * @return          true if it is a root back btn
 */
// llgo:link (*LvObjT).LvMenuBackButtonIsRoot C.lv_menu_back_button_is_root
func (recv_ *LvObjT) LvMenuBackButtonIsRoot(obj *LvObjT) bool {
	return false
}

/**
 * Clear menu history
 * @param obj       pointer to the menu
 */
// llgo:link (*LvObjT).LvMenuClearHistory C.lv_menu_clear_history
func (recv_ *LvObjT) LvMenuClearHistory() {
}

type LvChartTypeT c.Int

const (
	LV_CHART_TYPE_NONE    LvChartTypeT = 0
	LV_CHART_TYPE_LINE    LvChartTypeT = 1
	LV_CHART_TYPE_BAR     LvChartTypeT = 2
	LV_CHART_TYPE_SCATTER LvChartTypeT = 3
)

type LvChartUpdateModeT c.Int

const (
	LV_CHART_UPDATE_MODE_SHIFT    LvChartUpdateModeT = 0
	LV_CHART_UPDATE_MODE_CIRCULAR LvChartUpdateModeT = 1
)

type LvChartAxisT c.Int

const (
	LV_CHART_AXIS_PRIMARY_Y   LvChartAxisT = 0
	LV_CHART_AXIS_SECONDARY_Y LvChartAxisT = 1
	LV_CHART_AXIS_PRIMARY_X   LvChartAxisT = 2
	LV_CHART_AXIS_SECONDARY_X LvChartAxisT = 4
	LV_CHART_AXIS_LAST        LvChartAxisT = 5
)

/**
 * Create a chart object
 * @param parent    pointer to an object, it will be the parent of the new chart
 * @return          pointer to the created chart
 */
// llgo:link (*LvObjT).LvChartCreate C.lv_chart_create
func (recv_ *LvObjT) LvChartCreate() *LvObjT {
	return nil
}

/**
 * Set a new type for a chart
 * @param obj       pointer to a chart object
 * @param type      new type of the chart (from 'lv_chart_type_t' enum)
 */
// llgo:link (*LvObjT).LvChartSetType C.lv_chart_set_type
func (recv_ *LvObjT) LvChartSetType(type_ LvChartTypeT) {
}

/**
 * Set the number of points on a data line on a chart
 * @param obj       pointer to a chart object
 * @param cnt       new number of points on the data lines
 */
// llgo:link (*LvObjT).LvChartSetPointCount C.lv_chart_set_point_count
func (recv_ *LvObjT) LvChartSetPointCount(cnt c.Uint32T) {
}

/**
 * Set the minimal and maximal y values on an axis
 * @param obj       pointer to a chart object
 * @param axis      `LV_CHART_AXIS_PRIMARY_Y` or `LV_CHART_AXIS_SECONDARY_Y`
 * @param min       minimum value of the y axis
 * @param max       maximum value of the y axis
 */
// llgo:link (*LvObjT).LvChartSetAxisRange C.lv_chart_set_axis_range
func (recv_ *LvObjT) LvChartSetAxisRange(axis LvChartAxisT, min c.Int32T, max c.Int32T) {
}

/**
 * Set the minimal values on an axis
 * @param obj       pointer to a chart object
 * @param axis      `LV_CHART_AXIS_PRIMARY_Y` or `LV_CHART_AXIS_SECONDARY_Y`
 * @param min       minimal value of the y axis
 */
// llgo:link (*LvObjT).LvChartSetAxisMinValue C.lv_chart_set_axis_min_value
func (recv_ *LvObjT) LvChartSetAxisMinValue(axis LvChartAxisT, min c.Int32T) {
}

/**
 * Set the maximal y values on an axis
 * @param obj       pointer to a chart object
 * @param axis      `LV_CHART_AXIS_PRIMARY_Y` or `LV_CHART_AXIS_SECONDARY_Y`
 * @param max       maximum value of the y axis
 */
// llgo:link (*LvObjT).LvChartSetAxisMaxValue C.lv_chart_set_axis_max_value
func (recv_ *LvObjT) LvChartSetAxisMaxValue(axis LvChartAxisT, max c.Int32T) {
}

/**
 * Set update mode of the chart object. Affects
 * @param obj              pointer to a chart object
 * @param update_mode      the update mode
 */
// llgo:link (*LvObjT).LvChartSetUpdateMode C.lv_chart_set_update_mode
func (recv_ *LvObjT) LvChartSetUpdateMode(update_mode LvChartUpdateModeT) {
}

/**
 * Set the number of horizontal and vertical division lines
 * @param obj       pointer to a chart object
 * @param hdiv      number of horizontal division lines
 * @param vdiv      number of vertical division lines
 */
// llgo:link (*LvObjT).LvChartSetDivLineCount C.lv_chart_set_div_line_count
func (recv_ *LvObjT) LvChartSetDivLineCount(hdiv c.Uint32T, vdiv c.Uint32T) {
}

/**
 * Set the number of horizontal division lines
 * @param obj       pointer to a chart object
 * @param cnt       number of horizontal division lines
 */
// llgo:link (*LvObjT).LvChartSetHorDivLineCount C.lv_chart_set_hor_div_line_count
func (recv_ *LvObjT) LvChartSetHorDivLineCount(cnt c.Uint32T) {
}

/**
 * Set the number of vertical division lines
 * @param obj       pointer to a chart object
 * @param cnt       number of vertical division lines
 */
// llgo:link (*LvObjT).LvChartSetVerDivLineCount C.lv_chart_set_ver_div_line_count
func (recv_ *LvObjT) LvChartSetVerDivLineCount(cnt c.Uint32T) {
}

/**
 * Get the type of a chart
 * @param obj       pointer to chart object
 * @return          type of the chart (from 'lv_chart_t' enum)
 */
// llgo:link (*LvObjT).LvChartGetType C.lv_chart_get_type
func (recv_ *LvObjT) LvChartGetType() LvChartTypeT {
	return 0
}

/**
 * Get the data point number per data line on chart
 * @param obj       pointer to chart object
 * @return          point number on each data line
 */
// llgo:link (*LvObjT).LvChartGetPointCount C.lv_chart_get_point_count
func (recv_ *LvObjT) LvChartGetPointCount() c.Uint32T {
	return 0
}

/**
 * Get the current index of the x-axis start point in the data array
 * @param obj       pointer to a chart object
 * @param ser       pointer to a data series on 'chart'
 * @return          the index of the current x start point in the data array
 */
// llgo:link (*LvObjT).LvChartGetXStartPoint C.lv_chart_get_x_start_point
func (recv_ *LvObjT) LvChartGetXStartPoint(ser *LvChartSeriesT) c.Uint32T {
	return 0
}

/**
 * Get the position of a point to the chart.
 * @param obj       pointer to a chart object
 * @param ser       pointer to series
 * @param id        the index.
 * @param p_out     store the result position here
 */
// llgo:link (*LvObjT).LvChartGetPointPosById C.lv_chart_get_point_pos_by_id
func (recv_ *LvObjT) LvChartGetPointPosById(ser *LvChartSeriesT, id c.Uint32T, p_out *LvPointT) {
}

/**
 * Refresh a chart if its data line has changed
 * @param   obj   pointer to chart object
 */
// llgo:link (*LvObjT).LvChartRefresh C.lv_chart_refresh
func (recv_ *LvObjT) LvChartRefresh() {
}

/**
 * Allocate and add a data series to the chart
 * @param obj       pointer to a chart object
 * @param color     color of the data series
 * @param axis      the y axis to which the series should be attached (::LV_CHART_AXIS_PRIMARY_Y or ::LV_CHART_AXIS_SECONDARY_Y)
 * @return          pointer to the allocated data series or NULL on failure
 */
// llgo:link (*LvObjT).LvChartAddSeries C.lv_chart_add_series
func (recv_ *LvObjT) LvChartAddSeries(color LvColorT, axis LvChartAxisT) *LvChartSeriesT {
	return nil
}

/**
 * Deallocate and remove a data series from a chart
 * @param obj       pointer to a chart object
 * @param series    pointer to a data series on 'chart'
 */
// llgo:link (*LvObjT).LvChartRemoveSeries C.lv_chart_remove_series
func (recv_ *LvObjT) LvChartRemoveSeries(series *LvChartSeriesT) {
}

/**
 * Hide/Unhide a single series of a chart.
 * @param chart     pointer to a chart object.
 * @param series    pointer to a series object
 * @param hide      true: hide the series
 */
// llgo:link (*LvObjT).LvChartHideSeries C.lv_chart_hide_series
func (recv_ *LvObjT) LvChartHideSeries(series *LvChartSeriesT, hide bool) {
}

/**
 * Change the color of a series
 * @param chart     pointer to a chart object.
 * @param series    pointer to a series object
 * @param color     the new color of the series
 */
// llgo:link (*LvObjT).LvChartSetSeriesColor C.lv_chart_set_series_color
func (recv_ *LvObjT) LvChartSetSeriesColor(series *LvChartSeriesT, color LvColorT) {
}

/**
 * Get the color of a series
 * @param chart     pointer to a chart object.
 * @param series    pointer to a series object
 * @return          the color of the series
 */
// llgo:link (*LvObjT).LvChartGetSeriesColor C.lv_chart_get_series_color
func (recv_ *LvObjT) LvChartGetSeriesColor(series *LvChartSeriesT) LvColorT {
	return LvColorT{}
}

/**
 * Set the index of the x-axis start point in the data array.
 * This point will be considers the first (left) point and the other points will be drawn after it.
 * @param obj       pointer to a chart object
 * @param ser       pointer to a data series on 'chart'
 * @param id        the index of the x point in the data array
 */
// llgo:link (*LvObjT).LvChartSetXStartPoint C.lv_chart_set_x_start_point
func (recv_ *LvObjT) LvChartSetXStartPoint(ser *LvChartSeriesT, id c.Uint32T) {
}

/**
 * Get the next series.
 * @param chart     pointer to a chart
 * @param ser      the previous series or NULL to get the first
 * @return          the next series or NULL if there is no more.
 */
// llgo:link (*LvObjT).LvChartGetSeriesNext C.lv_chart_get_series_next
func (recv_ *LvObjT) LvChartGetSeriesNext(ser *LvChartSeriesT) *LvChartSeriesT {
	return nil
}

/**
 * Add a cursor with a given color
 * @param obj       pointer to chart object
 * @param color     color of the cursor
 * @param dir       direction of the cursor. `LV_DIR_RIGHT/LEFT/TOP/DOWN/HOR/VER/ALL`. OR-ed values are possible
 * @return          pointer to the created cursor
 */
// llgo:link (*LvObjT).LvChartAddCursor C.lv_chart_add_cursor
func (recv_ *LvObjT) LvChartAddCursor(color LvColorT, dir LvDirT) *LvChartCursorT {
	return nil
}

/**
 * Set the coordinate of the cursor with respect to the paddings
 * @param chart     pointer to a chart object
 * @param cursor    pointer to the cursor
 * @param pos       the new coordinate of cursor relative to the chart
 */
// llgo:link (*LvObjT).LvChartSetCursorPos C.lv_chart_set_cursor_pos
func (recv_ *LvObjT) LvChartSetCursorPos(cursor *LvChartCursorT, pos *LvPointT) {
}

/**
 * Set the X coordinate of the cursor with respect to the paddings
 * @param chart     pointer to a chart object
 * @param cursor    pointer to the cursor
 * @param x         the new X coordinate of cursor relative to the chart
 */
// llgo:link (*LvObjT).LvChartSetCursorPosX C.lv_chart_set_cursor_pos_x
func (recv_ *LvObjT) LvChartSetCursorPosX(cursor *LvChartCursorT, x c.Int32T) {
}

/**
 * Set the coordinate of the cursor with respect to the paddings
 * @param chart     pointer to a chart object
 * @param cursor    pointer to the cursor
 * @param y         the new Y coordinate of cursor relative to the chart
 */
// llgo:link (*LvObjT).LvChartSetCursorPosY C.lv_chart_set_cursor_pos_y
func (recv_ *LvObjT) LvChartSetCursorPosY(cursor *LvChartCursorT, y c.Int32T) {
}

/**
 * Stick the cursor to a point
 * @param chart     pointer to a chart object
 * @param cursor    pointer to the cursor
 * @param ser       pointer to a series
 * @param point_id  the point's index or `LV_CHART_POINT_NONE` to not assign to any points.
 */
// llgo:link (*LvObjT).LvChartSetCursorPoint C.lv_chart_set_cursor_point
func (recv_ *LvObjT) LvChartSetCursorPoint(cursor *LvChartCursorT, ser *LvChartSeriesT, point_id c.Uint32T) {
}

/**
 * Get the coordinate of the cursor with respect to the paddings
 * @param chart     pointer to a chart object
 * @param cursor    pointer to cursor
 * @return          coordinate of the cursor as lv_point_t
 */
// llgo:link (*LvObjT).LvChartGetCursorPoint C.lv_chart_get_cursor_point
func (recv_ *LvObjT) LvChartGetCursorPoint(cursor *LvChartCursorT) LvPointT {
	return LvPointT{}
}

/**
 * Initialize all data points of a series with a value
 * @param obj       pointer to chart object
 * @param ser       pointer to a data series on 'chart'
 * @param value     the new value for all points. `LV_CHART_POINT_NONE` can be used to hide the points.
 */
// llgo:link (*LvObjT).LvChartSetAllValues C.lv_chart_set_all_values
func (recv_ *LvObjT) LvChartSetAllValues(ser *LvChartSeriesT, value c.Int32T) {
}

/**
 * Set the next point's Y value according to the update mode policy.
 * @param obj       pointer to chart object
 * @param ser       pointer to a data series on 'chart'
 * @param value     the new value of the next data
 */
// llgo:link (*LvObjT).LvChartSetNextValue C.lv_chart_set_next_value
func (recv_ *LvObjT) LvChartSetNextValue(ser *LvChartSeriesT, value c.Int32T) {
}

/**
 * Set the next point's X and Y value according to the update mode policy.
 * @param obj       pointer to chart object
 * @param ser       pointer to a data series on 'chart'
 * @param x_value   the new X value of the next data
 * @param y_value   the new Y value of the next data
 */
// llgo:link (*LvObjT).LvChartSetNextValue2 C.lv_chart_set_next_value2
func (recv_ *LvObjT) LvChartSetNextValue2(ser *LvChartSeriesT, x_value c.Int32T, y_value c.Int32T) {
}

/**
 * Same as `lv_chart_set_next_value` but set the values from an array
 * @param obj           pointer to chart object
 * @param ser           pointer to a data series on 'chart'
 * @param values        the new values to set
 * @param values_cnt    number of items in `values`
 */
// llgo:link (*LvObjT).LvChartSetSeriesValues C.lv_chart_set_series_values
func (recv_ *LvObjT) LvChartSetSeriesValues(ser *LvChartSeriesT, values *c.Int32T, values_cnt c.SizeT) {
}

/**
 * Same as `lv_chart_set_next_value2` but set the values from an array
 * @param obj           pointer to chart object
 * @param ser           pointer to a data series on 'chart'
 * @param x_values      the new values to set on the X axis
 * @param y_values      the new values to set o nthe Y axis
 * @param values_cnt    number of items in `x_values` and `y_values`
 */
// llgo:link (*LvObjT).LvChartSetSeriesValues2 C.lv_chart_set_series_values2
func (recv_ *LvObjT) LvChartSetSeriesValues2(ser *LvChartSeriesT, x_values *c.Int32T, y_values *c.Int32T, values_cnt c.SizeT) {
}

/**
 * Set an individual point's y value of a chart's series directly based on its index
 * @param obj     pointer to a chart object
 * @param ser     pointer to a data series on 'chart'
 * @param id      the index of the x point in the array
 * @param value   value to assign to array point
 */
// llgo:link (*LvObjT).LvChartSetSeriesValueById C.lv_chart_set_series_value_by_id
func (recv_ *LvObjT) LvChartSetSeriesValueById(ser *LvChartSeriesT, id c.Uint32T, value c.Int32T) {
}

/**
 * Set an individual point's x and y value of a chart's series directly based on its index
 * Can be used only with `LV_CHART_TYPE_SCATTER`.
 * @param obj       pointer to chart object
 * @param ser       pointer to a data series on 'chart'
 * @param id        the index of the x point in the array
 * @param x_value   the new X value of the next data
 * @param y_value   the new Y value of the next data
 */
// llgo:link (*LvObjT).LvChartSetSeriesValueById2 C.lv_chart_set_series_value_by_id2
func (recv_ *LvObjT) LvChartSetSeriesValueById2(ser *LvChartSeriesT, id c.Uint32T, x_value c.Int32T, y_value c.Int32T) {
}

/**
 * Set an external array for the y data points to use for the chart
 * NOTE: It is the users responsibility to make sure the `point_cnt` matches the external array size.
 * @param obj       pointer to a chart object
 * @param ser       pointer to a data series on 'chart'
 * @param array     external array of points for chart
 */
// llgo:link (*LvObjT).LvChartSetSeriesExtYArray C.lv_chart_set_series_ext_y_array
func (recv_ *LvObjT) LvChartSetSeriesExtYArray(ser *LvChartSeriesT, array *c.Int32T) {
}

/**
 * Set an external array for the x data points to use for the chart
 * NOTE: It is the users responsibility to make sure the `point_cnt` matches the external array size.
 * @param obj       pointer to a chart object
 * @param ser       pointer to a data series on 'chart'
 * @param array     external array of points for chart
 */
// llgo:link (*LvObjT).LvChartSetSeriesExtXArray C.lv_chart_set_series_ext_x_array
func (recv_ *LvObjT) LvChartSetSeriesExtXArray(ser *LvChartSeriesT, array *c.Int32T) {
}

/**
 * Get the array of y values of a series
 * @param obj   pointer to a chart object
 * @param ser   pointer to a data series on 'chart'
 * @return      the array of values with 'point_count' elements
 */
// llgo:link (*LvObjT).LvChartGetSeriesYArray C.lv_chart_get_series_y_array
func (recv_ *LvObjT) LvChartGetSeriesYArray(ser *LvChartSeriesT) *c.Int32T {
	return nil
}

/**
 * Get the array of x values of a series
 * @param obj   pointer to a chart object
 * @param ser   pointer to a data series on 'chart'
 * @return      the array of values with 'point_count' elements
 */
// llgo:link (*LvObjT).LvChartGetSeriesXArray C.lv_chart_get_series_x_array
func (recv_ *LvObjT) LvChartGetSeriesXArray(ser *LvChartSeriesT) *c.Int32T {
	return nil
}

/**
 * Get the index of the currently pressed point. It's the same for every series.
 * @param obj       pointer to a chart object
 * @return          the index of the point [0 .. point count] or LV_CHART_POINT_ID_NONE if no point is being pressed
 */
// llgo:link (*LvObjT).LvChartGetPressedPoint C.lv_chart_get_pressed_point
func (recv_ *LvObjT) LvChartGetPressedPoint() c.Uint32T {
	return 0
}

/**
 * Get the overall offset from the chart's side to the center of the first point.
 * In case of a bar chart it will be the center of the first column group
 * @param obj       pointer to a chart object
 * @return          the offset of the center
 */
// llgo:link (*LvObjT).LvChartGetFirstPointCenterOffset C.lv_chart_get_first_point_center_offset
func (recv_ *LvObjT) LvChartGetFirstPointCenterOffset() c.Int32T {
	return 0
}

/**
 * Create a button object
 * @param parent    pointer to an object, it will be the parent of the new button
 * @return          pointer to the created button
 */
// llgo:link (*LvObjT).LvButtonCreate C.lv_button_create
func (recv_ *LvObjT) LvButtonCreate() *LvObjT {
	return nil
}

type LvScaleModeT c.Int

const (
	LV_SCALE_MODE_HORIZONTAL_TOP    LvScaleModeT = 0
	LV_SCALE_MODE_HORIZONTAL_BOTTOM LvScaleModeT = 1
	LV_SCALE_MODE_VERTICAL_LEFT     LvScaleModeT = 2
	LV_SCALE_MODE_VERTICAL_RIGHT    LvScaleModeT = 4
	LV_SCALE_MODE_ROUND_INNER       LvScaleModeT = 8
	LV_SCALE_MODE_ROUND_OUTER       LvScaleModeT = 16
	LV_SCALE_MODE_LAST              LvScaleModeT = 17
)

/**
 * Create an scale object
 * @param parent    pointer to an object, it will be the parent of the new scale
 * @return          pointer to created Scale Widget
 */
// llgo:link (*LvObjT).LvScaleCreate C.lv_scale_create
func (recv_ *LvObjT) LvScaleCreate() *LvObjT {
	return nil
}

/**
 * Set scale mode. See lv_scale_mode_t.
 * @param obj       pointer to Scale Widget
 * @param mode      the new scale mode
 */
// llgo:link (*LvObjT).LvScaleSetMode C.lv_scale_set_mode
func (recv_ *LvObjT) LvScaleSetMode(mode LvScaleModeT) {
}

/**
 * Set scale total tick count (including minor and major ticks).
 * @param obj       pointer to Scale Widget
 * @param total_tick_count    New total tick count
 */
// llgo:link (*LvObjT).LvScaleSetTotalTickCount C.lv_scale_set_total_tick_count
func (recv_ *LvObjT) LvScaleSetTotalTickCount(total_tick_count c.Uint32T) {
}

/**
 * Sets how often major ticks are drawn.
 * @param obj                 pointer to Scale Widget
 * @param major_tick_every    the new count for major tick drawing
 */
// llgo:link (*LvObjT).LvScaleSetMajorTickEvery C.lv_scale_set_major_tick_every
func (recv_ *LvObjT) LvScaleSetMajorTickEvery(major_tick_every c.Uint32T) {
}

/**
 * Sets label visibility.
 * @param obj           pointer to Scale Widget
 * @param show_label    true/false to enable tick label
 */
// llgo:link (*LvObjT).LvScaleSetLabelShow C.lv_scale_set_label_show
func (recv_ *LvObjT) LvScaleSetLabelShow(show_label bool) {
}

/**
 * Set minimum and maximum values on Scale.
 * @param obj       pointer to Scale Widget
 * @param min       minimum value of Scale
 * @param max       maximum value of Scale
 */
// llgo:link (*LvObjT).LvScaleSetRange C.lv_scale_set_range
func (recv_ *LvObjT) LvScaleSetRange(min c.Int32T, max c.Int32T) {
}

/**
 * Set minimum values on Scale.
 * @param obj       pointer to Scale Widget
 * @param min       minimum value of Scale
 */
// llgo:link (*LvObjT).LvScaleSetMinValue C.lv_scale_set_min_value
func (recv_ *LvObjT) LvScaleSetMinValue(min c.Int32T) {
}

/**
 * Set maximum values on Scale.
 * @param obj       pointer to Scale Widget
 * @param min       minimum value of Scale
 */
// llgo:link (*LvObjT).LvScaleSetMaxValue C.lv_scale_set_max_value
func (recv_ *LvObjT) LvScaleSetMaxValue(max c.Int32T) {
}

/**
 * Set angle between the low end and the high end of the Scale.
 * (Applies only to round Scales.)
 * @param obj         pointer to Scale Widget
 * @param max_angle   angle in degrees from Scale minimum where top end of Scale will be drawn
 */
// llgo:link (*LvObjT).LvScaleSetAngleRange C.lv_scale_set_angle_range
func (recv_ *LvObjT) LvScaleSetAngleRange(angle_range c.Uint32T) {
}

/**
 * Set angular offset from the 3-o'clock position of the low end of the Scale.
 * (Applies only to round Scales.)
 * @param obj       pointer to Scale Widget
 * @param rotation  clockwise angular offset (in degrees) from the 3-o'clock position
 *                  of the low end of the scale; negative and >360 values are first normalized
 *                  to range [0..360].
 *                  Examples:
 *                      -   0 = 3 o'clock (right side)
 *                      -  30 = 4 o'clock
 *                      -  60 = 5 o'clock
 *                      -  90 = 6 o'clock
 *                      - 135 = midway between 7 and 8 o'clock (default)
 *                      - 180 = 9 o'clock
 *                      - 270 = 12 o'clock
 *                      - 300 = 1 o'clock
 *                      - 330 = 2 o'clock
 *                      - -30 = 2 o'clock
 *                      - 390 = 4 o'clock
 */
// llgo:link (*LvObjT).LvScaleSetRotation C.lv_scale_set_rotation
func (recv_ *LvObjT) LvScaleSetRotation(rotation c.Int32T) {
}

/**
 * Point line needle to specified value.
 * @param obj              pointer to Scale Widget
 * @param needle_line      needle_line of the Scale. The line points will be allocated and
 *                         managed by the Scale unless the line point array was previously set
 *                         using `lv_line_set_points_mutable`.
 * @param needle_length    length of the needle
 *                         - needle_length>0: needle_length=needle_length;
 *                         - needle_length<0: needle_length=radius-|needle_length|;
 * @param value            Scale value needle will point to
 */
// llgo:link (*LvObjT).LvScaleSetLineNeedleValue C.lv_scale_set_line_needle_value
func (recv_ *LvObjT) LvScaleSetLineNeedleValue(needle_line *LvObjT, needle_length c.Int32T, value c.Int32T) {
}

/**
* Point image needle to specified value;
  image must point to the right. E.g. -O------>
* @param obj              pointer to Scale Widget
* @param needle_img       pointer to needle's Image
* @param value            Scale value needle will point to
*/
// llgo:link (*LvObjT).LvScaleSetImageNeedleValue C.lv_scale_set_image_needle_value
func (recv_ *LvObjT) LvScaleSetImageNeedleValue(needle_img *LvObjT, value c.Int32T) {
}

/**
 * Set custom text source for major ticks labels.
 * @param obj       pointer to Scale Widget
 * @param txt_src   pointer to an array of strings which will be display at major ticks;
 *                  last element must be a NULL pointer.
 */
// llgo:link (*LvObjT).LvScaleSetTextSrc C.lv_scale_set_text_src
func (recv_ *LvObjT) LvScaleSetTextSrc(txt_src **c.Char) {
}

/**
 * Draw Scale after all its children are drawn.
 * @param obj       pointer to Scale Widget
 * @param en        true: enable post draw
 */
// llgo:link (*LvObjT).LvScaleSetPostDraw C.lv_scale_set_post_draw
func (recv_ *LvObjT) LvScaleSetPostDraw(en bool) {
}

/**
 * Draw Scale ticks on top of all other parts.
 * @param obj       pointer to Scale Widget
 * @param en        true: enable draw ticks on top of all parts
 */
// llgo:link (*LvObjT).LvScaleSetDrawTicksOnTop C.lv_scale_set_draw_ticks_on_top
func (recv_ *LvObjT) LvScaleSetDrawTicksOnTop(en bool) {
}

/**
 * Add a Section to specified Scale.  Section will not be drawn until
 * a valid range is set for it using `lv_scale_set_section_range()`.
 * @param obj       pointer to Scale Widget
 * @return          pointer to new Section
 */
// llgo:link (*LvObjT).LvScaleAddSection C.lv_scale_add_section
func (recv_ *LvObjT) LvScaleAddSection() *LvScaleSectionT {
	return nil
}

/**
 * DEPRECATED, use lv_scale_set_section_range instead.
 * Set range for specified Scale Section
 * @param section       pointer to Section
 * @param range_min     Section new minimum value
 * @param range_max     Section new maximum value
 */
// llgo:link (*LvScaleSectionT).LvScaleSectionSetRange C.lv_scale_section_set_range
func (recv_ *LvScaleSectionT) LvScaleSectionSetRange(min c.Int32T, max c.Int32T) {
}

/**
 * Set the range of a scale section
 * @param scale         pointer to scale
 * @param section       pointer to section
 * @param range_min     the section's new minimum value
 * @param range_max     the section's new maximum value
 */
// llgo:link (*LvObjT).LvScaleSetSectionRange C.lv_scale_set_section_range
func (recv_ *LvObjT) LvScaleSetSectionRange(section *LvScaleSectionT, min c.Int32T, max c.Int32T) {
}

/**
 * Set the minimum value of a scale section
 * @param scale         pointer to scale
 * @param section       pointer to section
 * @param min           the section's new minimum value
 */
// llgo:link (*LvObjT).LvScaleSetSectionMinValue C.lv_scale_set_section_min_value
func (recv_ *LvObjT) LvScaleSetSectionMinValue(section *LvScaleSectionT, min c.Int32T) {
}

/**
 * Set the maximum value of a scale section
 * @param scale         pointer to scale
 * @param section       pointer to section
 * @param max           the section's new maximum value
 */
// llgo:link (*LvObjT).LvScaleSetSectionMaxValue C.lv_scale_set_section_max_value
func (recv_ *LvObjT) LvScaleSetSectionMaxValue(section *LvScaleSectionT, max c.Int32T) {
}

/**
 * DEPRECATED, use lv_scale_set_section_style_main/indicator/items instead.
 * Set style for specified part of Section.
 * @param section             pointer to Section
 * @param part                the part of the Scale the style will apply to, e.g. LV_PART_INDICATOR
 * @param section_part_style  pointer to style to apply
 */
// llgo:link (*LvScaleSectionT).LvScaleSectionSetStyle C.lv_scale_section_set_style
func (recv_ *LvScaleSectionT) LvScaleSectionSetStyle(part LvPartT, section_part_style *LvStyleT) {
}

/**
 * Set the style of the line on a section.
 * @param scale         pointer to scale
 * @param section       pointer to section
 * @param style         point to a style
 */
// llgo:link (*LvObjT).LvScaleSetSectionStyleMain C.lv_scale_set_section_style_main
func (recv_ *LvObjT) LvScaleSetSectionStyleMain(section *LvScaleSectionT, style *LvStyleT) {
}

/**
 * Set the style of the major ticks and label on a section.
 * @param scale         pointer to scale
 * @param section       pointer to section
 * @param style         point to a style
 */
// llgo:link (*LvObjT).LvScaleSetSectionStyleIndicator C.lv_scale_set_section_style_indicator
func (recv_ *LvObjT) LvScaleSetSectionStyleIndicator(section *LvScaleSectionT, style *LvStyleT) {
}

/**
 * Set the style of the minor ticks on a section.
 * @param scale         pointer to scale
 * @param section       pointer to section
 * @param style         point to a style
 */
// llgo:link (*LvObjT).LvScaleSetSectionStyleItems C.lv_scale_set_section_style_items
func (recv_ *LvObjT) LvScaleSetSectionStyleItems(section *LvScaleSectionT, style *LvStyleT) {
}

/**
 * Get scale mode. See lv_scale_mode_t
 * @param obj   pointer to Scale Widget
 * @return      Scale mode
 */
// llgo:link (*LvObjT).LvScaleGetMode C.lv_scale_get_mode
func (recv_ *LvObjT) LvScaleGetMode() LvScaleModeT {
	return 0
}

/**
 * Get scale total tick count (including minor and major ticks)
 * @param obj   pointer to Scale Widget
 * @return      Scale total tick count
 */
// llgo:link (*LvObjT).LvScaleGetTotalTickCount C.lv_scale_get_total_tick_count
func (recv_ *LvObjT) LvScaleGetTotalTickCount() c.Int32T {
	return 0
}

/**
 * Get how often the major tick will be drawn
 * @param obj   pointer to Scale Widget
 * @return      Scale major tick every count
 */
// llgo:link (*LvObjT).LvScaleGetMajorTickEvery C.lv_scale_get_major_tick_every
func (recv_ *LvObjT) LvScaleGetMajorTickEvery() c.Int32T {
	return 0
}

/**
 * Get angular location of low end of Scale.
 * @param obj   pointer to Scale Widget
 * @return      Scale low end anglular location
 */
// llgo:link (*LvObjT).LvScaleGetRotation C.lv_scale_get_rotation
func (recv_ *LvObjT) LvScaleGetRotation() c.Int32T {
	return 0
}

/**
 * Gets label visibility
 * @param obj   pointer to Scale Widget
 * @return      true if tick label is enabled, false otherwise
 */
// llgo:link (*LvObjT).LvScaleGetLabelShow C.lv_scale_get_label_show
func (recv_ *LvObjT) LvScaleGetLabelShow() bool {
	return false
}

/**
 * Get Scale's range in degrees
 * @param obj   pointer to Scale Widget
 * @return      Scale's angle_range
 */
// llgo:link (*LvObjT).LvScaleGetAngleRange C.lv_scale_get_angle_range
func (recv_ *LvObjT) LvScaleGetAngleRange() c.Uint32T {
	return 0
}

/**
 * Get minimum value for Scale
 * @param obj   pointer to Scale Widget
 * @return      Scale's minimum value
 */
// llgo:link (*LvObjT).LvScaleGetRangeMinValue C.lv_scale_get_range_min_value
func (recv_ *LvObjT) LvScaleGetRangeMinValue() c.Int32T {
	return 0
}

/**
 * Get maximum value for Scale
 * @param obj   pointer to Scale Widget
 * @return      Scale's maximum value
 */
// llgo:link (*LvObjT).LvScaleGetRangeMaxValue C.lv_scale_get_range_max_value
func (recv_ *LvObjT) LvScaleGetRangeMaxValue() c.Int32T {
	return 0
}

/**
 * Create a led object
 * @param parent    pointer to an object, it will be the parent of the new led
 * @return          pointer to the created led
 */
// llgo:link (*LvObjT).LvLedCreate C.lv_led_create
func (recv_ *LvObjT) LvLedCreate() *LvObjT {
	return nil
}

/**
 * Set the color of the LED
 * @param led       pointer to a LED object
 * @param color     the color of the LED
 */
// llgo:link (*LvObjT).LvLedSetColor C.lv_led_set_color
func (recv_ *LvObjT) LvLedSetColor(color LvColorT) {
}

/**
 * Set the brightness of a LED object
 * @param led       pointer to a LED object
 * @param bright    LV_LED_BRIGHT_MIN (max. dark) ... LV_LED_BRIGHT_MAX (max. light)
 */
// llgo:link (*LvObjT).LvLedSetBrightness C.lv_led_set_brightness
func (recv_ *LvObjT) LvLedSetBrightness(bright c.Uint8T) {
}

/**
 * Light on a LED
 * @param led       pointer to a LED object
 */
// llgo:link (*LvObjT).LvLedOn C.lv_led_on
func (recv_ *LvObjT) LvLedOn() {
}

/**
 * Light off a LED
 * @param led       pointer to a LED object
 */
// llgo:link (*LvObjT).LvLedOff C.lv_led_off
func (recv_ *LvObjT) LvLedOff() {
}

/**
 * Toggle the state of a LED
 * @param led       pointer to a LED object
 */
// llgo:link (*LvObjT).LvLedToggle C.lv_led_toggle
func (recv_ *LvObjT) LvLedToggle() {
}

/**
 * Get the brightness of a LED object
 * @param obj       pointer to LED object
 * @return bright   0 (max. dark) ... 255 (max. light)
 */
// llgo:link (*LvObjT).LvLedGetBrightness C.lv_led_get_brightness
func (recv_ *LvObjT) LvLedGetBrightness() c.Uint8T {
	return 0
}

type LvArcModeT c.Int

const (
	LV_ARC_MODE_NORMAL      LvArcModeT = 0
	LV_ARC_MODE_SYMMETRICAL LvArcModeT = 1
	LV_ARC_MODE_REVERSE     LvArcModeT = 2
)

/**
 * Create an arc object
 * @param parent    pointer to an object, it will be the parent of the new arc
 * @return          pointer to the created arc
 */
// llgo:link (*LvObjT).LvArcCreate C.lv_arc_create
func (recv_ *LvObjT) LvArcCreate() *LvObjT {
	return nil
}

/**
 * Set the start angle of an arc. 0 deg: right, 90 bottom, etc.
 * @param obj       pointer to an arc object
 * @param start     the start angle. (if `LV_USE_FLOAT` is enabled it can be fractional too.)
 */
// llgo:link (*LvObjT).LvArcSetStartAngle C.lv_arc_set_start_angle
func (recv_ *LvObjT) LvArcSetStartAngle(start LvValuePreciseT) {
}

/**
 * Set the end angle of an arc. 0 deg: right, 90 bottom, etc.
 * @param obj       pointer to an arc object
 * @param end       the end angle  (if `LV_USE_FLOAT` is enabled it can be fractional too.)
 */
// llgo:link (*LvObjT).LvArcSetEndAngle C.lv_arc_set_end_angle
func (recv_ *LvObjT) LvArcSetEndAngle(end LvValuePreciseT) {
}

/**
 * Set the start and end angles
 * @param obj       pointer to an arc object
 * @param start     the start angle  (if `LV_USE_FLOAT` is enabled it can be fractional too.)
 * @param end       the end angle  (if `LV_USE_FLOAT` is enabled it can be fractional too.)
 */
// llgo:link (*LvObjT).LvArcSetAngles C.lv_arc_set_angles
func (recv_ *LvObjT) LvArcSetAngles(start LvValuePreciseT, end LvValuePreciseT) {
}

/**
 * Set the start angle of an arc background. 0 deg: right, 90 bottom, etc.
 * @param obj       pointer to an arc object
 * @param start     the start angle  (if `LV_USE_FLOAT` is enabled it can be fractional too.)
 */
// llgo:link (*LvObjT).LvArcSetBgStartAngle C.lv_arc_set_bg_start_angle
func (recv_ *LvObjT) LvArcSetBgStartAngle(start LvValuePreciseT) {
}

/**
 * Set the start angle of an arc background. 0 deg: right, 90 bottom etc.
 * @param obj       pointer to an arc object
 * @param end       the end angle  (if `LV_USE_FLOAT` is enabled it can be fractional too.)
 */
// llgo:link (*LvObjT).LvArcSetBgEndAngle C.lv_arc_set_bg_end_angle
func (recv_ *LvObjT) LvArcSetBgEndAngle(end LvValuePreciseT) {
}

/**
 * Set the start and end angles of the arc background
 * @param obj       pointer to an arc object
 * @param start     the start angle  (if `LV_USE_FLOAT` is enabled it can be fractional too.)
 * @param end       the end angle  (if `LV_USE_FLOAT` is enabled it can be fractional too.)
 */
// llgo:link (*LvObjT).LvArcSetBgAngles C.lv_arc_set_bg_angles
func (recv_ *LvObjT) LvArcSetBgAngles(start LvValuePreciseT, end LvValuePreciseT) {
}

/**
 * Set the rotation for the whole arc
 * @param obj           pointer to an arc object
 * @param rotation      rotation angle
 */
// llgo:link (*LvObjT).LvArcSetRotation C.lv_arc_set_rotation
func (recv_ *LvObjT) LvArcSetRotation(rotation c.Int32T) {
}

/**
 * Set the type of arc.
 * @param obj       pointer to arc object
 * @param type      arc's mode
 */
// llgo:link (*LvObjT).LvArcSetMode C.lv_arc_set_mode
func (recv_ *LvObjT) LvArcSetMode(type_ LvArcModeT) {
}

/**
 * Set a new value on the arc
 * @param obj       pointer to an arc object
 * @param value     new value
 */
// llgo:link (*LvObjT).LvArcSetValue C.lv_arc_set_value
func (recv_ *LvObjT) LvArcSetValue(value c.Int32T) {
}

/**
 * Set minimum and the maximum values of an arc
 * @param obj       pointer to the arc object
 * @param min       minimum value
 * @param max       maximum value
 */
// llgo:link (*LvObjT).LvArcSetRange C.lv_arc_set_range
func (recv_ *LvObjT) LvArcSetRange(min c.Int32T, max c.Int32T) {
}

/**
 * Set the minimum values of an arc
 * @param obj       pointer to the arc object
 * @param min       minimum value
 */
// llgo:link (*LvObjT).LvArcSetMinValue C.lv_arc_set_min_value
func (recv_ *LvObjT) LvArcSetMinValue(min c.Int32T) {
}

/**
 * Set the maximum values of an arc
 * @param obj       pointer to the arc object
 * @param max       maximum value
 */
// llgo:link (*LvObjT).LvArcSetMaxValue C.lv_arc_set_max_value
func (recv_ *LvObjT) LvArcSetMaxValue(max c.Int32T) {
}

/**
 * Set a change rate to limit the speed how fast the arc should reach the pressed point.
 * @param obj       pointer to an arc object
 * @param rate      the change rate
 */
// llgo:link (*LvObjT).LvArcSetChangeRate C.lv_arc_set_change_rate
func (recv_ *LvObjT) LvArcSetChangeRate(rate c.Uint32T) {
}

/**
 * Set an offset angle for the knob
 * @param obj       pointer to an arc object
 * @param offset    knob offset from main arc in degrees
 */
// llgo:link (*LvObjT).LvArcSetKnobOffset C.lv_arc_set_knob_offset
func (recv_ *LvObjT) LvArcSetKnobOffset(offset c.Int32T) {
}

/**
 * Get the start angle of an arc.
 * @param obj       pointer to an arc object
 * @return          the start angle [0..360]  (if `LV_USE_FLOAT` is enabled it can be fractional too.)
 */
// llgo:link (*LvObjT).LvArcGetAngleStart C.lv_arc_get_angle_start
func (recv_ *LvObjT) LvArcGetAngleStart() LvValuePreciseT {
	return 0
}

/**
 * Get the end angle of an arc.
 * @param obj       pointer to an arc object
 * @return          the end angle [0..360]  (if `LV_USE_FLOAT` is enabled it can be fractional too.)
 */
// llgo:link (*LvObjT).LvArcGetAngleEnd C.lv_arc_get_angle_end
func (recv_ *LvObjT) LvArcGetAngleEnd() LvValuePreciseT {
	return 0
}

/**
 * Get the start angle of an arc background.
 * @param obj       pointer to an arc object
 * @return          the  start angle [0..360]  (if `LV_USE_FLOAT` is enabled it can be fractional too.)
 */
// llgo:link (*LvObjT).LvArcGetBgAngleStart C.lv_arc_get_bg_angle_start
func (recv_ *LvObjT) LvArcGetBgAngleStart() LvValuePreciseT {
	return 0
}

/**
 * Get the end angle of an arc background.
 * @param obj       pointer to an arc object
 * @return          the end angle [0..360]  (if `LV_USE_FLOAT` is enabled it can be fractional too.)
 */
// llgo:link (*LvObjT).LvArcGetBgAngleEnd C.lv_arc_get_bg_angle_end
func (recv_ *LvObjT) LvArcGetBgAngleEnd() LvValuePreciseT {
	return 0
}

/**
 * Get the value of an arc
 * @param obj       pointer to an arc object
 * @return          the value of the arc
 */
// llgo:link (*LvObjT).LvArcGetValue C.lv_arc_get_value
func (recv_ *LvObjT) LvArcGetValue() c.Int32T {
	return 0
}

/**
 * Get the minimum value of an arc
 * @param obj       pointer to an arc object
 * @return          the minimum value of the arc
 */
// llgo:link (*LvObjT).LvArcGetMinValue C.lv_arc_get_min_value
func (recv_ *LvObjT) LvArcGetMinValue() c.Int32T {
	return 0
}

/**
 * Get the maximum value of an arc
 * @param obj       pointer to an arc object
 * @return          the maximum value of the arc
 */
// llgo:link (*LvObjT).LvArcGetMaxValue C.lv_arc_get_max_value
func (recv_ *LvObjT) LvArcGetMaxValue() c.Int32T {
	return 0
}

/**
 * Get whether the arc is type or not.
 * @param obj       pointer to an arc object
 * @return          arc's mode
 */
// llgo:link (*LvObjT).LvArcGetMode C.lv_arc_get_mode
func (recv_ *LvObjT) LvArcGetMode() LvArcModeT {
	return 0
}

/**
 * Get the rotation for the whole arc
 * @param obj       pointer to an arc object
 * @return          arc's current rotation
 */
// llgo:link (*LvObjT).LvArcGetRotation C.lv_arc_get_rotation
func (recv_ *LvObjT) LvArcGetRotation() c.Int32T {
	return 0
}

/**
 * Get the current knob angle offset
 * @param obj       pointer to an arc object
 * @return          arc's current knob offset
 */
// llgo:link (*LvObjT).LvArcGetKnobOffset C.lv_arc_get_knob_offset
func (recv_ *LvObjT) LvArcGetKnobOffset() c.Int32T {
	return 0
}

/**
 * Align an object to the current position of the arc (knob)
 * @param obj           pointer to an arc object
 * @param obj_to_align  pointer to an object to align
 * @param r_offset      consider the radius larger with this value (< 0: for smaller radius)
 */
// llgo:link (*LvObjT).LvArcAlignObjToAngle C.lv_arc_align_obj_to_angle
func (recv_ *LvObjT) LvArcAlignObjToAngle(obj_to_align *LvObjT, r_offset c.Int32T) {
}

/**
 * Rotate an object to the current position of the arc (knob)
 * @param obj            pointer to an arc object
 * @param obj_to_rotate  pointer to an object to rotate
 * @param r_offset       consider the radius larger with this value (< 0: for smaller radius)
 */
// llgo:link (*LvObjT).LvArcRotateObjToAngle C.lv_arc_rotate_obj_to_angle
func (recv_ *LvObjT) LvArcRotateObjToAngle(obj_to_rotate *LvObjT, r_offset c.Int32T) {
}

/**
 * Create a tileview object
 * @param parent      pointer to an object, it will be the parent of the new tileview
 * @return            pointer to the created tileview
 */
// llgo:link (*LvObjT).LvTileviewCreate C.lv_tileview_create
func (recv_ *LvObjT) LvTileviewCreate() *LvObjT {
	return nil
}

/**
 * Add a tile to the tileview
 * @param tv          pointer to the tileview object
 * @param col_id      column id of the tile
 * @param row_id      row id of the tile
 * @param dir         direction to move to the next tile
 * @return            pointer to the added tile object
 */
// llgo:link (*LvObjT).LvTileviewAddTile C.lv_tileview_add_tile
func (recv_ *LvObjT) LvTileviewAddTile(col_id c.Uint8T, row_id c.Uint8T, dir LvDirT) *LvObjT {
	return nil
}

/**
 * Set the active tile in the tileview.
 * @param parent      pointer to the tileview object
 * @param tile_obj    pointer to the tile object to be set as active
 * @param anim_en     animation enable flag (LV_ANIM_ON or LV_ANIM_OFF)
 */
// llgo:link (*LvObjT).LvTileviewSetTile C.lv_tileview_set_tile
func (recv_ *LvObjT) LvTileviewSetTile(tile_obj *LvObjT, anim_en LvAnimEnableT) {
}

/**
 * Set the active tile by index in the tileview
 * @param tv          pointer to the tileview object
 * @param col_id      column id of the tile to be set as active
 * @param row_id      row id of the tile to be set as active
 * @param anim_en     animation enable flag (LV_ANIM_ON or LV_ANIM_OFF)
 */
// llgo:link (*LvObjT).LvTileviewSetTileByIndex C.lv_tileview_set_tile_by_index
func (recv_ *LvObjT) LvTileviewSetTileByIndex(col_id c.Uint32T, row_id c.Uint32T, anim_en LvAnimEnableT) {
}

/**
 * Get the currently active tile in the tileview
 * @param obj         pointer to the tileview object
 * @return            pointer to the currently active tile object
 */
// llgo:link (*LvObjT).LvTileviewGetTileActive C.lv_tileview_get_tile_active
func (recv_ *LvObjT) LvTileviewGetTileActive() *LvObjT {
	return nil
}

/**
 * Create a spinbox object
 * @param parent    pointer to an object, it will be the parent of the new spinbox
 * @return          pointer to the created spinbox
 */
// llgo:link (*LvObjT).LvSpinboxCreate C.lv_spinbox_create
func (recv_ *LvObjT) LvSpinboxCreate() *LvObjT {
	return nil
}

/**
 * Set spinbox value
 * @param obj   pointer to spinbox
 * @param v     value to be set
 */
// llgo:link (*LvObjT).LvSpinboxSetValue C.lv_spinbox_set_value
func (recv_ *LvObjT) LvSpinboxSetValue(v c.Int32T) {
}

/**
 * Set spinbox rollover function
 * @param obj       pointer to spinbox
 * @param rollover  true or false to enable or disable (default)
 */
// llgo:link (*LvObjT).LvSpinboxSetRollover C.lv_spinbox_set_rollover
func (recv_ *LvObjT) LvSpinboxSetRollover(rollover bool) {
}

/**
 * Set spinbox digit format (digit count and decimal format)
 * @param obj           pointer to spinbox
 * @param digit_count   number of digit excluding the decimal separator and the sign
 * @param sep_pos       number of digit before the decimal point. If 0, decimal point is not
 * shown
 */
// llgo:link (*LvObjT).LvSpinboxSetDigitFormat C.lv_spinbox_set_digit_format
func (recv_ *LvObjT) LvSpinboxSetDigitFormat(digit_count c.Uint32T, sep_pos c.Uint32T) {
}

/**
 * Set spinbox step
 * @param obj   pointer to spinbox
 * @param step  steps on increment/decrement. Can be 1, 10, 100, 1000, etc the digit that will change.
 */
// llgo:link (*LvObjT).LvSpinboxSetStep C.lv_spinbox_set_step
func (recv_ *LvObjT) LvSpinboxSetStep(step c.Uint32T) {
}

/**
 * Set spinbox value range
 * @param obj       pointer to spinbox
 * @param range_min maximum value, inclusive
 * @param range_max minimum value, inclusive
 */
// llgo:link (*LvObjT).LvSpinboxSetRange C.lv_spinbox_set_range
func (recv_ *LvObjT) LvSpinboxSetRange(range_min c.Int32T, range_max c.Int32T) {
}

/**
 * Set cursor position to a specific digit for edition
 * @param obj   pointer to spinbox
 * @param pos   selected position in spinbox
 */
// llgo:link (*LvObjT).LvSpinboxSetCursorPos C.lv_spinbox_set_cursor_pos
func (recv_ *LvObjT) LvSpinboxSetCursorPos(pos c.Uint32T) {
}

/**
 * Set direction of digit step when clicking an encoder button while in editing mode
 * @param obj           pointer to spinbox
 * @param direction     the direction (LV_DIR_RIGHT or LV_DIR_LEFT)
 */
// llgo:link (*LvObjT).LvSpinboxSetDigitStepDirection C.lv_spinbox_set_digit_step_direction
func (recv_ *LvObjT) LvSpinboxSetDigitStepDirection(direction LvDirT) {
}

/**
 * Get spinbox rollover function status
 * @param obj   pointer to spinbox
 */
// llgo:link (*LvObjT).LvSpinboxGetRollover C.lv_spinbox_get_rollover
func (recv_ *LvObjT) LvSpinboxGetRollover() bool {
	return false
}

/**
 * Get the spinbox numeral value (user has to convert to float according to its digit format)
 * @param obj   pointer to spinbox
 * @return      value integer value of the spinbox
 */
// llgo:link (*LvObjT).LvSpinboxGetValue C.lv_spinbox_get_value
func (recv_ *LvObjT) LvSpinboxGetValue() c.Int32T {
	return 0
}

/**
 * Get the spinbox step value (user has to convert to float according to its digit format)
 * @param obj   pointer to spinbox
 * @return      value integer step value of the spinbox
 */
// llgo:link (*LvObjT).LvSpinboxGetStep C.lv_spinbox_get_step
func (recv_ *LvObjT) LvSpinboxGetStep() c.Int32T {
	return 0
}

/**
 * Select next lower digit for edition by dividing the step by 10
 * @param obj   pointer to spinbox
 */
// llgo:link (*LvObjT).LvSpinboxStepNext C.lv_spinbox_step_next
func (recv_ *LvObjT) LvSpinboxStepNext() {
}

/**
 * Select next higher digit for edition by multiplying the step by 10
 * @param obj   pointer to spinbox
 */
// llgo:link (*LvObjT).LvSpinboxStepPrev C.lv_spinbox_step_prev
func (recv_ *LvObjT) LvSpinboxStepPrev() {
}

/**
 * Increment spinbox value by one step
 * @param obj   pointer to spinbox
 */
// llgo:link (*LvObjT).LvSpinboxIncrement C.lv_spinbox_increment
func (recv_ *LvObjT) LvSpinboxIncrement() {
}

/**
 * Decrement spinbox value by one step
 * @param obj   pointer to spinbox
 */
// llgo:link (*LvObjT).LvSpinboxDecrement C.lv_spinbox_decrement
func (recv_ *LvObjT) LvSpinboxDecrement() {
}

type LvSpanOverflowT c.Int

const (
	LV_SPAN_OVERFLOW_CLIP     LvSpanOverflowT = 0
	LV_SPAN_OVERFLOW_ELLIPSIS LvSpanOverflowT = 1
	LV_SPAN_OVERFLOW_LAST     LvSpanOverflowT = 2
)

type LvSpanModeT c.Int

const (
	LV_SPAN_MODE_FIXED  LvSpanModeT = 0
	LV_SPAN_MODE_EXPAND LvSpanModeT = 1
	LV_SPAN_MODE_BREAK  LvSpanModeT = 2
	LV_SPAN_MODE_LAST   LvSpanModeT = 3
)

/** Coords of a span */

type X_lvSpanCoordsT struct {
	Heading  LvAreaT
	Middle   LvAreaT
	Trailing LvAreaT
}
type LvSpanCoordsT X_lvSpanCoordsT

/**********************
 * GLOBAL PROTOTYPES
 **********************/
//go:linkname LvSpanStackInit C.lv_span_stack_init
func LvSpanStackInit()

//go:linkname LvSpanStackDeinit C.lv_span_stack_deinit
func LvSpanStackDeinit()

/**
 * Create a spangroup object
 * @param parent    pointer to an object, it will be the parent of the new spangroup
 * @return          pointer to the created spangroup
 */
// llgo:link (*LvObjT).LvSpangroupCreate C.lv_spangroup_create
func (recv_ *LvObjT) LvSpangroupCreate() *LvObjT {
	return nil
}

/**
 * Create a span string descriptor and add to spangroup.
 * @param obj       pointer to a spangroup object.
 * @return          pointer to the created span.
 */
// llgo:link (*LvObjT).LvSpangroupAddSpan C.lv_spangroup_add_span
func (recv_ *LvObjT) LvSpangroupAddSpan() *LvSpanT {
	return nil
}

/**
 * Remove the span from the spangroup and free memory.
 * @param obj   pointer to a spangroup object.
 * @param span  pointer to a span.
 */
// llgo:link (*LvObjT).LvSpangroupDeleteSpan C.lv_spangroup_delete_span
func (recv_ *LvObjT) LvSpangroupDeleteSpan(span *LvSpanT) {
}

/**
 * Set a new text for a span. Memory will be allocated to store the text by the span.
 * As the spangroup is not passed a redraw (invalidation) can't be triggered automatically.
 * Therefore `lv_spangroup_refresh(spangroup)` needs to be called manually,
 * @param span  pointer to a span.
 * @param text  pointer to a text.
 */
// llgo:link (*LvSpanT).LvSpanSetText C.lv_span_set_text
func (recv_ *LvSpanT) LvSpanSetText(text *c.Char) {
}

/**
 * Set a static text. It will not be saved by the span so the 'text' variable
 * has to be 'alive' while the span exist.
 * As the spangroup is not passed a redraw (invalidation) can't be triggered automatically.
 * Therefore `lv_spangroup_refresh(spangroup)` needs to be called manually,
 *
 * @param span  pointer to a span.
 * @param text  pointer to a text.
 */
// llgo:link (*LvSpanT).LvSpanSetTextStatic C.lv_span_set_text_static
func (recv_ *LvSpanT) LvSpanSetTextStatic(text *c.Char) {
}

/**
 * Set a new text for a span. Memory will be allocated to store the text by the span.
 * @param obj   pointer to a spangroup widget.
 * @param span  pointer to a span.
 * @param text  pointer to a text.
 */
// llgo:link (*LvObjT).LvSpangroupSetSpanText C.lv_spangroup_set_span_text
func (recv_ *LvObjT) LvSpangroupSetSpanText(span *LvSpanT, text *c.Char) {
}

/**
 * Set a new text for a span. Memory will be allocated to store the text by the span.
 * @param obj   pointer to a spangroup widget.
 * @param span  pointer to a span.
 * @param text  pointer to a text.
 */
// llgo:link (*LvObjT).LvSpangroupSetSpanTextStatic C.lv_spangroup_set_span_text_static
func (recv_ *LvObjT) LvSpangroupSetSpanTextStatic(span *LvSpanT, text *c.Char) {
}

/**
 * Copy all style properties of style to the bbuilt-in static style of the span.
 * @param obj       pointer_to a spangroup
 * @param span      pointer to a span.
 * @param style     pointer to a style to copy into the span's built-in style
 */
// llgo:link (*LvObjT).LvSpangroupSetSpanStyle C.lv_spangroup_set_span_style
func (recv_ *LvObjT) LvSpangroupSetSpanStyle(span *LvSpanT, style *LvStyleT) {
}

/**
 * DEPRECATED. Use the text_align style property instead
 * Set the align of the spangroup.
 * @param obj   pointer to a spangroup object.
 * @param align see lv_text_align_t for details.
 */
// llgo:link (*LvObjT).LvSpangroupSetAlign C.lv_spangroup_set_align
func (recv_ *LvObjT) LvSpangroupSetAlign(align LvTextAlignT) {
}

/**
 * Set the overflow of the spangroup.
 * @param obj       pointer to a spangroup object.
 * @param overflow  see lv_span_overflow_t for details.
 */
// llgo:link (*LvObjT).LvSpangroupSetOverflow C.lv_spangroup_set_overflow
func (recv_ *LvObjT) LvSpangroupSetOverflow(overflow LvSpanOverflowT) {
}

/**
 * Set the indent of the spangroup.
 * @param obj       pointer to a spangroup object.
 * @param indent    the first line indentation
 */
// llgo:link (*LvObjT).LvSpangroupSetIndent C.lv_spangroup_set_indent
func (recv_ *LvObjT) LvSpangroupSetIndent(indent c.Int32T) {
}

/**
 * DEPRECATED, set the width to LV_SIZE_CONTENT or fixed value to control expanding/wrapping"
 * Set the mode of the spangroup.
 * @param obj       pointer to a spangroup object.
 * @param mode      see lv_span_mode_t for details.
 */
// llgo:link (*LvObjT).LvSpangroupSetMode C.lv_spangroup_set_mode
func (recv_ *LvObjT) LvSpangroupSetMode(mode LvSpanModeT) {
}

/**
 * Set maximum lines of the spangroup.
 * @param obj       pointer to a spangroup object.
 * @param lines     max lines that can be displayed in LV_SPAN_MODE_BREAK mode. < 0 means no limit.
 */
// llgo:link (*LvObjT).LvSpangroupSetMaxLines C.lv_spangroup_set_max_lines
func (recv_ *LvObjT) LvSpangroupSetMaxLines(lines c.Int32T) {
}

/**
 * Get a pointer to the style of a span's built-in style.
 * Any lv_style_set_... functions can be applied on the returned style.
 * @param span  pointer to the span
 * @return      pointer to the style. (valid as long as the span is valid)
 */
// llgo:link (*LvSpanT).LvSpanGetStyle C.lv_span_get_style
func (recv_ *LvSpanT) LvSpanGetStyle() *LvStyleT {
	return nil
}

/**
 * Get a pointer to the text of a span
 * @param span  pointer to the span
 * @return      pointer to the text
 */
// llgo:link (*LvSpanT).LvSpanGetText C.lv_span_get_text
func (recv_ *LvSpanT) LvSpanGetText() *c.Char {
	return nil
}

/**
 * Get a spangroup child by its index.
 *
 * @param obj   The spangroup object
 * @param id    the index of the child.
 *              0: the oldest (firstly created) child
 *              1: the second oldest
 *              child count-1: the youngest
 *              -1: the youngest
 *              -2: the second youngest
 * @return      The child span at index `id`, or NULL if the ID does not exist
 */
// llgo:link (*LvObjT).LvSpangroupGetChild C.lv_spangroup_get_child
func (recv_ *LvObjT) LvSpangroupGetChild(id c.Int32T) *LvSpanT {
	return nil
}

/**
 * Get number of spans
 * @param obj   the spangroup object to get the child count of.
 * @return      the span count of the spangroup.
 */
// llgo:link (*LvObjT).LvSpangroupGetSpanCount C.lv_spangroup_get_span_count
func (recv_ *LvObjT) LvSpangroupGetSpanCount() c.Uint32T {
	return 0
}

/**
 * Get the align of the spangroup.
 * @param obj   pointer to a spangroup object.
 * @return      the align value.
 */
// llgo:link (*LvObjT).LvSpangroupGetAlign C.lv_spangroup_get_align
func (recv_ *LvObjT) LvSpangroupGetAlign() LvTextAlignT {
	return 0
}

/**
 * Get the overflow of the spangroup.
 * @param obj   pointer to a spangroup object.
 * @return      the overflow value.
 */
// llgo:link (*LvObjT).LvSpangroupGetOverflow C.lv_spangroup_get_overflow
func (recv_ *LvObjT) LvSpangroupGetOverflow() LvSpanOverflowT {
	return 0
}

/**
 * Get the indent of the spangroup.
 * @param obj   pointer to a spangroup object.
 * @return      the indent value.
 */
// llgo:link (*LvObjT).LvSpangroupGetIndent C.lv_spangroup_get_indent
func (recv_ *LvObjT) LvSpangroupGetIndent() c.Int32T {
	return 0
}

/**
 * Get the mode of the spangroup.
 * @param obj   pointer to a spangroup object.
 */
// llgo:link (*LvObjT).LvSpangroupGetMode C.lv_spangroup_get_mode
func (recv_ *LvObjT) LvSpangroupGetMode() LvSpanModeT {
	return 0
}

/**
 * Get maximum lines of the spangroup.
 * @param obj   pointer to a spangroup object.
 * @return      the max lines value.
 */
// llgo:link (*LvObjT).LvSpangroupGetMaxLines C.lv_spangroup_get_max_lines
func (recv_ *LvObjT) LvSpangroupGetMaxLines() c.Int32T {
	return 0
}

/**
 * Get max line height of all span in the spangroup.
 * @param obj   pointer to a spangroup object.
 */
// llgo:link (*LvObjT).LvSpangroupGetMaxLineHeight C.lv_spangroup_get_max_line_height
func (recv_ *LvObjT) LvSpangroupGetMaxLineHeight() c.Int32T {
	return 0
}

/**
 * Get the text content width when all span of spangroup on a line.
 * @param obj       pointer to a spangroup object.
 * @param max_width if text content width >= max_width, return max_width
 * to reduce computation, if max_width == 0, returns the text content width.
 * @return text     content width or max_width.
 */
// llgo:link (*LvObjT).LvSpangroupGetExpandWidth C.lv_spangroup_get_expand_width
func (recv_ *LvObjT) LvSpangroupGetExpandWidth(max_width c.Uint32T) c.Uint32T {
	return 0
}

/**
 * Get the text content height with width fixed.
 * @param obj       pointer to a spangroup object.
 * @param width     the width of the span group.

 */
// llgo:link (*LvObjT).LvSpangroupGetExpandHeight C.lv_spangroup_get_expand_height
func (recv_ *LvObjT) LvSpangroupGetExpandHeight(width c.Int32T) c.Int32T {
	return 0
}

/**
 * Get the span's coords in the spangroup.
 * @note Before calling this function, please make sure that the layout of span group has been updated.
 * Like calling lv_obj_update_layout() like function.
 *
 *     +--------+
 *     |Heading +--->------------------+
 *     |  Pos   |   |     Heading      |
 *     +--------+---+------------------+
 *     |                               |
 *     |                               |
 *     |                               |
 *     |            Middle   +--------+|
 *     |                     |Trailing||
 *     |                   +-|  Pos   ||
 *     |                   | +--------+|
 *     +-------------------v-----------+
 *     |     Trailing      |
 *     +-------------------+
 * @param obj       pointer to a spangroup object.
 * @param span      pointer to a span.
 * @return the span's coords in the spangroup.
 */
// llgo:link (*LvObjT).LvSpangroupGetSpanCoords C.lv_spangroup_get_span_coords
func (recv_ *LvObjT) LvSpangroupGetSpanCoords(span *LvSpanT) LvSpanCoordsT {
	return LvSpanCoordsT{}
}

/**
 * Get the span object by point.
 * @param obj       pointer to a spangroup object.
 * @param point     pointer to point containing absolute coordinates
 * @return          pointer to the span under the point or `NULL` if not found.
 */
// llgo:link (*LvObjT).LvSpangroupGetSpanByPoint C.lv_spangroup_get_span_by_point
func (recv_ *LvObjT) LvSpangroupGetSpanByPoint(point *LvPointT) *LvSpanT {
	return nil
}

/**
 * Update the mode of the spangroup.
 * @param obj   pointer to a spangroup object.
 */
// llgo:link (*LvObjT).LvSpangroupRefresh C.lv_spangroup_refresh
func (recv_ *LvObjT) LvSpangroupRefresh() {
}

/**
 * Create a tabview widget
 * @param parent    pointer to a parent widget
 * @return          the created tabview
 */
// llgo:link (*LvObjT).LvTabviewCreate C.lv_tabview_create
func (recv_ *LvObjT) LvTabviewCreate() *LvObjT {
	return nil
}

/**
 * Add a tab to the tabview
 * @param obj       pointer to a tabview widget
 * @param name      the name of the tab, it will be displayed on the tab bar
 * @return          the widget where the content of the tab can be created
 */
// llgo:link (*LvObjT).LvTabviewAddTab C.lv_tabview_add_tab
func (recv_ *LvObjT) LvTabviewAddTab(name *c.Char) *LvObjT {
	return nil
}

/**
 * Change the name of the tab
 * @param obj       pointer to a tabview widget
 * @param idx       the index of the tab to rename
 * @param new_name  the new name as a string
 */
// llgo:link (*LvObjT).LvTabviewRenameTab C.lv_tabview_rename_tab
func (recv_ *LvObjT) LvTabviewRenameTab(idx c.Uint32T, new_name *c.Char) {
}

/**
 * Show a tab
 * @param obj       pointer to a tabview widget
 * @param idx       the index of the tab to show
 * @param anim_en   LV_ANIM_ON/OFF
 */
// llgo:link (*LvObjT).LvTabviewSetActive C.lv_tabview_set_active
func (recv_ *LvObjT) LvTabviewSetActive(idx c.Uint32T, anim_en LvAnimEnableT) {
}

/**
 * Set the position of the tab bar
 * @param obj       pointer to a tabview widget
 * @param dir       LV_DIR_TOP/BOTTOM/LEFT/RIGHT
 */
// llgo:link (*LvObjT).LvTabviewSetTabBarPosition C.lv_tabview_set_tab_bar_position
func (recv_ *LvObjT) LvTabviewSetTabBarPosition(dir LvDirT) {
}

/**
 * Set the width or height of the tab bar
 * @param obj       pointer to tabview widget
 * @param size      size of the tab bar in pixels or percentage.
 *                  will be used as width or height based on the position of the tab bar)
 */
// llgo:link (*LvObjT).LvTabviewSetTabBarSize C.lv_tabview_set_tab_bar_size
func (recv_ *LvObjT) LvTabviewSetTabBarSize(size c.Int32T) {
}

/**
 * Get the number of tabs
 * @param obj       pointer to a tabview widget
 * @return          the number of tabs
 */
// llgo:link (*LvObjT).LvTabviewGetTabCount C.lv_tabview_get_tab_count
func (recv_ *LvObjT) LvTabviewGetTabCount() c.Uint32T {
	return 0
}

/**
 * Get the current tab's index
 * @param obj       pointer to a tabview widget
 * @return          the zero based index of the current tab
 */
// llgo:link (*LvObjT).LvTabviewGetTabActive C.lv_tabview_get_tab_active
func (recv_ *LvObjT) LvTabviewGetTabActive() c.Uint32T {
	return 0
}

/**
 * Get the widget where the container of each tab is created
 * @param obj       pointer to a tabview widget
 * @return          the main container widget
 */
// llgo:link (*LvObjT).LvTabviewGetContent C.lv_tabview_get_content
func (recv_ *LvObjT) LvTabviewGetContent() *LvObjT {
	return nil
}

/**
 * Get the tab bar where the buttons are created
 * @param obj       pointer to a tabview widget
 * @return          the tab bar
 */
// llgo:link (*LvObjT).LvTabviewGetTabBar C.lv_tabview_get_tab_bar
func (recv_ *LvObjT) LvTabviewGetTabBar() *LvObjT {
	return nil
}

/**
 * Initialize the OS layer
 */
//go:linkname LvOsInit C.lv_os_init
func LvOsInit()

/**
 * Initialize LVGL library.
 * Should be called before any other LVGL related function.
 */
//go:linkname LvInit C.lv_init
func LvInit()

/**
 * Deinit the 'lv' library
 */
//go:linkname LvDeinit C.lv_deinit
func LvDeinit()

/**
 * Returns whether the 'lv' library is currently initialized
 */
//go:linkname LvIsInitialized C.lv_is_initialized
func LvIsInitialized() bool

// llgo:type C
type LvAsyncCbT func(c.Pointer)

/**
 * Call an asynchronous function the next time lv_timer_handler() is run. This function is likely to return
 * **before** the call actually happens!
 * @param async_xcb a callback which is the task itself.
 *                 (the 'x' in the argument name indicates that it's not a fully generic function because it not follows
 *                  the `func_name(object, callback, ...)` convention)
 * @param user_data custom parameter
 */
//go:linkname LvAsyncCall C.lv_async_call
func LvAsyncCall(async_xcb LvAsyncCbT, user_data c.Pointer) LvResultT

/**
 * Cancel an asynchronous function call
 * @param async_xcb a callback which is the task itself.
 * @param user_data custom parameter
 */
//go:linkname LvAsyncCallCancel C.lv_async_call_cancel
func LvAsyncCallCancel(async_xcb LvAsyncCbT, user_data c.Pointer) LvResultT

type X_lvAnimTimelineT struct {
	Unused [8]uint8
}
type LvAnimTimelineT X_lvAnimTimelineT

/**
 * Create an animation timeline.
 * @return pointer to the animation timeline.
 */
//go:linkname LvAnimTimelineCreate C.lv_anim_timeline_create
func LvAnimTimelineCreate() *LvAnimTimelineT

/**
 * Delete animation timeline.
 * @param at    pointer to the animation timeline.
 */
// llgo:link (*LvAnimTimelineT).LvAnimTimelineDelete C.lv_anim_timeline_delete
func (recv_ *LvAnimTimelineT) LvAnimTimelineDelete() {
}

/**
 * Add animation to the animation timeline.
 * @param at            pointer to the animation timeline.
 * @param start_time    the time the animation started on the timeline, note that start_time will override the value of delay.
 * @param a             pointer to an animation.
 */
// llgo:link (*LvAnimTimelineT).LvAnimTimelineAdd C.lv_anim_timeline_add
func (recv_ *LvAnimTimelineT) LvAnimTimelineAdd(start_time c.Uint32T, a *LvAnimT) {
}

/**
 * Start the animation timeline.
 * @param at    pointer to the animation timeline.
 * @return total time spent in animation timeline.
 */
// llgo:link (*LvAnimTimelineT).LvAnimTimelineStart C.lv_anim_timeline_start
func (recv_ *LvAnimTimelineT) LvAnimTimelineStart() c.Uint32T {
	return 0
}

/**
 * Pause the animation timeline.
 * @param at    pointer to the animation timeline.
 */
// llgo:link (*LvAnimTimelineT).LvAnimTimelinePause C.lv_anim_timeline_pause
func (recv_ *LvAnimTimelineT) LvAnimTimelinePause() {
}

/**
 * Set the playback direction of the animation timeline.
 * @param at        pointer to the animation timeline.
 * @param reverse   whether to play in reverse.
 */
// llgo:link (*LvAnimTimelineT).LvAnimTimelineSetReverse C.lv_anim_timeline_set_reverse
func (recv_ *LvAnimTimelineT) LvAnimTimelineSetReverse(reverse bool) {
}

/**
 * Make the animation timeline repeat itself.
 * @param at        pointer to the animation timeline.
 * @param cnt       repeat count or `LV_ANIM_REPEAT_INFINITE` for infinite repetition. 0: to disable repetition.
 */
// llgo:link (*LvAnimTimelineT).LvAnimTimelineSetRepeatCount C.lv_anim_timeline_set_repeat_count
func (recv_ *LvAnimTimelineT) LvAnimTimelineSetRepeatCount(cnt c.Uint32T) {
}

/**
 * Set a delay before repeating the animation timeline.
 * @param at        pointer to the animation timeline.
 * @param delay     delay in milliseconds before repeating the animation timeline.
 */
// llgo:link (*LvAnimTimelineT).LvAnimTimelineSetRepeatDelay C.lv_anim_timeline_set_repeat_delay
func (recv_ *LvAnimTimelineT) LvAnimTimelineSetRepeatDelay(delay c.Uint32T) {
}

/**
 * Set the progress of the animation timeline.
 * @param at        pointer to the animation timeline.
 * @param progress  set value 0~65535 to map 0~100% animation progress.
 */
// llgo:link (*LvAnimTimelineT).LvAnimTimelineSetProgress C.lv_anim_timeline_set_progress
func (recv_ *LvAnimTimelineT) LvAnimTimelineSetProgress(progress c.Uint16T) {
}

/**
 * Get the time used to play the animation timeline.
 * @param at    pointer to the animation timeline.
 * @return total time spent in animation timeline.
 */
// llgo:link (*LvAnimTimelineT).LvAnimTimelineGetPlaytime C.lv_anim_timeline_get_playtime
func (recv_ *LvAnimTimelineT) LvAnimTimelineGetPlaytime() c.Uint32T {
	return 0
}

/**
 * Get whether the animation timeline is played in reverse.
 * @param at    pointer to the animation timeline.
 * @return return true if it is reverse playback.
 */
// llgo:link (*LvAnimTimelineT).LvAnimTimelineGetReverse C.lv_anim_timeline_get_reverse
func (recv_ *LvAnimTimelineT) LvAnimTimelineGetReverse() bool {
	return false
}

/**
 * Get the progress of the animation timeline.
 * @param at    pointer to the animation timeline.
 * @return return value 0~65535 to map 0~100% animation progress.
 */
// llgo:link (*LvAnimTimelineT).LvAnimTimelineGetProgress C.lv_anim_timeline_get_progress
func (recv_ *LvAnimTimelineT) LvAnimTimelineGetProgress() c.Uint16T {
	return 0
}

/**
 * Get repeat count of the animation timeline.
 * @param at    pointer to the animation timeline.
 */
// llgo:link (*LvAnimTimelineT).LvAnimTimelineGetRepeatCount C.lv_anim_timeline_get_repeat_count
func (recv_ *LvAnimTimelineT) LvAnimTimelineGetRepeatCount() c.Uint32T {
	return 0
}

/**
 * Get repeat delay of the animation timeline.
 * @param at    pointer to the animation timeline.
 */
// llgo:link (*LvAnimTimelineT).LvAnimTimelineGetRepeatDelay C.lv_anim_timeline_get_repeat_delay
func (recv_ *LvAnimTimelineT) LvAnimTimelineGetRepeatDelay() c.Uint32T {
	return 0
}

/** Searches base[0] to base[n - 1] for an item that matches *key.
 *
 * @note The function cmp must return negative if it's first
 *  argument (the search key) is less that it's second (a table entry),
 *  zero if equal, and positive if greater.
 *
 * @note Items in the array must be in ascending order.
 *
 * @param key    Pointer to item being searched for
 * @param base   Pointer to first element to search
 * @param n      Number of elements
 * @param size   Size of each element
 * @param cmp    Pointer to comparison function (see unicode_list_compare()
 *                 as a comparison function example)
 *
 * @return a pointer to a matching item, or NULL if none exists.
 */
//go:linkname LvUtilsBsearch C.lv_utils_bsearch
func LvUtilsBsearch(key c.Pointer, base c.Pointer, n c.SizeT, size c.SizeT, cmp func(c.Pointer, c.Pointer) c.Int) c.Pointer

/**
 * Save a draw buf to a file
 * @param draw_buf  pointer to a draw buffer
 * @param path      path to the file to save
 * @return          LV_RESULT_OK: success; LV_RESULT_INVALID: error
 */
// llgo:link (*LvDrawBufT).LvDrawBufSaveToFile C.lv_draw_buf_save_to_file
func (recv_ *LvDrawBufT) LvDrawBufSaveToFile(path *c.Char) LvResultT {
	return 0
}

// llgo:type C
type LvIterNextCb func(c.Pointer, c.Pointer, c.Pointer) LvResultT

// llgo:type C
type LvIterInspectCb func(c.Pointer)

/**
 * Create an iterator based on an instance, and then the next element of the iterator can be obtained through lv_iter_next,
 * In order to obtain the next operation in a unified and abstract way.
 * @param instance       The instance to be iterated
 * @param elem_size      The size of the element to be iterated in bytes
 * @param context_size   The size of the context to be passed to the next_cb in bytes
 * @param next_cb        The callback function to get the next element
 * @return               The iterator object
 */
//go:linkname LvIterCreate C.lv_iter_create
func LvIterCreate(instance c.Pointer, elem_size c.Uint32T, context_size c.Uint32T, next_cb LvIterNextCb) *LvIterT

/**
 * Get the context of the iterator. You can use it to store some temporary variables associated with current iterator..
 * @param iter           `lv_iter_t` object create before
 * @return the iter context
 */
// llgo:link (*LvIterT).LvIterGetContext C.lv_iter_get_context
func (recv_ *LvIterT) LvIterGetContext() c.Pointer {
	return nil
}

/**
 * Destroy the iterator object, and release the context. Other resources allocated by the user are not released.
 * The user needs to release it by itself.
 * @param iter          `lv_iter_t` object create before
 */
// llgo:link (*LvIterT).LvIterDestroy C.lv_iter_destroy
func (recv_ *LvIterT) LvIterDestroy() {
}

/**
 * Get the next element of the iterator.
 * @param iter          `lv_iter_t` object create before
 * @param elem          The pointer to store the next element
 * @return              LV_RESULT_OK: Get the next element successfully
 *                      LV_RESULT_INVALID: The next element is invalid
 */
// llgo:link (*LvIterT).LvIterNext C.lv_iter_next
func (recv_ *LvIterT) LvIterNext(elem c.Pointer) LvResultT {
	return 0
}

/**
 * Make the iterator peekable, which means that the user can peek the next element without advancing the iterator.
 * @param iter          `lv_iter_t` object create before
 * @param capacity      The capacity of the peek buffer
 */
// llgo:link (*LvIterT).LvIterMakePeekable C.lv_iter_make_peekable
func (recv_ *LvIterT) LvIterMakePeekable(capacity c.Uint32T) {
}

/**
 * Peek the next element of the iterator without advancing the iterator.
 * @param iter          `lv_iter_t` object create before
 * @param elem          The pointer to store the next element
 * @return              LV_RESULT_OK: Peek the next element successfully
 *                      LV_RESULT_INVALID: The next element is invalid
 */
// llgo:link (*LvIterT).LvIterPeek C.lv_iter_peek
func (recv_ *LvIterT) LvIterPeek(elem c.Pointer) LvResultT {
	return 0
}

/**
 * Only advance the iterator without getting the next element.
 * @param iter          `lv_iter_t` object create before
 * @return              LV_RESULT_OK: Peek the next element successfully
 *                      LV_RESULT_INVALID: The next element is invalid
 */
// llgo:link (*LvIterT).LvIterPeekAdvance C.lv_iter_peek_advance
func (recv_ *LvIterT) LvIterPeekAdvance() LvResultT {
	return 0
}

/**
 * Reset the peek cursor to the `next` cursor.
 * @param iter          `lv_iter_t` object create before
 * @return              LV_RESULT_OK: Reset the peek buffer successfully
 *                      LV_RESULT_INVALID: The peek buffer is invalid
 */
// llgo:link (*LvIterT).LvIterPeekReset C.lv_iter_peek_reset
func (recv_ *LvIterT) LvIterPeekReset() LvResultT {
	return 0
}

/**
 * Inspect the element of the iterator. The callback function will be called for each element of the iterator.
 * @param iter          `lv_iter_t` object create before
 * @param inspect_cb    The callback function to inspect the element
 */
// llgo:link (*LvIterT).LvIterInspect C.lv_iter_inspect
func (recv_ *LvIterT) LvIterInspect(inspect_cb LvIterInspectCb) {
}

// llgo:type C
type LvCircleBufFillCbT func(c.Pointer, c.Uint32T, c.Int32T, c.Pointer) bool

/**
 * Create a circle buffer
 * @param capacity the maximum number of elements in the buffer
 * @param element_size the size of an element in bytes
 * @return pointer to the created buffer
 */
//go:linkname LvCircleBufCreate C.lv_circle_buf_create
func LvCircleBufCreate(capacity c.Uint32T, element_size c.Uint32T) *LvCircleBufT

/**
 * Create a circle buffer from an existing buffer
 * @param buf pointer to a buffer
 * @param capacity the maximum number of elements in the buffer
 * @param element_size the size of an element in bytes
 * @return pointer to the created buffer
 */
//go:linkname LvCircleBufCreateFromBuf C.lv_circle_buf_create_from_buf
func LvCircleBufCreateFromBuf(buf c.Pointer, capacity c.Uint32T, element_size c.Uint32T) *LvCircleBufT

/**
 * Create a circle buffer from an existing array
 * @param array pointer to an array
 * @return pointer to the created buffer
 */
// llgo:link (*LvArrayT).LvCircleBufCreateFromArray C.lv_circle_buf_create_from_array
func (recv_ *LvArrayT) LvCircleBufCreateFromArray() *LvCircleBufT {
	return nil
}

/**
 * Resize the buffer
 * @param circle_buf pointer to a buffer
 * @param capacity the new capacity of the buffer
 * @return LV_RESULT_OK: the buffer is resized; LV_RESULT_INVALID: the buffer is not resized
 */
// llgo:link (*LvCircleBufT).LvCircleBufResize C.lv_circle_buf_resize
func (recv_ *LvCircleBufT) LvCircleBufResize(capacity c.Uint32T) LvResultT {
	return 0
}

/**
 * Destroy a circle buffer
 * @param circle_buf pointer to buffer
 */
// llgo:link (*LvCircleBufT).LvCircleBufDestroy C.lv_circle_buf_destroy
func (recv_ *LvCircleBufT) LvCircleBufDestroy() {
}

/**
 * Get the size of the buffer
 * @param circle_buf pointer to buffer
 * @return the number of elements in the buffer
 */
// llgo:link (*LvCircleBufT).LvCircleBufSize C.lv_circle_buf_size
func (recv_ *LvCircleBufT) LvCircleBufSize() c.Uint32T {
	return 0
}

/**
 * Get the capacity of the buffer
 * @param circle_buf pointer to buffer
 * @return the maximum number of elements in the buffer
 */
// llgo:link (*LvCircleBufT).LvCircleBufCapacity C.lv_circle_buf_capacity
func (recv_ *LvCircleBufT) LvCircleBufCapacity() c.Uint32T {
	return 0
}

/**
 * Get the remaining space in the buffer
 * @param circle_buf pointer to buffer
 * @return the number of elements that can be written to the buffer
 */
// llgo:link (*LvCircleBufT).LvCircleBufRemain C.lv_circle_buf_remain
func (recv_ *LvCircleBufT) LvCircleBufRemain() c.Uint32T {
	return 0
}

/**
 * Check if the buffer is empty
 * @param circle_buf pointer to buffer
 * @return true: the buffer is empty; false: the buffer is not empty
 */
// llgo:link (*LvCircleBufT).LvCircleBufIsEmpty C.lv_circle_buf_is_empty
func (recv_ *LvCircleBufT) LvCircleBufIsEmpty() bool {
	return false
}

/**
 * Check if the buffer is full
 * @param circle_buf pointer to buffer
 * @return true: the buffer is full; false: the buffer is not full
 */
// llgo:link (*LvCircleBufT).LvCircleBufIsFull C.lv_circle_buf_is_full
func (recv_ *LvCircleBufT) LvCircleBufIsFull() bool {
	return false
}

/**
 * Reset the buffer
 * @param circle_buf pointer to buffer
 * @return LV_RESULT_OK: the buffer is reset; LV_RESULT_INVALID: the buffer is not reset
 */
// llgo:link (*LvCircleBufT).LvCircleBufReset C.lv_circle_buf_reset
func (recv_ *LvCircleBufT) LvCircleBufReset() {
}

/**
 * Get the head of the buffer
 * @param circle_buf pointer to buffer
 * @return pointer to the head of the buffer
 */
// llgo:link (*LvCircleBufT).LvCircleBufHead C.lv_circle_buf_head
func (recv_ *LvCircleBufT) LvCircleBufHead() c.Pointer {
	return nil
}

/**
 * Get the tail of the buffer
 * @param circle_buf pointer to buffer
 * @return pointer to the tail of the buffer
 */
// llgo:link (*LvCircleBufT).LvCircleBufTail C.lv_circle_buf_tail
func (recv_ *LvCircleBufT) LvCircleBufTail() c.Pointer {
	return nil
}

/**
 * Read a value
 * @param circle_buf pointer to buffer
 * @param data pointer to a variable to store the read value
 * @return LV_RESULT_OK: the value is read; LV_RESULT_INVALID: the value is not read
 */
// llgo:link (*LvCircleBufT).LvCircleBufRead C.lv_circle_buf_read
func (recv_ *LvCircleBufT) LvCircleBufRead(data c.Pointer) LvResultT {
	return 0
}

/**
 * Write a value
 * @param circle_buf pointer to buffer
 * @param data pointer to the value to write
 * @return LV_RESULT_OK: the value is written; LV_RESULT_INVALID: the value is not written
 */
// llgo:link (*LvCircleBufT).LvCircleBufWrite C.lv_circle_buf_write
func (recv_ *LvCircleBufT) LvCircleBufWrite(data c.Pointer) LvResultT {
	return 0
}

/**
 * Fill the buffer with values
 * @param circle_buf pointer to buffer
 * @param count the number of values to fill
 * @param fill_cb the callback function to fill the buffer
 * @param user_data
 * @return the number of values filled
 */
// llgo:link (*LvCircleBufT).LvCircleBufFill C.lv_circle_buf_fill
func (recv_ *LvCircleBufT) LvCircleBufFill(count c.Uint32T, fill_cb LvCircleBufFillCbT, user_data c.Pointer) c.Uint32T {
	return 0
}

/**
 * Skip a value
 * @param circle_buf pointer to buffer
 * @return LV_RESULT_OK: the value is skipped; LV_RESULT_INVALID: the value is not skipped
 */
// llgo:link (*LvCircleBufT).LvCircleBufSkip C.lv_circle_buf_skip
func (recv_ *LvCircleBufT) LvCircleBufSkip() LvResultT {
	return 0
}

/**
 * Peek a value
 * @param circle_buf pointer to buffer
 * @param data pointer to a variable to store the peeked value
 * @return LV_RESULT_OK: the value is peeked; LV_RESULT_INVALID: the value is not peeked
 */
// llgo:link (*LvCircleBufT).LvCircleBufPeek C.lv_circle_buf_peek
func (recv_ *LvCircleBufT) LvCircleBufPeek(data c.Pointer) LvResultT {
	return 0
}

/**
 * Peek a value at an index
 * @param circle_buf pointer to buffer
 * @param index the index of the value to peek, if the index is greater than the size of the buffer, it will return looply.
 * @param data pointer to a variable to store the peeked value
 * @return LV_RESULT_OK: the value is peeked; LV_RESULT_INVALID: the value is not peeked
 */
// llgo:link (*LvCircleBufT).LvCircleBufPeekAt C.lv_circle_buf_peek_at
func (recv_ *LvCircleBufT) LvCircleBufPeekAt(index c.Uint32T, data c.Pointer) LvResultT {
	return 0
}

/**********************
 *      TYPEDEFS
 **********************/

type X_lvTreeNodeT struct {
	Parent   *X_lvTreeNodeT
	Children **X_lvTreeNodeT
	ChildCnt c.Uint32T
	ChildCap c.Uint32T
	ClassP   *X_lvTreeClassT
}

/**
 * Describe the common methods of every object.
 * Similar to a C++ class.
 */

type X_lvTreeClassT struct {
	BaseClass     *X_lvTreeClassT
	InstanceSize  c.Uint32T
	ConstructorCb c.Pointer
	DestructorCb  c.Pointer
}
type LvTreeClassT X_lvTreeClassT
type LvTreeNodeT X_lvTreeNodeT
type X_lvTreeWalkModeT c.Int

const (
	LV_TREE_WALK_PRE_ORDER  X_lvTreeWalkModeT = 0
	LV_TREE_WALK_POST_ORDER X_lvTreeWalkModeT = 1
)

type LvTreeWalkModeT c.Uint8T

// llgo:type C
type LvTreeTraverseCbT func(*LvTreeNodeT, c.Pointer) bool

// llgo:type C
type LvTreeBeforeCbT func(*LvTreeNodeT, c.Pointer) bool

// llgo:type C
type LvTreeAfterCbT func(*LvTreeNodeT, c.Pointer)

/**
 * @brief Create a tree node
 * @param class_p pointer to a class of the node
 * @param parent pointer to the parent node (or NULL if it's the root node)
 * @return pointer to the new node
 */
// llgo:link (*LvTreeClassT).LvTreeNodeCreate C.lv_tree_node_create
func (recv_ *LvTreeClassT) LvTreeNodeCreate(parent *LvTreeNodeT) *LvTreeNodeT {
	return nil
}

/**
 * @brief Delete a tree node and all its children recursively
 * @param node pointer to the node to delete
 */
// llgo:link (*LvTreeNodeT).LvTreeNodeDelete C.lv_tree_node_delete
func (recv_ *LvTreeNodeT) LvTreeNodeDelete() {
}

/**
 * @brief Walk the tree recursively and call a callback function on each node
 * @param node pointer to the root node of the tree
 * @param mode LV_TREE_WALK_PRE_ORDER or LV_TREE_WALK_POST_ORDER
 * @param cb callback function to call on each node
 * @param bcb callback function to call before visiting a node
 * @param acb callback function to call after visiting a node
 * @param user_data user data to pass to the callback functions
 * @return true: traversal is finished; false: traversal broken
 */
// llgo:link (*LvTreeNodeT).LvTreeWalk C.lv_tree_walk
func (recv_ *LvTreeNodeT) LvTreeWalk(mode LvTreeWalkModeT, cb LvTreeTraverseCbT, bcb LvTreeBeforeCbT, acb LvTreeAfterCbT, user_data c.Pointer) bool {
	return false
}

/**********************
 *      TYPEDEFS
 **********************/

type LvBinfontFontSrcT struct {
	FontSize   c.Uint32T
	Path       *c.Char
	Buffer     c.Pointer
	BufferSize c.Uint32T
}

/**
 * Loads a `lv_font_t` object from a binary font file
 * @param path   path to font file
 * @return  pointer to font where to load
 */
//go:linkname LvBinfontCreate C.lv_binfont_create
func LvBinfontCreate(path *c.Char) *LvFontT

/**
 * Frees the memory allocated by the `lv_binfont_create()` function
 * @param font          lv_font_t object created by the lv_binfont_create function
 */
// llgo:link (*LvFontT).LvBinfontDestroy C.lv_binfont_destroy
func (recv_ *LvFontT) LvBinfontDestroy() {
}

type LvArclabelDirT c.Int

const (
	LV_ARCLABEL_DIR_CLOCKWISE         LvArclabelDirT = 0
	LV_ARCLABEL_DIR_COUNTER_CLOCKWISE LvArclabelDirT = 1
)

type LvArclabelTextAlignT c.Int

const (
	LV_ARCLABEL_TEXT_ALIGN_DEFAULT  LvArclabelTextAlignT = 0
	LV_ARCLABEL_TEXT_ALIGN_LEADING  LvArclabelTextAlignT = 1
	LV_ARCLABEL_TEXT_ALIGN_CENTER   LvArclabelTextAlignT = 2
	LV_ARCLABEL_TEXT_ALIGN_TRAILING LvArclabelTextAlignT = 3
)

/**
 * Create an arc label object
 * @param parent    pointer to an object, it will be the parent of the new arc label
 * @return          pointer to the created arc label
 */
// llgo:link (*LvObjT).LvArclabelCreate C.lv_arclabel_create
func (recv_ *LvObjT) LvArclabelCreate() *LvObjT {
	return nil
}

/**
 * Set the text of the arc label.
 *
 * This function sets the text displayed by an arc label object.
 *
 * @param obj       Pointer to the arc label object.
 * @param text      Pointer to a null-terminated string containing the new text for the label.
 */
// llgo:link (*LvObjT).LvArclabelSetText C.lv_arclabel_set_text
func (recv_ *LvObjT) LvArclabelSetText(text *c.Char) {
}

/**
 * Set the formatted text of an arc label object.
 *
 * This function sets the text of an arc label object with support for
 * variable arguments formatting, similar to `printf`.
 *
 * @param obj       The arc label object to set the text for.
 * @param fmt       A format string that specifies how subsequent arguments are converted to text.
 * @param ...       Arguments following the format string that are used to replace format specifiers in the format string.
 */
// llgo:link (*LvObjT).LvArclabelSetTextFmt C.lv_arclabel_set_text_fmt
func (recv_ *LvObjT) LvArclabelSetTextFmt(fmt *c.Char, __llgo_va_list ...interface{}) {
}

/**
 * Sets a new static text for the arc label or refreshes it with the current text.
 * The 'text' must remain valid in memory; the arc label does not manage its lifecycle.
 *
 * @param obj       Pointer to the arc label object.
 * @param text      Pointer to the new text. If NULL, the label is refreshed with its current text.
 */
// llgo:link (*LvObjT).LvArclabelSetTextStatic C.lv_arclabel_set_text_static
func (recv_ *LvObjT) LvArclabelSetTextStatic(text *c.Char) {
}

/**
 * Set the start angle of an arc. 0 deg: right, 90 bottom, etc.
 * @param obj       pointer to an arc label object
 * @param start     the start angle. (if `LV_USE_FLOAT` is enabled it can be fractional too.)
 */
// llgo:link (*LvObjT).LvArclabelSetAngleStart C.lv_arclabel_set_angle_start
func (recv_ *LvObjT) LvArclabelSetAngleStart(start LvValuePreciseT) {
}

/**
 * Set the end angle of an arc. 0 deg: right, 90 bottom, etc.
 * @param obj       pointer to an arc label object
 * @param size      the angle size (if `LV_USE_FLOAT` is enabled it can be fractional too.)
 */
// llgo:link (*LvObjT).LvArclabelSetAngleSize C.lv_arclabel_set_angle_size
func (recv_ *LvObjT) LvArclabelSetAngleSize(size LvValuePreciseT) {
}

/**
 * Set the rotation for the whole arc
 * @param obj       pointer to an arc label object
 * @param offset    rotation angle
 */
// llgo:link (*LvObjT).LvArclabelSetOffset C.lv_arclabel_set_offset
func (recv_ *LvObjT) LvArclabelSetOffset(offset c.Int32T) {
}

/**
 * Set the type of arc.
 * @param obj       pointer to and arc label object
 * @param dir      arc label's direction
 */
// llgo:link (*LvObjT).LvArclabelSetDir C.lv_arclabel_set_dir
func (recv_ *LvObjT) LvArclabelSetDir(dir LvArclabelDirT) {
}

/**
 * Enable the recoloring by in-line commands
 * @param obj       pointer to an arc label object
 * @param en        true: enable recoloring, false: disable
 * Example: "This is a #ff0000 red# word"
 */
// llgo:link (*LvObjT).LvArclabelSetRecolor C.lv_arclabel_set_recolor
func (recv_ *LvObjT) LvArclabelSetRecolor(en bool) {
}

/**
 * Set the radius for an arc label object.
 *
 * @param obj       pointer to the arc label object.
 * @param radius    The radius value to set for the label's curvature, in pixels.
 */
// llgo:link (*LvObjT).LvArclabelSetRadius C.lv_arclabel_set_radius
func (recv_ *LvObjT) LvArclabelSetRadius(radius c.Uint32T) {
}

/**
 * Set the center offset x for an arc label object.
 * @param obj       pointer to an arc label object
 * @param x         the x offset
 */
// llgo:link (*LvObjT).LvArclabelSetCenterOffsetX C.lv_arclabel_set_center_offset_x
func (recv_ *LvObjT) LvArclabelSetCenterOffsetX(x c.Uint32T) {
}

/**
 * Set the center offset y for an arc label object.
 * @param obj       pointer to an arc label object
 * @param y         the y offset
 */
// llgo:link (*LvObjT).LvArclabelSetCenterOffsetY C.lv_arclabel_set_center_offset_y
func (recv_ *LvObjT) LvArclabelSetCenterOffsetY(y c.Uint32T) {
}

/**
 * Set the text vertical alignment for an arc label object.
 * @param obj       pointer to an arc label object
 * @param align     the vertical alignment
 */
// llgo:link (*LvObjT).LvArclabelSetTextVerticalAlign C.lv_arclabel_set_text_vertical_align
func (recv_ *LvObjT) LvArclabelSetTextVerticalAlign(align LvArclabelTextAlignT) {
}

/**
 * Set the text horizontal alignment for an arc label object.
 * @param obj       pointer to an arc label object
 * @param align     the horizontal alignment
 */
// llgo:link (*LvObjT).LvArclabelSetTextHorizontalAlign C.lv_arclabel_set_text_horizontal_align
func (recv_ *LvObjT) LvArclabelSetTextHorizontalAlign(align LvArclabelTextAlignT) {
}

/**
 * Get the start angle of an arc label.
 * @param obj       pointer to an arc label object
 * @return          the start angle [0..360]  (if `LV_USE_FLOAT` is enabled it can be fractional too.)
 */
// llgo:link (*LvObjT).LvArclabelGetAngleStart C.lv_arclabel_get_angle_start
func (recv_ *LvObjT) LvArclabelGetAngleStart() LvValuePreciseT {
	return 0
}

/**
 * Get the angle size of an arc label.
 * @param obj       pointer to an arc label object
 * @return          the end angle [0..360]  (if `LV_USE_FLOAT` is enabled it can be fractional too.)
 */
// llgo:link (*LvObjT).LvArclabelGetAngleSize C.lv_arclabel_get_angle_size
func (recv_ *LvObjT) LvArclabelGetAngleSize() LvValuePreciseT {
	return 0
}

/**
 * Get whether the arc label is type or not.
 * @param obj       pointer to an arc label object
 * @return          arc label's direction
 */
// llgo:link (*LvObjT).LvArclabelGetDir C.lv_arclabel_get_dir
func (recv_ *LvObjT) LvArclabelGetDir() LvArclabelDirT {
	return 0
}

/**
 * Enable the recoloring by in-line commands
 *
 * @see lv_arclabel_set_recolor
 *
 * @param obj       pointer to a label object
 * @return          true: enable recoloring, false: disable
 */
// llgo:link (*LvObjT).LvArclabelGetRecolor C.lv_arclabel_get_recolor
func (recv_ *LvObjT) LvArclabelGetRecolor() bool {
	return false
}

/**
 * Get the text of the arc label.
 * @param obj       pointer to an arc label object
 * @return          the radius of the arc label
 */
// llgo:link (*LvObjT).LvArclabelGetRadius C.lv_arclabel_get_radius
func (recv_ *LvObjT) LvArclabelGetRadius() c.Uint32T {
	return 0
}

/**
 * Get the center offset x for an arc label object.
 * @param obj       pointer to an arc label object
 * @return          the x offset
 */
// llgo:link (*LvObjT).LvArclabelGetCenterOffsetX C.lv_arclabel_get_center_offset_x
func (recv_ *LvObjT) LvArclabelGetCenterOffsetX() c.Uint32T {
	return 0
}

/**
 * Get the center offset y for an arc label object.
 * @param obj       pointer to an arc label object
 * @return          the y offset
 */
// llgo:link (*LvObjT).LvArclabelGetCenterOffsetY C.lv_arclabel_get_center_offset_y
func (recv_ *LvObjT) LvArclabelGetCenterOffsetY() c.Uint32T {
	return 0
}

/**
 * Get the text vertical alignment for an arc label object.
 * @param obj       pointer to an arc label object
 * @return          the vertical alignment
 */
// llgo:link (*LvObjT).LvArclabelGetTextVerticalAlign C.lv_arclabel_get_text_vertical_align
func (recv_ *LvObjT) LvArclabelGetTextVerticalAlign() LvArclabelTextAlignT {
	return 0
}

/**
 * Get the text horizontal alignment for an arc label object.
 * @param obj       pointer to an arc label object
 * @return          the horizontal alignment
 */
// llgo:link (*LvObjT).LvArclabelGetTextHorizontalAlign C.lv_arclabel_get_text_horizontal_align
func (recv_ *LvObjT) LvArclabelGetTextHorizontalAlign() LvArclabelTextAlignT {
	return 0
}

/**
 * Create a list object
 * @param parent    pointer to an object, it will be the parent of the new list
 * @return          pointer to the created list
 */
// llgo:link (*LvObjT).LvListCreate C.lv_list_create
func (recv_ *LvObjT) LvListCreate() *LvObjT {
	return nil
}

/**
 * Add text to a list
 * @param list      pointer to a list, it will be the parent of the new label
 * @param txt       text of the new label
 * @return          pointer to the created label
 */
// llgo:link (*LvObjT).LvListAddText C.lv_list_add_text
func (recv_ *LvObjT) LvListAddText(txt *c.Char) *LvObjT {
	return nil
}

/**
 * Add button to a list
 * @param list      pointer to a list, it will be the parent of the new button
 * @param icon      icon for the button, when NULL it will have no icon
 * @param txt       text of the new button, when NULL no text will be added
 * @return          pointer to the created button
 */
// llgo:link (*LvObjT).LvListAddButton C.lv_list_add_button
func (recv_ *LvObjT) LvListAddButton(icon c.Pointer, txt *c.Char) *LvObjT {
	return nil
}

/**
 * Get text of a given list button
 * @param list      pointer to a list
 * @param btn       pointer to the button
 * @return          text of btn, if btn doesn't have text "" will be returned
 */
// llgo:link (*LvObjT).LvListGetButtonText C.lv_list_get_button_text
func (recv_ *LvObjT) LvListGetButtonText(btn *LvObjT) *c.Char {
	return nil
}

/**
 * Set text of a given list button
 * @param list      pointer to a list
 * @param btn       pointer to the button
 * @param txt       pointer to the text
 */
// llgo:link (*LvObjT).LvListSetButtonText C.lv_list_set_button_text
func (recv_ *LvObjT) LvListSetButtonText(btn *LvObjT, txt *c.Char) {
}

/**
 * Create a spinner widget
 * @param parent    pointer to an object, it will be the parent of the new spinner.
 * @return          the created spinner
 */
// llgo:link (*LvObjT).LvSpinnerCreate C.lv_spinner_create
func (recv_ *LvObjT) LvSpinnerCreate() *LvObjT {
	return nil
}

/**
 * Set the animation time and arc length of the spinner
 * @param obj       pointer to a spinner
 * @param t         the animation time in milliseconds
 * @param angle     the angle of the arc in degrees
 */
// llgo:link (*LvObjT).LvSpinnerSetAnimParams C.lv_spinner_set_anim_params
func (recv_ *LvObjT) LvSpinnerSetAnimParams(t c.Uint32T, angle c.Uint32T) {
}

/**
 * Initialize the binary image decoder module
 */
//go:linkname LvBinDecoderInit C.lv_bin_decoder_init
func LvBinDecoderInit()

/**
 * Get info about a lvgl binary image
 * @param decoder the decoder where this function belongs
 * @param dsc image descriptor containing the source and type of the image and other info.
 * @param header store the image data here
 * @return LV_RESULT_OK: the info is successfully stored in `header`; LV_RESULT_INVALID: unknown format or other error.
 */
// llgo:link (*LvImageDecoderT).LvBinDecoderInfo C.lv_bin_decoder_info
func (recv_ *LvImageDecoderT) LvBinDecoderInfo(dsc *LvImageDecoderDscT, header *LvImageHeaderT) LvResultT {
	return 0
}

// llgo:link (*LvImageDecoderT).LvBinDecoderGetArea C.lv_bin_decoder_get_area
func (recv_ *LvImageDecoderT) LvBinDecoderGetArea(dsc *LvImageDecoderDscT, full_area *LvAreaT, decoded_area *LvAreaT) LvResultT {
	return 0
}

/**
 * Open a lvgl binary image
 * @param decoder the decoder where this function belongs
 * @param dsc pointer to decoder descriptor. `src`, `style` are already initialized in it.
 * @return LV_RESULT_OK: the info is successfully stored in `header`; LV_RESULT_INVALID: unknown format or other error.
 */
// llgo:link (*LvImageDecoderT).LvBinDecoderOpen C.lv_bin_decoder_open
func (recv_ *LvImageDecoderT) LvBinDecoderOpen(dsc *LvImageDecoderDscT) LvResultT {
	return 0
}

/**
 * Close the pending decoding. Free resources etc.
 * @param decoder pointer to the decoder the function associated with
 * @param dsc pointer to decoder descriptor
 */
// llgo:link (*LvImageDecoderT).LvBinDecoderClose C.lv_bin_decoder_close
func (recv_ *LvImageDecoderT) LvBinDecoderClose(dsc *LvImageDecoderDscT) {
}

/**********************
 * GLOBAL PROTOTYPES
 **********************/
//go:linkname LvSdlWindowCreate C.lv_sdl_window_create
func LvSdlWindowCreate(hor_res c.Int32T, ver_res c.Int32T) *LvDisplayT

// llgo:link (*LvDisplayT).LvSdlWindowSetResizeable C.lv_sdl_window_set_resizeable
func (recv_ *LvDisplayT) LvSdlWindowSetResizeable(value bool) {
}

// llgo:link (*LvDisplayT).LvSdlWindowSetZoom C.lv_sdl_window_set_zoom
func (recv_ *LvDisplayT) LvSdlWindowSetZoom(zoom c.Float) {
}

// llgo:link (*LvDisplayT).LvSdlWindowGetZoom C.lv_sdl_window_get_zoom
func (recv_ *LvDisplayT) LvSdlWindowGetZoom() c.Float {
	return 0
}

// llgo:link (*LvDisplayT).LvSdlWindowSetTitle C.lv_sdl_window_set_title
func (recv_ *LvDisplayT) LvSdlWindowSetTitle(title *c.Char) {
}

// llgo:link (*LvDisplayT).LvSdlWindowSetIcon C.lv_sdl_window_set_icon
func (recv_ *LvDisplayT) LvSdlWindowSetIcon(icon c.Pointer, width c.Int32T, height c.Int32T) {
}

// llgo:link (*LvDisplayT).LvSdlWindowGetRenderer C.lv_sdl_window_get_renderer
func (recv_ *LvDisplayT) LvSdlWindowGetRenderer() c.Pointer {
	return nil
}

//go:linkname LvSdlQuit C.lv_sdl_quit
func LvSdlQuit()

/**********************
 * GLOBAL PROTOTYPES
 **********************/
//go:linkname LvSdlMouseCreate C.lv_sdl_mouse_create
func LvSdlMouseCreate() *LvIndevT

/**********************
 * GLOBAL PROTOTYPES
 **********************/
//go:linkname LvSdlMousewheelCreate C.lv_sdl_mousewheel_create
func LvSdlMousewheelCreate() *LvIndevT

/**********************
 * GLOBAL PROTOTYPES
 **********************/
//go:linkname LvSdlKeyboardCreate C.lv_sdl_keyboard_create
func LvSdlKeyboardCreate() *LvIndevT

type LvCoordT c.Int32T
type LvResT LvResultT
type LvImgDscT LvImageDscT
type LvDispT LvDisplayT
type LvDispRotationT LvDisplayRotationT
type LvDispRenderT LvDisplayRenderModeT
type LvAnimReadyCbT LvAnimCompletedCbT
type LvScrLoadAnimT LvScreenLoadAnimT
type LvBtnmatrixCtrlT LvButtonmatrixCtrlT
