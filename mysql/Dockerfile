FROM mysql:8.0

# 환경 변수 설정
ENV MYSQL_ROOT_PASSWORD=my-secret-pw
ENV MYSQL_DATABASE=my_database
ENV MYSQL_USER=user
ENV MYSQL_PASSWORD=password

# 데이터베이스 초기화 스크립트 (필요한 경우)
COPY ./init.sql /docker-entrypoint-initdb.d/

# MySQL 포트 노출
EXPOSE 3306
