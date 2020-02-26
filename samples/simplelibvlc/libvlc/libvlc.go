/*

  这个文件是给github.com/ying32/govcl/samples/simaplelibvlc 测试使用
  翻译自fpc的libvlc.pp文件
  By:ying32
*/

package libvlc

import (
	"fmt"
	"math"
	"runtime"
	"unsafe"

	"github.com/ying32/dylib/floatpatch"

	"github.com/ying32/dylib"
)

type (
	Pstring = uintptr
	// type def
	libvlc_event_manager_t     uintptr
	libvlc_instance_t          uintptr
	libvlc_log_iterator_t      uintptr
	libvlc_log_t               uintptr
	libvlc_media_discoverer_t  uintptr
	libvlc_media_library_t     uintptr
	libvlc_media_list_player_t uintptr
	libvlc_media_list_t        uintptr
	libvlc_media_player_t      uintptr
	libvlc_media_t             uintptr

	plibvlc_audio_output_t       *libvlc_audio_output_t       //= uintptr               //^libvlc_audio_output_t;
	plibvlc_event_manager_t      uintptr                      //^libvlc_event_manager_t;
	plibvlc_event_t              *libvlc_event_t              //= uintptr               //^libvlc_event_t;
	plibvlc_instance_t           uintptr                      //^libvlc_instance_t;
	plibvlc_log_iterator_t       uintptr                      //^libvlc_log_iterator_t;
	plibvlc_log_message_t        *libvlc_log_message_t        //= uintptr               //^libvlc_log_message_t;
	plibvlc_log_t                uintptr                      //^libvlc_log_t;
	plibvlc_media_discoverer_t   uintptr                      //^libvlc_media_discoverer_t;
	plibvlc_media_library_t      uintptr                      //^libvlc_media_library_t;
	plibvlc_media_list_player_t  uintptr                      //^libvlc_media_list_player_t;
	plibvlc_media_list_t         uintptr                      // ^libvlc_media_list_t;
	plibvlc_media_player_t       uintptr                      //^libvlc_media_player_t;
	plibvlc_media_stats_t        *libvlc_media_stats_t        //= uintptr               //^libvlc_media_stats_t;
	plibvlc_media_t              uintptr                      //^libvlc_media_t;
	plibvlc_media_track_info_t   *libvlc_media_track_info_t   //= uintptr               //^libvlc_media_track_info_t;
	plibvlc_module_description_t *libvlc_module_description_t //= uintptr               //^libvlc_module_description_t;
	plibvlc_track_description_t  *libvlc_track_description_t  //= uintptr               // ^libvlc_track_description_t;

	libvlc_time_t        int64
	libvlc_log_message_t struct {
		i_severity  int     //: cint;
		psz_type    uintptr //: ^cchar;
		psz_name    uintptr //: ^cchar;
		psz_header  uintptr //: ^cchar;
		psz_message uintptr //: ^cchar;
	}

	libvlc_event_type_t int     //cint;
	libvlc_callback_t   uintptr //procedure (_para1:Plibvlc_event_t; _para2:pointer);cdecl;

	libvlc_module_description_t struct {
		psz_name      uintptr                      //: ^cchar;
		psz_shortname uintptr                      //: ^cchar;
		psz_longname  uintptr                      //: ^cchar;
		psz_help      uintptr                      //: ^cchar;
		p_next        *libvlc_module_description_t //: ^libvlc_module_description_t;
	}

	// callback
	libvlc_video_lock_cb       uintptr //= function (opaque:pointer; planes:Ppointer):pointer;cdecl;
	libvlc_video_unlock_cb     uintptr //= procedure (opaque:pointer; picture:pointer; planes:Ppointer);cdecl;
	libvlc_video_display_cb    uintptr //= procedure (opaque:pointer; picture:pointer);cdecl;
	libvlc_video_format_cb     uintptr //= function (opaque:Ppointer; chroma:pcchar; width:pcunsigned; height:pcunsigned; pitches:pcunsigned; lines:pcunsigned):cunsigned;cdecl;
	libvlc_video_cleanup_cb    uintptr //= procedure (opaque:pointer);cdecl;
	libvlc_audio_play_cb       uintptr //= procedure (data:pointer; samples:pointer; count:cunsigned; pts:int64_t);cdecl;
	libvlc_audio_pause_cb      uintptr //= procedure (data:pointer; pts:int64_t);cdecl;
	libvlc_audio_resume_cb     uintptr //= procedure (data:pointer; pts:int64_t);cdecl;
	libvlc_audio_flush_cb      uintptr //= procedure (data:pointer; pts:int64_t);cdecl;
	libvlc_audio_drain_cb      uintptr //= procedure (data:pointer);cdecl;
	libvlc_audio_set_volume_cb uintptr //= procedure (data:pointer; volume:cfloat; mute:_Bool);cdecl;
	libvlc_audio_setup_cb      uintptr //= function (data:Ppointer; format:pcchar; rate:pcunsigned; channels:pcunsigned):cint;cdecl;
	libvlc_audio_cleanup_cb    uintptr //= procedure (data:pointer);cdecl;

)

type libvlc_meta_t int32

const (
	libvlc_meta_Title libvlc_meta_t = iota + 0
	libvlc_meta_Artist
	libvlc_meta_Genre
	libvlc_meta_Copyright
	libvlc_meta_Album
	libvlc_meta_TrackNumber
	libvlc_meta_Description
	libvlc_meta_Rating
	libvlc_meta_Date
	libvlc_meta_Setting
	libvlc_meta_URL
	libvlc_meta_Language
	libvlc_meta_NowPlaying
	libvlc_meta_Publisher
	libvlc_meta_EncodedBy
	libvlc_meta_ArtworkURL
	libvlc_meta_TrackID
)

type libvlc_state_t int32

const (
	libvlc_NothingSpecial libvlc_state_t = iota + 0
	libvlc_Opening
	libvlc_Buffering
	libvlc_Playing
	libvlc_Paused
	libvlc_Stopped
	libvlc_Ended
	libvlc_Error
)

type libvlc_media_option_t int32

const (
	libvlc_media_option_trusted libvlc_media_option_t = iota + 0x2
	libvlc_media_option_unique                        = 0x100
)

type libvlc_track_type_t int32

const (
	libvlc_track_unknown libvlc_track_type_t = iota - 1 // -1
	libvlc_track_audio                                  //0
	libvlc_track_video                                  //1
	libvlc_track_text                                   //2
)

type (
	libvlc_media_stats_t struct {
		I_read_bytes          int     // cint
		F_input_bitrate       float32 // cfloat
		I_demux_read_bytes    int     // cint
		F_demux_bitrate       float32 // cfloat
		I_demux_corrupted     int     // cint
		I_demux_discontinuity int     // cint
		I_decoded_video       int     // cint
		I_decoded_audio       int     // cint
		I_displayed_pictures  int     // cint
		I_lost_pictures       int     // cint
		I_played_abuffers     int     // cint
		I_lost_abuffers       int     // cint
		I_sent_packets        int     // cint
		I_sent_bytes          int     // cint
		F_send_bitrate        float32 // cfloat
	}

	libvlc_media_track_info_t struct {
		i_codec   uint32 //: uint32_t;
		i_id      int    //: cint;
		i_type    libvlc_track_type_t
		i_profile int // : cint;
		i_level   int //: cint;
		u         struct {
			//0
			audio struct {
				i_channels uint //: cunsigned;
				i_rate     uint //: cunsigned;
			}

			//1
			//video struct {
			//	i_height uint //: cunsigned;
			//	i_width  uint //: cunsigned;
			//}
		}
	}

	libvlc_track_description_t struct {
		i_id     int     //: cint;
		psz_name uintptr //: ^cchar;
		p_next   *libvlc_track_description_t
	}

	libvlc_audio_output_t struct {
		psz_name        uintptr //: ^cchar;
		psz_description uintptr //: ^cchar;
		p_next          *libvlc_audio_output_t
	}

	libvlc_rectangle_t struct {
		top    int //: cint;
		left   int //: cint;
		bottom int //: cint;
		right  int //: cint;
	}
)

type libvlc_video_marquee_option_t int32

const (
	libvlc_marquee_Enable libvlc_video_marquee_option_t = iota + 0
	libvlc_marquee_Text
	libvlc_marquee_Color
	libvlc_marquee_Opacity
	libvlc_marquee_Position
	libvlc_marquee_Refresh
	libvlc_marquee_Size
	libvlc_marquee_Timeout
	libvlc_marquee_X
	libvlc_marquee_Y
)

type libvlc_navigate_mode_t int32

const (
	libvlc_navigate_activate libvlc_navigate_mode_t = iota + 0
	libvlc_navigate_up
	libvlc_navigate_down
	libvlc_navigate_left
	libvlc_navigate_right
)

type libvlc_video_logo_option_t int32

const (
	libvlc_logo_enable libvlc_video_logo_option_t = iota + 0
	libvlc_logo_file
	libvlc_logo_x
	libvlc_logo_y
	libvlc_logo_delay
	libvlc_logo_repeat
	libvlc_logo_opacity
	libvlc_logo_position
)

type libvlc_video_adjust_option_t int32

const (
	libvlc_adjust_Enable libvlc_video_adjust_option_t = iota + 0
	libvlc_adjust_Contrast
	libvlc_adjust_Brightness
	libvlc_adjust_Hue
	libvlc_adjust_Saturation
	libvlc_adjust_Gamma
)

type libvlc_audio_output_device_types_t int32

const (
	libvlc_AudioOutputDevice_Error  libvlc_audio_output_device_types_t = iota - 1
	libvlc_AudioOutputDevice_Mono                                      = 1
	libvlc_AudioOutputDevice_Stereo                                    = 2
	libvlc_AudioOutputDevice_2F2R                                      = 4
	libvlc_AudioOutputDevice_3F2R                                      = 5
	libvlc_AudioOutputDevice_5_1                                       = 6
	libvlc_AudioOutputDevice_6_1                                       = 7
	libvlc_AudioOutputDevice_7_1                                       = 8
	libvlc_AudioOutputDevice_SPDIF                                     = 10
)

type libvlc_audio_output_channel_t int32

const (
	libvlc_AudioChannel_Error   libvlc_audio_output_channel_t = iota - 1
	libvlc_AudioChannel_Stereo                                = 1
	libvlc_AudioChannel_RStereo                               = 2
	libvlc_AudioChannel_Left                                  = 3
	libvlc_AudioChannel_Right                                 = 4
	libvlc_AudioChannel_Dolbys                                = 5
)

type libvlc_playback_mode_t int32

const (
	libvlc_playback_mode_default libvlc_playback_mode_t = iota + 0
	libvlc_playback_mode_loop
	libvlc_playback_mode_repeat
)

type libvlc_event_e int32

