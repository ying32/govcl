package vcl

import (
	"syscall"
)

var (
	eventCallback      = syscall.NewCallback(eventCallbackProc)
	messageCallback    = syscall.NewCallback(messageCallbackProc)
	threadSyncCallback = syscall.NewCallback(threadSyncCallbackProc)
)
