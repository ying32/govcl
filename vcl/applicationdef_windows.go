package vcl

import (
	"gitee.com/ying32/govcl/vcl/win"
)

// SetIconResId 从资源中置图标的id
func (a *TApplication) SetIconResId(id int) {
	a.Icon().SetHandle(win.LoadIcon(win.GetSelfModuleHandle(), id))
}
