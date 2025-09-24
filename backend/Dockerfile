FROM golang:1.24.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o subscriptions-service ./cmd

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/subscriptions-service /app/subscriptions-service

COPY --from=builder /app/.env /app/.env

EXPOSE 8080

CMD ["/app/subscriptions-service"]