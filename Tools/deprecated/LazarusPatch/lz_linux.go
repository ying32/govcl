package main

import (
	"io/ioutil"
	"regexp"
)

func GetLazarusPath() string {
	envFile := "/etc/lazarus/environmentoptions.xml"
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
	return ""
}
