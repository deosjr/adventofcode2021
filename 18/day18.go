
package main

import (
    "fmt"
    "github.com/deosjr/adventofcode2021/lib"
)

type node struct {
    isLeft bool // otherwise isRight
    value int64
    parent *node
    left *node
    right *node
}

func (n *node) isLeaf() bool {
    return n.left == nil && n.right == nil
}

func (n *node) isRoot() bool {
    return n.parent == nil
}

func (n *node) reduce() {
    for {
        if n.explode() {
            continue
        }
        if n.split() {
            continue
        }
        break
    }
}

func (n *node) find(depth int, match func(*node, int) bool) (*node, bool) {
    if match(n, depth) {
        return n, true
    }
    if n.isLeaf() {
        return nil, false
    }
    if v, ok := n.left.find(depth+1, match); ok {
        return v, true
    }
    if v, ok := n.right.find(depth+1, match); ok {
        return v, true
    }
    return nil, false
}

func (n *node) explode() bool {
    candidate, ok := n.find(0, func(v *node, depth int) bool {
        return !v.isLeaf() && depth == 4
    })
    if !ok {
        return false
    }
    if v, ok := candidate.left.leftLeaf(); ok {
        v.value += candidate.left.value
    }
    if v, ok := candidate.right.rightLeaf(); ok {
        v.value += candidate.right.value
    }
    newnode := &node{parent: candidate.parent, isLeft: candidate.isLeft}
    if candidate.isLeft {
        candidate.parent.left = newnode
    } else {
        candidate.parent.right = newnode
    }
    return true
}

func (n *node) split() bool {
    candidate, ok := n.find(0, func(v *node, _ int) bool {
        return v.isLeaf() && v.value > 9
    })
    if !ok {
        return false
    }
    v := candidate.value / 2
    newnode := &node{parent: candidate.parent, isLeft: candidate.isLeft}
    newnode.left = &node{parent: newnode, value: v, isLeft: true}
    newnode.right = &node{parent: newnode, value: v}
    if candidate.value % 2 == 1 {
        newnode.right.value += 1
    }
    if candidate.isLeft {
        candidate.parent.left = newnode
    } else {
        candidate.parent.right = newnode
    }
    return true
}

func (n *node) leftLeaf() (*node, bool) {
    nn := n
    for !nn.isRoot() && nn.isLeft {
        nn = nn.parent
    }
    if nn.isRoot() {
        return nil, false
    }
    nn = nn.parent.left
    for !nn.isLeaf() {
        nn = nn.right
    }
    return nn, true
}

func (n *node) rightLeaf() (*node, bool) {
    nn := n
    for !nn.isRoot() && !nn.isLeft {
        nn = nn.parent
    }
    if nn.isRoot() {
        return nil, false
    }
    nn = nn.parent.right
    for !nn.isLeaf() {
        nn = nn.left
    }
    return nn, true
}

func add(left, right *node) *node {
    n := &node{left:left, right:right}
    n.left.isLeft = true
    n.left.parent = n
    n.right.parent = n
    return n
}

func (n *node) magnitude() int64 {
    if n.isLeaf() {
        return n.value
    }
    return 3*n.left.magnitude() + 2*n.right.magnitude()
}

func parseToken(s string, token rune) string {
    if rune(s[0]) != token {
        panic("invalid parse")
    }
    return s[1:]
}

func parseInternal(s string) (*node, string) {
    n := &node{}
    s = parseToken(s, '[')
    n.left, s = parse(s)
    n.left.isLeft = true
    n.left.parent = n
    s = parseToken(s, ',')
    n.right, s = parse(s)
    n.right.parent = n
    s = parseToken(s, ']')
    return n, s
}

func parse(s string) (*node, string) {
    if s[0] == '[' {
        return parseInternal(s)
    }
    v := lib.MustParseInt(string(s[0]))
    return &node{value: v}, s[1:]
}

func (n *node) String() string {
    if n.isLeaf() {
        return fmt.Sprint(n.value)
    }
    return fmt.Sprintf("[%s,%s]", n.left.String(), n.right.String())
}

func main() {
    var input []string
    lib.ReadFileByLine(18, func(s string) {
        input = append(input, s)
    })

    p1, _ := parse(input[0])
    for _, s := range input[1:] {
        n, _ := parse(s)
        p1 = add(p1, n)
        p1.reduce()
    }

    lib.WritePart1("%d", p1.magnitude())

    var p2 int64
    for i:=0; i<len(input); i++ {
        for j:=0; j<len(input); j++ {
            if i == j { continue }
            a, _ := parse(input[i])
            b, _ := parse(input[j])
            n := add(a, b)
            n.reduce()
            m := n.magnitude()
            if m > p2 {
                p2 = m
            }
        }
    }
    lib.WritePart2("%d", p2)
}
