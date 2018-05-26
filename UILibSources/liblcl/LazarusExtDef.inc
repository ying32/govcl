// 枚举都按4字节对齐
{$Z4+}

type
  // 重定定义
  TUnixDateTime = type LongInt; // int64改为LongInt。只精确到秒
  TPngImage = type TPortableNetworkGraphic;
  TMonthCalendar = type TCalendar;
  TJPEGPixelFormat = TPixelFormat;

  // 重定一个，主要是为了修改相关默认
  TGoForm = class(TForm)
  protected
    procedure ProcessResource; override;
  public
    constructor CreateNew(AOwner: TComponent; Num: Integer = 0); override;
  end;


constructor TGoForm.CreateNew(AOwner: TComponent; Num: Integer);
begin
  inherited CreateNew(AOwner, Num);
{$IFDEF WINDOWS}
  // Windows下统一VCL/LCL默认字体
  Font.Name := 'Tahoma';
  Font.Size := 8;
{$ENDIF}
end;

procedure TGoForm.ProcessResource;
begin
  // 没有使用窗口资源，不处理，处理就会报错的。
end;
