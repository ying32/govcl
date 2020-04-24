unit Unit1;

{$mode delphi}

{$I ExtDecl.inc}

interface


uses
  Interfaces, // this includes the LCL widgetset,
{$IFDEF WINDOWS}
  Windows,
  MultiMon,
  ShellAPI,
  ShlObj,
  ComObj,
  ActiveX,
  Win32Int,
  win32proc,
  CommCtrl,
  LazUTF8,
  InterfaceBase,
{$ELSE}
  LCLType,
  Types,
{$ENDIF}
  typinfo,
  LCLProc,
  LCLIntf,
  LCLStrConsts,
  SysUtils,
  Classes,
  LMessages,
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
  Grids,
  ValEdit,
  Spin,
  fgl,
  Printers,
  ComboEx,
  uFormDesignerFile,
{$I UserDefineComponentUses.inc}
  uLinkLabel in 'uLinkLabel.pas',
  uEventCallback in 'uEventCallback.pas',
  ImageButton in 'ImageButton.pas',
  Gauges in 'Gauges.pas',
  uMiniWebview in 'uMiniWebview.pas',
  xButton in 'xButton.pas',
  uComponents,
  uGoForm,
  uRichEdit,
{$IFDEF LCLCocoa}
  uMacOSPatchs,
{$ENDIF}
{$IFDEF LINUX}
  uLinuxPatchs,
{$ENDIF}
  uControlPatchs;





implementation

{$I LazarusExtDef.inc}
{$I LazarusDef.inc}

{$I MyLCL_Object.inc}
{$I MyLCL_Graphic.inc}
{$I MyLCL_PngImage.inc}
{$I MyLCL_JPEGImage.inc}
{$I MyLCL_GIFImage.inc}
{$I MyLCL_ActionList.inc}
{$I MyLCL_Action.inc}
{$I MyLCL_ToolButton.inc}
{$I MyLCL_IniFile.inc}
{$I MyLCL_Registry.inc}
{$I MyLCL_Clipboard.inc}
{$I MyLCL_Monitor.inc}
{$I MyLCL_PaintBox.inc}
{$I MyLCL_Timer.inc}
{$I MyLCL_List.inc}
{$I MyLCL_Component.inc}
{$I MyLCL_ParaAttributes.inc}
{$I MyLCL_TextAttributes.inc}
{$I MyLCL_IconOptions.inc}
{$I MyLCL_Exception.inc}
{$I MyLCL_ScrollBar.inc}
{$I MyLCL_MaskEdit.inc}
{$I MyLCL_Shape.inc}
{$I MyLCL_Bevel.inc}
{$I MyLCL_ScrollBox.inc}
{$I MyLCL_CheckListBox.inc}
{$I MyLCL_Gauge.inc}
{$I MyLCL_ImageButton.inc}
{$I MyLCL_FindDialog.inc}
{$I MyLCL_ReplaceDialog.inc}
{$I MyLCL_PrinterSetupDialog.inc}
{$I MyLCL_PageSetupDialog.inc}
{$I MyLCL_DragObject.inc}
{$I MyLCL_DragDockObject.inc}
{$I MyLCL_StringGrid.inc}
{$I MyLCL_DrawGrid.inc}
{$I MyLCL_ValueListEditor.inc}
{$I MyLCL_HeaderControl.inc}
{$I MyLCL_HeaderSection.inc}
{$I MyLCL_HeaderSections.inc}
{$I MyLCL_LabeledEdit.inc}
{$I MyLCL_BoundLabel.inc}
{$I MyLCL_FlowPanel.inc}
{$I MyLCL_CoolBar.inc}
{$I MyLCL_CoolBands.inc}
{$I MyLCL_CoolBand.inc}
{$I MyLCL_Collection.inc}
{$I MyLCL_Printer.inc}
{$I MyLCL_TaskDialog.inc}
{$I MyLCL_TaskDialogButtons.inc}
{$I MyLCL_TaskDialogButtonItem.inc}
{$I MyLCL_TaskDialogRadioButtonItem.inc}
{$I MyLCL_TaskDialogBaseButtonItem.inc}
{$I MyLCL_ComboBoxEx.inc}
{$I MyLCL_ComboExItems.inc}
{$I MyLCL_ComboExItem.inc}
{$I MyLCL_Frame.inc}
{$I MyLCL_ControlScrollBar.inc}
{$I MyLCL_SizeConstraints.inc}
{$I MyLCL_XButton.inc}

end.

