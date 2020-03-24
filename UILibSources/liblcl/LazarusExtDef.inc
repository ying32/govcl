// 枚举都按4字节对齐
{$Z4+}

type
  // 重定定义
  TUnixDateTime = type LongInt; // int64改为LongInt。只精确到秒
  TPngImage = type TPortableNetworkGraphic;
  TMonthCalendar = type TCalendar;
  TJPEGPixelFormat = TPixelFormat;
  TMargins = TControlBorderSpacing;
  //TPageMeasureUnits = type TMeasureUnits;
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
  {$IFDEF WINDOWS}
    // 重写这个方法，因为要拦截一些东西
    procedure MyRunLoop;
    procedure MyHandleMessage;
    function IsKeyMsg(var Msg: TMsg): Boolean;
  {$ENDIF}
  end;

procedure TApplicationHelper.RestoreTopMosts;
begin
  Self.RestoreStayOnTop(False);
end;

{$IFDEF WINDOWS}
function TApplicationHelper.IsKeyMsg(var Msg: TMsg): Boolean;

  function IsLCLWebviewControl(AhWnd: HWND; out AWeb: TMiniWebview): Boolean;
  var
    LCtl: TWinControl;
  begin
    Result := False;
    LCtl := FindControl(AhWnd);
    if Assigned(LCtl) then
       Result := LCtl is TMiniWebview;
    if Result then
       AWeb := LCtl as TMiniWebview;
  end;


var
  LWnd: HWND;
  LWebview: TMiniWebview;
  LMessage: TMessage;
begin
  Result := False;
  // 这里只拦截key的消息，因为只用于IE浏览器
  if (Msg.Message >= WM_KEYFIRST) and (Msg.Message <= WM_KEYLAST) then
  begin
     LWnd := GetCapture;
     if LWnd <> 0 then
       Exit;
     if not IsLCLWebviewControl(Msg.hwnd, LWebview) then
     begin
       LWnd := Msg.hwnd;
       repeat
         LWnd := GetParent(LWnd);
       until IsLCLWebviewControl(LWnd, LWebview) or (LWnd = 0);
       if LWnd <> 0 then
       begin
         //Windows.PeekMessageW(Msg, 0, WM_KEYFIRST, WM_KEYLAST, PM_REMOVE);
         SendMessageW(LWnd, Msg.message+CN_BASE, Msg.wParam, Msg.lParam);
       end;
     end;
  end;
end;

procedure TApplicationHelper.MyHandleMessage;
var
  LMsg: Windows.TMsg;
  LHandled: Boolean;
begin
  LHandled := True;
  if GWebviewCount > 0 then
  begin
    // 这里只拦截，不移除消息
    if Windows.PeekMessageW(LMsg, 0, WM_KEYFIRST, WM_KEYLAST, PM_NOREMOVE) then
    begin
      // 这里是要另处理的，先不处理吧。
      if IsKeyMsg(LMsg) then
        LHandled := False;
    end;    if not Terminated then
      Idle(True);
  end;
  if LHandled then HandleMessage;
end;

procedure TApplicationHelper.MyRunLoop;
begin
  repeat
    if CaptureExceptions then
      try // run with try..except
        MyHandleMessage;
      except
        HandleException(Self);
      end
    else
      MyHandleMessage; // run without try..except
  until Terminated;
end;
{$ENDIF}



