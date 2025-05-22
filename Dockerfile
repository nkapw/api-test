# Stage 1: Build
FROM golang:1.23.5 AS builder

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

WORKDIR /app

COPY . .

RUN go build -o hello-world-app main.go

# Stage 2: Final image with Ubuntu (tanpa certs)
FROM ubuntu:22.04

WORKDIR /root/

# Salin binary dari builder stage
COPY --from=builder /app/hello-world-app .

EXPOSE 8000

CMD ["./hello-world-app"]
