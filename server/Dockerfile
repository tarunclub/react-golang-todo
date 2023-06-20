# Use the official Golang base image
FROM golang:1.16-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download and install Go dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Expose port 8000
EXPOSE 8000

# Set the entry point for the container
CMD ["./main"]
