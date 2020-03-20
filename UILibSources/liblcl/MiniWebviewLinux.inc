

{*******************************************************}
{                                                       }
{       govcl Linux WebBrowser                          }
{                                                       }
{                                                       }
{                                                       }
{                                                       }
{                                                       }
{       版权所有 (C) ying32                             }
{                                                       }
{*******************************************************}


{$IFDEF LINUX}

//{$IFDEF FPC}
//  {$mode delphi}
//{$ENDIF}
interface

uses
  Classes, SysUtils, Forms, Controls, Graphics, Dialogs,
{$IFDEF LCLgtk2}
  gtk2, Gdk2, Glib2, Gtk2Proc
{$ENDIF}
{$IFDEF LClgtk3}
  LazGtk3, LazGObject2, LazGLib2, gtk3procs, gtk3widgets// ,LazWebKit3
{$ENDIF}
  , dynlibs;

type
   TReadyState = (rsUninitialized, rsLoading, rsLoaded, rsInterActive, rsComplete);

   TWebTitleChangeEvent = procedure(Sender: TObject; const Text: string) of object;
   TWebStatusTextChangeEvent = procedure(Sender: TObject; const Text: string) of object;
   TWebJSExternalEvent = procedure(Sender: TObject; const func: string; const args: WideString; var retval: WideString) of object;

   { TMiniWebview }

   WebKitWebView = PGtkWidget;

   TMiniWebview = class({$IFDEF LCLgtk3}TCustomControl{$ELSE}TWinControl{$ENDIF})
   private
     FOnJSExternal: TWebJSExternalEvent;
     FOnStatusTextChange: TWebStatusTextChangeEvent;
     FOnTitleChange: TWebTitleChangeEvent;
     FWebView: WebKitWebView;
   {$IFDEF LCLgtk2}
     FScrolledWindow: PGtkWidget;
   {$ENDIF}
   {$IFDEF LCLgtk3}
     FScrolledWindow: PGtkScrolledWindow;
   {$ENDIF}
     function GetReadyState: TReadyState;
   public
     constructor Create(AOwner: TComponent); override;
     destructor Destroy; override;
     procedure Resize; override;
     procedure CreateWnd; override;

     procedure Navigate(const AURL: string);
     procedure GoBack;
     procedure GoForward;
     procedure GoHome;  // not support
     procedure GoSearch;  // not support
     procedure Refresh;
     procedure Stop;

     procedure ExecuteScript(AScriptText, AScriptType: string);  // not support
     procedure ExecuteJS(AScriptText: string);   // not support

     procedure LoadHTML(const AStr: string);

     // propertys
     property ReadyState: TReadyState read GetReadyState;   // not support
     // events
     property OnStatusTextChange: TWebStatusTextChangeEvent read FOnStatusTextChange write FOnStatusTextChange;  // no support????
     property OnTitleChange: TWebTitleChangeEvent read FOnTitleChange write FOnTitleChange;  // not support
     //property OnDocumentComplete: TWebDocumentCompleteEvent read FOnOnDocumentComplete write FOnOnDocumentComplete;
     //property OnWindowClosing: TWebWindowClosingEvent read FOnWindowClosing write FOnWindowClosing;
     property OnJSExternal: TWebJSExternalEvent read FOnJSExternal write FOnJSExternal;// not support
   end;

implementation

uses
  Math;

const
{$IF Defined(LCLgtk2)}
  libwebkitgtkso = 'libwebkitgtk-1.0.so.0';
{$ELSEIF Defined(LCLGtk3) }
  libwebkitgtkso = 'libwebkitgtk-3.0.so.0';
{$ELSE}
  libwebkitgtkso = '';
{$ENDIF}

