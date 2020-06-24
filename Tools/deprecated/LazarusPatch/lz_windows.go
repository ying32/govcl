package main

import (
	"io/ioutil"
	"regexp"

	"github.com/ying32/govcl/vcl/win"
)

func GetLazarusPath() string {
	var ppidl *win.TItemIDList

	win.SHGetSpecialFolderLocation(0, win.CSIDL_LOCAL_APPDATA, &ppidl)
	res, szPath := win.SHGetPathFromIDList(ppidl)
	if res {
		envFile := szPath + "\\lazarus\\environmentoptions.xml"
		if checkFileExists(envFile) {
			bs, err := ioutil.ReadFile(envFile)
			if err == nil {
				reg := regexp.MustCompile(`\<LazarusDirectory Value\=\"(.+?)"\>`)
				matchs := reg.FindSubmatch(bs)
				if len(matchs) >= 2 {
					return string(matchs[1])
				}
			}
		}
	}
	return ""
}
