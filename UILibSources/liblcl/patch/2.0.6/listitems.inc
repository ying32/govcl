{%MainUnit ../comctrls.pp}
{ $Id: listitems.inc 51595 2016-02-11 05:39:03Z ondrej $

 *****************************************************************************
  This file is part of the Lazarus Component Library (LCL)

  See the file COPYING.modifiedLGPL.txt, included in this distribution,
  for details about the license.
 *****************************************************************************
}

{------------------------------------------------------------------------------}
{   TListItems Constructor                                                     }
{------------------------------------------------------------------------------}
constructor TListItems.Create(AOwner : TCustomListView);
begin
  Inherited Create;
  FItems := TFPList.Create;
  FOwner := AOwner;
  FCacheIndex := -1;
end;

{------------------------------------------------------------------------------}
{   TListItems GetCount                                                        }
{------------------------------------------------------------------------------}
function TListItems.GetCount: Integer;
begin
  Result := FItems.Count;
end;

procedure TListItems.SetCount(const ACount: Integer); 
begin
  //just ignore!
end;

{------------------------------------------------------------------------------}
{   TListItems GetItem                                                         }
{------------------------------------------------------------------------------}
function TListItems.GetItem(const AIndex: Integer): TListItem;
begin
  if (FCacheIndex <> -1) and (FCacheIndex = AIndex)
  then begin
    Result := FCacheItem;
    Exit;
  end;

  if FItems.Count - 1 < AIndex
  then Result := nil
  else begin
    Result := TListItem(FItems.Items[AIndex]);
    FCacheItem := Result;
    FCacheIndex := AIndex;
  end;
end;

function TListItems.GetOwner: TPersistent;
begin
  Result := FOwner;
end;

{------------------------------------------------------------------------------}
{   TListItems SetItem                                                         }
{------------------------------------------------------------------------------}
procedure TListItems.SetItem(const AIndex: Integer; const AValue: TListItem);
var
  OldItem: TListItem;
begin
  if Count - 1 < AIndex then Exit;
  
  OldItem := GetItem(AIndex);
  if OldItem = AValue then Exit;
  
  FItems.Items[AIndex] := AValue;
  FCacheIndex := AIndex;
  FCacheItem := AValue;

  if WSUpdateAllowed then
    WSUpdateItem(AIndex, AValue);

  if AValue.WSUpdateAllowed
  then begin
    AValue.WSUpdateText;
    AValue.WSUpdateImages;
    AValue.WSUpdateChecked;
    AValue.WSUpdateState;
  end;
end;

{------------------------------------------------------------------------------}
{   TListItems Add                                                             }
{------------------------------------------------------------------------------}
function TListItems.Add: TListItem;
begin
  if Assigned(Owner) then
    Result := Owner.CreateListItem
  else
    Result := TListItem.Create(Self);
  AddItem(Result);
end;

{------------------------------------------------------------------------------}
{   TListItems AddItem                                                         }
{------------------------------------------------------------------------------}
procedure TListItems.AddItem(AItem: TListItem);
begin
  FCacheIndex := FItems.Add(AItem);
  FCacheItem := AItem;

  if WSUpdateAllowed
  then WSCreateCacheItem;

  //Notify parent TListView that something was added.
  if FOwner <> nil
  then FOwner.ItemInserted(AItem);
end;

procedure TListItems.BeginUpdate;
begin
  Owner.BeginUpdate;
end;

{------------------------------------------------------------------------------
   TListItems Clear
------------------------------------------------------------------------------}
procedure TListItems.Clear;
begin
  while Count > 0 do Delete(Count-1)
end;

{------------------------------------------------------------------------------}
{   TListItems Delete                                                          }
{------------------------------------------------------------------------------}
procedure TListItems.Delete(const AIndex: Integer);
begin
  // Don't use GetItem, it updates the cache, which then will become invalid
  TListItem(FItems.Items[AIndex]).Delete;
end;

procedure TListItems.EndUpdate;
begin
  Owner.EndUpdate;
end;

procedure TListItems.Exchange(const AIndex1, AIndex2: Integer);
var
  AItem: TListItem;
  SelItem: TListItem;
  FocusItem: TListItem;