var
  LibWebkitGtkHandle: TLibHandle;

  webkit_web_view_new: function: PGtkWidget; cdecl;
  webkit_web_view_load_uri: procedure(webview: WebKitWebView; aurl: Pgchar); cdecl;
  webkit_web_view_get_title: function(webview: WebKitWebView): Pgchar; cdecl;
  webkit_web_view_can_go_back: function (webview: WebKitWebView): gboolean; cdecl;
  webkit_web_view_go_back: procedure(webview: WebKitWebView); cdecl;
  webkit_web_view_can_go_forward: function(webview: WebKitWebView): gboolean; cdecl;
  webkit_web_view_go_forward: procedure(webview: WebKitWebView); cdecl;
  webkit_web_view_reload: procedure(webview: WebKitWebView); cdecl;
  webkit_web_view_stop_loading: procedure(webview: WebKitWebView); cdecl;
  webkit_web_view_is_loading: function(webview: WebKitWebView): gboolean; cdecl;
  webkit_web_view_load_html: procedure(webview: WebKitWebView; content, base_uri: Pgchar); cdecl;

function WEBKIT_WEB_VIEW(webview: PGtkWidget): WebKitWebView; inline;
begin
  Result := WebKitWebView(webview);
end;

procedure LoadLibWebkitGtk;
begin
  if LibWebkitGtkHandle <> 0 then
    Exit;
  LibWebkitGtkHandle := dynlibs.SafeLoadLibrary(libwebkitgtkso);
  if LibWebkitGtkHandle <> 0 then
  begin
    {$IFDEF LCLgtk3}
      //dynlibs.SafeLoadLibrary('libjavascriptcoregtk-3.0.so.0');
    {$ENDIF}
    Pointer(webkit_web_view_new) := dynlibs.GetProcAddress(LibWebkitGtkHandle, 'webkit_web_view_new');
    Pointer(webkit_web_view_load_uri) := dynlibs.GetProcAddress(LibWebkitGtkHandle, 'webkit_web_view_load_uri');
    Pointer(webkit_web_view_get_title) := dynlibs.GetProcAddress(LibWebkitGtkHandle, 'webkit_web_view_get_title');
    Pointer(webkit_web_view_can_go_back) := dynlibs.GetProcAddress(LibWebkitGtkHandle, 'webkit_web_view_can_go_back');
    Pointer(webkit_web_view_go_back) := dynlibs.GetProcAddress(LibWebkitGtkHandle, 'webkit_web_view_go_back');
    Pointer(webkit_web_view_can_go_forward) := dynlibs.GetProcAddress(LibWebkitGtkHandle, 'webkit_web_view_can_go_forward');
    Pointer(webkit_web_view_go_forward) := dynlibs.GetProcAddress(LibWebkitGtkHandle, 'webkit_web_view_go_forward');
    Pointer(webkit_web_view_reload) := dynlibs.GetProcAddress(LibWebkitGtkHandle, 'webkit_web_view_reload');
    Pointer(webkit_web_view_stop_loading) := dynlibs.GetProcAddress(LibWebkitGtkHandle, 'webkit_web_view_stop_loading');
    Pointer(webkit_web_view_is_loading) := dynlibs.GetProcAddress(LibWebkitGtkHandle, 'webkit_web_view_is_loading');
    Pointer(webkit_web_view_load_html) := dynlibs.GetProcAddress(LibWebkitGtkHandle, 'webkit_web_view_load_html');
  end;
end;

procedure FreeLibWebkitGtk;
begin
  if LibWebkitGtkHandle <> 0 then
  begin
    dynlibs.UnloadLibrary(LibWebkitGtkHandle);
    LibWebkitGtkHandle := 0;
  end;
end;

{ TMiniWebview }

