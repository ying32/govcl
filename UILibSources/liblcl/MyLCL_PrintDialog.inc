
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function PrintDialog_Create(AOwner: TComponent): TPrintDialog; extdecl;
begin
  Result :=  TPrintDialog.Create(AOwner);
end;

procedure PrintDialog_Free(AObj: TPrintDialog); extdecl;
begin
  AObj.Free;
end;

function PrintDialog_Execute(AObj: TPrintDialog): LongBool; extdecl;
begin
  Result :=  AObj.Execute;
end;

function PrintDialog_FindComponent(AObj: TPrintDialog; AName: PChar): TComponent; extdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function PrintDialog_GetNamePath(AObj: TPrintDialog): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function PrintDialog_HasParent(AObj: TPrintDialog): LongBool; extdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure PrintDialog_Assign(AObj: TPrintDialog; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function PrintDialog_ClassName(AObj: TPrintDialog): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function PrintDialog_Equals(AObj: TPrintDialog; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function PrintDialog_GetHashCode(AObj: TPrintDialog): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function PrintDialog_ToString(AObj: TPrintDialog): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function PrintDialog_GetCollate(AObj: TPrintDialog): LongBool; extdecl;
begin
  Result :=  AObj.Collate;
end;

procedure PrintDialog_SetCollate(AObj: TPrintDialog; AValue: LongBool); extdecl;
begin
  AObj.Collate := AValue;
end;

function PrintDialog_GetCopies(AObj: TPrintDialog): Integer; extdecl;
begin
  Result :=  AObj.Copies;
end;

procedure PrintDialog_SetCopies(AObj: TPrintDialog; AValue: Integer); extdecl;
begin
  AObj.Copies := AValue;
end;

function PrintDialog_GetFromPage(AObj: TPrintDialog): Integer; extdecl;
begin
  Result :=  AObj.FromPage;
end;

procedure PrintDialog_SetFromPage(AObj: TPrintDialog; AValue: Integer); extdecl;
begin
  AObj.FromPage := AValue;
end;

function PrintDialog_GetMinPage(AObj: TPrintDialog): Integer; extdecl;
begin
  Result :=  AObj.MinPage;
end;

procedure PrintDialog_SetMinPage(AObj: TPrintDialog; AValue: Integer); extdecl;
begin
  AObj.MinPage := AValue;
end;

function PrintDialog_GetMaxPage(AObj: TPrintDialog): Integer; extdecl;
begin
  Result :=  AObj.MaxPage;
end;

procedure PrintDialog_SetMaxPage(AObj: TPrintDialog; AValue: Integer); extdecl;
begin
  AObj.MaxPage := AValue;
end;

function PrintDialog_GetOptions(AObj: TPrintDialog): TPrintDialogOptions; extdecl;
begin
  Result :=  AObj.Options;
end;

procedure PrintDialog_SetOptions(AObj: TPrintDialog; AValue: TPrintDialogOptions); extdecl;
begin
  AObj.Options := AValue;
end;

function PrintDialog_GetPrintToFile(AObj: TPrintDialog): LongBool; extdecl;
begin
  Result :=  AObj.PrintToFile;
end;

procedure PrintDialog_SetPrintToFile(AObj: TPrintDialog; AValue: LongBool); extdecl;
begin
  AObj.PrintToFile := AValue;
end;

function PrintDialog_GetPrintRange(AObj: TPrintDialog): TPrintRange; extdecl;
begin
  Result :=  AObj.PrintRange;
end;

procedure PrintDialog_SetPrintRange(AObj: TPrintDialog; AValue: TPrintRange); extdecl;
begin
  AObj.PrintRange := AValue;
end;

function PrintDialog_GetToPage(AObj: TPrintDialog): Integer; extdecl;
begin
  Result :=  AObj.ToPage;
end;

procedure PrintDialog_SetToPage(AObj: TPrintDialog; AValue: Integer); extdecl;
begin
  AObj.ToPage := AValue;
end;

function PrintDialog_GetHandle(AObj: TPrintDialog): HWND; extdecl;
begin
  Result :=  AObj.Handle;
end;

procedure PrintDialog_SetOnClose(AObj: TPrintDialog; AEventId: NativeUInt); extdecl;
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

procedure PrintDialog_SetOnShow(AObj: TPrintDialog; AEventId: NativeUInt); extdecl;
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

function PrintDialog_GetComponentCount(AObj: TPrintDialog): Integer; extdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function PrintDialog_GetComponentIndex(AObj: TPrintDialog): Integer; extdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure PrintDialog_SetComponentIndex(AObj: TPrintDialog; AValue: Integer); extdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function PrintDialog_GetOwner(AObj: TPrintDialog): TComponent; extdecl;
begin
  Result :=  AObj.Owner;
end;

function PrintDialog_GetName(AObj: TPrintDialog): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure PrintDialog_SetName(AObj: TPrintDialog; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function PrintDialog_GetTag(AObj: TPrintDialog): NativeInt; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure PrintDialog_SetTag(AObj: TPrintDialog; AValue: NativeInt); extdecl;
begin
  AObj.Tag := AValue;
end;

function PrintDialog_GetComponents(AObj: TPrintDialog; AIndex: Integer): TComponent; extdecl;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  PrintDialog_Create,
  PrintDialog_Free,
  PrintDialog_Execute,
  PrintDialog_FindComponent,
  PrintDialog_GetNamePath,
  PrintDialog_HasParent,
  PrintDialog_Assign,
  PrintDialog_ClassName,
  PrintDialog_Equals,
  PrintDialog_GetHashCode,
  PrintDialog_ToString,
  PrintDialog_GetCollate,
  PrintDialog_SetCollate,
  PrintDialog_GetCopies,
  PrintDialog_SetCopies,
  PrintDialog_GetFromPage,
  PrintDialog_SetFromPage,
  PrintDialog_GetMinPage,
  PrintDialog_SetMinPage,
  PrintDialog_GetMaxPage,
  PrintDialog_SetMaxPage,
  PrintDialog_GetOptions,
  PrintDialog_SetOptions,
  PrintDialog_GetPrintToFile,
  PrintDialog_SetPrintToFile,
  PrintDialog_GetPrintRange,
  PrintDialog_SetPrintRange,
  PrintDialog_GetToPage,
  PrintDialog_SetToPage,
  PrintDialog_GetHandle,
  PrintDialog_SetOnClose,
  PrintDialog_SetOnShow,
  PrintDialog_GetComponentCount,
  PrintDialog_GetComponentIndex,
  PrintDialog_SetComponentIndex,
  PrintDialog_GetOwner,
  PrintDialog_GetName,
  PrintDialog_SetName,
  PrintDialog_GetTag,
  PrintDialog_SetTag,
  PrintDialog_GetComponents;

