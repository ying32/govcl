{$ifdef windows}
unit uExport2;

{$mode delphi}

{$I ExtDecl.inc}

interface

implementation

uses
  Classes, SysUtils,
  {$I UseAll.inc},
  uControlPatchs, uExportTable;
  
{$endif windows}

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
{$I MyLCL_AnchorSide.inc}
{$I MyLCL_ControlBorderSpacing.inc}
{$I MyLCL_ControlChildSizing.inc}


{$ifdef windows}
end.
{$endif windows}
