package main

import (
    "strings"
    "github.com/deosjr/adventofcode2021/lib"
)

type coord struct {
    x, y int
}

type step struct {
    from coord
    to coord
    typ int
}

type amphipod struct {
    pos coord
    typ int
}

func (c coord) blockedInRoom(amphipods map[coord]int) bool {
    for y:=c.y-1; y>0; y-- {
        if _, ok := amphipods[coord{x:c.x, y:y}]; ok {
            return true
        }
    }
    return false
}

func (c coord) inHallway() bool {
    return c.y == 0
}

func (c coord) organized(amphipods map[coord]int, roomSize, typ int) bool {
    if c.inHallway() || c.x != 2*(typ+1) {
        return false
    }
    for y:=c.y+1; y<=roomSize; y++ {
        pos := coord{c.x, y}
        t, ok := amphipods[pos]
        if !ok {
            continue
        }
        if t != typ {
            return false
        }
    }
    return true
}

func organized(amphipods map[coord]int, roomSize int) bool {
    for pos, typ := range amphipods {
        if !pos.organized(amphipods, roomSize, typ) {
            return false
        }
    }
    return true
}

func destination(amphipods map[coord]int, typ, roomSize int) (coord, bool) {
    destx := 2*(typ+1)
    desty := roomSize
    for y:=desty; y>0; y-- {
        desty = y
        t, ok := amphipods[coord{destx, desty}]
        if !ok {
            break
        }
        if t != typ || y == 1 {
            return coord{}, false
        }
    }
    return coord{destx, desty}, true
}

func pathBlocked(amphipods map[coord]int, from, to coord) bool {
    if from.x < to.x {
        for i:=from.x+1; i<to.x; i++ {
            if _, ok := amphipods[coord{x:i, y:0}]; ok {
                return true
            }
        }
    } else {
        for i:=from.x-1; i>to.x; i-- {
            if _, ok := amphipods[coord{x:i, y:0}]; ok {
                return true
            }
        }
    }
    return false
}

func cost(solution []step) int {
    sum := 0
    for _, s := range solution {
        dx := s.to.x - s.from.x
        if dx < 0 { dx = -dx }
        dy := s.to.y - s.from.y
        if dy < 0 { dy = -dy }
        exp := 1
        for i:=0; i<s.typ; i++ {
            exp *= 10
        }
        sum += (dx+dy) * exp
    }
    return sum
}

func mincostremaining(amphipods map[coord] int, roomSize int) int {
    sum := 0
    for pos, typ := range amphipods {
        if pos.organized(amphipods, roomSize, typ) {
            continue
        }
        destx := 2*(typ+1)
        dx := destx - pos.x
        if dx < 0 { dx = -dx }
        dy := pos.y + 1
        exp := 1
        for i:=0; i<typ; i++ {
            exp *= 10
        }
        sum += (dx+dy) * exp
    }
    return sum
}

func recurse(amphipodmap map[coord]int, amphipods []*amphipod, a *amphipod, lenHallway, roomSize int, newstep step, answer, sofar []step) []step {
    delete(amphipodmap, newstep.from)
    amphipodmap[newstep.to] = newstep.typ
    a.pos = newstep.to
    newsofar := make([]step, len(sofar)+1)
    copy(newsofar, sofar)
    newsofar[len(sofar)] = newstep
    sol, ok := solve(amphipodmap, amphipods, lenHallway, roomSize, newsofar)
    delete(amphipodmap, newstep.to)
    amphipodmap[newstep.from] = newstep.typ
    a.pos = newstep.from
    if !ok {
        return answer
    }
    if len(answer) == 0 || cost(answer) > cost(sol) {
        answer = sol
    }
    return answer
}

