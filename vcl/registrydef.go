package vcl

import (
	"gitee.com/ying32/govcl/vcl/win"
)

// NewRegistryAllAccess 所有访问权限
func NewRegistryAllAccess() *TRegistry {
	return NewRegistry(win.KEY_ALL_ACCESS)
}
