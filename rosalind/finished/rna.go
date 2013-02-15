package main

import (
	"fmt"
	"os"
	"io"
)

func main() {
	filename := "RNA.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	dna := make(chan byte)
	go FnBytes(filename, dna, Replace)

	var bytes string
	for b := range dna {
		bytes += string(b)
	}
	fmt.Println(bytes)
}

func Replace(bs []byte, dna chan byte) {
	for _, b := range bs {
		switch b {
		case 'A': dna <- 'A'
		case 'G': dna <- 'G'
		case 'C': dna <- 'C'
		case 'T': dna <- 'U'
		}
	}
}

func FnBytes(filename string, dna chan byte, fn func ([]byte, chan byte)) {
	file, err := os.Open(filename)
	if err != nil { panic(err) }
	defer file.Close()

	// Buffer was not being filled

	for {
		buf := make([]byte, 10)
		n, err := file.Read(buf)
		if err != nil && err != io.EOF { panic(err) }
		if n == 0 { break }
		fn(buf, dna)
	}
	close(dna)
}


