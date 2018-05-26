
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function GIFImage_Create: TGIFImage; stdcall;
begin
  Result :=  TGIFImage.Create;
end;

procedure GIFImage_Free(AObj: TGIFImage); stdcall;
begin
  AObj.Free;
end;

procedure GIFImage_SaveToStream(AObj: TGIFImage; Stream: TStream); stdcall;
begin
  AObj.SaveToStream(Stream);
end;

procedure GIFImage_LoadFromStream(AObj: TGIFImage; Stream: TStream); stdcall;
begin
  AObj.LoadFromStream(Stream);
end;

function GIFImage_Add(AObj: TGIFImage; Source: TPersistent): TGIFFrame; stdcall;
begin
  Result :=  AObj.Add(Source);
end;

procedure GIFImage_Clear(AObj: TGIFImage); stdcall;
begin
  AObj.Clear;
end;

procedure GIFImage_Assign(AObj: TGIFImage; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

procedure GIFImage_StopDraw(AObj: TGIFImage); stdcall;
begin
  AObj.StopDraw;
end;

procedure GIFImage_SuspendDraw(AObj: TGIFImage); stdcall;
begin
  AObj.SuspendDraw;
end;

procedure GIFImage_ResumeDraw(AObj: TGIFImage); stdcall;
begin
  AObj.ResumeDraw;
end;

function GIFImage_Equals(AObj: TGIFImage; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

procedure GIFImage_LoadFromFile(AObj: TGIFImage; Filename: PChar); stdcall;
begin
  AObj.LoadFromFile(Filename);
end;

procedure GIFImage_SaveToFile(AObj: TGIFImage; Filename: PChar); stdcall;
begin
  AObj.SaveToFile(Filename);
end;

procedure GIFImage_SetSize(AObj: TGIFImage; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetSize(AWidth, AHeight);
end;

function GIFImage_GetNamePath(AObj: TGIFImage): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function GIFImage_ClassName(AObj: TGIFImage): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function GIFImage_GetHashCode(AObj: TGIFImage): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function GIFImage_ToString(AObj: TGIFImage): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function GIFImage_GetVersion(AObj: TGIFImage): TGIFVersion; stdcall;
begin
  Result :=  AObj.Version;
end;

function GIFImage_GetBitsPerPixel(AObj: TGIFImage): Integer; stdcall;
begin
  Result :=  AObj.BitsPerPixel;
end;

function GIFImage_GetBackgroundColor(AObj: TGIFImage): TColor; stdcall;
begin
  Result :=  AObj.BackgroundColor;
end;

procedure GIFImage_SetBackgroundColor(AObj: TGIFImage; AValue: TColor); stdcall;
begin
  AObj.BackgroundColor := AValue;
end;

function GIFImage_GetAspectRatio(AObj: TGIFImage): Byte; stdcall;
begin
  Result :=  AObj.AspectRatio;
end;

procedure GIFImage_SetAspectRatio(AObj: TGIFImage; AValue: Byte); stdcall;
begin
  AObj.AspectRatio := AValue;
end;

function GIFImage_GetIsTransparent(AObj: TGIFImage): LongBool; stdcall;
begin
  Result :=  AObj.IsTransparent;
end;

function GIFImage_GetAnimationSpeed(AObj: TGIFImage): Integer; stdcall;
begin
  Result :=  AObj.AnimationSpeed;
end;

procedure GIFImage_SetAnimationSpeed(AObj: TGIFImage; AValue: Integer); stdcall;
begin
  AObj.AnimationSpeed := AValue;
end;

function GIFImage_GetBitmap(AObj: TGIFImage): TBitmap; stdcall;
begin
  Result :=  AObj.Bitmap;
end;

procedure GIFImage_SetOnPaint(AObj: TGIFImage; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnPaint := nil;
    TEventClass.Remove(AObj, gePaint);
    Exit;
  end;
  AObj.OnPaint := TEventClass.OnPaint;
  TEventClass.Add(AObj, gePaint, AEventId);
end;

function GIFImage_GetAnimate(AObj: TGIFImage): LongBool; stdcall;
begin
  Result :=  AObj.Animate;
end;

procedure GIFImage_SetAnimate(AObj: TGIFImage; AValue: LongBool); stdcall;
begin
  AObj.Animate := AValue;
end;

function GIFImage_GetAnimateLoop(AObj: TGIFImage): TGIFAnimationLoop; stdcall;
begin
  Result :=  AObj.AnimateLoop;
end;

procedure GIFImage_SetAnimateLoop(AObj: TGIFImage; AValue: TGIFAnimationLoop); stdcall;
begin
  AObj.AnimateLoop := AValue;
end;

function GIFImage_GetShouldDither(AObj: TGIFImage): LongBool; stdcall;
begin
  Result :=  AObj.ShouldDither;
end;

function GIFImage_GetEmpty(AObj: TGIFImage): LongBool; stdcall;
begin
  Result :=  AObj.Empty;
end;

function GIFImage_GetHeight(AObj: TGIFImage): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure GIFImage_SetHeight(AObj: TGIFImage; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function GIFImage_GetModified(AObj: TGIFImage): LongBool; stdcall;
begin
  Result :=  AObj.Modified;
end;

procedure GIFImage_SetModified(AObj: TGIFImage; AValue: LongBool); stdcall;
begin
  AObj.Modified := AValue;
end;

function GIFImage_GetPaletteModified(AObj: TGIFImage): LongBool; stdcall;
begin
  Result :=  AObj.PaletteModified;
end;

procedure GIFImage_SetPaletteModified(AObj: TGIFImage; AValue: LongBool); stdcall;
begin
  AObj.PaletteModified := AValue;
end;

function GIFImage_GetTransparent(AObj: TGIFImage): LongBool; stdcall;
begin
  Result :=  AObj.Transparent;
end;

procedure GIFImage_SetTransparent(AObj: TGIFImage; AValue: LongBool); stdcall;
begin
  AObj.Transparent := AValue;
end;

function GIFImage_GetWidth(AObj: TGIFImage): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure GIFImage_SetWidth(AObj: TGIFImage; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

procedure GIFImage_SetOnChange(AObj: TGIFImage; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnChange := nil;
    TEventClass.Remove(AObj, geChange);
    Exit;
  end;
  AObj.OnChange := TEventClass.OnChange;
  TEventClass.Add(AObj, geChange, AEventId);
end;


exports
  GIFImage_Create,
  GIFImage_Free,
  GIFImage_SaveToStream,
  GIFImage_LoadFromStream,
  GIFImage_Add,
  GIFImage_Clear,
  GIFImage_Assign,
  GIFImage_StopDraw,
  GIFImage_SuspendDraw,
  GIFImage_ResumeDraw,
  GIFImage_Equals,
  GIFImage_LoadFromFile,
  GIFImage_SaveToFile,
  GIFImage_SetSize,
  GIFImage_GetNamePath,
  GIFImage_ClassName,
  GIFImage_GetHashCode,
  GIFImage_ToString,
  GIFImage_GetVersion,
  GIFImage_GetBitsPerPixel,
  GIFImage_GetBackgroundColor,
  GIFImage_SetBackgroundColor,
  GIFImage_GetAspectRatio,
  GIFImage_SetAspectRatio,
  GIFImage_GetIsTransparent,
  GIFImage_GetAnimationSpeed,
  GIFImage_SetAnimationSpeed,
  GIFImage_GetBitmap,
  GIFImage_SetOnPaint,
  GIFImage_GetAnimate,
  GIFImage_SetAnimate,
  GIFImage_GetAnimateLoop,
  GIFImage_SetAnimateLoop,
  GIFImage_GetShouldDither,
  GIFImage_GetEmpty,
  GIFImage_GetHeight,
  GIFImage_SetHeight,
  GIFImage_GetModified,
  GIFImage_SetModified,
  GIFImage_GetPaletteModified,
  GIFImage_SetPaletteModified,
  GIFImage_GetTransparent,
  GIFImage_SetTransparent,
  GIFImage_GetWidth,
  GIFImage_SetWidth,
  GIFImage_SetOnChange;

