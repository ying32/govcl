{%MainUnit ../comctrls.pp}

{******************************************************************************
                                  TTreeView
 ******************************************************************************

  Author: Mattias Gaertner

 *****************************************************************************
  This file is part of the Lazarus Component Library (LCL)

  See the file COPYING.modifiedLGPL.txt, included in this distribution,
  for details about the license.
 *****************************************************************************

  Abstract:
    TTreeView for LCL
}


{ $DEFINE TREEVIEW_DEBUG}

const
  TTreeNodeWithPointerStreamVersion : word = 1;
  TTreeNodeStreamVersion : word = 2;
  TVAutoHeightString = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789jgpq|\()^';
  MinNodeCapacity = 10;

function CompareExpandedNodes(Data1, Data2: Pointer): integer;
var
  Node1: TTreeNodeExpandedState;
  Node2: TTreeNodeExpandedState;
begin
  Node1:=TTreeNodeExpandedState(Data1);
  Node2:=TTreeNodeExpandedState(Data2);
  Result:=AnsiCompareText(Node1.NodeText,Node2.NodeText);
end;

function CompareTextWithExpandedNode(Key, Data: Pointer): integer;
var
  NodeText: String;
  Node: TTreeNodeExpandedState;
begin
  NodeText:=String(Key);
  Node:=TTreeNodeExpandedState(Data);
  Result:=AnsiCompareText(NodeText,Node.NodeText);
end;

procedure TreeViewError(const Msg: string);
begin
  raise ETreeViewError.Create(Msg);
end;

procedure TreeNodeError(const Msg: string);
begin
  raise ETreeNodeError.Create(Msg);
end;

procedure TreeNodeErrorFmt(const Msg: string; Format: array of const);
begin
  raise ETreeNodeError.CreateFmt(Msg, Format);
end;

function IndexOfNodeAtTop(NodeArray: TTreeNodeArray; Count, y: integer): integer;
// NodeArray must be sorted via Top
// returns index of Node with Node.Top <= y < Node[+1].Top
var
  l, m, r, VisibleCount: integer;
  VisibleNodesAlloc: Boolean;
  VisibleNodes: TTreeNodeArray;
begin
  if (Count = 0) or (NodeArray = nil) then
    Exit(-1);
  // Count the visible nodes
  VisibleCount := 0;
  VisibleNodesAlloc := False;
  for l := 0 to Count-1 do
    if NodeArray[l].Visible then
      Inc(VisibleCount);
  try
    // Make a temporary array of visible nodes if there are hidden nodes
    if VisibleCount < Count then begin
      GetMem(VisibleNodes,SizeOf(Pointer)*VisibleCount);
      m := 0;
      for l := 0 to Count-1 do
        if NodeArray[l].Visible then begin
          VisibleNodes[m] := NodeArray[l];
          Inc(m);
        end;
      Count := VisibleCount;
      VisibleNodesAlloc := True;
    end
    else
      VisibleNodes := NodeArray;
    // Binary search for the Y coordinate
    l := 0;
    r := Count - 1;
    while (l <= r) do
    begin
      m := (l + r) shr 1;
      //DebugLn(':0 [IndexOfNodeAtTop] m=',m,' y=',y,' ',NodeArray[m].Text,' NodeArray[m].Top=',NodeArray[m].Top,' NodeArray[m].BottomExpanded=',NodeArray[m].BottomExpanded);
      if VisibleNodes[m].Top > y then
        r := m - 1
      else if VisibleNodes[m].BottomExpanded <= y then
        l := m + 1
      else
        Exit(VisibleNodes[m].Index);
    end;
    Result := -1;
  finally
    if VisibleNodesAlloc then
      Freemem(VisibleNodes);
  end;
end;

// procedure for sorting a TTreeNodeArray
procedure Sort(Nodes: TTreeNodeArray; Count: integer;
  SortProc: TTreeNodeCompare; UpdateIndex: Boolean);
// Sorts the nodes using merge sort and updates the sibling links
var
  Buffer: TTreeNodeArray;
  i: Integer;

  procedure MergeNodeArrays(Pos1, Pos2, Pos3: integer);
  // merge two sorted arrays (result is in Src)
  // the first array ranges Pos1..Pos2-1, the second ranges Pos2..Pos3
  var Src1Pos,Src2Pos,DestPos,cmp,a:integer;
  begin
    if (Pos1>=Pos2) or (Pos2>Pos3) then exit;
    Src1Pos:=Pos2-1;
    Src2Pos:=Pos3;
    DestPos:=Pos3;
    while (Src2Pos>=Pos2) and (Src1Pos>=Pos1) do begin
      cmp:=SortProc(Nodes[Src1Pos],Nodes[Src2Pos]);
      if cmp>0 then begin
        Buffer[DestPos]:=Nodes[Src1Pos];
        dec(Src1Pos);
      end else begin
        Buffer[DestPos]:=Nodes[Src2Pos];
        dec(Src2Pos);
      end;
      dec(DestPos);
    end;
    while Src2Pos>=Pos2 do begin
      Buffer[DestPos]:=Nodes[Src2Pos];
      dec(Src2Pos);
      dec(DestPos);
    end;
    for a:=DestPos+1 to Pos3 do
      Nodes[a]:=Buffer[a];
  end;

  procedure MergeSort(StartPos, EndPos: integer);
  // sort Src from Position StartPos to EndPos (both included)
  var cmp,mid:integer;
  begin
    if StartPos>=EndPos then begin
      // sort one element -> very easy :)
    end else if StartPos+1=EndPos then begin
      // sort two elements -> quite easy :)
      cmp:=SortProc(Nodes[StartPos],Nodes[EndPos]);
      if cmp>0 then begin
        Buffer[StartPos]:=Nodes[StartPos];
        Nodes[StartPos]:=Nodes[EndPos];
        Nodes[EndPos]:=Buffer[StartPos];
      end;
    end else begin
      // sort more than two elements -> Mergesort
      mid:=(StartPos+EndPos) shr 1;
      MergeSort(StartPos,mid);
      MergeSort(mid+1,EndPos);
      MergeNodeArrays(StartPos,mid+1,EndPos);
    end;
  end;

begin
  if Count>0 then begin
    Buffer := GetMem(SizeOf(Pointer)*Count);
    MergeSort(0,Count-1);
    FreeMem(Buffer);
    // update sibling links
    Nodes[0].FPrevBrother := nil;
    Nodes[Count-1].FNextBrother := nil;
    if UpdateIndex then Nodes[0].FIndex:=0;
    for i:= 1 to Count-1 do begin
      Nodes[i].FPrevBrother := Nodes[i-1];
      Nodes[i-1].FNextBrother := Nodes[i];
      if UpdateIndex then Nodes[i].FIndex:=i;
    end;
  end;
end;

{ TTreeNodeExpandedState }

function TTreeNodeExpandedState.DefaultGetNodeText(Node: TTreeNode): string;
begin
  Result:=Node.Text;
end;

constructor TTreeNodeExpandedState.Create(FirstTreeNode: TTreeNode;
  const GetNodeTextEvent: TTVGetNodeText);
begin
  if GetNodeTextEvent<>nil then
    FOnGetNodeText:=GetNodeTextEvent
  else
    FOnGetNodeText:=@DefaultGetNodeText;
  CreateChildNodes(FirstTreeNode);
end;

constructor TTreeNodeExpandedState.Create(TreeView: TCustomTreeView;
  const GetNodeTextEvent: TTVGetNodeText);
begin
  if GetNodeTextEvent<>nil then
    FOnGetNodeText:=GetNodeTextEvent
  else
    FOnGetNodeText:=@DefaultGetNodeText;
  CreateChildNodes(TreeView.Items.GetFirstNode);
end;

destructor TTreeNodeExpandedState.Destroy;
begin
  Clear;
  inherited Destroy;
end;

procedure TTreeNodeExpandedState.Clear;
begin
  if Children<>nil then begin
    Children.FreeAndClear;
    FreeThenNil(Children);
  end;
end;

procedure TTreeNodeExpandedState.CreateChildNodes(FirstTreeNode: TTreeNode);
var
  ChildNode: TTreeNode;
  NewExpandedNode: TTreeNodeExpandedState;
begin
  Clear;
  if (FirstTreeNode=nil) then exit;
  if (FirstTreeNode.Parent<>nil) then
    NodeText:=OnGetNodeText(FirstTreeNode.Parent)
  else
    NodeText:='';
  ChildNode:=FirstTreeNode;
  while ChildNode<>nil do begin
    if ChildNode.Expanded then begin
      if Children=nil then Children:=TAvlTree.Create(@CompareExpandedNodes);
      NewExpandedNode:=TTreeNodeExpandedState.Create(ChildNode.GetFirstChild,OnGetNodeText);
      if ChildNode.GetFirstChild=nil then
        NewExpandedNode.NodeText:=OnGetNodeText(ChildNode);
      Children.Add(NewExpandedNode);
    end;
    ChildNode:=ChildNode.GetNextSibling;
  end;
end;

procedure TTreeNodeExpandedState.Apply(FirstTreeNode: TTreeNode; CollapseToo: boolean);
var
  ChildNode: TTreeNode;
  ANode: TAvlTreeNode;
  ChildNodeText: String;
begin
  if Children=nil then exit;
  ChildNode:=FirstTreeNode;
  while ChildNode<>nil do begin
    ChildNodeText:=OnGetNodeText(ChildNode);
    ANode:=Children.FindKey(PChar(ChildNodeText),@CompareTextWithExpandedNode);
    if ChildNodeText='' then
      debugln(['TTreeNodeExpandedState.Apply ',ChildNode.GetTextPath,' ChildNodeText="',ChildNodeText,'"']);
    if ANode<>nil then
      ChildNode.Expanded:=true
    else if CollapseToo then
      ChildNode.Expanded:=false;
    if ANode<>nil then
      TTreeNodeExpandedState(ANode.Data).Apply(ChildNode.GetFirstChild,CollapseToo);
    ChildNode:=ChildNode.GetNextSibling;
  end;
end;

procedure TTreeNodeExpandedState.Apply(TreeView: TCustomTreeView; CollapseToo: boolean);
begin
  Apply(TreeView.Items.GetFirstNode,CollapseToo);
end;

{ TTreeNode }

function TTreeNode.DefaultTreeViewSort(Node1, Node2: TTreeNode): Integer;
begin
  if (Node1.TreeView<>nil) and  Assigned(Node1.TreeView.OnCompare) then begin
    Result:=0;
    Node1.TreeView.OnCompare(Node1.TreeView,Node1, Node2, Result);
  end else
    Result := AnsiCompareStr(Node1.Text,Node2.Text);
end;

constructor TTreeNode.Create(AnOwner: TTreeNodes);
begin
  inherited Create;
  FNodeEffect := gdeNormal;
  FImageIndex := -1;
  FOverlayIndex := -1;
  FSelectedIndex := -1;
  FStateIndex := -1;
  FStates := [nsVisible];
  FOwner := AnOwner;
  FSubTreeCount := 1;
  FIndex := -1;
end;

destructor TTreeNode.Destroy;
var
  lOwnerAccessibleObject, lAccessibleObject: TLazAccessibleObject;
begin
  {$IFDEF TREEVIEW_DEBUG}
  DebugLn('[TTreeNode.Destroy] Self=',DbgS(Self),' Self.Text=',Text);
  {$ENDIF}
  Include(FStates,nsDeleting);

  // we must trigger TCustomTreeView.OnDeletion event before
  // unbinding. See issue #17832.
  if Assigned(Owner) and Assigned(Owner.Owner) then
    Owner.Owner.Delete(Self);

  // Remove the accessibility object too
  if Assigned(Owner) and Assigned(Owner.Owner) then
  begin
    lOwnerAccessibleObject := Owner.Owner.GetAccessibleObject();
    if lOwnerAccessibleObject<>nil then
    begin
      lAccessibleObject := lOwnerAccessibleObject.GetChildAccessibleObjectWithDataObject(Self);
      if lAccessibleObject <> nil then
        lOwnerAccessibleObject.RemoveChildAccessibleObject(lAccessibleObject);
    end;
  end;

  // delete children
  HasChildren := false;
  // unbind all references
  Unbind;
  if Assigned(Owner) then begin
    if FStates * [nsSelected, nsMultiSelected] <> [] then
      Owner.FSelection.Remove(Self);
  end;

  Data := nil;
  // free data
  if Assigned(FItems) then
  begin
    FreeMem(FItems);
    FItems := nil;
  end;
  inherited Destroy;
end;

function TTreeNode.GetHandle: THandle;
begin
  if TreeView<>nil then
    Result := TreeView.Handle
  else
    Result := 0;
end;

function TTreeNode.GetParentNodeOfAbsoluteLevel(
  TheAbsoluteLevel: integer): TTreeNode;
var
  i: integer;
  l: LongInt;
begin
  l:=Level;
  if (TheAbsoluteLevel > l) or (TheAbsoluteLevel < 0) then
    Result := nil
  else
  begin
    Result := Self;
    for i := TheAbsoluteLevel to l-1 do
      Result := Result.Parent;
  end;
end;

function TTreeNode.GetTreeNodes: TTreeNodes;
begin
  if (Owner<>nil) and (Owner is TTreeNodes) then
    Result:=TTreeNodes(Owner)
  else
    Result:=nil;
end;

function TTreeNode.GetTreeView: TCustomTreeView;
begin
  if Owner <> nil then
    Result := Owner.Owner
  else
    Result := nil;
end;

function TTreeNode.GetTop: integer;
begin
  if TreeView <> nil then
    TreeView.UpdateAllTops;
  Result := FTop;
end;

function TTreeNode.GetVisible: Boolean;
begin
  Result:=nsVisible in FStates;
end;

function TTreeNode.HasAsParent(AValue: TTreeNode): Boolean;
begin
  if AValue<>nil then begin
    if Parent=nil then Result := False
    else if Parent=AValue then Result := True
    else Result := Parent.HasAsParent(AValue);
  end
  else Result := True;
end;

procedure TTreeNode.SetText(const S: string);
var
  lSelfAX: TLazAccessibleObject;
begin
  if S=FText then exit;
  FText := S;
  if TreeView=nil then exit;
  Include(TreeView.FStates,tvsMaxRightNeedsUpdate);
  if (TreeView.SortType in [stText, stBoth])
  and (nsBound in FStates) then begin
    if (Parent <> nil) then Parent.AlphaSort
    else TreeView.AlphaSort;
  end;
  Update;
  Changed(ncTextChanged);
  // Update accessibility information
  lSelfAX := TreeView.GetAccessibleObject.GetChildAccessibleObjectWithDataObject(Self);
  if lSelfAX <> nil then
    lSelfAX.AccessibleValue := S;
end;

procedure TTreeNode.SetData(AValue: Pointer);
begin
  if FData=AValue then exit;
  FData := AValue;
  if (TreeView<>nil)
  and (TreeView.SortType in [stData, stBoth]) and Assigned(TreeView.OnCompare)
  and (not Deleting)
  and (nsBound in FStates) then
  begin
    if Parent <> nil then
      Parent.AlphaSort
    else
      TreeView.AlphaSort;
  end;
  Changed(ncDataChanged);
end;

function TTreeNode.GetState(NodeState: TNodeState): Boolean;
begin
  Result:=NodeState in FStates;
end;

procedure TTreeNode.SetHeight(AValue: integer);
begin
  if AValue<0 then AValue:=0;
  if AValue=FHeight then exit;
  FHeight:=AValue;
  if TreeView<>nil then
    TreeView.FStates:=TreeView.FStates+[tvsScrollbarChanged,tvsTopsNeedsUpdate];
  Update;
  Changed(ncHeightChanged);
end;

procedure TTreeNode.SetImageEffect(AValue: TGraphicsDrawEffect);
begin
  if FNodeEffect=AValue then exit;
  FNodeEffect := AValue;
  Update;
  Changed(ncImageEffect);
end;

procedure TTreeNode.SetImageIndex(AValue: TImageIndex);
begin
  if FImageIndex=AValue then exit;
  FImageIndex := AValue;
  Update;
  Changed(ncImageIndex);
end;

procedure TTreeNode.SetIndex(const AValue: Integer);

  procedure RaiseMissingTreeNodes;
  begin
    TreeViewError('TTreeNode.SetIndex missing Owner');
  end;

var
  OldIndex: LongInt;
begin
  OldIndex:=Index;
  if OldIndex=AValue then exit;
  if Parent=nil then begin
    // moving a top level node
    if Owner=nil then RaiseMissingTreeNodes;
    if AValue=0 then
      MoveTo(Owner.GetFirstNode,naInsert)
    else if AValue<OldIndex then
      MoveTo(Owner.TopLvlItems[AValue-1],naInsertBehind)
    else
      MoveTo(Owner.TopLvlItems[AValue],naInsertBehind);
  end else begin
    // moving a normal node
    if AValue=0 then
      MoveTo(Parent.GetFirstChild,naInsert)
    else if AValue<OldIndex then
      MoveTo(Parent.Items[AValue-1],naInsertBehind)
    else
      MoveTo(Parent.Items[AValue],naInsertBehind);
  end;
end;

procedure TTreeNode.SetSelectedIndex(AValue: Integer);
begin
  if FSelectedIndex = AValue then exit;
  FSelectedIndex := AValue;
  Update;
  Changed(ncSelectedIndex);
end;

procedure TTreeNode.SetVisible(const AValue: Boolean);
begin
  if Visible = AValue then exit;
  if AValue then
    Include(FStates,nsVisible)
  else
    Exclude(FStates,nsVisible);
  Selected := False;
  if TreeView<>nil then
    TreeView.FStates:=TreeView.FStates+[tvsScrollbarChanged,tvsTopsNeedsUpdate,
                                        tvsTopItemNeedsUpdate,tvsBottomItemNeedsUpdate,
                                        tvsMaxLvlNeedsUpdate,tvsMaxRightNeedsUpdate];
  Update;
  Changed(ncVisibility);
end;

procedure TTreeNode.SetOverlayIndex(AValue: Integer);
begin
  if FOverlayIndex = AValue then exit;
  FOverlayIndex := AValue;
  Update;
  Changed(ncOverlayIndex)
end;

procedure TTreeNode.SetStateIndex(AValue: Integer);
begin
  if FStateIndex = AValue then exit;
  FStateIndex := AValue;
  Update;
  Changed(ncStateIndex);
end;

function TTreeNode.AreParentsExpandedAndVisible: Boolean;
var
  ANode: TTreeNode;
begin
  Result:=false;
  ANode:=Parent;
  while ANode<>nil do begin
    if not (ANode.Expanded and ANode.Visible) then exit;
    ANode:=ANode.Parent;
  end;
  Result:=true;
end;

procedure TTreeNode.BindToMultiSelected;
var
  TheTreeNodes: TTreeNodes;
  CurNode: TTreeNode;
begin
  TheTreeNodes:=TreeNodes;
  if TheTreeNodes=nil then exit;

  // Get the first selected node of the tree
  CurNode := TheTreeNodes.FFirstMultiSelected;

  // Initialize self unbinded
  Self.FPrevMultiSelected := nil;
  Self.FNextMultiSelected := nil;

  // If there isn't any selected node, set self as first
  if CurNode = nil then
  begin
    TheTreeNodes.FFirstMultiSelected := Self;
    TheTreeNodes.FStartMultiSelected := Self;
  end
  else
  begin

    // if last selected node was the previous one
    if (TheTreeNodes.FLastMultiSelected.AbsoluteIndex+1=Self.AbsoluteIndex) and (TheTreeNodes.FLastMultiSelected.FNextMultiSelected=nil) then
    begin
      TheTreeNodes.FLastMultiSelected.FNextMultiSelected := Self;
      Self.FPrevMultiSelected := TheTreeNodes.FLastMultiSelected;
    end
    else
    begin
      // if last selected node was the next one
      if (TheTreeNodes.FLastMultiSelected.AbsoluteIndex=Self.AbsoluteIndex+1) and (TheTreeNodes.FLastMultiSelected.FPrevMultiSelected=nil) then
      begin
        TheTreeNodes.FLastMultiSelected.FPrevMultiSelected := Self;
        Self.FNextMultiSelected := TheTreeNodes.FLastMultiSelected;
        TheTreeNodes.FFirstMultiSelected := Self
      end
      else
      begin

        // Scan linked list of selected nodes until one has a lower absolute index or we reach the end
        While (CurNode.GetNextMultiSelected<>Nil) and (CurNode.AbsoluteIndex<Self.AbsoluteIndex) do
          CurNode := CurNode.GetNextMultiSelected;

        // last of the list
        if CurNode.AbsoluteIndex < Self.AbsoluteIndex then
        begin
           CurNode.FNextMultiSelected := Self;
           Self.FPrevMultiSelected    := CurNode;
        end
        else
        // insert between two nodes
        begin
           Self.FPrevMultiSelected    := CurNode.FPrevMultiSelected;
           Self.FNextMultiSelected    := CurNode;
           if CurNode.FPrevMultiSelected <> nil then
              CurNode.FPrevMultiSelected.FNextMultiSelected := Self;
           CurNode.FPrevMultiSelected := Self;
        end;
        // Set self as head of the list if needed
        if Self.FPrevMultiSelected = nil then
          TheTreeNodes.FFirstMultiSelected := Self;
      end;

    end;

  end;

  // Set self as last selected node
  TheTreeNodes.FLastMultiSelected := Self;
end;

function TTreeNode.CompareCount(CompareMe: Integer): Boolean;
Begin
  Result:=(CompareMe=Count);
end;

function TTreeNode.DoCanExpand(ExpandIt: Boolean): Boolean;
begin
  Result := False;
  if (TreeView<>nil) and HasChildren then begin
    if ExpandIt then
      Result := TreeView.CanExpand(Self)
    else
      Result := TreeView.CanCollapse(Self);
  end;
end;

procedure TTreeNode.DoExpand(ExpandIt: Boolean);
begin
//DebugLn('[TTreeNode.DoExpand] Self=',DbgS(Self),' Text=',Text,
//' HasChildren=',HasChildren,' ExpandIt=',ExpandIt,' Expanded=',Expanded);
  if HasChildren and (Expanded<>ExpandIt) then begin
    if ExpandIt then
      Include(FStates,nsExpanded)
    else begin
      Exclude(FStates,nsExpanded);
      if (not Owner.KeepCollapsedNodes) then begin
        while GetLastChild<>nil do
          GetLastChild.Free;
      end;
    end;
    if TreeView<>nil then begin
      TreeView.FStates:=(TreeView.FStates+[tvsTopsNeedsUpdate,
                tvsTopItemNeedsUpdate,tvsBottomItemNeedsUpdate,
                tvsScrollbarChanged,tvsMaxRightNeedsUpdate]);
      Update;
      if ExpandIt then
        TreeView.Expand(Self)
      else
        TreeView.Collapse(Self);
    end;
  end;
end;

procedure TTreeNode.ExpandItem(ExpandIt: Boolean; Recurse: Boolean);
var
  //Flag: Integer;
  ANode: TTreeNode;
begin
  if Recurse then begin
    ExpandItem(ExpandIt, False);
    ANode := GetFirstChild;
    while ANode<>nil do begin
      ANode.ExpandItem(ExpandIt, true);
      ANode := ANode.FNextBrother;
    end;
  end
  else begin
    if TreeView<>nil then
      Include(TreeView.FStates,tvsManualNotify);
    try
      if DoCanExpand(ExpandIt) then
        DoExpand(ExpandIt);
    finally
      if TreeView<>nil then
        Exclude(TreeView.FStates,tvsManualNotify);
    end;
  end;
end;

procedure TTreeNode.Expand(Recurse: Boolean);
begin
  ExpandItem(True, Recurse);
end;

procedure TTreeNode.ExpandParents;
var ANode: TTreeNode;
begin
  ANode:=Parent;
  while ANode<>nil do begin
    ANode.Expanded:=true;
    ANode:=ANode.Parent;
  end;
end;

procedure TTreeNode.FreeAllNodeData;
var
  i: Integer;
begin
  FreeAndNil(FData);
  for i:=0 to Count-1 do
    Items[i].FreeAllNodeData;
end;

procedure TTreeNode.Collapse(Recurse: Boolean);
begin
  ExpandItem(False, Recurse);
end;

function TTreeNode.GetExpanded: Boolean;
begin
  Result := GetState(nsExpanded);
end;

procedure TTreeNode.SetExpanded(AValue: Boolean);
begin
  if AValue=Expanded then exit;
  if AValue then
    Expand(False)
  else
    Collapse(False);
end;

function TTreeNode.GetSelected: Boolean;
begin
  Result := GetState(nsSelected);
  if (not Result) and (tvoAllowMultiSelect in TreeView.Options) then
    Result := GetState(nsMultiSelected);
end;

procedure TTreeNode.SetSelected(AValue: Boolean);
var
  MultiSelect: Boolean;
  TV: TCustomTreeView;
begin
  if AValue=GetSelected then exit;
  TV:=TreeView;
  if AValue then
  begin
    Include(FStates,nsSelected);
    if (TV<>nil) then
    begin
      TV.EndEditing(true);
      MultiSelect:=tvoAllowMultiselect in TV.Options;

      if not MultiSelect and Assigned(FOwner) then
        FOwner.SelectionsChanged(Self, True);

      if MultiSelect then TV.LockSelectionChangeEvent;
      try
        TV.Selected:=Self;
        if TV.Selected<>Self then
          Exclude(FStates,nsSelected);
        if (nsSelected in FStates) and MultiSelect then
          MultiSelected:=true;
      finally
        if MultiSelect then TV.UnlockSelectionChangeEvent;
      end;
    end;
  end else
  begin
    if not MultiSelected and Assigned(FOwner) then
      FOwner.SelectionsChanged(Self, False);
    Exclude(FStates,nsSelected);
    if (TV<>nil) and (TV.Selected=Self) then
    begin
      {$IFDEF TREEVIEW_DEBUG}
      DebugLn('TTreeNode.SetSelected: Removing selection from Node (but it does not work): ', Text);
      {$ENDIF}
//      TV.EndEditing(true);    // Done in TV.SetSelection
      TV.Selected:=nil;
      Assert(TV.Selected<>Self, 'Should not happen');
      //if TV.Selected=Self then
      //  Include(FStates,nsSelected);
    end;
  end;
  Update;
end;

function TTreeNode.GetCut: boolean;
begin
  Result := GetState(nsCut);
end;

procedure TTreeNode.SetCut(AValue: Boolean);
begin
  if AValue=Cut then exit;
  if AValue then
    Include(FStates,nsCut)
  else
    Exclude(FStates,nsCut);
end;

function TTreeNode.GetDropTarget: Boolean;
begin
  Result := GetState(nsDropHilited);
end;

procedure TTreeNode.SetDropTarget(AValue: Boolean);
begin
  if AValue then begin
    Include(FStates,nsDropHilited);
    if TreeView<>nil then
      TreeView.FLastDropTarget:=Self;
  end else begin
    Exclude(FStates,nsDropHilited);
    if TreeView<>nil then
      TreeView.FLastDropTarget:=nil;
  end;
end;

function TTreeNode.GetHasChildren: Boolean;
begin
  Result := GetState(nsHasChildren);
end;

procedure TTreeNode.SetFocused(AValue: Boolean);
begin
  if AValue=GetFocused then exit;
  if AValue then
    Include(FStates,nsFocused)
  else
    Exclude(FStates,nsFocused);
  Update;
end;

function TTreeNode.Bottom: integer;
begin
  Result := Top + Height;
end;

function TTreeNode.BottomExpanded: integer;
var
  Node: TTreeNode;
begin
  Node := GetNextVisibleSibling;
  if Node <> nil then
    Result := Node.Top
  else begin
    Node := GetLastVisibleChild;
    if Expanded and (Node <> nil) then
      Result := Node.BottomExpanded
    else
      Result := Bottom;
  end;
