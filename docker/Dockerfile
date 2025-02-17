FROM golang:1.23.4-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o ./tmp/main ./cmd/app/.

FROM alpine:latest

WORKDIR /app

COPY --from=builder app/tmp/main .

COPY entrypoint.sh .

COPY .env .env

RUN chmod +x entrypoint.sh
