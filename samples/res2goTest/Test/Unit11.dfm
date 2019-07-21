object MainForm: TMainForm
  Left = 0
  Top = 0
  Caption = #20027#31383#21475
  ClientHeight = 465
  ClientWidth = 822
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
  object Splitter1: TSplitter
    Left = 0
    Top = 0
    Height = 465
    ExplicitLeft = 186
    ExplicitTop = 29
    ExplicitHeight = 235
  end
  object Panel1: TPanel
    Left = 3
    Top = 0
    Width = 177
    Height = 465
    Align = alLeft
    TabOrder = 0
    ExplicitLeft = 56
    ExplicitTop = 144
    ExplicitHeight = 241
    object BtnShowAbout: TButton
      Left = 40
      Top = 88
      Width = 75
      Height = 25
      Caption = 'Show About'
      TabOrder = 0
      OnClick = BtnShowAboutClick
    end
    object BtnShowFrame1: TButton
      Left = 40
      Top = 152
      Width = 75
      Height = 25
      Caption = 'Show Frame1'
      TabOrder = 1
      OnClick = BtnShowFrame1Click
    end
    object BtnShowFrame2: TButton
      Left = 40
      Top = 210
      Width = 75
      Height = 25
      Caption = 'Show Frame2'
      TabOrder = 2
      OnClick = BtnShowFrame2Click
    end
  end
  object Panel2: TPanel
    Left = 180
    Top = 0
    Width = 642
    Height = 465
    Align = alClient
    TabOrder = 1
    ExplicitLeft = 528
    ExplicitTop = 216
    ExplicitWidth = 185
    ExplicitHeight = 41
  end
end
