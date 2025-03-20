# Dockerfile
# Stage 1: Build Frontend
FROM oven/bun:1.2.5 as frontend-builder
WORKDIR /app/frontend
COPY frontend/package.json frontend/bun.lockb ./
RUN bun install
COPY frontend .
RUN bun run build

# Stage 2: Build Backend
FROM golang:1.24-alpine as backend-builder
WORKDIR /app
COPY backend/go.* ./
RUN go mod download
COPY backend .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server ./cmd/server/main.go

# Stage 3: Final Runtime
FROM alpine:3.21
RUN apk add --no-cache ca-certificates tzdata
WORKDIR /app

# Copy backend binary
COPY --from=backend-builder /app/server ./

# Copy frontend build output
COPY --from=frontend-builder /app/frontend/build /app/web/static

EXPOSE 8080
CMD ["/app/server"]
