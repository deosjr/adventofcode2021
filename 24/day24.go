package main

import (
    "fmt"
    "strconv"
    "strings"
    "github.com/deosjr/adventofcode2021/lib"
)

type instr struct {
    op string
    a  int
    b  int
    value int64
}

func (i instr) hasValue() bool {
    return i.b == -1
}

func (i instr) String() string {
    a := rune(i.a + 119)
    b := string(rune(i.b + 119))
    if i.hasValue() {
        b = fmt.Sprintf("%d", i.value)
    }
    switch i.op {
    case "inp":
        return fmt.Sprintf("input %c", a)
    case "add":
        return fmt.Sprintf("%c += %s", a, b)
    case "mul":
        return fmt.Sprintf("%c *= %s", a, b)
    case "div":
        return fmt.Sprintf("%c /= %s", a, b)
    case "mod":
        return fmt.Sprintf("%c = %c mod %s", a, a, b)
    case "eql":
        return fmt.Sprintf("%c = %c eq %s", a, a, b)
    case "set":
        return fmt.Sprintf("%c = %s", a, b)
    }
    return ""
}

func main() {
    var input []instr
    lib.ReadFileByLine(24, func(line string) {
        s := strings.Split(line, " ")
        a := int(s[1][0]) - 119
        in := instr{op:s[0], a:a, b:-1}
        if in.op == "inp" {
            input = append(input, in)
            return
        }
        n, err := strconv.ParseInt(s[2], 10, 64)
        if err == nil {
            in.value = n
        } else {
            in.b = int(s[2][0]) - 119
        }
        input = append(input, in)
    })

    min := [4]int64{}
    max := [4]int64{}
    prev := input[0]
    var prevEqNever, eqNever, eqAlways bool

    effective := []instr{}

    for _, in := range input {
        prevmin, prevmax := min, max
        minv, maxv := in.value, in.value
        if !in.hasValue() {
            minv = min[in.b]
            maxv = max[in.b]
        }
        switch in.op {
        case "inp":
            min[in.a] = 1
            max[in.a] = 9
        case "add":
            min[in.a] = min[in.a] + minv
            max[in.a] = max[in.a] + maxv
        case "mul":
            min[in.a] = min[in.a] * minv
            max[in.a] = max[in.a] * maxv
        case "div":
            min[in.a] = min[in.a] / minv
            max[in.a] = max[in.a] / maxv
        case "mod":
            min[in.a] = min[in.a] % minv
            max[in.a] = max[in.a] % maxv
        case "eql":
            if max[in.a] < minv || min[in.a] > maxv {
                min[in.a] = 0
                max[in.a] = 0
                eqNever = true
            } else if min[in.a] == max[in.a] && minv == maxv && minv == min[in.a] {
                min[in.a] = 1
                max[in.a] = 1
                eqAlways = true
            } else {
                min[in.a] = 0
                max[in.a] = 1
            }
        }
        if prevmin == min && prevmax == max && in.op != "inp" && in.op != "eql" {
            continue
        }
        if prev.op == in.op && in.op == "add" && prev.a == in.a && prev.hasValue() && in.hasValue() {
            effective[len(effective)-1] = instr{op:"add", a:prev.a, b:-1, value: prev.value+in.value}
            continue
        }
        if prevEqNever && eqAlways && prev.a == in.a {
            for i:=len(effective)-1; i>=0; i-- {
                p := effective[i]
                if p.a != in.a {
                    break
                }
                effective = effective[:i]
            }
            effective = append(effective, instr{op:"set", a:in.a, b:-1, value: 1})
            prevEqNever = eqNever
            eqNever, eqAlways = false, false
            continue
        }
        if in.op == "mul" && in.hasValue() && in.value == 0 {
            in = instr{op:"set", a:in.a, b:-1, value: 0}
            for i:=len(effective)-1; i>=0; i-- {
                p := effective[i]
                if p.a != in.a {
                    break
                }
                effective = effective[:i]
            }
        }
        prev = in
        prevEqNever = eqNever
        eqNever, eqAlways = false, false
        effective = append(effective, in)
    }

    fmt.Println("func(V) :-")
    fmt.Println("\tV = [A,B,C,D,E,F,G,H,I,J,K,L,M,N],")
    fmt.Println("\t[W0, X0, Y0, Z0] ins 0..0,")
    uses := [4]int{}
    for _, in := range effective {
        runeA := rune(in.a+87)
        b := fmt.Sprintf("%d", in.value)
        if !in.hasValue() {
            b = fmt.Sprintf("%c%d", rune(in.b+87), uses[in.b])
        }
        switch in.op {
        case "inp":
            fmt.Printf("\tW%d #= %c,\n", uses[in.a]+1, rune(uses[in.a]+65))
            uses[in.a] = uses[in.a] + 1
        case "add":
            fmt.Printf("\t%c%d #= %c%d + %s,\n", runeA, uses[in.a]+1, runeA, uses[in.a], b)
            uses[in.a] = uses[in.a] + 1
        case "mul":
            fmt.Printf("\t%c%d #= %c%d * %s,\n", runeA, uses[in.a]+1, runeA, uses[in.a], b)
            uses[in.a] = uses[in.a] + 1
        case "div":
            fmt.Printf("\t%c%d #= %c%d div %s,\n", runeA, uses[in.a]+1, runeA, uses[in.a], b)
            uses[in.a] = uses[in.a] + 1
        case "mod":
            fmt.Printf("\t%c%d #= %c%d mod %s,\n", runeA, uses[in.a]+1, runeA, uses[in.a], b)
            uses[in.a] = uses[in.a] + 1
        case "eql":
            next := fmt.Sprintf("%c%d", runeA, uses[in.a]+1)
            fmt.Printf("\t( %c%d #= %s -> %s #= 1 ; %s #= 0 ),\n", runeA, uses[in.a], b, next, next)
            uses[in.a] = uses[in.a] + 1
        case "set":
            fmt.Printf("\t%c%d #= %s,\n", runeA, uses[in.a]+1, b)
            uses[in.a] = uses[in.a] + 1
        }
    }
    fmt.Printf("\tZ%d #= 0.\n", uses[3])

    // max value when generating all possibilities using prolog
    lib.WritePart1("%d", 39999698799429)

    // min value when generating all possibilities using prolog
    lib.WritePart2("%d", 18116121134117)
}
