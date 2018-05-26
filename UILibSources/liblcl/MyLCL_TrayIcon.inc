
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function TrayIcon_Create(AOwner: TComponent): TTrayIcon; extdecl;
begin
  Result :=  TTrayIcon.Create(AOwner);
end;

procedure TrayIcon_Free(AObj: TTrayIcon); extdecl;
begin
  AObj.Free;
end;

procedure TrayIcon_ShowBalloonHint(AObj: TTrayIcon); extdecl;
begin
  AObj.ShowBalloonHint;
end;

function TrayIcon_FindComponent(AObj: TTrayIcon; AName: PChar): TComponent; extdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function TrayIcon_GetNamePath(AObj: TTrayIcon): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function TrayIcon_HasParent(AObj: TTrayIcon): LongBool; extdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure TrayIcon_Assign(AObj: TTrayIcon; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function TrayIcon_ClassName(AObj: TTrayIcon): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function TrayIcon_Equals(AObj: TTrayIcon; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function TrayIcon_GetHashCode(AObj: TTrayIcon): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function TrayIcon_ToString(AObj: TTrayIcon): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function TrayIcon_GetHint(AObj: TTrayIcon): PChar; extdecl;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure TrayIcon_SetHint(AObj: TTrayIcon; AValue: PChar); extdecl;
begin
  AObj.Hint := AValue;
end;

function TrayIcon_GetBalloonHint(AObj: TTrayIcon): PChar; extdecl;
begin
  Result :=  PChar(AObj.BalloonHint);
end;

procedure TrayIcon_SetBalloonHint(AObj: TTrayIcon; AValue: PChar); extdecl;
begin
  AObj.BalloonHint := AValue;
end;

function TrayIcon_GetBalloonTitle(AObj: TTrayIcon): PChar; extdecl;
begin
  Result :=  PChar(AObj.BalloonTitle);
end;

procedure TrayIcon_SetBalloonTitle(AObj: TTrayIcon; AValue: PChar); extdecl;
begin
  AObj.BalloonTitle := AValue;
end;

function TrayIcon_GetBalloonTimeout(AObj: TTrayIcon): Integer; extdecl;
begin
  Result :=  AObj.BalloonTimeout;
end;

procedure TrayIcon_SetBalloonTimeout(AObj: TTrayIcon; AValue: Integer); extdecl;
begin
  AObj.BalloonTimeout := AValue;
end;

function TrayIcon_GetBalloonFlags(AObj: TTrayIcon): TBalloonFlags; extdecl;
begin
  Result :=  AObj.BalloonFlags;
end;

procedure TrayIcon_SetBalloonFlags(AObj: TTrayIcon; AValue: TBalloonFlags); extdecl;
begin
  AObj.BalloonFlags := AValue;
end;

function TrayIcon_GetIcon(AObj: TTrayIcon): TIcon; extdecl;
begin
  Result :=  AObj.Icon;
end;

procedure TrayIcon_SetIcon(AObj: TTrayIcon; AValue: TIcon); extdecl;
begin
  AObj.Icon := AValue;
end;

function TrayIcon_GetPopupMenu(AObj: TTrayIcon): TPopupMenu; extdecl;
begin
  Result :=  AObj.PopupMenu;
end;

procedure TrayIcon_SetPopupMenu(AObj: TTrayIcon; AValue: TPopupMenu); extdecl;
begin
  AObj.PopupMenu := AValue;
end;

function TrayIcon_GetVisible(AObj: TTrayIcon): LongBool; extdecl;
begin
  Result :=  AObj.Visible;
end;

procedure TrayIcon_SetVisible(AObj: TTrayIcon; AValue: LongBool); extdecl;
begin
  AObj.Visible := AValue;
end;

procedure TrayIcon_SetOnClick(AObj: TTrayIcon; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnClick := nil;
    EventClass.Remove(AObj, geClick);
    Exit;
  end;
  AObj.OnClick := @EventClass.OnClick;
  EventClass.Add(AObj, geClick, AEventId);
end;

procedure TrayIcon_SetOnDblClick(AObj: TTrayIcon; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnDblClick := nil;
    EventClass.Remove(AObj, geDblClick);
    Exit;
  end;
  AObj.OnDblClick := @EventClass.OnDblClick;
  EventClass.Add(AObj, geDblClick, AEventId);
end;

procedure TrayIcon_SetOnMouseMove(AObj: TTrayIcon; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseMove := nil;
    EventClass.Remove(AObj, geMouseMove);
    Exit;
  end;
  AObj.OnMouseMove := @EventClass.OnMouseMove;
  EventClass.Add(AObj, geMouseMove, AEventId);
end;

procedure TrayIcon_SetOnMouseUp(AObj: TTrayIcon; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseUp := nil;
    EventClass.Remove(AObj, geMouseUp);
    Exit;
  end;
  AObj.OnMouseUp := @EventClass.OnMouseUp;
  EventClass.Add(AObj, geMouseUp, AEventId);
end;

procedure TrayIcon_SetOnMouseDown(AObj: TTrayIcon; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseDown := nil;
    EventClass.Remove(AObj, geMouseDown);
    Exit;
  end;
  AObj.OnMouseDown := @EventClass.OnMouseDown;
  EventClass.Add(AObj, geMouseDown, AEventId);
end;

function TrayIcon_GetComponentCount(AObj: TTrayIcon): Integer; extdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function TrayIcon_GetComponentIndex(AObj: TTrayIcon): Integer; extdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure TrayIcon_SetComponentIndex(AObj: TTrayIcon; AValue: Integer); extdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function TrayIcon_GetOwner(AObj: TTrayIcon): TComponent; extdecl;
begin
  Result :=  AObj.Owner;
end;

function TrayIcon_GetName(AObj: TTrayIcon): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure TrayIcon_SetName(AObj: TTrayIcon; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function TrayIcon_GetTag(AObj: TTrayIcon): NativeInt; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure TrayIcon_SetTag(AObj: TTrayIcon; AValue: NativeInt); extdecl;
begin
  AObj.Tag := AValue;
end;

function TrayIcon_GetComponents(AObj: TTrayIcon; AIndex: Integer): TComponent; extdecl;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  TrayIcon_Create,
  TrayIcon_Free,
  TrayIcon_ShowBalloonHint,
  TrayIcon_FindComponent,
  TrayIcon_GetNamePath,
  TrayIcon_HasParent,
  TrayIcon_Assign,
  TrayIcon_ClassName,
  TrayIcon_Equals,
  TrayIcon_GetHashCode,
  TrayIcon_ToString,
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
  TrayIcon_GetPopupMenu,
  TrayIcon_SetPopupMenu,
  TrayIcon_GetVisible,
  TrayIcon_SetVisible,
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

