all:
	@for n in $$(seq -f "%02g" 1 22); do \
		echo "$$n\n--------"; \
		echo "Go"; \
		\time go run $$n/day$$n.go; \
		echo "";\
	done

go1:
	@go run 01/day01.go

go2:
	@go run 02/day02.go
	
go3:
	@go run 03/day03.go

go4:
	@go run 04/day04.go

go5:
	@go run 05/day05.go

go6:
	@go run 06/day06.go

go7:
	@go run 07/day07.go

go8:
	@go run 08/day08.go

go9:
	@go run 09/day09.go

go10:
	@go run 10/day10.go

go11:
	@go run 11/day11.go

go12:
	@go run 12/day12.go

go13:
	@go run 13/day13.go

go14:
	@go run 14/day14.go

go15:
	@go run 15/day15.go

go16:
	@go run 16/day16.go

go17:
	@go run 17/day17.go

go18:
	@go run 18/day18.go

go19:
	@go run 19/day19.go

go20:
	@go run 20/day20.go

go21:
	@go run 21/day21.go

go22:
	@go run 22/day22.go

go23:
	@go run 23/day23.go

go24:
	@go run 24/day24.go

go25:
	@go run 25/day25.go
