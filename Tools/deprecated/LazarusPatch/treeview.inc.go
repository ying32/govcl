package main

var (
	treeViewIncData = []PatchItem{
		{Find1: []byte("Filer.DefineBinaryProperty('Data', @ReadData, @WriteData, WriteNodes);"), Find2: nil, Content: []byte("  Filer.DefineBinaryProperty('NodeData', @ReadDelphiNodeData, nil, False);")},
		{Find1: []byte("procedure TTreeNodes.WriteData(Stream: TStream);"), Find2: nil, Content: []byte(`

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

`)},
	}
)

// Filer.DefineBinaryProperty('Data', @ReadData, @WriteData, WriteNodes);
// 添加 Filer.DefineBinaryProperty('NodeData', @ReadDelphiNodeData, nil, False);

// 查找
// procedure TTreeNodes.WriteData(Stream: TStream);

// 添加
/*

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

*/
