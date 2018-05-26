
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function ImageList_Create(AOwner: TComponent): TImageList; extdecl;
begin
  Result :=  TImageList.Create(AOwner);
end;

procedure ImageList_Free(AObj: TImageList); extdecl;
begin
  AObj.Free;
end;

procedure ImageList_GetHotSpot(AObj: TImageList; var Result: TPoint); extdecl;
begin
  Result :=  AObj.GetHotSpot;
end;

procedure ImageList_Assign(AObj: TImageList; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function ImageList_Add(AObj: TImageList; Image: TBitmap; Mask: TBitmap): Integer; extdecl;
begin
  Result :=  AObj.Add(Image, Mask);
end;

function ImageList_AddIcon(AObj: TImageList; Image: TIcon): Integer; extdecl;
begin
  Result :=  AObj.AddIcon(Image);
end;

procedure ImageList_AddImages(AObj: TImageList; Value: TCustomImageList); extdecl;
begin
  AObj.AddImages(Value);
end;

function ImageList_AddMasked(AObj: TImageList; Image: TBitmap; MaskColor: TColor): Integer; extdecl;
begin
  Result :=  AObj.AddMasked(Image, MaskColor);
end;

procedure ImageList_Clear(AObj: TImageList); extdecl;
begin
  AObj.Clear;
end;

procedure ImageList_Delete(AObj: TImageList; Index: Integer); extdecl;
begin
  AObj.Delete(Index);
end;

procedure ImageList_Insert(AObj: TImageList; Index: Integer; Image: TBitmap; Mask: TBitmap); extdecl;
begin
  AObj.Insert(Index, Image, Mask);
end;

procedure ImageList_InsertIcon(AObj: TImageList; Index: Integer; Image: TIcon); extdecl;
begin
  AObj.InsertIcon(Index, Image);
end;

procedure ImageList_InsertMasked(AObj: TImageList; Index: Integer; Image: TBitmap; MaskColor: TColor); extdecl;
begin
  AObj.InsertMasked(Index, Image, MaskColor);
end;

procedure ImageList_Move(AObj: TImageList; CurIndex: Integer; NewIndex: Integer); extdecl;
begin
  AObj.Move(CurIndex, NewIndex);
end;

procedure ImageList_Replace(AObj: TImageList; Index: Integer; Image: TBitmap; Mask: TBitmap); extdecl;
begin
  AObj.Replace(Index, Image, Mask);
end;

procedure ImageList_ReplaceMasked(AObj: TImageList; Index: Integer; NewImage: TBitmap; MaskColor: TColor); extdecl;
begin
  AObj.ReplaceMasked(Index, NewImage, MaskColor);
end;

procedure ImageList_BeginUpdate(AObj: TImageList); extdecl;
begin
  AObj.BeginUpdate;
end;

procedure ImageList_EndUpdate(AObj: TImageList); extdecl;
begin
  AObj.EndUpdate;
end;

function ImageList_FindComponent(AObj: TImageList; AName: PChar): TComponent; extdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function ImageList_GetNamePath(AObj: TImageList): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function ImageList_HasParent(AObj: TImageList): LongBool; extdecl;
begin
  Result :=  AObj.HasParent;
end;

function ImageList_ClassName(AObj: TImageList): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function ImageList_Equals(AObj: TImageList; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function ImageList_GetHashCode(AObj: TImageList): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function ImageList_ToString(AObj: TImageList): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function ImageList_GetBlendColor(AObj: TImageList): TColor; extdecl;
begin
  Result :=  AObj.BlendColor;
end;

procedure ImageList_SetBlendColor(AObj: TImageList; AValue: TColor); extdecl;
begin
  AObj.BlendColor := AValue;
end;

function ImageList_GetBkColor(AObj: TImageList): TColor; extdecl;
begin
  Result :=  AObj.BkColor;
end;

procedure ImageList_SetBkColor(AObj: TImageList; AValue: TColor); extdecl;
begin
  AObj.BkColor := AValue;
end;

function ImageList_GetAllocBy(AObj: TImageList): Integer; extdecl;
begin
  Result :=  AObj.AllocBy;
end;

procedure ImageList_SetAllocBy(AObj: TImageList; AValue: Integer); extdecl;
begin
  AObj.AllocBy := AValue;
end;

function ImageList_GetDrawingStyle(AObj: TImageList): TDrawingStyle; extdecl;
begin
  Result :=  AObj.DrawingStyle;
end;

procedure ImageList_SetDrawingStyle(AObj: TImageList; AValue: TDrawingStyle); extdecl;
begin
  AObj.DrawingStyle := AValue;
end;

function ImageList_GetHeight(AObj: TImageList): Integer; extdecl;
begin
  Result :=  AObj.Height;
end;

procedure ImageList_SetHeight(AObj: TImageList; AValue: Integer); extdecl;
begin
  AObj.Height := AValue;
end;

function ImageList_GetImageType(AObj: TImageList): TImageType; extdecl;
begin
  Result :=  AObj.ImageType;
end;

procedure ImageList_SetImageType(AObj: TImageList; AValue: TImageType); extdecl;
begin
  AObj.ImageType := AValue;
end;

function ImageList_GetMasked(AObj: TImageList): LongBool; extdecl;
begin
  Result :=  AObj.Masked;
end;

procedure ImageList_SetMasked(AObj: TImageList; AValue: LongBool); extdecl;
begin
  AObj.Masked := AValue;
end;

procedure ImageList_SetOnChange(AObj: TImageList; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnChange := nil;
    EventClass.Remove(AObj, geChange);
    Exit;
  end;
  AObj.OnChange := @EventClass.OnChange;
  EventClass.Add(AObj, geChange, AEventId);
end;

function ImageList_GetShareImages(AObj: TImageList): LongBool; extdecl;
begin
  Result :=  AObj.ShareImages;
end;

procedure ImageList_SetShareImages(AObj: TImageList; AValue: LongBool); extdecl;
begin
  AObj.ShareImages := AValue;
end;

function ImageList_GetWidth(AObj: TImageList): Integer; extdecl;
begin
  Result :=  AObj.Width;
end;

procedure ImageList_SetWidth(AObj: TImageList; AValue: Integer); extdecl;
begin
  AObj.Width := AValue;
end;

function ImageList_GetCount(AObj: TImageList): Integer; extdecl;
begin
  Result :=  AObj.Count;
end;

function ImageList_GetComponentCount(AObj: TImageList): Integer; extdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function ImageList_GetComponentIndex(AObj: TImageList): Integer; extdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure ImageList_SetComponentIndex(AObj: TImageList; AValue: Integer); extdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function ImageList_GetOwner(AObj: TImageList): TComponent; extdecl;
begin
  Result :=  AObj.Owner;
end;

function ImageList_GetName(AObj: TImageList): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure ImageList_SetName(AObj: TImageList; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function ImageList_GetTag(AObj: TImageList): NativeInt; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure ImageList_SetTag(AObj: TImageList; AValue: NativeInt); extdecl;
begin
  AObj.Tag := AValue;
end;

function ImageList_GetComponents(AObj: TImageList; AIndex: Integer): TComponent; extdecl;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  ImageList_Create,
  ImageList_Free,
  ImageList_GetHotSpot,
  ImageList_Assign,
  ImageList_Add,
  ImageList_AddIcon,
  ImageList_AddImages,
  ImageList_AddMasked,
  ImageList_Clear,
  ImageList_Delete,
  ImageList_Insert,
  ImageList_InsertIcon,
  ImageList_InsertMasked,
  ImageList_Move,
  ImageList_Replace,
  ImageList_ReplaceMasked,
  ImageList_BeginUpdate,
  ImageList_EndUpdate,
  ImageList_FindComponent,
  ImageList_GetNamePath,
  ImageList_HasParent,
  ImageList_ClassName,
  ImageList_Equals,
  ImageList_GetHashCode,
  ImageList_ToString,
  ImageList_GetBlendColor,
  ImageList_SetBlendColor,
  ImageList_GetBkColor,
  ImageList_SetBkColor,
  ImageList_GetAllocBy,
  ImageList_SetAllocBy,
  ImageList_GetDrawingStyle,
  ImageList_SetDrawingStyle,
  ImageList_GetHeight,
  ImageList_SetHeight,
  ImageList_GetImageType,
  ImageList_SetImageType,
  ImageList_GetMasked,
  ImageList_SetMasked,
  ImageList_SetOnChange,
  ImageList_GetShareImages,
  ImageList_SetShareImages,
  ImageList_GetWidth,
  ImageList_SetWidth,
  ImageList_GetCount,
  ImageList_GetComponentCount,
  ImageList_GetComponentIndex,
  ImageList_SetComponentIndex,
  ImageList_GetOwner,
  ImageList_GetName,
  ImageList_SetName,
  ImageList_GetTag,
  ImageList_SetTag,
  ImageList_GetComponents;

