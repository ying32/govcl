
//----------------------------------------
// Copyright © ying32. All Rights Reserved.
// 
// Licensed under Apache License 2.0
//
//----------------------------------------
//----------------------------------------
// 用来兼容Delphi的TMargins的，不是自动生成
//----------------------------------------

function Margins_Create(AOwner: TControl): TControlBorderSpacing; extdecl;
begin
  Result :=  TControlBorderSpacing.Create(AOwner);
end;

procedure Margins_Free(AObj: TControlBorderSpacing); extdecl;
begin
  AObj.Free;
end;

//procedure Margins_SetBounds(AObj: TControlBorderSpacing; ALeft: Integer; ATop: Integer; ARight: Integer; ABottom: Integer); extdecl;
//begin
  // 不支持此方法
  //AObj.SetBounds(ALeft, ATop, ARight, ABottom);
//end;

procedure Margins_Assign(AObj: TControlBorderSpacing; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function Margins_GetNamePath(AObj: TControlBorderSpacing): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

//procedure Margins_DisposeOf(AObj: TControlBorderSpacing); extdecl;
//begin
//  AObj.DisposeOf;
//end;

function Margins_ClassType(AObj: TControlBorderSpacing): TClass; extdecl;
begin
  Result :=  AObj.ClassType;
end;

function Margins_ClassName(AObj: TControlBorderSpacing): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function Margins_InstanceSize(AObj: TControlBorderSpacing): Integer; extdecl;
begin
  Result :=  AObj.InstanceSize;
end;

function Margins_InheritsFrom(AObj: TControlBorderSpacing; AClass: TClass): LongBool; extdecl;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function Margins_Equals(AObj: TControlBorderSpacing; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function Margins_GetHashCode(AObj: TControlBorderSpacing): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function Margins_ToString(AObj: TControlBorderSpacing): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function Margins_GetControlLeft(AObj: TControlBorderSpacing): Integer; extdecl;
begin
  Result :=  AObj.ControlLeft;
end;

function Margins_GetControlTop(AObj: TControlBorderSpacing): Integer; extdecl;
begin
  Result :=  AObj.ControlTop;
end;

function Margins_GetControlWidth(AObj: TControlBorderSpacing): Integer; extdecl;
begin
  Result :=  AObj.ControlWidth;
end;

function Margins_GetControlHeight(AObj: TControlBorderSpacing): Integer; extdecl;
begin
  Result :=  AObj.ControlHeight;
end;

//function Margins_GetExplicitLeft(AObj: TControlBorderSpacing): Integer; extdecl;
//begin
//  Result :=  AObj.ExplicitLeft;
//end;

//function Margins_GetExplicitTop(AObj: TControlBorderSpacing): Integer; extdecl;
//begin
//  Result :=  AObj.ExplicitTop;
//end;

//function Margins_GetExplicitWidth(AObj: TControlBorderSpacing): Integer; extdecl;
//begin
//  Result :=  AObj.ExplicitWidth;
//end;

//function Margins_GetExplicitHeight(AObj: TControlBorderSpacing): Integer; extdecl;
//begin
//  Result :=  AObj.ExplicitHeight;
//end;

procedure Margins_SetOnChange(AObj: TControlBorderSpacing; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnChange := nil;
    TEventClass.Remove(AObj, @TEventClass.OnChange);
    Exit;
  end;
  AObj.OnChange := TEventClass.OnChange;
  TEventClass.Add(AObj, @TEventClass.OnChange, AEventId);
end;

function Margins_GetLeft(AObj: TControlBorderSpacing): TSpacingSize; extdecl;
begin
  Result :=  AObj.Left;
end;

procedure Margins_SetLeft(AObj: TControlBorderSpacing; AValue: TSpacingSize); extdecl;
begin
  AObj.Left := AValue;
end;

function Margins_GetTop(AObj: TControlBorderSpacing): TSpacingSize; extdecl;
begin
  Result :=  AObj.Top;
end;

procedure Margins_SetTop(AObj: TControlBorderSpacing; AValue: TSpacingSize); extdecl;
begin
  AObj.Top := AValue;
end;

function Margins_GetRight(AObj: TControlBorderSpacing): TSpacingSize; extdecl;
begin
  Result :=  AObj.Right;
end;

procedure Margins_SetRight(AObj: TControlBorderSpacing; AValue: TSpacingSize); extdecl;
begin
  AObj.Right := AValue;
end;

function Margins_GetBottom(AObj: TControlBorderSpacing): TSpacingSize; extdecl;
begin
  Result :=  AObj.Bottom;
end;

procedure Margins_SetBottom(AObj: TControlBorderSpacing; AValue: TSpacingSize); extdecl;
begin
  AObj.Bottom := AValue;
end;

function Margins_StaticClassType: TClass; extdecl;
begin
  Result := TControlBorderSpacing;
end;

exports
  Margins_Create,
  Margins_Free,
  //Margins_SetBounds,
  Margins_Assign,
  Margins_GetNamePath,
  //Margins_DisposeOf,
  Margins_ClassType,
  Margins_ClassName,
  Margins_InstanceSize,
  Margins_InheritsFrom,
  Margins_Equals,
  Margins_GetHashCode,
  Margins_ToString,
  Margins_GetControlLeft,
  Margins_GetControlTop,
  Margins_GetControlWidth,
  Margins_GetControlHeight,
  //Margins_GetExplicitLeft,
  //Margins_GetExplicitTop,
  //Margins_GetExplicitWidth,
  //Margins_GetExplicitHeight,
  Margins_SetOnChange,
  Margins_GetLeft,
  Margins_SetLeft,
  Margins_GetTop,
  Margins_SetTop,
  Margins_GetRight,
  Margins_SetRight,
  Margins_GetBottom,
  Margins_SetBottom,
  Margins_StaticClassType;

