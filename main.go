package main

import "github.com/cacharle/tried/trie"

func main() {
    t := trie.New()
    t.AddString("bonjour")
    t.AddString("bonsoir")
    t.Print()
}
