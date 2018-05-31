// govcl project main.go
// go.exe build -i -ldflags="-H windowsgui"
package main

import (
	"fmt"

	"gitee.com/ying32/govcl/vcl"
	"gitee.com/ying32/govcl/vcl/exts/tools"
	"gitee.com/ying32/govcl/vcl/rtl"
	"gitee.com/ying32/govcl/vcl/rtl/version"
	"gitee.com/ying32/govcl/vcl/types"
	"gitee.com/ying32/govcl/vcl/win"
)

var (
	mainForm *vcl.TForm
	trayicon *vcl.TTrayIcon
)

func main() {

	// mac下记得发布时去掉
	tools.RunWithMacOSApp()
	// 异常捕获
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Exception: ", err)
			vcl.ShowMessage(err.(error).Error())
		}
	}()

	fmt.Println("main")
	vcl.Application.SetIconResId(3)
	vcl.Application.Initialize()

	vcl.Application.SetOnException(func(sender vcl.IObject, e *vcl.Exception) {
		fmt.Println("exception.")
	})

	vcl.Application.SetTitle("Hello World! 系统信息：" + version.OSVersion.ToString())
	vcl.Application.SetMainFormOnTaskBar(true)
	// 窗口自动根据系统绽放，默认为true
	//vcl.Application.SetFormScaled(false)

	mainForm = vcl.Application.CreateForm()
	mainForm.SetWidth(800)
	mainForm.SetHeight(600)
	mainForm.SetOnClose(func(Sender vcl.IObject, Action *types.TCloseAction) {
		fmt.Println("close")
	})

	// 窗口大小约束
	mainForm.SetOnConstrainedResize(func(sender vcl.IObject, minWidth, minHeight, maxWidth, maxHeight *int32) {
		*minWidth = 800
		*minHeight = 600
		*maxWidth = 800
		*maxHeight = 600
	})

	mainForm.SetOnDestroy(func(sender vcl.IObject) {
		fmt.Println("Form Destroy.")
	})

	fmt.Println("MainForm ClientRect: ", mainForm.ClientRect())
	filename := vcl.Application.ExeName()
	fmt.Println("application.ExeName: ", filename)
	fmt.Println("path: ", rtl.ExtractFilePath(filename))
	fmt.Println("fileExists: ", rtl.FileExists(filename))

	mainForm.SetOnCloseQuery(func(Sender vcl.IObject, CanClose *bool) {
		*CanClose = vcl.MessageDlg("是否退出?", types.MtInformation, types.MbYes, types.MbNo) == types.MrYes
		fmt.Println("OnCloseQuery")
	})

	mainForm.SetCaption(vcl.Application.Title())
	mainForm.EnabledMaximize(false)
	mainForm.SetDoubleBuffered(true)
	//mainForm.SetPosition(types.PoScreenCenter)
	//mainForm.ScreenCenter()
	mainForm.WorkAreaCenter()
	mainForm.SetKeyPreview(true)
	mainForm.SetOnKeyDown(func(Sender vcl.IObject, Key *types.Char, Shift types.TShiftState) {
		fmt.Println(rtl.InSets(uint32(Shift), types.SsCtrl))
		fmt.Println("key:", *Key)
	})

	mainForm.SetOnMouseDown(func(sender vcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
		fmt.Println("Button:", button == types.MbLeft, ", X:", x, ", y:", y)
		fmt.Println("OnMouseDown")
	})

	chk := vcl.NewCheckBox(mainForm)
	chk.SetParent(mainForm)
	chk.SetChecked(true)
	chk.SetCaption("测试")
	chk.SetLeft(1)
	chk.SetTop(60)
	chk.SetOnClick(func(vcl.IObject) {
		fmt.Println("chk.Checked=", chk.Checked())
	})

	// action
	action := vcl.NewAction(mainForm)
	action.SetCaption("action1")
	action.SetOnUpdate(func(sender vcl.IObject) {
		vcl.ActionFromObj(sender).SetEnabled(chk.Checked())
	})
	action.SetOnExecute(func(vcl.IObject) {
		fmt.Println("action execute")
	})
	btn := vcl.NewButton(mainForm)
	btn.SetParent(mainForm)
	btn.SetBounds(250, 30, 90, 25)
	btn.SetCaption("action")
	btn.SetAction(action)

	trayicon = vcl.NewTrayIcon(mainForm)
	if rtl.LcLLoaded() {
		trayicon.SetIcon(vcl.Application.Icon()) //不设置会自动使用Application.Icon
	}

	trayicon.SetHint(mainForm.Caption())
	trayicon.SetVisible(true)
	trayicon.SetOnClick(func(vcl.IObject) {
		trayicon.SetBalloonTitle("test")
		trayicon.SetBalloonTimeout(10000)
		trayicon.SetBalloonHint("我是提示正文啦")
		trayicon.ShowBalloonHint()
		fmt.Println("TrayIcon Click.")
	})

	// img
	img := vcl.NewImage(mainForm)
	img.SetBounds(132, 30, 156, 97)
	img.SetParent(mainForm)
	img.Picture().LoadFromFile("./imgs/1.jpg")
	//img.SetStretch(true)
	img.SetProportional(true)

	// linklabel
	linklbl := vcl.NewLinkLabel(mainForm)
	linklbl.SetAlign(types.AlBottom)
	linklbl.SetCaption("<a href=\"https://github.com/ying32/govcl\">govcl测试链接</a>")
	linklbl.SetParent(mainForm)
	linklbl.SetOnLinkClick(func(sender vcl.IObject, link string, linktype types.TSysLinkType) {
		fmt.Println("link label: ", link, ", type: ", linktype)
		rtl.SysOpen(link)
	})

	// menu
	mainMenu := vcl.NewMainMenu(mainForm)
	item := vcl.NewMenuItem(mainForm)
	item.SetCaption("File(&F)")
	mainMenu.Items().Add(item)

	item2 := vcl.NewMenuItem(mainForm)
	item2.SetCaption("MemoryStreamTest")
	item2.SetOnClick(func(vcl.IObject) {
		mem := vcl.NewMemoryStream()
		defer mem.Free()
		mem.Write([]byte("测试"))
		mem.SaveToFile("test.txt")

		mem.SetPosition(0)
		n, bs := mem.Read(int32(mem.Size()))
		fmt.Println("n:", n, ", bs:", bs, ", str:", string(bs))
	})
	item.Add(item2)

	item2 = vcl.NewMenuItem(mainForm)
	item2.SetCaption("Exit(&E)")
	item2.SetShortCutFromString("Ctrl+Q")
	item2.SetOnClick(func(vcl.IObject) {
		mainForm.Close()
	})
	item.Add(item2)

	//	mainForm.EnabledMinimize(false)
	//	mainForm.EnabledSystemMenu(false)

	button := vcl.NewButton(mainForm)

	button.SetCaption("消息")
	button.SetParent(mainForm)
	button.SetOnClick(func(vcl.IObject) {
		fmt.Println("button click")
		vcl.ShowMessage("这是一个消息")
		vcl.Application.MessageBox("Hello!", "Message", win.MB_YESNO+win.MB_ICONINFORMATION)
	})
	button.SetLeft(50)
	button.SetTop(50)
	button.SetAlign(types.AlRight)

	edit := vcl.NewEdit(mainForm)
	edit.SetParent(mainForm)
	edit.SetLeft(1)
	edit.SetTop(30)
	edit.SetTextHint("测试")
	edit.SetOnChange(func(vcl.IObject) {
		fmt.Println("edit OnChange")
	})

	button2 := vcl.NewButton(mainForm)
	button2.SetParent(mainForm)
	button2.SetCaption("a")
	button2.SetWidth(100)
	button2.SetHeight(28)
	button2.SetOnClick(func(vcl.IObject) {
		fmt.Println("button2 click")

		edit.SetText("Hello!")
		fmt.Println("ScreenWidth:", vcl.Screen.Width(), ", ScreenHeight:", vcl.Screen.Height())
	})
	button2.SetAlign(types.AlTop)

	combo := vcl.NewComboBox(mainForm)
	combo.SetAlign(types.AlBottom)
	combo.SetParent(mainForm)
	combo.SetText("ffff")
	combo.Items().Add("1")
	combo.Items().Add("2")
	combo.SetItemIndex(0)
	combo.SetOnChange(func(vcl.IObject) {
		if combo.ItemIndex() != -1 {
			fmt.Println("combo Change: ", combo.Items().Strings(combo.ItemIndex()))
		}

	})

	page := vcl.NewPageControl(mainForm)
	page.SetParent(mainForm)
	page.SetAlign(types.AlBottom)
	sheet := vcl.NewTabSheet(mainForm)
	sheet.SetPageControl(page)
	sheet.SetCaption("第一页")

	// 需要先将TabSheet设置了父窗口，TListView才可用，不然就会报错
	lv1 := vcl.NewListView(mainForm)
	lv1.SetAlign(types.AlClient)
	lv1.SetParent(sheet)

	lv1.SetViewStyle(types.VsReport)
	lv1.SetRowSelect(true)
	lv1.SetReadOnly(true)
	lv1.SetGridLines(true)
	col := lv1.Columns().Add()
	col.SetCaption("序号")
	col.SetWidth(100)
	// 强制柱头宽，即使被调整也会被还原
	col.SetMaxWidth(100)
	col.SetMinWidth(100)
	col = lv1.Columns().Add()
	col.SetCaption("名称")
	col.SetWidth(200)
	col = lv1.Columns().Add()
	col.SetCaption("内容")
	col.SetWidth(200)
	lv1.SetOnClick(func(vcl.IObject) {
		if lv1.ItemIndex() != -1 {
			item := lv1.Selected() // lv1.Items().Item(lv1.ItemIndex())
			fmt.Println(item.Caption(),
				item.SubItems().Strings(0),
				item.SubItems().Strings(1))
		}
	})

	lv1.Items().BeginUpdate()
	for i := 1; i <= 50; i++ {
		lstitem := lv1.Items().Add()
		lstitem.SetCaption(fmt.Sprintf("%d", i))
		lstitem.SubItems().Add(fmt.Sprintf("第%d", i))
		lstitem.SubItems().Add(fmt.Sprintf("内容%d", i))
	}
	lv1.Items().EndUpdate()

	sheet = vcl.NewTabSheet(mainForm)
	sheet.SetCaption("第二页")
	sheet.SetPageControl(page)

	// -----------TreeView 不同Node弹出不同菜单，两个右键例程不同使用

	tvpm1 := vcl.NewPopupMenu(mainForm)
	mItem := vcl.NewMenuItem(mainForm)
	mItem.SetCaption("第一种")
	tvpm1.Items().Add(mItem)

	tvpm2 := vcl.NewPopupMenu(mainForm)
	mItem = vcl.NewMenuItem(mainForm)
	mItem.SetCaption("第二种")
	tvpm2.Items().Add(mItem)

	tv1 := vcl.NewTreeView(mainForm)
	tv1.SetAutoExpand(true)
	tv1.SetParent(sheet)
	tv1.SetAlign(types.AlClient)
	//	tv1.SetRightClickSelect(true)
	tv1.SetOnClick(func(vcl.IObject) {
		if tv1.SelectionCount() > 0 {
			node := tv1.Selected()
			fmt.Println("text:", node.Text(), ", index:", node.Index())
		}
	})

	tv1.SetOnMouseDown(func(sender vcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
		if button == types.MbRight {
			node := tv1.GetNodeAt(x, y)
			if node != nil && node.IsValid() {
				// 自由决择是否选中
				node.SetSelected(true)
				// 根据Level来判断，这里只是做演示
				p := vcl.Mouse.CursorPos()
				switch node.Level() {
				case 0:
					tvpm1.Popup(p.X, p.Y)
				case 1:
					tvpm2.Popup(p.X, p.Y)
				}
				fmt.Println("node.Level():", node.Level(), ", text:", node.Text())
			}
		}
	})

	tv1.Items().BeginUpdate()
	node := tv1.Items().AddChild(nil, "首个")
	for i := 1; i <= 50; i++ {
		tv1.Items().AddChild(node, fmt.Sprintf("Node%d", i))
	}
	node = tv1.Items().AddChild(nil, "第二个")
	for i := 1; i <= 50; i++ {
		tv1.Items().AddChild(node, fmt.Sprintf("Node%d", i))
	}
	tv1.Items().EndUpdate()

	fmt.Println("Compoment Count:", mainForm.ComponentCount())
	//	mainForm.ScreenCenter()

	lbl := vcl.NewLabel(mainForm)
	lbl.SetCaption("标签")
	lbl.SetAlign(types.AlBottom)
	fmt.Println("InheritsFromControl:", rtl.InheritsFromControl(mainForm.Instance()))
	fmt.Println("InheritsFromWinControl:", rtl.InheritsFromWinControl(mainForm.Instance()))
	fmt.Println("InheritsFromComponent:", rtl.InheritsFromComponent(tv1.Instance()))
	fmt.Println("InheritsFromWinControl:", rtl.InheritsFromWinControl(lbl.Instance()))

	vcl.Application.Run()
}
