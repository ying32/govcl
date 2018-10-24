package api

// Printer
func Printer_Instance() uintptr {
	r, _, _ := printer_Instance.Call()
	return r
}
