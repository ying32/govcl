//----------------------------------------
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------
//----------------------------------------
// 用来兼容Delphi的TGauge
//----------------------------------------
unit uGauge;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Graphics, ATGauge;

type

  // 重新定义，目的为了兼容ATGauge的
  TGaugeKind = (gkText, gkHorizontalBar, gkVerticalBar, gkPie, gkNeedle, gkHalfPie);

  { TGauge }

  TGauge = class(TATGauge)
  private
    procedure DoFontChanged(Sender: TObject);

    function GetBackColor: TColor;
    function GetForeColor: TColor;
    function GetKind: TGaugeKind;
    procedure SetBackColor(AValue: TColor);
    procedure SetForeColor(AValue: TColor);
    procedure SetKind(AValue: TGaugeKind);
  public
    constructor Create(AOwner: TComponent); override;
    destructor Destroy; override;
  published
    property Align;
    property Anchors;
    property BorderStyle;
    property BorderSpacing;
    property Color;
    property Constraints;
    property DoubleBuffered;
    property ParentColor;
    property ParentShowHint;
    property PopupMenu;
    property ShowHint;
    property ParentFont;

    property ForeColor: TColor read GetForeColor write SetForeColor;
    property BackColor: TColor read GetBackColor write SetBackColor;
    property Kind: TGaugeKind read GetKind write SetKind default gkHorizontalBar;

    property Progress;
    property MinValue;
    property MaxValue;
    property ShowText;
    property ShowTextInverted;
    property OnClick;
    property OnDblClick;
    property OnResize;
    property OnContextPopup;
    property OnMouseDown;
    property OnMouseUp;
    property OnMouseMove;
    property OnMouseEnter;
    property OnMouseLeave;
    property OnMouseWheel;
    property OnMouseWheelDown;
    property OnMouseWheelUp;
  end;

implementation

{ TGauge }

procedure TGauge.DoFontChanged(Sender: TObject);
begin
  Theme^.FontName:=Font.Name;
  Theme^.FontSize:=Font.Size;
  Theme^.FontStyles:=Font.Style;
  Theme^.ColorFont:=Font.Color;
  Invalidate;
end;

function TGauge.GetBackColor: TColor;
begin
  Result := Theme^.ColorBgPassive;
end;

function TGauge.GetForeColor: TColor;
begin
  Result := Theme^.ColorBgOver;
end;

function TGauge.GetKind: TGaugeKind;
begin
  Result := TGaugeKind(inherited Kind);
end;

procedure TGauge.SetBackColor(AValue: TColor);
begin
  if Theme^.ColorBgPassive=AValue then Exit;
  Theme^.ColorBgPassive := AValue;
  Invalidate;
end;

procedure TGauge.SetForeColor(AValue: TColor);
begin
  if Theme^.ColorBgOver=AValue then Exit;
  Theme^.ColorBgOver := AValue;
  Invalidate;
end;

procedure TGauge.SetKind(AValue: TGaugeKind);
begin
  inherited Kind := TATGaugeKind(AValue);
end;

constructor TGauge.Create(AOwner: TComponent);
begin
  inherited Create(AOwner);
  Font.OnChange:=@DoFontChanged;
  ForeColor := clBlack;
  BackColor := clWhite;
  Color := clBtnFace;
end;

destructor TGauge.Destroy;
begin
  inherited Destroy;
end;

end.