const (
	libvlc_MediaMetaChanged libvlc_event_e = iota + 0
	libvlc_MediaSubItemAdded
	libvlc_MediaDurationChanged
	libvlc_MediaParsedChanged
	libvlc_MediaFreed
	libvlc_MediaStateChanged
	libvlc_MediaPlayerMediaChanged = 0x100
	libvlc_MediaPlayerNothingSpecial
	libvlc_MediaPlayerOpening
	libvlc_MediaPlayerBuffering
	libvlc_MediaPlayerPlaying
	libvlc_MediaPlayerPaused
	libvlc_MediaPlayerStopped
	libvlc_MediaPlayerForward
	libvlc_MediaPlayerBackward
	libvlc_MediaPlayerEndReached
	libvlc_MediaPlayerEncounteredError
	libvlc_MediaPlayerTimeChanged
	libvlc_MediaPlayerPositionChanged
	libvlc_MediaPlayerSeekableChanged
	libvlc_MediaPlayerPausableChanged
	libvlc_MediaPlayerTitleChanged
	libvlc_MediaPlayerSnapshotTaken
	libvlc_MediaPlayerLengthChanged
	libvlc_MediaPlayerVout
	libvlc_MediaListItemAdded = 0x200
	libvlc_MediaListWillAddItem
	libvlc_MediaListItemDeleted
	libvlc_MediaListWillDeleteItem
	libvlc_MediaListViewItemAdded = 0x300
	libvlc_MediaListViewWillAddItem
	libvlc_MediaListViewItemDeleted
	libvlc_MediaListViewWillDeleteItem
	libvlc_MediaListPlayerPlayed = 0x400
	libvlc_MediaListPlayerNextItemSet
	libvlc_MediaListPlayerStopped
	libvlc_MediaDiscovererStarted = 0x500
	libvlc_MediaDiscovererEnded
	libvlc_VlmMediaAdded = 0x600
	libvlc_VlmMediaRemoved
	libvlc_VlmMediaChanged
	libvlc_VlmMediaInstanceStarted
	libvlc_VlmMediaInstanceStopped
	libvlc_VlmMediaInstanceStatusInit
	libvlc_VlmMediaInstanceStatusOpening
	libvlc_VlmMediaInstanceStatusPlaying
	libvlc_VlmMediaInstanceStatusPause
	libvlc_VlmMediaInstanceStatusEnd
	libvlc_VlmMediaInstanceStatusError
)

type libvlc_event_t struct {
	_type int     //: cint;
	p_obj uintptr //: pointer;
}

/*
  libvlc_event_t struct {
    _type : cint;
    p_obj : pointer;
      case longint of
        0 : ( media_meta_changed : record
            meta_type : libvlc_meta_t;
          end );
        1 : ( media_subitem_added : record
            new_child : ^libvlc_media_t;
          end );
        2 : ( media_duration_changed : record
            new_duration : int64_t;
          end );
        3 : ( media_parsed_changed : record
            new_status : cint;
          end );
        4 : ( media_freed : record
            md : ^libvlc_media_t;
          end );
        5 : ( media_state_changed : record
            new_state : libvlc_state_t;
          end );
        6 : ( media_player_buffering : record
            new_cache : cfloat;
          end );
        7 : ( media_player_position_changed : record
            new_position : cfloat;
          end );
        8 : ( media_player_time_changed : record
            new_time : libvlc_time_t;
          end );
        9 : ( media_player_title_changed : record
            new_title : cint;
          end );
        10 : ( media_player_seekable_changed : record
            new_seekable : cint;
          end );
        11 : ( media_player_pausable_changed : record
            new_pausable : cint;
          end );
        12 : ( media_player_vout : record
            new_count : cint;
          end );
        13 : ( media_list_item_added : record
            item : ^libvlc_media_t;
            index : cint;
          end );
        14 : ( media_list_will_add_item : record
            item : ^libvlc_media_t;
            index : cint;
          end );
        15 : ( media_list_item_deleted : record
            item : ^libvlc_media_t;
            index : cint;
          end );
        16 : ( media_list_will_delete_item : record
            item : ^libvlc_media_t;
            index : cint;
          end );
        17 : ( media_list_player_next_item_set : record
            item : ^libvlc_media_t;
          end );
        18 : ( media_player_snapshot_taken : record
            psz_filename : ^cchar;
          end );
        19 : ( media_player_length_changed : record
            new_length : libvlc_time_t;
          end );
        20 : ( vlm_media_event : record
            psz_media_name : ^cchar;
            psz_instance_name : ^cchar;
          end );
        21 : ( media_player_media_changed : record
            new_media : ^libvlc_media_t;
          end );
    }

*/

type (
	pplibvlc_media_track_info_t *plibvlc_media_track_info_t
	cbtype1                     uintptr // = procedure (_para1:pointer); cdecl;
)

