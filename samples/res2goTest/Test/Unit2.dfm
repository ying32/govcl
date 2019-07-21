object Frame2: TFrame2
  Left = 0
  Top = 0
  Width = 320
  Height = 240
  Align = alClient
  TabOrder = 0
  OnClick = FrameClick
  OnDblClick = FrameDblClick
  object ToolBar1: TToolBar
    Left = 0
    Top = 0
    Width = 320
    Height = 29
    Caption = 'ToolBar1'
    TabOrder = 0
    ExplicitLeft = 144
    ExplicitTop = 96
    ExplicitWidth = 150
    object ToolButton1: TToolButton
      Left = 0
      Top = 0
      Caption = 'ToolButton1'
      ImageIndex = 0
    end
    object ToolButton2: TToolButton
      Left = 23
      Top = 0
      Caption = 'ToolButton2'
      ImageIndex = 1
    end
    object ToolButton3: TToolButton
      Left = 46
      Top = 0
      Caption = 'ToolButton3'
      ImageIndex = 2
    end
    object ToolButton4: TToolButton
      Left = 69
      Top = 0
      Width = 8
      Caption = 'ToolButton4'
      ImageIndex = 3
      Style = tbsSeparator
    end
    object ToolButton5: TToolButton
      Left = 77
      Top = 0
      Caption = 'ToolButton5'
      ImageIndex = 3
    end
    object ToolButton6: TToolButton
      Left = 100
      Top = 0
      Caption = 'ToolButton6'
      ImageIndex = 4
    end
    object ToolButton7: TToolButton
      Left = 123
      Top = 0
      Width = 8
      Caption = 'ToolButton7'
      ImageIndex = 5
      Style = tbsSeparator
    end
    object ToolButton8: TToolButton
      Left = 131
      Top = 0
      Caption = 'ToolButton8'
      ImageIndex = 5
    end
  end
  object Button1: TButton
    Left = 100
    Top = 104
    Width = 75
    Height = 25
    Caption = 'Button1'
    TabOrder = 1
    OnClick = Button1Click
  end
  object StatusBar1: TStatusBar
    Left = 0
    Top = 221
    Width = 320
    Height = 19
    Panels = <>
    ExplicitLeft = 272
    ExplicitTop = 112
    ExplicitWidth = 0
  end
end
