// +build windows

package main

import (
	"os"

	"github.com/ying32/govcl/vcl"
	_ "github.com/ying32/govcl/vcl/exts/winappres"
	"github.com/ying32/govcl/vcl/rtl"
	"github.com/ying32/govcl/vcl/win"
)

func main() {
	rtl.CreateURLShortCut(win.GetDesktopPath(), "govcl", "https://github.com/ying32/govcl")

	rtl.CreateShortCut(win.GetDesktopPath(), "shortcut", os.Args[0], "", "描述", "-b -c")

	vcl.ShowMessage("Hello!")

}
