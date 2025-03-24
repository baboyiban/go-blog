## Let's Encrypt 인증서 발급
Let's Encrypt에서 SSL 인증서를 발급받기 위해 `certbot`을 사용합니다.

### Certbot 설치 및 인증서 발급:
1. Certbot을 설치합니다.
   ```bash
   sudo apt update
   sudo apt install certbot
   ```

2. 인증서를 발급받습니다.
   ```bash
   sudo certbot certonly --standalone -d example.com
   ```
   - `example.com`을 실제 도메인으로 변경하세요.
   - 이 명령어는 도메인의 80번 포트를 사용하여 인증서를 발급받습니다.

3. 발급된 인증서는 `/etc/letsencrypt/live/example.com/` 경로에 저장됩니다.

### Certbot 자동 갱신 설정:
```bash
# 테스트 및 자동생성
sudo certbot renew --dry-run
# 확인
cat /etc/cron.d/certbot
```

### 파일을 빌드 컨텍스트 디렉토리로 복사
`/etc/letsencrypt/live/example.com/privkey.pem` 파일이 빌드 컨텍스트 디렉토리(현재 작업 디렉토리)에 없으면 오류가 발생합니다.
- Docker 빌드 컨텍스트 디렉토리(현재 작업 디렉토리)로 SSL 인증서 파일을 복사합니다.
- 예를 들어, 현재 작업 디렉토리(`~/server/go-blog/nginx`)로 파일을 복사합니다.

```bash
sudo cp /etc/letsencrypt/live/example.com/privkey.pem .
sudo cp /etc/letsencrypt/live/example.com/fullchain.pem .
```
