unit uGoForm;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, LMessages;

type
    // 消息过程定义
  TWndProcEvent = procedure(Sender: TObject; var TheMessage: TLMessage; var AHandled: Boolean) of object;
  // 重定一个，主要是为了修改相关默认
  TGoForm = class(TForm)
  private
    FOnWndProc: TWndProcEvent;
  protected
    procedure ProcessResource; override;
    procedure WndProc(var TheMessage: TLMessage); override;
  public
    constructor CreateNew(AOwner: TComponent; Num: Integer = 0); override;
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

procedure TGoForm.ProcessResource;
begin
  // 没有使用窗口资源，不处理，处理就会报错的。
end;


procedure TGoForm.WndProc(var TheMessage: TLMessage);
var
  LHandled: Boolean;
begin
  LHandled := True;
  if Assigned(FOnWndProc) then
    FOnWndProc(Self, TheMessage, LHandled);
  if LHandled then
    inherited WndProc(TheMessage);
end;

end.

