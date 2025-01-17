# Use the official Golang image for building the app
FROM golang:1.23.3 AS builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download and cache dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o ping-redis .

# Use a minimal image for running the application
FROM alpine:latest

# Install certificates for HTTPS connections (optional, for TLS-enabled Redis)
RUN apk add --no-cache ca-certificates

# Set the working directory
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/ping-redis .

# Expose the default port (optional, if you plan to make this an API server later)
EXPOSE 8080

# Run the application
CMD ["./ping-redis"]
