package libvlc

import "errors"

func ErrMsg() error {
	s := libvlc_errmsg()
	if s == "" {
		return nil
	}
	return errors.New(s)
}

func Version() {
	//libvlc_get_Version()
}
