package vcl

import (
	"fmt"

	. "github.com/ying32/govcl/vcl/api"
)

var globalFormScaled bool = false

/*
 CreateForm 一般不建议使用 NewForm，而优先使用CreateForm
 这里要做兼容
 len(fields) = 1
   这时，可不加载窗口的资源文件，自己构造一种，如下：

   type TForm1 struct {
      *vcl.TForm // 必须首个，无需命名，否则失败。
      Button1 *vcl.TButton
   }

   var form1 *TForm1

   vcl.Application.Create(&form1)

 len(fields) = 2, 首个参数为：文件名或者字节， 第二个参数为输出的目标
   目前只支持libvcl的，主要是解决dpi问题。

*/
func (a *TApplication) CreateForm(fields ...interface{}) *TForm {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	// 参数的个数决定，创建窗口时是否使用缩放，此值需要 vcl.Application.SetFormScaled(true) 后才能生效。
	form := FormFromInst(Application_CreateForm(a.instance, len(fields) != 2))

	// 当等于1时使用手动构造的一种
	if len(fields) == 1 {
		a.fullFiledVal(form, fields[0], false)
	} else if len(fields) == 2 { // 等于2时，使用资源中的
		switch fields[0].(type) {
		case string:
			ResFormLoadFromFile(fields[0].(string), CheckPtr(form))
		case []byte:
			mem := NewMemoryStreamFromBytes(fields[0].([]byte))
			defer mem.Free()
			ResFormLoadFromStream(CheckPtr(mem), CheckPtr(form))
		default:
			panic("error")
		}
		a.fullFiledVal(form, fields[1], true)
		// 使用参数，则不返回
		return nil
	}
	return form
}

// SetFormScaled 设置全局窗口的Scaled
func (a *TApplication) SetFormScaled(val bool) {
	globalFormScaled = val
	SetGlobalFormScaled(val)
}

//
func (a *TApplication) Run() {
	Application_Run(a.instance)
}

func (a *TApplication) Initialize() {
	Application_Initialize(a.instance)
}
