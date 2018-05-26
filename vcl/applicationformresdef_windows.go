//----------------------------------------
// 加载文件或者内存中的窗口资源文件功能
// CreateFormFromFile与CreateFormFromStream内部函数不在开源范围内
// 需要配合窗口设计器使用。
// 设计器是独立于govcl的，设计器的目的是用于简化窗口的创建。设计器不开源，不免费。
//
// 例：
//    type TMainForm struct {
//        *vcl.TForm  // 不要写名称，使用隐式的
//        Button1 *vcl.TButton // 注意，设计器中一定要首字母大写，否则不能使用反射填充
//    }
//
//	  var mainForm *TMainForm
//
//    func main() {
//        ...
//        // 文件加载方式
//        vcl.Application.CreateFormFromFile(filename, &mainForm)
//        // 流加载方式
//        // mem := vcl.NewMemoryStream()
//        // mem.Write([]byte(...))
//        // vcl.Application.CreateFormFromStream(mem, &mainForm)
//        // mem.Free()
//        // 资源加载方式
//        //vcl.Application.CreateFormFromResourceName("ResName", &mainForm)
//        ...
//     }

// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

package vcl

import (
	"gitee.com/ying32/govcl/vcl/api"
	"gitee.com/ying32/govcl/vcl/rtl"
)

// CreateFormFromResourceName 资源类型为 RT_RCDATA
func (a *TApplication) CreateFormFromResourceName(resName string, out interface{}) {
	f := a.CreateForm()
	api.ResFormLoadFromResourceName(rtl.MainInstance(), resName, CheckPtr(f))
	a.fullFiledVal(f, out)
}
