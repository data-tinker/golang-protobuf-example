# Build stage
FROM golang:1.21 AS builder

RUN apt-get update && apt-get install -y protobuf-compiler

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN protoc --go_out=. *.proto

RUN go build -o main

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main /app/main

CMD ["/app/main"]
