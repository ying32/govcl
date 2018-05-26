
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function Splitter_Create(AOwner: TComponent): TSplitter; stdcall;
begin
  Result :=  TSplitter.Create(AOwner);
end;

procedure Splitter_Free(AObj: TSplitter); stdcall;
begin
  AObj.Free;
end;

procedure Splitter_BringToFront(AObj: TSplitter); stdcall;
begin
  AObj.BringToFront;
end;

function Splitter_HasParent(AObj: TSplitter): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure Splitter_Hide(AObj: TSplitter); stdcall;
begin
  AObj.Hide;
end;

procedure Splitter_Invalidate(AObj: TSplitter); stdcall;
begin
  AObj.Invalidate;
end;

function Splitter_Perform(AObj: TSplitter; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure Splitter_Refresh(AObj: TSplitter); stdcall;
begin
  AObj.Refresh;
end;

procedure Splitter_Repaint(AObj: TSplitter); stdcall;
begin
  AObj.Repaint;
end;

procedure Splitter_SendToBack(AObj: TSplitter); stdcall;
begin
  AObj.SendToBack;
end;

procedure Splitter_SetBounds(AObj: TSplitter; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure Splitter_Show(AObj: TSplitter); stdcall;
begin
  AObj.Show;
end;

procedure Splitter_Update(AObj: TSplitter); stdcall;
begin
  AObj.Update;
end;

function Splitter_GetTextBuf(AObj: TSplitter; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function Splitter_GetTextLen(AObj: TSplitter): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

function Splitter_FindComponent(AObj: TSplitter; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function Splitter_GetNamePath(AObj: TSplitter): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure Splitter_Assign(AObj: TSplitter; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function Splitter_ClassName(AObj: TSplitter): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function Splitter_Equals(AObj: TSplitter; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function Splitter_GetHashCode(AObj: TSplitter): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function Splitter_ToString(AObj: TSplitter): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function Splitter_GetCanvas(AObj: TSplitter): TCanvas; stdcall;
begin
  Result :=  AObj.Canvas;
end;

function Splitter_GetAlign(AObj: TSplitter): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure Splitter_SetAlign(AObj: TSplitter; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function Splitter_GetColor(AObj: TSplitter): TColor; stdcall;
begin
  Result :=  AObj.Color;
end;

procedure Splitter_SetColor(AObj: TSplitter; AValue: TColor); stdcall;
begin
  AObj.Color := AValue;
end;

function Splitter_GetCursor(AObj: TSplitter): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure Splitter_SetCursor(AObj: TSplitter; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function Splitter_GetParentColor(AObj: TSplitter): LongBool; stdcall;
begin
  Result :=  AObj.ParentColor;
end;

procedure Splitter_SetParentColor(AObj: TSplitter; AValue: LongBool); stdcall;
begin
  AObj.ParentColor := AValue;
end;

function Splitter_GetVisible(AObj: TSplitter): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure Splitter_SetVisible(AObj: TSplitter; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

function Splitter_GetWidth(AObj: TSplitter): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure Splitter_SetWidth(AObj: TSplitter; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function Splitter_GetStyleElements(AObj: TSplitter): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure Splitter_SetStyleElements(AObj: TSplitter; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

procedure Splitter_SetOnPaint(AObj: TSplitter; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnPaint := nil;
    TEventClass.Remove(AObj, gePaint);
    Exit;
  end;
  AObj.OnPaint := TEventClass.OnPaint;
  TEventClass.Add(AObj, gePaint, AEventId);
end;

function Splitter_GetEnabled(AObj: TSplitter): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure Splitter_SetEnabled(AObj: TSplitter; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function Splitter_GetAction(AObj: TSplitter): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure Splitter_SetAction(AObj: TSplitter; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

function Splitter_GetAnchors(AObj: TSplitter): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure Splitter_SetAnchors(AObj: TSplitter; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function Splitter_GetBiDiMode(AObj: TSplitter): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure Splitter_SetBiDiMode(AObj: TSplitter; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

procedure Splitter_GetBoundsRect(AObj: TSplitter; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure Splitter_SetBoundsRect(AObj: TSplitter; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function Splitter_GetClientHeight(AObj: TSplitter): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure Splitter_SetClientHeight(AObj: TSplitter; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure Splitter_GetClientRect(AObj: TSplitter; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function Splitter_GetClientWidth(AObj: TSplitter): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure Splitter_SetClientWidth(AObj: TSplitter; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function Splitter_GetExplicitLeft(AObj: TSplitter): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function Splitter_GetExplicitTop(AObj: TSplitter): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function Splitter_GetExplicitWidth(AObj: TSplitter): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function Splitter_GetExplicitHeight(AObj: TSplitter): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function Splitter_GetShowHint(AObj: TSplitter): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure Splitter_SetShowHint(AObj: TSplitter; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function Splitter_GetParent(AObj: TSplitter): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure Splitter_SetParent(AObj: TSplitter; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function Splitter_GetAlignWithMargins(AObj: TSplitter): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure Splitter_SetAlignWithMargins(AObj: TSplitter; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function Splitter_GetLeft(AObj: TSplitter): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure Splitter_SetLeft(AObj: TSplitter; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function Splitter_GetTop(AObj: TSplitter): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure Splitter_SetTop(AObj: TSplitter; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function Splitter_GetHeight(AObj: TSplitter): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure Splitter_SetHeight(AObj: TSplitter; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function Splitter_GetHint(AObj: TSplitter): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure Splitter_SetHint(AObj: TSplitter; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function Splitter_GetMargins(AObj: TSplitter): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure Splitter_SetMargins(AObj: TSplitter; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function Splitter_GetCustomHint(AObj: TSplitter): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure Splitter_SetCustomHint(AObj: TSplitter; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function Splitter_GetComponentCount(AObj: TSplitter): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function Splitter_GetComponentIndex(AObj: TSplitter): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure Splitter_SetComponentIndex(AObj: TSplitter; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function Splitter_GetOwner(AObj: TSplitter): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function Splitter_GetName(AObj: TSplitter): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure Splitter_SetName(AObj: TSplitter; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function Splitter_GetTag(AObj: TSplitter): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure Splitter_SetTag(AObj: TSplitter; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function Splitter_GetComponents(AObj: TSplitter; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  Splitter_Create,
  Splitter_Free,
  Splitter_BringToFront,
  Splitter_HasParent,
  Splitter_Hide,
  Splitter_Invalidate,
  Splitter_Perform,
  Splitter_Refresh,
  Splitter_Repaint,
  Splitter_SendToBack,
  Splitter_SetBounds,
  Splitter_Show,
  Splitter_Update,
  Splitter_GetTextBuf,
  Splitter_GetTextLen,
  Splitter_FindComponent,
  Splitter_GetNamePath,
  Splitter_Assign,
  Splitter_ClassName,
  Splitter_Equals,
  Splitter_GetHashCode,
  Splitter_ToString,
  Splitter_GetCanvas,
  Splitter_GetAlign,
  Splitter_SetAlign,
  Splitter_GetColor,
  Splitter_SetColor,
  Splitter_GetCursor,
  Splitter_SetCursor,
  Splitter_GetParentColor,
  Splitter_SetParentColor,
  Splitter_GetVisible,
  Splitter_SetVisible,
  Splitter_GetWidth,
  Splitter_SetWidth,
  Splitter_GetStyleElements,
  Splitter_SetStyleElements,
  Splitter_SetOnPaint,
  Splitter_GetEnabled,
  Splitter_SetEnabled,
  Splitter_GetAction,
  Splitter_SetAction,
  Splitter_GetAnchors,
  Splitter_SetAnchors,
  Splitter_GetBiDiMode,
  Splitter_SetBiDiMode,
  Splitter_GetBoundsRect,
  Splitter_SetBoundsRect,
  Splitter_GetClientHeight,
  Splitter_SetClientHeight,
  Splitter_GetClientRect,
  Splitter_GetClientWidth,
  Splitter_SetClientWidth,
  Splitter_GetExplicitLeft,
  Splitter_GetExplicitTop,
  Splitter_GetExplicitWidth,
  Splitter_GetExplicitHeight,
  Splitter_GetShowHint,
  Splitter_SetShowHint,
  Splitter_GetParent,
  Splitter_SetParent,
  Splitter_GetAlignWithMargins,
  Splitter_SetAlignWithMargins,
  Splitter_GetLeft,
  Splitter_SetLeft,
  Splitter_GetTop,
  Splitter_SetTop,
  Splitter_GetHeight,
  Splitter_SetHeight,
  Splitter_GetHint,
  Splitter_SetHint,
  Splitter_GetMargins,
  Splitter_SetMargins,
  Splitter_GetCustomHint,
  Splitter_SetCustomHint,
  Splitter_GetComponentCount,
  Splitter_GetComponentIndex,
  Splitter_SetComponentIndex,
  Splitter_GetOwner,
  Splitter_GetName,
  Splitter_SetName,
  Splitter_GetTag,
  Splitter_SetTag,
  Splitter_GetComponents;

