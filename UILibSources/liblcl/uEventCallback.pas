//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
// 
// Licensed under Apache License 2.0
//
//----------------------------------------
// 事件回调

unit uEventCallback;

{$mode objfpc}{$H+}
//{$mode delphi}

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
  Graphics,
  StdCtrls,
  LMessages,
  Grids,
  uLinkLabel,
  fgl;

const
  // 线程同步消息
  //THREAD_SYNC_MESSAGE = WM_USER + $9090;
  // call最长参数数，与导出的MySyscall一致，暂定为12个
  CALL_MAX_PARAM = 12;

type
  TEventCallbackPtr = function(f: NativeUInt; args: Pointer; argcout: NativeInt): Pointer; extdecl;
  TMessageCallbackPtr = function(f: NativeUInt; msg: Pointer): Pointer; extdecl;
  TThreadSyncCallbackPtr = function{(f: NativeUInt)}: Pointer; extdecl;

var
  GEventCallbackPtr: TEventCallbackPtr;
  GMessageCallbackPtr: TMessageCallbackPtr;
  GThreadSyncCallbackPtr: TThreadSyncCallbackPtr;

type

  TEventKey = packed record
    Sender: TObject;
    Event: Pointer;
  end;
  PEventKey = ^TEventKey;

  { TEventList }

  TEventList = specialize  TFPGMap<NativeUInt, NativeUInt>;


  { TEventClass }

  TEventClass = class
  private class var
    FEvents: TEventList;
    FThreadEvtId: NativeUInt;
    class procedure SendEvent(Sender: TObject; AEvent: Pointer; AArgs: array of const);
  public
    class constructor Create;
    class destructor Destroy;
  public
    class procedure OnClick(Sender: TObject);

    class procedure FormOnClose(Sender: TObject; var Action: TCloseAction);
    class procedure FormOnCloseQuery(Sender: TObject; var CanClose: Boolean);
    class procedure FormOnDropFiles(Sender: TObject; const AFileNames: array of string);

    class procedure OnClose(Sender: TObject);

    class procedure OnChange(Sender: TObject);

    class procedure UpDownOnClick(Sender: TObject; Button: TUDBtnType);

    class procedure TreeViewOnChange(Sender: TObject; ANode: TTreeNode);
    class procedure TreeViewOnGetImageIndex(Sender: TObject; Node: TTreeNode);
    class procedure TreeViewOnGetSelectedIndex(Sender: TObject; Node: TTreeNode);
    class procedure TreeViewOnCompare(Sender: TObject; Node1, Node2: TTreeNode; var Compare: Integer);
    class procedure TreeViewOnAdvancedCustomDraw(Sender: TCustomTreeView;
      const ARect: TRect; Stage: TCustomDrawStage; var DefaultDraw: Boolean);
    class procedure TreeViewOnAdvancedCustomDrawItem(Sender: TCustomTreeView;
      Node: TTreeNode; State: TCustomDrawState; Stage: TCustomDrawStage;
      var PaintImages, DefaultDraw: Boolean);


    class procedure ListViewOnChange(Sender: TObject; AItem: TListItem; Change: TItemChange);
    class procedure ListViewOnColumnClick(Sender: TObject; Column: TListColumn);
    class procedure ListViewOnColumnRightClick(Sender: TObject; Column: TListColumn; Point: TPoint);
    class procedure ListViewOnGetImageIndex(Sender: TObject; Item: TListItem);
    class procedure ListViewOnSelectItem(Sender: TObject; Item: TListItem; Selected: Boolean);
    class procedure ListViewOnItemChecked(Sender: TObject; Item: TListItem);
    class procedure ListViewOnCompare(Sender: TObject; Item1, Item2: TListItem; Data: Integer; var Compare: Integer);
    class procedure ListViewOnAdvancedCustomDraw(Sender: TCustomListView;
       const ARect: TRect; Stage: TCustomDrawStage; var DefaultDraw: Boolean);
    class procedure ListViewOnAdvancedCustomDrawItem(Sender: TCustomListView;
      Item: TListItem; State: TCustomDrawState; Stage: TCustomDrawStage;
      var DefaultDraw: Boolean);
    class procedure ListViewOnAdvancedCustomDrawSubItem(Sender: TCustomListView;
      Item: TListItem; SubItem: Integer; State: TCustomDrawState;
      Stage: TCustomDrawStage; var DefaultDraw: Boolean);

    class procedure PageControlOnGetImageIndex(Sender: TObject; TabIndex: Integer; var ImageIndex: Integer);


    class procedure ToolBarOnAdvancedCustomDraw(Sender: TToolBar;
       const ARect: TRect; Stage: TCustomDrawStage; var DefaultDraw: Boolean);
    //class procedure ToolBarOnAdvancedCustomDrawButton(Sender: TToolBar;
      //Button: TToolButton; State: TCustomDrawState; Stage: TCustomDrawStage;
      //var Flags: TTBCustomDrawFlags; var DefaultDraw: Boolean);


    class procedure OnDblClick(Sender: TObject);

    class procedure OnPaint(Sender: TObject);
    class procedure OnResize(Sender: TObject);
    class procedure OnShow(Sender: TObject);
    class procedure OnEnter(Sender: TObject);
    class procedure OnExit(Sender: TObject);
    class procedure OnPopup(Sender: TObject);
    class procedure OnHint(Sender: TObject);
    class procedure OnClickCheck(Sender: TObject);

    class procedure OnExecute(Sender: TObject);
    class procedure OnUpdate(Sender: TObject);

    class procedure OnBalloonClick(Sender: TObject);

    class procedure OnException(Sender: TObject; E: Exception);
    class procedure OnTimer(Sender: TObject);

    class procedure OnMinimize(Sender: TObject);
    class procedure OnRestore(Sender: TObject);
    class procedure OnHide(Sender: TObject);

    class procedure OnDestroy(Sender: TObject);
    class procedure OnReplace(Sender: TObject);
    class procedure OnFind(Sender: TObject);

    class procedure OnActivate(Sender: TObject);
    class procedure OnDeactivate(Sender: TObject);
    class procedure OnConstrainedResize(Sender: TObject; var MinWidth, MinHeight, MaxWidth, MaxHeight: TConstraintSize);


    // new
    class function OnHelp(Command: Word; Data: PtrInt; var CallHelp: Boolean): Boolean;
    class procedure OnShortCut(var Msg: TLMKey; var Handled: Boolean);
    class procedure OnContextPopup(Sender: TObject; MousePos: TPoint; var Handled: Boolean);
    class procedure OnDockDrop(Sender: TObject; Source: TDragDockObject; X, Y: Integer);
    class procedure OnDragDrop(Sender, Source: TObject; X, Y: Integer);
    class procedure OnDragOver(Sender, Source: TObject; X, Y: Integer; State: TDragState; var Accept: Boolean);
    class procedure OnEndDock(Sender, Target: TObject; X, Y: Integer);
    class procedure OnGetSiteInfo(Sender: TObject; DockClient: TControl; var InfluenceRect: TRect; MousePos: TPoint; var CanDock: Boolean);
    class procedure OnMouseWheelDown(Sender: TObject; Shift: TShiftState; MousePos: TPoint; var Handled: Boolean);
    class procedure OnMouseWheelUp(Sender: TObject; Shift: TShiftState; MousePos: TPoint; var Handled: Boolean);
    class procedure OnStartDock(Sender: TObject; var DragObject: TDragDockObject);
    class procedure OnUnDock(Sender: TObject; Client: TControl; NewTarget: TWinControl; var Allow: Boolean);
    class procedure OnEndDrag(Sender, Target: TObject; X, Y: Integer);



    //class procedure OnGesture(Sender: TObject; const EventInfo: TGestureEventInfo; var Handled: Boolean);
    //class procedure OnMouseActivate(Sender: TObject; Button: TMouseButton; Shift: TShiftState; X, Y, HitTest: Integer; var MouseActivate: TMouseActivate);
    //class procedure ListBoxOnData(Control: TWinControl; Index: Integer; var Data: string);
    //class function ListBoxOnDataFind(Control: TWinControl; FindString: string): Integer;
    //class procedure ListBoxOnDataObject(Control: TWinControl; Index: Integer; var DataObject: TObject);
    class procedure ListBoxOnMeasureItem(Control: TWinControl; Index: Integer; var Height: Integer);
    class procedure OnChanging(Sender: TObject);
    class procedure UpDownOnChanging(Sender: TObject; var AllowChange: Boolean);

    class procedure ListViewOnChanging(Sender: TObject; Item: TListItem; Change: TItemChange; var AllowChange: Boolean);
    class procedure ListViewOnData(Sender: TObject; Item: TListItem);
    class procedure ListViewOnDataFind(Sender: TObject; Find: TItemFind;
      const FindString: string; const FindPosition: TPoint; FindData: Pointer;
      StartIndex: Integer; Direction: TSearchDirection; Wrap: Boolean; var Index: Integer);
    class procedure ListViewOnEdited(Sender: TObject; Item: TListItem; var S: string);
    class procedure ListViewOnEditing(Sender: TObject; Item: TListItem; var AllowEdit: Boolean);
    class procedure ListViewOnInsert(Sender: TObject; Item: TListItem);
    class procedure ListViewOnDeletion(Sender: TObject; Item: TListItem);

    class procedure ListViewOnCustomDraw(Sender: TCustomListView; const ARect: TRect; var DefaultDraw: Boolean);
    class procedure ListViewOnCustomDrawItem(Sender: TCustomListView; Item: TListItem; State: TCustomDrawState; var DefaultDraw: Boolean);
    class procedure ListViewOnCustomDrawSubItem(Sender: TCustomListView; Item: TListItem; SubItem: Integer; State: TCustomDrawState; var DefaultDraw: Boolean);
    class procedure ListViewOnDrawItem(Sender: TCustomListView; Item: TListItem; ARect: TRect; State: TOwnerDrawState);

    class procedure ListViewOnDataHint(Sender: TObject; StartIndex, EndIndex: Integer);




    class procedure TreeViewOnChanging(Sender: TObject; Node: TTreeNode; var AllowChange: Boolean);
    //class procedure TreeViewOnCancelEdit(Sender: TObject; Node: TTreeNode);
    class procedure TreeViewOnAddition(Sender: TObject; Node: TTreeNode);
    class procedure TreeViewOnCollapsed(Sender: TObject; Node: TTreeNode);
    class procedure TreeViewOnCollapsing(Sender: TObject; Node: TTreeNode; var AllowCollapse: Boolean);
    class procedure TreeViewOnDeletion(Sender: TObject; Node: TTreeNode);
    class procedure TreeViewOnEdited(Sender: TObject; Node: TTreeNode; var S: string);
    class procedure TreeViewOnEditing(Sender: TObject; Node: TTreeNode; var AllowEdit: Boolean);
    class procedure TreeViewOnExpanded(Sender: TObject; Node: TTreeNode);
    class procedure TreeViewOnExpanding(Sender: TObject; Node: TTreeNode; var AllowExpansion: Boolean);
    //class procedure TreeViewOnHint(Sender: TObject; const Node: TTreeNode; var Hint: string);
    class procedure TreeViewOnCustomDraw(Sender: TCustomTreeView; const ARect: TRect; var DefaultDraw: Boolean);
    class procedure TreeViewOnCustomDrawItem(Sender: TCustomTreeView; Node: TTreeNode; State: TCustomDrawState; var DefaultDraw: Boolean);



    class procedure MenuItemOnMeasureItem(Sender: TObject; ACanvas: TCanvas; var Width, Height: Integer);
    class procedure PageControlOnChanging(Sender: TObject; var AllowChange: Boolean);


    // grid

    class procedure OnColumnMoved(Sender: TObject; IsColumn: Boolean; FromIndex, ToIndex: Longint);
    class procedure OnDrawCell(Sender: TObject; ACol, ARow: Longint; ARect: TRect; State: TGridDrawState);
    class procedure OnFixedCellClick(Sender: TObject; ACol, ARow: Integer);
    class procedure OnGetEditMask(Sender: TObject; ACol, ARow: Integer; var Value: string);
    class procedure OnGetEditText(Sender: TObject; ACol, ARow: Integer; var Value: string);
    class procedure OnRowMoved(Sender: TObject; IsColumn: Boolean; FromIndex, ToIndex: Integer);
    class procedure OnSelectCell(Sender: TObject; ACol, ARow: Integer; var CanSelect: Boolean);
    class procedure OnSetEditText(Sender: TObject; ACol, ARow: Integer; const Value: string);
    class procedure OnTopLeftChanged(Sender: TObject);



    // headercontrol
    //class procedure OnDrawSection(HeaderControl: THeaderControl; Section: THeaderSection; const Rect: TRect; Pressed: Boolean);
    //class procedure OnSectionCheck(HeaderControl: TCustomHeaderControl; Section: THeaderSection);
    class procedure OnSectionClick(HeaderControl: TCustomHeaderControl; Section: THeaderSection);
    class procedure OnSectionDrag(Sender: TObject; FromSection, ToSection: THeaderSection; var AllowDrag: Boolean);
    class procedure OnSectionEndDrag(Sender: TObject);
    class procedure OnSectionResize(HeaderControl: TCustomHeaderControl; Section: THeaderSection);
    class procedure OnSectionTrack(HeaderControl: TCustomHeaderControl; Section: THeaderSection; Width: Integer; State: TSectionTrackState);



    class procedure OnKeyDown(Sender: TObject; var Key: Word; Shift: TShiftState);
    class procedure OnKeyUp(Sender: TObject; var Key: Word; Shift: TShiftState);
    class procedure OnKeyPress(Sender: TObject; var Key: Char);

    class procedure OnMouseDown(Sender: TObject; Button: TMouseButton;
           Shift: TShiftState; X, Y: Integer);
    class procedure OnMouseMove(Sender: TObject; Shift: TShiftState; X, Y: Integer);
    class procedure OnMouseUp(Sender: TObject; Button: TMouseButton;
          Shift: TShiftState; X, Y: Integer);
    class procedure OnMouseWheel(Sender: TObject; Shift: TShiftState;
          WheelDelta: Integer; MousePos: TPoint; var Handled: Boolean);

    class procedure OnMouseEnter(Sender: TObject);
    class procedure OnMouseLeave(Sender: TObject);


    class procedure ListBoxOnDrawItem(Control: TWinControl; Index: Integer;
          ARect: TRect; State: TOwnerDrawState);

    class procedure OnLinkClick(Sender: TObject; const Link: string; LinkType: TSysLinkType);

    class procedure MenuOnChange(Sender: TObject; Source: TMenuItem; Rebuild: Boolean);

    // webbrowser
    class procedure OnTitleChange(Sender: TObject; const Text: string);
    class procedure OnJSExternal(Sender: TObject; const Afunc: string; const AArgs: WideString; var ARetval: WideString);

    // TaskDialog
    class procedure TaskDialogOnButtonClicked(Sender: TObject; ModalResult: TModalResult; var CanClose: Boolean);
    // Lazarus不支持
    //class procedure TaskDialogOnDialogConstructed(Sender: TObject);
    //class procedure TaskDialogOnDialogCreated(Sender: TObject);
    //class procedure TaskDialogOnDialogDestroyed(Sender: TObject);
    //class procedure TaskDialogOnExpanded(Sender: TObject);
    //class procedure TaskDialogOnHyperlinkClicked(Sender: TObject);
    //class procedure TaskDialogOnNavigated(Sender: TObject);
    //class procedure TaskDialogOnRadioButtonClicked(Sender: TObject);
    //class procedure TaskDialogOnTimer(Sender: TObject; TickCount: Cardinal; var Reset: Boolean);
    //class procedure TaskDialogOnVerificationClicked(Sender: TObject);


    //    class function OnAlignInsertBefore(Sender: TWinControl; C1, C2: TControl): Boolean;
    class procedure OnAlignPosition(Sender: TWinControl; Control: TControl;
       var NewLeft, NewTop, NewWidth, NewHeight: Integer; var AlignRect: TRect; AlignInfo: TAlignInfo);

    class procedure OnDropDown(Sender: TObject);
    class procedure OnSelect(Sender: TObject);
    class procedure OnBeginEdit(Sender: TObject);
    class procedure OnEndEdit(Sender: TObject);


    class procedure Add(AObj: TObject; AEvent: Pointer; AId: NativeUInt);
    class procedure AddClick(Sender: TObject; AId: NativeUInt);
    class procedure Remove(AObj: TObject; AEvent: Pointer);
    class procedure ThreadProc;

    // 用户定义事件声明
    {$I UserDefineEventsDeclaration.inc}
  public
    class property ThreadEvtId: NativeUInt read FThreadEvtId write FThreadEvtId;
  end;

  // 窗口消息的，不与之前的事件混在一起。
  TMessageEventList = specialize  TFPGMap<NativeUInt, NativeUInt>;

  { TMessageEventClass }

  TMessageEventClass = class
  private class var
    FMsgEvents: TMessageEventList;
  public
    class constructor Create;
    class destructor Destroy;
    class procedure Add(AObj: TObject; AId: NativeUInt);
    class procedure Remove(AObj: TObject);

    class procedure OnWndProc(Sender: TObject; var TheMessage: TLMessage);
  end;

