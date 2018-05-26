
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function SaveTextFileDialog_Create(AOwner: TComponent): TSaveTextFileDialog; stdcall;
begin
  Result :=  TSaveTextFileDialog.Create(AOwner);
end;

procedure SaveTextFileDialog_Free(AObj: TSaveTextFileDialog); stdcall;
begin
  AObj.Free;
end;

function SaveTextFileDialog_Execute(AObj: TSaveTextFileDialog): LongBool; stdcall;
begin
  Result :=  AObj.Execute;
end;

function SaveTextFileDialog_FindComponent(AObj: TSaveTextFileDialog; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function SaveTextFileDialog_GetNamePath(AObj: TSaveTextFileDialog): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function SaveTextFileDialog_HasParent(AObj: TSaveTextFileDialog): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure SaveTextFileDialog_Assign(AObj: TSaveTextFileDialog; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function SaveTextFileDialog_ClassName(AObj: TSaveTextFileDialog): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function SaveTextFileDialog_Equals(AObj: TSaveTextFileDialog; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function SaveTextFileDialog_GetHashCode(AObj: TSaveTextFileDialog): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function SaveTextFileDialog_ToString(AObj: TSaveTextFileDialog): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function SaveTextFileDialog_GetFiles(AObj: TSaveTextFileDialog): TStrings; stdcall;
begin
  Result :=  AObj.Files;
end;

function SaveTextFileDialog_GetDefaultExt(AObj: TSaveTextFileDialog): PChar; stdcall;
begin
  Result :=  PChar(AObj.DefaultExt);
end;

procedure SaveTextFileDialog_SetDefaultExt(AObj: TSaveTextFileDialog; AValue: PChar); stdcall;
begin
  AObj.DefaultExt := AValue;
end;

function SaveTextFileDialog_GetFileName(AObj: TSaveTextFileDialog): PChar; stdcall;
begin
  Result :=  PChar(AObj.FileName);
end;

procedure SaveTextFileDialog_SetFileName(AObj: TSaveTextFileDialog; AValue: PChar); stdcall;
begin
  AObj.FileName := AValue;
end;

function SaveTextFileDialog_GetFilter(AObj: TSaveTextFileDialog): PChar; stdcall;
begin
  Result :=  PChar(AObj.Filter);
end;

procedure SaveTextFileDialog_SetFilter(AObj: TSaveTextFileDialog; AValue: PChar); stdcall;
begin
  AObj.Filter := AValue;
end;

function SaveTextFileDialog_GetFilterIndex(AObj: TSaveTextFileDialog): Integer; stdcall;
begin
  Result :=  AObj.FilterIndex;
end;

procedure SaveTextFileDialog_SetFilterIndex(AObj: TSaveTextFileDialog; AValue: Integer); stdcall;
begin
  AObj.FilterIndex := AValue;
end;

function SaveTextFileDialog_GetInitialDir(AObj: TSaveTextFileDialog): PChar; stdcall;
begin
  Result :=  PChar(AObj.InitialDir);
end;

procedure SaveTextFileDialog_SetInitialDir(AObj: TSaveTextFileDialog; AValue: PChar); stdcall;
begin
  AObj.InitialDir := AValue;
end;

function SaveTextFileDialog_GetOptions(AObj: TSaveTextFileDialog): TOpenOptions; stdcall;
begin
  Result :=  AObj.Options;
end;

procedure SaveTextFileDialog_SetOptions(AObj: TSaveTextFileDialog; AValue: TOpenOptions); stdcall;
begin
  AObj.Options := AValue;
end;

function SaveTextFileDialog_GetOptionsEx(AObj: TSaveTextFileDialog): TOpenOptionsEx; stdcall;
begin
  Result :=  AObj.OptionsEx;
end;

procedure SaveTextFileDialog_SetOptionsEx(AObj: TSaveTextFileDialog; AValue: TOpenOptionsEx); stdcall;
begin
  AObj.OptionsEx := AValue;
end;

function SaveTextFileDialog_GetTitle(AObj: TSaveTextFileDialog): PChar; stdcall;
begin
  Result :=  PChar(AObj.Title);
end;

procedure SaveTextFileDialog_SetTitle(AObj: TSaveTextFileDialog; AValue: PChar); stdcall;
begin
  AObj.Title := AValue;
end;

function SaveTextFileDialog_GetHandle(AObj: TSaveTextFileDialog): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

procedure SaveTextFileDialog_SetOnClose(AObj: TSaveTextFileDialog; AEventId: NativeUInt); stdcall;
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

procedure SaveTextFileDialog_SetOnShow(AObj: TSaveTextFileDialog; AEventId: NativeUInt); stdcall;
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

function SaveTextFileDialog_GetComponentCount(AObj: TSaveTextFileDialog): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function SaveTextFileDialog_GetComponentIndex(AObj: TSaveTextFileDialog): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure SaveTextFileDialog_SetComponentIndex(AObj: TSaveTextFileDialog; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function SaveTextFileDialog_GetOwner(AObj: TSaveTextFileDialog): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function SaveTextFileDialog_GetName(AObj: TSaveTextFileDialog): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure SaveTextFileDialog_SetName(AObj: TSaveTextFileDialog; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function SaveTextFileDialog_GetTag(AObj: TSaveTextFileDialog): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure SaveTextFileDialog_SetTag(AObj: TSaveTextFileDialog; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function SaveTextFileDialog_GetComponents(AObj: TSaveTextFileDialog; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  SaveTextFileDialog_Create,
  SaveTextFileDialog_Free,
  SaveTextFileDialog_Execute,
  SaveTextFileDialog_FindComponent,
  SaveTextFileDialog_GetNamePath,
  SaveTextFileDialog_HasParent,
  SaveTextFileDialog_Assign,
  SaveTextFileDialog_ClassName,
  SaveTextFileDialog_Equals,
  SaveTextFileDialog_GetHashCode,
  SaveTextFileDialog_ToString,
  SaveTextFileDialog_GetFiles,
  SaveTextFileDialog_GetDefaultExt,
  SaveTextFileDialog_SetDefaultExt,
  SaveTextFileDialog_GetFileName,
  SaveTextFileDialog_SetFileName,
  SaveTextFileDialog_GetFilter,
  SaveTextFileDialog_SetFilter,
  SaveTextFileDialog_GetFilterIndex,
  SaveTextFileDialog_SetFilterIndex,
  SaveTextFileDialog_GetInitialDir,
  SaveTextFileDialog_SetInitialDir,
  SaveTextFileDialog_GetOptions,
  SaveTextFileDialog_SetOptions,
  SaveTextFileDialog_GetOptionsEx,
  SaveTextFileDialog_SetOptionsEx,
  SaveTextFileDialog_GetTitle,
  SaveTextFileDialog_SetTitle,
  SaveTextFileDialog_GetHandle,
  SaveTextFileDialog_SetOnClose,
  SaveTextFileDialog_SetOnShow,
  SaveTextFileDialog_GetComponentCount,
  SaveTextFileDialog_GetComponentIndex,
  SaveTextFileDialog_SetComponentIndex,
  SaveTextFileDialog_GetOwner,
  SaveTextFileDialog_GetName,
  SaveTextFileDialog_SetName,
  SaveTextFileDialog_GetTag,
  SaveTextFileDialog_SetTag,
  SaveTextFileDialog_GetComponents;

