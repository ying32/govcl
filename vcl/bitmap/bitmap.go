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
	ErrBitmapInvalid         = errors.New("bitmap invalid")
	ErrObjIsInvalid          = errors.New("object is invalid")
	ErrUnknownError          = errors.New("unknown error")
)

// CN: 将Go的Image转为VCL/LCL的 TPngImage
//     返回的Png对象用完记得Free掉
// EN: Convert the image of go to TPngImage of VCL/LCL
//     Remember to free the returned png object
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

// CN: 32bit bmp
//     返回的Bmp对象用完记得Free掉
// EN: 32bit bmp
//     Remember to free the returned bmp object
func ToBitmap(img image.Image) (*vcl.TBitmap, error) {
	bmp := vcl.NewBitmap()
	if err := ToBitmap2(img, bmp); err != nil {
		defer bmp.Free()
		return nil, err
	}
	return bmp, nil
}

func ToBitmap2(img image.Image, bmp *vcl.TBitmap) error {
	if bmp == nil || !bmp.IsValid() {
		return ErrBitmapInvalid
	}
	switch img.(type) {
	case *image.RGBA:
		data, _ := img.(*image.RGBA)
		err := toBitmap(img.Bounds().Size().X, img.Bounds().Size().Y, data.Pix, bmp)
		if err != nil {
			return err
		}
	case *image.NRGBA:
		data, _ := img.(*image.NRGBA)
		err := toBitmap(img.Bounds().Size().X, img.Bounds().Size().Y, data.Pix, bmp)
		if err != nil {
			return err
		}
	default:
		return ErrUnsupportedDataFormat
	}
	return nil
}

// CN: 将Go的Image转为VCL/LCL的 TJPEGImage
//     返回的jpg对象用完记得Free掉
// EN: Convert the image of go to TJPEGImage of VCL/LCL
//     Remember to free the returned jpg object
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

// CN: 将Go的Image转为VCL/LCL的 TGIFImage
//     返回的gif对象用完记得Free掉
// EN: Convert the image of go to TGIFImage of VCL/LCL
//     Remember to free the returned GIF object
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

func toBitmap(width, height int, pix []uint8, bmp *vcl.TBitmap) error {
	if len(pix) == 0 {
		return ErrPixelDataEmpty
	}
	if bmp == nil || !bmp.IsValid() {
		return ErrBitmapInvalid
	}

	bmp.SetSize(int32(width), int32(height))
	// 总是32位，不然没办法透明。
	bmp.SetPixelFormat(types.Pf32bit)
	bmp.SetHandleType(types.BmDIB)

	bmp.BeginUpdate(false)
	defer bmp.EndUpdate(false)

	// 填充，左下角为起点
	for h := height - 1; h >= 0; h-- {
		ptr := bmp.ScanLine(int32(h))
		for w := 0; w < width; w++ {
			index := (h*width + w) * 4
			c := (*rgba)(unsafe.Pointer(ptr + uintptr(w*4)))
			c.R = pix[index+0]
			c.G = pix[index+1]
			c.B = pix[index+2]
			c.A = pix[index+3]
		}
	}
	return nil
}

// CN: 将vcl/lcl的Graphic对象转为Go的Image
// EN: Convert the graphic object of VCL/LCL to the image of go
func ToGoImage(obj *vcl.TGraphic) (image.Image, error) {
	if obj == nil {
		return nil, ErrObjIsInvalid
	}
	if !obj.IsValid() {
		return nil, ErrObjIsInvalid
	}
	buff := bytes.NewBuffer([]byte{})
	mem := vcl.NewMemoryStream()
	defer mem.Free()
	obj.SaveToStream(mem)
	mem.SetPosition(0)
	_, bs := mem.Read(int32(mem.Size()))
	buff.Write(bs)
	if obj.Is().Bitmap() {
		height := int(obj.Height())
		width := int(obj.Width())
		img := image.NewRGBA(image.Rect(0, 0, width, height))

		bmp := vcl.AsBitmap(obj)
		switch bmp.PixelFormat() {

		// 还有待测试。。。
		case types.Pf24bit:

			for h := height - 1; h >= 0; h-- {
				ptr := bmp.ScanLine(int32(h))
				for w := 0; w < width; w++ {
					index := (h*width + w) * 4
					c := (*rgb)(unsafe.Pointer(ptr + uintptr(w*3)))
					img.Pix[index+0] = c.R
					img.Pix[index+1] = c.G
					img.Pix[index+2] = c.B
					img.Pix[index+3] = 0
				}
			}

		case types.Pf32bit:
			for h := height - 1; h >= 0; h-- {
				ptr := bmp.ScanLine(int32(h))
				for w := 0; w < width; w++ {
					index := (h*width + w) * 4
					c := (*rgba)(unsafe.Pointer(ptr + uintptr(w*4)))
					img.Pix[index+0] = c.R
					img.Pix[index+1] = c.G
					img.Pix[index+2] = c.B
					img.Pix[index+3] = c.A
				}
			}
		default:
			//case types.Pf1bit:
			//case types.Pf4bit:
			//case types.Pf8bit:
			//case types.Pf15bit:
			//case types.Pf16bit:
			//case types.PfCustom:
			return nil, ErrUnsupportedDataFormat
		}
		return img, nil
	} else if obj.Is().PngImage() {
		img, err := png.Decode(buff)
		if err != nil {
			return nil, err
		}
		return img, nil
	} else if obj.Is().JPEGImage() {
		img, err := jpeg.Decode(buff)
		if err != nil {
			return nil, err
		}
		return img, nil
	} else if obj.Is().GIFImage() {
		img, err := gif.Decode(buff)
		if err != nil {
			return nil, err
		}
		return img, nil
	}
	return nil, ErrUnknownError
}
