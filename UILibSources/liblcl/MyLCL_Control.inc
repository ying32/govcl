
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function Control_Create(AOwner: TComponent): TControl; extdecl;
begin
  Result :=  TControl.Create(AOwner);
end;

procedure Control_Free(AObj: TControl); extdecl;
begin
  AObj.Free;
end;

procedure Control_BringToFront(AObj: TControl); extdecl;
begin
  AObj.BringToFront;
end;

function Control_HasParent(AObj: TControl): LongBool; extdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure Control_Hide(AObj: TControl); extdecl;
begin
  AObj.Hide;
end;

procedure Control_Invalidate(AObj: TControl); extdecl;
begin
  AObj.Invalidate;
end;

function Control_Perform(AObj: TControl; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; extdecl;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure Control_Refresh(AObj: TControl); extdecl;
begin
  AObj.Refresh;
end;

procedure Control_Repaint(AObj: TControl); extdecl;
begin
  AObj.Repaint;
end;

procedure Control_SendToBack(AObj: TControl); extdecl;
begin
  AObj.SendToBack;
end;

procedure Control_SetBounds(AObj: TControl; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); extdecl;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure Control_Show(AObj: TControl); extdecl;
begin
  AObj.Show;
end;

procedure Control_Update(AObj: TControl); extdecl;
begin
  AObj.Update;
end;

function Control_GetTextBuf(AObj: TControl; Buffer: PChar; BufSize: Integer): Integer; extdecl;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function Control_GetTextLen(AObj: TControl): Integer; extdecl;
begin
  Result :=  AObj.GetTextLen;
end;

function Control_FindComponent(AObj: TControl; AName: PChar): TComponent; extdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function Control_GetNamePath(AObj: TControl): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure Control_Assign(AObj: TControl; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function Control_ClassName(AObj: TControl): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function Control_Equals(AObj: TControl; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function Control_GetHashCode(AObj: TControl): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function Control_ToString(AObj: TControl): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function Control_GetEnabled(AObj: TControl): LongBool; extdecl;
begin
  Result :=  AObj.Enabled;
end;

procedure Control_SetEnabled(AObj: TControl; AValue: LongBool); extdecl;
begin
  AObj.Enabled := AValue;
end;

function Control_GetAction(AObj: TControl): TBasicAction; extdecl;
begin
  Result :=  AObj.Action;
end;

procedure Control_SetAction(AObj: TControl; AValue: TBasicAction); extdecl;
begin
  AObj.Action := AValue;
end;

function Control_GetAlign(AObj: TControl): TAlign; extdecl;
begin
  Result :=  AObj.Align;
end;

procedure Control_SetAlign(AObj: TControl; AValue: TAlign); extdecl;
begin
  AObj.Align := AValue;
end;

function Control_GetAnchors(AObj: TControl): TAnchors; extdecl;
begin
  Result :=  AObj.Anchors;
end;

procedure Control_SetAnchors(AObj: TControl; AValue: TAnchors); extdecl;
begin
  AObj.Anchors := AValue;
end;

function Control_GetBiDiMode(AObj: TControl): TBiDiMode; extdecl;
begin
  Result :=  AObj.BiDiMode;
end;

procedure Control_SetBiDiMode(AObj: TControl; AValue: TBiDiMode); extdecl;
begin
  AObj.BiDiMode := AValue;
end;

procedure Control_GetBoundsRect(AObj: TControl; var Result: TRect); extdecl;
begin
  Result :=  AObj.BoundsRect;
end;

procedure Control_SetBoundsRect(AObj: TControl; var AValue: TRect); extdecl;
begin
  AObj.BoundsRect := AValue;
end;

function Control_GetClientHeight(AObj: TControl): Integer; extdecl;
begin
  Result :=  AObj.ClientHeight;
end;

procedure Control_SetClientHeight(AObj: TControl; AValue: Integer); extdecl;
begin
  AObj.ClientHeight := AValue;
end;

procedure Control_GetClientRect(AObj: TControl; var Result: TRect); extdecl;
begin
  Result :=  AObj.ClientRect;
end;

function Control_GetClientWidth(AObj: TControl): Integer; extdecl;
begin
  Result :=  AObj.ClientWidth;
end;

procedure Control_SetClientWidth(AObj: TControl; AValue: Integer); extdecl;
begin
  AObj.ClientWidth := AValue;
end;

function Control_GetShowHint(AObj: TControl): LongBool; extdecl;
begin
  Result :=  AObj.ShowHint;
end;

procedure Control_SetShowHint(AObj: TControl; AValue: LongBool); extdecl;
begin
  AObj.ShowHint := AValue;
end;

function Control_GetVisible(AObj: TControl): LongBool; extdecl;
begin
  Result :=  AObj.Visible;
end;

procedure Control_SetVisible(AObj: TControl; AValue: LongBool); extdecl;
begin
  AObj.Visible := AValue;
end;

function Control_GetParent(AObj: TControl): TWinControl; extdecl;
begin
  Result :=  AObj.Parent;
end;

procedure Control_SetParent(AObj: TControl; AValue: TWinControl); extdecl;
begin
  AObj.Parent := AValue;
end;

function Control_GetLeft(AObj: TControl): Integer; extdecl;
begin
  Result :=  AObj.Left;
end;

procedure Control_SetLeft(AObj: TControl; AValue: Integer); extdecl;
begin
  AObj.Left := AValue;
end;

function Control_GetTop(AObj: TControl): Integer; extdecl;
begin
  Result :=  AObj.Top;
end;

procedure Control_SetTop(AObj: TControl; AValue: Integer); extdecl;
begin
  AObj.Top := AValue;
end;

function Control_GetWidth(AObj: TControl): Integer; extdecl;
begin
  Result :=  AObj.Width;
end;

procedure Control_SetWidth(AObj: TControl; AValue: Integer); extdecl;
begin
  AObj.Width := AValue;
end;

function Control_GetHeight(AObj: TControl): Integer; extdecl;
begin
  Result :=  AObj.Height;
end;

procedure Control_SetHeight(AObj: TControl; AValue: Integer); extdecl;
begin
  AObj.Height := AValue;
end;

function Control_GetCursor(AObj: TControl): TCursor; extdecl;
begin
  Result :=  AObj.Cursor;
end;

procedure Control_SetCursor(AObj: TControl; AValue: TCursor); extdecl;
begin
  AObj.Cursor := AValue;
end;

function Control_GetHint(AObj: TControl): PChar; extdecl;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure Control_SetHint(AObj: TControl; AValue: PChar); extdecl;
begin
  AObj.Hint := AValue;
end;

function Control_GetComponentCount(AObj: TControl): Integer; extdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function Control_GetComponentIndex(AObj: TControl): Integer; extdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure Control_SetComponentIndex(AObj: TControl; AValue: Integer); extdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function Control_GetOwner(AObj: TControl): TComponent; extdecl;
begin
  Result :=  AObj.Owner;
end;

function Control_GetName(AObj: TControl): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure Control_SetName(AObj: TControl; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function Control_GetTag(AObj: TControl): NativeInt; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure Control_SetTag(AObj: TControl; AValue: NativeInt); extdecl;
begin
  AObj.Tag := AValue;
end;

function Control_GetComponents(AObj: TControl; AIndex: Integer): TComponent; extdecl;
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
  Control_GetShowHint,
  Control_SetShowHint,
  Control_GetVisible,
  Control_SetVisible,
  Control_GetParent,
  Control_SetParent,
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
  Control_GetComponentCount,
  Control_GetComponentIndex,
  Control_SetComponentIndex,
  Control_GetOwner,
  Control_GetName,
  Control_SetName,
  Control_GetTag,
  Control_SetTag,
  Control_GetComponents;