begin
  if AIndex1 = AIndex2 then
    exit;
  if (AIndex1 < 0) or (AIndex1 >= FItems.Count) then
    raise Exception.CreateFmt(rsListIndexExceedsBounds, [AIndex1]);
  if (AIndex2 < 0) or (AIndex2 >= FItems.Count) then
    raise Exception.CreateFmt(rsListIndexExceedsBounds, [AIndex2]);
  AItem := Item[AIndex1];
  SelItem := Owner.Selected;
  FocusItem := Owner.ItemFocused;

  Include(Owner.FFlags, lffItemsMoving);
  try
    FItems.Exchange(AIndex1, AIndex2);
    FCacheIndex := AIndex1;
    FCacheItem := AItem;
    if not Owner.OwnerData and WSUpdateAllowed then
    begin
      TWSCustomListViewClass(FOwner.WidgetSetClass).ItemExchange(FOwner, AItem,
        AIndex1, AIndex2);
      Owner.InvalidateSelected;
      Owner.ItemFocused := nil;
      Owner.Selected := SelItem;
      Owner.ItemFocused := FocusItem;
    end;
  finally
    Exclude(Owner.FFlags, lffItemsMoving);
  end;
end;

procedure TListItems.Move(const AFromIndex, AToIndex: Integer);
var
  AItem: TListItem;
  SelItem: TListItem;
  FocusItem: TListItem;
begin
  if AFromIndex = AToIndex then
    exit;
  if (AFromIndex < 0) or (AFromIndex >= FItems.Count) then
    raise Exception.CreateFmt(rsListIndexExceedsBounds, [AFromIndex]);
  if (AToIndex < 0) or (AToIndex >= FItems.Count) then
    raise Exception.CreateFmt(rsListIndexExceedsBounds, [AToIndex]);
  AItem := Item[AFromIndex];
  SelItem := Owner.Selected;
  FocusItem := Owner.ItemFocused;
  Include(Owner.FFlags, lffItemsMoving);
  try
    FItems.Move(AFromIndex, AToIndex);
    FCacheIndex := AToIndex;
    FCacheItem := AItem;
    if not Owner.OwnerData and WSUpdateAllowed then
    begin
      TWSCustomListViewClass(FOwner.WidgetSetClass).ItemMove(FOwner, AItem,
        AFromIndex, AToIndex);
      Owner.InvalidateSelected;
      Owner.ItemFocused := nil;
      Owner.Selected := SelItem;
      Owner.ItemFocused := FocusItem;
    end;
  finally
    Exclude(Owner.FFlags, lffItemsMoving);
  end;
end;

{------------------------------------------------------------------------------}
{   TListItems IntfCreateItem                                                  }
{------------------------------------------------------------------------------}
procedure TListItems.WSCreateCacheItem;
var
  ACacheItem: TListItem;
begin
  ACacheItem := FCacheItem; // we must use local copy because FCacheItem changes in WSUpdateText on win32
  TWSCustomListViewClass(FOwner.WidgetSetClass).ItemInsert(FOwner, FCacheIndex, ACacheItem);
  Include(ACacheItem.FFlags, lifCreated);
  ACacheItem.WSUpdateText;
  ACacheItem.WSUpdateImages;
  ACacheItem.WSUpdateChecked;
  if not (lisfWSItemsCreated in FFlags) then
    ACacheItem.WSSetState;
  ACacheItem.WSUpdateState;
end;

{------------------------------------------------------------------------------}
{   TListItems IntfCreateItems                                                 }
{------------------------------------------------------------------------------}
procedure TListItems.WSCreateItems;
var
  n: integer;
begin
  for n := 0 to FItems.Count - 1 do
  begin
    FCacheItem := TListItem(FItems[n]);
    FCacheIndex := n;
    WSCreateCacheItem;
  end;
  Include(FFlags, lisfWSItemsCreated);
end;

procedure TListItems.DoFinalizeWnd;
begin
  Exclude(FFlags, lisfWSItemsCreated);
end;

{------------------------------------------------------------------------------}
{   TListItems IntfUpdateAllowed                                               }
{------------------------------------------------------------------------------}
function TListItems.WSUpdateAllowed: Boolean;
begin
  Result := (FOwner <> nil)
    and FOwner.HandleAllocated
    and not (csDestroying in FOwner.ComponentState);
end;

procedure TListItems.WSUpdateItem(const AIndex: Integer; const AValue: TListItem);
begin
  TWSCustomListViewClass(FOwner.WidgetSetClass).ItemUpdate(FOwner, AIndex, AValue);
end;

procedure TListItems.WSSetItemsCount(const ACount: Integer); 
begin
  TWSCustomListViewClass(FOwner.WidgetSetClass).SetItemsCount(FOwner, ACount);
end;

