package api

func Mouse_Instance() uintptr {
	ret, _, _ := mouse_Instance.Call()
	return ret
}
