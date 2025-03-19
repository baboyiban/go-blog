# Build stage
FROM golang:1.24-alpine3.21 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# 템플릿 복사 추가 (빌더 스테이지에서)
COPY web/templates ./web/templates
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-blog ./cmd/server/main.go

# Runtime stage
FROM alpine:3.21
WORKDIR /
COPY --from=builder /go-blog /go-blog
# 런타임에 템플릿 복사 (빌더 → 런타임)
COPY --from=builder /app/web/templates ./web/templates
# 정적 파일 복사 (빌더 → 런타임)
COPY --from=builder /app/web/static ./web/static

EXPOSE 8080
ENTRYPOINT ["/go-blog"]
