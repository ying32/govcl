
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function SavePictureDialog_Create(AOwner: TComponent): TSavePictureDialog; stdcall;
begin
  Result :=  TSavePictureDialog.Create(AOwner);
end;

procedure SavePictureDialog_Free(AObj: TSavePictureDialog); stdcall;
begin
  AObj.Free;
end;

function SavePictureDialog_Execute(AObj: TSavePictureDialog): LongBool; stdcall;
begin
  Result :=  AObj.Execute;
end;

function SavePictureDialog_FindComponent(AObj: TSavePictureDialog; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function SavePictureDialog_GetNamePath(AObj: TSavePictureDialog): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function SavePictureDialog_HasParent(AObj: TSavePictureDialog): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure SavePictureDialog_Assign(AObj: TSavePictureDialog; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function SavePictureDialog_ClassName(AObj: TSavePictureDialog): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function SavePictureDialog_Equals(AObj: TSavePictureDialog; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function SavePictureDialog_GetHashCode(AObj: TSavePictureDialog): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function SavePictureDialog_ToString(AObj: TSavePictureDialog): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function SavePictureDialog_GetFilter(AObj: TSavePictureDialog): PChar; stdcall;
begin
  Result :=  PChar(AObj.Filter);
end;

procedure SavePictureDialog_SetFilter(AObj: TSavePictureDialog; AValue: PChar); stdcall;
begin
  AObj.Filter := AValue;
end;

function SavePictureDialog_GetFiles(AObj: TSavePictureDialog): TStrings; stdcall;
begin
  Result :=  AObj.Files;
end;

function SavePictureDialog_GetDefaultExt(AObj: TSavePictureDialog): PChar; stdcall;
begin
  Result :=  PChar(AObj.DefaultExt);
end;

procedure SavePictureDialog_SetDefaultExt(AObj: TSavePictureDialog; AValue: PChar); stdcall;
begin
  AObj.DefaultExt := AValue;
end;

function SavePictureDialog_GetFileName(AObj: TSavePictureDialog): PChar; stdcall;
begin
  Result :=  PChar(AObj.FileName);
end;

procedure SavePictureDialog_SetFileName(AObj: TSavePictureDialog; AValue: PChar); stdcall;
begin
  AObj.FileName := AValue;
end;

function SavePictureDialog_GetFilterIndex(AObj: TSavePictureDialog): Integer; stdcall;
begin
  Result :=  AObj.FilterIndex;
end;

procedure SavePictureDialog_SetFilterIndex(AObj: TSavePictureDialog; AValue: Integer); stdcall;
begin
  AObj.FilterIndex := AValue;
end;

function SavePictureDialog_GetInitialDir(AObj: TSavePictureDialog): PChar; stdcall;
begin
  Result :=  PChar(AObj.InitialDir);
end;

procedure SavePictureDialog_SetInitialDir(AObj: TSavePictureDialog; AValue: PChar); stdcall;
begin
  AObj.InitialDir := AValue;
end;

function SavePictureDialog_GetOptions(AObj: TSavePictureDialog): TOpenOptions; stdcall;
begin
  Result :=  AObj.Options;
end;

procedure SavePictureDialog_SetOptions(AObj: TSavePictureDialog; AValue: TOpenOptions); stdcall;
begin
  AObj.Options := AValue;
end;

function SavePictureDialog_GetOptionsEx(AObj: TSavePictureDialog): TOpenOptionsEx; stdcall;
begin
  Result :=  AObj.OptionsEx;
end;

procedure SavePictureDialog_SetOptionsEx(AObj: TSavePictureDialog; AValue: TOpenOptionsEx); stdcall;
begin
  AObj.OptionsEx := AValue;
end;

function SavePictureDialog_GetTitle(AObj: TSavePictureDialog): PChar; stdcall;
begin
  Result :=  PChar(AObj.Title);
end;

procedure SavePictureDialog_SetTitle(AObj: TSavePictureDialog; AValue: PChar); stdcall;
begin
  AObj.Title := AValue;
end;

function SavePictureDialog_GetHandle(AObj: TSavePictureDialog): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

procedure SavePictureDialog_SetOnClose(AObj: TSavePictureDialog; AEventId: NativeUInt); stdcall;
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

procedure SavePictureDialog_SetOnShow(AObj: TSavePictureDialog; AEventId: NativeUInt); stdcall;
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

function SavePictureDialog_GetComponentCount(AObj: TSavePictureDialog): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function SavePictureDialog_GetComponentIndex(AObj: TSavePictureDialog): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure SavePictureDialog_SetComponentIndex(AObj: TSavePictureDialog; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function SavePictureDialog_GetOwner(AObj: TSavePictureDialog): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function SavePictureDialog_GetName(AObj: TSavePictureDialog): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure SavePictureDialog_SetName(AObj: TSavePictureDialog; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function SavePictureDialog_GetTag(AObj: TSavePictureDialog): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure SavePictureDialog_SetTag(AObj: TSavePictureDialog; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function SavePictureDialog_GetComponents(AObj: TSavePictureDialog; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  SavePictureDialog_Create,
  SavePictureDialog_Free,
  SavePictureDialog_Execute,
  SavePictureDialog_FindComponent,
  SavePictureDialog_GetNamePath,
  SavePictureDialog_HasParent,
  SavePictureDialog_Assign,
  SavePictureDialog_ClassName,
  SavePictureDialog_Equals,
  SavePictureDialog_GetHashCode,
  SavePictureDialog_ToString,
  SavePictureDialog_GetFilter,
  SavePictureDialog_SetFilter,
  SavePictureDialog_GetFiles,
  SavePictureDialog_GetDefaultExt,
  SavePictureDialog_SetDefaultExt,
  SavePictureDialog_GetFileName,
  SavePictureDialog_SetFileName,
  SavePictureDialog_GetFilterIndex,
  SavePictureDialog_SetFilterIndex,
  SavePictureDialog_GetInitialDir,
  SavePictureDialog_SetInitialDir,
  SavePictureDialog_GetOptions,
  SavePictureDialog_SetOptions,
  SavePictureDialog_GetOptionsEx,
  SavePictureDialog_SetOptionsEx,
  SavePictureDialog_GetTitle,
  SavePictureDialog_SetTitle,
  SavePictureDialog_GetHandle,
  SavePictureDialog_SetOnClose,
  SavePictureDialog_SetOnShow,
  SavePictureDialog_GetComponentCount,
  SavePictureDialog_GetComponentIndex,
  SavePictureDialog_SetComponentIndex,
  SavePictureDialog_GetOwner,
  SavePictureDialog_GetName,
  SavePictureDialog_SetName,
  SavePictureDialog_GetTag,
  SavePictureDialog_SetTag,
  SavePictureDialog_GetComponents;

