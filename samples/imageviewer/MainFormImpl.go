// 在这里写你的事件

package main

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"

	"sort"

	"math"

	"runtime"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/colors"
	"github.com/ying32/govcl/vcl/types/keys"
)

//::private::
type TMainFormFields struct {
	isMouseDown    bool
	mouseDownPos   types.TPoint
	canAcceptTypes []string
	isAutoCenter   bool // 首次加载图片后，如果不改变ImageViewer的位置，则窗口调整大小时居中显示
	imgMousePos    types.TPoint
	orgTitle       string // 保存
	//imgBuffer      *vcl.TPicture
	imgThumb      *vcl.TBitmap
	curDirImages  []string // 以当前打开的文件检索当前目标下的
	curImageIndex int      // 当前浏览的图片索引
}

func (f *TMainForm) OnFormCreate(sender vcl.IObject) {
	hookMainFormWndPrc()

	f.canAcceptTypes = []string{".jpg", ".gif", ".png", ".bmp", ".ico", ".psd"}

	f.SetAllowDropFiles(true)
	//imgBuffer = vcl.NewPicture()
	//imgBuffer.SetOnChange(f.OnPicChanged)
	f.imgThumb = vcl.NewBitmap()
	//f.setPBPreViewPos()
	// 调试下，先显示
	//f.PBPreview.Show()

	f.setWindowState(types.WsMaximized)

	f.ImgViewer.SetStretch(true)
	f.ImgViewer.SetAutoSize(false)
	f.ImgIcon.Picture().Assign(vcl.Application.Icon())
	if len(os.Args) > 1 {
		if f.canAccept(os.Args[1]) {
			f.loadImage(os.Args[1])
		}
	}
	// 调试用
}

func (f *TMainForm) OnFormDestroy(sender vcl.IObject) {
	unHookMainFormWndPrc()

	f.SetAllowDropFiles(false)
	f.imgThumb.Free()
	//imgBuffer.Free()
}

func (f *TMainForm) OnPicChanged(sender vcl.IObject) {
	fmt.Println("图片改变")
}

func (f *TMainForm) OnFormMouseWheel(sender vcl.IObject, shift types.TShiftState, wheelDelta, x, y int32, handled *bool) {
	if wheelDelta > 0 {
		f.zoom(1)
	} else {
		f.zoom(-1)
	}
}

func (f *TMainForm) setWindowState(state types.TWindowState) {
	f.SetWindowState(state)
	f.updateSysButtonState()
}

func (f *TMainForm) getFileIndex(aFileName string) int {
	for i, fname := range f.curDirImages {
		if strings.Compare(fname, aFileName) == 0 {
			return i
		}
	}
	return -1
}

func (f *TMainForm) getCurrentImages(aFileName string) {
	// curDirImages
	// 清零
	f.curDirImages = make([]string, 0)
	path := aFileName[:len(aFileName)-len(filepath.Base(aFileName))]
	fd, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		return
	}
	defer fd.Close()
	for {
		files, err := fd.Readdir(100)
		for _, file := range files {
			if !file.IsDir() {
				fname := path + file.Name()
				if f.canAccept(fname) {
					f.curDirImages = append(f.curDirImages, fname)
				}
			}
		}
		if err == io.EOF {
			break
		}
		if len(files) == 0 {
			break
		}
	}
	// 排序下
	sort.Strings(f.curDirImages)
}

func (f *TMainForm) getRatio() int {
	return int(float64(f.ImgViewer.Width()) / float64(f.ImgViewer.Picture().Width()) * 100)
}

