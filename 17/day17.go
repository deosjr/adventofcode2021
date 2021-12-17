package main

import (
    "fmt"
    "math"
    "github.com/deosjr/adventofcode2021/lib"
)

func sum(n int) int {
    return n*(n+1)/2
}

func xsteps(vx, step int) int {
    if step > vx {
        return sum(vx)
    }
    return sum(vx) - sum(vx-step)
}

func main() {
    in := lib.ReadFile(17)
    var xmin, xmax, ymin, ymax int
    fmt.Sscanf(in, "target area: x=%d..%d, y=%d..%d", &xmin, &xmax, &ymin, &ymax)

    maxyv := int(math.Abs(float64(ymin))-1)
    lib.WritePart1("%d", sum(maxyv))

    validys := map[int][]int{}
    for vy:=ymin; vy<=maxyv; vy++ {
        i:=0
        y := 0
        dy := vy
        validy := []int{}
        for y >= ymin {
            if y <= ymax {
                validy = append(validy, i)
            }
            i++
            y += dy
            dy -= 1
        }
        validys[vy] = validy
    }

    p2 := 0
    for _, ysteps := range validys {
        for vx:=0; vx<=xmax; vx++ {
            for _, step := range ysteps {
                x := xsteps(vx, step)
                if x > xmax || x < xmin {
                    continue
                }
                p2++
                break
            }
        }
    }
    lib.WritePart2("%d", p2)
}
