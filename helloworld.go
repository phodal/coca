package main

import (
	"fmt"
	"os"
	"strings"

	"./imp"
)

func main() {
	imp.Count(1)
	fmt.Println(strings.Join(os.Args[1:], " "))
}
