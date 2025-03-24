### 초기 인증서 발급 (수동으로 진행)
초기 인증서 발급은 수동으로 진행해야 합니다. 아래 명령어를 실행하여 인증서를 발급받습니다.

```sh
docker-compose up -d
docker exec -it go-blog-certbot certbot certonly --webroot --webroot-path /var/www/certbot --email ${EMAIL} --agree-tos --no-eff-email -d ${DOMAIN}
```

- `${EMAIL}`과 `${DOMAIN}`은 `.env` 파일에 정의된 값을 사용합니다.
- 인증서가 정상적으로 발급되면, `/etc/letsencrypt/live/${DOMAIN}/` 경로에 인증서 파일이 생성됩니다.

---

### Nginx 재시작
인증서 발급 후, Nginx를 재시작하여 HTTPS 설정을 적용합니다.

```sh
docker-compose restart nginx
```

---

### 자동 갱신 확인
`certbot` 컨테이너는 12시간마다 인증서를 자동으로 갱신합니다. 갱신 로그는 아래 명령어로 확인할 수 있습니다.

```sh
docker logs go-blog-certbot
```

---

### 결과 확인
Nginx 컨테이너 내부에서 생성된 `nginx.conf` 파일을 확인하여 변수가 정상적으로 대체되었는지 확인합니다.

```sh
docker exec -it go-blog-nginx cat /etc/nginx/conf.d/nginx.conf
```

---

### 요약
- 초기 인증서 발급은 수동으로 진행 (`docker exec` 명령어 사용).
- 갱신은 `certbot` 컨테이너가 자동으로 처리.
- Nginx는 인증서가 발급된 후 HTTPS 설정을 적용.

이렇게 설정하면 초기 인증서 발급은 수동으로, 갱신은 자동으로 처리할 수 있습니다.