{------------------------------------------------------------------------------}
{   TListItems ItemDeleted                                                     }
{------------------------------------------------------------------------------}
procedure TListItems.ItemDestroying(const AItem: TListItem);
var
  idx: Integer;
begin
  // Don't use IndexOf, it updates the cache, which then will become invalid
  //DebugLn('TListItems.ItemDestroying ',dbgs(AItem));
  if  (FCacheIndex <> -1) and (FCacheItem = AItem) then
    idx := FCacheIndex
  else
    idx := FItems.IndexOf(AItem);

  if idx = -1 then Exit; //?????

  if FOwner <> nil then
  begin

    FOwner.ItemDeleted(AItem);

    // trigger DoSelectItem before deletion - Delphi compat. issue #21335
    // but fire after ItemIndex is invalidated. issue #21346
    if FOwner.HandleAllocated and AItem.Selected and
      not (csDestroying in FOwner.Componentstate) then
      FOwner.DoSelectItem(AItem, False);

    if (FOwner.HandleAllocated) and (lifCreated in AItem.FFlags) then
    begin
      Exclude(AItem.FFlags, lifCreated);
      TWSCustomListViewClass(FOwner.WidgetSetClass).ItemDelete(FOwner, idx);
    end;
  end;

  if FCacheIndex >= idx then
    FCacheIndex := -1;
  FItems.Delete(idx);
end;

{------------------------------------------------------------------------------
   TListItems IndexOf
------------------------------------------------------------------------------}
function TListItems.IndexOf(const AItem: TListItem): Integer;
begin
  if  (FCacheIndex <> -1) 
  and (FCacheItem = AItem)
  then begin
    Result := FCacheIndex;
    Exit;
  end;
  
  Result := FItems.IndexOf(AItem);
  if Result = -1 then Exit;
  
  FCacheIndex := Result;
  FCacheItem := AItem;
end;

{------------------------------------------------------------------------------
   TListItems Insert
------------------------------------------------------------------------------}
function TListItems.Insert(const AIndex: Integer): TListItem;
begin
  if Assigned(Owner) then
    Result := Owner.CreateListItem
  else
    Result := TListItem.Create(Self);
  InsertItem(Result, AIndex);
end;

procedure TListItems.SelectAll;
var
  i: Integer;
begin
  for i := 0 to Count - 1 do
  begin
    if not (lisSelected in Item[i].FStates) then
    begin
      Self.Item[i].FStates := Self.Item[i].FStates + [lisSelected];
      if Assigned(Owner.OnSelectItem) then
        Owner.OnSelectItem(Owner, Item[i], True);
    end;
  end;
end;

procedure TListItems.ClearSelection;
var
  i: Integer;
begin
  for i := 0 to Count - 1 do
  begin
    if (lisSelected in Item[i].FStates) then
    begin
      Self.Item[i].FStates := Self.Item[i].FStates - [lisSelected];
      if Assigned(Owner.OnSelectItem) then
        Owner.OnSelectItem(Owner, Item[i], False);
    end;
  end;
end;

{------------------------------------------------------------------------------
   TListItems InsertItem
------------------------------------------------------------------------------}
procedure TListItems.InsertItem(AItem: TListItem; const AIndex: Integer);
begin
  FItems.Insert(AIndex, AItem);
  FCacheIndex := AIndex;
  FCacheItem := AItem;

  if WSUpdateAllowed
  then WSCreateCacheItem;

  //Notify parent TListView that something was added.
  if FOwner <> nil
  then FOwner.ItemInserted(AItem);
end;

{------------------------------------------------------------------------------}
{   TListItems Destructor                                                      }
{------------------------------------------------------------------------------}
destructor TListItems.Destroy;
begin
  if not FOwner.FOwnerData then  
  begin
    // to call Self.Clear instead?
    FCacheIndex := 0;
    while FItems.Count > 0 do
    begin
      FCacheItem := TListItem(FItems[0]);
      FCacheItem.Free;
    end;
  end;
  FCacheIndex := -1;
  FreeAndNil(FItems);
  inherited Destroy;
end;

{------------------------------------------------------------------------------}
{   TListItems FindCaption                                                     }
{------------------------------------------------------------------------------}
function TListItems.FindCaption(StartIndex: Integer; Value: string;
  Partial, Inclusive, Wrap: Boolean; PartStart: Boolean): TListItem;
var
  I: Integer;
  CaptionFound, AllChecked: Boolean;
