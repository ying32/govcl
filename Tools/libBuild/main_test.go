package main

import (
	"fmt"
	"testing"
)

func TestAll(t *testing.T) {

	testVCL(t)
	testLCL(t)
}

func testLCL(t *testing.T) {

	lazarusDir := GetLazarusDir()
	fpcExe := fmt.Sprintf("%s\\fpc\\3.0.4\\bin\\x86_64-win64\\fpc.exe", lazarusDir)
	libSrcFileName := GetLibProjectFile("/liblcl/lcl.lpr")

	// 32位的
	t.Log(buildLCL(fpcExe,
		"i386",
		"win32",
		lazarusDir,
		libSrcFileName,
		GetObjFileDir("liblcl"),
		"F:/Golang/src/github.com/ying32/govcl/Tools/libBuild"))

	// 64位的
	t.Log(buildLCL(fpcExe,
		"x86_64",
		"win64",
		lazarusDir,
		libSrcFileName,
		GetObjFileDir("liblcl"),
		"F:/Golang/src/github.com/ying32/govcl/Tools/libBuild/x64"))
}

func testVCL(t *testing.T) {

	bsdDir, userDir := GetBsdDir()
	libSrcFileName := GetLibProjectFile("/libvcl/vcl.dpr")

	//   f:\program files (x86)\embarcadero\studio\19.0\bin\cgrc.exe -c65001 vcl.vrc -fovcl.res
	t.Log(buildVCL(bsdDir, userDir,
		"32",
		libSrcFileName,
		GetObjFileDir("libvcl"),
		"F:\\Golang\\src\\github.com\\ying32\\govcl\\Tools\\libBuild"))

	//   f:\program files (x86)\embarcadero\studio\19.0\bin\cgrc.exe -c65001 vcl.vrc -fovcl.res
	t.Log(buildVCL(bsdDir, userDir,
		"64",
		libSrcFileName,
		GetObjFileDir("libvcl"),
		"F:\\Golang\\src\\github.com\\ying32\\govcl\\Tools\\libBuild"))
}
