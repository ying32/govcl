//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package api

import (
	"unsafe"

	"github.com/ying32/govcl/vcl/api/dllimports"
	. "github.com/ying32/govcl/vcl/types"
)

// -------------------- TApplication ---------------------------

func Application_Instance() uintptr {
	return defSyscallN(dllimports.APPLICATION_INSTANCE)
}

func Application_CreateForm(app uintptr) uintptr {
	return defSyscallN(dllimports.APPLICATION_CREATEFORM, app)
}

func Application_Run(app uintptr) {
	defer func() {
		// 开启了finalizerOn选项后，以防止关闭库后GC还没开始调用。
		callGC()
		// 运行结束后就结束close掉lib，不然他不会关掉的
		closeLib()
	}()
	defSyscallN(dllimports.APPLICATION_RUN, app)
}

func Application_Initialize(obj uintptr) {
	defSyscallN(dllimports.APPLICATION_INITIALIZE, obj)
}

// -------------------- TForm ---------------------------

func Form_Create2(owner uintptr) uintptr {
	return defSyscallN(dllimports.FORM_CREATE2, owner)
}

func Form_SetOnWndProc(obj uintptr, fn interface{}) {
	defSyscallN(dllimports.FORM_SETONWNDPROC, obj, MakeEventDataPtr(fn))
}

func Form_SetGoPtr(obj uintptr, ptr uintptr) {
	defSyscallN(dllimports.FORM_SETGOPTR, obj, ptr)
}

// -------------------- Form Resource ---------------------------
//----------------------------------------
// 加载文件或者内存中的窗口资源文件功能
// 需要配合窗口设计器使用
//----------------------------------------

// ResFormLoadFromStream
//
// 从流中加载窗口资源
func ResFormLoadFromStream(obj, root uintptr) {
	defSyscallN(dllimports.RESFORMLOADFROMSTREAM, obj, root)
}

// ResFormLoadFromFile
//
// 从文件中加载窗口资源
func ResFormLoadFromFile(filename string, root uintptr) {
	defSyscallN(dllimports.RESFORMLOADFROMFILE, PascalStr(filename), root)
}

// ResFormLoadFromResourceName
//
// 从指定资源中加载窗口资源
func ResFormLoadFromResourceName(instance uintptr, resName string, root uintptr) {
	defSyscallN(dllimports.RESFORMLOADFROMRESOURCENAME, instance, PascalStr(resName), root)
}

// -------------------- TCanvas ---------------------------

func Canvas_BrushCopy(obj uintptr, dest TRect, bitmap uintptr, source TRect, color TColor) {
	defSyscallN(dllimports.CANVAS_BRUSHCOPY, obj, uintptr(unsafe.Pointer(&dest)), bitmap, uintptr(unsafe.Pointer(&source)), uintptr(color))
}

func Canvas_CopyRect(obj uintptr, dest TRect, canvas uintptr, source TRect) {
	defSyscallN(dllimports.CANVAS_COPYRECT, obj, uintptr(unsafe.Pointer(&dest)), canvas, uintptr(unsafe.Pointer(&source)))
}

func Canvas_Draw1(obj uintptr, x, y int32, graphic uintptr) {
	defSyscallN(dllimports.CANVAS_DRAW1, obj, uintptr(x), uintptr(y), graphic)
}

func Canvas_DrawFocusRect(obj uintptr, aRect TRect) {
	defSyscallN(dllimports.CANVAS_DRAWFOCUSRECT, obj, uintptr(unsafe.Pointer(&aRect)))
}

func Canvas_FillRect(obj uintptr, aRect TRect) {
	defSyscallN(dllimports.CANVAS_FILLRECT, obj, uintptr(unsafe.Pointer(&aRect)))
}

func Canvas_FrameRect(obj uintptr, aRect TRect) {
	defSyscallN(dllimports.CANVAS_FRAMERECT, obj, uintptr(unsafe.Pointer(&aRect)))
}

