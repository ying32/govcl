unit Unit11;

interface

uses
  Winapi.Windows, Winapi.Messages, System.SysUtils, System.Variants, System.Classes, Vcl.Graphics,
  Vcl.Controls, Vcl.Forms, Vcl.Dialogs, System.ImageList, Vcl.ImgList,
  Vcl.StdCtrls, Vcl.ComCtrls;

type
  TForm11 = class(TForm)
    ComboBoxEx1: TComboBoxEx;
    ImageList1: TImageList;
    ComboBoxEx2: TComboBoxEx;
    ImageList2: TImageList;
    procedure FormCreate(Sender: TObject);
  private
    { Private declarations }
  public
    { Public declarations }
  end;

var
  Form11: TForm11;

implementation

{$R *.dfm}

procedure TForm11.FormCreate(Sender: TObject);
begin
//
end;

end.
