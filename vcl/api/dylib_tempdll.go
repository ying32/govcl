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
	tempDLLDir := fmt.Sprintf("%s/liblcl/%x", os.TempDir(), liblclbinres.CRC32Value)
	// create liblcl: $tempdir/liblcl/{crc32}/liblcl.{ext}
	if !fileExists(tempDLLDir) {
		if err := os.MkdirAll(tempDLLDir, 0775); err != nil {
			return false, ""
		}
	}
	tempDLLFileName := fmt.Sprintf("%s/%s", tempDLLDir, getDLLName())
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
			if os.Remove(tempDLLFileName) != nil {
				return false, ""
			}
		}
	}
	return true, tempDLLFileName
}
