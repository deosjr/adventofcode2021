package main

import (
    "github.com/deosjr/adventofcode2021/lib"
)

type coord struct {
    x, y int
}

type octopus struct {
    energy int
    flashed bool
    friends []*octopus
}

func (o *octopus) flash() {
    o.energy += 1
    if o.flashed || o.energy <= 9 {
        return
    }
    o.flashed = true
    for _, f := range o.friends {
        f.flash()
    }
}

func step(cavern map[coord]*octopus) {
    for _, o := range cavern {
        o.flash()
    }
}

func main() {
    cavern := map[coord]*octopus{}
    y := 0
    lib.ReadFileByLine(11, func(s string) {
        for x, c := range s {
            e := int(lib.MustParseInt(string(c)))
            cavern[coord{x, y}] = &octopus{energy:e, flashed:false}
        }
        y++
    })

    for y:=0; y<10; y++ {
        for x:=0; x<10; x++ {
            c := coord{x,y}
            o := cavern[c]
            if f, ok := cavern[coord{x-1,y-1}]; ok {
                o.friends = append(o.friends, f)
            }
            if f, ok := cavern[coord{x,y-1}]; ok {
                o.friends = append(o.friends, f)
            }
            if f, ok := cavern[coord{x+1,y-1}]; ok {
                o.friends = append(o.friends, f)
            }
            if f, ok := cavern[coord{x-1,y}]; ok {
                o.friends = append(o.friends, f)
            }
            if f, ok := cavern[coord{x+1,y}]; ok {
                o.friends = append(o.friends, f)
            }
            if f, ok := cavern[coord{x-1,y+1}]; ok {
                o.friends = append(o.friends, f)
            }
            if f, ok := cavern[coord{x,y+1}]; ok {
                o.friends = append(o.friends, f)
            }
            if f, ok := cavern[coord{x+1,y+1}]; ok {
                o.friends = append(o.friends, f)
            }
        }
    }

    p1 := 0
    p2 := 0
    i := 0
    for {
        i += 1
        step(cavern)
        sum := 0
        for _, o := range cavern {
            if !o.flashed {
                continue
            }
            o.energy = 0
            o.flashed = false
            sum += 1
        }
        if i <= 100 {
            p1 += sum
        }
        if sum == 100 {
            p2 = i
            break
        }
    }

    lib.WritePart1("%d", p1)

    lib.WritePart2("%d", p2)
}
