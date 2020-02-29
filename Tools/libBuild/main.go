package main

import (
	"flag"
	"log"
)

var (
	// 编译库类型
	libType = flag.String("lib", "", "要编译的目标库，可选为lcl、vcl")
	libArch = flag.String("arch", "all", "编译的目标库arch，可选为x86、x64、all")
)

func main() {
	flag.Parse()
	if *libType == "" {
		log.Fatal("lib参数不能为空。")
	}
}
