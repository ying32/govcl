
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

type TGIFFrame struct {
    IObject
    instance uintptr
}

func NewGIFFrame() *TGIFFrame {
    g := new(TGIFFrame)
    g.instance = GIFFrame_Create()
    return g
}

func GIFFrameFromInst(inst uintptr) *TGIFFrame {
    g := new(TGIFFrame)
    g.instance = inst
    return g
}

func GIFFrameFromObj(obj IObject) *TGIFFrame {
    g := new(TGIFFrame)
    g.instance = CheckPtr(obj)
    return g
}

func (g *TGIFFrame) Free() {
    if g.instance != 0 {
        GIFFrame_Free(g.instance)
        g.instance = 0
    }
}

func (g *TGIFFrame) Instance() uintptr {
    return g.instance
}

func (g *TGIFFrame) IsValid() bool {
    return g.instance != 0
}

func (g *TGIFFrame) Clear() {
    GIFFrame_Clear(g.instance)
}

func (g *TGIFFrame) SaveToStream(Stream IObject) {
    GIFFrame_SaveToStream(g.instance, CheckPtr(Stream))
}

func (g *TGIFFrame) LoadFromStream(Stream IObject) {
    GIFFrame_LoadFromStream(g.instance, CheckPtr(Stream))
}

func (g *TGIFFrame) Assign(Source IObject) {
    GIFFrame_Assign(g.instance, CheckPtr(Source))
}

func (g *TGIFFrame) SaveToFile(Filename string) {
    GIFFrame_SaveToFile(g.instance, Filename)
}

func (g *TGIFFrame) LoadFromFile(Filename string) {
    GIFFrame_LoadFromFile(g.instance, Filename)
}

func (g *TGIFFrame) GetNamePath() string {
    return GIFFrame_GetNamePath(g.instance)
}

func (g *TGIFFrame) ClassName() string {
    return GIFFrame_ClassName(g.instance)
}

func (g *TGIFFrame) Equals(Obj IObject) bool {
    return GIFFrame_Equals(g.instance, CheckPtr(Obj))
}

func (g *TGIFFrame) GetHashCode() int32 {
    return GIFFrame_GetHashCode(g.instance)
}

func (g *TGIFFrame) ToString() string {
    return GIFFrame_ToString(g.instance)
}

func (g *TGIFFrame) HasBitmap() bool {
    return GIFFrame_GetHasBitmap(g.instance)
}

func (g *TGIFFrame) SetHasBitmap(value bool) {
    GIFFrame_SetHasBitmap(g.instance, value)
}

func (g *TGIFFrame) Left() uint16 {
    return GIFFrame_GetLeft(g.instance)
}

func (g *TGIFFrame) SetLeft(value uint16) {
    GIFFrame_SetLeft(g.instance, value)
}

func (g *TGIFFrame) Top() uint16 {
    return GIFFrame_GetTop(g.instance)
}

func (g *TGIFFrame) SetTop(value uint16) {
    GIFFrame_SetTop(g.instance, value)
}

func (g *TGIFFrame) Width() uint16 {
    return GIFFrame_GetWidth(g.instance)
}

func (g *TGIFFrame) SetWidth(value uint16) {
    GIFFrame_SetWidth(g.instance, value)
}

func (g *TGIFFrame) Height() uint16 {
    return GIFFrame_GetHeight(g.instance)
}

func (g *TGIFFrame) SetHeight(value uint16) {
    GIFFrame_SetHeight(g.instance, value)
}

func (g *TGIFFrame) BoundsRect() TRect {
    return GIFFrame_GetBoundsRect(g.instance)
}

func (g *TGIFFrame) SetBoundsRect(value TRect) {
    GIFFrame_SetBoundsRect(g.instance, value)
}

func (g *TGIFFrame) ClientRect() TRect {
    return GIFFrame_GetClientRect(g.instance)
}

func (g *TGIFFrame) Data() uintptr {
    return GIFFrame_GetData(g.instance)
}

func (g *TGIFFrame) DataSize() int32 {
    return GIFFrame_GetDataSize(g.instance)
}

func (g *TGIFFrame) Version() TGIFVersion {
    return GIFFrame_GetVersion(g.instance)
}

func (g *TGIFFrame) BitsPerPixel() int32 {
    return GIFFrame_GetBitsPerPixel(g.instance)
}

func (g *TGIFFrame) Bitmap() *TBitmap {
    return BitmapFromInst(GIFFrame_GetBitmap(g.instance))
}

func (g *TGIFFrame) SetBitmap(value *TBitmap) {
    GIFFrame_SetBitmap(g.instance, CheckPtr(value))
}

func (g *TGIFFrame) Empty() bool {
    return GIFFrame_GetEmpty(g.instance)
}

func (g *TGIFFrame) Transparent() bool {
    return GIFFrame_GetTransparent(g.instance)
}

