
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function Bevel_Create(AOwner: TComponent): TBevel; stdcall;
begin
  Result :=  TBevel.Create(AOwner);
end;

procedure Bevel_Free(AObj: TBevel); stdcall;
begin
  AObj.Free;
end;

procedure Bevel_BringToFront(AObj: TBevel); stdcall;
begin
  AObj.BringToFront;
end;

function Bevel_HasParent(AObj: TBevel): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure Bevel_Hide(AObj: TBevel); stdcall;
begin
  AObj.Hide;
end;

procedure Bevel_Invalidate(AObj: TBevel); stdcall;
begin
  AObj.Invalidate;
end;

function Bevel_Perform(AObj: TBevel; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure Bevel_Refresh(AObj: TBevel); stdcall;
begin
  AObj.Refresh;
end;

procedure Bevel_Repaint(AObj: TBevel); stdcall;
begin
  AObj.Repaint;
end;

procedure Bevel_SendToBack(AObj: TBevel); stdcall;
begin
  AObj.SendToBack;
end;

procedure Bevel_SetBounds(AObj: TBevel; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure Bevel_Show(AObj: TBevel); stdcall;
begin
  AObj.Show;
end;

procedure Bevel_Update(AObj: TBevel); stdcall;
begin
  AObj.Update;
end;

function Bevel_GetTextBuf(AObj: TBevel; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function Bevel_GetTextLen(AObj: TBevel): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

function Bevel_FindComponent(AObj: TBevel; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function Bevel_GetNamePath(AObj: TBevel): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure Bevel_Assign(AObj: TBevel; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function Bevel_ClassName(AObj: TBevel): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function Bevel_Equals(AObj: TBevel; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function Bevel_GetHashCode(AObj: TBevel): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function Bevel_ToString(AObj: TBevel): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function Bevel_GetAlign(AObj: TBevel): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure Bevel_SetAlign(AObj: TBevel; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function Bevel_GetAnchors(AObj: TBevel): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure Bevel_SetAnchors(AObj: TBevel; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function Bevel_GetParentShowHint(AObj: TBevel): LongBool; stdcall;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure Bevel_SetParentShowHint(AObj: TBevel; AValue: LongBool); stdcall;
begin
  AObj.ParentShowHint := AValue;
end;

function Bevel_GetShape(AObj: TBevel): TBevelShape; stdcall;
begin
  Result :=  AObj.Shape;
end;

procedure Bevel_SetShape(AObj: TBevel; AValue: TBevelShape); stdcall;
begin
  AObj.Shape := AValue;
end;

function Bevel_GetShowHint(AObj: TBevel): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure Bevel_SetShowHint(AObj: TBevel; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function Bevel_GetStyle(AObj: TBevel): TBevelStyle; stdcall;
begin
  Result :=  AObj.Style;
end;

procedure Bevel_SetStyle(AObj: TBevel; AValue: TBevelStyle); stdcall;
begin
  AObj.Style := AValue;
end;

function Bevel_GetVisible(AObj: TBevel): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure Bevel_SetVisible(AObj: TBevel; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

function Bevel_GetEnabled(AObj: TBevel): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure Bevel_SetEnabled(AObj: TBevel; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function Bevel_GetAction(AObj: TBevel): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure Bevel_SetAction(AObj: TBevel; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

function Bevel_GetBiDiMode(AObj: TBevel): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure Bevel_SetBiDiMode(AObj: TBevel; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

procedure Bevel_GetBoundsRect(AObj: TBevel; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure Bevel_SetBoundsRect(AObj: TBevel; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function Bevel_GetClientHeight(AObj: TBevel): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure Bevel_SetClientHeight(AObj: TBevel; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure Bevel_GetClientRect(AObj: TBevel; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function Bevel_GetClientWidth(AObj: TBevel): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure Bevel_SetClientWidth(AObj: TBevel; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function Bevel_GetExplicitLeft(AObj: TBevel): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function Bevel_GetExplicitTop(AObj: TBevel): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function Bevel_GetExplicitWidth(AObj: TBevel): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function Bevel_GetExplicitHeight(AObj: TBevel): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function Bevel_GetParent(AObj: TBevel): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure Bevel_SetParent(AObj: TBevel; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function Bevel_GetStyleElements(AObj: TBevel): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure Bevel_SetStyleElements(AObj: TBevel; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

function Bevel_GetAlignWithMargins(AObj: TBevel): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure Bevel_SetAlignWithMargins(AObj: TBevel; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function Bevel_GetLeft(AObj: TBevel): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure Bevel_SetLeft(AObj: TBevel; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function Bevel_GetTop(AObj: TBevel): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure Bevel_SetTop(AObj: TBevel; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function Bevel_GetWidth(AObj: TBevel): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure Bevel_SetWidth(AObj: TBevel; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function Bevel_GetHeight(AObj: TBevel): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure Bevel_SetHeight(AObj: TBevel; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function Bevel_GetCursor(AObj: TBevel): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure Bevel_SetCursor(AObj: TBevel; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function Bevel_GetHint(AObj: TBevel): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure Bevel_SetHint(AObj: TBevel; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function Bevel_GetMargins(AObj: TBevel): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure Bevel_SetMargins(AObj: TBevel; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function Bevel_GetCustomHint(AObj: TBevel): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure Bevel_SetCustomHint(AObj: TBevel; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function Bevel_GetComponentCount(AObj: TBevel): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function Bevel_GetComponentIndex(AObj: TBevel): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure Bevel_SetComponentIndex(AObj: TBevel; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function Bevel_GetOwner(AObj: TBevel): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function Bevel_GetName(AObj: TBevel): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure Bevel_SetName(AObj: TBevel; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function Bevel_GetTag(AObj: TBevel): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure Bevel_SetTag(AObj: TBevel; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function Bevel_GetComponents(AObj: TBevel; AIndex: Integer): TComponent; stdcall;
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
  Bevel_GetExplicitLeft,
  Bevel_GetExplicitTop,
  Bevel_GetExplicitWidth,
  Bevel_GetExplicitHeight,
  Bevel_GetParent,
  Bevel_SetParent,
  Bevel_GetStyleElements,
  Bevel_SetStyleElements,
  Bevel_GetAlignWithMargins,
  Bevel_SetAlignWithMargins,
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
  Bevel_GetMargins,
  Bevel_SetMargins,
  Bevel_GetCustomHint,
  Bevel_SetCustomHint,
  Bevel_GetComponentCount,
  Bevel_GetComponentIndex,
  Bevel_SetComponentIndex,
  Bevel_GetOwner,
  Bevel_GetName,
  Bevel_SetName,
  Bevel_GetTag,
  Bevel_SetTag,
  Bevel_GetComponents;

