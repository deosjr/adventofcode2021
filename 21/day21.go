package main

import (
    "fmt"
    "strings"
    "github.com/deosjr/adventofcode2021/lib"
)

func modPlus1(n, m int) int {
    x := n % m
    if x == 0 {
        x = m
    }
    return x
}

type universe struct {
    p1pos int
    p1score int
    p2pos int
    p2score int
    p2turn bool // if false, its p1's turn
    num int64
}

func main() {
    input := strings.Split(lib.ReadFile(21), "\n")
    var p1, p2 int
    fmt.Sscanf(input[0], "Player 1 starting position: %d", &p1)
    fmt.Sscanf(input[1], "Player 2 starting position: %d", &p2)

    state := universe{p1pos: p1, p2pos: p2}
    var die int
    dierolls := 0
    for state.p1score < 1000 && state.p2score < 1000 {
        die = modPlus1(die+1, 100)
        dierolls++
        sum := die
        die = modPlus1(die+1, 100)
        dierolls++
        sum += die
        die = modPlus1(die+1, 100)
        dierolls++
        sum += die
        if state.p2turn {
            state.p2pos = modPlus1(state.p2pos + sum, 10)
            state.p2score += state.p2pos
            state.p2turn = false
            continue
        }
        state.p1pos = modPlus1(state.p1pos + sum, 10)
        state.p1score += state.p1pos
        state.p2turn = true
    }
    ans1 := dierolls*state.p1score
    if state.p2score < state.p1score {
        ans1 = dierolls*state.p2score
    }
    lib.WritePart1("%d", ans1)

    universes := []universe{{p1pos: p1, p2pos: p2, num: 1}}
    diedistr := []struct{
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
    var p1sum, p2sum int64
    for len(universes) > 0 {
        u := universes[0]
        universes = universes[1:]
        if u.p1score >= 21 {
            p1sum += u.num
            continue
        }
        if u.p2score >= 21 {
            p2sum += u.num
            continue
        }
        for _, dist := range diedistr {
            if u.p2turn {
                newpos := modPlus1(u.p2pos+dist.sum, 10)
                newu := universe{
                    p1pos:   u.p1pos,
                    p1score: u.p1score,
                    p2pos:   newpos,
                    p2score: u.p2score + newpos,
                    p2turn:  false,
                    num:     u.num * dist.num,
                }
                universes = append(universes, newu)
                continue
            }
            newpos := modPlus1(u.p1pos+dist.sum, 10)
            newu := universe{
                p1pos:   newpos,
                p1score: u.p1score + newpos,
                p2pos:   u.p2pos,
                p2score: u.p2score,
                p2turn:  true,
                num:     u.num * dist.num,
            }
            universes = append(universes, newu)
        }
    }
    ans2 := p1sum
    if p2sum > p1sum {
        ans2 = p2sum
    }
    lib.WritePart2("%d", ans2)
}