implementation

function HashOf(const Key: TEventKey): NativeUInt; inline;
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


function CreateEventKey(ASender: TObject; AEvent: Pointer): NativeUInt; inline;
var
  LEvent: TEventKey;
begin
  LEvent.Sender := ASender;
  LEvent.Event := AEvent;
  Result := HashOf(LEvent);
end;


{ TEvent}


class constructor TEventClass.Create;
begin
  FEvents := TEventList.Create;
end;

class destructor TEventClass.Destroy;
begin
  FEvents.Free;
end;

class procedure TEventClass.Add(AObj: TObject; AEvent: Pointer; AId: NativeUInt);
begin
  if AObj is TTrayIcon then
     AObj := Application;
  FEvents.AddOrSetData(CreateEventKey(AObj, AEvent), AId);
end;

class procedure TEventClass.Remove(AObj: TObject; AEvent: Pointer);
begin
  FEvents.delete(CreateEventKey(AObj, AEvent));
end;

class procedure TEventClass.ThreadProc;
begin
  GThreadSyncCallbackPtr();
end;

class procedure TEventClass.AddClick(Sender: TObject; AId: NativeUInt);
begin
  Add(Sender, @TEventClass.OnClick , AId);
end;

class procedure TEventClass.OnBalloonClick(Sender: TObject);
begin
  SendEvent(Sender, @TEventClass.OnBalloonClick, [Sender]);