func (f *TMainForm) zoom(val int) {
	var newW, newH, newX, newY int32
	img := f.ImgViewer
	scale := 1.1
	mPos := f.ScreenToClient(vcl.Mouse.CursorPos())

	r := types.TRect{img.Left(), img.Top(), img.Left() + img.Width(), img.Top() + img.Height()}

	if val > 0 {
		if f.getRatio() >= 500 {
			return
		}
	} else if val < 0 {
		if f.getRatio() <= 5 {
			return
		}
	} else if val == 0 {

		newW = img.Picture().Width()
		newH = img.Picture().Height()
		ratio := 1.0
		if img.Picture().Width() > f.Width()-20 {
			newW = f.Width() - 20
			ratio = float64(newW) / float64(img.Picture().Width())
			newH = int32(float64(img.Picture().Height()) * ratio)
		} else if img.Picture().Height() > f.Height()-f.PnlTitleBar.Height()-20 {
			newH = f.Height() - f.PnlTitleBar.Height() - 20
			ratio = float64(newH) / float64(img.Picture().Height())
			newW = int32(float64(img.Picture().Width()) * ratio)
		}
		// 这里调整
		f.ImgViewer.SetBounds((f.Width()-newW)/2, (f.Height()-f.PnlTitleBar.Height()-newH)/2, newW, newH)
		goto theEnd
	}

	if r.PtInRect(mPos) {
		if val > 0 {
			newW = int32(float64(img.Width()) * scale)
			newH = int32(float64(img.Height()) * scale)
			newX = img.Left() - int32(float64(mPos.X-img.Left())*(scale-1.0))
			newY = img.Top() - int32(float64(mPos.Y-img.Top())*(scale-1.0))
		} else {
			newW = int32(float64(img.Width()) / scale)
			newH = int32(float64(img.Height()) / scale)
			newX = img.Left() + int32(float64(mPos.X-img.Left())*(1.0-1.0/scale))
			newY = img.Top() + int32(float64(mPos.Y-img.Top())*(1.0-1.0/scale))
		}
	} else {
		if val > 0 {
			newW = int32(float64(img.Width()) * scale)
			newH = int32(float64(img.Height()) * scale)
		} else {
			newW = int32(float64(img.Width()) / scale)
			newH = int32(float64(img.Height()) / scale)
		}
		newX = img.Left()
		newY = img.Top()

	}
	img.SetBounds(newX, newY, newW, newH)

theEnd:
	f.updateTitle()

	// 决定是否显示
	f.PBPreview.SetVisible(img.Width() >= f.PnlClient.Width() || img.Height() >= f.PnlClient.Height())
	if f.PBPreview.Visible() {
		f.setPBPreViewPos()
		f.Invalidate()
	}
}

func (f *TMainForm) loadImage(aFileName string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err:", err)
		}
	}()
	if aFileName == "" {
		return
	}

	if strings.ToLower(path.Ext(aFileName)) == ".psd" {
		if err := PsdToBitmap(aFileName, f.ImgViewer.Picture().Bitmap()); err != nil {
			fmt.Println("err:", err)
			return
		}
	} else {
		f.ImgViewer.Picture().LoadFromFile(aFileName)
	}

	// 首次加载后居中
	f.isAutoCenter = true
	// 首次设置为true
	f.zoom(0.0)

	imgW := f.ImgViewer.Picture().Width()
	imgH := f.ImgViewer.Picture().Height()
	//f.ImgViewer.SetBounds((f.Width()-imgW)/2, (f.Height()-imgH)/2, imgW, imgH)

	if f.imgThumb.IsValid() {
		f.imgThumb.Free()
	}

	thumbW := f.PBPreview.Width()
	thumbH := int32(float64(thumbW) / float64(imgW) * float64(imgH))
	f.imgThumb = vcl.NewBitmap()
	f.imgThumb.SetSize(thumbW, thumbH)
	f.imgThumb.Canvas().StretchDraw(types.TRect{0, 0, thumbW, thumbH}, f.ImgViewer.Picture().Graphic())

	// libvcl下可以动的gif
	//if !rtl.LcLLoaded() && f.ImgViewer.Picture().Graphic().ClassName() == "TGIFImage" {
	//	vcl.AsGIFImage(f.ImgViewer.Picture().Graphic()).SetAnimate(true)
	//}

	// 获取本目录下的文件名
	f.getCurrentImages(aFileName)
	f.curImageIndex = f.getFileIndex(aFileName)
	stat, _ := os.Stat(aFileName)
	f.orgTitle = fmt.Sprintf("%s (%dKB, %dx%d像素) - 第%d/%d张 ", filepath.Base(aFileName), stat.Size()/1024, imgW, imgH, f.curImageIndex+1, len(f.curDirImages))
	f.updateTitle()

}

func (f *TMainForm) OnFormResize(sender vcl.IObject) {
	// 首次加载后居中
	if f.isAutoCenter {
		fmt.Println("isAutoCenter:", f.isAutoCenter)
		f.zoom(0)
	}
	f.setPBPreViewPos()
	f.resetBtnPrevNextPos()
}

func (f *TMainForm) OnFormKeyDown(sender vcl.IObject, key *types.Char, shift types.TShiftState) {
	switch *key {
	// Left
	case keys.VkLeft:
		f.BtnPrev.Click()
	// Right
	case keys.VkRight:
		f.BtnNext.Click()

	// +
	case keys.VkAdd, keys.VkUp:
		f.zoom(1)

	// -
	case keys.VkSubtract, keys.VkDown:
		f.zoom(-1)
	}
}

