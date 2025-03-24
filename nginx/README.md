### 초기 인증서 발급
1. Nginx 컨테이너를 실행합니다.
```bash
docker-compose up -d nginx
```

2. Certbot으로 초기 인증서를 발급합니다.
```bash
docker-compose run --rm certbot certonly --webroot --webroot-path /var/www/certbot --email your-email@example.com --agree-tos --no-eff-email -d yourdomain.com
```

3. Nginx 컨테이너를 재시작하여 SSL 설정을 적용합니다.
```bash
docker-compose restart nginx
```
