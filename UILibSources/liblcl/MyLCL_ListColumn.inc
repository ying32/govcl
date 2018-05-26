
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function ListColumn_Create(AOwner: TCollection): TListColumn; extdecl;
begin
  Result :=  TListColumn.Create(AOwner);
end;

procedure ListColumn_Free(AObj: TListColumn); extdecl;
begin
  AObj.Free;
end;

procedure ListColumn_Assign(AObj: TListColumn; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function ListColumn_GetNamePath(AObj: TListColumn): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function ListColumn_ClassName(AObj: TListColumn): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function ListColumn_Equals(AObj: TListColumn; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function ListColumn_GetHashCode(AObj: TListColumn): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function ListColumn_ToString(AObj: TListColumn): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function ListColumn_GetAlignment(AObj: TListColumn): TAlignment; extdecl;
begin
  Result :=  AObj.Alignment;
end;

procedure ListColumn_SetAlignment(AObj: TListColumn; AValue: TAlignment); extdecl;
begin
  AObj.Alignment := AValue;
end;

function ListColumn_GetAutoSize(AObj: TListColumn): LongBool; extdecl;
begin
  Result :=  AObj.AutoSize;
end;

procedure ListColumn_SetAutoSize(AObj: TListColumn; AValue: LongBool); extdecl;
begin
  AObj.AutoSize := AValue;
end;

function ListColumn_GetCaption(AObj: TListColumn): PChar; extdecl;
begin
  Result :=  PChar(AObj.Caption);
end;

procedure ListColumn_SetCaption(AObj: TListColumn; AValue: PChar); extdecl;
begin
  AObj.Caption := AValue;
end;

function ListColumn_GetImageIndex(AObj: TListColumn): TImageIndex; extdecl;
begin
  Result :=  AObj.ImageIndex;
end;

procedure ListColumn_SetImageIndex(AObj: TListColumn; AValue: TImageIndex); extdecl;
begin
  AObj.ImageIndex := AValue;
end;

function ListColumn_GetMaxWidth(AObj: TListColumn): TWidth; extdecl;
begin
  Result :=  AObj.MaxWidth;
end;

procedure ListColumn_SetMaxWidth(AObj: TListColumn; AValue: TWidth); extdecl;
begin
  AObj.MaxWidth := AValue;
end;

function ListColumn_GetMinWidth(AObj: TListColumn): TWidth; extdecl;
begin
  Result :=  AObj.MinWidth;
end;

procedure ListColumn_SetMinWidth(AObj: TListColumn; AValue: TWidth); extdecl;
begin
  AObj.MinWidth := AValue;
end;

function ListColumn_GetTag(AObj: TListColumn): Integer; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure ListColumn_SetTag(AObj: TListColumn; AValue: Integer); extdecl;
begin
  AObj.Tag := AValue;
end;

function ListColumn_GetWidth(AObj: TListColumn): TWidth; extdecl;
begin
  Result :=  AObj.Width;
end;

procedure ListColumn_SetWidth(AObj: TListColumn; AValue: TWidth); extdecl;
begin
  AObj.Width := AValue;
end;

function ListColumn_GetIndex(AObj: TListColumn): Integer; extdecl;
begin
  Result :=  AObj.Index;
end;

procedure ListColumn_SetIndex(AObj: TListColumn; AValue: Integer); extdecl;
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

