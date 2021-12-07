package main

import (
    "math"
    "strings"
    "github.com/deosjr/adventofcode2021/lib"
)

func main() {
    var input []int64
    lib.ReadFile(7)
    for _, s := range strings.Split(lib.ReadFile(7), ",") {
        input = append(input, lib.MustParseInt(strings.TrimSpace(s)))
    }

    var min int64 = math.MaxInt64
    for i:=0; i<1000; i++ {
        var fuel int64
        for _, n := range input {
            fuel += int64(math.Abs(float64(int64(i)-n)))
        }
        if fuel < min {
            min = fuel
        }
    }

    lib.WritePart1("%d", min)

    difs := map[int64]int64{}
    var i int64
    for i=0;i<2000; i++ {
        difs[i] = difs[i-1]+i
    }

    min = math.MaxInt64
    for i=0; i<1000; i++ {
        var fuel int64
        for _, n := range input {
            dif := int64(math.Abs(float64(i-n)))
            fuel += difs[dif]
        }
        if fuel < min {
            min = fuel
        }
    }
    lib.WritePart2("%d", min)
}
