package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	numWords := 10000
	for i := 0; i < numWords; i++ {
		rand.Seed(time.Now().UnixNano())
		word := ""
		for j := 0; j < (rand.Int() % 7) + 8; j++ {
			word += string((rand.Int() % 57) + 65)
		}
		fmt.Println(word)
	}
}
