package main

import (
	"fmt"
	"os"
)

func main() {
	var sep, s string
	for i := 1; i < len(os.Args); i++ {
		s += os.Args[i] + sep
		sep = " "
	}
	fmt.Println(s)
}
