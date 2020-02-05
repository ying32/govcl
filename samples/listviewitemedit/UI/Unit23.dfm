object Form1: TForm1
  Left = 0
  Top = 0
  Caption = 'Form1'
  ClientHeight = 605
  ClientWidth = 931
  Color = clBtnFace
  Font.Charset = DEFAULT_CHARSET
  Font.Color = clWindowText
  Font.Height = -11
  Font.Name = 'Tahoma'
  Font.Style = []
  OldCreateOrder = False
  Position = poScreenCenter
  OnCreate = FormCreate
  PixelsPerInch = 96
  TextHeight = 13
  object ListView1: TListView
    Left = 0
    Top = 0
    Width = 931
    Height = 605
    Align = alClient
    Columns = <
      item
        Caption = #24207
      end
      item
        Caption = #21015':TEdit'
        Width = 200
      end
      item
        Caption = #21015'TCombobox'
        Width = 200
      end>
    GridLines = True
    ReadOnly = True
    RowSelect = True
    SmallImages = ImageList1
    TabOrder = 0
    ViewStyle = vsReport
    OnDblClick = ListView1DblClick
  end
  object ComboBox1: TComboBox
    Left = 400
    Top = 240
    Width = 145
    Height = 21
    TabOrder = 1
    Text = 'ComboBox1'
    Visible = False
    OnExit = ComboBox1Exit
  end
  object Edit1: TEdit
    Left = 424
    Top = 360
    Width = 121
    Height = 21
    TabOrder = 2
    Text = 'Edit1'
    Visible = False
    OnExit = Edit1Exit
  end
  object ImageList1: TImageList
    Height = 21
    Width = 1
    Left = 432
    Top = 448
  end
end
