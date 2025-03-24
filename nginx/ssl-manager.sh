#!/bin/sh

# 초기 대기 시간
sleep 30

# 도메인 설정 (실제 도메인으로 변경 필요)
DOMAIN="choidaruhan.xyz"
EMAIL="chl11wq12@gmail.com"

# 인증서 없으면 초기 발급
if [ ! -d "/etc/letsencrypt/live/$DOMAIN" ]; then
    echo "$(date): 인증서 초기 발급 시도..."
    certbot certonly --webroot -w /var/www/certbot \
        -d $DOMAIN --email $EMAIL --agree-tos --non-interactive

    # 인증서 발급 성공 확인 및 Nginx 재시작
    if [ $? -eq 0 ]; then
        echo "$(date): 인증서 발급 성공! Nginx 재시작..."
        nginx -s reload
    else
        echo "$(date): 인증서 발급 실패. 로그를 확인하세요."
    fi
fi

# 정기적인 갱신 체크
while true; do
    echo "$(date): 인증서 갱신 확인..."
    certbot renew --quiet --webroot -w /var/www/certbot --deploy-hook "nginx -s reload"

    echo "$(date): 다음 갱신 체크는 12시간 후..."
    sleep 12h
done