end;

class procedure TEventClass.OnChange(Sender: TObject);
begin
  SendEvent(Sender, @TEventClass.OnChange, [Sender]);
end;

class procedure TEventClass.OnClick(Sender: TObject);
begin
  SendEvent(Sender, @TEventClass.OnClick, [Sender]);
end;

class procedure TEventClass.OnClickCheck(Sender: TObject);
begin
  SendEvent(Sender, @TEventClass.OnClickCheck, [Sender]);
end;
class procedure TEventClass.FormOnClose(Sender: TObject; var Action: TCloseAction);
begin
  SendEvent(Sender, @TEventClass.FormOnClose, [Sender, @Action]);
end;

class procedure TEventClass.OnReplace(Sender: TObject);
begin
  SendEvent(Sender, @TEventClass.OnReplace, [Sender]);
end;

class procedure TEventClass.OnFind(Sender: TObject);
begin
  SendEvent(Sender, @TEventClass.OnFind, [Sender]);
end;


class procedure TEventClass.OnActivate(Sender: TObject);
begin
  SendEvent(Sender, @TEventClass.OnActivate, [Sender]);
end;

class procedure TEventClass.OnDeactivate(Sender: TObject);
begin
  SendEvent(Sender, @TEventClass.OnDeactivate, [Sender]);
end;

class procedure TEventClass.OnConstrainedResize(Sender: TObject; var MinWidth, MinHeight, MaxWidth, MaxHeight: TConstraintSize);
begin
  SendEvent(Sender, @TEventClass.OnConstrainedResize, [Sender, @MinWidth, @MinHeight, @MaxWidth, @MaxHeight]);
end;


class function TEventClass.OnHelp(Command: Word; Data: PtrInt;
  var CallHelp: Boolean): Boolean;
var
  LResult: Boolean;
begin
  SendEvent(Application, @TEventClass.OnHelp, [Command, Data, Pointer(@CallHelp), Pointer(@LResult)]);
  Result := LResult;
end;

class procedure TEventClass.OnShortCut(var Msg: TLMKey; var Handled: Boolean);
begin
  SendEvent(Application, @TEventClass.OnShortCut, [Pointer(@Msg), Pointer(@Handled)]);
end;

class procedure TEventClass.OnContextPopup(Sender: TObject; MousePos: TPoint; var Handled: Boolean);
begin
  SendEvent(Sender, @TEventClass.OnContextPopup, [Sender, Pointer(@MousePos), Pointer(@Handled)]);
end;

class procedure TEventClass.OnDockDrop(Sender: TObject; Source: TDragDockObject; X, Y: Integer);
begin
  SendEvent(Sender, @TEventClass.OnDockDrop, [Sender, Source, X, Y]);
end;

class procedure TEventClass.OnDragDrop(Sender, Source: TObject; X, Y: Integer);
begin
  SendEvent(Sender, @TEventClass.OnDragDrop, [Sender, Source, X, Y]);
end;

class procedure TEventClass.OnDragOver(Sender, Source: TObject; X, Y: Integer; State: TDragState; var Accept: Boolean);
begin
  SendEvent(Sender, @TEventClass.OnDragOver, [Sender, Source, X, Y, Integer(State), Pointer(@Accept)]);
