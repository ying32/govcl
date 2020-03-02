package main

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"os"
	"runtime"
)

const (
	lclCompileCommandline = `"{{.fpc}}" -T{{.platform}} -P{{.arch}} -MObjFPC -Scghi -CX -Cg -O4 -XX{{if eq .platform "darwin"}} -k-framework -kCocoa{{end}} -l -vewnhibq -Fi{{.objFileDir}}/lib/{{.arch}}-{{.platform}} -Fu{{.lazarusDir}}/components/printers/lib/{{.arch}}-{{.platform}}/{{.lclType}} -Fu{{.lazarusDir}}/components/lazcontrols/lib/{{.arch}}-{{.platform}}/{{.lclType}} -Fu{{.lazarusDir}}/components/cairocanvas/lib/{{.arch}}-{{.platform}}/{{.lclType}} -Fu{{.lazarusDir}}/lcl/units/{{.arch}}-{{.platform}}/{{.lclType}} -Fu{{.lazarusDir}}/components/datetimectrls/lib/{{.arch}}-{{.platform}}/{{.lclType}} -Fu{{.lazarusDir}}/lcl/units/{{.arch}}-{{.platform}} -Fu{{.lazarusDir}}/components/lazutils/lib/{{.arch}}-{{.platform}} -Fu{{.lazarusDir}}/packager/units/{{.arch}}-{{.platform}} -Fu{{.objFileDir}}/ -FU{{.objFileDir}}/lib/{{.arch}}-{{.platform}}/ -FE{{.binFileDir}}/ -o{{.binFileDir}}/liblcl.{{.ext}} -dLCL -dLCL{{.lclType}} -dUSED_BY_LAZLOGGER_DUMMY "{{.projectFileName}}"`
)

func buildLCL(arch, binFileDir string) error {

	lazarusDir, fpcExe := GetLazarusDir()
	projectFileName := GetLibProjectFile("/liblcl/lcl.lpr")

	// lcl.lpi 这里面要修改版本号之类的
	// ProductVersion="1.2.8.0"

	if !checkFileExists(lazarusDir) {
		return errors.New("未找到lazarus目录。")
	}
	if !checkFileExists(projectFileName) {
		return errors.New("liblcl工程文件未找到。")
	}
	if arch != "i386" && arch != "x86_64" {
		return errors.New("arch错误，只能为i386或者x86_64")
	}

	lazarusDir = fixDirName(lazarusDir)
	objFileDir := fixDirName(GetObjFileDir("liblcl"))
	binFileDir = fixDirName(binFileDir)

	fmt.Println("lazarusDir:", lazarusDir)
	fmt.Println("projectFileName: ", projectFileName)
	fmt.Println("objFileDir: ", objFileDir)
	fmt.Println("binFileDir: ", binFileDir)

	extName := ""
	lclType := ""
	platform := ""
	switch runtime.GOOS {
	case "linux":
		extName = "so"
		lclType = "gtk2"
		platform = "linux"

	case "darwin":
		extName = "dylib"
		lclType = "cocoa"
		//if arch == "i386" {
		//	lclType = "carbon"
		//}
		platform = "darwin"

	default:
		extName = "dll"
		lclType = "win32"
		if arch == "i386" {
			platform = "win32"
		} else {
			platform = "win64"
		}
	}

	var buf bytes.Buffer
	tmp := template.New("lcl")
	tmp.Parse(lclCompileCommandline)
	tmp.Execute(&buf, map[string]interface{}{
		"fpc":             fpcExe,
		"arch":            arch,     // i386 x86_64
		"platform":        platform, // win32 win64 darwin linux
		"lazarusDir":      lazarusDir,
		"lclType":         lclType, // win32 cocoa carbon gtk2 gtk3
		"ext":             extName, // dll so dylib
		"projectFileName": projectFileName,
		"objFileDir":      objFileDir, // 编译时中间文件目录
		"binFileDir":      binFileDir,
	})

	// liblcl二进制输出目录不存在，则创建
	if !checkFileExists(binFileDir) {
		if err := os.MkdirAll(binFileDir, 0775); err != nil {
			return err
		}
	}

	// 中间文件输出
	libOutDir := objFileDir + "/lib/" + arch + "-" + platform
	if !checkFileExists(libOutDir) {
		if err := os.MkdirAll(libOutDir, 0775); err != nil {
			return err
		}
	}

	fmt.Println(buf.String())
	return execCmd(objFileDir, buf.String())
}
