package main

import (
	"fmt"
	"github.com/ggnextmap/go-ds/trie"
)

func main () {
	t := trie.NewTrie()
	t.Insert("tea")
	t.Insert("ten")
	t.Insert("tree")

	fmt.Println(t.Find("te"))

	fmt.Println(t)
}
