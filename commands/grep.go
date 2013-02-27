package main

import (
	"fmt"
	"flag"
	"os"
	"regexp"
)

func main() {
	flag.Parse()
	pattern := flag.Args()[0]
	re, err := regexp.Compile(pattern)
	if err != nil { panic(err) }
	fmt.Println(re)

	bytes := make(chan []byte)

	filename := flag.Args()[1]
	go ChunkFile(filename, bytes)

	for v := range bytes {
		fmt.Println(len(v))
	}

}

func ChunkFile(filename string, bytes chan []byte) {
	file, err := os.Open(filename)
	if err != nil { panic(err) }
	defer file.Close()

	chunks := make([]byte, 1024)
	for {
		numBytes, err := file.Read(chunks)
		if numBytes == 0 { break }
		if err != nil { panic(err) }
		bytes <- chunks[:numBytes]
	}
	close(bytes)
}
/*
, re *regexp.Regexp
		go func(chunks []byte) {
			for _, v := range re.FindAll(chunks, -1) {
				founds <- string(v)
			}
			
		}(chunks)
		*/
