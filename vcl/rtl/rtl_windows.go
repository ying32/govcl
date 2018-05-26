package rtl

import (
	"gitee.com/ying32/govcl/vcl/api"
	"gitee.com/ying32/govcl/vcl/win"
)

// MainInstance EXE自身的实例
func MainInstance() uintptr {
	return win.GetSelfModuleHandle()
}

// SetReportMemoryLeaksOnShutdown 程序结束时报告内存泄露，总有2-4字节的未知泄露位置
func SetReportMemoryLeaksOnShutdown(v bool) {
	api.DSetReportMemoryLeaksOnShutdown(v)
}
