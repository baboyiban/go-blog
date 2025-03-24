# Certbot 자동 갱신 설정
while true; do
    # Certbot 갱신 테스트
    certbot renew --dry-run

    # 인증서 갱신 시도
    certbot renew --quiet

    # Nginx 재시작
    nginx -s reload

    # 12시간마다 갱신 시도
    sleep 12 * 60 * 60
done
