//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
// 
// Licensed under Apache License 2.0
//
//----------------------------------------

unit uComponents;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, fgl;

type
  TClassLists = specialize  TFPGMap<string, TClass>;

  var uClassLists: TClassLists;

implementation

uses
  Interfaces, // this includes the LCL widgetset,
{$IFDEF WINDOWS}
  Windows,
  MultiMon,
  ShellAPI,
  CommCtrl,
  ActiveX,
{$ELSE}
  LCLType,
  Types,
{$ENDIF}
  typinfo,
  LCLProc,
  LCLIntf,
  DateUtils,
  IniFiles,
  Registry,
  Forms,
  StdCtrls,
  Dialogs,
  ExtCtrls,
  Graphics,
  Controls,
  Buttons,
  ComCtrls,
  ToolWin,
  ImgList,
  ExtDlgs,
  ActnList,
  ColorBox,
  PrintersDlgs,
  DateTimePicker,
  Calendar,
  Menus,
  Clipbrd,
  CheckLst,
  MaskEdit,
  uLinkLabel,
  ImageButton,
  Grids,
  ValEdit,
  Gauges,
  Spin,
  ComboEx,
  {$I UserDefineComponentUses.inc}
  uMiniWebview,
  XButton,
  uControlPatchs;

{$I LazarusExtDef.inc}


procedure AddComponentClass(AClass: TClass);
begin
  uClassLists.AddOrSetData(AClass.ClassName, AClass);
end;

procedure MyRegisterComponents(AClass: array of TClass);
var
  LB: TClass;
begin
  for LB in AClass do
    AddComponentClass(LB);
end;

procedure InitClassLists;
begin
  MyRegisterComponents([
      //TApplication,
      //TForm,
      TButton, TBitBtn, TMaskEdit, TEdit, TMainMenu, TPopupMenu, TMemo, TCheckBox,
      TRadioButton,TGroupBox,TLabel,TListBox,TComboBox,TPanel,TImage,TLinkLabel,
      TSpeedButton,TSplitter,TRadioGroup,TStaticText,TColorBox,TColorListBox,
      TTrayIcon,{TBalloonHint,TCategoryPanelGroup,TCategoryPanel,}TOpenDialog,
      TSaveDialog,TColorDialog,TFontDialog,TPrintDialog,TOpenPictureDialog,
      TSavePictureDialog{,TSaveTextFileDialog,TOpenTextFileDialog},
      {TRichEdit,}TTrackBar,TImageList,TUpDown,TProgressBar,
      {THotKey,}TDateTimePicker,TMonthCalendar,TListView,TTreeView,TStatusBar,
      TToolBar{,TIcon,TBitmap,TMemoryStream,TFont,TStrings,
      TStringList,TBrush,TPen},TMenuItem{,TListGroups,TPicture,
      TListColumns,TListItems,TTreeNodes,TListItem,TTreeNode},
      TPageControl,TTabSheet,TControl{,TScreen,TMouse,TListGroup,
      TListColumn,TCollectionItem},TStatusPanels,{TStatusPanel,
      TCanvas,TObject,TPngImage,TJPEGImage,TGIFImage,TGIFFrame,}
      TActionList,TAction,TToolButton,{TIniFile,TRegistry,TClipboard,Forms.TMonitor,}
      {TMargins,TPadding,}TPaintBox,TTimer,TComponent,{TList,TGraphic,TMonthCalColors,}
      {TParaAttributes,TTextAttributes,TIconOptions,}TScrollBar,TShape,TBevel,TScrollBox,
      TCheckListBox,TGauge{,TCustomHint},TImageButton,TFontDialog,TFindDialog,TReplaceDialog,TPageSetupDialog,
      TPrinterSetupDialog,
      TStringGrid, TDrawGrid, TValueListEditor, THeaderControl,
      {THeaderSection,THeaderSections,}TLabeledEdit,TBoundLabel,
      TFlowPanel,TCoolBar,TCoolBands,TCoolBand, TSpinEdit,TMiniWebview,
      TTaskDialog, {TTaskDialogButtons, TTaskDialogProgressBar, TTaskDialogButtonItem, TTaskDialogRadioButtonItem, TTaskDialogBaseButtonItem,}TCalendar,
      {TControlBorderSpacing,}TComboBoxEx,{TFrame,TControlScrollBar,TSizeConstraints,}TXButton
  ]);
  {$I UserDefineComponentsClass.inc}
end;



initialization
  uClassLists := TClassLists.Create;
  InitClassLists;

finalization
  uClassLists.Free;

end.

