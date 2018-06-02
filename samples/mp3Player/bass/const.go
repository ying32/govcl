package bass

const (
	// BASS_ChannelGetLength/GetPosition/SetPosition modes
	BASS_POS_BYTE        = 0          // byte position
	BASS_POS_MUSIC_ORDER = 1          // order.row position, MAKELONG(order,row)
	BASS_POS_OGG         = 3          // OGG bitstream number
	BASS_POS_INEXACT     = 0x8000000  // flag: allow seeking to inexact position
	BASS_POS_DECODE      = 0x10000000 // flag: get the decoding (not playing) position
	BASS_POS_DECODETO    = 0x20000000 // flag: decode to the position instead of seeking
	BASS_POS_SCAN        = 0x40000000 // flag: scan to the position

	// Channel attributes
	BASS_ATTRIB_FREQ = 1
	BASS_ATTRIB_VOL  = 2

	// Error codes returned by BASS_ErrorGetCode
	BASS_OK             = 0  // all is OK
	BASS_ERROR_MEM      = 1  // memory error
	BASS_ERROR_FILEOPEN = 2  // can't open the file
	BASS_ERROR_DRIVER   = 3  // can't find a free/valid driver
	BASS_ERROR_BUFLOST  = 4  // the sample buffer was lost
	BASS_ERROR_HANDLE   = 5  // invalid handle
	BASS_ERROR_FORMAT   = 6  // unsupported sample format
	BASS_ERROR_POSITION = 7  // invalid position
	BASS_ERROR_INIT     = 8  // BASS_Init has not been successfully called
	BASS_ERROR_START    = 9  // BASS_Start has not been successfully called
	BASS_ERROR_SSL      = 10 // SSL/HTTPS support isn't available
	BASS_ERROR_ALREADY  = 14 // already initialized/paused/whatever
	BASS_ERROR_NOCHAN   = 18 // can't get a free channel
	BASS_ERROR_ILLTYPE  = 19 // an illegal type was specified
	BASS_ERROR_ILLPARAM = 20 // an illegal parameter was specified
	BASS_ERROR_NO3D     = 21 // no 3D support
	BASS_ERROR_NOEAX    = 22 // no EAX support
	BASS_ERROR_DEVICE   = 23 // illegal device number
	BASS_ERROR_NOPLAY   = 24 // not playing
	BASS_ERROR_FREQ     = 25 // illegal sample rate
	BASS_ERROR_NOTFILE  = 27 // the stream is not a file stream
	BASS_ERROR_NOHW     = 29 // no hardware voices available
	BASS_ERROR_EMPTY    = 31 // the MOD music has no sequence data
	BASS_ERROR_NONET    = 32 // no internet connection could be opened
	BASS_ERROR_CREATE   = 33 // couldn't create the file
	BASS_ERROR_NOFX     = 34 // effects are not available
	BASS_ERROR_NOTAVAIL = 37 // requested data is not available
	BASS_ERROR_DECODE   = 38 // the channel is/isn't a "decoding channel"
	BASS_ERROR_DX       = 39 // a sufficient DirectX version is not installed
	BASS_ERROR_TIMEOUT  = 40 // connection timedout
	BASS_ERROR_FILEFORM = 41 // unsupported file format
	BASS_ERROR_SPEAKER  = 42 // unavailable speaker
	BASS_ERROR_VERSION  = 43 // invalid BASS version (used by add-ons)
	BASS_ERROR_CODEC    = 44 // codec is not available/supported
	BASS_ERROR_ENDED    = 45 // the channel/file has ended
	BASS_ERROR_BUSY     = 46 // the device is busy
	BASS_ERROR_UNKNOWN  = -1 // some other mystery problem
)
