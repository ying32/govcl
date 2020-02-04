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
     implementation
  {$ENDIF}
{$ELSEIF Defined(LINUX)}
  {$I MiniWebviewLinux.inc}
{$ELSE}
interface
implementation
{$ENDIF}

end.
