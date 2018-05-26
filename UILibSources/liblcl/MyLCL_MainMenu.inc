
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function MainMenu_Create(AOwner: TComponent): TMainMenu; extdecl;
begin
  Result :=  TMainMenu.Create(AOwner);
end;

procedure MainMenu_Free(AObj: TMainMenu); extdecl;
begin
  AObj.Free;
end;

function MainMenu_FindComponent(AObj: TMainMenu; AName: PChar): TComponent; extdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function MainMenu_GetNamePath(AObj: TMainMenu): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function MainMenu_HasParent(AObj: TMainMenu): LongBool; extdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure MainMenu_Assign(AObj: TMainMenu; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function MainMenu_ClassName(AObj: TMainMenu): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function MainMenu_Equals(AObj: TMainMenu; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function MainMenu_GetHashCode(AObj: TMainMenu): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function MainMenu_ToString(AObj: TMainMenu): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function MainMenu_GetBiDiMode(AObj: TMainMenu): TBiDiMode; extdecl;
begin
  Result :=  AObj.BiDiMode;
end;

procedure MainMenu_SetBiDiMode(AObj: TMainMenu; AValue: TBiDiMode); extdecl;
begin
  AObj.BiDiMode := AValue;
end;

function MainMenu_GetImages(AObj: TMainMenu): TCustomImageList; extdecl;
begin
  Result :=  AObj.Images;
end;

procedure MainMenu_SetImages(AObj: TMainMenu; AValue: TCustomImageList); extdecl;
begin
  AObj.Images := AValue;
end;

procedure MainMenu_SetOnChange(AObj: TMainMenu; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnChange := nil;
    EventClass.Remove(AObj, geMenuChange);
    Exit;
  end;
  AObj.OnChange := @EventClass.MenuOnChange;
  EventClass.Add(AObj, geMenuChange, AEventId);
end;

function MainMenu_GetHandle(AObj: TMainMenu): HMENU; extdecl;
begin
  Result :=  AObj.Handle;
end;

function MainMenu_GetItems(AObj: TMainMenu): TMenuItem; extdecl;
begin
  Result :=  AObj.Items;
end;

function MainMenu_GetComponentCount(AObj: TMainMenu): Integer; extdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function MainMenu_GetComponentIndex(AObj: TMainMenu): Integer; extdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure MainMenu_SetComponentIndex(AObj: TMainMenu; AValue: Integer); extdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function MainMenu_GetOwner(AObj: TMainMenu): TComponent; extdecl;
begin
  Result :=  AObj.Owner;
end;

function MainMenu_GetName(AObj: TMainMenu): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure MainMenu_SetName(AObj: TMainMenu; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function MainMenu_GetTag(AObj: TMainMenu): NativeInt; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure MainMenu_SetTag(AObj: TMainMenu; AValue: NativeInt); extdecl;
begin
  AObj.Tag := AValue;
end;

function MainMenu_GetComponents(AObj: TMainMenu; AIndex: Integer): TComponent; extdecl;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  MainMenu_Create,
  MainMenu_Free,
  MainMenu_FindComponent,
  MainMenu_GetNamePath,
  MainMenu_HasParent,
  MainMenu_Assign,
  MainMenu_ClassName,
  MainMenu_Equals,
  MainMenu_GetHashCode,
  MainMenu_ToString,
  MainMenu_GetBiDiMode,
  MainMenu_SetBiDiMode,
  MainMenu_GetImages,
  MainMenu_SetImages,
  MainMenu_SetOnChange,
  MainMenu_GetHandle,
  MainMenu_GetItems,
  MainMenu_GetComponentCount,
  MainMenu_GetComponentIndex,
  MainMenu_SetComponentIndex,
  MainMenu_GetOwner,
  MainMenu_GetName,
  MainMenu_SetName,
  MainMenu_GetTag,
  MainMenu_SetTag,
  MainMenu_GetComponents;

