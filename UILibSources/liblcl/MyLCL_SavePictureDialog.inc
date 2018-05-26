
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function SavePictureDialog_Create(AOwner: TComponent): TSavePictureDialog; extdecl;
begin
  Result :=  TSavePictureDialog.Create(AOwner);
end;

procedure SavePictureDialog_Free(AObj: TSavePictureDialog); extdecl;
begin
  AObj.Free;
end;

function SavePictureDialog_Execute(AObj: TSavePictureDialog): LongBool; extdecl;
begin
  Result :=  AObj.Execute;
end;

function SavePictureDialog_FindComponent(AObj: TSavePictureDialog; AName: PChar): TComponent; extdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function SavePictureDialog_GetNamePath(AObj: TSavePictureDialog): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function SavePictureDialog_HasParent(AObj: TSavePictureDialog): LongBool; extdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure SavePictureDialog_Assign(AObj: TSavePictureDialog; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function SavePictureDialog_ClassName(AObj: TSavePictureDialog): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function SavePictureDialog_Equals(AObj: TSavePictureDialog; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function SavePictureDialog_GetHashCode(AObj: TSavePictureDialog): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function SavePictureDialog_ToString(AObj: TSavePictureDialog): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function SavePictureDialog_GetFilter(AObj: TSavePictureDialog): PChar; extdecl;
begin
  Result :=  PChar(AObj.Filter);
end;

procedure SavePictureDialog_SetFilter(AObj: TSavePictureDialog; AValue: PChar); extdecl;
begin
  AObj.Filter := AValue;
end;

function SavePictureDialog_GetFiles(AObj: TSavePictureDialog): TStrings; extdecl;
begin
  Result :=  AObj.Files;
end;

function SavePictureDialog_GetDefaultExt(AObj: TSavePictureDialog): PChar; extdecl;
begin
  Result :=  PChar(AObj.DefaultExt);
end;

procedure SavePictureDialog_SetDefaultExt(AObj: TSavePictureDialog; AValue: PChar); extdecl;
begin
  AObj.DefaultExt := AValue;
end;

function SavePictureDialog_GetFileName(AObj: TSavePictureDialog): PChar; extdecl;
begin
  Result :=  PChar(AObj.FileName);
end;

procedure SavePictureDialog_SetFileName(AObj: TSavePictureDialog; AValue: PChar); extdecl;
begin
  AObj.FileName := AValue;
end;

function SavePictureDialog_GetFilterIndex(AObj: TSavePictureDialog): Integer; extdecl;
begin
  Result :=  AObj.FilterIndex;
end;

procedure SavePictureDialog_SetFilterIndex(AObj: TSavePictureDialog; AValue: Integer); extdecl;
begin
  AObj.FilterIndex := AValue;
end;

function SavePictureDialog_GetInitialDir(AObj: TSavePictureDialog): PChar; extdecl;
begin
  Result :=  PChar(AObj.InitialDir);
end;

procedure SavePictureDialog_SetInitialDir(AObj: TSavePictureDialog; AValue: PChar); extdecl;
begin
  AObj.InitialDir := AValue;
end;

function SavePictureDialog_GetOptions(AObj: TSavePictureDialog): TOpenOptions; extdecl;
begin
  Result :=  AObj.Options;
end;

procedure SavePictureDialog_SetOptions(AObj: TSavePictureDialog; AValue: TOpenOptions); extdecl;
begin
  AObj.Options := AValue;
end;

function SavePictureDialog_GetTitle(AObj: TSavePictureDialog): PChar; extdecl;
begin
  Result :=  PChar(AObj.Title);
end;

procedure SavePictureDialog_SetTitle(AObj: TSavePictureDialog; AValue: PChar); extdecl;
begin
  AObj.Title := AValue;
end;

function SavePictureDialog_GetHandle(AObj: TSavePictureDialog): HWND; extdecl;
begin
  Result :=  AObj.Handle;
end;

procedure SavePictureDialog_SetOnClose(AObj: TSavePictureDialog; AEventId: NativeUInt); extdecl;
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

procedure SavePictureDialog_SetOnShow(AObj: TSavePictureDialog; AEventId: NativeUInt); extdecl;
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

function SavePictureDialog_GetComponentCount(AObj: TSavePictureDialog): Integer; extdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function SavePictureDialog_GetComponentIndex(AObj: TSavePictureDialog): Integer; extdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure SavePictureDialog_SetComponentIndex(AObj: TSavePictureDialog; AValue: Integer); extdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function SavePictureDialog_GetOwner(AObj: TSavePictureDialog): TComponent; extdecl;
begin
  Result :=  AObj.Owner;
end;

function SavePictureDialog_GetName(AObj: TSavePictureDialog): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure SavePictureDialog_SetName(AObj: TSavePictureDialog; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function SavePictureDialog_GetTag(AObj: TSavePictureDialog): NativeInt; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure SavePictureDialog_SetTag(AObj: TSavePictureDialog; AValue: NativeInt); extdecl;
begin
  AObj.Tag := AValue;
end;

function SavePictureDialog_GetComponents(AObj: TSavePictureDialog; AIndex: Integer): TComponent; extdecl;
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

