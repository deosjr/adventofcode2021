package main

import (
    "fmt"
    "os"
    "os/exec"
    "strings"
    "github.com/deosjr/adventofcode2021/lib"
)

func main() {
    prologlines := []string{}
    lib.ReadFileByLine(12, func(s string) {
        split := strings.Split(s, "-")
        from, to := split[0], split[1]
        prologlines = append(prologlines, fmt.Sprintf("cave(%q,%q).", from, to))
        prologlines = append(prologlines, fmt.Sprintf("cave(%q,%q).", to, from))
    })

    pl := fmt.Sprintf(`%%%%%%
run :-
    findall(X, path(["start"], "end", X), Paths),
    length(Paths, Ans),
    write(Ans).

small_cave(S) :-
    string_code(1, S, C),
    char_type(C, lower).

path(From, To, Path) :-
    From = [H|T],
    cave(H, Next),
    (
        Next = To
    ->
        reverse([To|From], Path)
    ;
        (
            small_cave(Next)
        ->
            not(member(Next, T))
        ;
            true
        ),
        path([Next|From], To, Path)
    ).
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
    lib.WritePart1("%d", out)

    pl = fmt.Sprintf(`%%%%%%
run :-
    findall(X, path(["start"], "end", false, X), Paths),
    length(Paths, Ans),
    write(Ans).

small_cave(S) :-
    string_code(1, S, C),
    char_type(C, lower).

path(From, To, Doubled, Path) :-
    From = [H|T],
    cave(H, Next),
    Next \= "start",
    (
        Next = To
    ->
        reverse([To|From], Path)
    ;
        (
            small_cave(Next)
        ->
            (
                Doubled = true
            ->
                not(member(Next, T)),
                NewDoubled = true
            ;
                (
                    member(Next, T)
                ->
                    NewDoubled = true
                ;
                    NewDoubled = false
                )
            )
        ;
            NewDoubled = Doubled,
            true
        ),
        path([Next|From], To, NewDoubled, Path)
    ).
%s`, strings.Join(prologlines, "\n"))

    err = os.WriteFile("temp.pl", []byte(pl), 0644)
    if err != nil {
        panic(err)
    }

    rawout, err = exec.Command("swipl","-q","-l","temp.pl","-t","run").Output()
	if err != nil {
		panic(err)
	}

    err = os.Remove("temp.pl")
    if err != nil {
        panic(err)
    }

    out = lib.MustParseInt(string(rawout))
    lib.WritePart2("%d", out)
}
