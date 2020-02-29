// +build !arm

/*
   用于解决syscall不能返回浮点的补丁，用法参考REAMDME.md

*/

package floatpatch

func Getfloat32() float32

func Getfloat64() float64
