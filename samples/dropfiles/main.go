package main

import (
	"github.com/ying32/govcl/vcl"

	"fmt"

	"github.com/ying32/govcl/vcl/types"
)

func main() {

	vcl.Application.SetIconResId(3)
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)

	mainForm := vcl.Application.CreateForm()
	mainForm.SetCaption("DropFiles测试")
	mainForm.SetPosition(types.PoScreenCenter)
	mainForm.EnabledMaximize(false)
	mainForm.SetWidth(300)
	mainForm.SetHeight(200)

	// 注意，当在Windows上使用时如果使用了UAC，则无法收到消息
	// 需要使用未公开的winapi   ChangeWindowMessageFilter 或 ChangeWindowMessageFilterEx 根据系统版本不同使用其中的，然后添加
	// ChangeWindowMessageFilterEx(pnl_Drag.Handle, WM_DROPFILES, MSGFLT_ALLOW, LChangeFilterStruct);消息

	// Delphi代码举例
	/*
			  用法举例， 下面是一个使用了 Administrator 权限但又想不让系统过滤掉 WM_DROPFILES 消息的做法
		    // 判断操作系统的版本号，大于6.0为Vista，Xp不使用这个函数
			  if (TOSVersion.Major = 6) and (TOSVersion.Minor = 0) then
			  begin
				@ChangeWindowMessageFilter := GetProcAddress(GetModuleHandle(user32), 'ChangeWindowMessageFilter');
				if Assigned(ChangeWindowMessageFilter) then
				  ChangeWindowMessageFilter(WM_DROPFILES, MSGFLT_ADD);
			  end else
			  // 大于等于6.1为Win7及以上版本
			  if (TOSVersion.Major = 6) and (TOSVersion.Minor >= 1) then
			  begin
				TLog.D('当前为Win7或Win7以上操作系统');
				@ChangeWindowMessageFilterEx := GetProcAddress(GetModuleHandle(user32), 'ChangeWindowMessageFilterEx');
				if Assigned(ChangeWindowMessageFilterEx) then
				begin
				  LChangeFilterStruct.cbSize := SizeOf(TChangeFilterStruct);
				  ChangeWindowMessageFilterEx(pnl_Drag.Handle, WM_DROPFILES, MSGFLT_ALLOW, LChangeFilterStruct);
				end;
			  end;

	*/
	fmt.Println("allowDropFiles1:", mainForm.AllowDropFiles())
	mainForm.SetAllowDropFiles(true)
	fmt.Println("allowDropFiles2:", mainForm.AllowDropFiles())
	mainForm.SetOnDropFiles(func(sender vcl.IObject, aFileNames []string) {
		fmt.Println("当前拖放文件事件执行，文件数：", len(aFileNames))
		for i, s := range aFileNames {
			fmt.Println("index:", i, ", filename:", s)
		}
	})
	vcl.Application.Run()
}
