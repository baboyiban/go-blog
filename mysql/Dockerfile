FROM mysql:8.0-debian

# 기본 설정 오버라이드
COPY ./my.cnf /etc/mysql/conf.d/my.cnf
RUN chmod 644 /etc/mysql/conf.d/my.cnf

RUN apt-get update && apt-get install -y tzdata && rm -rf /var/lib/apt/lists/*

COPY ./init.sql /docker-entrypoint-initdb.d/
