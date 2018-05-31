package types

type TPoint struct {
	X, Y int32
}

type TRect struct {
	Left, Top, Right, Bottom int32
}

type TSize struct {
	Cx, Cy int32
}

type HWND uintptr

type HBITMAP uintptr

type TModalResult int32

type HMENU uintptr

type HICON uintptr

type HDC uintptr

type TColor uint32

type THelpEventData uintptr

type TTabOrder int16

type HFONT uintptr

type HBRUSH uintptr

type HPEN uintptr

type TFontCharset uint8

type HKEY uintptr

type HMONITOR uintptr

type Char uint16

type PFNLVCOMPARE uintptr

type PFNTVCOMPARE uintptr

type HGDIOBJ uintptr

//----------------------------------------------------------------------------------------------------------------------
// -- TRect

func (r *TRect) PtInRect(P TPoint) bool {
	return P.X >= r.Left && P.X < r.Right && P.Y >= r.Top && P.Y < r.Bottom
}

func (r *TRect) Width() int32 {
	return r.Right - r.Left
}

func (r *TRect) Height() int32 {
	return r.Bottom - r.Top
}

func (r *TRect) IsEmpty() bool {
	return r.Right <= r.Left || r.Bottom <= r.Top
}

func (r *TRect) Contains(aR TRect) bool {
	return r.Left <= aR.Left && r.Right >= aR.Right && r.Top <= aR.Top && r.Bottom >= aR.Bottom
}

func (r *TRect) IntersectsWith(aR TRect) bool {
	return r.Left < aR.Right && r.Right > aR.Left && r.Top < aR.Bottom && r.Bottom > aR.Top
}

func (r *TRect) CenterPoint() (ret TPoint) {
	ret.X = (r.Right-r.Left)/2 + r.Left
	ret.Y = (r.Bottom-r.Top)/2 + r.Top
	return
}

// -- TPoint

func (p *TPoint) IsZero() bool {
	return p.X == 0 && p.Y == 0
}

func (p *TPoint) Offset(dx, dy int32) {
	p.X += dx
	p.Y += dy
}

// TMsg: Only Windows,  tagMSG
type TMsg struct {
	Hwnd    HWND
	Message uint32
	WParam  uintptr
	LParam  uintptr
	Time    uint32
	Pt      TPoint
}
