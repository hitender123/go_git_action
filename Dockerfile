# Stage 1: Build the Go binary
FROM golang:1.22 AS builder

WORKDIR /api

# Copy go.mod and go.sum first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build statically linked binary (smaller size, no CGO dependencies)
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main .

# Stage 2: Create lightweight runtime image
FROM alpine:latest AS runner

WORKDIR /api

# Install certificates and timezone data for HTTPS and time handling
RUN apk --no-cache add ca-certificates tzdata

# Copy binary from builder stage
COPY --from=builder /api/main .

# Expose application port
EXPOSE 8080

# Run the binary
CMD ["./main"]