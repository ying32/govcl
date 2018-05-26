
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
    . "gitee.com/ying32/govcl/vcl/types"
)

type TCanvas struct {
    IObject
    instance uintptr
}

func NewCanvas() *TCanvas {
    c := new(TCanvas)
    c.instance = Canvas_Create()
    return c
}

func CanvasFromInst(inst uintptr) *TCanvas {
    c := new(TCanvas)
    c.instance = inst
    return c
}

func CanvasFromObj(obj IObject) *TCanvas {
    c := new(TCanvas)
    c.instance = CheckPtr(obj)
    return c
}

func (c *TCanvas) Free() {
    if c.instance != 0 {
        Canvas_Free(c.instance)
        c.instance = 0
    }
}

func (c *TCanvas) Instance() uintptr {
    return c.instance
}

func (c *TCanvas) IsValid() bool {
    return c.instance != 0
}

func (c *TCanvas) Arc(X1 int32, Y1 int32, X2 int32, Y2 int32, X3 int32, Y3 int32, X4 int32, Y4 int32) {
    Canvas_Arc(c.instance, X1 , Y1 , X2 , Y2 , X3 , Y3 , X4 , Y4)
}

func (c *TCanvas) ArcTo(X1 int32, Y1 int32, X2 int32, Y2 int32, X3 int32, Y3 int32, X4 int32, Y4 int32) {
    Canvas_ArcTo(c.instance, X1 , Y1 , X2 , Y2 , X3 , Y3 , X4 , Y4)
}

func (c *TCanvas) AngleArc(X int32, Y int32, Radius uint32, StartAngle float32, SweepAngle float32) {
    Canvas_AngleArc(c.instance, X , Y , Radius , StartAngle , SweepAngle)
}

func (c *TCanvas) Chord(X1 int32, Y1 int32, X2 int32, Y2 int32, X3 int32, Y3 int32, X4 int32, Y4 int32) {
    Canvas_Chord(c.instance, X1 , Y1 , X2 , Y2 , X3 , Y3 , X4 , Y4)
}

func (c *TCanvas) Ellipse(X1 int32, Y1 int32, X2 int32, Y2 int32) {
    Canvas_Ellipse(c.instance, X1 , Y1 , X2 , Y2)
}

func (c *TCanvas) FloodFill(X int32, Y int32, Color TColor, FillStyle TFillStyle) {
    Canvas_FloodFill(c.instance, X , Y , Color , FillStyle)
}

func (c *TCanvas) HandleAllocated() bool {
    return Canvas_HandleAllocated(c.instance)
}

func (c *TCanvas) LineTo(X int32, Y int32) {
    Canvas_LineTo(c.instance, X , Y)
}

func (c *TCanvas) MoveTo(X int32, Y int32) {
    Canvas_MoveTo(c.instance, X , Y)
}

func (c *TCanvas) Pie(X1 int32, Y1 int32, X2 int32, Y2 int32, X3 int32, Y3 int32, X4 int32, Y4 int32) {
    Canvas_Pie(c.instance, X1 , Y1 , X2 , Y2 , X3 , Y3 , X4 , Y4)
}

func (c *TCanvas) Rectangle(X1 int32, Y1 int32, X2 int32, Y2 int32) {
    Canvas_Rectangle(c.instance, X1 , Y1 , X2 , Y2)
}

func (c *TCanvas) Refresh() {
    Canvas_Refresh(c.instance)
}

func (c *TCanvas) RoundRect(X1 int32, Y1 int32, X2 int32, Y2 int32, X3 int32, Y3 int32) {
    Canvas_RoundRect(c.instance, X1 , Y1 , X2 , Y2 , X3 , Y3)
}

func (c *TCanvas) TextExtent(Text string) TSize {
    return Canvas_TextExtent(c.instance, Text)
}

func (c *TCanvas) TextOut(X int32, Y int32, Text string) {
    Canvas_TextOut(c.instance, X , Y , Text)
}

func (c *TCanvas) Lock() {
    Canvas_Lock(c.instance)
}

func (c *TCanvas) TextHeight(Text string) int32 {
    return Canvas_TextHeight(c.instance, Text)
}

func (c *TCanvas) TextWidth(Text string) int32 {
    return Canvas_TextWidth(c.instance, Text)
}

func (c *TCanvas) Assign(Source IObject) {
    Canvas_Assign(c.instance, CheckPtr(Source))
}

func (c *TCanvas) GetNamePath() string {
    return Canvas_GetNamePath(c.instance)
}

func (c *TCanvas) ClassName() string {
    return Canvas_ClassName(c.instance)
}

func (c *TCanvas) Equals(Obj IObject) bool {
    return Canvas_Equals(c.instance, CheckPtr(Obj))
}

func (c *TCanvas) GetHashCode() int32 {
    return Canvas_GetHashCode(c.instance)
}

func (c *TCanvas) ToString() string {
    return Canvas_ToString(c.instance)
}

func (c *TCanvas) Handle() HDC {
    return Canvas_GetHandle(c.instance)
}

func (c *TCanvas) SetHandle(value HDC) {
    Canvas_SetHandle(c.instance, value)
}

func (c *TCanvas) Brush() *TBrush {
    return BrushFromInst(Canvas_GetBrush(c.instance))
}

func (c *TCanvas) SetBrush(value *TBrush) {
    Canvas_SetBrush(c.instance, CheckPtr(value))
}

func (c *TCanvas) CopyMode() int32 {
    return Canvas_GetCopyMode(c.instance)
}

func (c *TCanvas) SetCopyMode(value int32) {
    Canvas_SetCopyMode(c.instance, value)
}

func (c *TCanvas) Font() *TFont {
    return FontFromInst(Canvas_GetFont(c.instance))
}

func (c *TCanvas) SetFont(value *TFont) {
    Canvas_SetFont(c.instance, CheckPtr(value))
}

func (c *TCanvas) Pen() *TPen {
    return PenFromInst(Canvas_GetPen(c.instance))
}

func (c *TCanvas) SetPen(value *TPen) {
    Canvas_SetPen(c.instance, CheckPtr(value))
}

func (c *TCanvas) SetOnChange(fn TNotifyEvent) {
    Canvas_SetOnChange(c.instance, fn)
}

