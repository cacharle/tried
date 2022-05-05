package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/cacharle/tried/trie"
)

const defaultDictFilePath = "/usr/share/dict/words"

func main() {
    dictFilePath := flag.String("dict", defaultDictFilePath, "file which contains the words registered for autocompletion")
    prefix := flag.String("prefix", "", "print words starting with the prefix")
    flag.Parse()

    t := trie.New()
    file, err := os.Open(*dictFilePath)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Loading dictionary at %v", *dictFilePath)
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        t.Insert(scanner.Text())
    }
    log.Printf("Created trie with %v nodes", t.NodeCount())

    fmt.Printf("Words starting with %#v\n", *prefix)
    for _, w := range t.AtPrefix(*prefix).Words() {
        fmt.Println(w)
    }
}
