
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function ListColumn_Create(AOwner: TCollection): TListColumn; stdcall;
begin
  Result :=  TListColumn.Create(AOwner);
end;

procedure ListColumn_Free(AObj: TListColumn); stdcall;
begin
  AObj.Free;
end;

procedure ListColumn_Assign(AObj: TListColumn; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function ListColumn_GetNamePath(AObj: TListColumn): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function ListColumn_ClassName(AObj: TListColumn): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function ListColumn_Equals(AObj: TListColumn; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function ListColumn_GetHashCode(AObj: TListColumn): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function ListColumn_ToString(AObj: TListColumn): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function ListColumn_GetAlignment(AObj: TListColumn): TAlignment; stdcall;
begin
  Result :=  AObj.Alignment;
end;

procedure ListColumn_SetAlignment(AObj: TListColumn; AValue: TAlignment); stdcall;
begin
  AObj.Alignment := AValue;
end;

function ListColumn_GetAutoSize(AObj: TListColumn): LongBool; stdcall;
begin
  Result :=  AObj.AutoSize;
end;

procedure ListColumn_SetAutoSize(AObj: TListColumn; AValue: LongBool); stdcall;
begin
  AObj.AutoSize := AValue;
end;

function ListColumn_GetCaption(AObj: TListColumn): PChar; stdcall;
begin
  Result :=  PChar(AObj.Caption);
end;

procedure ListColumn_SetCaption(AObj: TListColumn; AValue: PChar); stdcall;
begin
  AObj.Caption := AValue;
end;

function ListColumn_GetImageIndex(AObj: TListColumn): TImageIndex; stdcall;
begin
  Result :=  AObj.ImageIndex;
end;

procedure ListColumn_SetImageIndex(AObj: TListColumn; AValue: TImageIndex); stdcall;
begin
  AObj.ImageIndex := AValue;
end;

function ListColumn_GetMaxWidth(AObj: TListColumn): TWidth; stdcall;
begin
  Result :=  AObj.MaxWidth;
end;

procedure ListColumn_SetMaxWidth(AObj: TListColumn; AValue: TWidth); stdcall;
begin
  AObj.MaxWidth := AValue;
end;

function ListColumn_GetMinWidth(AObj: TListColumn): TWidth; stdcall;
begin
  Result :=  AObj.MinWidth;
end;

procedure ListColumn_SetMinWidth(AObj: TListColumn; AValue: TWidth); stdcall;
begin
  AObj.MinWidth := AValue;
end;

function ListColumn_GetTag(AObj: TListColumn): Integer; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure ListColumn_SetTag(AObj: TListColumn; AValue: Integer); stdcall;
begin
  AObj.Tag := AValue;
end;

function ListColumn_GetWidth(AObj: TListColumn): TWidth; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure ListColumn_SetWidth(AObj: TListColumn; AValue: TWidth); stdcall;
begin
  AObj.Width := AValue;
end;

function ListColumn_GetIndex(AObj: TListColumn): Integer; stdcall;
begin
  Result :=  AObj.Index;
end;

procedure ListColumn_SetIndex(AObj: TListColumn; AValue: Integer); stdcall;
begin
  AObj.Index := AValue;
end;


exports
  ListColumn_Create,
  ListColumn_Free,
  ListColumn_Assign,
  ListColumn_GetNamePath,
  ListColumn_ClassName,
  ListColumn_Equals,
  ListColumn_GetHashCode,
  ListColumn_ToString,
  ListColumn_GetAlignment,
  ListColumn_SetAlignment,
  ListColumn_GetAutoSize,
  ListColumn_SetAutoSize,
  ListColumn_GetCaption,
  ListColumn_SetCaption,
  ListColumn_GetImageIndex,
  ListColumn_SetImageIndex,
  ListColumn_GetMaxWidth,
  ListColumn_SetMaxWidth,
  ListColumn_GetMinWidth,
  ListColumn_SetMinWidth,
  ListColumn_GetTag,
  ListColumn_SetTag,
  ListColumn_GetWidth,
  ListColumn_SetWidth,
  ListColumn_GetIndex,
  ListColumn_SetIndex;

