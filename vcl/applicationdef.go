package vcl

import (
	"fmt"

	. "github.com/ying32/govcl/vcl/api"
)

// CreateForm 一般不建议使用 NewForm，而优先使用CreateForm
// 这里要做兼容
// len(fields) = 2, 首个参数为：文件名或者字节， 第二个参数为输出的目标
// 目前只支持libvcl的，主要是解决dpi问题。
func (a *TApplication) CreateForm(fields ...interface{}) *TForm {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	// 参数的个数决定，创建窗口时是否使用缩放，此值需要 vcl.Application.SetFormScaled(true) 后才能生效。
	form := FormFromInst(Application_CreateForm(a.instance, len(fields) != 2))
	if len(fields) == 2 {
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
		a.fullFiledVal(form, fields[1])
		// 使用参数，则不返回
		return nil
	}
	return form
}

// SetFormScaled 设置全局窗口的Scaled
func (a *TApplication) SetFormScaled(val bool) {
	SetGlobalFormScaled(val)
}

//
func (a *TApplication) Run() {
	Application_Run(a.instance)
}

func (a *TApplication) Initialize() {
	Application_Initialize(a.instance)
}
