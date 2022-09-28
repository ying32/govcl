//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

//go:build !hideversion
// +build !hideversion

package vcl

import "fmt"

func printVersion(version string) {
	fmt.Println("Library Version:", version)
}
