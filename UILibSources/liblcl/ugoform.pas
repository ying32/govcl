//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
// 
// Licensed under Apache License 2.0
//
//----------------------------------------

unit uGoForm;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, LMessages, LCLType;

type
    // 消息过程定义
  TWndProcEvent = procedure(Sender: TObject; var TheMessage: TLMessage) of object;
  // 重定一个，主要是为了修改相关默认

  { TGoForm }

  TGoForm = class(TForm)
  private
    FOnWndProc: TWndProcEvent;
  protected
    procedure ProcessResource; override;
    procedure WndProc(var TheMessage: TLMessage); override;
  public
    constructor CreateNew(AOwner: TComponent; Num: Integer = 0); override;
    procedure ScaleForPPI(ANewPPI: Integer);
    procedure ScaleForCurrentDpi;
    procedure InheritedWndProc(var TheMessage: TLMessage);
  published
    property OnWndProc: TWndProcEvent read FOnWndProc write FOnWndProc;
  end;


implementation

constructor TGoForm.CreateNew(AOwner: TComponent; Num: Integer);
begin
  inherited CreateNew(AOwner, Num);
{$IFDEF WINDOWS}
  // Windows下统一VCL/LCL默认字体
  Font.Name := 'Tahoma';
  Font.Size := 8;
{$ENDIF}
end;

procedure TGoForm.ScaleForPPI(ANewPPI: Integer);
begin
  if ANewPPI < 30 then
    Exit;
  if ANewPPI <> PixelsPerInch then
  begin
    AutoAdjustLayout(lapAutoAdjustForDPI, PixelsPerInch, ANewPPI,
      MulDiv(Width, ANewPPI, PixelsPerInch),
      MulDiv(Height, ANewPPI, PixelsPerInch));
  end;
end;

procedure TGoForm.ScaleForCurrentDpi;
begin
  if not Scaled then
  begin
    Scaled := True;
    Exit;
  end;
  if PixelsPerInch <> Monitor.PixelsPerInch then
  begin
    AutoAdjustLayout(lapAutoAdjustForDPI, PixelsPerInch, Monitor.PixelsPerInch,
      MulDiv(Width, Monitor.PixelsPerInch, PixelsPerInch),
      MulDiv(Height, Monitor.PixelsPerInch, PixelsPerInch));
  end;
end;

procedure TGoForm.InheritedWndProc(var TheMessage: TLMessage);
begin
  inherited WndProc(TheMessage);
end;

procedure TGoForm.ProcessResource;
begin
  Self.ClientHeight := 321;
  Self.ClientWidth := 678;
  // 没有使用窗口资源，不处理，处理就会报错的。
end;


procedure TGoForm.WndProc(var TheMessage: TLMessage);
begin
  if Assigned(FOnWndProc) then
    FOnWndProc(Self, TheMessage)
  else
    inherited WndProc(TheMessage);
end;



end.

