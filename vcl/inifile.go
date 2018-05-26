
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
    "time"
	. "gitee.com/ying32/govcl/vcl/api"
)

type TIniFile struct {
    IObject
    instance uintptr
}

func NewIniFile(filename string) *TIniFile {
    i := new(TIniFile)
    i.instance = IniFile_Create(filename)
    return i
}

func IniFileFromInst(inst uintptr) *TIniFile {
    i := new(TIniFile)
    i.instance = inst
    return i
}

func IniFileFromObj(obj IObject) *TIniFile {
    i := new(TIniFile)
    i.instance = CheckPtr(obj)
    return i
}

func (i *TIniFile) Free() {
    if i.instance != 0 {
        IniFile_Free(i.instance)
        i.instance = 0
    }
}

func (i *TIniFile) Instance() uintptr {
    return i.instance
}

func (i *TIniFile) IsValid() bool {
    return i.instance != 0
}

func (i *TIniFile) ReadString(Section string, Ident string, Default string) string {
    return IniFile_ReadString(i.instance, Section , Ident , Default)
}

func (i *TIniFile) WriteString(Section string, Ident string, Value string) {
    IniFile_WriteString(i.instance, Section , Ident , Value)
}

func (i *TIniFile) ReadSections(Strings IObject) {
    IniFile_ReadSections(i.instance, CheckPtr(Strings))
}

func (i *TIniFile) ReadSectionValues(Section string, Strings IObject) {
    IniFile_ReadSectionValues(i.instance, Section , CheckPtr(Strings))
}

func (i *TIniFile) EraseSection(Section string) {
    IniFile_EraseSection(i.instance, Section)
}

func (i *TIniFile) DeleteKey(Section string, Ident string) {
    IniFile_DeleteKey(i.instance, Section , Ident)
}

func (i *TIniFile) UpdateFile() {
    IniFile_UpdateFile(i.instance)
}

func (i *TIniFile) SectionExists(Section string) bool {
    return IniFile_SectionExists(i.instance, Section)
}

func (i *TIniFile) ReadInteger(Section string, Ident string, Default int32) int32 {
    return IniFile_ReadInteger(i.instance, Section , Ident , Default)
}

func (i *TIniFile) WriteInteger(Section string, Ident string, Value int32) {
    IniFile_WriteInteger(i.instance, Section , Ident , Value)
}

func (i *TIniFile) ReadBool(Section string, Ident string, Default bool) bool {
    return IniFile_ReadBool(i.instance, Section , Ident , Default)
}

func (i *TIniFile) WriteBool(Section string, Ident string, Value bool) {
    IniFile_WriteBool(i.instance, Section , Ident , Value)
}

func (i *TIniFile) ReadDate(Section string, Name string, Default time.Time) time.Time {
    return IniFile_ReadDate(i.instance, Section , Name , Default)
}

func (i *TIniFile) ReadDateTime(Section string, Name string, Default time.Time) time.Time {
    return IniFile_ReadDateTime(i.instance, Section , Name , Default)
}

func (i *TIniFile) ReadFloat(Section string, Name string, Default float64) float64 {
    return IniFile_ReadFloat(i.instance, Section , Name , Default)
}

func (i *TIniFile) ReadTime(Section string, Name string, Default time.Time) time.Time {
    return IniFile_ReadTime(i.instance, Section , Name , Default)
}

func (i *TIniFile) WriteDate(Section string, Name string, Value time.Time) {
    IniFile_WriteDate(i.instance, Section , Name , Value)
}

func (i *TIniFile) WriteDateTime(Section string, Name string, Value time.Time) {
    IniFile_WriteDateTime(i.instance, Section , Name , Value)
}

func (i *TIniFile) WriteFloat(Section string, Name string, Value float64) {
    IniFile_WriteFloat(i.instance, Section , Name , Value)
}

func (i *TIniFile) WriteTime(Section string, Name string, Value time.Time) {
    IniFile_WriteTime(i.instance, Section , Name , Value)
}

func (i *TIniFile) ReadSubSections(Section string, Strings IObject, Recurse bool) {
    IniFile_ReadSubSections(i.instance, Section , CheckPtr(Strings), Recurse)
}

func (i *TIniFile) ValueExists(Section string, Ident string) bool {
    return IniFile_ValueExists(i.instance, Section , Ident)
}

func (i *TIniFile) ClassName() string {
    return IniFile_ClassName(i.instance)
}

func (i *TIniFile) Equals(Obj IObject) bool {
    return IniFile_Equals(i.instance, CheckPtr(Obj))
}

func (i *TIniFile) GetHashCode() int32 {
    return IniFile_GetHashCode(i.instance)
}

func (i *TIniFile) ToString() string {
    return IniFile_ToString(i.instance)
}

func (i *TIniFile) FileName() string {
    return IniFile_GetFileName(i.instance)
}

