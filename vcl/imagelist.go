
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

type TImageList struct {
    IComponent
    instance uintptr
}

func NewImageList(owner IComponent) *TImageList {
    i := new(TImageList)
    i.instance = ImageList_Create(CheckPtr(owner))
    return i
}

func ImageListFromInst(inst uintptr) *TImageList {
    i := new(TImageList)
    i.instance = inst
    return i
}

func ImageListFromObj(obj IObject) *TImageList {
    i := new(TImageList)
    i.instance = CheckPtr(obj)
    return i
}

func (i *TImageList) Free() {
    if i.instance != 0 {
        ImageList_Free(i.instance)
        i.instance = 0
    }
}

func (i *TImageList) Instance() uintptr {
    return i.instance
}

func (i *TImageList) IsValid() bool {
    return i.instance != 0
}

func (i *TImageList) GetHotSpot() TPoint {
    return ImageList_GetHotSpot(i.instance)
}

func (i *TImageList) Assign(Source IObject) {
    ImageList_Assign(i.instance, CheckPtr(Source))
}

func (i *TImageList) Add(Image *TBitmap, Mask *TBitmap) int32 {
    return ImageList_Add(i.instance, CheckPtr(Image), CheckPtr(Mask))
}

func (i *TImageList) AddIcon(Image *TIcon) int32 {
    return ImageList_AddIcon(i.instance, CheckPtr(Image))
}

func (i *TImageList) AddImage(Value IComponent, Index int32) int32 {
    return ImageList_AddImage(i.instance, CheckPtr(Value), Index)
}

func (i *TImageList) AddImages(Value IComponent) {
    ImageList_AddImages(i.instance, CheckPtr(Value))
}

func (i *TImageList) AddMasked(Image *TBitmap, MaskColor TColor) int32 {
    return ImageList_AddMasked(i.instance, CheckPtr(Image), MaskColor)
}

func (i *TImageList) Clear() {
    ImageList_Clear(i.instance)
}

func (i *TImageList) Delete(Index int32) {
    ImageList_Delete(i.instance, Index)
}

func (i *TImageList) FileLoad(ResType TResType, Name string, MaskColor TColor) bool {
    return ImageList_FileLoad(i.instance, ResType , Name , MaskColor)
}

func (i *TImageList) GetBitmap(Index int32, Image *TBitmap) bool {
    return ImageList_GetBitmap(i.instance, Index , CheckPtr(Image))
}

func (i *TImageList) GetImageBitmap() HBITMAP {
    return ImageList_GetImageBitmap(i.instance)
}

func (i *TImageList) GetMaskBitmap() HBITMAP {
    return ImageList_GetMaskBitmap(i.instance)
}

func (i *TImageList) GetResource(ResType TResType, Name string, Width int32, LoadFlags TLoadResources, MaskColor TColor) bool {
    return ImageList_GetResource(i.instance, ResType , Name , Width , LoadFlags , MaskColor)
}

func (i *TImageList) HandleAllocated() bool {
    return ImageList_HandleAllocated(i.instance)
}

func (i *TImageList) Insert(Index int32, Image *TBitmap, Mask *TBitmap) {
    ImageList_Insert(i.instance, Index , CheckPtr(Image), CheckPtr(Mask))
}

func (i *TImageList) InsertIcon(Index int32, Image *TIcon) {
    ImageList_InsertIcon(i.instance, Index , CheckPtr(Image))
}

func (i *TImageList) InsertMasked(Index int32, Image *TBitmap, MaskColor TColor) {
    ImageList_InsertMasked(i.instance, Index , CheckPtr(Image), MaskColor)
}

func (i *TImageList) Move(CurIndex int32, NewIndex int32) {
    ImageList_Move(i.instance, CurIndex , NewIndex)
}

func (i *TImageList) Overlay(ImageIndex int32, Overlay uint8) bool {
    return ImageList_Overlay(i.instance, ImageIndex , Overlay)
}

func (i *TImageList) ResourceLoad(ResType TResType, Name string, MaskColor TColor) bool {
    return ImageList_ResourceLoad(i.instance, ResType , Name , MaskColor)
}

func (i *TImageList) ResInstLoad(Instance uintptr, ResType TResType, Name string, MaskColor TColor) bool {
    return ImageList_ResInstLoad(i.instance, Instance , ResType , Name , MaskColor)
}

func (i *TImageList) Replace(Index int32, Image *TBitmap, Mask *TBitmap) {
    ImageList_Replace(i.instance, Index , CheckPtr(Image), CheckPtr(Mask))
}

func (i *TImageList) ReplaceIcon(Index int32, Image *TIcon) {
    ImageList_ReplaceIcon(i.instance, Index , CheckPtr(Image))
}

func (i *TImageList) ReplaceMasked(Index int32, NewImage *TBitmap, MaskColor TColor) {
    ImageList_ReplaceMasked(i.instance, Index , CheckPtr(NewImage), MaskColor)
}