//------------------------- 函数导入-------------------------------------------------------------------------------------
var (
	libvlcdll                                        = dylib.NewLazyDLL(getDLLName())
	_libvlc_media_player_new                         = libvlcdll.NewProc("libvlc_media_player_new")
	_libvlc_media_player_new_from_media              = libvlcdll.NewProc("libvlc_media_player_new_from_media")
	_libvlc_media_player_release                     = libvlcdll.NewProc("libvlc_media_player_release")
	_libvlc_media_player_retain                      = libvlcdll.NewProc("libvlc_media_player_retain")
	_libvlc_media_player_set_media                   = libvlcdll.NewProc("libvlc_media_player_set_media")
	_libvlc_media_player_get_media                   = libvlcdll.NewProc("libvlc_media_player_get_media")
	_libvlc_media_player_event_manager               = libvlcdll.NewProc("libvlc_media_player_event_manager")
	_libvlc_media_player_is_playing                  = libvlcdll.NewProc("libvlc_media_player_is_playing")
	_libvlc_media_player_play                        = libvlcdll.NewProc("libvlc_media_player_play")
	_libvlc_media_player_set_pause                   = libvlcdll.NewProc("libvlc_media_player_set_pause")
	_libvlc_media_player_pause                       = libvlcdll.NewProc("libvlc_media_player_pause")
	_libvlc_media_player_stop                        = libvlcdll.NewProc("libvlc_media_player_stop")
	_libvlc_media_new_location                       = libvlcdll.NewProc("libvlc_media_new_location")
	_libvlc_media_new_path                           = libvlcdll.NewProc("libvlc_media_new_path")
	_libvlc_media_new_fd                             = libvlcdll.NewProc("libvlc_media_new_fd")
	_libvlc_media_new_as_node                        = libvlcdll.NewProc("libvlc_media_new_as_node")
	_libvlc_media_add_option                         = libvlcdll.NewProc("libvlc_media_add_option")
	_libvlc_media_add_option_flag                    = libvlcdll.NewProc("libvlc_media_add_option_flag")
	_libvlc_media_retain                             = libvlcdll.NewProc("libvlc_media_retain")
	_libvlc_media_release                            = libvlcdll.NewProc("libvlc_media_release")
	_libvlc_media_get_mrl                            = libvlcdll.NewProc("libvlc_media_get_mrl")
	_libvlc_media_duplicate                          = libvlcdll.NewProc("libvlc_media_duplicate")
	_libvlc_media_get_meta                           = libvlcdll.NewProc("libvlc_media_get_meta")
	_libvlc_media_set_meta                           = libvlcdll.NewProc("libvlc_media_set_meta")
	_libvlc_media_save_meta                          = libvlcdll.NewProc("libvlc_media_save_meta")
	_libvlc_media_get_state                          = libvlcdll.NewProc("libvlc_media_get_state")
	_libvlc_media_get_stats                          = libvlcdll.NewProc("libvlc_media_get_stats")
	_libvlc_media_subitems                           = libvlcdll.NewProc("libvlc_media_subitems")
	_libvlc_media_event_manager                      = libvlcdll.NewProc("libvlc_media_event_manager")
	_libvlc_media_get_duration                       = libvlcdll.NewProc("libvlc_media_get_duration")
	_libvlc_media_parse                              = libvlcdll.NewProc("libvlc_media_parse")
	_libvlc_media_parse_async                        = libvlcdll.NewProc("libvlc_media_parse_async")
	_libvlc_media_is_parsed                          = libvlcdll.NewProc("libvlc_media_is_parsed")
	_libvlc_media_set_user_data                      = libvlcdll.NewProc("libvlc_media_set_user_data")
	_libvlc_media_get_user_data                      = libvlcdll.NewProc("libvlc_media_get_user_data")
	_libvlc_media_get_tracks_info                    = libvlcdll.NewProc("libvlc_media_get_tracks_info")
	_libvlc_module_description_list_release          = libvlcdll.NewProc("libvlc_module_description_list_release")
	_libvlc_audio_filter_list_get                    = libvlcdll.NewProc("libvlc_audio_filter_list_get")
	_libvlc_video_filter_list_get                    = libvlcdll.NewProc("libvlc_video_filter_list_get")
	_libvlc_clock                                    = libvlcdll.NewProc("libvlc_clock")
	_libvlc_errmsg                                   = libvlcdll.NewProc("libvlc_errmsg")
	_libvlc_clearerr                                 = libvlcdll.NewProc("libvlc_clearerr")
	_libvlc_printerr                                 = libvlcdll.NewProc("libvlc_printerr")
	_libvlc_new                                      = libvlcdll.NewProc("libvlc_new")
	_libvlc_release                                  = libvlcdll.NewProc("libvlc_release")
	_libvlc_retain                                   = libvlcdll.NewProc("libvlc_retain")
	_libvlc_add_intf                                 = libvlcdll.NewProc("libvlc_add_intf")
	_libvlc_set_exit_handler                         = libvlcdll.NewProc("libvlc_set_exit_handler")
	_libvlc_wait                                     = libvlcdll.NewProc("libvlc_wait")
	_libvlc_set_user_agent                           = libvlcdll.NewProc("libvlc_set_user_agent")
	_libvlc_get_version                              = libvlcdll.NewProc("libvlc_get_version")
	_libvlc_get_compiler                             = libvlcdll.NewProc("libvlc_get_compiler")
	_libvlc_get_changeset                            = libvlcdll.NewProc("libvlc_get_changeset")
	_libvlc_free                                     = libvlcdll.NewProc("libvlc_free")
	_libvlc_event_attach                             = libvlcdll.NewProc("libvlc_event_attach")
	_libvlc_event_detach                             = libvlcdll.NewProc("libvlc_event_detach")
	_libvlc_event_type_name                          = libvlcdll.NewProc("libvlc_event_type_name")
	_libvlc_get_log_verbosity                        = libvlcdll.NewProc("libvlc_get_log_verbosity")
	_libvlc_set_log_verbosity                        = libvlcdll.NewProc("libvlc_set_log_verbosity")
	_libvlc_log_open                                 = libvlcdll.NewProc("libvlc_log_open")
	_libvlc_log_close                                = libvlcdll.NewProc("libvlc_log_close")
	_libvlc_log_count                                = libvlcdll.NewProc("libvlc_log_count")
	_libvlc_log_clear                                = libvlcdll.NewProc("libvlc_log_clear")
	_libvlc_log_get_iterator                         = libvlcdll.NewProc("libvlc_log_get_iterator")
	_libvlc_log_iterator_free                        = libvlcdll.NewProc("libvlc_log_iterator_free")
	_libvlc_log_iterator_has_next                    = libvlcdll.NewProc("libvlc_log_iterator_has_next")
	_libvlc_log_iterator_next                        = libvlcdll.NewProc("libvlc_log_iterator_next")
	_libvlc_audio_output_list_get                    = libvlcdll.NewProc("libvlc_audio_output_list_get")
	_libvlc_audio_output_list_release                = libvlcdll.NewProc("libvlc_audio_output_list_release")
	_libvlc_audio_output_set                         = libvlcdll.NewProc("libvlc_audio_output_set")
	_libvlc_audio_output_device_count                = libvlcdll.NewProc("libvlc_audio_output_device_count")
	_libvlc_audio_output_device_longname             = libvlcdll.NewProc("libvlc_audio_output_device_longname")
	_libvlc_audio_output_device_id                   = libvlcdll.NewProc("libvlc_audio_output_device_id")
	_libvlc_audio_output_device_set                  = libvlcdll.NewProc("libvlc_audio_output_device_set")
	_libvlc_audio_output_get_device_type             = libvlcdll.NewProc("libvlc_audio_output_get_device_type")
	_libvlc_audio_output_set_device_type             = libvlcdll.NewProc("libvlc_audio_output_set_device_type")
	_libvlc_audio_toggle_mute                        = libvlcdll.NewProc("libvlc_audio_toggle_mute")
	_libvlc_audio_get_mute                           = libvlcdll.NewProc("libvlc_audio_get_mute")
	_libvlc_audio_set_mute                           = libvlcdll.NewProc("libvlc_audio_set_mute")
	_libvlc_audio_get_volume                         = libvlcdll.NewProc("libvlc_audio_get_volume")
	_libvlc_audio_set_volume                         = libvlcdll.NewProc("libvlc_audio_set_volume")
	_libvlc_audio_get_track_count                    = libvlcdll.NewProc("libvlc_audio_get_track_count")
	_libvlc_audio_get_track_description              = libvlcdll.NewProc("libvlc_audio_get_track_description")
	_libvlc_audio_get_track                          = libvlcdll.NewProc("libvlc_audio_get_track")
	_libvlc_audio_set_track                          = libvlcdll.NewProc("libvlc_audio_set_track")
	_libvlc_audio_get_channel                        = libvlcdll.NewProc("libvlc_audio_get_channel")
	_libvlc_audio_set_channel                        = libvlcdll.NewProc("libvlc_audio_set_channel")
	_libvlc_audio_get_delay                          = libvlcdll.NewProc("libvlc_audio_get_delay")
	_libvlc_audio_set_delay                          = libvlcdll.NewProc("libvlc_audio_set_delay")
	_libvlc_media_list_new                           = libvlcdll.NewProc("libvlc_media_list_new")
	_libvlc_media_list_release                       = libvlcdll.NewProc("libvlc_media_list_release")
	_libvlc_media_list_retain                        = libvlcdll.NewProc("libvlc_media_list_retain")
	_libvlc_media_list_add_file_content              = libvlcdll.NewProc("libvlc_media_list_add_file_content")
	_libvlc_media_list_set_media                     = libvlcdll.NewProc("libvlc_media_list_set_media")
	_libvlc_media_list_media                         = libvlcdll.NewProc("libvlc_media_list_media")
	_libvlc_media_list_add_media                     = libvlcdll.NewProc("libvlc_media_list_add_media")
	_libvlc_media_list_insert_media                  = libvlcdll.NewProc("libvlc_media_list_insert_media")
	_libvlc_media_list_remove_index                  = libvlcdll.NewProc("libvlc_media_list_remove_index")
	_libvlc_media_list_count                         = libvlcdll.NewProc("libvlc_media_list_count")
	_libvlc_media_list_item_at_index                 = libvlcdll.NewProc("libvlc_media_list_item_at_index")
	_libvlc_media_list_index_of_item                 = libvlcdll.NewProc("libvlc_media_list_index_of_item")
	_libvlc_media_list_is_readonly                   = libvlcdll.NewProc("libvlc_media_list_is_readonly")
	_libvlc_media_list_lock                          = libvlcdll.NewProc("libvlc_media_list_lock")
	_libvlc_media_list_unlock                        = libvlcdll.NewProc("libvlc_media_list_unlock")
	_libvlc_media_list_event_manager                 = libvlcdll.NewProc("libvlc_media_list_event_manager")
	_libvlc_media_list_player_new                    = libvlcdll.NewProc("libvlc_media_list_player_new")
	_libvlc_media_list_player_release                = libvlcdll.NewProc("libvlc_media_list_player_release")
	_libvlc_media_list_player_retain                 = libvlcdll.NewProc("libvlc_media_list_player_retain")
	_libvlc_media_list_player_event_manager          = libvlcdll.NewProc("libvlc_media_list_player_event_manager")
	_libvlc_media_list_player_set_media_player       = libvlcdll.NewProc("libvlc_media_list_player_set_media_player")
	_libvlc_media_list_player_set_media_list         = libvlcdll.NewProc("libvlc_media_list_player_set_media_list")
	_libvlc_media_list_player_play                   = libvlcdll.NewProc("libvlc_media_list_player_play")
	_libvlc_media_list_player_pause                  = libvlcdll.NewProc("libvlc_media_list_player_pause")
	_libvlc_media_list_player_is_playing             = libvlcdll.NewProc("libvlc_media_list_player_is_playing")
	_libvlc_media_list_player_get_state              = libvlcdll.NewProc("libvlc_media_list_player_get_state")
	_libvlc_media_list_player_play_item_at_index     = libvlcdll.NewProc("libvlc_media_list_player_play_item_at_index")
	_libvlc_media_list_player_play_item              = libvlcdll.NewProc("libvlc_media_list_player_play_item")
	_libvlc_media_list_player_stop                   = libvlcdll.NewProc("libvlc_media_list_player_stop")
	_libvlc_media_list_player_next                   = libvlcdll.NewProc("libvlc_media_list_player_next")
	_libvlc_media_list_player_previous               = libvlcdll.NewProc("libvlc_media_list_player_previous")
	_libvlc_media_list_player_set_playback_mode      = libvlcdll.NewProc("libvlc_media_list_player_set_playback_mode")
	_libvlc_media_library_new                        = libvlcdll.NewProc("libvlc_media_library_new")
	_libvlc_media_library_release                    = libvlcdll.NewProc("libvlc_media_library_release")
	_libvlc_media_library_retain                     = libvlcdll.NewProc("libvlc_media_library_retain")
	_libvlc_media_library_load                       = libvlcdll.NewProc("libvlc_media_library_load")
	_libvlc_media_library_media_list                 = libvlcdll.NewProc("libvlc_media_library_media_list")
	_libvlc_video_get_adjust_int                     = libvlcdll.NewProc("libvlc_video_get_adjust_int")
	_libvlc_video_set_adjust_int                     = libvlcdll.NewProc("libvlc_video_set_adjust_int")
	_libvlc_video_get_adjust_float                   = libvlcdll.NewProc("libvlc_video_get_adjust_float")
	_libvlc_video_set_adjust_float                   = libvlcdll.NewProc("libvlc_video_set_adjust_float")
	_libvlc_video_get_logo_int                       = libvlcdll.NewProc("libvlc_video_get_logo_int")
	_libvlc_video_set_logo_int                       = libvlcdll.NewProc("libvlc_video_set_logo_int")
	_libvlc_video_set_logo_string                    = libvlcdll.NewProc("libvlc_video_set_logo_string")
	_libvlc_audio_set_format_callbacks               = libvlcdll.NewProc("libvlc_audio_set_format_callbacks")
	_libvlc_audio_set_format                         = libvlcdll.NewProc("libvlc_audio_set_format")
	_libvlc_media_player_get_length                  = libvlcdll.NewProc("libvlc_media_player_get_length")
	_libvlc_media_player_get_time                    = libvlcdll.NewProc("libvlc_media_player_get_time")
	_libvlc_media_player_set_time                    = libvlcdll.NewProc("libvlc_media_player_set_time")
	_libvlc_media_player_get_position                = libvlcdll.NewProc("libvlc_media_player_get_position")
	_libvlc_media_player_set_position                = libvlcdll.NewProc("libvlc_media_player_set_position")
	_libvlc_media_player_set_chapter                 = libvlcdll.NewProc("libvlc_media_player_set_chapter")
	_libvlc_media_player_get_chapter                 = libvlcdll.NewProc("libvlc_media_player_get_chapter")
	_libvlc_media_player_get_chapter_count           = libvlcdll.NewProc("libvlc_media_player_get_chapter_count")
	_libvlc_media_player_will_play                   = libvlcdll.NewProc("libvlc_media_player_will_play")
	_libvlc_media_player_get_chapter_count_for_title = libvlcdll.NewProc("libvlc_media_player_get_chapter_count_for_title")
	_libvlc_media_player_set_title                   = libvlcdll.NewProc("libvlc_media_player_set_title")
	_libvlc_media_player_get_title                   = libvlcdll.NewProc("libvlc_media_player_get_title")
	_libvlc_media_player_get_title_count             = libvlcdll.NewProc("libvlc_media_player_get_title_count")
	_libvlc_media_player_previous_chapter            = libvlcdll.NewProc("libvlc_media_player_previous_chapter")
	_libvlc_media_player_next_chapter                = libvlcdll.NewProc("libvlc_media_player_next_chapter")
	_libvlc_media_player_get_rate                    = libvlcdll.NewProc("libvlc_media_player_get_rate")
	_libvlc_media_player_set_rate                    = libvlcdll.NewProc("libvlc_media_player_set_rate")
	_libvlc_media_player_get_state                   = libvlcdll.NewProc("libvlc_media_player_get_state")
	_libvlc_media_player_get_fps                     = libvlcdll.NewProc("libvlc_media_player_get_fps")
	_libvlc_media_player_has_vout                    = libvlcdll.NewProc("libvlc_media_player_has_vout")
	_libvlc_media_player_is_seekable                 = libvlcdll.NewProc("libvlc_media_player_is_seekable")
	_libvlc_media_player_can_pause                   = libvlcdll.NewProc("libvlc_media_player_can_pause")
	_libvlc_media_player_next_frame                  = libvlcdll.NewProc("libvlc_media_player_next_frame")
	_libvlc_media_player_navigate                    = libvlcdll.NewProc("libvlc_media_player_navigate")
	_libvlc_track_description_list_release           = libvlcdll.NewProc("libvlc_track_description_list_release")
	_libvlc_track_description_release                = libvlcdll.NewProc("libvlc_track_description_release")
	_libvlc_toggle_fullscreen                        = libvlcdll.NewProc("libvlc_toggle_fullscreen")
	_libvlc_set_fullscreen                           = libvlcdll.NewProc("libvlc_set_fullscreen")
	_libvlc_get_fullscreen                           = libvlcdll.NewProc("libvlc_get_fullscreen")
	_libvlc_video_set_key_input                      = libvlcdll.NewProc("libvlc_video_set_key_input")
	_libvlc_video_set_mouse_input                    = libvlcdll.NewProc("libvlc_video_set_mouse_input")
	_libvlc_video_get_size                           = libvlcdll.NewProc("libvlc_video_get_size")
	_libvlc_video_get_height                         = libvlcdll.NewProc("libvlc_video_get_height")
	_libvlc_video_get_width                          = libvlcdll.NewProc("libvlc_video_get_width")
	_libvlc_video_get_cursor                         = libvlcdll.NewProc("libvlc_video_get_cursor")
	_libvlc_video_get_scale                          = libvlcdll.NewProc("libvlc_video_get_scale")
	_libvlc_video_set_scale                          = libvlcdll.NewProc("libvlc_video_set_scale")
	_libvlc_video_get_aspect_ratio                   = libvlcdll.NewProc("libvlc_video_get_aspect_ratio")
	_libvlc_video_set_aspect_ratio                   = libvlcdll.NewProc("libvlc_video_set_aspect_ratio")
	_libvlc_video_get_spu                            = libvlcdll.NewProc("libvlc_video_get_spu")
	_libvlc_video_get_spu_count                      = libvlcdll.NewProc("libvlc_video_get_spu_count")
	_libvlc_video_get_spu_description                = libvlcdll.NewProc("libvlc_video_get_spu_description")
	_libvlc_video_set_spu                            = libvlcdll.NewProc("libvlc_video_set_spu")
	_libvlc_video_set_subtitle_file                  = libvlcdll.NewProc("libvlc_video_set_subtitle_file")
	_libvlc_video_get_spu_delay                      = libvlcdll.NewProc("libvlc_video_get_spu_delay")
	_libvlc_video_set_spu_delay                      = libvlcdll.NewProc("libvlc_video_set_spu_delay")
	_libvlc_video_get_title_description              = libvlcdll.NewProc("libvlc_video_get_title_description")
	_libvlc_video_get_chapter_description            = libvlcdll.NewProc("libvlc_video_get_chapter_description")
	_libvlc_video_get_crop_geometry                  = libvlcdll.NewProc("libvlc_video_get_crop_geometry")
	_libvlc_video_set_crop_geometry                  = libvlcdll.NewProc("libvlc_video_set_crop_geometry")
	_libvlc_video_get_teletext                       = libvlcdll.NewProc("libvlc_video_get_teletext")
	_libvlc_video_set_teletext                       = libvlcdll.NewProc("libvlc_video_set_teletext")
	_libvlc_toggle_teletext                          = libvlcdll.NewProc("libvlc_toggle_teletext")
	_libvlc_video_get_track_count                    = libvlcdll.NewProc("libvlc_video_get_track_count")
	_libvlc_video_get_track_description              = libvlcdll.NewProc("libvlc_video_get_track_description")
	_libvlc_video_get_track                          = libvlcdll.NewProc("libvlc_video_get_track")
	_libvlc_video_set_track                          = libvlcdll.NewProc("libvlc_video_set_track")
	_libvlc_video_take_snapshot                      = libvlcdll.NewProc("libvlc_video_take_snapshot")
	_libvlc_video_set_deinterlace                    = libvlcdll.NewProc("libvlc_video_set_deinterlace")
	_libvlc_video_get_marquee_int                    = libvlcdll.NewProc("libvlc_video_get_marquee_int")
	_libvlc_video_get_marquee_string                 = libvlcdll.NewProc("libvlc_video_get_marquee_string")
	_libvlc_video_set_marquee_int                    = libvlcdll.NewProc("libvlc_video_set_marquee_int")
	_libvlc_video_set_marquee_string                 = libvlcdll.NewProc("libvlc_video_set_marquee_string")
	_libvlc_audio_set_callbacks                      = libvlcdll.NewProc("libvlc_audio_set_callbacks")
	_libvlc_audio_set_volume_callback                = libvlcdll.NewProc("libvlc_audio_set_volume_callback")
	_libvlc_video_set_callbacks                      = libvlcdll.NewProc("libvlc_video_set_callbacks")
	_libvlc_video_set_format                         = libvlcdll.NewProc("libvlc_video_set_format")
	_libvlc_video_set_format_callbacks               = libvlcdll.NewProc("libvlc_video_set_format_callbacks")
	_libvlc_media_player_set_nsobject                = libvlcdll.NewProc("libvlc_media_player_set_nsobject")
	_libvlc_media_player_get_nsobject                = libvlcdll.NewProc("libvlc_media_player_get_nsobject")
	_libvlc_media_player_set_agl                     = libvlcdll.NewProc("libvlc_media_player_set_agl")
	_libvlc_media_player_get_agl                     = libvlcdll.NewProc("libvlc_media_player_get_agl")
	_libvlc_media_player_set_xwindow                 = libvlcdll.NewProc("libvlc_media_player_set_xwindow")
	_libvlc_media_player_get_xwindow                 = libvlcdll.NewProc("libvlc_media_player_get_xwindow")
	_libvlc_media_player_set_hwnd                    = libvlcdll.NewProc("libvlc_media_player_set_hwnd")
	_libvlc_media_player_get_hwnd                    = libvlcdll.NewProc("libvlc_media_player_get_hwnd")
	_libvlc_media_discoverer_new_from_name           = libvlcdll.NewProc("libvlc_media_discoverer_new_from_name")
	_libvlc_media_discoverer_release                 = libvlcdll.NewProc("libvlc_media_discoverer_release")
	_libvlc_media_discoverer_localized_name          = libvlcdll.NewProc("libvlc_media_discoverer_localized_name")
	_libvlc_media_discoverer_media_list              = libvlcdll.NewProc("libvlc_media_discoverer_media_list")
	_libvlc_media_discoverer_event_manager           = libvlcdll.NewProc("libvlc_media_discoverer_event_manager")
	_libvlc_media_discoverer_is_running              = libvlcdll.NewProc("libvlc_media_discoverer_is_running")
	_libvlc_vlm_release                              = libvlcdll.NewProc("libvlc_vlm_release")
	_libvlc_vlm_add_broadcast                        = libvlcdll.NewProc("libvlc_vlm_add_broadcast")
	_libvlc_vlm_add_vod                              = libvlcdll.NewProc("libvlc_vlm_add_vod")
	_libvlc_vlm_del_media                            = libvlcdll.NewProc("libvlc_vlm_del_media")
	_libvlc_vlm_set_enabled                          = libvlcdll.NewProc("libvlc_vlm_set_enabled")
	_libvlc_vlm_set_output                           = libvlcdll.NewProc("libvlc_vlm_set_output")
	_libvlc_vlm_set_input                            = libvlcdll.NewProc("libvlc_vlm_set_input")
	_libvlc_vlm_add_input                            = libvlcdll.NewProc("libvlc_vlm_add_input")
	_libvlc_vlm_set_loop                             = libvlcdll.NewProc("libvlc_vlm_set_loop")
	_libvlc_vlm_set_mux                              = libvlcdll.NewProc("libvlc_vlm_set_mux")
	_libvlc_vlm_change_media                         = libvlcdll.NewProc("libvlc_vlm_change_media")
	_libvlc_vlm_play_media                           = libvlcdll.NewProc("libvlc_vlm_play_media")
	_libvlc_vlm_stop_media                           = libvlcdll.NewProc("libvlc_vlm_stop_media")
	_libvlc_vlm_pause_media                          = libvlcdll.NewProc("libvlc_vlm_pause_media")
	_libvlc_vlm_seek_media                           = libvlcdll.NewProc("libvlc_vlm_seek_media")
	_libvlc_vlm_show_media                           = libvlcdll.NewProc("libvlc_vlm_show_media")
	_libvlc_vlm_get_media_instance_position          = libvlcdll.NewProc("libvlc_vlm_get_media_instance_position")
	_libvlc_vlm_get_media_instance_time              = libvlcdll.NewProc("libvlc_vlm_get_media_instance_time")
	_libvlc_vlm_get_media_instance_length            = libvlcdll.NewProc("libvlc_vlm_get_media_instance_length")
	_libvlc_vlm_get_media_instance_rate              = libvlcdll.NewProc("libvlc_vlm_get_media_instance_rate")
	_libvlc_vlm_get_event_manager                    = libvlcdll.NewProc("libvlc_vlm_get_event_manager")
	_libvlc_playlist_play                            = libvlcdll.NewProc("libvlc_playlist_play")
)

