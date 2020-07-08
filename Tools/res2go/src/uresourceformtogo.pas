unit uresourceformtogo;

{$IFDEF FPC}
  {$mode objfpc}{$H+}
//{$IFDEF WINDOWS}
//  {$codepage UTF8}
//{$ENDIF}
{$ENDIF}

interface

procedure ConvertAll;

implementation

uses
{$IFDEF MSWINDOWS}
  Windows,
{$ENDIF}
  Classes, SysUtils, Math, StrUtils, uFormDesignerFile
{$IFNDEF FPC}
  , System.Generics.Collections
{$ELSE}
  , fgl
{$ENDIF};

const
  APPVERSION = '1.0.21';

type
  TComponentItem = record
    Name: string;
    ClassName: string;
  end;
  PComponentItem = ^TComponentItem;

  TSupportComponentItem = record
    ClassName: string;
    PkgName: string;  // 需要生成的包名，为空则为vcl
    UnitName: string  // 单元名，用于将gfm转lfm+pas所有
  end;

  TEventItem = record
    Name: string;
    RealName: string;
    ComponentName: string;
    ComponentClass: string;
  end;

  TEventType = record
    Name: string;
    // 指定此事件只有此控件应用（暂时未启用，待以后再弄）
    ControlClassName: string;
    
    ImportTypePkg: boolean;
    Params: string;
  end;
  PEventType = ^TEventType;

  { TStringStreamHelper }
{$IFDEF FPC}
  TStringStreamHelper = class Helper for TStringStream
  public
    procedure LoadFromFile(const AFileName: string);
    procedure SaveToFile(const AFileName: string);
    procedure Clear;
  end;

{$ENDIF}


