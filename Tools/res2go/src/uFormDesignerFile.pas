unit uFormDesignerFile;

interface

uses
  sysutils,
  Classes,
{$IFDEF FPC}
  zbase,
  typinfo,
  paszlib,
  Types
{$ELSE}
  Zlib
{$ENDIF}
  ;


type
{$IFDEF FPC}
  EZLibError = class(Exception);
  EZCompressionError = class(EZLibError);
  EZDecompressionError = class(EZLibError);
{$ENDIF}


  { TFormResFile }

  TFormResFile = class
  public const
    VERSION = 1;
    // GOVCLFORMZ
    HEADER: array[0..9] of byte = ($47, $4F, $56, $43, $4C, $46, $4F, $52, $4D, $5A);
  public
    class function Encrypt(AInStream: TStream; AOutStream: TStream): Boolean;
    class procedure XorStream(AStream: TMemoryStream);
  end;

implementation


{$IFDEF FPC}

type
  TZCompressionLevel = (zcNone, zcFastest, zcDefault, zcMax);

  // CG: Define old enum for compression level
  TCompressionLevel = (clNone = Integer(zcNone), clFastest, clDefault, clMax);

  TZStreamRec = z_stream;

const
  {** return code messages **************************************************}
  _z_errmsg: array [0..9] of PAnsiChar = (
    'need dictionary',      // Z_NEED_DICT      (2)  //do not localize
    'stream end',           // Z_STREAM_END     (1)  //do not localize
    '',                     // Z_OK             (0)  //do not localize
    'file error',           // Z_ERRNO          (-1) //do not localize
    'stream error',         // Z_STREAM_ERROR   (-2) //do not localize
    'data error',           // Z_DATA_ERROR     (-3) //do not localize
    'insufficient memory',  // Z_MEM_ERROR      (-4) //do not localize
    'buffer error',         // Z_BUF_ERROR      (-5) //do not localize
    'incompatible version', // Z_VERSION_ERROR  (-6) //do not localize
    ''                                               //do not localize
    );

  ZLevels: array[TZCompressionLevel] of Shortint = (
    Z_NO_COMPRESSION,
    Z_BEST_SPEED,
    Z_DEFAULT_COMPRESSION,
    Z_BEST_COMPRESSION
    );


function ZDecompressCheck(code: Integer): Integer;
begin
  Result := code;
  if code < 0 then
    raise EZDecompressionError.Create(string(_z_errmsg[2 - code]));
end;

function ZDecompressCheckWithoutBufferError(code: Integer): Integer;
begin
  Result := code;
  if code < 0 then
    if (code <> Z_BUF_ERROR) then
     raise EZDecompressionError.Create(string(_z_errmsg[2 - code]));
end;

function ZCompressCheck(code: Integer): Integer; overload;
begin
  Result := code;
  if code < 0 then
     raise EZCompressionError.Create(string(_z_errmsg[2 - code]));
end;

function ZCompressCheckWithoutBufferError(code: Integer): Integer; overload;
begin
  Result := code;
  if code < 0 then
    if (code <> Z_BUF_ERROR) then
     raise EZCompressionError.Create(string(_z_errmsg[2 - code]));
end;


procedure ZDecompressStream(inStream, outStream: TStream);
const
  bufferSize = 32768;
var
  zstream: TZStream;
  zresult: Integer;
  inBuffer: TBytes;
  outBuffer: TBytes;
  inSize: Integer;
  outSize: Integer;
begin
  SetLength(inBuffer, BufferSize);
  SetLength(outBuffer, BufferSize);
  FillChar(zstream, SizeOf(TZStream), 0);
  ZDecompressCheck(InflateInit(zstream));

  try
    inSize := inStream.Read(inBuffer[0], bufferSize);
    while inSize > 0 do
    begin
      zstream.next_in := @inBuffer[0];
      zstream.avail_in := inSize;

      repeat
        zstream.next_out := @outBuffer[0];
        zstream.avail_out := bufferSize;

        ZDecompressCheckWithoutBufferError(inflate(zstream, Z_NO_FLUSH));

        outSize := bufferSize - zstream.avail_out;

        outStream.Write(outBuffer[0], outSize);
      until (zstream.avail_in = 0) and (zstream.avail_out > 0);

      inSize := inStream.Read(inBuffer[0], bufferSize);
    end;

    repeat
      zstream.next_out := @outBuffer[0];
      zstream.avail_out := bufferSize;

      zresult := ZDecompressCheckWithoutBufferError(inflate(zstream, Z_FINISH));

      outSize := bufferSize - zstream.avail_out;

      outStream.Write(outBuffer[0], outSize);
    until (zresult = Z_STREAM_END) and (zstream.avail_out > 0);
  finally
    ZDecompressCheck(inflateEnd(zstream));
  end;
