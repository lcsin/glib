package iutil

import _ "unsafe"

//go:linkname Bool2Int runtime.bool2int
func Bool2Int(b bool) int

//go:linkname NanoTime runtime.nanotime
func NanoTime() int64
