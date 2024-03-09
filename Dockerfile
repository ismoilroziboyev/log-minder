# Build stage
FROM golang:1.21.5-alpine3.18 AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o log-minder cmd/main.go

# Run stage
FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/log-minder .

ENTRYPOINT [ "/app/log-minder" ]