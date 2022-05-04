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

func (t *Trie) String() string {
    ret := ""
    var rec func (t *Trie, depth int)
    rec = func (t *Trie, depth int) {
        for k, v := range t.children {
            for i := 0; i < depth; i++ {
                ret += "  "
            }
            ret += fmt.Sprintf("%c\n", k)
            rec(v, depth + 1)
        }
    }
    rec(t, 0)
    return ret
}
