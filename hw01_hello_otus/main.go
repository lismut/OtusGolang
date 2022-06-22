package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	s := "Hello, OTUS!"
	fmt.Printf("%s\n", stringutil.Reverse(s))
}
