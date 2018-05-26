
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function TrayIcon_Create(AOwner: TComponent): TTrayIcon; stdcall;
begin
  Result :=  TTrayIcon.Create(AOwner);
end;

procedure TrayIcon_Free(AObj: TTrayIcon); stdcall;
begin
  AObj.Free;
end;

procedure TrayIcon_Refresh(AObj: TTrayIcon); stdcall;
begin
  AObj.Refresh;
end;

procedure TrayIcon_SetDefaultIcon(AObj: TTrayIcon); stdcall;
begin
  AObj.SetDefaultIcon;
end;

procedure TrayIcon_ShowBalloonHint(AObj: TTrayIcon); stdcall;
begin
  AObj.ShowBalloonHint;
end;

function TrayIcon_FindComponent(AObj: TTrayIcon; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function TrayIcon_GetNamePath(AObj: TTrayIcon): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function TrayIcon_HasParent(AObj: TTrayIcon): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure TrayIcon_Assign(AObj: TTrayIcon; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function TrayIcon_ClassName(AObj: TTrayIcon): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function TrayIcon_Equals(AObj: TTrayIcon; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function TrayIcon_GetHashCode(AObj: TTrayIcon): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function TrayIcon_ToString(AObj: TTrayIcon): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function TrayIcon_GetAnimate(AObj: TTrayIcon): LongBool; stdcall;
begin
  Result :=  AObj.Animate;
end;

procedure TrayIcon_SetAnimate(AObj: TTrayIcon; AValue: LongBool); stdcall;
begin
  AObj.Animate := AValue;
end;

function TrayIcon_GetAnimateInterval(AObj: TTrayIcon): Cardinal; stdcall;
begin
  Result :=  AObj.AnimateInterval;
end;

procedure TrayIcon_SetAnimateInterval(AObj: TTrayIcon; AValue: Cardinal); stdcall;
begin
  AObj.AnimateInterval := AValue;
end;

function TrayIcon_GetHint(AObj: TTrayIcon): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure TrayIcon_SetHint(AObj: TTrayIcon; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function TrayIcon_GetBalloonHint(AObj: TTrayIcon): PChar; stdcall;
begin
  Result :=  PChar(AObj.BalloonHint);
end;

procedure TrayIcon_SetBalloonHint(AObj: TTrayIcon; AValue: PChar); stdcall;
begin
  AObj.BalloonHint := AValue;
end;

function TrayIcon_GetBalloonTitle(AObj: TTrayIcon): PChar; stdcall;
begin
  Result :=  PChar(AObj.BalloonTitle);
end;

procedure TrayIcon_SetBalloonTitle(AObj: TTrayIcon; AValue: PChar); stdcall;
begin
  AObj.BalloonTitle := AValue;
end;

function TrayIcon_GetBalloonTimeout(AObj: TTrayIcon): Integer; stdcall;
begin
  Result :=  AObj.BalloonTimeout;
end;

procedure TrayIcon_SetBalloonTimeout(AObj: TTrayIcon; AValue: Integer); stdcall;
begin
  AObj.BalloonTimeout := AValue;
end;

function TrayIcon_GetBalloonFlags(AObj: TTrayIcon): TBalloonFlags; stdcall;
begin
  Result :=  AObj.BalloonFlags;
end;

procedure TrayIcon_SetBalloonFlags(AObj: TTrayIcon; AValue: TBalloonFlags); stdcall;
begin
  AObj.BalloonFlags := AValue;
end;

function TrayIcon_GetIcon(AObj: TTrayIcon): TIcon; stdcall;
begin
  Result :=  AObj.Icon;
end;

procedure TrayIcon_SetIcon(AObj: TTrayIcon; AValue: TIcon); stdcall;
begin
  AObj.Icon := AValue;
end;

function TrayIcon_GetIconIndex(AObj: TTrayIcon): Integer; stdcall;
begin
  Result :=  AObj.IconIndex;
end;

procedure TrayIcon_SetIconIndex(AObj: TTrayIcon; AValue: Integer); stdcall;
begin
  AObj.IconIndex := AValue;
end;

function TrayIcon_GetPopupMenu(AObj: TTrayIcon): TPopupMenu; stdcall;
begin
  Result :=  AObj.PopupMenu;
end;

procedure TrayIcon_SetPopupMenu(AObj: TTrayIcon; AValue: TPopupMenu); stdcall;
begin
  AObj.PopupMenu := AValue;
end;

function TrayIcon_GetVisible(AObj: TTrayIcon): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure TrayIcon_SetVisible(AObj: TTrayIcon; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

procedure TrayIcon_SetOnBalloonClick(AObj: TTrayIcon; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnBalloonClick := nil;
    TEventClass.Remove(AObj, geBalloonClick);
    Exit;
  end;
  AObj.OnBalloonClick := TEventClass.OnBalloonClick;
  TEventClass.Add(AObj, geBalloonClick, AEventId);
end;

procedure TrayIcon_SetOnClick(AObj: TTrayIcon; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnClick := nil;
    TEventClass.Remove(AObj, geClick);
    Exit;
  end;
  AObj.OnClick := TEventClass.OnClick;
  TEventClass.Add(AObj, geClick, AEventId);
end;

procedure TrayIcon_SetOnDblClick(AObj: TTrayIcon; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnDblClick := nil;
    TEventClass.Remove(AObj, geDblClick);
    Exit;
  end;
  AObj.OnDblClick := TEventClass.OnDblClick;
  TEventClass.Add(AObj, geDblClick, AEventId);
end;

procedure TrayIcon_SetOnMouseMove(AObj: TTrayIcon; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseMove := nil;
    TEventClass.Remove(AObj, geMouseMove);
    Exit;
  end;
  AObj.OnMouseMove := TEventClass.OnMouseMove;
  TEventClass.Add(AObj, geMouseMove, AEventId);
end;

procedure TrayIcon_SetOnMouseUp(AObj: TTrayIcon; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseUp := nil;
    TEventClass.Remove(AObj, geMouseUp);
    Exit;
  end;
  AObj.OnMouseUp := TEventClass.OnMouseUp;
  TEventClass.Add(AObj, geMouseUp, AEventId);
end;

procedure TrayIcon_SetOnMouseDown(AObj: TTrayIcon; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseDown := nil;
    TEventClass.Remove(AObj, geMouseDown);
    Exit;
  end;
  AObj.OnMouseDown := TEventClass.OnMouseDown;
  TEventClass.Add(AObj, geMouseDown, AEventId);
end;

function TrayIcon_GetComponentCount(AObj: TTrayIcon): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function TrayIcon_GetComponentIndex(AObj: TTrayIcon): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure TrayIcon_SetComponentIndex(AObj: TTrayIcon; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function TrayIcon_GetOwner(AObj: TTrayIcon): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function TrayIcon_GetName(AObj: TTrayIcon): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure TrayIcon_SetName(AObj: TTrayIcon; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function TrayIcon_GetTag(AObj: TTrayIcon): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure TrayIcon_SetTag(AObj: TTrayIcon; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function TrayIcon_GetComponents(AObj: TTrayIcon; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  TrayIcon_Create,
  TrayIcon_Free,
  TrayIcon_Refresh,
  TrayIcon_SetDefaultIcon,
  TrayIcon_ShowBalloonHint,
  TrayIcon_FindComponent,
  TrayIcon_GetNamePath,
  TrayIcon_HasParent,
  TrayIcon_Assign,
  TrayIcon_ClassName,
  TrayIcon_Equals,
  TrayIcon_GetHashCode,
  TrayIcon_ToString,
  TrayIcon_GetAnimate,
  TrayIcon_SetAnimate,
  TrayIcon_GetAnimateInterval,
  TrayIcon_SetAnimateInterval,
  TrayIcon_GetHint,
  TrayIcon_SetHint,
  TrayIcon_GetBalloonHint,
  TrayIcon_SetBalloonHint,
  TrayIcon_GetBalloonTitle,
  TrayIcon_SetBalloonTitle,
  TrayIcon_GetBalloonTimeout,
  TrayIcon_SetBalloonTimeout,
  TrayIcon_GetBalloonFlags,
  TrayIcon_SetBalloonFlags,
  TrayIcon_GetIcon,
  TrayIcon_SetIcon,
  TrayIcon_GetIconIndex,
  TrayIcon_SetIconIndex,
  TrayIcon_GetPopupMenu,
  TrayIcon_SetPopupMenu,
  TrayIcon_GetVisible,
  TrayIcon_SetVisible,
  TrayIcon_SetOnBalloonClick,
  TrayIcon_SetOnClick,
  TrayIcon_SetOnDblClick,
  TrayIcon_SetOnMouseMove,
  TrayIcon_SetOnMouseUp,
  TrayIcon_SetOnMouseDown,
  TrayIcon_GetComponentCount,
  TrayIcon_GetComponentIndex,
  TrayIcon_SetComponentIndex,
  TrayIcon_GetOwner,
  TrayIcon_GetName,
  TrayIcon_SetName,
  TrayIcon_GetTag,
  TrayIcon_SetTag,
  TrayIcon_GetComponents;

