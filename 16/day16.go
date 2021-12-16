package main

import (
    "fmt"
    "math/big"
    "strconv"
    "github.com/deosjr/adventofcode2021/lib"
)

type packet struct {
    version int64
    ptype int64
    lit int64
    sub []*packet
}

func (p *packet) eval() *big.Int {
    switch p.ptype {
    case 0:
        sum := big.NewInt(0)
        for _, sp := range p.sub {
            sum.Add(sum, sp.eval())
        }
        return sum
    case 1:
        prod := big.NewInt(1)
        for _, sp := range p.sub {
            prod.Mul(prod, sp.eval())
        }
        return prod
    case 2:
        min := p.sub[0].eval()
        if len(p.sub) == 1 {
            return min
        }
        for _, sp := range p.sub[1:] {
            v := sp.eval()
            if v.Cmp(min) == -1 {
                min = v
            }
        }
        return min
    case 3:
        max := p.sub[0].eval()
        if len(p.sub) == 1 {
            return max
        }
        for _, sp := range p.sub[1:] {
            v := sp.eval()
            if v.Cmp(max) == 1 {
                max = v
            }
        }
        return max
    case 4:
        return big.NewInt(p.lit)
    case 5:
        v1 := p.sub[0].eval()
        v2 := p.sub[1].eval()
        if v1.Cmp(v2) == 1 {
            return big.NewInt(1)
        }
        return big.NewInt(0)
    case 6:
        v1 := p.sub[0].eval()
        v2 := p.sub[1].eval()
        if v1.Cmp(v2) == -1 {
            return big.NewInt(1)
        }
        return big.NewInt(0)
    case 7:
        v1 := p.sub[0].eval()
        v2 := p.sub[1].eval()
        if v1.Cmp(v2) == 0 {
            return big.NewInt(1)
        }
        return big.NewInt(0)
    }
    panic(p.ptype)
}

var p1 int64 = 0

func parsePacket(s string) (*packet, int) {
    p := &packet{}
    p.version, _ = strconv.ParseInt(s[0:3], 2, 64)
    p1 += p.version
    p.ptype, _ = strconv.ParseInt(s[3:6], 2, 64)
    n := 6
    var pn int
    if p.ptype == 4 {
        p.lit, pn = parseLiteral(s[6:])
        return p, n + pn
    }
    lengthType0 := s[6] == '0'
    n += 1
    if lengthType0 {
        length, _ := strconv.ParseInt(s[7:22], 2, 64)
        p.sub, pn = parseLenSubpackets(s[22:], length)
        n += pn + 15
    } else {
        num, _ := strconv.ParseInt(s[7:18], 2, 64)
        p.sub, pn = parseNumSubpackets(s[18:], num)
        n += pn + 11
    }
    return p, n
}

func parseLiteral(s string) (lit int64, i int) {
    for {
        next := s[i:i+5]
        x, _ := strconv.ParseInt(next[1:], 2, 64)
        lit = lit * 16 + x
        i += 5
        if next[0] == '0' {
            break
        }
    }
    return lit, i
}

func parseLenSubpackets(s string, length int64) (ps []*packet, n int) {
    offset := 0
    for offset < int(length) {
        p, plen := parsePacket(s[offset:])
        ps = append(ps, p)
        offset += plen
    }
    return ps, offset
}

func parseNumSubpackets(s string, num int64) (ps []*packet, n int) {
    offset := 0
    for i:=0; i<int(num); i++ {
        p, plen := parsePacket(s[offset:])
        ps = append(ps, p)
        offset += plen
    }
    return ps, offset
}

func hex2bin(hex string) string {
    bin := ""
    for _, r := range hex {
        dec, _ := strconv.ParseInt(string(r), 16, 64)
        bin += fmt.Sprintf("%04b", dec)
    }
    return bin
}

func main() {
    hex := lib.ReadFile(16)
    bin := hex2bin(hex)

    hierarchy, _ := parsePacket(bin)

    lib.WritePart1("%d", p1)

    lib.WritePart2("%d", hierarchy.eval())
}
