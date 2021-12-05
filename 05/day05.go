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
        if sx == ex {
            if ey < sy {
                sy, ey = ey, sy
            }
            for y:=sy; y<=ey; y++ {
                oceanfloor[coord{sx, y}] += 1
                oceanfloorp2[coord{sx, y}] += 1
            }
            continue
        }
        if sy == ey {
            if ex < sx {
                sx, ex = ex, sx
            }
            for x:=sx; x<=ex; x++ {
                oceanfloor[coord{x, sy}] += 1
                oceanfloorp2[coord{x, sy}] += 1
            }
            continue
        }
        //diagonal
        dx, dy := 1, 1
        if sx > ex {
            dx = -1
        }
        if sy > ey {
            dy = -1
        }
        for {
            oceanfloorp2[coord{sx, sy}] += 1
            if sx == ex {
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
