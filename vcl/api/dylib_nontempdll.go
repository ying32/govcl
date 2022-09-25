//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

//go:build !tempdll
// +build !tempdll

// 指令为：!tempdll

package api

func checkAndReleaseDLL() (bool, string) {
	return false, ""
}
