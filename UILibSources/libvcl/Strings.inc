
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function Strings_Create: TStrings; stdcall;
begin
  Result :=  TStrings.Create;
end;

procedure Strings_Free(AObj: TStrings); stdcall;
begin
  AObj.Free;
end;

function Strings_Add(AObj: TStrings; S: PChar): Integer; stdcall;
begin
  Result :=  AObj.Add(S);
end;

function Strings_AddObject(AObj: TStrings; S: PChar; AObject: TObject): Integer; stdcall;
begin
  Result :=  AObj.AddObject(S, AObject);
end;

procedure Strings_Append(AObj: TStrings; S: PChar); stdcall;
begin
  AObj.Append(S);
end;

procedure Strings_Assign(AObj: TStrings; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

procedure Strings_BeginUpdate(AObj: TStrings); stdcall;
begin
  AObj.BeginUpdate;
end;

procedure Strings_Clear(AObj: TStrings); stdcall;
begin
  AObj.Clear;
end;

procedure Strings_Delete(AObj: TStrings; Index: Integer); stdcall;
begin
  AObj.Delete(Index);
end;

procedure Strings_EndUpdate(AObj: TStrings); stdcall;
begin
  AObj.EndUpdate;
end;

function Strings_Equals(AObj: TStrings; Strings: TStrings): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Strings);
end;

function Strings_IndexOf(AObj: TStrings; S: PChar): Integer; stdcall;
begin
  Result :=  AObj.IndexOf(S);
end;

function Strings_IndexOfName(AObj: TStrings; Name: PChar): Integer; stdcall;
begin
  Result :=  AObj.IndexOfName(Name);
end;

function Strings_IndexOfObject(AObj: TStrings; AObject: TObject): Integer; stdcall;
begin
  Result :=  AObj.IndexOfObject(AObject);
end;

procedure Strings_Insert(AObj: TStrings; Index: Integer; S: PChar); stdcall;
begin
  AObj.Insert(Index, S);
end;

procedure Strings_InsertObject(AObj: TStrings; Index: Integer; S: PChar; AObject: TObject); stdcall;
begin
  AObj.InsertObject(Index, S, AObject);
end;

procedure Strings_LoadFromFile(AObj: TStrings; FileName: PChar); stdcall;
begin
  AObj.LoadFromFile(FileName);
end;

procedure Strings_LoadFromStream(AObj: TStrings; Stream: TStream); stdcall;
begin
  AObj.LoadFromStream(Stream);
end;

procedure Strings_Move(AObj: TStrings; CurIndex: Integer; NewIndex: Integer); stdcall;
begin
  AObj.Move(CurIndex, NewIndex);
end;

procedure Strings_SaveToFile(AObj: TStrings; FileName: PChar); stdcall;
begin
  AObj.SaveToFile(FileName);
end;

procedure Strings_SaveToStream(AObj: TStrings; Stream: TStream); stdcall;
begin
  AObj.SaveToStream(Stream);
end;

function Strings_GetNamePath(AObj: TStrings): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function Strings_ClassName(AObj: TStrings): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function Strings_GetHashCode(AObj: TStrings): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function Strings_ToString(AObj: TStrings): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function Strings_GetCapacity(AObj: TStrings): Integer; stdcall;
begin
  Result :=  AObj.Capacity;
end;

procedure Strings_SetCapacity(AObj: TStrings; AValue: Integer); stdcall;
begin
  AObj.Capacity := AValue;
end;

function Strings_GetCommaText(AObj: TStrings): PChar; stdcall;
begin
  Result :=  PChar(AObj.CommaText);
end;

procedure Strings_SetCommaText(AObj: TStrings; AValue: PChar); stdcall;
begin
  AObj.CommaText := AValue;
end;

function Strings_GetCount(AObj: TStrings): Integer; stdcall;
begin
  Result :=  AObj.Count;
end;

