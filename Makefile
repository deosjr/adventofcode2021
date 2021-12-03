all:
	@for n in $$(seq -f "%02g" 1 3); do \
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
