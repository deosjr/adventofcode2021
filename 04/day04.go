package main

import (
    "strconv"
    "strings"

    "github.com/deosjr/adventofcode2021/lib"
)

type board struct {
    nums []int64
    marked map[int]struct{}
    won bool
    score int64
}

func (b *board) mark(n int64) {
    for i, v := range b.nums {
        if v == n {
            b.marked[i] = struct{}{}
            break
        }
    }
}

func (b *board) wins() bool {
    rowsum := [5]int{}
    colsum := [5]int{}
    for k,_ := range b.marked {
        rowsum[k/5] += 1
        colsum[k%5] += 1
    }
    for _, v := range rowsum {
        if v == 5 {
            b.won = true
            return true
        }
    }
    for _, v := range colsum {
        if v == 5 {
            b.won = true
            return true
        }
    }
    return false
}

func (b *board) getScore(called int64) int64 {
    var sum int64
    for i, n := range b.nums {
        if _, ok := b.marked[i]; ok {
            continue
        }
        sum += n
    }
    return sum * called
}

func main() {
    split := strings.Split(lib.ReadFile(4), "\n")
    nums := []int64{}
    for _, s := range strings.Split(split[0], ",") {
        n, _ := strconv.ParseInt(s, 10, 64)
        nums = append(nums, n)
    }
    boards := []*board{}
    i := 0
    for _, s := range split[2:] {
        if s == "" {
            i++
            continue
        }
        if len(boards) == i {
            boards = append(boards, &board{
                nums:[]int64{},
                marked: map[int]struct{}{},
            })
        }
        for _, sn := range strings.Split(s, " ") {
            if sn == "" { continue }
            n, _ := strconv.ParseInt(sn, 10, 64)
            boards[i].nums = append(boards[i].nums, n)
        }
    }

    var p1 int64
    var p2 int64
    i = 0
Loop:
    for _, v := range nums {
        for _,b := range boards {
            b.mark(v)
            if !b.won && b.wins() {
                b.score = b.getScore(v)
                if i == 0 {
                    p1 = b.score
                }
                i++
            }
            if i == len(boards) {
                p2 = b.score
                break Loop
            }
        }
    }

    lib.WritePart1("%d", p1)
    lib.WritePart2("%d", p2)
}
