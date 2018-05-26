package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
    . "gitee.com/ying32/govcl/vcl/types"
)


func (i *TImageList) Draw (canvas IObject, x, y, index int32, enabled bool) {
    ImageList_Draw1(i.instance, CheckPtr(canvas), x, y, index, enabled)
}

func (i *TImageList) Draw2 (canvas IObject, x, y, index int32,  drawingStyle TDrawingStyle, imageType TImageType, enabled bool) {
    ImageList_Draw2(i.instance, CheckPtr(canvas), x, y, index, drawingStyle, imageType, enabled)
}

func (i *TImageList) DrawOverlay (canvas IObject, x, y, imageIndex int32, overlay uint8, enabled bool) {
    ImageList_DrawOverlay1(i.instance, CheckPtr(canvas), x, y, imageIndex, overlay, enabled)
}

func (i *TImageList) DrawOverlay2 (canvas IObject, x, y, imageIndex int32, overlay uint8, drawingStyle TDrawingStyle, imageType TImageType, enabled bool) {
    ImageList_DrawOverlay2(i.instance, CheckPtr(canvas), x, y, imageIndex, overlay, drawingStyle, imageType, enabled)
}

func (i *TImageList) GetIcon (index int32, image IObject) {
    ImageList_GetIcon1(i.instance, index, CheckPtr(image))
}

func (i *TImageList) GetIcon2 (index int32, image IObject, drawingStyle TDrawingStyle, imageType TImageType) {
    ImageList_GetIcon2(i.instance, index, CheckPtr(image), drawingStyle, imageType)
}