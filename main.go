package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/cacharle/tried/lev"
	"github.com/cacharle/tried/trie"
	"github.com/gdamore/tcell"
)


const defaultDictFilePath = "/usr/share/dict/words"

func putLineAt(screen tcell.Screen, style tcell.Style, line_num int, content string) {
    for i, c := range content {
        screen.SetContent(i + 1, line_num, c, nil, style)
    }
}

func main() {
    dictFilePath := flag.String("dict", defaultDictFilePath, "File which contains the words registered for autocompletion")
    prefixOption := flag.String("prefix", "", "Get a trie that only starts with a prefix")
    printDot := flag.Bool("dot", false, "Dump a dot representation of the trie for graphviz")
    printWords := flag.Bool("words", false, "Print the words contained in the trie")
    flag.Parse()

    t := trie.New()
    var file *os.File
    if *dictFilePath == "-" {
        file = os.Stdin
    } else {
        var err error
        file, err = os.Open(*dictFilePath)
        if err != nil {
            log.Fatal(err)
        }
    }
    log.Printf("Loading dictionary at %v", *dictFilePath)
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        t.Insert(scanner.Text())
    }
    log.Printf("Created trie with %v nodes", t.NodeCount())

    if *printDot {
        t.AtPrefix(*prefixOption).PrintDot()
        return
    }

    if *printWords {
        fmt.Printf("Words starting with prefix: %#v\n", *prefixOption)
        words := t.AtPrefix(*prefixOption).Words()
        sort.Slice(words, lev.DistanceCmpFuncFactory(*prefixOption, words))
        for _, w := range words {
            fmt.Println(w)
        }
        return
    }

    screen, err := tcell.NewScreen()
    if err != nil {
        log.Fatal(err)
    }
    err = screen.Init()
    if err != nil {
        log.Fatal(err)
    }
    defer screen.Fini()
    style := tcell.StyleDefault.Background(tcell.NewRGBColor(17, 17, 17)).Foreground(tcell.ColorWhite)
    screen.SetStyle(style)
    screen.HideCursor()
    _, height := screen.Size()
    prefix := *prefixOption
    running := true
    for running {
        screen.Clear()
        putLineAt(screen, style, 0, "Help: Escape or Ctrl+c to quit, Ctrl+u to clear prefix")
        putLineAt(screen, style, 1, ">>> " + prefix)
        screen.ShowCursor(len(prefix) + 5, 1)
        foundTrie := t.AtPrefix(prefix)
        if foundTrie != nil {
            words := foundTrie.Words()
            sort.Slice(words, lev.DistanceCmpFuncFactory(prefix, words))
            if len(words) > height {
                words = words[:height]
            }
            for i, w := range words {
                putLineAt(screen, style, i + 2, w)
            }
        }
        screen.Show()
        ev := screen.PollEvent()
        switch ev := ev.(type) {
        case *tcell.EventResize:
            screen.Sync()
        case *tcell.EventKey:
            if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
                running = false
            } else if ev.Key() == tcell.KeyBackspace || ev.Key() == tcell.KeyBackspace2 {
                if len(prefix) > 0 {
                    prefix = prefix[:len(prefix) - 1]
                }
            } else if ev.Key() == tcell.KeyCtrlU {
                prefix = ""
            } else {
                prefix += string(ev.Rune())
            }
        }
    }
}
