
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function Bevel_Create(AOwner: TComponent): TBevel; extdecl;
begin
  Result :=  TBevel.Create(AOwner);
end;

procedure Bevel_Free(AObj: TBevel); extdecl;
begin
  AObj.Free;
end;

procedure Bevel_BringToFront(AObj: TBevel); extdecl;
begin
  AObj.BringToFront;
end;

function Bevel_HasParent(AObj: TBevel): LongBool; extdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure Bevel_Hide(AObj: TBevel); extdecl;
begin
  AObj.Hide;
end;

procedure Bevel_Invalidate(AObj: TBevel); extdecl;
begin
  AObj.Invalidate;
end;

function Bevel_Perform(AObj: TBevel; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; extdecl;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure Bevel_Refresh(AObj: TBevel); extdecl;
begin
  AObj.Refresh;
end;

procedure Bevel_Repaint(AObj: TBevel); extdecl;
begin
  AObj.Repaint;
end;

procedure Bevel_SendToBack(AObj: TBevel); extdecl;
begin
  AObj.SendToBack;
end;

procedure Bevel_SetBounds(AObj: TBevel; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); extdecl;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure Bevel_Show(AObj: TBevel); extdecl;
begin
  AObj.Show;
end;

procedure Bevel_Update(AObj: TBevel); extdecl;
begin
  AObj.Update;
end;

function Bevel_GetTextBuf(AObj: TBevel; Buffer: PChar; BufSize: Integer): Integer; extdecl;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function Bevel_GetTextLen(AObj: TBevel): Integer; extdecl;
begin
  Result :=  AObj.GetTextLen;
end;

function Bevel_FindComponent(AObj: TBevel; AName: PChar): TComponent; extdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function Bevel_GetNamePath(AObj: TBevel): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure Bevel_Assign(AObj: TBevel; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function Bevel_ClassName(AObj: TBevel): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function Bevel_Equals(AObj: TBevel; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function Bevel_GetHashCode(AObj: TBevel): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function Bevel_ToString(AObj: TBevel): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function Bevel_GetAlign(AObj: TBevel): TAlign; extdecl;
begin
  Result :=  AObj.Align;
end;

procedure Bevel_SetAlign(AObj: TBevel; AValue: TAlign); extdecl;
begin
  AObj.Align := AValue;
end;

function Bevel_GetAnchors(AObj: TBevel): TAnchors; extdecl;
begin
  Result :=  AObj.Anchors;
end;

procedure Bevel_SetAnchors(AObj: TBevel; AValue: TAnchors); extdecl;
begin
  AObj.Anchors := AValue;
end;

function Bevel_GetParentShowHint(AObj: TBevel): LongBool; extdecl;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure Bevel_SetParentShowHint(AObj: TBevel; AValue: LongBool); extdecl;
begin
  AObj.ParentShowHint := AValue;
end;

function Bevel_GetShape(AObj: TBevel): TBevelShape; extdecl;
begin
  Result :=  AObj.Shape;
end;

procedure Bevel_SetShape(AObj: TBevel; AValue: TBevelShape); extdecl;
begin
  AObj.Shape := AValue;
end;

function Bevel_GetShowHint(AObj: TBevel): LongBool; extdecl;
begin
  Result :=  AObj.ShowHint;
end;

procedure Bevel_SetShowHint(AObj: TBevel; AValue: LongBool); extdecl;
begin
  AObj.ShowHint := AValue;
end;

function Bevel_GetStyle(AObj: TBevel): TBevelStyle; extdecl;
begin
  Result :=  AObj.Style;
end;

procedure Bevel_SetStyle(AObj: TBevel; AValue: TBevelStyle); extdecl;
begin
  AObj.Style := AValue;
end;

function Bevel_GetVisible(AObj: TBevel): LongBool; extdecl;
begin
  Result :=  AObj.Visible;
end;

procedure Bevel_SetVisible(AObj: TBevel; AValue: LongBool); extdecl;
begin
  AObj.Visible := AValue;
end;

function Bevel_GetEnabled(AObj: TBevel): LongBool; extdecl;
begin
  Result :=  AObj.Enabled;
end;

procedure Bevel_SetEnabled(AObj: TBevel; AValue: LongBool); extdecl;
begin
  AObj.Enabled := AValue;
end;

function Bevel_GetAction(AObj: TBevel): TBasicAction; extdecl;
begin
  Result :=  AObj.Action;
end;

procedure Bevel_SetAction(AObj: TBevel; AValue: TBasicAction); extdecl;
begin
  AObj.Action := AValue;
end;

function Bevel_GetBiDiMode(AObj: TBevel): TBiDiMode; extdecl;
begin
  Result :=  AObj.BiDiMode;
end;

procedure Bevel_SetBiDiMode(AObj: TBevel; AValue: TBiDiMode); extdecl;
begin
  AObj.BiDiMode := AValue;
end;

procedure Bevel_GetBoundsRect(AObj: TBevel; var Result: TRect); extdecl;
begin
  Result :=  AObj.BoundsRect;
end;

procedure Bevel_SetBoundsRect(AObj: TBevel; var AValue: TRect); extdecl;
begin
  AObj.BoundsRect := AValue;
end;

function Bevel_GetClientHeight(AObj: TBevel): Integer; extdecl;
begin
  Result :=  AObj.ClientHeight;
end;

procedure Bevel_SetClientHeight(AObj: TBevel; AValue: Integer); extdecl;
begin
  AObj.ClientHeight := AValue;
end;

procedure Bevel_GetClientRect(AObj: TBevel; var Result: TRect); extdecl;
begin
  Result :=  AObj.ClientRect;
end;

function Bevel_GetClientWidth(AObj: TBevel): Integer; extdecl;
begin
  Result :=  AObj.ClientWidth;
end;

procedure Bevel_SetClientWidth(AObj: TBevel; AValue: Integer); extdecl;
begin
  AObj.ClientWidth := AValue;
end;

function Bevel_GetParent(AObj: TBevel): TWinControl; extdecl;
begin
  Result :=  AObj.Parent;
end;

procedure Bevel_SetParent(AObj: TBevel; AValue: TWinControl); extdecl;
begin
  AObj.Parent := AValue;
end;

function Bevel_GetLeft(AObj: TBevel): Integer; extdecl;
begin
  Result :=  AObj.Left;
end;

procedure Bevel_SetLeft(AObj: TBevel; AValue: Integer); extdecl;
begin
  AObj.Left := AValue;
end;

function Bevel_GetTop(AObj: TBevel): Integer; extdecl;
begin
  Result :=  AObj.Top;
end;

procedure Bevel_SetTop(AObj: TBevel; AValue: Integer); extdecl;
begin
  AObj.Top := AValue;
end;

function Bevel_GetWidth(AObj: TBevel): Integer; extdecl;
begin
  Result :=  AObj.Width;
end;

procedure Bevel_SetWidth(AObj: TBevel; AValue: Integer); extdecl;
begin
  AObj.Width := AValue;
end;

function Bevel_GetHeight(AObj: TBevel): Integer; extdecl;
begin
  Result :=  AObj.Height;
end;

procedure Bevel_SetHeight(AObj: TBevel; AValue: Integer); extdecl;
begin
  AObj.Height := AValue;
end;

function Bevel_GetCursor(AObj: TBevel): TCursor; extdecl;
begin
  Result :=  AObj.Cursor;
end;

procedure Bevel_SetCursor(AObj: TBevel; AValue: TCursor); extdecl;
begin
  AObj.Cursor := AValue;
end;

function Bevel_GetHint(AObj: TBevel): PChar; extdecl;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure Bevel_SetHint(AObj: TBevel; AValue: PChar); extdecl;
begin
  AObj.Hint := AValue;
end;

function Bevel_GetComponentCount(AObj: TBevel): Integer; extdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function Bevel_GetComponentIndex(AObj: TBevel): Integer; extdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure Bevel_SetComponentIndex(AObj: TBevel; AValue: Integer); extdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function Bevel_GetOwner(AObj: TBevel): TComponent; extdecl;
begin
  Result :=  AObj.Owner;
end;

function Bevel_GetName(AObj: TBevel): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure Bevel_SetName(AObj: TBevel; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function Bevel_GetTag(AObj: TBevel): NativeInt; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure Bevel_SetTag(AObj: TBevel; AValue: NativeInt); extdecl;
begin
  AObj.Tag := AValue;
end;

function Bevel_GetComponents(AObj: TBevel; AIndex: Integer): TComponent; extdecl;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  Bevel_Create,
  Bevel_Free,
  Bevel_BringToFront,
  Bevel_HasParent,
  Bevel_Hide,
  Bevel_Invalidate,
  Bevel_Perform,
  Bevel_Refresh,
  Bevel_Repaint,
  Bevel_SendToBack,
  Bevel_SetBounds,
  Bevel_Show,
  Bevel_Update,
  Bevel_GetTextBuf,
  Bevel_GetTextLen,
  Bevel_FindComponent,
  Bevel_GetNamePath,
  Bevel_Assign,
  Bevel_ClassName,
  Bevel_Equals,
  Bevel_GetHashCode,
  Bevel_ToString,
  Bevel_GetAlign,
  Bevel_SetAlign,
  Bevel_GetAnchors,
  Bevel_SetAnchors,
  Bevel_GetParentShowHint,
  Bevel_SetParentShowHint,
  Bevel_GetShape,
  Bevel_SetShape,
  Bevel_GetShowHint,
  Bevel_SetShowHint,
  Bevel_GetStyle,
  Bevel_SetStyle,
  Bevel_GetVisible,
  Bevel_SetVisible,
  Bevel_GetEnabled,
  Bevel_SetEnabled,
  Bevel_GetAction,
  Bevel_SetAction,
  Bevel_GetBiDiMode,
  Bevel_SetBiDiMode,
  Bevel_GetBoundsRect,
  Bevel_SetBoundsRect,
  Bevel_GetClientHeight,
  Bevel_SetClientHeight,
  Bevel_GetClientRect,
  Bevel_GetClientWidth,
  Bevel_SetClientWidth,
  Bevel_GetParent,
  Bevel_SetParent,
  Bevel_GetLeft,
  Bevel_SetLeft,
  Bevel_GetTop,
  Bevel_SetTop,
  Bevel_GetWidth,
  Bevel_SetWidth,
  Bevel_GetHeight,
  Bevel_SetHeight,
  Bevel_GetCursor,
  Bevel_SetCursor,
  Bevel_GetHint,
  Bevel_SetHint,
  Bevel_GetComponentCount,
  Bevel_GetComponentIndex,
  Bevel_SetComponentIndex,
  Bevel_GetOwner,
  Bevel_GetName,
  Bevel_SetName,
  Bevel_GetTag,
  Bevel_SetTag,
  Bevel_GetComponents;

