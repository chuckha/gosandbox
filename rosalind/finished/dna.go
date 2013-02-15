package main

import (
	"fmt"
	"os"
	"io"
)



func main() {
	filename := "rosalind_dna.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	dna := make(chan byte)
	m := make(map[byte]int)
	go FnBytes(filename, dna, Sum)
	for b := range dna {
		m[b] += 1
	}
	fmt.Println(m)
	fmt.Printf("%v %v %v %v\n", m['A'], m['C'], m['G'], m['T'])
}

func Sum(bs []byte, dna chan byte) {
	for _, b := range bs {
		switch b {
		case 'A': dna <- 'A'
		case 'G': dna <- 'G'
		case 'C': dna <- 'C'
		case 'T': dna <- 'T'
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

