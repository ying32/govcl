// +build !windows
// +build amd64

package types

// 消息值参见 types/messages包
type TMessage struct {
	Msg        Cardinal
	_UnusedMsg Cardinal
	WParam     WAPRAM
	LParam     LPARAM
	Result     LRESULT
}
