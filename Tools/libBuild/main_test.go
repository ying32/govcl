package main

import (
	"testing"
)

func TestAll(t *testing.T) {
	testVCL(t)
	//testLCL(t)
}

func testLCL(t *testing.T) {
	// 32位的
	t.Log(buildLCL("F:/lazarus/fpc/3.0.4/bin/x86_64-win64/fpc.exe",
		"i386",
		"win32",
		"F:/lazarus",
		"F:/Golang/src/gitee.com/ying32/genliblcl2/liblcl/lcl.lpr",
		"F:/Golang/src/github.com/ying32/govcl/Tools/libBuild/liblcl",
		"F:/Golang/src/github.com/ying32/govcl/Tools/libBuild"))

	// 64位的
	t.Log(buildLCL("F:/lazarus/fpc/3.0.4/bin/x86_64-win64/fpc.exe",
		"x86_64",
		"win64",
		"F:/lazarus",
		"F:/Golang/src/gitee.com/ying32/genliblcl2/liblcl/lcl.lpr",
		"F:/Golang/src/github.com/ying32/govcl/Tools/libBuild/liblcl",
		"F:/Golang/src/github.com/ying32/govcl/Tools/libBuild/x64"))
}

func testVCL(t *testing.T) {

	//   f:\program files (x86)\embarcadero\studio\19.0\bin\cgrc.exe -c65001 vcl.vrc -fovcl.res
	buildVCL("f:\\program files (x86)\\embarcadero\\studio\\19.0",
		"F:\\Users\\Public\\Documents\\Embarcadero\\Studio\\19.0",
		"32",
		"F:\\Golang\\src\\gitee.com\\ying32\\genlibvcl\\libvcl\\vcl.dpr",
		"F:\\Golang\\src\\github.com\\ying32\\govcl\\Tools\\libBuild\\libvcl",
		"F:\\Golang\\src\\github.com\\ying32\\govcl\\Tools\\libBuild")

	//   f:\program files (x86)\embarcadero\studio\19.0\bin\cgrc.exe -c65001 vcl.vrc -fovcl.res
	buildVCL("f:\\program files (x86)\\embarcadero\\studio\\19.0",
		"F:\\Users\\Public\\Documents\\Embarcadero\\Studio\\19.0",
		"64",
		"F:\\Golang\\src\\gitee.com\\ying32\\genlibvcl\\libvcl\\vcl.dpr",
		"F:\\Golang\\src\\github.com\\ying32\\govcl\\Tools\\libBuild\\libvcl",
		"F:\\Golang\\src\\github.com\\ying32\\govcl\\Tools\\libBuild")
}
