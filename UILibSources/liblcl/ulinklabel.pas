unit uLinkLabel;

{$mode objfpc}{$H+}

interface

uses
  SysUtils, Classes, LCLType, LCLProc, LResources, Controls,
  Forms, StdCtrls, lMessages, GraphType, Graphics, LCLIntf,
  LCLClasses, Menus;

type
  TSysLinkType = (sltURL, sltID);

  TSysLinkEvent = procedure(Sender: TObject; const Link: string; LinkType: TSysLinkType) of object;


  { TLinkLabel }

  TLinkLabel = class(TCustomLabel)
  private
    FOnLinkClick: TSysLinkEvent;
    FDrawCaption: string;
    FLink: string;
    procedure DoOnLinkClick(const Link: string; LinkType: TSysLinkType);
    procedure ParserLink;
  protected
    procedure TextChanged; override;
    procedure CMMouseEnter(var Message :TLMessage); message CM_MOUSEENTER;
    procedure CMMouseLeave(var Message :TLMessage); message CM_MOUSELEAVE;
    procedure Click; override;
    function  GetLabelText: string; override;
  public
    constructor Create(AOwner: TComponent); override;
    procedure Paint; override;
  published
    property Align;
    property Alignment;
    property Anchors;
    property AutoSize;
    property Caption;
    property Color nodefault;
    property Constraints;
    property DragCursor;
    property DragKind;
    property DragMode;
    property Enabled;
    property Font;
    property ParentColor;
    property ParentFont;
    property ParentShowHint;
    property PopupMenu;
    property ShowHint;
    property Visible;
    property OnClick;
    property OnContextPopup;
    property OnDblClick;
    property OnDragDrop;
    property OnDragOver;
    property OnEndDock;
    property OnEndDrag;
    property OnMouseDown;
    property OnMouseEnter;
    property OnMouseLeave;
    property OnMouseMove;
    property OnMouseUp;
    property OnStartDock;
    property OnStartDrag;
    property OnLinkClick: TSysLinkEvent read FOnLinkClick write FOnLinkClick;
  end;




implementation

{ TLinkLabel }

constructor TLinkLabel.Create(AOwner: TComponent);
begin
  inherited;
  Width := 53;
  Height := 17;
  Self.Font.Style := [];
  Self.Font.Color := clBlue;
  Self.Cursor := crDefault;
end;

procedure TLinkLabel.Paint;
begin
  inherited Paint;
end;

procedure TLinkLabel.DoOnLinkClick(const Link: string; LinkType: TSysLinkType);
begin
  if Assigned(FOnLinkClick) then
    FOnLinkClick(Self, Link, LinkType);
end;

procedure TLinkLabel.ParserLink;
var
  iStart, iEnd: Integer;
begin
  FDrawCaption := '';
  iStart := Pos('">', Caption);
  if iStart > 0 then
  begin
    iEnd := Pos('</a>', Caption);
    if iEnd > 0 then
      FDrawCaption := Copy(Caption, iStart+2, iEnd - iStart - 2);
  end;
  FLink := '';
  iStart := Pos('href="', Caption);
  if iStart > 0 then
  begin
    iEnd := Pos('">', Caption);
    if iEnd > 0 then
      FLink := Copy(Caption, iStart + 6, iEnd - iStart - 6);
  end;
end;

procedure TLinkLabel.TextChanged;
begin
  ParserLink;
  inherited TextChanged;
end;

procedure TLinkLabel.CMMouseEnter(var Message: TLMessage);
begin
  Self.Font.Style := [fsUnderLine];
  Self.Font.Color := clRed;
  Self.Cursor := crHandPoint;
  inherited;
end;

procedure TLinkLabel.CMMouseLeave(var Message: TLMessage);
begin
  Self.Font.Style := [];
  Self.Font.Color := clBlue;
  Self.Cursor := crDefault;
  inherited;
end;

procedure TLinkLabel.Click;
begin
  DoOnLinkClick(FLink, sltURL);
  inherited Click;
end;

function TLinkLabel.GetLabelText: string;
begin
  Result:= FDrawCaption;
end;


end.

