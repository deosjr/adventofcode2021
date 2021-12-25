package main

import (
    "github.com/deosjr/adventofcode2021/lib"
)

type coord struct {
    x, y int
}

func main() {
    east := map[coord]struct{}{}
    south := map[coord]struct{}{}
    y := 0
    var maxx, maxy int
    lib.ReadFileByLine(25, func(s string) {
        for x, c := range s {
            if c == '>' {
                east[coord{x,y}] = struct{}{}
                continue
            }
            if c == 'v' {
                south[coord{x,y}] = struct{}{}
            }
        }
        y++
        maxy = y
        maxx = len(s)
    })

    step := 0
    num := 1
    tempmap := map[coord]struct{}{}
    for num > 0 {
        num = 0
        step++
        for k, _ := range east {
            next := coord{x: (k.x+1) % maxx, y:k.y}
            if _, ok := east[next]; ok {
                tempmap[k] = struct{}{}
                continue
            }
            if _, ok := south[next]; ok {
                tempmap[k] = struct{}{}
                continue
            }
            tempmap[next] = struct{}{}
            num++
        }
        east = tempmap
        tempmap = map[coord]struct{}{}
        for k, _ := range south {
            next := coord{x: k.x, y: (k.y+1) % maxy}
            if _, ok := east[next]; ok {
                tempmap[k] = struct{}{}
                continue
            }
            if _, ok := south[next]; ok {
                tempmap[k] = struct{}{}
                continue
            }
            tempmap[next] = struct{}{}
            num++
        }
        south = tempmap
        tempmap = map[coord]struct{}{}
    }

    lib.WritePart1("%d", step)
}
