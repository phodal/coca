package main

import (
	"fmt"
	"os"
	"strings"

	"./imp"
)

func main() {
	var a [3]int // array of 3 integers

	// Print the elements only.
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	imp.Count(1)
	fmt.Println(strings.Join(os.Args[1:], " "))
}
