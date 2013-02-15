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
	filename := "GC.txt"
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
	largest := dnas[0]
	for _, d := range dnas {
		if d.Gc > largest.Gc {
			largest = d
		}
	}

	fmt.Println(largest.Name)
	fmt.Println(largest.Gc)

}

func findGc (dna string) float64 {
	count := strings.Count(dna, "C") + strings.Count(dna, "G")
	return float64(count)/float64(len(dna)) * 100
}


