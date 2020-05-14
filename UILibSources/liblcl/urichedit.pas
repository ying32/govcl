//----------------------------------------
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------
//----------------------------------------
// 用来兼容Delphi的TRichEdit
//----------------------------------------
unit uRichEdit;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, RichMemo, richmemohelpers, Graphics;

type

  TSearchType = TSearchOption;
  TSearchTypes = TSearchOptions;

  TAttributeType = (atSelected, atDefaultText);

  TNumberingStyle = (nsNone, nsBullet);

  TRichEdit = class;

  { TTextAttributes }

  TTextAttributes = class(TPersistent)
  private
    FOwner: TRichMemo;
    FType: TAttributeType;
    function DoGetAttrs: TFontParams;
  private
    function GetCharset: TFontCharset;
    function GetColor: TColor;
    function GetHeight: Integer;
    function GetName: string;
    function GetPitch: TFontPitch;
    function GetSize: Integer;
    function GetStyle: TFontStyles;
    procedure SetCharset(AValue: TFontCharset);
    procedure SetColor(AValue: TColor);
    procedure SetHeight(AValue: Integer);
    procedure SetName(AValue: string);
    procedure SetPitch(AValue: TFontPitch);
    procedure SetSize(AValue: Integer);
    procedure SetStyle(AValue: TFontStyles);
  public
    constructor Create(AOwner: TRichMemo; AType: TAttributeType);
    procedure Assign(Source: TPersistent); override;

    property Name: string read GetName write SetName;
    property Pitch: TFontPitch read GetPitch write SetPitch;
    property Charset: TFontCharset read GetCharset write SetCharset;
    property Color: TColor read GetColor write SetColor;
    property Size: Integer read GetSize write SetSize;
    property Style: TFontStyles read GetStyle write SetStyle;
    property Height: Integer read GetHeight write SetHeight;
  end;


  { TParaAttributes }

  TParaAttributes = class(TPersistent)
  private
    FOwner: TRichMemo;
    function DoGetParaMetric: TParaMetric;
  private
    function GetAlignment: TAlignment;
    function GetFirstIndent: Integer;
    function GetLeftIndent: Integer;
    function GetNumbering: TNumberingStyle;
    function GetRightIndent: Integer;
    function GetTab(Index: Byte): Longint;
    function GetTabCount: Integer;
    procedure SetAlignment(AValue: TAlignment);
    procedure SetFirstIndent(AValue: Integer);
    procedure SetLeftIndent(AValue: Integer);
    procedure SetNumbering(AValue: TNumberingStyle);
    procedure SetRightIndent(AValue: Integer);
    procedure SetTab(Index: Byte; AValue: Longint);
    procedure SetTabCount(AValue: Integer);
  public
    constructor Create(AOwner: TRichMemo);
    procedure Assign(Source: TPersistent); override;

    property Alignment: TAlignment read GetAlignment write SetAlignment;
    property FirstIndent: Integer read GetFirstIndent write SetFirstIndent;
    property LeftIndent: Integer read GetLeftIndent write SetLeftIndent;
    property RightIndent: Integer read GetRightIndent write SetRightIndent;
    property Numbering: TNumberingStyle read GetNumbering write SetNumbering;
    property TabCount: Integer read GetTabCount write SetTabCount;
    property Tab[Index: Byte]: Longint read GetTab write SetTab;
  end;

  { TRichEdit }

  TRichEdit = class(TRichMemo)
  private
    FDefAttributes: TTextAttributes;
    FHideScrollBars: Boolean;
    FParagraph: TParaAttributes;
    FPlainText: Boolean;
    FSelAttributes: TTextAttributes;
    function GetZoom: Integer;
    procedure SetDefAttributes(AValue: TTextAttributes);
    procedure SetSelAttributes(AValue: TTextAttributes);
    procedure SetZoom(AValue: Integer);
  public
    constructor Create(AOnwer: TComponent); override;
    destructor Destroy; override;
    function FindText(ASearchStr: string; AStartPos: Integer; ALength: Integer; AOptions: TSearchTypes): Integer;

    property Zoom: Integer read GetZoom write SetZoom;
    property HideScrollBars: Boolean read FHideScrollBars write FHideScrollBars;
    property PlainText: Boolean read FPlainText write FPlainText;
    property DefAttributes: TTextAttributes read FDefAttributes write SetDefAttributes;
    property SelAttributes: TTextAttributes read FSelAttributes write SetSelAttributes;
    property Paragraph: TParaAttributes read FParagraph write FParagraph;
  public
    property Align;
    property Alignment;
    property Anchors;
    property BidiMode;
    property BorderSpacing;
    property BorderStyle;
    property Color;
    property Constraints;
    property DragCursor;
    property DragKind;
    property DragMode;
    property Enabled;
    property Font;
    property HideSelection;
    property Lines;
    property MaxLength;
    property OnChange;
    property OnClick;
    property OnContextPopup;
    property OnDblClick;
    property OnDragDrop;
    property OnDragOver;
    property OnEditingDone;
    property OnEndDrag;
    property OnEnter;
    property OnExit;
    property OnKeyDown;
    property OnKeyPress;
    property OnKeyUp;
    property OnLinkAction;
    property OnMouseDown;
    property OnMouseEnter;
    property OnMouseLeave;
    property OnMouseMove;
    property OnMouseUp;
    property OnMouseWheel;
    property OnMouseWheelDown;
    property OnMouseWheelUp;
    property OnSelectionChange;
    property OnStartDrag;
    property OnPrintAction;
    property OnUTF8KeyPress;
    property ParentBidiMode;
    property ParentColor;
    property ParentFont;
    property PopupMenu;
    property ParentShowHint;
    property ReadOnly;
    //property Rtf: string read GetRTF write SetRTF;
    property ScrollBars;
    property ShowHint;
    property TabOrder;
    property TabStop;
    property Visible;
    property WantReturns;
    property WantTabs;
    property WordWrap;
    //property ZoomFactor;
  end;

