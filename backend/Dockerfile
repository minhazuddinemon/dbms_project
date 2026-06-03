FROM golang:1.26-alpine
WORKDIR /app

COPY go.mod go.sum* ./
RUN go mod download

COPY . .
# ─── UPDATE THIS LINE TO USE COMPILER CACHE ───
RUN --mount=type=cache,target=/root/.cache/go-build \
    go build -o server main.go

EXPOSE 8080
CMD [ "./server" ]
