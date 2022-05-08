package lev

import "strings"

func min(xs... int) int {
    curr := xs[0]
    for _, x := range xs[1:] {
        if x < curr {
            curr = x
        }
    }
    return curr
}

func distanceRec(s1, s2 string, cache [][]int) int {
    if s1 == "" {
        return len(s2)
    }
    if s2 == "" {
        return len(s1)
    }
    if s1[0] == s2[0] {
        return distanceCache(s1[1:], s2[1:], cache)
    }
    return 1 + min(
        distanceCache(s1[1:], s2, cache),
        distanceCache(s1, s2[1:], cache),
        distanceCache(s1[1:], s2[1:], cache),
    )
}

func distanceCache(s1, s2 string, cache [][]int) int {
    cached := cache[len(s1)][len(s2)]
    if cached != -1 {
        return cached
    }
    ret := distanceRec(s1, s2, cache)
    cache[len(s1)][len(s2)] = ret
    return ret
}

func Distance(s1, s2 string) int {
    cache := make([][]int, len(s1) + 1)
    for i := range cache {
        cache[i] = make([]int, len(s2) + 1)
        for j := range cache[i] {
            cache[i][j] = -1
        }
    }
    return distanceCache(s1, s2, cache)
}

func DistanceCmpFuncFactory(prefix string, words []string) func (i, j int) bool {
    return func (i, j int) bool {
        prefix = strings.ToLower(prefix)
        iDist := Distance(prefix, strings.ToLower(words[i]))
        jDist := Distance(prefix, strings.ToLower(words[j]))
        return iDist < jDist
    }
}
