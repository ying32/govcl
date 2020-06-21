//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// +build windows linux
// +build tempdll
// +build !memorydll

// 指令为：target == windows || target == linux && tempdll && !memorydll

package api

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"hash/crc32"
	"io"
	"io/ioutil"
	"os"
	"runtime"

	"github.com/ying32/liblclbinres"
)

// $GOPATH/github.com/ying32/liblclbinres

func ZlibUnCompressToFile(destFileName string, input []byte) error {
	r, err := zlib.NewReader(bytes.NewReader(input))
	if err != nil {
		return err
	}
	defer r.Close()
	fi, err := os.Create(destFileName)
	if err != nil {
		return err
	}
	defer fi.Close()
	_, err = io.Copy(fi, r)
	if err != nil {
		return nil
	}
	return nil
}

var (
	spStr = string(os.PathSeparator)
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

func checkAndReleaseDLL() (bool, string) {
	tempDLLDir := os.TempDir() + spStr + "liblcl" + spStr + liblclbinres.Version + spStr + runtime.GOARCH
	// 创建目录 tempdir/liblcl/{version}/{arch}/liblcl.{ext}
	if !fileExists(tempDLLDir) {
		if err := os.MkdirAll(tempDLLDir, 0664); err != nil {
			return false, ""
		}
	} else {
		fmt.Println("目录已创建")
	}
	tempDLLFileName := tempDLLDir + spStr + getDLLName()
	// test crc32
	if fileExists(tempDLLFileName) {
		bs, err := ioutil.ReadFile(tempDLLFileName)
		if err == nil {
			if crc32.ChecksumIEEE(bs) != liblclbinres.CRC32Value {
				os.Remove(tempDLLFileName)
			}
		}
	}
	if !fileExists(tempDLLFileName) {
		if err := ZlibUnCompressToFile(tempDLLFileName, liblclbinres.LCLBinRes); err != nil {
			os.Remove(tempDLLFileName)
			return false, ""
		}
	} else {
		fmt.Println("liblcl已经存在")
	}
	return true, tempDLLFileName
}
