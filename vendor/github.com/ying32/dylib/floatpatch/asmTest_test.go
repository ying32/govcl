package floatpatch

import (
	"testing"

	"github.com/ying32/govcl/vcl/dylib"
)

var (
	libTest               = dylib.NewLazyDLL("Project12.dll")
	_testFloat32_MaxValue = libTest.NewProc("testFloat32_MaxValue")
	_testFloat32_MinValue = libTest.NewProc("testFloat32_MinValue")
	_testFloat32_Epsilon  = libTest.NewProc("testFloat32_Epsilon")

	_testFloat64_MaxValue = libTest.NewProc("testFloat64_MaxValue")
	_testFloat64_MinValue = libTest.NewProc("testFloat64_MinValue")
	_testFloat64_Epsilon  = libTest.NewProc("testFloat64_Epsilon")
)

func TestASM(t *testing.T) {

	t.Log("---------------float32----------------")
	/*
		Epsilon:Single = 1.4012984643248170709e-45;
		MaxValue:Single =  340282346638528859811704183484516925440.0;
		MinValue:Single = -340282346638528859811704183484516925440.0;
		PositiveInfinity:Single =  1.0 / 0.0;
		NegativeInfinity:Single = -1.0 / 0.0;
		NaN:Single = 0.0 / 0.0;
	*/

	t.Log("testFloat32_MaxValue:", testFloat32_MaxValue())
	t.Log("testFloat32_MinValue:", testFloat32_MinValue())
	t.Log("testFloat32_Epsilon:", testFloat32_Epsilon())

	t.Log("---------------float64----------------")
	/*
		Epsilon:Double = 4.9406564584124654418e-324;
		MaxValue:Double =  1.7976931348623157081e+308;
		MinValue:Double = -1.7976931348623157081e+308;
		PositiveInfinity:Double =  1.0 / 0.0;
		NegativeInfinity:Double = -1.0 / 0.0;
		NaN:Double = 0.0 / 0.0;
	*/

	t.Log("testFloat64_MaxValue:", testFloat64_MaxValue())
	t.Log("testFloat64_MinValue:", testFloat64_MinValue())
	t.Log("testFloat64_Epsilon:", testFloat64_Epsilon())

}

func testFloat32_MaxValue() float32 {
	_testFloat32_MaxValue.Call()
	return Getfloat32()
}

func testFloat32_MinValue() float32 {
	_testFloat32_MinValue.Call()
	return Getfloat32()
}

func testFloat32_Epsilon() float32 {
	_testFloat32_Epsilon.Call()
	return Getfloat32()
}

//-- float64
func testFloat64_MaxValue() float64 {
	_testFloat64_MaxValue.Call()
	return Getfloat64()
}

func testFloat64_MinValue() float64 {
	_testFloat64_MinValue.Call()
	return Getfloat64()
}

func testFloat64_Epsilon() float64 {
	_testFloat64_Epsilon.Call()
	return Getfloat64()
}
