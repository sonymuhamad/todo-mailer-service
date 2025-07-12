# Stage 1: Build the binary
FROM golang:1.24.1-alpine AS builder

WORKDIR /app

# Install git for private modules (optional)
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

# Build binary
RUN go build -o app ./cmd/main.go

# Stage 2: Run the binary
FROM alpine:latest

WORKDIR /root/

# Install CA certificates (for HTTPS)
RUN apk --no-cache add ca-certificates

# Copy the binary from builder
COPY --from=builder /app/app .

# Copy any required files (optional)
COPY templates/ /templates

# Command to run
CMD ["./app"]