implementation

procedure SafeFillChar(out x; ASize: SizeInt);
begin
  FillChar(x, ASize, #0);
end;


{ TParaAttributes }

function TParaAttributes.DoGetParaMetric: TParaMetric;
begin
  FillChar(Result, SizeOf(Result), #0);
  FOwner.GetParaMetric(FOwner.SelStart, Result);
end;

function TParaAttributes.GetAlignment: TAlignment;
begin
  Result := TAlignment(RichMemoHelpers.TParaAttributes(TObject(FOwner)).Alignment);
 //Result := TAlignment(FOwner.GetParaAlignment(FOwner.SelStart));
end;

function TParaAttributes.GetFirstIndent: Integer;
begin
  Result := RichMemoHelpers.TParaAttributes(TObject(FOwner)).FirstIndent;
  //Result := Trunc(DoGetParaMetric.FirstLine);
end;

function TParaAttributes.GetLeftIndent: Integer;
begin
  Result := RichMemoHelpers.TParaAttributes(TObject(FOwner)).LeftIndent;
end;

function TParaAttributes.GetNumbering: TNumberingStyle;
begin
  Result := nsNone;
end;

function TParaAttributes.GetRightIndent: Integer;
begin
  Result := RichMemoHelpers.TParaAttributes(TObject(FOwner)).RightIndent;
end;

function TParaAttributes.GetTab(Index: Byte): Longint;
begin
  Result := RichMemoHelpers.TParaAttributes(TObject(FOwner)).Tab[Index];
end;

function TParaAttributes.GetTabCount: Integer;
begin
  Result := RichMemoHelpers.TParaAttributes(TObject(FOwner)).TabCount;
end;

procedure TParaAttributes.SetAlignment(AValue: TAlignment);
begin
  RichMemoHelpers.TParaAttributes(TObject(FOwner)).Alignment := RichMemoHelpers.TRichEditAlignment(AValue);
  // TAlignment = (taLeftJustify, taRightJustify, taCenter);
end;

procedure TParaAttributes.SetFirstIndent(AValue: Integer);
begin
  RichMemoHelpers.TParaAttributes(TObject(FOwner)).FirstIndent:=AValue;
end;

procedure TParaAttributes.SetLeftIndent(AValue: Integer);
begin
  RichMemoHelpers.TParaAttributes(TObject(FOwner)).LeftIndent:=AValue;
end;

procedure TParaAttributes.SetNumbering(AValue: TNumberingStyle);
begin

end;

procedure TParaAttributes.SetRightIndent(AValue: Integer);
begin
  RichMemoHelpers.TParaAttributes(TObject(FOwner)).RightIndent:=AValue;
end;

procedure TParaAttributes.SetTab(Index: Byte; AValue: Longint);
begin
  RichMemoHelpers.TParaAttributes(TObject(FOwner)).Tab[Index] := AValue;
end;

procedure TParaAttributes.SetTabCount(AValue: Integer);
begin
  RichMemoHelpers.TParaAttributes(TObject(FOwner)).TabCount:=AValue;
end;

constructor TParaAttributes.Create(AOwner: TRichMemo);
begin
  inherited Create;
  FOwner := AOwner;
end;

procedure TParaAttributes.Assign(Source: TPersistent);
begin
  inherited Assign(Source);
end;

{ TTextAttributes }

function TTextAttributes.DoGetAttrs: TFontParams;
begin
  SafeFillChar(Result, SizeOf(Result));
  FOwner.GetTextAttributes(FOwner.SelStart, Result);
end;

function TTextAttributes.GetCharset: TFontCharset;
begin
  Result := 0;
  //if FType = atDefaultText then
    //Result := FOwner.GetTextAttributes();
end;

function TTextAttributes.GetColor: TColor;
begin
  Result := DoGetAttrs.Color;
end;

function TTextAttributes.GetHeight: Integer;
begin
  Result := 0;
end;

function TTextAttributes.GetName: string;
begin
  Result := DoGetAttrs.Name;
end;

function TTextAttributes.GetPitch: TFontPitch;
begin
  Result := fpDefault;
end;

function TTextAttributes.GetSize: Integer;
begin
  Result := DoGetAttrs.Size;
end;

function TTextAttributes.GetStyle: TFontStyles;
begin
  Result := DoGetAttrs.Style;
end;

procedure TTextAttributes.SetCharset(AValue: TFontCharset);
begin
  if FType = atDefaultText then
    FOwner.SetTextAttributes(FOwner.SelStart, FOwner.SelLength, FOwner.Font)
end;

procedure TTextAttributes.SetColor(AValue: TColor);
begin
  if FType = atDefaultText then
    FOwner.SetTextAttributes(FOwner.SelStart, FOwner.SelLength, FOwner.Font)
  else
    FOwner.SetRangeColor(FOwner.SelStart, FOwner.SelLength, AValue);
end;

procedure TTextAttributes.SetHeight(AValue: Integer);
begin
  if FType = atDefaultText then
    FOwner.SetTextAttributes(FOwner.SelStart, FOwner.SelLength, FOwner.Font)
end;


procedure TTextAttributes.SetName(AValue: string);
begin
  if FType = atDefaultText then
    FOwner.SetTextAttributes(FOwner.SelStart, FOwner.SelLength, FOwner.Font)
  else
    FOwner.SetRangeParams(FOwner.SelStart, FOwner.SelLength, [tmm_Name], AValue, 0, 0, [], []);
end;

procedure TTextAttributes.SetPitch(AValue: TFontPitch);
begin
  if FType = atDefaultText then
    FOwner.SetTextAttributes(FOwner.SelStart, FOwner.SelLength, FOwner.Font)
end;

procedure TTextAttributes.SetSize(AValue: Integer);
begin
  if FType = atDefaultText then
    FOwner.SetTextAttributes(FOwner.SelStart, FOwner.SelLength, FOwner.Font)
  else
    FOwner.SetRangeParams(FOwner.SelStart, FOwner.SelLength, [tmm_Size], '', AValue, 0, [], []);
end;

procedure TTextAttributes.SetStyle(AValue: TFontStyles);
begin
  if FType = atDefaultText then
    FOwner.SetTextAttributes(FOwner.SelStart, FOwner.SelLength, FOwner.Font)
  else
    FOwner.SetRangeParams(FOwner.SelStart, FOwner.SelLength, [tmm_Styles], '', 0, 0, AValue, []);
end;

constructor TTextAttributes.Create(AOwner: TRichMemo; AType: TAttributeType);
begin
  inherited Create;
  FOwner := AOwner;
  FType := AType;
end;

procedure TTextAttributes.Assign(Source: TPersistent);
begin
  inherited Assign(Source);
end;

{ TRichEdit }

function TRichEdit.GetZoom: Integer;
begin
  Result := Trunc(ZoomFactor);
end;

procedure TRichEdit.SetDefAttributes(AValue: TTextAttributes);
begin
  if FDefAttributes=AValue then Exit;
  FDefAttributes:=AValue;
end;

procedure TRichEdit.SetSelAttributes(AValue: TTextAttributes);
begin
  if FSelAttributes=AValue then Exit;
  FSelAttributes:=AValue;
end;

procedure TRichEdit.SetZoom(AValue: Integer);
begin
  ZoomFactor := AValue;
end;

constructor TRichEdit.Create(AOnwer: TComponent);
begin
  inherited Create(AOnwer);
  FDefAttributes := TTextAttributes.Create(Self, atDefaultText);
  FParagraph := TParaAttributes.Create(Self);
  FSelAttributes := TTextAttributes.Create(Self, atSelected);
end;

destructor TRichEdit.Destroy;
begin
  FSelAttributes.Free;
  FParagraph.Free;
  FDefAttributes.Free;
  inherited Destroy;
end;

function TRichEdit.FindText(ASearchStr: string; AStartPos: Integer;
  ALength: Integer; AOptions: TSearchTypes): Integer;
begin
  Result := Search(ASearchStr, AStartPos, ALength, AOptions);
end;

end.

