
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

type TListGroup struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewListGroup
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewListGroup() *TListGroup {
    l := new(TListGroup)
    l.instance = ListGroup_Create()
    l.ptr = unsafe.Pointer(l.instance)
    return l
}

// ListGroupFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func ListGroupFromInst(inst uintptr) *TListGroup {
    l := new(TListGroup)
    l.instance = inst
    l.ptr = unsafe.Pointer(inst)
    return l
}

// ListGroupFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func ListGroupFromObj(obj IObject) *TListGroup {
    l := new(TListGroup)
    l.instance = CheckPtr(obj)
    l.ptr = unsafe.Pointer(l.instance)
    return l
}

// ListGroupFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func ListGroupFromUnsafePointer(ptr unsafe.Pointer) *TListGroup {
    l := new(TListGroup)
    l.instance = uintptr(ptr)
    l.ptr = ptr
    return l
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (l *TListGroup) Free() {
    if l.instance != 0 {
        ListGroup_Free(l.instance)
        l.instance = 0
        l.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (l *TListGroup) Instance() uintptr {
    return l.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (l *TListGroup) UnsafeAddr() unsafe.Pointer {
    return l.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (l *TListGroup) IsValid() bool {
    return l.instance != 0
}

// TListGroupClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TListGroupClass() TClass {
    return ListGroup_StaticClassType()
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (l *TListGroup) Assign(Source IObject) {
    ListGroup_Assign(l.instance, CheckPtr(Source))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (l *TListGroup) GetNamePath() string {
    return ListGroup_GetNamePath(l.instance)
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (l *TListGroup) DisposeOf() {
    ListGroup_DisposeOf(l.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (l *TListGroup) ClassType() TClass {
    return ListGroup_ClassType(l.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (l *TListGroup) ClassName() string {
    return ListGroup_ClassName(l.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (l *TListGroup) InstanceSize() int32 {
    return ListGroup_InstanceSize(l.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (l *TListGroup) InheritsFrom(AClass TClass) bool {
    return ListGroup_InheritsFrom(l.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (l *TListGroup) Equals(Obj IObject) bool {
    return ListGroup_Equals(l.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (l *TListGroup) GetHashCode() int32 {
    return ListGroup_GetHashCode(l.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (l *TListGroup) ToString() string {
    return ListGroup_ToString(l.instance)
}

// Header
func (l *TListGroup) Header() string {
    return ListGroup_GetHeader(l.instance)
}

// SetHeader
func (l *TListGroup) SetHeader(value string) {
    ListGroup_SetHeader(l.instance, value)
}

// Footer
func (l *TListGroup) Footer() string {
    return ListGroup_GetFooter(l.instance)
}

// SetFooter
func (l *TListGroup) SetFooter(value string) {
    ListGroup_SetFooter(l.instance, value)
}

// GroupID
func (l *TListGroup) GroupID() int32 {
    return ListGroup_GetGroupID(l.instance)
}

// SetGroupID
func (l *TListGroup) SetGroupID(value int32) {
    ListGroup_SetGroupID(l.instance, value)
}

// State
func (l *TListGroup) State() TListGroupStateSet {
    return ListGroup_GetState(l.instance)
}

// SetState
func (l *TListGroup) SetState(value TListGroupStateSet) {
    ListGroup_SetState(l.instance, value)
}

// HeaderAlign
func (l *TListGroup) HeaderAlign() TAlignment {
    return ListGroup_GetHeaderAlign(l.instance)
}

// SetHeaderAlign
func (l *TListGroup) SetHeaderAlign(value TAlignment) {
    ListGroup_SetHeaderAlign(l.instance, value)
}

// FooterAlign
func (l *TListGroup) FooterAlign() TAlignment {
    return ListGroup_GetFooterAlign(l.instance)
}

// SetFooterAlign
func (l *TListGroup) SetFooterAlign(value TAlignment) {
    ListGroup_SetFooterAlign(l.instance, value)
}

// Subtitle
func (l *TListGroup) Subtitle() string {
    return ListGroup_GetSubtitle(l.instance)
}

// SetSubtitle
func (l *TListGroup) SetSubtitle(value string) {
    ListGroup_SetSubtitle(l.instance, value)
}

// TitleImage
func (l *TListGroup) TitleImage() int32 {
    return ListGroup_GetTitleImage(l.instance)
}

// SetTitleImage
func (l *TListGroup) SetTitleImage(value int32) {
    ListGroup_SetTitleImage(l.instance, value)
}

// Collection
func (l *TListGroup) Collection() *TCollection {
    return CollectionFromInst(ListGroup_GetCollection(l.instance))
}

// SetCollection
func (l *TListGroup) SetCollection(value *TCollection) {
    ListGroup_SetCollection(l.instance, CheckPtr(value))
}

// Index
func (l *TListGroup) Index() int32 {
    return ListGroup_GetIndex(l.instance)
}

// SetIndex
func (l *TListGroup) SetIndex(value int32) {
    ListGroup_SetIndex(l.instance, value)
}

// DisplayName
func (l *TListGroup) DisplayName() string {
    return ListGroup_GetDisplayName(l.instance)
}

// SetDisplayName
func (l *TListGroup) SetDisplayName(value string) {
    ListGroup_SetDisplayName(l.instance, value)
}

