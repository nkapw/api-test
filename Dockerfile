FROM golang:1.23.3-alpine AS builder

# Set working directory
WORKDIR /app

# Copy source code
COPY . .

# Build aplikasi
RUN go build -o hello-world-app main.go

# Tahap final dengan image minimal
FROM alpine:latest

# Install CA certificates untuk HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy binary dari builder stage
COPY --from=builder /app/hello-world-app .

# Set PORT environment variable
ENV PORT=8000

# Expose port
EXPOSE 8000

# Run aplikasi
CMD ["./hello-world-app"]
