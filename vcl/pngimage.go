
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

type TPngImage struct {
    IObject
    instance uintptr
}

func NewPngImage() *TPngImage {
    p := new(TPngImage)
    p.instance = PngImage_Create()
    return p
}

func PngImageFromInst(inst uintptr) *TPngImage {
    p := new(TPngImage)
    p.instance = inst
    return p
}

func PngImageFromObj(obj IObject) *TPngImage {
    p := new(TPngImage)
    p.instance = CheckPtr(obj)
    return p
}

func (p *TPngImage) Free() {
    if p.instance != 0 {
        PngImage_Free(p.instance)
        p.instance = 0
    }
}

func (p *TPngImage) Instance() uintptr {
    return p.instance
}

func (p *TPngImage) IsValid() bool {
    return p.instance != 0
}

func (p *TPngImage) Assign(Source IObject) {
    PngImage_Assign(p.instance, CheckPtr(Source))
}

func (p *TPngImage) LoadFromStream(Stream IObject) {
    PngImage_LoadFromStream(p.instance, CheckPtr(Stream))
}

func (p *TPngImage) SaveToStream(Stream IObject) {
    PngImage_SaveToStream(p.instance, CheckPtr(Stream))
}

func (p *TPngImage) LoadFromResourceName(Instance uintptr, Name string) {
    PngImage_LoadFromResourceName(p.instance, Instance , Name)
}

func (p *TPngImage) LoadFromResourceID(Instance uintptr, ResID int32) {
    PngImage_LoadFromResourceID(p.instance, Instance , ResID)
}

func (p *TPngImage) Equals(Obj IObject) bool {
    return PngImage_Equals(p.instance, CheckPtr(Obj))
}

func (p *TPngImage) LoadFromFile(Filename string) {
    PngImage_LoadFromFile(p.instance, Filename)
}

func (p *TPngImage) SaveToFile(Filename string) {
    PngImage_SaveToFile(p.instance, Filename)
}

func (p *TPngImage) SetSize(AWidth int32, AHeight int32) {
    PngImage_SetSize(p.instance, AWidth , AHeight)
}

func (p *TPngImage) GetNamePath() string {
    return PngImage_GetNamePath(p.instance)
}

func (p *TPngImage) ClassName() string {
    return PngImage_ClassName(p.instance)
}

func (p *TPngImage) GetHashCode() int32 {
    return PngImage_GetHashCode(p.instance)
}

func (p *TPngImage) ToString() string {
    return PngImage_ToString(p.instance)
}

func (p *TPngImage) TransparentColor() TColor {
    return PngImage_GetTransparentColor(p.instance)
}

func (p *TPngImage) SetTransparentColor(value TColor) {
    PngImage_SetTransparentColor(p.instance, value)
}

func (p *TPngImage) Canvas() *TCanvas {
    return CanvasFromInst(PngImage_GetCanvas(p.instance))
}

func (p *TPngImage) Width() int32 {
    return PngImage_GetWidth(p.instance)
}

func (p *TPngImage) Height() int32 {
    return PngImage_GetHeight(p.instance)
}

func (p *TPngImage) MaxIdatSize() int32 {
    return PngImage_GetMaxIdatSize(p.instance)
}

func (p *TPngImage) SetMaxIdatSize(value int32) {
    PngImage_SetMaxIdatSize(p.instance, value)
}

func (p *TPngImage) Empty() bool {
    return PngImage_GetEmpty(p.instance)
}

func (p *TPngImage) CompressionLevel() TCompressionLevel {
    return PngImage_GetCompressionLevel(p.instance)
}

func (p *TPngImage) SetCompressionLevel(value TCompressionLevel) {
    PngImage_SetCompressionLevel(p.instance, value)
}

func (p *TPngImage) Version() string {
    return PngImage_GetVersion(p.instance)
}

func (p *TPngImage) Modified() bool {
    return PngImage_GetModified(p.instance)
}

func (p *TPngImage) SetModified(value bool) {
    PngImage_SetModified(p.instance, value)
}

func (p *TPngImage) PaletteModified() bool {
    return PngImage_GetPaletteModified(p.instance)
}

func (p *TPngImage) SetPaletteModified(value bool) {
    PngImage_SetPaletteModified(p.instance, value)
}

func (p *TPngImage) Transparent() bool {
    return PngImage_GetTransparent(p.instance)
}

func (p *TPngImage) SetTransparent(value bool) {
    PngImage_SetTransparent(p.instance, value)
}

func (p *TPngImage) SetOnChange(fn TNotifyEvent) {
    PngImage_SetOnChange(p.instance, fn)
}

