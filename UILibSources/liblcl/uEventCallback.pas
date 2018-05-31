unit uEventCallback;

{$mode objfpc}{$H+}

{$I ExtDecl.inc}


interface

uses
  Controls,
  Forms,
  ComCtrls,
  Menus,
  ExtCtrls,
  Classes,
  SysUtils,
  StdCtrls,
  LMessages,
  uLinkLabel,
  fgl;

type
  TCallbackPtr = function(f: NativeUInt; args: Pointer; argcout: NativeInt): Pointer; extdecl;

var
  GCallbackPtr: TCallbackPtr;

type
  TGoParam = record
    &Type: Byte;
    Value: Pointer;
  end;
  PGoParam = ^TGoParam;

  TGoEvent = (geClick, geClose, geFormClose, geFormCloseQuery, geChange,
              geUpDownClick, geTreeViewChange, geListViewChange, geDblClick, gePaint,
              geResize, geShow, geMenuChange, geEnter, geExit, gePopup, geBalloonClick,
              geLinkClick, geExecute, geUpdate, geException, geTimer, geMinimize,
              geRestore, geHide, geKeyDown, geKeyPress, geKeyUp, geMouseDown,
              geMouseEnter, geMouseLeave, geMouseMove, geMouseUp, geMouseWheel,
              geListBoxDrawItem, geMenuItemDrawItem, geListViewColumnClick,
              geListViewColumnRightClick, geListViewGetImageIndex, geListViewSelectItem,
              geListViewItemChecked, geTreeViewGetSelectedIndex, geTreeViewGetImageIndex,
              gePageControlGetImageIndex, geListViewCompare, geTreeViewCompare,
              geListViewAdvancedCustomDraw, geListViewAdvancedCustomDrawItem,
              geListViewAdvancedCustomDrawSubItem,
              geTreeViewAdvancedCustomDraw, geTreeViewAdvancedCustomDrawItem,
              geToolBarAdvancedCustomDraw, geToolBarAdvancedCustomDrawButton,
              geHint, geClickCheck, geDropFiles, geDestroy, geFind, geReplace,
              geConstrainedResize, geDeactivate, geActivate,
              geHelp, geShortCut, geContextPopup, geDockDrop, geDragDrop,
              geDragOver, geEndDock, geGetSiteInfo, geMouseWheelDown,
              geMouseWheelUp, geStartDock, geUnDock, geMessage, geEndDrag);

  TEventKey = packed record
    Sender: TObject;
    Event: TGoEvent;
  end;
  PEventKey = ^TEventKey;

  { TEventList }

  TEventList = specialize  TFPGMap<NativeUInt, NativeUInt>;


  { TEventClass }

  TEventClass = class
  private
    FEvents: TEventList;
    FThreadEvtId: NativeUInt;
    procedure SendEvent(Sender: TObject; AEvent: TGoEvent; AArgs: array of const);
  public
    constructor Create;
    destructor Destroy; override;
  public
    procedure OnClick(Sender: TObject);

    procedure FormOnClose(Sender: TObject; var Action: TCloseAction);
    procedure FormOnCloseQuery(Sender: TObject; var CanClose: Boolean);
    procedure FormOnDropFiles(Sender: TObject; const AFileNames: array of string);

    procedure OnClose(Sender: TObject);

    procedure OnChange(Sender: TObject);

    procedure UpDownOnClick(Sender: TObject; Button: TUDBtnType);

    procedure TreeViewOnChange(Sender: TObject; ANode: TTreeNode);
    procedure TreeViewOnGetImageIndex(Sender: TObject; Node: TTreeNode);
    procedure TreeViewOnGetSelectedIndex(Sender: TObject; Node: TTreeNode);
    procedure TreeViewOnCompare(Sender: TObject; Node1, Node2: TTreeNode; var Compare: Integer);
    procedure TreeViewOnAdvancedCustomDraw(Sender: TCustomTreeView;
      const ARect: TRect; Stage: TCustomDrawStage; var DefaultDraw: Boolean);
    procedure TreeViewOnAdvancedCustomDrawItem(Sender: TCustomTreeView;
      Node: TTreeNode; State: TCustomDrawState; Stage: TCustomDrawStage;
      var PaintImages, DefaultDraw: Boolean);


    procedure ListViewOnChange(Sender: TObject; AItem: TListItem; Change: TItemChange);
    procedure ListViewOnColumnClick(Sender: TObject; Column: TListColumn);
    procedure ListViewOnColumnRightClick(Sender: TObject; Column: TListColumn; Point: TPoint);
    procedure ListViewOnGetImageIndex(Sender: TObject; Item: TListItem);
    procedure ListViewOnSelectItem(Sender: TObject; Item: TListItem; Selected: Boolean);
    procedure ListViewOnItemChecked(Sender: TObject; Item: TListItem);
    procedure ListViewOnCompare(Sender: TObject; Item1, Item2: TListItem; Data: Integer; var Compare: Integer);
    procedure ListViewOnAdvancedCustomDraw(Sender: TCustomListView;
       const ARect: TRect; Stage: TCustomDrawStage; var DefaultDraw: Boolean);
    procedure ListViewOnAdvancedCustomDrawItem(Sender: TCustomListView;
      Item: TListItem; State: TCustomDrawState; Stage: TCustomDrawStage;
      var DefaultDraw: Boolean);
    procedure ListViewOnAdvancedCustomDrawSubItem(Sender: TCustomListView;
      Item: TListItem; SubItem: Integer; State: TCustomDrawState;
      Stage: TCustomDrawStage; var DefaultDraw: Boolean);

    procedure PageControlOnGetImageIndex(Sender: TObject; TabIndex: Integer; var ImageIndex: Integer);


    procedure ToolBarOnAdvancedCustomDraw(Sender: TToolBar;
       const ARect: TRect; Stage: TCustomDrawStage; var DefaultDraw: Boolean);
    //procedure ToolBarOnAdvancedCustomDrawButton(Sender: TToolBar;
      //Button: TToolButton; State: TCustomDrawState; Stage: TCustomDrawStage;
      //var Flags: TTBCustomDrawFlags; var DefaultDraw: Boolean);


    procedure OnDblClick(Sender: TObject);

    procedure OnPaint(Sender: TObject);
    procedure OnResize(Sender: TObject);
    procedure OnShow(Sender: TObject);
    procedure OnEnter(Sender: TObject);
    procedure OnExit(Sender: TObject);
    procedure OnPopup(Sender: TObject);
    procedure OnHint(Sender: TObject);
    procedure OnClickCheck(Sender: TObject);

    procedure OnExecute(Sender: TObject);
    procedure OnUpdate(Sender: TObject);

    procedure OnBalloonClick(Sender: TObject);

    procedure OnException(Sender: TObject; E: Exception);
    procedure OnTimer(Sender: TObject);

    procedure OnMinimize(Sender: TObject);
    procedure OnRestore(Sender: TObject);
    procedure OnHide(Sender: TObject);

    procedure OnDestroy(Sender: TObject);
    procedure OnReplace(Sender: TObject);
    procedure OnFind(Sender: TObject);

    procedure OnActivate(Sender: TObject);
    procedure OnDeactivate(Sender: TObject);
    procedure OnConstrainedResize(Sender: TObject; var MinWidth, MinHeight, MaxWidth, MaxHeight: TConstraintSize);


    // new
    function OnHelp(Command: Word; Data: PtrInt; var CallHelp: Boolean): Boolean;
    procedure OnShortCut(var Msg: TLMKey; var Handled: Boolean);
    procedure OnContextPopup(Sender: TObject; MousePos: TPoint; var Handled: Boolean);
    procedure OnDockDrop(Sender: TObject; Source: TDragDockObject; X, Y: Integer);
    procedure OnDragDrop(Sender, Source: TObject; X, Y: Integer);
    procedure OnDragOver(Sender, Source: TObject; X, Y: Integer; State: TDragState; var Accept: Boolean);
    procedure OnEndDock(Sender, Target: TObject; X, Y: Integer);
    procedure OnGetSiteInfo(Sender: TObject; DockClient: TControl; var InfluenceRect: TRect; MousePos: TPoint; var CanDock: Boolean);
    procedure OnMouseWheelDown(Sender: TObject; Shift: TShiftState; MousePos: TPoint; var Handled: Boolean);
    procedure OnMouseWheelUp(Sender: TObject; Shift: TShiftState; MousePos: TPoint; var Handled: Boolean);
    procedure OnStartDock(Sender: TObject; var DragObject: TDragDockObject);
    procedure OnUnDock(Sender: TObject; Client: TControl; NewTarget: TWinControl; var Allow: Boolean);
    procedure OnEndDrag(Sender, Target: TObject; X, Y: Integer);

    procedure OnKeyDown(Sender: TObject; var Key: Word; Shift: TShiftState);
    procedure OnKeyUp(Sender: TObject; var Key: Word; Shift: TShiftState);
    procedure OnKeyPress(Sender: TObject; var Key: Char);

    procedure OnMouseDown(Sender: TObject; Button: TMouseButton;
           Shift: TShiftState; X, Y: Integer);
    procedure OnMouseMove(Sender: TObject; Shift: TShiftState; X, Y: Integer);
    procedure OnMouseUp(Sender: TObject; Button: TMouseButton;
          Shift: TShiftState; X, Y: Integer);
    procedure OnMouseWheel(Sender: TObject; Shift: TShiftState;
          WheelDelta: Integer; MousePos: TPoint; var Handled: Boolean);

    procedure OnMouseEnter(Sender: TObject);
    procedure OnMouseLeave(Sender: TObject);


    procedure ListBoxOnDrawItem(Control: TWinControl; Index: Integer;
          ARect: TRect; State: TOwnerDrawState);

    procedure OnLinkClick(Sender: TObject; const Link: string; LinkType: TSysLinkType);

    procedure MenuOnChange(Sender: TObject; Source: TMenuItem; Rebuild: Boolean);
    procedure Add(AObj: TObject; AEvent: TGoEvent; AId: NativeUInt);
    procedure AddClick(Sender: TObject; AId: NativeUInt);
    procedure Remove(AObj: TObject; AEvent: TGoEvent);
    procedure ThreadProc;

  public
    property ThreadEvtId: NativeUInt read FThreadEvtId write FThreadEvtId;
  end;


