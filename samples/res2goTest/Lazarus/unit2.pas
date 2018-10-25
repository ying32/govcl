unit Unit2;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, FileUtil, Forms, Controls, Graphics, Dialogs;

type

  { TAbout }

  TAbout = class(TForm)
    procedure FormCreate(Sender: TObject);
  private

  public

  end;

var
  About: TAbout;

implementation

{$R *.lfm}

{ TAbout }

procedure TAbout.FormCreate(Sender: TObject);
begin
  //
end;

end.

