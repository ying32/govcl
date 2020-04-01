
// 移植自Delphi组件,在lcl中算第三方组件了吧
// Ported from Delphi components

{$mode objfpc}{$H+}

unit Gauges;

interface

uses SysUtils, Classes, Graphics, Controls, Types;

type

  TGaugeKind = (gkText, gkHorizontalBar, gkVerticalBar, gkPie, gkNeedle);

  { TGauge }

  TGauge = class(TGraphicControl)
  private
    FMinValue: LongInt;
    FMaxValue: LongInt;
    FCurValue: LongInt;
    FKind: TGaugeKind;
    FShowText: boolean;
    FBorderStyle: TBorderStyle;
    FForeColor: TColor;
    FBackColor: TColor;
    procedure PaintBackground(AnImage: TBitmap);
    procedure PaintAsText(AnImage: TBitmap; APaintRect: TRect);
    procedure PaintAsNothing(AnImage: TBitmap; APaintRect: TRect);
    procedure PaintAsBar(AnImage: TBitmap; APaintRect: TRect);
    procedure PaintAsPie(AnImage: TBitmap; APaintRect: TRect);
    procedure PaintAsNeedle(AnImage: TBitmap; APaintRect: TRect);
    procedure SetGaugeKind(Value: TGaugeKind);
    procedure SetShowText(Value: Boolean);
    procedure SetBorderStyle(Value: TBorderStyle);
    procedure SetForeColor(Value: TColor);
    procedure SetBackColor(Value: TColor);
    procedure SetMinValue(Value: LongInt);
    procedure SetMaxValue(Value: LongInt);
    procedure SetProgress(Value: LongInt);
    function GetPercentDone: LongInt;

    procedure UpdateState;
  protected
    procedure Paint; override;
  public
    constructor Create(AOwner: TComponent); override;
    procedure AddProgress(Value: longint);
    property PercentDone: longint read GetPercentDone;
  published
    property Align;
    property Anchors;
    property BackColor: TColor read FBackColor write SetBackColor default clWhite;
    property BorderStyle: TBorderStyle
      read FBorderStyle write SetBorderStyle default bsSingle;
    property Color;
    property Constraints;
    property Enabled;
    property ForeColor: TColor read FForeColor write SetForeColor default clBlack;
    property Font;
    property Kind: TGaugeKind read FKind write SetGaugeKind default gkHorizontalBar;
    property MinValue: longint read FMinValue write SetMinValue default 0;
    property MaxValue: longint read FMaxValue write SetMaxValue default 100;
    property ParentColor;
    property ParentFont;
    property ParentShowHint;
    property PopupMenu;
    property Progress: longint read FCurValue write SetProgress;
    property ShowHint;
    property ShowText: boolean read FShowText write SetShowText default True;
    property Visible;
  end;

implementation

const
  SOutOfRange = 'Value must be between %d and %d';

type
  TBitmapHelper = class helper for TBitmap
  public
    procedure MakeLike(ATemplate: TBitmap);
  end;

{ TBitmapHelper }

procedure TBitmapHelper.MakeLike(ATemplate: TBitmap);
begin
  Width := ATemplate.Width;
  Height := ATemplate.Height;
  Canvas.Brush.Color := clWindowFrame;
  Canvas.Brush.Style := bsSolid;
  Canvas.FillRect(Rect(0, 0, Width, Height));
end;

function SolveForX(Y, Z: LongInt): LongInt;
begin
  Result := LongInt(Trunc(Z * (Y * 0.01)));
end;

function SolveForY(X, Z: LongInt): LongInt;
begin
  if Z = 0 then
    Result := 0
  else
    Result := LongInt(Trunc((X * 100.0) / Z));
end;

{ TGauge }

constructor TGauge.Create(AOwner: TComponent);
begin
  inherited Create(AOwner);
  ControlStyle := ControlStyle + [csFramed, csOpaque];
  FMinValue := 0;
  FMaxValue := 100;
  FCurValue := 0;
  FKind := gkHorizontalBar;
  FShowText := True;
  FBorderStyle := bsSingle;
  FForeColor := clBlack;
  FBackColor := clWhite;
  Width := 100;
  Height := 100;
end;

function TGauge.GetPercentDone: LongInt;
begin
  Result := SolveForY(FCurValue - FMinValue, FMaxValue - FMinValue);
end;

procedure TGauge.UpdateState;
begin
  // 不能用Repaint，linux下会卡
  Self.Invalidate;
end;

