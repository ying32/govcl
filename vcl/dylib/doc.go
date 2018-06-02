/*
   dylib是从govcl的api包中分离出来，可以方便其它使用。
   dylib实现现了一个通用的共享库调用，可以在windows、linux、macOS下不需要任何改变实现通用。
   Window下使用syscall.NewLazyDLL加载dll，linux与macOS下使用dlopen加载，然后使用dlsym来获取,
   通过封装，使用接口与Windows下的一模一样，如此达到通用，最多实现12个参数的过程。

   在被调用的共享库中如果导出了“MySyscall”函数，则会使用此函数来call， 一般导出此函数的目的是为了在共享库中捕获异常，不然一但共享库中出现异常
   程序就挂了。

   一个Pascal的实现例程：

// dylib
type
  TSyscall0 = function: UInt64; stdcall;
  TSyscall1 = function(A1: Pointer): UInt64; stdcall;
  TSyscall2 = function(A1, A2: Pointer): UInt64; stdcall;
  TSyscall3 = function(A1, A2, A3: Pointer): UInt64; stdcall;
  TSyscall4 = function(A1, A2, A3, A4: Pointer): UInt64; stdcall;
  TSyscall5 = function(A1, A2, A3, A4, A5: Pointer): UInt64; stdcall;
  TSyscall6 = function(A1, A2, A3, A4, A5, A6: Pointer): UInt64; stdcall;
  TSyscall7 = function(A1, A2, A3, A4, A5, A6, A7: Pointer): UInt64; stdcall;
  TSyscall8 = function(A1, A2, A3, A4, A5, A6, A7, A8: Pointer): UInt64; stdcall;
  TSyscall9 = function(A1, A2, A3, A4, A5, A6, A7, A8, A9: Pointer): UInt64; stdcall;
  TSyscall10 = function(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10: Pointer): UInt64; stdcall;
  TSyscall11 = function(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11: Pointer): UInt64; stdcall;
  TSyscall12 = function(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12: Pointer): UInt64; stdcall;

function MySyscall(AProc: Pointer; ALen: NativeInt; A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12: Pointer): UInt64; stdcall;
begin
  Result := 0;
  if AProc = nil then
    Exit;
  try
    case ALen of
      0: Result := TSyscall0(AProc)();
      1: Result := TSyscall1(AProc)(A1);
      2: Result := TSyscall2(AProc)(A1, A2);
      3: Result := TSyscall3(AProc)(A1, A2, A3);
      4: Result := TSyscall4(AProc)(A1, A2, A3, A4);
      5: Result := TSyscall5(AProc)(A1, A2, A3, A4, A5);
      6: Result := TSyscall6(AProc)(A1, A2, A3, A4, A5, A6);
      7: Result := TSyscall7(AProc)(A1, A2, A3, A4, A5, A6, A7);
      8: Result := TSyscall8(AProc)(A1, A2, A3, A4, A5, A6, A7, A8);
      9: Result := TSyscall9(AProc)(A1, A2, A3, A4, A5, A6, A7, A8, A9);
      10: Result := TSyscall10(AProc)(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10);
      11: Result := TSyscall11(AProc)(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11);
      12: Result := TSyscall12(AProc)(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12);
    else
      Exit;
    end;
  except
    on E: Exception do
    begin
      Writeln('Syscall', ALen, ' Error: ', E.Message);
    end;
  end;
end;

*/

package dylib
