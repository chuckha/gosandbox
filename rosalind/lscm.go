package main

import (
	"github.com/ggnextmap/go-bio/reader"
	"fmt"
	"index/suffixarray"
	"os"
)

func main () {
	dnas, err := reader.ReadSeqs("LSCM.fasta")
	if err != nil { panic(err) }

	dna := dnas[0]

	sa := suffixarray.New([]byte(dna.Dna))

	fmt.Println("the suffix array")
	sa.Write(os.Stdout)
}

