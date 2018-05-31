
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function Application_Create(AOwner: TComponent): TApplication; stdcall;
begin
  Result :=  TApplication.Create(AOwner);
end;

procedure Application_Free(AObj: TApplication); stdcall;
begin
  AObj.Free;
end;

procedure Application_ActivateHint(AObj: TApplication; var CursorPos: TPoint); stdcall;
begin
  AObj.ActivateHint(CursorPos);
end;

procedure Application_BringToFront(AObj: TApplication); stdcall;
begin
  AObj.BringToFront;
end;

procedure Application_CancelHint(AObj: TApplication); stdcall;
begin
  AObj.CancelHint;
end;

procedure Application_HandleMessage(AObj: TApplication); stdcall;
begin
  AObj.HandleMessage;
end;

procedure Application_HideHint(AObj: TApplication); stdcall;
begin
  AObj.HideHint;
end;

procedure Application_Minimize(AObj: TApplication); stdcall;
begin
  AObj.Minimize;
end;

procedure Application_ModalStarted(AObj: TApplication); stdcall;
begin
  AObj.ModalStarted;
end;

procedure Application_ModalFinished(AObj: TApplication); stdcall;
begin
  AObj.ModalFinished;
end;

procedure Application_NormalizeAllTopMosts(AObj: TApplication); stdcall;
begin
  AObj.NormalizeAllTopMosts;
end;

procedure Application_NormalizeTopMosts(AObj: TApplication); stdcall;
begin
  AObj.NormalizeTopMosts;
end;

procedure Application_ProcessMessages(AObj: TApplication); stdcall;
begin
  AObj.ProcessMessages;
end;

procedure Application_Restore(AObj: TApplication); stdcall;
begin
  AObj.Restore;
end;

procedure Application_RestoreTopMosts(AObj: TApplication); stdcall;
begin
  AObj.RestoreTopMosts;
end;

procedure Application_Terminate(AObj: TApplication); stdcall;
begin
  AObj.Terminate;
end;

function Application_MessageBox(AObj: TApplication; Text: PWideChar; Caption: PWideChar; Flags: Integer): Integer; stdcall;
begin
  Result :=  AObj.MessageBox(Text, Caption, Flags);
end;

