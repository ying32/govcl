// +build linux darwin

package vcl

import (
	"fmt"
)

func showError(err interface{}) {
	fmt.Println(err)
}
