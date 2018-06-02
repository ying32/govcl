// 在这里写你的事件

package main

import (
	"fmt"

	"github.com/ying32/govcl/samples/mp3Player/bass"
	"github.com/ying32/govcl/vcl"
)

var (
	bassPlayer *bass.TBass
)

func (f *TForm1) OnForm1Create(sender vcl.IObject) {
	bassPlayer = bass.NewBass()
	f.Timer1.SetEnabled(true)
}

func (f *TForm1) OnForm1Destroy(sender vcl.IObject) {
	bassPlayer.Close()
	bass.BassFree()
}

func (f *TForm1) OnBtnOpenClick(sender vcl.IObject) {
	if f.OpenDialog1.Execute() {
		fmt.Println(bassPlayer.OpenFile(f.OpenDialog1.FileName()))
		mLen, err := bassPlayer.GetLength()
		fmt.Println("length:", mLen, ", err:", err)
		f.TbPostion.SetMax(int32(mLen))
	}
}

func (f *TForm1) OnBtnPlayClick(sender vcl.IObject) {
	fmt.Println(bassPlayer.Play(false))
}

func (f *TForm1) OnTimer1Timer(sender vcl.IObject) {
	if bassPlayer.IsValid() && bassPlayer.State == bass.PsPlaying {
		f.LblTime.SetCaption(bassPlayer.TimeStrLabel())
		pos, _ := bassPlayer.GetPosition()
		f.TbPostion.SetPosition(int32(pos))

	}
}

func (f *TForm1) OnTbVolChange(sender vcl.IObject) {
	if bassPlayer.IsValid() {
		bassPlayer.SetVolume(int(f.TbVol.Position()))
	}
}

func (f *TForm1) OnTbPostionChange(sender vcl.IObject) {

}
