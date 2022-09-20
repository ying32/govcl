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

func PredefinedClipboardFormat(aFormat types.TPredefinedClipboardFormat) types.TClipboardFormat {
	return DPredefinedClipboardFormat(aFormat)
}

func (c *TClipboard) HasFormat(aFormatID types.TClipboardFormat) bool {
	return Clipboard_HasFormat(c._instance(), aFormatID)
}

func (c *TClipboard) GetAsHtml(ExtractFragmentOnly bool) string {
	return Clipboard_GetAsHtml(c._instance(), ExtractFragmentOnly)
}

func (c *TClipboard) GetTextBuf(Buffer *string, BufSize int32) int32 {
	return Clipboard_GetTextBuf(c._instance(), Buffer, BufSize)
}

func (c *TClipboard) AsText() string {
	mem := NewMemoryStream()
	defer mem.Free()
	if c.GetFormat(PredefinedClipboardFormat(types.PcfText), mem) {
		return getStreamText(mem)
	}
	return ""
}

func (c *TClipboard) SetAsText(value string) {
	Clipboard_SetAsText(c._instance(), value)
}
