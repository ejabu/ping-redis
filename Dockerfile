# Use the official Golang image
FROM golang:1.23.3

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

# Expose the default port (optional, if you plan to make this an API server later)
EXPOSE 8080

# Set the command to run the application
CMD ["./ping-redis"]
