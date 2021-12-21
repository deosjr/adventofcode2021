package main

import (
    "fmt"
    "strings"
    "github.com/deosjr/adventofcode2021/lib"
)

var mem = map[int]int{}

func mod10Plus1(n int) int {
    if v, ok := mem[n]; ok {
        return v
    }
    x := ((n-1) % 10) + 1
    mem[n] = x
    return x
}

type universe struct {
    activepos int
    activescore int
    passivepos int
    passivescore int
    p2turn bool // if false, its p1's turn
}

func part1(u universe) int {
    var die int
    dierolls := 0
    for u.passivescore < 1000 {
        sum := 0
        for i:=0; i<3; i++ {
            die = (die % 100) + 1
            dierolls++
            sum += die
        }
        newpos := mod10Plus1(u.activepos + sum)
        u.activepos, u.passivepos = u.passivepos, newpos
        u.activescore, u.passivescore = u.passivescore, u.activescore + u.passivepos
        u.p2turn = !u.p2turn
    }
    return dierolls * u.activescore
}

var diedistr = []struct{
    sum int
    num int64
}{
    {sum:3, num:1},
    {sum:4, num:3},
    {sum:5, num:6},
    {sum:6, num:7},
    {sum:7, num:6},
    {sum:8, num:3},
    {sum:9, num:1},
}

type in struct {
    u universe
    n int64
}

type out struct {
    p1 int64
    p2 int64
}

var cache = map[in]out{}

func part2(u universe, n int64) (int64, int64) {
    if u.passivescore >= 21 {
        if u.p2turn {
            return n, 0
        } else {
            return 0, n
        }
    }
    v, ok := cache[in{u, n}]
    if ok {
        return v.p1, v.p2
    }
    var p1sum, p2sum int64
    newu := universe{
        activepos: u.passivepos,
        activescore: u.passivescore,
    }
    for _, d := range diedistr {
        newu.passivepos = mod10Plus1(u.activepos + d.sum)
        newu.passivescore = u.activescore + newu.passivepos
        newu.p2turn = !u.p2turn
        p1, p2 := part2(newu, d.num)
        p1sum += p1
        p2sum += p2
    }
    cache[in{u, n}] = out{n*p1sum, n*p2sum}
    return n*p1sum, n*p2sum
}

func main() {
    input := strings.Split(lib.ReadFile(21), "\n")
    var p1, p2 int
    fmt.Sscanf(input[0], "Player 1 starting position: %d", &p1)
    fmt.Sscanf(input[1], "Player 2 starting position: %d", &p2)

    ans1 := part1(universe{activepos: p1, passivepos: p2})
    lib.WritePart1("%d", ans1)

    p1sum, p2sum := part2(universe{activepos: p1, passivepos: p2}, 1)
    ans2 := p1sum
    if p2sum > p1sum {
        ans2 = p2sum
    }
    lib.WritePart2("%d", ans2)
}
