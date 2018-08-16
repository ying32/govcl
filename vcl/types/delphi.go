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
