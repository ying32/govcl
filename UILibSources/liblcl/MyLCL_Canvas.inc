
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function Canvas_Create: TCanvas; extdecl;
begin
  Result :=  TCanvas.Create;
end;

procedure Canvas_Free(AObj: TCanvas); extdecl;
begin
  AObj.Free;
end;

procedure Canvas_Arc(AObj: TCanvas; X1: Integer; Y1: Integer; X2: Integer; Y2: Integer; X3: Integer; Y3: Integer; X4: Integer; Y4: Integer); extdecl;
begin
  AObj.Arc(X1, Y1, X2, Y2, X3, Y3, X4, Y4);
end;

procedure Canvas_ArcTo(AObj: TCanvas; X1: Integer; Y1: Integer; X2: Integer; Y2: Integer; X3: Integer; Y3: Integer; X4: Integer; Y4: Integer); extdecl;
begin
  AObj.ArcTo(X1, Y1, X2, Y2, X3, Y3, X4, Y4);
end;

procedure Canvas_AngleArc(AObj: TCanvas; X: Integer; Y: Integer; Radius: Cardinal; var StartAngle: Single; var SweepAngle: Single); extdecl;
begin
  AObj.AngleArc(X, Y, Radius, StartAngle, SweepAngle);
end;

procedure Canvas_Chord(AObj: TCanvas; X1: Integer; Y1: Integer; X2: Integer; Y2: Integer; X3: Integer; Y3: Integer; X4: Integer; Y4: Integer); extdecl;
begin
  AObj.Chord(X1, Y1, X2, Y2, X3, Y3, X4, Y4);
end;

procedure Canvas_Ellipse(AObj: TCanvas; X1: Integer; Y1: Integer; X2: Integer; Y2: Integer); extdecl;
begin
  AObj.Ellipse(X1, Y1, X2, Y2);
end;

procedure Canvas_FloodFill(AObj: TCanvas; X: Integer; Y: Integer; Color: TColor; FillStyle: TFillStyle); extdecl;
begin
  AObj.FloodFill(X, Y, Color, FillStyle);
end;

procedure Canvas_LineTo(AObj: TCanvas; X: Integer; Y: Integer); extdecl;
begin
  AObj.LineTo(X, Y);
end;

procedure Canvas_MoveTo(AObj: TCanvas; X: Integer; Y: Integer); extdecl;
begin
  AObj.MoveTo(X, Y);
end;

procedure Canvas_Pie(AObj: TCanvas; X1: Integer; Y1: Integer; X2: Integer; Y2: Integer; X3: Integer; Y3: Integer; X4: Integer; Y4: Integer); extdecl;
begin
  AObj.Pie(X1, Y1, X2, Y2, X3, Y3, X4, Y4);
end;

procedure Canvas_Rectangle(AObj: TCanvas; X1: Integer; Y1: Integer; X2: Integer; Y2: Integer); extdecl;
begin
  AObj.Rectangle(X1, Y1, X2, Y2);
end;

procedure Canvas_Refresh(AObj: TCanvas); extdecl;
begin
  AObj.Refresh;
end;

procedure Canvas_RoundRect(AObj: TCanvas; X1: Integer; Y1: Integer; X2: Integer; Y2: Integer; X3: Integer; Y3: Integer); extdecl;
begin
  AObj.RoundRect(X1, Y1, X2, Y2, X3, Y3);
end;

procedure Canvas_TextExtent(AObj: TCanvas; Text: PChar; var Result: TSize); extdecl;
begin
  Result :=  AObj.TextExtent(Text);
end;

procedure Canvas_TextOut(AObj: TCanvas; X: Integer; Y: Integer; Text: PChar); extdecl;
begin
  AObj.TextOut(X, Y, Text);
end;

procedure Canvas_Lock(AObj: TCanvas); extdecl;
begin
  AObj.Lock;
end;

function Canvas_TextHeight(AObj: TCanvas; Text: PChar): Integer; extdecl;
begin
  Result :=  AObj.TextHeight(Text);
end;

function Canvas_TextWidth(AObj: TCanvas; Text: PChar): Integer; extdecl;
begin
  Result :=  AObj.TextWidth(Text);
end;

procedure Canvas_Assign(AObj: TCanvas; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function Canvas_GetNamePath(AObj: TCanvas): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function Canvas_ClassName(AObj: TCanvas): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function Canvas_Equals(AObj: TCanvas; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function Canvas_GetHashCode(AObj: TCanvas): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function Canvas_ToString(AObj: TCanvas): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function Canvas_GetHandle(AObj: TCanvas): HDC; extdecl;
begin
  Result :=  AObj.Handle;
end;

procedure Canvas_SetHandle(AObj: TCanvas; AValue: HDC); extdecl;
begin
  AObj.Handle := AValue;
end;

function Canvas_GetBrush(AObj: TCanvas): TBrush; extdecl;
begin
  Result :=  AObj.Brush;
end;

procedure Canvas_SetBrush(AObj: TCanvas; AValue: TBrush); extdecl;
begin
  AObj.Brush := AValue;
end;

function Canvas_GetCopyMode(AObj: TCanvas): Integer; extdecl;
begin
  Result :=  AObj.CopyMode;
end;

procedure Canvas_SetCopyMode(AObj: TCanvas; AValue: Integer); extdecl;
begin
  AObj.CopyMode := AValue;
end;

function Canvas_GetFont(AObj: TCanvas): TFont; extdecl;
begin
  Result :=  AObj.Font;
end;

procedure Canvas_SetFont(AObj: TCanvas; AValue: TFont); extdecl;
begin
  AObj.Font := AValue;
end;

function Canvas_GetPen(AObj: TCanvas): TPen; extdecl;
begin
  Result :=  AObj.Pen;
end;

procedure Canvas_SetPen(AObj: TCanvas; AValue: TPen); extdecl;
begin
  AObj.Pen := AValue;
end;

procedure Canvas_SetOnChange(AObj: TCanvas; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnChange := nil;
    EventClass.Remove(AObj, geChange);
    Exit;
  end;
  AObj.OnChange := @EventClass.OnChange;
  EventClass.Add(AObj, geChange, AEventId);
end;


exports
  Canvas_Create,
  Canvas_Free,
  Canvas_Arc,
  Canvas_ArcTo,
  Canvas_AngleArc,
  Canvas_Chord,
  Canvas_Ellipse,
  Canvas_FloodFill,
  Canvas_LineTo,
  Canvas_MoveTo,
  Canvas_Pie,
  Canvas_Rectangle,
  Canvas_Refresh,
  Canvas_RoundRect,
  Canvas_TextExtent,
  Canvas_TextOut,
  Canvas_Lock,
  Canvas_TextHeight,
  Canvas_TextWidth,
  Canvas_Assign,
  Canvas_GetNamePath,
  Canvas_ClassName,
  Canvas_Equals,
  Canvas_GetHashCode,
  Canvas_ToString,
  Canvas_GetHandle,
  Canvas_SetHandle,
  Canvas_GetBrush,
  Canvas_SetBrush,
  Canvas_GetCopyMode,
  Canvas_SetCopyMode,
  Canvas_GetFont,
  Canvas_SetFont,
  Canvas_GetPen,
  Canvas_SetPen,
  Canvas_SetOnChange;

