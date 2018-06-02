library testdll;

{ Important note about DLL memory management: ShareMem must be the
  first unit in your library's USES clause AND your project's (select
  Project-View Source) USES clause if your DLL exports any procedures or
  functions that pass strings as parameters or function results. This
  applies to all strings passed to and from your DLL--even those that
  are nested in records and classes. ShareMem is the interface unit to
  the BORLNDMM.DLL shared memory manager, which must be deployed along
  with your DLL. To avoid using BORLNDMM.DLL, pass string information
  using PChar or ShortString parameters. }

uses
  System.SysUtils,
  System.Classes;

{$R *.res}



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


function ResultValue: Int64; stdcall;
begin
  Result := Int64.MaxValue;
end;

function ResultValAndParam(p1, p2: Int64): UInt64; stdcall;
begin
  Writeln('p1:', p1, ', p2:', p2);
  if p1 = $FF2233445500 then
     Writeln('---p1 ok.')
  else
     Writeln('---p1 fail.');

  if p2 = $FF6677889900 then
     Writeln('---p2 ok.')
  else
     Writeln('---p2 fail.');

  Result := UInt64.MaxValue;
end;


function ResultValue32: Integer; stdcall;
begin
  Result := Integer.MaxValue;
end;

function ResultValAndParam32(p1, p2: Integer): Cardinal; stdcall;
begin
  Writeln('p1:', p1, ', p2:', p2);
  if p1 = $0F112233 then
     Writeln('---p1 ok.')
  else
     Writeln('---p1 fail.');

  if p2 = $0F445566 then
     Writeln('---p2 ok.')
  else
     Writeln('---p2 fail.');

  Result := Cardinal.MaxValue;
end;

exports
  ResultValue,
  ResultValAndParam,

  ResultValue32,
  ResultValAndParam32,

  MySyscall;

begin
end.

