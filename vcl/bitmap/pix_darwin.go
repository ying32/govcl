//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------
// +build darwin

package bitmap

//var (
//	// A R G B
//	// Big Endian
//	pixIndex = [4]int{3, 0, 1, 2}
//)

// Big Endian
type bgra struct {
	A, R, G, B uint8
}
