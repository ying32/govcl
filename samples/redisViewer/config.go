//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/ying32/govcl/vcl/rtl"
)

type TConnConfig struct {
	Name     string
	Host     string
	Port     int
	Password string
}

// 全局配置
var (
	appConfFileName = rtl.ExtractFilePath(os.Args[0]) + "app.conf"
	config          = make(map[string]TConnConfig, 0)
)

func saveConfig() {
	bs, err := json.Marshal(config)
	if err == nil {
		ioutil.WriteFile(appConfFileName, bs, 0664)
	}
}

func loadConfig() {
	config = make(map[string]TConnConfig, 0)
	if rtl.FileExists(appConfFileName) {
		bs, err := ioutil.ReadFile(appConfFileName)
		if err == nil {
			json.Unmarshal(bs, &config)
		}
	}
}
