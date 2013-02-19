package main

import (
	"bufio"
	"fmt"
	"os"
	"io"
)

func main() {
	filename := "HAMM.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	file, err := os.Open(filename)
	if err != nil { panic(err) }
	defer file.Close()

	reader := bufio.NewReader(file)

	dnas := make([]string, 0)

	for {
		line, e := reader.ReadString('\n')
		if e == io.EOF { break }
		if e != nil { panic("GetLines: " + e.Error()) }
		dnas = append(dnas, line[:len(line)-1])
	}
	fmt.Println(hamm(dnas[0], dnas[1]))
}

func hamm(dna1, dna2 string) int {
	// both dna1 and dna2 are of equal length
	count := 0
	for i := 0; i < len(dna1); i++ {
		if dna1[i] != dna2[i] {
			count += 1
		}
	}
	return count
}
