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
  Self.TreeLinePenStyle := TPenStyle.psSolid;
{$IFDEF LCLCocoa}

{$ENDIF}
end;


{ TListView }

constructor TListView.Create(AOwner: TComponent);
begin
  inherited Create(AOwner);
  Self.ScrollBars := TScrollStyle.ssAutoBoth;
end;

end.

