package main

import (
    "fmt"
    "math"
    "github.com/deosjr/adventofcode2021/lib"
)

var trianglecache = map[int64]int64{}
var trianglelookup = map[int64]int64{}

func triangle(n int64) int64 {
    if v, ok := trianglecache[n]; ok {
        return v
    }
    v := n*(n+1)/2
    trianglecache[n] = v
    trianglelookup[v] = n
    return v
}

func xsteps(vx, step int64) int64 {
    if step > vx {
        return triangle(vx)
    }
    return triangle(vx) - triangle(vx-step)
}

type stepv struct {
    velocity int64
    step int64
}

type coord struct {
    x, y int64
}

func main() {
    in := lib.ReadFile(17)
    var xmin, xmax, ymin, ymax int64
    fmt.Sscanf(in, "target area: x=%d..%d, y=%d..%d", &xmin, &xmax, &ymin, &ymax)

    var i int64
    for i=0; i<=xmax; i++ {
        triangle(i)
    }

    maxyv := int64(math.Abs(float64(ymin)))

    // map of ycoord to list of options to get there: velocity+in which step
    validys := map[int64][]stepv{}
    for vy:=ymin; vy<=maxyv; vy++ {
        var i, y int64
        dy := vy
        for y >= ymin {
            if y <= ymax {
                validys[y] = append(validys[y], stepv{velocity:vy, step:i})
            }
            i++
            y += dy
            dy -= 1
        }
    }

    validxs := map[int64][]stepv{}
    for x:=xmin; x<=xmax; x++ {
        validxs[x] = append(validxs[x], stepv{velocity:x, step:1})
        var i int64
        for i=0; i<x; i++ {
            from := triangle(i)
            to := from + x
            n, ok := trianglelookup[to]
            if !ok {
                continue
            }
            validxs[x] = append(validxs[x], stepv{velocity:n, step:n-i})
        }
    }

    validvs := map[coord]struct{}{}
    for y:=ymin; y<=ymax; y++ {
        for x:=xmin; x<=xmax; x++ {
            // if theres overlap in steps between validxs and validys, weve found a hit
            vxs := validxs[x]
            vys := validys[y]
            for _, sx := range vxs {
                for _, sy := range vys {
                    if _, ok := validvs[coord{sx.velocity, sy.velocity}]; ok {
                        continue
                    }
                    if sx.step == sy.step {
                        validvs[coord{sx.velocity, sy.velocity}] = struct{}{}
                        continue
                    }
                    // if theres no overlap, theres still a chance. dx could be exactly 0 here
                    if sx.velocity == sx.step && sy.step > sx.step {
                        validvs[coord{sx.velocity, sy.velocity}] = struct{}{}
                    }
                }
            }
        }
    }

    var p1 int64
    for k, _ := range validvs {
        if k.y > p1 {
            p1 = k.y
        }
    }
    lib.WritePart1("%d", triangle(p1))

    lib.WritePart2("%d", len(validvs))
}
