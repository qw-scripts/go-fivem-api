# Builder stage
FROM golang:1.22.2 AS builder

WORKDIR /app

# Install build dependencies for CGO (for Debian-based systems)
RUN apt-get update && apt-get install -y gcc libc6-dev

# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code and build
COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -o go-fivem-api ./cmd/server/main.go

# Final stage
FROM debian:latest

# Set timezone
ENV TZ=UTC
RUN apt-get update && apt-get install -y tzdata && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

# Copy the binary from the builder stage
COPY --from=builder /app/go-fivem-api /app/go-fivem-api

# Copy the database file into the container
COPY test.db /test.db

# Set permissions for the database file, this is gross and should not exist
RUN chmod 777 /test.db

# Set the working directory to /app
WORKDIR /app

# Expose and run the binary
EXPOSE $PORT
CMD ["/app/go-fivem-api"]

# Health check (replace with your health check endpoint)
HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 CMD wget -qO- http://localhost:$PORT/ || exit 1
