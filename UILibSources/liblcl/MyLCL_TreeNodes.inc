
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function TreeNodes_Create(AOwner: TCustomTreeView): TTreeNodes; extdecl;
begin
  Result :=  TTreeNodes.Create(AOwner);
end;

procedure TreeNodes_Free(AObj: TTreeNodes); extdecl;
begin
  AObj.Free;
end;

function TreeNodes_AddChildFirst(AObj: TTreeNodes; Parent: TTreeNode; S: PChar): TTreeNode; extdecl;
begin
  Result :=  AObj.AddChildFirst(Parent, S);
end;

function TreeNodes_AddChild(AObj: TTreeNodes; Parent: TTreeNode; S: PChar): TTreeNode; extdecl;
begin
  Result :=  AObj.AddChild(Parent, S);
end;

function TreeNodes_AddChildObjectFirst(AObj: TTreeNodes; Parent: TTreeNode; S: PChar; Ptr: Pointer): TTreeNode; extdecl;
begin
  Result :=  AObj.AddChildObjectFirst(Parent, S, Ptr);
end;

function TreeNodes_AddChildObject(AObj: TTreeNodes; Parent: TTreeNode; S: PChar; Ptr: Pointer): TTreeNode; extdecl;
begin
  Result :=  AObj.AddChildObject(Parent, S, Ptr);
end;

function TreeNodes_AddObjectFirst(AObj: TTreeNodes; Sibling: TTreeNode; S: PChar; Ptr: Pointer): TTreeNode; extdecl;
begin
  Result :=  AObj.AddObjectFirst(Sibling, S, Ptr);
end;

function TreeNodes_AddObject(AObj: TTreeNodes; Sibling: TTreeNode; S: PChar; Ptr: Pointer): TTreeNode; extdecl;
begin
  Result :=  AObj.AddObject(Sibling, S, Ptr);
end;

function TreeNodes_AddNode(AObj: TTreeNodes; Node: TTreeNode; Relative: TTreeNode; S: PChar; Ptr: Pointer; Method: TNodeAttachMode): TTreeNode; extdecl;
begin
  Result :=  AObj.AddNode(Node, Relative, S, Ptr, Method);
end;

function TreeNodes_AddFirst(AObj: TTreeNodes; Sibling: TTreeNode; S: PChar): TTreeNode; extdecl;
begin
  Result :=  AObj.AddFirst(Sibling, S);
end;

function TreeNodes_Add(AObj: TTreeNodes; Sibling: TTreeNode; S: PChar): TTreeNode; extdecl;
begin
  Result :=  AObj.Add(Sibling, S);
end;

procedure TreeNodes_Assign(AObj: TTreeNodes; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

procedure TreeNodes_BeginUpdate(AObj: TTreeNodes); extdecl;
begin
  AObj.BeginUpdate;
end;

procedure TreeNodes_Clear(AObj: TTreeNodes); extdecl;
begin
  AObj.Clear;
end;

procedure TreeNodes_Delete(AObj: TTreeNodes; Node: TTreeNode); extdecl;
begin
  AObj.Delete(Node);
end;

procedure TreeNodes_EndUpdate(AObj: TTreeNodes); extdecl;
begin
  AObj.EndUpdate;
end;

function TreeNodes_GetFirstNode(AObj: TTreeNodes): TTreeNode; extdecl;
begin
  Result :=  AObj.GetFirstNode;
end;

function TreeNodes_Insert(AObj: TTreeNodes; Sibling: TTreeNode; S: PChar): TTreeNode; extdecl;
begin
  Result :=  AObj.Insert(Sibling, S);
end;

function TreeNodes_InsertObject(AObj: TTreeNodes; Sibling: TTreeNode; S: PChar; Ptr: Pointer): TTreeNode; extdecl;
begin
  Result :=  AObj.InsertObject(Sibling, S, Ptr);
end;

function TreeNodes_GetNamePath(AObj: TTreeNodes): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function TreeNodes_ClassName(AObj: TTreeNodes): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function TreeNodes_Equals(AObj: TTreeNodes; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function TreeNodes_GetHashCode(AObj: TTreeNodes): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function TreeNodes_ToString(AObj: TTreeNodes): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function TreeNodes_GetCount(AObj: TTreeNodes): Integer; extdecl;
begin
  Result :=  AObj.Count;
end;

function TreeNodes_GetOwner(AObj: TTreeNodes): TCustomTreeView; extdecl;
begin
  Result :=  AObj.Owner;
end;

function TreeNodes_GetItem(AObj: TTreeNodes; Index: Integer): TTreeNode; extdecl;
begin
  Result :=  AObj.Item[Index];
end;


exports
  TreeNodes_Create,
  TreeNodes_Free,
  TreeNodes_AddChildFirst,
  TreeNodes_AddChild,
  TreeNodes_AddChildObjectFirst,
  TreeNodes_AddChildObject,
  TreeNodes_AddObjectFirst,
  TreeNodes_AddObject,
  TreeNodes_AddNode,
  TreeNodes_AddFirst,
  TreeNodes_Add,
  TreeNodes_Assign,
  TreeNodes_BeginUpdate,
  TreeNodes_Clear,
  TreeNodes_Delete,
  TreeNodes_EndUpdate,
  TreeNodes_GetFirstNode,
  TreeNodes_Insert,
  TreeNodes_InsertObject,
  TreeNodes_GetNamePath,
  TreeNodes_ClassName,
  TreeNodes_Equals,
  TreeNodes_GetHashCode,
  TreeNodes_ToString,
  TreeNodes_GetCount,
  TreeNodes_GetOwner,
  TreeNodes_GetItem;

