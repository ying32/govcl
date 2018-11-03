unit uEventCallback;

interface

uses
  Winapi.Windows,
  Winapi.Messages,
  Vcl.Controls,
  Vcl.Forms,
  Vcl.ComCtrls,
  Vcl.Grids,
  Vcl.Menus,
  Vcl.ExtCtrls,
  Vcl.Graphics,
  System.Classes,
  System.SysUtils,
  System.Types,
  Vcl.JumpList,
  System.Win.TaskbarCore,
  Vcl.Taskbar,
  System.Generics.Collections;

var
  GEventCallbackPtr: function(f: NativeUInt; args: Pointer; argcout: NativeInt): Pointer; stdcall;
  GMessageCallbackPtr: function(f: NativeUInt; msg, handled: Pointer): Pointer; stdcall;

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
              geMouseWheelUp, geStartDock, geUnDock, geMessage, geEndDrag,

              geColumnMoved, geDrawCell, geFixedCellClick, geGetEditMask,
              geGetEditText, geRowMoved, geSelectCell, geSetEditText, geTopLeftChanged,

              geDrawSection, geSectionResize, geSectionTrack, geSectionDrag, geSectionEndDrag, geSectionCheck,
              geSectionClick,

              geGesture, geMouseActivate, geListBoxData, geListBoxDataFind, geListBoxDataObject, geListBoxMeasureItem,
              geChanging, geUpDownChanging,
              geListViewChanging, geListViewData, geListViewDataFind, geListViewEdited, geListViewEditing, geListViewInsert,
              geListViewDeletion, geListViewCustomDraw, geListViewCustomDrawItem, geListViewCustomDrawSubItem, geListViewDrawItem,

              geTreeViewChanging, geTreeViewCancelEdit, geTreeViewAddition, geTreeViewCollapsed, geTreeViewCollapsing,
              geTreeViewDeletion, geTreeViewEdited, geTreeViewEditing, geTreeViewExpanded, geTreeViewExpanding, geTreeViewHint,
              geTreeViewCustomDraw, geTreeViewCustomDrawItem,

              geMenuItemMeasureItem, gePageControlChanging,

              geJumpListItemDeleted, geJumpListListUpdateError, geJumpListItemsLoaded,
              geTaskbarThumbPreviewRequest, geTaskbarWindowPreviewItemRequest, geTaskbarThumbButtonClick,
              geStyleChanged
              );

  TEventKey = packed record
    Sender: TObject;
    Event: TGoEvent;
  public
    constructor Create(ASender: TObject; AEvent: TGoEvent);
  end;

  TEventList = TDictionary<TEventKey, NativeUInt>;

  TEventClass = class
  private class var
    FEvents: TEventList;
    class constructor Create;
    class destructor Destroy;

    class procedure SendEvent(Sender: TObject; AEvent: TGoEvent; AArgs: array of const);
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
    class procedure TreeViewOnCompare(Sender: TObject; Node1, Node2: TTreeNode; Data: Integer; var Compare: Integer);
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
    class procedure ToolBarOnAdvancedCustomDrawButton(Sender: TToolBar;
      Button: TToolButton; State: TCustomDrawState; Stage: TCustomDrawStage;
      var Flags: TTBCustomDrawFlags; var DefaultDraw: Boolean);


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
    class procedure OnConstrainedResize(Sender: TObject; var MinWidth, MinHeight, MaxWidth, MaxHeight: Integer);

    // new
    class function OnHelp(Command: Word; Data: THelpEventData; var CallHelp: Boolean): Boolean;
    class procedure OnShortCut(var Msg: TWMKey; var Handled: Boolean);
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
    class procedure OnMessage(var Msg: TMsg; var Handled: Boolean);
    class procedure OnEndDrag(Sender, Target: TObject; X, Y: Integer);




    class procedure OnGesture(Sender: TObject; const EventInfo: TGestureEventInfo; var Handled: Boolean);
    class procedure OnMouseActivate(Sender: TObject; Button: TMouseButton; Shift: TShiftState; X, Y, HitTest: Integer; var MouseActivate: TMouseActivate);
    class procedure ListBoxOnData(Control: TWinControl; Index: Integer; var Data: string);
    class function ListBoxOnDataFind(Control: TWinControl; FindString: string): Integer;
    class procedure ListBoxOnDataObject(Control: TWinControl; Index: Integer; var DataObject: TObject);
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


    class procedure TreeViewOnChanging(Sender: TObject; Node: TTreeNode; var AllowChange: Boolean);
    class procedure TreeViewOnCancelEdit(Sender: TObject; Node: TTreeNode);
    class procedure TreeViewOnAddition(Sender: TObject; Node: TTreeNode);
    class procedure TreeViewOnCollapsed(Sender: TObject; Node: TTreeNode);
    class procedure TreeViewOnCollapsing(Sender: TObject; Node: TTreeNode; var AllowCollapse: Boolean);
    class procedure TreeViewOnDeletion(Sender: TObject; Node: TTreeNode);
    class procedure TreeViewOnEdited(Sender: TObject; Node: TTreeNode; var S: string);
    class procedure TreeViewOnEditing(Sender: TObject; Node: TTreeNode; var AllowEdit: Boolean);
    class procedure TreeViewOnExpanded(Sender: TObject; Node: TTreeNode);
    class procedure TreeViewOnExpanding(Sender: TObject; Node: TTreeNode; var AllowExpansion: Boolean);
    class procedure TreeViewOnHint(Sender: TObject; const Node: TTreeNode; var Hint: string);
    class procedure TreeViewOnCustomDraw(Sender: TCustomTreeView; const ARect: TRect; var DefaultDraw: Boolean);
    class procedure TreeViewOnCustomDrawItem(Sender: TCustomTreeView; Node: TTreeNode; State: TCustomDrawState; var DefaultDraw: Boolean);

    class procedure MenuItemOnMeasureItem(Sender: TObject; ACanvas: TCanvas; var Width, Height: Integer);
    class procedure PageControlOnChanging(Sender: TObject; var AllowChange: Boolean);

    class procedure JumpListOnItemDeleted(Sender: TObject; const Item: TJumpListItem; const CategoryName: string; FromTasks: Boolean);
    class procedure JumpListOnListUpdateError(Sender: TObject; WinErrorCode: Cardinal; const ErrorDescription: string; var Handled: Boolean);
    class procedure JumpListOnItemsLoaded(Sender: TObject);

    class procedure TaskbarOnThumbPreviewRequest(Sender: TObject; APreviewHeight, APreviewWidth: Integer; PreviewBitmap: TBitmap);
    class procedure TaskbarOnWindowPreviewItemRequest(Sender: TObject; var Position: TPoint; PreviewBitmap: TBitmap);
    class procedure TaskbarOnThumbButtonClick(Sender: TObject; AButtonID: Integer);
    // grid

    class procedure OnColumnMoved(Sender: TObject; FromIndex, ToIndex: Longint);
    class procedure OnDrawCell(Sender: TObject; ACol, ARow: Longint; ARect: TRect; State: TGridDrawState);
    class procedure OnFixedCellClick(Sender: TObject; ACol, ARow: Integer);
    class procedure OnGetEditMask(Sender: TObject; ACol, ARow: Integer; var Value: string);
    class procedure OnGetEditText(Sender: TObject; ACol, ARow: Integer; var Value: string);
    class procedure OnRowMoved(Sender: TObject; FromIndex, ToIndex: Integer);
    class procedure OnSelectCell(Sender: TObject; ACol, ARow: Integer; var CanSelect: Boolean);
    class procedure OnSetEditText(Sender: TObject; ACol, ARow: Integer; const Value: string);
    class procedure OnTopLeftChanged(Sender: TObject);



    // headercontrol
    class procedure OnDrawSection(HeaderControl: THeaderControl; Section: THeaderSection; const Rect: TRect; Pressed: Boolean);
    class procedure OnSectionCheck(HeaderControl: TCustomHeaderControl; Section: THeaderSection);
    class procedure OnSectionClick(HeaderControl: THeaderControl; Section: THeaderSection);
    class procedure OnSectionDrag(Sender: TObject; FromSection, ToSection: THeaderSection; var AllowDrag: Boolean);
    class procedure OnSectionEndDrag(Sender: TObject);
    class procedure OnSectionResize(HeaderControl: THeaderControl; Section: THeaderSection);
    class procedure OnSectionTrack(HeaderControl: THeaderControl; Section: THeaderSection; Width: Integer; State: TSectionTrackState);



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
    class procedure MenuItemOnDrawItem(Sender: TObject; ACanvas: TCanvas;
          ARect: TRect; Selected: Boolean);


    class procedure OnLinkClick(Sender: TObject; const Link: string; LinkType: TSysLinkType);

    class procedure OnStyleChanged(Sender: TObject);

    class procedure MenuOnChange(Sender: TObject; Source: TMenuItem; Rebuild: Boolean);
    class procedure Add(AObj: TObject; AEvent: TGoEvent; AId: NativeUInt);
    class procedure AddClick(Sender: TObject; AId: NativeUInt);
    class procedure Remove(AObj: TObject; AEvent: TGoEvent);
  end;

  // 窗口消息的，不与之前的事件混在一起。 
  TMessageEventList =  TDictionary<TObject, NativeUInt>;

  TMessageEventClass = class
  private class var
    FMsgEvents: TMessageEventList;
    class constructor Create;
    class destructor Destroy;
  public
    class procedure Add(AObj: TObject; AId: NativeUInt);
    class procedure Remove(AObj: TObject);
    class procedure OnWndProc(Sender: TObject; var AMsg: TMessage; var AHandled: Boolean);
  end;

