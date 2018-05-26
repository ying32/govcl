package api

import (
	"unsafe"

	. "gitee.com/ying32/govcl/vcl/types"
)

func Canvas_BrushCopy(obj uintptr, dest TRect, bitmap uintptr, source TRect, color TColor) {
	canvas_BrushCopy.Call(obj, uintptr(unsafe.Pointer(&dest)), bitmap, uintptr(unsafe.Pointer(&source)), uintptr(color))
}

func Canvas_CopyRect(obj uintptr, dest TRect, canvas uintptr, source TRect) {
	canvas_CopyRect.Call(obj, uintptr(unsafe.Pointer(&dest)), canvas, uintptr(unsafe.Pointer(&source)))
}

func Canvas_Draw1(obj uintptr, x, y int32, graphic uintptr) {
	canvas_Draw1.Call(obj, uintptr(x), uintptr(y), graphic)
}

func Canvas_Draw2(obj uintptr, x, y int32, graphic uintptr, opacity uint8) {
	canvas_Draw2.Call(obj, uintptr(x), uintptr(y), graphic, uintptr(opacity))
}

func Canvas_DrawFocusRect(obj uintptr, aRect TRect) {
	canvas_DrawFocusRect.Call(obj, uintptr(unsafe.Pointer(&aRect)))
}

func Canvas_FillRect(obj uintptr, aRect TRect) {
	canvas_FillRect.Call(obj, uintptr(unsafe.Pointer(&aRect)))
}

func Canvas_FrameRect(obj uintptr, aRect TRect) {
	canvas_FrameRect.Call(obj, uintptr(unsafe.Pointer(&aRect)))
}

func Canvas_StretchDraw(obj uintptr, aRect TRect, graphic uintptr) {
	canvas_StretchDraw.Call(obj, uintptr(unsafe.Pointer(&aRect)), graphic)
}

func Canvas_TextRect2(obj uintptr, aRect *TRect, text *string, textFormat TTextFormat) {
	var pstr uintptr
	r, _, _ := canvas_TextRect2.Call(obj, uintptr(unsafe.Pointer(aRect)), GoStrToDStr(*text), uintptr(unsafe.Pointer(&pstr)), uintptr(textFormat))
	if r == 1 {
		*text = DStrToGoStr(pstr)
	}
}

func Canvas_TextRect1(obj uintptr, aRect TRect, x, y int32, text string) {
	canvas_TextRect1.Call(obj, uintptr(unsafe.Pointer(&aRect)), uintptr(x), uintptr(y), GoStrToDStr(text))
}
