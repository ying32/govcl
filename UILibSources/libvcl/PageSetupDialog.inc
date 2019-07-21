
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function PageSetupDialog_Create(AOwner: TComponent): TPageSetupDialog; stdcall;
begin
  Result :=  TPageSetupDialog.Create(AOwner);
end;

procedure PageSetupDialog_Free(AObj: TPageSetupDialog); stdcall;
begin
  AObj.Free;
end;

function PageSetupDialog_GetDefaults(AObj: TPageSetupDialog): LongBool; stdcall;
begin
  Result :=  AObj.GetDefaults;
end;

function PageSetupDialog_Execute(AObj: TPageSetupDialog): LongBool; stdcall;
begin
  Result :=  AObj.Execute;
end;

function PageSetupDialog_FindComponent(AObj: TPageSetupDialog; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function PageSetupDialog_GetNamePath(AObj: TPageSetupDialog): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function PageSetupDialog_HasParent(AObj: TPageSetupDialog): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure PageSetupDialog_Assign(AObj: TPageSetupDialog; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

procedure PageSetupDialog_DisposeOf(AObj: TPageSetupDialog); stdcall;
begin
  AObj.DisposeOf;
end;

function PageSetupDialog_ClassType(AObj: TPageSetupDialog): TClass; stdcall;
begin
  Result :=  AObj.ClassType;
end;

function PageSetupDialog_ClassName(AObj: TPageSetupDialog): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function PageSetupDialog_InstanceSize(AObj: TPageSetupDialog): Integer; stdcall;
begin
  Result :=  AObj.InstanceSize;
end;

function PageSetupDialog_InheritsFrom(AObj: TPageSetupDialog; AClass: TClass): LongBool; stdcall;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function PageSetupDialog_Equals(AObj: TPageSetupDialog; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function PageSetupDialog_GetHashCode(AObj: TPageSetupDialog): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function PageSetupDialog_ToString(AObj: TPageSetupDialog): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function PageSetupDialog_GetMinMarginLeft(AObj: TPageSetupDialog): Integer; stdcall;
begin
  Result :=  AObj.MinMarginLeft;
end;

procedure PageSetupDialog_SetMinMarginLeft(AObj: TPageSetupDialog; AValue: Integer); stdcall;
begin
  AObj.MinMarginLeft := AValue;
end;

function PageSetupDialog_GetMinMarginTop(AObj: TPageSetupDialog): Integer; stdcall;
begin
  Result :=  AObj.MinMarginTop;
end;

procedure PageSetupDialog_SetMinMarginTop(AObj: TPageSetupDialog; AValue: Integer); stdcall;
begin
  AObj.MinMarginTop := AValue;
end;

function PageSetupDialog_GetMinMarginRight(AObj: TPageSetupDialog): Integer; stdcall;
begin
  Result :=  AObj.MinMarginRight;
end;

procedure PageSetupDialog_SetMinMarginRight(AObj: TPageSetupDialog; AValue: Integer); stdcall;
begin
  AObj.MinMarginRight := AValue;
end;

function PageSetupDialog_GetMinMarginBottom(AObj: TPageSetupDialog): Integer; stdcall;
begin
  Result :=  AObj.MinMarginBottom;
end;

procedure PageSetupDialog_SetMinMarginBottom(AObj: TPageSetupDialog; AValue: Integer); stdcall;
begin
  AObj.MinMarginBottom := AValue;
end;

function PageSetupDialog_GetMarginLeft(AObj: TPageSetupDialog): Integer; stdcall;
begin
  Result :=  AObj.MarginLeft;
end;

procedure PageSetupDialog_SetMarginLeft(AObj: TPageSetupDialog; AValue: Integer); stdcall;
begin
  AObj.MarginLeft := AValue;
end;

function PageSetupDialog_GetMarginTop(AObj: TPageSetupDialog): Integer; stdcall;
begin
  Result :=  AObj.MarginTop;
end;

procedure PageSetupDialog_SetMarginTop(AObj: TPageSetupDialog; AValue: Integer); stdcall;
begin
  AObj.MarginTop := AValue;
end;

function PageSetupDialog_GetMarginRight(AObj: TPageSetupDialog): Integer; stdcall;
begin
  Result :=  AObj.MarginRight;
end;

procedure PageSetupDialog_SetMarginRight(AObj: TPageSetupDialog; AValue: Integer); stdcall;
begin
  AObj.MarginRight := AValue;
end;

function PageSetupDialog_GetMarginBottom(AObj: TPageSetupDialog): Integer; stdcall;
begin
  Result :=  AObj.MarginBottom;
end;

procedure PageSetupDialog_SetMarginBottom(AObj: TPageSetupDialog; AValue: Integer); stdcall;
begin
  AObj.MarginBottom := AValue;
end;

function PageSetupDialog_GetOptions(AObj: TPageSetupDialog): TPageSetupDialogOptions; stdcall;
begin
  Result :=  AObj.Options;
end;

procedure PageSetupDialog_SetOptions(AObj: TPageSetupDialog; AValue: TPageSetupDialogOptions); stdcall;
begin
  AObj.Options := AValue;
end;

function PageSetupDialog_GetPageWidth(AObj: TPageSetupDialog): Integer; stdcall;
begin
  Result :=  AObj.PageWidth;
end;

procedure PageSetupDialog_SetPageWidth(AObj: TPageSetupDialog; AValue: Integer); stdcall;
begin
  AObj.PageWidth := AValue;
end;

function PageSetupDialog_GetPageHeight(AObj: TPageSetupDialog): Integer; stdcall;
begin
  Result :=  AObj.PageHeight;
end;

procedure PageSetupDialog_SetPageHeight(AObj: TPageSetupDialog; AValue: Integer); stdcall;
begin
  AObj.PageHeight := AValue;
end;

function PageSetupDialog_GetUnits(AObj: TPageSetupDialog): TPageMeasureUnits; stdcall;
begin
  Result :=  AObj.Units;
end;

procedure PageSetupDialog_SetUnits(AObj: TPageSetupDialog; AValue: TPageMeasureUnits); stdcall;
begin
  AObj.Units := AValue;
end;

function PageSetupDialog_GetHandle(AObj: TPageSetupDialog): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

function PageSetupDialog_GetCtl3D(AObj: TPageSetupDialog): LongBool; stdcall;
begin
  Result :=  AObj.Ctl3D;
end;

procedure PageSetupDialog_SetCtl3D(AObj: TPageSetupDialog; AValue: LongBool); stdcall;
begin
  AObj.Ctl3D := AValue;
end;

procedure PageSetupDialog_SetOnClose(AObj: TPageSetupDialog; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnClose := nil;
    TEventClass.Remove(AObj, @TEventClass.OnClose);
    Exit;
  end;
  AObj.OnClose := TEventClass.OnClose;
  TEventClass.Add(AObj, @TEventClass.OnClose, AEventId);
end;

procedure PageSetupDialog_SetOnShow(AObj: TPageSetupDialog; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnShow := nil;
    TEventClass.Remove(AObj, @TEventClass.OnShow);
    Exit;
  end;
  AObj.OnShow := TEventClass.OnShow;
  TEventClass.Add(AObj, @TEventClass.OnShow, AEventId);
end;

function PageSetupDialog_GetComponentCount(AObj: TPageSetupDialog): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function PageSetupDialog_GetComponentIndex(AObj: TPageSetupDialog): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure PageSetupDialog_SetComponentIndex(AObj: TPageSetupDialog; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function PageSetupDialog_GetOwner(AObj: TPageSetupDialog): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function PageSetupDialog_GetName(AObj: TPageSetupDialog): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure PageSetupDialog_SetName(AObj: TPageSetupDialog; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function PageSetupDialog_GetTag(AObj: TPageSetupDialog): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure PageSetupDialog_SetTag(AObj: TPageSetupDialog; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function PageSetupDialog_GetComponents(AObj: TPageSetupDialog; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;

function PageSetupDialog_StaticClassType: TClass; stdcall;
begin
  Result := TPageSetupDialog;
end;

exports
  PageSetupDialog_Create,
  PageSetupDialog_Free,
  PageSetupDialog_GetDefaults,
  PageSetupDialog_Execute,
  PageSetupDialog_FindComponent,
  PageSetupDialog_GetNamePath,
  PageSetupDialog_HasParent,
  PageSetupDialog_Assign,
  PageSetupDialog_DisposeOf,
  PageSetupDialog_ClassType,
  PageSetupDialog_ClassName,
  PageSetupDialog_InstanceSize,
  PageSetupDialog_InheritsFrom,
  PageSetupDialog_Equals,
  PageSetupDialog_GetHashCode,
  PageSetupDialog_ToString,
  PageSetupDialog_GetMinMarginLeft,
  PageSetupDialog_SetMinMarginLeft,
  PageSetupDialog_GetMinMarginTop,
  PageSetupDialog_SetMinMarginTop,
  PageSetupDialog_GetMinMarginRight,
  PageSetupDialog_SetMinMarginRight,
  PageSetupDialog_GetMinMarginBottom,
  PageSetupDialog_SetMinMarginBottom,
  PageSetupDialog_GetMarginLeft,
  PageSetupDialog_SetMarginLeft,
  PageSetupDialog_GetMarginTop,
  PageSetupDialog_SetMarginTop,
  PageSetupDialog_GetMarginRight,
  PageSetupDialog_SetMarginRight,
  PageSetupDialog_GetMarginBottom,
  PageSetupDialog_SetMarginBottom,
  PageSetupDialog_GetOptions,
  PageSetupDialog_SetOptions,
  PageSetupDialog_GetPageWidth,
  PageSetupDialog_SetPageWidth,
  PageSetupDialog_GetPageHeight,
  PageSetupDialog_SetPageHeight,
  PageSetupDialog_GetUnits,
  PageSetupDialog_SetUnits,
  PageSetupDialog_GetHandle,
  PageSetupDialog_GetCtl3D,
  PageSetupDialog_SetCtl3D,
  PageSetupDialog_SetOnClose,
  PageSetupDialog_SetOnShow,
  PageSetupDialog_GetComponentCount,
  PageSetupDialog_GetComponentIndex,
  PageSetupDialog_SetComponentIndex,
  PageSetupDialog_GetOwner,
  PageSetupDialog_GetName,
  PageSetupDialog_SetName,
  PageSetupDialog_GetTag,
  PageSetupDialog_SetTag,
  PageSetupDialog_GetComponents,
  PageSetupDialog_StaticClassType;

