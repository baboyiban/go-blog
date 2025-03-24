FROM nginx:latest

# envsubst 설치
RUN apt-get update && apt-get install -y gettext-base

# 필요한 디렉토리 생성
RUN mkdir -p /etc/letsencrypt /var/www/certbot /etc/nginx/certs

# Nginx 설정 파일 복사
COPY ./nginx/conf.d/nginx.conf.template /etc/nginx/conf.d/nginx.conf.template

# 컨테이너의 80 포트 (HTTP)와 443 포트 (HTTPS)를 외부에 노출합니다.
EXPOSE 80 443

# Nginx 실행
CMD ["sh", "-c", "envsubst '$$DOMAIN $$EMAIL' < /etc/nginx/conf.d/nginx.conf.template > /etc/nginx/conf.d/nginx.conf && nginx -g 'daemon off;'"]
