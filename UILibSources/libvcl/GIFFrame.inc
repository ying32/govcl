
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function GIFFrame_Create(AOwner: TGIFImage): TGIFFrame; stdcall;
begin
  Result :=  TGIFFrame.Create(AOwner);
end;

procedure GIFFrame_Free(AObj: TGIFFrame); stdcall;
begin
  AObj.Free;
end;

procedure GIFFrame_Clear(AObj: TGIFFrame); stdcall;
begin
  AObj.Clear;
end;

procedure GIFFrame_SaveToStream(AObj: TGIFFrame; Stream: TStream); stdcall;
begin
  AObj.SaveToStream(Stream);
end;

procedure GIFFrame_LoadFromStream(AObj: TGIFFrame; Stream: TStream); stdcall;
begin
  AObj.LoadFromStream(Stream);
end;

procedure GIFFrame_Assign(AObj: TGIFFrame; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

procedure GIFFrame_SaveToFile(AObj: TGIFFrame; Filename: PChar); stdcall;
begin
  AObj.SaveToFile(Filename);
end;

procedure GIFFrame_LoadFromFile(AObj: TGIFFrame; Filename: PChar); stdcall;
begin
  AObj.LoadFromFile(Filename);
end;

function GIFFrame_GetNamePath(AObj: TGIFFrame): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function GIFFrame_ClassName(AObj: TGIFFrame): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function GIFFrame_Equals(AObj: TGIFFrame; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function GIFFrame_GetHashCode(AObj: TGIFFrame): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function GIFFrame_ToString(AObj: TGIFFrame): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function GIFFrame_GetHasBitmap(AObj: TGIFFrame): LongBool; stdcall;
begin
  Result :=  AObj.HasBitmap;
end;

procedure GIFFrame_SetHasBitmap(AObj: TGIFFrame; AValue: LongBool); stdcall;
begin
  AObj.HasBitmap := AValue;
end;

function GIFFrame_GetLeft(AObj: TGIFFrame): Word; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure GIFFrame_SetLeft(AObj: TGIFFrame; AValue: Word); stdcall;
begin
  AObj.Left := AValue;
end;

function GIFFrame_GetTop(AObj: TGIFFrame): Word; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure GIFFrame_SetTop(AObj: TGIFFrame; AValue: Word); stdcall;
begin
  AObj.Top := AValue;
end;

function GIFFrame_GetWidth(AObj: TGIFFrame): Word; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure GIFFrame_SetWidth(AObj: TGIFFrame; AValue: Word); stdcall;
begin
  AObj.Width := AValue;
end;

function GIFFrame_GetHeight(AObj: TGIFFrame): Word; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure GIFFrame_SetHeight(AObj: TGIFFrame; AValue: Word); stdcall;
begin
  AObj.Height := AValue;
end;

procedure GIFFrame_GetBoundsRect(AObj: TGIFFrame; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure GIFFrame_SetBoundsRect(AObj: TGIFFrame; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

procedure GIFFrame_GetClientRect(AObj: TGIFFrame; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function GIFFrame_GetData(AObj: TGIFFrame): PByte; stdcall;
begin
  Result :=  AObj.Data;
end;

function GIFFrame_GetDataSize(AObj: TGIFFrame): Integer; stdcall;
begin
  Result :=  AObj.DataSize;
end;

function GIFFrame_GetVersion(AObj: TGIFFrame): TGIFVersion; stdcall;
begin
  Result :=  AObj.Version;
end;

function GIFFrame_GetBitsPerPixel(AObj: TGIFFrame): Integer; stdcall;
begin
  Result :=  AObj.BitsPerPixel;
end;

function GIFFrame_GetBitmap(AObj: TGIFFrame): TBitmap; stdcall;
begin
  Result :=  AObj.Bitmap;
end;

procedure GIFFrame_SetBitmap(AObj: TGIFFrame; AValue: TBitmap); stdcall;
begin
  AObj.Bitmap := AValue;
end;

function GIFFrame_GetEmpty(AObj: TGIFFrame): LongBool; stdcall;
begin
  Result :=  AObj.Empty;
end;

function GIFFrame_GetTransparent(AObj: TGIFFrame): LongBool; stdcall;
begin
  Result :=  AObj.Transparent;
end;


exports
  GIFFrame_Create,
  GIFFrame_Free,
  GIFFrame_Clear,
  GIFFrame_SaveToStream,
  GIFFrame_LoadFromStream,
  GIFFrame_Assign,
  GIFFrame_SaveToFile,
  GIFFrame_LoadFromFile,
  GIFFrame_GetNamePath,
  GIFFrame_ClassName,
  GIFFrame_Equals,
  GIFFrame_GetHashCode,
  GIFFrame_ToString,
  GIFFrame_GetHasBitmap,
  GIFFrame_SetHasBitmap,
  GIFFrame_GetLeft,
  GIFFrame_SetLeft,
  GIFFrame_GetTop,
  GIFFrame_SetTop,
  GIFFrame_GetWidth,
  GIFFrame_SetWidth,
  GIFFrame_GetHeight,
  GIFFrame_SetHeight,
  GIFFrame_GetBoundsRect,
  GIFFrame_SetBoundsRect,
  GIFFrame_GetClientRect,
  GIFFrame_GetData,
  GIFFrame_GetDataSize,
  GIFFrame_GetVersion,
  GIFFrame_GetBitsPerPixel,
  GIFFrame_GetBitmap,
  GIFFrame_SetBitmap,
  GIFFrame_GetEmpty,
  GIFFrame_GetTransparent;

