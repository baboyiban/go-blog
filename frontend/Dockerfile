# Build stage
# 빌드 단계: 애플리케이션을 빌드하는 단계
FROM oven/bun:1.2.5 AS builder
# oven/bun:1.2.5 이미지를 기반으로 빌드 단계를 시작합니다.
# AS builder는 이 단계를 'builder'라는 이름으로 명명합니다.

WORKDIR /app
# 작업 디렉토리를 /app으로 설정합니다. 이후 모든 명령은 이 디렉토리에서 실행됩니다.

# Install dependencies
# 의존성 설치: package.json과 bun.lock 파일을 복사하고 의존성을 설치합니다.
COPY package.json bun.lock ./
# 호스트의 package.json과 bun.lock 파일을 컨테이너의 /app 디렉토리로 복사합니다.

RUN bun install
# bun install 명령어를 실행하여 의존성을 설치합니다.

# Copy source code
# 소스 코드 복사: 모든 소스 코드를 컨테이너로 복사합니다.
COPY . .
# 호스트의 모든 파일을 컨테이너의 /app 디렉토리로 복사합니다.

# Build application
# 애플리케이션 빌드: 소스 코드를 빌드합니다.
RUN bun run build || (echo "bun run build failed"; exit 1)
# bun run build 명령어를 실행하여 애플리케이션을 빌드합니다.
# 빌드가 실패하면 "bun run build failed" 메시지를 출력하고 종료합니다.

# Verify build directory
# 빌드 디렉토리 확인: 빌드된 파일이 제대로 생성되었는지 확인합니다.
RUN ls -la /app
# /app 디렉토리의 파일 목록을 출력하여 빌드 결과를 확인합니다.

# Runtime stage
# 런타임 단계: 빌드된 애플리케이션을 실행하는 단계
FROM oven/bun:1.2.5-slim
# oven/bun:1.2.5-slim 이미지를 기반으로 런타임 단계를 시작합니다.
# slim 버전은 불필요한 파일을 제거하여 경량화된 이미지입니다.

WORKDIR /app
# 작업 디렉토리를 /app으로 설정합니다.

# Copy package.json and built assets
# package.json과 빌드된 파일 복사: 빌드 단계에서 생성된 파일을 런타임 단계로 복사합니다.
COPY --from=builder /app/package.json ./
# 빌드 단계에서 생성된 package.json 파일을 런타임 단계의 /app 디렉토리로 복사합니다.

COPY --from=builder /app/dist ./dist
# 빌드 단계에서 생성된 dist 디렉토리를 런타임 단계의 /app/dist 디렉토리로 복사합니다.

# Install serve package
# serve 패키지 설치: 정적 파일을 서빙하기 위해 serve 패키지를 설치합니다.
RUN bun add serve
# bun add serve 명령어를 실행하여 serve 패키지를 설치합니다.

EXPOSE 3000
# 컨테이너의 3000 포트를 외부에 노출합니다.

CMD ["bun", "run", "start"]
# 컨테이너가 실행될 때 실행할 명령어를 지정합니다.
# bun run start 명령어를 실행하여 애플리케이션을 시작합니다.
