package main

import (
	"fmt"
	"bio/reader"
)

func main() {
	dnas, err := bio.ReadSeqs("GRPH.txt")
	if err != nil {panic(err)}

	pre := make(map[string][]bio.Dna, 0)
	suf := make(map[string][]bio.Dna, 0)

	for _, dna := range dnas {
		suffix, prefix := DnaSuffix(dna), DnaPrefix(dna)
		suf[suffix] = append(suf[suffix], dna)
		pre[prefix] = append(pre[prefix], dna)
	}

	for suffix, dna := range suf {
		for _, suffixDna := range dna {
			for _, prefixDna := range pre[suffix] {
				if suffixDna != prefixDna {
					fmt.Printf("%v %v\n", suffixDna.Name, prefixDna.Name)
				}
			}
		}
	}
}

func DnaPrefix(dna bio.Dna) string {
	return dna.Dna[:3]
}

func DnaSuffix(dna bio.Dna) string {
	return dna.Dna[len(dna.Dna)-3:]
}
