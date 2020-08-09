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
	"github.com/ying32/govcl/vcl/types"
)

func SetClipboard(newClipboard IObject) *TClipboard {
	return AsClipboard(DSetClipboard(CheckPtr(newClipboard)))
}

func RegisterClipboardFormat(aFormat string) types.TClipboardFormat {
	return DRegisterClipboardFormat(aFormat)
}

func (c *TClipboard) HasFormat(aFormatID types.TClipboardFormat) bool {
	return Clipboard_HasFormat(c.instance, aFormatID)
}
