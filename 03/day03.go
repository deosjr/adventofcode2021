package main

import (
    "math"
    "strconv"

    "github.com/deosjr/adventofcode2021/lib"
)

var bitlength = 12 // set to 5 when using lib.Test()

func getCommon(in []uint64) uint64 {
    var common uint64
    threshold := float64(len(in))/2.0
    for i:=0; i<bitlength; i++ {
        sum := 0
        for _, n := range in {
            if (n>>i)&1 == 1 {
                sum++
            }
        }
        if float64(sum) >= threshold {
            common += 1<<i
        }
    }
    return common
}

func inv(in uint64) uint64 {
    return in^uint64(math.Exp2(float64(bitlength))-1)
}

func part2(input []uint64, inverse bool) uint64 {
    prev := input
    next := []uint64{}
    for i:=bitlength-1; i>=0; i-- {
        common := getCommon(prev)
        if inverse {
            common = inv(common)
        }
        var bit uint64 = (common>>i)&1
        for _, n := range prev {
            if (n>>i)&1 == bit {
                next = append(next, n)
            }
        }
        if len(next) == 1 {
            break
        }
        prev = next
        next = []uint64{}
    }
    return next[0]
}

func main() {
    input := []uint64{}
    lib.ReadFileByLine(3, func(s string) {
        n, _ := strconv.ParseUint(s, 2, 64)
        input = append(input, n)
    })

    common := getCommon(input)
    lib.WritePart1("%d", common * inv(common))

    lib.WritePart2("%d", part2(input, false) * part2(input, true))
}
