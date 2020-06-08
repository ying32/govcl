object Form11: TForm11
  Left = 619
  Top = 399
  Caption = 'Form11'
  ClientHeight = 321
  ClientWidth = 678
  Color = clBtnFace
  Font.Charset = DEFAULT_CHARSET
  Font.Color = clWindowText
  Font.Height = -11
  Font.Name = 'Tahoma'
  Font.Style = []
  OldCreateOrder = False
  Position = poDesigned
  OnCreate = FormCreate
  PixelsPerInch = 96
  TextHeight = 13
  object ListView1: TListView
    Left = 0
    Top = 0
    Width = 681
    Height = 289
    Columns = <
      item
        Caption = #21517#31216
        Width = 60
      end
      item
        Caption = #21487#35265
        Width = 80
      end
      item
        Caption = #39068#33394
        Width = 100
      end
      item
        Caption = #20540
        Width = 120
      end>
    GridLines = True
    ReadOnly = True
    RowSelect = True
    SmallImages = ImageList1
    TabOrder = 0
    ViewStyle = vsReport
  end
  object Button1: TButton
    Left = 0
    Top = 295
    Width = 75
    Height = 25
    Caption = 'Button1'
    TabOrder = 1
  end
  object ImageList1: TImageList
    Height = 24
    Width = 24
    Left = 336
    Top = 168
  end
  object ColorDialog1: TColorDialog
    Left = 344
    Top = 176
  end
end
