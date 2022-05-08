package trie

import "fmt"

type Trie struct {
    children map[rune]*Trie
    end bool
}

func New() *Trie {
    return &Trie{
        children: map[rune]*Trie{},
        end: false,
    }
}

func (t *Trie) Insert(s string) {
    t.InsertRunes([]rune(s))
}

func (t *Trie) InsertRunes(s []rune) {
    if len(s) == 0 {
        t.end = true
        return
    }
    sub, ok := t.children[s[0]]
    if !ok {
        t.children[s[0]] = New()
        sub = t.children[s[0]]
    }
    sub.InsertRunes(s[1:])
}

func (t *Trie) Delete(s string) {
    t.DeleteRunes([]rune(s))
}

// TODO: that deletes the entier prefix
func (t *Trie) deleteRunesRec(s []rune) bool {
    if len(s) == 0 {
        if t.end {
            t.end = false
        }
        return true
    }
    found := t.children[s[0]].deleteRunesRec(s[1:])
    if found {
        delete(t.children, s[0])
    }
    return false
}

func (t *Trie) DeleteRunes(s []rune) {
    t.deleteRunesRec(s)
}

func (t *Trie) Contains(s string) bool {
    return t.ContainsRunes([]rune(s))
}

func (t *Trie) ContainsRunes(s []rune) bool {
    if len(s) == 0 && t.end {
        return true
    }
    sub, ok := t.children[s[0]]
    if !ok {
        return false
    }
    return sub.ContainsRunes(s[1:])
}

func (t *Trie) wordsRec(current []rune, acc *[]string) {
    for c, sub := range t.children {
        word := append(current, c)
        if sub.end {
            *acc = append(*acc, string(word))
        }
        sub.wordsRec(word, acc)
    }
}

func (t *Trie) Words() (acc []string) {
    t.wordsRec([]rune{}, &acc)
    return
}

func (t *Trie) AtPrefix(prefix string) *Trie {
    return t.AtPrefixRunes([]rune(prefix))
}

func (t *Trie) AtPrefixRunes(prefix []rune) *Trie {
    // TODO: use unicode.ToLower to make it case insensitive
    if len(prefix) == 0 {
        return t
    }
    sub, ok := t.children[prefix[0]]
    if !ok {
        return nil
    }
    subPrefix := sub.AtPrefixRunes(prefix[1:])
    if subPrefix == nil {
        return nil
    }
    prefixed := New()
    prefixed.children[prefix[0]] = subPrefix
    return prefixed
}

func (t *Trie) NodeCount() uint {
    var count uint = 1
    for _, sub := range t.children {
        count += sub.NodeCount()
    }
    return count
}

func (t *Trie) String() string {
    ret := ""
    var rec func (t *Trie, depth int)
    rec = func (t *Trie, depth int) {
        for k, v := range t.children {
            for i := 0; i < depth; i++ {
                ret += "  "
            }
            if v.end {
                ret += fmt.Sprintf("%c*\n", k)
            } else {
                ret += fmt.Sprintf("%c \n", k)
            }
            rec(v, depth + 1)
        }
    }
    rec(t, 0)
    return ret
}

func (t *Trie) PrintDot() {
    fmt.Println("digraph trie {")
    fmt.Println("    color=white;")
    fmt.Println("    bgcolor=\"#111111\";")
    t.printDotRec('_')
    fmt.Println("}")
}

func (t *Trie) printDotRec(label rune) {
    fmt.Printf("    node_%p [bgcolor=white] [fontsize=25] [fontcolor=white] [label=%c]", t, label)
    if !t.end {
        fmt.Print(" [color=white]")
    } else {
        fmt.Print(" [color=green]")
    }
    fmt.Println(";")
    for k, sub := range t.children {
        fmt.Printf("    node_%p -> node_%p [bgcolor=white] [color=white];\n", t, sub)
        sub.printDotRec(k)
    }
}
