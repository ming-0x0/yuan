# Development stage
FROM golang:1.23-alpine AS development

WORKDIR /app

# Install necessary tools for development
RUN apk add --no-cache git make gcc musl-dev

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Install air for hot reload (optional)
RUN go install github.com/air-verse/air@latest

# Expose application and debugger ports
EXPOSE 8080 2345

# Default command for development (can be overridden)
CMD ["go", "run", "cmd/standalone/main.go"]

# Builder stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git make gcc musl-dev

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o standalone cmd/standalone/main.go

# Production stage
FROM alpine:latest AS production

RUN apk --no-cache add ca-certificates wget

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/standalone .

# Copy configs
COPY configs ./configs

# Expose application port
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --quiet --tries=1 --spider http://localhost:8080/health || exit 1

# Run the application
CMD ["./standalone"]
