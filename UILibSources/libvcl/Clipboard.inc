
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function Clipboard_Create: TClipboard; stdcall;
begin
  Result :=  TClipboard.Create;
end;

procedure Clipboard_Free(AObj: TClipboard); stdcall;
begin
  AObj.Free;
end;

procedure Clipboard_Assign(AObj: TClipboard; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

procedure Clipboard_Clear(AObj: TClipboard); stdcall;
begin
  AObj.Clear;
end;

procedure Clipboard_Close(AObj: TClipboard); stdcall;
begin
  AObj.Close;
end;

function Clipboard_GetAsHandle(AObj: TClipboard; Format: Word): NativeUInt; stdcall;
begin
  Result :=  AObj.GetAsHandle(Format);
end;

function Clipboard_HasFormat(AObj: TClipboard; Format: Word): LongBool; stdcall;
begin
  Result :=  AObj.HasFormat(Format);
end;

procedure Clipboard_Open(AObj: TClipboard); stdcall;
begin
  AObj.Open;
end;

procedure Clipboard_SetAsHandle(AObj: TClipboard; Format: Word; Value: NativeUInt); stdcall;
begin
  AObj.SetAsHandle(Format, Value);
end;

function Clipboard_GetTextBuf(AObj: TClipboard; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function Clipboard_GetNamePath(AObj: TClipboard): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function Clipboard_ClassName(AObj: TClipboard): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function Clipboard_Equals(AObj: TClipboard; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function Clipboard_GetHashCode(AObj: TClipboard): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function Clipboard_ToString(AObj: TClipboard): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function Clipboard_GetAsText(AObj: TClipboard): PChar; stdcall;
begin
  Result :=  PChar(AObj.AsText);
end;

procedure Clipboard_SetAsText(AObj: TClipboard; AValue: PChar); stdcall;
begin
  AObj.AsText := AValue;
end;

function Clipboard_GetFormatCount(AObj: TClipboard): Integer; stdcall;
begin
  Result :=  AObj.FormatCount;
end;

function Clipboard_GetFormats(AObj: TClipboard; Index: Integer): Word; stdcall;
begin
  Result :=  AObj.Formats[Index];
end;


exports
  Clipboard_Create,
  Clipboard_Free,
  Clipboard_Assign,
  Clipboard_Clear,
  Clipboard_Close,
  Clipboard_GetAsHandle,
  Clipboard_HasFormat,
  Clipboard_Open,
  Clipboard_SetAsHandle,
  Clipboard_GetTextBuf,
  Clipboard_GetNamePath,
  Clipboard_ClassName,
  Clipboard_Equals,
  Clipboard_GetHashCode,
  Clipboard_ToString,
  Clipboard_GetAsText,
  Clipboard_SetAsText,
  Clipboard_GetFormatCount,
  Clipboard_GetFormats;