func toCStr(s string) uintptr {
	return uintptr(unsafe.Pointer(&([]byte(s + "\x00")[0])))
}

func strLen(sptr uintptr) int {
	if sptr <= 0 {
		return 0
	}
	i := 0
	for *(*byte)(unsafe.Pointer(sptr + uintptr(i))) != '0' {
		i++
	}
	return i
}

func toGoBool(i int) bool {
	return i == 1
}

func getDLLName() string {

	switch runtime.GOOS {
	case "windows":
		return "libvlc.dll"
	case "linux":
		return "libvlc.so.5"
	case "darwin":
		return "./lib/libvlc.dylib"
	}
	return ""
}

func libvlc_media_player_new(p_libvlc_instance plibvlc_instance_t) plibvlc_media_player_t {
	r, _, _ := _libvlc_media_player_new.Call(uintptr(p_libvlc_instance))
	return plibvlc_media_player_t(r)
}

func libvlc_media_player_new_from_media(p_md plibvlc_media_t) plibvlc_media_player_t {
	r, _, _ := _libvlc_media_player_new_from_media.Call(uintptr(p_md))
	return plibvlc_media_player_t(r)
}

func libvlc_media_player_release(p_mi plibvlc_media_player_t) {
	_libvlc_media_player_release.Call(uintptr(p_mi))
}

func libvlc_media_player_retain(p_mi plibvlc_media_player_t) {
	_libvlc_media_player_retain.Call(uintptr(p_mi))
}

func libvlc_media_player_set_media(p_mi plibvlc_media_player_t, p_md plibvlc_media_t) {
	_libvlc_media_player_set_media.Call(uintptr(p_mi), uintptr(p_md))
}

func libvlc_media_player_get_media(p_mi plibvlc_media_player_t) plibvlc_media_t {
	r, _, _ := _libvlc_media_player_get_media.Call(uintptr(p_mi))
	return plibvlc_media_t(r)
}

func libvlc_media_player_event_manager(p_mi plibvlc_media_player_t) plibvlc_event_manager_t {
	r, _, _ := _libvlc_media_player_event_manager.Call(uintptr(p_mi))
	return plibvlc_event_manager_t(r)
}

func libvlc_media_player_is_playing(p_mi plibvlc_media_player_t) int {
	r, _, _ := _libvlc_media_player_is_playing.Call(uintptr(p_mi))
	return int(r)
}

func libvlc_media_player_play(p_mi plibvlc_media_player_t) int {
	r, _, _ := _libvlc_media_player_play.Call(uintptr(p_mi))
	return int(r)
}

func libvlc_media_player_set_pause(mp plibvlc_media_player_t, do_pause int) {
	_libvlc_media_player_set_pause.Call(uintptr(mp), uintptr(do_pause))
}

func libvlc_media_player_pause(p_mi plibvlc_media_player_t) {
	_libvlc_media_player_pause.Call(uintptr(p_mi))
}

func libvlc_media_player_stop(p_mi plibvlc_media_player_t) {
	_libvlc_media_player_stop.Call(uintptr(p_mi))
}

func libvlc_media_new_location(p_instance plibvlc_instance_t, psz_mrl string) plibvlc_media_t {
	r, _, _ := _libvlc_media_new_location.Call(uintptr(p_instance), toCStr(psz_mrl))
	return plibvlc_media_t(r)
}

func libvlc_media_new_path(p_instance plibvlc_instance_t, path string) plibvlc_media_t {
	r, _, _ := _libvlc_media_new_path.Call(uintptr(p_instance), toCStr(path))
	return plibvlc_media_t(r)
}

func libvlc_media_new_fd(p_instance plibvlc_instance_t, fd int) plibvlc_media_t {
	r, _, _ := _libvlc_media_new_fd.Call(uintptr(p_instance), uintptr(fd))
	return plibvlc_media_t(r)
}

func libvlc_media_new_as_node(p_instance plibvlc_instance_t, psz_name string) plibvlc_media_t {
	r, _, _ := _libvlc_media_new_as_node.Call(uintptr(p_instance), toCStr(psz_name))
	return plibvlc_media_t(r)
}

func libvlc_media_add_option(p_md plibvlc_media_t, ppsz_options string) {
	_libvlc_media_add_option.Call(uintptr(p_md), toCStr(ppsz_options))
}

func libvlc_media_add_option_flag(p_md plibvlc_media_t, ppsz_options string, i_flags uint) {
	_libvlc_media_add_option_flag.Call(uintptr(p_md), toCStr(ppsz_options), uintptr(i_flags))
}

func libvlc_media_retain(p_md plibvlc_media_t) {
	_libvlc_media_retain.Call(uintptr(p_md))
}

func libvlc_media_release(p_md plibvlc_media_t) {
	_libvlc_media_release.Call(uintptr(p_md))
}

func libvlc_media_get_mrl(p_md plibvlc_media_t) string {
	r, _, _ := _libvlc_media_get_mrl.Call(uintptr(p_md))
	return string(r)
}

func libvlc_media_duplicate(p_md plibvlc_media_t) plibvlc_media_t {
	r, _, _ := _libvlc_media_duplicate.Call(uintptr(p_md))
	return plibvlc_media_t(r)
}

func libvlc_media_get_meta(p_md plibvlc_media_t, e_meta libvlc_meta_t) string {
	r, _, _ := _libvlc_media_get_meta.Call(uintptr(p_md), uintptr(e_meta))
	return string(r)
}

func libvlc_media_set_meta(p_md plibvlc_media_t, e_meta libvlc_meta_t, psz_value string) {
	_libvlc_media_set_meta.Call(uintptr(p_md), uintptr(e_meta), toCStr(psz_value))
}

func libvlc_media_save_meta(p_md plibvlc_media_t) int {
	r, _, _ := _libvlc_media_save_meta.Call(uintptr(p_md))
	return int(r)
}

func libvlc_media_get_state(p_md plibvlc_media_t) libvlc_state_t {
	r, _, _ := _libvlc_media_get_state.Call(uintptr(p_md))
	return libvlc_state_t(r)
}

func libvlc_media_get_stats(p_md plibvlc_media_t, p_stats plibvlc_media_stats_t) int {
	r, _, _ := _libvlc_media_get_stats.Call(uintptr(p_md), uintptr(unsafe.Pointer(p_stats)))
	return int(r)
}

