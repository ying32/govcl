
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function MemoryStream_Create: TMemoryStream; stdcall;
begin
  Result :=  TMemoryStream.Create;
end;

procedure MemoryStream_Free(AObj: TMemoryStream); stdcall;
begin
  AObj.Free;
end;

procedure MemoryStream_Clear(AObj: TMemoryStream); stdcall;
begin
  AObj.Clear;
end;

procedure MemoryStream_LoadFromStream(AObj: TMemoryStream; Stream: TStream); stdcall;
begin
  AObj.LoadFromStream(Stream);
end;

procedure MemoryStream_LoadFromFile(AObj: TMemoryStream; FileName: PChar); stdcall;
begin
  AObj.LoadFromFile(FileName);
end;

procedure MemoryStream_Seek(AObj: TMemoryStream; var Offset: Int64; Origin: TSeekOrigin; var Result: Int64); stdcall;
begin
  Result :=  AObj.Seek(Offset, Origin);
end;

procedure MemoryStream_SaveToStream(AObj: TMemoryStream; Stream: TStream); stdcall;
begin
  AObj.SaveToStream(Stream);
end;

procedure MemoryStream_SaveToFile(AObj: TMemoryStream; FileName: PChar); stdcall;
begin
  AObj.SaveToFile(FileName);
end;

procedure MemoryStream_CopyFrom(AObj: TMemoryStream; Source: TStream; var Count: Int64; var Result: Int64); stdcall;
begin
  Result :=  AObj.CopyFrom(Source, Count);
end;

function MemoryStream_ClassName(AObj: TMemoryStream): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function MemoryStream_Equals(AObj: TMemoryStream; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function MemoryStream_GetHashCode(AObj: TMemoryStream): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function MemoryStream_ToString(AObj: TMemoryStream): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function MemoryStream_GetMemory(AObj: TMemoryStream): Pointer; stdcall;
begin
  Result :=  AObj.Memory;
end;

procedure MemoryStream_GetPosition(AObj: TMemoryStream; var Result: Int64); stdcall;
begin
  Result :=  AObj.Position;
end;

procedure MemoryStream_SetPosition(AObj: TMemoryStream; var AValue: Int64); stdcall;
begin
  AObj.Position := AValue;
end;

procedure MemoryStream_GetSize(AObj: TMemoryStream; var Result: Int64); stdcall;
begin
  Result :=  AObj.Size;
end;

procedure MemoryStream_SetSize(AObj: TMemoryStream; var AValue: Int64); stdcall;
begin
  AObj.Size := AValue;
end;


exports
  MemoryStream_Create,
  MemoryStream_Free,
  MemoryStream_Clear,
  MemoryStream_LoadFromStream,
  MemoryStream_LoadFromFile,
  MemoryStream_Seek,
  MemoryStream_SaveToStream,
  MemoryStream_SaveToFile,
  MemoryStream_CopyFrom,
  MemoryStream_ClassName,
  MemoryStream_Equals,
  MemoryStream_GetHashCode,
  MemoryStream_ToString,
  MemoryStream_GetMemory,
  MemoryStream_GetPosition,
  MemoryStream_SetPosition,
  MemoryStream_GetSize,
  MemoryStream_SetSize;

