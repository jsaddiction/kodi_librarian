# syntax=docker/dockerfile:1

# Build Stage
FROM golang:1.22.5 as builder
WORKDIR /app
COPY ../go.mod ../go.sum ./
RUN go mod download
COPY ../src ./src
WORKDIR /app/src
RUN go build -o kodi_librarian .

# Final Stage
FROM mcr.microsoft.com/vscode/devcontainers/go:1
RUN apt-get update && apt-get install -y git
WORKDIR /root/
COPY --from=builder /app/src/kodi_librarian .
CMD ["./kodi_librarian"]
