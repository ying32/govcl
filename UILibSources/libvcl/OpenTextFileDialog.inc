
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function OpenTextFileDialog_Create(AOwner: TComponent): TOpenTextFileDialog; stdcall;
begin
  Result :=  TOpenTextFileDialog.Create(AOwner);
end;

procedure OpenTextFileDialog_Free(AObj: TOpenTextFileDialog); stdcall;
begin
  AObj.Free;
end;

function OpenTextFileDialog_Execute(AObj: TOpenTextFileDialog): LongBool; stdcall;
begin
  Result :=  AObj.Execute;
end;

function OpenTextFileDialog_FindComponent(AObj: TOpenTextFileDialog; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function OpenTextFileDialog_GetNamePath(AObj: TOpenTextFileDialog): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function OpenTextFileDialog_HasParent(AObj: TOpenTextFileDialog): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure OpenTextFileDialog_Assign(AObj: TOpenTextFileDialog; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function OpenTextFileDialog_ClassName(AObj: TOpenTextFileDialog): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function OpenTextFileDialog_Equals(AObj: TOpenTextFileDialog; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function OpenTextFileDialog_GetHashCode(AObj: TOpenTextFileDialog): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function OpenTextFileDialog_ToString(AObj: TOpenTextFileDialog): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function OpenTextFileDialog_GetFiles(AObj: TOpenTextFileDialog): TStrings; stdcall;
begin
  Result :=  AObj.Files;
end;

function OpenTextFileDialog_GetDefaultExt(AObj: TOpenTextFileDialog): PChar; stdcall;
begin
  Result :=  PChar(AObj.DefaultExt);
end;

procedure OpenTextFileDialog_SetDefaultExt(AObj: TOpenTextFileDialog; AValue: PChar); stdcall;
begin
  AObj.DefaultExt := AValue;
end;

function OpenTextFileDialog_GetFileName(AObj: TOpenTextFileDialog): PChar; stdcall;
begin
  Result :=  PChar(AObj.FileName);
end;

procedure OpenTextFileDialog_SetFileName(AObj: TOpenTextFileDialog; AValue: PChar); stdcall;
begin
  AObj.FileName := AValue;
end;

function OpenTextFileDialog_GetFilter(AObj: TOpenTextFileDialog): PChar; stdcall;
begin
  Result :=  PChar(AObj.Filter);
end;

procedure OpenTextFileDialog_SetFilter(AObj: TOpenTextFileDialog; AValue: PChar); stdcall;
begin
  AObj.Filter := AValue;
end;

function OpenTextFileDialog_GetFilterIndex(AObj: TOpenTextFileDialog): Integer; stdcall;
begin
  Result :=  AObj.FilterIndex;
end;

procedure OpenTextFileDialog_SetFilterIndex(AObj: TOpenTextFileDialog; AValue: Integer); stdcall;
begin
  AObj.FilterIndex := AValue;
end;

function OpenTextFileDialog_GetInitialDir(AObj: TOpenTextFileDialog): PChar; stdcall;
begin
  Result :=  PChar(AObj.InitialDir);
end;

procedure OpenTextFileDialog_SetInitialDir(AObj: TOpenTextFileDialog; AValue: PChar); stdcall;
begin
  AObj.InitialDir := AValue;
end;

function OpenTextFileDialog_GetOptions(AObj: TOpenTextFileDialog): TOpenOptions; stdcall;
begin
  Result :=  AObj.Options;
end;

procedure OpenTextFileDialog_SetOptions(AObj: TOpenTextFileDialog; AValue: TOpenOptions); stdcall;
begin
  AObj.Options := AValue;
end;

function OpenTextFileDialog_GetOptionsEx(AObj: TOpenTextFileDialog): TOpenOptionsEx; stdcall;
begin
  Result :=  AObj.OptionsEx;
end;

procedure OpenTextFileDialog_SetOptionsEx(AObj: TOpenTextFileDialog; AValue: TOpenOptionsEx); stdcall;
begin
  AObj.OptionsEx := AValue;
end;

function OpenTextFileDialog_GetTitle(AObj: TOpenTextFileDialog): PChar; stdcall;
begin
  Result :=  PChar(AObj.Title);
end;

procedure OpenTextFileDialog_SetTitle(AObj: TOpenTextFileDialog; AValue: PChar); stdcall;
begin
  AObj.Title := AValue;
end;

function OpenTextFileDialog_GetHandle(AObj: TOpenTextFileDialog): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

procedure OpenTextFileDialog_SetOnClose(AObj: TOpenTextFileDialog; AEventId: NativeUInt); stdcall;
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

procedure OpenTextFileDialog_SetOnShow(AObj: TOpenTextFileDialog; AEventId: NativeUInt); stdcall;
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

function OpenTextFileDialog_GetComponentCount(AObj: TOpenTextFileDialog): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function OpenTextFileDialog_GetComponentIndex(AObj: TOpenTextFileDialog): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure OpenTextFileDialog_SetComponentIndex(AObj: TOpenTextFileDialog; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function OpenTextFileDialog_GetOwner(AObj: TOpenTextFileDialog): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function OpenTextFileDialog_GetName(AObj: TOpenTextFileDialog): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure OpenTextFileDialog_SetName(AObj: TOpenTextFileDialog; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function OpenTextFileDialog_GetTag(AObj: TOpenTextFileDialog): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure OpenTextFileDialog_SetTag(AObj: TOpenTextFileDialog; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function OpenTextFileDialog_GetComponents(AObj: TOpenTextFileDialog; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  OpenTextFileDialog_Create,
  OpenTextFileDialog_Free,
  OpenTextFileDialog_Execute,
  OpenTextFileDialog_FindComponent,
  OpenTextFileDialog_GetNamePath,
  OpenTextFileDialog_HasParent,
  OpenTextFileDialog_Assign,
  OpenTextFileDialog_ClassName,
  OpenTextFileDialog_Equals,
  OpenTextFileDialog_GetHashCode,
  OpenTextFileDialog_ToString,
  OpenTextFileDialog_GetFiles,
  OpenTextFileDialog_GetDefaultExt,
  OpenTextFileDialog_SetDefaultExt,
  OpenTextFileDialog_GetFileName,
  OpenTextFileDialog_SetFileName,
  OpenTextFileDialog_GetFilter,
  OpenTextFileDialog_SetFilter,
  OpenTextFileDialog_GetFilterIndex,
  OpenTextFileDialog_SetFilterIndex,
  OpenTextFileDialog_GetInitialDir,
  OpenTextFileDialog_SetInitialDir,
  OpenTextFileDialog_GetOptions,
  OpenTextFileDialog_SetOptions,
  OpenTextFileDialog_GetOptionsEx,
  OpenTextFileDialog_SetOptionsEx,
  OpenTextFileDialog_GetTitle,
  OpenTextFileDialog_SetTitle,
  OpenTextFileDialog_GetHandle,
  OpenTextFileDialog_SetOnClose,
  OpenTextFileDialog_SetOnShow,
  OpenTextFileDialog_GetComponentCount,
  OpenTextFileDialog_GetComponentIndex,
  OpenTextFileDialog_SetComponentIndex,
  OpenTextFileDialog_GetOwner,
  OpenTextFileDialog_GetName,
  OpenTextFileDialog_SetName,
  OpenTextFileDialog_GetTag,
  OpenTextFileDialog_SetTag,
  OpenTextFileDialog_GetComponents;