end;

function TTreeNode.GetFocused: Boolean;
begin
  Result := GetState(nsFocused);
end;

procedure TTreeNode.SetHasChildren(AValue: Boolean);
begin
  if AValue=HasChildren then exit;
  //DebugLn('[TTreeNode.SetHasChildren] Self=',DbgS(Self),' Self.Text=',Text,' AValue=',AValue);
  if AValue then
    Include(FStates,nsHasChildren)
  else begin
    while GetLastChild<>nil do
      GetLastChild.Free;
    Exclude(FStates,nsHasChildren)
  end;
  Update;
end;

function TTreeNode.GetNextSibling: TTreeNode;
begin
  Result := FNextBrother;
end;

function TTreeNode.GetPrevSibling: TTreeNode;
begin
  Result := FPrevBrother;
end;

function TTreeNode.GetNextVisible: TTreeNode;
begin
  if Visible and Expanded and (GetFirstVisibleChild<>nil) then
    Result:=GetFirstVisibleChild
  else begin
    Result:=Self;
    while (Result<>nil) and (Result.GetNextVisibleSibling=nil) do
      Result:=Result.Parent;
    if Result<>nil then
      Result:=Result.GetNextVisibleSibling;
  end;
  if (Result<>nil) and ( (not Result.Visible) or (not AreParentsExpandedAndVisible) ) then
    Result:=nil;
end;

function TTreeNode.GetNextVisibleSibling: TTreeNode;
begin
  Result := Self;
  repeat
    Result := Result.GetNextSibling;
  until ((Result=nil) or (Result.Visible));
  if (Result<>nil) and (not Result.Visible) then // Result := nil ... will be removed
    Assert(False,'TTreeNode.GetNextVisibleSibling: (Result<>nil) and (not Result.Visible)');
end;

function TTreeNode.GetPrevVisible: TTreeNode;
var
  ANode: TTreeNode;
begin
  Result:=GetPrevVisibleSibling;
  if Result <> nil then begin
    while Result.Visible and Result.Expanded do begin
      ANode:=Result.GetLastVisibleChild;
      if ANode=nil then break;
      Result:=ANode;
    end;
  end
  else
    Result := Parent;
  if (Result<>nil) and ( (not Result.Visible) or (not AreParentsExpandedAndVisible) ) then
    Result:=nil;
end;

function TTreeNode.GetPrevVisibleSibling: TTreeNode;
begin
  Result := Self;
  repeat
    Result := Result.GetPrevSibling;
  until ((Result=nil) or (Result.Visible));
  if (Result<>nil) and (not Result.Visible) then  // Result := nil ... will be removed
    Assert(False,'TTreeNode.GetPrevVisibleSibling: (Result<>nil) and (not Result.Visible)');
end;

function TTreeNode.GetPrevExpanded: TTreeNode;
var
  ANode: TTreeNode;
begin
  Result:=GetPrevVisibleSibling;
  if Result <> nil then begin
    while Result.Visible and Result.Expanded do begin
      ANode:=Result.GetLastVisibleChild;
      if ANode=nil then break;
      Result:=ANode;
    end;
  end
  else
    Result:=Parent;
end;

function TTreeNode.GetNextChild(AValue: TTreeNode): TTreeNode;
begin
  if AValue <> nil then
    Result := AValue.GetNextSibling
  else
    Result := nil;
end;

function TTreeNode.GetNextExpanded: TTreeNode;
var
  ANode: TTreeNode;
begin
  ANode := GetFirstVisibleChild;
  if Expanded and (ANode<>nil) then
    Result:=ANode
  else begin
    Result:=Self;
    while (Result<>nil) and (Result.GetNextVisibleSibling=nil) do
      Result:=Result.Parent;
    if Result<>nil then
      Result:=Result.GetNextVisibleSibling;
  end;
end;

function TTreeNode.GetNextMultiSelected: TTreeNode;
begin
  Result:=FNextMultiSelected;
end;

function TTreeNode.GetPrevChild(AValue: TTreeNode): TTreeNode;
begin
  if AValue <> nil then
    Result := AValue.GetPrevSibling
  else
    Result := nil;
end;

function TTreeNode.GetPrevMultiSelected: TTreeNode;
begin
  Result:=FPrevMultiSelected;
end;

function TTreeNode.GetFirstChild: TTreeNode;
begin
  if Count>0 then
    Result:=FItems[0]
  else
    Result:=nil;
end;

function TTreeNode.GetFirstVisibleChild: TTreeNode;
begin
  Result := GetFirstChild;
  if (Result<>nil) and (not Result.Visible) then
    Result := Result.GetNextVisibleSibling;
end;

function TTreeNode.GetLastSibling: TTreeNode;
begin
  if Parent<>nil then
    Result:=Parent.GetLastChild
  else begin
    Result:=Self;
    while Result.FNextBrother<>nil do
      Result:=Result.FNextBrother;
  end;
end;

function TTreeNode.GetLastChild: TTreeNode;
begin
  if Count > 0 then
    Result := FItems[Count - 1]
  else
    Result := nil;
end;

function TTreeNode.GetLastVisibleChild: TTreeNode;
begin
  Result := GetLastChild;
  if Assigned(Result) and not Result.Visible then begin
    Result := Result.GetPrevVisible;
    if Result = Self then begin                     // No visible nodes found.
      Assert(Visible, 'TTreeNode.GetLastVisibleChild: Node is not Visible');
      Result := Nil;
    end;
  end;
end;

function TTreeNode.GetLastSubChild: TTreeNode;
var
  Node: TTreeNode;
begin
  Result:=GetLastChild;
  if Result<>nil then begin
    Node:=Result.GetLastSubChild;
    if Node<>nil then
      Result:=Node;
  end;
end;

function TTreeNode.GetNext: TTreeNode;
begin
  Result:=GetFirstChild;
  if Result=nil then
    Result:=GetNextSkipChildren;
end;

function TTreeNode.GetNextSkipChildren: TTreeNode;
begin
  Result:=Self;
  while (Result<>nil) and (Result.FNextBrother=nil) do
    Result:=Result.Parent;
  if Result<>nil then
    Result:=Result.FNextBrother;
end;

function TTreeNode.GetPrev: TTreeNode;
var
  ANode: TTreeNode;
begin
  Result := GetPrevSibling;
  if Result <> nil then begin
    ANode := Result;
    repeat
      Result := ANode;
      ANode := Result.GetLastChild;
    until ANode = nil;
  end else
    Result := Parent;
end;

function TTreeNode.GetAbsoluteIndex: Integer;
// - first node has index 0
// - the first child of a node has an index one bigger than its parent
// - a node without children has an index one bigger than its previous brother
var
  ANode: TTreeNode;
begin
  Result:=-1;
  ANode:=Self;
  repeat
    inc(Result);
    while ANode.FPrevBrother<>nil do begin
      ANode:=ANode.FPrevBrother;
      inc(Result,ANode.FSubTreeCount);
    end;
    ANode:=ANode.Parent;
  until ANode=nil;
end;

function TTreeNode.GetDeleting: Boolean;
begin
  Result := nsDeleting in FStates;
end;

function TTreeNode.GetHeight: integer;
begin
  if FHeight <= 0 then
  begin
    if TreeView <> nil then
      Result := TreeView.FDefItemHeight
    else
      Result := DefaultTreeNodeHeight;
  end
  else
    Result := FHeight;
end;

function TTreeNode.GetIndex: Integer;
// returns number of previous siblings (nodes on same lvl with same parent)
var
  ANode: TTreeNode;
begin
  if self = nil then
    exit(-1);
  if FIndex>=0 then
    exit(FIndex);
  // many algorithms uses the last sibling, so we check that first for speed
  if (Parent<>nil) and (Parent[Parent.Count-1]=Self) then begin
    Result:=Parent.Count-1;
    FIndex:=Result;
    exit;
  end;
  // count previous siblings
  Result := -1;
  ANode := Self;
  while ANode <> nil do begin
    Inc(Result);
    if ANode.FIndex>=0 then begin
      inc(Result,ANode.FIndex);
      break;
    end;
    ANode := ANode.GetPrevSibling;
  end;
  FIndex:=Result;
end;

function TTreeNode.GetItems(AnIndex: Integer): TTreeNode;
begin
  if (AnIndex<0) or (AnIndex>=Count) then
    TreeNodeErrorFmt(rsIndexOutOfBounds,[ClassName, AnIndex, Count-1]);
  Result:=FItems[AnIndex];
end;

procedure TTreeNode.SetItems(AnIndex: Integer; AValue: TTreeNode);
begin
  if (AnIndex<0) or (AnIndex>=Count) then
    TreeNodeErrorFmt(rsIndexOutOfBounds, [ClassName, AnIndex, Count-1]);
  Items[AnIndex].Assign(AValue);
end;

procedure TTreeNode.SetMultiSelected(const AValue: Boolean);
begin
  if AValue=GetMultiSelected then exit;
  if AValue then begin
    if (Treeview<>nil) and (not (tvoAllowMultiselect in TreeView.Options)) then
      exit;
    if Assigned(FOwner) then
      FOwner.SelectionsChanged(Self, True);
    Include(FStates,nsMultiSelected);
    if TreeNodes<>nil then BindToMultiSelected;
  end else begin
    if Assigned(FOwner) then
      FOwner.SelectionsChanged(Self, False);
    Exclude(FStates,nsMultiSelected);
    if TreeNodes<>nil then UnbindFromMultiSelected;
  end;
  if TreeView<>nil then TreeView.InternalSelectionChanged;
  Update;
end;

function TTreeNode.IndexOf(AValue: TTreeNode): Integer;
begin
  if (AValue = nil) or (AValue.FParent <> Self) then
  begin
    Result:=-1;
    exit;
  end;
  Result := AValue.GetIndex;
end;

function TTreeNode.IndexOfText(const NodeText: string): Integer;
begin
  Result := Count - 1;
  while Result >= 0 do
  begin
    if FItems[Result].Text = NodeText then exit;
    dec(Result);
  end;
end;

function TTreeNode.FindNode(const NodeText: string): TTreeNode;
begin
  Result:=GetFirstChild;
  while (Result<>nil) and (Result.Text<>NodeText) do
    Result:=Result.GetNextSibling;
end;

function TTreeNode.GetTextPath: string;
var
  Node: TTreeNode;
begin
  Result := '';
  Node := Self;
  while Assigned(Node) do
  begin
    if Result <> '' then
      Result := '/' + Result;
    Result := Node.Text + Result;
    Node := Node.Parent;
  end;
end;

function TTreeNode.GetCount: Integer;
begin
  Result := FCount;
end;

procedure TTreeNode.EndEdit(Cancel: Boolean);
var
  TV: TCustomTreeView;
begin
  TV:=TreeView;
  if TV=nil then exit;
  TV.EndEditing(Cancel);
end;

procedure TTreeNode.Unbind;
// unbind from parent and neighbor siblings, but not from owner
var OldIndex, i: integer;
  HigherNode: TTreeNode;
  TheTreeView: TCustomTreeView;
begin
  {$IFDEF TREEVIEW_DEBUG}
  DebugLn('[TTreeNode.Unbind] Self=',DbgS(Self),' Self.Text=',Text);
  {$ENDIF}
  if not (nsBound in FStates) then exit;
  Exclude(FStates,nsBound);
  if Owner<>nil then dec(Owner.FCount);

  // remove single select
  Selected:=false;
  // invalidate caches of TreeView and if root item, remove from TreeView.Items
  if Owner<>nil then begin
    Owner.ClearCache;
    if FParent=nil then
      Owner.MoveTopLvlNode(Owner.IndexOfTopLvlItem(Self),-1,Self);
    TheTreeView:=Owner.Owner;
    if TheTreeView<>nil then begin
      TheTreeView.FStates:=TheTreeView.FStates+[tvsMaxRightNeedsUpdate,
        tvsTopsNeedsUpdate,tvsTopItemNeedsUpdate,tvsBottomItemNeedsUpdate];
      if TheTreeView.FLastDropTarget=Self then
        TheTreeView.FLastDropTarget:=nil;
      if TheTreeView.FInsertMarkNode=Self then
        TheTreeView.FInsertMarkNode:=nil;
    end;
  end;
  // unmultiselect (keeping MultiSelected, but leaving multiselection list)
  UnbindFromMultiSelected;
  // remove from sibling list
  if FPrevBrother<>nil then FPrevBrother.FNextBrother:=FNextBrother;
  if FNextBrother<>nil then FNextBrother.FPrevBrother:=FPrevBrother;
  FPrevBrother:=nil;
  FNextBrother:=nil;
  // remove from parent
  if FParent<>nil then begin
    // update all FSubTreeCount
    HigherNode:=FParent;
    while HigherNode<>nil do begin
      dec(HigherNode.FSubTreeCount,FSubTreeCount);
      HigherNode:=HigherNode.Parent;
    end;
    //if TreeNodes<>nil then Dec(TreeNodes.FCount,FSubTreeCount);
    // remove from parents list
    OldIndex:=FIndex;
    if OldIndex<0 then
      RaiseGDBException('');
    for i:=OldIndex to FParent.FCount-2 do begin
      FParent.FItems[i]:=FParent.FItems[i+1];
      FParent.FItems[i].FIndex:=i;
    end;
    dec(FParent.FCount);
    if (FParent.FCapacity>15) and (FParent.FCount<(FParent.FCapacity shr 2))
    then begin
      // shrink FParent.FItems
      FParent.FCapacity:=FParent.FCapacity shr 1;
      ReAllocMem(FParent.FItems,SizeOf(Pointer)*FParent.FCapacity);
    end;
    if FParent.Count=0 then begin
      FParent.Expanded:=false;
      FParent.HasChildren:=false;
    end;
    FParent:=nil;
  end;
end;

procedure TTreeNode.UnbindFromMultiSelected;
var
  TheTreeNodes: TTreeNodes;
begin
  TheTreeNodes:=TreeNodes;
  if TheTreeNodes=nil then exit;
  if TheTreeNodes.FFirstMultiSelected=Self then
    TheTreeNodes.FFirstMultiSelected:=FNextMultiSelected;
  // Reset last multiselected node
  if TheTreeNodes.FLastMultiSelected=Self then
  begin
    if Self.FNextMultiSelected <> nil then
      TheTreeNodes.FLastMultiSelected := Self.FNextMultiSelected
    else
      if Self.FPrevMultiSelected <> nil then
        TheTreeNodes.FLastMultiSelected := Self.FPrevMultiSelected
      else
        TheTreeNodes.FLastMultiSelected := nil;
  end;
  if FNextMultiSelected<>nil then
    FNextMultiSelected.FPrevMultiSelected:=FPrevMultiSelected;
  if FPrevMultiSelected<>nil then
    FPrevMultiSelected.FNextMultiSelected:=FNextMultiSelected;
  FNextMultiSelected:=nil;
  FPrevMultiSelected:=nil;
end;

function AddModeStr(AddMode: TAddMode): string;
begin
  WriteStr(Result, AddMode);
end;

procedure TTreeNode.InternalMove(ANode: TTreeNode; AddMode: TAddMode);
{
  TAddMode = (taAddFirst, taAdd, taInsert);

  taAdd:      add Self as last child of ANode
  taAddFirst: add Self as first child of ANode
  taInsert:   add Self in front of ANode
}
var
  HigherNode: TTreeNode;
  NewIndex, NewParentItemSize, i: integer;
  WasSelected: Boolean;
begin
  {$IFDEF TREEVIEW_DEBUG}
  DbgOut('[TTreeNode.InternalMove]  Self=',DbgS(Self),' Self.Text=',Text
         ,' ANode=',ANode<>nil,' AddMode=', AddModeStr(AddMode));
  if ANode<>nil then
    DbgOut(' ANode.Text=',ANode.Text);
  DebugLn('');
  {$ENDIF}
  if TreeView<>nil then
    TreeView.BeginUpdate;
  try
    WasSelected:=Selected;
    Unbind;
    if Owner<>nil then
      Owner.ClearCache;
    Include(FStates,nsBound);
    if Owner<>nil then inc(Owner.FCount);
    // set parent
    if AddMode in [taAdd, taAddFirst] then
      FParent:=ANode
    else begin // taInsert
      if (ANode=nil) then
        TreeNodeError('TTreeNode.InternalMove AddMode=taInsert but ANode=nil');
      FParent:=ANode.Parent;
      FPrevBrother:=ANode.FPrevBrother;
      FNextBrother:=ANode;
    end;
    if FParent<>nil then begin
      FParent.HasChildren:=true;
      if (FParent.FCount=FParent.FCapacity) then begin
        // grow FParent.FItems
        if FParent.FCapacity=0 then
          FParent.FCapacity:=5
        else
          FParent.FCapacity:=FParent.FCapacity shl 1;
        NewParentItemSize:=SizeOf(Pointer)*FParent.FCapacity;
        if FParent.FItems=nil then
          GetMem(FParent.FItems,NewParentItemSize)
        else
          ReAllocMem(FParent.FItems,NewParentItemSize);
      end;
      inc(FParent.FCount);
      // calculate new Index
      case AddMode of
      taAdd: NewIndex:=FParent.Count-1;
      taAddFirst: NewIndex:=0;
      else
        // taInsert
        NewIndex:=ANode.Index;
      end;
      // move next siblings
      for i:=FParent.FCount-1 downto NewIndex+1 do begin
        FParent.FItems[i]:=FParent.FItems[i-1];
        FParent.FItems[i].FIndex:=i;
      end;
      // insert this node to parent's items
      FParent.FItems[NewIndex]:=Self;
      FIndex:=NewIndex;
      // set Next and Prev sibling
      if NewIndex>0 then
        FPrevBrother:=FParent.FItems[NewIndex-1]
      else
        FPrevBrother:=nil;
      if NewIndex<FParent.Count-1 then
        FNextBrother:=FParent.FItems[NewIndex+1]
      else
        FNextBrother:=nil;
      // update total node count of all parents
      HigherNode:=FParent;
      while HigherNode<>nil do begin
        inc(HigherNode.FSubTreeCount,FSubTreeCount);
        HigherNode:=HigherNode.Parent;
      end;
      //if TreeNodes<>nil then inc(TreeNodes.FCount,FSubTreeCount);
    end else begin
      // add as top level node
      case AddMode of
      taAdd:
        begin
          // add as last top level node
          if Owner<>nil then begin
            FPrevBrother:=Owner.GetLastNode;
            Owner.MoveTopLvlNode(-1,Owner.FTopLvlCount,Self);
          end;
        end;
      taAddFirst:
        begin
          // add as first top level node = root node
          if Owner<>nil then begin
            FNextBrother:=Owner.GetFirstNode;
            Owner.MoveTopLvlNode(-1,0,Self);
          end;
        end;
      taInsert:
        begin
          // insert node in front of ANode
          //DebugLn('[TTreeNode.InternalMove] ANode.Index=',ANode.Index,' ANode=',DbgS(ANode));
          FNextBrother:=ANode;
          FPrevBrother:=ANode.GetPrevSibling;
          if Owner<>nil then begin
            Owner.MoveTopLvlNode(-1,ANode.Index,Self);
          end;
        end;
      end;
    end;
    // connect Next and Prev sibling
    if FPrevBrother<>nil then FPrevBrother.FNextBrother:=Self;
    if FNextBrother<>nil then FNextBrother.FPrevBrother:=Self;
    if Owner.Owner<>nil then
      Owner.Owner.FStates:=Owner.Owner.FStates+[tvsMaxRightNeedsUpdate,
        tvsTopsNeedsUpdate,tvsTopItemNeedsUpdate,tvsBottomItemNeedsUpdate];
    // re-add to multiselection list
    if MultiSelected then
      BindToMultiSelected;
    if WasSelected then
      Selected:=true;
  finally
    if TreeView<>nil then TreeView.EndUpdate;
  end;

  {$IFDEF TREEVIEW_DEBUG}
  DbgOut('[TTreeNode.InternalMove] END Self=',DbgS(Self),' Self.Text=',Text
         ,' ANode=',DbgS(ANode<>nil),' AddMode=',AddModeStr(AddMode));
  if ANode<>nil then
    DbgOut(' ANode.Text=',ANode.Text);
  DebugLn('');
  {$ENDIF}
end;

procedure TTreeNode.MoveTo(Destination: TTreeNode; Mode: TNodeAttachMode);
{
  TNodeAttachMode = (naAdd, naAddFirst, naAddChild, naAddChildFirst, naInsert);

  naAdd:           add as last sibling of Destination
  naAddFirst:      add as first sibling of Destnation
  naAddChild:      add as last child of Destination
  naAddChildFirst: add as first child of Destination
  naInsert:        insert in front of Destination
  naInsertBehind:  insert behind Destination
}
var
  AddMode: TAddMode;
  //ANode: TTreeNode;
  //HItem: HTreeItem;
  OldOnChanging: TTVChangingEvent;
  OldOnChange: TTVChangedEvent;
begin
  if (Destination=nil) and not(Mode in [naAdd,naAddFirst]) then
    TreeNodeError('TTreeNode.MoveTo Destination=nil');
  if Mode=naInsertBehind then begin      // convert naInsertBehind
    if Destination.GetNextSibling=nil then
      Mode:=naAdd
    else begin
      Mode:=naInsert;
      Destination:=Destination.GetNextSibling;
    end;
  end;
  if (Destination = nil) or not Destination.HasAsParent(Self) then begin
    OldOnChanging := TreeView.OnChanging;
    OldOnChange := TreeView.OnChange;
    TreeView.OnChanging := nil;
    TreeView.OnChange := nil;
    try
      if (Destination <> nil) and (Mode in [naAdd, naAddFirst]) then
        Destination := Destination.Parent;
      case Mode of
        naAddFirst,
        naAddChildFirst: AddMode := taAddFirst;
        naInsert:        AddMode := taInsert;
      else
        AddMode:=taAdd;
      end;
      if (Destination <> Self) then
        InternalMove(Destination, AddMode);
    finally
      TreeView.OnChanging := OldOnChanging;
      TreeView.OnChange := OldOnChange;
    end;
  end;
  Changed(ncParentChanged)
end;

procedure TTreeNode.MultiSelectGroup;
var
  FirstNode, LastNode, ANode: TTreeNode;
begin
  if Assigned(TreeView) and not (tvoAllowMultiselect in TreeView.Options) then
    Exit;
  if Assigned(TreeView) then TreeView.LockSelectionChangeEvent;
  try
    // We need to select the nodes between the selected node and the current node
    FirstNode := GetPrevSibling;
    while Assigned(FirstNode) and not FirstNode.Selected do
      FirstNode := FirstNode.GetPrevSibling;
    if not Assigned(FirstNode) then FirstNode := Self;
    LastNode := GetNextSibling;
    while Assigned(LastNode) and not LastNode.Selected do
      LastNode := LastNode.GetNextSibling;
    if not Assigned(LastNode) then LastNode := Self;
    ANode := FirstNode;
    while Assigned(ANode) do
    begin
      ANode.MultiSelected := True;
      if ANode = LastNode then Break;
      ANode := ANode.GetNextSibling;
    end;
  finally
    if Assigned(TreeView) then TreeView.UnlockSelectionChangeEvent;
  end;
end;

procedure TTreeNode.MakeVisible;
begin
  if Assigned(TreeView) then
    TreeView.EnsureNodeIsVisible(Self)
  else
    ExpandParents;
end;

function TTreeNode.GetLevel: Integer;
// root is on level 0
var
  ANode: TTreeNode;
begin
  Result := 0;
  ANode := Parent;
  while Assigned(ANode) do
  begin
    Inc(Result);
    ANode := ANode.Parent;
  end;
end;

function TTreeNode.GetMultiSelected: Boolean;
begin
  Result := GetState(nsMultiSelected);
end;

function TTreeNode.IsNodeVisible: Boolean;
begin
  if Assigned(TreeView) then
    Result := TreeView.IsNodeVisible(Self)
  else
    Result := AreParentsExpandedAndVisible;
end;

function TTreeNode.IsNodeHeightFullVisible: Boolean;
begin
  if Assigned(TreeView) then
    Result := TreeView.IsNodeHeightFullVisible(Self)
  else
    Result := AreParentsExpandedAndVisible;
end;

procedure TTreeNode.Update;
var
  TV: TCustomTreeView;
begin
  TV := TreeView;
  if Assigned(TV) and (Owner.FUpdateCount = 0) and (not (csLoading in TV.ComponentState)) then
    TV.Invalidate;
end;

function TTreeNode.EditText: Boolean;
var
  TV: TCustomTreeView;
begin
  TV := TreeView;
  Result := Assigned(TV) and (tvsIsEditing in TreeView.FStates);
  TV.BeginEditing(Self);
end;

function TTreeNode.DisplayRect(TextOnly: Boolean): TRect;
begin
  FillChar(Result, SizeOf(Result), 0);
  if TreeView <> nil then
  begin
    Result.Left := TreeView.BorderWidth;
    Result.Top := Top - TreeView.ScrolledTop + TreeView.BorderWidth;
    Result.Right := TreeView.ClientWidth - TreeView.BorderWidth;
    Result.Bottom := Result.Top + Height;
    if TextOnly then
    begin
      Result.Left := DisplayTextLeft;
      if Result.Left > Result.Right then
        Result.Left := Result.Right;
      Result.Right := DisplayTextRight;
      if Result.Right < Result.Left then
        Result.Right := Result.Left;
    end;
  end;
end;

function TTreeNode.DisplayExpandSignLeft: integer;
var
  TV: TCustomTreeView;
  l: LongInt;
  RealIndent: Integer;
begin
  Result := 0;
  TV := TreeView;
  RealIndent := TV.Indent;
  if TV = nil then Exit;
  l := Level;
  if not (tvoShowRoot in TV.Options) then
    inc(Result, RealIndent * (l - 1) + (RealIndent shr 2) + TV.BorderWidth - TV.FScrolledLeft)
  else
    inc(Result, RealIndent * l + TV.BorderWidth - TV.FScrolledLeft);
end;

function TTreeNode.DisplayExpandSignRect: TRect;
begin
  FillChar(Result, SizeOf(Result), 0);
  if TreeView <> nil then
  begin
    Result.Left := DisplayExpandSignLeft;
    Result.Top := Top;
    Result.Right := Result.Left + TreeView.Indent;
    Result.Bottom := Top + Height;
  end;
end;

function TTreeNode.DisplayExpandSignRight: integer;
begin
  Result := DisplayExpandSignLeft;
  if TreeView <> nil then
    inc(Result, TreeView.Indent);
end;

function TTreeNode.DisplayIconLeft: integer;
var
  TV: TCustomTreeView;
begin
  Result := DisplayStateIconLeft;
  TV := TreeView;
  if (TV = nil) or (TV.StateImages = nil) then Exit;
  if (StateIndex < 0) or (StateIndex >= TV.StateImages.Count) then Exit;
  Inc(Result, TV.StateImages.WidthForPPI[TV.StateImagesWidth, TV.Font.PixelsPerInch] + TV.FDefItemSpace);
end;

function TTreeNode.DisplayStateIconLeft: integer;
begin
  Result := DisplayExpandSignRight;
end;

function TTreeNode.DisplayTextLeft: integer;
var
  TV: TCustomTreeView;
  ImgIndex: TImageIndex;
begin
  Result := DisplayIconLeft;
  TV := TreeView;
  if (TV = nil) or (TV.Images = nil) then Exit;
  if (TV.Selected = Self) then
    ImgIndex:=SelectedIndex
  else
    ImgIndex:=ImageIndex;
  if (ImgIndex<0) or (ImgIndex>=TV.Images.Count) then Exit;
  Inc(Result, TV.Images.WidthForPPI[TV.StateImagesWidth, TV.Font.PixelsPerInch] + TV.FDefItemSpace);
