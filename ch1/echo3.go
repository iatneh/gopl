package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(strings.Join(os.Args, " "))
	for i, v := range os.Args {
		fmt.Printf("index=%d,value=%s\n", i, v)
	}
}
