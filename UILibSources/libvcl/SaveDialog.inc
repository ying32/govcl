
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function SaveDialog_Create(AOwner: TComponent): TSaveDialog; stdcall;
begin
  Result :=  TSaveDialog.Create(AOwner);
end;

procedure SaveDialog_Free(AObj: TSaveDialog); stdcall;
begin
  AObj.Free;
end;

function SaveDialog_Execute(AObj: TSaveDialog): LongBool; stdcall;
begin
  Result :=  AObj.Execute;
end;

function SaveDialog_FindComponent(AObj: TSaveDialog; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function SaveDialog_GetNamePath(AObj: TSaveDialog): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function SaveDialog_HasParent(AObj: TSaveDialog): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure SaveDialog_Assign(AObj: TSaveDialog; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function SaveDialog_ClassName(AObj: TSaveDialog): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function SaveDialog_Equals(AObj: TSaveDialog; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function SaveDialog_GetHashCode(AObj: TSaveDialog): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function SaveDialog_ToString(AObj: TSaveDialog): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function SaveDialog_GetFiles(AObj: TSaveDialog): TStrings; stdcall;
begin
  Result :=  AObj.Files;
end;

function SaveDialog_GetDefaultExt(AObj: TSaveDialog): PChar; stdcall;
begin
  Result :=  PChar(AObj.DefaultExt);
end;

procedure SaveDialog_SetDefaultExt(AObj: TSaveDialog; AValue: PChar); stdcall;
begin
  AObj.DefaultExt := AValue;
end;

function SaveDialog_GetFileName(AObj: TSaveDialog): PChar; stdcall;
begin
  Result :=  PChar(AObj.FileName);
end;

procedure SaveDialog_SetFileName(AObj: TSaveDialog; AValue: PChar); stdcall;
begin
  AObj.FileName := AValue;
end;

function SaveDialog_GetFilter(AObj: TSaveDialog): PChar; stdcall;
begin
  Result :=  PChar(AObj.Filter);
end;

procedure SaveDialog_SetFilter(AObj: TSaveDialog; AValue: PChar); stdcall;
begin
  AObj.Filter := AValue;
end;

function SaveDialog_GetFilterIndex(AObj: TSaveDialog): Integer; stdcall;
begin
  Result :=  AObj.FilterIndex;
end;

procedure SaveDialog_SetFilterIndex(AObj: TSaveDialog; AValue: Integer); stdcall;
begin
  AObj.FilterIndex := AValue;
end;

function SaveDialog_GetInitialDir(AObj: TSaveDialog): PChar; stdcall;
begin
  Result :=  PChar(AObj.InitialDir);
end;

procedure SaveDialog_SetInitialDir(AObj: TSaveDialog; AValue: PChar); stdcall;
begin
  AObj.InitialDir := AValue;
end;

function SaveDialog_GetOptions(AObj: TSaveDialog): TOpenOptions; stdcall;
begin
  Result :=  AObj.Options;
end;

procedure SaveDialog_SetOptions(AObj: TSaveDialog; AValue: TOpenOptions); stdcall;
begin
  AObj.Options := AValue;
end;

function SaveDialog_GetOptionsEx(AObj: TSaveDialog): TOpenOptionsEx; stdcall;
begin
  Result :=  AObj.OptionsEx;
end;

procedure SaveDialog_SetOptionsEx(AObj: TSaveDialog; AValue: TOpenOptionsEx); stdcall;
begin
  AObj.OptionsEx := AValue;
end;

function SaveDialog_GetTitle(AObj: TSaveDialog): PChar; stdcall;
begin
  Result :=  PChar(AObj.Title);
end;

procedure SaveDialog_SetTitle(AObj: TSaveDialog; AValue: PChar); stdcall;
begin
  AObj.Title := AValue;
end;

function SaveDialog_GetHandle(AObj: TSaveDialog): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

procedure SaveDialog_SetOnClose(AObj: TSaveDialog; AEventId: NativeUInt); stdcall;
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

procedure SaveDialog_SetOnShow(AObj: TSaveDialog; AEventId: NativeUInt); stdcall;
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

function SaveDialog_GetComponentCount(AObj: TSaveDialog): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function SaveDialog_GetComponentIndex(AObj: TSaveDialog): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure SaveDialog_SetComponentIndex(AObj: TSaveDialog; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function SaveDialog_GetOwner(AObj: TSaveDialog): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function SaveDialog_GetName(AObj: TSaveDialog): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure SaveDialog_SetName(AObj: TSaveDialog; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function SaveDialog_GetTag(AObj: TSaveDialog): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure SaveDialog_SetTag(AObj: TSaveDialog; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function SaveDialog_GetComponents(AObj: TSaveDialog; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  SaveDialog_Create,
  SaveDialog_Free,
  SaveDialog_Execute,
  SaveDialog_FindComponent,
  SaveDialog_GetNamePath,
  SaveDialog_HasParent,
  SaveDialog_Assign,
  SaveDialog_ClassName,
  SaveDialog_Equals,
  SaveDialog_GetHashCode,
  SaveDialog_ToString,
  SaveDialog_GetFiles,
  SaveDialog_GetDefaultExt,
  SaveDialog_SetDefaultExt,
  SaveDialog_GetFileName,
  SaveDialog_SetFileName,
  SaveDialog_GetFilter,
  SaveDialog_SetFilter,
  SaveDialog_GetFilterIndex,
  SaveDialog_SetFilterIndex,
  SaveDialog_GetInitialDir,
  SaveDialog_SetInitialDir,
  SaveDialog_GetOptions,
  SaveDialog_SetOptions,
  SaveDialog_GetOptionsEx,
  SaveDialog_SetOptionsEx,
  SaveDialog_GetTitle,
  SaveDialog_SetTitle,
  SaveDialog_GetHandle,
  SaveDialog_SetOnClose,
  SaveDialog_SetOnShow,
  SaveDialog_GetComponentCount,
  SaveDialog_GetComponentIndex,
  SaveDialog_SetComponentIndex,
  SaveDialog_GetOwner,
  SaveDialog_GetName,
  SaveDialog_SetName,
  SaveDialog_GetTag,
  SaveDialog_SetTag,
  SaveDialog_GetComponents;

