ARG GO_VERSION=1.24.1

# 개발 빌드
FROM golang:${GO_VERSION}-alpine AS dev
WORKDIR /app
RUN go install github.com/air-verse/air@latest
COPY go.mod go.sum ./
RUN go mod download
COPY . .
EXPOSE 8080
CMD ["air", "-c", ".air.toml"]

# 프로덕션 빌드
FROM golang:${GO_VERSION}-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server/main.go

# 프로덕션 런타임
FROM alpine:3.19
RUN apk --no-cache add ca-certificates tzdata
ENV TZ=Asia/Seoul
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["./server"]
