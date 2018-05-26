
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

type TGIFImage struct {
    IObject
    instance uintptr
}

func NewGIFImage() *TGIFImage {
    g := new(TGIFImage)
    g.instance = GIFImage_Create()
    return g
}

func GIFImageFromInst(inst uintptr) *TGIFImage {
    g := new(TGIFImage)
    g.instance = inst
    return g
}

func GIFImageFromObj(obj IObject) *TGIFImage {
    g := new(TGIFImage)
    g.instance = CheckPtr(obj)
    return g
}

func (g *TGIFImage) Free() {
    if g.instance != 0 {
        GIFImage_Free(g.instance)
        g.instance = 0
    }
}

func (g *TGIFImage) Instance() uintptr {
    return g.instance
}

func (g *TGIFImage) IsValid() bool {
    return g.instance != 0
}

func (g *TGIFImage) SaveToStream(Stream IObject) {
    GIFImage_SaveToStream(g.instance, CheckPtr(Stream))
}

func (g *TGIFImage) LoadFromStream(Stream IObject) {
    GIFImage_LoadFromStream(g.instance, CheckPtr(Stream))
}

func (g *TGIFImage) Add(Source IObject) *TGIFFrame {
    return GIFFrameFromInst(GIFImage_Add(g.instance, CheckPtr(Source)))
}

func (g *TGIFImage) Clear() {
    GIFImage_Clear(g.instance)
}

func (g *TGIFImage) Assign(Source IObject) {
    GIFImage_Assign(g.instance, CheckPtr(Source))
}

func (g *TGIFImage) StopDraw() {
    GIFImage_StopDraw(g.instance)
}

func (g *TGIFImage) SuspendDraw() {
    GIFImage_SuspendDraw(g.instance)
}

func (g *TGIFImage) ResumeDraw() {
    GIFImage_ResumeDraw(g.instance)
}

func (g *TGIFImage) Equals(Obj IObject) bool {
    return GIFImage_Equals(g.instance, CheckPtr(Obj))
}

func (g *TGIFImage) LoadFromFile(Filename string) {
    GIFImage_LoadFromFile(g.instance, Filename)
}

func (g *TGIFImage) SaveToFile(Filename string) {
    GIFImage_SaveToFile(g.instance, Filename)
}

func (g *TGIFImage) SetSize(AWidth int32, AHeight int32) {
    GIFImage_SetSize(g.instance, AWidth , AHeight)
}

func (g *TGIFImage) GetNamePath() string {
    return GIFImage_GetNamePath(g.instance)
}

func (g *TGIFImage) ClassName() string {
    return GIFImage_ClassName(g.instance)
}

func (g *TGIFImage) GetHashCode() int32 {
    return GIFImage_GetHashCode(g.instance)
}

func (g *TGIFImage) ToString() string {
    return GIFImage_ToString(g.instance)
}

func (g *TGIFImage) Version() TGIFVersion {
    return GIFImage_GetVersion(g.instance)
}

func (g *TGIFImage) BitsPerPixel() int32 {
    return GIFImage_GetBitsPerPixel(g.instance)
}

func (g *TGIFImage) BackgroundColor() TColor {
    return GIFImage_GetBackgroundColor(g.instance)
}

func (g *TGIFImage) SetBackgroundColor(value TColor) {
    GIFImage_SetBackgroundColor(g.instance, value)
}

func (g *TGIFImage) AspectRatio() uint8 {
    return GIFImage_GetAspectRatio(g.instance)
}

func (g *TGIFImage) SetAspectRatio(value uint8) {
    GIFImage_SetAspectRatio(g.instance, value)
}

func (g *TGIFImage) IsTransparent() bool {
    return GIFImage_GetIsTransparent(g.instance)
}

func (g *TGIFImage) AnimationSpeed() int32 {
    return GIFImage_GetAnimationSpeed(g.instance)
}

func (g *TGIFImage) SetAnimationSpeed(value int32) {
    GIFImage_SetAnimationSpeed(g.instance, value)
}

func (g *TGIFImage) Bitmap() *TBitmap {
    return BitmapFromInst(GIFImage_GetBitmap(g.instance))
}

func (g *TGIFImage) SetOnPaint(fn TNotifyEvent) {
    GIFImage_SetOnPaint(g.instance, fn)
}

func (g *TGIFImage) Animate() bool {
    return GIFImage_GetAnimate(g.instance)
}

func (g *TGIFImage) SetAnimate(value bool) {
    GIFImage_SetAnimate(g.instance, value)
}

func (g *TGIFImage) AnimateLoop() TGIFAnimationLoop {
    return GIFImage_GetAnimateLoop(g.instance)
}

func (g *TGIFImage) SetAnimateLoop(value TGIFAnimationLoop) {
    GIFImage_SetAnimateLoop(g.instance, value)
}

func (g *TGIFImage) ShouldDither() bool {
    return GIFImage_GetShouldDither(g.instance)
}

func (g *TGIFImage) Empty() bool {
    return GIFImage_GetEmpty(g.instance)
}

func (g *TGIFImage) Height() int32 {
    return GIFImage_GetHeight(g.instance)
}

func (g *TGIFImage) SetHeight(value int32) {
    GIFImage_SetHeight(g.instance, value)
}

func (g *TGIFImage) Modified() bool {
    return GIFImage_GetModified(g.instance)
}

func (g *TGIFImage) SetModified(value bool) {
    GIFImage_SetModified(g.instance, value)
}

func (g *TGIFImage) PaletteModified() bool {
    return GIFImage_GetPaletteModified(g.instance)
}

func (g *TGIFImage) SetPaletteModified(value bool) {
    GIFImage_SetPaletteModified(g.instance, value)
}

func (g *TGIFImage) Transparent() bool {
    return GIFImage_GetTransparent(g.instance)
}

func (g *TGIFImage) SetTransparent(value bool) {
    GIFImage_SetTransparent(g.instance, value)
}

func (g *TGIFImage) Width() int32 {
    return GIFImage_GetWidth(g.instance)
}

func (g *TGIFImage) SetWidth(value int32) {
    GIFImage_SetWidth(g.instance, value)
}

func (g *TGIFImage) SetOnChange(fn TNotifyEvent) {
    GIFImage_SetOnChange(g.instance, fn)
}

