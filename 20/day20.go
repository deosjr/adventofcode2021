package main

import (
    "math"
    "strings"
    "github.com/deosjr/adventofcode2021/lib"
)

type coord struct {
    x, y int
}

// inf value flips every step: note example input doesnt have this!
func neighbourvalue(image map[coord]bool, c coord, step int) (v int) {
    neighbours := []coord{
        {c.x+1, c.y+1},
        {c.x, c.y+1},
        {c.x-1, c.y+1},
        {c.x+1, c.y},
        {c.x, c.y},
        {c.x-1, c.y},
        {c.x+1, c.y-1},
        {c.x, c.y-1},
        {c.x-1, c.y-1},
    }
    for i, n := range neighbours {
        light, ok := image[n]
        if ok && !light {
            continue
        }
        if !ok && step % 2 == 1 {
            continue
        }
        v += int(math.Exp2(float64(i)))
    }
    return v
}

func main() {
    input := strings.Split(lib.ReadFile(20), "\n\n")
    algorithm := map[int]bool{}
    for i, c := range input[0] {
        algorithm[i] = c == '#'
    }

    image := map[coord]bool{}
    for y, line := range strings.Split(input[1], "\n") {
        if line == "" { continue }
        for x, c := range line {
            image[coord{x,y}] = c == '#'
        }
    }

    var p1image = map[coord]bool{}
    for i:=1; i<=50; i++ {
        newimage := map[coord]bool{}
        for y:=-160;y<160;y++ {
            for x:=-160;x<160;x++ {
                c := coord{x, y}
                v := neighbourvalue(image, c, i)
                newimage[c] = algorithm[v]
            }
        }
        image = newimage
        if i == 2 {
            p1image = newimage
        }
    }

    p1 := 0
    for _, v := range p1image {
        if v {
            p1++
        }
    }
    lib.WritePart1("%d", p1)

    p2 := 0
    for _, v := range image {
        if v {
            p2++
        }
    }
    lib.WritePart2("%d", p2)
}
