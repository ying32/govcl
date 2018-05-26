
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function TreeNode_Create(AOwner: TTreeNodes): TTreeNode; extdecl;
begin
  Result :=  TTreeNode.Create(AOwner);
end;

procedure TreeNode_Free(AObj: TTreeNode); extdecl;
begin
  AObj.Free;
end;

procedure TreeNode_Assign(AObj: TTreeNode; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

procedure TreeNode_Collapse(AObj: TTreeNode; Recurse: LongBool); extdecl;
begin
  AObj.Collapse(Recurse);
end;

procedure TreeNode_Delete(AObj: TTreeNode); extdecl;
begin
  AObj.Delete;
end;

procedure TreeNode_DisplayRect(AObj: TTreeNode; TextOnly: LongBool; var Result: TRect); extdecl;
begin
  Result :=  AObj.DisplayRect(TextOnly);
end;

function TreeNode_EditText(AObj: TTreeNode): LongBool; extdecl;
begin
  Result :=  AObj.EditText;
end;

procedure TreeNode_Expand(AObj: TTreeNode; Recurse: LongBool); extdecl;
begin
  AObj.Expand(Recurse);
end;

function TreeNode_IndexOf(AObj: TTreeNode; Value: TTreeNode): Integer; extdecl;
begin
  Result :=  AObj.IndexOf(Value);
end;

procedure TreeNode_MakeVisible(AObj: TTreeNode); extdecl;
begin
  AObj.MakeVisible;
end;

procedure TreeNode_MoveTo(AObj: TTreeNode; Destination: TTreeNode; Mode: TNodeAttachMode); extdecl;
begin
  AObj.MoveTo(Destination, Mode);
end;

function TreeNode_GetNamePath(AObj: TTreeNode): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function TreeNode_ClassName(AObj: TTreeNode): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function TreeNode_Equals(AObj: TTreeNode; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function TreeNode_GetHashCode(AObj: TTreeNode): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function TreeNode_ToString(AObj: TTreeNode): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function TreeNode_GetAbsoluteIndex(AObj: TTreeNode): Integer; extdecl;
begin
  Result :=  AObj.AbsoluteIndex;
end;

function TreeNode_GetCount(AObj: TTreeNode): Integer; extdecl;
begin
  Result :=  AObj.Count;
end;

function TreeNode_GetCut(AObj: TTreeNode): LongBool; extdecl;
begin
  Result :=  AObj.Cut;
end;

procedure TreeNode_SetCut(AObj: TTreeNode; AValue: LongBool); extdecl;
begin
  AObj.Cut := AValue;
end;

function TreeNode_GetData(AObj: TTreeNode): Pointer; extdecl;
begin
  Result :=  AObj.Data;
end;

procedure TreeNode_SetData(AObj: TTreeNode; AValue: Pointer); extdecl;
begin
  AObj.Data := AValue;
end;

function TreeNode_GetDeleting(AObj: TTreeNode): LongBool; extdecl;
begin
  Result :=  AObj.Deleting;
end;

function TreeNode_GetFocused(AObj: TTreeNode): LongBool; extdecl;
begin
  Result :=  AObj.Focused;
end;

procedure TreeNode_SetFocused(AObj: TTreeNode; AValue: LongBool); extdecl;
begin
  AObj.Focused := AValue;
end;

function TreeNode_GetSelected(AObj: TTreeNode): LongBool; extdecl;
begin
  Result :=  AObj.Selected;
end;

procedure TreeNode_SetSelected(AObj: TTreeNode; AValue: LongBool); extdecl;
begin
  AObj.Selected := AValue;
end;

function TreeNode_GetExpanded(AObj: TTreeNode): LongBool; extdecl;
begin
  Result :=  AObj.Expanded;
end;

procedure TreeNode_SetExpanded(AObj: TTreeNode; AValue: LongBool); extdecl;
begin
  AObj.Expanded := AValue;
end;

function TreeNode_GetHandle(AObj: TTreeNode): HWND; extdecl;
begin
  Result :=  AObj.Handle;
end;

function TreeNode_GetHasChildren(AObj: TTreeNode): LongBool; extdecl;
begin
  Result :=  AObj.HasChildren;
end;

procedure TreeNode_SetHasChildren(AObj: TTreeNode; AValue: LongBool); extdecl;
begin
  AObj.HasChildren := AValue;
end;

function TreeNode_GetImageIndex(AObj: TTreeNode): TImageIndex; extdecl;
begin
  Result :=  AObj.ImageIndex;
end;

procedure TreeNode_SetImageIndex(AObj: TTreeNode; AValue: TImageIndex); extdecl;
begin
  AObj.ImageIndex := AValue;
end;

function TreeNode_GetIndex(AObj: TTreeNode): Integer; extdecl;
begin
  Result :=  AObj.Index;
end;

function TreeNode_GetIsVisible(AObj: TTreeNode): LongBool; extdecl;
begin
  Result :=  AObj.IsVisible;
end;

function TreeNode_GetLevel(AObj: TTreeNode): Integer; extdecl;
begin
  Result :=  AObj.Level;
end;

function TreeNode_GetOwner(AObj: TTreeNode): TTreeNodes; extdecl;
begin
  Result :=  AObj.Owner;
end;

function TreeNode_GetParent(AObj: TTreeNode): TTreeNode; extdecl;
begin
  Result :=  AObj.Parent;
end;

function TreeNode_GetSelectedIndex(AObj: TTreeNode): Integer; extdecl;
begin
  Result :=  AObj.SelectedIndex;
end;

procedure TreeNode_SetSelectedIndex(AObj: TTreeNode; AValue: Integer); extdecl;
begin
  AObj.SelectedIndex := AValue;
end;

function TreeNode_GetStateIndex(AObj: TTreeNode): Integer; extdecl;
begin
  Result :=  AObj.StateIndex;
end;

procedure TreeNode_SetStateIndex(AObj: TTreeNode; AValue: Integer); extdecl;
begin
  AObj.StateIndex := AValue;
end;

function TreeNode_GetText(AObj: TTreeNode): PChar; extdecl;
begin
  Result :=  PChar(AObj.Text);
end;

procedure TreeNode_SetText(AObj: TTreeNode; AValue: PChar); extdecl;
begin
  AObj.Text := AValue;
end;

function TreeNode_GetItem(AObj: TTreeNode; Index: Integer): TTreeNode; extdecl;
begin
  Result :=  AObj.Items[Index];
end;

procedure TreeNode_SetItem(AObj: TTreeNode; Index: Integer; AValue: TTreeNode); extdecl;
begin
  AObj.Items[Index] := AValue;
end;


exports
  TreeNode_Create,
  TreeNode_Free,
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
  TreeNode_GetSelected,
  TreeNode_SetSelected,
  TreeNode_GetExpanded,
  TreeNode_SetExpanded,
  TreeNode_GetHandle,
  TreeNode_GetHasChildren,
  TreeNode_SetHasChildren,
  TreeNode_GetImageIndex,
  TreeNode_SetImageIndex,
  TreeNode_GetIndex,
  TreeNode_GetIsVisible,
  TreeNode_GetLevel,
  TreeNode_GetOwner,
  TreeNode_GetParent,
  TreeNode_GetSelectedIndex,
  TreeNode_SetSelectedIndex,
  TreeNode_GetStateIndex,
  TreeNode_SetStateIndex,
  TreeNode_GetText,
  TreeNode_SetText,
  TreeNode_GetItem,
  TreeNode_SetItem;

