package main

import (
    "fmt"
    "math"
    "strings"
    "github.com/deosjr/adventofcode2021/lib"
)

type pairstep struct {
    pair string
    step int
}

var (
    mem = map[pairstep]map[rune]int64{}
    rules = map[string]string{}
)

func steps(pair string, n int) map[rune]int64 {
    if occ, ok := mem[pairstep{pair, n}]; ok {
        return occ
    }
    occ := map[rune]int64{}
    if n == 0 {
        occ[rune(pair[0])] += 1
        occ[rune(pair[1])] += 1
        mem[pairstep{pair, n}] = occ
        return occ
    }
    insert := rules[pair]
    occLeft := steps(fmt.Sprintf("%c%s", pair[0], insert), n-1)
    occRight := steps(fmt.Sprintf("%s%c", insert, pair[1]), n-1)
    for k,v := range occLeft {
        occ[k] += v
    }
    for k,v := range occRight {
        occ[k] += v
    }
    occ[rune(insert[0])] -= 1
    mem[pairstep{pair, n}] = occ
    return occ
}

func mostMinLeast(occ map[rune]int64) int64 {
    var most int64 = math.MinInt64
    var least int64 = math.MaxInt64
    for _,v := range occ {
        if v > most {
            most = v
        }
        if v < least {
            least = v
        }
    }
    return most-least
}

func main() {
    input := strings.Split(lib.ReadFile(14), "\n\n")
    polymer := input[0]
    for _, rule := range strings.Split(input[1], "\n") {
        if rule == "" { continue }
        var adjacent, insert string
        fmt.Sscanf(rule, "%s -> %s", &adjacent, &insert)
        rules[adjacent] = insert
    }

    occurences := map[rune]int64{}
    for i:=0; i<len(polymer)-1; i++ {
        pair := polymer[i:i+2]
        occ := steps(pair, 10)
        for k,v := range occ {
            occurences[k] += v
        }
        if i != 0 {
            // prevent double counting
            occurences[rune(pair[0])] -= 1
        }
    }
    lib.WritePart1("%d", mostMinLeast(occurences))

    occurences = map[rune]int64{}
    for i:=0; i<len(polymer)-1; i++ {
        pair := polymer[i:i+2]
        occ := steps(pair, 40)
        for k,v := range occ {
            occurences[k] += v
        }
        if i != 0 {
            // prevent double counting
            occurences[rune(pair[0])] -= 1
        }
    }
    lib.WritePart2("%d", mostMinLeast(occurences))
}
