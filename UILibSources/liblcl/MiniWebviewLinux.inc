

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
  Classes, SysUtils, Forms, Controls, Graphics, Dialogs, gtk2, Gdk2, Glib2, dynlibs;

type
   TReadyState = (rsUninitialized, rsLoading, rsLoaded, rsInterActive, rsComplete);

   TWebTitleChangeEvent = procedure(Sender: TObject; const Text: string) of object;
   TWebStatusTextChangeEvent = procedure(Sender: TObject; const Text: string) of object;
   TWebJSExternalEvent = procedure(Sender: TObject; const func: string; const args: WideString; var retval: WideString) of object;

   { TMiniWebview }

   WebKitWebView = PGtkWidget;

   TMiniWebview = class(TWinControl)
   private
     FOnJSExternal: TWebJSExternalEvent;
     FOnStatusTextChange: TWebStatusTextChangeEvent;
     FOnTitleChange: TWebTitleChangeEvent;
     FWebView: WebKitWebView;
     FScrolledWindow: PGtkWidget;
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
  libwebkitgtk1_0 = 'libwebkitgtk-1.0.so.0';

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
  LibWebkitGtkHandle := dynlibs.SafeLoadLibrary(libwebkitgtk1_0);
  if LibWebkitGtkHandle <> 0 then
  begin
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
  gtk_scrolled_window_set_policy(GTK_SCROLLED_WINDOW(FScrolledWindow), GTK_POLICY_AUTOMATIC, GTK_POLICY_AUTOMATIC);
  gtk_container_add(GTK_CONTAINER(FScrolledWindow), GTK_WIDGET(FWebView));
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
    gtk_widget_set_usize(FScrolledWindow, Width, Height);
end;

procedure TMiniWebview.CreateWnd;
var
  LChild: PGtkWidget;
begin
  inherited CreateWnd;
  if (Handle <> 0) and (FWebView <> nil) and (FScrolledWindow <> nil) and GTK_IS_BIN(Pointer(Handle)) then
  begin
    LChild := gtk_bin_get_child(GTK_BIN(Pointer(Handle)));
    if LChild <> nil then
    begin
      gtk_container_add(GTK_CONTAINER(LChild), GTK_WIDGET(FScrolledWindow));
      gtk_widget_show_all(LChild);
    end;
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
