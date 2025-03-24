# Build stage
# 빌드 단계: Go 애플리케이션을 빌드하는 단계
FROM golang:1.24-alpine AS builder
# golang:1.24-alpine 이미지를 기반으로 빌드 단계를 시작합니다.
# alpine 버전은 경량화된 Linux 배포판을 사용하여 이미지 크기를 줄입니다.
# AS builder는 이 단계를 'builder'라는 이름으로 명명합니다.

WORKDIR /app
# 작업 디렉토리를 /app으로 설정합니다. 이후 모든 명령은 이 디렉토리에서 실행됩니다.

# Install dependencies
# 의존성 설치: go.mod와 go.sum 파일을 복사하고 의존성을 다운로드합니다.
COPY go.* ./
# 호스트의 go.mod와 go.sum 파일을 컨테이너의 /app 디렉토리로 복사합니다.

RUN go mod download
# go mod download 명령어를 실행하여 Go 모듈 의존성을 다운로드합니다.

# Copy source code
# 소스 코드 복사: 모든 소스 코드를 컨테이너로 복사합니다.
COPY . .
# 호스트의 모든 파일을 컨테이너의 /app 디렉토리로 복사합니다.

# Build
# 애플리케이션 빌드: Go 애플리케이션을 빌드합니다.
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server/main.go
# CGO_ENABLED=0: CGO를 비활성화하여 정적 링크를 사용합니다.
# GOOS=linux: Linux용으로 빌드합니다.
# -o server: 빌드 결과물을 server라는 이름으로 저장합니다.
# ./cmd/server/main.go: 빌드할 메인 Go 파일의 경로입니다.

# Final stage
# 최종 단계: 빌드된 애플리케이션을 실행하는 단계
FROM alpine:3.19
# alpine:3.19 이미지를 기반으로 최종 단계를 시작합니다.
# alpine은 경량화된 Linux 배포판으로, 최종 이미지 크기를 줄이기 위해 사용됩니다.

RUN apk --no-cache add ca-certificates tzdata
# 필요한 패키지를 설치합니다.
# ca-certificates: SSL 인증서를 지원합니다.
# tzdata: 시간대 데이터를 지원합니다.
# --no-cache: 패키지 캐시를 저장하지 않아 이미지 크기를 줄입니다.

WORKDIR /app
# 작업 디렉토리를 /app으로 설정합니다.

COPY --from=builder /app/server .
# 빌드 단계에서 생성된 server 실행 파일을 최종 단계의 /app 디렉토리로 복사합니다.

# Create config directory and set environment
# config 디렉토리 생성 및 환경 변수 설정
RUN mkdir -p /app/config
# /app/config 디렉토리를 생성합니다. 설정 파일 등을 저장하기 위한 디렉토리입니다.

ENV GIN_MODE=release
# GIN_MODE 환경 변수를 release로 설정합니다.
# Gin 프레임워크가 release 모드로 동작하도록 합니다.

EXPOSE 8080
# 컨테이너의 8080 포트를 외부에 노출합니다.

CMD ["./server"]
# 컨테이너가 실행될 때 실행할 명령어를 지정합니다.
# ./server 명령어를 실행하여 애플리케이션을 시작합니다.