implementation

{ TEventClass }

class procedure TEventClass.Add(AObj: TObject; AEvent: TGoEvent; AId: NativeUInt);
begin
  FEvents.AddOrSetValue(TEventKey.Create(AObj, AEvent), AId);
end;

class procedure TEventClass.AddClick(Sender: TObject; AId: NativeUInt);
begin
  Add(Sender, geClick, AId);
end;

class constructor TEventClass.Create;
begin
  FEvents := TEventList.Create;
end;

class destructor TEventClass.Destroy;
begin
  FreeAndNil(FEvents);
end;

class procedure TEventClass.OnBalloonClick(Sender: TObject);
begin
  SendEvent(Sender, geBalloonClick, [Sender]);
end;

class procedure TEventClass.OnChange(Sender: TObject);
begin
  SendEvent(Sender, geChange, [Sender]);
end;

class procedure TEventClass.OnClick(Sender: TObject);
begin
  SendEvent(Sender, geClick, [Sender]);
end;

class procedure TEventClass.OnClickCheck(Sender: TObject);
begin
  SendEvent(Sender, geClickCheck, [Sender]);
end;

class procedure TEventClass.FormOnClose(Sender: TObject; var Action: TCloseAction);
begin
  SendEvent(Sender, geFormClose, [Sender, @Action]);
