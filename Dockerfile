# Stage 1: Build the Go binary
FROM golang:1.22 AS builder

# Set working directory inside container
WORKDIR /api

# Copy go.mod and go.sum first (better caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go binary
RUN go build -o main .

# Stage 2: Create a lightweight image
FROM alpine:latest

# Set working directory inside container
WORKDIR /api

# Copy binary from builder stage
COPY --from=builder /api/main .

# Expose port (adjust if your app uses a different one)
EXPOSE 8080

# Run the binary
CMD ["./main"]