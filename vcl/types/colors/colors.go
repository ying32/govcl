//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package colors

// Lazarus中的颜色表，Lazarus中的TColor与一般HTML的RGB有点不一样，反过来的以BGR表示.

// 原Delphi中的定义
const (
	ClClSysNone  = 0x1FFFFFFF
	ClSysDefault = 0x20000000
	// Actual colors
	ClAliceblue            = 0xFFF8F0
	ClAntiquewhite         = 0xD7EBFA
	ClAqua                 = 0xFFFF00
	ClAquamarine           = 0xD4FF7F
	ClAzure                = 0xFFFFF0
	ClBeige                = 0xDCF5F5
	ClBisque               = 0xC4E4FF
	ClBlack                = 0x000000
	ClBlanchedalmond       = 0xCDEBFF
	ClBlue                 = 0xFF0000
	ClBlueviolet           = 0xE22B8A
	ClBrown                = 0x2A2AA5
	ClBurlywood            = 0x87B8DE
	ClCadetblue            = 0xA09E5F
	ClChartreuse           = 0x00FF7F
	ClChocolate            = 0x1E69D2
	ClCoral                = 0x507FFF
	ClCornflowerblue       = 0xED9564
	ClCornsilk             = 0xDCF8FF
	ClCrimson              = 0x3C14DC
	ClCyan                 = 0xFFFF00
	ClDarkblue             = 0x8B0000
	ClDarkcyan             = 0x8B8B00
	ClDarkgoldenrod        = 0x0B86B8
	ClDarkgray             = 0xA9A9A9
	ClDarkgreen            = 0x006400
	ClDarkgrey             = 0xA9A9A9
	ClDarkkhaki            = 0x6BB7BD
	ClDarkmagenta          = 0x8B008B
	ClDarkolivegreen       = 0x2F6B55
	ClDarkorange           = 0x008CFF
	ClDarkorchid           = 0xCC3299
	ClDarkred              = 0x00008B
	ClDarksalmon           = 0x7A96E9
	ClDarkseagreen         = 0x8FBC8F
	ClDarkslateblue        = 0x8B3D48
	ClDarkslategray        = 0x4F4F2F
	ClDarkslategrey        = 0x4F4F2F
	ClDarkturquoise        = 0xD1CE00
	ClDarkviolet           = 0xD30094
	ClDeeppink             = 0x9314FF
	ClDeepskyblue          = 0xFFBF00
	ClDimgray              = 0x696969
	ClDimgrey              = 0x696969
	ClDodgerblue           = 0xFF901E
	ClFirebrick            = 0x2222B2
	ClFloralwhite          = 0xF0FAFF
	ClForestgreen          = 0x228B22
	ClFuchsia              = 0xFF00FF
	ClGainsboro            = 0xDCDCDC
	ClGhostwhite           = 0xFFF8F8
	ClGold                 = 0x00D7FF
	ClGoldenrod            = 0x20A5DA
	ClGray                 = 0x808080
	ClGreen                = 0x008000
	ClGreenyellow          = 0x2FFFAD
	ClGrey                 = 0x808080
	ClHoneydew             = 0xF0FFF0
	ClHotpink              = 0xB469FF
	ClIndianred            = 0x5C5CCD
	ClIndigo               = 0x82004B
	ClIvory                = 0xF0FFFF
	ClKhaki                = 0x8CE6F0
	ClLavender             = 0xFAE6E6
	ClLavenderblush        = 0xF5F0FF
	ClLawngreen            = 0x00FC7C
	ClLemonchiffon         = 0xCDFAFF
	ClLightblue            = 0xE6D8AD
	ClLightcoral           = 0x8080F0
	ClLightcyan            = 0xFFFFE0
	ClLightgoldenrodyellow = 0xD2FAFA
	ClLightgray            = 0xD3D3D3
	ClLightgreen           = 0x90EE90
	ClLightgrey            = 0xD3D3D3
	ClLightpink            = 0xC1B6FF
	ClLightsalmon          = 0x7AA0FF
	ClLightseagreen        = 0xAAB220
	ClLightskyblue         = 0xFACE87
	ClLightslategray       = 0x998877
	ClLightslategrey       = 0x998877
	ClLightsteelblue       = 0xDEC4B0
	ClLightyellow          = 0xE0FFFF
	ClLtGray               = 0xC0C0C0
	ClMedGray              = 0xA4A0A0
	ClDkGray               = 0x808080
	ClMoneyGreen           = 0xC0DCC0
	ClLegacySkyBlue        = 0xF0CAA6
	ClCream                = 0xF0FBFF
	ClLime                 = 0x00FF00
	ClLimegreen            = 0x32CD32
	ClLinen                = 0xE6F0FA
	ClMagenta              = 0xFF00FF
	ClMaroon               = 0x000080
	ClMediumaquamarine     = 0xAACD66
	ClMediumblue           = 0xCD0000
	ClMediumorchid         = 0xD355BA
	ClMediumpurple         = 0xDB7093
	ClMediumseagreen       = 0x71B33C
	ClMediumslateblue      = 0xEE687B
	ClMediumspringgreen    = 0x9AFA00
	ClMediumturquoise      = 0xCCD148
	ClMediumvioletred      = 0x8515C7
	ClMidnightblue         = 0x701919
	ClMintcream            = 0xFAFFF5
	ClMistyrose            = 0xE1E4FF
	ClMoccasin             = 0xB5E4FF
	ClNavajowhite          = 0xADDEFF
	ClNavy                 = 0x800000
	ClOldlace              = 0xE6F5FD
	ClOlive                = 0x008080
	ClOlivedrab            = 0x238E6B
	ClOrange               = 0x00A5FF
	ClOrangered            = 0x0045FF
	ClOrchid               = 0xD670DA
	ClPalegoldenrod        = 0xAAE8EE
	ClPalegreen            = 0x98FB98
	ClPaleturquoise        = 0xEEEEAF
	ClPalevioletred        = 0x9370DB
	ClPapayawhip           = 0xD5EFFF
	ClPeachpuff            = 0xB9DAFF
	ClPeru                 = 0x3F85CD
	ClPink                 = 0xCBC0FF
	ClPlum                 = 0xDDA0DD
	ClPowderblue           = 0xE6E0B0
	ClPurple               = 0x800080
	ClRed                  = 0x0000FF
	ClRosybrown            = 0x8F8FBC
	ClRoyalblue            = 0xE16941
	ClSaddlebrown          = 0x13458B
	ClSalmon               = 0x7280FA
	ClSandybrown           = 0x60A4F4
	ClSeagreen             = 0x578B2E
	ClSeashell             = 0xEEF5FF
	ClSienna               = 0x2D52A0
	ClSilver               = 0xC0C0C0
	ClSkyblue              = 0xEBCE87
	ClSlateblue            = 0xCD5A6A
	ClSlategray            = 0x908070
	ClSlategrey            = 0x908070
	ClSnow                 = 0xFAFAFF
	ClSpringgreen          = 0x7FFF00
	ClSteelblue            = 0xB48246
	ClTan                  = 0x8CB4D2
	ClTeal                 = 0x808000
	ClThistle              = 0xD8BFD8
	ClTomato               = 0x4763FF
	ClTurquoise            = 0xD0E040
	ClViolet               = 0xEE82EE
	ClWheat                = 0xB3DEF5
	ClWhite                = 0xFFFFFF
	ClWhitesmoke           = 0xF5F5F5
	ClYellow               = 0x00FFFF
	ClYellowgreen          = 0x32CD9A
	ClBtnFace              = 0xFF00000F
	ClNull                 = 0x00000000
)