end;

class procedure TEventClass.OnEndDock(Sender, Target: TObject; X, Y: Integer);
begin
  SendEvent(Sender, @TEventClass.OnEndDock, [Sender, Target, X, Y]);
end;

class procedure TEventClass.OnGetSiteInfo(Sender: TObject; DockClient: TControl;
  var InfluenceRect: TRect; MousePos: TPoint; var CanDock: Boolean);
begin
  SendEvent(Sender, @TEventClass.OnGetSiteInfo, [Sender, DockClient, Pointer(@InfluenceRect), Pointer(@MousePos), Pointer(@CanDock)]);
end;

class procedure TEventClass.OnMouseWheelDown(Sender: TObject; Shift: TShiftState; MousePos: TPoint; var Handled: Boolean);
begin
  SendEvent(Sender, @TEventClass.OnMouseWheelDown, [Sender, PWord(@Shift)^, Pointer(@MousePos), Pointer(@Handled)]);
end;

class procedure TEventClass.OnMouseWheelUp(Sender: TObject; Shift: TShiftState; MousePos: TPoint; var Handled: Boolean);
begin
  SendEvent(Sender, @TEventClass.OnMouseWheelUp, [Sender, PWord(@Shift)^, Pointer(@MousePos), Pointer(@Handled)]);
end;

class procedure TEventClass.OnStartDock(Sender: TObject; var DragObject: TDragDockObject);
begin
  SendEvent(Sender, @TEventClass.OnStartDock, [Sender, DragObject]);
end;

class procedure TEventClass.OnUnDock(Sender: TObject; Client: TControl; NewTarget: TWinControl; var Allow: Boolean);
begin
  SendEvent(Sender, @TEventClass.OnUnDock, [Sender, Client, NewTarget, Pointer(@Allow)]);
end;

class procedure TEventClass.OnEndDrag(Sender, Target: TObject; X, Y: Integer);
begin
  SendEvent(Sender, @TEventClass.OnEndDrag, [Sender, Target, X, Y]);
end;


//class procedure TEventClass.OnGesture(Sender: TObject; const EventInfo: TGestureEventInfo; var Handled: Boolean);
//begin
//  SendEvent(Sender, @TEventClass.OnGesture, [Sender, @EventInfo, @Handled]);
//end;

//class procedure TEventClass.OnMouseActivate(Sender: TObject; Button: TMouseButton; Shift: TShiftState;
//  X, Y, HitTest: Integer; var MouseActivate: TMouseActivate);
//begin
//  SendEvent(Sender, @TEventClass.OnMouseActivate, [Sender, Ord(Button), PWord(@Shift)^, X, Y, HitTest, @MouseActivate]);
//end;


//class procedure TEventClass.ListBoxOnData(Control: TWinControl; Index: Integer; var Data: string);
//var
//  LData: PChar;
//begin
//  LData := PChar(Data);
//  SendEvent(Control, @TEventClass.ListBoxOnData, [Control,Index, @LData]);
//  Data := LData;
//end;
//
//
//class function TEventClass.ListBoxOnDataFind(Control: TWinControl; FindString: string): Integer;
//begin
//  SendEvent(Control, @TEventClass.ListBoxOnDataFind, [Control, PChar(FindString), @Result]);
//end;
//
//class procedure TEventClass.ListBoxOnDataObject(Control: TWinControl; Index: Integer; var DataObject: TObject);
//begin
//  SendEvent(Control, @TEventClass.ListBoxOnDataObject, [Control, Index, @DataObject]);
//end;

class procedure TEventClass.ListBoxOnMeasureItem(Control: TWinControl; Index: Integer; var Height: Integer);
begin
  SendEvent(Control, @TEventClass.ListBoxOnMeasureItem, [Control, Index, @Height]);
end;

class procedure TEventClass.OnChanging(Sender: TObject);
begin
  SendEvent(Sender, @TEventClass.OnChanging, [Sender]);
end;

class procedure TEventClass.UpDownOnChanging(Sender: TObject; var AllowChange: Boolean);
begin
  SendEvent(Sender, @TEventClass.UpDownOnChanging, [Sender, @AllowChange]);
end;

class procedure TEventClass.ListViewOnChanging(Sender: TObject; Item: TListItem; Change: TItemChange; var AllowChange: Boolean);
begin
  SendEvent(Sender, @TEventClass.ListViewOnChanging, [Sender, Item, Ord(Change), @AllowChange]);
end;

class procedure TEventClass.ListViewOnData(Sender: TObject; Item: TListItem);
begin
  SendEvent(Sender, @TEventClass.ListViewOnData, [Sender, Item]);
end;

class procedure TEventClass.ListViewOnDataFind(Sender: TObject; Find: TItemFind;
  const FindString: string; const FindPosition: TPoint; FindData: Pointer;
  StartIndex: Integer; Direction: TSearchDirection; Wrap: Boolean; var Index: Integer);
begin
  SendEvent(Sender, @TEventClass.ListViewOnDataFind, [Sender, Ord(Find), PChar(FindString), @FindPosition, FindData, StartIndex,
    Ord(Direction), Integer(Wrap), @Index]);
end;

class procedure TEventClass.ListViewOnEdited(Sender: TObject; Item: TListItem; var S: string);
var
  LS: PChar;
begin
  LS := PChar(S);
  SendEvent(Sender, @TEventClass.ListViewOnEdited, [Sender, Item, @LS]);
  S := LS;
end;

class procedure TEventClass.ListViewOnEditing(Sender: TObject; Item: TListItem; var AllowEdit: Boolean);
begin
  SendEvent(Sender, @TEventClass.ListViewOnEditing, [Sender, Item, @AllowEdit]);
end;

class procedure TEventClass.ListViewOnInsert(Sender: TObject; Item: TListItem);
begin
  SendEvent(Sender, @TEventClass.ListViewOnInsert, [Sender, Item]);
end;

class procedure TEventClass.ListViewOnDeletion(Sender: TObject; Item: TListItem);
begin
  SendEvent(Sender, @TEventClass.ListViewOnDeletion, [Sender, Item]);
end;


class procedure TEventClass.ListViewOnCustomDraw(Sender: TCustomListView; const ARect: TRect; var DefaultDraw: Boolean);
begin
  SendEvent(Sender, @TEventClass.ListViewOnCustomDraw, [Sender, @ARect, @DefaultDraw]);
end;

class procedure TEventClass.ListViewOnCustomDrawItem(Sender: TCustomListView;
  Item: TListItem; State: TCustomDrawState; var DefaultDraw: Boolean);
begin
  SendEvent(Sender, @TEventClass.ListViewOnCustomDrawItem, [Sender, Item, PWord(@State)^, @DefaultDraw]);
end;

class procedure TEventClass.ListViewOnCustomDrawSubItem(Sender: TCustomListView;
  Item: TListItem; SubItem: Integer; State: TCustomDrawState; var DefaultDraw: Boolean);
begin
  SendEvent(Sender, @TEventClass.ListViewOnCustomDrawSubItem, [Sender, Item, SubItem, PWord(@State)^, @DefaultDraw]);
end;

class procedure TEventClass.ListViewOnDrawItem(Sender: TCustomListView; Item: TListItem; ARect: TRect; State: TOwnerDrawState);
begin
  SendEvent(Sender, @TEventClass.ListViewOnDrawItem, [Sender, Item, @ARect, PWord(@State)^]);
