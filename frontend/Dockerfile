# 개발용 빌드
FROM oven/bun:1.2.5 AS dev
WORKDIR /app
COPY package.json bun.lock ./
RUN bun install
COPY . .
EXPOSE 3000
CMD bun run dev --host

# 프로덕션 빌드
FROM oven/bun:1.2.5 AS builder
WORKDIR /app
COPY package.json bun.lock ./
RUN bun install
ARG DOMAIN
ENV DOMAIN=$DOMAIN
COPY . .
RUN bun run build

# 프로덕션 런타임
FROM oven/bun:1.2.5-slim
WORKDIR /app
COPY --from=builder /app/package.json ./
COPY --from=builder /app/dist ./dist
RUN bun add serve
EXPOSE 3000
CMD ["serve", "-s", "dist", "-l", "3000"]
