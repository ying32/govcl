package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func GetLibProjectFile(file string) string {
	gopath := strings.Split(os.Getenv("GOPATH"), ";")
	for _, path := range gopath {
		pp := path + "/src/github.com/ying32/govcl/UILibSources"
		if checkFileExists(pp) {
			pp += file
			if runtime.GOOS == "windows" {
				return strings.Replace(pp, "/", "\\", -1)
			} else {
				return strings.Replace(pp, "\\", "/", -1)
			}
		}
	}
	return ""
}

func GetObjFileDir(path string) string {
	return os.TempDir() + string(os.PathSeparator) + path
}

func checkFileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func fixDirName(dir string) string {
	if strings.HasSuffix(dir, "/") || strings.HasSuffix(dir, "\\") {
		dir = dir[:len(dir)-1]
	}
	return dir
}

func execCmd(objFileDir, cmdStr string) error {
	extName := ""
	if runtime.GOOS == "windows" {
		extName = ".bat"
	}
	bashFileName := objFileDir + string(os.PathSeparator) + "compile" + extName
	if err := ioutil.WriteFile(bashFileName, []byte(cmdStr), 0755); err != nil {
		os.Chmod(bashFileName, 0755)
	}
	cmd := exec.Command(bashFileName)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