end;

//----------------------- TListView

class procedure TEventClass.ListViewOnChange(Sender: TObject; AItem: TListItem; Change: TItemChange);
begin
  SendEvent(Sender, geListViewChange, [Sender, AItem, Ord(Change)]);
end;

class procedure TEventClass.ListViewOnColumnClick(Sender: TObject; Column: TListColumn);
begin
  SendEvent(Sender, geListViewColumnClick, [Sender, Column]);
end;

class procedure TEventClass.ListViewOnColumnRightClick(Sender: TObject; Column: TListColumn; Point: TPoint);
begin
  SendEvent(Sender, geListViewColumnRightClick, [Sender, Column, Point.X, Point.Y]);
end;

class procedure TEventClass.ListViewOnGetImageIndex(Sender: TObject; Item: TListItem);
begin
  SendEvent(Sender, geListViewGetImageIndex, [Sender, Item]);
end;

class procedure TEventClass.ListViewOnSelectItem(Sender: TObject; Item: TListItem; Selected: Boolean);
begin
  SendEvent(Sender, geListViewSelectItem, [Sender, Item, Selected]);
end;

class procedure TEventClass.ListViewOnItemChecked(Sender: TObject; Item: TListItem);
begin
  SendEvent(Sender, geListViewItemChecked, [Sender, Item]);
end;

class procedure TEventClass.ListViewOnCompare(Sender: TObject; Item1, Item2: TListItem; Data: Integer; var Compare: Integer);
begin
  SendEvent(Sender, geListViewCompare, [Sender, Item1, Item2, Data, @Compare]);
end;

class procedure TEventClass.ListViewOnAdvancedCustomDraw(Sender: TCustomListView;
  const ARect: TRect; Stage: TCustomDrawStage; var DefaultDraw: Boolean);
begin
  SendEvent(Sender, geListViewAdvancedCustomDraw, [
    Sender, @ARect, Ord(Stage), @DefaultDraw
  ]);
end;

class procedure TEventClass.ListViewOnAdvancedCustomDrawItem(
  Sender: TCustomListView; Item: TListItem; State: TCustomDrawState;
  Stage: TCustomDrawStage; var DefaultDraw: Boolean);
begin
  SendEvent(Sender, geListViewAdvancedCustomDrawItem, [
    Sender, Item, PWord(@State)^, Ord(Stage), @DefaultDraw
  ]);
end;

class procedure TEventClass.ListViewOnAdvancedCustomDrawSubItem(
  Sender: TCustomListView; Item: TListItem; SubItem: Integer;
  State: TCustomDrawState; Stage: TCustomDrawStage; var DefaultDraw: Boolean);
begin
  SendEvent(Sender, geListViewAdvancedCustomDrawSubItem, [
    Sender, Item, SubItem, PWord(@State)^, Ord(Stage), @DefaultDraw
  ]);
end;

//----------------------------------------- TPageControl

class procedure TEventClass.PageControlOnGetImageIndex(Sender: TObject; TabIndex: Integer; var ImageIndex: Integer);
begin
  SendEvent(Sender, gePageControlGetImageIndex, [Sender, TabIndex, @ImageIndex]);
end;

class procedure TEventClass.MenuOnChange(Sender: TObject; Source: TMenuItem;
  Rebuild: Boolean);
begin
  SendEvent(Sender, geMenuChange, [Sender, Source, Rebuild]);
end;

class procedure TEventClass.OnClose(Sender: TObject);
begin
  SendEvent(Sender, geClose, [Sender]);
end;

class procedure TEventClass.FormOnCloseQuery(Sender: TObject;
  var CanClose: Boolean);
begin
  SendEvent(Sender, geFormCloseQuery, [Sender, @CanClose]);
end;

class procedure TEventClass.FormOnDropFiles(Sender: TObject;
  const AFileNames: array of string);
begin
  SendEvent(Sender, geDropFiles, [Sender, @AFileNames[0], Length(AFileNames)]);
end;

class procedure TEventClass.OnDblClick(Sender: TObject);
begin
  SendEvent(Sender, geDblClick, [Sender]);
end;

class procedure TEventClass.OnDestroy(Sender: TObject);
begin
  SendEvent(Sender, geDestroy, [Sender]);
end;

class procedure TEventClass.OnEnter(Sender: TObject);
begin
  SendEvent(Sender, geEnter, [Sender]);
end;

class procedure TEventClass.OnException(Sender: TObject; E: Exception);
begin
  SendEvent(Sender, geException, [Sender, E]);
end;

class procedure TEventClass.OnExecute(Sender: TObject);
begin
  SendEvent(Sender, geExecute, [Sender]);
end;

class procedure TEventClass.OnExit(Sender: TObject);
begin
  SendEvent(Sender, geExit, [Sender]);
end;

class procedure TEventClass.OnFind(Sender: TObject);
begin
  SendEvent(Sender, geFind, [Sender]);
end;


class procedure TEventClass.OnActivate(Sender: TObject);
begin
  SendEvent(Sender, geActivate, [Sender]);
end;

class procedure TEventClass.OnDeactivate(Sender: TObject);
begin
  SendEvent(Sender, geDeactivate, [Sender]);
end;

class procedure TEventClass.OnConstrainedResize(Sender: TObject; var MinWidth, MinHeight, MaxWidth, MaxHeight: Integer);
begin
  SendEvent(Sender, geConstrainedResize, [Sender, @MinWidth, @MinHeight, @MaxWidth, @MaxHeight]);