func libvlc_media_subitems(p_md plibvlc_media_t) plibvlc_media_list_t {
	r, _, _ := _libvlc_media_subitems.Call(uintptr(p_md))
	return plibvlc_media_list_t(r)
}

func libvlc_media_event_manager(p_md plibvlc_media_t) plibvlc_event_manager_t {
	r, _, _ := _libvlc_media_event_manager.Call(uintptr(p_md))
	return plibvlc_event_manager_t(r)
}

func libvlc_media_get_duration(p_md plibvlc_media_t) libvlc_time_t {
	r, _, _ := _libvlc_media_get_duration.Call(uintptr(p_md))
	return libvlc_time_t(r)
}

func libvlc_media_parse(p_md plibvlc_media_t) {
	_libvlc_media_parse.Call(uintptr(p_md))
}

func libvlc_media_parse_async(p_md plibvlc_media_t) {
	_libvlc_media_parse_async.Call(uintptr(p_md))
}

func libvlc_media_is_parsed(p_md plibvlc_media_t) int {
	r, _, _ := _libvlc_media_is_parsed.Call(uintptr(p_md))
	return int(r)
}

func libvlc_media_set_user_data(p_md plibvlc_media_t, p_new_user_data uintptr) {
	_libvlc_media_set_user_data.Call(uintptr(p_md), uintptr(p_new_user_data))
}

func libvlc_media_get_user_data(p_md plibvlc_media_t) uintptr {
	r, _, _ := _libvlc_media_get_user_data.Call(uintptr(p_md))
	return uintptr(r)
}

func libvlc_media_get_tracks_info(p_md plibvlc_media_t, tracks pplibvlc_media_track_info_t) int {
	r, _, _ := _libvlc_media_get_tracks_info.Call(uintptr(p_md), uintptr(unsafe.Pointer(tracks)))
	return int(r)
}

func libvlc_module_description_list_release(p_list plibvlc_module_description_t) {
	_libvlc_module_description_list_release.Call(uintptr(unsafe.Pointer(p_list)))
}

func libvlc_audio_filter_list_get(p_instance plibvlc_instance_t) plibvlc_module_description_t {
	r, _, _ := _libvlc_audio_filter_list_get.Call(uintptr(p_instance))
	return plibvlc_module_description_t(unsafe.Pointer(r))
}

func libvlc_video_filter_list_get(p_instance plibvlc_instance_t) plibvlc_module_description_t {
	r, _, _ := _libvlc_video_filter_list_get.Call(uintptr(p_instance))
	return plibvlc_module_description_t(unsafe.Pointer(r))
}

func libvlc_clock() int64 {
	r, _, _ := _libvlc_clock.Call()
	return int64(r)
}

func libvlc_errmsg() string {
	r, r2, r3 := _libvlc_errmsg.Call()
	fmt.Println(r, r2, r3, "strLen:")
	// 这个问题还有待解决。
	return ""
}

func libvlc_clearerr() {
	_libvlc_clearerr.Call()
}

func libvlc_printerr(fmtStr string) string {
	r, _, _ := _libvlc_printerr.Call(toCStr(fmtStr))
	return string(r)
}

func libvlc_new(args ...string) plibvlc_instance_t {
	var ptr uintptr
	if len(args) > 0 {
		ps := make([]uintptr, len(args))
		for i := 0; i < len(ps); i++ {
			ps[i] = toCStr(args[i])
		}
		ptr = uintptr(unsafe.Pointer(&ps[0]))
	}
	ret, _, _ := _libvlc_new.Call(uintptr(len(args)), ptr)
	return plibvlc_instance_t(ret)
}

func libvlc_release(p_instance plibvlc_instance_t) {
	_libvlc_release.Call(uintptr(p_instance))
}

func libvlc_retain(p_instance plibvlc_instance_t) {
	_libvlc_retain.Call(uintptr(p_instance))
}

func libvlc_add_intf(p_instance plibvlc_instance_t, name string) int {
	r, _, _ := _libvlc_add_intf.Call(uintptr(p_instance), toCStr(name))
	return int(r)
}

func libvlc_set_exit_handler(p_instance plibvlc_instance_t, cb cbtype1, opaque uintptr) {
	_libvlc_set_exit_handler.Call(uintptr(p_instance), uintptr(cb), uintptr(opaque))
}

func libvlc_wait(p_instance plibvlc_instance_t) {
	_libvlc_wait.Call(uintptr(p_instance))
}

func libvlc_set_user_agent(p_instance plibvlc_instance_t, name string, http string) {
	_libvlc_set_user_agent.Call(uintptr(p_instance), toCStr(name), toCStr(http))
}

func libvlc_get_version() string {
	r, _, _ := _libvlc_get_version.Call()
	fmt.Println(r)
	return ""
}

func libvlc_get_compiler() string {
	r, _, _ := _libvlc_get_compiler.Call()
	fmt.Println(r)
	return ""
}

func libvlc_get_changeset() string {
	r, _, _ := _libvlc_get_changeset.Call()
	fmt.Println(r)
	return ""
}

func libvlc_free(ptr uintptr) {
	_libvlc_free.Call(uintptr(ptr))
}

func libvlc_event_attach(p_event_manager plibvlc_event_manager_t, i_event_type libvlc_event_type_t, f_callback libvlc_callback_t, user_data uintptr) int {
	r, _, _ := _libvlc_event_attach.Call(uintptr(p_event_manager), uintptr(i_event_type), uintptr(f_callback), uintptr(user_data))
	return int(r)
}

func libvlc_event_detach(p_event_manager plibvlc_event_manager_t, i_event_type libvlc_event_type_t, f_callback libvlc_callback_t, p_user_data uintptr) {
	_libvlc_event_detach.Call(uintptr(p_event_manager), uintptr(i_event_type), uintptr(f_callback), uintptr(p_user_data))
}

func libvlc_event_type_name(event_type libvlc_event_type_t) string {
	r, _, _ := _libvlc_event_type_name.Call(uintptr(event_type))
	return string(r)
}

func libvlc_get_log_verbosity(p_instance plibvlc_instance_t) uint {
	r, _, _ := _libvlc_get_log_verbosity.Call(uintptr(p_instance))
	return uint(r)
}

func libvlc_set_log_verbosity(p_instance plibvlc_instance_t, level uint) {
	_libvlc_set_log_verbosity.Call(uintptr(p_instance), uintptr(level))
}

func libvlc_log_open(p_instance plibvlc_instance_t) plibvlc_log_t {
	r, _, _ := _libvlc_log_open.Call(uintptr(p_instance))
	return plibvlc_log_t(r)
}

func libvlc_log_close(p_log plibvlc_log_t) {
	_libvlc_log_close.Call(uintptr(p_log))
}

func libvlc_log_count(p_log plibvlc_log_t) uint {
	r, _, _ := _libvlc_log_count.Call(uintptr(p_log))
	return uint(r)
}

func libvlc_log_clear(p_log plibvlc_log_t) {
	_libvlc_log_clear.Call(uintptr(p_log))
}

func libvlc_log_get_iterator(p_log plibvlc_log_t) plibvlc_log_iterator_t {
	r, _, _ := _libvlc_log_get_iterator.Call(uintptr(p_log))
	return plibvlc_log_iterator_t(r)
}

func libvlc_log_iterator_free(p_iter plibvlc_log_iterator_t) {
	_libvlc_log_iterator_free.Call(uintptr(p_iter))
}

func libvlc_log_iterator_has_next(p_iter plibvlc_log_iterator_t) int {
	r, _, _ := _libvlc_log_iterator_has_next.Call(uintptr(p_iter))
	return int(r)
}

func libvlc_log_iterator_next(p_iter plibvlc_log_iterator_t, p_buffer plibvlc_log_message_t) plibvlc_log_message_t {
	r, _, _ := _libvlc_log_iterator_next.Call(uintptr(p_iter), uintptr(unsafe.Pointer(p_buffer)))
	return plibvlc_log_message_t(unsafe.Pointer(r))
}

func libvlc_audio_output_list_get(p_instance plibvlc_instance_t) plibvlc_audio_output_t {
	r, _, _ := _libvlc_audio_output_list_get.Call(uintptr(p_instance))
	return plibvlc_audio_output_t(unsafe.Pointer(r))
}

func libvlc_audio_output_list_release(p_list plibvlc_audio_output_t) {
	_libvlc_audio_output_list_release.Call(uintptr(unsafe.Pointer(p_list)))
}

func libvlc_audio_output_set(p_mi plibvlc_media_player_t, psz_name string) int {
	r, _, _ := _libvlc_audio_output_set.Call(uintptr(p_mi), toCStr(psz_name))
	return int(r)
}

func libvlc_audio_output_device_count(p_instance plibvlc_instance_t, psz_audio_output string) int {
	r, _, _ := _libvlc_audio_output_device_count.Call(uintptr(p_instance), toCStr(psz_audio_output))
	return int(r)
}

func libvlc_audio_output_device_longname(p_instance plibvlc_instance_t, psz_audio_output string, i_device int) string {
	r, _, _ := _libvlc_audio_output_device_longname.Call(uintptr(p_instance), toCStr(psz_audio_output), uintptr(i_device))
	return string(r)
}

func libvlc_audio_output_device_id(p_instance plibvlc_instance_t, psz_audio_output string, i_device int) string {
	r, _, _ := _libvlc_audio_output_device_id.Call(uintptr(p_instance), toCStr(psz_audio_output), uintptr(i_device))
	return string(r)
}

func libvlc_audio_output_device_set(p_mi plibvlc_media_player_t, psz_audio_output string, psz_device_id string) {
	_libvlc_audio_output_device_set.Call(uintptr(p_mi), toCStr(psz_audio_output), toCStr(psz_device_id))
}

func libvlc_audio_output_get_device_type(p_mi plibvlc_media_player_t) int {
	r, _, _ := _libvlc_audio_output_get_device_type.Call(uintptr(p_mi))
	return int(r)
}

func libvlc_audio_output_set_device_type(p_mi plibvlc_media_player_t, device_type int) {
	_libvlc_audio_output_set_device_type.Call(uintptr(p_mi), uintptr(device_type))
}

func libvlc_audio_toggle_mute(p_mi plibvlc_media_player_t) {
	_libvlc_audio_toggle_mute.Call(uintptr(p_mi))
}

func libvlc_audio_get_mute(p_mi plibvlc_media_player_t) int {
	r, _, _ := _libvlc_audio_get_mute.Call(uintptr(p_mi))
	return int(r)
}

func libvlc_audio_set_mute(p_mi plibvlc_media_player_t, status int) {
	_libvlc_audio_set_mute.Call(uintptr(p_mi), uintptr(status))
}

func libvlc_audio_get_volume(p_mi plibvlc_media_player_t) int {
	r, _, _ := _libvlc_audio_get_volume.Call(uintptr(p_mi))
	return int(r)
}

func libvlc_audio_set_volume(p_mi plibvlc_media_player_t, i_volume int) int {
	r, _, _ := _libvlc_audio_set_volume.Call(uintptr(p_mi), uintptr(i_volume))
	return int(r)
}

func libvlc_audio_get_track_count(p_mi plibvlc_media_player_t) int {
	r, _, _ := _libvlc_audio_get_track_count.Call(uintptr(p_mi))
	return int(r)
}

func libvlc_audio_get_track_description(p_mi plibvlc_media_player_t) plibvlc_track_description_t {
	r, _, _ := _libvlc_audio_get_track_description.Call(uintptr(p_mi))
	return plibvlc_track_description_t(unsafe.Pointer(r))
}

func libvlc_audio_get_track(p_mi plibvlc_media_player_t) int {
	r, _, _ := _libvlc_audio_get_track.Call(uintptr(p_mi))
	return int(r)
}

func libvlc_audio_set_track(p_mi plibvlc_media_player_t, i_track int) int {
	r, _, _ := _libvlc_audio_set_track.Call(uintptr(p_mi), uintptr(i_track))
	return int(r)
}

