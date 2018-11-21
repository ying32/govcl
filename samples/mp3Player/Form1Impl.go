// 在这里写你的事件

package main

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/ying32/govcl/vcl/rtl/version"

	"github.com/ying32/govcl/vcl/rtl"

	"github.com/ying32/govcl/vcl/types"

	"github.com/ying32/govcl/samples/mp3Player/bass"
	"github.com/ying32/govcl/vcl"
)

//::private::
type TForm1Fields struct {
	bassPlayer *bass.TBass
	playCtl    *TPlayControl
	progress   *TImageTrackBar
	volbar     *TImageTrackBar

	// 仅windows
	taskbar1         *vcl.TTaskbar
	isSupportTaskbar bool
	thumbPreviewBmp  *vcl.TBitmap
}

func (f *TForm1) OnFormCreate(sender vcl.IObject) {
	f.SetDoubleBuffered(true)
	f.EnabledMaximize(false)
	f.SetColor(0x39302c)

	f.bassPlayer = bass.NewBass()
	f.Timer1.SetEnabled(true)
	f.playCtl = NewPlayControl(f)
	f.playCtl.SetParent(f.Panel2)
	f.playCtl.SetAlign(types.AlClient)
	f.playCtl.OnSelect = f.OnPlayListSelect
	f.playCtl.SingerPic = vcl.BitmapFromObj(f.ImgSinger.Picture().Graphic())

	f.progress = NewImageTrackBar(f)
	f.progress.SetParent(f)
	f.progress.SetImages(f.ImgThumb, f.ImgBk, f.ImgFore)
	f.progress.SetLeft(58)
	f.progress.SetTop(288)
	f.progress.SetWidth(200)
	f.progress.SetPosition(0)

	f.volbar = NewImageTrackBar(f)
	f.volbar.SetParent(f)
	f.volbar.SetImages(f.ImgThumb, f.ImgBk, f.ImgFore)
	f.volbar.SetLeft(82)
	f.volbar.SetTop(318)
	f.volbar.SetWidth(80)
	f.volbar.SetPosition(80)
	f.volbar.OnTrackChange = f.OnVolChange

	f.addFoler("F:\\KuGou\\")

	f.isSupportTaskbar = runtime.GOOS == "windows" && version.OSVersion.CheckMajorMinor(6, 1) && !rtl.LcLLoaded()
	if f.isSupportTaskbar {
		f.thumbPreviewBmp = vcl.NewBitmap()
		f.thumbPreviewBmp.SetPixelFormat(types.Pf32bit)
		f.thumbPreviewBmp.SetSize(f.ImgSinger.Width(), f.ImgSinger.Height())

		f.thumbPreviewBmp.Assign(f.ImgSinger.Picture().Graphic())
		f.taskbar1 = vcl.NewTaskbar(f)
		f.taskbar1.SetOnThumbPreviewRequest(f.onTaskbar1ThumbPreviewRequest)
		f.taskbar1.SetOnThumbButtonClick(f.onTaskbar1ThumbButtonClick)
		f.taskbar1.SetTabProperties(rtl.Include(0, types.CustomizedPreview))

		// buttons
		tbtn := f.taskbar1.TaskBarButtons().Add()
		tbtn.SetHint("上一曲")
		f.ImageList1.GetIcon(0, tbtn.Icon())

		tbtn = f.taskbar1.TaskBarButtons().Add()
		tbtn.SetHint("播放")
		f.ImageList1.GetIcon(1, tbtn.Icon())

		tbtn = f.taskbar1.TaskBarButtons().Add()
		tbtn.SetButtonState(rtl.Include(tbtn.ButtonState(), types.TbsHidden))
		tbtn.SetHint("暂停")
		f.ImageList1.GetIcon(2, tbtn.Icon())

		tbtn = f.taskbar1.TaskBarButtons().Add()
		tbtn.SetHint("下一曲")
		f.ImageList1.GetIcon(3, tbtn.Icon())

	}

}

func (f *TForm1) onTaskbar1ThumbPreviewRequest(sender vcl.IObject, aPreviewHeight, aPreviewWidth int32, previewBitmap *vcl.TBitmap) {
	//fmt.Println("OnTaskbar1ThumbPreviewRequest")
	previewBitmap.Assign(f.thumbPreviewBmp)
}

func (f *TForm1) onTaskbar1ThumbButtonClick(sender vcl.IObject, aButtonID int32) {

	switch aButtonID {
	case 0:
		f.BtnPrev.Click()
	case 1:
		f.BtnPlay.Click()
	case 2:
		f.BtnPause.Click()
	case 3:
		f.BtnNext.Click()
	}
}

func (f *TForm1) setTaskbarHint(str string) {
	if f.isSupportTaskbar {
		f.taskbar1.SetToolTip(str)
		f.taskbar1.SetOverlayHint(str)
	}
}