begin
  result := nil;
  if (Count = 0) or (StartIndex >= Count) or (not Inclusive and (count = 1)) then Exit;
  CaptionFound := False;
  AllChecked := False;
  if Inclusive then
    I := StartIndex
  else begin
    I := succ(StartIndex);
    if I >= Count then I := 0;
  end;
  if Wrap then Wrap := (StartIndex <> 0);
  repeat
    if Partial then begin
      if PartStart then
        CaptionFound := pos(Value, Item[I].Caption) = 1
      else
        CaptionFound := pos(Value, Item[I].Caption) <> 0;
    end else
      CaptionFound := Value = Item[I].Caption;
    if not CaptionFound then begin
      Inc(I);
      if Wrap then begin
        if I = Count then
          I := 0
        else
          if I = StartIndex then
            AllChecked := True;
      end else begin
        if I = Count then AllChecked := True;
      end;
    end;
  until CaptionFound or AllChecked;
  if CaptionFound then result := Item[I];
end;


{------------------------------------------------------------------------------}
{   TListItems FindData                                                        }
{------------------------------------------------------------------------------}
function TListItems.FindData(const AData: Pointer): TListItem;
var
  n: Integer;
begin
  if  (FCacheIndex <> -1) 
  and (FCacheItem <> nil) 
  and (FCacheItem.Data = AData)
  then begin
    Result := FCacheItem;
    Exit;
  end;
  
  for n := 0 to FItems.Count - 1 do
  begin
    Result := TListItem(FItems[n]);
    if Result.Data = AData 
    then begin
      FCacheIndex := n;
      FCacheItem := Result;
      Exit;
    end;
  end;
  Result := nil;
end;

function TListItems.FindData(StartIndex: Integer; Value: Pointer; Inclusive,
  Wrap: Boolean): TListItem;
var
  AnItem: TListItem;
  i: Integer;
begin
  Result := nil;
  if Inclusive then Dec(StartIndex);
  for i := StartIndex + 1 to Count - 1 do
  begin
    AnItem := Item[i];
    if (AnItem <> nil) and (AnItem.Data = Value) then
    begin
      Result := AnItem;
      exit;
    end;
  end;
  if Wrap then
  begin
    if Inclusive then Inc(StartIndex);
    for i := 0 to StartIndex - 1 do
    begin
      AnItem := Item[i];
      if (AnItem <> nil) and (AnItem.Data = Value) then
      begin
        Result := AnItem;
        exit;
      end;
    end;
  end;
end;

function TListItems.GetEnumerator: TListItemsEnumerator;
begin
  Result := TListItemsEnumerator.Create(Self);
end;

{------------------------------------------------------------------------------}
{   TListItems DefineProperties                                                }
{------------------------------------------------------------------------------}
procedure TListItems.DefineProperties(Filer: TFiler);

  function WriteItems: Boolean;
  var
    I: Integer;
    Items: TListItems;
  begin
    Items := TListItems(Filer.Ancestor);
    if not Assigned(Items) then
      Result := Count > 0
    else if (Items.Count <> Count) then
             Result := True
    else
    begin
      Result := False;
      for I := 0 to Count - 1 do
      begin
        Result := not Item[I].IsEqual(Items[I]);
        if Result then Break;
      end
    end;
  end;

begin
  inherited DefineProperties(Filer);
  Filer.DefineBinaryProperty('Data', @ReadData, nil, false);
  // Delphi 10.2.1
  Filer.DefineBinaryProperty('ItemData', @ReadItemData, nil, False);
  Filer.DefineBinaryProperty('LazData', @ReadLazData, @WriteLazData, WriteItems);
end;

type
  TItemHeader = record // packing is not needed (and not wanted since it controls also how this record is stored)
    Size, Count: Integer;
    Items: record end;
  end;
  // for reading Delphi compatible TListItem.Data
  TItemInfo = record // packing is not needed (and not wanted since it controls also how this record is stored)
    ImageIndex: Integer;
    StateIndex: Integer;
    OverlayIndex: Integer;
    SubItemCount: Integer;
    Data: Integer; // pointer field on 32 bits computers
  end;
  // Lazarus does not store ListItem.Data field, because a pointer field cannot
  // be made 64 bits safe
  TLazItemInfo = record // packing is not needed and not wanted
    ImageIndex: Integer;
    StateIndex: Integer;
    OverlayIndex: Integer;
    SubItemCount: Integer;
  end;

procedure TListItems.ReadData(Stream: TStream);
  function ReadString: String;
  var
    Len: Byte;
  begin
    Len := Stream.ReadByte;
    SetLength(Result, Len);
    Stream.ReadBuffer(Result[1], Len);
  end;

