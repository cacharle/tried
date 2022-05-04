package trie

import "fmt"

type Trie struct {
    children map[rune]*Trie
}

func New() *Trie {
    return &Trie{ map[rune]*Trie{} }
}

func (t *Trie) AddString(s string) {
    t.AddRunes([]rune(s))
}

func (t *Trie) AddRunes(s []rune) {
    if len(s) == 0 {
        return
    }
    sub, ok := t.children[s[0]]
    if !ok {
        t.children[s[0]] = New()
        sub = t.children[s[0]]
    }
    sub.AddRunes(s[1:])
}

func (t *Trie) printRec(depth int) {
    for k, v := range t.children {
        for i := 0; i < depth; i++ {
            fmt.Print("  ")
        }
        fmt.Printf("%c\n", k)
        v.printRec(depth + 1)
    }
}

func (t *Trie) Print() {
    t.printRec(0)
}
