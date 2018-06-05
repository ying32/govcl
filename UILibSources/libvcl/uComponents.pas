unit uComponents;

interface

uses
  System.Classes,
  System.SysUtils,
  System.Generics.Collections;

var
  uClassLists: TDictionary<string, TClass>;

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
  ImageButton;

//--------------------- Ô­DelphiResForm.incÄÚÈÝ

const
  ClassRefArrs: array[0..123] of TClass = (
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
    TFlowPanel,TCoolBar,TCoolBands,TCoolBand
  );


procedure InitClassLists;
var
  LB: TClass;
begin
  for LB in ClassRefArrs do
    uClassLists.AddOrSetValue(LB.ClassName, LB);
end;

initialization
  uClassLists := TDictionary<string, TClass>.Create;
  InitClassLists;

finalization
  uClassLists.Free;


end.
