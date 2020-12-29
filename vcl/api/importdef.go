//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package api

var (
	application_Instance   = libvcl.NewProc("Application_Instance")
	application_CreateForm = libvcl.NewProc("Application_CreateForm")
	application_Run        = libvcl.NewProc("Application_Run")
	application_Initialize = libvcl.NewProc("Application_Initialize")

	form_Create2      = libvcl.NewProc("Form_Create2")
	form_SetOnWndProc = libvcl.NewProc("Form_SetOnWndProc")
	form_SetGoPtr     = libvcl.NewProc("Form_SetGoPtr")

	setEventCallback                   = libvcl.NewProc("SetEventCallback")
	setMessageCallback                 = libvcl.NewProc("SetMessageCallback")
	setThreadSyncCallback              = libvcl.NewProc("SetThreadSyncCallback")
	setRequestCallCreateParamsCallback = libvcl.NewProc("SetRequestCallCreateParamsCallback")

	dGetStringArrOf = libvcl.NewProc("DGetStringArrOf")
	dStrLen         = libvcl.NewProc("DStrLen")
	dMove           = libvcl.NewProc("DMove")

	dShowMessage = libvcl.NewProc("DShowMessage")
	dMessageDlg  = libvcl.NewProc("DMessageDlg")

	mouse_Instance  = libvcl.NewProc("Mouse_Instance")
	screen_Instance = libvcl.NewProc("Screen_Instance")

	dSynchronize = libvcl.NewProc("DSynchronize")

	// TMenuItem
	dTextToShortCut = libvcl.NewProc("DTextToShortCut")
	dShortCutToText = libvcl.NewProc("DShortCutToText")

	// TClipboard
	clipboard_Instance         = libvcl.NewProc("Clipboard_Instance")
	clipboard_HasFormat        = libvcl.NewProc("Clipboard_HasFormat")
	dSetClipboard              = libvcl.NewProc("DSetClipboard")
	dRegisterClipboardFormat   = libvcl.NewProc("DRegisterClipboardFormat")
	clipboard_GetAsHtml        = libvcl.NewProc("Clipboard_GetAsHtml")
	clipboard_GetTextBuf       = libvcl.NewProc("Clipboard_GetTextBuf")
	clipboard_GetAsText        = libvcl.NewProc("Clipboard_GetAsText")
	clipboard_SetAsText        = libvcl.NewProc("Clipboard_SetAsText")
	dPredefinedClipboardFormat = libvcl.NewProc("DPredefinedClipboardFormat")

	// DSysOpen
	dSysOpen = libvcl.NewProc("DSysOpen")

	// TMemoryStream
	memoryStream_Read  = libvcl.NewProc("MemoryStream_Read")
	memoryStream_Write = libvcl.NewProc("MemoryStream_Write")

	// TCanvas
	canvas_BrushCopy     = libvcl.NewProc("Canvas_BrushCopy")
	canvas_CopyRect      = libvcl.NewProc("Canvas_CopyRect")
	canvas_Draw1         = libvcl.NewProc("Canvas_Draw1")
	canvas_Draw2         = libvcl.NewProc("Canvas_Draw2")
	canvas_DrawFocusRect = libvcl.NewProc("Canvas_DrawFocusRect")
	canvas_FillRect      = libvcl.NewProc("Canvas_FillRect")
	canvas_FrameRect     = libvcl.NewProc("Canvas_FrameRect")
	canvas_TextRect1     = libvcl.NewProc("Canvas_TextRect1")
	canvas_TextRect2     = libvcl.NewProc("Canvas_TextRect2")
	canvas_Polygon       = libvcl.NewProc("Canvas_Polygon")
	canvas_Polyline      = libvcl.NewProc("Canvas_Polyline")
	canvas_PolyBezier    = libvcl.NewProc("Canvas_PolyBezier")

	// TImageList
	imageList_Draw1        = libvcl.NewProc("ImageList_Draw1")
	imageList_Draw2        = libvcl.NewProc("ImageList_Draw2")
	imageList_DrawOverlay1 = libvcl.NewProc("ImageList_DrawOverlay1")
	imageList_DrawOverlay2 = libvcl.NewProc("ImageList_DrawOverlay2")
	imageList_GetIcon1     = libvcl.NewProc("ImageList_GetIcon1")
	imageList_GetIcon2     = libvcl.NewProc("ImageList_GetIcon2")

	dExtractFilePath = libvcl.NewProc("DExtractFilePath")
	dFileExists      = libvcl.NewProc("DFileExists")

	dSelectDirectory1 = libvcl.NewProc("DSelectDirectory1")
	dSelectDirectory2 = libvcl.NewProc("DSelectDirectory2")
	dInputBox         = libvcl.NewProc("DInputBox")
	dInputQuery       = libvcl.NewProc("DInputQuery")
	dPasswordBox      = libvcl.NewProc("DPasswordBox")
	dInputCombo       = libvcl.NewProc("DInputCombo")
	dInputComboEx     = libvcl.NewProc("DInputComboEx")

	// TSysLocaled
	dSysLocale = libvcl.NewProc("DSysLocale")

	// Shortcut
	dCreateURLShortCut = libvcl.NewProc("DCreateURLShortCut")
	dCreateShortCut    = libvcl.NewProc("DCreateShortCut")

	// SetProperty
	dSetPropertyValue    = libvcl.NewProc("DSetPropertyValue")
	dSetPropertySecValue = libvcl.NewProc("DSetPropertySecValue")

	// Printer
	printer_Instance   = libvcl.NewProc("Printer_Instance")
	printer_SetPrinter = libvcl.NewProc("Printer_SetPrinter")

	// guid
	dGUIDToString = libvcl.NewProc("DGUIDToString")
	dStringToGUID = libvcl.NewProc("DStringToGUID")
	dCreateGUID   = libvcl.NewProc("DCreateGUID")

	// libResources
	dGetLibResourceCount = libvcl.NewProc("DGetLibResourceCount")
	dGetLibResourceItem  = libvcl.NewProc("DGetLibResourceItem")
	dModifyLibResource   = libvcl.NewProc("DModifyLibResource")
	dLibAbout            = libvcl.NewProc("DLibAbout")

	// 库的信息
	dLibStringEncoding = libvcl.NewProc("DLibStringEncoding")
	dLibVersion        = libvcl.NewProc("DLibVersion")

	dMainThreadId    = libvcl.NewProc("DMainThreadId")
	dCurrentThreadId = libvcl.NewProc("DCurrentThreadId")

	dInitGoDll = libvcl.NewProc("DInitGoDll")

	dFindControl           = libvcl.NewProc("DFindControl")
	dFindLCLControl        = libvcl.NewProc("DFindLCLControl")
	dFindOwnerControl      = libvcl.NewProc("DFindOwnerControl")
	dFindControlAtPosition = libvcl.NewProc("DFindControlAtPosition")
	dFindLCLWindow         = libvcl.NewProc("DFindLCLWindow")
	dFindDragTarget        = libvcl.NewProc("DFindDragTarget")
)
