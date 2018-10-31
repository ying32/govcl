unit uComponents;

interface

uses
  Winapi.GDIPAPI,
  Winapi.GDIPOBJ,
  System.Classes,
  System.SysUtils,
  System.Generics.Collections;

var
  uClassLists: TDictionary<string, TClass>;

  procedure SetGlobalFormScaled(AValue: LongBool); stdcall;
  function GetGlobalFormScaled: LongBool; stdcall;

  procedure ScaledForm(AForm: TComponent);

implementation


uses
  Winapi.Windows,
  Winapi.MultiMon,
  Winapi.ShellAPI,
  Winapi.CommCtrl,
  System.UITypes,
  System.DateUtils,
  System.Types,
  System.IniFiles,
  System.Win.Registry,
  System.Actions,

  Vcl.Forms,
  Vcl.StdCtrls,
  Vcl.Dialogs,
  Vcl.ExtCtrls,
  Vcl.Graphics,
  Vcl.Controls,
  Vcl.Buttons,
  Vcl.ComCtrls,
  Vcl.ToolWin,
  Vcl.ImgList,
  Vcl.ExtDlgs,
  Vcl.ActnList,
  Vcl.Menus,
  Vcl.Imaging.jpeg,
  Vcl.Imaging.GIFImg,
  Vcl.GraphUtil,
  Vcl.Imaging.pngimage,
  Vcl.Clipbrd,
  Vcl.Themes,
  Vcl.Styles,
  Vcl.Mask,
  Vcl.CheckLst,
  Vcl.Grids,
  Vcl.ValEdit,
  Vcl.Samples.Gauges,
  Vcl.FileCtrl,
  Vcl.JumpList,
  Vcl.Taskbar,
  Vcl.Samples.Spin,
  ImageButton;

var
  uGlobalFormScaled: Boolean = False;


//--------------------- Ô­DelphiResForm.incÄÚÈÝ

const
  ClassRefArrs: array[0..126] of TClass = (
    TApplication,TForm,TButton,TBitBtn,TMaskEdit,TEdit,TMainMenu,TPopupMenu,TMemo,TCheckBox,
    TRadioButton,TGroupBox,TLabel,TListBox,TComboBox,TPanel,TImage,TLinkLabel,
    TSpeedButton,TSplitter,TRadioGroup,TStaticText,TColorBox,TColorListBox,
    TTrayIcon,TBalloonHint,TCategoryPanelGroup,TCategoryPanel,TOpenDialog,
    TSaveDialog,TColorDialog,TFontDialog,TPrintDialog,TOpenPictureDialog,
    TSavePictureDialog,TSaveTextFileDialog,TOpenTextFileDialog,
    TRichEdit,TTrackBar,TImageList,TUpDown,TProgressBar,
    THotKey,TDateTimePicker,TMonthCalendar,TListView,TTreeView,TStatusBar,
    TToolBar,TIcon,TBitmap,TMemoryStream,TFont,TStrings,
    TStringList,TBrush,TPen,TMenuItem,TListGroups,TPicture,
    TListColumns,TListItems,TTreeNodes,TListItem,TTreeNode,
    TPageControl,TTabSheet,TControl,TScreen,TMouse,TListGroup,
    TListColumn,TCollectionItem,TStatusPanels,TStatusPanel,
    TCanvas,TObject,TPngImage,TJPEGImage,TGIFImage,TGIFFrame,
    TActionList,TAction,TToolButton,TIniFile,TRegistry,TClipboard,Vcl.Forms.TMonitor,
    TMargins,TPadding,TPaintBox,TTimer,TList,TGraphic,TComponent,TMonthCalColors,
    TParaAttributes,TTextAttributes,TIconOptions,TScrollBar,TShape,TBevel,TScrollBox,
    TCheckListBox,TGauge,TCustomHint,TImageButton,TFontDialog,TFindDialog, TReplaceDialog,TPageSetupDialog,
    TPrinterSetupDialog,
    TStringGrid, TDrawGrid, TValueListEditor, THeaderControl,
    THeaderSection,THeaderSections,TLabeledEdit,TBoundLabel,
    TFlowPanel,TCoolBar,TCoolBands,TCoolBand,TTaskbar,TJumpList,TSpinEdit
  );


procedure SetGlobalFormScaled(AValue: LongBool);
begin
  uGlobalFormScaled := AValue;
end;

function GetGlobalFormScaled: LongBool;
begin
  Result := uGlobalFormScaled;
end;

procedure ScaledForm(AForm: TComponent);
begin
  if uGlobalFormScaled then
  begin
    if AForm is TForm then
      (AForm as TForm).ScaleForPPI(Screen.PixelsPerInch)
  end;
end;

procedure InitClassLists;
var
  LB: TClass;
begin
  for LB in ClassRefArrs do
    uClassLists.AddOrSetValue(LB.ClassName, LB);
end;

exports
  SetGlobalFormScaled,
  GetGlobalFormScaled;


procedure GidpInit;
begin
  // Initialize StartupInput structure
  StartupInput.DebugEventCallback := nil;
  StartupInput.SuppressBackgroundThread := False;
  StartupInput.SuppressExternalCodecs   := False;
  StartupInput.GdiplusVersion := 1;

  GdiplusStartup(gdiplusToken, @StartupInput, nil);
end;

procedure GdipUnInit;
begin
  if Assigned(GenericSansSerifFontFamily) then
    GenericSansSerifFontFamily.Free;
  if Assigned(GenericSerifFontFamily) then
    GenericSerifFontFamily.Free;
  if Assigned(GenericMonospaceFontFamily) then
    GenericMonospaceFontFamily.Free;
  if Assigned(GenericTypographicStringFormatBuffer) then
    GenericTypographicStringFormatBuffer.free;
  if Assigned(GenericDefaultStringFormatBuffer) then
    GenericDefaultStringFormatBuffer.Free;

  GdiplusShutdown(gdiplusToken);
end;


initialization
  GidpInit;
  uClassLists := TDictionary<string, TClass>.Create;
  InitClassLists;

finalization
  uClassLists.Free;
  GdipUnInit;


end.
