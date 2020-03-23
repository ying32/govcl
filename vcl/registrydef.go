//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package vcl

import (
	"github.com/ying32/govcl/vcl/win"
)

// NewRegistryAllAccess 所有访问权限
func NewRegistryAllAccess() *TRegistry {
	return NewRegistry(win.KEY_ALL_ACCESS)
}
