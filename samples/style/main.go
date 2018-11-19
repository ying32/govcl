// +build windows

// +build windows

package main

import (
	"fmt"
	"time"

	"io"
	"os"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/rtl"
	"github.com/ying32/govcl/vcl/types"
)

var styleNames = make(map[string]string, 0)

func main() {

	if rtl.LcLLoaded() {
		vcl.ShowMessage("样式不支持liblcl。")
		return
	}

	vcl.Application.SetIconResId(3)
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)

	vcl.StyleManager.SetStyle(vcl.StyleManager.LoadFromFile("..\\..\\bin\\styles\\TabletLight.vsf"))

	mainForm := vcl.Application.CreateForm()
	mainForm.SetCaption("Hello")
	mainForm.SetPosition(types.PoScreenCenter)
	mainForm.SetWidth(500)
	mainForm.SetHeight(700)

	var top int32 = 40
	// TButton
	btn := vcl.NewButton(mainForm)
	btn.SetParent(mainForm)
	btn.SetLeft(10)
	btn.SetTop(top)
	btn.SetWidth(150)
	btn.SetCaption("打印已注册的样式")
	btn.SetOnClick(func(vcl.IObject) {
		fmt.Println("按钮1单击")
		styleNames := vcl.StyleManager.StyleNames()
		fmt.Println("len:", len(styleNames))
		for _, s := range styleNames {
			fmt.Println("已注册样式：", s)
		}
	})

	top += btn.Height() + 5

	// TEdit
	edit := vcl.NewEdit(mainForm)
	edit.SetParent(mainForm)
	edit.SetLeft(10)
	edit.SetTop(top)
	edit.SetTextHint("提示")
	//edit.SetText("文字")
	//	edit.SetReadOnly(true)
	edit.SetOnChange(func(vcl.IObject) {
		fmt.Println("文字改变了")
	})

	top += edit.Height() + 5
	// TEdit Password
	edit = vcl.NewEdit(mainForm)
	edit.SetParent(mainForm)
	edit.SetLeft(10)
	edit.SetTop(top)
	edit.SetText("文字")
	edit.SetPasswordChar('*')

	top += edit.Height() + 5

	// TLabel
	lbl := vcl.NewLabel(mainForm)
	lbl.SetParent(mainForm)
	lbl.SetLeft(10)
	lbl.SetTop(top)
	lbl.SetCaption("标签1")
	// 解除样式对Label Font的Hook
	lbl.SetStyleElements(rtl.Include(0, types.SeClient, types.SeBorder))
	lbl.Font().SetColor(255)

	top += lbl.Height() + 5

	// TCheckBox
	chk := vcl.NewCheckBox(mainForm)
	chk.SetParent(mainForm)
	chk.SetLeft(10)
	chk.SetTop(top)
	chk.SetCaption("选择框1")
	chk.SetOnClick(func(vcl.IObject) {
		fmt.Println("checked: ", chk.Checked())
	})

	// TStatusBar
	stat := vcl.NewStatusBar(mainForm)
	stat.SetParent(mainForm)
	//stat.SetSizeGrip(false) // 右解是否有可调的
	spnl := stat.Panels().Add()
	spnl.SetText("第一个")
	spnl.SetWidth(80)

	spnl = stat.Panels().Add()
	spnl.SetText("第二个")
	spnl.SetWidth(80)

	// TToolBar
	tlbar := vcl.NewToolBar(mainForm)
	tlbar.SetParent(mainForm)
	tlbar.SetShowCaptions(true)

	// 倒过来创建
	tlbtn := vcl.NewToolButton(mainForm)
	tlbtn.SetParent(tlbar)
	tlbtn.SetCaption("2")
	tlbtn.SetStyle(types.TbsDropDown)

	tlbtn = vcl.NewToolButton(mainForm)
	tlbtn.SetParent(tlbar)
	tlbtn.SetStyle(types.TbsSeparator)

	tlbtn = vcl.NewToolButton(mainForm)
	tlbtn.SetParent(tlbar)
	tlbtn.SetCaption("1")

	top += chk.Height() + 5
	// TRadioButton
	rd := vcl.NewRadioButton(mainForm)
	rd.SetParent(mainForm)
	rd.SetLeft(10)
	rd.SetTop(top)
	rd.SetCaption("选项1")

	var left int32 = rd.Left() + rd.Width() + 5

	rd = vcl.NewRadioButton(mainForm)
	rd.SetParent(mainForm)
	rd.SetLeft(left)
	rd.SetTop(top)
	rd.SetCaption("选项2")

	top += rd.Height() + 5
	// TMemo
	mmo := vcl.NewMemo(mainForm)
	mmo.SetParent(mainForm)
	mmo.SetBounds(10, top, 167, 50)
	//    mmo.Text()
	mmo.Lines().Add("1")
	mmo.Lines().Add("2")

	top += mmo.Height() + 5
	// TComboBox
	cb := vcl.NewComboBox(mainForm)
	cb.SetParent(mainForm)
	cb.SetLeft(10)
	cb.SetTop(top)
	cb.SetStyle(types.CsDropDownList)
	cb.Items().Add("1")
	cb.Items().Add("2")
	cb.Items().Add("3")
	cb.SetItemIndex(0)
	cb.SetOnChange(func(vcl.IObject) {
		if cb.ItemIndex() != -1 {
			fmt.Println(cb.Items().Strings(cb.ItemIndex()))
		}
	})

	// TListBox
	top += cb.Height() + 5
	lst := vcl.NewListBox(mainForm)
	lst.SetParent(mainForm)
	lst.SetBounds(10, top, 167, 50)
	lst.Items().Add("1")
	lst.Items().Add("2")
	lst.Items().Add("3")

	// TPanel
	top += lst.Height() + 5
	pnl := vcl.NewPanel(mainForm)
	pnl.SetParent(mainForm)
	pnl.SetCaption("fff")
	//    pnl.SetShowCaption(false)
	pnl.SetBounds(10, top, 167, 50)

	// color
	top += pnl.Height() + 5
	clr := vcl.NewColorBox(mainForm)
	clr.SetParent(mainForm)
	clr.SetLeft(10)
	clr.SetTop(top)
	clr.SetOnChange(func(vcl.IObject) {
		if clr.ItemIndex() != -1 {
			lbl.Font().SetColor(clr.Selected())
		}
	})

	// TPageControl
	top += clr.Height() + 5
	pgc := vcl.NewPageControl(mainForm)
	pgc.SetParent(mainForm)
	pgc.SetBounds(10, top, 167, 100)
	pgc.SetOnChange(func(vcl.IObject) {
		fmt.Println("当前索引:", pgc.ActivePageIndex())
	})

	sheet := vcl.NewTabSheet(mainForm)
	sheet.SetPageControl(pgc)
	sheet.SetCaption("一")
	btn = vcl.NewButton(mainForm)
	btn.SetParent(sheet)
	btn.SetLeft(10)
	btn.SetTop(10)
	btn.SetCaption("按钮1")

	sheet = vcl.NewTabSheet(mainForm)
	sheet.SetPageControl(pgc)
	sheet.SetCaption("二")
	btn = vcl.NewButton(mainForm)
	btn.SetParent(sheet)
	btn.SetLeft(10)
	btn.SetTop(10)
	btn.SetCaption("按钮2")

	sheet = vcl.NewTabSheet(mainForm)
	sheet.SetPageControl(pgc)
	sheet.SetCaption("三")
	btn = vcl.NewButton(mainForm)
	btn.SetParent(sheet)
	btn.SetLeft(10)
	btn.SetTop(10)
	btn.SetCaption("按钮3")

	// TImage
	top += pgc.Height() + 5
	img := vcl.NewImage(mainForm)
	img.SetBounds(10, top, 167, 97)
	img.SetParent(mainForm)
	img.Picture().LoadFromFile("1.jpg")
	//img.SetStretch(true)
	img.SetProportional(true)

	left = 210
	top = 10
	// TTrackBar
	trkbar := vcl.NewTrackBar(mainForm)
	trkbar.SetParent(mainForm)
	trkbar.SetBounds(left, top, 167, 20)
	trkbar.SetMax(100)
	trkbar.SetMin(0)
	trkbar.SetPosition(50)

	// TProgressBar
	top += trkbar.Height() + 10
	prgbar := vcl.NewProgressBar(mainForm)
	prgbar.SetParent(mainForm)
	prgbar.SetBounds(left, top, 10, 167)
	prgbar.SetMax(100)
	prgbar.SetMin(0)
	prgbar.SetPosition(1)
	prgbar.SetOrientation(types.PbVertical)

	trkbar.SetOnChange(func(vcl.IObject) {
		prgbar.SetPosition(trkbar.Position())
	})

	stylelist := vcl.NewListBox(mainForm)
	stylelist.SetParent(mainForm)
	stylelist.SetLeft(prgbar.Left() + prgbar.Width() + 10)
	stylelist.SetTop(prgbar.Top())
	stylelist.SetHeight(prgbar.Height())
	stylelist.SetWidth(240)
	stylelist.SetOnDblClick(func(vcl.IObject) {
		index := stylelist.ItemIndex()
		if index != -1 {
			// 这里直接替换是因为原本文件名就是样式名，只是简单下，实际样式名要通过相关函数取得，但这里没有给出相关函数
			//styleName := strings.Replace(stylelist.Items().Strings(stylelist.ItemIndex()), ".vsf", "", 1)
			//styleHandle := vcl.StyleManager.Style(styleName)
			text := stylelist.Items().Strings(index)
			if name, ok := styleNames[text]; ok {
				vcl.StyleManager.SetStyle2(name)
				return
			}
			styleFileName := "..\\..\\bin\\styles\\" + text
			if rtl.FileExists(styleFileName) {
				if ok, name := vcl.StyleManager.IsValidStyle2(styleFileName); ok {
					styleNames[text] = name
					vcl.StyleManager.SetStyleFromFileName(styleFileName)
				} else {
					fmt.Println("样式无效")
				}
			} else {
				fmt.Println("样式文件不存在")
			}
		}
	})
	addStyleFileName(stylelist)

	top += prgbar.Height() + 10

	dtp := vcl.NewDateTimePicker(mainForm)
	dtp.SetParent(mainForm)
	dtp.SetBounds(left, top, 167, 25)
	dtp.SetFormat("yyyy-MM-dd HH:mm:ss")

	// 在xp下应用了style需要解除一个样式属性，应该是冲突引起的
	dtp.SetStyleElements(rtl.Include(0, types.SeFont))

	top += dtp.Height() + 10

	mdtp := vcl.NewMonthCalendar(mainForm)
	mdtp.SetParent(mainForm)
	mdtp.SetBounds(left, top, 250, 250)
	mdtp.SetOnClick(func(vcl.IObject) {
		fmt.Println(mdtp.Date())
	})

	top += mdtp.Height() + 10
	dtp.SetDateTime(time.Now().Add(time.Hour * 48))
	dtp.SetDate(time.Now().AddDate(1, 0, 0))
	fmt.Println("time: ", mdtp.Date())

	btn = vcl.NewButton(mainForm)
	btn.SetParent(mainForm)
	btn.SetLeft(left)
	btn.SetTop(top)
	btn.SetCaption("改变日期")
	btn.SetOnClick(func(vcl.IObject) {
		mdtp.SetDate(time.Now().AddDate(-20, 0, 0))
	})

	// 样式已改变
	mainForm.SetOnStyleChanged(func(sender vcl.IObject) {
		fmt.Println("样式已经改变")
	})

	mainForm.SetAllowDropFiles(true)
	mainForm.SetOnDropFiles(func(sender vcl.IObject, aFileNames []string) {
		fmt.Println("当前拖放文件事件执行，文件数：", len(aFileNames))
		for i, s := range aFileNames {
			fmt.Println("index:", i, ", filename:", s)
		}
	})

	// run
	vcl.Application.Run()
}

func addStyleFileName(list *vcl.TListBox) {
	fd, err := os.Open("..\\..\\bin\\styles\\")
	if err != nil {
		fmt.Println(err)
		if os.IsNotExist(err) {
			return
		}
		return
	}
	defer fd.Close()
	list.Items().BeginUpdate()
	defer list.Items().EndUpdate()
	for {
		files, err := fd.Readdir(100)
		for _, f := range files {
			if !f.IsDir() {
				list.Items().Add(f.Name())
			}
		}
		if err == io.EOF {
			break
		}
		if len(files) == 0 {
			break
		}
	}
}
