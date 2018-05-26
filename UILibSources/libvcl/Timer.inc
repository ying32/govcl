
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function Timer_Create(AOwner: TComponent): TTimer; stdcall;
begin
  Result :=  TTimer.Create(AOwner);
end;

procedure Timer_Free(AObj: TTimer); stdcall;
begin
  AObj.Free;
end;

function Timer_FindComponent(AObj: TTimer; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function Timer_GetNamePath(AObj: TTimer): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function Timer_HasParent(AObj: TTimer): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure Timer_Assign(AObj: TTimer; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function Timer_ClassName(AObj: TTimer): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function Timer_Equals(AObj: TTimer; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function Timer_GetHashCode(AObj: TTimer): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function Timer_ToString(AObj: TTimer): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function Timer_GetEnabled(AObj: TTimer): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure Timer_SetEnabled(AObj: TTimer; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function Timer_GetInterval(AObj: TTimer): Cardinal; stdcall;
begin
  Result :=  AObj.Interval;
end;

procedure Timer_SetInterval(AObj: TTimer; AValue: Cardinal); stdcall;
begin
  AObj.Interval := AValue;
end;

procedure Timer_SetOnTimer(AObj: TTimer; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnTimer := nil;
    TEventClass.Remove(AObj, geTimer);
    Exit;
  end;
  AObj.OnTimer := TEventClass.OnTimer;
  TEventClass.Add(AObj, geTimer, AEventId);
end;

function Timer_GetComponentCount(AObj: TTimer): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function Timer_GetComponentIndex(AObj: TTimer): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure Timer_SetComponentIndex(AObj: TTimer; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function Timer_GetOwner(AObj: TTimer): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function Timer_GetName(AObj: TTimer): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure Timer_SetName(AObj: TTimer; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function Timer_GetTag(AObj: TTimer): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure Timer_SetTag(AObj: TTimer; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function Timer_GetComponents(AObj: TTimer; AIndex: Integer): TComponent; stdcall;
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