var
  EventClass: TEventClass;

implementation

function HashOf(const Key: TEventKey): NativeUInt;
var
  I: Integer;
  P: PByte;
begin
   Result := 0;
   P := @Key;
    for I := 1 to SizeOf(Key) do
    begin
      Result := ((Result shl 2) or (Result shr (SizeOf(Result) * 8 - 2))) xor P^;
      Inc(P);
    end;
end;


function CreateEventKey(ASender: TObject; AEvent: TGoEvent): NativeUInt;
var
  LEvent: TEventKey;
begin
  LEvent.Sender := ASender;
  LEvent.Event := AEvent;
  Result := HashOf(LEvent);
end;


{ TEvent}


constructor TEventClass.Create;
begin
  inherited;
  FEvents := TEventList.Create;
end;

destructor TEventClass.Destroy;
begin
  FEvents.Free;
end;

procedure TEventClass.Add(AObj: TObject; AEvent: TGoEvent; AId: NativeUInt);
begin
  //writeln('addevent: id:', aid, ', event:', ord(Aevent), ', sender:', AObj.ToString);
  // I服了U
  if AObj is TTrayIcon then
     AObj := Application;
  FEvents.AddOrSetData(CreateEventKey(AObj, AEvent), AId);
  //FEvents.insert(CreateEventKey(AObj, AEvent), AId);
