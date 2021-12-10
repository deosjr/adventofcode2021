all:
	@for n in $$(seq -f "%02g" 1 10); do \
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
