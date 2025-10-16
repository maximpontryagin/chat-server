FROM golang:1.24.8-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o ./bin/chat_server cmd/main.go

FROM alpine:latest


WORKDIR /app
COPY --from=builder /app/bin/chat_server .
EXPOSE 50061

CMD ["./chat_server"]