end;

procedure TEventClass.Remove(AObj: TObject; AEvent: TGoEvent);
begin
  FEvents.delete(CreateEventKey(AObj, AEvent));
end;

procedure TEventClass.ThreadProc;
begin
  GCallbackPtr(ThreadEvtId, nil, 0);
end;

procedure TEventClass.AddClick(Sender: TObject; AId: NativeUInt);
begin
  Add(Sender, geClick, AId);
end;

procedure TEventClass.OnBalloonClick(Sender: TObject);
begin
  SendEvent(Sender, geBalloonClick, [Sender]);
end;

procedure TEventClass.OnChange(Sender: TObject);
begin
  SendEvent(Sender, geChange, [Sender]);
end;

procedure TEventClass.OnClick(Sender: TObject);
begin
  SendEvent(Sender, geClick, [Sender]);
end;

procedure TEventClass.OnClickCheck(Sender: TObject);
begin
  SendEvent(Sender, geClickCheck, [Sender]);
end;
procedure TEventClass.FormOnClose(Sender: TObject; var Action: TCloseAction);
begin
  SendEvent(Sender, geFormClose, [Sender, @Action]);
end;

procedure TEventClass.OnReplace(Sender: TObject);
begin
  SendEvent(Sender, geReplace, [Sender]);
