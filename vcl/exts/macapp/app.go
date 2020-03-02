//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------
// +build darwin

package macapp

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
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
func runWithMacOSApp() error {

	if len(os.Args) == 1 {
		if strings.Contains(os.Args[0], "Contents/MacOS") {
			return fmt.Errorf("_")
		}

		execName := os.Args[0][strings.LastIndex(os.Args[0], "/")+1:]
		macContentsDir := os.Args[0] + ".app/Contents"
		macOSDir := macContentsDir + "/MacOS"
		macResources := macContentsDir + "/Resources"
		execFile := macOSDir + "/" + execName
		if !fileExists(macOSDir) {
			if err := os.MkdirAll(macOSDir, 0755); err != nil {
				return err
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
			ioutil.WriteFile(plistFileName, []byte(fmt.Sprintf(infoplist, execName, execName, execName, execName)), 0666)
		}

		pkgInfoFileName := macContentsDir + "/PkgInfo"
		if !fileExists(pkgInfoFileName) {
			ioutil.WriteFile(pkgInfoFileName, pkgInfo, 0666)
		}

		if copyFile(os.Args[0], execFile) == nil {
			os.Chmod(execFile, 0755)
			args := os.Args[:len(os.Args)-1]
			cmd := exec.Command(execFile, args...)
			cmd.Run()
			os.Exit(0)
		}
	}

	return nil
}
