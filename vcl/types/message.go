// +build windows arm linux,386 darwin,386

package types

// 消息值参见 types/messages包
type TMessage struct {
	Msg    Cardinal
	WParam WAPRAM
	LParam LPARAM
	Result LRESULT
}
