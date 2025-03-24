# go-blog/mysql/Dockerfile
FROM mysql:8.0

# 데이터베이스 초기화 스크립트 복사
COPY ./init.sql /docker-entrypoint-initdb.d/
