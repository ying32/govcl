
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

type TStringList struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewStringList
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewStringList() *TStringList {
    s := new(TStringList)
    s.instance = StringList_Create()
    s.ptr = unsafe.Pointer(s.instance)
    return s
}

// StringListFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func StringListFromInst(inst uintptr) *TStringList {
    s := new(TStringList)
    s.instance = inst
    s.ptr = unsafe.Pointer(inst)
    return s
}

// StringListFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func StringListFromObj(obj IObject) *TStringList {
    s := new(TStringList)
    s.instance = CheckPtr(obj)
    s.ptr = unsafe.Pointer(s.instance)
    return s
}

// StringListFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func StringListFromUnsafePointer(ptr unsafe.Pointer) *TStringList {
    s := new(TStringList)
    s.instance = uintptr(ptr)
    s.ptr = ptr
    return s
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (s *TStringList) Free() {
    if s.instance != 0 {
        StringList_Free(s.instance)
        s.instance = 0
        s.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (s *TStringList) Instance() uintptr {
    return s.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (s *TStringList) UnsafeAddr() unsafe.Pointer {
    return s.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (s *TStringList) IsValid() bool {
    return s.instance != 0
}

// TStringListClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TStringListClass() TClass {
    return StringList_StaticClassType()
}

// Add
func (s *TStringList) Add(S string) int32 {
    return StringList_Add(s.instance, S)
}

// AddObject
func (s *TStringList) AddObject(S string, AObject IObject) int32 {
    return StringList_AddObject(s.instance, S , CheckPtr(AObject))
}

// Assign
func (s *TStringList) Assign(Source IObject) {
    StringList_Assign(s.instance, CheckPtr(Source))
}

// Clear
func (s *TStringList) Clear() {
    StringList_Clear(s.instance)
}

// Delete
func (s *TStringList) Delete(Index int32) {
    StringList_Delete(s.instance, Index)
}

// IndexOf
func (s *TStringList) IndexOf(S string) int32 {
    return StringList_IndexOf(s.instance, S)
}

// Insert
func (s *TStringList) Insert(Index int32, S string) {
    StringList_Insert(s.instance, Index , S)
}

// InsertObject
func (s *TStringList) InsertObject(Index int32, S string, AObject IObject) {
    StringList_InsertObject(s.instance, Index , S , CheckPtr(AObject))
}

// Append
func (s *TStringList) Append(S string) {
    StringList_Append(s.instance, S)
}

// BeginUpdate
func (s *TStringList) BeginUpdate() {
    StringList_BeginUpdate(s.instance)
}

// EndUpdate
func (s *TStringList) EndUpdate() {
    StringList_EndUpdate(s.instance)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (s *TStringList) Equals(Strings IObject) bool {
    return StringList_Equals(s.instance, CheckPtr(Strings))
}

// IndexOfName
func (s *TStringList) IndexOfName(Name string) int32 {
    return StringList_IndexOfName(s.instance, Name)
}

// IndexOfObject
func (s *TStringList) IndexOfObject(AObject IObject) int32 {
    return StringList_IndexOfObject(s.instance, CheckPtr(AObject))
}

// LoadFromFile
func (s *TStringList) LoadFromFile(FileName string) {
    StringList_LoadFromFile(s.instance, FileName)
}

// LoadFromStream
func (s *TStringList) LoadFromStream(Stream IObject) {
    StringList_LoadFromStream(s.instance, CheckPtr(Stream))
}

// Move
func (s *TStringList) Move(CurIndex int32, NewIndex int32) {
    StringList_Move(s.instance, CurIndex , NewIndex)
}

// SaveToFile
func (s *TStringList) SaveToFile(FileName string) {
    StringList_SaveToFile(s.instance, FileName)
}

// SaveToStream
func (s *TStringList) SaveToStream(Stream IObject) {
    StringList_SaveToStream(s.instance, CheckPtr(Stream))
}

// GetNamePath
func (s *TStringList) GetNamePath() string {
    return StringList_GetNamePath(s.instance)
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (s *TStringList) DisposeOf() {
    StringList_DisposeOf(s.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (s *TStringList) ClassType() TClass {
    return StringList_ClassType(s.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (s *TStringList) ClassName() string {
    return StringList_ClassName(s.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (s *TStringList) InstanceSize() int32 {
    return StringList_InstanceSize(s.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (s *TStringList) InheritsFrom(AClass TClass) bool {
    return StringList_InheritsFrom(s.instance, AClass)
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (s *TStringList) GetHashCode() int32 {
    return StringList_GetHashCode(s.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (s *TStringList) ToString() string {
    return StringList_ToString(s.instance)
}

// Sorted
func (s *TStringList) Sorted() bool {
    return StringList_GetSorted(s.instance)
}

// SetSorted
func (s *TStringList) SetSorted(value bool) {
    StringList_SetSorted(s.instance, value)
}

// SetOnChange
func (s *TStringList) SetOnChange(fn TNotifyEvent) {
    StringList_SetOnChange(s.instance, fn)
}

// Capacity
func (s *TStringList) Capacity() int32 {
    return StringList_GetCapacity(s.instance)
}

// SetCapacity
func (s *TStringList) SetCapacity(value int32) {
    StringList_SetCapacity(s.instance, value)
}

// CommaText
func (s *TStringList) CommaText() string {
    return StringList_GetCommaText(s.instance)
}

// SetCommaText
func (s *TStringList) SetCommaText(value string) {
    StringList_SetCommaText(s.instance, value)
}

// Count
func (s *TStringList) Count() int32 {
    return StringList_GetCount(s.instance)
}

// Delimiter
func (s *TStringList) Delimiter() uint16 {
    return StringList_GetDelimiter(s.instance)
}

// SetDelimiter
func (s *TStringList) SetDelimiter(value uint16) {
    StringList_SetDelimiter(s.instance, value)
}

// Text
func (s *TStringList) Text() string {
    return StringList_GetText(s.instance)
}

// SetText
func (s *TStringList) SetText(value string) {
    StringList_SetText(s.instance, value)
}

// WriteBOM
func (s *TStringList) WriteBOM() bool {
    return StringList_GetWriteBOM(s.instance)
}

// SetWriteBOM
func (s *TStringList) SetWriteBOM(value bool) {
    StringList_SetWriteBOM(s.instance, value)
}

// Options
func (s *TStringList) Options() TStringsOptions {
    return StringList_GetOptions(s.instance)
}

// SetOptions
func (s *TStringList) SetOptions(value TStringsOptions) {
    StringList_SetOptions(s.instance, value)
}

// Objects
func (s *TStringList) Objects(Index int32) *TObject {
    return ObjectFromInst(StringList_GetObjects(s.instance, Index))
}

// Objects
func (s *TStringList) SetObjects(Index int32, value IObject) {
    StringList_SetObjects(s.instance, Index, CheckPtr(value))
}

// Values
func (s *TStringList) Values(Name string) string {
    return StringList_GetValues(s.instance, Name)
}

// Values
func (s *TStringList) SetValues(Name string, value string) {
    StringList_SetValues(s.instance, Name, value)
}

// ValueFromIndex
func (s *TStringList) ValueFromIndex(Index int32) string {
    return StringList_GetValueFromIndex(s.instance, Index)
}

// ValueFromIndex
func (s *TStringList) SetValueFromIndex(Index int32, value string) {
    StringList_SetValueFromIndex(s.instance, Index, value)
}

// Strings
func (s *TStringList) Strings(Index int32) string {
    return StringList_GetStrings(s.instance, Index)
}

// Strings
func (s *TStringList) SetStrings(Index int32, value string) {
    StringList_SetStrings(s.instance, Index, value)
}

