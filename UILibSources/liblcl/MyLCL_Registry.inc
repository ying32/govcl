
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function Registry_Create(AAccess: LongWord): TRegistry; extdecl;
begin
  Result :=  TRegistry.Create(AAccess);
end;

procedure Registry_Free(AObj: TRegistry); extdecl;
begin
  AObj.Free;
end;

procedure Registry_CloseKey(AObj: TRegistry); extdecl;
begin
  AObj.CloseKey;
end;

function Registry_CreateKey(AObj: TRegistry; Key: PChar): LongBool; extdecl;
begin
  Result :=  AObj.CreateKey(Key);
end;

function Registry_DeleteKey(AObj: TRegistry; Key: PChar): LongBool; extdecl;
begin
  Result :=  AObj.DeleteKey(Key);
end;

function Registry_DeleteValue(AObj: TRegistry; Name: PChar): LongBool; extdecl;
begin
  Result :=  AObj.DeleteValue(Name);
end;

function Registry_HasSubKeys(AObj: TRegistry): LongBool; extdecl;
begin
  Result :=  AObj.HasSubKeys;
end;

function Registry_KeyExists(AObj: TRegistry; Key: PChar): LongBool; extdecl;
begin
  Result :=  AObj.KeyExists(Key);
end;

function Registry_LoadKey(AObj: TRegistry; Key: PChar; FileName: PChar): LongBool; extdecl;
begin
  Result :=  AObj.LoadKey(Key, FileName);
end;

procedure Registry_MoveKey(AObj: TRegistry; OldName: PChar; NewName: PChar; Delete: LongBool); extdecl;
begin
  AObj.MoveKey(OldName, NewName, Delete);
end;

function Registry_OpenKey(AObj: TRegistry; Key: PChar; CanCreate: LongBool): LongBool; extdecl;
begin
  Result :=  AObj.OpenKey(Key, CanCreate);
end;

function Registry_OpenKeyReadOnly(AObj: TRegistry; Key: PChar): LongBool; extdecl;
begin
  Result :=  AObj.OpenKeyReadOnly(Key);
end;

function Registry_ReadBool(AObj: TRegistry; Name: PChar): LongBool; extdecl;
begin
  Result :=  AObj.ReadBool(Name);
end;

function Registry_ReadDate(AObj: TRegistry; Name: PChar): TUnixDateTime; extdecl;
begin
  Result :=  ToUnixTime(AObj.ReadDate(Name));
end;

function Registry_ReadDateTime(AObj: TRegistry; Name: PChar): TUnixDateTime; extdecl;
begin
  Result :=  ToUnixTime(AObj.ReadDateTime(Name));
end;

procedure Registry_ReadFloat(AObj: TRegistry; Name: PChar; var Result: Double); extdecl;
begin
  Result :=  AObj.ReadFloat(Name);
end;

function Registry_ReadInteger(AObj: TRegistry; Name: PChar): Integer; extdecl;
begin
  Result :=  AObj.ReadInteger(Name);
end;

function Registry_ReadString(AObj: TRegistry; Name: PChar): PChar; extdecl;
begin
  Result :=  PChar(AObj.ReadString(Name));
end;

function Registry_ReadTime(AObj: TRegistry; Name: PChar): TUnixDateTime; extdecl;
begin
  Result :=  ToUnixTime(AObj.ReadTime(Name));
end;

function Registry_RegistryConnect(AObj: TRegistry; UNCName: PChar): LongBool; extdecl;
begin
  Result :=  AObj.RegistryConnect(UNCName);
end;

procedure Registry_RenameValue(AObj: TRegistry; OldName: PChar; NewName: PChar); extdecl;
begin
  AObj.RenameValue(OldName, NewName);
end;

function Registry_ReplaceKey(AObj: TRegistry; Key: PChar; FileName: PChar; BackUpFileName: PChar): LongBool; extdecl;
begin
  Result :=  AObj.ReplaceKey(Key, FileName, BackUpFileName);
end;

function Registry_RestoreKey(AObj: TRegistry; Key: PChar; FileName: PChar): LongBool; extdecl;
begin
  Result :=  AObj.RestoreKey(Key, FileName);
end;

function Registry_SaveKey(AObj: TRegistry; Key: PChar; FileName: PChar): LongBool; extdecl;
begin
  Result :=  AObj.SaveKey(Key, FileName);
end;

function Registry_UnLoadKey(AObj: TRegistry; Key: PChar): LongBool; extdecl;
begin
  Result :=  AObj.UnLoadKey(Key);
end;

function Registry_ValueExists(AObj: TRegistry; Name: PChar): LongBool; extdecl;
begin
  Result :=  AObj.ValueExists(Name);
end;

