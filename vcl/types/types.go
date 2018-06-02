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

type HMENU uintptr

type HICON uintptr

type HDC uintptr

type HFONT uintptr

type HBRUSH uintptr

type HPEN uintptr

type HKEY uintptr

type HMONITOR uintptr

type HGDIOBJ uintptr

type HMODULE uintptr

type COLORREF uint32

type DWORD uint32

type HCURSOR HICON

type HINST uintptr

type LPCWSTR uintptr

type HRGN uintptr

type UINT uint32

type LPARAM uintptr

type WAPRAM uintptr

//----------------------------------------------------------------------------------------------------------------------
// -- TRect

func (r *TRect) PtInRect(P TPoint) bool {
	return P.X >= r.Left && P.X < r.Right && P.Y >= r.Top && P.Y < r.Bottom
}

func (r *TRect) Width() int32 {
	return r.Right - r.Left
}

func (r *TRect) SetWidth(val int32) {
	r.Right = r.Left + val
}

func (r *TRect) Height() int32 {
	return r.Bottom - r.Top
}

func (r *TRect) SetHeight(val int32) {
	r.Bottom = r.Top + val
}

func (r *TRect) IsEmpty() bool {
	return r.Right <= r.Left || r.Bottom <= r.Top
}

func (r *TRect) Empty() {
	r.Left = 0
	r.Top = 0
	r.Right = 0
	r.Bottom = 0
}

func (r *TRect) Size() TSize {
	s := TSize{r.Width(), r.Height()}
	return s
}

func (r *TRect) SetSize(w, h int32) {
	r.SetWidth(w)
	r.SetHeight(h)
}

func (r *TRect) Inflate(dx, dy int32) {
	r.Left += -dx
	r.Top += -dy
	r.Right += dx
	r.Bottom += dy
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

func (r *TRect) Scale(val float64) {
	r.Left = int32(float64(r.Left) * val)
	r.Top = int32(float64(r.Top) * val)
	r.Right = int32(float64(r.Right) * val)
	r.Bottom = int32(float64(r.Bottom) * val)
}

func (r *TRect) Scale2(val int) {
	r.Scale(float64(val))
}

// -- TPoint

func (p *TPoint) IsZero() bool {
	return p.X == 0 && p.Y == 0
}

func (p *TPoint) Offset(dx, dy int32) {
	p.X += dx
	p.Y += dy
}

func (p *TPoint) Scale(val float64) {
	p.X = int32(float64(p.X) * val)
	p.Y = int32(float64(p.Y) * val)
}

func (p *TPoint) Scale2(val int) {
	p.Scale(float64(val))
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

// TCursorInfo
type TCursorInfo struct {
	CbSize      uint32
	Flags       uint32
	HCursor     HCURSOR
	PtScreenPos TPoint
}

// TWndClass
type TWndClass struct {
	Style         uint32
	LpfnWndProc   uintptr
	CbClsExtra    int32
	CbWndExtra    int32
	HInstance     uintptr
	HIcon         HICON
	HCursor       HCURSOR
	HbrBackground HBRUSH
	LpszMenuName  LPCWSTR
	LpszClassName LPCWSTR
}
