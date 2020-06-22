package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"archive/zip"
)

var (
	gPath string
)

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func main() {
	gopaths := os.Getenv("GOPATH")
	if gopaths == "" {
		panic("GOPATH为空！")
	}
	sp := strings.Split(gopaths, ";")
	libLCLBinResDir := strings.TrimSpace(sp[0])
	for _, s := range sp {
		s = strings.TrimSpace(s) + "/src/github.com/ying32/liblclbinres"
		if fileExists(s) {
			libLCLBinResDir = s
			break
		}
	}

	fmt.Println("找到路径")
	if !fileExists(libLCLBinResDir) {
		if err := os.MkdirAll(libLCLBinResDir, 0666); err != nil {
			panic(err)
		}
	}

	if len(os.Args) > 1 {
		//
		zipFileName := os.Args[1]
		fmt.Println(zipFileName)
		if strings.ToLower(path.Ext(zipFileName)) != ".zip" {
			panic("输入正确的zip包，如：“liblcl-2.0.2.zip”")
		}
		zz, err := zip.OpenReader(zipFileName)
		if err != nil {
			panic(err)
		}
		defer zz.Close()
		for _, ff := range zz.File {
			//fmt.Println(ff.Name)
			switch ff.Name {
			case "linux64-gtk2/liblcl.so":
				genresByte(readZipData(ff), libLCLBinResDir+"/liblcl_linux_amd64.go")
			case "win32/liblcl.dll":
				genresByte(readZipData(ff), libLCLBinResDir+"/liblcl_windows_386.go")
			case "win64/liblcl.dll":
				genresByte(readZipData(ff), libLCLBinResDir+"/liblcl_windows_amd64.go")
			}
		}

	} else {
		genresFile("../../Librarys/liblcl/win32/liblcl.dll", libLCLBinResDir+"/liblcl_windows_386.go")
		genresFile("../../Librarys/liblcl/win64/liblcl.dll", libLCLBinResDir+"/liblcl_windows_amd64.go")
		genresFile("../../Librarys/liblcl/linux64-gtk2/liblcl.so", libLCLBinResDir+"/liblcl_linux_amd64.go")
	}

	// macOS不支持这种，也不需要支持这种
	//genres("../liblcl/macOS64-cocoa/liblcl.dylib", "liblcl_darwin_amd64")
}

func readZipData(ff *zip.File) []byte {
	if rr, err := ff.Open(); err == nil {
		defer rr.Close()
		bs, err := ioutil.ReadAll(rr)
		if err != nil {
			return nil
		}
		return bs
	}
	return nil
}

//  zlib压缩
func ZlibCompress(input []byte) ([]byte, error) {
	var in bytes.Buffer
	w, err := zlib.NewWriterLevel(&in, zlib.BestCompression)
	if err != nil {
		return nil, err
	}
	_, err = w.Write(input)
	if err != nil {
		return nil, err
	}
	err = w.Close()
	if err != nil {
		return nil, err
	}
	return in.Bytes(), nil
}

func genresByte(input []byte, newFileName string) {
	fmt.Println("genFile: ", newFileName)
	if len(input) == 0 {
		fmt.Println("000000")
		return
	}

	crc32Val := crc32.ChecksumIEEE(input)

	//压缩
	bs, err := ZlibCompress(input)
	if err != nil {
		panic(err)
	}
	code := bytes.NewBuffer(nil)
	code.WriteString("package liblclbinres")
	code.WriteString("\r\n\r\n")
	code.WriteString(fmt.Sprintf("const CRC32Value uint32 = 0x%x\r\n\r\n", crc32Val))

	code.WriteString("var LCLBinRes = []byte(\"")
	for _, b := range bs {
		code.WriteString("\\x" + fmt.Sprintf("%.2x", b))
	}
	code.WriteString("\")\r\n")
	ioutil.WriteFile(newFileName, code.Bytes(), 0666)
}

// 生成字节的单元
func genresFile(fileName, newFileName string) {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	genresByte(bs, newFileName)
}
