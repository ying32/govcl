package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

func GetGoVCLDir() string {
	getPath := func(key string) string {
		for _, path := range strings.Split(os.Getenv(key), ";") {
			pp := path + "/src/github.com/ying32/govcl"
			if checkFileExists(pp) {
				return pp
			}
		}
		return ""
	}
	ret := getPath("GOPATH")
	if ret == "" {
		return getPath("GOROOT")
	}
	return ret
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
	} else {
		extName = ".sh"
		cmdStr = "#!/bin/sh\n" + cmdStr
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

func readLazarusEnvFile(szPath string) (lazarusDir string, fpcExe string) {
	envFile := szPath + "\\lazarus\\environmentoptions.xml"
	if runtime.GOOS != "windows" {
		envFile = strings.Replace(envFile, "\\", "/", -1)
	}
	if checkFileExists(envFile) {
		bs, err := ioutil.ReadFile(envFile)
		if err == nil {
			reg := regexp.MustCompile(`\<LazarusDirectory Value\=\"(.+?)"\>`)
			matchs := reg.FindSubmatch(bs)
			if len(matchs) >= 2 {
				lazarusDir = string(matchs[1])
			}
			reg = regexp.MustCompile(`\<CompilerFilename Value\=\"(.+?)\"\>`)
			matchs = reg.FindSubmatch(bs)
			if len(matchs) >= 2 {
				fpcExe = string(matchs[1])
			}
		}
	}
	return
}
