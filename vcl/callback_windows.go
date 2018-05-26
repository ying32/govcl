package vcl

import (
	"syscall"
)

var (
	callbackStdcall = syscall.NewCallback(callbackProc)
)
