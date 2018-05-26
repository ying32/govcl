
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

type TStringList struct {
    IObject
    instance uintptr
}

func NewStringList() *TStringList {
    s := new(TStringList)
    s.instance = StringList_Create()
    return s
}

func StringListFromInst(inst uintptr) *TStringList {
    s := new(TStringList)
    s.instance = inst
    return s
}

func StringListFromObj(obj IObject) *TStringList {
    s := new(TStringList)
    s.instance = CheckPtr(obj)
    return s
}

func (s *TStringList) Free() {
    if s.instance != 0 {
        StringList_Free(s.instance)
        s.instance = 0
    }
}

func (s *TStringList) Instance() uintptr {
    return s.instance
}

func (s *TStringList) IsValid() bool {
    return s.instance != 0
}

func (s *TStringList) Add(S string) int32 {
    return StringList_Add(s.instance, S)
}

func (s *TStringList) AddObject(S string, AObject IObject) int32 {
    return StringList_AddObject(s.instance, S , CheckPtr(AObject))
}

func (s *TStringList) Assign(Source IObject) {
    StringList_Assign(s.instance, CheckPtr(Source))
}

func (s *TStringList) Clear() {
    StringList_Clear(s.instance)
}

func (s *TStringList) Delete(Index int32) {
    StringList_Delete(s.instance, Index)
}

func (s *TStringList) IndexOf(S string) int32 {
    return StringList_IndexOf(s.instance, S)
}

func (s *TStringList) Insert(Index int32, S string) {
    StringList_Insert(s.instance, Index , S)
}

func (s *TStringList) InsertObject(Index int32, S string, AObject IObject) {
    StringList_InsertObject(s.instance, Index , S , CheckPtr(AObject))
}

func (s *TStringList) Append(S string) {
    StringList_Append(s.instance, S)
}

func (s *TStringList) BeginUpdate() {
    StringList_BeginUpdate(s.instance)
}

func (s *TStringList) EndUpdate() {
    StringList_EndUpdate(s.instance)
}

func (s *TStringList) Equals(Strings IObject) bool {
    return StringList_Equals(s.instance, CheckPtr(Strings))
}

func (s *TStringList) IndexOfName(Name string) int32 {
    return StringList_IndexOfName(s.instance, Name)
}

func (s *TStringList) IndexOfObject(AObject IObject) int32 {
    return StringList_IndexOfObject(s.instance, CheckPtr(AObject))
}

func (s *TStringList) LoadFromFile(FileName string) {
    StringList_LoadFromFile(s.instance, FileName)
}

func (s *TStringList) LoadFromStream(Stream IObject) {
    StringList_LoadFromStream(s.instance, CheckPtr(Stream))
}

func (s *TStringList) Move(CurIndex int32, NewIndex int32) {
    StringList_Move(s.instance, CurIndex , NewIndex)
}

func (s *TStringList) SaveToFile(FileName string) {
    StringList_SaveToFile(s.instance, FileName)
}

func (s *TStringList) SaveToStream(Stream IObject) {
    StringList_SaveToStream(s.instance, CheckPtr(Stream))
}

func (s *TStringList) GetNamePath() string {
    return StringList_GetNamePath(s.instance)
}

func (s *TStringList) ClassName() string {
    return StringList_ClassName(s.instance)
}

func (s *TStringList) GetHashCode() int32 {
    return StringList_GetHashCode(s.instance)
}

func (s *TStringList) ToString() string {
    return StringList_ToString(s.instance)
}

func (s *TStringList) Sorted() bool {
    return StringList_GetSorted(s.instance)
}

func (s *TStringList) SetSorted(value bool) {
    StringList_SetSorted(s.instance, value)
}

func (s *TStringList) SetOnChange(fn TNotifyEvent) {
    StringList_SetOnChange(s.instance, fn)
}

func (s *TStringList) Capacity() int32 {
    return StringList_GetCapacity(s.instance)
}

func (s *TStringList) SetCapacity(value int32) {
    StringList_SetCapacity(s.instance, value)
}

func (s *TStringList) CommaText() string {
    return StringList_GetCommaText(s.instance)
}

func (s *TStringList) SetCommaText(value string) {
    StringList_SetCommaText(s.instance, value)
}

func (s *TStringList) Count() int32 {
    return StringList_GetCount(s.instance)
}

func (s *TStringList) Delimiter() uint16 {
    return StringList_GetDelimiter(s.instance)
}

func (s *TStringList) SetDelimiter(value uint16) {
    StringList_SetDelimiter(s.instance, value)
}

func (s *TStringList) Text() string {
    return StringList_GetText(s.instance)
}

func (s *TStringList) SetText(value string) {
    StringList_SetText(s.instance, value)
}

func (s *TStringList) WriteBOM() bool {
    return StringList_GetWriteBOM(s.instance)
}

func (s *TStringList) SetWriteBOM(value bool) {
    StringList_SetWriteBOM(s.instance, value)
}

func (s *TStringList) Options() TStringsOptions {
    return StringList_GetOptions(s.instance)
}

func (s *TStringList) SetOptions(value TStringsOptions) {
    StringList_SetOptions(s.instance, value)
}

func (s *TStringList) Objects(Index int32) *TObject {
    return ObjectFromInst(StringList_GetObjects(s.instance, Index))
}

func (s *TStringList) SetObjects(Index int32, value IObject) {
    StringList_SetObjects(s.instance, Index, CheckPtr(value))
}

func (s *TStringList) Values(Name string) string {
    return StringList_GetValues(s.instance, Name)
}

func (s *TStringList) SetValues(Name string, value string) {
    StringList_SetValues(s.instance, Name, value)
}

func (s *TStringList) ValueFromIndex(Index int32) string {
    return StringList_GetValueFromIndex(s.instance, Index)
}

func (s *TStringList) SetValueFromIndex(Index int32, value string) {
    StringList_SetValueFromIndex(s.instance, Index, value)
}

func (s *TStringList) Strings(Index int32) string {
    return StringList_GetStrings(s.instance, Index)
}

func (s *TStringList) SetStrings(Index int32, value string) {
    StringList_SetStrings(s.instance, Index, value)
}