end;




class function TEventClass.OnHelp(Command: Word; Data: THelpEventData;
  var CallHelp: Boolean): Boolean;
var
  LResult: Boolean;
begin
  SendEvent(Application, geHelp, [Command, Data, Pointer(@CallHelp), Pointer(@LResult)]);
  Result := LResult;
end;

class procedure TEventClass.OnShortCut(var Msg: TWMKey; var Handled: Boolean);
begin
  SendEvent(Application, geShortCut, [Pointer(@Msg), Pointer(@Handled)]);
end;

class procedure TEventClass.OnContextPopup(Sender: TObject; MousePos: TPoint; var Handled: Boolean);
begin
  SendEvent(Sender, geContextPopup, [Sender, Pointer(@MousePos), Pointer(@Handled)]);
end;

class procedure TEventClass.OnDockDrop(Sender: TObject; Source: TDragDockObject; X, Y: Integer);
begin
  SendEvent(Sender, geDockDrop, [Sender, Source, X, Y]);
end;

class procedure TEventClass.OnDragDrop(Sender, Source: TObject; X, Y: Integer);
begin
  SendEvent(Sender, geDragDrop, [Sender, Source, X, Y]);
end;

class procedure TEventClass.OnDragOver(Sender, Source: TObject; X, Y: Integer; State: TDragState; var Accept: Boolean);
begin
  SendEvent(Sender, geDragOver, [Sender, Source, X, Y, Integer(State), Pointer(@Accept)]);
end;

class procedure TEventClass.OnEndDock(Sender, Target: TObject; X, Y: Integer);
begin
  SendEvent(Sender, geEndDock, [Sender, Target, X, Y]);
end;

class procedure TEventClass.OnGetSiteInfo(Sender: TObject; DockClient: TControl;
  var InfluenceRect: TRect; MousePos: TPoint; var CanDock: Boolean);
begin
  SendEvent(Sender, geGetSiteInfo, [Sender, DockClient, Pointer(@InfluenceRect), Pointer(@MousePos), Pointer(@CanDock)]);
end;

class procedure TEventClass.OnMouseWheelDown(Sender: TObject; Shift: TShiftState; MousePos: TPoint; var Handled: Boolean);
begin
  SendEvent(Sender, geMouseWheelDown, [Sender, PWord(@Shift)^, Pointer(@MousePos), Pointer(@Handled)]);
end;

class procedure TEventClass.OnMouseWheelUp(Sender: TObject; Shift: TShiftState; MousePos: TPoint; var Handled: Boolean);
begin
  SendEvent(Sender, geMouseWheelUp, [Sender, PWord(@Shift)^, Pointer(@MousePos), Pointer(@Handled)]);
end;

class procedure TEventClass.OnStartDock(Sender: TObject; var DragObject: TDragDockObject);
begin
  SendEvent(Sender, geStartDock, [Sender, DragObject]);
end;



class procedure TEventClass.OnUnDock(Sender: TObject; Client: TControl; NewTarget: TWinControl; var Allow: Boolean);
begin
  SendEvent(Sender, geUnDock, [Sender, Client, NewTarget, Pointer(@Allow)]);
end;


class procedure TEventClass.OnMessage(var Msg: TMsg; var Handled: Boolean);
begin
  SendEvent(Application, geMessage, [Pointer(@Msg), Pointer(@Handled)]);
end;


class procedure TEventClass.OnEndDrag(Sender, Target: TObject; X, Y: Integer);
begin
  SendEvent(Sender, geEndDrag, [Sender, Target, X, Y]);
end;




class procedure TEventClass.OnGesture(Sender: TObject; const EventInfo: TGestureEventInfo; var Handled: Boolean);
begin
  SendEvent(Sender, geGesture, [Sender, @EventInfo, @Handled]);
end;

class procedure TEventClass.OnMouseActivate(Sender: TObject; Button: TMouseButton; Shift: TShiftState;
  X, Y, HitTest: Integer; var MouseActivate: TMouseActivate);
begin
  SendEvent(Sender, geMouseActivate, [Sender, Ord(Button), PWord(@Shift)^, X, Y, HitTest, @MouseActivate]);
end;


class procedure TEventClass.ListBoxOnData(Control: TWinControl; Index: Integer; var Data: string);
var
  LData: PChar;
begin
  LData := PChar(Data);
  SendEvent(Control, geListBoxData, [Control,Index, @LData]);
  Data := LData;
end;


class function TEventClass.ListBoxOnDataFind(Control: TWinControl; FindString: string): Integer;
begin
  SendEvent(Control, geListBoxDataFind, [Control, PChar(FindString), @Result]);
end;

class procedure TEventClass.ListBoxOnDataObject(Control: TWinControl; Index: Integer; var DataObject: TObject);
begin
  SendEvent(Control, geListBoxDataObject, [Control, Index, @DataObject]);
end;

class procedure TEventClass.ListBoxOnMeasureItem(Control: TWinControl; Index: Integer; var Height: Integer);
begin
  SendEvent(Control, geListBoxMeasureItem, [Control, Index, @Height]);
end;

class procedure TEventClass.OnChanging(Sender: TObject);
begin
  SendEvent(Sender, geChanging, [Sender]);
end;

class procedure TEventClass.UpDownOnChanging(Sender: TObject; var AllowChange: Boolean);
begin
  SendEvent(Sender, geUpDownChanging, [Sender, @AllowChange]);
end;

