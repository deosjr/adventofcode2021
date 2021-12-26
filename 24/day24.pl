:- use_module(library(clpfd)).

func(V) :-
    V = [A,B,C,D,E,F,G,H,I,J,K,L,M,N],
    N #= A + 6,
    M #= B - 7,
    L #= E - 5,
    K #= F + 3,
    J #= I + 2,
    H #= G - 1,
    C #= D.

find(V) :-
    length(V, 14),
    V ins 1..9,
    func(V),
    label(V).

run2 :-
    findall(V, find(V), List),
    length(List, X),
    writeln(X).

run :-
    p1,
    p2.

p1 :-
    length(V, 14),
    V ins 1..9,
    func2(V),
    V = [A,B,C,D,E,F,G,H,I,J,K,L,M,N],
    labeling([up], V),
    format("Part 1: ~w\n", [V]).

p2 :-
    length(V, 14),
    V ins 1..9,
    func2(V),
    V = [A,B,C,D,E,F,G,H,I,J,K,L,M,N],
    labeling([down], V),
    format("Part 2: ~w\n", [V]).
    
func2(V) :-
	V = [A,B,C,D,E,F,G,H,I,J,K,L,M,N],
	[W0, X0, Y0, Z0] ins 0..0,
	W1 #= A,
	X1 #= 1,
	Y1 #= 0,
	Y2 #= Y1 + W1,
	Y3 #= Y2 + 12,
	Z1 #= Z0 + Y3,
	W2 #= B,
	X2 #= 1,
	Y4 #= 0,
	Y5 #= Y4 + 26,
	Z2 #= Z1 * Y5,
	Y6 #= 0,
	Y7 #= Y6 + W2,
	Y8 #= Y7 + 9,
	Z3 #= Z2 + Y8,
	W3 #= C,
	X3 #= 1,
	Y9 #= 0,
	Y10 #= Y9 + 26,
	Z4 #= Z3 * Y10,
	Y11 #= 0,
	Y12 #= Y11 + W3,
	Y13 #= Y12 + 8,
	Z5 #= Z4 + Y13,
	W4 #= D,
	X4 #= 0,
	X5 #= X4 + Z5,
	X6 #= X5 mod 26,
	Z6 #= Z5 div 26,
	X7 #= X6 + -8,
	( X7 #= W4 -> X8 #= 1 ; X8 #= 0 ),
	( X8 #= 0 -> X9 #= 1 ; X9 #= 0 ),
	Y14 #= 0,
	Y15 #= Y14 + 25,
	Y16 #= Y15 * X9,
	Y17 #= Y16 + 1,
	Z7 #= Z6 * Y17,
	Y18 #= 0,
	Y19 #= Y18 + W4,
	Y20 #= Y19 + 3,
	Y21 #= Y20 * X9,
	Z8 #= Z7 + Y21,
	W5 #= E,
	X10 #= 1,
	Y22 #= 0,
	Y23 #= Y22 + 26,
	Z9 #= Z8 * Y23,
	Y24 #= 0,
	Y25 #= Y24 + W5,
	Z10 #= Z9 + Y25,
	W6 #= F,
	X11 #= 1,
	Y26 #= 0,
	Y27 #= Y26 + 26,
	Z11 #= Z10 * Y27,
	Y28 #= 0,
	Y29 #= Y28 + W6,
	Y30 #= Y29 + 11,
	Z12 #= Z11 + Y30,
	W7 #= G,
	X12 #= 1,
	Y31 #= 0,
	Y32 #= Y31 + 26,
	Z13 #= Z12 * Y32,
	Y33 #= 0,
	Y34 #= Y33 + W7,
	Y35 #= Y34 + 10,
	Z14 #= Z13 + Y35,
	W8 #= H,
	X13 #= 0,
	X14 #= X13 + Z14,
	X15 #= X14 mod 26,
	Z15 #= Z14 div 26,
	X16 #= X15 + -11,
	( X16 #= W8 -> X17 #= 1 ; X17 #= 0 ),
	( X17 #= 0 -> X18 #= 1 ; X18 #= 0 ),
	Y36 #= 0,
	Y37 #= Y36 + 25,
	Y38 #= Y37 * X18,
	Y39 #= Y38 + 1,
	Z16 #= Z15 * Y39,
	Y40 #= 0,
	Y41 #= Y40 + W8,
	Y42 #= Y41 + 13,
	Y43 #= Y42 * X18,
	Z17 #= Z16 + Y43,
	W9 #= I,
	X19 #= 1,
	Y44 #= 0,
	Y45 #= Y44 + 26,
	Z18 #= Z17 * Y45,
	Y46 #= 0,
	Y47 #= Y46 + W9,
	Y48 #= Y47 + 3,
	Z19 #= Z18 + Y48,
	W10 #= J,
	X20 #= 0,
	X21 #= X20 + Z19,
	X22 #= X21 mod 26,
	Z20 #= Z19 div 26,
	X23 #= X22 + -1,
	( X23 #= W10 -> X24 #= 1 ; X24 #= 0 ),
	( X24 #= 0 -> X25 #= 1 ; X25 #= 0 ),
	Y49 #= 0,
	Y50 #= Y49 + 25,
	Y51 #= Y50 * X25,
	Y52 #= Y51 + 1,
	Z21 #= Z20 * Y52,
	Y53 #= 0,
	Y54 #= Y53 + W10,
	Y55 #= Y54 + 10,
	Y56 #= Y55 * X25,
	Z22 #= Z21 + Y56,
	W11 #= K,
	X26 #= 0,
	X27 #= X26 + Z22,
	X28 #= X27 mod 26,
	Z23 #= Z22 div 26,
	X29 #= X28 + -8,
	( X29 #= W11 -> X30 #= 1 ; X30 #= 0 ),
	( X30 #= 0 -> X31 #= 1 ; X31 #= 0 ),
	Y57 #= 0,
	Y58 #= Y57 + 25,
	Y59 #= Y58 * X31,
	Y60 #= Y59 + 1,
	Z24 #= Z23 * Y60,
	Y61 #= 0,
	Y62 #= Y61 + W11,
	Y63 #= Y62 + 10,
	Y64 #= Y63 * X31,
	Z25 #= Z24 + Y64,
	W12 #= L,
	X32 #= 0,
	X33 #= X32 + Z25,
	X34 #= X33 mod 26,
	Z26 #= Z25 div 26,
	X35 #= X34 + -5,
	( X35 #= W12 -> X36 #= 1 ; X36 #= 0 ),
	( X36 #= 0 -> X37 #= 1 ; X37 #= 0 ),
	Y65 #= 0,
	Y66 #= Y65 + 25,
	Y67 #= Y66 * X37,
	Y68 #= Y67 + 1,
	Z27 #= Z26 * Y68,
	Y69 #= 0,
	Y70 #= Y69 + W12,
	Y71 #= Y70 + 14,
	Y72 #= Y71 * X37,
	Z28 #= Z27 + Y72,
	W13 #= M,
	X38 #= 0,
	X39 #= X38 + Z28,
	X40 #= X39 mod 26,
	Z29 #= Z28 div 26,
	X41 #= X40 + -16,
	( X41 #= W13 -> X42 #= 1 ; X42 #= 0 ),
	( X42 #= 0 -> X43 #= 1 ; X43 #= 0 ),
	Y73 #= 0,
	Y74 #= Y73 + 25,
	Y75 #= Y74 * X43,
	Y76 #= Y75 + 1,
	Z30 #= Z29 * Y76,
	Y77 #= 0,
	Y78 #= Y77 + W13,
	Y79 #= Y78 + 6,
	Y80 #= Y79 * X43,
	Z31 #= Z30 + Y80,
	W14 #= N,
	X44 #= 0,
	X45 #= X44 + Z31,
	X46 #= X45 mod 26,
	Z32 #= Z31 div 26,
	X47 #= X46 + -6,
	( X47 #= W14 -> X48 #= 1 ; X48 #= 0 ),
	( X48 #= 0 -> X49 #= 1 ; X49 #= 0 ),
	Y81 #= 0,
	Y82 #= Y81 + 25,
	Y83 #= Y82 * X49,
	Y84 #= Y83 + 1,
	Z33 #= Z32 * Y84,
	Y85 #= 0,
	Y86 #= Y85 + W14,
	Y87 #= Y86 + 5,
	Y88 #= Y87 * X49,
	Z34 #= Z33 + Y88,
	Z34 #= 0.
