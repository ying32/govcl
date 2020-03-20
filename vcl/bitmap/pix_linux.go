//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package bitmap

//var (
//	// R G B A
//	// Little Endian
//	pixIndex = [4]int{0, 1, 2, 3}
//)

// Little Endian
type rgba struct {
	R, G, B, A uint8
}

type rgb struct {
	R, G, B uint8
}
