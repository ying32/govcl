
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function CustomHint_Create(AOwner: TComponent): TCustomHint; stdcall;
begin
  Result :=  TCustomHint.Create(AOwner);
end;

procedure CustomHint_Free(AObj: TCustomHint); stdcall;
begin
  AObj.Free;
end;

procedure CustomHint_ShowHint(AObj: TCustomHint); stdcall;
begin
  AObj.ShowHint;
end;

procedure CustomHint_HideHint(AObj: TCustomHint); stdcall;
begin
  AObj.HideHint;
end;

function CustomHint_FindComponent(AObj: TCustomHint; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function CustomHint_GetNamePath(AObj: TCustomHint): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function CustomHint_HasParent(AObj: TCustomHint): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure CustomHint_Assign(AObj: TCustomHint; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function CustomHint_ClassName(AObj: TCustomHint): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function CustomHint_Equals(AObj: TCustomHint; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function CustomHint_GetHashCode(AObj: TCustomHint): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function CustomHint_ToString(AObj: TCustomHint): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function CustomHint_GetShowingHint(AObj: TCustomHint): LongBool; stdcall;
begin
  Result :=  AObj.ShowingHint;
end;

function CustomHint_GetTitle(AObj: TCustomHint): PChar; stdcall;
begin
  Result :=  PChar(AObj.Title);
end;

procedure CustomHint_SetTitle(AObj: TCustomHint; AValue: PChar); stdcall;
begin
  AObj.Title := AValue;
end;

function CustomHint_GetDescription(AObj: TCustomHint): PChar; stdcall;
begin
  Result :=  PChar(AObj.Description);
end;

procedure CustomHint_SetDescription(AObj: TCustomHint; AValue: PChar); stdcall;
begin
  AObj.Description := AValue;
end;

function CustomHint_GetImageIndex(AObj: TCustomHint): TImageIndex; stdcall;
begin
  Result :=  AObj.ImageIndex;
end;

procedure CustomHint_SetImageIndex(AObj: TCustomHint; AValue: TImageIndex); stdcall;
begin
  AObj.ImageIndex := AValue;
end;

function CustomHint_GetImages(AObj: TCustomHint): TImageList; stdcall;
begin
  Result :=  AObj.Images;
end;

procedure CustomHint_SetImages(AObj: TCustomHint; AValue: TImageList); stdcall;
begin
  AObj.Images := AValue;
end;

function CustomHint_GetStyle(AObj: TCustomHint): TBalloonHintStyle; stdcall;
begin
  Result :=  AObj.Style;
end;

procedure CustomHint_SetStyle(AObj: TCustomHint; AValue: TBalloonHintStyle); stdcall;
begin
  AObj.Style := AValue;
end;

function CustomHint_GetDelay(AObj: TCustomHint): Cardinal; stdcall;
begin
  Result :=  AObj.Delay;
end;

procedure CustomHint_SetDelay(AObj: TCustomHint; AValue: Cardinal); stdcall;
begin
  AObj.Delay := AValue;
end;

function CustomHint_GetHideAfter(AObj: TCustomHint): Integer; stdcall;
begin
  Result :=  AObj.HideAfter;
end;

procedure CustomHint_SetHideAfter(AObj: TCustomHint; AValue: Integer); stdcall;
begin
  AObj.HideAfter := AValue;
end;

function CustomHint_GetComponentCount(AObj: TCustomHint): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function CustomHint_GetComponentIndex(AObj: TCustomHint): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure CustomHint_SetComponentIndex(AObj: TCustomHint; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function CustomHint_GetOwner(AObj: TCustomHint): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function CustomHint_GetName(AObj: TCustomHint): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure CustomHint_SetName(AObj: TCustomHint; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function CustomHint_GetTag(AObj: TCustomHint): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure CustomHint_SetTag(AObj: TCustomHint; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function CustomHint_GetComponents(AObj: TCustomHint; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  CustomHint_Create,
  CustomHint_Free,
  CustomHint_ShowHint,
  CustomHint_HideHint,
  CustomHint_FindComponent,
  CustomHint_GetNamePath,
  CustomHint_HasParent,
  CustomHint_Assign,
  CustomHint_ClassName,
  CustomHint_Equals,
  CustomHint_GetHashCode,
  CustomHint_ToString,
  CustomHint_GetShowingHint,
  CustomHint_GetTitle,
  CustomHint_SetTitle,
  CustomHint_GetDescription,
  CustomHint_SetDescription,
  CustomHint_GetImageIndex,
  CustomHint_SetImageIndex,
  CustomHint_GetImages,
  CustomHint_SetImages,
  CustomHint_GetStyle,
  CustomHint_SetStyle,
  CustomHint_GetDelay,
  CustomHint_SetDelay,
  CustomHint_GetHideAfter,
  CustomHint_SetHideAfter,
  CustomHint_GetComponentCount,
  CustomHint_GetComponentIndex,
  CustomHint_SetComponentIndex,
  CustomHint_GetOwner,
  CustomHint_GetName,
  CustomHint_SetName,
  CustomHint_GetTag,
  CustomHint_SetTag,
  CustomHint_GetComponents;

