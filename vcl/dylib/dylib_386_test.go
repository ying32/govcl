package dylib

import (
	"runtime"
	"testing"
)

func ToUInt64(r1, r2 uintptr) uint64 {
	ret := uint64(r2)
	ret = uint64(ret<<32) + uint64(r1)
	return ret
}

func TestAll(t *testing.T) {
	testdll := NewLazyDLL("testdll_" + runtime.GOARCH)
	defer testdll.Close()

	var int64Max int64 = 9223372036854775807
	var uint64Max uint64 = 18446744073709551615
	var int64Val1 int64 = 0xFF2233445500
	var int64Val2 int64 = 0xFF6677889900

	t.Log(int64Max, uint64Max, int64Val1, int64Val2)
	t.Log("--------------------------------------------------------------")

	// 无参数，但返回值为int64
	ResultValue := testdll.NewProc("ResultValue")

	r1, r2, _ := ResultValue.Call()
	t.Log("r1:", r1, ", r2:", r2)
	if int64(ToUInt64(r1, r2)) == int64Max {
		t.Log("测试结果正确")
	} else {
		t.Fatal("测试结果错误，正确值为", int64Max, "，当前值:", int64(ToUInt64(r1, r2)))
	}

	// 2个参数，类型为int64，返回为uint64
	ResultValAndParam := testdll.NewProc("ResultValAndParam")
	r1, r2, _ = ResultValAndParam.Call(uintptr(int64Val1), uintptr(int64Val1>>32), uintptr(int64Val2), uintptr(int64Val2>>32))
	t.Log("r1:", r1, ", r2:", r2)
	if ToUInt64(r1, r2) == uint64Max {
		t.Log("测试结果正确")
	} else {
		t.Fatal("测试结果错误，正确值为", uint64Max, "，当前值:", ToUInt64(r1, r2))
	}

	t.Log("------------------------------32bit-----------------------------")

	var intMax int32 = 2147483647
	var uintMax uint32 = 4294967295
	var intVal1 int32 = 0x0F112233
	var intVal2 int32 = 0x0F445566

	t.Log(intMax, uintMax, intVal1, intVal2)
	t.Log("--------------------------------------------------------------")

	// 无参数，但返回值为int64
	ResultValue32 := testdll.NewProc("ResultValue32")

	r1, r2, _ = ResultValue32.Call()
	t.Log("r1:", r1, ", r2:", r2)
	if int32(r1) == intMax {
		t.Log("测试结果正确")
	} else {
		t.Fatal("测试结果错误，正确值为", intMax, "，当前值:", int32(r1))
	}

	// 2个参数，类型为int64，返回为uint64
	ResultValAndParam32 := testdll.NewProc("ResultValAndParam32")
	r1, r2, _ = ResultValAndParam32.Call(uintptr(intVal1), uintptr(intVal2))
	t.Log("r1:", r1, ", r2:", r2)
	if uint32(r1) == uintMax {
		t.Log("测试结果正确")
	} else {
		t.Fatal("测试结果错误，正确值为", uintMax, "，当前值:", uint32(r1))
	}

}