var
  I, J: Integer;
  ItemInfo: TItemInfo;
  ListItem: TListItem;
  Size, ItemCount, SubCount: Integer;
  StartPos: Int64;
begin
  Clear;
  StartPos := Stream.Position;
  Size := LEtoN(Integer(Stream.ReadDWord));
  ItemCount := LEtoN(Integer(Stream.ReadDWord));
  Owner.BeginUpdate;
  try
    for I := 0 to ItemCount - 1 do
    begin
      Stream.ReadBuffer(ItemInfo, SizeOf(ItemInfo));
      ListItem := Add;
      ListItem.Caption := ReadString;
      ListItem.ImageIndex := LEtoN(ItemInfo.ImageIndex);
      // ListItem.StateIndex := LEtoN(ItemInfo.StateIndex);
      // ListItem.OverlayIndex := LEtoN(ItemInfo.OverlayIndex);

      //TODO: check if we need to stream a data pointer
      ListItem.Data := Pointer(LEtoN(PtrInt(ItemInfo.Data)));
      
      SubCount := LEtoN(ItemInfo.SubItemCount);
      for J := 0 to SubCount - 1 do
      begin
        ListItem.SubItems.Add(ReadString);
      end;
    end;
    
    //read subitem images
    if Stream.Position < StartPos + Size
    then begin
      for I := 0 to Count - 1 do
      begin
        ListItem := Item[I];
        if ListItem.FSubItems = nil then Continue;

        for J := 0 to ListItem.SubItems.Count - 1 do
          ListItem.SubItemImages[J] := LEtoN(Integer(Stream.ReadDWord));
      end;
    end;
  finally
    Owner.EndUpdate;
  end;
end;


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

procedure TListItems.ReadLazData(Stream: TStream);
var
  I, J: Integer;
  ItemInfo: TLazItemInfo;
  ListItem: TListItem;
  Size, ItemCount, SubCount: Integer;
  StartPos: Int64;

  {$ifdef FPC_BIG_ENDIAN}

    { This is a local redefinition of TStream.ReadAnsiString(), which ensures
      that the string length is read in little-endian order to conform to the
      convention of the remaining resources.
    }
    function stream_readAnsiStringLE(): AnsiString;

    Var
      TheSize : Longint;
      P : PByte ;
    begin
      Stream.ReadBuffer (TheSize,SizeOf(TheSize));
      TheSize := LEtoN(TheSize);
      SetLength(Result,TheSize);
      // Illegal typecast if no AnsiStrings defined.
      if TheSize>0 then
       begin
         Stream.ReadBuffer (Pointer(Result)^,TheSize);
         P:=Pointer(Result)+TheSize;
         p^:=0;
       end
    end { stream_readAnsiStringLE } ;

  {$endif FPC_BIG_ENDIAN}

begin
  Clear;
  StartPos := Stream.Position;
  Size := LEtoN(Integer(Stream.ReadDWord));
  ItemCount := LEtoN(Integer(Stream.ReadDWord));
  Owner.BeginUpdate;
  try
    for I := 0 to ItemCount - 1 do
    begin
      Stream.ReadBuffer(ItemInfo, SizeOf(ItemInfo));
      ListItem := Add;
      {$ifndef FPC_BIG_ENDIAN}
        ListItem.Caption := Stream.ReadAnsiString();
      {$else}
        ListItem.Caption := stream_readAnsiStringLE;
      {$endif FPC_BIG_ENDIAN}
      ListItem.ImageIndex := LEtoN(ItemInfo.ImageIndex);
      // ListItem.StateIndex := LEtoN(ItemInfo.StateIndex);
      // ListItem.OverlayIndex := LEtoN(ItemInfo.OverlayIndex);

      SubCount := LEtoN(ItemInfo.SubItemCount);
      for J := 0 to SubCount - 1 do
      begin
        {$ifndef FPC_BIG_ENDIAN}
          ListItem.SubItems.Add(Stream.ReadAnsiString);
        {$else}
          ListItem.SubItems.Add(stream_readAnsiStringLE())
        {$endif FPC_BIG_ENDIAN}
      end;
    end;

    //read subitem images
    if Stream.Position < StartPos + Size
    then begin
      for I := 0 to Count - 1 do
      begin
        ListItem := Item[I];
        if ListItem.FSubItems = nil then Continue;

        for J := 0 to ListItem.SubItems.Count - 1 do
          ListItem.SubItemImages[J] := LEtoN(Integer(Stream.ReadDWord));
      end;
    end;
  finally
    Owner.EndUpdate;
  end;