procedure TGauge.Paint;
var
  TheImage: TBitmap;
  OverlayImage: TBitmap;
  PaintRect: TRect;
begin
  with Canvas do
  begin
    TheImage := TBitmap.Create;
    try
      TheImage.Height := Height;
      TheImage.Width := Width;
      PaintBackground(TheImage);
      PaintRect := ClientRect;
      if FBorderStyle = bsSingle then
        InflateRect(PaintRect, -1, -1);
      OverlayImage := TBitmap.Create;
      try
        OverlayImage.MakeLike(TheImage);
        PaintBackground(OverlayImage);
        case FKind of
          gkText: PaintAsNothing(OverlayImage, PaintRect);
          gkHorizontalBar, gkVerticalBar: PaintAsBar(OverlayImage, PaintRect);
          gkPie: PaintAsPie(OverlayImage, PaintRect);
          gkNeedle: PaintAsNeedle(OverlayImage, PaintRect);
        end;
        TheImage.Canvas.CopyMode := cmSrcInvert;
        TheImage.Canvas.Draw(0, 0, OverlayImage);
        TheImage.Canvas.CopyMode := cmSrcCopy;
        if ShowText then
          PaintAsText(TheImage, PaintRect);
      finally
        OverlayImage.Free;
      end;
      Canvas.CopyMode := cmSrcCopy;
      Canvas.Draw(0, 0, TheImage);
    finally
      TheImage.Destroy;
    end;
  end;
end;

procedure TGauge.PaintBackground(AnImage: TBitmap);
var
  LRect: TRect;
begin
  with AnImage.Canvas do
  begin
    CopyMode := cmBlackness;
    LRect := Rect(0, 0, Width, Height);
    CopyRect(LRect, AnImage.Canvas, LRect);
    CopyMode := cmSrcCopy;
  end;
end;

procedure TGauge.PaintAsText(AnImage: TBitmap; APaintRect: TRect);
var
  LS: string;
  LX, LY: integer;
  LOverBmp: TBitmap;
begin
  LOverBmp := TBitmap.Create;
  try
    LOverBmp.MakeLike(AnImage);
    PaintBackground(LOverBmp);
    LS := Format('%d%%', [PercentDone]);
    with LOverBmp.Canvas do
    begin
      Brush.Style := bsClear;
      Font := Self.Font;
      Font.Color := clWhite;
      LX := (APaintRect.Right - APaintRect.Left + 1 - TextWidth(LS)) div 2;
      LY := (APaintRect.Bottom - APaintRect.Top + 1 - TextHeight(LS)) div 2;
      TextRect(APaintRect, LX, LY, LS);
    end;
    AnImage.Canvas.CopyMode := cmSrcInvert;
    AnImage.Canvas.Draw(0, 0, LOverBmp);
  finally
    LOverBmp.Free;
  end;
end;

procedure TGauge.PaintAsNothing(AnImage: TBitmap; APaintRect: TRect);
begin
  with AnImage do
  begin
    Canvas.Brush.Color := BackColor;
    Canvas.FillRect(APaintRect);
  end;
end;

procedure TGauge.PaintAsBar(AnImage: TBitmap; APaintRect: TRect);
var
  FillSize: LongInt;
  W, H: integer;
begin
  W := APaintRect.Right - APaintRect.Left + 1;
  H := APaintRect.Bottom - APaintRect.Top + 1;
  with AnImage.Canvas do
  begin
    Brush.Color := BackColor;
    FillRect(APaintRect);
    Pen.Color := ForeColor;
    Pen.Width := 1;
    Brush.Color := ForeColor;
    case FKind of
      gkHorizontalBar:
      begin
        FillSize := SolveForX(PercentDone, W);
        if FillSize > W then
          FillSize := W;
        if FillSize > 0 then
          FillRect(Rect(APaintRect.Left, APaintRect.Top, FillSize, H));
      end;
      gkVerticalBar:
      begin
        FillSize := SolveForX(PercentDone, H);
        if FillSize >= H then
          FillSize := H - 1;
        FillRect(Rect(APaintRect.Left, H - FillSize, W, H));
      end;
    end;
  end;
end;

procedure TGauge.PaintAsPie(AnImage: TBitmap; APaintRect: TRect);
var
  MiddleX, MiddleY: integer;
  Angle: double;
  W, H: integer;
