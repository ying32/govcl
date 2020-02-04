{*******************************************************}
{                                                       }
{       govcl windows WebBrowser                        }
{                                                       }
{       根据VCL原有的TWebBrowser精简的一个，用于在      }
{       Lazarus中也能使用                               }
{                                                       }
{       版权所有 (C) ying32                             }
{                                                       }
{*******************************************************}

unit uMiniWebview;

{$IF Defined(MSWINDOWS)}
  {$I MiniWebviewWin.inc}
{$ELSEIF Defined(DARWIN) or Defined(MACOS)}
  {$IFDEF LCLcocoa}
    {$I MiniWebviewMac.inc}
  {$ELSE}
     interface
       type TMiniWebview = class end;
     implementation
  {$ENDIF}
{$ELSEIF Defined(LINUX)}
     {$IFDEF LCLgtk2}
        {$I MiniWebviewLinux.inc}
     {$ELSE}
        interface
          type TMiniWebview = class end;
        implementation
     {$ENDIF}
{$ELSE}
interface
  type TMiniWebview = class end;
implementation
{$ENDIF}

end.
