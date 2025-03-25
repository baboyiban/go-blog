#!/bin/sh

# 초기 대기 시간
sleep 30

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