func libvlc_audio_get_channel(p_mi plibvlc_media_player_t) int {
	r, _, _ := _libvlc_audio_get_channel.Call(uintptr(p_mi))
	return int(r)
}

func libvlc_audio_set_channel(p_mi plibvlc_media_player_t, channel int) int {
	r, _, _ := _libvlc_audio_set_channel.Call(uintptr(p_mi), uintptr(channel))
	return int(r)
}

func libvlc_audio_get_delay(p_mi plibvlc_media_player_t) int64 {
	r, _, _ := _libvlc_audio_get_delay.Call(uintptr(p_mi))
	return int64(r)
}

func libvlc_audio_set_delay(p_mi plibvlc_media_player_t, i_delay int64) int {
	r, _, _ := _libvlc_audio_set_delay.Call(uintptr(p_mi), uintptr(i_delay))
	return int(r)
}

func libvlc_media_list_new(p_instance plibvlc_instance_t) plibvlc_media_list_t {
	r, _, _ := _libvlc_media_list_new.Call(uintptr(p_instance))
	return plibvlc_media_list_t(r)
}

func libvlc_media_list_release(p_ml plibvlc_media_list_t) {
	_libvlc_media_list_release.Call(uintptr(p_ml))
}

func libvlc_media_list_retain(p_ml plibvlc_media_list_t) {
	_libvlc_media_list_retain.Call(uintptr(p_ml))
}

func libvlc_media_list_add_file_content(p_ml plibvlc_media_list_t, psz_uri string) int {
	r, _, _ := _libvlc_media_list_add_file_content.Call(uintptr(p_ml), toCStr(psz_uri))
	return int(r)
}

func libvlc_media_list_set_media(p_ml plibvlc_media_list_t, p_md plibvlc_media_t) {
	_libvlc_media_list_set_media.Call(uintptr(p_ml), uintptr(p_md))
}

func libvlc_media_list_media(p_ml plibvlc_media_list_t) plibvlc_media_t {
	r, _, _ := _libvlc_media_list_media.Call(uintptr(p_ml))
	return plibvlc_media_t(r)
}

func libvlc_media_list_add_media(p_ml plibvlc_media_list_t, p_md plibvlc_media_t) int {
	r, _, _ := _libvlc_media_list_add_media.Call(uintptr(p_ml), uintptr(p_md))
	return int(r)
}

func libvlc_media_list_insert_media(p_ml plibvlc_media_list_t, p_md plibvlc_media_t, i_pos int) int {
	r, _, _ := _libvlc_media_list_insert_media.Call(uintptr(p_ml), uintptr(p_md), uintptr(i_pos))
	return int(r)
}

func libvlc_media_list_remove_index(p_ml plibvlc_media_list_t, i_pos int) int {
	r, _, _ := _libvlc_media_list_remove_index.Call(uintptr(p_ml), uintptr(i_pos))
	return int(r)
}

func libvlc_media_list_count(p_ml plibvlc_media_list_t) int {
	r, _, _ := _libvlc_media_list_count.Call(uintptr(p_ml))
	return int(r)
}

func libvlc_media_list_item_at_index(p_ml plibvlc_media_list_t, i_pos int) plibvlc_media_t {
	r, _, _ := _libvlc_media_list_item_at_index.Call(uintptr(p_ml), uintptr(i_pos))
	return plibvlc_media_t(r)
}

func libvlc_media_list_index_of_item(p_ml plibvlc_media_list_t, p_md plibvlc_media_t) int {
	r, _, _ := _libvlc_media_list_index_of_item.Call(uintptr(p_ml), uintptr(p_md))
	return int(r)
}

func libvlc_media_list_is_readonly(p_ml plibvlc_media_list_t) int {
	r, _, _ := _libvlc_media_list_is_readonly.Call(uintptr(p_ml))
	return int(r)
}

func libvlc_media_list_lock(p_ml plibvlc_media_list_t) {
	_libvlc_media_list_lock.Call(uintptr(p_ml))
}

func libvlc_media_list_unlock(p_ml plibvlc_media_list_t) {
	_libvlc_media_list_unlock.Call(uintptr(p_ml))
}

func libvlc_media_list_event_manager(p_ml plibvlc_media_list_t) plibvlc_event_manager_t {
	r, _, _ := _libvlc_media_list_event_manager.Call(uintptr(p_ml))
	return plibvlc_event_manager_t(r)
}

func libvlc_media_list_player_new(p_instance plibvlc_instance_t) plibvlc_media_list_player_t {
	r, _, _ := _libvlc_media_list_player_new.Call(uintptr(p_instance))
	return plibvlc_media_list_player_t(r)
}

func libvlc_media_list_player_release(p_mlp plibvlc_media_list_player_t) {
	_libvlc_media_list_player_release.Call(uintptr(p_mlp))
}

func libvlc_media_list_player_retain(p_mlp plibvlc_media_list_player_t) {
	_libvlc_media_list_player_retain.Call(uintptr(p_mlp))
}

func libvlc_media_list_player_event_manager(p_mlp plibvlc_media_list_player_t) plibvlc_event_manager_t {
	r, _, _ := _libvlc_media_list_player_event_manager.Call(uintptr(p_mlp))
	return plibvlc_event_manager_t(r)
}

func libvlc_media_list_player_set_media_player(p_mlp plibvlc_media_list_player_t, p_mi plibvlc_media_player_t) {
	_libvlc_media_list_player_set_media_player.Call(uintptr(p_mlp), uintptr(p_mi))
}

func libvlc_media_list_player_set_media_list(p_mlp plibvlc_media_list_player_t, p_mlist plibvlc_media_list_t) {
	_libvlc_media_list_player_set_media_list.Call(uintptr(p_mlp), uintptr(p_mlist))
}

func libvlc_media_list_player_play(p_mlp plibvlc_media_list_player_t) {
	_libvlc_media_list_player_play.Call(uintptr(p_mlp))
}

func libvlc_media_list_player_pause(p_mlp plibvlc_media_list_player_t) {
	_libvlc_media_list_player_pause.Call(uintptr(p_mlp))
}

func libvlc_media_list_player_is_playing(p_mlp plibvlc_media_list_player_t) int {
	r, _, _ := _libvlc_media_list_player_is_playing.Call(uintptr(p_mlp))
	return int(r)
}

func libvlc_media_list_player_get_state(p_mlp plibvlc_media_list_player_t) libvlc_state_t {
	r, _, _ := _libvlc_media_list_player_get_state.Call(uintptr(p_mlp))
	return libvlc_state_t(r)
}

func libvlc_media_list_player_play_item_at_index(p_mlp plibvlc_media_list_player_t, i_index int) int {
	r, _, _ := _libvlc_media_list_player_play_item_at_index.Call(uintptr(p_mlp), uintptr(i_index))
	return int(r)
}

func libvlc_media_list_player_play_item(p_mlp plibvlc_media_list_player_t, p_md plibvlc_media_t) int {
	r, _, _ := _libvlc_media_list_player_play_item.Call(uintptr(p_mlp), uintptr(p_md))
	return int(r)
}

func libvlc_media_list_player_stop(p_mlp plibvlc_media_list_player_t) {
	_libvlc_media_list_player_stop.Call(uintptr(p_mlp))
}

func libvlc_media_list_player_next(p_mlp plibvlc_media_list_player_t) int {
	r, _, _ := _libvlc_media_list_player_next.Call(uintptr(p_mlp))
	return int(r)
}

func libvlc_media_list_player_previous(p_mlp plibvlc_media_list_player_t) int {
	r, _, _ := _libvlc_media_list_player_previous.Call(uintptr(p_mlp))
	return int(r)
}

func libvlc_media_list_player_set_playback_mode(p_mlp plibvlc_media_list_player_t, e_mode libvlc_playback_mode_t) {
	_libvlc_media_list_player_set_playback_mode.Call(uintptr(p_mlp), uintptr(e_mode))
}

func libvlc_media_library_new(p_instance plibvlc_instance_t) plibvlc_media_library_t {
	r, _, _ := _libvlc_media_library_new.Call(uintptr(p_instance))
	return plibvlc_media_library_t(r)
}

func libvlc_media_library_release(p_mlib plibvlc_media_library_t) {
	_libvlc_media_library_release.Call(uintptr(p_mlib))
}

func libvlc_media_library_retain(p_mlib plibvlc_media_library_t) {
	_libvlc_media_library_retain.Call(uintptr(p_mlib))
}

func libvlc_media_library_load(p_mlib plibvlc_media_library_t) int {
	r, _, _ := _libvlc_media_library_load.Call(uintptr(p_mlib))
	return int(r)
}

func libvlc_media_library_media_list(p_mlib plibvlc_media_library_t) plibvlc_media_list_t {
	r, _, _ := _libvlc_media_library_media_list.Call(uintptr(p_mlib))
	return plibvlc_media_list_t(r)
}

func libvlc_video_get_adjust_int(p_mi plibvlc_media_player_t, option uint) int {
	r, _, _ := _libvlc_video_get_adjust_int.Call(uintptr(p_mi), uintptr(option))
	return int(r)
}

func libvlc_video_set_adjust_int(p_mi plibvlc_media_player_t, option uint, value int) {
	_libvlc_video_set_adjust_int.Call(uintptr(p_mi), uintptr(option), uintptr(value))
}

func libvlc_video_get_adjust_float(p_mi plibvlc_media_player_t, option uint) float32 {
	_libvlc_video_get_adjust_float.Call(uintptr(p_mi), uintptr(option))
	return floatpatch.Getfloat32()
}

func libvlc_video_set_adjust_float(p_mi plibvlc_media_player_t, option uint, value float32) {
	_libvlc_video_set_adjust_float.Call(uintptr(p_mi), uintptr(option), uintptr(math.Float32bits(value)))
}

func libvlc_video_get_logo_int(p_mi plibvlc_media_player_t, option uint) int {
	r, _, _ := _libvlc_video_get_logo_int.Call(uintptr(p_mi), uintptr(option))
	return int(r)
}

func libvlc_video_set_logo_int(p_mi plibvlc_media_player_t, option uint, value int) {
	_libvlc_video_set_logo_int.Call(uintptr(p_mi), uintptr(option), uintptr(value))
}

func libvlc_video_set_logo_string(p_mi plibvlc_media_player_t, option uint, psz_value string) {
	_libvlc_video_set_logo_string.Call(uintptr(p_mi), uintptr(option), toCStr(psz_value))
}

func libvlc_audio_set_format_callbacks(mp plibvlc_media_player_t, setup libvlc_audio_setup_cb, cleanup libvlc_audio_cleanup_cb) {
	_libvlc_audio_set_format_callbacks.Call(uintptr(mp), uintptr(setup), uintptr(cleanup))
}

func libvlc_audio_set_format(mp plibvlc_media_player_t, format string, rate uint, channels uint) {
	_libvlc_audio_set_format.Call(uintptr(mp), toCStr(format), uintptr(rate), uintptr(channels))
}

func libvlc_media_player_get_length(p_mi plibvlc_media_player_t) libvlc_time_t {
	r, _, _ := _libvlc_media_player_get_length.Call(uintptr(p_mi))
	return libvlc_time_t(r)
}

func libvlc_media_player_get_time(p_mi plibvlc_media_player_t) libvlc_time_t {
	r, _, _ := _libvlc_media_player_get_time.Call(uintptr(p_mi))
	return libvlc_time_t(r)
}

func libvlc_media_player_set_time(p_mi plibvlc_media_player_t, i_time libvlc_time_t) {
	_libvlc_media_player_set_time.Call(uintptr(p_mi), uintptr(i_time))
}

func libvlc_media_player_get_position(p_mi plibvlc_media_player_t) float32 {
	_libvlc_media_player_get_position.Call(uintptr(p_mi))
	return floatpatch.Getfloat32()
}

func libvlc_media_player_set_position(p_mi plibvlc_media_player_t, f_pos float32) {
	_libvlc_media_player_set_position.Call(uintptr(p_mi), uintptr(math.Float32bits(f_pos)))
}

