unit uLinuxPatchs;

{$mode objfpc}{$H+}
{$I ExtDecl.inc}

interface

uses
  Classes, SysUtils, Forms, LCLType
{$IFDEF LCLGTK2}
  ,gtk2, Gtk2Proc, gdk2, glib2 {$IFDEF UNIX},Gdk2x, X {$ENDIF}
{$ENDIF}
{$IFDEF LClgtk3}
  ,LazGtk3, LazGObject2, LazGdk3, LazGLib2, gtk3procs, gtk3widgets, Lazxlib2
{$ENDIF}
  ;


  function GtkWidget_GetGtkFixed(Ah: HWND): PGtkFixed; extdecl;
  function GdkWindow_FromForm(AForm: TForm): PGdkWindow; extdecl;
  procedure GdkWindow_GetXId(AW: PGdkWindow; out AXId: TXId); extdecl;
  //function GtkWidget_Window(Ah: HWND): PGdkWindow; extdecl;
implementation

function GdkWindow_FromForm(AForm: TForm): PGdkWindow; extdecl;
begin
 Result := nil;
 if AForm = nil then
   Exit;
{$IFDEF LCLgtk2}
  Result := PGtkWidget(AForm.Handle)^.window;
{$ENDIF}
{$IFDEF LClgtk3}
  Result := TGtk3Widget(AForm.Handle).GetWindow;
{$ENDIF}
end;

function GtkWidget_GetGtkFixed(Ah: HWND): PGtkFixed; extdecl;
begin
 Result := nil;
 if Ah <= 0 then
   Exit;
{$IFDEF LCLgtk2}
  Result := PGtkFixed(GetFixedWidget(PGtkWidget(Ah)));
{$ENDIF}
{$IFDEF LClgtk3}
  Result := PGtkFixed(TGtk3Widget(Ah).GetContainerWidget);
{$ENDIF}
end;


procedure GdkWindow_GetXId(AW: PGdkWindow; out AXId: TXId); extdecl;
begin
{$IFDEF LCLGTK2}
  AXId := GDK_WINDOW_XID(AW);
  // GDK_WINDOW_XWINDOW
{$ENDIF}
{$IFDEF LClgtk3}

{$ENDIF}
end;


end.

