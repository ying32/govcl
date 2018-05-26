
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function MonthCalColors_Create(AOwner: TComponent): TMonthCalColors; extdecl;
begin
  Result :=  TMonthCalColors.Create(AOwner);
end;

procedure MonthCalColors_Free(AObj: TMonthCalColors); extdecl;
begin
  AObj.Free;
end;

procedure MonthCalColors_Assign(AObj: TMonthCalColors; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function MonthCalColors_GetNamePath(AObj: TMonthCalColors): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function MonthCalColors_Equals(AObj: TMonthCalColors; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function MonthCalColors_GetHashCode(AObj: TMonthCalColors): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function MonthCalColors_ToString(AObj: TMonthCalColors): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;


exports
  MonthCalColors_Create,
  MonthCalColors_Free,
  MonthCalColors_Assign,
  MonthCalColors_GetNamePath,
  MonthCalColors_Equals,
  MonthCalColors_GetHashCode,
  MonthCalColors_ToString;

