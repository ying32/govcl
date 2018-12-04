// 枚举都按4字节对齐
{$Z4+}

type
  // 重定定义
  TUnixDateTime = type LongInt; // int64改为LongInt。只精确到秒
  TPngImage = type TPortableNetworkGraphic;
  TMonthCalendar = type TCalendar;
  TJPEGPixelFormat = TPixelFormat;
  TPageMeasureUnits = type TMeasureUnits;
{$ifndef windows}
  PFNLVCOMPARE = Pointer;
  PFNTVCOMPARE = Pointer;
{$endif}


// 补丁集合


// --------------- TColorBox ----------------------
type

  { TColorBoxHelper }

  TColorBoxHelper = class helper for TColorBox
  public
    procedure DeleteSelected;
  end;


procedure TColorBoxHelper.DeleteSelected;
begin
  if ItemIndex <> -1 then
    Items.Delete(ItemIndex);
end;

// --------------- TListView ------------------
type

  { TListViewHelper }

  TListViewHelper = class helper for TListView
  public
    procedure DeleteSelected;
    function CustomSort(SortProc: PFNLVCOMPARE; lParam: NativeInt): Boolean;
  end;

procedure TListViewHelper.DeleteSelected;
begin
if ItemIndex <> -1 then
  Items.Delete(ItemIndex);
end;

// 兼容版
function TListViewHelper.CustomSort(SortProc: PFNLVCOMPARE; lParam: NativeInt): Boolean;
begin
  SortType := stText;
  SortColumn := lParam;
  Result := True;
end;

// --------------- TTreeView ------------------
type

  { TTreeViewHelper }

  TTreeViewHelper = class helper for TTreeView
  public
    procedure DeleteSelected;
    function CustomSort(SortProc: PFNTVCOMPARE; Data: NativeInt; ARecurse: LongBool): Boolean;
    function AlphaSort(ARecurse: Boolean): Boolean;
  end;

procedure TTreeViewHelper.DeleteSelected;
var
  LNode: TTreeNode;
begin
  LNode := Selected;
  if LNode <> nil then
     LNode.Delete;
end;

function TTreeViewHelper.CustomSort(SortProc: PFNTVCOMPARE; Data: NativeInt; ARecurse: LongBool): Boolean;
begin
  // no support
  Result := True;
end;

function TTreeViewHelper.AlphaSort(ARecurse: Boolean): Boolean;
begin
  Result := inherited AlphaSort();
end;

// --------------- TTreeNodes ------------------
type

  { TTreeNodesHelper }

  TTreeNodesHelper = class helper for TTreeNodes
  public
    function CustomSort(SortProc: PFNTVCOMPARE; Data: NativeInt; ARecurse: Boolean): Boolean;
  end;

function TTreeNodesHelper.CustomSort(SortProc: PFNTVCOMPARE; Data: NativeInt;
  ARecurse: Boolean): Boolean;
begin
  Result := False;
end;


// --------------- TTreeNode ------------------
type

  { TTreeNodeHelper }

  TTreeNodeHelper = class helper for TTreeNode
  public
    function CustomSort(SortProc: PFNTVCOMPARE; Data: NativeInt; ARecurse: Boolean): Boolean;
  end;

function TTreeNodeHelper.CustomSort(SortProc: PFNTVCOMPARE; Data: NativeInt;
  ARecurse: Boolean): Boolean;
begin
  Result := False;// inherited CustomSort(TTreeNodeCompare(SortProc));
end;

// --------------- TComboBox ------------------
type

  { TComboBoxHelper }

  TComboBoxHelper = class helper for TComboBox
  public
    procedure DeleteSelected;
  end;

procedure TComboBoxHelper.DeleteSelected;
begin
if ItemIndex <> -1 then
  Items.Delete(ItemIndex);
end;


// --------------- TList ------------------
type

 TListAssignOp = (laCopy, laAnd, laOr, laXor, laSrcUnique, laDestUnique);
  { TListHelper }

  TListHelper = class helper for TList
  public
    procedure Assign(ListA: TList; AOperator: TListAssignOp = laCopy; ListB: TList = nil);
  end;

procedure TListHelper.Assign(ListA: TList; AOperator: TListAssignOp = laCopy; ListB: TList = nil);
begin
  // no support
end;


// --------------- TMonthCalendar ------------------
type

  { TMonthCalendarHelper }

  TMonthCalendarHelper = class helper for TMonthCalendar
  private
    function GetDate: TDateTime;
    procedure SetDate(AValue: TDateTime);
  public
    property Date: TDateTime read GetDate write SetDate;
  end;

function TMonthCalendarHelper.GetDate: TDateTime;
begin
  Result := DateTime;
end;

procedure TMonthCalendarHelper.SetDate(AValue: TDateTime);
begin
  DateTime := AValue;
end;

// --------------- TDragObject ------------------
type

  { TDragObjectHelper }

  TDragObjectHelper = class helper for TDragObject
  private
    function GetDragTarget: Pointer;
    procedure SetDragTarget(AValue: Pointer);
  public
    procedure Assign(Source: TDragObject);
    property DragTarget: Pointer read GetDragTarget write SetDragTarget;
  end;

