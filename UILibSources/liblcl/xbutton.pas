//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
// 
// Licensed under Apache License 2.0
//
//----------------------------------------
unit xButton;

{$IFDEF FPC}
{$mode objfpc}{$H+}
{$ENDIF}

interface

uses
  Classes, SysUtils, Controls, Graphics
  {$IFNDEF FPC}
    ,Winapi.Messages, System.Types
  {$ELSE}
    ,LMessages
  {$ENDIF};

type

{$IFDEF FPC}
  TMessage = TLMessage;
{$ENDIF}

  { TXButton }

  TDrawImageMode = (dimNormal, dimCenter, dimStretch);

  TXButtonState = (xbsNone, xbsHot, xbsDown, xbsDisabled);

  TXButton = class(TGraphicControl)
  private
    FCaption: string;
    FDownFontColor: TColor;
    FDrawMode: TDrawImageMode;
    FHoverFontColor: TColor;
    FNormalFontColor: TColor;
    FShowCaption: Boolean;
    FState: TXButtonState;
    FBackColor: TColor;
    FBorderColor: TColor;
    FBorderWidth: Integer;
    FDownColor: TColor;
    FHoverColor: TColor;
    FPicture: TPicture;
    FBufferBmp: TBitmap;

    FMouseInControl: Boolean;

    procedure StateChanged;
    procedure DoPictureChange(Sender: TObject);
    procedure DrawBmp;

    procedure SetBackColor(AValue: TColor);
    procedure SetBorderColor(AValue: TColor);
    procedure SetBorderWidth(AValue: Integer);
    procedure SetCaption(AValue: string);
    procedure SetDownColor(AValue: TColor);
    procedure SetDownFontColor(AValue: TColor);
    procedure SetDrawMode(AValue: TDrawImageMode);
    procedure SetHoverColor(AValue: TColor);
    procedure SetHoverFontColor(AValue: TColor);
    procedure SetNormalFontColor(AValue: TColor);
    procedure SetPicture(AValue: TPicture);
    procedure SetShowCaption(AValue: Boolean);
  protected
    procedure CMMouseEnter(var Message: TMessage); message CM_MOUSEENTER;
    procedure CMMouseLeave(var Message: TMessage); message CM_MOUSELEAVE;
    procedure CMEnabledChanged(var Message: TMessage); message CM_ENABLEDCHANGED;
    procedure CMFontChanged(var Message: TMessage); message CM_FONTCHANGED;


    procedure MouseDown(Button: TMouseButton; Shift:TShiftState; X,Y:Integer); override;
    procedure MouseUp(Button: TMouseButton; Shift:TShiftState; X,Y:Integer); override;

  public
    constructor Create(AOwner: TComponent); override;
    destructor Destroy; override;
    procedure Paint; override;
    procedure Resize; override;
  published
    property Caption: string read FCaption write SetCaption;
    property ShowCaption: Boolean read FShowCaption write SetShowCaption;
    property BackColor: TColor read FBackColor write SetBackColor;
    property HoverColor: TColor read FHoverColor write SetHoverColor;
    property DownColor: TColor read FDownColor write SetDownColor;
    property BorderWidth: Integer read FBorderWidth write SetBorderWidth;
    property BorderColor: TColor read FBorderColor write SetBorderColor;
    property Picture: TPicture read FPicture write SetPicture;
    property DrawMode: TDrawImageMode read FDrawMode write SetDrawMode;

    property NormalFontColor: TColor read FNormalFontColor write SetNormalFontColor;
    property DownFontColor: TColor read FDownFontColor write SetDownFontColor;
    property HoverFontColor: TColor read FHoverFontColor write SetHoverFontColor;

    property Action;
    property Align;
    property Anchors;
    property BiDiMode;
    property Constraints;
    property Enabled;
    property Font;
    property ParentFont;
    property ParentShowHint;
    property ParentBiDiMode;
    property PopupMenu;
    property ShowHint;
    property Visible;
    property OnClick;
    property OnDblClick;
    property OnMouseDown;
    property OnMouseEnter;
    property OnMouseLeave;
    property OnMouseMove;
    property OnMouseUp;
  end;

implementation

{ TXButton }

constructor TXButton.Create(AOwner: TComponent);
begin
  inherited Create(AOwner);
  Width := 80;
  Height := 40;
  FBufferBmp := TBitmap.Create;
  FPicture := TPicture.Create;
  FPicture.OnChange := {$IFDEF FPC}@{$ENDIF}DoPictureChange;
  FShowCaption := True;
  FDrawMode := dimCenter;
  FBackColor := clBtnFace;
  FState := xbsNone;