function Application_FindComponent(AObj: TApplication; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function Application_GetNamePath(AObj: TApplication): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function Application_HasParent(AObj: TApplication): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure Application_Assign(AObj: TApplication; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function Application_ClassName(AObj: TApplication): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function Application_Equals(AObj: TApplication; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function Application_GetHashCode(AObj: TApplication): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function Application_ToString(AObj: TApplication): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function Application_GetDialogHandle(AObj: TApplication): HWND; stdcall;
begin
  Result :=  AObj.DialogHandle;
end;

procedure Application_SetDialogHandle(AObj: TApplication; AValue: HWND); stdcall;
begin
  AObj.DialogHandle := AValue;
end;

function Application_GetExeName(AObj: TApplication): PChar; stdcall;
begin
  Result :=  PChar(AObj.ExeName);
end;

function Application_GetHint(AObj: TApplication): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure Application_SetHint(AObj: TApplication; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function Application_GetHintColor(AObj: TApplication): TColor; stdcall;
begin
  Result :=  AObj.HintColor;
end;

procedure Application_SetHintColor(AObj: TApplication; AValue: TColor); stdcall;
begin
  AObj.HintColor := AValue;
end;

function Application_GetHintHidePause(AObj: TApplication): Integer; stdcall;
begin
  Result :=  AObj.HintHidePause;
end;

procedure Application_SetHintHidePause(AObj: TApplication; AValue: Integer); stdcall;
begin
  AObj.HintHidePause := AValue;
end;

function Application_GetHintPause(AObj: TApplication): Integer; stdcall;
begin
  Result :=  AObj.HintPause;
end;

procedure Application_SetHintPause(AObj: TApplication; AValue: Integer); stdcall;
begin
  AObj.HintPause := AValue;
end;

function Application_GetHintShortCuts(AObj: TApplication): LongBool; stdcall;
begin
  Result :=  AObj.HintShortCuts;
end;

procedure Application_SetHintShortCuts(AObj: TApplication; AValue: LongBool); stdcall;
begin
  AObj.HintShortCuts := AValue;
end;

function Application_GetHintShortPause(AObj: TApplication): Integer; stdcall;
begin
  Result :=  AObj.HintShortPause;
end;

procedure Application_SetHintShortPause(AObj: TApplication; AValue: Integer); stdcall;
begin
  AObj.HintShortPause := AValue;
end;

function Application_GetIcon(AObj: TApplication): TIcon; stdcall;
begin
  Result :=  AObj.Icon;
end;

procedure Application_SetIcon(AObj: TApplication; AValue: TIcon); stdcall;
begin
  AObj.Icon := AValue;
end;

function Application_GetIsMetropolisUI(AObj: TApplication): LongBool; stdcall;
begin
  Result :=  AObj.IsMetropolisUI;
end;

function Application_GetMainForm(AObj: TApplication): TForm; stdcall;
begin
  Result :=  AObj.MainForm;
end;

function Application_GetMainFormHandle(AObj: TApplication): HWND; stdcall;
begin
  Result :=  AObj.MainFormHandle;
end;

function Application_GetMainFormOnTaskBar(AObj: TApplication): LongBool; stdcall;
begin
  Result :=  AObj.MainFormOnTaskBar;
end;

procedure Application_SetMainFormOnTaskBar(AObj: TApplication; AValue: LongBool); stdcall;
begin
  AObj.MainFormOnTaskBar := AValue;
end;

function Application_GetBiDiMode(AObj: TApplication): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure Application_SetBiDiMode(AObj: TApplication; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

function Application_GetShowHint(AObj: TApplication): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure Application_SetShowHint(AObj: TApplication; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function Application_GetShowMainForm(AObj: TApplication): LongBool; stdcall;
begin
  Result :=  AObj.ShowMainForm;
end;

procedure Application_SetShowMainForm(AObj: TApplication; AValue: LongBool); stdcall;
begin
  AObj.ShowMainForm := AValue;
end;

function Application_GetTitle(AObj: TApplication): PChar; stdcall;
begin
  Result :=  PChar(AObj.Title);
end;

procedure Application_SetTitle(AObj: TApplication; AValue: PChar); stdcall;
begin
  AObj.Title := AValue;
end;

procedure Application_SetOnException(AObj: TApplication; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnException := nil;
    TEventClass.Remove(AObj, geException);
    Exit;
  end;
  AObj.OnException := TEventClass.OnException;
  TEventClass.Add(AObj, geException, AEventId);
end;

procedure Application_SetOnHelp(AObj: TApplication; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnHelp := nil;
    TEventClass.Remove(AObj, geHelp);
    Exit;
  end;
  AObj.OnHelp := TEventClass.OnHelp;
  TEventClass.Add(AObj, geHelp, AEventId);
end;

procedure Application_SetOnHint(AObj: TApplication; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnHint := nil;
    TEventClass.Remove(AObj, geHint);
    Exit;
  end;
  AObj.OnHint := TEventClass.OnHint;
  TEventClass.Add(AObj, geHint, AEventId);
end;

procedure Application_SetOnMessage(AObj: TApplication; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnMessage := nil;
    TEventClass.Remove(AObj, geMessage);
    Exit;
  end;
  AObj.OnMessage := TEventClass.OnMessage;
  TEventClass.Add(AObj, geMessage, AEventId);
end;

procedure Application_SetOnMinimize(AObj: TApplication; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnMinimize := nil;
    TEventClass.Remove(AObj, geMinimize);
    Exit;
  end;
  AObj.OnMinimize := TEventClass.OnMinimize;
  TEventClass.Add(AObj, geMinimize, AEventId);
end;

procedure Application_SetOnRestore(AObj: TApplication; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnRestore := nil;
    TEventClass.Remove(AObj, geRestore);
    Exit;
  end;
  AObj.OnRestore := TEventClass.OnRestore;
  TEventClass.Add(AObj, geRestore, AEventId);
end;

procedure Application_SetOnShortCut(AObj: TApplication; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnShortCut := nil;
    TEventClass.Remove(AObj, geShortCut);
    Exit;
  end;
  AObj.OnShortCut := TEventClass.OnShortCut;
  TEventClass.Add(AObj, geShortCut, AEventId);
end;

function Application_GetHandle(AObj: TApplication): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

procedure Application_SetHandle(AObj: TApplication; AValue: HWND); stdcall;
begin
  AObj.Handle := AValue;
end;

function Application_GetComponentCount(AObj: TApplication): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function Application_GetComponentIndex(AObj: TApplication): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure Application_SetComponentIndex(AObj: TApplication; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function Application_GetOwner(AObj: TApplication): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function Application_GetName(AObj: TApplication): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure Application_SetName(AObj: TApplication; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function Application_GetTag(AObj: TApplication): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure Application_SetTag(AObj: TApplication; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function Application_GetComponents(AObj: TApplication; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  Application_Create,
  Application_Free,
  Application_ActivateHint,
  Application_BringToFront,
  Application_CancelHint,
  Application_HandleMessage,
  Application_HideHint,
  Application_Minimize,
  Application_ModalStarted,
  Application_ModalFinished,
  Application_NormalizeAllTopMosts,
  Application_NormalizeTopMosts,
  Application_ProcessMessages,
  Application_Restore,
  Application_RestoreTopMosts,
  Application_Terminate,
  Application_MessageBox,
  Application_FindComponent,
  Application_GetNamePath,
  Application_HasParent,
  Application_Assign,
  Application_ClassName,
  Application_Equals,
  Application_GetHashCode,
  Application_ToString,
  Application_GetDialogHandle,
  Application_SetDialogHandle,
  Application_GetExeName,
  Application_GetHint,
  Application_SetHint,
  Application_GetHintColor,
  Application_SetHintColor,
  Application_GetHintHidePause,
  Application_SetHintHidePause,
  Application_GetHintPause,
  Application_SetHintPause,
  Application_GetHintShortCuts,
  Application_SetHintShortCuts,
  Application_GetHintShortPause,
  Application_SetHintShortPause,
  Application_GetIcon,
  Application_SetIcon,
  Application_GetIsMetropolisUI,
  Application_GetMainForm,
  Application_GetMainFormHandle,
  Application_GetMainFormOnTaskBar,
  Application_SetMainFormOnTaskBar,
  Application_GetBiDiMode,
  Application_SetBiDiMode,
  Application_GetShowHint,
  Application_SetShowHint,
  Application_GetShowMainForm,
  Application_SetShowMainForm,
  Application_GetTitle,
  Application_SetTitle,
  Application_SetOnException,
  Application_SetOnHelp,
  Application_SetOnHint,
  Application_SetOnMessage,
  Application_SetOnMinimize,
  Application_SetOnRestore,
  Application_SetOnShortCut,
  Application_GetHandle,
  Application_SetHandle,
  Application_GetComponentCount,
  Application_GetComponentIndex,
  Application_SetComponentIndex,
  Application_GetOwner,
  Application_GetName,
  Application_SetName,
  Application_GetTag,
  Application_SetTag,
  Application_GetComponents;

