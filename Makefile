.PHONY: run format test

run:
	go run .

format:
	gofmt -s -w .

test:
	go test -v ./...
