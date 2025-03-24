#!/bin/sh

# 환경 변수 불러오기 (상위 폴더의 .env 파일을 참조)
if [ -f ../.env ]; then
    export $(cat ../.env | xargs)
else
    echo "$(date): .env 파일을 찾을 수 없습니다."
    exit 1
fi

# 환경 변수가 설정되었는지 확인
if [ -z "$DOMAIN" ] || [ -z "$EMAIL" ]; then
    echo "$(date): DOMAIN 또는 EMAIL 환경 변수가 설정되지 않았습니다."
    exit 1
fi

# 초기 인증서 발급
if [ ! -f /etc/letsencrypt/live/${DOMAIN}/fullchain.pem ]; then
  certbot certonly --webroot --webroot-path /var/www/certbot --email ${EMAIL} --agree-tos --no-eff-email -d ${DOMAIN}
fi

# 인증서 갱신 스케줄링
while :; do
  certbot renew --quiet
  sleep 12h
done
