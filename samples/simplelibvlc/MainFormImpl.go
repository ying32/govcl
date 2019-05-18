// 在这里写你的事件

package main

import (
	"fmt"

	"github.com/ying32/govcl/vcl"
)

//::private::
type TMainFormFields struct {
	player *TVLCMediaPlayer
}

func (f *TMainForm) OnFormCreate(sender vcl.IObject) {
	f.ScreenCenter()
	f.player = NewVLCMediaPlayer()
	if f.player == nil {
		vcl.ShowMessage("创建MediaPlayer失败。")
		return
	}
	f.player.SethWnd(f.PnlVideo.Handle())
	f.player.LoadFromFile("I:\\30分钟教你学会超火的日语歌曲- PLANET.mp4")

}

func (f *TMainForm) OnFormDestroy(sender vcl.IObject) {
	if f.player != nil {
		f.player.Free()
	}
}

func (f *TMainForm) OnActPlayExecute(sender vcl.IObject) {
	f.Timer1.SetEnabled(true)
	f.player.Play()
}

func (f *TMainForm) OnActPlayUpdate(sender vcl.IObject) {
	vcl.ActionFromObj(sender).SetEnabled(f.player != nil && !f.player.Playing())
}

func (f *TMainForm) OnActStopExecute(sender vcl.IObject) {
	f.Timer1.SetEnabled(false)
	f.player.Stop()
	f.LblCurTime.SetCaption("00:00:00")
	f.LblTotalTime.SetCaption("00:00:00")
	f.TrackBar1.SetPosition(0)
}

func (f *TMainForm) OnActStopUpdate(sender vcl.IObject) {
	vcl.ActionFromObj(sender).SetEnabled(f.player != nil && f.player.Playing())
}

func (f *TMainForm) OnActPauseExecute(sender vcl.IObject) {
	f.Timer1.SetEnabled(false)
	f.player.Pause()
}

func (f *TMainForm) OnActPauseUpdate(sender vcl.IObject) {
	vcl.ActionFromObj(sender).SetEnabled(f.player != nil && f.player.Playing())
}

func (f *TMainForm) OnTimer1Timer(sender vcl.IObject) {
	if f.player != nil && f.player.checkMediaPlayer() {
		f.LblCurTime.SetCaption(f.player.MediaTimeString())
		f.LblTotalTime.SetCaption(f.player.MediaLengthString())
		fmt.Println("pos:", f.player.Position())
	}
}
