
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function Control_Create(AOwner: TComponent): TControl; stdcall;
begin
  Result :=  TControl.Create(AOwner);
end;

procedure Control_Free(AObj: TControl); stdcall;
begin
  AObj.Free;
end;

procedure Control_BringToFront(AObj: TControl); stdcall;
begin
  AObj.BringToFront;
end;

function Control_HasParent(AObj: TControl): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure Control_Hide(AObj: TControl); stdcall;
begin
  AObj.Hide;
end;

procedure Control_Invalidate(AObj: TControl); stdcall;
begin
  AObj.Invalidate;
end;

function Control_Perform(AObj: TControl; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure Control_Refresh(AObj: TControl); stdcall;
begin
  AObj.Refresh;
end;

procedure Control_Repaint(AObj: TControl); stdcall;
begin
  AObj.Repaint;
end;

procedure Control_SendToBack(AObj: TControl); stdcall;
begin
  AObj.SendToBack;
end;

procedure Control_SetBounds(AObj: TControl; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure Control_Show(AObj: TControl); stdcall;
begin
  AObj.Show;
end;

procedure Control_Update(AObj: TControl); stdcall;
begin
  AObj.Update;
end;

function Control_GetTextBuf(AObj: TControl; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function Control_GetTextLen(AObj: TControl): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

function Control_FindComponent(AObj: TControl; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function Control_GetNamePath(AObj: TControl): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure Control_Assign(AObj: TControl; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function Control_ClassName(AObj: TControl): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function Control_Equals(AObj: TControl; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function Control_GetHashCode(AObj: TControl): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function Control_ToString(AObj: TControl): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function Control_GetEnabled(AObj: TControl): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure Control_SetEnabled(AObj: TControl; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function Control_GetAction(AObj: TControl): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure Control_SetAction(AObj: TControl; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

function Control_GetAlign(AObj: TControl): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure Control_SetAlign(AObj: TControl; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function Control_GetAnchors(AObj: TControl): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure Control_SetAnchors(AObj: TControl; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function Control_GetBiDiMode(AObj: TControl): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure Control_SetBiDiMode(AObj: TControl; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

procedure Control_GetBoundsRect(AObj: TControl; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure Control_SetBoundsRect(AObj: TControl; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function Control_GetClientHeight(AObj: TControl): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure Control_SetClientHeight(AObj: TControl; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure Control_GetClientRect(AObj: TControl; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function Control_GetClientWidth(AObj: TControl): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure Control_SetClientWidth(AObj: TControl; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function Control_GetExplicitLeft(AObj: TControl): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function Control_GetExplicitTop(AObj: TControl): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function Control_GetExplicitWidth(AObj: TControl): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function Control_GetExplicitHeight(AObj: TControl): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function Control_GetShowHint(AObj: TControl): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure Control_SetShowHint(AObj: TControl; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function Control_GetVisible(AObj: TControl): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure Control_SetVisible(AObj: TControl; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

function Control_GetParent(AObj: TControl): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure Control_SetParent(AObj: TControl; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function Control_GetStyleElements(AObj: TControl): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure Control_SetStyleElements(AObj: TControl; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

function Control_GetAlignWithMargins(AObj: TControl): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure Control_SetAlignWithMargins(AObj: TControl; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function Control_GetLeft(AObj: TControl): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure Control_SetLeft(AObj: TControl; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function Control_GetTop(AObj: TControl): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure Control_SetTop(AObj: TControl; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function Control_GetWidth(AObj: TControl): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure Control_SetWidth(AObj: TControl; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function Control_GetHeight(AObj: TControl): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure Control_SetHeight(AObj: TControl; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function Control_GetCursor(AObj: TControl): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure Control_SetCursor(AObj: TControl; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function Control_GetHint(AObj: TControl): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure Control_SetHint(AObj: TControl; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function Control_GetMargins(AObj: TControl): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure Control_SetMargins(AObj: TControl; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function Control_GetCustomHint(AObj: TControl): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure Control_SetCustomHint(AObj: TControl; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function Control_GetComponentCount(AObj: TControl): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function Control_GetComponentIndex(AObj: TControl): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure Control_SetComponentIndex(AObj: TControl; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function Control_GetOwner(AObj: TControl): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function Control_GetName(AObj: TControl): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure Control_SetName(AObj: TControl; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function Control_GetTag(AObj: TControl): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure Control_SetTag(AObj: TControl; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function Control_GetComponents(AObj: TControl; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  Control_Create,
  Control_Free,
  Control_BringToFront,
  Control_HasParent,
  Control_Hide,
  Control_Invalidate,
  Control_Perform,
  Control_Refresh,
  Control_Repaint,
  Control_SendToBack,
  Control_SetBounds,
  Control_Show,
  Control_Update,
  Control_GetTextBuf,
  Control_GetTextLen,
  Control_FindComponent,
  Control_GetNamePath,
  Control_Assign,
  Control_ClassName,
  Control_Equals,
  Control_GetHashCode,
  Control_ToString,
  Control_GetEnabled,
  Control_SetEnabled,
  Control_GetAction,
  Control_SetAction,
  Control_GetAlign,
  Control_SetAlign,
  Control_GetAnchors,
  Control_SetAnchors,
  Control_GetBiDiMode,
  Control_SetBiDiMode,
  Control_GetBoundsRect,
  Control_SetBoundsRect,
  Control_GetClientHeight,
  Control_SetClientHeight,
  Control_GetClientRect,
  Control_GetClientWidth,
  Control_SetClientWidth,
  Control_GetExplicitLeft,
  Control_GetExplicitTop,
  Control_GetExplicitWidth,
  Control_GetExplicitHeight,
  Control_GetShowHint,
  Control_SetShowHint,
  Control_GetVisible,
  Control_SetVisible,
  Control_GetParent,
  Control_SetParent,
  Control_GetStyleElements,
  Control_SetStyleElements,
  Control_GetAlignWithMargins,
  Control_SetAlignWithMargins,
  Control_GetLeft,
  Control_SetLeft,
  Control_GetTop,
  Control_SetTop,
  Control_GetWidth,
  Control_SetWidth,
  Control_GetHeight,
  Control_SetHeight,
  Control_GetCursor,
  Control_SetCursor,
  Control_GetHint,
  Control_SetHint,
  Control_GetMargins,
  Control_SetMargins,
  Control_GetCustomHint,
  Control_SetCustomHint,
  Control_GetComponentCount,
  Control_GetComponentIndex,
  Control_SetComponentIndex,
  Control_GetOwner,
  Control_GetName,
  Control_SetName,
  Control_GetTag,
  Control_SetTag,
  Control_GetComponents;