end;

procedure TEventClass.OnFind(Sender: TObject);
begin
  SendEvent(Sender, geFind, [Sender]);
end;



procedure TEventClass.OnActivate(Sender: TObject);
begin
  SendEvent(Sender, geActivate, [Sender]);
end;

procedure TEventClass.OnDeactivate(Sender: TObject);
begin
  SendEvent(Sender, geDeactivate, [Sender]);
end;

procedure TEventClass.OnConstrainedResize(Sender: TObject; var MinWidth, MinHeight, MaxWidth, MaxHeight: TConstraintSize);
begin
  SendEvent(Sender, geConstrainedResize, [Sender, @MinWidth, @MinHeight, @MaxWidth, @MaxHeight]);
end;





function TEventClass.OnHelp(Command: Word; Data: PtrInt;
  var CallHelp: Boolean): Boolean;
var
  LResult: Boolean;
begin
  SendEvent(Application, geHelp, [Command, Data, Pointer(@CallHelp), Pointer(@LResult)]);
  Result := LResult;
end;

procedure TEventClass.OnShortCut(var Msg: TLMKey; var Handled: Boolean);
begin
  SendEvent(Application, geShortCut, [Pointer(@Msg), Pointer(@Handled)]);
end;

procedure TEventClass.OnContextPopup(Sender: TObject; MousePos: TPoint; var Handled: Boolean);
begin
  SendEvent(Sender, geContextPopup, [Sender, Pointer(@MousePos), Pointer(@Handled)]);
end;

procedure TEventClass.OnDockDrop(Sender: TObject; Source: TDragDockObject; X, Y: Integer);
begin
  SendEvent(Sender, geDockDrop, [Sender, Source, X, Y]);
end;

procedure TEventClass.OnDragDrop(Sender, Source: TObject; X, Y: Integer);
begin
  SendEvent(Sender, geDragDrop, [Sender, Source, X, Y]);
end;

procedure TEventClass.OnDragOver(Sender, Source: TObject; X, Y: Integer; State: TDragState; var Accept: Boolean);
begin
  SendEvent(Sender, geDragOver, [Sender, Source, X, Y, Integer(State), Pointer(@Accept)]);
end;

procedure TEventClass.OnEndDock(Sender, Target: TObject; X, Y: Integer);
begin
  SendEvent(Sender, geEndDock, [Sender, Target, X, Y]);
end;

procedure TEventClass.OnGetSiteInfo(Sender: TObject; DockClient: TControl;
  var InfluenceRect: TRect; MousePos: TPoint; var CanDock: Boolean);