end;

function TTreeNode.DisplayTextRight: integer;
var
  TV: TCustomTreeView;
begin
  Result := DisplayTextLeft;
  TV := TreeView;
  if TV <> nil then
    Inc(Result, TV.Canvas.TextWidth(Text) + TV.Indent div 2);
end;

function TTreeNode.AlphaSort: Boolean;
begin
  Result := CustomSort(nil);
end;

function TTreeNode.CustomSort(SortProc: TTreeNodeCompare): Boolean;
begin
  if FCount>0 then begin
    if Owner<>nil then Owner.ClearCache;
    if not Assigned(SortProc) then SortProc:=@DefaultTreeViewSort;
    Sort(FItems, FCount, SortProc, true);
  end;
  if (TreeView <> nil) then
    Include(TreeView.FStates, tvsTopsNeedsUpdate);
  Result:=true;
end;

procedure TTreeNode.Delete;
begin
  if not Deleting then Free;
end;

procedure TTreeNode.DeleteChildren;
begin
  if Owner<>nil then Owner.ClearCache;
  Collapse(true);
  HasChildren := False;
end;

procedure TTreeNode.Assign(Source: TPersistent);
var
  ANode: TTreeNode;
begin
  if Owner<>nil then Owner.ClearCache;
  if Source is TTreeNode then
  begin
    ANode := TTreeNode(Source);
    Text := ANode.Text;
    Data := ANode.Data;
    ImageIndex := ANode.ImageIndex;
    SelectedIndex := ANode.SelectedIndex;
    StateIndex := ANode.StateIndex;
    OverlayIndex := ANode.OverlayIndex;
    Height := ANode.Height;
    Focused := ANode.Focused;
    //DropTarget := ANode.DropTarget;
    Cut := ANode.Cut;
    HasChildren := ANode.HasChildren;
  end
  else inherited Assign(Source);
end;

function TTreeNode.IsEqual(Node: TTreeNode): Boolean;
begin
  Result := (Text = Node.Text) and (Data = Node.Data);
end;

procedure TTreeNode.ReadData(Stream: TStream; StreamVersion: integer);
var
  I, ItemCount: Integer;
  NewExpanded: boolean;
  OldInfo: TOldTreeNodeInfo;
  Info: TTreeNodeInfo;
  Node: TTreeNode;
  lSelfAX: TLazAccessibleObject;
begin
  if Owner<>nil then Owner.ClearCache;
  if StreamVersion=TTreeNodeWithPointerStreamVersion then
  begin
    Stream.ReadBuffer(OldInfo, SizeOf(TOldTreeNodeInfo));
    ImageIndex := OldInfo.ImageIndex;
    SelectedIndex := OldInfo.SelectedIndex;
    StateIndex := OldInfo.StateIndex;
    OverlayIndex := OldInfo.OverlayIndex;
    Data := Pointer(OldInfo.Data);
    Height := OldInfo.Height;
    NewExpanded := OldInfo.Expanded;
    ItemCount := OldInfo.Count;
    SetLength(FText,OldInfo.TextLen)
  end
  else
  begin
    Stream.ReadBuffer(Info, SizeOf(TTreeNodeInfo));
    ImageIndex := Info.ImageIndex;
    SelectedIndex := Info.SelectedIndex;
    StateIndex := Info.StateIndex;
    OverlayIndex := Info.OverlayIndex;
    Height := Info.Height;
    NewExpanded := Info.Expanded;
    ItemCount := Info.Count;
    SetLength(FText,Info.TextLen);
  end;
  if FText<>'' then
  begin
    Stream.Read(FText[1],length(FText));
    // Update accessibility information
    if TreeView<>nil then
    begin
     lSelfAX := TreeView.GetAccessibleObject.GetChildAccessibleObjectWithDataObject(Self);
     if lSelfAX <> nil then
       lSelfAX.AccessibleValue := FText;
    end;
  end;
  if Owner<>nil then begin
    for I := 0 to ItemCount - 1 do begin
      Node:=Owner.AddChild(Self, '');
      Node.ReadData(Stream, StreamVersion);
      Owner.Owner.Added(Node);
    end;
  end;
  Expanded := NewExpanded;
end;

procedure TTreeNode.ReadDelphiData(Stream: TStream; Info: PDelphiNodeInfo);
var
  I, Size, ItemCount: Integer;
begin
  if Owner<>nil then Owner.ClearCache;
  Stream.ReadBuffer(Size, SizeOf(Size));
  Stream.ReadBuffer(Info^, Size);
  Text := Info^.Text;
  ImageIndex := Info^.ImageIndex;
  SelectedIndex := Info^.SelectedIndex;
  StateIndex := Info^.StateIndex;
  OverlayIndex := Info^.OverlayIndex;
  Data := {%H-}Pointer(Info^.Data);
  if Owner<>nil then begin
    ItemCount := Info^.Count;
    for I := 0 to ItemCount - 1 do
      Owner.AddChild(Self, '').ReadDelphiData(Stream, Info);
  end;
end;

procedure TTreeNode.WriteData(Stream: TStream; StreamVersion: integer);
var
  i: integer;
  OldInfo: TOldTreeNodeInfo;
  Info: TTreeNodeInfo;
begin
  if StreamVersion=TTreeNodeWithPointerStreamVersion then
  begin
    OldInfo.ImageIndex := ImageIndex;
    OldInfo.SelectedIndex := SelectedIndex;
    OldInfo.OverlayIndex := OverlayIndex;
    OldInfo.StateIndex := StateIndex;
    OldInfo.Height := FHeight;
    OldInfo.Data := PtrUInt(Data);
    OldInfo.Count := Count;
    OldInfo.Expanded := Expanded;
    OldInfo.TextLen := Length(Text);
    Stream.WriteBuffer(OldInfo, SizeOf(TOldTreeNodeInfo));
  end
  else
  begin
    Info.ImageIndex := ImageIndex;
    Info.SelectedIndex := SelectedIndex;
    Info.OverlayIndex := OverlayIndex;
    Info.StateIndex := StateIndex;
    Info.Height := FHeight;
    Info.Count := Count;
    Info.Expanded := Expanded;
    Info.TextLen := Length(Text);
    Stream.WriteBuffer(Info, SizeOf(TTreeNodeInfo));
  end;
  if Text<>'' then
    Stream.Write(FText[1],length(Text));
  for i := 0 to Count - 1 do
    Items[i].WriteData(Stream, StreamVersion);
end;

procedure TTreeNode.WriteDelphiData(Stream: TStream; Info: PDelphiNodeInfo);
var
  I, Size, L, ItemCount: Integer;
begin
  L := Length(Text);
  if L > 255 then L := 255;
  Size := SizeOf(TDelphiNodeInfo) + L - 255;
  Info^.Text := Text;
  Info^.ImageIndex := ImageIndex;
  Info^.SelectedIndex := SelectedIndex;
  Info^.OverlayIndex := OverlayIndex;
  Info^.StateIndex := StateIndex;
  Info^.Data := {%H-}Cardinal(Data);
  ItemCount := Count;
  Info^.Count := ItemCount;
  Stream.WriteBuffer(Size, SizeOf(Size));
  Stream.WriteBuffer(Info^, Size);
  for I := 0 to ItemCount - 1 do
    Items[I].WriteDelphiData(Stream, Info);
end;

procedure TTreeNode.Changed(ChangeReason: TTreeNodeChangeReason);
begin
  TreeView.NodeChanged(self,ChangeReason);
end;

function TTreeNode.GetOwner: TPersistent;
begin
  Result := FOwner;
end;

procedure TTreeNode.ConsistencyCheck;
var
  RealSubTreeCount: integer;
  i: integer;
  Node1: TTreeNode;
begin
  if FCapacity<0 then
    RaiseGDBException('');
  if FCapacity<FCount then
    RaiseGDBException('');
  if FCount<0 then
    RaiseGDBException('');
  if FHeight<0 then
    RaiseGDBException('');
  if (FItems<>nil) and (FCapacity<=0) then
    RaiseGDBException('');
  if (FCapacity>0) and (FItems=nil) then
    RaiseGDBException('');
  if (FNextBrother<>nil) and (FNextBrother.FPrevBrother<>Self) then
    RaiseGDBException('');
  if (FPrevBrother<>nil) and (FPrevBrother.FNextBrother<>Self) then
    RaiseGDBException('');
  if (FNextMultiSelected<>nil) and (FNextMultiSelected.FPrevMultiSelected<>Self) then
    RaiseGDBException('');
  if (FPrevMultiSelected<>nil) and (FPrevMultiSelected.FNextMultiSelected<>Self) then
    RaiseGDBException('');
  if MultiSelected then begin
    Node1:=TreeView.GetFirstMultiSelected;
    while (Node1<>nil) and (Node1<>Self) do Node1:=Node1.FNextMultiSelected;
    if Node1=nil then
      RaiseGDBException('');
  end;
  if Selected and (TreeView<>nil) and (tvoAllowMultiselect in TreeView.Options)
  and (not MultiSelected) then
    RaiseGDBException('');// selected, but not multiselected

  // check children
  RealSubTreeCount:=1;
  for i:=0 to FCount-1 do begin
    if (Items[i]=nil) then RaiseGDBException('');
    Node1:=Items[i];
    if Node1.FParent<>Self then RaiseGDBException('');
    if (i=0) and (Node1.FPrevBrother<>nil) then
      RaiseGDBException('');
    if (i>0) and (Node1.FPrevBrother=nil) then
      RaiseGDBException('');
    if (i>0) and (Node1.FPrevBrother<>Items[i-1]) then
      RaiseGDBException('');
    if (i<FCount-1) and (Node1.FNextBrother=nil) then
      RaiseGDBException('');
    if (i<FCount-1) and (Node1.FNextBrother<>Items[i+1]) then
      RaiseGDBException('');
    if (i=FCount-1) and (Node1.FNextBrother<>nil) then
      RaiseGDBException('');
    if Node1.FIndex<>i then
      RaiseGDBException('');
    Node1.ConsistencyCheck;
    inc(RealSubTreeCount,Node1.SubTreeCount);
  end;
  if FParent<>nil then begin
    if FParent.IndexOf(Self)<0 then RaiseGDBException('');
  end;
  if RealSubTreeCount<>SubTreeCount then RaiseGDBException('');
  if FTop<0 then RaiseGDBException('');
  // check for circles
  if FNextBrother=Self then RaiseGDBException('');
  if FPrevBrother=Self then RaiseGDBException('');
  if FParent=Self then RaiseGDBException('');
  Node1:=FParent;
  while Node1<>nil do begin
    if (Node1=Self) then RaiseGDBException('');
    Node1:=Node1.FParent;
  end;
end;

procedure TTreeNode.WriteDebugReport(const Prefix: string; Recurse: boolean);
var i: integer;
begin
  DbgOut('%s%s.WriteDebugReport Self=%p',[Prefix, ClassName, Pointer(Self)]);
  ConsistencyCheck;
  DebugLn(' Text=',Text);
  if Recurse then begin
    for i:=0 to FCount-1 do
      Items[i].WriteDebugReport(Prefix+'  ',true);
  end;
end;


{ TTreeNodes }

constructor TTreeNodes.Create(AnOwner: TCustomTreeView);
begin
  inherited Create;
  FSelection := TFPList.Create;
  FOwner := AnOwner;
end;

destructor TTreeNodes.Destroy;
begin
  Clear;
  FreeThenNil(FSelection);
  inherited Destroy;
end;

function TTreeNodes.GetCount: Integer;
begin
  Result:=FCount;
end;

function TTreeNodes.GetOwner: TPersistent;
begin
  Result := FOwner;
end;

function TTreeNodes.GetHandle: THandle;
begin
  if Owner<>nil then
    Result:=Owner.Handle
  else
    Result:=0;
end;

procedure TTreeNodes.Delete(Node: TTreeNode);
begin
  Node.Delete;
  if (FUpdateCount=0) and (Owner<>nil) then
    Owner.Invalidate;
end;

procedure TTreeNodes.Clear;
var
  Node: TTreeNode;
begin
  BeginUpdate;
  ClearCache;
  Node := GetLastNode;
  if Assigned(Node) then
  begin
    while Assigned(Node) do
    begin
      Node.Delete;
      Node := GetLastNode;
    end;
  end;
  FSelection.Clear;
  if (FOwner <> nil) then
    FOwner.GetAccessibleObject().ClearChildAccessibleObjects();
  EndUpdate;
end;

procedure TTreeNodes.ClearMultiSelection(ClearSelected: boolean = false);
var
  ANode, OldNode: TTreeNode;
begin
  if Assigned(Owner) then Owner.LockSelectionChangeEvent;
  try
    ANode := FFirstMultiSelected;
    while Assigned(ANode) do
    begin
      OldNode := ANode;
      ANode := ANode.GetNextMultiSelected;
      OldNode.MultiSelected := false;
    end;
    if ClearSelected then
      Owner.Selected := nil;
  finally
    if Assigned(Owner) then Owner.UnlockSelectionChangeEvent;
  end;
end;

procedure TTreeNodes.SelectOnlyThis(Node: TTreeNode);
begin
  if Owner<>nil then Owner.LockSelectionChangeEvent;
  try
    ClearMultiSelection(true);
    Node.Selected:=true;
  finally
    if Owner<>nil then Owner.UnlockSelectionChangeEvent;
  end;
end;

function TTreeNodes.IsMultiSelection: boolean;
begin
  Result:=(FFirstMultiSelected<>nil)
          and (FFirstMultiSelected.GetNextMultiSelected<>nil);
end;

function TTreeNodes.AddChildFirst(ParentNode: TTreeNode; const S: string): TTreeNode;
begin
  Result := AddChildObjectFirst(ParentNode, S, nil);
end;

function TTreeNodes.AddChildObjectFirst(ParentNode: TTreeNode; const S: string;
  Data: Pointer): TTreeNode;
begin
  Result := InternalAddObject(ParentNode, S, Data, taAddFirst);
end;

function TTreeNodes.AddChild(ParentNode: TTreeNode; const S: string): TTreeNode;
begin
  Result := AddChildObject(ParentNode, S, nil);
end;

function TTreeNodes.AddChildObject(ParentNode: TTreeNode; const S: string;
  Data: Pointer): TTreeNode;
begin
  Result := InternalAddObject(ParentNode, S, Data, taAdd);
end;

function TTreeNodes.AddFirst(SiblingNode: TTreeNode; const S: string): TTreeNode;
begin
  Result := AddObjectFirst(SiblingNode, S, nil);
end;

function TTreeNodes.AddObjectFirst(SiblingNode: TTreeNode; const S: string;
  Data: Pointer): TTreeNode;
var ParentNode: TTreeNode;
begin
  if SiblingNode <> nil then
    ParentNode := SiblingNode.Parent
  else
    ParentNode := nil;
  Result := InternalAddObject(ParentNode, S, Data, taAddFirst);
end;

function TTreeNodes.AddNode(Node: TTreeNode; Relative: TTreeNode;
  const S: string; Ptr: Pointer; Method: TNodeAttachMode): TTreeNode;
var
   AddMode: TAddMode;
begin
  if (Relative=nil) and not (Method in [naAdd,naAddFirst]) then
    TreeNodeError('TTreeNode.AddNode Relative=nil');
  if Method=naInsertBehind then begin    // convert naInsertBehind
    if Relative.GetNextSibling=nil then
      Method:=naAdd
    else begin
      Method:=naInsert;
      Relative:=Relative.GetNextSibling;
    end;
  end;
  if (Relative <> nil) and (Method in [naAdd, naAddFirst]) then
    Relative := Relative.Parent;
  // Convert TNodeAttachMode to TAddMode
  case Method of
    naAddFirst,naAddChildFirst: AddMode := taAddFirst;
    naInsert: AddMode := taInsert;
  else
    AddMode:=taAdd;
  end;
  fNewNodeToBeAdded := Node;
  Result := InternalAddObject(Relative, S, Ptr, AddMode);
end;

procedure TTreeNodes.SelectionsChanged(ANode: TTreeNode; const AIsSelected: Boolean);
begin
  if ANode <> nil then
  begin
    if AIsSelected then
      FSelection.Add(ANode)
    else
      FSelection.Remove(ANode);
  end;
end;

function TTreeNodes.GetSelections(const AIndex: Integer): TTreeNode;

  procedure RaiseOutOfBounds;
  begin
    TreeNodeError('TTreeNodes.GetSelections Index '+IntToStr(AIndex)
        +' out of bounds (Count='+IntToStr(FSelection.Count)+')');
  end;

begin
  if (AIndex < 0) or (AIndex > FSelection.Count - 1) then
    RaiseOutOfBounds;
  Result := TTreeNode(FSelection.Items[AIndex]);
end;

function TTreeNodes.Add(SiblingNode: TTreeNode; const S: string): TTreeNode;
begin
  Result := AddObject(SiblingNode, S, nil);
end;

procedure TTreeNodes.Repaint(ANode: TTreeNode);
var
  R: TRect;
begin
  if (FUpdateCount < 1) and (Owner<>nil) then begin
    while (ANode <> nil) and not ANode.IsVisible do ANode := ANode.Parent;
    if ANode <> nil then begin
      R := ANode.DisplayRect(False);
      InvalidateRect(Owner.Handle, @R, True);
    end;
  end;
end;

function TTreeNodes.AddObject(SiblingNode: TTreeNode; const S: string;
  Data: Pointer): TTreeNode;
var ParentNode: TTreeNode;
begin
  if SiblingNode <> nil then
    ParentNode := SiblingNode.Parent
  else
    ParentNode := nil;
  Result := InternalAddObject(ParentNode, S, Data, taAdd);
end;

function TTreeNodes.Insert(NextNode: TTreeNode; const S: string): TTreeNode;
begin
  Result := InsertObject(NextNode, S, nil);
end;

function TTreeNodes.InsertObject(NextNode: TTreeNode; const S: string;
  Data: Pointer): TTreeNode;
// create a new node with Text=S and Data=Data and insert in front of
// NextNode (as sibling with same parent).
begin
  Result:=InternalAddObject(NextNode,S,Data,taInsert);
end;

function TTreeNodes.InsertBehind(PrevNode: TTreeNode; const S: string): TTreeNode;
begin
  Result := InsertObjectBehind(PrevNode, S, nil);
end;

function TTreeNodes.InsertObjectBehind(PrevNode: TTreeNode; const S: string;
  Data: Pointer): TTreeNode;
// create a new node with Text=S and Data=Data and insert in front of
// NextNode (as sibling with same parent).
begin
  if (PrevNode<>nil) and (PrevNode.GetNextSibling<>nil) then
    Result:=InternalAddObject(PrevNode.GetNextSibling,S,Data,taInsert)
  else
    Result:=AddObject(PrevNode,S,Data);
end;

procedure TTreeNodes.SortTopLevelNodes(SortProc: TTreeNodeCompare);
begin
  Sort(FTopLvlItems, FTopLvlCount, SortProc, true);
end;

function TTreeNodes.InternalAddObject(Node: TTreeNode; const S: string;
  Data: Pointer; AddMode: TAddMode): TTreeNode;
{
  TAddMode = (taAddFirst, taAdd, taInsert);

  taAdd:      add Result as last child of Node
  taAddFirst: add Result as first child of Node
  taInsert:   add Result in front of Node
}
var
  ok: boolean;
  // Item: HTreeItem;
  lAccessibleObject: TLazAccessibleObject;
begin
  if Owner=nil then
    TreeNodeError('TTreeNodes.InternalAddObject Owner=nil');
  {$IFDEF TREEVIEW_DEBUG}
  write('[TTreeNodes.InternalAddObject] Node=',Node<>nil,' S=',S,
        ' AddMode=',AddModeStr(AddMode));
  if Node<>nil then
    DbgOut(' Node.Text=',Node.Text);
  DebugLn('');
  {$ENDIF}
  Result := fNewNodeToBeAdded; // Used by AddNode to pass an existing node.
  if Result = Nil then
    Result := Owner.CreateNode;
  fNewNodeToBeAdded := nil;
  ok:=false;
  try
    Result.Data := Data;
    Result.Text := S;
    // move node in tree (tree of TTreeNode)
    Result.InternalMove(Node,AddMode);
    if (Owner<>nil) and Owner.AutoExpand and (Result.Parent<>nil) then
      Result.Parent.Expanded:=true;
    if (Owner<>nil) and (not (csReading in Owner.ComponentState)) then
      Owner.Added(Result);
    ok:=true;
    if ok and (Owner<>nil) and (Owner.AccessibilityOn) then
    begin
      lAccessibleObject := FOwner.GetAccessibleObject().AddChildAccessibleObject();
      lAccessibleObject.AccessibleDescription := 'Item';
      lAccessibleObject.AccessibleValue := S;
      lAccessibleObject.AccessibleRole := larTreeItem;
      lAccessibleObject.DataObject := Result;
    end;
  finally
    // this construction creates nicer exception output
    if not ok then
      Result.Free;
  end;
end;

function TTreeNodes.GetFirstNode: TTreeNode;
begin
  if Assigned(FTopLvlItems) then
    Result := FTopLvlItems[0]
  else
    Result := nil;
end;

function TTreeNodes.GetFirstVisibleNode: TTreeNode;
var
  Node: TTreeNode;
  i: Integer;
begin
  Result := nil;
  if Assigned(FTopLvlItems) then
    for i := 0 to FTopLvlCount-1 do begin
      Node := FTopLvlItems[i];
      if Node.Visible then begin
        Result := Node;
        Break;
      end;
    end;
end;

function TTreeNodes.GetLastNode: TTreeNode;
begin
  if Assigned(FTopLvlItems) then
    Result := FTopLvlItems[FTopLvlCount - 1]
  else
    Result := nil;
end;

function TTreeNodes.GetLastVisibleNode: TTreeNode;
var
  Node: TTreeNode;
  i: Integer;
begin
  Result := nil;
  if Assigned(FTopLvlItems) then
    for i := FTopLvlCount-1 downto 0 do begin
      Node := FTopLvlItems[i];
      if Node.Visible then begin
        Result := Node;
        Break;
      end;
    end;
end;

function TTreeNodes.GetLastSubNode: TTreeNode;
// absolute last node
var
  Node: TTreeNode;
begin
  Result := GetLastNode;
  if Assigned(Result) then
  begin
    Node := Result.GetLastSubChild;
    if Assigned(Node) then Result := Node;
  end;
end;

function TTreeNodes.GetLastExpandedSubNode: TTreeNode;
// absolute last expanded node
var
  Node: TTreeNode;
begin
  Result := GetLastVisibleNode;
  while Assigned(Result) and (Result.Expanded) do
  begin
    Node := Result.GetLastVisibleChild;
    if Assigned(Node) then
      Result := Node
    else
      exit;
  end;
end;

function TTreeNodes.FindTopLvlNode(const NodeText: string): TTreeNode;
begin
  Result := GetFirstNode;
  while Assigned(Result) and (Result.Text <> NodeText) do
    Result := Result.GetNextSibling;
end;

function TTreeNodes.FindNodeWithText(const NodeText: string): TTreeNode;
begin
  Result := GetFirstNode;
  while Assigned(Result) and (Result.Text <> NodeText) do
    Result := Result.GetNext;
end;

function TTreeNodes.FindNodeWithTextPath(TextPath: string): TTreeNode;
var
  p: SizeInt;
  CurText: String;
begin
  Result:=nil;
  repeat
    p:=System.Pos('/',TextPath);
    if p>0 then begin
      CurText:=LeftStr(TextPath,p-1);
      System.Delete(TextPath,1,p);
    end else begin
      CurText:=TextPath;
      TextPath:='';
    end;
    //debugln(['TTreeNodes.FindNodeWithTextPath CurText=',CurText,' Rest=',TextPath]);
    if Result=nil then
      Result:=FindTopLvlNode(CurText)
    else
      Result:=Result.FindNode(CurText);
  until (Result=nil) or (TextPath='');
end;

function TTreeNodes.FindNodeWithData(const NodeData: Pointer): TTreeNode;
begin
  Result := GetFirstNode;
  while Assigned(Result) and (Result.Data <> NodeData) do
    Result := Result.GetNext;
end;

function TTreeNodes.GetNodeFromIndex(Index: Integer): TTreeNode;
// find node with absolute index in ALL nodes (even collapsed)

  procedure RaiseIndexOutOfBounds;
  begin
    TreeNodeError('TTreeNodes.GetNodeFromIndex Index '+IntToStr(Index)
           +' out of bounds (Count='+IntToStr(FCount)+')');
  end;

  procedure RaiseSubTreeCount0;
  begin
    TreeNodeError(
      'TTreeNodes.GetNodeFromIndex: Consistency error - SubTreeCount=0');
  end;

  procedure RaiseSubTreeCountTooBig;
  begin
    TreeNodeError(
      'TTreeNodes.GetNodeFromIndex: Consistency error - invalid SubTreeCount');
  end;

  procedure RaiseCountTooBig;
  begin
    TreeNodeError(
      'TTreeNodes.GetNodeFromIndex: Consistency Error - Count too big');
  end;

var
  I, J: Integer;
begin
  if (Index < 0) or (Index >= FCount) then
    RaiseIndexOutOfBounds;
  if (FNodeCache.CacheNode <> nil) and (Abs(FNodeCache.CacheIndex - Index) <= 1)
  then begin
    with FNodeCache do
    begin
      if Index = CacheIndex then Result := CacheNode
      else if Index < CacheIndex then Result := CacheNode.GetPrev
      else Result := CacheNode.GetNext;
    end;
  end
  else if Index>Count-5 then begin
    // optimization for the last nodes
    Result:=GetLastSubNode;
    i:=Count-1;
    while (Index<i) do begin
      Result:=Result.GetPrev;
      dec(i);
    end;
  end
  else begin
    Result := GetFirstNode;
    I:=0;
    while (Result<>nil) and (Index>I) do begin
      Repeat
        // calculate the absolute index of the next sibling
        J:=I+Result.FSubTreeCount;
        if J=I then RaiseSubTreeCount0;
        if J<=Index then begin
          // Index > absolute index of next sibling -> search in next sibling
          Result:=Result.GetNextSibling;
          I:=J;
        end else
          break;
      until false;
      if (Result<>nil) and (Index>I) then begin
        // Index is somewhere in subtree of Result
        Result:=Result.GetFirstChild;
        if Result=nil then RaiseSubTreeCountTooBig;
        inc(I);
      end;
    end;
  end;
  if Result = nil then RaiseCountTooBig;
  FNodeCache.CacheNode := Result;
  FNodeCache.CacheIndex := Index;
end;

function TTreeNodes.GetSelectionCount: Cardinal;
begin
  Result := Cardinal(FSelection.Count);
end;

procedure TTreeNodes.SetItem(Index: Integer; AValue: TTreeNode);
begin
  GetNodeFromIndex(Index).Assign(AValue);
end;

procedure TTreeNodes.SetTopLvlItems(Index: integer; AValue: TTreeNode);
begin
  GetTopLvlItems(Index).Assign(AValue);
end;

procedure TTreeNodes.BeginUpdate;
begin
  if FUpdateCount = 0 then SetUpdateState(True);
  Inc(FUpdateCount);
end;

procedure TTreeNodes.SetUpdateState(Updating: Boolean);
begin
  //SendMessage(Handle, WM_SETREDRAW, Ord(not Updating), 0);
  if Updating then
    Include(Owner.FStates,tvsUpdating)
  else
    Exclude(Owner.FStates,tvsUpdating);
  if not Updating then
    Owner.Invalidate;
end;

procedure TTreeNodes.EndUpdate;
begin
  Dec(FUpdateCount);
  if FUpdateCount = 0 then SetUpdateState(False);
end;

procedure TTreeNodes.FreeAllNodeData;
var
  i: Integer;
