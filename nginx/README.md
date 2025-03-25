### Let's Encrypt 인증서 발급
Let's Encrypt에서 SSL 인증서를 발급받기 위해 `certbot`을 사용합니다.

#### Certbot 설치 및 인증서 발급:
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

#### 파일을 빌드 컨텍스트 디렉토리로 복사.
`/etc/letsencrypt/live/choidaruhan.xyz/privkey.pem` 파일이 빌드 컨텍스트 디렉토리(현재 작업 디렉토리)에 없으면 오류가 발생합니다.
- Docker 빌드 컨텍스트 디렉토리(현재 작업 디렉토리)로 SSL 인증서 파일을 복사합니다.
- 예를 들어, 현재 작업 디렉토리(`~/server/go-blog/nginx`)로 파일을 복사합니다.

```bash
sudo cp /etc/letsencrypt/live/choidaruhan.xyz/privkey.pem .
sudo cp /etc/letsencrypt/live/choidaruhan.xyz/fullchain.pem .
```

#### Certbot 자동 갱신 설정:
```bash
# 테스트 및 자동생성
sudo certbot renew --dry-run
# 확인
cat /etc/cron.d/certbot
```

### Docker 배포

#### 빌드:
```bash
docker build -t nginx .
```

#### 실행:
```bash
# Docker 컨테이너를 백그라운드에서 실행 (-d 옵션)
# 호스트의 80 포트를 컨테이너의 80 포트로 매핑 (-p 80:80)
# 호스트의 443 포트를 컨테이너의 443 포트로 매핑 (-p 443:443)
# 컨테이너 이름을 "nginx"로 설정 (--name nginx)
# 사용할 이미지는 "nginx" (nginx)
docker run -d -p 80:80 -p 443:443 --name nginx nginx
```

#### 삭제:
```bash
docker rm nginx
```
#### 이미지 삭제:
```bash
docker rmi nginx
```

#### 재빌드:
```bash
docker build -t nginx .
```

#### 재실행:
```bash
docker run -d -p 80:80 -p 443:443 --name nginx nginx
```
