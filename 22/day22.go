package main

import (
    "fmt"
    "github.com/deosjr/adventofcode2021/lib"
)

type vector struct {
    x, y, z int
}

type cube struct {
    pmin vector
    pmax vector
    on bool
}

func (c cube) volume() int64 {
    return int64(c.pmax.x - c.pmin.x) * int64(c.pmax.y - c.pmin.y) * int64(c.pmax.z - c.pmin.z)
}

func (a cube) intersect(b cube) (cube, bool) {
    if a.pmin.x > b.pmax.x || a.pmax.x < b.pmin.x {
        return cube{}, false
    }
    if a.pmin.y > b.pmax.y || a.pmax.y < b.pmin.y {
        return cube{}, false
    }
    if a.pmin.z > b.pmax.z || a.pmax.z < b.pmin.z {
        return cube{}, false
    }
    return cube{
        pmin: vector{x: max(a.pmin.x, b.pmin.x), y: max(a.pmin.y, b.pmin.y), z: max(a.pmin.z, b.pmin.z)},
        pmax: vector{x: min(a.pmax.x, b.pmax.x), y: min(a.pmax.y, b.pmax.y), z: min(a.pmax.z, b.pmax.z)},
    }, true
}

// since were subtraction intersection from larger cube, we can make some assumptions
// b will be wholly contained within a, and remainder will keep a's on/off state
// amount of cubes returned depends on amount of overlapping planes
// this can be more efficient i.e. return less cubes (max 3), naive will do for now
func (a cube) sub(b cube) []cube {
    cubes := []cube{}
    if b.pmax.x < a.pmax.x {
        cubes = append(cubes, cube{pmin:vector{b.pmax.x, a.pmin.y, a.pmin.z}, pmax:a.pmax, on:a.on})
        a.pmax.x = b.pmax.x
    }
    if b.pmax.y < a.pmax.y {
        cubes = append(cubes, cube{pmin:vector{a.pmin.x, b.pmax.y, a.pmin.z}, pmax:a.pmax, on:a.on})
        a.pmax.y = b.pmax.y
    }
    if b.pmax.z < a.pmax.z {
        cubes = append(cubes, cube{pmin:vector{a.pmin.x, a.pmin.y, b.pmax.z}, pmax:a.pmax, on:a.on})
        a.pmax.z = b.pmax.z
    }
    if b.pmin.x > a.pmin.x {
        cubes = append(cubes, cube{pmin:a.pmin, pmax:vector{b.pmin.x, a.pmax.y, a.pmax.z}, on:a.on})
        a.pmin.x = b.pmin.x
    }
    if b.pmin.y > a.pmin.y {
        cubes = append(cubes, cube{pmin:a.pmin, pmax:vector{a.pmax.x, b.pmin.y, a.pmax.z}, on:a.on})
        a.pmin.y = b.pmin.y
    }
    if b.pmin.z > a.pmin.z {
        cubes = append(cubes, cube{pmin:a.pmin, pmax:vector{a.pmax.x, a.pmax.y, b.pmin.z}, on:a.on})
    }
    return cubes
}

func max(x, y int) int {
    if x > y { return x } else { return y }
}

func min(x, y int) int {
    if x < y { return x } else { return y }
}

func reboot(current cube, cubes []cube) (sum int64) {
    for i, c := range cubes {
        intersection, ok := current.intersect(c)
        if !ok {
            continue
        }
        // add or substract 'on' volume
        if current.on && !c.on {
            sum += intersection.volume()
        }
        if !current.on && c.on {
            sum -= intersection.volume()
        }
        if i == len(cubes)-1 {
            return sum
        }
        // continue with the nonoverlapping cube(s) if any
        remainder := current.sub(intersection)
        for _, r := range remainder {
            sum += reboot(r, cubes[i+1:])
        }
        return sum
    }
    return 0
}

func main() {
    input := []cube{}
    lib.ReadFileByLine(22, func(s string) {
        var state string
        var xmin, xmax, ymin, ymax, zmin, zmax int
        fmt.Sscanf(s, "%s x=%d..%d,y=%d..%d,z=%d..%d", &state, &xmin, &xmax, &ymin, &ymax, &zmin, &zmax)
        c := cube{pmin: vector{xmin, ymin, zmin}, pmax: vector{xmax+1, ymax+1, zmax+1}, on: state == "on"}
        input = append(input, c)
    })

    init := cube{pmin:vector{-50, -50, -50}, pmax:vector{51, 51, 51}}
    //init := cube{pmin:vector{-99999, -99999, -99999}, pmax:vector{99999, 99999, 99999}}
    cubes := []cube{init}
    var sum int64
    for _, c := range input {
        if c.pmin.x < -50 || c.pmin.y < -50 || c.pmin.z < -50 {
            continue
        }
        if c.pmax.x > 51 || c.pmax.y > 51 || c.pmax.z > 51 {
            continue
        }
        sum += reboot(c, cubes)
        // then add the original cube to the list (NOTE could reuse input instead)
        cubes = append([]cube{c}, cubes...)
    }
    lib.WritePart1("%d", sum)

    init = cube{pmin:vector{-99999, -99999, -99999}, pmax:vector{99999, 99999, 99999}}
    cubes = []cube{init}
    sum = 0
    for _, c := range input {
        sum += reboot(c, cubes)
        cubes = append([]cube{c}, cubes...)
    }
    lib.WritePart2("%d", sum)
}
