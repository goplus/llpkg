package lvgl

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const OS_NONE = 0
const OS_PTHREAD = 1
const OS_FREERTOS = 2
const OS_CMSIS_RTOS2 = 3
const OS_RTTHREAD = 4
const OS_WINDOWS = 5
const OS_MQX = 6
const OS_SDL2 = 7
const OS_CUSTOM = 255
const STDLIB_BUILTIN = 0
const STDLIB_CLIB = 1
const STDLIB_MICROPYTHON = 2
const STDLIB_RTTHREAD = 3
const STDLIB_CUSTOM = 255
const DRAW_SW_ASM_NONE = 0
const DRAW_SW_ASM_NEON = 1
const DRAW_SW_ASM_HELIUM = 2
const DRAW_SW_ASM_CUSTOM = 255
const NEMA_HAL_CUSTOM = 0
const NEMA_HAL_STM32 = 1
const LOG_TRACE_MEM = 0
const LOG_TRACE_TIMER = 0
const LOG_TRACE_INDEV = 0
const LOG_TRACE_DISP_REFR = 0
const LOG_TRACE_EVENT = 0
const LOG_TRACE_OBJ_CREATE = 0
const LOG_TRACE_LAYOUT = 0
const LOG_TRACE_ANIM = 0
const WAYLAND_USE_DMABUF = 0
const WAYLAND_WINDOW_DECORATIONS = 0
const WAYLAND_WL_SHELL = 0
const USE_PERF_MONITOR = 0
const USE_MEM_MONITOR = 0
const USE_PERF_MONITOR_LOG_MODE = 0
const USE_LZ4 = 0
const USE_THORVG = 0
const SYMBOL_BULLET = "•"
const SYMBOL_AUDIO = "\uf001"
const SYMBOL_VIDEO = "\uf008"
const SYMBOL_LIST = "\uf00b"
const SYMBOL_OK = "\uf00c"
const SYMBOL_CLOSE = "\uf00d"
const SYMBOL_POWER = "\uf011"
const SYMBOL_SETTINGS = "\uf013"
const SYMBOL_HOME = "\uf015"
const SYMBOL_DOWNLOAD = "\uf019"
const SYMBOL_DRIVE = "\uf01c"
const SYMBOL_REFRESH = "\uf021"
const SYMBOL_MUTE = "\uf026"
const SYMBOL_VOLUME_MID = "\uf027"
const SYMBOL_VOLUME_MAX = "\uf028"
const SYMBOL_IMAGE = "\uf03e"
const SYMBOL_TINT = "\uf043"
const SYMBOL_PREV = "\uf048"
const SYMBOL_PLAY = "\uf04b"
const SYMBOL_PAUSE = "\uf04c"
const SYMBOL_STOP = "\uf04d"
const SYMBOL_NEXT = "\uf051"
const SYMBOL_EJECT = "\uf052"
const SYMBOL_LEFT = "\uf053"
const SYMBOL_RIGHT = "\uf054"
const SYMBOL_PLUS = "\uf067"
const SYMBOL_MINUS = "\uf068"
const SYMBOL_EYE_OPEN = "\uf06e"
const SYMBOL_EYE_CLOSE = "\uf070"
const SYMBOL_WARNING = "\uf071"
const SYMBOL_SHUFFLE = "\uf074"
const SYMBOL_UP = "\uf077"
const SYMBOL_DOWN = "\uf078"
const SYMBOL_LOOP = "\uf079"
const SYMBOL_DIRECTORY = "\uf07b"
const SYMBOL_UPLOAD = "\uf093"
const SYMBOL_CALL = "\uf095"
const SYMBOL_CUT = "\uf0c4"
const SYMBOL_COPY = "\uf0c5"
const SYMBOL_SAVE = "\uf0c7"
const SYMBOL_BARS = "\uf0c9"
const SYMBOL_ENVELOPE = "\uf0e0"
const SYMBOL_CHARGE = "\uf0e7"
const SYMBOL_PASTE = "\uf0ea"
const SYMBOL_BELL = "\uf0f3"
const SYMBOL_KEYBOARD = "\uf11c"
const SYMBOL_GPS = "\uf124"
const SYMBOL_FILE = "\uf15b"
const SYMBOL_WIFI = "\uf1eb"
const SYMBOL_BATTERY_FULL = "\uf240"
const SYMBOL_BATTERY_3 = "\uf241"
const SYMBOL_BATTERY_2 = "\uf242"
const SYMBOL_BATTERY_1 = "\uf243"
const SYMBOL_BATTERY_EMPTY = "\uf244"
const SYMBOL_USB = "\uf287"
const SYMBOL_BLUETOOTH = "\uf293"
const SYMBOL_TRASH = "\uf2ed"
const SYMBOL_EDIT = "\uf304"
const SYMBOL_BACKSPACE = "\uf55a"
const SYMBOL_SD_CARD = "\uf7c2"
const SYMBOL_NEW_LINE = "\uf8a2"
const SYMBOL_DUMMY = "\uf8ff"
const TRIGO_SIN_MAX = 32768
const TRIGO_SHIFT = 15
const BEZIER_VAL_SHIFT = 10
const LOG_LEVEL_TRACE = 0
const LOG_LEVEL_INFO = 1
const LOG_LEVEL_WARN = 2
const LOG_LEVEL_ERROR = 3
const LOG_LEVEL_USER = 4
const LOG_LEVEL_NONE = 5
const LOG_LEVEL_NUM = 5
const COLOR_NATIVE_WITH_ALPHA_SIZE = 3
const OPA_MIN = 2
const OPA_MAX = 253
const STRIDE_AUTO = 0
const NO_TIMER_READY = 0xFFFFFFFF
const ANIM_REPEAT_INFINITE = 0xFFFFFFFF
const ANIM_PLAYTIME_INFINITE = 0xFFFFFFFF
const ANIM_PAUSE_FOREVER = 0xFFFFFFFF
const TXT_ENC_UTF8 = 1
const TXT_ENC_ASCII = 2
const BIDI_LRO = "\u202d"
const BIDI_RLO = "\u202e"
const STYLE_SENTINEL_VALUE = 0xAABBCCDD
const SCALE_NONE = 256
const FS_MAX_FN_LENGTH = 64
const FS_MAX_PATH_LENGTH = 256
const DRAW_UNIT_NONE = 0
const ARRAY_DEFAULT_CAPACITY = 4
const ARRAY_DEFAULT_SHRINK_RATIO = 2
const RADIUS_CIRCLE = 0x7FFF
const MASK_MAX_NUM = 16
const ZERO_MEM_SENTINEL = 0xa1b2c3d4
const INV_BUF_SIZE = 32
const INDEV_VECT_HIST_SIZE = 8
const BUTTONMATRIX_BUTTON_NONE = 0xFFFF
const LABEL_DOT_NUM = 3
const LABEL_POS_LAST = 0xFFFF
const LABEL_DEFAULT_TEXT = "Text"
const SWITCH_KNOB_EXT_AREA_CORRECTION = 2
const TABLE_CELL_NONE = 0xFFFF
const DROPDOWN_POS_LAST = 0xFFFF
const SCALE_LABEL_ROTATE_MATCH_TICKS = 0x100000
const SCALE_LABEL_ROTATE_KEEP_UPRIGHT = 0x80000
const SCALE_ROTATION_ANGLE_MASK = 0x7FFFF
const LED_BRIGHT_MIN = 80
const LED_BRIGHT_MAX = 255
const SPINBOX_MAX_DIGIT_COUNT = 10
const ANIM_TIMELINE_PROGRESS_MAX = 0xFFFF
const ARCLABEL_DOT_NUM = 3
const ARCLABEL_DEFAULT_TEXT = "Arced Text"
const FS_MAX_PATH_LEN = 256
const SDL_MOUSEWHEEL_MODE_ENCODER = 0
const SDL_MOUSEWHEEL_MODE_CROWN = 1
const KEYBOARD_BUFFER_SIZE = 32

type X_silenceGccWarning struct {
	Unused [8]uint8
}
type ResultT c.Int

const (
	RESULT_INVALID ResultT = 0
	RESULT_OK      ResultT = 1
)

type UintptrT c.UintptrT
type IntptrT c.IntptrT
type ValuePreciseT c.Int32T

type X_lvObjT struct {
	ClassP                    *ObjClassT
	Parent                    *ObjT
	SpecAttr                  *ObjSpecAttrT
	Styles                    *ObjStyleT
	UserData                  c.Pointer
	Coords                    AreaT
	Flags                     ObjFlagT
	State                     StateT
	LayoutInv                 c.Uint16T
	ReadjustScrollAfterLayout c.Uint16T
	ScrLayoutInv              c.Uint16T
	SkipTrans                 c.Uint16T
	StyleCnt                  c.Uint16T
	HLayout                   c.Uint16T
	WLayout                   c.Uint16T
	IsDeleting                c.Uint16T
}
type ObjT X_lvObjT

// llgo:type C
type ScreenCreateCbT func() *ObjT
type StateT c.Uint16T
type PartT c.Uint32T
type OpaT c.Uint8T
type StylePropT c.Uint8T

type X_lvObjClassT struct {
	BaseClass        *ObjClassT
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
type ObjClassT X_lvObjClassT

type X_lvGroupT struct {
	ObjLl         LlT
	ObjFocus      **ObjT
	FocusCb       GroupFocusCbT
	EdgeCb        GroupEdgeCbT
	UserData      c.Pointer
	Frozen        c.Uint8T
	Editing       c.Uint8T
	RefocusPolicy c.Uint8T
	Wrap          c.Uint8T
}
type GroupT X_lvGroupT

type X_lvDisplayT struct {
	HorRes              c.Int32T
	VerRes              c.Int32T
	PhysicalHorRes      c.Int32T
	PhysicalVerRes      c.Int32T
	OffsetX             c.Int32T
	OffsetY             c.Int32T
	Dpi                 c.Uint32T
	Buf1                *DrawBufT
	Buf2                *DrawBufT
	Buf3                *DrawBufT
	BufAct              *DrawBufT
	FlushCb             DisplayFlushCbT
	FlushWaitCb         DisplayFlushWaitCbT
	Flushing            c.Int
	FlushingLast        c.Int
	LastArea            c.Uint32T
	LastPart            c.Uint32T
	RenderMode          DisplayRenderModeT
	Antialiasing        c.Uint32T
	TileCnt             c.Uint32T
	StrideIsAuto        c.Uint32T
	RenderingInProgress c.Uint32T
	ColorFormat         ColorFormatT
	InvAreas            [32]AreaT
	InvAreaJoined       [32]c.Uint8T
	InvP                c.Uint32T
	InvEnCnt            c.Int32T
	SyncAreas           LlT
	X_staticBuf1        DrawBufT
	X_staticBuf2        DrawBufT
	LayerHead           *LayerT
	LayerInit           c.Pointer
	LayerDeinit         c.Pointer
	Screens             **ObjT
	SysLayer            *ObjT
	TopLayer            *ObjT
	ActScr              *ObjT
	BottomLayer         *ObjT
	PrevScr             *ObjT
	ScrToLoad           *ObjT
	ScreenCnt           c.Uint32T
	DrawPrevOverAct     c.Uint8T
	DelPrev             c.Uint8T
	DriverData          c.Pointer
	UserData            c.Pointer
	EventList           EventListT
	Rotation            c.Uint32T
	MatrixRotation      c.Uint32T
	Theme               *ThemeT
	RefrTimer           *TimerT
	LastActivityTime    c.Uint32T
	RefreshedArea       AreaT
	VsyncCount          c.Uint32T
}
type DisplayT X_lvDisplayT

type X_lvLayerT struct {
	DrawBuf        *DrawBufT
	DrawTaskHead   *DrawTaskT
	Parent         *LayerT
	Next           *LayerT
	UserData       c.Pointer
	BufArea        AreaT
	PhyClipArea    AreaT
	X_clipArea     AreaT
	PartialYOffset c.Int32T
	Recolor        Color32T
	ColorFormat    ColorFormatT
	AllTasksAdded  bool
	Opa            OpaT
}
type LayerT X_lvLayerT

type X_lvDrawUnitT struct {
	Next            *DrawUnitT
	Name            *c.Char
	Idx             c.Int32T
	DispatchCb      c.Pointer
	EvaluateCb      c.Pointer
	WaitForFinishCb c.Pointer
	DeleteCb        c.Pointer
}
type DrawUnitT X_lvDrawUnitT

type X_lvDrawTaskT struct {
	Next                *DrawTaskT
	Type                DrawTaskTypeT
	Area                AreaT
	X_realArea          AreaT
	ClipAreaOriginal    AreaT
	ClipArea            AreaT
	TargetLayer         *LayerT
	DrawUnit            *DrawUnitT
	State               c.Int
	DrawDsc             c.Pointer
	Opa                 OpaT
	PreferredDrawUnitId c.Uint8T
	PreferenceScore     c.Uint8T
}
type DrawTaskT X_lvDrawTaskT

type X_lvIndevT struct {
	Type                IndevTypeT
	ReadCb              IndevReadCbT
	State               IndevStateT
	PrevState           IndevStateT
	Mode                IndevModeT
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
	Disp                *DisplayT
	ReadTimer           *TimerT
	ScrollLimit         c.Uint8T
	ScrollThrow         c.Uint8T
	GestureMinVelocity  c.Uint8T
	GestureLimit        c.Uint8T
	LongPressTime       c.Uint16T
	LongPressRepeatTime c.Uint16T
	RotarySensitivity   c.Int32T
	Pointer             struct {
		ActPoint                PointT
		LastPoint               PointT
		LastRawPoint            PointT
		Vect                    PointT
		VectHist                [8]PointT
		VectHistTimestamp       [8]c.Uint32T
		VectHistIndex           c.Uint8T
		ScrollSum               PointT
		ScrollThrowVect         PointT
		ScrollThrowVectOri      PointT
		ActObj                  *ObjT
		LastObj                 *ObjT
		ScrollObj               *ObjT
		LastPressed             *ObjT
		LastHovered             *ObjT
		ScrollArea              AreaT
		GestureSum              PointT
		Diff                    c.Int32T
		ShortClickStreak        c.Uint8T
		LastShortClickPoint     PointT
		LastShortClickTimestamp c.Uint32T
		ScrollDir               c.Uint8T
		GestureDir              c.Uint8T
		GestureSent             c.Uint8T
		PressMoved              c.Uint8T
		Pressed                 c.Uint8T
	}
	Keypad struct {
		LastState IndevStateT
		LastKey   c.Uint32T
	}
	Cursor          *ObjT
	Group           *GroupT
	BtnPoints       *PointT
	EventList       EventListT
	ScrollThrowAnim *AnimT
}
type IndevT X_lvIndevT

type X_lvEventT struct {
	CurrentTarget  c.Pointer
	OriginalTarget c.Pointer
	Code           EventCodeT
	UserData       c.Pointer
	Param          c.Pointer
	Prev           *EventT
	Deleted        c.Uint8T
	StopProcessing c.Uint8T
	StopBubbling   c.Uint8T
}
type EventT X_lvEventT

type X_lvTimerT struct {
	Period      c.Uint32T
	LastRun     c.Uint32T
	TimerCb     TimerCbT
	UserData    c.Pointer
	RepeatCount c.Int32T
	Paused      c.Uint32T
	AutoDelete  c.Uint32T
}
type TimerT X_lvTimerT

type X_lvThemeT struct {
	ApplyCb        ThemeApplyCbT
	Parent         *ThemeT
	UserData       c.Pointer
	Disp           *DisplayT
	ColorPrimary   ColorT
	ColorSecondary ColorT
	FontSmall      *FontT
	FontNormal     *FontT
	FontLarge      *FontT
	Flags          c.Uint32T
}
type ThemeT X_lvThemeT

type X_lvAnimT struct {
	Var                   c.Pointer
	ExecCb                AnimExecXcbT
	CustomExecCb          AnimCustomExecCbT
	StartCb               AnimStartCbT
	CompletedCb           AnimCompletedCbT
	DeletedCb             AnimDeletedCbT
	GetValueCb            AnimGetValueCbT
	UserData              c.Pointer
	PathCb                AnimPathCbT
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
type AnimT X_lvAnimT

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
	Fallback           *FontT
	UserData           c.Pointer
}
type FontT X_lvFontT

type X_lvFontClassT struct {
	CreateCb  c.Pointer
	DeleteCb  c.Pointer
	DupSrcCb  c.Pointer
	FreeSrcCb c.Pointer
}
type FontClassT X_lvFontClassT

type X_lvFontInfoT struct {
	Name       *c.Char
	ClassP     *FontClassT
	Size       c.Uint32T
	RenderMode c.Uint32T
	Style      c.Uint32T
	Kerning    FontKerningT
}
type FontInfoT X_lvFontInfoT

type X_lvFontManagerT struct {
	Unused [8]uint8
}
type FontManagerT X_lvFontManagerT

type X_lvImageDecoderT struct {
	InfoCb       ImageDecoderInfoFT
	OpenCb       ImageDecoderOpenFT
	GetAreaCb    ImageDecoderGetAreaCbT
	CloseCb      ImageDecoderCloseFT
	CustomDrawCb ImageDecoderCustomDrawT
	Name         *c.Char
	UserData     c.Pointer
}
type ImageDecoderT X_lvImageDecoderT

type X_lvImageDecoderDscT struct {
	Decoder     *ImageDecoderT
	Args        ImageDecoderArgsT
	Src         c.Pointer
	SrcType     ImageSrcT
	File        FsFileT
	Header      ImageHeaderT
	Decoded     *DrawBufT
	Palette     *Color32T
	PaletteSize c.Uint32T
	TimeToOpen  c.Uint32T
	ErrorMsg    *c.Char
	Cache       *CacheT
	CacheEntry  *CacheEntryT
	UserData    c.Pointer
}
type ImageDecoderDscT X_lvImageDecoderDscT

type X_lvDrawImageDscT struct {
	Base          DrawDscBaseT
	Src           c.Pointer
	Header        ImageHeaderT
	ClipRadius    c.Int32T
	Rotation      c.Int32T
	ScaleX        c.Int32T
	ScaleY        c.Int32T
	SkewX         c.Int32T
	SkewY         c.Int32T
	Pivot         PointT
	Recolor       ColorT
	RecolorOpa    OpaT
	Opa           OpaT
	BlendMode     BlendModeT
	Antialias     c.Uint16T
	Tile          c.Uint16T
	Sup           *DrawImageSupT
	ImageArea     AreaT
	BitmapMaskSrc *ImageDscT
}
type DrawImageDscT X_lvDrawImageDscT

type X_lvFragmentT struct {
	Unused [8]uint8
}
type FragmentT X_lvFragmentT

type X_lvFragmentClassT struct {
	Unused [8]uint8
}
type FragmentClassT X_lvFragmentClassT

type X_lvFragmentManagedStatesT struct {
	Unused [8]uint8
}
type FragmentManagedStatesT X_lvFragmentManagedStatesT

type X_lvProfilerBuiltinConfigT struct {
	Unused [8]uint8
}
type ProfilerBuiltinConfigT X_lvProfilerBuiltinConfigT

type X_lvRbNodeT struct {
	Parent *X_lvRbNodeT
	Left   *X_lvRbNodeT
	Right  *X_lvRbNodeT
	Color  RbColorT
	Data   c.Pointer
}
type RbNodeT X_lvRbNodeT

type X_lvRbT struct {
	Root    *RbNodeT
	Compare RbCompareT
	Size    c.SizeT
}
type RbT X_lvRbT

type X_lvColorFilterDscT struct {
	FilterCb ColorFilterCbT
	UserData c.Pointer
}
type ColorFilterDscT X_lvColorFilterDscT

type X_lvEventDscT struct {
	Cb       EventCbT
	UserData c.Pointer
	Filter   c.Uint32T
}
type EventDscT X_lvEventDscT

type X_lvCacheT struct {
	Clz      *CacheClassT
	NodeSize c.Uint32T
	MaxSize  c.Uint32T
	Size     c.Uint32T
	Ops      CacheOpsT
	Lock     MutexT
	Name     *c.Char
}
type CacheT X_lvCacheT

type X_lvCacheEntryT struct {
	Unused [8]uint8
}
type CacheEntryT X_lvCacheEntryT

type X_lvFsFileCacheT struct {
	Start        c.Uint32T
	End          c.Uint32T
	FilePosition c.Uint32T
	Buffer       c.Pointer
}
type FsFileCacheT X_lvFsFileCacheT

type X_lvFsPathExT struct {
	Path   [4]c.Char
	Buffer c.Pointer
	Size   c.Uint32T
}
type FsPathExT X_lvFsPathExT

type X_lvImageDecoderArgsT struct {
	StrideAlign bool
	Premultiply bool
	NoCache     bool
	UseIndexed  bool
	FlushCache  bool
}
type ImageDecoderArgsT X_lvImageDecoderArgsT

type X_lvImageCacheDataT struct {
	Slot     CacheSlotSizeT
	Src      c.Pointer
	SrcType  ImageSrcT
	Decoded  *DrawBufT
	Decoder  *ImageDecoderT
	UserData c.Pointer
}
type ImageCacheDataT X_lvImageCacheDataT

type X_lvImageHeaderCacheDataT struct {
	Src     c.Pointer
	SrcType ImageSrcT
	Header  ImageHeaderT
	Decoder *ImageDecoderT
}
type ImageHeaderCacheDataT X_lvImageHeaderCacheDataT

type X_lvDrawMaskT struct {
	UserData c.Pointer
}
type DrawMaskT X_lvDrawMaskT

type X_lvDrawLabelHintT struct {
	LineStart c.Int32T
	Y         c.Int32T
	CoordY    c.Int32T
}
type DrawLabelHintT X_lvDrawLabelHintT

type X_lvDrawGlyphDscT struct {
	GlyphData          c.Pointer
	Format             FontGlyphFormatT
	LetterCoords       *AreaT
	BgCoords           *AreaT
	G                  *FontGlyphDscT
	Color              ColorT
	Opa                OpaT
	OutlineStrokeColor ColorT
	OutlineStrokeOpa   OpaT
	OutlineStrokeWidth c.Int32T
	Rotation           c.Int32T
	Pivot              PointT
	X_drawBuf          *DrawBufT
}
type DrawGlyphDscT X_lvDrawGlyphDscT

type X_lvDrawImageSupT struct {
	AlphaColor  ColorT
	Palette     *Color32T
	PaletteSize c.Uint32T
}
type DrawImageSupT X_lvDrawImageSupT

type X_lvDrawMaskRectDscT struct {
	Base        DrawDscBaseT
	Area        AreaT
	Radius      c.Int32T
	KeepOutside c.Uint32T
}
type DrawMaskRectDscT X_lvDrawMaskRectDscT

type X_lvObjStyleT struct {
	Style      *StyleT
	Selector   c.Uint32T
	IsLocal    c.Uint32T
	IsTrans    c.Uint32T
	IsDisabled c.Uint32T
}
type ObjStyleT X_lvObjStyleT

type X_lvObjStyleTransitionDscT struct {
	Time     c.Uint16T
	Delay    c.Uint16T
	Selector StyleSelectorT
	Prop     StylePropT
	PathCb   AnimPathCbT
	UserData c.Pointer
}
type ObjStyleTransitionDscT X_lvObjStyleTransitionDscT

type X_lvHitTestInfoT struct {
	Point *PointT
	Res   bool
}
type HitTestInfoT X_lvHitTestInfoT

type X_lvCoverCheckInfoT struct {
	Res  CoverResT
	Area *AreaT
}
type CoverCheckInfoT X_lvCoverCheckInfoT

type X_lvObjSpecAttrT struct {
	Children      **ObjT
	GroupP        *GroupT
	EventList     EventListT
	Scroll        PointT
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
type ObjSpecAttrT X_lvObjSpecAttrT

type X_lvImageT struct {
	Obj           ObjT
	Src           c.Pointer
	BitmapMaskSrc *ImageDscT
	Offset        PointT
	W             c.Int32T
	H             c.Int32T
	Rotation      c.Uint32T
	ScaleX        c.Uint32T
	ScaleY        c.Uint32T
	Pivot         PointT
	SrcType       c.Uint32T
	Cf            c.Uint32T
	Antialias     c.Uint32T
	Align         c.Uint32T
	BlendMode     c.Uint32T
}
type ImageT X_lvImageT

type X_lvAnimimgT struct {
	Img      ImageT
	Anim     AnimT
	Dsc      *c.Pointer
	PicCount c.Int8T
}
type AnimimgT X_lvAnimimgT

type X_lvArcT struct {
	Obj             ObjT
	Rotation        c.Int32T
	IndicAngleStart ValuePreciseT
	IndicAngleEnd   ValuePreciseT
	BgAngleStart    ValuePreciseT
	BgAngleEnd      ValuePreciseT
	Value           c.Int32T
	MinValue        c.Int32T
	MaxValue        c.Int32T
	Dragging        c.Uint32T
	Type            c.Uint32T
	MinClose        c.Uint32T
	InOut           c.Uint32T
	ChgRate         c.Uint32T
	LastTick        c.Uint32T
	LastAngle       ValuePreciseT
	KnobOffset      c.Int16T
}
type ArcT X_lvArcT

type X_lvArclabelT struct {
	Unused [8]uint8
}
type ArclabelT X_lvArclabelT

type X_lvLabelT struct {
	Obj              ObjT
	Text             *c.Char
	Dot              [4]c.Char
	DotBegin         c.Uint32T
	Hint             DrawLabelHintT
	SelStart         c.Uint32T
	SelEnd           c.Uint32T
	SizeCache        PointT
	Offset           PointT
	LongMode         LabelLongModeT
	StaticTxt        c.Uint8T
	Recolor          c.Uint8T
	Expand           c.Uint8T
	InvalidSizeCache c.Uint8T
	TextSize         PointT
}
type LabelT X_lvLabelT

type X_lvBarAnimT struct {
	Bar       *ObjT
	AnimStart c.Int32T
	AnimEnd   c.Int32T
	AnimState c.Int32T
}
type BarAnimT X_lvBarAnimT

type X_lvBarT struct {
	Obj            ObjT
	CurValue       c.Int32T
	MinValue       c.Int32T
	MaxValue       c.Int32T
	StartValue     c.Int32T
	IndicArea      AreaT
	ValReversed    bool
	CurValueAnim   BarAnimT
	StartValueAnim BarAnimT
	Mode           BarModeT
	Orientation    BarOrientationT
}
type BarT X_lvBarT

type X_lvButtonT struct {
	Obj ObjT
}
type ButtonT X_lvButtonT

type X_lvButtonmatrixT struct {
	Obj         ObjT
	MapP        **c.Char
	ButtonAreas *AreaT
	CtrlBits    *ButtonmatrixCtrlT
	BtnCnt      c.Uint32T
	RowCnt      c.Uint32T
	BtnIdSel    c.Uint32T
	OneCheck    c.Uint32T
	AutoFreeMap c.Uint32T
}
type ButtonmatrixT X_lvButtonmatrixT

type X_lvCalendarT struct {
	Obj                 ObjT
	Btnm                *ObjT
	Today               CalendarDateT
	ShowedDate          CalendarDateT
	HighlightedDates    *CalendarDateT
	HighlightedDatesNum c.SizeT
	Map                 [56]*c.Char
	UseChineseCalendar  bool
	Nums                [42][20]c.Char
}
type CalendarT X_lvCalendarT

type X_lvCanvasT struct {
	Img       ImageT
	DrawBuf   *DrawBufT
	StaticBuf DrawBufT
}
type CanvasT X_lvCanvasT

type X_lvChartSeriesT struct {
	XPoints         *c.Int32T
	YPoints         *c.Int32T
	Color           ColorT
	StartPoint      c.Uint32T
	Hidden          c.Uint32T
	XExtBufAssigned c.Uint32T
	YExtBufAssigned c.Uint32T
	XAxisSec        c.Uint32T
	YAxisSec        c.Uint32T
}
type ChartSeriesT X_lvChartSeriesT

type X_lvChartCursorT struct {
	Pos     PointT
	PointId c.Int32T
	Color   ColorT
	Ser     *ChartSeriesT
	Dir     DirT
	PosSet  c.Uint32T
}
type ChartCursorT X_lvChartCursorT

type X_lvChartT struct {
	Obj            ObjT
	SeriesLl       LlT
	CursorLl       LlT
	Ymin           [2]c.Int32T
	Ymax           [2]c.Int32T
	Xmin           [2]c.Int32T
	Xmax           [2]c.Int32T
	PressedPointId c.Int32T
	HdivCnt        c.Uint32T
	VdivCnt        c.Uint32T
	PointCnt       c.Uint32T
	Type           ChartTypeT
	UpdateMode     ChartUpdateModeT
}
type ChartT X_lvChartT

type X_lvCheckboxT struct {
	Obj       ObjT
	Txt       *c.Char
	StaticTxt c.Uint32T
}
type CheckboxT X_lvCheckboxT

type X_lvDropdownT struct {
	Obj               ObjT
	List              *ObjT
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
type DropdownT X_lvDropdownT

type X_lvDropdownListT struct {
	Obj      ObjT
	Dropdown *ObjT
}
type DropdownListT X_lvDropdownListT

type X_lvImagebuttonSrcInfoT struct {
	ImgSrc c.Pointer
	Header ImageHeaderT
}
type ImagebuttonSrcInfoT X_lvImagebuttonSrcInfoT

type X_lvImagebuttonT struct {
	Obj      ObjT
	SrcMid   [6]ImagebuttonSrcInfoT
	SrcLeft  [6]ImagebuttonSrcInfoT
	SrcRight [6]ImagebuttonSrcInfoT
}
type ImagebuttonT X_lvImagebuttonT

type X_lvKeyboardT struct {
	Btnm     ButtonmatrixT
	Ta       *ObjT
	Mode     KeyboardModeT
	Popovers c.Uint8T
}
type KeyboardT X_lvKeyboardT

type X_lvLedT struct {
	Obj    ObjT
	Color  ColorT
	Bright c.Uint8T
}
type LedT X_lvLedT

type X_lvLineT struct {
	Obj        ObjT
	PointArray struct {
		Constant *PointPreciseT
	}
	PointNum            c.Uint32T
	YInv                c.Uint32T
	PointArrayIsMutable c.Uint32T
}
type LineT X_lvLineT

type X_lvMenuLoadPageEventDataT struct {
	Menu *ObjT
	Page *ObjT
}
type MenuLoadPageEventDataT X_lvMenuLoadPageEventDataT

type X_lvMenuHistoryT struct {
	Page *ObjT
}
type MenuHistoryT X_lvMenuHistoryT

type X_lvMenuT struct {
	Obj                  ObjT
	Storage              *ObjT
	Main                 *ObjT
	MainPage             *ObjT
	MainHeader           *ObjT
	MainHeaderBackBtn    *ObjT
	MainHeaderTitle      *ObjT
	Sidebar              *ObjT
	SidebarPage          *ObjT
	SidebarHeader        *ObjT
	SidebarHeaderBackBtn *ObjT
	SidebarHeaderTitle   *ObjT
	SelectedTab          *ObjT
	HistoryLl            LlT
	CurDepth             c.Uint8T
	PrevDepth            c.Uint8T
	SidebarGenerated     c.Uint8T
	ModeHeader           MenuModeHeaderT
	ModeRootBackBtn      MenuModeRootBackButtonT
}
type MenuT X_lvMenuT

type X_lvMenuPageT struct {
	Obj         ObjT
	Title       *c.Char
	StaticTitle bool
}
type MenuPageT X_lvMenuPageT

type X_lvMsgboxT struct {
	Obj     ObjT
	Header  *ObjT
	Content *ObjT
	Footer  *ObjT
	Title   *ObjT
}
type MsgboxT X_lvMsgboxT

type X_lvRollerT struct {
	Obj         ObjT
	OptionCnt   c.Uint32T
	SelOptId    c.Uint32T
	SelOptIdOri c.Uint32T
	InfPageCnt  c.Uint32T
	Mode        RollerModeT
	Moved       c.Uint32T
}
type RollerT X_lvRollerT

type X_lvScaleSectionT struct {
	MainStyle               *StyleT
	IndicatorStyle          *StyleT
	ItemsStyle              *StyleT
	RangeMin                c.Int32T
	RangeMax                c.Int32T
	FirstTickIdxInSection   c.Uint32T
	LastTickIdxInSection    c.Uint32T
	FirstTickInSectionWidth c.Int32T
	LastTickInSectionWidth  c.Int32T
	FirstTickInSection      PointT
	LastTickInSection       PointT
	FirstTickIdxIsMajor     c.Uint32T
	LastTickIdxIsMajor      c.Uint32T
}
type ScaleSectionT X_lvScaleSectionT

type X_lvScaleT struct {
	Obj            ObjT
	SectionLl      LlT
	TxtSrc         **c.Char
	Mode           ScaleModeT
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
type ScaleT X_lvScaleT

type X_lvSliderT struct {
	Bar           BarT
	LeftKnobArea  AreaT
	RightKnobArea AreaT
	PressedPoint  PointT
	ValueToSet    *c.Int32T
	Dragging      c.Uint8T
	LeftKnobFocus c.Uint8T
}
type SliderT X_lvSliderT

type X_lvSpanT struct {
	Txt            *c.Char
	Style          StyleT
	StaticFlag     c.Uint32T
	TrailingPos    PointT
	TrailingHeight c.Int32T
}
type SpanT X_lvSpanT

type X_lvSpangroupT struct {
	Obj      ObjT
	Lines    c.Int32T
	Indent   c.Int32T
	CacheW   c.Int32T
	CacheH   c.Int32T
	ChildLl  LlT
	Overflow c.Uint32T
	Refresh  c.Uint32T
}
type SpangroupT X_lvSpangroupT

type X_lvTextareaT struct {
	Obj            ObjT
	Label          *ObjT
	PlaceholderTxt *c.Char
	PwdTmp         *c.Char
	PwdBullet      *c.Char
	AcceptedChars  *c.Char
	MaxLength      c.Uint32T
	PwdShowTime    c.Uint32T
	Cursor         struct {
		ValidX     c.Int32T
		Pos        c.Uint32T
		Area       AreaT
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
type TextareaT X_lvTextareaT

type X_lvSpinboxT struct {
	Ta           TextareaT
	Value        c.Int32T
	RangeMax     c.Int32T
	RangeMin     c.Int32T
	Step         c.Int32T
	DigitCount   c.Uint32T
	DecPointPos  c.Uint32T
	Rollover     c.Uint32T
	DigitStepDir c.Uint32T
}
type SpinboxT X_lvSpinboxT

type X_lvSwitchT struct {
	Obj         ObjT
	AnimState   c.Int32T
	Orientation SwitchOrientationT
}
type SwitchT X_lvSwitchT

type X_lvTableCellT struct {
	Ctrl     TableCellCtrlT
	UserData c.Pointer
	Txt      [1]c.Char
}
type TableCellT X_lvTableCellT

type X_lvTableT struct {
	Obj      ObjT
	ColCnt   c.Uint32T
	RowCnt   c.Uint32T
	CellData **TableCellT
	RowH     *c.Int32T
	ColW     *c.Int32T
	ColAct   c.Uint32T
	RowAct   c.Uint32T
}
type TableT X_lvTableT

type X_lvTabviewT struct {
	Obj    ObjT
	TabCur c.Uint32T
	TabPos DirT
}
type TabviewT X_lvTabviewT

type X_lvTileviewT struct {
	Obj     ObjT
	TileAct *ObjT
}
type TileviewT X_lvTileviewT

type X_lvTileviewTileT struct {
	Obj ObjT
	Dir DirT
}
type TileviewTileT X_lvTileviewTileT

type X_lvWinT struct {
	Obj ObjT
}
type WinT X_lvWinT

type X_lv3dtextureT struct {
	Unused [8]uint8
}
type X3dtextureT X_lv3dtextureT

type X_lvObserverT struct {
	Subject          *SubjectT
	Cb               ObserverCbT
	Target           c.Pointer
	UserData         c.Pointer
	AutoFreeUserData c.Uint32T
	Notified         c.Uint32T
	ForObj           c.Uint32T
}
type ObserverT X_lvObserverT

type X_lvMonkeyConfigT struct {
	Unused [8]uint8
}
type MonkeyConfigT X_lvMonkeyConfigT

type X_lvImePinyinT struct {
	Unused [8]uint8
}
type ImePinyinT X_lvImePinyinT

type X_lvFileExplorerT struct {
	Unused [8]uint8
}
type FileExplorerT X_lvFileExplorerT

type X_lvBarcodeT struct {
	Unused [8]uint8
}
type BarcodeT X_lvBarcodeT

type X_lvGifT struct {
	Unused [8]uint8
}
type GifT X_lvGifT

type X_lvQrcodeT struct {
	Unused [8]uint8
}
type QrcodeT X_lvQrcodeT

type X_lvFreetypeOutlineVectorT struct {
	Unused [8]uint8
}
type FreetypeOutlineVectorT X_lvFreetypeOutlineVectorT

type X_lvFreetypeOutlineEventParamT struct {
	Unused [8]uint8
}
type FreetypeOutlineEventParamT X_lvFreetypeOutlineEventParamT

type X_lvFpointT struct {
	Unused [8]uint8
}
type FpointT X_lvFpointT

type X_lvMatrixT struct {
	Unused [8]uint8
}
type MatrixT X_lvMatrixT

type X_lvVectorPathT struct {
	Unused [8]uint8
}
type VectorPathT X_lvVectorPathT

type X_lvVectorGradientT struct {
	Unused [8]uint8
}
type VectorGradientT X_lvVectorGradientT

type X_lvVectorFillDscT struct {
	Unused [8]uint8
}
type VectorFillDscT X_lvVectorFillDscT

type X_lvVectorStrokeDscT struct {
	Unused [8]uint8
}
type VectorStrokeDscT X_lvVectorStrokeDscT

type X_lvVectorDrawDscT struct {
	Unused [8]uint8
}
type VectorDrawDscT X_lvVectorDrawDscT

type X_lvDrawVectorTaskDscT struct {
	Unused [8]uint8
}
type DrawVectorTaskDscT X_lvDrawVectorTaskDscT

type X_lvVectorDscT struct {
	Unused [8]uint8
}
type VectorDscT X_lvVectorDscT

type X_lvXkbT struct {
	Unused [8]uint8
}
type XkbT X_lvXkbT

type X_lvLibinputEventT struct {
	Unused [8]uint8
}
type LibinputEventT X_lvLibinputEventT

type X_lvLibinputT struct {
	Unused [8]uint8
}
type LibinputT X_lvLibinputT

type X_lvDrawSwUnitT struct {
	BaseUnit DrawUnitT
	TaskAct  *DrawTaskT
}
type DrawSwUnitT X_lvDrawSwUnitT

type X_lvDrawSwMaskCommonDscT struct {
	Cb   DrawSwMaskXcbT
	Type DrawSwMaskTypeT
}
type DrawSwMaskCommonDscT X_lvDrawSwMaskCommonDscT

type X_lvDrawSwMaskLineParamT struct {
	Dsc DrawSwMaskCommonDscT
	Cfg struct {
		P1   PointT
		P2   PointT
		Side DrawSwMaskLineSideT
	}
	Origo   PointT
	XySteep c.Int32T
	YxSteep c.Int32T
	Steep   c.Int32T
	Spx     c.Int32T
	Flat    c.Uint8T
	Inv     c.Uint8T
}
type DrawSwMaskLineParamT X_lvDrawSwMaskLineParamT

type X_lvDrawSwMaskAngleParamT struct {
	Dsc DrawSwMaskCommonDscT
	Cfg struct {
		VertexP    PointT
		StartAngle c.Int32T
		EndAngle   c.Int32T
	}
	StartLine DrawSwMaskLineParamT
	EndLine   DrawSwMaskLineParamT
	DeltaDeg  c.Uint16T
}
type DrawSwMaskAngleParamT X_lvDrawSwMaskAngleParamT

type X_lvDrawSwMaskRadiusParamT struct {
	Dsc DrawSwMaskCommonDscT
	Cfg struct {
		Rect   AreaT
		Radius c.Int32T
		Outer  c.Uint8T
	}
	Circle *DrawSwMaskRadiusCircleDscT
}
type DrawSwMaskRadiusParamT X_lvDrawSwMaskRadiusParamT

type X_lvDrawSwMaskFadeParamT struct {
	Dsc DrawSwMaskCommonDscT
	Cfg struct {
		Coords    AreaT
		YTop      c.Int32T
		YBottom   c.Int32T
		OpaTop    OpaT
		OpaBottom OpaT
	}
}
type DrawSwMaskFadeParamT X_lvDrawSwMaskFadeParamT

type X_lvDrawSwMaskMapParamT struct {
	Dsc DrawSwMaskCommonDscT
	Cfg struct {
		Coords AreaT
		Map    *OpaT
	}
}
type DrawSwMaskMapParamT X_lvDrawSwMaskMapParamT

type X_lvDrawSwBlendDscT struct {
	BlendArea      *AreaT
	SrcBuf         c.Pointer
	SrcStride      c.Uint32T
	SrcColorFormat ColorFormatT
	SrcArea        *AreaT
	Opa            OpaT
	Color          ColorT
	MaskBuf        *OpaT
	MaskRes        DrawSwMaskResT
	MaskArea       *AreaT
	MaskStride     c.Int32T
	BlendMode      BlendModeT
}
type DrawSwBlendDscT X_lvDrawSwBlendDscT

type X_lvDrawSwBlendFillDscT struct {
	DestBuf      c.Pointer
	DestW        c.Int32T
	DestH        c.Int32T
	DestStride   c.Int32T
	MaskBuf      *OpaT
	MaskStride   c.Int32T
	Color        ColorT
	Opa          OpaT
	RelativeArea AreaT
}
type DrawSwBlendFillDscT X_lvDrawSwBlendFillDscT

type X_lvDrawSwBlendImageDscT struct {
	DestBuf        c.Pointer
	DestW          c.Int32T
	DestH          c.Int32T
	DestStride     c.Int32T
	MaskBuf        *OpaT
	MaskStride     c.Int32T
	SrcBuf         c.Pointer
	SrcStride      c.Int32T
	SrcColorFormat ColorFormatT
	Opa            OpaT
	BlendMode      BlendModeT
	RelativeArea   AreaT
	SrcArea        AreaT
}
type DrawSwBlendImageDscT X_lvDrawSwBlendImageDscT

type X_lvDrawBufHandlersT struct {
	BufMallocCb       DrawBufMallocCb
	BufFreeCb         DrawBufFreeCb
	AlignPointerCb    DrawBufAlignCb
	InvalidateCacheCb DrawBufCacheOperationCb
	FlushCacheCb      DrawBufCacheOperationCb
	WidthToStrideCb   DrawBufWidthToStrideCb
}
type DrawBufHandlersT X_lvDrawBufHandlersT

type X_lvRlottieT struct {
	Unused [8]uint8
}
type RlottieT X_lvRlottieT

type X_lvFfmpegPlayerT struct {
	Unused [8]uint8
}
type FfmpegPlayerT X_lvFfmpegPlayerT

type X_lvGlfwWindowT struct {
	Unused [8]uint8
}
type GlfwWindowT X_lvGlfwWindowT

type X_lvGlfwTextureT struct {
	Unused [8]uint8
}
type GlfwTextureT X_lvGlfwTextureT
type PropIdT c.Uint32T

type X_lvArrayT struct {
	Data        *c.Uint8T
	Size        c.Uint32T
	Capacity    c.Uint32T
	ElementSize c.Uint32T
	InnerAlloc  bool
}
type ArrayT X_lvArrayT

type X_lvIterT struct {
	Unused [8]uint8
}
type IterT X_lvIterT

type X_lvCircleBufT struct {
	Unused [8]uint8
}
type CircleBufT X_lvCircleBufT

type X_lvDrawBufT struct {
	Header        ImageHeaderT
	DataSize      c.Uint32T
	Data          *c.Uint8T
	UnalignedData c.Pointer
	Handlers      *DrawBufHandlersT
}
type DrawBufT X_lvDrawBufT

type X_lvXmlComponentScopeT struct {
	Unused [8]uint8
}
type XmlComponentScopeT X_lvXmlComponentScopeT

type X_lvXmlParserStateT struct {
	Unused [8]uint8
}
type XmlParserStateT X_lvXmlParserStateT
type X_lvStrSymbolIdT c.Int

const (
	STR_SYMBOL_BULLET        X_lvStrSymbolIdT = 0
	STR_SYMBOL_AUDIO         X_lvStrSymbolIdT = 1
	STR_SYMBOL_VIDEO         X_lvStrSymbolIdT = 2
	STR_SYMBOL_LIST          X_lvStrSymbolIdT = 3
	STR_SYMBOL_OK            X_lvStrSymbolIdT = 4
	STR_SYMBOL_CLOSE         X_lvStrSymbolIdT = 5
	STR_SYMBOL_POWER         X_lvStrSymbolIdT = 6
	STR_SYMBOL_SETTINGS      X_lvStrSymbolIdT = 7
	STR_SYMBOL_HOME          X_lvStrSymbolIdT = 8
	STR_SYMBOL_DOWNLOAD      X_lvStrSymbolIdT = 9
	STR_SYMBOL_DRIVE         X_lvStrSymbolIdT = 10
	STR_SYMBOL_REFRESH       X_lvStrSymbolIdT = 11
	STR_SYMBOL_MUTE          X_lvStrSymbolIdT = 12
	STR_SYMBOL_VOLUME_MID    X_lvStrSymbolIdT = 13
	STR_SYMBOL_VOLUME_MAX    X_lvStrSymbolIdT = 14
	STR_SYMBOL_IMAGE         X_lvStrSymbolIdT = 15
	STR_SYMBOL_TINT          X_lvStrSymbolIdT = 16
	STR_SYMBOL_PREV          X_lvStrSymbolIdT = 17
	STR_SYMBOL_PLAY          X_lvStrSymbolIdT = 18
	STR_SYMBOL_PAUSE         X_lvStrSymbolIdT = 19
	STR_SYMBOL_STOP          X_lvStrSymbolIdT = 20
	STR_SYMBOL_NEXT          X_lvStrSymbolIdT = 21
	STR_SYMBOL_EJECT         X_lvStrSymbolIdT = 22
	STR_SYMBOL_LEFT          X_lvStrSymbolIdT = 23
	STR_SYMBOL_RIGHT         X_lvStrSymbolIdT = 24
	STR_SYMBOL_PLUS          X_lvStrSymbolIdT = 25
	STR_SYMBOL_MINUS         X_lvStrSymbolIdT = 26
	STR_SYMBOL_EYE_OPEN      X_lvStrSymbolIdT = 27
	STR_SYMBOL_EYE_CLOSE     X_lvStrSymbolIdT = 28
	STR_SYMBOL_WARNING       X_lvStrSymbolIdT = 29
	STR_SYMBOL_SHUFFLE       X_lvStrSymbolIdT = 30
	STR_SYMBOL_UP            X_lvStrSymbolIdT = 31
	STR_SYMBOL_DOWN          X_lvStrSymbolIdT = 32
	STR_SYMBOL_LOOP          X_lvStrSymbolIdT = 33
	STR_SYMBOL_DIRECTORY     X_lvStrSymbolIdT = 34
	STR_SYMBOL_UPLOAD        X_lvStrSymbolIdT = 35
	STR_SYMBOL_CALL          X_lvStrSymbolIdT = 36
	STR_SYMBOL_CUT           X_lvStrSymbolIdT = 37
	STR_SYMBOL_COPY          X_lvStrSymbolIdT = 38
	STR_SYMBOL_SAVE          X_lvStrSymbolIdT = 39
	STR_SYMBOL_BARS          X_lvStrSymbolIdT = 40
	STR_SYMBOL_ENVELOPE      X_lvStrSymbolIdT = 41
	STR_SYMBOL_CHARGE        X_lvStrSymbolIdT = 42
	STR_SYMBOL_PASTE         X_lvStrSymbolIdT = 43
	STR_SYMBOL_BELL          X_lvStrSymbolIdT = 44
	STR_SYMBOL_KEYBOARD      X_lvStrSymbolIdT = 45
	STR_SYMBOL_GPS           X_lvStrSymbolIdT = 46
	STR_SYMBOL_FILE          X_lvStrSymbolIdT = 47
	STR_SYMBOL_WIFI          X_lvStrSymbolIdT = 48
	STR_SYMBOL_BATTERY_FULL  X_lvStrSymbolIdT = 49
	STR_SYMBOL_BATTERY_3     X_lvStrSymbolIdT = 50
	STR_SYMBOL_BATTERY_2     X_lvStrSymbolIdT = 51
	STR_SYMBOL_BATTERY_1     X_lvStrSymbolIdT = 52
	STR_SYMBOL_BATTERY_EMPTY X_lvStrSymbolIdT = 53
	STR_SYMBOL_USB           X_lvStrSymbolIdT = 54
	STR_SYMBOL_BLUETOOTH     X_lvStrSymbolIdT = 55
	STR_SYMBOL_TRASH         X_lvStrSymbolIdT = 56
	STR_SYMBOL_EDIT          X_lvStrSymbolIdT = 57
	STR_SYMBOL_BACKSPACE     X_lvStrSymbolIdT = 58
	STR_SYMBOL_SD_CARD       X_lvStrSymbolIdT = 59
	STR_SYMBOL_NEW_LINE      X_lvStrSymbolIdT = 60
	STR_SYMBOL_DUMMY         X_lvStrSymbolIdT = 61
)

/**********************
 *      TYPEDEFS
 **********************/

type SqrtResT struct {
	I c.Uint16T
	F c.Uint16T
}

//go:linkname TrigoSin C.lv_trigo_sin
func TrigoSin(angle c.Int16T) c.Int32T

//go:linkname TrigoCos C.lv_trigo_cos
func TrigoCos(angle c.Int16T) c.Int32T

/**
 * Calculate the y value of cubic-bezier(x1, y1, x2, y2) function as specified x.
 * @param x time in range of [0..LV_BEZIER_VAL_MAX]
 * @param x1 x of control point 1 in range of [0..LV_BEZIER_VAL_MAX]
 * @param y1 y of control point 1 in range of [0..LV_BEZIER_VAL_MAX]
 * @param x2 x of control point 2 in range of [0..LV_BEZIER_VAL_MAX]
 * @param y2 y of control point 2 in range of [0..LV_BEZIER_VAL_MAX]
 * @return the value calculated
 */
//go:linkname CubicBezier C.lv_cubic_bezier
func CubicBezier(x c.Int32T, x1 c.Int32T, y1 c.Int32T, x2 c.Int32T, y2 c.Int32T) c.Int32T

/**
 * Calculate a value of a Cubic Bezier function.
 * @param t time in range of [0..LV_BEZIER_VAL_MAX]
 * @param u0 must be 0
 * @param u1 control value 1 values in range of [0..LV_BEZIER_VAL_MAX]
 * @param u2 control value 2 in range of [0..LV_BEZIER_VAL_MAX]
 * @param u3 must be LV_BEZIER_VAL_MAX
 * @return the value calculated from the given parameters in range of [0..LV_BEZIER_VAL_MAX]
 */
//go:linkname Bezier3 C.lv_bezier3
func Bezier3(t c.Int32T, u0 c.Int32T, u1 c.Uint32T, u2 c.Int32T, u3 c.Int32T) c.Int32T

/**
 * Calculate the atan2 of a vector.
 * @param x
 * @param y
 * @return the angle in degree calculated from the given parameters in range of [0..360]
 */
//go:linkname Atan2 C.lv_atan2
func Atan2(x c.Int, y c.Int) c.Uint16T

//go:linkname Sqrt C.lv_sqrt
func Sqrt(x c.Uint32T, q *SqrtResT, mask c.Uint32T)

//go:linkname Sqrt32 C.lv_sqrt32
func Sqrt32(x c.Uint32T) c.Int32T

/**
 * Calculate the integer exponents.
 * @param base
 * @param exp
 * @return base raised to the power exponent
 */
//go:linkname Pow C.lv_pow
func Pow(base c.Int64T, exp c.Int8T) c.Int64T

/**
 * Get the mapped of a number given an input and output range
 * @param x integer which mapped value should be calculated
 * @param min_in min input range
 * @param max_in max input range
 * @param min_out max output range
 * @param max_out max output range
 * @return the mapped number
 */
//go:linkname Map C.lv_map
func Map(x c.Int32T, min_in c.Int32T, max_in c.Int32T, min_out c.Int32T, max_out c.Int32T) c.Int32T

/**
 * Set the seed of the pseudo random number generator
 * @param seed a number to initialize the random generator
 */
//go:linkname RandSetSeed C.lv_rand_set_seed
func RandSetSeed(seed c.Uint32T)

/**
 * Get a pseudo random number in the given range
 * @param min   the minimum value
 * @param max   the maximum value
 * @return return the random number. min <= return_value <= max
 */
//go:linkname Rand C.lv_rand
func Rand(min c.Uint32T, max c.Uint32T) c.Uint32T

/**
 * Represents a point on the screen.
 */

type PointT struct {
	X c.Int32T
	Y c.Int32T
}

type PointPreciseT struct {
	X ValuePreciseT
	Y ValuePreciseT
}

/** Represents an area of the screen.*/

type AreaT struct {
	X1 c.Int32T
	Y1 c.Int32T
	X2 c.Int32T
	Y2 c.Int32T
}
type AlignT c.Int

const (
	ALIGN_DEFAULT          AlignT = 0
	ALIGN_TOP_LEFT         AlignT = 1
	ALIGN_TOP_MID          AlignT = 2
	ALIGN_TOP_RIGHT        AlignT = 3
	ALIGN_BOTTOM_LEFT      AlignT = 4
	ALIGN_BOTTOM_MID       AlignT = 5
	ALIGN_BOTTOM_RIGHT     AlignT = 6
	ALIGN_LEFT_MID         AlignT = 7
	ALIGN_RIGHT_MID        AlignT = 8
	ALIGN_CENTER           AlignT = 9
	ALIGN_OUT_TOP_LEFT     AlignT = 10
	ALIGN_OUT_TOP_MID      AlignT = 11
	ALIGN_OUT_TOP_RIGHT    AlignT = 12
	ALIGN_OUT_BOTTOM_LEFT  AlignT = 13
	ALIGN_OUT_BOTTOM_MID   AlignT = 14
	ALIGN_OUT_BOTTOM_RIGHT AlignT = 15
	ALIGN_OUT_LEFT_TOP     AlignT = 16
	ALIGN_OUT_LEFT_MID     AlignT = 17
	ALIGN_OUT_LEFT_BOTTOM  AlignT = 18
	ALIGN_OUT_RIGHT_TOP    AlignT = 19
	ALIGN_OUT_RIGHT_MID    AlignT = 20
	ALIGN_OUT_RIGHT_BOTTOM AlignT = 21
)

type DirT c.Int

const (
	DIR_NONE   DirT = 0
	DIR_LEFT   DirT = 1
	DIR_RIGHT  DirT = 2
	DIR_TOP    DirT = 4
	DIR_BOTTOM DirT = 8
	DIR_HOR    DirT = 3
	DIR_VER    DirT = 12
	DIR_ALL    DirT = 15
)

/**
 * Initialize an area
 * @param area_p pointer to an area
 * @param x1 left coordinate of the area
 * @param y1 top coordinate of the area
 * @param x2 right coordinate of the area
 * @param y2 bottom coordinate of the area
 */
// llgo:link (*AreaT).AreaSet C.lv_area_set
func (recv_ *AreaT) AreaSet(x1 c.Int32T, y1 c.Int32T, x2 c.Int32T, y2 c.Int32T) {
}

/**
 * Get the width of an area
 * @param area_p pointer to an area
 * @return the width of the area (if x1 == x2 -> width = 1)
 */
// llgo:link (*AreaT).AreaGetWidth C.lv_area_get_width
func (recv_ *AreaT) AreaGetWidth() c.Int32T {
	return 0
}

/**
 * Get the height of an area
 * @param area_p pointer to an area
 * @return the height of the area (if y1 == y2 -> height = 1)
 */
// llgo:link (*AreaT).AreaGetHeight C.lv_area_get_height
func (recv_ *AreaT) AreaGetHeight() c.Int32T {
	return 0
}

/**
 * Set the width of an area
 * @param area_p pointer to an area
 * @param w the new width of the area (w == 1 makes x1 == x2)
 */
// llgo:link (*AreaT).AreaSetWidth C.lv_area_set_width
func (recv_ *AreaT) AreaSetWidth(w c.Int32T) {
}

/**
 * Set the height of an area
 * @param area_p pointer to an area
 * @param h the new height of the area (h == 1 makes y1 == y2)
 */
// llgo:link (*AreaT).AreaSetHeight C.lv_area_set_height
func (recv_ *AreaT) AreaSetHeight(h c.Int32T) {
}

/**
 * Return with area of an area (x * y)
 * @param area_p pointer to an area
 * @return size of area
 */
// llgo:link (*AreaT).AreaGetSize C.lv_area_get_size
func (recv_ *AreaT) AreaGetSize() c.Uint32T {
	return 0
}

// llgo:link (*AreaT).AreaIncrease C.lv_area_increase
func (recv_ *AreaT) AreaIncrease(w_extra c.Int32T, h_extra c.Int32T) {
}

// llgo:link (*AreaT).AreaMove C.lv_area_move
func (recv_ *AreaT) AreaMove(x_ofs c.Int32T, y_ofs c.Int32T) {
}

/**
 * Align an area to another
 * @param base an area where the other will be aligned
 * @param to_align the area to align
 * @param align `LV_ALIGN_...`
 * @param ofs_x X offset
 * @param ofs_y Y offset
 */
// llgo:link (*AreaT).AreaAlign C.lv_area_align
func (recv_ *AreaT) AreaAlign(to_align *AreaT, align AlignT, ofs_x c.Int32T, ofs_y c.Int32T) {
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
// llgo:link (*PointT).PointTransform C.lv_point_transform
func (recv_ *PointT) PointTransform(angle c.Int32T, scale_x c.Int32T, scale_y c.Int32T, pivot *PointT, zoom_first bool) {
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
// llgo:link (*PointT).PointArrayTransform C.lv_point_array_transform
func (recv_ *PointT) PointArrayTransform(count c.SizeT, angle c.Int32T, scale_x c.Int32T, scale_y c.Int32T, pivot *PointT, zoom_first bool) {
}

// llgo:link (*PointPreciseT).PointFromPrecise C.lv_point_from_precise
func (recv_ *PointPreciseT) PointFromPrecise() PointT {
	return PointT{}
}

// llgo:link (*PointT).PointToPrecise C.lv_point_to_precise
func (recv_ *PointT) PointToPrecise() PointPreciseT {
	return PointPreciseT{}
}

// llgo:link (*PointT).PointSet C.lv_point_set
func (recv_ *PointT) PointSet(x c.Int32T, y c.Int32T) {
}

// llgo:link (*PointPreciseT).PointPreciseSet C.lv_point_precise_set
func (recv_ *PointPreciseT) PointPreciseSet(x ValuePreciseT, y ValuePreciseT) {
}

// llgo:link (*PointT).PointSwap C.lv_point_swap
func (recv_ *PointT) PointSwap(p2 *PointT) {
}

// llgo:link (*PointPreciseT).PointPreciseSwap C.lv_point_precise_swap
func (recv_ *PointPreciseT) PointPreciseSwap(p2 *PointPreciseT) {
}

/**
 * Convert a percentage value to `int32_t`.
 * Percentage values are stored in special range
 * @param x the percentage (0..1000)
 * @return a coordinate that stores the percentage
 */
//go:linkname Pct C.lv_pct
func Pct(x c.Int32T) c.Int32T

//go:linkname PctToPx C.lv_pct_to_px
func PctToPx(v c.Int32T, base c.Int32T) c.Int32T

type LogLevelT c.Int8T

/**
 * @brief Copies a block of memory from a source address to a destination address.
 * @param dst Pointer to the destination array where the content is to be copied.
 * @param src Pointer to the source of data to be copied.
 * @param len Number of bytes to copy.
 * @return Pointer to the destination array.
 * @note The function does not check for any overlapping of the source and destination memory blocks.
 */
//go:linkname Memcpy C.lv_memcpy
func Memcpy(dst c.Pointer, src c.Pointer, len c.SizeT) c.Pointer

/**
 * @brief Fills a block of memory with a specified value.
 * @param dst Pointer to the destination array to fill with the specified value.
 * @param v Value to be set. The value is passed as an int, but the function fills
 *          the block of memory using the unsigned char conversion of this value.
 * @param len Number of bytes to be set to the value.
 */
//go:linkname Memset C.lv_memset
func Memset(dst c.Pointer, v c.Uint8T, len c.SizeT)

/**
 * @brief Move a block of memory from source to destination
 * @param dst Pointer to the destination array where the content is to be copied.
 * @param src Pointer to the source of data to be copied.
 * @param len Number of bytes to copy
 * @return Pointer to the destination array.
 */
//go:linkname Memmove C.lv_memmove
func Memmove(dst c.Pointer, src c.Pointer, len c.SizeT) c.Pointer

/**
 * @brief This function will compare two memory blocks
 * @param p1 Pointer to the first memory block
 * @param p2 Pointer to the second memory block
 * @param len Number of bytes to compare
 * @return The difference between the value of the first unmatching byte.
 */
//go:linkname Memcmp C.lv_memcmp
func Memcmp(p1 c.Pointer, p2 c.Pointer, len c.SizeT) c.Int

/**
 * @brief Computes the length of the string str up to (but not including) the terminating null character.
 * @param str Pointer to the null-terminated byte string to be examined.
 * @return The length of the string in bytes.
 */
//go:linkname Strlen C.lv_strlen
func Strlen(str *c.Char) c.SizeT

/**
 * @brief Computes the length of the string str up to (but not including) the terminating null character,
 *        or the given maximum length.
 * @param str Pointer to byte string that is null-terminated or at least max_len bytes long.
 * @param max_len Maximum number of characters to examine.
 * @return The length of the string in bytes.
 */
//go:linkname Strnlen C.lv_strnlen
func Strnlen(str *c.Char, max_len c.SizeT) c.SizeT

/**
 * @brief Copies up to dst_size-1 (non-null) characters from src to dst. A null terminator is always added.
 * @param dst Pointer to the destination array where the content is to be copied.
 * @param src Pointer to the source of data to be copied.
 * @param dst_size Maximum number of characters to be copied to dst, including the null character.
 * @return The length of src. The return value is equivalent to the value returned by lv_strlen(src)
 */
//go:linkname Strlcpy C.lv_strlcpy
func Strlcpy(dst *c.Char, src *c.Char, dst_size c.SizeT) c.SizeT

/**
 * @brief Copies up to dest_size characters from the string pointed to by src to the character array pointed to by dst
 *        and fills the remaining length with null bytes.
 * @param dst Pointer to the destination array where the content is to be copied.
 * @param src Pointer to the source of data to be copied.
 * @param dest_size Maximum number of characters to be copied to dst.
 * @return A pointer to the destination array, which is dst.
 * @note dst will not be null terminated if dest_size bytes were copied from src before the end of src was reached.
 */
//go:linkname Strncpy C.lv_strncpy
func Strncpy(dst *c.Char, src *c.Char, dest_size c.SizeT) *c.Char

/**
 * @brief Copies the string pointed to by src, including the terminating null character,
 *        to the character array pointed to by dst.
 * @param dst Pointer to the destination array where the content is to be copied.
 * @param src Pointer to the source of data to be copied.
 * @return A pointer to the destination array, which is dst.
 */
//go:linkname Strcpy C.lv_strcpy
func Strcpy(dst *c.Char, src *c.Char) *c.Char

/**
 * @brief  This function will compare two strings without specified length.
 * @param s1    pointer to the first string
 * @param s2    pointer to the second string
 * @return      the difference between the value of the first unmatching character.
 */
//go:linkname Strcmp C.lv_strcmp
func Strcmp(s1 *c.Char, s2 *c.Char) c.Int

/**
 * @brief  This function will compare two strings up to the given length.
 * @param s1    pointer to the first string
 * @param s2    pointer to the second string
 * @param len   the maximum amount of characters to compare
 * @return      the difference between the value of the first unmatching character.
 */
//go:linkname Strncmp C.lv_strncmp
func Strncmp(s1 *c.Char, s2 *c.Char, len c.SizeT) c.Int

/**
 * @brief Duplicate a string by allocating a new one and copying the content.
 * @param src Pointer to the source of data to be copied.
 * @return A pointer to the new allocated string. NULL if failed.
 */
//go:linkname Strdup C.lv_strdup
func Strdup(src *c.Char) *c.Char

/**
 * @brief Duplicate a string by allocating a new one and copying the content
 *        up to the end or the specified maximum length, whichever comes first.
 * @param src Pointer to the source of data to be copied.
 * @param max_len Maximum number of characters to be copied.
 * @return Pointer to a newly allocated null-terminated string. NULL if failed.
 */
//go:linkname Strndup C.lv_strndup
func Strndup(src *c.Char, max_len c.SizeT) *c.Char

/**
 * @brief Copies the string pointed to by src, including the terminating null character,
 *        to the end of the string pointed to by dst.
 * @param dst Pointer to the destination string where the content is to be appended.
 * @param src Pointer to the source of data to be copied.
 * @return A pointer to the destination string, which is dst.
 */
//go:linkname Strcat C.lv_strcat
func Strcat(dst *c.Char, src *c.Char) *c.Char

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
//go:linkname Strncat C.lv_strncat
func Strncat(dst *c.Char, src *c.Char, src_len c.SizeT) *c.Char

/**
 * @brief Searches for the first occurrence of character c in the string str.
 * @param str Pointer to the null-terminated byte string to be searched.
 * @param c The character to be searched for.
 * @return A pointer to the first occurrence of character c in the string str, or a null pointer if c is not found.
 */
//go:linkname Strchr C.lv_strchr
func Strchr(str *c.Char, c c.Int) *c.Char

type MemPoolT c.Pointer

/**
 * Heap information structure.
 */

type MemMonitorT struct {
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
//go:linkname MemInit C.lv_mem_init
func MemInit()

/**
 * Drop all dynamically allocated memory and reset the memory pools' state
 */
//go:linkname MemDeinit C.lv_mem_deinit
func MemDeinit()

//go:linkname MemAddPool C.lv_mem_add_pool
func MemAddPool(mem c.Pointer, bytes c.SizeT) MemPoolT

//go:linkname MemRemovePool C.lv_mem_remove_pool
func MemRemovePool(pool MemPoolT)

/**
 * Allocate memory dynamically
 * @param size requested size in bytes
 * @return pointer to allocated uninitialized memory, or NULL on failure
 */
//go:linkname Malloc C.lv_malloc
func Malloc(size c.SizeT) c.Pointer

/**
 * Allocate a block of zeroed memory dynamically
 * @param num requested number of element to be allocated.
 * @param size requested size of each element in bytes.
 * @return pointer to allocated zeroed memory, or NULL on failure
 */
//go:linkname Calloc C.lv_calloc
func Calloc(num c.SizeT, size c.SizeT) c.Pointer

/**
 * Allocate zeroed memory dynamically
 * @param size requested size in bytes
 * @return pointer to allocated zeroed memory, or NULL on failure
 */
//go:linkname Zalloc C.lv_zalloc
func Zalloc(size c.SizeT) c.Pointer

/**
 * Allocate zeroed memory dynamically
 * @param size requested size in bytes
 * @return pointer to allocated zeroed memory, or NULL on failure
 */
//go:linkname MallocZeroed C.lv_malloc_zeroed
func MallocZeroed(size c.SizeT) c.Pointer

/**
 * Free an allocated data
 * @param data pointer to an allocated memory
 */
//go:linkname Free C.lv_free
func Free(data c.Pointer)

/**
 * Reallocate a memory with a new size. The old content will be kept.
 * @param data_p pointer to an allocated memory.
 *               Its content will be copied to the new memory block and freed
 * @param new_size the desired new size in byte
 * @return pointer to the new memory, NULL on failure
 */
//go:linkname Realloc C.lv_realloc
func Realloc(data_p c.Pointer, new_size c.SizeT) c.Pointer

/**
 * Reallocate a memory with a new size. The old content will be kept.
 * In case of failure, the old pointer is free'd.
 * @param data_p pointer to an allocated memory.
 *               Its content will be copied to the new memory block and freed
 * @param new_size the desired new size in byte
 * @return pointer to the new memory, NULL on failure
 */
//go:linkname Reallocf C.lv_reallocf
func Reallocf(data_p c.Pointer, new_size c.SizeT) c.Pointer

/**
 * Used internally to execute a plain `malloc` operation
 * @param size      size in bytes to `malloc`
 */
//go:linkname MallocCore C.lv_malloc_core
func MallocCore(size c.SizeT) c.Pointer

/**
 * Used internally to execute a plain `free` operation
 * @param p      memory address to free
 */
//go:linkname FreeCore C.lv_free_core
func FreeCore(p c.Pointer)

/**
 * Used internally to execute a plain realloc operation
 * @param p         memory address to realloc
 * @param new_size  size in bytes to realloc
 */
//go:linkname ReallocCore C.lv_realloc_core
func ReallocCore(p c.Pointer, new_size c.SizeT) c.Pointer

/**
 * Used internally by lv_mem_monitor() to gather LVGL heap state information.
 * @param mon_p      pointer to lv_mem_monitor_t object to be populated.
 */
// llgo:link (*MemMonitorT).MemMonitorCore C.lv_mem_monitor_core
func (recv_ *MemMonitorT) MemMonitorCore() {
}

//go:linkname MemTestCore C.lv_mem_test_core
func MemTestCore() ResultT

/**
 * @brief Tests the memory allocation system by allocating and freeing a block of memory.
 * @return LV_RESULT_OK if the memory allocation system is working properly, or LV_RESULT_INVALID if there is an error.
 */
//go:linkname MemTest C.lv_mem_test
func MemTest() ResultT

/**
 * Give information about the work memory of dynamic allocation
 * @param mon_p pointer to a lv_mem_monitor_t variable,
 *              the result of the analysis will be stored here
 */
// llgo:link (*MemMonitorT).MemMonitor C.lv_mem_monitor
func (recv_ *MemMonitorT) MemMonitor() {
}

/**********************
 * GLOBAL PROTOTYPES
 **********************/
//go:linkname AssertHandler C.lv_assert_handler
func AssertHandler()

type X_lvOpacityLevelT c.Int

const (
	OPA_TRANSP X_lvOpacityLevelT = 0
	OPA_0      X_lvOpacityLevelT = 0
	OPA_10     X_lvOpacityLevelT = 25
	OPA_20     X_lvOpacityLevelT = 51
	OPA_30     X_lvOpacityLevelT = 76
	OPA_40     X_lvOpacityLevelT = 102
	OPA_50     X_lvOpacityLevelT = 127
	OPA_60     X_lvOpacityLevelT = 153
	OPA_70     X_lvOpacityLevelT = 178
	OPA_80     X_lvOpacityLevelT = 204
	OPA_90     X_lvOpacityLevelT = 229
	OPA_100    X_lvOpacityLevelT = 255
	OPA_COVER  X_lvOpacityLevelT = 255
)

/**********************
 *      TYPEDEFS
 **********************/

type ColorT struct {
	Blue  c.Uint8T
	Green c.Uint8T
	Red   c.Uint8T
}

type Color16T struct {
	Blue  c.Uint16T
	Green c.Uint16T
	Red   c.Uint16T
}

type Color32T struct {
	Blue  c.Uint8T
	Green c.Uint8T
	Red   c.Uint8T
	Alpha c.Uint8T
}

type ColorHsvT struct {
	H c.Uint16T
	S c.Uint8T
	V c.Uint8T
}

type Color16aT struct {
	Lumi  c.Uint8T
	Alpha c.Uint8T
}
type ColorFormatT c.Int

const (
	COLOR_FORMAT_UNKNOWN                ColorFormatT = 0
	COLOR_FORMAT_RAW                    ColorFormatT = 1
	COLOR_FORMAT_RAW_ALPHA              ColorFormatT = 2
	COLOR_FORMAT_L8                     ColorFormatT = 6
	COLOR_FORMAT_I1                     ColorFormatT = 7
	COLOR_FORMAT_I2                     ColorFormatT = 8
	COLOR_FORMAT_I4                     ColorFormatT = 9
	COLOR_FORMAT_I8                     ColorFormatT = 10
	COLOR_FORMAT_A8                     ColorFormatT = 14
	COLOR_FORMAT_RGB565                 ColorFormatT = 18
	COLOR_FORMAT_ARGB8565               ColorFormatT = 19
	COLOR_FORMAT_RGB565A8               ColorFormatT = 20
	COLOR_FORMAT_AL88                   ColorFormatT = 21
	COLOR_FORMAT_RGB565_SWAPPED         ColorFormatT = 27
	COLOR_FORMAT_RGB888                 ColorFormatT = 15
	COLOR_FORMAT_ARGB8888               ColorFormatT = 16
	COLOR_FORMAT_XRGB8888               ColorFormatT = 17
	COLOR_FORMAT_ARGB8888_PREMULTIPLIED ColorFormatT = 26
	COLOR_FORMAT_A1                     ColorFormatT = 11
	COLOR_FORMAT_A2                     ColorFormatT = 12
	COLOR_FORMAT_A4                     ColorFormatT = 13
	COLOR_FORMAT_ARGB1555               ColorFormatT = 22
	COLOR_FORMAT_ARGB4444               ColorFormatT = 23
	COLOR_FORMAT_ARGB2222               ColorFormatT = 24
	COLOR_FORMAT_YUV_START              ColorFormatT = 32
	COLOR_FORMAT_I420                   ColorFormatT = 32
	COLOR_FORMAT_I422                   ColorFormatT = 33
	COLOR_FORMAT_I444                   ColorFormatT = 34
	COLOR_FORMAT_I400                   ColorFormatT = 35
	COLOR_FORMAT_NV21                   ColorFormatT = 36
	COLOR_FORMAT_NV12                   ColorFormatT = 37
	COLOR_FORMAT_YUY2                   ColorFormatT = 38
	COLOR_FORMAT_UYVY                   ColorFormatT = 39
	COLOR_FORMAT_YUV_END                ColorFormatT = 39
	COLOR_FORMAT_PROPRIETARY_START      ColorFormatT = 48
	COLOR_FORMAT_NEMA_TSC_START         ColorFormatT = 48
	COLOR_FORMAT_NEMA_TSC4              ColorFormatT = 48
	COLOR_FORMAT_NEMA_TSC6              ColorFormatT = 49
	COLOR_FORMAT_NEMA_TSC6A             ColorFormatT = 50
	COLOR_FORMAT_NEMA_TSC6AP            ColorFormatT = 51
	COLOR_FORMAT_NEMA_TSC12             ColorFormatT = 52
	COLOR_FORMAT_NEMA_TSC12A            ColorFormatT = 53
	COLOR_FORMAT_NEMA_TSC_END           ColorFormatT = 53
	COLOR_FORMAT_NATIVE                 ColorFormatT = 18
	COLOR_FORMAT_NATIVE_WITH_ALPHA      ColorFormatT = 20
)

/**
 * Get the pixel size of a color format in bits, bpp
 * @param cf        a color format (`LV_COLOR_FORMAT_...`)
 * @return          the pixel size in bits
 * @sa              LV_COLOR_FORMAT_GET_BPP
 */
// llgo:link ColorFormatT.ColorFormatGetBpp C.lv_color_format_get_bpp
func (recv_ ColorFormatT) ColorFormatGetBpp() c.Uint8T {
	return 0
}

/**
 * Get the pixel size of a color format in bytes
 * @param cf        a color format (`LV_COLOR_FORMAT_...`)
 * @return          the pixel size in bytes
 * @sa              LV_COLOR_FORMAT_GET_SIZE
 */
// llgo:link ColorFormatT.ColorFormatGetSize C.lv_color_format_get_size
func (recv_ ColorFormatT) ColorFormatGetSize() c.Uint8T {
	return 0
}

/**
 * Check if a color format has alpha channel or not
 * @param src_cf    a color format (`LV_COLOR_FORMAT_...`)
 * @return          true: has alpha channel; false: doesn't have alpha channel
 */
// llgo:link ColorFormatT.ColorFormatHasAlpha C.lv_color_format_has_alpha
func (recv_ ColorFormatT) ColorFormatHasAlpha() bool {
	return false
}

/**
 * Create an ARGB8888 color from RGB888 + alpha
 * @param color     an RGB888 color
 * @param opa       the alpha value
 * @return          the ARGB8888 color
 */
// llgo:link ColorT.ColorTo32 C.lv_color_to_32
func (recv_ ColorT) ColorTo32(opa OpaT) Color32T {
	return Color32T{}
}

/**
 * Convert an RGB888 color to an integer
 * @param c     an RGB888 color
 * @return      `c` as an integer
 */
// llgo:link ColorT.ColorToInt C.lv_color_to_int
func (recv_ ColorT) ColorToInt() c.Uint32T {
	return 0
}

/**
 * Check if two RGB888 color are equal
 * @param c1    the first color
 * @param c2    the second color
 * @return      true: equal
 */
// llgo:link ColorT.ColorEq C.lv_color_eq
func (recv_ ColorT) ColorEq(c2 ColorT) bool {
	return false
}

/**
 * Check if two ARGB8888 color are equal
 * @param c1    the first color
 * @param c2    the second color
 * @return      true: equal
 */
// llgo:link Color32T.Color32Eq C.lv_color32_eq
func (recv_ Color32T) Color32Eq(c2 Color32T) bool {
	return false
}

/**
 * Create a color from 0x000000..0xffffff input
 * @param c     the hex input
 * @return      the color
 */
//go:linkname ColorHex C.lv_color_hex
func ColorHex(c c.Uint32T) ColorT

/**
 * Create an RGB888 color
 * @param r     the red channel (0..255)
 * @param g     the green channel (0..255)
 * @param b     the blue channel (0..255)
 * @return      the color
 */
//go:linkname ColorMake C.lv_color_make
func ColorMake(r c.Uint8T, g c.Uint8T, b c.Uint8T) ColorT

/**
 * Create an ARGB8888 color
 * @param r     the red channel (0..255)
 * @param g     the green channel (0..255)
 * @param b     the blue channel (0..255)
 * @param a     the alpha channel (0..255)
 * @return      the color
 */
//go:linkname Color32Make C.lv_color32_make
func Color32Make(r c.Uint8T, g c.Uint8T, b c.Uint8T, a c.Uint8T) Color32T

/**
 * Create a color from 0x000..0xfff input
 * @param c     the hex input (e.g. 0x123 will be 0x112233)
 * @return      the color
 */
//go:linkname ColorHex3 C.lv_color_hex3
func ColorHex3(c c.Uint32T) ColorT

/**
 * Convert am RGB888 color to RGB565 stored in `uint16_t`
 * @param color     and RGB888 color
 * @return          `color` as RGB565 on `uin16_t`
 */
// llgo:link ColorT.ColorToU16 C.lv_color_to_u16
func (recv_ ColorT) ColorToU16() c.Uint16T {
	return 0
}

/**
 * Convert am RGB888 color to XRGB8888 stored in `uint32_t`
 * @param color     and RGB888 color
 * @return          `color` as XRGB8888 on `uin32_t` (the alpha channel is always set to 0xFF)
 */
// llgo:link ColorT.ColorToU32 C.lv_color_to_u32
func (recv_ ColorT) ColorToU32() c.Uint32T {
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
//go:linkname Color1616Mix C.lv_color_16_16_mix
func Color1616Mix(c1 c.Uint16T, c2 c.Uint16T, mix c.Uint8T) c.Uint16T

/**
 * Mix white to a color
 * @param c     the base color
 * @param lvl   the intensity of white (0: no change, 255: fully white)
 * @return      the mixed color
 */
// llgo:link ColorT.ColorLighten C.lv_color_lighten
func (recv_ ColorT) ColorLighten(lvl OpaT) ColorT {
	return ColorT{}
}

/**
 * Mix black to a color
 * @param c     the base color
 * @param lvl   the intensity of black (0: no change, 255: fully black)
 * @return      the mixed color
 */
// llgo:link ColorT.ColorDarken C.lv_color_darken
func (recv_ ColorT) ColorDarken(lvl OpaT) ColorT {
	return ColorT{}
}

/**
 * Convert a HSV color to RGB
 * @param h hue [0..359]
 * @param s saturation [0..100]
 * @param v value [0..100]
 * @return the given RGB color in RGB (with LV_COLOR_DEPTH depth)
 */
//go:linkname ColorHsvToRgb C.lv_color_hsv_to_rgb
func ColorHsvToRgb(h c.Uint16T, s c.Uint8T, v c.Uint8T) ColorT

/**
 * Convert a 32-bit RGB color to HSV
 * @param r8 8-bit red
 * @param g8 8-bit green
 * @param b8 8-bit blue
 * @return the given RGB color in HSV
 */
//go:linkname ColorRgbToHsv C.lv_color_rgb_to_hsv
func ColorRgbToHsv(r8 c.Uint8T, g8 c.Uint8T, b8 c.Uint8T) ColorHsvT

/**
 * Convert a color to HSV
 * @param color color
 * @return the given color in HSV
 */
// llgo:link ColorT.ColorToHsv C.lv_color_to_hsv
func (recv_ ColorT) ColorToHsv() ColorHsvT {
	return ColorHsvT{}
}

/**
 * A helper for white color
 * @return      a white color
 */
//go:linkname ColorWhite C.lv_color_white
func ColorWhite() ColorT

/**
 * A helper for black color
 * @return      a black color
 */
//go:linkname ColorBlack C.lv_color_black
func ColorBlack() ColorT

// llgo:link (*Color32T).ColorPremultiply C.lv_color_premultiply
func (recv_ *Color32T) ColorPremultiply() {
}

// llgo:link (*Color16T).Color16Premultiply C.lv_color16_premultiply
func (recv_ *Color16T) Color16Premultiply(a OpaT) {
}

/**
 * Get the luminance of a color: luminance = 0.3 R + 0.59 G + 0.11 B
 * @param c a color
 * @return the brightness [0..255]
 */
// llgo:link ColorT.ColorLuminance C.lv_color_luminance
func (recv_ ColorT) ColorLuminance() c.Uint8T {
	return 0
}

/**
 * Get the luminance of a color16: luminance = 0.3 R + 0.59 G + 0.11 B
 * @param c a color
 * @return the brightness [0..255]
 */
// llgo:link Color16T.Color16Luminance C.lv_color16_luminance
func (recv_ Color16T) Color16Luminance() c.Uint8T {
	return 0
}

/**
 * Get the luminance of a color24: luminance = 0.3 R + 0.59 G + 0.11 B
 * @param c a color
 * @return the brightness [0..255]
 */
//go:linkname Color24Luminance C.lv_color24_luminance
func Color24Luminance(c *c.Uint8T) c.Uint8T

/**
 * Get the luminance of a color32: luminance = 0.3 R + 0.59 G + 0.11 B
 * @param c a color
 * @return the brightness [0..255]
 */
// llgo:link Color32T.Color32Luminance C.lv_color32_luminance
func (recv_ Color32T) Color32Luminance() c.Uint8T {
	return 0
}

type PaletteT c.Int

const (
	PALETTE_RED         PaletteT = 0
	PALETTE_PINK        PaletteT = 1
	PALETTE_PURPLE      PaletteT = 2
	PALETTE_DEEP_PURPLE PaletteT = 3
	PALETTE_INDIGO      PaletteT = 4
	PALETTE_BLUE        PaletteT = 5
	PALETTE_LIGHT_BLUE  PaletteT = 6
	PALETTE_CYAN        PaletteT = 7
	PALETTE_TEAL        PaletteT = 8
	PALETTE_GREEN       PaletteT = 9
	PALETTE_LIGHT_GREEN PaletteT = 10
	PALETTE_LIME        PaletteT = 11
	PALETTE_YELLOW      PaletteT = 12
	PALETTE_AMBER       PaletteT = 13
	PALETTE_ORANGE      PaletteT = 14
	PALETTE_DEEP_ORANGE PaletteT = 15
	PALETTE_BROWN       PaletteT = 16
	PALETTE_BLUE_GREY   PaletteT = 17
	PALETTE_GREY        PaletteT = 18
	PALETTE_LAST        PaletteT = 19
	PALETTE_NONE        PaletteT = 255
)

/*Source: https://vuetifyjs.com/en/styles/colors/#material-colors*/
// llgo:link PaletteT.PaletteMain C.lv_palette_main
func (recv_ PaletteT) PaletteMain() ColorT {
	return ColorT{}
}

// llgo:link PaletteT.PaletteLighten C.lv_palette_lighten
func (recv_ PaletteT) PaletteLighten(lvl c.Uint8T) ColorT {
	return ColorT{}
}

// llgo:link PaletteT.PaletteDarken C.lv_palette_darken
func (recv_ PaletteT) PaletteDarken(lvl c.Uint8T) ColorT {
	return ColorT{}
}

// llgo:type C
type ColorFilterCbT func(*X_lvColorFilterDscT, ColorT, OpaT) ColorT

/**
 * Mix two colors with a given ratio.
 * @param c1 the first color to mix (usually the foreground)
 * @param c2 the second color to mix (usually the background)
 * @param mix The ratio of the colors. 0: full `c2`, 255: full `c1`, 127: half `c1` and half`c2`
 * @return the mixed color
 */
// llgo:link ColorT.ColorMix C.lv_color_mix
func (recv_ ColorT) ColorMix(c2 ColorT, mix c.Uint8T) ColorT {
	return ColorT{}
}

/**
 *
 * @param fg
 * @param bg
 * @return
 * @note Use bg.alpha in the return value
 * @note Use fg.alpha as mix ratio
 */
// llgo:link Color32T.ColorMix32 C.lv_color_mix32
func (recv_ Color32T) ColorMix32(bg Color32T) Color32T {
	return Color32T{}
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
// llgo:link Color32T.ColorMix32Premultiplied C.lv_color_mix32_premultiplied
func (recv_ Color32T) ColorMix32Premultiplied(bg Color32T) Color32T {
	return Color32T{}
}

/**
 * Get the brightness of a color
 * @param c   a color
 * @return brightness in range [0..255]
 */
// llgo:link ColorT.ColorBrightness C.lv_color_brightness
func (recv_ ColorT) ColorBrightness() c.Uint8T {
	return 0
}

// llgo:link (*ColorFilterDscT).ColorFilterDscInit C.lv_color_filter_dsc_init
func (recv_ *ColorFilterDscT) ColorFilterDscInit(cb ColorFilterCbT) {
}

/**
 * Blend two colors that have not been pre-multiplied using their alpha values
 * @param fg the foreground color
 * @param bg the background color
 * @return result color
 */
// llgo:link Color32T.ColorOver32 C.lv_color_over32
func (recv_ Color32T) ColorOver32(bg Color32T) Color32T {
	return Color32T{}
}

type X_lvimageFlagsT c.Int

const (
	IMAGE_FLAGS_PREMULTIPLIED X_lvimageFlagsT = 1
	IMAGE_FLAGS_COMPRESSED    X_lvimageFlagsT = 8
	IMAGE_FLAGS_ALLOCATED     X_lvimageFlagsT = 16
	IMAGE_FLAGS_MODIFIABLE    X_lvimageFlagsT = 32
	IMAGE_FLAGS_CUSTOM_DRAW   X_lvimageFlagsT = 64
	IMAGE_FLAGS_USER1         X_lvimageFlagsT = 256
	IMAGE_FLAGS_USER2         X_lvimageFlagsT = 512
	IMAGE_FLAGS_USER3         X_lvimageFlagsT = 1024
	IMAGE_FLAGS_USER4         X_lvimageFlagsT = 2048
	IMAGE_FLAGS_USER5         X_lvimageFlagsT = 4096
	IMAGE_FLAGS_USER6         X_lvimageFlagsT = 8192
	IMAGE_FLAGS_USER7         X_lvimageFlagsT = 16384
	IMAGE_FLAGS_USER8         X_lvimageFlagsT = 32768
)

type ImageFlagsT X_lvimageFlagsT
type ImageCompressT c.Int

const (
	IMAGE_COMPRESS_NONE ImageCompressT = 0
	IMAGE_COMPRESS_RLE  ImageCompressT = 1
	IMAGE_COMPRESS_LZ4  ImageCompressT = 2
)

type ImageHeaderT struct {
	Magic     c.Uint32T
	Cf        c.Uint32T
	Flags     c.Uint32T
	W         c.Uint32T
	H         c.Uint32T
	Stride    c.Uint32T
	Reserved2 c.Uint32T
}

type YuvPlaneT struct {
	Buf    c.Pointer
	Stride c.Uint32T
}

type YuvBufT struct {
	Planar struct {
		Y YuvPlaneT
		U YuvPlaneT
		V YuvPlaneT
	}
}

/**
 * Struct to describe a constant image resource.
 * It's similar to lv_draw_buf_t, but the data is constant.
 */

type ImageDscT struct {
	Header    ImageHeaderT
	DataSize  c.Uint32T
	Data      *c.Uint8T
	Reserved  c.Pointer
	Reserved2 c.Pointer
}

// llgo:type C
type DrawBufMallocCb func(c.SizeT, ColorFormatT) c.Pointer

// llgo:type C
type DrawBufFreeCb func(c.Pointer)

// llgo:type C
type DrawBufAlignCb func(c.Pointer, ColorFormatT) c.Pointer

// llgo:type C
type DrawBufCacheOperationCb func(*DrawBufT, *AreaT)

// llgo:type C
type DrawBufWidthToStrideCb func(c.Uint32T, ColorFormatT) c.Uint32T

/**
 * Initialize the draw buffer with the default handlers.
 *
 * @param handlers  the draw buffer handlers to set
 */
// llgo:link (*DrawBufHandlersT).DrawBufInitWithDefaultHandlers C.lv_draw_buf_init_with_default_handlers
func (recv_ *DrawBufHandlersT) DrawBufInitWithDefaultHandlers() {
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
// llgo:link (*DrawBufHandlersT).DrawBufHandlersInit C.lv_draw_buf_handlers_init
func (recv_ *DrawBufHandlersT) DrawBufHandlersInit(buf_malloc_cb DrawBufMallocCb, buf_free_cb DrawBufFreeCb, align_pointer_cb DrawBufAlignCb, invalidate_cache_cb DrawBufCacheOperationCb, flush_cache_cb DrawBufCacheOperationCb, width_to_stride_cb DrawBufWidthToStrideCb) {
}

/**
 * Get the struct which holds the callbacks for draw buf management.
 * Custom callback can be set on the returned value
 * @return                  pointer to the struct of handlers
 */
//go:linkname DrawBufGetHandlers C.lv_draw_buf_get_handlers
func DrawBufGetHandlers() *DrawBufHandlersT

//go:linkname DrawBufGetFontHandlers C.lv_draw_buf_get_font_handlers
func DrawBufGetFontHandlers() *DrawBufHandlersT

//go:linkname DrawBufGetImageHandlers C.lv_draw_buf_get_image_handlers
func DrawBufGetImageHandlers() *DrawBufHandlersT

/**
 * Align the address of a buffer. The buffer needs to be large enough for the real data after alignment
 * @param buf           the data to align
 * @param color_format  the color format of the buffer
 * @return              the aligned buffer
 */
//go:linkname DrawBufAlign C.lv_draw_buf_align
func DrawBufAlign(buf c.Pointer, color_format ColorFormatT) c.Pointer

/**
 * Align the address of a buffer with custom draw buffer handlers.
 * The buffer needs to be large enough for the real data after alignment
 * @param handlers      the draw buffer handlers
 * @param buf           the data to align
 * @param color_format  the color format of the buffer
 * @return              the aligned buffer
 */
// llgo:link (*DrawBufHandlersT).DrawBufAlignEx C.lv_draw_buf_align_ex
func (recv_ *DrawBufHandlersT) DrawBufAlignEx(buf c.Pointer, color_format ColorFormatT) c.Pointer {
	return nil
}

/**
 * Invalidate the cache of the buffer
 * @param draw_buf     the draw buffer needs to be invalidated
 * @param area         the area to invalidate in the buffer,
 *                     use NULL to invalidate the whole draw buffer address range
 */
// llgo:link (*DrawBufT).DrawBufInvalidateCache C.lv_draw_buf_invalidate_cache
func (recv_ *DrawBufT) DrawBufInvalidateCache(area *AreaT) {
}

/**
 * Flush the cache of the buffer
 * @param draw_buf     the draw buffer needs to be flushed
 * @param area         the area to flush in the buffer,
 *                     use NULL to flush the whole draw buffer address range
 */
// llgo:link (*DrawBufT).DrawBufFlushCache C.lv_draw_buf_flush_cache
func (recv_ *DrawBufT) DrawBufFlushCache(area *AreaT) {
}

/**
 * Calculate the stride in bytes based on a width and color format
 * @param w                 the width in pixels
 * @param color_format      the color format
 * @return                  the stride in bytes
 */
//go:linkname DrawBufWidthToStride C.lv_draw_buf_width_to_stride
func DrawBufWidthToStride(w c.Uint32T, color_format ColorFormatT) c.Uint32T

/**
 * Calculate the stride in bytes based on a width and color format
 * @param handlers          the draw buffer handlers
 * @param w                 the width in pixels
 * @param color_format      the color format
 * @return                  the stride in bytes
 */
// llgo:link (*DrawBufHandlersT).DrawBufWidthToStrideEx C.lv_draw_buf_width_to_stride_ex
func (recv_ *DrawBufHandlersT) DrawBufWidthToStrideEx(w c.Uint32T, color_format ColorFormatT) c.Uint32T {
	return 0
}

/**
 * Clear an area on the buffer
 * @param draw_buf          pointer to draw buffer
 * @param a                 the area to clear, or NULL to clear the whole buffer
 */
// llgo:link (*DrawBufT).DrawBufClear C.lv_draw_buf_clear
func (recv_ *DrawBufT) DrawBufClear(a *AreaT) {
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
// llgo:link (*DrawBufT).DrawBufCopy C.lv_draw_buf_copy
func (recv_ *DrawBufT) DrawBufCopy(dest_area *AreaT, src *DrawBufT, src_area *AreaT) {
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
//go:linkname DrawBufCreate C.lv_draw_buf_create
func DrawBufCreate(w c.Uint32T, h c.Uint32T, cf ColorFormatT, stride c.Uint32T) *DrawBufT

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
// llgo:link (*DrawBufHandlersT).DrawBufCreateEx C.lv_draw_buf_create_ex
func (recv_ *DrawBufHandlersT) DrawBufCreateEx(w c.Uint32T, h c.Uint32T, cf ColorFormatT, stride c.Uint32T) *DrawBufT {
	return nil
}

/**
 * Duplicate a draw buf with same image size, stride and color format. Copy the image data too.
 * @param draw_buf  the draw buf to duplicate
 * @return          the duplicated draw buf on success, NULL if failed
 */
// llgo:link (*DrawBufT).DrawBufDup C.lv_draw_buf_dup
func (recv_ *DrawBufT) DrawBufDup() *DrawBufT {
	return nil
}

/**
 * Duplicate a draw buf with same image size, stride and color format. Copy the image data too.
 * @param handlers  the draw buffer handlers
 * @param draw_buf  the draw buf to duplicate
 * @return          the duplicated draw buf on success, NULL if failed
 */
// llgo:link (*DrawBufHandlersT).DrawBufDupEx C.lv_draw_buf_dup_ex
func (recv_ *DrawBufHandlersT) DrawBufDupEx(draw_buf *DrawBufT) *DrawBufT {
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
// llgo:link (*DrawBufT).DrawBufInit C.lv_draw_buf_init
func (recv_ *DrawBufT) DrawBufInit(w c.Uint32T, h c.Uint32T, cf ColorFormatT, stride c.Uint32T, data c.Pointer, data_size c.Uint32T) ResultT {
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
// llgo:link (*DrawBufT).DrawBufReshape C.lv_draw_buf_reshape
func (recv_ *DrawBufT) DrawBufReshape(cf ColorFormatT, w c.Uint32T, h c.Uint32T, stride c.Uint32T) *DrawBufT {
	return nil
}

/**
 * Destroy a draw buf by freeing the actual buffer if it's marked as LV_IMAGE_FLAGS_ALLOCATED in header.
 * Then free the lv_draw_buf_t struct.
 *
 * @param draw_buf  the draw buffer to destroy
 */
// llgo:link (*DrawBufT).DrawBufDestroy C.lv_draw_buf_destroy
func (recv_ *DrawBufT) DrawBufDestroy() {
}

/**
 * Return pointer to the buffer at the given coordinates
 */
// llgo:link (*DrawBufT).DrawBufGotoXy C.lv_draw_buf_goto_xy
func (recv_ *DrawBufT) DrawBufGotoXy(x c.Uint32T, y c.Uint32T) c.Pointer {
	return nil
}

/**
 * Adjust the stride of a draw buf in place.
 * @param src       pointer to a draw buffer
 * @param stride    the new stride in bytes for image. Use LV_STRIDE_AUTO for automatic calculation.
 * @return          LV_RESULT_OK: success or LV_RESULT_INVALID: failed
 */
// llgo:link (*DrawBufT).DrawBufAdjustStride C.lv_draw_buf_adjust_stride
func (recv_ *DrawBufT) DrawBufAdjustStride(stride c.Uint32T) ResultT {
	return 0
}

/**
 * Premultiply draw buffer color with alpha channel.
 * If it's already premultiplied, return directly.
 * Only color formats with alpha channel will be processed.
 *
 * @return LV_RESULT_OK: premultiply success
 */
// llgo:link (*DrawBufT).DrawBufPremultiply C.lv_draw_buf_premultiply
func (recv_ *DrawBufT) DrawBufPremultiply() ResultT {
	return 0
}

// llgo:link (*DrawBufT).DrawBufHasFlag C.lv_draw_buf_has_flag
func (recv_ *DrawBufT) DrawBufHasFlag(flag ImageFlagsT) bool {
	return false
}

// llgo:link (*DrawBufT).DrawBufSetFlag C.lv_draw_buf_set_flag
func (recv_ *DrawBufT) DrawBufSetFlag(flag ImageFlagsT) {
}

// llgo:link (*DrawBufT).DrawBufClearFlag C.lv_draw_buf_clear_flag
func (recv_ *DrawBufT) DrawBufClearFlag(flag ImageFlagsT) {
}

/**
 * As of now, draw buf share same definition as `lv_image_dsc_t`.
 * And is interchangeable with `lv_image_dsc_t`.
 */
// llgo:link (*DrawBufT).DrawBufFromImage C.lv_draw_buf_from_image
func (recv_ *DrawBufT) DrawBufFromImage(img *ImageDscT) ResultT {
	return 0
}

// llgo:link (*DrawBufT).DrawBufToImage C.lv_draw_buf_to_image
func (recv_ *DrawBufT) DrawBufToImage(img *ImageDscT) {
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
// llgo:link (*DrawBufT).DrawBufSetPalette C.lv_draw_buf_set_palette
func (recv_ *DrawBufT) DrawBufSetPalette(index c.Uint8T, color Color32T) {
}

/**
 * @deprecated Use lv_draw_buf_set_palette instead.
 */
// llgo:link (*ImageDscT).ImageBufSetPalette C.lv_image_buf_set_palette
func (recv_ *ImageDscT) ImageBufSetPalette(id c.Uint8T, c Color32T) {
}

/**
 * @deprecated Use lv_draw_buffer_create/destroy instead.
 * Free the data pointer and dsc struct of an image.
 */
// llgo:link (*ImageDscT).ImageBufFree C.lv_image_buf_free
func (recv_ *ImageDscT) ImageBufFree() {
}

type FontGlyphFormatT c.Int

const (
	FONT_GLYPH_FORMAT_NONE   FontGlyphFormatT = 0
	FONT_GLYPH_FORMAT_A1     FontGlyphFormatT = 1
	FONT_GLYPH_FORMAT_A2     FontGlyphFormatT = 2
	FONT_GLYPH_FORMAT_A3     FontGlyphFormatT = 3
	FONT_GLYPH_FORMAT_A4     FontGlyphFormatT = 4
	FONT_GLYPH_FORMAT_A8     FontGlyphFormatT = 8
	FONT_GLYPH_FORMAT_IMAGE  FontGlyphFormatT = 25
	FONT_GLYPH_FORMAT_VECTOR FontGlyphFormatT = 26
	FONT_GLYPH_FORMAT_SVG    FontGlyphFormatT = 27
	FONT_GLYPH_FORMAT_CUSTOM FontGlyphFormatT = 255
)

/** Describes the properties of a glyph.*/

type FontGlyphDscT struct {
	ResolvedFont       *FontT
	AdvW               c.Uint16T
	BoxW               c.Uint16T
	BoxH               c.Uint16T
	OfsX               c.Int16T
	OfsY               c.Int16T
	Stride             c.Uint16T
	Format             FontGlyphFormatT
	IsPlaceholder      c.Uint8T
	ReqRawBitmap       c.Uint8T
	OutlineStrokeWidth c.Int32T
	Gid                struct {
		Src c.Pointer
	}
	Entry *CacheEntryT
}
type FontSubpxT c.Int

const (
	FONT_SUBPX_NONE FontSubpxT = 0
	FONT_SUBPX_HOR  FontSubpxT = 1
	FONT_SUBPX_VER  FontSubpxT = 2
	FONT_SUBPX_BOTH FontSubpxT = 3
)

type FontKerningT c.Int

const (
	FONT_KERNING_NORMAL FontKerningT = 0
	FONT_KERNING_NONE   FontKerningT = 1
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
// llgo:link (*FontGlyphDscT).FontGetGlyphBitmap C.lv_font_get_glyph_bitmap
func (recv_ *FontGlyphDscT) FontGetGlyphBitmap(draw_buf *DrawBufT) c.Pointer {
	return nil
}

/**
 * Return the bitmap as it is. It works only if the font stores the bitmap in
 * a non-volitile memory.
 * @param g_dsc         the glyph descriptor including which font to use, which supply the glyph_index
 *                      and the format.
 * @return              the bitmap as it is
 */
// llgo:link (*FontGlyphDscT).FontGetGlyphStaticBitmap C.lv_font_get_glyph_static_bitmap
func (recv_ *FontGlyphDscT) FontGetGlyphStaticBitmap() c.Pointer {
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
// llgo:link (*FontT).FontGetGlyphDsc C.lv_font_get_glyph_dsc
func (recv_ *FontT) FontGetGlyphDsc(dsc_out *FontGlyphDscT, letter c.Uint32T, letter_next c.Uint32T) bool {
	return false
}

/**
 * Release the bitmap of a font.
 * @note You must call lv_font_get_glyph_dsc() to get `g_dsc` (lv_font_glyph_dsc_t) before you can call this function.
 * @param g_dsc         the glyph descriptor including which font to use, which supply the glyph_index and the format.
 */
// llgo:link (*FontGlyphDscT).FontGlyphReleaseDrawData C.lv_font_glyph_release_draw_data
func (recv_ *FontGlyphDscT) FontGlyphReleaseDrawData() {
}

/**
 * Get the width of a glyph with kerning
 * @param font          pointer to a font
 * @param letter        a UNICODE letter
 * @param letter_next   the next letter after `letter`. Used for kerning
 * @return the width of the glyph
 */
// llgo:link (*FontT).FontGetGlyphWidth C.lv_font_get_glyph_width
func (recv_ *FontT) FontGetGlyphWidth(letter c.Uint32T, letter_next c.Uint32T) c.Uint16T {
	return 0
}

/**
 * Get the line height of a font. All characters fit into this height
 * @param font      pointer to a font
 * @return the height of a font
 */
// llgo:link (*FontT).FontGetLineHeight C.lv_font_get_line_height
func (recv_ *FontT) FontGetLineHeight() c.Int32T {
	return 0
}

/**
 * Configure the use of kerning information stored in a font
 * @param font    pointer to a font
 * @param kerning `LV_FONT_KERNING_NORMAL` (default) or `LV_FONT_KERNING_NONE`
 */
// llgo:link (*FontT).FontSetKerning C.lv_font_set_kerning
func (recv_ *FontT) FontSetKerning(kerning FontKerningT) {
}

/**
 * Get the default font, defined by LV_FONT_DEFAULT
 * @return  return      pointer to the default font
 */
//go:linkname FontGetDefault C.lv_font_get_default
func FontGetDefault() *FontT

/**
 * Compare font information.
 * @param ft_info_1 font information 1.
 * @param ft_info_2 font information 2.
 * @return return true if the fonts are equal.
 */
// llgo:link (*FontInfoT).FontInfoIsEqual C.lv_font_info_is_equal
func (recv_ *FontInfoT) FontInfoIsEqual(ft_info_2 *FontInfoT) bool {
	return false
}

/**
 * Checks if a font has a static rendering bitmap.
 * @param font    pointer to a font
 * @return return true if the font has a bitmap generated for static rendering.
 */
// llgo:link (*FontT).FontHasStaticBitmap C.lv_font_has_static_bitmap
func (recv_ *FontT) FontHasStaticBitmap() bool {
	return false
}

// llgo:type C
type TickGetCbT func() c.Uint32T

// llgo:type C
type DelayCbT func(c.Uint32T)

/**
 * You have to call this function periodically
 * @param tick_period   the call period of this function in milliseconds
 */
//go:linkname TickInc C.lv_tick_inc
func TickInc(tick_period c.Uint32T)

/**
 * Get the elapsed milliseconds since start up
 * @return          the elapsed milliseconds
 */
//go:linkname TickGet C.lv_tick_get
func TickGet() c.Uint32T

/**
 * Get the elapsed milliseconds since a previous time stamp
 * @param prev_tick     a previous time stamp (return value of lv_tick_get() )
 * @return              the elapsed milliseconds since 'prev_tick'
 */
//go:linkname TickElaps C.lv_tick_elaps
func TickElaps(prev_tick c.Uint32T) c.Uint32T

/**
 * Get the elapsed milliseconds between two time stamps
 * @param tick          a time stamp
 * @param prev_tick     a time stamp before `tick`
 * @return              the elapsed milliseconds between `prev_tick` and `tick`
 */
//go:linkname TickDiff C.lv_tick_diff
func TickDiff(tick c.Uint32T, prev_tick c.Uint32T) c.Uint32T

/**
 * Delay for the given milliseconds.
 * By default it's a blocking delay, but with `lv_delay_set_cb()`
 * a custom delay function can be set too
 * @param ms        the number of milliseconds to delay
 */
//go:linkname DelayMs C.lv_delay_ms
func DelayMs(ms c.Uint32T)

/**
 * Set a callback for a blocking delay
 * @param cb        pointer to a callback
 */
//go:linkname DelaySetCb C.lv_delay_set_cb
func DelaySetCb(cb DelayCbT)

/**
 * Set the custom callback for 'lv_tick_get'
 * @param cb        call this callback on 'lv_tick_get'
 */
//go:linkname TickSetCb C.lv_tick_set_cb
func TickSetCb(cb TickGetCbT)

/**
 * Get the custom callback for 'lv_tick_get'
 * @return      call this callback on 'lv_tick_get'
 */
//go:linkname TickGetCb C.lv_tick_get_cb
func TickGetCb() TickGetCbT

type LlNodeT c.Uint8T

/** Description of a linked list*/

type LlT struct {
	NSize c.Uint32T
	Head  *LlNodeT
	Tail  *LlNodeT
}

/**
 * Initialize linked list
 * @param ll_p pointer to lv_ll_t variable
 * @param node_size the size of 1 node in bytes
 */
// llgo:link (*LlT).LlInit C.lv_ll_init
func (recv_ *LlT) LlInit(node_size c.Uint32T) {
}

/**
 * Add a new head to a linked list
 * @param ll_p pointer to linked list
 * @return pointer to the new head
 */
// llgo:link (*LlT).LlInsHead C.lv_ll_ins_head
func (recv_ *LlT) LlInsHead() c.Pointer {
	return nil
}

/**
 * Insert a new node in front of the n_act node
 * @param ll_p pointer to linked list
 * @param n_act pointer a node
 * @return pointer to the new node
 */
// llgo:link (*LlT).LlInsPrev C.lv_ll_ins_prev
func (recv_ *LlT) LlInsPrev(n_act c.Pointer) c.Pointer {
	return nil
}

/**
 * Add a new tail to a linked list
 * @param ll_p pointer to linked list
 * @return pointer to the new tail
 */
// llgo:link (*LlT).LlInsTail C.lv_ll_ins_tail
func (recv_ *LlT) LlInsTail() c.Pointer {
	return nil
}

/**
 * Remove the node 'node_p' from 'll_p' linked list.
 * It does not free the memory of node.
 * @param ll_p pointer to the linked list of 'node_p'
 * @param node_p pointer to node in 'll_p' linked list
 */
// llgo:link (*LlT).LlRemove C.lv_ll_remove
func (recv_ *LlT) LlRemove(node_p c.Pointer) {
}

// llgo:link (*LlT).LlClearCustom C.lv_ll_clear_custom
func (recv_ *LlT) LlClearCustom(cleanup func(c.Pointer)) {
}

/**
 * Remove and free all elements from a linked list. The list remain valid but become empty.
 * @param ll_p pointer to linked list
 */
// llgo:link (*LlT).LlClear C.lv_ll_clear
func (recv_ *LlT) LlClear() {
}

/**
 * Move a node to a new linked list
 * @param ll_ori_p pointer to the original (old) linked list
 * @param ll_new_p pointer to the new linked list
 * @param node pointer to a node
 * @param head true: be the head in the new list
 *             false be the tail in the new list
 */
// llgo:link (*LlT).LlChgList C.lv_ll_chg_list
func (recv_ *LlT) LlChgList(ll_new_p *LlT, node c.Pointer, head bool) {
}

/**
 * Return with head node of the linked list
 * @param ll_p pointer to linked list
 * @return pointer to the head of 'll_p'
 */
// llgo:link (*LlT).LlGetHead C.lv_ll_get_head
func (recv_ *LlT) LlGetHead() c.Pointer {
	return nil
}

/**
 * Return with tail node of the linked list
 * @param ll_p pointer to linked list
 * @return pointer to the tail of 'll_p'
 */
// llgo:link (*LlT).LlGetTail C.lv_ll_get_tail
func (recv_ *LlT) LlGetTail() c.Pointer {
	return nil
}

/**
 * Return with the pointer of the next node after 'n_act'
 * @param ll_p pointer to linked list
 * @param n_act pointer a node
 * @return pointer to the next node
 */
// llgo:link (*LlT).LlGetNext C.lv_ll_get_next
func (recv_ *LlT) LlGetNext(n_act c.Pointer) c.Pointer {
	return nil
}

/**
 * Return with the pointer of the previous node after 'n_act'
 * @param ll_p pointer to linked list
 * @param n_act pointer a node
 * @return pointer to the previous node
 */
// llgo:link (*LlT).LlGetPrev C.lv_ll_get_prev
func (recv_ *LlT) LlGetPrev(n_act c.Pointer) c.Pointer {
	return nil
}

/**
 * Return the length of the linked list.
 * @param ll_p pointer to linked list
 * @return length of the linked list
 */
// llgo:link (*LlT).LlGetLen C.lv_ll_get_len
func (recv_ *LlT) LlGetLen() c.Uint32T {
	return 0
}

/**
 * Move a node before another node in the same linked list
 *
 * @param ll_p pointer to a linked list
 * @param n_act pointer to node to move
 * @param n_after pointer to a node which should be after `n_act`
 */
// llgo:link (*LlT).LlMoveBefore C.lv_ll_move_before
func (recv_ *LlT) LlMoveBefore(n_act c.Pointer, n_after c.Pointer) {
}

/**
 * Check if a linked list is empty
 * @param ll_p pointer to a linked list
 * @return true: the linked list is empty; false: not empty
 */
// llgo:link (*LlT).LlIsEmpty C.lv_ll_is_empty
func (recv_ *LlT) LlIsEmpty() bool {
	return false
}

// llgo:type C
type TimerCbT func(*TimerT)

// llgo:type C
type TimerHandlerResumeCbT func(c.Pointer)

/**
 * Call it periodically to handle lv_timers.
 * @return time till it needs to be run next (in ms)
 */
//go:linkname TimerHandler C.lv_timer_handler
func TimerHandler() c.Uint32T

/**
 * Call it in the super-loop of main() or threads. It will run lv_timer_handler()
 * with a given period in ms. You can use it with sleep or delay in OS environment.
 * This function is used to simplify the porting.
 * @param period the period for running lv_timer_handler()
 * @return the time after which it must be called again
 */
//go:linkname TimerHandlerRunInPeriod C.lv_timer_handler_run_in_period
func TimerHandlerRunInPeriod(period c.Uint32T) c.Uint32T

/**
 * Call it in the super-loop of main() or threads. It will automatically call lv_timer_handler() at the right time.
 * This function is used to simplify the porting.
 */
//go:linkname TimerPeriodicHandler C.lv_timer_periodic_handler
func TimerPeriodicHandler()

/**
 * Set the resume callback to the timer handler
 * @param cb the function to call when timer handler is resumed
 * @param data pointer to a resume data
 */
//go:linkname TimerHandlerSetResumeCb C.lv_timer_handler_set_resume_cb
func TimerHandlerSetResumeCb(cb TimerHandlerResumeCbT, data c.Pointer)

/**
 * Create an "empty" timer. It needs to be initialized with at least
 * `lv_timer_set_cb` and `lv_timer_set_period`
 * @return pointer to the created timer
 */
//go:linkname TimerCreateBasic C.lv_timer_create_basic
func TimerCreateBasic() *TimerT

/**
 * Create a new lv_timer
 * @param timer_xcb a callback to call periodically.
 *                 (the 'x' in the argument name indicates that it's not a fully generic function because it not follows
 *                  the `func_name(object, callback, ...)` convention)
 * @param period call period in ms unit
 * @param user_data custom parameter
 * @return pointer to the new timer
 */
//go:linkname TimerCreate C.lv_timer_create
func TimerCreate(timer_xcb TimerCbT, period c.Uint32T, user_data c.Pointer) *TimerT

/**
 * Delete a lv_timer
 * @param timer pointer to an lv_timer
 */
// llgo:link (*TimerT).TimerDelete C.lv_timer_delete
func (recv_ *TimerT) TimerDelete() {
}

/**
 * Pause a timer.
 * @param timer pointer to an lv_timer
 */
// llgo:link (*TimerT).TimerPause C.lv_timer_pause
func (recv_ *TimerT) TimerPause() {
}

/**
 * Resume a timer.
 * @param timer pointer to an lv_timer
 */
// llgo:link (*TimerT).TimerResume C.lv_timer_resume
func (recv_ *TimerT) TimerResume() {
}

/**
 * Set the callback to the timer (the function to call periodically)
 * @param timer pointer to a timer
 * @param timer_cb the function to call periodically
 */
// llgo:link (*TimerT).TimerSetCb C.lv_timer_set_cb
func (recv_ *TimerT) TimerSetCb(timer_cb TimerCbT) {
}

/**
 * Set new period for a lv_timer
 * @param timer pointer to a lv_timer
 * @param period the new period
 */
// llgo:link (*TimerT).TimerSetPeriod C.lv_timer_set_period
func (recv_ *TimerT) TimerSetPeriod(period c.Uint32T) {
}

/**
 * Make a lv_timer ready. It will not wait its period.
 * @param timer pointer to a lv_timer.
 */
// llgo:link (*TimerT).TimerReady C.lv_timer_ready
func (recv_ *TimerT) TimerReady() {
}

/**
 * Set the number of times a timer will repeat.
 * @param timer pointer to a lv_timer.
 * @param repeat_count -1 : infinity;  0 : stop ;  n>0: residual times
 */
// llgo:link (*TimerT).TimerSetRepeatCount C.lv_timer_set_repeat_count
func (recv_ *TimerT) TimerSetRepeatCount(repeat_count c.Int32T) {
}

/**
 * Set whether a lv_timer will be deleted automatically when it is called `repeat_count` times.
 * @param timer pointer to a lv_timer.
 * @param auto_delete true: auto delete; false: timer will be paused when it is called `repeat_count` times.
 */
// llgo:link (*TimerT).TimerSetAutoDelete C.lv_timer_set_auto_delete
func (recv_ *TimerT) TimerSetAutoDelete(auto_delete bool) {
}

/**
 * Set custom parameter to the lv_timer.
 * @param timer pointer to a lv_timer.
 * @param user_data custom parameter
 */
// llgo:link (*TimerT).TimerSetUserData C.lv_timer_set_user_data
func (recv_ *TimerT) TimerSetUserData(user_data c.Pointer) {
}

/**
 * Reset a lv_timer.
 * It will be called the previously set period milliseconds later.
 * @param timer pointer to a lv_timer.
 */
// llgo:link (*TimerT).TimerReset C.lv_timer_reset
func (recv_ *TimerT) TimerReset() {
}

/**
 * Enable or disable the whole lv_timer handling
 * @param en true: lv_timer handling is running, false: lv_timer handling is suspended
 */
//go:linkname TimerEnable C.lv_timer_enable
func TimerEnable(en bool)

/**
 * Get idle percentage
 * @return the lv_timer idle in percentage
 */
//go:linkname TimerGetIdle C.lv_timer_get_idle
func TimerGetIdle() c.Uint32T

/**
 * Get the time remaining until the next timer will run
 * @return the time remaining in ms
 */
//go:linkname TimerGetTimeUntilNext C.lv_timer_get_time_until_next
func TimerGetTimeUntilNext() c.Uint32T

/**
 * Iterate through the timers
 * @param timer NULL to start iteration or the previous return value to get the next timer
 * @return the next timer or NULL if there is no more timer
 */
// llgo:link (*TimerT).TimerGetNext C.lv_timer_get_next
func (recv_ *TimerT) TimerGetNext() *TimerT {
	return nil
}

/**
 * Get the user_data passed when the timer was created
 * @param timer pointer to the lv_timer
 * @return pointer to the user_data
 */
// llgo:link (*TimerT).TimerGetUserData C.lv_timer_get_user_data
func (recv_ *TimerT) TimerGetUserData() c.Pointer {
	return nil
}

/**
 * Get the pause state of a timer
 * @param timer pointer to a lv_timer
 * @return true: timer is paused; false: timer is running
 */
// llgo:link (*TimerT).TimerGetPaused C.lv_timer_get_paused
func (recv_ *TimerT) TimerGetPaused() bool {
	return false
}

type AnimEnableT bool

// llgo:type C
type AnimPathCbT func(*AnimT) c.Int32T

// llgo:type C
type AnimExecXcbT func(c.Pointer, c.Int32T)

// llgo:type C
type AnimCustomExecCbT func(*AnimT, c.Int32T)

// llgo:type C
type AnimCompletedCbT func(*AnimT)

// llgo:type C
type AnimStartCbT func(*AnimT)

// llgo:type C
type AnimGetValueCbT func(*AnimT) c.Int32T

// llgo:type C
type AnimDeletedCbT func(*AnimT)

/** Parameter used when path is custom_bezier */

type AnimBezier3ParaT struct {
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
// llgo:link (*AnimT).AnimInit C.lv_anim_init
func (recv_ *AnimT) AnimInit() {
}

/**
 * Set a variable to animate
 * @param a     pointer to an initialized `lv_anim_t` variable
 * @param var   pointer to a variable to animate
 */
// llgo:link (*AnimT).AnimSetVar C.lv_anim_set_var
func (recv_ *AnimT) AnimSetVar(var_ c.Pointer) {
}

/**
 * Set a function to animate `var`
 * @param a         pointer to an initialized `lv_anim_t` variable
 * @param exec_cb   a function to execute during animation
 *                  LVGL's built-in functions can be used.
 *                  E.g. lv_obj_set_x
 */
// llgo:link (*AnimT).AnimSetExecCb C.lv_anim_set_exec_cb
func (recv_ *AnimT) AnimSetExecCb(exec_cb AnimExecXcbT) {
}

/**
 * Set the duration of an animation
 * @param a         pointer to an initialized `lv_anim_t` variable
 * @param duration  duration of the animation in milliseconds
 */
// llgo:link (*AnimT).AnimSetDuration C.lv_anim_set_duration
func (recv_ *AnimT) AnimSetDuration(duration c.Uint32T) {
}

/**
 * Set a delay before starting the animation
 * @param a         pointer to an initialized `lv_anim_t` variable
 * @param delay     delay before the animation in milliseconds
 */
// llgo:link (*AnimT).AnimSetDelay C.lv_anim_set_delay
func (recv_ *AnimT) AnimSetDelay(delay c.Uint32T) {
}

/**
 * Resumes a paused animation
 * @param a         pointer to an initialized `lv_anim_t` variable
 */
// llgo:link (*AnimT).AnimResume C.lv_anim_resume
func (recv_ *AnimT) AnimResume() {
}

/**
 * Pauses the animation
 * @param a         pointer to an initialized `lv_anim_t` variable
 */
// llgo:link (*AnimT).AnimPause C.lv_anim_pause
func (recv_ *AnimT) AnimPause() {
}

/**
 * Pauses the animation for ms milliseconds
 * @param a         pointer to an initialized `lv_anim_t` variable
 * @param ms        the pause time in milliseconds
 */
// llgo:link (*AnimT).AnimPauseFor C.lv_anim_pause_for
func (recv_ *AnimT) AnimPauseFor(ms c.Uint32T) {
}

/**
 * Check if the animation is paused
 * @param a         pointer to an initialized `lv_anim_t` variable
 * @return          true if the animation is paused else false
 */
// llgo:link (*AnimT).AnimIsPaused C.lv_anim_is_paused
func (recv_ *AnimT) AnimIsPaused() bool {
	return false
}

/**
 * Set the start and end values of an animation
 * @param a         pointer to an initialized `lv_anim_t` variable
 * @param start     the start value
 * @param end       the end value
 */
// llgo:link (*AnimT).AnimSetValues C.lv_anim_set_values
func (recv_ *AnimT) AnimSetValues(start c.Int32T, end c.Int32T) {
}

/**
 * Similar to `lv_anim_set_exec_cb` but `lv_anim_custom_exec_cb_t` receives
 * `lv_anim_t * ` as its first parameter instead of `void *`.
 * This function might be used when LVGL is bound to other languages because
 * it's more consistent to have `lv_anim_t *` as first parameter.
 * @param a         pointer to an initialized `lv_anim_t` variable
 * @param exec_cb   a function to execute.
 */
// llgo:link (*AnimT).AnimSetCustomExecCb C.lv_anim_set_custom_exec_cb
func (recv_ *AnimT) AnimSetCustomExecCb(exec_cb AnimCustomExecCbT) {
}

/**
 * Set the path (curve) of the animation.
 * @param a         pointer to an initialized `lv_anim_t` variable
 * @param path_cb a function to set the current value of the animation.
 */
// llgo:link (*AnimT).AnimSetPathCb C.lv_anim_set_path_cb
func (recv_ *AnimT) AnimSetPathCb(path_cb AnimPathCbT) {
}

/**
 * Set a function call when the animation really starts (considering `delay`)
 * @param a         pointer to an initialized `lv_anim_t` variable
 * @param start_cb  a function call when the animation starts
 */
// llgo:link (*AnimT).AnimSetStartCb C.lv_anim_set_start_cb
func (recv_ *AnimT) AnimSetStartCb(start_cb AnimStartCbT) {
}

/**
 * Set a function to use the current value of the variable and make start and end value
 * relative to the returned current value.
 * @param a             pointer to an initialized `lv_anim_t` variable
 * @param get_value_cb  a function call when the animation starts
 */
// llgo:link (*AnimT).AnimSetGetValueCb C.lv_anim_set_get_value_cb
func (recv_ *AnimT) AnimSetGetValueCb(get_value_cb AnimGetValueCbT) {
}

/**
 * Set a function call when the animation is completed
 * @param a             pointer to an initialized `lv_anim_t` variable
 * @param completed_cb  a function call when the animation is fully completed
 */
// llgo:link (*AnimT).AnimSetCompletedCb C.lv_anim_set_completed_cb
func (recv_ *AnimT) AnimSetCompletedCb(completed_cb AnimCompletedCbT) {
}

/**
 * Set a function call when the animation is deleted.
 * @param a         pointer to an initialized `lv_anim_t` variable
 * @param deleted_cb  a function call when the animation is deleted
 */
// llgo:link (*AnimT).AnimSetDeletedCb C.lv_anim_set_deleted_cb
func (recv_ *AnimT) AnimSetDeletedCb(deleted_cb AnimDeletedCbT) {
}

/**
 * Make the animation to play back to when the forward direction is ready
 * @param a         pointer to an initialized `lv_anim_t` variable
 * @param duration  duration of playback animation in milliseconds. 0: disable playback
 */
// llgo:link (*AnimT).AnimSetReverseDuration C.lv_anim_set_reverse_duration
func (recv_ *AnimT) AnimSetReverseDuration(duration c.Uint32T) {
}

/**
 * Legacy `lv_anim_set_reverse_time` API will be removed soon, use `lv_anim_set_reverse_duration` instead.
 */
// llgo:link (*AnimT).AnimSetReverseTime C.lv_anim_set_reverse_time
func (recv_ *AnimT) AnimSetReverseTime(duration c.Uint32T) {
}

/**
 * Make the animation to play back to when the forward direction is ready
 * @param a         pointer to an initialized `lv_anim_t` variable
 * @param delay     delay in milliseconds before starting the playback animation.
 */
// llgo:link (*AnimT).AnimSetReverseDelay C.lv_anim_set_reverse_delay
func (recv_ *AnimT) AnimSetReverseDelay(delay c.Uint32T) {
}

/**
 * Make the animation repeat itself.
 * @param a         pointer to an initialized `lv_anim_t` variable
 * @param cnt       repeat count or `LV_ANIM_REPEAT_INFINITE` for infinite repetition. 0: to disable repetition.
 */
// llgo:link (*AnimT).AnimSetRepeatCount C.lv_anim_set_repeat_count
func (recv_ *AnimT) AnimSetRepeatCount(cnt c.Uint32T) {
}

/**
 * Set a delay before repeating the animation.
 * @param a         pointer to an initialized `lv_anim_t` variable
 * @param delay     delay in milliseconds before repeating the animation.
 */
// llgo:link (*AnimT).AnimSetRepeatDelay C.lv_anim_set_repeat_delay
func (recv_ *AnimT) AnimSetRepeatDelay(delay c.Uint32T) {
}

/**
 * Set a whether the animation's should be applied immediately or only when the delay expired.
 * @param a         pointer to an initialized `lv_anim_t` variable
 * @param en        true: apply the start value immediately in `lv_anim_start`;
 *                  false: apply the start value only when `delay` ms is elapsed and the animations really starts
 */
// llgo:link (*AnimT).AnimSetEarlyApply C.lv_anim_set_early_apply
func (recv_ *AnimT) AnimSetEarlyApply(en bool) {
}

/**
 * Set the custom user data field of the animation.
 * @param a           pointer to an initialized `lv_anim_t` variable
 * @param user_data   pointer to the new user_data.
 */
// llgo:link (*AnimT).AnimSetUserData C.lv_anim_set_user_data
func (recv_ *AnimT) AnimSetUserData(user_data c.Pointer) {
}

/**
 * Set parameter for cubic bezier path
 * @param a         pointer to an initialized `lv_anim_t` variable
 * @param x1        first control point X
 * @param y1        first control point Y
 * @param x2        second control point X
 * @param y2        second control point Y
 */
// llgo:link (*AnimT).AnimSetBezier3Param C.lv_anim_set_bezier3_param
func (recv_ *AnimT) AnimSetBezier3Param(x1 c.Int16T, y1 c.Int16T, x2 c.Int16T, y2 c.Int16T) {
}

/**
 * Create an animation
 * @param a         an initialized 'anim_t' variable. Not required after call.
 * @return          pointer to the created animation (different from the `a` parameter)
 */
// llgo:link (*AnimT).AnimStart C.lv_anim_start
func (recv_ *AnimT) AnimStart() *AnimT {
	return nil
}

/**
 * Get a delay before starting the animation
 * @param a pointer to an initialized `lv_anim_t` variable
 * @return delay before the animation in milliseconds
 */
// llgo:link (*AnimT).AnimGetDelay C.lv_anim_get_delay
func (recv_ *AnimT) AnimGetDelay() c.Uint32T {
	return 0
}

/**
 * Get the time used to play the animation.
 * @param a pointer to an animation.
 * @return the play time in milliseconds.
 */
// llgo:link (*AnimT).AnimGetPlaytime C.lv_anim_get_playtime
func (recv_ *AnimT) AnimGetPlaytime() c.Uint32T {
	return 0
}

/**
 * Get the duration of an animation
 * @param a         pointer to an initialized `lv_anim_t` variable
 * @return the duration of the animation in milliseconds
 */
// llgo:link (*AnimT).AnimGetTime C.lv_anim_get_time
func (recv_ *AnimT) AnimGetTime() c.Uint32T {
	return 0
}

/**
 * Get the repeat count of the animation.
 * @param a         pointer to an initialized `lv_anim_t` variable
 * @return the repeat count or `LV_ANIM_REPEAT_INFINITE` for infinite repetition. 0: disabled repetition.
 */
// llgo:link (*AnimT).AnimGetRepeatCount C.lv_anim_get_repeat_count
func (recv_ *AnimT) AnimGetRepeatCount() c.Uint32T {
	return 0
}

/**
 * Get the user_data field of the animation
 * @param   a pointer to an initialized `lv_anim_t` variable
 * @return  the pointer to the custom user_data of the animation
 */
// llgo:link (*AnimT).AnimGetUserData C.lv_anim_get_user_data
func (recv_ *AnimT) AnimGetUserData() c.Pointer {
	return nil
}

/**
 * Delete animation(s) of a variable with a given animator function
 * @param var       pointer to variable
 * @param exec_cb   a function pointer which is animating 'var',
 *                  or NULL to ignore it and delete all the animations of 'var
 * @return          true: at least 1 animation is deleted, false: no animation is deleted
 */
//go:linkname AnimDelete C.lv_anim_delete
func AnimDelete(var_ c.Pointer, exec_cb AnimExecXcbT) bool

/**
 * Delete all the animations
 */
//go:linkname AnimDeleteAll C.lv_anim_delete_all
func AnimDeleteAll()

/**
 * Get the animation of a variable and its `exec_cb`.
 * @param var       pointer to variable
 * @param exec_cb   a function pointer which is animating 'var', or NULL to return first matching 'var'
 * @return          pointer to the animation.
 */
//go:linkname AnimGet C.lv_anim_get
func AnimGet(var_ c.Pointer, exec_cb AnimExecXcbT) *AnimT

/**
 * Get global animation refresher timer.
 * @return pointer to the animation refresher timer.
 */
//go:linkname AnimGetTimer C.lv_anim_get_timer
func AnimGetTimer() *TimerT

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
// llgo:link (*AnimT).AnimCustomDelete C.lv_anim_custom_delete
func (recv_ *AnimT) AnimCustomDelete(exec_cb AnimCustomExecCbT) bool {
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
// llgo:link (*AnimT).AnimCustomGet C.lv_anim_custom_get
func (recv_ *AnimT) AnimCustomGet(exec_cb AnimCustomExecCbT) *AnimT {
	return nil
}

/**
 * Get the number of currently running animations
 * @return      the number of running animations
 */
//go:linkname AnimCountRunning C.lv_anim_count_running
func AnimCountRunning() c.Uint16T

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
//go:linkname AnimSpeed C.lv_anim_speed
func AnimSpeed(speed c.Uint32T) c.Uint32T

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
//go:linkname AnimSpeedClamped C.lv_anim_speed_clamped
func AnimSpeedClamped(speed c.Uint32T, min_time c.Uint32T, max_time c.Uint32T) c.Uint32T

/**
 * Resolve the speed (created with `lv_anim_speed` or `lv_anim_speed_clamped`) to time
 * based on start and end values.
 * @param speed     return values of `lv_anim_speed` or `lv_anim_speed_clamped`
 * @param start     the start value of the animation
 * @param end       the end value of the animation
 * @return          the time required to get from `start` to `end` with the given `speed` setting
 */
//go:linkname AnimResolveSpeed C.lv_anim_resolve_speed
func AnimResolveSpeed(speed c.Uint32T, start c.Int32T, end c.Int32T) c.Uint32T

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
//go:linkname AnimSpeedToTime C.lv_anim_speed_to_time
func AnimSpeedToTime(speed c.Uint32T, start c.Int32T, end c.Int32T) c.Uint32T

/**
 * Manually refresh the state of the animations.
 * Useful to make the animations running in a blocking process where
 * `lv_timer_handler` can't run for a while.
 * Shouldn't be used directly because it is called in `lv_refr_now()`.
 */
//go:linkname AnimRefrNow C.lv_anim_refr_now
func AnimRefrNow()

/**
 * Calculate the current value of an animation applying linear characteristic
 * @param a     pointer to an animation
 * @return      the current value to set
 */
// llgo:link (*AnimT).AnimPathLinear C.lv_anim_path_linear
func (recv_ *AnimT) AnimPathLinear() c.Int32T {
	return 0
}

/**
 * Calculate the current value of an animation slowing down the start phase
 * @param a     pointer to an animation
 * @return      the current value to set
 */
// llgo:link (*AnimT).AnimPathEaseIn C.lv_anim_path_ease_in
func (recv_ *AnimT) AnimPathEaseIn() c.Int32T {
	return 0
}

/**
 * Calculate the current value of an animation slowing down the end phase
 * @param a     pointer to an animation
 * @return      the current value to set
 */
// llgo:link (*AnimT).AnimPathEaseOut C.lv_anim_path_ease_out
func (recv_ *AnimT) AnimPathEaseOut() c.Int32T {
	return 0
}

/**
 * Calculate the current value of an animation applying an "S" characteristic (cosine)
 * @param a     pointer to an animation
 * @return      the current value to set
 */
// llgo:link (*AnimT).AnimPathEaseInOut C.lv_anim_path_ease_in_out
func (recv_ *AnimT) AnimPathEaseInOut() c.Int32T {
	return 0
}

/**
 * Calculate the current value of an animation with overshoot at the end
 * @param a     pointer to an animation
 * @return      the current value to set
 */
// llgo:link (*AnimT).AnimPathOvershoot C.lv_anim_path_overshoot
func (recv_ *AnimT) AnimPathOvershoot() c.Int32T {
	return 0
}

/**
 * Calculate the current value of an animation with 3 bounces
 * @param a     pointer to an animation
 * @return      the current value to set
 */
// llgo:link (*AnimT).AnimPathBounce C.lv_anim_path_bounce
func (recv_ *AnimT) AnimPathBounce() c.Int32T {
	return 0
}

/**
 * Calculate the current value of an animation applying step characteristic.
 * (Set end value on the end of the animation)
 * @param a     pointer to an animation
 * @return      the current value to set
 */
// llgo:link (*AnimT).AnimPathStep C.lv_anim_path_step
func (recv_ *AnimT) AnimPathStep() c.Int32T {
	return 0
}

/**
 * A custom cubic bezier animation path, need to specify cubic-parameters in a->parameter.bezier3
 * @param a     pointer to an animation
 * @return      the current value to set
 */
// llgo:link (*AnimT).AnimPathCustomBezier3 C.lv_anim_path_custom_bezier3
func (recv_ *AnimT) AnimPathCustomBezier3() c.Int32T {
	return 0
}

//go:linkname Snprintf C.lv_snprintf
func Snprintf(buffer *c.Char, count c.SizeT, format *c.Char, __llgo_va_list ...interface{}) c.Int

//go:linkname Vsnprintf C.lv_vsnprintf
func Vsnprintf(buffer *c.Char, count c.SizeT, format *c.Char, va c.VaList) c.Int

type TextFlagT c.Int

const (
	TEXT_FLAG_NONE      TextFlagT = 0
	TEXT_FLAG_EXPAND    TextFlagT = 1
	TEXT_FLAG_FIT       TextFlagT = 2
	TEXT_FLAG_BREAK_ALL TextFlagT = 4
	TEXT_FLAG_RECOLOR   TextFlagT = 8
)

type TextAlignT c.Int

const (
	TEXT_ALIGN_AUTO   TextAlignT = 0
	TEXT_ALIGN_LEFT   TextAlignT = 1
	TEXT_ALIGN_CENTER TextAlignT = 2
	TEXT_ALIGN_RIGHT  TextAlignT = 3
)

type TextCmdStateT c.Int

const (
	TEXT_CMD_STATE_WAIT TextCmdStateT = 0
	TEXT_CMD_STATE_PAR  TextCmdStateT = 1
	TEXT_CMD_STATE_IN   TextCmdStateT = 2
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
// llgo:link (*PointT).TextGetSize C.lv_text_get_size
func (recv_ *PointT) TextGetSize(text *c.Char, font *FontT, letter_space c.Int32T, line_space c.Int32T, max_width c.Int32T, flag TextFlagT) {
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
//go:linkname TextGetWidth C.lv_text_get_width
func TextGetWidth(txt *c.Char, length c.Uint32T, font *FontT, letter_space c.Int32T) c.Int32T

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
//go:linkname TextGetWidthWithFlags C.lv_text_get_width_with_flags
func TextGetWidthWithFlags(txt *c.Char, length c.Uint32T, font *FontT, letter_space c.Int32T, flags TextFlagT) c.Int32T

/**
 * Check if c is command state
 * @param state
 * @param c
 * @return True if c is state
 */
// llgo:link (*TextCmdStateT).TextIsCmd C.lv_text_is_cmd
func (recv_ *TextCmdStateT) TextIsCmd(c c.Uint32T) bool {
	return false
}

type BaseDirT c.Int

const (
	BASE_DIR_LTR     BaseDirT = 0
	BASE_DIR_RTL     BaseDirT = 1
	BASE_DIR_AUTO    BaseDirT = 2
	BASE_DIR_NEUTRAL BaseDirT = 32
	BASE_DIR_WEAK    BaseDirT = 33
)

type GradDirT c.Int

const (
	GRAD_DIR_NONE    GradDirT = 0
	GRAD_DIR_VER     GradDirT = 1
	GRAD_DIR_HOR     GradDirT = 2
	GRAD_DIR_LINEAR  GradDirT = 3
	GRAD_DIR_RADIAL  GradDirT = 4
	GRAD_DIR_CONICAL GradDirT = 5
)

type GradExtendT c.Int

const (
	GRAD_EXTEND_PAD     GradExtendT = 0
	GRAD_EXTEND_REPEAT  GradExtendT = 1
	GRAD_EXTEND_REFLECT GradExtendT = 2
)

/** A gradient stop definition.
 *  This matches a color and a position in a virtual 0-255 scale.
 */

type GradStopT struct {
	Color ColorT
	Opa   OpaT
	Frac  c.Uint8T
}

/** A descriptor of a gradient. */

type GradDscT struct {
	Stops      [2]GradStopT
	StopsCount c.Uint8T
	Dir        GradDirT
	Extend     GradExtendT
}

/**
 * Initialize gradient color map from a table
 * @param grad      pointer to a gradient descriptor
 * @param colors    color array
 * @param fracs     position array (0..255): if NULL, then colors are distributed evenly
 * @param opa       opacity array: if NULL, then LV_OPA_COVER is assumed
 * @param num_stops number of gradient stops (1..LV_GRADIENT_MAX_STOPS)
 */
// llgo:link (*GradDscT).GradInitStops C.lv_grad_init_stops
func (recv_ *GradDscT) GradInitStops(colors *ColorT, opa *OpaT, fracs *c.Uint8T, num_stops c.Int) {
}

/**
 * Helper function to initialize a horizontal gradient.
 * @param dsc      gradient descriptor
 */
// llgo:link (*GradDscT).GradHorizontalInit C.lv_grad_horizontal_init
func (recv_ *GradDscT) GradHorizontalInit() {
}

/**
 * Helper function to initialize a vertical gradient.
 * @param dsc      gradient descriptor
 */
// llgo:link (*GradDscT).GradVerticalInit C.lv_grad_vertical_init
func (recv_ *GradDscT) GradVerticalInit() {
}

// llgo:type C
type LayoutUpdateCbT func(*ObjT, c.Pointer)
type LayoutT c.Int

const (
	LAYOUT_NONE LayoutT = 0
	LAYOUT_FLEX LayoutT = 1
	LAYOUT_GRID LayoutT = 2
	LAYOUT_LAST LayoutT = 3
)

/**
 * Register a new layout
 * @param cb        the layout update callback
 * @param user_data custom data that will be passed to `cb`
 * @return          the ID of the new layout
 */
//go:linkname LayoutRegister C.lv_layout_register
func LayoutRegister(cb LayoutUpdateCbT, user_data c.Pointer) c.Uint32T

type FlexAlignT c.Int

const (
	FLEX_ALIGN_START         FlexAlignT = 0
	FLEX_ALIGN_END           FlexAlignT = 1
	FLEX_ALIGN_CENTER        FlexAlignT = 2
	FLEX_ALIGN_SPACE_EVENLY  FlexAlignT = 3
	FLEX_ALIGN_SPACE_AROUND  FlexAlignT = 4
	FLEX_ALIGN_SPACE_BETWEEN FlexAlignT = 5
)

type FlexFlowT c.Int

const (
	FLEX_FLOW_ROW                 FlexFlowT = 0
	FLEX_FLOW_COLUMN              FlexFlowT = 1
	FLEX_FLOW_ROW_WRAP            FlexFlowT = 4
	FLEX_FLOW_ROW_REVERSE         FlexFlowT = 8
	FLEX_FLOW_ROW_WRAP_REVERSE    FlexFlowT = 12
	FLEX_FLOW_COLUMN_WRAP         FlexFlowT = 5
	FLEX_FLOW_COLUMN_REVERSE      FlexFlowT = 9
	FLEX_FLOW_COLUMN_WRAP_REVERSE FlexFlowT = 13
)

/**
 * Initialize a flex layout to default values
 */
//go:linkname FlexInit C.lv_flex_init
func FlexInit()

/**
 * Set how the item should flow
 * @param obj pointer to an object. The parent must have flex layout else nothing will happen.
 * @param flow an element of `lv_flex_flow_t`.
 */
// llgo:link (*ObjT).ObjSetFlexFlow C.lv_obj_set_flex_flow
func (recv_ *ObjT) ObjSetFlexFlow(flow FlexFlowT) {
}

/**
 * Set how to place (where to align) the items and tracks
 * @param obj pointer to an object. The parent must have flex layout else nothing will happen.
 * @param main_place where to place the items on main axis (in their track). Any value of `lv_flex_align_t`.
 * @param cross_place where to place the item in their track on the cross axis. `LV_FLEX_ALIGN_START/END/CENTER`
 * @param track_cross_place where to place the tracks in the cross direction. Any value of `lv_flex_align_t`.
 */
// llgo:link (*ObjT).ObjSetFlexAlign C.lv_obj_set_flex_align
func (recv_ *ObjT) ObjSetFlexAlign(main_place FlexAlignT, cross_place FlexAlignT, track_cross_place FlexAlignT) {
}

/**
 * Sets the width or height (on main axis) to grow the object in order fill the free space
 * @param obj pointer to an object. The parent must have flex layout else nothing will happen.
 * @param grow a value to set how much free space to take proportionally to other growing items.
 */
// llgo:link (*ObjT).ObjSetFlexGrow C.lv_obj_set_flex_grow
func (recv_ *ObjT) ObjSetFlexGrow(grow c.Uint8T) {
}

type GridAlignT c.Int

const (
	GRID_ALIGN_START         GridAlignT = 0
	GRID_ALIGN_CENTER        GridAlignT = 1
	GRID_ALIGN_END           GridAlignT = 2
	GRID_ALIGN_STRETCH       GridAlignT = 3
	GRID_ALIGN_SPACE_EVENLY  GridAlignT = 4
	GRID_ALIGN_SPACE_AROUND  GridAlignT = 5
	GRID_ALIGN_SPACE_BETWEEN GridAlignT = 6
)

/**********************
 * GLOBAL PROTOTYPES
 **********************/
//go:linkname GridInit C.lv_grid_init
func GridInit()

// llgo:link (*ObjT).ObjSetGridDscArray C.lv_obj_set_grid_dsc_array
func (recv_ *ObjT) ObjSetGridDscArray(col_dsc *c.Int32T, row_dsc *c.Int32T) {
}

// llgo:link (*ObjT).ObjSetGridAlign C.lv_obj_set_grid_align
func (recv_ *ObjT) ObjSetGridAlign(column_align GridAlignT, row_align GridAlignT) {
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
// llgo:link (*ObjT).ObjSetGridCell C.lv_obj_set_grid_cell
func (recv_ *ObjT) ObjSetGridCell(column_align GridAlignT, col_pos c.Int32T, col_span c.Int32T, row_align GridAlignT, row_pos c.Int32T, row_span c.Int32T) {
}

/**
 * Just a wrapper to `LV_GRID_FR` for bindings.
 */
//go:linkname GridFr C.lv_grid_fr
func GridFr(x c.Uint8T) c.Int32T

type BlendModeT c.Int

const (
	BLEND_MODE_NORMAL      BlendModeT = 0
	BLEND_MODE_ADDITIVE    BlendModeT = 1
	BLEND_MODE_SUBTRACTIVE BlendModeT = 2
	BLEND_MODE_MULTIPLY    BlendModeT = 3
	BLEND_MODE_DIFFERENCE  BlendModeT = 4
)

type TextDecorT c.Int

const (
	TEXT_DECOR_NONE          TextDecorT = 0
	TEXT_DECOR_UNDERLINE     TextDecorT = 1
	TEXT_DECOR_STRIKETHROUGH TextDecorT = 2
)

type BorderSideT c.Int

const (
	BORDER_SIDE_NONE     BorderSideT = 0
	BORDER_SIDE_BOTTOM   BorderSideT = 1
	BORDER_SIDE_TOP      BorderSideT = 2
	BORDER_SIDE_LEFT     BorderSideT = 4
	BORDER_SIDE_RIGHT    BorderSideT = 8
	BORDER_SIDE_FULL     BorderSideT = 15
	BORDER_SIDE_INTERNAL BorderSideT = 16
)

/**
 * A common type to handle all the property types in the same way.
 */

type StyleValueT struct {
	Ptr c.Pointer
}
type X_lvStyleIdT c.Int

const (
	STYLE_PROP_INV                  X_lvStyleIdT = 0
	STYLE_WIDTH                     X_lvStyleIdT = 1
	STYLE_HEIGHT                    X_lvStyleIdT = 2
	STYLE_LENGTH                    X_lvStyleIdT = 3
	STYLE_MIN_WIDTH                 X_lvStyleIdT = 4
	STYLE_MAX_WIDTH                 X_lvStyleIdT = 5
	STYLE_MIN_HEIGHT                X_lvStyleIdT = 6
	STYLE_MAX_HEIGHT                X_lvStyleIdT = 7
	STYLE_X                         X_lvStyleIdT = 8
	STYLE_Y                         X_lvStyleIdT = 9
	STYLE_ALIGN                     X_lvStyleIdT = 10
	STYLE_RADIUS                    X_lvStyleIdT = 12
	STYLE_RADIAL_OFFSET             X_lvStyleIdT = 13
	STYLE_PAD_RADIAL                X_lvStyleIdT = 14
	STYLE_PAD_TOP                   X_lvStyleIdT = 16
	STYLE_PAD_BOTTOM                X_lvStyleIdT = 17
	STYLE_PAD_LEFT                  X_lvStyleIdT = 18
	STYLE_PAD_RIGHT                 X_lvStyleIdT = 19
	STYLE_PAD_ROW                   X_lvStyleIdT = 20
	STYLE_PAD_COLUMN                X_lvStyleIdT = 21
	STYLE_LAYOUT                    X_lvStyleIdT = 22
	STYLE_MARGIN_TOP                X_lvStyleIdT = 24
	STYLE_MARGIN_BOTTOM             X_lvStyleIdT = 25
	STYLE_MARGIN_LEFT               X_lvStyleIdT = 26
	STYLE_MARGIN_RIGHT              X_lvStyleIdT = 27
	STYLE_BG_COLOR                  X_lvStyleIdT = 28
	STYLE_BG_OPA                    X_lvStyleIdT = 29
	STYLE_BG_GRAD_DIR               X_lvStyleIdT = 32
	STYLE_BG_MAIN_STOP              X_lvStyleIdT = 33
	STYLE_BG_GRAD_STOP              X_lvStyleIdT = 34
	STYLE_BG_GRAD_COLOR             X_lvStyleIdT = 35
	STYLE_BG_MAIN_OPA               X_lvStyleIdT = 36
	STYLE_BG_GRAD_OPA               X_lvStyleIdT = 37
	STYLE_BG_GRAD                   X_lvStyleIdT = 38
	STYLE_BASE_DIR                  X_lvStyleIdT = 39
	STYLE_BG_IMAGE_SRC              X_lvStyleIdT = 40
	STYLE_BG_IMAGE_OPA              X_lvStyleIdT = 41
	STYLE_BG_IMAGE_RECOLOR          X_lvStyleIdT = 42
	STYLE_BG_IMAGE_RECOLOR_OPA      X_lvStyleIdT = 43
	STYLE_BG_IMAGE_TILED            X_lvStyleIdT = 44
	STYLE_CLIP_CORNER               X_lvStyleIdT = 45
	STYLE_BORDER_WIDTH              X_lvStyleIdT = 48
	STYLE_BORDER_COLOR              X_lvStyleIdT = 49
	STYLE_BORDER_OPA                X_lvStyleIdT = 50
	STYLE_BORDER_SIDE               X_lvStyleIdT = 52
	STYLE_BORDER_POST               X_lvStyleIdT = 53
	STYLE_OUTLINE_WIDTH             X_lvStyleIdT = 56
	STYLE_OUTLINE_COLOR             X_lvStyleIdT = 57
	STYLE_OUTLINE_OPA               X_lvStyleIdT = 58
	STYLE_OUTLINE_PAD               X_lvStyleIdT = 59
	STYLE_SHADOW_WIDTH              X_lvStyleIdT = 60
	STYLE_SHADOW_COLOR              X_lvStyleIdT = 61
	STYLE_SHADOW_OPA                X_lvStyleIdT = 62
	STYLE_SHADOW_OFFSET_X           X_lvStyleIdT = 64
	STYLE_SHADOW_OFFSET_Y           X_lvStyleIdT = 65
	STYLE_SHADOW_SPREAD             X_lvStyleIdT = 66
	STYLE_IMAGE_OPA                 X_lvStyleIdT = 68
	STYLE_IMAGE_RECOLOR             X_lvStyleIdT = 69
	STYLE_IMAGE_RECOLOR_OPA         X_lvStyleIdT = 70
	STYLE_LINE_WIDTH                X_lvStyleIdT = 72
	STYLE_LINE_DASH_WIDTH           X_lvStyleIdT = 73
	STYLE_LINE_DASH_GAP             X_lvStyleIdT = 74
	STYLE_LINE_ROUNDED              X_lvStyleIdT = 75
	STYLE_LINE_COLOR                X_lvStyleIdT = 76
	STYLE_LINE_OPA                  X_lvStyleIdT = 77
	STYLE_ARC_WIDTH                 X_lvStyleIdT = 80
	STYLE_ARC_ROUNDED               X_lvStyleIdT = 81
	STYLE_ARC_COLOR                 X_lvStyleIdT = 82
	STYLE_ARC_OPA                   X_lvStyleIdT = 83
	STYLE_ARC_IMAGE_SRC             X_lvStyleIdT = 84
	STYLE_TEXT_COLOR                X_lvStyleIdT = 88
	STYLE_TEXT_OPA                  X_lvStyleIdT = 89
	STYLE_TEXT_FONT                 X_lvStyleIdT = 90
	STYLE_TEXT_LETTER_SPACE         X_lvStyleIdT = 91
	STYLE_TEXT_LINE_SPACE           X_lvStyleIdT = 92
	STYLE_TEXT_DECOR                X_lvStyleIdT = 93
	STYLE_TEXT_ALIGN                X_lvStyleIdT = 94
	STYLE_TEXT_OUTLINE_STROKE_WIDTH X_lvStyleIdT = 95
	STYLE_TEXT_OUTLINE_STROKE_OPA   X_lvStyleIdT = 96
	STYLE_TEXT_OUTLINE_STROKE_COLOR X_lvStyleIdT = 97
	STYLE_OPA                       X_lvStyleIdT = 98
	STYLE_OPA_LAYERED               X_lvStyleIdT = 99
	STYLE_COLOR_FILTER_DSC          X_lvStyleIdT = 100
	STYLE_COLOR_FILTER_OPA          X_lvStyleIdT = 101
	STYLE_ANIM                      X_lvStyleIdT = 102
	STYLE_ANIM_DURATION             X_lvStyleIdT = 103
	STYLE_TRANSITION                X_lvStyleIdT = 104
	STYLE_BLEND_MODE                X_lvStyleIdT = 105
	STYLE_TRANSFORM_WIDTH           X_lvStyleIdT = 106
	STYLE_TRANSFORM_HEIGHT          X_lvStyleIdT = 107
	STYLE_TRANSLATE_X               X_lvStyleIdT = 108
	STYLE_TRANSLATE_Y               X_lvStyleIdT = 109
	STYLE_TRANSFORM_SCALE_X         X_lvStyleIdT = 110
	STYLE_TRANSFORM_SCALE_Y         X_lvStyleIdT = 111
	STYLE_TRANSFORM_ROTATION        X_lvStyleIdT = 112
	STYLE_TRANSFORM_PIVOT_X         X_lvStyleIdT = 113
	STYLE_TRANSFORM_PIVOT_Y         X_lvStyleIdT = 114
	STYLE_TRANSFORM_SKEW_X          X_lvStyleIdT = 115
	STYLE_TRANSFORM_SKEW_Y          X_lvStyleIdT = 116
	STYLE_BITMAP_MASK_SRC           X_lvStyleIdT = 117
	STYLE_ROTARY_SENSITIVITY        X_lvStyleIdT = 118
	STYLE_TRANSLATE_RADIAL          X_lvStyleIdT = 119
	STYLE_RECOLOR                   X_lvStyleIdT = 120
	STYLE_RECOLOR_OPA               X_lvStyleIdT = 121
	STYLE_FLEX_FLOW                 X_lvStyleIdT = 122
	STYLE_FLEX_MAIN_PLACE           X_lvStyleIdT = 123
	STYLE_FLEX_CROSS_PLACE          X_lvStyleIdT = 124
	STYLE_FLEX_TRACK_PLACE          X_lvStyleIdT = 125
	STYLE_FLEX_GROW                 X_lvStyleIdT = 126
	STYLE_GRID_COLUMN_ALIGN         X_lvStyleIdT = 127
	STYLE_GRID_ROW_ALIGN            X_lvStyleIdT = 128
	STYLE_GRID_ROW_DSC_ARRAY        X_lvStyleIdT = 129
	STYLE_GRID_COLUMN_DSC_ARRAY     X_lvStyleIdT = 130
	STYLE_GRID_CELL_COLUMN_POS      X_lvStyleIdT = 131
	STYLE_GRID_CELL_COLUMN_SPAN     X_lvStyleIdT = 132
	STYLE_GRID_CELL_X_ALIGN         X_lvStyleIdT = 133
	STYLE_GRID_CELL_ROW_POS         X_lvStyleIdT = 134
	STYLE_GRID_CELL_ROW_SPAN        X_lvStyleIdT = 135
	STYLE_GRID_CELL_Y_ALIGN         X_lvStyleIdT = 136
	STYLE_LAST_BUILT_IN_PROP        X_lvStyleIdT = 137
	STYLE_NUM_BUILT_IN_PROPS        X_lvStyleIdT = 138
	STYLE_PROP_ANY                  X_lvStyleIdT = 255
	STYLE_PROP_CONST                X_lvStyleIdT = 255
)

type StyleResT c.Int

const (
	STYLE_RES_NOT_FOUND StyleResT = 0
	STYLE_RES_FOUND     StyleResT = 1
)

/**
 * Descriptor for style transitions
 */

type StyleTransitionDscT struct {
	Props    *StylePropT
	UserData c.Pointer
	PathXcb  AnimPathCbT
	Time     c.Uint32T
	Delay    c.Uint32T
}

/**
 * Descriptor of a constant style property.
 */

type StyleConstPropT struct {
	Prop  StylePropT
	Value StyleValueT
}

/**
 * Descriptor of a style (a collection of properties and values).
 */

type StyleT struct {
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
// llgo:link (*StyleT).StyleInit C.lv_style_init
func (recv_ *StyleT) StyleInit() {
}

/**
 * Clear all properties from a style and free all allocated memories.
 * @param style pointer to a style
 */
// llgo:link (*StyleT).StyleReset C.lv_style_reset
func (recv_ *StyleT) StyleReset() {
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
// llgo:link (*StyleT).StyleCopy C.lv_style_copy
func (recv_ *StyleT) StyleCopy(src *StyleT) {
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
//go:linkname StyleRegisterProp C.lv_style_register_prop
func StyleRegisterProp(flag c.Uint8T) StylePropT

/**
 * Get the number of custom properties that have been registered thus far.
 */
//go:linkname StyleGetNumCustomProps C.lv_style_get_num_custom_props
func StyleGetNumCustomProps() StylePropT

/**
 * Remove a property from a style
 * @param style pointer to a style
 * @param prop  a style property ORed with a state.
 * @return true: the property was found and removed; false: the property wasn't found
 */
// llgo:link (*StyleT).StyleRemoveProp C.lv_style_remove_prop
func (recv_ *StyleT) StyleRemoveProp(prop StylePropT) bool {
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
// llgo:link (*StyleT).StyleSetProp C.lv_style_set_prop
func (recv_ *StyleT) StyleSetProp(prop StylePropT, value StyleValueT) {
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
// llgo:link (*StyleT).StyleGetProp C.lv_style_get_prop
func (recv_ *StyleT) StyleGetProp(prop StylePropT, value *StyleValueT) StyleResT {
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
// llgo:link (*StyleTransitionDscT).StyleTransitionDscInit C.lv_style_transition_dsc_init
func (recv_ *StyleTransitionDscT) StyleTransitionDscInit(props *StylePropT, path_cb AnimPathCbT, time c.Uint32T, delay c.Uint32T, user_data c.Pointer) {
}

/**
 * Get the default value of a property
 * @param prop the ID of a property
 * @return the default value
 */
// llgo:link StylePropT.StylePropGetDefault C.lv_style_prop_get_default
func (recv_ StylePropT) StylePropGetDefault() StyleValueT {
	return StyleValueT{}
}

/**
 * Checks if a style is empty (has no properties)
 * @param style pointer to a style
 * @return true if the style is empty
 */
// llgo:link (*StyleT).StyleIsEmpty C.lv_style_is_empty
func (recv_ *StyleT) StyleIsEmpty() bool {
	return false
}

/**
 * Get the flags of a built-in or custom property.
 *
 * @param prop a style property
 * @return the flags of the property
 */
// llgo:link StylePropT.StylePropLookupFlags C.lv_style_prop_lookup_flags
func (recv_ StylePropT) StylePropLookupFlags() c.Uint8T {
	return 0
}

// llgo:link (*StyleT).StyleSetWidth C.lv_style_set_width
func (recv_ *StyleT) StyleSetWidth(value c.Int32T) {
}

// llgo:link (*StyleT).StyleSetMinWidth C.lv_style_set_min_width
func (recv_ *StyleT) StyleSetMinWidth(value c.Int32T) {
}

// llgo:link (*StyleT).StyleSetMaxWidth C.lv_style_set_max_width
func (recv_ *StyleT) StyleSetMaxWidth(value c.Int32T) {
}

// llgo:link (*StyleT).StyleSetHeight C.lv_style_set_height
func (recv_ *StyleT) StyleSetHeight(value c.Int32T) {
}

// llgo:link (*StyleT).StyleSetMinHeight C.lv_style_set_min_height
func (recv_ *StyleT) StyleSetMinHeight(value c.Int32T) {
}

// llgo:link (*StyleT).StyleSetMaxHeight C.lv_style_set_max_height
func (recv_ *StyleT) StyleSetMaxHeight(value c.Int32T) {
}

// llgo:link (*StyleT).StyleSetLength C.lv_style_set_length
func (recv_ *StyleT) StyleSetLength(value c.Int32T) {
}

// llgo:link (*StyleT).StyleSetX C.lv_style_set_x
func (recv_ *StyleT) StyleSetX(value c.Int32T) {
}

// llgo:link (*StyleT).StyleSetY C.lv_style_set_y
func (recv_ *StyleT) StyleSetY(value c.Int32T) {
}

// llgo:link (*StyleT).StyleSetAlign C.lv_style_set_align
func (recv_ *StyleT) StyleSetAlign(value AlignT) {
}

// llgo:link (*StyleT).StyleSetTransformWidth C.lv_style_set_transform_width
func (recv_ *StyleT) StyleSetTransformWidth(value c.Int32T) {
}

// llgo:link (*StyleT).StyleSetTransformHeight C.lv_style_set_transform_height
func (recv_ *StyleT) StyleSetTransformHeight(value c.Int32T) {
}

// llgo:link (*StyleT).StyleSetTranslateX C.lv_style_set_translate_x
func (recv_ *StyleT) StyleSetTranslateX(value c.Int32T) {
}

// llgo:link (*StyleT).StyleSetTranslateY C.lv_style_set_translate_y
func (recv_ *StyleT) StyleSetTranslateY(value c.Int32T) {
}

// llgo:link (*StyleT).StyleSetTranslateRadial C.lv_style_set_translate_radial
func (recv_ *StyleT) StyleSetTranslateRadial(value c.Int32T) {
}

// llgo:link (*StyleT).StyleSetTransformScaleX C.lv_style_set_transform_scale_x
func (recv_ *StyleT) StyleSetTransformScaleX(value c.Int32T) {
}

// llgo:link (*StyleT).StyleSetTransformScaleY C.lv_style_set_transform_scale_y
func (recv_ *StyleT) StyleSetTransformScaleY(value c.Int32T) {
}

// llgo:link (*StyleT).StyleSetTransformRotation C.lv_style_set_transform_rotation
func (recv_ *StyleT) StyleSetTransformRotation(value c.Int32T) {
}

// llgo:link (*StyleT).StyleSetTransformPivotX C.lv_style_set_transform_pivot_x
func (recv_ *StyleT) StyleSetTransformPivotX(value c.Int32T) {
}

// llgo:link (*StyleT).StyleSetTransformPivotY C.lv_style_set_transform_pivot_y
func (recv_ *StyleT) StyleSetTransformPivotY(value c.Int32T) {
}

// llgo:link (*StyleT).StyleSetTransformSkewX C.lv_style_set_transform_skew_x
func (recv_ *StyleT) StyleSetTransformSkewX(value c.Int32T) {
}

// llgo:link (*StyleT).StyleSetTransformSkewY C.lv_style_set_transform_skew_y
func (recv_ *StyleT) StyleSetTransformSkewY(value c.Int32T) {
}

// llgo:link (*StyleT).StyleSetPadTop C.lv_style_set_pad_top
func (recv_ *StyleT) StyleSetPadTop(value c.Int32T) {
}

// llgo:link (*StyleT).StyleSetPadBottom C.lv_style_set_pad_bottom
func (recv_ *StyleT) StyleSetPadBottom(value c.Int32T) {
}

// llgo:link (*StyleT).StyleSetPadLeft C.lv_style_set_pad_left
func (recv_ *StyleT) StyleSetPadLeft(value c.Int32T) {
}

// llgo:link (*StyleT).StyleSetPadRight C.lv_style_set_pad_right
func (recv_ *StyleT) StyleSetPadRight(value c.Int32T) {
}

// llgo:link (*StyleT).StyleSetPadRow C.lv_style_set_pad_row
func (recv_ *StyleT) StyleSetPadRow(value c.Int32T) {
}

// llgo:link (*StyleT).StyleSetPadColumn C.lv_style_set_pad_column
func (recv_ *StyleT) StyleSetPadColumn(value c.Int32T) {
}

// llgo:link (*StyleT).StyleSetPadRadial C.lv_style_set_pad_radial
func (recv_ *StyleT) StyleSetPadRadial(value c.Int32T) {
}

// llgo:link (*StyleT).StyleSetMarginTop C.lv_style_set_margin_top
func (recv_ *StyleT) StyleSetMarginTop(value c.Int32T) {
}

// llgo:link (*StyleT).StyleSetMarginBottom C.lv_style_set_margin_bottom
func (recv_ *StyleT) StyleSetMarginBottom(value c.Int32T) {
}

// llgo:link (*StyleT).StyleSetMarginLeft C.lv_style_set_margin_left
func (recv_ *StyleT) StyleSetMarginLeft(value c.Int32T) {
}

// llgo:link (*StyleT).StyleSetMarginRight C.lv_style_set_margin_right
func (recv_ *StyleT) StyleSetMarginRight(value c.Int32T) {
}

// llgo:link (*StyleT).StyleSetBgColor C.lv_style_set_bg_color
func (recv_ *StyleT) StyleSetBgColor(value ColorT) {
}

// llgo:link (*StyleT).StyleSetBgOpa C.lv_style_set_bg_opa
func (recv_ *StyleT) StyleSetBgOpa(value OpaT) {
}

// llgo:link (*StyleT).StyleSetBgGradColor C.lv_style_set_bg_grad_color
func (recv_ *StyleT) StyleSetBgGradColor(value ColorT) {
}

// llgo:link (*StyleT).StyleSetBgGradDir C.lv_style_set_bg_grad_dir
func (recv_ *StyleT) StyleSetBgGradDir(value GradDirT) {
}

// llgo:link (*StyleT).StyleSetBgMainStop C.lv_style_set_bg_main_stop
func (recv_ *StyleT) StyleSetBgMainStop(value c.Int32T) {
}

// llgo:link (*StyleT).StyleSetBgGradStop C.lv_style_set_bg_grad_stop
func (recv_ *StyleT) StyleSetBgGradStop(value c.Int32T) {
}

// llgo:link (*StyleT).StyleSetBgMainOpa C.lv_style_set_bg_main_opa
func (recv_ *StyleT) StyleSetBgMainOpa(value OpaT) {
}

// llgo:link (*StyleT).StyleSetBgGradOpa C.lv_style_set_bg_grad_opa
func (recv_ *StyleT) StyleSetBgGradOpa(value OpaT) {
}

// llgo:link (*StyleT).StyleSetBgGrad C.lv_style_set_bg_grad
func (recv_ *StyleT) StyleSetBgGrad(value *GradDscT) {
}

// llgo:link (*StyleT).StyleSetBgImageSrc C.lv_style_set_bg_image_src
func (recv_ *StyleT) StyleSetBgImageSrc(value c.Pointer) {
}

// llgo:link (*StyleT).StyleSetBgImageOpa C.lv_style_set_bg_image_opa
func (recv_ *StyleT) StyleSetBgImageOpa(value OpaT) {
}

// llgo:link (*StyleT).StyleSetBgImageRecolor C.lv_style_set_bg_image_recolor
func (recv_ *StyleT) StyleSetBgImageRecolor(value ColorT) {
}

// llgo:link (*StyleT).StyleSetBgImageRecolorOpa C.lv_style_set_bg_image_recolor_opa
func (recv_ *StyleT) StyleSetBgImageRecolorOpa(value OpaT) {
}

// llgo:link (*StyleT).StyleSetBgImageTiled C.lv_style_set_bg_image_tiled
func (recv_ *StyleT) StyleSetBgImageTiled(value bool) {
}

// llgo:link (*StyleT).StyleSetBorderColor C.lv_style_set_border_color
func (recv_ *StyleT) StyleSetBorderColor(value ColorT) {
}

// llgo:link (*StyleT).StyleSetBorderOpa C.lv_style_set_border_opa
func (recv_ *StyleT) StyleSetBorderOpa(value OpaT) {
}

// llgo:link (*StyleT).StyleSetBorderWidth C.lv_style_set_border_width
func (recv_ *StyleT) StyleSetBorderWidth(value c.Int32T) {
}

// llgo:link (*StyleT).StyleSetBorderSide C.lv_style_set_border_side
func (recv_ *StyleT) StyleSetBorderSide(value BorderSideT) {
}

// llgo:link (*StyleT).StyleSetBorderPost C.lv_style_set_border_post
func (recv_ *StyleT) StyleSetBorderPost(value bool) {
}

// llgo:link (*StyleT).StyleSetOutlineWidth C.lv_style_set_outline_width
func (recv_ *StyleT) StyleSetOutlineWidth(value c.Int32T) {
}

// llgo:link (*StyleT).StyleSetOutlineColor C.lv_style_set_outline_color
func (recv_ *StyleT) StyleSetOutlineColor(value ColorT) {
}

// llgo:link (*StyleT).StyleSetOutlineOpa C.lv_style_set_outline_opa
func (recv_ *StyleT) StyleSetOutlineOpa(value OpaT) {
}

// llgo:link (*StyleT).StyleSetOutlinePad C.lv_style_set_outline_pad
func (recv_ *StyleT) StyleSetOutlinePad(value c.Int32T) {
}

// llgo:link (*StyleT).StyleSetShadowWidth C.lv_style_set_shadow_width
func (recv_ *StyleT) StyleSetShadowWidth(value c.Int32T) {
}

// llgo:link (*StyleT).StyleSetShadowOffsetX C.lv_style_set_shadow_offset_x
func (recv_ *StyleT) StyleSetShadowOffsetX(value c.Int32T) {
}

// llgo:link (*StyleT).StyleSetShadowOffsetY C.lv_style_set_shadow_offset_y
func (recv_ *StyleT) StyleSetShadowOffsetY(value c.Int32T) {
}

// llgo:link (*StyleT).StyleSetShadowSpread C.lv_style_set_shadow_spread
func (recv_ *StyleT) StyleSetShadowSpread(value c.Int32T) {
}

// llgo:link (*StyleT).StyleSetShadowColor C.lv_style_set_shadow_color
func (recv_ *StyleT) StyleSetShadowColor(value ColorT) {
}

// llgo:link (*StyleT).StyleSetShadowOpa C.lv_style_set_shadow_opa
func (recv_ *StyleT) StyleSetShadowOpa(value OpaT) {
}

// llgo:link (*StyleT).StyleSetImageOpa C.lv_style_set_image_opa
func (recv_ *StyleT) StyleSetImageOpa(value OpaT) {
}

// llgo:link (*StyleT).StyleSetImageRecolor C.lv_style_set_image_recolor
func (recv_ *StyleT) StyleSetImageRecolor(value ColorT) {
}

// llgo:link (*StyleT).StyleSetImageRecolorOpa C.lv_style_set_image_recolor_opa
func (recv_ *StyleT) StyleSetImageRecolorOpa(value OpaT) {
}

// llgo:link (*StyleT).StyleSetLineWidth C.lv_style_set_line_width
func (recv_ *StyleT) StyleSetLineWidth(value c.Int32T) {
}

// llgo:link (*StyleT).StyleSetLineDashWidth C.lv_style_set_line_dash_width
func (recv_ *StyleT) StyleSetLineDashWidth(value c.Int32T) {
}

// llgo:link (*StyleT).StyleSetLineDashGap C.lv_style_set_line_dash_gap
func (recv_ *StyleT) StyleSetLineDashGap(value c.Int32T) {
}

// llgo:link (*StyleT).StyleSetLineRounded C.lv_style_set_line_rounded
func (recv_ *StyleT) StyleSetLineRounded(value bool) {
}

// llgo:link (*StyleT).StyleSetLineColor C.lv_style_set_line_color
func (recv_ *StyleT) StyleSetLineColor(value ColorT) {
}

// llgo:link (*StyleT).StyleSetLineOpa C.lv_style_set_line_opa
func (recv_ *StyleT) StyleSetLineOpa(value OpaT) {
}

// llgo:link (*StyleT).StyleSetArcWidth C.lv_style_set_arc_width
func (recv_ *StyleT) StyleSetArcWidth(value c.Int32T) {
}

// llgo:link (*StyleT).StyleSetArcRounded C.lv_style_set_arc_rounded
func (recv_ *StyleT) StyleSetArcRounded(value bool) {
}

// llgo:link (*StyleT).StyleSetArcColor C.lv_style_set_arc_color
func (recv_ *StyleT) StyleSetArcColor(value ColorT) {
}

// llgo:link (*StyleT).StyleSetArcOpa C.lv_style_set_arc_opa
func (recv_ *StyleT) StyleSetArcOpa(value OpaT) {
}

// llgo:link (*StyleT).StyleSetArcImageSrc C.lv_style_set_arc_image_src
func (recv_ *StyleT) StyleSetArcImageSrc(value c.Pointer) {
}

// llgo:link (*StyleT).StyleSetTextColor C.lv_style_set_text_color
func (recv_ *StyleT) StyleSetTextColor(value ColorT) {
}

// llgo:link (*StyleT).StyleSetTextOpa C.lv_style_set_text_opa
func (recv_ *StyleT) StyleSetTextOpa(value OpaT) {
}

// llgo:link (*StyleT).StyleSetTextFont C.lv_style_set_text_font
func (recv_ *StyleT) StyleSetTextFont(value *FontT) {
}

// llgo:link (*StyleT).StyleSetTextLetterSpace C.lv_style_set_text_letter_space
func (recv_ *StyleT) StyleSetTextLetterSpace(value c.Int32T) {
}

// llgo:link (*StyleT).StyleSetTextLineSpace C.lv_style_set_text_line_space
func (recv_ *StyleT) StyleSetTextLineSpace(value c.Int32T) {
}

// llgo:link (*StyleT).StyleSetTextDecor C.lv_style_set_text_decor
func (recv_ *StyleT) StyleSetTextDecor(value TextDecorT) {
}

// llgo:link (*StyleT).StyleSetTextAlign C.lv_style_set_text_align
func (recv_ *StyleT) StyleSetTextAlign(value TextAlignT) {
}

// llgo:link (*StyleT).StyleSetTextOutlineStrokeColor C.lv_style_set_text_outline_stroke_color
func (recv_ *StyleT) StyleSetTextOutlineStrokeColor(value ColorT) {
}

// llgo:link (*StyleT).StyleSetTextOutlineStrokeWidth C.lv_style_set_text_outline_stroke_width
func (recv_ *StyleT) StyleSetTextOutlineStrokeWidth(value c.Int32T) {
}

// llgo:link (*StyleT).StyleSetTextOutlineStrokeOpa C.lv_style_set_text_outline_stroke_opa
func (recv_ *StyleT) StyleSetTextOutlineStrokeOpa(value OpaT) {
}

// llgo:link (*StyleT).StyleSetRadius C.lv_style_set_radius
func (recv_ *StyleT) StyleSetRadius(value c.Int32T) {
}

// llgo:link (*StyleT).StyleSetRadialOffset C.lv_style_set_radial_offset
func (recv_ *StyleT) StyleSetRadialOffset(value c.Int32T) {
}

// llgo:link (*StyleT).StyleSetClipCorner C.lv_style_set_clip_corner
func (recv_ *StyleT) StyleSetClipCorner(value bool) {
}

// llgo:link (*StyleT).StyleSetOpa C.lv_style_set_opa
func (recv_ *StyleT) StyleSetOpa(value OpaT) {
}

// llgo:link (*StyleT).StyleSetOpaLayered C.lv_style_set_opa_layered
func (recv_ *StyleT) StyleSetOpaLayered(value OpaT) {
}

// llgo:link (*StyleT).StyleSetColorFilterDsc C.lv_style_set_color_filter_dsc
func (recv_ *StyleT) StyleSetColorFilterDsc(value *ColorFilterDscT) {
}

// llgo:link (*StyleT).StyleSetColorFilterOpa C.lv_style_set_color_filter_opa
func (recv_ *StyleT) StyleSetColorFilterOpa(value OpaT) {
}

// llgo:link (*StyleT).StyleSetRecolor C.lv_style_set_recolor
func (recv_ *StyleT) StyleSetRecolor(value ColorT) {
}

// llgo:link (*StyleT).StyleSetRecolorOpa C.lv_style_set_recolor_opa
func (recv_ *StyleT) StyleSetRecolorOpa(value OpaT) {
}

// llgo:link (*StyleT).StyleSetAnim C.lv_style_set_anim
func (recv_ *StyleT) StyleSetAnim(value *AnimT) {
}

// llgo:link (*StyleT).StyleSetAnimDuration C.lv_style_set_anim_duration
func (recv_ *StyleT) StyleSetAnimDuration(value c.Uint32T) {
}

// llgo:link (*StyleT).StyleSetTransition C.lv_style_set_transition
func (recv_ *StyleT) StyleSetTransition(value *StyleTransitionDscT) {
}

// llgo:link (*StyleT).StyleSetBlendMode C.lv_style_set_blend_mode
func (recv_ *StyleT) StyleSetBlendMode(value BlendModeT) {
}

// llgo:link (*StyleT).StyleSetLayout C.lv_style_set_layout
func (recv_ *StyleT) StyleSetLayout(value c.Uint16T) {
}

// llgo:link (*StyleT).StyleSetBaseDir C.lv_style_set_base_dir
func (recv_ *StyleT) StyleSetBaseDir(value BaseDirT) {
}

// llgo:link (*StyleT).StyleSetBitmapMaskSrc C.lv_style_set_bitmap_mask_src
func (recv_ *StyleT) StyleSetBitmapMaskSrc(value c.Pointer) {
}

// llgo:link (*StyleT).StyleSetRotarySensitivity C.lv_style_set_rotary_sensitivity
func (recv_ *StyleT) StyleSetRotarySensitivity(value c.Uint32T) {
}

// llgo:link (*StyleT).StyleSetFlexFlow C.lv_style_set_flex_flow
func (recv_ *StyleT) StyleSetFlexFlow(value FlexFlowT) {
}

// llgo:link (*StyleT).StyleSetFlexMainPlace C.lv_style_set_flex_main_place
func (recv_ *StyleT) StyleSetFlexMainPlace(value FlexAlignT) {
}

// llgo:link (*StyleT).StyleSetFlexCrossPlace C.lv_style_set_flex_cross_place
func (recv_ *StyleT) StyleSetFlexCrossPlace(value FlexAlignT) {
}

// llgo:link (*StyleT).StyleSetFlexTrackPlace C.lv_style_set_flex_track_place
func (recv_ *StyleT) StyleSetFlexTrackPlace(value FlexAlignT) {
}

// llgo:link (*StyleT).StyleSetFlexGrow C.lv_style_set_flex_grow
func (recv_ *StyleT) StyleSetFlexGrow(value c.Uint8T) {
}

// llgo:link (*StyleT).StyleSetGridColumnDscArray C.lv_style_set_grid_column_dsc_array
func (recv_ *StyleT) StyleSetGridColumnDscArray(value *c.Int32T) {
}

// llgo:link (*StyleT).StyleSetGridColumnAlign C.lv_style_set_grid_column_align
func (recv_ *StyleT) StyleSetGridColumnAlign(value GridAlignT) {
}

// llgo:link (*StyleT).StyleSetGridRowDscArray C.lv_style_set_grid_row_dsc_array
func (recv_ *StyleT) StyleSetGridRowDscArray(value *c.Int32T) {
}

// llgo:link (*StyleT).StyleSetGridRowAlign C.lv_style_set_grid_row_align
func (recv_ *StyleT) StyleSetGridRowAlign(value GridAlignT) {
}

// llgo:link (*StyleT).StyleSetGridCellColumnPos C.lv_style_set_grid_cell_column_pos
func (recv_ *StyleT) StyleSetGridCellColumnPos(value c.Int32T) {
}

// llgo:link (*StyleT).StyleSetGridCellXAlign C.lv_style_set_grid_cell_x_align
func (recv_ *StyleT) StyleSetGridCellXAlign(value GridAlignT) {
}

// llgo:link (*StyleT).StyleSetGridCellColumnSpan C.lv_style_set_grid_cell_column_span
func (recv_ *StyleT) StyleSetGridCellColumnSpan(value c.Int32T) {
}

// llgo:link (*StyleT).StyleSetGridCellRowPos C.lv_style_set_grid_cell_row_pos
func (recv_ *StyleT) StyleSetGridCellRowPos(value c.Int32T) {
}

// llgo:link (*StyleT).StyleSetGridCellYAlign C.lv_style_set_grid_cell_y_align
func (recv_ *StyleT) StyleSetGridCellYAlign(value GridAlignT) {
}

// llgo:link (*StyleT).StyleSetGridCellRowSpan C.lv_style_set_grid_cell_row_span
func (recv_ *StyleT) StyleSetGridCellRowSpan(value c.Int32T) {
}

type FsResT c.Int

const (
	FS_RES_OK         FsResT = 0
	FS_RES_HW_ERR     FsResT = 1
	FS_RES_FS_ERR     FsResT = 2
	FS_RES_NOT_EX     FsResT = 3
	FS_RES_FULL       FsResT = 4
	FS_RES_LOCKED     FsResT = 5
	FS_RES_DENIED     FsResT = 6
	FS_RES_BUSY       FsResT = 7
	FS_RES_TOUT       FsResT = 8
	FS_RES_NOT_IMP    FsResT = 9
	FS_RES_OUT_OF_MEM FsResT = 10
	FS_RES_INV_PARAM  FsResT = 11
	FS_RES_UNKNOWN    FsResT = 12
)

type FsModeT c.Int

const (
	FS_MODE_WR FsModeT = 1
	FS_MODE_RD FsModeT = 2
)

type FsWhenceT c.Int

const (
	FS_SEEK_SET FsWhenceT = 0
	FS_SEEK_CUR FsWhenceT = 1
	FS_SEEK_END FsWhenceT = 2
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
type FsDrvT X_lvFsDrvT

type FsFileT struct {
	FileD c.Pointer
	Drv   *FsDrvT
	Cache *FsFileCacheT
}

type FsDirT struct {
	DirD c.Pointer
	Drv  *FsDrvT
}

/**
 * Initialize a file system driver with default values.
 * It is used to ensure all fields have known values and not memory junk.
 * After it you can set the fields.
 * @param drv     pointer to driver variable to initialize
 */
// llgo:link (*FsDrvT).FsDrvInit C.lv_fs_drv_init
func (recv_ *FsDrvT) FsDrvInit() {
}

/**
 * Add a new drive
 * @param drv       pointer to an lv_fs_drv_t structure which is inited with the
 *                  corresponding function pointers. Only pointer is saved, so the
 *                  driver should be static or dynamically allocated.
 */
// llgo:link (*FsDrvT).FsDrvRegister C.lv_fs_drv_register
func (recv_ *FsDrvT) FsDrvRegister() {
}

/**
 * Give a pointer to a driver from its letter
 * @param letter    the driver-identifier letter
 * @return          pointer to a driver or NULL if not found
 */
//go:linkname FsGetDrv C.lv_fs_get_drv
func FsGetDrv(letter c.Char) *FsDrvT

/**
 * Test if a drive is ready or not. If the `ready` function was not initialized `true` will be
 * returned.
 * @param letter    letter of the drive
 * @return          true: drive is ready; false: drive is not ready
 */
//go:linkname FsIsReady C.lv_fs_is_ready
func FsIsReady(letter c.Char) bool

/**
 * Open a file
 * @param file_p    pointer to a lv_fs_file_t variable
 * @param path      path to the file beginning with the driver letter (e.g. S:/folder/file.txt)
 * @param mode      read: FS_MODE_RD, write: FS_MODE_WR, both: FS_MODE_RD | FS_MODE_WR
 * @return          LV_FS_RES_OK or any error from lv_fs_res_t enum
 */
// llgo:link (*FsFileT).FsOpen C.lv_fs_open
func (recv_ *FsFileT) FsOpen(path *c.Char, mode FsModeT) FsResT {
	return 0
}

/**
 * Make a path object for the memory-mapped file compatible with the file system interface
 * @param path      path to a lv_fs_path_ex object
 * @param letter    the identifier letter of the driver. E.g. `LV_FS_MEMFS_LETTER`
 * @param buf       address of the memory buffer
 * @param size      size of the memory buffer in bytes
 */
// llgo:link (*FsPathExT).FsMakePathFromBuffer C.lv_fs_make_path_from_buffer
func (recv_ *FsPathExT) FsMakePathFromBuffer(letter c.Char, buf c.Pointer, size c.Uint32T) {
}

/**
 * Close an already opened file
 * @param file_p    pointer to a lv_fs_file_t variable
 * @return          LV_FS_RES_OK or any error from lv_fs_res_t enum
 */
// llgo:link (*FsFileT).FsClose C.lv_fs_close
func (recv_ *FsFileT) FsClose() FsResT {
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
// llgo:link (*FsFileT).FsRead C.lv_fs_read
func (recv_ *FsFileT) FsRead(buf c.Pointer, btr c.Uint32T, br *c.Uint32T) FsResT {
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
// llgo:link (*FsFileT).FsWrite C.lv_fs_write
func (recv_ *FsFileT) FsWrite(buf c.Pointer, btw c.Uint32T, bw *c.Uint32T) FsResT {
	return 0
}

/**
 * Set the position of the 'cursor' (read write pointer) in a file
 * @param file_p    pointer to a lv_fs_file_t variable
 * @param pos       the new position expressed in bytes index (0: start of file)
 * @param whence    tells from where to set position. See lv_fs_whence_t
 * @return          LV_FS_RES_OK or any error from lv_fs_res_t enum
 */
// llgo:link (*FsFileT).FsSeek C.lv_fs_seek
func (recv_ *FsFileT) FsSeek(pos c.Uint32T, whence FsWhenceT) FsResT {
	return 0
}

/**
 * Give the position of the read write pointer
 * @param file_p    pointer to a lv_fs_file_t variable
 * @param pos       pointer to store the position of the read write pointer
 * @return          LV_FS_RES_OK or any error from 'fs_res_t'
 */
// llgo:link (*FsFileT).FsTell C.lv_fs_tell
func (recv_ *FsFileT) FsTell(pos *c.Uint32T) FsResT {
	return 0
}

/**
 * Initialize a 'fs_dir_t' variable for directory reading
 * @param rddir_p   pointer to a 'lv_fs_dir_t' variable
 * @param path      path to a directory
 * @return          LV_FS_RES_OK or any error from lv_fs_res_t enum
 */
// llgo:link (*FsDirT).FsDirOpen C.lv_fs_dir_open
func (recv_ *FsDirT) FsDirOpen(path *c.Char) FsResT {
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
// llgo:link (*FsDirT).FsDirRead C.lv_fs_dir_read
func (recv_ *FsDirT) FsDirRead(fn *c.Char, fn_len c.Uint32T) FsResT {
	return 0
}

/**
 * Close the directory reading
 * @param rddir_p   pointer to an initialized 'fs_dir_t' variable
 * @return          LV_FS_RES_OK or any error from lv_fs_res_t enum
 */
// llgo:link (*FsDirT).FsDirClose C.lv_fs_dir_close
func (recv_ *FsDirT) FsDirClose() FsResT {
	return 0
}

/**
 * Fill a buffer with the letters of existing drivers
 * @param buf       buffer to store the letters ('\0' added after the last letter)
 * @return          the buffer
 */
//go:linkname FsGetLetters C.lv_fs_get_letters
func FsGetLetters(buf *c.Char) *c.Char

/**
 * Return with the extension of the filename
 * @param fn        string with a filename
 * @return          pointer to the beginning extension or empty string if no extension
 */
//go:linkname FsGetExt C.lv_fs_get_ext
func FsGetExt(fn *c.Char) *c.Char

/**
 * Step up one level
 * @param path      pointer to a file name
 * @return          the truncated file name
 */
//go:linkname FsUp C.lv_fs_up
func FsUp(path *c.Char) *c.Char

/**
 * Get the last element of a path (e.g. U:/folder/file -> file)
 * @param path      pointer to a file name
 * @return          pointer to the beginning of the last element in the path
 */
//go:linkname FsGetLast C.lv_fs_get_last
func FsGetLast(path *c.Char) *c.Char

type ImageSrcT c.Int

const (
	IMAGE_SRC_VARIABLE ImageSrcT = 0
	IMAGE_SRC_FILE     ImageSrcT = 1
	IMAGE_SRC_SYMBOL   ImageSrcT = 2
	IMAGE_SRC_UNKNOWN  ImageSrcT = 3
)

// llgo:type C
type ImageDecoderInfoFT func(*ImageDecoderT, *ImageDecoderDscT, *ImageHeaderT) ResultT

// llgo:type C
type ImageDecoderOpenFT func(*ImageDecoderT, *ImageDecoderDscT) ResultT

// llgo:type C
type ImageDecoderGetAreaCbT func(*ImageDecoderT, *ImageDecoderDscT, *AreaT, *AreaT) ResultT

// llgo:type C
type ImageDecoderCloseFT func(*ImageDecoderT, *ImageDecoderDscT)

// llgo:type C
type ImageDecoderCustomDrawT func(*LayerT, *ImageDecoderDscT, *AreaT, *DrawImageDscT, *AreaT)

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
//go:linkname ImageDecoderGetInfo C.lv_image_decoder_get_info
func ImageDecoderGetInfo(src c.Pointer, header *ImageHeaderT) ResultT

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
// llgo:link (*ImageDecoderDscT).ImageDecoderOpen C.lv_image_decoder_open
func (recv_ *ImageDecoderDscT) ImageDecoderOpen(src c.Pointer, args *ImageDecoderArgsT) ResultT {
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
// llgo:link (*ImageDecoderDscT).ImageDecoderGetArea C.lv_image_decoder_get_area
func (recv_ *ImageDecoderDscT) ImageDecoderGetArea(full_area *AreaT, decoded_area *AreaT) ResultT {
	return 0
}

/**
 * Close a decoding session
 * @param dsc pointer to `lv_image_decoder_dsc_t` used in `lv_image_decoder_open`
 */
// llgo:link (*ImageDecoderDscT).ImageDecoderClose C.lv_image_decoder_close
func (recv_ *ImageDecoderDscT) ImageDecoderClose() {
}

/**
 * Create a new image decoder
 * @return pointer to the new image decoder
 */
//go:linkname ImageDecoderCreate C.lv_image_decoder_create
func ImageDecoderCreate() *ImageDecoderT

/**
 * Delete an image decoder
 * @param decoder pointer to an image decoder
 */
// llgo:link (*ImageDecoderT).ImageDecoderDelete C.lv_image_decoder_delete
func (recv_ *ImageDecoderT) ImageDecoderDelete() {
}

/**
 * Get the next image decoder in the linked list of image decoders
 * @param decoder pointer to an image decoder or NULL to get the first one
 * @return the next image decoder or NULL if no more image decoder exists
 */
// llgo:link (*ImageDecoderT).ImageDecoderGetNext C.lv_image_decoder_get_next
func (recv_ *ImageDecoderT) ImageDecoderGetNext() *ImageDecoderT {
	return nil
}

/**
 * Set a callback to get information about the image
 * @param decoder pointer to an image decoder
 * @param info_cb a function to collect info about an image (fill an `lv_image_header_t` struct)
 */
// llgo:link (*ImageDecoderT).ImageDecoderSetInfoCb C.lv_image_decoder_set_info_cb
func (recv_ *ImageDecoderT) ImageDecoderSetInfoCb(info_cb ImageDecoderInfoFT) {
}

/**
 * Set a callback to open an image
 * @param decoder pointer to an image decoder
 * @param open_cb a function to open an image
 */
// llgo:link (*ImageDecoderT).ImageDecoderSetOpenCb C.lv_image_decoder_set_open_cb
func (recv_ *ImageDecoderT) ImageDecoderSetOpenCb(open_cb ImageDecoderOpenFT) {
}

/**
 * Set a callback to a decoded line of an image
 * @param decoder pointer to an image decoder
 * @param read_line_cb a function to read a line of an image
 */
// llgo:link (*ImageDecoderT).ImageDecoderSetGetAreaCb C.lv_image_decoder_set_get_area_cb
func (recv_ *ImageDecoderT) ImageDecoderSetGetAreaCb(read_line_cb ImageDecoderGetAreaCbT) {
}

/**
 * Set a callback to close a decoding session. E.g. close files and free other resources.
 * @param decoder pointer to an image decoder
 * @param close_cb a function to close a decoding session
 */
// llgo:link (*ImageDecoderT).ImageDecoderSetCloseCb C.lv_image_decoder_set_close_cb
func (recv_ *ImageDecoderT) ImageDecoderSetCloseCb(close_cb ImageDecoderCloseFT) {
}

// llgo:link (*ImageDecoderT).ImageDecoderAddToCache C.lv_image_decoder_add_to_cache
func (recv_ *ImageDecoderT) ImageDecoderAddToCache(search_key *ImageCacheDataT, decoded *DrawBufT, user_data c.Pointer) *CacheEntryT {
	return nil
}

/**
 * Check the decoded image, make any modification if decoder `args` requires.
 * @note A new draw buf will be allocated if provided `decoded` is not modifiable or stride mismatch etc.
 * @param dsc       pointer to a decoder descriptor
 * @param decoded   pointer to a decoded image to post process to meet dsc->args requirement.
 * @return          post processed draw buffer, when it differs with `decoded`, it's newly allocated.
 */
// llgo:link (*ImageDecoderDscT).ImageDecoderPostProcess C.lv_image_decoder_post_process
func (recv_ *ImageDecoderDscT) ImageDecoderPostProcess(decoded *DrawBufT) *DrawBufT {
	return nil
}

type DrawTaskTypeT c.Int

const (
	DRAW_TASK_TYPE_NONE           DrawTaskTypeT = 0
	DRAW_TASK_TYPE_FILL           DrawTaskTypeT = 1
	DRAW_TASK_TYPE_BORDER         DrawTaskTypeT = 2
	DRAW_TASK_TYPE_BOX_SHADOW     DrawTaskTypeT = 3
	DRAW_TASK_TYPE_LETTER         DrawTaskTypeT = 4
	DRAW_TASK_TYPE_LABEL          DrawTaskTypeT = 5
	DRAW_TASK_TYPE_IMAGE          DrawTaskTypeT = 6
	DRAW_TASK_TYPE_LAYER          DrawTaskTypeT = 7
	DRAW_TASK_TYPE_LINE           DrawTaskTypeT = 8
	DRAW_TASK_TYPE_ARC            DrawTaskTypeT = 9
	DRAW_TASK_TYPE_TRIANGLE       DrawTaskTypeT = 10
	DRAW_TASK_TYPE_MASK_RECTANGLE DrawTaskTypeT = 11
	DRAW_TASK_TYPE_MASK_BITMAP    DrawTaskTypeT = 12
)

type DrawTaskStateT c.Int

const (
	DRAW_TASK_STATE_WAITING     DrawTaskStateT = 0
	DRAW_TASK_STATE_QUEUED      DrawTaskStateT = 1
	DRAW_TASK_STATE_IN_PROGRESS DrawTaskStateT = 2
	DRAW_TASK_STATE_READY       DrawTaskStateT = 3
)

type DrawDscBaseT struct {
	Obj      *ObjT
	Part     PartT
	Id1      c.Uint32T
	Id2      c.Uint32T
	Layer    *LayerT
	DscSize  c.SizeT
	UserData c.Pointer
}

/**
 * Used internally to initialize the drawing module
 */
//go:linkname DrawInit C.lv_draw_init
func DrawInit()

/**
 * Deinitialize the drawing module
 */
//go:linkname DrawDeinit C.lv_draw_deinit
func DrawDeinit()

/**
 * Allocate a new draw unit with the given size and appends it to the list of draw units
 * @param size      the size to allocate. E.g. `sizeof(my_draw_unit_t)`,
 *                  where the first element of `my_draw_unit_t` is `lv_draw_unit_t`.
 */
//go:linkname DrawCreateUnit C.lv_draw_create_unit
func DrawCreateUnit(size c.SizeT) c.Pointer

/**
 * Add an empty draw task to the draw task list of a layer.
 * @param layer     pointer to a layer
 * @param coords    the coordinates of the draw task
 * @return          the created draw task which needs to be
 *                  further configured e.g. by added a draw descriptor
 */
// llgo:link (*LayerT).DrawAddTask C.lv_draw_add_task
func (recv_ *LayerT) DrawAddTask(coords *AreaT, type_ DrawTaskTypeT) *DrawTaskT {
	return nil
}

/**
 * Needs to be called when a draw task is created and configured.
 * It will send an event about the new draw task to the widget
 * and assign it to a draw unit.
 * @param layer     pointer to a layer
 * @param t         pointer to a draw task
 */
// llgo:link (*LayerT).DrawFinalizeTaskCreation C.lv_draw_finalize_task_creation
func (recv_ *LayerT) DrawFinalizeTaskCreation(t *DrawTaskT) {
}

/**
 * Try dispatching draw tasks to draw units
 */
//go:linkname DrawDispatch C.lv_draw_dispatch
func DrawDispatch()

/**
 * Used internally to try dispatching draw tasks of a specific layer
 * @param disp      pointer to a display on which the dispatching was requested
 * @param layer     pointer to a layer
 * @return          at least one draw task is being rendered (maybe it was taken earlier)
 */
// llgo:link (*DisplayT).DrawDispatchLayer C.lv_draw_dispatch_layer
func (recv_ *DisplayT) DrawDispatchLayer(layer *LayerT) bool {
	return false
}

/**
 * Wait for a new dispatch request.
 * It's blocking if `LV_USE_OS == 0` else it yields
 */
//go:linkname DrawDispatchWaitForRequest C.lv_draw_dispatch_wait_for_request
func DrawDispatchWaitForRequest()

/**
 * Wait for draw finish in case of asynchronous task execution.
 * If `LV_USE_OS == 0` it just return.
 */
//go:linkname DrawWaitForFinish C.lv_draw_wait_for_finish
func DrawWaitForFinish()

/**
 * When a draw unit finished a draw task it needs to request dispatching
 * to let LVGL assign a new draw task to it
 */
//go:linkname DrawDispatchRequest C.lv_draw_dispatch_request
func DrawDispatchRequest()

/**
 * Get the total number of draw units.
 */
//go:linkname DrawGetUnitCount C.lv_draw_get_unit_count
func DrawGetUnitCount() c.Uint32T

/**
 * If there is only one draw unit check the first draw task if it's available.
 * If there are multiple draw units call `lv_draw_get_next_available_task` to find a task.
 * @param layer             the draw layer to search in
 * @param t_prev            continue searching from this task
 * @param draw_unit_id      check the task where `preferred_draw_unit_id` equals this value or `LV_DRAW_UNIT_NONE`
 * @return                  an available draw task or NULL if there is not any
 */
// llgo:link (*LayerT).DrawGetAvailableTask C.lv_draw_get_available_task
func (recv_ *LayerT) DrawGetAvailableTask(t_prev *DrawTaskT, draw_unit_id c.Uint8T) *DrawTaskT {
	return nil
}

/**
 * Find and available draw task
 * @param layer             the draw layer to search in
 * @param t_prev            continue searching from this task
 * @param draw_unit_id      check the task where `preferred_draw_unit_id` equals this value or `LV_DRAW_UNIT_NONE`
 * @return                  an available draw task or NULL if there is not any
 */
// llgo:link (*LayerT).DrawGetNextAvailableTask C.lv_draw_get_next_available_task
func (recv_ *LayerT) DrawGetNextAvailableTask(t_prev *DrawTaskT, draw_unit_id c.Uint8T) *DrawTaskT {
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
// llgo:link (*DrawTaskT).DrawGetDependentCount C.lv_draw_get_dependent_count
func (recv_ *DrawTaskT) DrawGetDependentCount() c.Uint32T {
	return 0
}

/**
 * Initialize a layer
 * @param layer pointer to a layer to initialize
 */
// llgo:link (*LayerT).LayerInit C.lv_layer_init
func (recv_ *LayerT) LayerInit() {
}

/**
 * Reset the layer to a drawable state
 * @param layer pointer to a layer to reset
 */
// llgo:link (*LayerT).LayerReset C.lv_layer_reset
func (recv_ *LayerT) LayerReset() {
}

/**
 * Create (allocate) a new layer on a parent layer
 * @param parent_layer      the parent layer to which the layer will be merged when it's rendered
 * @param color_format      the color format of the layer
 * @param area              the areas of the layer (absolute coordinates)
 * @return                  the new target_layer or NULL on error
 */
// llgo:link (*LayerT).DrawLayerCreate C.lv_draw_layer_create
func (recv_ *LayerT) DrawLayerCreate(color_format ColorFormatT, area *AreaT) *LayerT {
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
// llgo:link (*LayerT).DrawLayerInit C.lv_draw_layer_init
func (recv_ *LayerT) DrawLayerInit(parent_layer *LayerT, color_format ColorFormatT, area *AreaT) {
}

/**
 * Try to allocate a buffer for the layer.
 * @param layer             pointer to a layer
 * @return                  pointer to the allocated aligned buffer or NULL on failure
 */
// llgo:link (*LayerT).DrawLayerAllocBuf C.lv_draw_layer_alloc_buf
func (recv_ *LayerT) DrawLayerAllocBuf() c.Pointer {
	return nil
}

/**
 * Got to a pixel at X and Y coordinate on a layer
 * @param layer             pointer to a layer
 * @param x                 the target X coordinate
 * @param y                 the target X coordinate
 * @return                  `buf` offset to point to the given X and Y coordinate
 */
// llgo:link (*LayerT).DrawLayerGoToXy C.lv_draw_layer_go_to_xy
func (recv_ *LayerT) DrawLayerGoToXy(x c.Int32T, y c.Int32T) c.Pointer {
	return nil
}

/**
 * Get the type of a draw task
 * @param t   the draw task to get the type of
 * @return    the draw task type
 */
// llgo:link (*DrawTaskT).DrawTaskGetType C.lv_draw_task_get_type
func (recv_ *DrawTaskT) DrawTaskGetType() DrawTaskTypeT {
	return 0
}

/**
 * Get the draw descriptor of a draw task
 * @param t   the draw task to get the draw descriptor of
 * @return    a void pointer to the draw descriptor
 */
// llgo:link (*DrawTaskT).DrawTaskGetDrawDsc C.lv_draw_task_get_draw_dsc
func (recv_ *DrawTaskT) DrawTaskGetDrawDsc() c.Pointer {
	return nil
}

/**
 * Get the draw area of a draw task
 * @param t      the draw task to get the draw area of
 * @param area   the destination where the draw area will be stored
 */
// llgo:link (*DrawTaskT).DrawTaskGetArea C.lv_draw_task_get_area
func (recv_ *DrawTaskT) DrawTaskGetArea(area *AreaT) {
}

/**
 * Init an array.
 * @param array pointer to an `lv_array_t` variable to initialize
 * @param capacity the initial capacity of the array
 * @param element_size the size of an element in bytes
 */
// llgo:link (*ArrayT).ArrayInit C.lv_array_init
func (recv_ *ArrayT) ArrayInit(capacity c.Uint32T, element_size c.Uint32T) {
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
// llgo:link (*ArrayT).ArrayInitFromBuf C.lv_array_init_from_buf
func (recv_ *ArrayT) ArrayInitFromBuf(buf c.Pointer, capacity c.Uint32T, element_size c.Uint32T) {
}

/**
 * Resize the array to the given capacity.
 * @note if the new capacity is smaller than the current size, the array will be truncated.
 * @param array pointer to an `lv_array_t` variable
 * @param new_capacity the new capacity of the array
 */
// llgo:link (*ArrayT).ArrayResize C.lv_array_resize
func (recv_ *ArrayT) ArrayResize(new_capacity c.Uint32T) bool {
	return false
}

/**
 * Deinit the array, and free the allocated memory
 * @param array pointer to an `lv_array_t` variable to deinitialize
 */
// llgo:link (*ArrayT).ArrayDeinit C.lv_array_deinit
func (recv_ *ArrayT) ArrayDeinit() {
}

/**
 * Copy an array to another.
 * @note this will create a new array with the same capacity and size as the source array.
 * @param target pointer to an `lv_array_t` variable to copy to
 * @param source pointer to an `lv_array_t` variable to copy from
 */
// llgo:link (*ArrayT).ArrayCopy C.lv_array_copy
func (recv_ *ArrayT) ArrayCopy(source *ArrayT) {
}

/**
 * Shrink the memory capacity of array if necessary.
 * @param array pointer to an `lv_array_t` variable
 */
// llgo:link (*ArrayT).ArrayShrink C.lv_array_shrink
func (recv_ *ArrayT) ArrayShrink() {
}

/**
 * Remove the element at the specified position in the array.
 * @param array pointer to an `lv_array_t` variable
 * @param index the index of the element to remove
 * @return LV_RESULT_OK: success, otherwise: error
 */
// llgo:link (*ArrayT).ArrayRemove C.lv_array_remove
func (recv_ *ArrayT) ArrayRemove(index c.Uint32T) ResultT {
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
// llgo:link (*ArrayT).ArrayErase C.lv_array_erase
func (recv_ *ArrayT) ArrayErase(start c.Uint32T, end c.Uint32T) ResultT {
	return 0
}

/**
 * Concatenate two arrays. Adds new elements to the end of the array.
 * @note The destination array is automatically expanded as necessary.
 * @param array pointer to an `lv_array_t` variable
 * @param other pointer to the array to concatenate
 * @return LV_RESULT_OK: success, otherwise: error
 */
// llgo:link (*ArrayT).ArrayConcat C.lv_array_concat
func (recv_ *ArrayT) ArrayConcat(other *ArrayT) ResultT {
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
// llgo:link (*ArrayT).ArrayPushBack C.lv_array_push_back
func (recv_ *ArrayT) ArrayPushBack(element c.Pointer) ResultT {
	return 0
}

/**
 * Assigns one content to the array, replacing its current content.
 * @param array pointer to an `lv_array_t` variable
 * @param index the index of the element to replace
 * @param value pointer to the elements to add
 * @return true: success; false: error
 */
// llgo:link (*ArrayT).ArrayAssign C.lv_array_assign
func (recv_ *ArrayT) ArrayAssign(index c.Uint32T, value c.Pointer) ResultT {
	return 0
}

/**
 * Returns a pointer to the element at position n in the array.
 * @param array pointer to an `lv_array_t` variable
 * @param index the index of the element to return
 * @return a pointer to the requested element, NULL if `index` is out of range
 */
// llgo:link (*ArrayT).ArrayAt C.lv_array_at
func (recv_ *ArrayT) ArrayAt(index c.Uint32T) c.Pointer {
	return nil
}

// llgo:type C
type EventCbT func(*EventT)
type EventCodeT c.Int

const (
	EVENT_ALL                  EventCodeT = 0
	EVENT_PRESSED              EventCodeT = 1
	EVENT_PRESSING             EventCodeT = 2
	EVENT_PRESS_LOST           EventCodeT = 3
	EVENT_SHORT_CLICKED        EventCodeT = 4
	EVENT_SINGLE_CLICKED       EventCodeT = 5
	EVENT_DOUBLE_CLICKED       EventCodeT = 6
	EVENT_TRIPLE_CLICKED       EventCodeT = 7
	EVENT_LONG_PRESSED         EventCodeT = 8
	EVENT_LONG_PRESSED_REPEAT  EventCodeT = 9
	EVENT_CLICKED              EventCodeT = 10
	EVENT_RELEASED             EventCodeT = 11
	EVENT_SCROLL_BEGIN         EventCodeT = 12
	EVENT_SCROLL_THROW_BEGIN   EventCodeT = 13
	EVENT_SCROLL_END           EventCodeT = 14
	EVENT_SCROLL               EventCodeT = 15
	EVENT_GESTURE              EventCodeT = 16
	EVENT_KEY                  EventCodeT = 17
	EVENT_ROTARY               EventCodeT = 18
	EVENT_FOCUSED              EventCodeT = 19
	EVENT_DEFOCUSED            EventCodeT = 20
	EVENT_LEAVE                EventCodeT = 21
	EVENT_HIT_TEST             EventCodeT = 22
	EVENT_INDEV_RESET          EventCodeT = 23
	EVENT_HOVER_OVER           EventCodeT = 24
	EVENT_HOVER_LEAVE          EventCodeT = 25
	EVENT_COVER_CHECK          EventCodeT = 26
	EVENT_REFR_EXT_DRAW_SIZE   EventCodeT = 27
	EVENT_DRAW_MAIN_BEGIN      EventCodeT = 28
	EVENT_DRAW_MAIN            EventCodeT = 29
	EVENT_DRAW_MAIN_END        EventCodeT = 30
	EVENT_DRAW_POST_BEGIN      EventCodeT = 31
	EVENT_DRAW_POST            EventCodeT = 32
	EVENT_DRAW_POST_END        EventCodeT = 33
	EVENT_DRAW_TASK_ADDED      EventCodeT = 34
	EVENT_VALUE_CHANGED        EventCodeT = 35
	EVENT_INSERT               EventCodeT = 36
	EVENT_REFRESH              EventCodeT = 37
	EVENT_READY                EventCodeT = 38
	EVENT_CANCEL               EventCodeT = 39
	EVENT_CREATE               EventCodeT = 40
	EVENT_DELETE               EventCodeT = 41
	EVENT_CHILD_CHANGED        EventCodeT = 42
	EVENT_CHILD_CREATED        EventCodeT = 43
	EVENT_CHILD_DELETED        EventCodeT = 44
	EVENT_SCREEN_UNLOAD_START  EventCodeT = 45
	EVENT_SCREEN_LOAD_START    EventCodeT = 46
	EVENT_SCREEN_LOADED        EventCodeT = 47
	EVENT_SCREEN_UNLOADED      EventCodeT = 48
	EVENT_SIZE_CHANGED         EventCodeT = 49
	EVENT_STYLE_CHANGED        EventCodeT = 50
	EVENT_LAYOUT_CHANGED       EventCodeT = 51
	EVENT_GET_SELF_SIZE        EventCodeT = 52
	EVENT_INVALIDATE_AREA      EventCodeT = 53
	EVENT_RESOLUTION_CHANGED   EventCodeT = 54
	EVENT_COLOR_FORMAT_CHANGED EventCodeT = 55
	EVENT_REFR_REQUEST         EventCodeT = 56
	EVENT_REFR_START           EventCodeT = 57
	EVENT_REFR_READY           EventCodeT = 58
	EVENT_RENDER_START         EventCodeT = 59
	EVENT_RENDER_READY         EventCodeT = 60
	EVENT_FLUSH_START          EventCodeT = 61
	EVENT_FLUSH_FINISH         EventCodeT = 62
	EVENT_FLUSH_WAIT_START     EventCodeT = 63
	EVENT_FLUSH_WAIT_FINISH    EventCodeT = 64
	EVENT_VSYNC                EventCodeT = 65
	EVENT_VSYNC_REQUEST        EventCodeT = 66
	EVENT_LAST                 EventCodeT = 67
	EVENT_PREPROCESS           EventCodeT = 32768
	EVENT_MARKED_DELETING      EventCodeT = 65536
)

type EventListT struct {
	Array             ArrayT
	IsTraversing      c.Uint8T
	HasMarkedDeleting c.Uint8T
}

/**
 * @brief Event callback.
 * Events are used to notify the user of some action being taken on Widget.
 * For details, see ::lv_event_t.
 */
// llgo:link (*EventListT).EventSend C.lv_event_send
func (recv_ *EventListT) EventSend(e *EventT, preprocess bool) ResultT {
	return 0
}

// llgo:link (*EventListT).EventAdd C.lv_event_add
func (recv_ *EventListT) EventAdd(cb EventCbT, filter EventCodeT, user_data c.Pointer) *EventDscT {
	return nil
}

// llgo:link (*EventListT).EventRemoveDsc C.lv_event_remove_dsc
func (recv_ *EventListT) EventRemoveDsc(dsc *EventDscT) bool {
	return false
}

// llgo:link (*EventListT).EventGetCount C.lv_event_get_count
func (recv_ *EventListT) EventGetCount() c.Uint32T {
	return 0
}

// llgo:link (*EventListT).EventGetDsc C.lv_event_get_dsc
func (recv_ *EventListT) EventGetDsc(index c.Uint32T) *EventDscT {
	return nil
}

// llgo:link (*EventDscT).EventDscGetCb C.lv_event_dsc_get_cb
func (recv_ *EventDscT) EventDscGetCb() EventCbT {
	return nil
}

// llgo:link (*EventDscT).EventDscGetUserData C.lv_event_dsc_get_user_data
func (recv_ *EventDscT) EventDscGetUserData() c.Pointer {
	return nil
}

// llgo:link (*EventListT).EventRemove C.lv_event_remove
func (recv_ *EventListT) EventRemove(index c.Uint32T) bool {
	return false
}

// llgo:link (*EventListT).EventRemoveAll C.lv_event_remove_all
func (recv_ *EventListT) EventRemoveAll() {
}

/**
 * Get Widget originally targeted by the event. It's the same even if event was bubbled.
 * @param e     pointer to the event descriptor
 * @return      the target of the event_code
 */
// llgo:link (*EventT).EventGetTarget C.lv_event_get_target
func (recv_ *EventT) EventGetTarget() c.Pointer {
	return nil
}

/**
 * Get current target of the event. It's the Widget for which the event handler being called.
 * If the event is not bubbled it's the same as "normal" target.
 * @param e     pointer to the event descriptor
 * @return      pointer to the current target of the event_code
 */
// llgo:link (*EventT).EventGetCurrentTarget C.lv_event_get_current_target
func (recv_ *EventT) EventGetCurrentTarget() c.Pointer {
	return nil
}

/**
 * Get event code of an event.
 * @param e     pointer to the event descriptor
 * @return      the event code. (E.g. `LV_EVENT_CLICKED`, `LV_EVENT_FOCUSED`, etc)
 */
// llgo:link (*EventT).EventGetCode C.lv_event_get_code
func (recv_ *EventT) EventGetCode() EventCodeT {
	return 0
}

/**
 * Get parameter passed when event was sent.
 * @param e     pointer to the event descriptor
 * @return      pointer to the parameter
 */
// llgo:link (*EventT).EventGetParam C.lv_event_get_param
func (recv_ *EventT) EventGetParam() c.Pointer {
	return nil
}

/**
 * Get user_data passed when event was registered on Widget.
 * @param e     pointer to the event descriptor
 * @return      pointer to the user_data
 */
// llgo:link (*EventT).EventGetUserData C.lv_event_get_user_data
func (recv_ *EventT) EventGetUserData() c.Pointer {
	return nil
}

/**
 * Stop event from bubbling.
 * This is only valid when called in the middle of an event processing chain.
 * @param e     pointer to the event descriptor
 */
// llgo:link (*EventT).EventStopBubbling C.lv_event_stop_bubbling
func (recv_ *EventT) EventStopBubbling() {
}

/**
 * Stop processing this event.
 * This is only valid when called in the middle of an event processing chain.
 * @param e     pointer to the event descriptor
 */
// llgo:link (*EventT).EventStopProcessing C.lv_event_stop_processing
func (recv_ *EventT) EventStopProcessing() {
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
//go:linkname EventRegisterId C.lv_event_register_id
func EventRegisterId() c.Uint32T

/**
 * Get the name of an event code.
 * @param code  the event code
 * @return      the name of the event code as a string
 */
// llgo:link EventCodeT.EventCodeGetName C.lv_event_code_get_name
func (recv_ EventCodeT) EventCodeGetName() *c.Char {
	return nil
}

type DisplayRotationT c.Int

const (
	DISPLAY_ROTATION_0   DisplayRotationT = 0
	DISPLAY_ROTATION_90  DisplayRotationT = 1
	DISPLAY_ROTATION_180 DisplayRotationT = 2
	DISPLAY_ROTATION_270 DisplayRotationT = 3
)

type DisplayRenderModeT c.Int

const (
	DISPLAY_RENDER_MODE_PARTIAL DisplayRenderModeT = 0
	DISPLAY_RENDER_MODE_DIRECT  DisplayRenderModeT = 1
	DISPLAY_RENDER_MODE_FULL    DisplayRenderModeT = 2
)

type ScreenLoadAnimT c.Int

const (
	SCREEN_LOAD_ANIM_NONE        ScreenLoadAnimT = 0
	SCREEN_LOAD_ANIM_OVER_LEFT   ScreenLoadAnimT = 1
	SCREEN_LOAD_ANIM_OVER_RIGHT  ScreenLoadAnimT = 2
	SCREEN_LOAD_ANIM_OVER_TOP    ScreenLoadAnimT = 3
	SCREEN_LOAD_ANIM_OVER_BOTTOM ScreenLoadAnimT = 4
	SCREEN_LOAD_ANIM_MOVE_LEFT   ScreenLoadAnimT = 5
	SCREEN_LOAD_ANIM_MOVE_RIGHT  ScreenLoadAnimT = 6
	SCREEN_LOAD_ANIM_MOVE_TOP    ScreenLoadAnimT = 7
	SCREEN_LOAD_ANIM_MOVE_BOTTOM ScreenLoadAnimT = 8
	SCREEN_LOAD_ANIM_FADE_IN     ScreenLoadAnimT = 9
	SCREEN_LOAD_ANIM_FADE_ON     ScreenLoadAnimT = 9
	SCREEN_LOAD_ANIM_FADE_OUT    ScreenLoadAnimT = 10
	SCREEN_LOAD_ANIM_OUT_LEFT    ScreenLoadAnimT = 11
	SCREEN_LOAD_ANIM_OUT_RIGHT   ScreenLoadAnimT = 12
	SCREEN_LOAD_ANIM_OUT_TOP     ScreenLoadAnimT = 13
	SCREEN_LOAD_ANIM_OUT_BOTTOM  ScreenLoadAnimT = 14
)

// llgo:type C
type DisplayFlushCbT func(*DisplayT, *AreaT, *c.Uint8T)

// llgo:type C
type DisplayFlushWaitCbT func(*DisplayT)

/**
 * Create a new display with the given resolution
 * @param hor_res   horizontal resolution in pixels
 * @param ver_res   vertical resolution in pixels
 * @return          pointer to a display object or `NULL` on error
 */
//go:linkname DisplayCreate C.lv_display_create
func DisplayCreate(hor_res c.Int32T, ver_res c.Int32T) *DisplayT

/**
 * Remove a display
 * @param disp      pointer to display
 */
// llgo:link (*DisplayT).DisplayDelete C.lv_display_delete
func (recv_ *DisplayT) DisplayDelete() {
}

/**
 * Set a default display. The new screens will be created on it by default.
 * @param disp      pointer to a display
 */
// llgo:link (*DisplayT).DisplaySetDefault C.lv_display_set_default
func (recv_ *DisplayT) DisplaySetDefault() {
}

/**
 * Get the default display
 * @return          pointer to the default display
 */
//go:linkname DisplayGetDefault C.lv_display_get_default
func DisplayGetDefault() *DisplayT

/**
 * Get the next display.
 * @param disp      pointer to the current display. NULL to initialize.
 * @return          the next display or NULL if no more. Gives the first display when the parameter is NULL.
 */
// llgo:link (*DisplayT).DisplayGetNext C.lv_display_get_next
func (recv_ *DisplayT) DisplayGetNext() *DisplayT {
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
// llgo:link (*DisplayT).DisplaySetResolution C.lv_display_set_resolution
func (recv_ *DisplayT) DisplaySetResolution(hor_res c.Int32T, ver_res c.Int32T) {
}

/**
 * It's not mandatory to use the whole display for LVGL, however in some cases physical resolution is important.
 * For example the touchpad still sees whole resolution and the values needs to be converted
 * to the active LVGL display area.
 * @param disp      pointer to a display
 * @param hor_res   the new physical horizontal resolution, or -1 to assume it's the same as the normal hor. res.
 * @param ver_res   the new physical vertical resolution, or -1 to assume it's the same as the normal hor. res.
 */
// llgo:link (*DisplayT).DisplaySetPhysicalResolution C.lv_display_set_physical_resolution
func (recv_ *DisplayT) DisplaySetPhysicalResolution(hor_res c.Int32T, ver_res c.Int32T) {
}

/**
 * If physical resolution is not the same as the normal resolution
 * the offset of the active display area can be set here.
 * @param disp      pointer to a display
 * @param x         X offset
 * @param y         Y offset
 */
// llgo:link (*DisplayT).DisplaySetOffset C.lv_display_set_offset
func (recv_ *DisplayT) DisplaySetOffset(x c.Int32T, y c.Int32T) {
}

/**
 * Set the rotation of this display. LVGL will swap the horizontal and vertical resolutions internally.
 * @param disp      pointer to a display (NULL to use the default display)
 * @param rotation  `LV_DISPLAY_ROTATION_0/90/180/270`
 */
// llgo:link (*DisplayT).DisplaySetRotation C.lv_display_set_rotation
func (recv_ *DisplayT) DisplaySetRotation(rotation DisplayRotationT) {
}

/**
 * Use matrix rotation for the display. This function is depended on `LV_DRAW_TRANSFORM_USE_MATRIX`
 * @param disp      pointer to a display (NULL to use the default display)
 * @param enable    true: enable matrix rotation, false: disable
 */
// llgo:link (*DisplayT).DisplaySetMatrixRotation C.lv_display_set_matrix_rotation
func (recv_ *DisplayT) DisplaySetMatrixRotation(enable bool) {
}

/**
 * Set the DPI (dot per inch) of the display.
 * dpi = sqrt(hor_res^2 + ver_res^2) / diagonal"
 * @param disp      pointer to a display
 * @param dpi       the new DPI
 */
// llgo:link (*DisplayT).DisplaySetDpi C.lv_display_set_dpi
func (recv_ *DisplayT) DisplaySetDpi(dpi c.Int32T) {
}

/**
 * Get the horizontal resolution of a display.
 * @param disp      pointer to a display (NULL to use the default display)
 * @return          the horizontal resolution of the display.
 */
// llgo:link (*DisplayT).DisplayGetHorizontalResolution C.lv_display_get_horizontal_resolution
func (recv_ *DisplayT) DisplayGetHorizontalResolution() c.Int32T {
	return 0
}

/**
 * Get the vertical resolution of a display
 * @param disp      pointer to a display (NULL to use the default display)
 * @return          the vertical resolution of the display
 */
// llgo:link (*DisplayT).DisplayGetVerticalResolution C.lv_display_get_vertical_resolution
func (recv_ *DisplayT) DisplayGetVerticalResolution() c.Int32T {
	return 0
}

/**
 * Get the original horizontal resolution of a display without considering rotation
 * @param disp      pointer to a display (NULL to use the default display)
 * @return          the horizontal resolution of the display.
 */
// llgo:link (*DisplayT).DisplayGetOriginalHorizontalResolution C.lv_display_get_original_horizontal_resolution
func (recv_ *DisplayT) DisplayGetOriginalHorizontalResolution() c.Int32T {
	return 0
}

/**
 * Get the original vertical resolution of a display without considering rotation
 * @param disp      pointer to a display (NULL to use the default display)
 * @return          the vertical resolution of the display
 */
// llgo:link (*DisplayT).DisplayGetOriginalVerticalResolution C.lv_display_get_original_vertical_resolution
func (recv_ *DisplayT) DisplayGetOriginalVerticalResolution() c.Int32T {
	return 0
}

/**
 * Get the physical horizontal resolution of a display
 * @param disp      pointer to a display (NULL to use the default display)
 * @return the      physical horizontal resolution of the display
 */
// llgo:link (*DisplayT).DisplayGetPhysicalHorizontalResolution C.lv_display_get_physical_horizontal_resolution
func (recv_ *DisplayT) DisplayGetPhysicalHorizontalResolution() c.Int32T {
	return 0
}

/**
 * Get the physical vertical resolution of a display
 * @param disp      pointer to a display (NULL to use the default display)
 * @return          the physical vertical resolution of the display
 */
// llgo:link (*DisplayT).DisplayGetPhysicalVerticalResolution C.lv_display_get_physical_vertical_resolution
func (recv_ *DisplayT) DisplayGetPhysicalVerticalResolution() c.Int32T {
	return 0
}

/**
 * Get the horizontal offset from the full / physical display
 * @param disp      pointer to a display (NULL to use the default display)
 * @return          the horizontal offset from the physical display
 */
// llgo:link (*DisplayT).DisplayGetOffsetX C.lv_display_get_offset_x
func (recv_ *DisplayT) DisplayGetOffsetX() c.Int32T {
	return 0
}

/**
 * Get the vertical offset from the full / physical display
 * @param disp      pointer to a display (NULL to use the default display)
 * @return          the horizontal offset from the physical display
 */
// llgo:link (*DisplayT).DisplayGetOffsetY C.lv_display_get_offset_y
func (recv_ *DisplayT) DisplayGetOffsetY() c.Int32T {
	return 0
}

/**
 * Get the current rotation of this display.
 * @param disp      pointer to a display (NULL to use the default display)
 * @return          the current rotation
 */
// llgo:link (*DisplayT).DisplayGetRotation C.lv_display_get_rotation
func (recv_ *DisplayT) DisplayGetRotation() DisplayRotationT {
	return 0
}

/**
 * Get if matrix rotation is enabled for a display or not
 * @param disp      pointer to a display (NULL to use the default display)
 * @return          true: matrix rotation is enabled; false: disabled
 */
// llgo:link (*DisplayT).DisplayGetMatrixRotation C.lv_display_get_matrix_rotation
func (recv_ *DisplayT) DisplayGetMatrixRotation() bool {
	return false
}

/**
 * Get the DPI of the display
 * @param disp      pointer to a display (NULL to use the default display)
 * @return          dpi of the display
 */
// llgo:link (*DisplayT).DisplayGetDpi C.lv_display_get_dpi
func (recv_ *DisplayT) DisplayGetDpi() c.Int32T {
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
// llgo:link (*DisplayT).DisplaySetBuffers C.lv_display_set_buffers
func (recv_ *DisplayT) DisplaySetBuffers(buf1 c.Pointer, buf2 c.Pointer, buf_size c.Uint32T, render_mode DisplayRenderModeT) {
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
// llgo:link (*DisplayT).DisplaySetBuffersWithStride C.lv_display_set_buffers_with_stride
func (recv_ *DisplayT) DisplaySetBuffersWithStride(buf1 c.Pointer, buf2 c.Pointer, buf_size c.Uint32T, stride c.Uint32T, render_mode DisplayRenderModeT) {
}

/**
 * Set the buffers for a display, accept a draw buffer pointer.
 * Normally use `lv_display_set_buffers` is enough for most cases.
 * Use this function when an existing lv_draw_buf_t is available.
 * @param disp              pointer to a display
 * @param buf1              first buffer
 * @param buf2              second buffer (can be `NULL`)
 */
// llgo:link (*DisplayT).DisplaySetDrawBuffers C.lv_display_set_draw_buffers
func (recv_ *DisplayT) DisplaySetDrawBuffers(buf1 *DrawBufT, buf2 *DrawBufT) {
}

/**
 * Set the third draw buffer for a display.
 * @param disp              pointer to a display
 * @param buf3              third buffer
 */
// llgo:link (*DisplayT).DisplaySet3rdDrawBuffer C.lv_display_set_3rd_draw_buffer
func (recv_ *DisplayT) DisplaySet3rdDrawBuffer(buf3 *DrawBufT) {
}

/**
 * Set display render mode
 * @param disp              pointer to a display
 * @param render_mode       LV_DISPLAY_RENDER_MODE_PARTIAL/DIRECT/FULL
 */
// llgo:link (*DisplayT).DisplaySetRenderMode C.lv_display_set_render_mode
func (recv_ *DisplayT) DisplaySetRenderMode(render_mode DisplayRenderModeT) {
}

/**
 * Set the flush callback which will be called to copy the rendered image to the display.
 * @param disp      pointer to a display
 * @param flush_cb  the flush callback (`px_map` contains the rendered image as raw pixel map and it should be copied to `area` on the display)
 */
// llgo:link (*DisplayT).DisplaySetFlushCb C.lv_display_set_flush_cb
func (recv_ *DisplayT) DisplaySetFlushCb(flush_cb DisplayFlushCbT) {
}

/**
 * Set a callback to be used while LVGL is waiting flushing to be finished.
 * It can do any complex logic to wait, including semaphores, mutexes, polling flags, etc.
 * If not set the `disp->flushing` flag is used which can be cleared with `lv_display_flush_ready()`
 * @param disp      pointer to a display
 * @param wait_cb   a callback to call while LVGL is waiting for flush ready.
 *                  If NULL `lv_display_flush_ready()` can be used to signal that flushing is ready.
 */
// llgo:link (*DisplayT).DisplaySetFlushWaitCb C.lv_display_set_flush_wait_cb
func (recv_ *DisplayT) DisplaySetFlushWaitCb(wait_cb DisplayFlushWaitCbT) {
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
// llgo:link (*DisplayT).DisplaySetColorFormat C.lv_display_set_color_format
func (recv_ *DisplayT) DisplaySetColorFormat(color_format ColorFormatT) {
}

/**
 * Get the color format of the display
 * @param disp              pointer to a display
 * @return                  the color format
 */
// llgo:link (*DisplayT).DisplayGetColorFormat C.lv_display_get_color_format
func (recv_ *DisplayT) DisplayGetColorFormat() ColorFormatT {
	return 0
}

/**
 * Set the number of tiles for parallel rendering.
 * @param disp              pointer to a display
 * @param tile_cnt          number of tiles (1 =< tile_cnt < 256)
 */
// llgo:link (*DisplayT).DisplaySetTileCnt C.lv_display_set_tile_cnt
func (recv_ *DisplayT) DisplaySetTileCnt(tile_cnt c.Uint32T) {
}

/**
 * Get the number of tiles used for parallel rendering
 * @param disp              pointer to a display
 * @return                  number of tiles
 */
// llgo:link (*DisplayT).DisplayGetTileCnt C.lv_display_get_tile_cnt
func (recv_ *DisplayT) DisplayGetTileCnt() c.Uint32T {
	return 0
}

/**
 * Enable anti-aliasing for the render engine
 * @param disp      pointer to a display
 * @param en        true/false
 */
// llgo:link (*DisplayT).DisplaySetAntialiasing C.lv_display_set_antialiasing
func (recv_ *DisplayT) DisplaySetAntialiasing(en bool) {
}

/**
 * Get if anti-aliasing is enabled for a display or not
 * @param disp      pointer to a display (NULL to use the default display)
 * @return          true/false
 */
// llgo:link (*DisplayT).DisplayGetAntialiasing C.lv_display_get_antialiasing
func (recv_ *DisplayT) DisplayGetAntialiasing() bool {
	return false
}

/**
 * Call from the display driver when the flushing is finished
 * @param disp      pointer to display whose `flush_cb` was called
 */
// llgo:link (*DisplayT).DisplayFlushReady C.lv_display_flush_ready
func (recv_ *DisplayT) DisplayFlushReady() {
}

/**
 * Tell if it's the last area of the refreshing process.
 * Can be called from `flush_cb` to execute some special display refreshing if needed when all areas area flushed.
 * @param disp      pointer to display
 * @return          true: it's the last area to flush;
 *                  false: there are other areas too which will be refreshed soon
 */
// llgo:link (*DisplayT).DisplayFlushIsLast C.lv_display_flush_is_last
func (recv_ *DisplayT) DisplayFlushIsLast() bool {
	return false
}

// llgo:link (*DisplayT).DisplayIsDoubleBuffered C.lv_display_is_double_buffered
func (recv_ *DisplayT) DisplayIsDoubleBuffered() bool {
	return false
}

/**
 * Return a pointer to the active screen on a display
 * @param disp      pointer to display which active screen should be get.
 *                  (NULL to use the default screen)
 * @return          pointer to the active screen object (loaded by 'lv_screen_load()')
 */
// llgo:link (*DisplayT).DisplayGetScreenActive C.lv_display_get_screen_active
func (recv_ *DisplayT) DisplayGetScreenActive() *ObjT {
	return nil
}

/**
 * Return with a pointer to the previous screen. Only used during screen transitions.
 * @param disp      pointer to display which previous screen should be get.
 *                  (NULL to use the default screen)
 * @return          pointer to the previous screen object or NULL if not used now
 */
// llgo:link (*DisplayT).DisplayGetScreenPrev C.lv_display_get_screen_prev
func (recv_ *DisplayT) DisplayGetScreenPrev() *ObjT {
	return nil
}

/**
 * Return the top layer. The top layer is the same on all screens and it is above the normal screen layer.
 * @param disp      pointer to display which top layer should be get. (NULL to use the default screen)
 * @return          pointer to the top layer object
 */
// llgo:link (*DisplayT).DisplayGetLayerTop C.lv_display_get_layer_top
func (recv_ *DisplayT) DisplayGetLayerTop() *ObjT {
	return nil
}

/**
 * Return the sys. layer. The system layer is the same on all screen and it is above the normal screen and the top layer.
 * @param disp      pointer to display which sys. layer should be retrieved. (NULL to use the default screen)
 * @return          pointer to the sys layer object
 */
// llgo:link (*DisplayT).DisplayGetLayerSys C.lv_display_get_layer_sys
func (recv_ *DisplayT) DisplayGetLayerSys() *ObjT {
	return nil
}

/**
 * Return the bottom layer. The bottom layer is the same on all screen and it is under the normal screen layer.
 * It's visible only if the screen is transparent.
 * @param disp      pointer to display (NULL to use the default screen)
 * @return          pointer to the bottom layer object
 */
// llgo:link (*DisplayT).DisplayGetLayerBottom C.lv_display_get_layer_bottom
func (recv_ *DisplayT) DisplayGetLayerBottom() *ObjT {
	return nil
}

/**
 * Load a screen on the default display
 * @param scr       pointer to a screen
 */
// llgo:link (*X_lvObjT).ScreenLoad C.lv_screen_load
func (recv_ *X_lvObjT) ScreenLoad() {
}

/**
 * Switch screen with animation
 * @param scr       pointer to the new screen to load
 * @param anim_type type of the animation from `lv_screen_load_anim_t`, e.g. `LV_SCREEN_LOAD_ANIM_MOVE_LEFT`
 * @param time      time of the animation
 * @param delay     delay before the transition
 * @param auto_del  true: automatically delete the old screen
 */
// llgo:link (*ObjT).ScreenLoadAnim C.lv_screen_load_anim
func (recv_ *ObjT) ScreenLoadAnim(anim_type ScreenLoadAnimT, time c.Uint32T, delay c.Uint32T, auto_del bool) {
}

/**
 * Get the active screen of the default display
 * @return          pointer to the active screen
 */
//go:linkname ScreenActive C.lv_screen_active
func ScreenActive() *ObjT

/**
 * Get the top layer  of the default display
 * @return          pointer to the top layer
 */
//go:linkname LayerTop C.lv_layer_top
func LayerTop() *ObjT

/**
 * Get the system layer  of the default display
 * @return          pointer to the sys layer
 */
//go:linkname LayerSys C.lv_layer_sys
func LayerSys() *ObjT

/**
 * Get the bottom layer  of the default display
 * @return          pointer to the bottom layer
 */
//go:linkname LayerBottom C.lv_layer_bottom
func LayerBottom() *ObjT

/**
 * Add an event handler to the display
 * @param disp          pointer to a display
 * @param event_cb      an event callback
 * @param filter        event code to react or `LV_EVENT_ALL`
 * @param user_data     optional user_data
 */
// llgo:link (*DisplayT).DisplayAddEventCb C.lv_display_add_event_cb
func (recv_ *DisplayT) DisplayAddEventCb(event_cb EventCbT, filter EventCodeT, user_data c.Pointer) {
}

/**
 * Get the number of event attached to a display
 * @param disp          pointer to a display
 * @return              number of events
 */
// llgo:link (*DisplayT).DisplayGetEventCount C.lv_display_get_event_count
func (recv_ *DisplayT) DisplayGetEventCount() c.Uint32T {
	return 0
}

/**
 * Get an event descriptor for an event
 * @param disp          pointer to a display
 * @param index         the index of the event
 * @return              the event descriptor
 */
// llgo:link (*DisplayT).DisplayGetEventDsc C.lv_display_get_event_dsc
func (recv_ *DisplayT) DisplayGetEventDsc(index c.Uint32T) *EventDscT {
	return nil
}

/**
 * Remove an event
 * @param disp          pointer to a display
 * @param index         the index of the event to remove
 * @return              true: and event was removed; false: no event was removed
 */
// llgo:link (*DisplayT).DisplayDeleteEvent C.lv_display_delete_event
func (recv_ *DisplayT) DisplayDeleteEvent(index c.Uint32T) bool {
	return false
}

/**
 * Remove an event_cb with user_data
 * @param disp          pointer to a display
 * @param event_cb      the event_cb of the event to remove
 * @param user_data     user_data
 * @return              the count of the event removed
 */
// llgo:link (*DisplayT).DisplayRemoveEventCbWithUserData C.lv_display_remove_event_cb_with_user_data
func (recv_ *DisplayT) DisplayRemoveEventCbWithUserData(event_cb EventCbT, user_data c.Pointer) c.Uint32T {
	return 0
}

/**
 * Send an event to a display
 * @param disp          pointer to a display
 * @param code          an event code. LV_EVENT_...
 * @param param         optional param
 * @return              LV_RESULT_OK: disp wasn't deleted in the event.
 */
// llgo:link (*DisplayT).DisplaySendEvent C.lv_display_send_event
func (recv_ *DisplayT) DisplaySendEvent(code EventCodeT, param c.Pointer) ResultT {
	return 0
}

/**
 * Get the area to be invalidated. Can be used in `LV_EVENT_INVALIDATE_AREA`
 * @param e     pointer to an event
 * @return      the area to invalidated (can be modified as required)
 */
// llgo:link (*EventT).EventGetInvalidatedArea C.lv_event_get_invalidated_area
func (recv_ *EventT) EventGetInvalidatedArea() *AreaT {
	return nil
}

/**
 * Set the theme of a display. If there are no user created widgets yet the screens' theme will be updated
 * @param disp      pointer to a display
 * @param th        pointer to a theme
 */
// llgo:link (*DisplayT).DisplaySetTheme C.lv_display_set_theme
func (recv_ *DisplayT) DisplaySetTheme(th *ThemeT) {
}

/**
 * Get the theme of a display
 * @param disp      pointer to a display
 * @return          the display's theme (can be NULL)
 */
// llgo:link (*DisplayT).DisplayGetTheme C.lv_display_get_theme
func (recv_ *DisplayT) DisplayGetTheme() *ThemeT {
	return nil
}

/**
 * Get elapsed time since last user activity on a display (e.g. click)
 * @param disp      pointer to a display (NULL to get the overall smallest inactivity)
 * @return          elapsed ticks (milliseconds) since the last activity
 */
// llgo:link (*DisplayT).DisplayGetInactiveTime C.lv_display_get_inactive_time
func (recv_ *DisplayT) DisplayGetInactiveTime() c.Uint32T {
	return 0
}

/**
 * Manually trigger an activity on a display
 * @param disp      pointer to a display (NULL to use the default display)
 */
// llgo:link (*DisplayT).DisplayTriggerActivity C.lv_display_trigger_activity
func (recv_ *DisplayT) DisplayTriggerActivity() {
}

/**
 * Temporarily enable and disable the invalidation of the display.
 * @param disp      pointer to a display (NULL to use the default display)
 * @param en        true: enable invalidation; false: invalidation
 */
// llgo:link (*DisplayT).DisplayEnableInvalidation C.lv_display_enable_invalidation
func (recv_ *DisplayT) DisplayEnableInvalidation(en bool) {
}

/**
 * Get display invalidation is enabled.
 * @param disp      pointer to a display (NULL to use the default display)
 * @return return   true if invalidation is enabled
 */
// llgo:link (*DisplayT).DisplayIsInvalidationEnabled C.lv_display_is_invalidation_enabled
func (recv_ *DisplayT) DisplayIsInvalidationEnabled() bool {
	return false
}

/**
 * Get a pointer to the screen refresher timer to
 * modify its parameters with `lv_timer_...` functions.
 * @param disp      pointer to a display
 * @return          pointer to the display refresher timer. (NULL on error)
 */
// llgo:link (*DisplayT).DisplayGetRefrTimer C.lv_display_get_refr_timer
func (recv_ *DisplayT) DisplayGetRefrTimer() *TimerT {
	return nil
}

/**
 * Delete screen refresher timer
 * @param disp      pointer to a display
 */
// llgo:link (*DisplayT).DisplayDeleteRefrTimer C.lv_display_delete_refr_timer
func (recv_ *DisplayT) DisplayDeleteRefrTimer() {
}

/**
 * Register vsync event of a display. `LV_EVENT_VSYNC` event will be sent periodically.
 * Please don't use it in display event listeners, as it may cause memory leaks and illegal access issues.
 *
 * @param disp      pointer to a display
 * @param event_cb      an event callback
 * @param user_data     optional user_data
 */
// llgo:link (*DisplayT).DisplayRegisterVsyncEvent C.lv_display_register_vsync_event
func (recv_ *DisplayT) DisplayRegisterVsyncEvent(event_cb EventCbT, user_data c.Pointer) bool {
	return false
}

/**
 * Unregister vsync event of a display. `LV_EVENT_VSYNC` event won't be sent periodically.
 * Please don't use it in display event listeners, as it may cause memory leaks and illegal access issues.
 * @param disp      pointer to a display
 * @param event_cb      an event callback
 * @param user_data     optional user_data
 */
// llgo:link (*DisplayT).DisplayUnregisterVsyncEvent C.lv_display_unregister_vsync_event
func (recv_ *DisplayT) DisplayUnregisterVsyncEvent(event_cb EventCbT, user_data c.Pointer) bool {
	return false
}

/**
 * Send an vsync event to a display
 * @param disp          pointer to a display
 * @param param         optional param
 * @return              LV_RESULT_OK: disp wasn't deleted in the event.
 */
// llgo:link (*DisplayT).DisplaySendVsyncEvent C.lv_display_send_vsync_event
func (recv_ *DisplayT) DisplaySendVsyncEvent(param c.Pointer) ResultT {
	return 0
}

// llgo:link (*DisplayT).DisplaySetUserData C.lv_display_set_user_data
func (recv_ *DisplayT) DisplaySetUserData(user_data c.Pointer) {
}

// llgo:link (*DisplayT).DisplaySetDriverData C.lv_display_set_driver_data
func (recv_ *DisplayT) DisplaySetDriverData(driver_data c.Pointer) {
}

// llgo:link (*DisplayT).DisplayGetUserData C.lv_display_get_user_data
func (recv_ *DisplayT) DisplayGetUserData() c.Pointer {
	return nil
}

// llgo:link (*DisplayT).DisplayGetDriverData C.lv_display_get_driver_data
func (recv_ *DisplayT) DisplayGetDriverData() c.Pointer {
	return nil
}

// llgo:link (*DisplayT).DisplayGetBufActive C.lv_display_get_buf_active
func (recv_ *DisplayT) DisplayGetBufActive() *DrawBufT {
	return nil
}

/**
 * Rotate an area in-place according to the display's rotation
 * @param disp      pointer to a display
 * @param area      pointer to an area to rotate
 */
// llgo:link (*DisplayT).DisplayRotateArea C.lv_display_rotate_area
func (recv_ *DisplayT) DisplayRotateArea(area *AreaT) {
}

/**
 * Get the size of the draw buffers
 * @param disp      pointer to a display
 * @return          the size of the draw buffer in bytes for valid display, 0 otherwise
 */
// llgo:link (*DisplayT).DisplayGetDrawBufSize C.lv_display_get_draw_buf_size
func (recv_ *DisplayT) DisplayGetDrawBufSize() c.Uint32T {
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
// llgo:link (*DisplayT).DisplayGetInvalidatedDrawBufSize C.lv_display_get_invalidated_draw_buf_size
func (recv_ *DisplayT) DisplayGetInvalidatedDrawBufSize(width c.Uint32T, height c.Uint32T) c.Uint32T {
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
//go:linkname Dpx C.lv_dpx
func Dpx(n c.Int32T) c.Int32T

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
// llgo:link (*DisplayT).DisplayDpx C.lv_display_dpx
func (recv_ *DisplayT) DisplayDpx(n c.Int32T) c.Int32T {
	return 0
}

type MutexT c.Int
type ThreadT c.Int
type ThreadSyncT c.Int
type ThreadPrioT c.Int

const (
	THREAD_PRIO_LOWEST  ThreadPrioT = 0
	THREAD_PRIO_LOW     ThreadPrioT = 1
	THREAD_PRIO_MID     ThreadPrioT = 2
	THREAD_PRIO_HIGH    ThreadPrioT = 3
	THREAD_PRIO_HIGHEST ThreadPrioT = 4
)

/**
 * Set it for `LV_SYSMON_GET_IDLE` to show the CPU usage
 * @return the idle percentage since the last call
 */
//go:linkname OsGetIdlePercent C.lv_os_get_idle_percent
func OsGetIdlePercent() c.Uint32T

// llgo:type C
type DrawImageCoreCb func(*DrawTaskT, *DrawImageDscT, *ImageDecoderDscT, *DrawImageSupT, *AreaT, *AreaT)

/**
 * Initialize an image draw descriptor.
 * @param dsc       pointer to a draw descriptor
 */
// llgo:link (*DrawImageDscT).DrawImageDscInit C.lv_draw_image_dsc_init
func (recv_ *DrawImageDscT) DrawImageDscInit() {
}

/**
 * Try to get an image draw descriptor from a draw task.
 * @param task      draw task
 * @return          the task's draw descriptor or NULL if the task is not of type LV_DRAW_TASK_TYPE_IMAGE
 */
// llgo:link (*DrawTaskT).DrawTaskGetImageDsc C.lv_draw_task_get_image_dsc
func (recv_ *DrawTaskT) DrawTaskGetImageDsc() *DrawImageDscT {
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
// llgo:link (*LayerT).DrawImage C.lv_draw_image
func (recv_ *LayerT) DrawImage(dsc *DrawImageDscT, coords *AreaT) {
}

/**
 * Create a draw task to blend a layer to another layer
 * @param layer         pointer to a layer
 * @param dsc           pointer to an initialized draw descriptor. `src` must be set to the layer to blend
 * @param coords        the coordinates of the layer.
 * @note                `coords` can be small than the total widget area from which the layer is created
 *                      (if only a part of the widget was rendered to a layer)
 */
// llgo:link (*LayerT).DrawLayer C.lv_draw_layer
func (recv_ *LayerT) DrawLayer(dsc *DrawImageDscT, coords *AreaT) {
}

/**
 * Get the type of an image source
 * @param src pointer to an image source:
 *  - pointer to an 'lv_image_t' variable (image stored internally and compiled into the code)
 *  - a path to a file (e.g. "S:/folder/image.bin")
 *  - or a symbol (e.g. LV_SYMBOL_CLOSE)
 * @return type of the image source LV_IMAGE_SRC_VARIABLE/FILE/SYMBOL/UNKNOWN
 */
//go:linkname ImageSrcGetType C.lv_image_src_get_type
func ImageSrcGetType(src c.Pointer) ImageSrcT

/**********************
 *      TYPEDEFS
 **********************/

type DrawRectDscT struct {
	Base              DrawDscBaseT
	Radius            c.Int32T
	BgImageSrc        c.Pointer
	BgImageSymbolFont c.Pointer
	BgImageRecolor    ColorT
	BgImageOpa        OpaT
	BgImageRecolorOpa OpaT
	BgImageTiled      c.Uint8T
	BgOpa             OpaT
	BorderOpa         OpaT
	OutlineOpa        OpaT
	ShadowOpa         OpaT
	BgColor           ColorT
	BgGrad            GradDscT
	BorderColor       ColorT
	BorderWidth       c.Int32T
	BorderSide        BorderSideT
	BorderPost        c.Uint8T
	OutlineColor      ColorT
	OutlineWidth      c.Int32T
	OutlinePad        c.Int32T
	ShadowColor       ColorT
	ShadowWidth       c.Int32T
	ShadowOffsetX     c.Int32T
	ShadowOffsetY     c.Int32T
	ShadowSpread      c.Int32T
}

type DrawFillDscT struct {
	Base   DrawDscBaseT
	Radius c.Int32T
	Opa    OpaT
	Color  ColorT
	Grad   GradDscT
}

type DrawBorderDscT struct {
	Base   DrawDscBaseT
	Radius c.Int32T
	Color  ColorT
	Width  c.Int32T
	Opa    OpaT
	Side   BorderSideT
}

type DrawBoxShadowDscT struct {
	Base    DrawDscBaseT
	Radius  c.Int32T
	Color   ColorT
	Width   c.Int32T
	Spread  c.Int32T
	OfsX    c.Int32T
	OfsY    c.Int32T
	Opa     OpaT
	BgCover c.Uint8T
}

// llgo:link (*DrawRectDscT).DrawRectDscInit C.lv_draw_rect_dsc_init
func (recv_ *DrawRectDscT) DrawRectDscInit() {
}

/**
 * Initialize a fill draw descriptor.
 * @param dsc       pointer to a draw descriptor
 */
// llgo:link (*DrawFillDscT).DrawFillDscInit C.lv_draw_fill_dsc_init
func (recv_ *DrawFillDscT) DrawFillDscInit() {
}

/**
 * Try to get a fill draw descriptor from a draw task.
 * @param task      draw task
 * @return          the task's draw descriptor or NULL if the task is not of type LV_DRAW_TASK_TYPE_FILL
 */
// llgo:link (*DrawTaskT).DrawTaskGetFillDsc C.lv_draw_task_get_fill_dsc
func (recv_ *DrawTaskT) DrawTaskGetFillDsc() *DrawFillDscT {
	return nil
}

/**
 * Fill an area
 * @param layer         pointer to a layer
 * @param dsc           pointer to an initialized draw descriptor variable
 * @param coords        the coordinates of the rectangle
 */
// llgo:link (*LayerT).DrawFill C.lv_draw_fill
func (recv_ *LayerT) DrawFill(dsc *DrawFillDscT, coords *AreaT) {
}

/**
 * Initialize a border draw descriptor.
 * @param dsc       pointer to a draw descriptor
 */
// llgo:link (*DrawBorderDscT).DrawBorderDscInit C.lv_draw_border_dsc_init
func (recv_ *DrawBorderDscT) DrawBorderDscInit() {
}

/**
 * Try to get a border draw descriptor from a draw task.
 * @param task      draw task
 * @return          the task's draw descriptor or NULL if the task is not of type LV_DRAW_TASK_TYPE_BORDER
 */
// llgo:link (*DrawTaskT).DrawTaskGetBorderDsc C.lv_draw_task_get_border_dsc
func (recv_ *DrawTaskT) DrawTaskGetBorderDsc() *DrawBorderDscT {
	return nil
}

/**
 * Draw a border
 * @param layer         pointer to a layer
 * @param dsc           pointer to an initialized draw descriptor variable
 * @param coords        the coordinates of the rectangle
 */
// llgo:link (*LayerT).DrawBorder C.lv_draw_border
func (recv_ *LayerT) DrawBorder(dsc *DrawBorderDscT, coords *AreaT) {
}

/**
 * Initialize a box shadow draw descriptor.
 * @param dsc       pointer to a draw descriptor
 */
// llgo:link (*DrawBoxShadowDscT).DrawBoxShadowDscInit C.lv_draw_box_shadow_dsc_init
func (recv_ *DrawBoxShadowDscT) DrawBoxShadowDscInit() {
}

/**
 * Try to get a box shadow draw descriptor from a draw task.
 * @param task      draw task
 * @return          the task's draw descriptor or NULL if the task is not of type LV_DRAW_TASK_TYPE_BOX_SHADOW
 */
// llgo:link (*DrawTaskT).DrawTaskGetBoxShadowDsc C.lv_draw_task_get_box_shadow_dsc
func (recv_ *DrawTaskT) DrawTaskGetBoxShadowDsc() *DrawBoxShadowDscT {
	return nil
}

/**
 * Draw a box shadow
 * @param layer         pointer to a layer
 * @param dsc           pointer to an initialized draw descriptor variable
 * @param coords        the coordinates of the rectangle
 */
// llgo:link (*LayerT).DrawBoxShadow C.lv_draw_box_shadow
func (recv_ *LayerT) DrawBoxShadow(dsc *DrawBoxShadowDscT, coords *AreaT) {
}

/**
 * The rectangle is a wrapper for fill, border, bg. image and box shadow.
 * Internally fill, border, image and box shadow draw tasks will be created.
 * @param layer         pointer to a layer
 * @param dsc           pointer to an initialized draw descriptor variable
 * @param coords        the coordinates of the rectangle
 */
// llgo:link (*LayerT).DrawRect C.lv_draw_rect
func (recv_ *LayerT) DrawRect(dsc *DrawRectDscT, coords *AreaT) {
}

/**********************
 *      TYPEDEFS
 **********************/

type DrawTriangleDscT struct {
	Base  DrawDscBaseT
	P     [3]PointPreciseT
	Color ColorT
	Opa   OpaT
	Grad  GradDscT
}

/**
 * Initialize a triangle draw descriptor
 * @param draw_dsc  pointer to a draw descriptor
 */
// llgo:link (*DrawTriangleDscT).DrawTriangleDscInit C.lv_draw_triangle_dsc_init
func (recv_ *DrawTriangleDscT) DrawTriangleDscInit() {
}

/**
 * Try to get a triangle draw descriptor from a draw task.
 * @param task      draw task
 * @return          the task's draw descriptor or NULL if the task is not of type LV_DRAW_TASK_TYPE_TRIANGLE
 */
// llgo:link (*DrawTaskT).DrawTaskGetTriangleDsc C.lv_draw_task_get_triangle_dsc
func (recv_ *DrawTaskT) DrawTaskGetTriangleDsc() *DrawTriangleDscT {
	return nil
}

/**
 * Create a triangle draw task
 * @param layer     pointer to a layer
 * @param draw_dsc  pointer to an initialized `lv_draw_triangle_dsc_t` object
 */
// llgo:link (*LayerT).DrawTriangle C.lv_draw_triangle
func (recv_ *LayerT) DrawTriangle(draw_dsc *DrawTriangleDscT) {
}

/**********************
 *      TYPEDEFS
 **********************/

type DrawLabelDscT struct {
	Base               DrawDscBaseT
	Text               *c.Char
	TextSize           PointT
	Font               *FontT
	Color              ColorT
	LineSpace          c.Int32T
	LetterSpace        c.Int32T
	OfsX               c.Int32T
	OfsY               c.Int32T
	Rotation           c.Int32T
	SelStart           c.Uint32T
	SelEnd             c.Uint32T
	SelColor           ColorT
	SelBgColor         ColorT
	TextLength         c.Uint32T
	Align              TextAlignT
	BidiDir            BaseDirT
	Opa                OpaT
	OutlineStrokeOpa   OpaT
	Decor              TextDecorT
	Flag               TextFlagT
	TextLocal          c.Uint8T
	TextStatic         c.Uint8T
	HasBided           c.Uint8T
	Hint               *DrawLabelHintT
	OutlineStrokeColor ColorT
	OutlineStrokeWidth c.Int32T
}

type DrawLetterDscT struct {
	Base               DrawDscBaseT
	Unicode            c.Uint32T
	Font               *FontT
	Color              ColorT
	Rotation           c.Int32T
	ScaleX             c.Int32T
	ScaleY             c.Int32T
	SkewX              c.Int32T
	SkewY              c.Int32T
	Pivot              PointT
	Opa                OpaT
	Decor              TextDecorT
	BlendMode          BlendModeT
	OutlineStrokeOpa   OpaT
	OutlineStrokeWidth c.Int32T
	OutlineStrokeColor ColorT
}

// llgo:type C
type DrawGlyphCbT func(*DrawTaskT, *DrawGlyphDscT, *DrawFillDscT, *AreaT)

// llgo:link (*DrawLetterDscT).DrawLetterDscInit C.lv_draw_letter_dsc_init
func (recv_ *DrawLetterDscT) DrawLetterDscInit() {
}

// llgo:link (*DrawLabelDscT).DrawLabelDscInit C.lv_draw_label_dsc_init
func (recv_ *DrawLabelDscT) DrawLabelDscInit() {
}

/**
 * Try to get a label draw descriptor from a draw task.
 * @param task      draw task
 * @return          the task's draw descriptor or NULL if the task is not of type LV_DRAW_TASK_TYPE_LABEL
 */
// llgo:link (*DrawTaskT).DrawTaskGetLabelDsc C.lv_draw_task_get_label_dsc
func (recv_ *DrawTaskT) DrawTaskGetLabelDsc() *DrawLabelDscT {
	return nil
}

/**
 * Initialize a glyph draw descriptor.
 * Used internally.
 * @param dsc       pointer to a draw descriptor
 */
// llgo:link (*DrawGlyphDscT).DrawGlyphDscInit C.lv_draw_glyph_dsc_init
func (recv_ *DrawGlyphDscT) DrawGlyphDscInit() {
}

// llgo:link (*LayerT).DrawLabel C.lv_draw_label
func (recv_ *LayerT) DrawLabel(dsc *DrawLabelDscT, coords *AreaT) {
}

// llgo:link (*LayerT).DrawCharacter C.lv_draw_character
func (recv_ *LayerT) DrawCharacter(dsc *DrawLabelDscT, point *PointT, unicode_letter c.Uint32T) {
}

// llgo:link (*LayerT).DrawLetter C.lv_draw_letter
func (recv_ *LayerT) DrawLetter(dsc *DrawLetterDscT, point *PointT) {
}

/**
 * Should be used during rendering the characters to get the position and other
 * parameters of the characters
 * @param t             pointer to a draw task
 * @param dsc           pointer to draw descriptor
 * @param coords        coordinates of the label
 * @param cb            a callback to call to draw each glyphs one by one
 */
// llgo:link (*DrawTaskT).DrawLabelIterateCharacters C.lv_draw_label_iterate_characters
func (recv_ *DrawTaskT) DrawLabelIterateCharacters(dsc *DrawLabelDscT, coords *AreaT, cb DrawGlyphCbT) {
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
// llgo:link (*DrawTaskT).DrawUnitDrawLetter C.lv_draw_unit_draw_letter
func (recv_ *DrawTaskT) DrawUnitDrawLetter(dsc *DrawGlyphDscT, pos *PointT, font *FontT, letter c.Uint32T, cb DrawGlyphCbT) {
}

/**********************
 *      TYPEDEFS
 **********************/

type DrawLineDscT struct {
	Base       DrawDscBaseT
	P1         PointPreciseT
	P2         PointPreciseT
	Color      ColorT
	Width      c.Int32T
	DashWidth  c.Int32T
	DashGap    c.Int32T
	Opa        OpaT
	RoundStart c.Uint8T
	RoundEnd   c.Uint8T
	RawEnd     c.Uint8T
}

/**
 * Initialize a line draw descriptor
 * @param dsc       pointer to a draw descriptor
 */
// llgo:link (*DrawLineDscT).DrawLineDscInit C.lv_draw_line_dsc_init
func (recv_ *DrawLineDscT) DrawLineDscInit() {
}

/**
 * Try to get a line draw descriptor from a draw task.
 * @param task      draw task
 * @return          the task's draw descriptor or NULL if the task is not of type LV_DRAW_TASK_TYPE_LINE
 */
// llgo:link (*DrawTaskT).DrawTaskGetLineDsc C.lv_draw_task_get_line_dsc
func (recv_ *DrawTaskT) DrawTaskGetLineDsc() *DrawLineDscT {
	return nil
}

/**
 * Create a line draw task
 * @param layer     pointer to a layer
 * @param dsc       pointer to an initialized `lv_draw_line_dsc_t` variable
 */
// llgo:link (*LayerT).DrawLine C.lv_draw_line
func (recv_ *LayerT) DrawLine(dsc *DrawLineDscT) {
}

/**********************
 *      TYPEDEFS
 **********************/

type DrawArcDscT struct {
	Base       DrawDscBaseT
	Color      ColorT
	Width      c.Int32T
	StartAngle ValuePreciseT
	EndAngle   ValuePreciseT
	Center     PointT
	ImgSrc     c.Pointer
	Radius     c.Uint16T
	Opa        OpaT
	Rounded    c.Uint8T
}

/**
 * Initialize an arc draw descriptor.
 * @param dsc       pointer to a draw descriptor
 */
// llgo:link (*DrawArcDscT).DrawArcDscInit C.lv_draw_arc_dsc_init
func (recv_ *DrawArcDscT) DrawArcDscInit() {
}

/**
 * Try to get an arc draw descriptor from a draw task.
 * @param task      draw task
 * @return          the task's draw descriptor or NULL if the task is not of type LV_DRAW_TASK_TYPE_ARC
 */
// llgo:link (*DrawTaskT).DrawTaskGetArcDsc C.lv_draw_task_get_arc_dsc
func (recv_ *DrawTaskT) DrawTaskGetArcDsc() *DrawArcDscT {
	return nil
}

/**
 * Create an arc draw task.
 * @param layer         pointer to a layer
 * @param dsc           pointer to an initialized draw descriptor variable
 */
// llgo:link (*LayerT).DrawArc C.lv_draw_arc
func (recv_ *LayerT) DrawArc(dsc *DrawArcDscT) {
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
//go:linkname DrawArcGetArea C.lv_draw_arc_get_area
func DrawArcGetArea(x c.Int32T, y c.Int32T, radius c.Uint16T, start_angle ValuePreciseT, end_angle ValuePreciseT, w c.Int32T, rounded bool, area *AreaT)

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
//go:linkname DrawSwI1ToArgb8888 C.lv_draw_sw_i1_to_argb8888
func DrawSwI1ToArgb8888(buf_i1 c.Pointer, buf_argb8888 c.Pointer, width c.Uint32T, height c.Uint32T, buf_i1_stride c.Uint32T, buf_argb8888_stride c.Uint32T, index0_color c.Uint32T, index1_color c.Uint32T)

/**
 * Swap the upper and lower byte of an RGB565 buffer.
 * Might be required if a 8bit parallel port or an SPI port send the bytes in the wrong order.
 * The bytes will be swapped in place.
 * @param buf           pointer to buffer
 * @param buf_size_px   number of pixels in the buffer
 */
//go:linkname DrawSwRgb565Swap C.lv_draw_sw_rgb565_swap
func DrawSwRgb565Swap(buf c.Pointer, buf_size_px c.Uint32T)

/**
 * Invert a draw buffer in the I1 color format.
 * Conventionally, a bit is set to 1 during blending if the luminance is greater than 127.
 * Depending on the display controller used, you might want to have different behavior.
 * The inversion will be performed in place.
 * @param buf          pointer to the buffer to be inverted
 * @param buf_size     size of the buffer in bytes
 */
//go:linkname DrawSwI1Invert C.lv_draw_sw_i1_invert
func DrawSwI1Invert(buf c.Pointer, buf_size c.Uint32T)

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
//go:linkname DrawSwI1ConvertToVtiled C.lv_draw_sw_i1_convert_to_vtiled
func DrawSwI1ConvertToVtiled(buf c.Pointer, buf_size c.Uint32T, width c.Uint32T, height c.Uint32T, out_buf c.Pointer, out_buf_size c.Uint32T, bit_order_lsb bool)

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
//go:linkname DrawSwRotate C.lv_draw_sw_rotate
func DrawSwRotate(src c.Pointer, dest c.Pointer, src_width c.Int32T, src_height c.Int32T, src_stride c.Int32T, dest_stride c.Int32T, rotation DisplayRotationT, color_format ColorFormatT)

type DrawSwMaskResT c.Int

const (
	DRAW_SW_MASK_RES_TRANSP     DrawSwMaskResT = 0
	DRAW_SW_MASK_RES_FULL_COVER DrawSwMaskResT = 1
	DRAW_SW_MASK_RES_CHANGED    DrawSwMaskResT = 2
	DRAW_SW_MASK_RES_UNKNOWN    DrawSwMaskResT = 3
)

type DrawSwMaskTypeT c.Int

const (
	DRAW_SW_MASK_TYPE_LINE   DrawSwMaskTypeT = 0
	DRAW_SW_MASK_TYPE_ANGLE  DrawSwMaskTypeT = 1
	DRAW_SW_MASK_TYPE_RADIUS DrawSwMaskTypeT = 2
	DRAW_SW_MASK_TYPE_FADE   DrawSwMaskTypeT = 3
	DRAW_SW_MASK_TYPE_MAP    DrawSwMaskTypeT = 4
)

type DrawSwMaskLineSideT c.Int

const (
	DRAW_SW_MASK_LINE_SIDE_LEFT   DrawSwMaskLineSideT = 0
	DRAW_SW_MASK_LINE_SIDE_RIGHT  DrawSwMaskLineSideT = 1
	DRAW_SW_MASK_LINE_SIDE_TOP    DrawSwMaskLineSideT = 2
	DRAW_SW_MASK_LINE_SIDE_BOTTOM DrawSwMaskLineSideT = 3
)

// llgo:type C
type DrawSwMaskXcbT func(*OpaT, c.Int32T, c.Int32T, c.Int32T, c.Pointer) DrawSwMaskResT

/**********************
 * GLOBAL PROTOTYPES
 **********************/
//go:linkname DrawSwMaskInit C.lv_draw_sw_mask_init
func DrawSwMaskInit()

//go:linkname DrawSwMaskDeinit C.lv_draw_sw_mask_deinit
func DrawSwMaskDeinit()

//go:linkname DrawSwMaskApply C.lv_draw_sw_mask_apply
func DrawSwMaskApply(masks *c.Pointer, mask_buf *OpaT, abs_x c.Int32T, abs_y c.Int32T, len c.Int32T) DrawSwMaskResT

/**
 * Free the data from the parameter.
 * It's called inside `lv_draw_sw_mask_remove_id` and `lv_draw_sw_mask_remove_custom`
 * Needs to be called only in special cases when the mask is not added by `lv_draw_mask_add`
 * and not removed by `lv_draw_mask_remove_id` or `lv_draw_mask_remove_custom`
 * @param p pointer to a mask parameter
 */
//go:linkname DrawSwMaskFreeParam C.lv_draw_sw_mask_free_param
func DrawSwMaskFreeParam(p c.Pointer)

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
// llgo:link (*DrawSwMaskLineParamT).DrawSwMaskLinePointsInit C.lv_draw_sw_mask_line_points_init
func (recv_ *DrawSwMaskLineParamT) DrawSwMaskLinePointsInit(p1x c.Int32T, p1y c.Int32T, p2x c.Int32T, p2y c.Int32T, side DrawSwMaskLineSideT) {
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
// llgo:link (*DrawSwMaskLineParamT).DrawSwMaskLineAngleInit C.lv_draw_sw_mask_line_angle_init
func (recv_ *DrawSwMaskLineParamT) DrawSwMaskLineAngleInit(px c.Int32T, py c.Int32T, angle c.Int16T, side DrawSwMaskLineSideT) {
}

/**
 * Initialize an angle mask.
 * @param param pointer to a `lv_draw_mask_param_t` to initialize
 * @param vertex_x X coordinate of the angle vertex (absolute coordinates)
 * @param vertex_y Y coordinate of the angle vertex (absolute coordinates)
 * @param start_angle start angle in degrees. 0 deg on the right, 90 deg, on the bottom
 * @param end_angle end angle
 */
// llgo:link (*DrawSwMaskAngleParamT).DrawSwMaskAngleInit C.lv_draw_sw_mask_angle_init
func (recv_ *DrawSwMaskAngleParamT) DrawSwMaskAngleInit(vertex_x c.Int32T, vertex_y c.Int32T, start_angle c.Int32T, end_angle c.Int32T) {
}

/**
 * Initialize a fade mask.
 * @param param pointer to an `lv_draw_mask_radius_param_t` to initialize
 * @param rect coordinates of the rectangle to affect (absolute coordinates)
 * @param radius radius of the rectangle
 * @param inv true: keep the pixels inside the rectangle; keep the pixels outside of the rectangle
 */
// llgo:link (*DrawSwMaskRadiusParamT).DrawSwMaskRadiusInit C.lv_draw_sw_mask_radius_init
func (recv_ *DrawSwMaskRadiusParamT) DrawSwMaskRadiusInit(rect *AreaT, radius c.Int32T, inv bool) {
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
// llgo:link (*DrawSwMaskFadeParamT).DrawSwMaskFadeInit C.lv_draw_sw_mask_fade_init
func (recv_ *DrawSwMaskFadeParamT) DrawSwMaskFadeInit(coords *AreaT, opa_top OpaT, y_top c.Int32T, opa_bottom OpaT, y_bottom c.Int32T) {
}

/**
 * Initialize a map mask.
 * @param param pointer to a `lv_draw_mask_param_t` to initialize
 * @param coords coordinates of the map (absolute coordinates)
 * @param map array of bytes with the mask values
 */
// llgo:link (*DrawSwMaskMapParamT).DrawSwMaskMapInit C.lv_draw_sw_mask_map_init
func (recv_ *DrawSwMaskMapParamT) DrawSwMaskMapInit(coords *AreaT, map_ *OpaT) {
}

// llgo:type C
type DrawSwBlendHandlerT func(*DrawTaskT, *DrawSwBlendDscT)

type DrawSwCustomBlendHandlerT struct {
	DestCf  ColorFormatT
	Handler DrawSwBlendHandlerT
}

/**
 * Call the blend function of the `layer`.
 * @param draw_unit     pointer to a draw unit
 * @param dsc           pointer to an initialized blend descriptor
 */
// llgo:link (*DrawTaskT).DrawSwBlend C.lv_draw_sw_blend
func (recv_ *DrawTaskT) DrawSwBlend(dsc *DrawSwBlendDscT) {
}

/**
 * Initialize the SW renderer. Called in internally.
 * It creates as many SW renderers as defined in LV_DRAW_SW_DRAW_UNIT_CNT
 */
//go:linkname DrawSwInit C.lv_draw_sw_init
func DrawSwInit()

/**
 * Deinitialize the SW renderers
 */
//go:linkname DrawSwDeinit C.lv_draw_sw_deinit
func DrawSwDeinit()

/**
 * Fill an area using SW render. Handle gradient and radius.
 * @param t             pointer to a draw task
 * @param dsc           the draw descriptor
 * @param coords        the coordinates of the rectangle
 */
// llgo:link (*DrawTaskT).DrawSwFill C.lv_draw_sw_fill
func (recv_ *DrawTaskT) DrawSwFill(dsc *DrawFillDscT, coords *AreaT) {
}

/**
 * Draw border with SW render.
 * @param t             pointer to a draw task
 * @param dsc           the draw descriptor
 * @param coords        the coordinates of the rectangle
 */
// llgo:link (*DrawTaskT).DrawSwBorder C.lv_draw_sw_border
func (recv_ *DrawTaskT) DrawSwBorder(dsc *DrawBorderDscT, coords *AreaT) {
}

/**
 * Draw box shadow with SW render.
 * @param t             pointer to a draw task
 * @param dsc           the draw descriptor
 * @param coords        the coordinates of the rectangle for which the box shadow should be drawn
 */
// llgo:link (*DrawTaskT).DrawSwBoxShadow C.lv_draw_sw_box_shadow
func (recv_ *DrawTaskT) DrawSwBoxShadow(dsc *DrawBoxShadowDscT, coords *AreaT) {
}

/**
 * Draw an image with SW render. It handles image decoding, tiling, transformations, and recoloring.
 * @param t             pointer to a draw task
 * @param draw_dsc      the draw descriptor
 * @param coords        the coordinates of the image
 */
// llgo:link (*DrawTaskT).DrawSwImage C.lv_draw_sw_image
func (recv_ *DrawTaskT) DrawSwImage(draw_dsc *DrawImageDscT, coords *AreaT) {
}

// llgo:link (*DrawTaskT).DrawSwLetter C.lv_draw_sw_letter
func (recv_ *DrawTaskT) DrawSwLetter(dsc *DrawLetterDscT, coords *AreaT) {
}

/**
 * Draw a label with SW render.
 * @param t             pointer to a draw task
 * @param dsc           the draw descriptor
 * @param coords        the coordinates of the label
 */
// llgo:link (*DrawTaskT).DrawSwLabel C.lv_draw_sw_label
func (recv_ *DrawTaskT) DrawSwLabel(dsc *DrawLabelDscT, coords *AreaT) {
}

/**
 * Draw an arc with SW render.
 * @param t             pointer to a draw task
 * @param dsc           the draw descriptor
 * @param coords        the coordinates of the arc
 */
// llgo:link (*DrawTaskT).DrawSwArc C.lv_draw_sw_arc
func (recv_ *DrawTaskT) DrawSwArc(dsc *DrawArcDscT, coords *AreaT) {
}

/**
 * Draw a line with SW render.
 * @param t             pointer to a draw task
 * @param dsc           the draw descriptor
 */
// llgo:link (*DrawTaskT).DrawSwLine C.lv_draw_sw_line
func (recv_ *DrawTaskT) DrawSwLine(dsc *DrawLineDscT) {
}

/**
 * Blend a layer with SW render
 * @param t             pointer to a draw task
 * @param draw_dsc      the draw descriptor
 * @param coords        the coordinates of the layer
 */
// llgo:link (*DrawTaskT).DrawSwLayer C.lv_draw_sw_layer
func (recv_ *DrawTaskT) DrawSwLayer(draw_dsc *DrawImageDscT, coords *AreaT) {
}

/**
 * Draw a triangle with SW render.
 * @param t             pointer to a draw task
 * @param dsc           the draw descriptor
 */
// llgo:link (*DrawTaskT).DrawSwTriangle C.lv_draw_sw_triangle
func (recv_ *DrawTaskT) DrawSwTriangle(dsc *DrawTriangleDscT) {
}

/**
 * Mask out a rectangle with radius from a current layer
 * @param t             pointer to a draw task
 * @param dsc           the draw descriptor
 * @param coords        the coordinates of the mask
 */
// llgo:link (*DrawTaskT).DrawSwMaskRect C.lv_draw_sw_mask_rect
func (recv_ *DrawTaskT) DrawSwMaskRect(dsc *DrawMaskRectDscT) {
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
// llgo:link (*AreaT).DrawSwTransform C.lv_draw_sw_transform
func (recv_ *AreaT) DrawSwTransform(src_buf c.Pointer, src_w c.Int32T, src_h c.Int32T, src_stride c.Int32T, draw_dsc *DrawImageDscT, sup *DrawImageSupT, cf ColorFormatT, dest_buf c.Pointer) {
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
// llgo:link (*DrawSwCustomBlendHandlerT).DrawSwRegisterBlendHandler C.lv_draw_sw_register_blend_handler
func (recv_ *DrawSwCustomBlendHandlerT) DrawSwRegisterBlendHandler() bool {
	return false
}

/**
 * Unregister a custom blend handler for a color format.
 * @param dest_cf color format
 * @return true if a handler was unregistered, false if no handler was registered
 */
// llgo:link ColorFormatT.DrawSwUnregisterBlendHandler C.lv_draw_sw_unregister_blend_handler
func (recv_ ColorFormatT) DrawSwUnregisterBlendHandler() bool {
	return false
}

/**
 * Get the blend handler for a color format.
 * @param dest_cf color format
 * @return pointer to the blend handler or NULL if no handler is registered
 */
// llgo:link ColorFormatT.DrawSwGetBlendHandler C.lv_draw_sw_get_blend_handler
func (recv_ ColorFormatT) DrawSwGetBlendHandler() DrawSwBlendHandlerT {
	return nil
}

type ObjTreeWalkResT c.Int

const (
	OBJ_TREE_WALK_NEXT          ObjTreeWalkResT = 0
	OBJ_TREE_WALK_SKIP_CHILDREN ObjTreeWalkResT = 1
	OBJ_TREE_WALK_END           ObjTreeWalkResT = 2
)

// llgo:type C
type ObjTreeWalkCbT func(*ObjT, c.Pointer) ObjTreeWalkResT

/**
 * Delete an object and all of its children.
 * Also remove the objects from their group and remove all animations (if any).
 * Send `LV_EVENT_DELETED` to deleted objects.
 * @param obj       pointer to an object
 */
// llgo:link (*ObjT).ObjDelete C.lv_obj_delete
func (recv_ *ObjT) ObjDelete() {
}

/**
 * Delete all children of an object.
 * Also remove the objects from their group and remove all animations (if any).
 * Send `LV_EVENT_DELETED` to deleted objects.
 * @param obj       pointer to an object
 */
// llgo:link (*ObjT).ObjClean C.lv_obj_clean
func (recv_ *ObjT) ObjClean() {
}

/**
 * Delete an object after some delay
 * @param obj       pointer to an object
 * @param delay_ms  time to wait before delete in milliseconds
 */
// llgo:link (*ObjT).ObjDeleteDelayed C.lv_obj_delete_delayed
func (recv_ *ObjT) ObjDeleteDelayed(delay_ms c.Uint32T) {
}

/**
 * A function to be easily used in animation ready callback to delete an object when the animation is ready
 * @param a         pointer to the animation
 */
// llgo:link (*AnimT).ObjDeleteAnimCompletedCb C.lv_obj_delete_anim_completed_cb
func (recv_ *AnimT) ObjDeleteAnimCompletedCb() {
}

/**
 * Helper function for asynchronously deleting objects.
 * Useful for cases where you can't delete an object directly in an `LV_EVENT_DELETE` handler (i.e. parent).
 * @param obj       object to delete
 * @see lv_async_call
 */
// llgo:link (*ObjT).ObjDeleteAsync C.lv_obj_delete_async
func (recv_ *ObjT) ObjDeleteAsync() {
}

/**
 * Move the parent of an object. The relative coordinates will be kept.
 *
 * @param obj       pointer to an object whose parent needs to be changed
 * @param parent pointer to the new parent
 */
// llgo:link (*ObjT).ObjSetParent C.lv_obj_set_parent
func (recv_ *ObjT) ObjSetParent(parent *ObjT) {
}

/**
 * Swap the positions of two objects.
 * When used in listboxes, it can be used to sort the listbox items.
 * @param obj1  pointer to the first object
 * @param obj2  pointer to the second object
 */
// llgo:link (*ObjT).ObjSwap C.lv_obj_swap
func (recv_ *ObjT) ObjSwap(obj2 *ObjT) {
}

/**
 * moves the object to the given index in its parent.
 * When used in listboxes, it can be used to sort the listbox items.
 * @param obj  pointer to the object to be moved.
 * @param index  new index in parent. -1 to count from the back
 * @note to move to the background: lv_obj_move_to_index(obj, 0)
 * @note to move forward (up): lv_obj_move_to_index(obj, lv_obj_get_index(obj) - 1)
 */
// llgo:link (*ObjT).ObjMoveToIndex C.lv_obj_move_to_index
func (recv_ *ObjT) ObjMoveToIndex(index c.Int32T) {
}

/**
 * Get the screen of an object
 * @param obj       pointer to an object
 * @return          pointer to the object's screen
 */
// llgo:link (*ObjT).ObjGetScreen C.lv_obj_get_screen
func (recv_ *ObjT) ObjGetScreen() *ObjT {
	return nil
}

/**
 * Get the display of the object
 * @param obj       pointer to an object
 * @return          pointer to the object's display
 */
// llgo:link (*ObjT).ObjGetDisplay C.lv_obj_get_display
func (recv_ *ObjT) ObjGetDisplay() *DisplayT {
	return nil
}

/**
 * Get the parent of an object
 * @param obj       pointer to an object
 * @return          the parent of the object. (NULL if `obj` was a screen)
 */
// llgo:link (*ObjT).ObjGetParent C.lv_obj_get_parent
func (recv_ *ObjT) ObjGetParent() *ObjT {
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
// llgo:link (*ObjT).ObjGetChild C.lv_obj_get_child
func (recv_ *ObjT) ObjGetChild(idx c.Int32T) *ObjT {
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
// llgo:link (*ObjT).ObjGetChildByType C.lv_obj_get_child_by_type
func (recv_ *ObjT) ObjGetChildByType(idx c.Int32T, class_p *ObjClassT) *ObjT {
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
// llgo:link (*ObjT).ObjGetSibling C.lv_obj_get_sibling
func (recv_ *ObjT) ObjGetSibling(idx c.Int32T) *ObjT {
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
// llgo:link (*ObjT).ObjGetSiblingByType C.lv_obj_get_sibling_by_type
func (recv_ *ObjT) ObjGetSiblingByType(idx c.Int32T, class_p *ObjClassT) *ObjT {
	return nil
}

/**
 * Get the number of children
 * @param obj       pointer to an object
 * @return          the number of children
 */
// llgo:link (*ObjT).ObjGetChildCount C.lv_obj_get_child_count
func (recv_ *ObjT) ObjGetChildCount() c.Uint32T {
	return 0
}

/**
 * Get the number of children having a given type.
 * @param obj       pointer to an object
 * @param class_p   the type of the children to check
 * @return          the number of children
 */
// llgo:link (*ObjT).ObjGetChildCountByType C.lv_obj_get_child_count_by_type
func (recv_ *ObjT) ObjGetChildCountByType(class_p *ObjClassT) c.Uint32T {
	return 0
}

/**
 * Get the index of a child.
 * @param obj       pointer to an object
 * @return          the child index of the object.
 *                  E.g. 0: the oldest (firstly created child).
 *                  (-1 if child could not be found or no parent exists)
 */
// llgo:link (*ObjT).ObjGetIndex C.lv_obj_get_index
func (recv_ *ObjT) ObjGetIndex() c.Int32T {
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
// llgo:link (*ObjT).ObjGetIndexByType C.lv_obj_get_index_by_type
func (recv_ *ObjT) ObjGetIndexByType(class_p *ObjClassT) c.Int32T {
	return 0
}

/**
 * Iterate through all children of any object.
 * @param start_obj     start integrating from this object
 * @param cb            call this callback on the objects
 * @param user_data     pointer to any user related data (will be passed to `cb`)
 */
// llgo:link (*ObjT).ObjTreeWalk C.lv_obj_tree_walk
func (recv_ *ObjT) ObjTreeWalk(cb ObjTreeWalkCbT, user_data c.Pointer) {
}

/**
 * Iterate through all children of any object and print their ID.
 * @param start_obj     start integrating from this object
 */
// llgo:link (*ObjT).ObjDumpTree C.lv_obj_dump_tree
func (recv_ *ObjT) ObjDumpTree() {
}

type ObjPointTransformFlagT c.Int

const (
	OBJ_POINT_TRANSFORM_FLAG_NONE              ObjPointTransformFlagT = 0
	OBJ_POINT_TRANSFORM_FLAG_RECURSIVE         ObjPointTransformFlagT = 1
	OBJ_POINT_TRANSFORM_FLAG_INVERSE           ObjPointTransformFlagT = 2
	OBJ_POINT_TRANSFORM_FLAG_INVERSE_RECURSIVE ObjPointTransformFlagT = 3
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
// llgo:link (*ObjT).ObjSetPos C.lv_obj_set_pos
func (recv_ *ObjT) ObjSetPos(x c.Int32T, y c.Int32T) {
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
// llgo:link (*ObjT).ObjSetX C.lv_obj_set_x
func (recv_ *ObjT) ObjSetX(x c.Int32T) {
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
// llgo:link (*ObjT).ObjSetY C.lv_obj_set_y
func (recv_ *ObjT) ObjSetY(y c.Int32T) {
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
// llgo:link (*ObjT).ObjSetSize C.lv_obj_set_size
func (recv_ *ObjT) ObjSetSize(w c.Int32T, h c.Int32T) {
}

/**
 * Recalculate the size of the object
 * @param obj       pointer to an object
 * @return          true: the size has been changed
 */
// llgo:link (*ObjT).ObjRefrSize C.lv_obj_refr_size
func (recv_ *ObjT) ObjRefrSize() bool {
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
// llgo:link (*ObjT).ObjSetWidth C.lv_obj_set_width
func (recv_ *ObjT) ObjSetWidth(w c.Int32T) {
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
// llgo:link (*ObjT).ObjSetHeight C.lv_obj_set_height
func (recv_ *ObjT) ObjSetHeight(h c.Int32T) {
}

/**
 * Set the width reduced by the left and right padding and the border width.
 * @param obj       pointer to an object
 * @param w         the width without paddings in pixels
 */
// llgo:link (*ObjT).ObjSetContentWidth C.lv_obj_set_content_width
func (recv_ *ObjT) ObjSetContentWidth(w c.Int32T) {
}

/**
 * Set the height reduced by the top and bottom padding and the border width.
 * @param obj       pointer to an object
 * @param h         the height without paddings in pixels
 */
// llgo:link (*ObjT).ObjSetContentHeight C.lv_obj_set_content_height
func (recv_ *ObjT) ObjSetContentHeight(h c.Int32T) {
}

/**
 * Set a layout for an object
 * @param obj       pointer to an object
 * @param layout    pointer to a layout descriptor to set
 */
// llgo:link (*ObjT).ObjSetLayout C.lv_obj_set_layout
func (recv_ *ObjT) ObjSetLayout(layout c.Uint32T) {
}

/**
 * Test whether the and object is positioned by a layout or not
 * @param obj       pointer to an object to test
 * @return true:    positioned by a layout; false: not positioned by a layout
 */
// llgo:link (*ObjT).ObjIsLayoutPositioned C.lv_obj_is_layout_positioned
func (recv_ *ObjT) ObjIsLayoutPositioned() bool {
	return false
}

/**
 * Mark the object for layout update.
 * @param obj      pointer to an object whose children need to be updated
 */
// llgo:link (*ObjT).ObjMarkLayoutAsDirty C.lv_obj_mark_layout_as_dirty
func (recv_ *ObjT) ObjMarkLayoutAsDirty() {
}

/**
 * Update the layout of an object.
 * @param obj      pointer to an object whose position and size needs to be updated
 */
// llgo:link (*ObjT).ObjUpdateLayout C.lv_obj_update_layout
func (recv_ *ObjT) ObjUpdateLayout() {
}

/**
 * Change the alignment of an object.
 * @param obj       pointer to an object to align
 * @param align     type of alignment (see 'lv_align_t' enum) `LV_ALIGN_OUT_...` can't be used.
 */
// llgo:link (*ObjT).ObjSetAlign C.lv_obj_set_align
func (recv_ *ObjT) ObjSetAlign(align AlignT) {
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
// llgo:link (*ObjT).ObjAlign C.lv_obj_align
func (recv_ *ObjT) ObjAlign(align AlignT, x_ofs c.Int32T, y_ofs c.Int32T) {
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
// llgo:link (*ObjT).ObjAlignTo C.lv_obj_align_to
func (recv_ *ObjT) ObjAlignTo(base *ObjT, align AlignT, x_ofs c.Int32T, y_ofs c.Int32T) {
}

/**
 * Align an object to the center on its parent.
 * @param obj       pointer to an object to align
 * @note            if the parent size changes `obj` needs to be aligned manually again
 */
// llgo:link (*ObjT).ObjCenter C.lv_obj_center
func (recv_ *ObjT) ObjCenter() {
}

/**
 * Set the transform matrix of an object
 * @param obj       pointer to an object
 * @param matrix    pointer to a matrix to set
 * @note `LV_DRAW_TRANSFORM_USE_MATRIX` needs to be enabled.
 */
// llgo:link (*ObjT).ObjSetTransform C.lv_obj_set_transform
func (recv_ *ObjT) ObjSetTransform(matrix *MatrixT) {
}

/**
 * Reset the transform matrix of an object to identity matrix
 * @param obj       pointer to an object
 * @note `LV_DRAW_TRANSFORM_USE_MATRIX` needs to be enabled.
 */
// llgo:link (*ObjT).ObjResetTransform C.lv_obj_reset_transform
func (recv_ *ObjT) ObjResetTransform() {
}

/**
 * Copy the coordinates of an object to an area
 * @param obj       pointer to an object
 * @param coords    pointer to an area to store the coordinates
 */
// llgo:link (*ObjT).ObjGetCoords C.lv_obj_get_coords
func (recv_ *ObjT) ObjGetCoords(coords *AreaT) {
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
// llgo:link (*ObjT).ObjGetX C.lv_obj_get_x
func (recv_ *ObjT) ObjGetX() c.Int32T {
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
// llgo:link (*ObjT).ObjGetX2 C.lv_obj_get_x2
func (recv_ *ObjT) ObjGetX2() c.Int32T {
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
// llgo:link (*ObjT).ObjGetY C.lv_obj_get_y
func (recv_ *ObjT) ObjGetY() c.Int32T {
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
// llgo:link (*ObjT).ObjGetY2 C.lv_obj_get_y2
func (recv_ *ObjT) ObjGetY2() c.Int32T {
	return 0
}

/**
 * Get the actually set x coordinate of object, i.e. the offset from the set alignment
 * @param obj       pointer to an object
 * @return          the set x coordinate
 */
// llgo:link (*ObjT).ObjGetXAligned C.lv_obj_get_x_aligned
func (recv_ *ObjT) ObjGetXAligned() c.Int32T {
	return 0
}

/**
 * Get the actually set y coordinate of object, i.e. the offset from the set alignment
 * @param obj       pointer to an object
 * @return          the set y coordinate
 */
// llgo:link (*ObjT).ObjGetYAligned C.lv_obj_get_y_aligned
func (recv_ *ObjT) ObjGetYAligned() c.Int32T {
	return 0
}

/**
 * Get the width of an object
 * @param obj       pointer to an object
 * @note            The position of the object is recalculated only on the next redraw. To force coordinate recalculation
 *                  call `lv_obj_update_layout(obj)`.
 * @return          the width in pixels
 */
// llgo:link (*ObjT).ObjGetWidth C.lv_obj_get_width
func (recv_ *ObjT) ObjGetWidth() c.Int32T {
	return 0
}

/**
 * Get the height of an object
 * @param obj       pointer to an object
 * @note            The position of the object is recalculated only on the next redraw. To force coordinate recalculation
 *                  call `lv_obj_update_layout(obj)`.
 * @return          the height in pixels
 */
// llgo:link (*ObjT).ObjGetHeight C.lv_obj_get_height
func (recv_ *ObjT) ObjGetHeight() c.Int32T {
	return 0
}

/**
 * Get the width reduced by the left and right padding and the border width.
 * @param obj       pointer to an object
 * @note            The position of the object is recalculated only on the next redraw. To force coordinate recalculation
 *                  call `lv_obj_update_layout(obj)`.
 * @return          the width which still fits into its parent without causing overflow (making the parent scrollable)
 */
// llgo:link (*ObjT).ObjGetContentWidth C.lv_obj_get_content_width
func (recv_ *ObjT) ObjGetContentWidth() c.Int32T {
	return 0
}

/**
 * Get the height reduced by the top and bottom padding and the border width.
 * @param obj       pointer to an object
 * @note            The position of the object is recalculated only on the next redraw. To force coordinate recalculation
 *                  call `lv_obj_update_layout(obj)`.
 * @return          the height which still fits into the parent without causing overflow (making the parent scrollable)
 */
// llgo:link (*ObjT).ObjGetContentHeight C.lv_obj_get_content_height
func (recv_ *ObjT) ObjGetContentHeight() c.Int32T {
	return 0
}

/**
 * Get the area reduced by the paddings and the border width.
 * @param obj       pointer to an object
 * @note            The position of the object is recalculated only on the next redraw. To force coordinate recalculation
 *                  call `lv_obj_update_layout(obj)`.
 * @param area      the area which still fits into the parent without causing overflow (making the parent scrollable)
 */
// llgo:link (*ObjT).ObjGetContentCoords C.lv_obj_get_content_coords
func (recv_ *ObjT) ObjGetContentCoords(area *AreaT) {
}

/**
 * Get the width occupied by the "parts" of the widget. E.g. the width of all columns of a table.
 * @param obj       pointer to an objects
 * @return          the width of the virtually drawn content
 * @note            This size independent from the real size of the widget.
 *                  It just tells how large the internal ("virtual") content is.
 */
// llgo:link (*ObjT).ObjGetSelfWidth C.lv_obj_get_self_width
func (recv_ *ObjT) ObjGetSelfWidth() c.Int32T {
	return 0
}

/**
 * Get the height occupied by the "parts" of the widget. E.g. the height of all rows of a table.
 * @param obj       pointer to an objects
 * @return          the width of the virtually drawn content
 * @note            This size independent from the real size of the widget.
 *                  It just tells how large the internal ("virtual") content is.
 */
// llgo:link (*ObjT).ObjGetSelfHeight C.lv_obj_get_self_height
func (recv_ *ObjT) ObjGetSelfHeight() c.Int32T {
	return 0
}

/**
 * Handle if the size of the internal ("virtual") content of an object has changed.
 * @param obj       pointer to an object
 * @return          false: nothing happened; true: refresh happened
 */
// llgo:link (*ObjT).ObjRefreshSelfSize C.lv_obj_refresh_self_size
func (recv_ *ObjT) ObjRefreshSelfSize() bool {
	return false
}

// llgo:link (*ObjT).ObjRefrPos C.lv_obj_refr_pos
func (recv_ *ObjT) ObjRefrPos() {
}

// llgo:link (*ObjT).ObjMoveTo C.lv_obj_move_to
func (recv_ *ObjT) ObjMoveTo(x c.Int32T, y c.Int32T) {
}

// llgo:link (*ObjT).ObjMoveChildrenBy C.lv_obj_move_children_by
func (recv_ *ObjT) ObjMoveChildrenBy(x_diff c.Int32T, y_diff c.Int32T, ignore_floating bool) {
}

/**
 * Get the transform matrix of an object
 * @param obj       pointer to an object
 * @return          pointer to the transform matrix or NULL if not set
 */
// llgo:link (*ObjT).ObjGetTransform C.lv_obj_get_transform
func (recv_ *ObjT) ObjGetTransform() *MatrixT {
	return nil
}

/**
 * Transform a point using the angle and zoom style properties of an object
 * @param obj           pointer to an object whose style properties should be used
 * @param p             a point to transform, the result will be written back here too
 * @param flags         OR-ed valued of :cpp:enum:`lv_obj_point_transform_flag_t`
 */
// llgo:link (*ObjT).ObjTransformPoint C.lv_obj_transform_point
func (recv_ *ObjT) ObjTransformPoint(p *PointT, flags ObjPointTransformFlagT) {
}

/**
 * Transform an array of points using the angle and zoom style properties of an object
 * @param obj           pointer to an object whose style properties should be used
 * @param points        the array of points to transform, the result will be written back here too
 * @param count         number of points in the array
 * @param flags         OR-ed valued of :cpp:enum:`lv_obj_point_transform_flag_t`
 */
// llgo:link (*ObjT).ObjTransformPointArray C.lv_obj_transform_point_array
func (recv_ *ObjT) ObjTransformPointArray(points *PointT, count c.SizeT, flags ObjPointTransformFlagT) {
}

/**
 * Transform an area using the angle and zoom style properties of an object
 * @param obj           pointer to an object whose style properties should be used
 * @param area          an area to transform, the result will be written back here too
 * @param flags         OR-ed valued of :cpp:enum:`lv_obj_point_transform_flag_t`
 */
// llgo:link (*ObjT).ObjGetTransformedArea C.lv_obj_get_transformed_area
func (recv_ *ObjT) ObjGetTransformedArea(area *AreaT, flags ObjPointTransformFlagT) {
}

/**
 * Mark an area of an object as invalid.
 * The area will be truncated to the object's area and marked for redraw.
 * @param obj       pointer to an object
 * @param           area the area to redraw
 */
// llgo:link (*ObjT).ObjInvalidateArea C.lv_obj_invalidate_area
func (recv_ *ObjT) ObjInvalidateArea(area *AreaT) {
}

/**
 * Mark the object as invalid to redrawn its area
 * @param obj       pointer to an object
 */
// llgo:link (*ObjT).ObjInvalidate C.lv_obj_invalidate
func (recv_ *ObjT) ObjInvalidate() {
}

/**
 * Tell whether an area of an object is visible (even partially) now or not
 * @param obj       pointer to an object
 * @param area      the are to check. The visible part of the area will be written back here.
 * @return true     visible; false not visible (hidden, out of parent, on other screen, etc)
 */
// llgo:link (*ObjT).ObjAreaIsVisible C.lv_obj_area_is_visible
func (recv_ *ObjT) ObjAreaIsVisible(area *AreaT) bool {
	return false
}

/**
 * Tell whether an object is visible (even partially) now or not
 * @param obj       pointer to an object
 * @return      true: visible; false not visible (hidden, out of parent, on other screen, etc)
 */
// llgo:link (*ObjT).ObjIsVisible C.lv_obj_is_visible
func (recv_ *ObjT) ObjIsVisible() bool {
	return false
}

/**
 * Set the size of an extended clickable area
 * @param obj       pointer to an object
 * @param size      extended clickable area in all 4 directions [px]
 */
// llgo:link (*ObjT).ObjSetExtClickArea C.lv_obj_set_ext_click_area
func (recv_ *ObjT) ObjSetExtClickArea(size c.Int32T) {
}

/**
 * Get the an area where to object can be clicked.
 * It's the object's normal area plus the extended click area.
 * @param obj       pointer to an object
 * @param area      store the result area here
 */
// llgo:link (*ObjT).ObjGetClickArea C.lv_obj_get_click_area
func (recv_ *ObjT) ObjGetClickArea(area *AreaT) {
}

/**
 * Hit-test an object given a particular point in screen space.
 * @param obj       object to hit-test
 * @param point     screen-space point (absolute coordinate)
 * @return          true: if the object is considered under the point
 */
// llgo:link (*ObjT).ObjHitTest C.lv_obj_hit_test
func (recv_ *ObjT) ObjHitTest(point *PointT) bool {
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
//go:linkname ClampWidth C.lv_clamp_width
func ClampWidth(width c.Int32T, min_width c.Int32T, max_width c.Int32T, ref_width c.Int32T) c.Int32T

/**
 * Clamp a height between min and max height. If the min/max height is in percentage value use the ref_height
 * @param height         height to clamp
 * @param min_height     the minimal height
 * @param max_height     the maximal height
 * @param ref_height     the reference height used when min/max height is in percentage
 * @return              the clamped height
 */
//go:linkname ClampHeight C.lv_clamp_height
func ClampHeight(height c.Int32T, min_height c.Int32T, max_height c.Int32T, ref_height c.Int32T) c.Int32T

type ScrollbarModeT c.Int

const (
	SCROLLBAR_MODE_OFF    ScrollbarModeT = 0
	SCROLLBAR_MODE_ON     ScrollbarModeT = 1
	SCROLLBAR_MODE_ACTIVE ScrollbarModeT = 2
	SCROLLBAR_MODE_AUTO   ScrollbarModeT = 3
)

type ScrollSnapT c.Int

const (
	SCROLL_SNAP_NONE   ScrollSnapT = 0
	SCROLL_SNAP_START  ScrollSnapT = 1
	SCROLL_SNAP_END    ScrollSnapT = 2
	SCROLL_SNAP_CENTER ScrollSnapT = 3
)

/**
 * Set how the scrollbars should behave.
 * @param obj       pointer to Widget
 * @param mode      LV_SCROLL_MODE_ON/OFF/AUTO/ACTIVE
 */
// llgo:link (*ObjT).ObjSetScrollbarMode C.lv_obj_set_scrollbar_mode
func (recv_ *ObjT) ObjSetScrollbarMode(mode ScrollbarModeT) {
}

/**
 * Set direction Widget can be scrolled
 * @param obj       pointer to Widget
 * @param dir       one or more bit-wise OR-ed values of `lv_dir_t` enumeration
 */
// llgo:link (*ObjT).ObjSetScrollDir C.lv_obj_set_scroll_dir
func (recv_ *ObjT) ObjSetScrollDir(dir DirT) {
}

/**
 * Set where to snap the children when scrolling ends horizontally
 * @param obj       pointer to Widget
 * @param align     value from `lv_scroll_snap_t` enumeration
 */
// llgo:link (*ObjT).ObjSetScrollSnapX C.lv_obj_set_scroll_snap_x
func (recv_ *ObjT) ObjSetScrollSnapX(align ScrollSnapT) {
}

/**
 * Set where to snap the children when scrolling ends vertically
 * @param obj       pointer to Widget
 * @param align     value from `lv_scroll_snap_t` enumeration
 */
// llgo:link (*ObjT).ObjSetScrollSnapY C.lv_obj_set_scroll_snap_y
func (recv_ *ObjT) ObjSetScrollSnapY(align ScrollSnapT) {
}

/**
 * Get the current scroll mode (when to hide the scrollbars)
 * @param obj       pointer to Widget
 * @return          the current scroll mode from `lv_scrollbar_mode_t`
 */
// llgo:link (*ObjT).ObjGetScrollbarMode C.lv_obj_get_scrollbar_mode
func (recv_ *ObjT) ObjGetScrollbarMode() ScrollbarModeT {
	return 0
}

/**
 * Get directions Widget can be scrolled (set with `lv_obj_set_scroll_dir()`)
 * @param obj       pointer to Widget
 * @return          current scroll direction bit(s)
 */
// llgo:link (*ObjT).ObjGetScrollDir C.lv_obj_get_scroll_dir
func (recv_ *ObjT) ObjGetScrollDir() DirT {
	return 0
}

/**
 * Get where to snap child Widgets when horizontal scrolling ends.
 * @param obj       pointer to Widget
 * @return          current snap value from `lv_scroll_snap_t`
 */
// llgo:link (*ObjT).ObjGetScrollSnapX C.lv_obj_get_scroll_snap_x
func (recv_ *ObjT) ObjGetScrollSnapX() ScrollSnapT {
	return 0
}

/**
 * Get where to snap child Widgets when vertical scrolling ends.
 * @param  obj      pointer to Widget
 * @return          current snap value from `lv_scroll_snap_t`
 */
// llgo:link (*ObjT).ObjGetScrollSnapY C.lv_obj_get_scroll_snap_y
func (recv_ *ObjT) ObjGetScrollSnapY() ScrollSnapT {
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
// llgo:link (*ObjT).ObjGetScrollX C.lv_obj_get_scroll_x
func (recv_ *ObjT) ObjGetScrollX() c.Int32T {
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
// llgo:link (*ObjT).ObjGetScrollY C.lv_obj_get_scroll_y
func (recv_ *ObjT) ObjGetScrollY() c.Int32T {
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
// llgo:link (*ObjT).ObjGetScrollTop C.lv_obj_get_scroll_top
func (recv_ *ObjT) ObjGetScrollTop() c.Int32T {
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
// llgo:link (*ObjT).ObjGetScrollBottom C.lv_obj_get_scroll_bottom
func (recv_ *ObjT) ObjGetScrollBottom() c.Int32T {
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
// llgo:link (*ObjT).ObjGetScrollLeft C.lv_obj_get_scroll_left
func (recv_ *ObjT) ObjGetScrollLeft() c.Int32T {
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
// llgo:link (*ObjT).ObjGetScrollRight C.lv_obj_get_scroll_right
func (recv_ *ObjT) ObjGetScrollRight() c.Int32T {
	return 0
}

/**
 * Get the X and Y coordinates where the scrolling will end for Widget if a scrolling animation is in progress.
 * If no scrolling animation, give the current `x` or `y` scroll position.
 * @param obj       pointer to scrollable Widget
 * @param end       pointer to `lv_point_t` object in which to store result
 */
// llgo:link (*ObjT).ObjGetScrollEnd C.lv_obj_get_scroll_end
func (recv_ *ObjT) ObjGetScrollEnd(end *PointT) {
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
// llgo:link (*ObjT).ObjScrollBy C.lv_obj_scroll_by
func (recv_ *ObjT) ObjScrollBy(dx c.Int32T, dy c.Int32T, anim_en AnimEnableT) {
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
// llgo:link (*ObjT).ObjScrollByBounded C.lv_obj_scroll_by_bounded
func (recv_ *ObjT) ObjScrollByBounded(dx c.Int32T, dy c.Int32T, anim_en AnimEnableT) {
}

/**
 * Scroll to given coordinate on Widget.
 * `x` and `y` will be limited internally to allow scrolling only on the content area.
 * @param obj       pointer to scrollable Widget to scroll
 * @param x         pixels to scroll horizontally
 * @param y         pixels to scroll vertically
 * @param anim_en   LV_ANIM_ON: scroll with animation; LV_ANIM_OFF: scroll immediately
 */
// llgo:link (*ObjT).ObjScrollTo C.lv_obj_scroll_to
func (recv_ *ObjT) ObjScrollTo(x c.Int32T, y c.Int32T, anim_en AnimEnableT) {
}

/**
 * Scroll to X coordinate on Widget.
 * `x` will be limited internally to allow scrolling only on the content area.
 * @param obj       pointer to scrollable Widget to scroll
 * @param x         pixels to scroll horizontally
 * @param anim_en   LV_ANIM_ON: scroll with animation; LV_ANIM_OFF: scroll immediately
 */
// llgo:link (*ObjT).ObjScrollToX C.lv_obj_scroll_to_x
func (recv_ *ObjT) ObjScrollToX(x c.Int32T, anim_en AnimEnableT) {
}

/**
 * Scroll to Y coordinate on Widget.
 * `y` will be limited internally to allow scrolling only on the content area.
 * @param obj       pointer to scrollable Widget to scroll
 * @param y         pixels to scroll vertically
 * @param anim_en   LV_ANIM_ON: scroll with animation; LV_ANIM_OFF: scroll immediately
 */
// llgo:link (*ObjT).ObjScrollToY C.lv_obj_scroll_to_y
func (recv_ *ObjT) ObjScrollToY(y c.Int32T, anim_en AnimEnableT) {
}

/**
 * Scroll `obj`'s parent Widget until `obj` becomes visible.
 * @param obj       pointer to Widget to scroll into view
 * @param anim_en   LV_ANIM_ON: scroll with animation; LV_ANIM_OFF: scroll immediately
 */
// llgo:link (*ObjT).ObjScrollToView C.lv_obj_scroll_to_view
func (recv_ *ObjT) ObjScrollToView(anim_en AnimEnableT) {
}

/**
 * Scroll `obj`'s parent Widgets recursively until `obj` becomes visible.
 * Widget will be scrolled into view even it has nested scrollable parents.
 * @param obj       pointer to Widget to scroll into view
 * @param anim_en   LV_ANIM_ON: scroll with animation; LV_ANIM_OFF: scroll immediately
 */
// llgo:link (*ObjT).ObjScrollToViewRecursive C.lv_obj_scroll_to_view_recursive
func (recv_ *ObjT) ObjScrollToViewRecursive(anim_en AnimEnableT) {
}

/**
 * Tell whether Widget is being scrolled or not at this moment
 * @param obj   pointer to Widget
 * @return      true: `obj` is being scrolled
 */
// llgo:link (*ObjT).ObjIsScrolling C.lv_obj_is_scrolling
func (recv_ *ObjT) ObjIsScrolling() bool {
	return false
}

/**
 * Stop scrolling the current object
 *
 * @param obj The object being scrolled
 */
// llgo:link (*ObjT).ObjStopScrollAnim C.lv_obj_stop_scroll_anim
func (recv_ *ObjT) ObjStopScrollAnim() {
}

/**
 * Check children of `obj` and scroll `obj` to fulfill scroll_snap settings.
 * @param obj       Widget whose children need to be checked and snapped
 * @param anim_en   LV_ANIM_ON/OFF
 */
// llgo:link (*ObjT).ObjUpdateSnap C.lv_obj_update_snap
func (recv_ *ObjT) ObjUpdateSnap(anim_en AnimEnableT) {
}

/**
 * Get the area of the scrollbars
 * @param obj   pointer to Widget
 * @param hor   pointer to store the area of the horizontal scrollbar
 * @param ver   pointer to store the area of the vertical  scrollbar
 */
// llgo:link (*ObjT).ObjGetScrollbarArea C.lv_obj_get_scrollbar_area
func (recv_ *ObjT) ObjGetScrollbarArea(hor *AreaT, ver *AreaT) {
}

/**
 * Invalidate the area of the scrollbars
 * @param obj       pointer to Widget
 */
// llgo:link (*ObjT).ObjScrollbarInvalidate C.lv_obj_scrollbar_invalidate
func (recv_ *ObjT) ObjScrollbarInvalidate() {
}

/**
 * Checks if the content is scrolled "in" and adjusts it to a normal position.
 * @param obj       pointer to Widget
 * @param anim_en   LV_ANIM_ON/OFF
 */
// llgo:link (*ObjT).ObjReadjustScroll C.lv_obj_readjust_scroll
func (recv_ *ObjT) ObjReadjustScroll(anim_en AnimEnableT) {
}

type StyleStateCmpT c.Int

const (
	STYLE_STATE_CMP_SAME          StyleStateCmpT = 0
	STYLE_STATE_CMP_DIFF_REDRAW   StyleStateCmpT = 1
	STYLE_STATE_CMP_DIFF_DRAW_PAD StyleStateCmpT = 2
	STYLE_STATE_CMP_DIFF_LAYOUT   StyleStateCmpT = 3
)

type StyleSelectorT c.Uint32T

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
// llgo:link (*ObjT).ObjAddStyle C.lv_obj_add_style
func (recv_ *ObjT) ObjAddStyle(style *StyleT, selector StyleSelectorT) {
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
// llgo:link (*ObjT).ObjReplaceStyle C.lv_obj_replace_style
func (recv_ *ObjT) ObjReplaceStyle(old_style *StyleT, new_style *StyleT, selector StyleSelectorT) bool {
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
// llgo:link (*ObjT).ObjRemoveStyle C.lv_obj_remove_style
func (recv_ *ObjT) ObjRemoveStyle(style *StyleT, selector StyleSelectorT) {
}

/**
 * Remove all styles from an object
 * @param obj       pointer to an object
 */
// llgo:link (*ObjT).ObjRemoveStyleAll C.lv_obj_remove_style_all
func (recv_ *ObjT) ObjRemoveStyleAll() {
}

/**
 * Notify all object if a style is modified
 * @param style     pointer to a style. Only the objects with this style will be notified
 *                  (NULL to notify all objects)
 */
// llgo:link (*StyleT).ObjReportStyleChange C.lv_obj_report_style_change
func (recv_ *StyleT) ObjReportStyleChange() {
}

/**
 * Notify an object and its children about its style is modified.
 * @param obj       pointer to an object
 * @param part      the part whose style was changed. E.g. `LV_PART_ANY`, `LV_PART_MAIN`
 * @param prop      `LV_STYLE_PROP_ANY` or an `LV_STYLE_...` property.
 *                  It is used to optimize what needs to be refreshed.
 *                  `LV_STYLE_PROP_INV` to perform only a style cache update
 */
// llgo:link (*ObjT).ObjRefreshStyle C.lv_obj_refresh_style
func (recv_ *ObjT) ObjRefreshStyle(part PartT, prop StylePropT) {
}

/**
 * Temporary disable a style for a selector. It will look like is the style wasn't added
 * @param obj       pointer to an object
 * @param style     pointer to a style
 * @param selector  the selector of a style (e.g. LV_STATE_PRESSED | LV_PART_KNOB)
 * @param dis       true: disable the style, false: enable the style
 */
// llgo:link (*ObjT).ObjStyleSetDisabled C.lv_obj_style_set_disabled
func (recv_ *ObjT) ObjStyleSetDisabled(style *StyleT, selector StyleSelectorT, dis bool) {
}

/**
 * Get if a given style is disabled on an object.
 * @param obj       pointer to an object
 * @param style     pointer to a style
 * @param selector  the selector of a style (e.g. LV_STATE_PRESSED | LV_PART_KNOB)
 * @return          true: disable the style, false: enable the style
 */
// llgo:link (*ObjT).ObjStyleGetDisabled C.lv_obj_style_get_disabled
func (recv_ *ObjT) ObjStyleGetDisabled(style *StyleT, selector StyleSelectorT) bool {
	return false
}

/**
 * Enable or disable automatic style refreshing when a new style is added/removed to/from an object
 * or any other style change happens.
 * @param en        true: enable refreshing; false: disable refreshing
 */
//go:linkname ObjEnableStyleRefresh C.lv_obj_enable_style_refresh
func ObjEnableStyleRefresh(en bool)

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
// llgo:link (*ObjT).ObjGetStyleProp C.lv_obj_get_style_prop
func (recv_ *ObjT) ObjGetStyleProp(part PartT, prop StylePropT) StyleValueT {
	return StyleValueT{}
}

/**
 * Check if an object has a specified style property for a given style selector.
 * @param obj       pointer to an object
 * @param selector  the style selector to be checked, defining the scope of the style to be examined.
 * @param prop      the property to be checked.
 * @return          true if the object has the specified selector and property, false otherwise.
 */
// llgo:link (*ObjT).ObjHasStyleProp C.lv_obj_has_style_prop
func (recv_ *ObjT) ObjHasStyleProp(selector StyleSelectorT, prop StylePropT) bool {
	return false
}

/**
 * Set local style property on an object's part and state.
 * @param obj       pointer to an object
 * @param prop      the property
 * @param value     value of the property. The correct element should be set according to the type of the property
 * @param selector  OR-ed value of parts and state for which the style should be set
 */
// llgo:link (*ObjT).ObjSetLocalStyleProp C.lv_obj_set_local_style_prop
func (recv_ *ObjT) ObjSetLocalStyleProp(prop StylePropT, value StyleValueT, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjGetLocalStyleProp C.lv_obj_get_local_style_prop
func (recv_ *ObjT) ObjGetLocalStyleProp(prop StylePropT, value *StyleValueT, selector StyleSelectorT) StyleResT {
	return 0
}

/**
 * Remove a local style property from a part of an object with a given state.
 * @param obj       pointer to an object
 * @param prop      a style property to remove.
 * @param selector  OR-ed value of parts and state for which the style should be removed
 * @return true     the property was found and removed; false: the property was not found
 */
// llgo:link (*ObjT).ObjRemoveLocalStyleProp C.lv_obj_remove_local_style_prop
func (recv_ *ObjT) ObjRemoveLocalStyleProp(prop StylePropT, selector StyleSelectorT) bool {
	return false
}

/**
 * Used internally for color filtering
 */
// llgo:link (*ObjT).ObjStyleApplyColorFilter C.lv_obj_style_apply_color_filter
func (recv_ *ObjT) ObjStyleApplyColorFilter(part PartT, v StyleValueT) StyleValueT {
	return StyleValueT{}
}

/**
 * Fade in an an object and all its children.
 * @param obj       the object to fade in
 * @param time      time of fade
 * @param delay     delay to start the animation
 */
// llgo:link (*ObjT).ObjFadeIn C.lv_obj_fade_in
func (recv_ *ObjT) ObjFadeIn(time c.Uint32T, delay c.Uint32T) {
}

/**
 * Fade out an an object and all its children.
 * @param obj       the object to fade out
 * @param time      time of fade
 * @param delay     delay to start the animation
 */
// llgo:link (*ObjT).ObjFadeOut C.lv_obj_fade_out
func (recv_ *ObjT) ObjFadeOut(time c.Uint32T, delay c.Uint32T) {
}

// llgo:link (*ObjT).ObjSetStyleWidth C.lv_obj_set_style_width
func (recv_ *ObjT) ObjSetStyleWidth(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleMinWidth C.lv_obj_set_style_min_width
func (recv_ *ObjT) ObjSetStyleMinWidth(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleMaxWidth C.lv_obj_set_style_max_width
func (recv_ *ObjT) ObjSetStyleMaxWidth(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleHeight C.lv_obj_set_style_height
func (recv_ *ObjT) ObjSetStyleHeight(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleMinHeight C.lv_obj_set_style_min_height
func (recv_ *ObjT) ObjSetStyleMinHeight(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleMaxHeight C.lv_obj_set_style_max_height
func (recv_ *ObjT) ObjSetStyleMaxHeight(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleLength C.lv_obj_set_style_length
func (recv_ *ObjT) ObjSetStyleLength(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleX C.lv_obj_set_style_x
func (recv_ *ObjT) ObjSetStyleX(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleY C.lv_obj_set_style_y
func (recv_ *ObjT) ObjSetStyleY(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleAlign C.lv_obj_set_style_align
func (recv_ *ObjT) ObjSetStyleAlign(value AlignT, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleTransformWidth C.lv_obj_set_style_transform_width
func (recv_ *ObjT) ObjSetStyleTransformWidth(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleTransformHeight C.lv_obj_set_style_transform_height
func (recv_ *ObjT) ObjSetStyleTransformHeight(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleTranslateX C.lv_obj_set_style_translate_x
func (recv_ *ObjT) ObjSetStyleTranslateX(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleTranslateY C.lv_obj_set_style_translate_y
func (recv_ *ObjT) ObjSetStyleTranslateY(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleTranslateRadial C.lv_obj_set_style_translate_radial
func (recv_ *ObjT) ObjSetStyleTranslateRadial(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleTransformScaleX C.lv_obj_set_style_transform_scale_x
func (recv_ *ObjT) ObjSetStyleTransformScaleX(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleTransformScaleY C.lv_obj_set_style_transform_scale_y
func (recv_ *ObjT) ObjSetStyleTransformScaleY(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleTransformRotation C.lv_obj_set_style_transform_rotation
func (recv_ *ObjT) ObjSetStyleTransformRotation(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleTransformPivotX C.lv_obj_set_style_transform_pivot_x
func (recv_ *ObjT) ObjSetStyleTransformPivotX(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleTransformPivotY C.lv_obj_set_style_transform_pivot_y
func (recv_ *ObjT) ObjSetStyleTransformPivotY(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleTransformSkewX C.lv_obj_set_style_transform_skew_x
func (recv_ *ObjT) ObjSetStyleTransformSkewX(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleTransformSkewY C.lv_obj_set_style_transform_skew_y
func (recv_ *ObjT) ObjSetStyleTransformSkewY(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStylePadTop C.lv_obj_set_style_pad_top
func (recv_ *ObjT) ObjSetStylePadTop(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStylePadBottom C.lv_obj_set_style_pad_bottom
func (recv_ *ObjT) ObjSetStylePadBottom(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStylePadLeft C.lv_obj_set_style_pad_left
func (recv_ *ObjT) ObjSetStylePadLeft(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStylePadRight C.lv_obj_set_style_pad_right
func (recv_ *ObjT) ObjSetStylePadRight(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStylePadRow C.lv_obj_set_style_pad_row
func (recv_ *ObjT) ObjSetStylePadRow(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStylePadColumn C.lv_obj_set_style_pad_column
func (recv_ *ObjT) ObjSetStylePadColumn(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStylePadRadial C.lv_obj_set_style_pad_radial
func (recv_ *ObjT) ObjSetStylePadRadial(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleMarginTop C.lv_obj_set_style_margin_top
func (recv_ *ObjT) ObjSetStyleMarginTop(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleMarginBottom C.lv_obj_set_style_margin_bottom
func (recv_ *ObjT) ObjSetStyleMarginBottom(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleMarginLeft C.lv_obj_set_style_margin_left
func (recv_ *ObjT) ObjSetStyleMarginLeft(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleMarginRight C.lv_obj_set_style_margin_right
func (recv_ *ObjT) ObjSetStyleMarginRight(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleBgColor C.lv_obj_set_style_bg_color
func (recv_ *ObjT) ObjSetStyleBgColor(value ColorT, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleBgOpa C.lv_obj_set_style_bg_opa
func (recv_ *ObjT) ObjSetStyleBgOpa(value OpaT, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleBgGradColor C.lv_obj_set_style_bg_grad_color
func (recv_ *ObjT) ObjSetStyleBgGradColor(value ColorT, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleBgGradDir C.lv_obj_set_style_bg_grad_dir
func (recv_ *ObjT) ObjSetStyleBgGradDir(value GradDirT, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleBgMainStop C.lv_obj_set_style_bg_main_stop
func (recv_ *ObjT) ObjSetStyleBgMainStop(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleBgGradStop C.lv_obj_set_style_bg_grad_stop
func (recv_ *ObjT) ObjSetStyleBgGradStop(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleBgMainOpa C.lv_obj_set_style_bg_main_opa
func (recv_ *ObjT) ObjSetStyleBgMainOpa(value OpaT, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleBgGradOpa C.lv_obj_set_style_bg_grad_opa
func (recv_ *ObjT) ObjSetStyleBgGradOpa(value OpaT, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleBgGrad C.lv_obj_set_style_bg_grad
func (recv_ *ObjT) ObjSetStyleBgGrad(value *GradDscT, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleBgImageSrc C.lv_obj_set_style_bg_image_src
func (recv_ *ObjT) ObjSetStyleBgImageSrc(value c.Pointer, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleBgImageOpa C.lv_obj_set_style_bg_image_opa
func (recv_ *ObjT) ObjSetStyleBgImageOpa(value OpaT, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleBgImageRecolor C.lv_obj_set_style_bg_image_recolor
func (recv_ *ObjT) ObjSetStyleBgImageRecolor(value ColorT, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleBgImageRecolorOpa C.lv_obj_set_style_bg_image_recolor_opa
func (recv_ *ObjT) ObjSetStyleBgImageRecolorOpa(value OpaT, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleBgImageTiled C.lv_obj_set_style_bg_image_tiled
func (recv_ *ObjT) ObjSetStyleBgImageTiled(value bool, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleBorderColor C.lv_obj_set_style_border_color
func (recv_ *ObjT) ObjSetStyleBorderColor(value ColorT, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleBorderOpa C.lv_obj_set_style_border_opa
func (recv_ *ObjT) ObjSetStyleBorderOpa(value OpaT, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleBorderWidth C.lv_obj_set_style_border_width
func (recv_ *ObjT) ObjSetStyleBorderWidth(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleBorderSide C.lv_obj_set_style_border_side
func (recv_ *ObjT) ObjSetStyleBorderSide(value BorderSideT, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleBorderPost C.lv_obj_set_style_border_post
func (recv_ *ObjT) ObjSetStyleBorderPost(value bool, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleOutlineWidth C.lv_obj_set_style_outline_width
func (recv_ *ObjT) ObjSetStyleOutlineWidth(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleOutlineColor C.lv_obj_set_style_outline_color
func (recv_ *ObjT) ObjSetStyleOutlineColor(value ColorT, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleOutlineOpa C.lv_obj_set_style_outline_opa
func (recv_ *ObjT) ObjSetStyleOutlineOpa(value OpaT, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleOutlinePad C.lv_obj_set_style_outline_pad
func (recv_ *ObjT) ObjSetStyleOutlinePad(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleShadowWidth C.lv_obj_set_style_shadow_width
func (recv_ *ObjT) ObjSetStyleShadowWidth(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleShadowOffsetX C.lv_obj_set_style_shadow_offset_x
func (recv_ *ObjT) ObjSetStyleShadowOffsetX(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleShadowOffsetY C.lv_obj_set_style_shadow_offset_y
func (recv_ *ObjT) ObjSetStyleShadowOffsetY(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleShadowSpread C.lv_obj_set_style_shadow_spread
func (recv_ *ObjT) ObjSetStyleShadowSpread(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleShadowColor C.lv_obj_set_style_shadow_color
func (recv_ *ObjT) ObjSetStyleShadowColor(value ColorT, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleShadowOpa C.lv_obj_set_style_shadow_opa
func (recv_ *ObjT) ObjSetStyleShadowOpa(value OpaT, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleImageOpa C.lv_obj_set_style_image_opa
func (recv_ *ObjT) ObjSetStyleImageOpa(value OpaT, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleImageRecolor C.lv_obj_set_style_image_recolor
func (recv_ *ObjT) ObjSetStyleImageRecolor(value ColorT, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleImageRecolorOpa C.lv_obj_set_style_image_recolor_opa
func (recv_ *ObjT) ObjSetStyleImageRecolorOpa(value OpaT, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleLineWidth C.lv_obj_set_style_line_width
func (recv_ *ObjT) ObjSetStyleLineWidth(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleLineDashWidth C.lv_obj_set_style_line_dash_width
func (recv_ *ObjT) ObjSetStyleLineDashWidth(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleLineDashGap C.lv_obj_set_style_line_dash_gap
func (recv_ *ObjT) ObjSetStyleLineDashGap(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleLineRounded C.lv_obj_set_style_line_rounded
func (recv_ *ObjT) ObjSetStyleLineRounded(value bool, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleLineColor C.lv_obj_set_style_line_color
func (recv_ *ObjT) ObjSetStyleLineColor(value ColorT, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleLineOpa C.lv_obj_set_style_line_opa
func (recv_ *ObjT) ObjSetStyleLineOpa(value OpaT, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleArcWidth C.lv_obj_set_style_arc_width
func (recv_ *ObjT) ObjSetStyleArcWidth(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleArcRounded C.lv_obj_set_style_arc_rounded
func (recv_ *ObjT) ObjSetStyleArcRounded(value bool, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleArcColor C.lv_obj_set_style_arc_color
func (recv_ *ObjT) ObjSetStyleArcColor(value ColorT, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleArcOpa C.lv_obj_set_style_arc_opa
func (recv_ *ObjT) ObjSetStyleArcOpa(value OpaT, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleArcImageSrc C.lv_obj_set_style_arc_image_src
func (recv_ *ObjT) ObjSetStyleArcImageSrc(value c.Pointer, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleTextColor C.lv_obj_set_style_text_color
func (recv_ *ObjT) ObjSetStyleTextColor(value ColorT, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleTextOpa C.lv_obj_set_style_text_opa
func (recv_ *ObjT) ObjSetStyleTextOpa(value OpaT, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleTextFont C.lv_obj_set_style_text_font
func (recv_ *ObjT) ObjSetStyleTextFont(value *FontT, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleTextLetterSpace C.lv_obj_set_style_text_letter_space
func (recv_ *ObjT) ObjSetStyleTextLetterSpace(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleTextLineSpace C.lv_obj_set_style_text_line_space
func (recv_ *ObjT) ObjSetStyleTextLineSpace(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleTextDecor C.lv_obj_set_style_text_decor
func (recv_ *ObjT) ObjSetStyleTextDecor(value TextDecorT, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleTextAlign C.lv_obj_set_style_text_align
func (recv_ *ObjT) ObjSetStyleTextAlign(value TextAlignT, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleTextOutlineStrokeColor C.lv_obj_set_style_text_outline_stroke_color
func (recv_ *ObjT) ObjSetStyleTextOutlineStrokeColor(value ColorT, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleTextOutlineStrokeWidth C.lv_obj_set_style_text_outline_stroke_width
func (recv_ *ObjT) ObjSetStyleTextOutlineStrokeWidth(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleTextOutlineStrokeOpa C.lv_obj_set_style_text_outline_stroke_opa
func (recv_ *ObjT) ObjSetStyleTextOutlineStrokeOpa(value OpaT, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleRadius C.lv_obj_set_style_radius
func (recv_ *ObjT) ObjSetStyleRadius(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleRadialOffset C.lv_obj_set_style_radial_offset
func (recv_ *ObjT) ObjSetStyleRadialOffset(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleClipCorner C.lv_obj_set_style_clip_corner
func (recv_ *ObjT) ObjSetStyleClipCorner(value bool, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleOpa C.lv_obj_set_style_opa
func (recv_ *ObjT) ObjSetStyleOpa(value OpaT, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleOpaLayered C.lv_obj_set_style_opa_layered
func (recv_ *ObjT) ObjSetStyleOpaLayered(value OpaT, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleColorFilterDsc C.lv_obj_set_style_color_filter_dsc
func (recv_ *ObjT) ObjSetStyleColorFilterDsc(value *ColorFilterDscT, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleColorFilterOpa C.lv_obj_set_style_color_filter_opa
func (recv_ *ObjT) ObjSetStyleColorFilterOpa(value OpaT, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleRecolor C.lv_obj_set_style_recolor
func (recv_ *ObjT) ObjSetStyleRecolor(value ColorT, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleRecolorOpa C.lv_obj_set_style_recolor_opa
func (recv_ *ObjT) ObjSetStyleRecolorOpa(value OpaT, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleAnim C.lv_obj_set_style_anim
func (recv_ *ObjT) ObjSetStyleAnim(value *AnimT, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleAnimDuration C.lv_obj_set_style_anim_duration
func (recv_ *ObjT) ObjSetStyleAnimDuration(value c.Uint32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleTransition C.lv_obj_set_style_transition
func (recv_ *ObjT) ObjSetStyleTransition(value *StyleTransitionDscT, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleBlendMode C.lv_obj_set_style_blend_mode
func (recv_ *ObjT) ObjSetStyleBlendMode(value BlendModeT, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleLayout C.lv_obj_set_style_layout
func (recv_ *ObjT) ObjSetStyleLayout(value c.Uint16T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleBaseDir C.lv_obj_set_style_base_dir
func (recv_ *ObjT) ObjSetStyleBaseDir(value BaseDirT, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleBitmapMaskSrc C.lv_obj_set_style_bitmap_mask_src
func (recv_ *ObjT) ObjSetStyleBitmapMaskSrc(value c.Pointer, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleRotarySensitivity C.lv_obj_set_style_rotary_sensitivity
func (recv_ *ObjT) ObjSetStyleRotarySensitivity(value c.Uint32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleFlexFlow C.lv_obj_set_style_flex_flow
func (recv_ *ObjT) ObjSetStyleFlexFlow(value FlexFlowT, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleFlexMainPlace C.lv_obj_set_style_flex_main_place
func (recv_ *ObjT) ObjSetStyleFlexMainPlace(value FlexAlignT, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleFlexCrossPlace C.lv_obj_set_style_flex_cross_place
func (recv_ *ObjT) ObjSetStyleFlexCrossPlace(value FlexAlignT, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleFlexTrackPlace C.lv_obj_set_style_flex_track_place
func (recv_ *ObjT) ObjSetStyleFlexTrackPlace(value FlexAlignT, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleFlexGrow C.lv_obj_set_style_flex_grow
func (recv_ *ObjT) ObjSetStyleFlexGrow(value c.Uint8T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleGridColumnDscArray C.lv_obj_set_style_grid_column_dsc_array
func (recv_ *ObjT) ObjSetStyleGridColumnDscArray(value *c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleGridColumnAlign C.lv_obj_set_style_grid_column_align
func (recv_ *ObjT) ObjSetStyleGridColumnAlign(value GridAlignT, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleGridRowDscArray C.lv_obj_set_style_grid_row_dsc_array
func (recv_ *ObjT) ObjSetStyleGridRowDscArray(value *c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleGridRowAlign C.lv_obj_set_style_grid_row_align
func (recv_ *ObjT) ObjSetStyleGridRowAlign(value GridAlignT, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleGridCellColumnPos C.lv_obj_set_style_grid_cell_column_pos
func (recv_ *ObjT) ObjSetStyleGridCellColumnPos(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleGridCellXAlign C.lv_obj_set_style_grid_cell_x_align
func (recv_ *ObjT) ObjSetStyleGridCellXAlign(value GridAlignT, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleGridCellColumnSpan C.lv_obj_set_style_grid_cell_column_span
func (recv_ *ObjT) ObjSetStyleGridCellColumnSpan(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleGridCellRowPos C.lv_obj_set_style_grid_cell_row_pos
func (recv_ *ObjT) ObjSetStyleGridCellRowPos(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleGridCellYAlign C.lv_obj_set_style_grid_cell_y_align
func (recv_ *ObjT) ObjSetStyleGridCellYAlign(value GridAlignT, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjSetStyleGridCellRowSpan C.lv_obj_set_style_grid_cell_row_span
func (recv_ *ObjT) ObjSetStyleGridCellRowSpan(value c.Int32T, selector StyleSelectorT) {
}

// llgo:link (*ObjT).ObjCalculateStyleTextAlign C.lv_obj_calculate_style_text_align
func (recv_ *ObjT) ObjCalculateStyleTextAlign(part PartT, txt *c.Char) TextAlignT {
	return 0
}

/**
 * Get the `opa` style property from all parents and multiply and `>> 8` them.
 * @param obj       the object whose opacity should be get
 * @param part      the part whose opacity should be get. Non-MAIN parts will consider the `opa` of the MAIN part too
 * @return          the final opacity considering the parents' opacity too
 */
// llgo:link (*ObjT).ObjGetStyleOpaRecursive C.lv_obj_get_style_opa_recursive
func (recv_ *ObjT) ObjGetStyleOpaRecursive(part PartT) OpaT {
	return 0
}

/**
 * Apply recolor effect to the input color based on the object's style properties.
 * @param obj       the target object containing recolor style properties
 * @param part      the part to retrieve recolor styles.
 * @param color     the original color to be modified
 * @return          the blended color after applying recolor and opacity
 */
// llgo:link (*ObjT).ObjStyleApplyRecolor C.lv_obj_style_apply_recolor
func (recv_ *ObjT) ObjStyleApplyRecolor(part PartT, color Color32T) Color32T {
	return Color32T{}
}

/**
 * Get the `recolor` style property from all parents and blend them recursively.
 * @param obj       the object whose recolor value should be retrieved
 * @param part      the target part to check. Non-MAIN parts will also consider
 *                  the `recolor` value from the MAIN part during calculation
 * @return          the final blended recolor value combining all parent's recolor values
 */
// llgo:link (*ObjT).ObjGetStyleRecolorRecursive C.lv_obj_get_style_recolor_recursive
func (recv_ *ObjT) ObjGetStyleRecolorRecursive(part PartT) Color32T {
	return Color32T{}
}

type LayerTypeT c.Int

const (
	LAYER_TYPE_NONE      LayerTypeT = 0
	LAYER_TYPE_SIMPLE    LayerTypeT = 1
	LAYER_TYPE_TRANSFORM LayerTypeT = 2
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
// llgo:link (*ObjT).ObjInitDrawRectDsc C.lv_obj_init_draw_rect_dsc
func (recv_ *ObjT) ObjInitDrawRectDsc(part PartT, draw_dsc *DrawRectDscT) {
}

/**
 * Initialize a label draw descriptor from an object's styles in its current state
 * @param obj       pointer to an object
 * @param part      part of the object, e.g. `LV_PART_MAIN`, `LV_PART_SCROLLBAR`, `LV_PART_KNOB`, etc
 * @param draw_dsc  the descriptor to initialize.
 *                  If the `opa` field is set to or the property is equal to `LV_OPA_TRANSP` the rest won't be initialized.
 *                  Should be initialized with `lv_draw_label_dsc_init(draw_dsc)`.
 */
// llgo:link (*ObjT).ObjInitDrawLabelDsc C.lv_obj_init_draw_label_dsc
func (recv_ *ObjT) ObjInitDrawLabelDsc(part PartT, draw_dsc *DrawLabelDscT) {
}

/**
 * Initialize an image draw descriptor from an object's styles in its current state
 * @param obj       pointer to an object
 * @param part      part of the object, e.g. `LV_PART_MAIN`, `LV_PART_SCROLLBAR`, `LV_PART_KNOB`, etc
 * @param draw_dsc  the descriptor to initialize.
 *                  Should be initialized with `lv_draw_image_dsc_init(draw_dsc)`.
 */
// llgo:link (*ObjT).ObjInitDrawImageDsc C.lv_obj_init_draw_image_dsc
func (recv_ *ObjT) ObjInitDrawImageDsc(part PartT, draw_dsc *DrawImageDscT) {
}

/**
 * Initialize a line draw descriptor from an object's styles in its current state
 * @param obj pointer to an object
 * @param part      part of the object, e.g. `LV_PART_MAIN`, `LV_PART_SCROLLBAR`, `LV_PART_KNOB`, etc
 * @param draw_dsc  the descriptor to initialize.
 *                  Should be initialized with `lv_draw_line_dsc_init(draw_dsc)`.
 */
// llgo:link (*ObjT).ObjInitDrawLineDsc C.lv_obj_init_draw_line_dsc
func (recv_ *ObjT) ObjInitDrawLineDsc(part PartT, draw_dsc *DrawLineDscT) {
}

/**
 * Initialize an arc draw descriptor from an object's styles in its current state
 * @param obj       pointer to an object
 * @param part      part of the object, e.g. `LV_PART_MAIN`, `LV_PART_SCROLLBAR`, `LV_PART_KNOB`, etc
 * @param draw_dsc  the descriptor to initialize.
 *                  Should be initialized with `lv_draw_arc_dsc_init(draw_dsc)`.
 */
// llgo:link (*ObjT).ObjInitDrawArcDsc C.lv_obj_init_draw_arc_dsc
func (recv_ *ObjT) ObjInitDrawArcDsc(part PartT, draw_dsc *DrawArcDscT) {
}

/**
 * Get the required extra size (around the object's part) to draw shadow, outline, value etc.
 * @param obj       pointer to an object
 * @param part      part of the object
 * @return          the extra size required around the object
 */
// llgo:link (*ObjT).ObjCalculateExtDrawSize C.lv_obj_calculate_ext_draw_size
func (recv_ *ObjT) ObjCalculateExtDrawSize(part PartT) c.Int32T {
	return 0
}

/**
 * Send a 'LV_EVENT_REFR_EXT_DRAW_SIZE' Call the ancestor's event handler to the object to refresh the value of the extended draw size.
 * The result will be saved in `obj`.
 * @param obj       pointer to an object
 */
// llgo:link (*ObjT).ObjRefreshExtDrawSize C.lv_obj_refresh_ext_draw_size
func (recv_ *ObjT) ObjRefreshExtDrawSize() {
}

type ObjClassEditableT c.Int

const (
	OBJ_CLASS_EDITABLE_INHERIT ObjClassEditableT = 0
	OBJ_CLASS_EDITABLE_TRUE    ObjClassEditableT = 1
	OBJ_CLASS_EDITABLE_FALSE   ObjClassEditableT = 2
)

type ObjClassGroupDefT c.Int

const (
	OBJ_CLASS_GROUP_DEF_INHERIT ObjClassGroupDefT = 0
	OBJ_CLASS_GROUP_DEF_TRUE    ObjClassGroupDefT = 1
	OBJ_CLASS_GROUP_DEF_FALSE   ObjClassGroupDefT = 2
)

type ObjClassThemeInheritableT c.Int

const (
	OBJ_CLASS_THEME_INHERITABLE_FALSE ObjClassThemeInheritableT = 0
	OBJ_CLASS_THEME_INHERITABLE_TRUE  ObjClassThemeInheritableT = 1
)

// llgo:type C
type ObjClassEventCbT func(*ObjClassT, *EventT)

/**
 * Create an object form a class descriptor
 * @param class_p   pointer to a class
 * @param parent    pointer to an object where the new object should be created
 * @return          pointer to the created object
 */
// llgo:link (*ObjClassT).ObjClassCreateObj C.lv_obj_class_create_obj
func (recv_ *ObjClassT) ObjClassCreateObj(parent *ObjT) *ObjT {
	return nil
}

// llgo:link (*ObjT).ObjClassInitObj C.lv_obj_class_init_obj
func (recv_ *ObjT) ObjClassInitObj() {
}

// llgo:link (*ObjT).ObjIsEditable C.lv_obj_is_editable
func (recv_ *ObjT) ObjIsEditable() bool {
	return false
}

// llgo:link (*ObjT).ObjIsGroupDef C.lv_obj_is_group_def
func (recv_ *ObjT) ObjIsGroupDef() bool {
	return false
}

type KeyT c.Int

const (
	KEY_UP        KeyT = 17
	KEY_DOWN      KeyT = 18
	KEY_RIGHT     KeyT = 19
	KEY_LEFT      KeyT = 20
	KEY_ESC       KeyT = 27
	KEY_DEL       KeyT = 127
	KEY_BACKSPACE KeyT = 8
	KEY_ENTER     KeyT = 10
	KEY_NEXT      KeyT = 9
	KEY_PREV      KeyT = 11
	KEY_HOME      KeyT = 2
	KEY_END       KeyT = 3
)

// llgo:type C
type GroupFocusCbT func(*GroupT)

// llgo:type C
type GroupEdgeCbT func(*GroupT, bool)
type GroupRefocusPolicyT c.Int

const (
	GROUP_REFOCUS_POLICY_NEXT GroupRefocusPolicyT = 0
	GROUP_REFOCUS_POLICY_PREV GroupRefocusPolicyT = 1
)

/**
 * Create new Widget group.
 * @return          pointer to the new Widget group
 */
//go:linkname GroupCreate C.lv_group_create
func GroupCreate() *GroupT

/**
 * Delete group object.
 * @param group     pointer to a group
 */
// llgo:link (*GroupT).GroupDelete C.lv_group_delete
func (recv_ *GroupT) GroupDelete() {
}

/**
 * Set default group. New Widgets will be added to this group if it's enabled in
 * their class with `add_to_def_group = true`.
 * @param group     pointer to a group (can be `NULL`)
 */
// llgo:link (*GroupT).GroupSetDefault C.lv_group_set_default
func (recv_ *GroupT) GroupSetDefault() {
}

/**
 * Get default group.
 * @return          pointer to the default group
 */
//go:linkname GroupGetDefault C.lv_group_get_default
func GroupGetDefault() *GroupT

/**
 * Add an Widget to group.
 * @param group     pointer to a group
 * @param obj       pointer to a Widget to add
 */
// llgo:link (*GroupT).GroupAddObj C.lv_group_add_obj
func (recv_ *GroupT) GroupAddObj(obj *ObjT) {
}

/**
 * Swap 2 Widgets in group.  Widgets must be in the same group.
 * @param obj1  pointer to a Widget
 * @param obj2  pointer to another Widget
 */
// llgo:link (*ObjT).GroupSwapObj C.lv_group_swap_obj
func (recv_ *ObjT) GroupSwapObj(obj2 *ObjT) {
}

/**
 * Remove a Widget from its group.
 * @param obj       pointer to Widget to remove
 */
// llgo:link (*ObjT).GroupRemoveObj C.lv_group_remove_obj
func (recv_ *ObjT) GroupRemoveObj() {
}

/**
 * Remove all Widgets from a group.
 * @param group     pointer to a group
 */
// llgo:link (*GroupT).GroupRemoveAllObjs C.lv_group_remove_all_objs
func (recv_ *GroupT) GroupRemoveAllObjs() {
}

/**
 * Focus on a Widget (defocus the current).
 * @param obj       pointer to Widget to focus on
 */
// llgo:link (*ObjT).GroupFocusObj C.lv_group_focus_obj
func (recv_ *ObjT) GroupFocusObj() {
}

/**
 * Focus on next Widget in a group (defocus the current).
 * @param group     pointer to a group
 */
// llgo:link (*GroupT).GroupFocusNext C.lv_group_focus_next
func (recv_ *GroupT) GroupFocusNext() {
}

/**
 * Focus on previous Widget in a group (defocus the current).
 * @param group     pointer to a group
 */
// llgo:link (*GroupT).GroupFocusPrev C.lv_group_focus_prev
func (recv_ *GroupT) GroupFocusPrev() {
}

/**
 * Do not allow changing focus from current Widget.
 * @param group     pointer to a group
 * @param en        true: freeze, false: release freezing (normal mode)
 */
// llgo:link (*GroupT).GroupFocusFreeze C.lv_group_focus_freeze
func (recv_ *GroupT) GroupFocusFreeze(en bool) {
}

/**
 * Send a control character to Widget that has focus in a group.
 * @param group     pointer to a group
 * @param c         a character (use LV_KEY_.. to navigate)
 * @return          result of Widget with focus in group.
 */
// llgo:link (*GroupT).GroupSendData C.lv_group_send_data
func (recv_ *GroupT) GroupSendData(c c.Uint32T) ResultT {
	return 0
}

/**
 * Set a function for a group which will be called when a new Widget has focus.
 * @param group         pointer to a group
 * @param focus_cb      the call back function or NULL if unused
 */
// llgo:link (*GroupT).GroupSetFocusCb C.lv_group_set_focus_cb
func (recv_ *GroupT) GroupSetFocusCb(focus_cb GroupFocusCbT) {
}

/**
 * Set a function for a group which will be called when a focus edge is reached
 * @param group         pointer to a group
 * @param edge_cb      the call back function or NULL if unused
 */
// llgo:link (*GroupT).GroupSetEdgeCb C.lv_group_set_edge_cb
func (recv_ *GroupT) GroupSetEdgeCb(edge_cb GroupEdgeCbT) {
}

/**
 * Set whether the next or previous Widget in a group gets focus when Widget that has
 * focus is deleted.
 * @param group         pointer to a group
 * @param policy        new refocus policy enum
 */
// llgo:link (*GroupT).GroupSetRefocusPolicy C.lv_group_set_refocus_policy
func (recv_ *GroupT) GroupSetRefocusPolicy(policy GroupRefocusPolicyT) {
}

/**
 * Manually set the current mode (edit or navigate).
 * @param group         pointer to group
 * @param edit          true: edit mode; false: navigate mode
 */
// llgo:link (*GroupT).GroupSetEditing C.lv_group_set_editing
func (recv_ *GroupT) GroupSetEditing(edit bool) {
}

/**
 * Set whether moving focus to next/previous Widget will allow wrapping from
 * first->last or last->first Widget.
 * @param group         pointer to group
 * @param               en true: wrapping enabled; false: wrapping disabled
 */
// llgo:link (*GroupT).GroupSetWrap C.lv_group_set_wrap
func (recv_ *GroupT) GroupSetWrap(en bool) {
}

/**
 * Get Widget that has focus, or NULL if there isn't one.
 * @param group         pointer to a group
 * @return              pointer to Widget with focus
 */
// llgo:link (*GroupT).GroupGetFocused C.lv_group_get_focused
func (recv_ *GroupT) GroupGetFocused() *ObjT {
	return nil
}

/**
 * Get focus callback function of a group.
 * @param group pointer to a group
 * @return the call back function or NULL if not set
 */
// llgo:link (*GroupT).GroupGetFocusCb C.lv_group_get_focus_cb
func (recv_ *GroupT) GroupGetFocusCb() GroupFocusCbT {
	return nil
}

/**
 * Get edge callback function of a group.
 * @param group pointer to a group
 * @return the call back function or NULL if not set
 */
// llgo:link (*GroupT).GroupGetEdgeCb C.lv_group_get_edge_cb
func (recv_ *GroupT) GroupGetEdgeCb() GroupEdgeCbT {
	return nil
}

/**
 * Get current mode (edit or navigate).
 * @param group         pointer to group
 * @return              true: edit mode; false: navigate mode
 */
// llgo:link (*GroupT).GroupGetEditing C.lv_group_get_editing
func (recv_ *GroupT) GroupGetEditing() bool {
	return false
}

/**
 * Get whether moving focus to next/previous Widget will allow wrapping from
 * first->last or last->first Widget.
 * @param group         pointer to group
 */
// llgo:link (*GroupT).GroupGetWrap C.lv_group_get_wrap
func (recv_ *GroupT) GroupGetWrap() bool {
	return false
}

/**
 * Get number of Widgets in group.
 * @param group         pointer to a group
 * @return              number of Widgets in the group
 */
// llgo:link (*GroupT).GroupGetObjCount C.lv_group_get_obj_count
func (recv_ *GroupT) GroupGetObjCount() c.Uint32T {
	return 0
}

/**
 * Get nth Widget within group.
 * @param group         pointer to a group
 * @param index         index of Widget within the group
 * @return              pointer to Widget
 */
// llgo:link (*GroupT).GroupGetObjByIndex C.lv_group_get_obj_by_index
func (recv_ *GroupT) GroupGetObjByIndex(index c.Uint32T) *ObjT {
	return nil
}

/**
 * Get the number of groups.
 * @return              number of groups
 */
//go:linkname GroupGetCount C.lv_group_get_count
func GroupGetCount() c.Uint32T

/**
 * Get a group by its index.
 * @param index         index of the group
 * @return              pointer to the group
 */
//go:linkname GroupByIndex C.lv_group_by_index
func GroupByIndex(index c.Uint32T) *GroupT

type IndevTypeT c.Int

const (
	INDEV_TYPE_NONE    IndevTypeT = 0
	INDEV_TYPE_POINTER IndevTypeT = 1
	INDEV_TYPE_KEYPAD  IndevTypeT = 2
	INDEV_TYPE_BUTTON  IndevTypeT = 3
	INDEV_TYPE_ENCODER IndevTypeT = 4
)

type IndevStateT c.Int

const (
	INDEV_STATE_RELEASED IndevStateT = 0
	INDEV_STATE_PRESSED  IndevStateT = 1
)

type IndevModeT c.Int

const (
	INDEV_MODE_NONE  IndevModeT = 0
	INDEV_MODE_TIMER IndevModeT = 1
	INDEV_MODE_EVENT IndevModeT = 2
)

type IndevGestureTypeT c.Int

const (
	INDEV_GESTURE_NONE              IndevGestureTypeT = 0
	INDEV_GESTURE_PINCH             IndevGestureTypeT = 1
	INDEV_GESTURE_SWIPE             IndevGestureTypeT = 2
	INDEV_GESTURE_ROTATE            IndevGestureTypeT = 3
	INDEV_GESTURE_TWO_FINGERS_SWIPE IndevGestureTypeT = 4
	INDEV_GESTURE_SCROLL            IndevGestureTypeT = 5
	INDEV_GESTURE_CNT               IndevGestureTypeT = 6
)

/** Data structure passed to an input driver to fill*/

type IndevDataT struct {
	GestureType     [6]IndevGestureTypeT
	GestureData     [6]c.Pointer
	State           IndevStateT
	Point           PointT
	Key             c.Uint32T
	BtnId           c.Uint32T
	EncDiff         c.Int16T
	Timestamp       c.Uint32T
	ContinueReading bool
}

// llgo:type C
type IndevReadCbT func(*IndevT, *IndevDataT)

/**
 * Create an indev
 * @return Pointer to the created indev or NULL when allocation failed
 */
//go:linkname IndevCreate C.lv_indev_create
func IndevCreate() *IndevT

/**
 * Remove the provided input device. Make sure not to use the provided input device afterwards anymore.
 * @param indev pointer to delete
 */
// llgo:link (*IndevT).IndevDelete C.lv_indev_delete
func (recv_ *IndevT) IndevDelete() {
}

/**
 * Get the next input device.
 * @param indev pointer to the current input device. NULL to initialize.
 * @return the next input device or NULL if there are no more. Provide the first input device when
 * the parameter is NULL
 */
// llgo:link (*IndevT).IndevGetNext C.lv_indev_get_next
func (recv_ *IndevT) IndevGetNext() *IndevT {
	return nil
}

/**
 * Read data from an input device.
 * @param indev pointer to an input device
 */
// llgo:link (*IndevT).IndevRead C.lv_indev_read
func (recv_ *IndevT) IndevRead() {
}

/**
 * Called periodically to read the input devices
 * @param timer pointer to a timer to read
 */
// llgo:link (*TimerT).IndevReadTimerCb C.lv_indev_read_timer_cb
func (recv_ *TimerT) IndevReadTimerCb() {
}

/**
 * Enable or disable one or all input devices (default enabled)
 * @param indev pointer to an input device or NULL to enable/disable all of them
 * @param enable true to enable, false to disable
 */
// llgo:link (*IndevT).IndevEnable C.lv_indev_enable
func (recv_ *IndevT) IndevEnable(enable bool) {
}

/**
 * Get the currently processed input device. Can be used in action functions too.
 * @return pointer to the currently processed input device or NULL if no input device processing
 * right now
 */
//go:linkname IndevActive C.lv_indev_active
func IndevActive() *IndevT

/**
 * Set the type of an input device
 * @param indev pointer to an input device
 * @param indev_type the type of the input device from `lv_indev_type_t` (`LV_INDEV_TYPE_...`)
 */
// llgo:link (*IndevT).IndevSetType C.lv_indev_set_type
func (recv_ *IndevT) IndevSetType(indev_type IndevTypeT) {
}

/**
 * Set a callback function to read input device data to the indev
 * @param indev pointer to an input device
 * @param read_cb pointer to callback function to read input device data
 */
// llgo:link (*IndevT).IndevSetReadCb C.lv_indev_set_read_cb
func (recv_ *IndevT) IndevSetReadCb(read_cb IndevReadCbT) {
}

/**
 * Set user data to the indev
 * @param indev pointer to an input device
 * @param user_data pointer to user data
 */
// llgo:link (*IndevT).IndevSetUserData C.lv_indev_set_user_data
func (recv_ *IndevT) IndevSetUserData(user_data c.Pointer) {
}

/**
 * Set driver data to the indev
 * @param indev pointer to an input device
 * @param driver_data pointer to driver data
 */
// llgo:link (*IndevT).IndevSetDriverData C.lv_indev_set_driver_data
func (recv_ *IndevT) IndevSetDriverData(driver_data c.Pointer) {
}

/**
 * Assign a display to the indev
 * @param indev pointer to an input device
 * @param disp pointer to an display
 */
// llgo:link (*IndevT).IndevSetDisplay C.lv_indev_set_display
func (recv_ *IndevT) IndevSetDisplay(disp *X_lvDisplayT) {
}

/**
 * Set long press time to indev
 * @param  indev            pointer to input device
 * @param  long_press_time  time long press time in ms
 */
// llgo:link (*IndevT).IndevSetLongPressTime C.lv_indev_set_long_press_time
func (recv_ *IndevT) IndevSetLongPressTime(long_press_time c.Uint16T) {
}

/**
 * Set long press repeat time to indev
 * @param  indev            pointer to input device
 * @param  long_press_repeat_time  long press repeat time in ms
 */
// llgo:link (*IndevT).IndevSetLongPressRepeatTime C.lv_indev_set_long_press_repeat_time
func (recv_ *IndevT) IndevSetLongPressRepeatTime(long_press_repeat_time c.Uint16T) {
}

/**
 * Set scroll limit to the input device
 * @param indev pointer to an input device
 * @param scroll_limit the number of pixels to slide before actually drag the object
 */
// llgo:link (*IndevT).IndevSetScrollLimit C.lv_indev_set_scroll_limit
func (recv_ *IndevT) IndevSetScrollLimit(scroll_limit c.Uint8T) {
}

/**
 * Set scroll throw slow-down to the indev. Greater value means faster slow-down
 * @param indev pointer to an input device
 * @param scroll_throw the slow-down in [%]
 */
// llgo:link (*IndevT).IndevSetScrollThrow C.lv_indev_set_scroll_throw
func (recv_ *IndevT) IndevSetScrollThrow(scroll_throw c.Uint8T) {
}

/**
 * Get the type of an input device
 * @param indev pointer to an input device
 * @return the type of the input device from `lv_hal_indev_type_t` (`LV_INDEV_TYPE_...`)
 */
// llgo:link (*IndevT).IndevGetType C.lv_indev_get_type
func (recv_ *IndevT) IndevGetType() IndevTypeT {
	return 0
}

/**
 * Get the callback function to read input device data to the indev
 * @param indev pointer to an input device
 * @return Pointer to callback function to read input device data or NULL if indev is NULL
 */
// llgo:link (*IndevT).IndevGetReadCb C.lv_indev_get_read_cb
func (recv_ *IndevT) IndevGetReadCb() IndevReadCbT {
	return nil
}

/**
 * Get the indev state
 * @param indev pointer to an input device
 * @return Indev state or LV_INDEV_STATE_RELEASED if indev is NULL
 */
// llgo:link (*IndevT).IndevGetState C.lv_indev_get_state
func (recv_ *IndevT) IndevGetState() IndevStateT {
	return 0
}

/**
 * Get the indev assigned group
 * @param indev pointer to an input device
 * @return Pointer to indev assigned group or NULL if indev is NULL
 */
// llgo:link (*IndevT).IndevGetGroup C.lv_indev_get_group
func (recv_ *IndevT) IndevGetGroup() *GroupT {
	return nil
}

/**
 * Get a pointer to the assigned display of the indev
 * @param indev pointer to an input device
 * @return pointer to the assigned display or NULL if indev is NULL
 */
// llgo:link (*IndevT).IndevGetDisplay C.lv_indev_get_display
func (recv_ *IndevT) IndevGetDisplay() *DisplayT {
	return nil
}

/**
 * Get a pointer to the user data of the indev
 * @param indev pointer to an input device
 * @return pointer to the user data or NULL if indev is NULL
 */
// llgo:link (*IndevT).IndevGetUserData C.lv_indev_get_user_data
func (recv_ *IndevT) IndevGetUserData() c.Pointer {
	return nil
}

/**
 * Get a pointer to the driver data of the indev
 * @param indev pointer to an input device
 * @return pointer to the driver data or NULL if indev is NULL
 */
// llgo:link (*IndevT).IndevGetDriverData C.lv_indev_get_driver_data
func (recv_ *IndevT) IndevGetDriverData() c.Pointer {
	return nil
}

/**
 * Get whether indev is moved while pressed
 * @param indev pointer to an input device
 * @return true: indev is moved while pressed; false: indev is not moved while pressed
 */
// llgo:link (*IndevT).IndevGetPressMoved C.lv_indev_get_press_moved
func (recv_ *IndevT) IndevGetPressMoved() bool {
	return false
}

/**
 * Reset one or all input devices
 * @param indev pointer to an input device to reset or NULL to reset all of them
 * @param obj pointer to an object which triggers the reset.
 */
// llgo:link (*IndevT).IndevReset C.lv_indev_reset
func (recv_ *IndevT) IndevReset(obj *ObjT) {
}

/**
 * Touch and key related events are sent to the input device first and to the widget after that.
 * If this functions called in an indev event, the event won't be sent to the widget.
 * @param indev pointer to an input device
 */
// llgo:link (*IndevT).IndevStopProcessing C.lv_indev_stop_processing
func (recv_ *IndevT) IndevStopProcessing() {
}

/**
 * Reset the long press state of an input device
 * @param indev pointer to an input device
 */
// llgo:link (*IndevT).IndevResetLongPress C.lv_indev_reset_long_press
func (recv_ *IndevT) IndevResetLongPress() {
}

/**
 * Set a cursor for a pointer input device (for LV_INPUT_TYPE_POINTER and LV_INPUT_TYPE_BUTTON)
 * @param indev pointer to an input device
 * @param cur_obj pointer to an object to be used as cursor
 */
// llgo:link (*IndevT).IndevSetCursor C.lv_indev_set_cursor
func (recv_ *IndevT) IndevSetCursor(cur_obj *ObjT) {
}

/**
 * Set a destination group for a keypad input device (for LV_INDEV_TYPE_KEYPAD)
 * @param indev pointer to an input device
 * @param group pointer to a group
 */
// llgo:link (*IndevT).IndevSetGroup C.lv_indev_set_group
func (recv_ *IndevT) IndevSetGroup(group *GroupT) {
}

/**
 * Set the an array of points for LV_INDEV_TYPE_BUTTON.
 * These points will be assigned to the buttons to press a specific point on the screen
 * @param indev pointer to an input device
 * @param points array of points
 */
// llgo:link (*IndevT).IndevSetButtonPoints C.lv_indev_set_button_points
func (recv_ *IndevT) IndevSetButtonPoints(points *PointT) {
}

/**
 * Get the last point of an input device (for LV_INDEV_TYPE_POINTER and LV_INDEV_TYPE_BUTTON)
 * @param indev pointer to an input device
 * @param point pointer to a point to store the result
 */
// llgo:link (*IndevT).IndevGetPoint C.lv_indev_get_point
func (recv_ *IndevT) IndevGetPoint(point *PointT) {
}

/**
* Get the current gesture direct
* @param indev pointer to an input device
* @return current gesture direct
 */
// llgo:link (*IndevT).IndevGetGestureDir C.lv_indev_get_gesture_dir
func (recv_ *IndevT) IndevGetGestureDir() DirT {
	return 0
}

/**
 * Get the last pressed key of an input device (for LV_INDEV_TYPE_KEYPAD)
 * @param indev pointer to an input device
 * @return the last pressed key (0 on error)
 */
// llgo:link (*IndevT).IndevGetKey C.lv_indev_get_key
func (recv_ *IndevT) IndevGetKey() c.Uint32T {
	return 0
}

/**
 * Get the counter for consecutive clicks within a short distance and time.
 * The counter is updated before LV_EVENT_SHORT_CLICKED is fired.
 * @param indev pointer to an input device
 * @return short click streak counter
 */
// llgo:link (*IndevT).IndevGetShortClickStreak C.lv_indev_get_short_click_streak
func (recv_ *IndevT) IndevGetShortClickStreak() c.Uint8T {
	return 0
}

/**
 * Check the current scroll direction of an input device (for LV_INDEV_TYPE_POINTER and
 * LV_INDEV_TYPE_BUTTON)
 * @param indev pointer to an input device
 * @return LV_DIR_NONE: no scrolling now
 *         LV_DIR_HOR/VER
 */
// llgo:link (*IndevT).IndevGetScrollDir C.lv_indev_get_scroll_dir
func (recv_ *IndevT) IndevGetScrollDir() DirT {
	return 0
}

/**
 * Get the currently scrolled object (for LV_INDEV_TYPE_POINTER and
 * LV_INDEV_TYPE_BUTTON)
 * @param indev pointer to an input device
 * @return pointer to the currently scrolled object or NULL if no scrolling by this indev
 */
// llgo:link (*IndevT).IndevGetScrollObj C.lv_indev_get_scroll_obj
func (recv_ *IndevT) IndevGetScrollObj() *ObjT {
	return nil
}

/**
 * Get the movement vector of an input device (for LV_INDEV_TYPE_POINTER and
 * LV_INDEV_TYPE_BUTTON)
 * @param indev pointer to an input device
 * @param point pointer to a point to store the types.pointer.vector
 */
// llgo:link (*IndevT).IndevGetVect C.lv_indev_get_vect
func (recv_ *IndevT) IndevGetVect(point *PointT) {
}

/**
 * Get the cursor object of an input device (for LV_INDEV_TYPE_POINTER only)
 * @param indev pointer to an input device
 * @return pointer to the cursor object
 */
// llgo:link (*IndevT).IndevGetCursor C.lv_indev_get_cursor
func (recv_ *IndevT) IndevGetCursor() *ObjT {
	return nil
}

/**
 * Do nothing until the next release
 * @param indev pointer to an input device
 */
// llgo:link (*IndevT).IndevWaitRelease C.lv_indev_wait_release
func (recv_ *IndevT) IndevWaitRelease() {
}

/**
 * Gets a pointer to the currently active object in the currently processed input device.
 * @return pointer to currently active object or NULL if no active object
 */
//go:linkname IndevGetActiveObj C.lv_indev_get_active_obj
func IndevGetActiveObj() *ObjT

/**
 * Get a pointer to the indev read timer to
 * modify its parameters with `lv_timer_...` functions.
 * @param indev pointer to an input device
 * @return pointer to the indev read refresher timer. (NULL on error)
 */
// llgo:link (*IndevT).IndevGetReadTimer C.lv_indev_get_read_timer
func (recv_ *IndevT) IndevGetReadTimer() *TimerT {
	return nil
}

/**
* Set the input device's event model: event-driven mode or timer mode.
* @param indev pointer to an input device
* @param mode the mode of input device
 */
// llgo:link (*IndevT).IndevSetMode C.lv_indev_set_mode
func (recv_ *IndevT) IndevSetMode(mode IndevModeT) {
}

/**
 * Get the input device's running mode.
 * @param indev pointer to an input device
 * @return the running mode for the specified input device.
 */
// llgo:link (*IndevT).IndevGetMode C.lv_indev_get_mode
func (recv_ *IndevT) IndevGetMode() IndevModeT {
	return 0
}

/**
 * Search the most top, clickable object by a point
 * @param obj pointer to a start object, typically the screen
 * @param point pointer to a point for searching the most top child
 * @return pointer to the found object or NULL if there was no suitable object
 */
// llgo:link (*ObjT).IndevSearchObj C.lv_indev_search_obj
func (recv_ *ObjT) IndevSearchObj(point *PointT) *ObjT {
	return nil
}

/**
 * Add an event handler to the indev
 * @param indev          pointer to an indev
 * @param event_cb      an event callback
 * @param filter        event code to react or `LV_EVENT_ALL`
 * @param user_data     optional user_data
 */
// llgo:link (*IndevT).IndevAddEventCb C.lv_indev_add_event_cb
func (recv_ *IndevT) IndevAddEventCb(event_cb EventCbT, filter EventCodeT, user_data c.Pointer) {
}

/**
 * Get the number of event attached to an indev
 * @param indev          pointer to an indev
 * @return              number of events
 */
// llgo:link (*IndevT).IndevGetEventCount C.lv_indev_get_event_count
func (recv_ *IndevT) IndevGetEventCount() c.Uint32T {
	return 0
}

/**
 * Get an event descriptor for an event
 * @param indev          pointer to an indev
 * @param index         the index of the event
 * @return              the event descriptor
 */
// llgo:link (*IndevT).IndevGetEventDsc C.lv_indev_get_event_dsc
func (recv_ *IndevT) IndevGetEventDsc(index c.Uint32T) *EventDscT {
	return nil
}

/**
 * Remove an event
 * @param indev         pointer to an indev
 * @param index         the index of the event to remove
 * @return              true: and event was removed; false: no event was removed
 */
// llgo:link (*IndevT).IndevRemoveEvent C.lv_indev_remove_event
func (recv_ *IndevT) IndevRemoveEvent(index c.Uint32T) bool {
	return false
}

/**
 * Remove an event_cb with user_data
 * @param indev         pointer to a indev
 * @param event_cb      the event_cb of the event to remove
 * @param user_data     user_data
 * @return              the count of the event removed
 */
// llgo:link (*IndevT).IndevRemoveEventCbWithUserData C.lv_indev_remove_event_cb_with_user_data
func (recv_ *IndevT) IndevRemoveEventCbWithUserData(event_cb EventCbT, user_data c.Pointer) c.Uint32T {
	return 0
}

/**
 * Send an event to an indev
 * @param indev         pointer to an indev
 * @param code          an event code. LV_EVENT_...
 * @param param         optional param
 * @return              LV_RESULT_OK: indev wasn't deleted in the event.
 */
// llgo:link (*IndevT).IndevSendEvent C.lv_indev_send_event
func (recv_ *IndevT) IndevSendEvent(code EventCodeT, param c.Pointer) ResultT {
	return 0
}

type CoverResT c.Int

const (
	COVER_RES_COVER     CoverResT = 0
	COVER_RES_NOT_COVER CoverResT = 1
	COVER_RES_MASKED    CoverResT = 2
)

/**
 * Send an event to the object
 * @param obj           pointer to an object
 * @param event_code    the type of the event from `lv_event_t`
 * @param param         arbitrary data depending on the widget type and the event. (Usually `NULL`)
 * @return LV_RESULT_OK: `obj` was not deleted in the event; LV_RESULT_INVALID: `obj` was deleted in the event_code
 */
// llgo:link (*ObjT).ObjSendEvent C.lv_obj_send_event
func (recv_ *ObjT) ObjSendEvent(event_code EventCodeT, param c.Pointer) ResultT {
	return 0
}

/**
 * Used by the widgets internally to call the ancestor widget types's event handler
 * @param class_p   pointer to the class of the widget (NOT the ancestor class)
 * @param e         pointer to the event descriptor
 * @return          LV_RESULT_OK: the target object was not deleted in the event; LV_RESULT_INVALID: it was deleted in the event_code
 */
// llgo:link (*ObjClassT).ObjEventBase C.lv_obj_event_base
func (recv_ *ObjClassT) ObjEventBase(e *EventT) ResultT {
	return 0
}

/**
 * Get the current target of the event. It's the object which event handler being called.
 * If the event is not bubbled it's the same as "original" target.
 * @param e     pointer to the event descriptor
 * @return      the target of the event_code
 */
// llgo:link (*EventT).EventGetCurrentTargetObj C.lv_event_get_current_target_obj
func (recv_ *EventT) EventGetCurrentTargetObj() *ObjT {
	return nil
}

/**
 * Get the object originally targeted by the event. It's the same even if the event is bubbled.
 * @param e     pointer to the event descriptor
 * @return      pointer to the original target of the event_code
 */
// llgo:link (*EventT).EventGetTargetObj C.lv_event_get_target_obj
func (recv_ *EventT) EventGetTargetObj() *ObjT {
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
// llgo:link (*ObjT).ObjAddEventCb C.lv_obj_add_event_cb
func (recv_ *ObjT) ObjAddEventCb(event_cb EventCbT, filter EventCodeT, user_data c.Pointer) *EventDscT {
	return nil
}

// llgo:link (*ObjT).ObjGetEventCount C.lv_obj_get_event_count
func (recv_ *ObjT) ObjGetEventCount() c.Uint32T {
	return 0
}

// llgo:link (*ObjT).ObjGetEventDsc C.lv_obj_get_event_dsc
func (recv_ *ObjT) ObjGetEventDsc(index c.Uint32T) *EventDscT {
	return nil
}

// llgo:link (*ObjT).ObjRemoveEvent C.lv_obj_remove_event
func (recv_ *ObjT) ObjRemoveEvent(index c.Uint32T) bool {
	return false
}

// llgo:link (*ObjT).ObjRemoveEventDsc C.lv_obj_remove_event_dsc
func (recv_ *ObjT) ObjRemoveEventDsc(dsc *EventDscT) bool {
	return false
}

/**
 * Remove an event_cb from an object
 * @param obj           pointer to a obj
 * @param event_cb      the event_cb of the event to remove
 * @return              the count of the event removed
 */
// llgo:link (*ObjT).ObjRemoveEventCb C.lv_obj_remove_event_cb
func (recv_ *ObjT) ObjRemoveEventCb(event_cb EventCbT) c.Uint32T {
	return 0
}

/**
 * Remove an event_cb with user_data
 * @param obj           pointer to a obj
 * @param event_cb      the event_cb of the event to remove
 * @param user_data     user_data
 * @return              the count of the event removed
 */
// llgo:link (*ObjT).ObjRemoveEventCbWithUserData C.lv_obj_remove_event_cb_with_user_data
func (recv_ *ObjT) ObjRemoveEventCbWithUserData(event_cb EventCbT, user_data c.Pointer) c.Uint32T {
	return 0
}

/**
 * Get the input device passed as parameter to indev related events.
 * @param e     pointer to an event
 * @return      the indev that triggered the event or NULL if called on a not indev related event
 */
// llgo:link (*EventT).EventGetIndev C.lv_event_get_indev
func (recv_ *EventT) EventGetIndev() *IndevT {
	return nil
}

/**
 * Get the draw context which should be the first parameter of the draw functions.
 * Namely: `LV_EVENT_DRAW_MAIN/POST`, `LV_EVENT_DRAW_MAIN/POST_BEGIN`, `LV_EVENT_DRAW_MAIN/POST_END`
 * @param e     pointer to an event
 * @return      pointer to a draw context or NULL if called on an unrelated event
 */
// llgo:link (*EventT).EventGetLayer C.lv_event_get_layer
func (recv_ *EventT) EventGetLayer() *LayerT {
	return nil
}

/**
 * Get the old area of the object before its size was changed. Can be used in `LV_EVENT_SIZE_CHANGED`
 * @param e     pointer to an event
 * @return      the old absolute area of the object or NULL if called on an unrelated event
 */
// llgo:link (*EventT).EventGetOldSize C.lv_event_get_old_size
func (recv_ *EventT) EventGetOldSize() *AreaT {
	return nil
}

/**
 * Get the key passed as parameter to an event. Can be used in `LV_EVENT_KEY`
 * @param e     pointer to an event
 * @return      the triggering key or NULL if called on an unrelated event
 */
// llgo:link (*EventT).EventGetKey C.lv_event_get_key
func (recv_ *EventT) EventGetKey() c.Uint32T {
	return 0
}

/**
 * Get the signed rotary encoder diff. passed as parameter to an event. Can be used in `LV_EVENT_ROTARY`
 * @param e     pointer to an event
 * @return      the triggering key or NULL if called on an unrelated event
 */
// llgo:link (*EventT).EventGetRotaryDiff C.lv_event_get_rotary_diff
func (recv_ *EventT) EventGetRotaryDiff() c.Int32T {
	return 0
}

/**
 * Get the animation descriptor of a scrolling. Can be used in `LV_EVENT_SCROLL_BEGIN`
 * @param e     pointer to an event
 * @return      the animation that will scroll the object. (can be modified as required)
 */
// llgo:link (*EventT).EventGetScrollAnim C.lv_event_get_scroll_anim
func (recv_ *EventT) EventGetScrollAnim() *AnimT {
	return nil
}

/**
 * Set the new extra draw size. Can be used in `LV_EVENT_REFR_EXT_DRAW_SIZE`
 * @param e     pointer to an event
 * @param size  The new extra draw size
 */
// llgo:link (*EventT).EventSetExtDrawSize C.lv_event_set_ext_draw_size
func (recv_ *EventT) EventSetExtDrawSize(size c.Int32T) {
}

/**
 * Get a pointer to an `lv_point_t` variable in which the self size should be saved (width in `point->x` and height `point->y`).
 * Can be used in `LV_EVENT_GET_SELF_SIZE`
 * @param e     pointer to an event
 * @return      pointer to `lv_point_t` or NULL if called on an unrelated event
 */
// llgo:link (*EventT).EventGetSelfSizeInfo C.lv_event_get_self_size_info
func (recv_ *EventT) EventGetSelfSizeInfo() *PointT {
	return nil
}

/**
 * Get a pointer to an `lv_hit_test_info_t` variable in which the hit test result should be saved. Can be used in `LV_EVENT_HIT_TEST`
 * @param e     pointer to an event
 * @return      pointer to `lv_hit_test_info_t` or NULL if called on an unrelated event
 */
// llgo:link (*EventT).EventGetHitTestInfo C.lv_event_get_hit_test_info
func (recv_ *EventT) EventGetHitTestInfo() *HitTestInfoT {
	return nil
}

/**
 * Get a pointer to an area which should be examined whether the object fully covers it or not.
 * Can be used in `LV_EVENT_HIT_TEST`
 * @param e     pointer to an event
 * @return      an area with absolute coordinates to check
 */
// llgo:link (*EventT).EventGetCoverArea C.lv_event_get_cover_area
func (recv_ *EventT) EventGetCoverArea() *AreaT {
	return nil
}

/**
 * Set the result of cover checking. Can be used in `LV_EVENT_COVER_CHECK`
 * @param e     pointer to an event
 * @param res   an element of ::lv_cover_check_info_t
 */
// llgo:link (*EventT).EventSetCoverRes C.lv_event_set_cover_res
func (recv_ *EventT) EventSetCoverRes(res CoverResT) {
}

/**
 * Get the draw task which was just added.
 * Can be used in `LV_EVENT_DRAW_TASK_ADDED event`
 * @param e     pointer to an event
 * @return      the added draw task
 */
// llgo:link (*EventT).EventGetDrawTask C.lv_event_get_draw_task
func (recv_ *EventT) EventGetDrawTask() *DrawTaskT {
	return nil
}

type X_lvStateT c.Int

const (
	STATE_DEFAULT   X_lvStateT = 0
	STATE_CHECKED   X_lvStateT = 1
	STATE_FOCUSED   X_lvStateT = 2
	STATE_FOCUS_KEY X_lvStateT = 4
	STATE_EDITED    X_lvStateT = 8
	STATE_HOVERED   X_lvStateT = 16
	STATE_PRESSED   X_lvStateT = 32
	STATE_SCROLLED  X_lvStateT = 64
	STATE_DISABLED  X_lvStateT = 128
	STATE_USER_1    X_lvStateT = 4096
	STATE_USER_2    X_lvStateT = 8192
	STATE_USER_3    X_lvStateT = 16384
	STATE_USER_4    X_lvStateT = 32768
	STATE_ANY       X_lvStateT = 65535
)

type X_lvPartT c.Int

const (
	PART_MAIN         X_lvPartT = 0
	PART_SCROLLBAR    X_lvPartT = 65536
	PART_INDICATOR    X_lvPartT = 131072
	PART_KNOB         X_lvPartT = 196608
	PART_SELECTED     X_lvPartT = 262144
	PART_ITEMS        X_lvPartT = 327680
	PART_CURSOR       X_lvPartT = 393216
	PART_CUSTOM_FIRST X_lvPartT = 524288
	PART_ANY          X_lvPartT = 983040
)

type ObjFlagT c.Int

const (
	OBJ_FLAG_HIDDEN                ObjFlagT = 1
	OBJ_FLAG_CLICKABLE             ObjFlagT = 2
	OBJ_FLAG_CLICK_FOCUSABLE       ObjFlagT = 4
	OBJ_FLAG_CHECKABLE             ObjFlagT = 8
	OBJ_FLAG_SCROLLABLE            ObjFlagT = 16
	OBJ_FLAG_SCROLL_ELASTIC        ObjFlagT = 32
	OBJ_FLAG_SCROLL_MOMENTUM       ObjFlagT = 64
	OBJ_FLAG_SCROLL_ONE            ObjFlagT = 128
	OBJ_FLAG_SCROLL_CHAIN_HOR      ObjFlagT = 256
	OBJ_FLAG_SCROLL_CHAIN_VER      ObjFlagT = 512
	OBJ_FLAG_SCROLL_CHAIN          ObjFlagT = 768
	OBJ_FLAG_SCROLL_ON_FOCUS       ObjFlagT = 1024
	OBJ_FLAG_SCROLL_WITH_ARROW     ObjFlagT = 2048
	OBJ_FLAG_SNAPPABLE             ObjFlagT = 4096
	OBJ_FLAG_PRESS_LOCK            ObjFlagT = 8192
	OBJ_FLAG_EVENT_BUBBLE          ObjFlagT = 16384
	OBJ_FLAG_GESTURE_BUBBLE        ObjFlagT = 32768
	OBJ_FLAG_ADV_HITTEST           ObjFlagT = 65536
	OBJ_FLAG_IGNORE_LAYOUT         ObjFlagT = 131072
	OBJ_FLAG_FLOATING              ObjFlagT = 262144
	OBJ_FLAG_SEND_DRAW_TASK_EVENTS ObjFlagT = 524288
	OBJ_FLAG_OVERFLOW_VISIBLE      ObjFlagT = 1048576
	OBJ_FLAG_FLEX_IN_NEW_TRACK     ObjFlagT = 2097152
	OBJ_FLAG_LAYOUT_1              ObjFlagT = 8388608
	OBJ_FLAG_LAYOUT_2              ObjFlagT = 16777216
	OBJ_FLAG_WIDGET_1              ObjFlagT = 33554432
	OBJ_FLAG_WIDGET_2              ObjFlagT = 67108864
	OBJ_FLAG_USER_1                ObjFlagT = 134217728
	OBJ_FLAG_USER_2                ObjFlagT = 268435456
	OBJ_FLAG_USER_3                ObjFlagT = 536870912
	OBJ_FLAG_USER_4                ObjFlagT = 1073741824
)

/**
 * Create a base object (a rectangle)
 * @param parent    pointer to a parent object. If NULL then a screen will be created.
 * @return          pointer to the new object
 */
// llgo:link (*ObjT).ObjCreate C.lv_obj_create
func (recv_ *ObjT) ObjCreate() *ObjT {
	return nil
}

/**
 * Set one or more flags
 * @param obj   pointer to an object
 * @param f     OR-ed values from `lv_obj_flag_t` to set.
 */
// llgo:link (*ObjT).ObjAddFlag C.lv_obj_add_flag
func (recv_ *ObjT) ObjAddFlag(f ObjFlagT) {
}

/**
 * Remove one or more flags
 * @param obj   pointer to an object
 * @param f     OR-ed values from `lv_obj_flag_t` to clear.
 */
// llgo:link (*ObjT).ObjRemoveFlag C.lv_obj_remove_flag
func (recv_ *ObjT) ObjRemoveFlag(f ObjFlagT) {
}

/**
 * Set add or remove one or more flags.
 * @param obj   pointer to an object
 * @param f     OR-ed values from `lv_obj_flag_t` to update.
 * @param v     true: add the flags; false: remove the flags
 */
// llgo:link (*ObjT).ObjSetFlag C.lv_obj_set_flag
func (recv_ *ObjT) ObjSetFlag(f ObjFlagT, v bool) {
}

/**
 * Add one or more states to the object. The other state bits will remain unchanged.
 * If specified in the styles, transition animation will be started from the previous state to the current.
 * @param obj       pointer to an object
 * @param state     the states to add. E.g `LV_STATE_PRESSED | LV_STATE_FOCUSED`
 */
// llgo:link (*ObjT).ObjAddState C.lv_obj_add_state
func (recv_ *ObjT) ObjAddState(state StateT) {
}

/**
 * Remove one or more states to the object. The other state bits will remain unchanged.
 * If specified in the styles, transition animation will be started from the previous state to the current.
 * @param obj       pointer to an object
 * @param state     the states to add. E.g `LV_STATE_PRESSED | LV_STATE_FOCUSED`
 */
// llgo:link (*ObjT).ObjRemoveState C.lv_obj_remove_state
func (recv_ *ObjT) ObjRemoveState(state StateT) {
}

/**
 * Add or remove one or more states to the object. The other state bits will remain unchanged.
 * @param obj       pointer to an object
 * @param state     the states to add. E.g `LV_STATE_PRESSED | LV_STATE_FOCUSED`
 * @param v         true: add the states; false: remove the states
 */
// llgo:link (*ObjT).ObjSetState C.lv_obj_set_state
func (recv_ *ObjT) ObjSetState(state StateT, v bool) {
}

/**
 * Set the user_data field of the object
 * @param obj   pointer to an object
 * @param user_data   pointer to the new user_data.
 */
// llgo:link (*ObjT).ObjSetUserData C.lv_obj_set_user_data
func (recv_ *ObjT) ObjSetUserData(user_data c.Pointer) {
}

/**
 * Check if a given flag or all the given flags are set on an object.
 * @param obj   pointer to an object
 * @param f     the flag(s) to check (OR-ed values can be used)
 * @return      true: all flags are set; false: not all flags are set
 */
// llgo:link (*ObjT).ObjHasFlag C.lv_obj_has_flag
func (recv_ *ObjT) ObjHasFlag(f ObjFlagT) bool {
	return false
}

/**
 * Check if a given flag or any of the flags are set on an object.
 * @param obj   pointer to an object
 * @param f     the flag(s) to check (OR-ed values can be used)
 * @return      true: at least one flag is set; false: none of the flags are set
 */
// llgo:link (*ObjT).ObjHasFlagAny C.lv_obj_has_flag_any
func (recv_ *ObjT) ObjHasFlagAny(f ObjFlagT) bool {
	return false
}

/**
 * Get the state of an object
 * @param obj   pointer to an object
 * @return      the state (OR-ed values from `lv_state_t`)
 */
// llgo:link (*ObjT).ObjGetState C.lv_obj_get_state
func (recv_ *ObjT) ObjGetState() StateT {
	return 0
}

/**
 * Check if the object is in a given state or not.
 * @param obj       pointer to an object
 * @param state     a state or combination of states to check
 * @return          true: `obj` is in `state`; false: `obj` is not in `state`
 */
// llgo:link (*ObjT).ObjHasState C.lv_obj_has_state
func (recv_ *ObjT) ObjHasState(state StateT) bool {
	return false
}

/**
 * Get the group of the object
 * @param       obj pointer to an object
 * @return      the pointer to group of the object
 */
// llgo:link (*ObjT).ObjGetGroup C.lv_obj_get_group
func (recv_ *ObjT) ObjGetGroup() *GroupT {
	return nil
}

/**
 * Get the user_data field of the object
 * @param obj   pointer to an object
 * @return      the pointer to the user_data of the object
 */
// llgo:link (*ObjT).ObjGetUserData C.lv_obj_get_user_data
func (recv_ *ObjT) ObjGetUserData() c.Pointer {
	return nil
}

/**
 * Allocate special data for an object if not allocated yet.
 * @param obj   pointer to an object
 */
// llgo:link (*ObjT).ObjAllocateSpecAttr C.lv_obj_allocate_spec_attr
func (recv_ *ObjT) ObjAllocateSpecAttr() {
}

/**
 * Check the type of obj.
 * @param obj       pointer to an object
 * @param class_p   a class to check (e.g. `lv_slider_class`)
 * @return          true: `class_p` is the `obj` class.
 */
// llgo:link (*ObjT).ObjCheckType C.lv_obj_check_type
func (recv_ *ObjT) ObjCheckType(class_p *ObjClassT) bool {
	return false
}

/**
 * Check if any object has a given class (type).
 * It checks the ancestor classes too.
 * @param obj       pointer to an object
 * @param class_p   a class to check (e.g. `lv_slider_class`)
 * @return          true: `obj` has the given class
 */
// llgo:link (*ObjT).ObjHasClass C.lv_obj_has_class
func (recv_ *ObjT) ObjHasClass(class_p *ObjClassT) bool {
	return false
}

/**
 * Get the class (type) of the object
 * @param obj   pointer to an object
 * @return      the class (type) of the object
 */
// llgo:link (*ObjT).ObjGetClass C.lv_obj_get_class
func (recv_ *ObjT) ObjGetClass() *ObjClassT {
	return nil
}

/**
 * Check if any object is still "alive".
 * @param obj       pointer to an object
 * @return          true: valid
 */
// llgo:link (*ObjT).ObjIsValid C.lv_obj_is_valid
func (recv_ *ObjT) ObjIsValid() bool {
	return false
}

/**
 * Utility to set an object reference to NULL when it gets deleted.
 * The reference should be in a location that will not become invalid
 * during the object's lifetime, i.e. static or allocated.
 * @param obj_ptr   a pointer to a pointer to an object
 */
//go:linkname ObjNullOnDelete C.lv_obj_null_on_delete
func ObjNullOnDelete(obj_ptr **ObjT)

/**
 * Add an event handler to a widget that will load a screen on a trigger.
 * @param obj           pointer to widget which should load the screen
 * @param trigger       an event code, e.g. `LV_EVENT_CLICKED`
 * @param screen        the screen to load (must be a valid widget)
 * @param anim_type     element of `lv_screen_load_anim_t` the screen load animation
 * @param duration      duration of the animation in milliseconds
 * @param delay         delay before the screen load in milliseconds
 */
// llgo:link (*ObjT).ObjAddScreenLoadEvent C.lv_obj_add_screen_load_event
func (recv_ *ObjT) ObjAddScreenLoadEvent(trigger EventCodeT, screen *ObjT, anim_type ScreenLoadAnimT, duration c.Uint32T, delay c.Uint32T) {
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
// llgo:link (*ObjT).ObjAddScreenCreateEvent C.lv_obj_add_screen_create_event
func (recv_ *ObjT) ObjAddScreenCreateEvent(trigger EventCodeT, screen_create_cb ScreenCreateCbT, anim_type ScreenLoadAnimT, duration c.Uint32T, delay c.Uint32T) {
}

type SubjectTypeT c.Int

const (
	SUBJECT_TYPE_INVALID SubjectTypeT = 0
	SUBJECT_TYPE_NONE    SubjectTypeT = 1
	SUBJECT_TYPE_INT     SubjectTypeT = 2
	SUBJECT_TYPE_FLOAT   SubjectTypeT = 3
	SUBJECT_TYPE_POINTER SubjectTypeT = 4
	SUBJECT_TYPE_COLOR   SubjectTypeT = 5
	SUBJECT_TYPE_GROUP   SubjectTypeT = 6
	SUBJECT_TYPE_STRING  SubjectTypeT = 7
)

/**
 * A common type to handle all the various observable types in the same way
 */

type SubjectValueT struct {
	Pointer c.Pointer
}

/**
 * The Subject (an observable value)
 */

type SubjectT struct {
	SubsLl             LlT
	Value              SubjectValueT
	PrevValue          SubjectValueT
	UserData           c.Pointer
	Type               c.Uint32T
	Size               c.Uint32T
	NotifyRestartQuery c.Uint32T
}

// llgo:type C
type ObserverCbT func(*ObserverT, *SubjectT)

/**
 * Initialize an integer-type Subject.
 * @param subject   pointer to Subject
 * @param value     initial value
 */
// llgo:link (*SubjectT).SubjectInitInt C.lv_subject_init_int
func (recv_ *SubjectT) SubjectInitInt(value c.Int32T) {
}

/**
 * Set value of an integer Subject and notify Observers.
 * @param subject   pointer to Subject
 * @param value     new value
 */
// llgo:link (*SubjectT).SubjectSetInt C.lv_subject_set_int
func (recv_ *SubjectT) SubjectSetInt(value c.Int32T) {
}

/**
 * Get current value of an integer Subject.
 * @param subject   pointer to Subject
 * @return          current value
 */
// llgo:link (*SubjectT).SubjectGetInt C.lv_subject_get_int
func (recv_ *SubjectT) SubjectGetInt() c.Int32T {
	return 0
}

/**
 * Get previous value of an integer Subject.
 * @param subject   pointer to Subject
 * @return          current value
 */
// llgo:link (*SubjectT).SubjectGetPreviousInt C.lv_subject_get_previous_int
func (recv_ *SubjectT) SubjectGetPreviousInt() c.Int32T {
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
// llgo:link (*SubjectT).SubjectInitString C.lv_subject_init_string
func (recv_ *SubjectT) SubjectInitString(buf *c.Char, prev_buf *c.Char, size c.SizeT, value *c.Char) {
}

/**
 * Copy a string to a Subject and notify Observers if it changed.
 * @param subject   pointer to Subject
 * @param buf       new string
 */
// llgo:link (*SubjectT).SubjectCopyString C.lv_subject_copy_string
func (recv_ *SubjectT) SubjectCopyString(buf *c.Char) {
}

/**
 * Format a new string, updating Subject, and notify Observers if it changed.
 * @param subject   pointer to Subject
 * @param format    format string
 */
// llgo:link (*SubjectT).SubjectSnprintf C.lv_subject_snprintf
func (recv_ *SubjectT) SubjectSnprintf(format *c.Char, __llgo_va_list ...interface{}) {
}

/**
 * Get current value of a string Subject.
 * @param subject   pointer to Subject
 * @return          pointer to buffer containing current value
 */
// llgo:link (*SubjectT).SubjectGetString C.lv_subject_get_string
func (recv_ *SubjectT) SubjectGetString() *c.Char {
	return nil
}

/**
 * Get previous value of a string Subject.
 * @param subject   pointer to Subject
 * @return          pointer to buffer containing previous value
 * @note            NULL will be returned if NULL was passed in `lv_subject_init_string()`
 *                  as `prev_buf`.
 */
// llgo:link (*SubjectT).SubjectGetPreviousString C.lv_subject_get_previous_string
func (recv_ *SubjectT) SubjectGetPreviousString() *c.Char {
	return nil
}

/**
 * Initialize a pointer-type Subject.
 * @param subject   pointer to Subject
 * @param value     initial value
 */
// llgo:link (*SubjectT).SubjectInitPointer C.lv_subject_init_pointer
func (recv_ *SubjectT) SubjectInitPointer(value c.Pointer) {
}

/**
 * Set value of a pointer Subject and notify Observers (regardless of whether it changed).
 * @param subject   pointer to Subject
 * @param ptr       new value
 */
// llgo:link (*SubjectT).SubjectSetPointer C.lv_subject_set_pointer
func (recv_ *SubjectT) SubjectSetPointer(ptr c.Pointer) {
}

/**
 * Get current value of a pointer Subject.
 * @param subject   pointer to Subject
 * @return          current value
 */
// llgo:link (*SubjectT).SubjectGetPointer C.lv_subject_get_pointer
func (recv_ *SubjectT) SubjectGetPointer() c.Pointer {
	return nil
}

/**
 * Get previous value of a pointer Subject.
 * @param subject   pointer to Subject
 * @return          previous value
 */
// llgo:link (*SubjectT).SubjectGetPreviousPointer C.lv_subject_get_previous_pointer
func (recv_ *SubjectT) SubjectGetPreviousPointer() c.Pointer {
	return nil
}

/**
 * Initialize a color-type Subject.
 * @param subject   pointer to Subject
 * @param color     initial value
 */
// llgo:link (*SubjectT).SubjectInitColor C.lv_subject_init_color
func (recv_ *SubjectT) SubjectInitColor(color ColorT) {
}

/**
 * Set value of a color Subject and notify Observers if it changed.
 * @param subject   pointer to Subject
 * @param color     new value
 */
// llgo:link (*SubjectT).SubjectSetColor C.lv_subject_set_color
func (recv_ *SubjectT) SubjectSetColor(color ColorT) {
}

/**
 * Get current value of a color Subject.
 * @param subject   pointer to Subject
 * @return          current value
 */
// llgo:link (*SubjectT).SubjectGetColor C.lv_subject_get_color
func (recv_ *SubjectT) SubjectGetColor() ColorT {
	return ColorT{}
}

/**
 * Get previous value of a color Subject.
 * @param subject   pointer to Subject
 * @return          previous value
 */
// llgo:link (*SubjectT).SubjectGetPreviousColor C.lv_subject_get_previous_color
func (recv_ *SubjectT) SubjectGetPreviousColor() ColorT {
	return ColorT{}
}

/**
* Initialize a Group-type Subject.
* @param group_subject  pointer to Group-type Subject
* @param list           list of other Subject addresses; when any of these have values
                            updated, Observers of `group_subject` will be notified.
* @param list_len       number of elements in `list[]`
*/
// llgo:link (*SubjectT).SubjectInitGroup C.lv_subject_init_group
func (recv_ *SubjectT) SubjectInitGroup(list **SubjectT, list_len c.Uint32T) {
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
// llgo:link (*SubjectT).SubjectDeinit C.lv_subject_deinit
func (recv_ *SubjectT) SubjectDeinit() {
}

/**
 * Get an element from Subject Group's list.
 * @param subject   pointer to Group-type Subject
 * @param index     index of element to get
 * @return          pointer to indexed Subject from list, or NULL if index is out of bounds
 */
// llgo:link (*SubjectT).SubjectGetGroupElement C.lv_subject_get_group_element
func (recv_ *SubjectT) SubjectGetGroupElement(index c.Int32T) *SubjectT {
	return nil
}

/**
 * Add Observer to Subject. When Subject's value changes `observer_cb` will be called.
 * @param subject       pointer to Subject
 * @param observer_cb   notification callback
 * @param user_data     optional user data
 * @return              pointer to newly-created Observer
 */
// llgo:link (*SubjectT).SubjectAddObserver C.lv_subject_add_observer
func (recv_ *SubjectT) SubjectAddObserver(observer_cb ObserverCbT, user_data c.Pointer) *ObserverT {
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
// llgo:link (*SubjectT).SubjectAddObserverObj C.lv_subject_add_observer_obj
func (recv_ *SubjectT) SubjectAddObserverObj(observer_cb ObserverCbT, obj *ObjT, user_data c.Pointer) *ObserverT {
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
// llgo:link (*SubjectT).SubjectAddObserverWithTarget C.lv_subject_add_observer_with_target
func (recv_ *SubjectT) SubjectAddObserverWithTarget(observer_cb ObserverCbT, target c.Pointer, user_data c.Pointer) *ObserverT {
	return nil
}

/**
 * Remove Observer from its Subject.
 * @param observer      pointer to Observer
 */
// llgo:link (*ObserverT).ObserverRemove C.lv_observer_remove
func (recv_ *ObserverT) ObserverRemove() {
}

/**
 * Remove Observers associated with Widget `obj` from specified `subject` or all Subjects.
 * @param obj       pointer to Widget whose Observers should be removed
 * @param subject   Subject to remove Widget from, or NULL to remove from all Subjects
 * @note            This function can be used e.g. when a Widget's Subject(s) needs to
 *                  be replaced by other Subject(s)
 */
// llgo:link (*ObjT).ObjRemoveFromSubject C.lv_obj_remove_from_subject
func (recv_ *ObjT) ObjRemoveFromSubject(subject *SubjectT) {
}

/**
 * Get target of an Observer.
 * @param observer      pointer to Observer
 * @return              pointer to saved target
 */
// llgo:link (*ObserverT).ObserverGetTarget C.lv_observer_get_target
func (recv_ *ObserverT) ObserverGetTarget() c.Pointer {
	return nil
}

/**
 * Get target Widget of Observer.
 * This is the same as `lv_observer_get_target()`, except it returns `target`
 * as an `lv_obj_t *`.
 * @param observer      pointer to Observer
 * @return              pointer to saved Widget target
 */
// llgo:link (*ObserverT).ObserverGetTargetObj C.lv_observer_get_target_obj
func (recv_ *ObserverT) ObserverGetTargetObj() *ObjT {
	return nil
}

/**
 * Get Observer's user data.
 * @param observer      pointer to Observer
 * @return              void pointer to saved user data
 */
// llgo:link (*ObserverT).ObserverGetUserData C.lv_observer_get_user_data
func (recv_ *ObserverT) ObserverGetUserData() c.Pointer {
	return nil
}

/**
 * Notify all Observers of Subject.
 * @param subject       pointer to Subject
 */
// llgo:link (*SubjectT).SubjectNotify C.lv_subject_notify
func (recv_ *SubjectT) SubjectNotify() {
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
// llgo:link (*ObjT).ObjAddSubjectIncrementEvent C.lv_obj_add_subject_increment_event
func (recv_ *ObjT) ObjAddSubjectIncrementEvent(subject *SubjectT, trigger EventCodeT, step c.Int32T, min c.Int32T, max c.Int32T) {
}

/**
 * Set the value of an integer subject.
 * @param obj       pointer to a widget
 * @param subject   pointer to a subject to change
 * @param trigger   the trigger on which the subject should be changed
 * @param value     the value to set
 */
// llgo:link (*ObjT).ObjAddSubjectSetIntEvent C.lv_obj_add_subject_set_int_event
func (recv_ *ObjT) ObjAddSubjectSetIntEvent(subject *SubjectT, trigger EventCodeT, value c.Int32T) {
}

/**
 * Set the value of a string subject.
 * @param obj       pointer to a widget
 * @param subject   pointer to a subject to change
 * @param trigger   the trigger on which the subject should be changed
 * @param value     the value to set
 */
// llgo:link (*ObjT).ObjAddSubjectSetStringEvent C.lv_obj_add_subject_set_string_event
func (recv_ *ObjT) ObjAddSubjectSetStringEvent(subject *SubjectT, trigger EventCodeT, value *c.Char) {
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
// llgo:link (*ObjT).ObjBindStyle C.lv_obj_bind_style
func (recv_ *ObjT) ObjBindStyle(style *StyleT, selector StyleSelectorT, subject *SubjectT, ref_value c.Int32T) *ObserverT {
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
// llgo:link (*ObjT).ObjBindFlagIfEq C.lv_obj_bind_flag_if_eq
func (recv_ *ObjT) ObjBindFlagIfEq(subject *SubjectT, flag ObjFlagT, ref_value c.Int32T) *ObserverT {
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
// llgo:link (*ObjT).ObjBindFlagIfNotEq C.lv_obj_bind_flag_if_not_eq
func (recv_ *ObjT) ObjBindFlagIfNotEq(subject *SubjectT, flag ObjFlagT, ref_value c.Int32T) *ObserverT {
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
// llgo:link (*ObjT).ObjBindFlagIfGt C.lv_obj_bind_flag_if_gt
func (recv_ *ObjT) ObjBindFlagIfGt(subject *SubjectT, flag ObjFlagT, ref_value c.Int32T) *ObserverT {
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
// llgo:link (*ObjT).ObjBindFlagIfGe C.lv_obj_bind_flag_if_ge
func (recv_ *ObjT) ObjBindFlagIfGe(subject *SubjectT, flag ObjFlagT, ref_value c.Int32T) *ObserverT {
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
// llgo:link (*ObjT).ObjBindFlagIfLt C.lv_obj_bind_flag_if_lt
func (recv_ *ObjT) ObjBindFlagIfLt(subject *SubjectT, flag ObjFlagT, ref_value c.Int32T) *ObserverT {
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
// llgo:link (*ObjT).ObjBindFlagIfLe C.lv_obj_bind_flag_if_le
func (recv_ *ObjT) ObjBindFlagIfLe(subject *SubjectT, flag ObjFlagT, ref_value c.Int32T) *ObserverT {
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
// llgo:link (*ObjT).ObjBindStateIfEq C.lv_obj_bind_state_if_eq
func (recv_ *ObjT) ObjBindStateIfEq(subject *SubjectT, state StateT, ref_value c.Int32T) *ObserverT {
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
// llgo:link (*ObjT).ObjBindStateIfNotEq C.lv_obj_bind_state_if_not_eq
func (recv_ *ObjT) ObjBindStateIfNotEq(subject *SubjectT, state StateT, ref_value c.Int32T) *ObserverT {
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
// llgo:link (*ObjT).ObjBindStateIfGt C.lv_obj_bind_state_if_gt
func (recv_ *ObjT) ObjBindStateIfGt(subject *SubjectT, state StateT, ref_value c.Int32T) *ObserverT {
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
// llgo:link (*ObjT).ObjBindStateIfGe C.lv_obj_bind_state_if_ge
func (recv_ *ObjT) ObjBindStateIfGe(subject *SubjectT, state StateT, ref_value c.Int32T) *ObserverT {
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
// llgo:link (*ObjT).ObjBindStateIfLt C.lv_obj_bind_state_if_lt
func (recv_ *ObjT) ObjBindStateIfLt(subject *SubjectT, state StateT, ref_value c.Int32T) *ObserverT {
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
// llgo:link (*ObjT).ObjBindStateIfLe C.lv_obj_bind_state_if_le
func (recv_ *ObjT) ObjBindStateIfLe(subject *SubjectT, state StateT, ref_value c.Int32T) *ObserverT {
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
// llgo:link (*ObjT).ObjBindChecked C.lv_obj_bind_checked
func (recv_ *ObjT) ObjBindChecked(subject *SubjectT) *ObserverT {
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
// llgo:link (*ObjT).LabelBindText C.lv_label_bind_text
func (recv_ *ObjT) LabelBindText(subject *SubjectT, fmt *c.Char) *ObserverT {
	return nil
}

/**
 * Bind an integer subject to an Arc's value.
 * @param obj       pointer to Arc
 * @param subject   pointer to Subject
 * @return          pointer to newly-created Observer
 */
// llgo:link (*ObjT).ArcBindValue C.lv_arc_bind_value
func (recv_ *ObjT) ArcBindValue(subject *SubjectT) *ObserverT {
	return nil
}

/**
 * Bind an integer Subject to a Slider's value.
 * @param obj       pointer to Slider
 * @param subject   pointer to Subject
 * @return          pointer to newly-created Observer
 */
// llgo:link (*ObjT).SliderBindValue C.lv_slider_bind_value
func (recv_ *ObjT) SliderBindValue(subject *SubjectT) *ObserverT {
	return nil
}

/**
 * Bind an integer Subject to a Roller's value.
 * @param obj       pointer to Roller
 * @param subject   pointer to Subject
 * @return          pointer to newly-created Observer
 */
// llgo:link (*ObjT).RollerBindValue C.lv_roller_bind_value
func (recv_ *ObjT) RollerBindValue(subject *SubjectT) *ObserverT {
	return nil
}

/**
 * Bind an integer Subject to a Dropdown's value.
 * @param obj       pointer to Dropdown
 * @param subject   pointer to Subject
 * @return          pointer to newly-created Observer
 */
// llgo:link (*ObjT).DropdownBindValue C.lv_dropdown_bind_value
func (recv_ *ObjT) DropdownBindValue(subject *SubjectT) *ObserverT {
	return nil
}

type TlsfT c.Pointer
type PoolT c.Pointer

/* Create/destroy a memory pool. */
//go:linkname TlsfCreate C.lv_tlsf_create
func TlsfCreate(mem c.Pointer) TlsfT

//go:linkname TlsfCreateWithPool C.lv_tlsf_create_with_pool
func TlsfCreateWithPool(mem c.Pointer, bytes c.SizeT) TlsfT

//go:linkname TlsfDestroy C.lv_tlsf_destroy
func TlsfDestroy(tlsf TlsfT)

//go:linkname TlsfGetPool C.lv_tlsf_get_pool
func TlsfGetPool(tlsf TlsfT) PoolT

/* Add/remove memory pools. */
//go:linkname TlsfAddPool C.lv_tlsf_add_pool
func TlsfAddPool(tlsf TlsfT, mem c.Pointer, bytes c.SizeT) PoolT

//go:linkname TlsfRemovePool C.lv_tlsf_remove_pool
func TlsfRemovePool(tlsf TlsfT, pool PoolT)

/* malloc/memalign/realloc/free replacements. */
//go:linkname TlsfMalloc C.lv_tlsf_malloc
func TlsfMalloc(tlsf TlsfT, bytes c.SizeT) c.Pointer

//go:linkname TlsfMemalign C.lv_tlsf_memalign
func TlsfMemalign(tlsf TlsfT, align c.SizeT, bytes c.SizeT) c.Pointer

//go:linkname TlsfRealloc C.lv_tlsf_realloc
func TlsfRealloc(tlsf TlsfT, ptr c.Pointer, size c.SizeT) c.Pointer

//go:linkname TlsfFree C.lv_tlsf_free
func TlsfFree(tlsf TlsfT, ptr c.Pointer) c.SizeT

/* Returns internal block size, not original request size */
//go:linkname TlsfBlockSize C.lv_tlsf_block_size
func TlsfBlockSize(ptr c.Pointer) c.SizeT

/* Overheads/limits of internal structures. */
//go:linkname TlsfSize C.lv_tlsf_size
func TlsfSize() c.SizeT

//go:linkname TlsfAlignSize C.lv_tlsf_align_size
func TlsfAlignSize() c.SizeT

//go:linkname TlsfBlockSizeMin C.lv_tlsf_block_size_min
func TlsfBlockSizeMin() c.SizeT

//go:linkname TlsfBlockSizeMax C.lv_tlsf_block_size_max
func TlsfBlockSizeMax() c.SizeT

//go:linkname TlsfPoolOverhead C.lv_tlsf_pool_overhead
func TlsfPoolOverhead() c.SizeT

//go:linkname TlsfAllocOverhead C.lv_tlsf_alloc_overhead
func TlsfAllocOverhead() c.SizeT

// llgo:type C
type TlsfWalker func(c.Pointer, c.SizeT, c.Int, c.Pointer)

//go:linkname TlsfWalkPool C.lv_tlsf_walk_pool
func TlsfWalkPool(pool PoolT, walker TlsfWalker, user c.Pointer)

/* Returns nonzero if any internal consistency check fails. */
//go:linkname TlsfCheck C.lv_tlsf_check
func TlsfCheck(tlsf TlsfT) c.Int

//go:linkname TlsfCheckPool C.lv_tlsf_check_pool
func TlsfCheckPool(pool PoolT) c.Int

type TimerStateT struct {
	TimerLl            LlT
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
	ResumeCb           TimerHandlerResumeCbT
	ResumeData         c.Pointer
}

/**
 * Init the lv_timer module
 */
//go:linkname TimerCoreInit C.lv_timer_core_init
func TimerCoreInit()

/**
 * Deinit the lv_timer module
 */
//go:linkname TimerCoreDeinit C.lv_timer_core_deinit
func TimerCoreDeinit()

/**********************
 *      TYPEDEFS
 **********************/

type AnimStateT struct {
	AnimListChanged     bool
	AnimRunRound        bool
	AnimVsyncRegistered bool
	Timer               *TimerT
	AnimLl              LlT
}

/**
 * Init the animation module
 */
//go:linkname AnimCoreInit C.lv_anim_core_init
func AnimCoreInit()

/**
 * Deinit the animation module
 */
//go:linkname AnimCoreDeinit C.lv_anim_core_deinit
func AnimCoreDeinit()

/*
 * Set animation use vsync mode.
 * @param enable true: use vsync mode, false: use timer mode.
 */
//go:linkname AnimEnableVsyncMode C.lv_anim_enable_vsync_mode
func AnimEnableVsyncMode(enable bool)

/**********************
 *      TYPEDEFS
 **********************/

type TickStateT struct {
	SysTime    c.Uint32T
	SysIrqFlag c.Uint8T
	TickGetCb  TickGetCbT
	DelayCb    DelayCbT
}

/**
 * Called internally to initialize the draw_buf_handlers in lv_global
 */
//go:linkname DrawBufInitHandlers C.lv_draw_buf_init_handlers
func DrawBufInitHandlers()

/**
 * Get the size of a cache entry.
 * @param node_size     The size of the node in the cache.
 * @return              The size of the cache entry.
 */
//go:linkname CacheEntryGetSize C.lv_cache_entry_get_size
func CacheEntryGetSize(node_size c.Uint32T) c.Uint32T

/**
 * Get the reference count of a cache entry.
 * @param entry        The cache entry to get the reference count of.
 * @return             The reference count of the cache entry.
 */
// llgo:link (*CacheEntryT).CacheEntryGetRef C.lv_cache_entry_get_ref
func (recv_ *CacheEntryT) CacheEntryGetRef() c.Int32T {
	return 0
}

/**
 * Get the node size of a cache entry. Which is the same size with lv_cache_entry_get_size()'s node_size parameter.
 * @param entry        The cache entry to get the node size of.
 * @return             The node size of the cache entry.
 */
// llgo:link (*CacheEntryT).CacheEntryGetNodeSize C.lv_cache_entry_get_node_size
func (recv_ *CacheEntryT) CacheEntryGetNodeSize() c.Uint32T {
	return 0
}

/**
 * Check if a cache entry is invalid.
 * @param entry        The cache entry to check.
 * @return             True: the cache entry is invalid. False: the cache entry is valid.
 */
// llgo:link (*CacheEntryT).CacheEntryIsInvalid C.lv_cache_entry_is_invalid
func (recv_ *CacheEntryT) CacheEntryIsInvalid() bool {
	return false
}

/**
 * Get the data of a cache entry.
 * @param entry        The cache entry to get the data of.
 * @return             The pointer to the data of the cache entry.
 */
// llgo:link (*CacheEntryT).CacheEntryGetData C.lv_cache_entry_get_data
func (recv_ *CacheEntryT) CacheEntryGetData() c.Pointer {
	return nil
}

/**
 * Get the cache instance of a cache entry.
 * @param entry        The cache entry to get the cache instance of.
 * @return             The pointer to the cache instance of the cache entry.
 */
// llgo:link (*CacheEntryT).CacheEntryGetCache C.lv_cache_entry_get_cache
func (recv_ *CacheEntryT) CacheEntryGetCache() *CacheT {
	return nil
}

/**
 * Get the cache entry of a data. The data should be allocated by the cache instance.
 * @param data         The data to get the cache entry of.
 * @param node_size    The size of the node in the cache.
 * @return             The pointer to the cache entry of the data.
 */
//go:linkname CacheEntryGetEntry C.lv_cache_entry_get_entry
func CacheEntryGetEntry(data c.Pointer, node_size c.Uint32T) *CacheEntryT

/**
 * Allocate a cache entry.
 * @param node_size    The size of the node in the cache.
 * @param cache        The cache instance to allocate the cache entry from.
 * @return             The pointer to the allocated cache entry.
 */
//go:linkname CacheEntryAlloc C.lv_cache_entry_alloc
func CacheEntryAlloc(node_size c.Uint32T, cache *CacheT) *CacheEntryT

/**
 * Initialize a cache entry.
 * @param entry        The cache entry to initialize.
 * @param cache        The cache instance to allocate the cache entry from.
 * @param node_size    The size of the node in the cache.
 */
// llgo:link (*CacheEntryT).CacheEntryInit C.lv_cache_entry_init
func (recv_ *CacheEntryT) CacheEntryInit(cache *CacheT, node_size c.Uint32T) {
}

/**
 * Deallocate a cache entry. And the data of the cache entry will be freed.
 * @param entry        The cache entry to deallocate.
 */
// llgo:link (*CacheEntryT).CacheEntryDelete C.lv_cache_entry_delete
func (recv_ *CacheEntryT) CacheEntryDelete() {
}

type CacheReserveCondResT c.Int

const (
	CACHE_RESERVE_COND_OK          CacheReserveCondResT = 0
	CACHE_RESERVE_COND_TOO_LARGE   CacheReserveCondResT = 1
	CACHE_RESERVE_COND_NEED_VICTIM CacheReserveCondResT = 2
	CACHE_RESERVE_COND_ERROR       CacheReserveCondResT = 3
)

type X_lvCacheOpsT struct {
	CompareCb CacheCompareCbT
	CreateCb  CacheCreateCbT
	FreeCb    CacheFreeCbT
}

type X_lvCacheClassT struct {
	AllocCb       CacheAllocCbT
	InitCb        CacheInitCbT
	DestroyCb     CacheDestroyCbT
	GetCb         CacheGetCbT
	AddCb         CacheAddCbT
	RemoveCb      CacheRemoveCbT
	DropCb        CacheDropCbT
	DropAllCb     CacheDropAllCbT
	GetVictimCb   CacheGetVictimCb
	ReserveCondCb CacheReserveCondCb
	IterCreateCb  CacheIterCreateCb
}
type CacheOpsT X_lvCacheOpsT
type CacheClassT X_lvCacheClassT
type CacheCompareResT c.Int8T

// llgo:type C
type CacheCreateCbT func(c.Pointer, c.Pointer) bool

// llgo:type C
type CacheFreeCbT func(c.Pointer, c.Pointer)

// llgo:type C
type CacheCompareCbT func(c.Pointer, c.Pointer) CacheCompareResT

// llgo:type C
type CacheAllocCbT func() c.Pointer

// llgo:type C
type CacheInitCbT func(*CacheT) bool

// llgo:type C
type CacheDestroyCbT func(*CacheT, c.Pointer)

// llgo:type C
type CacheGetCbT func(*CacheT, c.Pointer, c.Pointer) *CacheEntryT

// llgo:type C
type CacheAddCbT func(*CacheT, c.Pointer, c.Pointer) *CacheEntryT

// llgo:type C
type CacheRemoveCbT func(*CacheT, *CacheEntryT, c.Pointer)

// llgo:type C
type CacheDropCbT func(*CacheT, c.Pointer, c.Pointer)

// llgo:type C
type CacheDropAllCbT func(*CacheT, c.Pointer)

// llgo:type C
type CacheGetVictimCb func(*CacheT, c.Pointer) *CacheEntryT

// llgo:type C
type CacheReserveCondCb func(*CacheT, c.Pointer, c.SizeT, c.Pointer) CacheReserveCondResT

// llgo:type C
type CacheIterCreateCb func(*CacheT) *IterT

/*-----------------
 * Cache entry slot
 *----------------*/

type X_lvCacheSlotSizeT struct {
	Size c.SizeT
}
type CacheSlotSizeT X_lvCacheSlotSizeT

/**
 * Initialize image header cache.
 * @param  count initial size of the cache in count of image headers.
 * @return LV_RESULT_OK: initialization succeeded, LV_RESULT_INVALID: failed.
 */
//go:linkname ImageHeaderCacheInit C.lv_image_header_cache_init
func ImageHeaderCacheInit(count c.Uint32T) ResultT

/**
 * Resize image header cache.
 * If set to 0, the cache is disabled.
 * @param count  new max count of cached image headers.
 * @param evict_now true: evict the image headers should be removed by the eviction policy, false: wait for the next cache cleanup.
 */
//go:linkname ImageHeaderCacheResize C.lv_image_header_cache_resize
func ImageHeaderCacheResize(count c.Uint32T, evict_now bool)

/**
 * Invalidate image header cache. Use NULL to invalidate all image headers.
 * It's also automatically called when an image is invalidated.
 * @param src pointer to an image source.
 */
//go:linkname ImageHeaderCacheDrop C.lv_image_header_cache_drop
func ImageHeaderCacheDrop(src c.Pointer)

/**
 * Return true if the image header cache is enabled.
 * @return true: enabled, false: disabled.
 */
//go:linkname ImageHeaderCacheIsEnabled C.lv_image_header_cache_is_enabled
func ImageHeaderCacheIsEnabled() bool

/**
 * Create an iterator to iterate over the image header cache.
 * @return an iterator to iterate over the image header cache.
 */
//go:linkname ImageHeaderCacheIterCreate C.lv_image_header_cache_iter_create
func ImageHeaderCacheIterCreate() *IterT

/**
 * Dump the content of the image header cache in a human-readable format with cache order.
 */
//go:linkname ImageHeaderCacheDump C.lv_image_header_cache_dump
func ImageHeaderCacheDump()

/**
 * Initialize image cache.
 * @param  size size of the cache in bytes.
 * @return LV_RESULT_OK: initialization succeeded, LV_RESULT_INVALID: failed.
 */
//go:linkname ImageCacheInit C.lv_image_cache_init
func ImageCacheInit(size c.Uint32T) ResultT

/**
 * Resize image cache.
 * If set to 0, the cache will be disabled.
 * @param new_size  new size of the cache in bytes.
 * @param evict_now true: evict the images should be removed by the eviction policy, false: wait for the next cache cleanup.
 */
//go:linkname ImageCacheResize C.lv_image_cache_resize
func ImageCacheResize(new_size c.Uint32T, evict_now bool)

/**
 * Invalidate image cache. Use NULL to invalidate all images.
 * @param src pointer to an image source.
 */
//go:linkname ImageCacheDrop C.lv_image_cache_drop
func ImageCacheDrop(src c.Pointer)

/**
 * Return true if the image cache is enabled.
 * @return true: enabled, false: disabled.
 */
//go:linkname ImageCacheIsEnabled C.lv_image_cache_is_enabled
func ImageCacheIsEnabled() bool

/**
 * Create an iterator to iterate over the image cache.
 * @return an iterator to iterate over the image cache.
 */
//go:linkname ImageCacheIterCreate C.lv_image_cache_iter_create
func ImageCacheIterCreate() *IterT

/**
 * Dump the content of the image cache in a human-readable format with cache order.
 */
//go:linkname ImageCacheDump C.lv_image_cache_dump
func ImageCacheDump()

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
// llgo:link (*CacheClassT).CacheCreate C.lv_cache_create
func (recv_ *CacheClassT) CacheCreate(node_size c.SizeT, max_size c.SizeT, ops CacheOpsT) *CacheT {
	return nil
}

/**
 * Destroy a cache object.
 * @param cache         The cache object pointer to destroy.
 * @param user_data     A user data pointer that will be passed to the free callback.
 */
// llgo:link (*CacheT).CacheDestroy C.lv_cache_destroy
func (recv_ *CacheT) CacheDestroy(user_data c.Pointer) {
}

/**
 * Acquire a cache entry with the given key. If entry not in cache, it will return `NULL` (not found).
 * If the entry is found, it's priority will be changed by the cache's policy. And the `lv_cache_entry_t::ref_cnt` will be incremented.
 * @param cache         The cache object pointer to acquire the entry.
 * @param key           The key of the entry to acquire.
 * @param user_data     A user data pointer that will be passed to the create callback.
 * @return              Returns a pointer to the acquired cache entry on success with `lv_cache_entry_t::ref_cnt` incremented, `NULL` on error.
 */
// llgo:link (*CacheT).CacheAcquire C.lv_cache_acquire
func (recv_ *CacheT) CacheAcquire(key c.Pointer, user_data c.Pointer) *CacheEntryT {
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
// llgo:link (*CacheT).CacheAcquireOrCreate C.lv_cache_acquire_or_create
func (recv_ *CacheT) CacheAcquireOrCreate(key c.Pointer, user_data c.Pointer) *CacheEntryT {
	return nil
}

/**
 * Add a new cache entry with the given key and data. If the cache is full, the cache's policy will be used to evict an entry.
 * @param cache         The cache object pointer to add the entry.
 * @param key           The key of the entry to add.
 * @param user_data     A user data pointer that will be passed to the create callback.
 * @return              Returns a pointer to the added cache entry on success with `lv_cache_entry_t::ref_cnt` incremented, `NULL` on error.
 */
// llgo:link (*CacheT).CacheAdd C.lv_cache_add
func (recv_ *CacheT) CacheAdd(key c.Pointer, user_data c.Pointer) *CacheEntryT {
	return nil
}

/**
 * Release a cache entry. The `lv_cache_entry_t::ref_cnt` will be decremented. If the `lv_cache_entry_t::ref_cnt` is zero, it will issue an error.
 * If the entry passed to this function is the last reference to the data and the entry is marked as invalid, the cache's policy will be used to evict the entry.
 * @param cache         The cache object pointer to release the entry.
 * @param entry         The cache entry pointer to release.
 * @param user_data     A user data pointer that will be passed to the free callback.
 */
// llgo:link (*CacheT).CacheRelease C.lv_cache_release
func (recv_ *CacheT) CacheRelease(entry *CacheEntryT, user_data c.Pointer) {
}

/**
 * Reserve a certain amount of memory/count in the cache. This function is useful when you want to reserve a certain amount of memory/count in advance,
 * for example, when you know that you will need it later.
 * When the current cache size is max than the reserved size, the function will evict entries until the reserved size is reached.
 * @param cache         The cache object pointer to reserve.
 * @param reserved_size The amount of memory/count to reserve.
 * @param user_data     A user data pointer that will be passed to the free callback.
 */
// llgo:link (*CacheT).CacheReserve C.lv_cache_reserve
func (recv_ *CacheT) CacheReserve(reserved_size c.Uint32T, user_data c.Pointer) {
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
// llgo:link (*CacheT).CacheDrop C.lv_cache_drop
func (recv_ *CacheT) CacheDrop(key c.Pointer, user_data c.Pointer) {
}

/**
 * Drop all cache entries. All entries will be removed from the cache and their data will be freed when the last reference to them is released.
 * @note If some entries are still referenced by other objects, it will issue an error. And this case shouldn't happen in normal cases..
 * @param cache         The cache object pointer to drop all entries.
 * @param user_data     A user data pointer that will be passed to the free callback.
 */
// llgo:link (*CacheT).CacheDropAll C.lv_cache_drop_all
func (recv_ *CacheT) CacheDropAll(user_data c.Pointer) {
}

/**
 * Evict one entry from the cache. The eviction policy will be used to select the entry to evict.
 * @param cache         The cache object pointer to evict an entry.
 * @param user_data     A user data pointer that will be passed to the free callback.
 * @return              Returns true if an entry is evicted, false if no entry is evicted.
 */
// llgo:link (*CacheT).CacheEvictOne C.lv_cache_evict_one
func (recv_ *CacheT) CacheEvictOne(user_data c.Pointer) bool {
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
// llgo:link (*CacheT).CacheSetMaxSize C.lv_cache_set_max_size
func (recv_ *CacheT) CacheSetMaxSize(max_size c.SizeT, user_data c.Pointer) {
}

/**
 * Get the maximum size of the cache.
 * @param cache         The cache object pointer to get the maximum size.
 * @param user_data     A user data pointer that will be passed to the free callback.
 * @return              Returns the maximum size of the cache.
 */
// llgo:link (*CacheT).CacheGetMaxSize C.lv_cache_get_max_size
func (recv_ *CacheT) CacheGetMaxSize(user_data c.Pointer) c.SizeT {
	return 0
}

/**
 * Get the current size of the cache.
 * @param cache         The cache object pointer to get the current size.
 * @param user_data     A user data pointer that will be passed to the free callback.
 * @return              Returns the current size of the cache.
 */
// llgo:link (*CacheT).CacheGetSize C.lv_cache_get_size
func (recv_ *CacheT) CacheGetSize(user_data c.Pointer) c.SizeT {
	return 0
}

/**
 * Get the free size of the cache.
 * @param cache         The cache object pointer to get the free size.
 * @param user_data     A user data pointer that will be passed to the free callback.
 * @return              Returns the free size of the cache.
 */
// llgo:link (*CacheT).CacheGetFreeSize C.lv_cache_get_free_size
func (recv_ *CacheT) CacheGetFreeSize(user_data c.Pointer) c.SizeT {
	return 0
}

/**
 * Return true if the cache is enabled.
 * Disabled cache means that when the max_size of the cache is 0. In this case, all cache operations will be no-op.
 * @param cache         The cache object pointer to check if it's disabled.
 * @return              Returns true if the cache is enabled, false otherwise.
 */
// llgo:link (*CacheT).CacheIsEnabled C.lv_cache_is_enabled
func (recv_ *CacheT) CacheIsEnabled() bool {
	return false
}

/**
 * Set the compare callback of the cache.
 * @param cache         The cache object pointer to set the compare callback.
 * @param compare_cb    The compare callback to set.
 * @param user_data     A user data pointer.
 */
// llgo:link (*CacheT).CacheSetCompareCb C.lv_cache_set_compare_cb
func (recv_ *CacheT) CacheSetCompareCb(compare_cb CacheCompareCbT, user_data c.Pointer) {
}

/**
 * Set the create callback of the cache.
 * @param cache         The cache object pointer to set the create callback.
 * @param alloc_cb      The create callback to set.
 * @param user_data     A user data pointer.
 */
// llgo:link (*CacheT).CacheSetCreateCb C.lv_cache_set_create_cb
func (recv_ *CacheT) CacheSetCreateCb(alloc_cb CacheCreateCbT, user_data c.Pointer) {
}

/**
 * Set the free callback of the cache.
 * @param cache         The cache object pointer to set the free callback.
 * @param free_cb       The free callback to set.
 * @param user_data     A user data pointer.
 */
// llgo:link (*CacheT).CacheSetFreeCb C.lv_cache_set_free_cb
func (recv_ *CacheT) CacheSetFreeCb(free_cb CacheFreeCbT, user_data c.Pointer) {
}

/**
 * Give a name for a cache object. Only the pointer of the string is saved.
 * @param cache         The cache object pointer to set the name.
 * @param name          The name of the cache.
 */
// llgo:link (*CacheT).CacheSetName C.lv_cache_set_name
func (recv_ *CacheT) CacheSetName(name *c.Char) {
}

/**
 * Get the name of a cache object.
 * @param cache         The cache object pointer to get the name.
 * @return              Returns the name of the cache.
 */
// llgo:link (*CacheT).CacheGetName C.lv_cache_get_name
func (recv_ *CacheT) CacheGetName() *c.Char {
	return nil
}

/**
 * Create an iterator for the cache object. The iterator is used to iterate over all cache entries.
 * @param cache         The cache object pointer to create the iterator.
 * @return              Returns a pointer to the created iterator on success, `NULL` on error.
 */
// llgo:link (*CacheT).CacheIterCreate C.lv_cache_iter_create
func (recv_ *CacheT) CacheIterCreate() *IterT {
	return nil
}

type DrawGlobalInfoT struct {
	UnitHead            *DrawUnitT
	UnitCnt             c.Uint32T
	UsedMemoryForLayers c.Uint32T
	DispatchReq         c.Int
	CircleCacheMutex    MutexT
	TaskRunning         bool
}

/**********************
 *      TYPEDEFS
 **********************/

type DrawSwThreadDscT struct {
	TaskAct    *DrawTaskT
	Thread     ThreadT
	Sync       ThreadSyncT
	DrawUnit   *DrawUnitT
	Idx        c.Uint32T
	Inited     bool
	ExitStatus bool
}

/**********************
 *      TYPEDEFS
 **********************/

type DrawSwMaskRadiusCircleDscT struct {
	Buf         *c.Uint8T
	CirOpa      *OpaT
	XStartOnY   *c.Uint16T
	OpaStartOnY *c.Uint16T
	Life        c.Int32T
	UsedCnt     c.Uint32T
	Radius      c.Int32T
}
type DrawSwMaskRadiusCircleDscArrT [4]DrawSwMaskRadiusCircleDscT

/**
 * Called by LVGL the rendering of a screen is ready to clean up
 * the temporal (cache) data of the masks
 */
//go:linkname DrawSwMaskCleanup C.lv_draw_sw_mask_cleanup
func DrawSwMaskCleanup()

/**********************
 *      TYPEDEFS
 **********************/

type TlsfStateT struct {
	Tlsf    TlsfT
	CurUsed c.SizeT
	MaxUsed c.SizeT
	PoolLl  LlT
}

/**********************
 *      TYPEDEFS
 **********************/

type LayoutDscT struct {
	Cb       LayoutUpdateCbT
	UserData c.Pointer
}

/**********************
 * GLOBAL PROTOTYPES
 **********************/
//go:linkname LayoutInit C.lv_layout_init
func LayoutInit()

//go:linkname LayoutDeinit C.lv_layout_deinit
func LayoutDeinit()

/**
 * Update the layout of a widget
 * @param obj   pointer to a widget
 */
// llgo:link (*ObjT).LayoutApply C.lv_layout_apply
func (recv_ *ObjT) LayoutApply() {
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
	DispLl                         LlT
	DispRefresh                    *DisplayT
	DispDefault                    *DisplayT
	StyleTransLl                   LlT
	StyleRefresh                   bool
	StyleCustomTableSize           c.Uint32T
	StyleLastCustomPropId          c.Uint32T
	StyleCustomPropFlagLookupTable *c.Uint8T
	GroupLl                        LlT
	GroupDefault                   *GroupT
	IndevLl                        LlT
	IndevActive                    *IndevT
	IndevObjActive                 *ObjT
	LayoutCount                    c.Uint32T
	LayoutList                     *LayoutDscT
	LayoutUpdateMutex              bool
	MemoryZero                     c.Uint32T
	MathRandSeed                   c.Uint32T
	EventHeader                    *EventT
	EventLastRegisterId            c.Uint32T
	TimerState                     TimerStateT
	AnimState                      AnimStateT
	TickState                      TickStateT
	DrawBufHandlers                DrawBufHandlersT
	FontDrawBufHandlers            DrawBufHandlersT
	ImageCacheDrawBufHandlers      DrawBufHandlersT
	ImgDecoderLl                   LlT
	ImgCache                       *CacheT
	ImgHeaderCache                 *CacheT
	DrawInfo                       DrawGlobalInfoT
	DrawSwBlendHandlerLl           LlT
	SwCircleCache                  DrawSwMaskRadiusCircleDscArrT
	ThemeSimple                    c.Pointer
	ThemeDefault                   c.Pointer
	ThemeMono                      c.Pointer
	TlsfState                      TlsfStateT
	FsdrvLl                        LlT
	SpanSnippetStack               *X_snippetStack
	ObjidArray                     c.Pointer
	ObjidCount                     c.Uint32T
	UserData                       c.Pointer
}
type GlobalT X_lvGlobalT

/**
 * Handle scrolling. Called by LVGL during input device processing
 * @param indev      pointer to an input device
 */
// llgo:link (*IndevT).IndevScrollHandler C.lv_indev_scroll_handler
func (recv_ *IndevT) IndevScrollHandler() {
}

/**
 * Handle throwing after scrolling. Called by LVGL during input device processing
 * @param indev      pointer to an input device
 */
// llgo:link (*IndevT).IndevScrollThrowHandler C.lv_indev_scroll_throw_handler
func (recv_ *IndevT) IndevScrollThrowHandler() {
}

/**
 * Predict where would a scroll throw end
 * @param indev     pointer to an input device
 * @param dir `     LV_DIR_VER` or `LV_DIR_HOR`
 * @return          the difference compared to the current position when the throw would be finished
 */
// llgo:link (*IndevT).IndevScrollThrowPredict C.lv_indev_scroll_throw_predict
func (recv_ *IndevT) IndevScrollThrowPredict(dir DirT) c.Int32T {
	return 0
}

/**
 * Get the distance of the nearest snap point
 * @param obj       the object on which snap points should be found
 * @param p         save the distance of the found snap point there
 */
// llgo:link (*ObjT).IndevScrollGetSnapDist C.lv_indev_scroll_get_snap_dist
func (recv_ *ObjT) IndevScrollGetSnapDist(p *PointT) {
}

/**
 * Find a scrollable object based on the current scroll vector in the indev.
 * In handles scroll propagation to the parent if needed, and scroll directions too.
 * @param indev     pointer to an indev
 * @return          the found scrollable object or NULL if not found.
 */
// llgo:link (*IndevT).IndevFindScrollObj C.lv_indev_find_scroll_obj
func (recv_ *IndevT) IndevFindScrollObj() *ObjT {
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
//go:linkname TextGetNextLine C.lv_text_get_next_line
func TextGetNextLine(txt *c.Char, len c.Uint32T, font *FontT, letter_space c.Int32T, max_width c.Int32T, used_width *c.Int32T, flag TextFlagT) c.Uint32T

/**
 * Insert a string into another
 * @param txt_buf the original text (must be big enough for the result text and NULL terminated)
 * @param pos position to insert (0: before the original text, 1: after the first char etc.)
 * @param ins_txt text to insert, must be '\0' terminated
 */
//go:linkname TextIns C.lv_text_ins
func TextIns(txt_buf *c.Char, pos c.Uint32T, ins_txt *c.Char)

/**
 * Delete a part of a string
 * @param txt string to modify, must be '\0' terminated and should point to a heap or stack frame, not read-only memory.
 * @param pos position where to start the deleting (0: before the first char, 1: after the first
 * char etc.)
 * @param len number of characters to delete
 */
//go:linkname TextCut C.lv_text_cut
func TextCut(txt *c.Char, pos c.Uint32T, len c.Uint32T)

/**
 * return a new formatted text. Memory will be allocated to store the text.
 * @param fmt `printf`-like format
 * @param ap items to print

 * @return pointer to the allocated text string.
 */
//go:linkname TextSetTextVfmt C.lv_text_set_text_vfmt
func TextSetTextVfmt(fmt *c.Char, ap c.VaList) *c.Char

/**
 * Decode two encoded character from a string.
 * @param txt pointer to '\0' terminated string
 * @param letter the first decoded Unicode character or 0 on invalid data code
 * @param letter_next the second decoded Unicode character or 0 on invalid data code
 * @param ofs start index in 'txt' where to start.
 *                After the call it will point to the next encoded char in 'txt'.
 *                NULL to use txt[0] as index
 */
//go:linkname TextEncodedLetterNext2 C.lv_text_encoded_letter_next_2
func TextEncodedLetterNext2(txt *c.Char, letter *c.Uint32T, letter_next *c.Uint32T, ofs *c.Uint32T)

/**********************
 * GLOBAL PROTOTYPES
 **********************/
// llgo:link (*CacheEntryT).CacheEntryResetRef C.lv_cache_entry_reset_ref
func (recv_ *CacheEntryT) CacheEntryResetRef() {
}

// llgo:link (*CacheEntryT).CacheEntryIncRef C.lv_cache_entry_inc_ref
func (recv_ *CacheEntryT) CacheEntryIncRef() {
}

// llgo:link (*CacheEntryT).CacheEntryDecRef C.lv_cache_entry_dec_ref
func (recv_ *CacheEntryT) CacheEntryDecRef() {
}

// llgo:link (*CacheEntryT).CacheEntrySetNodeSize C.lv_cache_entry_set_node_size
func (recv_ *CacheEntryT) CacheEntrySetNodeSize(node_size c.Uint32T) {
}

// llgo:link (*CacheEntryT).CacheEntrySetInvalid C.lv_cache_entry_set_invalid
func (recv_ *CacheEntryT) CacheEntrySetInvalid(is_invalid bool) {
}

// llgo:link (*CacheEntryT).CacheEntrySetCache C.lv_cache_entry_set_cache
func (recv_ *CacheEntryT) CacheEntrySetCache(cache *CacheT) {
}

// llgo:link (*CacheEntryT).CacheEntryAcquireData C.lv_cache_entry_acquire_data
func (recv_ *CacheEntryT) CacheEntryAcquireData() c.Pointer {
	return nil
}

// llgo:link (*CacheEntryT).CacheEntryReleaseData C.lv_cache_entry_release_data
func (recv_ *CacheEntryT) CacheEntryReleaseData(user_data c.Pointer) {
}

type ImageAlignT c.Int

const (
	IMAGE_ALIGN_DEFAULT        ImageAlignT = 0
	IMAGE_ALIGN_TOP_LEFT       ImageAlignT = 1
	IMAGE_ALIGN_TOP_MID        ImageAlignT = 2
	IMAGE_ALIGN_TOP_RIGHT      ImageAlignT = 3
	IMAGE_ALIGN_BOTTOM_LEFT    ImageAlignT = 4
	IMAGE_ALIGN_BOTTOM_MID     ImageAlignT = 5
	IMAGE_ALIGN_BOTTOM_RIGHT   ImageAlignT = 6
	IMAGE_ALIGN_LEFT_MID       ImageAlignT = 7
	IMAGE_ALIGN_RIGHT_MID      ImageAlignT = 8
	IMAGE_ALIGN_CENTER         ImageAlignT = 9
	IMAGE_ALIGN_AUTO_TRANSFORM ImageAlignT = 10
	IMAGE_ALIGN_STRETCH        ImageAlignT = 11
	IMAGE_ALIGN_TILE           ImageAlignT = 12
	IMAGE_ALIGN_CONTAIN        ImageAlignT = 13
	IMAGE_ALIGN_COVER          ImageAlignT = 14
)

/**
 * Create an image object
 * @param parent pointer to an object, it will be the parent of the new image
 * @return pointer to the created image
 */
// llgo:link (*ObjT).ImageCreate C.lv_image_create
func (recv_ *ObjT) ImageCreate() *ObjT {
	return nil
}

/**
 * Set the image data to display on the object
 * @param obj       pointer to an image object
 * @param src       1) pointer to an ::lv_image_dsc_t descriptor (converted by LVGL's image converter) (e.g. &my_img) or
 *                  2) path to an image file (e.g. "S:/dir/img.bin")or
 *                  3) a SYMBOL (e.g. LV_SYMBOL_OK)
 */
// llgo:link (*ObjT).ImageSetSrc C.lv_image_set_src
func (recv_ *ObjT) ImageSetSrc(src c.Pointer) {
}

/**
 * Set an offset for the source of an image so the image will be displayed from the new origin.
 * @param obj       pointer to an image
 * @param x         the new offset along x axis.
 */
// llgo:link (*ObjT).ImageSetOffsetX C.lv_image_set_offset_x
func (recv_ *ObjT) ImageSetOffsetX(x c.Int32T) {
}

/**
 * Set an offset for the source of an image.
 * so the image will be displayed from the new origin.
 * @param obj       pointer to an image
 * @param y         the new offset along y axis.
 */
// llgo:link (*ObjT).ImageSetOffsetY C.lv_image_set_offset_y
func (recv_ *ObjT) ImageSetOffsetY(y c.Int32T) {
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
// llgo:link (*ObjT).ImageSetRotation C.lv_image_set_rotation
func (recv_ *ObjT) ImageSetRotation(angle c.Int32T) {
}

/**
 * Set the rotation center of the image.
 * The image will be rotated around this point.
 * x, y can be set with value of LV_PCT, lv_image_get_pivot will return the true pixel coordinate of pivot in this case.
 * @param obj       pointer to an image object
 * @param x         rotation center x of the image
 * @param y         rotation center y of the image
 */
// llgo:link (*ObjT).ImageSetPivot C.lv_image_set_pivot
func (recv_ *ObjT) ImageSetPivot(x c.Int32T, y c.Int32T) {
}

/**
 * Set the rotation horizontal center of the image.
 * @param obj       pointer to an image object
 * @param x         rotation center x of the image, or lv_pct()
 */
// llgo:link (*ObjT).ImageSetPivotX C.lv_image_set_pivot_x
func (recv_ *ObjT) ImageSetPivotX(x c.Int32T) {
}

/**
 * Set the rotation vertical center of the image.
 * @param obj       pointer to an image object
 * @param y         rotation center y of the image, or lv_pct()
 */
// llgo:link (*ObjT).ImageSetPivotY C.lv_image_set_pivot_y
func (recv_ *ObjT) ImageSetPivotY(y c.Int32T) {
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
// llgo:link (*ObjT).ImageSetScale C.lv_image_set_scale
func (recv_ *ObjT) ImageSetScale(zoom c.Uint32T) {
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
// llgo:link (*ObjT).ImageSetScaleX C.lv_image_set_scale_x
func (recv_ *ObjT) ImageSetScaleX(zoom c.Uint32T) {
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
// llgo:link (*ObjT).ImageSetScaleY C.lv_image_set_scale_y
func (recv_ *ObjT) ImageSetScaleY(zoom c.Uint32T) {
}

/**
 * Set the blend mode of an image.
 * @param obj           pointer to an image object
 * @param blend_mode    the new blend mode
 */
// llgo:link (*ObjT).ImageSetBlendMode C.lv_image_set_blend_mode
func (recv_ *ObjT) ImageSetBlendMode(blend_mode BlendModeT) {
}

/**
 * Enable/disable anti-aliasing for the transformations (rotate, zoom) or not.
 * The quality is better with anti-aliasing looks better but slower.
 * @param obj       pointer to an image object
 * @param antialias true: anti-aliased; false: not anti-aliased
 */
// llgo:link (*ObjT).ImageSetAntialias C.lv_image_set_antialias
func (recv_ *ObjT) ImageSetAntialias(antialias bool) {
}

/**
 * Set the image object size mode.
 * @param obj       pointer to an image object
 * @param align     the new align mode.
 * @note            if image_align is `LV_IMAGE_ALIGN_STRETCH` or `LV_IMAGE_ALIGN_FIT`
 *                  rotation, scale and pivot will be overwritten and controlled internally.
 */
// llgo:link (*ObjT).ImageSetInnerAlign C.lv_image_set_inner_align
func (recv_ *ObjT) ImageSetInnerAlign(align ImageAlignT) {
}

/**
 * Set an A8 bitmap mask for the image.
 * @param obj       pointer to an image object
 * @param src       an lv_image_dsc_t bitmap mask source.
 */
// llgo:link (*ObjT).ImageSetBitmapMapSrc C.lv_image_set_bitmap_map_src
func (recv_ *ObjT) ImageSetBitmapMapSrc(src *ImageDscT) {
}

/**
 * Get the source of the image
 * @param obj       pointer to an image object
 * @return          the image source (symbol, file name or ::lv-img_dsc_t for C arrays)
 */
// llgo:link (*ObjT).ImageGetSrc C.lv_image_get_src
func (recv_ *ObjT) ImageGetSrc() c.Pointer {
	return nil
}

/**
 * Get the offset's x attribute of the image object.
 * @param obj       pointer to an image
 * @return          offset X value.
 */
// llgo:link (*ObjT).ImageGetOffsetX C.lv_image_get_offset_x
func (recv_ *ObjT) ImageGetOffsetX() c.Int32T {
	return 0
}

/**
 * Get the offset's y attribute of the image object.
 * @param obj       pointer to an image
 * @return          offset Y value.
 */
// llgo:link (*ObjT).ImageGetOffsetY C.lv_image_get_offset_y
func (recv_ *ObjT) ImageGetOffsetY() c.Int32T {
	return 0
}

/**
 * Get the rotation of the image.
 * @param obj       pointer to an image object
 * @return          rotation in 0.1 degrees (0..3600)
 * @note            if image_align is `LV_IMAGE_ALIGN_STRETCH` or  `LV_IMAGE_ALIGN_FIT`
 *                  rotation will be set to 0 automatically.
 */
// llgo:link (*ObjT).ImageGetRotation C.lv_image_get_rotation
func (recv_ *ObjT) ImageGetRotation() c.Int32T {
	return 0
}

/**
 * Get the pivot (rotation center) of the image.
 * If pivot is set with LV_PCT, convert it to px before return.
 * @param obj       pointer to an image object
 * @param pivot     store the rotation center here
 */
// llgo:link (*ObjT).ImageGetPivot C.lv_image_get_pivot
func (recv_ *ObjT) ImageGetPivot(pivot *PointT) {
}

/**
 * Get the zoom factor of the image.
 * @param obj       pointer to an image object
 * @return          zoom factor (256: no zoom)
 */
// llgo:link (*ObjT).ImageGetScale C.lv_image_get_scale
func (recv_ *ObjT) ImageGetScale() c.Int32T {
	return 0
}

/**
 * Get the horizontal zoom factor of the image.
 * @param obj       pointer to an image object
 * @return          zoom factor (256: no zoom)
 */
// llgo:link (*ObjT).ImageGetScaleX C.lv_image_get_scale_x
func (recv_ *ObjT) ImageGetScaleX() c.Int32T {
	return 0
}

/**
 * Get the vertical zoom factor of the image.
 * @param obj       pointer to an image object
 * @return          zoom factor (256: no zoom)
 */
// llgo:link (*ObjT).ImageGetScaleY C.lv_image_get_scale_y
func (recv_ *ObjT) ImageGetScaleY() c.Int32T {
	return 0
}

/**
 * Get the width of an image before any transformations.
 * @param obj Pointer to an image object.
 * @return The width of the image.
 */
// llgo:link (*ObjT).ImageGetSrcWidth C.lv_image_get_src_width
func (recv_ *ObjT) ImageGetSrcWidth() c.Int32T {
	return 0
}

/**
 * Get the height of an image before any transformations.
 * @param obj Pointer to an image object.
 * @return The height of the image.
 */
// llgo:link (*ObjT).ImageGetSrcHeight C.lv_image_get_src_height
func (recv_ *ObjT) ImageGetSrcHeight() c.Int32T {
	return 0
}

/**
 * Get the transformed width of an image object.
 * @param obj Pointer to an image object.
 * @return The transformed width of the image.
 */
// llgo:link (*ObjT).ImageGetTransformedWidth C.lv_image_get_transformed_width
func (recv_ *ObjT) ImageGetTransformedWidth() c.Int32T {
	return 0
}

/**
 * Get the transformed height of an image object.
 * @param obj Pointer to an image object.
 * @return The transformed height of the image.
 */
// llgo:link (*ObjT).ImageGetTransformedHeight C.lv_image_get_transformed_height
func (recv_ *ObjT) ImageGetTransformedHeight() c.Int32T {
	return 0
}

/**
 * Get the current blend mode of the image
 * @param obj       pointer to an image object
 * @return          the current blend mode
 */
// llgo:link (*ObjT).ImageGetBlendMode C.lv_image_get_blend_mode
func (recv_ *ObjT) ImageGetBlendMode() BlendModeT {
	return 0
}

/**
 * Get whether the transformations (rotate, zoom) are anti-aliased or not
 * @param obj       pointer to an image object
 * @return          true: anti-aliased; false: not anti-aliased
 */
// llgo:link (*ObjT).ImageGetAntialias C.lv_image_get_antialias
func (recv_ *ObjT) ImageGetAntialias() bool {
	return false
}

/**
 * Get the size mode of the image
 * @param obj       pointer to an image object
 * @return          element of `lv_image_align_t`
 */
// llgo:link (*ObjT).ImageGetInnerAlign C.lv_image_get_inner_align
func (recv_ *ObjT) ImageGetInnerAlign() ImageAlignT {
	return 0
}

/**
 * Get the bitmap mask source.
 * @param obj       pointer to an image object
 * @return          an lv_image_dsc_t bitmap mask source.
 */
// llgo:link (*ObjT).ImageGetBitmapMapSrc C.lv_image_get_bitmap_map_src
func (recv_ *ObjT) ImageGetBitmapMapSrc() *ImageDscT {
	return nil
}

/**
 * Create a canvas object
 * @param parent     pointer to an object, it will be the parent of the new canvas
 * @return           pointer to the created canvas
 */
// llgo:link (*ObjT).CanvasCreate C.lv_canvas_create
func (recv_ *ObjT) CanvasCreate() *ObjT {
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
// llgo:link (*ObjT).CanvasSetBuffer C.lv_canvas_set_buffer
func (recv_ *ObjT) CanvasSetBuffer(buf c.Pointer, w c.Int32T, h c.Int32T, cf ColorFormatT) {
}

/**
 * Set a draw buffer for the canvas. A draw buffer either can be allocated by `lv_draw_buf_create()`
 * or defined statically by `LV_DRAW_BUF_DEFINE_STATIC`. When buffer start address and stride has alignment
 * requirement, it's recommended to use `lv_draw_buf_create`.
 * @param obj       pointer to a canvas object
 * @param draw_buf  pointer to a draw buffer
 */
// llgo:link (*ObjT).CanvasSetDrawBuf C.lv_canvas_set_draw_buf
func (recv_ *ObjT) CanvasSetDrawBuf(draw_buf *DrawBufT) {
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
// llgo:link (*ObjT).CanvasSetPx C.lv_canvas_set_px
func (recv_ *ObjT) CanvasSetPx(x c.Int32T, y c.Int32T, color ColorT, opa OpaT) {
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
// llgo:link (*ObjT).CanvasSetPalette C.lv_canvas_set_palette
func (recv_ *ObjT) CanvasSetPalette(index c.Uint8T, color Color32T) {
}

/*=====================
 * Getter functions
 *====================*/
// llgo:link (*ObjT).CanvasGetDrawBuf C.lv_canvas_get_draw_buf
func (recv_ *ObjT) CanvasGetDrawBuf() *DrawBufT {
	return nil
}

/**
 * Get a pixel's color and opacity
 * @param obj   pointer to a canvas
 * @param x     X coordinate of the pixel
 * @param y     Y coordinate of the pixel
 * @return      ARGB8888 color of the pixel
 */
// llgo:link (*ObjT).CanvasGetPx C.lv_canvas_get_px
func (recv_ *ObjT) CanvasGetPx(x c.Int32T, y c.Int32T) Color32T {
	return Color32T{}
}

/**
 * Get the image of the canvas as a pointer to an `lv_image_dsc_t` variable.
 * @param canvas    pointer to a canvas object
 * @return          pointer to the image descriptor.
 */
// llgo:link (*ObjT).CanvasGetImage C.lv_canvas_get_image
func (recv_ *ObjT) CanvasGetImage() *ImageDscT {
	return nil
}

/**
 * Return the pointer for the buffer.
 * It's recommended to use this function instead of the buffer form the
 * return value of lv_canvas_get_image() as is can be aligned
 * @param canvas    pointer to a canvas object
 * @return          pointer to the buffer
 */
// llgo:link (*ObjT).CanvasGetBuf C.lv_canvas_get_buf
func (recv_ *ObjT) CanvasGetBuf() c.Pointer {
	return nil
}

/**
 * Copy a buffer to the canvas
 * @param obj           pointer to a canvas object
 * @param canvas_area   the area of the canvas to copy
 * @param dest_buf      pointer to a buffer to store the copied data
 * @param dest_area     the area of the destination buffer to copy to. If omitted NULL, copy to the whole `dest_buf`
 */
// llgo:link (*ObjT).CanvasCopyBuf C.lv_canvas_copy_buf
func (recv_ *ObjT) CanvasCopyBuf(canvas_area *AreaT, dest_buf *DrawBufT, dest_area *AreaT) {
}

/**
 * Fill the canvas with color
 * @param obj       pointer to a canvas
 * @param color     the background color
 * @param opa       the desired opacity
 */
// llgo:link (*ObjT).CanvasFillBg C.lv_canvas_fill_bg
func (recv_ *ObjT) CanvasFillBg(color ColorT, opa OpaT) {
}

/**
 * Initialize a layer to use LVGL's generic draw functions (lv_draw_rect/label/...) on the canvas.
 * Needs to be usd in pair with `lv_canvas_finish_layer`.
 * @param canvas    pointer to a canvas
 * @param layer     pointer to a layer variable to initialize
 */
// llgo:link (*ObjT).CanvasInitLayer C.lv_canvas_init_layer
func (recv_ *ObjT) CanvasInitLayer(layer *LayerT) {
}

/**
 * Wait until all the drawings are finished on layer.
 * Needs to be usd in pair with `lv_canvas_init_layer`.
 * @param canvas    pointer to a canvas
 * @param layer     pointer to a layer to finalize
 */
// llgo:link (*ObjT).CanvasFinishLayer C.lv_canvas_finish_layer
func (recv_ *ObjT) CanvasFinishLayer(layer *LayerT) {
}

/**
 * Just a wrapper to `LV_CANVAS_BUF_SIZE` for bindings.
 */
//go:linkname CanvasBufSize C.lv_canvas_buf_size
func CanvasBufSize(w c.Int32T, h c.Int32T, bpp c.Uint8T, stride c.Uint8T) c.Uint32T

/**
 * Can be used by draw units to handle the decoding and
 * prepare everything for the actual image rendering
 * @param t             pointer to a draw task
 * @param draw_dsc      the draw descriptor of the image
 * @param coords        the absolute coordinates of the image
 * @param draw_core_cb  a callback to perform the actual rendering
 */
// llgo:link (*DrawTaskT).DrawImageNormalHelper C.lv_draw_image_normal_helper
func (recv_ *DrawTaskT) DrawImageNormalHelper(draw_dsc *DrawImageDscT, coords *AreaT, draw_core_cb DrawImageCoreCb) {
}

/**
 * Can be used by draw units for TILED images to handle the decoding and
 * prepare everything for the actual image rendering
 * @param t             pointer to a draw task
 * @param draw_dsc      the draw descriptor of the image
 * @param coords        the absolute coordinates of the image
 * @param draw_core_cb  a callback to perform the actual rendering
 */
// llgo:link (*DrawTaskT).DrawImageTiledHelper C.lv_draw_image_tiled_helper
func (recv_ *DrawTaskT) DrawImageTiledHelper(draw_dsc *DrawImageDscT, coords *AreaT, draw_core_cb DrawImageCoreCb) {
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
// llgo:link (*AreaT).ImageBufGetTransformedArea C.lv_image_buf_get_transformed_area
func (recv_ *AreaT) ImageBufGetTransformedArea(w c.Int32T, h c.Int32T, angle c.Int32T, scale_x c.Uint16T, scale_y c.Uint16T, pivot *PointT) {
}

/**
 * Initialize the image decoder module
 * @param image_cache_size    Image cache size in bytes. 0 to disable cache.
 * @param image_header_count  Number of header cache entries. 0 to disable header cache.
 */
//go:linkname ImageDecoderInit C.lv_image_decoder_init
func ImageDecoderInit(image_cache_size c.Uint32T, image_header_count c.Uint32T)

/**
 * Deinitialize the image decoder module
 */
//go:linkname ImageDecoderDeinit C.lv_image_decoder_deinit
func ImageDecoderDeinit()

// llgo:link (*DrawMaskRectDscT).DrawMaskRectDscInit C.lv_draw_mask_rect_dsc_init
func (recv_ *DrawMaskRectDscT) DrawMaskRectDscInit() {
}

/**
 * Try to get a rectangle mask draw descriptor from a draw task.
 * @param task      draw task
 * @return          the task's draw descriptor or NULL if the task is not of type LV_DRAW_TASK_TYPE_MASK_RECTANGLE
 */
// llgo:link (*DrawTaskT).DrawTaskGetMaskRectDsc C.lv_draw_task_get_mask_rect_dsc
func (recv_ *DrawTaskT) DrawTaskGetMaskRectDsc() *DrawMaskRectDscT {
	return nil
}

/**
 * Create a draw task to mask a rectangle from the buffer
 * @param layer     pointer to a layer
 * @param dsc       pointer to a draw descriptor
 */
// llgo:link (*LayerT).DrawMaskRect C.lv_draw_mask_rect
func (recv_ *LayerT) DrawMaskRect(dsc *DrawMaskRectDscT) {
}

/** This describes a glyph.*/

type FontFmtTxtGlyphDscT struct {
	BitmapIndex c.Uint32T
	AdvW        c.Uint32T
	BoxW        c.Uint8T
	BoxH        c.Uint8T
	OfsX        c.Int8T
	OfsY        c.Int8T
}
type FontFmtTxtCmapTypeT c.Int

const (
	FONT_FMT_TXT_CMAP_FORMAT0_FULL FontFmtTxtCmapTypeT = 0
	FONT_FMT_TXT_CMAP_SPARSE_FULL  FontFmtTxtCmapTypeT = 1
	FONT_FMT_TXT_CMAP_FORMAT0_TINY FontFmtTxtCmapTypeT = 2
	FONT_FMT_TXT_CMAP_SPARSE_TINY  FontFmtTxtCmapTypeT = 3
)

/**
 * Map codepoints to a `glyph_dsc`s
 * Several formats are supported to optimize memory usage
 * See https://github.com/lvgl/lv_font_conv/blob/master/doc/font_spec.md
 */

type FontFmtTxtCmapT struct {
	RangeStart     c.Uint32T
	RangeLength    c.Uint16T
	GlyphIdStart   c.Uint16T
	UnicodeList    *c.Uint16T
	GlyphIdOfsList c.Pointer
	ListLength     c.Uint16T
	Type           FontFmtTxtCmapTypeT
}

/** A simple mapping of kern values from pairs*/

type FontFmtTxtKernPairT struct {
	GlyphIds     c.Pointer
	Values       *c.Int8T
	PairCnt      c.Uint32T
	GlyphIdsSize c.Uint32T
}

/** More complex but more optimal class based kern value storage*/

type FontFmtTxtKernClassesT struct {
	ClassPairValues   *c.Int8T
	LeftClassMapping  *c.Uint8T
	RightClassMapping *c.Uint8T
	LeftClassCnt      c.Uint8T
	RightClassCnt     c.Uint8T
}
type FontFmtTxtBitmapFormatT c.Int

const (
	FONT_FMT_TXT_PLAIN                   FontFmtTxtBitmapFormatT = 0
	FONT_FMT_TXT_COMPRESSED              FontFmtTxtBitmapFormatT = 1
	FONT_FMT_TXT_COMPRESSED_NO_PREFILTER FontFmtTxtBitmapFormatT = 2
)

/** Describe store for additional data for fonts */

type FontFmtTxtDscT struct {
	GlyphBitmap  *c.Uint8T
	GlyphDsc     *FontFmtTxtGlyphDscT
	Cmaps        *FontFmtTxtCmapT
	KernDsc      c.Pointer
	KernScale    c.Uint16T
	CmapNum      c.Uint16T
	Bpp          c.Uint16T
	KernClasses  c.Uint16T
	BitmapFormat c.Uint16T
	Stride       c.Uint8T
}

type BuiltinFontSrcT struct {
	FontP *FontT
	Size  c.Uint32T
}

/**
 * Used as `get_glyph_bitmap` callback in lvgl's native font format if the font is uncompressed.
 * @param g_dsc         the glyph descriptor including which font to use, which supply the glyph_index and format.
 * @param draw_buf      a draw buffer that can be used to store the bitmap of the glyph, it's OK not to use it.
 * @return pointer to an A8 bitmap (not necessarily bitmap_out) or NULL if `unicode_letter` not found
 */
// llgo:link (*FontGlyphDscT).FontGetBitmapFmtTxt C.lv_font_get_bitmap_fmt_txt
func (recv_ *FontGlyphDscT) FontGetBitmapFmtTxt(draw_buf *DrawBufT) c.Pointer {
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
// llgo:link (*FontT).FontGetGlyphDscFmtTxt C.lv_font_get_glyph_dsc_fmt_txt
func (recv_ *FontT) FontGetGlyphDscFmtTxt(dsc_out *FontGlyphDscT, unicode_letter c.Uint32T, unicode_letter_next c.Uint32T) bool {
	return false
}

// llgo:type C
type ThemeApplyCbT func(*ThemeT, *ObjT)

/**
 * Get the theme assigned to the display of the object
 * @param obj       pointer to a theme object
 * @return          the theme of the object's display (can be NULL)
 */
// llgo:link (*ObjT).ThemeGetFromObj C.lv_theme_get_from_obj
func (recv_ *ObjT) ThemeGetFromObj() *ThemeT {
	return nil
}

/**
 * Apply the active theme on an object
 * @param obj pointer to an object
 */
// llgo:link (*ObjT).ThemeApply C.lv_theme_apply
func (recv_ *ObjT) ThemeApply() {
}

/**
 * Set a base theme for a theme.
 * The styles from the base them will be added before the styles of the current theme.
 * Arbitrary long chain of themes can be created by setting base themes.
 * @param new_theme pointer to theme which base should be set
 * @param parent pointer to the base theme
 */
// llgo:link (*ThemeT).ThemeSetParent C.lv_theme_set_parent
func (recv_ *ThemeT) ThemeSetParent(parent *ThemeT) {
}

/**
 * Set an apply callback for a theme.
 * The apply callback is used to add styles to different objects
 * @param theme pointer to theme which callback should be set
 * @param apply_cb pointer to the callback
 */
// llgo:link (*ThemeT).ThemeSetApplyCb C.lv_theme_set_apply_cb
func (recv_ *ThemeT) ThemeSetApplyCb(apply_cb ThemeApplyCbT) {
}

/**
 * Get the small font of the theme
 * @param obj pointer to an object
 * @return pointer to the font
 */
// llgo:link (*ObjT).ThemeGetFontSmall C.lv_theme_get_font_small
func (recv_ *ObjT) ThemeGetFontSmall() *FontT {
	return nil
}

/**
 * Get the normal font of the theme
 * @param obj pointer to an object
 * @return pointer to the font
 */
// llgo:link (*ObjT).ThemeGetFontNormal C.lv_theme_get_font_normal
func (recv_ *ObjT) ThemeGetFontNormal() *FontT {
	return nil
}

/**
 * Get the subtitle font of the theme
 * @param obj pointer to an object
 * @return pointer to the font
 */
// llgo:link (*ObjT).ThemeGetFontLarge C.lv_theme_get_font_large
func (recv_ *ObjT) ThemeGetFontLarge() *FontT {
	return nil
}

/**
 * Get the primary color of the theme
 * @param obj pointer to an object
 * @return the color
 */
// llgo:link (*ObjT).ThemeGetColorPrimary C.lv_theme_get_color_primary
func (recv_ *ObjT) ThemeGetColorPrimary() ColorT {
	return ColorT{}
}

/**
 * Get the secondary color of the theme
 * @param obj pointer to an object
 * @return the color
 */
// llgo:link (*ObjT).ThemeGetColorSecondary C.lv_theme_get_color_secondary
func (recv_ *ObjT) ThemeGetColorSecondary() ColorT {
	return ColorT{}
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
// llgo:link (*DisplayT).ThemeDefaultInit C.lv_theme_default_init
func (recv_ *DisplayT) ThemeDefaultInit(color_primary ColorT, color_secondary ColorT, dark bool, font *FontT) *ThemeT {
	return nil
}

/**
 * Check if default theme is initialized
 * @return true if default theme is initialized, false otherwise
 */
//go:linkname ThemeDefaultIsInited C.lv_theme_default_is_inited
func ThemeDefaultIsInited() bool

/**
 * Get default theme
 * @return a pointer to default theme, or NULL if this is not initialized
 */
//go:linkname ThemeDefaultGet C.lv_theme_default_get
func ThemeDefaultGet() *ThemeT

/**
 * Deinitialize the default theme
 */
//go:linkname ThemeDefaultDeinit C.lv_theme_default_deinit
func ThemeDefaultDeinit()

/**
 * Initialize the theme
 * @param disp pointer to display
 * @param dark_bg
 * @param font pointer to a font to use.
 * @return a pointer to reference this theme later
 */
// llgo:link (*DisplayT).ThemeMonoInit C.lv_theme_mono_init
func (recv_ *DisplayT) ThemeMonoInit(dark_bg bool, font *FontT) *ThemeT {
	return nil
}

/**
* Check if the theme is initialized
* @return true if default theme is initialized, false otherwise
 */
//go:linkname ThemeMonoIsInited C.lv_theme_mono_is_inited
func ThemeMonoIsInited() bool

/**
 * Get mono theme
 * @return a pointer to mono theme, or NULL if this is not initialized
 */
//go:linkname ThemeMonoGet C.lv_theme_mono_get
func ThemeMonoGet() *ThemeT

/**
 * Deinitialize the mono theme
 */
//go:linkname ThemeMonoDeinit C.lv_theme_mono_deinit
func ThemeMonoDeinit()

/**
 * Initialize the theme
 * @param disp pointer to display
 * @return a pointer to reference this theme later
 */
// llgo:link (*DisplayT).ThemeSimpleInit C.lv_theme_simple_init
func (recv_ *DisplayT) ThemeSimpleInit() *ThemeT {
	return nil
}

/**
* Check if the theme is initialized
* @return true if default theme is initialized, false otherwise
 */
//go:linkname ThemeSimpleIsInited C.lv_theme_simple_is_inited
func ThemeSimpleIsInited() bool

/**
 * Get simple theme
 * @return a pointer to simple theme, or NULL if this is not initialized
 */
//go:linkname ThemeSimpleGet C.lv_theme_simple_get
func ThemeSimpleGet() *ThemeT

/**
 * Deinitialize the simple theme
 */
//go:linkname ThemeSimpleDeinit C.lv_theme_simple_deinit
func ThemeSimpleDeinit()

/**
 * Redraw the invalidated areas now.
 * Normally the redrawing is periodically executed in `lv_timer_handler` but a long blocking process
 * can prevent the call of `lv_timer_handler`. In this case if the GUI is updated in the process
 * (e.g. progress bar) this function can be called when the screen should be updated.
 * @param disp pointer to display to refresh. NULL to refresh all displays.
 */
// llgo:link (*DisplayT).RefrNow C.lv_refr_now
func (recv_ *DisplayT) RefrNow() {
}

/**
 * Redrawn on object and all its children using the passed draw context
 * @param layer pointer to a layer where to draw.
 * @param obj   the start object from the redraw should start
 */
// llgo:link (*LayerT).ObjRedraw C.lv_obj_redraw
func (recv_ *LayerT) ObjRedraw(obj *ObjT) {
}

/**
 * Called periodically to handle the refreshing
 * @param timer pointer to the timer itself, or `NULL`
 */
// llgo:link (*TimerT).DisplayRefrTimer C.lv_display_refr_timer
func (recv_ *TimerT) DisplayRefrTimer() {
}

/**
 * Initialize the screen refresh subsystem
 */
//go:linkname RefrInit C.lv_refr_init
func RefrInit()

/**
 * Deinitialize the screen refresh subsystem
 */
//go:linkname RefrDeinit C.lv_refr_deinit
func RefrDeinit()

/**
 * Invalidate an area on display to redraw it
 * @param area_p pointer to area which should be invalidated (NULL: delete the invalidated areas)
 * @param disp pointer to display where the area should be invalidated (NULL can be used if there is
 * only one display)
 */
// llgo:link (*DisplayT).InvArea C.lv_inv_area
func (recv_ *DisplayT) InvArea(area_p *AreaT) {
}

/**
 * Get the display which is being refreshed
 * @return the display being refreshed
 */
//go:linkname RefrGetDispRefreshing C.lv_refr_get_disp_refreshing
func RefrGetDispRefreshing() *DisplayT

/**
 * Set the display which is being refreshed
 * @param disp the display being refreshed
 */
// llgo:link (*DisplayT).RefrSetDispRefreshing C.lv_refr_set_disp_refreshing
func (recv_ *DisplayT) RefrSetDispRefreshing() {
}

/**
 * Search the most top object which fully covers an area
 * @param area_p pointer to an area
 * @param obj the first object to start the searching (typically a screen)
 * @return
 */
// llgo:link (*AreaT).RefrGetTopObj C.lv_refr_get_top_obj
func (recv_ *AreaT) RefrGetTopObj(obj *ObjT) *ObjT {
	return nil
}

/**
 * Initialize the object related style manager module.
 * Called by LVGL in `lv_init()`
 */
//go:linkname ObjStyleInit C.lv_obj_style_init
func ObjStyleInit()

/**
 * Deinitialize the object related style manager module.
 * Called by LVGL in `lv_deinit()`
 */
//go:linkname ObjStyleDeinit C.lv_obj_style_deinit
func ObjStyleDeinit()

/**
 * Used internally to create a style transition
 * @param obj
 * @param part
 * @param prev_state
 * @param new_state
 * @param tr
 */
// llgo:link (*ObjT).ObjStyleCreateTransition C.lv_obj_style_create_transition
func (recv_ *ObjT) ObjStyleCreateTransition(part PartT, prev_state StateT, new_state StateT, tr *ObjStyleTransitionDscT) {
}

/**
 * Used internally to compare the appearance of an object in 2 states
 * @param obj
 * @param state1
 * @param state2
 * @return
 */
// llgo:link (*ObjT).ObjStyleStateCompare C.lv_obj_style_state_compare
func (recv_ *ObjT) ObjStyleStateCompare(state1 StateT, state2 StateT) StyleStateCmpT {
	return 0
}

/**
 * Update the layer type of a widget bayed on its current styles.
 * The result will be stored in `obj->spec_attr->layer_type`
 * @param obj       the object whose layer should be updated
 */
// llgo:link (*ObjT).ObjUpdateLayerType C.lv_obj_update_layer_type
func (recv_ *ObjT) ObjUpdateLayerType() {
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
// llgo:link (*ObjT).ObjScrollByRaw C.lv_obj_scroll_by_raw
func (recv_ *ObjT) ObjScrollByRaw(x c.Int32T, y c.Int32T) ResultT {
	return 0
}

/**
 * Get the extended draw area of an object.
 * @param obj       pointer to an object
 * @return          the size extended draw area around the real coordinates
 */
// llgo:link (*ObjT).ObjGetExtDrawSize C.lv_obj_get_ext_draw_size
func (recv_ *ObjT) ObjGetExtDrawSize() c.Int32T {
	return 0
}

// llgo:link (*ObjT).ObjGetLayerType C.lv_obj_get_layer_type
func (recv_ *ObjT) ObjGetLayerType() LayerTypeT {
	return 0
}

/**********************
 * GLOBAL PROTOTYPES
 **********************/
// llgo:link (*ObjT).ObjDestruct C.lv_obj_destruct
func (recv_ *ObjT) ObjDestruct() {
}

/**
 * Init the group module
 * @remarks Internal function, do not call directly.
 */
//go:linkname GroupInit C.lv_group_init
func GroupInit()

/**
 * Deinit the group module
 * @remarks Internal function, do not call directly.
 */
//go:linkname GroupDeinit C.lv_group_deinit
func GroupDeinit()

/**
 * Set the position of an area (width and height will be kept)
 * @param area_p pointer to an area
 * @param x the new x coordinate of the area
 * @param y the new y coordinate of the area
 */
// llgo:link (*AreaT).AreaSetPos C.lv_area_set_pos
func (recv_ *AreaT) AreaSetPos(x c.Int32T, y c.Int32T) {
}

/**
 * Get the common parts of two areas
 * @param res_p pointer to an area, the result will be stored her
 * @param a1_p pointer to the first area
 * @param a2_p pointer to the second area
 * @return false: the two area has NO common parts, res_p is invalid
 */
// llgo:link (*AreaT).AreaIntersect C.lv_area_intersect
func (recv_ *AreaT) AreaIntersect(a1_p *AreaT, a2_p *AreaT) bool {
	return false
}

/**
 * Get resulting sub areas after removing the common parts of two areas from the first area
 * @param res_p pointer to an array of areas with a count of 4, the resulting areas will be stored here
 * @param a1_p pointer to the first area
 * @param a2_p pointer to the second area
 * @return number of results (max 4) or -1 if no intersect
 */
//go:linkname AreaDiff C.lv_area_diff
func AreaDiff(res_p *AreaT, a1_p *AreaT, a2_p *AreaT) c.Int8T

/**
 * Join two areas into a third which involves the other two
 * @param a_res_p pointer to an area, the result will be stored here
 * @param a1_p pointer to the first area
 * @param a2_p pointer to the second area
 */
// llgo:link (*AreaT).AreaJoin C.lv_area_join
func (recv_ *AreaT) AreaJoin(a1_p *AreaT, a2_p *AreaT) {
}

/**
 * Check if a point is on an area
 * @param a_p pointer to an area
 * @param p_p pointer to a point
 * @param radius radius of area (e.g. for rounded rectangle)
 * @return false:the point is out of the area
 */
// llgo:link (*AreaT).AreaIsPointOn C.lv_area_is_point_on
func (recv_ *AreaT) AreaIsPointOn(p_p *PointT, radius c.Int32T) bool {
	return false
}

/**
 * Check if two area has common parts
 * @param a1_p pointer to an area.
 * @param a2_p pointer to another area
 * @return false: a1_p and a2_p has no common parts
 */
// llgo:link (*AreaT).AreaIsOn C.lv_area_is_on
func (recv_ *AreaT) AreaIsOn(a2_p *AreaT) bool {
	return false
}

/**
 * Check if an area is fully on another
 * @param ain_p pointer to an area which could be in 'aholder_p'
 * @param aholder_p pointer to an area which could involve 'ain_p'
 * @param radius radius of `aholder_p` (e.g. for rounded rectangle)
 * @return true: `ain_p` is fully inside `aholder_p`
 */
// llgo:link (*AreaT).AreaIsIn C.lv_area_is_in
func (recv_ *AreaT) AreaIsIn(aholder_p *AreaT, radius c.Int32T) bool {
	return false
}

/**
 * Check if an area is fully out of another
 * @param aout_p pointer to an area which could be in 'aholder_p'
 * @param aholder_p pointer to an area which could involve 'ain_p'
 * @param radius radius of `aholder_p` (e.g. for rounded rectangle)
 * @return true: `aout_p` is fully outside `aholder_p`
 */
// llgo:link (*AreaT).AreaIsOut C.lv_area_is_out
func (recv_ *AreaT) AreaIsOut(aholder_p *AreaT, radius c.Int32T) bool {
	return false
}

/**
 * Check if 2 area is the same
 * @param a pointer to an area
 * @param b pointer to another area
 */
// llgo:link (*AreaT).AreaIsEqual C.lv_area_is_equal
func (recv_ *AreaT) AreaIsEqual(b *AreaT) bool {
	return false
}

/**
 * Initialize the File system interface
 */
//go:linkname FsInit C.lv_fs_init
func FsInit()

/**
 * Deinitialize the File system interface
 */
//go:linkname FsDeinit C.lv_fs_deinit
func FsDeinit()

/**********************
 * GLOBAL PROTOTYPES
 **********************/
// llgo:link (*EventT).EventPush C.lv_event_push
func (recv_ *EventT) EventPush() {
}

// llgo:link (*EventT).EventPop C.lv_event_pop
func (recv_ *EventT) EventPop() {
}

/**
 * Nested events can be called and one of them might belong to an object that is being deleted.
 * Mark this object's `event_temp_data` deleted to know that its `lv_obj_send_event` should return `LV_RESULT_INVALID`
 * @param target     pointer to an event target which was deleted
 */
//go:linkname EventMarkDeleted C.lv_event_mark_deleted
func EventMarkDeleted(target c.Pointer)

type RbColorT c.Int

const (
	RB_COLOR_RED   RbColorT = 0
	RB_COLOR_BLACK RbColorT = 1
)

type RbCompareResT c.Int8T

// llgo:type C
type RbCompareT func(c.Pointer, c.Pointer) RbCompareResT

/**********************
 * GLOBAL PROTOTYPES
 **********************/
// llgo:link (*RbT).RbInit C.lv_rb_init
func (recv_ *RbT) RbInit(compare RbCompareT, node_size c.SizeT) bool {
	return false
}

// llgo:link (*RbT).RbInsert C.lv_rb_insert
func (recv_ *RbT) RbInsert(key c.Pointer) *RbNodeT {
	return nil
}

// llgo:link (*RbT).RbFind C.lv_rb_find
func (recv_ *RbT) RbFind(key c.Pointer) *RbNodeT {
	return nil
}

// llgo:link (*RbT).RbRemoveNode C.lv_rb_remove_node
func (recv_ *RbT) RbRemoveNode(node *RbNodeT) c.Pointer {
	return nil
}

// llgo:link (*RbT).RbRemove C.lv_rb_remove
func (recv_ *RbT) RbRemove(key c.Pointer) c.Pointer {
	return nil
}

// llgo:link (*RbT).RbDropNode C.lv_rb_drop_node
func (recv_ *RbT) RbDropNode(node *RbNodeT) bool {
	return false
}

// llgo:link (*RbT).RbDrop C.lv_rb_drop
func (recv_ *RbT) RbDrop(key c.Pointer) bool {
	return false
}

// llgo:link (*RbT).RbMinimum C.lv_rb_minimum
func (recv_ *RbT) RbMinimum() *RbNodeT {
	return nil
}

// llgo:link (*RbT).RbMaximum C.lv_rb_maximum
func (recv_ *RbT) RbMaximum() *RbNodeT {
	return nil
}

// llgo:link (*RbNodeT).RbMinimumFrom C.lv_rb_minimum_from
func (recv_ *RbNodeT) RbMinimumFrom() *RbNodeT {
	return nil
}

// llgo:link (*RbNodeT).RbMaximumFrom C.lv_rb_maximum_from
func (recv_ *RbNodeT) RbMaximumFrom() *RbNodeT {
	return nil
}

// llgo:link (*RbT).RbDestroy C.lv_rb_destroy
func (recv_ *RbT) RbDestroy() {
}

/**
 * Create an empty message box
 * @param parent        the parent or NULL to create a modal msgbox
 * @return              the created message box
 */
// llgo:link (*ObjT).MsgboxCreate C.lv_msgbox_create
func (recv_ *ObjT) MsgboxCreate() *ObjT {
	return nil
}

/**
 * Add title to the message box. It also creates a header for the title.
 * @param obj           pointer to a message box
 * @param title         the text of the tile
 * @return              the created title label
 */
// llgo:link (*ObjT).MsgboxAddTitle C.lv_msgbox_add_title
func (recv_ *ObjT) MsgboxAddTitle(title *c.Char) *ObjT {
	return nil
}

/**
 * Add a button to the header of to the message box. It also creates a header.
 * @param obj           pointer to a message box
 * @param icon          the icon of the button
 * @return              the created button
 */
// llgo:link (*ObjT).MsgboxAddHeaderButton C.lv_msgbox_add_header_button
func (recv_ *ObjT) MsgboxAddHeaderButton(icon c.Pointer) *ObjT {
	return nil
}

/**
 * Add a text to the content area of message box. Multiple texts will be created below each other.
 * @param obj           pointer to a message box
 * @param text          text to add
 * @return              the created button
 */
// llgo:link (*ObjT).MsgboxAddText C.lv_msgbox_add_text
func (recv_ *ObjT) MsgboxAddText(text *c.Char) *ObjT {
	return nil
}

/**
 * Add a button to the footer of to the message box. It also creates a footer.
 * @param obj           pointer to a message box
 * @param text          the text of the button
 * @return              the created button
 */
// llgo:link (*ObjT).MsgboxAddFooterButton C.lv_msgbox_add_footer_button
func (recv_ *ObjT) MsgboxAddFooterButton(text *c.Char) *ObjT {
	return nil
}

/**
 * Add a close button to the message box. It also creates a header.
 * @param obj           pointer to a message box
 * @return              the created close button
 */
// llgo:link (*ObjT).MsgboxAddCloseButton C.lv_msgbox_add_close_button
func (recv_ *ObjT) MsgboxAddCloseButton() *ObjT {
	return nil
}

/**
 * Get the header widget
 * @param obj           pointer to a message box
 * @return              the header, or NULL if not exists
 */
// llgo:link (*ObjT).MsgboxGetHeader C.lv_msgbox_get_header
func (recv_ *ObjT) MsgboxGetHeader() *ObjT {
	return nil
}

/**
 * Get the footer widget
 * @param obj           pointer to a message box
 * @return              the footer, or NULL if not exists
 */
// llgo:link (*ObjT).MsgboxGetFooter C.lv_msgbox_get_footer
func (recv_ *ObjT) MsgboxGetFooter() *ObjT {
	return nil
}

/**
 * Get the content widget
 * @param obj           pointer to a message box
 * @return              the content
 */
// llgo:link (*ObjT).MsgboxGetContent C.lv_msgbox_get_content
func (recv_ *ObjT) MsgboxGetContent() *ObjT {
	return nil
}

/**
 * Get the title label
 * @param obj           pointer to a message box
 * @return              the title, or NULL if it does not exist
 */
// llgo:link (*ObjT).MsgboxGetTitle C.lv_msgbox_get_title
func (recv_ *ObjT) MsgboxGetTitle() *ObjT {
	return nil
}

/**
 * Close a message box
 * @param mbox           pointer to a message box
 */
// llgo:link (*ObjT).MsgboxClose C.lv_msgbox_close
func (recv_ *ObjT) MsgboxClose() {
}

/**
 * Close a message box in the next call of the message box
 * @param mbox           pointer to a message box
 */
// llgo:link (*ObjT).MsgboxCloseAsync C.lv_msgbox_close_async
func (recv_ *ObjT) MsgboxCloseAsync() {
}

type ButtonmatrixCtrlT c.Int

const (
	BUTTONMATRIX_CTRL_NONE       ButtonmatrixCtrlT = 0
	BUTTONMATRIX_CTRL_WIDTH_1    ButtonmatrixCtrlT = 1
	BUTTONMATRIX_CTRL_WIDTH_2    ButtonmatrixCtrlT = 2
	BUTTONMATRIX_CTRL_WIDTH_3    ButtonmatrixCtrlT = 3
	BUTTONMATRIX_CTRL_WIDTH_4    ButtonmatrixCtrlT = 4
	BUTTONMATRIX_CTRL_WIDTH_5    ButtonmatrixCtrlT = 5
	BUTTONMATRIX_CTRL_WIDTH_6    ButtonmatrixCtrlT = 6
	BUTTONMATRIX_CTRL_WIDTH_7    ButtonmatrixCtrlT = 7
	BUTTONMATRIX_CTRL_WIDTH_8    ButtonmatrixCtrlT = 8
	BUTTONMATRIX_CTRL_WIDTH_9    ButtonmatrixCtrlT = 9
	BUTTONMATRIX_CTRL_WIDTH_10   ButtonmatrixCtrlT = 10
	BUTTONMATRIX_CTRL_WIDTH_11   ButtonmatrixCtrlT = 11
	BUTTONMATRIX_CTRL_WIDTH_12   ButtonmatrixCtrlT = 12
	BUTTONMATRIX_CTRL_WIDTH_13   ButtonmatrixCtrlT = 13
	BUTTONMATRIX_CTRL_WIDTH_14   ButtonmatrixCtrlT = 14
	BUTTONMATRIX_CTRL_WIDTH_15   ButtonmatrixCtrlT = 15
	BUTTONMATRIX_CTRL_HIDDEN     ButtonmatrixCtrlT = 16
	BUTTONMATRIX_CTRL_NO_REPEAT  ButtonmatrixCtrlT = 32
	BUTTONMATRIX_CTRL_DISABLED   ButtonmatrixCtrlT = 64
	BUTTONMATRIX_CTRL_CHECKABLE  ButtonmatrixCtrlT = 128
	BUTTONMATRIX_CTRL_CHECKED    ButtonmatrixCtrlT = 256
	BUTTONMATRIX_CTRL_CLICK_TRIG ButtonmatrixCtrlT = 512
	BUTTONMATRIX_CTRL_POPOVER    ButtonmatrixCtrlT = 1024
	BUTTONMATRIX_CTRL_RECOLOR    ButtonmatrixCtrlT = 2048
	BUTTONMATRIX_CTRL_RESERVED_1 ButtonmatrixCtrlT = 4096
	BUTTONMATRIX_CTRL_RESERVED_2 ButtonmatrixCtrlT = 8192
	BUTTONMATRIX_CTRL_CUSTOM_1   ButtonmatrixCtrlT = 16384
	BUTTONMATRIX_CTRL_CUSTOM_2   ButtonmatrixCtrlT = 32768
)

// llgo:type C
type ButtonmatrixButtonDrawCbT func(*ObjT, c.Uint32T, *AreaT, *AreaT) bool

/**
 * Create a button matrix object
 * @param parent    pointer to an object, it will be the parent of the new button matrix
 * @return          pointer to the created button matrix
 */
// llgo:link (*ObjT).ButtonmatrixCreate C.lv_buttonmatrix_create
func (recv_ *ObjT) ButtonmatrixCreate() *ObjT {
	return nil
}

/**
 * Set a new map. Buttons will be created/deleted according to the map. The
 * button matrix keeps a reference to the map and so the string array must not
 * be deallocated during the life of the matrix.
 * @param obj       pointer to a button matrix object
 * @param map       pointer a string array. The last string has to be: "". Use "\n" to make a line break.
 */
// llgo:link (*ObjT).ButtonmatrixSetMap C.lv_buttonmatrix_set_map
func (recv_ *ObjT) ButtonmatrixSetMap(map_ **c.Char) {
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
// llgo:link (*ObjT).ButtonmatrixSetCtrlMap C.lv_buttonmatrix_set_ctrl_map
func (recv_ *ObjT) ButtonmatrixSetCtrlMap(ctrl_map *ButtonmatrixCtrlT) {
}

/**
 * Set the selected buttons
 * @param obj        pointer to button matrix object
 * @param btn_id     0 based index of the button to modify. (Not counting new lines)
 */
// llgo:link (*ObjT).ButtonmatrixSetSelectedButton C.lv_buttonmatrix_set_selected_button
func (recv_ *ObjT) ButtonmatrixSetSelectedButton(btn_id c.Uint32T) {
}

/**
 * Set the attributes of a button of the button matrix
 * @param obj       pointer to button matrix object
 * @param btn_id    0 based index of the button to modify. (Not counting new lines)
 * @param ctrl      OR-ed attributes. E.g. `LV_BUTTONMATRIX_CTRL_NO_REPEAT | LV_BUTTONMATRIX_CTRL_CHECKABLE`
 */
// llgo:link (*ObjT).ButtonmatrixSetButtonCtrl C.lv_buttonmatrix_set_button_ctrl
func (recv_ *ObjT) ButtonmatrixSetButtonCtrl(btn_id c.Uint32T, ctrl ButtonmatrixCtrlT) {
}

/**
 * Clear the attributes of a button of the button matrix
 * @param obj       pointer to button matrix object
 * @param btn_id    0 based index of the button to modify. (Not counting new lines)
 * @param ctrl      OR-ed attributes. E.g. `LV_BUTTONMATRIX_CTRL_NO_REPEAT | LV_BUTTONMATRIX_CTRL_CHECKABLE`
 */
// llgo:link (*ObjT).ButtonmatrixClearButtonCtrl C.lv_buttonmatrix_clear_button_ctrl
func (recv_ *ObjT) ButtonmatrixClearButtonCtrl(btn_id c.Uint32T, ctrl ButtonmatrixCtrlT) {
}

/**
 * Set attributes of all buttons of a button matrix
 * @param obj       pointer to a button matrix object
 * @param ctrl      attribute(s) to set from `lv_buttonmatrix_ctrl_t`. Values can be ORed.
 */
// llgo:link (*ObjT).ButtonmatrixSetButtonCtrlAll C.lv_buttonmatrix_set_button_ctrl_all
func (recv_ *ObjT) ButtonmatrixSetButtonCtrlAll(ctrl ButtonmatrixCtrlT) {
}

/**
 * Clear the attributes of all buttons of a button matrix
 * @param obj       pointer to a button matrix object
 * @param ctrl      attribute(s) to set from `lv_buttonmatrix_ctrl_t`. Values can be ORed.
 */
// llgo:link (*ObjT).ButtonmatrixClearButtonCtrlAll C.lv_buttonmatrix_clear_button_ctrl_all
func (recv_ *ObjT) ButtonmatrixClearButtonCtrlAll(ctrl ButtonmatrixCtrlT) {
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
// llgo:link (*ObjT).ButtonmatrixSetButtonWidth C.lv_buttonmatrix_set_button_width
func (recv_ *ObjT) ButtonmatrixSetButtonWidth(btn_id c.Uint32T, width c.Uint32T) {
}

/**
 * Make the button matrix like a selector widget (only one button may be checked at a time).
 * `LV_BUTTONMATRIX_CTRL_CHECKABLE` must be enabled on the buttons to be selected using
 * `lv_buttonmatrix_set_ctrl()` or `lv_buttonmatrix_set_button_ctrl_all()`.
 * @param obj       pointer to a button matrix object
 * @param en        whether "one check" mode is enabled
 */
// llgo:link (*ObjT).ButtonmatrixSetOneChecked C.lv_buttonmatrix_set_one_checked
func (recv_ *ObjT) ButtonmatrixSetOneChecked(en bool) {
}

/**
 * Get the current map of a button matrix
 * @param obj       pointer to a button matrix object
 * @return          the current map
 */
// llgo:link (*ObjT).ButtonmatrixGetMap C.lv_buttonmatrix_get_map
func (recv_ *ObjT) ButtonmatrixGetMap() **c.Char {
	return nil
}

/**
 * Get the index of the lastly "activated" button by the user (pressed, released, focused etc)
 * Useful in the `event_cb` to get the text of the button, check if hidden etc.
 * @param obj       pointer to button matrix object
 * @return          index of the last released button (LV_BUTTONMATRIX_BUTTON_NONE: if unset)
 */
// llgo:link (*ObjT).ButtonmatrixGetSelectedButton C.lv_buttonmatrix_get_selected_button
func (recv_ *ObjT) ButtonmatrixGetSelectedButton() c.Uint32T {
	return 0
}

/**
 * Get the button's text
 * @param obj       pointer to button matrix object
 * @param btn_id    the index a button not counting new line characters.
 * @return          text of btn_index` button
 */
// llgo:link (*ObjT).ButtonmatrixGetButtonText C.lv_buttonmatrix_get_button_text
func (recv_ *ObjT) ButtonmatrixGetButtonText(btn_id c.Uint32T) *c.Char {
	return nil
}

/**
 * Get the whether a control value is enabled or disabled for button of a button matrix
 * @param obj       pointer to a button matrix object
 * @param btn_id    the index of a button not counting new line characters.
 * @param ctrl      control values to check (ORed value can be used)
 * @return          true: the control attribute is enabled false: disabled
 */
// llgo:link (*ObjT).ButtonmatrixHasButtonCtrl C.lv_buttonmatrix_has_button_ctrl
func (recv_ *ObjT) ButtonmatrixHasButtonCtrl(btn_id c.Uint32T, ctrl ButtonmatrixCtrlT) bool {
	return false
}

/**
 * Tell whether "one check" mode is enabled or not.
 * @param obj       Button matrix object
 * @return          true: "one check" mode is enabled; false: disabled
 */
// llgo:link (*ObjT).ButtonmatrixGetOneChecked C.lv_buttonmatrix_get_one_checked
func (recv_ *ObjT) ButtonmatrixGetOneChecked() bool {
	return false
}

type LabelLongModeT c.Int

const (
	LABEL_LONG_MODE_WRAP            LabelLongModeT = 0
	LABEL_LONG_MODE_DOTS            LabelLongModeT = 1
	LABEL_LONG_MODE_SCROLL          LabelLongModeT = 2
	LABEL_LONG_MODE_SCROLL_CIRCULAR LabelLongModeT = 3
	LABEL_LONG_MODE_CLIP            LabelLongModeT = 4
)

/**
 * Create a label object
 * @param parent    pointer to an object, it will be the parent of the new label.
 * @return          pointer to the created button
 */
// llgo:link (*ObjT).LabelCreate C.lv_label_create
func (recv_ *ObjT) LabelCreate() *ObjT {
	return nil
}

/**
 * Set a new text for a label. Memory will be allocated to store the text by the label.
 * @param obj           pointer to a label object
 * @param text          '\0' terminated character string. NULL to refresh with the current text.
 */
// llgo:link (*ObjT).LabelSetText C.lv_label_set_text
func (recv_ *ObjT) LabelSetText(text *c.Char) {
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
// llgo:link (*ObjT).LabelSetTextFmt C.lv_label_set_text_fmt
func (recv_ *ObjT) LabelSetTextFmt(fmt *c.Char, __llgo_va_list ...interface{}) {
}

/**
 * Set a static text. It will not be saved by the label so the 'text' variable
 * has to be 'alive' while the label exists.
 * @param obj           pointer to a label object
 * @param text          pointer to a text. NULL to refresh with the current text.
 */
// llgo:link (*ObjT).LabelSetTextStatic C.lv_label_set_text_static
func (recv_ *ObjT) LabelSetTextStatic(text *c.Char) {
}

/**
 * Set the behavior of the label with text longer than the object size
 * @param obj           pointer to a label object
 * @param long_mode     the new mode from 'lv_label_long_mode' enum.
 *                      In LV_LONG_WRAP/DOT/SCROLL/SCROLL_CIRC the size of the label should be set AFTER this function
 */
// llgo:link (*ObjT).LabelSetLongMode C.lv_label_set_long_mode
func (recv_ *ObjT) LabelSetLongMode(long_mode LabelLongModeT) {
}

/**
 * Set where text selection should start
 * @param obj       pointer to a label object
 * @param index     character index from where selection should start. `LV_LABEL_TEXT_SELECTION_OFF` for no selection
 */
// llgo:link (*ObjT).LabelSetTextSelectionStart C.lv_label_set_text_selection_start
func (recv_ *ObjT) LabelSetTextSelectionStart(index c.Uint32T) {
}

/**
 * Set where text selection should end
 * @param obj       pointer to a label object
 * @param index     character index where selection should end. `LV_LABEL_TEXT_SELECTION_OFF` for no selection
 */
// llgo:link (*ObjT).LabelSetTextSelectionEnd C.lv_label_set_text_selection_end
func (recv_ *ObjT) LabelSetTextSelectionEnd(index c.Uint32T) {
}

/**
 * Enable the recoloring by in-line commands
 * @param obj           pointer to a label object
 * @param en            true: enable recoloring, false: disable
 * Example: "This is a #ff0000 red# word"
 */
// llgo:link (*ObjT).LabelSetRecolor C.lv_label_set_recolor
func (recv_ *ObjT) LabelSetRecolor(en bool) {
}

/**
 * Get the text of a label
 * @param obj       pointer to a label object
 * @return          the text of the label
 */
// llgo:link (*ObjT).LabelGetText C.lv_label_get_text
func (recv_ *ObjT) LabelGetText() *c.Char {
	return nil
}

/**
 * Get the long mode of a label
 * @param obj       pointer to a label object
 * @return          the current long mode
 */
// llgo:link (*ObjT).LabelGetLongMode C.lv_label_get_long_mode
func (recv_ *ObjT) LabelGetLongMode() LabelLongModeT {
	return 0
}

/**
 * Get the relative x and y coordinates of a letter
 * @param obj       pointer to a label object
 * @param char_id   index of the character [0 ... text length - 1].
 *                  Expressed in character index, not byte index (different in UTF-8)
 * @param pos       store the result here (E.g. index = 0 gives 0;0 coordinates if the text if aligned to the left)
 */
// llgo:link (*ObjT).LabelGetLetterPos C.lv_label_get_letter_pos
func (recv_ *ObjT) LabelGetLetterPos(char_id c.Uint32T, pos *PointT) {
}

/**
 * Get the index of letter on a relative point of a label.
 * @param obj       pointer to label object
 * @param pos_in    pointer to point with coordinates on a the label
 * @param bidi      whether to use bidi processed
 * @return          The index of the letter on the 'pos_p' point (E.g. on 0;0 is the 0. letter if aligned to the left)
 *                  Expressed in character index and not byte index (different in UTF-8)
 */
// llgo:link (*ObjT).LabelGetLetterOn C.lv_label_get_letter_on
func (recv_ *ObjT) LabelGetLetterOn(pos_in *PointT, bidi bool) c.Uint32T {
	return 0
}

/**
 * Check if a character is drawn under a point.
 * @param obj       pointer to a label object
 * @param pos       Point to check for character under
 * @return          whether a character is drawn under the point
 */
// llgo:link (*ObjT).LabelIsCharUnderPos C.lv_label_is_char_under_pos
func (recv_ *ObjT) LabelIsCharUnderPos(pos *PointT) bool {
	return false
}

/**
 * @brief Get the selection start index.
 * @param obj       pointer to a label object.
 * @return          selection start index. `LV_LABEL_TEXT_SELECTION_OFF` if nothing is selected.
 */
// llgo:link (*ObjT).LabelGetTextSelectionStart C.lv_label_get_text_selection_start
func (recv_ *ObjT) LabelGetTextSelectionStart() c.Uint32T {
	return 0
}

/**
 * @brief Get the selection end index.
 * @param obj       pointer to a label object.
 * @return          selection end index. `LV_LABEL_TXT_SEL_OFF` if nothing is selected.
 */
// llgo:link (*ObjT).LabelGetTextSelectionEnd C.lv_label_get_text_selection_end
func (recv_ *ObjT) LabelGetTextSelectionEnd() c.Uint32T {
	return 0
}

/**
 * @brief Get the recoloring attribute
 * @param obj       pointer to a label object.
 * @return          true: recoloring is enabled, false: recoloring is disabled
 */
// llgo:link (*ObjT).LabelGetRecolor C.lv_label_get_recolor
func (recv_ *ObjT) LabelGetRecolor() bool {
	return false
}

/**
 * Insert a text to a label. The label text cannot be static.
 * @param obj       pointer to a label object
 * @param pos       character index to insert. Expressed in character index and not byte index.
 *                  0: before first char. LV_LABEL_POS_LAST: after last char.
 * @param txt       pointer to the text to insert
 */
// llgo:link (*ObjT).LabelInsText C.lv_label_ins_text
func (recv_ *ObjT) LabelInsText(pos c.Uint32T, txt *c.Char) {
}

/**
 * Delete characters from a label. The label text cannot be static.
 * @param obj       pointer to a label object
 * @param pos       character index from where to cut. Expressed in character index and not byte index.
 *                  0: start in front of the first character
 * @param cnt       number of characters to cut
 */
// llgo:link (*ObjT).LabelCutText C.lv_label_cut_text
func (recv_ *ObjT) LabelCutText(pos c.Uint32T, cnt c.Uint32T) {
}

type BarModeT c.Int

const (
	BAR_MODE_NORMAL      BarModeT = 0
	BAR_MODE_SYMMETRICAL BarModeT = 1
	BAR_MODE_RANGE       BarModeT = 2
)

type BarOrientationT c.Int

const (
	BAR_ORIENTATION_AUTO       BarOrientationT = 0
	BAR_ORIENTATION_HORIZONTAL BarOrientationT = 1
	BAR_ORIENTATION_VERTICAL   BarOrientationT = 2
)

/**
 * Create a bar object
 * @param parent        pointer to an object, it will be the parent of the new bar
 * @return              pointer to the created bar
 */
// llgo:link (*ObjT).BarCreate C.lv_bar_create
func (recv_ *ObjT) BarCreate() *ObjT {
	return nil
}

/**
 * Set a new value on the bar
 * @param obj           pointer to a bar object
 * @param value         new value
 * @param anim          LV_ANIM_ON: set the value with an animation; LV_ANIM_OFF: change the value immediately
 */
// llgo:link (*ObjT).BarSetValue C.lv_bar_set_value
func (recv_ *ObjT) BarSetValue(value c.Int32T, anim AnimEnableT) {
}

/**
 * Set a new start value on the bar
 * @param obj             pointer to a bar object
 * @param start_value     new start value
 * @param anim            LV_ANIM_ON: set the value with an animation; LV_ANIM_OFF: change the value immediately
 */
// llgo:link (*ObjT).BarSetStartValue C.lv_bar_set_start_value
func (recv_ *ObjT) BarSetStartValue(start_value c.Int32T, anim AnimEnableT) {
}

/**
 * Set minimum and the maximum values of a bar
 * @param obj       pointer to the bar object
 * @param min       minimum value
 * @param max       maximum value
 * @note If min is greater than max, the drawing direction becomes to the opposite direction.
 */
// llgo:link (*ObjT).BarSetRange C.lv_bar_set_range
func (recv_ *ObjT) BarSetRange(min c.Int32T, max c.Int32T) {
}

/**
 * Set minimum value of a bar
 * @param obj       pointer to the bar object
 * @param min       minimum value
 */
// llgo:link (*ObjT).BarSetMinValue C.lv_bar_set_min_value
func (recv_ *ObjT) BarSetMinValue(min c.Int32T) {
}

/**
 * Set maximum value of a bar
 * @param obj       pointer to the bar object
 * @param max       maximum value
 */
// llgo:link (*ObjT).BarSetMaxValue C.lv_bar_set_max_value
func (recv_ *ObjT) BarSetMaxValue(max c.Int32T) {
}

/**
 * Set the type of bar.
 * @param obj       pointer to bar object
 * @param mode      bar type from `lv_bar_mode_t`
 */
// llgo:link (*ObjT).BarSetMode C.lv_bar_set_mode
func (recv_ *ObjT) BarSetMode(mode BarModeT) {
}

/**
 * Set the orientation of bar.
 * @param obj           pointer to bar object
 * @param orientation   bar orientation from `lv_bar_orientation_t`
 */
// llgo:link (*ObjT).BarSetOrientation C.lv_bar_set_orientation
func (recv_ *ObjT) BarSetOrientation(orientation BarOrientationT) {
}

/**
 * Get the value of a bar
 * @param obj       pointer to a bar object
 * @return          the value of the bar
 */
// llgo:link (*ObjT).BarGetValue C.lv_bar_get_value
func (recv_ *ObjT) BarGetValue() c.Int32T {
	return 0
}

/**
 * Get the start value of a bar
 * @param obj       pointer to a bar object
 * @return          the start value of the bar
 */
// llgo:link (*ObjT).BarGetStartValue C.lv_bar_get_start_value
func (recv_ *ObjT) BarGetStartValue() c.Int32T {
	return 0
}

/**
 * Get the minimum value of a bar
 * @param obj       pointer to a bar object
 * @return          the minimum value of the bar
 */
// llgo:link (*ObjT).BarGetMinValue C.lv_bar_get_min_value
func (recv_ *ObjT) BarGetMinValue() c.Int32T {
	return 0
}

/**
 * Get the maximum value of a bar
 * @param obj       pointer to a bar object
 * @return          the maximum value of the bar
 */
// llgo:link (*ObjT).BarGetMaxValue C.lv_bar_get_max_value
func (recv_ *ObjT) BarGetMaxValue() c.Int32T {
	return 0
}

/**
 * Get the type of bar.
 * @param obj       pointer to bar object
 * @return          bar type from `lv_bar_mode_t`
 */
// llgo:link (*ObjT).BarGetMode C.lv_bar_get_mode
func (recv_ *ObjT) BarGetMode() BarModeT {
	return 0
}

/**
 * Get the orientation of bar.
 * @param obj       pointer to bar object
 * @return          bar orientation from `lv_bar_orientation_t`
 */
// llgo:link (*ObjT).BarGetOrientation C.lv_bar_get_orientation
func (recv_ *ObjT) BarGetOrientation() BarOrientationT {
	return 0
}

/**
 * Give the bar is in symmetrical mode or not
 * @param obj       pointer to bar object
 * @return          true: in symmetrical mode false : not in
 */
// llgo:link (*ObjT).BarIsSymmetrical C.lv_bar_is_symmetrical
func (recv_ *ObjT) BarIsSymmetrical() bool {
	return false
}

type SliderModeT c.Int

const (
	SLIDER_MODE_NORMAL      SliderModeT = 0
	SLIDER_MODE_SYMMETRICAL SliderModeT = 1
	SLIDER_MODE_RANGE       SliderModeT = 2
)

type SliderOrientationT c.Int

const (
	SLIDER_ORIENTATION_AUTO       SliderOrientationT = 0
	SLIDER_ORIENTATION_HORIZONTAL SliderOrientationT = 1
	SLIDER_ORIENTATION_VERTICAL   SliderOrientationT = 2
)

/**
 * Create a slider object
 * @param parent    pointer to an object, it will be the parent of the new slider.
 * @return          pointer to the created slider
 */
// llgo:link (*ObjT).SliderCreate C.lv_slider_create
func (recv_ *ObjT) SliderCreate() *ObjT {
	return nil
}

/**
 * Set a new value on the slider
 * @param obj       pointer to a slider object
 * @param value     the new value
 * @param anim      LV_ANIM_ON: set the value with an animation; LV_ANIM_OFF: change the value immediately
 */
// llgo:link (*ObjT).SliderSetValue C.lv_slider_set_value
func (recv_ *ObjT) SliderSetValue(value c.Int32T, anim AnimEnableT) {
}

/**
 * Set a new value for the left knob of a slider
 * @param obj       pointer to a slider object
 * @param value     new value
 * @param anim      LV_ANIM_ON: set the value with an animation; LV_ANIM_OFF: change the value immediately
 */
// llgo:link (*ObjT).SliderSetStartValue C.lv_slider_set_start_value
func (recv_ *ObjT) SliderSetStartValue(value c.Int32T, anim AnimEnableT) {
}

/**
 * Set the minimum and the maximum values of a bar
 * @param obj       pointer to the slider object
 * @param min       minimum value
 * @param max       maximum value
 */
// llgo:link (*ObjT).SliderSetRange C.lv_slider_set_range
func (recv_ *ObjT) SliderSetRange(min c.Int32T, max c.Int32T) {
}

/**
 * Set the minimum values of a bar
 * @param obj       pointer to the slider object
 * @param min       minimum value
 */
// llgo:link (*ObjT).SliderSetMinValue C.lv_slider_set_min_value
func (recv_ *ObjT) SliderSetMinValue(min c.Int32T) {
}

/**
 * Set the maximum values of a bar
 * @param obj       pointer to the slider object
 * @param max       maximum value
 */
// llgo:link (*ObjT).SliderSetMaxValue C.lv_slider_set_max_value
func (recv_ *ObjT) SliderSetMaxValue(max c.Int32T) {
}

/**
 * Set the mode of slider.
 * @param obj       pointer to a slider object
 * @param mode      the mode of the slider. See `lv_slider_mode_t`
 */
// llgo:link (*ObjT).SliderSetMode C.lv_slider_set_mode
func (recv_ *ObjT) SliderSetMode(mode SliderModeT) {
}

/**
 * Set the orientation of slider.
 * @param obj           pointer to a slider object
 * @param orientation   slider  orientation from `lv_slider_orientation_t`
 */
// llgo:link (*ObjT).SliderSetOrientation C.lv_slider_set_orientation
func (recv_ *ObjT) SliderSetOrientation(orientation SliderOrientationT) {
}

/**
 * Get the value of the main knob of a slider
 * @param obj       pointer to a slider object
 * @return          the value of the main knob of the slider
 */
// llgo:link (*ObjT).SliderGetValue C.lv_slider_get_value
func (recv_ *ObjT) SliderGetValue() c.Int32T {
	return 0
}

/**
 * Get the value of the left knob of a slider
 * @param obj       pointer to a slider object
 * @return          the value of the left knob of the slider
 */
// llgo:link (*ObjT).SliderGetLeftValue C.lv_slider_get_left_value
func (recv_ *ObjT) SliderGetLeftValue() c.Int32T {
	return 0
}

/**
 * Get the minimum value of a slider
 * @param obj       pointer to a slider object
 * @return          the minimum value of the slider
 */
// llgo:link (*ObjT).SliderGetMinValue C.lv_slider_get_min_value
func (recv_ *ObjT) SliderGetMinValue() c.Int32T {
	return 0
}

/**
 * Get the maximum value of a slider
 * @param obj       pointer to a slider object
 * @return          the maximum value of the slider
 */
// llgo:link (*ObjT).SliderGetMaxValue C.lv_slider_get_max_value
func (recv_ *ObjT) SliderGetMaxValue() c.Int32T {
	return 0
}

/**
 * Give the slider is being dragged or not
 * @param obj       pointer to a slider object
 * @return          true: drag in progress false: not dragged
 */
// llgo:link (*ObjT).SliderIsDragged C.lv_slider_is_dragged
func (recv_ *ObjT) SliderIsDragged() bool {
	return false
}

/**
 * Get the mode of the slider.
 * @param slider       pointer to a slider object
 * @return          see `lv_slider_mode_t`
 */
// llgo:link (*ObjT).SliderGetMode C.lv_slider_get_mode
func (recv_ *ObjT) SliderGetMode() SliderModeT {
	return 0
}

/**
 * Get the orientation of slider.
 * @param obj       pointer to a slider object
 * @return          slider orientation from `lv_slider_orientation_t`
 */
// llgo:link (*ObjT).SliderGetOrientation C.lv_slider_get_orientation
func (recv_ *ObjT) SliderGetOrientation() SliderOrientationT {
	return 0
}

/**
 * Give the slider is in symmetrical mode or not
 * @param obj       pointer to slider object
 * @return          true: in symmetrical mode false : not in
 */
// llgo:link (*ObjT).SliderIsSymmetrical C.lv_slider_is_symmetrical
func (recv_ *ObjT) SliderIsSymmetrical() bool {
	return false
}

type SwitchOrientationT c.Int

const (
	SWITCH_ORIENTATION_AUTO       SwitchOrientationT = 0
	SWITCH_ORIENTATION_HORIZONTAL SwitchOrientationT = 1
	SWITCH_ORIENTATION_VERTICAL   SwitchOrientationT = 2
)

/**
 * Create a switch object
 * @param parent    pointer to an object, it will be the parent of the new switch
 * @return          pointer to the created switch
 */
// llgo:link (*ObjT).SwitchCreate C.lv_switch_create
func (recv_ *ObjT) SwitchCreate() *ObjT {
	return nil
}

/**
 * Set the orientation of switch.
 * @param obj           pointer to switch object
 * @param orientation   switch orientation from `lv_switch_orientation_t`
 */
// llgo:link (*ObjT).SwitchSetOrientation C.lv_switch_set_orientation
func (recv_ *ObjT) SwitchSetOrientation(orientation SwitchOrientationT) {
}

/**
 * Get the orientation of switch.
 * @param obj       pointer to switch object
 * @return          switch orientation from ::lv_switch_orientation_t
 */
// llgo:link (*ObjT).SwitchGetOrientation C.lv_switch_get_orientation
func (recv_ *ObjT) SwitchGetOrientation() SwitchOrientationT {
	return 0
}

/**
 * Represents a date on the calendar object (platform-agnostic).
 */

type CalendarDateT struct {
	Year  c.Uint16T
	Month c.Uint8T
	Day   c.Uint8T
}

/**
 * Create a calendar widget
 * @param parent    pointer to an object, it will be the parent of the new calendar
 * @return          pointer the created calendar
 */
// llgo:link (*ObjT).CalendarCreate C.lv_calendar_create
func (recv_ *ObjT) CalendarCreate() *ObjT {
	return nil
}

/**
 * Set the today's year, month and day at once
 * @param obj  pointer to a calendar object
 * @param year      today's year
 * @param month     today's month [1..12]
 * @param day       today's day [1..31]
 */
// llgo:link (*ObjT).CalendarSetTodayDate C.lv_calendar_set_today_date
func (recv_ *ObjT) CalendarSetTodayDate(year c.Uint32T, month c.Uint32T, day c.Uint32T) {
}

/**
 * Set the today's year
 * @param obj  pointer to a calendar object
 * @param year      today's year
 */
// llgo:link (*ObjT).CalendarSetTodayYear C.lv_calendar_set_today_year
func (recv_ *ObjT) CalendarSetTodayYear(year c.Uint32T) {
}

/**
 * Set the today's year
 * @param obj  pointer to a calendar object
 * @param month     today's month [1..12]
 */
// llgo:link (*ObjT).CalendarSetTodayMonth C.lv_calendar_set_today_month
func (recv_ *ObjT) CalendarSetTodayMonth(month c.Uint32T) {
}

/**
 * Set the today's year
 * @param obj  pointer to a calendar object
 * @param day       today's day [1..31]
 */
// llgo:link (*ObjT).CalendarSetTodayDay C.lv_calendar_set_today_day
func (recv_ *ObjT) CalendarSetTodayDay(day c.Uint32T) {
}

/**
 * Set the currently shown year and month at once
 * @param obj           pointer to a calendar object
 * @param year          shown year
 * @param month         shown month [1..12]
 */
// llgo:link (*ObjT).CalendarSetMonthShown C.lv_calendar_set_month_shown
func (recv_ *ObjT) CalendarSetMonthShown(year c.Uint32T, month c.Uint32T) {
}

/**
 * Set the currently shown year
 * @param obj           pointer to a calendar object
 * @param year          shown year
 */
// llgo:link (*ObjT).CalendarSetShownYear C.lv_calendar_set_shown_year
func (recv_ *ObjT) CalendarSetShownYear(year c.Uint32T) {
}

/**
 * Set the currently shown month
 * @param obj           pointer to a calendar object
 * @param month         shown month [1..12]
 */
// llgo:link (*ObjT).CalendarSetShownMonth C.lv_calendar_set_shown_month
func (recv_ *ObjT) CalendarSetShownMonth(month c.Uint32T) {
}

/**
 * Set the highlighted dates
 * @param obj           pointer to a calendar object
 * @param highlighted   pointer to an `lv_calendar_date_t` array containing the dates.
 *                      Only the pointer will be saved so this variable can't be local which will be destroyed later.
 * @param date_num number of dates in the array
 */
// llgo:link (*ObjT).CalendarSetHighlightedDates C.lv_calendar_set_highlighted_dates
func (recv_ *ObjT) CalendarSetHighlightedDates(highlighted *CalendarDateT, date_num c.SizeT) {
}

/**
 * Set the name of the days
 * @param obj           pointer to a calendar object
 * @param day_names     pointer to an array with the names.
 *                      E.g. `const char * days[7] = {"Sun", "Mon", ...}`
 *                      Only the pointer will be saved so this variable can't be local which will be destroyed later.
 */
// llgo:link (*ObjT).CalendarSetDayNames C.lv_calendar_set_day_names
func (recv_ *ObjT) CalendarSetDayNames(day_names **c.Char) {
}

/**
 * Get the button matrix object of the calendar.
 * It shows the dates and day names.
 * @param obj       pointer to a calendar object
 * @return          pointer to a the button matrix
 */
// llgo:link (*ObjT).CalendarGetBtnmatrix C.lv_calendar_get_btnmatrix
func (recv_ *ObjT) CalendarGetBtnmatrix() *ObjT {
	return nil
}

/**
 * Get the today's date
 * @param calendar  pointer to a calendar object
 * @return          return pointer to an `lv_calendar_date_t` variable containing the date of today.
 */
// llgo:link (*ObjT).CalendarGetTodayDate C.lv_calendar_get_today_date
func (recv_ *ObjT) CalendarGetTodayDate() *CalendarDateT {
	return nil
}

/**
 * Get the currently showed
 * @param calendar  pointer to a calendar object
 * @return          pointer to an `lv_calendar_date_t` variable containing the date is being shown.
 */
// llgo:link (*ObjT).CalendarGetShowedDate C.lv_calendar_get_showed_date
func (recv_ *ObjT) CalendarGetShowedDate() *CalendarDateT {
	return nil
}

/**
 * Get the highlighted dates
 * @param calendar  pointer to a calendar object
 * @return          pointer to an `lv_calendar_date_t` array containing the dates.
 */
// llgo:link (*ObjT).CalendarGetHighlightedDates C.lv_calendar_get_highlighted_dates
func (recv_ *ObjT) CalendarGetHighlightedDates() *CalendarDateT {
	return nil
}

/**
 * Get the number of the highlighted dates
 * @param calendar  pointer to a calendar object
 * @return          number of highlighted days
 */
// llgo:link (*ObjT).CalendarGetHighlightedDatesNum C.lv_calendar_get_highlighted_dates_num
func (recv_ *ObjT) CalendarGetHighlightedDatesNum() c.SizeT {
	return 0
}

/**
 * Get the currently pressed day
 * @param calendar  pointer to a calendar object
 * @param date      store the pressed date here
 * @return          LV_RESULT_OK: there is a valid pressed date
 *                  LV_RESULT_INVALID: there is no pressed data
 */
// llgo:link (*ObjT).CalendarGetPressedDate C.lv_calendar_get_pressed_date
func (recv_ *ObjT) CalendarGetPressedDate(date *CalendarDateT) ResultT {
	return 0
}

/**
 * Create a calendar header with drop-drowns to select the year and month
 * @param parent    pointer to a calendar object.
 * @return          the created header
 */
// llgo:link (*ObjT).CalendarAddHeaderArrow C.lv_calendar_add_header_arrow
func (recv_ *ObjT) CalendarAddHeaderArrow() *ObjT {
	return nil
}

/**
 * Create a calendar header with drop-drowns to select the year and month
 * @param parent    pointer to a calendar object.
 * @return          the created header
 */
// llgo:link (*ObjT).CalendarAddHeaderDropdown C.lv_calendar_add_header_dropdown
func (recv_ *ObjT) CalendarAddHeaderDropdown() *ObjT {
	return nil
}

/**
 * Sets a custom calendar year list
 * @param parent        pointer to a calendar object
 * @param years_list    pointer to an const char array with the years list, see lv_dropdown set_options for more information.
 *                      E.g. `const char * years = "2023\n2022\n2021\n2020\n2019"
 *                      Only the pointer will be saved so this variable can't be local which will be destroyed later.
 */
// llgo:link (*ObjT).CalendarHeaderDropdownSetYearList C.lv_calendar_header_dropdown_set_year_list
func (recv_ *ObjT) CalendarHeaderDropdownSetYearList(years_list *c.Char) {
}

type ImagebuttonStateT c.Int

const (
	IMAGEBUTTON_STATE_RELEASED         ImagebuttonStateT = 0
	IMAGEBUTTON_STATE_PRESSED          ImagebuttonStateT = 1
	IMAGEBUTTON_STATE_DISABLED         ImagebuttonStateT = 2
	IMAGEBUTTON_STATE_CHECKED_RELEASED ImagebuttonStateT = 3
	IMAGEBUTTON_STATE_CHECKED_PRESSED  ImagebuttonStateT = 4
	IMAGEBUTTON_STATE_CHECKED_DISABLED ImagebuttonStateT = 5
	IMAGEBUTTON_STATE_NUM              ImagebuttonStateT = 6
)

/**
 * Create an image button object
 * @param parent pointer to an object, it will be the parent of the new image button
 * @return pointer to the created image button
 */
// llgo:link (*ObjT).ImagebuttonCreate C.lv_imagebutton_create
func (recv_ *ObjT) ImagebuttonCreate() *ObjT {
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
// llgo:link (*ObjT).ImagebuttonSetSrc C.lv_imagebutton_set_src
func (recv_ *ObjT) ImagebuttonSetSrc(state ImagebuttonStateT, src_left c.Pointer, src_mid c.Pointer, src_right c.Pointer) {
}

/**
 * Use this function instead of `lv_obj_add/remove_state` to set a state manually
 * @param imagebutton   pointer to an image button object
 * @param state         the new state
 */
// llgo:link (*ObjT).ImagebuttonSetState C.lv_imagebutton_set_state
func (recv_ *ObjT) ImagebuttonSetState(state ImagebuttonStateT) {
}

/**
 * Get the left image in a given state
 * @param imagebutton   pointer to an image button object
 * @param state         the state where to get the image (from `lv_button_state_t`) `
 * @return              pointer to the left image source (a C array or path to a file)
 */
// llgo:link (*ObjT).ImagebuttonGetSrcLeft C.lv_imagebutton_get_src_left
func (recv_ *ObjT) ImagebuttonGetSrcLeft(state ImagebuttonStateT) c.Pointer {
	return nil
}

/**
 * Get the middle image in a given state
 * @param imagebutton   pointer to an image button object
 * @param state         the state where to get the image (from `lv_button_state_t`) `
 * @return              pointer to the middle image source (a C array or path to a file)
 */
// llgo:link (*ObjT).ImagebuttonGetSrcMiddle C.lv_imagebutton_get_src_middle
func (recv_ *ObjT) ImagebuttonGetSrcMiddle(state ImagebuttonStateT) c.Pointer {
	return nil
}

/**
 * Get the right image in a given state
 * @param imagebutton   pointer to an image button object
 * @param state         the state where to get the image (from `lv_button_state_t`) `
 * @return              pointer to the left image source (a C array or path to a file)
 */
// llgo:link (*ObjT).ImagebuttonGetSrcRight C.lv_imagebutton_get_src_right
func (recv_ *ObjT) ImagebuttonGetSrcRight(state ImagebuttonStateT) c.Pointer {
	return nil
}

type X_lvPartTextareaIdT c.Int

const PART_TEXTAREA_PLACEHOLDER X_lvPartTextareaIdT = 524288

/**
 * Create a text area object
 * @param parent    pointer to an object, it will be the parent of the new text area
 * @return          pointer to the created text area
 */
// llgo:link (*ObjT).TextareaCreate C.lv_textarea_create
func (recv_ *ObjT) TextareaCreate() *ObjT {
	return nil
}

/**
 * Insert a character to the current cursor position.
 * To add a wide char, e.g. 'Á' use `lv_text_encoded_conv_wc('Á')`
 * @param obj       pointer to a text area object
 * @param c         a character (e.g. 'a')
 */
// llgo:link (*ObjT).TextareaAddChar C.lv_textarea_add_char
func (recv_ *ObjT) TextareaAddChar(c c.Uint32T) {
}

/**
 * Insert a text to the current cursor position
 * @param obj       pointer to a text area object
 * @param txt       a '\0' terminated string to insert
 */
// llgo:link (*ObjT).TextareaAddText C.lv_textarea_add_text
func (recv_ *ObjT) TextareaAddText(txt *c.Char) {
}

/**
 * Delete a the left character from the current cursor position
 * @param obj       pointer to a text area object
 */
// llgo:link (*ObjT).TextareaDeleteChar C.lv_textarea_delete_char
func (recv_ *ObjT) TextareaDeleteChar() {
}

/**
 * Delete the right character from the current cursor position
 * @param obj       pointer to a text area object
 */
// llgo:link (*ObjT).TextareaDeleteCharForward C.lv_textarea_delete_char_forward
func (recv_ *ObjT) TextareaDeleteCharForward() {
}

/**
 * Set the text of a text area
 * @param obj       pointer to a text area object
 * @param txt       pointer to the text
 */
// llgo:link (*ObjT).TextareaSetText C.lv_textarea_set_text
func (recv_ *ObjT) TextareaSetText(txt *c.Char) {
}

/**
 * Set the placeholder text of a text area
 * @param obj       pointer to a text area object
 * @param txt       pointer to the text
 */
// llgo:link (*ObjT).TextareaSetPlaceholderText C.lv_textarea_set_placeholder_text
func (recv_ *ObjT) TextareaSetPlaceholderText(txt *c.Char) {
}

/**
 * Set the cursor position
 * @param obj       pointer to a text area object
 * @param pos       the new cursor position in character index
 *                  < 0 : index from the end of the text
 *                  LV_TEXTAREA_CURSOR_LAST: go after the last character
 */
// llgo:link (*ObjT).TextareaSetCursorPos C.lv_textarea_set_cursor_pos
func (recv_ *ObjT) TextareaSetCursorPos(pos c.Int32T) {
}

/**
 * Enable/Disable the positioning of the cursor by clicking the text on the text area.
 * @param obj       pointer to a text area object
 * @param en        true: enable click positions; false: disable
 */
// llgo:link (*ObjT).TextareaSetCursorClickPos C.lv_textarea_set_cursor_click_pos
func (recv_ *ObjT) TextareaSetCursorClickPos(en bool) {
}

/**
 * Enable/Disable password mode
 * @param obj       pointer to a text area object
 * @param en        true: enable, false: disable
 */
// llgo:link (*ObjT).TextareaSetPasswordMode C.lv_textarea_set_password_mode
func (recv_ *ObjT) TextareaSetPasswordMode(en bool) {
}

/**
 * Set the replacement characters to show in password mode
 * @param obj       pointer to a text area object
 * @param bullet    pointer to the replacement text
 */
// llgo:link (*ObjT).TextareaSetPasswordBullet C.lv_textarea_set_password_bullet
func (recv_ *ObjT) TextareaSetPasswordBullet(bullet *c.Char) {
}

/**
 * Configure the text area to one line or back to normal
 * @param obj       pointer to a text area object
 * @param en        true: one line, false: normal
 */
// llgo:link (*ObjT).TextareaSetOneLine C.lv_textarea_set_one_line
func (recv_ *ObjT) TextareaSetOneLine(en bool) {
}

/**
 * Set a list of characters. Only these characters will be accepted by the text area
 * @param obj       pointer to a text area object
 * @param list      list of characters. Only the pointer is saved. E.g. "+-.,0123456789"
 */
// llgo:link (*ObjT).TextareaSetAcceptedChars C.lv_textarea_set_accepted_chars
func (recv_ *ObjT) TextareaSetAcceptedChars(list *c.Char) {
}

/**
 * Set max length of a Text Area.
 * @param obj       pointer to a text area object
 * @param num       the maximal number of characters can be added (`lv_textarea_set_text` ignores it)
 */
// llgo:link (*ObjT).TextareaSetMaxLength C.lv_textarea_set_max_length
func (recv_ *ObjT) TextareaSetMaxLength(num c.Uint32T) {
}

/**
 * In `LV_EVENT_INSERT` the text which planned to be inserted can be replaced by another text.
 * It can be used to add automatic formatting to the text area.
 * @param obj       pointer to a text area object
 * @param txt       pointer to a new string to insert. If `""` no text will be added.
 *                  The variable must be live after the `event_cb` exists. (Should be `global` or `static`)
 */
// llgo:link (*ObjT).TextareaSetInsertReplace C.lv_textarea_set_insert_replace
func (recv_ *ObjT) TextareaSetInsertReplace(txt *c.Char) {
}

/**
 * Enable/disable selection mode.
 * @param obj       pointer to a text area object
 * @param en        true or false to enable/disable selection mode
 */
// llgo:link (*ObjT).TextareaSetTextSelection C.lv_textarea_set_text_selection
func (recv_ *ObjT) TextareaSetTextSelection(en bool) {
}

/**
 * Set how long show the password before changing it to '*'
 * @param obj       pointer to a text area object
 * @param time      show time in milliseconds. 0: hide immediately.
 */
// llgo:link (*ObjT).TextareaSetPasswordShowTime C.lv_textarea_set_password_show_time
func (recv_ *ObjT) TextareaSetPasswordShowTime(time c.Uint32T) {
}

/**
 * @deprecated Use the normal text_align style property instead
 * Set the label's alignment.
 * It sets where the label is aligned (in one line mode it can be smaller than the text area)
 * and how the lines of the area align in case of multiline text area
 * @param obj       pointer to a text area object
 * @param align     the align mode from ::lv_text_align_t
 */
// llgo:link (*ObjT).TextareaSetAlign C.lv_textarea_set_align
func (recv_ *ObjT) TextareaSetAlign(align TextAlignT) {
}

/**
 * Get the text of a text area. In password mode it gives the real text (not '*'s).
 * @param obj       pointer to a text area object
 * @return          pointer to the text
 */
// llgo:link (*ObjT).TextareaGetText C.lv_textarea_get_text
func (recv_ *ObjT) TextareaGetText() *c.Char {
	return nil
}

/**
 * Get the placeholder text of a text area
 * @param obj       pointer to a text area object
 * @return          pointer to the text
 */
// llgo:link (*ObjT).TextareaGetPlaceholderText C.lv_textarea_get_placeholder_text
func (recv_ *ObjT) TextareaGetPlaceholderText() *c.Char {
	return nil
}

/**
 * Get the label of a text area
 * @param obj       pointer to a text area object
 * @return          pointer to the label object
 */
// llgo:link (*ObjT).TextareaGetLabel C.lv_textarea_get_label
func (recv_ *ObjT) TextareaGetLabel() *ObjT {
	return nil
}

/**
 * Get the current cursor position in character index
 * @param obj       pointer to a text area object
 * @return          the cursor position
 */
// llgo:link (*ObjT).TextareaGetCursorPos C.lv_textarea_get_cursor_pos
func (recv_ *ObjT) TextareaGetCursorPos() c.Uint32T {
	return 0
}

/**
 * Get whether the cursor click positioning is enabled or not.
 * @param obj       pointer to a text area object
 * @return          true: enable click positions; false: disable
 */
// llgo:link (*ObjT).TextareaGetCursorClickPos C.lv_textarea_get_cursor_click_pos
func (recv_ *ObjT) TextareaGetCursorClickPos() bool {
	return false
}

/**
 * Get the password mode attribute
 * @param obj       pointer to a text area object
 * @return          true: password mode is enabled, false: disabled
 */
// llgo:link (*ObjT).TextareaGetPasswordMode C.lv_textarea_get_password_mode
func (recv_ *ObjT) TextareaGetPasswordMode() bool {
	return false
}

/**
 * Get the replacement characters to show in password mode
 * @param obj       pointer to a text area object
 * @return          pointer to the replacement text
 */
// llgo:link (*ObjT).TextareaGetPasswordBullet C.lv_textarea_get_password_bullet
func (recv_ *ObjT) TextareaGetPasswordBullet() *c.Char {
	return nil
}

/**
 * Get the one line configuration attribute
 * @param obj       pointer to a text area object
 * @return          true: one line configuration is enabled, false: disabled
 */
// llgo:link (*ObjT).TextareaGetOneLine C.lv_textarea_get_one_line
func (recv_ *ObjT) TextareaGetOneLine() bool {
	return false
}

/**
 * Get a list of accepted characters.
 * @param obj       pointer to a text area object
 * @return          list of accented characters.
 */
// llgo:link (*ObjT).TextareaGetAcceptedChars C.lv_textarea_get_accepted_chars
func (recv_ *ObjT) TextareaGetAcceptedChars() *c.Char {
	return nil
}

/**
 * Get max length of a Text Area.
 * @param obj       pointer to a text area object
 * @return          the maximal number of characters to be add
 */
// llgo:link (*ObjT).TextareaGetMaxLength C.lv_textarea_get_max_length
func (recv_ *ObjT) TextareaGetMaxLength() c.Uint32T {
	return 0
}

/**
 * Find whether text is selected or not.
 * @param obj       pointer to a text area object
 * @return          whether text is selected or not
 */
// llgo:link (*ObjT).TextareaTextIsSelected C.lv_textarea_text_is_selected
func (recv_ *ObjT) TextareaTextIsSelected() bool {
	return false
}

/**
 * Find whether selection mode is enabled.
 * @param obj       pointer to a text area object
 * @return          true: selection mode is enabled, false: disabled
 */
// llgo:link (*ObjT).TextareaGetTextSelection C.lv_textarea_get_text_selection
func (recv_ *ObjT) TextareaGetTextSelection() bool {
	return false
}

/**
 * Set how long show the password before changing it to '*'
 * @param obj       pointer to a text area object
 * @return          show time in milliseconds. 0: hide immediately.
 */
// llgo:link (*ObjT).TextareaGetPasswordShowTime C.lv_textarea_get_password_show_time
func (recv_ *ObjT) TextareaGetPasswordShowTime() c.Uint32T {
	return 0
}

/**
 * Get a the character from the current cursor position
 * @param obj       pointer to a text area object
 * @return          a the character or 0
 */
// llgo:link (*ObjT).TextareaGetCurrentChar C.lv_textarea_get_current_char
func (recv_ *ObjT) TextareaGetCurrentChar() c.Uint32T {
	return 0
}

/**
 * Clear the selection on the text area.
 * @param obj       pointer to a text area object
 */
// llgo:link (*ObjT).TextareaClearSelection C.lv_textarea_clear_selection
func (recv_ *ObjT) TextareaClearSelection() {
}

/**
 * Move the cursor one character right
 * @param obj       pointer to a text area object
 */
// llgo:link (*ObjT).TextareaCursorRight C.lv_textarea_cursor_right
func (recv_ *ObjT) TextareaCursorRight() {
}

/**
 * Move the cursor one character left
 * @param obj       pointer to a text area object
 */
// llgo:link (*ObjT).TextareaCursorLeft C.lv_textarea_cursor_left
func (recv_ *ObjT) TextareaCursorLeft() {
}

/**
 * Move the cursor one line down
 * @param obj       pointer to a text area object
 */
// llgo:link (*ObjT).TextareaCursorDown C.lv_textarea_cursor_down
func (recv_ *ObjT) TextareaCursorDown() {
}

/**
 * Move the cursor one line up
 * @param obj       pointer to a text area object
 */
// llgo:link (*ObjT).TextareaCursorUp C.lv_textarea_cursor_up
func (recv_ *ObjT) TextareaCursorUp() {
}

type TableCellCtrlT c.Int

const (
	TABLE_CELL_CTRL_NONE        TableCellCtrlT = 0
	TABLE_CELL_CTRL_MERGE_RIGHT TableCellCtrlT = 1
	TABLE_CELL_CTRL_TEXT_CROP   TableCellCtrlT = 2
	TABLE_CELL_CTRL_CUSTOM_1    TableCellCtrlT = 16
	TABLE_CELL_CTRL_CUSTOM_2    TableCellCtrlT = 32
	TABLE_CELL_CTRL_CUSTOM_3    TableCellCtrlT = 64
	TABLE_CELL_CTRL_CUSTOM_4    TableCellCtrlT = 128
)

/**
 * Create a table object
 * @param parent        pointer to an object, it will be the parent of the new table
 * @return              pointer to the created table
 */
// llgo:link (*ObjT).TableCreate C.lv_table_create
func (recv_ *ObjT) TableCreate() *ObjT {
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
// llgo:link (*ObjT).TableSetCellValue C.lv_table_set_cell_value
func (recv_ *ObjT) TableSetCellValue(row c.Uint32T, col c.Uint32T, txt *c.Char) {
}

/**
 * Set the value of a cell.  Memory will be allocated to store the text by the table.
 * @param obj           pointer to a Table object
 * @param row           id of the row [0 .. row_cnt -1]
 * @param col           id of the column [0 .. col_cnt -1]
 * @param fmt           `printf`-like format
 * @note                New roes/columns are added automatically if required
 */
// llgo:link (*ObjT).TableSetCellValueFmt C.lv_table_set_cell_value_fmt
func (recv_ *ObjT) TableSetCellValueFmt(row c.Uint32T, col c.Uint32T, fmt *c.Char, __llgo_va_list ...interface{}) {
}

/**
 * Set the number of rows
 * @param obj           table pointer to a Table object
 * @param row_cnt       number of rows
 */
// llgo:link (*ObjT).TableSetRowCount C.lv_table_set_row_count
func (recv_ *ObjT) TableSetRowCount(row_cnt c.Uint32T) {
}

/**
 * Set the number of columns
 * @param obj       table pointer to a Table object
 * @param col_cnt   number of columns.
 */
// llgo:link (*ObjT).TableSetColumnCount C.lv_table_set_column_count
func (recv_ *ObjT) TableSetColumnCount(col_cnt c.Uint32T) {
}

/**
 * Set the width of a column
 * @param obj       table pointer to a Table object
 * @param col_id    id of the column [0 .. LV_TABLE_COL_MAX -1]
 * @param w         width of the column
 */
// llgo:link (*ObjT).TableSetColumnWidth C.lv_table_set_column_width
func (recv_ *ObjT) TableSetColumnWidth(col_id c.Uint32T, w c.Int32T) {
}

/**
 * Add control bits to the cell.
 * @param obj       pointer to a Table object
 * @param row       id of the row [0 .. row_cnt -1]
 * @param col       id of the column [0 .. col_cnt -1]
 * @param ctrl      OR-ed values from ::lv_table_cell_ctrl_t
 */
// llgo:link (*ObjT).TableSetCellCtrl C.lv_table_set_cell_ctrl
func (recv_ *ObjT) TableSetCellCtrl(row c.Uint32T, col c.Uint32T, ctrl TableCellCtrlT) {
}

/**
 * Clear control bits of the cell.
 * @param obj       pointer to a Table object
 * @param row       id of the row [0 .. row_cnt -1]
 * @param col       id of the column [0 .. col_cnt -1]
 * @param ctrl      OR-ed values from ::lv_table_cell_ctrl_t
 */
// llgo:link (*ObjT).TableClearCellCtrl C.lv_table_clear_cell_ctrl
func (recv_ *ObjT) TableClearCellCtrl(row c.Uint32T, col c.Uint32T, ctrl TableCellCtrlT) {
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
// llgo:link (*ObjT).TableSetCellUserData C.lv_table_set_cell_user_data
func (recv_ *ObjT) TableSetCellUserData(row c.Uint16T, col c.Uint16T, user_data c.Pointer) {
}

/**
 * Set the selected cell
 * @param obj       pointer to a table object
 * @param row       id of the cell row to select
 * @param col       id of the cell column to select
 */
// llgo:link (*ObjT).TableSetSelectedCell C.lv_table_set_selected_cell
func (recv_ *ObjT) TableSetSelectedCell(row c.Uint16T, col c.Uint16T) {
}

/**
 * Get the value of a cell.
 * @param obj       pointer to a Table object
 * @param row       id of the row [0 .. row_cnt -1]
 * @param col       id of the column [0 .. col_cnt -1]
 * @return          text in the cell
 */
// llgo:link (*ObjT).TableGetCellValue C.lv_table_get_cell_value
func (recv_ *ObjT) TableGetCellValue(row c.Uint32T, col c.Uint32T) *c.Char {
	return nil
}

/**
 * Get the number of rows.
 * @param obj       table pointer to a Table object
 * @return          number of rows.
 */
// llgo:link (*ObjT).TableGetRowCount C.lv_table_get_row_count
func (recv_ *ObjT) TableGetRowCount() c.Uint32T {
	return 0
}

/**
 * Get the number of columns.
 * @param obj       table pointer to a Table object
 * @return          number of columns.
 */
// llgo:link (*ObjT).TableGetColumnCount C.lv_table_get_column_count
func (recv_ *ObjT) TableGetColumnCount() c.Uint32T {
	return 0
}

/**
 * Get the width of a column
 * @param obj       table pointer to a Table object
 * @param col       id of the column [0 .. LV_TABLE_COL_MAX -1]
 * @return          width of the column
 */
// llgo:link (*ObjT).TableGetColumnWidth C.lv_table_get_column_width
func (recv_ *ObjT) TableGetColumnWidth(col c.Uint32T) c.Int32T {
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
// llgo:link (*ObjT).TableHasCellCtrl C.lv_table_has_cell_ctrl
func (recv_ *ObjT) TableHasCellCtrl(row c.Uint32T, col c.Uint32T, ctrl TableCellCtrlT) bool {
	return false
}

/**
 * Get the selected cell (pressed and or focused)
 * @param obj       pointer to a table object
 * @param row       pointer to variable to store the selected row (LV_TABLE_CELL_NONE: if no cell selected)
 * @param col       pointer to variable to store the selected column  (LV_TABLE_CELL_NONE: if no cell selected)
 */
// llgo:link (*ObjT).TableGetSelectedCell C.lv_table_get_selected_cell
func (recv_ *ObjT) TableGetSelectedCell(row *c.Uint32T, col *c.Uint32T) {
}

/**
 * Get custom user data to the cell.
 * @param obj       pointer to a Table object
 * @param row       id of the row [0 .. row_cnt -1]
 * @param col       id of the column [0 .. col_cnt -1]
 */
// llgo:link (*ObjT).TableGetCellUserData C.lv_table_get_cell_user_data
func (recv_ *ObjT) TableGetCellUserData(row c.Uint16T, col c.Uint16T) c.Pointer {
	return nil
}

/**
 * Create a check box object
 * @param parent    pointer to an object, it will be the parent of the new button
 * @return          pointer to the created check box
 */
// llgo:link (*ObjT).CheckboxCreate C.lv_checkbox_create
func (recv_ *ObjT) CheckboxCreate() *ObjT {
	return nil
}

/**
 * Set the text of a check box. `txt` will be copied and may be deallocated
 * after this function returns.
 * @param obj   pointer to a check box
 * @param txt   the text of the check box. NULL to refresh with the current text.
 */
// llgo:link (*ObjT).CheckboxSetText C.lv_checkbox_set_text
func (recv_ *ObjT) CheckboxSetText(txt *c.Char) {
}

/**
 * Set the text of a check box. `txt` must not be deallocated during the life
 * of this checkbox.
 * @param obj   pointer to a check box
 * @param txt   the text of the check box.
 */
// llgo:link (*ObjT).CheckboxSetTextStatic C.lv_checkbox_set_text_static
func (recv_ *ObjT) CheckboxSetTextStatic(txt *c.Char) {
}

/**
 * Get the text of a check box
 * @param obj   pointer to check box object
 * @return      pointer to the text of the check box
 */
// llgo:link (*ObjT).CheckboxGetText C.lv_checkbox_get_text
func (recv_ *ObjT) CheckboxGetText() *c.Char {
	return nil
}

type RollerModeT c.Int

const (
	ROLLER_MODE_NORMAL   RollerModeT = 0
	ROLLER_MODE_INFINITE RollerModeT = 1
)

/**
 * Create a roller object
 * @param parent    pointer to an object, it will be the parent of the new roller.
 * @return          pointer to the created roller
 */
// llgo:link (*ObjT).RollerCreate C.lv_roller_create
func (recv_ *ObjT) RollerCreate() *ObjT {
	return nil
}

/**
 * Set the options on a roller
 * @param obj       pointer to roller object
 * @param options   a string with '\n' separated options. E.g. "One\nTwo\nThree"
 * @param mode      `LV_ROLLER_MODE_NORMAL` or `LV_ROLLER_MODE_INFINITE`
 */
// llgo:link (*ObjT).RollerSetOptions C.lv_roller_set_options
func (recv_ *ObjT) RollerSetOptions(options *c.Char, mode RollerModeT) {
}

/**
 * Set the selected option
 * @param obj       pointer to a roller object
 * @param sel_opt   index of the selected option (0 ... number of option - 1);
 * @param anim   LV_ANIM_ON: set with animation; LV_ANIM_OFF set immediately
 */
// llgo:link (*ObjT).RollerSetSelected C.lv_roller_set_selected
func (recv_ *ObjT) RollerSetSelected(sel_opt c.Uint32T, anim AnimEnableT) {
}

/**
 * Sets the given string as the selection on the roller. Does not alter the current selection on failure.
 * @param obj               pointer to roller object
 * @param sel_opt   pointer to the string you want to set as an option
 * @param anim          LV_ANIM_ON: set with animation; LV_ANIM_OFF set immediately
 * @return                  `true` if set successfully and `false` if the given string does not exist as an option in the roller
 */
// llgo:link (*ObjT).RollerSetSelectedStr C.lv_roller_set_selected_str
func (recv_ *ObjT) RollerSetSelectedStr(sel_opt *c.Char, anim AnimEnableT) bool {
	return false
}

/**
 * Set the height to show the given number of rows (options)
 * @param obj       pointer to a roller object
 * @param row_cnt   number of desired visible rows
 */
// llgo:link (*ObjT).RollerSetVisibleRowCount C.lv_roller_set_visible_row_count
func (recv_ *ObjT) RollerSetVisibleRowCount(row_cnt c.Uint32T) {
}

/**
 * Get the index of the selected option
 * @param obj       pointer to a roller object
 * @return          index of the selected option (0 ... number of option - 1);
 */
// llgo:link (*ObjT).RollerGetSelected C.lv_roller_get_selected
func (recv_ *ObjT) RollerGetSelected() c.Uint32T {
	return 0
}

/**
 * Get the current selected option as a string.
 * @param obj       pointer to roller object
 * @param buf       pointer to an array to store the string
 * @param buf_size  size of `buf` in bytes. 0: to ignore it.
 */
// llgo:link (*ObjT).RollerGetSelectedStr C.lv_roller_get_selected_str
func (recv_ *ObjT) RollerGetSelectedStr(buf *c.Char, buf_size c.Uint32T) {
}

/**
 * Get the options of a roller
 * @param obj       pointer to roller object
 * @return          the options separated by '\n'-s (E.g. "Option1\nOption2\nOption3")
 */
// llgo:link (*ObjT).RollerGetOptions C.lv_roller_get_options
func (recv_ *ObjT) RollerGetOptions() *c.Char {
	return nil
}

/**
 * Get the total number of options
 * @param obj   pointer to a roller object
 * @return      the total number of options
 */
// llgo:link (*ObjT).RollerGetOptionCount C.lv_roller_get_option_count
func (recv_ *ObjT) RollerGetOptionCount() c.Uint32T {
	return 0
}

/**
 * Create a window widget
 * @param parent    pointer to a parent widget
 * @return          the created window
 */
// llgo:link (*ObjT).WinCreate C.lv_win_create
func (recv_ *ObjT) WinCreate() *ObjT {
	return nil
}

/**
 * Add a title to the window
 * @param obj       pointer to a window widget
 * @param txt       the text of the title
 * @return          the widget where the content of the title can be created
 */
// llgo:link (*ObjT).WinAddTitle C.lv_win_add_title
func (recv_ *ObjT) WinAddTitle(txt *c.Char) *ObjT {
	return nil
}

/**
 * Add a button to the window
 * @param obj       pointer to a window widget
 * @param icon      an icon to be displayed on the button
 * @param btn_w     width of the button
 * @return          the widget where the content of the button can be created
 */
// llgo:link (*ObjT).WinAddButton C.lv_win_add_button
func (recv_ *ObjT) WinAddButton(icon c.Pointer, btn_w c.Int32T) *ObjT {
	return nil
}

/**
 * Get the header of the window
 * @param win       pointer to a window widget
 * @return          the header of the window
 */
// llgo:link (*ObjT).WinGetHeader C.lv_win_get_header
func (recv_ *ObjT) WinGetHeader() *ObjT {
	return nil
}

/**
 * Get the content of the window
 * @param win       pointer to a window widget
 * @return          the content of the window
 */
// llgo:link (*ObjT).WinGetContent C.lv_win_get_content
func (recv_ *ObjT) WinGetContent() *ObjT {
	return nil
}

type KeyboardModeT c.Int

const (
	KEYBOARD_MODE_TEXT_LOWER KeyboardModeT = 0
	KEYBOARD_MODE_TEXT_UPPER KeyboardModeT = 1
	KEYBOARD_MODE_SPECIAL    KeyboardModeT = 2
	KEYBOARD_MODE_NUMBER     KeyboardModeT = 3
	KEYBOARD_MODE_USER_1     KeyboardModeT = 4
	KEYBOARD_MODE_USER_2     KeyboardModeT = 5
	KEYBOARD_MODE_USER_3     KeyboardModeT = 6
	KEYBOARD_MODE_USER_4     KeyboardModeT = 7
)

/**
 * Create a Keyboard object
 * @param parent    pointer to an object, it will be the parent of the new keyboard
 * @return          pointer to the created keyboard
 */
// llgo:link (*ObjT).KeyboardCreate C.lv_keyboard_create
func (recv_ *ObjT) KeyboardCreate() *ObjT {
	return nil
}

/**
 * Assign a Text Area to the Keyboard. The pressed characters will be put there.
 * @param kb        pointer to a Keyboard object
 * @param ta        pointer to a Text Area object to write there
 */
// llgo:link (*ObjT).KeyboardSetTextarea C.lv_keyboard_set_textarea
func (recv_ *ObjT) KeyboardSetTextarea(ta *ObjT) {
}

/**
 * Set a new a mode (text or number map)
 * @param kb        pointer to a Keyboard object
 * @param mode      the mode from 'lv_keyboard_mode_t'
 */
// llgo:link (*ObjT).KeyboardSetMode C.lv_keyboard_set_mode
func (recv_ *ObjT) KeyboardSetMode(mode KeyboardModeT) {
}

/**
 * Show the button title in a popover when pressed.
 * @param kb        pointer to a Keyboard object
 * @param en        whether "popovers" mode is enabled
 */
// llgo:link (*ObjT).KeyboardSetPopovers C.lv_keyboard_set_popovers
func (recv_ *ObjT) KeyboardSetPopovers(en bool) {
}

/**
 * Set a new map for the keyboard
 * @param kb        pointer to a Keyboard object
 * @param mode      keyboard map to alter 'lv_keyboard_mode_t'
 * @param map       pointer to a string array to describe the map.
 *                  See 'lv_buttonmatrix_set_map()' for more info.
 * @param ctrl_map  See 'lv_buttonmatrix_set_ctrl_map()' for more info.

 */
// llgo:link (*ObjT).KeyboardSetMap C.lv_keyboard_set_map
func (recv_ *ObjT) KeyboardSetMap(mode KeyboardModeT, map_ **c.Char, ctrl_map *ButtonmatrixCtrlT) {
}

/**
 * Assign a Text Area to the Keyboard. The pressed characters will be put there.
 * @param kb        pointer to a Keyboard object
 * @return          pointer to the assigned Text Area object
 */
// llgo:link (*ObjT).KeyboardGetTextarea C.lv_keyboard_get_textarea
func (recv_ *ObjT) KeyboardGetTextarea() *ObjT {
	return nil
}

/**
 * Set a new a mode (text or number map)
 * @param kb        pointer to a Keyboard object
 * @return          the current mode from 'lv_keyboard_mode_t'
 */
// llgo:link (*ObjT).KeyboardGetMode C.lv_keyboard_get_mode
func (recv_ *ObjT) KeyboardGetMode() KeyboardModeT {
	return 0
}

/**
 * Tell whether "popovers" mode is enabled or not.
 * @param obj       pointer to a Keyboard object
 * @return          true: "popovers" mode is enabled; false: disabled
 */
// llgo:link (*ObjT).KeyboardGetPopovers C.lv_keyboard_get_popovers
func (recv_ *ObjT) KeyboardGetPopovers() bool {
	return false
}

/**
 * Get the current map of a keyboard
 * @param kb        pointer to a keyboard object
 * @return          the current map
 */
// llgo:link (*ObjT).KeyboardGetMapArray C.lv_keyboard_get_map_array
func (recv_ *ObjT) KeyboardGetMapArray() **c.Char {
	return nil
}

/**
 * Get the index of the lastly "activated" button by the user (pressed, released, focused etc)
 * Useful in the `event_cb` to get the text of the button, check if hidden etc.
 * @param obj       pointer to button matrix object
 * @return          index of the last released button (LV_BUTTONMATRIX_BUTTON_NONE: if unset)
 */
// llgo:link (*ObjT).KeyboardGetSelectedButton C.lv_keyboard_get_selected_button
func (recv_ *ObjT) KeyboardGetSelectedButton() c.Uint32T {
	return 0
}

/**
 * Get the button's text
 * @param obj       pointer to button matrix object
 * @param btn_id    the index a button not counting new line characters.
 * @return          text of btn_index` button
 */
// llgo:link (*ObjT).KeyboardGetButtonText C.lv_keyboard_get_button_text
func (recv_ *ObjT) KeyboardGetButtonText(btn_id c.Uint32T) *c.Char {
	return nil
}

/**
 * Default keyboard event to add characters to the Text area and change the map.
 * If a custom `event_cb` is added to the keyboard this function can be called from it to handle the
 * button clicks
 * @param e the triggering event
 */
// llgo:link (*EventT).KeyboardDefEventCb C.lv_keyboard_def_event_cb
func (recv_ *EventT) KeyboardDefEventCb() {
}

/**
 * Create a line object
 * @param parent pointer to an object, it will be the parent of the new line
 * @return pointer to the created line
 */
// llgo:link (*ObjT).LineCreate C.lv_line_create
func (recv_ *ObjT) LineCreate() *ObjT {
	return nil
}

/**
 * Set an array of points. The line object will connect these points.
 * @param obj           pointer to a line object
 * @param points        an array of points. Only the address is saved, so the array needs to be alive while the line exists
 * @param point_num     number of points in 'point_a'
 */
// llgo:link (*ObjT).LineSetPoints C.lv_line_set_points
func (recv_ *ObjT) LineSetPoints(points *PointPreciseT, point_num c.Uint32T) {
}

/**
 * Set a non-const array of points. Identical to `lv_line_set_points` except the array may be retrieved by `lv_line_get_points_mutable`.
 * @param obj           pointer to a line object
 * @param points        a non-const array of points. Only the address is saved, so the array needs to be alive while the line exists.
 * @param point_num     number of points in 'point_a'
 */
// llgo:link (*ObjT).LineSetPointsMutable C.lv_line_set_points_mutable
func (recv_ *ObjT) LineSetPointsMutable(points *PointPreciseT, point_num c.Uint32T) {
}

/**
 * Enable (or disable) the y coordinate inversion.
 * If enabled then y will be subtracted from the height of the object,
 * therefore the y = 0 coordinate will be on the bottom.
 * @param obj       pointer to a line object
 * @param en        true: enable the y inversion, false:disable the y inversion
 */
// llgo:link (*ObjT).LineSetYInvert C.lv_line_set_y_invert
func (recv_ *ObjT) LineSetYInvert(en bool) {
}

/**
 * Get the pointer to the array of points.
 * @param obj           pointer to a line object
 * @return              const pointer to the array of points
 */
// llgo:link (*ObjT).LineGetPoints C.lv_line_get_points
func (recv_ *ObjT) LineGetPoints() *PointPreciseT {
	return nil
}

/**
 * Get the number of points in the array of points.
 * @param obj           pointer to a line object
 * @return              number of points in array of points
 */
// llgo:link (*ObjT).LineGetPointCount C.lv_line_get_point_count
func (recv_ *ObjT) LineGetPointCount() c.Uint32T {
	return 0
}

/**
 * Check the mutability of the stored point array pointer.
 * @param obj           pointer to a line object
 * @return              true: the point array pointer is mutable, false: constant
 */
// llgo:link (*ObjT).LineIsPointArrayMutable C.lv_line_is_point_array_mutable
func (recv_ *ObjT) LineIsPointArrayMutable() bool {
	return false
}

/**
 * Get a pointer to the mutable array of points or NULL if it is not mutable
 * @param obj           pointer to a line object
 * @return              pointer to the array of points. NULL if not mutable.
 */
// llgo:link (*ObjT).LineGetPointsMutable C.lv_line_get_points_mutable
func (recv_ *ObjT) LineGetPointsMutable() *PointPreciseT {
	return nil
}

/**
 * Get the y inversion attribute
 * @param obj       pointer to a line object
 * @return          true: y inversion is enabled, false: disabled
 */
// llgo:link (*ObjT).LineGetYInvert C.lv_line_get_y_invert
func (recv_ *ObjT) LineGetYInvert() bool {
	return false
}

type AnimimgPartT c.Int

const ANIM_IMAGE_PART_MAIN AnimimgPartT = 0

/**
 * Create an animation image objects
 * @param parent pointer to an object, it will be the parent of the new button
 * @return pointer to the created animation image object
 */
// llgo:link (*ObjT).AnimimgCreate C.lv_animimg_create
func (recv_ *ObjT) AnimimgCreate() *ObjT {
	return nil
}

/**
 * Set the image animation images source.
 * @param obj       pointer to an animation image object
 * @param dsc       pointer to a series images
 * @param num       images' number
 */
// llgo:link (*ObjT).AnimimgSetSrc C.lv_animimg_set_src
func (recv_ *ObjT) AnimimgSetSrc(dsc *c.Pointer, num c.SizeT) {
}

/**
 *  Set the images source for flip playback of animation image.
 * @param obj       pointer to an animation image object
 * @param dsc       pointer to a series images
 * @param num       images' number
 */
// llgo:link (*ObjT).AnimimgSetSrcReverse C.lv_animimg_set_src_reverse
func (recv_ *ObjT) AnimimgSetSrcReverse(dsc *c.Pointer, num c.SizeT) {
}

/**
 * Startup the image animation.
 * @param obj   pointer to an animation image object
 */
// llgo:link (*ObjT).AnimimgStart C.lv_animimg_start
func (recv_ *ObjT) AnimimgStart() {
}

/**
 * Delete the image animation.
 * @param obj   pointer to an animation image object
 */
// llgo:link (*ObjT).AnimimgDelete C.lv_animimg_delete
func (recv_ *ObjT) AnimimgDelete() bool {
	return false
}

/**
 * Set the image animation duration time. unit:ms
 * @param obj       pointer to an animation image object
 * @param duration  the duration in milliseconds
 */
// llgo:link (*ObjT).AnimimgSetDuration C.lv_animimg_set_duration
func (recv_ *ObjT) AnimimgSetDuration(duration c.Uint32T) {
}

/**
 * Set the image animation repeatedly play times.
 * @param obj       pointer to an animation image object
 * @param count     the number of times to repeat the animation
 */
// llgo:link (*ObjT).AnimimgSetRepeatCount C.lv_animimg_set_repeat_count
func (recv_ *ObjT) AnimimgSetRepeatCount(count c.Uint32T) {
}

/**
 * Make the image animation to play back to when the forward direction is ready.
 * @param obj   pointer to an animation image object
 * @param duration   the duration of the playback image animation in milliseconds. 0: disable playback
 */
// llgo:link (*ObjT).AnimimgSetReverseDuration C.lv_animimg_set_reverse_duration
func (recv_ *ObjT) AnimimgSetReverseDuration(duration c.Uint32T) {
}

/**
 * Make the image animation to play back to when the forward direction is ready.
 * @param obj   pointer to an animation image object
 * @param duration   delay in milliseconds before starting the playback image animation.
 */
// llgo:link (*ObjT).AnimimgSetReverseDelay C.lv_animimg_set_reverse_delay
func (recv_ *ObjT) AnimimgSetReverseDelay(duration c.Uint32T) {
}

/**
 * Set a function call when the animation image really starts (considering `delay`)
 * @param obj   pointer to an animation image object
 * @param start_cb   a function call when the animation is start
 */
// llgo:link (*ObjT).AnimimgSetStartCb C.lv_animimg_set_start_cb
func (recv_ *ObjT) AnimimgSetStartCb(start_cb AnimStartCbT) {
}

/**
 * Set a function call when the animation is completed
 * @param obj pointer to an animation image object
 * @param completed_cb  a function call when the animation is completed
 */
// llgo:link (*ObjT).AnimimgSetCompletedCb C.lv_animimg_set_completed_cb
func (recv_ *ObjT) AnimimgSetCompletedCb(completed_cb AnimCompletedCbT) {
}

/**
 * Get the image animation images source.
 * @param obj   pointer to an animation image object
 * @return a     pointer that will point to a series images
 */
// llgo:link (*ObjT).AnimimgGetSrc C.lv_animimg_get_src
func (recv_ *ObjT) AnimimgGetSrc() *c.Pointer {
	return nil
}

/**
 * Get the image animation images source.
 * @param obj   pointer to an animation image object
 * @return      the number of source images
 */
// llgo:link (*ObjT).AnimimgGetSrcCount C.lv_animimg_get_src_count
func (recv_ *ObjT) AnimimgGetSrcCount() c.Uint8T {
	return 0
}

/**
 * Get the image animation duration time. unit:ms
 * @param obj   pointer to an animation image object
 * @return      the animation duration time
 */
// llgo:link (*ObjT).AnimimgGetDuration C.lv_animimg_get_duration
func (recv_ *ObjT) AnimimgGetDuration() c.Uint32T {
	return 0
}

/**
 * Get the image animation repeat play times.
 * @param obj   pointer to an animation image object
 * @return      the repeat count
 */
// llgo:link (*ObjT).AnimimgGetRepeatCount C.lv_animimg_get_repeat_count
func (recv_ *ObjT) AnimimgGetRepeatCount() c.Uint32T {
	return 0
}

/**
 * Get the image animation underlying animation.
 * @param obj   pointer to an animation image object
 * @return      the animation reference
 */
// llgo:link (*ObjT).AnimimgGetAnim C.lv_animimg_get_anim
func (recv_ *ObjT) AnimimgGetAnim() *AnimT {
	return nil
}

/**
 * Create a drop-down list object
 * @param parent pointer to an object, it will be the parent of the new drop-down list
 * @return pointer to the created drop-down list
 */
// llgo:link (*ObjT).DropdownCreate C.lv_dropdown_create
func (recv_ *ObjT) DropdownCreate() *ObjT {
	return nil
}

/**
 * Set text of the drop-down list's button.
 * If set to `NULL` the selected option's text will be displayed on the button.
 * If set to a specific text then that text will be shown regardless of the selected option.
 * @param obj       pointer to a drop-down list object
 * @param txt       the text as a string (Only its pointer is saved)
 */
// llgo:link (*ObjT).DropdownSetText C.lv_dropdown_set_text
func (recv_ *ObjT) DropdownSetText(txt *c.Char) {
}

/**
 * Set the options in a drop-down list from a string.
 * The options will be copied and saved in the object so the `options` can be destroyed after calling this function
 * @param obj       pointer to drop-down list object
 * @param options   a string with '\n' separated options. E.g. "One\nTwo\nThree"
 */
// llgo:link (*ObjT).DropdownSetOptions C.lv_dropdown_set_options
func (recv_ *ObjT) DropdownSetOptions(options *c.Char) {
}

/**
 * Set the options in a drop-down list from a static string (global, static or dynamically allocated).
 * Only the pointer of the option string will be saved.
 * @param obj       pointer to drop-down list object
 * @param options   a static string with '\n' separated options. E.g. "One\nTwo\nThree"
 */
// llgo:link (*ObjT).DropdownSetOptionsStatic C.lv_dropdown_set_options_static
func (recv_ *ObjT) DropdownSetOptionsStatic(options *c.Char) {
}

/**
 * Add an options to a drop-down list from a string.  Only works for non-static options.
 * @param obj       pointer to drop-down list object
 * @param option    a string without '\n'. E.g. "Four"
 * @param pos       the insert position, indexed from 0, LV_DROPDOWN_POS_LAST = end of string
 */
// llgo:link (*ObjT).DropdownAddOption C.lv_dropdown_add_option
func (recv_ *ObjT) DropdownAddOption(option *c.Char, pos c.Uint32T) {
}

/**
 * Clear all options in a drop-down list.  Works with both static and dynamic options.
 * @param obj       pointer to drop-down list object
 */
// llgo:link (*ObjT).DropdownClearOptions C.lv_dropdown_clear_options
func (recv_ *ObjT) DropdownClearOptions() {
}

/**
 * Set the selected option
 * @param obj       pointer to drop-down list object
 * @param sel_opt   id of the selected option (0 ... number of option - 1);
 */
// llgo:link (*ObjT).DropdownSetSelected C.lv_dropdown_set_selected
func (recv_ *ObjT) DropdownSetSelected(sel_opt c.Uint32T) {
}

/**
 * Set the direction of the a drop-down list
 * @param obj       pointer to a drop-down list object
 * @param dir       LV_DIR_LEFT/RIGHT/TOP/BOTTOM
 */
// llgo:link (*ObjT).DropdownSetDir C.lv_dropdown_set_dir
func (recv_ *ObjT) DropdownSetDir(dir DirT) {
}

/**
 * Set an arrow or other symbol to display when on drop-down list's button. Typically a down caret or arrow.
 * @param obj       pointer to drop-down list object
 * @param symbol    a text like `LV_SYMBOL_DOWN`, an image (pointer or path) or NULL to not draw symbol icon
 * @note angle and zoom transformation can be applied if the symbol is an image.
 * E.g. when drop down is checked (opened) rotate the symbol by 180 degree
 */
// llgo:link (*ObjT).DropdownSetSymbol C.lv_dropdown_set_symbol
func (recv_ *ObjT) DropdownSetSymbol(symbol c.Pointer) {
}

/**
 * Set whether the selected option in the list should be highlighted or not
 * @param obj       pointer to drop-down list object
 * @param en        true: highlight enabled; false: disabled
 */
// llgo:link (*ObjT).DropdownSetSelectedHighlight C.lv_dropdown_set_selected_highlight
func (recv_ *ObjT) DropdownSetSelectedHighlight(en bool) {
}

/**
 * Get the list of a drop-down to allow styling or other modifications
 * @param obj       pointer to a drop-down list object
 * @return          pointer to the list of the drop-down
 */
// llgo:link (*ObjT).DropdownGetList C.lv_dropdown_get_list
func (recv_ *ObjT) DropdownGetList() *ObjT {
	return nil
}

/**
 * Get text of the drop-down list's button.
 * @param obj   pointer to a drop-down list object
 * @return      the text as string, `NULL` if no text
 */
// llgo:link (*ObjT).DropdownGetText C.lv_dropdown_get_text
func (recv_ *ObjT) DropdownGetText() *c.Char {
	return nil
}

/**
 * Get the options of a drop-down list
 * @param obj       pointer to drop-down list object
 * @return          the options separated by '\n'-s (E.g. "Option1\nOption2\nOption3")
 */
// llgo:link (*ObjT).DropdownGetOptions C.lv_dropdown_get_options
func (recv_ *ObjT) DropdownGetOptions() *c.Char {
	return nil
}

/**
 * Get the index of the selected option
 * @param obj       pointer to drop-down list object
 * @return          index of the selected option (0 ... number of option - 1);
 */
// llgo:link (*ObjT).DropdownGetSelected C.lv_dropdown_get_selected
func (recv_ *ObjT) DropdownGetSelected() c.Uint32T {
	return 0
}

/**
 * Get the total number of options
 * @param obj       pointer to drop-down list object
 * @return          the total number of options in the list
 */
// llgo:link (*ObjT).DropdownGetOptionCount C.lv_dropdown_get_option_count
func (recv_ *ObjT) DropdownGetOptionCount() c.Uint32T {
	return 0
}

/**
 * Get the current selected option as a string
 * @param obj       pointer to drop-down object
 * @param buf       pointer to an array to store the string
 * @param buf_size  size of `buf` in bytes. 0: to ignore it.
 */
// llgo:link (*ObjT).DropdownGetSelectedStr C.lv_dropdown_get_selected_str
func (recv_ *ObjT) DropdownGetSelectedStr(buf *c.Char, buf_size c.Uint32T) {
}

/**
 * Get the index of an option.
 * @param obj       pointer to drop-down object
 * @param option    an option as string
 * @return          index of `option` in the list of all options. -1 if not found.
 */
// llgo:link (*ObjT).DropdownGetOptionIndex C.lv_dropdown_get_option_index
func (recv_ *ObjT) DropdownGetOptionIndex(option *c.Char) c.Int32T {
	return 0
}

/**
 * Get the symbol on the drop-down list. Typically a down caret or arrow.
 * @param obj       pointer to drop-down list object
 * @return          the symbol or NULL if not enabled
 */
// llgo:link (*ObjT).DropdownGetSymbol C.lv_dropdown_get_symbol
func (recv_ *ObjT) DropdownGetSymbol() *c.Char {
	return nil
}

/**
 * Get whether the selected option in the list should be highlighted or not
 * @param obj       pointer to drop-down list object
 * @return          true: highlight enabled; false: disabled
 */
// llgo:link (*ObjT).DropdownGetSelectedHighlight C.lv_dropdown_get_selected_highlight
func (recv_ *ObjT) DropdownGetSelectedHighlight() bool {
	return false
}

/**
 * Get the direction of the drop-down list
 * @param obj       pointer to a drop-down list object
 * @return          LV_DIR_LEF/RIGHT/TOP/BOTTOM
 */
// llgo:link (*ObjT).DropdownGetDir C.lv_dropdown_get_dir
func (recv_ *ObjT) DropdownGetDir() DirT {
	return 0
}

/**
 * Open the drop.down list
 * @param dropdown_obj       pointer to drop-down list object
 */
// llgo:link (*ObjT).DropdownOpen C.lv_dropdown_open
func (recv_ *ObjT) DropdownOpen() {
}

/**
 * Close (Collapse) the drop-down list
 * @param obj       pointer to drop-down list object
 */
// llgo:link (*ObjT).DropdownClose C.lv_dropdown_close
func (recv_ *ObjT) DropdownClose() {
}

/**
 * Tells whether the list is opened or not
 * @param obj       pointer to a drop-down list object
 * @return          true if the list os opened
 */
// llgo:link (*ObjT).DropdownIsOpen C.lv_dropdown_is_open
func (recv_ *ObjT) DropdownIsOpen() bool {
	return false
}

type MenuModeHeaderT c.Int

const (
	MENU_HEADER_TOP_FIXED    MenuModeHeaderT = 0
	MENU_HEADER_TOP_UNFIXED  MenuModeHeaderT = 1
	MENU_HEADER_BOTTOM_FIXED MenuModeHeaderT = 2
)

type MenuModeRootBackButtonT c.Int

const (
	MENU_ROOT_BACK_BUTTON_DISABLED MenuModeRootBackButtonT = 0
	MENU_ROOT_BACK_BUTTON_ENABLED  MenuModeRootBackButtonT = 1
)

/**
 * Create a menu object
 * @param parent    pointer to an object, it will be the parent of the new menu
 * @return          pointer to the created menu
 */
// llgo:link (*ObjT).MenuCreate C.lv_menu_create
func (recv_ *ObjT) MenuCreate() *ObjT {
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
// llgo:link (*ObjT).MenuPageCreate C.lv_menu_page_create
func (recv_ *ObjT) MenuPageCreate(title *c.Char) *ObjT {
	return nil
}

/**
 * Create a menu cont object
 * @param parent    pointer to a menu page object, it will be the parent of the new menu cont object
 * @return          pointer to the created menu cont
 */
// llgo:link (*ObjT).MenuContCreate C.lv_menu_cont_create
func (recv_ *ObjT) MenuContCreate() *ObjT {
	return nil
}

/**
 * Create a menu section object
 * @param parent    pointer to a menu page object, it will be the parent of the new menu section object
 * @return          pointer to the created menu section
 */
// llgo:link (*ObjT).MenuSectionCreate C.lv_menu_section_create
func (recv_ *ObjT) MenuSectionCreate() *ObjT {
	return nil
}

/**
 * Create a menu separator object
 * @param parent    pointer to a menu page object, it will be the parent of the new menu separator object
 * @return          pointer to the created menu separator
 */
// llgo:link (*ObjT).MenuSeparatorCreate C.lv_menu_separator_create
func (recv_ *ObjT) MenuSeparatorCreate() *ObjT {
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
// llgo:link (*ObjT).MenuSetPage C.lv_menu_set_page
func (recv_ *ObjT) MenuSetPage(page *ObjT) {
}

/**
 * Set menu page title
 * @param page      pointer to the menu page
 * @param title     pointer to text for title in header (NULL to not display title)
 */
// llgo:link (*ObjT).MenuSetPageTitle C.lv_menu_set_page_title
func (recv_ *ObjT) MenuSetPageTitle(title *c.Char) {
}

/**
 * Set menu page title with a static text. It will not be saved by the label so the 'text' variable
 * has to be 'alive' while the page exists.
 * @param page      pointer to the menu page
 * @param title     pointer to text for title in header (NULL to not display title)
 */
// llgo:link (*ObjT).MenuSetPageTitleStatic C.lv_menu_set_page_title_static
func (recv_ *ObjT) MenuSetPageTitleStatic(title *c.Char) {
}

/**
 * Set menu page to display in sidebar
 * @param obj       pointer to the menu
 * @param page      pointer to the menu page to set (NULL to clear sidebar)
 */
// llgo:link (*ObjT).MenuSetSidebarPage C.lv_menu_set_sidebar_page
func (recv_ *ObjT) MenuSetSidebarPage(page *ObjT) {
}

/**
 * Set the how the header should behave and its position
 * @param obj       pointer to a menu
 * @param mode      LV_MENU_HEADER_TOP_FIXED/TOP_UNFIXED/BOTTOM_FIXED
 */
// llgo:link (*ObjT).MenuSetModeHeader C.lv_menu_set_mode_header
func (recv_ *ObjT) MenuSetModeHeader(mode MenuModeHeaderT) {
}

/**
 * Set whether back button should appear at root
 * @param obj       pointer to a menu
 * @param mode      LV_MENU_ROOT_BACK_BUTTON_DISABLED/ENABLED
 */
// llgo:link (*ObjT).MenuSetModeRootBackButton C.lv_menu_set_mode_root_back_button
func (recv_ *ObjT) MenuSetModeRootBackButton(mode MenuModeRootBackButtonT) {
}

/**
 * Add menu to the menu item
 * @param menu      pointer to the menu
 * @param obj       pointer to the obj
 * @param page      pointer to the page to load when obj is clicked
 */
// llgo:link (*ObjT).MenuSetLoadPageEvent C.lv_menu_set_load_page_event
func (recv_ *ObjT) MenuSetLoadPageEvent(obj *ObjT, page *ObjT) {
}

/*=====================
 * Getter functions
 *====================*/
/**
* Get a pointer to menu page that is currently displayed in main
* @param obj        pointer to the menu
* @return           pointer to current page
 */
// llgo:link (*ObjT).MenuGetCurMainPage C.lv_menu_get_cur_main_page
func (recv_ *ObjT) MenuGetCurMainPage() *ObjT {
	return nil
}

/**
* Get a pointer to menu page that is currently displayed in sidebar
* @param obj        pointer to the menu
* @return           pointer to current page
 */
// llgo:link (*ObjT).MenuGetCurSidebarPage C.lv_menu_get_cur_sidebar_page
func (recv_ *ObjT) MenuGetCurSidebarPage() *ObjT {
	return nil
}

/**
* Get a pointer to main header obj
* @param obj        pointer to the menu
* @return           pointer to main header obj
 */
// llgo:link (*ObjT).MenuGetMainHeader C.lv_menu_get_main_header
func (recv_ *ObjT) MenuGetMainHeader() *ObjT {
	return nil
}

/**
* Get a pointer to main header back btn obj
* @param obj        pointer to the menu
* @return           pointer to main header back btn obj
 */
// llgo:link (*ObjT).MenuGetMainHeaderBackButton C.lv_menu_get_main_header_back_button
func (recv_ *ObjT) MenuGetMainHeaderBackButton() *ObjT {
	return nil
}

/**
* Get a pointer to sidebar header obj
* @param obj        pointer to the menu
* @return           pointer to sidebar header obj
 */
// llgo:link (*ObjT).MenuGetSidebarHeader C.lv_menu_get_sidebar_header
func (recv_ *ObjT) MenuGetSidebarHeader() *ObjT {
	return nil
}

/**
* Get a pointer to sidebar header obj
* @param obj        pointer to the menu
* @return           pointer to sidebar header back btn obj
 */
// llgo:link (*ObjT).MenuGetSidebarHeaderBackButton C.lv_menu_get_sidebar_header_back_button
func (recv_ *ObjT) MenuGetSidebarHeaderBackButton() *ObjT {
	return nil
}

/**
 * Check if an obj is a root back btn
 * @param menu      pointer to the menu
 * @param obj       pointer to the back button
 * @return          true if it is a root back btn
 */
// llgo:link (*ObjT).MenuBackButtonIsRoot C.lv_menu_back_button_is_root
func (recv_ *ObjT) MenuBackButtonIsRoot(obj *ObjT) bool {
	return false
}

/**
 * Clear menu history
 * @param obj       pointer to the menu
 */
// llgo:link (*ObjT).MenuClearHistory C.lv_menu_clear_history
func (recv_ *ObjT) MenuClearHistory() {
}

type ChartTypeT c.Int

const (
	CHART_TYPE_NONE    ChartTypeT = 0
	CHART_TYPE_LINE    ChartTypeT = 1
	CHART_TYPE_BAR     ChartTypeT = 2
	CHART_TYPE_SCATTER ChartTypeT = 3
)

type ChartUpdateModeT c.Int

const (
	CHART_UPDATE_MODE_SHIFT    ChartUpdateModeT = 0
	CHART_UPDATE_MODE_CIRCULAR ChartUpdateModeT = 1
)

type ChartAxisT c.Int

const (
	CHART_AXIS_PRIMARY_Y   ChartAxisT = 0
	CHART_AXIS_SECONDARY_Y ChartAxisT = 1
	CHART_AXIS_PRIMARY_X   ChartAxisT = 2
	CHART_AXIS_SECONDARY_X ChartAxisT = 4
	CHART_AXIS_LAST        ChartAxisT = 5
)

/**
 * Create a chart object
 * @param parent    pointer to an object, it will be the parent of the new chart
 * @return          pointer to the created chart
 */
// llgo:link (*ObjT).ChartCreate C.lv_chart_create
func (recv_ *ObjT) ChartCreate() *ObjT {
	return nil
}

/**
 * Set a new type for a chart
 * @param obj       pointer to a chart object
 * @param type      new type of the chart (from 'lv_chart_type_t' enum)
 */
// llgo:link (*ObjT).ChartSetType C.lv_chart_set_type
func (recv_ *ObjT) ChartSetType(type_ ChartTypeT) {
}

/**
 * Set the number of points on a data line on a chart
 * @param obj       pointer to a chart object
 * @param cnt       new number of points on the data lines
 */
// llgo:link (*ObjT).ChartSetPointCount C.lv_chart_set_point_count
func (recv_ *ObjT) ChartSetPointCount(cnt c.Uint32T) {
}

/**
 * Set the minimal and maximal y values on an axis
 * @param obj       pointer to a chart object
 * @param axis      `LV_CHART_AXIS_PRIMARY_Y` or `LV_CHART_AXIS_SECONDARY_Y`
 * @param min       minimum value of the y axis
 * @param max       maximum value of the y axis
 */
// llgo:link (*ObjT).ChartSetAxisRange C.lv_chart_set_axis_range
func (recv_ *ObjT) ChartSetAxisRange(axis ChartAxisT, min c.Int32T, max c.Int32T) {
}

/**
 * Set the minimal values on an axis
 * @param obj       pointer to a chart object
 * @param axis      `LV_CHART_AXIS_PRIMARY_Y` or `LV_CHART_AXIS_SECONDARY_Y`
 * @param min       minimal value of the y axis
 */
// llgo:link (*ObjT).ChartSetAxisMinValue C.lv_chart_set_axis_min_value
func (recv_ *ObjT) ChartSetAxisMinValue(axis ChartAxisT, min c.Int32T) {
}

/**
 * Set the maximal y values on an axis
 * @param obj       pointer to a chart object
 * @param axis      `LV_CHART_AXIS_PRIMARY_Y` or `LV_CHART_AXIS_SECONDARY_Y`
 * @param max       maximum value of the y axis
 */
// llgo:link (*ObjT).ChartSetAxisMaxValue C.lv_chart_set_axis_max_value
func (recv_ *ObjT) ChartSetAxisMaxValue(axis ChartAxisT, max c.Int32T) {
}

/**
 * Set update mode of the chart object. Affects
 * @param obj              pointer to a chart object
 * @param update_mode      the update mode
 */
// llgo:link (*ObjT).ChartSetUpdateMode C.lv_chart_set_update_mode
func (recv_ *ObjT) ChartSetUpdateMode(update_mode ChartUpdateModeT) {
}

/**
 * Set the number of horizontal and vertical division lines
 * @param obj       pointer to a chart object
 * @param hdiv      number of horizontal division lines
 * @param vdiv      number of vertical division lines
 */
// llgo:link (*ObjT).ChartSetDivLineCount C.lv_chart_set_div_line_count
func (recv_ *ObjT) ChartSetDivLineCount(hdiv c.Uint32T, vdiv c.Uint32T) {
}

/**
 * Set the number of horizontal division lines
 * @param obj       pointer to a chart object
 * @param cnt       number of horizontal division lines
 */
// llgo:link (*ObjT).ChartSetHorDivLineCount C.lv_chart_set_hor_div_line_count
func (recv_ *ObjT) ChartSetHorDivLineCount(cnt c.Uint32T) {
}

/**
 * Set the number of vertical division lines
 * @param obj       pointer to a chart object
 * @param cnt       number of vertical division lines
 */
// llgo:link (*ObjT).ChartSetVerDivLineCount C.lv_chart_set_ver_div_line_count
func (recv_ *ObjT) ChartSetVerDivLineCount(cnt c.Uint32T) {
}

/**
 * Get the type of a chart
 * @param obj       pointer to chart object
 * @return          type of the chart (from 'lv_chart_t' enum)
 */
// llgo:link (*ObjT).ChartGetType C.lv_chart_get_type
func (recv_ *ObjT) ChartGetType() ChartTypeT {
	return 0
}

/**
 * Get the data point number per data line on chart
 * @param obj       pointer to chart object
 * @return          point number on each data line
 */
// llgo:link (*ObjT).ChartGetPointCount C.lv_chart_get_point_count
func (recv_ *ObjT) ChartGetPointCount() c.Uint32T {
	return 0
}

/**
 * Get the current index of the x-axis start point in the data array
 * @param obj       pointer to a chart object
 * @param ser       pointer to a data series on 'chart'
 * @return          the index of the current x start point in the data array
 */
// llgo:link (*ObjT).ChartGetXStartPoint C.lv_chart_get_x_start_point
func (recv_ *ObjT) ChartGetXStartPoint(ser *ChartSeriesT) c.Uint32T {
	return 0
}

/**
 * Get the position of a point to the chart.
 * @param obj       pointer to a chart object
 * @param ser       pointer to series
 * @param id        the index.
 * @param p_out     store the result position here
 */
// llgo:link (*ObjT).ChartGetPointPosById C.lv_chart_get_point_pos_by_id
func (recv_ *ObjT) ChartGetPointPosById(ser *ChartSeriesT, id c.Uint32T, p_out *PointT) {
}

/**
 * Refresh a chart if its data line has changed
 * @param   obj   pointer to chart object
 */
// llgo:link (*ObjT).ChartRefresh C.lv_chart_refresh
func (recv_ *ObjT) ChartRefresh() {
}

/**
 * Allocate and add a data series to the chart
 * @param obj       pointer to a chart object
 * @param color     color of the data series
 * @param axis      the y axis to which the series should be attached (::LV_CHART_AXIS_PRIMARY_Y or ::LV_CHART_AXIS_SECONDARY_Y)
 * @return          pointer to the allocated data series or NULL on failure
 */
// llgo:link (*ObjT).ChartAddSeries C.lv_chart_add_series
func (recv_ *ObjT) ChartAddSeries(color ColorT, axis ChartAxisT) *ChartSeriesT {
	return nil
}

/**
 * Deallocate and remove a data series from a chart
 * @param obj       pointer to a chart object
 * @param series    pointer to a data series on 'chart'
 */
// llgo:link (*ObjT).ChartRemoveSeries C.lv_chart_remove_series
func (recv_ *ObjT) ChartRemoveSeries(series *ChartSeriesT) {
}

/**
 * Hide/Unhide a single series of a chart.
 * @param chart     pointer to a chart object.
 * @param series    pointer to a series object
 * @param hide      true: hide the series
 */
// llgo:link (*ObjT).ChartHideSeries C.lv_chart_hide_series
func (recv_ *ObjT) ChartHideSeries(series *ChartSeriesT, hide bool) {
}

/**
 * Change the color of a series
 * @param chart     pointer to a chart object.
 * @param series    pointer to a series object
 * @param color     the new color of the series
 */
// llgo:link (*ObjT).ChartSetSeriesColor C.lv_chart_set_series_color
func (recv_ *ObjT) ChartSetSeriesColor(series *ChartSeriesT, color ColorT) {
}

/**
 * Get the color of a series
 * @param chart     pointer to a chart object.
 * @param series    pointer to a series object
 * @return          the color of the series
 */
// llgo:link (*ObjT).ChartGetSeriesColor C.lv_chart_get_series_color
func (recv_ *ObjT) ChartGetSeriesColor(series *ChartSeriesT) ColorT {
	return ColorT{}
}

/**
 * Set the index of the x-axis start point in the data array.
 * This point will be considers the first (left) point and the other points will be drawn after it.
 * @param obj       pointer to a chart object
 * @param ser       pointer to a data series on 'chart'
 * @param id        the index of the x point in the data array
 */
// llgo:link (*ObjT).ChartSetXStartPoint C.lv_chart_set_x_start_point
func (recv_ *ObjT) ChartSetXStartPoint(ser *ChartSeriesT, id c.Uint32T) {
}

/**
 * Get the next series.
 * @param chart     pointer to a chart
 * @param ser      the previous series or NULL to get the first
 * @return          the next series or NULL if there is no more.
 */
// llgo:link (*ObjT).ChartGetSeriesNext C.lv_chart_get_series_next
func (recv_ *ObjT) ChartGetSeriesNext(ser *ChartSeriesT) *ChartSeriesT {
	return nil
}

/**
 * Add a cursor with a given color
 * @param obj       pointer to chart object
 * @param color     color of the cursor
 * @param dir       direction of the cursor. `LV_DIR_RIGHT/LEFT/TOP/DOWN/HOR/VER/ALL`. OR-ed values are possible
 * @return          pointer to the created cursor
 */
// llgo:link (*ObjT).ChartAddCursor C.lv_chart_add_cursor
func (recv_ *ObjT) ChartAddCursor(color ColorT, dir DirT) *ChartCursorT {
	return nil
}

/**
 * Set the coordinate of the cursor with respect to the paddings
 * @param chart     pointer to a chart object
 * @param cursor    pointer to the cursor
 * @param pos       the new coordinate of cursor relative to the chart
 */
// llgo:link (*ObjT).ChartSetCursorPos C.lv_chart_set_cursor_pos
func (recv_ *ObjT) ChartSetCursorPos(cursor *ChartCursorT, pos *PointT) {
}

/**
 * Set the X coordinate of the cursor with respect to the paddings
 * @param chart     pointer to a chart object
 * @param cursor    pointer to the cursor
 * @param x         the new X coordinate of cursor relative to the chart
 */
// llgo:link (*ObjT).ChartSetCursorPosX C.lv_chart_set_cursor_pos_x
func (recv_ *ObjT) ChartSetCursorPosX(cursor *ChartCursorT, x c.Int32T) {
}

/**
 * Set the coordinate of the cursor with respect to the paddings
 * @param chart     pointer to a chart object
 * @param cursor    pointer to the cursor
 * @param y         the new Y coordinate of cursor relative to the chart
 */
// llgo:link (*ObjT).ChartSetCursorPosY C.lv_chart_set_cursor_pos_y
func (recv_ *ObjT) ChartSetCursorPosY(cursor *ChartCursorT, y c.Int32T) {
}

/**
 * Stick the cursor to a point
 * @param chart     pointer to a chart object
 * @param cursor    pointer to the cursor
 * @param ser       pointer to a series
 * @param point_id  the point's index or `LV_CHART_POINT_NONE` to not assign to any points.
 */
// llgo:link (*ObjT).ChartSetCursorPoint C.lv_chart_set_cursor_point
func (recv_ *ObjT) ChartSetCursorPoint(cursor *ChartCursorT, ser *ChartSeriesT, point_id c.Uint32T) {
}

/**
 * Get the coordinate of the cursor with respect to the paddings
 * @param chart     pointer to a chart object
 * @param cursor    pointer to cursor
 * @return          coordinate of the cursor as lv_point_t
 */
// llgo:link (*ObjT).ChartGetCursorPoint C.lv_chart_get_cursor_point
func (recv_ *ObjT) ChartGetCursorPoint(cursor *ChartCursorT) PointT {
	return PointT{}
}

/**
 * Initialize all data points of a series with a value
 * @param obj       pointer to chart object
 * @param ser       pointer to a data series on 'chart'
 * @param value     the new value for all points. `LV_CHART_POINT_NONE` can be used to hide the points.
 */
// llgo:link (*ObjT).ChartSetAllValues C.lv_chart_set_all_values
func (recv_ *ObjT) ChartSetAllValues(ser *ChartSeriesT, value c.Int32T) {
}

/**
 * Set the next point's Y value according to the update mode policy.
 * @param obj       pointer to chart object
 * @param ser       pointer to a data series on 'chart'
 * @param value     the new value of the next data
 */
// llgo:link (*ObjT).ChartSetNextValue C.lv_chart_set_next_value
func (recv_ *ObjT) ChartSetNextValue(ser *ChartSeriesT, value c.Int32T) {
}

/**
 * Set the next point's X and Y value according to the update mode policy.
 * @param obj       pointer to chart object
 * @param ser       pointer to a data series on 'chart'
 * @param x_value   the new X value of the next data
 * @param y_value   the new Y value of the next data
 */
// llgo:link (*ObjT).ChartSetNextValue2 C.lv_chart_set_next_value2
func (recv_ *ObjT) ChartSetNextValue2(ser *ChartSeriesT, x_value c.Int32T, y_value c.Int32T) {
}

/**
 * Same as `lv_chart_set_next_value` but set the values from an array
 * @param obj           pointer to chart object
 * @param ser           pointer to a data series on 'chart'
 * @param values        the new values to set
 * @param values_cnt    number of items in `values`
 */
// llgo:link (*ObjT).ChartSetSeriesValues C.lv_chart_set_series_values
func (recv_ *ObjT) ChartSetSeriesValues(ser *ChartSeriesT, values *c.Int32T, values_cnt c.SizeT) {
}

/**
 * Same as `lv_chart_set_next_value2` but set the values from an array
 * @param obj           pointer to chart object
 * @param ser           pointer to a data series on 'chart'
 * @param x_values      the new values to set on the X axis
 * @param y_values      the new values to set o nthe Y axis
 * @param values_cnt    number of items in `x_values` and `y_values`
 */
// llgo:link (*ObjT).ChartSetSeriesValues2 C.lv_chart_set_series_values2
func (recv_ *ObjT) ChartSetSeriesValues2(ser *ChartSeriesT, x_values *c.Int32T, y_values *c.Int32T, values_cnt c.SizeT) {
}

/**
 * Set an individual point's y value of a chart's series directly based on its index
 * @param obj     pointer to a chart object
 * @param ser     pointer to a data series on 'chart'
 * @param id      the index of the x point in the array
 * @param value   value to assign to array point
 */
// llgo:link (*ObjT).ChartSetSeriesValueById C.lv_chart_set_series_value_by_id
func (recv_ *ObjT) ChartSetSeriesValueById(ser *ChartSeriesT, id c.Uint32T, value c.Int32T) {
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
// llgo:link (*ObjT).ChartSetSeriesValueById2 C.lv_chart_set_series_value_by_id2
func (recv_ *ObjT) ChartSetSeriesValueById2(ser *ChartSeriesT, id c.Uint32T, x_value c.Int32T, y_value c.Int32T) {
}

/**
 * Set an external array for the y data points to use for the chart
 * NOTE: It is the users responsibility to make sure the `point_cnt` matches the external array size.
 * @param obj       pointer to a chart object
 * @param ser       pointer to a data series on 'chart'
 * @param array     external array of points for chart
 */
// llgo:link (*ObjT).ChartSetSeriesExtYArray C.lv_chart_set_series_ext_y_array
func (recv_ *ObjT) ChartSetSeriesExtYArray(ser *ChartSeriesT, array *c.Int32T) {
}

/**
 * Set an external array for the x data points to use for the chart
 * NOTE: It is the users responsibility to make sure the `point_cnt` matches the external array size.
 * @param obj       pointer to a chart object
 * @param ser       pointer to a data series on 'chart'
 * @param array     external array of points for chart
 */
// llgo:link (*ObjT).ChartSetSeriesExtXArray C.lv_chart_set_series_ext_x_array
func (recv_ *ObjT) ChartSetSeriesExtXArray(ser *ChartSeriesT, array *c.Int32T) {
}

/**
 * Get the array of y values of a series
 * @param obj   pointer to a chart object
 * @param ser   pointer to a data series on 'chart'
 * @return      the array of values with 'point_count' elements
 */
// llgo:link (*ObjT).ChartGetSeriesYArray C.lv_chart_get_series_y_array
func (recv_ *ObjT) ChartGetSeriesYArray(ser *ChartSeriesT) *c.Int32T {
	return nil
}

/**
 * Get the array of x values of a series
 * @param obj   pointer to a chart object
 * @param ser   pointer to a data series on 'chart'
 * @return      the array of values with 'point_count' elements
 */
// llgo:link (*ObjT).ChartGetSeriesXArray C.lv_chart_get_series_x_array
func (recv_ *ObjT) ChartGetSeriesXArray(ser *ChartSeriesT) *c.Int32T {
	return nil
}

/**
 * Get the index of the currently pressed point. It's the same for every series.
 * @param obj       pointer to a chart object
 * @return          the index of the point [0 .. point count] or LV_CHART_POINT_ID_NONE if no point is being pressed
 */
// llgo:link (*ObjT).ChartGetPressedPoint C.lv_chart_get_pressed_point
func (recv_ *ObjT) ChartGetPressedPoint() c.Uint32T {
	return 0
}

/**
 * Get the overall offset from the chart's side to the center of the first point.
 * In case of a bar chart it will be the center of the first column group
 * @param obj       pointer to a chart object
 * @return          the offset of the center
 */
// llgo:link (*ObjT).ChartGetFirstPointCenterOffset C.lv_chart_get_first_point_center_offset
func (recv_ *ObjT) ChartGetFirstPointCenterOffset() c.Int32T {
	return 0
}

/**
 * Create a button object
 * @param parent    pointer to an object, it will be the parent of the new button
 * @return          pointer to the created button
 */
// llgo:link (*ObjT).ButtonCreate C.lv_button_create
func (recv_ *ObjT) ButtonCreate() *ObjT {
	return nil
}

type ScaleModeT c.Int

const (
	SCALE_MODE_HORIZONTAL_TOP    ScaleModeT = 0
	SCALE_MODE_HORIZONTAL_BOTTOM ScaleModeT = 1
	SCALE_MODE_VERTICAL_LEFT     ScaleModeT = 2
	SCALE_MODE_VERTICAL_RIGHT    ScaleModeT = 4
	SCALE_MODE_ROUND_INNER       ScaleModeT = 8
	SCALE_MODE_ROUND_OUTER       ScaleModeT = 16
	SCALE_MODE_LAST              ScaleModeT = 17
)

/**
 * Create an scale object
 * @param parent    pointer to an object, it will be the parent of the new scale
 * @return          pointer to created Scale Widget
 */
// llgo:link (*ObjT).ScaleCreate C.lv_scale_create
func (recv_ *ObjT) ScaleCreate() *ObjT {
	return nil
}

/**
 * Set scale mode. See lv_scale_mode_t.
 * @param obj       pointer to Scale Widget
 * @param mode      the new scale mode
 */
// llgo:link (*ObjT).ScaleSetMode C.lv_scale_set_mode
func (recv_ *ObjT) ScaleSetMode(mode ScaleModeT) {
}

/**
 * Set scale total tick count (including minor and major ticks).
 * @param obj       pointer to Scale Widget
 * @param total_tick_count    New total tick count
 */
// llgo:link (*ObjT).ScaleSetTotalTickCount C.lv_scale_set_total_tick_count
func (recv_ *ObjT) ScaleSetTotalTickCount(total_tick_count c.Uint32T) {
}

/**
 * Sets how often major ticks are drawn.
 * @param obj                 pointer to Scale Widget
 * @param major_tick_every    the new count for major tick drawing
 */
// llgo:link (*ObjT).ScaleSetMajorTickEvery C.lv_scale_set_major_tick_every
func (recv_ *ObjT) ScaleSetMajorTickEvery(major_tick_every c.Uint32T) {
}

/**
 * Sets label visibility.
 * @param obj           pointer to Scale Widget
 * @param show_label    true/false to enable tick label
 */
// llgo:link (*ObjT).ScaleSetLabelShow C.lv_scale_set_label_show
func (recv_ *ObjT) ScaleSetLabelShow(show_label bool) {
}

/**
 * Set minimum and maximum values on Scale.
 * @param obj       pointer to Scale Widget
 * @param min       minimum value of Scale
 * @param max       maximum value of Scale
 */
// llgo:link (*ObjT).ScaleSetRange C.lv_scale_set_range
func (recv_ *ObjT) ScaleSetRange(min c.Int32T, max c.Int32T) {
}

/**
 * Set minimum values on Scale.
 * @param obj       pointer to Scale Widget
 * @param min       minimum value of Scale
 */
// llgo:link (*ObjT).ScaleSetMinValue C.lv_scale_set_min_value
func (recv_ *ObjT) ScaleSetMinValue(min c.Int32T) {
}

/**
 * Set maximum values on Scale.
 * @param obj       pointer to Scale Widget
 * @param min       minimum value of Scale
 */
// llgo:link (*ObjT).ScaleSetMaxValue C.lv_scale_set_max_value
func (recv_ *ObjT) ScaleSetMaxValue(max c.Int32T) {
}

/**
 * Set angle between the low end and the high end of the Scale.
 * (Applies only to round Scales.)
 * @param obj         pointer to Scale Widget
 * @param max_angle   angle in degrees from Scale minimum where top end of Scale will be drawn
 */
// llgo:link (*ObjT).ScaleSetAngleRange C.lv_scale_set_angle_range
func (recv_ *ObjT) ScaleSetAngleRange(angle_range c.Uint32T) {
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
// llgo:link (*ObjT).ScaleSetRotation C.lv_scale_set_rotation
func (recv_ *ObjT) ScaleSetRotation(rotation c.Int32T) {
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
// llgo:link (*ObjT).ScaleSetLineNeedleValue C.lv_scale_set_line_needle_value
func (recv_ *ObjT) ScaleSetLineNeedleValue(needle_line *ObjT, needle_length c.Int32T, value c.Int32T) {
}

/**
* Point image needle to specified value;
  image must point to the right. E.g. -O------>
* @param obj              pointer to Scale Widget
* @param needle_img       pointer to needle's Image
* @param value            Scale value needle will point to
*/
// llgo:link (*ObjT).ScaleSetImageNeedleValue C.lv_scale_set_image_needle_value
func (recv_ *ObjT) ScaleSetImageNeedleValue(needle_img *ObjT, value c.Int32T) {
}

/**
 * Set custom text source for major ticks labels.
 * @param obj       pointer to Scale Widget
 * @param txt_src   pointer to an array of strings which will be display at major ticks;
 *                  last element must be a NULL pointer.
 */
// llgo:link (*ObjT).ScaleSetTextSrc C.lv_scale_set_text_src
func (recv_ *ObjT) ScaleSetTextSrc(txt_src **c.Char) {
}

/**
 * Draw Scale after all its children are drawn.
 * @param obj       pointer to Scale Widget
 * @param en        true: enable post draw
 */
// llgo:link (*ObjT).ScaleSetPostDraw C.lv_scale_set_post_draw
func (recv_ *ObjT) ScaleSetPostDraw(en bool) {
}

/**
 * Draw Scale ticks on top of all other parts.
 * @param obj       pointer to Scale Widget
 * @param en        true: enable draw ticks on top of all parts
 */
// llgo:link (*ObjT).ScaleSetDrawTicksOnTop C.lv_scale_set_draw_ticks_on_top
func (recv_ *ObjT) ScaleSetDrawTicksOnTop(en bool) {
}

/**
 * Add a Section to specified Scale.  Section will not be drawn until
 * a valid range is set for it using `lv_scale_set_section_range()`.
 * @param obj       pointer to Scale Widget
 * @return          pointer to new Section
 */
// llgo:link (*ObjT).ScaleAddSection C.lv_scale_add_section
func (recv_ *ObjT) ScaleAddSection() *ScaleSectionT {
	return nil
}

/**
 * DEPRECATED, use lv_scale_set_section_range instead.
 * Set range for specified Scale Section
 * @param section       pointer to Section
 * @param range_min     Section new minimum value
 * @param range_max     Section new maximum value
 */
// llgo:link (*ScaleSectionT).ScaleSectionSetRange C.lv_scale_section_set_range
func (recv_ *ScaleSectionT) ScaleSectionSetRange(min c.Int32T, max c.Int32T) {
}

/**
 * Set the range of a scale section
 * @param scale         pointer to scale
 * @param section       pointer to section
 * @param range_min     the section's new minimum value
 * @param range_max     the section's new maximum value
 */
// llgo:link (*ObjT).ScaleSetSectionRange C.lv_scale_set_section_range
func (recv_ *ObjT) ScaleSetSectionRange(section *ScaleSectionT, min c.Int32T, max c.Int32T) {
}

/**
 * Set the minimum value of a scale section
 * @param scale         pointer to scale
 * @param section       pointer to section
 * @param min           the section's new minimum value
 */
// llgo:link (*ObjT).ScaleSetSectionMinValue C.lv_scale_set_section_min_value
func (recv_ *ObjT) ScaleSetSectionMinValue(section *ScaleSectionT, min c.Int32T) {
}

/**
 * Set the maximum value of a scale section
 * @param scale         pointer to scale
 * @param section       pointer to section
 * @param max           the section's new maximum value
 */
// llgo:link (*ObjT).ScaleSetSectionMaxValue C.lv_scale_set_section_max_value
func (recv_ *ObjT) ScaleSetSectionMaxValue(section *ScaleSectionT, max c.Int32T) {
}

/**
 * DEPRECATED, use lv_scale_set_section_style_main/indicator/items instead.
 * Set style for specified part of Section.
 * @param section             pointer to Section
 * @param part                the part of the Scale the style will apply to, e.g. LV_PART_INDICATOR
 * @param section_part_style  pointer to style to apply
 */
// llgo:link (*ScaleSectionT).ScaleSectionSetStyle C.lv_scale_section_set_style
func (recv_ *ScaleSectionT) ScaleSectionSetStyle(part PartT, section_part_style *StyleT) {
}

/**
 * Set the style of the line on a section.
 * @param scale         pointer to scale
 * @param section       pointer to section
 * @param style         point to a style
 */
// llgo:link (*ObjT).ScaleSetSectionStyleMain C.lv_scale_set_section_style_main
func (recv_ *ObjT) ScaleSetSectionStyleMain(section *ScaleSectionT, style *StyleT) {
}

/**
 * Set the style of the major ticks and label on a section.
 * @param scale         pointer to scale
 * @param section       pointer to section
 * @param style         point to a style
 */
// llgo:link (*ObjT).ScaleSetSectionStyleIndicator C.lv_scale_set_section_style_indicator
func (recv_ *ObjT) ScaleSetSectionStyleIndicator(section *ScaleSectionT, style *StyleT) {
}

/**
 * Set the style of the minor ticks on a section.
 * @param scale         pointer to scale
 * @param section       pointer to section
 * @param style         point to a style
 */
// llgo:link (*ObjT).ScaleSetSectionStyleItems C.lv_scale_set_section_style_items
func (recv_ *ObjT) ScaleSetSectionStyleItems(section *ScaleSectionT, style *StyleT) {
}

/**
 * Get scale mode. See lv_scale_mode_t
 * @param obj   pointer to Scale Widget
 * @return      Scale mode
 */
// llgo:link (*ObjT).ScaleGetMode C.lv_scale_get_mode
func (recv_ *ObjT) ScaleGetMode() ScaleModeT {
	return 0
}

/**
 * Get scale total tick count (including minor and major ticks)
 * @param obj   pointer to Scale Widget
 * @return      Scale total tick count
 */
// llgo:link (*ObjT).ScaleGetTotalTickCount C.lv_scale_get_total_tick_count
func (recv_ *ObjT) ScaleGetTotalTickCount() c.Int32T {
	return 0
}

/**
 * Get how often the major tick will be drawn
 * @param obj   pointer to Scale Widget
 * @return      Scale major tick every count
 */
// llgo:link (*ObjT).ScaleGetMajorTickEvery C.lv_scale_get_major_tick_every
func (recv_ *ObjT) ScaleGetMajorTickEvery() c.Int32T {
	return 0
}

/**
 * Get angular location of low end of Scale.
 * @param obj   pointer to Scale Widget
 * @return      Scale low end anglular location
 */
// llgo:link (*ObjT).ScaleGetRotation C.lv_scale_get_rotation
func (recv_ *ObjT) ScaleGetRotation() c.Int32T {
	return 0
}

/**
 * Gets label visibility
 * @param obj   pointer to Scale Widget
 * @return      true if tick label is enabled, false otherwise
 */
// llgo:link (*ObjT).ScaleGetLabelShow C.lv_scale_get_label_show
func (recv_ *ObjT) ScaleGetLabelShow() bool {
	return false
}

/**
 * Get Scale's range in degrees
 * @param obj   pointer to Scale Widget
 * @return      Scale's angle_range
 */
// llgo:link (*ObjT).ScaleGetAngleRange C.lv_scale_get_angle_range
func (recv_ *ObjT) ScaleGetAngleRange() c.Uint32T {
	return 0
}

/**
 * Get minimum value for Scale
 * @param obj   pointer to Scale Widget
 * @return      Scale's minimum value
 */
// llgo:link (*ObjT).ScaleGetRangeMinValue C.lv_scale_get_range_min_value
func (recv_ *ObjT) ScaleGetRangeMinValue() c.Int32T {
	return 0
}

/**
 * Get maximum value for Scale
 * @param obj   pointer to Scale Widget
 * @return      Scale's maximum value
 */
// llgo:link (*ObjT).ScaleGetRangeMaxValue C.lv_scale_get_range_max_value
func (recv_ *ObjT) ScaleGetRangeMaxValue() c.Int32T {
	return 0
}

/**
 * Create a led object
 * @param parent    pointer to an object, it will be the parent of the new led
 * @return          pointer to the created led
 */
// llgo:link (*ObjT).LedCreate C.lv_led_create
func (recv_ *ObjT) LedCreate() *ObjT {
	return nil
}

/**
 * Set the color of the LED
 * @param led       pointer to a LED object
 * @param color     the color of the LED
 */
// llgo:link (*ObjT).LedSetColor C.lv_led_set_color
func (recv_ *ObjT) LedSetColor(color ColorT) {
}

/**
 * Set the brightness of a LED object
 * @param led       pointer to a LED object
 * @param bright    LV_LED_BRIGHT_MIN (max. dark) ... LV_LED_BRIGHT_MAX (max. light)
 */
// llgo:link (*ObjT).LedSetBrightness C.lv_led_set_brightness
func (recv_ *ObjT) LedSetBrightness(bright c.Uint8T) {
}

/**
 * Light on a LED
 * @param led       pointer to a LED object
 */
// llgo:link (*ObjT).LedOn C.lv_led_on
func (recv_ *ObjT) LedOn() {
}

/**
 * Light off a LED
 * @param led       pointer to a LED object
 */
// llgo:link (*ObjT).LedOff C.lv_led_off
func (recv_ *ObjT) LedOff() {
}

/**
 * Toggle the state of a LED
 * @param led       pointer to a LED object
 */
// llgo:link (*ObjT).LedToggle C.lv_led_toggle
func (recv_ *ObjT) LedToggle() {
}

/**
 * Get the brightness of a LED object
 * @param obj       pointer to LED object
 * @return bright   0 (max. dark) ... 255 (max. light)
 */
// llgo:link (*ObjT).LedGetBrightness C.lv_led_get_brightness
func (recv_ *ObjT) LedGetBrightness() c.Uint8T {
	return 0
}

type ArcModeT c.Int

const (
	ARC_MODE_NORMAL      ArcModeT = 0
	ARC_MODE_SYMMETRICAL ArcModeT = 1
	ARC_MODE_REVERSE     ArcModeT = 2
)

/**
 * Create an arc object
 * @param parent    pointer to an object, it will be the parent of the new arc
 * @return          pointer to the created arc
 */
// llgo:link (*ObjT).ArcCreate C.lv_arc_create
func (recv_ *ObjT) ArcCreate() *ObjT {
	return nil
}

/**
 * Set the start angle of an arc. 0 deg: right, 90 bottom, etc.
 * @param obj       pointer to an arc object
 * @param start     the start angle. (if `LV_USE_FLOAT` is enabled it can be fractional too.)
 */
// llgo:link (*ObjT).ArcSetStartAngle C.lv_arc_set_start_angle
func (recv_ *ObjT) ArcSetStartAngle(start ValuePreciseT) {
}

/**
 * Set the end angle of an arc. 0 deg: right, 90 bottom, etc.
 * @param obj       pointer to an arc object
 * @param end       the end angle  (if `LV_USE_FLOAT` is enabled it can be fractional too.)
 */
// llgo:link (*ObjT).ArcSetEndAngle C.lv_arc_set_end_angle
func (recv_ *ObjT) ArcSetEndAngle(end ValuePreciseT) {
}

/**
 * Set the start and end angles
 * @param obj       pointer to an arc object
 * @param start     the start angle  (if `LV_USE_FLOAT` is enabled it can be fractional too.)
 * @param end       the end angle  (if `LV_USE_FLOAT` is enabled it can be fractional too.)
 */
// llgo:link (*ObjT).ArcSetAngles C.lv_arc_set_angles
func (recv_ *ObjT) ArcSetAngles(start ValuePreciseT, end ValuePreciseT) {
}

/**
 * Set the start angle of an arc background. 0 deg: right, 90 bottom, etc.
 * @param obj       pointer to an arc object
 * @param start     the start angle  (if `LV_USE_FLOAT` is enabled it can be fractional too.)
 */
// llgo:link (*ObjT).ArcSetBgStartAngle C.lv_arc_set_bg_start_angle
func (recv_ *ObjT) ArcSetBgStartAngle(start ValuePreciseT) {
}

/**
 * Set the start angle of an arc background. 0 deg: right, 90 bottom etc.
 * @param obj       pointer to an arc object
 * @param end       the end angle  (if `LV_USE_FLOAT` is enabled it can be fractional too.)
 */
// llgo:link (*ObjT).ArcSetBgEndAngle C.lv_arc_set_bg_end_angle
func (recv_ *ObjT) ArcSetBgEndAngle(end ValuePreciseT) {
}

/**
 * Set the start and end angles of the arc background
 * @param obj       pointer to an arc object
 * @param start     the start angle  (if `LV_USE_FLOAT` is enabled it can be fractional too.)
 * @param end       the end angle  (if `LV_USE_FLOAT` is enabled it can be fractional too.)
 */
// llgo:link (*ObjT).ArcSetBgAngles C.lv_arc_set_bg_angles
func (recv_ *ObjT) ArcSetBgAngles(start ValuePreciseT, end ValuePreciseT) {
}

/**
 * Set the rotation for the whole arc
 * @param obj           pointer to an arc object
 * @param rotation      rotation angle
 */
// llgo:link (*ObjT).ArcSetRotation C.lv_arc_set_rotation
func (recv_ *ObjT) ArcSetRotation(rotation c.Int32T) {
}

/**
 * Set the type of arc.
 * @param obj       pointer to arc object
 * @param type      arc's mode
 */
// llgo:link (*ObjT).ArcSetMode C.lv_arc_set_mode
func (recv_ *ObjT) ArcSetMode(type_ ArcModeT) {
}

/**
 * Set a new value on the arc
 * @param obj       pointer to an arc object
 * @param value     new value
 */
// llgo:link (*ObjT).ArcSetValue C.lv_arc_set_value
func (recv_ *ObjT) ArcSetValue(value c.Int32T) {
}

/**
 * Set minimum and the maximum values of an arc
 * @param obj       pointer to the arc object
 * @param min       minimum value
 * @param max       maximum value
 */
// llgo:link (*ObjT).ArcSetRange C.lv_arc_set_range
func (recv_ *ObjT) ArcSetRange(min c.Int32T, max c.Int32T) {
}

/**
 * Set the minimum values of an arc
 * @param obj       pointer to the arc object
 * @param min       minimum value
 */
// llgo:link (*ObjT).ArcSetMinValue C.lv_arc_set_min_value
func (recv_ *ObjT) ArcSetMinValue(min c.Int32T) {
}

/**
 * Set the maximum values of an arc
 * @param obj       pointer to the arc object
 * @param max       maximum value
 */
// llgo:link (*ObjT).ArcSetMaxValue C.lv_arc_set_max_value
func (recv_ *ObjT) ArcSetMaxValue(max c.Int32T) {
}

/**
 * Set a change rate to limit the speed how fast the arc should reach the pressed point.
 * @param obj       pointer to an arc object
 * @param rate      the change rate
 */
// llgo:link (*ObjT).ArcSetChangeRate C.lv_arc_set_change_rate
func (recv_ *ObjT) ArcSetChangeRate(rate c.Uint32T) {
}

/**
 * Set an offset angle for the knob
 * @param obj       pointer to an arc object
 * @param offset    knob offset from main arc in degrees
 */
// llgo:link (*ObjT).ArcSetKnobOffset C.lv_arc_set_knob_offset
func (recv_ *ObjT) ArcSetKnobOffset(offset c.Int32T) {
}

/**
 * Get the start angle of an arc.
 * @param obj       pointer to an arc object
 * @return          the start angle [0..360]  (if `LV_USE_FLOAT` is enabled it can be fractional too.)
 */
// llgo:link (*ObjT).ArcGetAngleStart C.lv_arc_get_angle_start
func (recv_ *ObjT) ArcGetAngleStart() ValuePreciseT {
	return 0
}

/**
 * Get the end angle of an arc.
 * @param obj       pointer to an arc object
 * @return          the end angle [0..360]  (if `LV_USE_FLOAT` is enabled it can be fractional too.)
 */
// llgo:link (*ObjT).ArcGetAngleEnd C.lv_arc_get_angle_end
func (recv_ *ObjT) ArcGetAngleEnd() ValuePreciseT {
	return 0
}

/**
 * Get the start angle of an arc background.
 * @param obj       pointer to an arc object
 * @return          the  start angle [0..360]  (if `LV_USE_FLOAT` is enabled it can be fractional too.)
 */
// llgo:link (*ObjT).ArcGetBgAngleStart C.lv_arc_get_bg_angle_start
func (recv_ *ObjT) ArcGetBgAngleStart() ValuePreciseT {
	return 0
}

/**
 * Get the end angle of an arc background.
 * @param obj       pointer to an arc object
 * @return          the end angle [0..360]  (if `LV_USE_FLOAT` is enabled it can be fractional too.)
 */
// llgo:link (*ObjT).ArcGetBgAngleEnd C.lv_arc_get_bg_angle_end
func (recv_ *ObjT) ArcGetBgAngleEnd() ValuePreciseT {
	return 0
}

/**
 * Get the value of an arc
 * @param obj       pointer to an arc object
 * @return          the value of the arc
 */
// llgo:link (*ObjT).ArcGetValue C.lv_arc_get_value
func (recv_ *ObjT) ArcGetValue() c.Int32T {
	return 0
}

/**
 * Get the minimum value of an arc
 * @param obj       pointer to an arc object
 * @return          the minimum value of the arc
 */
// llgo:link (*ObjT).ArcGetMinValue C.lv_arc_get_min_value
func (recv_ *ObjT) ArcGetMinValue() c.Int32T {
	return 0
}

/**
 * Get the maximum value of an arc
 * @param obj       pointer to an arc object
 * @return          the maximum value of the arc
 */
// llgo:link (*ObjT).ArcGetMaxValue C.lv_arc_get_max_value
func (recv_ *ObjT) ArcGetMaxValue() c.Int32T {
	return 0
}

/**
 * Get whether the arc is type or not.
 * @param obj       pointer to an arc object
 * @return          arc's mode
 */
// llgo:link (*ObjT).ArcGetMode C.lv_arc_get_mode
func (recv_ *ObjT) ArcGetMode() ArcModeT {
	return 0
}

/**
 * Get the rotation for the whole arc
 * @param obj       pointer to an arc object
 * @return          arc's current rotation
 */
// llgo:link (*ObjT).ArcGetRotation C.lv_arc_get_rotation
func (recv_ *ObjT) ArcGetRotation() c.Int32T {
	return 0
}

/**
 * Get the current knob angle offset
 * @param obj       pointer to an arc object
 * @return          arc's current knob offset
 */
// llgo:link (*ObjT).ArcGetKnobOffset C.lv_arc_get_knob_offset
func (recv_ *ObjT) ArcGetKnobOffset() c.Int32T {
	return 0
}

/**
 * Align an object to the current position of the arc (knob)
 * @param obj           pointer to an arc object
 * @param obj_to_align  pointer to an object to align
 * @param r_offset      consider the radius larger with this value (< 0: for smaller radius)
 */
// llgo:link (*ObjT).ArcAlignObjToAngle C.lv_arc_align_obj_to_angle
func (recv_ *ObjT) ArcAlignObjToAngle(obj_to_align *ObjT, r_offset c.Int32T) {
}

/**
 * Rotate an object to the current position of the arc (knob)
 * @param obj            pointer to an arc object
 * @param obj_to_rotate  pointer to an object to rotate
 * @param r_offset       consider the radius larger with this value (< 0: for smaller radius)
 */
// llgo:link (*ObjT).ArcRotateObjToAngle C.lv_arc_rotate_obj_to_angle
func (recv_ *ObjT) ArcRotateObjToAngle(obj_to_rotate *ObjT, r_offset c.Int32T) {
}

/**
 * Create a tileview object
 * @param parent      pointer to an object, it will be the parent of the new tileview
 * @return            pointer to the created tileview
 */
// llgo:link (*ObjT).TileviewCreate C.lv_tileview_create
func (recv_ *ObjT) TileviewCreate() *ObjT {
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
// llgo:link (*ObjT).TileviewAddTile C.lv_tileview_add_tile
func (recv_ *ObjT) TileviewAddTile(col_id c.Uint8T, row_id c.Uint8T, dir DirT) *ObjT {
	return nil
}

/**
 * Set the active tile in the tileview.
 * @param parent      pointer to the tileview object
 * @param tile_obj    pointer to the tile object to be set as active
 * @param anim_en     animation enable flag (LV_ANIM_ON or LV_ANIM_OFF)
 */
// llgo:link (*ObjT).TileviewSetTile C.lv_tileview_set_tile
func (recv_ *ObjT) TileviewSetTile(tile_obj *ObjT, anim_en AnimEnableT) {
}

/**
 * Set the active tile by index in the tileview
 * @param tv          pointer to the tileview object
 * @param col_id      column id of the tile to be set as active
 * @param row_id      row id of the tile to be set as active
 * @param anim_en     animation enable flag (LV_ANIM_ON or LV_ANIM_OFF)
 */
// llgo:link (*ObjT).TileviewSetTileByIndex C.lv_tileview_set_tile_by_index
func (recv_ *ObjT) TileviewSetTileByIndex(col_id c.Uint32T, row_id c.Uint32T, anim_en AnimEnableT) {
}

/**
 * Get the currently active tile in the tileview
 * @param obj         pointer to the tileview object
 * @return            pointer to the currently active tile object
 */
// llgo:link (*ObjT).TileviewGetTileActive C.lv_tileview_get_tile_active
func (recv_ *ObjT) TileviewGetTileActive() *ObjT {
	return nil
}

/**
 * Create a spinbox object
 * @param parent    pointer to an object, it will be the parent of the new spinbox
 * @return          pointer to the created spinbox
 */
// llgo:link (*ObjT).SpinboxCreate C.lv_spinbox_create
func (recv_ *ObjT) SpinboxCreate() *ObjT {
	return nil
}

/**
 * Set spinbox value
 * @param obj   pointer to spinbox
 * @param v     value to be set
 */
// llgo:link (*ObjT).SpinboxSetValue C.lv_spinbox_set_value
func (recv_ *ObjT) SpinboxSetValue(v c.Int32T) {
}

/**
 * Set spinbox rollover function
 * @param obj       pointer to spinbox
 * @param rollover  true or false to enable or disable (default)
 */
// llgo:link (*ObjT).SpinboxSetRollover C.lv_spinbox_set_rollover
func (recv_ *ObjT) SpinboxSetRollover(rollover bool) {
}

/**
 * Set spinbox digit format (digit count and decimal format)
 * @param obj           pointer to spinbox
 * @param digit_count   number of digit excluding the decimal separator and the sign
 * @param sep_pos       number of digit before the decimal point. If 0, decimal point is not
 * shown
 */
// llgo:link (*ObjT).SpinboxSetDigitFormat C.lv_spinbox_set_digit_format
func (recv_ *ObjT) SpinboxSetDigitFormat(digit_count c.Uint32T, sep_pos c.Uint32T) {
}

/**
 * Set spinbox step
 * @param obj   pointer to spinbox
 * @param step  steps on increment/decrement. Can be 1, 10, 100, 1000, etc the digit that will change.
 */
// llgo:link (*ObjT).SpinboxSetStep C.lv_spinbox_set_step
func (recv_ *ObjT) SpinboxSetStep(step c.Uint32T) {
}

/**
 * Set spinbox value range
 * @param obj       pointer to spinbox
 * @param range_min maximum value, inclusive
 * @param range_max minimum value, inclusive
 */
// llgo:link (*ObjT).SpinboxSetRange C.lv_spinbox_set_range
func (recv_ *ObjT) SpinboxSetRange(range_min c.Int32T, range_max c.Int32T) {
}

/**
 * Set cursor position to a specific digit for edition
 * @param obj   pointer to spinbox
 * @param pos   selected position in spinbox
 */
// llgo:link (*ObjT).SpinboxSetCursorPos C.lv_spinbox_set_cursor_pos
func (recv_ *ObjT) SpinboxSetCursorPos(pos c.Uint32T) {
}

/**
 * Set direction of digit step when clicking an encoder button while in editing mode
 * @param obj           pointer to spinbox
 * @param direction     the direction (LV_DIR_RIGHT or LV_DIR_LEFT)
 */
// llgo:link (*ObjT).SpinboxSetDigitStepDirection C.lv_spinbox_set_digit_step_direction
func (recv_ *ObjT) SpinboxSetDigitStepDirection(direction DirT) {
}

/**
 * Get spinbox rollover function status
 * @param obj   pointer to spinbox
 */
// llgo:link (*ObjT).SpinboxGetRollover C.lv_spinbox_get_rollover
func (recv_ *ObjT) SpinboxGetRollover() bool {
	return false
}

/**
 * Get the spinbox numeral value (user has to convert to float according to its digit format)
 * @param obj   pointer to spinbox
 * @return      value integer value of the spinbox
 */
// llgo:link (*ObjT).SpinboxGetValue C.lv_spinbox_get_value
func (recv_ *ObjT) SpinboxGetValue() c.Int32T {
	return 0
}

/**
 * Get the spinbox step value (user has to convert to float according to its digit format)
 * @param obj   pointer to spinbox
 * @return      value integer step value of the spinbox
 */
// llgo:link (*ObjT).SpinboxGetStep C.lv_spinbox_get_step
func (recv_ *ObjT) SpinboxGetStep() c.Int32T {
	return 0
}

/**
 * Select next lower digit for edition by dividing the step by 10
 * @param obj   pointer to spinbox
 */
// llgo:link (*ObjT).SpinboxStepNext C.lv_spinbox_step_next
func (recv_ *ObjT) SpinboxStepNext() {
}

/**
 * Select next higher digit for edition by multiplying the step by 10
 * @param obj   pointer to spinbox
 */
// llgo:link (*ObjT).SpinboxStepPrev C.lv_spinbox_step_prev
func (recv_ *ObjT) SpinboxStepPrev() {
}

/**
 * Increment spinbox value by one step
 * @param obj   pointer to spinbox
 */
// llgo:link (*ObjT).SpinboxIncrement C.lv_spinbox_increment
func (recv_ *ObjT) SpinboxIncrement() {
}

/**
 * Decrement spinbox value by one step
 * @param obj   pointer to spinbox
 */
// llgo:link (*ObjT).SpinboxDecrement C.lv_spinbox_decrement
func (recv_ *ObjT) SpinboxDecrement() {
}

type SpanOverflowT c.Int

const (
	SPAN_OVERFLOW_CLIP     SpanOverflowT = 0
	SPAN_OVERFLOW_ELLIPSIS SpanOverflowT = 1
	SPAN_OVERFLOW_LAST     SpanOverflowT = 2
)

type SpanModeT c.Int

const (
	SPAN_MODE_FIXED  SpanModeT = 0
	SPAN_MODE_EXPAND SpanModeT = 1
	SPAN_MODE_BREAK  SpanModeT = 2
	SPAN_MODE_LAST   SpanModeT = 3
)

/** Coords of a span */

type X_lvSpanCoordsT struct {
	Heading  AreaT
	Middle   AreaT
	Trailing AreaT
}
type SpanCoordsT X_lvSpanCoordsT

/**********************
 * GLOBAL PROTOTYPES
 **********************/
//go:linkname SpanStackInit C.lv_span_stack_init
func SpanStackInit()

//go:linkname SpanStackDeinit C.lv_span_stack_deinit
func SpanStackDeinit()

/**
 * Create a spangroup object
 * @param parent    pointer to an object, it will be the parent of the new spangroup
 * @return          pointer to the created spangroup
 */
// llgo:link (*ObjT).SpangroupCreate C.lv_spangroup_create
func (recv_ *ObjT) SpangroupCreate() *ObjT {
	return nil
}

/**
 * Create a span string descriptor and add to spangroup.
 * @param obj       pointer to a spangroup object.
 * @return          pointer to the created span.
 */
// llgo:link (*ObjT).SpangroupAddSpan C.lv_spangroup_add_span
func (recv_ *ObjT) SpangroupAddSpan() *SpanT {
	return nil
}

/**
 * Remove the span from the spangroup and free memory.
 * @param obj   pointer to a spangroup object.
 * @param span  pointer to a span.
 */
// llgo:link (*ObjT).SpangroupDeleteSpan C.lv_spangroup_delete_span
func (recv_ *ObjT) SpangroupDeleteSpan(span *SpanT) {
}

/**
 * Set a new text for a span. Memory will be allocated to store the text by the span.
 * As the spangroup is not passed a redraw (invalidation) can't be triggered automatically.
 * Therefore `lv_spangroup_refresh(spangroup)` needs to be called manually,
 * @param span  pointer to a span.
 * @param text  pointer to a text.
 */
// llgo:link (*SpanT).SpanSetText C.lv_span_set_text
func (recv_ *SpanT) SpanSetText(text *c.Char) {
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
// llgo:link (*SpanT).SpanSetTextStatic C.lv_span_set_text_static
func (recv_ *SpanT) SpanSetTextStatic(text *c.Char) {
}

/**
 * Set a new text for a span. Memory will be allocated to store the text by the span.
 * @param obj   pointer to a spangroup widget.
 * @param span  pointer to a span.
 * @param text  pointer to a text.
 */
// llgo:link (*ObjT).SpangroupSetSpanText C.lv_spangroup_set_span_text
func (recv_ *ObjT) SpangroupSetSpanText(span *SpanT, text *c.Char) {
}

/**
 * Set a new text for a span. Memory will be allocated to store the text by the span.
 * @param obj   pointer to a spangroup widget.
 * @param span  pointer to a span.
 * @param text  pointer to a text.
 */
// llgo:link (*ObjT).SpangroupSetSpanTextStatic C.lv_spangroup_set_span_text_static
func (recv_ *ObjT) SpangroupSetSpanTextStatic(span *SpanT, text *c.Char) {
}

/**
 * Copy all style properties of style to the bbuilt-in static style of the span.
 * @param obj       pointer_to a spangroup
 * @param span      pointer to a span.
 * @param style     pointer to a style to copy into the span's built-in style
 */
// llgo:link (*ObjT).SpangroupSetSpanStyle C.lv_spangroup_set_span_style
func (recv_ *ObjT) SpangroupSetSpanStyle(span *SpanT, style *StyleT) {
}

/**
 * DEPRECATED. Use the text_align style property instead
 * Set the align of the spangroup.
 * @param obj   pointer to a spangroup object.
 * @param align see lv_text_align_t for details.
 */
// llgo:link (*ObjT).SpangroupSetAlign C.lv_spangroup_set_align
func (recv_ *ObjT) SpangroupSetAlign(align TextAlignT) {
}

/**
 * Set the overflow of the spangroup.
 * @param obj       pointer to a spangroup object.
 * @param overflow  see lv_span_overflow_t for details.
 */
// llgo:link (*ObjT).SpangroupSetOverflow C.lv_spangroup_set_overflow
func (recv_ *ObjT) SpangroupSetOverflow(overflow SpanOverflowT) {
}

/**
 * Set the indent of the spangroup.
 * @param obj       pointer to a spangroup object.
 * @param indent    the first line indentation
 */
// llgo:link (*ObjT).SpangroupSetIndent C.lv_spangroup_set_indent
func (recv_ *ObjT) SpangroupSetIndent(indent c.Int32T) {
}

/**
 * DEPRECATED, set the width to LV_SIZE_CONTENT or fixed value to control expanding/wrapping"
 * Set the mode of the spangroup.
 * @param obj       pointer to a spangroup object.
 * @param mode      see lv_span_mode_t for details.
 */
// llgo:link (*ObjT).SpangroupSetMode C.lv_spangroup_set_mode
func (recv_ *ObjT) SpangroupSetMode(mode SpanModeT) {
}

/**
 * Set maximum lines of the spangroup.
 * @param obj       pointer to a spangroup object.
 * @param lines     max lines that can be displayed in LV_SPAN_MODE_BREAK mode. < 0 means no limit.
 */
// llgo:link (*ObjT).SpangroupSetMaxLines C.lv_spangroup_set_max_lines
func (recv_ *ObjT) SpangroupSetMaxLines(lines c.Int32T) {
}

/**
 * Get a pointer to the style of a span's built-in style.
 * Any lv_style_set_... functions can be applied on the returned style.
 * @param span  pointer to the span
 * @return      pointer to the style. (valid as long as the span is valid)
 */
// llgo:link (*SpanT).SpanGetStyle C.lv_span_get_style
func (recv_ *SpanT) SpanGetStyle() *StyleT {
	return nil
}

/**
 * Get a pointer to the text of a span
 * @param span  pointer to the span
 * @return      pointer to the text
 */
// llgo:link (*SpanT).SpanGetText C.lv_span_get_text
func (recv_ *SpanT) SpanGetText() *c.Char {
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
// llgo:link (*ObjT).SpangroupGetChild C.lv_spangroup_get_child
func (recv_ *ObjT) SpangroupGetChild(id c.Int32T) *SpanT {
	return nil
}

/**
 * Get number of spans
 * @param obj   the spangroup object to get the child count of.
 * @return      the span count of the spangroup.
 */
// llgo:link (*ObjT).SpangroupGetSpanCount C.lv_spangroup_get_span_count
func (recv_ *ObjT) SpangroupGetSpanCount() c.Uint32T {
	return 0
}

/**
 * Get the align of the spangroup.
 * @param obj   pointer to a spangroup object.
 * @return      the align value.
 */
// llgo:link (*ObjT).SpangroupGetAlign C.lv_spangroup_get_align
func (recv_ *ObjT) SpangroupGetAlign() TextAlignT {
	return 0
}

/**
 * Get the overflow of the spangroup.
 * @param obj   pointer to a spangroup object.
 * @return      the overflow value.
 */
// llgo:link (*ObjT).SpangroupGetOverflow C.lv_spangroup_get_overflow
func (recv_ *ObjT) SpangroupGetOverflow() SpanOverflowT {
	return 0
}

/**
 * Get the indent of the spangroup.
 * @param obj   pointer to a spangroup object.
 * @return      the indent value.
 */
// llgo:link (*ObjT).SpangroupGetIndent C.lv_spangroup_get_indent
func (recv_ *ObjT) SpangroupGetIndent() c.Int32T {
	return 0
}

/**
 * Get the mode of the spangroup.
 * @param obj   pointer to a spangroup object.
 */
// llgo:link (*ObjT).SpangroupGetMode C.lv_spangroup_get_mode
func (recv_ *ObjT) SpangroupGetMode() SpanModeT {
	return 0
}

/**
 * Get maximum lines of the spangroup.
 * @param obj   pointer to a spangroup object.
 * @return      the max lines value.
 */
// llgo:link (*ObjT).SpangroupGetMaxLines C.lv_spangroup_get_max_lines
func (recv_ *ObjT) SpangroupGetMaxLines() c.Int32T {
	return 0
}

/**
 * Get max line height of all span in the spangroup.
 * @param obj   pointer to a spangroup object.
 */
// llgo:link (*ObjT).SpangroupGetMaxLineHeight C.lv_spangroup_get_max_line_height
func (recv_ *ObjT) SpangroupGetMaxLineHeight() c.Int32T {
	return 0
}

/**
 * Get the text content width when all span of spangroup on a line.
 * @param obj       pointer to a spangroup object.
 * @param max_width if text content width >= max_width, return max_width
 * to reduce computation, if max_width == 0, returns the text content width.
 * @return text     content width or max_width.
 */
// llgo:link (*ObjT).SpangroupGetExpandWidth C.lv_spangroup_get_expand_width
func (recv_ *ObjT) SpangroupGetExpandWidth(max_width c.Uint32T) c.Uint32T {
	return 0
}

/**
 * Get the text content height with width fixed.
 * @param obj       pointer to a spangroup object.
 * @param width     the width of the span group.

 */
// llgo:link (*ObjT).SpangroupGetExpandHeight C.lv_spangroup_get_expand_height
func (recv_ *ObjT) SpangroupGetExpandHeight(width c.Int32T) c.Int32T {
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
// llgo:link (*ObjT).SpangroupGetSpanCoords C.lv_spangroup_get_span_coords
func (recv_ *ObjT) SpangroupGetSpanCoords(span *SpanT) SpanCoordsT {
	return SpanCoordsT{}
}

/**
 * Get the span object by point.
 * @param obj       pointer to a spangroup object.
 * @param point     pointer to point containing absolute coordinates
 * @return          pointer to the span under the point or `NULL` if not found.
 */
// llgo:link (*ObjT).SpangroupGetSpanByPoint C.lv_spangroup_get_span_by_point
func (recv_ *ObjT) SpangroupGetSpanByPoint(point *PointT) *SpanT {
	return nil
}

/**
 * Update the mode of the spangroup.
 * @param obj   pointer to a spangroup object.
 */
// llgo:link (*ObjT).SpangroupRefresh C.lv_spangroup_refresh
func (recv_ *ObjT) SpangroupRefresh() {
}

/**
 * Create a tabview widget
 * @param parent    pointer to a parent widget
 * @return          the created tabview
 */
// llgo:link (*ObjT).TabviewCreate C.lv_tabview_create
func (recv_ *ObjT) TabviewCreate() *ObjT {
	return nil
}

/**
 * Add a tab to the tabview
 * @param obj       pointer to a tabview widget
 * @param name      the name of the tab, it will be displayed on the tab bar
 * @return          the widget where the content of the tab can be created
 */
// llgo:link (*ObjT).TabviewAddTab C.lv_tabview_add_tab
func (recv_ *ObjT) TabviewAddTab(name *c.Char) *ObjT {
	return nil
}

/**
 * Change the name of the tab
 * @param obj       pointer to a tabview widget
 * @param idx       the index of the tab to rename
 * @param new_name  the new name as a string
 */
// llgo:link (*ObjT).TabviewRenameTab C.lv_tabview_rename_tab
func (recv_ *ObjT) TabviewRenameTab(idx c.Uint32T, new_name *c.Char) {
}

/**
 * Show a tab
 * @param obj       pointer to a tabview widget
 * @param idx       the index of the tab to show
 * @param anim_en   LV_ANIM_ON/OFF
 */
// llgo:link (*ObjT).TabviewSetActive C.lv_tabview_set_active
func (recv_ *ObjT) TabviewSetActive(idx c.Uint32T, anim_en AnimEnableT) {
}

/**
 * Set the position of the tab bar
 * @param obj       pointer to a tabview widget
 * @param dir       LV_DIR_TOP/BOTTOM/LEFT/RIGHT
 */
// llgo:link (*ObjT).TabviewSetTabBarPosition C.lv_tabview_set_tab_bar_position
func (recv_ *ObjT) TabviewSetTabBarPosition(dir DirT) {
}

/**
 * Set the width or height of the tab bar
 * @param obj       pointer to tabview widget
 * @param size      size of the tab bar in pixels or percentage.
 *                  will be used as width or height based on the position of the tab bar)
 */
// llgo:link (*ObjT).TabviewSetTabBarSize C.lv_tabview_set_tab_bar_size
func (recv_ *ObjT) TabviewSetTabBarSize(size c.Int32T) {
}

/**
 * Get the number of tabs
 * @param obj       pointer to a tabview widget
 * @return          the number of tabs
 */
// llgo:link (*ObjT).TabviewGetTabCount C.lv_tabview_get_tab_count
func (recv_ *ObjT) TabviewGetTabCount() c.Uint32T {
	return 0
}

/**
 * Get the current tab's index
 * @param obj       pointer to a tabview widget
 * @return          the zero based index of the current tab
 */
// llgo:link (*ObjT).TabviewGetTabActive C.lv_tabview_get_tab_active
func (recv_ *ObjT) TabviewGetTabActive() c.Uint32T {
	return 0
}

/**
 * Get the widget where the container of each tab is created
 * @param obj       pointer to a tabview widget
 * @return          the main container widget
 */
// llgo:link (*ObjT).TabviewGetContent C.lv_tabview_get_content
func (recv_ *ObjT) TabviewGetContent() *ObjT {
	return nil
}

/**
 * Get the tab bar where the buttons are created
 * @param obj       pointer to a tabview widget
 * @return          the tab bar
 */
// llgo:link (*ObjT).TabviewGetTabBar C.lv_tabview_get_tab_bar
func (recv_ *ObjT) TabviewGetTabBar() *ObjT {
	return nil
}

/**
 * Initialize the OS layer
 */
//go:linkname OsInit C.lv_os_init
func OsInit()

/**
 * Initialize LVGL library.
 * Should be called before any other LVGL related function.
 */
//go:linkname Init C.lv_init
func Init()

/**
 * Deinit the 'lv' library
 */
//go:linkname Deinit C.lv_deinit
func Deinit()

/**
 * Returns whether the 'lv' library is currently initialized
 */
//go:linkname IsInitialized C.lv_is_initialized
func IsInitialized() bool

// llgo:type C
type AsyncCbT func(c.Pointer)

/**
 * Call an asynchronous function the next time lv_timer_handler() is run. This function is likely to return
 * **before** the call actually happens!
 * @param async_xcb a callback which is the task itself.
 *                 (the 'x' in the argument name indicates that it's not a fully generic function because it not follows
 *                  the `func_name(object, callback, ...)` convention)
 * @param user_data custom parameter
 */
//go:linkname AsyncCall C.lv_async_call
func AsyncCall(async_xcb AsyncCbT, user_data c.Pointer) ResultT

/**
 * Cancel an asynchronous function call
 * @param async_xcb a callback which is the task itself.
 * @param user_data custom parameter
 */
//go:linkname AsyncCallCancel C.lv_async_call_cancel
func AsyncCallCancel(async_xcb AsyncCbT, user_data c.Pointer) ResultT

type X_lvAnimTimelineT struct {
	Unused [8]uint8
}
type AnimTimelineT X_lvAnimTimelineT

/**
 * Create an animation timeline.
 * @return pointer to the animation timeline.
 */
//go:linkname AnimTimelineCreate C.lv_anim_timeline_create
func AnimTimelineCreate() *AnimTimelineT

/**
 * Delete animation timeline.
 * @param at    pointer to the animation timeline.
 */
// llgo:link (*AnimTimelineT).AnimTimelineDelete C.lv_anim_timeline_delete
func (recv_ *AnimTimelineT) AnimTimelineDelete() {
}

/**
 * Add animation to the animation timeline.
 * @param at            pointer to the animation timeline.
 * @param start_time    the time the animation started on the timeline, note that start_time will override the value of delay.
 * @param a             pointer to an animation.
 */
// llgo:link (*AnimTimelineT).AnimTimelineAdd C.lv_anim_timeline_add
func (recv_ *AnimTimelineT) AnimTimelineAdd(start_time c.Uint32T, a *AnimT) {
}

/**
 * Start the animation timeline.
 * @param at    pointer to the animation timeline.
 * @return total time spent in animation timeline.
 */
// llgo:link (*AnimTimelineT).AnimTimelineStart C.lv_anim_timeline_start
func (recv_ *AnimTimelineT) AnimTimelineStart() c.Uint32T {
	return 0
}

/**
 * Pause the animation timeline.
 * @param at    pointer to the animation timeline.
 */
// llgo:link (*AnimTimelineT).AnimTimelinePause C.lv_anim_timeline_pause
func (recv_ *AnimTimelineT) AnimTimelinePause() {
}

/**
 * Set the playback direction of the animation timeline.
 * @param at        pointer to the animation timeline.
 * @param reverse   whether to play in reverse.
 */
// llgo:link (*AnimTimelineT).AnimTimelineSetReverse C.lv_anim_timeline_set_reverse
func (recv_ *AnimTimelineT) AnimTimelineSetReverse(reverse bool) {
}

/**
 * Make the animation timeline repeat itself.
 * @param at        pointer to the animation timeline.
 * @param cnt       repeat count or `LV_ANIM_REPEAT_INFINITE` for infinite repetition. 0: to disable repetition.
 */
// llgo:link (*AnimTimelineT).AnimTimelineSetRepeatCount C.lv_anim_timeline_set_repeat_count
func (recv_ *AnimTimelineT) AnimTimelineSetRepeatCount(cnt c.Uint32T) {
}

/**
 * Set a delay before repeating the animation timeline.
 * @param at        pointer to the animation timeline.
 * @param delay     delay in milliseconds before repeating the animation timeline.
 */
// llgo:link (*AnimTimelineT).AnimTimelineSetRepeatDelay C.lv_anim_timeline_set_repeat_delay
func (recv_ *AnimTimelineT) AnimTimelineSetRepeatDelay(delay c.Uint32T) {
}

/**
 * Set the progress of the animation timeline.
 * @param at        pointer to the animation timeline.
 * @param progress  set value 0~65535 to map 0~100% animation progress.
 */
// llgo:link (*AnimTimelineT).AnimTimelineSetProgress C.lv_anim_timeline_set_progress
func (recv_ *AnimTimelineT) AnimTimelineSetProgress(progress c.Uint16T) {
}

/**
 * Get the time used to play the animation timeline.
 * @param at    pointer to the animation timeline.
 * @return total time spent in animation timeline.
 */
// llgo:link (*AnimTimelineT).AnimTimelineGetPlaytime C.lv_anim_timeline_get_playtime
func (recv_ *AnimTimelineT) AnimTimelineGetPlaytime() c.Uint32T {
	return 0
}

/**
 * Get whether the animation timeline is played in reverse.
 * @param at    pointer to the animation timeline.
 * @return return true if it is reverse playback.
 */
// llgo:link (*AnimTimelineT).AnimTimelineGetReverse C.lv_anim_timeline_get_reverse
func (recv_ *AnimTimelineT) AnimTimelineGetReverse() bool {
	return false
}

/**
 * Get the progress of the animation timeline.
 * @param at    pointer to the animation timeline.
 * @return return value 0~65535 to map 0~100% animation progress.
 */
// llgo:link (*AnimTimelineT).AnimTimelineGetProgress C.lv_anim_timeline_get_progress
func (recv_ *AnimTimelineT) AnimTimelineGetProgress() c.Uint16T {
	return 0
}

/**
 * Get repeat count of the animation timeline.
 * @param at    pointer to the animation timeline.
 */
// llgo:link (*AnimTimelineT).AnimTimelineGetRepeatCount C.lv_anim_timeline_get_repeat_count
func (recv_ *AnimTimelineT) AnimTimelineGetRepeatCount() c.Uint32T {
	return 0
}

/**
 * Get repeat delay of the animation timeline.
 * @param at    pointer to the animation timeline.
 */
// llgo:link (*AnimTimelineT).AnimTimelineGetRepeatDelay C.lv_anim_timeline_get_repeat_delay
func (recv_ *AnimTimelineT) AnimTimelineGetRepeatDelay() c.Uint32T {
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
//go:linkname UtilsBsearch C.lv_utils_bsearch
func UtilsBsearch(key c.Pointer, base c.Pointer, n c.SizeT, size c.SizeT, cmp func(c.Pointer, c.Pointer) c.Int) c.Pointer

/**
 * Save a draw buf to a file
 * @param draw_buf  pointer to a draw buffer
 * @param path      path to the file to save
 * @return          LV_RESULT_OK: success; LV_RESULT_INVALID: error
 */
// llgo:link (*DrawBufT).DrawBufSaveToFile C.lv_draw_buf_save_to_file
func (recv_ *DrawBufT) DrawBufSaveToFile(path *c.Char) ResultT {
	return 0
}

// llgo:type C
type IterNextCb func(c.Pointer, c.Pointer, c.Pointer) ResultT

// llgo:type C
type IterInspectCb func(c.Pointer)

/**
 * Create an iterator based on an instance, and then the next element of the iterator can be obtained through lv_iter_next,
 * In order to obtain the next operation in a unified and abstract way.
 * @param instance       The instance to be iterated
 * @param elem_size      The size of the element to be iterated in bytes
 * @param context_size   The size of the context to be passed to the next_cb in bytes
 * @param next_cb        The callback function to get the next element
 * @return               The iterator object
 */
//go:linkname IterCreate C.lv_iter_create
func IterCreate(instance c.Pointer, elem_size c.Uint32T, context_size c.Uint32T, next_cb IterNextCb) *IterT

/**
 * Get the context of the iterator. You can use it to store some temporary variables associated with current iterator..
 * @param iter           `lv_iter_t` object create before
 * @return the iter context
 */
// llgo:link (*IterT).IterGetContext C.lv_iter_get_context
func (recv_ *IterT) IterGetContext() c.Pointer {
	return nil
}

/**
 * Destroy the iterator object, and release the context. Other resources allocated by the user are not released.
 * The user needs to release it by itself.
 * @param iter          `lv_iter_t` object create before
 */
// llgo:link (*IterT).IterDestroy C.lv_iter_destroy
func (recv_ *IterT) IterDestroy() {
}

/**
 * Get the next element of the iterator.
 * @param iter          `lv_iter_t` object create before
 * @param elem          The pointer to store the next element
 * @return              LV_RESULT_OK: Get the next element successfully
 *                      LV_RESULT_INVALID: The next element is invalid
 */
// llgo:link (*IterT).IterNext C.lv_iter_next
func (recv_ *IterT) IterNext(elem c.Pointer) ResultT {
	return 0
}

/**
 * Make the iterator peekable, which means that the user can peek the next element without advancing the iterator.
 * @param iter          `lv_iter_t` object create before
 * @param capacity      The capacity of the peek buffer
 */
// llgo:link (*IterT).IterMakePeekable C.lv_iter_make_peekable
func (recv_ *IterT) IterMakePeekable(capacity c.Uint32T) {
}

/**
 * Peek the next element of the iterator without advancing the iterator.
 * @param iter          `lv_iter_t` object create before
 * @param elem          The pointer to store the next element
 * @return              LV_RESULT_OK: Peek the next element successfully
 *                      LV_RESULT_INVALID: The next element is invalid
 */
// llgo:link (*IterT).IterPeek C.lv_iter_peek
func (recv_ *IterT) IterPeek(elem c.Pointer) ResultT {
	return 0
}

/**
 * Only advance the iterator without getting the next element.
 * @param iter          `lv_iter_t` object create before
 * @return              LV_RESULT_OK: Peek the next element successfully
 *                      LV_RESULT_INVALID: The next element is invalid
 */
// llgo:link (*IterT).IterPeekAdvance C.lv_iter_peek_advance
func (recv_ *IterT) IterPeekAdvance() ResultT {
	return 0
}

/**
 * Reset the peek cursor to the `next` cursor.
 * @param iter          `lv_iter_t` object create before
 * @return              LV_RESULT_OK: Reset the peek buffer successfully
 *                      LV_RESULT_INVALID: The peek buffer is invalid
 */
// llgo:link (*IterT).IterPeekReset C.lv_iter_peek_reset
func (recv_ *IterT) IterPeekReset() ResultT {
	return 0
}

/**
 * Inspect the element of the iterator. The callback function will be called for each element of the iterator.
 * @param iter          `lv_iter_t` object create before
 * @param inspect_cb    The callback function to inspect the element
 */
// llgo:link (*IterT).IterInspect C.lv_iter_inspect
func (recv_ *IterT) IterInspect(inspect_cb IterInspectCb) {
}

// llgo:type C
type CircleBufFillCbT func(c.Pointer, c.Uint32T, c.Int32T, c.Pointer) bool

/**
 * Create a circle buffer
 * @param capacity the maximum number of elements in the buffer
 * @param element_size the size of an element in bytes
 * @return pointer to the created buffer
 */
//go:linkname CircleBufCreate C.lv_circle_buf_create
func CircleBufCreate(capacity c.Uint32T, element_size c.Uint32T) *CircleBufT

/**
 * Create a circle buffer from an existing buffer
 * @param buf pointer to a buffer
 * @param capacity the maximum number of elements in the buffer
 * @param element_size the size of an element in bytes
 * @return pointer to the created buffer
 */
//go:linkname CircleBufCreateFromBuf C.lv_circle_buf_create_from_buf
func CircleBufCreateFromBuf(buf c.Pointer, capacity c.Uint32T, element_size c.Uint32T) *CircleBufT

/**
 * Create a circle buffer from an existing array
 * @param array pointer to an array
 * @return pointer to the created buffer
 */
// llgo:link (*ArrayT).CircleBufCreateFromArray C.lv_circle_buf_create_from_array
func (recv_ *ArrayT) CircleBufCreateFromArray() *CircleBufT {
	return nil
}

/**
 * Resize the buffer
 * @param circle_buf pointer to a buffer
 * @param capacity the new capacity of the buffer
 * @return LV_RESULT_OK: the buffer is resized; LV_RESULT_INVALID: the buffer is not resized
 */
// llgo:link (*CircleBufT).CircleBufResize C.lv_circle_buf_resize
func (recv_ *CircleBufT) CircleBufResize(capacity c.Uint32T) ResultT {
	return 0
}

/**
 * Destroy a circle buffer
 * @param circle_buf pointer to buffer
 */
// llgo:link (*CircleBufT).CircleBufDestroy C.lv_circle_buf_destroy
func (recv_ *CircleBufT) CircleBufDestroy() {
}

/**
 * Get the size of the buffer
 * @param circle_buf pointer to buffer
 * @return the number of elements in the buffer
 */
// llgo:link (*CircleBufT).CircleBufSize C.lv_circle_buf_size
func (recv_ *CircleBufT) CircleBufSize() c.Uint32T {
	return 0
}

/**
 * Get the capacity of the buffer
 * @param circle_buf pointer to buffer
 * @return the maximum number of elements in the buffer
 */
// llgo:link (*CircleBufT).CircleBufCapacity C.lv_circle_buf_capacity
func (recv_ *CircleBufT) CircleBufCapacity() c.Uint32T {
	return 0
}

/**
 * Get the remaining space in the buffer
 * @param circle_buf pointer to buffer
 * @return the number of elements that can be written to the buffer
 */
// llgo:link (*CircleBufT).CircleBufRemain C.lv_circle_buf_remain
func (recv_ *CircleBufT) CircleBufRemain() c.Uint32T {
	return 0
}

/**
 * Check if the buffer is empty
 * @param circle_buf pointer to buffer
 * @return true: the buffer is empty; false: the buffer is not empty
 */
// llgo:link (*CircleBufT).CircleBufIsEmpty C.lv_circle_buf_is_empty
func (recv_ *CircleBufT) CircleBufIsEmpty() bool {
	return false
}

/**
 * Check if the buffer is full
 * @param circle_buf pointer to buffer
 * @return true: the buffer is full; false: the buffer is not full
 */
// llgo:link (*CircleBufT).CircleBufIsFull C.lv_circle_buf_is_full
func (recv_ *CircleBufT) CircleBufIsFull() bool {
	return false
}

/**
 * Reset the buffer
 * @param circle_buf pointer to buffer
 * @return LV_RESULT_OK: the buffer is reset; LV_RESULT_INVALID: the buffer is not reset
 */
// llgo:link (*CircleBufT).CircleBufReset C.lv_circle_buf_reset
func (recv_ *CircleBufT) CircleBufReset() {
}

/**
 * Get the head of the buffer
 * @param circle_buf pointer to buffer
 * @return pointer to the head of the buffer
 */
// llgo:link (*CircleBufT).CircleBufHead C.lv_circle_buf_head
func (recv_ *CircleBufT) CircleBufHead() c.Pointer {
	return nil
}

/**
 * Get the tail of the buffer
 * @param circle_buf pointer to buffer
 * @return pointer to the tail of the buffer
 */
// llgo:link (*CircleBufT).CircleBufTail C.lv_circle_buf_tail
func (recv_ *CircleBufT) CircleBufTail() c.Pointer {
	return nil
}

/**
 * Read a value
 * @param circle_buf pointer to buffer
 * @param data pointer to a variable to store the read value
 * @return LV_RESULT_OK: the value is read; LV_RESULT_INVALID: the value is not read
 */
// llgo:link (*CircleBufT).CircleBufRead C.lv_circle_buf_read
func (recv_ *CircleBufT) CircleBufRead(data c.Pointer) ResultT {
	return 0
}

/**
 * Write a value
 * @param circle_buf pointer to buffer
 * @param data pointer to the value to write
 * @return LV_RESULT_OK: the value is written; LV_RESULT_INVALID: the value is not written
 */
// llgo:link (*CircleBufT).CircleBufWrite C.lv_circle_buf_write
func (recv_ *CircleBufT) CircleBufWrite(data c.Pointer) ResultT {
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
// llgo:link (*CircleBufT).CircleBufFill C.lv_circle_buf_fill
func (recv_ *CircleBufT) CircleBufFill(count c.Uint32T, fill_cb CircleBufFillCbT, user_data c.Pointer) c.Uint32T {
	return 0
}

/**
 * Skip a value
 * @param circle_buf pointer to buffer
 * @return LV_RESULT_OK: the value is skipped; LV_RESULT_INVALID: the value is not skipped
 */
// llgo:link (*CircleBufT).CircleBufSkip C.lv_circle_buf_skip
func (recv_ *CircleBufT) CircleBufSkip() ResultT {
	return 0
}

/**
 * Peek a value
 * @param circle_buf pointer to buffer
 * @param data pointer to a variable to store the peeked value
 * @return LV_RESULT_OK: the value is peeked; LV_RESULT_INVALID: the value is not peeked
 */
// llgo:link (*CircleBufT).CircleBufPeek C.lv_circle_buf_peek
func (recv_ *CircleBufT) CircleBufPeek(data c.Pointer) ResultT {
	return 0
}

/**
 * Peek a value at an index
 * @param circle_buf pointer to buffer
 * @param index the index of the value to peek, if the index is greater than the size of the buffer, it will return looply.
 * @param data pointer to a variable to store the peeked value
 * @return LV_RESULT_OK: the value is peeked; LV_RESULT_INVALID: the value is not peeked
 */
// llgo:link (*CircleBufT).CircleBufPeekAt C.lv_circle_buf_peek_at
func (recv_ *CircleBufT) CircleBufPeekAt(index c.Uint32T, data c.Pointer) ResultT {
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
type TreeClassT X_lvTreeClassT
type TreeNodeT X_lvTreeNodeT
type X_lvTreeWalkModeT c.Int

const (
	TREE_WALK_PRE_ORDER  X_lvTreeWalkModeT = 0
	TREE_WALK_POST_ORDER X_lvTreeWalkModeT = 1
)

type TreeWalkModeT c.Uint8T

// llgo:type C
type TreeTraverseCbT func(*TreeNodeT, c.Pointer) bool

// llgo:type C
type TreeBeforeCbT func(*TreeNodeT, c.Pointer) bool

// llgo:type C
type TreeAfterCbT func(*TreeNodeT, c.Pointer)

/**
 * @brief Create a tree node
 * @param class_p pointer to a class of the node
 * @param parent pointer to the parent node (or NULL if it's the root node)
 * @return pointer to the new node
 */
// llgo:link (*TreeClassT).TreeNodeCreate C.lv_tree_node_create
func (recv_ *TreeClassT) TreeNodeCreate(parent *TreeNodeT) *TreeNodeT {
	return nil
}

/**
 * @brief Delete a tree node and all its children recursively
 * @param node pointer to the node to delete
 */
// llgo:link (*TreeNodeT).TreeNodeDelete C.lv_tree_node_delete
func (recv_ *TreeNodeT) TreeNodeDelete() {
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
// llgo:link (*TreeNodeT).TreeWalk C.lv_tree_walk
func (recv_ *TreeNodeT) TreeWalk(mode TreeWalkModeT, cb TreeTraverseCbT, bcb TreeBeforeCbT, acb TreeAfterCbT, user_data c.Pointer) bool {
	return false
}

/**********************
 *      TYPEDEFS
 **********************/

type BinfontFontSrcT struct {
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
//go:linkname BinfontCreate C.lv_binfont_create
func BinfontCreate(path *c.Char) *FontT

/**
 * Frees the memory allocated by the `lv_binfont_create()` function
 * @param font          lv_font_t object created by the lv_binfont_create function
 */
// llgo:link (*FontT).BinfontDestroy C.lv_binfont_destroy
func (recv_ *FontT) BinfontDestroy() {
}

type ArclabelDirT c.Int

const (
	ARCLABEL_DIR_CLOCKWISE         ArclabelDirT = 0
	ARCLABEL_DIR_COUNTER_CLOCKWISE ArclabelDirT = 1
)

type ArclabelTextAlignT c.Int

const (
	ARCLABEL_TEXT_ALIGN_DEFAULT  ArclabelTextAlignT = 0
	ARCLABEL_TEXT_ALIGN_LEADING  ArclabelTextAlignT = 1
	ARCLABEL_TEXT_ALIGN_CENTER   ArclabelTextAlignT = 2
	ARCLABEL_TEXT_ALIGN_TRAILING ArclabelTextAlignT = 3
)

/**
 * Create an arc label object
 * @param parent    pointer to an object, it will be the parent of the new arc label
 * @return          pointer to the created arc label
 */
// llgo:link (*ObjT).ArclabelCreate C.lv_arclabel_create
func (recv_ *ObjT) ArclabelCreate() *ObjT {
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
// llgo:link (*ObjT).ArclabelSetText C.lv_arclabel_set_text
func (recv_ *ObjT) ArclabelSetText(text *c.Char) {
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
// llgo:link (*ObjT).ArclabelSetTextFmt C.lv_arclabel_set_text_fmt
func (recv_ *ObjT) ArclabelSetTextFmt(fmt *c.Char, __llgo_va_list ...interface{}) {
}

/**
 * Sets a new static text for the arc label or refreshes it with the current text.
 * The 'text' must remain valid in memory; the arc label does not manage its lifecycle.
 *
 * @param obj       Pointer to the arc label object.
 * @param text      Pointer to the new text. If NULL, the label is refreshed with its current text.
 */
// llgo:link (*ObjT).ArclabelSetTextStatic C.lv_arclabel_set_text_static
func (recv_ *ObjT) ArclabelSetTextStatic(text *c.Char) {
}

/**
 * Set the start angle of an arc. 0 deg: right, 90 bottom, etc.
 * @param obj       pointer to an arc label object
 * @param start     the start angle. (if `LV_USE_FLOAT` is enabled it can be fractional too.)
 */
// llgo:link (*ObjT).ArclabelSetAngleStart C.lv_arclabel_set_angle_start
func (recv_ *ObjT) ArclabelSetAngleStart(start ValuePreciseT) {
}

/**
 * Set the end angle of an arc. 0 deg: right, 90 bottom, etc.
 * @param obj       pointer to an arc label object
 * @param size      the angle size (if `LV_USE_FLOAT` is enabled it can be fractional too.)
 */
// llgo:link (*ObjT).ArclabelSetAngleSize C.lv_arclabel_set_angle_size
func (recv_ *ObjT) ArclabelSetAngleSize(size ValuePreciseT) {
}

/**
 * Set the rotation for the whole arc
 * @param obj       pointer to an arc label object
 * @param offset    rotation angle
 */
// llgo:link (*ObjT).ArclabelSetOffset C.lv_arclabel_set_offset
func (recv_ *ObjT) ArclabelSetOffset(offset c.Int32T) {
}

/**
 * Set the type of arc.
 * @param obj       pointer to and arc label object
 * @param dir      arc label's direction
 */
// llgo:link (*ObjT).ArclabelSetDir C.lv_arclabel_set_dir
func (recv_ *ObjT) ArclabelSetDir(dir ArclabelDirT) {
}

/**
 * Enable the recoloring by in-line commands
 * @param obj       pointer to an arc label object
 * @param en        true: enable recoloring, false: disable
 * Example: "This is a #ff0000 red# word"
 */
// llgo:link (*ObjT).ArclabelSetRecolor C.lv_arclabel_set_recolor
func (recv_ *ObjT) ArclabelSetRecolor(en bool) {
}

/**
 * Set the radius for an arc label object.
 *
 * @param obj       pointer to the arc label object.
 * @param radius    The radius value to set for the label's curvature, in pixels.
 */
// llgo:link (*ObjT).ArclabelSetRadius C.lv_arclabel_set_radius
func (recv_ *ObjT) ArclabelSetRadius(radius c.Uint32T) {
}

/**
 * Set the center offset x for an arc label object.
 * @param obj       pointer to an arc label object
 * @param x         the x offset
 */
// llgo:link (*ObjT).ArclabelSetCenterOffsetX C.lv_arclabel_set_center_offset_x
func (recv_ *ObjT) ArclabelSetCenterOffsetX(x c.Uint32T) {
}

/**
 * Set the center offset y for an arc label object.
 * @param obj       pointer to an arc label object
 * @param y         the y offset
 */
// llgo:link (*ObjT).ArclabelSetCenterOffsetY C.lv_arclabel_set_center_offset_y
func (recv_ *ObjT) ArclabelSetCenterOffsetY(y c.Uint32T) {
}

/**
 * Set the text vertical alignment for an arc label object.
 * @param obj       pointer to an arc label object
 * @param align     the vertical alignment
 */
// llgo:link (*ObjT).ArclabelSetTextVerticalAlign C.lv_arclabel_set_text_vertical_align
func (recv_ *ObjT) ArclabelSetTextVerticalAlign(align ArclabelTextAlignT) {
}

/**
 * Set the text horizontal alignment for an arc label object.
 * @param obj       pointer to an arc label object
 * @param align     the horizontal alignment
 */
// llgo:link (*ObjT).ArclabelSetTextHorizontalAlign C.lv_arclabel_set_text_horizontal_align
func (recv_ *ObjT) ArclabelSetTextHorizontalAlign(align ArclabelTextAlignT) {
}

/**
 * Get the start angle of an arc label.
 * @param obj       pointer to an arc label object
 * @return          the start angle [0..360]  (if `LV_USE_FLOAT` is enabled it can be fractional too.)
 */
// llgo:link (*ObjT).ArclabelGetAngleStart C.lv_arclabel_get_angle_start
func (recv_ *ObjT) ArclabelGetAngleStart() ValuePreciseT {
	return 0
}

/**
 * Get the angle size of an arc label.
 * @param obj       pointer to an arc label object
 * @return          the end angle [0..360]  (if `LV_USE_FLOAT` is enabled it can be fractional too.)
 */
// llgo:link (*ObjT).ArclabelGetAngleSize C.lv_arclabel_get_angle_size
func (recv_ *ObjT) ArclabelGetAngleSize() ValuePreciseT {
	return 0
}

/**
 * Get whether the arc label is type or not.
 * @param obj       pointer to an arc label object
 * @return          arc label's direction
 */
// llgo:link (*ObjT).ArclabelGetDir C.lv_arclabel_get_dir
func (recv_ *ObjT) ArclabelGetDir() ArclabelDirT {
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
// llgo:link (*ObjT).ArclabelGetRecolor C.lv_arclabel_get_recolor
func (recv_ *ObjT) ArclabelGetRecolor() bool {
	return false
}

/**
 * Get the text of the arc label.
 * @param obj       pointer to an arc label object
 * @return          the radius of the arc label
 */
// llgo:link (*ObjT).ArclabelGetRadius C.lv_arclabel_get_radius
func (recv_ *ObjT) ArclabelGetRadius() c.Uint32T {
	return 0
}

/**
 * Get the center offset x for an arc label object.
 * @param obj       pointer to an arc label object
 * @return          the x offset
 */
// llgo:link (*ObjT).ArclabelGetCenterOffsetX C.lv_arclabel_get_center_offset_x
func (recv_ *ObjT) ArclabelGetCenterOffsetX() c.Uint32T {
	return 0
}

/**
 * Get the center offset y for an arc label object.
 * @param obj       pointer to an arc label object
 * @return          the y offset
 */
// llgo:link (*ObjT).ArclabelGetCenterOffsetY C.lv_arclabel_get_center_offset_y
func (recv_ *ObjT) ArclabelGetCenterOffsetY() c.Uint32T {
	return 0
}

/**
 * Get the text vertical alignment for an arc label object.
 * @param obj       pointer to an arc label object
 * @return          the vertical alignment
 */
// llgo:link (*ObjT).ArclabelGetTextVerticalAlign C.lv_arclabel_get_text_vertical_align
func (recv_ *ObjT) ArclabelGetTextVerticalAlign() ArclabelTextAlignT {
	return 0
}

/**
 * Get the text horizontal alignment for an arc label object.
 * @param obj       pointer to an arc label object
 * @return          the horizontal alignment
 */
// llgo:link (*ObjT).ArclabelGetTextHorizontalAlign C.lv_arclabel_get_text_horizontal_align
func (recv_ *ObjT) ArclabelGetTextHorizontalAlign() ArclabelTextAlignT {
	return 0
}

/**
 * Create a list object
 * @param parent    pointer to an object, it will be the parent of the new list
 * @return          pointer to the created list
 */
// llgo:link (*ObjT).ListCreate C.lv_list_create
func (recv_ *ObjT) ListCreate() *ObjT {
	return nil
}

/**
 * Add text to a list
 * @param list      pointer to a list, it will be the parent of the new label
 * @param txt       text of the new label
 * @return          pointer to the created label
 */
// llgo:link (*ObjT).ListAddText C.lv_list_add_text
func (recv_ *ObjT) ListAddText(txt *c.Char) *ObjT {
	return nil
}

/**
 * Add button to a list
 * @param list      pointer to a list, it will be the parent of the new button
 * @param icon      icon for the button, when NULL it will have no icon
 * @param txt       text of the new button, when NULL no text will be added
 * @return          pointer to the created button
 */
// llgo:link (*ObjT).ListAddButton C.lv_list_add_button
func (recv_ *ObjT) ListAddButton(icon c.Pointer, txt *c.Char) *ObjT {
	return nil
}

/**
 * Get text of a given list button
 * @param list      pointer to a list
 * @param btn       pointer to the button
 * @return          text of btn, if btn doesn't have text "" will be returned
 */
// llgo:link (*ObjT).ListGetButtonText C.lv_list_get_button_text
func (recv_ *ObjT) ListGetButtonText(btn *ObjT) *c.Char {
	return nil
}

/**
 * Set text of a given list button
 * @param list      pointer to a list
 * @param btn       pointer to the button
 * @param txt       pointer to the text
 */
// llgo:link (*ObjT).ListSetButtonText C.lv_list_set_button_text
func (recv_ *ObjT) ListSetButtonText(btn *ObjT, txt *c.Char) {
}

/**
 * Create a spinner widget
 * @param parent    pointer to an object, it will be the parent of the new spinner.
 * @return          the created spinner
 */
// llgo:link (*ObjT).SpinnerCreate C.lv_spinner_create
func (recv_ *ObjT) SpinnerCreate() *ObjT {
	return nil
}

/**
 * Set the animation time and arc length of the spinner
 * @param obj       pointer to a spinner
 * @param t         the animation time in milliseconds
 * @param angle     the angle of the arc in degrees
 */
// llgo:link (*ObjT).SpinnerSetAnimParams C.lv_spinner_set_anim_params
func (recv_ *ObjT) SpinnerSetAnimParams(t c.Uint32T, angle c.Uint32T) {
}

/**
 * Initialize the binary image decoder module
 */
//go:linkname BinDecoderInit C.lv_bin_decoder_init
func BinDecoderInit()

/**
 * Get info about a lvgl binary image
 * @param decoder the decoder where this function belongs
 * @param dsc image descriptor containing the source and type of the image and other info.
 * @param header store the image data here
 * @return LV_RESULT_OK: the info is successfully stored in `header`; LV_RESULT_INVALID: unknown format or other error.
 */
// llgo:link (*ImageDecoderT).BinDecoderInfo C.lv_bin_decoder_info
func (recv_ *ImageDecoderT) BinDecoderInfo(dsc *ImageDecoderDscT, header *ImageHeaderT) ResultT {
	return 0
}

// llgo:link (*ImageDecoderT).BinDecoderGetArea C.lv_bin_decoder_get_area
func (recv_ *ImageDecoderT) BinDecoderGetArea(dsc *ImageDecoderDscT, full_area *AreaT, decoded_area *AreaT) ResultT {
	return 0
}

/**
 * Open a lvgl binary image
 * @param decoder the decoder where this function belongs
 * @param dsc pointer to decoder descriptor. `src`, `style` are already initialized in it.
 * @return LV_RESULT_OK: the info is successfully stored in `header`; LV_RESULT_INVALID: unknown format or other error.
 */
// llgo:link (*ImageDecoderT).BinDecoderOpen C.lv_bin_decoder_open
func (recv_ *ImageDecoderT) BinDecoderOpen(dsc *ImageDecoderDscT) ResultT {
	return 0
}

/**
 * Close the pending decoding. Free resources etc.
 * @param decoder pointer to the decoder the function associated with
 * @param dsc pointer to decoder descriptor
 */
// llgo:link (*ImageDecoderT).BinDecoderClose C.lv_bin_decoder_close
func (recv_ *ImageDecoderT) BinDecoderClose(dsc *ImageDecoderDscT) {
}

/**********************
 * GLOBAL PROTOTYPES
 **********************/
//go:linkname SdlWindowCreate C.lv_sdl_window_create
func SdlWindowCreate(hor_res c.Int32T, ver_res c.Int32T) *DisplayT

// llgo:link (*DisplayT).SdlWindowSetResizeable C.lv_sdl_window_set_resizeable
func (recv_ *DisplayT) SdlWindowSetResizeable(value bool) {
}

// llgo:link (*DisplayT).SdlWindowSetZoom C.lv_sdl_window_set_zoom
func (recv_ *DisplayT) SdlWindowSetZoom(zoom c.Float) {
}

// llgo:link (*DisplayT).SdlWindowGetZoom C.lv_sdl_window_get_zoom
func (recv_ *DisplayT) SdlWindowGetZoom() c.Float {
	return 0
}

// llgo:link (*DisplayT).SdlWindowSetTitle C.lv_sdl_window_set_title
func (recv_ *DisplayT) SdlWindowSetTitle(title *c.Char) {
}

// llgo:link (*DisplayT).SdlWindowSetIcon C.lv_sdl_window_set_icon
func (recv_ *DisplayT) SdlWindowSetIcon(icon c.Pointer, width c.Int32T, height c.Int32T) {
}

// llgo:link (*DisplayT).SdlWindowGetRenderer C.lv_sdl_window_get_renderer
func (recv_ *DisplayT) SdlWindowGetRenderer() c.Pointer {
	return nil
}

//go:linkname SdlQuit C.lv_sdl_quit
func SdlQuit()

/**********************
 * GLOBAL PROTOTYPES
 **********************/
//go:linkname SdlMouseCreate C.lv_sdl_mouse_create
func SdlMouseCreate() *IndevT

/**********************
 * GLOBAL PROTOTYPES
 **********************/
//go:linkname SdlMousewheelCreate C.lv_sdl_mousewheel_create
func SdlMousewheelCreate() *IndevT

/**********************
 * GLOBAL PROTOTYPES
 **********************/
//go:linkname SdlKeyboardCreate C.lv_sdl_keyboard_create
func SdlKeyboardCreate() *IndevT

type CoordT c.Int32T
type ResT ResultT
type ImgDscT ImageDscT
type DispT DisplayT
type DispRotationT DisplayRotationT
type DispRenderT DisplayRenderModeT
type AnimReadyCbT AnimCompletedCbT
type ScrLoadAnimT ScreenLoadAnimT
type BtnmatrixCtrlT ButtonmatrixCtrlT