begin
  BeginUpdate;
  for i:=0 to TopLvlCount-1 do
    TopLvlItems[i].FreeAllNodeData;
  EndUpdate;
end;

function TTreeNodes.GetEnumerator: TTreeNodesEnumerator;
begin
  Result := TTreeNodesEnumerator.Create(Self);
end;

procedure TTreeNodes.GrowTopLvlItems;
begin
  if FTopLvlItems<>nil then begin
    FTopLvlCapacity:=FTopLvlCapacity shl 1;
    ReAllocMem(FTopLvlItems,SizeOf(Pointer)*FTopLvlCapacity);
  end else begin
    FTopLvlCapacity:=MinNodeCapacity;
    GetMem(FTopLvlItems,SizeOf(Pointer)*FTopLvlCapacity);
  end;
  //debugln('TTreeNodes.GrowTopLvlItems END FTopLvlCapacity=',FTopLvlCapacity,' FTopLvlCount=',FTopLvlCount,' ',FTopLvlItems<>nil);
end;

function TTreeNodes.GetTopLvlItems(Index: integer): TTreeNode;
begin
  Result:=FTopLvlItems[Index];
end;

procedure TTreeNodes.ShrinkTopLvlItems;
var
  OldCapacity: LongInt;
begin
  if FTopLvlItems<>nil then begin
    OldCapacity:=FTopLvlCapacity;
    FTopLvlCapacity:=FTopLvlCapacity shr 1;
    if FTopLvlCapacity<FTopLvlCount then FTopLvlCapacity:=FTopLvlCount;
    if (FTopLvlCapacity<MinNodeCapacity) then begin
      if (FTopLvlCount>0) then
        FTopLvlCapacity:=MinNodeCapacity
      else
        FTopLvlCapacity:=0;
    end;
    if OldCapacity=FTopLvlCapacity then exit;
    //debugln('TTreeNodes.ShrinkTopLvlItems A FTopLvlCapacity=',FTopLvlCapacity,' FTopLvlCount=',FTopLvlCount,' ',FTopLvlItems<>nil);
    ReAllocMem(FTopLvlItems,SizeOf(Pointer)*FTopLvlCapacity);
    //debugln('TTreeNodes.ShrinkTopLvlItems B FTopLvlCapacity=',FTopLvlCapacity,' FTopLvlCount=',FTopLvlCount,' ',FTopLvlItems<>nil);
  end else begin
    if (FTopLvlCapacity>0) then
      TreeNodeError('TTreeNodes.ShrinkTopLvlItems FTopLvlCapacity>0');
  end;
end;

function TTreeNodes.IndexOfTopLvlItem(Node: TTreeNode): integer;
begin
  if (Node<>nil) and (Node.Owner=Self) then
    Result:=Node.FIndex
  else
    Result:=-1;
end;

procedure TTreeNodes.MoveTopLvlNode(TopLvlFromIndex, TopLvlToIndex: integer;
  Node: TTreeNode);
// TopLvlFromIndex = -1 and is insert
// TopLvlToIndex = -1 is remove
var i: integer;
begin
  {$IFDEF TREEVIEW_DEBUG}
  DebugLn('[TTreeNodes.MoveTopLvlNode] TopLvlFromIndex=',TopLvlFromIndex,
  ' TopLvlToIndex=',TopLvlToIndex,' Node.Text=',Node.Text);
  {$ENDIF}
  if TopLvlFromIndex=TopLvlToIndex then exit;
  if (TopLvlFromIndex>=FTopLvlCount) then
    TreeNodeError('TTreeNodes.MoveTopLvlNode TopLvlFromIndex>FTopLvlCount');
  if (TopLvlToIndex>FTopLvlCount) then
    TreeNodeError('TTreeNodes.MoveTopLvlNode TopLvlFromIndex>FTopLvlCount');
  if (TopLvlFromIndex>=0) then begin
    Node:=FTopLvlItems[TopLvlFromIndex];
    if (TopLvlToIndex>=0) then begin
      // move node
      if TopLvlFromIndex<TopLvlToIndex then begin
        // move forward
        for i:=TopLvlFromIndex to TopLvlToIndex-1 do begin
          FTopLvlItems[i]:=FTopLvlItems[i+1];
          FTopLvlItems[i].FIndex:=i;
        end;
      end else begin
        // move backward
        for i:=TopLvlToIndex downto TopLvlFromIndex+1 do begin
          FTopLvlItems[i]:=FTopLvlItems[i-1];
          FTopLvlItems[i].FIndex:=i;
        end;
      end;
      FTopLvlItems[TopLvlToIndex]:=Node;
      FTopLvlItems[TopLvlToIndex].FIndex:=TopLvlToIndex;
    end else begin
      // remove node
      if FTopLvlItems<>nil then begin
        for i:=TopLvlFromIndex to FTopLvlCount-2 do begin
          FTopLvlItems[i]:=FTopLvlItems[i+1];
          FTopLvlItems[i].FIndex:=i;
        end;
      end;
      Dec(FTopLvlCount);
      if FTopLvlCount<0 then
        TreeNodeError('TTreeNodes.MoveTopLvlNode FTopLvlCount<0');
      if FTopLvlCount<(FTopLvlCapacity shr 2) then ShrinkTopLvlItems;
    end;
  end else begin
    if (TopLvlToIndex>=0) then begin
      if Node=nil then
        TreeNodeError('TTreeNodes.MoveTopLvlNode inserting nil');
      // insert node
      if (FTopLvlCount=FTopLvlCapacity) then GrowTopLvlItems;
      inc(FTopLvlCount);
      if FTopLvlItems<>nil then begin
        for i:=FTopLvlCount-1 downto TopLvlToIndex+1 do begin
          FTopLvlItems[i]:=FTopLvlItems[i-1];
          FTopLvlItems[i].FIndex:=i;
        end;
        FTopLvlItems[TopLvlToIndex]:=Node;
        FTopLvlItems[TopLvlToIndex].FIndex:=TopLvlToIndex;
      end;
    end else begin
      // nothing to do
    end;
  end;
end;

procedure TTreeNodes.MultiSelect(Node: TTreeNode; ClearWholeSelection: Boolean);
var
  bGoNext, bOnlySiblings, bOnlyVisible: Boolean;
  //
  procedure _TakeNext(var N: TTreeNode);
  begin
    if bGoNext then
    begin
      if bOnlySiblings and bOnlyVisible then
        N := N.GetNextVisibleSibling
      else
      if bOnlySiblings and not bOnlyVisible then
        N := N.GetNextSibling
      else
      if not bOnlySiblings and bOnlyVisible then
        N := N.GetNextVisible
      else
        N := N.GetNext;
    end
    else
    begin
      if bOnlySiblings and bOnlyVisible then
        N := N.GetPrevVisibleSibling
      else
      if bOnlySiblings and not bOnlyVisible then
        N := N.GetPrevSibling
      else
      if not bOnlySiblings and bOnlyVisible then
        N := N.GetPrevVisible
      else
        N := N.GetPrev;
    end;
  end;
  //
var
  I, FirstNode, LastNode: TTreeNode;
begin
  if Owner<>nil then Owner.LockSelectionChangeEvent;
  bOnlySiblings := Assigned(Owner) and (msSiblingOnly in Owner.MultiSelectStyle);
  bOnlyVisible := Assigned(Owner) and (msVisibleOnly in Owner.MultiSelectStyle);

  try
    if FStartMultiSelected=nil then
    begin
      FirstNode := Node;
      FStartMultiSelected := Node;
    end else
      FirstNode := FStartMultiSelected;

    if ClearWholeSelection then
    begin
      ClearMultiSelection(True);
    end else
    begin
      //clear only last selection
      if Assigned(FLastMultiSelected) then
      begin
        LastNode := FLastMultiSelected;
        bGoNext := (FirstNode.Index <= LastNode.Index);
        I := FirstNode;
        I.MultiSelected:=False;
        while (I<>LastNode) do
        begin
          _TakeNext(I);
          if I=nil then Break;
          I.MultiSelected:=False;
        end;
      end;
      if Assigned(Owner) then
        Owner.Selected := nil;
    end;

    //select again
    bGoNext := (FirstNode.Index <= Node.Index);
    I := FirstNode;
    I.MultiSelected:=True;
    while (I<>Node) do
    begin
      _TakeNext(I);
      if I=nil then Break;
      I.MultiSelected:=True;
    end;

    FStartMultiSelected := FirstNode;
    FLastMultiSelected := Node;
  finally
    if Owner<>nil then Owner.UnlockSelectionChangeEvent;
  end;
end;

procedure TTreeNodes.Assign(Source: TPersistent);
var
  SrcNodes: TTreeNodes;
  SrcStream: TMemoryStream;
begin
  ClearCache;
  if Source is TTreeNodes then begin
    SrcNodes := TTreeNodes(Source);
    Clear;
    SrcStream := TMemoryStream.Create;
    try
      SrcNodes.WriteData(SrcStream, true);
      SrcStream.Position := 0;
      ReadData(SrcStream);
    finally
      SrcStream.Free;
    end;
  end
  else
    inherited Assign(Source);
end;

procedure TTreeNodes.DefineProperties(Filer: TFiler);

  function WriteNodes: Boolean;
  var
    I: Integer;
    Nodes: TTreeNodes;
  begin
    Nodes := TTreeNodes(Filer.Ancestor);
    if Nodes = nil then
      Result := Count > 0
    else if Nodes.Count <> Count then
      Result := True
    else
    begin
      Result := False;
      for I := 0 to Count - 1 do
      begin
        Result := not Item[I].IsEqual(Nodes[I]);
        if Result then Break;
      end
    end;
  end;

begin
  inherited DefineProperties(Filer);
  // Delphi 10.2.1
  Filer.DefineBinaryProperty('NodeData', @ReadDelphiNodeData, nil, False);
  Filer.DefineBinaryProperty('Data', @ReadData, @WriteData, WriteNodes);
end;

procedure TTreeNodes.ReadData(Stream: TStream);
var
  I, NewCount, MagicNumber: Integer;
  DelphiNodeInfo: TDelphiNodeInfo;
  StreamVersion: word;
begin
  Clear;
  // -7 for lcl stream
  Stream.ReadBuffer(MagicNumber,SizeOf(Integer));
  if MagicNumber=LCLStreamID then begin
    // read stream version
    Stream.ReadBuffer(StreamVersion,SizeOf(StreamVersion));
    // read top level node count
    Stream.ReadBuffer(NewCount, SizeOf(NewCount));
    for I := 0 to NewCount - 1 do
      Add(nil, '').ReadData(Stream, StreamVersion);
  end else begin
    // delphi stream
    NewCount:=MagicNumber;
    for I := 0 to NewCount - 1 do
      Add(nil, '').ReadDelphiData(Stream, @DelphiNodeInfo);
  end;
end;

// from Delphi Vcl.ComCtrls.pas
const
  TreeNodeStreamVersion1    = $01;  // 32-bit struct size version 1
  TreeNodeStreamVersion1x64 = $02;  // 64-bit struct size version 1
  TreeNodeStreamVersion2    = $03;  // 32-bit struct size version 2
  TreeNodeStreamVersion2x64 = $04;  // 64-bit struct size version 2

type
  TNodeDataType = (ndtDefault, ndtDefault2, ndt32bit, ndt64bit, ndt232bit, ndt264bit);

procedure TTreeNodes.ReadDelphiNodeData(Stream: TStream);


  function ReadInteger: Integer;
  begin
    Stream.Read(Result, SizeOf(Integer));
  end;

 procedure ReadNodeData(ANode: TTreeNode; Stream: TStream; NodeDataType: TNodeDataType);
 var
   I, LSize, LItemCount: Integer;
   LNode: TTreeNode;
   LText: UnicodeString;

   // 读取一定要按这个顺序
   //LImageIndex: Integer;
   //LSelectedIndex: Integer;
   //LStateIndex: Integer;
   //LOverlayIndex: Integer;
   //LExpandedIndex: Integer;
   //LData: Integer;
   //LCount: Integer;
   //LEnabled: Byte;
   LTextLen: Byte;
 begin
   ANode.OWner.ClearCache;
   // 只读这一种
   case NodeDataType of
     ndt232bit:
       begin
         LSize := ReadInteger;

         //LImageIndex: Integer;
         //LSelectedIndex: Integer;
         //LStateIndex: Integer;
         //LOverlayIndex: Integer;
         //LExpandedIndex: Integer;
         //LData: Integer;
         //LCount: Integer;
         //LEnabled: Byte;
         //LTextLen: Byte;

         ANode.ImageIndex := ReadInteger;
         ANode.SelectedIndex := ReadInteger;
         ANode.StateIndex := ReadInteger;
         ANode.OverlayIndex := ReadInteger;
         //ANode.ExpandedImageIndex := ReadInteger;
         ReadInteger;
         ANode.Data := Pointer(ReadInteger);
         LItemCount := ReadInteger;
         //ANode.Enabled := Stream.ReadByte <> 0;
         Stream.ReadByte;

         LTextLen := Stream.ReadByte;
         // 读文本
         SetLength(LText, LTextLen);
         Stream.ReadBuffer(LText[1], LTextLen * SizeOf(WideChar));
         ANode.Text := UTF16ToUTF8(LText);

         ANode.HasChildren := LItemCount <> 0;
         if LItemCount > 0 then
         begin
           for I := 0 to LItemCount - 1 do
           begin
             LNode := ANode.Owner.AddChild(ANode, '');
             ReadNodeData(LNode, Stream, NodeDataType);
             ANode.Owner.Owner.Added(LNode);
           end;
         end;
       end;
   else
     Exit;
   end;
 end;

var
  LStreamVersion: Byte;
  I, LCount: Integer;
  LNode: TTreeNode;
  LHandleAllocated: Boolean;
  LNodeDataType: TNodeDataType;
begin
  LHandleAllocated := Owner.HandleAllocated;
  if LHandleAllocated then
    BeginUpdate;
  try
    Clear;
    LStreamVersion := Stream.ReadByte;
    case LStreamVersion of
      TreeNodeStreamVersion1:
        LNodeDataType := ndtDefault;
      TreeNodeStreamVersion2:
        LNodeDataType := ndt232bit;
      TreeNodeStreamVersion2x64:
        LNodeDataType := ndt264bit;
    else
      Exit;
    end;

    LCount := ReadInteger;
    for I := 0 to LCount - 1 do
    begin
      LNode := Add(nil, '');
      ReadNodeData(LNode, Stream, LNodeDataType);
      Owner.Added(LNode);
    end;
  finally
    if LHandleAllocated then
      EndUpdate;
  end;
end;

procedure TTreeNodes.WriteData(Stream: TStream);
begin
   WriteData(Stream, false);
end;

procedure TTreeNodes.WriteData(Stream: TStream; WriteDataPointer: boolean);
var
  ANode: TTreeNode;
  MagicNumber: integer;
  StreamVersion: word;
begin
  // -7 for lcl stream
  MagicNumber:=LCLStreamID;
  Stream.WriteBuffer(MagicNumber,SizeOf(MagicNumber));
  // write stream version
  if WriteDataPointer then
    StreamVersion:=TTreeNodeWithPointerStreamVersion
  else
    StreamVersion:=TTreeNodeStreamVersion;
  Stream.WriteBuffer(StreamVersion,SizeOf(Word));
  // write top level node count
  Stream.WriteBuffer(FTopLvlCount, SizeOf(Integer));
  // write all nodes recursively
  ANode := GetFirstNode;
  while ANode <> nil do begin
    ANode.WriteData(Stream, StreamVersion);
    ANode := ANode.GetNextSibling;
  end;
end;

procedure TTreeNodes.ReadExpandedState(Stream: TStream);
var
  ItemCount,
  Index: Integer;
  Node: TTreeNode;
  NodeExpanded: Boolean;
begin
  if Stream.Position < Stream.Size then
    Stream.ReadBuffer(ItemCount, SizeOf(ItemCount))
  else Exit;
  Index := 0;
  Node := GetFirstNode;
  while (Index < ItemCount) and (Node <> nil) do begin
    Stream.ReadBuffer(NodeExpanded, SizeOf(NodeExpanded));
    Node.Expanded := NodeExpanded;
    Inc(Index);
    Node := Node.GetNext;
  end;
end;

procedure TTreeNodes.WriteExpandedState(Stream: TStream);
var
  Size: Integer;
  ANode: TTreeNode;
  NodeExpanded: Boolean;
begin
  Size := SizeOf(Boolean) * Count;
  Stream.WriteBuffer(Size, SizeOf(Size));
  ANode := GetFirstNode;
  while (ANode <> nil) do begin
    NodeExpanded := ANode.Expanded;
    Stream.WriteBuffer(NodeExpanded, SizeOf(Boolean));
    ANode := ANode.GetNext;
  end;
end;

procedure TTreeNodes.ClearCache;
begin
  FNodeCache.CacheNode := nil;
end;

procedure TTreeNodes.ConsistencyCheck;
var
  Node: TTreeNode;
  RealCount, i: integer;
  OldCache: TNodeCache;
begin
  if FUpdateCount<0 then
    RaiseGDBException('FUpdateCount<0');
  RealCount:=0;
  Node:=GetFirstNode;
  while Node<>nil do begin
    Node.ConsistencyCheck;
    inc(RealCount,Node.SubTreeCount);
    //DebugLn(' ConsistencyCheck: B  ',RealCount,',',Node.SubTreeCount);
    Node:=Node.FNextBrother;
  end;
  //DebugLn(' ConsistencyCheck: B  ',RealCount,',',FCount);
  if RealCount<>FCount then
    RaiseGDBException('RealCount<>FCount');
  if (FTopLvlCapacity<=0) and (FTopLvlItems<>nil) then
    RaiseGDBException('');
  if (FTopLvlCapacity>0) and (FTopLvlItems=nil) then
    RaiseGDBException('');
  if FTopLvlCapacity<FTopLvlCount then
    RaiseGDBException('');
  if (FTopLvlCount<0) then
    RaiseGDBException('');
  for i:=0 to FTopLvlCount-1 do begin
    Node:=FTopLvlItems[i];
    if (i=0) and (Node.FPrevBrother<>nil) then
      RaiseGDBException('');
    if (i>0) and (Node.FPrevBrother<>FTopLvlItems[i-1]) then
      RaiseGDBException('');
    if (i<FTopLvlCount-1) and (Node.FNextBrother<>FTopLvlItems[i+1])
    then begin
      DebugLn(' CONSISTENCY i=%d FTopLvlCount=%d FTopLvlItems[i]=%p FTopLvlItems[i].FNextBrother=%p FTopLvlItems[i+1]=%p',
             [i, FTopLvlCount, Pointer(Node), Pointer(Node.FNextBrother), Pointer(FTopLvlItems[i+1])]);
      RaiseGDBException('');
    end;
    if (i=FTopLvlCount-1) and (Node.FNextBrother<>nil) then
      RaiseGDBException('');
    if Node.FIndex<>i then
      RaiseGDBException('');
  end;
  if FNodeCache.CacheNode<>nil then begin
    OldCache:=FNodeCache;
    ClearCache;
    if GetNodeFromIndex(OldCache.CacheIndex)<>OldCache.CacheNode then
      RaiseGDBException('');
  end;
end;

procedure TTreeNodes.WriteDebugReport(const Prefix: string; AllNodes: boolean);
var
  Node: TTreeNode;
begin
  DbgOut('%s%s.WriteDebugReport Self=%p', [Prefix, ClassName, Pointer(Self)]);
  ConsistencyCheck;
  DebugLn('');
  if AllNodes then begin
    Node:=GetFirstNode;
    while Node<>nil do begin
      Node.WriteDebugReport(Prefix+'  ',true);
      Node:=Node.GetNextSibling;
    end;
  end;
end;


type
  TTreeStrings = class(TStrings)
  private
    FOwner: TTreeNodes;
  protected
    function Get(Index: Integer): string; override;
    function GetBufStart(Buffer: PChar; var Level: Integer): PChar;
    function GetCount: Integer; override;
    function GetObject(Index: Integer): TObject; override;
    procedure PutObject(Index: Integer; AObject: TObject); override;
    procedure SetUpdateState(Updating: Boolean); override;
  public
    constructor Create(AnOwner: TTreeNodes);
    function Add(const S: string): Integer; override;
    procedure Clear; override;
    procedure Delete(Index: Integer); override;
    procedure Insert(Index: Integer; const S: string); override;
    procedure LoadTreeFromStream(Stream: TStream);
    procedure SaveTreeToStream(Stream: TStream);
    function ConsistencyCheck: integer;
    procedure WriteDebugReport(const Prefix: string);
    property Owner: TTreeNodes read FOwner;
  end;

constructor TTreeStrings.Create(AnOwner: TTreeNodes);
begin
  inherited Create;
  FOwner := AnOwner;
end;

function TTreeStrings.Get(Index: Integer): string;
const
  TabChar = #9;
var
  Level, I: Integer;
  Node: TTreeNode;
begin
  Result := '';
  Node := Owner.GetNodeFromIndex(Index);
  Level := Node.Level;
  for I := 0 to Level - 1 do
    Result := Result + TabChar;
  Result := Result + Node.Text;
end;

