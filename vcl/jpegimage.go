
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

type TJPEGImage struct {
    IObject
    instance uintptr
}

func NewJPEGImage() *TJPEGImage {
    j := new(TJPEGImage)
    j.instance = JPEGImage_Create()
    return j
}

func JPEGImageFromInst(inst uintptr) *TJPEGImage {
    j := new(TJPEGImage)
    j.instance = inst
    return j
}

func JPEGImageFromObj(obj IObject) *TJPEGImage {
    j := new(TJPEGImage)
    j.instance = CheckPtr(obj)
    return j
}

func (j *TJPEGImage) Free() {
    if j.instance != 0 {
        JPEGImage_Free(j.instance)
        j.instance = 0
    }
}

func (j *TJPEGImage) Instance() uintptr {
    return j.instance
}

func (j *TJPEGImage) IsValid() bool {
    return j.instance != 0
}

func (j *TJPEGImage) Assign(Source IObject) {
    JPEGImage_Assign(j.instance, CheckPtr(Source))
}

func (j *TJPEGImage) LoadFromStream(Stream IObject) {
    JPEGImage_LoadFromStream(j.instance, CheckPtr(Stream))
}

func (j *TJPEGImage) SaveToStream(Stream IObject) {
    JPEGImage_SaveToStream(j.instance, CheckPtr(Stream))
}

func (j *TJPEGImage) Equals(Obj IObject) bool {
    return JPEGImage_Equals(j.instance, CheckPtr(Obj))
}

func (j *TJPEGImage) LoadFromFile(Filename string) {
    JPEGImage_LoadFromFile(j.instance, Filename)
}

func (j *TJPEGImage) SaveToFile(Filename string) {
    JPEGImage_SaveToFile(j.instance, Filename)
}

func (j *TJPEGImage) SetSize(AWidth int32, AHeight int32) {
    JPEGImage_SetSize(j.instance, AWidth , AHeight)
}

func (j *TJPEGImage) GetNamePath() string {
    return JPEGImage_GetNamePath(j.instance)
}

func (j *TJPEGImage) ClassName() string {
    return JPEGImage_ClassName(j.instance)
}

func (j *TJPEGImage) GetHashCode() int32 {
    return JPEGImage_GetHashCode(j.instance)
}

func (j *TJPEGImage) ToString() string {
    return JPEGImage_ToString(j.instance)
}

func (j *TJPEGImage) PixelFormat() TJPEGPixelFormat {
    return JPEGImage_GetPixelFormat(j.instance)
}

func (j *TJPEGImage) SetPixelFormat(value TJPEGPixelFormat) {
    JPEGImage_SetPixelFormat(j.instance, value)
}

func (j *TJPEGImage) ProgressiveDisplay() bool {
    return JPEGImage_GetProgressiveDisplay(j.instance)
}

func (j *TJPEGImage) SetProgressiveDisplay(value bool) {
    JPEGImage_SetProgressiveDisplay(j.instance, value)
}

func (j *TJPEGImage) Performance() TJPEGPerformance {
    return JPEGImage_GetPerformance(j.instance)
}

func (j *TJPEGImage) SetPerformance(value TJPEGPerformance) {
    JPEGImage_SetPerformance(j.instance, value)
}

func (j *TJPEGImage) Scale() TJPEGScale {
    return JPEGImage_GetScale(j.instance)
}

func (j *TJPEGImage) SetScale(value TJPEGScale) {
    JPEGImage_SetScale(j.instance, value)
}

func (j *TJPEGImage) Smoothing() bool {
    return JPEGImage_GetSmoothing(j.instance)
}

func (j *TJPEGImage) SetSmoothing(value bool) {
    JPEGImage_SetSmoothing(j.instance, value)
}

func (j *TJPEGImage) Canvas() *TCanvas {
    return CanvasFromInst(JPEGImage_GetCanvas(j.instance))
}

func (j *TJPEGImage) Empty() bool {
    return JPEGImage_GetEmpty(j.instance)
}

func (j *TJPEGImage) Height() int32 {
    return JPEGImage_GetHeight(j.instance)
}

func (j *TJPEGImage) SetHeight(value int32) {
    JPEGImage_SetHeight(j.instance, value)
}

func (j *TJPEGImage) Modified() bool {
    return JPEGImage_GetModified(j.instance)
}

func (j *TJPEGImage) SetModified(value bool) {
    JPEGImage_SetModified(j.instance, value)
}

func (j *TJPEGImage) PaletteModified() bool {
    return JPEGImage_GetPaletteModified(j.instance)
}

func (j *TJPEGImage) SetPaletteModified(value bool) {
    JPEGImage_SetPaletteModified(j.instance, value)
}

func (j *TJPEGImage) Transparent() bool {
    return JPEGImage_GetTransparent(j.instance)
}

func (j *TJPEGImage) SetTransparent(value bool) {
    JPEGImage_SetTransparent(j.instance, value)
}

func (j *TJPEGImage) Width() int32 {
    return JPEGImage_GetWidth(j.instance)
}

func (j *TJPEGImage) SetWidth(value int32) {
    JPEGImage_SetWidth(j.instance, value)
}

func (j *TJPEGImage) SetOnChange(fn TNotifyEvent) {
    JPEGImage_SetOnChange(j.instance, fn)
}

