package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	//	"syscall"

	"gitee.com/ying32/govcl/vcl"
	"gitee.com/ying32/govcl/vcl/rtl"
	"gitee.com/ying32/govcl/vcl/types"
)

//var (
//	kernel32            = syscall.NewLazyDLL("kernel32.dll")
//	_GetCurrentThreadId = kernel32.NewProc("GetCurrentThreadId")
//)

func GetCurrentThreadId() uintptr {
	//r, _, _ := _GetCurrentThreadId.Call()
	//return r
	return 0
}

func main() {

	fmt.Println("main:currentThreadId:", GetCurrentThreadId())
	icon := vcl.NewIcon()
	defer icon.Free()
	icon.LoadFromResourceID(rtl.MainInstance(), 3)

	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.SetIcon(icon)

	mainForm := vcl.Application.CreateForm()
	mainForm.SetCaption("Hello")
	mainForm.SetPosition(types.PoScreenCenter)
	mainForm.EnabledMaximize(false)
	mainForm.SetWidth(500)
	mainForm.SetHeight(600)

	img := vcl.NewImage(mainForm)
	img.SetParent(mainForm)
	// 本地加载
	mem := vcl.NewMemoryStream()
	defer mem.Free()
	mem.LoadFromFile("..\\..\\imgs\\1.jpg")
	img.Picture().LoadFromStream(mem)
	// 网络图片加载
	img2 := vcl.NewImage(mainForm)
	img2.SetParent(mainForm)
	img2.SetTop(img.Height() + 10)
	img2.SetAutoSize(true)

	// 异步加载，一般来说不要在非主线程中访问UI组件,需要在线程中访问ui组件请使用 vcl.ThreadSync
	go func() {
		fmt.Println("main:currentThreadId2:", GetCurrentThreadId())
		resp, err := http.Get("http://ww2.sinaimg.cn/large/df780e95jw1egxm06uxerj20cs05hjs8.jpg")
		if err == nil {
			defer resp.Body.Close()
			bs, err := ioutil.ReadAll(resp.Body)
			if err == nil {
				mem := vcl.NewMemoryStream()
				defer mem.Free()
				mem.Write(bs)
				mem.SetPosition(0)
				// 让以下代码运行在主线程中
				vcl.ThreadSync(func() {
					fmt.Println("main:currentThreadId3:", GetCurrentThreadId())
					img2.Picture().LoadFromStream(mem)
				})
				fmt.Println("测试运行到此。")
			} else {
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
		}
	}()
	vcl.Application.Run()
}
