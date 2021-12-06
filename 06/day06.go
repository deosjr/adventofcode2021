package main

import (
    "strings"
    "github.com/deosjr/adventofcode2021/lib"
)

func simulateDay(fish map[int64]int64) {
    temp := fish[0]
    var i int64
    for i=1; i<=8; i++ {
        fish[i-1] = fish[i]
    }
    fish[8] = temp
    fish[6] += temp
}

func main() {
    input := map[int64]int64{}
    for _, s := range strings.Split(lib.ReadFile(6), ",") {
        n := lib.MustParseInt(strings.TrimSpace(s))
        input[n] += 1
    }
    for i:=0;i<80;i++ {
        simulateDay(input)
    }
    var p1 int64
    for _, v := range input {
        p1 += v
    }

    lib.WritePart1("%d", p1)

    for i:=0;i<256-80;i++ {
        simulateDay(input)
    }
    var p2 int64
    for _, v := range input {
        p2 += v
    }
    lib.WritePart2("%d", p2)
}
