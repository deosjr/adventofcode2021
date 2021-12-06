package main

import (
    "strings"
    "github.com/deosjr/adventofcode2021/lib"
)

func simDay(before map[int64]int64) map[int64]int64 {
    after := map[int64]int64{}
    for k, v := range before {
        if k == 0 {
            after[6] += v
            after[8] = v
            continue
        }
        after[k-1] += v
    }
    return after
}

func main() {
    input := map[int64]int64{}
    for _, s := range strings.Split(lib.ReadFile(6), ",") {
        n := lib.MustParseInt(strings.TrimSpace(s))
        input[n] += 1
    }
    for i:=0;i<80;i++ {
        input = simDay(input)
    }
    var p1 int64
    for _, v := range input {
        p1 += v
    }

    lib.WritePart1("%d", p1)

    for i:=0;i<256-80;i++ {
        input = simDay(input)
    }
    var p2 int64
    for _, v := range input {
        p2 += v
    }
    lib.WritePart2("%d", p2)
}