begin
  SendEvent(Sender, geGetSiteInfo, [Sender, DockClient, Pointer(@InfluenceRect), Pointer(@MousePos), Pointer(@CanDock)]);
end;

procedure TEventClass.OnMouseWheelDown(Sender: TObject; Shift: TShiftState; MousePos: TPoint; var Handled: Boolean);
begin
  SendEvent(Sender, geMouseWheelDown, [Sender, PWord(@Shift)^, Pointer(@MousePos), Pointer(@Handled)]);
end;

procedure TEventClass.OnMouseWheelUp(Sender: TObject; Shift: TShiftState; MousePos: TPoint; var Handled: Boolean);
begin
  SendEvent(Sender, geMouseWheelUp, [Sender, PWord(@Shift)^, Pointer(@MousePos), Pointer(@Handled)]);
end;

procedure TEventClass.OnStartDock(Sender: TObject; var DragObject: TDragDockObject);
begin
  SendEvent(Sender, geStartDock, [Sender, DragObject]);
end;

procedure TEventClass.OnUnDock(Sender: TObject; Client: TControl; NewTarget: TWinControl; var Allow: Boolean);
begin
  SendEvent(Sender, geUnDock, [Sender, Client, NewTarget, Pointer(@Allow)]);
end;

procedure TEventClass.OnEndDrag(Sender, Target: TObject; X, Y: Integer);
begin
  SendEvent(Sender, geEndDrag, [Sender, Target, X, Y]);
end;


//procedure TEventClass.OnMessage(var Msg: TMsg; var Handled: Boolean);
//begin
//  SendEvent(Application, geMessage, [Pointer(@Msg), Pointer(@Handled)]);
//end;




//----------------------- TListView

procedure TEventClass.ListViewOnChange(Sender: TObject; AItem: TListItem; Change: TItemChange);
begin
  SendEvent(Sender, geListViewChange, [Sender, AItem, Ord(Change)]);
end;

procedure TEventClass.ListViewOnColumnClick(Sender: TObject; Column: TListColumn);
begin
  SendEvent(Sender, geListViewColumnClick, [Sender, Column]);
end;

procedure TEventClass.ListViewOnColumnRightClick(Sender: TObject; Column: TListColumn; Point: TPoint);
begin
  SendEvent(Sender, geListViewColumnRightClick, [Sender, Column, Point.X, Point.Y]);
end;

procedure TEventClass.ListViewOnGetImageIndex(Sender: TObject; Item: TListItem);
begin
  SendEvent(Sender, geListViewGetImageIndex, [Sender, Item]);
end;

procedure TEventClass.ListViewOnSelectItem(Sender: TObject; Item: TListItem; Selected: Boolean);
begin
  SendEvent(Sender, geListViewSelectItem, [Sender, Item, Selected]);
end;

procedure TEventClass.ListViewOnItemChecked(Sender: TObject; Item: TListItem);
begin
  SendEvent(Sender, geListViewItemChecked, [Sender, Item]);
end;

procedure TEventClass.ListViewOnCompare(Sender: TObject; Item1, Item2: TListItem; Data: Integer; var Compare: Integer);
begin
  SendEvent(Sender, geListViewCompare, [Sender, Item1, Item2, Data, @Compare]);
end;

procedure TEventClass.ListViewOnAdvancedCustomDraw(Sender: TCustomListView;
  const ARect: TRect; Stage: TCustomDrawStage; var DefaultDraw: Boolean);
begin
  SendEvent(Sender, geListViewAdvancedCustomDraw, [
    Sender, @ARect, Ord(Stage), @DefaultDraw
  ]);
end;

procedure TEventClass.ListViewOnAdvancedCustomDrawItem(
  Sender: TCustomListView; Item: TListItem; State: TCustomDrawState;
  Stage: TCustomDrawStage; var DefaultDraw: Boolean);
begin
  SendEvent(Sender, geListViewAdvancedCustomDrawItem, [
    Sender, Item, PWord(@State)^, Ord(Stage), @DefaultDraw
  ]);
