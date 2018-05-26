
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function StatusPanels_Create(AOwner: TStatusBar): TStatusPanels; extdecl;
begin
  Result :=  TStatusPanels.Create(AOwner);
end;

procedure StatusPanels_Free(AObj: TStatusPanels); extdecl;
begin
  AObj.Free;
end;

function StatusPanels_Add(AObj: TStatusPanels): TStatusPanel; extdecl;
begin
  Result :=  AObj.Add;
end;

function StatusPanels_Insert(AObj: TStatusPanels; Index: Integer): TStatusPanel; extdecl;
begin
  Result :=  AObj.Insert(Index) as TStatusPanel;
end;

function StatusPanels_Owner(AObj: TStatusPanels): TPersistent; extdecl;
begin
  Result :=  AObj.Owner;
end;

procedure StatusPanels_Assign(AObj: TStatusPanels; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

procedure StatusPanels_BeginUpdate(AObj: TStatusPanels); extdecl;
begin
  AObj.BeginUpdate;
end;

procedure StatusPanels_Clear(AObj: TStatusPanels); extdecl;
begin
  AObj.Clear;
end;

procedure StatusPanels_Delete(AObj: TStatusPanels; Index: Integer); extdecl;
begin
  AObj.Delete(Index);
end;

procedure StatusPanels_EndUpdate(AObj: TStatusPanels); extdecl;
begin
  AObj.EndUpdate;
end;

function StatusPanels_GetNamePath(AObj: TStatusPanels): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function StatusPanels_ClassName(AObj: TStatusPanels): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function StatusPanels_Equals(AObj: TStatusPanels; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function StatusPanels_GetHashCode(AObj: TStatusPanels): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function StatusPanels_ToString(AObj: TStatusPanels): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function StatusPanels_GetCount(AObj: TStatusPanels): Integer; extdecl;
begin
  Result :=  AObj.Count;
end;

function StatusPanels_GetItems(AObj: TStatusPanels; Index: Integer): TStatusPanel; extdecl;
begin
  Result :=  AObj.Items[Index];
end;

procedure StatusPanels_SetItems(AObj: TStatusPanels; Index: Integer; AValue: TStatusPanel); extdecl;
begin
  AObj.Items[Index] := AValue;
end;


exports
  StatusPanels_Create,
  StatusPanels_Free,
  StatusPanels_Add,
  StatusPanels_Insert,
  StatusPanels_Owner,
  StatusPanels_Assign,
  StatusPanels_BeginUpdate,
  StatusPanels_Clear,
  StatusPanels_Delete,
  StatusPanels_EndUpdate,
  StatusPanels_GetNamePath,
  StatusPanels_ClassName,
  StatusPanels_Equals,
  StatusPanels_GetHashCode,
  StatusPanels_ToString,
  StatusPanels_GetCount,
  StatusPanels_GetItems,
  StatusPanels_SetItems;