end;

class procedure TEventClass.ListViewOnDataHint(Sender: TObject; StartIndex,
  EndIndex: Integer);
begin
  SendEvent(Sender, @TEventClass.ListViewOnDataHint, [Sender, StartIndex, EndIndex]);
end;



class procedure TEventClass.TreeViewOnChanging(Sender: TObject; Node: TTreeNode; var AllowChange: Boolean);
begin
  SendEvent(Sender, @TEventClass.TreeViewOnChanging, [Sender, Node, @AllowChange]);
end;

//class procedure TEventClass.TreeViewOnCancelEdit(Sender: TObject; Node: TTreeNode);
//begin
//  SendEvent(Sender, @TEventClass.TreeViewOnCancelEdit, [Sender, Node]);
//end;

class procedure TEventClass.TreeViewOnAddition(Sender: TObject; Node: TTreeNode);
begin
  SendEvent(Sender, @TEventClass.TreeViewOnAddition, [Sender, Node]);
end;

class procedure TEventClass.TreeViewOnCollapsed(Sender: TObject; Node: TTreeNode);
begin
  SendEvent(Sender, @TEventClass.TreeViewOnCollapsed, [Sender, Node]);
end;

class procedure TEventClass.TreeViewOnCollapsing(Sender: TObject; Node: TTreeNode; var AllowCollapse: Boolean);
begin
  SendEvent(Sender, @TEventClass.TreeViewOnCollapsing, [Sender, Node, @AllowCollapse]);
end;

class procedure TEventClass.TreeViewOnDeletion(Sender: TObject; Node: TTreeNode);
begin
  SendEvent(Sender, @TEventClass.TreeViewOnDeletion, [Sender, Node]);
end;

class procedure TEventClass.TreeViewOnEdited(Sender: TObject; Node: TTreeNode; var S: string);
var
  LS: PChar;
begin
  LS := PChar(S);
  SendEvent(Sender, @TEventClass.TreeViewOnEdited, [Sender, Node, @LS]);
  S := LS;
end;

class procedure TEventClass.TreeViewOnEditing(Sender: TObject; Node: TTreeNode; var AllowEdit: Boolean);
begin
  SendEvent(Sender, @TEventClass.TreeViewOnEditing, [Sender, Node, @AllowEdit]);
end;

class procedure TEventClass.TreeViewOnExpanded(Sender: TObject; Node: TTreeNode);
begin
  SendEvent(Sender, @TEventClass.TreeViewOnExpanded, [Sender, Node]);
end;

class procedure TEventClass.TreeViewOnExpanding(Sender: TObject; Node: TTreeNode; var AllowExpansion: Boolean);
begin
  SendEvent(Sender, @TEventClass.TreeViewOnExpanding, [Sender, Node, @AllowExpansion]);
end;

//class procedure TEventClass.TreeViewOnHint(Sender: TObject; const Node: TTreeNode; var Hint: string);
//var
//  LHint: PChar;
//begin
//  LHint := PChar(Hint);
//  SendEvent(Sender, @TEventClass.TreeViewOnHint, [Sender, Node, @LHint]);
//  Hint := LHint;
//end;

class procedure TEventClass.TreeViewOnCustomDraw(Sender: TCustomTreeView; const ARect: TRect; var DefaultDraw: Boolean);
begin
  SendEvent(Sender, @TEventClass.TreeViewOnCustomDraw, [Sender, @ARect, @DefaultDraw]);
end;

class procedure TEventClass.TreeViewOnCustomDrawItem(Sender: TCustomTreeView; Node: TTreeNode; State: TCustomDrawState; var DefaultDraw: Boolean);
begin
  SendEvent(Sender, @TEventClass.TreeViewOnCustomDrawItem, [Sender, Node, PWord(@State)^, @DefaultDraw]);
end;


class procedure TEventClass.MenuItemOnMeasureItem(Sender: TObject; ACanvas: TCanvas; var Width, Height: Integer);
begin
  SendEvent(Sender, @TEventClass.MenuItemOnMeasureItem, [Sender, ACanvas, @Width, @Height]);
end;

class procedure TEventClass.PageControlOnChanging(Sender: TObject; var AllowChange: Boolean);
begin
  SendEvent(Sender, @TEventClass.PageControlOnChanging, [Sender, @AllowChange]);
end;


//class procedure TEventClass.OnMessage(var Msg: TMsg; var Handled: Boolean);
//begin
//  SendEvent(Application, @TEventClass.OnMessage, [Pointer(@Msg), Pointer(@Handled)]);
//end;


//----------------------- TListView

class procedure TEventClass.ListViewOnChange(Sender: TObject; AItem: TListItem; Change: TItemChange);
begin
  SendEvent(Sender, @TEventClass.ListViewOnChange, [Sender, AItem, Ord(Change)]);
end;

class procedure TEventClass.ListViewOnColumnClick(Sender: TObject; Column: TListColumn);
begin
  SendEvent(Sender, @TEventClass.ListViewOnColumnClick, [Sender, Column]);
end;

class procedure TEventClass.ListViewOnColumnRightClick(Sender: TObject; Column: TListColumn; Point: TPoint);
begin
  SendEvent(Sender, @TEventClass.ListViewOnColumnRightClick, [Sender, Column, Point.X, Point.Y]);
end;

class procedure TEventClass.ListViewOnGetImageIndex(Sender: TObject; Item: TListItem);
begin
  SendEvent(Sender, @TEventClass.ListViewOnGetImageIndex, [Sender, Item]);
end;

class procedure TEventClass.ListViewOnSelectItem(Sender: TObject; Item: TListItem; Selected: Boolean);
begin
  SendEvent(Sender, @TEventClass.ListViewOnSelectItem, [Sender, Item, Selected]);
end;

class procedure TEventClass.ListViewOnItemChecked(Sender: TObject; Item: TListItem);
begin
  SendEvent(Sender, @TEventClass.ListViewOnItemChecked, [Sender, Item]);
end;

class procedure TEventClass.ListViewOnCompare(Sender: TObject; Item1, Item2: TListItem; Data: Integer; var Compare: Integer);
begin
  SendEvent(Sender, @TEventClass.ListViewOnCompare, [Sender, Item1, Item2, Data, @Compare]);
end;

class procedure TEventClass.ListViewOnAdvancedCustomDraw(Sender: TCustomListView;
  const ARect: TRect; Stage: TCustomDrawStage; var DefaultDraw: Boolean);
begin
  SendEvent(Sender, @TEventClass.ListViewOnAdvancedCustomDraw, [
    Sender, @ARect, Ord(Stage), @DefaultDraw
  ]);
end;

class procedure TEventClass.ListViewOnAdvancedCustomDrawItem(
  Sender: TCustomListView; Item: TListItem; State: TCustomDrawState;
  Stage: TCustomDrawStage; var DefaultDraw: Boolean);
begin
  SendEvent(Sender, @TEventClass.ListViewOnAdvancedCustomDrawItem, [
    Sender, Item, PWord(@State)^, Ord(Stage), @DefaultDraw
  ]);
end;

class procedure TEventClass.ListViewOnAdvancedCustomDrawSubItem(
  Sender: TCustomListView; Item: TListItem; SubItem: Integer;
  State: TCustomDrawState; Stage: TCustomDrawStage; var DefaultDraw: Boolean);