func Canvas_TextRect1(obj uintptr, aRect TRect, x, y int32, text string) {
	defSyscallN(dllimports.CANVAS_TEXTRECT1, obj, uintptr(unsafe.Pointer(&aRect)), uintptr(x), uintptr(y), PascalStr(text))
}

func Canvas_TextRect2(obj uintptr, aRect *TRect, text string, textFormat TTextFormat) {
	defSyscallN(dllimports.CANVAS_TEXTRECT2, obj, uintptr(unsafe.Pointer(aRect)), PascalStr(text), uintptr(textFormat))
}

func Canvas_Polygon(obj uintptr, points []TPoint) {
	if len(points) == 0 {
		return
	}
	defSyscallN(dllimports.CANVAS_POLYGON, obj, uintptr(unsafe.Pointer(&points[0])), uintptr(len(points)))
}

func Canvas_Polyline(obj uintptr, points []TPoint) {
	if len(points) == 0 {
		return
	}
	defSyscallN(dllimports.CANVAS_POLYLINE, obj, uintptr(unsafe.Pointer(&points[0])), uintptr(len(points)))
}

func Canvas_PolyBezier(obj uintptr, points []TPoint) {
	if len(points) == 0 {
		return
	}
	defSyscallN(dllimports.CANVAS_POLYBEZIER, obj, uintptr(unsafe.Pointer(&points[0])), uintptr(len(points)))
}

// -------------------- TClipboard ---------------------------

func Clipboard_Instance() uintptr {
	return defSyscallN(dllimports.CLIPBOARD_INSTANCE)
}

func Clipboard_HasFormat(obj uintptr, aFormatID TClipboardFormat) bool {
	return GoBool(defSyscallN(dllimports.CLIPBOARD_HASFORMAT, obj, uintptr(aFormatID)))
}

func Clipboard_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if BufSize <= 0 {
		return 0
	}
	buff := make([]byte, BufSize)
	ret := defSyscallN(dllimports.CLIPBOARD_GETTEXTBUF, obj, uintptr(unsafe.Pointer(&buff[0])), uintptr(BufSize))
	if int(ret) < len(buff) {
		*Buffer = string(buff[:ret])
	}
	return int32(ret)
}

func Clipboard_GetAsText(obj uintptr) string {
	return GoStr(defSyscallN(dllimports.CLIPBOARD_GETASTEXT, obj))
}

func Clipboard_SetAsText(obj uintptr, value string) {
	defSyscallN(dllimports.CLIPBOARD_SETASTEXT, obj, PascalStr(value))
}

func Clipboard_GetAsHtml(obj uintptr, ExtractFragmentOnly bool) string {
	return GoStr(defSyscallN(dllimports.CLIPBOARD_GETASHTML, obj, PascalBool(ExtractFragmentOnly)))
}

func DSetClipboard(obj uintptr) uintptr {
	return defSyscallN(dllimports.DSETCLIPBOARD, obj)
}

func DRegisterClipboardFormat(aFormat string) TClipboardFormat {
	return TClipboardFormat(defSyscallN(dllimports.DREGISTERCLIPBOARDFORMAT, PascalStr(aFormat)))
}

func DPredefinedClipboardFormat(aFormat TPredefinedClipboardFormat) TClipboardFormat {
	return TClipboardFormat(defSyscallN(dllimports.DPREDEFINEDCLIPBOARDFORMAT, uintptr(aFormat)))
}

// -------------------- TImageList ---------------------------

func ImageList_Draw1(obj uintptr, canvas uintptr, x, y, index int32, enabled bool) {
	defSyscallN(dllimports.IMAGELIST_DRAW1, obj, canvas, uintptr(x), uintptr(y), uintptr(index), PascalBool(enabled))
}

func ImageList_Draw2(obj uintptr, canvas uintptr, x, y, index int32, drawingStyle TDrawingStyle, imageType TImageType, enabled bool) {
	defSyscallN(dllimports.IMAGELIST_DRAW2, obj, canvas, uintptr(x), uintptr(y), uintptr(index), uintptr(drawingStyle), uintptr(imageType), PascalBool(enabled))
}