procedure Registry_WriteBool(AObj: TRegistry; Name: PChar; Value: LongBool); extdecl;
begin
  AObj.WriteBool(Name, Value);
end;

procedure Registry_WriteDate(AObj: TRegistry; Name: PChar; Value: TUnixDateTime); extdecl;
begin
  AObj.WriteDate(Name, UnixToTime(Value));
end;

procedure Registry_WriteDateTime(AObj: TRegistry; Name: PChar; Value: TUnixDateTime); extdecl;
begin
  AObj.WriteDateTime(Name, UnixToTime(Value));
end;

procedure Registry_WriteFloat(AObj: TRegistry; Name: PChar; var Value: Double); extdecl;
begin
  AObj.WriteFloat(Name, Value);
end;

procedure Registry_WriteInteger(AObj: TRegistry; Name: PChar; Value: Integer); extdecl;
begin
  AObj.WriteInteger(Name, Value);
end;

procedure Registry_WriteString(AObj: TRegistry; Name: PChar; Value: PChar); extdecl;
begin
  AObj.WriteString(Name, Value);
end;

procedure Registry_WriteExpandString(AObj: TRegistry; Name: PChar; Value: PChar); extdecl;
begin
  AObj.WriteExpandString(Name, Value);
end;

procedure Registry_WriteTime(AObj: TRegistry; Name: PChar; Value: TUnixDateTime); extdecl;
begin
  AObj.WriteTime(Name, UnixToTime(Value));
end;

function Registry_ClassName(AObj: TRegistry): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function Registry_Equals(AObj: TRegistry; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function Registry_GetHashCode(AObj: TRegistry): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function Registry_ToString(AObj: TRegistry): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function Registry_GetCurrentKey(AObj: TRegistry): HKEY; extdecl;
begin
  Result :=  AObj.CurrentKey;
end;

function Registry_GetCurrentPath(AObj: TRegistry): PChar; extdecl;
begin
  Result :=  PChar(AObj.CurrentPath);
end;

function Registry_GetLazyWrite(AObj: TRegistry): LongBool; extdecl;
begin
  Result :=  AObj.LazyWrite;
end;

procedure Registry_SetLazyWrite(AObj: TRegistry; AValue: LongBool); extdecl;
begin
  AObj.LazyWrite := AValue;
end;

function Registry_GetLastError(AObj: TRegistry): Integer; extdecl;
begin
  Result :=  AObj.LastError;
end;

function Registry_GetLastErrorMsg(AObj: TRegistry): PChar; extdecl;
begin
  Result :=  PChar(AObj.LastErrorMsg);
end;

function Registry_GetRootKey(AObj: TRegistry): HKEY; extdecl;
begin
  Result :=  AObj.RootKey;
end;

procedure Registry_SetRootKey(AObj: TRegistry; AValue: HKEY); extdecl;
begin
  AObj.RootKey := AValue;
end;

function Registry_GetAccess(AObj: TRegistry): Cardinal; extdecl;
begin
  Result :=  AObj.Access;
end;

procedure Registry_SetAccess(AObj: TRegistry; AValue: Cardinal); extdecl;
begin
  AObj.Access := AValue;
end;


exports
  Registry_Create,
  Registry_Free,
  Registry_CloseKey,
  Registry_CreateKey,
  Registry_DeleteKey,
  Registry_DeleteValue,
  Registry_HasSubKeys,
  Registry_KeyExists,
  Registry_LoadKey,
  Registry_MoveKey,
  Registry_OpenKey,
  Registry_OpenKeyReadOnly,
  Registry_ReadBool,
  Registry_ReadDate,
  Registry_ReadDateTime,
  Registry_ReadFloat,
  Registry_ReadInteger,
  Registry_ReadString,
  Registry_ReadTime,
  Registry_RegistryConnect,
  Registry_RenameValue,
  Registry_ReplaceKey,
  Registry_RestoreKey,
  Registry_SaveKey,
  Registry_UnLoadKey,
  Registry_ValueExists,
  Registry_WriteBool,
  Registry_WriteDate,
  Registry_WriteDateTime,
  Registry_WriteFloat,
  Registry_WriteInteger,
  Registry_WriteString,
  Registry_WriteExpandString,
  Registry_WriteTime,
  Registry_ClassName,
  Registry_Equals,
  Registry_GetHashCode,
  Registry_ToString,
  Registry_GetCurrentKey,
  Registry_GetCurrentPath,
  Registry_GetLazyWrite,
  Registry_SetLazyWrite,
  Registry_GetLastError,
  Registry_GetLastErrorMsg,
  Registry_GetRootKey,
  Registry_SetRootKey,
  Registry_GetAccess,
  Registry_SetAccess;

