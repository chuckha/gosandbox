package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
)

func main() {
	filename := "SUBS.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	file, err := os.Open(filename)
	if err != nil { panic(err) }
	defer file.Close()

	reader := bufio.NewReader(file)

	count := 0

	var dna string
	var substr string
	for {
		line, e := reader.ReadString('\n')
		if e == io.EOF { break }
		if e != nil { panic("GetLines: " + e.Error()) }
		if count == 0 {
			dna = line[:len(line)-1]
			count++
		} else {
			substr = line[:len(line)-1]
		}
	}
	for i := 0; i<= len(dna) - len(substr); i++ {
		if dna[i:i+len(substr)] == substr {
			fmt.Printf("%d ", i+1)
		}
	}
	fmt.Println()
}
