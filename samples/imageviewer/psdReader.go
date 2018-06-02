//=======================================================
// 功能：转换photoshop PSD格式为TBITMAP
//
// 日期：2018-06-01
// 作者：ying32
// QQ  ：1444386932
// 注  ：参考出处
//       Photoshop文件格式参考官网
//         http://www.adobe.com/devnet-apps/photoshop/fileformatashtml/#50577409_72092
//
//  这个只是 github.com/ying32/govcl中imageviewer例程所写，解析psd。
//=======================================================
/*

以下为基本格式，

File header section
The file header contains the basic properties of the image.
Table 2–7: File header
|--------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Length    |    Name       |    Description                                                                                                                         |
|-----------|---------------|----------------------------------------------------------------------------------------------------------------------------------------|
| 4 bytes   |    Signature  |    Always equal to 8BPS. Do not try to read the ﬁle if the signature does not match this value.                                        |
|-----------|---------------|----------------------------------------------------------------------------------------------------------------------------------------|
| 2 bytes   |    Version    |    Always equal to 1. Do not try to read the ﬁle if the version does not match this value.                                             |
|-----------|---------------|----------------------------------------------------------------------------------------------------------------------------------------|
| 6 bytes   |    Reserved   |    Must be zero.                                                                                                                       |
|-----------|---------------|----------------------------------------------------------------------------------------------------------------------------------------|
| 2 bytes   |    Channels   |    The number of channels in the image, including any alpha chan-nels. Supported range is 1 to 24.                                     |
|-----------|---------------|----------------------------------------------------------------------------------------------------------------------------------------|
| 4 bytes   |    Rows       |    The height of the image in pixels. Supported range is 1 to 30,000.                                                                  |
|-----------|---------------|----------------------------------------------------------------------------------------------------------------------------------------|
| 4 bytes   |    Columns    |    The width of the image in pixels. Supported range is 1 to 30,000.                                                                   |
|-----------|---------------|----------------------------------------------------------------------------------------------------------------------------------------|
| 2 bytes   |    Depth      |    The number of bits per channel. Supported values are 1, 8, and 16.                                                                  |
|-----------|---------------|----------------------------------------------------------------------------------------------------------------------------------------|
| 2 bytes   |    Mode       |    The color mode of the ﬁle. Supported values are: Bitmap=0; Grayscale=1; Indexed=2; RGB=3; CMYK=4; Multichannel=7; Duotone=8; Lab=9. |
|--------------------------------------------------------------------------------------------------------------------------------------------------------------------|


Table 2–8: Color mode data
|---------------------------------------------------------------------|
| Length   |  Name       | Description                                |
|----------|-------------|--------------------------------------------|
| 4 bytes  |  Length     | The length of the following color data.    |
|----------|-------------|--------------------------------------------|
| Variable |  Color data | The color data.                            |
|---------------------------------------------------------------------|


Table 2–9: Image resources
|---------------------------------------------------------------------|
| Length    | Name       | Description                                |
|-----------|------------|--------------------------------------------|
| 4 bytes   | Length     | Length of image resource section.          |
|-----------|------------|--------------------------------------------|
| Variable  | Resources Image resources                               |
|---------------------------------------------------------------------|


Table 2–10: Layer and mask information
|--------------------------------------------------------------------------|
| Length    | Name       | Description                                     |
|-----------|------------|-------------------------------------------------|
| 4 bytes   | Length     | Length of the miscellaneous information section.|
|-----------|------------|-------------------------------------------------|
| Variable  | Layers     | Layer info. See table 2–12.                     |
|-----------|------------|-------------------------------------------------|
| Variable  | Global     | Global layer mask info. See table 2–19.         |
|           | layer mask |                                                 |
|--------------------------------------------------------------------------|


Table 2–11: Image data
|----------------------------------------------------------------------------------|
| Length    | Name         | Description                                           |
|-----------|--------------|-------------------------------------------------------|
| 2 bytes   | Compression  | Compression method. Raw data = 0, RLE compressed = 1. |
|-----------|--------------|-------------------------------------------------------|
| Variable  | Data         | The image data. Planar order = RRR GGG BBB, etc       |
|----------------------------------------------------------------------------------|

*/

package main