// Lazarus中的定义
const (
	//CLR_NONE    = 0xFFFFFFFF
	//CLR_DEFAULT = 0xFF000000
	//CLR_INVALID = 0xFFFFFFFF

	cOLOR_SCROLLBAR           = 0
	cOLOR_BACKGROUND          = 1
	cOLOR_ACTIVECAPTION       = 2
	cOLOR_INACTIVECAPTION     = 3
	cOLOR_MENU                = 4
	cOLOR_WINDOW              = 5
	cOLOR_WINDOWFRAME         = 6
	cOLOR_MENUTEXT            = 7
	cOLOR_WINDOWTEXT          = 8
	cOLOR_CAPTIONTEXT         = 9
	cOLOR_ACTIVEBORDER        = 10
	cOLOR_INACTIVEBORDER      = 11
	cOLOR_APPWORKSPACE        = 12
	cOLOR_HIGHLIGHT           = 13
	cOLOR_HIGHLIGHTTEXT       = 14
	cOLOR_BTNFACE             = 15
	cOLOR_BTNSHADOW           = 16
	cOLOR_GRAYTEXT            = 17
	cOLOR_BTNTEXT             = 18
	cOLOR_INACTIVECAPTIONTEXT = 19
	cOLOR_BTNHIGHLIGHT        = 20
	cOLOR_3DDKSHADOW          = 21
	cOLOR_3DLIGHT             = 22
	cOLOR_INFOTEXT            = 23
	cOLOR_INFOBK              = 24
	// PBD: 25 is unassigned in all the docs I can find
	//      if someone finds what this is supposed to be then fill it in
	//      note defaults below, and cl[ColorConst] in graphics
	cOLOR_HOTLIGHT                = 26
	cOLOR_GRADIENTACTIVECAPTION   = 27
	cOLOR_GRADIENTINACTIVECAPTION = 28
	cOLOR_MENUHILIGHT             = 29
	cOLOR_MENUBAR                 = 30
	cOLOR_FORM                    = 31
	cOLOR_ENDCOLORS               = cOLOR_FORM
	cOLOR_DESKTOP                 = cOLOR_BACKGROUND
	cOLOR_3DFACE                  = cOLOR_BTNFACE
	cOLOR_3DSHADOW                = cOLOR_BTNSHADOW
	cOLOR_3DHIGHLIGHT             = cOLOR_BTNHIGHLIGHT
	cOLOR_3DHILIGHT               = cOLOR_BTNHIGHLIGHT
	cOLOR_BTNHILIGHT              = cOLOR_BTNHIGHLIGHT
	mAX_SYS_COLORS                = cOLOR_ENDCOLORS
	sYS_COLOR_BASE                = 0x80000000

	// The following colors match the predefined Delphi Colors
	// standard colors
	//ClBlack   = 0x000000
	//ClMaroon  = 0x000080
	//ClGreen   = 0x008000
	//ClOlive   = 0x008080
	//ClNavy    = 0x800000
	//ClPurple  = 0x800080
	//ClTeal    = 0x808000
	//ClGray    = 0x808080
	//ClSilver  = 0xC0C0C0
	//ClRed     = 0x0000FF
	//ClLime    = 0x00FF00
	//ClYellow  = 0x00FFFF
	//ClBlue    = 0xFF0000
	//ClFuchsia = 0xFF00FF
	//ClAqua    = 0xFFFF00
	//ClLtGray  = 0xC0C0C0 // ClSilver alias
	//ClDkGray  = 0x808080 // ClGray alias
	//ClWhite   = 0xFFFFFF

	// extended colors
	//ClMoneyGreen = 0xC0DCC0
	//ClSkyBlue    = 0xF0CAA6
	//ClCream      = 0xF0FBFF
	//ClMedGray    = 0xA4A0A0

	// special colors
	ClNone    = 0x1FFFFFFF
	ClDefault = 0x20000000

	// system colors
	ClScrollBar       = sYS_COLOR_BASE | cOLOR_SCROLLBAR
	ClBackground      = sYS_COLOR_BASE | cOLOR_BACKGROUND
	ClActiveCaption   = sYS_COLOR_BASE | cOLOR_ACTIVECAPTION
	ClInactiveCaption = sYS_COLOR_BASE | cOLOR_INACTIVECAPTION
	ClMenu            = sYS_COLOR_BASE | cOLOR_MENU
	ClWindow          = sYS_COLOR_BASE | cOLOR_WINDOW
	ClWindowFrame     = sYS_COLOR_BASE | cOLOR_WINDOWFRAME
	ClMenuText        = sYS_COLOR_BASE | cOLOR_MENUTEXT
	ClWindowText      = sYS_COLOR_BASE | cOLOR_WINDOWTEXT
	ClCaptionText     = sYS_COLOR_BASE | cOLOR_CAPTIONTEXT
	ClActiveBorder    = sYS_COLOR_BASE | cOLOR_ACTIVEBORDER
	ClInactiveBorder  = sYS_COLOR_BASE | cOLOR_INACTIVEBORDER
	ClAppWorkspace    = sYS_COLOR_BASE | cOLOR_APPWORKSPACE
	ClHighlight       = sYS_COLOR_BASE | cOLOR_HIGHLIGHT
	ClHighlightText   = sYS_COLOR_BASE | cOLOR_HIGHLIGHTTEXT
	//ClBtnFace             = sYS_COLOR_BASE | cOLOR_BTNFACE
	ClBtnShadow           = sYS_COLOR_BASE | cOLOR_BTNSHADOW
	ClGrayText            = sYS_COLOR_BASE | cOLOR_GRAYTEXT
	ClBtnText             = sYS_COLOR_BASE | cOLOR_BTNTEXT
	ClInactiveCaptionText = sYS_COLOR_BASE | cOLOR_INACTIVECAPTIONTEXT
	ClBtnHighlight        = sYS_COLOR_BASE | cOLOR_BTNHIGHLIGHT
	Cl3DDkShadow          = sYS_COLOR_BASE | cOLOR_3DDKSHADOW
	Cl3DLight             = sYS_COLOR_BASE | cOLOR_3DLIGHT
	ClInfoText            = sYS_COLOR_BASE | cOLOR_INFOTEXT
	ClInfoBk              = sYS_COLOR_BASE | cOLOR_INFOBK

	ClHotLight                = sYS_COLOR_BASE | cOLOR_HOTLIGHT
	ClGradientActiveCaption   = sYS_COLOR_BASE | cOLOR_GRADIENTACTIVECAPTION
	ClGradientInactiveCaption = sYS_COLOR_BASE | cOLOR_GRADIENTINACTIVECAPTION
	ClMenuHighlight           = sYS_COLOR_BASE | cOLOR_MENUHILIGHT
	ClMenuBar                 = sYS_COLOR_BASE | cOLOR_MENUBAR
	ClForm                    = sYS_COLOR_BASE | cOLOR_FORM

	// synonyms: do not show them in color lists
	ClColorDesktop = sYS_COLOR_BASE | cOLOR_DESKTOP
	Cl3DFace       = sYS_COLOR_BASE | cOLOR_3DFACE
	Cl3DShadow     = sYS_COLOR_BASE | cOLOR_3DSHADOW
	Cl3DHiLight    = sYS_COLOR_BASE | cOLOR_3DHIGHLIGHT
	ClBtnHiLight   = sYS_COLOR_BASE | cOLOR_BTNHILIGHT

	ClFirstSpecialColor = ClBtnHiLight
	ClMask              = ClWhite
	ClDontMask          = ClBlack
)

// RGB
func RGB(r, g, b byte) uint32 {
	return uint32(r) | (uint32(g) << 8) | (uint32(b) << 16)
}

// RGBToBGR
func RGBToBGR(rgb uint32) uint32 {
	return uint32(byte(rgb>>16)) | (uint32(byte(rgb>>8)) << 8) | (uint32(byte(rgb)) << 16)
}
