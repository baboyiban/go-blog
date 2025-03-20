# Stage 1: Build Go Backend
FROM golang:1.24-alpine3.21 AS go-builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-blog ./cmd/server/main.go

# Stage 2: Build Svelte Frontend
FROM node:20-alpine AS svelte-builder
WORKDIR /app
COPY web/svelte/package*.json ./web/svelte/
RUN cd web/svelte && npm install
COPY web/svelte ./web/svelte
RUN cd web/svelte && npm run build

# Stage 3: Runtime Image
FROM alpine:3.21
RUN apk add --no-cache ca-certificates tzdata
WORKDIR /app

# Copy Go Binary
COPY --from=go-builder /go-blog .

# Copy Svelte Build Output
COPY --from=svelte-builder /app/web/svelte/public ./web/static

EXPOSE 8080
ENTRYPOINT ["/app/go-blog"]
