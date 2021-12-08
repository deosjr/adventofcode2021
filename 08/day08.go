package main

import (
    "fmt"
    "os"
    "os/exec"
    "strings"
    "github.com/deosjr/adventofcode2021/lib"
)

type line struct {
    signals []string
    output []string
}

func main() {
    var input []line
    lib.ReadFileByLine(8, func(s string) {
        ss := strings.Split(s, " | ")
        sig := strings.Split(ss[0], " ")
        out := strings.Split(ss[1], " ")
        input = append(input, line{sig, out})
    })
    var sum int64
    for _, l := range input {
        for _, o := range l.output {
            n := len(o)
            if n == 2 || n == 3 || n == 4 || n == 7 {
                sum++
            }
        }
    }

    lib.WritePart1("%d", sum)

    // for part2 we escape to prolog
    prologlines := []string{}
    for _, testline := range input {
        prologline := "line([A,B,C,D,E,F,G], [["
        signals := make([]string, len(testline.signals))
        for i, s := range testline.signals {
            signals[i] = strings.Join(strings.Split(strings.ToUpper(s), ""), ",")
        }
        prologline += strings.Join(signals, "],[")
        prologline += "]],[["
        outputs := make([]string, len(testline.output))
        for i, s := range testline.output {
            outputs[i] = strings.Join(strings.Split(strings.ToUpper(s), ""), ",")
        }
        prologline += strings.Join(outputs, "],[")
        prologline += "]])."
        prologlines = append(prologlines, prologline)
    }

    pl := fmt.Sprintf(`%%%%%%
:- use_module(library(clpfd)).

run :-
    findall(Vars-Input-Output, line(Vars, Input, Output), Lines),
    maplist(match_wires, Lines, Wires),
    sum(Wires, #=, Ans),
    format("~w", [Ans]).

match_wires(Vars-Input-Output, Ans) :-
    length(Vars, 7),
    Vars ins 0..6,
    all_distinct(Vars),
    line(Vars, Input, Output),
    partition([X]>>(length(X,2)), Input, [One], Rest1),
    match_digit(1, One),
    partition([X]>>(length(X,3)), Rest1, [Seven], Rest7),
    match_digit(7, Seven),
    partition([X]>>(length(X,4)), Rest7, [Four], Rest4),
    match_digit(4, Four),
    include([X]>>length(X,6), Rest4, Len6s),
    %% each number in the len6 list allows us to collapse one option:
    %% 0 gives us b(1) vs d(3), 6 gives f(5) vs c(2) and 9 gives g(6) vs e(4)
    maplist(unique_domain, Len6s),
    %% output
    maplist(sort, Output, Sorted),
    maplist([X,Y]>>digit(Y,X), Sorted, Digits),
    list2int(Digits, Ans).

match_digit(N, Vars) :-
    digit(N, List),
    list2domain(List, Domain),
    Vars ins Domain.

unique_domain(Vars) :-
    exclude([X]>>fd_size(X,1), Vars, AmbiguousVars),
    maplist([X,Y]>>fd_dom(X,Y), AmbiguousVars, Domains),
    include([X]>>(fd_dom(X,D), unique(D,Domains)), AmbiguousVars, [Unique]),
    fd_dom(Unique, Dom),
    lookup(Dom, Digit),
    Unique #= Digit.

unique(Elem, List) :-
    include([X]>>(X=Elem), List, [_]).

lookup(1\/3, 1).
lookup(2\/5, 5).
lookup(4\/6, 6).

list2int(List, N) :-
    foldl([X,Y,Z]>>(Z#=Y*10+X), List, 0, N).

list2domain([H], H).
list2domain([H|T], Dom) :-
    list2domain(T, D),
    Dom = D\/H.

digit(0, [0,1,2,4,5,6]).
digit(1, [2,5]).
digit(2, [0,2,3,4,6]).
digit(3, [0,2,3,5,6]).
digit(4, [1,2,3,5]).
digit(5, [0,1,3,5,6]).
digit(6, [0,1,3,4,5,6]).
digit(7, [0,2,5]).
digit(8, [0,1,2,3,4,5,6]).
digit(9, [0,1,2,3,5,6]).

%s`, strings.Join(prologlines, "\n"))

    err := os.WriteFile("temp.pl", []byte(pl), 0644)
    if err != nil {
        panic(err)
    }

    rawout, err := exec.Command("swipl","-q","-l","temp.pl","-t","run").Output()
	if err != nil {
		panic(err)
	}

    err = os.Remove("temp.pl")
    if err != nil {
        panic(err)
    }

    out := lib.MustParseInt(string(rawout))
    lib.WritePart2("%d", out)
}
