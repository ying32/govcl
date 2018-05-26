package api

func Application_Instance() uintptr {
	ret, _, _ := application_Instance.Call()
	return ret
}

func Application_CreateForm(app uintptr) uintptr {
	ret, _, _ := application_CreateForm.Call(app)
	return ret
}

func Application_Run(app uintptr) {
	application_Run.Call(app)
}
