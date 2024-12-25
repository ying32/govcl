package main

import (
	_ "github.com/ying32/govcl/pkgs/winappres"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

type TMainForm struct {
	*vcl.TForm
	NoteBook1        *vcl.TNotebook
	BtnPrev, BtnNext *vcl.TButton
	Pnl1             *vcl.TPanel
}

var (
	mainForm *TMainForm
)

func main() {
	vcl.DEBUG = true
	vcl.RunApp(&mainForm)
}

func (f *TMainForm) OnFormCreate(sender vcl.IObject) {
	f.SetCaption("Hello")
	f.EnabledMaximize(false)
	f.SetWidth(600)
	f.SetHeight(400)
	f.ScreenCenter()

	f.NoteBook1 = vcl.NewNotebook(f)
	f.NoteBook1.SetParent(f)
	f.NoteBook1.SetAlign(types.AlClient)
	f.newPage("Page1")
	f.newPage("Page2")
	f.newPage("Page3")

	f.Pnl1 = vcl.NewPanel(f)
	f.Pnl1.SetParent(f)
	f.Pnl1.SetAlign(types.AlBottom)
	f.Pnl1.SetHeight(80)

	f.BtnPrev = vcl.NewButton(f)
	f.BtnPrev.SetParent(f.Pnl1)
	f.BtnPrev.SetAlign(types.AlLeft)
	f.BtnPrev.SetCaption("上一页")
	f.BtnPrev.SetOnClick(f.onBtnPrevClick)

	f.BtnNext = vcl.NewButton(f)
	f.BtnNext.SetParent(f.Pnl1)
	f.BtnNext.SetAlign(types.AlRight)
	f.BtnNext.SetCaption("下一页")
	f.BtnNext.SetOnClick(f.onBtnNextClick)
}

func (f *TMainForm) newPage(title string) {
	idx := f.NoteBook1.Pages().Add(title)
	page := f.NoteBook1.Page(idx)
	//page := f.NoteBook1.Page(f.NoteBook1.PageCount() - 1)
	btn := vcl.NewButton(f)
	btn.SetParent(page)
	btn.SetCaption(title)
}

func (f *TMainForm) onBtnPrevClick(sender vcl.IObject) {
	if f.NoteBook1.PageIndex()-1 >= 0 {
		f.NoteBook1.SetPageIndex(f.NoteBook1.PageIndex() - 1)
	}
}

func (f *TMainForm) onBtnNextClick(sender vcl.IObject) {
	if f.NoteBook1.PageIndex()+1 < f.NoteBook1.PageCount() {
		f.NoteBook1.SetPageIndex(f.NoteBook1.PageIndex() + 1)
	}
}
