package main

import "testing"

func TestExplode(t *testing.T) {
    for _, tt := range []struct{
        input string
        want string
    }{
        {
            input: "[[[[[9,8],1],2],3],4]",
            want:  "[[[[0,9],2],3],4]",
        },
        {
            input: "[7,[6,[5,[4,[3,2]]]]]",
            want:  "[7,[6,[5,[7,0]]]]",
        },
        {
            input: "[[6,[5,[4,[3,2]]]],1]",
            want:  "[[6,[5,[7,0]]],3]",
        },
        {
            input: "[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]",
            want:  "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
        },
        {
            input: "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
            want:  "[[3,[2,[8,0]]],[9,[5,[7,0]]]]",
        },
    }{
        n, _ := parseInternal(tt.input)
        n.explode()
        if n.String() != tt.want {
            t.Errorf("got %s but want %s", n.String(), tt.want)
        }
    }
}

func TestSplit(t *testing.T) {
    n1, _ := parse("[[[[4,3],4],4],[7,[[8,4],9]]]")
    n2, _ := parse("[1,1]")
    n := add(n1, n2)
    if n.String() != "[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]" {
        t.Fatalf("wrong after addition: %s", n)
    }
    n.explode()
    if n.String() != "[[[[0,7],4],[7,[[8,4],9]]],[1,1]]" {
        t.Fatalf("wrong after explode1: %s", n)
    }
    n.explode()
    if n.String() != "[[[[0,7],4],[15,[0,13]]],[1,1]]" {
        t.Fatalf("wrong after explode2: %s", n)
    }
    n.split()
    if n.String() != "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]" {
        t.Fatalf("wrong after split1: %s", n)
    }
    n.split()
    if n.String() != "[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]" {
        t.Fatalf("wrong after split2: %s", n)
    }
    n.explode()
    if n.String() != "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]" {
        t.Fatalf("wrong after explode3: %s", n)
    }
    n1, _ = parse("[[[[4,3],4],4],[7,[[8,4],9]]]")
    n2, _ = parse("[1,1]")
    n = add(n1, n2)
    n.reduce()
    if n.String() != "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]" {
        t.Fatalf("wrong after reduce: %s", n)
    }
}

func TestMagnitude(t *testing.T) {
    for _, tt := range []struct{
        input string
        want int64
    }{
        {
            input: "[[1,2],[[3,4],5]]",
            want:  143,
        },
        {
            input: "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
            want:  1384,
        },
        {
            input: "[[[[1,1],[2,2]],[3,3]],[4,4]]",
            want:  445,
        },
        {
            input: "[[[[3,0],[5,3]],[4,4]],[5,5]]",
            want:  791,
        },
        {
            input: "[[[[5,0],[7,4]],[5,5]],[6,6]]",
            want:  1137,
        },
        {
            input: "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
            want:  3488,
        },
    }{
        n, _ := parseInternal(tt.input)
        got := n.magnitude()
        if got != tt.want {
            t.Errorf("got %d but want %d", got, tt.want)
        }
    }
}
