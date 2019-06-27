package main

import "github.com/ying32/govcl/vcl/win"

func GetCurrentThreadId() uintptr {
	return win.GetCurrentThreadId()
}