func ImageList_DrawOverlay1(obj uintptr, canvas uintptr, x, y, imageIndex int32, overlay uint8, enabled bool) {
	defSyscallN(dllimports.IMAGELIST_DRAWOVERLAY1, obj, canvas, uintptr(x), uintptr(y), uintptr(imageIndex), uintptr(overlay), PascalBool(enabled))
}

func ImageList_DrawOverlay2(obj, canvas uintptr, x, y, imageIndex int32, overlay uint8, drawingStyle TDrawingStyle, imageType TImageType, enabled bool) {
	defSyscallN(dllimports.IMAGELIST_DRAWOVERLAY2, obj, canvas, uintptr(x), uintptr(y), uintptr(imageIndex), uintptr(overlay), uintptr(drawingStyle), uintptr(imageType), PascalBool(enabled))
}

func ImageList_GetIcon1(obj uintptr, index int32, image uintptr) {
	defSyscallN(dllimports.IMAGELIST_GETICON1, obj, uintptr(index), image)
}

func ImageList_GetIcon2(obj uintptr, index int32, image uintptr, drawingStyle TDrawingStyle, imageType TImageType) {
	defSyscallN(dllimports.IMAGELIST_GETICON2, obj, uintptr(index), image, uintptr(drawingStyle), uintptr(imageType))
}

// -------------------- TMemoryStream ---------------------------

func MemoryStream_Read(obj uintptr, count int32) (int32, []byte) {
	if count <= 0 {
		return 0, nil
	}
	bs := make([]byte, count)
	r := defSyscallN(dllimports.MEMORYSTREAM_READ, obj, uintptr(unsafe.Pointer(&bs[0])), uintptr(count))
	return int32(r), bs
}

func MemoryStream_Write(obj uintptr, buffer []byte) int32 {
	return int32(defSyscallN(dllimports.MEMORYSTREAM_WRITE, obj, uintptr(unsafe.Pointer(&buffer[0])), uintptr(len(buffer))))
}

// -------------------- TMouse ---------------------------

func Mouse_Instance() uintptr {
	return defSyscallN(dllimports.MOUSE_INSTANCE)
}

// -------------------- TPrinter ---------------------------

func Printer_Instance() uintptr {
	return defSyscallN(dllimports.PRINTER_INSTANCE)
}

func Printer_SetPrinter(obj uintptr, aName string) {
	defSyscallN(dllimports.PRINTER_SETPRINTER, obj, PascalStr(aName))
}

// -------------------- TScreen ---------------------------

func Screen_Instance() uintptr {
	return defSyscallN(dllimports.SCREEN_INSTANCE)
}

// -------------------- Procedures/Functions ---------------------------

func SetEventCallback(ptr uintptr) {
	defSyscallN(dllimports.SETEVENTCALLBACK, ptr)
}

func SetMessageCallback(ptr uintptr) {
	defSyscallN(dllimports.SETMESSAGECALLBACK, ptr)
}

func SetThreadSyncCallback(ptr uintptr) {
	defSyscallN(dllimports.SETTHREADSYNCCALLBACK, ptr)
}

func SetRequestCallCreateParamsCallback(ptr uintptr) {
	defSyscallN(dllimports.SETREQUESTCALLCREATEPARAMSCALLBACK, ptr)
}

func SetRemoveEventCallback(ptr uintptr) {
	defSyscallN(dllimports.SETREMOVEEVENTCALLBACK, ptr)
}

func DGetStringArrOf(p uintptr, index int) string {
	return GoStr(defSyscallN(dllimports.DGETSTRINGARROF, p, uintptr(index)))
}

func DStrLen(p uintptr) int {
	return int(defSyscallN(dllimports.DSTRLEN, p))
}

func DMove(src, dest uintptr, nLen int) {
	defSyscallN(dllimports.DMOVE, src, dest, uintptr(nLen))
}

func DShowMessage(s string) {
	defSyscallN(dllimports.DSHOWMESSAGE, PascalStr(s))
}

