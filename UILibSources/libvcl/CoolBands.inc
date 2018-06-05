
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function CoolBands_Create(AOwner: TCoolBar): TCoolBands; stdcall;
begin
  Result :=  TCoolBands.Create(AOwner);
end;

procedure CoolBands_Free(AObj: TCoolBands); stdcall;
begin
  AObj.Free;
end;

function CoolBands_Add(AObj: TCoolBands): TCoolBand; stdcall;
begin
  Result :=  AObj.Add;
end;

function CoolBands_FindBand(AObj: TCoolBands; AControl: TControl): TCoolBand; stdcall;
begin
  Result :=  AObj.FindBand(AControl);
end;

function CoolBands_Owner(AObj: TCoolBands): TPersistent; stdcall;
begin
  Result :=  AObj.Owner;
end;

procedure CoolBands_Assign(AObj: TCoolBands; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

procedure CoolBands_BeginUpdate(AObj: TCoolBands); stdcall;
begin
  AObj.BeginUpdate;
end;

procedure CoolBands_Clear(AObj: TCoolBands); stdcall;
begin
  AObj.Clear;
end;

procedure CoolBands_Delete(AObj: TCoolBands; Index: Integer); stdcall;
begin
  AObj.Delete(Index);
end;

procedure CoolBands_EndUpdate(AObj: TCoolBands); stdcall;
begin
  AObj.EndUpdate;
end;

function CoolBands_GetNamePath(AObj: TCoolBands): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function CoolBands_Insert(AObj: TCoolBands; Index: Integer): TCollectionItem; stdcall;
begin
  Result :=  AObj.Insert(Index);
end;

procedure CoolBands_DisposeOf(AObj: TCoolBands); stdcall;
begin
  AObj.DisposeOf;
end;

function CoolBands_ClassType(AObj: TCoolBands): TClass; stdcall;
begin
  Result :=  AObj.ClassType;
end;

function CoolBands_ClassName(AObj: TCoolBands): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function CoolBands_InstanceSize(AObj: TCoolBands): Integer; stdcall;
begin
  Result :=  AObj.InstanceSize;
end;

function CoolBands_InheritsFrom(AObj: TCoolBands; AClass: TClass): LongBool; stdcall;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function CoolBands_Equals(AObj: TCoolBands; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function CoolBands_GetHashCode(AObj: TCoolBands): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function CoolBands_ToString(AObj: TCoolBands): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function CoolBands_GetCoolBar(AObj: TCoolBands): TCoolBar; stdcall;
begin
  Result :=  AObj.CoolBar;
end;

function CoolBands_GetCapacity(AObj: TCoolBands): Integer; stdcall;
begin
  Result :=  AObj.Capacity;
end;

procedure CoolBands_SetCapacity(AObj: TCoolBands; AValue: Integer); stdcall;
begin
  AObj.Capacity := AValue;
end;

function CoolBands_GetCount(AObj: TCoolBands): Integer; stdcall;
begin
  Result :=  AObj.Count;
end;

function CoolBands_GetItems(AObj: TCoolBands; Index: Integer): TCoolBand; stdcall;
begin
  Result :=  AObj.Items[Index];
end;

procedure CoolBands_SetItems(AObj: TCoolBands; Index: Integer; AValue: TCoolBand); stdcall;
begin
  AObj.Items[Index] := AValue;
end;

function CoolBands_StaticClassType: TClass; stdcall;
begin
  Result := TCoolBands;
end;

exports
  CoolBands_Create,
  CoolBands_Free,
  CoolBands_Add,
  CoolBands_FindBand,
  CoolBands_Owner,
  CoolBands_Assign,
  CoolBands_BeginUpdate,
  CoolBands_Clear,
  CoolBands_Delete,
  CoolBands_EndUpdate,
  CoolBands_GetNamePath,
  CoolBands_Insert,
  CoolBands_DisposeOf,
  CoolBands_ClassType,
  CoolBands_ClassName,
  CoolBands_InstanceSize,
  CoolBands_InheritsFrom,
  CoolBands_Equals,
  CoolBands_GetHashCode,
  CoolBands_ToString,
  CoolBands_GetCoolBar,
  CoolBands_GetCapacity,
  CoolBands_SetCapacity,
  CoolBands_GetCount,
  CoolBands_GetItems,
  CoolBands_SetItems,
  CoolBands_StaticClassType;

