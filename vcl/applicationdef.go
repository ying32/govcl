package vcl

import . "github.com/ying32/govcl/vcl/api"

var globalFormScaled bool = false

/*
 TApplication.CreateForm 一般不建议使用NewForm，而优先使用CreateForm

 用法：
  1、mainForm := vcl.Application.CreateForm()    // 直接返回
  2、vcl.Application.CreateForm(&mainForm)       // 无资源加载，只会绑定窗口的事件，不会绑定子组件事件
  3、vcl.Application.CreateForm(&mainForm, true) // 无资源加载，当第二个参数为true时在OnFormCreate调用完后会绑定子组件事件

  4、vcl.Application.CreateForm("form1.gfm", &mainForm)  // 从资源文件中填充子组件，并绑定所有事件
  5、vcl.Application.CreateForm(form1Bytes, &mainForm)   // 从字节中填充子组件，并绑定所有事件
*/
func (a *TApplication) CreateForm(fields ...interface{}) *TForm {
	return FormFromObj(resObjtBuild(0, nil, a.instance, fields...))
	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Println("TApplication.CreateForm Error: ", err)
	//	}
	//}()
	//
	//var fullSubComponent bool
	//var afterBindSubComponentsEvents bool
	//var field1 interface{}
	//
	//// 初始创建时是否使用缩放
	//initScale := len(fields) != 2
	//if len(fields) >= 2 {
	//	switch fields[1].(type) {
	//	case bool:
	//		initScale = true
	//	default:
	//		initScale = false
	//	}
	//}
	//
	//// 由参数的个数决定，创建窗口时是否使用缩放，此值需要 vcl.Application.SetFormScaled(true) 后才能生效。
	//form := FormFromInst(Application_CreateForm(a.instance, initScale))
	//
	//switch len(fields) {
	//case 1:
	//	field1 = fields[0]
	//	fullSubComponent = false
	//	afterBindSubComponentsEvents = false
	//
	//case 2:
	//	switch fields[1].(type) {
	//	// 当第二个参数为bool时，表示不填充子组件，为true表示之后绑定事件
	//	case bool:
	//		field1 = fields[0]
	//		fullSubComponent = false
	//		afterBindSubComponentsEvents = fields[1].(bool)
	//	default:
	//		// 第二个参数类型不为bool时，填充子组件为true，之后绑定事件为false
	//		field1 = fields[1]
	//		fullSubComponent = true
	//		afterBindSubComponentsEvents = false
	//		switch fields[0].(type) {
	//		case string:
	//			ResFormLoadFromFile(fields[0].(string), CheckPtr(form))
	//		case []byte:
	//			mem := NewMemoryStreamFromBytes(fields[0].([]byte))
	//			defer mem.Free()
	//			ResFormLoadFromStream(CheckPtr(mem), CheckPtr(form))
	//		}
	//	}
	//default:
	//	return form
	//}
	//fullFiledVal(form, field1, fullSubComponent, afterBindSubComponentsEvents)
	//return nil
}

// SetFormScaled 设置全局窗口的Scaled
func (a *TApplication) SetFormScaled(val bool) {
	globalFormScaled = val
	SetGlobalFormScaled(val)
}

// Run 运行APP
func (a *TApplication) Run() {
	Application_Run(a.instance)
}

// Initialize 初始APP信息
func (a *TApplication) Initialize() {
	Application_Initialize(a.instance)
}
