// 在这里写你的事件。
// Write your event here.

package main

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/exts/skinh"
)

func (f *TForm6) OnFormCreate(sender vcl.IObject) {
	f.TrackBar1.SetOnChange(f.OnHSBChange)
	f.TrackBar2.SetOnChange(f.OnHSBChange)
	f.TrackBar3.SetOnChange(f.OnHSBChange)

	f.TrackBar5.SetOnChange(f.OnAeroChange)
	f.TrackBar6.SetOnChange(f.OnAeroChange)
	f.TrackBar7.SetOnChange(f.OnAeroChange)
	f.TrackBar8.SetOnChange(f.OnAeroChange)
	f.TrackBar9.SetOnChange(f.OnAeroChange)
	f.TrackBar10.SetOnChange(f.OnAeroChange)
	f.TrackBar11.SetOnChange(f.OnAeroChange)

	f.TrackBar4.SetOnChange(f.OnMenuAlphaChange)
}

func (f *TForm6) OnHSBChange(sender vcl.IObject) {
	skinh.AdjustHSV(f.TrackBar1.Position(), f.TrackBar2.Position(), f.TrackBar3.Position())
}

func (f *TForm6) OnMenuAlphaChange(sender vcl.IObject) {
	skinh.SetMenuAlpha(f.TrackBar4.Position())
}

func (f *TForm6) OnAeroChange(sender vcl.IObject) {
	skinh.AdjustAero(f.TrackBar5.Position(), f.TrackBar8.Position(), f.TrackBar7.Position(), f.TrackBar6.Position(), 0, 0,
		f.TrackBar9.Position(), f.TrackBar10.Position(), f.TrackBar11.Position())
}
