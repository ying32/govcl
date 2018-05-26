
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function OpenDialog_Create(AOwner: TComponent): TOpenDialog; extdecl;
begin
  Result :=  TOpenDialog.Create(AOwner);
end;

procedure OpenDialog_Free(AObj: TOpenDialog); extdecl;
begin
  AObj.Free;
end;

function OpenDialog_Execute(AObj: TOpenDialog): LongBool; extdecl;
begin
  Result :=  AObj.Execute;
end;

function OpenDialog_FindComponent(AObj: TOpenDialog; AName: PChar): TComponent; extdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function OpenDialog_GetNamePath(AObj: TOpenDialog): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function OpenDialog_HasParent(AObj: TOpenDialog): LongBool; extdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure OpenDialog_Assign(AObj: TOpenDialog; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function OpenDialog_ClassName(AObj: TOpenDialog): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function OpenDialog_Equals(AObj: TOpenDialog; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function OpenDialog_GetHashCode(AObj: TOpenDialog): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function OpenDialog_ToString(AObj: TOpenDialog): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function OpenDialog_GetFiles(AObj: TOpenDialog): TStrings; extdecl;
begin
  Result :=  AObj.Files;
end;

function OpenDialog_GetDefaultExt(AObj: TOpenDialog): PChar; extdecl;
begin
  Result :=  PChar(AObj.DefaultExt);
end;

procedure OpenDialog_SetDefaultExt(AObj: TOpenDialog; AValue: PChar); extdecl;
begin
  AObj.DefaultExt := AValue;
end;

function OpenDialog_GetFileName(AObj: TOpenDialog): PChar; extdecl;
begin
  Result :=  PChar(AObj.FileName);
end;

procedure OpenDialog_SetFileName(AObj: TOpenDialog; AValue: PChar); extdecl;
begin
  AObj.FileName := AValue;
end;

function OpenDialog_GetFilter(AObj: TOpenDialog): PChar; extdecl;
begin
  Result :=  PChar(AObj.Filter);
end;

procedure OpenDialog_SetFilter(AObj: TOpenDialog; AValue: PChar); extdecl;
begin
  AObj.Filter := AValue;
end;

function OpenDialog_GetFilterIndex(AObj: TOpenDialog): Integer; extdecl;
begin
  Result :=  AObj.FilterIndex;
end;

procedure OpenDialog_SetFilterIndex(AObj: TOpenDialog; AValue: Integer); extdecl;
begin
  AObj.FilterIndex := AValue;
end;

function OpenDialog_GetInitialDir(AObj: TOpenDialog): PChar; extdecl;
begin
  Result :=  PChar(AObj.InitialDir);
end;

procedure OpenDialog_SetInitialDir(AObj: TOpenDialog; AValue: PChar); extdecl;
begin
  AObj.InitialDir := AValue;
end;

function OpenDialog_GetOptions(AObj: TOpenDialog): TOpenOptions; extdecl;
begin
  Result :=  AObj.Options;
end;

procedure OpenDialog_SetOptions(AObj: TOpenDialog; AValue: TOpenOptions); extdecl;
begin
  AObj.Options := AValue;
end;

function OpenDialog_GetTitle(AObj: TOpenDialog): PChar; extdecl;
begin
  Result :=  PChar(AObj.Title);
end;

procedure OpenDialog_SetTitle(AObj: TOpenDialog; AValue: PChar); extdecl;
begin
  AObj.Title := AValue;
end;

function OpenDialog_GetHandle(AObj: TOpenDialog): HWND; extdecl;
begin
  Result :=  AObj.Handle;
end;

procedure OpenDialog_SetOnClose(AObj: TOpenDialog; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnClose := nil;
    EventClass.Remove(AObj, geClose);
    Exit;
  end;
  AObj.OnClose := @EventClass.OnClose;
  EventClass.Add(AObj, geClose, AEventId);
end;

procedure OpenDialog_SetOnShow(AObj: TOpenDialog; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnShow := nil;
    EventClass.Remove(AObj, geShow);
    Exit;
  end;
  AObj.OnShow := @EventClass.OnShow;
  EventClass.Add(AObj, geShow, AEventId);
end;

function OpenDialog_GetComponentCount(AObj: TOpenDialog): Integer; extdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function OpenDialog_GetComponentIndex(AObj: TOpenDialog): Integer; extdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure OpenDialog_SetComponentIndex(AObj: TOpenDialog; AValue: Integer); extdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function OpenDialog_GetOwner(AObj: TOpenDialog): TComponent; extdecl;
begin
  Result :=  AObj.Owner;
end;

function OpenDialog_GetName(AObj: TOpenDialog): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure OpenDialog_SetName(AObj: TOpenDialog; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function OpenDialog_GetTag(AObj: TOpenDialog): NativeInt; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure OpenDialog_SetTag(AObj: TOpenDialog; AValue: NativeInt); extdecl;
begin
  AObj.Tag := AValue;
end;

function OpenDialog_GetComponents(AObj: TOpenDialog; AIndex: Integer): TComponent; extdecl;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  OpenDialog_Create,
  OpenDialog_Free,
  OpenDialog_Execute,
  OpenDialog_FindComponent,
  OpenDialog_GetNamePath,
  OpenDialog_HasParent,
  OpenDialog_Assign,
  OpenDialog_ClassName,
  OpenDialog_Equals,
  OpenDialog_GetHashCode,
  OpenDialog_ToString,
  OpenDialog_GetFiles,
  OpenDialog_GetDefaultExt,
  OpenDialog_SetDefaultExt,
  OpenDialog_GetFileName,
  OpenDialog_SetFileName,
  OpenDialog_GetFilter,
  OpenDialog_SetFilter,
  OpenDialog_GetFilterIndex,
  OpenDialog_SetFilterIndex,
  OpenDialog_GetInitialDir,
  OpenDialog_SetInitialDir,
  OpenDialog_GetOptions,
  OpenDialog_SetOptions,
  OpenDialog_GetTitle,
  OpenDialog_SetTitle,
  OpenDialog_GetHandle,
  OpenDialog_SetOnClose,
  OpenDialog_SetOnShow,
  OpenDialog_GetComponentCount,
  OpenDialog_GetComponentIndex,
  OpenDialog_SetComponentIndex,
  OpenDialog_GetOwner,
  OpenDialog_GetName,
  OpenDialog_SetName,
  OpenDialog_GetTag,
  OpenDialog_SetTag,
  OpenDialog_GetComponents;

