package main

import (
	"fmt"
	"flag"
	"math/big"
)

var ki = flag.Int("k", 0, "")
var mi = flag.Int("m", 0, "")
var ni = flag.Int("n", 0, "")

func main() {
	flag.Parse()
	// Input the dataset
	k := Big(*ki)
	m := Big(*mi)
	n := Big(*ni)

	// Get the total number of combos
	total := Big(0)
	total.Add(total, k)
	total.Add(total, m)
	total.Add(total, n)

	// How many gene pairings have a dominant trait?
	numK := NumK(k, m, n, total)

	// How many possible gene pairings are there?
	perms := Permutation(total, Big(2))

	// Get the number of possible Homozygous genes
	mm := Homozygous(m)
	nn := Homozygous(n)
	// Get the number of possible Heterozygous genes
	mn := Heterozygous(perms, numK, mm, nn)

	// For Aa x Aa { AA, Aa, Aa, aa } 
	// So there is only 1 aa
	numMm := new(big.Int).Mul(mm, Big(1))
	// For aa x aa  there are 4 aa
	numNn := new(big.Int).Mul(nn, Big(4))
	// For Aa x aa there are 2 aa
	numMn := new(big.Int).Mul(mn, Big(2))
	numMm.Add(numMm, numNn)
	numMm.Add(numMm, numMn)
	// numMm is how many possible "aa"s there are

	genes := new(big.Int).Mul(perms, Big(4))
	// This is the inverse -- the number of A*s there are
	genes.Sub(genes, numMm)

	fmt.Println(genes)
	// This is how many possibilities there are
	all := new(big.Int).Mul(perms, Big(4))

	ratio := new(big.Rat)
	ratio.SetFrac(genes, all)
	fmt.Println(ratio)

}

func NumK (k, m, n, total *big.Int) *big.Int {
	l := new(big.Int).Sub(total, Big(1))
	l.Mul(l, k)
	r := new(big.Int).Add(m, n)
	r.Mul(r, k)
	l.Add(l, r)
	return l
}

func Permutation (n, k *big.Int) *big.Int {
	t := Fact(n)
	t.Div(t, Fact(new(big.Int).Sub(n, k)))
	return t
}

func Heterozygous (perms, numK, mm, nn *big.Int) *big.Int {
	hetero := new(big.Int).Sub(perms, numK)
	hetero.Sub(hetero, mm)
	hetero.Sub(hetero, nn)
	return hetero
}

func Homozygous (x *big.Int) *big.Int {
	// copy x and find its factorial
	c := Fact(new(big.Int).Add(x, Big(0)))
	d := Fact(new(big.Int).Sub(x, Big(2)))
	c.Div(c,d)
	return c
}


func Big (x int) *big.Int {
	return big.NewInt(int64(x))
}

/// Have to use bigint :(
func Fact (x *big.Int) *big.Int {
	factorial := big.NewInt(1)
	tmp := new(big.Int).Add(x, Big(0))
	if x.Cmp(Big(0)) == -1 {
		return Big(1)
	} else {
		for i := tmp; i.Cmp(Big(1))==1; i.Sub(i, Big(1)) {
			factorial = factorial.Mul(factorial, i)
		}
	}
	return factorial
}
