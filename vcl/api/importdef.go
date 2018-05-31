package api

var (
	application_Instance   = libvcl.NewProc("Application_Instance")
	application_CreateForm = libvcl.NewProc("Application_CreateForm")
	application_Run        = libvcl.NewProc("Application_Run")
	application_Initialize = libvcl.NewProc("Application_Initialize")

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

	setEventCallback = libvcl.NewProc("SetEventCallback")
	dGetParam        = libvcl.NewProc("DGetParam")
	dGetStringArrOf  = libvcl.NewProc("DGetStringArrOf")
	dStrLen          = libvcl.NewProc("DStrLen")
	dMove            = libvcl.NewProc("DMove")

	dShowMessage     = libvcl.NewProc("DShowMessage")
	dGetMainInstance = libvcl.NewProc("DGetMainInstance")
	dMessageDlg      = libvcl.NewProc("DMessageDlg")

	mouse_Instance  = libvcl.NewProc("Mouse_Instance")
	screen_Instance = libvcl.NewProc("Screen_Instance")

	dSetReportMemoryLeaksOnShutdown = libvcl.NewProc("DSetReportMemoryLeaksOnShutdown")
	dSynchronize                    = libvcl.NewProc("DSynchronize")

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

	// TImageList
	imageList_Draw1        = libvcl.NewProc("ImageList_Draw1")
	imageList_Draw2        = libvcl.NewProc("ImageList_Draw2")
	imageList_DrawOverlay1 = libvcl.NewProc("ImageList_DrawOverlay1")
	imageList_DrawOverlay2 = libvcl.NewProc("ImageList_DrawOverlay2")
	imageList_GetIcon1     = libvcl.NewProc("ImageList_GetIcon1")
	imageList_GetIcon2     = libvcl.NewProc("ImageList_GetIcon2")

	dExtractFilePath = libvcl.NewProc("DExtractFilePath")
	dFileExists      = libvcl.NewProc("DFileExists")

	dInheritsFromControl    = libvcl.NewProc("DInheritsFromControl")
	dInheritsFromWinControl = libvcl.NewProc("DInheritsFromWinControl")
	dInheritsFromComponent  = libvcl.NewProc("DInheritsFromComponent")

	dSelectDirectory1 = libvcl.NewProc("DSelectDirectory1")
	dSelectDirectory2 = libvcl.NewProc("DSelectDirectory2")
	dInputBox         = libvcl.NewProc("DInputBox")
	dInputQuery       = libvcl.NewProc("DInputQuery")

	// TForm相关设置
	setGlobalFormScaled      = libvcl.NewProc("SetGlobalFormScaled")
	form_ScaleForPPI         = libvcl.NewProc("Form_ScaleForPPI")
	form_ScaleControlsForDpi = libvcl.NewProc("Form_ScaleControlsForDpi")
)
