//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package vcl

import (
	"reflect"
	"unsafe"

	"fmt"

	"github.com/ying32/govcl/vcl/api"
	. "github.com/ying32/govcl/vcl/types"
)

// ShowMessage 显示一个消息框
func ShowMessage(msg string) {
	api.DShowMessage(msg)
}

func ShowMessageFmt(format string, args ...interface{}) {
	ShowMessage(fmt.Sprintf(format, args...))
}

// MessageDlg 消息框，Buttons为按钮样式，祥见types.TMsgDlgButtons
func MessageDlg(Msg string, DlgType TMsgDlgType, Buttons ...uint8) int32 {
	return api.DMessageDlg(Msg, DlgType, NewSet(Buttons...), 0)
}

// CheckPtr 检测接口是否被实例化，如果已经实例化则返回实例指针
func CheckPtr(value interface{}) uintptr {
	switch value.(type) {
	case IObject:
		if reflect.ValueOf(value).Pointer() == 0 {
			return 0
		}
		return value.(IObject).Instance()
	}
	return 0
}

// As操作的简化
func getInstance(value interface{}) (uintptr, unsafe.Pointer) {
	var ptr uintptr
	switch value.(type) {
	case uintptr:
		// 一个对象来自已经存在的对象实例指针
		ptr = value.(uintptr)
	case unsafe.Pointer:
		// 一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
		ptr = uintptr(value.(unsafe.Pointer))
	case IObject:
		// 一个对象来自已经存在的对象实例
		ptr = CheckPtr(value)
	}
	return ptr, unsafe.Pointer(ptr)
}

// SelectDirectory1 选择目录
func SelectDirectory1(options TSelectDirOpts) (bool, string) {
	return api.DSelectDirectory1(options)
}

// SelectDirectory2 选择目录，一般 options默认是SdNewUI，parent默认为nil
func SelectDirectory2(caption, root string, options TSelectDirExtOpts, parent IObject) (bool, string) {
	return api.DSelectDirectory2(caption, root, options, CheckPtr(parent))
}

// SelectDirectory3 选择目录， options默认是SdNewUI，parent默认为nil
func SelectDirectory3(caption, root string, options ...uint8) (bool, string) {
	opts := NewSet(options...)
	if len(options) == 0 {
		opts = opts.Include(SdNewUI)
	}
	return SelectDirectory2(caption, root, opts, nil)
}

// ThreadSync 主线程中执行
func ThreadSync(fn TThreadProc) {
	api.DSynchronize(fn, 1)
}

// ThreadSyncVcl 主线程中执行，第二个参数决定是否使用Delphi自带的，此也只对libvcl生效，1使用消息，0使用Delphi自带的线程同步方法。
func ThreadSyncVcl(fn TThreadProc) {
	api.DSynchronize(fn, 0)
}

// InputBox 输入框
func InputBox(aCaption, aPrompt, aDefault string) string {
	return api.DInputBox(aCaption, aPrompt, aDefault)
}

// InputQuery 输入框
func InputQuery(aCaption, aPrompt string, value *string) bool {
	return api.DInputQuery(aCaption, aPrompt, value)
}

// 简化运行
func RunApp(forms ...interface{}) {
	Application.Initialize()
	Application.SetMainFormOnTaskBar(true)
	for i := 0; i < len(forms); i++ {
		Application.CreateForm(forms[i])
	}
	Application.Run()
}

// 不必引用rtl包来判断是否为lcl库
func LclLoaded() bool {
	return api.IsloadedLcl
}
