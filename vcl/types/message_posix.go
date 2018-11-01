// +build linux,amd64 darwin,amd64

package types

type TMessage struct {
	Msg        Cardinal
	_UnusedMsg Cardinal
	WParam     WAPRAM
	LParam     LPARAM
	Result     LRESULT
}
