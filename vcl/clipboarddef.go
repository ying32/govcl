package vcl

import (
	. "gitee.com/ying32/govcl/vcl/api"
)

func SetClipboard(newClipboard IObject) *TClipboard {
	return ClipboardFromInst(Clipboard_SetClipboard(CheckPtr(newClipboard)))
}
