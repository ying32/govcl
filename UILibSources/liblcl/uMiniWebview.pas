//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
// 
// Licensed under Apache License 2.0
//
//----------------------------------------

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
     {$IF Defined(LCLgtk2) or Defined(LCLgtk3)}
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
