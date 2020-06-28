package main

import (
	_ "github.com/ying32/govcl/pkgs/winappres"
	"github.com/ying32/govcl/vcl"
)

func main() {
	vcl.RunApp(&MainForm,
		&AppSettings,
		&OSSSaveSettings,
		&NewAliyunAudioEngine,
		&NewBaiduTranslateEngine,
		&NewTencentTranslateEngine)
}
