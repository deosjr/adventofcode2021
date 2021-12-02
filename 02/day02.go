package main

import (
    "strings"

    "github.com/deosjr/adventofcode2021/lib"
)

type coord struct {
    x,y,aim int64
}

func main() {
    var input []coord
    lib.ReadFileByLine(2, func(s string) {
        split := strings.Split(s, " ")
        n := lib.MustParseInt(split[1])
        var instr coord
        switch split[0] {
        case "up":
            instr.y = -n
            instr.aim = -n
        case "down":
            instr.y = n
            instr.aim = n
        case "forward":
            instr.x = n
        }
        input = append(input, instr)
    })

    var x,y int64
    for _, c := range input {
        x += c.x
        y += c.y
    }

    lib.WritePart1("%d", x*y)

    x, y = 0, 0
    var aim int64
    for _, c := range input {
        aim += c.aim
        x += c.x
        y += c.x * aim
    }

    lib.WritePart2("%d", x*y)
}
