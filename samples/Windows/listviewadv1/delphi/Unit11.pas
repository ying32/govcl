unit Unit11;

interface

uses
  Winapi.Windows, Winapi.Messages, System.SysUtils, System.Variants, System.Classes, Vcl.Graphics,
  Vcl.Controls, Vcl.Forms, Vcl.Dialogs, Vcl.StdCtrls, Vcl.ComCtrls, Winapi.CommCtrl,
  Vcl.ImgList, Vcl.Samples.Spin;

type
  TForm11 = class(TForm)
    ListView1: TListView;
    Button1: TButton;
    ImageList1: TImageList;
    ColorDialog1: TColorDialog;
    procedure FormCreate(Sender: TObject);
  private
    { Private declarations }
    procedure ColorLblClick(Sender: TObject);
  public
    { Public declarations }
  end;

var
  Form11: TForm11;

implementation

{$R *.dfm}

procedure TForm11.ColorLblClick(Sender: TObject);
begin
 if ColorDialog1.Execute then
    TLabel(Sender).Color := ColorDialog1.Color;
end;

procedure TForm11.FormCreate(Sender: TObject);
  procedure AddItem;
  const Colors: array[0..3] of TColor = (clRed, clGreen, clBlue, clLtGray);
  var
    LItem: TListItem;
    LCbb: TComboBox;
    LSpin: TSpinEdit;
    LColorLbl: TLabel;
    LR: TRect;
  begin
    LItem := ListView1.Items.Add;
    LItem.Caption := IntToStr(ListView1.Items.Count);
    LItem.SubItems.Add('');
    LItem.SubItems.Add('');
    LItem.SubItems.Add('');

    ListView_GetSubItemRect(ListView1.Handle, LItem.Index, 1, LVIR_BOUNDS, @LR);
    LCbb := TComboBox.Create(ListView1);
    LCbb.Parent := ListView1;
    LCbb.Items.Add('ÊÇ');
    LCbb.Items.Add('·ñ');
    LCbb.ItemIndex := 0;
    LCbb.Style := csDropDownList;
    LR.Inflate(-2, -2);
    LCbb.SetBounds(LR.Left, LR.Top, LR.Width, LR.Height);


    ListView_GetSubItemRect(ListView1.Handle, LItem.Index, 2, LVIR_BOUNDS, @LR);
    LR.Inflate(-2, -2);
    LColorLbl := TLabel.Create(ListView1);
    LColorLbl.Parent := ListView1;
    LColorLbl.AutoSize := False;
    LColorLbl.Color := Colors[Random(High(Colors))];
    LColorLbl.StyleElements := [];
    LColorLbl.Transparent := False;
    LColorLbl.SetBounds(LR.Left, LR.Top, LR.Width, LR.Height);
    LColorLbl.OnClick := ColorLblClick;

    ListView_GetSubItemRect(ListView1.Handle, LItem.Index, 3, LVIR_BOUNDS, @LR);
    LR.Inflate(-2, -2);
    LSpin := TSpinEdit.Create(ListView1);
    LSpin.Parent := ListView1;
    LSpin.Value := Random(10);
    LSpin.SetBounds(LR.Left, LR.Top, LR.Width, LR.Height);
  end;
var
  I: Integer;
begin
  Randomize;
  for I := 0 to 10 do
    AddItem;
end;

end.
