
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function Clipboard_Create: TClipboard; extdecl;
begin
  Result :=  TClipboard.Create;
end;

procedure Clipboard_Free(AObj: TClipboard); extdecl;
begin
  AObj.Free;
end;

procedure Clipboard_Assign(AObj: TClipboard; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

procedure Clipboard_Clear(AObj: TClipboard); extdecl;
begin
  AObj.Clear;
end;

procedure Clipboard_Close(AObj: TClipboard); extdecl;
begin
  AObj.Close;
end;

function Clipboard_HasFormat(AObj: TClipboard; Format: Word): LongBool; extdecl;
begin
  Result :=  AObj.HasFormat(Format);
end;

procedure Clipboard_Open(AObj: TClipboard); extdecl;
begin
  AObj.Open;
end;

function Clipboard_GetTextBuf(AObj: TClipboard; Buffer: PChar; BufSize: Integer): Integer; extdecl;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function Clipboard_GetNamePath(AObj: TClipboard): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function Clipboard_ClassName(AObj: TClipboard): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function Clipboard_Equals(AObj: TClipboard; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function Clipboard_GetHashCode(AObj: TClipboard): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function Clipboard_ToString(AObj: TClipboard): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function Clipboard_GetAsText(AObj: TClipboard): PChar; extdecl;
begin
  Result :=  PChar(AObj.AsText);
end;

procedure Clipboard_SetAsText(AObj: TClipboard; AValue: PChar); extdecl;
begin
  AObj.AsText := AValue;
end;

function Clipboard_GetFormatCount(AObj: TClipboard): Integer; extdecl;
begin
  Result :=  AObj.FormatCount;
end;

function Clipboard_GetFormats(AObj: TClipboard; Index: Integer): Word; extdecl;
begin
  Result :=  AObj.Formats[Index];
end;


exports
  Clipboard_Create,
  Clipboard_Free,
  Clipboard_Assign,
  Clipboard_Clear,
  Clipboard_Close,
  Clipboard_HasFormat,
  Clipboard_Open,
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

