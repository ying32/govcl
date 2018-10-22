package types

type TModalResult int32

type TColor uint32

type THelpEventData uintptr

type TTabOrder int16

type PFNLVCOMPARE uintptr

type PFNTVCOMPARE uintptr

type Integer int32

type Cardinal uint32

type Single float32

type Word uint16

type Byte uint8

type TFontCharset byte

// Unicode
type Char uint16

type TClass uintptr

type TGridCoord struct {
	X int32
	Y int32
}

type TCustomData uintptr

type TGridRect TRect

// 伪造，实际为一个接口类型
type IObjectArray uintptr

type TSysLocale struct {
	//Delphi compat fields
	DefaultLCID int32
	PriLangID   int32
	SubLangID   int32

	// win32 names
	FarEast    bool
	MiddleEast bool

	// LCL
	// real meaning  2: (MBCS: boolean; RightToLeft: Boolean);
}

type TSmallPoint struct {
	X int16
	Y int16
}

type TGUID struct {
	D1 uint32
	D2 uint16
	D3 uint16
	D4 [8]uint8
}
