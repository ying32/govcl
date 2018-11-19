

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
  Classes, SysUtils, Forms, Controls, Graphics, Dialogs;

type
   TReadyState = (rsUninitialized, rsLoading, rsLoaded, rsInterActive, rsComplete);
   
   TWebTitleChangeEvent = procedure(Sender: TObject; const Text: string) of object;

   { TMiniWebView }

   TMiniWebView = class(TWinControl)
   end;

implementation


{$ENDIF LINUX}
