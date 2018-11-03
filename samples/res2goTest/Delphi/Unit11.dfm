object MainForm: TMainForm
  Left = 0
  Top = 0
  BorderStyle = bsDialog
  Caption = #20027#31383#21475
  ClientHeight = 286
  ClientWidth = 486
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
    Left = 185
    Top = 29
    Height = 235
    ExplicitLeft = 186
  end
  object ToolBar1: TToolBar
    Left = 0
    Top = 0
    Width = 486
    Height = 29
    Caption = 'ToolBar1'
    TabOrder = 0
  end
  object StatusBar1: TStatusBar
    Left = 0
    Top = 264
    Width = 486
    Height = 22
    Panels = <>
    ExplicitTop = 0
    ExplicitWidth = 8
  end
  object Panel1: TPanel
    Left = 0
    Top = 29
    Width = 185
    Height = 235
    Align = alLeft
    BevelOuter = bvNone
    TabOrder = 2
    ExplicitLeft = 152
    ExplicitTop = 136
    ExplicitHeight = 41
  end
  object Panel2: TPanel
    Left = 188
    Top = 29
    Width = 298
    Height = 235
    Align = alClient
    BevelOuter = bvNone
    TabOrder = 3
    ExplicitLeft = 152
    ExplicitTop = 136
    ExplicitWidth = 185
    ExplicitHeight = 41
    object Button1: TButton
      Left = 80
      Top = 96
      Width = 75
      Height = 25
      Caption = 'Show About'
      TabOrder = 0
      OnClick = Button1Click
    end
  end
end