end;

procedure TListItems.WriteLazData(Stream: TStream);
var
  I, J, Size, L : Integer;
  ItemHeader    : TItemHeader;
  ItemInfo      : TLazItemInfo;
  ListItem      : TListItem;

  {$ifdef FPC_BIG_ENDIAN}

    { This is a local redefinition of TStream.WriteAnsiString(), which ensures
      that the string length is written in little-endian order to conform to the
      convention of the remaining resources.
    }
    procedure stream_writeAnsiStringLE(const S: AnsiString);

    Var L : Longint;

    begin
      L:=NtoLE(Length(S));
      Stream.WriteBuffer (L,SizeOf(L));
      Stream.WriteBuffer (Pointer(S)^,L)
    end { stream_writeAnsiStringLE } ;

  {$endif FPC_BIG_ENDIAN}

begin
  Size := SizeOf(ItemHeader);
  for I := 0 to Count - 1 do
  begin
    L := Length(Item[I].Caption) + 4;
    for J := 0 to Item[I].SubItems.Count - 1 do
    begin
      Inc(L, Length(Item[I].SubItems[J]) + 4);
      Inc(L, SizeOf(DWORD));
    end;
    Inc(Size, SizeOf(TLazItemInfo) + L);
  end;

  ItemHeader.Size := NtoLE(Size);
  ItemHeader.Count := NtoLE(Count);
  Stream.WriteBuffer(ItemHeader, SizeOf(ItemHeader));
  
  for I := 0 to Count - 1 do
  begin
    ListItem := Item[I];

    ItemInfo.ImageIndex := NtoLE(ListItem.ImageIndex);
    ItemInfo.StateIndex := NtoLE(Integer(-1)) {StateIndex};
    ItemInfo.OverlayIndex := NtoLE(Integer(-1)) {OverlayIndex};
    // don't acces SubItems directly, they will be created
    if ListItem.FSubItems = nil
    then ItemInfo.SubItemCount := 0
    else ItemInfo.SubItemCount := NtoLE(ListItem.SubItems.Count);

    Stream.WriteBuffer(ItemInfo, SizeOf(ItemInfo));

    // Write the strings
    {$ifndef FPC_BIG_ENDIAN}
      Stream.WriteAnsiString(ListItem.Caption);
    {$else}
      stream_writeAnsiStringLE(ListItem.Caption);
    {$endif FPC_BIG_ENDIAN}

    for J := 0 to ItemInfo.SubItemCount - 1 do
    begin
      {$ifndef FPC_BIG_ENDIAN}
        Stream.WriteAnsiString(ListItem.SubItems[J]);
      {$else}
        stream_writeAnsiStringLE(ListItem.SubItems[J])
      {$endif FPC_BIG_ENDIAN}
    end;
  end;
  
  //write SubItem images.
  for I := 0 to Count - 1 do
  begin
    ListItem := Item[I];
    // do not force subitem creation
    if ListItem.FSubItems = nil then Continue;
    for J := 0 to  ListItem.SubItems.Count - 1 do
    begin
      Stream.WriteDWord(DWord(ListItem.SubItemImages[J]));
    end;
  end;
end;

{ TOwnerDataListItems }

function TOwnerDataListItems.GetCount: Integer;
begin
  Result:=fItemsCount;
end;

procedure TOwnerDataListItems.SetCount(const ACount: Integer);
begin
  if (ACount<0) or (ACount=fItemsCount) then Exit;
  fItemsCount:=ACount;
  if WSUpdateAllowed then WSSetItemsCount(fItemsCount);
  // reset ownerdata
  if (FOwner.FOwnerDataItem.Index >= fItemsCount) then
    FOwner.FOwnerDataItem.SetDataIndex(-1);
  // invalidate selection
  FOwner.InvalidateSelected;
end;

function TOwnerDataListItems.GetItem(const AIndex: Integer): TListItem;
begin
  if (AIndex >= FItemsCount) then
  begin
    Result := nil;
    Exit;
  end;
  FOwner.FOwnerDataItem.SetDataIndex(AIndex);
  Result := FOwner.FOwnerDataItem;
end;

procedure TOwnerDataListItems.Clear;
var
  i : Integer;
begin
  for i := 0 to FItems.Count - 1 do
  begin
    TListItem(FItems[i]).Free;
    FItems[i]:=nil;
  end;
  Count := 0;
  fItemsCount := 0;
end;

