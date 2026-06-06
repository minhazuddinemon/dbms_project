FROM golang:1.26-alpine
WORKDIR /app

# 1. Install the Air hot-reloader globally inside the container
RUN go install github.com/air-verse/air@latest

# 2. Cache your dependencies
COPY go.mod go.sum* ./
RUN go mod download

# 3. Copy source code
COPY . .

EXPOSE 8080

# 4. Change the startup command from your compiled binary to Air
CMD ["air"]