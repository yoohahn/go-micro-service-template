build:
	@go build -o bin/joke main.go

run: build
	./bin/joke

test:
	go test -v ./...