func (i *TImageList) SetSize(AWidth int32, AHeight int32) {
    ImageList_SetSize(i.instance, AWidth , AHeight)
}

func (i *TImageList) BeginUpdate() {
    ImageList_BeginUpdate(i.instance)
}

func (i *TImageList) EndUpdate() {
    ImageList_EndUpdate(i.instance)
}

func (i *TImageList) FindComponent(AName string) *TComponent {
    return ComponentFromInst(ImageList_FindComponent(i.instance, AName))
}

func (i *TImageList) GetNamePath() string {
    return ImageList_GetNamePath(i.instance)
}

func (i *TImageList) HasParent() bool {
    return ImageList_HasParent(i.instance)
}

func (i *TImageList) ClassName() string {
    return ImageList_ClassName(i.instance)
}

func (i *TImageList) Equals(Obj IObject) bool {
    return ImageList_Equals(i.instance, CheckPtr(Obj))
}

func (i *TImageList) GetHashCode() int32 {
    return ImageList_GetHashCode(i.instance)
}

func (i *TImageList) ToString() string {
    return ImageList_ToString(i.instance)
}

func (i *TImageList) BlendColor() TColor {
    return ImageList_GetBlendColor(i.instance)
}

func (i *TImageList) SetBlendColor(value TColor) {
    ImageList_SetBlendColor(i.instance, value)
}

func (i *TImageList) BkColor() TColor {
    return ImageList_GetBkColor(i.instance)
}

func (i *TImageList) SetBkColor(value TColor) {
    ImageList_SetBkColor(i.instance, value)
}

func (i *TImageList) AllocBy() int32 {
    return ImageList_GetAllocBy(i.instance)
}

func (i *TImageList) SetAllocBy(value int32) {
    ImageList_SetAllocBy(i.instance, value)
}

func (i *TImageList) ColorDepth() TColorDepth {
    return ImageList_GetColorDepth(i.instance)
}

func (i *TImageList) SetColorDepth(value TColorDepth) {
    ImageList_SetColorDepth(i.instance, value)
}

func (i *TImageList) DrawingStyle() TDrawingStyle {
    return ImageList_GetDrawingStyle(i.instance)
}

func (i *TImageList) SetDrawingStyle(value TDrawingStyle) {
    ImageList_SetDrawingStyle(i.instance, value)
}

func (i *TImageList) GrayscaleFactor() uint8 {
    return ImageList_GetGrayscaleFactor(i.instance)
}

func (i *TImageList) SetGrayscaleFactor(value uint8) {
    ImageList_SetGrayscaleFactor(i.instance, value)
}

func (i *TImageList) Height() int32 {
    return ImageList_GetHeight(i.instance)
}

func (i *TImageList) SetHeight(value int32) {
    ImageList_SetHeight(i.instance, value)
}

func (i *TImageList) ImageType() TImageType {
    return ImageList_GetImageType(i.instance)
}

func (i *TImageList) SetImageType(value TImageType) {
    ImageList_SetImageType(i.instance, value)
}

func (i *TImageList) Masked() bool {
    return ImageList_GetMasked(i.instance)
}

func (i *TImageList) SetMasked(value bool) {
    ImageList_SetMasked(i.instance, value)
}

func (i *TImageList) SetOnChange(fn TNotifyEvent) {
    ImageList_SetOnChange(i.instance, fn)
}

func (i *TImageList) ShareImages() bool {
    return ImageList_GetShareImages(i.instance)
}

func (i *TImageList) SetShareImages(value bool) {
    ImageList_SetShareImages(i.instance, value)
}

func (i *TImageList) Width() int32 {
    return ImageList_GetWidth(i.instance)
}

func (i *TImageList) SetWidth(value int32) {
    ImageList_SetWidth(i.instance, value)
}

func (i *TImageList) Handle() uintptr {
    return ImageList_GetHandle(i.instance)
}

func (i *TImageList) SetHandle(value uintptr) {
    ImageList_SetHandle(i.instance, value)
}

func (i *TImageList) Count() int32 {
    return ImageList_GetCount(i.instance)
}

func (i *TImageList) ComponentCount() int32 {
    return ImageList_GetComponentCount(i.instance)
}

func (i *TImageList) ComponentIndex() int32 {
    return ImageList_GetComponentIndex(i.instance)
}

func (i *TImageList) SetComponentIndex(value int32) {
    ImageList_SetComponentIndex(i.instance, value)
}

func (i *TImageList) Owner() *TComponent {
    return ComponentFromInst(ImageList_GetOwner(i.instance))
}

func (i *TImageList) Name() string {
    return ImageList_GetName(i.instance)
}

func (i *TImageList) SetName(value string) {
    ImageList_SetName(i.instance, value)
}

func (i *TImageList) Tag() int {
    return ImageList_GetTag(i.instance)
}

func (i *TImageList) SetTag(value int) {
    ImageList_SetTag(i.instance, value)
}

func (i *TImageList) Components(AIndex int32) *TComponent {
    return ComponentFromInst(ImageList_GetComponents(i.instance, AIndex))
}

