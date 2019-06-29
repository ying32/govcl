package msrdp

type IMsRdpClient9 struct {
	instance uintptr
}

func IMsRdpClient9FromInst(inst uintptr) *IMsRdpClient9 {
	m := new(IMsRdpClient9)
	m.instance = inst
	return m
}

func (m *IMsRdpClient9) Instance() uintptr {
	return m.instance
}
