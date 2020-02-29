package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func GetGoVCLDir() string {
	for _, path := range strings.Split(os.Getenv("GOPATH"), ";") {
		pp := path + "/src/github.com/ying32/govcl"
		if checkFileExists(pp) {
			return pp
		}
	}
	return ""
}

func GetLibProjectFile(file string) string {
	path := GetGoVCLDir() + "/UILibSources" + file
	if checkFileExists(path) {
		if runtime.GOOS == "windows" {
			return strings.Replace(path, "/", "\\", -1)
		} else {
			return strings.Replace(path, "\\", "/", -1)
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

func ExtractFilePath(path string) string {
	filename := filepath.Base(path)
	return path[:len(path)-len(filename)]
}

func execCmd(objFileDir, cmdStr string) error {
	extName := ""
	if runtime.GOOS == "windows" {
		extName = ".bat"
	}
	bashFileName := objFileDir + string(os.PathSeparator) + "compile" + extName
	if err := ioutil.WriteFile(bashFileName, []byte(cmdStr), 0755); err != nil {
		return err
	}
	if err := os.Chmod(bashFileName, 0755); err != nil {
		return err
	}
	cmd := exec.Command(bashFileName)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
