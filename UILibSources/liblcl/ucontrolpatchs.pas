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
  Classes, SysUtils, ExtCtrls, Forms;

{$IFDEF WINDOWS}
 type
   TTrayIcon = class(ExtCtrls.TTrayIcon)
   public
     constructor Create(TheOwner: TComponent); override;
   end;

{$ENDIF}

implementation

{$IFDEF WINDOWS}
constructor TTrayIcon.Create(TheOwner: TComponent);
begin
  inherited Create(TheOwner);
  if not Application.Icon.Empty then
    Icon.Assign(Application.Icon);
end;

{$ENDIF}

end.

