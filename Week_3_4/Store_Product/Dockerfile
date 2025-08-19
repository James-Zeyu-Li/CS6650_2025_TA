FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY src ./src
COPY internal ./internal

# 指定主入口文件编译
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o server src/server/main.go

FROM alpine:latest
RUN apk add --no-cache ca-certificates

WORKDIR /app
COPY --from=builder /app/server .

RUN chmod +x server

EXPOSE 8080

CMD ["./server"]