

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

   TWebStatusTextChangeEvent = procedure(Sender: TObject; const AText: string) of object;
   TWebTitleChangeEvent = procedure(Sender: TObject; const AText: string) of object;
   TWebDocumentCompleteEvent = procedure(Sender: TObject; const pDisp: Pointer; const AURL: string) of object;
   TWebWindowClosingEvent = procedure(Sender: TObject; AIsChildWindow: Boolean; var ACancel: Boolean) of object;
   TWebJSExternalEvent = procedure(Sender: TObject; const AFunc: string; const AArgs: WideString; var ARetval: WideString) of object;

   { TMiniWebview }

   WebKitWebView = PGtkWidget;

   TMiniWebview = class(TCustomControl)
   private
     FOnStatusTextChange: TWebStatusTextChangeEvent;
     FOnWindowClosing: TWebWindowClosingEvent;
     FOnOnDocumentComplete: TWebDocumentCompleteEvent;
     FOnTitleChange: TWebTitleChangeEvent;
     FOnJSExternal: TWebJSExternalEvent;
     FWebView: WebKitWebView;
   {$IFDEF LCLgtk2}
     FScrolledWindow: PGtkWidget;
   {$ENDIF}
   {$IFDEF LCLgtk3}
     FScrolledWindow: PGtkScrolledWindow;
   {$ENDIF}
     function GetReadyState: TReadyState;
   protected
     procedure Paint; override;
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

     function ExecuteScript(AScriptText, AScriptType: string): string;  // not support
     function ExecuteJS(AScriptText: string): string;   // not support

     procedure LoadHTML(const AStr: string);

     // propertys
     property ReadyState: TReadyState read GetReadyState;   // not support
   published
     property Align;
     property Anchors;
     property BorderSpacing;
     property ChildSizing;
     property Constraints;
     property Enabled;
     property Visible;
     // events
     property OnStatusTextChange: TWebStatusTextChangeEvent read FOnStatusTextChange write FOnStatusTextChange;  // no support????
     property OnTitleChange: TWebTitleChangeEvent read FOnTitleChange write FOnTitleChange;  // not support
     property OnDocumentComplete: TWebDocumentCompleteEvent read FOnOnDocumentComplete write FOnOnDocumentComplete;
     property OnWindowClosing: TWebWindowClosingEvent read FOnWindowClosing write FOnWindowClosing;
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
  // gtk2: webkit_web_view_load_html_string
  webkit_web_view_load_html: procedure(webview: WebKitWebView; content, base_uri: Pgchar); cdecl;
{$ifdef LCLGtk2}
  webkit_web_view_execute_script: procedure(web_view: WebKitWebView; script: Pgchar); cdecl;
{$endif}

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
    Pointer(webkit_web_view_load_html) := dynlibs.GetProcAddress(LibWebkitGtkHandle, {$ifdef LCLgtk3}'webkit_web_view_load_html'{$else}'webkit_web_view_load_html_string'{$endif});
  {$ifdef LCLGtk2}
    Pointer(webkit_web_view_execute_script) := dynlibs.GetProcAddress(LibWebkitGtkHandle, 'webkit_web_view_execute_script');
  {$endif}
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

// event
procedure doWebviewTitleChanged(web_view: WebKitWebView; frame: {WebKitWebFrame}Pointer; title: Pgchar; user_data: gpointer); cdecl;
begin
  if (user_data <> nil) and Assigned(TMiniWebview(user_data).OnTitleChange) then
    TMiniWebview(user_data).OnTitleChange(TObject(user_data), title);
end;

{ TMiniWebview }

constructor TMiniWebview.Create(AOwner: TComponent);
begin
  inherited Create(AOwner);
  ControlStyle := ControlStyle - [csAcceptsControls, csSetCaption];
  Width := 300;
  Height := 200;
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
  g_signal_connect(FWebView, 'title-changed', G_CALLBACK(@doWebviewTitleChanged), Self);
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
  if csDesigning in ComponentState then
    Exit;
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
  if csDesigning in ComponentState then
    Exit;
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
    webkit_web_view_load_uri(FWebView, PChar(AURL));
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

function TMiniWebview.ExecuteScript(AScriptText, AScriptType: string): string;
begin
  Result := '';
{$ifdef LCLGtk2}
  if Assigned(webkit_web_view_execute_script) and Assigned(FWebview) and (Length(AScriptText) > 0) then
    webkit_web_view_execute_script(FWebview, Pgchar(AScriptText));
{$endif}
end;

function TMiniWebview.ExecuteJS(AScriptText: string): string;
begin
  Result := ExecuteScript(AScriptText, '');
end;

procedure TMiniWebview.LoadHTML(const AStr: string);
begin
  if Assigned(webkit_web_view_load_html) and Assigned(FWebview) and (Length(AStr) > 0)  then
    webkit_web_view_load_html(FWebview, PChar(AStr), 'about:blank');
end;

procedure TMiniWebview.Paint;
begin
  inherited Paint;
  if csDesigning in ComponentState then
  begin
    Canvas.Brush.Color:=clWhite;
    Canvas.FillRect(ClientRect);
  end;
end;

initialization
    // 屏蔽浮点异常
  SetExceptionMask([exInvalidOp..exPrecision]);

finalization
  FreeLibWebkitGtk;

{$ENDIF LINUX}
