# === 1. 빌드 단계 ===
FROM oven/bun:1.2.5 AS builder

WORKDIR /app
COPY package.json bun.lock ./
# 정확한 버전 강제
RUN bun install --frozen-lockfile

# 빌드 시 전달되는 변수 (예: API 주소)
ARG DOMAIN
# Vite 환경 변수 예시
ENV DOMAIN=${DOMAIN}

COPY . .
RUN bun run build  # 프로덕션 빌드 (dist 생성)

# === 2. 런타임 단계 ===
# 경량 이미지
FROM oven/bun:1.2.5-slim

WORKDIR /app
COPY --from=builder /app/package.json ./
# 빌드 결과물만 복사
COPY --from=builder /app/dist ./dist
# 정적 파일 서버 설치
RUN bun add serve

EXPOSE 3000
# 빌드된 파일 서빙
CMD ["serve", "-s", "dist", "-l", "3000"]
