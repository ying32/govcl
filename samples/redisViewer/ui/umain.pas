unit uMain;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, Graphics, Dialogs, Menus, ExtCtrls,
  StdCtrls, Buttons, ComCtrls;

type

  { TMainForm }

  TMainForm = class(TForm)
    MainMenu1: TMainMenu;
    MIClose: TMenuItem;
    PgMain: TPageControl;
    Panel1: TPanel;
    Panel2: TPanel;
    Panel3: TPanel;
    BtnNewConn: TSpeedButton;
    PopupMenu1: TPopupMenu;
    Splitter1: TSplitter;
    StatusBar1: TStatusBar;
    TvConns: TTreeView;
    procedure BtnNewConnClick(Sender: TObject);
    procedure FormCreate(Sender: TObject);
    procedure FormDestroy(Sender: TObject);
    procedure MICloseClick(Sender: TObject);
    procedure TvConnsDblClick(Sender: TObject);
  private

  public

  end;

var
  MainForm: TMainForm;

implementation

{$R *.lfm}

{ TMainForm }

procedure TMainForm.FormCreate(Sender: TObject);
begin
  //
end;

procedure TMainForm.FormDestroy(Sender: TObject);
begin
  //
end;

procedure TMainForm.MICloseClick(Sender: TObject);
begin
  //
end;

procedure TMainForm.TvConnsDblClick(Sender: TObject);
begin
  //
end;

procedure TMainForm.BtnNewConnClick(Sender: TObject);
begin
  //
end;

end.

