package main

import (
    "github.com/deosjr/adventofcode2021/lib"
)

func main() {
    var input []int64
    lib.ReadFileByLine(1, func(s string) {
        input = append(input, lib.MustParseInt(s))
    })

    count := 0
    prev := input[0]
    for _, n := range input[1:] {
        if n > prev { count++ }
        prev = n
    }
    lib.WritePart1("%d", count)

    count = 0
    prev = input[0] + input[1] + input[2]
    for i:=3; i<len(input);i++ {
        n := input[i-2] + input[i-1] + input[i]
        if n > prev { count++ }
        prev = n
    }
    lib.WritePart2("%d", count)
}
