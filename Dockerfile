FROM golang:1.23.6-alpine AS builder

COPY . /app
WORKDIR /app

RUN go mod download
RUN go build -o ./bin/avitoMerch ./cmd/http_server/http_server.go

EXPOSE 8080

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/bin/avitoMerch .
COPY --from=builder /app/.env .

EXPOSE 8080

CMD ["./avitoMerch"]