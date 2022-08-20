run:
	cd src && go run cmd/main.go

build:
	cd src && go build -o main cmd/main.go 

test:
	cd src && go test -v -cover ./...
