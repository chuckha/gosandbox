package main

import (
	"fmt"
	"os"
	"io/ioutil"
)

func main() {
	filename := "REVC.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	file, err := ioutil.ReadFile(filename)
	if err != nil { panic(err) }

	for i := len(file)-1; i > -1; i-- {
		switch file[i] {
			case 'A': fmt.Printf("T")
			case 'T': fmt.Printf("A")
			case 'G': fmt.Printf("C")
			case 'C': fmt.Printf("G")
		}
	}
	fmt.Println()
}

