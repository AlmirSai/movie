# builder stage
FROM golang:1.24.1 AS builder

WORKDIR /app

# Copy module files and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the remaining source code and build the binary
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server/main.go

# final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/server .

# Optionally copy the config file if your app requires it at runtime
COPY config/.env ./config/.env

EXPOSE 8080

CMD ["./server"]
