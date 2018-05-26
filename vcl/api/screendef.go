package api

func Screen_Instance() uintptr {
	ret, _, _ := screen_Instance.Call()
	return ret
}
