//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package api

import "runtime"

var (

	// 专用于判断是否为Windows系统
	IsWindows = runtime.GOOS == "windows"

	// 全局导入库
	libvcl = loadUILib()

	// 初始导入字符串指针函数
	_ = initStdString()
)

// VCL或者LCL
type TLibType int32

const (
	LtVCL TLibType = iota + 0
	LtLCL
)