func (f *TForm1) setTasbarButtonState(aPlayVisible, aPauseVisible bool) {

	if f.isSupportTaskbar {
		var state types.TThumbButtonStates
		// play
		tbtn := f.taskbar1.TaskBarButtons().Items(1)
		state = tbtn.ButtonState()
		if aPlayVisible {
			state = rtl.Exclude(state, types.TbsHidden)
		} else {
			state = rtl.Include(state, types.TbsHidden)
		}
		tbtn.SetButtonState(state)

		// pause
		tbtn = f.taskbar1.TaskBarButtons().Items(2)
		state = tbtn.ButtonState()
		if aPauseVisible {
			state = rtl.Exclude(state, types.TbsHidden)
		} else {
			state = rtl.Include(state, types.TbsHidden)
		}
		tbtn.SetButtonState(state)
	}
}

func (f *TForm1) OnFormDestroy(sender vcl.IObject) {
	if f.isSupportTaskbar {
		f.thumbPreviewBmp.Free()
	}
	f.bassPlayer.Close()
	bass.BassFree()
}

func (f *TForm1) OnMIAddFileClick(sender vcl.IObject) {
	if f.OpenDialog1.Execute() {
		f.addFile(f.OpenDialog1.FileName())
	}
}

func (f *TForm1) OnMIAddFolderClick(sender vcl.IObject) {
	if ok, str := vcl.SelectDirectory2("选择目录", "", rtl.Include(0, types.SdNewUI, types.SdShowEdit), nil); ok {
		f.addFoler(str)
	}
}

func (f *TForm1) addFile(fileName string) {
	name := filepath.Base(fileName)
	nameArr := strings.Split(name[:len(name)-4], "-")
	if len(nameArr) >= 2 {
		lenVal := int32(bass.GetFileLength(fileName))
		f.playCtl.Add(TPlayListItem{strings.TrimSpace(nameArr[1]), strings.TrimSpace(nameArr[0]), lenVal, "", fileName})
	}
}

func (f *TForm1) addFoler(rootPath string) {
	filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(info.Name()) == ".mp3" {
			f.addFile(path)
		}
		return nil
	})
}

func (f *TForm1) OnPlayListSelect(sender vcl.IObject, item TPlayListItem) {
	f.stopPlay()
	f.bassPlayer.Close()
	f.bassPlayer.OpenFile(item.FileName)
	f.bassPlayer.SetVolume(f.volbar.Position())
	f.play()
	str := item.Caption + " - " + item.Singer
	f.SetCaption(str)
	f.setTaskbarHint(str)
}

func (f *TForm1) OnBtnPlayClick(sender vcl.IObject) {
	f.play()
}

func (f *TForm1) OnTimer1Timer(sender vcl.IObject) {
	if f.bassPlayer.IsValid() && f.bassPlayer.State == bass.PsPlaying {
		f.LblTime.SetCaption(f.bassPlayer.TimeStrLabel())
		pos, _ := f.bassPlayer.GetPosition()
		mLen, _ := f.bassPlayer.GetLength()

		f.progress.SetPosition(int(float32(pos) / float32(mLen) * 100))

		caption := []rune(f.Caption())
		if len(caption) > 0 {
			temp := caption
			c := temp[:1]
			temp = temp[1:]
			f.SetCaption(string(temp) + string(c))
		}
	}
}

func (f *TForm1) OnVolChange(sender vcl.IObject) {
	f.bassPlayer.SetVolume(f.volbar.Position())
}

func (f *TForm1) OnBtnPauseClick(sender vcl.IObject) {
	f.pause()
}

func (f *TForm1) OnBtnMinClick(sender vcl.IObject) {
	vcl.Application.Minimize()
}

func (f *TForm1) OnBtnCloseClick(sender vcl.IObject) {
	vcl.Application.Terminate()
}

func (f *TForm1) stopPlay() {
	f.bassPlayer.Stop()
	f.Timer1.SetEnabled(false)
	f.bassPlayer.SetPosition(0)
	f.progress.SetPosition(0)

	f.BtnPause.Hide()
	f.BtnPlay.Show()

	f.setTasbarButtonState(true, false)
}

func (f *TForm1) play() {
	f.bassPlayer.Play(false)

	f.BtnPlay.Hide()
	f.BtnPause.Show()
	f.Timer1.SetEnabled(true)

	f.setTasbarButtonState(false, true)
}

func (f *TForm1) pause() {
	f.bassPlayer.Pause()
	f.BtnPause.Hide()
	f.BtnPlay.Show()
	f.Timer1.SetEnabled(false)

	f.setTasbarButtonState(true, false)
}

func (f *TForm1) OnPanel1MouseDown(sender vcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
	if button == types.MbLeft {
		//f.Perform(messages.WM_SYSCOMMAND, message)
	}
}

func (f *TForm1) OnBtnPrevClick(sender vcl.IObject) {
	f.playCtl.Prev()
}

func (f *TForm1) OnBtnNextClick(sender vcl.IObject) {
	f.playCtl.Next()
}
