package main

import (
	"fmt"
	"flag"
)

var a = flag.Int("a", 0, "AA-AA couple")
var b = flag.Int("b", 0, "AA-Aa couple")
var c = flag.Int("c", 0, "AA-aa couple")
var d = flag.Int("d", 0, "Aa-Aa couple")
var e = flag.Int("e", 0, "Aa-aa couple")
var f = flag.Int("f", 0, "aa-aa couple")

func main() {
	flag.Parse()
	fmt.Printf("%f\n", 2 * (float64(*a) * 1 + 1 * float64(*b) + 1 * float64(*c) + 0.75 * float64(*d) + 0.5 * float64(*e)))
}
