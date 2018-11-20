// 作者： ying32
// bass包初用于govcl例程 mp3Player
// bass二进制下载：http://www.un4seen.com/
//  windows: http://www.un4seen.com/download.php?bass24
//  macOS: http://www.un4seen.com/download.php?bass24-osx
//  linux: http://www.un4seen.com/download.php?bass24-linux

package bass

import (
	"errors"
	"fmt"
)

type TPlayState int

const (
	PsUnkown = iota + 0
	PsPlaying
	PsStoped
	PsPaused
)

type TBass struct {
	hstream HSTREAM
	State   TPlayState
}

// 新建播放类
func NewBass() *TBass {
	if BASS_GetVersion()>>16 != BASSVERSION {
		panic("版本与头文件不符合！")
	}
	bass := new(TBass)
	BASS_Init(-1, 44100, 0, 0, 0)
	return bass
}

// 释放库
func BassFree() {
	BASS_Free()
}

// 打开一个文件，用作播放的
func (c *TBass) OpenFile(filename string) error {
	c.Close()
	c.hstream = BASS_StreamCreateFile(false, filename, 0, 0, 0)
	return errorFromCode()
}

// 停止并释放一个流
func (c *TBass) Close() error {
	c.Stop()
	if c.hstream != 0 {
		BASS_StreamFree(c.hstream)
		c.hstream = 0
	}
	return errorFromCode()
}

// 播放
func (c *TBass) Play(loop bool) error {
	if c.hstream != 0 {
		if c.State == PsPlaying {
			return nil
		}
		if BASS_ChannelPlay(c.hstream, loop) {
			c.State = PsPlaying
		}
	}
	return errorFromCode()
}

// 停止播放
func (c *TBass) Stop() error {
	if c.hstream == 0 || c.State == PsStoped {
		return nil
	}
	if BASS_ChannelStop(c.hstream) {
		c.State = PsStoped
	}
	return errorFromCode()
}

// 暂停播放
func (c *TBass) Pause() error {
	if c.hstream == 0 || c.State != PsPlaying {
		return nil
	}
	if BASS_ChannelPause(c.hstream) {
		c.State = PsPaused
	}
	return errorFromCode()
}

// 文件是否效
func (c *TBass) IsValid() bool {
	return c.hstream != 0
}

// 在取媒体文件长度
func (c *TBass) GetLength() (uint, error) {
	return uint(BASS_ChannelBytes2Seconds(c.hstream, BASS_ChannelGetLength(c.hstream, BASS_POS_BYTE)) * 1000),
		errorFromCode()
}

// 获取媒体当前播放的位置
func (c *TBass) GetPosition() (uint, error) {
	return uint(BASS_ChannelBytes2Seconds(c.hstream, BASS_ChannelGetPosition(c.hstream, BASS_POS_BYTE)) * 1000),
		errorFromCode()
}

// 设置媒体播放位置
func (c *TBass) SetPosition(val uint) error {
	BASS_ChannelSetPosition(c.hstream, BASS_ChannelSeconds2Bytes(c.hstream, float64(val)/1000.0), BASS_POS_BYTE)
	return errorFromCode()
}

// 获取播放音量
func (c *TBass) GetVolume() (int, error) {
	var val float32
	BASS_ChannelGetAttribute(c.hstream, BASS_ATTRIB_VOL, &val)
	return int(val * 100), errorFromCode()
}

// 设置播放音量
func (c *TBass) SetVolume(val int) error {
	BASS_ChannelSetAttribute(c.hstream, BASS_ATTRIB_VOL, float32(val)/100.0)
	return errorFromCode()
}

func (c *TBass) TimeStrLabel() string {
	lPosi, _ := c.GetPosition()
	lLen, _ := c.GetLength()

	return fmt.Sprintf("%.2d:%.2d/%.2d:%.2d", lPosi/1000/60, lPosi/1000%60, lLen/1000/60, lLen/1000%60)
}

func GetFileLength(fileName string) int {
	hstream := BASS_StreamCreateFile(false, fileName, 0, 0, 0)
	if hstream > 0 {
		defer BASS_StreamFree(hstream)
		return int(BASS_ChannelBytes2Seconds(hstream, BASS_ChannelGetLength(hstream, BASS_POS_BYTE)) * 1000)
	}
	return 0
}

func errorFromCode() error {
	switch BASS_ErrorGetCode() {
	case BASS_OK:
		return nil //errors.New("all is OK")
	case BASS_ERROR_MEM:
		return errors.New("memory error")
	case BASS_ERROR_FILEOPEN:
		return errors.New("can't open the file")
	case BASS_ERROR_DRIVER:
		return errors.New("can't find a free/valid driver")
	case BASS_ERROR_BUFLOST:
		return errors.New("the sample buffer was lost")
	case BASS_ERROR_HANDLE:
		return errors.New("invalid handle")
	case BASS_ERROR_FORMAT:
		return errors.New("unsupported sample format")
	case BASS_ERROR_POSITION:
		return errors.New("invalid position")
	case BASS_ERROR_INIT:
		return errors.New("BASS_Init has not been successfully called")
	case BASS_ERROR_START:
		return errors.New("BASS_Start has not been successfully called")
	case BASS_ERROR_SSL:
		return errors.New("SSL/HTTPS support isn't available")
	case BASS_ERROR_ALREADY:
		return errors.New("already initialized/paused/whatever")
	case BASS_ERROR_NOCHAN:
		return errors.New("can't get a free channel")
	case BASS_ERROR_ILLTYPE:
		return errors.New("an illegal type was specified")
	case BASS_ERROR_ILLPARAM:
		return errors.New("an illegal parameter was specified")
	case BASS_ERROR_NO3D:
		return errors.New("no 3D support")
	case BASS_ERROR_NOEAX:
		return errors.New("no EAX support")
	case BASS_ERROR_DEVICE:
		return errors.New("illegal device number")
	case BASS_ERROR_NOPLAY:
		return errors.New("not playing")
	case BASS_ERROR_FREQ:
		return errors.New("illegal sample rate")
	case BASS_ERROR_NOTFILE:
		return errors.New("the stream is not a file stream")
	case BASS_ERROR_NOHW:
		return errors.New("no hardware voices available")
	case BASS_ERROR_EMPTY:
		return errors.New("the MOD music has no sequence data")
	case BASS_ERROR_NONET:
		return errors.New("no internet connection could be opened")
	case BASS_ERROR_CREATE:
		return errors.New("couldn't create the file")
	case BASS_ERROR_NOFX:
		return errors.New("effects are not available")
	case BASS_ERROR_NOTAVAIL:
		return errors.New("requested data is not available")
	case BASS_ERROR_DECODE:
		return errors.New("the channel is/isn't a \"decoding channel\"")
	case BASS_ERROR_DX:
		return errors.New("a sufficient DirectX version is not installed")
	case BASS_ERROR_TIMEOUT:
		return errors.New("connection timedout")
	case BASS_ERROR_FILEFORM:
		return errors.New("unsupported file format")
	case BASS_ERROR_SPEAKER:
		return errors.New("unavailable speaker")
	case BASS_ERROR_VERSION:
		return errors.New("invalid BASS version (used by add-ons)")
	case BASS_ERROR_CODEC:
		return errors.New("codec is not available/supported")
	case BASS_ERROR_ENDED:
		return errors.New("the channel/file has ended")
	case BASS_ERROR_BUSY:
		return errors.New("the device is busy")
	case BASS_ERROR_UNKNOWN:
		return errors.New("some other mystery problem")
	default:
		return errors.New("some other mystery problem")
	}
}
