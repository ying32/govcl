
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function IniFile_Create(AFileName: PChar): TIniFile; stdcall;
begin
  Result :=  TIniFile.Create(AFileName);
end;

procedure IniFile_Free(AObj: TIniFile); stdcall;
begin
  AObj.Free;
end;

function IniFile_ReadString(AObj: TIniFile; Section: PChar; Ident: PChar; Default: PChar): PChar; stdcall;
begin
  Result :=  PChar(AObj.ReadString(Section, Ident, Default));
end;

procedure IniFile_WriteString(AObj: TIniFile; Section: PChar; Ident: PChar; Value: PChar); stdcall;
begin
  AObj.WriteString(Section, Ident, Value);
end;

procedure IniFile_ReadSections(AObj: TIniFile; Strings: TStrings); stdcall;
begin
  AObj.ReadSections(Strings);
end;

procedure IniFile_ReadSectionValues(AObj: TIniFile; Section: PChar; Strings: TStrings); stdcall;
begin
  AObj.ReadSectionValues(Section, Strings);
end;

procedure IniFile_EraseSection(AObj: TIniFile; Section: PChar); stdcall;
begin
  AObj.EraseSection(Section);
end;

procedure IniFile_DeleteKey(AObj: TIniFile; Section: PChar; Ident: PChar); stdcall;
begin
  AObj.DeleteKey(Section, Ident);
end;

procedure IniFile_UpdateFile(AObj: TIniFile); stdcall;
begin
  AObj.UpdateFile;
end;

function IniFile_SectionExists(AObj: TIniFile; Section: PChar): LongBool; stdcall;
begin
  Result :=  AObj.SectionExists(Section);
end;

function IniFile_ReadInteger(AObj: TIniFile; Section: PChar; Ident: PChar; Default: Integer): Integer; stdcall;
begin
  Result :=  AObj.ReadInteger(Section, Ident, Default);
end;

procedure IniFile_WriteInteger(AObj: TIniFile; Section: PChar; Ident: PChar; Value: Integer); stdcall;
begin
  AObj.WriteInteger(Section, Ident, Value);
end;

function IniFile_ReadBool(AObj: TIniFile; Section: PChar; Ident: PChar; Default: LongBool): LongBool; stdcall;
begin
  Result :=  AObj.ReadBool(Section, Ident, Default);
end;

procedure IniFile_WriteBool(AObj: TIniFile; Section: PChar; Ident: PChar; Value: LongBool); stdcall;
begin
  AObj.WriteBool(Section, Ident, Value);
end;

function IniFile_ReadDate(AObj: TIniFile; Section: PChar; Name: PChar; Default: TUnixDateTime): TUnixDateTime; stdcall;
begin
  Result :=  ToUnixTime(AObj.ReadDate(Section, Name, UnixToTime(Default)));
end;

function IniFile_ReadDateTime(AObj: TIniFile; Section: PChar; Name: PChar; Default: TUnixDateTime): TUnixDateTime; stdcall;
begin
  Result :=  ToUnixTime(AObj.ReadDateTime(Section, Name, UnixToTime(Default)));
end;

procedure IniFile_ReadFloat(AObj: TIniFile; Section: PChar; Name: PChar; var Default: Double; var Result: Double); stdcall;
begin
  Result :=  AObj.ReadFloat(Section, Name, Default);
end;

function IniFile_ReadTime(AObj: TIniFile; Section: PChar; Name: PChar; Default: TUnixDateTime): TUnixDateTime; stdcall;
begin
  Result :=  ToUnixTime(AObj.ReadTime(Section, Name, UnixToTime(Default)));
end;

procedure IniFile_WriteDate(AObj: TIniFile; Section: PChar; Name: PChar; Value: TUnixDateTime); stdcall;
begin
  AObj.WriteDate(Section, Name, UnixToTime(Value));
end;

procedure IniFile_WriteDateTime(AObj: TIniFile; Section: PChar; Name: PChar; Value: TUnixDateTime); stdcall;
begin
  AObj.WriteDateTime(Section, Name, UnixToTime(Value));
end;

procedure IniFile_WriteFloat(AObj: TIniFile; Section: PChar; Name: PChar; var Value: Double); stdcall;
begin
  AObj.WriteFloat(Section, Name, Value);
end;

procedure IniFile_WriteTime(AObj: TIniFile; Section: PChar; Name: PChar; Value: TUnixDateTime); stdcall;
begin
  AObj.WriteTime(Section, Name, UnixToTime(Value));
end;

procedure IniFile_ReadSubSections(AObj: TIniFile; Section: PChar; Strings: TStrings; Recurse: LongBool); stdcall;
begin
  AObj.ReadSubSections(Section, Strings, Recurse);
end;

function IniFile_ValueExists(AObj: TIniFile; Section: PChar; Ident: PChar): LongBool; stdcall;
begin
  Result :=  AObj.ValueExists(Section, Ident);
end;

function IniFile_ClassName(AObj: TIniFile): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function IniFile_Equals(AObj: TIniFile; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function IniFile_GetHashCode(AObj: TIniFile): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function IniFile_ToString(AObj: TIniFile): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function IniFile_GetFileName(AObj: TIniFile): PChar; stdcall;
begin
  Result :=  PChar(AObj.FileName);
end;


exports
  IniFile_Create,
  IniFile_Free,
  IniFile_ReadString,
  IniFile_WriteString,
  IniFile_ReadSections,
  IniFile_ReadSectionValues,
  IniFile_EraseSection,
  IniFile_DeleteKey,
  IniFile_UpdateFile,
  IniFile_SectionExists,
  IniFile_ReadInteger,
  IniFile_WriteInteger,
  IniFile_ReadBool,
  IniFile_WriteBool,
  IniFile_ReadDate,
  IniFile_ReadDateTime,
  IniFile_ReadFloat,
  IniFile_ReadTime,
  IniFile_WriteDate,
  IniFile_WriteDateTime,
  IniFile_WriteFloat,
  IniFile_WriteTime,
  IniFile_ReadSubSections,
  IniFile_ValueExists,
  IniFile_ClassName,
  IniFile_Equals,
  IniFile_GetHashCode,
  IniFile_ToString,
  IniFile_GetFileName;

