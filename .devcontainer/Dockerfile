# syntax=docker/dockerfile:1

# Build Stage
FROM golang:1.22.5 as builder

# Download packages
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy source files and build binary
COPY . .
WORKDIR /app/src

# Build binary
RUN go build -o /bin/kodi_librarian .

# Final Stage
FROM mcr.microsoft.com/vscode/devcontainers/go:1

# Install packages
RUN apt-get update && apt-get install -y git

# Copy binary
WORKDIR /bin
COPY --from=builder /bin/kodi_librarian .

# Copy all files into workspace
WORKDIR /app
COPY --from=builder /app .

# Run the application
CMD ["/bin/kodi_librarian"]
