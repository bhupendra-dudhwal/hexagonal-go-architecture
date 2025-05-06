# Stage 1: Build the Go binary
FROM golang:1.24.2-alpine AS builder

WORKDIR /app

# Install necessary tools
RUN apk add --no-cache git

# Copy go.mod and go.sum for dependency caching
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the entire source code
COPY . .

# Build the Go app
RUN go build -o main ./cmd/http

# Stage 2: Minimal image
FROM alpine:latest

# ✅ Install timezone data
RUN apk add --no-cache tzdata

WORKDIR /app

# Copy only the compiled binary from builder
COPY --from=builder /app/main .

# Expose application port
EXPOSE 8080

# Run the application
CMD ["./main"]