func solve(amphipodmap map[coord]int, amphipods []*amphipod, lenHallway, roomSize int, sofar []step) ([]step, bool) {
    c := cost(sofar)
    if c >= min {
        return nil, false
    }
    // if we are in a correct configuration, return
    if organized(amphipodmap, roomSize) {
        if c < min {
            min = c
            return sofar, true
        }
        return nil, false
    }
    if c + mincostremaining(amphipodmap, roomSize) >= min {
        return nil, false
    }

    // else take all possible steps, and find best answer recursively
    var answer []step
    // all possible steps into the hallway
    for _, a := range amphipods {
        pos, typ := a.pos, a.typ
        if pos.inHallway() || pos.organized(amphipodmap, roomSize, typ) || pos.blockedInRoom(amphipodmap) {
            continue
        }
        for i:=pos.x-1; i>=0; i-- {
            if i%2 == 0 && i > 0 && i < 10 {
                continue
            }
            newpos := coord{x:i, y:0}
            if _, ok := amphipodmap[newpos]; ok {
                break
            }
            newstep := step{from:pos, to:newpos, typ:typ}
            answer = recurse(amphipodmap, amphipods, a, lenHallway, roomSize, newstep, answer, sofar)
        }
        for i:=pos.x+1; i<lenHallway; i++ {
            if i%2 == 0 && i > 0 && i < 10 {
                continue
            }
            newpos := coord{x:i, y:0}
            if _, ok := amphipodmap[newpos]; ok {
                break
            }
            newstep := step{from:pos, to:newpos, typ:typ}
            answer = recurse(amphipodmap, amphipods, a, lenHallway, roomSize, newstep, answer, sofar)
        }
    }
    // all possible steps from hallway to room
    for _, a := range amphipods {
        pos, typ := a.pos, a.typ
        if !pos.inHallway() {
            continue
        }
        dest, ok := destination(amphipodmap, typ, roomSize)
        if !ok {
            continue
        }
        if pathBlocked(amphipodmap, pos, dest) {
            continue
        }
        newstep := step{from:pos, to:dest, typ:typ}
        answer = recurse(amphipodmap, amphipods, a, lenHallway, roomSize, newstep, answer, sofar)
    }
    if cost(answer) > min {
        return nil, false
    }
    if answer == nil {
        return nil, false
    }
    return answer, true
}

func solveP1(amphipods []*amphipod, hallway int) int {
    m := map[coord]int{}
    for _, a := range amphipods {
        m[a.pos] = a.typ
    }
    sol, _ := solve(m, amphipods, hallway, 2, []step{})
    return cost(sol)
}

func solveP2(amphipods []*amphipod, hallway int) int {
    m := map[coord]int{}
    for _, a := range amphipods {
        if a.pos.y == 2 {
            a.pos.y = 4
        }
        m[a.pos] = a.typ
    }
    // add:   #D#C#B#A#
    //        #D#B#A#C#
    for _, a := range []*amphipod{
        {pos: coord{2,2}, typ: 3},
        {pos: coord{2,3}, typ: 3},
        {pos: coord{4,2}, typ: 2},
        {pos: coord{4,3}, typ: 1},
        {pos: coord{6,2}, typ: 1},
        {pos: coord{6,3}, typ: 0},
        {pos: coord{8,2}, typ: 0},
        {pos: coord{8,3}, typ: 2},
    }{
        m[a.pos] = a.typ
        amphipods = append(amphipods, a)
    }
    sol, _ := solve(m, amphipods, hallway, 4, []step{})
    return cost(sol)
}

var min int = 9999999

func main() {
    input := strings.Split(lib.ReadFile(23), "\n")
    hallway := len(input[1]) - 2
    amphipods := []*amphipod{}
    for i, c := range input[2] {
        if c == '#' {
            continue
        }
        a := &amphipod{pos:coord{x:i-1,y:1}, typ:int(c)-65}
        amphipods = append(amphipods, a)
    }
    for i, c := range input[3] {
        if c == '#' || c == ' ' {
            continue
        }
        a := &amphipod{pos:coord{x:i-1,y:2}, typ:int(c)-65}
        amphipods = append(amphipods, a)
    }

    p1 := solveP1(amphipods, hallway)
    lib.WritePart1("%d", p1)

    min = 9999999
    p2 := solveP2(amphipods, hallway)
    lib.WritePart2("%d", p2)
}