func DMessageDlg(Msg string, DlgType TMsgDlgType, Buttons TMsgDlgButtons, HelpCtx int32) int32 {
	return int32(defSyscallN(dllimports.DMESSAGEDLG, PascalStr(Msg), uintptr(DlgType), uintptr(Buttons), uintptr(HelpCtx)))
}

func DTextToShortCut(val string) TShortCut {
	return TShortCut(defSyscallN(dllimports.DTEXTTOSHORTCUT, PascalStr(val)))
}

func DShortCutToText(val TShortCut) string {
	return GoStr(defSyscallN(dllimports.DSHORTCUTTOTEXT, uintptr(val)))
}

func DSysOpen(filename string) {
	defSyscallN(dllimports.DSYSOPEN, PascalStr(filename))
}

func DExtractFilePath(filename string) string {
	return GoStr(defSyscallN(dllimports.DEXTRACTFILEPATH, PascalStr(filename)))
}

func DFileExists(filename string) bool {
	return GoBool(defSyscallN(dllimports.DFILEEXISTS, PascalStr(filename)))
}

func DSelectDirectory1(options TSelectDirOpts) (bool, string) {
	var ptr uintptr
	v := GoBool(defSyscallN(dllimports.DSELECTDIRECTORY1, uintptr(unsafe.Pointer(&ptr)), uintptr(options), 0))
	if v {
		return true, GoStr(ptr)
	}
	return false, ""
}

func DSelectDirectory2(caption, root string, showHidden bool) (bool, string) {
	var ptr uintptr
	v := GoBool(defSyscallN(dllimports.DSELECTDIRECTORY2, PascalStr(caption), PascalStr(root), PascalBool(showHidden), uintptr(unsafe.Pointer(&ptr))))
	if v {
		return true, GoStr(ptr)
	}
	return false, ""
}

func DSynchronize(fn func(), useMsg uintptr) {
	threadSyncRef.Lock()
	defer threadSyncRef.Unlock()
	threadSyncProc = fn
	defSyscallN(dllimports.DSYNCHRONIZE, useMsg)
	threadSyncProc = nil
}

func DInputBox(aCaption, aPrompt, aDefault string) string {
	return GoStr(defSyscallN(dllimports.DINPUTBOX, PascalStr(aCaption), PascalStr(aPrompt), PascalStr(aDefault)))
}

func DInputQuery(aCaption, aPrompt string, value *string) bool {
	if value == nil {
		return false
	}
	var strPtr uintptr
	r := defSyscallN(dllimports.DINPUTQUERY, PascalStr(aCaption), PascalStr(aPrompt), PascalStr(*value), uintptr(unsafe.Pointer(&strPtr)))
	if strPtr != 0 {
		*value = GoStr(strPtr)
	}
	return GoBool(r)
}

func DPasswordBox(aCaption, aPrompt string) string {
	return GoStr(defSyscallN(dllimports.DPASSWORDBOX, PascalStr(aCaption), PascalStr(aPrompt)))
}

func DInputCombo(aCaption, aPrompt string, aList uintptr) int32 {
	return int32(defSyscallN(dllimports.DINPUTCOMBO, PascalStr(aCaption), PascalStr(aPrompt), aList))
}

func DInputComboEx(aCaption, aPrompt string, aList uintptr, allowCustomText bool) string {
	return GoStr(defSyscallN(dllimports.DINPUTCOMBOEX, PascalStr(aCaption), PascalStr(aPrompt), aList, PascalBool(allowCustomText)))
}

func DSysLocale(aInfo *TSysLocale) {
	defSyscallN(dllimports.DSYSLOCALE, uintptr(unsafe.Pointer(aInfo)))
}

func DCreateURLShortCut(aDestPath, aShortCutName, aURL string) {
	defSyscallN(dllimports.DCREATEURLSHORTCUT, PascalStr(aDestPath), PascalStr(aShortCutName), PascalStr(aURL))
}

