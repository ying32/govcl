
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TPicture struct {
    IObject
    instance uintptr
}

func NewPicture() *TPicture {
    p := new(TPicture)
    p.instance = Picture_Create()
    return p
}

func PictureFromInst(inst uintptr) *TPicture {
    p := new(TPicture)
    p.instance = inst
    return p
}

func PictureFromObj(obj IObject) *TPicture {
    p := new(TPicture)
    p.instance = CheckPtr(obj)
    return p
}

func (p *TPicture) Free() {
    if p.instance != 0 {
        Picture_Free(p.instance)
        p.instance = 0
    }
}

func (p *TPicture) Instance() uintptr {
    return p.instance
}

func (p *TPicture) IsValid() bool {
    return p.instance != 0
}

func (p *TPicture) LoadFromFile(Filename string) {
    Picture_LoadFromFile(p.instance, Filename)
}

func (p *TPicture) SaveToFile(Filename string) {
    Picture_SaveToFile(p.instance, Filename)
}

func (p *TPicture) LoadFromStream(Stream IObject) {
    Picture_LoadFromStream(p.instance, CheckPtr(Stream))
}

func (p *TPicture) SaveToStream(Stream IObject) {
    Picture_SaveToStream(p.instance, CheckPtr(Stream))
}

func (p *TPicture) Assign(Source IObject) {
    Picture_Assign(p.instance, CheckPtr(Source))
}

func (p *TPicture) GetNamePath() string {
    return Picture_GetNamePath(p.instance)
}

func (p *TPicture) ClassName() string {
    return Picture_ClassName(p.instance)
}

func (p *TPicture) Equals(Obj IObject) bool {
    return Picture_Equals(p.instance, CheckPtr(Obj))
}

func (p *TPicture) GetHashCode() int32 {
    return Picture_GetHashCode(p.instance)
}

func (p *TPicture) ToString() string {
    return Picture_ToString(p.instance)
}

func (p *TPicture) Bitmap() *TBitmap {
    return BitmapFromInst(Picture_GetBitmap(p.instance))
}

func (p *TPicture) SetBitmap(value *TBitmap) {
    Picture_SetBitmap(p.instance, CheckPtr(value))
}

func (p *TPicture) Graphic() *TGraphic {
    return GraphicFromInst(Picture_GetGraphic(p.instance))
}

func (p *TPicture) SetGraphic(value *TGraphic) {
    Picture_SetGraphic(p.instance, CheckPtr(value))
}

func (p *TPicture) Height() int32 {
    return Picture_GetHeight(p.instance)
}

func (p *TPicture) Icon() *TIcon {
    return IconFromInst(Picture_GetIcon(p.instance))
}

func (p *TPicture) SetIcon(value *TIcon) {
    Picture_SetIcon(p.instance, CheckPtr(value))
}

func (p *TPicture) Width() int32 {
    return Picture_GetWidth(p.instance)
}

func (p *TPicture) SetOnChange(fn TNotifyEvent) {
    Picture_SetOnChange(p.instance, fn)
}

