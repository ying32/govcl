
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function ColorDialog_Create(AOwner: TComponent): TColorDialog; stdcall;
begin
  Result :=  TColorDialog.Create(AOwner);
end;

procedure ColorDialog_Free(AObj: TColorDialog); stdcall;
begin
  AObj.Free;
end;

function ColorDialog_Execute(AObj: TColorDialog): LongBool; stdcall;
begin
  Result :=  AObj.Execute;
end;

function ColorDialog_FindComponent(AObj: TColorDialog; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function ColorDialog_GetNamePath(AObj: TColorDialog): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function ColorDialog_HasParent(AObj: TColorDialog): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure ColorDialog_Assign(AObj: TColorDialog; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function ColorDialog_ClassName(AObj: TColorDialog): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function ColorDialog_Equals(AObj: TColorDialog; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function ColorDialog_GetHashCode(AObj: TColorDialog): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function ColorDialog_ToString(AObj: TColorDialog): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function ColorDialog_GetColor(AObj: TColorDialog): TColor; stdcall;
begin
  Result :=  AObj.Color;
end;

procedure ColorDialog_SetColor(AObj: TColorDialog; AValue: TColor); stdcall;
begin
  AObj.Color := AValue;
end;

function ColorDialog_GetOptions(AObj: TColorDialog): TColorDialogOptions; stdcall;
begin
  Result :=  AObj.Options;
end;

procedure ColorDialog_SetOptions(AObj: TColorDialog; AValue: TColorDialogOptions); stdcall;
begin
  AObj.Options := AValue;
end;

function ColorDialog_GetHandle(AObj: TColorDialog): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

procedure ColorDialog_SetOnClose(AObj: TColorDialog; AEventId: NativeUInt); stdcall;
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

procedure ColorDialog_SetOnShow(AObj: TColorDialog; AEventId: NativeUInt); stdcall;
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

function ColorDialog_GetComponentCount(AObj: TColorDialog): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function ColorDialog_GetComponentIndex(AObj: TColorDialog): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure ColorDialog_SetComponentIndex(AObj: TColorDialog; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function ColorDialog_GetOwner(AObj: TColorDialog): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function ColorDialog_GetName(AObj: TColorDialog): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure ColorDialog_SetName(AObj: TColorDialog; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function ColorDialog_GetTag(AObj: TColorDialog): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure ColorDialog_SetTag(AObj: TColorDialog; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function ColorDialog_GetComponents(AObj: TColorDialog; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  ColorDialog_Create,
  ColorDialog_Free,
  ColorDialog_Execute,
  ColorDialog_FindComponent,
  ColorDialog_GetNamePath,
  ColorDialog_HasParent,
  ColorDialog_Assign,
  ColorDialog_ClassName,
  ColorDialog_Equals,
  ColorDialog_GetHashCode,
  ColorDialog_ToString,
  ColorDialog_GetColor,
  ColorDialog_SetColor,
  ColorDialog_GetOptions,
  ColorDialog_SetOptions,
  ColorDialog_GetHandle,
  ColorDialog_SetOnClose,
  ColorDialog_SetOnShow,
  ColorDialog_GetComponentCount,
  ColorDialog_GetComponentIndex,
  ColorDialog_SetComponentIndex,
  ColorDialog_GetOwner,
  ColorDialog_GetName,
  ColorDialog_SetName,
  ColorDialog_GetTag,
  ColorDialog_SetTag,
  ColorDialog_GetComponents;

