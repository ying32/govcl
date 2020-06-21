//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// +build windows linux
// +build !tempdll
// +build !memorydll

// 指令为：target == windows || target == linux && !tempdll && !memorydll

package api

func checkAndReleaseDLL() (bool, string) {
	return false, ""
}
