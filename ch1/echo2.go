package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var s, sep string
	for _, v := range os.Args[1:] {
		s += sep + v
		sep = " "
	}
	fmt.Println(s)
}
