# Use the official Golang image to build the application
FROM golang:1.22 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod go.sum ./

# Download the Go Modules. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Navigate to the cmd/my-go-api directory to build the application
WORKDIR /app/cmd/my-go-api

# Build the Go app
RUN go build -o /app/main .

# Start a new stage using the same base image to ensure compatibility
FROM golang:1.22

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Copy the .env file from the build context
COPY .env .

# Expose port 8081 to the outside world
EXPOSE 8081

# Command to run the executable
CMD ["./main"]
