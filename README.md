# govcl

#### 目录
* [项目介绍](#项目介绍)
* [重要说明](#重要说明)
* [使用方法](#使用方法)
* [项目中的包说明](#项目中的包说明)
* [实例类说明](#实例类说明)
* [支持的组件列表](#支持的组件列表)
* [截图](#截图)
* [备注](#备注)
* [作者信息](#作者信息)
* [代码示例](#代码示例)

#### 项目介绍

> 1、由于现有第三方的Go UI库不是太宠大就是用的不习惯，或者组件太少。就萌生了自己写一个UI库的相法  
> Delphi有些许多优秀的VCL组件，不拿来使用太可惜了。所以就索性做了一套。目前支持Win32跟Win64，  
> 只需要带上一个libvcl.dll即可。  

> 2、项目现在支持VCL标准控件中的大部分，足以满足日常操作了，具体见[支持的组件列表](#支持的组件列表)。  
> 事件方面也支持部分，如下：  
```delphi
 TGoEvent = (geClick, geClose, geFormClose, geFormCloseQuery, geChange,
              geUpDownClick, geTreeViewChange, geListViewChange, geDblClick, gePaint,
              geResize, geShow, geMenuChange, geEnter, geExit, gePopup, geBalloonClick,
              geLinkClick, geExecute, geUpdate, geException, geTimer, geMinimize,
              geRestore, geHide, geKeyDown, geKeyPress, geKeyUp, geMouseDown,
              geMouseEnter, geMouseLeave, geMouseMove, geMouseUp, geMouseWheel);
```

> 3、关于不能跨平台问题，原本是有打算借助CrossVcl这个库做到跨三个平台的，但在尝试中发现在共享  
> 库中创建会报错，调试了下，没有找到原因，所以暂时就放弃这块了，后期视情况看能不能解决这个问题。  

#### 重要说明
**所有的代码只会存储在OSC的[码云](https://gitee.com/ying32/govcl)中，原因在于go包路径的问题。**  
**至于github上会建一个同名的项目[govcl](https://github.com/ying32/govcl)，但不会提交任何代码**  


#### 使用方法
> go get gitee.com/ying32/govcl  

```golang
package main

import (
   "gitee.com/ying32/govcl/vcl"
   "gitee.com/ying32/govcl/api" // 这个视情况导入
)

var (
   mainForm *vcl.TForm
)

func main() {
    vcl.Application.Initialize()
    mainForm = vcl.Application.CreateForm()
    mainForm.SetCaption("Hello")
    mainForm.EnabledMaximize(false)
    mainForm.SetDoubleBuffered(true)
    mainForm.ScreenCenter()
    vcl.Application.Run()
}

```  

#### 项目中的包说明

* api  
  包含各种类型定义、枚举值、DLL函数申明与重新包装  
* dylib  
  仅针对Linux及MacOS，模拟windows下动态调用，需要用到cgo  
* rtl  
  包含Delphi中Set类型操作、内存操作等其它函数  
* win  
  包含windows下的常量、函数、类型定义  
* xui  
  包含一个使用xml创建UI的类  


#### 实例类说明

> 按照Delphi中的Application、 Screen、 Mouse、Clipboard四个类实例是可以直接访问的，不需要释放  
其实组件带有Onwer参数的一般指定TFrom的就好了，这样就不需要手动释放，反之Owner填   
写nil则需要手动调用Free，就像其它非组件类的。  

#### 支持的组件列表
> 现支持组件和非组件类列表：  
>
TApplication    
TForm    
TButton    
TEdit    
TMainMenu    
TPopupMenu    
TMemo    
TCheckBox    
TRadioButton    
TGroupBox    
TLabel    
TListBox    
TComboBox    
TPanel    
TImage    
TLinkLabel    
TSpeedButton   
TSplitter    
TRadioGroup    
TStaticText    
TColorBox    
TColorListBox    
TTrayIcon    
TBalloonHint    
TCategoryPanelGroup    
TOpenDialog    
TSaveDialog    
TColorDialog    
TFontDialog    
TPrintDialog    
TOpenPictureDialog    
TSavePictureDialog    
TSaveTextFileDialog    
TOpenTextFileDialog    
TRichEdit    
TTrackBar    
TImageList    
TUpDown    
TProgressBar    
THotKey    
TDateTimePicker    
TMonthCalendar    
TListView    
TTreeView    
TStatusBar    
TToolBar    
TPageControl  
TTabSheet    
TControl 
TActionList  
TToolButton     
TPaintBox  
TTimer  
> 
TIcon    
TBitmap    
TMemoryStream    
TFont    
TStrings    
TStringList    
TBrush    
TPen    
TMenuItem    
TListGroups    
TPicture    
TListColumns    
TListItems    
TTreeNodes    
TListItem    
TTreeNode      
TScreen    
TMouse    
TListGroup    
TListColumn    
TCollectionItem    
TStatusPanels    
TStatusPanel    
TCanvas    
TObject    
TPngImage    
TJPEGImage    
TGIFImage    
TGIFFrame    
TIniFile  
TRegistry  
TClipboard  
TMonitor  
TMargins  
TList  
TGraphic  
TComponent  

#### 截图

![截图1](https://raw.githubusercontent.com/ying32/govcl/master/Screenshot/1.png)   
![截图2](https://raw.githubusercontent.com/ying32/govcl/master/Screenshot/2.png)    


#### 备注
**文件名后面带有def的为手动编写**   

#### 作者信息
by: ying32  


#### 代码示例  


```golang
// govcl project main.go
// go.exe build -i -ldflags="-H windowsgui"
package main

import (
	"fmt"

	"gitee.com/ying32/govcl/vcl"
	"gitee.com/ying32/govcl/vcl/api"
	"gitee.com/ying32/govcl/vcl/rtl"
	"gitee.com/ying32/govcl/vcl/win"
)

var (
	mainForm *vcl.TForm
	trayicon *vcl.TTrayIcon
)

func main() {

	// 异常捕获
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Exception: ", err)
			vcl.ShowMessage(err.(error).Error())
		}
	}()

	fmt.Println("main")
	icon := vcl.NewIcon()
	//icon.LoadFromFile(".\\imgs\\0.ico")
	icon.LoadFromResourceID(rtl.MainInstance(), 3)
	defer icon.Free()
	vcl.Application.Initialize()

	vcl.Application.SetOnException(func(vcl.IObject, vcl.IObject) {
		fmt.Println("exception.")
	})

	vcl.Application.SetIcon(icon)
	vcl.Application.SetTitle("Hello World!")
	vcl.Application.SetMainFormOnTaskBar(true)
	mainForm = vcl.Application.CreateForm()
	mainForm.SetWidth(600)
	mainForm.SetHeight(400)
	mainForm.SetOnClose(func(Sender vcl.IObject, Action uintptr) {
		fmt.Println("close")
	})

	fmt.Println("MainForm ClientRect: ", mainForm.ClientRect())

	mainForm.SetOnCloseQuery(func(Sender vcl.IObject, CanClose uintptr) {
		rtl.SetFormCanClose(CanClose, vcl.MessageDlg("是否退出?", api.MtInformation, api.MbYes, api.MbNo) == vcl.MrYes)
		fmt.Println("OnCloseQuery")
	})

	mainForm.SetCaption(vcl.Application.Title())
	mainForm.EnabledMaximize(false)
	mainForm.SetDoubleBuffered(true)
	mainForm.SetPosition(api.PoScreenCenter)
	mainForm.SetKeyPreview(true)
	mainForm.SetOnKeyDown(func(Sender vcl.IObject, Key uintptr, Shift int32) {
		fmt.Println(rtl.InSets(uint32(Shift), api.SsCtrl))
		fmt.Println(rtl.GetKey(Key))
	})

	mainForm.SetOnMouseDown(func(sender vcl.IObject, button, shift, x, y int32) {
		fmt.Println("Button:", button == api.MbLeft, ", X:", x, ", y:", y)
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
	trayicon.SetIcon(icon)
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
	img.Picture().LoadFromFile(".\\imgs\\1.jpg")
	//img.SetStretch(true)
	img.SetProportional(true)

	// linklabel
	linklbl := vcl.NewLinkLabel(mainForm)
	linklbl.SetAlign(api.AlBottom)
	linklbl.SetCaption("<a href=\"https://github.com/ying32/govcl\">govcl测试链接</a>")
	linklbl.SetParent(mainForm)
	linklbl.SetOnLinkClick(func(sender vcl.IObject, link string, linktype int32) {
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
		s := api.GoStrToDStr("ffff中f")
		mem.Write(s, int32(api.DStrLen(s)*2))
		mem.SaveToFile("test.txt")
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
	button.SetAlign(api.AlRight)

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
	button2.SetAlign(api.AlTop)

	combo := vcl.NewComboBox(mainForm)
	combo.SetAlign(api.AlBottom)
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
	page.SetAlign(api.AlBottom)
	sheet := vcl.NewTabSheet(mainForm)
	sheet.SetPageControl(page)
	sheet.SetCaption("第一页")

	// 需要先将TabSheet设置了父窗口，TListView才可用，不然就会报错
	lv1 := vcl.NewListView(mainForm)
	lv1.SetAlign(api.AlClient)
	lv1.SetParent(sheet)

	lv1.SetViewStyle(api.VsReport)
	lv1.SetRowSelect(true)
	lv1.SetReadOnly(true)
	lv1.SetGridLines(true)
	col := lv1.Columns().Add()
	col.SetCaption("序号")
	col.SetWidth(100)
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

	tv1 := vcl.NewTreeView(mainForm)
	tv1.SetAutoExpand(true)
	tv1.SetParent(sheet)
	tv1.SetAlign(api.AlClient)
	tv1.SetOnClick(func(vcl.IObject) {
		if tv1.SelectionCount() > 0 {
			node := tv1.Selected()
			fmt.Println("text:", node.Text(), ", index:", node.Index())
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

	vcl.Application.Run()
}

```  