unit Unit11;

interface

uses
  Windows, Messages, SysUtils, Variants, Classes, Graphics,
  Controls, Forms, Dialogs, ImageList, ImgList,
  StdCtrls, ComCtrls;

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

{$R *.lfm}

procedure TForm11.FormCreate(Sender: TObject);
begin
//
end;

end.
