package main

import (
	"flag"
	"os"
	"strings"
)

var (
	// 编译库类型
	libType = flag.String("lib", "", "要编译的目标库，可选为lcl、vcl")
	libArch = flag.String("arch", "all", "编译的目标库arch，可选为x86、x64、all")
)

func main() {
	// HKEY_CURRENT_USER\Software\Embarcadero\BDS\19.0

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
