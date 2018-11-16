
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

library vcl;


uses
  Winapi.Windows,
  Winapi.MultiMon,
  Winapi.ShellAPI,
  Winapi.CommCtrl,
  Winapi.ActiveX,
  Winapi.ShlObj,
  Winapi.ObjectArray,
  System.SysUtils,
  System.Classes,
  System.UITypes,
  System.DateUtils,
  System.Types,
  //System.Variants,
  System.IniFiles,
  System.Win.Registry,
  System.Win.ComObj,
  System.Win.TaskbarCore,
  System.Actions,
  System.Generics.Collections,
  System.TypInfo,
  Vcl.Consts,
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
  Vcl.Samples.Gauges,
  Vcl.FileCtrl,
  Vcl.Grids,
  Vcl.ValEdit,
  Vcl.JumpList,
  Vcl.Taskbar,
  Vcl.Printers,
  Vcl.Samples.Spin,
  uFormDesignerFile,
  uEventCallback in 'uEventCallback.pas',
  ufrmGo in 'ufrmGo.pas' {GoForm},
  ImageButton in 'ImageButton.pas',
  uImages in 'uImages.pas',
  uComponents in 'uComponents.pas',
  uMiniWebview in 'uMiniWebview.pas';

{$R *.res}



{$IFNDEF DEBUG}
  {$WARNINGS OFF}
{$ENDIF}

{$I DelphiDef.inc}

{$I Application.inc}
{$I Form.inc}
{$I Button.inc}
{$I Edit.inc}
{$I MainMenu.inc}
{$I PopupMenu.inc}
{$I Memo.inc}
{$I CheckBox.inc}
{$I RadioButton.inc}
{$I GroupBox.inc}
{$I Label.inc}
{$I ListBox.inc}
{$I ComboBox.inc}
{$I Panel.inc}
{$I Image.inc}
{$I LinkLabel.inc}
{$I SpeedButton.inc}
{$I Splitter.inc}
{$I RadioGroup.inc}
{$I StaticText.inc}
{$I ColorBox.inc}
{$I ColorListBox.inc}
{$I TrayIcon.inc}
{$I BalloonHint.inc}
{$I CategoryPanelGroup.inc}
{$I CategoryPanel.inc}
{$I OpenDialog.inc}
{$I SaveDialog.inc}
{$I ColorDialog.inc}
{$I FontDialog.inc}
{$I PrintDialog.inc}
{$I OpenPictureDialog.inc}
{$I SavePictureDialog.inc}
{$I SaveTextFileDialog.inc}
{$I OpenTextFileDialog.inc}
{$I RichEdit.inc}
{$I TrackBar.inc}
{$I ImageList.inc}
{$I UpDown.inc}
{$I ProgressBar.inc}
{$I HotKey.inc}
{$I DateTimePicker.inc}
{$I MonthCalendar.inc}
{$I ListView.inc}
{$I TreeView.inc}
{$I StatusBar.inc}
{$I ToolBar.inc}
{$I MaskEdit.inc}
{$I BitBtn.inc}
{$I Icon.inc}
{$I Bitmap.inc}
{$I MemoryStream.inc}
{$I Font.inc}
{$I Strings.inc}
{$I StringList.inc}
{$I Brush.inc}
{$I Pen.inc}
{$I MenuItem.inc}
{$I ListGroups.inc}
{$I Picture.inc}
{$I ListColumns.inc}
{$I ListItems.inc}
{$I TreeNodes.inc}
{$I ListItem.inc}
{$I TreeNode.inc}
{$I PageControl.inc}
{$I TabSheet.inc}
{$I Control.inc}
{$I WinControl.inc}
{$I Screen.inc}
{$I Mouse.inc}
{$I ListGroup.inc}
{$I ListColumn.inc}
{$I CollectionItem.inc}
{$I StatusPanels.inc}
{$I StatusPanel.inc}
{$I SpinEdit.inc}
{$I MiniWebview.inc}
{$I Canvas.inc}
{$I Object.inc}
{$I PngImage.inc}
{$I JPEGImage.inc}
{$I GIFImage.inc}
{$I GIFFrame.inc}
{$I ActionList.inc}
{$I Action.inc}
{$I ToolButton.inc}
{$I IniFile.inc}
{$I Registry.inc}
{$I Clipboard.inc}
{$I Monitor.inc}
{$I Margins.inc}
{$I PaintBox.inc}
{$I Timer.inc}
{$I List.inc}
{$I Graphic.inc}
{$I Component.inc}
{$I MonthCalColors.inc}
{$I ParaAttributes.inc}
{$I TextAttributes.inc}
{$I IconOptions.inc}
{$I Exception.inc}
{$I ScrollBar.inc}
{$I CustomHint.inc}
{$I Shape.inc}
{$I Bevel.inc}
{$I ScrollBox.inc}
{$I CheckListBox.inc}
{$I Gauge.inc}
{$I ImageButton.inc}
{$I FindDialog.inc}
{$I ReplaceDialog.inc}
{$I PrinterSetupDialog.inc}
{$I PageSetupDialog.inc}
{$I DragObject.inc}
{$I DragDockObject.inc}
{$I StringGrid.inc}
{$I DrawGrid.inc}
{$I ValueListEditor.inc}
{$I HeaderControl.inc}
{$I HeaderSection.inc}
{$I HeaderSections.inc}
{$I LabeledEdit.inc}
{$I BoundLabel.inc}
{$I FlowPanel.inc}
{$I CoolBar.inc}
{$I CoolBands.inc}
{$I CoolBand.inc}
{$I JumpList.inc}
{$I JumpListItem.inc}
{$I Taskbar.inc}
{$I ThumbBarButtonList.inc}
{$I PreviewClipRegion.inc}
{$I ThumbBarButton.inc}
{$I JumpListCollection.inc}
{$I JumpCategories.inc}
{$I JumpCategoryItem.inc}
{$I Collection.inc}
{$I Printer.inc}

begin
end.
