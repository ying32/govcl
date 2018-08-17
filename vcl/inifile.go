
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
    "time"
	. "github.com/ying32/govcl/vcl/api"
    . "github.com/ying32/govcl/vcl/types"
    "unsafe"
)

type TIniFile struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewIniFile
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewIniFile(filename string) *TIniFile {
    i := new(TIniFile)
    i.instance = IniFile_Create(filename)
    i.ptr = unsafe.Pointer(i.instance)
    return i
}

// IniFileFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func IniFileFromInst(inst uintptr) *TIniFile {
    i := new(TIniFile)
    i.instance = inst
    i.ptr = unsafe.Pointer(inst)
    return i
}

// IniFileFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func IniFileFromObj(obj IObject) *TIniFile {
    i := new(TIniFile)
    i.instance = CheckPtr(obj)
    i.ptr = unsafe.Pointer(i.instance)
    return i
}

// IniFileFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func IniFileFromUnsafePointer(ptr unsafe.Pointer) *TIniFile {
    i := new(TIniFile)
    i.instance = uintptr(ptr)
    i.ptr = ptr
    return i
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (i *TIniFile) Free() {
    if i.instance != 0 {
        IniFile_Free(i.instance)
        i.instance = 0
        i.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (i *TIniFile) Instance() uintptr {
    return i.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (i *TIniFile) UnsafeAddr() unsafe.Pointer {
    return i.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (i *TIniFile) IsValid() bool {
    return i.instance != 0
}

// TIniFileClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TIniFileClass() TClass {
    return IniFile_StaticClassType()
}

// ReadString
func (i *TIniFile) ReadString(Section string, Ident string, Default string) string {
    return IniFile_ReadString(i.instance, Section , Ident , Default)
}

// WriteString
func (i *TIniFile) WriteString(Section string, Ident string, Value string) {
    IniFile_WriteString(i.instance, Section , Ident , Value)
}

// ReadSections
func (i *TIniFile) ReadSections(Strings IObject) {
    IniFile_ReadSections(i.instance, CheckPtr(Strings))
}

// ReadSectionValues
func (i *TIniFile) ReadSectionValues(Section string, Strings IObject) {
    IniFile_ReadSectionValues(i.instance, Section , CheckPtr(Strings))
}

// EraseSection
func (i *TIniFile) EraseSection(Section string) {
    IniFile_EraseSection(i.instance, Section)
}

// DeleteKey
func (i *TIniFile) DeleteKey(Section string, Ident string) {
    IniFile_DeleteKey(i.instance, Section , Ident)
}

// UpdateFile
func (i *TIniFile) UpdateFile() {
    IniFile_UpdateFile(i.instance)
}

// SectionExists
func (i *TIniFile) SectionExists(Section string) bool {
    return IniFile_SectionExists(i.instance, Section)
}

// ReadInteger
func (i *TIniFile) ReadInteger(Section string, Ident string, Default int32) int32 {
    return IniFile_ReadInteger(i.instance, Section , Ident , Default)
}

// WriteInteger
func (i *TIniFile) WriteInteger(Section string, Ident string, Value int32) {
    IniFile_WriteInteger(i.instance, Section , Ident , Value)
}

// ReadBool
func (i *TIniFile) ReadBool(Section string, Ident string, Default bool) bool {
    return IniFile_ReadBool(i.instance, Section , Ident , Default)
}

// WriteBool
func (i *TIniFile) WriteBool(Section string, Ident string, Value bool) {
    IniFile_WriteBool(i.instance, Section , Ident , Value)
}

// ReadDate
func (i *TIniFile) ReadDate(Section string, Name string, Default time.Time) time.Time {
    return IniFile_ReadDate(i.instance, Section , Name , Default)
}

// ReadDateTime
func (i *TIniFile) ReadDateTime(Section string, Name string, Default time.Time) time.Time {
    return IniFile_ReadDateTime(i.instance, Section , Name , Default)
}

// ReadFloat
func (i *TIniFile) ReadFloat(Section string, Name string, Default float64) float64 {
    return IniFile_ReadFloat(i.instance, Section , Name , Default)
}

// ReadTime
func (i *TIniFile) ReadTime(Section string, Name string, Default time.Time) time.Time {
    return IniFile_ReadTime(i.instance, Section , Name , Default)
}

// WriteDate
func (i *TIniFile) WriteDate(Section string, Name string, Value time.Time) {
    IniFile_WriteDate(i.instance, Section , Name , Value)
}

// WriteDateTime
func (i *TIniFile) WriteDateTime(Section string, Name string, Value time.Time) {
    IniFile_WriteDateTime(i.instance, Section , Name , Value)
}

// WriteFloat
func (i *TIniFile) WriteFloat(Section string, Name string, Value float64) {
    IniFile_WriteFloat(i.instance, Section , Name , Value)
}

// WriteTime
func (i *TIniFile) WriteTime(Section string, Name string, Value time.Time) {
    IniFile_WriteTime(i.instance, Section , Name , Value)
}

// ReadSubSections
func (i *TIniFile) ReadSubSections(Section string, Strings IObject, Recurse bool) {
    IniFile_ReadSubSections(i.instance, Section , CheckPtr(Strings), Recurse)
}

// ValueExists
func (i *TIniFile) ValueExists(Section string, Ident string) bool {
    return IniFile_ValueExists(i.instance, Section , Ident)
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (i *TIniFile) DisposeOf() {
    IniFile_DisposeOf(i.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (i *TIniFile) ClassType() TClass {
    return IniFile_ClassType(i.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (i *TIniFile) ClassName() string {
    return IniFile_ClassName(i.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (i *TIniFile) InstanceSize() int32 {
    return IniFile_InstanceSize(i.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (i *TIniFile) InheritsFrom(AClass TClass) bool {
    return IniFile_InheritsFrom(i.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (i *TIniFile) Equals(Obj IObject) bool {
    return IniFile_Equals(i.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (i *TIniFile) GetHashCode() int32 {
    return IniFile_GetHashCode(i.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (i *TIniFile) ToString() string {
    return IniFile_ToString(i.instance)
}

// FileName
func (i *TIniFile) FileName() string {
    return IniFile_GetFileName(i.instance)
}

