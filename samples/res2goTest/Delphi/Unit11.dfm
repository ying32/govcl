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
  end
  object Panel1: TPanel
    Left = 0
    Top = 29
    Width = 185
    Height = 235
    Align = alLeft
    BevelOuter = bvNone
    TabOrder = 2
    object Panel3: TPanel
      Left = 26
      Top = 47
      Width = 87
      Height = 41
      Caption = 'Panel3'
      Color = clRed
      ParentBackground = False
      TabOrder = 0
      OnMouseEnter = Panel3MouseEnter
      OnMouseLeave = Panel3MouseLeave
    end
    object Panel4: TPanel
      Left = 26
      Top = 94
      Width = 87
      Height = 41
      Caption = 'Panel4'
      Color = clGreen
      ParentBackground = False
      TabOrder = 1
      OnMouseEnter = Panel3MouseEnter
      OnMouseLeave = Panel3MouseLeave
    end
    object Panel5: TPanel
      Left = 26
      Top = 141
      Width = 87
      Height = 41
      Caption = 'Panel5'
      Color = clBlue
      ParentBackground = False
      TabOrder = 2
      OnMouseEnter = Panel3MouseEnter
      OnMouseLeave = Panel3MouseLeave
    end
  end
  object Panel2: TPanel
    Left = 188
    Top = 29
    Width = 298
    Height = 235
    Align = alClient
    BevelOuter = bvNone
    TabOrder = 3
    object Label1: TLabel
      Left = 104
      Top = 128
      Width = 65
      Height = 13
      Caption = 'Shared Event'
    end
    object Button1: TButton
      Left = 88
      Top = 56
      Width = 75
      Height = 25
      Caption = 'Show About'
      TabOrder = 0
      OnClick = Button1Click
    end
    object Button2: TButton
      Left = 48
      Top = 168
      Width = 75
      Height = 25
      Caption = 'Button2'
      TabOrder = 1
      OnClick = Button2Click
    end
    object Button3: TButton
      Left = 129
      Top = 168
      Width = 75
      Height = 25
      Caption = 'Button3'
      TabOrder = 2
      OnClick = Button2Click
    end
    object Button4: TButton
      Left = 210
      Top = 168
      Width = 75
      Height = 25
      Caption = 'Button4'
      TabOrder = 3
      OnClick = Button2Click
    end
  end
end
