all:
	@for n in $$(seq -f "%02g" 1 2); do \
		echo "$$n\n--------"; \
		echo "Go"; \
		\time go run $$n/day$$n.go; \
		echo "";\
	done

go1:
	@go run 01/day01.go

go2:
	@go run 02/day02.go
