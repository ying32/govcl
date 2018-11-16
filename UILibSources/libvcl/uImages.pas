// 为了兼容Lazarus与Delphi对于资源窗口数据中图片读取的处理
// 比如在Delphi中TPngImage是没有Size标识的，但Lazarus中一定有
// 于是只能进行修改，统一都有图片Size标识，但在Delphi读取时经过判断后跳过相关的

unit uImages;

interface

uses
  System.Classes,
  System.SysUtils,
  Vcl.Graphics,
  Vcl.Imaging.pngimage,
  Vcl.Imaging.GIFImg,
  Vcl.Imaging.GIFConsts;

type

  TPngImage = class(Vcl.Imaging.pngimage.TPngImage)
  protected
    procedure ReadData(Stream: TStream); override;
    procedure WriteData(Stream: TStream); override;
  end;


  TGIFImage = class(Vcl.Imaging.GIFImg.TGIFImage)
  protected
    procedure ReadData(Stream: TStream); override;
    procedure WriteData(Stream: TStream); override;
  end;

  // 与lazarus保持一致
  TPortableNetworkGraphic = class(TPngImage)
  end;

implementation



{ TPngImage }

procedure TPngImage.ReadData(Stream: TStream);
const
  PngHeader: Array[0..7] of AnsiChar = (#137, #80, #78, #71, #13, #10, #26, #10);
var
  LHeader: array[0..7] of AnsiChar;
begin
  Stream.Read(LHeader[0], 8);
  if CompareMem(@PngHeader[0], @LHeader[0], 8) then
    Stream.Position := Stream.Position - 8
  else
    Stream.Position := Stream.Position - 4;
  inherited;
end;

procedure TPngImage.WriteData(Stream: TStream);
var
  Size: Longint;
  LMem: TMemoryStream;
begin
  LMem := TMemoryStream.Create;
  try
    SaveToStream(LMem);
    Size := LMem.Size;
    Stream.Write(Size, Sizeof(Size));
    if Size > 0 then
      Stream.Write(LMem.Memory^, Size);
  finally
    LMem.Free;
  end;
end;

{ TGIFImage }

procedure TGIFImage.ReadData(Stream: TStream);
const
  GifHeader: array[0..2] of Byte = ($47, $49, $46);
var
  LHeader: array[0..2] of Byte;
begin
  Stream.Read(LHeader[0], 3);

  if CompareMem(@GifHeader[0], @LHeader[0], 3) then
    Stream.Position := Stream.Position - 3
  else
    Stream.Position := Stream.Position + 1;
  inherited;
end;

procedure TGIFImage.WriteData(Stream: TStream);
var
  Size: Longint;
  LMem: TMemoryStream;
begin
  LMem := TMemoryStream.Create;
  try
    SaveToStream(LMem);
    Size := LMem.Size;
    Stream.Write(Size, Sizeof(Size));
    if Size > 0 then
      Stream.Write(LMem.Memory^, Size);
  finally
    LMem.Free;
  end;
end;

initialization
   TPicture.UnregisterGraphicClass(Vcl.Imaging.pngimage.TPngImage);
   TPicture.UnregisterGraphicClass(Vcl.Imaging.GIFImg.TGIFImage);

   // 重新注册
   TPicture.RegisterFileFormat('PNG', 'Portable Network Graphics', TPortableNetworkGraphic);
   TPicture.RegisterFileFormat('PNG', 'Portable Network Graphics', TPngImage);
   TPicture.RegisterFileFormat('GIF', sGIFImageFile, TGIFImage);

finalization


end.
