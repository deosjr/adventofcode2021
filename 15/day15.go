package main

import (
    "math"
    "github.com/deosjr/adventofcode2021/lib"
    "github.com/deosjr/Pathfinding/path"
)

type coord struct {
    x, y int
}

type cave struct {
    coords map[coord]int64
}

func (c *cave) Neighbours(n path.Node) []path.Node {
    p := n.(coord)
	x, y := p.x, p.y
	points := []path.Node{}
	points2d := []coord{
		{x - 1, y},
		{x, y - 1},
		{x, y + 1},
		{x + 1, y},
    }
    for _, p2d := range points2d {
		if _, ok := c.coords[p2d]; ok {
			points = append(points, p2d)
		}
	}
	return points
}

func (c *cave) G(pn, qn path.Node) float64 {
    return float64(c.coords[qn.(coord)])
}

func (c *cave) H(pn, qn path.Node) float64 {
    p, q := pn.(coord), qn.(coord)
	return euclidian2d(p, q)
}

func euclidian2d(p, q coord) float64 {
	dx := float64(q.x - p.x)
	dy := float64(q.y - p.y)
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
    input := map[coord]int64{}
    y:=0
    var maxx, maxy int
    lib.ReadFileByLine(15, func(s string) {
        maxy = y
        maxx = len(s) - 1
        for x, c := range s {
            input[coord{x, y}] = lib.MustParseInt(string(c))
        }
        y++
    })

    cave := &cave{coords:input}
    start := coord{0,0}
    end := coord{maxx,maxy}
    route, _ := path.FindRoute(cave, start, end)

    var p1 int64
    for _, c := range route[:len(route)-1] {
        p1 += input[c.(coord)]
    }

    lib.WritePart1("%d", p1)

    // embiggen
    for ry:=0; ry<5; ry++ {
        for rx:=0; rx<5; rx++ {
            if rx==0 && ry==0 { continue }
            for y:=0; y<=maxy; y++ {
                for x:=0; x<=maxx; x++ {
                    c := coord{x:x+rx*(maxx+1), y:y+ry*(maxy+1)}
                    v := input[coord{x, y}]
                    v += int64(rx + ry)
                    if v > 9 {
                        v -= 9
                    }
                    if v > 9 {
                        v -= 9
                    }
                    input[c] = v
                }
            }
        }
    }

    end = coord{5*(maxx+1)-1,5*(maxy+1)-1}
    route, _ = path.FindRoute(cave, start, end)

    var p2 int64
    for _, c := range route[:len(route)-1] {
        p2 += input[c.(coord)]
    }
    lib.WritePart2("%d", p2)
}
