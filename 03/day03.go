package main

import (
    "math"

    "github.com/deosjr/adventofcode2021/lib"
)

func getCommon(list [][12]int64) [12]int64 {
    sum := [12]int64{}
    for _, line := range list {
        for i:=0;i<12;i++ {
            sum[i] += line[i]
        }
    }

    for i:=0;i<12;i++{
        if float64(sum[i]) >= float64(len(list))/2.0 {
            sum[i] = 1
        } else {
            sum[i] = 0
        }
    }

    return sum
}

func inv(in [12]int64) [12]int64 {
    out := [12]int64{}
    for i:=0; i<12; i++ {
        if in[i] == 0 {
            out[i] = 1
        }
    }
    return out
}

func toInt(in [12]int64) int64 {
    var out int64
    for i:=0; i<12; i++ {
        if in[11-i] == 0 {
            continue
        }
        out += int64(math.Exp2(float64(i)))
    }
    return out
}

func part2(input [][12]int64, inverse bool) int64 {
    prev := input
    next := [][12]int64{}
    for i:=0; i<12; i++ {
        common := getCommon(prev)
        if inverse {
            common = inv(common)
        }
        var bit int64 = common[i]
        for _, line := range prev {
            if line[i] == bit {
                next = append(next, line)
            }
        }
        if len(next) == 1 {
            return toInt(next[0])
        }
        prev = next
        next = [][12]int64{}
    }
    panic("incorrect")
}

func main() {
    input := [][12]int64{}
    lib.ReadFileByLine(3, func(s string) {
        line := [12]int64{}
        for i, v := range s {
            line[i] = lib.MustParseInt(string(v))
        }
        input = append(input, line)
    })

    sum := getCommon(input)
    lib.WritePart1("%d", toInt(sum) * toInt(inv(sum)))

    lib.WritePart2("%d", part2(input, false) * part2(input, true))
}
