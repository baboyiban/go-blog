# 사용 방법

## 개발 환경
1. 개발 환경 실행:
```bash
docker-compose up --build
```

2. 프로덕션 환경 실행:
```bash
# 애플리케이션 빌드
docker-compose -f docker-compose.prod.yml up --build
# 코드를 수정한 후 재빌드 및 재실행
docker-compose -f docker-compose.prod.yml up -d --build
# 애플리케이션을 중지
docker-compose -f docker-compose.prod.yml down
# 볼륨까지 삭제하며 중지
docker-compose -f docker-compose.prod.yml down -v
```

3. 개별 서비스 빌드:
```bash
# 백엔드만 빌드
docker-compose build backend
# 프론트엔드만 빌드
docker-compose build frontend
```

4. 로그 확인
```bash
# 전체 로그 확인:
docker-compose -f docker-compose.prod.yml logs -f

# 개별 서비스 로그 확인:
## 프론트엔드:
docker-compose -f docker-compose.prod.yml logs -f frontend
## 백엔드:
docker-compose -f docker-compose.prod.yml logs -f backend
## Nginx:
docker-compose -f docker-compose.prod.yml logs -f nginx
```

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
sudo certbot renew --dry-run
```

#### 자체 서명된 SSL 인증서 생성:
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

2. 생성된 파일을 Nginx 설정에 사용합니다.

### 인증서 발급 명령어:
```bash
docker-compose -f docker-compose.prod.yml run --rm certbot certonly --webroot --webroot-path /var/www/certbot -d choidaruhan.xyz
```

### Nginx 재시작
인증서가 발급된 후 Nginx를 재시작하여 SSL 설정을 적용합니다.
```bash
docker-compose -f docker-compose.prod.yml restart nginx
```

### 갱신 테스트
갱신이 정상적으로 동작하는지 테스트하려면 다음 명령어를 실행하세요:
```bash
docker-compose -f docker-compose.prod.yml run --rm certbot certbot renew --dry-run
```

## 포트 설정
```bash
sudo ufw allow 80
sudo ufw allow 443
sudo ufw allow 22
sudo ufw enable
```

#### Linux (curl 사용)
```bash
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube
```
