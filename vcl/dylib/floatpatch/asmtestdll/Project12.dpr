library Project12;

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

{
      Epsilon:Single = 1.4012984643248170709e-45;
      MaxValue:Single =  340282346638528859811704183484516925440.0;
      MinValue:Single = -340282346638528859811704183484516925440.0;
      PositiveInfinity:Single =  1.0 / 0.0;
      NegativeInfinity:Single = -1.0 / 0.0;
      NaN:Single = 0.0 / 0.0;
}

function testFloat32_MaxValue: Single; cdecl;
begin
  Result := Single.MaxValue;
end;

function testFloat32_MinValue: Single; cdecl;
begin
  Result := Single.MinValue;
end;

function testFloat32_Epsilon: Single; cdecl;
begin
  Result := Single.Epsilon;
end;

//function testFloat32_Value(A1, A2: Single): Single; cdecl;
//begin
//  Result := Single.Epsilon;
//end;


{
      Epsilon:Double = 4.9406564584124654418e-324;
      MaxValue:Double =  1.7976931348623157081e+308;
      MinValue:Double = -1.7976931348623157081e+308;
      PositiveInfinity:Double =  1.0 / 0.0;
      NegativeInfinity:Double = -1.0 / 0.0;
      NaN:Double = 0.0 / 0.0;
}

function testFloat64_MaxValue: Double; cdecl;
begin
  Result := Double.MaxValue;
end;

function testFloat64_MinValue: Double; cdecl;
begin
  Result := Double.MinValue;
end;

function testFloat64_Epsilon: Double; cdecl;
begin
  Result := Double.Epsilon;
end;

//function testFloat64_Value(A1, A2: Double): Single; cdecl;
//begin
//  Result := Single.Epsilon;
//end;




exports
   testFloat32_MaxValue,
   testFloat32_MinValue,
   testFloat32_Epsilon,

   testFloat64_MaxValue,
   testFloat64_MinValue,
   testFloat64_Epsilon;

begin
end.
