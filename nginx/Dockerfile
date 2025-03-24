# 베이스 이미지: nginx 1.25 버전의 alpine 버전을 사용합니다.
FROM nginx:1.25-alpine

# Certbot 설치
RUN apk add --no-cache certbot

# 자체 서명된 인증서 및 설정 파일 복사
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Let's Encrypt 인증서 복사
COPY fullchain.pem /etc/nginx/certs/cert.pem
COPY privkey.pem /etc/nginx/certs/key.pem

# Certbot 자동 갱신 스크립트 복사
COPY certbot-renew.sh /usr/local/bin/certbot-renew.sh
RUN chmod +x /usr/local/bin/certbot-renew.sh

# 컨테이너의 80 포트 (HTTP)와 443 포트 (HTTPS)를 외부에 노출합니다.
EXPOSE 80 443

# Nginx와 Certbot 자동 갱신 스크립트 실행
CMD ["sh", "-c", "nginx -g 'daemon off;' & /usr/local/bin/certbot-renew.sh"]
