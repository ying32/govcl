
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "github.com/ying32/govcl/vcl/api"
    . "github.com/ying32/govcl/vcl/types"
)

type TCoolBand struct {
    IObject
    instance uintptr
}

func NewCoolBand() *TCoolBand {
    c := new(TCoolBand)
    c.instance = CoolBand_Create()
    return c
}

func CoolBandFromInst(inst uintptr) *TCoolBand {
    c := new(TCoolBand)
    c.instance = inst
    return c
}

func CoolBandFromObj(obj IObject) *TCoolBand {
    c := new(TCoolBand)
    c.instance = CheckPtr(obj)
    return c
}

func (c *TCoolBand) Free() {
    if c.instance != 0 {
        CoolBand_Free(c.instance)
        c.instance = 0
    }
}

func (c *TCoolBand) Instance() uintptr {
    return c.instance
}

func (c *TCoolBand) IsValid() bool {
    return c.instance != 0
}

func TCoolBandClass() TClass {
    return CoolBand_StaticClassType()
}

func (c *TCoolBand) Assign(Source IObject) {
    CoolBand_Assign(c.instance, CheckPtr(Source))
}

func (c *TCoolBand) GetNamePath() string {
    return CoolBand_GetNamePath(c.instance)
}

func (c *TCoolBand) DisposeOf() {
    CoolBand_DisposeOf(c.instance)
}

func (c *TCoolBand) ClassType() TClass {
    return CoolBand_ClassType(c.instance)
}

func (c *TCoolBand) ClassName() string {
    return CoolBand_ClassName(c.instance)
}

func (c *TCoolBand) InstanceSize() int32 {
    return CoolBand_InstanceSize(c.instance)
}

func (c *TCoolBand) InheritsFrom(AClass TClass) bool {
    return CoolBand_InheritsFrom(c.instance, AClass)
}

func (c *TCoolBand) Equals(Obj IObject) bool {
    return CoolBand_Equals(c.instance, CheckPtr(Obj))
}

func (c *TCoolBand) GetHashCode() int32 {
    return CoolBand_GetHashCode(c.instance)
}

func (c *TCoolBand) ToString() string {
    return CoolBand_ToString(c.instance)
}

func (c *TCoolBand) Height() int32 {
    return CoolBand_GetHeight(c.instance)
}

func (c *TCoolBand) Bitmap() *TBitmap {
    return BitmapFromInst(CoolBand_GetBitmap(c.instance))
}

func (c *TCoolBand) SetBitmap(value *TBitmap) {
    CoolBand_SetBitmap(c.instance, CheckPtr(value))
}

func (c *TCoolBand) BorderStyle() TBorderStyle {
    return CoolBand_GetBorderStyle(c.instance)
}

func (c *TCoolBand) SetBorderStyle(value TBorderStyle) {
    CoolBand_SetBorderStyle(c.instance, value)
}

func (c *TCoolBand) Break() bool {
    return CoolBand_GetBreak(c.instance)
}

func (c *TCoolBand) SetBreak(value bool) {
    CoolBand_SetBreak(c.instance, value)
}

func (c *TCoolBand) Color() TColor {
    return CoolBand_GetColor(c.instance)
}

func (c *TCoolBand) SetColor(value TColor) {
    CoolBand_SetColor(c.instance, value)
}

func (c *TCoolBand) Control() *TWinControl {
    return WinControlFromInst(CoolBand_GetControl(c.instance))
}

func (c *TCoolBand) SetControl(value IWinControl) {
    CoolBand_SetControl(c.instance, CheckPtr(value))
}

func (c *TCoolBand) FixedBackground() bool {
    return CoolBand_GetFixedBackground(c.instance)
}

func (c *TCoolBand) SetFixedBackground(value bool) {
    CoolBand_SetFixedBackground(c.instance, value)
}

func (c *TCoolBand) FixedSize() bool {
    return CoolBand_GetFixedSize(c.instance)
}

func (c *TCoolBand) SetFixedSize(value bool) {
    CoolBand_SetFixedSize(c.instance, value)
}

func (c *TCoolBand) HorizontalOnly() bool {
    return CoolBand_GetHorizontalOnly(c.instance)
}

func (c *TCoolBand) SetHorizontalOnly(value bool) {
    CoolBand_SetHorizontalOnly(c.instance, value)
}

func (c *TCoolBand) ImageIndex() int32 {
    return CoolBand_GetImageIndex(c.instance)
}

func (c *TCoolBand) SetImageIndex(value int32) {
    CoolBand_SetImageIndex(c.instance, value)
}

func (c *TCoolBand) MinHeight() int32 {
    return CoolBand_GetMinHeight(c.instance)
}

func (c *TCoolBand) SetMinHeight(value int32) {
    CoolBand_SetMinHeight(c.instance, value)
}

func (c *TCoolBand) MinWidth() int32 {
    return CoolBand_GetMinWidth(c.instance)
}

func (c *TCoolBand) SetMinWidth(value int32) {
    CoolBand_SetMinWidth(c.instance, value)
}

func (c *TCoolBand) ParentColor() bool {
    return CoolBand_GetParentColor(c.instance)
}

func (c *TCoolBand) SetParentColor(value bool) {
    CoolBand_SetParentColor(c.instance, value)
}

func (c *TCoolBand) ParentBitmap() bool {
    return CoolBand_GetParentBitmap(c.instance)
}

func (c *TCoolBand) SetParentBitmap(value bool) {
    CoolBand_SetParentBitmap(c.instance, value)
}

func (c *TCoolBand) Text() string {
    return CoolBand_GetText(c.instance)
}

func (c *TCoolBand) SetText(value string) {
    CoolBand_SetText(c.instance, value)
}

func (c *TCoolBand) Visible() bool {
    return CoolBand_GetVisible(c.instance)
}

func (c *TCoolBand) SetVisible(value bool) {
    CoolBand_SetVisible(c.instance, value)
}

func (c *TCoolBand) Width() int32 {
    return CoolBand_GetWidth(c.instance)
}

func (c *TCoolBand) SetWidth(value int32) {
    CoolBand_SetWidth(c.instance, value)
}

func (c *TCoolBand) Index() int32 {
    return CoolBand_GetIndex(c.instance)
}

func (c *TCoolBand) SetIndex(value int32) {
    CoolBand_SetIndex(c.instance, value)
}

