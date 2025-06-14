# Stage 1: Build aplikasi
FROM golang:1.21-alpine as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o server ./cmd/main.go

FROM alpine:3.21.3
WORKDIR /app

COPY --from=builder /app/server .
RUN chmod +x server

EXPOSE 8080
CMD ["./server"]