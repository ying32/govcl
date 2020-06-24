package main

import (
	"os"
	"runtime"
	"testing"
)

func TestAll(t *testing.T) {
	if runtime.GOOS == "windows" {
		testVCL(t)
	}
	testLCL(t)
}

func testLCL(t *testing.T) {

	switch runtime.GOOS {
	case "windows":
		// 32位的
		t.Log(buildLCL("i386", "F:/Golang/src/github.com/ying32/govcl/Tools/libBuild"))

		// 64位的
		t.Log(buildLCL("x86_64", "F:/Golang/src/github.com/ying32/govcl/Tools/libBuild/x64"))

	case "linux":
		// 64位的 linux
		t.Log(buildLCL("x86_64", "/home/ying32/genliblcl2/bin"))

	case "darwin":
		usrHome := os.Getenv("HOME")
		// 64位的 macOS
		t.Log(buildLCL("x86_64", usrHome+"/godev/gosrc/bin"))
	}

}

func testVCL(t *testing.T) {

	//   f:\program files (x86)\embarcadero\studio\19.0\bin\cgrc.exe -c65001 vcl.vrc -fovcl.res
	t.Log(buildVCL("32", "F:\\Golang\\src\\github.com\\ying32\\govcl\\Tools\\libBuild"))

	//   f:\program files (x86)\embarcadero\studio\19.0\bin\cgrc.exe -c65001 vcl.vrc -fovcl.res
	t.Log(buildVCL("64", "F:\\Golang\\src\\github.com\\ying32\\govcl\\Tools\\libBuild"))
}