const
  // 事件
  {
     还要确定独有事件的处理
     先不管了，
  }
  // 特殊
  CommonEventType: array[0..116] of TEventType = (

  (Name: 'Exception'; ControlClassName: 'TApplication'; ImportTypePkg: False; Params: 'sender vcl.IObject, e vcl.Exception'),
  (Name: 'Destroy'; ControlClassName: 'TForm'; ImportTypePkg: False; Params: 'sender vcl.IObject'),
  (Name: 'Create'; ControlClassName: 'TForm'; ImportTypePkg: False; Params: 'sender vcl.IObject'),
  (Name: 'Help'; ControlClassName: 'TForm'; ImportTypePkg: True; Params: 'command uint16, data types.THelpEventData, callhelp, result *bool'),
  (Name: 'ShortCut'; ControlClassName: 'TForm'; ImportTypePkg: True; Params: 'msg *types.TWMKey, handled *bool'),
  (Name: 'Close'; ControlClassName: 'TForm'; ImportTypePkg: True; Params: 'sender vcl.IObject, action *types.TCloseAction'),
  (Name: 'CloseQuery'; ControlClassName: 'TForm'; ImportTypePkg: False; Params: 'sender vcl.IObject, canClose *bool'),
  (Name: 'DropFiles'; ControlClassName: 'TForm'; ImportTypePkg: False; Params: 'sender vcl.IObject, aFileNames []string'),
  (Name: 'LinkClick'; ControlClassName: 'TLinkLabel'; ImportTypePkg: True; Params: 'sender vcl.IObject, link string, linkType types.TSysLinkType'),
  (Name: 'Changing'; ControlClassName: 'TUpDown'; ImportTypePkg: False; Params: 'sender vcl.IObject, allowChange *bool'),
  (Name: 'Click'; ControlClassName: 'TUpDown'; ImportTypePkg: True; Params: 'sender vcl.IObject, button types.TUDBtnType'),
  (Name: 'AdvancedCustomDraw'; ControlClassName: 'TListView'; ImportTypePkg: True; Params: 'sender *vcl.TListView, aRect types.TRect, stage types.TCustomDrawStage, defaultDraw *bool'),
  (Name: 'AdvancedCustomDrawItem'; ControlClassName: 'TListView'; ImportTypePkg: True; Params: 'sender *vcl.TListView, item *vcl.TListItem, state types.TCustomDrawState, Stage types.TCustomDrawStage, defaultDraw *bool'),
  (Name: 'AdvancedCustomDrawSubItem'; ControlClassName: 'TListView'; ImportTypePkg: True; Params: 'sender *vcl.TListView, item *vcl.TListItem, subItem int32, state types.TCustomDrawState, stage types.TCustomDrawStage, defaultDraw *bool'),
  (Name: 'Change'; ControlClassName: 'TListView'; ImportTypePkg: True; Params: 'sender vcl.IObject, item *vcl.TListItem, change types.TItemChange'),
  (Name: 'ColumnClick'; ControlClassName: 'TListView'; ImportTypePkg: False; Params: 'sender vcl.IObject, column *vcl.TListColumn'),
  (Name: 'Compare'; ControlClassName: 'TListView'; ImportTypePkg: False; Params: 'sender vcl.IObject, item1, item2 *vcl.TListItem, data int32, compare *int32'),
  (Name: 'CustomDraw'; ControlClassName: 'TListView'; ImportTypePkg: True; Params: 'sender *vcl.TListView, aRect types.TRect, defaultDraw *bool'),
  (Name: 'CustomDrawItem'; ControlClassName: 'TListView'; ImportTypePkg: True; Params: 'sender *vcl.TListView, item *vcl.TListItem, state types.TCustomDrawStage, defaultDraw *bool'),
  (Name: 'CustomDrawSubItem'; ControlClassName: 'TListView'; ImportTypePkg: True; Params: 'sender *vcl.TListView, item *vcl.TListItem, subItem int32, state types.TCustomDrawStage, defaultDraw *bool'),
  (Name: 'Data'; ControlClassName: 'TListView'; ImportTypePkg: False; Params: 'sender vcl.IObject, item *vcl.TListItem'),
  (Name: 'DataFind'; ControlClassName: 'TListView'; ImportTypePkg: True; Params: 'sender vcl.IObject, find types.TItemFind, findString string, findPosition types.TPoint, findData types.TCustomData, startIndex int32, direction types.TSearchDirection, warp bool, index *int32'),
  (Name: 'DataHint'; ControlClassName: 'TListView'; ImportTypePkg: False; Params: 'sender vcl.IObject, startIndex, endIndex int32'),
  (Name: 'Deletion'; ControlClassName: 'TListView'; ImportTypePkg: False; Params: 'sender vcl.IObject, item *vcl.TListItem'),
  (Name: 'Edited'; ControlClassName: 'TListView'; ImportTypePkg: False; Params: 'sender vcl.IObject, item *vcl.TListItem, s *string'),
  (Name: 'Editing'; ControlClassName: 'TListView'; ImportTypePkg: False; Params: 'sender vcl.IObject, item *vcl.TListItem, allowEdit *bool'),
  (Name: 'Insert'; ControlClassName: 'TListView'; ImportTypePkg: False; Params: 'sender vcl.IObject, item *vcl.TListItem'),
  (Name: 'SelectItem'; ControlClassName: 'TListView'; ImportTypePkg: False; Params: 'sender vcl.IObject, item *vcl.TListItem, selected bool'),
  (Name: 'ItemChecked'; ControlClassName: 'TListView'; ImportTypePkg: False; Params: 'sender vcl.IObject, item *vcl.TListItem'),
  (Name: 'Addition'; ControlClassName: 'TTreeView'; ImportTypePkg: False; Params: 'sender vcl.IObject, node *vcl.TTreeNode'),
  (Name: 'Collapsed'; ControlClassName: 'TTreeView'; ImportTypePkg: False; Params: 'sender vcl.IObject, node *vcl.TTreeNode'),
  (Name: 'Deletion'; ControlClassName: 'TTreeView'; ImportTypePkg: False; Params: 'sender vcl.IObject, node *vcl.TTreeNode'),
  (Name: 'Expanded'; ControlClassName: 'TTreeView'; ImportTypePkg: False; Params: 'sender vcl.IObject, node *vcl.TTreeNode'),
  (Name: 'GetSelectedIndex'; ControlClassName: 'TTreeView'; ImportTypePkg: False; Params: 'sender vcl.IObject, node *vcl.TTreeNode'),
  (Name: 'AdvancedCustomDraw'; ControlClassName: 'TTreeView'; ImportTypePkg: True; Params: 'sender *vcl.TTreeView, aRect types.TRect, stage types.TCustomDrawStage, defaultDraw *bool'),
  (Name: 'AdvancedCustomDrawItem'; ControlClassName: 'TTreeView'; ImportTypePkg: True; Params: 'sender *vcl.TTreeView, node *vcl.TTreeNode, state types.TCustomDrawState, stage types.TCustomDrawStage, paintImages, defaultDraw *bool'),
  (Name: 'Change'; ControlClassName: 'TTreeView'; ImportTypePkg: False; Params: 'sender vcl.IObject, node *vcl.TTreeNode'),
  (Name: 'Changing'; ControlClassName: 'TTreeView'; ImportTypePkg: False; Params: 'sender vcl.IObject, node *vcl.TTreeNode, allowChange *bool'),
  (Name: 'Collapsing'; ControlClassName: 'TTreeView'; ImportTypePkg: False; Params: 'sender vcl.IObject, node *vcl.TTreeNode, allowCollapse *bool'),
  (Name: 'Compare'; ControlClassName: 'TTreeView'; ImportTypePkg: False; Params: 'sender vcl.IObject, node1, node2 *vcl.TTreeNode, data int32, compare *int32'),
  (Name: 'CustomDraw'; ControlClassName: 'TTreeView'; ImportTypePkg: True; Params: 'sender *vcl.TTreeView, aRect types.TRect, defaultDraw *bool'),
  (Name: 'CustomDrawItem'; ControlClassName: 'TTreeView'; ImportTypePkg: True; Params: 'sender *vcl.TTreeView, node *vcl.TTreeNode, state types.TCustomDrawStage, defaultDraw *bool'),
  (Name: 'Edited'; ControlClassName: 'TTreeView'; ImportTypePkg: False; Params: 'sender vcl.IObject, node *vcl.TTreeNode, s *string'),
  (Name: 'Editing'; ControlClassName: 'TTreeView'; ImportTypePkg: False; Params: 'sender vcl.IObject, node *vcl.TTreeNode, allowEdit *bool'),
  (Name: 'Expanding'; ControlClassName: 'TTreeView'; ImportTypePkg: False; Params: 'sender vcl.IObject, node *vcl.TTreeNode, allowExpansion *bool'),
  (Name: 'DrawItem'; ControlClassName: 'TListBox'; ImportTypePkg: True; Params: 'control vcl.IWinControl, index int32, aRect types.TRect, state types.TOwnerDrawState'),
  (Name: 'MeasureItem'; ControlClassName: 'TListBox'; ImportTypePkg: False; Params: 'control *vcl.TWinControl, index int32, height *int32'),
  (Name: 'DrawItem'; ControlClassName: 'TComboBox'; ImportTypePkg: True; Params: 'control vcl.IWinControl, index int32, aRect types.TRect, state types.TOwnerDrawState'),
  (Name: 'MeasureItem'; ControlClassName: 'TComboBox'; ImportTypePkg: False; Params: 'control *vcl.TWinControl, index int32, height *int32'),
  (Name: 'Change'; ControlClassName: 'TMainMenu'; ImportTypePkg: False; Params: 'sender vcl.IObject, source *vcl.TMenuItem, rebuild bool'),
  (Name: 'MeasureItem'; ControlClassName: 'TMainMenu'; ImportTypePkg: False; Params: 'sender vcl.IObject, aCanvas *vcl.TCanvas, width, height *int32'),
  (Name: 'DrawItem'; ControlClassName: 'TMainMenu'; ImportTypePkg: True; Params: 'sender vcl.IObject, aCanvas *vcl.TCanvas, aRect types.TRect, selected bool'),
  (Name: 'MeasureItem'; ControlClassName: 'TPopupMenu'; ImportTypePkg: False; Params: 'sender vcl.IObject, aCanvas *vcl.TCanvas, width, height *int32'),
  (Name: 'DrawItem'; ControlClassName: 'TPopupMenu'; ImportTypePkg: True; Params: 'sender vcl.IObject, aCanvas *vcl.TCanvas, aRect types.TRect, selected bool'),
  (Name: 'Changing'; ControlClassName: 'TPageControl'; ImportTypePkg: False; Params: 'sender vcl.IObject, allowChange *bool'),
  (Name: 'TitleChange'; ControlClassName: 'TMiniWebview'; ImportTypePkg: False; Params: 'sender vcl.IObject, text string'),
  (Name: 'JSExternal'; ControlClassName: 'TMiniWebview'; ImportTypePkg: False; Params: 'sender vcl.IObject, funcName, args string, retVal *string'),
  (Name: 'SectionClick'; ControlClassName: 'THeaderControl'; ImportTypePkg: False; Params: 'headerControl *vcl.THeaderControl, section *vcl.THeaderSection'),
  (Name: 'SectionResize'; ControlClassName: 'THeaderControl'; ImportTypePkg: False; Params: 'headerControl *vcl.THeaderControl, section *vcl.THeaderSection'),
  (Name: 'SectionTrack'; ControlClassName: 'THeaderControl'; ImportTypePkg: True; Params: 'headerControl *vcl.THeaderControl, section *vcl.THeaderSection, width int32, state types.TSectionTrackState'),
  (Name: 'SectionDrag'; ControlClassName: 'THeaderControl'; ImportTypePkg: False; Params: 'sender vcl.IObject, fromSection, toSection *vcl.THeaderSection, allowDrag *bool'),
  (Name: 'ButtonClicked'; ControlClassName: 'TTaskDialog'; ImportTypePkg: True; Params: 'sender vcl.IObject, modalResult types.TModalResult, canClose *bool'),
  (Name: 'ItemClick'; ControlClassName: 'TCheckGroup'; ImportTypePkg: False; Params: 'sender vcl.IObject, index int32'),
  (Name: 'Activate'; ControlClassName: ''; ImportTypePkg: False; Params: 'sender vcl.IObject'),
  (Name: 'Deactivate'; ControlClassName: ''; ImportTypePkg: False; Params: 'sender vcl.IObject'),
  (Name: 'Hint'; ControlClassName: ''; ImportTypePkg: False; Params: 'sender vcl.IObject'),
  (Name: 'Minimize'; ControlClassName: ''; ImportTypePkg: False; Params: 'sender vcl.IObject'),
  (Name: 'Restore'; ControlClassName: ''; ImportTypePkg: False; Params: 'sender vcl.IObject'),
  (Name: 'Click'; ControlClassName: ''; ImportTypePkg: False; Params: 'sender vcl.IObject'),
  (Name: 'DblClick'; ControlClassName: ''; ImportTypePkg: False; Params: 'sender vcl.IObject'),
  (Name: 'Hide'; ControlClassName: ''; ImportTypePkg: False; Params: 'sender vcl.IObject'),
  (Name: 'MouseEnter'; ControlClassName: ''; ImportTypePkg: False; Params: 'sender vcl.IObject'),
  (Name: 'MouseLeave'; ControlClassName: ''; ImportTypePkg: False; Params: 'sender vcl.IObject'),
  (Name: 'Paint'; ControlClassName: ''; ImportTypePkg: False; Params: 'sender vcl.IObject'),
  (Name: 'Resize'; ControlClassName: ''; ImportTypePkg: False; Params: 'sender vcl.IObject'),
  (Name: 'Show'; ControlClassName: ''; ImportTypePkg: False; Params: 'sender vcl.IObject'),
  (Name: 'Enter'; ControlClassName: ''; ImportTypePkg: False; Params: 'sender vcl.IObject'),
  (Name: 'Exit'; ControlClassName: ''; ImportTypePkg: False; Params: 'sender vcl.IObject'),
  (Name: 'Change'; ControlClassName: ''; ImportTypePkg: False; Params: 'sender vcl.IObject'),
  (Name: 'Popup'; ControlClassName: ''; ImportTypePkg: False; Params: 'sender vcl.IObject'),
  (Name: 'Select'; ControlClassName: ''; ImportTypePkg: False; Params: 'sender vcl.IObject'),
  (Name: 'Close'; ControlClassName: ''; ImportTypePkg: False; Params: 'sender vcl.IObject'),
  (Name: 'Changing'; ControlClassName: ''; ImportTypePkg: False; Params: 'sender vcl.IObject'),
  (Name: 'Execute'; ControlClassName: ''; ImportTypePkg: False; Params: 'sender vcl.IObject'),
  (Name: 'Update'; ControlClassName: ''; ImportTypePkg: False; Params: 'sender vcl.IObject'),
  (Name: 'Timer'; ControlClassName: ''; ImportTypePkg: False; Params: 'sender vcl.IObject'),
  (Name: 'ClickCheck'; ControlClassName: ''; ImportTypePkg: False; Params: 'sender vcl.IObject'),
  (Name: 'Find'; ControlClassName: ''; ImportTypePkg: False; Params: 'sender vcl.IObject'),
  (Name: 'Replace'; ControlClassName: ''; ImportTypePkg: False; Params: 'sender vcl.IObject'),
  (Name: 'TopLeftChanged'; ControlClassName: ''; ImportTypePkg: False; Params: 'sender vcl.IObject'),
  (Name: 'SectionEndDrag'; ControlClassName: ''; ImportTypePkg: False; Params: 'sender vcl.IObject'),
  (Name: 'ConstrainedResize'; ControlClassName: ''; ImportTypePkg: False; Params: 'sender vcl.IObject, minWidth, minHeight, maxWidth, maxHeight *int32'),
  (Name: 'AlignPosition'; ControlClassName: ''; ImportTypePkg: True; Params: 'sender *vcl.TWinControl, control *vcl.TControl, newLeft, newTop, newWidth, newHeight *int32, alignRect *types.TRect, alignInfo types.TAlignInfo'),
  (Name: 'ContextPopup'; ControlClassName: ''; ImportTypePkg: True; Params: 'sender vcl.IObject, mousePos types.TPoint, handled *bool'),
  (Name: 'DockDrop'; ControlClassName: ''; ImportTypePkg: False; Params: 'sender vcl.IObject, source *vcl.TDragDockObject, x, y int32'),
  (Name: 'DragDrop'; ControlClassName: ''; ImportTypePkg: False; Params: 'sender, source vcl.IObject, x, y int32'),
  (Name: 'DragOver'; ControlClassName: ''; ImportTypePkg: True; Params: 'sender, source vcl.IObject, x, y int32, state types.TDragState, accept *bool'),
  (Name: 'EndDrag'; ControlClassName: ''; ImportTypePkg: False; Params: 'sender, target vcl.IObject, x, y int32'),
  (Name: 'EndDock'; ControlClassName: ''; ImportTypePkg: False; Params: 'sender, target vcl.IObject, x, y int32'),
  (Name: 'GetSiteInfo'; ControlClassName: ''; ImportTypePkg: True; Params: 'sender vcl.IObject, dockClient *vcl.TControl, influenceRect *types.TRect, mousePos types.TPoint, canDock *bool'),
  (Name: 'KeyDown'; ControlClassName: ''; ImportTypePkg: True; Params: 'sender vcl.IObject, key *types.Char, shift types.TShiftState'),
  (Name: 'KeyUp'; ControlClassName: ''; ImportTypePkg: True; Params: 'sender vcl.IObject, key *types.Char, shift types.TShiftState'),
  (Name: 'KeyPress'; ControlClassName: ''; ImportTypePkg: True; Params: 'sender vcl.IObject, key *types.Char'),
  (Name: 'MouseDown'; ControlClassName: ''; ImportTypePkg: True; Params: 'sender vcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32'),
  (Name: 'MouseUp'; ControlClassName: ''; ImportTypePkg: True; Params: 'sender vcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32'),
  (Name: 'MouseMove'; ControlClassName: ''; ImportTypePkg: True; Params: 'sender vcl.IObject, shift types.TShiftState, x, y int32'),
  (Name: 'MouseWheel'; ControlClassName: ''; ImportTypePkg: True; Params: 'sender vcl.IObject, shift types.TShiftState, wheelDelta, x, y int32, handled *bool'),
  (Name: 'MouseWheelDown'; ControlClassName: ''; ImportTypePkg: True; Params: 'sender vcl.IObject, shift types.TShiftState, mousePos types.TPoint, handled *bool'),
  (Name: 'MouseWheelUp'; ControlClassName: ''; ImportTypePkg: True; Params: 'sender vcl.IObject, shift types.TShiftState, mousePos types.TPoint, handled *bool'),
  (Name: 'StartDock'; ControlClassName: ''; ImportTypePkg: False; Params: 'sender vcl.IObject, dragObject *vcl.TDragDockObject'),
  (Name: 'UnDock'; ControlClassName: ''; ImportTypePkg: False; Params: 'sender vcl.IObject, client *vcl.TControl, newTarget *vcl.TControl, allow *bool'),
  (Name: 'ColRowMoved'; ControlClassName: ''; ImportTypePkg: False; Params: 'sender vcl.IObject, isColumn bool, sIndex, tIndex int32'),
  (Name: 'DrawCell'; ControlClassName: ''; ImportTypePkg: True; Params: 'sender vcl.IObject, aCol, aRow int32, aRect types.TRect, state types.TGridDrawState'),
  (Name: 'GetEditMask'; ControlClassName: ''; ImportTypePkg: False; Params: 'sender vcl.IObject, aCol, aRow int32, value *string'),
  (Name: 'GetEditText'; ControlClassName: ''; ImportTypePkg: False; Params: 'sender vcl.IObject, aCol, aRow int32, value *string'),
  (Name: 'SetEditText'; ControlClassName: ''; ImportTypePkg: False; Params: 'sender vcl.IObject, aCol, aRow int32, value string'),
  (Name: 'SelectCell'; ControlClassName: ''; ImportTypePkg: False; Params: 'sender vcl.IObject, aCol, aRow int32, canSelect *bool')

 );

