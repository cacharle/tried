# Tried

A trie data structure implementation in Go.

## Usage

```
$ go run main.go -h
```

### Graphviz

![trie-graph](trie-graph.png)

```
$ go run main.go | dot -Tpng > graph.png
```

## References

- [Wikipedia][1]
- [Jacob Sorber's video][2]
- [Tsoding's live][3]

## TODO

- [ ] Sort the prefix result by edit distance
- [x] Convert trie to a graphviz dot file
- [x] Make autocompletion with ncurses
- [ ] Optimize by triming the nodes with only one child (radix tree)
      ```
      b -> a -> l -> l
             -> s -> e
      ```
      to:
      ```
      ba -> ll
         -> se
      ```

[1]: https://en.wikipedia.org/wiki/Trie
[2]: https://www.youtube.com/watch?v=3CbFFVHQrk4
[3]: https://www.youtube.com/watch?v=2fosrL7I7oc
