FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o udp-relay main.go

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/udp-relay .
COPY .env .

ENV LOCAL_PORT=2454
ENV REMOTE_HOST=127.0.0.1
ENV REMOTE_PORT=24454

EXPOSE 24454/udp

ENTRYPOINT ["./udp-relay"]
