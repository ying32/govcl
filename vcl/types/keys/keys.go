//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package keys

// 虚拟键
const (
	// Virtual Keys, Standard Set
	VkLButton    = 0x01 //   1
	VkRButton    = 0x02 //   2
	VkCancel     = 0x03 //   3
	VkMButton    = 0x04 //   4
	VkXButton1   = 0x05 //   5
	VkXButton2   = 0x06 //   6
	VkBack       = 0x08 //   8
	VkTab        = 0x09 //   9
	VkLineFeed   = 0x0A //  10
	VkClear      = 0x0C //  12
	VkReturn     = 0x0D //  13
	VkShift      = 0x10 //  16
	VkControl    = 0x11 //  17
	VkMenu       = 0x12 //  18
	VkPause      = 0x13 //  19
	VkCapital    = 0x14 //  20
	VkKana       = 0x15 //  21
	VkHangul     = 0x15 //  21
	VkJunja      = 0x17 //  23
	VkFinal      = 0x18 //  24
	VkHanja      = 0x19 //  25
	VkKanji      = 0x19 //  25
	VkConvert    = 0x1C //  28
	VkNonConvert = 0x1D //  29
	VkAccept     = 0x1E //  30
	VkModeChange = 0x1F //  31
	VkEscape     = 0x1B //  27
	VkSpace      = 0x20 //  32
	VkPrior      = 0x21 //  33
	VkNext       = 0x22 //  34
	VkEnd        = 0x23 //  35
	VkHome       = 0x24 //  36
	VkLeft       = 0x25 //  37
	VkUp         = 0x26 //  38
	VkRight      = 0x27 //  39
	VkDown       = 0x28 //  40
	VkSelect     = 0x29 //  41
	VkPrint      = 0x2A //  42
	VkExecute    = 0x2B //  43
	VkSnapshot   = 0x2C //  44
	VkInsert     = 0x2D //  45
	VkDelete     = 0x2E //  46
	VkHelp       = 0x2F //  47
	// vk0 thru vk9 are the same as ASCII '0' thru '9' (0x30 - 0x39)
	Vk0 = 0x30 //  48
	Vk1 = 0x31 //  49
	Vk2 = 0x32 //  50
	Vk3 = 0x33 //  51
	Vk4 = 0x34 //  52
	Vk5 = 0x35 //  53
	Vk6 = 0x36 //  54
	Vk7 = 0x37 //  55
	Vk8 = 0x38 //  56
	Vk9 = 0x39 //  57
	// vkA thru vkZ are the same as ASCII 'A' thru 'Z' (0x41 - 0x5A)
	VkA         = 0x41 //  65
	VkB         = 0x42 //  66
	VkC         = 0x43 //  67
	VkD         = 0x44 //  68
	VkE         = 0x45 //  69
	VkF         = 0x46 //  70
	VkG         = 0x47 //  71
	VkH         = 0x48 //  72
	VkI         = 0x49 //  73
	VkJ         = 0x4A //  74
	VkK         = 0x4B //  75
	VkL         = 0x4C //  76
	VkM         = 0x4D //  77
	VkN         = 0x4E //  78
	VkO         = 0x4F //  79
	VkP         = 0x50 //  80
	VkQ         = 0x51 //  81
	VkR         = 0x52 //  82
	VkS         = 0x53 //  83
	VkT         = 0x54 //  84
	VkU         = 0x55 //  85
	VkV         = 0x56 //  86
	VkW         = 0x57 //  87
	VkX         = 0x58 //  88
	VkY         = 0x59 //  89
	VkZ         = 0x5A //  90
	VkLWin      = 0x5B //  91
	VkRWin      = 0x5C //  92
	VkApps      = 0x5D //  93
	VkSleep     = 0x5F //  95
	VkNumpad0   = 0x60 //  96
	VkNumpad1   = 0x61 //  97
	VkNumpad2   = 0x62 //  98
	VkNumpad3   = 0x63 //  99
	VkNumpad4   = 0x64 // 100
	VkNumpad5   = 0x65 // 101
	VkNumpad6   = 0x66 // 102
	VkNumpad7   = 0x67 // 103
	VkNumpad8   = 0x68 // 104
	VkNumpad9   = 0x69 // 105
	VkMultiply  = 0x6A // 106
	VkAdd       = 0x6B // 107
	VkSeparator = 0x6C // 108
	VkSubtract  = 0x6D // 109
	VkDecimal   = 0x6E // 110
	VkDivide    = 0x6F // 111
	VkF1        = 0x70 // 112
	VkF2        = 0x71 // 113
	VkF3        = 0x72 // 114
	VkF4        = 0x73 // 115
	VkF5        = 0x74 // 116
	VkF6        = 0x75 // 117
	VkF7        = 0x76 // 118
	VkF8        = 0x77 // 119
	VkF9        = 0x78 // 120
	VkF10       = 0x79 // 121
	VkF11       = 0x7A // 122
	VkF12       = 0x7B // 123
	VkF13       = 0x7C // 124
	VkF14       = 0x7D // 125
	VkF15       = 0x7E // 126
	VkF16       = 0x7F // 127
	VkF17       = 0x80 // 128
	VkF18       = 0x81 // 129
	VkF19       = 0x82 // 130
	VkF20       = 0x83 // 131
	VkF21       = 0x84 // 132
	VkF22       = 0x85 // 133
	VkF23       = 0x86 // 134
	VkF24       = 0x87 // 135

	VkCamera       = 0x88 // 136
	VkHardwareBack = 0x89 // 137

	VkNumLock  = 0x90 // 144
	VkScroll   = 0x91 // 145
	VkLShift   = 0xA0 // 160
	VkRShift   = 0xA1 // 161
	VkLControl = 0xA2 // 162
	VkRControl = 0xA3 // 163
	VkLMenu    = 0xA4 // 164
	VkRMenu    = 0xA5 // 165

	VkBrowserBack       = 0xA6 // 166
	VkBrowserForward    = 0xA7 // 167
	VkBrowserRefresh    = 0xA8 // 168
	VkBrowserStop       = 0xA9 // 169
	VkBrowserSearch     = 0xAA // 170
	VkBrowserFavorites  = 0xAB // 171
	VkBrowserHome       = 0xAC // 172
	VkVolumeMute        = 0xAD // 173
	VkVolumeDown        = 0xAE // 174
	VkVolumeUp          = 0xAF // 175
	VkMediaNextTrack    = 0xB0 // 176
	VkMediaPrevTrack    = 0xB1 // 177
	VkMediaStop         = 0xB2 // 178
	VkMediaPlayPause    = 0xB3 // 179
	VkLaunchMail        = 0xB4 // 180
	VkLaunchMediaSelect = 0xB5 // 181
	VkLaunchApp1        = 0xB6 // 182
	VkLaunchApp2        = 0xB7 // 183

	VkSemicolon    = 0xBA // 186
	VkEqual        = 0xBB // 187
	VkComma        = 0xBC // 188
	VkMinus        = 0xBD // 189
	VkPeriod       = 0xBE // 190
	VkSlash        = 0xBF // 191
	VkTilde        = 0xC0 // 192
	VkLeftBracket  = 0xDB // 219
	VkBackslash    = 0xDC // 220
	VkRightBracket = 0xDD // 221
	VkQuote        = 0xDE // 222
	VkPara         = 0xDF // 223

	VkOem102     = 0xE2 // 226
	VkIcoHelp    = 0xE3 // 227
	VkIco00      = 0xE4 // 228
	VkProcessKey = 0xE5 // 229
	VkIcoClear   = 0xE6 // 230
	VkPacket     = 0xE7 // 231
	VkAttn       = 0xF6 // 246
	VkCrsel      = 0xF7 // 247
	VkExsel      = 0xF8 // 248
	VkErEof      = 0xF9 // 249
	VkPlay       = 0xFA // 250
	VkZoom       = 0xFB // 251
	VkNoname     = 0xFC // 252
	VkPA1        = 0xFD // 253
	VkOemClear   = 0xFE // 254
	VkNone       = 0xFF // 255

)
