//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

//+build finalizerOn

package vcl

import "runtime"

func setFinalizer(obj interface{}, finalizer interface{}) {
	runtime.SetFinalizer(obj, finalizer)
}
