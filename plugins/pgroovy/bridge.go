package main

// #cgo CFLAGS: -DPNG_DEBUG=1
// #include "wrap.h"
import "C"

func main() {
	C.output()
}