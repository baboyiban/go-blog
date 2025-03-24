# 베이스 이미지: nginx 1.25 버전의 alpine 버전을 사용합니다.
# alpine 버전은 경량화된 Linux 배포판을 사용하여 이미지 크기를 줄입니다.
FROM nginx:1.25-alpine

# 자체 서명된 인증서 및 설정 파일 복사
# nginx 설정 파일을 컨테이너의 Nginx 설정 디렉토리로 복사합니다.
COPY nginx.conf /etc/nginx/conf.d/default.conf
# nginx.conf 파일은 Nginx의 기본 설정 파일로, 이 파일을 덮어쓰면 사용자 정의 설정이 적용됩니다.

# 자체 서명된 SSL 인증서와 키 파일을 컨테이너의 Nginx 인증서 디렉토리로 복사합니다.
COPY certs/cert.pem /etc/nginx/certs/cert.pem
COPY certs/key.pem /etc/nginx/certs/key.pem
# cert.pem: SSL 인증서 파일
# key.pem: SSL 키 파일
# 이 파일들은 HTTPS 통신을 위해 사용됩니다.

# 컨테이너의 80 포트 (HTTP)와 443 포트 (HTTPS)를 외부에 노출합니다.
EXPOSE 80 443

# Nginx를 실행하는 명령어
# nginx -g "daemon off;": Nginx를 포그라운드에서 실행합니다.
# daemon off; 옵션은 Nginx가 백그라운드로 실행되지 않도록 합니다.
# 이는 Docker 컨테이너가 실행 중인 동안 Nginx가 계속 실행되도록 하기 위함입니다.
CMD ["nginx", "-g", "daemon off;"]
