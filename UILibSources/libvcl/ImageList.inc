
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function ImageList_Create(AOwner: TComponent): TImageList; stdcall;
begin
  Result :=  TImageList.Create(AOwner);
end;

procedure ImageList_Free(AObj: TImageList); stdcall;
begin
  AObj.Free;
end;

procedure ImageList_GetHotSpot(AObj: TImageList; var Result: TPoint); stdcall;
begin
  Result :=  AObj.GetHotSpot;
end;

procedure ImageList_Assign(AObj: TImageList; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function ImageList_Add(AObj: TImageList; Image: TBitmap; Mask: TBitmap): Integer; stdcall;
begin
  Result :=  AObj.Add(Image, Mask);
end;

function ImageList_AddIcon(AObj: TImageList; Image: TIcon): Integer; stdcall;
begin
  Result :=  AObj.AddIcon(Image);
end;

function ImageList_AddImage(AObj: TImageList; Value: TCustomImageList; Index: Integer): Integer; stdcall;
begin
  Result :=  AObj.AddImage(Value, Index);
end;

procedure ImageList_AddImages(AObj: TImageList; Value: TCustomImageList); stdcall;
begin
  AObj.AddImages(Value);
end;

function ImageList_AddMasked(AObj: TImageList; Image: TBitmap; MaskColor: TColor): Integer; stdcall;
begin
  Result :=  AObj.AddMasked(Image, MaskColor);
end;

procedure ImageList_Clear(AObj: TImageList); stdcall;
begin
  AObj.Clear;
end;

procedure ImageList_Delete(AObj: TImageList; Index: Integer); stdcall;
begin
  AObj.Delete(Index);
end;

function ImageList_FileLoad(AObj: TImageList; ResType: TResType; Name: PChar; MaskColor: TColor): LongBool; stdcall;
begin
  Result :=  AObj.FileLoad(ResType, Name, MaskColor);
end;

function ImageList_GetBitmap(AObj: TImageList; Index: Integer; Image: TBitmap): LongBool; stdcall;
begin
  Result :=  AObj.GetBitmap(Index, Image);
end;

function ImageList_GetImageBitmap(AObj: TImageList): HBITMAP; stdcall;
begin
  Result :=  AObj.GetImageBitmap;
end;

function ImageList_GetMaskBitmap(AObj: TImageList): HBITMAP; stdcall;
begin
  Result :=  AObj.GetMaskBitmap;
end;

function ImageList_GetResource(AObj: TImageList; ResType: TResType; Name: PChar; Width: Integer; LoadFlags: TLoadResources; MaskColor: TColor): LongBool; stdcall;
begin
  Result :=  AObj.GetResource(ResType, Name, Width, LoadFlags, MaskColor);
end;

function ImageList_HandleAllocated(AObj: TImageList): LongBool; stdcall;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure ImageList_Insert(AObj: TImageList; Index: Integer; Image: TBitmap; Mask: TBitmap); stdcall;
begin
  AObj.Insert(Index, Image, Mask);
end;

procedure ImageList_InsertIcon(AObj: TImageList; Index: Integer; Image: TIcon); stdcall;
begin
  AObj.InsertIcon(Index, Image);
end;

procedure ImageList_InsertMasked(AObj: TImageList; Index: Integer; Image: TBitmap; MaskColor: TColor); stdcall;
begin
  AObj.InsertMasked(Index, Image, MaskColor);
end;

procedure ImageList_Move(AObj: TImageList; CurIndex: Integer; NewIndex: Integer); stdcall;
begin
  AObj.Move(CurIndex, NewIndex);
end;

function ImageList_Overlay(AObj: TImageList; ImageIndex: Integer; Overlay: TOverlay): LongBool; stdcall;
begin
  Result :=  AObj.Overlay(ImageIndex, Overlay);
end;

function ImageList_ResourceLoad(AObj: TImageList; ResType: TResType; Name: PChar; MaskColor: TColor): LongBool; stdcall;
begin
  Result :=  AObj.ResourceLoad(ResType, Name, MaskColor);
end;

function ImageList_ResInstLoad(AObj: TImageList; Instance: NativeUInt; ResType: TResType; Name: PChar; MaskColor: TColor): LongBool; stdcall;
begin
  Result :=  AObj.ResInstLoad(Instance, ResType, Name, MaskColor);
end;

procedure ImageList_Replace(AObj: TImageList; Index: Integer; Image: TBitmap; Mask: TBitmap); stdcall;
begin
  AObj.Replace(Index, Image, Mask);
end;

procedure ImageList_ReplaceIcon(AObj: TImageList; Index: Integer; Image: TIcon); stdcall;
begin
  AObj.ReplaceIcon(Index, Image);
end;

procedure ImageList_ReplaceMasked(AObj: TImageList; Index: Integer; NewImage: TBitmap; MaskColor: TColor); stdcall;
begin
  AObj.ReplaceMasked(Index, NewImage, MaskColor);
end;

procedure ImageList_SetSize(AObj: TImageList; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetSize(AWidth, AHeight);
end;

procedure ImageList_BeginUpdate(AObj: TImageList); stdcall;
begin
  AObj.BeginUpdate;
end;

procedure ImageList_EndUpdate(AObj: TImageList); stdcall;
begin
  AObj.EndUpdate;
end;

function ImageList_FindComponent(AObj: TImageList; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function ImageList_GetNamePath(AObj: TImageList): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function ImageList_HasParent(AObj: TImageList): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

function ImageList_ClassName(AObj: TImageList): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function ImageList_Equals(AObj: TImageList; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function ImageList_GetHashCode(AObj: TImageList): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function ImageList_ToString(AObj: TImageList): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function ImageList_GetBlendColor(AObj: TImageList): TColor; stdcall;
begin
  Result :=  AObj.BlendColor;
end;

procedure ImageList_SetBlendColor(AObj: TImageList; AValue: TColor); stdcall;
begin
  AObj.BlendColor := AValue;
end;

function ImageList_GetBkColor(AObj: TImageList): TColor; stdcall;
begin
  Result :=  AObj.BkColor;
end;

procedure ImageList_SetBkColor(AObj: TImageList; AValue: TColor); stdcall;
begin
  AObj.BkColor := AValue;
end;

function ImageList_GetAllocBy(AObj: TImageList): Integer; stdcall;
begin
  Result :=  AObj.AllocBy;
end;

procedure ImageList_SetAllocBy(AObj: TImageList; AValue: Integer); stdcall;
begin
  AObj.AllocBy := AValue;
end;

function ImageList_GetColorDepth(AObj: TImageList): TColorDepth; stdcall;
begin
  Result :=  AObj.ColorDepth;
end;

procedure ImageList_SetColorDepth(AObj: TImageList; AValue: TColorDepth); stdcall;
begin
  AObj.ColorDepth := AValue;
end;

function ImageList_GetDrawingStyle(AObj: TImageList): TDrawingStyle; stdcall;
begin
  Result :=  AObj.DrawingStyle;
end;

procedure ImageList_SetDrawingStyle(AObj: TImageList; AValue: TDrawingStyle); stdcall;
begin
  AObj.DrawingStyle := AValue;
end;

function ImageList_GetGrayscaleFactor(AObj: TImageList): Byte; stdcall;
begin
  Result :=  AObj.GrayscaleFactor;
end;

procedure ImageList_SetGrayscaleFactor(AObj: TImageList; AValue: Byte); stdcall;
begin
  AObj.GrayscaleFactor := AValue;
end;

function ImageList_GetHeight(AObj: TImageList): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure ImageList_SetHeight(AObj: TImageList; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function ImageList_GetImageType(AObj: TImageList): TImageType; stdcall;
begin
  Result :=  AObj.ImageType;
end;

procedure ImageList_SetImageType(AObj: TImageList; AValue: TImageType); stdcall;
begin
  AObj.ImageType := AValue;
end;

function ImageList_GetMasked(AObj: TImageList): LongBool; stdcall;
begin
  Result :=  AObj.Masked;
end;

procedure ImageList_SetMasked(AObj: TImageList; AValue: LongBool); stdcall;
begin
  AObj.Masked := AValue;
end;

procedure ImageList_SetOnChange(AObj: TImageList; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnChange := nil;
    TEventClass.Remove(AObj, geChange);
    Exit;
  end;
  AObj.OnChange := TEventClass.OnChange;
  TEventClass.Add(AObj, geChange, AEventId);
end;

function ImageList_GetShareImages(AObj: TImageList): LongBool; stdcall;
begin
  Result :=  AObj.ShareImages;
end;

procedure ImageList_SetShareImages(AObj: TImageList; AValue: LongBool); stdcall;
begin
  AObj.ShareImages := AValue;
end;

function ImageList_GetWidth(AObj: TImageList): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure ImageList_SetWidth(AObj: TImageList; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function ImageList_GetHandle(AObj: TImageList): NativeUInt; stdcall;
begin
  Result :=  AObj.Handle;
end;

procedure ImageList_SetHandle(AObj: TImageList; AValue: NativeUInt); stdcall;
begin
  AObj.Handle := AValue;
end;

function ImageList_GetCount(AObj: TImageList): Integer; stdcall;
begin
  Result :=  AObj.Count;
end;

function ImageList_GetComponentCount(AObj: TImageList): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function ImageList_GetComponentIndex(AObj: TImageList): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure ImageList_SetComponentIndex(AObj: TImageList; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function ImageList_GetOwner(AObj: TImageList): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function ImageList_GetName(AObj: TImageList): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure ImageList_SetName(AObj: TImageList; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function ImageList_GetTag(AObj: TImageList): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure ImageList_SetTag(AObj: TImageList; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function ImageList_GetComponents(AObj: TImageList; AIndex: Integer): TComponent; stdcall;
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
  ImageList_AddImage,
  ImageList_AddImages,
  ImageList_AddMasked,
  ImageList_Clear,
  ImageList_Delete,
  ImageList_FileLoad,
  ImageList_GetBitmap,
  ImageList_GetImageBitmap,
  ImageList_GetMaskBitmap,
  ImageList_GetResource,
  ImageList_HandleAllocated,
  ImageList_Insert,
  ImageList_InsertIcon,
  ImageList_InsertMasked,
  ImageList_Move,
  ImageList_Overlay,
  ImageList_ResourceLoad,
  ImageList_ResInstLoad,
  ImageList_Replace,
  ImageList_ReplaceIcon,
  ImageList_ReplaceMasked,
  ImageList_SetSize,
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
  ImageList_GetColorDepth,
  ImageList_SetColorDepth,
  ImageList_GetDrawingStyle,
  ImageList_SetDrawingStyle,
  ImageList_GetGrayscaleFactor,
  ImageList_SetGrayscaleFactor,
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
  ImageList_GetHandle,
  ImageList_SetHandle,
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

