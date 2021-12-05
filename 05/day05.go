package main

import (
    "fmt"

    "github.com/deosjr/adventofcode2021/lib"
)

type coord struct {
    x, y int
}

type line struct {
    start, end coord
}

func main() {
    input := []line{}
    lib.ReadFileByLine(5, func(s string) {
        var x1,y1,x2,y2 int
        fmt.Sscanf(s, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
        input = append(input, line{start:coord{x1,y1}, end:coord{x2,y2}})
    })

    oceanfloor := map[coord]int{}
    oceanfloorp2 := map[coord]int{}
    for _, line := range input {
        sx,sy := line.start.x, line.start.y
        ex,ey := line.end.x, line.end.y
        dx, dy := 1, 1
        if sx > ex {
            dx = -1
        }
        if sy > ey {
            dy = -1
        }
        if sx == ex {
            dx = 0
        }
        if sy == ey {
            dy = 0
        }
        for {
            if dx == 0 || dy == 0 {
                oceanfloor[coord{sx, sy}] += 1
            }
            oceanfloorp2[coord{sx, sy}] += 1
            if sx == ex && sy == ey {
                break
            }
            sx += dx
            sy += dy
        }
    }

    p1 := 0
    for _, v := range oceanfloor {
        if v > 1 {
            p1++
        }
    }

    lib.WritePart1("%d", p1)

    p2 := 0
    for _, v := range oceanfloorp2 {
        if v > 1 {
            p2++
        }
    }
    lib.WritePart2("%d", p2)
}
