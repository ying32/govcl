
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function ListItem_Create(AOwner: TListItems): TListItem; extdecl;
begin
  Result :=  TListItem.Create(AOwner);
end;

procedure ListItem_Free(AObj: TListItem); extdecl;
begin
  AObj.Free;
end;

procedure ListItem_Assign(AObj: TListItem; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

procedure ListItem_Delete(AObj: TListItem); extdecl;
begin
  AObj.Delete;
end;

procedure ListItem_DisplayRect(AObj: TListItem; Code: TDisplayCode; var Result: TRect); extdecl;
begin
  Result :=  AObj.DisplayRect(Code);
end;

function ListItem_EditCaption(AObj: TListItem): LongBool; extdecl;
begin
  Result :=  AObj.EditCaption;
end;

procedure ListItem_MakeVisible(AObj: TListItem; PartialOK: LongBool); extdecl;
begin
  AObj.MakeVisible(PartialOK);
end;

function ListItem_GetNamePath(AObj: TListItem): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function ListItem_ClassName(AObj: TListItem): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function ListItem_Equals(AObj: TListItem; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function ListItem_GetHashCode(AObj: TListItem): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function ListItem_ToString(AObj: TListItem): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function ListItem_GetCaption(AObj: TListItem): PChar; extdecl;
begin
  Result :=  PChar(AObj.Caption);
end;

procedure ListItem_SetCaption(AObj: TListItem; AValue: PChar); extdecl;
begin
  AObj.Caption := AValue;
end;

function ListItem_GetChecked(AObj: TListItem): LongBool; extdecl;
begin
  Result :=  AObj.Checked;
end;

procedure ListItem_SetChecked(AObj: TListItem; AValue: LongBool); extdecl;
begin
  AObj.Checked := AValue;
end;

function ListItem_GetCut(AObj: TListItem): LongBool; extdecl;
begin
  Result :=  AObj.Cut;
end;

procedure ListItem_SetCut(AObj: TListItem; AValue: LongBool); extdecl;
begin
  AObj.Cut := AValue;
end;

function ListItem_GetData(AObj: TListItem): Pointer; extdecl;
begin
  Result :=  AObj.Data;
end;

procedure ListItem_SetData(AObj: TListItem; AValue: Pointer); extdecl;
begin
  AObj.Data := AValue;
end;

function ListItem_GetFocused(AObj: TListItem): LongBool; extdecl;
begin
  Result :=  AObj.Focused;
end;

procedure ListItem_SetFocused(AObj: TListItem; AValue: LongBool); extdecl;
begin
  AObj.Focused := AValue;
end;

function ListItem_GetImageIndex(AObj: TListItem): TImageIndex; extdecl;
begin
  Result :=  AObj.ImageIndex;
end;

procedure ListItem_SetImageIndex(AObj: TListItem; AValue: TImageIndex); extdecl;
begin
  AObj.ImageIndex := AValue;
end;

function ListItem_GetIndex(AObj: TListItem): Integer; extdecl;
begin
  Result :=  AObj.Index;
end;

function ListItem_GetLeft(AObj: TListItem): Integer; extdecl;
begin
  Result :=  AObj.Left;
end;

procedure ListItem_SetLeft(AObj: TListItem; AValue: Integer); extdecl;
begin
  AObj.Left := AValue;
end;

function ListItem_GetOwner(AObj: TListItem): TListItems; extdecl;
begin
  Result :=  AObj.Owner;
end;

procedure ListItem_GetPosition(AObj: TListItem; var Result: TPoint); extdecl;
begin
  Result :=  AObj.Position;
end;

procedure ListItem_SetPosition(AObj: TListItem; var AValue: TPoint); extdecl;
begin
  AObj.Position := AValue;
end;

function ListItem_GetSelected(AObj: TListItem): LongBool; extdecl;
begin
  Result :=  AObj.Selected;
end;

procedure ListItem_SetSelected(AObj: TListItem; AValue: LongBool); extdecl;
begin
  AObj.Selected := AValue;
end;

function ListItem_GetStateIndex(AObj: TListItem): TImageIndex; extdecl;
begin
  Result :=  AObj.StateIndex;
end;

procedure ListItem_SetStateIndex(AObj: TListItem; AValue: TImageIndex); extdecl;
begin
  AObj.StateIndex := AValue;
end;

function ListItem_GetSubItems(AObj: TListItem): TStrings; extdecl;
begin
  Result :=  AObj.SubItems;
end;

procedure ListItem_SetSubItems(AObj: TListItem; AValue: TStrings); extdecl;
begin
  AObj.SubItems := AValue;
end;

function ListItem_GetTop(AObj: TListItem): Integer; extdecl;
begin
  Result :=  AObj.Top;
end;

procedure ListItem_SetTop(AObj: TListItem; AValue: Integer); extdecl;
begin
  AObj.Top := AValue;
end;

function ListItem_GetSubItemImages(AObj: TListItem; Index: Integer): Integer; extdecl;
begin
  Result :=  AObj.SubItemImages[Index];
end;

procedure ListItem_SetSubItemImages(AObj: TListItem; Index: Integer; AValue: Integer); extdecl;
begin
  AObj.SubItemImages[Index] := AValue;
end;


exports
  ListItem_Create,
  ListItem_Free,
  ListItem_Assign,
  ListItem_Delete,
  ListItem_DisplayRect,
  ListItem_EditCaption,
  ListItem_MakeVisible,
  ListItem_GetNamePath,
  ListItem_ClassName,
  ListItem_Equals,
  ListItem_GetHashCode,
  ListItem_ToString,
  ListItem_GetCaption,
  ListItem_SetCaption,
  ListItem_GetChecked,
  ListItem_SetChecked,
  ListItem_GetCut,
  ListItem_SetCut,
  ListItem_GetData,
  ListItem_SetData,
  ListItem_GetFocused,
  ListItem_SetFocused,
  ListItem_GetImageIndex,
  ListItem_SetImageIndex,
  ListItem_GetIndex,
  ListItem_GetLeft,
  ListItem_SetLeft,
  ListItem_GetOwner,
  ListItem_GetPosition,
  ListItem_SetPosition,
  ListItem_GetSelected,
  ListItem_SetSelected,
  ListItem_GetStateIndex,
  ListItem_SetStateIndex,
  ListItem_GetSubItems,
  ListItem_SetSubItems,
  ListItem_GetTop,
  ListItem_SetTop,
  ListItem_GetSubItemImages,
  ListItem_SetSubItemImages;

