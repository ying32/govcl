package main

import (
	"fmt"
	"unsafe"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/dylib"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/colors"
)

type (
	plibvlc_instance_t     = uintptr
	plibvlc_media_player_t = uintptr
	plibvlc_media_t        = uintptr
)

var (
	//libvlccoredll = dylib.NewLazyDLL("libvlccore.dll")
	libvlc = dylib.NewLazyDLL("libvlc.dll")

	_libvlc_new     = libvlc.NewProc("libvlc_new")
	_libvlc_release = libvlc.NewProc("libvlc_release")
	_libvlc_errmsg  = libvlc.NewProc("libvlc_errmsg")

	_libvlc_media_player_new        = libvlc.NewProc("libvlc_media_player_new")
	_libvlc_media_player_release    = libvlc.NewProc("libvlc_media_player_release")
	_libvlc_media_player_play       = libvlc.NewProc("libvlc_media_player_play")
	_libvlc_media_player_is_playing = libvlc.NewProc("libvlc_media_player_is_playing")
	_libvlc_media_player_stop       = libvlc.NewProc("libvlc_media_player_stop")
	_libvlc_media_player_set_pause  = libvlc.NewProc("libvlc_media_player_set_pause")
	_libvlc_media_player_pause      = libvlc.NewProc("libvlc_media_player_pause")

	_libvlc_media_player_set_xwindow = libvlc.NewProc("libvlc_media_player_set_xwindow")
	_libvlc_media_player_set_hwnd    = libvlc.NewProc("libvlc_media_player_set_hwnd")

	_libvlc_media_new_path         = libvlc.NewProc("libvlc_media_new_path")
	_libvlc_media_new_location     = libvlc.NewProc("libvlc_media_new_location")
	_libvlc_media_release          = libvlc.NewProc("libvlc_media_release")
	_libvlc_media_player_set_media = libvlc.NewProc("libvlc_media_player_set_media")
	_libvlc_media_player_get_media = libvlc.NewProc("libvlc_media_player_get_media")
)

type TMainForm struct {
	*vcl.TForm
	BtnPlay   *vcl.TButton
	BtnStop   *vcl.TButton
	BtnPause  *vcl.TButton
	PnlClient *vcl.TPanel
	PnlLeft   *vcl.TPanel

	libvlcInstance      plibvlc_instance_t
	mediaplayerInstance plibvlc_media_player_t
	//media               plibvlc_media_t
}

var MainForm *TMainForm

func (f *TMainForm) OnFormCreate(sender vcl.IObject) {
	f.SetCaption("libvlc测试")
	f.ScreenCenter()
	fmt.Println("Create")

	f.PnlLeft = vcl.NewPanel(f)
	f.PnlLeft.SetParent(f)
	f.PnlLeft.SetAlign(types.AlLeft)
	f.PnlLeft.SetWidth(140)
	f.PnlLeft.SetShowCaption(false)
	f.PnlLeft.SetParentBackground(false)

	f.BtnPlay = vcl.NewButton(f)
	f.BtnPlay.SetParent(f.PnlLeft)
	f.BtnPlay.SetCaption("播放")
	f.BtnPlay.SetTop(10)
	f.BtnPlay.SetLeft(10)

	f.BtnPause = vcl.NewButton(f)
	f.BtnPause.SetParent(f.PnlLeft)
	f.BtnPause.SetCaption("暂停")
	f.BtnPause.SetTop(f.BtnPlay.Top() + f.BtnPlay.Height() + 10)
	f.BtnPause.SetLeft(10)

	f.BtnStop = vcl.NewButton(f)
	f.BtnStop.SetParent(f.PnlLeft)
	f.BtnStop.SetCaption("停止")
	f.BtnStop.SetTop(f.BtnPause.Top() + f.BtnPause.Height() + 10)
	f.BtnStop.SetLeft(10)

	f.PnlClient = vcl.NewPanel(f)
	f.PnlClient.SetParent(f)
	f.PnlClient.SetAlign(types.AlClient)
	f.PnlClient.SetShowCaption(false)
	f.PnlClient.SetColor(colors.ClBlack)
	f.PnlClient.SetParentBackground(false)

	f.libvlcInstance = libvlc_new()
	if f.libvlcInstance != 0 {
		f.mediaplayerInstance = libvlc_media_player_new(f.libvlcInstance)
		if f.mediaplayerInstance != 0 {
			libvlc_media_player_set_hwnd(f.mediaplayerInstance, f.PnlClient.Handle())

			//f.media = libvlc_media_new_location(f.libvlcInstance, "http://")
			media := libvlc_media_new_path(f.libvlcInstance, "I:\\30分钟教你学会超火的日语歌曲- PLANET.mp4")
			if media != 0 {
				defer libvlc_media_release(media)
				libvlc_media_player_set_media(f.mediaplayerInstance, media)

			}
		}
	}
}