{$I supportsComponents.inc}


var
{$IFDEF MSWINDOWS}
  uConsoleHandle: THandle;
{$ENDIF}
  uGoPkgName: string; // 使用全局的，懒得添加额外的改变了


//控制台文本的颜色常量
const
  tfGreen = 2;
  tfRed = 4;
  tfIntensity = 8;
  tfWhite = $f;

  PrivateFiledsFlagStr = '//::private::';
  PrivateFiledsStr = 'T%sFields';

type
   TEventFindList = specialize  TFPGMap<string, string>;

var
  uErrorPause, uWaringPause: boolean;
  uEventFindList: TEventFindList;



procedure WriteHelp; forward;

procedure CrtTextColor(AColor: byte);
begin
{$IFDEF MSWINDOWS}
  SetConsoleTextAttribute(uConsoleHandle, AColor);
{$ELSE}
  //crt.TextColor(AColor);
  case AColor of
    tfRed: Write(#27'[0;31m');  // red
    tfGreen: Write(#27'[0;32m');  // green
    tfWhite: Write(#27'[0;37m');  // white
  end;
{$ENDIF}
end;

procedure CtlWriteln(const AFmt: string; AArgs: array of const); inline; overload;
begin
{$IFDEF FPC}
  {$IFDEF WINDOWS}
    Writeln(UTF8Decode(Format(AFmt, AArgs))); //UTF8ToUTF16
  {$ELSE}
    Writeln(Format(AFmt, AArgs));
  {$ENDIF}
{$ELSE}
  Writeln(Format(AFmt, AArgs));
{$ENDIF}
end;

procedure CtlWriteln(const AFmt: string); inline; overload;
begin
  CtlWriteln('%s', [AFmt]);
end;

procedure TextColorRed;
begin
  CrtTextColor(
{$IFDEF MSWINDOWS}
    tfIntensity or
{$ENDIF}
    tfRed);
end;

procedure TextColorGreen;
begin
  CrtTextColor(tfGreen);
end;

procedure TextColorWhite;
begin
  CrtTextColor(tfWhite);
end;


// 做Lazarus和Delphi两边兼容的，同一种组件但名称不一样的。
function FixClass(AOrgiName: string): string;
begin
  if AOrgiName = 'TCalendar' then
    Result := 'TMonthCalendar'
  else
    Result := AOrgiName;
end;


// 中文环境，不区分简/繁
{
  "zh-CN":  2052, // 简体
  "zh-HK":  3076, // 繁体
  "zh-MO":  5124, // 繁体
  "zh-SG":  4100, // 简体
  "zh-TW":  1028, // 繁体
}
function IsZhLang: Boolean;
{$IFDEF WINDOWS}
const
   Langs:array[0..4] of Integer = (2052, 3076, 5124, 4100, 1028);
{$ELSE}
const
   Langs:array[0..4] of string = ('zh_CN', 'zh_HK', 'zh_MO', 'zh_SG', 'zh_TW');
{$ENDIF}
var
  I: Integer;
begin
  Result := False;
  for I := 0 to High(Langs) do
  begin
  {$IFDEF MSWINDOWS}
    if Langs[I] = SysLocale.DefaultLCID then
  {$ELSE}
    if Pos(Langs[I], GetEnvironmentVariable('LANG')) <> -1 then
  {$ENDIF}
      Exit(True);
  end;
end;

// 获取指命令行的一下个参数
function GetNextParam(ASwitch: string): string;
var
  I, L: integer;
  S, T: string;
begin
  Result := '';
  S := ASwitch;
  I := ParamCount;
  while I > 0 do
  begin
    L := Length(ParamStr(I));
    if (L > 0) and (ParamStr(I)[1] in SwitchChars) then
    begin
      T := Copy(ParamStr(I), 2, L - 1);
      if S = T then
        Exit(ParamStr(I + 1));
    end;
    Dec(I);
  end;
end;

function GetEventParam(AItem: TEventItem): string;
begin
  Result := '';
  if not uEventFindList.TryGetData(AItem.ComponentClass+AItem.Name, Result) then
    if not uEventFindList.TryGetData(AItem.Name, Result) then
       Result := '';
end;

function IsSupportsComponent(AClassName: string): boolean;
var
  LItem: TSupportComponentItem;
begin
  Result := False;
  for LItem in supportsComponents do
  begin
    if LItem.ClassName = AClassName then
      Exit(True);
  end;
end;

function IsComponentPackageName(AClassName: string): string;
var
  LItem: TSupportComponentItem;
begin
  Result := 'vcl';
  for LItem in supportsComponents do
  begin
    if (LItem.ClassName = AClassName) and (LItem.PkgName <> '') then
      Exit(LItem.PkgName);
  end;
end;


function GetNeedTypesPkg(AItem: TEventItem): boolean;
var
  LItem: TEventType;
begin
  Result := False;
  for LItem in CommonEventType do
    if (SameText(LItem.Name, AItem.Name) or SameText(LItem.Name, AItem.RealName)) and
      LItem.ImportTypePkg then
      Exit(True);
end;

procedure CreateImplFile(AFileName: string; AEvents: array of TEventItem;
  AFormName: string);
var
  LImplFileName, LMName, LTemp, LCode, LPrivateName, LFlags: string;
  LItem: TEventItem;
  LStream: TStringStream;
  LExists, LB: boolean;
  LListStr: TStringList;
  I: integer;
begin
  LImplFileName := AFileName;
  Insert('Impl', LImplFileName, Length(LImplFileName) -
    Length(ExtractFileExt(AFileName)) + 1);

  LStream := TStringStream.Create(''
{$IFNDEF FPC}
    , TEncoding.UTF8
{$ENDIF}
    );
  try
    LExists := FileExists(LImplFileName);
    LListStr := TStringList.Create;
    try
      // 不存在，则添加
      if not LExists then
      begin
        if IsZhLang then
        begin
          LListStr.Add('// 由res2go自动生成。');
          LListStr.Add('// 在这里写你的事件。');
        end
        else
        begin
          LListStr.Add('// Automatically generated by the res2go.');
          LListStr.Add('// Write your event here.');
        end;
        LListStr.Add('');
        LListStr.Add('package ' + uGoPkgName);
        LListStr.Add('');
        if Length(AEvents) > 0 then
        begin
          LListStr.Add('import (');
          LListStr.Add('    "github.com/ying32/govcl/vcl"');
          LListStr.Add(')');
        end;
      end
      else
      begin
        // 反之加载
        LStream.LoadFromFile(LImplFileName);
        LTemp := LStream.DataString;
        LListStr.Text := LTemp;

        // 有事件时检查下有没有添加govcl包
        if Length(AEvents) > 0 then
        begin
          if Pos('import', LTemp) = 0 then
          begin
            I := 0;
            while I < LListStr.Count do
            begin
              if Trim(LListStr[I]).StartsWith('package') then
              begin
                Inc(I);
                LListStr.Insert(I, ')');
                LListStr.Insert(I, '    "github.com/ying32/govcl/vcl"');
                LListStr.Insert(I, 'import (');
                LListStr.Insert(I, '');
                Break;
              end;
              Inc(I);
            end;
          end;
        end;
      end;


      // 添加事件
      for LItem in AEvents do
      begin
        LMName := Format('On%s', [LItem.RealName]);
        //CtlWriteln('method name: %s', [LMName]);
        LCode := Format(#13#10'func (f *T%s) %s(%s) {'#13#10#13#10'}'#13#10,
          [AFormName, LMName, GetEventParam(LItem)]);
        // 不存在不查找了
        if not LExists then
        begin
          if Pos(LMName, LListStr.Text) = 0 then
            LListStr.Add(LCode);
        end
        else
        begin
          // 没有找到，则添加
          if Pos(LMName, LListStr.Text) = 0 then
            LListStr.Add(LCode);
        end;
      end;
      // 这里检查下是否需要类型包
      if not LListStr.Text.Contains('"github.com/ying32/govcl/vcl/types"') then
      begin
        for LItem in AEvents do
        begin
          if GetNeedTypesPkg(LItem) then
          begin
            I := 0;
            for LTemp in LListStr do
            begin
              if LTemp.Contains('"github.com/ying32/govcl/vcl"') then
              begin
                LListStr.Insert(I + 1, '    "github.com/ying32/govcl/vcl/types"');
                Break;
              end;
              Inc(I);
            end;
            Break;
          end;
        end;
      end;

      // 检查私有变量结构是否存在
      LPrivateName := Format(PrivateFiledsStr, [AFormName]);
      // 不存在，则添加
      if Pos(PrivateFiledsFlagStr, LListStr.Text) = 0 then
      begin
        I := 0;
        while I < LListStr.Count do
        begin
          // 在首个func前几行插入
          LFlags := 'import';
          LB := (not LExists) and (Length(AEvents) = 0);
          if LB then
            LFlags := 'package';
          if Trim(LListStr[I]).StartsWith(LFlags) then
          begin
            if not LB then
            begin
              repeat
                Inc(I);
              until Trim(LListStr[I]).StartsWith(')');
            end;
            Inc(I);
            LListStr.Insert(I, '');
            LListStr.Insert(I, '}');
            LListStr.Insert(I, 'type ' + LPrivateName + ' struct {');
            LListStr.Insert(I, PrivateFiledsFlagStr);
            LListStr.Insert(I, '');
            Break;
          end;
          Inc(I);
        end;
      end
      else
      begin
        // 如果存在，则更新，因为防止把窗口名称改了，这里同步更新
        for I := 0 to LListStr.Count - 1 do
        begin
          // 在首个func前几行插入
          if Trim(LListStr[I]).CompareTo(PrivateFiledsFlagStr) = 0 then
          begin
            LListStr[I + 1] := 'type ' + LPrivateName + ' struct {';
            Break;
          end;
        end;
      end;

      // 这里是不是还得处理下，将窗口名称做一次替换
      //f *TFrmMain


      LStream.Clear;
      LStream.WriteString(LListStr.Text);
    finally
      LListStr.Free;
    end;
    LStream.SaveToFile(LImplFileName);
  finally
    LStream.Free;
  end;
end;


procedure SaveToGoFile(AComponents: TList; AEvents: array of TEventItem;
  const AResFileName, AOutPath, AOrigFileName: string; AMem: TMemoryStream);
var
  LStrStream, LBuffer: TStringStream;
{$IFDEF FPC}
  LLines: TStringList;
{$ENDIF}

  procedure WLine(s: string = '');
  begin
  {$IFDEF FPC}
    LLines.Add(S);
  {$ELSE}
    LStrStream.WriteString(s + sLineBreak);
  {$ENDIF}
  end;

  function GetMaxLength: integer;
  var
    I: integer;
    C: PComponentItem;
  begin
    Result := 0;
    for I := 0 to AComponents.Count - 1 do
    begin
      C := AComponents[I];
      Result := Max(Result, Length(C^.Name));
    end;
  end;

  function GetIsFrame(AClassName: string): boolean;
  var
    LStrs: TStringList;
    LFileName: string;
  begin
    Result := False;
    LFileName := AResFileName.Remove(AResFileName.LastIndexOf('.'), 4) + '.pas';
    if FileExists(LFileName) then
    begin
      LStrs := TStringList.Create;
      try
        LStrs.LoadFromFile(LFileName);
        Result := LStrs.Text.Contains(Format('%s = class(TFrame)', [AClassName]));
      finally
        LStrs.Free;
      end;
    end;
  end;

var
  I, LMaxLen: integer;
  C: PComponentItem;
  LVarName, LFormName, LFileName, LTempName: string;
  LUseStr: boolean;
  LItem: TEventItem;
  LFindEvent: boolean;
  LReadEventName: string;
  LSCPkgName: string;
  LIsFrame: boolean;
begin
  LStrStream := TStringStream.Create(''
{$IFNDEF FPC}
    , TEncoding.UTF8
{$ENDIF}
    );
  LBuffer := TStringStream.Create(''
{$IFNDEF FPC}
    , TEncoding.UTF8
{$ENDIF}
    );
{$IFDEF FPC}
  LLines := TStringList.Create;
{$ENDIF}
  try
    //-usestr
    LUseStr := True;
    if FindCmdLineSwitch('usestr') then
      LUseStr := SameText(GetNextParam('usestr'), 'True');

    if IsZhLang then
      WLine('// 由res2go自动生成，不要编辑。')
    else
      WLine('// Automatically generated by the res2go, do not edit.');
    WLine('package ' + uGoPkgName);
    WLine;
    WLine('import (');
    WLine('    "github.com/ying32/govcl/vcl"');
    WLine(')');
    WLine;
    LFormName := PComponentItem(AComponents[0])^.Name;

    LIsFrame := False;
    //if PComponentItem(AComponents[0])^.ClassName.StartsWith('TFrame') then
    if GetIsFrame(PComponentItem(AComponents[0])^.ClassName) then
      LIsFrame := True;

    WLine(Format('type T%s struct {', [LFormName]));
    if LIsFrame then
      WLine('    *vcl.TFrame')
    else
      WLine('    *vcl.TForm');
    LMaxLen := GetMaxLength;
    for I := 1 to AComponents.Count - 1 do
    begin
      C := PComponentItem(AComponents[I]);

      if not IsSupportsComponent(C^.ClassName) then
      begin
        uErrorPause := True;
        TextColorRed;
        if IsZhLang then
          CtlWriteln('错误：“%s:%s”不被支持。', [C^.Name, C^.ClassName])
        else
          CtlWriteln('Error: "%s:%s" is not supported.', [C^.Name, C^.ClassName]);
        TextColorWhite;
        // error: exit;
        Exit;
      end;

      if C^.Name = '' then
        Continue;
      if CharInSet(C^.Name[1], ['a'..'z', '_']) then
      begin
        uWaringPause := True;
        TextColorGreen;
        if IsZhLang then
          CtlWriteln('提示：“%s:%s”必须首字母大写才能被导出。',
            [C^.Name, C^.ClassName])
        else
          CtlWriteln('Hint: "%s:%s" must be capitalized first to be exported.',
            [C^.Name, C^.ClassName]);
        TextColorWhite;
        Continue;
      end;
      //CtlWriteln('%s: %s', [C^.Name, C^.ClassName]);
      // 这里查找下，当前组件有事件，但是这个事件是共享的。
      LReadEventName := '';
      LFindEvent := False;
      for LItem in AEvents do
      begin
        if LItem.ComponentName = C^.Name then
        begin
          // 当前实际关联的事件不是自己的，比如  Button2Click != Button1Click
          if C^.Name + LItem.Name <> LItem.RealName then
          begin
            LFindEvent := True;
            if LReadEventName <> '' then
              LReadEventName := LReadEventName + ',';
            LReadEventName := LReadEventName + 'On' + LItem.RealName;
          end;
        end;
      end;
      // CtlWriteln('LReadEventName: %s', [LReadEventName]);
      LSCPkgName := IsComponentPackageName(C^.ClassName);
      LTempName := Copy(C^.Name + DupeString(#32, LMaxLen), 1, LMaxLen);
      if LFindEvent and (LReadEventName <> '') then
        WLine(Format('    %s *%s.%s `events:"%s"`',
          [LTempName, LSCPkgName, C^.ClassName, LReadEventName]))
      else
        WLine(Format('    %s *%s.%s', [LTempName, LSCPkgName, C^.ClassName]));
    end;
    WLine;
    // 添加一个隐式字段，用于私有，方便写一些结构定自定义的变量什么的
    WLine('    ' + PrivateFiledsFlagStr); // 这是一个查找标识
    WLine(Format('    ' + PrivateFiledsStr, [LFormName]));
    WLine('}');
    WLine;
    if not LIsFrame then
    begin
      WLine(Format('var %s *T%s', [LFormName, LFormName]));
      WLine;
      WLine;
      WLine;
    end;
    WLine;
    // AMem = nil表示不以字节输出到go文件
    if AMem = nil then
    begin
      // 不支持这种方式了输出了
      //if not LIsFrame then
      //begin
      //  if SysIsZhCN then
      //    Wline('// 以文件形式加载')
      //  else
      //    WLine('// Loaded as a file.');
      //  WLine(Format('// vcl.Application.CreateForm("%s.gfm", &%s)', [LFormName, LFormName]));
      //end;

      //// 添加一个默认构建的，不使用Application.CreateForm
      //WLine(Format('func New%s(owner vcl.IComponent) (root *T%s)  {', [LFormName, LFormName]));
      //if not LIsFrame then
      //  WLine(Format('    vcl.CreateResForm(owner, "%s.gfm", &root)', [LFormName]))
      //else
      //  WLine(Format('    vcl.CreateResFrame(owner, "%s.gfm", &root)', [LFormName]));
      //WLine('    return');
      //WLine('}');
      //WLine('');
    end
    else
    begin
      LVarName := LFormName + 'Bytes';

      // 包名不为main时，起始不变为小写。
      if uGoPkgName = 'main' then
      begin
    {$IFDEF FPC}
        LVarName[1] := LowerCase(LVarName[1]);
    {$ELSE}
        LVarName[1] := LowerCase(LVarName[1])[1];
    {$ENDIF}
      end;

      if not LIsFrame then
      begin
        if IsZhLang then
          WLine('// 以字节形式加载')
        else
          WLine('// Loaded in bytes.');
        WLine(Format('// vcl.Application.CreateForm(&%s)', [LFormName]));
        WLine;
      end;
      // 添加一个默认构建的，不使用Application.CreateForm
      WLine(Format('func New%s(owner vcl.IComponent) (root *T%s)  {',
        [LFormName, LFormName]));
      if not LIsFrame then
        WLine(Format('    vcl.CreateResForm(owner, &root)', []))
      else
        WLine(Format('    vcl.CreateResFrame(owner, &root)', []));
      WLine('    return');
      WLine('}');
      WLine('');


      if not LUseStr then
      begin
        WLine('var (');
        LBuffer.WriteString(Format('    %s = []byte ' + '{' + sLineBreak, [LVarName]));
      end
      else
      begin
        LBuffer.WriteString(Format('var %s = []byte("', [LVarName]));
      end;
      for I := 0 to AMem.Size - 1 do
      begin
        if (I > 0) and (not LUseStr) then
          LBuffer.WriteString(', ');
        if (I mod 12 = 0) and (not LUseStr) then
          LBuffer.WriteString(IfThen(I > 0, sLineBreak, '') + '        ');
        if not LUseStr then
          LBuffer.WriteString('0x')
        else
          LBuffer.WriteString('\x');
        LBuffer.WriteString(PByte(PByte(AMem.Memory) + I)^.ToHexString(2));
      end;
      LBuffer.WriteString(IfThen(LUseStr, '")', '}'));
      WLine(LBuffer.DataString);
      if not LUseStr then
        WLine(')');

      WLine('');
      if IsZhLang then
        WLine('// 注册Form资源')
      else
        WLine('// Register Form Resources');
      if LIsFrame then
        WLine(Format('var _ = vcl.RegisterFormResource(T%s{}, &%s)',
          [LFormName, LVarName]))
      else
        WLine(Format('var _ = vcl.RegisterFormResource(%s, &%s)',
          [LFormName, LVarName]));
    end;
    if AOrigFileName <> '' then
      LFileName := AOutPath + AOrigFileName + '.go'
    else
      LFileName := AOutPath + LFormName + '.go';
  {$IFDEF FPC}
    LStrStream.WriteString(LLines.Text);
  {$ENDIF}
    LStrStream.SaveToFile(LFileName);
  finally
  {$IFDEF FPC}
    LLines.Free;
  {$ENDIF}
    LBuffer.Free;
    LStrStream.Free;
  end;
  // 一定创建，因为多加了个
  CreateImplFile(LFileName, AEvents, LFormName);
end;


procedure ResouceFormToGo(ASrcFileName, AOutPath: string);
var
  LParser: TParser;
  LComponents: TList;
  LEventList: array of TEventItem;
  LCurObjectName, LCurObjectClass: string;

  // 摘自lazarus classes.pas单元中的代码
  procedure ProcessProperty;
  var
    LPropName, LEventName: string;
    stream: TMemoryStream;
  begin
    try
      LParser.CheckToken(toSymbol);
    except
      LParser.NextToken;
      Exit;
    end;
    LPropName := LParser.TokenString;
    while True do
    begin
      LParser.NextToken;
      if LParser.Token <> '.' then
        break;
      LParser.NextToken;
      LParser.CheckToken(toSymbol);
      LPropName := LPropName + '.' + LParser.TokenString;
    end;
    LParser.CheckToken('=');
    LParser.NextToken;
    case LParser.Token of
      toSymbol:
      begin
        if (CompareText(LParser.TokenString, 'True') <> 0) and
          (CompareText(LParser.TokenString, 'False') <> 0) and
          (CompareText(LParser.TokenString, 'nil') <> 0) and
          (CompareText(Copy(LPropName, 1, 2), 'On') = 0) then
        begin
          LEventName := LParser.TokenComponentIdent;
          if LEventName[1] in ['A'..'Z'] then
          begin
            SetLength(LEventList, Length(LEventList) + 1);
            LEventList[High(LEventList)].Name := LPropName.Substring(2);
            LEventList[High(LEventList)].RealName := LParser.TokenComponentIdent;
            LEventList[High(LEventList)].ComponentName := LCurObjectName;
            if 'Form'+LEventList[High(LEventList)].Name = LEventList[High(LEventList)].RealName then
              LEventList[High(LEventList)].ComponentClass := 'TForm'
            else
              LEventList[High(LEventList)].ComponentClass := LCurObjectClass;
              //writeln(Length(LEventList), '   ',  LEventList[High(LEventList)].RealName, '   ', LEventList[High(LEventList)].Name);
          end;
        end;
      end;
      '[':
      begin
        LParser.NextToken;
        while LParser.Token <> ']' do
          LParser.NextToken;
      end;
      '(':
      begin
        LParser.NextToken;
        while LParser.Token <> ')' do
          LParser.NextToken;
      end;
      '<':
      begin
        LParser.NextToken;
        while LParser.Token <> '>' do
          LParser.NextToken;
      end;
      '{':
      begin
        stream := TMemoryStream.Create;
        try
          LParser.HexToBinary(stream);
        finally
          stream.Free;
        end;
      end;
    end;
    LParser.NextToken;
  end;

  procedure ProcessObject;
  var
    ObjectName, ObjectType: string;
    LItem: PComponentItem;
  begin
    LParser.NextToken;
    LParser.CheckToken(toSymbol);
    ObjectName := '';
    ObjectType := LParser.TokenString;
    LParser.NextToken;
    if LParser.Token = ':' then
    begin
      LParser.NextToken;
      LParser.CheckToken(toSymbol);
      ObjectName := ObjectType;
      ObjectType := LParser.TokenString;
      LParser.NextToken;
      if LParser.Token = '[' then
      begin
        LParser.NextToken;
        LParser.NextToken;
        LParser.CheckToken(']');
        LParser.NextToken;
      end;
    end;
    // 保存对象名称
    LCurObjectName := ObjectName;
    LCurObjectClass := ObjectType;
    //CtlWriteln('%s: %s', [ObjectName, ObjectType]);
    New(LItem);
    LItem^.Name := ObjectName;
    // 修复类
    LItem^.ClassName := FixClass(ObjectType);
    LComponents.Add(LItem);
    while not (LParser.TokenSymbolIs('END') or LParser.TokenSymbolIs('OBJECT') or
        LParser.TokenSymbolIs('INHERITED') or LParser.TokenSymbolIs('INLINE')) do
    begin
      ProcessProperty;
    end;
    while not LParser.TokenSymbolIs('END') do
      ProcessObject;
    LParser.NextToken;
  end;


  procedure ClearList;
  var
    I: integer;
  begin
    for I := 0 to LComponents.Count - 1 do
      Dispose(PComponentItem(LComponents[I]));
  end;

var
  LOutput, LEnStream: TMemoryStream;
  LInput: TFileStream;
  LUseEncrypt, LOutbytes, LOrigfn: boolean;
  LGfmFileName, LTempFileName: string;
begin
  LInput := TFileStream.Create(ASrcFileName, fmOpenRead or fmShareDenyNone);
  try
    //LInput.LoadFromFile(ASrcFileName);
    LOutput := TMemoryStream.Create;
    try
      try
        ObjectTextToBinary(LInput, LOutput);
        LInput.Position := 0;
        LParser := TParser.Create(LInput);
        try
          LComponents := TList.Create;
          try
            ProcessObject;
            LEnStream := TMemoryStream.Create;
            try
              LOutput.Position := 0;

              LUseEncrypt := False;
              if FindCmdLineSwitch('encrypt') then
                LUseEncrypt := SameText(GetNextParam('encrypt'), 'True');

              // 总是为True
              LOutbytes := True;
              //if FindCmdLineSwitch('outbytes') then
              //  LOutbytes := SameText(GetNextParam('outbytes'), 'True');

              // 是否使用原始的dfm/lfm文件名。
              LOrigfn := FindCmdLineSwitch('origfn');
              // 提取单元名称
              LtempFileName := '';
              if LOrigfn then
              begin
                LTempFileName := ExtractFileName(ASrcFileName);
                LTempFileName :=
                  Copy(LTempFileName, 1, Length(LTempFileName) - Length(ExtractFileExt(LTempFileName)));
              end;
              // 使用加密格式的
              if LUseEncrypt then
              begin
                TFormResFile.Encrypt(LOutput, LEnStream);
                if LOutbytes then
                  SaveToGoFile(LComponents, LEventList, ASrcFileName,
                    AOutPath, LtempFileName, LEnStream)
                else
                  SaveToGoFile(LComponents, LEventList, ASrcFileName,
                    AOutPath, LtempFileName, nil);
              end
              else
              begin
                if LOutbytes then
                  SaveToGoFile(LComponents, LEventList, ASrcFileName,
                    AOutPath, LtempFileName, LOutput)
                else
                  SaveToGoFile(LComponents, LEventList, ASrcFileName,
                    AOutPath, LtempFileName, nil);
              end;
              // 保存gfm文件
              //if not LOutbytes then
              //begin
              if LOrigfn then
                LGfmFileName := AOutPath + LTempFileName + '.gfm'
              else
                LGfmFileName :=
                  AOutPath + PComponentItem(LComponents[0])^.Name + '.gfm';
              if LUseEncrypt then
              begin
                LEnStream.Position := 0;
                LEnStream.SaveToFile(LGfmFileName);
              end
              else
              begin
                LOutput.Position := 0;
                LOutput.SaveToFile(LGfmFileName);
              end;
              //end;
            finally
              LEnStream.Free;
            end;
          finally
            ClearList;
            LComponents.Free;
          end;
        finally
          LParser.Free;
        end;
      except
        on E: Exception do
        begin
          if IsZhLang then
            CtlWriteln('错误: %s', [E.message])
          else
            CtlWriteln('Error: %s', [E.message]);
        end;
      end;
    finally
      LOutput.Free;
    end;
  finally
    LInput.Free;
  end;
end;


procedure ProjectFileToMainDotGo(AFileName, AOutPath: string);
var
  LStrs, LMainDotGo: TStringList;
  S, LVarName, LFormName, LSaveFileName: string;
  LP: integer;
  LFile: TStringStream;
  LMainFileExists, LOutWinRes: boolean;
  LForms: array of string;
  LIndex, I: integer;
  LPkg: string;
  LProjFile: TFileStream;
begin
  LStrs := TStringList.Create;
  LMainDotGo := TStringList.Create;
  try
    LSaveFileName := AOutPath + 'main.go';
    LMainFileExists := FileExists(LSaveFileName);
    LProjFile := TFileStream.Create(AFileName, fmOpenRead or fmShareDenyNone);
    try
      LProjFile.Position := 0;
      LStrs.LoadFromStream(LProjFile);
    finally
      LProjFile.Free;
    end;
    // 如果不存在 main.go文件，则新建一个
    if not LMainFileExists then
    begin
      if IsZhLang then
        LMainDotGo.Add('// 由res2go自动生成。')
      else
        LMainDotGo.Add('// Automatically generated by the res2go.');
      LMainDotGo.Add('package main');
      LMainDotGo.Add('');
      LMainDotGo.Add('import (');
      LMainDotGo.Add('    "github.com/ying32/govcl/vcl"');

      LOutWinRes := True;
      if FindCmdLineSwitch('outres') then
        LOutWinRes := SameText(GetNextParam('outres'), 'True');
      // winappres
      if LOutWinRes then
        LMainDotGo.Add('    _ "github.com/ying32/govcl/pkgs/winappres"');
      LMainDotGo.Add(')');
      LMainDotGo.Add('');
      LMainDotGo.Add('func main() {');
      if FindCmdLineSwitch('scale') then
        LMainDotGo.Add('    vcl.Application.SetScaled(true)');
      LMainDotGo.Add('    vcl.Application.Initialize()');
      if SameText(ExtractFileExt(AFileName), '.dpr') then
        LMainDotGo.Add('    vcl.Application.SetMainFormOnTaskBar(true)');
    end
    else
      // 存在则加载此文件
      LMainDotGo.LoadFromFile(LSaveFileName);

    // 如果不是main包，则输出的需要加上包名
    LPkg := '';
    if uGoPkgName <> 'main' then
      LPkg := uGoPkgName + '.';

    for S in LStrs do
    begin
      // 开始提取 Application.CreateForm的
      if S.Trim.StartsWith('Application.CreateForm(') then
      begin
        LP := S.IndexOf(',');
        LFormName := Trim(S.Substring(LP + 1, S.IndexOf(')') - LP - 1));
        LVarName := LFormName + 'Bytes';
        // 包名不为main时，起始不变为小写。
        if uGoPkgName = 'main' then
        begin
        {$IFDEF FPC}
          LVarName[1] := LowerCase(LVarName[1]);
        {$ELSE}
          LVarName[1] := LowerCase(LVarName[1])[1];
        {$ENDIF}
        end;
        SetLength(LForms, Length(LForms) + 1);
        LForms[High(LForms)] :=
          Format('    vcl.Application.CreateForm(&%s)', [{LPkg+LVarName, }LPkg + LFormName]);
        // main.go文件不存在则直接添加
        if not LMainFileExists then
          LMainDotGo.Add(LForms[High(LForms)]);
      end;
    end;

    // main.go文件存在的处理方式
    if LMainFileExists then
    begin
      LIndex := -1;
      for I := LMainDotGo.Count - 1 downto 0 do
      begin
        // 查找并移除现有的
        if LMainDotGo[I].Trim.StartsWith('vcl.Application.CreateForm(') then
          LMainDotGo.Delete(I);
      end;

      for I := 0 to LMainDotGo.Count - 1 do
      begin
        // 找初始语句
        if LMainDotGo[I].Trim.StartsWith('vcl.Application.Initialize') then
        begin
          // 找到了则 I+1为插入起始行
          LIndex := I + 1;
          // 判断下一行是不是 vcl.Application.SetMainFormOnTaskBar
          if LMainDotGo[I + 1].Trim.StartsWith('vcl.Application.SetMainFormOnTaskBar') then
            Inc(LIndex);
          Break;
        end;
      end;

      // 将前面找到的附加进去
      if LIndex <> -1 then
      begin
        for I := High(LForms) downto 0 do
          LMainDotGo.Insert(LIndex, LForms[I]);
      end;
    end;

    if not LMainFileExists then
    begin
      LMainDotGo.Add('    vcl.Application.Run()');
      LMainDotGo.Add('}');
    end;

    LFile := TStringStream.Create(''
{$IFNDEF FPC}
      , TEncoding.UTF8
{$ENDIF}
      );
    try
      LFile.WriteString(LMainDotGo.Text);
      LFile.SaveToFile(LSaveFileName);
    finally
      LFile.Free;
    end;
  finally
    LMainDotGo.Free;
    LStrs.Free;
  end;
end;

procedure ConvertAll;
type
{$IFDEF FPC}
  TWatchFileList = specialize  TFPGMap<string, longint>;
{$ELSE}
  TWatchFileList = TDictionary<string, longint>;
{$ENDIF}

var
{$IFDEF FPC}
  LRec: TRawbyteSearchRec;
{$ELSE}
  LRec: TSearchRec;
{$ENDIF}
  LPath, LOutPath, LExt, LFileName, LPause: string;
  LConvPro, LWatch: boolean;
  LWatchList: TWatchFileList;

  // 从监视列表中查找
  function WatchFile(const AFileName: string; ACurTime: longint): boolean;
  var
    LTime: longint;
  begin
    Result := False;
    if (not LWatch) or (LWatchList = nil) then
      Exit;
  {$IFDEF FPC}
    if LWatchList.TryGetData(AFileName, LTime) then
  {$ELSE}
      if LWatchList.TryGetValue(AFileName, LTime) then
  {$ENDIF}
      begin
        if LTime = ACurTime then
          Result := True;
        //CtlWriteln('Result: %s, AFileName: %s, Time: %d, CurTime: %d', [Result, AFileName, LTime, ACurTime]);
      end;
    // 当前列表中没有或者有，则更新或添加此文件及时间
  {$IFDEF FPC}
    LWatchList.AddOrSetData(AFileName, ACurTime);
  {$ELSE}
    LWatchList.AddOrSetValue(AFileName, ACurTime);
  {$ENDIF}
  end;

begin
  if FindCmdLineSwitch('help') or FindCmdLineSwitch('h') then
  begin
    WriteHelp;
    Exit;
  end;
  if FindCmdLineSwitch('version') or FindCmdLineSwitch('v') then
  begin
    if IsZhLang then
      CtlWriteln('版本：%s', [APPVERSION])
    else
      CtlWriteln('Version: %s', [APPVERSION]);
    Exit;
  end;
{$IFDEF MSWINDOWS}
  uConsoleHandle := GetStdHandle(STD_OUTPUT_HANDLE);
  try
{$ENDIF}
    // 包名
    uGoPkgName := 'main';
    if FindCmdLineSwitch('pkgname') then
    begin
      uGoPkgName := GetNextParam('pkgname');
      if uGoPkgName = '' then
        uGoPkgName := 'main';
    end;

    LConvPro := True;
    if FindCmdLineSwitch('outmain') then
      LConvPro := SameText(GetNextParam('outmain'), 'True');

    LPath := '.' + PathDelim; //{$IFDEF FPC}DirectorySeparator{$ELSE}PathDelim{$ENDIF};
    if FindCmdLineSwitch('path') then
    begin
      LPath := GetNextParam('path');
      if not DirectoryExists(LPath) then
      begin
        TextColorWhite;
        if IsZhLang then
          CtlWriteln('“-path”目录不存在。')
        else
          CtlWriteln('The "-path" directory does not exist.');
        ExitCode := 1;
        Exit;
      end;
      if not LPath.EndsWith(PathDelim) then
        LPath := LPath + PathDelim;
    end;

    LOutPath := '.' + PathDelim;//{$IFDEF FPC}DirectorySeparator{$ELSE}PathDelim{$ENDIF};
    if FindCmdLineSwitch('outpath') then
    begin
      LOutPath := GetNextParam('outpath');
      if not DirectoryExists(LOutPath) then
        CreateDir(LOutPath);
      if not LOutPath.EndsWith(PathDelim) then
        LOutPath := LOutPath + PathDelim;
    end;

    //if LOutWinRes then
    //  WriteDefaultWindowsRes(LOutPath);
    if IsZhLang then
      CtlWriteln('转换完成。')
    else
      CtlWriteln('Done.');

    // 是否监视指定文件夹，有这个参数则表示监视，没有则不监视
    LWatch := FindCmdLineSwitch('watch');
    if LWatch then
      LWatchList := TWatchFileList.Create;
    try
      repeat
        // 搜索文件
        if FindFirst(LPath + '*.*', faAnyFile, LRec) = 0 then
        begin
          repeat
            LExt := ExtractFileExt(LRec.Name);
            LFileName := LPath + LRec.Name;
            if SameText(LExt, '.lfm') or SameText(LExt, '.dfm') then
            begin
              if WatchFile(LFileName, LRec.Time) then
                Continue;

              TextColorWhite;
              if IsZhLang then
                CtlWriteln('------转换文件：%s', [LFileName])
              else
                CtlWriteln('------Transform file: %s', [LFileName]);
              ResouceFormToGo(LFileName, LOutPath);
            end
            else if LConvPro and (SameText(LExt, '.lpr') or SameText(LExt, '.dpr')) and
              (not SameText(LRec.Name, 'res2go.lpr') and not
              SameText(LRec.Name, 'res2go.dpr')) then
            begin
              if WatchFile(LFileName, LRec.Time) then
                Continue;

              TextColorWhite;
              if IsZhLang then
                CtlWriteln('------转换文件：%s', [LFileName])
              else
                CtlWriteln('------Transform file: %s', [LFileName]);
              ProjectFileToMainDotGo(LFileName, LOutPath);
            end;
          until FindNext(LRec) <> 0;
          FindClose(LRec);
        end;
        // 嗯。。。1000ms频率吧
        if LWatch then
          Sleep(1000);
      until not LWatch;
    finally
      if LWatch then
        LWatchList.Free;
    end;

    if FindCmdLineSwitch('pause') then
    begin
      LPause := GetNextParam('pause');
      if (Pos('a', LPause) <> 0) or (uErrorPause and (Pos('e', LPause) <> 0)) or
        (uWaringPause and (Pos('w', LPause) <> 0)) then
      begin
        if IsZhLang then
          CtlWriteln('请按回车键退出。')
        else
          CtlWriteln('Please press Enter to exit.');
        Readln;
      end;
    end;

{$IFDEF MSWINDOWS}
  finally
    if uConsoleHandle > 0 then
      CloseHandle(uConsoleHandle);
  end;
{$ENDIF}
end;

procedure WriteHelp;
begin
  //CtlWriteln('---------------Chinese------------------');

  if IsZhLang then
  begin
    CtlWriteln('');
    TextColorWhite;
    CtlWriteln(
      'res2go是一个将Lazarus/Delphi资源窗口转go工具，可自动解析lfm、dfm中的组件名、组件类型、事件名称。解析lpr、dpr文件中窗口信息。');
    CtlWriteln('');
    CtlWriteln('用法：res2go [-path "C:\project\"] [-outpath "C:\xxx\"] [-outmain true] [-outres true] [-scale]');
    CtlWriteln('  -path       待转换的工程路径，可为空，默认以当前目录为准。');
    CtlWriteln('  -outpath    输出目录，可为空，默认为当前目录。');
    CtlWriteln('  -outmain    是否输出“main.go”，此为解析lpr或者dpr文件，默认为true。');
    CtlWriteln('  -outres     输出一个Windows默认资源文件，如果存在则不创建，默认为true。');
    CtlWriteln('  -scale      缩放窗口选项，默认为false，即不缩放。');
    CtlWriteln('  -encrypt    使用加密格式的*.gfm文件，默认为false。');
    CtlWriteln(
      '  -usestr     当-outbytes标识为true时，加上此参数会以字符形式输出字节，默认为true。 ');
    CtlWriteln('  -origfn     生成的.go文件使用原始的delphi/lazarus单元名，默认为false。 ');
    CtlWriteln(
      '  -pause      结束后根据选项暂停，比如： -pause "ew"，表示有错或者警告，可选为“e”,“w”,“a” e=错误，w=警告，a=忽略其它选项，总是显示。');
    CtlWriteln('  -pkgname    指定生成的go文件包名，默认为main。');
    CtlWriteln('  -watch      监视“-path”目录的文件，如果有改变则进行转换。');
    CtlWriteln('  -h -help    显示帮助。');
    CtlWriteln('  -v -version 显示版本号');
    CtlWriteln('');
  end
  else
  begin
    //CtlWriteln('---------------English------------------');
    CtlWriteln('');
    TextColorWhite;
    CtlWriteln(
      'res2go is a Lazarus/Delphi resource window to go tool, can automatically resolve the lfm, dfm component name, component type and event name. Parse window information in lpr, dpr file.');
    CtlWriteln('');
    CtlWriteln('usage: res2go [-path "C:\project\"] [-outpath "C:\xxx\"] [-outmain true] [-outres true] [-scale]');
    CtlWriteln('  -path       The project path to be converted can be empty. The default is the current directory.');
    CtlWriteln('  -outpath    Output directory, can be empty, the default is the current directory.');
    CtlWriteln('  -outmain    Whether to output "main.go", this is parsing lpr or dpr file, the default is true.');
    CtlWriteln('  -outres     Outputs a Windows default resource file, if it does not exist, the default is true.');
    CtlWriteln('  -scale      The windoscale option, the default is false.');
    CtlWriteln('  -encrypt    Using the encrypted format of the *.gfm file, the default is false.');
    CtlWriteln(
      '  -usestr     When the -outbytes flag is true, adding this parameter will output the bytes as characters, the default is true.');
    CtlWriteln('  -origfn     The generated .go file uses the original delphi/lazarus unit name, the default is false.');
    CtlWriteln(
      '  -pause      After the end, pause according to the option, for example: -pause "ew", indicating that there is a fault or warning, you can choose "e", "w", "a" e=error, w=warning, a=ignore other options, always display.');
    CtlWriteln('  -pkgname    Specifies the name of the generated go file package. The default is main.');
    CtlWriteln('  -watch      Monitor files in the "-path" directory and convert if there are changes.');
    CtlWriteln('  -h -help    Show help.');
    CtlWriteln('  -v -version Show Version.');
    CtlWriteln('');
  end;

end;

{ TStringStreamHelper }
{$IFDEF FPC}
procedure TStringStreamHelper.LoadFromFile(const AFileName: string);
var
  LFileStream: TFileStream;
begin
  LFileStream := TFileStream.Create(AFileName, fmOpenRead);
  try
    Clear;
    LFileStream.Position := 0;
    Self.CopyFrom(LFileStream, LFileStream.Size);
  finally
    LFileStream.Free;
  end;
end;

procedure TStringStreamHelper.SaveToFile(const AFileName: string);
var
  LFileStream: TFileStream;
  OldPos: int64;
begin
  LFileStream := TFileStream.Create(AFileName, fmCreate);
  try
    OldPos := Self.Position;
    Self.Position := 0;
    LFileStream.CopyFrom(Self, Self.Size);
    Self.Position := OldPos;
  finally
    LFileStream.Free;
  end;
end;

procedure TStringStreamHelper.Clear;
begin
  Self.Size := 0;
end;

{$ENDIF}

procedure InitEventFindList;
var
  LItem: TEventType;
begin
  for LItem in CommonEventType do
    uEventFindList.Add(LItem.ControlClassName+LItem.Name, LItem.Params);
end;

initialization
  uEventFindList := TEventFindList.Create;
  InitEventFindList;

finalization
  uEventFindList.Free;


end.
