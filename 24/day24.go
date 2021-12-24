package main

import (
    "fmt"
    "strconv"
    "strings"
    "github.com/deosjr/adventofcode2021/lib"
)

// NOTES:
// inp only ever takes w and w is never otherwise updated
// x and y are local vars per input
// z is kept between inputs
// z can only be 0 if none of the if statements pass (?)

// written by hand after looking at effective lines
func validate(n int64) bool {
    w := n % 10
    n = n / 10
    z := w+12
    w = n % 10
    n = n / 10
    z *= 26
    z += w+9
    w = n % 10
    n = n / 10
    z *= 26
    z += w+8
    w = n % 10
    n = n / 10
    x := (z % 26) - 8
    z = z / 26
    if x != w {
        z *= 26
        z += w+3
    }
    w = n % 10
    n = n / 10
    z *= 26
    z += w
    w = n % 10
    n = n / 10
    z *= 26
    z += w+11
    w = n % 10
    n = n / 10
    z *= 26
    z += w+10
    w = n % 10
    n = n / 10
    x = (z % 26) - 11
    z = z / 26
    if x != w {
        z *= 26
        z += w+13
    }
    w = n % 10
    n = n / 10
    z *= 26
    z += w+3
    w = n % 10
    n = n / 10
    x = (z % 26) -1
    z = z / 26
    if x != w {
        z *= 26
        z += w+10
    }
    w = n % 10
    n = n / 10
    x = (z % 26) - 8
    z = z / 26
    if x != w {
        z *= 26
        z += w+10
    }
    w = n % 10
    n = n / 10
    x = (z % 26) - 5
    z = z / 26
    if x != w {
        z *= 26
        z += w+14
    }
    w = n % 10
    n = n / 10
    x = (z % 26) - 16
    z = z / 26
    if x != w {
        z *= 26
        z += w+6
    }
    w = n % 10
    n = n / 10
    x = (z % 26) - 6
    z = z / 26
    if x != w {
        z *= 26
        z += w+5
    }
    return z == 0
}

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

    //fmt.Println(validate(92499789699993)) equals true
    // evaluated by hand
    lib.WritePart1("%d", 39999698799429)

    //fmt.Println(validate(71143112161181)) equals true
    // evaluated by hand
    lib.WritePart2("%d", 18116121134117)
}
