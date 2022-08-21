# build state
FROM golang:1.19-alpine AS builder
WORKDIR /app
COPY ./src /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-w -s" -o main cmd/main.go

# Run stage
FROM gcr.io/distroless/static-debian11
COPY --from=builder /app/main /
COPY .env /

EXPOSE 8000
CMD [ "/main" ]