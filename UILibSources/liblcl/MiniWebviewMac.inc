

{*******************************************************}
{                                                       }
{       govcl MacOS WebBrowser                          }
{                                                       }
{       目前只支持cocoa                                 }
{                                                       }
{                                                       }
{       版权所有 (C) ying32                             }
{                                                       }
{*******************************************************}

{$IFDEF DARWIN}

{$mode objfpc}{$H+}
{$modeswitch objectivec1}

interface

uses
  Classes, SysUtils, Forms, Controls, Graphics, Dialogs,
  MacOSAll, CocoaAll, WebKit, CocoaUtils, CocoaPrivate,
  CocoaWSCommon;

type
   TReadyState = (rsUninitialized, rsLoading, rsLoaded, rsInterActive, rsComplete);
   
   TWebStatusTextChangeEvent = procedure(Sender: TObject; const AText: string) of object;
   TWebTitleChangeEvent = procedure(Sender: TObject; const AText: string) of object;
   TWebDocumentCompleteEvent = procedure(Sender: TObject; const pDisp: Pointer; const AURL: string) of object;
   TWebWindowClosingEvent = procedure(Sender: TObject; AIsChildWindow: Boolean; var ACancel: Boolean) of object;
   TWebJSExternalEvent = procedure(Sender: TObject; const AFunc: string; const AArgs: WideString; var ARetval: WideString) of object;

   { TWebviewUIDelegate }

   TWebviewUIDelegate = objcclass(NSObject) //external (WebUIDelegateCategory)
   private
     FWeb: Pointer;
   public
     procedure webView_runJavaScriptAlertPanelWithMessage_initiatedByFrame(sender: WebView; message: NSString; frame: WebFrame); message 'webView:runJavaScriptAlertPanelWithMessage:initiatedByFrame:'; override;
     procedure webView_didReceiveTitle_forFrame(sender: WebView; title: NSString; frame: WebFrame); message 'webView:didReceiveTitle:forFrame:'; override;
     procedure webView_didFinishLoadForFrame(sender: WebView; frame: WebFrame); message 'webView:didFinishLoadForFrame:'; override;
   end;



   { TMiniWebView }

   TMiniWebView = class(TCustomControl)
   private
     FWebview: Webview;
     FURL: string;
     FOnStatusTextChange: TWebStatusTextChangeEvent;
     FOnWindowClosing: TWebWindowClosingEvent;
     FOnOnDocumentComplete: TWebDocumentCompleteEvent;
     FOnTitleChange: TWebTitleChangeEvent;
     FOnJSExternal: TWebJSExternalEvent;

     function GetCurWebview: Webview;
   private
     procedure CreateCocoaWebView;
     function GetReadyState: TReadyState;
     procedure UpdateBounds;
   private
     property CurWebView: Webview read GetCurWebview;
   protected
     procedure VisibleChanged; override;
     procedure RealizeBounds; override;
     procedure Paint; override;
   public
     procedure Resize; override;
     procedure CreateWnd; override;

   public
     constructor Create(AOwner: TComponent); override;
     destructor Destroy; override;

     procedure Navigate(const AURL: string);
     procedure GoBack;
     procedure GoForward;
     procedure GoHome;
     procedure GoSearch;
     procedure Refresh;
     procedure Stop;

     procedure ExecuteScript(AScriptText, AScriptType: string);
     procedure ExecuteJS(AScriptText: string);

     procedure LoadHTML(const AStr: string);

     // propertys
     property ReadyState: TReadyState read GetReadyState;
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
     property OnTitleChange: TWebTitleChangeEvent read FOnTitleChange write FOnTitleChange;
     property OnDocumentComplete: TWebDocumentCompleteEvent read FOnOnDocumentComplete write FOnOnDocumentComplete;
     property OnWindowClosing: TWebWindowClosingEvent read FOnWindowClosing write FOnWindowClosing;
     property OnJSExternal: TWebJSExternalEvent read FOnJSExternal write FOnJSExternal;
   end;

implementation

{ TWebviewUIDelegate }

procedure TWebviewUIDelegate.webView_runJavaScriptAlertPanelWithMessage_initiatedByFrame(sender: WebView; message: NSString; frame: WebFrame);
begin
  ShowMessage(NSStringToString(message));
end;

procedure TWebviewUIDelegate.webView_didReceiveTitle_forFrame(sender: WebView; title: NSString; frame: WebFrame);
begin
  if (FWeb <> nil) and Assigned(TMiniWebView(FWeb).OnTitleChange) then
    TMiniWebView(Self.FWeb).OnTitleChange(TMiniWebView(FWeb), NSStringToString(title));
end;

procedure TWebviewUIDelegate.webView_didFinishLoadForFrame(sender: WebView; frame: WebFrame);
var
  context: id;
begin
  //contextm:= sender.valueForKeyPath(NSStringUtf8('documentView.webView.mainFrame.javasScriptContext'));
  //sender.mainFrame.globalContext;
end;

{ TMiniWebView }

