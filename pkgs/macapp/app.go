//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------
// +build darwin

/*
  使用方法: 根据go同一个包中的执行顺序，创建一个最小文件名的go文件，然后在里面写入，如：

package main

import _ "github.com/ying32/govcl/pkgs/macapp"
*/

package macapp

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

// 自定义包信息
type pkgFile struct {
	Copyright    string // copy right ....
	Locale       string // zh_CN en_US
	DevRegion    string // China En English
	IconFileName string // .icns file
	Files        []struct {
		Src  string // C:/xxx/a.txt
		Dest string // /Contents/MacOS/a.txt or /Contents/a.txt or /Contents/Resources/a.txt
	}
}

const (
	defaultPkgName = "apppkg.conf" // 打包的文件，一个json格式的。
)

func copyFile(src, dest string) error {
	filedest, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer filedest.Close()
	filesrc, err := os.Open(src)
	if err != nil {
		return err
	}
	defer filesrc.Close()
	_, err = io.Copy(filedest, filesrc)
	return err
}

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

func getdylib() string {
	env := os.Getenv("GOPATH")
	if env == "" {
		return ""
	}
	for _, s := range strings.Split(env, ":") {
		s += "/bin/liblcl.dylib"
		if fileExists(s) {
			return s
		}
	}
	return ""
}

func init() {
	runWithMacOSApp()
}

// 以一个Mac下的app形式运行
// 调试下使用，正式发布的时候虽然可 以不用去掉，但也不咋好
func runWithMacOSApp() {

	if strings.Contains(os.Args[0], ".app/Contents/MacOS") {
		return
	}

	execName := os.Args[0][strings.LastIndex(os.Args[0], "/")+1:]
	macContentsDir := os.Args[0] + ".app/Contents"
	macOSDir := macContentsDir + "/MacOS"
	macResources := macContentsDir + "/Resources"
	execFile := macOSDir + "/" + execName
	if !fileExists(macOSDir) {
		if err := os.MkdirAll(macOSDir, 0755); err != nil {
			return
		}
	}

	if !fileExists(macResources) {
		os.MkdirAll(macResources, 0755)
	}

	resName := fmt.Sprintf("%s/%s.icns", macResources, execName)
	if !fileExists(resName) {
		ioutil.WriteFile(resName, macOSAppIcon, 0666)
	}

	liblclFileName := macOSDir + "/liblcl.dylib"
	// 文件不存在，复制
	if !fileExists(liblclFileName) {
		libFileName := getdylib()
		if fileExists(libFileName) {
			copyFile(libFileName, liblclFileName)
		}
	} else {
		// 文件存在，对比后更新
		libFileName := getdylib()
		if fileExists(libFileName) {
			f1, _ := os.Stat(libFileName)
			f2, _ := os.Stat(liblclFileName)
			if f1.Size() != f2.Size() {
				copyFile(libFileName, liblclFileName)
			}
		}
	}

	plistFileName := macContentsDir + "/Info.plist"
	if !fileExists(plistFileName) {
		datas := map[string]string{
			"execName":  execName,
			"devRegion": "China", // China English
			"locale":    "zh_CN", //os.Getenv("LANG"),
			"copyright": "copyright xxxx",
		}
		buff := bytes.NewBuffer([]byte{})
		tmp := template.New("file")
		tmp.Parse(infoplist)
		tmp.Execute(buff, datas)
		ioutil.WriteFile(plistFileName, buff.Bytes(), 0666)
	}

	pkgInfoFileName := macContentsDir + "/PkgInfo"
	if !fileExists(pkgInfoFileName) {
		ioutil.WriteFile(pkgInfoFileName, pkgInfo, 0666)
	}

	if copyFile(os.Args[0], execFile) == nil {
		os.Chmod(execFile, 0755)
		var args []string
		if len(os.Args) > 1 {
			args = os.Args[1:]
		}
		cmd := exec.Command(execFile, args...)
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			os.Exit(1)
		} else {
			os.Exit(0)
		}
	}
}
