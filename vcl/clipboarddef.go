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
	return Clipboard_HasFormat(c.instance, aFormatID)
}

func (c *TClipboard) GetAsHtml(ExtractFragmentOnly bool) string {
	return Clipboard_GetAsHtml(c.instance, ExtractFragmentOnly)
}

func (c *TClipboard) GetTextBuf(Buffer *string, BufSize int32) int32 {
	return Clipboard_GetTextBuf(c.instance, Buffer, BufSize)
}

func (c *TClipboard) AsText() string {
	mem := NewMemoryStream()
	defer mem.Free()
	if c.GetFormat(PredefinedClipboardFormat(types.PcfText), mem) {
		size := mem.Size()
		mem.SetPosition(0)
		if size > 0 {
			n, buff := mem.Read(int32(size))
			if n > 0 && buff[len(buff)-1] == 0 {
				return string(buff[:len(buff)-1])
			} else {
				return string(buff)
			}
		}
	}
	return ""
	//return Clipboard_GetAsText(c.instance)
}

func (c *TClipboard) SetAsText(value string) {
	Clipboard_SetAsText(c.instance, value)
}
