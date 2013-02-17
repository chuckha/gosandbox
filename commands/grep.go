package main

import (
	"bufio"
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
		ChunkFile(filename, re)
	}
}


func ChunkFile(filename string, re *regexp.Regexp) {
	file, err := os.Open(filename)
	if err != nil { panic(err) }
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, e := reader.ReadString('\n')
		if e == io.EOF { break }
		if e != nil { panic(e) }
		for _, v := range re.FindAllString(line, -1) {
			fmt.Println(v)
		}
	}
}
