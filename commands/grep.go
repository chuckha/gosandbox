package main

import (
	"fmt"
	"flag"
	"io"
	"os"
	"regexp"
)

func main() {
	flag.Parse()
	pattern := flag.Args()[0]
	re, err := regexp.Compile(pattern)
	if err != nil { panic(err) }

	for _, filename := range flag.Args()[1:] {
		matches := Search(filename, re)
		for _, match := range matches {
			fmt.Printf("%v: %s\n", filename, match)
		}
	}
}

func Search(filename string, pattern *regexp.Regexp) []string {
	file, err := os.Open(filename)
	if err != nil { panic(err) }
	defer file.Close()

	finds := []string{}

	buf := make([]byte, 1024)

	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF { panic(err) }
		if n == 0 { break }
		for _, v := range pattern.FindAll(buf, -1) {
			finds = append(finds, string(v))
		}
	}

	return finds
}
