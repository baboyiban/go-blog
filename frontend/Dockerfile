# Build stage
FROM oven/bun:1.2.5 AS builder
WORKDIR /app

# Install dependencies
COPY package.json bun.lock ./
RUN bun install --frozen-lockfile

# Copy source code
COPY . .

# Build application
RUN bun run build || (echo "Build failed"; exit 1)

# Verify build directory
RUN ls -la /app

# Runtime stage
FROM oven/bun:1.2.5-slim
WORKDIR /app

# Copy package.json and built assets
COPY --from=builder /app/package.json ./
COPY --from=builder /app/dist ./dist

# Install serve package
RUN bun add serve

EXPOSE 3000
CMD ["bun", "run", "start"]
