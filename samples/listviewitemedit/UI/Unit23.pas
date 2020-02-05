unit Unit23;

interface

uses
  Winapi.Windows, Winapi.Messages, System.SysUtils, System.Variants, System.Classes, Vcl.Graphics,
  Vcl.Controls, Vcl.Forms, Vcl.Dialogs, Vcl.StdCtrls, Vcl.ComCtrls,
  System.ImageList, Vcl.ImgList, CommCtrl;

type
  TForm1 = class(TForm)
    ListView1: TListView;
    ComboBox1: TComboBox;
    Edit1: TEdit;
    ImageList1: TImageList;
    procedure FormCreate(Sender: TObject);
    procedure ComboBox1Exit(Sender: TObject);
    procedure Edit1Exit(Sender: TObject);
    procedure ListView1DblClick(Sender: TObject);
  private
    { Private declarations }
  public
    { Public declarations }
  end;

var
  Form1: TForm1;

implementation

{$R *.dfm}

procedure TForm1.ComboBox1Exit(Sender: TObject);
begin
  //
end;

procedure TForm1.Edit1Exit(Sender: TObject);
begin
 //
end;

procedure TForm1.FormCreate(Sender: TObject);
begin
//
end;

procedure TForm1.ListView1DblClick(Sender: TObject);
begin
//
end;

end.
