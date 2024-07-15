# Use the official Golang image to create a build artifact
FROM golang:1.22 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod tidy

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
# RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY src/ ./src/

# Build the Go app
RUN go build -o main ./src/main.go

# Start a new stage from scratch
FROM mcr.microsoft.com/vscode/devcontainers/go:1

WORKDIR /root/

# install git
RUN apt-get update && apt-get install -y git

# Install sqlite3
# RUN apk add --no-cache sqlite-libs

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main /usr/local/bin

# Set executable perms
RUN chmod +x ./main

# Expose port 8080 to the outside world
# EXPOSE 8080

# Command to run the executable
CMD ["main"]
