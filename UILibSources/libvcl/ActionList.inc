
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function ActionList_Create(AOwner: TComponent): TActionList; stdcall;
begin
  Result :=  TActionList.Create(AOwner);
end;

procedure ActionList_Free(AObj: TActionList); stdcall;
begin
  AObj.Free;
end;

function ActionList_FindComponent(AObj: TActionList; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function ActionList_GetNamePath(AObj: TActionList): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function ActionList_HasParent(AObj: TActionList): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure ActionList_Assign(AObj: TActionList; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function ActionList_ClassName(AObj: TActionList): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function ActionList_Equals(AObj: TActionList; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function ActionList_GetHashCode(AObj: TActionList): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function ActionList_ToString(AObj: TActionList): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function ActionList_GetImages(AObj: TActionList): TCustomImageList; stdcall;
begin
  Result :=  AObj.Images;
end;

procedure ActionList_SetImages(AObj: TActionList; AValue: TCustomImageList); stdcall;
begin
  AObj.Images := AValue;
end;

function ActionList_GetState(AObj: TActionList): TActionListState; stdcall;
begin
  Result :=  AObj.State;
end;

procedure ActionList_SetState(AObj: TActionList; AValue: TActionListState); stdcall;
begin
  AObj.State := AValue;
end;

procedure ActionList_SetOnChange(AObj: TActionList; AEventId: NativeUInt); stdcall;
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

function ActionList_GetComponentCount(AObj: TActionList): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function ActionList_GetComponentIndex(AObj: TActionList): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure ActionList_SetComponentIndex(AObj: TActionList; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function ActionList_GetOwner(AObj: TActionList): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function ActionList_GetName(AObj: TActionList): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure ActionList_SetName(AObj: TActionList; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function ActionList_GetTag(AObj: TActionList): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure ActionList_SetTag(AObj: TActionList; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function ActionList_GetComponents(AObj: TActionList; AIndex: Integer): TComponent; stdcall;
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