import (
	"bytes"
	"encoding/binary"
	"io"
	"io/ioutil"

	"unsafe"

	"runtime"

	"errors"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

// psd文件读取，将psd转为一个TBitmap

const (
	MODE_BITMAP       = 0
	MODE_GRAYSCALE    = 1
	MODE_INDEXED      = 2
	MODE_RGB          = 3
	MODE_CMYK         = 4
	MODE_MULTICHANNEL = 7
	MODE_DUOTONE      = 8
	MODE_LAB          = 9

	MAX_PSD_CHANNELS = 24
)

// PsdToBitmap psd文件转bmp格式
func PsdToBitmap(aFileName string, bmp *vcl.TBitmap) error {
	if bmp == nil || !bmp.IsValid() {
		return errors.New("bmp不能为空！")
	}
	// 文件签名
	var psdSignature = []byte{0x38, 0x42, 0x50, 0x53} // 8BPS
	// 读入文件
	bs, err := ioutil.ReadFile(aFileName)
	if err != nil {
		return err
	}

	buffer := bytes.NewReader(bs)
	// 签名
	signature := make([]byte, 4)
	binary.Read(buffer, binary.BigEndian, signature)
	if bytes.Compare(signature, psdSignature) != 0 {
		return errors.New("签名错误！")
	}
	// 版本
	var version uint16
	binary.Read(buffer, binary.BigEndian, &version)
	if version != 1 {
		return errors.New("版本错误，必须等于1！")
	}
	// 跳过
	buffer.Seek(6, io.SeekCurrent)
	// 通道 支持 1 to 24.
	var ChannelCount uint16
	binary.Read(buffer, binary.BigEndian, &ChannelCount)
	if ChannelCount <= 0 || ChannelCount > MAX_PSD_CHANNELS {
		return errors.New("通道总数不正确！")
	}

	// 读size
	var width, height uint32
	binary.Read(buffer, binary.BigEndian, &height)
	binary.Read(buffer, binary.BigEndian, &width)

	if width <= 0 || height <= 0 {
		return errors.New("读宽度或者高度错误！")
	}

	// 每个通道的位数， 支持1,8和16
	var bitsPerChannel uint16
	binary.Read(buffer, binary.BigEndian, &bitsPerChannel)
	if bitsPerChannel != 8 {
		return errors.New("每个通道的位数错误，只支持8！")
	}

	// 颜色模式
	var colorMode uint16
	binary.Read(buffer, binary.BigEndian, &colorMode)
	if colorMode != MODE_RGB {
		return errors.New("颜色模式错误，只支持RGB！")
	}

	// 颜色数据，跳过
	var colorDataLen uint32
	binary.Read(buffer, binary.BigEndian, &colorDataLen)
	buffer.Seek(int64(colorDataLen), io.SeekCurrent)

	// 图片资源，跳过
	var imageResoucesLen uint32
	binary.Read(buffer, binary.BigEndian, &imageResoucesLen)
	buffer.Seek(int64(imageResoucesLen), io.SeekCurrent)

	// 全局层掩码信息，跳过
	var globalLayerMaskInfoLen uint32
	binary.Read(buffer, binary.BigEndian, &globalLayerMaskInfoLen)
	buffer.Seek(int64(globalLayerMaskInfoLen), io.SeekCurrent)

	// 压缩模式 Raw=0, RLE=1
	var compressionType uint16
	binary.Read(buffer, binary.BigEndian, &compressionType)
	if compressionType != 0 && compressionType != 1 {
		return errors.New("不支持的压缩格式，当前只支持模式为0或者1")
	}
	// 读数据
	psdPixels := make([]byte, width*height*4)
	unPackPSD(buffer, width, height, psdPixels, ChannelCount, compressionType)

	bmp.SetPixelFormat(types.Pf32bit)
	bmp.SetSize(int32(width), int32(height))

	if runtime.GOOS == "windows" || runtime.GOOS == "linux" {
		for h := int(height - 1); h >= 0; h-- {
			ptr := bmp.ScanLine(int32(h))
			for w := 0; w < int(width*4); w++ {
				index := h*(int(width)*4) + w
				*(*byte)(unsafe.Pointer(ptr + uintptr(w))) = psdPixels[index]
			}
		}
	} else {
		// 填充，左下角为起点
		for h := int(height - 1); h >= 0; h-- {
			ptr := bmp.ScanLine(int32(h))
			for w := 0; w < int(width); w++ {
				index := (h*int(width) + w) * 4
				*(*byte)(unsafe.Pointer(ptr + uintptr(w*4))) = psdPixels[index+3]
				*(*byte)(unsafe.Pointer(ptr + uintptr(w*4+1))) = psdPixels[index+2]
				*(*byte)(unsafe.Pointer(ptr + uintptr(w*4+2))) = psdPixels[index+1]
				*(*byte)(unsafe.Pointer(ptr + uintptr(w*4+3))) = psdPixels[index]
			}
		}
	}
	return nil
}

// 解包psd文件
func unPackPSD(buffer *bytes.Reader, width, height uint32, pixels []byte, channelCount, compressionType uint16) {
	var (
		channelArr = []byte{2, 1, 0, 3}
		defaultArr = []byte{0, 0, 0, 255}
	)
	pixelCount := width * height
	if compressionType == 1 {
		buffer.Seek(int64(height*uint32(channelCount)*2), io.SeekCurrent)
		for c := 0; c <= 3; c++ {
			n := 0
			channel := channelArr[c]
			if uint16(channel) >= channelCount {
				for n = 0; n < int(pixelCount); n++ {
					pixels[n*4+int(channel)] = defaultArr[channel]
				}
			} else { // 非压缩
				count := 0
				for count < int(pixelCount) {
					nlen, _ := buffer.ReadByte()
					if nlen == 128 {
						// 啥也不做
					} else if nlen < 128 { // 非RLE
						nlen++
						count += int(nlen)
						for nlen != 0 {
							pixels[n*4+int(channel)], _ = buffer.ReadByte()
							n++
							nlen--
						}
					} else if nlen > 128 { // RLE
						nlen = nlen ^ 0xFF
						nlen += 2
						val, _ := buffer.ReadByte()
						count += int(nlen)
						for nlen != 0 {
							pixels[n*4+int(channel)] = val
							n++
							nlen--
						}
					}
				}
			}
		}
	} else {
		for c := 0; c <= 3; c++ {
			channel := channelArr[c]
			if uint16(channel) > channelCount {
				for n := 0; n < int(pixelCount); n++ {
					pixels[n*4+int(channel)] = defaultArr[channel]
				}
			} else {
				for n := 0; n < int(pixelCount); n++ {
					pixels[n*4+int(channel)], _ = buffer.ReadByte()
				}
			}
		}
	}
}
