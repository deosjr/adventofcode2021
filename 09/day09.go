package main

import (
    "sort"
    "strings"
    "github.com/deosjr/adventofcode2021/lib"
)

type coord struct {
    x, y int
}

type basin struct {
    size int
}

func main() {
    heightmap := map[coord]int64{}
    var maxx, maxy int
    for y, line := range strings.Split(lib.ReadFile(9), "\n") {
        if line == "" { continue }
        maxx = len(line)
        maxy = y
        for x, c := range strings.Split(line, "") {
            n := lib.MustParseInt(c)
            heightmap[coord{x,y}] = n
        }
    }

    candidates := map[coord]struct{}{}
    basinmap := map[coord]*basin{}
    basins := []*basin{}

    var p1 int64
    for y:=0; y<=maxy; y++ {
        for x:=0; x<maxx; x++ {
            // p1
            c := coord{x,y}
            cup := coord{x, y-1}
            cleft := coord{x-1, y}
            height := heightmap[c]
            up, upok := heightmap[cup]
            left, leftok := heightmap[cleft]
            // check only left and up. if we are smaller than both, we are a candidate
            checkup := !upok || height < up
            checkleft := !leftok || height < left
            if checkleft && checkup {
                // if theres nothing below to confirm, immediately confirm
                candidates[c] = struct{}{}
            }
            // if up exists, is smaller than us, and is a candidate, it is now a confirmed basin
            _, ok := candidates[cup]
            if upok && up < height && ok {
                p1 += up+1
            }
            // if left exists, is not smaller than us, and is a candidate, remove it from candidates
            _, ok = candidates[cleft]
            if leftok && left >= height && ok {
                delete(candidates, cleft)
            }

            // p2
            if height == 9 {
                continue
            }
            upcheck := !upok || up==9
            leftcheck := !leftok || left==9
            if upcheck && leftcheck {
                // new basin
                basinmap[c] = &basin{size:1}
                basins = append(basins, basinmap[c])
                continue
            }
            bup, bupok := basinmap[cup]
            bleft, bleftok := basinmap[cleft]
            if bupok && bleftok {
                basinmap[c] = bup
                bup.size += 1
                // if basins up and left, merge
                if bup == bleft {
                    continue
                }
                // bup and left are different, actual merge
                bup.size += bleft.size
                for k, v := range basinmap {
                    if v == bleft {
                        basinmap[k] = bup
                    }
                }
                continue }
            if bupok {
                bup.size += 1
                basinmap[c] = bup
                continue
            }
            if bleftok {
                bleft.size += 1
                basinmap[c] = bleft
                continue
            }
        }
    }
    // confirm the last row
    for x:=0; x<maxx; x++ {
        if _, ok := candidates[coord{x, maxy}]; !ok {
            continue
        }
        p1 += heightmap[coord{x, maxy}] + 1
    }

    lib.WritePart1("%d", p1)

    sort.Slice(basins, func(i, j int) bool {
		return basins[i].size > basins[j].size
	})

    var p2 int = 1
    for _, v := range basins[:3] {
        p2 *= v.size
    }

    lib.WritePart2("%d", p2)
}