end;

procedure TEventClass.ListViewOnAdvancedCustomDrawSubItem(
  Sender: TCustomListView; Item: TListItem; SubItem: Integer;
  State: TCustomDrawState; Stage: TCustomDrawStage; var DefaultDraw: Boolean);
begin
  SendEvent(Sender, geListViewAdvancedCustomDrawSubItem, [
    Sender, Item, SubItem, PWord(@State)^, Ord(Stage), @DefaultDraw
  ]);
end;

//----------------------------------------- TPageControl
procedure TEventClass.PageControlOnGetImageIndex(Sender: TObject; TabIndex: Integer; var ImageIndex: Integer);
begin
  SendEvent(Sender, gePageControlGetImageIndex, [Sender, TabIndex, @ImageIndex]);
end;
procedure TEventClass.MenuOnChange(Sender: TObject; Source: TMenuItem;
  Rebuild: Boolean);
begin
  SendEvent(Sender, geMenuChange, [Sender, Source, Rebuild]);
end;

procedure TEventClass.OnClose(Sender: TObject);
begin
  SendEvent(Sender, geClose, [Sender]);
end;

procedure TEventClass.FormOnCloseQuery(Sender: TObject;
  var CanClose: Boolean);
begin
  SendEvent(Sender, geFormCloseQuery, [Sender, @CanClose]);
end;

procedure TEventClass.FormOnDropFiles(Sender: TObject;
  const AFileNames: array of string);
var
  LLen: Integer;
begin
  LLen := Length(AFileNames);
  if LLen > 0 then
    SendEvent(Sender, geDropFiles, [Sender, @AFileNames[0], LLen]);
end;

procedure TEventClass.OnDblClick(Sender: TObject);
begin
  SendEvent(Sender, geDblClick, [Sender]);
end;

procedure TEventClass.OnEnter(Sender: TObject);
begin
  SendEvent(Sender, geEnter, [Sender]);
end;

procedure TEventClass.OnException(Sender: TObject; E: Exception);
begin
  SendEvent(Sender, geException, [Sender, E]);
end;

procedure TEventClass.OnExecute(Sender: TObject);
begin
  SendEvent(Sender, geExecute, [Sender]);
end;

procedure TEventClass.OnExit(Sender: TObject);
begin
  SendEvent(Sender, geExit, [Sender]);
end;

procedure TEventClass.OnHide(Sender: TObject);
begin
  SendEvent(Sender, geHide, [Sender]);
end;

procedure TEventClass.OnDestroy(Sender: TObject);
begin
  SendEvent(Sender, geDestroy, [Sender]);
end;

procedure TEventClass.OnHint(Sender: TObject);
begin
   SendEvent(Sender, geHint, [Sender]);
end;
procedure TEventClass.OnKeyDown(Sender: TObject; var Key: Word;
  Shift: TShiftState);
begin
  SendEvent(Sender, geKeyDown, [Sender, @Key, PWord(@Shift)^]);
end;

procedure TEventClass.OnKeyPress(Sender: TObject; var Key: Char);
begin
  SendEvent(Sender, geKeyPress, [Sender, @Key]);
end;

procedure TEventClass.OnKeyUp(Sender: TObject; var Key: Word;
  Shift: TShiftState);
begin
  SendEvent(Sender, geKeyUp, [Sender, @Key, PWord(@Shift)^]);
end;

procedure TEventClass.OnMinimize(Sender: TObject);
begin
  SendEvent(Sender, geMinimize, [Sender]);
end;

procedure TEventClass.OnMouseDown(Sender: TObject; Button: TMouseButton;
  Shift: TShiftState; X, Y: Integer);
begin
  SendEvent(Sender, geMouseDown, [Sender, Ord(Button), PWord(@Shift)^, X, Y]);
end;

procedure TEventClass.OnMouseEnter(Sender: TObject);
begin
  SendEvent(Sender, geMouseEnter, [Sender]);
end;

