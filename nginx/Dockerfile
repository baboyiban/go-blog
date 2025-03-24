# 베이스 이미지: nginx 1.25 버전의 alpine 버전을 사용합니다.
FROM nginx:1.25-alpine

# 필요한 디렉토리 생성
RUN mkdir -p /etc/letsencrypt /var/www/certbot /etc/nginx/certs

# Nginx 설정 파일 복사
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Certbot 자동 갱신 스크립트 복사
COPY ssl-manager.sh /usr/local/bin/ssl-manager.sh
RUN chmod +x /usr/local/bin/ssl-manager.sh

# 컨테이너의 80 포트 (HTTP)와 443 포트 (HTTPS)를 외부에 노출합니다.
EXPOSE 80 443

# Nginx 실행
CMD ["nginx", "-g", "daemon off;"]
