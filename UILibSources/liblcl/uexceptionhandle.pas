unit uExceptionHandle;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Dialogs;

  procedure CallException(E: Exception);
implementation

// 显示异常
procedure CallException(E: Exception);
begin
  if Assigned(Application) then
  begin
    if Assigned(Application.OnException) then
      Application.OnException(Application, E)
    else
      MessageDlg(E.Message,  mtError, [mbOK], 0);
  end;
end;

end.