func libvlc_media_player_set_chapter(p_mi plibvlc_media_player_t, i_chapter int) {
	_libvlc_media_player_set_chapter.Call(uintptr(p_mi), uintptr(i_chapter))
}

func libvlc_media_player_get_chapter(p_mi plibvlc_media_player_t) int {
	r, _, _ := _libvlc_media_player_get_chapter.Call(uintptr(p_mi))
	return int(r)
}

func libvlc_media_player_get_chapter_count(p_mi plibvlc_media_player_t) int {
	r, _, _ := _libvlc_media_player_get_chapter_count.Call(uintptr(p_mi))
	return int(r)
}

func libvlc_media_player_will_play(p_mi plibvlc_media_player_t) int {
	r, _, _ := _libvlc_media_player_will_play.Call(uintptr(p_mi))
	return int(r)
}

func libvlc_media_player_get_chapter_count_for_title(p_mi plibvlc_media_player_t, i_title int) int {
	r, _, _ := _libvlc_media_player_get_chapter_count_for_title.Call(uintptr(p_mi), uintptr(i_title))
	return int(r)
}

func libvlc_media_player_set_title(p_mi plibvlc_media_player_t, i_title int) {
	_libvlc_media_player_set_title.Call(uintptr(p_mi), uintptr(i_title))
}

func libvlc_media_player_get_title(p_mi plibvlc_media_player_t) int {
	r, _, _ := _libvlc_media_player_get_title.Call(uintptr(p_mi))
	return int(r)
}

func libvlc_media_player_get_title_count(p_mi plibvlc_media_player_t) int {
	r, _, _ := _libvlc_media_player_get_title_count.Call(uintptr(p_mi))
	return int(r)
}

func libvlc_media_player_previous_chapter(p_mi plibvlc_media_player_t) {
	_libvlc_media_player_previous_chapter.Call(uintptr(p_mi))
}

func libvlc_media_player_next_chapter(p_mi plibvlc_media_player_t) {
	_libvlc_media_player_next_chapter.Call(uintptr(p_mi))
}

func libvlc_media_player_get_rate(p_mi plibvlc_media_player_t) float32 {
	_libvlc_media_player_get_rate.Call(uintptr(p_mi))
	return floatpatch.Getfloat32()
}

func libvlc_media_player_set_rate(p_mi plibvlc_media_player_t, rate float32) int {
	r, _, _ := _libvlc_media_player_set_rate.Call(uintptr(p_mi), uintptr(math.Float32bits(rate)))
	return int(r)
}

func libvlc_media_player_get_state(p_mi plibvlc_media_player_t) libvlc_state_t {
	r, _, _ := _libvlc_media_player_get_state.Call(uintptr(p_mi))
	return libvlc_state_t(r)
}

func libvlc_media_player_get_fps(p_mi plibvlc_media_player_t) float32 {
	_libvlc_media_player_get_fps.Call(uintptr(p_mi))
	return floatpatch.Getfloat32()
}

func libvlc_media_player_has_vout(p_mi plibvlc_media_player_t) uint {
	r, _, _ := _libvlc_media_player_has_vout.Call(uintptr(p_mi))
	return uint(r)
}

func libvlc_media_player_is_seekable(p_mi plibvlc_media_player_t) int {
	r, _, _ := _libvlc_media_player_is_seekable.Call(uintptr(p_mi))
	return int(r)
}

func libvlc_media_player_can_pause(p_mi plibvlc_media_player_t) int {
	r, _, _ := _libvlc_media_player_can_pause.Call(uintptr(p_mi))
	return int(r)
}

func libvlc_media_player_next_frame(p_mi plibvlc_media_player_t) {
	_libvlc_media_player_next_frame.Call(uintptr(p_mi))
}

func libvlc_media_player_navigate(p_mi plibvlc_media_player_t, navigate uint) {
	_libvlc_media_player_navigate.Call(uintptr(p_mi), uintptr(navigate))
}

func libvlc_track_description_list_release(p_track_description plibvlc_track_description_t) {
	_libvlc_track_description_list_release.Call(uintptr(unsafe.Pointer(p_track_description)))
}

func libvlc_track_description_release(p_track_description plibvlc_track_description_t) {
	_libvlc_track_description_release.Call(uintptr(unsafe.Pointer(p_track_description)))
}

func libvlc_toggle_fullscreen(p_mi plibvlc_media_player_t) {
	_libvlc_toggle_fullscreen.Call(uintptr(p_mi))
}

func libvlc_set_fullscreen(p_mi plibvlc_media_player_t, b_fullscreen int) {
	_libvlc_set_fullscreen.Call(uintptr(p_mi), uintptr(b_fullscreen))
}

func libvlc_get_fullscreen(p_mi plibvlc_media_player_t) int {
	r, _, _ := _libvlc_get_fullscreen.Call(uintptr(p_mi))
	return int(r)
}

func libvlc_video_set_key_input(p_mi plibvlc_media_player_t, on uint) {
	_libvlc_video_set_key_input.Call(uintptr(p_mi), uintptr(on))
}

func libvlc_video_set_mouse_input(p_mi plibvlc_media_player_t, on uint) {
	_libvlc_video_set_mouse_input.Call(uintptr(p_mi), uintptr(on))
}

func libvlc_video_get_size(p_mi plibvlc_media_player_t, num uint, px *uint, py *uint) int {
	r, _, _ := _libvlc_video_get_size.Call(uintptr(p_mi), uintptr(num), uintptr(unsafe.Pointer(px)), uintptr(unsafe.Pointer(py)))
	return int(r)
}

func libvlc_video_get_height(p_mi plibvlc_media_player_t) int {
	r, _, _ := _libvlc_video_get_height.Call(uintptr(p_mi))
	return int(r)
}

func libvlc_video_get_width(p_mi plibvlc_media_player_t) int {
	r, _, _ := _libvlc_video_get_width.Call(uintptr(p_mi))
	return int(r)
}

func libvlc_video_get_cursor(p_mi plibvlc_media_player_t, num uint, px *int, py *int) int {
	r, _, _ := _libvlc_video_get_cursor.Call(uintptr(p_mi), uintptr(num), uintptr(unsafe.Pointer(px)), uintptr(unsafe.Pointer(py)))
	return int(r)
}

func libvlc_video_get_scale(p_mi plibvlc_media_player_t) float32 {
	_libvlc_video_get_scale.Call(uintptr(p_mi))
	return floatpatch.Getfloat32()
}

func libvlc_video_set_scale(p_mi plibvlc_media_player_t, f_factor float32) {
	_libvlc_video_set_scale.Call(uintptr(p_mi), uintptr(math.Float32bits(f_factor)))
}

func libvlc_video_get_aspect_ratio(p_mi plibvlc_media_player_t) string {
	r, _, _ := _libvlc_video_get_aspect_ratio.Call(uintptr(p_mi))
	return string(r)
}

func libvlc_video_set_aspect_ratio(p_mi plibvlc_media_player_t, psz_aspect string) {
	_libvlc_video_set_aspect_ratio.Call(uintptr(p_mi), toCStr(psz_aspect))
}

func libvlc_video_get_spu(p_mi plibvlc_media_player_t) int {
	r, _, _ := _libvlc_video_get_spu.Call(uintptr(p_mi))
	return int(r)
}

func libvlc_video_get_spu_count(p_mi plibvlc_media_player_t) int {
	r, _, _ := _libvlc_video_get_spu_count.Call(uintptr(p_mi))
	return int(r)
}

func libvlc_video_get_spu_description(p_mi plibvlc_media_player_t) plibvlc_track_description_t {
	r, _, _ := _libvlc_video_get_spu_description.Call(uintptr(p_mi))
	return plibvlc_track_description_t(unsafe.Pointer(r))
}

func libvlc_video_set_spu(p_mi plibvlc_media_player_t, i_spu uint) int {
	r, _, _ := _libvlc_video_set_spu.Call(uintptr(p_mi), uintptr(i_spu))
	return int(r)
}

func libvlc_video_set_subtitle_file(p_mi plibvlc_media_player_t, psz_subtitle string) int {
	r, _, _ := _libvlc_video_set_subtitle_file.Call(uintptr(p_mi), toCStr(psz_subtitle))
	return int(r)
}

func libvlc_video_get_spu_delay(p_mi plibvlc_media_player_t) int64 {
	r, _, _ := _libvlc_video_get_spu_delay.Call(uintptr(p_mi))
	return int64(r)
}

func libvlc_video_set_spu_delay(p_mi plibvlc_media_player_t, i_delay int64) int {
	r, _, _ := _libvlc_video_set_spu_delay.Call(uintptr(p_mi), uintptr(i_delay))
	return int(r)
}

func libvlc_video_get_title_description(p_mi plibvlc_media_player_t) plibvlc_track_description_t {
	r, _, _ := _libvlc_video_get_title_description.Call(uintptr(p_mi))
	return plibvlc_track_description_t(unsafe.Pointer(r))
}

func libvlc_video_get_chapter_description(p_mi plibvlc_media_player_t, i_title int) plibvlc_track_description_t {
	r, _, _ := _libvlc_video_get_chapter_description.Call(uintptr(p_mi), uintptr(i_title))
	return plibvlc_track_description_t(unsafe.Pointer(r))
}

func libvlc_video_get_crop_geometry(p_mi plibvlc_media_player_t) string {
	r, _, _ := _libvlc_video_get_crop_geometry.Call(uintptr(p_mi))
	return string(r)
}

func libvlc_video_set_crop_geometry(p_mi plibvlc_media_player_t, psz_geometry string) {
	_libvlc_video_set_crop_geometry.Call(uintptr(p_mi), toCStr(psz_geometry))
}

func libvlc_video_get_teletext(p_mi plibvlc_media_player_t) int {
	r, _, _ := _libvlc_video_get_teletext.Call(uintptr(p_mi))
	return int(r)
}

func libvlc_video_set_teletext(p_mi plibvlc_media_player_t, i_page int) {
	_libvlc_video_set_teletext.Call(uintptr(p_mi), uintptr(i_page))
}

func libvlc_toggle_teletext(p_mi plibvlc_media_player_t) {
	_libvlc_toggle_teletext.Call(uintptr(p_mi))
}

func libvlc_video_get_track_count(p_mi plibvlc_media_player_t) int {
	r, _, _ := _libvlc_video_get_track_count.Call(uintptr(p_mi))
	return int(r)
}

func libvlc_video_get_track_description(p_mi plibvlc_media_player_t) plibvlc_track_description_t {
	r, _, _ := _libvlc_video_get_track_description.Call(uintptr(p_mi))
	return plibvlc_track_description_t(unsafe.Pointer(r))
}

func libvlc_video_get_track(p_mi plibvlc_media_player_t) int {
	r, _, _ := _libvlc_video_get_track.Call(uintptr(p_mi))
	return int(r)
}

func libvlc_video_set_track(p_mi plibvlc_media_player_t, i_track int) int {
	r, _, _ := _libvlc_video_set_track.Call(uintptr(p_mi), uintptr(i_track))
	return int(r)
}

func libvlc_video_take_snapshot(p_mi plibvlc_media_player_t, num uint, psz_filepath string, i_width uint, i_height uint) int {
	r, _, _ := _libvlc_video_take_snapshot.Call(uintptr(p_mi), uintptr(num), toCStr(psz_filepath), uintptr(i_width), uintptr(i_height))
	return int(r)
}

func libvlc_video_set_deinterlace(p_mi plibvlc_media_player_t, psz_mode string) {
	_libvlc_video_set_deinterlace.Call(uintptr(p_mi), toCStr(psz_mode))
}

func libvlc_video_get_marquee_int(p_mi plibvlc_media_player_t, option uint) int {
	r, _, _ := _libvlc_video_get_marquee_int.Call(uintptr(p_mi), uintptr(option))
	return int(r)
}

