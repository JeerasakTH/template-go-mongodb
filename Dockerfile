FROM golang:1.19.7-alpine3.17 AS builder
WORKDIR /app
COPY . /app

RUN go build -o main main.go

# Build small image
FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8080

CMD ["/app/main"]