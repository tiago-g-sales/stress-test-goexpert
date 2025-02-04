FROM golang:latest as builder
WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build --ldflags="-w -s" -o stress-test-cli main.go

FROM alpine:latest
COPY --from=builder /app/stress-test-cli .



CMD ["./stress-test-cli"]