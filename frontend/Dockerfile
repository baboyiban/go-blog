ARG BUN_VERSION=1.2.6

# 개발 단계 (소스 마운트 + 핫 리로드)
FROM oven/bun:${BUN_VERSION} AS dev

WORKDIR /usr/src/app
COPY package.json ./
RUN bun install --no-optional

COPY . .
EXPOSE 4321
CMD ["bun", "run", "dev", "--host"]

# 빌드 단계
FROM oven/bun:${BUN_VERSION} AS builder

WORKDIR /usr/src/app
RUN rm -rf node_modules bun.lock
COPY package.json bun.lock ./
RUN bun install --frozen-lockfile --no-optional
COPY . .
RUN bun run build

# 런타임 단계
FROM oven/bun:${BUN_VERSION}-slim
WORKDIR /usr/src/app
COPY --from=builder /usr/src/app/package.json ./
RUN bun add serve
COPY --from=builder /usr/src/app/dist ./dist

EXPOSE 4321
CMD ["serve", "-s", "dist", "-l", "4321"]
