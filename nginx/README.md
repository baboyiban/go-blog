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
   sudo certbot certonly --standalone -d choidaruhan.xyz
   ```
   - `choidaruhan.xyz`을 실제 도메인으로 변경하세요.
   - 이 명령어는 도메인의 80번 포트를 사용하여 인증서를 발급받습니다.

3. 발급된 인증서는 `/etc/letsencrypt/live/choidaruhan.xyz/` 경로에 저장됩니다.

### Certbot 자동 갱신 설정:
```bash
# 테스트
sudo certbot renew --dry-run

cat /etc/cron.d/certbot
```

### 자체 서명된 SSL 인증서 생성:
1. OpenSSL을 사용하여 인증서를 생성합니다.
   ```bash
   sudo apt install openssl
   mkdir -p certs
   cd certs
   openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -nodes -subj "/CN=localhost"
   ```
   - `key.pem`: 개인 키 파일
   - `cert.pem`: 인증서 파일
   - `-days 365`: 인증서 유효 기간 (1년)
