package api

import (
	. "gitee.com/ying32/govcl/vcl/types"
)
 
func ImageList_Draw1(obj uintptr, canvas uintptr, x, y, index int32, enabled bool) {
	imageList_Draw1.Call(obj, canvas, uintptr(x), uintptr(y), uintptr(index), GoBoolToDBool(enabled))
}

func ImageList_Draw2(obj uintptr, canvas uintptr, x, y, index int32, drawingStyle TDrawingStyle, 
	imageType TImageType, enabled bool) {
    imageList_Draw2.Call(obj, canvas, uintptr(x), uintptr(y), uintptr(index), uintptr(drawingStyle), uintptr(imageType), GoBoolToDBool(enabled))
}

func ImageList_DrawOverlay1(obj uintptr, canvas uintptr, x, y, imageIndex int32, overlay uint8, enabled bool) {
	imageList_DrawOverlay1.Call(obj, canvas, uintptr(x), uintptr(y), uintptr(imageIndex), uintptr(overlay), GoBoolToDBool(enabled))
}


func ImageList_DrawOverlay2(obj, canvas uintptr, x, y, imageIndex int32, overlay uint8,
	 drawingStyle TDrawingStyle, imageType TImageType, enabled bool) {
	imageList_DrawOverlay2.Call(obj, canvas, uintptr(x), uintptr(y), uintptr(imageIndex), uintptr(overlay), uintptr(drawingStyle), uintptr(imageType), GoBoolToDBool(enabled))
}

func ImageList_GetIcon1(obj uintptr, index int32, image uintptr) {
	imageList_GetIcon1.Call(obj, uintptr(index), image)
}

func ImageList_GetIcon2(obj uintptr, index int32, image uintptr, drawingStyle TDrawingStyle, imageType TImageType) {
	imageList_GetIcon2.Call(obj, uintptr(index), image, uintptr(drawingStyle), uintptr(imageType))
}