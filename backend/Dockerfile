FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod tidy

COPY main.go ./

RUN go build -o backend

FROM alpine:latest
WORKDIR /root/

COPY --from=builder /app/backend .

EXPOSE 3000

CMD ["./backend"]