begin
  SendEvent(Sender, @TEventClass.ListViewOnAdvancedCustomDrawSubItem, [
    Sender, Item, SubItem, PWord(@State)^, Ord(Stage), @DefaultDraw
  ]);
end;

//----------------------------------------- TPageControl
class procedure TEventClass.PageControlOnGetImageIndex(Sender: TObject; TabIndex: Integer; var ImageIndex: Integer);
begin
  SendEvent(Sender, @TEventClass.PageControlOnGetImageIndex, [Sender, TabIndex, @ImageIndex]);
end;
class procedure TEventClass.MenuOnChange(Sender: TObject; Source: TMenuItem;
  Rebuild: Boolean);
begin
  SendEvent(Sender, @TEventClass.MenuOnChange, [Sender, Source, Rebuild]);
end;

class procedure TEventClass.OnTitleChange(Sender: TObject; const Text: string);
begin
  SendEvent(Sender, @TEventClass.OnTitleChange, [Sender, Text]);
end;

class procedure TEventClass.OnJSExternal(Sender: TObject; const Afunc: string;
  const AArgs: WideString; var ARetval: WideString);
var
  LRet: PChar;
begin
  LRet := PChar(ARetval);
  SendEvent(Sender, @TEventClass.OnJSExternal, [Sender, string(Afunc), string(AArgs), @LRet]);
  ARetval := WideString(LRet);
end;


// TaskDialog
class procedure TEventClass.TaskDialogOnButtonClicked(Sender: TObject; ModalResult: TModalResult; var CanClose: Boolean);
begin
  SendEvent(Sender, @TEventClass.TaskDialogOnButtonClicked, [Sender, ModalResult, @CanClose]);
end;

class procedure TEventClass.OnAlignPosition(Sender: TWinControl;
  Control: TControl; var NewLeft, NewTop, NewWidth, NewHeight: Integer;
  var AlignRect: TRect; AlignInfo: TAlignInfo);
begin
  SendEvent(Sender, @TEventClass.OnAlignPosition, [Sender, Control, @NewLeft, @NewTop, @NewWidth, @NewHeight, @AlignRect, @AlignInfo]);
end;

//class procedure TEventClass.TaskDialogOnDialogConstructed(Sender: TObject);
//begin
//  SendEvent(Sender, @TEventClass.TaskDialogOnDialogConstructed, [Sender]);
//end;
//
//class procedure TEventClass.TaskDialogOnDialogCreated(Sender: TObject);
//begin
//  SendEvent(Sender, @TEventClass.TaskDialogOnDialogCreated, [Sender]);
//end;
//
//class procedure TEventClass.TaskDialogOnDialogDestroyed(Sender: TObject);
//begin
//  SendEvent(Sender, @TEventClass.TaskDialogOnDialogDestroyed, [Sender]);
//end;
//
//class procedure TEventClass.TaskDialogOnExpanded(Sender: TObject);
//begin
//  SendEvent(Sender, @TEventClass.TaskDialogOnExpanded, [Sender]);
//end;
//
//class procedure TEventClass.TaskDialogOnHyperlinkClicked(Sender: TObject);
//begin
//  SendEvent(Sender, @TEventClass.TaskDialogOnHyperlinkClicked, [Sender]);
//end;
//
//class procedure TEventClass.TaskDialogOnNavigated(Sender: TObject);
//begin
//  SendEvent(Sender, @TEventClass.TaskDialogOnNavigated, [Sender]);
//end;
//
//class procedure TEventClass.TaskDialogOnRadioButtonClicked(Sender: TObject);
//begin
//  SendEvent(Sender, @TEventClass.TaskDialogOnRadioButtonClicked, [Sender]);
//end;
//
//class procedure TEventClass.TaskDialogOnTimer(Sender: TObject; TickCount: Cardinal; var Reset: Boolean);
//begin
//  SendEvent(Sender, @TEventClass.TaskDialogOnTimer, [Sender, TickCount, @Reset]);
//end;
//
//class procedure TEventClass.TaskDialogOnVerificationClicked(Sender: TObject);
//begin
//  SendEvent(Sender, @TEventClass.TaskDialogOnVerificationClicked, [Sender]);
//end;

// -------------- end TaskDialog -----------------------

//class function TEventClass.OnAlignInsertBefore(Sender: TWinControl; C1, C2: TControl): Boolean;
//begin
//  SendEvent(Sender, @TEventClass.OnAlignInsertBefore, [Sender, C1, C2, @Result]);
//end;

class procedure TEventClass.OnDropDown(Sender: TObject);
begin
  SendEvent(Sender, @TEventClass.OnDropDown, [Sender]);
end;

class procedure TEventClass.OnSelect(Sender: TObject);
begin
  SendEvent(Sender, @TEventClass.OnSelect, [Sender]);
end;

class procedure TEventClass.OnBeginEdit(Sender: TObject);
begin
  SendEvent(Sender, @TEventClass.OnBeginEdit, [Sender]);
end;

class procedure TEventClass.OnEndEdit(Sender: TObject);
begin
  SendEvent(Sender, @TEventClass.OnEndEdit, [Sender]);
end;



class procedure TEventClass.OnClose(Sender: TObject);
begin
  SendEvent(Sender, @TEventClass.OnClose, [Sender]);
end;

class procedure TEventClass.FormOnCloseQuery(Sender: TObject;
  var CanClose: Boolean);
begin
  SendEvent(Sender, @TEventClass.FormOnCloseQuery, [Sender, @CanClose]);
end;

class procedure TEventClass.FormOnDropFiles(Sender: TObject;
  const AFileNames: array of string);
var
  LLen: Integer;
begin
  LLen := Length(AFileNames);
  if LLen > 0 then
    SendEvent(Sender, @TEventClass.FormOnDropFiles, [Sender, @AFileNames[0], LLen]);
end;

class procedure TEventClass.OnDblClick(Sender: TObject);
begin
  SendEvent(Sender, @TEventClass.OnDblClick, [Sender]);
end;

class procedure TEventClass.OnEnter(Sender: TObject);
begin
  SendEvent(Sender, @TEventClass.OnEnter, [Sender]);
end;

class procedure TEventClass.OnException(Sender: TObject; E: Exception);
begin
  SendEvent(Sender, @TEventClass.OnException, [Sender, E]);
end;

class procedure TEventClass.OnExecute(Sender: TObject);
begin
  SendEvent(Sender, @TEventClass.OnExecute, [Sender]);
end;

class procedure TEventClass.OnExit(Sender: TObject);
begin
  SendEvent(Sender, @TEventClass.OnExit, [Sender]);
end;

class procedure TEventClass.OnHide(Sender: TObject);
begin
  SendEvent(Sender, @TEventClass.OnHide, [Sender]);
end;


// grid
class procedure TEventClass.OnColumnMoved(Sender: TObject; IsColumn: Boolean; FromIndex, ToIndex: Longint);
begin
  if IsColumn then
    SendEvent(Sender, @TEventClass.OnColumnMoved, [Sender, FromIndex, ToIndex]);
end;

class procedure TEventClass.OnDrawCell(Sender: TObject; ACol, ARow: Longint; ARect: TRect; State: TGridDrawState);
begin
  SendEvent(Sender, @TEventClass.OnDrawCell, [Sender, ACol, ARow, Pointer(@ARect), PWord(@State)^]);
