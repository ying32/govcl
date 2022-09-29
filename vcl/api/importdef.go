//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package api

var (
	//application_Instance   = newDLLProc("Application_Instance")
	//application_CreateForm = newDLLProc("Application_CreateForm")
	//application_Run        = newDLLProc("Application_Run")
	//application_Initialize = newDLLProc("Application_Initialize")
	//
	//form_Create2      = newDLLProc("Form_Create2")
	//form_SetOnWndProc = newDLLProc("Form_SetOnWndProc")
	//form_SetGoPtr     = newDLLProc("Form_SetGoPtr")

	setEventCallback                   = newDLLProc("SetEventCallback")
	setMessageCallback                 = newDLLProc("SetMessageCallback")
	setThreadSyncCallback              = newDLLProc("SetThreadSyncCallback")
	setRequestCallCreateParamsCallback = newDLLProc("SetRequestCallCreateParamsCallback")
	setRemoveEventCallback             = newDLLProc("SetRemoveEventCallback")

	dGetStringArrOf = newDLLProc("DGetStringArrOf")
	dStrLen         = newDLLProc("DStrLen")
	dMove           = newDLLProc("DMove")

	dShowMessage = newDLLProc("DShowMessage")
	dMessageDlg  = newDLLProc("DMessageDlg")

	//mouse_Instance  = newDLLProc("Mouse_Instance")
	//screen_Instance = newDLLProc("Screen_Instance")

	dSynchronize = newDLLProc("DSynchronize")

	// TMenuItem
	dTextToShortCut = newDLLProc("DTextToShortCut")
	dShortCutToText = newDLLProc("DShortCutToText")

	// TClipboard
	clipboard_Instance         = newDLLProc("Clipboard_Instance")
	clipboard_HasFormat        = newDLLProc("Clipboard_HasFormat")
	dSetClipboard              = newDLLProc("DSetClipboard")
	dRegisterClipboardFormat   = newDLLProc("DRegisterClipboardFormat")
	clipboard_GetAsHtml        = newDLLProc("Clipboard_GetAsHtml")
	clipboard_GetTextBuf       = newDLLProc("Clipboard_GetTextBuf")
	clipboard_GetAsText        = newDLLProc("Clipboard_GetAsText")
	clipboard_SetAsText        = newDLLProc("Clipboard_SetAsText")
	dPredefinedClipboardFormat = newDLLProc("DPredefinedClipboardFormat")

	// DSysOpen
	dSysOpen = newDLLProc("DSysOpen")

	// TMemoryStream
	//memoryStream_Read  = newDLLProc("MemoryStream_Read")
	//memoryStream_Write = newDLLProc("MemoryStream_Write")

	// TCanvas
	canvas_BrushCopy     = newDLLProc("Canvas_BrushCopy")
	canvas_CopyRect      = newDLLProc("Canvas_CopyRect")
	canvas_Draw1         = newDLLProc("Canvas_Draw1")
	canvas_Draw2         = newDLLProc("Canvas_Draw2")
	canvas_DrawFocusRect = newDLLProc("Canvas_DrawFocusRect")
	canvas_FillRect      = newDLLProc("Canvas_FillRect")
	canvas_FrameRect     = newDLLProc("Canvas_FrameRect")
	canvas_TextRect1     = newDLLProc("Canvas_TextRect1")
	canvas_TextRect2     = newDLLProc("Canvas_TextRect2")
	canvas_Polygon       = newDLLProc("Canvas_Polygon")
	canvas_Polyline      = newDLLProc("Canvas_Polyline")
	canvas_PolyBezier    = newDLLProc("Canvas_PolyBezier")

	// TImageList
	imageList_Draw1        = newDLLProc("ImageList_Draw1")
	imageList_Draw2        = newDLLProc("ImageList_Draw2")
	imageList_DrawOverlay1 = newDLLProc("ImageList_DrawOverlay1")
	imageList_DrawOverlay2 = newDLLProc("ImageList_DrawOverlay2")
	imageList_GetIcon1     = newDLLProc("ImageList_GetIcon1")
	imageList_GetIcon2     = newDLLProc("ImageList_GetIcon2")

	dExtractFilePath = newDLLProc("DExtractFilePath")
	dFileExists      = newDLLProc("DFileExists")

	dSelectDirectory1 = newDLLProc("DSelectDirectory1")
	dSelectDirectory2 = newDLLProc("DSelectDirectory2")
	dInputBox         = newDLLProc("DInputBox")
	dInputQuery       = newDLLProc("DInputQuery")
	dPasswordBox      = newDLLProc("DPasswordBox")
	dInputCombo       = newDLLProc("DInputCombo")
	dInputComboEx     = newDLLProc("DInputComboEx")

	// TSysLocaled
	dSysLocale = newDLLProc("DSysLocale")

	// Shortcut
	dCreateURLShortCut = newDLLProc("DCreateURLShortCut")
	dCreateShortCut    = newDLLProc("DCreateShortCut")

	// SetProperty
	dSetPropertyValue    = newDLLProc("DSetPropertyValue")
	dSetPropertySecValue = newDLLProc("DSetPropertySecValue")

	// Printer
	//printer_Instance   = newDLLProc("Printer_Instance")
	//printer_SetPrinter = newDLLProc("Printer_SetPrinter")

	// guid
	dGUIDToString = newDLLProc("DGUIDToString")
	dStringToGUID = newDLLProc("DStringToGUID")
	dCreateGUID   = newDLLProc("DCreateGUID")

	// libResources
	dGetLibResourceCount = newDLLProc("DGetLibResourceCount")
	dGetLibResourceItem  = newDLLProc("DGetLibResourceItem")
	dModifyLibResource   = newDLLProc("DModifyLibResource")
	dLibAbout            = newDLLProc("DLibAbout")

	// 库的信息
	dLibStringEncoding = newDLLProc("DLibStringEncoding")
	dLibVersion        = newDLLProc("DLibVersion")

	dMainThreadId    = newDLLProc("DMainThreadId")
	dCurrentThreadId = newDLLProc("DCurrentThreadId")

	dInitGoDll = newDLLProc("DInitGoDll")

	dFindControl           = newDLLProc("DFindControl")
	dFindLCLControl        = newDLLProc("DFindLCLControl")
	dFindOwnerControl      = newDLLProc("DFindOwnerControl")
	dFindControlAtPosition = newDLLProc("DFindControlAtPosition")
	dFindLCLWindow         = newDLLProc("DFindLCLWindow")
	dFindDragTarget        = newDLLProc("DFindDragTarget")
)