constructor TMiniWebview.Create(AOwner: TComponent);
begin
  inherited Create(AOwner);
  if csDesigning in ComponentState then
    Exit;
  LoadLibWebkitGtk;
  FWebView := WEBKIT_WEB_VIEW(webkit_web_view_new());
  if FWebView = nil then
    raise Exception.Create('WebView Create Failed.');
  FScrolledWindow := gtk_scrolled_window_new(nil, nil);
{$IFDEF LCLgtk2}
  gtk_scrolled_window_set_policy(GTK_SCROLLED_WINDOW(FScrolledWindow), GTK_POLICY_AUTOMATIC, GTK_POLICY_AUTOMATIC);
  gtk_container_add(GTK_CONTAINER(FScrolledWindow), GTK_WIDGET(FWebView));
{$ENDIF}
{$IFDEF LClgtk3}
  gtk_scrolled_window_set_policy(FScrolledWindow, GTK_POLICY_AUTOMATIC, GTK_POLICY_AUTOMATIC);
  gtk_container_add(FScrolledWindow, FWebView);
{$ENDIF}
end;

destructor TMiniWebview.Destroy;
begin
  //if FWebView <> nil then
     //gtk_widget_destroy(FWebView);
  //if FScrolledWindow <> nil then
     //gtk_widget_destroy(FScrolledWindow);
  inherited Destroy;
end;

procedure TMiniWebview.Resize;
begin
  inherited Resize;
  if FScrolledWindow <> nil then
{$IFDEF LCLgtk3}
  gtk_widget_set_size_request(FScrolledWindow, Width, Height);
{$ENDIF}
{$IFDEF LCLgtk2}
  gtk_widget_set_usize(FScrolledWindow, Width, Height);
{$ENDIF}
end;

procedure TMiniWebview.CreateWnd;
var
  LChild: PGtkWidget;
begin
  inherited CreateWnd;
  if (Handle <> 0) and (FWebView <> nil) and (FScrolledWindow <> nil) then
  begin
{$IFDEF LCLgtk2}
  LChild := GetFixedWidget(PGtkWidget(Handle));
  FixedPutControl(LChild, FScrolledWindow, 0, 0);
  gtk_widget_show_all(LChild);
{$ENDIF}
{$IFDEF LCLgtk3}
  PGtkFixed(TGtk3Widget(Handle).GetContainerWidget)^.put(FScrolledWindow, 0, 0);
{$ENDIF}
  end;
end;

function TMiniWebview.GetReadyState: TReadyState;
begin
  Result := rsUninitialized;
end;

procedure TMiniWebview.Navigate(const AURL: string);
begin
  if Assigned(FWebView) then
  begin
    webkit_web_view_load_uri(FWebView, PChar(AURL));
    //gtk_widget_grab_focus(GTK_WIDGET(FWebView));
  end;
end;

procedure TMiniWebview.GoBack;
begin
  if Assigned(FWebview) and webkit_web_view_can_go_back(FWebview) then
     webkit_web_view_go_back(FWebview);
end;

procedure TMiniWebview.GoForward;
begin
  if Assigned(FWebview) and webkit_web_view_can_go_forward(FWebview) then
     webkit_web_view_go_forward(FWebview);
end;

procedure TMiniWebview.GoHome;
begin

end;

procedure TMiniWebview.GoSearch;
begin

end;

procedure TMiniWebview.Refresh;
begin
  if Assigned(FWebview) then
     webkit_web_view_reload(FWebview);
end;

procedure TMiniWebview.Stop;
begin
  if Assigned(FWebview) and webkit_web_view_is_loading(FWebview) then
     webkit_web_view_stop_loading(FWebview);
end;

procedure TMiniWebview.ExecuteScript(AScriptText, AScriptType: string);
begin

end;

procedure TMiniWebview.ExecuteJS(AScriptText: string);
begin

end;

procedure TMiniWebview.LoadHTML(const AStr: string);
begin
  //if Assigned(FWebview) and (Length(AStr) > 0) then
    //webkit_web_view_load_html(FWebview, PChar(AStr), 'about:blank');
end;

initialization
    // 屏蔽浮点异常
  SetExceptionMask([exInvalidOp..exPrecision]);

finalization
  FreeLibWebkitGtk;

{$ENDIF LINUX}
