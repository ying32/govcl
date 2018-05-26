
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
    "time"
	. "gitee.com/ying32/govcl/vcl/api"
    . "gitee.com/ying32/govcl/vcl/types"
)

type TRegistry struct {
    IObject
    instance uintptr
}

func NewRegistry(aAccess uint32) *TRegistry {
    r := new(TRegistry)
    r.instance = Registry_Create(aAccess)
    return r
}

func RegistryFromInst(inst uintptr) *TRegistry {
    r := new(TRegistry)
    r.instance = inst
    return r
}

func RegistryFromObj(obj IObject) *TRegistry {
    r := new(TRegistry)
    r.instance = CheckPtr(obj)
    return r
}

func (r *TRegistry) Free() {
    if r.instance != 0 {
        Registry_Free(r.instance)
        r.instance = 0
    }
}

func (r *TRegistry) Instance() uintptr {
    return r.instance
}

func (r *TRegistry) IsValid() bool {
    return r.instance != 0
}

func (r *TRegistry) CloseKey() {
    Registry_CloseKey(r.instance)
}

func (r *TRegistry) CreateKey(Key string) bool {
    return Registry_CreateKey(r.instance, Key)
}

func (r *TRegistry) DeleteKey(Key string) bool {
    return Registry_DeleteKey(r.instance, Key)
}

func (r *TRegistry) DeleteValue(Name string) bool {
    return Registry_DeleteValue(r.instance, Name)
}

func (r *TRegistry) HasSubKeys() bool {
    return Registry_HasSubKeys(r.instance)
}

func (r *TRegistry) KeyExists(Key string) bool {
    return Registry_KeyExists(r.instance, Key)
}

func (r *TRegistry) LoadKey(Key string, FileName string) bool {
    return Registry_LoadKey(r.instance, Key , FileName)
}

func (r *TRegistry) MoveKey(OldName string, NewName string, Delete bool) {
    Registry_MoveKey(r.instance, OldName , NewName , Delete)
}

func (r *TRegistry) OpenKey(Key string, CanCreate bool) bool {
    return Registry_OpenKey(r.instance, Key , CanCreate)
}

func (r *TRegistry) OpenKeyReadOnly(Key string) bool {
    return Registry_OpenKeyReadOnly(r.instance, Key)
}

func (r *TRegistry) ReadBool(Name string) bool {
    return Registry_ReadBool(r.instance, Name)
}

func (r *TRegistry) ReadDate(Name string) time.Time {
    return Registry_ReadDate(r.instance, Name)
}

func (r *TRegistry) ReadDateTime(Name string) time.Time {
    return Registry_ReadDateTime(r.instance, Name)
}

func (r *TRegistry) ReadFloat(Name string) float64 {
    return Registry_ReadFloat(r.instance, Name)
}

func (r *TRegistry) ReadInteger(Name string) int32 {
    return Registry_ReadInteger(r.instance, Name)
}

func (r *TRegistry) ReadString(Name string) string {
    return Registry_ReadString(r.instance, Name)
}

func (r *TRegistry) ReadTime(Name string) time.Time {
    return Registry_ReadTime(r.instance, Name)
}

func (r *TRegistry) RegistryConnect(UNCName string) bool {
    return Registry_RegistryConnect(r.instance, UNCName)
}

func (r *TRegistry) RenameValue(OldName string, NewName string) {
    Registry_RenameValue(r.instance, OldName , NewName)
}

func (r *TRegistry) ReplaceKey(Key string, FileName string, BackUpFileName string) bool {
    return Registry_ReplaceKey(r.instance, Key , FileName , BackUpFileName)
}

func (r *TRegistry) RestoreKey(Key string, FileName string) bool {
    return Registry_RestoreKey(r.instance, Key , FileName)
}

func (r *TRegistry) SaveKey(Key string, FileName string) bool {
    return Registry_SaveKey(r.instance, Key , FileName)
}

func (r *TRegistry) UnLoadKey(Key string) bool {
    return Registry_UnLoadKey(r.instance, Key)
}

func (r *TRegistry) ValueExists(Name string) bool {
    return Registry_ValueExists(r.instance, Name)
}

func (r *TRegistry) WriteBool(Name string, Value bool) {
    Registry_WriteBool(r.instance, Name , Value)
}

func (r *TRegistry) WriteDate(Name string, Value time.Time) {
    Registry_WriteDate(r.instance, Name , Value)
}

func (r *TRegistry) WriteDateTime(Name string, Value time.Time) {
    Registry_WriteDateTime(r.instance, Name , Value)
}

func (r *TRegistry) WriteFloat(Name string, Value float64) {
    Registry_WriteFloat(r.instance, Name , Value)
}

func (r *TRegistry) WriteInteger(Name string, Value int32) {
    Registry_WriteInteger(r.instance, Name , Value)
}

func (r *TRegistry) WriteString(Name string, Value string) {
    Registry_WriteString(r.instance, Name , Value)
}

func (r *TRegistry) WriteExpandString(Name string, Value string) {
    Registry_WriteExpandString(r.instance, Name , Value)
}

func (r *TRegistry) WriteTime(Name string, Value time.Time) {
    Registry_WriteTime(r.instance, Name , Value)
}

func (r *TRegistry) ClassName() string {
    return Registry_ClassName(r.instance)
}

func (r *TRegistry) Equals(Obj IObject) bool {
    return Registry_Equals(r.instance, CheckPtr(Obj))
}

func (r *TRegistry) GetHashCode() int32 {
    return Registry_GetHashCode(r.instance)
}

func (r *TRegistry) ToString() string {
    return Registry_ToString(r.instance)
}

func (r *TRegistry) CurrentKey() HKEY {
    return Registry_GetCurrentKey(r.instance)
}

func (r *TRegistry) CurrentPath() string {
    return Registry_GetCurrentPath(r.instance)
}

func (r *TRegistry) LazyWrite() bool {
    return Registry_GetLazyWrite(r.instance)
}

func (r *TRegistry) SetLazyWrite(value bool) {
    Registry_SetLazyWrite(r.instance, value)
}

func (r *TRegistry) LastError() int32 {
    return Registry_GetLastError(r.instance)
}

func (r *TRegistry) LastErrorMsg() string {
    return Registry_GetLastErrorMsg(r.instance)
}

func (r *TRegistry) RootKey() HKEY {
    return Registry_GetRootKey(r.instance)
}

func (r *TRegistry) SetRootKey(value HKEY) {
    Registry_SetRootKey(r.instance, value)
}

func (r *TRegistry) RootKeyName() string {
    return Registry_GetRootKeyName(r.instance)
}

func (r *TRegistry) Access() uint32 {
    return Registry_GetAccess(r.instance)
}

func (r *TRegistry) SetAccess(value uint32) {
    Registry_SetAccess(r.instance, value)
}

