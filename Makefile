run:
	go run src/cmd/main.go

build:
	go build -o main src/cmd/main.go 

test:
	go test -v -cover src/...
