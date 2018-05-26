
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function Action_Create(AOwner: TComponent): TAction; extdecl;
begin
  Result :=  TAction.Create(AOwner);
end;

procedure Action_Free(AObj: TAction); extdecl;
begin
  AObj.Free;
end;

function Action_Execute(AObj: TAction): LongBool; extdecl;
begin
  Result :=  AObj.Execute;
end;

function Action_Update(AObj: TAction): LongBool; extdecl;
begin
  Result :=  AObj.Update;
end;

function Action_HasParent(AObj: TAction): LongBool; extdecl;
begin
  Result :=  AObj.HasParent;
end;

function Action_FindComponent(AObj: TAction; AName: PChar): TComponent; extdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function Action_GetNamePath(AObj: TAction): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure Action_Assign(AObj: TAction; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function Action_ClassName(AObj: TAction): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function Action_Equals(AObj: TAction; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function Action_GetHashCode(AObj: TAction): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function Action_ToString(AObj: TAction): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function Action_GetCaption(AObj: TAction): PChar; extdecl;
begin
  Result :=  PChar(AObj.Caption);
end;

procedure Action_SetCaption(AObj: TAction; AValue: PChar); extdecl;
begin
  AObj.Caption := AValue;
end;

function Action_GetChecked(AObj: TAction): LongBool; extdecl;
begin
  Result :=  AObj.Checked;
end;

procedure Action_SetChecked(AObj: TAction; AValue: LongBool); extdecl;
begin
  AObj.Checked := AValue;
end;

function Action_GetEnabled(AObj: TAction): LongBool; extdecl;
begin
  Result :=  AObj.Enabled;
end;

procedure Action_SetEnabled(AObj: TAction; AValue: LongBool); extdecl;
begin
  AObj.Enabled := AValue;
end;

function Action_GetGroupIndex(AObj: TAction): Integer; extdecl;
begin
  Result :=  AObj.GroupIndex;
end;

procedure Action_SetGroupIndex(AObj: TAction; AValue: Integer); extdecl;
begin
  AObj.GroupIndex := AValue;
end;

function Action_GetHint(AObj: TAction): PChar; extdecl;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure Action_SetHint(AObj: TAction; AValue: PChar); extdecl;
begin
  AObj.Hint := AValue;
end;

function Action_GetImageIndex(AObj: TAction): TImageIndex; extdecl;
begin
  Result :=  AObj.ImageIndex;
end;

procedure Action_SetImageIndex(AObj: TAction; AValue: TImageIndex); extdecl;
begin
  AObj.ImageIndex := AValue;
end;

function Action_GetShortCut(AObj: TAction): TShortCut; extdecl;
begin
  Result :=  AObj.ShortCut;
end;

procedure Action_SetShortCut(AObj: TAction; AValue: TShortCut); extdecl;
begin
  AObj.ShortCut := AValue;
end;

function Action_GetVisible(AObj: TAction): LongBool; extdecl;
begin
  Result :=  AObj.Visible;
end;

procedure Action_SetVisible(AObj: TAction; AValue: LongBool); extdecl;
begin
  AObj.Visible := AValue;
end;

procedure Action_SetOnExecute(AObj: TAction; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnExecute := nil;
    EventClass.Remove(AObj, geExecute);
    Exit;
  end;
  AObj.OnExecute := @EventClass.OnExecute;
  EventClass.Add(AObj, geExecute, AEventId);
end;

procedure Action_SetOnUpdate(AObj: TAction; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnUpdate := nil;
    EventClass.Remove(AObj, geUpdate);
    Exit;
  end;
  AObj.OnUpdate := @EventClass.OnUpdate;
  EventClass.Add(AObj, geUpdate, AEventId);
end;

function Action_GetIndex(AObj: TAction): Integer; extdecl;
begin
  Result :=  AObj.Index;
end;

procedure Action_SetIndex(AObj: TAction; AValue: Integer); extdecl;
begin
  AObj.Index := AValue;
end;

function Action_GetComponentCount(AObj: TAction): Integer; extdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function Action_GetComponentIndex(AObj: TAction): Integer; extdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure Action_SetComponentIndex(AObj: TAction; AValue: Integer); extdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function Action_GetOwner(AObj: TAction): TComponent; extdecl;
begin
  Result :=  AObj.Owner;
end;

function Action_GetName(AObj: TAction): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure Action_SetName(AObj: TAction; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function Action_GetTag(AObj: TAction): NativeInt; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure Action_SetTag(AObj: TAction; AValue: NativeInt); extdecl;
begin
  AObj.Tag := AValue;
end;

function Action_GetComponents(AObj: TAction; AIndex: Integer): TComponent; extdecl;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  Action_Create,
  Action_Free,
  Action_Execute,
  Action_Update,
  Action_HasParent,
  Action_FindComponent,
  Action_GetNamePath,
  Action_Assign,
  Action_ClassName,
  Action_Equals,
  Action_GetHashCode,
  Action_ToString,
  Action_GetCaption,
  Action_SetCaption,
  Action_GetChecked,
  Action_SetChecked,
  Action_GetEnabled,
  Action_SetEnabled,
  Action_GetGroupIndex,
  Action_SetGroupIndex,
  Action_GetHint,
  Action_SetHint,
  Action_GetImageIndex,
  Action_SetImageIndex,
  Action_GetShortCut,
  Action_SetShortCut,
  Action_GetVisible,
  Action_SetVisible,
  Action_SetOnExecute,
  Action_SetOnUpdate,
  Action_GetIndex,
  Action_SetIndex,
  Action_GetComponentCount,
  Action_GetComponentIndex,
  Action_SetComponentIndex,
  Action_GetOwner,
  Action_GetName,
  Action_SetName,
  Action_GetTag,
  Action_SetTag,
  Action_GetComponents;

