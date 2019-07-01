package msrdp

type IMsRdpClientAdvancedSettings8 struct {
	instance uintptr
}

func IMsRdpClientAdvancedSettings8FromInst(inst uintptr) *IMsRdpClientAdvancedSettings8 {
	m := new(IMsRdpClientAdvancedSettings8)
	m.instance = inst
	return m
}

func (m *IMsRdpClientAdvancedSettings8) Instance() uintptr {
	return m.instance
}
