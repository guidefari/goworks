package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
	"golang.org/x/example/stringutil"
)

func main() {
    fmt.Println(reverse.String("Hello"))
	fmt.Println(stringutil.ToUpper("Hello"))
}