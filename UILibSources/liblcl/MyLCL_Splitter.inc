
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function Splitter_Create(AOwner: TComponent): TSplitter; extdecl;
begin
  Result :=  TSplitter.Create(AOwner);
end;

procedure Splitter_Free(AObj: TSplitter); extdecl;
begin
  AObj.Free;
end;

procedure Splitter_BringToFront(AObj: TSplitter); extdecl;
begin
  AObj.BringToFront;
end;

function Splitter_HasParent(AObj: TSplitter): LongBool; extdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure Splitter_Hide(AObj: TSplitter); extdecl;
begin
  AObj.Hide;
end;

procedure Splitter_Invalidate(AObj: TSplitter); extdecl;
begin
  AObj.Invalidate;
end;

function Splitter_Perform(AObj: TSplitter; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; extdecl;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure Splitter_Refresh(AObj: TSplitter); extdecl;
begin
  AObj.Refresh;
end;

procedure Splitter_Repaint(AObj: TSplitter); extdecl;
begin
  AObj.Repaint;
end;

procedure Splitter_SendToBack(AObj: TSplitter); extdecl;
begin
  AObj.SendToBack;
end;

procedure Splitter_SetBounds(AObj: TSplitter; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); extdecl;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure Splitter_Show(AObj: TSplitter); extdecl;
begin
  AObj.Show;
end;

procedure Splitter_Update(AObj: TSplitter); extdecl;
begin
  AObj.Update;
end;

function Splitter_GetTextBuf(AObj: TSplitter; Buffer: PChar; BufSize: Integer): Integer; extdecl;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function Splitter_GetTextLen(AObj: TSplitter): Integer; extdecl;
begin
  Result :=  AObj.GetTextLen;
end;

function Splitter_FindComponent(AObj: TSplitter; AName: PChar): TComponent; extdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function Splitter_GetNamePath(AObj: TSplitter): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure Splitter_Assign(AObj: TSplitter; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function Splitter_ClassName(AObj: TSplitter): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function Splitter_Equals(AObj: TSplitter; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function Splitter_GetHashCode(AObj: TSplitter): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function Splitter_ToString(AObj: TSplitter): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function Splitter_GetCanvas(AObj: TSplitter): TCanvas; extdecl;
begin
  Result :=  AObj.Canvas;
end;

function Splitter_GetAlign(AObj: TSplitter): TAlign; extdecl;
begin
  Result :=  AObj.Align;
end;

procedure Splitter_SetAlign(AObj: TSplitter; AValue: TAlign); extdecl;
begin
  AObj.Align := AValue;
end;

function Splitter_GetColor(AObj: TSplitter): TColor; extdecl;
begin
  Result :=  AObj.Color;
end;

procedure Splitter_SetColor(AObj: TSplitter; AValue: TColor); extdecl;
begin
  AObj.Color := AValue;
end;

function Splitter_GetCursor(AObj: TSplitter): TCursor; extdecl;
begin
  Result :=  AObj.Cursor;
end;

procedure Splitter_SetCursor(AObj: TSplitter; AValue: TCursor); extdecl;
begin
  AObj.Cursor := AValue;
end;

function Splitter_GetParentColor(AObj: TSplitter): LongBool; extdecl;
begin
  Result :=  AObj.ParentColor;
end;

procedure Splitter_SetParentColor(AObj: TSplitter; AValue: LongBool); extdecl;
begin
  AObj.ParentColor := AValue;
end;

function Splitter_GetVisible(AObj: TSplitter): LongBool; extdecl;
begin
  Result :=  AObj.Visible;
end;

procedure Splitter_SetVisible(AObj: TSplitter; AValue: LongBool); extdecl;
begin
  AObj.Visible := AValue;
end;

function Splitter_GetWidth(AObj: TSplitter): Integer; extdecl;
begin
  Result :=  AObj.Width;
end;

procedure Splitter_SetWidth(AObj: TSplitter; AValue: Integer); extdecl;
begin
  AObj.Width := AValue;
end;

procedure Splitter_SetOnPaint(AObj: TSplitter; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnPaint := nil;
    EventClass.Remove(AObj, gePaint);
    Exit;
  end;
  AObj.OnPaint := @EventClass.OnPaint;
  EventClass.Add(AObj, gePaint, AEventId);
end;

function Splitter_GetEnabled(AObj: TSplitter): LongBool; extdecl;
begin
  Result :=  AObj.Enabled;
end;

procedure Splitter_SetEnabled(AObj: TSplitter; AValue: LongBool); extdecl;
begin
  AObj.Enabled := AValue;
end;

function Splitter_GetAction(AObj: TSplitter): TBasicAction; extdecl;
begin
  Result :=  AObj.Action;
end;

procedure Splitter_SetAction(AObj: TSplitter; AValue: TBasicAction); extdecl;
begin
  AObj.Action := AValue;
end;

function Splitter_GetAnchors(AObj: TSplitter): TAnchors; extdecl;
begin
  Result :=  AObj.Anchors;
end;

procedure Splitter_SetAnchors(AObj: TSplitter; AValue: TAnchors); extdecl;
begin
  AObj.Anchors := AValue;
end;

function Splitter_GetBiDiMode(AObj: TSplitter): TBiDiMode; extdecl;
begin
  Result :=  AObj.BiDiMode;
end;

procedure Splitter_SetBiDiMode(AObj: TSplitter; AValue: TBiDiMode); extdecl;
begin
  AObj.BiDiMode := AValue;
end;

procedure Splitter_GetBoundsRect(AObj: TSplitter; var Result: TRect); extdecl;
begin
  Result :=  AObj.BoundsRect;
end;

procedure Splitter_SetBoundsRect(AObj: TSplitter; var AValue: TRect); extdecl;
begin
  AObj.BoundsRect := AValue;
end;

function Splitter_GetClientHeight(AObj: TSplitter): Integer; extdecl;
begin
  Result :=  AObj.ClientHeight;
end;

procedure Splitter_SetClientHeight(AObj: TSplitter; AValue: Integer); extdecl;
begin
  AObj.ClientHeight := AValue;
end;

procedure Splitter_GetClientRect(AObj: TSplitter; var Result: TRect); extdecl;
begin
  Result :=  AObj.ClientRect;
end;

function Splitter_GetClientWidth(AObj: TSplitter): Integer; extdecl;
begin
  Result :=  AObj.ClientWidth;
end;

procedure Splitter_SetClientWidth(AObj: TSplitter; AValue: Integer); extdecl;
begin
  AObj.ClientWidth := AValue;
end;

function Splitter_GetShowHint(AObj: TSplitter): LongBool; extdecl;
begin
  Result :=  AObj.ShowHint;
end;

procedure Splitter_SetShowHint(AObj: TSplitter; AValue: LongBool); extdecl;
begin
  AObj.ShowHint := AValue;
end;

function Splitter_GetParent(AObj: TSplitter): TWinControl; extdecl;
begin
  Result :=  AObj.Parent;
end;

procedure Splitter_SetParent(AObj: TSplitter; AValue: TWinControl); extdecl;
begin
  AObj.Parent := AValue;
end;

function Splitter_GetLeft(AObj: TSplitter): Integer; extdecl;
begin
  Result :=  AObj.Left;
end;

procedure Splitter_SetLeft(AObj: TSplitter; AValue: Integer); extdecl;
begin
  AObj.Left := AValue;
end;

function Splitter_GetTop(AObj: TSplitter): Integer; extdecl;
begin
  Result :=  AObj.Top;
end;

procedure Splitter_SetTop(AObj: TSplitter; AValue: Integer); extdecl;
begin
  AObj.Top := AValue;
end;

function Splitter_GetHeight(AObj: TSplitter): Integer; extdecl;
begin
  Result :=  AObj.Height;
end;

procedure Splitter_SetHeight(AObj: TSplitter; AValue: Integer); extdecl;
begin
  AObj.Height := AValue;
end;

function Splitter_GetHint(AObj: TSplitter): PChar; extdecl;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure Splitter_SetHint(AObj: TSplitter; AValue: PChar); extdecl;
begin
  AObj.Hint := AValue;
end;

function Splitter_GetComponentCount(AObj: TSplitter): Integer; extdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function Splitter_GetComponentIndex(AObj: TSplitter): Integer; extdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure Splitter_SetComponentIndex(AObj: TSplitter; AValue: Integer); extdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function Splitter_GetOwner(AObj: TSplitter): TComponent; extdecl;
begin
  Result :=  AObj.Owner;
end;

function Splitter_GetName(AObj: TSplitter): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure Splitter_SetName(AObj: TSplitter; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function Splitter_GetTag(AObj: TSplitter): NativeInt; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure Splitter_SetTag(AObj: TSplitter; AValue: NativeInt); extdecl;
begin
  AObj.Tag := AValue;
end;

function Splitter_GetComponents(AObj: TSplitter; AIndex: Integer): TComponent; extdecl;
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
  Splitter_GetShowHint,
  Splitter_SetShowHint,
  Splitter_GetParent,
  Splitter_SetParent,
  Splitter_GetLeft,
  Splitter_SetLeft,
  Splitter_GetTop,
  Splitter_SetTop,
  Splitter_GetHeight,
  Splitter_SetHeight,
  Splitter_GetHint,
  Splitter_SetHint,
  Splitter_GetComponentCount,
  Splitter_GetComponentIndex,
  Splitter_SetComponentIndex,
  Splitter_GetOwner,
  Splitter_GetName,
  Splitter_SetName,
  Splitter_GetTag,
  Splitter_SetTag,
  Splitter_GetComponents;

