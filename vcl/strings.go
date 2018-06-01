
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "github.com/ying32/govcl/vcl/api"
    . "github.com/ying32/govcl/vcl/types"
)

type TStrings struct {
    IObject
    instance uintptr
}

func NewStrings() *TStrings {
    s := new(TStrings)
    s.instance = Strings_Create()
    return s
}

func StringsFromInst(inst uintptr) *TStrings {
    s := new(TStrings)
    s.instance = inst
    return s
}

func StringsFromObj(obj IObject) *TStrings {
    s := new(TStrings)
    s.instance = CheckPtr(obj)
    return s
}

func (s *TStrings) Free() {
    if s.instance != 0 {
        Strings_Free(s.instance)
        s.instance = 0
    }
}

func (s *TStrings) Instance() uintptr {
    return s.instance
}

func (s *TStrings) IsValid() bool {
    return s.instance != 0
}

func (s *TStrings) Add(S string) int32 {
    return Strings_Add(s.instance, S)
}

func (s *TStrings) AddObject(S string, AObject IObject) int32 {
    return Strings_AddObject(s.instance, S , CheckPtr(AObject))
}

func (s *TStrings) Append(S string) {
    Strings_Append(s.instance, S)
}

func (s *TStrings) Assign(Source IObject) {
    Strings_Assign(s.instance, CheckPtr(Source))
}

func (s *TStrings) BeginUpdate() {
    Strings_BeginUpdate(s.instance)
}

func (s *TStrings) Clear() {
    Strings_Clear(s.instance)
}

func (s *TStrings) Delete(Index int32) {
    Strings_Delete(s.instance, Index)
}

func (s *TStrings) EndUpdate() {
    Strings_EndUpdate(s.instance)
}

func (s *TStrings) Equals(Strings IObject) bool {
    return Strings_Equals(s.instance, CheckPtr(Strings))
}

func (s *TStrings) IndexOf(S string) int32 {
    return Strings_IndexOf(s.instance, S)
}

func (s *TStrings) IndexOfName(Name string) int32 {
    return Strings_IndexOfName(s.instance, Name)
}

func (s *TStrings) IndexOfObject(AObject IObject) int32 {
    return Strings_IndexOfObject(s.instance, CheckPtr(AObject))
}

func (s *TStrings) Insert(Index int32, S string) {
    Strings_Insert(s.instance, Index , S)
}

func (s *TStrings) InsertObject(Index int32, S string, AObject IObject) {
    Strings_InsertObject(s.instance, Index , S , CheckPtr(AObject))
}

func (s *TStrings) LoadFromFile(FileName string) {
    Strings_LoadFromFile(s.instance, FileName)
}

func (s *TStrings) LoadFromStream(Stream IObject) {
    Strings_LoadFromStream(s.instance, CheckPtr(Stream))
}

func (s *TStrings) Move(CurIndex int32, NewIndex int32) {
    Strings_Move(s.instance, CurIndex , NewIndex)
}

func (s *TStrings) SaveToFile(FileName string) {
    Strings_SaveToFile(s.instance, FileName)
}

func (s *TStrings) SaveToStream(Stream IObject) {
    Strings_SaveToStream(s.instance, CheckPtr(Stream))
}

func (s *TStrings) GetNamePath() string {
    return Strings_GetNamePath(s.instance)
}

func (s *TStrings) ClassName() string {
    return Strings_ClassName(s.instance)
}

func (s *TStrings) GetHashCode() int32 {
    return Strings_GetHashCode(s.instance)
}

func (s *TStrings) ToString() string {
    return Strings_ToString(s.instance)
}

func (s *TStrings) Capacity() int32 {
    return Strings_GetCapacity(s.instance)
}

func (s *TStrings) SetCapacity(value int32) {
    Strings_SetCapacity(s.instance, value)
}

func (s *TStrings) CommaText() string {
    return Strings_GetCommaText(s.instance)
}

func (s *TStrings) SetCommaText(value string) {
    Strings_SetCommaText(s.instance, value)
}

func (s *TStrings) Count() int32 {
    return Strings_GetCount(s.instance)
}

func (s *TStrings) Delimiter() uint16 {
    return Strings_GetDelimiter(s.instance)
}

func (s *TStrings) SetDelimiter(value uint16) {
    Strings_SetDelimiter(s.instance, value)
}

func (s *TStrings) Text() string {
    return Strings_GetText(s.instance)
}

func (s *TStrings) SetText(value string) {
    Strings_SetText(s.instance, value)
}

func (s *TStrings) WriteBOM() bool {
    return Strings_GetWriteBOM(s.instance)
}

func (s *TStrings) SetWriteBOM(value bool) {
    Strings_SetWriteBOM(s.instance, value)
}

func (s *TStrings) Options() TStringsOptions {
    return Strings_GetOptions(s.instance)
}

func (s *TStrings) SetOptions(value TStringsOptions) {
    Strings_SetOptions(s.instance, value)
}

func (s *TStrings) Objects(Index int32) *TObject {
    return ObjectFromInst(Strings_GetObjects(s.instance, Index))
}

func (s *TStrings) SetObjects(Index int32, value IObject) {
    Strings_SetObjects(s.instance, Index, CheckPtr(value))
}

func (s *TStrings) Values(Name string) string {
    return Strings_GetValues(s.instance, Name)
}

func (s *TStrings) SetValues(Name string, value string) {
    Strings_SetValues(s.instance, Name, value)
}

func (s *TStrings) ValueFromIndex(Index int32) string {
    return Strings_GetValueFromIndex(s.instance, Index)
}

func (s *TStrings) SetValueFromIndex(Index int32, value string) {
    Strings_SetValueFromIndex(s.instance, Index, value)
}

func (s *TStrings) Strings(Index int32) string {
    return Strings_GetStrings(s.instance, Index)
}

func (s *TStrings) SetStrings(Index int32, value string) {
    Strings_SetStrings(s.instance, Index, value)
}

