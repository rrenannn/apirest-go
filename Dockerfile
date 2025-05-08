FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/main ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

RUN apk add --no-cache libc6-compat

COPY --from=builder /app/main /root/

RUN chmod +x /root/main

EXPOSE 8080

CMD ["./main"]