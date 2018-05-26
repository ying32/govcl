
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function Margins_Create(AOwner: TControl): TMargins; stdcall;
begin
  Result :=  TMargins.Create(AOwner);
end;

procedure Margins_Free(AObj: TMargins); stdcall;
begin
  AObj.Free;
end;

procedure Margins_SetBounds(AObj: TMargins; ALeft: Integer; ATop: Integer; ARight: Integer; ABottom: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, ARight, ABottom);
end;

procedure Margins_Assign(AObj: TMargins; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function Margins_GetNamePath(AObj: TMargins): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function Margins_ClassName(AObj: TMargins): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function Margins_Equals(AObj: TMargins; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function Margins_GetHashCode(AObj: TMargins): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function Margins_ToString(AObj: TMargins): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function Margins_GetControlLeft(AObj: TMargins): Integer; stdcall;
begin
  Result :=  AObj.ControlLeft;
end;

function Margins_GetControlTop(AObj: TMargins): Integer; stdcall;
begin
  Result :=  AObj.ControlTop;
end;

function Margins_GetControlWidth(AObj: TMargins): Integer; stdcall;
begin
  Result :=  AObj.ControlWidth;
end;

function Margins_GetControlHeight(AObj: TMargins): Integer; stdcall;
begin
  Result :=  AObj.ControlHeight;
end;

function Margins_GetExplicitLeft(AObj: TMargins): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function Margins_GetExplicitTop(AObj: TMargins): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function Margins_GetExplicitWidth(AObj: TMargins): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function Margins_GetExplicitHeight(AObj: TMargins): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

procedure Margins_SetOnChange(AObj: TMargins; AEventId: NativeUInt); stdcall;
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

function Margins_GetLeft(AObj: TMargins): TMarginSize; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure Margins_SetLeft(AObj: TMargins; AValue: TMarginSize); stdcall;
begin
  AObj.Left := AValue;
end;

function Margins_GetTop(AObj: TMargins): TMarginSize; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure Margins_SetTop(AObj: TMargins; AValue: TMarginSize); stdcall;
begin
  AObj.Top := AValue;
end;

function Margins_GetRight(AObj: TMargins): TMarginSize; stdcall;
begin
  Result :=  AObj.Right;
end;

procedure Margins_SetRight(AObj: TMargins; AValue: TMarginSize); stdcall;
begin
  AObj.Right := AValue;
end;

function Margins_GetBottom(AObj: TMargins): TMarginSize; stdcall;
begin
  Result :=  AObj.Bottom;
end;

procedure Margins_SetBottom(AObj: TMargins; AValue: TMarginSize); stdcall;
begin
  AObj.Bottom := AValue;
end;


exports
  Margins_Create,
  Margins_Free,
  Margins_SetBounds,
  Margins_Assign,
  Margins_GetNamePath,
  Margins_ClassName,
  Margins_Equals,
  Margins_GetHashCode,
  Margins_ToString,
  Margins_GetControlLeft,
  Margins_GetControlTop,
  Margins_GetControlWidth,
  Margins_GetControlHeight,
  Margins_GetExplicitLeft,
  Margins_GetExplicitTop,
  Margins_GetExplicitWidth,
  Margins_GetExplicitHeight,
  Margins_SetOnChange,
  Margins_GetLeft,
  Margins_SetLeft,
  Margins_GetTop,
  Margins_SetTop,
  Margins_GetRight,
  Margins_SetRight,
  Margins_GetBottom,
  Margins_SetBottom;

