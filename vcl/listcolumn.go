
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

type TListColumn struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewListColumn
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewListColumn() *TListColumn {
    l := new(TListColumn)
    l.instance = ListColumn_Create()
    l.ptr = unsafe.Pointer(l.instance)
    return l
}

// ListColumnFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func ListColumnFromInst(inst uintptr) *TListColumn {
    l := new(TListColumn)
    l.instance = inst
    l.ptr = unsafe.Pointer(inst)
    return l
}

// ListColumnFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func ListColumnFromObj(obj IObject) *TListColumn {
    l := new(TListColumn)
    l.instance = CheckPtr(obj)
    l.ptr = unsafe.Pointer(l.instance)
    return l
}

// ListColumnFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func ListColumnFromUnsafePointer(ptr unsafe.Pointer) *TListColumn {
    l := new(TListColumn)
    l.instance = uintptr(ptr)
    l.ptr = ptr
    return l
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (l *TListColumn) Free() {
    if l.instance != 0 {
        ListColumn_Free(l.instance)
        l.instance = 0
        l.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (l *TListColumn) Instance() uintptr {
    return l.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (l *TListColumn) UnsafeAddr() unsafe.Pointer {
    return l.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (l *TListColumn) IsValid() bool {
    return l.instance != 0
}

// TListColumnClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TListColumnClass() TClass {
    return ListColumn_StaticClassType()
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (l *TListColumn) Assign(Source IObject) {
    ListColumn_Assign(l.instance, CheckPtr(Source))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (l *TListColumn) GetNamePath() string {
    return ListColumn_GetNamePath(l.instance)
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (l *TListColumn) DisposeOf() {
    ListColumn_DisposeOf(l.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (l *TListColumn) ClassType() TClass {
    return ListColumn_ClassType(l.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (l *TListColumn) ClassName() string {
    return ListColumn_ClassName(l.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (l *TListColumn) InstanceSize() int32 {
    return ListColumn_InstanceSize(l.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (l *TListColumn) InheritsFrom(AClass TClass) bool {
    return ListColumn_InheritsFrom(l.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (l *TListColumn) Equals(Obj IObject) bool {
    return ListColumn_Equals(l.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (l *TListColumn) GetHashCode() int32 {
    return ListColumn_GetHashCode(l.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (l *TListColumn) ToString() string {
    return ListColumn_ToString(l.instance)
}

// Alignment
// CN: 获取文字对齐。
// EN: Get Text alignment.
func (l *TListColumn) Alignment() TAlignment {
    return ListColumn_GetAlignment(l.instance)
}

// SetAlignment
// CN: 设置文字对齐。
// EN: Set Text alignment.
func (l *TListColumn) SetAlignment(value TAlignment) {
    ListColumn_SetAlignment(l.instance, value)
}

// AutoSize
// CN: 获取自动调整大小。
// EN: .
func (l *TListColumn) AutoSize() bool {
    return ListColumn_GetAutoSize(l.instance)
}

// SetAutoSize
// CN: 设置自动调整大小。
// EN: .
func (l *TListColumn) SetAutoSize(value bool) {
    ListColumn_SetAutoSize(l.instance, value)
}

// Caption
// CN: 获取控件标题。
// EN: Get the control title.
func (l *TListColumn) Caption() string {
    return ListColumn_GetCaption(l.instance)
}

// SetCaption
// CN: 设置控件标题。
// EN: Set the control title.
func (l *TListColumn) SetCaption(value string) {
    ListColumn_SetCaption(l.instance, value)
}

// ImageIndex
// CN: 获取图像在images中的索引。
// EN: .
func (l *TListColumn) ImageIndex() int32 {
    return ListColumn_GetImageIndex(l.instance)
}

// SetImageIndex
// CN: 设置图像在images中的索引。
// EN: .
func (l *TListColumn) SetImageIndex(value int32) {
    ListColumn_SetImageIndex(l.instance, value)
}

// MaxWidth
func (l *TListColumn) MaxWidth() int32 {
    return ListColumn_GetMaxWidth(l.instance)
}

// SetMaxWidth
func (l *TListColumn) SetMaxWidth(value int32) {
    ListColumn_SetMaxWidth(l.instance, value)
}

// MinWidth
func (l *TListColumn) MinWidth() int32 {
    return ListColumn_GetMinWidth(l.instance)
}

// SetMinWidth
func (l *TListColumn) SetMinWidth(value int32) {
    ListColumn_SetMinWidth(l.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (l *TListColumn) Tag() int32 {
    return ListColumn_GetTag(l.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (l *TListColumn) SetTag(value int32) {
    ListColumn_SetTag(l.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (l *TListColumn) Width() int32 {
    return ListColumn_GetWidth(l.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (l *TListColumn) SetWidth(value int32) {
    ListColumn_SetWidth(l.instance, value)
}

// Collection
func (l *TListColumn) Collection() *TCollection {
    return CollectionFromInst(ListColumn_GetCollection(l.instance))
}

// SetCollection
func (l *TListColumn) SetCollection(value *TCollection) {
    ListColumn_SetCollection(l.instance, CheckPtr(value))
}

// Index
func (l *TListColumn) Index() int32 {
    return ListColumn_GetIndex(l.instance)
}

// SetIndex
func (l *TListColumn) SetIndex(value int32) {
    ListColumn_SetIndex(l.instance, value)
}

