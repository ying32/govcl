package main

import (
	"math/rand"

	_ "github.com/ying32/govcl/pkgs/winappres"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/colors"
)

var (
	mainForm *vcl.TForm
	richEdit *vcl.TRichEdit
)

func main() {

	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)

	initMainForm()
	initMainMenu()
	tlbar := vcl.NewToolBar(mainForm)
	tlbar.SetParent(mainForm)
	tlbar.SetShowCaptions(true)
	btn1 := vcl.NewToolButton(mainForm)
	btn1.SetParent(tlbar)
	btn1.SetCaption("吉")
	btn1.SetOnClick(func(sender vcl.IObject) {
		//richEdit.SetSelStart(-1)
		//vcl.ShowMessage(strconv.Itoa(int(richEdit.SelStart())))

		//start := richEdit.SelStart()
		//
		//richEdit.SetSelStart(start)
		richEdit.SetSelLength(1)
		x := rand.Int31n(3)
		switch x {
		case 0:
			richEdit.SelAttributes().SetColor(colors.ClRed)
		case 1:
			richEdit.SelAttributes().SetColor(colors.ClGreen)
		case 2:
			richEdit.SelAttributes().SetColor(colors.ClBlue)
		default:
			richEdit.SelAttributes().SetColor(colors.ClYellow)
		}
		richEdit.Lines().Add("数据1")
		//richEdit.SetSelLength(-1)
		//println(richEdit.SelStart())
		//vcl.ShowMessage(strconv.Itoa(int(richEdit.SelStart())))
	})

	richEdit = vcl.NewRichEdit(mainForm)
	richEdit.SetParent(mainForm)
	richEdit.SetAlign(types.AlClient)
	richEdit.Lines().Add("这是一段文字红色，粗体，斜體")
	richEdit.SetSelStart(6)
	richEdit.SetSelLength(2)
	richEdit.SelAttributes().SetColor(colors.ClRed)

	richEdit.SetSelStart(9)
	richEdit.SetSelLength(2)

	richEdit.SelAttributes().SetStyle(types.NewSet(types.FsBold))

	richEdit.SetSelStart(12)
	richEdit.SetSelLength(2)

	richEdit.SelAttributes().SetStyle(types.NewSet(types.FsItalic))

	richEdit.SetSelStart(15)

	initRichEditPopupMenu()

	stabar := vcl.NewStatusBar(mainForm)
	stabar.SetParent(mainForm)

	vcl.Application.Run()
}

func initMainForm() {
	mainForm = vcl.Application.CreateForm()
	mainForm.SetCaption("Hello")
	mainForm.SetPosition(types.PoScreenCenter)
	mainForm.EnabledMaximize(false)
	mainForm.SetWidth(600)
	mainForm.SetHeight(400)
}

func initMainMenu() {
	mainMenu := vcl.NewMainMenu(mainForm)

	item := vcl.NewMenuItem(mainForm)
	item.SetCaption("&File")
	mainMenu.Items().Add(item)

	item = vcl.NewMenuItem(mainForm)
	item.SetCaption("&Help")
	mainMenu.Items().Add(item)
}

func initRichEditPopupMenu() {
	pm := vcl.NewPopupMenu(mainForm)
	item := vcl.NewMenuItem(mainForm)
	item.SetCaption("&Clear")
	pm.Items().Add(item)

	richEdit.SetPopupMenu(pm)
}
