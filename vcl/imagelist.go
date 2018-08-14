
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "github.com/ying32/govcl/vcl/api"
    . "github.com/ying32/govcl/vcl/types"
    "unsafe"
)

type TImageList struct {
    IComponent
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewImageList
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewImageList(owner IComponent) *TImageList {
    i := new(TImageList)
    i.instance = ImageList_Create(CheckPtr(owner))
    i.ptr = unsafe.Pointer(i.instance)
    return i
}

// ImageListFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func ImageListFromInst(inst uintptr) *TImageList {
    i := new(TImageList)
    i.instance = inst
    i.ptr = unsafe.Pointer(inst)
    return i
}

// ImageListFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func ImageListFromObj(obj IObject) *TImageList {
    i := new(TImageList)
    i.instance = CheckPtr(obj)
    i.ptr = unsafe.Pointer(i.instance)
    return i
}

// ImageListFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func ImageListFromUnsafePointer(ptr unsafe.Pointer) *TImageList {
    i := new(TImageList)
    i.instance = uintptr(ptr)
    i.ptr = ptr
    return i
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (i *TImageList) Free() {
    if i.instance != 0 {
        ImageList_Free(i.instance)
        i.instance = 0
        i.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (i *TImageList) Instance() uintptr {
    return i.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (i *TImageList) UnsafeAddr() unsafe.Pointer {
    return i.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (i *TImageList) IsValid() bool {
    return i.instance != 0
}

// TImageListClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TImageListClass() TClass {
    return ImageList_StaticClassType()
}

// GetHotSpot
func (i *TImageList) GetHotSpot() TPoint {
    return ImageList_GetHotSpot(i.instance)
}

// HideDragImage
func (i *TImageList) HideDragImage() {
    ImageList_HideDragImage(i.instance)
}

// ShowDragImage
func (i *TImageList) ShowDragImage() {
    ImageList_ShowDragImage(i.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (i *TImageList) Assign(Source IObject) {
    ImageList_Assign(i.instance, CheckPtr(Source))
}

// Add
func (i *TImageList) Add(Image *TBitmap, Mask *TBitmap) int32 {
    return ImageList_Add(i.instance, CheckPtr(Image), CheckPtr(Mask))
}

// AddIcon
func (i *TImageList) AddIcon(Image *TIcon) int32 {
    return ImageList_AddIcon(i.instance, CheckPtr(Image))
}

// AddImage
func (i *TImageList) AddImage(Value IComponent, Index int32) int32 {
    return ImageList_AddImage(i.instance, CheckPtr(Value), Index)
}

// AddImages
func (i *TImageList) AddImages(Value IComponent) {
    ImageList_AddImages(i.instance, CheckPtr(Value))
}

// AddMasked
func (i *TImageList) AddMasked(Image *TBitmap, MaskColor TColor) int32 {
    return ImageList_AddMasked(i.instance, CheckPtr(Image), MaskColor)
}

// Clear
// CN: 清除。
// EN: .
func (i *TImageList) Clear() {
    ImageList_Clear(i.instance)
}

// Delete
func (i *TImageList) Delete(Index int32) {
    ImageList_Delete(i.instance, Index)
}

// FileLoad
func (i *TImageList) FileLoad(ResType TResType, Name string, MaskColor TColor) bool {
    return ImageList_FileLoad(i.instance, ResType , Name , MaskColor)
}

// GetBitmap
func (i *TImageList) GetBitmap(Index int32, Image *TBitmap) bool {
    return ImageList_GetBitmap(i.instance, Index , CheckPtr(Image))
}

// GetImageBitmap
func (i *TImageList) GetImageBitmap() HBITMAP {
    return ImageList_GetImageBitmap(i.instance)
}

// GetMaskBitmap
func (i *TImageList) GetMaskBitmap() HBITMAP {
    return ImageList_GetMaskBitmap(i.instance)
}

// GetResource
func (i *TImageList) GetResource(ResType TResType, Name string, Width int32, LoadFlags TLoadResources, MaskColor TColor) bool {
    return ImageList_GetResource(i.instance, ResType , Name , Width , LoadFlags , MaskColor)
}

// HandleAllocated
// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (i *TImageList) HandleAllocated() bool {
    return ImageList_HandleAllocated(i.instance)
}

// Insert
func (i *TImageList) Insert(Index int32, Image *TBitmap, Mask *TBitmap) {
    ImageList_Insert(i.instance, Index , CheckPtr(Image), CheckPtr(Mask))
}

// InsertIcon
func (i *TImageList) InsertIcon(Index int32, Image *TIcon) {
    ImageList_InsertIcon(i.instance, Index , CheckPtr(Image))
}

// InsertMasked
func (i *TImageList) InsertMasked(Index int32, Image *TBitmap, MaskColor TColor) {
    ImageList_InsertMasked(i.instance, Index , CheckPtr(Image), MaskColor)
}

// Move
func (i *TImageList) Move(CurIndex int32, NewIndex int32) {
    ImageList_Move(i.instance, CurIndex , NewIndex)
}

// Overlay
func (i *TImageList) Overlay(ImageIndex int32, Overlay uint8) bool {
    return ImageList_Overlay(i.instance, ImageIndex , Overlay)
}

// ResourceLoad
func (i *TImageList) ResourceLoad(ResType TResType, Name string, MaskColor TColor) bool {
    return ImageList_ResourceLoad(i.instance, ResType , Name , MaskColor)
}

// ResInstLoad
func (i *TImageList) ResInstLoad(Instance uintptr, ResType TResType, Name string, MaskColor TColor) bool {
    return ImageList_ResInstLoad(i.instance, Instance , ResType , Name , MaskColor)
}

// Replace
func (i *TImageList) Replace(Index int32, Image *TBitmap, Mask *TBitmap) {
    ImageList_Replace(i.instance, Index , CheckPtr(Image), CheckPtr(Mask))
}

// ReplaceIcon
func (i *TImageList) ReplaceIcon(Index int32, Image *TIcon) {
    ImageList_ReplaceIcon(i.instance, Index , CheckPtr(Image))
}

// ReplaceMasked
func (i *TImageList) ReplaceMasked(Index int32, NewImage *TBitmap, MaskColor TColor) {
    ImageList_ReplaceMasked(i.instance, Index , CheckPtr(NewImage), MaskColor)
}

// SetSize
func (i *TImageList) SetSize(AWidth int32, AHeight int32) {
    ImageList_SetSize(i.instance, AWidth , AHeight)
}

// BeginUpdate
func (i *TImageList) BeginUpdate() {
    ImageList_BeginUpdate(i.instance)
}

// EndUpdate
func (i *TImageList) EndUpdate() {
    ImageList_EndUpdate(i.instance)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (i *TImageList) FindComponent(AName string) *TComponent {
    return ComponentFromInst(ImageList_FindComponent(i.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (i *TImageList) GetNamePath() string {
    return ImageList_GetNamePath(i.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (i *TImageList) HasParent() bool {
    return ImageList_HasParent(i.instance)
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (i *TImageList) DisposeOf() {
    ImageList_DisposeOf(i.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (i *TImageList) ClassType() TClass {
    return ImageList_ClassType(i.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (i *TImageList) ClassName() string {
    return ImageList_ClassName(i.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (i *TImageList) InstanceSize() int32 {
    return ImageList_InstanceSize(i.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (i *TImageList) InheritsFrom(AClass TClass) bool {
    return ImageList_InheritsFrom(i.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (i *TImageList) Equals(Obj IObject) bool {
    return ImageList_Equals(i.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (i *TImageList) GetHashCode() int32 {
    return ImageList_GetHashCode(i.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (i *TImageList) ToString() string {
    return ImageList_ToString(i.instance)
}

// BlendColor
func (i *TImageList) BlendColor() TColor {
    return ImageList_GetBlendColor(i.instance)
}

// SetBlendColor
func (i *TImageList) SetBlendColor(value TColor) {
    ImageList_SetBlendColor(i.instance, value)
}

// BkColor
func (i *TImageList) BkColor() TColor {
    return ImageList_GetBkColor(i.instance)
}

// SetBkColor
func (i *TImageList) SetBkColor(value TColor) {
    ImageList_SetBkColor(i.instance, value)
}

// AllocBy
func (i *TImageList) AllocBy() int32 {
    return ImageList_GetAllocBy(i.instance)
}

// SetAllocBy
func (i *TImageList) SetAllocBy(value int32) {
    ImageList_SetAllocBy(i.instance, value)
}

// ColorDepth
func (i *TImageList) ColorDepth() TColorDepth {
    return ImageList_GetColorDepth(i.instance)
}

// SetColorDepth
func (i *TImageList) SetColorDepth(value TColorDepth) {
    ImageList_SetColorDepth(i.instance, value)
}

// DrawingStyle
func (i *TImageList) DrawingStyle() TDrawingStyle {
    return ImageList_GetDrawingStyle(i.instance)
}

// SetDrawingStyle
func (i *TImageList) SetDrawingStyle(value TDrawingStyle) {
    ImageList_SetDrawingStyle(i.instance, value)
}

// GrayscaleFactor
func (i *TImageList) GrayscaleFactor() uint8 {
    return ImageList_GetGrayscaleFactor(i.instance)
}

// SetGrayscaleFactor
func (i *TImageList) SetGrayscaleFactor(value uint8) {
    ImageList_SetGrayscaleFactor(i.instance, value)
}

// Height
// CN: 获取高度。
// EN: Get height.
func (i *TImageList) Height() int32 {
    return ImageList_GetHeight(i.instance)
}

// SetHeight
// CN: 设置高度。
// EN: Set height.
func (i *TImageList) SetHeight(value int32) {
    ImageList_SetHeight(i.instance, value)
}

// ImageType
func (i *TImageList) ImageType() TImageType {
    return ImageList_GetImageType(i.instance)
}

// SetImageType
func (i *TImageList) SetImageType(value TImageType) {
    ImageList_SetImageType(i.instance, value)
}

// Masked
func (i *TImageList) Masked() bool {
    return ImageList_GetMasked(i.instance)
}

// SetMasked
func (i *TImageList) SetMasked(value bool) {
    ImageList_SetMasked(i.instance, value)
}

// SetOnChange
// CN: 设置改变事件。
// EN: Set changed event.
func (i *TImageList) SetOnChange(fn TNotifyEvent) {
    ImageList_SetOnChange(i.instance, fn)
}

// ShareImages
func (i *TImageList) ShareImages() bool {
    return ImageList_GetShareImages(i.instance)
}

// SetShareImages
func (i *TImageList) SetShareImages(value bool) {
    ImageList_SetShareImages(i.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (i *TImageList) Width() int32 {
    return ImageList_GetWidth(i.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (i *TImageList) SetWidth(value int32) {
    ImageList_SetWidth(i.instance, value)
}

// DragCursor
// CN: 获取设置控件拖拽时的光标。
// EN: Get Set the cursor when the control is dragged.
func (i *TImageList) DragCursor() TCursor {
    return ImageList_GetDragCursor(i.instance)
}

// SetDragCursor
// CN: 设置设置控件拖拽时的光标。
// EN: Set Set the cursor when the control is dragged.
func (i *TImageList) SetDragCursor(value TCursor) {
    ImageList_SetDragCursor(i.instance, value)
}

// Dragging
// CN: 获取是否在拖拽中。
// EN: Get Is it in the middle of dragging.
func (i *TImageList) Dragging() bool {
    return ImageList_GetDragging(i.instance)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (i *TImageList) Handle() uintptr {
    return ImageList_GetHandle(i.instance)
}

// SetHandle
// CN: 设置控件句柄。
// EN: Set Control handle.
func (i *TImageList) SetHandle(value uintptr) {
    ImageList_SetHandle(i.instance, value)
}

// Count
func (i *TImageList) Count() int32 {
    return ImageList_GetCount(i.instance)
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (i *TImageList) ComponentCount() int32 {
    return ImageList_GetComponentCount(i.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (i *TImageList) ComponentIndex() int32 {
    return ImageList_GetComponentIndex(i.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (i *TImageList) SetComponentIndex(value int32) {
    ImageList_SetComponentIndex(i.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (i *TImageList) Owner() *TComponent {
    return ComponentFromInst(ImageList_GetOwner(i.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (i *TImageList) Name() string {
    return ImageList_GetName(i.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (i *TImageList) SetName(value string) {
    ImageList_SetName(i.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (i *TImageList) Tag() int {
    return ImageList_GetTag(i.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (i *TImageList) SetTag(value int) {
    ImageList_SetTag(i.instance, value)
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (i *TImageList) Components(AIndex int32) *TComponent {
    return ComponentFromInst(ImageList_GetComponents(i.instance, AIndex))
}

