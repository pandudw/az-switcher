FROM golang:1.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o azclt ./cmd/commands

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/azclt .

ENTRYPOINT ["./azclt"]
