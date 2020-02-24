package main

var (
	listItemsIncData = []PatchItem{
		{Find1: []byte("Filer.DefineBinaryProperty('Data', @ReadData, nil, false);"), Find2: nil, Content: []byte("  Filer.DefineBinaryProperty('ItemData', @ReadItemData, nil, False);")},
		{Find1: []byte("procedure TListItems.ReadLazData(Stream: TStream);"), Find2: nil, Content: []byte(`

// from Delphi Vcl.ComCtrl.pas, 10.2.1
const
  //ListItemStreamVersion1    = $01; // 32-bit struct size version 1
  //ListItemStreamVersion1x64 = $02; // 64-bit struct size version 1
  ListItemStreamVersion2    = $03; // 32-bit struct size version 2
  //ListItemStreamVersion2x64 = $04; // 64-bit struct size version 2
  ListItemStreamVersion3    = $05; // 32-bit struct size version 3
  //ListItemStreamVersion3x64 = $06; // 64-bit struct size version 3


procedure TListItems.ReadItemData(Stream: TStream);
var
  I, J, LSize, LItemCount: Integer;
  LImageIndex: SmallInt;
  LStreamVersion, LLen: Byte;
  LWideStr: UnicodeString;
  LSubItemData: Pointer;

  // 按顺序读
  //LImageIndex: Integer;
  //LStateIndex: Integer;
  //LOverlayIndex: Integer;
  LSubItemCount: Integer;
  //LGroupID: Integer;
  //LData: Integer;
  LCaptionLen: Byte;

  function ReadInteger: Integer;
  begin
    Stream.Read(Result, SizeOf(Integer));
  end;

begin
  Clear;
  if Stream.Size = 0 then
    Exit;
  LStreamVersion := Stream.ReadByte;
  case LStreamVersion of

    ListItemStreamVersion2,
    ListItemStreamVersion3:
      begin
        Stream.ReadBuffer(LSize, SizeOf(Integer)); // 这个值不要的
        Stream.ReadBuffer(LItemCount, SizeOf(Integer));
        for I := 0 to LItemCount - 1 do
        begin
          with Add do
          begin
            ImageIndex := ReadInteger;
            StateIndex := ReadInteger;
            //OverlayIndex
            ReadInteger;
            LSubItemCount := ReadInteger;
            // GroupID
            ReadInteger;
            Data := Pointer(ReadInteger);

            // 读标题长度
            LCaptionLen := Stream.ReadByte;
            // 读标题
            SetLength(LWideStr, LCaptionLen);
            Stream.ReadBuffer(LWideStr[1], LCaptionLen * SizeOf(WideChar));
            Caption := UTF16ToUTF8(LWideStr);

            // 读子项目的文本
            for J := 0 to LSubItemCount - 1 do
            begin
              LLen := Stream.ReadByte;
              SetLength(LWideStr, LLen);
              Stream.ReadBuffer(LWideStr[1], LLen * SizeOf(WideChar));
              LSubItemData := nil;
              if LStreamVersion = ListItemStreamVersion3 then
                LSubItemData := Pointer(Stream.ReadDWORD);
              SubItems.AddObject(UTF16ToUTF8(LWideStr), TObject(LSubItemData));
            end;
          end;
        end;
      end;
  else
    Exit;
  end;

  // Read subitem images, if present
  for I := 0 to Count - 1 do
    with Item[I] do
      for J := 0 to SubItems.Count - 1 do
      begin
        Stream.ReadBuffer(LImageIndex, SizeOf(SmallInt));
        SubItemImages[J] := LImageIndex;
      end;
end;

`)},
	}
)

// 查找
// Filer.DefineBinaryProperty('Data', @ReadData, nil, false);

// 添加
// Filer.DefineBinaryProperty('ItemData', @ReadItemData, nil, False);

// 查找
// procedure TListItems.ReadLazData(Stream: TStream);

// 添加
/*

// from Delphi Vcl.ComCtrl.pas, 10.2.1
const
  //ListItemStreamVersion1    = $01; // 32-bit struct size version 1
  //ListItemStreamVersion1x64 = $02; // 64-bit struct size version 1
  ListItemStreamVersion2    = $03; // 32-bit struct size version 2
  //ListItemStreamVersion2x64 = $04; // 64-bit struct size version 2
  ListItemStreamVersion3    = $05; // 32-bit struct size version 3
  //ListItemStreamVersion3x64 = $06; // 64-bit struct size version 3


procedure TListItems.ReadItemData(Stream: TStream);
var
  I, J, LSize, LItemCount: Integer;
  LImageIndex: SmallInt;
  LStreamVersion, LLen: Byte;
  LWideStr: UnicodeString;
  LSubItemData: Pointer;

  // 按顺序读
  //LImageIndex: Integer;
  //LStateIndex: Integer;
  //LOverlayIndex: Integer;
  LSubItemCount: Integer;
  //LGroupID: Integer;
  //LData: Integer;
  LCaptionLen: Byte;

  function ReadInteger: Integer;
  begin
    Stream.Read(Result, SizeOf(Integer));
  end;

begin
  Clear;
  if Stream.Size = 0 then
    Exit;
  LStreamVersion := Stream.ReadByte;
  case LStreamVersion of

    ListItemStreamVersion2,
    ListItemStreamVersion3:
      begin
        Stream.ReadBuffer(LSize, SizeOf(Integer)); // 这个值不要的
        Stream.ReadBuffer(LItemCount, SizeOf(Integer));
        for I := 0 to LItemCount - 1 do
        begin
          with Add do
          begin
            ImageIndex := ReadInteger;
            StateIndex := ReadInteger;
            //OverlayIndex
            ReadInteger;
            LSubItemCount := ReadInteger;
            // GroupID
            ReadInteger;
            Data := Pointer(ReadInteger);

            // 读标题长度
            LCaptionLen := Stream.ReadByte;
            // 读标题
            SetLength(LWideStr, LCaptionLen);
            Stream.ReadBuffer(LWideStr[1], LCaptionLen * SizeOf(WideChar));
            Caption := UTF16ToUTF8(LWideStr);

            // 读子项目的文本
            for J := 0 to LSubItemCount - 1 do
            begin
              LLen := Stream.ReadByte;
              SetLength(LWideStr, LLen);
              Stream.ReadBuffer(LWideStr[1], LLen * SizeOf(WideChar));
              LSubItemData := nil;
              if LStreamVersion = ListItemStreamVersion3 then
                LSubItemData := Pointer(Stream.ReadDWORD);
              SubItems.AddObject(UTF16ToUTF8(LWideStr), TObject(LSubItemData));
            end;
          end;
        end;
      end;
  else
    Exit;
  end;

  // Read subitem images, if present
  for I := 0 to Count - 1 do
    with Item[I] do
      for J := 0 to SubItems.Count - 1 do
      begin
        Stream.ReadBuffer(LImageIndex, SizeOf(SmallInt));
        SubItemImages[J] := LImageIndex;
      end;
end;
*/
