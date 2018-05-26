
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function ListGroup_Create(AOwner: TCollection): TListGroup; stdcall;
begin
  Result :=  TListGroup.Create(AOwner);
end;

procedure ListGroup_Free(AObj: TListGroup); stdcall;
begin
  AObj.Free;
end;

procedure ListGroup_Assign(AObj: TListGroup; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function ListGroup_GetNamePath(AObj: TListGroup): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function ListGroup_ClassName(AObj: TListGroup): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function ListGroup_Equals(AObj: TListGroup; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function ListGroup_GetHashCode(AObj: TListGroup): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function ListGroup_ToString(AObj: TListGroup): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function ListGroup_GetHeader(AObj: TListGroup): PChar; stdcall;
begin
  Result :=  PChar(AObj.Header);
end;

procedure ListGroup_SetHeader(AObj: TListGroup; AValue: PChar); stdcall;
begin
  AObj.Header := AValue;
end;

function ListGroup_GetFooter(AObj: TListGroup): PChar; stdcall;
begin
  Result :=  PChar(AObj.Footer);
end;

procedure ListGroup_SetFooter(AObj: TListGroup; AValue: PChar); stdcall;
begin
  AObj.Footer := AValue;
end;

function ListGroup_GetGroupID(AObj: TListGroup): Integer; stdcall;
begin
  Result :=  AObj.GroupID;
end;

procedure ListGroup_SetGroupID(AObj: TListGroup; AValue: Integer); stdcall;
begin
  AObj.GroupID := AValue;
end;

function ListGroup_GetState(AObj: TListGroup): TListGroupStateSet; stdcall;
begin
  Result :=  AObj.State;
end;

procedure ListGroup_SetState(AObj: TListGroup; AValue: TListGroupStateSet); stdcall;
begin
  AObj.State := AValue;
end;

function ListGroup_GetHeaderAlign(AObj: TListGroup): TAlignment; stdcall;
begin
  Result :=  AObj.HeaderAlign;
end;

procedure ListGroup_SetHeaderAlign(AObj: TListGroup; AValue: TAlignment); stdcall;
begin
  AObj.HeaderAlign := AValue;
end;

function ListGroup_GetFooterAlign(AObj: TListGroup): TAlignment; stdcall;
begin
  Result :=  AObj.FooterAlign;
end;

procedure ListGroup_SetFooterAlign(AObj: TListGroup; AValue: TAlignment); stdcall;
begin
  AObj.FooterAlign := AValue;
end;

function ListGroup_GetSubtitle(AObj: TListGroup): PChar; stdcall;
begin
  Result :=  PChar(AObj.Subtitle);
end;

procedure ListGroup_SetSubtitle(AObj: TListGroup; AValue: PChar); stdcall;
begin
  AObj.Subtitle := AValue;
end;

function ListGroup_GetTitleImage(AObj: TListGroup): TImageIndex; stdcall;
begin
  Result :=  AObj.TitleImage;
end;

procedure ListGroup_SetTitleImage(AObj: TListGroup; AValue: TImageIndex); stdcall;
begin
  AObj.TitleImage := AValue;
end;

function ListGroup_GetIndex(AObj: TListGroup): Integer; stdcall;
begin
  Result :=  AObj.Index;
end;

procedure ListGroup_SetIndex(AObj: TListGroup; AValue: Integer); stdcall;
begin
  AObj.Index := AValue;
end;


exports
  ListGroup_Create,
  ListGroup_Free,
  ListGroup_Assign,
  ListGroup_GetNamePath,
  ListGroup_ClassName,
  ListGroup_Equals,
  ListGroup_GetHashCode,
  ListGroup_ToString,
  ListGroup_GetHeader,
  ListGroup_SetHeader,
  ListGroup_GetFooter,
  ListGroup_SetFooter,
  ListGroup_GetGroupID,
  ListGroup_SetGroupID,
  ListGroup_GetState,
  ListGroup_SetState,
  ListGroup_GetHeaderAlign,
  ListGroup_SetHeaderAlign,
  ListGroup_GetFooterAlign,
  ListGroup_SetFooterAlign,
  ListGroup_GetSubtitle,
  ListGroup_SetSubtitle,
  ListGroup_GetTitleImage,
  ListGroup_SetTitleImage,
  ListGroup_GetIndex,
  ListGroup_SetIndex;

