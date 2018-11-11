{$MACRO ON}
{$IFDEF Windows}
  {$DEFINE extdecl := stdcall}
{$ELSE}
  {$DEFINE extdecl := cdecl}
{$ENDIF}

(*
   EventMethodName 实际的方法名称 TEventClass中的
   EventName 组件的方法名称
   EventMethod 方法
   EventMethodAddr 方法的地址
   EventMethodCode 代码段
*)
{$DEFINE EventMethod := TEventClass.EventMethodName}
{$DEFINE EventMethodAddr := @EventMethod}
{$DEFINE EventMethodCode := if AEventId = 0 then begin AObj.EventName := nil; TEventClass.Remove(AObj, EventMethodAddr); Exit; end; AObj.EventName := EventMethod; TEventClass.Add(AObj, EventMethodAddr, AEventId);}


(*
   事件中简化，暂时没用到。
*)
{$DEFINE SendEventParams := Sender, @TEventClass.EventMethodName}