end;

class procedure TEventClass.OnFixedCellClick(Sender: TObject; ACol, ARow: Integer);
begin
  SendEvent(Sender, @TEventClass.OnFixedCellClick, [Sender, ACol, ARow]);
end;

class procedure TEventClass.OnGetEditMask(Sender: TObject; ACol, ARow: Integer; var Value: string);
var
  LS: PChar;
begin
  LS := PChar(Value);
  SendEvent(Sender, @TEventClass.OnGetEditMask, [Sender, ACol, ARow, Pointer(@LS)]);
  Value := LS;
end;

class procedure TEventClass.OnGetEditText(Sender: TObject; ACol, ARow: Integer; var Value: string);
var
  LS: PChar;
begin
  LS := PChar(Value);
  SendEvent(Sender, @TEventClass.OnGetEditText, [Sender, ACol, ARow, Pointer(@LS)]);
  Value := LS;
end;

class procedure TEventClass.OnRowMoved(Sender: TObject; IsColumn: Boolean; FromIndex, ToIndex: Integer);
begin
  if not IsColumn then
    SendEvent(Sender, @TEventClass.OnRowMoved, [Sender, FromIndex, ToIndex]);
end;

class procedure TEventClass.OnSelectCell(Sender: TObject; ACol, ARow: Integer; var CanSelect: Boolean);
begin
  SendEvent(Sender, @TEventClass.OnSelectCell, [Sender, ACol, ARow, Pointer(@CanSelect)]);
end;

class procedure TEventClass.OnSetEditText(Sender: TObject; ACol, ARow: Integer; const Value: string);
begin
  SendEvent(Sender, @TEventClass.OnSetEditText, [Sender, ACol, ARow, PChar(Value)]);
end;

class procedure TEventClass.OnTopLeftChanged(Sender: TObject);
begin
  SendEvent(Sender, @TEventClass.OnTopLeftChanged, [Sender]);
end;


// headercontrol
//class procedure TEventClass.OnDrawSection(HeaderControl: THeaderControl; Section: THeaderSection; const Rect: TRect; Pressed: Boolean);
//begin
//  SendEvent(HeaderControl, @TEventClass.OnDrawSection, [HeaderControl, Section, Pointer(@Rect), Pressed]);
//end;

//class procedure TEventClass.OnSectionCheck(HeaderControl: TCustomHeaderControl; Section: THeaderSection);
//begin
//  SendEvent(HeaderControl, @TEventClass.OnSectionCheck, [HeaderControl, Section]);
//end;

class procedure TEventClass.OnSectionClick(HeaderControl: TCustomHeaderControl; Section: THeaderSection);
begin
  SendEvent(HeaderControl, @TEventClass.OnSectionClick, [HeaderControl, Section]);
end;

class procedure TEventClass.OnSectionDrag(Sender: TObject; FromSection, ToSection: THeaderSection; var AllowDrag: Boolean);
begin
  SendEvent(Sender, @TEventClass.OnSectionDrag, [Sender, FromSection, ToSection, Pointer(@AllowDrag)]);
end;

class procedure TEventClass.OnSectionEndDrag(Sender: TObject);
begin
  SendEvent(Sender, @TEventClass.OnSectionEndDrag, [Sender]);
end;

class procedure TEventClass.OnSectionResize(HeaderControl: TCustomHeaderControl; Section: THeaderSection);
begin
  SendEvent(HeaderControl, @TEventClass.OnSectionResize, [HeaderControl, Section]);
end;

class procedure TEventClass.OnSectionTrack(HeaderControl: TCustomHeaderControl; Section: THeaderSection; Width: Integer; State: TSectionTrackState);
begin
  SendEvent(HeaderControl, @TEventClass.OnSectionTrack, [HeaderControl, Section, Width, Integer(State)]);
end;


class procedure TEventClass.OnDestroy(Sender: TObject);
begin
  SendEvent(Sender, @TEventClass.OnDestroy, [Sender]);
end;

class procedure TEventClass.OnHint(Sender: TObject);
begin
   SendEvent(Sender, @TEventClass.OnHint, [Sender]);
end;
class procedure TEventClass.OnKeyDown(Sender: TObject; var Key: Word;
  Shift: TShiftState);
begin
  SendEvent(Sender, @TEventClass.OnKeyDown, [Sender, @Key, PWord(@Shift)^]);
end;

class procedure TEventClass.OnKeyPress(Sender: TObject; var Key: Char);
var
  LKey: Word;
begin
  // 这里要修复下
  LKey := Ord(Key);
  SendEvent(Sender, @TEventClass.OnKeyPress, [Sender, @LKey]);
  Key := Char(LKey);
end;

class procedure TEventClass.OnKeyUp(Sender: TObject; var Key: Word;
  Shift: TShiftState);
begin
  SendEvent(Sender, @TEventClass.OnKeyUp, [Sender, @Key, PWord(@Shift)^]);
end;

class procedure TEventClass.OnMinimize(Sender: TObject);
begin
  SendEvent(Sender, @TEventClass.OnMinimize, [Sender]);
end;

class procedure TEventClass.OnMouseDown(Sender: TObject; Button: TMouseButton;
  Shift: TShiftState; X, Y: Integer);
begin
  SendEvent(Sender, @TEventClass.OnMouseDown, [Sender, Ord(Button), PWord(@Shift)^, X, Y]);
end;

class procedure TEventClass.OnMouseEnter(Sender: TObject);
begin
  SendEvent(Sender, @TEventClass.OnMouseEnter, [Sender]);
end;

class procedure TEventClass.OnMouseLeave(Sender: TObject);
begin
  SendEvent(Sender, @TEventClass.OnMouseLeave, [Sender]);
end;

class procedure TEventClass.ListBoxOnDrawItem(Control: TWinControl; Index: Integer;
  ARect: TRect; State: TOwnerDrawState);
begin
  SendEvent(Control, @TEventClass.ListBoxOnDrawItem, [Control, Index, @ARect, PWord(@State)^]);
end;

class procedure TEventClass.OnMouseMove(Sender: TObject; Shift: TShiftState; X,
  Y: Integer);
begin
  SendEvent(Sender, @TEventClass.OnMouseMove, [Sender, PWord(@Shift)^, X, Y]);
end;

class procedure TEventClass.OnMouseUp(Sender: TObject; Button: TMouseButton;
  Shift: TShiftState; X, Y: Integer);
begin
  SendEvent(Sender, @TEventClass.OnMouseUp, [Sender, Ord(Button), PWord(@Shift)^, X, Y]);
end;

class procedure TEventClass.OnMouseWheel(Sender: TObject; Shift: TShiftState;
  WheelDelta: Integer; MousePos: TPoint; var Handled: Boolean);
begin
  SendEvent(Sender, @TEventClass.OnMouseWheel, [Sender, PWord(@Shift)^, WheelDelta, MousePos.X, MousePos.Y, @Handled]);
end;

class procedure TEventClass.OnLinkClick(Sender: TObject; const Link: string;
  LinkType: TSysLinkType);
begin
  SendEvent(Sender, @TEventClass.OnLinkClick, [Sender, Link, Ord(LinkType)]);
end;

