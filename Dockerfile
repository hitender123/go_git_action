# -------- Build Stage --------
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Install git for modules
RUN apk add --no-cache git

# Copy go mod first for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .

# -------- Runtime Stage --------
FROM alpine:3.19

WORKDIR /app

RUN apk add --no-cache ca-certificates

# Copy binary from builder
COPY --from=builder /app/app .

# Expose port if API
EXPOSE 8080

CMD ["./app"]