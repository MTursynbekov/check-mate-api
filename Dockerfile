FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /app/check-mate ./cmd/main

FROM alpine:3.9
WORKDIR /root/
COPY --from=builder /app .
EXPOSE 8080
CMD ./chek-mate
