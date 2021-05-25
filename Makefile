.PHONY: coverage test build

.DEFAULT_GOAL = build

build:
	bash hack/build-ui.sh

test:
	go test -v -race .

coverage:
	go test -v -coverprofile=coverage.txt -covermode=count .
	go tool cover --html=coverage.txt -o coverage.html

