//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
// 
// Licensed under Apache License 2.0
//
//----------------------------------------

unit uControlPatchs;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, ExtCtrls, Forms, ComCtrls, StdCtrls, Graphics;


type
  PFNLVCOMPARE = Pointer;
  PFNTVCOMPARE = Pointer;



{$IFDEF WINDOWS}
 type
   TTrayIcon = class(ExtCtrls.TTrayIcon)
   public
     constructor Create(TheOwner: TComponent); override;
   end;

{$ENDIF}

type

  { TListView }

  TListView = class(ComCtrls.TListView)
  public
    constructor Create(AOwner: TComponent); override;
    procedure DeleteSelected;
    function CustomSort(ASortProc: PFNLVCOMPARE; AOptionalParam: PtrInt): Boolean;
  end;

  { TTreeView }

  TTreeView = class(ComCtrls.TTreeView)
  public
    constructor Create(AOwner: TComponent); override;
  end;

implementation

{$IFDEF WINDOWS}
constructor TTrayIcon.Create(TheOwner: TComponent);
begin
  inherited Create(TheOwner);
  if not Application.Icon.Empty then
    Icon.Assign(Application.Icon);
end;

{$ENDIF}

{ TTreeView }

constructor TTreeView.Create(AOwner: TComponent);
begin
  inherited Create(AOwner);
  Self.ScrollBars := TScrollStyle.ssAutoBoth;
  Self.ExpandSignType := tvestArrowFill;
{$IFNDEF WINDOWS}
  Self.TreeLinePenStyle := TPenStyle.psSolid;
{$ENDIF}
end;


{ TListView }

constructor TListView.Create(AOwner: TComponent);
begin
  inherited Create(AOwner);
  Self.ScrollBars := TScrollStyle.ssAutoBoth;
end;

procedure TListView.DeleteSelected;
begin
  if ItemIndex <> -1 then
    Items.Delete(ItemIndex);
end;

// 兼容版，因为内部Lazarus2.0起后有个兼容版本，但实际使用中并不兼容。
// TLVCompare = function(Item1, Item2: TListItem; AOptionalParam: PtrInt): Integer stdcall;
function DefaultListViewSortProc(Item1, Item2: TListItem; AOptionalParam: PtrInt): Integer stdcall;
begin
  with Item1 do
    if Assigned(TListView(ListView).OnCompare) then
      TListView(ListView).OnCompare(ListView, Item1, Item2, AOptionalParam, Result)
    else Result := CompareText(Item1.Caption, Item2.Caption);
end;



function TListView.CustomSort(ASortProc: PFNLVCOMPARE; AOptionalParam: PtrInt): Boolean;
begin
  //if not Assigned(ASortProc) then ASortProc := @DefaultListViewSortProc;
  Result := inherited CustomSort(@DefaultListViewSortProc, AOptionalParam);
end;

end.

