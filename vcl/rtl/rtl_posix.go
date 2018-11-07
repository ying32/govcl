// +build !windows

package rtl

// MainInstance EXE自身的实例
func MainInstance() uintptr {
	return 0
}
