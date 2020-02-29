package main

import (
	"flag"
	"log"
	"runtime"
)

var (
	// 编译库类型
	libType = flag.String("lib", "all", "要编译的目标库，可选为lcl、vcl、all，仅Windows下有效。")
	libArch = flag.String("arch", "all", "编译的目标库arch，可选为x86、x64、all，仅Windows下有效")
)

func main() {
	flag.Parse()
	switch runtime.GOOS {
	case "windows":
	case "linux":

	case "darwin":
	default:
		log.Fatal("不支持：", runtime.GOOS)
	}
}
