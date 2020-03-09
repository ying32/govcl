package main

import (
	"fmt"

	_ "github.com/ying32/govcl/pkgs/winappres"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

func main() {

	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)

	mainForm := vcl.Application.CreateForm()
	mainForm.SetCaption("Hello")
	mainForm.SetPosition(types.PoScreenCenter)
	mainForm.EnabledMaximize(false)
	mainForm.SetWidth(300)
	mainForm.SetHeight(500)

	dlgOpen := vcl.NewOpenDialog(mainForm)
	dlgOpen.SetFilter("文本文件(*.txt)|*.txt|所有文件(*.*)|*.*")
	//    dlgOpen.SetInitialDir()
	//	dlgOpen.SetFilterIndex()

	dlgOpen.SetOptions(dlgOpen.Options().Include(types.OfShowHelp, types.OfAllowMultiSelect)) //rtl.Include(, types.OfShowHelp))
	dlgOpen.SetTitle("打开")

	btn := vcl.NewButton(mainForm)
	btn.SetAlign(types.AlTop)
	btn.SetParent(mainForm)
	btn.SetCaption("Open Dialog")
	btn.SetOnClick(func(vcl.IObject) {
		if dlgOpen.Execute() {
			fmt.Println("filename: ", dlgOpen.FileName())
		}
	})

	dlSave := vcl.NewSaveDialog(mainForm)
	dlSave.SetFilter("文本文件(*.txt)|*.txt|所有文件(*.*)|*.*")
	dlSave.SetOptions(dlSave.Options().Include(types.OfShowHelp))
	dlSave.SetTitle("保存")

	btn = vcl.NewButton(mainForm)
	btn.SetAlign(types.AlTop)
	btn.SetParent(mainForm)
	btn.SetCaption("Save Dialog")
	btn.SetOnClick(func(vcl.IObject) {
		if dlSave.Execute() {
			fmt.Println("filename: ", dlSave.FileName())
		}
	})

	dlFont := vcl.NewFontDialog(mainForm)

	btn = vcl.NewButton(mainForm)
	btn.SetAlign(types.AlTop)
	btn.SetParent(mainForm)
	btn.SetCaption("Font Dialog")
	btn.SetOnClick(func(vcl.IObject) {
		if dlFont.Execute() {
			fmt.Println("Name: ", dlFont.Font().Name())
		}
	})

	dlColor := vcl.NewColorDialog(mainForm)
	btn = vcl.NewButton(mainForm)
	btn.SetAlign(types.AlTop)
	btn.SetParent(mainForm)
	btn.SetCaption("Color Dialog")
	btn.SetOnClick(func(vcl.IObject) {
		if dlColor.Execute() {
			fmt.Println("Color: ", dlColor.Color())
		}
	})

	dlPicOpen := vcl.NewOpenPictureDialog(mainForm)
	btn = vcl.NewButton(mainForm)
	btn.SetAlign(types.AlTop)
	btn.SetParent(mainForm)
	btn.SetCaption("OpenPic Dialog")
	btn.SetOnClick(func(vcl.IObject) {
		if dlPicOpen.Execute() {
			fmt.Println("Name: ", dlPicOpen.FileName())
		}
	})

	dlPicSave := vcl.NewSavePictureDialog(mainForm)
	btn = vcl.NewButton(mainForm)
	btn.SetAlign(types.AlTop)
	btn.SetParent(mainForm)
	btn.SetCaption("SavePic Dialog")
	btn.SetOnClick(func(vcl.IObject) {
		if dlPicSave.Execute() {
			fmt.Println("Name: ", dlPicSave.FileName())
		}
	})

	dlTxtOpen := vcl.NewOpenTextFileDialog(mainForm)
	btn = vcl.NewButton(mainForm)
	btn.SetAlign(types.AlTop)
	btn.SetParent(mainForm)
	btn.SetCaption("Open Text Dialog")
	btn.SetOnClick(func(vcl.IObject) {
		if dlTxtOpen.Execute() {
			fmt.Println("Name: ", dlTxtOpen.FileName())
		}
	})

	dlTxtSave := vcl.NewSaveTextFileDialog(mainForm)
	btn = vcl.NewButton(mainForm)
	btn.SetAlign(types.AlTop)
	btn.SetParent(mainForm)
	btn.SetCaption("Save Text Dialog")
	btn.SetOnClick(func(vcl.IObject) {
		if dlTxtSave.Execute() {
			fmt.Println("Name: ", dlTxtSave.FileName())
		}
	})

	btn = vcl.NewButton(mainForm)
	btn.SetAlign(types.AlTop)
	btn.SetParent(mainForm)
	btn.SetCaption("SelectDirectory1")
	btn.SetOnClick(func(vcl.IObject) {
		if ok, dir := vcl.SelectDirectory1(0); ok {
			fmt.Println("选择的目录为：", dir)
		}
	})

	btn = vcl.NewButton(mainForm)
	btn.SetAlign(types.AlTop)
	btn.SetParent(mainForm)
	btn.SetCaption("SelectDirectory2")
	btn.SetOnClick(func(vcl.IObject) {
		options := types.NewSet(types.SdNewFolder, types.SdShowEdit, types.SdNewUI)
		if ok, dir := vcl.SelectDirectory2("标题了", "C:/", options, nil); ok {
			fmt.Println("选择的目录为：", dir)
		}
	})

	btn = vcl.NewButton(mainForm)
	btn.SetAlign(types.AlTop)
	btn.SetParent(mainForm)
	btn.SetCaption("FindDialog")
	findDialog := vcl.NewFindDialog(mainForm)
	findDialog.SetOnFind(func(sender vcl.IObject) {
		fmt.Println("FindText: ", findDialog.FindText())
		opt := findDialog.Options()
		/*
			FrDown = iota + 0
			FrFindNext
			FrHideMatchCase
			FrHideWholeWord
			FrHideUpDown
			FrMatchCase
			FrDisableMatchCase
			FrDisableUpDown
			FrDisableWholeWord
			FrReplace
			FrReplaceAll
			FrWholeWord
			FrShowHelp
		*/
		if opt.In(types.FrDown) {
			fmt.Println("向下")
		} else {
			fmt.Println("向上")
		}
		if opt.In(types.FrFindNext) {
			fmt.Println("查找下一个")
		}
		if opt.In(types.FrMatchCase) {
			fmt.Println("区分大小写")
		}
	})
	btn.SetOnClick(func(vcl.IObject) {
		findDialog.Execute()
	})

	btn = vcl.NewButton(mainForm)
	btn.SetAlign(types.AlTop)
	btn.SetParent(mainForm)
	btn.SetCaption("ReplaceDialog")
	replaceDialog := vcl.NewReplaceDialog(mainForm)
	replaceDialog.SetOnFind(func(sender vcl.IObject) {
		fmt.Println("FindText:", replaceDialog.FindText(), ", Relpace: ", replaceDialog.ReplaceText())
		opt := replaceDialog.Options()
		/*
			FrDown = iota + 0
			FrFindNext
			FrHideMatchCase
			FrHideWholeWord
			FrHideUpDown
			FrMatchCase
			FrDisableMatchCase
			FrDisableUpDown
			FrDisableWholeWord
			FrReplace
			FrReplaceAll
			FrWholeWord
			FrShowHelp
		*/
		if opt.In(types.FrDown) {
			fmt.Println("向下")
		} else {
			fmt.Println("向上")
		}
		if opt.In(types.FrFindNext) {
			fmt.Println("查找下一个")
		}
		if opt.In(types.FrMatchCase) {
			fmt.Println("区分大小写")
		}
	})

	replaceDialog.SetOnReplace(func(sender vcl.IObject) {
		opt := replaceDialog.Options()
		if opt.In(types.FrReplaceAll) {
			fmt.Println("替换全部")
		}
		if opt.In(types.FrReplace) {
			fmt.Println("替换一次")
		}
		fmt.Println("替换字符：", replaceDialog.ReplaceText())
	})

	btn.SetOnClick(func(vcl.IObject) {
		replaceDialog.Execute()
	})

	btn = vcl.NewButton(mainForm)
	btn.SetAlign(types.AlTop)
	btn.SetParent(mainForm)
	btn.SetCaption("InputBox")
	btn.SetOnClick(func(vcl.IObject) {
		s := vcl.InputBox("标题", "提示", "默认值")
		if s != "" {
			fmt.Println("结果：", s)
		}
	})

	btn = vcl.NewButton(mainForm)
	btn.SetAlign(types.AlTop)
	btn.SetParent(mainForm)
	btn.SetCaption("InpuQuery")
	btn.SetOnClick(func(vcl.IObject) {
		s := "default"
		if vcl.InputQuery("标题", "提示", &s) {
			fmt.Println("结果：", s)
		}
	})

	dlPrinterSetupDialog := vcl.NewPrinterSetupDialog(mainForm)
	btn = vcl.NewButton(mainForm)
	btn.SetAlign(types.AlTop)
	btn.SetParent(mainForm)
	btn.SetCaption("PrinterSetupDialog")
	btn.SetOnClick(func(vcl.IObject) {
		dlPrinterSetupDialog.Execute()
	})

	dlPageSetupDialog := vcl.NewPageSetupDialog(mainForm)
	btn = vcl.NewButton(mainForm)
	btn.SetAlign(types.AlTop)
	btn.SetParent(mainForm)
	btn.SetCaption("PageSetupDialog")
	btn.SetOnClick(func(vcl.IObject) {
		dlPageSetupDialog.Execute()
	})

	vcl.Application.Run()
}
