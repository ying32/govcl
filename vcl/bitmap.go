
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

type TBitmap struct {
    IObject
    instance uintptr
}

func NewBitmap() *TBitmap {
    b := new(TBitmap)
    b.instance = Bitmap_Create()
    return b
}

func BitmapFromInst(inst uintptr) *TBitmap {
    b := new(TBitmap)
    b.instance = inst
    return b
}

func BitmapFromObj(obj IObject) *TBitmap {
    b := new(TBitmap)
    b.instance = CheckPtr(obj)
    return b
}

func (b *TBitmap) Free() {
    if b.instance != 0 {
        Bitmap_Free(b.instance)
        b.instance = 0
    }
}

func (b *TBitmap) Instance() uintptr {
    return b.instance
}

func (b *TBitmap) IsValid() bool {
    return b.instance != 0
}

func (b *TBitmap) Assign(Source IObject) {
    Bitmap_Assign(b.instance, CheckPtr(Source))
}

func (b *TBitmap) HandleAllocated() bool {
    return Bitmap_HandleAllocated(b.instance)
}

func (b *TBitmap) LoadFromStream(Stream IObject) {
    Bitmap_LoadFromStream(b.instance, CheckPtr(Stream))
}

func (b *TBitmap) SaveToStream(Stream IObject) {
    Bitmap_SaveToStream(b.instance, CheckPtr(Stream))
}

func (b *TBitmap) SetSize(AWidth int32, AHeight int32) {
    Bitmap_SetSize(b.instance, AWidth , AHeight)
}

func (b *TBitmap) LoadFromResourceName(Instance uintptr, ResName string) {
    Bitmap_LoadFromResourceName(b.instance, Instance , ResName)
}

func (b *TBitmap) LoadFromResourceID(Instance uintptr, ResID int32) {
    Bitmap_LoadFromResourceID(b.instance, Instance , ResID)
}

func (b *TBitmap) Equals(Obj IObject) bool {
    return Bitmap_Equals(b.instance, CheckPtr(Obj))
}

func (b *TBitmap) LoadFromFile(Filename string) {
    Bitmap_LoadFromFile(b.instance, Filename)
}

func (b *TBitmap) SaveToFile(Filename string) {
    Bitmap_SaveToFile(b.instance, Filename)
}

func (b *TBitmap) GetNamePath() string {
    return Bitmap_GetNamePath(b.instance)
}

func (b *TBitmap) ClassName() string {
    return Bitmap_ClassName(b.instance)
}

func (b *TBitmap) GetHashCode() int32 {
    return Bitmap_GetHashCode(b.instance)
}

func (b *TBitmap) ToString() string {
    return Bitmap_ToString(b.instance)
}

func (b *TBitmap) Canvas() *TCanvas {
    return CanvasFromInst(Bitmap_GetCanvas(b.instance))
}

func (b *TBitmap) Handle() HBITMAP {
    return Bitmap_GetHandle(b.instance)
}

func (b *TBitmap) SetHandle(value HBITMAP) {
    Bitmap_SetHandle(b.instance, value)
}

func (b *TBitmap) PixelFormat() TPixelFormat {
    return Bitmap_GetPixelFormat(b.instance)
}

func (b *TBitmap) SetPixelFormat(value TPixelFormat) {
    Bitmap_SetPixelFormat(b.instance, value)
}

func (b *TBitmap) TransparentColor() TColor {
    return Bitmap_GetTransparentColor(b.instance)
}

func (b *TBitmap) SetTransparentColor(value TColor) {
    Bitmap_SetTransparentColor(b.instance, value)
}

func (b *TBitmap) Empty() bool {
    return Bitmap_GetEmpty(b.instance)
}

func (b *TBitmap) Height() int32 {
    return Bitmap_GetHeight(b.instance)
}

func (b *TBitmap) SetHeight(value int32) {
    Bitmap_SetHeight(b.instance, value)
}

func (b *TBitmap) Modified() bool {
    return Bitmap_GetModified(b.instance)
}

func (b *TBitmap) SetModified(value bool) {
    Bitmap_SetModified(b.instance, value)
}

func (b *TBitmap) PaletteModified() bool {
    return Bitmap_GetPaletteModified(b.instance)
}

func (b *TBitmap) SetPaletteModified(value bool) {
    Bitmap_SetPaletteModified(b.instance, value)
}

func (b *TBitmap) Transparent() bool {
    return Bitmap_GetTransparent(b.instance)
}

func (b *TBitmap) SetTransparent(value bool) {
    Bitmap_SetTransparent(b.instance, value)
}

func (b *TBitmap) Width() int32 {
    return Bitmap_GetWidth(b.instance)
}

func (b *TBitmap) SetWidth(value int32) {
    Bitmap_SetWidth(b.instance, value)
}

func (b *TBitmap) SetOnChange(fn TNotifyEvent) {
    Bitmap_SetOnChange(b.instance, fn)
}

func (b *TBitmap) ScanLine(Row int32) uintptr {
    return Bitmap_GetScanLine(b.instance, Row)
}

