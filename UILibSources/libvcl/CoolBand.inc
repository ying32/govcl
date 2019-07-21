
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function CoolBand_Create(AOwner: TCollection): TCoolBand; stdcall;
begin
  Result :=  TCoolBand.Create(AOwner);
end;

procedure CoolBand_Free(AObj: TCoolBand); stdcall;
begin
  AObj.Free;
end;

procedure CoolBand_Assign(AObj: TCoolBand; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function CoolBand_GetNamePath(AObj: TCoolBand): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure CoolBand_DisposeOf(AObj: TCoolBand); stdcall;
begin
  AObj.DisposeOf;
end;

function CoolBand_ClassType(AObj: TCoolBand): TClass; stdcall;
begin
  Result :=  AObj.ClassType;
end;

function CoolBand_ClassName(AObj: TCoolBand): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function CoolBand_InstanceSize(AObj: TCoolBand): Integer; stdcall;
begin
  Result :=  AObj.InstanceSize;
end;

function CoolBand_InheritsFrom(AObj: TCoolBand; AClass: TClass): LongBool; stdcall;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function CoolBand_Equals(AObj: TCoolBand; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function CoolBand_GetHashCode(AObj: TCoolBand): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function CoolBand_ToString(AObj: TCoolBand): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function CoolBand_GetHeight(AObj: TCoolBand): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

function CoolBand_GetBitmap(AObj: TCoolBand): TBitmap; stdcall;
begin
  Result :=  AObj.Bitmap;
end;

procedure CoolBand_SetBitmap(AObj: TCoolBand; AValue: TBitmap); stdcall;
begin
  AObj.Bitmap := AValue;
end;

function CoolBand_GetBorderStyle(AObj: TCoolBand): TBorderStyle; stdcall;
begin
  Result :=  AObj.BorderStyle;
end;

procedure CoolBand_SetBorderStyle(AObj: TCoolBand; AValue: TBorderStyle); stdcall;
begin
  AObj.BorderStyle := AValue;
end;

function CoolBand_GetBreak(AObj: TCoolBand): LongBool; stdcall;
begin
  Result :=  AObj.Break;
end;

procedure CoolBand_SetBreak(AObj: TCoolBand; AValue: LongBool); stdcall;
begin
  AObj.Break := AValue;
end;

function CoolBand_GetColor(AObj: TCoolBand): TColor; stdcall;
begin
  Result :=  AObj.Color;
end;

procedure CoolBand_SetColor(AObj: TCoolBand; AValue: TColor); stdcall;
begin
  AObj.Color := AValue;
end;

function CoolBand_GetControl(AObj: TCoolBand): TWinControl; stdcall;
begin
  Result :=  AObj.Control;
end;

procedure CoolBand_SetControl(AObj: TCoolBand; AValue: TWinControl); stdcall;
begin
  AObj.Control := AValue;
end;

function CoolBand_GetFixedBackground(AObj: TCoolBand): LongBool; stdcall;
begin
  Result :=  AObj.FixedBackground;
end;

procedure CoolBand_SetFixedBackground(AObj: TCoolBand; AValue: LongBool); stdcall;
begin
  AObj.FixedBackground := AValue;
end;

function CoolBand_GetFixedSize(AObj: TCoolBand): LongBool; stdcall;
begin
  Result :=  AObj.FixedSize;
end;

procedure CoolBand_SetFixedSize(AObj: TCoolBand; AValue: LongBool); stdcall;
begin
  AObj.FixedSize := AValue;
end;

function CoolBand_GetHorizontalOnly(AObj: TCoolBand): LongBool; stdcall;
begin
  Result :=  AObj.HorizontalOnly;
end;

procedure CoolBand_SetHorizontalOnly(AObj: TCoolBand; AValue: LongBool); stdcall;
begin
  AObj.HorizontalOnly := AValue;
end;

function CoolBand_GetImageIndex(AObj: TCoolBand): TImageIndex; stdcall;
begin
  Result :=  AObj.ImageIndex;
end;

procedure CoolBand_SetImageIndex(AObj: TCoolBand; AValue: TImageIndex); stdcall;
begin
  AObj.ImageIndex := AValue;
end;

function CoolBand_GetMinHeight(AObj: TCoolBand): Integer; stdcall;
begin
  Result :=  AObj.MinHeight;
end;

procedure CoolBand_SetMinHeight(AObj: TCoolBand; AValue: Integer); stdcall;
begin
  AObj.MinHeight := AValue;
end;

function CoolBand_GetMinWidth(AObj: TCoolBand): Integer; stdcall;
begin
  Result :=  AObj.MinWidth;
end;

procedure CoolBand_SetMinWidth(AObj: TCoolBand; AValue: Integer); stdcall;
begin
  AObj.MinWidth := AValue;
end;

function CoolBand_GetParentColor(AObj: TCoolBand): LongBool; stdcall;
begin
  Result :=  AObj.ParentColor;
end;

procedure CoolBand_SetParentColor(AObj: TCoolBand; AValue: LongBool); stdcall;
begin
  AObj.ParentColor := AValue;
end;

function CoolBand_GetParentBitmap(AObj: TCoolBand): LongBool; stdcall;
begin
  Result :=  AObj.ParentBitmap;
end;

procedure CoolBand_SetParentBitmap(AObj: TCoolBand; AValue: LongBool); stdcall;
begin
  AObj.ParentBitmap := AValue;
end;

function CoolBand_GetText(AObj: TCoolBand): PChar; stdcall;
begin
  Result :=  PChar(AObj.Text);
end;

procedure CoolBand_SetText(AObj: TCoolBand; AValue: PChar); stdcall;
begin
  AObj.Text := AValue;
end;

function CoolBand_GetVisible(AObj: TCoolBand): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure CoolBand_SetVisible(AObj: TCoolBand; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

function CoolBand_GetWidth(AObj: TCoolBand): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure CoolBand_SetWidth(AObj: TCoolBand; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function CoolBand_GetCollection(AObj: TCoolBand): TCollection; stdcall;
begin
  Result :=  AObj.Collection;
end;

procedure CoolBand_SetCollection(AObj: TCoolBand; AValue: TCollection); stdcall;
begin
  AObj.Collection := AValue;
end;

function CoolBand_GetIndex(AObj: TCoolBand): Integer; stdcall;
begin
  Result :=  AObj.Index;
end;

procedure CoolBand_SetIndex(AObj: TCoolBand; AValue: Integer); stdcall;
begin
  AObj.Index := AValue;
end;

function CoolBand_GetDisplayName(AObj: TCoolBand): PChar; stdcall;
begin
  Result :=  PChar(AObj.DisplayName);
end;

procedure CoolBand_SetDisplayName(AObj: TCoolBand; AValue: PChar); stdcall;
begin
  AObj.DisplayName := AValue;
end;

function CoolBand_StaticClassType: TClass; stdcall;
begin
  Result := TCoolBand;
end;

exports
  CoolBand_Create,
  CoolBand_Free,
  CoolBand_Assign,
  CoolBand_GetNamePath,
  CoolBand_DisposeOf,
  CoolBand_ClassType,
  CoolBand_ClassName,
  CoolBand_InstanceSize,
  CoolBand_InheritsFrom,
  CoolBand_Equals,
  CoolBand_GetHashCode,
  CoolBand_ToString,
  CoolBand_GetHeight,
  CoolBand_GetBitmap,
  CoolBand_SetBitmap,
  CoolBand_GetBorderStyle,
  CoolBand_SetBorderStyle,
  CoolBand_GetBreak,
  CoolBand_SetBreak,
  CoolBand_GetColor,
  CoolBand_SetColor,
  CoolBand_GetControl,
  CoolBand_SetControl,
  CoolBand_GetFixedBackground,
  CoolBand_SetFixedBackground,
  CoolBand_GetFixedSize,
  CoolBand_SetFixedSize,
  CoolBand_GetHorizontalOnly,
  CoolBand_SetHorizontalOnly,
  CoolBand_GetImageIndex,
  CoolBand_SetImageIndex,
  CoolBand_GetMinHeight,
  CoolBand_SetMinHeight,
  CoolBand_GetMinWidth,
  CoolBand_SetMinWidth,
  CoolBand_GetParentColor,
  CoolBand_SetParentColor,
  CoolBand_GetParentBitmap,
  CoolBand_SetParentBitmap,
  CoolBand_GetText,
  CoolBand_SetText,
  CoolBand_GetVisible,
  CoolBand_SetVisible,
  CoolBand_GetWidth,
  CoolBand_SetWidth,
  CoolBand_GetCollection,
  CoolBand_SetCollection,
  CoolBand_GetIndex,
  CoolBand_SetIndex,
  CoolBand_GetDisplayName,
  CoolBand_SetDisplayName,
  CoolBand_StaticClassType;

