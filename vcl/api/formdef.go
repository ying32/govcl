package api

func Form_EnabledMaximize(obj uintptr, val bool) {
	form_EnabledMaximize.Call(obj, GoBoolToDBool(val))
}

func Form_EnabledMinimize(obj uintptr, val bool) {
	form_EnabledMinimize.Call(obj, GoBoolToDBool(val))
}

func Form_EnabledSystemMenu(obj uintptr, val bool) {
	form_EnabledSystemMenu.Call(obj, GoBoolToDBool(val))
}

func Form_SetAllowDropFiles(obj uintptr, val bool) {
	form_SetAllowDropFiles.Call(obj, GoBoolToDBool(val))
}

func Form_GetAllowDropFiles(obj uintptr) bool {
	r, _, _ := form_GetAllowDropFiles.Call(obj)
	return DBoolToGoBool(r)
}

func Form_SetOnDropFiles(obj uintptr, fn interface{}) {
	form_SetOnDropFiles.Call(obj, addEventToMap(fn))
}

func Form_SetOnDestroy(obj uintptr, fn interface{}) {
	form_SetOnDestroy.Call(obj, addEventToMap(fn))
}

func Form_SetOnConstrainedResize(obj uintptr, fn interface{}) {
	form_SetOnConstrainedResize.Call(obj, addEventToMap(fn))
}

func Form_SetOnDeactivate(obj uintptr, fn interface{}) {
	form_SetOnDeactivate.Call(obj, addEventToMap(fn))
}

func Form_SetOnActivate(obj uintptr, fn interface{}) {
	form_SetOnActivate.Call(obj, addEventToMap(fn))
}

// 下面两个函数放在Application下面吧，直接调用实例类

func SetGlobalFormScaled(val bool) {
	setGlobalFormScaled.Call(GoBoolToDBool(val))
}

func Form_ScaleForPPI(obj uintptr, newPPI int32) {
	form_ScaleForPPI.Call(obj, uintptr(newPPI))
}

func Form_ScaleControlsForDpi(obj uintptr, newPPI int32) {
	form_ScaleControlsForDpi.Call(obj, uintptr(newPPI))
}
