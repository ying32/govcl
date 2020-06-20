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
	application_GetHandle  = libvcl.NewProc("Application_GetHandle")
	application_SetHandle  = libvcl.NewProc("Application_SetHandle")

	form_Create2                = libvcl.NewProc("Form_Create2")
	form_EnabledMaximize        = libvcl.NewProc("Form_EnabledMaximize")
	form_EnabledMinimize        = libvcl.NewProc("Form_EnabledMinimize")
	form_EnabledSystemMenu      = libvcl.NewProc("Form_EnabledSystemMenu")
	form_SetAllowDropFiles      = libvcl.NewProc("Form_SetAllowDropFiles")
	form_GetAllowDropFiles      = libvcl.NewProc("Form_GetAllowDropFiles")
	form_SetOnDropFiles         = libvcl.NewProc("Form_SetOnDropFiles")
	form_SetOnDestroy           = libvcl.NewProc("Form_SetOnDestroy")
	form_SetOnConstrainedResize = libvcl.NewProc("Form_SetOnConstrainedResize")
	form_SetOnDeactivate        = libvcl.NewProc("Form_SetOnDeactivate")
	form_SetOnActivate          = libvcl.NewProc("Form_SetOnActivate")
	form_SetOnStyleChanged      = libvcl.NewProc("Form_SetOnStyleChanged")
	form_SetOnWndProc           = libvcl.NewProc("Form_SetOnWndProc")
	form_SetShowInTaskBar       = libvcl.NewProc("Form_SetShowInTaskBar")
	form_ShowInTaskBar          = libvcl.NewProc("Form_ShowInTaskBar")
	form_ScaleForCurrentDpi     = libvcl.NewProc("Form_ScaleForCurrentDpi")
	form_InheritedWndProc       = libvcl.NewProc("Form_InheritedWndProc")

	setEventCallback      = libvcl.NewProc("SetEventCallback")
	setMessageCallback    = libvcl.NewProc("SetMessageCallback")
	setThreadSyncCallback = libvcl.NewProc("SetThreadSyncCallback")

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
	clipboard_Instance     = libvcl.NewProc("Clipboard_Instance")
	clipboard_SetClipboard = libvcl.NewProc("Clipboard_SetClipboard")

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
	canvas_StretchDraw   = libvcl.NewProc("Canvas_StretchDraw")
	canvas_TextRect1     = libvcl.NewProc("Canvas_TextRect1")
	canvas_TextRect2     = libvcl.NewProc("Canvas_TextRect2")
	canvas_Polygon       = libvcl.NewProc("Canvas_Polygon")
	canvas_Polyline      = libvcl.NewProc("Canvas_Polyline")
	canvas_PolyBezier    = libvcl.NewProc("Canvas_PolyBezier")
	canvas_PolyBezierTo  = libvcl.NewProc("Canvas_PolyBezierTo")
	canvas_Pixels        = libvcl.NewProc("Canvas_Pixels")
	canvas_SetPixels     = libvcl.NewProc("Canvas_SetPixels")

	// TImageList
	imageList_Draw1        = libvcl.NewProc("ImageList_Draw1")
	imageList_Draw2        = libvcl.NewProc("ImageList_Draw2")
	imageList_DrawOverlay1 = libvcl.NewProc("ImageList_DrawOverlay1")
	imageList_DrawOverlay2 = libvcl.NewProc("ImageList_DrawOverlay2")
	imageList_GetIcon1     = libvcl.NewProc("ImageList_GetIcon1")
	imageList_GetIcon2     = libvcl.NewProc("ImageList_GetIcon2")

	// TBitmap
	bitmap_Clear          = libvcl.NewProc("Bitmap_Clear")
	bitmap_BeginUpdate    = libvcl.NewProc("Bitmap_BeginUpdate")
	bitmap_EndUpdate      = libvcl.NewProc("Bitmap_EndUpdate")
	bitmap_LoadFromDevice = libvcl.NewProc("Bitmap_LoadFromDevice")

	dExtractFilePath = libvcl.NewProc("DExtractFilePath")
	dFileExists      = libvcl.NewProc("DFileExists")

	dSelectDirectory1 = libvcl.NewProc("DSelectDirectory1")
	dSelectDirectory2 = libvcl.NewProc("DSelectDirectory2")
	dInputBox         = libvcl.NewProc("DInputBox")
	dInputQuery       = libvcl.NewProc("DInputQuery")

	// TForm相关设置
	form_ScaleForPPI         = libvcl.NewProc("Form_ScaleForPPI")
	form_ScaleControlsForDpi = libvcl.NewProc("Form_ScaleControlsForDpi")

	// TSysLocaled
	dSysLocale = libvcl.NewProc("DSysLocale")

	// Shortcut
	dCreateURLShortCut = libvcl.NewProc("DCreateURLShortCut")
	dCreateShortCut    = libvcl.NewProc("DCreateShortCut")

	// SetProperty
	dSetPropertyValue    = libvcl.NewProc("DSetPropertyValue")
	dSetPropertySecValue = libvcl.NewProc("DSetPropertySecValue")

	// Printer
	printer_Instance = libvcl.NewProc("Printer_Instance")

	// guid
	dGUIDToString = libvcl.NewProc("DGUIDToString")
	dStringToGUID = libvcl.NewProc("DStringToGUID")
	dCreateGUID   = libvcl.NewProc("DCreateGUID")

	// libResouces
	dGetLibResouceCount = libvcl.NewProc("DGetLibResouceCount")
	dGetLibResouceItem  = libvcl.NewProc("DGetLibResouceItem")
	dModifyLibResouce   = libvcl.NewProc("DModifyLibResouce")
	dLibAbout           = libvcl.NewProc("DLibAbout")

	// 库的信息
	dLibStringEncoding = libvcl.NewProc("DLibStringEncoding")
	dLibVersion        = libvcl.NewProc("DLibVersion")

	dMainThreadId    = libvcl.NewProc("DMainThreadId")
	dCurrentThreadId = libvcl.NewProc("DCurrentThreadId")

	dInitGoDll = libvcl.NewProc("DInitGoDll")
)
