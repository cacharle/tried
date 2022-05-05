package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/cacharle/tried/trie"
)

func main() {
    t := trie.New()

    file, err := os.Open("/usr/share/dict/words")
    if err != nil {
        log.Fatal(err)
    }
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        t.Insert(scanner.Text())
    }

    // t.Insert("bonsoir")
    // t.Insert("bonjour")
    // t.Insert("aurevoir")
    // t.Insert("good")
    // t.Insert("goodbye")
    // // fmt.Print(t)
    //
    // t.Delete("goodbye")
    // t.Delete("bonsoir")

    // fmt.Println(t.NodeCount())

    fmt.Println(len(t.Words()))
    t.Delete("hell")
    fmt.Println(len(t.Words()))

    // for _, w := range t.Words() {
    //     fmt.Println(w)
    // }



    // fmt.Println(t.Contains("bonjour"))
}
