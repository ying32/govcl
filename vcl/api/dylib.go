package api

import "runtime"

var (
	// 全局导入库
	libvcl = loadUILib()

	// 判断是否加载lcl库，一般用于windows上
	IsloadedLcl = false

	// 专用于判断是否为Windows系统
	IsWindows = runtime.GOOS == "windows"
)

// VCL或者LCL
type TLibType int32

const (
	LtVCL TLibType = iota + 0
	LtLCL
)