class procedure TEventClass.ListViewOnChanging(Sender: TObject; Item: TListItem; Change: TItemChange; var AllowChange: Boolean);
begin
  SendEvent(Sender, geListViewChanging, [Sender, Item, Ord(Change), @AllowChange]);
end;

class procedure TEventClass.ListViewOnData(Sender: TObject; Item: TListItem);
begin
  SendEvent(Sender, geListViewData, [Sender, Item]);
end;

class procedure TEventClass.ListViewOnDataFind(Sender: TObject; Find: TItemFind;
  const FindString: string; const FindPosition: TPoint; FindData: Pointer;
  StartIndex: Integer; Direction: TSearchDirection; Wrap: Boolean; var Index: Integer);
begin
  SendEvent(Sender, geListViewDataFind, [Sender, Ord(Find), PChar(FindString), @FindPosition, FindData, StartIndex,
    Ord(Direction), Integer(Wrap), @Index]);
end;

class procedure TEventClass.ListViewOnEdited(Sender: TObject; Item: TListItem; var S: string);
var
  LS: PChar;
begin
  LS := PChar(S);
  SendEvent(Sender, geListViewEdited, [Sender, Item, @LS]);
  S := LS;
end;

class procedure TEventClass.ListViewOnEditing(Sender: TObject; Item: TListItem; var AllowEdit: Boolean);
begin
  SendEvent(Sender, geListViewEditing, [Sender, Item, @AllowEdit]);
end;

class procedure TEventClass.ListViewOnInsert(Sender: TObject; Item: TListItem);
begin
  SendEvent(Sender, geListViewInsert, [Sender, Item]);
end;

class procedure TEventClass.ListViewOnDeletion(Sender: TObject; Item: TListItem);
begin
  SendEvent(Sender, geListViewDeletion, [Sender, Item]);
end;


class procedure TEventClass.ListViewOnCustomDraw(Sender: TCustomListView; const ARect: TRect; var DefaultDraw: Boolean);
begin
  SendEvent(Sender, geListViewCustomDraw, [Sender, @ARect, @DefaultDraw]);
end;

class procedure TEventClass.ListViewOnCustomDrawItem(Sender: TCustomListView;
  Item: TListItem; State: TCustomDrawState; var DefaultDraw: Boolean);
begin
  SendEvent(Sender, geListViewCustomDrawItem, [Sender, Item, PWord(@State)^, @DefaultDraw]);
end;

class procedure TEventClass.ListViewOnCustomDrawSubItem(Sender: TCustomListView;
  Item: TListItem; SubItem: Integer; State: TCustomDrawState; var DefaultDraw: Boolean);
begin
  SendEvent(Sender, geListViewCustomDrawSubItem, [Sender, Item, SubItem, PWord(@State)^, @DefaultDraw]);
end;

class procedure TEventClass.ListViewOnDrawItem(Sender: TCustomListView; Item: TListItem; ARect: TRect; State: TOwnerDrawState);
begin
  SendEvent(Sender, geListViewDrawItem, [Sender, Item, @ARect, PWord(@State)^]);
end;




class procedure TEventClass.TreeViewOnChanging(Sender: TObject; Node: TTreeNode; var AllowChange: Boolean);
begin
  SendEvent(Sender, geTreeViewChanging, [Sender, Node, @AllowChange]);
end;

class procedure TEventClass.TreeViewOnCancelEdit(Sender: TObject; Node: TTreeNode);
begin
  SendEvent(Sender, geTreeViewCancelEdit, [Sender, Node]);
end;

class procedure TEventClass.TreeViewOnAddition(Sender: TObject; Node: TTreeNode);
begin
  SendEvent(Sender, geTreeViewAddition, [Sender, Node]);
end;

class procedure TEventClass.TreeViewOnCollapsed(Sender: TObject; Node: TTreeNode);
begin
  SendEvent(Sender, geTreeViewCollapsed, [Sender, Node]);
end;

class procedure TEventClass.TreeViewOnCollapsing(Sender: TObject; Node: TTreeNode; var AllowCollapse: Boolean);
begin
  SendEvent(Sender, geTreeViewCollapsing, [Sender, Node, @AllowCollapse]);
end;

class procedure TEventClass.TreeViewOnDeletion(Sender: TObject; Node: TTreeNode);
begin
  SendEvent(Sender, geTreeViewDeletion, [Sender, Node]);
end;

class procedure TEventClass.TreeViewOnEdited(Sender: TObject; Node: TTreeNode; var S: string);
var
  LS: PChar;
begin
  LS := PChar(S);
  SendEvent(Sender, geTreeViewEdited, [Sender, Node, @LS]);
  S := LS;
end;

class procedure TEventClass.TreeViewOnEditing(Sender: TObject; Node: TTreeNode; var AllowEdit: Boolean);
begin
  SendEvent(Sender, geTreeViewEditing, [Sender, Node, @AllowEdit]);
end;

class procedure TEventClass.TreeViewOnExpanded(Sender: TObject; Node: TTreeNode);
begin
  SendEvent(Sender, geTreeViewExpanded, [Sender, Node]);
end;

class procedure TEventClass.TreeViewOnExpanding(Sender: TObject; Node: TTreeNode; var AllowExpansion: Boolean);
begin
  SendEvent(Sender, geTreeViewExpanding, [Sender, Node, @AllowExpansion]);
end;

class procedure TEventClass.TreeViewOnHint(Sender: TObject; const Node: TTreeNode; var Hint: string);
var
  LHint: PChar;