constructor TMiniWebView.Create(AOwner: TComponent);
begin
  inherited Create(AOwner);
  ControlStyle := ControlStyle - [csAcceptsControls, csSetCaption];
  Width := 300;
  Height := 200;
  Visible := True;
  if csDesigning in ComponentState then
    Exit;
  CreateCocoaWebView;
  Navigate(FURL);
end;

destructor TMiniWebView.Destroy;
begin
  if not(csDesigning in ComponentState) then
  begin
    if FWebview <> nil then
      FWebview.release;
  end;
  inherited Destroy;
end;

function TMiniWebView.GetReadyState: TReadyState;
begin
  Result := rsUninitialized;  //?????
end;

procedure TMiniWebView.UpdateBounds;
var
  LR: NSRect;
begin
  if not(csDesigning in ComponentState) then
  begin
    if CurWebView <> nil then
    begin
      LR.origin.x := 0;
      LR.origin.y := 0;
      LR.size.width := Width;
      LR.size.height := Height;
      CurWebView.setFrame(LR);
    end;
  end;
end;

procedure TMiniWebView.Navigate(const AURL: string);
var
  AutoPool: NSAutoreleasePool;
  LReq: NSURLRequest;
begin
  if csDesigning in ComponentState then
    Exit;
  AutoPool := NSAutoreleasePool.alloc.init;
  try
    if CurWebview <> nil then
    begin
      if AURL = '' then
      begin
        FURL := '';
        Exit;
      end;
      FURL := AURL;
      LReq := NSURLRequest.requestWithURL(NSURL.URLWithString(NSStringUTF8(AURL)));
      CurWebview.mainFrame.loadRequest(LReq);
    end;
  finally
    AutoPool.release;
  end;
end;

procedure TMiniWebView.GoBack;
begin
  if CurWebview <> nil then
    CurWebview.goBack;
end;

procedure TMiniWebView.GoForward;
begin
  if CurWebview <> nil then
    CurWebview.goForward;
end;

procedure TMiniWebView.GoHome;
begin
  // no support
end;

procedure TMiniWebView.GoSearch;
begin
  // no support
end;

procedure TMiniWebView.Refresh;
begin
  if CurWebview <> nil then
    CurWebview.mainFrame.reload;
end;

procedure TMiniWebView.Stop;
begin
  if CurWebview <> nil then
    CurWebview.mainFrame.stopLoading;
end;

// AScriptType no support
procedure TMiniWebView.ExecuteScript(AScriptText, AScriptType: string);
begin
  if not HandleAllocated then
    HandleNeeded;
  if CurwebView = nil then
    Exit;
  Curwebview.stringByEvaluatingJavaScriptFromString(NSStringUtf8(AScriptText));
end;

procedure TMiniWebView.ExecuteJS(AScriptText: string);
begin
  ExecuteScript(AScriptText, 'javascript');
end;

procedure TMiniWebView.LoadHTML(const AStr: string);
begin
  if not HandleAllocated then
    HandleNeeded;
  if Curwebview = nil then
    Exit;
  Curwebview.mainFrame.loadHTMLString_baseURL(NSStringUtf8(AStr), nil);
end;

procedure TMiniWebView.CreateCocoaWebView;
var
  AutoPool: NSAutoreleasePool;
  LR: NSRect;
  LDelegate: TWebviewUIDelegate;
begin
  if csDesigning in ComponentState then
    Exit;
  AutoPool := NSAutoreleasePool.alloc.init;
  try
    LR.origin.x:=0;
    LR.origin.y:=0;
    LR.size.width := Width;
    LR.size.height := Height;
    FWebview := WebView.alloc.initWithFrame(LR);
    LDelegate := TWebviewUIDelegate.alloc.init;
    LDelegate.FWeb := Self;
    FWebview.setUIDelegate(LDelegate);
    FWebview.setFrameLoadDelegate(LDelegate);
  finally
    AutoPool.release;
  end;
end;


function TMiniWebView.GetCurWebview: Webview;
begin
  Result := FWebView;
end;

procedure TMiniWebView.VisibleChanged;
begin
  inherited VisibleChanged;
  if csDesigning in ComponentState then
    Exit;
  if FWebview <> nil then
    FWebView.setHidden(not Visible);
end;

procedure TMiniWebView.RealizeBounds;
begin
  inherited RealizeBounds;
end;


procedure TMiniWebView.Resize;
begin
  inherited Resize;
  UpdateBounds;
end;

procedure TMiniWebView.CreateWnd;
var
  view: NSView;
  AutoPool: NSAutoreleasePool;
begin
  inherited CreateWnd;
  if csDesigning in ComponentState then
    Exit;
  AutoPool := NSAutoreleasePool.alloc.init;
  try
    view := NSView(Self.Handle).lclContentView;//GetNSObjectView(NSObject(Handle));
    //writeln(NSObjectDebugStr(view));
    view.addSubview(FWebview);
    UpdateBounds;
  finally
    AutoPool.release;
  end;
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


{$ENDIF DARWIN}
