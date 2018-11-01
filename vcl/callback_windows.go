package vcl

import "C"
import (
	"syscall"
)

var (
	eventCallback   = syscall.NewCallback(eventCallbackProc)
	messageCallback = syscall.NewCallback(messageCallbackProc)
)
