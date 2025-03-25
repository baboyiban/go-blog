<<<<<<< HEAD
FROM mysql:8.0-debian

RUN apt-get update && apt-get install -y tzdata && rm -rf /var/lib/apt/lists/*

# 초기화 스크립트 복사
=======
# go-blog/mysql/Dockerfile
FROM mysql:8.0

# 데이터베이스 초기화 스크립트 복사
>>>>>>> main
COPY ./init.sql /docker-entrypoint-initdb.d/
