//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package bitmap

//var (
//	// A R G B
//	// Big Endian
//	pixIndex = [4]int{3, 0, 1, 2}
//)

// Big Endian
type rgba struct {
	A, R, G, B uint8
}

type rgb struct {
	R, G, B uint8
}
