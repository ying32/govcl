
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function Timer_Create(AOwner: TComponent): TTimer; extdecl;
begin
  Result :=  TTimer.Create(AOwner);
end;

procedure Timer_Free(AObj: TTimer); extdecl;
begin
  AObj.Free;
end;

function Timer_FindComponent(AObj: TTimer; AName: PChar): TComponent; extdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function Timer_GetNamePath(AObj: TTimer): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function Timer_HasParent(AObj: TTimer): LongBool; extdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure Timer_Assign(AObj: TTimer; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function Timer_ClassName(AObj: TTimer): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function Timer_Equals(AObj: TTimer; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function Timer_GetHashCode(AObj: TTimer): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function Timer_ToString(AObj: TTimer): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function Timer_GetEnabled(AObj: TTimer): LongBool; extdecl;
begin
  Result :=  AObj.Enabled;
end;

procedure Timer_SetEnabled(AObj: TTimer; AValue: LongBool); extdecl;
begin
  AObj.Enabled := AValue;
end;

function Timer_GetInterval(AObj: TTimer): Cardinal; extdecl;
begin
  Result :=  AObj.Interval;
end;

procedure Timer_SetInterval(AObj: TTimer; AValue: Cardinal); extdecl;
begin
  AObj.Interval := AValue;
end;

procedure Timer_SetOnTimer(AObj: TTimer; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnTimer := nil;
    EventClass.Remove(AObj, geTimer);
    Exit;
  end;
  AObj.OnTimer := @EventClass.OnTimer;
  EventClass.Add(AObj, geTimer, AEventId);
end;

function Timer_GetComponentCount(AObj: TTimer): Integer; extdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function Timer_GetComponentIndex(AObj: TTimer): Integer; extdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure Timer_SetComponentIndex(AObj: TTimer; AValue: Integer); extdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function Timer_GetOwner(AObj: TTimer): TComponent; extdecl;
begin
  Result :=  AObj.Owner;
end;

function Timer_GetName(AObj: TTimer): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure Timer_SetName(AObj: TTimer; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function Timer_GetTag(AObj: TTimer): NativeInt; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure Timer_SetTag(AObj: TTimer; AValue: NativeInt); extdecl;
begin
  AObj.Tag := AValue;
end;

function Timer_GetComponents(AObj: TTimer; AIndex: Integer): TComponent; extdecl;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  Timer_Create,
  Timer_Free,
  Timer_FindComponent,
  Timer_GetNamePath,
  Timer_HasParent,
  Timer_Assign,
  Timer_ClassName,
  Timer_Equals,
  Timer_GetHashCode,
  Timer_ToString,
  Timer_GetEnabled,
  Timer_SetEnabled,
  Timer_GetInterval,
  Timer_SetInterval,
  Timer_SetOnTimer,
  Timer_GetComponentCount,
  Timer_GetComponentIndex,
  Timer_SetComponentIndex,
  Timer_GetOwner,
  Timer_GetName,
  Timer_SetName,
  Timer_GetTag,
  Timer_SetTag,
  Timer_GetComponents;