class procedure TEventClass.OnPaint(Sender: TObject);
begin
  SendEvent(Sender, @TEventClass.OnPaint, [Sender]);
end;

class procedure TEventClass.OnPopup(Sender: TObject);
begin
  SendEvent(Sender, @TEventClass.OnPopup, [Sender]);
end;

class procedure TEventClass.OnResize(Sender: TObject);
begin
  SendEvent(Sender, @TEventClass.OnResize, [Sender]);
end;

class procedure TEventClass.OnRestore(Sender: TObject);
begin
  SendEvent(Sender, @TEventClass.OnRestore, [Sender]);
end;

class procedure TEventClass.OnShow(Sender: TObject);
begin
  SendEvent(Sender, @TEventClass.OnShow, [Sender]);
end;

class procedure TEventClass.OnTimer(Sender: TObject);
begin
  SendEvent(Sender, @TEventClass.OnTimer, [Sender]);
end;

class procedure TEventClass.OnUpdate(Sender: TObject);
begin
  SendEvent(Sender, @TEventClass.OnUpdate, [Sender]);
end;


class procedure TEventClass.SendEvent(Sender: TObject; AEvent: Pointer; AArgs: array of const);


  procedure SendEventSrc(EventId: NativeUInt; AArgs: array of const);
  var
    LParams: array[0..CALL_MAX_PARAM-1] of Pointer;
    LArgLen: Integer;
    LV: TVarRec;
    I: Integer;
  begin
    if Assigned(GEventCallbackPtr) and (EventId > 0) then
    begin
      LArgLen := Length(AArgs);
      if LArgLen <= Length(LParams) then
      begin
        for I := 0 to LArgLen - 1 do
        begin
          LV := AArgs[I];
          case LV.VType of
            vtInteger       : LParams[I] := Pointer(LV.VInteger);
            vtBoolean       : LParams[I] := Pointer(Byte(LV.VBoolean));
            vtChar          : LParams[I] := Pointer(Ord(LV.VChar));
            vtExtended      : LParams[I] := LV.VExtended;

            vtString        : LParams[I] := {$IFDEF MSWINDOWS}LV.VString{$ELSE}LV.VAnsiString{$ENDIF};

            vtPointer       : LParams[I] := LV.VPointer;
            vtPChar         : LParams[I] := LV.VPChar;
            vtObject        : LParams[I] := LV.VObject;
            vtClass         : LParams[I] := LV.VClass;
            vtWideChar      : LParams[I] := Pointer(Ord(LV.VWideChar));
            vtPWideChar     : LParams[I] := LV.VPWideChar;
            vtAnsiString    : LParams[I] := LV.VAnsiString;
  //          vtCurrency      = 12;
  //          vtVariant       = 13;
            vtInterface     : LParams[I] := LV.VInterface;
            vtWideString    : LParams[I] := LV.VWideString;
            vtInt64         : LParams[I] := LV.VInt64;
            vtUnicodeString : LParams[I] := LV.VUnicodeString;
          end;
        end;
        GEventCallbackPtr(EventId, @LParams[0], LArgLen);
      end;
    end;
  end;

var
  LEventId: NativeUInt;
begin
  if FEvents.TryGetData(CreateEventKey(Sender, AEvent), LEventId) then
    SendEventSrc(LEventId, AArgs)
  else writeln('can''t found id, sender:', sender.ToString, ', event:', Cardinal(AEvent));
end;

// ---------------------- TTreeView

class procedure TEventClass.TreeViewOnAdvancedCustomDraw(
  Sender: TCustomTreeView; const ARect: TRect; Stage: TCustomDrawStage;
  var DefaultDraw: Boolean);
begin
  SendEvent(Sender, @TEventClass.TreeViewOnAdvancedCustomDraw,
    [Sender, @ARect, Ord(Stage), DefaultDraw]);
end;

class procedure TEventClass.TreeViewOnAdvancedCustomDrawItem(
  Sender: TCustomTreeView; Node: TTreeNode; State: TCustomDrawState;
  Stage: TCustomDrawStage; var PaintImages, DefaultDraw: Boolean);
begin
  SendEvent(Sender, @TEventClass.TreeViewOnAdvancedCustomDrawItem,
    [Sender, Node, PWord(@State)^, Ord(Stage), @PaintImages, DefaultDraw]);
end;

class procedure TEventClass.TreeViewOnChange(Sender: TObject; ANode: TTreeNode);
begin
  SendEvent(Sender, @TEventClass.TreeViewOnChange, [Sender, ANode]);
end;

class procedure TEventClass.TreeViewOnGetImageIndex(Sender: TObject; Node: TTreeNode);
begin
  SendEvent(Sender, @TEventClass.TreeViewOnGetImageIndex, [Sender, Node]);
end;

class procedure TEventClass.TreeViewOnGetSelectedIndex(Sender: TObject; Node: TTreeNode);
begin
  SendEvent(Sender, @TEventClass.TreeViewOnGetSelectedIndex, [Sender, Node]);
end;

class procedure TEventClass.TreeViewOnCompare(Sender: TObject; Node1, Node2: TTreeNode; var Compare: Integer);
begin
  SendEvent(Sender, @TEventClass.TreeViewOnCompare, [Sender, Node1, Node2, 0, @Compare]);
end;
//----------- TToolBar
class procedure TEventClass.ToolBarOnAdvancedCustomDraw(Sender: TToolBar;
  const ARect: TRect; Stage: TCustomDrawStage; var DefaultDraw: Boolean);
begin
  SendEvent(Sender, @TEventClass.ToolBarOnAdvancedCustomDraw,
    [Sender, @ARect, Ord(Stage), @DefaultDraw]);
end;

//class procedure TEventClass.ToolBarOnAdvancedCustomDrawButton(Sender: TToolBar;
//  Button: TToolButton; State: TCustomDrawState; Stage: TCustomDrawStage;
//  var Flags: TTBCustomDrawFlags; var DefaultDraw: Boolean);
//begin
//  SendEvent(Sender, @TEventClass.ToolBarOnAdvancedCustomDrawButton,
//    [Sender, Button, PWord(@State)^, Ord(Stage), PWord(@Flags), @DefaultDraw]);
//end;


class procedure TEventClass.UpDownOnClick(Sender: TObject; Button: TUDBtnType);
begin
  SendEvent(Sender, @TEventClass.UpDownOnClick, [Sender, Ord(Button)]);
end;

// 用户定义事件实现引入
{$I UserDefineEventsImplement.inc}

{ TMessageEventClass }

class constructor TMessageEventClass.Create;
begin
  FMsgEvents := TEventList.Create;
end;

class destructor TMessageEventClass.Destroy;
begin
  FMsgEvents.Free;
end;

class procedure TMessageEventClass.OnWndProc(Sender: TObject; var TheMessage: TLMessage);
var
  LId: NativeUInt;
begin
  if Assigned(GMessageCallbackPtr) then
  begin
    if FMsgEvents.TryGetData(NativeUInt(Sender), LId) then
      GMessageCallbackPtr(LId, @TheMessage);
  end;
end;

class procedure TMessageEventClass.Remove(AObj: TObject);
begin
  FMsgEvents.Remove(NativeUInt(AObj));
end;

class procedure TMessageEventClass.Add(AObj: TObject; AId: NativeUInt);
begin
  FMsgEvents.AddOrSetData(NativeUInt(AObj), AId);
end;


end.
