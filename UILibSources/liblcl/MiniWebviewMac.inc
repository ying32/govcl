

{*******************************************************}
{                                                       }
{       govcl MacOS WebBrowser                          }
{                                                       }
{       目前只对cocoa框架友好，carbon也能将就，但只能   }
{       使用Root也是就Form作为父容器                    }
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
  MacOSAll, CocoaAll, WebKit,
  {$IFDEF LCLcocoa}
    CocoaUtils
  {$ELSE}
    //CarbonDef,
    CarbonProc,
    CarbonPrivate
  {$ENDIF};

type

   { TMiniWebView }

   TMiniWebView = class(TWinControl)
   private
   {$IFDEF LCLCocoa}
     FWebview: Webview;
   {$ENDIF}
     FURL: string;
     function GetCurWebview: Webview;
   private
   {$IFDEF LCLCocoa}
     procedure CreateCocoaWebView;
   {$ELSE}
     FCarbonWeb: HIViewRef;
     procedure CreateCarbonWebView;
   {$ENDIF}

     property CurWebView: Webview read GetCurWebview;
   protected
     procedure VisibleChanged; override;
     procedure RealizeBounds; override;

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
   end;

implementation

{$IFDEF LCLCarbon}
  function HIWebViewCreate(var outControl: HIViewRef): OSStatus; cdecl; external name 'HIWebViewCreate';
  function HIWebViewGetWebView(inView: HIViewRef): WebView; cdecl; external name 'HIWebViewGetWebView';
  procedure WebInitForCarbon; cdecl; external name 'WebInitForCarbon';

  // LCLCocoa
  function NSStringUTF8(AStr: string): NSString;
  var
    cf: CFStringRef;
  begin
    cf := CFStringCreateWithCString(nil, Pointer(PChar(AStr)), kCFStringEncodingUTF8);
    Result := NSString(cf);
  end;
{$ENDIF}


constructor TMiniWebView.Create(AOwner: TComponent);
begin
  inherited Create(AOwner);
  Width := 300;
  Height := 200;
  Visible := True;
{$IFDEF LCLCarbon}
  WebInitForCarbon;
{$ENDIF}
{$IFDEF LCLCocoa}
  CreateCocoaWebView;
{$ENDIF}
{$IFDEF LCLCarbon}
  CreateCarbonWebView;
{$ENDIF}
  Navigate(FURL);
end;

destructor TMiniWebView.Destroy;
begin
{$IFDEF LCLCocoa}
  if FWebview <> nil then
    FWebview.release;
{$ENDIF}
{$IFDEF LCLCarbon}
  if FCarbonWeb <> nil then
    CFRelease(FCarbonWeb);
{$ENDIF}
  inherited Destroy;
end;

procedure TMiniWebView.Navigate(const AURL: string);
var
  AutoPool: NSAutoreleasePool;
  LReq: NSURLRequest;
begin
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
end;

procedure TMiniWebView.GoSearch;
begin
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

{$IFDEF LCLCocoa}
procedure TMiniWebView.CreateCocoaWebView;
var
  AutoPool: NSAutoreleasePool;
  LR: NSRect;
  view: NSView;
begin
  AutoPool := NSAutoreleasePool.alloc.init;
  try
    LR.origin.x:=0;
    LR.origin.y:=0;
    LR.size.width := Width;
    LR.size.height := Height;
    FWebview := WebView.alloc.initWithFrame(LR);
  finally
    AutoPool.release;
  end;
end;
{$ENDIF}

{$IFDEF LCLCarbon}
procedure TMiniWebView.CreateCarbonWebView;
begin
  HIWebViewCreate(FCarbonWeb);
end;
{$ENDIF}

function TMiniWebView.GetCurWebview: Webview;
begin
  Result := nil;
{$IFDEF LCLCocoa}
   Result := FWebView;
{$ENDIF}
{$IFDEF LCLCarbon}
  if FCarbonWeb <> nil then
    Result := HIWebViewGetWebView(FCarbonWeb);
{$ENDIF}
end;

procedure TMiniWebView.VisibleChanged;
begin
  inherited VisibleChanged;
{$IFDEF LCLCocoa}
  if FWebview <> nil then
    FWebView.setHidden(Visible);
{$ENDIF}
{$IFDEF LCLCarbon}
  //if FCarbonWeb <> nil then
    //HIViewSetVisible(FCarbonWeb, Visible);
{$ENDIF}
end;

procedure TMiniWebView.RealizeBounds;
begin
  inherited RealizeBounds;
  {$IFDEF LCLCarbon}
    //if FCarbonWeb <> nil then
    //  HIViewSetNeedsDisplay(FCarbonWeb, True);
    //Self.Invalidate;
  {$ENDIF}
end;


procedure TMiniWebView.Resize;
{$IFDEF LCLCocoa}
var
  LR: NSRect;
{$ENDIF}
begin
  inherited Resize;
{$IFDEF LCLCocoa}
  if CurWebView <> nil then
  begin
    LR.origin.x := 0;
    LR.origin.y := 0;
    LR.size.width := Width;
    LR.size.height := Height;
    CurWebView.setFrame(LR);
  end;
{$ENDIF}

{$IFDEF LCLCarbon}
  if FCarbonWeb <> nil then
  begin
    HIViewSetFrame(FCarbonWeb, RectToCGRect(Self.BoundsRect));
  end;
{$ENDIF}
end;

procedure TMiniWebView.CreateWnd;
{$IFDEF LCLCocoa}
var
  view: NSView;
  AutoPool: NSAutoreleasePool;
{$ENDIF}
{$IFDEF LCLCarbon}
 var
  LContentView: HIViewRef;
  LBounds: HIRect;
{$ENDIF}
begin
  inherited CreateWnd;
{$IFDEF LCLCocoa}
  AutoPool := NSAutoreleasePool.alloc.init;
  try
    view := GetNSObjectView(NSObject(Handle));
    view.addSubview(FWebview);
    //view.setHidden(false);
  finally
    AutoPool.release;
  end;
{$ENDIF}
{$IFDEF LCLCarbon}
  //LContentView := TCarbonControl(Handle).Content;
  HIViewFindByID(HIViewGetRoot(TCarbonWindow(Parent.Handle).Window), kHIViewWindowContentID, LContentView);
  if LContentView = nil then
    Exit;

  //FCarbonWidget.AddToWidget(TCarbonControl(Handle));
  HIViewGetBounds(LContentView, LBounds);
                      //showmessagefmt('x=%f,y=%f,w=%f,h=%f',
                      //[lbounds.origin.x, lbounds.origin.y, lbounds.size.width,lbounds.size.height]);
  HIViewAddSubview(LContentView, FCarbonWeb);

  HIViewSetFrame(FCarbonWeb, LBounds);
  HIViewSetVisible(FCarbonWeb, True);
{$ENDIF}
end;


{$ENDIF DARWIN}