func (f *TMainForm) OnFormDestroy(sender vcl.IObject) {
	if f.mediaplayerInstance != 0 && libvlc_media_player_is_playing(f.mediaplayerInstance) != 0 {
		libvlc_media_player_stop(f.mediaplayerInstance)
	}

	if f.mediaplayerInstance != 0 {
		libvlc_media_player_release(f.mediaplayerInstance)
	}
	if f.libvlcInstance != 0 {
		libvlc_release(f.libvlcInstance)
	}

	fmt.Println("Free")
}

func (f *TMainForm) OnBtnPlayClick(sender vcl.IObject) {
	fmt.Println("播放单击:", libvlc_media_player_is_playing(f.mediaplayerInstance))
	if libvlc_media_player_is_playing(f.mediaplayerInstance) != 0 {
		return
	}
	// play
	libvlc_media_player_play(f.mediaplayerInstance)
}

func (f *TMainForm) OnBtnStopClick(sender vcl.IObject) {
	fmt.Println("停止单击:", libvlc_media_player_is_playing(f.mediaplayerInstance))
	if libvlc_media_player_is_playing(f.mediaplayerInstance) == 0 {
		return
	}
	// play
	libvlc_media_player_stop(f.mediaplayerInstance)
}

func (f *TMainForm) OnBtnPauseClick(sender vcl.IObject) {
	fmt.Println("停止暂停:", libvlc_media_player_is_playing(f.mediaplayerInstance))
	if f.mediaplayerInstance == 0 {
		return
	}
	libvlc_media_player_pause(f.mediaplayerInstance)
}

func main() {
	//fmt.Println(libvlccoredll.Load())
	fmt.Println(libvlc.Load())

	vcl.Application.Initialize()
	vcl.Application.CreateForm(&MainForm, true)
	vcl.Application.Run()

}

func libvlc_new(args ...string) plibvlc_instance_t {
	var ptr uintptr
	if len(args) > 0 {
		ps := make([]uintptr, len(args))
		for i := 0; i < len(ps); i++ {
			ps[i] = uintptr(unsafe.Pointer(&([]byte(args[i] + "\x00")[0])))
		}
		ptr = uintptr(unsafe.Pointer(&ps[0]))
	}
	ret, _, _ := _libvlc_new.Call(uintptr(len(args)), ptr)
	return ret
}

func libvlc_release(p_instance plibvlc_instance_t) {
	_libvlc_release.Call(p_instance)
}

func libvlc_errmsg() {

	fmt.Println(_libvlc_errmsg.Call())

}

// player

func libvlc_media_player_new(p_libvlc_instance plibvlc_instance_t) plibvlc_media_player_t {
	fmt.Println("p_libvlc_instance:", p_libvlc_instance)
	ret, _, _ := _libvlc_media_player_new.Call(p_libvlc_instance)
	return ret
}

func libvlc_media_player_release(p_mi plibvlc_media_player_t) {
	_libvlc_media_player_release.Call(p_mi)
}

func libvlc_media_player_play(p_mi plibvlc_media_player_t) int {
	ret, _, _ := _libvlc_media_player_play.Call(p_mi)
	return int(ret)
}

func libvlc_media_player_is_playing(p_mi plibvlc_media_player_t) int {
	ret, _, _ := _libvlc_media_player_is_playing.Call(p_mi)
	return int(ret)
}

func libvlc_media_player_stop(p_mi plibvlc_media_player_t) {
	_libvlc_media_player_stop.Call(p_mi)
}

func libvlc_media_player_set_pause(mp plibvlc_media_player_t, do_pause int) {
	_libvlc_media_player_set_pause.Call(mp, uintptr(do_pause))
}

func libvlc_media_player_pause(p_mi plibvlc_media_player_t) {
	_libvlc_media_player_pause.Call(p_mi)
}

func libvlc_media_player_set_xwindow(p_mi plibvlc_media_player_t, drawable uint32) {
	_libvlc_media_player_set_xwindow.Call(p_mi, uintptr(drawable))
}

func libvlc_media_player_set_hwnd(p_mi plibvlc_media_player_t, drawable uintptr) {
	_libvlc_media_player_set_hwnd.Call(p_mi, drawable)
}

func libvlc_media_new_path(p_instance plibvlc_instance_t, path string) plibvlc_media_t {
	ret, _, _ := _libvlc_media_new_path.Call(p_instance, uintptr(unsafe.Pointer(&([]byte(path + "\x00")[0]))))
	return ret
}

func libvlc_media_new_location(p_instance plibvlc_instance_t, psz_mrl string) plibvlc_media_t {
	ret, _, _ := _libvlc_media_new_location.Call(p_instance, uintptr(unsafe.Pointer(&([]byte(psz_mrl + "\x00")[0]))))
	return ret
}

func libvlc_media_release(p_md plibvlc_media_t) {
	_libvlc_media_release.Call(p_md)
}

func libvlc_media_player_set_media(p_mi plibvlc_media_player_t, p_md plibvlc_media_t) {
	_libvlc_media_player_set_media.Call(p_mi, p_md)
}

func libvlc_media_player_get_media(p_mi plibvlc_media_player_t) plibvlc_media_t {
	ret, _, _ := _libvlc_media_player_get_media.Call(p_mi)
	return ret
}
