// TMsRdpClient9NotSafeForScriptingOnDisconnected
class procedure TEventClass.OnDisconnected(ASender: TObject; discReason: Integer);
begin
  SendEvent(ASender, @TEventClass.OnDisconnected, [ASender, discReason]);
end;

