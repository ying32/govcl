// 在这里写你的事件

package main

import (
	"fmt"

	"github.com/ying32/govcl/samples/mp3Player/bass"
	"github.com/ying32/govcl/vcl"
)

//::private::
type TForm1Fields struct {
	bassPlayer *bass.TBass
}

func (f *TForm1) OnFormCreate(sender vcl.IObject) {
	f.bassPlayer = bass.NewBass()
	f.Timer1.SetEnabled(true)
}

func (f *TForm1) OnFormDestroy(sender vcl.IObject) {
	f.bassPlayer.Close()
	bass.BassFree()
}

func (f *TForm1) OnBtnOpenClick(sender vcl.IObject) {
	if f.OpenDialog1.Execute() {
		fmt.Println(f.bassPlayer.OpenFile(f.OpenDialog1.FileName()))
		mLen, err := f.bassPlayer.GetLength()
		fmt.Println("length:", mLen, ", err:", err)
		f.TbPostion.SetMax(int32(mLen))
		f.bassPlayer.SetVolume(int(f.TbVol.Position()))
	}
}

func (f *TForm1) OnBtnPlayClick(sender vcl.IObject) {
	fmt.Println(f.bassPlayer.Play(false))
}

func (f *TForm1) OnTimer1Timer(sender vcl.IObject) {
	if f.bassPlayer.IsValid() && f.bassPlayer.State == bass.PsPlaying {
		f.LblTime.SetCaption(f.bassPlayer.TimeStrLabel())
		pos, _ := f.bassPlayer.GetPosition()
		f.TbPostion.SetPosition(int32(pos))
	}
}

func (f *TForm1) OnTbVolChange(sender vcl.IObject) {
	f.bassPlayer.SetVolume(int(f.TbVol.Position()))
}

func (f *TForm1) OnTbPostionChange(sender vcl.IObject) {

}

func (f *TForm1) OnBtnPauseClick(sender vcl.IObject) {
	f.bassPlayer.Pause()
}

func (f *TForm1) OnBtnStopClick(sender vcl.IObject) {
	f.bassPlayer.Stop()
	f.Timer1.SetEnabled(false)
	f.bassPlayer.SetPosition(0)
	f.TbPostion.SetPosition(0)
}