end;


procedure ZCompressStream(inStream, outStream: TStream; level: TZCompressionLevel = zcDefault);
const
  bufferSize = 32768;
var
  zstream: TZStreamRec;
  zresult: Integer;
  inBuffer: TBytes;
  outBuffer: TBytes;
  inSize: Integer;
  outSize: Integer;
begin
  SetLength(inBuffer, BufferSize);
  SetLength(outBuffer, BufferSize);
  FillChar(zstream, SizeOf(TZStreamRec), 0);

  ZCompressCheck(DeflateInit(zstream, ZLevels[level]));

  try
    inSize := inStream.Read(inBuffer[0], bufferSize);
    while inSize > 0 do
    begin
      zstream.next_in := @inBuffer[0];
      zstream.avail_in := inSize;

      repeat
        zstream.next_out := @outBuffer[0];
        zstream.avail_out := bufferSize;

        ZCompressCheckWithoutBufferError(deflate(zstream, Z_NO_FLUSH));

        // outSize := zstream.next_out - outBuffer;
        outSize := bufferSize - zstream.avail_out;

        outStream.Write(outBuffer[0], outSize);
      until (zstream.avail_in = 0) and (zstream.avail_out > 0);

      inSize := inStream.Read(inBuffer[0], bufferSize);
    end;

    repeat
      zstream.next_out := @outBuffer[0];
      zstream.avail_out := bufferSize;

      zresult := ZCompressCheck(deflate(zstream, Z_FINISH));

      // outSize := zstream.next_out - outBuffer;
      outSize := bufferSize - zstream.avail_out;

      outStream.Write(outBuffer[0], outSize);
    until (zresult = Z_STREAM_END) and (zstream.avail_out > 0);
  finally
    ZCompressCheck(deflateEnd(zstream));
  end;
end;




{$ENDIF}

{ TFormResFile }

class function TFormResFile.Encrypt(AInStream: TStream; AOutStream: TStream): Boolean;
var
  LVer: Word;
  LMem, LZipMem: TMemoryStream;
begin
  LMem := TMemoryStream.Create;
  try
    LMem.Write(HEADER, Length(HEADER));
    LVer := VERSION;
    LMem.Write(LVer, SizeOf(Word));
    LZipMem := TMemoryStream.Create;
    try
      AInStream.Position := 0;
      ZCompressStream(AInStream, LZipMem);
      XorStream(LZipMem);
      LMem.Write(LZipMem.Memory^, LZipMem.Size);
      AOutStream.Position := 0;
      LMem.SaveToStream(AOutStream);
      AOutStream.Position := 0;
      Result := True;
    finally
      LZipMem.Free;
    end;
  finally
    LMem.Free;
  end;
end;

class procedure TFormResFile.XorStream(AStream: TMemoryStream);
var
  EncodedBytes, ZipBytes:array of Byte; //编码字节
  BytesLength: Integer;
  I, L:integer;
const
  EncKey: array[0..15] of Byte = (
	$EB, $D4, $93, $48, $97, $0E, $38, $5B, $89, $B8, $2F, $D2, $31, $11, $AD, $BC);

begin
  AStream.Position := 0;
  BytesLength := AStream.Size;
  SetLength(EncodedBytes, BytesLength);
  SetLength(ZipBytes, BytesLength);
  AStream.Read(EncodedBytes[0], BytesLength);
  for I := 0 to BytesLength - 1 do
  begin
    L := I mod 16;
    ZipBytes[i] := EncodedBytes[i] xor EncKey[L];
  end;
  AStream.Clear; //清除以前的流
  AStream.Write(ZipBytes[0], BytesLength); //将异或后的字节集写入原流中
  AStream.Position := 0;
end;






end.
