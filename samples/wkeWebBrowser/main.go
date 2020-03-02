// +build windows,386

package main

/*
	wke浏览器测试，wke暂时只能在windows上跑，也只支持32位
	现选择的wke头文件及二进制为cexer维护的一个支
	https://github.com/cexer/wke

    其实可以直接调用wke.h吧，没有试过

	注：目前这个只是一个例程，暂不会做封装正式入到govcl中。只支持win32，并且限libvcl库。
	    wke也许不能满足你，要求高的无解了，cef暂时不打算弄进来。
*/

import (
	"github.com/ying32/govcl/vcl"

	"fmt"

	_ "github.com/ying32/govcl/vcl/exts/winappres"
	"github.com/ying32/govcl/vcl/exts/wke"
	"github.com/ying32/govcl/vcl/types"
)

var wkeBrw *wke.TWkeWebBrowser

func main() {
	// 一定要调用的初始
	wke.Initialize()
	defer wke.Finalize()

	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)

	mainForm := vcl.Application.CreateForm()
	mainForm.SetCaption("Wke测试")
	mainForm.SetPosition(types.PoScreenCenter)
	mainForm.SetWidth(800)
	mainForm.SetHeight(600)
	mainForm.SetShowHint(true)

	pnl := vcl.NewPanel(mainForm)
	pnl.SetParent(mainForm)
	pnl.SetAlign(types.AlTop)

	// 导航按钮区域
	lPnl := vcl.NewPanel(mainForm)
	lPnl.SetParent(pnl)
	lPnl.SetWidth(110)
	lPnl.SetBevelOuter(types.BvNone)
	lPnl.SetAlign(types.AlLeft)

	// 返回按钮
	btnBack := vcl.NewButton(mainForm)
	btnBack.SetParent(lPnl)
	btnBack.SetCaption("<")
	btnBack.SetHint("后退")
	btnBack.SetBounds(5, (lPnl.Height()-30)/2, 30, 30)
	btnBack.SetOnClick(func(sender vcl.IObject) {
		wkeBrw.GoBack()
	})

	// 前进按钮
	btnForward := vcl.NewButton(mainForm)
	btnForward.SetParent(lPnl)
	btnForward.SetCaption(">")
	btnForward.SetHint("前进")
	btnForward.SetBounds(btnBack.Left()+btnBack.Width()+5, (lPnl.Height()-30)/2, 30, 30)
	btnForward.SetOnClick(func(sender vcl.IObject) {
		wkeBrw.GoForward()
	})

	// 重载按钮
	btnReload := vcl.NewButton(mainForm)
	btnReload.SetParent(lPnl)
	btnReload.SetCaption("O")
	btnReload.SetHint("重载")
	btnReload.SetBounds(btnForward.Left()+btnForward.Width()+5, (lPnl.Height()-30)/2, 30, 30)
	btnReload.SetOnClick(func(sender vcl.IObject) {
		wkeBrw.Reload()
	})

	// 中间地址栏
	cPnl := vcl.NewPanel(mainForm)
	cPnl.SetParent(pnl)
	cPnl.SetBevelOuter(types.BvNone)
	cPnl.SetAlign(types.AlClient)
	edit := vcl.NewEdit(mainForm)
	edit.SetParent(cPnl)
	edit.SetAlign(types.AlClient)
	edit.Font().SetSize(10)
	//edit.SetName("editURL")
	// https://github.com/ying32
	edit.SetText("https://www.baidu.com")
	edit.SetAlignWithMargins(true)
	edit.Margins().SetBounds(10, 5, 10, 5)

	// 跳转按钮区域
	rPnl := vcl.NewPanel(mainForm)
	rPnl.SetParent(pnl)
	rPnl.SetWidth(80)
	rPnl.SetBevelOuter(types.BvNone)
	rPnl.SetAlign(types.AlRight)

	btnGo := vcl.NewButton(mainForm)
	btnGo.SetParent(rPnl)
	btnGo.SetCaption("转到")
	btnGo.SetLeft(1)
	btnGo.SetHeight(30)
	btnGo.SetTop(int32((rPnl.Height() - btnGo.Height()) / 2))

	btnGo.SetOnClick(func(sender vcl.IObject) {
		if edit.Text() != "" {
			wkeBrw.Load(edit.Text())
		}
	})

	pnl = vcl.NewPanel(mainForm)
	pnl.SetParent(mainForm)
	pnl.SetAlign(types.AlClient)
	pnl.SetOnResize(func(sender vcl.IObject) {
		if wkeBrw != nil {
			wkeBrw.MoveWindow(0, 0, pnl.Width(), pnl.Height())
		}
	})

	// 优先接收键盘消息
	mainForm.SetKeyPreview(true)
	mainForm.SetOnKeyPress(func(sender vcl.IObject, key *types.Char) {
		if *key == '\r' {
			btnGo.Click()
		}
	})

	// 底部的状态条
	statusbar := vcl.NewStatusBar(mainForm)
	statusbar.SetParent(mainForm)
	statusbar.SetSimplePanel(true)

	wkeBrw = wke.NewWkeWebBrowser(pnl.Handle())

	mainForm.SetOnDestroy(func(vcl.IObject) {
		if wkeBrw.IsVaild() {
			wkeBrw.Free()
		}
	})
	if wkeBrw.IsVaild() {
		fmt.Println("wke浏览器创建成功")

		wkeBrw.SetOnTitleChanged(func(title string) {
			mainForm.SetCaption(title + " - Wke测试")
		})

		wkeBrw.SetOnURLChanged(func(url string) {
			edit.SetText(url)
		})

		wkeBrw.SetOnNavigation(func(navigationType wke.NavigationType, url string) {
			fmt.Println("nav:", navigationType, ", url:", url)
			switch navigationType {
			case wke.WKE_NAVIGATION_TYPE_LINKCLICK, wke.WKE_NAVIGATION_TYPE_BACKFORWARD:
				statusbar.SetSimpleText("正在跳转：" + url)

			case wke.WKE_NAVIGATION_TYPE_FORMSUBMITTE:
				statusbar.SetSimpleText("提交表单中：" + url)

			case wke.WKE_NAVIGATION_TYPE_RELOAD:
				statusbar.SetSimpleText("正在重新载入：" + url)

			case wke.WKE_NAVIGATION_TYPE_FORMRESUBMITT:
				statusbar.SetSimpleText("重新提交表单：" + url)

			case wke.WKE_NAVIGATION_TYPE_OTHER:
				// 不知道啥类型了
				statusbar.SetSimpleText("正在载入：" + url)
			}

		})

		wkeBrw.SetOnLoadingFinish(func(url string, result wke.LoadingResult, failedReason string) {
			fmt.Println("LoadingFinish:", result)
			switch result {
			case wke.WKE_LOADING_SUCCEEDED:
				statusbar.SetSimpleText("加载完成.")
			case wke.WKE_LOADING_FAILED:
				statusbar.SetSimpleText("加载失败.")
			case wke.WKE_LOADING_CANCELED:
				statusbar.SetSimpleText("加载缓存.")
			}
		})

		wkeBrw.SetOnDocumentReady(func(info uintptr) {
			statusbar.SetSimpleText("文档已经准备")
		})

		wkeBrw.Show(true)

		// 以前的wke是可以自己处理刷新的，后来作者改了，需要手动刷了，一般是放在消息中，不过我没有提供相关的消息处理的，用个计时器组件也行吧
		refWebbrowserTimer := vcl.NewTimer(mainForm)
		refWebbrowserTimer.SetInterval(80)
		refWebbrowserTimer.SetEnabled(true)
		refWebbrowserTimer.SetOnTimer(func(sender vcl.IObject) {
			wke.RepaintAllNeeded()
		})

		btnGo.Click()
	} else {
		fmt.Println("wke浏览器创建失败。")
	}

	vcl.Application.Run()
}