func (f *TMainForm) setPBPreViewPos() {
	if f.PBPreview.Visible() {
		f.PBPreview.SetLeft(f.Width() - f.PBPreview.Width() - 20)
		f.PBPreview.SetTop(f.Height() - f.PBPreview.Height() - 50)
	}
}

func (f *TMainForm) updateTitle() {
	title := f.orgTitle
	ratio := f.getRatio()
	if ratio != 100 {
		title += fmt.Sprintf(" - %d%%", f.getRatio())
	}
	title += " - ying32图片浏览器"

	f.SetCaption(title)
	f.LblCaption.SetCaption(title)
}

func (f *TMainForm) canAccept(aFileName string) bool {
	ext := strings.ToLower(filepath.Ext(aFileName))
	for _, s := range f.canAcceptTypes {
		if ext == s {
			return true
		}
	}
	return false
}

func (f *TMainForm) updateSysButtonState() {
	if f.isMax() {
		f.BtnMax.Hide()
		f.BtnRestore.Show()
	} else {
		f.BtnMax.Show()
		f.BtnRestore.Hide()
	}
}

func (f *TMainForm) isMax() bool {
	return f.WindowState() == types.WsMaximized
}

func (f *TMainForm) OnBtnMinClick(sender vcl.IObject) {
	f.setWindowState(types.WsMinimized)
}

func (f *TMainForm) OnBtnCloseClick(sender vcl.IObject) {
	f.Close()
}

func (f *TMainForm) OnBtnMaxClick(sender vcl.IObject) {
	f.setWindowState(types.WsMaximized)
}

func (f *TMainForm) OnBtnRestoreClick(sender vcl.IObject) {
	f.setWindowState(types.WsNormal)
}

func (f *TMainForm) OnBtnMenuClick(sender vcl.IObject) {
	vcl.ShowMessage("别点了，点了我我也不干活！")
}

func (f *TMainForm) OnLblCaptionMouseUp(sender vcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
	f.isMouseDown = false
}

func (f *TMainForm) OnLblCaptionMouseDown(sender vcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
	if button == types.MbLeft {
		// windows下可以用这种，貌似只对libvcl生效，lcl不生效的
		//win.ReleaseCapture()
		//f.Perform(win.WM_SYSCOMMAND, win.SC_MOVE+1, 0)
		f.isMouseDown = true
		f.mouseDownPos.X = x
		f.mouseDownPos.Y = y
	}
}

func (f *TMainForm) OnLblCaptionMouseMove(sender vcl.IObject, shift types.TShiftState, x, y int32) {
	if f.isMouseDown {
		if f.isMax() {
			f.setWindowState(types.WsNormal)
			// 这里做，主要是修复还原后鼠标位置
			ratio := float64(f.Width()) / float64(vcl.Screen.Width())
			x = int32(float64(x) * ratio)
			f.mouseDownPos.X = x
		}
		f.SetLeft(f.Left() + (x - f.mouseDownPos.X))
		f.SetTop(f.Top() + (y - f.mouseDownPos.Y))
	}
}

func (f *TMainForm) OnLblCaptionDblClick(sender vcl.IObject) {
	//if runtime.GOOS == "windows" {
	//	win.ReleaseCapture()
	//	f.Perform(win.WM_NCLBUTTONDBLCLK, win.HTCAPTION, 0)
	//}
	//if f.isMax() {
	//	f.setWindowState(types.WsNormal)
	//} else {
	//	f.setWindowState(types.WsMaximized)
	//}
}

func (f *TMainForm) OnImgViewerMouseDown(sender vcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
	if button == types.MbLeft {
		f.isMouseDown = true
		f.mouseDownPos.X = x
		f.mouseDownPos.Y = y
	}
}

func (f *TMainForm) OnImgViewerMouseUp(sender vcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
	f.isMouseDown = false
}

func (f *TMainForm) OnImgViewerMouseMove(sender vcl.IObject, shift types.TShiftState, x, y int32) {
	f.imgMousePos.X = x
	f.imgMousePos.Y = y
	if f.isMouseDown {
		f.isAutoCenter = false
		f.ImgViewer.SetLeft(f.ImgViewer.Left() + (x - f.mouseDownPos.X))
		f.ImgViewer.SetTop(f.ImgViewer.Top() + (y - f.mouseDownPos.Y))
		f.PBPreview.Invalidate()
	}
}

