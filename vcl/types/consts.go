//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package types

// MessageBox or MessageDlg 返回值
const (
	IdOK       = 1
	IdCancel   = 2
	IdAbort    = 3
	IdRetry    = 4
	IdIgnore   = 5
	IdYes      = 6
	IdNo       = 7
	IdClose    = 8
	IdHelp     = 9
	IdTryAgain = 10
	IdContinue = 11
	MrNone     = 0
	MrOk       = IdOK
	MrCancel   = IdCancel
	MrAbort    = IdAbort
	MrRetry    = IdRetry
	MrIgnore   = IdIgnore
	MrYes      = IdYes
	MrNo       = IdNo
	MrClose    = IdClose
	MrHelp     = IdHelp
	MrTryAgain = IdTryAgain
	MrContinue = IdContinue
	MrAll      = MrContinue + 1
	MrNoToAll  = MrAll + 1
	MrYesToAll = MrNoToAll + 1
)

// Predefined Clipboard Formats
const (
	CF_BITMAP          = 2
	CF_DIB             = 8
	CF_PALETTE         = 9
	CF_ENHMETAFILE     = 14
	CF_METAFILEPICT    = 3
	CF_OEMTEXT         = 7
	CF_TEXT            = 1
	CF_UNICODETEXT     = 13
	CF_DIF             = 5
	CF_DSPBITMAP       = 130
	CF_DSPENHMETAFILE  = 142
	CF_DSPMETAFILEPICT = 131
	CF_DSPTEXT         = 129
	CF_GDIOBJFIRST     = 768
	CF_GDIOBJLAST      = 1023
	CF_HDROP           = 15
	CF_LOCALE          = 16
	CF_OWNERDISPLAY    = 128
	CF_PENDATA         = 10
	CF_PRIVATEFIRST    = 512
	CF_PRIVATELAST     = 767
	CF_RIFF            = 11
	CF_SYLK            = 4
	CF_WAVE            = 12
	CF_TIFF            = 6

	// custom
	CF_PICTURE   = 700
	CF_HTML      = 701
	CF_COMPONENT = 702
)
