// +build windows,cgo

package miniblink

import "github.com/ying32/govcl/vcl/rtl"

var (
	isLcl bool
)

func init() {
	isLcl = rtl.LcLLoaded()
	//Initialize()

}
