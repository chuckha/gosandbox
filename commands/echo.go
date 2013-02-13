package main

import (
	"fmt"
	"strings"
	"flag"
)

var newline = flag.Bool("n", false, "No newline character will be printed if this is set")

func main() {
	flag.Parse()
	fmt.Printf(strings.Join(flag.Args(), " "))
	if !*newline {
		fmt.Printf("\n")
	}
}
