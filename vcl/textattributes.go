
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

type TTextAttributes struct {
    IObject
    instance uintptr
}

func TextAttributesFromInst(inst uintptr) *TTextAttributes {
    t := new(TTextAttributes)
    t.instance = inst
    return t
}

func TextAttributesFromObj(obj IObject) *TTextAttributes {
    t := new(TTextAttributes)
    t.instance = CheckPtr(obj)
    return t
}

func (t *TTextAttributes) Instance() uintptr {
    return t.instance
}

func (t *TTextAttributes) IsValid() bool {
    return t.instance != 0
}

func (t *TTextAttributes) Assign(Source IObject) {
    TextAttributes_Assign(t.instance, CheckPtr(Source))
}

func (t *TTextAttributes) GetNamePath() string {
    return TextAttributes_GetNamePath(t.instance)
}

func (t *TTextAttributes) ClassName() string {
    return TextAttributes_ClassName(t.instance)
}

func (t *TTextAttributes) Equals(Obj IObject) bool {
    return TextAttributes_Equals(t.instance, CheckPtr(Obj))
}

func (t *TTextAttributes) GetHashCode() int32 {
    return TextAttributes_GetHashCode(t.instance)
}

func (t *TTextAttributes) ToString() string {
    return TextAttributes_ToString(t.instance)
}

func (t *TTextAttributes) Charset() TFontCharset {
    return TextAttributes_GetCharset(t.instance)
}

func (t *TTextAttributes) SetCharset(value TFontCharset) {
    TextAttributes_SetCharset(t.instance, value)
}

func (t *TTextAttributes) Color() TColor {
    return TextAttributes_GetColor(t.instance)
}

func (t *TTextAttributes) SetColor(value TColor) {
    TextAttributes_SetColor(t.instance, value)
}

func (t *TTextAttributes) ConsistentAttributes() TConsistentAttributes {
    return TextAttributes_GetConsistentAttributes(t.instance)
}

func (t *TTextAttributes) Name() string {
    return TextAttributes_GetName(t.instance)
}

func (t *TTextAttributes) SetName(value string) {
    TextAttributes_SetName(t.instance, value)
}

func (t *TTextAttributes) Pitch() TFontPitch {
    return TextAttributes_GetPitch(t.instance)
}

func (t *TTextAttributes) SetPitch(value TFontPitch) {
    TextAttributes_SetPitch(t.instance, value)
}

func (t *TTextAttributes) Protected() bool {
    return TextAttributes_GetProtected(t.instance)
}

func (t *TTextAttributes) SetProtected(value bool) {
    TextAttributes_SetProtected(t.instance, value)
}

func (t *TTextAttributes) Size() int32 {
    return TextAttributes_GetSize(t.instance)
}

func (t *TTextAttributes) SetSize(value int32) {
    TextAttributes_SetSize(t.instance, value)
}

func (t *TTextAttributes) Style() TFontStyles {
    return TextAttributes_GetStyle(t.instance)
}

func (t *TTextAttributes) SetStyle(value TFontStyles) {
    TextAttributes_SetStyle(t.instance, value)
}

func (t *TTextAttributes) Height() int32 {
    return TextAttributes_GetHeight(t.instance)
}

func (t *TTextAttributes) SetHeight(value int32) {
    TextAttributes_SetHeight(t.instance, value)
}

