run:
	go run src/cmd/main.go

build:
	go build -o main src/cmd/main.go 

test:
	cd src && go test -v -cover ./...
