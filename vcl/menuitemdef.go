
//----------------------------------------
// 
// Copyright © ying32. All Rights Reserved.
// 
// Licensed under Apache License 2.0
//
//----------------------------------------


package vcl

import (
	. "github.com/ying32/govcl/vcl/api"
)

// SetShortCutFromString 设置快捷键字符
func (m *TMenuItem) SetShortCutFromString(s string) {
	MenuItem_SetShortCut(m.instance, DTextToShortCut(s))
}

// ShortCutFromString 获取快捷键字符
func (m *TMenuItem) ShortCutFromString() string {
	return DShortCutToText(MenuItem_GetShortCut(m.instance))
}
