//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// 实现一些Go的Image转vcl/lcl的
package bitmap

import (
	"bytes"
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"unsafe"

	"github.com/ying32/govcl/vcl/types"

	"github.com/ying32/govcl/vcl"
)

var (
	ErrPixelDataEmpty        = errors.New("the pixel data is empty")
	ErrUnsupportedDataFormat = errors.New("unsupported pixel data format")
)

// 将Go的Image转为VCL/LCL的 TPngImage
// 返回的Png对象用完记得Free掉
func ToPngImage(img image.Image) (*vcl.TPngImage, error) {
	buff := bytes.NewBuffer([]byte{})
	if err := png.Encode(buff, img); err != nil {
		return nil, err
	}
	mem := vcl.NewMemoryStreamFromBytes(buff.Bytes())
	defer mem.Free()
	mem.SetPosition(0)
	obj := vcl.NewPngImage()
	obj.LoadFromStream(mem)
	return obj, nil
}

// 32bit bmp，丢失透明度
// 返回的Bmp对象用完记得Free掉
// LCL貌似不会丢失透明度，VCL会。。。。
func ToBitmap(img image.Image) (*vcl.TBitmap, error) {
	switch img.(type) {
	case *image.RGBA:
		data, _ := img.(*image.RGBA)
		return toBitmap(img.Bounds().Size().X, img.Bounds().Size().Y, data.Pix)

	case *image.NRGBA:
		data, _ := img.(*image.NRGBA)
		return toBitmap(img.Bounds().Size().X, img.Bounds().Size().Y, data.Pix)

	default:
		return nil, ErrUnsupportedDataFormat
	}
}

// 将Go的Image转为VCL/LCL的 TJPEGImage
// 返回的jpg对象用完记得Free掉
func ToJPEGImage(img image.Image, quality int) (*vcl.TJPEGImage, error) {
	buff := bytes.NewBuffer([]byte{})
	if err := jpeg.Encode(buff, img, &jpeg.Options{quality}); err != nil {
		return nil, err
	}
	mem := vcl.NewMemoryStreamFromBytes(buff.Bytes())
	defer mem.Free()
	mem.SetPosition(0)
	obj := vcl.NewJPEGImage()
	obj.LoadFromStream(mem)
	return obj, nil
}

// 将Go的Image转为VCL/LCL的 TGIFImage
// 返回的gif对象用完记得Free掉
func ToGIFImage(img image.Image) (*vcl.TGIFImage, error) {
	buff := bytes.NewBuffer([]byte{})
	if err := gif.Encode(buff, img, &gif.Options{NumColors: 256}); err != nil {
		return nil, err
	}
	mem := vcl.NewMemoryStreamFromBytes(buff.Bytes())
	defer mem.Free()
	mem.SetPosition(0)
	obj := vcl.NewGIFImage()
	obj.LoadFromStream(mem)
	return obj, nil
}

func toBitmap(width, height int, pix []uint8) (*vcl.TBitmap, error) {
	if len(pix) == 0 {
		return nil, ErrPixelDataEmpty
	}
	bmp := vcl.NewBitmap()
	bmp.SetPixelFormat(types.Pf32bit)
	bmp.SetSize(int32(width), int32(height))
	// 填充，左下角为起点
	for h := height - 1; h >= 0; h-- {
		ptr := bmp.ScanLine(int32(h))
		for w := 0; w < width; w++ {
			index := (h*width + w) * 4
			*(*byte)(unsafe.Pointer(ptr + uintptr(w*4))) = pix[index+pixIndex[0]]
			*(*byte)(unsafe.Pointer(ptr + uintptr(w*4+1))) = pix[index+pixIndex[1]]
			*(*byte)(unsafe.Pointer(ptr + uintptr(w*4+2))) = pix[index+pixIndex[2]]
			*(*byte)(unsafe.Pointer(ptr + uintptr(w*4+3))) = pix[index+pixIndex[3]]
		}
	}
	return bmp, nil
}

// 将vcl/lcl的Graphic对象转为Go的Image
func ToGoImage(obj *vcl.TGraphic) (image.Image, error) {
	if obj == nil || !obj.IsValid() {
		return nil, errors.New("obj is invalid")
	}
	buff := bytes.NewBuffer([]byte{})
	mem := vcl.NewMemoryStream()
	defer mem.Free()
	obj.SaveToStream(mem)
	mem.SetPosition(0)
	_, bs := mem.Read(int32(mem.Size()))
	buff.Write(bs)
	if obj.InheritsFrom(vcl.TBitmapClass()) {
		img := image.NewRGBA(image.Rect(0, 0, int(obj.Width()), int(obj.Height())))
		// 复制数据到Pix中。。。
		// 先不管他了，以后再说
		//img.Pix
		return img, nil
	} else if obj.InheritsFrom(vcl.TPngImageClass()) {
		img, err := png.Decode(buff)
		if err != nil {
			return nil, err
		}
		return img, nil
	} else if obj.InheritsFrom(vcl.TJPEGImageClass()) {
		img, err := jpeg.Decode(buff)
		if err != nil {
			return nil, err
		}
		return img, nil
	} else if obj.InheritsFrom(vcl.TGIFImageClass()) {
		img, err := gif.Decode(buff)
		if err != nil {
			return nil, err
		}
		return img, nil
	}
	return nil, errors.New("unknown error")
}
