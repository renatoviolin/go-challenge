# build state
FROM golang:1.19 AS builder
WORKDIR /app
COPY ./src /app
COPY .env /app
RUN go version
RUN go test -v -cover ./application/... ./controllers/... ./repository/postgres/... ./infra/http-server/...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-w -s" -o main cmd/main.go

# Run stage
FROM gcr.io/distroless/static-debian11
COPY --from=builder /app/main /
COPY .env /

EXPOSE 8000
CMD [ "/main" ]

# docker build -t unico-app -f Dockerfile.app --no-cache . && docker run --name unico-app -p 8000:8000 --env-file .env -d unico-app