func (f *TMainForm) OnFormDropFiles(sender vcl.IObject, aFileNames []string) {
	// 只第一个
	if f.canAccept(aFileNames[0]) {
		f.loadImage(aFileNames[0])
	}
}

func (f *TMainForm) OnPBPreviewMouseDown(sender vcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {

}

func (f *TMainForm) OnPBPreviewPaint(sender vcl.IObject) {
	if f.imgThumb.IsValid() && f.PBPreview.Visible() {
		pb := f.PBPreview
		canvas := pb.Canvas()
		canvas.Pen().SetColor(colors.ClBlack)
		canvas.Brush().SetStyle(types.BsClear)
		canvas.Rectangle(0, 0, pb.Width(), pb.Height())
		canvas.Draw(0 /* pb.Height()-ImgThumb.Height()*/, 0, f.imgThumb)

		img := f.ImgViewer

		vRect := types.TRect{}
		if img.Left() <= 0 {
			vRect.Left = int32(math.Abs(float64(img.Left())))
		}
		if img.Top() <= 0 {
			vRect.Top = int32(math.Abs(float64(img.Top())))
		}
		if img.Left()+img.Width() >= f.PnlClient.Width() {
			vRect.SetWidth(f.PnlClient.Width())
		} else {
			vRect.SetWidth(img.Width())
		}
		if img.Top()+img.Height() >= f.PnlClient.Height() {
			vRect.SetHeight(f.PnlClient.Height())
		} else {
			vRect.SetHeight(img.Height())
		}

		wRatio := float64(f.imgThumb.Width()) / float64(img.Width())
		hRatio := float64(f.imgThumb.Height()) / float64(img.Height())

		r := types.TRect{
			int32(float64(vRect.Left) * wRatio),
			int32(float64(vRect.Top) * hRatio),
			int32(float64(vRect.Right) * wRatio),
			int32(float64(vRect.Bottom) * hRatio)}

		canvas.Pen().SetColor(colors.ClGreen)
		canvas.Rectangle(r.Left, r.Top, r.Right, r.Bottom)
	}
}

func (f *TMainForm) OnPBPreviewMouseMove(sender vcl.IObject, shift types.TShiftState, x, y int32) {

}

func (f *TMainForm) OnPBPreviewMouseUp(sender vcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {

}

func (f *TMainForm) resetBtnPrevNextPos() {
	f.BtnPrev.SetLeft(5)
	f.BtnPrev.SetTop((f.PnlClient.Height() - f.BtnPrev.Height()) / 2)
	f.BtnNext.SetLeft(f.PnlClient.Width() - f.BtnNext.Width() - 5)
	f.BtnNext.SetTop((f.PnlClient.Height() - f.BtnNext.Height()) / 2)
}

func (f *TMainForm) showBtnPreNext(val bool) {

	if f.BtnPrev.Visible() == val || f.BtnNext.Visible() == val {
		return
	}

	f.BtnNext.SetVisible(val)
	f.BtnPrev.SetVisible(val)
	f.resetBtnPrevNextPos()
}

func (f *TMainForm) OnPnlClientMouseMove(sender vcl.IObject, shift types.TShiftState, x, y int32) {
	f.showBtnPreNext(x <= f.BtnPrev.Left()+f.BtnPrev.Width() || x >= f.BtnNext.Left())

}

func (f *TMainForm) OnBtnPrevClick(sender vcl.IObject) {
	if len(f.curDirImages) == 0 {
		return
	}
	f.curImageIndex--
	if f.curImageIndex < 0 {
		f.curImageIndex = len(f.curDirImages) - 1
	}
	f.loadImage(f.curDirImages[f.curImageIndex])
}

func (f *TMainForm) OnBtnNextClick(sender vcl.IObject) {
	if len(f.curDirImages) == 0 {
		return
	}

	f.curImageIndex++
	if f.curImageIndex >= len(f.curDirImages) {
		f.curImageIndex = 0
	}
	f.loadImage(f.curDirImages[f.curImageIndex])
}

func (f *TMainForm) OnFormConstrainedResize(sender vcl.IObject, minWidth, minHeight, maxWidth, maxHeight *int32) {
	if runtime.GOOS == "windows" {
		if f.WindowState() == types.WsMaximized {
			*maxHeight = vcl.Screen.WorkAreaHeight()
		}
	}
	*minWidth = 400
	*minHeight = 400
}
