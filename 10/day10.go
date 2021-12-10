
package main

import (
    "sort"
    "github.com/deosjr/adventofcode2021/lib"
)

type stack []rune

func (s *stack) push(r rune) {
    *s = append(*s, r)
}

func (s *stack) pop() rune {
    t := *s
    news, r := t[:len(t)-1], t[len(t)-1]
    *s = news
    return r
}

func main() {
    var input []string
    lib.ReadFileByLine(10, func(s string) {
        input = append(input, s)
    })

    var p1 int64
    p2 := []int{}
Loop:
    for _, line := range input {
        s := &stack{}
        for _, c := range line {
            switch c {
            case '(','[','{','<':
                s.push(c)
            case ')':
                last := s.pop()
                if last != '(' {
                    p1 += 3
                    continue Loop
                }
            case ']':
                last := s.pop()
                if last != '[' {
                    p1 += 57
                    continue Loop
                }
            case '}':
                last := s.pop()
                if last != '{' {
                    p1 += 1197
                    continue Loop
                }
            case '>':
                last := s.pop()
                if last != '<' {
                    p1 += 25137
                    continue Loop
                }
            }
        }
        var sum int
        for len(*s) != 0 {
            sum *= 5
            switch s.pop() {
            case '(':
                sum += 1
            case '[':
                sum += 2
            case '{':
                sum += 3
            case '<':
                sum += 4
            }
        }
        p2 = append(p2, sum)
    }

    lib.WritePart1("%d", p1)

    sort.Ints(p2)
    lib.WritePart2("%d", p2[len(p2)/2])
}
