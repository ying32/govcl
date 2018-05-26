
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TGraphic struct {
    IObject
    instance uintptr
}

func NewGraphic() *TGraphic {
    g := new(TGraphic)
    g.instance = Graphic_Create()
    return g
}

func GraphicFromInst(inst uintptr) *TGraphic {
    g := new(TGraphic)
    g.instance = inst
    return g
}

func GraphicFromObj(obj IObject) *TGraphic {
    g := new(TGraphic)
    g.instance = CheckPtr(obj)
    return g
}

func (g *TGraphic) Free() {
    if g.instance != 0 {
        Graphic_Free(g.instance)
        g.instance = 0
    }
}

func (g *TGraphic) Instance() uintptr {
    return g.instance
}

func (g *TGraphic) IsValid() bool {
    return g.instance != 0
}

func (g *TGraphic) Equals(Obj IObject) bool {
    return Graphic_Equals(g.instance, CheckPtr(Obj))
}

func (g *TGraphic) LoadFromFile(Filename string) {
    Graphic_LoadFromFile(g.instance, Filename)
}

func (g *TGraphic) SaveToFile(Filename string) {
    Graphic_SaveToFile(g.instance, Filename)
}

func (g *TGraphic) LoadFromStream(Stream IObject) {
    Graphic_LoadFromStream(g.instance, CheckPtr(Stream))
}

func (g *TGraphic) SaveToStream(Stream IObject) {
    Graphic_SaveToStream(g.instance, CheckPtr(Stream))
}

func (g *TGraphic) SetSize(AWidth int32, AHeight int32) {
    Graphic_SetSize(g.instance, AWidth , AHeight)
}

func (g *TGraphic) Assign(Source IObject) {
    Graphic_Assign(g.instance, CheckPtr(Source))
}

func (g *TGraphic) GetNamePath() string {
    return Graphic_GetNamePath(g.instance)
}

func (g *TGraphic) ClassName() string {
    return Graphic_ClassName(g.instance)
}

func (g *TGraphic) GetHashCode() int32 {
    return Graphic_GetHashCode(g.instance)
}

func (g *TGraphic) ToString() string {
    return Graphic_ToString(g.instance)
}

func (g *TGraphic) Empty() bool {
    return Graphic_GetEmpty(g.instance)
}

func (g *TGraphic) Height() int32 {
    return Graphic_GetHeight(g.instance)
}

func (g *TGraphic) SetHeight(value int32) {
    Graphic_SetHeight(g.instance, value)
}

func (g *TGraphic) Modified() bool {
    return Graphic_GetModified(g.instance)
}

func (g *TGraphic) SetModified(value bool) {
    Graphic_SetModified(g.instance, value)
}

func (g *TGraphic) PaletteModified() bool {
    return Graphic_GetPaletteModified(g.instance)
}

func (g *TGraphic) SetPaletteModified(value bool) {
    Graphic_SetPaletteModified(g.instance, value)
}

func (g *TGraphic) Transparent() bool {
    return Graphic_GetTransparent(g.instance)
}

func (g *TGraphic) SetTransparent(value bool) {
    Graphic_SetTransparent(g.instance, value)
}

func (g *TGraphic) Width() int32 {
    return Graphic_GetWidth(g.instance)
}

func (g *TGraphic) SetWidth(value int32) {
    Graphic_SetWidth(g.instance, value)
}

func (g *TGraphic) SetOnChange(fn TNotifyEvent) {
    Graphic_SetOnChange(g.instance, fn)
}