func DCreateShortCut(aDestPath, aShortCutName, aSrcFileName, aIconFileName, aDescription, aCmdArgs string) bool {
	return GoBool(defSyscallN(dllimports.DCREATESHORTCUT, PascalStr(aDestPath), PascalStr(aShortCutName), PascalStr(aSrcFileName),
		PascalStr(aIconFileName), PascalStr(aDescription), PascalStr(aCmdArgs)))
}

func DSetPropertyValue(instance uintptr, propName, value string) {
	defSyscallN(dllimports.DSETPROPERTYVALUE, instance, PascalStr(propName), PascalStr(value))
}

func DSetPropertySecValue(instance uintptr, propName, secPropName, value string) {
	defSyscallN(dllimports.DSETPROPERTYSECVALUE, instance, PascalStr(propName), PascalStr(secPropName), PascalStr(value))
}

func DGUIDToString(guid TGUID) string {
	return GoStr(defSyscallN(dllimports.DGUIDTOSTRING, uintptr(unsafe.Pointer(&guid))))
}

func DStringToGUID(str string) (guid TGUID) {
	defSyscallN(dllimports.DSTRINGTOGUID, PascalStr(str), uintptr(unsafe.Pointer(&guid)))
	return
}

func DCreateGUID() (guid TGUID) {
	defSyscallN(dllimports.DCREATEGUID, uintptr(unsafe.Pointer(&guid)))
	return
}

func DGetLibResourceCount() int32 {
	return int32(defSyscallN(dllimports.DGETLIBRESOURCECOUNT))
}

func DGetLibResourceItem(aIndex int32) (ret TLibResource) {
	item := struct {
		Name     uintptr
		ValuePtr uintptr
	}{}
	defSyscallN(dllimports.DGETLIBRESOURCEITEM, uintptr(aIndex), uintptr(unsafe.Pointer(&item)))
	ret.Name = GoStr(item.Name)
	ret.Ptr = item.ValuePtr
	return
}

func DModifyLibResource(aPtr uintptr, aValue string) {
	defSyscallN(dllimports.DMODIFYLIBRESOURCE, aPtr, PascalStr(aValue))
}

func DLibStringEncoding() TStringEncoding {
	return TStringEncoding(defSyscallN(dllimports.DLIBSTRINGENCODING))
}

func DLibVersion() uint32 {
	return uint32(defSyscallN(dllimports.DLIBVERSION))
}

func DLibAbout() string {
	return GoStr(defSyscallN(dllimports.DLIBABOUT))
}

func DMainThreadId() uintptr {
	return defSyscallN(dllimports.DMAINTHREADID)
}

func DCurrentThreadId() uintptr {
	return defSyscallN(dllimports.DCURRENTTHREADID)
}

func DInitGoDll(aMainThreadId uintptr) {
	defSyscallN(dllimports.DINITGODLL, aMainThreadId)
}

func DFindControl(handle HWND) uintptr {
	return defSyscallN(dllimports.DFINDCONTROL, handle)
}

func DFindLCLControl(screenPos TPoint) uintptr {
	return defSyscallN(dllimports.DFINDLCLCONTROL, uintptr(unsafe.Pointer(&screenPos)))
}

func DFindOwnerControl(handle HWND) uintptr {
	return defSyscallN(dllimports.DFINDOWNERCONTROL, handle)
}

func DFindControlAtPosition(position TPoint, allowDisabled bool) uintptr {
	return defSyscallN(dllimports.DFINDCONTROLATPOSITION, uintptr(unsafe.Pointer(&position)), PascalBool(allowDisabled))
}

func DFindLCLWindow(screenPos TPoint, allowDisabled bool) uintptr {
	return defSyscallN(dllimports.DFINDLCLWINDOW, uintptr(unsafe.Pointer(&screenPos)), PascalBool(allowDisabled))
}

func DFindDragTarget(position TPoint, allowDisabled bool) uintptr {
	return defSyscallN(dllimports.DFINDDRAGTARGET, uintptr(unsafe.Pointer(&position)), PascalBool(allowDisabled))
}
