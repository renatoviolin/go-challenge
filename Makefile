run-dev:
	cd src && go run cmd/main.go

build-dev:
	cd src && go build -o main cmd/main.go 

test-dev:
	cd src && go clean -testcache && go test -v -cover ./...

make run-db:
	docker rm -f go-challenge-db && docker build -t go-challenge-db -f Dockerfile.postgres --no-cache . && docker run --name go-challenge-db -p 5432:5432 --env-file .env -d go-challenge-db

make run-app:
	docker rm -f go-challenge-app && docker build -t go-challenge-app -f Dockerfile.app --no-cache . && docker run --name go-challenge-app -p 8000:8000 --env-file .env -d go-challenge-app && docker logs go-challenge-app --follow