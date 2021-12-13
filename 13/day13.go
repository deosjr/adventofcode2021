package main

import (
    "fmt"
    "strings"
    "github.com/deosjr/adventofcode2021/lib"
)

type coord struct {
    x, y int
}

func foldup(m map[coord]struct{}, n int) map[coord]struct{} {
    newm := map[coord]struct{}{}
    for k,_ := range m {
        if k.y < n {
            newm[k] = struct{}{}
            continue
        }
        newm[coord{x:k.x, y:n - (k.y - n)}] = struct{}{}
    }
    return newm
}

func foldleft(m map[coord]struct{}, n int) map[coord]struct{} {
    newm := map[coord]struct{}{}
    for k,_ := range m {
        if k.x < n {
            newm[k] = struct{}{}
            continue
        }
        newm[coord{x:n - (k.x - n), y:k.y}] = struct{}{}
    }
    return newm
}

func main() {
    input := strings.Split(lib.ReadFile(13), "\n\n")
    dots := map[coord]struct{}{}
    var maxx, maxy int
    for _, line := range strings.Split(input[0], "\n") {
        if line == "" { continue }
        var x, y int
        fmt.Sscanf(line, "%d,%d", &x, &y)
        if x > maxx { maxx = x }
        if y > maxy { maxy = y }
        dots[coord{x, y}] = struct{}{}
    }

    var p1 int
    for i, line := range strings.Split(input[1], "\n") {
        var xory rune
        var n int
        fmt.Sscanf(line, "fold along %c=%d", &xory, &n)
        switch xory {
        case 'x' :
            dots = foldleft(dots, n)
        case 'y' :
            dots = foldup(dots, n)
        }
        if i == 0 {
            p1 = len(dots)
        }
    }

    lib.WritePart1("%d", p1)

    lib.WritePart2("")
    for y:=0; y<6; y++ {
        for x:=0; x<39; x++ {
            if _, ok := dots[coord{x,y}]; ok {
                fmt.Print("\u2588")
            } else {
                fmt.Print(".")
            }
        }
        fmt.Println()
    }
}