end;

destructor TXButton.Destroy;
begin
  FPicture.Free;
  FBufferBmp.Free;;
  inherited Destroy;
end;

procedure TXButton.Paint;
begin
  inherited Paint;

  if csDesigning in ComponentState then
  begin
    Canvas.Draw(1, 1,FBufferBmp);
    Canvas.Pen.Style := psDot;
    Canvas.Pen.Color := clBlack;
    Canvas.Brush.Style := bsClear;
    Canvas.Rectangle(Rect(0, 0, Width -1, Height - 1)); ///???
  end else
    Canvas.Draw(0, 0,FBufferBmp);
end;

procedure TXButton.Resize;
begin
  inherited Resize;
  if Assigned(FBufferBmp) then
  begin
    FBufferBmp.SetSize(Width, Height);
    FBufferBmp.PixelFormat := pf24bit;
    StateChanged;
  end;
end;

procedure TXButton.StateChanged;
type
  TRGB = record
  {$if Defined(WINDOWS) or Defined(MSWINDOWS)}
    B, G, R: Byte;
  {$elseif Defined(LINUX) or Defined(DARWIN)}
    R, G, B: Byte;
  {$ENDIF}
  end;
  PRGB = ^TRGB;

var
  X, Y: Integer;
  Ptr: PByte;
  LGray: Byte;
  LC: PRGB;
begin
  DrawBmp;
  if not Enabled then
  begin
  {$IFDEF FPC}
    FBufferBmp.BeginUpdate;
    try
  {$ENDIF}
     // 灰度化图像
     for Y := 0 to FBufferBmp.Height - 1 do
     begin
       Ptr := PByte(FBufferBmp.ScanLine[Y]);
       for X := 0 to FBufferBmp.Width - 1 do
       begin
         LC := PRGB(Ptr + X * 3);
         LGray := Byte(Round(0.299 * LC^.R + 0.587 * LC^.G + 0.114 * LC^.B + 0.1));
         LC^.R := LGray;
         LC^.G := LGray;
         LC^.B := LGray;
       end;
     end;
  {$IFDEF FPC}
    finally
      FBufferBmp.EndUpdate;
    end;
  {$ENDIF}
  end;
  Invalidate;
end;


procedure TXButton.DoPictureChange(Sender: TObject);
begin
  StateChanged;
end;

procedure TXButton.DrawBmp;
var
  LR: TRect;
{$IFDEF FPC}
  LTextStyle: TTextStyle;
{$ENDIF}
  LBrushColor, LFontColor: TColor;
begin
  LR := ClientRect;//Rect(0, 0, Width, Height);
  with FBufferBmp do
  begin
    FreeImage;
    if (FBufferBmp.Width <> LR.Width) or (FBufferBmp.Height <> LR.Height) then
      FBufferBmp.SetSize(LR.Width, LR.Height);
    // pen清除
    Canvas.Pen.Style := psClear;

    // 画背景 鼠标背景
    Canvas.Brush.Style:= bsSolid;

    case FState of
      xbsDown:
       begin
         LBrushColor := FDownColor;
         LFontColor := FDownFontColor;
       end;
      xbsHot:
        begin
          LBrushColor := FHoverColor;
          LFontColor := FHoverFontColor;
        end;
    else
      LBrushColor := FBackColor;
      LFontColor := FNormalFontColor
    end;
    Canvas.Brush.Color:= LBrushColor;
    Canvas.Rectangle(LR);

    // 画图片背景
    if Assigned(FPicture) and Assigned(FPicture.Graphic) then
    begin
      case FDrawMode of
        dimNormal:
         Canvas.Draw(0, 0, FPicture.Graphic);
        dimCenter:
         Canvas.Draw((LR.Width - FPicture.Graphic.Width) div 2, (LR.Height - FPicture.Graphic.Height) div 2, FPicture.Graphic);
        dimStretch:
         Canvas.StretchDraw(LR, FPicture.Graphic);
      end;
    end;

    // 画边框
    if FBorderWidth > 0 then
    begin
      Canvas.Pen.Width := FBorderWidth;
      Canvas.Pen.Color := FBorderColor;
      Canvas.Pen.Style := psSolid;
      Canvas.Brush.Style := bsClear;
      Canvas.Rectangle(LR);
    end;

    // 画文字
    if FShowCaption and (FCaption <> '') then
    begin
      Canvas.Brush.Style := bsClear;
      Canvas.Font.Color:= LFontColor;
    {$IFDEF FPC}
      LTextStyle := Canvas.TextStyle;
      LTextStyle.Alignment := taCenter;
      LTextStyle.Layout := tlCenter;
      Canvas.TextRect(LR, 0, 0, FCaption, LTextStyle);
    {$ELSE}
      Canvas.TextRect(LR, FCaption, [tfCenter, tfVerticalCenter, tfSingleLine]);
    {$ENDIF}
    end;
  end;
