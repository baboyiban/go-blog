# === 1. 빌드 단계 ===
FROM oven/bun:1.2.5 AS builder

WORKDIR /app
COPY package.json bun.lock ./
RUN bun install --frozen-lockfile

COPY . .
RUN bun run build  # 프로덕션 빌드 (dist 생성)

# === 2. 런타임 단계 ===
# 경량 이미지
FROM oven/bun:1.2.5-slim

WORKDIR /app
COPY --from=builder /app/package.json ./
# 런타임 단계에서 serve 설치
RUN bun add serve
# 빌드 결과물만 복사
COPY --from=builder /app/dist ./dist

EXPOSE 3000
# 빌드된 파일 서빙
CMD ["serve", "-s", "dist", "-l", "3000"]
