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
	ErrPixelDataEmpty        = errors.New("The pixel data is empty")
	ErrUnsupportedDataFormat = errors.New("Unsupported pixel data format")
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
// LCL貌似不分丢失透明度，VCL会。。。。
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
func ToGIFImage(img image.Image, quality int) (*vcl.TGIFImage, error) {
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
