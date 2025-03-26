# 개발 환경용 (소스 마운트 + 핫 리로드)
FROM oven/bun:1.2.5

WORKDIR /app
COPY package.json bun.lock ./
RUN bun install  # 의존성 설치

EXPOSE 3000
CMD ["bun", "run", "dev", "--host"]  # 개발 서버 실행
