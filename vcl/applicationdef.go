//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package vcl

import (
	. "github.com/ying32/govcl/vcl/api"
)

/*
 TApplication.CreateForm 一般不建议使用NewForm，而优先使用CreateForm

  -------------------- 用法一--------------------------------------
  1、mainForm := vcl.Application.CreateForm()    // 直接返回，不推荐使用
  例:
    mainForm := vcl.Application.CreateForm()
    mainForm.SetOnClick(func(sender vcl.IObject) {
        vcl.ShowMessage("msg")
    })


  -------------------- 用法二--------------------------------------
  2、vcl.Application.CreateForm(&mainForm)       // 无资源或者自动加载对应类名的资源，当无资源时只会绑定窗口的事件，不会绑定子组件事件，有资源则自动绑定所有事件
  例：
    type TMainForm struct {
        *vcl.TForm
    }

    var mainForm *TMainForm
    vcl.Application.CreateForm(&mainForm)

    func (f *TMainForm)OnFormCreate(sender vcl.IObject) {
        fmt.Println("FormCreate")
    }

    func (f *TMainForm)OnFormClick(sender vcl.IObject) {
        vcl.ShowMessage("click")
    }


  -------------------- 用法三--------------------------------------
  3、vcl.Application.CreateForm(&mainForm, true) // 无资源或者自动加载对应类名的资源，当第二个参数为true时在OnFormCreate调用完后会绑定子组件事件(当查找到对应的资源则第二个参数无效)
  例：
    type TMainForm struct {
        *vcl.TForm
        Btn1 *vcl.TButton
    }

    var mainForm *TMainForm
    vcl.Application.CreateForm(&mainForm, true)

    func (f *TMainForm)OnFormCreate(sender vcl.IObject) {
        fmt.Println("FormCreate")
        f.Btn1 = vcl.NewButton(f)
        f.Btn1.SetParent(f)
    }

    func (f *TMainForm)OnFormClick(sender vcl.IObject) {
        vcl.ShowMessage("click")
    }

    func (f *TMainForm)OnBtn1Click(Sender vcl.IObject) {
        vcl.ShowMessage("Btn1 Click")
    }


  -------------------- 用法四--------------------------------------
  // 将来准备废弃
  4、vcl.Application.CreateForm("form1.gfm", &mainForm)  // 从资源文件中填充子组件，并绑定所有事件
  -------------------- 用法五--------------------------------------
  // 将来准备废弃
  5、vcl.Application.CreateForm(form1Bytes, &mainForm)   // 从字节中填充子组件，并绑定所有事件
*/

// 创建一个TForm。
//
// Create a TForm.
func (a *TApplication) CreateForm(fields ...interface{}) *TForm {
	return AsForm(resObjectBuild(0, nil, a.instance, fields...))
}

// 运行APP。
//
// Run the app.
func (a *TApplication) Run() {
	Application_Run(a.instance)
}

// 初始APP信息。
//
// Initial APP information.
func (a *TApplication) Initialize() {
	Application_Initialize(a.instance)
}
