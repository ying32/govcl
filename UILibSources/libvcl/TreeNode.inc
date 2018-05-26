
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function TreeNode_Create(AOwner: TTreeNodes): TTreeNode; stdcall;
begin
  Result :=  TTreeNode.Create(AOwner);
end;

procedure TreeNode_Free(AObj: TTreeNode); stdcall;
begin
  AObj.Free;
end;

function TreeNode_AlphaSort(AObj: TTreeNode; ARecurse: LongBool): LongBool; stdcall;
begin
  Result :=  AObj.AlphaSort(ARecurse);
end;

procedure TreeNode_Assign(AObj: TTreeNode; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

procedure TreeNode_Collapse(AObj: TTreeNode; Recurse: LongBool); stdcall;
begin
  AObj.Collapse(Recurse);
end;

procedure TreeNode_Delete(AObj: TTreeNode); stdcall;
begin
  AObj.Delete;
end;

procedure TreeNode_DisplayRect(AObj: TTreeNode; TextOnly: LongBool; var Result: TRect); stdcall;
begin
  Result :=  AObj.DisplayRect(TextOnly);
end;

function TreeNode_EditText(AObj: TTreeNode): LongBool; stdcall;
begin
  Result :=  AObj.EditText;
end;

procedure TreeNode_Expand(AObj: TTreeNode; Recurse: LongBool); stdcall;
begin
  AObj.Expand(Recurse);
end;

function TreeNode_IndexOf(AObj: TTreeNode; Value: TTreeNode): Integer; stdcall;
begin
  Result :=  AObj.IndexOf(Value);
end;

procedure TreeNode_MakeVisible(AObj: TTreeNode); stdcall;
begin
  AObj.MakeVisible;
end;

procedure TreeNode_MoveTo(AObj: TTreeNode; Destination: TTreeNode; Mode: TNodeAttachMode); stdcall;
begin
  AObj.MoveTo(Destination, Mode);
end;

function TreeNode_GetNamePath(AObj: TTreeNode): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function TreeNode_ClassName(AObj: TTreeNode): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function TreeNode_Equals(AObj: TTreeNode; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function TreeNode_GetHashCode(AObj: TTreeNode): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function TreeNode_ToString(AObj: TTreeNode): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function TreeNode_GetAbsoluteIndex(AObj: TTreeNode): Integer; stdcall;
begin
  Result :=  AObj.AbsoluteIndex;
end;

function TreeNode_GetCount(AObj: TTreeNode): Integer; stdcall;
begin
  Result :=  AObj.Count;
end;

function TreeNode_GetCut(AObj: TTreeNode): LongBool; stdcall;
begin
  Result :=  AObj.Cut;
end;

procedure TreeNode_SetCut(AObj: TTreeNode; AValue: LongBool); stdcall;
begin
  AObj.Cut := AValue;
end;

function TreeNode_GetData(AObj: TTreeNode): Pointer; stdcall;
begin
  Result :=  AObj.Data;
end;

procedure TreeNode_SetData(AObj: TTreeNode; AValue: Pointer); stdcall;
begin
  AObj.Data := AValue;
end;

function TreeNode_GetDeleting(AObj: TTreeNode): LongBool; stdcall;
begin
  Result :=  AObj.Deleting;
end;

function TreeNode_GetFocused(AObj: TTreeNode): LongBool; stdcall;
begin
  Result :=  AObj.Focused;
end;

procedure TreeNode_SetFocused(AObj: TTreeNode; AValue: LongBool); stdcall;
begin
  AObj.Focused := AValue;
end;

function TreeNode_GetDropTarget(AObj: TTreeNode): LongBool; stdcall;
begin
  Result :=  AObj.DropTarget;
end;

procedure TreeNode_SetDropTarget(AObj: TTreeNode; AValue: LongBool); stdcall;
begin
  AObj.DropTarget := AValue;
end;

function TreeNode_GetSelected(AObj: TTreeNode): LongBool; stdcall;
begin
  Result :=  AObj.Selected;
end;

procedure TreeNode_SetSelected(AObj: TTreeNode; AValue: LongBool); stdcall;
begin
  AObj.Selected := AValue;
end;

function TreeNode_GetExpanded(AObj: TTreeNode): LongBool; stdcall;
begin
  Result :=  AObj.Expanded;
end;

procedure TreeNode_SetExpanded(AObj: TTreeNode; AValue: LongBool); stdcall;
begin
  AObj.Expanded := AValue;
end;

function TreeNode_GetExpandedImageIndex(AObj: TTreeNode): TImageIndex; stdcall;
begin
  Result :=  AObj.ExpandedImageIndex;
end;

procedure TreeNode_SetExpandedImageIndex(AObj: TTreeNode; AValue: TImageIndex); stdcall;
begin
  AObj.ExpandedImageIndex := AValue;
end;

function TreeNode_GetHandle(AObj: TTreeNode): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

function TreeNode_GetHasChildren(AObj: TTreeNode): LongBool; stdcall;
begin
  Result :=  AObj.HasChildren;
end;

procedure TreeNode_SetHasChildren(AObj: TTreeNode; AValue: LongBool); stdcall;
begin
  AObj.HasChildren := AValue;
end;

function TreeNode_GetImageIndex(AObj: TTreeNode): TImageIndex; stdcall;
begin
  Result :=  AObj.ImageIndex;
end;

procedure TreeNode_SetImageIndex(AObj: TTreeNode; AValue: TImageIndex); stdcall;
begin
  AObj.ImageIndex := AValue;
end;

function TreeNode_GetIndex(AObj: TTreeNode): Integer; stdcall;
begin
  Result :=  AObj.Index;
end;

function TreeNode_GetIsVisible(AObj: TTreeNode): LongBool; stdcall;
begin
  Result :=  AObj.IsVisible;
end;

function TreeNode_GetItemId(AObj: TTreeNode): HTREEITEM; stdcall;
begin
  Result :=  AObj.ItemId;
end;

function TreeNode_GetLevel(AObj: TTreeNode): Integer; stdcall;
begin
  Result :=  AObj.Level;
end;

function TreeNode_GetOverlayIndex(AObj: TTreeNode): Integer; stdcall;
begin
  Result :=  AObj.OverlayIndex;
end;

procedure TreeNode_SetOverlayIndex(AObj: TTreeNode; AValue: Integer); stdcall;
begin
  AObj.OverlayIndex := AValue;
end;

function TreeNode_GetOwner(AObj: TTreeNode): TTreeNodes; stdcall;
begin
  Result :=  AObj.Owner;
end;

function TreeNode_GetParent(AObj: TTreeNode): TTreeNode; stdcall;
begin
  Result :=  AObj.Parent;
end;

function TreeNode_GetSelectedIndex(AObj: TTreeNode): Integer; stdcall;
begin
  Result :=  AObj.SelectedIndex;
end;

procedure TreeNode_SetSelectedIndex(AObj: TTreeNode; AValue: Integer); stdcall;
begin
  AObj.SelectedIndex := AValue;
end;

function TreeNode_GetEnabled(AObj: TTreeNode): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure TreeNode_SetEnabled(AObj: TTreeNode; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function TreeNode_GetStateIndex(AObj: TTreeNode): Integer; stdcall;
begin
  Result :=  AObj.StateIndex;
end;

procedure TreeNode_SetStateIndex(AObj: TTreeNode; AValue: Integer); stdcall;
begin
  AObj.StateIndex := AValue;
end;

function TreeNode_GetText(AObj: TTreeNode): PChar; stdcall;
begin
  Result :=  PChar(AObj.Text);
end;

procedure TreeNode_SetText(AObj: TTreeNode; AValue: PChar); stdcall;
begin
  AObj.Text := AValue;
end;

function TreeNode_GetItem(AObj: TTreeNode; Index: Integer): TTreeNode; stdcall;
begin
  Result :=  AObj.Item[Index];
end;

procedure TreeNode_SetItem(AObj: TTreeNode; Index: Integer; AValue: TTreeNode); stdcall;
begin
  AObj.Item[Index] := AValue;
end;


exports
  TreeNode_Create,
  TreeNode_Free,
  TreeNode_AlphaSort,
  TreeNode_Assign,
  TreeNode_Collapse,
  TreeNode_Delete,
  TreeNode_DisplayRect,
  TreeNode_EditText,
  TreeNode_Expand,
  TreeNode_IndexOf,
  TreeNode_MakeVisible,
  TreeNode_MoveTo,
  TreeNode_GetNamePath,
  TreeNode_ClassName,
  TreeNode_Equals,
  TreeNode_GetHashCode,
  TreeNode_ToString,
  TreeNode_GetAbsoluteIndex,
  TreeNode_GetCount,
  TreeNode_GetCut,
  TreeNode_SetCut,
  TreeNode_GetData,
  TreeNode_SetData,
  TreeNode_GetDeleting,
  TreeNode_GetFocused,
  TreeNode_SetFocused,
  TreeNode_GetDropTarget,
  TreeNode_SetDropTarget,
  TreeNode_GetSelected,
  TreeNode_SetSelected,
  TreeNode_GetExpanded,
  TreeNode_SetExpanded,
  TreeNode_GetExpandedImageIndex,
  TreeNode_SetExpandedImageIndex,
  TreeNode_GetHandle,
  TreeNode_GetHasChildren,
  TreeNode_SetHasChildren,
  TreeNode_GetImageIndex,
  TreeNode_SetImageIndex,
  TreeNode_GetIndex,
  TreeNode_GetIsVisible,
  TreeNode_GetItemId,
  TreeNode_GetLevel,
  TreeNode_GetOverlayIndex,
  TreeNode_SetOverlayIndex,
  TreeNode_GetOwner,
  TreeNode_GetParent,
  TreeNode_GetSelectedIndex,
  TreeNode_SetSelectedIndex,
  TreeNode_GetEnabled,
  TreeNode_SetEnabled,
  TreeNode_GetStateIndex,
  TreeNode_SetStateIndex,
  TreeNode_GetText,
  TreeNode_SetText,
  TreeNode_GetItem,
  TreeNode_SetItem;

