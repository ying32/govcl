// +build windows

package win

var (
	_GetWindowLongPtr = user32dll.NewProc("GetWindowLongW")
	_SetWindowLongPtr = user32dll.NewProc("SetWindowLongW")
)