func libvlc_video_get_marquee_string(p_mi plibvlc_media_player_t, option uint) string {
	r, _, _ := _libvlc_video_get_marquee_string.Call(uintptr(p_mi), uintptr(option))
	return string(r)
}

func libvlc_video_set_marquee_int(p_mi plibvlc_media_player_t, option uint, i_val int) {
	_libvlc_video_set_marquee_int.Call(uintptr(p_mi), uintptr(option), uintptr(i_val))
}

func libvlc_video_set_marquee_string(p_mi plibvlc_media_player_t, option uint, psz_text string) {
	_libvlc_video_set_marquee_string.Call(uintptr(p_mi), uintptr(option), toCStr(psz_text))
}

func libvlc_audio_set_callbacks(mp plibvlc_media_player_t, play libvlc_audio_play_cb, pause libvlc_audio_pause_cb, resume libvlc_audio_resume_cb, flush libvlc_audio_flush_cb, drain libvlc_audio_drain_cb, opaque uintptr) {
	_libvlc_audio_set_callbacks.Call(uintptr(mp), uintptr(play), uintptr(pause), uintptr(resume), uintptr(flush), uintptr(drain), uintptr(opaque))
}

func libvlc_audio_set_volume_callback(mp plibvlc_media_player_t, set_volume libvlc_audio_set_volume_cb) {
	_libvlc_audio_set_volume_callback.Call(uintptr(mp), uintptr(set_volume))
}

func libvlc_video_set_callbacks(mp plibvlc_media_player_t, lock libvlc_video_lock_cb, unlock libvlc_video_unlock_cb, display libvlc_video_display_cb, opaque uintptr) {
	_libvlc_video_set_callbacks.Call(uintptr(mp), uintptr(lock), uintptr(unlock), uintptr(display), uintptr(opaque))
}

func libvlc_video_set_format(mp plibvlc_media_player_t, chroma string, width uint, height uint, pitch uint) {
	_libvlc_video_set_format.Call(uintptr(mp), toCStr(chroma), uintptr(width), uintptr(height), uintptr(pitch))
}

func libvlc_video_set_format_callbacks(mp plibvlc_media_player_t, setup libvlc_video_format_cb, cleanup libvlc_video_cleanup_cb) {
	_libvlc_video_set_format_callbacks.Call(uintptr(mp), uintptr(setup), uintptr(cleanup))
}

func libvlc_media_player_set_nsobject(p_mi plibvlc_media_player_t, drawable uintptr) {
	_libvlc_media_player_set_nsobject.Call(uintptr(p_mi), uintptr(drawable))
}

func libvlc_media_player_get_nsobject(p_mi plibvlc_media_player_t) uintptr {
	r, _, _ := _libvlc_media_player_get_nsobject.Call(uintptr(p_mi))
	return uintptr(r)
}

func libvlc_media_player_set_agl(p_mi plibvlc_media_player_t, drawable uint32) {
	_libvlc_media_player_set_agl.Call(uintptr(p_mi), uintptr(drawable))
}

func libvlc_media_player_get_agl(p_mi plibvlc_media_player_t) uint32 {
	r, _, _ := _libvlc_media_player_get_agl.Call(uintptr(p_mi))
	return uint32(r)
}

func libvlc_media_player_set_xwindow(p_mi plibvlc_media_player_t, drawable uint32) {
	_libvlc_media_player_set_xwindow.Call(uintptr(p_mi), uintptr(drawable))
}

func libvlc_media_player_get_xwindow(p_mi plibvlc_media_player_t) uint32 {
	r, _, _ := _libvlc_media_player_get_xwindow.Call(uintptr(p_mi))
	return uint32(r)
}

func libvlc_media_player_set_hwnd(p_mi plibvlc_media_player_t, drawable uintptr) {
	_libvlc_media_player_set_hwnd.Call(uintptr(p_mi), uintptr(drawable))
}

func libvlc_media_player_get_hwnd(p_mi plibvlc_media_player_t) uintptr {
	r, _, _ := _libvlc_media_player_get_hwnd.Call(uintptr(p_mi))
	return uintptr(r)
}

func libvlc_media_discoverer_new_from_name(p_inst plibvlc_instance_t, psz_name string) plibvlc_media_discoverer_t {
	r, _, _ := _libvlc_media_discoverer_new_from_name.Call(uintptr(p_inst), toCStr(psz_name))
	return plibvlc_media_discoverer_t(r)
}

func libvlc_media_discoverer_release(p_mdis plibvlc_media_discoverer_t) {
	_libvlc_media_discoverer_release.Call(uintptr(p_mdis))
}

func libvlc_media_discoverer_localized_name(p_mdis plibvlc_media_discoverer_t) string {
	r, _, _ := _libvlc_media_discoverer_localized_name.Call(uintptr(p_mdis))
	return string(r)
}

func libvlc_media_discoverer_media_list(p_mdis plibvlc_media_discoverer_t) plibvlc_media_list_t {
	r, _, _ := _libvlc_media_discoverer_media_list.Call(uintptr(p_mdis))
	return plibvlc_media_list_t(r)
}

func libvlc_media_discoverer_event_manager(p_mdis plibvlc_media_discoverer_t) plibvlc_event_manager_t {
	r, _, _ := _libvlc_media_discoverer_event_manager.Call(uintptr(p_mdis))
	return plibvlc_event_manager_t(r)
}

func libvlc_media_discoverer_is_running(p_mdis plibvlc_media_discoverer_t) int {
	r, _, _ := _libvlc_media_discoverer_is_running.Call(uintptr(p_mdis))
	return int(r)
}

func libvlc_vlm_release(p_instance plibvlc_instance_t) {
	_libvlc_vlm_release.Call(uintptr(p_instance))
}

func libvlc_vlm_add_broadcast(p_instance plibvlc_instance_t, psz_name string, psz_input string, psz_output string, i_options int, ppsz_options Pstring, b_enabled int, b_loop int) int {
	r, _, _ := _libvlc_vlm_add_broadcast.Call(uintptr(p_instance), toCStr(psz_name), toCStr(psz_input), toCStr(psz_output), uintptr(i_options), uintptr(ppsz_options), uintptr(b_enabled), uintptr(b_loop))
	return int(r)
}

func libvlc_vlm_add_vod(p_instance plibvlc_instance_t, psz_name string, psz_input string, i_options int, ppsz_options Pstring, b_enabled int, psz_mux string) int {
	r, _, _ := _libvlc_vlm_add_vod.Call(uintptr(p_instance), toCStr(psz_name), toCStr(psz_input), uintptr(i_options), uintptr(ppsz_options), uintptr(b_enabled), toCStr(psz_mux))
	return int(r)
}

func libvlc_vlm_del_media(p_instance plibvlc_instance_t, psz_name string) int {
	r, _, _ := _libvlc_vlm_del_media.Call(uintptr(p_instance), toCStr(psz_name))
	return int(r)
}

func libvlc_vlm_set_enabled(p_instance plibvlc_instance_t, psz_name string, b_enabled int) int {
	r, _, _ := _libvlc_vlm_set_enabled.Call(uintptr(p_instance), toCStr(psz_name), uintptr(b_enabled))
	return int(r)
}

func libvlc_vlm_set_output(p_instance plibvlc_instance_t, psz_name string, psz_output string) int {
	r, _, _ := _libvlc_vlm_set_output.Call(uintptr(p_instance), toCStr(psz_name), toCStr(psz_output))
	return int(r)
}

func libvlc_vlm_set_input(p_instance plibvlc_instance_t, psz_name string, psz_input string) int {
	r, _, _ := _libvlc_vlm_set_input.Call(uintptr(p_instance), toCStr(psz_name), toCStr(psz_input))
	return int(r)
}

func libvlc_vlm_add_input(p_instance plibvlc_instance_t, psz_name string, psz_input string) int {
	r, _, _ := _libvlc_vlm_add_input.Call(uintptr(p_instance), toCStr(psz_name), toCStr(psz_input))
	return int(r)
}

func libvlc_vlm_set_loop(p_instance plibvlc_instance_t, psz_name string, b_loop int) int {
	r, _, _ := _libvlc_vlm_set_loop.Call(uintptr(p_instance), toCStr(psz_name), uintptr(b_loop))
	return int(r)
}

func libvlc_vlm_set_mux(p_instance plibvlc_instance_t, psz_name string, psz_mux string) int {
	r, _, _ := _libvlc_vlm_set_mux.Call(uintptr(p_instance), toCStr(psz_name), toCStr(psz_mux))
	return int(r)
}

func libvlc_vlm_change_media(p_instance plibvlc_instance_t, psz_name string, psz_input string, psz_output string, i_options int, ppsz_options Pstring, b_enabled int, b_loop int) int {
	r, _, _ := _libvlc_vlm_change_media.Call(uintptr(p_instance), toCStr(psz_name), toCStr(psz_input), toCStr(psz_output), uintptr(i_options), uintptr(ppsz_options), uintptr(b_enabled), uintptr(b_loop))
	return int(r)
}

func libvlc_vlm_play_media(p_instance plibvlc_instance_t, psz_name string) int {
	r, _, _ := _libvlc_vlm_play_media.Call(uintptr(p_instance), toCStr(psz_name))
	return int(r)
}

func libvlc_vlm_stop_media(p_instance plibvlc_instance_t, psz_name string) int {
	r, _, _ := _libvlc_vlm_stop_media.Call(uintptr(p_instance), toCStr(psz_name))
	return int(r)
}

func libvlc_vlm_pause_media(p_instance plibvlc_instance_t, psz_name string) int {
	r, _, _ := _libvlc_vlm_pause_media.Call(uintptr(p_instance), toCStr(psz_name))
	return int(r)
}

func libvlc_vlm_seek_media(p_instance plibvlc_instance_t, psz_name string, f_percentage float32) int {
	r, _, _ := _libvlc_vlm_seek_media.Call(uintptr(p_instance), toCStr(psz_name), uintptr(math.Float32bits(f_percentage)))
	return int(r)
}

func libvlc_vlm_show_media(p_instance plibvlc_instance_t, psz_name string) string {
	r, _, _ := _libvlc_vlm_show_media.Call(uintptr(p_instance), toCStr(psz_name))
	return string(r)
}

func libvlc_vlm_get_media_instance_position(p_instance plibvlc_instance_t, psz_name string, i_instance int) float32 {
	_libvlc_vlm_get_media_instance_position.Call(uintptr(p_instance), toCStr(psz_name), uintptr(i_instance))
	return floatpatch.Getfloat32()
}

func libvlc_vlm_get_media_instance_time(p_instance plibvlc_instance_t, psz_name string, i_instance int) int {
	r, _, _ := _libvlc_vlm_get_media_instance_time.Call(uintptr(p_instance), toCStr(psz_name), uintptr(i_instance))
	return int(r)
}

func libvlc_vlm_get_media_instance_length(p_instance plibvlc_instance_t, psz_name string, i_instance int) int {
	r, _, _ := _libvlc_vlm_get_media_instance_length.Call(uintptr(p_instance), toCStr(psz_name), uintptr(i_instance))
	return int(r)
}

func libvlc_vlm_get_media_instance_rate(p_instance plibvlc_instance_t, psz_name string, i_instance int) int {
	r, _, _ := _libvlc_vlm_get_media_instance_rate.Call(uintptr(p_instance), toCStr(psz_name), uintptr(i_instance))
	return int(r)
}

func libvlc_vlm_get_event_manager(p_instance plibvlc_instance_t) plibvlc_event_manager_t {
	r, _, _ := _libvlc_vlm_get_event_manager.Call(uintptr(p_instance))
	return plibvlc_event_manager_t(r)
}

func libvlc_playlist_play(p_instance plibvlc_instance_t, i_id int, i_options int, ppsz_options Pstring) {
	_libvlc_playlist_play.Call(uintptr(p_instance), uintptr(i_id), uintptr(i_options), uintptr(ppsz_options))
}
