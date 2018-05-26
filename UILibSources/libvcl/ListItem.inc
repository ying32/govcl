
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function ListItem_Create(AOwner: TListItems): TListItem; stdcall;
begin
  Result :=  TListItem.Create(AOwner);
end;

procedure ListItem_Free(AObj: TListItem); stdcall;
begin
  AObj.Free;
end;

procedure ListItem_Assign(AObj: TListItem; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

procedure ListItem_CancelEdit(AObj: TListItem); stdcall;
begin
  AObj.CancelEdit;
end;

procedure ListItem_Delete(AObj: TListItem); stdcall;
begin
  AObj.Delete;
end;

procedure ListItem_DisplayRect(AObj: TListItem; Code: TDisplayCode; var Result: TRect); stdcall;
begin
  Result :=  AObj.DisplayRect(Code);
end;

function ListItem_EditCaption(AObj: TListItem): LongBool; stdcall;
begin
  Result :=  AObj.EditCaption;
end;

procedure ListItem_MakeVisible(AObj: TListItem; PartialOK: LongBool); stdcall;
begin
  AObj.MakeVisible(PartialOK);
end;

procedure ListItem_Update(AObj: TListItem); stdcall;
begin
  AObj.Update;
end;

function ListItem_WorkArea(AObj: TListItem): Integer; stdcall;
begin
  Result :=  AObj.WorkArea;
end;

function ListItem_GetNamePath(AObj: TListItem): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function ListItem_ClassName(AObj: TListItem): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function ListItem_Equals(AObj: TListItem; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function ListItem_GetHashCode(AObj: TListItem): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function ListItem_ToString(AObj: TListItem): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function ListItem_GetCaption(AObj: TListItem): PChar; stdcall;
begin
  Result :=  PChar(AObj.Caption);
end;

procedure ListItem_SetCaption(AObj: TListItem; AValue: PChar); stdcall;
begin
  AObj.Caption := AValue;
end;

function ListItem_GetChecked(AObj: TListItem): LongBool; stdcall;
begin
  Result :=  AObj.Checked;
end;

procedure ListItem_SetChecked(AObj: TListItem; AValue: LongBool); stdcall;
begin
  AObj.Checked := AValue;
end;

function ListItem_GetCut(AObj: TListItem): LongBool; stdcall;
begin
  Result :=  AObj.Cut;
end;

procedure ListItem_SetCut(AObj: TListItem; AValue: LongBool); stdcall;
begin
  AObj.Cut := AValue;
end;

function ListItem_GetData(AObj: TListItem): Pointer; stdcall;
begin
  Result :=  AObj.Data;
end;

procedure ListItem_SetData(AObj: TListItem; AValue: Pointer); stdcall;
begin
  AObj.Data := AValue;
end;

function ListItem_GetDeleting(AObj: TListItem): LongBool; stdcall;
begin
  Result :=  AObj.Deleting;
end;

function ListItem_GetDropTarget(AObj: TListItem): LongBool; stdcall;
begin
  Result :=  AObj.DropTarget;
end;

procedure ListItem_SetDropTarget(AObj: TListItem; AValue: LongBool); stdcall;
begin
  AObj.DropTarget := AValue;
end;

function ListItem_GetFocused(AObj: TListItem): LongBool; stdcall;
begin
  Result :=  AObj.Focused;
end;

procedure ListItem_SetFocused(AObj: TListItem; AValue: LongBool); stdcall;
begin
  AObj.Focused := AValue;
end;

function ListItem_GetGroupID(AObj: TListItem): Integer; stdcall;
begin
  Result :=  AObj.GroupID;
end;

procedure ListItem_SetGroupID(AObj: TListItem; AValue: Integer); stdcall;
begin
  AObj.GroupID := AValue;
end;

function ListItem_GetHandle(AObj: TListItem): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

function ListItem_GetImageIndex(AObj: TListItem): TImageIndex; stdcall;
begin
  Result :=  AObj.ImageIndex;
end;

procedure ListItem_SetImageIndex(AObj: TListItem; AValue: TImageIndex); stdcall;
begin
  AObj.ImageIndex := AValue;
end;

function ListItem_GetIndent(AObj: TListItem): Integer; stdcall;
begin
  Result :=  AObj.Indent;
end;

procedure ListItem_SetIndent(AObj: TListItem; AValue: Integer); stdcall;
begin
  AObj.Indent := AValue;
end;

function ListItem_GetIndex(AObj: TListItem): Integer; stdcall;
begin
  Result :=  AObj.Index;
end;

function ListItem_GetLeft(AObj: TListItem): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure ListItem_SetLeft(AObj: TListItem; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function ListItem_GetOwner(AObj: TListItem): TListItems; stdcall;
begin
  Result :=  AObj.Owner;
end;

function ListItem_GetOverlayIndex(AObj: TListItem): TImageIndex; stdcall;
begin
  Result :=  AObj.OverlayIndex;
end;

procedure ListItem_SetOverlayIndex(AObj: TListItem; AValue: TImageIndex); stdcall;
begin
  AObj.OverlayIndex := AValue;
end;

procedure ListItem_GetPosition(AObj: TListItem; var Result: TPoint); stdcall;
begin
  Result :=  AObj.Position;
end;

procedure ListItem_SetPosition(AObj: TListItem; var AValue: TPoint); stdcall;
begin
  AObj.Position := AValue;
end;

function ListItem_GetSelected(AObj: TListItem): LongBool; stdcall;
begin
  Result :=  AObj.Selected;
end;

procedure ListItem_SetSelected(AObj: TListItem; AValue: LongBool); stdcall;
begin
  AObj.Selected := AValue;
end;

function ListItem_GetStateIndex(AObj: TListItem): TImageIndex; stdcall;
begin
  Result :=  AObj.StateIndex;
end;

procedure ListItem_SetStateIndex(AObj: TListItem; AValue: TImageIndex); stdcall;
begin
  AObj.StateIndex := AValue;
end;

function ListItem_GetSubItems(AObj: TListItem): TStrings; stdcall;
begin
  Result :=  AObj.SubItems;
end;

procedure ListItem_SetSubItems(AObj: TListItem; AValue: TStrings); stdcall;
begin
  AObj.SubItems := AValue;
end;

function ListItem_GetTop(AObj: TListItem): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure ListItem_SetTop(AObj: TListItem; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function ListItem_GetSubItemImages(AObj: TListItem; Index: Integer): Integer; stdcall;
begin
  Result :=  AObj.SubItemImages[Index];
end;

procedure ListItem_SetSubItemImages(AObj: TListItem; Index: Integer; AValue: Integer); stdcall;
begin
  AObj.SubItemImages[Index] := AValue;
end;


exports
  ListItem_Create,
  ListItem_Free,
  ListItem_Assign,
  ListItem_CancelEdit,
  ListItem_Delete,
  ListItem_DisplayRect,
  ListItem_EditCaption,
  ListItem_MakeVisible,
  ListItem_Update,
  ListItem_WorkArea,
  ListItem_GetNamePath,
  ListItem_ClassName,
  ListItem_Equals,
  ListItem_GetHashCode,
  ListItem_ToString,
  ListItem_GetCaption,
  ListItem_SetCaption,
  ListItem_GetChecked,
  ListItem_SetChecked,
  ListItem_GetCut,
  ListItem_SetCut,
  ListItem_GetData,
  ListItem_SetData,
  ListItem_GetDeleting,
  ListItem_GetDropTarget,
  ListItem_SetDropTarget,
  ListItem_GetFocused,
  ListItem_SetFocused,
  ListItem_GetGroupID,
  ListItem_SetGroupID,
  ListItem_GetHandle,
  ListItem_GetImageIndex,
  ListItem_SetImageIndex,
  ListItem_GetIndent,
  ListItem_SetIndent,
  ListItem_GetIndex,
  ListItem_GetLeft,
  ListItem_SetLeft,
  ListItem_GetOwner,
  ListItem_GetOverlayIndex,
  ListItem_SetOverlayIndex,
  ListItem_GetPosition,
  ListItem_SetPosition,
  ListItem_GetSelected,
  ListItem_SetSelected,
  ListItem_GetStateIndex,
  ListItem_SetStateIndex,
  ListItem_GetSubItems,
  ListItem_SetSubItems,
  ListItem_GetTop,
  ListItem_SetTop,
  ListItem_GetSubItemImages,
  ListItem_SetSubItemImages;

