FROM mysql:8.0

# 타임존 데이터 설치
RUN apt-get update && apt-get install -y tzdata && rm -rf /var/lib/apt/lists/*

# 초기화 스크립트 복사
COPY ./init.sql /docker-entrypoint-initdb.d/
