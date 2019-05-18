package main

import (
	"fmt"

	"github.com/ying32/govcl/vcl/types"
)

type TVLCMediaPlayer struct {
	vlcInstance         plibvlc_instance_t
	mediaPlayerInstance plibvlc_media_player_t
}

func NewVLCMediaPlayer(args ...string) *TVLCMediaPlayer {
	v := new(TVLCMediaPlayer)
	v.vlcInstance = libvlc_new(args...)
	if !v.checkVLCInstance() {
		return nil
	}
	v.mediaPlayerInstance = libvlc_media_player_new(v.vlcInstance)
	if !v.checkMediaPlayer() {
		return nil
	}
	return v
}

func (v *TVLCMediaPlayer) Free() {
	v.Stop()
	if v.mediaPlayerInstance != 0 {
		libvlc_media_player_release(v.mediaPlayerInstance)
	}
	if v.vlcInstance != 0 {
		libvlc_release(v.vlcInstance)
	}
}

func (v *TVLCMediaPlayer) checkVLCInstance() bool {
	return v.vlcInstance != 0
}

func (v *TVLCMediaPlayer) checkMediaPlayer() bool {
	return v.mediaPlayerInstance != 0
}

func (v *TVLCMediaPlayer) Playing() bool {
	if !v.checkMediaPlayer() {
		return false
	}
	return toGoBool(libvlc_media_player_is_playing(v.mediaPlayerInstance))
}

func (v *TVLCMediaPlayer) Play() {
	if !v.checkMediaPlayer() {
		return
	}
	libvlc_media_player_play(v.mediaPlayerInstance)
}

func (v *TVLCMediaPlayer) Pause() {
	if !v.checkMediaPlayer() {
		return
	}
	libvlc_media_player_pause(v.mediaPlayerInstance)
}

func (v *TVLCMediaPlayer) Stop() {
	if !v.checkMediaPlayer() {
		return
	}
	libvlc_media_player_stop(v.mediaPlayerInstance)
}

func (v *TVLCMediaPlayer) SethWnd(parenthWnd types.HWND) {
	if !v.checkMediaPlayer() {
		return
	}
	libvlc_media_player_set_hwnd(v.mediaPlayerInstance, parenthWnd)
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
	if !v.checkMediaPlayer() {
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
	if !v.checkMediaPlayer() {
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
	if !v.checkMediaPlayer() {
		return 0
	}
	return libvlc_media_player_get_position(v.mediaPlayerInstance)
}

func (v *TVLCMediaPlayer) SetPosition(pos float32) {
	if !v.checkMediaPlayer() {
		return
	}
	libvlc_media_player_set_position(v.mediaPlayerInstance, pos)
}
