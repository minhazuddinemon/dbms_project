# Build Stage
FROM golang:1.26-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker layer caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire backend source code
COPY . .

# Build the Go application
# CGO_ENABLED=0 ensures a statically linked binary, perfect for Alpine
# -o main names the output binary "main"
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api/main.go

# Run Stage
FROM alpine:latest

# Add CA certificates (needed if your API makes external HTTPS requests) and tzdata for timezones
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

# Copy only the compiled binary from the builder stage
COPY --from=builder /app/main .
# Copy migration files if your Go app runs them on startup
COPY --from=builder /app/database/migrations ./database/migrations

# Expose the port your REST API will run on (adjust if your server uses a different port)
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