begin
  LHint := PChar(Hint);
  SendEvent(Sender, geTreeViewHint, [Sender, Node, @LHint]);
  Hint := LHint;
end;

class procedure TEventClass.TreeViewOnCustomDraw(Sender: TCustomTreeView; const ARect: TRect; var DefaultDraw: Boolean);
begin
  SendEvent(Sender, geTreeViewCustomDraw, [Sender, @ARect, @DefaultDraw]);
end;

class procedure TEventClass.TreeViewOnCustomDrawItem(Sender: TCustomTreeView; Node: TTreeNode; State: TCustomDrawState; var DefaultDraw: Boolean);
begin
  SendEvent(Sender, geTreeViewCustomDrawItem, [Sender, Node, PWord(@State)^, @DefaultDraw]);
end;



class procedure TEventClass.MenuItemOnMeasureItem(Sender: TObject; ACanvas: TCanvas; var Width, Height: Integer);
begin
  SendEvent(Sender, geMenuItemMeasureItem, [Sender, ACanvas, @Width, @Height]);
end;

class procedure TEventClass.PageControlOnChanging(Sender: TObject; var AllowChange: Boolean);
begin
  SendEvent(Sender, gePageControlChanging, [Sender, @AllowChange]);
end;






class procedure TEventClass.JumpListOnItemDeleted(Sender: TObject;
  const Item: TJumpListItem; const CategoryName: string; FromTasks: Boolean);
begin
  SendEvent(Sender, geJumpListItemDeleted, [Sender, Item, PChar(CategoryName), Integer(FromTasks)]);
end;

class procedure TEventClass.JumpListOnListUpdateError(Sender: TObject;
  WinErrorCode: Cardinal; const ErrorDescription: string; var Handled: Boolean);
begin
  SendEvent(Sender, geJumpListListUpdateError, [Sender, WinErrorCode, PChar(ErrorDescription), @Handled]);
end;

class procedure TEventClass.JumpListOnItemsLoaded(Sender: TObject);
begin
  SendEvent(Sender, geJumpListItemsLoaded, [Sender]);
end;


class procedure TEventClass.TaskbarOnThumbPreviewRequest(Sender: TObject; APreviewHeight, APreviewWidth: Integer; PreviewBitmap: TBitmap);
begin
  SendEvent(Sender, geTaskbarThumbPreviewRequest, [Sender, APreviewHeight, APreviewWidth, PreviewBitmap]);
end;

class procedure TEventClass.TaskbarOnWindowPreviewItemRequest(Sender: TObject; var Position: TPoint; PreviewBitmap: TBitmap);
begin
  SendEvent(Sender, geTaskbarWindowPreviewItemRequest, [Sender, @Position, PreviewBitmap]);
end;

class procedure TEventClass.TaskbarOnThumbButtonClick(Sender: TObject; AButtonID: Integer);
begin
  // 这里要处理下
  if Sender = nil then
    Exit;
  SendEvent(TThumbBarButton(Sender).Collection.Owner, geTaskbarThumbButtonClick, [Sender, AButtonID]);
end;

// grid
class procedure TEventClass.OnColumnMoved(Sender: TObject; FromIndex, ToIndex: Longint);
begin
  SendEvent(Sender, geColumnMoved, [Sender, FromIndex, ToIndex]);
end;

class procedure TEventClass.OnDrawCell(Sender: TObject; ACol, ARow: Longint; ARect: TRect; State: TGridDrawState);
begin
  SendEvent(Sender, geDrawCell, [Sender, ACol, ARow, Pointer(@ARect), PWord(@State)^]);
end;

class procedure TEventClass.OnFixedCellClick(Sender: TObject; ACol, ARow: Integer);
begin
  SendEvent(Sender, geFixedCellClick, [Sender, ACol, ARow]);
end;

class procedure TEventClass.OnGetEditMask(Sender: TObject; ACol, ARow: Integer; var Value: string);
var
  LS: PChar;
begin
  LS := PChar(Value);
  SendEvent(Sender, geGetEditMask, [Sender, ACol, ARow, Pointer(@LS)]);
  Value := LS;
end;

class procedure TEventClass.OnGetEditText(Sender: TObject; ACol, ARow: Integer; var Value: string);
var
  LS: PChar;
begin
  LS := PChar(Value);
  SendEvent(Sender, geGetEditText, [Sender, ACol, ARow, Pointer(@LS)]);
  Value := LS;
end;

class procedure TEventClass.OnRowMoved(Sender: TObject; FromIndex, ToIndex: Integer);
begin
  SendEvent(Sender, geRowMoved, [Sender, FromIndex, ToIndex]);
end;

class procedure TEventClass.OnSelectCell(Sender: TObject; ACol, ARow: Integer; var CanSelect: Boolean);
begin
  SendEvent(Sender, geSelectCell, [Sender, ACol, ARow, Pointer(@CanSelect)]);
end;

class procedure TEventClass.OnSetEditText(Sender: TObject; ACol, ARow: Integer; const Value: string);
begin
  SendEvent(Sender, geSetEditText, [Sender, ACol, ARow, PChar(Value)]);
end;

class procedure TEventClass.OnTopLeftChanged(Sender: TObject);
begin
  SendEvent(Sender, geTopLeftChanged, [Sender]);
end;


// headercontrol
class procedure TEventClass.OnDrawSection(HeaderControl: THeaderControl; Section: THeaderSection; const Rect: TRect; Pressed: Boolean);
begin
  SendEvent(HeaderControl, geDrawSection, [HeaderControl, Section, Pointer(@Rect), Pressed]);
