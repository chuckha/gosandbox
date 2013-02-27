package main

import (
	"fmt"
	"flag"
	"os"
	"regexp"
	"sync"
)


var wg sync.WaitGroup

func main() {
	flag.Parse()
	pattern := flag.Args()[0]
	re, err := regexp.Compile(pattern)
	if err != nil { panic(err) }

	results := make(chan string)

	for _, filename := range flag.Args()[1:] {
		ChunkFile(filename, re, results)
	}

	for v := range results {
		fmt.Println(v)
	}

}


func ChunkFile(filename string, re *regexp.Regexp, founds chan string) {
	file, err := os.Open(filename)
	if err != nil { panic(err) }
	defer file.Close()

	chunks := make([]byte, 1024)
	for {
		numBytes, err := file.Read(chunks)
		if numBytes == 0 { break }
		if err != nil { panic(err) }
		wg.Add(1)
		go func(chunks []byte) {
			for _, v := range re.FindAll(chunks, -1) {
				founds <- string(v)
			}
			wg.Done()
		}(chunks)
	}
}
