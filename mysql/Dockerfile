# 베이스 이미지: 공식 MySQL 8.0 이미지를 사용합니다.
FROM mysql:8.0

# 환경 변수 설정
# MYSQL_ROOT_PASSWORD: root 사용자의 비밀번호를 설정합니다.
ENV MYSQL_ROOT_PASSWORD=rootpassword

# MYSQL_DATABASE: 초기 데이터베이스를 생성합니다.
ENV MYSQL_DATABASE=godatabase

# MYSQL_USER: 초기 사용자를 생성합니다.
ENV MYSQL_USER=gouser

# MYSQL_PASSWORD: 초기 사용자의 비밀번호를 설정합니다.
ENV MYSQL_PASSWORD=gopassword

# 초기 SQL 스크립트 및 설정 파일 복사
# 초기 데이터베이스 스키마 또는 데이터를 포함하는 SQL 파일을 복사합니다.
COPY ./init.sql /docker-entrypoint-initdb.d/
# /docker-entrypoint-initdb.d/ 디렉토리에 있는 SQL 파일은 컨테이너가 처음 실행될 때 자동으로 실행됩니다.

# MySQL 설정 파일 복사 (선택 사항)
COPY ./my.cnf /etc/mysql/conf.d/
# my.cnf 파일은 MySQL의 사용자 정의 설정을 포함할 수 있습니다.

# MySQL의 기본 포트 (3306)를 외부에 노출합니다.
EXPOSE 3306

# MySQL 서버를 실행합니다.
CMD ["mysqld"]