end;

class procedure TEventClass.OnSectionCheck(HeaderControl: TCustomHeaderControl; Section: THeaderSection);
begin
  SendEvent(HeaderControl, geSectionCheck, [HeaderControl, Section]);
end;

class procedure TEventClass.OnSectionClick(HeaderControl: THeaderControl; Section: THeaderSection);
begin
  SendEvent(HeaderControl, geSectionClick, [HeaderControl, Section]);
end;

class procedure TEventClass.OnSectionDrag(Sender: TObject; FromSection, ToSection: THeaderSection; var AllowDrag: Boolean);
begin
  SendEvent(Sender, geSectionClick, [Sender, FromSection, ToSection, Pointer(@AllowDrag)]);
end;

class procedure TEventClass.OnSectionEndDrag(Sender: TObject);
begin
  SendEvent(Sender, geSectionEndDrag, [Sender]);
end;

class procedure TEventClass.OnSectionResize(HeaderControl: THeaderControl; Section: THeaderSection);
begin
  SendEvent(HeaderControl, geSectionResize, [HeaderControl, Section]);
end;

class procedure TEventClass.OnSectionTrack(HeaderControl: THeaderControl; Section: THeaderSection; Width: Integer; State: TSectionTrackState);
begin
  SendEvent(HeaderControl, geSectionTrack, [HeaderControl, Section, Width, Integer(State)]);
end;





class procedure TEventClass.OnHide(Sender: TObject);
begin
  SendEvent(Sender, geHide, [Sender]);
end;

class procedure TEventClass.OnHint(Sender: TObject);
begin
   SendEvent(Sender, geHint, [Sender]);
end;

class procedure TEventClass.OnKeyDown(Sender: TObject; var Key: Word;
  Shift: TShiftState);
begin
  SendEvent(Sender, geKeyDown, [Sender, @Key, PWord(@Shift)^]);
end;

class procedure TEventClass.OnKeyPress(Sender: TObject; var Key: Char);
begin
  SendEvent(Sender, geKeyPress, [Sender, @Key]);
end;

class procedure TEventClass.OnKeyUp(Sender: TObject; var Key: Word;
  Shift: TShiftState);
begin
  SendEvent(Sender, geKeyUp, [Sender, @Key, PWord(@Shift)^]);
end;

class procedure TEventClass.OnMinimize(Sender: TObject);
begin
  SendEvent(Sender, geMinimize, [Sender]);
end;

class procedure TEventClass.OnMouseDown(Sender: TObject; Button: TMouseButton;
  Shift: TShiftState; X, Y: Integer);
begin
  SendEvent(Sender, geMouseDown, [Sender, Ord(Button), PWord(@Shift)^, X, Y]);
end;

class procedure TEventClass.OnMouseEnter(Sender: TObject);
begin
  SendEvent(Sender, geMouseEnter, [Sender]);
end;

class procedure TEventClass.OnMouseLeave(Sender: TObject);
begin
  SendEvent(Sender, geMouseLeave, [Sender]);
end;

class procedure TEventClass.OnMouseMove(Sender: TObject; Shift: TShiftState; X,
  Y: Integer);
begin
  SendEvent(Sender, geMouseMove, [Sender, PWord(@Shift)^, X, Y]);
end;

class procedure TEventClass.OnMouseUp(Sender: TObject; Button: TMouseButton;
  Shift: TShiftState; X, Y: Integer);
begin
  SendEvent(Sender, geMouseUp, [Sender, Ord(Button), PWord(@Shift)^, X, Y]);
end;

class procedure TEventClass.OnMouseWheel(Sender: TObject; Shift: TShiftState;
  WheelDelta: Integer; MousePos: TPoint; var Handled: Boolean);
begin
  SendEvent(Sender, geMouseWheel, [Sender, PWord(@Shift)^, WheelDelta, MousePos.X, MousePos.Y, @Handled]);
end;

class procedure TEventClass.OnLinkClick(Sender: TObject; const Link: string;
  LinkType: TSysLinkType);
begin
  SendEvent(Sender, geLinkClick, [Sender, Link, Ord(LinkType)]);
end;

class procedure TEventClass.OnStyleChanged(Sender: TObject);
begin
  SendEvent(Sender, geStyleChanged, [Sender]);
end;

class procedure TEventClass.OnPaint(Sender: TObject);
begin
  SendEvent(Sender, gePaint, [Sender]);
end;

class procedure TEventClass.OnPopup(Sender: TObject);
begin
  SendEvent(Sender, gePopup, [Sender]);
end;

class procedure TEventClass.OnReplace(Sender: TObject);
begin
  SendEvent(Sender, geReplace, [Sender]);
end;

class procedure TEventClass.OnResize(Sender: TObject);
begin
  SendEvent(Sender, geResize, [Sender]);
end;

class procedure TEventClass.OnRestore(Sender: TObject);
begin
  SendEvent(Sender, geRestore, [Sender]);
end;

class procedure TEventClass.OnShow(Sender: TObject);
begin
  SendEvent(Sender, geShow, [Sender]);
end;

class procedure TEventClass.OnTimer(Sender: TObject);
begin
  SendEvent(Sender, geTimer, [Sender]);
end;

class procedure TEventClass.OnUpdate(Sender: TObject);
begin
  SendEvent(Sender, geUpdate, [Sender]);
end;

class procedure TEventClass.Remove(AObj: TObject; AEvent: TGoEvent);
begin
  FEvents.Remove(TEventKey.Create(AObj, AEvent));
