
//----------------------------------------
// 代码由genvcllib工具自动生成。
// 时间：2017/9/21
// 作者：ying32
//----------------------------------------

function GIFImageList_Create(AOwner: TComponent): TGIFImageList; stdcall;
begin
  Result :=  TGIFImageList.Create(AOwner);
end;

procedure GIFImageList_Free(AObj: TGIFImageList); stdcall;
begin
  AObj.Free;
end;

procedure GIFImageList_LoadFromStream(AObj: TGIFImageList; Stream: TStream); stdcall;
begin
  AObj.LoadFromStream(Stream);
end;

procedure GIFImageList_SaveToStream(AObj: TGIFImageList; Stream: TStream); stdcall;
begin
  AObj.SaveToStream(Stream);
end;

function GIFImageList_Add(AObj: TGIFImageList; Item: TGIFItem): Integer; stdcall;
begin
  Result :=  AObj.Add(Item);
end;

procedure GIFImageList_Clear(AObj: TGIFImageList); stdcall;
begin
  AObj.Clear;
end;

procedure GIFImageList_Delete(AObj: TGIFImageList; Index: Integer); stdcall;
begin
  AObj.Delete(Index);
end;

function GIFImageList_IndexOf(AObj: TGIFImageList; Item: TGIFItem): Integer; stdcall;
begin
  Result :=  AObj.IndexOf(Item);
end;

procedure GIFImageList_Insert(AObj: TGIFImageList; Index: Integer; Item: TGIFItem); stdcall;
begin
  AObj.Insert(Index, Item);
end;

procedure GIFImageList_Move(AObj: TGIFImageList; CurIndex: Integer; NewIndex: Integer); stdcall;
begin
  AObj.Move(CurIndex, NewIndex);
end;

procedure GIFImageList_Assign(AObj: TGIFImageList; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function GIFImageList_GetNamePath(AObj: TGIFImageList): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function GIFImageList_ClassName(AObj: TGIFImageList): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function GIFImageList_Equals(AObj: TGIFImageList; Obj: TObject): Boolean; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function GIFImageList_GetHashCode(AObj: TGIFImageList): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function GIFImageList_ToString(AObj: TGIFImageList): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function GIFImageList_GetList(AObj: TGIFImageList): TList; stdcall;
begin
  Result :=  AObj.List;
end;


exports
  GIFImageList_Create,
  GIFImageList_Free,
  GIFImageList_LoadFromStream,
  GIFImageList_SaveToStream,
  GIFImageList_Add,
  GIFImageList_Clear,
  GIFImageList_Delete,
  GIFImageList_IndexOf,
  GIFImageList_Insert,
  GIFImageList_Move,
  GIFImageList_Assign,
  GIFImageList_GetNamePath,
  GIFImageList_ClassName,
  GIFImageList_Equals,
  GIFImageList_GetHashCode,
  GIFImageList_ToString,
  GIFImageList_GetList;

