package types

type TModalResult = int32

// 常用值请见 types/colors 包
type TColor uint32

type THelpEventData = uintptr

type TTabOrder = int16

type PFNLVCOMPARE uintptr

type PFNTVCOMPARE uintptr

type Integer = int32

type Cardinal = uint32

type Single = float32

type Word = uint16

type Byte = uint8

type TFontCharset = uint8

// Unicode
type Char = uint16

type TClass uintptr

type TGridCoord struct {
	X int32
	Y int32
}

type TCustomData = uintptr

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

// LibResouces
type TLibResouce struct {
	Name string
	Ptr  uintptr
}

// TConstraintSize = 0..MaxInt;
type TConstraintSize int32

// TColor
func (c TColor) R() byte {
	return byte(c)
}

func (c TColor) G() byte {
	return byte(c >> 8)
}

func (c TColor) B() byte {
	return byte(c >> 16)
}

func (c TColor) RGB(r, g, b byte) TColor {
	return TColor(uint32(r) | (uint32(g) << 8) | (uint32(b) << 16))
}
