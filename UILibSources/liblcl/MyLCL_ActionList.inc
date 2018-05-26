
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function ActionList_Create(AOwner: TComponent): TActionList; extdecl;
begin
  Result :=  TActionList.Create(AOwner);
end;

procedure ActionList_Free(AObj: TActionList); extdecl;
begin
  AObj.Free;
end;

function ActionList_FindComponent(AObj: TActionList; AName: PChar): TComponent; extdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function ActionList_GetNamePath(AObj: TActionList): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function ActionList_HasParent(AObj: TActionList): LongBool; extdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure ActionList_Assign(AObj: TActionList; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function ActionList_ClassName(AObj: TActionList): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function ActionList_Equals(AObj: TActionList; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function ActionList_GetHashCode(AObj: TActionList): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function ActionList_ToString(AObj: TActionList): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function ActionList_GetImages(AObj: TActionList): TCustomImageList; extdecl;
begin
  Result :=  AObj.Images;
end;

procedure ActionList_SetImages(AObj: TActionList; AValue: TCustomImageList); extdecl;
begin
  AObj.Images := AValue;
end;

function ActionList_GetState(AObj: TActionList): TActionListState; extdecl;
begin
  Result :=  AObj.State;
end;

procedure ActionList_SetState(AObj: TActionList; AValue: TActionListState); extdecl;
begin
  AObj.State := AValue;
end;

procedure ActionList_SetOnChange(AObj: TActionList; AEventId: NativeUInt); extdecl;
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

function ActionList_GetComponentCount(AObj: TActionList): Integer; extdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function ActionList_GetComponentIndex(AObj: TActionList): Integer; extdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure ActionList_SetComponentIndex(AObj: TActionList; AValue: Integer); extdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function ActionList_GetOwner(AObj: TActionList): TComponent; extdecl;
begin
  Result :=  AObj.Owner;
end;

function ActionList_GetName(AObj: TActionList): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure ActionList_SetName(AObj: TActionList; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function ActionList_GetTag(AObj: TActionList): NativeInt; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure ActionList_SetTag(AObj: TActionList; AValue: NativeInt); extdecl;
begin
  AObj.Tag := AValue;
end;

function ActionList_GetComponents(AObj: TActionList; AIndex: Integer): TComponent; extdecl;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  ActionList_Create,
  ActionList_Free,
  ActionList_FindComponent,
  ActionList_GetNamePath,
  ActionList_HasParent,
  ActionList_Assign,
  ActionList_ClassName,
  ActionList_Equals,
  ActionList_GetHashCode,
  ActionList_ToString,
  ActionList_GetImages,
  ActionList_SetImages,
  ActionList_GetState,
  ActionList_SetState,
  ActionList_SetOnChange,
  ActionList_GetComponentCount,
  ActionList_GetComponentIndex,
  ActionList_SetComponentIndex,
  ActionList_GetOwner,
  ActionList_GetName,
  ActionList_SetName,
  ActionList_GetTag,
  ActionList_SetTag,
  ActionList_GetComponents;