procedure TEventClass.OnMouseLeave(Sender: TObject);
begin
  SendEvent(Sender, geMouseLeave, [Sender]);
end;

procedure TEventClass.ListBoxOnDrawItem(Control: TWinControl; Index: Integer;
  ARect: TRect; State: TOwnerDrawState);
begin
  SendEvent(Control, geListBoxDrawItem, [Control, Index, @ARect, PWord(@State)^]);
end;

procedure TEventClass.OnMouseMove(Sender: TObject; Shift: TShiftState; X,
  Y: Integer);
begin
  SendEvent(Sender, geMouseMove, [Sender, PWord(@Shift)^, X, Y]);
end;

procedure TEventClass.OnMouseUp(Sender: TObject; Button: TMouseButton;
  Shift: TShiftState; X, Y: Integer);
begin
  SendEvent(Sender, geMouseUp, [Sender, Ord(Button), PWord(@Shift)^, X, Y]);
end;

procedure TEventClass.OnMouseWheel(Sender: TObject; Shift: TShiftState;
  WheelDelta: Integer; MousePos: TPoint; var Handled: Boolean);
begin
  SendEvent(Sender, geMouseWheel, [Sender, PWord(@Shift)^, WheelDelta, MousePos.X, MousePos.Y, @Handled]);
end;

procedure TEventClass.OnLinkClick(Sender: TObject; const Link: string;
  LinkType: TSysLinkType);
begin
  SendEvent(Sender, geLinkClick, [Sender, Link, Ord(LinkType)]);
end;

procedure TEventClass.OnPaint(Sender: TObject);
begin
  SendEvent(Sender, gePaint, [Sender]);
end;

procedure TEventClass.OnPopup(Sender: TObject);
begin
  SendEvent(Sender, gePopup, [Sender]);
end;

procedure TEventClass.OnResize(Sender: TObject);
begin
  SendEvent(Sender, geResize, [Sender]);
end;

procedure TEventClass.OnRestore(Sender: TObject);
begin
  SendEvent(Sender, geRestore, [Sender]);
end;

procedure TEventClass.OnShow(Sender: TObject);
begin
  SendEvent(Sender, geShow, [Sender]);
end;

procedure TEventClass.OnTimer(Sender: TObject);
begin
  SendEvent(Sender, geTimer, [Sender]);
end;

procedure TEventClass.OnUpdate(Sender: TObject);
begin
  SendEvent(Sender, geUpdate, [Sender]);
end;


procedure TEventClass.SendEvent(Sender: TObject; AEvent: TGoEvent; AArgs: array of const);


  procedure SendEventSrc(EventId: NativeUInt; AArgs: array of const);
  var
    LParams: array[0..29] of TGoParam;
    LArgLen: Integer;
    LV: TVarRec;
    I: Integer;
  begin
    if Assigned(GCallbackPtr) and (EventId > 0) then
    begin
      LArgLen := Length(AArgs);
      if LArgLen <= Length(LParams) then
      begin
        for I := 0 to LArgLen - 1 do
        begin
          LV := AArgs[I];
          LParams[I].&Type := LV.VType;
          case LV.VType of
            vtInteger       : LParams[I].Value := Pointer(LV.VInteger);
            vtBoolean       : LParams[I].Value := Pointer(Byte(LV.VBoolean));
            vtChar          : LParams[I].Value := Pointer(Ord(LV.VChar));
            vtExtended      : LParams[I].Value := LV.VExtended;

            vtString        : LParams[I].Value := {$IFDEF MSWINDOWS}LV.VString{$ELSE}LV.VAnsiString{$ENDIF};

            vtPointer       : LParams[I].Value := LV.VPointer;
            vtPChar         : LParams[I].Value := LV.VPChar;
            vtObject        : LParams[I].Value := LV.VObject;
            vtClass         : LParams[I].Value := LV.VClass;
            vtWideChar      : LParams[I].Value := Pointer(Ord(LV.VWideChar));
            vtPWideChar     : LParams[I].Value := LV.VPWideChar;
            vtAnsiString    : LParams[I].Value := LV.VAnsiString;
  //          vtCurrency      = 12;
  //          vtVariant       = 13;
            vtInterface     : LParams[I].Value := LV.VInterface;
            vtWideString    : LParams[I].Value := LV.VWideString;
            vtInt64         : LParams[I].Value := LV.VInt64;
            vtUnicodeString : LParams[I].Value := LV.VUnicodeString;
          end;
        end;
        GCallbackPtr(EventId, @LParams[0], LArgLen);
      end;
    end;
  end;