function Strings_GetDelimiter(AObj: TStrings): Char; stdcall;
begin
  Result :=  AObj.Delimiter;
end;

procedure Strings_SetDelimiter(AObj: TStrings; AValue: Char); stdcall;
begin
  AObj.Delimiter := AValue;
end;

function Strings_GetText(AObj: TStrings): PChar; stdcall;
begin
  Result :=  PChar(AObj.Text);
end;

procedure Strings_SetText(AObj: TStrings; AValue: PChar); stdcall;
begin
  AObj.Text := AValue;
end;

function Strings_GetWriteBOM(AObj: TStrings): LongBool; stdcall;
begin
  Result :=  AObj.WriteBOM;
end;

procedure Strings_SetWriteBOM(AObj: TStrings; AValue: LongBool); stdcall;
begin
  AObj.WriteBOM := AValue;
end;

function Strings_GetOptions(AObj: TStrings): TStringsOptions; stdcall;
begin
  Result :=  AObj.Options;
end;

procedure Strings_SetOptions(AObj: TStrings; AValue: TStringsOptions); stdcall;
begin
  AObj.Options := AValue;
end;

function Strings_GetObjects(AObj: TStrings; Index: Integer): TObject; stdcall;
begin
  Result :=  AObj.Objects[Index];
end;

procedure Strings_SetObjects(AObj: TStrings; Index: Integer; AValue: TObject); stdcall;
begin
  AObj.Objects[Index] := AValue;
end;

function Strings_GetValues(AObj: TStrings; Name: PChar): PChar; stdcall;
begin
  Result :=  PChar(AObj.Values[Name]);
end;

procedure Strings_SetValues(AObj: TStrings; Name: string; AValue: PChar); stdcall;
begin
  AObj.Values[Name] := AValue;
end;

function Strings_GetValueFromIndex(AObj: TStrings; Index: Integer): PChar; stdcall;
begin
  Result :=  PChar(AObj.ValueFromIndex[Index]);
end;

procedure Strings_SetValueFromIndex(AObj: TStrings; Index: Integer; AValue: PChar); stdcall;
begin
  AObj.ValueFromIndex[Index] := AValue;
end;

function Strings_GetStrings(AObj: TStrings; Index: Integer): PChar; stdcall;
begin
  Result :=  PChar(AObj.Strings[Index]);
end;

procedure Strings_SetStrings(AObj: TStrings; Index: Integer; AValue: PChar); stdcall;
begin
  AObj.Strings[Index] := AValue;
end;


exports
  Strings_Create,
  Strings_Free,
  Strings_Add,
  Strings_AddObject,
  Strings_Append,
  Strings_Assign,
  Strings_BeginUpdate,
  Strings_Clear,
  Strings_Delete,
  Strings_EndUpdate,
  Strings_Equals,
  Strings_IndexOf,
  Strings_IndexOfName,
  Strings_IndexOfObject,
  Strings_Insert,
  Strings_InsertObject,
  Strings_LoadFromFile,
  Strings_LoadFromStream,
  Strings_Move,
  Strings_SaveToFile,
  Strings_SaveToStream,
  Strings_GetNamePath,
  Strings_ClassName,
  Strings_GetHashCode,
  Strings_ToString,
  Strings_GetCapacity,
  Strings_SetCapacity,
  Strings_GetCommaText,
  Strings_SetCommaText,
  Strings_GetCount,
  Strings_GetDelimiter,
  Strings_SetDelimiter,
  Strings_GetText,
  Strings_SetText,
  Strings_GetWriteBOM,
  Strings_SetWriteBOM,
  Strings_GetOptions,
  Strings_SetOptions,
  Strings_GetObjects,
  Strings_SetObjects,
  Strings_GetValues,
  Strings_SetValues,
  Strings_GetValueFromIndex,
  Strings_SetValueFromIndex,
  Strings_GetStrings,
  Strings_SetStrings;

