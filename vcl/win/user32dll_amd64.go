// +build windows

package win

var (
	_GetWindowLongPtr = user32dll.NewProc("GetWindowLongPtrW")
	_SetWindowLongPtr = user32dll.NewProc("SetWindowLongPtrW")
)