function TTreeStrings.GetBufStart(Buffer: PChar; var Level: Integer): PChar;
begin
  Level := 0;
  while Buffer^ in [' ', #9] do
  begin
    Inc(Buffer);
    Inc(Level);
  end;
  Result := Buffer;
end;

function TTreeStrings.GetObject(Index: Integer): TObject;
begin
  Result := TObject(Owner.GetNodeFromIndex(Index).Data);
end;

procedure TTreeStrings.PutObject(Index: Integer; AObject: TObject);
begin
  Owner.GetNodeFromIndex(Index).Data := AObject;
end;

function TTreeStrings.GetCount: Integer;
begin
  Result := Owner.Count;
end;

procedure TTreeStrings.Clear;
begin
  Owner.Clear;
end;

procedure TTreeStrings.Delete(Index: Integer);
begin
  Owner.GetNodeFromIndex(Index).Delete;
end;

procedure TTreeStrings.SetUpdateState(Updating: Boolean);
begin
  //SendMessage(Owner.Handle, WM_SETREDRAW, Ord(not Updating), 0);
  if not Updating then
    Owner.Owner.Invalidate;
end;

function TTreeStrings.Add(const S: string): Integer;
var
  Level, OldLevel, I: Integer;
  NewStr: string;
  Node: TTreeNode;
begin
  Result := GetCount;
  if (Length(S) = 1) and (S[1] = Chr($1A)) then Exit;
  Node := nil;
  OldLevel := 0;
  NewStr := GetBufStart(PChar(S), Level);
  if Result > 0 then
  begin
    Node := Owner.GetNodeFromIndex(Result - 1);
    OldLevel := Node.Level;
  end;
  if (Level > OldLevel) or (Node = nil) then
  begin
    if Level - OldLevel > 1 then
      TreeViewError('TTreeStrings.Add: Invalid level, Level='+IntToStr(Level)
        +' OldLevel='+IntToStr(OldLevel));
  end
  else begin
    for I := OldLevel downto Level do
    begin
      Node := Node.Parent;
      if (Node = nil) and (I - Level > 0) then
        TreeViewError('TTreeStrings.Add: Invalid level, Node=nil I='+IntToStr(I)
          +' Level='+IntToStr(Level));
    end;
  end;
  Owner.AddChild(Node, NewStr);
end;

procedure TTreeStrings.Insert(Index: Integer; const S: string);
begin
  with Owner do
    Insert(GetNodeFromIndex(Index), S);
end;

procedure TTreeStrings.LoadTreeFromStream(Stream: TStream);
var
  List: TStringList;
  ANode, NextNode: TTreeNode;
  ALevel, i: Integer;
  CurrStr: string;
begin
  List := TStringList.Create;
  Owner.BeginUpdate;
  try
    Clear;
    List.LoadFromStream(Stream);
    ANode := nil;
    for i := 0 to List.Count - 1 do
    begin
      CurrStr := GetBufStart(PChar(List[i]), ALevel);
      if ANode = nil then
        ANode := Owner.AddChild(nil, CurrStr)
      else if ANode.Level = ALevel then
        ANode := Owner.AddChild(ANode.Parent, CurrStr)
      else if ANode.Level = (ALevel - 1) then
        ANode := Owner.AddChild(ANode, CurrStr)
      else if ANode.Level > ALevel then
      begin
        NextNode := ANode.Parent;
        while NextNode.Level > ALevel do
          NextNode := NextNode.Parent;
        ANode := Owner.AddChild(NextNode.Parent, CurrStr);
      end
      else TreeViewError('TTreeStrings.LoadTreeFromStream: Level='
        +IntToStr(ALevel)+' CuurStr="'+CurrStr+'"');
    end;
  finally
    Owner.EndUpdate;
    List.Free;
  end;
end;

procedure TTreeStrings.SaveTreeToStream(Stream: TStream);
const
  TabChar = #9;
  EndOfLine = #13#10;
var
  i: Integer;
  ANode: TTreeNode;
  NodeStr: string;
begin
  if Count > 0 then
  begin
    ANode := Owner[0];
    while ANode <> nil do
    begin
      NodeStr := '';
      for i := 0 to ANode.Level - 1 do NodeStr := NodeStr + TabChar;
      NodeStr := NodeStr + ANode.Text + EndOfLine;
      Stream.Write(Pointer(NodeStr)^, Length(NodeStr));
      ANode := ANode.GetNext;
    end;
  end;
end;

function TTreeStrings.ConsistencyCheck: integer;
begin
  Result:=0;
end;

procedure TTreeStrings.WriteDebugReport(const Prefix: string);
begin
  DebugLn('%sTTreeStrings.WriteDebugReport Self=%p Consistency=%d', [Prefix, Pointer(Self), ConsistencyCheck]);
end;


{ TCustomTreeView }

constructor TCustomTreeView.Create(AnOwner: TComponent);
var
  Details: TThemedElementDetails;
begin
  inherited Create(AnOwner);
  ControlStyle := ControlStyle - [csCaptureMouse]
                               + [csDisplayDragImage, csReflector];
  Width := 121;
  Height := 97;
  Color := clWindow;

  FSBVertShowing:=-1;
  FSBHorzShowing:=-1;
  TabStop := True;
  ParentColor := False;
//  FBackgroundColor := clWindow;
  FDefItemHeight := DefaultTreeNodeHeight;
  FDefItemSpace := ScaleY(2, 96);
  FExpandSignType := tvestTheme;
  FExpandSignSize := -1;
  Details := ThemeServices.GetElementDetails(ttGlyphOpened);
  FThemeExpandSignSize := ThemeServices.GetDetailSize(Details).cx;
  FTreeNodes := CreateNodes;
  BorderStyle := bsSingle;
  BorderWidth := 0;
  FMultiSelectStyle := DefaultMultiSelectStyle;
  FOptions := DefaultTreeViewOptions;
  Items.KeepCollapsedNodes:=KeepCollapsedNodes;
  FScrollBars:=ssBoth;
  FDragImage := TDragImageList.CreateSize(32, 32);
  FIndent:=-1;
  FChangeTimer := TTimer.Create(Self);
  FChangeTimer.Enabled := False;
  FChangeTimer.Interval := 1;
  FChangeTimer.OnTimer := @OnChangeTimer;
  FImageChangeLink := TChangeLink.Create;
  FImageChangeLink.OnChange := @ImageListChange;
  FSelectedColor:=clHighlight;
  FSelectedFontColor:=clWhite;
  FSelectedFontColorUsed:=false;
  fSeparatorColor:=clGray;
  FStateChangeLink := TChangeLink.Create;
  FStateChangeLink.OnChange := @ImageListChange;
  FStates:=[tvsMaxLvlNeedsUpdate,tvsMaxRightNeedsUpdate,tvsScrollbarChanged];
  FTreeLineColor := clWindowFrame;
  FTreeLinePenStyle := psPattern;
  SetLength(FTreeLinePenPattern, 2);
  FTreeLinePenPattern[0] := 1;
  FTreeLinePenPattern[1] := 1;
  FExpandSignColor := clWindowFrame;
  FHotTrackColor := clNone;
  // Accessibility
  AccessibleDescription := rsTTreeViewAccessibilityDescription;
  AccessibleRole := larTreeView;
  FAccessibilityOn := WidgetSet.GetLCLCapability(lcAccessibilitySupport) = LCL_CAPABILITY_YES;

  FDragScrollMargin := 20; // Height of 2 areas, on top/bottom of treeview, which auto-scroll treeview up/down
  FDragScrollTimer := TTimer.Create(Self);
  FDragScrollTimer.Enabled := false;
  FDragScrollTimer.Interval := 150;
  FDragScrollTimer.OnTimer := @DragScrollTimerTick;
end;

destructor TCustomTreeView.Destroy;
begin
  LockSelectionChangeEvent; // prevent change event during destroying

  Images:=nil;
  FreeThenNil(FTreeNodes);
  FreeThenNil(FSaveItems);
  FreeThenNil(FDragImage);
  FreeThenNil(FImageChangeLink);
  FreeThenNil(FStateChangeLink);
  inherited Destroy;
end;

procedure TCustomTreeView.CreateWnd;
begin
  Exclude(FStates,tvsStateChanging);
  FSBHorzShowing:=-1;
  FSBVertShowing:=-1;
  inherited CreateWnd;
end;

procedure TCustomTreeView.InitializeWnd;
begin
  inherited InitializeWnd;
  UpdateDefaultItemHeight;
end;

procedure TCustomTreeView.Invalidate;
begin
  if tvsPainting in FStates then exit;
  inherited Invalidate;
end;

procedure TCustomTreeView.EraseBackground(DC: HDC);
begin
  // everything is painted, so erasing the background is not needed
end;

procedure TCustomTreeView.DestroyWnd;
begin
  Include(FStates, tvsStateChanging);
  inherited DestroyWnd;
  if Canvas <> nil then
    TControlCanvas(Canvas).FreeHandle;
  FLastHorzScrollInfo.cbSize := 0;
  FLastVertScrollInfo.cbSize := 0;
end;

procedure TCustomTreeView.DoAutoAdjustLayout(
  const AMode: TLayoutAdjustmentPolicy; const AXProportion, AYProportion: Double
  );
begin
  inherited DoAutoAdjustLayout(AMode, AXProportion, AYProportion);

  if AMode in [lapAutoAdjustWithoutHorizontalScrolling, lapAutoAdjustForDPI] then
  begin
    DisableAutoSizing{$IFDEF DebugDisableAutoSizing}('TCustomTreeView.DoAutoAdjustLayout'){$ENDIF};
    try
      if DefaultItemHeightIsStored then
        DefaultItemHeight := Round(DefaultItemHeight*AYProportion);
      if IndentIsStored then
        FIndent := Round(FIndent*AXProportion);
      if ExpandSignSizeIsStored then
        FExpandSignSize := Round(FExpandSignSize*AXProportion);
    finally
      EnableAutoSizing{$IFDEF DebugDisableAutoSizing}('TCustomTreeView.DoAutoAdjustLayout'){$ENDIF};
    end;
  end;
end;

procedure TCustomTreeView.DoCreateNodeClass(var NewNodeClass: TTreeNodeClass);
begin
  if Assigned(OnCreateNodeClass) then
    OnCreateNodeClass(Self,NewNodeClass);
end;

procedure TCustomTreeView.BeginAutoDrag;
begin
  BeginDrag(False);
end;

procedure TCustomTreeView.BeginEditing(ANode: TTreeNode);
var
  ARect: TRect;
begin
  //DebugLn(['TCustomTreeView.BeginEditing tvsIsEditing=',tvsIsEditing in FStates,' Selected=',Selected<>nil]);
  if (tvsIsEditing in FStates) or (ANode=nil) then exit;
  if (not CanEdit(ANode)) or (not ANode.Visible) then exit;
  // if we are asked to edit another node while one is already being edited then
  // stop editing that node
  if FEditingItem <> nil then
    EndEditing;
  FEditingItem := ANode;
  // make node visible (this will cancel editing, so call this first)
  EnsureNodeIsVisible(ANode);

  Include(FStates,tvsIsEditing);
  if FEditor=nil then begin
    FEditor:=TEdit.Create(Self);
    FEditor.OnEditingDone:=@EditorEditingDone;
    FEditor.OnKeyDown:=@EditorKeyDown;
  end;
  ARect:=Rect(Max(BorderWidth,ANode.DisplayTextLeft),ANode.Top-ScrolledTop,
              ClientWidth-BorderWidth,ANode.Bottom-ScrolledTop);
  FEditor.BoundsRect:=ARect;
  FEditor.AnchorParallel(akLeft,ARect.Left,Self);
  FEditor.AnchorParallel(akRight,BorderWidth,Self);
  FEditor.Visible:=true;
  FEditor.Parent:=Self;
  FEditor.Text:=ANode.Text;
  FEditor.SelectAll;
  FEditor.SetFocus;
end;

procedure TCustomTreeView.BeginUpdate;
begin
  Items.BeginUpdate;
  LockSelectionChangeEvent;
end;

procedure TCustomTreeView.EndUpdate;
begin
  UnlockSelectionChangeEvent;
  if Items.FUpdateCount<=0 then RaiseGDBException('TCustomTreeView.EndUpdate');
  Items.EndUpdate;
  if Items.FUpdateCount=0 then begin
    // ToDo: only refresh if something changed
    UpdateScrollBars;
  end;
end;

function TCustomTreeView.AlphaSort: Boolean;
begin
  Result := CustomSort(nil);
end;

function TCustomTreeView.CustomSort(SortProc: TTreeNodeCompare): Boolean;
var Node: TTreeNode;
begin
  Result := False;
  if FTreeNodes.Count>0 then begin
    BeginUpdate;
    if not assigned(SortProc) then SortProc := @DefaultTreeViewSort;
    FTreeNodes.SortTopLevelNodes(SortProc);

    Node := FTreeNodes.GetFirstNode;
    while Node <> nil do begin
      if (Node.GetFirstChild<>nil) then Node.CustomSort(SortProc);
      Node := Node.GetNext;
    end;
    Items.ClearCache;
    FStates:= FStates+[tvsTopsNeedsUpdate, tvsTopItemNeedsUpdate,
                       tvsBottomItemNeedsUpdate,tvsScrollbarChanged];
    EndUpdate;
  end;
end;

function TCustomTreeView.DefaultItemHeightIsStored: Boolean;
begin
  Result := not(tvoAutoItemHeight in Options);
end;

function TCustomTreeView.DefaultTreeViewSort(Node1, Node2: TTreeNode): Integer;
begin
  if Assigned(OnCompare) then begin
    Result:=0;
    OnCompare(Node1.TreeView,Node1, Node2, Result);
  end else
    Result := Utf8CompareStr(Node1.Text,Node2.Text);
end;

procedure TCustomTreeView.SetAutoExpand(Value: Boolean);
begin
  if AutoExpand <> Value then begin
    if Value then
      Include(FOptions,tvoAutoExpand)
    else
      Exclude(FOptions,tvoAutoExpand);
  end;
end;

procedure TCustomTreeView.SetHotTrack(Value: Boolean);
begin
  if HotTrack <> Value then begin
    if Value then
      Include(FOptions,tvoHotTrack)
    else
      Exclude(FOptions,tvoHotTrack);
  end;
end;

procedure TCustomTreeView.SetRowSelect(Value: Boolean);
begin
  if RowSelect <> Value then begin
    if Value then
      Include(FOptions,tvoRowSelect)
    else
      Exclude(FOptions,tvoRowSelect);
    if FSelectedNode<>nil then
      Invalidate;
  end;
end;

procedure TCustomTreeView.SetScrollBars(const Value: TScrollStyle);
begin
  if (FScrollBars <> Value) then begin
    FScrollBars := Value;
    Include(FStates,tvsScrollbarChanged);
    UpdateScrollBars;
  end;
end;

procedure TCustomTreeView.SetScrolledLeft(AValue: integer);
var
  OldScrolledLeft: Integer;
begin
  OldScrolledLeft := FScrolledLeft;
  if AValue<0 then AValue:=0;
  if AValue=FScrolledLeft then exit;
  if AValue>GetMaxScrollLeft then AValue:=GetMaxScrollLeft;
  if AValue=FScrolledLeft then exit;
  EndEditing(true);
  FScrolledLeft:=AValue;
  ScrollView(OldScrolledLeft-FScrolledLeft, 0);
end;

procedure TCustomTreeView.SetScrolledTop(AValue: integer);
var
  OldScrolledTop: Integer;
begin
  OldScrolledTop:=FScrolledTop;
  if FScrolledTop=AValue then exit;
  if AValue<0 then AValue:=0;
  if AValue>GetMaxScrollTop then AValue:=GetMaxScrollTop;
  if AValue=FScrolledTop then exit;
  EndEditing(true);
  FScrolledTop:=AValue;
  FStates:=FStates+[tvsTopItemNeedsUpdate,tvsBottomItemNeedsUpdate];
  ScrollView(0, OldScrolledTop-FScrolledTop);
end;

procedure TCustomTreeView.SetToolTips(Value: Boolean);
begin
  if ToolTips <> Value then begin
    if Value then
      Include(FOptions,tvoToolTips)
    else
      Exclude(FOptions,tvoToolTips);
  end;
end;

procedure TCustomTreeView.SetSortType(Value: TSortType);
begin
  if SortType <> Value then begin
    FSortType := Value;
    if ((SortType in [stData, stBoth]) and Assigned(OnCompare)) or
      (SortType in [stText, stBoth]) then
      AlphaSort;
  end;
end;

procedure TCustomTreeView.SetBackgroundColor(Value: TColor);
begin
  if Color<>Value then begin
    Color:=Value;
    Invalidate;
  end;
end;

procedure TCustomTreeView.SetSelectedColor(Value: TColor);
begin
  if FSelectedColor<>Value then begin
    FSelectedColor:=Value;
    Invalidate;
  end;
end;

procedure TCustomTreeView.SetSelectedFontColor(Value: TColor);
begin
  if FSelectedFontColor<>Value then begin
    FSelectedFontColor:=Value;
    Invalidate;
  end;
end;

procedure TCustomTreeView.Paint;
begin
  DoPaint;
end;

procedure TCustomTreeView.SetDragMode(Value: TDragMode);
begin
  inherited SetDragMode(Value);
end;

procedure TCustomTreeView.SetOptions(NewOptions: TTreeViewOptions);
var ChangedOptions: TTreeViewOptions;
begin
  if FOptions=NewOptions then exit;
  ChangedOptions:=(FOptions-NewOptions)+(NewOptions-FOptions);
  FOptions:=NewOptions;
  if tvoKeepCollapsedNodes in ChangedOptions then
    Items.KeepCollapsedNodes:=(tvoKeepCollapsedNodes in FOptions);
  if (tvoReadOnly in ChangedOptions) and (not (tvoReadOnly in FOptions)) then
    EndEditing;
  if (tvoAllowMultiselect in ChangedOptions) then begin
    if (tvoAllowMultiselect in FOptions) then begin
      if Selected<>nil then
        Selected.MultiSelected:=true;
    end else begin
      Items.ClearMultiSelection;
    end;
  end;
  if tvoAutoItemHeight in ChangedOptions then
    UpdateDefaultItemHeight;
  if ([tvoHideSelection,tvoReadOnly,tvoShowButtons,tvoShowRoot,tvoShowLines]
    * ChangedOptions)<>[]
  then
    Invalidate;
end;

procedure TCustomTreeView.UpdateDefaultItemHeight;
var
  NewDefItemHeight: Integer;
  ImageSize, StateImageSize: TSize;
begin
  if (tvoAutoItemHeight in FOptions)
  and HandleAllocated and Canvas.HandleAllocated then begin
    NewDefItemHeight:=Canvas.TextHeight(TVAutoHeightString)+FDefItemSpace;
    if Assigned(FImages) then
      ImageSize := Images.SizeForPPI[ImagesWidth, Font.PixelsPerInch];
    if Assigned(FStateImages) then
      StateImageSize := StateImages.SizeForPPI[StateImagesWidth, Font.PixelsPerInch];
    if NewDefItemHeight<FDefItemSpace then NewDefItemHeight:=FDefItemSpace;
    if (Images<>nil) and (ImageSize.cy+FDefItemSpace>NewDefItemHeight) then
      NewDefItemHeight:=ImageSize.cy+FDefItemSpace;
    if (StateImages<>nil) and (StateImageSize.cy+FDefItemSpace>NewDefItemHeight) then
      NewDefItemHeight:=StateImageSize.cy+FDefItemSpace;
    if Odd(NewDefItemHeight) then Inc(NewDefItemHeight);
    if NewDefItemHeight<>FDefItemHeight then begin
      FDefItemHeight:=NewDefItemHeight;
      FStates:=FStates+[tvsTopsNeedsUpdate,tvsTopItemNeedsUpdate,tvsBottomItemNeedsUpdate];
      Invalidate;
    end;
  end;
end;

procedure TCustomTreeView.UpdateAllTops;
var
  CurTop: integer;

  procedure CalculateTops(Node: TTreeNode);
  begin
    while Node<>nil do begin
      if Node.Visible then begin
        Node.fTop:=CurTop;
        inc(CurTop, Node.Height);
        if Node.Expanded then
          CalculateTops(Node.GetFirstChild);
      end;
      Node:=Node.GetNextSibling;
    end;
  end;

begin
  if not (tvsTopsNeedsUpdate in FStates) then exit;
  CurTop:=0;
  CalculateTops(Items.GetFirstVisibleNode);
  Exclude(FStates,tvsTopsNeedsUpdate);
  Include(FStates,tvsScrollbarChanged);
end;

procedure TCustomTreeView.UpdateMaxLvl;

  procedure LookInChildsAndBrothers(Node: TTreeNode; CurLvl: integer);
  begin
    if Node=nil then exit;
    if CurLvl>FMaxLvl then FMaxLvl:=CurLvl;
    LookInChildsAndBrothers(Node.GetFirstChild,CurLvl+1);
    LookInChildsAndBrothers(Node.GetNextSibling,CurLvl);
  end;

begin
  if not (tvsMaxLvlNeedsUpdate in FStates) then exit;
  FMaxLvl:=0;
  LookInChildsAndBrothers(Items.GetFirstNode,1);
  Exclude(FStates,tvsMaxRightNeedsUpdate);
end;

procedure TCustomTreeView.UpdateMaxRight;
const
  LargeItemCount = 100;
  ReservedWidth = 100;
var
  Node: TTreeNode;
  i: integer;
  FMaxTextLen, AIndent: Integer;
  Cnt: Integer;
begin
  if not (tvsMaxRightNeedsUpdate in FStates) then exit;
  FMaxRight := 0;
  FMaxTextLen := 0;
  Node := Items.GetFirstNode;
  Cnt := 0;
  AIndent := Indent;
  while Node <> nil do
  begin
    if not Node.AreParentsExpandedAndVisible then
    begin
      Node := Node.GetNext;
      Continue;
    end;
    inc(Cnt);
    if (Cnt < LargeItemCount) then
    begin
      i := Node.DisplayTextRight + ScrolledLeft + AIndent div 2;
    end else
    begin
      // computing DisplayTextRight is too expensive when the tree
      // has hundreds of nodes
      // => use a heuristic
      if length(Node.Text) > FMaxTextLen then
        i := Node.DisplayTextRight + ScrolledLeft + ReservedWidth
      else
        i := FMaxRight;
    end;
    if FMaxRight < i then
    begin
      FMaxRight := i;
      FMaxTextLen := length(Node.Text);
    end;
    Node := Node.GetNext;
  end;
  Exclude(FStates, tvsMaxRightNeedsUpdate);
  Include(FStates, tvsScrollbarChanged);
end;

procedure TCustomTreeView.UpdateTopItem;
begin
  //DebugLn('TCustomTreeView.UpdateTopItem tvsTopItemNeedsUpdate in FStates=',tvsTopItemNeedsUpdate in FStates);
  if (FStates*[tvsTopItemNeedsUpdate,tvsTopsNeedsUpdate]=[]) then exit;
  FTopItem:=GetNodeAtY(BorderWidth);
  Exclude(FStates,tvsTopItemNeedsUpdate);
end;

procedure TCustomTreeView.UpdateBottomItem;
begin
  if (FStates*[tvsTopItemNeedsUpdate,tvsTopsNeedsUpdate,tvsBottomItemNeedsUpdate]=[])
  then exit;
//  if not (tvsBottomItemNeedsUpdate in FStates) then exit;  already above
  FBottomItem:=TopItem;
  while (FBottomItem<>nil) and (FBottomItem.GetNextVisible<>nil) do
    FBottomItem:=FBottomItem.GetNextVisible;
  Exclude(FStates,tvsBottomItemNeedsUpdate);
end;

procedure TCustomTreeView.SetBottomItem(Value: TTreeNode);
begin
  if HandleAllocated and (Value <> nil) then begin
    Value.MakeVisible;
    ScrolledTop:=Value.Top+Value.Height-ClientHeight;
  end;
end;

procedure TCustomTreeView.SetSeparatorColor(const AValue: TColor);
begin
  if fSeparatorColor=AValue then exit;
  fSeparatorColor:=AValue;
  if tvoShowSeparators in Options then
    Invalidate;
end;

procedure TCustomTreeView.SetShowButton(Value: Boolean);
begin
  if ShowButtons <> Value then begin
    if Value then
      Include(FOptions,tvoShowButtons)
    else
      Exclude(FOptions,tvoShowButtons);
    Invalidate;
  end;
end;

procedure TCustomTreeView.SetShowLines(Value: Boolean);
begin
  if ShowLines <> Value then begin
    if Value then
      Include(FOptions,tvoShowLines)
    else
      Exclude(FOptions,tvoShowLines);
    Invalidate;
  end;
end;

procedure TCustomTreeView.SetShowRoot(Value: Boolean);
begin
  if ShowRoot <> Value then begin
    if Value then
      Include(FOptions,tvoShowRoot)
    else
      Exclude(FOptions,tvoShowRoot);
    Invalidate;
  end;
end;

procedure TCustomTreeView.SetShowScrollBar(Which: Integer; AShow: Boolean);
begin
  if ((Which in [SB_Horz, SB_BOTH]) and (FSBHorzShowing<>Ord(AShow)))
  or ((Which in [SB_Vert, SB_BOTH]) and (FSBVertShowing<>Ord(AShow)))
  then
    ShowScrollBar(Handle, Which, AShow);

  if Which in [SB_Horz, SB_BOTH] then
    FSBHorzShowing:=Ord(AShow);
  if Which in [SB_Vert, SB_BOTH] then
    FSBVertShowing:=Ord(AShow);
end;

procedure TCustomTreeView.SetShowSeparators(Value: Boolean);
begin
  if ShowSeparators <> Value then begin
    if Value then
      Include(FOptions,tvoShowSeparators)
    else
      Exclude(FOptions,tvoShowSeparators);
    Invalidate;
  end;
end;

procedure TCustomTreeView.SetKeepCollapsedNodes(Value: Boolean);
begin
  if KeepCollapsedNodes=Value then exit;
  if Value then
    Include(FOptions,tvoKeepCollapsedNodes)
  else
    Exclude(FOptions,tvoKeepCollapsedNodes);
  Items.KeepCollapsedNodes:=Value;
end;

procedure TCustomTreeView.SetMultiSelect(const AValue: Boolean);
begin
  if MultiSelect <> AValue then
  begin
    if AValue then
      Include(FOptions,tvoAllowMultiselect)
    else
      Exclude(FOptions,tvoAllowMultiselect);
  end;
end;

procedure TCustomTreeView.SetMultiSelectStyle(const AValue: TMultiSelectStyle);
begin
  if FMultiSelectStyle=AValue then exit;
  FMultiSelectStyle:=AValue;
  // there must be at least one multiselectstyle according to docs
  // http://docs.embarcadero.com/products/rad_studio/delphiAndcpp2009/HelpUpdate2/EN/html/delphivclwin32/ComCtrls_TCustomTreeView_MultiSelectStyle.html
  if FMultiSelectStyle = [] then
    FMultiSelectStyle := DefaultMultiSelectStyle;
end;

procedure TCustomTreeView.SetReadOnly(Value: Boolean);
begin
  if ReadOnly = Value then exit;
  if Value then
    Include(FOptions,tvoReadOnly)
  else
    Exclude(FOptions,tvoReadOnly);
  if not Value then EndEditing;
end;

procedure TCustomTreeView.SetRightClickSelect(Value: Boolean);
begin
  if Value then
    Include(FOptions,tvoRightClickSelect)
  else
    Exclude(FOptions,tvoRightClickSelect);
end;

procedure TCustomTreeView.SetHideSelection(Value: Boolean);
begin
  if HideSelection <> Value then begin
    if Value then
      Include(FOptions,tvoHideSelection)
    else
      Exclude(FOptions,tvoHideSelection);
    if FSelectedNode<>nil then Invalidate;
  end;
end;

function TCustomTreeView.GetMaxLvl: integer;
begin
  UpdateMaxRight;
  Result:=FMaxRight;
end;

function TCustomTreeView.GetMaxScrollLeft: integer;
begin
  UpdateMaxRight;
  Result:=FMaxRight-(ClientWidth-2*BorderWidth);
  if Result<0 then Result:=0;
end;

function TCustomTreeView.GetMaxScrollTop: integer;
var
  LastVisibleNode: TTreeNode;
begin
  LastVisibleNode:=Items.GetLastExpandedSubNode;
  if LastVisibleNode=nil then
    Result:=0
  else begin
    Result:=LastVisibleNode.Top+LastVisibleNode.Height
             -ClientHeight+2*integer(BorderWidth);
    //DebugLn('>>> ',LastVisibleNode.Text,' ',Result);
    if Result<0 then Result:=0;
  end;
end;

function TCustomTreeView.GetNodeAtY(Y: Integer): TTreeNode;
// search in all expanded nodes for the node at the screen coordinate Y
var
  i: integer;
begin
  Result := nil;
  if not Assigned(Items) then
    Exit;
  if (Y >= BorderWidth) and (Y < ClientHeight - BorderWidth) then
  begin
    inc(Y, FScrolledTop - BorderWidth);
    i := IndexOfNodeAtTop(Items.FTopLvlItems, Items.FTopLvlCount, Y);
    if i >= 0 then
    begin
      Result := Items.FTopLvlItems[i];
      while Result.Visible and Result.Expanded do
      begin
        i := IndexOfNodeAtTop(Result.FItems, Result.FCount, Y);
        if i >= 0 then
          Result := Result.Items[i]
        else
          break;
      end;
    end;
  end;
end;

function TCustomTreeView.GetNodeDrawAreaWidth: integer;
begin
  Result:=ClientWidth-BorderWidth*2;
end;

function TCustomTreeView.GetNodeDrawAreaHeight: integer;
begin
  Result:=ClientHeight-BorderWidth*2;
end;

function TCustomTreeView.GetNodeAt(X, Y: Integer): TTreeNode;
begin
  if (X >= BorderWidth) and (X < ClientWidth - BorderWidth) then
    Result := GetNodeAtY(Y)
  else
    Result := nil;
end;

procedure TCustomTreeView.GetInsertMarkAt(X, Y: Integer;
  out AnInsertMarkNode: TTreeNode; out AnInsertMarkType: TTreeViewInsertMarkType
  );
var
  ANode: TTreeNode;
  NodeRect: TRect;
  NodeMidY: integer;
begin
  AnInsertMarkNode:=nil;
  AnInsertMarkType:=tvimNone;
  if Y<0 then Y:=0;
  if Y>=ClientHeight then Y:=ClientHeight-1;
  ANode:=GetNodeAtY(Y);
  if ANode<>nil then
  begin
    NodeRect:=ANode.DisplayRect(false);
    NodeMidY:=NodeRect.Top + (NodeRect.Bottom-NodeRect.Top) div 2;
    AnInsertMarkNode:=ANode;
    if (X>AnInsertMarkNode.DisplayExpandSignRight) then
      if Y>=NodeMidY then begin
        // insert as first child of pointed node
        AnInsertMarkType:=tvimAsFirstChild;
      end else begin
        // insert as previous sibling of pointed node
        AnInsertMarkType:=tvimAsPrevSibling;
    end else begin
      if Y>=NodeMidY then begin
        if ANode.Expanded and ANode.HasChildren then begin
          // insert as first child of pointed node
          AnInsertMarkType:=tvimAsFirstChild;
        end else begin
          // insert as next sibling of pointed node
          AnInsertMarkType:=tvimAsNextSibling;
        end;
      end else begin
        // insert as previous sibling of pointed node
        AnInsertMarkType:=tvimAsPrevSibling;
      end;
    end;
  end else begin
    // insert behind all nodes
    ANode:=Items.GetLastExpandedSubNode;
    if ANode<>nil then begin
     AnInsertMarkNode:=ANode;
     if X>AnInsertMarkNode.DisplayExpandSignRight then
       // insert as first child of last visible node
       AnInsertMarkType:=tvimAsFirstChild
     else
       // insert as next sibling of last visible node
       AnInsertMarkType:=tvimAsNextSibling;
    end else begin
      // insert as new root
      AnInsertMarkNode:=nil;
      AnInsertMarkType:=tvimAsFirstChild;
    end;
  end;

  // normalize (try to replace tvimAsPrevSibling)
  if (AnInsertMarkType=tvimAsPrevSibling) and (AnInsertMarkNode<>nil) then begin
    if (AnInsertMarkNode.GetPrevSibling<>nil) then begin
      if (AnInsertMarkNode.GetPrevSibling.Expanded=false)
      and (AnInsertMarkNode.GetPrevSibling.IsVisible) then begin
        AnInsertMarkNode:=AnInsertMarkNode.GetPrevSibling;
        AnInsertMarkType:=tvimAsNextSibling;
      end;
    end else if (AnInsertMarkNode.Parent<>nil)
    and (AnInsertMarkNode.Parent.IsVisible) then begin
      AnInsertMarkNode:=AnInsertMarkNode.Parent;
      AnInsertMarkType:=tvimAsFirstChild;
    end;
  end;
end;

procedure TCustomTreeView.SetInsertMark(AnInsertMarkNode: TTreeNode;
  AnInsertMarkType: TTreeViewInsertMarkType);
begin
  InsertMarkNode:=AnInsertMarkNode;
  InsertMarkType:=AnInsertMarkType;
end;

procedure TCustomTreeView.SetInsertMarkAt(X, Y: integer);
var
  AnInsertMarkNode: TTreeNode;
  AnInsertMarkType: TTreeViewInsertMarkType;
begin
  GetInsertMarkAt(X,Y,AnInsertMarkNode,AnInsertMarkType);
  SetInsertMark(AnInsertMarkNode,AnInsertMarkType);
end;

function TCustomTreeView.GetHitTestInfoAt(X, Y: Integer): THitTests;
// ToDo: Set also flags [htAbove, htBelow, htOnRight, htToLeft, htToRight];
var
  Node: TTreeNode;
begin
  Result := [];
  if (X>=0) and (X<ClientWidth) and (Y>=0) and (Y<ClientHeight)
  then begin
    Node:=GetNodeAtY(Y);
    if Node<>nil then begin
      Include(Result,htOnItem);
      if X<Node.DisplayExpandSignLeft then
        Include(Result,htOnIndent)
      else if X<Node.DisplayStateIconLeft then
        Include(Result,htOnButton)
      else if X<Node.DisplayIconLeft then
        Include(Result,htOnStateIcon)
      else if X<Node.DisplayTextLeft then
        Include(Result,htOnIcon)
      else if X<Node.DisplayTextRight then
        Include(Result,htOnLabel);
    end else
      Include(Result,htNowhere);
  end;
end;

procedure TCustomTreeView.SetTreeLineColor(Value: TColor);
begin
  if FTreeLineColor<>Value then begin
    FTreeLineColor:=Value;
    Invalidate;
  end;
end;

procedure TCustomTreeView.SetTreeNodes(Value: TTreeNodes);
begin
  Items.Assign(Value);
end;

procedure TCustomTreeView.SetIndent(Value: Integer);
begin
  if Value <> Indent then begin
    FIndent := Value;
    Invalidate;
  end;
end;

procedure TCustomTreeView.FullExpand;
var
  Node: TTreeNode;
begin
  Node := Items.GetFirstNode;
  while Node <> nil do begin
    Node.Expand(True);
    Node := Node.GetNextSibling;
  end;
end;

procedure TCustomTreeView.FullCollapse;
var
  Node: TTreeNode;
begin
  Node := Items.GetFirstNode;
  while Node <> nil do begin
    Node.Collapse(True);
    Node := Node.GetNextSibling;
  end;
end;

function TCustomTreeView.IsNodeVisible(ANode: TTreeNode): Boolean;
begin
  Result:=(ANode<>nil) and (ANode.Visible) and (ANode.AreParentsExpandedAndVisible);
  if Result then begin
    //DebugLn('[TCustomTreeView.IsNodeVisible] B Node=',DbgS(ANode),
    //  ' ',dbgs(FScrolledTop)+'>=',dbgs(ANode.Top+ANode.Height)+' or =',dbgs(FScrolledTop),'+'+dbgs(ClientHeight)+'<',dbgs(ANode.Top));
    if (FScrolledTop>=ANode.Top+ANode.Height)
    or (FScrolledTop+ClientHeight-2*BorderWidth<ANode.Top)
    then
      Result:=false;
  end;
  //DebugLn('[TCustomTreeView.IsNodeVisible] END Node=',DbgS(ANode),
  //' Node.Text=',ANode.Text,' Visible=',Result);
end;

function TCustomTreeView.IsNodeHeightFullVisible(ANode: TTreeNode): Boolean;
begin
  Result:=(ANode<>nil) and (ANode.AreParentsExpandedAndVisible);
  if Result then begin
    //DebugLn('[TCustomTreeView.IsNodeVisible] B Node=',DbgS(ANode),
    //' ',FScrolledTop,'>=',ANode.Top,'+',ANode.Height,' or ',FScrolledTop,'+',ClientHeight,'<',ANode.Top);
    if (FScrolledTop>ANode.Top)
    or (FScrolledTop+ClientHeight-2*BorderWidth
      <ANode.Top+ANode.Height)
    then
      Result:=false;
  end;
  //DebugLn('[TCustomTreeView.IsNodeVisible] END Node=',DbgS(ANode),
  //' Node.Text=',ANode.Text,' Visible=',Result);
end;

procedure TCustomTreeView.KeyDown(var Key: Word; Shift: TShiftState);
const
  EditKey = VK_F2;
  EditKeyShift = [];
var
  lNode: TTreeNode;
begin
  inherited KeyDown(Key, Shift);

  case Key of

  VK_DOWN:
  begin
    MoveToNextNode(AllowMultiSelectWithShift(Shift));
    Key:=VK_UNKNOWN;
  end;

  VK_UP:
  begin
    MoveToPrevNode(AllowMultiSelectWithShift(Shift));
    Key:=VK_UNKNOWN;
  end;

  VK_HOME:
  begin
    MoveHome(AllowMultiSelectWithShift(Shift));
    Key:=VK_UNKNOWN;
  end;

  VK_END:
  begin
    MoveEnd(AllowMultiSelectWithShift(Shift));
    Key:=VK_UNKNOWN;
  end;

  VK_PRIOR: // Page Up
  begin
    MovePageUp(AllowMultiSelectWithShift(Shift));
    Key:=VK_UNKNOWN;
  end;

  VK_NEXT: // Page Down
  begin
    MovePageDown(AllowMultiSelectWithShift(Shift));
    Key:=VK_UNKNOWN;
  end;

  else
    if not (Key in [VK_LEFT,VK_RIGHT,VK_ADD,VK_SUBTRACT]) then
      Exit;

    if (tvoAllowMultiSelect in FOptions) and AllowMultiSelectWithShift(Shift) then
      lNode := FTreeNodes.FLastMultiSelected
    else
      lNode := Selected;

    case Key of
    VK_LEFT:
    if lNode <> nil then
    begin
      if lNode.Expanded then
        lNode.Expanded := False
      else
      if lNode.Parent <> nil then
        lNode := lNode.Parent;
      Key:=VK_UNKNOWN;
    end;

    VK_RIGHT:
    if lNode <> nil then
    begin
      if lNode.Expanded then
        lNode := lNode.GetNextExpanded
      else
        lNode.Expanded := True;
      Key:=VK_UNKNOWN;
    end;

    VK_ADD:
    if lNode <> nil then
      lNode.Expanded := True;

    VK_SUBTRACT:
    if lNode <> nil then
      lNode.Expanded := False;

    else
      if (Key=EditKey) and (Shift=EditKeyShift) and (not ReadOnly) then
        BeginEditing(Selected);
      lNode := nil; { No change in selection }
    end;

    if lNode <> nil then
      MoveSelection(lNode, AllowMultiSelectWithShift(Shift));
  end;
end;

procedure TCustomTreeView.Loaded;
begin
  inherited Loaded;
  if csDesigning in ComponentState then FullExpand;
  UpdateDefaultItemHeight;
end;

function TCustomTreeView.GetTopItem: TTreeNode;
begin
  if HandleAllocated then
  begin
    UpdateTopItem;
    Result := FTopItem;
  end
  else
    Result := FTopItem;
end;

procedure TCustomTreeView.HintMouseLeave(Sender: TObject);
begin
  if FindLCLControl(Mouse.CursorPos)<>Self then
    FHintWnd.Hide;
end;

function TCustomTreeView.IsStoredBackgroundColor: Boolean;
begin
  result := Color <> clWindow;
end;

procedure TCustomTreeView.SetTopItem(Value: TTreeNode);
begin
  if HandleAllocated and (Value <> nil) then
  begin
    Value.MakeVisible;
    ScrolledTop:=Value.Top;
  end
  else
    FTopItem := Value;
end;

procedure TCustomTreeView.OnChangeTimer(Sender: TObject);
begin
  FChangeTimer.Enabled := False;
  //debugln('TCustomTreeView.OnChangeTimer');
  FCallingChange := True;
  try
    Change(FSelectedNode);
  finally
    FCallingChange := False;
  end;
end;

procedure TCustomTreeView.UpdateScrollbars;
var
  ScrollInfo: TScrollInfo;
  MaxScrollLeft, MaxScrollTop: Integer;
begin
  if not (tvsScrollbarChanged in FStates) then exit;

  if not HandleAllocated or (Items.FUpdateCount>0) then
    exit;

  MaxScrollLeft := GetMaxScrollLeft;
  MaxScrollTop := GetMaxScrollTop;

  //DebugLn('* TCustomTreeView.UpdateScrollbars Enter *');
  if ScrolledLeft>MaxScrollLeft then ScrolledLeft:=MaxScrollLeft;
  if ScrolledTop>MaxScrollTop then ScrolledTop:=MaxScrollTop;
  Exclude(FStates,tvsScrollbarChanged);

  if fScrollBars in [ssBoth, ssHorizontal, ssAutoBoth, ssAutoHorizontal] then
  begin
    // horizontal scrollbar
    FillChar(ScrollInfo,SizeOf(ScrollInfo),0);
    ScrollInfo.cbSize := SizeOf(ScrollInfo);
    ScrollInfo.fMask := SIF_ALL or SIF_DISABLENOSCROLL;
    ScrollInfo.nTrackPos := 0;
    ScrollInfo.nMin := 0;
    ScrollInfo.nPage := Max(1,ClientWidth-2*BorderWidth);
    ScrollInfo.nMax := Max(1,MaxScrollLeft+integer(ScrollInfo.nPage));
    ScrollInfo.nPos := Max(FScrolledLeft,0);
    if not CompareMem(@ScrollInfo,@FLastHorzScrollInfo,SizeOf(TScrollInfo))
    then begin
      if (fScrollBars in [ssAutoBoth, ssAutoHorizontal])
      and (ScrollInfo.nPage>=cardinal(ScrollInfo.nMax)) then begin
        //DebugLn(['TCustomTreeView.UpdateScrollbars Hide Horizontal.']);
        FLastHorzScrollInfo.cbSize:=0;
        SetShowScrollBar(SB_HORZ, false);
      end else begin
        //DebugLn(['TCustomTreeView.UpdateScrollbars Show Horizontal: nMin=',ScrollInfo.nMin,
        //' nMax=',ScrollInfo.nMax,' nPage=',ScrollInfo.nPage,
        //' nPos=',ScrollInfo.nPos,' GetMaxScrollLeft=',MaxScrollLeft,
        //' ClientW=',ClientWidth, ' MaxRight=',FMaxRight]);
        FLastHorzScrollInfo:=ScrollInfo;
        SetShowScrollBar(SB_HORZ, true);
        SetScrollInfo(Handle, SB_HORZ, ScrollInfo, true);
      end;
    end;
  end else begin
    FLastHorzScrollInfo.cbSize:=0;
    SetShowScrollBar(SB_HORZ,false);
  end;

  if fScrollBars in [ssBoth, ssVertical, ssAutoBoth, ssAutoVertical] then begin
    // vertical scrollbar
    ScrollInfo.cbSize := SizeOf(ScrollInfo);
    ScrollInfo.fMask := SIF_ALL or SIF_DISABLENOSCROLL;
    ScrollInfo.nTrackPos := 0;
    ScrollInfo.nMin := 0;
    ScrollInfo.nPage := Max(1,ClientHeight-FDefItemHeight);
    ScrollInfo.nMax := Max(1,MaxScrollTop+integer(ScrollInfo.nPage));
    ScrollInfo.nTrackPos := 0;
    ScrollInfo.nPos := Max(0,FScrolledTop);
    if not CompareMem(@ScrollInfo,@FLastVertScrollInfo,SizeOf(TScrollInfo))
    then begin
      if (fScrollBars in [ssAutoBoth, ssAutoVertical])
      and (ScrollInfo.nPage>=cardinal(ScrollInfo.nMax)) then begin
        //DebugLn(['TCustomTreeView.UpdateScrollbars Hide Vertical.']);
        FLastVertScrollInfo.cbSize:=0;
        SetShowScrollBar(SB_VERT, false);
      end else begin
        //DebugLn(['TCustomTreeView.UpdateScrollbars Show Vertical: nMin=',ScrollInfo.nMin,
        //' nMax=',ScrollInfo.nMax,' nPage=',ScrollInfo.nPage,
        //' nPos=',ScrollInfo.nPos,' GetMaxScrollTop=',MaxScrollTop,
        //' ClientH=',ClientHeight]);
        FLastVertScrollInfo:=ScrollInfo;
        SetShowScrollBar(SB_VERT, true);
        SetScrollInfo(Handle, SB_VERT, ScrollInfo, true);
      end;
    end;
  end else begin
    FLastVertScrollInfo.cbSize:=0;
    SetShowScrollBar(SB_VERT,false);
  end;
end;

procedure TCustomTreeView.UpdateTooltip(X, Y: integer);
var
  Node: TTreeNode;
  PHint, PLeft: TPoint;
  R, TextRect, IntRect: TRect;
  CurMonitor: TMonitor;
begin
  if not (tvoToolTips in FOptions) then exit;

  if FHintWnd=nil then
  begin
    FHintWnd:=THintWindow.Create(Self);
    FHintWnd.OnMouseLeave:=@HintMouseLeave;
  end;

  Node := GetNodeAt(X, Y);
  if Node=nil then
  begin
    FHintWnd.Hide;
    exit;
  end;

  TextRect := Rect(Node.DisplayTextLeft, Node.Top, Node.DisplayTextRight, Node.Top + Node.Height);
  OffsetRect(TextRect, 0, -ScrolledTop);
  if not PtInRect(TextRect, Point(X, Y))
  or (IntersectRect(IntRect, TextRect, ClientRect) and EqualRect(IntRect, TextRect)) then
  begin
    FHintWnd.Hide;
    exit;
  end;

  // Get max width for hint from monitor's work area.
  CurMonitor := GetParentForm(Self).Monitor;
  R := CurMonitor.WorkareaRect;
  R := FHintWnd.CalcHintRect(R.Right-R.Left, Node.Text, nil);

  if WidgetSet.GetLCLCapability(lcTransparentWindow) = LCL_CAPABILITY_YES then
  begin
    // Font is explicitly set for transparent hints, otherwise default font is used.
    if not FHintWnd.Visible then
    begin
      FHintWnd.Font.Assign(Self.Font);
      FHintWnd.Font.Color := Screen.HintFont.Color;
    end;
    // Put transparent hint exactly on the node.
    PHint := ClientToScreen(Point(TextRect.Left-1, TextRect.Top-3+BorderWidth));
  end
  else begin
    // By default put hint to the right side of node.
    PHint := ClientToScreen(Point(ClientWidth, TextRect.Top-3+BorderWidth));
    if PHint.X + R.Right > CurMonitor.BoundsRect.Right then
    begin                      // No space on the right? Put it to the left side.
      PLeft := ClientToScreen(Point(ClientRect.Left, ClientRect.Top));
      if PLeft.X >= R.Right then  // enough space on left?
        PHint.X := PLeft.X - R.Right;
    end;
  end;
  OffsetRect(R, PHint.X, PHint.Y);
  FHintWnd.ActivateHint(R, Node.Text)
end;

function TCustomTreeView.GetSelection: TTreeNode;
begin
  if RightClickSelect and Assigned(FRClickNode) then
    Result := FRClickNode
  else
    Result := FSelectedNode;
end;

function TCustomTreeView.GetSelectionCount: Cardinal;
begin
  Result := Items.SelectionCount;
end;

function TCustomTreeView.GetSelections(AIndex: Integer): TTreeNode;
begin
  if (AIndex >= 0) and (AIndex < Items.SelectionCount) then
    Result := Items.GetSelections(AIndex)
  else
    Result := nil;
end;

procedure TCustomTreeView.SetSelection(Value: TTreeNode);
var
  OldNode: TTreeNode;
begin
  if FSelectedNode = Value then Exit;
  if not CanChange(Value) then
    exit;
  {$IFDEF TREEVIEW_DEBUG}
  DebugLn('TCustomTreeView.SetSelection: Changing selection for Node: ', Text);
  {$ENDIF}
  EndEditing(true); // end editing before FSelectedNode change
  OldNode := FSelectedNode;
  FSelectedNode := Value;
  if Assigned(OldNode) then
    OldNode.Selected := False;
  if Assigned(Value) then
  begin
    Value.Selected := True;
    Value.MakeVisible;
  end;
  InternalSelectionChanged;
end;

function TCustomTreeView.GetShowButtons: boolean;
begin
  Result:=(tvoShowButtons in FOptions);
end;

function TCustomTreeView.GetShowLines: boolean;
begin
  Result:=(tvoShowLines in FOptions);
end;

function TCustomTreeView.GetShowRoot: boolean;
begin
  Result:=(tvoShowRoot in FOptions);
end;

function TCustomTreeView.GetShowSeparators: boolean;
begin
  Result:=(tvoShowSeparators in FOptions);
end;

function TCustomTreeView.GetToolTips: boolean;
begin
  Result:=(tvoToolTips in FOptions);
end;

procedure TCustomTreeView.SetExpandSignType(Value: TTreeViewExpandSignType);
begin
  if Value <> FExpandSignType then
  begin
    FExpandSignType := Value;
    Invalidate;
  end;
end;

procedure TCustomTreeView.SetDefaultItemHeight(Value: integer);
begin
  if (tvoAutoItemHeight in FOptions) and (not (csLoading in ComponentState))
  then exit;
  if Value<=0 then Value:=DefaultTreeNodeHeight;
  if Value=FDefItemHeight then exit;
  FDefItemHeight:=Value;
  Include(FStates,tvsTopsNeedsUpdate);
  Invalidate;
end;

function TCustomTreeView.GetAutoExpand: boolean;
begin
  Result:=(tvoAutoExpand in FOptions);
end;

function TCustomTreeView.GetBackgroundColor: TColor;
begin
  Result := Color;
end;

function TCustomTreeView.GetBottomItem: TTreeNode;
begin
  if HandleAllocated then begin
    UpdateBottomItem;
    Result := FBottomItem;
  end else
    Result := nil;
end;

function TCustomTreeView.GetDropTarget: TTreeNode;
begin
  if HandleAllocated then
  begin
    Result := FLastDropTarget;
  end
  else
    Result := nil;
end;

function TCustomTreeView.GetHideSelection: boolean;
begin
  Result:=(tvoHideSelection in FOptions);
end;

function TCustomTreeView.GetHotTrack: boolean;
begin
  Result:=(tvoHotTrack in FOptions);
end;

function TCustomTreeView.GetKeepCollapsedNodes: boolean;
begin
  Result:=(tvoKeepCollapsedNodes in FOptions);
end;

function TCustomTreeView.GetMultiSelect: Boolean;
begin
  Result := (tvoAllowMultiSelect in FOptions);
end;

function TCustomTreeView.GetReadOnly: boolean;
begin
  Result:=(tvoReadOnly in FOptions);
end;

function TCustomTreeView.GetExpandSignSize: integer;
begin
  if FExpandSignSize>=0 then
    Result := FExpandSignSize
  else
  if ExpandSignType = tvestTheme then
    Result := ScaleScreenToFont(FThemeExpandSignSize)
  else
    Result := Scale96ToFont(DefaultTreeNodeExpandSignSize);
end;

function TCustomTreeView.GetIndent: Integer;
begin
  if FIndent<0 then
    Result := Scale96ToFont(15)
  else
    Result := FIndent;
end;

function TCustomTreeView.GetRightClickSelect: boolean;
begin
  Result:=(tvoRightClickSelect in FOptions);
end;

function TCustomTreeView.GetRowSelect: boolean;
begin
  Result:=(tvoRowSelect in FOptions);
end;

procedure TCustomTreeView.SetDropTarget(Value: TTreeNode);
begin
  if HandleAllocated then
    if Value <> nil then
      Value.DropTarget := True;
end;

procedure TCustomTreeView.SetExpandSignSize(const AExpandSignSize: integer);
begin
  if FExpandSignSize = AExpandSignSize then Exit;
  FExpandSignSize := AExpandSignSize;
  Invalidate;
end;

function TCustomTreeView.IsEditing: Boolean;
begin
  Result:=tvsIsEditing in FStates;
end;

function TCustomTreeView.GetDragImages: TDragImageList;
begin
  if Assigned(FDragImage) and (FDragImage.Count > 0) then
    Result := FDragImage
  else
    Result := nil;
end;

procedure TCustomTreeView.UpdateInsertMark(X,Y: integer);
begin
  if (tvoAutoInsertMark in Options) and (not (csDesigning in ComponentState))
  then
    SetInsertMarkAt(X,Y)
  else
    SetInsertMark(nil,tvimNone);
end;

procedure TCustomTreeView.DoSelectionChanged;
var
  lAccessibleObject: TLazAccessibleObject;
  lSelection: TTreeNode;
  lSelectedText: string;
begin
  // Update the accessibility information
  lAccessibleObject := GetAccessibleObject();
  lSelection := Self.Selected;
  if lSelection = nil then lSelectedText := ''
  else lSelectedText := lSelection.Text;
  lAccessibleObject.AccessibleValue := lSelectedText;

  if Assigned(OnSelectionChanged) then OnSelectionChanged(Self);
end;

function TCustomTreeView.IsInsertMarkVisible: boolean;
begin
  Result:=(FInsertMarkType<>tvimNone) and (FInsertMarkNode<>nil)
           and (FInsertMarkNode.IsVisible);
end;

procedure TCustomTreeView.DoStartDrag(var DragObject: TDragObject);
var
  P: TPoint;
begin
  {$IFDEF VerboseDrag}
  DebugLn('TCustomTreeView.DoStartDrag A ',Name,':',ClassName);
  {$ENDIF}
  inherited DoStartDrag(DragObject);
  FLastDropTarget := nil;
  if FDragNode = nil then begin
    GetCursorPos(P);
    with ScreenToClient(P) do FDragNode := GetNodeAt(X, Y);
    {$IFDEF VerboseDrag}
    if FDragNode<>nil then
      DebugLn('TCustomTreeView.DoStartDrag DragNode=',FDragNode.Text)
    else
      DebugLn('TCustomTreeView.DoStartDrag DragNode=nil');
    {$ENDIF}
  end;
  FPrevToolTips := ToolTips;
  ToolTips := False;
  FDragScrollTimer.Enabled := true;
end;

procedure TCustomTreeView.DoEndDrag(Target: TObject; X, Y: Integer);
begin
  {$IFDEF VerboseDrag}
  DebugLn('TCustomTreeView.DoEndDrag A ',Name,':',ClassName);
  {$ENDIF}
  inherited DoEndDrag(Target, X, Y);
  FLastDropTarget := nil;
  FDragScrollTimer.Enabled := false;
  ToolTips := FPrevToolTips;
end;

function TCustomTreeView.DoMouseWheel(Shift: TShiftState; WheelDelta: Integer;
  MousePos: TPoint): Boolean;
var
  NDelta: integer;
begin
  Result:=inherited DoMouseWheel(Shift, WheelDelta, MousePos);
  if not Result then
  begin
    NDelta := (WheelDelta * Mouse.WheelScrollLines * DefaultItemHeight) div 120;
    ScrolledTop := ScrolledTop - NDelta;
    Result := true;
  end;
  UpdateTooltip(MousePos.X, MousePos.Y);
end;

function TCustomTreeView.DoMouseWheelHorz(Shift: TShiftState;
  WheelDelta: Integer; MousePos: TPoint): Boolean;
var
  NDelta: integer;
const
  cScrollStep = 50;
begin
  Result:=inherited DoMouseWheelHorz(Shift, WheelDelta, MousePos);
  if not Result then
  begin
    NDelta := (WheelDelta * cScrollStep) div 120;
    ScrolledLeft := ScrolledLeft + NDelta;
    Result := true;
  end;
  UpdateTooltip(MousePos.X, MousePos.Y);
end;

function TCustomTreeView.DoDragMsg(ADragMessage: TDragMessage; APosition: TPoint; ADragObject: TDragObject; ATarget: TControl; ADocking: Boolean):LRESULT;
begin
  Result:=inherited;
  {$IFDEF VerboseDrag}
  DebugLn('TCustomTreeView.DoDragMsg ',Name,':',ClassName,' ',IntToStr(ord(ADragMessage)));
  {$ENDIF}
  case ADragMessage of
    {dmDragMove:
      begin
        P:=ScreenToClient(Pos);
        DoDragOver(Source, P.X, P.Y, AMessage.Result <> 0);
      end;}
    dmDragLeave:
      begin
        ADragObject.HideDragImage;
        FLastDropTarget := DropTarget;
        DropTarget := nil;
        ADragObject.ShowDragImage;
      end;
    dmDragDrop: FLastDropTarget := nil;
  end;
end;

procedure TCustomTreeView.DragOver(Source: TObject; X,Y: Integer;
  State: TDragState; var Accept: Boolean);
var
  Node: TTreeNode;
begin
  inherited DragOver(Source,X,Y,State,Accept);
  Node := GetNodeAt(X, Y);
  {$IFDEF VerboseDrag}
  DebugLn(['TCustomTreeView.DragOver ',Name,':',ClassName,' ',Node<>nil,' ',Node <> DropTarget,' ',Node = FLastDropTarget]);
  DebugLn(['TCustomTreeView.DragOver Source ',Source,':',Source.ClassName]);
  {$ENDIF}
  if (Node <> nil)
  and ((Node <> DropTarget) or (Node = FLastDropTarget)) then
  begin
    FLastDropTarget := nil;
    Node.DropTarget := True;
  end;
end;

procedure TCustomTreeView.DoPaint;
var
  a,HalfBorderWidth:integer;
  SpaceRect, DrawRect: TRect;
  Node: TTreeNode;
  InsertMarkRect: TRect;
begin
  if [tvsUpdating,tvsPainting] * FStates <> [] then Exit;
  Include(FStates, tvsPainting);
  try
    if Focused then
      Include(FStates,tvoFocusedPainting)
    else
      Exclude(FStates,tvoFocusedPainting);
    if (tvoAutoItemHeight in fOptions) then
      UpdateDefaultItemHeight;
    UpdateScrollbars;
    with Canvas do
    begin
      if IsCustomDrawn(dtControl, cdPrePaint) then
      begin
        DrawRect := ClientRect;
        if not CustomDraw(DrawRect, cdPrePaint) then exit;
      end;
      // draw nodes
      Node := TopItem;
      //write('[TCustomTreeView.DoPaint] A Node=',DbgS(Node));
      //if Node<>nil then DebugLn(' Node.Text=',Node.Text) else DebugLn('');
      while Node <> nil do
      begin
        if Node.Visible then
          DoPaintNode(Node);
        Node := Node.GetNextVisible;
        //write('[TCustomTreeView.DoPaint] B Node=',DbgS(Node));
        //if Node<>nil then DebugLn(' Node.Text=',Node.Text) else DebugLn('');
      end;
      SpaceRect := Rect(BorderWidth, BorderWidth,
                        ClientWidth - BorderWidth,
                        ClientHeight - BorderWidth);
      // draw insert mark for new root node
      if (InsertMarkType = tvimAsFirstChild) and (Items.Count = 0) then
      begin
        Pen.Color := FTreeLineColor;
        Brush.Color := FSelectedColor;
        InsertMarkRect := SpaceRect;
        InsertMarkRect.Bottom := InsertMarkRect.Top + 2;
        Rectangle(InsertMarkRect);
        SpaceRect.Top := InsertMarkRect.Bottom;
      end;
      // draw unused space below nodes
      Node := BottomItem;
      if Node <> nil then
        SpaceRect.Top := Node.Top + Node.Height - FScrolledTop + BorderWidth;
      //if Node<>nil then DebugLn('BottomItem=',BottomItem.text) else DebugLn('NO BOTTOMITEM!!!!!!!!!');
      // TWinControl(Parent).InvalidateRect(Self,SpaceRect,true);
      if (Color <> clNone) and (SpaceRect.Top < SpaceRect.Bottom) then
      begin
        //DebugLn('  SpaceRect=',SpaceRect.Left,',',SpaceRect.Top,',',SpaceRect.Right,',',SpaceRect.Bottom);
        Brush.Color := Color;
        FillRect(SpaceRect);
      end;
      // draw border
      HalfBorderWidth := BorderWidth shr 1;
      Pen.Color := clGray;
      for a := 0 to BorderWidth - 1 do
      begin
        if a = HalfBorderWidth then
          Pen.Color := clBlack;
        MoveTo(a, ClientHeight - 1 - a);
        LineTo(a, a);
        LineTo(ClientWidth - 1 - a, a);
      end;
      Pen.Color := clWhite;
      for a := 0 to BorderWidth - 1 do
      begin
        if a = HalfBorderWidth then
          Pen.Color := clLtGray;
        MoveTo(ClientWidth - 1 - a, a);
        LineTo(ClientWidth - 1 - a, ClientHeight - 1 - a);
        LineTo(a, ClientHeight - 1 - a);
      end;
      if IsCustomDrawn(dtControl, cdPostPaint) then
      begin
        DrawRect := ClientRect;
        if not CustomDraw(DrawRect, cdPostPaint) then exit;
      end;
    end;
  finally
    Exclude(FStates, tvsPainting);
  end;
end;

procedure TCustomTreeView.DoPaintNode(Node: TTreeNode);
var
  NodeRect: TRect;
  VertMid, VertDelta, RealExpandSignSize, RealIndent: integer;
  NodeSelected, HasExpandSign: boolean;

  function InvertColor(AColor: TColor): TColor;
  var Red, Green, Blue: integer;
  begin
    if AColor<>clHighlight then begin
      Result:=clWhite;
      Red:=(AColor shr 16) and $ff;
      Green:=(AColor shr 8) and $ff;
      Blue:=AColor and $ff;
      if Red+Green+Blue>$180 then
        Result:=clBlack;
      //DebugLn(['[TCustomTreeView.DoPaintNode.InvertColor] Result=',Result,' ',Red,',',Green,',',Blue]);
    end
    else
      Result := clHighlightText;
  end;

  procedure DrawVertLine(X, Y1, Y2: Integer);
  begin
    if Y1 > Y2 then
      Exit;
    if TreeLinePenStyle = psPattern then
    begin
      // TODO: implement psPattern support in the LCL
      Y1 := Y1 + VertDelta;
      while Y1 < Y2 do
      begin
        Canvas.Pixels[X, Y1] := TreeLineColor;
        inc(Y1, 2);
      end;
    end
    else
    begin
      Canvas.MoveTo(X, Y1);
      Canvas.LineTo(X, Y2);
    end;
  end;

  procedure DrawHorzLine(Y, X1, X2: Integer);
  begin
    if X1 > X2 then
      Exit;
    if TreeLinePenStyle = psPattern then
    begin
      // TODO: implement psPattern support in the LCL
      while X1 < X2 do
      begin
        Canvas.Pixels[X1, Y] := TreeLineColor;
        inc(X1, 2);
      end;
    end
    else
    begin
      Canvas.MoveTo(X1, Y);
      Canvas.LineTo(X2, Y);
    end;
  end;

  function DrawTreeLines(CurNode: TTreeNode): integer;
  // paints tree lines, returns indent
  var
    CurMid: integer;
  begin
    if (CurNode <> nil) and ((tvoShowRoot in Options) or (CurNode.Parent<>nil)) then
    begin
      Result := DrawTreeLines(CurNode.Parent);
      if ShowLines then
      begin
        CurMid := Result + (RealIndent shr 1);
        if CurNode = Node then
        begin
          // draw horizontal line
          if HasExpandSign then
            DrawHorzLine(VertMid, CurMid + RealExpandSignSize div 2, Result + RealIndent)
          else
            DrawHorzLine(VertMid, CurMid, Result + RealIndent);
        end;

        if (CurNode.GetNextVisibleSibling <> nil) then
        begin
          // draw vertical line to next brother
          if (CurNode = Node) and HasExpandSign then
          begin
            if (Node.Parent = nil) and (Node.GetPrevSibling = nil) then
              DrawVertLine(CurMid, VertMid + RealExpandSignSize div 2, NodeRect.Bottom)
            else
            begin
              DrawVertLine(CurMid, NodeRect.Top, VertMid);
              DrawVertLine(CurMid, VertMid + RealExpandSignSize div 2 + VertDelta, NodeRect.Bottom);
            end;
          end
          else
          if (Node.Parent = nil) and (Node.GetPrevSibling = nil) then
            DrawVertLine(CurMid, VertMid + VertDelta, NodeRect.Bottom)
          else
            DrawVertLine(CurMid, NodeRect.Top, NodeRect.Bottom);
        end else
        if (CurNode = Node) then
        begin
          // draw vertical line from top to horizontal line
          if HasExpandSign then
          begin
            if ((InsertMarkNode = Node) and (InsertMarkType = tvimAsNextSibling)) then
            begin
              DrawVertLine(CurMid, NodeRect.Top, VertMid);
              DrawVertLine(CurMid, VertMid + RealExpandSignSize div 2, NodeRect.Bottom - 1);
            end
            else
              DrawVertLine(CurMid, NodeRect.Top, VertMid);
          end
          else
          if ((InsertMarkNode = Node) and (InsertMarkType = tvimAsNextSibling)) then
            DrawVertLine(CurMid, NodeRect.Top, NodeRect.Bottom - 1)
          else
            DrawVertLine(CurMid, NodeRect.Top, VertMid);
        end;
      end;
      inc(Result, RealIndent);
    end else
    begin
      Result := BorderWidth - FScrolledLeft;
      if CurNode <> nil then // indent first level of tree with ShowRoot = false a bit
        inc(Result, RealIndent shr 2);
    end;
  end;

  procedure DrawExpandSign(MidX, MidY: integer; CollapseSign: boolean);
  const
    PlusMinusDetail: array[Boolean {Hot}, Boolean {Expanded}] of TThemedTreeview =
    (
      (ttGlyphClosed, ttGlyphOpened),
      (ttHotGlyphClosed, ttHotGlyphOpened)
    );
  var
    HalfSize, ALeft, ATop, ARight, ABottom, SmallIndent: integer;
    Points: array[0..2] of TPoint; // for triangle
    Details: TThemedElementDetails;
    R: TRect;
    PrevColor: TColor;
  const
    cShiftHorzArrow = 2; //paint horz arrow N pixels upper than MidY
  begin
    HalfSize := RealExpandSignSize div 2;
    //if not Odd(RealExpandSignSize) then
    //  Dec(HalfSize);
    ALeft := MidX - HalfSize;
    ARight := MidX + HalfSize;
    ATop := MidY - HalfSize;
    ABottom := MidY + HalfSize;

    if Assigned(FOnCustomDrawArrow) then
    begin
      FOnCustomDrawArrow(Self, Rect(ALeft, ATop, ARight, ABottom), not CollapseSign);
      Exit
    end;

    with Canvas do
    begin
      Pen.Color := FExpandSignColor;
      Pen.Style := psSolid;
      case ExpandSignType of
      tvestTheme:
        begin
          // draw a themed expand sign. Todo: track hot
          R := Rect(ALeft, ATop, ARight, ABottom);
          Details := ThemeServices.GetElementDetails(PlusMinusDetail[False, CollapseSign]);
          ThemeServices.DrawElement(Canvas.Handle, Details, R, nil);
        end;
      tvestPlusMinus:
        begin
          // draw a plus or a minus sign
          R := Rect(ALeft, ATop, ARight+1, ABottom+1); //+1 for centering of line in square
          Rectangle(R);
          SmallIndent := Scale96ToFont(2);
          MoveTo(R.Left + SmallIndent, MidY);
          LineTo(R.Right - SmallIndent, MidY);
          if not CollapseSign then
          begin
            MoveTo(MidX, R.Top + SmallIndent);
            LineTo(MidX, R.Bottom - SmallIndent);
          end;
        end;
      tvestArrow,
      tvestArrowFill:
        begin
          // draw an arrow. down for collapse and right for expand
          R := Rect(ALeft, ATop, ARight+1, ABottom+1); //+1 for simmetry of arrow
          if CollapseSign then
          begin
            // draw an arrow down
            Points[0] := Point(R.Left, MidY - cShiftHorzArrow);
            Points[1] := Point(R.Right - 1, MidY - cShiftHorzArrow);
            Points[2] := Point(MidX, R.Bottom - 1 - cShiftHorzArrow);
          end else
          begin
            // draw an arrow right
            Points[0] := Point(MidX - 1, ATop);
            Points[1] := Point(R.Right - 2, MidY);
            Points[2] := Point(MidX - 1, R.Bottom - 1);
          end;

          if ExpandSignType = tvestArrowFill then
          begin
            PrevColor := Brush.Color;
            Brush.Color := ExpandSignColor;
          end;
          Polygon(Points, 3, False);
          if ExpandSignType = tvestArrowFill then
          begin
            Brush.Color := PrevColor;
          end;
        end;
      end;
    end;
  end;

  procedure DrawInsertMark;
  var
    InsertMarkRect: TRect;
    x: Integer;
  begin
    case InsertMarkType of

    tvimAsFirstChild:
      if InsertMarkNode=Node then begin
        // draw insert mark for new first child
        with Canvas do begin
          // draw virtual tree line
          Pen.Color:=TreeLineColor;
          // Pen.Style:=TreeLinePenStyle; ToDo: not yet implemented in all widgetsets
          x:=Node.DisplayExpandSignRight+RealIndent div 2;
          MoveTo(x,NodeRect.Bottom-3);
          LineTo(x,NodeRect.Bottom-2);
          x:=Node.DisplayExpandSignRight+RealIndent;
          LineTo(x,NodeRect.Bottom-2);
          Pen.Style:=psSolid;

          // draw virtual rectangle
          Pen.Color:=TreeLineColor;
          Brush.Color:=FSelectedColor;
          InsertMarkRect:=Rect(x,NodeRect.Bottom-3,
                               NodeRect.Right,NodeRect.Bottom-1);
          Rectangle(InsertMarkRect);
        end;
      end;

    tvimAsPrevSibling:
      if InsertMarkNode=Node then begin
        // draw insert mark for new previous sibling
        with Canvas do begin
          // draw virtual tree line
          Pen.Color:=TreeLineColor;
          //Pen.Style:=TreeLinePenStyle; ToDo: not yet implemented in all widgetsets
          x:=Node.DisplayExpandSignLeft+RealIndent div 2;
          MoveTo(x,NodeRect.Top+1);
          x:=Node.DisplayExpandSignRight;
          LineTo(x,NodeRect.Top+1);
          Pen.Style:=psSolid;

          // draw virtual rectangle
          Pen.Color:=TreeLineColor;
          Brush.Color:=FSelectedColor;
          InsertMarkRect:=Rect(x,NodeRect.Top,
                               NodeRect.Right,NodeRect.Top+2);
          Rectangle(InsertMarkRect);
        end;
      end;

    tvimAsNextSibling:
      if InsertMarkNode=Node then begin
        // draw insert mark for new next sibling
        with Canvas do begin
          // draw virtual tree line
          Pen.Color:=TreeLineColor;
          //Pen.Style:=TreeLinePenStyle; ToDo: not yet implemented in all widgetsets
          x:=Node.DisplayExpandSignLeft+RealIndent div 2;
          MoveTo(x,NodeRect.Bottom-3);
          LineTo(x,NodeRect.Bottom-2);
          x:=Node.DisplayExpandSignRight;
          LineTo(x,NodeRect.Bottom-2);
          Pen.Style:=psSolid;

          // draw virtual rectangle
          Pen.Color:=TreeLineColor;
          Brush.Color:=FSelectedColor;
          InsertMarkRect:=Rect(x,NodeRect.Bottom-3,
                               NodeRect.Right,NodeRect.Bottom-1);
          Rectangle(InsertMarkRect);
        end;
      end;

    end;
  end;

  procedure DrawBackground(IsSelected: Boolean; ARect: TRect);
  var
    Details: TThemedElementDetails;
    CurBackgroundColor,bclr: TColor;
  begin
    bclr:=Canvas.Brush.Color;
    try
      if (tvoRowSelect in Options) and IsSelected then
        if tvoThemedDraw in Options then
        begin
          if tvoFocusedPainting in FStates then
            Details := ThemeServices.GetElementDetails(ttItemSelected)
          else
            Details := ThemeServices.GetElementDetails(ttItemSelectedNotFocus);
          if ThemeServices.HasTransparentParts(Details) then
          begin
            Canvas.Brush.Color := Color;
            Canvas.FillRect(ARect);
          end;
          ThemeServices.DrawElement(Canvas.Handle, Details, ARect, nil);
          Exit;
        end
        else
          CurBackgroundColor := FSelectedColor
      else
        CurBackgroundColor := Color;
      if CurBackgroundColor <> clNone then
      begin
        Canvas.Brush.Color := CurBackgroundColor;
        Canvas.FillRect(ARect);
      end;
    finally
      Canvas.Brush.Color := bclr;
    end;
  end;

  procedure DrawNodeText(IsSelected: Boolean; NodeRect: TRect; AText: String);
  var
    Details: TThemedElementDetails;
    NeedUnderline: Boolean;
    PrevFontStyle: TFontStyles;
    PrevFontColor: TColor;
  begin
    if IsSelected then
    begin
      if tvoFocusedPainting in FStates then
        Details := ThemeServices.GetElementDetails(ttItemSelected)
      else
        Details := ThemeServices.GetElementDetails(ttItemSelectedNotFocus);
      if not (tvoRowSelect in Options) then
        if (tvoThemedDraw in Options) then
          ThemeServices.DrawElement(Canvas.Handle, Details, NodeRect, nil)
        else
        begin
          Canvas.Brush.Color := FSelectedColor;
          Canvas.Font.Color := IfThen(FSelectedFontColorUsed, FSelectedFontColor, InvertColor(FSelectedColor));
          Canvas.FillRect(NodeRect);
        end
      else
      if not (tvoThemedDraw in Options) then
      begin
        Canvas.Brush.Color := FSelectedColor;
        Canvas.Font.Color := IfThen(FSelectedFontColorUsed, FSelectedFontColor, InvertColor(FSelectedColor));
        Canvas.FillRect(NodeRect);
      end;
    end
    else
      Details := ThemeServices.GetElementDetails(ttItemNormal);

    NeedUnderline := (tvoHotTrack in FOptions) and (Node=FNodeUnderCursor);
    if NeedUnderline then
    begin
      PrevFontStyle := Canvas.Font.Style;
      PrevFontColor := Canvas.Font.Color;
      Canvas.Font.Style := [fsUnderline];
      if FHotTrackColor<>clNone then
        Canvas.Font.Color := FHotTrackColor;
    end;

    if (tvoThemedDraw in Options) then
      ThemeServices.DrawText(Canvas, Details, AText, NodeRect, DT_CENTER or DT_VCENTER or DT_SINGLELINE or DT_NOPREFIX, 0)
    else
      DrawText(Canvas.Handle, PChar(AText), -1, NodeRect, DT_CENTER or DT_VCENTER or DT_SINGLELINE or DT_NOPREFIX);

    if NeedUnderline then
    begin
      Canvas.Font.Style := PrevFontStyle;
      Canvas.Font.Color := PrevFontColor;
    end;
  end;


var
  x, ImgIndex: integer;
  CurTextRect: TRect;
  DrawState: TCustomDrawState;
  PaintImages: boolean;
  OverlayIndex: Integer;
  ImageRes, StateImageRes: TScaledImageListResolution;
begin
  if Assigned(FImages) then
    ImageRes := Images.ResolutionForPPI[ImagesWidth, Font.PixelsPerInch, GetCanvasScaleFactor];
  if Assigned(FStateImages) then
    StateImageRes := StateImages.ResolutionForPPI[StateImagesWidth, Font.PixelsPerInch, GetCanvasScaleFactor];
  RealExpandSignSize := ExpandSignSize;
  RealIndent := Indent;
  NodeRect := Node.DisplayRect(False);
  if (NodeRect.Bottom < 0) or (NodeRect.Top >= ClientHeight) then
    Exit;
  NodeSelected := (Node.Selected) or (Node.MultiSelected);
  Canvas.Font.Color := Font.Color;
  Canvas.Brush.Color := Color;
  PaintImages := True;
  if IsCustomDrawn(dtItem, cdPrePaint) then
  begin
    DrawState := [];
    if NodeSelected then
      Include(DrawState, cdsSelected);
    if Node.Focused then
      Include(DrawState, cdsFocused);
    if Node.MultiSelected then
      Include(DrawState, cdsMarked);
    if not CustomDrawItem(Node, DrawState, cdPrePaint, PaintImages) then Exit;
  end;

  VertMid := NodeRect.Top + (NodeRect.Bottom - NodeRect.Top) div 2;
  HasExpandSign := ShowButtons and Node.HasChildren and ((tvoShowRoot in Options) or (Node.Parent <> nil));
  VertDelta := Ord(FDefItemHeight and 3 = 2);
  //DebugLn(['[TCustomTreeView.DoPaintNode] Node=',DbgS(Node),' Node.Text=',Node.Text,' NodeRect=',NodeRect.Left,',',NodeRect.Top,',',NodeRect.Right,',',NodeRect.Bottom,' VertMid=',VertMid]);
  with Canvas do
  begin
    // draw background
    DrawBackground(NodeSelected, NodeRect);

    // draw tree lines
    Pen.Color := TreeLineColor;
    Pen.Style := TreeLinePenStyle;
    if Pen.Style = psPattern then
      Pen.SetPattern(FTreeLinePenPattern);
    x := DrawTreeLines(Node);
    Pen.Style := psSolid;

    // draw expand sign
    if HasExpandSign then
      DrawExpandSign(x - RealIndent + (RealIndent shr 1), VertMid, Node.Expanded);

    // draw state icon
    if (StateImages <> nil) then
    begin
      if (Node.StateIndex >= 0) and (Node.StateIndex < StateImages.Count) then
      begin
        if PaintImages then
          StateImageRes.Draw(Canvas, x + 1, NodeRect.Top +(NodeRect.Bottom - NodeRect.Top - StateImageRes.Height) div 2,
            Node.StateIndex, True);
        Inc(x, StateImageRes.Width + FDefItemSpace);
      end;
    end;

    // draw icon
    if (Images <> nil) then
    begin
      if FSelectedNode <> Node then
      begin
      	GetImageIndex(Node);
        ImgIndex := Node.ImageIndex
      end
      else
      begin
      	GetSelectedIndex(Node);
        ImgIndex := Node.SelectedIndex;
      end;
      if (ImgIndex >= 0) and (ImgIndex < Images.Count) then
      begin
        if PaintImages then
        begin
      	  if (Node.OverlayIndex >= 0) then begin
            OverlayIndex:=Node.OverlayIndex;
            if Images.HasOverlays then begin
              ImageRes.DrawOverlay(Canvas, x + 1, NodeRect.Top + (NodeRect.Bottom - NodeRect.Top - ImageRes.Height) div 2,
                 ImgIndex, OverlayIndex, Node.FNodeEffect);
            end else begin
              // draw the Overlay using the image from the list
              // set an Overlay
              Images.OverLay(OverlayIndex,0);
              // draw overlay
              ImageRes.DrawOverlay(Canvas, x + 1, NodeRect.Top + (NodeRect.Bottom - NodeRect.Top - ImageRes.Height) div 2,
                 ImgIndex, 0, Node.FNodeEffect);
              // reset the Overlay
              Images.OverLay(-1,0);
            end;
          end
          else begin
            ImageRes.Draw(Canvas, x + 1, NodeRect.Top + (NodeRect.Bottom - NodeRect.Top - ImageRes.Height) div 2,
               ImgIndex, Node.FNodeEffect);
          end;
        end;
        Inc(x, ImageRes.Width + FDefItemSpace);
      end;
    end;

    // draw text
    if Node.Text <> '' then
    begin
      CurTextRect := NodeRect;
      CurTextRect.Left := x;
      CurTextRect.Right := x + TextWidth(Node.Text) + RealIndent div 2;
      DrawNodeText(NodeSelected, CurTextRect, Node.Text);
    end;

    // draw separator
    if (tvoShowSeparators in FOptions) then
    begin
      Pen.Color:=SeparatorColor;
      MoveTo(NodeRect.Left,NodeRect.Bottom-1);
      LineTo(NodeRect.Right,NodeRect.Bottom-1);
    end;

    // draw insert mark
    DrawInsertMark;
  end;
  PaintImages := true;
  if IsCustomDrawn(dtItem, cdPostPaint) then
  begin
    DrawState:=[];
    if Node.Selected then
      Include(DrawState,cdsSelected);
    if Node.Focused then
      Include(DrawState,cdsFocused);
    if Node.MultiSelected then
      Include(DrawState,cdsMarked);
    if not CustomDrawItem(Node,DrawState,cdPostPaint,PaintImages) then exit;
  end;
end;

procedure TCustomTreeView.GetImageIndex(Node: TTreeNode);
begin
  if Assigned(FOnGetImageIndex) then FOnGetImageIndex(Self, Node);
end;

procedure TCustomTreeView.GetSelectedIndex(Node: TTreeNode);
begin
  if Assigned(FOnGetSelectedIndex) then FOnGetSelectedIndex(Self, Node);
end;

function TCustomTreeView.CanChange(Node: TTreeNode): Boolean;
begin
  Result := True;
  if Assigned(Node) and Assigned(FOnChanging) then FOnChanging(Self, Node, Result);
end;

procedure TCustomTreeView.Change(Node: TTreeNode);
begin
  if Assigned(FOnChange) then
    FOnChange(Self, Node);
end;

procedure TCustomTreeView.Delete(Node: TTreeNode);
begin
  if Assigned(FOnDeletion) then FOnDeletion(Self, Node);
end;

procedure TCustomTreeView.Expand(Node: TTreeNode);
begin
  if Assigned(FOnExpanded) then FOnExpanded(Self, Node);
end;

function TCustomTreeView.ExpandSignSizeIsStored: Boolean;
begin
  Result := FExpandSignSize >= 0;
end;

function TCustomTreeView.CanExpand(Node: TTreeNode): Boolean;
begin
  Result := True;
  if Assigned(FOnExpanding) then FOnExpanding(Self, Node, Result);
end;

procedure TCustomTreeView.Collapse(Node: TTreeNode);
begin
  if Assigned(FOnCollapsed) then FOnCollapsed(Self, Node);
end;

function TCustomTreeView.CanCollapse(Node: TTreeNode): Boolean;
begin
  Result := True;
  if Assigned(FOnCollapsing) then FOnCollapsing(Self, Node, Result);
end;

function TCustomTreeView.CanEdit(Node: TTreeNode): Boolean;
begin
  Result := True;
  if Assigned(FOnEditing) then FOnEditing(Self, Node, Result);
end;

procedure TCustomTreeView.EndEditing(Cancel: boolean);
var
  NewText: String;
  Node: TTreeNode;
begin
  //DebugLn(['TCustomTreeView.EndEditing ',DbgSName(Self),' ',tvsIsEditing in FStates,' ',DbgSName(FEditor)]);
  if not (tvsIsEditing in FStates) then exit;
  Exclude(FStates,tvsIsEditing);
  if FEditor<>nil then begin
    // get new value fom edit control and hide it
    NewText:='';
    if not Cancel then
      NewText:=FEditor.Text;
    FEditor.Parent:=nil;

    // commit new value
    if not Cancel then begin
      Node:=FEditingItem;
      if (Node<>nil) then begin
        if Assigned(OnEdited) then
          OnEdited(Self,Node,NewText);
        Node.Text:=NewText;
      end;
    end;
    if Assigned(FOnEditingEnd) then FOnEditingEnd(Self, FEditingItem, Cancel);
  end;
  FEditingItem := nil;
  Invalidate;
end;

procedure TCustomTreeView.EnsureNodeIsVisible(ANode: TTreeNode);
var b: integer;
begin
  if ANode=nil then exit;
  ANode.ExpandParents;
  if ANode.Top<ScrolledTop then
    ScrolledTop:=ANode.Top
  else begin
    b:=ANode.Top+ANode.Height-GetNodeDrawAreaHeight;
    if ScrolledTop<b then ScrolledTop:=b;
  end;
end;

function TCustomTreeView.CreateNode: TTreeNode;
var
  NewNodeClass: TTreeNodeClass;
begin
  Result := nil;
  if Assigned(FOnCustomCreateItem) then
    FOnCustomCreateItem(Self, Result);
  if Result = nil then
  begin
    NewNodeClass:=TTreeNode;
    DoCreateNodeClass(NewNodeClass);
    Result := NewNodeClass.Create(Items);
  end;
end;

function TCustomTreeView.CreateNodes: TTreeNodes;
begin
  Result := TTreeNodes.Create(Self);
end;

procedure TCustomTreeView.ImageListChange(Sender: TObject);
begin
  Invalidate;
end;

function TCustomTreeView.IndentIsStored: Boolean;
begin
  Result := FIndent >= 0;
end;

procedure TCustomTreeView.MouseDown(Button: TMouseButton; Shift: TShiftState;
  X, Y: Integer);
var
  CursorNode: TTreeNode;
  LogicalX: Integer;
  CursorNodeSelected: Boolean;
begin
  {$IFDEF VerboseDrag}
  DebugLn('TCustomTreeView.MouseDown A ',DbgSName(Self),' ');
  {$ENDIF}
  fMouseDownPos := Point(X,Y);
  FStates:=FStates-[tvsEditOnMouseUp,tvsSingleSelectOnMouseUp];

  CursorNode := GetNodeAt(X, Y);
  CursorNodeSelected := (CursorNode<>nil)
    and (CursorNode.Selected
      or ((tvoAllowMultiselect in Options) and CursorNode.MultiSelected));
  LogicalX:=X;

  //change selection on right click
  if (Button = mbRight) and RightClickSelect and//right click
     (([ssDouble, ssTriple, ssQuad] * Shift) = []) and//single or first of a multi click
     not AllowMultiSelectWithCtrl(Shift) and//only when CTRL is not pressed
     (CursorNode <> nil) and
     (LogicalX >= CursorNode.DisplayStateIconLeft)//only after expand sign
  then
  begin
    if not (tvoAllowMultiselect in Options) then
      Selected := CursorNode
    else
    if not CursorNodeSelected then
      Items.SelectOnlyThis(CursorNode);
  end;

  if not Focused and CanFocus then
    SetFocus;

  inherited MouseDown(Button, Shift, X, Y);

  //CursorNode must be reassigned again - e.g. in OnMouseDown the node can be deleted or moved.
  CursorNode := GetNodeAt(X, Y);
  CursorNodeSelected := (CursorNode<>nil)
    and (CursorNode.Selected
      or ((tvoAllowMultiselect in Options) and CursorNode.MultiSelected));

  //change selection on left click
  if (Button = mbLeft) and//left click
     (([ssDouble, ssTriple, ssQuad] * Shift) = []) and//single or first of a multi click
     (CursorNode <> nil) then
  begin
    if CursorNode.HasChildren and ShowButtons and
       (LogicalX >= CursorNode.DisplayExpandSignLeft) and
       (LogicalX < CursorNode.DisplayExpandSignRight) then
    begin
      // mousedown occured on expand sign -> expand/collapse
      CursorNode.Expanded := not CursorNode.Expanded;
    end
    else if LogicalX >= CursorNode.DisplayStateIconLeft then
    begin
      // mousedown occured in text or icon
      // -> select node and begin drag operation
      {$IFDEF VerboseDrag}
      DebugLn(['TCustomTreeView.MouseDown In Text ',DbgSName(Self),' MouseCapture=',MouseCapture]);
      {$ENDIF}
      if (Selected = CursorNode) and (LogicalX >= CursorNode.DisplayTextLeft) then
        Include(FStates, tvsEditOnMouseUp);
      if not (tvoAllowMultiselect in Options) then
        Selected := CursorNode
      else
      begin
        if AllowMultiSelectWithShift(Shift) then
        begin
          Exclude(FStates,tvsEditOnMouseUp);
          LockSelectionChangeEvent;
          try
            Items.MultiSelect(CursorNode, not AllowMultiSelectWithCtrl(Shift));
          finally
            UnlockSelectionChangeEvent;
          end;
        end
        else if AllowMultiSelectWithCtrl(Shift) then
        begin
          Exclude(FStates,tvsEditOnMouseUp);
          CursorNode.MultiSelected:=not CursorNode.MultiSelected;
          if CursorNode.MultiSelected then
            FTreeNodes.FStartMultiSelected := CursorNode;
        end
        else
        begin
          if not CursorNodeSelected then
            Items.SelectOnlyThis(CursorNode)
          else
            Include(FStates, tvsSingleSelectOnMouseUp);
        end;
      end;
    end;
  end
  else// multi click
  if not (tvoNoDoubleClickExpand in Options) and (ssDouble in Shift)
  and (Button = mbLeft) and (CursorNode<>nil) then
    CursorNode.Expanded := not CursorNode.Expanded;
end;

procedure TCustomTreeView.MouseMove(Shift: TShiftState; X, Y: Integer);
begin
  inherited MouseMove(Shift, x, y);
  if (tvoAutoInsertMark in FOptions) then
    UpdateInsertMark(X,Y);
  UpdateTooltip(X, Y);
  UpdateHotTrack(X, Y);
end;

procedure TCustomTreeView.UpdateHotTrack(X, Y: Integer);
begin
  FNodeUnderCursor := nil;
  if Cursor = crHandPoint then
    Cursor := crDefault;
  if not (tvoHotTrack in FOptions) then Exit;

  FNodeUnderCursor := GetNodeAt(X, Y);
  if Assigned(FNodeUnderCursor) then
    Cursor := crHandPoint;
  Invalidate;
end;

procedure TCustomTreeView.MouseUp(Button: TMouseButton; Shift: TShiftState;
  X, Y: Integer);
var
  aMouseDownNode, aMouseUpNode: TTreeNode;
begin
  if (FHintWnd<>nil) and FHintWnd.Visible then // must hide hint window in mouse up to receive redirected mouse up messages
    FHintWnd.Hide;
  inherited MouseUp(Button, Shift, X, Y);
  if (Button = mbRight) and (Shift = [ssRight]) and Assigned(PopupMenu) then
    exit;
  if Button=mbLeft then
    MouseCapture := False;
  if (Button=mbLeft)
  and (FStates * [tvsDblClicked, tvsTripleClicked, tvsQuadClicked] = [])
  then begin
    //AquirePrimarySelection;
    aMouseDownNode:=GetNodeAt(fMouseDownPos.X,fMouseDownPos.Y);
    aMouseUpNode:=GetNodeAt(X,Y);
    if (abs(fMouseDownPos.X-X)+abs(fMouseDownPos.Y-Y)<10)
    and (aMouseDownNode=aMouseUpNode) then
    begin
      // mouse up on mouse-down node
      if (tvsEditOnMouseUp in FStates) and (not ReadOnly) then
        BeginEditing(Selected)
      else if (tvsSingleSelectOnMouseUp in FStates) then
        Items.SelectOnlyThis(aMouseUpNode);
    end;
  end;
  FStates:=FStates-[tvsDblClicked,tvsTripleClicked,tvsQuadClicked,
                    tvsEditOnMouseUp,tvsSingleSelectOnMouseUp];
end;

procedure TCustomTreeView.MoveEnd(ASelect: Boolean);
var
  lNode: TTreeNode;
begin
  lNode := Items.GetLastExpandedSubNode;
  if lNode <> nil then
    MoveSelection(lNode, ASelect);
end;

procedure TCustomTreeView.MoveHome(ASelect: Boolean);
var
  lNode: TTreeNode;
begin
  lNode := Items.GetFirstVisibleNode;
  if lNode <> nil then
    MoveSelection(lNode, ASelect);
end;

procedure TCustomTreeView.MovePageDown(ASelect: Boolean);
var
  I: Integer;
  lNode: TTreeNode;
begin
  if tvoAllowMultiSelect in FOptions then
    lNode := FTreeNodes.FLastMultiSelected
  else
    lNode := Selected;
  if lNode = nil then
    lNode := Items.GetFirstVisibleNode;

  if lNode <> nil then
  begin
    I := Pred(ClientHeight div DefaultItemHeight);

    while (I > 0) do
      if lNode.GetNextExpanded <> nil then
      begin
        lNode := lNode.GetNextExpanded;
        Dec(I);
      end
      else Break;

  end;
  if lNode <> nil then
    MoveSelection(lNode, ASelect);
end;

procedure TCustomTreeView.MovePageUp(ASelect: Boolean);
var
  I: Integer;
  lNode: TTreeNode;
begin
  if tvoAllowMultiSelect in FOptions then
    lNode := FTreeNodes.FLastMultiSelected
  else
    lNode := Selected;
  if lNode = nil then
    lNode := Items.GetFirstVisibleNode;

  if lNode <> nil then
  begin
    I := Pred(ClientHeight div DefaultItemHeight);

    while (I > 0) do
      if lNode.GetPrevExpanded <> nil then
      begin
        lNode := lNode.GetPrevExpanded;
        Dec(I);
      end
      else Break;

  end;
  if lNode <> nil then
    MoveSelection(lNode, ASelect);
end;

procedure TCustomTreeView.MoveSelection(ANewNode: TTreeNode; ASelect: Boolean);
begin
  if tvoAllowMultiSelect in FOptions then
  begin
    if ASelect then
      FTreeNodes.MultiSelect(ANewNode, False)
    else begin
      FTreeNodes.SelectOnlyThis(ANewNode);
    end;
  end else
    Selected := ANewNode;
  ANewNode.MakeVisible;
end;

procedure TCustomTreeView.MouseLeave;
begin
  FStates:=FStates-[tvsDblClicked,tvsTripleClicked,tvsQuadClicked,
                    tvsEditOnMouseUp,tvsSingleSelectOnMouseUp];
  if Assigned(FHintWnd) and FHintWnd.Visible
  and ((WidgetSet.GetLCLCapability(lcTransparentWindow) = LCL_CAPABILITY_YES)
       or not PtInRect(ClientRect, ScreenToClient(Mouse.CursorPos))) then
    FHintWnd.Hide;

  if tvoHotTrack in FOptions then
  begin
    FNodeUnderCursor:=nil;
    Cursor:=crDefault;
    Invalidate;
  end;

  inherited MouseLeave;
end;

procedure TCustomTreeView.NodeChanged(Node: TTreeNode; ChangeReason: TTreeNodeChangeReason);
begin
  if assigned(FOnNodeChanged) then
     OnNodeChanged(self,Node,ChangeReason);
end;

procedure TCustomTreeView.Notification(AComponent: TComponent; Operation: TOperation);
begin
  inherited Notification(AComponent, Operation);
  if Operation = opRemove then begin
    if AComponent = Images then Images := nil;
    if AComponent = StateImages then StateImages := nil;
  end;
end;

procedure TCustomTreeView.SetImages(Value: TCustomImageList);
var
  AImageHeight: Integer;
begin
  if Images = Value then exit;
  if Images <> nil then
    Images.UnRegisterChanges(FImageChangeLink);
  FImages := Value;
  if Images <> nil then begin
    Images.RegisterChanges(FImageChangeLink);
    Images.FreeNotification(Self);
    AImageHeight := Images.HeightForPPI[ImagesWidth, Font.PixelsPerInch];
    if DefaultItemHeight<AImageHeight+FDefItemSpace then
      DefaultItemHeight:=AImageHeight+FDefItemSpace;
  end;
  Invalidate;
end;

procedure TCustomTreeView.SetImagesWidth(const aImagesWidth: Integer);
begin
  if FImagesWidth = aImagesWidth then Exit;
  FImagesWidth := aImagesWidth;
  Invalidate;
end;

procedure TCustomTreeView.SetInsertMarkNode(const AValue: TTreeNode);
var
  InvalidateNeeded: Boolean;
begin
  if FInsertMarkNode=AValue then exit;
  InvalidateNeeded:=IsInsertMarkVisible;
  FInsertMarkNode:=AValue;
  InvalidateNeeded:=InvalidateNeeded or IsInsertMarkVisible;
  if InvalidateNeeded then Invalidate;
end;

procedure TCustomTreeView.SetInsertMarkType(
  const AValue: TTreeViewInsertMarkType);
var
  InvalidateNeeded: Boolean;
begin
  if FInsertMarkType=AValue then exit;
  InvalidateNeeded:=IsInsertMarkVisible;
  FInsertMarkType:=AValue;
  InvalidateNeeded:=InvalidateNeeded or IsInsertMarkVisible;
  if InvalidateNeeded then Invalidate;
end;

procedure TCustomTreeView.SetStateImages(Value: TCustomImageList);
var
  AStateImageHeight: Integer;
begin
  if FStateImages=Value then exit;
  if StateImages <> nil then
    StateImages.UnRegisterChanges(FStateChangeLink);
  FStateImages := Value;
  if StateImages <> nil then begin
    StateImages.RegisterChanges(FStateChangeLink);
    StateImages.FreeNotification(Self);
    AStateImageHeight := StateImages.HeightForPPI[StateImagesWidth, Font.PixelsPerInch];
    if DefaultItemHeight<AStateImageHeight+FDefItemSpace then
      DefaultItemHeight:=AStateImageHeight+FDefItemSpace;
  end;
  Invalidate;
end;

procedure TCustomTreeView.SetStateImagesWidth(const aStateImagesWidth: Integer);
begin
  if FStateImagesWidth = aStateImagesWidth then Exit;
  FStateImagesWidth := aStateImagesWidth;
  Invalidate;
end;

procedure TCustomTreeView.LoadFromFile(const FileName: string);
var
  Stream: TStream;
begin
  Stream := TFileStreamUTF8.Create(FileName, fmOpenRead or fmShareDenyWrite);
  try
    LoadFromStream(Stream);
  finally
    Stream.Free;
  end;
end;

procedure TCustomTreeView.LoadFromStream(Stream: TStream);
begin
  with TTreeStrings.Create(Items) do
    try
      LoadTreeFromStream(Stream);
    finally
      Free;
  end;
end;

procedure TCustomTreeView.SaveToFile(const FileName: string);
var
  Stream: TStream;
begin
  Stream := TFileStreamUTF8.Create(FileName, fmCreate);
  try
    SaveToStream(Stream);
  finally
    Stream.Free;
  end;
end;

procedure TCustomTreeView.SaveToStream(Stream: TStream);
begin
  with TTreeStrings.Create(Items) do
    try
      SaveTreeToStream(Stream);
    finally
      Free;
  end;
end;

procedure TCustomTreeView.ScrollView(DeltaX, DeltaY: Integer);
var
  ScrollArea: TRect;
  ScrollFlags: Integer;
begin
  if (DeltaX=0) and (DeltaY=0) then
    Exit;

  Include(FStates,tvsScrollbarChanged);
  ScrollFlags := SW_INVALIDATE or SW_ERASE;
  ScrollArea := ClientRect;
  InflateRect(ScrollArea, -BorderWidth, -BorderWidth);
  ScrollWindowEx(Handle, DeltaX, DeltaY, @ScrollArea, @ScrollArea, 0, nil, ScrollFlags);
end;

procedure TCustomTreeView.WMVScroll(var Msg: TLMScroll);
begin
  case Msg.ScrollCode of
      // Scrolls to start / end of the text
    SB_TOP:        ScrolledTop := 0;
    SB_BOTTOM:     ScrolledTop := GetMaxScrollTop;
      // Scrolls one line up / down
    SB_LINEDOWN:   ScrolledTop := ScrolledTop + FDefItemHeight;
    SB_LINEUP:     ScrolledTop := ScrolledTop - FDefItemHeight;
      // Scrolls one page of lines up / down
    SB_PAGEDOWN:   ScrolledTop := ScrolledTop + ClientHeight
                                     - FDefItemHeight;
    SB_PAGEUP:     ScrolledTop := ScrolledTop - ClientHeight
                                     + FDefItemHeight;
      // Scrolls to the current scroll bar position
    SB_THUMBPOSITION,
    SB_THUMBTRACK: ScrolledTop := Msg.Pos;

    SB_ENDSCROLL: ; // Ends scrolling
  end;
end;

procedure TCustomTreeView.WMHScroll(var Msg: TLMScroll);
begin
  case Msg.ScrollCode of
      // Scrolls to start / end of the text
    SB_LEFT:       ScrolledLeft := 0;
    SB_RIGHT:      ScrolledLeft := GetMaxScrollLeft;
      // Scrolls one line left / right
    SB_LINERIGHT:  ScrolledLeft := ScrolledLeft + FDefItemHeight div 2;
    SB_LINELEFT:   ScrolledLeft := ScrolledLeft - FDefItemHeight div 2;
      // Scrolls one page of lines left / right
    SB_PAGERIGHT:  ScrolledLeft := ScrolledLeft + ClientHeight
                                       - FDefItemHeight;
    SB_PAGELEFT:   ScrolledLeft := ScrolledLeft - ClientHeight
                                       + FDefItemHeight;
      // Scrolls to the current scroll bar position
    SB_THUMBPOSITION,
    SB_THUMBTRACK: ScrolledLeft := Msg.Pos;

    SB_ENDSCROLL: ;// Ends scrolling
  end;
end;

procedure TCustomTreeView.WMLButtonDown(var AMessage: TLMLButtonDown);
begin
  {$IFDEF VerboseDrag}
  DebugLn('TCustomTreeView.WMLButtonDown A ',Name,':',ClassName,' ');
  {$ENDIF}
  Exclude(FStates,tvsDragged);
  inherited WMLButtonDown(AMessage);
  {$IFDEF VerboseDrag}
  DebugLn('TCustomTreeView.WMLButtonDown END ',Name,':',ClassName,' ');
  {$ENDIF}
end;

procedure TCustomTreeView.WMSetFocus(var Message: TLMSetFocus);
begin
  Invalidate;
  inherited;
end;

procedure TCustomTreeView.WMKillFocus(var Message: TLMKillFocus);
begin
  Invalidate;
  inherited;
end;

procedure TCustomTreeView.Resize;
begin
  FStates:=FStates+[tvsScrollbarChanged,tvsBottomItemNeedsUpdate];
  inherited Resize;
end;

function TCustomTreeView.GetSelectedChildAccessibleObject: TLazAccessibleObject;
var
  lNode: TTreeNode;
begin
  Result := nil;
  lNode := GetSelection();
  if lNode = nil then Exit;
  Result := FAccessibleObject.GetChildAccessibleObjectWithDataObject(lNode);
end;

function TCustomTreeView.GetChildAccessibleObjectAtPos(APos: TPoint): TLazAccessibleObject;
var
  lNode: TTreeNode;
begin
  Result := nil;
  lNode := GetNodeAt(APos.X, APos.Y);
  //if lNode = nil then DebugLn('[TCustomTreeView.GetChildAccessibleObjectAtPos] lNode=nil')
  //else DebugLn('[TCustomTreeView.GetChildAccessibleObjectAtPos] lNode=' + lNode.Text);
  if lNode = nil then Exit;
  Result := FAccessibleObject.GetChildAccessibleObjectWithDataObject(lNode);
end;

procedure TCustomTreeView.InternalSelectionChanged;
begin
  if FSelectionChangeEventLock > 0 then
    Include(FStates, tvsSelectionChanged)
  else
  begin
    Exclude(FStates, tvsSelectionChanged);
    DoSelectionChanged;
    FChangeTimer.Enabled := False;
    if not FCallingChange then  // Prevent recursive loop from OnChange handler.
      FChangeTimer.Enabled := True;
  end;
end;

function TCustomTreeView.AllowMultiSelectWithCtrl(AState: TShiftState): Boolean;
begin
  Result := (ssCtrl in AState) and (msControlSelect in FMultiSelectStyle);
end;

function TCustomTreeView.AllowMultiSelectWithShift(AState: TShiftState): Boolean;
begin
  Result := (ssShift in AState) and (msShiftSelect in FMultiSelectStyle);
end;

class procedure TCustomTreeView.WSRegisterClass;
begin
  inherited WSRegisterClass;
  RegisterCustomTreeView;
end;

class function TCustomTreeView.GetControlClassDefaultSize: TSize;
begin
  Result.CX := 121;
  Result.CY := 97;
end;

procedure TCustomTreeView.Added(Node: TTreeNode);
begin
  if Assigned(OnAddition) then OnAddition(Self,Node);
end;

{ CustomDraw support }

procedure TCustomTreeView.EditorEditingDone(Sender: TObject);
var
  WasFocused: Boolean;
begin
  WasFocused := (FEditor<>nil) and FEditor.Focused;
  EndEditing;
  if WasFocused then
    SetFocus;
end;

procedure TCustomTreeView.EditorKeyDown(Sender: TObject; var Key: Word;
  Shift: TShiftState);
var
  WasFocused: Boolean;
begin
  if (Key = VK_ESCAPE) or (Key = VK_RETURN) then
  begin
    WasFocused := Assigned(FEditor) and FEditor.Focused;
    EndEditing(Key = VK_ESCAPE);
    if WasFocused then
      SetFocus;
    Key := 0; // key was handled
  end;
end;

procedure TCustomTreeView.CanvasChanged(Sender: TObject);
begin
  Include(FStates,tvsCanvasChanged);
end;

procedure TCustomTreeView.DragScrollTimerTick(Sender: TObject);
const
  cScrollDelta = 10;
var
  Pnt: TPoint;
begin
  Pnt := ScreenToClient(Mouse.CursorPos);
  if (Pnt.X >= 0) and (Pnt.X < ClientWidth) then
  begin
    if (Pnt.Y >= 0) and (Pnt.Y < FDragScrollMargin) then
      ScrolledTop := ScrolledTop - cScrollDelta
    else
    if (Pnt.Y >= ClientHeight-FDragScrollMargin) and (Pnt.Y < ClientHeight) then
      ScrolledTop := ScrolledTop + cScrollDelta;
  end;
end;

procedure TCustomTreeView.ClearSelection(KeepPrimary: Boolean);
begin
 if tvoAllowMultiSelect in FOptions then
   Items.ClearMultiSelection(not KeepPrimary)
 else
   if not KeepPrimary then Selected := nil;
end;

function TCustomTreeView.IsCustomDrawn(Target: TCustomDrawTarget;
  Stage: TCustomDrawStage): Boolean;
begin
  { Tree view doesn't support erase notifications }
  if Stage = cdPrePaint then begin
    if Target = dtItem then
      Result := Assigned(FOnCustomDrawItem)
                or Assigned(FOnAdvancedCustomDrawItem)
    else if Target = dtControl then
      Result := Assigned(FOnCustomDraw) or Assigned(FOnAdvancedCustomDraw) or
              Assigned(FOnCustomDrawItem) or Assigned(FOnAdvancedCustomDrawItem)
    else
      Result := False;
  end else begin
    if Target = dtItem then
      Result := Assigned(FOnAdvancedCustomDrawItem)
    else if Target = dtControl then
      Result := Assigned(FOnAdvancedCustomDraw)
                or Assigned(FOnAdvancedCustomDrawItem)
    else
      Result := False;
  end;
end;

function TCustomTreeView.CustomDraw(const ARect: TRect;
  Stage: TCustomDrawStage): Boolean;
begin
  Result := True;
  if (Stage = cdPrePaint) and Assigned(FOnCustomDraw) then
    FOnCustomDraw(Self, ARect, Result);
  if Assigned(FOnAdvancedCustomDraw) then
    FOnAdvancedCustomDraw(Self, ARect, Stage, Result);
end;

function TCustomTreeView.CustomDrawItem(Node: TTreeNode;
  State: TCustomDrawState;
  Stage: TCustomDrawStage; var PaintImages: Boolean): Boolean;
begin
  Result := True;
  PaintImages := True;
  if (Stage = cdPrePaint) and Assigned(FOnCustomDrawItem) then
    FOnCustomDrawItem(Self, Node, State, Result);
  if Assigned(FOnAdvancedCustomDrawItem) then
    FOnAdvancedCustomDrawItem(Self, Node, State, Stage, PaintImages, Result);
end;

procedure TCustomTreeView.ConsistencyCheck;
var OldMaxRight, OldLastTop, OldMaxLvl: integer;
  OldTopItem, OldBottomItem: TTreeNode;
begin
  if Canvas=nil then
    RaiseGDBException('Canvas=nil');
  if FDefItemHeight<0 then
    RaiseGDBException('FDefItemHeight='+IntToStr(FDefItemHeight));
  if FIndent<0 then
    RaiseGDBException('FIndent='+IntToStr(FIndent));
  if FMaxRight<0 then
    RaiseGDBException('FMaxRight='+IntToStr(FMaxRight));
  if FTreeNodes=nil then
    RaiseGDBException('FTreeNodes=nil');
  FTreeNodes.ConsistencyCheck;
  if Items.FUpdateCount<0 then
    RaiseGDBException('FUpdateCount='+IntToStr(Items.FUpdateCount));
  if (not (tvsTopsNeedsUpdate in FStates)) then begin
    if Items.GetLastSubNode<>nil then begin
      OldLastTop:=Items.GetLastSubNode.Top;
      Include(FStates,tvsTopsNeedsUpdate);
      UpdateAllTops;
      if OldLastTop<>Items.GetLastSubNode.Top then
        RaiseGDBException('OldLastTop='+DbgS(OldLastTop)
        +'<>Items.GetLastSubNode.Top='+DbgS(Items.GetLastSubNode.Top));
    end;
  end;
  if not (tvsMaxRightNeedsUpdate in FStates) then begin
    OldMaxRight:=FMaxRight;
    Include(FStates,tvsMaxRightNeedsUpdate);
    UpdateMaxRight;
    if OldMaxRight<>FMaxRight then
      RaiseGDBException('OldMaxRight<>FMaxRight');
  end;
  if not (tvsMaxLvlNeedsUpdate in FStates) then begin
    OldMaxLvl:=FMaxLvl;
    Include(FStates,tvsMaxLvlNeedsUpdate);
    UpdateMaxLvl;
    if OldMaxLvl<>FMaxLvl then
      RaiseGDBException('OldMaxLvl<>FMaxLvl');
  end;
  if (tvsIsEditing in FStates) and (FSelectedNode=nil) then
    RaiseGDBException('');
  if (FSelectedNode<>nil) then begin
    if not FSelectedNode.IsVisible then
      RaiseGDBException('not FSelectedNode.IsVisible');
  end;
  if not (tvsTopItemNeedsUpdate in FStates) then begin
    OldTopItem:=FTopItem;
    UpdateTopItem;
    if FTopItem<>OldTopItem then
      RaiseGDBException('FTopItem<>OldTopItem');
  end;
  if not (tvsBottomItemNeedsUpdate in FStates) then begin
    OldBottomItem:=FBottomItem;
    UpdateBottomItem;
    if FBottomItem<>OldBottomItem then
      RaiseGDBException('FBottomItem<>OldBottomItem');
  end;
end;

procedure TCustomTreeView.WriteDebugReport(const Prefix: string; AllNodes: boolean);
begin
  DbgOut('%s%s.WriteDebugReport Self=%p', [Prefix, ClassName, Pointer(Self)]);
  ConsistencyCheck;
  DebugLn('');
  if AllNodes then begin
    Items.WriteDebugReport(Prefix+'  ',true);
  end;
end;

procedure TCustomTreeView.LockSelectionChangeEvent;
begin
  inc(FSelectionChangeEventLock);
end;

procedure TCustomTreeView.UnlockSelectionChangeEvent;
begin
  dec(FSelectionChangeEventLock);
  if FSelectionChangeEventLock<0 then
    RaiseGDBException('TCustomTreeView.UnlockSelectionChangeEvent');
  if (FSelectionChangeEventLock=0) and (tvsSelectionChanged in FStates) then
    InternalSelectionChanged;
end;

function TCustomTreeView.GetFirstMultiSelected: TTreeNode;
begin
  Result := Items.FFirstMultiSelected;
end;

function TCustomTreeView.GetLastMultiSelected: TTreeNode;
begin
  Result := Items.FLastMultiSelected;
end;

procedure TCustomTreeView.Select(Node: TTreeNode; ShiftState: TShiftState = []);
begin
  if (tvoAllowMultiSelect in FOptions) and AllowMultiSelectWithCtrl(ShiftState) then
    Node.Selected := True
  else begin
    ClearSelection;
    Selected := Node;
    if (tvoAllowMultiSelect in FOptions) then
      Node.Selected := True;
  end;
end;

procedure TCustomTreeView.Select(const Nodes: array of TTreeNode);
var
  I: Integer;
begin
  ClearSelection;
  if Length(Nodes)>0 then begin
    Selected := Nodes[0];
    if tvoAllowMultiSelect in FOptions then
      for I := Low(Nodes) to High(Nodes) do
        Nodes[I].Selected := True;
  end;
end;

procedure TCustomTreeView.Select(Nodes: TList);
var
  I: Integer;
begin
  ClearSelection;
  if Nodes.Count>0 then begin
    Selected := TTreeNode(Nodes[0]);
    if tvoAllowMultiSelect in FOptions then
      for I := 0 to Nodes.Count - 1 do
        TTreeNode(Nodes[I]).Selected := True;
  end;
end;

function TCustomTreeView.SelectionVisible: boolean;
var
  ANode: TTreeNode;
begin
  ANode:=GetFirstMultiSelected;
  if (ANode<>nil) and (ANode.GetNextMultiSelected<>nil) then begin
    // 2 or more elements => a real multi selection =>
    // is visible if even one of its nodes is partly visible
    while (ANode<>nil) do begin
      if ANode.IsVisible then begin
        Result:=true;
        exit;
      end;
      ANode:=ANode.GetNextMultiSelected;
    end;
    Result:=false;
  end else begin
    if ANode=nil then
      ANode:=Selected;
    Result:=(ANode<>nil) and (ANode.IsFullHeightVisible);
  end;
end;

procedure TCustomTreeView.MakeSelectionVisible;
var
  ANode: TTreeNode;
begin
  if SelectionVisible then exit;
  ANode:=GetFirstMultiSelected;
  if (ANode=nil) then
    ANode:=Selected;
  if ANode=nil then exit;
  ANode.MakeVisible;
end;

procedure TCustomTreeView.ClearInvisibleSelection;
var
  ANode: TTreeNode;
begin
  if tvoAllowMultiSelect in FOptions then begin
    Items.ClearMultiSelection(True);       // Now clears all multi-selected
  end
  else begin
    ANode := Selected;            // Clear a single selection only if not visible
    if Assigned(ANode) and not ANode.Visible then
      ANode.Selected:=False;       // Selected := nil;
  end;
end;

procedure TCustomTreeView.MoveToNextNode(ASelect: Boolean);
var
  ANode: TTreeNode;
begin
  if tvoAllowMultiSelect in FOptions then
    ANode := FTreeNodes.FLastMultiSelected
  else
    ANode := Selected;
  if ANode <> nil then
    ANode := ANode.GetNextVisible
  else
    ANode := FTreeNodes.GetFirstVisibleNode;
  if ANode <> nil then
    MoveSelection(ANode, ASelect);
end;

procedure TCustomTreeView.MoveToPrevNode(ASelect: Boolean);
var
  ANode: TTreeNode;
begin
  if tvoAllowMultiSelect in FOptions then
    ANode := FTreeNodes.FLastMultiSelected
  else
    ANode := Selected;
  if ANode <> nil then
    ANode := ANode.GetPrevVisible
  else
    ANode := Items.GetFirstVisibleNode;
  if ANode <> nil then
    MoveSelection(ANode, ASelect);
end;

function TCustomTreeView.StoreCurrentSelection: TStringList;
var
  ANode: TTreeNode;
begin
  Result:=TStringList.Create;
  ANode:=Selected;
  while ANode<>nil do begin
    Result.Insert(0,ANode.Text);
    ANode:=ANode.Parent;
  end;
end;

procedure TCustomTreeView.ApplyStoredSelection(ASelection: TStringList; FreeList: boolean);
var
  ANode: TTreeNode;
  CurText: string;
begin
  ANode:=nil;
  while ASelection.Count>0 do begin
    CurText:=ASelection[0];
    if ANode=nil then
      ANode:=Items.GetFirstNode
    else
      ANode:=ANode.GetFirstChild;
    while (ANode<>nil) and (ANode.Text<>CurText) do
      ANode:=ANode.GetNextSibling;
    if ANode=nil then break;
    ASelection.Delete(0);
  end;
  if ANode<>nil then
    Selected:=ANode;
  if FreeList then
    ASelection.Free;
end;

// back to comctrls.pp
