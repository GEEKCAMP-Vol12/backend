# Start from the official Golang base image with Go 1.20
FROM golang:1.20-alpine

# Install Git (required for go mod download)
RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Navigate to the directory containing main.go
WORKDIR /app/cmd/server

# Build the Go app
RUN go build -o /app/main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/app/main"]

