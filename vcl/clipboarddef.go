package vcl

import (
	. "github.com/ying32/govcl/vcl/api"
)

func SetClipboard(newClipboard IObject) *TClipboard {
	return ClipboardFromInst(Clipboard_SetClipboard(CheckPtr(newClipboard)))
}
