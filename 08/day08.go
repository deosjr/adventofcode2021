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

    sum = 0
    for _, testline := range input {

    // for part2 we escape to prolog
    prologline := "line([["
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

    pl := fmt.Sprintf(`%%%%%%
:- use_module(library(clpfd)).

run :-
    line(Input, Output),
    partition([X]>>(length(X,2)), Input, [One], Rest1),
    match_digit(1, One),
    partition([X]>>(length(X,3)), Rest1, [Seven], Rest7),
    match_digit(7, Seven),
    partition([X]>>(length(X,4)), Rest7, [Four], Rest4),
    match_digit(4, Four),
    %%%% 9 should give us 'g', only signal of len 6 with only one unbound variable now
    include([X]>>length(X,6), Rest4, Len6s),
    partition(numvars(1), Len6s, [Nine], _Rest9),
    match_digit(9, Nine),
    %%%% similarly 2 should now give us 'e'
    include([X]>>length(X,5), Rest4, Len5s),
    partition(numvars(1), Len5s, [Two], _Rest2),
    match_digit(2, Two),
    %% sanity check 8 but we should have everyting now
    partition([X]>>(length(X,7)), Rest4, [Eight], _Rest8),
    match_digit(8, Eight),
    maplist(sort, Output, Sorted),
    maplist([X,Y]>>digit(Y,X), Sorted, Digits),
    list2int(Digits, Ans),
    format("~w", [Ans]).

match_digit(N, Vars) :-
    permutation(Vars, Perm),
    digit(N, Perm).

numvars(0, []).
numvars(N, [H|T]) :-
    numvars(X, T),
    ( var(H) -> N#=X+1 ; N#=X ).

list2int(List, N) :-
    list2int(List, 0, N).
list2int([], X, X).
list2int([H|T], X, Y) :-
    N #= X*10+H,
    list2int(T, N, Y).

digit(0, [a,b,c,e,f,g]).
digit(1, [c,f]).
digit(2, [a,c,d,e,g]).
digit(3, [a,c,d,f,g]).
digit(4, [b,c,d,f]).
digit(5, [a,b,d,f,g]).
digit(6, [a,b,d,e,f,g]).
digit(7, [a,c,f]).
digit(8, [a,b,c,d,e,f,g]).
digit(9, [a,b,c,d,f,g]).
%s`, prologline)

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
    sum += out
    }

    lib.WritePart2("%d", sum)
}