end;

procedure TXButton.SetBackColor(AValue: TColor);
begin
  if FBackColor=AValue then Exit;
  FBackColor:=AValue;
  StateChanged;
end;

procedure TXButton.SetBorderColor(AValue: TColor);
begin
  if FBorderColor=AValue then Exit;
  FBorderColor:=AValue;
  StateChanged;
end;

procedure TXButton.SetBorderWidth(AValue: Integer);
begin
  if FBorderWidth=AValue then Exit;
  FBorderWidth:=AValue;
  StateChanged;
end;

procedure TXButton.SetCaption(AValue: string);
begin
  if FCaption=AValue then Exit;
  FCaption:=AValue;
  StateChanged;
end;

procedure TXButton.SetDownColor(AValue: TColor);
begin
  if FDownColor=AValue then Exit;
  FDownColor:=AValue;
  StateChanged;
end;

procedure TXButton.SetDownFontColor(AValue: TColor);
begin
  if FDownFontColor=AValue then Exit;
  FDownFontColor:=AValue;
  StateChanged;
end;

procedure TXButton.SetDrawMode(AValue: TDrawImageMode);
begin
  if FDrawMode=AValue then Exit;
  FDrawMode:=AValue;
  StateChanged;
end;

procedure TXButton.SetHoverColor(AValue: TColor);
begin
  if FHoverColor=AValue then Exit;
  FHoverColor:=AValue;
  StateChanged;
end;

procedure TXButton.SetHoverFontColor(AValue: TColor);
begin
  if FHoverFontColor=AValue then Exit;
  FHoverFontColor:=AValue;
  StateChanged;
end;

procedure TXButton.SetNormalFontColor(AValue: TColor);
begin
  if FNormalFontColor=AValue then Exit;
  FNormalFontColor:=AValue;
  StateChanged;
end;

procedure TXButton.SetPicture(AValue: TPicture);
begin
  FPicture.Assign(AValue);
end;

procedure TXButton.SetShowCaption(AValue: Boolean);
begin
  if FShowCaption=AValue then Exit;
  FShowCaption:=AValue;
  StateChanged;
end;

procedure TXButton.CMMouseEnter(var Message: TMessage);
begin
  inherited;
  if csDesigning in ComponentState then
    Exit;
  if not Enabled then
    Exit;
  FMouseInControl := True;
  FState := xbsHot;
  StateChanged;
end;

procedure TXButton.CMMouseLeave(var Message: TMessage);
begin
  inherited;
  if csDesigning in ComponentState then
    Exit;
  if not Enabled then
    Exit;
  FState := xbsNone;
  StateChanged;
  FMouseInControl := False;
end;

procedure TXButton.MouseDown(Button: TMouseButton; Shift: TShiftState; X,
  Y: Integer);
begin
  inherited MouseDown(Button, Shift, X, Y);
  if csDesigning in ComponentState then
    Exit;
  if Button = mbLeft then
  begin
    if not Enabled then
      Exit;
    FState := xbsDown;
    StateChanged;
  end;
end;

procedure TXButton.CMEnabledChanged(var Message: TMessage);
begin
  inherited;
  if Enabled then
    FState := xbsNone
  else FState := xbsDisabled;
  StateChanged;
end;

procedure TXButton.CMFontChanged(var Message: TMessage);
begin
  inherited;
  StateChanged;
end;

procedure TXButton.MouseUp(Button: TMouseButton; Shift: TShiftState; X,
  Y: Integer);
begin
  inherited MouseUp(Button, Shift, X, Y);
  if csDesigning in ComponentState then
    Exit;
  if Button = mbLeft then
  begin
    if not Enabled then
      Exit;
    if FMouseInControl then
      FState := xbsHot
    else FState := xbsNone;
    StateChanged;
  end;
end;




end.

