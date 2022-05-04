package main

import (
	"fmt"

	"github.com/cacharle/tried/trie"
)

func main() {
    t := trie.New()
    t.AddString("bonjour")
    t.AddString("bonsoir")
    fmt.Print(t)
}
