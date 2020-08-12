unit uMain;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, Graphics, Dialogs, Menus, ComCtrls,
  ExtCtrls, StdCtrls, ActnList, Buttons;

type

  { TMainForm }

  TMainForm = class(TForm)
    ActAudioEngineDelete: TAction;
    ActGenerateSubtitles: TAction;
    ActionList1: TActionList;
    ActTranslateEngineDelete: TAction;
    BtnDelAudioEngine: TButton;
    BtnDelTranslateEngine: TButton;
    Button1: TSpeedButton;
    CbbAudioEngines: TComboBox;
    CbbInputLang: TComboBox;
    CbbOutputAudioTrack: TComboBox;
    CbbOutputEncoding: TComboBox;
    CbbOutputLang: TComboBox;
    CbbTranslateEngines: TComboBox;
    ChkDoubleSubtitle: TCheckBox;
    ChkEnabledTranslate: TCheckBox;
    ChkMainSubtitle: TCheckBox;
    ChkOutputLRCFile: TCheckBox;
    ChkOutputSRTFile: TCheckBox;
    ChkOutputTxtFile: TCheckBox;
    ImageList1: TImageList;
    Label1: TLabel;
    Label2: TLabel;
    Label3: TLabel;
    Label4: TLabel;
    Label5: TLabel;
    Label6: TLabel;
    Label7: TLabel;
    Label8: TLabel;
    MiAppSettings: TMenuItem;
    MiGitee: TMenuItem;
    MiGithub: TMenuItem;
    MiHelpText: TMenuItem;
    MiNewAliyunAudoEngine: TMenuItem;
    MiNewBaiduTranslateEngine: TMenuItem;
    MiNewTencentTranslateEngine: TMenuItem;
    MiOpenMediaFile: TMenuItem;
    MiOSSSaveSettings: TMenuItem;
    MiQQGroup: TMenuItem;
    MiSponsor: TMenuItem;
    MmoInpuFiles: TMemo;
    MmoOutputLog: TMemo;
    OpenDialog1: TOpenDialog;
    Panel1: TPanel;
    Panel2: TPanel;
    Panel3: TPanel;
    Panel4: TPanel;
    Panel5: TPanel;
    Panel6: TPanel;
    Panel7: TPanel;
    PmOpen: TPopupMenu;
    PmNew: TPopupMenu;
    PmSettings: TPopupMenu;
    PmHelp: TPopupMenu;
    Splitter1: TSplitter;
    StatusBar1: TStatusBar;
    ToolBar1: TToolBar;
    TBtnOpen: TToolButton;
    TBtnNew: TToolButton;
    TBtnSettings: TToolButton;
    TBtnHelp: TToolButton;
    ToolButton5: TToolButton;
    ToolButton6: TToolButton;
    ToolButton7: TToolButton;
    procedure FormCreate(Sender: TObject);
    procedure FormDropFiles(Sender: TObject; const FileNames: array of String);
    procedure MiAppSettingsClick(Sender: TObject);
    procedure MiGiteeClick(Sender: TObject);
    procedure MiGithubClick(Sender: TObject);
    procedure MiHelpTextClick(Sender: TObject);
    procedure MiNewAliyunAudoEngineClick(Sender: TObject);
    procedure MiNewBaiduTranslateEngineClick(Sender: TObject);
    procedure MiNewTencentTranslateEngineClick(Sender: TObject);
    procedure MiOpenMediaFileClick(Sender: TObject);
    procedure MiOSSSaveSettingsClick(Sender: TObject);
    procedure MiQQGroupClick(Sender: TObject);
    procedure MiSponsorClick(Sender: TObject);
    procedure TBtnHelpClick(Sender: TObject);
    procedure TBtnNewClick(Sender: TObject);
    procedure TBtnOpenClick(Sender: TObject);
    procedure TBtnSettingsClick(Sender: TObject);
  private

  public

  end;

var
  MainForm: TMainForm;

implementation

{$R *.lfm}

{ TMainForm }

procedure TMainForm.FormDropFiles(Sender: TObject;
  const FileNames: array of String);
begin

end;

procedure TMainForm.MiAppSettingsClick(Sender: TObject);
begin

end;

procedure TMainForm.MiGiteeClick(Sender: TObject);
begin
   //
end;

procedure TMainForm.MiGithubClick(Sender: TObject);
begin
  //
end;

procedure TMainForm.MiHelpTextClick(Sender: TObject);
begin
  //
end;

procedure TMainForm.MiNewAliyunAudoEngineClick(Sender: TObject);
begin
 //
end;

procedure TMainForm.MiNewBaiduTranslateEngineClick(Sender: TObject);
begin

end;

procedure TMainForm.MiNewTencentTranslateEngineClick(Sender: TObject);
begin

end;

procedure TMainForm.MiOpenMediaFileClick(Sender: TObject);
begin

end;

procedure TMainForm.MiOSSSaveSettingsClick(Sender: TObject);
begin

end;

procedure TMainForm.MiQQGroupClick(Sender: TObject);
begin

end;

procedure TMainForm.MiSponsorClick(Sender: TObject);
begin

end;

procedure TMainForm.TBtnHelpClick(Sender: TObject);
begin

end;

procedure TMainForm.TBtnNewClick(Sender: TObject);
begin

end;

procedure TMainForm.TBtnOpenClick(Sender: TObject);
begin

end;

procedure TMainForm.TBtnSettingsClick(Sender: TObject);
begin

end;

procedure TMainForm.FormCreate(Sender: TObject);
begin
  TBtnSettings.Click;
end;




end.

