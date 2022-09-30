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
