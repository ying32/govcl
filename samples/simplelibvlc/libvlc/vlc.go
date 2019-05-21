package libvlc

import (
	"fmt"
	"runtime"

	"github.com/ying32/govcl/vcl/types"
)

type TVLCMediaPlayer struct {
	vlcInstance         plibvlc_instance_t
	mediaPlayerInstance plibvlc_media_player_t
}

func NewVLCMediaPlayer(args ...string) *TVLCMediaPlayer {
	v := new(TVLCMediaPlayer)
	v.vlcInstance = libvlc_new(args...)
	if !v.VLCValid() {
		return nil
	}
	v.mediaPlayerInstance = libvlc_media_player_new(v.vlcInstance)
	if !v.MediaPlayerValid() {
		return nil
	}
	return v
}

func (v *TVLCMediaPlayer) Free() {
	v.Stop()
	if v.MediaPlayerValid() {
		libvlc_media_player_release(v.mediaPlayerInstance)
	}
	if v.VLCValid() {
		libvlc_release(v.vlcInstance)
	}
}

func (v *TVLCMediaPlayer) VLCValid() bool {
	return v.vlcInstance != 0
}

func (v *TVLCMediaPlayer) MediaPlayerValid() bool {
	return v.mediaPlayerInstance != 0
}

func (v *TVLCMediaPlayer) Playing() bool {
	if !v.MediaPlayerValid() {
		return false
	}
	return toGoBool(libvlc_media_player_is_playing(v.mediaPlayerInstance))
}

func (v *TVLCMediaPlayer) Play() {
	if !v.MediaPlayerValid() {
		return
	}
	libvlc_media_player_play(v.mediaPlayerInstance)
}

func (v *TVLCMediaPlayer) Pause() {
	if !v.MediaPlayerValid() {
		return
	}
	libvlc_media_player_pause(v.mediaPlayerInstance)
}

func (v *TVLCMediaPlayer) Stop() {
	if !v.MediaPlayerValid() {
		return
	}
	libvlc_media_player_stop(v.mediaPlayerInstance)
}

func (v *TVLCMediaPlayer) SethWnd(parenthWnd types.HWND) {
	if !v.MediaPlayerValid() {
		return
	}
	switch runtime.GOOS {
	case "windows":
		libvlc_media_player_set_hwnd(v.mediaPlayerInstance, parenthWnd)
	case "linux":
		libvlc_media_player_set_xwindow(v.mediaPlayerInstance, uint32(parenthWnd))
	case "darwin":
		libvlc_media_player_set_nsobject(v.mediaPlayerInstance, parenthWnd)
	}
}

func (v *TVLCMediaPlayer) LoadFromFile(aFileName string) {
	media := libvlc_media_new_path(v.vlcInstance, aFileName)
	if media != 0 {
		defer libvlc_media_release(media)
		libvlc_media_player_set_media(v.mediaPlayerInstance, media)
	}
}

func (v *TVLCMediaPlayer) LoadFromURL(aURL string) {
	media := libvlc_media_new_location(v.vlcInstance, aURL)
	if media != 0 {
		defer libvlc_media_release(media)
		libvlc_media_player_set_media(v.mediaPlayerInstance, media)
	}
}

func (v *TVLCMediaPlayer) MediaTime() int64 {
	if !v.MediaPlayerValid() {
		return 0
	}
	return int64(libvlc_media_player_get_time(v.mediaPlayerInstance))
}

func (v *TVLCMediaPlayer) MediaTimeString() string {
	l := v.MediaTime()
	if l == 0 {
		return "00:00:00"
	}
	l = l / 1000
	return fmt.Sprintf("%.2d:%.2d:%.2d", l/3600, (l%3600)/60, l%60)
}

func (v *TVLCMediaPlayer) MediaLength() int64 {
	if !v.MediaPlayerValid() {
		return 0
	}
	return int64(libvlc_media_player_get_length(v.mediaPlayerInstance))
}

func (v *TVLCMediaPlayer) MediaLengthString() string {
	l := v.MediaLength()
	if l == 0 {
		return "00:00:00"
	}
	l = l / 1000
	return fmt.Sprintf("%.2d:%.2d:%.2d", l/3600, (l%3600)/60, l%60)
}

func (v *TVLCMediaPlayer) Position() float32 {
	if !v.MediaPlayerValid() {
		return 0
	}
	psi := libvlc_media_player_get_position(v.mediaPlayerInstance)
	return psi
}

func (v *TVLCMediaPlayer) SetPosition(pos float32) {
	if !v.MediaPlayerValid() {
		return
	}
	libvlc_media_player_set_position(v.mediaPlayerInstance, pos)
}
