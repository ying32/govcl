
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function BalloonHint_Create(AOwner: TComponent): TBalloonHint; stdcall;
begin
  Result :=  TBalloonHint.Create(AOwner);
end;

procedure BalloonHint_Free(AObj: TBalloonHint); stdcall;
begin
  AObj.Free;
end;

procedure BalloonHint_ShowHint(AObj: TBalloonHint); stdcall;
begin
  AObj.ShowHint;
end;

procedure BalloonHint_HideHint(AObj: TBalloonHint); stdcall;
begin
  AObj.HideHint;
end;

function BalloonHint_FindComponent(AObj: TBalloonHint; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function BalloonHint_GetNamePath(AObj: TBalloonHint): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function BalloonHint_HasParent(AObj: TBalloonHint): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure BalloonHint_Assign(AObj: TBalloonHint; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function BalloonHint_ClassName(AObj: TBalloonHint): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function BalloonHint_Equals(AObj: TBalloonHint; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function BalloonHint_GetHashCode(AObj: TBalloonHint): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function BalloonHint_ToString(AObj: TBalloonHint): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function BalloonHint_GetShowingHint(AObj: TBalloonHint): LongBool; stdcall;
begin
  Result :=  AObj.ShowingHint;
end;

function BalloonHint_GetTitle(AObj: TBalloonHint): PChar; stdcall;
begin
  Result :=  PChar(AObj.Title);
end;

procedure BalloonHint_SetTitle(AObj: TBalloonHint; AValue: PChar); stdcall;
begin
  AObj.Title := AValue;
end;

function BalloonHint_GetDescription(AObj: TBalloonHint): PChar; stdcall;
begin
  Result :=  PChar(AObj.Description);
end;

procedure BalloonHint_SetDescription(AObj: TBalloonHint; AValue: PChar); stdcall;
begin
  AObj.Description := AValue;
end;

function BalloonHint_GetImageIndex(AObj: TBalloonHint): TImageIndex; stdcall;
begin
  Result :=  AObj.ImageIndex;
end;

procedure BalloonHint_SetImageIndex(AObj: TBalloonHint; AValue: TImageIndex); stdcall;
begin
  AObj.ImageIndex := AValue;
end;

function BalloonHint_GetImages(AObj: TBalloonHint): TImageList; stdcall;
begin
  Result :=  AObj.Images;
end;

procedure BalloonHint_SetImages(AObj: TBalloonHint; AValue: TImageList); stdcall;
begin
  AObj.Images := AValue;
end;

function BalloonHint_GetStyle(AObj: TBalloonHint): TBalloonHintStyle; stdcall;
begin
  Result :=  AObj.Style;
end;

procedure BalloonHint_SetStyle(AObj: TBalloonHint; AValue: TBalloonHintStyle); stdcall;
begin
  AObj.Style := AValue;
end;

function BalloonHint_GetDelay(AObj: TBalloonHint): Cardinal; stdcall;
begin
  Result :=  AObj.Delay;
end;

procedure BalloonHint_SetDelay(AObj: TBalloonHint; AValue: Cardinal); stdcall;
begin
  AObj.Delay := AValue;
end;

function BalloonHint_GetHideAfter(AObj: TBalloonHint): Integer; stdcall;
begin
  Result :=  AObj.HideAfter;
end;

procedure BalloonHint_SetHideAfter(AObj: TBalloonHint; AValue: Integer); stdcall;
begin
  AObj.HideAfter := AValue;
end;

function BalloonHint_GetComponentCount(AObj: TBalloonHint): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function BalloonHint_GetComponentIndex(AObj: TBalloonHint): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure BalloonHint_SetComponentIndex(AObj: TBalloonHint; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function BalloonHint_GetOwner(AObj: TBalloonHint): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function BalloonHint_GetName(AObj: TBalloonHint): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure BalloonHint_SetName(AObj: TBalloonHint; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function BalloonHint_GetTag(AObj: TBalloonHint): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure BalloonHint_SetTag(AObj: TBalloonHint; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function BalloonHint_GetComponents(AObj: TBalloonHint; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  BalloonHint_Create,
  BalloonHint_Free,
  BalloonHint_ShowHint,
  BalloonHint_HideHint,
  BalloonHint_FindComponent,
  BalloonHint_GetNamePath,
  BalloonHint_HasParent,
  BalloonHint_Assign,
  BalloonHint_ClassName,
  BalloonHint_Equals,
  BalloonHint_GetHashCode,
  BalloonHint_ToString,
  BalloonHint_GetShowingHint,
  BalloonHint_GetTitle,
  BalloonHint_SetTitle,
  BalloonHint_GetDescription,
  BalloonHint_SetDescription,
  BalloonHint_GetImageIndex,
  BalloonHint_SetImageIndex,
  BalloonHint_GetImages,
  BalloonHint_SetImages,
  BalloonHint_GetStyle,
  BalloonHint_SetStyle,
  BalloonHint_GetDelay,
  BalloonHint_SetDelay,
  BalloonHint_GetHideAfter,
  BalloonHint_SetHideAfter,
  BalloonHint_GetComponentCount,
  BalloonHint_GetComponentIndex,
  BalloonHint_SetComponentIndex,
  BalloonHint_GetOwner,
  BalloonHint_GetName,
  BalloonHint_SetName,
  BalloonHint_GetTag,
  BalloonHint_SetTag,
  BalloonHint_GetComponents;

