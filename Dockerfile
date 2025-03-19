# Build stage
FROM golang:1.24-alpine3.21 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

# 전체 소스 복사 (템플릿 포함)
COPY . .

# 빌드 설정
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-blog ./cmd/server/main.go

# Runtime stage
FROM alpine:3.21

# 파일 감시 도구 설치 (inotify-tools)
RUN apk add --no-cache ca-certificates tzdata inotify-tools

WORKDIR /app

# 바이너리 복사
COPY --from=builder /go-blog .

# 템플릿 & 정적 파일 복사 (상대 경로 유지)
COPY --from=builder /app/web/templates ./web/templates
COPY --from=builder /app/web/static ./web/static

EXPOSE 8080
ENTRYPOINT ["/app/go-blog"]
