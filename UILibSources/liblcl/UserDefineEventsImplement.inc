// 在这里引入你的你的事件实现部分的inc文件。
// 比如：
// {$I MsRdpClient9NotSafeForScriptingEventsDeclaration.inc}
// ...
// ...
// 那么MsRdpClient9NotSafeForScriptingEventsImplement.inc里面是些啥呢？
// 就是下面这类东西
//
//class procedure TEventClass.UpDownOnClick(Sender: TObject; Button: TUDBtnType);
//begin
//  SendEvent(Sender, @TEventClass.UpDownOnClick, [Sender, Ord(Button)]);
//end;
//
//