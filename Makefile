.PHONY: coverage test

.DEFAULT_GOAL = test

test:
	go test -v -race .

coverage:
	go test -v -coverprofile=coverage.txt -covermode=count .
	go tool cover --html=coverage.txt -o coverage.html

