package msrdp

import "github.com/ying32/govcl/vcl"

// 要注册的事件回调
func eventCallback(fn interface{}, getVal func(idx int) uintptr) bool {
	switch fn.(type) {
	case TMsRdpClient9NotSafeForScriptingOnDisconnected:
		fn.(TMsRdpClient9NotSafeForScriptingOnDisconnected)(
			vcl.ObjectFromInst(getVal(0)),
			int32(getVal(1)))
	default:

	}
	return false
}
