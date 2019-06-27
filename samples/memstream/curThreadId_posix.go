// +build !windows

package main

func GetCurrentThreadId() uintptr {
	return 0
}
