FROM golang:1.20-alpine3.16 AS builder 
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main /app/main.go

FROM alpine:3.16.0
WORKDIR /app
COPY --from=builder /app/main .
RUN apk add --no-cache bash curl
EXPOSE 3000
CMD ["/app/main"]