var
  LEventId: NativeUInt;
begin
  if FEvents.TryGetData(CreateEventKey(Sender, AEvent), LEventId) then
  //if FEvents.GetValue(CreateEventKey(Sender, AEvent), LEventId) then
    SendEventSrc(LEventId, AArgs)
  else writeln('can''t found id, sender:', sender.ToString, ', event:', ord(AEvent));
end;

// ---------------------- TTreeView

procedure TEventClass.TreeViewOnAdvancedCustomDraw(
  Sender: TCustomTreeView; const ARect: TRect; Stage: TCustomDrawStage;
  var DefaultDraw: Boolean);
begin
  SendEvent(Sender, geTreeViewAdvancedCustomDraw,
    [Sender, @ARect, Ord(Stage), DefaultDraw]);
end;

procedure TEventClass.TreeViewOnAdvancedCustomDrawItem(
  Sender: TCustomTreeView; Node: TTreeNode; State: TCustomDrawState;
  Stage: TCustomDrawStage; var PaintImages, DefaultDraw: Boolean);
begin
  SendEvent(Sender, geTreeViewAdvancedCustomDrawItem,
    [Sender, Node, PWord(@State)^, Ord(Stage), @PaintImages, DefaultDraw]);
end;

procedure TEventClass.TreeViewOnChange(Sender: TObject; ANode: TTreeNode);
begin
  SendEvent(Sender, geTreeViewChange, [Sender, ANode]);
end;

procedure TEventClass.TreeViewOnGetImageIndex(Sender: TObject; Node: TTreeNode);
begin
  SendEvent(Sender, geTreeViewGetImageIndex, [Sender, Node]);
end;

procedure TEventClass.TreeViewOnGetSelectedIndex(Sender: TObject; Node: TTreeNode);
begin
  SendEvent(Sender, geTreeViewGetSelectedIndex, [Sender, Node]);
end;

procedure TEventClass.TreeViewOnCompare(Sender: TObject; Node1, Node2: TTreeNode; var Compare: Integer);
begin
  SendEvent(Sender, geTreeViewCompare, [Sender, Node1, Node2, 0, @Compare]);
end;
//----------- TToolBar
procedure TEventClass.ToolBarOnAdvancedCustomDraw(Sender: TToolBar;
  const ARect: TRect; Stage: TCustomDrawStage; var DefaultDraw: Boolean);
begin
  SendEvent(Sender, geToolBarAdvancedCustomDraw,
    [Sender, @ARect, Ord(Stage), @DefaultDraw]);
end;

//procedure TEventClass.ToolBarOnAdvancedCustomDrawButton(Sender: TToolBar;
//  Button: TToolButton; State: TCustomDrawState; Stage: TCustomDrawStage;
//  var Flags: TTBCustomDrawFlags; var DefaultDraw: Boolean);
//begin
//  SendEvent(Sender, geToolBarAdvancedCustomDrawButton,
//    [Sender, Button, PWord(@State)^, Ord(Stage), PWord(@Flags), @DefaultDraw]);
//end;


procedure TEventClass.UpDownOnClick(Sender: TObject; Button: TUDBtnType);
begin
  SendEvent(Sender, geUpDownClick, [Sender, Ord(Button)]);
end;


initialization

  EventClass := TEventClass.Create;

finalization
  EventClass.Free;

end.
