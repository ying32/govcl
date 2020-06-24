package main

import (
	"io/ioutil"
	"regexp"

	"github.com/ying32/govcl/vcl/win"
)

func GetLazarusDir() (lazarusDir string, fpcExe string) {
	var ppidl *win.TItemIDList

	win.SHGetSpecialFolderLocation(0, win.CSIDL_LOCAL_APPDATA, &ppidl)
	res, szPath := win.SHGetPathFromIDList(ppidl)
	if res {
		lazarusDir, fpcExe = readLazarusEnvFile(szPath)
	}
	return
}

func GetBsdDir() (bsdDir string, userDir string) {
	var ppidl *win.TItemIDList

	win.SHGetSpecialFolderLocation(0, win.CSIDL_APPDATA, &ppidl)
	res, szPath := win.SHGetPathFromIDList(ppidl)
	if res {
		// 10.2.1 = 19.0
		envFile := szPath + "\\Embarcadero\\BDS\\19.0\\environment.proj"
		if checkFileExists(envFile) {
			bs, err := ioutil.ReadFile(envFile)
			if err == nil {
				reg := regexp.MustCompile(`\<BDS Condition=\"\'\$\(BDS\)\'\=\=\'\'\"\>(.+?)\<\/BDS\>`)
				matchs := reg.FindSubmatch(bs)
				if len(matchs) >= 2 {
					bsdDir = string(matchs[1])
				}
				reg = regexp.MustCompile(`\<BDSCOMMONDIR Condition=\"\'\$\(BDSCOMMONDIR\)\'\=\=\'\'\"\>(.+?)\<\/BDSCOMMONDIR\>`)
				matchs = reg.FindSubmatch(bs)
				if len(matchs) >= 2 {
					userDir = string(matchs[1])
				}
			}
		}
	}
	return
}