end;

class procedure TEventClass.ListBoxOnDrawItem(Control: TWinControl; Index: Integer;
   ARect: TRect; State: TOwnerDrawState);
begin
  SendEvent(Control, geListBoxDrawItem, [Control, Index, @ARect, PWord(@State)^]);
end;

class procedure TEventClass.MenuItemOnDrawItem(Sender: TObject; ACanvas: TCanvas;
  ARect: TRect; Selected: Boolean);
begin
  SendEvent(Sender, geListBoxDrawItem, [Sender, ACanvas, @ARect, Selected]);
end;


class procedure TEventClass.SendEvent(Sender: TObject; AEvent: TGoEvent; AArgs: array of const);


  procedure SendEventSrc(EventId: NativeUInt; AArgs: array of const);
  var
    LParams: array[0..29] of TGoParam;
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
          LParams[I].&Type := LV.VType;
          case LV.VType of
            vtInteger       : LParams[I].Value := Pointer(LV.VInteger);
            vtBoolean       : LParams[I].Value := Pointer(Byte(LV.VBoolean));
  //          vtChar          = 2;
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
        GEventCallbackPtr(EventId, @LParams[0], LArgLen);
      end;
    end;
  end;

var
  LEventId: NativeUInt;
begin
  if FEvents <> nil then
    if FEvents.TryGetValue(TEventKey.Create(Sender, AEvent), LEventId) then
      SendEventSrc(LEventId, AArgs);
end;


// ---------------------- TTreeView

class procedure TEventClass.TreeViewOnAdvancedCustomDraw(
  Sender: TCustomTreeView; const ARect: TRect; Stage: TCustomDrawStage;
  var DefaultDraw: Boolean);
begin
  SendEvent(Sender, geTreeViewAdvancedCustomDraw,
    [Sender, @ARect, Ord(Stage), DefaultDraw]);
end;

class procedure TEventClass.TreeViewOnAdvancedCustomDrawItem(
  Sender: TCustomTreeView; Node: TTreeNode; State: TCustomDrawState;
  Stage: TCustomDrawStage; var PaintImages, DefaultDraw: Boolean);
begin
  SendEvent(Sender, geTreeViewAdvancedCustomDrawItem,
    [Sender, Node, PWord(@State)^, Ord(Stage), @PaintImages, DefaultDraw]);
end;

class procedure TEventClass.TreeViewOnChange(Sender: TObject; ANode: TTreeNode);
begin
  SendEvent(Sender, geTreeViewChange, [Sender, ANode]);
end;

class procedure TEventClass.TreeViewOnGetImageIndex(Sender: TObject; Node: TTreeNode);
begin
  SendEvent(Sender, geTreeViewGetImageIndex, [Sender, Node]);
end;

class procedure TEventClass.TreeViewOnGetSelectedIndex(Sender: TObject; Node: TTreeNode);
begin
  SendEvent(Sender, geTreeViewGetSelectedIndex, [Sender, Node]);
end;

class procedure TEventClass.TreeViewOnCompare(Sender: TObject; Node1, Node2: TTreeNode; Data: Integer; var Compare: Integer);
begin
  SendEvent(Sender, geTreeViewCompare, [Sender, Node1, Node2, Data, @Compare]);
end;


//----------- TToolBar
class procedure TEventClass.ToolBarOnAdvancedCustomDraw(Sender: TToolBar;
  const ARect: TRect; Stage: TCustomDrawStage; var DefaultDraw: Boolean);
begin
  SendEvent(Sender, geToolBarAdvancedCustomDraw,
    [Sender, @ARect, Ord(Stage), @DefaultDraw]);
end;

class procedure TEventClass.ToolBarOnAdvancedCustomDrawButton(Sender: TToolBar;
  Button: TToolButton; State: TCustomDrawState; Stage: TCustomDrawStage;
  var Flags: TTBCustomDrawFlags; var DefaultDraw: Boolean);
begin
  SendEvent(Sender, geToolBarAdvancedCustomDrawButton,
    [Sender, Button, PWord(@State)^, Ord(Stage), PWord(@Flags), @DefaultDraw]);
end;

class procedure TEventClass.UpDownOnClick(Sender: TObject; Button: TUDBtnType);
begin
  SendEvent(Sender, geUpDownClick, [Sender, Ord(Button)]);
end;

{ TEventKey }

constructor TEventKey.Create(ASender: TObject; AEvent: TGoEvent);
begin
  Sender := ASender;
  Event := AEvent;
end;


{ TMessageEventClass }

class constructor TMessageEventClass.Create;
begin
  FMsgEvents := TMessageEventList.Create;
end;

class destructor TMessageEventClass.Destroy;
begin
  FreeAndNil(FMsgEvents);
end;

class procedure TMessageEventClass.OnWndProc(Sender: TObject; var AMsg: TMessage;
  var AHandled: Boolean);
var
  LId: NativeUInt;
begin
  if Assigned(GMessageCallbackPtr) then
  begin
    if FMsgEvents.TryGetValue(Sender, LId) then
      GMessageCallbackPtr(LId, @AMsg, @AHandled);
  end;
end;

class procedure TMessageEventClass.Remove(AObj: TObject);
begin
  FMsgEvents.Remove(AObj);
end;

class procedure TMessageEventClass.Add(AObj: TObject; AId: NativeUInt);
begin
  FMsgEvents.AddOrSetValue(AObj, AId);
end;

end.
