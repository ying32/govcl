
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

type TFont struct {
    IObject
    instance uintptr
}

func NewFont() *TFont {
    f := new(TFont)
    f.instance = Font_Create()
    return f
}

func FontFromInst(inst uintptr) *TFont {
    f := new(TFont)
    f.instance = inst
    return f
}

func FontFromObj(obj IObject) *TFont {
    f := new(TFont)
    f.instance = CheckPtr(obj)
    return f
}

func (f *TFont) Free() {
    if f.instance != 0 {
        Font_Free(f.instance)
        f.instance = 0
    }
}

func (f *TFont) Instance() uintptr {
    return f.instance
}

func (f *TFont) IsValid() bool {
    return f.instance != 0
}

func (f *TFont) Assign(Source IObject) {
    Font_Assign(f.instance, CheckPtr(Source))
}

func (f *TFont) HandleAllocated() bool {
    return Font_HandleAllocated(f.instance)
}

func (f *TFont) GetNamePath() string {
    return Font_GetNamePath(f.instance)
}

func (f *TFont) ClassName() string {
    return Font_ClassName(f.instance)
}

func (f *TFont) Equals(Obj IObject) bool {
    return Font_Equals(f.instance, CheckPtr(Obj))
}

func (f *TFont) GetHashCode() int32 {
    return Font_GetHashCode(f.instance)
}

func (f *TFont) ToString() string {
    return Font_ToString(f.instance)
}

func (f *TFont) Handle() HFONT {
    return Font_GetHandle(f.instance)
}

func (f *TFont) SetHandle(value HFONT) {
    Font_SetHandle(f.instance, value)
}

func (f *TFont) PixelsPerInch() int32 {
    return Font_GetPixelsPerInch(f.instance)
}

func (f *TFont) SetPixelsPerInch(value int32) {
    Font_SetPixelsPerInch(f.instance, value)
}

func (f *TFont) Charset() TFontCharset {
    return Font_GetCharset(f.instance)
}

func (f *TFont) SetCharset(value TFontCharset) {
    Font_SetCharset(f.instance, value)
}

func (f *TFont) Color() TColor {
    return Font_GetColor(f.instance)
}

func (f *TFont) SetColor(value TColor) {
    Font_SetColor(f.instance, value)
}

func (f *TFont) Height() int32 {
    return Font_GetHeight(f.instance)
}

func (f *TFont) SetHeight(value int32) {
    Font_SetHeight(f.instance, value)
}

func (f *TFont) Name() string {
    return Font_GetName(f.instance)
}

func (f *TFont) SetName(value string) {
    Font_SetName(f.instance, value)
}

func (f *TFont) Orientation() int32 {
    return Font_GetOrientation(f.instance)
}

func (f *TFont) SetOrientation(value int32) {
    Font_SetOrientation(f.instance, value)
}

func (f *TFont) Pitch() TFontPitch {
    return Font_GetPitch(f.instance)
}

func (f *TFont) SetPitch(value TFontPitch) {
    Font_SetPitch(f.instance, value)
}

func (f *TFont) Size() int32 {
    return Font_GetSize(f.instance)
}

func (f *TFont) SetSize(value int32) {
    Font_SetSize(f.instance, value)
}

func (f *TFont) Style() TFontStyles {
    return Font_GetStyle(f.instance)
}

func (f *TFont) SetStyle(value TFontStyles) {
    Font_SetStyle(f.instance, value)
}

func (f *TFont) Quality() TFontQuality {
    return Font_GetQuality(f.instance)
}

func (f *TFont) SetQuality(value TFontQuality) {
    Font_SetQuality(f.instance, value)
}

func (f *TFont) SetOnChange(fn TNotifyEvent) {
    Font_SetOnChange(f.instance, fn)
}

