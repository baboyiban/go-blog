# Build stage
FROM golang:1.24-alpine AS builder
WORKDIR /app

# Install dependencies
COPY go.* ./
RUN go mod download

# Copy source code
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server/main.go

# Final stage
FROM alpine:3.19
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app
COPY --from=builder /app/server .

# Create config directory and set environment
RUN mkdir -p /app/config
ENV GIN_MODE=release

EXPOSE 8080
CMD ["./server"]
