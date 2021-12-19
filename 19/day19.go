package main

import (
    "fmt"
    "math"
    "strings"
    "github.com/deosjr/adventofcode2021/lib"
)

type coord struct {
    x, y, z int64
}

// turn left with respect to z (i.e. keeping z fixed)
func (c coord) turnLeft() coord {
    return coord{-c.y, c.x, c.z}
}

func (c coord) add(d coord) coord {
    return coord{c.x+d.x, c.y+d.y, c.z+d.z}
}

func manhattan(a, b coord) int64 {
    return int64(math.Abs(float64(a.x - b.x)) + math.Abs(float64(a.y - b.y)) + math.Abs(float64(a.z - b.z)))
}

// keeping z fixed, generate the permutations of beacons wrt rotation
func permutationsUp(beacons []coord) [][]coord {
    leftOne := make([]coord, len(beacons))
    leftTwo := make([]coord, len(beacons))
    leftThree := make([]coord, len(beacons))
    for i, c := range beacons {
        c = c.turnLeft()
        leftOne[i] = c
        c = c.turnLeft()
        leftTwo[i] = c
        c = c.turnLeft()
        leftThree[i] = c
    }
    return [][]coord{beacons, leftOne, leftTwo, leftThree}
}

// keeping rotation fixed, generate the permutations of beacons wrt facing
func permutationsFacing(beacons []coord) [][]coord {
    negZ := make([]coord, len(beacons))
    posY := make([]coord, len(beacons))
    negY := make([]coord, len(beacons))
    posX := make([]coord, len(beacons))
    negX := make([]coord, len(beacons))
    for i, c := range beacons {
        negZ[i] = coord{-c.x, c.y, -c.z}
        posY[i] = coord{c.x, -c.z, c.y}
        negY[i] = coord{c.x, c.z, -c.y}
        posX[i] = coord{-c.z, c.y, c.x}
        negX[i] = coord{c.z, c.y, -c.x}
    }
    return [][]coord{beacons, negZ, posY, negY, posX, negX}
}

func orient(beacons []coord) [][]coord {
    orientations := [][]coord{}
    for _, perm := range permutationsFacing(beacons) {
        orientations = append(orientations, permutationsUp(perm)...)
    }
    return orientations
}

// returns all beacons oriented to from-scanner if overlap is found
func findOverlap(from, to []coord) ([]coord, coord, bool) {
    for _, orientatedBeacons := range orient(to) {
        overlap := map[coord]int{}
        for _, s0 := range from {
            for _, s1 := range orientatedBeacons {
                dx := s0.x - s1.x
                dy := s0.y - s1.y
                dz := s0.z - s1.z
                overlap[coord{dx, dy, dz}] += 1
            }
        }
        for k,v := range overlap {
            if v < 12 {
                continue
            }
            oriented := make([]coord, len(to))
            for i, b := range orientatedBeacons {
                oriented[i] = b.add(k)
            }
            return oriented, k, true
        }
    }
    return nil, coord{}, false
}

func main() {
    scanners := [][]coord{}
    for _, scanner := range strings.Split(lib.ReadFile(19), "\n\n") {
        beacons := []coord{}
        for _, line := range strings.Split(scanner, "\n")[1:] {
            if line == "" { continue }
            var x, y, z int64
            fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
            beacons = append(beacons, coord{x, y, z})
        }
        scanners = append(scanners, beacons)
    }

    orientedScanners := map[int][]coord{}
    orientedScanners[0] = scanners[0]
    scannerLocations := []coord{{0,0,0}}

    toCheck := []int{0}
    for len(toCheck) > 0 {
        fromIndex := toCheck[0]
        toCheck = toCheck[1:]
        for i:=0; i<len(scanners); i++ {
            if _, ok := orientedScanners[i]; ok {
                continue
            }
            beacons, d, ok := findOverlap(orientedScanners[fromIndex], scanners[i])
            if !ok {
                continue
            }
            orientedScanners[i] = beacons
            scannerLocations = append(scannerLocations, d)
            toCheck = append(toCheck, i)
        }
    }

    beacons := map[coord]struct{}{}
    for _, v := range orientedScanners {
        for _, b := range v {
            beacons[b] = struct{}{}
        }
    }
    lib.WritePart1("%d", len(beacons))

    var p2 int64
    for i:=0; i<len(scanners); i++ {
        for j:=i+1;j<len(scanners); j++ {
            m := manhattan(scannerLocations[i], scannerLocations[j])
            if m > p2 {
                p2 = m
            }
        }
    }
    lib.WritePart2("%d", p2)
}