begin
  W := APaintRect.Right - APaintRect.Left;
  H := APaintRect.Bottom - APaintRect.Top;
  if FBorderStyle = bsSingle then
  begin
    Inc(W);
    Inc(H);
  end;
  with AnImage.Canvas do
  begin
    Brush.Color := Color;
    FillRect(APaintRect);
    Brush.Color := BackColor;
    Pen.Color := ForeColor;
    Pen.Width := 1;
    Ellipse(APaintRect.Left, APaintRect.Top, W, H);
    if PercentDone > 0 then
    begin
      Brush.Color := ForeColor;
      MiddleX := W div 2;
      MiddleY := H div 2;
      Angle := (Pi * ((PercentDone / 50) + 0.5));
      Pie(APaintRect.Left, APaintRect.Top, W, H,
        integer(Round(MiddleX * (1 - Cos(Angle)))),
        integer(Round(MiddleY * (1 - Sin(Angle)))), MiddleX, 0);
    end;
  end;
end;

procedure TGauge.PaintAsNeedle(AnImage: TBitmap; APaintRect: TRect);
var
  MiddleX: integer;
  Angle: double;
  X, Y, W, H: integer;
begin
  with APaintRect do
  begin
    X := Left;
    Y := Top;
    W := Right - Left;
    H := Bottom - Top;
    if FBorderStyle = bsSingle then
    begin
      Inc(W);
      Inc(H);
    end;
  end;
  with AnImage.Canvas do
  begin
    Brush.Color := Color;
    FillRect(APaintRect);
    Brush.Color := BackColor;
    Pen.Color := ForeColor;
    Pen.Width := 1;
    Pie(X, Y, W, H * 2 - 1, X + W, APaintRect.Bottom - 1, X, APaintRect.Bottom - 1);
    MoveTo(X, APaintRect.Bottom);
    LineTo(X + W, APaintRect.Bottom);
    if PercentDone > 0 then
    begin
      Pen.Color := ForeColor;
      MiddleX := Width div 2;
      MoveTo(MiddleX, APaintRect.Bottom - 1);
      Angle := (Pi * ((PercentDone / 100)));
      LineTo(integer(Round(MiddleX * (1 - Cos(Angle)))),
        integer(Round((APaintRect.Bottom - 1) * (1 - Sin(Angle)))));
    end;
  end;
end;

procedure TGauge.SetGaugeKind(Value: TGaugeKind);
begin
  if Value <> FKind then
  begin
    FKind := Value;
    UpdateState;
  end;
end;

procedure TGauge.SetShowText(Value: Boolean);
begin
  if Value <> FShowText then
  begin
    FShowText := Value;
    UpdateState;
  end;
end;

procedure TGauge.SetBorderStyle(Value: TBorderStyle);
begin
  if Value <> FBorderStyle then
  begin
    FBorderStyle := Value;
    UpdateState;
  end;
end;

procedure TGauge.SetForeColor(Value: TColor);
begin
  if Value <> FForeColor then
  begin
    FForeColor := Value;
    UpdateState;
  end;
end;

procedure TGauge.SetBackColor(Value: TColor);
begin
  if Value <> FBackColor then
  begin
    FBackColor := Value;
    UpdateState;
  end;
end;

procedure TGauge.SetMinValue(Value: LongInt);
begin
  if Value <> FMinValue then
  begin
    if Value > FMaxValue then
      if not (csLoading in ComponentState) then
        raise EInvalidOperation.CreateFmt(SOutOfRange, [-MaxInt, FMaxValue - 1]);
    FMinValue := Value;
    if FCurValue < Value then
      FCurValue := Value;
    UpdateState;
  end;
end;

procedure TGauge.SetMaxValue(Value: LongInt);
begin
  if Value <> FMaxValue then
  begin
    if Value < FMinValue then
      if not (csLoading in ComponentState) then
        raise EInvalidOperation.CreateFmt(SOutOfRange, [FMinValue + 1, MaxInt]);
    FMaxValue := Value;
    if FCurValue > Value then
      FCurValue := Value;
    UpdateState;
  end;
end;

procedure TGauge.SetProgress(Value: LongInt);
var
  TempPercent: LongInt;
begin
  TempPercent := GetPercentDone;
  if Value < FMinValue then
    Value := FMinValue
  else if Value > FMaxValue then
    Value := FMaxValue;
  if FCurValue <> Value then
  begin
    FCurValue := Value;
    if TempPercent <> GetPercentDone then
      UpdateState;
  end;
end;

procedure TGauge.AddProgress(Value: LongInt);
begin
  Progress := FCurValue + Value;
  UpdateState;
end;

end.
