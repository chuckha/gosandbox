package main

import (
	"bufio"
	"fmt"
	"os"
	"io"
	"strings"
)

type Dna struct {
	Name string
	Dna string
	Gc float64
}

func main() {
	filename := "CONS.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	file, err := os.Open(filename)
	if err != nil { panic(err) }
	defer file.Close()

	reader := bufio.NewReader(file)

	dnas := make([]Dna, 0)

	var tn string
	var dna string

	for {
		line, e := reader.ReadString('\n')
		if e == io.EOF {
			dnas = append(dnas, Dna{tn, dna, findGc(dna)})
			break
		}
		if e != nil { panic("GetLines: " + e.Error()) }
		if line[0] == '>' {
			if tn != "" {
				dnas = append(dnas, Dna{tn, dna, findGc(dna)})
			}
			tn = line[1:len(line)-1]
			dna = ""
		} else {
			dna += line[:len(line)-1]
		}
	}

	length := len(dnas[0].Dna)

	cons := make([][]int, 4)
	for i := 0; i < 4; i++ {
		cons[i] = make([]int, length)
	}
	// 0 => A
	// 1 => C
	// 2 => G
	// 3 => T

	// Assumption DNAs are same length

	for i := 0; i < length; i++ {
		for _, d := range dnas {
			switch d.Dna[i] {
			case 'A': cons[0][i] += 1
			case 'C': cons[1][i] += 1
			case 'G': cons[2][i] += 1
			case 'T': cons[3][i] += 1
			}
		}
	}
	for j := 0; j < length; j++ {
		largest := cons[0][j]
		val := "A"
		if cons[1][j] > largest {
			largest = cons[1][j]
			val = "C"
		}
		if cons[2][j] > largest {
			largest = cons[2][j]
			val = "G"
		}
		if cons[3][j] > largest {
			largest = cons[3][j]
			val = "T"
		}
		fmt.Printf("%s", val)
	}
	fmt.Println()

	PrintSlice("A", cons[0])
	PrintSlice("C", cons[1])
	PrintSlice("G", cons[2])
	PrintSlice("T", cons[3])
}

func findGc (dna string) float64 {
	count := strings.Count(dna, "C") + strings.Count(dna, "G")
	return float64(count)/float64(len(dna)) * 100
}

func PrintSlice (label string, s []int) {
	fmt.Printf("%s:", label)
	for _, v := range s {
		fmt.Printf(" %d", v)
	}
	fmt.Println()
}
