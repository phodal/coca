package main

import "fmt"

// #cgo CFLAGS: -DPNG_DEBUG=1
// #include "hello.h"
import "C"

func main() {
	fmt.Printf("Invoking c library...\n")
	C.output()
	fmt.Printf("Done\n")
}