function TDragObjectHelper.GetDragTarget: Pointer;
begin
  Result := inherited DragTarget;
end;

procedure TDragObjectHelper.SetDragTarget(AValue: Pointer);
begin
  inherited DragTarget := TControl(AValue);
end;

procedure TDragObjectHelper.Assign(Source: TDragObject);
begin
  // no support
end;


// --------------- TStatusPanels ------------------
type

  { TStatusPanelsHelper }

  TStatusPanelsHelper = class helper for TStatusPanels
  private
    function GetCapacity: Integer;
    procedure SetCapacity(AValue: Integer);
  public
    property Capacity: Integer read GetCapacity write SetCapacity;
  end;

function TStatusPanelsHelper.GetCapacity: Integer;
begin
  Result := 0;
end;

procedure TStatusPanelsHelper.SetCapacity(AValue: Integer);
begin
end;


// --------------- TPageSetupDialog ------------------
type

  TPageSetupDialogOptions = type Cardinal;
  { TPageSetupDialogHelper }

  TPageSetupDialogHelper = class helper for TPageSetupDialog
  private
    function GetOptions: TPageSetupDialogOptions;
    function GetPageHeight: Integer;
    function GetPageWidth: Integer;
    procedure SetOptions(AValue: TPageSetupDialogOptions);
    procedure SetPageHeight(AValue: Integer);
    procedure SetPageWidth(AValue: Integer);
    function GetMarginLeft: Integer;
    procedure SetMarginLeft(AValue: Integer);
    function GetMarginTop: Integer;
    procedure SetMarginTop(AValue: Integer);
    function GetMarginRight: Integer;
    procedure SetMarginRight(AValue: Integer);
    function GetMarginBottom: Integer;
    procedure SetMarginBottom(AValue: Integer);
  public
    property Options: TPageSetupDialogOptions read GetOptions write SetOptions;
    property PageWidth: Integer read GetPageWidth write SetPageWidth;
    property PageHeight: Integer read GetPageHeight write SetPageHeight;
    property MarginLeft: Integer read GetMarginLeft write SetMarginLeft;
    property MarginTop: Integer read GetMarginTop write SetMarginTop;
    property MarginRight: Integer read GetMarginRight write SetMarginRight;
    property MarginBottom: Integer read GetMarginBottom write SetMarginBottom;
  end;

function TPageSetupDialogHelper.GetOptions: TPageSetupDialogOptions;
begin
  Result := 0;
end;

function TPageSetupDialogHelper.GetPageHeight: Integer;
begin
  Result := 0;
end;

function TPageSetupDialogHelper.GetPageWidth: Integer;
begin
  Result := 0;
end;

procedure TPageSetupDialogHelper.SetOptions(AValue: TPageSetupDialogOptions);
begin
end;

procedure TPageSetupDialogHelper.SetPageHeight(AValue: Integer);
begin

end;

procedure TPageSetupDialogHelper.SetPageWidth(AValue: Integer);
begin

end;

function TPageSetupDialogHelper.GetMarginLeft: Integer;
begin
  Result := Margins.Left;
end;

procedure TPageSetupDialogHelper.SetMarginLeft(AValue: Integer);
var
  LR: TRect;
begin
  LR := Margins;
  LR.Left := AValue;
  Margins := LR;
end;

function TPageSetupDialogHelper.GetMarginTop: Integer;
begin
  Result := Margins.Top;
end;

procedure TPageSetupDialogHelper.SetMarginTop(AValue: Integer);
var
  LR: TRect;
begin
  LR := Margins;
  LR.Top := AValue;
  Margins := LR;
end;

function TPageSetupDialogHelper.GetMarginRight: Integer;
begin
  Result := Margins.Right;
end;

procedure TPageSetupDialogHelper.SetMarginRight(AValue: Integer);
var
  LR: TRect;
begin
  LR := Margins;
  LR.Right := AValue;
  Margins := LR;
end;

function TPageSetupDialogHelper.GetMarginBottom: Integer;
begin
  Result :=  Margins.Bottom;
end;

procedure TPageSetupDialogHelper.SetMarginBottom(AValue: Integer);
var
  LR: TRect;
begin
  LR := Margins;
  LR.Bottom := AValue;
  Margins := LR;
end;


// --------------- TPrinter ------------------
//type
//
//  { TPrinterHelper }
//
//  TPrinterHelper = class helper for TPrinter
//  private
//    function GetHandle: HDC;
//  public
//    property Handle: HDC read GetHandle;
//  end;
//
//function TPrinterHelper.GetHandle: HDC;
//begin
//  Result := 0;
//end;

// --------------- TApplication ------------------
type

  { TApplicationHelper }

  TApplicationHelper = class helper for TApplication
  public
    procedure RestoreTopMosts;
  end;

procedure TApplicationHelper.RestoreTopMosts;
begin
  Self.RestoreStayOnTop(False);
end;

