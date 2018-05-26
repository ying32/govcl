
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function OpenDialog_Create(AOwner: TComponent): TOpenDialog; stdcall;
begin
  Result :=  TOpenDialog.Create(AOwner);
end;

procedure OpenDialog_Free(AObj: TOpenDialog); stdcall;
begin
  AObj.Free;
end;

function OpenDialog_Execute(AObj: TOpenDialog): LongBool; stdcall;
begin
  Result :=  AObj.Execute;
end;

function OpenDialog_FindComponent(AObj: TOpenDialog; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function OpenDialog_GetNamePath(AObj: TOpenDialog): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function OpenDialog_HasParent(AObj: TOpenDialog): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure OpenDialog_Assign(AObj: TOpenDialog; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function OpenDialog_ClassName(AObj: TOpenDialog): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function OpenDialog_Equals(AObj: TOpenDialog; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function OpenDialog_GetHashCode(AObj: TOpenDialog): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function OpenDialog_ToString(AObj: TOpenDialog): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function OpenDialog_GetFiles(AObj: TOpenDialog): TStrings; stdcall;
begin
  Result :=  AObj.Files;
end;

function OpenDialog_GetDefaultExt(AObj: TOpenDialog): PChar; stdcall;
begin
  Result :=  PChar(AObj.DefaultExt);
end;

procedure OpenDialog_SetDefaultExt(AObj: TOpenDialog; AValue: PChar); stdcall;
begin
  AObj.DefaultExt := AValue;
end;

function OpenDialog_GetFileName(AObj: TOpenDialog): PChar; stdcall;
begin
  Result :=  PChar(AObj.FileName);
end;

procedure OpenDialog_SetFileName(AObj: TOpenDialog; AValue: PChar); stdcall;
begin
  AObj.FileName := AValue;
end;

function OpenDialog_GetFilter(AObj: TOpenDialog): PChar; stdcall;
begin
  Result :=  PChar(AObj.Filter);
end;

procedure OpenDialog_SetFilter(AObj: TOpenDialog; AValue: PChar); stdcall;
begin
  AObj.Filter := AValue;
end;

function OpenDialog_GetFilterIndex(AObj: TOpenDialog): Integer; stdcall;
begin
  Result :=  AObj.FilterIndex;
end;

procedure OpenDialog_SetFilterIndex(AObj: TOpenDialog; AValue: Integer); stdcall;
begin
  AObj.FilterIndex := AValue;
end;

function OpenDialog_GetInitialDir(AObj: TOpenDialog): PChar; stdcall;
begin
  Result :=  PChar(AObj.InitialDir);
end;

procedure OpenDialog_SetInitialDir(AObj: TOpenDialog; AValue: PChar); stdcall;
begin
  AObj.InitialDir := AValue;
end;

function OpenDialog_GetOptions(AObj: TOpenDialog): TOpenOptions; stdcall;
begin
  Result :=  AObj.Options;
end;

procedure OpenDialog_SetOptions(AObj: TOpenDialog; AValue: TOpenOptions); stdcall;
begin
  AObj.Options := AValue;
end;

function OpenDialog_GetOptionsEx(AObj: TOpenDialog): TOpenOptionsEx; stdcall;
begin
  Result :=  AObj.OptionsEx;
end;

procedure OpenDialog_SetOptionsEx(AObj: TOpenDialog; AValue: TOpenOptionsEx); stdcall;
begin
  AObj.OptionsEx := AValue;
end;

function OpenDialog_GetTitle(AObj: TOpenDialog): PChar; stdcall;
begin
  Result :=  PChar(AObj.Title);
end;

procedure OpenDialog_SetTitle(AObj: TOpenDialog; AValue: PChar); stdcall;
begin
  AObj.Title := AValue;
end;

function OpenDialog_GetHandle(AObj: TOpenDialog): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

procedure OpenDialog_SetOnClose(AObj: TOpenDialog; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnClose := nil;
    TEventClass.Remove(AObj, geClose);
    Exit;
  end;
  AObj.OnClose := TEventClass.OnClose;
  TEventClass.Add(AObj, geClose, AEventId);
end;

procedure OpenDialog_SetOnShow(AObj: TOpenDialog; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnShow := nil;
    TEventClass.Remove(AObj, geShow);
    Exit;
  end;
  AObj.OnShow := TEventClass.OnShow;
  TEventClass.Add(AObj, geShow, AEventId);
end;

function OpenDialog_GetComponentCount(AObj: TOpenDialog): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function OpenDialog_GetComponentIndex(AObj: TOpenDialog): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure OpenDialog_SetComponentIndex(AObj: TOpenDialog; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function OpenDialog_GetOwner(AObj: TOpenDialog): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function OpenDialog_GetName(AObj: TOpenDialog): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure OpenDialog_SetName(AObj: TOpenDialog; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function OpenDialog_GetTag(AObj: TOpenDialog): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure OpenDialog_SetTag(AObj: TOpenDialog; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function OpenDialog_GetComponents(AObj: TOpenDialog; AIndex: Integer): TComponent; stdcall;
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
  OpenDialog_GetOptionsEx,
  OpenDialog_SetOptionsEx,
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

